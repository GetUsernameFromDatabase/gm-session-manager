package web

import (
	"fmt"
	"net/http"
	"testing"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// benchSpells is the size of the benchmark caster's spell list; the first
// benchPinned are pinned to the outranked source.
const (
	benchSpells = 15
	benchPinned = 2
)

func (e testEnv) spell(tb testing.TB, source, slug string, level int) {
	tb.Helper()
	data := fmt.Sprintf(`{"level":%d,"school":"evocation","classes":["wizard"],"desc":"%s from %s"}`, level, slug, source)
	if err := e.q().UpsertCompendiumEntry(tb.Context(), sqlc.UpsertCompendiumEntryParams{Kind: enum.KindSpell, Slug: slug, Name: slug, SourceKey: source, DataJson: data, ContentHash: source + slug}); err != nil {
		tb.Fatal(err)
	}
}

// BenchmarkSheetState is one sheet publish for one device: a level-5 wizard
// with 15 spells, two pinned to a non-priority source.
func BenchmarkSheetState(b *testing.B) {
	f := newLiveFixture(b)
	if _, err := compendium.Import(b.Context(), f.app.Store, "../compendium/testdata"); err != nil {
		b.Fatal(err)
	}
	f.source(b, oldSrc, "2014")
	f.source(b, newSrc, "2024")
	sheet := dnd5e.Sheet{Classes: []dnd5e.ClassLevel{{Slug: "wizard", Name: "Wizard", Level: 5, HitDie: 6, Caster: dnd5e.CasterFull}}}
	for i := range benchSpells {
		slug, level := fmt.Sprintf("bench-spell-%d", i), i%4
		f.spell(b, newSrc, slug, level)
		sp := dnd5e.Spell{Slug: slug, Name: slug, Level: level, Prepared: level > 0}
		if i < benchPinned {
			f.spell(b, oldSrc, slug, level)
			sp.Source = oldSrc
		}
		sheet.Spells = append(sheet.Spells, sp)
	}
	f.allowEntrySources(b)
	f.setSheet(b, sheet)
	client := hub.NewClient(f.campaignID, perm.Player(f.characterID), 0, "bench")
	view := views.LiveKey(views.LiveCharacter, f.characterID)
	b.ReportAllocs()
	for b.Loop() {
		if _, ok, err := f.app.State(b.Context(), client, view); err != nil || !ok {
			b.Fatal(ok, err)
		}
	}
}

// BenchmarkEncounterState is one tracker publish for the GM: 4 PCs and 8
// monsters.
func BenchmarkEncounterState(b *testing.B) {
	f := newLiveFixture(b) // Alice and the Ogre are already in the fight
	ctx := b.Context()
	for i := range 3 {
		ch, err := f.q().CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: f.campaignID, Name: fmt.Sprintf("PC %d", i), System: "dnd5e"})
		if err != nil {
			b.Fatal(err)
		}
		if _, err := f.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: f.encounter, CharacterID: nullID(ch.ID), Name: ch.Name}); err != nil {
			b.Fatal(err)
		}
	}
	for i := range 7 {
		if _, err := f.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: f.encounter, Name: fmt.Sprintf("Goblin %d", i), Hp: 7, MaxHp: 7, Ac: 15}); err != nil {
			b.Fatal(err)
		}
	}
	client := hub.NewClient(f.campaignID, perm.GM, 0, "bench")
	view := views.LiveKey(views.LiveEncounter, f.encounter)
	b.ReportAllocs()
	for b.Loop() {
		if _, ok, err := f.app.State(ctx, client, view); err != nil || !ok {
			b.Fatal(ok, err)
		}
	}
}

// BenchmarkCompendiumList is the GM's spell list page.
func BenchmarkCompendiumList(b *testing.B) {
	e := newTestEnv(b)
	if _, err := compendium.Import(b.Context(), e.app.Store, "../compendium/testdata"); err != nil {
		b.Fatal(err)
	}
	gm := e.token(b, enum.RoleAdmin)
	path := "/compendium?" + views.KindKey + "=" + string(enum.KindSpell)
	b.ReportAllocs()
	for b.Loop() {
		if w := e.do(b, http.MethodGet, path, gm, nil, false); w.Code != http.StatusOK {
			b.Fatal(w.Code)
		}
	}
}
