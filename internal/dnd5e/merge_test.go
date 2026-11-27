package dnd5e

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

// TestMergeGroupsCoverSheet: every key of Sheet is in exactly one group and no
// group lists a key Sheet lacks.
func TestMergeGroupsCoverSheet(t *testing.T) {
	var keys []string
	for f := range reflect.TypeFor[Sheet]().Fields() {
		keys = append(keys, strings.Split(f.Tag.Get("json"), ",")[0])
	}
	for _, k := range keys {
		n := 0
		for _, g := range MergeGroups {
			if slices.Contains(GroupKeys(g), k) {
				n++
			}
		}
		if n != 1 {
			t.Errorf("sheet key %q is in %d groups, want 1", k, n)
		}
	}
	for _, g := range MergeGroups {
		for _, k := range GroupKeys(g) {
			if !slices.Contains(keys, k) {
				t.Errorf("group %s lists %q, which Sheet does not have", g, k)
			}
		}
		if GroupLabel(g) == "" {
			t.Errorf("group %s has no label", g)
		}
	}
}

func mergeChar(t *testing.T, change func(*Character)) Character {
	t.Helper()
	s, err := ParseSheet("")
	if err != nil {
		t.Fatal(err)
	}
	s.Classes = []ClassLevel{{Slug: "wizard", Name: "Wizard", Level: 4, HitDie: 6, Caster: CasterFull}}
	s.Spells = []Spell{{Slug: "shield", Name: "Shield", Level: 1}}
	s.Inventory = []Item{{Name: "Rope", Qty: 1}}
	s.Coins.GP = 30
	c := Character{Name: "Aria", HP: 20, MaxHP: 27, Sheet: s}
	if change != nil {
		change(&c)
	}
	return c
}

func TestMerge(t *testing.T) {
	base := mergeChar(t, nil)
	levelUp := func(c *Character) { c.Sheet.Classes[0].Level = 5; c.MaxHP = 33 }
	spend := func(c *Character) { c.Sheet.Coins.GP = 12 }
	rename := func(c *Character) { c.Name = "Aria Vale" }
	heal := func(c *Character) { c.HP = 33 }
	hurt := func(c *Character) { c.HP = 12 }

	tests := []struct {
		name        string
		table, home func(*Character)
		picks       map[string]string
		want        map[string]string // group -> state; others unchanged
		check       func(t *testing.T, m Character)
	}{
		{"nothing changed", nil, nil, nil, nil, func(t *testing.T, m Character) {
			if m.Name != "Aria" || m.HP != 20 {
				t.Errorf("merged %+v", m)
			}
		}},
		{"home only", nil, func(c *Character) { levelUp(c); spend(c) }, nil,
			map[string]string{GroupBuild: SideHome, GroupGear: SideHome},
			func(t *testing.T, m Character) {
				if m.Sheet.Level() != 5 || m.MaxHP != 33 || m.Sheet.Coins.GP != 12 {
					t.Errorf("merged %+v", m)
				}
			}},
		{"table only", func(c *Character) { spend(c); rename(c) }, nil, nil,
			map[string]string{GroupGear: SideTable, GroupStory: SideTable},
			func(t *testing.T, m Character) {
				if m.Name != "Aria Vale" || m.Sheet.Coins.GP != 12 {
					t.Errorf("merged %+v", m)
				}
			}},
		{"both alike is unchanged", spend, spend, nil, nil, func(t *testing.T, m Character) {
			if m.Sheet.Coins.GP != 12 {
				t.Errorf("gp %d", m.Sheet.Coins.GP)
			}
		}},
		{"different groups on each side", hurt, levelUp, nil,
			map[string]string{GroupHealth: SideTable, GroupBuild: SideHome},
			func(t *testing.T, m Character) {
				if m.HP != 12 || m.MaxHP != 33 {
					t.Errorf("hp %d/%d", m.HP, m.MaxHP)
				}
			}},
		{"conflict defaults to home", hurt, heal, nil, map[string]string{GroupHealth: MergeConflict}, func(t *testing.T, m Character) {
			if m.HP != 27 { // 33 clamped to the table's unchanged max
				t.Errorf("hp %d, want 27", m.HP)
			}
		}},
		{"conflict picks the table", hurt, heal, map[string]string{GroupHealth: SideTable}, map[string]string{GroupHealth: MergeConflict}, func(t *testing.T, m Character) {
			if m.HP != 12 {
				t.Errorf("hp %d, want 12", m.HP)
			}
		}},
		{"conflict picks home", hurt, heal, map[string]string{GroupHealth: SideHome}, map[string]string{GroupHealth: MergeConflict}, func(t *testing.T, m Character) {
			if m.HP != 27 {
				t.Errorf("hp %d, want 27", m.HP)
			}
		}},
		{"home max HP lets home HP stand", hurt, func(c *Character) { levelUp(c); heal(c) }, map[string]string{GroupHealth: SideHome},
			map[string]string{GroupHealth: MergeConflict, GroupBuild: SideHome},
			func(t *testing.T, m Character) {
				if m.HP != 33 {
					t.Errorf("hp %d, want 33", m.HP)
				}
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table, home := mergeChar(t, tt.table), mergeChar(t, tt.home)
			merged, diffs, err := Merge(base, table, home, tt.picks)
			if err != nil {
				t.Fatal(err)
			}
			for _, d := range diffs {
				want := tt.want[d.Group]
				if want == "" {
					want = MergeUnchanged
				}
				if d.State != want {
					t.Errorf("group %s: state %s, want %s", d.Group, d.State, want)
				}
				if d.State != MergeUnchanged && len(d.Home)+len(d.Table) == 0 {
					t.Errorf("group %s changed but says nothing", d.Group)
				}
			}
			tt.check(t, merged)
		})
	}
}

