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
	one := []dnd5e.ClassLevel{{Slug: "x", Level: 1}}
	two := append(slices.Clone(one), dnd5e.ClassLevel{Slug: "y", Level: 1})
	spells := []dnd5e.Spell{
		{Slug: "c1", Class: "x"}, {Slug: "c2", Ability: "wis"}, // class, lineage
		{Slug: "s1", Level: 1, Feat: "Magic Initiate (Wizard)"}, {Slug: "s2", Level: 1}, // feat, untagged
	}
	for _, tt := range []struct {
		name  string
		cols  map[string]string
		sheet dnd5e.Sheet
		class string
		want  views.SpellNeed
	}{
		{"no columns", map[string]string{"Rages": "2"}, dnd5e.Sheet{}, "x", views.SpellNeed{Cantrips: -1, Spells: -1}},
		{"prepared, at most the offered", map[string]string{"Cantrips": "3", "Prepared Spells": "4"}, dnd5e.Sheet{}, "x", views.SpellNeed{Cantrips: 2, Spells: 3, SpellsUpTo: true}},
		{"known, less the class's", known, dnd5e.Sheet{Classes: one, Spells: spells}, "x", views.SpellNeed{Cantrips: 1, Spells: 2}},
		// A multiclass's untagged spell may be either class's: upper limits only.
		{"untagged spells of a multiclass", known, dnd5e.Sheet{Classes: two, Spells: spells}, "x", views.SpellNeed{Cantrips: 1, Spells: 3, CantripsUpTo: true, SpellsUpTo: true}},
		{"another class's spells", known, dnd5e.Sheet{Classes: two, Spells: spells}, "y", views.SpellNeed{Cantrips: 2, Spells: 3, CantripsUpTo: true, SpellsUpTo: true}},
		{"a new class", known, dnd5e.Sheet{Classes: one, Spells: spells}, "z", views.SpellNeed{Cantrips: 2, Spells: 3}},
		{"more than the total", map[string]string{"Cantrips": "1"}, dnd5e.Sheet{Classes: one, Spells: []dnd5e.Spell{{Slug: "c1", Class: "x"}, {Slug: "c3", Class: "x"}}}, "x", views.SpellNeed{Cantrips: 0, Spells: -1}},
	} {
		if got := spellNeed(tt.cols, tt.sheet, tt.class, offered); got != tt.want {
			t.Errorf("%s: %+v, want %+v", tt.name, got, tt.want)
		}
	}
}

// TestSpellCounts: creation and level-up refuse cantrips the class table
// does not leave, too few cantrips, more prepared spells than it allows and
// spells off the list, keeping the picks; fewer prepared spells are fine.
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
	for _, step := range []struct {
		url   string
		form  url.Values
		spell []string
		code  int
		err   string
	}{
		// Level 1: no cantrips, up to one spell.
		{create + "choices", choices, []string{"test-blast"}, http.StatusUnprocessableEntity, "choose up to 1 level 1+ spell: you ticked 1 cantrip and 0 level 1+ spells"},
		{create + "choices", choices, []string{"test-ward", "test-hex"}, http.StatusUnprocessableEntity, "you ticked 0 cantrips and 2 level 1+ spells"},
		{create + "choices", choices, []string{"test-ward", "test-ward"}, http.StatusUnprocessableEntity, "choose spells from the list"},
		{create + "choices", choices, []string{"test-ward"}, http.StatusSeeOther, ""},
		// Level 2: one cantrip, up to one more spell.
		{levelUp, level2, nil, http.StatusUnprocessableEntity, "choose 1 cantrip and up to 1 level 1+ spell: you ticked 0"},
		{levelUp, level2, []string{"test-blast", "test-ward"}, http.StatusUnprocessableEntity, "choose spells from the list"}, // Test Ward is known
		{levelUp, level2, []string{"test-blast"}, http.StatusSeeOther, ""},
	} {
		form := maps.Clone(step.form)
		form["spell"] = step.spell
		w := e.do(t, http.MethodPost, step.url, player, form, false)
		if w.Code != step.code || !strings.Contains(w.Body.String(), step.err) {
			t.Fatalf("%v: %d, want %d with %q", form, w.Code, step.code, step.err)
		}
		if lost := lostPicks(w.Body.String(), form); step.code == http.StatusUnprocessableEntity && len(lost) > 0 {
			t.Errorf("%v: refused page lost %v", form, lost)
		}
	}
	ch, _ := e.q().GetCharacter(t.Context(), e.characterID)
	sh, _ := dnd5e.ParseSheet(ch.SheetJson)
	if got := slices.DeleteFunc(slices.Clone(sh.Spells), func(sp dnd5e.Spell) bool { return sp.Class != "testlock" }); len(got) != 2 {
		t.Errorf("class spells %+v, want Test Ward and Test Blast", got)
	}
}
