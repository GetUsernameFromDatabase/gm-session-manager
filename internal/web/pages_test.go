package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"slices"
	"strings"
	"testing"

	"github.com/a-h/templ"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

func TestContentSecurityPolicy(t *testing.T) {
	e := newTestEnv(t)
	w := e.do(t, http.MethodGet, "/", e.token(t, enum.RoleAdmin), nil, false)
	got := w.Header().Get("Content-Security-Policy")
	if got != contentSecurityPolicy {
		t.Errorf("Content-Security-Policy = %q, want %q", got, contentSecurityPolicy)
	}
	if !strings.HasPrefix(got, "default-src 'self'") {
		t.Errorf("policy %q does not start with default-src 'self'", got)
	}
	for _, d := range []string{"script-src 'self'", "frame-ancestors 'none'", "object-src 'none'"} {
		if !strings.Contains(got, d) {
			t.Errorf("policy %q lacks %s", got, d)
		}
	}
}

// TestPagesRender requests every page as the GM and as a player and
// checks each one renders.
func TestPagesRender(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	e.allowEntrySources(t)
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)

	note, _ := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "game_session", Title: "Session 1"})
	if _, err := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: "players"}); err != nil {
		t.Fatal(err)
	}
	ent, _ := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Mira", IsRevealed: true, FieldsJson: "{}"})
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	goblin, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	if err != nil {
		t.Fatal(err)
	}
	for _, kind := range []string{"pc", "monster", "entity"} {
		id := map[string]int64{"pc": e.characterID, "monster": goblin.ID, "entity": ent.ID}[kind]
		w := e.do(t, http.MethodPost, route.Path(route.EncounterCombatants, enc.ID), gm, url.Values{"kind": {kind}, "id": {fmt.Sprint(id)}}, true)
		if w.Code != http.StatusOK {
			t.Fatalf("add %s: %d %s", kind, w.Code, w.Body)
		}
	}
	start := route.Path(route.EncounterAction, enc.ID, "start")
	if w := e.do(t, http.MethodPost, start, gm, nil, true); w.Code == http.StatusOK {
		t.Fatal("start ran without initiatives")
	}
	rows, _ := e.q().ListCombatants(ctx, enc.ID)
	for i, r := range rows {
		if w := e.do(t, http.MethodPost, route.Path(route.CombatantAction, r.ID, "initiative"), gm, url.Values{"initiative": {fmt.Sprint(10 + i)}}, true); w.Code != http.StatusOK {
			t.Fatalf("initiative: %d %s", w.Code, w.Body)
		}
		if r.CharacterID.Valid { // a PC's resistance halves damage of its type only
			hpURL, id := route.Path(route.CombatantAction, r.ID, "hp"), r.CharacterID.Int64
			if err := e.app.updateCharacter(ctx, id, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
				dc.Sheet.Resistances, dc.HP, dc.MaxHP = []string{"fire"}, 30, 30
				return nil
			}); err != nil {
				t.Fatal(err)
			}
			for _, tt := range []struct {
				kind string
				hp   int64
			}{{"fire", 26}, {" ", 17}, {"cold", 8}} {
				e.do(t, http.MethodPost, hpURL, gm, url.Values{"op": {"damage"}, "amount": {"9"}, "type": {tt.kind}}, true)
				if ch, _ := e.q().GetCharacter(ctx, id); ch.Hp != tt.hp {
					t.Fatalf("%q damage: HP %d, want %d", tt.kind, ch.Hp, tt.hp)
				}
			}
			if w := e.do(t, http.MethodPost, hpURL, gm, url.Values{"op": {"damage"}, "amount": {"9"}, "type": {"pizza"}}, true); w.Code == http.StatusOK {
				t.Fatal("unknown damage type accepted")
			}
			e.do(t, http.MethodPost, hpURL, gm, url.Values{"op": {"heal"}, "amount": {"30"}, "type": {"fire"}}, true)
		}
		if r.EntityID.Valid && !r.MonsterRef.Valid { // Mira has no stats: max HP starts her at full
			e.do(t, http.MethodPost, route.Path(route.CombatantAction, r.ID, "hp"), gm, url.Values{"op": {"set"}, "max_hp": {"12"}}, true)
			e.do(t, http.MethodPost, route.Path(route.CombatantAction, r.ID, "hp"), gm, url.Values{"op": {"damage"}, "amount": {"15"}}, true)
			if cb, _ := e.q().GetCombatant(ctx, r.ID); cb.MaxHp != 12 || cb.Hp != -3 {
				t.Fatalf("Mira HP %d/%d, want -3/12", cb.Hp, cb.MaxHp)
			}
			// Set fills only the given fields; HP is capped at max.
			e.do(t, http.MethodPost, route.Path(route.CombatantAction, r.ID, "hp"), gm, url.Values{"op": {"set"}, "hp": {"20"}, "temp_hp": {"4"}}, true)
			if cb, _ := e.q().GetCombatant(ctx, r.ID); cb.Hp != 12 || cb.MaxHp != 12 || cb.TempHp != 4 {
				t.Fatalf("Mira HP %d/%d+%d, want 12/12+4", cb.Hp, cb.MaxHp, cb.TempHp)
			}
			if w := e.do(t, http.MethodPost, route.Path(route.CombatantAction, r.ID, "hp"), gm, url.Values{"op": {"set"}}, true); w.Code == http.StatusOK {
				t.Fatal("empty set succeeded")
			}
			e.do(t, http.MethodPost, route.Path(route.CombatantAction, r.ID, "hp"), gm, url.Values{"op": {"damage"}, "amount": {"19"}}, true)
		}
	}
	if w := e.do(t, http.MethodPost, start, gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("start: %d %s", w.Code, w.Body)
	}
	next := route.Path(route.EncounterAction, enc.ID, "next")
	if w := e.do(t, http.MethodPost, next, gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("next: %d %s", w.Code, w.Body)
	}

	// An effect on the PC shows on the sheet and in the tracker.
	effectMods, err := json.Marshal(dnd5e.EffectMods("Shield of Faith", []dnd5e.CustomMod{{Target: dnd5e.ModAC, Value: 2}}))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := e.q().CreateCondition(ctx, sqlc.CreateConditionParams{CharacterID: nullID(e.characterID), Slug: "custom", Note: "Shield of Faith", EndsOn: enum.EffectEndLongRest, ModsJson: string(effectMods)}); err != nil {
		t.Fatal(err)
	}

	cb := route.Path(route.CampaignCompendium, e.campaignID) // the campaign's compendium
	pages := []struct {
		path      string
		gm, other bool // who may load it
		contains  string
	}{
		{"/", true, false, "Campaigns"},
		{route.Path(route.Campaign, e.campaignID), true, false, "Alice"},
		{route.Path(route.Character, e.characterID), true, true, "Alice"},
		{route.Path(route.Character, e.characterID), true, true, "Add effect to Alice"},
		{route.Path(route.Character, e.characterID), true, true, "Armor class +2"},
		{route.Path(route.Character, e.characterID), true, true, "until long rest"},
		{route.Path(route.CampaignCombat, e.campaignID), true, true, "Armor class +2"},
		{route.Path(route.CampaignCombat, e.campaignID), true, true, "until long rest"},
		{route.Path(route.CharacterLevelup, e.characterID) + "?class=wizard", false, true, "Wizard 1"},
		{route.Path(route.CampaignCombat, e.campaignID), true, true, "Goblin Warrior"},
		{route.Path(route.CampaignNotes, e.campaignID) + "?q=Session", true, true, "Session 1"},
		{route.Path(route.Note, note.ID), true, true, "Session 1"},
		{route.Path(route.CampaignEntities, e.campaignID), true, true, "Mira"},
		{route.Path(route.Entity, ent.ID), true, true, "Mira"},
		{route.Path(route.CampaignSwitcher, e.campaignID) + "?q=mi", true, true, "Mira"},
		{cb + "?q=fire", true, true, "Fire Bolt"},
		{cb, true, true, "Goblin Warrior"}, // first batch of every kind
		{fmt.Sprintf("%s/entries/%d", cb, goblin.ID), true, true, "Scimitar"},
		{fmt.Sprintf("%s/entries/%d?preview=1", cb, goblin.ID), true, true, "Scimitar"},
		{cb + "?kind=spell", true, true, "Fire Bolt"},
		{cb + "?kind=spell", true, true, ">sorcerer</span>"},                       // class badges on spell rows
		{cb, true, true, `href="` + cb + `?kind=spell"`},                           // kind counts link to their list
		{cb + "?kind=spell&q=fire", true, true, `href="` + cb + `?q=fire"`},        // the active kind links back to all
		{cb + "?kind=spell", true, true, `href="` + cb + `?kind=monster%2Cspell"`}, // inactive kinds join the filter
		{cb + "?kind=monster,spell", true, true, "Goblin Warrior"},
		{cb + "?kind=monster,spell", true, true, "Fire Bolt"},
		{cb + "?after_name=Fire+Bolt&after_id=0", true, true, "Fire Bolt"},           // cursor is (name, id): same name, later id
		{cb + "?after_name=Fire+Bolt&after_id=999999", true, true, "Goblin Warrior"}, // next batch after the cursor
		{cb + route.CompendiumSearch + "?q=gob", true, true, "Goblin Warrior"},
		{cb + route.CompendiumSearch + "?kind=monster&q=gob&encounter=1", true, true, "Goblin Warrior"},
		{cb + route.CompendiumHomebrewNew, true, true, "New homebrew"}, // players write their own spells and items
	}
	if w := e.do(t, http.MethodGet, cb+route.CompendiumRandom+"?kind=monster", player, nil, false); w.Code != http.StatusSeeOther || !strings.HasPrefix(w.Header().Get("Location"), cb+route.Path(route.CompendiumEntry, "")) {
		t.Errorf("random: %d %q", w.Code, w.Header().Get("Location"))
	}
	if w := e.do(t, http.MethodGet, cb+route.CompendiumRandom+"?kind=monster", player, nil, true); !strings.HasPrefix(w.Header().Get("HX-Redirect"), cb+route.Path(route.CompendiumEntry, "")) {
		t.Errorf("htmx random: %d %q", w.Code, w.Header().Get("HX-Redirect"))
	}
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, cb+"?kind=spell", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", views.CompendiumResultsID)
	req.RemoteAddr = lanAddr
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: player})
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	if body := w.Body.String(); w.Code != http.StatusOK || !strings.Contains(body, "Fire Bolt") || strings.Contains(body, "<html") ||
		!strings.Contains(body, `href="`+cb+`?kind=monster%2Cspell"`) { // badges re-rendered without a stale query
		t.Errorf("results fragment: %d %s", w.Code, body)
	}
	for _, p := range pages {
		for _, as := range []struct {
			name   string
			cookie string
			ok     bool
		}{{"gm", gm, p.gm}, {"player", player, p.other}} {
			t.Run(as.name+p.path, func(t *testing.T) {
				w := e.do(t, http.MethodGet, p.path, as.cookie, nil, false)
				if !as.ok {
					if w.Code == http.StatusOK && p.path != "/" {
						t.Errorf("status 200 for a page this viewer may not load")
					}
					return
				}
				if w.Code != http.StatusOK {
					t.Fatalf("status %d: %s", w.Code, w.Body)
				}
				if !strings.Contains(w.Body.String(), p.contains) {
					t.Errorf("page lacks %q", p.contains)
				}
				// A spread event handler given a string renders blank: use templ.JSUnsafeFuncCall.
				if strings.Contains(w.Body.String(), templ.FailedSanitizationJS) {
					t.Errorf("page has an event handler templ blanked")
				}
			})
		}
	}
}

