package views

import (
	"net/url"
	"testing"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/dnd5e"
)

// TestDrawsField: a refusal's field counts as drawn only where the templates
// render an at for it.
func TestDrawsField(t *testing.T) {
	spells := []Option{{Slug: "light", Name: "Light"}}
	groups := []ChoiceGroup{{Key: "trait_lineage", Options: []compendium.NameDesc{{Name: "Drow"}, {Name: "High Elf"}},
		Sub: map[string][]ChoiceGroup{"Drow": {{Key: "trait_lineage_ability"}}, "Gone": {{Key: "orphan"}}}}}
	feats := map[string][]ChoiceGroup{"grappler": {{Key: "feat_grappler_ability"}}}
	asi, plain := &dnd5e.LevelUpPlan{ASI: true}, &dnd5e.LevelUpPlan{}
	for _, tt := range []struct {
		name, key string
		wizard    WizardView
		levelUp   LevelUpView
		want      bool
	}{
		{"group", "trait_lineage", WizardView{Choices: &CreationChoices{Groups: groups}}, LevelUpView{Plan: plain, Groups: groups}, true},
		{"sub-group", "trait_lineage_ability", WizardView{Choices: &CreationChoices{Groups: groups}}, LevelUpView{Plan: plain, Groups: groups}, true},
		{"sub-group of no option", "orphan", WizardView{Choices: &CreationChoices{Groups: groups}}, LevelUpView{Plan: plain, Groups: groups}, false},
		{"no groups", "trait_lineage", WizardView{}, LevelUpView{Groups: groups}, false},
		{"spell with a checklist", "spell", WizardView{Choices: &CreationChoices{Spells: spells}}, LevelUpView{Plan: plain, Spells: spells}, true},
		{"spell without", "spell", WizardView{Choices: &CreationChoices{}}, LevelUpView{Plan: plain}, false},
		{"unknown key", "nope", WizardView{Choices: &CreationChoices{Groups: groups}}, LevelUpView{Plan: plain, Groups: groups}, false},
		{"empty key", "", WizardView{Choices: &CreationChoices{Groups: groups}}, LevelUpView{Plan: plain, Groups: groups}, false},
	} {
		if got := tt.wizard.DrawsField(tt.key); got != tt.want {
			t.Errorf("wizard %s: %v, want %v", tt.name, got, tt.want)
		}
		if got := tt.levelUp.DrawsField(tt.key); got != tt.want {
			t.Errorf("level-up %s: %v, want %v", tt.name, got, tt.want)
		}
	}
	// Only the creation class step draws skills, and only while it asks for some.
	for _, tt := range []struct {
		name string
		v    WizardView
		want bool
	}{
		{"class step asking", WizardView{Step: "class", Choice: &WizardChoice{SkillCount: 2}}, true},
		{"class step not asking", WizardView{Step: "class", Choice: &WizardChoice{}}, false},
		{"other step", WizardView{Step: "species", Choice: &WizardChoice{SkillCount: 2}}, false},
		{"nothing picked", WizardView{Step: "class"}, false},
	} {
		if got := tt.v.DrawsField("skill"); got != tt.want {
			t.Errorf("skill, %s: %v, want %v", tt.name, got, tt.want)
		}
	}
	// A feat's choices are drawn only with an ASI and the feat on offer.
	for _, tt := range []struct {
		name string
		v    LevelUpView
		want bool
	}{
		{"ASI and feat offered", LevelUpView{Plan: asi, Feats: []Option{{Slug: "grappler"}}, FeatChoices: feats}, true},
		{"no ASI", LevelUpView{Plan: plain, Feats: []Option{{Slug: "grappler"}}, FeatChoices: feats}, false},
		{"feat not offered", LevelUpView{Plan: asi, FeatChoices: feats}, false},
	} {
		if got := tt.v.DrawsField("feat_grappler_ability"); got != tt.want {
			t.Errorf("feat choice, %s: %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestSpellCounts: cantrip and spell totals only, never slot columns.
func TestSpellCounts(t *testing.T) {
	for _, tt := range []struct {
		name string
		cols map[string]string
		want string
	}{
		{"warlock", map[string]string{"Cantrips": "2", "Eldritch Invocations": "3", "Prepared Spells": "3", "Slot Level": "1st", "Spell Slots": "2"}, "Cantrips 2 · Prepared Spells 3"},
		{"spells known", map[string]string{"Cantrips Known": "4", "Spells Known": "5"}, "Cantrips Known 4 · Spells Known 5"},
		{"no casting", map[string]string{"Rages": "2"}, ""},
		{"empty", nil, ""},
	} {
		if got := spellCounts(tt.cols); got != tt.want {
			t.Errorf("%s: %q, want %q", tt.name, got, tt.want)
		}
	}
}

// TestTaughtPicked: a spell a lineage on the form teaches is hidden only while
// that lineage option is picked.
func TestTaughtPicked(t *testing.T) {
	rules := []string{"trait_elven-lineage=High Elf"}
	for _, tt := range []struct {
		name string
		form url.Values
		want bool
	}{
		{"the lineage's option", url.Values{"trait_elven-lineage": {"High Elf"}}, true},
		{"another option", url.Values{"trait_elven-lineage": {"Drow"}}, false},
		{"nothing picked", url.Values{}, false},
	} {
		if got := taughtPicked(rules, tt.form); got != tt.want {
			t.Errorf("%s: got %v, want %v", tt.name, got, tt.want)
		}
	}
	if taughtPicked(nil, url.Values{"trait_elven-lineage": {"High Elf"}}) {
		t.Error("a spell no lineage teaches is hidden")
	}
	if got := taughtRules(nil); got != "[]" {
		t.Errorf("no rules: %s, want []", got)
	}
}
