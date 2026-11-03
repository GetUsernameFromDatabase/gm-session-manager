package web

import (
	"database/sql"
	"errors"
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

// TestCreationWizard walks a new character through every step and checks
// what each one put on the sheet, and that steps adding to it run once.
func TestCreationWizard(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	for _, entry := range []sqlc.UpsertCompendiumEntryParams{
		{Kind: "species", Slug: "goliath", Name: "Goliath", DataJson: `{"desc":"","parts":[{"name":"Speed","desc":"35 feet","type":"speed"},{"name":"Powerful Build","desc":"You count as one size larger when carrying."},
			{"name":"Giant Ancestry","desc":"Choose one of the following benefits; Intelligence, Wisdom, or Charisma is your spellcasting ability for its spells:\n\n- **Cloud's Jaunt (Cloud Giant)**. You teleport.\n- **Fire's Burn (Fire Giant)**. You know the Fire Bolt cantrip."},
			{"name":"Keen Senses","desc":"You have proficiency in the Insight, Perception, or Survival skill."},
			{"name":"Versatile","desc":"You gain an Origin feat of your choice."}]}`},
		{Kind: "feat", Slug: "alert", Name: "Alert", DataJson: `{"desc":"You add your Proficiency Bonus to Initiative.","category":"Origin"}`},
		{Kind: "feat", Slug: "grappler", Name: "Grappler", DataJson: `{"desc":"You grapple well.","category":"General"}`},
		{Kind: "background", Slug: "sage", Name: "Sage", DataJson: `{"desc":"","parts":[
			{"name":"Ability Scores","desc":"Constitution, Intelligence, Wisdom","type":"ability_score"},
			{"name":"Skill Proficiencies","desc":"Arcana and History","type":"skill_proficiency"},
			{"name":"Tool Proficiency","desc":"Calligrapher's Supplies","type":"tool_proficiency"},
			{"name":"Feat","desc":"Magic Initiate (Wizard)","type":"feat"},
			{"name":"Equipment","desc":"*Choose A or B:* (A) Quarterstaff, Book, and 8 GP; or (B) 50 GP","type":"equipment"}]}`},
	} {
		entry.SourceKey, entry.ContentHash = "test", "x"
		if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
			t.Fatal(err)
		}
	}
	player := e.token(t, auth.RolePlayer)
	base := fmt.Sprintf("/characters/%d/create", e.characterID)
	if w := e.do(t, http.MethodGet, base, player, nil, false); w.Header().Get("Location") != base+"/class" {
		t.Fatalf("start: %d %v", w.Code, w.Header())
	}
	if w := e.do(t, http.MethodGet, base+"/class?pick=wizard", player, nil, false); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Choose Wizard") {
		t.Fatalf("class page: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodGet, base+"/background?pick=sage", player, nil, false); !strings.Contains(w.Body.String(), "Ability score increases") {
		t.Errorf("background page lacks the increases: %s", w.Body)
	}
	scores := url.Values{"str": {"8"}, "dex": {"14"}, "con": {"14"}, "int": {"15"}, "wis": {"12"}, "cha": {"10"}}
	// Choices saved before a class and species change nothing, so the step
	// still asks for them later (the languages checked below).
	e.do(t, http.MethodPost, base+"/choices", player, url.Values{"language": {"Orc"}}, false)
	for _, step := range []struct {
		step string
		form url.Values
		next string
	}{
		{"class", url.Values{"pick": {"wizard"}}, "/species"},
		{"species", url.Values{"pick": {"goliath"}}, "/abilities"},
		{"species", url.Values{"pick": {"goliath"}}, "/abilities"}, // once only
		{"abilities", scores, "/background"},
		{"background", url.Values{"pick": {"sage"}, "mode": {"two"}, "plus2": {"int"}, "plus1": {"con"}}, "/choices"},
		{"choices", url.Values{"trait_giant-ancestry": {"Cloud's Jaunt (Cloud Giant)"}, "trait_giant-ancestry_ability": {"Wisdom"}, "trait_keen-senses": {"Perception"},
			"trait_versatile": {"Alert"}, "language": {"Elvish", "Dwarvish", "Elvish"}, "spell": {"fire-bolt"}}, "/equipment"},
		{"equipment", url.Values{"option_background": {"B"}}, "/details"},
		{"details", url.Values{"name": {"Wren"}, "backstory": {"Raised in a library."}, "size": {"medium"}, "vision": {"Darkvision 60 ft"}}, "/review"},
		{"review", nil, fmt.Sprintf("/characters/%d", e.characterID)},
	} {
		if step.step == "choices" {
			for _, bad := range []url.Values{{"language": {"Elvish", "Dwarvish", "Orc"}}, {"trait_keen-senses": {"Arcana"}}} {
				if w := e.do(t, http.MethodPost, base+"/choices", player, bad, false); w.Code != http.StatusUnprocessableEntity {
					t.Errorf("choices %v: %d", bad, w.Code)
				}
			}
		}
		w := e.do(t, http.MethodPost, base+"/"+step.step, player, step.form, false)
		if got := w.Header().Get("Location"); w.Code != http.StatusSeeOther || !strings.HasSuffix(got, step.next) {
			t.Fatalf("%s: %d %q %s", step.step, w.Code, got, w.Body)
		}
	}
	// Choosing the background again, or base scores after it, changes
	// nothing: that would undo or repeat its increases.
	e.do(t, http.MethodPost, base+"/background", player, url.Values{"pick": {"sage"}, "mode": {"three"}}, false)
	e.do(t, http.MethodPost, base+"/abilities", player, scores, false)

	row, err := e.q().GetCharacter(ctx, e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	sh, err := dnd5e.ParseSheet(row.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	// Wizard d6 + Con: 6 at Con 10, 8 once Con is 14 (15 after the background).
	if row.Name != "Wren" || sh.ClassSummary() != "Wizard 1" || row.MaxHp != 8 {
		t.Errorf("name %q classes %q max hp %d", row.Name, sh.ClassSummary(), row.MaxHp)
	}
	if sh.Size != "Medium" || sh.Vision != "Darkvision 60 ft" {
		t.Errorf("size %q vision %q", sh.Size, sh.Vision)
	}
	if sh.Species != "Goliath" || sh.Speed != 35 || sh.Background != "Sage" || sh.Scores["int"] != 17 || sh.Scores["con"] != 15 || sh.Coins.GP != 50 {
		t.Errorf("species %q speed %d background %q int %d con %d gp %d", sh.Species, sh.Speed, sh.Background, sh.Scores["int"], sh.Scores["con"], sh.Coins.GP)
	}
	// The choices: the ancestry names its feature. Fire Bolt, picked from the
	// class list, is another ancestry's cantrip, so it keeps the class's ability.
	if sh.SpeciesAbility != "wis" {
		t.Errorf("species ability %q", sh.SpeciesAbility)
	}
	if !slices.Equal(sh.Languages, []string{"Common", "Dwarvish", "Elvish"}) || !slices.Contains(sh.SkillProfs, "perception") ||
		len(sh.Spells) != 1 || sh.Spells[0].Slug != "fire-bolt" || sh.Spells[0].Ability != "" || !slices.Contains(sh.Feats, "Alert") {
		t.Errorf("languages %q skills %q spells %+v feats %q", sh.Languages, sh.SkillProfs, sh.Spells, sh.Feats)
	}
	for _, f := range []string{"Powerful Build", "Giant Ancestry (Cloud's Jaunt (Cloud Giant))", "Alert", "Magic Initiate (Wizard)", "Arcane Recovery"} {
		if !strings.Contains(row.SheetJson, `"name":"`+f+`"`) {
			t.Errorf("missing feature %s: %s", f, row.SheetJson)
		}
	}
	if strings.Count(row.SheetJson, `"name":"Powerful Build"`) != 1 || !sh.SetUp("review") || !strings.Contains(strings.Join(sh.SkillProfs, ","), "arcana") {
		t.Errorf("sheet after the wizard: %s", row.SheetJson)
	}
}

// featFixtures are feats with choices, the spells and tools they offer, and
// a species and background that grant them.
var featFixtures = []sqlc.UpsertCompendiumEntryParams{
	{Kind: "species", Slug: "human", Name: "Human", DataJson: `{"desc":"","parts":[{"name":"Versatile","desc":"You gain an Origin feat of your choice."}]}`},
	{Kind: "feat", Slug: "alert", Name: "Alert", DataJson: `{"desc":"You add your Proficiency Bonus to Initiative.","category":"Origin"}`},
	{Kind: "feat", Slug: "magic-initiate", Name: "Magic Initiate", DataJson: `{"desc":"You gain the following benefits.","category":"Origin","parts":[
		{"name":"Two Cantrips","desc":"You learn two cantrips of your choice from the Cleric, Druid, or Wizard spell list. Intelligence, Wisdom, or Charisma is your spellcasting ability for this feat's spells."},
		{"name":"Level One Spell","desc":"Choose a level 1 spell from the same list you selected for this feat's cantrips. You always have that spell prepared."}]}`},
	{Kind: "feat", Slug: "skilled", Name: "Skilled", DataJson: `{"desc":"","category":"Origin","parts":[{"name":"","desc":"You gain proficiency in any combination of three skills or tools of your choice."}]}`},
	{Kind: "feat", Slug: "veteran", Name: "Veteran", DataJson: `{"desc":"An Origin feat with a prerequisite.","category":"Origin","prerequisite":"Level 4+"}`},
	{Kind: "feat", Slug: "grappler", Name: "Grappler", DataJson: `{"desc":"","category":"General","prerequisite":"Level 4+, Strength or Dexterity 13+","parts":[{"name":"Ability Score Increase","desc":"Increase your Strength or Dexterity score by 1, to a maximum of 20."}]}`},
	{Kind: "feat", Slug: "archery", Name: "Archery", DataJson: `{"desc":"","category":"Fighting Style","prerequisite":"Fighting Style Feature"}`},
	{Kind: "feat", Slug: "boon-of-fate", Name: "Boon of Fate", DataJson: `{"desc":"","category":"Epic Boon","prerequisite":"Level 19+","parts":[{"name":"Ability Score Increase","desc":"Increase one ability score of your choice by 1, to a maximum of 30."}]}`},
	{Kind: "feat", Slug: "ability-score-improvement", Name: "Ability Score Improvement", DataJson: `{"desc":"","category":"General"}`},
	{Kind: "background", Slug: "sage", Name: "Sage", DataJson: `{"desc":"","parts":[{"name":"Skill Proficiencies","desc":"Arcana and History","type":"skill_proficiency"},{"name":"Feat","desc":"Magic Initiate (Wizard)","type":"feat"}]}`},
	{Kind: "spell", Slug: "light", Name: "Light", DataJson: `{"level":0,"classes":["cleric","wizard"]}`},
	{Kind: "spell", Slug: "mage-hand", Name: "Mage Hand", DataJson: `{"level":0,"classes":["wizard"]}`},
	{Kind: "spell", Slug: "guidance", Name: "Guidance", DataJson: `{"level":0,"classes":["cleric"]}`},
	{Kind: "spell", Slug: "magic-missile", Name: "Magic Missile", DataJson: `{"level":1,"classes":["wizard"]}`},
	{Kind: "spell", Slug: "shield", Name: "Shield", DataJson: `{"level":1,"classes":["wizard"]}`},
	{Kind: "spell", Slug: "bless", Name: "Bless", DataJson: `{"level":1,"classes":["cleric"]}`},
	{Kind: "item", Slug: "smiths-tools", Name: "Smith's Tools (20 GP)", DataJson: `{"category":"tools","desc":"Ability: Strength. Utilize: Pry open a door (DC 20)."}`},
	{Kind: "item", Slug: "climbers-kit", Name: "Climber's Kit", DataJson: `{"category":"tools","desc":"A Climber's Kit includes boot tips and pitons."}`},
}

// loadFeatFixtures imports the test SRD and featFixtures.
func loadFeatFixtures(t *testing.T, e testEnv) {
	t.Helper()
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	for _, entry := range featFixtures {
		entry.SourceKey, entry.ContentHash = "test", "x"
		if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
			t.Fatal(err)
		}
	}
}

// TestFeatChoices: the background's Magic Initiate and Human Versatile's
// Origin feat ask for their own spells, ability and proficiencies; only the
// picked feat's choices apply, and a second Magic Initiate takes another list.
func TestFeatChoices(t *testing.T) {
	mi := "trait_versatile_magic-initiate"
	for _, tc := range []struct {
		name  string
		form  url.Values
		code  int
		check func(sh dnd5e.Sheet) bool
	}{
		{"background Magic Initiate and Skilled", url.Values{
			"feat_magic-initiate_wizard_cantrips": {"Light", "Mage Hand"}, "feat_magic-initiate_wizard_spell": {"Magic Missile"}, "feat_magic-initiate_ability": {"Intelligence"},
			"trait_versatile": {"Skilled"}, "trait_versatile_skilled_profs": {"Stealth", "Smith's Tools", "Insight"},
		}, http.StatusSeeOther, func(sh dnd5e.Sheet) bool {
			mm := slices.IndexFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == "magic-missile" })
			return len(sh.Spells) == 3 && mm >= 0 && sh.Spells[mm].Prepared && sh.Spells[mm].Ability == "int" && sh.Spells[mm].Feat == "Magic Initiate (Wizard)" && sh.SpeciesAbility == "" &&
				slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == "Magic Missile (free cast)" && f.Uses.Fixed == 1 }) &&
				slices.Contains(sh.SkillProfs, "stealth") && slices.Contains(sh.SkillProfs, "insight") && slices.Contains(sh.OtherProfs, "Smith's Tools")
		}},
		{"Versatile Magic Initiate takes another list", url.Values{
			"trait_versatile": {"Magic Initiate"}, mi + "_list": {"Cleric"}, mi + "_cleric_cantrips": {"Guidance"}, mi + "_cleric_spell": {"Bless"}, mi + "_ability": {"Wisdom"},
			// Choices of the feat not picked are ignored.
			"trait_versatile_skilled_profs": {"Stealth"},
		}, http.StatusSeeOther, func(sh dnd5e.Sheet) bool {
			bless := slices.IndexFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == "bless" })
			return bless >= 0 && sh.Spells[bless].Ability == "wis" && slices.Contains(sh.Feats, "Magic Initiate (Cleric)") &&
				slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == "Magic Initiate (Cleric)" }) && !slices.Contains(sh.SkillProfs, "stealth")
		}},
		{"the background's list again", url.Values{"trait_versatile": {"Magic Initiate"}, mi + "_list": {"Wizard"}}, http.StatusUnprocessableEntity, nil},
		{"an Origin feat whose prerequisite is unmet", url.Values{"trait_versatile": {"Veteran"}}, http.StatusUnprocessableEntity, nil},
		{"too many cantrips", url.Values{"feat_magic-initiate_wizard_cantrips": {"Light", "Mage Hand", "Guidance"}}, http.StatusUnprocessableEntity, nil},
		{"a kit is not a tool", url.Values{"trait_versatile": {"Skilled"}, "trait_versatile_skilled_profs": {"Climber's Kit"}}, http.StatusUnprocessableEntity, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := newTestEnv(t)
			ctx := t.Context()
			loadFeatFixtures(t, e)
			player := e.token(t, auth.RolePlayer)
			base := fmt.Sprintf("/characters/%d/create/", e.characterID)
			for _, step := range []string{"class=wizard", "species=human", "background=sage"} {
				step, pick, _ := strings.Cut(step, "=")
				if w := e.do(t, http.MethodPost, base+step, player, url.Values{"pick": {pick}}, false); w.Code != http.StatusSeeOther {
					t.Fatalf("%s: %d %s", step, w.Code, w.Body)
				}
			}
			if w := e.do(t, http.MethodPost, base+"choices", player, tc.form, false); w.Code != tc.code {
				t.Fatalf("choices: %d %s", w.Code, w.Body)
			}
			if tc.check == nil {
				return
			}
			row, err := e.q().GetCharacter(ctx, e.characterID)
			if err != nil {
				t.Fatal(err)
			}
			if sh, err := dnd5e.ParseSheet(row.SheetJson); err != nil || !tc.check(sh) {
				t.Errorf("sheet: %v %s", err, row.SheetJson)
			}
		})
	}
}

