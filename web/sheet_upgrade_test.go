package web

import (
	"testing"

	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
)

// TestUpgradeSheets: a feature saved before the parser read its uses gets
// them once; uses already set stay, and a second run changes nothing.
func TestUpgradeSheets(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	s := dnd5e.NewSheet()
	s.Classes = []dnd5e.ClassLevel{{Slug: "wizard", Name: "Wizard", Level: 2, HitDie: 6, Caster: dnd5e.CasterFull}}
	const desc = "You can use it a number of times equal to your proficiency bonus."
	s.Features = []dnd5e.Feature{
		{Name: "Old", Source: "Wizard", Description: desc},
		{Name: "Set", Source: "Wizard", Description: desc, Uses: dnd5e.Uses{Fixed: 1, Used: 1}},
	}
	raw, err := s.JSON()
	if err != nil {
		t.Fatal(err)
	}
	if err := e.q().UpdateCharacterSheet(ctx, sqlc.UpdateCharacterSheetParams{ID: e.characterID, SheetJson: raw}); err != nil {
		t.Fatal(err)
	}

	for run, want := range []int{1, 0} {
		if n, err := UpgradeSheets(ctx, e.app.Store); err != nil || n != want {
			t.Fatalf("run %d: changed %d, err %v", run, n, err)
		}
	}
	row, err := e.q().GetCharacter(ctx, e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	got, err := dnd5e.ParseSheet(row.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Features[0].Uses.Prof || got.Features[1].Uses.Prof || got.Features[1].Uses.Fixed != 1 || got.Features[1].Uses.Used != 1 || got.EffectsVersion != dnd5e.EffectsVersion {
		t.Errorf("features %+v, version %d", got.Features, got.EffectsVersion)
	}
}
