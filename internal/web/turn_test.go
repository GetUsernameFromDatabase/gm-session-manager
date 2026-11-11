package web

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/perm"
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
		{"monster with shown HP", sqlc.ListCombatantsRow{MonsterRef: nullID(4)}, true, "/campaigns/9/compendium/entries/4"},
		{"monster with hidden HP", sqlc.ListCombatantsRow{MonsterRef: nullID(4)}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := playerCombatantLink(player, tt.row, 9, tt.showHP); got != tt.want {
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
	gm := e.token(t, enum.RoleAdmin)
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

// TestRollAllInitiative checks the roll action fills only monsters and NPCs
// without an initiative, leaving PCs and set values alone.
func TestRollAllInitiative(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	npc, _ := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Mira", FieldsJson: "{}"})
	npc2, _ := e.q().CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: e.campaignID, Kind: "npc", Name: "Vex", FieldsJson: "{}"})
	// Statblock initiative bonuses by combatant name; the range check below must include them.
	bonus := map[string]int64{"Bandit": 3, "Vex": -1}
	refs := map[string]int64{}
	for name, b := range bonus {
		slug := "test-init-" + strings.ToLower(name)
		if err := e.q().UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: "monster", Slug: slug, Name: name, SourceKey: "test", ContentHash: "x",
			DataJson: fmt.Sprintf(`{"initiative":%d}`, b)}); err != nil {
			t.Fatal(err)
		}
		m, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: slug})
		if err != nil {
			t.Fatal(err)
		}
		refs[name] = m.ID
	}
	for _, p := range []sqlc.CreateCombatantParams{
		{Name: "Hero", CharacterID: nullID(e.characterID)},
		{Name: "Ogre", Hp: 59, MaxHp: 59},
		{Name: "Guard", Hp: 11, MaxHp: 11, Initiative: sql.NullInt64{Int64: 7, Valid: true}},
		{Name: "Mira", EntityID: nullID(npc.ID)},
		{Name: "Bandit", MonsterRef: nullID(refs["Bandit"])},
		{Name: "Vex", EntityID: nullID(npc2.ID), MonsterRef: nullID(refs["Vex"])},
	} {
		p.EncounterID = enc.ID
		if _, err := e.q().CreateCombatant(ctx, p); err != nil {
			t.Fatal(err)
		}
	}
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/encounters/%d/roll", enc.ID), e.token(t, enum.RoleAdmin), nil, true); w.Code != http.StatusOK {
		t.Fatalf("roll: %d %s", w.Code, w.Body)
	}
	rows, _ := e.q().ListCombatants(ctx, enc.ID)
	for _, r := range rows {
		switch lo := int64(1 + initiativeMod(r)); {
		case r.CharacterID.Valid:
			if r.Initiative.Valid {
				t.Errorf("%s: PC initiative %v, want unset", r.Name, r.Initiative)
			}
		case r.Name == "Guard":
			if r.Initiative.Int64 != 7 {
				t.Errorf("Guard initiative %v, want 7", r.Initiative)
			}
		default:
			if want, ok := bonus[r.Name]; ok && initiativeMod(r) != int(want) {
				t.Errorf("%s: initiativeMod %d, want %d", r.Name, initiativeMod(r), want)
			}
			if !r.Initiative.Valid || r.Initiative.Int64 < lo || r.Initiative.Int64 > lo+19 {
				t.Errorf("%s: initiative %v outside %d..%d", r.Name, r.Initiative, lo, lo+19)
			}
		}
	}
}

