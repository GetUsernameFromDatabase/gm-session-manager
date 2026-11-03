package web

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
)

// TestRevertLevelUp checks level-ups revert newest first, back to the sheet
// and max HP from before each, keeping current HP up to that max.
func TestRevertLevelUp(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	start, _ := e.q().GetCharacter(ctx, e.characterID)
	var after []int64
	for range 2 {
		if w := e.do(t, http.MethodPost, base, player, url.Values{"class": {"wizard"}}, false); w.Code != http.StatusSeeOther {
			t.Fatalf("level up: %d %s", w.Code, w.Body)
		}
		ch, _ := e.q().GetCharacter(ctx, e.characterID)
		after = append(after, ch.MaxHp)
	}
	if after[1] <= after[0] {
		t.Fatalf("max HP did not grow: %v", after)
	}
	for i, wantMax := range []int64{after[0], start.MaxHp} {
		if w := e.do(t, http.MethodPost, base+"/revert", player, nil, true); w.Code != http.StatusOK {
			t.Fatalf("revert %d: %d %s", i+1, w.Code, w.Body)
		}
		ch, _ := e.q().GetCharacter(ctx, e.characterID)
		if ch.MaxHp != wantMax || ch.Hp > ch.MaxHp {
			t.Errorf("revert %d: hp %d/%d, want max %d", i+1, ch.Hp, ch.MaxHp, wantMax)
		}
	}
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	startSheet, _ := dnd5e.ParseSheet(start.SheetJson)
	if want, _ := startSheet.JSON(); ch.SheetJson != want {
		t.Errorf("sheet not restored:\n got %s\nwant %s", ch.SheetJson, want)
	}
	if w := e.do(t, http.MethodPost, base+"/revert", player, nil, true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("revert with nothing left: %d, want 422", w.Code)
	}
}

