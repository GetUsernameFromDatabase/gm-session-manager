package dnd5e

import "testing"

// TestSpellbook reads a spellbook's size from the class's features.
func TestSpellbook(t *testing.T) {
	wizard := ClassInfo{Features: []ClassFeature{{Name: "Spellcasting", Description: "The book contains the level 1+ spells you know. It starts with six level 1 Wizard spells of your choice.\n\nWhenever you gain a Wizard level after 1, add two Wizard spells of your choice to your spellbook."}}}
	lock := ClassInfo{Features: []ClassFeature{{Name: "Pact Magic", Description: "You know two Warlock cantrips of your choice."}}}
	for _, tt := range []struct {
		name  string
		info  ClassInfo
		level int
		want  int
	}{
		{"wizard 1", wizard, 1, 6},
		{"wizard 2", wizard, 2, 8},
		{"wizard 5", wizard, 5, 14},
		{"no spellbook", lock, 5, 0},
	} {
		if got := Spellbook(tt.info, tt.level); got != tt.want {
			t.Errorf("%s: %d, want %d", tt.name, got, tt.want)
		}
	}
}
