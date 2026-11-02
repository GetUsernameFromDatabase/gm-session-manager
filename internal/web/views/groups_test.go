package views

import (
	"slices"
	"testing"
)

// TestPlayerGroups checks players see others grouped by kind and sorted by
// name, never in initiative order.
func TestPlayerGroups(t *testing.T) {
	e := EncounterView{Combatants: []CombatantView{
		{Name: "Bandit 10", Kind: KindMonster},
		{Name: "Me", Kind: KindPC, IsOwn: true},
		{Name: "bandit 2", Kind: KindMonster},
		{Name: "Zed", Kind: KindPC},
		{Name: "Bandit", Kind: KindMonster},
		{Name: "Ann", Kind: KindPC},
	}}
	got := map[string][]string{}
	var kinds []string
	for _, g := range e.playerGroups() {
		kinds = append(kinds, g.Kind)
		for _, c := range g.Cards {
			got[g.Kind] = append(got[g.Kind], c.Name)
		}
	}
	tests := []struct {
		kind string
		want []string
	}{
		{KindPC, []string{"Ann", "Zed"}},
		{KindNPC, nil},
		{KindMonster, []string{"Bandit", "bandit 2", "Bandit 10"}},
	}
	for _, tt := range tests {
		if !slices.Equal(got[tt.kind], tt.want) {
			t.Errorf("%s: got %v, want %v", tt.kind, got[tt.kind], tt.want)
		}
	}
	if !slices.Equal(kinds, []string{KindPC, KindMonster}) {
		t.Errorf("groups %v, want pc then monster with no empty npc group", kinds)
	}
}
