package web

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/web/views"
)

// TestWizardStepSaves: every step saves through the table; one left out
// would save nothing.
func TestWizardStepSaves(t *testing.T) {
	for _, s := range wizardSteps {
		if wizardStepSaves[s.Key].command == nil {
			t.Errorf("step %s has no command in wizardStepSaves", s.Key)
		}
	}
	if len(wizardStepSaves) != len(wizardSteps) {
		t.Errorf("wizardStepSaves has %d steps, wizardSteps %d", len(wizardStepSaves), len(wizardSteps))
	}
}

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
	e.allowEntrySources(t)
	player := e.token(t, enum.RolePlayer)
	base := fmt.Sprintf("/characters/%d/create", e.characterID)
	if w := e.do(t, http.MethodGet, base, player, nil, false); w.Header().Get("Location") != base+"/class" {
		t.Fatalf("start: %d %v", w.Code, w.Header())
	}
	if w := e.do(t, http.MethodGet, base+"/class?pick=wizard", player, nil, false); w.Code != http.StatusOK || !strings.Contains(w.Body.String(), "Choose Wizard") {
		t.Fatalf("class page: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodGet, base+"/abilities", player, nil, false); !strings.Contains(w.Body.String(), "Choose these first") {
		t.Errorf("abilities page before a background does not ask for it: %s", w.Body)
	}
	scores := url.Values{"str": {"8"}, "dex": {"14"}, "con": {"14"}, "int": {"15"}, "wis": {"12"}, "cha": {"10"}}
	raised := maps.Clone(scores)
	raised.Set("mode", "two")
	raised.Set("plus2", "int")
	raised.Set("plus1", "con")
	// Choices and scores saved before what they need change nothing, so the
	// steps still ask for them later (the languages checked below).
	e.do(t, http.MethodPost, base+"/choices", player, url.Values{"language": {"Orc"}}, false)
	e.do(t, http.MethodPost, base+"/abilities", player, scores, false)
	for _, step := range []struct {
		step string
		form url.Values
		next string
	}{
		{"class", url.Values{"pick": {"wizard"}}, "/background"},
		{"background", url.Values{"pick": {"sage"}}, "/species"},
		{"species", url.Values{"pick": {"goliath"}}, "/abilities"},
		{"species", url.Values{"pick": {"goliath"}}, "/abilities"}, // once only
		{"abilities", raised, "/choices"},
		{"choices", url.Values{"trait_giant-ancestry": {"Cloud's Jaunt (Cloud Giant)"}, "trait_giant-ancestry_ability": {"Wisdom"}, "trait_keen-senses": {"Perception"},
			"trait_versatile": {"Alert"}, "language": {"Elvish", "Dwarvish", "Elvish"}, "spell": {"fire-bolt"}}, "/equipment"},
		{"equipment", url.Values{"option_background": {"B"}}, "/details"},
		{"details", url.Values{"name": {"Wren"}, "backstory": {"Raised in a library."}, "size": {"medium"}, "vision": {"Darkvision 60 ft"}}, "/review"},
		{"review", nil, fmt.Sprintf("/characters/%d", e.characterID)},
	} {
		if step.step == "abilities" {
			if w := e.do(t, http.MethodGet, base+"/abilities", player, nil, false); !strings.Contains(w.Body.String(), "Sage: ability score increases") {
				t.Errorf("abilities page lacks the background's increases: %s", w.Body)
			}
			// Refused, the page keeps the scores and increases asked for and
			// shows every error: two increases the same and Wisdom above 20;
			// +1 to all three taking Wisdom above 20; two scores out of range.
			same := maps.Clone(scores)
			same.Set("mode", "two")
			same.Set("plus2", "wis")
			same.Set("plus1", "wis")
			same.Set("wis", "21")
			high := maps.Clone(scores)
			high.Set("wis", "20")
			high.Set("mode", "three")
			low := maps.Clone(scores)
			low.Set("str", "2")
			low.Set("cha", "21")
			for _, bad := range []struct {
				form url.Values
				at   []string
			}{
				{form: same, at: []string{views.FieldMode, "wis"}},
				{form: high, at: []string{"wis"}},
				{form: low, at: []string{"str", "cha"}},
			} {
				w := e.do(t, http.MethodPost, base+"/abilities", player, bad.form, false)
				if lost := lostPicks(w.Body.String(), bad.form); w.Code != http.StatusUnprocessableEntity || len(lost) > 0 {
					t.Errorf("abilities %v: %d, lost %v", bad.form, w.Code, lost)
				}
				for _, at := range bad.at {
					if !errorAt(w.Body.String(), at) {
						t.Errorf("abilities %v: no error at %s", bad.form, at)
					}
				}
			}
		}
		if step.step == "details" {
			// A blank name and an unknown size are refused at their fields
			// together, and the page keeps the rest as typed.
			blank := url.Values{"name": {"  "}, "backstory": {"Raised in a library."}, "size": {"Huge-ish"}, "vision": {"Darkvision 60 ft"}}
			w := e.do(t, http.MethodPost, base+"/details", player, blank, false)
			if w.Code != http.StatusUnprocessableEntity || !errorAt(w.Body.String(), views.FieldName) || !errorAt(w.Body.String(), views.FieldSize) {
				t.Errorf("details %v: %d, error at name %v, at size %v", blank, w.Code, errorAt(w.Body.String(), views.FieldName), errorAt(w.Body.String(), views.FieldSize))
			}
			for _, kept := range []string{"Raised in a library.", `value="Darkvision 60 ft"`} {
				if !strings.Contains(w.Body.String(), kept) {
					t.Errorf("refused details page lost %q", kept)
				}
			}
		}
		if step.step == "choices" {
			for _, bad := range []url.Values{{"language": {"Elvish", "Dwarvish", "Orc"}}, {"trait_keen-senses": {"Arcana"}}} {
				if w := e.do(t, http.MethodPost, base+"/choices", player, bad, false); w.Code != http.StatusUnprocessableEntity {
					t.Errorf("choices %v: %d", bad, w.Code)
				}
			}
			// Every wrong group shows its own error in one refusal.
			w := e.do(t, http.MethodPost, base+"/choices", player, url.Values{"language": {"Orc"}, "trait_keen-senses": {"Arcana"}}, false)
			for _, at := range []string{"language", "trait_keen-senses"} {
				if !errorAt(w.Body.String(), at) {
					t.Errorf("choices: no error at %s", at)
				}
			}
			if !strings.Contains(w.Body.String(), "Common (every character knows it)") {
				t.Error("choices page does not show Common as known")
			}
		}
		w := e.do(t, http.MethodPost, base+"/"+step.step, player, step.form, false)
		if got := w.Header().Get("Location"); w.Code != http.StatusSeeOther || !strings.HasSuffix(got, step.next) {
			t.Fatalf("%s: %d %q %s", step.step, w.Code, got, w.Body)
		}
	}
	// Choosing the background or the scores again changes nothing: that
	// would repeat the increases.
	e.do(t, http.MethodPost, base+"/background", player, url.Values{"pick": {"sage"}}, false)
	e.do(t, http.MethodPost, base+"/abilities", player, raised, false)

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
	if strings.Count(row.SheetJson, `"name":"Powerful Build"`) != 1 || !sh.SetUp("review") || !strings.Contains(strings.Join(sh.SkillProfs, ","), "arcana") || sh.SetupSnapshots != nil {
		t.Errorf("sheet after the wizard: %s", row.SheetJson)
	}
	// Finished: no step can be redone in the wizard any more.
	if w := e.do(t, http.MethodPost, base+"/species/change", player, nil, false); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("change after the review: %d", w.Code)
	}
}

