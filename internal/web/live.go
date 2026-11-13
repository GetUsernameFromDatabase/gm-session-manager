package web

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"hash/fnv"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// liveView is one live view loaded for a viewer.
type liveView struct {
	// html is the view's fragment(s) as the page shows them; nil renders
	// nothing.
	html templ.Component
	// shape is html with the bound values held fixed, hashed to notice
	// every other change; nil means html.
	shape templ.Component
	// model, when set, is hashed as JSON instead of rendering shape (about
	// three quarters of a publish): shape's markup must read nothing else.
	model any
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

// liveKind is one kind of live view: which published topics may change it,
// and how to load it for a viewer.
type liveKind struct {
	own  hub.TopicKind   // topics of this kind with the view's id
	any  []hub.TopicKind // topics of these kinds with any id
	load func(a *App, ctx context.Context, vc viewerCtx, k liveKey) (liveView, error)
}

// liveKinds registers every live view. Publishes are already scoped to the
// campaign, so campaign-wide views take any topic of their kind.
var liveKinds = map[string]liveKind{
	// The roster shows every character's hit points.
	views.LiveRoster:   {any: []hub.TopicKind{hub.TopicCampaign, hub.TopicCharacter}, load: (*App).loadRosterLive},
	views.LiveRequests: {any: []hub.TopicKind{hub.TopicPending}, load: (*App).loadRequestsLive},
	views.LiveWaiting:  {any: []hub.TopicKind{hub.TopicPending}, load: (*App).loadWaitingLive},
	// An import, priority move or book removal can change the sheet's source warning.
	views.LiveCharacter:  {own: hub.TopicCharacter, any: []hub.TopicKind{hub.TopicCompendium}, load: (*App).loadCharacterLive},
	views.LiveEncounters: {any: []hub.TopicKind{hub.TopicEncounter}, load: (*App).loadEncountersLive},
	// Trackers show NPC portraits, which change through their NPC.
	views.LiveEncounter:    {own: hub.TopicEncounter, any: []hub.TopicKind{hub.TopicEntity}, load: (*App).loadEncounterLive},
	views.LivePlayerCombat: {any: []hub.TopicKind{hub.TopicEncounter, hub.TopicEntity}, load: (*App).loadPlayerCombatLive},
	views.LiveNotes:        {any: []hub.TopicKind{hub.TopicNotes}, load: (*App).loadNotesLive},
	views.LiveNote:         {own: hub.TopicNote, load: (*App).loadNoteLive},
	views.LiveEntities:     {any: []hub.TopicKind{hub.TopicEntities}, load: (*App).loadEntitiesLive},
	views.LiveEntity:       {own: hub.TopicEntity, load: (*App).loadEntityLive},
	views.LiveImport:       {any: []hub.TopicKind{hub.TopicImport}, load: (*App).loadImportLive},
}

// Affects implements hub.Handler.
func (a *App) Affects(view string, topics []hub.Topic) bool {
	k := parseLiveKey(view)
	kind, ok := liveKinds[k.kind]
	if !ok {
		return false
	}
	return slices.ContainsFunc(topics, func(t hub.Topic) bool {
		return t.Kind == kind.own && t.ID == k.num() || slices.Contains(kind.any, t.Kind)
	})
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
	buf := shapeBufs.Get().(*bytes.Buffer)
	defer shapeBufs.Put(buf)
	buf.Reset()
	var hash string
	if lv.model != nil {
		if err := json.NewEncoder(buf).Encode(lv.model); err != nil {
			return nil, false, err
		}
		// No random ids here, only user text that may look like one.
		h := fnv.New64a()
		_, _ = h.Write(buf.Bytes()) // fnv never fails
		hash = strconv.FormatUint(h.Sum64(), 36)
	} else {
		shape := lv.shape
		if shape == nil {
			shape = lv.html
		}
		if err := shape.Render(ctx, buf); err != nil {
			return nil, false, err
		}
		hash = shapeHash(buf.Bytes())
	}
	st := hub.State{hub.HashKey: hash}
	if lv.hot != nil {
		lv.hot(st)
	}
	return st, true, nil
}

// shapeBufs keeps render buffers between publishes: a sheet is ~1 MB of markup.
var shapeBufs = sync.Pool{New: func() any { return new(bytes.Buffer) }}

// randomIDPrefix and randomIDLen describe utils.RandomID ("id-" +
// rand.Text()), which components use for unnamed elements: new on every
// render, so not part of the shape.
const (
	randomIDPrefix = "id-"
	randomIDLen    = 26
)

// shapeHash hashes markup without its random ids, in place: a regexp
// replace copied the whole markup on every publish.
func shapeHash(b []byte) string {
	h := fnv.New64a()
	for {
		i := bytes.Index(b, []byte(randomIDPrefix))
		if i < 0 {
			break
		}
		end := i + len(randomIDPrefix)
		if isRandomID(b[end:]) {
			_, _ = h.Write(b[:i]) // fnv never fails
			b = b[end+randomIDLen:]
			continue
		}
		_, _ = h.Write(b[:end])
		b = b[end:]
	}
	_, _ = h.Write(b)
	return strconv.FormatUint(h.Sum64(), 36)
}

// isRandomID reports whether b starts with rand.Text's base32 characters.
func isRandomID(b []byte) bool {
	if len(b) < randomIDLen {
		return false
	}
	for _, r := range b[:randomIDLen] {
		if (r < 'A' || r > 'Z') && (r < '2' || r > '7') {
			return false
		}
	}
	return true
}

// liveHTML serves a view's fragments, for live.js to refetch a view whose
// markup changed shape. Visitors name their request's secret instead of a
// session.
func (a *App) liveHTML(c *gin.Context) error {
	w, err := a.whoami(c)
	if err != nil {
		return err
	}
	key := c.Query("v")
	vc := viewerCtx{viewer: w.viewer, campaignID: w.campaignID()}
	if k := parseLiveKey(key); k.kind == views.LiveWaiting && w.id.Session == nil {
		vc.pending, _ = a.Hub.PendingBySecret(k.id)
	}
	lv, err := a.loadLive(c.Request.Context(), vc, key)
	if err != nil {
		return err
	}
	if lv.html == nil {
		c.Status(http.StatusNoContent)
		return nil
	}
	// Compendium links follow the page's campaign: the GM's names it, a
	// player's is always their own.
	campaign := vc.campaignID
	if w.isGM() {
		campaign, _ = strconv.ParseInt(c.Query("campaign"), 10, 64)
	}
	c.Request = c.Request.WithContext(views.WithCampaign(c.Request.Context(), campaign))
	a.render(c, http.StatusOK, lv.html)
	return nil
}

// loadLive loads a view as vc may see it. Players only get views of their
// own campaign; the id-keyed builders apply the finer rules.
func (a *App) loadLive(ctx context.Context, vc viewerCtx, key string) (liveView, error) {
	k := parseLiveKey(key)
	kind, ok := liveKinds[k.kind]
	if !ok {
		return liveView{}, nil
	}
	return kind.load(a, ctx, vc, k)
}

// member reports whether vc may see campaign-wide views of campaignID.
func (vc viewerCtx) member(campaignID int64) bool {
	return vc.viewer.IsGM || (vc.viewer.IsPlayer() && campaignID == vc.campaignID)
}

// inCampaign reports whether a row of campaignID is in vc's reach.
func (vc viewerCtx) inCampaign(campaignID int64) bool {
	return vc.viewer.IsGM || campaignID == vc.campaignID
}

func (a *App) loadRosterLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	if !vc.viewer.IsGM {
		return liveView{}, nil
	}
	r, err := a.roster(ctx, k.num())
	if err != nil {
		return liveView{}, notFoundIsNil(err)
	}
	return liveView{html: views.CampaignRoster(r), shape: views.CampaignRoster(r.Shape()), hot: r.Hot, ok: true}, nil
}

