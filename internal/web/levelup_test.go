package web

import (
	"fmt"
	"maps"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/web/views"
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
		if w := e.do(t, http.MethodPost, base, player, wizardLevelUp(t, e, url.Values{"class": {"wizard"}}), false); w.Code != http.StatusSeeOther {
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

// TestLevelUpTwice: the same form posted twice (a double click, another
// device) takes one level, and the second post is refused.
func TestLevelUpTwice(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	level := func() int {
		ch, _ := e.q().GetCharacter(ctx, e.characterID)
		s, _ := dnd5e.ParseSheet(ch.SheetJson)
		return s.Level()
	}
	start := level()
	form := wizardLevelUp(t, e, url.Values{"class": {"wizard"}, "level": {fmt.Sprint(start + 1)}})
	if w := e.do(t, http.MethodPost, base, player, form, false); w.Code != http.StatusSeeOther {
		t.Fatalf("level up: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodPost, base, player, form, false); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("second post: %d, want 422", w.Code)
	}
	if got := level(); got != start+1 {
		t.Errorf("level %d, want %d", got, start+1)
	}
}

// TestLevelUpInvocations: each level asks for the invocations its column
// adds, only from the options whose prerequisite (level, another invocation,
// a damaging cantrip) holds on the result, Eldritch Spear asks for its
// cantrip, and a later level may replace one.
func TestLevelUpInvocations(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	// A cantrip that deals no damage, so ticking it does not meet Eldritch Spear's prerequisite.
	spark := sqlc.UpsertCompendiumEntryParams{Kind: "spell", Slug: "test-spark", Name: "Test Spark", SourceKey: "test", ContentHash: "x", DataJson: `{"level":0,"classes":["testlock"],"desc":"A harmless light."}`}
	for _, entry := range []sqlc.UpsertCompendiumEntryParams{testlock, testBlast, spark} {
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
		{url.Values{"class": {"testlock"}, pick: {"Armor of Shadows"}}, http.StatusUnprocessableEntity, "@" + pick + ": choose 2"},                                                         // taken
		{url.Values{"class": {"testlock"}, pick: {"Eldritch Spear"}, "hp_roll": {"3"}}, http.StatusUnprocessableEntity, "@" + pick + ": choose 2"},                                         // level 2 adds two
		{url.Values{"class": {"testlock"}, pick: {"Eldritch Spear", "Agonizing Blast"}, spear: {"Test Blast"}}, http.StatusUnprocessableEntity, "choose 1 cantrip: you ticked 0 cantrips"}, // Test Blast not ticked
		{url.Values{"class": {"testlock"}, pick: {"Eldritch Spear", "Agonizing Blast"}, spear: {"Test Blast"}, "spell": {"test-spark"}}, http.StatusUnprocessableEntity, "Eldritch Spear needs Level 2+ Warlock, a Testlock Cantrip"},
		{url.Values{"class": {"testlock"}, pick: {"Eldritch Spear", "Agonizing Blast"}, spear: {"Mage Hand"}, "spell": {"test-blast"}}, http.StatusUnprocessableEntity, "@" + spear + ": choose one"},
		// Shadow Step needs Armor of Shadows, which the same level-up replaces.
		{url.Values{"class": {"testlock"}, pick: {"Shadow Step", "Agonizing Blast"}, replace: {"Armor of Shadows"}, "replace_eldritch-invocation_armor-of-shadows": {"Agonizing Blast"}, "spell": {"test-blast"}}, http.StatusUnprocessableEntity, "Shadow Step needs"},
		{url.Values{"class": {"testlock"}, pick: {"Shadow Step", "Eldritch Spear"}, spear: {"Test Blast"}, "spell": {"test-blast"}}, http.StatusSeeOther, ""},
		// Shadow Step needs Armor of Shadows, so it cannot be replaced.
		{url.Values{"class": {"testlock"}, replace: {"Armor of Shadows"}, "replace_eldritch-invocation_armor-of-shadows": {"Agonizing Blast"}}, http.StatusUnprocessableEntity, "@" + replace + ": choose one"},
		{url.Values{"class": {"testlock"}, replace: {"Shadow Step"}, "replace_eldritch-invocation_shadow-step": {"Agonizing Blast"}}, http.StatusSeeOther, ""},
	} {
		w := e.do(t, http.MethodPost, base, player, step.form, false)
		if w.Code != step.code || !hasError(w.Body.String(), step.err) {
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
// free cast per Long Rest, and the class list does not offer it then.
func TestLevelUpLineage(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	for _, entry := range []sqlc.UpsertCompendiumEntryParams{
		testlock,
		{Kind: "species", Slug: "testelf", Name: "Testelf", SourceKey: "test", ContentHash: "x", DataJson: `{"desc":"","parts":[
			{"name":"Lineage","desc":"Choose a lineage.\n\n|Lineage|Level 1|Level 3|\n|---|---|---|\n|Drow|You know the Dancing Lights cantrip.|Faerie Fire|"}]}`},
		{Kind: "spell", Slug: "faerie-fire", Name: "Faerie Fire", SourceKey: "test", ContentHash: "x", DataJson: `{"level":1,"classes":["testlock","testmage"]}`},
		{Kind: "spell", Slug: "test-ward", Name: "Test Ward", SourceKey: "test", ContentHash: "x", DataJson: `{"level":1,"classes":["testmage"]}`},
		{Kind: "spell", Slug: "test-hex", Name: "Test Hex", SourceKey: "test", ContentHash: "x", DataJson: `{"level":1,"classes":["testmage"]}`},
	} {
		if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
			t.Fatal(err)
		}
	}
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	sh, _ := dnd5e.ParseSheet(ch.SheetJson)
	sh.Species, sh.SpeciesAbility = "Testelf", "wis"
	sh.Features = append(sh.Features, dnd5e.Feature{Name: "Lineage (Drow)", Source: "Testelf"})
	mi, change := "Magic Initiate (Testmage)", "change_magic-initiate-testmage"
	sh.Feats = append(sh.Feats, mi)
	sh.Spells = append(sh.Spells, dnd5e.Spell{Slug: "test-ward", Name: "Test Ward", Level: 1, Prepared: true, Ability: "wis", Feat: mi})
	raw, _ := sh.JSON()
	if err := e.q().UpdateCharacterSheet(ctx, sqlc.UpdateCharacterSheetParams{ID: e.characterID, SheetJson: raw}); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	// Each level's new invocations: one, two, none.
	for i, picks := range [][]string{{"Armor of Shadows"}, {"Agonizing Blast", "Shadow Step"}, nil} {
		if i == 2 {
			ch, _ := e.q().GetCharacter(ctx, e.characterID)
			v, err := e.app.levelUpView(ctx, ch, "testlock", false)
			if err != nil || slices.ContainsFunc(v.Spells, func(o views.Option) bool { return o.Slug == "faerie-fire" }) {
				t.Errorf("class list offers the lineage's Faerie Fire: %v %+v", err, v.Spells)
			}
			form := url.Values{"class": {"testlock"}, "spell": {"faerie-fire"}}
			// No checklist on the page: the refusal shows at the top.
			if w := e.do(t, http.MethodPost, base, player, form, false); w.Code != http.StatusUnprocessableEntity || errorAt(w.Body.String(), views.FieldSpell) ||
				!strings.Contains(w.Body.String(), "choose spells from the list") || len(lostPicks(w.Body.String(), form)) > 0 {
				t.Errorf("Faerie Fire as a class pick: %d, message not at the top or picks lost", w.Code)
			}
			// Nor does the feat's Spell Change, which would drop the pick.
			j := slices.IndexFunc(v.Groups, func(g views.ChoiceGroup) bool { return g.Key == change })
			if j < 0 {
				t.Fatalf("no Spell Change: %+v", v.Groups)
			}
			with := v.Groups[j].Sub["Test Ward"][0].Options
			has := func(n string) bool {
				return slices.ContainsFunc(with, func(o compendium.NameDesc) bool { return o.Name == n })
			}
			if has("Faerie Fire") || !has("Test Hex") {
				t.Errorf("Spell Change offers %+v", with)
			}
			form = url.Values{"class": {"testlock"}, change: {"Test Ward"}, change + "_test-ward": {"Faerie Fire"}}
			w := e.do(t, http.MethodPost, base, player, form, false)
			if body := w.Body.String(); w.Code != http.StatusUnprocessableEntity || !errorAt(body, change+"_test-ward") || strings.Count(body, `aria-invalid="true"`) != 1 || strings.Count(body, `aria-describedby="error-`+change+`_test-ward"`) != 1 || len(lostPicks(body, form)) > 0 {
				t.Errorf("Faerie Fire as a Spell Change: %d, lost %v", w.Code, lostPicks(body, form))
			}
		}
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
		msg   string // a refusal's message
		check func(sh dnd5e.Sheet) bool
	}{
		{"Grappler raises Strength", 15, url.Values{"feat": {"grappler"}, "feat_grappler_raise": {"Strength"}}, http.StatusSeeOther, "",
			func(sh dnd5e.Sheet) bool { return sh.Scores["str"] == 16 && slices.Contains(sh.Feats, "Grappler") }},
		{"Grappler at the maximum", 20, url.Values{"feat": {"grappler"}, "feat_grappler_raise": {"Strength"}}, http.StatusUnprocessableEntity, "", nil},
		{"Grappler without Strength or Dexterity 13", 8, url.Values{"feat": {"grappler"}, "feat_grappler_raise": {"Strength"}}, http.StatusUnprocessableEntity,
			"Grappler needs Level 4+, Strength or Dexterity 13+", nil},
		{"an Epic Boon before level 19", 15, url.Values{"feat": {"boon-of-fate"}, "feat_boon-of-fate_raise": {"Charisma"}}, http.StatusUnprocessableEntity,
			"Boon of Fate needs Level 19+", nil},
		{"an Epic Boon before level 19, its choice left out", 15, url.Values{"feat": {"boon-of-fate"}}, http.StatusUnprocessableEntity,
			"Boon of Fate needs Level 19+", nil},
		{"a Fighting Style feat without the feature", 15, url.Values{"feat": {"archery"}}, http.StatusUnprocessableEntity,
			"Archery needs Fighting Style Feature", nil},
		{"a second Magic Initiate takes another list", 15, url.Values{"feat": {"magic-initiate"}, mi + "_list": {"Cleric"}, mi + "_cleric_cantrips": {"Guidance"}, mi + "_cleric_spell": {"Bless"}, mi + "_ability": {"Wisdom"}}, http.StatusSeeOther, "",
			func(sh dnd5e.Sheet) bool {
				return slices.Contains(sh.Feats, "Magic Initiate (Cleric)") && slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool {
					return sp.Slug == "bless" && sp.Ability == "wis" && sp.Feat == "Magic Initiate (Cleric)"
				})
			}},
		{"not the list taken already", 15, url.Values{"feat": {"magic-initiate"}, mi + "_list": {"Wizard"}}, http.StatusUnprocessableEntity, "", nil},
		{"a feat the compendium lacks", 15, url.Values{"feat": {"nope"}}, http.StatusUnprocessableEntity, "choose a feat from the list", nil},
		{"a feat and an increase", 15, url.Values{"feat": {"grappler"}, "feat_grappler_raise": {"Strength"}, "asi1": {"str"}}, http.StatusUnprocessableEntity,
			"choose a feat or an ability score increase, not both", nil},
		{"an increase above 20", 20, url.Values{"asi1": {"str"}}, http.StatusUnprocessableEntity, "Strength cannot go above 20", nil},
		{"Spell Change", 15, url.Values{change: {"Magic Missile"}, change + "_magic-missile": {"Shield"}}, http.StatusSeeOther, "",
			func(sh dnd5e.Sheet) bool {
				return !slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == "magic-missile" }) &&
					slices.ContainsFunc(sh.Spells, func(sp dnd5e.Spell) bool {
						return sp.Slug == "shield" && sp.Prepared && sp.Ability == "int" && sp.Feat != ""
					}) &&
					slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == "Shield (free cast)" && f.Uses.Fixed == 1 })
			}},
		{"Spell Change to a spell ticked in the class list", 15, url.Values{"spell": {"shield"}, change: {"Magic Missile"}, change + "_magic-missile": {"Shield"}}, http.StatusUnprocessableEntity,
			"Shield comes from your species or feat already", nil},
		{"a new Magic Initiate's spell ticked in the class list", 15, url.Values{"spell": {"detect-magic"}, "feat": {"magic-initiate"}, mi + "_list": {"Cleric"}, mi + "_cleric_cantrips": {"Guidance"}, mi + "_cleric_spell": {"Detect Magic"}, mi + "_ability": {"Wisdom"}}, http.StatusUnprocessableEntity,
			"Detect Magic comes from your species or feat already", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := newTestEnv(t)
			ctx := t.Context()
			loadFeatFixtures(t, e)
			// On both lists, so a Cleric Magic Initiate can pick a Wizard class spell.
			if err := e.q().UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: "spell", Slug: "detect-magic", Name: "Detect Magic", SourceKey: "test", ContentHash: "x", DataJson: `{"level":1,"classes":["cleric","wizard"]}`}); err != nil {
				t.Fatal(err)
			}
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
				if w := e.do(t, http.MethodPost, base, player, wizardLevelUp(t, e, url.Values{"class": {"wizard"}}), false); w.Code != http.StatusSeeOther {
					t.Fatalf("level up: %d %s", w.Code, w.Body)
				}
			}
			// The Ability Score Improvement feat is the form's own selects; feats
			// whose prerequisite the sheet does not meet are left out.
			body := e.do(t, http.MethodGet, base+"?class=wizard", player, nil, false).Body.String()
			// Every select's list is wide by default: the class select's, in its own form.
			beforeClass, _, _ := strings.Cut(body, `name="class"`)
			classList := beforeClass[strings.LastIndex(beforeClass, "<form"):]
			for _, c := range []string{"min-w-(--anchor-width)", "max-w-[calc(100vw-2rem)]"} {
				if !strings.Contains(classList, c) {
					t.Errorf("class select list lacks %s: %s", c, classList)
				}
			}
			offered := func(slug string) bool { return strings.Contains(body, `value="`+slug+`"`) }
			if offered("grappler") != (tc.str >= 13) || offered("ability-score-improvement") || offered("archery") || offered("boon-of-fate") {
				t.Errorf("level 4 feats: grappler %v, Ability Score Improvement %v, archery %v, boon %v",
					offered("grappler"), offered("ability-score-improvement"), offered("archery"), offered("boon-of-fate"))
			}
			tc.form.Set("class", "wizard")
			// A class spell clash shows at the class spells; a feat's prerequisite at the top.
			atSpells := tc.form.Has("spell")
			w := e.do(t, http.MethodPost, base, player, wizardLevelUp(t, e, tc.form), false)
			if w.Code != tc.code {
				t.Fatalf("level 4: %d %s", w.Code, w.Body)
			}
			body = w.Body.String()
			if tc.code == http.StatusUnprocessableEntity && !strings.Contains(body, "Gain level") {
				t.Errorf("refusal did not re-render the form: %s", body)
			}
			kept := tc.form
			if (strings.Contains(tc.msg, " needs ") || tc.msg == "choose a feat from the list") && tc.form.Get("feat") != "" {
				// A feat that is not offered falls back to the select's first item.
				kept = maps.Clone(tc.form)
				kept.Del("feat")
				for _, tag := range reInput.FindAllString(body, -1) {
					if strings.Contains(tag, ` name="feat"`) && strings.Contains(tag, `value="`+tc.form.Get("feat")+`"`) {
						t.Errorf("refused page kept the unoffered feat: %s", tag)
					}
				}
			}
			if lost := lostPicks(body, kept); tc.code == http.StatusUnprocessableEntity && len(lost) > 0 {
				t.Errorf("refused page lost %v", lost)
			}
			// Every other refusal here is about the increase or feat pick.
			at := views.FieldASI
			if atSpells {
				at = views.FieldSpell
			}
			if tc.msg != "" && (!strings.Contains(body, tc.msg) || !errorAt(body, at)) {
				t.Errorf("no %q at %s: %s", tc.msg, at, body)
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

// TestLevelUpFeatWithoutASI: a feat posted at a level with no ASI draws no feat choices, so its refusal shows at the top.
func TestLevelUpFeatWithoutASI(t *testing.T) {
	mi := "feat_magic-initiate"
	e := newTestEnv(t)
	ctx := t.Context()
	loadFeatFixtures(t, e)
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	sh, _ := dnd5e.ParseSheet(ch.SheetJson)
	sh.Feats = []string{"Magic Initiate (Wizard)"}
	sh.Spells = []dnd5e.Spell{{Slug: "light", Name: "Light", Ability: "int", Feat: "Magic Initiate (Wizard)"}}
	raw, _ := sh.JSON()
	if err := e.q().UpdateCharacterSheet(ctx, sqlc.UpdateCharacterSheetParams{ID: e.characterID, SheetJson: raw}); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	form := url.Values{"class": {"wizard"}, "feat": {"magic-initiate"}}
	w := e.do(t, http.MethodPost, base, player, wizardLevelUp(t, e, form), false)
	body := w.Body.String()
	if w.Code != http.StatusUnprocessableEntity || errorAt(body, mi+"_list") || !strings.Contains(body, ": choose one") {
		t.Errorf("feat choice without an ASI: %d, message not at the top", w.Code)
	}
}

// TestLevelUpEpicBoon checks the Epic Boon level (the test SRD has none): the
// feat is offered, and its own maximum of 30 beats the usual 20.
func TestLevelUpEpicBoon(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	loadFeatFixtures(t, e)
	boonClass := sqlc.UpsertCompendiumEntryParams{Kind: "class", Slug: "boonling", Name: "Boonling", SourceKey: "test", ContentHash: "x",
		DataJson: `{"caster":"none","hit_die":8,"features":[{"name":"Epic Boon","desc":"","levels":[19]}]}`}
	if err := e.q().UpsertCompendiumEntry(ctx, boonClass); err != nil {
		t.Fatal(err)
	}
	ch, _ := e.q().GetCharacter(ctx, e.characterID)
	sh, _ := dnd5e.ParseSheet(ch.SheetJson)
	sh.Classes = []dnd5e.ClassLevel{{Slug: "boonling", Name: "Boonling", Level: 18, HitDie: 8, Caster: "none"}}
	sh.Scores["cha"] = 20
	raw, _ := sh.JSON()
	if err := e.q().UpdateCharacterSheet(ctx, sqlc.UpdateCharacterSheetParams{ID: e.characterID, SheetJson: raw}); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	if body := e.do(t, http.MethodGet, base+"?class=boonling", player, nil, false).Body.String(); !strings.Contains(body, `value="boon-of-fate"`) {
		t.Fatalf("Boon of Fate not offered: %s", body)
	}
	form := url.Values{"class": {"boonling"}, "feat": {"boon-of-fate"}, "feat_boon-of-fate_raise": {"Charisma"}}
	if w := e.do(t, http.MethodPost, base, player, form, false); w.Code != http.StatusSeeOther {
		t.Fatalf("level up: %d %s", w.Code, w.Body)
	}
	ch, _ = e.q().GetCharacter(ctx, e.characterID)
	if sh, err := dnd5e.ParseSheet(ch.SheetJson); err != nil || sh.Level() != 19 || !slices.Contains(sh.Feats, "Boon of Fate") || sh.Scores["cha"] != 21 {
		t.Errorf("sheet: %v %s", err, ch.SheetJson)
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
		if w := e.do(t, http.MethodPost, base, player, wizardLevelUp(t, e, url.Values{"class": {"wizard"}}), false); w.Code != http.StatusSeeOther {
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
