package web

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// seedMastery adds a class mastering 2 weapon kinds at level 1 and 3 at
// level 3, three weapons with a mastery and one without, and the Sap rule.
func seedMastery(t *testing.T, e testEnv) {
	t.Helper()
	e.source(t, "test", "2024")
	entry := func(kind, slug, name, data string) sqlc.UpsertCompendiumEntryParams {
		return sqlc.UpsertCompendiumEntryParams{Kind: kind, Slug: slug, Name: name, SourceKey: "test", ContentHash: "x", DataJson: data}
	}
	for _, p := range []sqlc.UpsertCompendiumEntryParams{
		entry("class", "testfighter", "Testfighter", `{"hit_die":10,"caster":"none","table":{"Weapon Mastery":{"1":"2","2":"2","3":"3"}}}`),
		entry("item", "longsword", "Longsword", `{"category":"weapon","damage":"1d8 slashing","mastery":"Sap","desc":""}`),
		entry("item", "dagger", "Dagger", `{"category":"weapon","damage":"1d4 piercing","mastery":"Nick","desc":""}`),
		entry("item", "club", "Club", `{"category":"weapon","damage":"1d4 bludgeoning","mastery":"Slow","desc":""}`),
		entry("item", "torch", "Torch", `{"category":"gear","desc":""}`),
		entry("rule", "sap", "Sap", `{"category":"Weapon Mastery","desc":"The target has Disadvantage."}`),
	} {
		if err := e.q().UpsertCompendiumEntry(t.Context(), p); err != nil {
			t.Fatal(err)
		}
	}
}

// TestWeaponMastery: level 1 and 3 ask for the kinds the class table adds,
// the commands keep to the limit and only the owner runs them, and the sheet
// shows each weapon's mastery and the count.
func TestWeaponMastery(t *testing.T) {
	e := newTestEnv(t)
	seedMastery(t, e)
	player, gm := e.token(t, auth.RolePlayer), e.token(t, auth.RoleAdmin)
	levelUp := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	for _, step := range []struct {
		form url.Values
		code int
		want []string // masteries after
	}{
		{url.Values{"class": {"testfighter"}, views.FieldMastery: {"Longsword"}}, http.StatusUnprocessableEntity, nil},
		{url.Values{"class": {"testfighter"}, views.FieldMastery: {"Longsword", "Torch"}}, http.StatusUnprocessableEntity, nil},
		{url.Values{"class": {"testfighter"}, views.FieldMastery: {"Longsword", "Dagger"}}, http.StatusSeeOther, []string{"Dagger", "Longsword"}},
		{url.Values{"class": {"testfighter"}}, http.StatusSeeOther, []string{"Dagger", "Longsword"}}, // level 2 adds none
		{url.Values{"class": {"testfighter"}, views.FieldMastery: {"Dagger"}}, http.StatusUnprocessableEntity, []string{"Dagger", "Longsword"}},
		{url.Values{"class": {"testfighter"}, views.FieldMastery: {"Club"}}, http.StatusSeeOther, []string{"Dagger", "Longsword", "Club"}},
	} {
		if w := e.do(t, http.MethodPost, levelUp, player, step.form, false); w.Code != step.code {
			t.Fatalf("%v: %d, want %d: %s", step.form, w.Code, step.code, w.Body)
		}
		if got := e.sheet(t).Masteries; !slices.Equal(got, step.want) {
			t.Fatalf("%v: masteries %v, want %v", step.form, got, step.want)
		}
	}

	cmdURL := fmt.Sprintf("/characters/%d/cmd", e.characterID)
	for _, tt := range []struct {
		name, cookie, cmd string
		code              int
	}{
		{"over the limit", player, "mastery add Torch", http.StatusUnprocessableEntity},           // no mastery: unknown
		{"the GM may not choose", gm, "mastery remove Longsword", http.StatusUnprocessableEntity}, // the flash says why
		{"remove", player, "mastery remove Club", http.StatusOK},
		{"add", player, "mastery add club", http.StatusOK},
		{"full", player, "mastery remove Dagger", http.StatusOK},
	} {
		if w := e.do(t, http.MethodPost, cmdURL, tt.cookie, url.Values{"cmd": {tt.cmd}}, true); w.Code != tt.code {
			t.Errorf("%s: %d, want %d: %s", tt.name, w.Code, tt.code, w.Body)
		}
	}
	if got := e.sheet(t).Masteries; !slices.Equal(got, []string{"Longsword", "Club"}) {
		t.Errorf("masteries %v", got)
	}
	if w := e.do(t, http.MethodPost, cmdURL, player, url.Values{"cmd": {"mastery add Dagger"}}, true); w.Code != http.StatusOK {
		t.Fatalf("add a third: %d %s", w.Code, w.Body)
	}
	if w := e.do(t, http.MethodPost, cmdURL, player, url.Values{"cmd": {"mastery add Torch"}}, true); w.Code != http.StatusUnprocessableEntity || !strings.Contains(w.Body.String(), "mastery property") {
		t.Errorf("unknown weapon: %d %s", w.Code, w.Body)
	}
	sh := e.sheet(t)
	sh.Masteries = sh.Masteries[:2]
	sh.Weapons = append(sh.Weapons, dnd5e.Weapon{Name: "+1 Longsword", Ability: "str", Dice: "1d8", Mastery: "Sap"}, dnd5e.Weapon{Name: "Spear", Ability: "str", Dice: "1d6", Mastery: "Sap"})
	e.setSheet(t, sh)
	page := e.do(t, http.MethodGet, fmt.Sprintf("/characters/%d", e.characterID), player, nil, false).Body.String()
	for _, want := range []string{"2/3 chosen", "/compendium/find/rule/sap", ">mastered<", "not mastered"} {
		if !strings.Contains(page, want) {
			t.Errorf("sheet lacks %q", want)
		}
	}
	form := e.do(t, http.MethodGet, fmt.Sprintf("/characters/%d/builder?verb=weapons&op=%s", e.characterID, views.OpMastery), player, nil, true).Body.String()
	if !strings.Contains(form, `mastery add \&#34;Dagger`) || !strings.Contains(form, `mastery remove \&#34;Club`) || strings.Contains(form, "Torch") {
		t.Errorf("mastery form: %s", form)
	}
}