// TestMeetsPrerequisite checks each clause of a feat's prerequisite must hold.
func TestMeetsPrerequisite(t *testing.T) {
	style := []dnd5e.Feature{{Name: "Fighting Style"}}
	for _, tc := range []struct {
		name     string
		prereq   string
		level    int
		str, dex int
		features []dnd5e.Feature
		want     bool
	}{
		{"none", "", 1, 10, 10, nil, true},
		{"level too low", "Level 4+", 3, 10, 10, nil, false},
		{"level reached", "Level 4+", 4, 10, 10, nil, true},
		{"Epic Boon at 18", "Level 19+", 18, 10, 10, nil, false},
		{"Epic Boon at 19", "Level 19+", 19, 10, 10, nil, true},
		{"either score", "Level 4+, Strength or Dexterity 13+", 4, 12, 14, nil, true},
		{"neither score", "Level 4+, Strength or Dexterity 13+", 4, 12, 12, nil, false},
		{"feature missing", "Fighting Style Feature", 4, 10, 10, nil, false},
		{"feature held", "Fighting Style Feature", 1, 10, 10, style, true},
		{"one of two features", "Spellcasting or Pact Magic Feature", 4, 10, 10, []dnd5e.Feature{{Name: "Pact Magic"}}, true},
		{"level met, feature not", "Level 19+, Spellcasting Feature", 19, 10, 10, style, false},
		{"unknown clause passes", "Proficiency with Heavy armor", 1, 10, 10, nil, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sh := dnd5e.Sheet{Scores: map[string]int{"str": tc.str, "dex": tc.dex}, Features: tc.features}
			if got := meetsPrerequisite(tc.prereq, sh, tc.level); got != tc.want {
				t.Errorf("meetsPrerequisite(%q) = %v, want %v", tc.prereq, got, tc.want)
			}
		})
	}
}

