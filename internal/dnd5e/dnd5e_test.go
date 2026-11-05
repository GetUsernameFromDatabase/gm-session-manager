package dnd5e

import (
	"errors"
	"reflect"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/perm"
)

func fighter() Character {
	s := NewSheet()
	s.Scores["con"] = 14
	s.Classes = []ClassLevel{{Slug: "fighter", Name: "Fighter", Level: 3, HitDie: 10, Caster: CasterNone}}
	s.Slots[0] = Slot{Max: 2}
	s.Inventory = []Item{{Name: "Torch", Qty: 3}}
	return Character{Name: "Tess", HP: 20, MaxHP: 28, TempHP: 5, Sheet: s}
}

func TestCommands(t *testing.T) {
	tests := []struct {
		input   string
		scope   perm.CharacterScope
		check   func(c Character, fx Effects) bool
		wantErr bool
	}{
		{"hp -3", perm.ScopeHP, func(c Character, _ Effects) bool { return c.HP == 20 && c.TempHP == 2 }, false},
		{"hp -8", perm.ScopeHP, func(c Character, _ Effects) bool { return c.HP == 17 && c.TempHP == 0 }, false},
		{"hp -100", perm.ScopeHP, func(c Character, _ Effects) bool { return c.HP == -75 && c.TempHP == 0 }, false},
		{"hp +100", perm.ScopeHP, func(c Character, _ Effects) bool { return c.HP == 28 }, false},
		{"hp =7", perm.ScopeHP, func(c Character, _ Effects) bool { return c.HP == 7 }, false},
		{"hp 5", 0, nil, true},
		{`hp ""`, 0, nil, true},
		{"hp", 0, nil, true},
		{"temp 8", perm.ScopeHP, func(c Character, _ Effects) bool { return c.TempHP == 8 }, false},
		{"temp -10", perm.ScopeHP, func(c Character, _ Effects) bool { return c.TempHP == 0 }, false},
		{"maxhp +2", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.MaxHP == 30 }, false},
		{"maxhp =10", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.MaxHP == 10 && c.HP == 10 }, false},
		{"toggle prone", perm.ScopeConditions, func(_ Character, fx Effects) bool {
			return len(fx.Conditions) == 1 && fx.Conditions[0] == CondOp{Op: CondToggle, Slug: "prone"}
		}, false},
		{"toggle condition Prone", perm.ScopeConditions, func(_ Character, fx Effects) bool { return fx.Conditions[0].Slug == "prone" }, false},
		{"toggle flying", 0, nil, true},
		{"condition add poisoned spider bite", perm.ScopeConditions, func(_ Character, fx Effects) bool {
			return fx.Conditions[0] == CondOp{Op: CondAdd, Slug: "poisoned", Note: "spider bite"}
		}, false},
		{"concentrate Bless", perm.ScopeConditions, func(_ Character, fx Effects) bool {
			return fx.Conditions[0] == CondOp{Op: CondAdd, Slug: TrackConcentration, Note: "Bless"}
		}, false},
		{"exhaustion +1", perm.ScopeConditions, func(_ Character, fx Effects) bool { return fx.Conditions[0].Delta == 1 }, false},
		{"exhaustion =3", 0, nil, true},
		{"deathsave fail", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.DeathSaves.Failures == 1 }, false},
		{"slot use 1", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Slots[0].Used == 1 }, false},
		{"slot use 2", 0, nil, true},
		{"slot use 10", 0, nil, true},
		{"slot max 2 3", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Slots[1].Max == 3 }, false},
		{"rest long", perm.ScopeSheet, func(c Character, fx Effects) bool {
			return c.HP == 28 && c.TempHP == 0 && fx.Conditions[0].Op == CondExhaustion
		}, false},
		{`item add "Rope (50 ft)" 2`, perm.ScopeSheet, func(c Character, _ Effects) bool {
			return len(c.Sheet.Inventory) == 2 && c.Sheet.Inventory[1].Name == "Rope (50 ft)" && c.Sheet.Inventory[1].Qty == 2
		}, false},
		{"item add torch", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Inventory[0].Qty == 4 }, false},
		{"item remove Torch 3", perm.ScopeSheet, func(c Character, _ Effects) bool { return len(c.Sheet.Inventory) == 0 }, false},
		{"item remove Shield", 0, nil, true},
		{`item add "unclosed`, 0, nil, true},
		// Names quoted by strconv.Quote, as the sheet's buttons send them.
		{`item add "Ring of \"Luck\" \\ x y"`, perm.ScopeSheet, func(c Character, _ Effects) bool {
			return c.Sheet.Inventory[1].Name == "Ring of \"Luck\" \\ x y"
		}, false},
		{"coin gp +10", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Coins.GP == 10 }, false},
		{"coin gp -1", 0, nil, true},
		{"coin gp +10 SP =3", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Coins.GP == 10 && c.Sheet.Coins.SP == 3 }, false},
		{"coin gp +10 sp -1", 0, nil, true},
		{"coin gp", 0, nil, true},
		{"item set torch 7", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Inventory[0].Qty == 7 }, false},
		{"item set Torch 0", perm.ScopeSheet, func(c Character, _ Effects) bool { return len(c.Sheet.Inventory) == 0 }, false},
		{"item set Torch", 0, nil, true},
		{"item add Torch 0", 0, nil, true},
		{"set ability str 16", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Scores["str"] == 16 }, false},
		{"set ability str 31", 0, nil, true},
		{"set name Tess the Bold", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Name == "Tess the Bold" }, false},
		{"prof add skill athletics", perm.ScopeSheet, func(c Character, _ Effects) bool { return len(c.Sheet.SkillProfs) == 1 }, false},
		{"prof add skill flying", 0, nil, true},
		{"prof set skill stealth expertise", perm.ScopeSheet, func(c Character, _ Effects) bool {
			return c.Sheet.ProfLevel("stealth") == ProfExpertise && slices.Contains(c.Sheet.SkillProfs, "stealth")
		}, false},
		{"prof set skill stealth none", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.ProfLevel("stealth") == ProfNone }, false},
		{"prof set save con proficient", perm.ScopeSheet, func(c Character, _ Effects) bool { return slices.Contains(c.Sheet.SaveProfs, "con") }, false},
		{"prof set save con expertise", 0, nil, true},
		{"prof set skill flying proficient", 0, nil, true},
		{"set backstory Raised by wolves", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Backstory == "Raised by wolves" }, false},
		{"weapon add Longsword str 1d8 slashing prof", perm.ScopeSheet, func(c Character, _ Effects) bool {
			return c.Sheet.Weapons[0] == Weapon{Name: "Longsword", Ability: "str", Dice: "1d8", DamageType: "slashing", Proficient: true}
		}, false},
		{"weapon add Club luck 1d4", 0, nil, true},
		{"weapon remove Longsword", 0, nil, true},
		{`mod add skill.stealth +2 "Cloak of Elvenkind"`, perm.ScopeSheet, func(c Character, _ Effects) bool {
			return c.Sheet.Bonus("skill.stealth") == 2 && c.Sheet.Modifiers[0].Note == "Cloak of Elvenkind"
		}, false},
		{"mod add ac +1", 0, nil, true},
		{"mod add ac 0 nothing", 0, nil, true},
		{"mod add flying +1 wings", 0, nil, true},
		{"mod remove ac Ring", 0, nil, true},
		{"inspiration", perm.ScopeSheet, func(c Character, _ Effects) bool { return c.Sheet.Inspiration }, false},
		{"dance", 0, nil, true},
		{"", 0, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			c := fighter()
			cmd, err := Parse(tt.input)
			if err == nil {
				var fx Effects
				fx, err = cmd.Apply(&c, Env{})
				if err == nil && cmd.Scope != tt.scope {
					t.Errorf("scope %v, want %v", cmd.Scope, tt.scope)
				}
				if err == nil && !tt.check(c, fx) {
					t.Errorf("unexpected result: %+v %+v", c, fx)
				}
			}
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				if c.HP != fighter().HP || len(c.Sheet.Inventory) != 1 {
					t.Error("failed command changed the character")
				}
			}
		})
	}
}

