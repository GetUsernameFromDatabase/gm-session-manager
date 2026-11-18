package views

import (
	"database/sql"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/route"
)

// TestChipRemovesOwnEffects: a player sees a remove x on the effects they
// applied to an ally, in the chip and in its "+N" popover, and on no other.
func TestChipRemovesOwnEffects(t *testing.T) {
	applied := sql.NullInt64{Int64: 7, Valid: true}
	conds := []sqlc.Condition{
		{ID: 11, Slug: "prone"},
		{ID: 12, Slug: "custom", Note: "first", AppliedBy: applied},
		{ID: 13, Slug: "custom", Note: "overflow", AppliedBy: applied},
		{ID: 14, Slug: "poisoned"},
	}
	html := renderHTML(t, combatantChip(EncounterView{}, CombatantView{ID: 1, Name: "Ally", Kind: KindPC, CharacterID: 2, Conditions: conds}))
	// The add button reads as an add slot after the conditions, not one more of them.
	// The button's base class is border-transparent: the dash needs its own colour.
	if !strings.Contains(html, "+ Effect") || !strings.Contains(html, "border-dashed") || !strings.Contains(html, "border-muted-foreground/50") {
		t.Error("the ally's add-effect button is not a dashed \"+ Effect\" pill")
	}
	// As tall as the condition pills on phones too, not the sheet button's 32px.
	if strings.Contains(html, "max-sm:h-8") {
		t.Error("the ally's add-effect pill grows to 32px on phones")
	}
	for _, c := range conds {
		has := strings.Contains(html, `"`+route.Path(route.ConditionDelete, c.ID)+`"`)
		if has != c.AppliedBy.Valid {
			t.Errorf("condition %d: remove x shown = %v, want %v", c.ID, has, c.AppliedBy.Valid)
		}
	}
}

// TestCloseDialogOnlyOnOwnRequest: a field inside a dialog form that loads
// itself (the effect form's fields) must not close the dialog.
func TestCloseDialogOnlyOnOwnRequest(t *testing.T) {
	if js := closeDialogAfter("").Call; !strings.Contains(js, "event.detail.elt === this") {
		t.Errorf("closes on any request inside the form: %s", js)
	}
}

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
