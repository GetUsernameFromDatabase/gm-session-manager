package web

import (
	"slices"
	"testing"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
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

// TestUpgradeRepeatedFeature: a sheet holding three copies of a feature
// listed at several levels, each with its +10, keeps one whose bonus follows
// the class table column, and gets the profs its features grant.
func TestUpgradeRepeatedFeature(t *testing.T) {
	e := newTestEnv(t)
	if _, err := compendium.Import(t.Context(), e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	s := dnd5e.NewSheet()
	s.EffectsVersion = 3
	s.Classes = []dnd5e.ClassLevel{{Slug: "wizard", Name: "Wizard", Level: 12, HitDie: 6, Caster: dnd5e.CasterFull}}
	mod := dnd5e.CustomMod{Target: dnd5e.ModSpeed, Value: 10, Note: "Fast Feet", Source: "Fast Feet"}
	for _, lvl := range []int{9, 10, 11} {
		s.Features = append(s.Features, dnd5e.Feature{Name: "Fast Feet", Source: "Wizard", Level: lvl})
		s.Modifiers = append(s.Modifiers, mod)
	}
	s.Features = append(s.Features, dnd5e.Feature{Name: "Herb Lore", Source: "Wizard", Level: 1, Description: "You gain proficiency with the herbalism kit."})
	e.setSheet(t, s)
	for run, want := range []int{1, 0} {
		if n, err := UpgradeSheets(t.Context(), e.app.Store); err != nil || n != want {
			t.Fatalf("run %d: changed %d, err %v", run, n, err)
		}
	}
	got := e.sheet(t)
	if len(got.Features) != 2 || got.Features[0].Level != 9 || len(got.Modifiers) != 1 || got.Modifiers[0].Value != 15 {
		t.Errorf("features %+v, mods %+v", got.Features, got.Modifiers)
	}
	if !slices.Equal(got.OtherProfLabels(), []string{"herbalism kit (Herb Lore)"}) {
		t.Errorf("other profs %q", got.OtherProfLabels())
	}
}

// TestUpgradeThirdCaster: a subclass chosen before third casters, unpinned,
// brings its slots, and weapons added before mastery get theirs; a second run
// changes nothing.
func TestUpgradeThirdCaster(t *testing.T) {
	e := newTestEnv(t)
	seedMastery(t, e)
	if err := e.q().UpsertCompendiumEntry(t.Context(), testKnight); err != nil {
		t.Fatal(err)
	}
	s := dnd5e.NewSheet()
	s.EffectsVersion = 2
	s.Classes = []dnd5e.ClassLevel{{Slug: "testfighter", Name: "Testfighter", Subclass: "Test Knight", Level: 4, HitDie: 10, Caster: dnd5e.CasterNone}}
	s.Weapons = []dnd5e.Weapon{{Name: "+1 Longsword", Ability: "str", Dice: "1d8"}, {Name: "Stick", Ability: "str", Dice: "1d4"}}
	e.setSheet(t, s)
	for run, want := range []int{1, 0} {
		if n, err := UpgradeSheets(t.Context(), e.app.Store); err != nil || n != want {
			t.Fatalf("run %d: changed %d, err %v", run, n, err)
		}
	}
	got := e.sheet(t)
	if got.Classes[0].Caster != dnd5e.CasterThird || got.Slots[0].Max != 3 || got.Weapons[0].Mastery != "Sap" || got.Weapons[1].Mastery != "" {
		t.Errorf("classes %+v, slots %+v, weapons %+v", got.Classes, got.Slots[0], got.Weapons)
	}
}

// TestCatchUpResistances: an old sheet gets its species' fixed resistances
// and those of the lineage it picked; one with resistances keeps them.
func TestCatchUpResistances(t *testing.T) {
	d, err := compendium.Decode[compendium.TextData](enum.KindSpecies, `{"desc":"","parts":[
		{"name":"Draconic Ancestry","desc":"Choose the kind of dragon.\n\n| Dragon | Damage Type |\n|---|---|\n| Red | Fire |\n| White | Cold |"},
		{"name":"Damage Resistance","desc":"You have Resistance to the damage type determined by your Draconic Ancestry trait."},
		{"name":"Fiendish Legacy","desc":"Choose a legacy.\n\n| Legacy | Level 1 |\n|---|---|\n| Abyssal | You have Resistance to Poison damage. |\n| Infernal | You have Resistance to Fire damage. |"},
		{"name":"Tough Hide","desc":"You have Resistance to Acid damage."}]}`)
	if err != nil {
		t.Fatal(err)
	}
	parts := d.Parts
	picked := func(names ...string) []dnd5e.Feature {
		var fs []dnd5e.Feature
		for _, n := range names {
			fs = append(fs, dnd5e.Feature{Name: n, Source: "Testborn"})
		}
		return fs
	}
	for _, tt := range []struct {
		name     string
		features []dnd5e.Feature
		want     []string
	}{
		{"no picks", nil, []string{"acid"}},
		{"ancestry", picked("Draconic Ancestry (White)"), []string{"cold", "acid"}},
		{"legacy", picked("Fiendish Legacy (Abyssal)"), []string{"poison", "acid"}},
		{"other species' feature", []dnd5e.Feature{{Name: "Draconic Ancestry (Red)", Source: "Elf"}}, []string{"acid"}},
	} {
		s := dnd5e.Sheet{Species: "Testborn", Features: tt.features}
		if !catchUpResistances(&s, parts) || !slices.Equal(s.Resistances, tt.want) {
			t.Errorf("%s: %q, want %q", tt.name, s.Resistances, tt.want)
		}
	}
	if s := (dnd5e.Sheet{Species: "Testborn"}); catchUpResistances(&s, nil) || len(s.Resistances) != 0 {
		t.Errorf("unknown species: %q", s.Resistances)
	}
}