// testKnight is a third-caster subclass of seedMastery's Testfighter.
var testKnight = sqlc.UpsertCompendiumEntryParams{Kind: "subclass", Slug: "test-knight", Name: "Test Knight", SourceKey: "test", ContentHash: "x", DataJson: `{"caster":"third","subclass_of":"testfighter"}`}

// TestLevelUpSubclassCaster: choosing a third-caster subclass gives a class
// without spellcasting its slots, and picking an already mastered kind adds
// it once.
func TestLevelUpSubclassCaster(t *testing.T) {
	e := newTestEnv(t)
	seedMastery(t, e)
	if err := e.q().UpsertCompendiumEntry(t.Context(), testKnight); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, auth.RolePlayer)
	levelUp := fmt.Sprintf("/characters/%d/levelup", e.characterID)
	for _, form := range []url.Values{
		{"class": {"testfighter"}, views.FieldMastery: {"Longsword", "Dagger"}},
		{"class": {"testfighter"}},
		{"class": {"testfighter"}, "subclass": {"Test Knight"}, views.FieldMastery: {"Club"}},
	} {
		if w := e.do(t, http.MethodPost, levelUp, player, form, false); w.Code != http.StatusSeeOther {
			t.Fatalf("%v: %d %s", form, w.Code, w.Body)
		}
	}
	sh := e.sheet(t)
	if cl := sh.Classes[0]; cl.Caster != dnd5e.CasterThird || cl.SubclassSlug != "test-knight" || sh.Slots[0].Max != 2 {
		t.Errorf("class %+v, slots %+v", cl, sh.Slots[0])
	}

	ch, err := e.q().GetCharacter(t.Context(), e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	info, _, err := classEntry(t.Context(), e.q(), "testfighter", "")
	if err != nil {
		t.Fatal(err)
	}
	sh.Masteries = []string{"Longsword"}
	groups, err := e.app.masteryGroups(t.Context(), sh, info, 3, ch.CampaignID)
	if err != nil || len(groups) != 1 {
		t.Fatalf("groups %v %v", groups, err)
	}
	groups[0].apply(&sh, []string{"longsword", "Club", "Club"})
	if !slices.Equal(sh.Masteries, []string{"Longsword", "Club"}) {
		t.Errorf("masteries %v", sh.Masteries)
	}
}

