package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"testing"

	"gm-session-manager/perm"
)

func TestDiff(t *testing.T) {
	tests := []struct {
		name       string
		prev, next State
		wantSet    State
		wantDel    []string
	}{
		{"unchanged", State{"hp": "5", "#": "a"}, State{"hp": "5", "#": "a"}, nil, nil},
		{"changed value", State{"hp": "5"}, State{"hp": "3"}, State{"hp": "3"}, nil},
		{"new key", State{"hp": "5"}, State{"hp": "5", "temp": "+2 temp"}, State{"temp": "+2 temp"}, nil},
		{"dropped keys sorted", State{"b": 1, "a": true, "c": 2}, State{"c": 2}, nil, []string{"a", "b"}},
		{"bool flip", State{"cond.prone": false}, State{"cond.prone": true}, State{"cond.prone": true}, nil},
		{"from nothing", nil, State{"x": 1}, State{"x": 1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set, del := Diff(tt.prev, tt.next)
			if !reflect.DeepEqual(set, tt.wantSet) || !slices.Equal(del, tt.wantDel) {
				t.Errorf("Diff = %v, %v; want %v, %v", set, del, tt.wantSet, tt.wantDel)
			}
		})
	}
}

// stateHandler serves states from a map; a missing view is not visible.
type stateHandler struct {
	states  map[string]State
	affects func(view string, topics []Topic) bool
}

func (h *stateHandler) Affects(view string, topics []Topic) bool { return h.affects(view, topics) }
func (h *stateHandler) State(_ context.Context, _ *Client, view string) (State, bool, error) {
	s, ok := h.states[view]
	return s, ok, nil
}
func (*stateHandler) CanSubscribe(context.Context, *Client, int64) bool { return true }

// manyKeys is a state with one more key than a patch may carry.
func manyKeys(v int) State {
	s := State{"#": "a"}
	for i := range maxPatchKeys + 1 {
		s[fmt.Sprintf("slot.%d", i)] = v
	}
	return s
}

func TestPatch(t *testing.T) {
	ch := func(id int64) Topic { return Topic{Kind: TopicCharacter, ID: id} }
	byID := func(view string, topics []Topic) bool {
		return slices.ContainsFunc(topics, func(t Topic) bool { return fmt.Sprintf("character:%d", t.ID) == view })
	}
	type step struct {
		states map[string]State
		msg    message
		want   []viewPatch // nil: no message at all
	}
	tests := []struct {
		name  string
		watch []string
		steps []step
	}{
		{
			name:  "hello baselines quietly, then only changes are sent",
			watch: []string{"character:1"},
			steps: []step{
				{map[string]State{"character:1": {"hp": "5", "#": "a"}}, message{hello: true}, []viewPatch{}},
				{map[string]State{"character:1": {"hp": "3", "#": "a"}}, message{topics: []Topic{ch(1)}}, []viewPatch{{View: "character:1", Set: State{"hp": "3"}}}},
				{map[string]State{"character:1": {"hp": "3", "#": "a"}}, message{topics: []Topic{ch(1)}}, nil},
			},
		},
		{
			name:  "a changed shape hash orders a refetch, not values",
			watch: []string{"character:1"},
			steps: []step{
				{map[string]State{"character:1": {"hp": "5", "#": "a"}}, message{hello: true}, []viewPatch{}},
				{map[string]State{"character:1": {"hp": "3", "#": "b"}}, message{topics: []Topic{ch(1)}}, []viewPatch{{View: "character:1", Refetch: true}}},
			},
		},
		{
			name:  "a diff over the key cap orders a refetch",
			watch: []string{"character:1"},
			steps: []step{
				{map[string]State{"character:1": manyKeys(0)}, message{hello: true}, []viewPatch{}},
				{map[string]State{"character:1": manyKeys(1)}, message{topics: []Topic{ch(1)}}, []viewPatch{{View: "character:1", Refetch: true}}},
			},
		},
		{
			name:  "views missed during catch-up are refetched",
			watch: []string{"character:1", "character:2"},
			steps: []step{
				{map[string]State{"character:1": {"#": "a"}, "character:2": {"#": "b"}}, message{hello: true, topics: []Topic{ch(2)}}, []viewPatch{{View: "character:2", Refetch: true}}},
			},
		},
		{
			name:  "unwatched and unaffected views are skipped",
			watch: []string{"character:1"},
			steps: []step{
				{map[string]State{"character:1": {"#": "a"}}, message{hello: true}, []viewPatch{}},
				{map[string]State{"character:1": {"#": "z"}, "character:2": {"#": "q"}}, message{topics: []Topic{ch(2)}}, nil},
			},
		},
		{
			name:  "a view becoming hidden or visible is refetched, never diffed",
			watch: []string{"character:1"},
			steps: []step{
				{map[string]State{"character:1": {"secret": "x"}}, message{hello: true}, []viewPatch{}},
				{map[string]State{}, message{topics: []Topic{ch(1)}}, []viewPatch{{View: "character:1", Refetch: true}}},
				{map[string]State{}, message{topics: []Topic{ch(1)}}, nil},
				{map[string]State{"character:1": {"secret": "y"}}, message{topics: []Topic{ch(1)}}, []viewPatch{{View: "character:1", Refetch: true}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &stateHandler{affects: byID}
			h := New(sh)
			c := NewClient(1, perm.GM, 0, "test")
			c.setWatch(tt.watch)
			for i, st := range tt.steps {
				sh.states = st.states
				raw, err := h.patch(t.Context(), c, st.msg)
				if err != nil {
					t.Fatal(err)
				}
				if st.want == nil {
					if raw != nil {
						t.Fatalf("step %d: sent %s, want nothing", i, raw)
					}
					continue
				}
				var got patchMessage
				if err := json.Unmarshal(raw, &got); err != nil {
					t.Fatal(err)
				}
				if got.Type != msgPatch || got.Epoch != h.epoch {
					t.Errorf("step %d: header %+v", i, got)
				}
				if len(got.Views) != len(st.want) {
					t.Fatalf("step %d: views %+v, want %+v", i, got.Views, st.want)
				}
				for j := range st.want {
					// JSON turns the State's values back into any; compare as JSON.
					w, _ := json.Marshal(st.want[j])
					g, _ := json.Marshal(got.Views[j])
					if string(w) != string(g) {
						t.Errorf("step %d: view %s, want %s", i, g, w)
					}
				}
			}
		})
	}
}
