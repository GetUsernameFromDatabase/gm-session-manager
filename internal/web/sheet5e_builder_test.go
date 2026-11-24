package web

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/route"
)

// TestBuilderInventory covers the coin form (one change per coin, as a
// difference from the amount shown) and setting an item's count.
func TestBuilderInventory(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, enum.RolePlayer)
	cmdPath := route.Path(route.CharacterCmd, e.characterID)
	coins := func(gp, wasGP string) url.Values {
		f := url.Values{}
		for _, k := range []string{"pp", "ep", "sp", "cp"} {
			f.Set("coin_"+k, "0")
			f.Set("was_"+k, "0")
		}
		f.Set("coin_gp", gp)
		f.Set("was_gp", wasGP)
		return f
	}
	for _, form := range []url.Values{
		coins("12", "0"),
		// A second device still showing 0 adds its own 5 on top.
		coins("5", "0"),
		// Nothing changed: accepted, nothing happens.
		coins("0", "0"),
		{"prefix": {"item add"}, "subject": {"Rope (50 ft)"}, "sign": {""}, "n": {"1"}},
		{"prefix": {"item set"}, "subject": {"Rope (50 ft)"}, "sign": {""}, "n": {"3"}},
	} {
		if w := e.do(t, http.MethodPost, cmdPath, player, form, true); w.Code != http.StatusOK {
			t.Fatalf("%v: status %d: %s", form, w.Code, w.Body)
		}
	}
	if w := e.do(t, http.MethodPost, cmdPath, player, coins("-30", "0"), true); w.Code != http.StatusUnprocessableEntity {
		t.Errorf("overspending: status %d", w.Code)
	}
	ch, _ := e.q().GetCharacter(t.Context(), e.characterID)
	for _, want := range []string{`"gp":17`, `"name":"Rope (50 ft)","qty":3`} {
		if !strings.Contains(ch.SheetJson, want) {
			t.Errorf("sheet lacks %s: %s", want, ch.SheetJson)
		}
	}
	base := route.Path(route.CharacterBuilder, e.characterID) + "?verb="
	for path, want := range map[string]string{
		"coins&op=coin-edit":                          `name="coin_gp"`,
		"items&op=item-edit&subject=Rope+%2850+ft%29": `value="3"`,
	} {
		if w := e.do(t, http.MethodGet, base+path, player, nil, true); !strings.Contains(w.Body.String(), want) {
			t.Errorf("%s lacks %q: %s", path, want, w.Body)
		}
	}
}

// TestBuilderFeat: a feat picked in the feature look-up joins the sheet's
// feats pinned to its source, and removing the feature takes both off.
func TestBuilderFeat(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if err := e.upsertEntry(ctx, sqlc.UpsertCompendiumEntryParams{Kind: "feat", Slug: "test-alert", Name: "Test Alert", SourceKey: "test", ContentHash: "x", DataJson: `{"desc":"Add your Proficiency Bonus to Initiative."}`}); err != nil {
		t.Fatal(err)
	}
	feat, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "feat", Slug: "test-alert"})
	if err != nil {
		t.Fatal(err)
	}
	player := e.token(t, enum.RolePlayer)
	form := e.do(t, http.MethodGet, fmt.Sprintf("%s?verb=features&op=feature-add&entry=%d", route.Path(route.CharacterBuilder, e.characterID), feat.ID), player, nil, true).Body.String()
	if !strings.Contains(form, `name="feat_source" value="test"`) {
		t.Fatalf("form lacks the feat's source: %s", form)
	}
	cmd := route.Path(route.CharacterCmd, e.characterID)
	add := url.Values{"prefix": {"feature add"}, "subject": {"Test Alert"}, "source": {"Feat"}, "text": {"Initiative."}, "feat_source": {"test"}}
	if w := e.do(t, http.MethodPost, cmd, player, add, true); w.Code != http.StatusOK {
		t.Fatalf("add: %d %s", w.Code, w.Body)
	}
	if sh := e.sheet(t); !slices.Equal(sh.Feats, []string{"Test Alert"}) || sh.FeatSources["test-alert"] != "test" {
		t.Errorf("after add: feats %v, pins %v", sh.Feats, sh.FeatSources)
	}
	if w := e.do(t, http.MethodPost, cmd, player, url.Values{"cmd": {`feature remove "Test Alert"`}}, true); w.Code != http.StatusOK {
		t.Fatalf("remove: %d %s", w.Code, w.Body)
	}
	if sh := e.sheet(t); len(sh.Feats) != 0 || len(sh.FeatSources) != 0 || len(sh.Features) != 0 {
		t.Errorf("after remove: feats %v, pins %v, features %v", sh.Feats, sh.FeatSources, sh.Features)
	}
}