// TestWizardChange: redoing a step puts the character back to before it,
// undoing the steps saved after it but keeping the earlier ones redoable.
func TestWizardChange(t *testing.T) {
	e := newTestEnv(t)
	loadFeatFixtures(t, e)
	player := e.token(t, enum.RolePlayer)
	base := fmt.Sprintf("/characters/%d/create", e.characterID)
	for _, step := range []struct {
		step string
		form url.Values
	}{
		{"class", url.Values{"pick": {"wizard"}}},
		{"background", url.Values{"pick": {"sage"}}},
		{"species", url.Values{"pick": {"human"}}},
		{"abilities", url.Values{"str": {"8"}, "dex": {"14"}, "con": {"14"}, "int": {"15"}, "wis": {"12"}, "cha": {"10"}}},
	} {
		if w := e.do(t, http.MethodPost, base+"/"+step.step, player, step.form, false); w.Code != http.StatusSeeOther {
			t.Fatalf("%s: %d %s", step.step, w.Code, w.Body)
		}
	}
	// A level-up taken mid-creation: undone with the step, so Revert must not bring it back.
	if err := e.q().CreateLevelUp(t.Context(), sqlc.CreateLevelUpParams{CharacterID: e.characterID, MaxHp: 8, SheetJson: "{}"}); err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		step string
		code int
	}{
		{"nonsense", http.StatusNotFound},
		{"details", http.StatusUnprocessableEntity}, // no snapshot: edited directly
		{"background", http.StatusSeeOther},
		{"species", http.StatusUnprocessableEntity}, // undone with the background
	} {
		if w := e.do(t, http.MethodPost, base+"/"+c.step+"/change", player, nil, false); w.Code != c.code {
			t.Errorf("change %s: %d, want %d", c.step, w.Code, c.code)
		}
	}
	row, err := e.q().GetCharacter(t.Context(), e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	sh, err := dnd5e.ParseSheet(row.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	if sh.Background != "" || sh.Species != "" || sh.Scores["con"] != 10 || !slices.Equal(sh.SetupDone, []string{"class"}) ||
		!slices.Equal(slices.Sorted(maps.Keys(sh.SetupSnapshots)), []string{"class"}) || sh.ClassSummary() != "Wizard 1" || row.MaxHp != 6 {
		t.Errorf("after changing the background: max hp %d, sheet %s", row.MaxHp, row.SheetJson)
	}
	if n, err := e.q().CountLevelUps(t.Context(), e.characterID); err != nil || n != 0 {
		t.Errorf("level-ups left after the change: %d %v", n, err)
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
	// A background that asks nothing, so the choices step opens with only the class's and species' picks.
	{Kind: "background", Slug: "testfolk", Name: "Testfolk", DataJson: `{"desc":"","parts":[]}`},
	{Kind: "spell", Slug: "light", Name: "Light", DataJson: `{"level":0,"classes":["cleric","wizard"]}`},
	{Kind: "spell", Slug: "mage-hand", Name: "Mage Hand", DataJson: `{"level":0,"classes":["wizard"]}`},
	{Kind: "spell", Slug: "guidance", Name: "Guidance", DataJson: `{"level":0,"classes":["cleric"]}`},
	// Enough Wizard cantrips besides the feat's Light and Mage Hand.
	{Kind: "spell", Slug: "ray-of-frost", Name: "Ray of Frost", DataJson: `{"level":0,"classes":["wizard"]}`},
	{Kind: "spell", Slug: "minor-illusion", Name: "Minor Illusion", DataJson: `{"level":0,"classes":["wizard"]}`},
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
	e.allowEntrySources(t)
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
			"trait_versatile": {"Skilled"}, "trait_versatile_skilled_profs": {"Stealth", "Smith's Tools", "Insight"},
		}, http.StatusSeeOther, func(sh dnd5e.Sheet) bool {
			mm := slices.IndexFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == "magic-missile" })
			feat := slices.DeleteFunc(slices.Clone(sh.Spells), func(sp dnd5e.Spell) bool { return sp.Feat == "" })
			return len(feat) == 3 && mm >= 0 && sh.Spells[mm].Prepared && sh.Spells[mm].Ability == "int" && sh.Spells[mm].Feat == "Magic Initiate (Wizard)" && sh.SpeciesAbility == "" &&
				slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == "Magic Missile (free cast)" && f.Uses.Fixed == 1 }) &&
				slices.Contains(sh.SkillProfs, "stealth") && slices.Contains(sh.SkillProfs, "insight") && slices.Contains(sh.OtherProfs, "Smith's Tools") &&
				sh.FeatSources["skilled"] == "test"
		}},
		{"Versatile Magic Initiate takes another list", url.Values{
			"trait_versatile": {"Magic Initiate"}, mi + "_list": {"Cleric"}, mi + "_cleric_cantrips": {"Guidance", "Light"}, mi + "_cleric_spell": {"Bless"}, mi + "_ability": {"Wisdom"},
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
		{"a kit is not a tool", url.Values{"trait_versatile": {"Skilled"}, "trait_versatile_skilled_profs": {"Stealth", "Insight", "Climber's Kit"}}, http.StatusUnprocessableEntity, nil},
		{"too few cantrips", url.Values{"feat_magic-initiate_wizard_cantrips": {"Light"}}, http.StatusUnprocessableEntity, nil},
		{"one language of two", url.Values{"language": {"Elvish"}}, http.StatusUnprocessableEntity, nil},
	} {
		// Every choice the step asks for, so each case fails only on its own.
		form := url.Values{"language": {"Elvish", "Dwarvish"}, "trait_versatile": {"Alert"},
			"feat_magic-initiate_wizard_cantrips": {"Light", "Mage Hand"}, "feat_magic-initiate_wizard_spell": {"Magic Missile"}, "feat_magic-initiate_ability": {"Intelligence"}}
		maps.Copy(form, tc.form)
		t.Run(tc.name, func(t *testing.T) {
			e := newTestEnv(t)
			ctx := t.Context()
			loadFeatFixtures(t, e)
			player := e.token(t, enum.RolePlayer)
			base := fmt.Sprintf("/characters/%d/create/", e.characterID)
			for _, step := range []string{"class=wizard", "species=human", "background=sage"} {
				step, pick, _ := strings.Cut(step, "=")
				if w := e.do(t, http.MethodPost, base+step, player, url.Values{"pick": {pick}}, false); w.Code != http.StatusSeeOther {
					t.Fatalf("%s: %d %s", step, w.Code, w.Body)
				}
			}
			w := e.do(t, http.MethodPost, base+"choices", player, creationSpells(t, e, form), false)
			if w.Code != tc.code {
				t.Fatalf("choices: %d %s", w.Code, w.Body)
			}
			if lost := lostPicks(w.Body.String(), form); tc.code == http.StatusUnprocessableEntity && len(lost) > 0 {
				t.Errorf("refused page lost %v", lost)
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

// TestCreationSpeciesSpells: a class pick the species teaches in the same
// step is refused (a lineage picked in the form) or not offered (a fixed
// cantrip), so the class pick is not lost to the species' copy.
func TestCreationSpeciesSpells(t *testing.T) {
	const lineage = "trait_gnomish-lineage"
	for _, tc := range []struct {
		name, species string
		form          url.Values
		code          int
	}{
		{"the picked lineage's cantrip", "testgnome", url.Values{lineage: {"Forest"}, "spell": {"light"}}, http.StatusUnprocessableEntity},
		{"another lineage's cantrip", "testgnome", url.Values{lineage: {"Rock"}, "spell": {"light"}}, http.StatusSeeOther},
		{"a fixed cantrip", "testglow", url.Values{}, http.StatusSeeOther},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := newTestEnv(t)
			ctx := t.Context()
			loadFeatFixtures(t, e)
			for _, entry := range []sqlc.UpsertCompendiumEntryParams{
				{Kind: "species", Slug: "testgnome", Name: "Testgnome", DataJson: `{"desc":"","parts":[
					{"name":"Gnomish Lineage","desc":"Choose a lineage.\n\n|Lineage|Level 1|\n|---|---|\n|Forest|You know the Light cantrip.|\n|Rock|You know the Guidance cantrip.|"}]}`},
				{Kind: "species", Slug: "testglow", Name: "Testglow", DataJson: `{"desc":"","parts":[{"name":"Glow","desc":"You know the Light cantrip."}]}`},
			} {
				entry.SourceKey, entry.ContentHash = "test", "x"
				if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
					t.Fatal(err)
				}
			}
			player := e.token(t, enum.RolePlayer)
			base := fmt.Sprintf("/characters/%d/create/", e.characterID)
			for _, step := range []string{"class=wizard", "species=" + tc.species, "background=testfolk"} {
				step, pick, _ := strings.Cut(step, "=")
				if w := e.do(t, http.MethodPost, base+step, player, url.Values{"pick": {pick}}, false); w.Code != http.StatusSeeOther {
					t.Fatalf("%s: %d %s", step, w.Code, w.Body)
				}
			}
			ch, _ := e.q().GetCharacter(ctx, e.characterID)
			sh, _ := dnd5e.ParseSheet(ch.SheetJson)
			cr, err := e.app.creationChoices(ctx, sh, ch.CampaignID)
			if err != nil {
				t.Fatal(err)
			}
			if offersLight := slices.ContainsFunc(cr.view.Spells, func(o views.Option) bool { return o.Slug == "light" }); offersLight == (tc.species == "testglow") {
				t.Errorf("class list offers Light: %v", offersLight)
			}
			// The lineage that teaches Light hides it from the class list while picked.
			if i := slices.IndexFunc(cr.view.Spells, func(o views.Option) bool { return o.Slug == "light" }); tc.species == "testgnome" && i >= 0 {
				if got := cr.view.Spells[i].TaughtBy; !slices.Equal(got, []string{"trait_gnomish-lineage=Forest"}) {
					t.Errorf("Light taught by %v, want the Forest lineage", got)
				}
			}
			tc.form["language"] = []string{"Elvish", "Dwarvish"}
			w := e.do(t, http.MethodPost, base+"choices", player, creationSpells(t, e, tc.form), false)
			if w.Code != tc.code {
				t.Fatalf("choices: %d %s", w.Code, w.Body)
			}
			if tc.code == http.StatusUnprocessableEntity {
				body := w.Body.String()
				// Light is dropped on purpose: the picked Forest lineage teaches it.
				lost := slices.DeleteFunc(lostPicks(body, tc.form), func(l string) bool { return l == "spell=light" })
				if len(lost) > 0 || !errorAt(body, views.FieldSpell) {
					t.Errorf("refused page lost %v, or the error is not at the spells", lost)
				}
				// Unticked, not disabled, so it can be picked once shown again.
				if row := reHiddenRow.FindStringSubmatch(body); row == nil || !strings.Contains(row[1], `value="light"`) || reChecked.MatchString(row[1]) || reDisabled.MatchString(row[1]) {
					t.Errorf("Light row not hidden and unticked: %v", row)
				}
				return
			}
			ch, _ = e.q().GetCharacter(ctx, e.characterID)
			sh, _ = dnd5e.ParseSheet(ch.SheetJson)
			if lights := slices.DeleteFunc(slices.Clone(sh.Spells), func(sp dnd5e.Spell) bool { return sp.Slug != "light" }); len(lights) != 1 {
				t.Errorf("Light %d times: %+v", len(lights), sh.Spells)
			}
		})
	}
}

// TestCheckNewOptions: an option a level-up added must meet its prerequisite
// on the result and name a cantrip the sheet knows; older ones are not checked.
func TestCheckNewOptions(t *testing.T) {
	info, err := classInfo(sqlc.CompendiumEntry{Slug: testlock.Slug, Name: testlock.Name, DataJson: testlock.DataJson})
	if err != nil {
		t.Fatal(err)
	}
	opt := func(name string) dnd5e.Feature {
		for _, f := range info.Features {
			for _, o := range compendium.Options(f.Description) {
				if strings.HasPrefix(name, o.Name) {
					return dnd5e.Feature{Name: "Eldritch Invocation: " + name, Source: "Testlock", Description: o.Desc}
				}
			}
		}
		t.Fatalf("no option %s", name)
		return dnd5e.Feature{}
	}
	blast := []dnd5e.Cantrip{{Name: "Test Blast", Classes: []string{"testlock"}, Range: "120 feet", Desc: "takes 1d10 Force damage"}}
	armor, step, spear := opt("Armor of Shadows"), opt("Shadow Step"), opt("Eldritch Spear (Test Blast)")
	for _, tc := range []struct {
		name          string
		before, after []dnd5e.Feature
		cantrips      []dnd5e.Cantrip
		err           string
	}{
		{"picked together", nil, []dnd5e.Feature{armor, step}, nil, ""},
		{"prerequisite missing", nil, []dnd5e.Feature{step}, nil, "Shadow Step needs"},
		{"held before is not checked", []dnd5e.Feature{step}, []dnd5e.Feature{step}, nil, ""},
		{"cantrip known", nil, []dnd5e.Feature{spear}, blast, ""},
		{"cantrip not known", nil, []dnd5e.Feature{opt("Eldritch Spear (Fire Bolt)")}, blast, "needs Fire Bolt"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := checkNewOptions(info, dnd5e.Sheet{Features: tc.before}, dnd5e.Sheet{Features: tc.after}, 2, tc.cantrips)
			if (tc.err == "") != (err == nil) || err != nil && !strings.Contains(err.Error(), tc.err) {
				t.Errorf("checkNewOptions = %v, want %q", err, tc.err)
			}
		})
	}
}

// TestBackgroundRepicksClassSkill: a Monk picks Insight; an Acolyte
// background, which gives Insight, asks for one more Monk skill, and
// Insight is held once. Redone afterwards, the class step shows the
// background's skills known and offers the rest.
func TestBackgroundRepicksClassSkill(t *testing.T) {
	e := newTestEnv(t)
	for _, p := range []sqlc.UpsertCompendiumEntryParams{
		{Kind: "class", Slug: "testmonk", Name: "Testmonk", SourceKey: "test", ContentHash: "x",
			DataJson: `{"hit_die":8,"desc":"|Skill Proficiencies|Choose 2: Acrobatics, Athletics, History, Insight, Religion, or Stealth|"}`},
		{Kind: "background", Slug: "testacolyte", Name: "Testacolyte", SourceKey: "test", ContentHash: "x",
			DataJson: `{"desc":"","parts":[{"name":"Skill Proficiencies","type":"skill_proficiency","desc":"Insight, Religion"}]}`},
	} {
		if err := e.q().UpsertCompendiumEntry(t.Context(), p); err != nil {
			t.Fatal(err)
		}
	}
	base, player := fmt.Sprintf("/characters/%d/create/", e.characterID), e.token(t, enum.RolePlayer)
	if w := e.do(t, http.MethodPost, base+"class", player, url.Values{"pick": {"testmonk"}, "skill": {"insight", "stealth"}}, false); w.Code != http.StatusSeeOther {
		t.Fatalf("class: %d %s", w.Code, w.Body)
	}
	page := e.do(t, http.MethodGet, base+"background?pick=testacolyte", player, nil, false).Body.String()
	if !strings.Contains(page, "choose one more") || !strings.Contains(page, `value="athletics"`) || strings.Contains(page, `value="insight"`) {
		t.Errorf("background page does not ask for one more class skill")
	}
	for _, tc := range []struct {
		name   string
		skills []string
		code   int
	}{
		{"none", nil, http.StatusUnprocessableEntity},
		{"a background skill", []string{"stealth", "religion"}, http.StatusUnprocessableEntity},
		{"not a class skill", []string{"stealth", "perception"}, http.StatusUnprocessableEntity},
		{"one more", []string{"stealth", "athletics"}, http.StatusSeeOther},
	} {
		w := e.do(t, http.MethodPost, base+"background", player, url.Values{"pick": {"testacolyte"}, "skill": tc.skills}, false)
		if w.Code != tc.code {
			t.Errorf("%s: %d, want %d", tc.name, w.Code, tc.code)
		}
	}
	sh := e.sheet(t)
	if got := slices.Sorted(slices.Values(sh.SkillProfs)); !slices.Equal(got, []string{"athletics", "insight", "religion", "stealth"}) ||
		sh.ProfSources["insight"] != "Testacolyte" || sh.ProfSources["athletics"] != "Testmonk" {
		t.Errorf("skills %q, sources %v", sh.SkillProfs, sh.ProfSources)
	}

	// A class step after the background shows its skills known.
	s := dnd5e.NewSheet()
	s.Background, s.SkillProfs = "Testacolyte", []string{"insight", "religion"}
	p := classSkillPick(s, []string{"acrobatics", "athletics", "history", "insight", "religion", "stealth"}, 2, nil)
	if !slices.Equal(p.Known, []string{"insight", "religion"}) || len(p.Options) != 4 || p.Count != 2 {
		t.Errorf("class pick after background: %+v", p)
	}
}

// TestClassSkillCount: the class step takes exactly its skill count, of its
// listed skills, each once.
func TestClassSkillCount(t *testing.T) {
	class := sqlc.UpsertCompendiumEntryParams{Kind: "class", Slug: "testfighter", Name: "Testfighter", SourceKey: "test", ContentHash: "x",
		DataJson: `{"hit_die":10,"desc":"|Skill Proficiencies|Choose 2: Athletics, Insight, or Stealth|"}`}
	// A class with no skills table draws no skills field.
	e0 := newTestEnv(t)
	if err := e0.q().UpsertCompendiumEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{Kind: "class", Slug: "testcommoner", Name: "Testcommoner", SourceKey: "test", ContentHash: "x", DataJson: `{"hit_die":8}`}); err != nil {
		t.Fatal(err)
	}
	w0 := e0.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/create/class", e0.characterID), e0.token(t, enum.RolePlayer), url.Values{"pick": {"testcommoner"}, "skill": {"athletics"}}, false)
	if body := w0.Body.String(); w0.Code != http.StatusUnprocessableEntity || errorAt(body, views.FieldSkill) || !strings.Contains(body, ">choose none<") {
		t.Errorf("skill on a class without skills: %d, message not at the top", w0.Code)
	}
	for _, tc := range []struct {
		skills []string
		code   int
	}{
		{[]string{"athletics", "insight"}, http.StatusSeeOther},
		{[]string{"athletics"}, http.StatusUnprocessableEntity},
		{[]string{"athletics", "athletics"}, http.StatusUnprocessableEntity},
		{[]string{"athletics", "perception"}, http.StatusUnprocessableEntity},
		{[]string{"athletics", "insight", "stealth"}, http.StatusUnprocessableEntity},
	} {
		e := newTestEnv(t)
		if err := e.q().UpsertCompendiumEntry(t.Context(), class); err != nil {
			t.Fatal(err)
		}
		form := url.Values{"pick": {"testfighter"}, "skill": tc.skills}
		path, player := fmt.Sprintf("/characters/%d/create/class", e.characterID), e.token(t, enum.RolePlayer)
		w := e.do(t, http.MethodPost, path, player, form, false)
		if w.Code != tc.code {
			t.Errorf("%v: %d, want %d", tc.skills, w.Code, tc.code)
		}
		if lost := lostPicks(w.Body.String(), form); tc.code == http.StatusUnprocessableEntity && (len(lost) > 0 || !errorAt(w.Body.String(), views.FieldSkill)) {
			t.Errorf("%v: refused page lost %v or its error is not at the skills", tc.skills, lost)
		}
		// Chosen already, the step changes nothing and asks for no skills.
		if w := e.do(t, http.MethodPost, path, player, url.Values{"pick": {"testfighter"}}, false); tc.code == http.StatusSeeOther && w.Code != http.StatusSeeOther {
			t.Errorf("class again: %d", w.Code)
		}
		if tc.code != http.StatusSeeOther {
			continue
		}
		// A crafted spell on a non-caster's choices, which draw no checklist,
		// is refused at the top.
		loadFeatFixtures(t, e)
		base := fmt.Sprintf("/characters/%d/create/", e.characterID)
		for _, step := range []string{"species=human", "background=testfolk"} {
			step, pick, _ := strings.Cut(step, "=")
			if w := e.do(t, http.MethodPost, base+step, player, url.Values{"pick": {pick}}, false); w.Code != http.StatusSeeOther {
				t.Fatalf("%s: %d %s", step, w.Code, w.Body)
			}
		}
		form = url.Values{"language": {"Elvish", "Dwarvish"}, "trait_versatile": {"Alert"}, "spell": {"light"}}
		w = e.do(t, http.MethodPost, base+"choices", player, form, false)
		if body := w.Body.String(); w.Code != http.StatusUnprocessableEntity || errorAt(body, views.FieldSpell) || !strings.Contains(body, "choose spells from the list") || len(lostPicks(body, form)) > 0 {
			t.Errorf("spell on a non-caster: %d, message not at the top or picks lost", w.Code)
		}
	}
}