// TestMasteryMulticlass: a level in a second class adds its table value to
// the first class's (limits add up across classes).
func TestMasteryMulticlass(t *testing.T) {
	e := newTestEnv(t)
	seedMastery(t, e)
	if err := e.q().UpsertCompendiumEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{
		Kind: "class", Slug: "testrogue", Name: "Testrogue", SourceKey: "test", ContentHash: "x",
		DataJson: `{"hit_die":8,"caster":"none","table":{"Weapon Mastery":{"1":"1"}}}`,
	}); err != nil {
		t.Fatal(err)
	}
	ch, err := e.q().GetCharacter(t.Context(), e.characterID)
	if err != nil {
		t.Fatal(err)
	}
	info, _, err := classEntry(t.Context(), e.q(), "testfighter", "")
	if err != nil {
		t.Fatal(err)
	}
	sh := dnd5e.Sheet{Classes: []dnd5e.ClassLevel{{Slug: "testrogue", Level: 1}}, Masteries: []string{"Club"}}
	groups, err := e.app.masteryGroups(t.Context(), sh, info, 1, ch.CampaignID)
	// Rogue 1 (1) + Fighter 1 (2) = 3, one already chosen: the highest alone would offer 1.
	if err != nil || len(groups) != 1 || groups[0].Count != 2 {
		t.Fatalf("groups %+v %v, want one group of 2", groups, err)
	}
}

// TestCompendiumFind checks a slug link opens the entry, and the item page
// links its mastery that way.
func TestCompendiumFind(t *testing.T) {
	e := newTestEnv(t)
	seedMastery(t, e)
	player := e.token(t, auth.RolePlayer)
	rule, _ := e.q().GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "rule", Slug: "sap"})
	w := e.do(t, http.MethodGet, "/compendium/find/rule/sap", player, nil, false)
	if w.Code != http.StatusSeeOther || w.Header().Get("Location") != fmt.Sprintf("/compendium/entries/%d", rule.ID) {
		t.Errorf("find: %d %s", w.Code, w.Header().Get("Location"))
	}
	if w := e.do(t, http.MethodGet, "/compendium/find/rule/vex", player, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("missing rule: %d", w.Code)
	}
	sword, _ := e.q().GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "item", Slug: "longsword"})
	if page := e.do(t, http.MethodGet, fmt.Sprintf("/compendium/entries/%d", sword.ID), player, nil, false).Body.String(); !strings.Contains(page, `href="/compendium/find/rule/sap"`) {
		t.Error("item page does not link its mastery")
	}
}

// TestLiveMastery checks choosing a mastery refetches the sheet: it is part
// of the sheet's shape, not a bound value.
func TestLiveMastery(t *testing.T) {
	e := newTestEnv(t)
	seedMastery(t, e)
	sh := e.sheet(t)
	sh.Classes = []dnd5e.ClassLevel{{Slug: "testfighter", Name: "Testfighter", Level: 1, HitDie: 10, Caster: dnd5e.CasterNone}}
	e.setSheet(t, sh)
	view := views.LiveKey(views.LiveCharacter, e.characterID)
	hash := func() any {
		st, ok, err := e.app.State(t.Context(), hub.NewClient(e.campaignID, perm.Player(e.characterID), 0, "test"), view)
		if err != nil || !ok {
			t.Fatal(ok, err)
		}
		return st[hub.HashKey]
	}
	before := hash()
	if w := e.do(t, http.MethodPost, fmt.Sprintf("/characters/%d/cmd", e.characterID), e.token(t, auth.RolePlayer), url.Values{"cmd": {"mastery add Longsword"}}, true); w.Code != http.StatusOK {
		t.Fatalf("%d %s", w.Code, w.Body)
	}
	if hash() == before {
		t.Error("a new mastery does not change the sheet's hash")
	}
}