// A fixed use count must be a whole number of at least one.
func TestBuilderFeatUsesN(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, enum.RolePlayer)
	cmd := route.Path(route.CharacterCmd, e.characterID)
	for _, n := range []string{"-2", "x", "0"} {
		add := url.Values{"prefix": {"feature add"}, "subject": {"Surge"}, "uses": {"n"}, "uses_n": {n}, "reset": {"long"}}
		if w := e.do(t, http.MethodPost, cmd, player, add, true); w.Code != http.StatusUnprocessableEntity {
			t.Errorf("uses_n=%s: %d, want 422", n, w.Code)
		}
	}
}

// TestBuilderCompendium covers filling forms from the compendium: weapon
// and feature look-ups, suggestions and spell details.
func TestBuilderCompendium(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	if _, err := compendium.Import(ctx, e.app.Store, "../compendium/testdata"); err != nil {
		t.Fatal(err)
	}
	if err := e.upsertEntry(ctx, sqlc.UpsertCompendiumEntryParams{
		Kind: "item", Slug: "rapier", Name: "Rapier", SourceKey: "test", BodyText: "rapier",
		DataJson: `{"category":"weapon","damage":"1d8 piercing","properties":"Finesse, Vex","desc":""}`,
	}); err != nil {
		t.Fatal(err)
	}
	rapier, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "item", Slug: "rapier"})
	if err != nil {
		t.Fatal(err)
	}
	wizard, err := e.q().GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: "class", Slug: "wizard"})
	if err != nil {
		t.Fatal(err)
	}
	e.allowEntrySources(t)
	player := e.token(t, enum.RolePlayer)
	id := e.characterID
	for _, tt := range []struct{ path, want string }{
		{route.Path(route.CharacterPicks, id) + "?pick=weapon&q=rapi", fmt.Sprintf("entry=%d", rapier.ID)},
		{route.Path(route.CharacterPicks, id) + "?pick=feature&q=arcane", "Arcane Recovery"},
		{fmt.Sprintf("%s?verb=weapons&op=weapon-add&entry=%d", route.Path(route.CharacterBuilder, id), rapier.ID), `value="1d8"`},
		{fmt.Sprintf("%s?verb=weapons&op=weapon-add&entry=%d", route.Path(route.CharacterBuilder, id), rapier.ID), `value="finesse"`},
		{fmt.Sprintf("%s?verb=features&op=feature-add&entry=%d&feature=Arcane+Recovery", route.Path(route.CharacterBuilder, id), wizard.ID), `value="Arcane Recovery"`},
		{route.Path(route.CharacterBuilder, id) + "?verb=details&op=set-species", "data-combobox-create-input"},
		{route.Path(route.CharacterPicks, id) + "?pick=feature&format=combobox&q=arcane", fmt.Sprintf("entry=%d&amp;feature=Arcane+Recovery", wizard.ID)},
		{route.Path(route.CharacterBuilder, id) + "?verb=features&op=feature-add", "pick=feature&amp;format=combobox"},
		{route.Path(route.CharacterBuilder, id) + "?verb=features&op=feature-remove", "data-reload="},
		{route.Compendium + route.CompendiumSearch + "?kind=item&format=combobox&q=rap", `data-templ-value="Rapier"`},
		{route.Compendium + route.CompendiumLookup + "?kind=spell&slug=fire-bolt", "mote of fire"},
		{route.Path(route.CharacterComplete, id) + "?words=spell+add&slot=%3Cslug%3E&q=fire", `"fire-bolt"`},
		{route.Path(route.CharacterComplete, id) + "?words=toggle&slot=%3Ccondition%3E&q=PR", `["prone"]`},
	} {
		w := e.do(t, http.MethodGet, tt.path, player, nil, true)
		if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), tt.want) {
			t.Errorf("%s: status %d, want %q in %s", tt.path, w.Code, tt.want, w.Body)
		}
	}
	// The coin tapped on the sheet is the one the coin form focuses.
	w := e.do(t, http.MethodGet, route.Path(route.CharacterBuilder, id)+"?verb=coins&op=coin-edit&subject=sp", player, nil, true)
	for _, coin := range []string{"sp", "gp"} {
		tag := regexp.MustCompile(`<input[^>]*id="coin-` + coin + `"[^>]*>`).FindString(w.Body.String())
		if strings.Contains(tag, "autofocus") != (coin == "sp") {
			t.Errorf("coin %s input: %s", coin, tag)
		}
	}
}