// TestClassOptionChoices: a "… Options" feature is offered as many times as
// its table column says, without options above the character's level, and
// each pick becomes a feature.
func TestClassOptionChoices(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if err := e.q().UpsertCompendiumEntry(ctx, testlock); err != nil {
		t.Fatal(err)
	}
	row, err := e.q().GetCharacter(ctx, e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	sh := dnd5e.NewSheet()
	sh.Classes = []dnd5e.ClassLevel{{Slug: "testlock", Name: "Testlock", Level: 1, Caster: "pact"}}
	cr, err := e.app.creationChoices(ctx, sh, row.CampaignID)
	if err != nil {
		t.Fatal(err)
	}
	i := slices.IndexFunc(cr.groups, func(g choiceGroup) bool { return g.Key == "class_eldritch-invocation" })
	if i < 0 {
		t.Fatalf("no invocation group: %+v", cr.view.Groups)
	}
	g := cr.groups[i]
	if g.Count != 1 || g.Title != "Eldritch Invocations" || len(g.Options) != 1 || g.Options[0].Name != "Armor of Shadows" {
		t.Fatalf("group %+v", g.ChoiceGroup)
	}
	g.apply(&sh, []string{"Armor of Shadows"})
	if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool {
		return f.Name == "Eldritch Invocation: Armor of Shadows" && f.Source == "Testlock" && f.Description == "Mage Armor."
	}) {
		t.Errorf("features %+v", sh.Features)
	}
	// At level 2 the column grows by 2; the level 2 option opens up and the
	// one taken is gone.
	info, err := classInfo(sqlc.CompendiumEntry{Slug: testlock.Slug, Name: testlock.Name, DataJson: testlock.DataJson})
	if err != nil {
		t.Fatal(err)
	}
	if gs := classOptionGroups(info, 2, 1, sh.Features); len(gs) != 1 || gs[0].Count != 2 || len(gs[0].Options) != 1 || gs[0].Options[0].Name != "Agonizing Blast" {
		t.Errorf("level 2: %+v", groupViews(gs))
	}
	// A Repeatable option stays on offer and is added again under a new name.
	for range 2 {
		gs := classOptionGroups(info, 2, 1, sh.Features)
		if len(gs) != 1 || len(gs[0].Options) != 1 {
			t.Fatalf("repeatable: %+v", groupViews(gs))
		}
		gs[0].apply(&sh, []string{"Agonizing Blast"})
	}
	for _, name := range []string{"Eldritch Invocation: Agonizing Blast", "Eldritch Invocation: Agonizing Blast (2)"} {
		if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == name }) {
			t.Errorf("no %q in %+v", name, sh.Features)
		}
	}
}

