package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/web/views"
)

// TestPagesRender requests every page as the GM and as a player and
// checks each one renders.
func TestPagesRender(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)

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
		w := e.do(t, http.MethodPost, fmt.Sprintf("/encounters/%d/combatants", enc.ID), gm, url.Values{"kind": {kind}, "id": {fmt.Sprint(id)}}, true)
		if w.Code != http.StatusOK {
			t.Fatalf("add %s: %d %s", kind, w.Code, w.Body)
		}
	}
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/encounters/%d/start", enc.ID), gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("start: %d %s", w.Code, w.Body)
	}
	next := fmt.Sprintf("/encounters/%d/next", enc.ID)
	if w := e.do(t, http.MethodPost, next, gm, nil, true); w.Code == http.StatusOK {
		t.Fatal("next turn ran without initiatives")
	}
	rows, _ := e.q().ListCombatants(ctx, enc.ID)
	for i, r := range rows {
		if w := e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/initiative", r.ID), gm, url.Values{"initiative": {fmt.Sprint(10 + i)}}, true); w.Code != http.StatusOK {
			t.Fatalf("initiative: %d %s", w.Code, w.Body)
		}
		if r.EntityID.Valid && !r.MonsterRef.Valid { // Mira has no stats: max HP starts her at full
			e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/hp", r.ID), gm, url.Values{"op": {"set"}, "max_hp": {"12"}}, true)
			e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/hp", r.ID), gm, url.Values{"op": {"damage"}, "amount": {"15"}}, true)
			if cb, _ := e.q().GetCombatant(ctx, r.ID); cb.MaxHp != 12 || cb.Hp != -3 {
				t.Fatalf("Mira HP %d/%d, want -3/12", cb.Hp, cb.MaxHp)
			}
			// Set fills only the given fields; HP is capped at max.
			e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/hp", r.ID), gm, url.Values{"op": {"set"}, "hp": {"20"}, "temp_hp": {"4"}}, true)
			if cb, _ := e.q().GetCombatant(ctx, r.ID); cb.Hp != 12 || cb.MaxHp != 12 || cb.TempHp != 4 {
				t.Fatalf("Mira HP %d/%d+%d, want 12/12+4", cb.Hp, cb.MaxHp, cb.TempHp)
			}
			if w := e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/hp", r.ID), gm, url.Values{"op": {"set"}}, true); w.Code == http.StatusOK {
				t.Fatal("empty set succeeded")
			}
			e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/hp", r.ID), gm, url.Values{"op": {"damage"}, "amount": {"19"}}, true)
		}
	}
	if w := e.do(t, http.MethodPost, next, gm, nil, true); w.Code != http.StatusOK {
		t.Fatalf("next: %d %s", w.Code, w.Body)
	}

	pages := []struct {
		path      string
		gm, other bool // who may load it
		contains  string
	}{
		{"/", true, false, "Campaigns"},
		{fmt.Sprintf("/campaigns/%d", e.campaignID), true, false, "Alice"},
		{fmt.Sprintf("/characters/%d", e.characterID), true, true, "Alice"},
		{fmt.Sprintf("/characters/%d/levelup?class=wizard", e.characterID), false, true, "Wizard 1"},
		{fmt.Sprintf("/campaigns/%d/combat", e.campaignID), true, true, "Goblin Warrior"},
		{fmt.Sprintf("/campaigns/%d/notes?q=Session", e.campaignID), true, true, "Session 1"},
		{fmt.Sprintf("/notes/%d", note.ID), true, true, "Session 1"},
		{fmt.Sprintf("/campaigns/%d/entities", e.campaignID), true, true, "Mira"},
		{fmt.Sprintf("/entities/%d", ent.ID), true, true, "Mira"},
		{fmt.Sprintf("/campaigns/%d/switcher?q=mi", e.campaignID), true, true, "Mira"},
		{"/compendium?q=fire", true, true, "Fire Bolt"},
		{"/compendium", true, true, "Goblin Warrior"}, // first batch of every kind
		{fmt.Sprintf("/compendium/entries/%d", goblin.ID), true, true, "Scimitar"},
		{fmt.Sprintf("/compendium/entries/%d?preview=1", goblin.ID), true, true, "Scimitar"},
		{"/compendium?kind=spell", true, true, "Fire Bolt"},
		{"/compendium?kind=spell", true, true, ">sorcerer</span>"},                        // class badges on spell rows
		{"/compendium", true, true, `href="/compendium?kind=spell"`},                      // kind counts link to their list
		{"/compendium?kind=spell&q=fire", true, true, `href="/compendium?q=fire"`},        // the active kind links back to all
		{"/compendium?kind=spell", true, true, `href="/compendium?kind=monster%2Cspell"`}, // inactive kinds join the filter
		{"/compendium?kind=monster,spell", true, true, "Goblin Warrior"},
		{"/compendium?kind=monster,spell", true, true, "Fire Bolt"},
		{"/compendium?after_name=Fire+Bolt&after_id=0", true, true, "Fire Bolt"},           // cursor is (name, id): same name, later id
		{"/compendium?after_name=Fire+Bolt&after_id=999999", true, true, "Goblin Warrior"}, // next batch after the cursor
		{"/compendium/search?q=gob", true, true, "Goblin Warrior"},
		{"/compendium/search?kind=monster&q=gob&encounter=1", true, true, "Goblin Warrior"},
		{"/compendium/homebrew/new", true, true, "New homebrew"}, // players write their own spells and items
	}
	if w := e.do(t, http.MethodGet, "/compendium/random?kind=monster", player, nil, false); w.Code != http.StatusSeeOther || !strings.HasPrefix(w.Header().Get("Location"), "/compendium/entries/") {
		t.Errorf("random: %d %q", w.Code, w.Header().Get("Location"))
	}
	if w := e.do(t, http.MethodGet, "/compendium/random?kind=monster", player, nil, true); !strings.HasPrefix(w.Header().Get("HX-Redirect"), "/compendium/entries/") {
		t.Errorf("htmx random: %d %q", w.Code, w.Header().Get("HX-Redirect"))
	}
	req := httptest.NewRequestWithContext(ctx, http.MethodGet, "/compendium?kind=spell", nil)
	req.Header.Set("HX-Request", "true")
	req.Header.Set("HX-Target", views.CompendiumResultsID)
	req.RemoteAddr = lanAddr
	req.AddCookie(&http.Cookie{Name: auth.CookieName, Value: player})
	w := httptest.NewRecorder()
	e.router.ServeHTTP(w, req)
	if body := w.Body.String(); w.Code != http.StatusOK || !strings.Contains(body, "Fire Bolt") || strings.Contains(body, "<html") ||
		!strings.Contains(body, `href="/compendium?kind=monster%2Cspell"`) { // badges re-rendered without a stale query
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
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)
	goblin, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	spell, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "spell", Slug: "fire-bolt"})
	page := fmt.Sprintf("/compendium/entries/%d", goblin.ID)
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
		w := e.do(t, http.MethodPost, fmt.Sprintf("/compendium/entries/%d/image", c.id), c.cookie, url.Values{"image_url": {c.link}}, false)
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

