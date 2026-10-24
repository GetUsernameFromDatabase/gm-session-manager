package hub

import (
	"context"
	"slices"
	"testing"

	"gm-session-manager/perm"
)

type fakeHandler struct{}

func (fakeHandler) Affects(string, []Topic) bool { return true }
func (fakeHandler) State(context.Context, *Client, string) (State, bool, error) {
	return State{}, true, nil
}
func (fakeHandler) CanSubscribe(context.Context, *Client, int64) bool { return true }

func drain(c *Client) []message {
	var out []message
	for {
		select {
		case m := <-c.out:
			out = append(out, m)
		default:
			return out
		}
	}
}

func topicIDs(m message) []int64 {
	var ids []int64
	for _, t := range m.topics {
		ids = append(ids, t.ID)
	}
	slices.Sort(ids)
	return ids
}

func TestCatchUp(t *testing.T) {
	ch := func(id int64) Topic { return Topic{Kind: TopicCharacter, ID: id} }
	tests := []struct {
		name string
		// publish runs before the client says hello.
		publish    func(h *Hub)
		helloEpoch func(h *Hub) string
		helloSeq   uint64
		wantReload bool
		wantTopics []int64
	}{
		{
			name:       "fresh page, nothing missed",
			publish:    func(*Hub) {},
			helloEpoch: func(h *Hub) string { return h.epoch },
			wantTopics: nil,
		},
		{
			name: "missed events in this campaign only, deduplicated",
			publish: func(h *Hub) {
				h.Publish(1, ch(10))
				h.Publish(2, ch(20)) // other campaign
				h.Publish(1, ch(11), ch(10))
			},
			helloEpoch: func(h *Hub) string { return h.epoch },
			wantTopics: []int64{10, 11},
		},
		{
			name: "only events after the client's seq",
			publish: func(h *Hub) {
				h.Publish(1, ch(10))
				h.Publish(1, ch(11))
			},
			helloEpoch: func(h *Hub) string { return h.epoch },
			helloSeq:   1,
			wantTopics: []int64{11},
		},
		{
			name:       "host restarted",
			publish:    func(*Hub) {},
			helloEpoch: func(*Hub) string { return "old-epoch" },
			wantReload: true,
		},
		{
			name: "gap older than the ring",
			publish: func(h *Hub) {
				for i := range ringSize + 5 {
					h.Publish(1, ch(int64(i)))
				}
			},
			helloEpoch: func(h *Hub) string { return h.epoch },
			helloSeq:   1,
			wantReload: true,
		},
		{
			name:       "seq from the future",
			publish:    func(*Hub) {},
			helloEpoch: func(h *Hub) string { return h.epoch },
			helloSeq:   99,
			wantReload: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(fakeHandler{})
			c := NewClient(1, perm.GM, 0, "test")
			h.register(c)
			tt.publish(h)
			if got := drain(c); len(got) != 0 {
				t.Fatalf("unsynced client received %d messages", len(got))
			}

			h.sync(c, tt.helloEpoch(h), tt.helloSeq)
			msgs := drain(c)
			if len(msgs) != 1 {
				t.Fatalf("got %d messages, want 1", len(msgs))
			}
			if gotReload := msgs[0].control == controlReload; gotReload != tt.wantReload {
				t.Fatalf("reload = %v, want %v", gotReload, tt.wantReload)
			}
			if !tt.wantReload {
				if ids := topicIDs(msgs[0]); !slices.Equal(ids, tt.wantTopics) {
					t.Errorf("topics %v, want %v", ids, tt.wantTopics)
				}
				_, seq := h.Position()
				if msgs[0].seq != seq {
					t.Errorf("position %d, want %d", msgs[0].seq, seq)
				}
			}

			// Once synced, live events arrive directly.
			h.Publish(1, ch(42))
			if live := drain(c); len(live) != 1 || topicIDs(live[0])[0] != 42 {
				t.Errorf("live delivery: %+v", live)
			}
		})
	}
}

func TestCloseSessionAndCharacter(t *testing.T) {
	h := New(fakeHandler{})
	alice := NewClient(1, perm.Player(7), 100, "phone")
	bob := NewClient(1, perm.Player(8), 200, "tablet")
	gm := NewClient(1, perm.GM, 0, "pc")
	all := []*Client{alice, bob, gm}
	for _, c := range all {
		h.register(c)
	}
	for _, c := range all {
		h.sync(c, h.epoch, 0)
		drain(c) // catch-up, including the device announcements
	}

	h.CloseSession(100)
	if m := drain(alice); len(m) != 1 || m[0].control != controlReload || !m[0].close {
		t.Errorf("alice: %+v", m)
	}
	if len(drain(bob))+len(drain(gm)) != 0 {
		t.Error("kick reached other devices")
	}
	h.CloseSession(0) // localhost admin has no session; must not match it
	if len(drain(gm)) != 0 {
		t.Error("session 0 matched the localhost admin")
	}
	h.CloseCharacter(8)
	if m := drain(bob); len(m) != 1 || !m[0].close {
		t.Errorf("bob: %+v", m)
	}
}

