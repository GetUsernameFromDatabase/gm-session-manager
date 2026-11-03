package views

import (
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
)

// TestCombatPageEmpty: the empty state shows only when there are no
// encounters, not when the one asked for is missing.
func TestCombatPageEmpty(t *testing.T) {
	for _, tt := range []struct {
		name       string
		encounters []sqlc.Encounter
		want       bool
	}{
		{"no encounters", nil, true},
		{"asked-for encounter missing", []sqlc.Encounter{{ID: 1, Name: "Fight"}}, false},
	} {
		var b strings.Builder
		if err := CombatPage(Layout{}, CombatView{IsGM: true, CampaignID: 1, Encounters: tt.encounters}).Render(t.Context(), &b); err != nil {
			t.Fatal(err)
		}
		if got := strings.Contains(b.String(), "No encounters yet"); got != tt.want {
			t.Errorf("%s: empty state %v, want %v", tt.name, got, tt.want)
		}
	}
}

// TestTurnNamesFirstTurn: Previous stays empty until someone has acted.
func TestTurnNamesFirstTurn(t *testing.T) {
	fight := func(round int64, active int, out ...int) EncounterView {
		e := EncounterView{Encounter: sqlc.Encounter{Round: round}}
		for i, name := range []string{"A", "B", "C"} {
			e.Combatants = append(e.Combatants, CombatantView{Name: name, Active: i == active})
		}
		for _, i := range out {
			e.Combatants[i].Out = true
		}
		return e
	}
	for _, tt := range []struct {
		name string
		e    EncounterView
		prev string
	}{
		{"round 1, first turn", fight(1, 0), "—"},
		{"round 1, first in the fight after one out", fight(1, 1, 0), "—"},
		{"round 1, second turn", fight(1, 1), "A"},
		{"round 2 wraps", fight(2, 0), "C"},
	} {
		if prev, _, _ := tt.e.turnNames(); prev != tt.prev {
			t.Errorf("%s: prev %q, want %q", tt.name, prev, tt.prev)
		}
	}
}