func TestMonsterImageLink(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	goblin, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	spell, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "spell", Slug: "fire-bolt"})
	page := route.Path(route.CampaignCompendium+route.CompendiumEntry, e.campaignID, goblin.ID)
	if w := e.do(t, http.MethodGet, page, player, nil, false); !strings.Contains(w.Body.String(), compendium.MonsterImagePrefix+"goblin-warrior.png") {
		t.Errorf("player page lacks the bound picture: %s", w.Body)
	}
	if body := e.do(t, http.MethodGet, page, player, nil, false).Body.String(); strings.Contains(body, `name="image_url"`) || strings.Contains(body, "data-pic") {
		t.Error("player page has the picture form")
	}
	if body := e.do(t, http.MethodGet, page, gm, nil, false).Body.String(); !strings.Contains(body, `name="image_url"`) || !strings.Contains(body, "data-pic") {
		t.Errorf("gm page lacks the picture form: %s", body)
	}

	cases := []struct {
		name   string
		cookie string
		id     int64
		link   string
		ok     bool
		want   string // goblin's link afterwards
	}{
		{"gm sets", gm, goblin.ID, "https://example.com/g.png", true, "https://example.com/g.png"},
		{"player refused", player, goblin.ID, "https://evil.example/x.png", false, "https://example.com/g.png"},
		{"http refused", gm, goblin.ID, "http://example.com/g.png", false, "https://example.com/g.png"},
		{"javascript refused", gm, goblin.ID, "javascript:alert(1)", false, "https://example.com/g.png"},
		{"not a monster", gm, spell.ID, "https://example.com/s.png", false, "https://example.com/g.png"},
		{"gm clears", gm, goblin.ID, "", true, ""},
	}
	for _, c := range cases {
		w := e.do(t, http.MethodPost, route.Compendium+route.Path(route.CompendiumEntryImage, c.id), c.cookie, url.Values{"image_url": {c.link}}, false)
		if ok := w.Code == http.StatusSeeOther; ok != c.ok {
			t.Errorf("%s: status %d", c.name, w.Code)
		}
		if got, _ := e.q().GetCompendiumEntry(ctx, goblin.ID); got.ImageUrl.String != c.want {
			t.Errorf("%s: link %q, want %q", c.name, got.ImageUrl.String, c.want)
		}
	}
	if s, _ := e.q().GetCompendiumEntry(ctx, spell.ID); s.ImageUrl.Valid {
		t.Errorf("spell got a link %q", s.ImageUrl.String)
	}
}