// testlock is a pact caster with invocations: one at level 1, three at 2.
var testlock = sqlc.UpsertCompendiumEntryParams{Kind: "class", Slug: "testlock", Name: "Testlock", SourceKey: "test", ContentHash: "x", DataJson: `{"caster":"pact","hit_die":8,
	"table":{"Eldritch Invocations":{"1":"1","2":"3"}},
	"features":[{"name":"Eldritch Invocations","desc":"You gain one invocation.","levels":[1]},
		{"name":"Eldritch Invocation Options","desc":"In order.\n\n### Agonizing Blast\n\n*Prerequisite: Level 2+ Warlock*\n\n**Repeatable.** More damage.\n\n### Armor of Shadows\n\nMage Armor."}]}`}

// TestApplyTraitText: trait text sets darkvision, a speed increase and
// resistances, but "the damage type determined by…" names none.
func TestApplyTraitText(t *testing.T) {
	for _, tt := range []struct {
		text   string
		vision string
		speed  int
		resist []string
	}{
		{"The range of your Darkvision increases to 120 feet.", "Darkvision 120 ft", 30, nil},
		{"Your Speed increases to 35 feet. You also know the Druidcraft cantrip.", "", 35, nil},
		{"You have Resistance to Poison damage. You also know the Poison Spray cantrip.", "", 30, []string{"poison"}},
		{"You have Resistance to the damage type determined by your Draconic Ancestry trait.", "", 30, nil},
	} {
		s := dnd5e.Sheet{Speed: 30}
		applyTraitText(&s, tt.text)
		if s.Vision != tt.vision || s.Speed != tt.speed || !slices.Equal(s.Resistances, tt.resist) {
			t.Errorf("%q: vision %q speed %d resist %q", tt.text, s.Vision, s.Speed, s.Resistances)
		}
	}
}

