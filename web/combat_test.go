package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
)

// TestConditionExpiresOnTurn checks a timed condition lasts through its
// rounds and ends as its creature's turn starts, not when the round turns.
func TestConditionExpiresOnTurn(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm := e.token(t, auth.RoleAdmin)
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	for _, name := range []string{"Ana", "Bo"} {
		ent, _ := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: name, FieldsJson: "{}"})
		if w := e.do(t, http.MethodPost, fmt.Sprintf("/encounters/%d/combatants", enc.ID), gm, url.Values{"kind": {"entity"}, "id": {fmt.Sprint(ent.ID)}}, true); w.Code != http.StatusOK {
			t.Fatalf("add %s: %d %s", name, w.Code, w.Body)
		}
	}
	rows, _ := e.q().ListCombatants(ctx, enc.ID)
	for i, r := range rows { // Ana goes first
		e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/initiative", r.ID), gm, url.Values{"initiative": {fmt.Sprint(20 - i)}}, true)
	}
	bo := rows[1].ID
	e.do(t, http.MethodPost, fmt.Sprintf("/encounters/%d/start", enc.ID), gm, nil, true)
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/condition", bo), gm, url.Values{"slug": {"prone"}, "rounds": {"1"}}, true); w.Code != http.StatusOK {
		t.Fatalf("condition: %d %s", w.Code, w.Body)
	}
	// Turns after start: Bo r1, Ana r2, Bo r2.
	for i, want := range []bool{true, true, false} {
		if w := e.do(t, http.MethodPost, fmt.Sprintf("/encounters/%d/next", enc.ID), gm, nil, true); w.Code != http.StatusOK {
			t.Fatalf("next: %d %s", w.Code, w.Body)
		}
		conds, _ := e.q().ListConditionsByEncounter(ctx, enc.ID)
		if got := len(conds) == 1; got != want {
			t.Fatalf("turn %d: prone present %v, want %v", i+1, got, want)
		}
	}
}

// TestCombatantNamesFollowSource checks renaming a character or NPC renames
// its combatants, while monsters keep their numbered names.
func TestCombatantNamesFollowSource(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	ent, _ := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Ana", FieldsJson: "{}"})
	for _, p := range []sqlc.CreateCombatantParams{
		{EncounterID: enc.ID, CharacterID: nullID(ch.ID), Name: ch.Name},
		{EncounterID: enc.ID, EntityID: nullID(ent.ID), Name: ent.Name},
		{EncounterID: enc.ID, Name: "Goblin 2"},
	} {
		if _, err := e.q().CreateCombatant(ctx, p); err != nil {
			t.Fatal(err)
		}
	}
	if err := e.q().UpdateCharacter(ctx, sqlc.UpdateCharacterParams{Name: "Renamed PC", Hp: ch.Hp, MaxHp: ch.MaxHp, TempHp: ch.TempHp, SheetJson: ch.SheetJson, ID: ch.ID}); err != nil {
		t.Fatal(err)
	}
	if err := e.q().UpdateEntity(ctx, sqlc.UpdateEntityParams{Name: "Renamed NPC", FieldsJson: "{}", ID: ent.ID}); err != nil {
		t.Fatal(err)
	}
	rows, _ := e.q().ListCombatants(ctx, enc.ID)
	got := map[string]bool{}
	for _, r := range rows {
		got[r.Name] = true
	}
	for _, want := range []string{"Renamed PC", "Renamed NPC", "Goblin 2"} {
		if !got[want] {
			t.Errorf("missing %q in %v", want, got)
		}
	}
}

// TestMonsterPortrait: a monster's compendium picture shows in the tracker
// for everyone, since every player may read the compendium.
func TestMonsterPortrait(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	const pic = "https://example.com/goblin.png"
	if err := e.q().UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: "monster", Slug: "goblin", Name: "Goblin", SourceKey: "test", DataJson: "{}"}); err != nil {
		t.Fatal(err)
	}
	m, _ := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin"})
	if _, err := e.q().SetCompendiumImage(ctx, sqlc.SetCompendiumImageParams{ImageUrl: sql.NullString{String: pic, Valid: true}, ID: m.ID}); err != nil {
		t.Fatal(err)
	}
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	for _, p := range []sqlc.CreateCombatantParams{
		{EncounterID: enc.ID, CharacterID: nullID(e.characterID), Name: "pc"},
		{EncounterID: enc.ID, MonsterRef: nullID(m.ID), Name: "Goblin"},
	} {
		if _, err := e.q().CreateCombatant(ctx, p); err != nil {
			t.Fatal(err)
		}
	}
	if err := e.q().UpdateEncounterState(ctx, sqlc.UpdateEncounterStateParams{ID: enc.ID, Status: "active", Round: 1}); err != nil {
		t.Fatal(err)
	}
	combat := fmt.Sprintf("/campaigns/%d/combat", e.campaignID)
	for _, role := range []auth.Role{auth.RoleAdmin, auth.RolePlayer} {
		if body := e.do(t, http.MethodGet, combat, e.token(t, role), nil, false).Body.String(); !strings.Contains(body, pic) {
			t.Errorf("%v: combat misses the monster picture", role)
		}
	}
}
