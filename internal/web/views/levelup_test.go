package views

import "testing"

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