// TestTextMarkdown: unnamed parts (a feat's body) get no empty bold label.
func TestTextMarkdown(t *testing.T) {
	got := textMarkdown(&compendium.TextData{Parts: []compendium.NameDesc{{Desc: "You strike twice."}, {Name: "Repeatable", Desc: "Again."}}})
	if want := "You strike twice.\n\n**Repeatable.** Again."; got != want {
		t.Errorf("got %q", got)
	}
}

// TestSpeciesResistance: a species' fixed resistance applies with the
// species, and an ancestry's damage type once it is picked.
func TestSpeciesResistance(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	species := sqlc.UpsertCompendiumEntryParams{Kind: "species", Slug: "testborn", Name: "Testborn", SourceKey: "test", ContentHash: "x", DataJson: `{"desc":"","parts":[
		{"name":"Draconic Ancestry","desc":"Choose the kind of dragon.\n\n| Dragon | Damage Type |\n|---|---|\n| Red | Fire |\n| White | Cold |"},
		{"name":"Damage Resistance","desc":"You have Resistance to the damage type determined by your Draconic Ancestry trait."},
		{"name":"Tough Hide","desc":"You have Resistance to Poison damage."}]}`}
	if err := e.q().UpsertCompendiumEntry(ctx, species); err != nil {
		t.Fatal(err)
	}
	d, err := compendium.DecodeData(compendium.KindSpecies, species.DataJson)
	if err != nil {
		t.Fatal(err)
	}
	sh := dnd5e.NewSheet()
	applySpecies(&sh, "Testborn", d.(*compendium.TextData))
	cr, err := e.app.creationChoices(ctx, sh, 1)
	if err != nil {
		t.Fatal(err)
	}
	i := slices.IndexFunc(cr.groups, func(g choiceGroup) bool { return g.Key == "trait_draconic-ancestry" })
	if i < 0 {
		t.Fatalf("no ancestry group: %+v", cr.view.Groups)
	}
	cr.groups[i].apply(&sh, []string{"Red"})
	if !slices.Equal(sh.Resistances, []string{"poison", "fire"}) {
		t.Errorf("resistances %q", sh.Resistances)
	}
}