// The command bar hints from Help and allows only Verbs: they must agree.
func TestHelpVerbs(t *testing.T) {
	for _, h := range Help {
		if verb, _, _ := strings.Cut(h.Syntax, " "); !slices.Contains(Verbs(), verb) {
			t.Errorf("help %q: %q is not a verb", h.Syntax, verb)
		}
	}
}

func TestUnknownCommand(t *testing.T) {
	if _, err := Parse("dance wildly"); !errors.Is(err, ErrUnknownCommand) {
		t.Errorf("err = %v", err)
	}
}

func TestHealingFromZeroResetsDeathSaves(t *testing.T) {
	c := fighter()
	c.HP = 0
	c.Sheet.DeathSaves = DeathSaves{Successes: 1, Failures: 2}
	c.ApplyHP(Heal(3))
	if c.HP != 3 || c.Sheet.DeathSaves != (DeathSaves{}) {
		t.Errorf("got hp %d saves %+v", c.HP, c.Sheet.DeathSaves)
	}
}

func TestCustomMods(t *testing.T) {
	c := fighter()
	for _, in := range []string{
		"weapon add Rapier finesse 1d8 piercing prof", "set ability dex 16",
		`mod add dex +2 "Gloves"`, `mod add attack.Rapier +1 "+1 weapon"`, `mod add damage.Rapier +1 "+1 weapon"`,
		`mod add ac +1 "Ring"`, `mod add ac +1 "Ring"`, `mod add init +5 "Alert"`,
	} {
		cmd, err := Parse(in)
		if err == nil {
			_, err = cmd.Apply(&c, Env{})
		}
		if err != nil {
			t.Fatalf("%s: %v", in, err)
		}
	}
	s, w := c.Sheet, c.Sheet.Weapons[0]
	tests := []struct {
		name      string
		got, want any
	}{
		{"score", s.Score("dex"), 18},
		{"save", s.SaveBonus("dex"), 4},
		{"same note replaces", s.ArmorClass(), 10 + 4 + 1},
		{"initiative", s.Initiative(), 9},
		{"to hit", s.AttackBonus(w), 4 + 2 + 1},
		{"damage", s.DamageLabel(w), "1d8+5 piercing"},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Errorf("%s = %v, want %v", tt.name, tt.got, tt.want)
		}
	}
	for _, tt := range []struct {
		in      string
		wantErr bool
	}{
		{`mod edit ac Ring ac +2 "Ring +2"`, false},
		{`mod edit ac Ring ac +2 again`, true},        // gone after the edit
		{`mod edit init Alert ac +2 "Ring +2"`, true}, // would duplicate
		{`mod edit init Alert flying +2 Wings`, true}, // bad new target
		{`mod edit init Alert init -1 "Slow reflexes"`, false},
	} {
		cmd, err := Parse(tt.in)
		if err == nil {
			_, err = cmd.Apply(&c, Env{})
		}
		if (err != nil) != tt.wantErr {
			t.Errorf("%s: err = %v", tt.in, err)
		}
	}
	if got := c.Sheet.Mods(ModAC); len(got) != 1 || got[0] != (CustomMod{Target: ModAC, Value: 2, Note: "Ring +2"}) || c.Sheet.Bonus(ModInit) != -1 {
		t.Errorf("after edits: %+v", c.Sheet.Modifiers)
	}
	cmd, _ := Parse("weapon remove rapier")
	if _, err := cmd.Apply(&c, Env{}); err != nil || len(c.Sheet.Weapons) != 0 || len(c.Sheet.Mods("attack.Rapier")) != 0 {
		t.Errorf("weapon remove left %+v %+v (%v)", c.Sheet.Weapons, c.Sheet.Modifiers, err)
	}
}

