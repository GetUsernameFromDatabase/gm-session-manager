package web

import (
	"database/sql"
	"fmt"
	"hash/fnv"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// liveFixture is a campaign with a hidden-HP monster in a running fight, an
// NPC with a secret field, a GM-only note, and a player in another
// campaign.
type liveFixture struct {
	testEnv
	ogre, entity, gmNote, encounter int64
	outsider                        perm.Viewer
	outsiderCampaign                int64
}

func newLiveFixture(t testing.TB) liveFixture {
	t.Helper()
	e := newTestEnv(t)
	ctx := t.Context()
	f := liveFixture{testEnv: e}
	ent, err := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{
		CampaignID: e.campaignID, Kind: "npc", Name: "Mira", IsRevealed: true,
		FieldsJson: `{"role":{"value":"Innkeeper","is_revealed":true},"secret":{"value":"Is the lich","is_revealed":false}}`,
	})
	if err != nil {
		t.Fatal(err)
	}
	f.entity = ent.ID
	note, _ := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "generic", Title: "Plans"})
	if _, err := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: "gm"}); err != nil {
		t.Fatal(err)
	}
	f.gmNote = note.ID
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	ogre, err := e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: enc.ID, Name: "Ogre", Hp: 59, MaxHp: 59, Ac: 11})
	if err != nil {
		t.Fatal(err)
	}
	// Players only see fights their character is in.
	if _, err := e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: enc.ID, CharacterID: nullID(e.characterID), Name: "Alice"}); err != nil {
		t.Fatal(err)
	}
	if err := e.q().UpdateEncounterState(ctx, sqlc.UpdateEncounterStateParams{ID: enc.ID, Status: "active", Round: 1, ActiveCombatantID: nullID(ogre.ID)}); err != nil {
		t.Fatal(err)
	}
	f.ogre, f.encounter = ogre.ID, enc.ID
	other, _ := e.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Elsewhere", System: "dnd5e"})
	stranger, err := e.q().CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: other.ID, Name: "Bob", System: "dnd5e"})
	if err != nil {
		t.Fatal(err)
	}
	f.outsider, f.outsiderCampaign = perm.Player(stranger.ID), other.ID
	return f
}

func (f liveFixture) state(t *testing.T, v perm.Viewer, campaignID int64, view string) (hub.State, bool) {
	t.Helper()
	st, ok, err := f.app.State(t.Context(), hub.NewClient(campaignID, v, 0, "test"), view)
	if err != nil {
		t.Fatal(err)
	}
	return st, ok
}

