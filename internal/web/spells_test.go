package web

import (
	"fmt"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/web/views"
)

// TestSpellNeed: the class table's totals less the class's spells on the
// sheet, at most what is offered; lineage and feat spells do not count.
func TestSpellNeed(t *testing.T) {
	offered := []views.Option{{Slug: "a"}, {Slug: "b"}, {Slug: "c", Level: 1}, {Slug: "d", Level: 1}, {Slug: "e", Level: 2}}
	known := map[string]string{"Cantrips Known": "2", "Spells Known": "3"}
	prepared := map[string]string{"Cantrips": "3", "Prepared Spells": "1"}
	one := []dnd5e.ClassLevel{{Slug: "x", Level: 1}}
	two := append(slices.Clone(one), dnd5e.ClassLevel{Slug: "y", Level: 1})
	spells := []dnd5e.Spell{
		{Slug: "c1", Class: "x"}, {Slug: "c2", Ability: "wis"}, // class, lineage
		{Slug: "s1", Level: 1, Feat: "Magic Initiate (Wizard)"}, {Slug: "s2", Level: 1}, // feat, untagged
	}
	for _, tt := range []struct {
		name  string
		cols  map[string]string
		book  int
		sheet dnd5e.Sheet
		class string
		want  views.SpellNeed
	}{
		{"no columns", map[string]string{"Rages": "2"}, 0, dnd5e.Sheet{}, "x", views.SpellNeed{Cantrips: -1, Spells: -1}},
		{"prepared, at most the offered", map[string]string{"Cantrips": "3", "Prepared Spells": "4"}, 0, dnd5e.Sheet{}, "x", views.SpellNeed{Cantrips: 2, Spells: 3, SpellsUpTo: true}},
		{"known, less the class's", known, 0, dnd5e.Sheet{Classes: one, Spells: spells}, "x", views.SpellNeed{Cantrips: 1, Spells: 2}},
		// A multiclass's untagged spell may be either class's: upper limits only.
		{"untagged spells of a multiclass", known, 0, dnd5e.Sheet{Classes: two, Spells: spells}, "x", views.SpellNeed{Cantrips: 1, Spells: 3, CantripsUpTo: true, SpellsUpTo: true}},
		{"another class's spells", known, 0, dnd5e.Sheet{Classes: two, Spells: spells}, "y", views.SpellNeed{Cantrips: 2, Spells: 3, CantripsUpTo: true, SpellsUpTo: true}},
		{"a new class", known, 0, dnd5e.Sheet{Classes: one, Spells: spells}, "z", views.SpellNeed{Cantrips: 2, Spells: 3}},
		{"more than the total", map[string]string{"Cantrips": "1"}, 0, dnd5e.Sheet{Classes: one, Spells: []dnd5e.Spell{{Slug: "c1", Class: "x"}, {Slug: "c3", Class: "x"}}}, "x", views.SpellNeed{Cantrips: 0, Spells: -1}},
		// A spellbook replaces the prepared total, exactly.
		{"spellbook", prepared, 3, dnd5e.Sheet{Classes: one, Spells: spells}, "x", views.SpellNeed{Cantrips: 2, Spells: 2, Book: true, BookHave: 1, BookSize: 3}},
		{"spellbook of a multiclass", prepared, 3, dnd5e.Sheet{Classes: two, Spells: spells}, "x", views.SpellNeed{Cantrips: 2, Spells: 3, CantripsUpTo: true, SpellsUpTo: true, Book: true, BookSize: 3}},
	} {
		info := dnd5e.ClassInfo{Slug: tt.class, Table: map[string]map[int]string{}}
		for col, v := range tt.cols {
			info.Table[col] = map[int]string{1: v}
		}
		if tt.book > 0 {
			info.Features = []dnd5e.ClassFeature{{Description: fmt.Sprintf("It starts with %d level 1 X spells.", tt.book)}}
		}
		if got := spellNeed(info, 1, tt.sheet, offered); got != tt.want {
			t.Errorf("%s: %+v, want %+v", tt.name, got, tt.want)
		}
	}
	// A book already short at the previous level gets the hint, a new class's none.
	info := dnd5e.ClassInfo{Slug: "x", Features: []dnd5e.ClassFeature{{Description: "It starts with 3 level 1 X spells."}}}
	if n := spellNeed(info, 2, dnd5e.Sheet{Classes: one, Spells: spells}, offered); !n.BookShort || n.BookHave != 1 || n.BookSize != 3 {
		t.Errorf("short book at level 2: %+v", n)
	}
	if n := spellNeed(info, 1, dnd5e.Sheet{Classes: one, Spells: spells}, offered); n.BookShort {
		t.Errorf("short book at level 1: %+v", n)
	}
}