// TestLineageSpells: the picked lineage's "Level 3"/"Level 5" spells come at
// those character levels, prepared, with the ability of its cantrip.
func TestLineageSpells(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	for _, entry := range []sqlc.UpsertCompendiumEntryParams{
		{Kind: "species", Slug: "testelf", Name: "Testelf", DataJson: `{"desc":"","parts":[
			{"name":"Lineage","desc":"Choose a lineage.\n\n|Lineage|Level 1|Level 3|Level 5|\n|---|---|---|---|\n|Drow|You also know the Dancing Lights cantrip.|Faerie Fire|Darkness|\n|Wood|You know the Druidcraft cantrip.|Longstrider|Pass without Trace|"}]}`},
		{Kind: "spell", Slug: "faerie-fire", Name: "Faerie Fire", DataJson: `{"level":1}`},
		{Kind: "spell", Slug: "darkness", Name: "Darkness", DataJson: `{"level":2}`},
	} {
		entry.SourceKey, entry.ContentHash = "test", "x"
		if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
			t.Fatal(err)
		}
	}
	drow := dnd5e.NewSheet()
	drow.Species = "Testelf"
	drow.Features = []dnd5e.Feature{{Name: "Lineage (Drow)", Source: "Testelf"}}
	drow.Spells = []dnd5e.Spell{{Slug: "dancing-lights", Name: "Dancing Lights", Ability: "int"}}
	// Faerie Fire from the class list, so its free cast is still owed.
	knowsFire := drow
	knowsFire.Spells = append(slices.Clone(drow.Spells), dnd5e.Spell{Slug: "faerie-fire", Name: "Faerie Fire", Level: 1})
	withFire := knowsFire
	withFire.Features = append(slices.Clone(drow.Features), dnd5e.Feature{Name: "Faerie Fire (free cast)", Source: "Testelf"})
	noPick := drow
	noPick.Features = nil
	// The cantrip was swapped; the ability picked at creation is kept.
	swapped := drow
	swapped.Spells, swapped.SpeciesAbility = nil, "wis"
	for _, tt := range []struct {
		name    string
		sheet   dnd5e.Sheet
		level   int
		want    []string
		ability string
	}{
		{"level 2", drow, 2, nil, "int"},
		{"level 3", drow, 3, []string{"faerie-fire"}, "int"},
		{"level 5 catches up", drow, 5, []string{"faerie-fire", "darkness"}, "int"},
		{"level 5 has level 3's", withFire, 5, []string{"darkness"}, "int"},
		{"known from the class list", knowsFire, 3, []string{"faerie-fire"}, "int"},
		{"no lineage picked", noPick, 5, nil, "int"},
		{"cantrip swapped", swapped, 3, []string{"faerie-fire"}, "wis"},
	} {
		got, err := e.app.lineageSpells(ctx, tt.sheet, tt.level, 1)
		if err != nil {
			t.Fatal(err)
		}
		var slugs []string
		for _, sp := range got {
			slugs = append(slugs, sp.Slug)
			if !sp.Prepared || sp.Ability != tt.ability {
				t.Errorf("%s: %+v not prepared with %s", tt.name, sp, tt.ability)
			}
		}
		if !slices.Equal(slugs, tt.want) {
			t.Errorf("%s: %q, want %q", tt.name, slugs, tt.want)
		}
	}
}

