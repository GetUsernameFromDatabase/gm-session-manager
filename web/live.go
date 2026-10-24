package web

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"hash/fnv"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

// liveView is one live view loaded for a viewer.
type liveView struct {
	// html is the view's fragment(s) as the page shows them; nil renders
	// nothing.
	html templ.Component
	// shape is html with the bound values held fixed, hashed to notice
	// every other change; nil means html.
	shape templ.Component
	hot   func(hub.State)
	// ok is false when the viewer may not see the view (any more) or it is
	// gone; html may still clear what the page showed.
	ok bool
}

// liveKey is a parsed view key: kind:id[:extra].
type liveKey struct {
	kind, id, extra string
}

func parseLiveKey(key string) liveKey {
	parts := strings.SplitN(key, ":", 3)
	parts = append(parts, "", "")
	return liveKey{kind: parts[0], id: parts[1], extra: parts[2]}
}

func (k liveKey) num() int64 {
	n, _ := strconv.ParseInt(k.id, 10, 64)
	return n
}

// liveAffects maps published topics to the views they may change.
var liveAffects = map[string]struct {
	kind  hub.TopicKind
	anyID bool // every topic of the kind, not just the view's id
}{
	views.LiveRoster:       {hub.TopicCampaign, true},
	views.LiveRequests:     {hub.TopicPending, true},
	views.LiveWaiting:      {hub.TopicPending, true},
	views.LiveCharacter:    {hub.TopicCharacter, false},
	views.LiveEncounters:   {hub.TopicEncounter, true},
	views.LiveEncounter:    {hub.TopicEncounter, false},
	views.LivePlayerCombat: {hub.TopicEncounter, true},
	views.LiveNotes:        {hub.TopicNotes, true},
	views.LiveNote:         {hub.TopicNote, false},
	views.LiveEntities:     {hub.TopicEntities, true},
	views.LiveEntity:       {hub.TopicEntity, false},
}

// Affects implements hub.Handler. Publishes are already scoped to the
// campaign, so campaign-wide views take any topic of their kind.
func (a *App) Affects(view string, topics []hub.Topic) bool {
	k := parseLiveKey(view)
	rule, ok := liveAffects[k.kind]
	for _, t := range topics {
		switch {
		case !ok:
		case t.Kind == rule.kind && (rule.anyID || t.ID == k.num()):
			return true
		// The roster shows every character's hit points.
		case k.kind == views.LiveRoster && t.Kind == hub.TopicCharacter:
			return true
		// Trackers show NPC portraits, which change through their NPC.
		case (k.kind == views.LiveEncounter || k.kind == views.LivePlayerCombat) && t.Kind == hub.TopicEntity:
			return true
		}
	}
	return false
}

// State implements hub.Handler.
func (a *App) State(ctx context.Context, c *hub.Client, view string) (hub.State, bool, error) {
	lv, err := a.loadLive(ctx, viewerCtx{viewer: c.Viewer, campaignID: c.CampaignID, pending: c.Pending}, view)
	if err != nil || !lv.ok {
		return nil, false, err
	}
	return lv.state(ctx)
}

func (lv liveView) state(ctx context.Context) (hub.State, bool, error) {
	shape := lv.shape
	if shape == nil {
		shape = lv.html
	}
	var buf bytes.Buffer
	if err := shape.Render(ctx, &buf); err != nil {
		return nil, false, err
	}
	h := fnv.New64a()
	_, _ = h.Write(randomIDs.ReplaceAll(buf.Bytes(), nil)) // fnv never fails
	st := hub.State{hub.HashKey: strconv.FormatUint(h.Sum64(), 36)}
	if lv.hot != nil {
		lv.hot(st)
	}
	return st, true, nil
}

// randomIDs matches utils.RandomID ("id-" + rand.Text()), which components
// use for unnamed elements: new on every render, so not part of the shape.
var randomIDs = regexp.MustCompile(`id-[A-Z2-7]{26}`)

// liveHTML serves a view's fragments, for live.js to refetch a view whose
// markup changed shape. Visitors name their request's secret instead of a
// session.
func (a *App) liveHTML(c *gin.Context) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return
	}
	key := c.Query("v")
	vc := viewerCtx{viewer: w.viewer, campaignID: w.campaignID()}
	if k := parseLiveKey(key); k.kind == views.LiveWaiting && w.id.Session == nil {
		vc.pending, _ = a.Hub.PendingBySecret(k.id)
	}
	lv, err := a.loadLive(c.Request.Context(), vc, key)
	if err != nil {
		a.serverError(c, err)
		return
	}
	if lv.html == nil {
		c.Status(http.StatusNoContent)
		return
	}
	a.render(c, http.StatusOK, lv.html)
}

