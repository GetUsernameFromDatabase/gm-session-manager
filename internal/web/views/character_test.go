package views

import (
	"testing"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/dnd5e"
)

// TestFreeCastKind: a lineage spell's free cast sorts by the spell's casting time.
func TestFreeCastKind(t *testing.T) {
	for _, tt := range []struct {
		name, casting, want, detail string
		noData                      bool
	}{
		{"reaction", "Reaction, which you take in response to taking damage", "reaction", "Reaction · Tiefling", false},
		{"action", "Action", "action", "Action · Tiefling", false},
		{"bonus", "Bonus Action", "bonus", "Bonus Action · Tiefling", false},
		{"long cast", "1 minute", "other", "Tiefling", false},
		{"spell not in compendium", "", "other", "Tiefling", true},
	} {
		s := SheetView{C: dnd5e.Character{Sheet: dnd5e.Sheet{
			Spells:   []dnd5e.Spell{{Slug: "hellish-rebuke", Name: "Hellish Rebuke", Level: 1}},
			Features: []dnd5e.Feature{{Name: "Hellish Rebuke (free cast)", Source: "Tiefling", Uses: dnd5e.Uses{Fixed: 1, Reset: "long"}}},
		}}}
		if !tt.noData {
			s.SpellData = map[string]*compendium.SpellData{"hellish-rebuke": {Level: 1, CastingTime: tt.casting}}
		}
		got := s.actions()["features"]
		if len(got) != 1 || got[0].Kind != tt.want || got[0].Detail != tt.detail {
			t.Errorf("%s: %+v, want kind %q detail %q", tt.name, got, tt.want, tt.detail)
		}
	}
}
