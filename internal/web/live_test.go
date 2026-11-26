package web

import (
	"bytes"
	"database/sql"
	"encoding"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"io/fs"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
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
	stranger, err := e.q().CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: other.ID, Name: "Bob", System: "dnd5e", SheetJson: "{}"})
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
		{"GM sees the starting sets", perm.GM, 0, views.LiveKey(views.LiveImport, views.StarterLiveCompendium), true, nil, nil},
		{"player may not see the starting sets", player, f.campaignID, views.LiveKey(views.LiveImport, views.StarterLiveCompendium), false, nil, nil},
		{"GM sees the source priority", perm.GM, 0, views.LiveKey(views.LiveSources, 0), true, nil, nil},
		{"GM sees a campaign's source priority", perm.GM, f.campaignID, views.LiveKey(views.LiveSources, f.campaignID), true, nil, nil},
		{"player may not see the source priority", player, f.campaignID, views.LiveKey(views.LiveSources, f.campaignID), false, nil, nil},
		{"GM sees the Sources card", perm.GM, f.campaignID, views.LiveKey(views.LiveCampaignSources, f.campaignID), true, nil, nil},
		{"player may not see the Sources card", player, f.campaignID, views.LiveKey(views.LiveCampaignSources, f.campaignID), false, nil, nil},
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
	// The tracker row's Dead flag is a hot key for everyone who sees the row,
	// the player's view included (a separate loop: the table's columns are positional).
	for _, tt := range []struct {
		name string
		v    perm.Viewer
		view string
	}{
		{"GM tracker", perm.GM, views.LiveKey(views.LiveEncounter, f.encounter)},
		{"player combat", player, views.LiveKey(views.LivePlayerCombat, f.campaignID)},
	} {
		st, _ := f.state(t, tt.v, f.campaignID, tt.view)
		for _, id := range []int64{f.ogre} {
			if _, has := st[fmt.Sprintf("c.%d.dead", id)]; !has {
				t.Errorf("%s: no dead key in %v", tt.name, st)
			}
		}
	}
	// Every live view is registered and has at least one row above.
	all := []string{
		views.LiveRoster, views.LiveRequests, views.LiveWaiting, views.LiveCharacter,
		views.LiveEncounters, views.LiveEncounter, views.LivePlayerCombat, views.LiveNotes,
		views.LiveNote, views.LiveEntities, views.LiveEntity, views.LiveImport,
		views.LiveSources, views.LiveCampaignSources,
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

// TestCombatHidesUnrevealedNPC checks that a player's turn order never names
// an NPC the GM has not revealed, nor shows its linked monster's picture,
// even after the NPC is deleted.
func TestCombatHidesUnrevealedNPC(t *testing.T) {
	f := newLiveFixture(t)
	ctx := t.Context()
	const pic = "https://example.com/vampire.png"
	if err := f.upsertEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: "monster", Slug: "vampire", Name: "Vampire", SourceKey: "test", DataJson: "{}"}); err != nil {
		t.Fatal(err)
	}
	m, err := f.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "vampire"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.q().SetCompendiumImage(ctx, sqlc.SetCompendiumImageParams{ImageUrl: sql.NullString{String: pic, Valid: true}, ID: m.ID}); err != nil {
		t.Fatal(err)
	}
	secret, err := f.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: f.campaignID, Kind: "npc", Name: "Vampire Mayor", FieldsJson: "{}", CompendiumRef: nullID(m.ID)})
	if err != nil {
		t.Fatal(err)
	}
	ids := map[string]int64{}
	for _, p := range []sqlc.CreateCombatantParams{
		{EncounterID: f.encounter, EntityID: nullID(secret.ID), MonsterRef: nullID(m.ID), Name: "Vampire Mayor"},
		{EncounterID: f.encounter, EntityID: nullID(f.entity), Name: "Mira"},
	} {
		cb, err := f.q().CreateCombatant(ctx, p)
		if err != nil {
			t.Fatal(err)
		}
		ids[p.Name] = cb.ID
	}
	enc, err := f.q().GetEncounter(ctx, f.encounter)
	if err != nil {
		t.Fatal(err)
	}
	player := perm.Player(f.characterID)
	row := func(t *testing.T, v perm.Viewer, combatantID int64) views.CombatantView {
		t.Helper()
		ev, err := f.app.encounterView(ctx, v, enc)
		if err != nil {
			t.Fatal(err)
		}
		i := slices.IndexFunc(ev.Combatants, func(c views.CombatantView) bool { return c.ID == combatantID })
		if i < 0 {
			t.Fatal("combatant not in the turn order")
		}
		return ev.Combatants[i]
	}
	tests := []struct {
		name       string
		viewer     perm.Viewer
		combatant  string
		want, wPic string
	}{
		{"player sees no hidden name or picture", player, "Vampire Mayor", hiddenNPCName, ""},
		{"player sees a revealed name", player, "Mira", "Mira", ""},
		{"GM sees the hidden name and picture", perm.GM, "Vampire Mayor", "Vampire Mayor", pic},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := row(t, tt.viewer, ids[tt.combatant]); got.Name != tt.want || got.Portrait != tt.wPic {
				t.Errorf("name, portrait = %q, %q, want %q, %q", got.Name, got.Portrait, tt.want, tt.wPic)
			}
		})
	}
	// Deleting the NPC leaves its combatant without a link: the stored name stays hidden too.
	gm := f.token(t, enum.RoleAdmin)
	if w := f.do(t, http.MethodPost, route.Path(route.EntityDelete, secret.ID), gm, url.Values{}, true); w.Code >= http.StatusBadRequest {
		t.Fatalf("delete NPC: %d %s", w.Code, w.Body)
	}
	if got := row(t, player, ids["Vampire Mayor"]); got.Name != hiddenNPCName {
		t.Errorf("after delete, player sees %q", got.Name)
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
		return f.do(t, http.MethodGet, route.Live+"?v="+url.QueryEscape(view), player, nil, true)
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
	if b := body(views.LiveKey(views.LiveNotes, f.campaignID)); !strings.Contains(b, "Rumours") || strings.Contains(b, route.Path(route.NoteRename, shared.ID)) {
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
		b := f.do(t, http.MethodGet, route.Live+"?v="+url.QueryEscape(view), cookie, nil, true).Body.String()
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
		if w := f.do(t, http.MethodPost, route.Path(route.CharacterCmd, f.characterID), player, url.Values{"cmd": {c}}, true); w.Code != http.StatusOK {
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
		{"GM compendium import", route.Compendium + "?import=" + views.ImportTabFiveTools, gm, perm.GM},
		{"GM empty compendium", route.Compendium, gm, perm.GM}, // the starting sets
		{"player sheet", route.Path(route.Character, f.characterID), player, perm.Player(f.characterID)},
		{"GM sheet", route.Path(route.Character, f.characterID), gm, perm.GM},
		{"player combat", route.Path(route.CampaignCombat, f.campaignID), player, perm.Player(f.characterID)},
		{"GM combat", fmt.Sprintf("%s?encounter=%d", route.Path(route.CampaignCombat, f.campaignID), f.encounter), gm, perm.GM},
		{"GM overview", route.Path(route.Campaign, f.campaignID), gm, perm.GM}, // the live Sources card too
	}
	// The source cards are live views, so the loop below checks them too.
	for path, kind := range map[string]string{route.Compendium + "?import=" + views.ImportTabFiveTools: views.LiveSources, route.Path(route.Campaign, f.campaignID): views.LiveCampaignSources} {
		if page := f.do(t, http.MethodGet, path, gm, nil, false).Body.String(); !strings.Contains(page, `data-live="`+kind+`:`) {
			t.Errorf("%s does not show the live %s view", path, kind)
		}
	}
	for _, dead := range []bool{false, true} { // the dead sheet swaps its buttons and the tracker its badge
		if dead {
			for range 2 {
				f.sheetCmd(t, player, "deathsave fail")
			}
		}
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s dead=%v", tt.name, dead), func(t *testing.T) {
				page := f.do(t, http.MethodGet, tt.path, tt.cookie, nil, false).Body.String()
				var bound []string
				for _, m := range bindAttr.FindAllStringSubmatch(page, -1) {
					for spec := range strings.FieldsSeq(m[1]) {
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
}

// TestShadcnPortalOwner pins the private field live.js follows from a
// dialog portaled to <body> back to its view; a bundle update that renames
// it would silently stop live values reaching opened dialogs.
func TestShadcnPortalOwner(t *testing.T) {
	names, err := fs.Glob(assetsFS, "assets/js/shadcn-templ-*.js")
	if err != nil || len(names) != 1 {
		t.Fatalf("shadcn bundle: %v %v", names, err)
	}
	js, err := assetsFS.ReadFile(names[0])
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(js, []byte("_templPortalOwner")) {
		t.Errorf("%s no longer sets _templPortalOwner; update within() in live.js", names[0])
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
	sheetCmd := func(c string) { f.sheetCmd(t, player, c) }
	post := func(path string, form url.Values) { f.postOK(t, gm, path, form) }
	sh := dnd5e.NewSheet() // a class, for hit dice
	sh.Classes = []dnd5e.ClassLevel{{Slug: "wizard", Name: "Wizard", Level: 1, HitDie: 6, Caster: dnd5e.CasterFull}}
	f.setSheet(t, sh)
	for _, c := range []string{"maxhp =12", "slot max 1 2", "feature add Lucky", "feature set Lucky uses=prof"} {
		sheetCmd(c)
	}
	f.readyTurns(t)
	hashes := func() map[string]string {
		out := map[string]string{}
		for view, v := range f.modelViews() {
			_, out[view] = f.hashes(t, v, view)
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
		"hit die":   func() { sheetCmd("hitdie use") },
		"next turn": func() { post(route.Path(route.EncounterAction, f.encounter, "next"), nil) },
		"ogre hp":   func() { post(route.Path(route.CombatantAction, f.ogre, "hp"), url.Values{"amount": {"5"}}) },
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
	// Dropping to 0 shows the death saves; from then on they are bound.
	sheetCmd("hp =0")
	before = hashes()
	sheetCmd("deathsave fail")
	sheetCmd("deathsave success")
	if after := hashes(); after[sheet] != before[sheet] {
		t.Errorf("death saves moved the sheet hash")
	}
	// Dying for good swaps the sheet's buttons and the tracker's badge: a refetch.
	before = hashes()
	sheetCmd("deathsave fail")
	sheetCmd("deathsave fail")
	for view, h := range hashes() {
		if h == before[view] {
			t.Errorf("%s: dying left the hash unchanged", view)
		}
	}
	before = hashes()
	sheetCmd("item add Rope")
	post(route.Path(route.CombatantAction, f.ogre, "marker"), url.Values{"marker": {"red"}})
	for view, h := range hashes() {
		if h == before[view] {
			t.Errorf("%s: a new item or marker left the hash unchanged", view)
		}
	}
}

// TestLiveTurnValsFollowTurn checks the turn buttons' bound payload: Next
// moves the turn without a refetch, so the GM tracker's turn.vals must name
// the new turn, or every later press would be refused as stale.
func TestLiveTurnValsFollowTurn(t *testing.T) {
	f := newLiveFixture(t)
	gm := f.token(t, enum.RoleAdmin)
	f.readyTurns(t)
	tracker := views.LiveKey(views.LiveEncounter, f.encounter)
	next := route.Path(route.EncounterAction, f.encounter, "next")
	vals := func() url.Values {
		st, _ := f.state(t, perm.GM, f.campaignID, tracker)
		s, _ := st["turn.vals"].(string)
		var m map[string]int64
		if err := json.Unmarshal([]byte(s), &m); err != nil {
			t.Fatalf("turn.vals %q: %v", st["turn.vals"], err)
		}
		form := url.Values{}
		for k, v := range m {
			form.Set(k, fmt.Sprint(v))
		}
		return form
	}
	first := vals()
	f.postOK(t, gm, next, first)
	f.postOK(t, gm, next, vals()) // the live payload, as the page now holds it
	if w := f.do(t, http.MethodPost, next, gm, first, true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("stale turn: %d, want %d", w.Code, http.StatusUnprocessableEntity)
	}
}

// TestLiveModelHashFollowsMarkup checks the JSON-hashed views against their
// markup: whenever a change shows on the page, it must move the hash too,
// or other devices would keep the old view. The page markup includes hot
// values, so no step may change one (TestLiveHotKeepsHash covers those).
func TestLiveModelHashFollowsMarkup(t *testing.T) {
	f := newLiveFixture(t)
	player := f.token(t, enum.RolePlayer)
	gm := f.token(t, enum.RoleAdmin)
	f.readyTurns(t)
	ogre := func(action string, form url.Values) func() {
		return func() { f.postOK(t, gm, route.Path(route.CombatantAction, f.ogre, action), form) }
	}
	steps := []struct {
		name string
		do   func()
	}{
		{"item", func() { f.sheetCmd(t, player, "item add Rope") }},
		{"rename", func() { f.sheetCmd(t, player, "set name Alicia") }},
		{"feature", func() { f.sheetCmd(t, player, "feature add Lucky") }},
		{"concentration", func() { f.sheetCmd(t, player, "concentrate Bless") }},
		{"effect", func() {
			f.postOK(t, player, route.Path(route.CharacterEffect, f.characterID), url.Values{"slug": {"custom"}, "name": {"Shield of Faith"}, "mod_target": {"ac"}, "mod_value": {"2"}})
		}},
		{"skill", func() { f.sheetCmd(t, player, "prof add skill athletics") }},
		{"death", func() {
			for range 3 {
				f.sheetCmd(t, player, "deathsave fail")
			}
		}},
		{"marker", ogre("marker", url.Values{"marker": {"red"}})},
		{"monster condition", ogre("condition", url.Values{"slug": {"blinded"}})},
		{"hp shown", ogre("visible", nil)},
		{"new NPC", func() {
			if _, err := f.q().CreateEntity(t.Context(), sqlc.CreateEntityParams{CampaignID: f.campaignID, Kind: "npc", Name: "Bram", FieldsJson: "{}"}); err != nil {
				t.Fatal(err)
			}
		}},
		{"new combatant", func() {
			if _, err := f.q().CreateCombatant(t.Context(), sqlc.CreateCombatantParams{EncounterID: f.encounter, Name: "Wolf", Hp: 11, MaxHp: 11, Ac: 13}); err != nil {
				t.Fatal(err)
			}
		}},
	}
	type pair struct{ markup, model string }
	snapshot := func() map[string]pair {
		out := map[string]pair{}
		for view, v := range f.modelViews() {
			markup, model := f.hashes(t, v, view)
			out[view] = pair{markup, model}
		}
		return out
	}
	shown := map[string]bool{}
	before := snapshot()
	for _, s := range steps {
		s.do()
		after := snapshot()
		for view, a := range after {
			b := before[view]
			if a.markup != b.markup {
				shown[view] = true
				if a.model == b.model {
					t.Errorf("%s: %s changed the page but not the hash", s.name, view)
				}
			}
		}
		before = after
	}
	for view := range f.modelViews() {
		if !shown[view] {
			t.Errorf("%s: no step changed its page", view)
		}
	}
}

// TestLiveModelsEncodeWhole guards the JSON-hashed models: a field JSON
// skips would be shown but not hashed, so its changes would never refetch.
func TestLiveModelsEncodeWhole(t *testing.T) {
	custom := []reflect.Type{reflect.TypeFor[json.Marshaler](), reflect.TypeFor[encoding.TextMarshaler]()}
	seen := map[reflect.Type]bool{}
	var walk func(ty reflect.Type, path string)
	walk = func(ty reflect.Type, path string) {
		if seen[ty] || ty == reflect.TypeFor[time.Time]() || ty == reflect.TypeFor[json.RawMessage]() { // both encode whole
			return
		}
		seen[ty] = true
		for _, m := range custom {
			if ty.Implements(m) || reflect.PointerTo(ty).Implements(m) {
				t.Errorf("%s (%s) encodes itself: check that it keeps every shown value", path, ty)
				return
			}
		}
		switch ty.Kind() {
		case reflect.Pointer, reflect.Slice, reflect.Array:
			walk(ty.Elem(), path+"[]")
		case reflect.Map:
			walk(ty.Key(), path+"{key}")
			walk(ty.Elem(), path+"{}")
		case reflect.Struct:
			for fd := range ty.Fields() {
				if !fd.IsExported() || fd.Tag.Get("json") == "-" {
					t.Errorf("%s.%s is not encoded", path, fd.Name)
				}
				walk(fd.Type, path+"."+fd.Name)
			}
		case reflect.Func, reflect.Chan, reflect.Interface, reflect.UnsafePointer:
			t.Errorf("%s (%s) does not encode by value", path, ty)
		default: // numbers, strings, bools
		}
	}
	walk(reflect.TypeFor[views.SheetView](), "SheetView")
	walk(reflect.TypeFor[views.EncounterView](), "EncounterView")
}

// modelViews are the live views hashed by model (liveView.model), with the
// viewer each test watches them as.
func (f liveFixture) modelViews() map[string]perm.Viewer {
	return map[string]perm.Viewer{
		views.LiveKey(views.LiveCharacter, f.characterID): perm.Player(f.characterID),
		views.LiveKey(views.LiveEncounter, f.encounter):   perm.GM,
	}
}

// hashes returns a view's markup hash (as the page shows it) and its live
// state hash.
func (f liveFixture) hashes(t *testing.T, v perm.Viewer, view string) (markup, state string) {
	t.Helper()
	lv, err := f.app.loadLive(t.Context(), viewerCtx{viewer: v, campaignID: f.campaignID}, view)
	if err != nil || !lv.ok {
		t.Fatal(view, lv.ok, err)
	}
	var buf bytes.Buffer
	if err := lv.html.Render(t.Context(), &buf); err != nil {
		t.Fatal(err)
	}
	st, _, err := lv.state(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	return shapeHash(buf.Bytes()), st[hub.HashKey].(string)
}

// sheetCmd runs a sheet command on the fixture's character.
func (f liveFixture) sheetCmd(t *testing.T, cookie, cmd string) {
	t.Helper()
	if w := f.do(t, http.MethodPost, route.Path(route.CharacterCmd, f.characterID), cookie, url.Values{"cmd": {cmd}}, true); w.Code != http.StatusOK {
		t.Fatalf("%s: %d %s", cmd, w.Code, w.Body)
	}
}

// postOK posts form as htmx and fails the test on a 4xx or 5xx.
func (f liveFixture) postOK(t *testing.T, cookie, path string, form url.Values) {
	t.Helper()
	if w := f.do(t, http.MethodPost, path, cookie, form, true); w.Code >= http.StatusBadRequest {
		t.Fatalf("%s: %d %s", path, w.Code, w.Body)
	}
}

// readyTurns gives every combatant an initiative, which turns need.
func (f liveFixture) readyTurns(t *testing.T) {
	t.Helper()
	rows, err := f.q().ListCombatants(t.Context(), f.encounter)
	if err != nil {
		t.Fatal(err)
	}
	for i, r := range rows {
		if err := f.q().SetCombatantInitiative(t.Context(), sqlc.SetCombatantInitiativeParams{ID: r.ID, Initiative: sql.NullInt64{Int64: int64(10 + i), Valid: true}}); err != nil {
			t.Fatal(err)
		}
	}
}
