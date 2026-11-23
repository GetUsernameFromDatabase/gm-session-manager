package dnd5e

import (
	"slices"
	"testing"
)

// TestParseArmorTraining reads every armor training wording the sources use.
func TestParseArmorTraining(t *testing.T) {
	all := []string{ArmorLight, ArmorMedium, ArmorHeavy, ArmorShield}
	for _, tc := range []struct {
		text string
		want []string
	}{
		{"Light, Medium, and Heavy armor and Shields", all},
		{"All armor, shields", all},
		{"Light and medium armor, shields", []string{ArmorLight, ArmorMedium, ArmorShield}},
		{"Light armor", []string{ArmorLight}},
		{"Shields", []string{ArmorShield}},
		{"None", nil},
		{"Daggers, darts, slings, quarterstaffs, light crossbows", nil},
		{"Simple weapons, heavy crossbows, Medium armor", []string{ArmorMedium}},
		{"{@filter Heavy armor|items|type=Heavy Armor}", []string{ArmorHeavy}},
	} {
		if got := ParseArmorTraining(tc.text); !slices.Equal(got, tc.want) {
			t.Errorf("ParseArmorTraining(%q) = %q, want %q", tc.text, got, tc.want)
		}
	}
}

// TestArmorTrainingTextReadsBack writes every set of kinds and reads it back.
func TestArmorTrainingTextReadsBack(t *testing.T) {
	for mask := range 1 << len(ArmorKinds) {
		var kinds []string
		for i, k := range ArmorKinds {
			if mask&(1<<i) != 0 {
				kinds = append(kinds, k)
			}
		}
		text := ArmorTrainingText(kinds)
		if got := ParseArmorTraining(text); !slices.Equal(got, kinds) {
			t.Errorf("%q reads back as %q, want %q", text, got, kinds)
		}
	}
	if got := ArmorTrainingText([]string{ArmorLight, ArmorMedium, ArmorShield}); got != "Light and Medium armor and Shields" {
		t.Errorf("wording %q", got)
	}
}

// TestArmorGrants checks what feature and feat texts train, and that a
// feature offering options grants none of them.
func TestArmorGrants(t *testing.T) {
	for _, tc := range []struct {
		desc string
		want []string
	}{
		{"Increase your Strength or Dexterity score by 1. You gain training with Light armor and Shields.", []string{ArmorLight, ArmorShield}},
		{"You gain proficiency with heavy armor and smith's tools.", []string{ArmorHeavy}},
		{"While you aren't wearing Heavy armor, your speed increases by 10 feet.", nil},
		{"You have dedicated yourself to one of the following sacred roles of your choice.\n\n**Protector.** You gain proficiency with Martial weapons and training with Heavy armor.", nil},
	} {
		if got := ParseEffects("Test", tc.desc, nil, "").Armor; !slices.Equal(got, tc.want) {
			t.Errorf("%q trains %q, want %q", tc.desc, got, tc.want)
		}
	}
}

// TestArmorUntrained checks the warning reads the training, not weapon text.
func TestArmorUntrained(t *testing.T) {
	wizard := Sheet{OtherProfs: []string{"Daggers, darts, slings, quarterstaffs, light crossbows"}, Armor: &WornArmor{Name: "Leather Armor", Type: ArmorLight}, Shield: true}
	if got := wizard.ArmorUntrained(); !slices.Equal(got, []string{"light armor", "shields"}) {
		t.Errorf("wizard untrained %q", got)
	}
	knight := Sheet{ArmorTraining: []string{ArmorHeavy, ArmorShield}, Armor: &WornArmor{Name: "Plate Armor", Type: ArmorHeavy}, Shield: true}
	if got := knight.ArmorUntrained(); got != nil {
		t.Errorf("trained knight untrained %q", got)
	}
}

// TestArmorTrainingGrantsAndRemoval follows training from a feat, a
// multiclass level and a command, and its removal with its origin.
func TestArmorTrainingGrantsAndRemoval(t *testing.T) {
	c := fighter()
	c.Sheet.Scores[Wis] = 13
	AddFeat(&c.Sheet, Feature{Name: "Heavily Armored", Source: SourceFeat, Description: "You gain training with Heavy armor."})
	if err := c.LevelUp(LevelUpChoice{Class: ClassInfo{Slug: "cleric", Name: "Cleric", HitDie: 8}}); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(c.Sheet.ArmorTraining, []string{ArmorLight, ArmorMedium, ArmorHeavy, ArmorShield}) {
		t.Fatalf("training %q", c.Sheet.ArmorTraining)
	}
	if got := c.Sheet.ArmorTrainingLabels(); !slices.Equal(got, []string{"Light (Cleric)", "Medium (Cleric)", "Heavy (Heavily Armored)", "Shields (Cleric)"}) {
		t.Errorf("labels %q", got)
	}
	c.Sheet.RemoveProfsFrom("Heavily Armored")
	if c.Sheet.TrainedIn(ArmorHeavy) || !c.Sheet.TrainedIn(ArmorMedium) {
		t.Errorf("after removing the feat: %q", c.Sheet.ArmorTraining)
	}
	for _, line := range []string{"prof add armor Heavy", "prof remove armor shields"} {
		cmd, err := Parse(line)
		if err == nil {
			_, err = cmd.Apply(&c, Env{})
		}
		if err != nil {
			t.Fatalf("%s: %v", line, err)
		}
	}
	if !slices.Equal(c.Sheet.ArmorTraining, []string{ArmorLight, ArmorMedium, ArmorHeavy}) || c.Sheet.ProfSources["armor:shield"] != "" {
		t.Errorf("after commands: %q, sources %v", c.Sheet.ArmorTraining, c.Sheet.ProfSources)
	}
	if _, err := Parse("prof add armor plate"); err == nil {
		t.Error("plate is no armor training kind")
	}
}

// TestCatchUpArmorTraining moves training kept as text and gives a second
// class its multiclass training, leaving weapon text alone.
func TestCatchUpArmorTraining(t *testing.T) {
	s := NewSheet()
	s.Classes = []ClassLevel{{Slug: "fighter", Name: "Fighter", Level: 2}, {Slug: "rogue", Name: "Rogue", Level: 1}}
	s.OtherProfs = []string{"All armor, shields", "Simple weapons, martial weapons", "None", "Daggers, darts, slings, quarterstaffs, light crossbows"}
	s.ProfSources = map[string]string{"All armor, shields": "Fighter"}
	if !s.CatchUpArmorTraining() {
		t.Fatal("nothing changed")
	}
	if !slices.Equal(s.ArmorTraining, []string{ArmorLight, ArmorMedium, ArmorHeavy, ArmorShield}) {
		t.Errorf("training %q", s.ArmorTraining)
	}
	if !slices.Equal(s.OtherProfs, []string{"Simple weapons, martial weapons", "Daggers, darts, slings, quarterstaffs, light crossbows"}) {
		t.Errorf("other profs %q", s.OtherProfs)
	}
	if s.ProfSources["armor:heavy"] != "Fighter" || s.ProfSources["All armor, shields"] != "" {
		t.Errorf("sources %v", s.ProfSources)
	}
	if s.CatchUpArmorTraining() {
		t.Error("a second catch-up changed the sheet")
	}
}
