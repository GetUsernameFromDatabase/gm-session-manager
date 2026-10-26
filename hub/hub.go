// Package hub is the real-time layer: one websocket per device, grouped by
// campaign. State changes are published as topics. Each page names the live
// views it shows; for every view a topic touches, the client gets only the
// data that changed since what it last saw, built for it (so the GM and each
// player see only what they may), as a JSON patch.
//
// Events carry (epoch, seq). The epoch changes on every restart. A small
// ring buffer lets a reconnecting client catch up from its last seen seq;
// on an epoch mismatch or a gap older than the buffer the client reloads to
// get fresh state.
package hub

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"slices"
	"sync"
	"time"

	"gm-session-manager/perm"
)

// TopicKind names what changed.
type TopicKind string

// Topics. ID is the row id, or the campaign id for campaign-wide topics.
const (
	// TopicCampaign: roster, connected devices, pending requests, open_join.
	TopicCampaign  TopicKind = "campaign"
	TopicCharacter TopicKind = "character"
	TopicEncounter TopicKind = "encounter"
	// TopicNotes: the campaign's note list.
	TopicNotes TopicKind = "notes"
	// TopicNote: one note's title and block layout.
	TopicNote TopicKind = "note"
	// TopicEntities: the campaign's entity list; TopicEntity: one entity.
	TopicEntities TopicKind = "entities"
	TopicEntity   TopicKind = "entity"
	// TopicPending: a join request changed state.
	TopicPending TopicKind = "pending"
)

// Topic is one thing that changed.
type Topic struct {
	Kind TopicKind
	ID   int64
}

// Handler is implemented by the web layer.
type Handler interface {
	// Affects reports whether any of topics may change the live view.
	Affects(view string, topics []Topic) bool
	// State is the view as client may see it; ok is false when the client
	// may not see it (any more) or it is gone.
	State(ctx context.Context, c *Client, view string) (s State, ok bool, err error)
	// CanSubscribe reports whether client may receive a note block's
	// collaboration steps.
	CanSubscribe(ctx context.Context, c *Client, blockID int64) bool
}

// ringSize is how many events a reconnecting client can catch up on.
const ringSize = 512

// outBuffer is how many messages may queue for one client before it is
// considered stuck and dropped (it reconnects and catches up).
const outBuffer = 64

type event struct {
	seq        uint64
	campaignID int64
	topics     []Topic
}

// Hub fans out events to connected clients.
type Hub struct {
	handler Handler
	epoch   string

	mu      sync.Mutex
	seq     uint64
	ring    []event // oldest first, at most ringSize
	clients map[*Client]struct{}
	pending map[string]*Pending // by ID
}

// New returns a hub with a fresh epoch.
func New(handler Handler) *Hub {
	return &Hub{
		handler: handler,
		epoch:   randomID(8),
		clients: map[*Client]struct{}{},
		pending: map[string]*Pending{},
	}
}

// Position is the current (epoch, seq). Pages render it so the websocket
// can catch up on everything published after the page's data was read;
// read it before querying the data.
func (h *Hub) Position() (epoch string, seq uint64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.epoch, h.seq
}

// Publish records an event for a campaign and pushes it to every synced
// client in that campaign.
func (h *Hub) Publish(campaignID int64, topics ...Topic) {
	if len(topics) == 0 {
		return
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	h.publishLocked(campaignID, topics...)
}

// SendRaw pushes a JSON message (collaboration steps) to the campaign's
// clients for which match returns true. Raw messages are not sequenced:
// the collab protocol has its own versions and recovers on reconnect.
func (h *Hub) SendRaw(campaignID int64, raw []byte, match func(c *Client) bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for c := range h.clients {
		if c.CampaignID == campaignID && c.synced && match(c) {
			h.enqueueLocked(c, message{raw: raw})
		}
	}
}

// CloseSession sends every device on a session back to the start screen
// and disconnects it. Kick and logout both end here.
func (h *Hub) CloseSession(sessionID int64) {
	h.closeMatching(func(c *Client) bool { return c.SessionID != 0 && c.SessionID == sessionID })
}

// CloseCharacter disconnects every device playing a character, after the
// character (and so its sessions) was deleted.
func (h *Hub) CloseCharacter(characterID int64) {
	h.closeMatching(func(c *Client) bool { return c.Viewer.CharacterID == characterID })
}

// CloseCampaign disconnects the players and waiting visitors of a deleted
// campaign; dropping a visitor drops its join request. GM tabs stay, so the
// deleting tab follows its own redirect.
func (h *Hub) CloseCampaign(campaignID int64) {
	h.closeMatching(func(c *Client) bool {
		return c.CampaignID == campaignID && !c.Viewer.IsGM || c.Pending != nil && c.Pending.CampaignID == campaignID
	})
}

func (h *Hub) closeMatching(match func(c *Client) bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for c := range h.clients {
		if match(c) {
			h.enqueueLocked(c, message{control: controlReload, close: true})
		}
	}
}

// Connected lists the clients connected to a campaign.
func (h *Hub) Connected(campaignID int64) []ClientInfo {
	h.mu.Lock()
	defer h.mu.Unlock()
	var out []ClientInfo
	for c := range h.clients {
		if c.CampaignID == campaignID && c.Pending == nil {
			out = append(out, ClientInfo{SessionID: c.SessionID, CharacterID: c.Viewer.CharacterID, IsGM: c.Viewer.IsGM, UserAgent: c.UserAgent, Since: c.connected})
		}
	}
	slices.SortFunc(out, func(a, b ClientInfo) int { return a.Since.Compare(b.Since) })
	return out
}

// ClientInfo describes a connected device for the GM's list.
type ClientInfo struct {
	SessionID   int64
	CharacterID int64
	IsGM        bool
	UserAgent   string
	Since       time.Time
}

// Close disconnects everyone, for shutdown.
func (h *Hub) Close() {
	h.mu.Lock()
	defer h.mu.Unlock()
	for c := range h.clients {
		c.stop()
	}
}

func (h *Hub) register(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[c] = struct{}{}
	if c.Pending != nil {
		c.Pending.attached = true
	}
	h.announceDevice(c)
}

func (h *Hub) unregister(c *Client) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, c)
	c.stop()
	h.announceDevice(c)
	if p := c.Pending; p != nil && h.pending[p.ID] == p {
		// A dropped connection drops the request; the visitor asks again.
		delete(h.pending, p.ID)
		h.publishLocked(p.CampaignID, Topic{Kind: TopicPending, ID: p.CampaignID})
	}
}