// TestApplySpeciesSizeVision: a single size and a darkvision range are set;
// "Medium or Small" leaves the size for the player.
func TestApplySpeciesSizeVision(t *testing.T) {
	tests := []struct {
		size, wantSize string
		parts          []compendium.NameDesc
		wantVision     string
	}{
		{"Medium", "Medium", []compendium.NameDesc{{Name: "Darkvision", Desc: "You can see in dim light within 60 feet of you."}}, "Darkvision 60 ft"},
		{"Medium or Small", "", nil, ""},
		{"Small.", "Small", []compendium.NameDesc{{Name: "Superior Darkvision", Desc: "Your darkvision has a range of 120 feet."}}, "Darkvision 120 ft"},
	}
	for _, tt := range tests {
		var s dnd5e.Sheet
		parts := append([]compendium.NameDesc{{Name: "Size", Desc: tt.size, Type: "size"}}, tt.parts...)
		applySpecies(&s, "Test", &compendium.TextData{Parts: parts})
		if s.Size != tt.wantSize || s.Vision != tt.wantVision {
			t.Errorf("%q: size %q vision %q", tt.size, s.Size, s.Vision)
		}
	}
}

// TestFindItemPlurals: equipment lines match their entry singular too.
func TestFindItemPlurals(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	for _, slug := range []string{"dagger", "pouch", "glass"} {
		entry := sqlc.UpsertCompendiumEntryParams{Kind: string(compendium.KindItem), Slug: slug, Name: slug, DataJson: "{}", SourceKey: "test", ContentHash: "x"}
		if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
			t.Fatal(err)
		}
	}
	for _, tt := range []struct{ line, want string }{
		{"Dagger", "dagger"}, {"Daggers", "dagger"}, {"Pouches", "pouch"}, {"Glasses", "glass"}, {"Ropes", ""},
	} {
		got, err := e.app.findItem(ctx, tt.line)
		if tt.want == "" {
			if !errors.Is(err, sql.ErrNoRows) {
				t.Errorf("%s: got %q, %v; want no rows", tt.line, got.Slug, err)
			}
		} else if err != nil || got.Slug != tt.want {
			t.Errorf("%s: got %q, %v; want %q", tt.line, got.Slug, err, tt.want)
		}
	}
}

