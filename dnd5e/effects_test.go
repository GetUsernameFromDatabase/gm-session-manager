package dnd5e

import (
	"reflect"
	"testing"
)

// TestParseEffects runs the parser over the 2024 SRD's own sentences.
func TestParseEffects(t *testing.T) {
	table := map[string]map[int]string{"Rages": {1: "2", 3: "3", 6: "4"}, "Focus Points": {1: "-", 2: "2", 3: "3"}, "Channel Divinity": {2: "2", 6: "3"}}
	byLevel := func(m map[int]int) Uses { return Uses{ByLevel: m, Class: "Barbarian"} }
	tests := []struct {
		name, desc string
		want       FeatureEffects
	}{
		{"Bardic Inspiration", "You can supernaturally inspire others. **Using Bardic Inspiration.** As a Bonus Action, you can inspire another creature within 60 feet of yourself. **Number of Uses.** You can confer a Bardic Inspiration die a number of times equal to your Charisma modifier (minimum of once), and you regain all expended uses when you finish a Long Rest.",
			FeatureEffects{Action: "bonus", Uses: Uses{Ability: "cha", Reset: "long"}}},
		{"Rage", "You can enter it as a Bonus Action if you aren't wearing Heavy armor. You can enter your Rage the number of times shown for your Barbarian level in the Rages column of the Barbarian Features table. You regain one expended use when you finish a Short Rest, and you regain all expended uses when you finish a Long Rest.",
			FeatureEffects{Action: "bonus", Uses: Uses{ByLevel: map[int]int{1: 2, 3: 3, 6: 4}, Class: "Barbarian", Reset: "short-one"}}},
		{"Second Wind", "As a Bonus Action, you can use it to regain Hit Points equal to 1d10 plus your Fighter level. You can use this feature twice. You regain one expended use when you finish a Short Rest, and you regain all expended uses when you finish a Long Rest.",
			FeatureEffects{Action: "bonus", Uses: Uses{Fixed: 2, Reset: "short-one"}}},
		{"Action Surge", "On your turn, you can take one additional action. Once you use this feature, you can't do so again until you finish a Short or Long Rest.",
			FeatureEffects{Uses: Uses{Fixed: 1, Reset: "short"}}},
		{"Stonecunning", "As a Bonus Action, you gain Tremorsense. You can use this Bonus Action a number of times equal to your Proficiency Bonus, and you regain all expended uses when you finish a Long Rest.",
			FeatureEffects{Action: "bonus", Uses: Uses{Prof: true, Reset: "long"}}},
		{"Fast Movement", "Your speed increases by 10 feet while you aren't wearing Heavy armor.",
			FeatureEffects{Mods: []CustomMod{{Target: ModSpeed, Value: 10, Note: "Fast Movement", Source: "Fast Movement", When: "no-heavy-armor"}}}},
		{"Unarmored Defense", "While you aren't wearing any armor, your base Armor Class equals 10 plus your Dexterity and Constitution modifiers. You can use a Shield and still gain this benefit.",
			FeatureEffects{Unarmored: "con"}},
		{"Unarmored Defense", "While you aren't wearing armor or wielding a Shield, your base Armor Class equals 10 plus your Dexterity and Wisdom modifiers.",
			FeatureEffects{Unarmored: "wis", NoShield: true}},
		{"Divine Spark", "As a Magic action, you point your Holy Symbol at another creature.", FeatureEffects{Action: "magic"}},
		{"Darkvision", "You have Darkvision with a range of 60 feet.", FeatureEffects{}},
		{"Monk's Focus", "This energy is represented by Focus Points. Your Monk level determines the number of points you have, as shown in the Focus Points column of the Monk Features table. When you expend a Focus Point, it is unavailable until you finish a Short or Long Rest, at the end of which you regain all your expended points.",
			FeatureEffects{Uses: func() Uses { u := byLevel(map[int]int{2: 2, 3: 3}); u.Reset = "short"; return u }()}},
		{"Channel Divinity", "You can use this class's Channel Divinity twice. You regain one of its expended uses when you finish a Short Rest, and you regain all expended uses when you finish a Long Rest. You gain additional uses when you reach certain Cleric levels, as shown in the Channel Divinity column of the Cleric Features table.",
			FeatureEffects{Uses: func() Uses { u := byLevel(map[int]int{2: 2, 6: 3}); u.Reset = "short-one"; return u }()}},
		{"Martial Arts", "You can roll 1d6 in place of the normal damage of your Unarmed Strike. This die changes as you gain Monk levels, as shown in the Martial Arts column of the Monk Features table.", FeatureEffects{}},
		{"Rage Damage", "You gain a bonus to the damage that increases as you gain levels as a Barbarian, as shown in the Rage Damage column of the Barbarian Features table.", FeatureEffects{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseEffects(tt.name, tt.desc, table, "Barbarian"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %+v\nwant %+v", got, tt.want)
			}
		})
	}
}