// TestSpellbook reads a spellbook's size from the class's features.
func TestSpellbook(t *testing.T) {
	wizard := dnd5e.ClassInfo{Features: []dnd5e.ClassFeature{{Name: "Spellcasting", Description: "The book contains the level 1+ spells you know. It starts with six level 1 Wizard spells of your choice.\n\nWhenever you gain a Wizard level after 1, add two Wizard spells of your choice to your spellbook."}}}
	lock := dnd5e.ClassInfo{Features: []dnd5e.ClassFeature{{Name: "Pact Magic", Description: "You know two Warlock cantrips of your choice."}}}
	for _, tt := range []struct {
		name  string
		info  dnd5e.ClassInfo
		level int
		want  int
	}{
		{"wizard 1", wizard, 1, 6},
		{"wizard 2", wizard, 2, 8},
		{"wizard 5", wizard, 5, 14},
		{"no spellbook", lock, 5, 0},
	} {
		if got := spellbook(tt.info, tt.level); got != tt.want {
			t.Errorf("%s: %d, want %d", tt.name, got, tt.want)
		}
	}
}

// TestSpellCounts: creation and level-up refuse cantrips the class table
// does not leave, too few cantrips, more prepared spells than it allows and
// spells off the list, keeping the picks; fewer prepared spells are fine,
// fewer than a spellbook holds are not.
func TestSpellCounts(t *testing.T) {
	lock := testlock
	lock.DataJson = strings.Replace(testlock.DataJson, `"table":{`, `"table":{"Prepared Spells":{"1":"1","2":"2"},`, 1)
	ward := sqlc.UpsertCompendiumEntryParams{Kind: "spell", Slug: "test-ward", Name: "Test Ward", SourceKey: "test", ContentHash: "x", DataJson: `{"level":1,"classes":["testlock"]}`}
	hex := ward
	hex.Slug, hex.Name = "test-hex", "Test Hex"
	e := newTestEnv(t)
	loadFeatFixtures(t, e)
	for _, entry := range []sqlc.UpsertCompendiumEntryParams{lock, testBlast, ward, hex} {
		if err := e.q().UpsertCompendiumEntry(t.Context(), entry); err != nil {
			t.Fatal(err)
		}
	}
	player := e.token(t, auth.RolePlayer)
	create := fmt.Sprintf("/characters/%d/create/", e.characterID)
	for _, step := range []string{"class=testlock", "species=human"} {
		step, p, _ := strings.Cut(step, "=")
		if w := e.do(t, http.MethodPost, create+step, player, url.Values{"pick": {p}}, false); w.Code != http.StatusSeeOther {
			t.Fatalf("%s: %d %s", step, w.Code, w.Body)
		}
	}
	const pick = "class_eldritch-invocation"
	choices := url.Values{"language": {"Elvish", "Dwarvish"}, "trait_versatile": {"Alert"}, pick: {"Armor of Shadows"}}
	levelUp, level2 := fmt.Sprintf("/characters/%d/levelup", e.characterID), url.Values{"class": {"testlock"}, pick: {"Agonizing Blast", "Shadow Step"}}
	// From level 3 the class keeps a spellbook: one spell, plus one per level.
	book := lock
	book.ContentHash = "book" // the upsert skips an unchanged hash
	book.DataJson = strings.Replace(lock.DataJson, `"features":[`, `"features":[{"name":"Spellbook","desc":"It starts with one level 1 Testlock spells. Whenever you gain a Testlock level after 1, add one Testlock spells to it.","levels":[1]},`, 1)
	for _, step := range []struct {
		class *sqlc.UpsertCompendiumEntryParams // replaces the class first
		url   string
		form  url.Values
		spell []string
		code  int
		err   string
	}{
		// Level 1: no cantrips, up to one spell.
		{nil, create + "choices", choices, []string{"test-blast"}, http.StatusUnprocessableEntity, "choose up to 1 level 1+ spell: you ticked 1 cantrip and 0 level 1+ spells"},
		{nil, create + "choices", choices, []string{"test-ward", "test-hex"}, http.StatusUnprocessableEntity, "you ticked 0 cantrips and 2 level 1+ spells"},
		{nil, create + "choices", choices, []string{"test-ward", "test-ward"}, http.StatusUnprocessableEntity, "choose spells from the list"},
		{nil, create + "choices", choices, []string{"test-ward"}, http.StatusSeeOther, ""},
		// Level 2: one cantrip, up to one more spell.
		{nil, levelUp, level2, nil, http.StatusUnprocessableEntity, "choose 1 cantrip and up to 1 level 1+ spell: you ticked 0"},
		{nil, levelUp, level2, []string{"test-blast", "test-ward"}, http.StatusUnprocessableEntity, "choose spells from the list"}, // Test Ward is known
		{nil, levelUp, level2, []string{"test-blast"}, http.StatusSeeOther, ""},
		// Level 3: a book of three holding Test Ward, so exactly one more (only Test Hex is left).
		{&book, levelUp, url.Values{"class": {"testlock"}}, nil, http.StatusUnprocessableEntity, "choose 1 level 1+ spell for your spellbook (your spellbook has 1 of 3): you ticked 0"},
		{nil, levelUp, url.Values{"class": {"testlock"}}, []string{"test-hex"}, http.StatusSeeOther, ""},
	} {
		if step.class != nil {
			if err := e.q().UpsertCompendiumEntry(t.Context(), *step.class); err != nil {
				t.Fatal(err)
			}
		}
		form := maps.Clone(step.form)
		form["spell"] = step.spell
		w := e.do(t, http.MethodPost, step.url, player, form, false)
		if w.Code != step.code || !strings.Contains(w.Body.String(), step.err) {
			t.Fatalf("%v: %d, want %d with %q", form, w.Code, step.code, step.err)
		}
		if lost := lostPicks(w.Body.String(), form); step.code == http.StatusUnprocessableEntity && (len(lost) > 0 || !errorAt(w.Body.String(), "spell")) {
			t.Errorf("%v: refused page lost %v or its error is not at the spells", form, lost)
		}
	}
	ch, _ := e.q().GetCharacter(t.Context(), e.characterID)
	sh, _ := dnd5e.ParseSheet(ch.SheetJson)
	if got := slices.DeleteFunc(slices.Clone(sh.Spells), func(sp dnd5e.Spell) bool { return sp.Class != "testlock" }); len(got) != 3 {
		t.Errorf("class spells %+v, want Test Ward, Test Blast and Test Hex", got)
	}
}