func TestParseRawArgs(t *testing.T) {
	c := fighter()
	const story = "Said \"hi\".\nLeft."
	cmd, err := Parse("set backstory", story)
	if err == nil {
		_, err = cmd.Apply(&c, Env{})
	}
	if err != nil || c.Sheet.Backstory != story {
		t.Errorf("backstory %q (%v)", c.Sheet.Backstory, err)
	}
}

func TestModifier(t *testing.T) {
	for score, want := range map[int]int{1: -5, 7: -2, 8: -1, 9: -1, 10: 0, 11: 0, 12: 1, 15: 2, 20: 5, 30: 10} {
		if got := Modifier(score); got != want {
			t.Errorf("Modifier(%d) = %d, want %d", score, got, want)
		}
	}
}

func TestSlotMaxima(t *testing.T) {
	tests := []struct {
		name      string
		classes   []ClassLevel
		first     int
		highest   int
		pact      int
		pactLevel int
	}{
		{"none", []ClassLevel{{Level: 5, Caster: CasterNone}}, 0, 0, 0, 0},
		{"wizard 1", []ClassLevel{{Level: 1, Caster: CasterFull}}, 2, 1, 0, 0},
		{"wizard 5", []ClassLevel{{Level: 5, Caster: CasterFull}}, 4, 3, 0, 0},
		{"paladin 1 rounds up", []ClassLevel{{Level: 1, Caster: CasterHalf}}, 2, 1, 0, 0},
		{"paladin 5", []ClassLevel{{Level: 5, Caster: CasterHalf}}, 4, 2, 0, 0},
		{"wizard 3 paladin 2", []ClassLevel{{Level: 3, Caster: CasterFull}, {Level: 2, Caster: CasterHalf}}, 4, 2, 0, 0},
		{"wizard 20", []ClassLevel{{Level: 20, Caster: CasterFull}}, 4, 9, 0, 0},
		{"warlock 5", []ClassLevel{{Level: 5, Caster: CasterPact}}, 0, 0, 2, 3},
		{"warlock 11", []ClassLevel{{Level: 11, Caster: CasterPact}}, 0, 0, 3, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slots, pact, pactLevel := SlotMaxima(tt.classes)
			highest := 0
			for i, n := range slots {
				if n > 0 {
					highest = i + 1
				}
			}
			if slots[0] != tt.first || highest != tt.highest || pact != tt.pact || pactLevel != tt.pactLevel {
				t.Errorf("slots %v pact %d@%d", slots, pact, pactLevel)
			}
		})
	}
}

