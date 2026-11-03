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