func TestPendingFlow(t *testing.T) {
	h := New(fakeHandler{})
	p := h.AddPending(Pending{CampaignID: 1, CampaignName: "Camp", Kind: PendingClaim, CharacterID: 5, Name: "Alice"})
	visitor := NewClient(1, perm.Viewer{}, 0, "phone")
	pp, ok := h.PendingBySecret(p.Secret)
	if !ok {
		t.Fatal("request not found by secret")
	}
	visitor.Pending = pp
	h.register(visitor)

	if got := h.PendingList(1); len(got) != 1 || got[0].ID != p.ID {
		t.Fatalf("list %+v", got)
	}
	if len(h.PendingList(2)) != 0 {
		t.Error("request listed in another campaign")
	}
	if _, ok := h.TakeApproved(p.Secret); ok {
		t.Error("redeemed before approval")
	}
	if h.Decide(2, p.ID, true) {
		t.Error("approved from another campaign")
	}
	if !h.Decide(1, p.ID, true) {
		t.Fatal("approve failed")
	}
	if h.Decide(1, p.ID, false) {
		t.Error("decided twice")
	}
	if len(h.PendingList(1)) != 0 {
		t.Error("approved request still listed as waiting")
	}
	if _, ok := h.TakeApproved("wrong"); ok {
		t.Error("redeemed with the wrong secret")
	}
	got, ok := h.TakeApproved(p.Secret)
	if !ok || got.CharacterID != 5 {
		t.Fatalf("redeem: %+v %v", got, ok)
	}
	if _, ok := h.TakeApproved(p.Secret); ok {
		t.Error("redeemed twice")
	}
}

func TestDisconnectDropsRequest(t *testing.T) {
	h := New(fakeHandler{})
	p := h.AddPending(Pending{CampaignID: 1, CampaignName: "Camp", Kind: PendingCreate, Name: "Newbie"})
	visitor := NewClient(1, perm.Viewer{}, 0, "phone")
	visitor.Pending, _ = h.PendingBySecret(p.Secret)
	h.register(visitor)
	h.unregister(visitor)
	if len(h.PendingList(1)) != 0 {
		t.Error("request survived its connection")
	}
}

func TestDeviceAnnouncements(t *testing.T) {
	h := New(fakeHandler{})
	gm := NewClient(1, perm.GM, 0, "pc")
	h.register(gm)
	h.sync(gm, h.epoch, 0)
	drain(gm)

	phone := NewClient(1, perm.Player(7), 100, "phone")
	h.register(phone)
	if m := drain(gm); len(m) != 1 || m[0].topics[0].Kind != TopicCampaign {
		t.Errorf("connect not announced: %+v", m)
	}
	h.unregister(phone)
	if m := drain(gm); len(m) != 1 {
		t.Errorf("disconnect not announced: %+v", m)
	}
	visitor := NewClient(1, perm.Viewer{}, 0, "phone")
	h.register(visitor)
	if m := drain(gm); len(m) != 0 {
		t.Errorf("visitor announced as a device: %+v", m)
	}
}

func TestDMGetsEveryCampaignsRequests(t *testing.T) {
	h := New(fakeHandler{})
	gm := NewClient(2, perm.GM, 0, "pc")
	player := NewClient(2, perm.Player(7), 100, "phone")
	for _, c := range []*Client{gm, player} {
		h.register(c)
		h.sync(c, h.epoch, 0)
	}
	drain(gm)
	drain(player)
	h.AddPending(Pending{CampaignID: 1, CampaignName: "Camp", Kind: PendingCreate, Name: "Newbie"})
	if m := drain(gm); len(m) != 1 || m[0].topics[0].Kind != TopicPending {
		t.Errorf("gm: %+v", m)
	}
	if m := drain(player); len(m) != 0 {
		t.Errorf("player in another campaign: %+v", m)
	}

	// A GM page loaded before the request catches up on it.
	late := NewClient(0, perm.GM, 0, "pc")
	h.register(late)
	h.sync(late, h.epoch, 0)
	if m := drain(late); len(m) != 1 || len(m[0].topics) != 1 {
		t.Errorf("catch-up: %+v", m)
	}
}