// publishLocked is Publish for callers already holding h.mu.
func (h *Hub) publishLocked(campaignID int64, topics ...Topic) {
	h.seq++
	h.ring = append(h.ring, event{seq: h.seq, campaignID: campaignID, topics: topics})
	if len(h.ring) > ringSize {
		h.ring = slices.Delete(h.ring, 0, len(h.ring)-ringSize)
	}
	for c := range h.clients {
		if c.synced && c.receives(campaignID, topics) {
			h.enqueueLocked(c, message{topics: topics, seq: h.seq})
		}
	}
}

// sync handles a client's hello: catch up from (epoch, seq) if the ring
// still covers it, otherwise tell the client to reload.
func (h *Hub) sync(c *Client, epoch string, seq uint64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if c.synced {
		return
	}
	c.synced = true

	oldest := h.seq + 1
	if len(h.ring) > 0 {
		oldest = h.ring[0].seq
	}
	if epoch != h.epoch || seq > h.seq || seq+1 < oldest {
		h.enqueueLocked(c, message{control: controlReload})
		return
	}
	var topics []Topic
	for _, ev := range h.ring {
		if ev.seq > seq && c.receives(ev.campaignID, ev.topics) {
			for _, t := range ev.topics {
				if !slices.Contains(topics, t) {
					topics = append(topics, t)
				}
			}
		}
	}
	h.enqueueLocked(c, message{topics: topics, seq: h.seq, hello: true})
}

// receives reports whether c gets an event: its own campaign's, and for the
// GM every campaign's join requests, shown on whatever page they are on.
func (c *Client) receives(campaignID int64, topics []Topic) bool {
	return c.CampaignID == campaignID || c.Viewer.IsGM && slices.ContainsFunc(topics, func(t Topic) bool { return t.Kind == TopicPending })
}

func (h *Hub) enqueueLocked(c *Client, m message) {
	select {
	case c.out <- m:
	default:
		slog.Warn("websocket client too slow, dropping", "campaign", c.CampaignID)
		c.stop()
	}
}

func randomID(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b) // crypto/rand.Read never fails
	return hex.EncodeToString(b)
}

// Client is one connected device.
type Client struct {
	CampaignID int64
	Viewer     perm.Viewer
	// SessionID is the sessions row, or 0 for localhost admin and visitors.
	SessionID int64
	// Pending is set for a visitor waiting on a join request.
	Pending   *Pending
	UserAgent string

	connected time.Time
	out       chan message
	done      chan struct{}
	stopOnce  sync.Once
	synced    bool // guarded by Hub.mu

	subMu  sync.Mutex
	blocks map[int64]bool
	watch  []string // live views on the page, guarded by subMu
	// last is the state each watched view had when last sent; only the
	// connection's writer touches it.
	last map[string]State
}

// NewClient prepares a client; Serve connects it.
func NewClient(campaignID int64, v perm.Viewer, sessionID int64, userAgent string) *Client {
	return &Client{
		CampaignID: campaignID, Viewer: v, SessionID: sessionID, UserAgent: userAgent,
		connected: time.Now(), out: make(chan message, outBuffer), done: make(chan struct{}),
		blocks: map[int64]bool{}, last: map[string]State{},
	}
}

// Subscribed reports whether the client asked for a block's steps.
func (c *Client) Subscribed(blockID int64) bool {
	c.subMu.Lock()
	defer c.subMu.Unlock()
	return c.blocks[blockID]
}

func (c *Client) subscribe(blockID int64, on bool) {
	c.subMu.Lock()
	defer c.subMu.Unlock()
	if on {
		c.blocks[blockID] = true
	} else {
		delete(c.blocks, blockID)
	}
}

func (c *Client) stop() {
	c.stopOnce.Do(func() { close(c.done) })
}

type message struct {
	topics  []Topic
	seq     uint64
	raw     []byte
	control string
	close   bool
	// hello is the catch-up reply: it baselines every watched view and
	// always carries the position.
	hello bool
}

const controlReload = "reload"

// announceDevice tells the GM's device list that a player device came or
// went. Callers hold h.mu.
func (h *Hub) announceDevice(c *Client) {
	if c.Pending == nil && c.SessionID != 0 && c.Viewer.IsPlayer() {
		h.publishLocked(c.CampaignID, Topic{Kind: TopicCampaign, ID: c.CampaignID})
	}
}
