package dnd5e

import (
	"testing"
)

// geared carries chain mail, a shield and a stowed longsword, and wears
// leather with a readied dagger.
func geared() Character {
	c := fighter()
	leather, chain := WornArmor{Name: "Leather", Type: "light", Base: 11, AddDex: true}, WornArmor{Name: "Chain Mail", Type: "heavy", Base: 16}
	sword, dagger := Weapon{Name: "Longsword", Ability: "str", Dice: "1d8"}, Weapon{Name: "Dagger", Ability: "finesse", Dice: "1d4"}
	c.Sheet.Armor, c.Sheet.Weapons = &leather, []Weapon{dagger}
	c.Sheet.Modifiers = []CustomMod{{Target: ModAttack + ".Longsword", Value: 1, Note: "+1 sword"}}
	c.Sheet.Inventory = append(c.Sheet.Inventory,
		Item{Name: "Leather", Qty: 1, Armor: &leather}, Item{Name: "Chain Mail", Qty: 1, Armor: &chain},
		Item{Name: "Shield", Qty: 1, Shield: true},
		Item{Name: "Longsword", Qty: 1, Weapon: &sword}, Item{Name: "Dagger", Qty: 2, Weapon: &dagger})
	return c
}

func TestGearCommands(t *testing.T) {
	chainEnv := Env{Armor: WornArmor{Name: "Plate", Type: "heavy", Base: 18}}
	tests := []struct {
		input   string
		env     Env
		check   func(s Sheet) bool
		wantErr bool
	}{
		{`armor wear "chain mail"`, Env{}, func(s Sheet) bool { return s.Armor.Name == "Chain Mail" && s.ArmorClass() == 16 }, false},
		{"armor wear plate", chainEnv, func(s Sheet) bool {
			i := findItem(s.Inventory, "Plate")
			return s.Armor.Name == "Plate" && i >= 0 && s.Worn(s.Inventory[i])
		}, false},
		{"armor wear shield", Env{}, func(s Sheet) bool { return s.Shield && len(s.Inventory) == 6 }, false},
		{"armor off", Env{}, func(s Sheet) bool { return s.Armor == nil && findItem(s.Inventory, "Leather") >= 0 }, false},
		{"shield on", Env{}, func(s Sheet) bool { return s.Shield && len(s.Inventory) == 6 }, false},
		{"weapon ready longsword", Env{}, func(s Sheet) bool { return len(s.Weapons) == 2 && s.AttackBonus(s.Weapons[1]) == 0+1 }, false},
		{"weapon ready torch", Env{}, nil, true},
		{"weapon stow dagger", Env{}, func(s Sheet) bool { return len(s.Weapons) == 0 && findItem(s.Inventory, "Dagger") >= 0 }, false},
		{"weapon add Spear str 1d6 piercing", Env{}, func(s Sheet) bool {
			i := findItem(s.Inventory, "Spear")
			return len(s.Weapons) == 2 && i >= 0 && s.Ready(s.Inventory[i])
		}, false},
		{"item remove Leather", Env{}, func(s Sheet) bool { return s.Armor == nil }, false},
		{"item remove Dagger", Env{}, func(s Sheet) bool { return len(s.Weapons) == 1 }, false},
		{"item set Dagger 0", Env{}, func(s Sheet) bool { return len(s.Weapons) == 0 }, false},
		{"set ac 2", Env{}, func(s Sheet) bool { return s.ArmorClass() == 11+0+2 }, false},
		{"set ac -1", Env{}, func(s Sheet) bool { return s.ArmorClass() == 10 }, false},
		{"set size small", Env{}, func(s Sheet) bool { return s.Size == "Small" }, false},
		{"set size tall", Env{}, nil, true},
		{"set vision Darkvision 60 ft", Env{}, func(s Sheet) bool { return s.Vision == "Darkvision 60 ft" }, false},
		{"set vision none", Env{}, func(s Sheet) bool { return s.Vision == "" }, false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			c := geared()
			cmd, err := Parse(tt.input)
			if err == nil {
				_, err = cmd.Apply(&c, tt.env)
			}
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && !tt.check(c.Sheet) {
				t.Errorf("unexpected sheet: %+v", c.Sheet)
			}
		})
	}
}

func TestStowKeepsMods(t *testing.T) {
	c := geared()
	for _, in := range []string{"weapon ready Longsword", "weapon stow Longsword", "weapon ready Longsword"} {
		cmd, err := Parse(in)
		if err == nil {
			_, err = cmd.Apply(&c, Env{})
		}
		if err != nil {
			t.Fatal(in, err)
		}
	}
	if len(c.Sheet.Modifiers) != 1 || c.Sheet.AttackBonus(c.Sheet.Weapons[1]) != 1 {
		t.Errorf("mods lost: %+v", c.Sheet)
	}
}