var (
	reInput   = regexp.MustCompile(`<input[^>]*>`)
	reChecked = regexp.MustCompile(`\schecked[\s/>]`)
	// reHiddenRow is a spell row hidden because a picked lineage teaches it.
	reHiddenRow = regexp.MustCompile(`(?s)data-hide-for="[^"]*" hidden>(.*?)</div>`)
	// reDisabled is a disabled attribute (not a disabled: class variant).
	reDisabled = regexp.MustCompile(`\s(aria-|data-)?disabled[\s=/>]`)
)

// errorAt reports whether a refused page shows its error at the field key
// instead of the top.
// hasError reports whether body shows want: "@key: msg" is msg in the error
// drawn at key (legends say "choose …" too), anything else is in the page.
func hasError(body, want string) bool {
	key, msg, ok := strings.Cut(strings.TrimPrefix(want, "@"), ": ")
	if !strings.HasPrefix(want, "@") || !ok {
		return strings.Contains(body, want)
	}
	re := regexp.MustCompile(`<div id="error-` + regexp.QuoteMeta(html.EscapeString(key)) + `"[^>]*>` + regexp.QuoteMeta(html.EscapeString(msg)) + `</div>`)
	return re.MatchString(body)
}

func errorAt(body, key string) bool {
	k := html.EscapeString(key)
	return strings.Contains(body, `data-error-for="`+k+`"`) &&
		strings.Contains(body, `id="error-`+k+`"`) &&
		strings.Contains(body, `aria-describedby="error-`+k+`"`)
}