// TestLiveVisibility checks what each viewer's live state may hold: views
// they may not see are not served at all, and hidden values are never keys.
func TestLiveVisibility(t *testing.T) {
	f := newLiveFixture(t)
	player := perm.Player(f.characterID)
	ogreHP := fmt.Sprintf("c.%d.hp", f.ogre)
	tests := []struct {
		name       string
		viewer     perm.Viewer
		campaignID int64
		view       string
		wantOK     bool
		wantKeys   []string
		noKeys     []string
	}{
		{"GM sees monster HP", perm.GM, f.campaignID, views.LiveKey(views.LiveEncounter, f.encounter), true, []string{ogreHP, "turn.cur"}, nil},
		{"player's fight hides monster HP", player, f.campaignID, views.LiveKey(views.LivePlayerCombat, f.campaignID), true, []string{fmt.Sprintf("c.%d.on", f.ogre), fmt.Sprintf("c.%d.down", f.ogre)}, []string{ogreHP}},
		{"player may not open the GM tracker", player, f.campaignID, views.LiveKey(views.LiveEncounter, f.encounter), false, nil, nil},
		{"player may not see the roster", player, f.campaignID, views.LiveKey(views.LiveRoster, f.campaignID), false, nil, nil},
		{"player may not see the encounter list", player, f.campaignID, views.LiveKey(views.LiveEncounters, f.campaignID), false, nil, nil},
		{"GM sees join requests outside a campaign", perm.GM, 0, views.LiveKey(views.LiveRequests, 0), true, nil, nil},
		{"player may not see join requests", player, f.campaignID, views.LiveKey(views.LiveRequests, 0), false, nil, nil},
		{"GM sees the 5etools import", perm.GM, 0, views.LiveKey(views.LiveImport, 0), true, nil, nil},
		{"player may not see the 5etools import", player, f.campaignID, views.LiveKey(views.LiveImport, 0), false, nil, nil},
		{"player sees a revealed NPC", player, f.campaignID, views.LiveKey(views.LiveEntity, f.entity), true, nil, nil},
		{"player may not see a GM-only note", player, f.campaignID, views.LiveKey(views.LiveNote, f.gmNote), false, nil, nil},
		{"player sees own sheet with HP", player, f.campaignID, views.LiveKey(views.LiveCharacter, f.characterID), true, []string{"hp"}, nil},
		{"outsider may not see another campaign's sheet", f.outsider, f.outsiderCampaign, views.LiveKey(views.LiveCharacter, f.characterID), false, nil, nil},
		{"outsider may not list another campaign's notes", f.outsider, f.outsiderCampaign, views.LiveKey(views.LiveNotes, f.campaignID), false, nil, nil},
		{"outsider may not list another campaign's NPCs", f.outsider, f.outsiderCampaign, views.LiveKey(views.LiveEntities, f.campaignID), false, nil, nil},
		{"outsider may not see another campaign's NPC", f.outsider, f.outsiderCampaign, views.LiveKey(views.LiveEntity, f.entity), false, nil, nil},
		{"visitor sees nothing", perm.Viewer{}, f.campaignID, views.LiveKey(views.LiveNotes, f.campaignID), false, nil, nil},
		{"visitor needs the right secret", perm.Viewer{}, f.campaignID, views.LiveKey(views.LiveWaiting, "guess"), false, nil, nil},
		{"unknown view", perm.GM, f.campaignID, "bogus:1", false, nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st, ok := f.state(t, tt.viewer, tt.campaignID, tt.view)
			if ok != tt.wantOK {
				t.Fatalf("ok = %v, want %v (%v)", ok, tt.wantOK, st)
			}
			for _, k := range tt.wantKeys {
				if _, has := st[k]; !has {
					t.Errorf("missing %s in %v", k, st)
				}
			}
			for _, k := range tt.noKeys {
				if _, has := st[k]; has {
					t.Errorf("leaked %s in %v", k, st)
				}
			}
		})
	}
	// Every live view is registered and has at least one row above.
	all := []string{
		views.LiveRoster, views.LiveRequests, views.LiveWaiting, views.LiveCharacter,
		views.LiveEncounters, views.LiveEncounter, views.LivePlayerCombat, views.LiveNotes,
		views.LiveNote, views.LiveEntities, views.LiveEntity, views.LiveImport,
	}
	if len(liveKinds) != len(all) {
		t.Errorf("liveKinds has %d kinds, this test lists %d", len(liveKinds), len(all))
	}
	covered := map[string]bool{}
	for _, tt := range tests {
		covered[parseLiveKey(tt.view).kind] = true
	}
	for _, kind := range all {
		if _, ok := liveKinds[kind]; !ok {
			t.Errorf("%s is not in liveKinds", kind)
		}
		if !covered[kind] {
			t.Errorf("%s has no TestLiveVisibility row", kind)
		}
	}
}