var wizard = ClassInfo{
	Slug: "wizard", Name: "Wizard", HitDie: 6, Caster: CasterFull, Saves: []string{"int", "wis"},
	Features: []ClassFeature{
		{Name: "Arcane Recovery", Levels: []int{1}},
		{Name: "Wizard Subclass", Levels: []int{3}},
		{Name: ASIFeature, Levels: []int{4, 8}},
	},
	Table: map[string]map[int]string{"Cantrips": {1: "3", 4: "4"}},
}

func TestLevelUp(t *testing.T) {
	newbie := Character{Name: "Wren", Sheet: NewSheet()}
	newbie.Sheet.Scores["con"] = 14
	newbie.Sheet.Scores["int"] = 19

	// Level 1: maximum hit die plus Constitution, class saving throws.
	c := newbie
	if err := c.LevelUp(LevelUpChoice{Class: wizard}); err != nil {
		t.Fatal(err)
	}
	if c.MaxHP != 8 || c.HP != 8 || c.Sheet.Level() != 1 || len(c.Sheet.SaveProfs) != 2 || c.Sheet.Slots[0].Max != 2 {
		t.Fatalf("level 1: %+v", c)
	}
	if len(c.Sheet.Features) != 1 || c.Sheet.Features[0].Name != "Arcane Recovery" {
		t.Errorf("features %+v", c.Sheet.Features)
	}

	// Level 2 with a roll, level 3 with a subclass.
	if err := c.LevelUp(LevelUpChoice{Class: wizard, HPRoll: 6}); err != nil {
		t.Fatal(err)
	}
	if c.MaxHP != 16 {
		t.Errorf("rolled hp: max %d", c.MaxHP)
	}
	plan, err := c.Sheet.Plan(wizard)
	if err != nil || !plan.NeedsSubclass || plan.ASI {
		t.Fatalf("plan 3: %+v %v", plan, err)
	}
	if err := c.LevelUp(LevelUpChoice{Class: wizard, Subclass: "Evoker"}); err != nil {
		t.Fatal(err)
	}
	if c.MaxHP != 22 || c.Sheet.Classes[0].Subclass != "Evoker" {
		t.Errorf("level 3: max %d %+v", c.MaxHP, c.Sheet.Classes)
	}

	// Level 4 is an ASI level; refusing to choose both, and the cap at 20.
	plan, _ = c.Sheet.Plan(wizard)
	if !plan.ASI || plan.Columns["Cantrips"] != "4" {
		t.Errorf("plan 4: %+v", plan)
	}
	if err := c.LevelUp(LevelUpChoice{Class: wizard, ASI: []string{"int"}, Feat: "Alert"}); err == nil {
		t.Error("feat and ASI together accepted")
	}
	if err := c.LevelUp(LevelUpChoice{Class: wizard, ASI: []string{"int"}}); err == nil {
		t.Error("score above 20 accepted")
	}
	if err := c.LevelUp(LevelUpChoice{Class: wizard, ASI: []string{"int", "con"}}); err != nil {
		t.Fatal(err)
	}
	if c.Sheet.Scores["int"] != 20 || c.Sheet.Scores["con"] != 15 || c.Sheet.Level() != 4 {
		t.Errorf("level 4 scores %+v", c.Sheet.Scores)
	}

	// Multiclassing adds a class without its saving throws.
	fighterClass := ClassInfo{Slug: "fighter", Name: "Fighter", HitDie: 10, Caster: CasterNone, Saves: []string{"str", "con"}}
	// Unmet prerequisites warn (Fighter needs STR or DEX 13) but do not block.
	if plan, _ := c.Sheet.Plan(fighterClass); len(plan.Warnings) != 1 || !strings.Contains(plan.Warnings[0], "Strength or Dexterity") {
		t.Errorf("multiclass warnings %q", plan.Warnings)
	}
	if plan, _ := c.Sheet.Plan(wizard); len(plan.Warnings) != 0 {
		t.Errorf("same class warned: %q", plan.Warnings)
	}
	if got := c.Sheet.DefaultClass(); got != "wizard" {
		t.Errorf("default class %q", got)
	}
	if err := c.LevelUp(LevelUpChoice{Class: fighterClass}); err != nil {
		t.Fatal(err)
	}
	if got := c.Sheet.DefaultClass(); got != "fighter" {
		t.Errorf("default class after multiclass %q", got)
	}
	if len(c.Sheet.Classes) != 2 || len(c.Sheet.SaveProfs) != 2 || c.Sheet.ClassSummary() != "Wizard 4 (Evoker) / Fighter 1" {
		t.Errorf("multiclass: %+v saves %v", c.Sheet.Classes, c.Sheet.SaveProfs)
	}
	if err := c.LevelUp(LevelUpChoice{Class: wizard, ASI: []string{"str"}}); err == nil {
		t.Error("ASI at a non-ASI level accepted")
	}
}

