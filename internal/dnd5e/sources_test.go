package dnd5e

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"gm-session-manager/internal/perm"
)

// oldFighter and newFighter are two copies of a class, as two books print it.
var (
	oldFighter = ClassInfo{
		Slug: "fighter", Name: "Fighter", Source: "PHB", SubclassSlug: "champion", SubclassSource: "PHB", HitDie: 10, Caster: CasterNone,
		Features: []ClassFeature{
			{Name: "Second Wind", Description: "old wind", Levels: []int{1}},
			{Name: "Action Surge", Description: "old surge", Levels: []int{2}},
			{Name: "Old Only", Description: "gone", Levels: []int{2}},
			{Name: "Fighter Subclass", Description: "pick", Levels: []int{3}, Subclass: true},
		},
		SubclassFeatures: []ClassFeature{{Name: "Improved Critical", Description: "old crit", Levels: []int{3}}},
	}
	newFighter = ClassInfo{
		Slug: "fighter", Name: "Fighter", Source: "XPHB", SubclassSlug: "champion", SubclassSource: "XPHB", HitDie: 12,
		Features: []ClassFeature{
			{Name: "Second Wind", Description: "new wind", Levels: []int{1}},
			{Name: "Action Surge", Description: "new surge", Levels: []int{2}},
			{Name: "Tactical Mind", Description: "new mind", Levels: []int{2}},
			{Name: "Fighter Subclass", Description: "pick", Levels: []int{3}, Subclass: true},
			{Name: "Extra Attack", Description: "two attacks", Levels: []int{5}},
		},
		SubclassFeatures: []ClassFeature{
			{Name: "Improved Critical", Description: "new crit", Levels: []int{3}},
			{Name: "Remarkable Athlete", Description: "athletic", Levels: []int{3}},
		},
	}
)

// switchFighter has the old copy's features, one of them edited.
func switchFighter() Character {
	c := fighter()
	s := &c.Sheet
	s.Classes[0].Subclass, s.Classes[0].Source = "Champion", "PHB"
	from := func(name, desc string, lvl int) Feature {
		return Feature{Name: name, Source: "Fighter", Level: lvl, Description: desc}
	}
	s.Features = []Feature{
		{Name: "Second Wind", Source: "Fighter", Level: 1, Description: "old wind", Uses: Uses{Fixed: 2, Used: 1, Reset: "short"}},
		from("Action Surge", "my surge", 2),
		from("Old Only", "gone", 2),
		from("Fighter Subclass", "pick", 3),
		from("Improved Critical", "old crit", 3),
		{Name: "Alert", Source: "Feat"},
	}
	s.Feats = []string{"Alert"}
	s.Species, s.Background = "Elf", "Soldier"
	s.Spells = []Spell{{Slug: "fire-bolt", Name: "Firebolt", Level: 0}}
	s.Armor = &WornArmor{Slug: "chain-mail", Name: "Chain Mail", Type: "heavy", Base: 16}
	return c
}

func names(fs []SwitchFeature) []string {
	out := []string{}
	for _, f := range fs {
		s := fmt.Sprintf("%s/%d", f.Name, f.Level)
		if f.Dropped {
			s += "!"
		}
		out = append(out, s)
	}
	return out
}

func TestPlanSourceSwitch(t *testing.T) {
	old := oldFighter
	tests := []struct {
		name                string
		kind, slug          string
		old                 *ClassInfo
		replace, add, keeps []string
	}{
		{"class", RefClass, "fighter", &old,
			[]string{"Second Wind/1", "Old Only/2!", "Fighter Subclass/3"}, []string{"Tactical Mind/2"}, []string{"Action Surge/2"}},
		{"old copy gone", RefClass, "fighter", nil,
			[]string{}, []string{"Tactical Mind/2"}, []string{"Second Wind/1", "Action Surge/2", "Old Only/2", "Fighter Subclass/3"}},
		{"subclass", RefSubclass, "champion", &old,
			[]string{"Improved Critical/3"}, []string{"Remarkable Athlete/3"}, []string{}},
		{"pin-only kind", RefSpell, "fire-bolt", &old, []string{}, []string{}, []string{}},
		{"unknown class", RefClass, "wizard", &old, []string{}, []string{}, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PlanSourceSwitch(switchFighter().Sheet, tt.kind, tt.slug, tt.old, newFighter)
			for _, got := range []struct {
				what      string
				got, want []string
			}{{"replace", names(p.Replace), tt.replace}, {"add", names(p.Add), tt.add}, {"keep", names(p.Keep), tt.keeps}} {
				if !reflect.DeepEqual(got.got, got.want) {
					t.Errorf("%s = %v, want %v", got.what, got.got, got.want)
				}
			}
		})
	}
}