func joinHabitats(hs []enum.Habitat, sep string) string {
	names := make([]string, len(hs))
	for i, h := range hs {
		names[i] = string(h)
	}
	return strings.Join(names, sep)
}

func TestMonsterHabitats(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	goblin, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	spell, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "spell", Slug: "fire-bolt"})
	giant, err := e.q().CreateHomebrew(ctx, sqlc.CreateHomebrewParams{Kind: "monster", Slug: "hill-giant", Name: "Hill Giant", DataJson: "{}", ContentHash: "x"})
	if err != nil {
		t.Fatal(err)
	}
	habitats := func(id int64) string {
		h, _ := e.q().ListEntryHabitats(ctx, id)
		slices.Sort(h)
		return joinHabitats(h, ",")
	}
	cell := func(id int64, habitat string, on bool) url.Values {
		v := url.Values{"entry_id": {fmt.Sprint(id)}, "habitat": {habitat}}
		if on {
			v.Set("on", "true")
		}
		return v
	}
	batch := func(q, habitat string, on bool) url.Values {
		return url.Values{"q": {q}, "habitat": {habitat}, "on": {fmt.Sprint(on)}}
	}
	anyForm := func(id int64) url.Values { return url.Values{"entry_id": {fmt.Sprint(id)}} }
	every := slices.Clone(compendium.Habitats)
	slices.Sort(every)
	all := joinHabitats(every, ",")

	// Every listed habitat passes the form's oneof and the table's CHECK.
	for _, h := range compendium.Habitats {
		if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumHabitatCell, gm, cell(goblin.ID, string(h), true), false); w.Code != http.StatusOK {
			t.Fatalf("habitat %s: status %d", h, w.Code)
		}
	}
	if got := habitats(goblin.ID); strings.Count(got, ",") != len(compendium.Habitats)-1 {
		t.Fatalf("goblin habitats %q", got)
	}
	for _, h := range compendium.Habitats {
		e.do(t, http.MethodPost, route.Compendium+route.CompendiumHabitatCell, gm, cell(goblin.ID, string(h), false), false)
	}

	cases := []struct {
		name   string
		cookie string
		path   string
		form   url.Values
		ok     bool
		goblin string // habitats afterwards
		giant  string
	}{
		{"gm ticks", gm, route.Compendium + route.CompendiumHabitatCell, cell(goblin.ID, "Forest", true), true, "Forest", ""},
		{"gm ticks again", gm, route.Compendium + route.CompendiumHabitatCell, cell(goblin.ID, "Forest", true), true, "Forest", ""},
		{"player refused", player, route.Compendium + route.CompendiumHabitatCell, cell(goblin.ID, "Hill", true), false, "Forest", ""},
		{"unknown habitat", gm, route.Compendium + route.CompendiumHabitatCell, cell(goblin.ID, "Moon", true), false, "Forest", ""},
		{"not a monster", gm, route.Compendium + route.CompendiumHabitatCell, cell(spell.ID, "Hill", true), false, "Forest", ""},
		{"batch adds to matches only", gm, route.Compendium + route.CompendiumHabitatBatch, batch("HILL", "Hill", true), true, "Forest", "Hill"},
		{"batch with no filter adds to all", gm, route.Compendium + route.CompendiumHabitatBatch, batch("", "Mountain", true), true, "Forest,Mountain", "Hill,Mountain"},
		{"player batch refused", player, route.Compendium + route.CompendiumHabitatBatch, batch("", "Urban", true), false, "Forest,Mountain", "Hill,Mountain"},
		{"batch removes from matches only", gm, route.Compendium + route.CompendiumHabitatBatch, batch("goblin", "Mountain", false), true, "Forest", "Hill,Mountain"},
		{"gm unticks", gm, route.Compendium + route.CompendiumHabitatCell, cell(goblin.ID, "Forest", false), true, "", "Hill,Mountain"},
		{"gm sets any", gm, route.Compendium + route.CompendiumHabitatAny, anyForm(goblin.ID), true, all, "Hill,Mountain"},
		{"player any refused", player, route.Compendium + route.CompendiumHabitatAny, anyForm(goblin.ID), false, all, "Hill,Mountain"},
		{"any on a spell refused", gm, route.Compendium + route.CompendiumHabitatAny, anyForm(spell.ID), false, all, "Hill,Mountain"},
		{"any from some sets the rest", gm, route.Compendium + route.CompendiumHabitatAny, anyForm(giant.ID), true, all, all},
		{"gm clears any", gm, route.Compendium + route.CompendiumHabitatAny, anyForm(goblin.ID), true, "", all},
		{"gm clears the giant's any", gm, route.Compendium + route.CompendiumHabitatAny, anyForm(giant.ID), true, "", ""},
		{"batch restores the giant", gm, route.Compendium + route.CompendiumHabitatBatch, batch("giant", "Hill", true), true, "", "Hill"},
		{"and its mountain", gm, route.Compendium + route.CompendiumHabitatBatch, batch("giant", "Mountain", true), true, "", "Hill,Mountain"},
	}
	for _, c := range cases {
		w := e.do(t, http.MethodPost, c.path, c.cookie, c.form, true)
		if ok := w.Code == http.StatusOK; ok != c.ok {
			t.Errorf("%s: status %d", c.name, w.Code)
		}
		if got := habitats(goblin.ID); got != c.goblin {
			t.Errorf("%s: goblin %q, want %q", c.name, got, c.goblin)
		}
		if got := habitats(giant.ID); got != c.giant {
			t.Errorf("%s: giant %q, want %q", c.name, got, c.giant)
		}
	}
	if got := habitats(spell.ID); got != "" {
		t.Errorf("spell got habitats %q", got)
	}

	// Every habitat reads as the Monster Manual's "Any".
	e.do(t, http.MethodPost, route.Compendium+route.CompendiumHabitatAny, gm, anyForm(goblin.ID), true)
	if body := e.do(t, http.MethodGet, route.Path(route.CampaignCompendium+route.CompendiumEntry, e.campaignID, goblin.ID), player, nil, false).Body.String(); !strings.Contains(body, "</strong> Any") {
		t.Error("entry page with every habitat does not read Any")
	}
	e.do(t, http.MethodPost, route.Compendium+route.CompendiumHabitatAny, gm, anyForm(goblin.ID), true)

	// Re-importing keeps what the GM typed.
	e.do(t, http.MethodPost, route.Compendium+route.CompendiumHabitatCell, gm, cell(goblin.ID, "Hill", true), false)
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	if got := habitats(goblin.ID); got != "Hill" {
		t.Errorf("after re-import: goblin %q", got)
	}

	page := route.Path(route.CampaignCompendium+route.CompendiumEntry, e.campaignID, giant.ID)
	if body := e.do(t, http.MethodGet, page, player, nil, false).Body.String(); !strings.Contains(body, "Hill, Mountain") || strings.Contains(body, route.Compendium+route.CompendiumHabitats) {
		t.Errorf("player page: want habitats, no grid link: %s", body)
	}
	if body := e.do(t, http.MethodGet, page, gm, nil, false).Body.String(); !strings.Contains(body, route.Compendium+route.CompendiumHabitats+"?q=Hill+Giant") {
		t.Errorf("gm page lacks the grid link")
	}
	if w := e.do(t, http.MethodGet, route.Compendium+route.CompendiumHabitats+"?q=giant", player, nil, false); w.Code == http.StatusOK {
		t.Error("player opened the grid")
	}
	if body := e.do(t, http.MethodGet, route.Compendium+route.CompendiumHabitats+"?q=giant", gm, nil, false).Body.String(); !strings.Contains(body, "Hill Giant") || strings.Contains(body, "Goblin Warrior") {
		t.Errorf("grid filter: %s", body)
	}
}