// loadLive loads a view as vc may see it. Players only get views of their
// own campaign; the id-keyed builders apply the finer rules.
func (a *App) loadLive(ctx context.Context, vc viewerCtx, key string) (liveView, error) {
	k := parseLiveKey(key)
	v := vc.viewer
	member := v.IsGM || (v.IsPlayer() && k.num() == vc.campaignID)
	inCampaign := func(campaignID int64) bool { return v.IsGM || campaignID == vc.campaignID }
	switch k.kind {
	case views.LiveRoster:
		if !v.IsGM {
			break
		}
		r, err := a.roster(ctx, k.num())
		if err != nil {
			return liveView{}, notFoundIsNil(err)
		}
		return liveView{html: views.CampaignRoster(r), shape: views.CampaignRoster(r.Shape()), hot: r.Hot, ok: true}, nil
	case views.LiveRequests:
		if v.IsGM {
			return liveView{html: views.PendingRequests(a.pendingViews()), ok: true}, nil
		}
	case views.LiveWaiting:
		if p := vc.pending; p != nil && p.Secret == k.id {
			return liveView{html: views.PendingStatus(a.waiting(p)), ok: true}, nil
		}
	case views.LiveCharacter:
		if !perm.CanCharacter(v, k.num(), perm.ScopeView) {
			break
		}
		row, err := a.Store.GetCharacter(ctx, k.num())
		if err != nil || !inCampaign(row.CampaignID) {
			return liveView{}, notFoundIsNil(err)
		}
		s, err := a.sheetView(ctx, v, row)
		if err != nil {
			return liveView{}, err
		}
		return liveView{html: views.CharacterSheet(s), shape: views.CharacterSheet(s.Shape()), hot: s.Hot, ok: true}, nil
	case views.LiveEncounters:
		if !v.IsGM {
			break
		}
		list, err := a.Store.ListEncountersByCampaign(ctx, k.num())
		if err != nil {
			return liveView{}, err
		}
		current, _ := strconv.ParseInt(k.extra, 10, 64)
		return liveView{html: views.EncounterList(k.num(), list, current), ok: true}, nil
	case views.LiveEncounter:
		if !v.IsGM {
			break
		}
		enc, err := a.Store.GetEncounter(ctx, k.num())
		if err != nil {
			return liveView{}, notFoundIsNil(err)
		}
		ev, err := a.encounterView(ctx, v, enc)
		if err != nil {
			return liveView{}, err
		}
		return liveView{html: views.Encounter(ev), shape: views.Encounter(ev.Shape()), hot: ev.Hot, ok: true}, nil
	case views.LivePlayerCombat:
		if !v.IsPlayer() || !member {
			break
		}
		want, _ := strconv.ParseInt(k.extra, 10, 64)
		list, ev, err := a.playerCombat(ctx, v, k.num(), want)
		if err != nil {
			return liveView{}, err
		}
		lv := liveView{html: views.PlayerCombat(k.num(), want, list, ev), ok: true}
		if ev != nil {
			shaped := ev.Shape()
			lv.shape, lv.hot = views.PlayerCombat(k.num(), want, list, &shaped), ev.Hot
		}
		return lv, nil
	case views.LiveNotes:
		if !member {
			break
		}
		notes, err := a.noteList(ctx, v, k.num())
		if err != nil {
			return liveView{}, err
		}
		return liveView{html: views.NoteList(k.num(), notes), ok: true}, nil
	case views.LiveNote:
		note, err := a.Store.GetNote(ctx, k.num())
		if err != nil || !inCampaign(note.CampaignID) {
			return liveView{}, notFoundIsNil(err)
		}
		nv, visible, err := a.noteView(ctx, v, note)
		if err != nil || !visible {
			return liveView{}, err
		}
		return liveView{html: templ.Join(views.NoteMeta(nv), views.NoteBlocks(nv)), ok: true}, nil
	case views.LiveEntities:
		if !member {
			break
		}
		list, err := a.entityList(ctx, v, k.num())
		if err != nil {
			return liveView{}, err
		}
		return liveView{html: views.EntityList(list), ok: true}, nil
	case views.LiveEntity:
		ev, ok, err := a.entityView(ctx, v, k.num())
		if err != nil {
			return liveView{}, err
		}
		if !ok || !inCampaign(ev.CampaignID) {
			// Clears a page that showed the entity before it was hidden.
			return liveView{html: views.EntityGone(k.num())}, nil
		}
		return liveView{html: views.EntityDetail(ev), ok: true}, nil
	}
	return liveView{}, nil
}

// notFoundIsNil turns a missing row into "not visible".
func notFoundIsNil(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}
