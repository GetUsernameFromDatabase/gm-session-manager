package web

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
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
		{Kind: "species", Slug: "goliath", Name: "Goliath", DataJson: `{"desc":"","parts":[{"name":"Speed","desc":"35 feet","type":"speed"},{"name":"Powerful Build","desc":"You count as one size larger when carrying."}]}`},
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
	for _, step := range []struct {
		step string
		form url.Values
		next string
	}{
		{"class", url.Values{"pick": {"wizard"}}, "/species"},
		{"species", url.Values{"pick": {"goliath"}}, "/abilities"},
		{"species", url.Values{"pick": {"goliath"}}, "/abilities"}, // once only
		{"abilities", scores, "/background"},
		{"background", url.Values{"pick": {"sage"}, "mode": {"two"}, "plus2": {"int"}, "plus1": {"con"}}, "/equipment"},
		{"equipment", url.Values{"option_background": {"B"}}, "/details"},
		{"details", url.Values{"name": {"Wren"}, "backstory": {"Raised in a library."}}, "/review"},
		{"review", nil, fmt.Sprintf("/characters/%d", e.characterID)},
	} {
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
	if sh.Species != "Goliath" || sh.Speed != 35 || sh.Background != "Sage" || sh.Scores["int"] != 17 || sh.Scores["con"] != 15 || sh.Coins.GP != 50 {
		t.Errorf("species %q speed %d background %q int %d con %d gp %d", sh.Species, sh.Speed, sh.Background, sh.Scores["int"], sh.Scores["con"], sh.Coins.GP)
	}
	for _, f := range []string{"Powerful Build", "Magic Initiate (Wizard)", "Arcane Recovery"} {
		if !strings.Contains(row.SheetJson, `"name":"`+f+`"`) {
			t.Errorf("missing feature %s: %s", f, row.SheetJson)
		}
	}
	if strings.Count(row.SheetJson, `"name":"Powerful Build"`) != 1 || !sh.SetUp("review") || !strings.Contains(strings.Join(sh.SkillProfs, ","), "arcana") {
		t.Errorf("sheet after the wizard: %s", row.SheetJson)
	}
}