// TestHabitatCopies: a monster with no habitats takes those of its copy in
// another source, in bulk from the grid or one by one from its page; one with
// habitats keeps its own. A link import's habitats seed the new homebrew.
func TestHabitatCopies(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, enum.RoleAdmin)
	player := e.token(t, enum.RolePlayer)
	srd, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	// A lower-ranked copy, made first so that ordering by id alone would pick it.
	if err := e.q().UpsertCompendiumSource(ctx, sqlc.UpsertCompendiumSourceParams{SourceKey: "old-book", Edition: "2014"}); err != nil {
		t.Fatal(err)
	}
	if err := e.q().UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: "monster", Slug: "goblin-warrior", Name: "Goblin Warrior", SourceKey: "old-book", DataJson: "{}", ContentHash: "x"}); err != nil {
		t.Fatal(err)
	}
	old, err := e.q().ListEntryCopies(ctx, sqlc.ListEntryCopiesParams{Kind: "monster", Slug: "goblin-warrior"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range old {
		if c.SourceKey == "old-book" {
			if _, err := e.q().AddMonsterHabitat(ctx, sqlc.AddMonsterHabitatParams{EntryID: c.ID, Habitat: "Arctic"}); err != nil {
				t.Fatal(err)
			}
		}
	}
	brew, err := e.q().CreateHomebrew(ctx, sqlc.CreateHomebrewParams{Kind: "monster", Slug: "goblin-warrior", Name: "Goblin Warrior", DataJson: "{}", ContentHash: "x"})
	if err != nil {
		t.Fatal(err)
	}
	giant, err := e.q().CreateHomebrew(ctx, sqlc.CreateHomebrewParams{Kind: "monster", Slug: "hill-giant", Name: "Hill Giant", DataJson: "{}", ContentHash: "x"})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := e.q().AddMonsterHabitat(ctx, sqlc.AddMonsterHabitatParams{EntryID: brew.ID, Habitat: "Swamp"}); err != nil {
		t.Fatal(err)
	}
	habitats := func(id int64) string {
		h, _ := e.q().ListEntryHabitats(ctx, id)
		slices.Sort(h)
		return joinHabitats(h, ",")
	}
	fill := func(q string) url.Values { return url.Values{"q": {q}} }
	copyForm := func(to, from int64) url.Values {
		return url.Values{"entry_id": {fmt.Sprint(to)}, "from_id": {fmt.Sprint(from)}}
	}
	clearSwamp := func(id int64) {
		if err := e.q().RemoveMonsterHabitat(ctx, sqlc.RemoveMonsterHabitatParams{EntryID: id, Habitat: "Swamp"}); err != nil {
			t.Fatal(err)
		}
	}

	cases := []struct {
		name   string
		cookie string
		path   string
		form   url.Values
		ok     bool
		srd    string // habitats afterwards
		giant  string
	}{
		{"player fill refused", player, route.Compendium + route.CompendiumHabitatFill, fill(""), false, "", ""},
		{"fill outside the filter", gm, route.Compendium + route.CompendiumHabitatFill, fill("giant"), true, "", ""},
		{"fill takes the copy's", gm, route.Compendium + route.CompendiumHabitatFill, fill("goblin"), true, "Swamp", ""},
		{"fill again changes nothing", gm, route.Compendium + route.CompendiumHabitatFill, fill(""), true, "Swamp", ""},
		{"player copy refused", player, route.Compendium + route.CompendiumHabitatCopy, copyForm(giant.ID, brew.ID), false, "Swamp", ""},
		{"copy from any monster", gm, route.Compendium + route.CompendiumHabitatCopy, copyForm(giant.ID, brew.ID), true, "Swamp", "Swamp"},
		{"copy needs both ids", gm, route.Compendium + route.CompendiumHabitatCopy, url.Values{"entry_id": {fmt.Sprint(srd.ID)}}, false, "Swamp", "Swamp"},
	}
	for _, c := range cases {
		w := e.do(t, http.MethodPost, c.path, c.cookie, c.form, true)
		if ok := w.Code == http.StatusOK; ok != c.ok {
			t.Errorf("%s: status %d", c.name, w.Code)
		}
		if got := habitats(srd.ID); got != c.srd {
			t.Errorf("%s: srd %q, want %q", c.name, got, c.srd)
		}
		if got := habitats(giant.ID); got != c.giant {
			t.Errorf("%s: giant %q, want %q", c.name, got, c.giant)
		}
	}

	// A copy onto a monster that has habitats changes nothing.
	if _, err := e.q().AddMonsterHabitat(ctx, sqlc.AddMonsterHabitatParams{EntryID: brew.ID, Habitat: "Hill"}); err != nil {
		t.Fatal(err)
	}
	e.do(t, http.MethodPost, route.Compendium+route.CompendiumHabitatCopy, gm, copyForm(srd.ID, brew.ID), true)
	if got := habitats(srd.ID); got != "Swamp" {
		t.Errorf("copy over habitats: srd %q", got)
	}

	// The page offers the copy to the GM only, while the monster has none.
	clearSwamp(srd.ID)
	page := route.Path(route.CampaignCompendium+route.CompendiumEntry, e.campaignID, srd.ID)
	if body := e.do(t, http.MethodGet, page, gm, nil, false).Body.String(); !strings.Contains(body, route.Compendium+route.CompendiumHabitatCopy) || !strings.Contains(body, "Hill, Swamp") {
		t.Error("gm page lacks the copy button")
	}
	if body := e.do(t, http.MethodGet, page, player, nil, false).Body.String(); strings.Contains(body, route.Compendium+route.CompendiumHabitatCopy) {
		t.Error("player page has the copy button")
	}

	// A new homebrew monster takes the habitats its form carries.
	form := url.Values{"kind": {"monster"}, "name": {"Howler"}, "habitats": {"Planar", "Urban"}}
	if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumHomebrew, gm, form, false); w.Code != http.StatusSeeOther {
		t.Fatalf("homebrew create: status %d", w.Code)
	}
	howler, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "howler"})
	if err != nil {
		t.Fatal(err)
	}
	if got := habitats(howler.ID); got != "Planar,Urban" {
		t.Errorf("homebrew habitats %q", got)
	}
	form.Set("name", "Moonling")
	form["habitats"] = []string{"Moon"}
	if w := e.do(t, http.MethodPost, route.Compendium+route.CompendiumHomebrew, gm, form, false); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("unknown habitat: status %d", w.Code)
	}
}