// TestApplyEquipmentMergesGear: an equipment line found in the compendium is
// stored under the entry's name, so the weapon and the stack are one item.
func TestApplyEquipmentMergesGear(t *testing.T) {
	picked := []compendium.EquipmentOption{{Items: []compendium.EquipmentItem{{Name: "Daggers", Qty: 4}, {Name: "Pouches", Qty: 2}, {Name: "Rope", Qty: 1}}}}
	dagger := dnd5e.Weapon{Name: "Dagger", Ability: "finesse", Dice: "1d4"}
	g := gear{weapons: []dnd5e.Weapon{dagger}, names: map[string]string{"Daggers": "Dagger", "Pouches": "Pouch"}}
	var s dnd5e.Sheet
	applyEquipment(&s, picked, g)
	sheetJSON, err := s.JSON()
	if err != nil {
		t.Fatal(err)
	}
	s, err = dnd5e.ParseSheet(sheetJSON)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, it := range s.Inventory {
		got = append(got, fmt.Sprintf("%s %d %v", it.Name, it.Qty, it.Weapon != nil))
	}
	if want := []string{"Dagger 4 true", "Pouch 2 false", "Rope 1 false"}; !slices.Equal(got, want) {
		t.Errorf("inventory %v, want %v", got, want)
	}
}