// TestLevelUpSubclassAndPools covers 2014-style subclass features (marked
// rather than named "… Subclass") and pools read from the class table.
func TestLevelUpSubclassAndPools(t *testing.T) {
	monk := ClassInfo{
		Slug: "monk", Name: "Monk", HitDie: 8, Caster: CasterNone,
		Table: map[string]map[int]string{"Focus Points": {2: "2", 3: "3"}},
		Features: []ClassFeature{
			{Name: "Monk's Focus", Description: "Your Monk level determines the number of points you have, as shown in the Focus Points column of the Monk Features table.", Levels: []int{2}},
			{Name: "Monastic Tradition", Levels: []int{3}, Subclass: true},
		},
		SubclassFeatures: []ClassFeature{{Name: "Hands of Healing", Levels: []int{3}}},
	}
	c := Character{Name: "Tam", Sheet: NewSheet()}
	c.Sheet.Classes = []ClassLevel{{Slug: "monk", Name: "Monk", Level: 2, HitDie: 8, Caster: CasterNone}}
	// Added before the parser knew Focus Points: no uses yet.
	c.Sheet.Features = []Feature{{Name: "Monk's Focus", Source: "Monk", Description: monk.Features[0].Description}}

	plan, err := c.Sheet.Plan(monk)
	if err != nil || !plan.NeedsSubclass || len(plan.Features) != 2 {
		t.Fatalf("plan: %+v %v", plan, err)
	}
	if err := c.LevelUp(LevelUpChoice{Class: monk, Subclass: "Way of Mercy"}); err != nil {
		t.Fatal(err)
	}
	if got := c.Sheet.MaxUses(c.Sheet.Features[0].Uses); got != 3 {
		t.Errorf("focus points %d, uses %+v", got, c.Sheet.Features[0].Uses)
	}
	if !slices.ContainsFunc(c.Sheet.Features, func(f Feature) bool { return f.Name == "Hands of Healing" }) {
		t.Errorf("subclass feature missing: %+v", c.Sheet.Features)
	}
}