// TestLevelUpInvocations: each level asks for the invocations its column
// adds, only from the options whose prerequisite (level, another invocation,
// a damaging cantrip) holds on the result, Eldritch Spear asks for its
// cantrip, and a later level may replace one.
func TestLevelUpInvocations(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	for _, entry := range []sqlc.UpsertCompendiumEntryParams{testlock, testBlast} {
		if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
			t.Fatal(err)
		}
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	if w := e.do(t, http.MethodGet, base+"?class=testlock", player, nil, false); !strings.Contains(w.Body.String(), "Eldritch Invocations") || strings.Contains(w.Body.String(), "Agonizing Blast") {
		t.Errorf("level 1 form lacks the pick or offers a level 2 option")
	}
	const pick, spear, replace = "class_eldritch-invocation", "class_eldritch-invocation_eldritch-spear", "replace_eldritch-invocation"
	// Level 2 grows the Cantrips column, so Eldritch Spear is offered for
	// Test Blast, learned in the same level-up.
	for _, step := range []struct {
		form url.Values
		code int
		err  string // in the refused page
	}{
		{url.Values{"class": {"testlock"}, pick: {"Armor of Shadows"}}, http.StatusSeeOther, ""},
		{url.Values{"class": {"testlock"}, pick: {"Armor of Shadows"}}, http.StatusUnprocessableEntity, "listed eldritch invocations"},                                                                       // taken
		{url.Values{"class": {"testlock"}, pick: {"Eldritch Spear"}, "hp_roll": {"3"}}, http.StatusUnprocessableEntity, "choose 2 of the listed eldritch invocations"},                                       // level 2 adds two
		{url.Values{"class": {"testlock"}, pick: {"Eldritch Spear", "Agonizing Blast"}, spear: {"Test Blast"}}, http.StatusUnprocessableEntity, "Eldritch Spear needs Level 2+ Warlock, a Testlock Cantrip"}, // Test Blast not ticked
		{url.Values{"class": {"testlock"}, pick: {"Eldritch Spear", "Agonizing Blast"}, spear: {"Mage Hand"}, "spell": {"test-blast"}}, http.StatusUnprocessableEntity, "listed eldritch spear: cantrip"},
		// Shadow Step needs Armor of Shadows, which the same level-up replaces.
		{url.Values{"class": {"testlock"}, pick: {"Shadow Step", "Agonizing Blast"}, replace: {"Armor of Shadows"}, "replace_eldritch-invocation_armor-of-shadows": {"Agonizing Blast"}}, http.StatusUnprocessableEntity, "Shadow Step needs"},
		{url.Values{"class": {"testlock"}, pick: {"Shadow Step", "Eldritch Spear"}, spear: {"Test Blast"}, "spell": {"test-blast"}}, http.StatusSeeOther, ""},
		// Shadow Step needs Armor of Shadows, so it cannot be replaced.
		{url.Values{"class": {"testlock"}, replace: {"Armor of Shadows"}, "replace_eldritch-invocation_armor-of-shadows": {"Agonizing Blast"}}, http.StatusUnprocessableEntity, "listed replace one eldritch invocation"},
		{url.Values{"class": {"testlock"}, replace: {"Shadow Step"}, "replace_eldritch-invocation_shadow-step": {"Agonizing Blast"}}, http.StatusSeeOther, ""},
	} {
		w := e.do(t, http.MethodPost, base, player, step.form, false)
		if w.Code != step.code || !strings.Contains(w.Body.String(), step.err) {
			t.Fatalf("%v: %d, want %d with %q", step.form, w.Code, step.code, step.err)
		}
		if lost := lostPicks(w.Body.String(), step.form); step.code == http.StatusUnprocessableEntity && len(lost) > 0 {
			t.Errorf("%v: refused page lost %v", step.form, lost)
		}
	}
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	sh, _ := dnd5e.ParseSheet(ch.SheetJson)
	for _, name := range []string{"Eldritch Invocation: Armor of Shadows", "Eldritch Invocation: Eldritch Spear (Test Blast)", "Eldritch Invocation: Agonizing Blast"} {
		if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == name }) {
			t.Errorf("missing %s: %s", name, ch.SheetJson)
		}
	}
	if slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == "Eldritch Invocation: Shadow Step" }) {
		t.Errorf("Shadow Step not replaced: %s", ch.SheetJson)
	}
	if !slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == "test-blast" && sp.Class == "testlock" }) {
		t.Errorf("Test Blast without its class: %+v", sh.Spells)
	}
	// Learned from another class's list, Test Blast is no Testlock cantrip.
	const prereq = "a Testlock Cantrip That Deals Damage"
	for class, want := range map[string]bool{"": true, "testlock": true, "sorcerer": false} {
		cs, err := e.app.cantripData(ctx, []dnd5e.Spell{{Slug: "test-blast", Name: "Test Blast", Class: class}}, ch.CampaignID)
		if err != nil {
			t.Fatal(err)
		}
		if got := meetsPrerequisite(prereq, dnd5e.Sheet{}, 2, cs); got != want {
			t.Errorf("taught by %q: %v, want %v", class, got, want)
		}
	}
	if sh.ClassSummary() != "Testlock 3" {
		t.Errorf("classes %q", sh.ClassSummary())
	}
}

// TestLevelUpLineage checks character level 3 adds the lineage spell and its
// free cast per Long Rest.
func TestLevelUpLineage(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	for _, entry := range []sqlc.UpsertCompendiumEntryParams{
		testlock,
		{Kind: "species", Slug: "testelf", Name: "Testelf", SourceKey: "test", ContentHash: "x", DataJson: `{"desc":"","parts":[
			{"name":"Lineage","desc":"Choose a lineage.\n\n|Lineage|Level 1|Level 3|\n|---|---|---|\n|Drow|You know the Dancing Lights cantrip.|Faerie Fire|"}]}`},
		{Kind: "spell", Slug: "faerie-fire", Name: "Faerie Fire", SourceKey: "test", ContentHash: "x", DataJson: `{"level":1}`},
	} {
		if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
			t.Fatal(err)
		}
	}
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	sh, _ := dnd5e.ParseSheet(ch.SheetJson)
	sh.Species, sh.SpeciesAbility = "Testelf", "wis"
	sh.Features = append(sh.Features, dnd5e.Feature{Name: "Lineage (Drow)", Source: "Testelf"})
	raw, _ := sh.JSON()
	if err := e.q().UpdateCharacterSheet(ctx, sqlc.UpdateCharacterSheetParams{ID: e.characterID, SheetJson: raw}); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	// Each level's new invocations: one, two, none.
	for _, picks := range [][]string{{"Armor of Shadows"}, {"Agonizing Blast", "Shadow Step"}, nil} {
		if w := e.do(t, http.MethodPost, base, player, url.Values{"class": {"testlock"}, "class_eldritch-invocation": picks}, false); w.Code != http.StatusSeeOther {
			t.Fatalf("level up: %d %s", w.Code, w.Body)
		}
	}
	ch, _ = e.q().GetCharacter(ctx, e.characterID)
	sh, _ = dnd5e.ParseSheet(ch.SheetJson)
	if !slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == "faerie-fire" && sp.Prepared && sp.Ability == "wis" }) {
		t.Errorf("no prepared Faerie Fire with wis: %+v", sh.Spells)
	}
	i := slices.IndexFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == "Faerie Fire (free cast)" })
	if i < 0 || sh.Features[i].Uses.Fixed != 1 || sh.Features[i].Uses.Reset != "long" {
		t.Errorf("free cast feature: %+v", sh.Features)
	}
}