// TestUpgradeSheet checks old sheets get gear in the inventory and their
// hand-set AC as a bonus, keeping the AC they had, and that it runs once.
func TestUpgradeSheet(t *testing.T) {
	tests := []struct {
		name, raw string
		ac, bonus int
		items     int
	}{
		{"unarmored base", `{"abilities":{"dex":14},"ac":13}`, 13, 1, 0},
		{"default base gets dex", `{"abilities":{"dex":14},"ac":10}`, 12, 0, 0},
		{"armor ignored base", `{"abilities":{"dex":14},"ac":15,"armor":{"name":"Chain Mail","type":"heavy","base":16},"shield":true}`, 18, 0, 2},
		{"weapons listed", `{"weapons":[{"name":"Club","ability":"str","dice":"1d4"}],"inventory":[{"name":"Club","qty":1}]}`, 10, 0, 1},
		{"current sheet", `{"abilities":{"dex":14},"ac_bonus":1}`, 13, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := ParseSheet(tt.raw)
			if err != nil {
				t.Fatal(err)
			}
			if s.ArmorClass() != tt.ac || s.ACBonus != tt.bonus || len(s.Inventory) != tt.items || s.AC != 0 {
				t.Errorf("AC %d bonus %d items %+v", s.ArmorClass(), s.ACBonus, s.Inventory)
			}
			raw, _ := s.JSON()
			again, _ := ParseSheet(raw)
			if again.ArmorClass() != tt.ac || len(again.Inventory) != tt.items {
				t.Errorf("not idempotent: %+v", again)
			}
		})
	}
}

func TestExplainDamage(t *testing.T) {
	c := geared()
	c.Sheet.Scores["dex"] = 16
	c.Sheet.Modifiers = append(c.Sheet.Modifiers, CustomMod{Target: ModDamage + ".Dagger", Value: 2, Note: "Sharp"})
	parts := c.Sheet.Explain(ModDamage + ".Dagger")
	if len(parts) != 2 || Sum(parts) != 5 || c.Sheet.DamageLabel(c.Sheet.Weapons[0]) != "1d4+5" {
		t.Errorf("damage parts %+v, label %q", parts, c.Sheet.DamageLabel(c.Sheet.Weapons[0]))
	}
	for _, tt := range []struct {
		a    WornArmor
		dex  int
		want string
	}{
		{WornArmor{Base: 16}, 3, "AC 16, no Dex"},
		{WornArmor{Base: 11, AddDex: true}, 3, "AC 11 + Dex (+3)"},
		{WornArmor{Base: 14, AddDex: true, DexCap: 2}, 3, "AC 14 + Dex, at most +2 (+2)"},
	} {
		if got := ArmorEffect(tt.a, tt.dex); got != tt.want {
			t.Errorf("ArmorEffect = %q, want %q", got, tt.want)
		}
	}
}

func TestDualWield(t *testing.T) {
	dagger := Weapon{Name: "Dagger", Ability: "finesse", Dice: "1d4", Properties: "Finesse, Light, Thrown"}
	short := Weapon{Name: "Shortsword", Ability: "finesse", Dice: "1d6", Properties: "Finesse, Light"}
	long := Weapon{Name: "Longsword", Ability: "str", Dice: "1d8", Properties: "Versatile"}
	tests := []struct {
		name     string
		weapons  []Weapon
		qty, dex int
		style    bool
		can      bool
		offhand  string
	}{
		{"one dagger", []Weapon{dagger}, 1, 16, false, false, ""},
		{"two daggers", []Weapon{dagger}, 2, 16, false, true, "1d4"},
		{"dagger and shortsword", []Weapon{dagger, short}, 1, 16, false, true, "1d4"},
		{"dagger and longsword", []Weapon{dagger, long}, 1, 16, false, false, ""},
		{"negative dex still counts", []Weapon{dagger}, 2, 8, false, true, "1d4-1"},
		{"two-weapon fighting style", []Weapon{dagger}, 2, 16, true, true, "1d4+3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSheet()
			s.Scores["dex"], s.Scores["str"], s.Weapons = tt.dex, min(tt.dex, 10), tt.weapons
			s.Inventory = []Item{{Name: "Dagger", Qty: tt.qty, Weapon: &dagger}}
			if tt.style {
				s.Features = []Feature{{Name: "Two-Weapon Fighting"}}
			}
			if got := s.CanDualWield(); got != tt.can {
				t.Fatalf("CanDualWield = %v, want %v", got, tt.can)
			}
			if tt.can {
				if got := s.OffHandDamage(dagger); got != tt.offhand {
					t.Errorf("off-hand %q, want %q: %+v", got, tt.offhand, s.Explain("offhand.Dagger"))
				}
				if s.DamageLabel(dagger) == tt.offhand && !tt.style && tt.dex > 10 {
					t.Error("main-hand damage lost its modifier")
				}
			}
		})
	}
}