func TestSwitchSource(t *testing.T) {
	old := oldFighter
	classEnv := Env{SourceOld: &old, SourceNew: newFighter}
	feature := func(c Character, name string) *Feature {
		i := slices.IndexFunc(c.Sheet.Features, func(f Feature) bool { return f.Name == name })
		if i < 0 {
			return nil
		}
		return &c.Sheet.Features[i]
	}
	tests := []struct {
		input   string
		env     Env
		check   func(c Character) bool
		wantErr bool
	}{
		{"source class fighter XPHB", classEnv, func(c Character) bool {
			sw, cl := feature(c, "Second Wind"), c.Sheet.Classes[0]
			return sw.Description == "new wind" && sw.Uses.Used == 1 && // untouched: replaced, uses spent kept
				feature(c, "Action Surge").Description == "my surge" && // edited: kept
				feature(c, "Old Only") == nil && // not in the new copy: dropped
				feature(c, "Tactical Mind").Level == 2 && feature(c, "Extra Attack") == nil &&
				feature(c, "Improved Critical").Description == "old crit" &&
				cl.HitDie == 12 && cl.Caster == CasterNone && cl.Source == "XPHB" && cl.SubclassSource == ""
		}, false},
		{"source class fighter XPHB", Env{SourceNew: newFighter}, func(c Character) bool {
			return feature(c, "Second Wind").Description == "old wind" && feature(c, "Old Only") != nil &&
				feature(c, "Tactical Mind") != nil && c.Sheet.Classes[0].Source == "XPHB"
		}, false},
		{"source subclass champion XPHB", classEnv, func(c Character) bool {
			cl := c.Sheet.Classes[0]
			return feature(c, "Improved Critical").Description == "new crit" && feature(c, "Remarkable Athlete").Level == 3 &&
				feature(c, "Second Wind").Description == "old wind" &&
				cl.SubclassSlug == "champion" && cl.SubclassSource == "XPHB" && cl.Source == "PHB" && cl.HitDie == 10
		}, false},
		{"source spell fire-bolt XPHB", Env{SourceSpell: Spell{Slug: "fire-bolt", Name: "Fire Bolt"}}, func(c Character) bool {
			return c.Sheet.Spells[0] == Spell{Slug: "fire-bolt", Name: "Fire Bolt", Source: "XPHB"}
		}, false},
		{"source species elf XPHB", Env{}, func(c Character) bool { return c.Sheet.SpeciesSource == "XPHB" }, false},
		{"source background soldier XPHB", Env{}, func(c Character) bool { return c.Sheet.BackgroundSource == "XPHB" }, false},
		{"source feat alert XPHB", Env{}, func(c Character) bool { return c.Sheet.FeatSources["alert"] == "XPHB" }, false},
		{"source armor chain-mail XPHB", Env{}, func(c Character) bool { return c.Sheet.Armor.Source == "XPHB" }, false},
		{"source class wizard XPHB", Env{SourceNew: ClassInfo{Slug: "wizard"}}, nil, true},
		{"source class fighter XPHB", Env{}, nil, true},
		{"source spell fire-bolt XPHB", Env{}, nil, true},
		{"source spell shield XPHB", Env{SourceSpell: Spell{Slug: "shield", Name: "Shield"}}, nil, true},
		{"source feat lucky XPHB", Env{}, nil, true},
		{"source armor plate XPHB", Env{}, nil, true},
		{"source", Env{}, nil, true},
		{"source class fighter", Env{}, nil, true},
		{"source race elf XPHB", Env{}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			c := switchFighter()
			cmd, err := Parse(tt.input)
			if err == nil {
				if cmd.Scope != perm.ScopeSources || cmd.SourceRef.Kind == "" {
					t.Errorf("scope %v ref %+v", cmd.Scope, cmd.SourceRef)
				}
				_, err = cmd.Apply(&c, tt.env)
			}
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !tt.check(c) {
				t.Errorf("unexpected result: %+v", c.Sheet)
			}
		})
	}
}