func TestCompendiumHabitatFilter(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, enum.RoleAdmin)
	goblin, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	giant, err := e.q().CreateHomebrew(ctx, sqlc.CreateHomebrewParams{Kind: "monster", Slug: "hill-giant", Name: "Hill Giant", DataJson: "{}", BodyText: "giant", ContentHash: "x"})
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range []sqlc.AddMonsterHabitatParams{{EntryID: giant.ID, Habitat: "Hill"}, {EntryID: goblin.ID, Habitat: "Forest"}} {
		if _, err := e.q().AddMonsterHabitat(ctx, p); err != nil {
			t.Fatal(err)
		}
	}

	cases := []struct {
		name         string
		query        string
		giant        bool
		goblin       bool
		habitatBadge bool
	}{
		{"browse by habitat", "kind=monster&habitat=Hill", true, false, true},
		{"either habitat", "kind=monster&habitat=Hill,Forest", true, true, true},
		{"search with habitat", "kind=monster&habitat=Hill&q=giant", true, false, true},
		{"search outside habitat", "kind=monster&habitat=Hill&q=goblin", false, false, true},
		{"monsters unfiltered", "kind=monster", true, true, true},
		{"habitat ignored for other kinds", "kind=monster,spell&habitat=Hill", true, true, false},
	}
	for _, c := range cases {
		body := e.do(t, http.MethodGet, route.Compendium+"?"+c.query, gm, nil, false).Body.String()
		// Only the results, not the GM's homebrew tab after them.
		_, results, _ := strings.Cut(body, `id="`+views.CompendiumResultsID+`"`)
		results, _, _ = strings.Cut(results, `role="tabpanel"`)
		if got := strings.Contains(results, ">Hill Giant<"); got != c.giant {
			t.Errorf("%s: giant listed %v, want %v", c.name, got, c.giant)
		}
		if got := strings.Contains(results, ">Goblin Warrior<"); got != c.goblin {
			t.Errorf("%s: goblin listed %v, want %v", c.name, got, c.goblin)
		}
		if got := strings.Contains(body, `data-habitat="Hill"`); got != c.habitatBadge {
			t.Errorf("%s: habitat badges %v, want %v", c.name, got, c.habitatBadge)
		}
	}

	w := e.do(t, http.MethodGet, route.Compendium+route.CompendiumRandom+"?kind=monster&habitat=Hill", gm, nil, false)
	if want := route.Compendium + route.Path(route.CompendiumEntry, giant.ID); w.Header().Get("Location") != want {
		t.Errorf("random: %d %q, want %q", w.Code, w.Header().Get("Location"), want)
	}
}