// lostPicks are the posted values a refused page offers but does not show
// again (a box left unticked, a select or field holding another value), or
// whose own choices it keeps hidden. Attribute order does not matter.
func lostPicks(body string, form url.Values) []string {
	tags := reInput.FindAllString(body, -1)
	var lost []string
	for k, vs := range form {
		name := ` name="` + html.EscapeString(k) + `"`
		for _, v := range vs {
			value := ` value="` + html.EscapeString(v) + `"`
			offered, kept := false, false
			for _, tag := range tags {
				if !strings.Contains(tag, name) {
					continue
				}
				box, has := strings.Contains(tag, `type="checkbox"`) || strings.Contains(tag, `type="radio"`), strings.Contains(tag, value)
				offered = offered || has || !box
				kept = kept || has && (!box || reChecked.MatchString(tag))
			}
			if offered && !kept {
				lost = append(lost, k+"="+v)
			}
			if strings.Contains(body, `data-show-for="`+html.EscapeString(k+"="+v)+`" hidden`) {
				lost = append(lost, k+"="+v+" (its choices)")
			}
		}
	}
	return lost
}

// withSpells ticks the first offered spells the form lacks (also by name, as
// a feat's pick) until it has as many cantrips and level 1+ spells as need
// asks, so a test fails only on its own picks.
func withSpells(form url.Values, offered []views.Option, need views.SpellNeed) url.Values {
	var named []string
	for _, vs := range form {
		named = append(named, vs...)
	}
	n := [2]int{need.Cantrips, need.Spells}
	for _, o := range offered {
		if slices.Contains(form["spell"], o.Slug) {
			n[min(o.Level, 1)]--
		}
	}
	for _, o := range offered {
		if k := min(o.Level, 1); n[k] > 0 && !slices.Contains(form["spell"], o.Slug) && !slices.Contains(named, o.Name) {
			form.Add("spell", o.Slug)
			n[k]--
		}
	}
	return form
}