// TestCombatDefenses: the GM's damage follows a statblock's resistances,
// immunities, vulnerabilities and conditional defenses, and Petrified; a
// player damages their own character by type, and never sees another
// combatant's defenses.
func TestCombatDefenses(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	gm, player := e.token(t, enum.RoleAdmin), e.token(t, enum.RolePlayer)
	if err := e.q().UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: "monster", Slug: "testgoyle", Name: "Testgoyle", SourceKey: "test", ContentHash: "x",
		DataJson: `{"ac":15,"hp":60,"resistances":"cold; bludgeoning from nonmagical attacks","immunities":"poison; poisoned","vulnerabilities":"radiant"}`}); err != nil {
		t.Fatal(err)
	}
	m, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "testgoyle"})
	if err != nil {
		t.Fatal(err)
	}
	enc, _ := e.q().CreateEncounter(ctx, sqlc.CreateEncounterParams{CampaignID: e.campaignID, Name: "Ambush"})
	mon, err := e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: enc.ID, MonsterRef: nullID(m.ID), Name: "Testgoyle", Hp: 60, MaxHp: 60})
	if err != nil {
		t.Fatal(err)
	}
	if _, err := e.q().CreateCombatant(ctx, sqlc.CreateCombatantParams{EncounterID: enc.ID, CharacterID: nullID(e.characterID), Name: "pc"}); err != nil {
		t.Fatal(err)
	}
	if err := e.app.updateCharacter(ctx, e.characterID, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
		dc.Sheet.Resistances, dc.HP, dc.MaxHP = []string{"fire"}, 30, 30
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	hpURL := fmt.Sprintf("/combatants/%d/hp", mon.ID)
	hp := int64(60)
	for _, tt := range []struct {
		kind  string
		taken int64 // -1: refused
	}{
		{" ", 9}, {"cold", 4}, {"poison", 0}, {"radiant", 18},
		{"bludgeoning", 9}, {"bludgeoning@1", 4},
		{"bludgeoning@0", -1}, {"cold@1", -1}, {"pizza", -1},
	} {
		w := e.do(t, http.MethodPost, hpURL, gm, url.Values{"op": {"damage"}, "amount": {"9"}, "type": {tt.kind}}, true)
		if tt.taken < 0 {
			if w.Code == http.StatusOK {
				t.Errorf("%q accepted", tt.kind)
			}
			continue
		}
		hp -= tt.taken
		if cb, _ := e.q().GetCombatant(ctx, mon.ID); cb.Hp != hp {
			t.Errorf("%q: HP %d, want %d", tt.kind, cb.Hp, hp)
		}
	}
	e.do(t, http.MethodPost, fmt.Sprintf("/combatants/%d/condition", mon.ID), gm, url.Values{"slug": {"petrified"}}, true)
	e.do(t, http.MethodPost, hpURL, gm, url.Values{"op": {"damage"}, "amount": {"9"}}, true)
	if cb, _ := e.q().GetCombatant(ctx, mon.ID); cb.Hp != hp-4 {
		t.Errorf("petrified: HP %d, want %d", cb.Hp, hp-4)
	}

	cmd := fmt.Sprintf("/characters/%d/cmd", e.characterID)
	e.do(t, http.MethodPost, cmd, player, url.Values{"hp_op": {"damage"}, "hp_amount": {"9"}, "hp_type": {"fire"}}, true)
	if ch, _ := e.q().GetCharacter(ctx, e.characterID); ch.Hp != 26 {
		t.Errorf("player's fire damage: HP %d, want 26", ch.Hp)
	}
	if w := e.do(t, http.MethodPost, cmd, player, url.Values{"hp_op": {"damage"}, "hp_amount": {"9"}, "hp_type": {"pizza"}}, true); w.Code == http.StatusOK {
		t.Error("player's unknown damage type accepted")
	}

	if err := e.q().UpdateEncounterState(ctx, sqlc.UpdateEncounterStateParams{ID: enc.ID, Status: "active", Round: 1}); err != nil {
		t.Fatal(err)
	}
	combat := fmt.Sprintf("/campaigns/%d/combat", e.campaignID)
	view := e.do(t, http.MethodGet, combat, player, nil, false).Body.String()
	if !strings.Contains(view, `name="hp_type"`) || !strings.Contains(view, "Fire (resisted, half)") {
		t.Error("player's own row has no damage type")
	}
	for _, leak := range []string{"Cold (resisted", "immune, none", "nonmagical", "Petrified: all damage"} {
		if strings.Contains(view, leak) {
			t.Errorf("player view shows the monster's %q", leak)
		}
	}
	gmView := e.do(t, http.MethodGet, combat, gm, nil, false).Body.String()
	for _, want := range []string{"Cold (resisted, half)", "Poison (immune, none)", "Radiant (vulnerable, double)", "Bludgeoning from nonmagical attacks (resisted, half)", "Petrified: all damage is halved."} {
		if !strings.Contains(gmView, want) {
			t.Errorf("GM view misses %q", want)
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
	player := e.do(t, http.MethodGet, combat, e.token(t, enum.RolePlayer), nil, false).Body.String()
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
	gm := e.do(t, http.MethodGet, combat, e.token(t, enum.RoleAdmin), nil, false).Body.String()
	if !strings.Contains(gm, `aria-label="Initiative"`) {
		t.Error("GM view lost the initiative field")
	}
	// Only the pc and Goblin are unset; the Ogre's is set, so it has no roll.
	if n := strings.Count(gm, `aria-label="Roll initiative`); n != 2 {
		t.Errorf("%d roll buttons, want 2", n)
	}
}
