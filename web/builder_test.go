package web

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
)

// TestBuilderInventory covers the coin form (one change per coin, as a
// difference from the amount shown) and setting an item's count.
func TestBuilderInventory(t *testing.T) {
	e := newTestEnv(t)
	player := e.token(t, auth.RolePlayer)
	cmdPath := fmt.Sprintf("/characters/%d/cmd", e.characterID)
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
	base := fmt.Sprintf("/characters/%d/builder?verb=", e.characterID)
	for path, want := range map[string]string{
		"coins&op=coin-edit":                          `name="coin_gp"`,
		"items&op=item-edit&subject=Rope+%2850+ft%29": `value="3"`,
	} {
		if w := e.do(t, http.MethodGet, base+path, player, nil, true); !strings.Contains(w.Body.String(), want) {
			t.Errorf("%s lacks %q: %s", path, want, w.Body)
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
	if err := e.q().UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{
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
	player := e.token(t, auth.RolePlayer)
	id := e.characterID
	for _, tt := range []struct{ path, want string }{
		{fmt.Sprintf("/characters/%d/picks?pick=weapon&q=rapi", id), fmt.Sprintf("entry=%d", rapier.ID)},
		{fmt.Sprintf("/characters/%d/picks?pick=feature&q=arcane", id), "Arcane Recovery"},
		{fmt.Sprintf("/characters/%d/builder?verb=weapons&op=weapon-add&entry=%d", id, rapier.ID), `value="1d8"`},
		{fmt.Sprintf("/characters/%d/builder?verb=weapons&op=weapon-add&entry=%d", id, rapier.ID), `value="finesse"`},
		{fmt.Sprintf("/characters/%d/builder?verb=features&op=feature-add&entry=%d&feature=Arcane+Recovery", id, wizard.ID), `value="Arcane Recovery"`},
		{fmt.Sprintf("/characters/%d/builder?verb=details&op=set-species", id), "data-combobox-create-input"},
		{fmt.Sprintf("/characters/%d/picks?pick=feature&format=combobox&q=arcane", id), fmt.Sprintf("entry=%d&amp;feature=Arcane+Recovery", wizard.ID)},
		{fmt.Sprintf("/characters/%d/builder?verb=features&op=feature-add", id), "pick=feature&amp;format=combobox"},
		{fmt.Sprintf("/characters/%d/builder?verb=features&op=feature-remove", id), "data-reload="},
		{"/compendium/search?kind=item&format=combobox&q=rap", `data-templ-value="Rapier"`},
		{"/compendium/lookup?kind=spell&slug=fire-bolt", "mote of fire"},
		{fmt.Sprintf("/characters/%d/complete?words=spell+add&slot=%%3Cslug%%3E&q=fire", id), `"fire-bolt"`},
		{fmt.Sprintf("/characters/%d/complete?words=toggle&slot=%%3Ccondition%%3E&q=PR", id), `["prone"]`},
	} {
		w := e.do(t, http.MethodGet, tt.path, player, nil, true)
		if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), tt.want) {
			t.Errorf("%s: status %d, want %q in %s", tt.path, w.Code, tt.want, w.Body)
		}
	}
	// The coin tapped on the sheet is the one the coin form focuses.
	w := e.do(t, http.MethodGet, fmt.Sprintf("/characters/%d/builder?verb=coins&op=coin-edit&subject=sp", id), player, nil, true)
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
	} {
		if got := compendiumWeapon("w", &compendium.ItemData{Damage: tt.damage, Properties: tt.props}); got != tt.want {
			t.Errorf("%s %s: %+v", tt.damage, tt.props, got)
		}
	}
}