// TestLevelUpFeats: a feat taken at Wizard 4 asks for its own choices, and
// Magic Initiate's Spell Change swaps one of its spells on any level-up.
func TestLevelUpFeats(t *testing.T) {
	mi := "feat_magic-initiate"
	change := "change_magic-initiate-wizard"
	for _, tc := range []struct {
		name  string
		str   int
		form  url.Values
		code  int
		check func(sh dnd5e.Sheet) bool
	}{
		{"Grappler raises Strength", 15, url.Values{"feat": {"grappler"}, "feat_grappler_raise": {"Strength"}}, http.StatusSeeOther,
			func(sh dnd5e.Sheet) bool { return sh.Scores["str"] == 16 && slices.Contains(sh.Feats, "Grappler") }},
		{"Grappler at the maximum", 20, url.Values{"feat": {"grappler"}, "feat_grappler_raise": {"Strength"}}, http.StatusUnprocessableEntity, nil},
		{"Grappler without Strength or Dexterity 13", 8, url.Values{"feat": {"grappler"}, "feat_grappler_raise": {"Strength"}}, http.StatusUnprocessableEntity, nil},
		{"an Epic Boon before level 19", 15, url.Values{"feat": {"boon-of-fate"}, "feat_boon-of-fate_raise": {"Charisma"}}, http.StatusUnprocessableEntity, nil},
		{"a Fighting Style feat without the feature", 15, url.Values{"feat": {"archery"}}, http.StatusUnprocessableEntity, nil},
		{"a second Magic Initiate takes another list", 15, url.Values{"feat": {"magic-initiate"}, mi + "_list": {"Cleric"}, mi + "_cleric_cantrips": {"Guidance"}, mi + "_cleric_spell": {"Bless"}, mi + "_ability": {"Wisdom"}}, http.StatusSeeOther,
			func(sh dnd5e.Sheet) bool {
				return slices.Contains(sh.Feats, "Magic Initiate (Cleric)") && slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool {
					return sp.Slug == "bless" && sp.Ability == "wis" && sp.Feat == "Magic Initiate (Cleric)"
				})
			}},
		{"not the list taken already", 15, url.Values{"feat": {"magic-initiate"}, mi + "_list": {"Wizard"}}, http.StatusUnprocessableEntity, nil},
		{"Spell Change", 15, url.Values{change: {"Magic Missile"}, change + "_magic-missile": {"Shield"}}, http.StatusSeeOther,
			func(sh dnd5e.Sheet) bool {
				return !slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == "magic-missile" }) &&
					slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool {
						return sp.Slug == "shield" && sp.Prepared && sp.Ability == "int" && sp.Feat != ""
					}) &&
					slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == "Shield (free cast)" && f.Uses.Fixed == 1 })
			}},
		{"Spell Change to a spell ticked in the class list", 15, url.Values{"spell": {"shield"}, change: {"Magic Missile"}, change + "_magic-missile": {"Shield"}}, http.StatusSeeOther,
			func(sh dnd5e.Sheet) bool {
				shields := slices.DeleteFunc(slices.Clone(sh.Spells), func(sp dnd5e.Spell) bool { return sp.Slug != "shield" })
				return len(shields) == 1 && slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == "magic-missile" })
			}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := newTestEnv(t)
			ctx := t.Context()
			loadFeatFixtures(t, e)
			ch, _ := e.q().GetCharacter(ctx, e.characterID)
			sh, _ := dnd5e.ParseSheet(ch.SheetJson)
			sh.Scores["str"], sh.Scores["dex"] = tc.str, 10 // Grappler needs either at 13
			sh.Feats = []string{"Magic Initiate (Wizard)"}
			sh.Spells = []dnd5e.Spell{
				{Slug: "light", Name: "Light", Ability: "int", Feat: "Magic Initiate (Wizard)"},
				{Slug: "magic-missile", Name: "Magic Missile", Level: 1, Prepared: true, Ability: "int", Feat: "Magic Initiate (Wizard)"},
			}
			sh.Features = append(sh.Features, freeCast(sh.Spells[1], "Magic Initiate (Wizard)"))
			raw, _ := sh.JSON()
			if err := e.q().UpdateCharacterSheet(ctx, sqlc.UpdateCharacterSheetParams{ID: e.characterID, SheetJson: raw}); err != nil {
				t.Fatal(err)
			}
			player := e.token(t, auth.RolePlayer)
			base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
			for range 3 {
				if w := e.do(t, http.MethodPost, base, player, url.Values{"class": {"wizard"}}, false); w.Code != http.StatusSeeOther {
					t.Fatalf("level up: %d %s", w.Code, w.Body)
				}
			}
			// The Ability Score Improvement feat is the form's own selects; feats
			// whose prerequisite the sheet does not meet are left out.
			body := e.do(t, http.MethodGet, base+"?class=wizard", player, nil, false).Body.String()
			offered := func(slug string) bool { return strings.Contains(body, `value="`+slug+`"`) }
			if offered("grappler") != (tc.str >= 13) || offered("ability-score-improvement") || offered("archery") || offered("boon-of-fate") {
				t.Errorf("level 4 feats: grappler %v, Ability Score Improvement %v, archery %v, boon %v",
					offered("grappler"), offered("ability-score-improvement"), offered("archery"), offered("boon-of-fate"))
			}
			tc.form.Set("class", "wizard")
			w := e.do(t, http.MethodPost, base, player, tc.form, false)
			if w.Code != tc.code {
				t.Fatalf("level 4: %d %s", w.Code, w.Body)
			}
			if lost := lostPicks(w.Body.String(), tc.form); tc.code == http.StatusUnprocessableEntity && len(lost) > 0 {
				t.Errorf("refused page lost %v", lost)
			}
			if tc.check == nil {
				return
			}
			ch, _ = e.q().GetCharacter(ctx, e.characterID)
			if sh, err := dnd5e.ParseSheet(ch.SheetJson); err != nil || !tc.check(sh) || sh.Level() != 4 {
				t.Errorf("sheet: %v %s", err, ch.SheetJson)
			}
		})
	}
}

