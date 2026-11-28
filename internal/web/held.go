package web

import (
	"context"
	"database/sql"
	"sync"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/hub"
)

// An effect can hang off a concentration row (conditions.held_by) and goes
// when that row does, by cascade. The owners of those effects are not part
// of the request, so before any delete that can cascade the handler calls
// noteHeld, and publish adds the topics it collected.

// heldTopics are the topics noted during one request.
type heldTopics struct {
	mu  sync.Mutex
	set map[hub.Topic]struct{}
}

func (h *heldTopics) add(t hub.Topic) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.set == nil {
		h.set = map[hub.Topic]struct{}{}
	}
	h.set[t] = struct{}{}
}

// collectHeld gives each request a place to note the topics of cascaded effects.
func collectHeld(c *gin.Context) {
	c.Request = c.Request.WithContext(
		context.WithValue(c.Request.Context(), heldTopicsKey, &heldTopics{}),
	)
}

// heldCollected are the topics noted so far.
func heldCollected(ctx context.Context) []hub.Topic {
	h, _ := ctx.Value(heldTopicsKey).(*heldTopics)
	if h == nil {
		return nil
	}
	h.mu.Lock()
	defer h.mu.Unlock()
	out := make([]hub.Topic, 0, len(h.set))
	for t := range h.set {
		out = append(out, t)
	}
	return out
}

// noteHeld records the owners of the effects held by the concentration rows
// holders, which are about to be deleted: a character and the fights it is
// in, or a combatant's fight. A rolled-back delete only costs a refresh.
func noteHeld(ctx context.Context, q *sqlc.Queries, holders ...int64) error {
	h, _ := ctx.Value(heldTopicsKey).(*heldTopics)
	if h == nil || len(holders) == 0 {
		return nil
	}
	ids := make([]sql.NullInt64, len(holders))
	for i, id := range holders {
		ids[i] = nullID(id)
	}
	rows, err := q.ListHeldBy(ctx, ids)
	if err != nil {
		return err
	}
	for _, r := range rows {
		if r.EncounterID.Valid {
			h.add(hub.Topic{Kind: hub.TopicEncounter, ID: r.EncounterID.Int64})
		}
		if !r.CharacterID.Valid {
			continue
		}
		h.add(hub.Topic{Kind: hub.TopicCharacter, ID: r.CharacterID.Int64})
		encounters, err := q.ListEncounterIDsByCharacter(ctx, r.CharacterID)
		if err != nil {
			return err
		}
		for _, e := range encounters {
			h.add(hub.Topic{Kind: hub.TopicEncounter, ID: e})
		}
	}
	return nil
}

// concentrationIDs are the ids of the concentration rows among conds.
func concentrationIDs(conds []sqlc.Condition) []int64 {
	var ids []int64
	for _, c := range conds {
		if c.Slug == slugConcentration {
			ids = append(ids, c.ID)
		}
	}
	return ids
}

// publish sends topics to the campaign with those of the effects the request
// deleted by cascade.
func (a *App) publish(c *gin.Context, campaignID int64, topics ...hub.Topic) {
	a.Hub.Publish(campaignID, append(topics, heldCollected(c.Request.Context())...)...)
}
