package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
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

func newLiveFixture(t *testing.T) liveFixture {
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
		{"GM sees join requests outside a campaign", perm.GM, 0, views.LiveKey(views.LiveRequests, 0), true, nil, nil},
		{"player may not see join requests", player, f.campaignID, views.LiveKey(views.LiveRequests, 0), false, nil, nil},
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
}

// TestLiveRefetchStrips checks the HTML a view refetches with, as a player.
func TestLiveRefetchStrips(t *testing.T) {
	f := newLiveFixture(t)
	player := f.token(t, auth.RolePlayer)
	get := func(view string) *httptest.ResponseRecorder {
		return f.do(t, http.MethodGet, "/live?v="+url.QueryEscape(view), player, nil, true)
	}
	body := func(view string) string { return get(view).Body.String() }
	if b := body(views.LiveKey(views.LiveEntity, f.entity)); !strings.Contains(b, "Innkeeper") || strings.Contains(b, "lich") {
		t.Errorf("NPC refetch leaks or misses fields: %s", b)
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
	for _, view := range []string{views.LiveKey(views.LiveNote, f.gmNote), views.LiveKey(views.LiveRoster, f.campaignID), views.LiveKey(views.LiveEncounter, f.encounter)} {
		if w := get(view); w.Code != http.StatusNoContent {
			t.Errorf("%s: %d, want nothing", view, w.Code)
		}
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
	player := f.token(t, auth.RolePlayer)
	gm := f.token(t, auth.RoleAdmin)
	for _, c := range []string{"maxhp =12", "hp =0", "temp 3", "slot max 1 2", "exhaustion +1", "toggle prone", "deathsave fail", "feature add Lucky", "feature set Lucky uses=prof"} {
		if w := f.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/cmd", f.characterID), player, url.Values{"cmd": {c}}, true); w.Code != http.StatusOK {
			t.Fatalf("%s: %d %s", c, w.Code, w.Body)
		}
	}
	tests := []struct {
		name, path, cookie string
		viewer             perm.Viewer
	}{
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
