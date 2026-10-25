package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/perm"
)

// TestStepTurn checks next and previous turn pass over combatants out of the
// fight and wrap rounds, never before round 1.
func TestStepTurn(t *testing.T) {
	alive := sqlc.ListCombatantsRow{Name: "Ana", Hp: 5, MaxHp: 5}
	down := sqlc.ListCombatantsRow{Name: "Goblin", Hp: 0, MaxHp: 7}
	pcDown := sqlc.ListCombatantsRow{Name: "Bo", Hp: 0, MaxHp: 9, CharacterID: nullID(1), SheetJson: sql.NullString{String: "{}", Valid: true}}
	pcDead := sqlc.ListCombatantsRow{Name: "Cy", CharacterID: nullID(2), SheetJson: sql.NullString{String: `{"death_saves":{"failures":3}}`, Valid: true}}
	tests := []struct {
		name        string
		rows        []sqlc.ListCombatantsRow
		from        int
		round       int64
		dir         int
		want        int
		wantRound   int64
		wantSkipped []string
		wantErr     bool
	}{
		{"next", []sqlc.ListCombatantsRow{alive, alive}, 0, 1, 1, 1, 1, nil, false},
		{"next wraps", []sqlc.ListCombatantsRow{alive, alive}, 1, 1, 1, 0, 2, nil, false},
		{"next skips a downed monster", []sqlc.ListCombatantsRow{alive, down, alive}, 0, 1, 1, 2, 1, []string{"Goblin"}, false},
		{"a PC at 0 HP keeps their turn", []sqlc.ListCombatantsRow{alive, pcDown}, 0, 1, 1, 1, 1, nil, false},
		{"a dead PC is skipped", []sqlc.ListCombatantsRow{alive, pcDead, alive}, 0, 1, 1, 2, 1, []string{"Cy"}, false},
		{"prev wraps to the last round", []sqlc.ListCombatantsRow{alive, down, alive}, 0, 2, -1, 2, 1, nil, false},
		{"prev skips back", []sqlc.ListCombatantsRow{alive, down, alive}, 2, 1, -1, 0, 1, []string{"Goblin"}, false},
		{"prev stops at the first turn", []sqlc.ListCombatantsRow{alive, alive}, 0, 1, -1, 0, 0, nil, true},
		{"everyone down", []sqlc.ListCombatantsRow{down, down}, 0, 1, 1, 0, 0, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i, round, skipped, err := stepTurn(tt.rows, tt.from, tt.round, tt.dir)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, want error %v", err, tt.wantErr)
			}
			if err == nil && (i != tt.want || round != tt.wantRound || !slices.Equal(skipped, tt.wantSkipped)) {
				t.Errorf("got %d round %d skipped %v, want %d round %d skipped %v", i, round, skipped, tt.want, tt.wantRound, tt.wantSkipped)
			}
		})
	}
}

// TestPlayerCombatantLink checks players get a combatant's page only when
// they may see it: their own sheet, a revealed NPC, a monster with shown HP.
func TestPlayerCombatantLink(t *testing.T) {
	player := perm.Player(1)
	tests := []struct {
		name   string
		row    sqlc.ListCombatantsRow
		showHP bool
		want   string
	}{
		{"own character", sqlc.ListCombatantsRow{CharacterID: nullID(1)}, true, "/characters/1"},
		{"another character", sqlc.ListCombatantsRow{CharacterID: nullID(2)}, true, ""},
		{"revealed NPC", sqlc.ListCombatantsRow{EntityID: nullID(3), EntityRevealed: sql.NullBool{Bool: true, Valid: true}}, false, "/entities/3"},
		{"hidden NPC", sqlc.ListCombatantsRow{EntityID: nullID(3), EntityRevealed: sql.NullBool{Valid: true}}, true, ""},
		{"statted NPC never links its monster", sqlc.ListCombatantsRow{EntityID: nullID(3), MonsterRef: nullID(4)}, true, ""},
		{"monster with shown HP", sqlc.ListCombatantsRow{MonsterRef: nullID(4)}, true, "/compendium/entries/4"},
		{"monster with hidden HP", sqlc.ListCombatantsRow{MonsterRef: nullID(4)}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playerCombatantLink(player, tt.row, tt.showHP); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// TestRollInitiative checks a roll lands on d20 plus the bonus (none here).
func TestRollInitiative(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	cb, err := e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: enc.ID, Name: "Ogre", Hp: 59, MaxHp: 59})
	if err != nil {
		t.Fatal(err)
	}
	gm := e.token(t, auth.RoleAdmin)
	for range 20 {
		if w := e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/initiative", cb.ID), gm, url.Values{"roll": {"1"}}, true); w.Code != http.StatusOK {
			t.Fatalf("roll: %d %s", w.Code, w.Body)
		}
		got, _ := e.q().GetCombatant(ctx, cb.ID)
		if !got.Initiative.Valid || got.Initiative.Int64 < 1 || got.Initiative.Int64 > 20 {
			t.Fatalf("initiative %v outside 1..20", got.Initiative)
		}
	}
}

// TestPlayerCombatView checks players get their own card with the HP form
// and everyone else as chips without initiative, which the GM still sets,
// and see who is down even when their HP is hidden.
func TestPlayerCombatView(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	for _, p := range []sqlc.CreateCombatantParams{
		{EncounterID: enc.ID, CharacterID: nullID(e.characterID), Name: "pc"},
		{EncounterID: enc.ID, Name: "Ogre", Hp: 59, MaxHp: 59},
		{EncounterID: enc.ID, Name: "Goblin", Hp: 0, MaxHp: 7},
	} {
		if _, err := e.q().CreateCombatant(ctx, p); err != nil {
			t.Fatal(err)
		}
	}
	cbs, _ := e.q().ListCombatants(ctx, enc.ID)
	for _, cb := range cbs {
		if cb.Name == "Ogre" {
			if err := e.q().SetCombatantMarker(ctx, sqlc.SetCombatantMarkerParams{ID: cb.ID, Marker: "mini 5"}); err != nil {
				t.Fatal(err)
			}
			if err := e.q().SetCombatantInitiative(ctx, sqlc.SetCombatantInitiativeParams{ID: cb.ID, Initiative: nullID(12)}); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err := e.q().UpdateEncounterState(ctx, sqlc.UpdateEncounterStateParams{ID: enc.ID, Status: "active", Round: 1}); err != nil {
		t.Fatal(err)
	}
	combat := fmt.Sprintf("/campaigns/%d/combat", e.campaignID)
	player := e.do(t, http.MethodGet, combat, e.token(t, auth.RolePlayer), nil, false).Body.String()
	for _, want := range []string{`name="hp_amount"`, "Ogre", "Monsters", "mini 5", `data-down="true"`} {
		if !strings.Contains(player, want) {
			t.Errorf("player view misses %q", want)
		}
	}
	for _, leak := range []string{`aria-label="Initiative"`, `aria-label="Table marker"`, "59/59", "0/7"} {
		if strings.Contains(player, leak) {
			t.Errorf("player view shows %q", leak)
		}
	}
	gm := e.do(t, http.MethodGet, combat, e.token(t, auth.RoleAdmin), nil, false).Body.String()
	if !strings.Contains(gm, `aria-label="Initiative"`) {
		t.Error("GM view lost the initiative field")
	}
	// Only the pc and Goblin are unset; the Ogre's is set, so it has no roll.
	if n := strings.Count(gm, `aria-label="Roll initiative`); n != 2 {
		t.Errorf("%d roll buttons, want 2", n)
	}
}
