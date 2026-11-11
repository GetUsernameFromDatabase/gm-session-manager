package compendium

import (
	"encoding/json"
	"reflect"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
)

// fiveToolsClasses are the fixture's classes and subclasses by kind/slug.
func fiveToolsClasses(t *testing.T) map[string]Entry {
	t.Helper()
	_, entries, err := parseFiveToolsBook("5etools-xphb", readBook(t, fiveToolsFixture))
	if err != nil {
		t.Fatal(err)
	}
	out := map[string]Entry{}
	for _, e := range entries {
		if e.Kind == enum.KindClass || e.Kind == enum.KindSubclass {
			out[string(e.Kind)+"/"+e.Slug] = e
		}
	}
	return out
}

func feature(d ClassData, name string) (ClassFeature, bool) {
	i := slices.IndexFunc(d.Features, func(f ClassFeature) bool { return f.Name == name })
	if i < 0 {
		return ClassFeature{}, false
	}
	return d.Features[i], true
}

func TestFiveToolsClasses(t *testing.T) {
	got := fiveToolsClasses(t)
	if len(got) != 4 {
		t.Fatalf("%d classes and subclasses, want 4 (the PHB fighter is another book's)", len(got))
	}
	for _, tt := range []struct {
		key        string
		hitDie     int
		caster     string
		saves      []string
		subclassOf string
		levels     map[string][]int
		subclass   []string
	}{
		{"class/fighter", 10, "none", []string{"str", "con"}, "", map[string][]int{
			"Second Wind": {1}, "Action Surge": {2}, "Fighter Subclass": {3},
			"Ability Score Improvement": {4, 6, 8}, "Subclass Feature": {7},
		}, []string{"Fighter Subclass", "Subclass Feature"}},
		{"class/wizard", 6, "full", []string{"int", "wis"}, "", map[string][]int{
			"Arcane Recovery": {1}, "Wizard Subclass": {3}, "Ability Score Improvement": {4},
		}, []string{"Wizard Subclass"}},
		{"subclass/champion", 0, "none", nil, "fighter", map[string][]int{"Champion": {3}, "Additional Fighting Style": {7}}, nil},
		{"subclass/evoker", 0, "none", nil, "wizard", map[string][]int{"Evoker": {3}, "Sculpt Spells": {6}}, nil},
	} {
		e, ok := got[tt.key]
		if !ok {
			t.Errorf("no %s", tt.key)
			continue
		}
		d := e.Data.(ClassData)
		if d.HitDie != tt.hitDie || d.Caster != tt.caster || !reflect.DeepEqual(d.Saves, tt.saves) || d.SubclassOf != tt.subclassOf {
			t.Errorf("%s: %+v", tt.key, d)
		}
		for name, levels := range tt.levels {
			if f, ok := feature(d, name); !ok || !reflect.DeepEqual(f.Levels, levels) {
				t.Errorf("%s %s: %+v", tt.key, name, f)
			}
		}
		for _, f := range d.Features {
			if f.Name == "" {
				t.Errorf("%s: nameless feature %+v", tt.key, f)
			}
			if f.Subclass != slices.Contains(tt.subclass, f.Name) {
				t.Errorf("%s %s: subclass %v", tt.key, f.Name, f.Subclass)
			}
		}
	}

	fighter, wizard := got["class/fighter"].Data.(ClassData), got["class/wizard"].Data.(ClassData)
	for _, tt := range []struct {
		name string
		d    ClassData
		col  string
		lvl  int
		want string
	}{
		{"fighter", fighter, "Second Wind", 4, "3"},
		{"fighter", fighter, "Weapon Mastery", 1, "3"},
		{"wizard", wizard, "Cantrips", 4, "4"},
		{"wizard", wizard, "Prepared Spells", 5, "9"},
		{"wizard", wizard, "Slots 1st", 1, "2"},
		{"wizard", wizard, "Slots 2nd", 3, "2"},
		{"wizard", wizard, "Slots 3rd", 5, "2"},
		{"wizard", wizard, "Slots 3rd", 4, ""}, // zero slots are left out
	} {
		if v := tt.d.Table[tt.col][tt.lvl]; v != tt.want {
			t.Errorf("%s %s at %d: %q, want %q (table %v)", tt.name, tt.col, tt.lvl, v, tt.want, tt.d.Table)
		}
	}

	// References: a class feature's text, a subclass feature's, an
	// optional feature's name.
	for _, tt := range []struct {
		key, feature, want string
	}{
		{"class/fighter", "Tactical Mind", "**Second Wind.** As a Bonus Action"},
		{"class/fighter", "Fighting Style", "\n\nArchery"},
		{"subclass/champion", "Champion", "**Improved Critical.** Your attack rolls"},
		{"subclass/champion", "Champion", "**Remarkable Athlete.** You have Advantage"},
	} {
		f, _ := feature(got[tt.key].Data.(ClassData), tt.feature)
		if !strings.Contains(f.Desc, tt.want) {
			t.Errorf("%s %s: %q, want %q", tt.key, tt.feature, f.Desc, tt.want)
		}
	}
}