// TestLiveAffects checks which published topics reach which views.
func TestLiveAffects(t *testing.T) {
	a := &App{}
	tests := []struct {
		name  string
		view  string
		topic hub.Topic
		want  bool
	}{
		{"own id", views.LiveKey(views.LiveCharacter, 1), hub.Topic{Kind: hub.TopicCharacter, ID: 1}, true},
		{"other id", views.LiveKey(views.LiveCharacter, 1), hub.Topic{Kind: hub.TopicCharacter, ID: 2}, false},
		{"sheet on compendium", views.LiveKey(views.LiveCharacter, 1), hub.Topic{Kind: hub.TopicCompendium}, true},
		{"roster on any campaign topic", views.LiveKey(views.LiveRoster, 1), hub.Topic{Kind: hub.TopicCampaign, ID: 9}, true},
		{"roster on character HP", views.LiveKey(views.LiveRoster, 1), hub.Topic{Kind: hub.TopicCharacter, ID: 9}, true},
		{"tracker on its encounter", views.LiveKey(views.LiveEncounter, 1), hub.Topic{Kind: hub.TopicEncounter, ID: 1}, true},
		{"tracker on another encounter", views.LiveKey(views.LiveEncounter, 1), hub.Topic{Kind: hub.TopicEncounter, ID: 2}, false},
		{"tracker on NPC portrait", views.LiveKey(views.LiveEncounter, 1), hub.Topic{Kind: hub.TopicEntity, ID: 9}, true},
		{"player combat on NPC portrait", views.LiveKey(views.LivePlayerCombat, 1), hub.Topic{Kind: hub.TopicEntity, ID: 9}, true},
		{"player combat on any encounter", views.LiveKey(views.LivePlayerCombat, 1), hub.Topic{Kind: hub.TopicEncounter, ID: 9}, true},
		{"note list on a note", views.LiveKey(views.LiveNotes, 1), hub.Topic{Kind: hub.TopicNote, ID: 1}, false},
		{"entity on another entity", views.LiveKey(views.LiveEntity, 1), hub.Topic{Kind: hub.TopicEntity, ID: 2}, false},
		{"unknown view", "bogus:1", hub.Topic{Kind: hub.TopicCampaign, ID: 1}, false},
	}
	for _, tt := range tests {
		if got := a.Affects(tt.view, []hub.Topic{tt.topic}); got != tt.want {
			t.Errorf("%s: Affects = %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestLiveRefetchStrips checks the HTML a view refetches with, as a player.
func TestLiveRefetchStrips(t *testing.T) {
	f := newLiveFixture(t)
	player := f.token(t, enum.RolePlayer)
	get := func(view string) *httptest.ResponseRecorder {
		return f.do(t, http.MethodGet, "/live?v="+url.QueryEscape(view), player, nil, true)
	}
	body := func(view string) string { return get(view).Body.String() }
	if b := body(views.LiveKey(views.LiveEntity, f.entity)); !strings.Contains(b, "Innkeeper") || strings.Contains(b, "lich") {
		t.Errorf("NPC refetch leaks or misses fields: %s", b)
	}
	// The NPC list shows a portrait only once its look is revealed.
	mask, _ := f.q().CreateImage(t.Context(), sqlc.CreateImageParams{CampaignID: f.campaignID, Name: "mask", Url: sql.NullString{String: "https://example.com/mask.png", Valid: true}})
	look := sqlc.ToggleEntityImageRevealedParams{EntityID: f.entity, ImageID: mask.ID}
	if err := f.q().LinkEntityImage(t.Context(), sqlc.LinkEntityImageParams{EntityID: f.entity, ImageID: mask.ID}); err != nil {
		t.Fatal(err)
	}
	src := views.ImageSrc(mask.ID)
	if b := body(views.LiveKey(views.LiveEntities, f.campaignID)); !strings.Contains(b, "Innkeeper") || strings.Contains(b, src) {
		t.Errorf("NPC list leaks a hidden portrait or misses the role: %s", b)
	}
	if err := f.q().ToggleEntityImageRevealed(t.Context(), look); err != nil {
		t.Fatal(err)
	}
	if b := body(views.LiveKey(views.LiveEntities, f.campaignID)); !strings.Contains(b, src) {
		t.Errorf("NPC list misses a revealed portrait: %s", b)
	}
	if err := f.q().SetCombatantMarker(t.Context(), sqlc.SetCombatantMarkerParams{ID: f.ogre, Marker: "mini 5"}); err != nil {
		t.Fatal(err)
	}
	if b := body(views.LiveKey(views.LivePlayerCombat, f.campaignID)); !strings.Contains(b, "Ogre") || strings.Contains(b, "59/59") || !strings.Contains(b, "mini 5") {
		t.Errorf("combat refetch leaks monster HP, or misses the monster or its table marker")
	}
	// A running fight without the player, and one with them still being
	// prepared, stay hidden even when asked for by id.
	brawl, _ := f.q().CreateEncounter(t.Context(), sqlc.CreateEncounterParams{CampaignID: f.campaignID, Name: "Brawl"})
	if err := f.q().UpdateEncounterState(t.Context(), sqlc.UpdateEncounterStateParams{ID: brawl.ID, Status: "active", Round: 1}); err != nil {
		t.Fatal(err)
	}
	siege, _ := f.q().CreateEncounter(t.Context(), sqlc.CreateEncounterParams{CampaignID: f.campaignID, Name: "Siege"})
	if _, err := f.q().CreateCombatant(t.Context(), sqlc.CreateCombatantParams{EncounterID: siege.ID, CharacterID: nullID(f.characterID), Name: "Alice"}); err != nil {
		t.Fatal(err)
	}
	for _, want := range []int64{0, brawl.ID, siege.ID} {
		if b := body(views.LiveKey(views.LivePlayerCombat, f.campaignID, want)); !strings.Contains(b, "Ambush") || strings.Contains(b, "Brawl") || strings.Contains(b, "Siege") {
			t.Errorf("player combat asking for %d shows another fight or misses their own", want)
		}
	}
	// A player reads the GM's shared note but may not manage it from the index.
	shared, _ := f.q().CreateNote(t.Context(), sqlc.CreateNoteParams{CampaignID: f.campaignID, Kind: "generic", Title: "Rumours"})
	if _, err := f.q().CreateBlock(t.Context(), sqlc.CreateBlockParams{NoteID: shared.ID, Visibility: "players"}); err != nil {
		t.Fatal(err)
	}
	if b := body(views.LiveKey(views.LiveNotes, f.campaignID)); !strings.Contains(b, "Rumours") || strings.Contains(b, fmt.Sprintf("/notes/%d/rename", shared.ID)) {
		t.Errorf("notes refetch misses the shared note or lets the player manage it")
	}
	// A hidden note clears the page that showed it, naming nothing.
	if b := body(views.LiveKey(views.LiveNote, f.gmNote)); !strings.Contains(b, "not available") || strings.Contains(b, "Plans") {
		t.Errorf("hidden note refetch leaks or does not clear: %s", b)
	}
	for _, view := range []string{views.LiveKey(views.LiveRoster, f.campaignID), views.LiveKey(views.LiveEncounter, f.encounter)} {
		if w := get(view); w.Code != http.StatusNoContent {
			t.Errorf("%s: %d, want nothing", view, w.Code)
		}
	}
}

// outrankPin gives the fixture's character a feat pinned to a copy that a
// newer source outranks, so its sheet shows the sources alert.
func (f liveFixture) outrankPin(t *testing.T) {
	t.Helper()
	f.source(t, oldSrc, "2014")
	f.feat(t, oldSrc, "tough", "old")
	f.setSheet(t, dnd5e.Sheet{Feats: []string{"Tough"}, FeatSources: map[string]string{"tough": oldSrc}})
	f.source(t, newSrc, "2024")
	f.feat(t, newSrc, "tough", "new")
}

// TestLiveSourcesAlert checks who sees the "Review sources" alert on a sheet
// with an outranked pin: the owner and the GM, not another player.
func TestLiveSourcesAlert(t *testing.T) {
	f := newLiveFixture(t)
	f.outrankPin(t)
	view := views.LiveKey(views.LiveCharacter, f.characterID)
	for name, cookie := range map[string]string{"owner": f.token(t, enum.RolePlayer), "GM": f.token(t, enum.RoleAdmin)} {
		b := f.do(t, http.MethodGet, "/live?v="+url.QueryEscape(view), cookie, nil, true).Body.String()
		if !strings.Contains(b, "1 source to review") || !strings.Contains(b, "Review sources") {
			t.Errorf("%s misses the alert: %s", name, b)
		}
	}
	if _, ok := f.state(t, f.outsider, f.outsiderCampaign, view); ok {
		t.Error("another player can load the sheet")
	}
}

var (
	liveAttr = regexp.MustCompile(`data-live="([^"]+)"`)
	bindAttr = regexp.MustCompile(`data-bind="([^"]+)"`)
)

// TestLiveBindsMatchState renders pages with every bound value showing and
// checks that the keys the markup binds are exactly the hot keys of the
// page's live states. A bound key missing from the state would never update;
// a state key nothing binds would refetch on every change.
func TestLiveBindsMatchState(t *testing.T) {
	f := newLiveFixture(t)
	f.outrankPin(t) // the sheets show the sources alert, which binds nothing
	player := f.token(t, enum.RolePlayer)
	gm := f.token(t, enum.RoleAdmin)
	for _, c := range []string{"maxhp =12", "hp =0", "temp 3", "slot max 1 2", "exhaustion +1", "toggle prone", "deathsave fail", "feature add Lucky", "feature set Lucky uses=prof"} {
		if w := f.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/cmd", f.characterID), player, url.Values{"cmd": {c}}, true); w.Code != http.StatusOK {
			t.Fatalf("%s: %d %s", c, w.Code, w.Body)
		}
	}
	// A 5etools import is running, so the compendium binds its progress.
	f.app.fiveTools.start(3)
	f.app.fiveTools.update(func(j *views.FiveToolsJob) { j.Current = "Test Book" })
	tests := []struct {
		name, path, cookie string
		viewer             perm.Viewer
	}{
		{"GM compendium import", "/compendium?import=" + views.ImportTabFiveTools, gm, perm.GM},
		{"player sheet", fmt.Sprintf("/characters/%d", f.characterID), player, perm.Player(f.characterID)},
		{"GM sheet", fmt.Sprintf("/characters/%d", f.characterID), gm, perm.GM},
		{"player combat", fmt.Sprintf("/campaigns/%d/combat", f.campaignID), player, perm.Player(f.characterID)},
		{"GM combat", fmt.Sprintf("/campaigns/%d/combat?encounter=%d", f.campaignID, f.encounter), gm, perm.GM},
		{"GM overview", fmt.Sprintf("/campaigns/%d", f.campaignID), gm, perm.GM},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			page := f.do(t, http.MethodGet, tt.path, tt.cookie, nil, false).Body.String()
			var bound []string
			for _, m := range bindAttr.FindAllStringSubmatch(page, -1) {
				for _, spec := range strings.Fields(m[1]) {
					bound = append(bound, spec[strings.Index(spec, "=")+1:])
				}
			}
			var hot []string
			for _, m := range liveAttr.FindAllStringSubmatch(page, -1) {
				st, ok := f.state(t, tt.viewer, f.campaignID, m[1])
				if !ok {
					t.Fatalf("page shows %s but its state is not served", m[1])
				}
				for k := range st {
					if k != hub.HashKey {
						hot = append(hot, k)
					}
				}
			}
			if len(bound) == 0 {
				t.Fatal("no bound elements on the page")
			}
			slices.Sort(bound)
			slices.Sort(hot)
			if bound, hot = slices.Compact(bound), slices.Compact(hot); !slices.Equal(bound, hot) {
				t.Errorf("bound %v\nstate %v", bound, hot)
			}
		})
	}
}

// TestShapeHashSkipsRandomIDs pins shapeHash to hashing the markup with
// every utils.RandomID removed, as the regexp it replaced did.
func TestShapeHashSkipsRandomIDs(t *testing.T) {
	ids := regexp.MustCompile(`id-[A-Z2-7]{26}`)
	id := "id-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, in := range []string{
		"",
		"no ids",
		`<div id="` + id + `">x</div>`,
		id + id,
		"id-id-" + id,
		"id-ABC",
		"id-abcdefghijklmnopqrstuvwxyz",
		`aria-labelledby="` + id + `" data-x="id-2345672345672345672345672"`,
		"id-ABCDEFGHIJKLMNOPQRSTUVWXY1",
	} {
		h := fnv.New64a()
		_, _ = h.Write(ids.ReplaceAll([]byte(in), nil))
		if got, want := shapeHash([]byte(in)), strconv.FormatUint(h.Sum64(), 36); got != want {
			t.Errorf("shapeHash(%q) = %s, want %s", in, got, want)
		}
	}
}