// TestRevertLevelUpInterleaved checks level-ups and reverts may alternate:
// each revert returns to the sheet from just before the newest level-up.
func TestRevertLevelUpInterleaved(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	start, _ := e.q().GetCharacter(ctx, e.characterID)
	startSheet, _ := dnd5e.ParseSheet(start.SheetJson)
	wantStart, _ := startSheet.JSON()
	maxHP := func() int64 {
		ch, _ := e.q().GetCharacter(ctx, e.characterID)
		return ch.MaxHp
	}
	levelUp := func() {
		t.Helper()
		if w := e.do(t, http.MethodPost, base, player, url.Values{"class": {"wizard"}}, false); w.Code != http.StatusSeeOther {
			t.Fatalf("level up: %d %s", w.Code, w.Body)
		}
	}
	revert := func() {
		t.Helper()
		if w := e.do(t, http.MethodPost, base+"/revert", player, nil, true); w.Code != http.StatusOK {
			t.Fatalf("revert: %d %s", w.Code, w.Body)
		}
	}

	levelUp()
	revert()
	if n, _ := e.q().CountLevelUps(ctx, e.characterID); n != 0 || maxHP() != start.MaxHp {
		t.Fatalf("after up, revert: %d level-ups left, max HP %d, want 0 and %d", n, maxHP(), start.MaxHp)
	}

	levelUp()
	once := maxHP()
	levelUp()
	revert()
	if got := maxHP(); got != once {
		t.Errorf("after up, up, revert: max HP %d, want %d", got, once)
	}

	levelUp()
	revert()
	revert()
	if ch, _ := e.q().GetCharacter(ctx, e.characterID); ch.SheetJson != wantStart || ch.MaxHp != start.MaxHp {
		t.Errorf("not back at the start:\n got %s (max %d)\nwant %s (max %d)", ch.SheetJson, ch.MaxHp, wantStart, start.MaxHp)
	}
	if w := e.do(t, http.MethodPost, base+"/revert", player, nil, true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("revert with nothing left: %d, want 422", w.Code)
	}
}