func TestFiveToolsClassCell(t *testing.T) {
	for _, tt := range []struct{ in, want string }{
		{`2`, "2"},
		{`"{@dice 1d4}"`, "1d4"},
		{`{"type":"bonus","value":2}`, "+2"},
		{`{"type":"bonusSpeed","value":10}`, "+10 ft."},
		{`{"type":"dice","toRoll":[{"number":1,"faces":6}]}`, "1d6"},
		{`"—"`, "—"},
	} {
		if got := ftClassCell(json.RawMessage(tt.in)); got != tt.want {
			t.Errorf("%s: %q, want %q", tt.in, got, tt.want)
		}
	}
}

// TestFiveToolsClassLevelUp feeds the classes to level-up the way the web
// handlers do (classFeatures and withSubclass in internal/web).
func TestFiveToolsClassLevelUp(t *testing.T) {
	got := fiveToolsClasses(t)
	info := func(key string) dnd5e.ClassInfo {
		e := got[key]
		d := e.Data.(ClassData)
		return dnd5e.ClassInfo{Slug: e.Slug, Name: e.Name, HitDie: d.HitDie, Caster: d.Caster, Saves: d.Saves, Table: d.Table, Features: dnd5eFeatures(d)}
	}
	fighter := info("class/fighter")
	sh := dnd5e.Sheet{Classes: []dnd5e.ClassLevel{{Slug: "fighter", Name: "Fighter", Level: 1, HitDie: 10, Caster: "none"}}}
	plan, err := sh.Plan(fighter)
	if err != nil || !slices.ContainsFunc(plan.Features, func(f dnd5e.ClassFeature) bool { return f.Name == "Action Surge" }) {
		t.Errorf("fighter 2: %+v %v", plan.Features, err)
	}
	sh.Classes[0].Level = 3
	if plan, _ := sh.Plan(fighter); !plan.ASI {
		t.Errorf("fighter 4: no ASI %+v", plan.Features)
	}

	wizard := info("class/wizard")
	wizard.SubclassFeatures = dnd5eFeatures(got["subclass/evoker"].Data.(ClassData))
	sh = dnd5e.Sheet{Classes: []dnd5e.ClassLevel{{Slug: "wizard", Name: "Wizard", Level: 2, HitDie: 6, Caster: "full"}}}
	plan, err = sh.Plan(wizard)
	if err != nil || !plan.NeedsSubclass || plan.Slots != [9]int{4, 2} || plan.Columns["Prepared Spells"] != "6" {
		t.Errorf("wizard 3: %+v %v", plan, err)
	}
	if !slices.ContainsFunc(plan.Features, func(f dnd5e.ClassFeature) bool { return f.Name == "Evoker" }) {
		t.Errorf("wizard 3: no subclass feature %+v", plan.Features)
	}
}

// dnd5eFeatures is classFeatures in internal/web.
func dnd5eFeatures(d ClassData) []dnd5e.ClassFeature {
	out := make([]dnd5e.ClassFeature, 0, len(d.Features))
	for _, f := range d.Features {
		out = append(out, dnd5e.ClassFeature{Name: f.Name, Description: f.Desc, Levels: f.Levels, Subclass: f.Subclass})
	}
	return out
}