func (a *App) loadImportLive(_ context.Context, vc viewerCtx, _ liveKey) (liveView, error) {
	if !vc.viewer.IsGM {
		return liveView{}, nil
	}
	j := a.fiveTools.get()
	return liveView{html: views.FiveToolsProgress(j, true), shape: views.FiveToolsProgress(j.Shape(), true), hot: j.Hot, ok: true}, nil
}

func (a *App) loadRequestsLive(_ context.Context, vc viewerCtx, _ liveKey) (liveView, error) {
	if !vc.viewer.IsGM {
		return liveView{}, nil
	}
	return liveView{html: views.PendingRequests(a.pendingViews()), ok: true}, nil
}

func (a *App) loadWaitingLive(_ context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	if p := vc.pending; p != nil && p.Secret == k.id {
		return liveView{html: views.PendingStatus(a.waiting(p)), ok: true}, nil
	}
	return liveView{}, nil
}

func (a *App) loadCharacterLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	if !perm.CanCharacter(vc.viewer, k.num(), perm.ScopeView) {
		return liveView{}, nil
	}
	row, err := a.Store.GetCharacter(ctx, k.num())
	if err != nil || !vc.inCampaign(row.CampaignID) {
		return liveView{}, notFoundIsNil(err)
	}
	s, err := a.sheetView(ctx, vc.viewer, row)
	if err != nil {
		return liveView{}, err
	}
	return liveView{html: views.CharacterSheet(s), model: s.Shape(), hot: s.Hot, ok: true}, nil
}