func TestLanguages(t *testing.T) {
	c := Character{Name: "Tam", Sheet: NewSheet()}
	for _, line := range []string{"prof add language Elvish", "prof add language elvish", "prof remove language Common"} {
		cmd, err := Parse(line)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := cmd.Apply(&c, Env{}); err != nil {
			t.Fatal(err)
		}
	}
	if !slices.Equal(c.Sheet.Languages, []string{"Elvish"}) {
		t.Errorf("%q", c.Sheet.Languages)
	}
}

// TestDefenses: immunity takes all damage, resistance halves it (rounding
// down) and vulnerability doubles it after that; a conditional defense
// counts only when picked, and Petrified resists untyped damage too.
func TestDefenses(t *testing.T) {
	d := Defenses{List: []Defense{
		{Kind: DefResist, Types: []string{"cold", "fire"}},
		{Kind: DefImmune, Types: []string{"poison"}},
		{Kind: DefVulnerable, Types: []string{"fire", "radiant"}},
		{Kind: DefResist, Types: []string{"bludgeoning"}, When: "from nonmagical attacks"},
	}}
	for _, tt := range []struct {
		n         int
		kind      string
		when      int
		resistAll bool
		want      int
	}{
		{9, "cold", -1, false, 4},
		{9, "poison", -1, false, 0},
		{9, "radiant", -1, false, 18},
		{9, "fire", -1, false, 8}, // resisted and vulnerable
		{9, "", -1, false, 9},
		{9, "bludgeoning", -1, false, 9},
		{9, "bludgeoning", 3, false, 4},
		{9, "cold", 3, false, 4}, // the pick covers bludgeoning only
		{9, "", -1, true, 4},
		{9, "cold", -1, true, 4}, // resistance does not stack
		{9, "radiant", -1, true, 8},
	} {
		d.ResistAll = tt.resistAll
		if got := d.Taken(tt.n, tt.kind, tt.when); got != tt.want {
			t.Errorf("Taken(%d, %q, %d) petrified %v = %d, want %d", tt.n, tt.kind, tt.when, tt.resistAll, got, tt.want)
		}
	}

	for _, tt := range []struct {
		text string
		want []Defense
	}{
		{"cold, fire; bludgeoning, piercing, and slashing from nonmagical attacks", []Defense{
			{Kind: DefResist, Types: []string{"cold", "fire"}},
			{Kind: DefResist, Types: []string{"bludgeoning", "piercing", "slashing"}, When: "from nonmagical attacks"},
		}},
		{"Poison; frightened, poisoned", []Defense{{Kind: DefResist, Types: []string{"poison"}}}},
		{"poison, poisoned, charmed", []Defense{{Kind: DefResist, Types: []string{"poison"}}}},
		{"fire damage", []Defense{{Kind: DefResist, Types: []string{"fire"}}}},
		{"bludgeoning from nonmagical attacks that aren't adamantine", []Defense{
			{Kind: DefResist, Types: []string{"bludgeoning"}, When: "from nonmagical attacks that aren't adamantine"},
		}},
		{"", nil},
	} {
		if got := ParseDefenses(DefResist, tt.text); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ParseDefenses(%q) = %+v, want %+v", tt.text, got, tt.want)
		}
	}

	c := Character{Name: "Tam", HP: 30, MaxHP: 30, Sheet: NewSheet()}
	c.Sheet.Resistances, c.Sheet.Vulnerabilities = []string{"fire"}, []string{"cold"}
	for _, tt := range []struct {
		line      string
		petrified bool
		hp        int
	}{
		{"hp -9 fire", false, 26},
		{"hp -9 Cold", false, 8},
		{"hp -9", true, 4},
		{"hp =30", false, 30},
		{"hp -9", false, 21},
	} {
		cmd, err := Parse(tt.line)
		if err == nil {
			_, err = cmd.Apply(&c, Env{Petrified: tt.petrified})
		}
		if err != nil || c.HP != tt.hp {
			t.Errorf("%s: HP %d, want %d (%v)", tt.line, c.HP, tt.hp, err)
		}
	}
	for _, bad := range []string{"hp -9 pizza", "hp +9 fire", "hp =9 fire"} {
		if _, err := Parse(bad); err == nil {
			t.Errorf("%s parsed", bad)
		}
	}
}