// TestEntryBackKeepsFilters checks an entry's back link returns to the
// compendium list it was opened from, and never leaves the app.
func TestEntryBackKeepsFilters(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, enum.RoleAdmin)
	goblin, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ name, referer, back string }{
		{"filtered list", "https://192.168.1.10:8080" + route.Compendium + "?kind=monster&q=gob", `href="` + route.Compendium + `?kind=monster&amp;q=gob"`},
		{"other site", "https://evil.example" + route.Compendium + "?kind=spell", `href="` + route.Compendium + `?kind=spell"`},
		{"other page", "https://192.168.1.10:8080" + route.Path(route.CampaignNotes, 1), `href="` + route.Compendium + `"`},
		{"no referer", "", `href="` + route.Compendium + `"`},
	}
	for _, c := range cases {
		req := httptest.NewRequestWithContext(ctx, http.MethodGet, route.Compendium+route.Path(route.CompendiumEntry, goblin.ID), nil)
		req.RemoteAddr = lanAddr
		req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: gm})
		if c.referer != "" {
			req.Header.Set("Referer", c.referer)
		}
		w := httptest.NewRecorder()
		e.router.ServeHTTP(w, req)
		if body := w.Body.String(); !strings.Contains(body, c.back) {
			t.Errorf("%s: back link %s missing", c.name, c.back)
		}
	}
}

// TestHabitatsMatchSchema keeps the habitat forms' oneof lists in step with
// compendium.Habitats.
func TestHabitatsMatchSchema(t *testing.T) {
	want := "oneof=" + joinHabitats(compendium.Habitats, " ")
	for _, form := range []any{habitatCellForm{}, habitatBatchForm{}} {
		f, _ := reflect.TypeOf(form).FieldByName("Habitat")
		if !strings.Contains(f.Tag.Get("validate"), want) {
			t.Errorf("%T: validate %q, want %s", form, f.Tag.Get("validate"), want)
		}
	}
	if f, _ := reflect.TypeFor[compendium.HomebrewInput]().FieldByName("Habitats"); !strings.Contains(f.Tag.Get("validate"), want) {
		t.Errorf("HomebrewInput: validate %q, want %s", f.Tag.Get("validate"), want)
	}
}
