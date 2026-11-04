package views

import (
	"net/url"
	"testing"
)

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
