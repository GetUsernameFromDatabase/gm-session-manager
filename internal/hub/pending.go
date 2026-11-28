package hub

import (
	"crypto/subtle"
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
	// PendingNewCharacter brings a character made at home, as a file waiting
	// for the GM's review; only adding it approves the request.
	PendingNewCharacter PendingKind = "new-character"
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
// A waiting request lives only while the visitor is connected.
const attachTimeout = time.Minute

// approvedTimeout is how long an approval waits to be redeemed, so a phone
// that slept (dropping its websocket) can still join when it wakes.
const approvedTimeout = 10 * time.Minute

// Pending is a join request waiting for the GM. Requests live in memory
// only; a dropped connection drops a waiting or denied request, but an
// approved one is kept until redeemed or approvedTimeout.
type Pending struct {
	// ID is shown to the GM and used in approve/deny URLs.
	ID string
	// Secret is known only to the visitor and redeems an approval.
	Secret     string
	CampaignID int64
	// CampaignName is copied in so every GM page can name it without a query.
	CampaignName string
	Kind         PendingKind
	CharacterID  int64  // PendingClaim; PendingNewCharacter once added
	Name         string // character name (claimed or to create)
	FileID       string // PendingNewCharacter: the file waiting for review
	// Player, Device (OS · browser), Phone and Host tell the GM who is asking.
	Player  string
	Device  string
	Phone   bool
	Host    string // LAN hostname; empty when the network does not know it
	Created time.Time
	State   PendingState
	// Addr is the visitor's address without port, for the per-address cap.
	Addr string

	attached bool
	decided  time.Time
}

// MaxPending bounds memory and GM spam: waiting requests across all
// campaigns. Answered ones do not count, so the GM can always clear it.
const MaxPending = 10

// MaxPendingPerAddr keeps one device from filling MaxPending alone.
const MaxPendingPerAddr = 3

// dropStaleLocked deletes requests whose visitor never connected and
// approvals nobody redeemed.
func (h *Hub) dropStaleLocked() {
	for id, p := range h.pending {
		if (!p.attached && time.Since(p.Created) > attachTimeout) ||
			(p.State == Approved && time.Since(p.decided) > approvedTimeout) {
			delete(h.pending, id)
		}
	}
}

// AddPending records a request, giving it its ID, secret and time, and
// tells the GM. It reports false, publishing nothing, once MaxPending
// requests wait host-wide.
func (h *Hub) AddPending(req Pending) (Pending, bool) {
	p := &req
	p.ID, p.Secret, p.Created, p.State = randomID(idLen), randomID(secretLen), time.Now(), Waiting
	h.mu.Lock()
	defer h.mu.Unlock()
	h.dropStaleLocked()
	waiting, fromAddr := 0, 0
	for _, q := range h.pending {
		if q.State == Waiting {
			waiting++
			if p.Addr != "" && q.Addr == p.Addr {
				fromAddr++
			}
		}
	}
	if waiting >= MaxPending || fromAddr >= MaxPendingPerAddr {
		return Pending{}, false
	}
	h.pending[p.ID] = p
	h.publishLocked(p.CampaignID, Topic{Kind: TopicPending, ID: p.CampaignID})
	return *p, true
}

// PendingList returns a campaign's waiting requests (every campaign's for
// 0), oldest first, dropping any whose visitor never connected.
func (h *Hub) PendingList(campaignID int64) []Pending {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.dropStaleLocked()
	var out []Pending
	for _, p := range h.pending {
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
		if secret != "" && subtle.ConstantTimeCompare([]byte(p.Secret), []byte(secret)) == 1 {
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
	if !ok || p.CampaignID != campaignID || p.State != Waiting ||
		approve && p.Kind == PendingNewCharacter {
		return false
	}
	if approve {
		p.State = Approved
		p.decided = time.Now()
	} else {
		p.State = Denied
	}
	h.publishLocked(campaignID, Topic{Kind: TopicPending, ID: campaignID})
	return true
}

// waitingFileLocked is the waiting new-character request of a file.
func (h *Hub) waitingFileLocked(campaignID int64, fileID string) *Pending {
	for _, p := range h.pending {
		if p.Kind == PendingNewCharacter && p.FileID == fileID && p.CampaignID == campaignID &&
			p.State == Waiting {
			return p
		}
	}
	return nil
}

// FileRequestWaiting reports whether the request that brought a file still waits.
func (h *Hub) FileRequestWaiting(campaignID int64, fileID string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.waitingFileLocked(campaignID, fileID) != nil
}

// ApproveNewCharacter approves the waiting request that brought a file, now
// that the GM added its character. It reports false when the visitor is gone.
func (h *Hub) ApproveNewCharacter(campaignID int64, fileID string, characterID int64) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	p := h.waitingFileLocked(campaignID, fileID)
	if p == nil {
		return false
	}
	p.State, p.CharacterID, p.decided = Approved, characterID, time.Now()
	h.publishLocked(campaignID, Topic{Kind: TopicPending, ID: campaignID})
	return true
}

// DenyFile denies the waiting request that brought a file; false when none waits.
func (h *Hub) DenyFile(campaignID int64, fileID string) bool {
	h.mu.Lock()
	defer h.mu.Unlock()
	p := h.waitingFileLocked(campaignID, fileID)
	if p == nil {
		return false
	}
	p.State = Denied
	h.publishLocked(campaignID, Topic{Kind: TopicPending, ID: campaignID})
	return true
}

// TakeApproved removes and returns an approved request, so its visitor can
// redeem it exactly once.
func (h *Hub) TakeApproved(secret string) (Pending, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.dropStaleLocked() // an approval older than approvedTimeout is spent
	for id, p := range h.pending {
		if secret != "" && subtle.ConstantTimeCompare([]byte(p.Secret), []byte(secret)) == 1 &&
			p.State == Approved {
			delete(h.pending, id)
			return *p, true
		}
	}
	return Pending{}, false
}