func (a *App) loadEncountersLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	if !vc.viewer.IsGM {
		return liveView{}, nil
	}
	list, err := a.Store.ListEncountersByCampaign(ctx, k.num())
	if err != nil {
		return liveView{}, err
	}
	current, _ := strconv.ParseInt(k.extra, 10, 64)
	return liveView{html: views.EncounterList(k.num(), list, current), ok: true}, nil
}

func (a *App) loadEncounterLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	if !vc.viewer.IsGM {
		return liveView{}, nil
	}
	enc, err := a.Store.GetEncounter(ctx, k.num())
	if err != nil {
		return liveView{}, notFoundIsNil(err)
	}
	ev, err := a.encounterView(ctx, vc.viewer, enc)
	if err != nil {
		return liveView{}, err
	}
	return liveView{html: views.Encounter(ev), model: ev.Shape(), hot: ev.Hot, ok: true}, nil
}

func (a *App) loadPlayerCombatLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	if !vc.viewer.IsPlayer() || !vc.member(k.num()) {
		return liveView{}, nil
	}
	want, _ := strconv.ParseInt(k.extra, 10, 64)
	list, ev, err := a.playerCombat(ctx, vc.viewer, k.num(), want)
	if err != nil {
		return liveView{}, err
	}
	lv := liveView{html: views.PlayerCombat(k.num(), want, list, ev), ok: true}
	if ev != nil {
		shaped := ev.Shape()
		lv.shape, lv.hot = views.PlayerCombat(k.num(), want, list, &shaped), ev.Hot
	}
	return lv, nil
}

func (a *App) loadNotesLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	if !vc.member(k.num()) {
		return liveView{}, nil
	}
	notes, err := a.noteList(ctx, vc.viewer, k.num())
	if err != nil {
		return liveView{}, err
	}
	return liveView{html: views.NoteList(k.num(), notes), ok: true}, nil
}

func (a *App) loadNoteLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	note, err := a.Store.GetNote(ctx, k.num())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return liveView{}, err
	}
	gone := liveView{html: views.NoteGone(k.num())} // clears a page that showed the note
	if err != nil || !vc.inCampaign(note.CampaignID) {
		return gone, nil
	}
	nv, visible, err := a.noteView(ctx, vc.viewer, note)
	if err != nil {
		return liveView{}, err
	}
	if !visible {
		return gone, nil
	}
	return liveView{html: templ.Join(views.NoteMeta(nv), views.NoteBlocks(nv)), ok: true}, nil
}

func (a *App) loadEntitiesLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	if !vc.member(k.num()) {
		return liveView{}, nil
	}
	list, err := a.entityList(ctx, vc.viewer, k.num())
	if err != nil {
		return liveView{}, err
	}
	return liveView{html: views.EntityList(list), ok: true}, nil
}

func (a *App) loadEntityLive(ctx context.Context, vc viewerCtx, k liveKey) (liveView, error) {
	ev, ok, err := a.entityView(ctx, vc.viewer, k.num())
	if err != nil {
		return liveView{}, err
	}
	if !ok || !vc.inCampaign(ev.CampaignID) {
		// Clears a page that showed the entity before it was hidden.
		return liveView{html: views.EntityGone(k.num())}, nil
	}
	return liveView{html: views.EntityDetail(ev), ok: true}, nil
}

// notFoundIsNil turns a missing row into "not visible".
func notFoundIsNil(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}
