package hub

import (
	"slices"
	"time"
)

// PendingKind is what a visitor asked for.
type PendingKind string

// Request kinds.
const (
	// PendingClaim asks to play an existing character.
	PendingClaim PendingKind = "claim"
	// PendingCreate asks to create a character in a campaign without
	// open_join.
	PendingCreate PendingKind = "create"
)

// PendingState is where a request stands.
type PendingState int

// Request states.
const (
	Waiting PendingState = iota
	Approved
	Denied
)

// attachTimeout is how long a request may wait for its visitor's websocket.
// Requests live only while the visitor is connected.
const attachTimeout = time.Minute

// Pending is a join request waiting for the GM. Requests live in memory
// only; a dropped connection drops the request.
type Pending struct {
	// ID is shown to the GM and used in approve/deny URLs.
	ID string
	// Secret is known only to the visitor and redeems an approval.
	Secret     string
	CampaignID int64
	// CampaignName is copied in so every GM page can name it without a query.
	CampaignName string
	Kind         PendingKind
	CharacterID  int64  // PendingClaim
	Name         string // character name (claimed or to create)
	// Player, OS, Phone and Host tell the GM who is asking.
	Player  string
	OS      string
	Phone   bool
	Host    string // LAN hostname; empty when the network does not know it
	Created time.Time
	State   PendingState

	attached bool
}

// AddPending records a request, giving it its ID, secret and time, and
// tells the GM.
func (h *Hub) AddPending(req Pending) Pending {
	p := &req
	p.ID, p.Secret, p.Created, p.State = randomID(8), randomID(16), time.Now(), Waiting
	h.mu.Lock()
	defer h.mu.Unlock()
	h.pending[p.ID] = p
	h.publishLocked(p.CampaignID, Topic{Kind: TopicPending, ID: p.CampaignID})
	return *p
}

// PendingList returns a campaign's waiting requests (every campaign's for
// 0), oldest first, dropping any whose visitor never connected.
func (h *Hub) PendingList(campaignID int64) []Pending {
	h.mu.Lock()
	defer h.mu.Unlock()
	var out []Pending
	for id, p := range h.pending {
		if !p.attached && time.Since(p.Created) > attachTimeout {
			delete(h.pending, id)
			continue
		}
		if (campaignID == 0 || p.CampaignID == campaignID) && p.State == Waiting {
			out = append(out, *p)
		}
	}
	slices.SortFunc(out, func(a, b Pending) int { return a.Created.Compare(b.Created) })
	return out
}

// PendingBySecret looks a request up for its visitor.
func (h *Hub) PendingBySecret(secret string) (*Pending, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, p := range h.pending {
		if secret != "" && p.Secret == secret {
			return p, true
		}
	}
	return nil, false
}

// PendingState returns a request's current state for rendering.
func (h *Hub) PendingState(p *Pending) PendingState {
	h.mu.Lock()
	defer h.mu.Unlock()
	return p.State
}

// Decide approves or denies a waiting request in a campaign and notifies
// both the GM and the visitor. It reports false if no such request waits.
func (h *Hub) Decide(campaignID int64, id string, approve bool) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	p, ok := h.pending[id]
	if !ok || p.CampaignID != campaignID || p.State != Waiting {
		return false
	}
	if approve {
		p.State = Approved
	} else {
		p.State = Denied
	}
	h.publishLocked(campaignID, Topic{Kind: TopicPending, ID: campaignID})
	return true
}

// TakeApproved removes and returns an approved request, so its visitor can
// redeem it exactly once.
func (h *Hub) TakeApproved(secret string) (Pending, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for id, p := range h.pending {
		if secret != "" && p.Secret == secret && p.State == Approved {
			delete(h.pending, id)
			return *p, true
		}
	}
	return Pending{}, false
}