// TestLiveHotKeepsHash checks that changes the page patches in place leave
// the hash alone (else every hp change would refetch the whole view), and
// that a change the page cannot patch moves it.
func TestLiveHotKeepsHash(t *testing.T) {
	f := newLiveFixture(t)
	player := f.token(t, enum.RolePlayer)
	gm := f.token(t, enum.RoleAdmin)
	sheetCmd := func(c string) {
		if w := f.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/cmd", f.characterID), player, url.Values{"cmd": {c}}, true); w.Code != http.StatusOK {
			t.Fatalf("%s: %d %s", c, w.Code, w.Body)
		}
	}
	post := func(path string, form url.Values) {
		if w := f.do(t, http.MethodPost, path, gm, form, true); w.Code >= http.StatusBadRequest {
			t.Fatalf("%s: %d %s", path, w.Code, w.Body)
		}
	}
	for _, c := range []string{"maxhp =12", "slot max 1 2", "feature add Lucky", "feature set Lucky uses=prof"} {
		sheetCmd(c)
	}
	rows, err := f.q().ListCombatants(t.Context(), f.encounter)
	if err != nil {
		t.Fatal(err)
	}
	for i, r := range rows { // turns need everyone's initiative
		if err := f.q().SetCombatantInitiative(t.Context(), sqlc.SetCombatantInitiativeParams{ID: r.ID, Initiative: sql.NullInt64{Int64: int64(10 + i), Valid: true}}); err != nil {
			t.Fatal(err)
		}
	}
	watched := map[string]perm.Viewer{
		views.LiveKey(views.LiveCharacter, f.characterID): perm.Player(f.characterID),
		views.LiveKey(views.LiveEncounter, f.encounter):   perm.GM,
	}
	hashes := func() map[string]any {
		out := map[string]any{}
		for view, v := range watched {
			st, _ := f.state(t, v, f.campaignID, view)
			out[view] = st[hub.HashKey]
		}
		return out
	}
	before := hashes()
	for name, step := range map[string]func(){
		"hp":        func() { sheetCmd("hp -3") },
		"temp":      func() { sheetCmd("temp 2") },
		"slot":      func() { sheetCmd("slot use 1") },
		"use":       func() { sheetCmd("use Lucky") },
		"insp":      func() { sheetCmd("inspiration") },
		"next turn": func() { post(fmt.Sprintf("/encounters/%d/next", f.encounter), nil) },
		"ogre hp":   func() { post(fmt.Sprintf("/combatants/%d/hp", f.ogre), url.Values{"amount": {"5"}}) },
	} {
		step()
		if after := hashes(); !maps.Equal(before, after) {
			t.Errorf("%s moved the hash: %v → %v", name, before, after)
			before = after
		}
	}
	// A toggled condition is bound on the sheet, but the tracker's rows list it.
	sheet := views.LiveKey(views.LiveCharacter, f.characterID)
	sheetCmd("toggle prone")
	if after := hashes(); after[sheet] != before[sheet] {
		t.Errorf("prone moved the sheet hash")
	}
	sheetCmd("item add Rope")
	post(fmt.Sprintf("/combatants/%d/marker", f.ogre), url.Values{"marker": {"red"}})
	for view, h := range hashes() {
		if h == before[view] {
			t.Errorf("%s: a new item or marker left the hash unchanged", view)
		}
	}
}