// TestExplainSums checks each stat equals its breakdown and armor rules.
func TestExplainSums(t *testing.T) {
	base := NewSheet()
	base.Scores["dex"], base.Scores["con"], base.Scores["wis"], base.Scores["str"] = 16, 14, 12, 10
	chain := &WornArmor{Name: "Chain Mail", Type: "heavy", Base: 16, StrReq: 13, Stealth: true}
	studded := &WornArmor{Name: "Studded Leather", Type: "light", Base: 12, AddDex: true}
	breast := &WornArmor{Name: "Breastplate", Type: "medium", Base: 14, AddDex: true, DexCap: 2}
	barb := Feature{Name: "Unarmored Defense", Unarmored: "con"}
	monk := Feature{Name: "Unarmored Defense", Unarmored: "wis", NoShield: true}
	fast := CustomMod{Target: ModSpeed, Value: 10, Note: "Fast Movement", Source: "Fast Movement", When: "no-heavy-armor"}
	tests := []struct {
		name         string
		armor        *WornArmor
		shield       bool
		features     []Feature
		mods         []CustomMod
		ac, speed    int
		tooHeavy     bool
		stealthWorse bool
	}{
		{"unarmored adds dex", nil, false, nil, nil, 13, 30, false, false},
		{"chain mail, weak", chain, false, nil, nil, 16, 20, true, true},
		{"studded + shield", studded, true, nil, nil, 17, 30, false, false},
		{"breastplate caps dex", breast, false, nil, nil, 16, 30, false, false},
		{"barbarian unarmored + shield", nil, true, []Feature{barb}, nil, 17, 30, false, false},
		{"monk with shield loses it", nil, true, []Feature{monk}, nil, 15, 30, false, false},
		{"fast movement in light armor", studded, false, nil, []CustomMod{fast}, 15, 40, false, false},
		{"fast movement off in heavy armor", chain, false, nil, []CustomMod{fast}, 16, 20, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := base
			s.Armor, s.Shield, s.Features, s.Modifiers = tt.armor, tt.shield, tt.features, tt.mods
			if got := s.ArmorClass(); got != tt.ac {
				t.Errorf("AC %d, want %d: %+v", got, tt.ac, s.Explain(ModAC))
			}
			if got := s.WalkSpeed(); got != tt.speed {
				t.Errorf("speed %d, want %d: %+v", got, tt.speed, s.Explain(ModSpeed))
			}
			if s.ArmorTooHeavy() != tt.tooHeavy || s.StealthDisadvantage() != tt.stealthWorse {
				t.Errorf("too heavy %v, stealth %v", s.ArmorTooHeavy(), s.StealthDisadvantage())
			}
		})
	}
	s := base
	s.SkillProfs, s.Expertise, s.SaveProfs = []string{"perception"}, []string{"stealth"}, []string{"dex"}
	s.Weapons = []Weapon{{Name: "Rapier", Ability: "finesse", Proficient: true}}
	s.Modifiers = []CustomMod{{Target: "skill.stealth", Value: 1, Note: "Cloak"}, {Target: ModPassive, Value: 5, Note: "Observant"}}
	for target, want := range map[string]int{
		ModInit: 3, ModPassive: 18, "save.dex": 5, "skill.stealth": 8, "attack.Rapier": 5,
	} {
		if got := Sum(s.Explain(target)); got != want {
			t.Errorf("%s = %d, want %d: %+v", target, got, want, s.Explain(target))
		}
	}
}

func TestArmorAndFeatureUses(t *testing.T) {
	c := Character{Name: "Tess", MaxHP: 10, HP: 10, Sheet: NewSheet()}
	c.Sheet.Scores["cha"] = 14
	c.Sheet.Features = []Feature{{Name: "Bardic Inspiration"}, {Name: "Second Wind", Uses: Uses{Fixed: 2, Reset: "short-one"}}}
	c.Sheet.Modifiers = []CustomMod{{Target: ModSpeed, Value: 10, Note: "Bardic Inspiration", Source: "Bardic Inspiration"}}
	run := func(line string, env Env) error {
		t.Helper()
		cmd, err := Parse(line)
		if err != nil {
			return err
		}
		_, err = cmd.Apply(&c, env)
		return err
	}
	chain := WornArmor{Slug: "chain-mail", Name: "Chain Mail", Type: "heavy", Base: 16}
	for _, step := range []struct {
		line string
		env  Env
		ok   bool
	}{
		{"armor wear chain-mail", Env{Armor: chain}, true},
		{"armor wear shield", Env{IsShield: true}, true},
		{`feature set "Bardic Inspiration" action=bonus uses=cha`, Env{}, true},
		{`feature set "Bardic Inspiration" uses=often`, Env{}, false},
		{`use "Bardic Inspiration"`, Env{}, true},
		{`use "Bardic Inspiration"`, Env{}, true},
		{`use "Bardic Inspiration"`, Env{}, false}, // Charisma +2: two uses
		{`use "Second Wind"`, Env{}, true},
		{`use "Second Wind"`, Env{}, true},
		{"rest short", Env{}, true},
	} {
		if err := run(step.line, step.env); (err == nil) != step.ok {
			t.Fatalf("%s: %v", step.line, err)
		}
	}
	bi, sw := c.Sheet.Features[0], c.Sheet.Features[1]
	if c.Sheet.ArmorClass() != 18 || bi.Action != "bonus" || bi.Uses.Used != 2 || bi.Uses.Reset != "long" || sw.Uses.Used != 1 {
		t.Errorf("after short rest: AC %d %+v %+v", c.Sheet.ArmorClass(), bi, sw)
	}
	for _, line := range []string{"rest long", "armor off", "shield off", `feature remove "Bardic Inspiration"`} {
		if err := run(line, Env{}); err != nil {
			t.Fatalf("%s: %v", line, err)
		}
	}
	if c.Sheet.Features[0].Uses.Used != 0 || c.Sheet.ArmorClass() != 10 || len(c.Sheet.Modifiers) != 0 {
		t.Errorf("after long rest, armor off, feature removed: %+v AC %d mods %+v", c.Sheet.Features, c.Sheet.ArmorClass(), c.Sheet.Modifiers)
	}
}