func TestMonsterHabitats(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, auth.RoleAdmin)
	player := e.token(t, auth.RolePlayer)
	goblin, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	spell, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "spell", Slug: "fire-bolt"})
	giant, err := e.q().CreateHomebrew(ctx, sqlc.CreateHomebrewParams{Kind: "monster", Slug: "hill-giant", Name: "Hill Giant", DataJson: "{}", ContentHash: "x"})
	if err != nil {
		t.Fatal(err)
	}
	habitats := func(id int64) string {
		h, _ := e.q().ListEntryHabitats(ctx, id)
		slices.Sort(h)
		return strings.Join(h, ",")
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
	all := strings.Join(every, ",")

	// Every listed habitat passes the form's oneof and the table's CHECK.
	for _, h := range compendium.Habitats {
		if w := e.do(t, http.MethodPost, "/compendium/habitats/cell", gm, cell(goblin.ID, h, true), false); w.Code != http.StatusOK {
			t.Fatalf("habitat %s: status %d", h, w.Code)
		}
	}
	if got := habitats(goblin.ID); strings.Count(got, ",") != len(compendium.Habitats)-1 {
		t.Fatalf("goblin habitats %q", got)
	}
	for _, h := range compendium.Habitats {
		e.do(t, http.MethodPost, "/compendium/habitats/cell", gm, cell(goblin.ID, h, false), false)
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
		{"gm ticks", gm, "/compendium/habitats/cell", cell(goblin.ID, "Forest", true), true, "Forest", ""},
		{"gm ticks again", gm, "/compendium/habitats/cell", cell(goblin.ID, "Forest", true), true, "Forest", ""},
		{"player refused", player, "/compendium/habitats/cell", cell(goblin.ID, "Hill", true), false, "Forest", ""},
		{"unknown habitat", gm, "/compendium/habitats/cell", cell(goblin.ID, "Moon", true), false, "Forest", ""},
		{"not a monster", gm, "/compendium/habitats/cell", cell(spell.ID, "Hill", true), false, "Forest", ""},
		{"batch adds to matches only", gm, "/compendium/habitats/batch", batch("HILL", "Hill", true), true, "Forest", "Hill"},
		{"batch with no filter adds to all", gm, "/compendium/habitats/batch", batch("", "Mountain", true), true, "Forest,Mountain", "Hill,Mountain"},
		{"player batch refused", player, "/compendium/habitats/batch", batch("", "Urban", true), false, "Forest,Mountain", "Hill,Mountain"},
		{"batch removes from matches only", gm, "/compendium/habitats/batch", batch("goblin", "Mountain", false), true, "Forest", "Hill,Mountain"},
		{"gm unticks", gm, "/compendium/habitats/cell", cell(goblin.ID, "Forest", false), true, "", "Hill,Mountain"},
		{"gm sets any", gm, "/compendium/habitats/any", anyForm(goblin.ID), true, all, "Hill,Mountain"},
		{"player any refused", player, "/compendium/habitats/any", anyForm(goblin.ID), false, all, "Hill,Mountain"},
		{"any on a spell refused", gm, "/compendium/habitats/any", anyForm(spell.ID), false, all, "Hill,Mountain"},
		{"any from some sets the rest", gm, "/compendium/habitats/any", anyForm(giant.ID), true, all, all},
		{"gm clears any", gm, "/compendium/habitats/any", anyForm(goblin.ID), true, "", all},
		{"gm clears the giant's any", gm, "/compendium/habitats/any", anyForm(giant.ID), true, "", ""},
		{"batch restores the giant", gm, "/compendium/habitats/batch", batch("giant", "Hill", true), true, "", "Hill"},
		{"and its mountain", gm, "/compendium/habitats/batch", batch("giant", "Mountain", true), true, "", "Hill,Mountain"},
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
	e.do(t, http.MethodPost, "/compendium/habitats/any", gm, anyForm(goblin.ID), true)
	if body := e.do(t, http.MethodGet, fmt.Sprintf("/compendium/entries/%d", goblin.ID), player, nil, false).Body.String(); !strings.Contains(body, "</strong> Any") {
		t.Error("entry page with every habitat does not read Any")
	}
	e.do(t, http.MethodPost, "/compendium/habitats/any", gm, anyForm(goblin.ID), true)

	// Re-importing keeps what the GM typed.
	e.do(t, http.MethodPost, "/compendium/habitats/cell", gm, cell(goblin.ID, "Hill", true), false)
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	if got := habitats(goblin.ID); got != "Hill" {
		t.Errorf("after re-import: goblin %q", got)
	}

	page := fmt.Sprintf("/compendium/entries/%d", giant.ID)
	if body := e.do(t, http.MethodGet, page, player, nil, false).Body.String(); !strings.Contains(body, "Hill, Mountain") || strings.Contains(body, "/compendium/habitats") {
		t.Errorf("player page: want habitats, no grid link: %s", body)
	}
	if body := e.do(t, http.MethodGet, page, gm, nil, false).Body.String(); !strings.Contains(body, "/compendium/habitats?q=Hill+Giant") {
		t.Errorf("gm page lacks the grid link")
	}
	if w := e.do(t, http.MethodGet, "/compendium/habitats?q=giant", player, nil, false); w.Code == http.StatusOK {
		t.Error("player opened the grid")
	}
	if body := e.do(t, http.MethodGet, "/compendium/habitats?q=giant", gm, nil, false).Body.String(); !strings.Contains(body, "Hill Giant") || strings.Contains(body, "Goblin Warrior") {
		t.Errorf("grid filter: %s", body)
	}
}

func TestCompendiumHabitatFilter(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, auth.RoleAdmin)
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
		body := e.do(t, http.MethodGet, "/compendium?"+c.query, gm, nil, false).Body.String()
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

	w := e.do(t, http.MethodGet, "/compendium/random?kind=monster&habitat=Hill", gm, nil, false)
	if want := fmt.Sprintf("/compendium/entries/%d", giant.ID); w.Header().Get("Location") != want {
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
	gm := e.token(t, auth.RoleAdmin)
	goblin, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ name, referer, back string }{
		{"filtered list", "https://192.168.1.10:8080/compendium?kind=monster&q=gob", `href="/compendium?kind=monster&amp;q=gob"`},
		{"other site", "https://evil.example/compendium?kind=spell", `href="/compendium?kind=spell"`},
		{"other page", "https://192.168.1.10:8080/campaigns/1/notes", `href="/compendium"`},
		{"no referer", "", `href="/compendium"`},
	}
	for _, c := range cases {
		req := httptest.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("/compendium/entries/%d", goblin.ID), nil)
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
	want := "oneof=" + strings.Join(compendium.Habitats, " ")
	for _, form := range []any{habitatCellForm{}, habitatBatchForm{}} {
		f, _ := reflect.TypeOf(form).FieldByName("Habitat")
		if !strings.Contains(f.Tag.Get("validate"), want) {
			t.Errorf("%T: validate %q, want %s", form, f.Tag.Get("validate"), want)
		}
	}
}