func TestResistances(t *testing.T) {
	c := Character{Name: "Tam", Sheet: NewSheet()}
	for _, tt := range []struct {
		line string
		ok   bool
	}{
		{"prof add resistance Fire", true},
		{"prof add resistance cold", true},
		{"prof add resistance fire", true}, // already there
		{"prof remove resistance cold", true},
		{"prof add resistance pizza", false},
	} {
		cmd, err := Parse(tt.line)
		if err == nil {
			_, err = cmd.Apply(&c, Env{})
		}
		if (err == nil) != tt.ok {
			t.Errorf("%s: %v", tt.line, err)
		}
	}
	if !slices.Equal(c.Sheet.Resistances, []string{"fire"}) {
		t.Errorf("%q", c.Sheet.Resistances)
	}
	for _, tt := range []struct {
		line string
		ok   bool
	}{
		{"prof add immunity poison", true},
		{"prof add vulnerability Radiant", true},
		{"prof add immunity pizza", false},
	} {
		cmd, err := Parse(tt.line)
		if err == nil {
			_, err = cmd.Apply(&c, Env{})
		}
		if (err == nil) != tt.ok {
			t.Errorf("%s: %v", tt.line, err)
		}
	}
	if !slices.Equal(c.Sheet.Immunities, []string{"poison"}) || !slices.Equal(c.Sheet.Vulnerabilities, []string{"radiant"}) {
		t.Errorf("immunities %q, vulnerabilities %q", c.Sheet.Immunities, c.Sheet.Vulnerabilities)
	}
	s := Sheet{Classes: []ClassLevel{{Name: "Warlock"}}}
	if s.SpellAbilityOf(Spell{}) != "cha" || s.SpellAbilityOf(Spell{Ability: "int"}) != "int" {
		t.Error("SpellAbilityOf")
	}
}

func TestCompletions(t *testing.T) {
	s := NewSheet()
	s.Inventory = []Item{{Name: "Torch", Qty: 1}, {Name: "Rope (50 ft)", Qty: 1}}
	s.Weapons = []Weapon{{Name: "Dagger"}}
	s.Features = []Feature{{Name: "Second Wind"}}
	s.Spells = []Spell{{Slug: "fire-bolt", Name: "Fire Bolt"}}
	tests := []struct {
		words      string
		slot, want string
	}{
		{"toggle", "<condition>", "prone"},
		{"item remove", `"<name>"`, "Rope (50 ft)"},
		{"weapon ready", "<name>", "Torch"},
		{"weapon stow", "<name>", "Dagger"},
		{"feature remove", `"<name>"`, "Second Wind"},
		{"use", `"<feature>"`, "Second Wind"},
		{"armor wear", `<slug|"item">`, "Torch"},
		{"prof add skill", "<name>", "stealth"},
		{"prof add save", "<name>", "con"},
		{"spell remove", "<slug>", "fire-bolt"},
		{"concentrate", "<spell>", "Fire Bolt"},
		{"set ability", "<ab>", "str"},
	}
	for _, tt := range tests {
		if got := Completions(s, strings.Fields(tt.words), tt.slot); !slices.Contains(got, tt.want) {
			t.Errorf("%s %s: %q lacks %q", tt.words, tt.slot, got, tt.want)
		}
	}
	for _, tt := range []struct{ words, slot string }{{"prof add other", "<name>"}, {"set name", "<text>"}, {"hp", "<n>"}} {
		if got := Completions(s, strings.Fields(tt.words), tt.slot); got != nil {
			t.Errorf("%s %s: want none, got %q", tt.words, tt.slot, got)
		}
	}
}

