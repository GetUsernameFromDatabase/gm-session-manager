package hub

import (
	"context"
	"encoding/json"
	"slices"
)

// State is a live view's data as one client may see it: flat keys to
// strings, numbers or booleans. Pages bind the keys that change often; any
// other key (by convention "#", a hash of the rest of the view) makes the
// browser refetch the view's HTML.
type State map[string]any

// HashKey holds the hash of everything in a view that is not bound.
const HashKey = "#"

// Diff returns the keys of next that are new or changed since prev, and the
// keys of prev that next dropped.
func Diff(prev, next State) (set State, del []string) {
	for k, v := range next {
		if old, ok := prev[k]; !ok || old != v {
			if set == nil {
				set = State{}
			}
			set[k] = v
		}
	}
	for k := range prev {
		if _, ok := next[k]; !ok {
			del = append(del, k)
		}
	}
	slices.Sort(del)
	return set, del
}

// viewPatch is one view's entry in a patch message. Refetch asks the browser
// to load the view's HTML, because the view appeared, went away or was
// missed while disconnected.
type viewPatch struct {
	View    string   `json:"v"`
	Set     State    `json:"set,omitempty"`
	Del     []string `json:"del,omitempty"`
	Refetch bool     `json:"refetch,omitempty"`
}

type patchMessage struct {
	Type  string      `json:"type"`
	Epoch string      `json:"epoch"`
	Seq   uint64      `json:"seq"`
	Views []viewPatch `json:"views,omitempty"`
}

const msgPatch = "patch"

// maxPatchKeys is the most changed keys one view's patch carries; beyond
// it (a long rest, a level-up) the page refetches the view instead.
const maxPatchKeys = 8

// patch builds the client's message for m: the changes, since what it last
// saw, of every watched view the topics touch. On hello every watched view
// is (re)baselined, and the ones that changed while the page was loading or
// the device was away are refetched.
func (h *Hub) patch(ctx context.Context, c *Client, m message) ([]byte, error) {
	var views []viewPatch
	for _, v := range c.watching() {
		if !m.hello && !h.handler.Affects(v, m.topics) {
			continue
		}
		next, ok, err := h.handler.State(ctx, c, v)
		if err != nil {
			return nil, err
		}
		prev, had := c.last[v]
		if !ok {
			next = nil
		}
		c.last[v] = next
		switch {
		case m.hello:
			if h.handler.Affects(v, m.topics) {
				views = append(views, viewPatch{View: v, Refetch: true})
			}
		case !had || (prev == nil) != (next == nil):
			views = append(views, viewPatch{View: v, Refetch: true})
		default:
			set, del := Diff(prev, next)
			_, reshaped := set[HashKey]
			switch {
			case set == nil && del == nil:
			case reshaped || len(set)+len(del) > maxPatchKeys:
				// A wholesale change is cheaper as one refetch order than as
				// values the page would refetch over anyway.
				views = append(views, viewPatch{View: v, Refetch: true})
			default:
				views = append(views, viewPatch{View: v, Set: set, Del: del})
			}
		}
	}
	if len(views) == 0 && !m.hello {
		return nil, nil
	}
	return json.Marshal(patchMessage{Type: msgPatch, Epoch: h.epoch, Seq: m.seq, Views: views})
}

// setWatch records the live views the client's page shows.
func (c *Client) setWatch(views []string) {
	c.subMu.Lock()
	defer c.subMu.Unlock()
	c.watch = slices.Compact(slices.Sorted(slices.Values(views)))
}

func (c *Client) watching() []string {
	c.subMu.Lock()
	defer c.subMu.Unlock()
	return slices.Clone(c.watch)
}