func TestMergeSummaries(t *testing.T) {
	base := mergeChar(t, nil)
	home := mergeChar(t, func(c *Character) {
		c.Sheet.Classes[0].Level = 5
		c.MaxHP = 33
		c.Sheet.Spells = append(c.Sheet.Spells, Spell{Slug: "fireball", Name: "Fireball", Level: 3})
		c.Sheet.Coins.GP = 12
		c.Sheet.Inventory = []Item{{Name: "Rope", Qty: 1}, {Name: "Lantern", Qty: 1}}
	})
	_, diffs, err := Merge(base, base, home, nil)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string][]string{}
	for _, d := range diffs {
		got[d.Group] = d.Home
	}
	for group, want := range map[string][]string{
		GroupBuild: {"Level 4 → 5", "Max HP 27 → 33", "Spells added: Fireball"},
		GroupGear:  {"Items added: Lantern", "30 gp → 12 gp"},
	} {
		if !slices.Equal(got[group], want) {
			t.Errorf("%s lines %q, want %q", group, got[group], want)
		}
	}
}

func TestMergeInvalid(t *testing.T) {
	base := mergeChar(t, nil)
	home := mergeChar(t, func(c *Character) { c.Name = "" })
	if _, _, err := Merge(base, base, home, nil); err == nil {
		t.Error("an empty name merged")
	}
}

func TestTakeBuild(t *testing.T) {
	c := mergeChar(t, func(c *Character) { c.HP = 27; c.Sheet.Coins.GP = 5 })
	fb := mergeChar(t, func(c *Character) { c.MaxHP = 20; c.Sheet.Classes[0].Level = 3 })
	got, err := TakeBuild(c, fb)
	if err != nil {
		t.Fatal(err)
	}
	if got.Sheet.Level() != 3 || got.MaxHP != 20 || got.HP != 20 || got.Sheet.Coins.GP != 5 {
		t.Errorf("got level %d, hp %d/%d, gp %d", got.Sheet.Level(), got.HP, got.MaxHP, got.Sheet.Coins.GP)
	}
}

func TestSourceUsesAndDrop(t *testing.T) {
	s := mergeChar(t, nil).Sheet
	s.Classes[0].SubclassSource, s.Classes[0].Subclass = "book", "Bladesinging"
	s.SpeciesSource, s.Species = "book", "Elf"
	s.FeatSources = map[string]string{"alert": "book", "tough": "other", "lucky": "book"}
	s.Feats = append(s.Feats, "Alert")
	s.Spells = append(s.Spells, Spell{Slug: "boom", Name: "Booming Blade", Source: "book"}, Spell{Slug: "zap", Name: "Zap", Source: "other"})
	plate := WornArmor{Name: "Plate", Source: "book", Type: "heavy", Base: 18}
	s.Armor = &plate
	s.Inventory = append(s.Inventory, Item{Name: "Plate", Qty: 1, Armor: &plate}, Item{Name: "Chain", Qty: 1, Armor: &WornArmor{Name: "Chain", Source: "other"}})

	want := []string{"subclass Bladesinging", "species Elf", "feat Alert", "feat lucky", "spell Booming Blade", "armor Plate"}
	if got := s.SourceUses("book"); !slices.Equal(got, want) {
		t.Errorf("SourceUses = %q, want %q", got, want)
	}
	if got := s.SourceUses("none"); len(got) != 0 {
		t.Errorf("SourceUses of an unused key = %q", got)
	}
	removed := s.DropSource("book")
	if !slices.Equal(removed, []string{"spell Booming Blade", "armor Plate"}) {
		t.Errorf("DropSource = %q", removed)
	}
	if len(s.Spells) != 2 || s.Armor != nil || len(s.Inventory) != 2 || !s.BuildPinsSource("book") {
		t.Errorf("after drop: %d spells, armor %v, %d items, build pins %v", len(s.Spells), s.Armor, len(s.Inventory), s.BuildPinsSource("book"))
	}
	if got := s.DropSource("book"); len(got) != 0 {
		t.Errorf("second DropSource removed %q", got)
	}
}

func TestRenameSpells(t *testing.T) {
	tests := []struct {
		name   string
		spells []Spell
		want   []string
	}{
		{"unpinned", []Spell{{Slug: "old"}}, []string{"new"}},
		{"homebrew pinned", []Spell{{Slug: "old", Source: "homebrew"}}, []string{"new"}},
		{"another source", []Spell{{Slug: "old", Source: "book"}}, []string{"old"}},
		{"another slug", []Spell{{Slug: "other"}, {Slug: "old"}}, []string{"other", "new"}},
		// "new" is renamed too: the spell that was "old" must not follow it.
		{"chain", []Spell{{Slug: "old"}, {Slug: "new"}}, []string{"new", "new-2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sheet{Spells: tt.spells}
			s.RenameSpells(map[string]string{"old": "new", "new": "new-2"})
			var got []string
			for _, sp := range s.Spells {
				got = append(got, sp.Slug)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("slugs %v, want %v", got, tt.want)
			}
		})
	}
}