// TestPrepareCapCommand: the class table's Prepared Spells reaches the
// prepare command, which refuses one more.
func TestPrepareCapCommand(t *testing.T) {
	lock := testlock
	lock.DataJson = strings.Replace(testlock.DataJson, `"table":{`, `"table":{"Prepared Spells":{"1":"1"},`, 1)
	e := newTestEnv(t)
	loadFeatFixtures(t, e)
	if err := e.q().UpsertCompendiumEntry(t.Context(), lock); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/create/class", e.characterID), player, url.Values{"pick": {"testlock"}}, false); w.Code != http.StatusSeeOther {
		t.Fatalf("class: %d %s", w.Code, w.Body)
	}
	ch, _ := e.q().GetCharacter(t.Context(), e.characterID)
	sh, _ := dnd5e.ParseSheet(ch.SheetJson)
	sh.Spells = []dnd5e.Spell{{Slug: "test-ward", Name: "Test Ward", Level: 1, Class: "testlock"}, {Slug: "test-hex", Name: "Test Hex", Level: 1, Class: "testlock"}}
	raw, _ := sh.JSON()
	if err := e.q().UpdateCharacterSheet(t.Context(), sqlc.UpdateCharacterSheetParams{ID: e.characterID, SheetJson: raw}); err != nil {
		t.Fatal(err)
	}
	cmdPath := fmt.Sprintf("/characters/%d/cmd", e.characterID)
	if w := e.do(t, http.MethodPost, cmdPath, player, url.Values{"cmd": {"spell prepare test-ward"}}, true); w.Code != http.StatusOK {
		t.Fatalf("first prepare: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodPost, cmdPath, player, url.Values{"cmd": {"spell prepare test-hex"}}, true); w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "prepares 1 spell") {
		t.Errorf("past the cap: %d %s", w.Code, w.Body)
	}
}