func TestPrepareCap(t *testing.T) {
	wizard := ClassLevel{Slug: "wizard", Name: "Wizard", Level: 3, HitDie: 6, Caster: CasterFull}
	cleric := ClassLevel{Slug: "cleric", Name: "Cleric", Level: 1, HitDie: 8, Caster: CasterFull}
	caps := map[string]int{"wizard": 2, "cleric": 2}
	const full = "Wizard prepares 2 spells at this level: unprepare one first"
	for _, tt := range []struct {
		name    string
		classes []ClassLevel
		spells  []Spell
		cmd     string
		caps    map[string]int
		want    string // the refusal, or "" for none
	}{
		{"under the cap", []ClassLevel{wizard}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true, Class: "wizard"}, {Slug: "b", Name: "B", Level: 1, Class: "wizard"}}, "spell prepare b", caps, ""},
		{"at the cap", []ClassLevel{wizard}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true, Class: "wizard"}, {Slug: "c", Name: "C", Level: 2, Prepared: true, Class: "wizard"}, {Slug: "b", Name: "B", Level: 1, Class: "wizard"}}, "spell prepare b", caps, full},
		{"prepared already", []ClassLevel{wizard}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true, Class: "wizard"}, {Slug: "c", Name: "C", Level: 2, Prepared: true, Class: "wizard"}}, "spell prepare a", caps, ""},
		{"feat and lineage spells are not counted", []ClassLevel{wizard}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true, Feat: "Magic Initiate (Wizard)"}, {Slug: "c", Name: "C", Level: 1, Prepared: true, Ability: "cha"}, {Slug: "b", Name: "B", Level: 1, Class: "wizard"}}, "spell prepare b", map[string]int{"wizard": 1}, ""},
		{"a feat spell is never capped", []ClassLevel{wizard}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true, Class: "wizard"}, {Slug: "b", Name: "B", Level: 1, Feat: "Magic Initiate (Wizard)"}}, "spell prepare b", map[string]int{"wizard": 1}, ""},
		{"untagged spells of a single class count", []ClassLevel{wizard}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true}, {Slug: "c", Name: "C", Level: 1, Prepared: true}, {Slug: "b", Name: "B", Level: 1}}, "spell prepare b", caps, full},
		{"another class's spells are not counted", []ClassLevel{wizard, cleric}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true, Class: "cleric"}, {Slug: "c", Name: "C", Level: 1, Prepared: true, Class: "cleric"}, {Slug: "b", Name: "B", Level: 1, Class: "wizard"}}, "spell prepare b", caps, ""},
		{"untagged spells of a multiclass are uncapped", []ClassLevel{wizard, cleric}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true}, {Slug: "c", Name: "C", Level: 1, Prepared: true}, {Slug: "b", Name: "B", Level: 1}}, "spell prepare b", caps, ""},
		{"no cap", []ClassLevel{wizard}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true, Class: "wizard"}, {Slug: "b", Name: "B", Level: 1, Class: "wizard"}}, "spell prepare b", nil, ""},
		{"unprepare over the cap", []ClassLevel{wizard}, []Spell{{Slug: "a", Name: "A", Level: 1, Prepared: true, Class: "wizard"}, {Slug: "c", Name: "C", Level: 1, Prepared: true, Class: "wizard"}, {Slug: "b", Name: "B", Level: 1, Prepared: true, Class: "wizard"}}, "spell unprepare b", caps, ""},
	} {
		c := fighter()
		c.Sheet.Classes, c.Sheet.Spells = tt.classes, tt.spells
		cmd, err := Parse(tt.cmd)
		if err != nil {
			t.Fatalf("%s: %v", tt.name, err)
		}
		_, err = cmd.Apply(&c, Env{PreparedCaps: tt.caps})
		got := ""
		if err != nil {
			got = err.Error()
		}
		if got != tt.want {
			t.Errorf("%s: %q, want %q", tt.name, got, tt.want)
		}
	}
}