func TestCompendiumWeapon(t *testing.T) {
	for _, tt := range []struct {
		damage, props string
		want          dnd5e.Weapon
	}{
		{"1d8 slashing", "Topple, Versatile", dnd5e.Weapon{Name: "w", Ability: "str", Dice: "1d8", DamageType: "slashing", Proficient: true, Properties: "Topple, Versatile"}},
		{"1d8 piercing", "Finesse, Vex", dnd5e.Weapon{Name: "w", Ability: "finesse", Dice: "1d8", DamageType: "piercing", Proficient: true, Properties: "Finesse, Vex"}},
		{"1 piercing", "Ammunition, Loading, Vex", dnd5e.Weapon{Name: "w", Ability: "dex", Dice: "1", DamageType: "piercing", Proficient: true, Properties: "Ammunition, Loading, Vex"}},
		{"1d8 slashing", "Versatile", dnd5e.Weapon{Name: "w", Ability: "str", Dice: "1d8", DamageType: "slashing", Proficient: true, Properties: "Versatile", Mastery: "Sap"}},
	} {
		if got := compendiumWeapon("w", &compendium.ItemData{Damage: tt.damage, Properties: tt.props, Mastery: tt.want.Mastery}); got != tt.want {
			t.Errorf("%s %s: %+v", tt.damage, tt.props, got)
		}
	}
}

// TestItemLookupByName checks an inventory item finds its compendium entry by
// name in any case, and a name the compendium lacks gets an empty answer.
func TestItemLookupByName(t *testing.T) {
	e := newTestEnv(t)
	if err := e.upsertEntry(t.Context(), sqlc.UpsertCompendiumEntryParams{
		Kind: "item", Slug: "rapier", Name: "Rapier", SourceKey: "test", BodyText: "rapier",
		DataJson: `{"category":"weapon","damage":"1d8 piercing","properties":"Finesse, Vex","desc":""}`,
	}); err != nil {
		t.Fatal(err)
	}
	player := e.token(t, enum.RolePlayer)
	for _, tt := range []struct{ name, want string }{{"rapier", "1d8 piercing"}, {"Lucky pebble", ""}} {
		w := e.do(t, http.MethodGet, route.Compendium+route.CompendiumLookup+"?kind=item&name="+url.QueryEscape(tt.name), player, nil, true)
		if body := w.Body.String(); w.Code != http.StatusOK || (tt.want == "") != (body == "") || !strings.Contains(body, tt.want) {
			t.Errorf("%s: %d %q, want %q", tt.name, w.Code, body, tt.want)
		}
	}
}

// TestLanguagesForm checks the languages tab lists each language with its
// own remove button above the add field, with no separate remove form.
func TestLanguagesForm(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, enum.RolePlayer)
	body := e.do(t, http.MethodGet, route.Path(route.CharacterBuilder, e.characterID)+"?verb=profs&op=lang-add", player, nil, true).Body.String()
	for _, want := range []string{`aria-label="Remove Common"`, "prof remove language", "Add a language"} {
		if !strings.Contains(body, want) {
			t.Errorf("missing %q", want)
		}
	}
	if strings.Contains(body, "Remove language") {
		t.Error("the separate remove form is still offered")
	}
}