// wizardLevelUp is form with the spells a Wizard level-up asks for.
func wizardLevelUp(t *testing.T, e testEnv, form url.Values) url.Values {
	t.Helper()
	ch, err := e.q().GetCharacter(t.Context(), e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	e.allowEntrySources(t)
	v, err := e.app.levelUpView(t.Context(), ch, "wizard", false)
	if err != nil {
		t.Fatal(err)
	}
	return withSpells(form, v.Spells, v.SpellNeed)
}

// creationSpells is form with the spells the choices step asks for.
func creationSpells(t *testing.T, e testEnv, form url.Values) url.Values {
	t.Helper()
	ch, err := e.q().GetCharacter(t.Context(), e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	sh, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	cr, err := e.app.creationChoices(t.Context(), sh, ch.CampaignID)
	if err != nil {
		t.Fatal(err)
	}
	return withSpells(form, cr.view.Spells, cr.view.SpellNeed)
}

// TestReadPicks: a group takes exactly its count, or all of fewer options; a
// group offering "No change" may stay empty.
func TestReadPicks(t *testing.T) {
	opts := func(names ...string) []compendium.NameDesc {
		var out []compendium.NameDesc
		for _, n := range names {
			out = append(out, compendium.NameDesc{Name: n})
		}
		return out
	}
	for _, tc := range []struct {
		name  string
		group views.ChoiceGroup
		form  url.Values
		ok    bool
	}{
		{"full count", views.ChoiceGroup{Key: "k", Count: 2, Options: opts("A", "B", "C")}, url.Values{"k": {"A", "B"}}, true},
		{"too few", views.ChoiceGroup{Key: "k", Count: 2, Options: opts("A", "B", "C")}, url.Values{"k": {"A"}}, false},
		{"a pick twice", views.ChoiceGroup{Key: "k", Count: 2, Options: opts("A", "B", "C")}, url.Values{"k": {"A", "A"}}, false},
		{"fewer options than the count", views.ChoiceGroup{Key: "k", Count: 2, Options: opts("A")}, url.Values{"k": {"A"}}, true},
		{"no options", views.ChoiceGroup{Key: "k", Count: 1}, nil, true},
		{"No change left empty", views.ChoiceGroup{Key: "k", Count: 1, Options: opts(keepSpells, "A")}, nil, true},
		{"empty without No change", views.ChoiceGroup{Key: "k", Count: 1, Options: opts("A")}, nil, false},
	} {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", strings.NewReader(tc.form.Encode()))
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if _, err := readPicks(c, []choiceGroup{{tc.group, nil, nil}}); (err == nil) != tc.ok {
			t.Errorf("%s: %v", tc.name, err)
		}
	}
	// A class level's and a feat's mastery groups may not pick one kind twice.
	feat := "feat_wm_" + views.FieldMastery
	for _, tc := range []struct {
		feat string
		ok   bool
	}{{"Mace", false}, {"Club", true}} {
		form := url.Values{views.FieldMastery: {"Mace"}, feat: {tc.feat}}
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Request = httptest.NewRequestWithContext(t.Context(), http.MethodPost, "/", strings.NewReader(form.Encode()))
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		groups := []choiceGroup{
			{views.ChoiceGroup{Key: views.FieldMastery, Count: 1, Options: opts("Mace", "Club")}, nil, nil},
			{views.ChoiceGroup{Key: feat, Count: 1, Options: opts("Mace", "Club")}, nil, nil},
		}
		if _, err := readPicks(c, groups); (err == nil) != tc.ok {
			t.Errorf("feat picks %s: %v", tc.feat, err)
		}
	}
}

// TestCreationOptionCheck: creation refuses an invocation whose prerequisite
// the finished step does not meet, like a level-up. Shadow Cloak, a level 1
// option needing Armor of Shadows, is offered next to it but not alone.
func TestCreationOptionCheck(t *testing.T) {
	cloak := testlock
	cloak.DataJson = strings.Replace(testlock.DataJson, "### Shadow Step", `### Shadow Cloak\n\n*Prerequisite: Armor of Shadows Invocation*\n\nCloak.\n\n### Shadow Step`, 1)
	const pick = "class_eldritch-invocation"
	for _, tc := range []struct {
		name string
		form url.Values
		code int
		err  string
	}{
		{"prerequisite held", url.Values{pick: {"Armor of Shadows"}}, http.StatusSeeOther, ""},
		{"prerequisite missing", url.Values{pick: {"Shadow Cloak"}}, http.StatusUnprocessableEntity, "Shadow Cloak needs Armor of Shadows Invocation"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			e := newTestEnv(t)
			loadFeatFixtures(t, e)
			if err := e.q().UpsertCompendiumEntry(t.Context(), cloak); err != nil {
				t.Fatal(err)
			}
			player := e.token(t, enum.RolePlayer)
			base := fmt.Sprintf("/characters/%d/create/", e.characterID)
			for _, step := range []string{"class=testlock", "species=human", "background=testfolk"} {
				step, pick, _ := strings.Cut(step, "=")
				if w := e.do(t, http.MethodPost, base+step, player, url.Values{"pick": {pick}}, false); w.Code != http.StatusSeeOther {
					t.Fatalf("%s: %d %s", step, w.Code, w.Body)
				}
			}
			form := url.Values{"language": {"Elvish", "Dwarvish"}, "trait_versatile": {"Alert"}}
			maps.Copy(form, tc.form)
			if w := e.do(t, http.MethodPost, base+"choices", player, form, false); w.Code != tc.code || !strings.Contains(w.Body.String(), tc.err) {
				t.Errorf("choices: %d, want %d with %q", w.Code, tc.code, tc.err)
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
	// At level 2 the column grows by 2; the level 2 options whose prerequisite
	// holds open up (not Eldritch Spear: no cantrip) and the one taken is gone.
	info, err := classInfo(sqlc.CompendiumEntry{Slug: testlock.Slug, Name: testlock.Name, DataJson: testlock.DataJson})
	if err != nil {
		t.Fatal(err)
	}
	names := func(g choiceGroup) []string {
		var out []string
		for _, o := range g.Options {
			out = append(out, o.Name)
		}
		return out
	}
	if gs := classOptionGroups(info, 2, 1, sh, nil); len(gs) != 2 || gs[0].Count != 2 || !slices.Equal(names(gs[0]), []string{"Agonizing Blast", "Shadow Step"}) {
		t.Errorf("level 2: %+v", groupViews(gs))
	}
	// Without Armor of Shadows, Shadow Step is offered next to it: picked together, they hold.
	if gs := classOptionGroups(info, 2, 1, dnd5e.NewSheet(), nil); len(gs) != 1 || !slices.Equal(names(gs[0]), []string{"Agonizing Blast", "Armor of Shadows", "Shadow Step"}) {
		t.Errorf("level 2, same group: %+v", groupViews(gs))
	}
	// A Repeatable option stays on offer and is added again under a new name.
	for range 2 {
		gs := classOptionGroups(info, 2, 1, sh, nil)
		if len(gs) == 0 || !slices.Contains(names(gs[0]), "Agonizing Blast") {
			t.Fatalf("repeatable: %+v", groupViews(gs))
		}
		gs[0].apply(&sh, []string{"Agonizing Blast"})
	}
	for _, name := range []string{"Eldritch Invocation: Agonizing Blast", "Eldritch Invocation: Agonizing Blast (2)"} {
		if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == name }) {
			t.Errorf("no %q in %+v", name, sh.Features)
		}
	}
	// An Eldritch Spear held without its cantrip (taken before the check) does
	// not make Armor of Shadows look needed: both stay replaceable.
	spearDesc := compendium.Options(info.Features[1].Description)[2].Desc
	old := dnd5e.Sheet{Features: []dnd5e.Feature{{Name: "Eldritch Invocation: Armor of Shadows", Description: "Mage Armor."}, {Name: "Eldritch Invocation: Eldritch Spear", Description: spearDesc}}}
	gs := classOptionGroups(info, 3, 2, old, nil)
	if i := slices.IndexFunc(gs, func(g choiceGroup) bool { return g.Key == "replace_eldritch-invocation" }); i < 0 || !slices.Equal(names(gs[i]), []string{keepSpells, "Armor of Shadows", "Eldritch Spear"}) {
		t.Errorf("swap with a stale option: %+v", groupViews(gs))
	}
}

// testlock is a pact caster with invocations: one at level 1, three at 2.
// Shadow Step needs Armor of Shadows; Eldritch Spear a damaging cantrip
// (testBlast).
var testlock = sqlc.UpsertCompendiumEntryParams{Kind: "class", Slug: "testlock", Name: "Testlock", SourceKey: "test", ContentHash: "x", DataJson: `{"caster":"pact","hit_die":8,
	"table":{"Cantrips":{"1":"0","2":"1"},"Eldritch Invocations":{"1":"1","2":"3"}},
	"features":[{"name":"Eldritch Invocations","desc":"You gain one invocation. Whenever you gain a Testlock level, you can replace one of your invocations with another one for which you qualify.","levels":[1]},
		{"name":"Eldritch Invocation Options","desc":"In order.\n\n### Agonizing Blast\n\n*Prerequisite: Level 2+ Warlock*\n\n**Repeatable.** More damage.\n\n### Armor of Shadows\n\nMage Armor.\n\n### Eldritch Spear\n\n*Prerequisite: Level 2+ Warlock, a Testlock Cantrip That Deals Damage*\n\nChoose one of your known Testlock cantrips that deals damage and has a range of 10+ feet.\n\n### Shadow Step\n\n*Prerequisite: Level 2+ Warlock, Armor of Shadows Invocation*\n\nTeleport."}]}`}

var testBlast = sqlc.UpsertCompendiumEntryParams{Kind: "spell", Slug: "test-blast", Name: "Test Blast", SourceKey: "test", ContentHash: "x",
	DataJson: `{"level":0,"classes":["testlock"],"range":"120 feet","desc":"Make a ranged spell attack. On a hit, the target takes 1d10 Force damage."}`}

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
	d, err := compendium.Decode[compendium.TextData](enum.KindSpecies, species.DataJson)
	if err != nil {
		t.Fatal(err)
	}
	sh := dnd5e.NewSheet()
	applySpecies(&sh, "Testborn", d)
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
		entry := sqlc.UpsertCompendiumEntryParams{Kind: enum.KindItem, Slug: slug, Name: slug, DataJson: "{}", SourceKey: "test", ContentHash: "x"}
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

// TestFeatRaiseText: every form the 5etools import writes a feat's ability
// increase in offers those abilities, below the maximum.
func TestFeatRaiseText(t *testing.T) {
	e := newTestEnv(t)
	sh := dnd5e.Sheet{Scores: map[string]int{"str": 20, "dex": 14, "con": 12, "int": 10, "wis": 10, "cha": 10}}
	for _, tt := range []struct {
		text string
		want []string
	}{
		{"Increase your Strength or Dexterity score by 1, to a maximum of 20.", []string{"Dexterity"}},
		{"Increase your Strength, Dexterity or Constitution score by 1, to a maximum of 20.", []string{"Dexterity", "Constitution"}},
		{"Increase your Constitution score by 1, to a maximum of 20.", []string{"Constitution"}},
		{"Increase your Strength score by 1, to a maximum of 20.", nil}, // at the maximum: no pick, the feat is still taken
		{"Increase one ability score of your choice by 1, to a maximum of 30.", []string{"Strength", "Dexterity", "Constitution", "Intelligence", "Wisdom", "Charisma"}},
	} {
		td := &compendium.TextData{Parts: []compendium.NameDesc{{Name: "Ability Score Increase", Desc: tt.text}}}
		groups, err := e.app.featChoices(t.Context(), "feat_x", "X", td, "", sh, 0)
		if err != nil || len(groups) != min(len(tt.want), 1) {
			t.Fatalf("%q: groups %+v %v", tt.text, groups, err)
		}
		if tt.want == nil {
			continue
		}
		var got []string
		for _, o := range groups[0].Options {
			got = append(got, o.Name)
		}
		if !slices.Equal(got, tt.want) {
			t.Errorf("%q: options %v, want %v", tt.text, got, tt.want)
		}
	}
}

// TestWizardSpeciesVariant: a species' subraces from any source are offered
// as variants, the preview follows the variant, and saving one stores it.
func TestWizardSpeciesVariant(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	for _, entry := range []sqlc.UpsertCompendiumEntryParams{
		{Kind: "species", Slug: "half-elf", Name: "Half-Elf", SourceKey: "book-a", DataJson: `{"desc":"","parts":[{"name":"Speed","desc":"30 feet","type":"speed"},{"name":"Skill Versatility","desc":"You gain two skills."}]}`},
		{Kind: "species", Slug: "half-elf-drow-descent", Name: "Half-Elf (Drow Descent)", SourceKey: "book-a", DataJson: `{"desc":"","subrace_of":"half-elf","parts":[{"name":"Speed","desc":"30 feet","type":"speed"},{"name":"Drow Magic","desc":"You know Dancing Lights."}]}`},
		{Kind: "species", Slug: "half-elf-aquatic-elf-descent", Name: "Half-Elf (Aquatic Elf Descent)", SourceKey: "book-b", DataJson: `{"desc":"","subrace_of":"half-elf","parts":[{"name":"Speed","desc":"35 feet","type":"speed"},{"name":"Swim","desc":"You swim."}]}`},
	} {
		entry.ContentHash = "x"
		if err := e.q().UpsertCompendiumEntry(ctx, entry); err != nil {
			t.Fatal(err)
		}
	}
	e.allowEntrySources(t)
	player := e.token(t, enum.RolePlayer)
	base := fmt.Sprintf("/characters/%d/create/species", e.characterID)
	page := e.do(t, http.MethodGet, base, player, nil, false).Body.String()
	if strings.Contains(page, "Drow Descent") {
		t.Error("a subrace is offered as a species of its own")
	}
	for _, tt := range []struct {
		query string
		want  []string
		not   []string
	}{
		{"?pick=half-elf", []string{"Standard", "Half-Elf (Drow Descent) · book-a", "Half-Elf (Aquatic Elf Descent) · book-b", "Skill Versatility"}, []string{"You swim."}},
		{"?pick=half-elf&variant=half-elf-aquatic-elf-descent", []string{"You swim.", `name="pick" value="half-elf-aquatic-elf-descent"`}, []string{"Skill Versatility"}},
		// A variant of another species is ignored.
		{"?pick=half-elf&variant=elf-high", []string{"Skill Versatility", `name="pick" value="half-elf"`}, nil},
	} {
		body := html.UnescapeString(e.do(t, http.MethodGet, base+tt.query, player, nil, false).Body.String())
		for _, s := range tt.want {
			if !strings.Contains(body, s) {
				t.Errorf("%s: no %q", tt.query, s)
			}
		}
		for _, s := range tt.not {
			if strings.Contains(body, s) {
				t.Errorf("%s: shows %q", tt.query, s)
			}
		}
	}
	e.do(t, http.MethodPost, base, player, url.Values{"pick": {"half-elf-aquatic-elf-descent"}}, false)
	row, err := e.q().GetCharacter(ctx, e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	sh, err := dnd5e.ParseSheet(row.SheetJson)
	if err != nil {
		t.Fatal(err)
	}
	if sh.Species != "Half-Elf (Aquatic Elf Descent)" || sh.SpeciesSource != "book-b" || sh.Speed != 35 ||
		!slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == "Swim" }) {
		t.Errorf("saved species %q from %q, speed %d, features %v", sh.Species, sh.SpeciesSource, sh.Speed, sh.Features)
	}
}
