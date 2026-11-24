package views

import (
	"regexp"
	"strings"
	"testing"

	"gm-session-manager/internal/dnd5e"
)

// TestModTargetKeepsStored: a modifier on a target the sheet no longer lists
// (a renamed or merged weapon) stays selected in both editors.
func TestModTargetKeepsStored(t *testing.T) {
	s := SheetView{C: dnd5e.Character{Sheet: dnd5e.Sheet{
		Weapons:   []dnd5e.Weapon{{Name: "Dagger 1"}},
		Modifiers: []dnd5e.CustomMod{{Target: "attack.Daggers 4", Value: 1, Note: "Magic", Source: "Blessed"}},
		Features:  []dnd5e.Feature{{Name: "Blessed"}},
	}}}
	for _, tt := range []struct{ name, verb, op, subject, field string }{
		{"modifier dialog", "mods", "mod-edit", "0", "subject"},
		{"feature editor", "features", "feature-add", "Blessed", "mod_target"},
	} {
		var b strings.Builder
		if err := CommandBuilder(s, tt.verb, tt.op, tt.subject, BuilderExtra{}).Render(t.Context(), &b); err != nil {
			t.Fatal(err)
		}
		html := b.String()
		re := regexp.MustCompile(`<input[^>]*name="` + tt.field + `"[^>]*value="attack\.Daggers 4"`)
		if !re.MatchString(html) {
			t.Errorf("%s: select does not carry the stored target in %s", tt.name, html)
		}
		if !strings.Contains(html, "Daggers 4 to hit") {
			t.Errorf("%s: no label for the stored target", tt.name)
		}
	}
}

// TestBuilderEveryOp renders every builder form of a sheet that offers them
// all, and checks each posts its own command prefix: a form that crashes, or
// an op the builder no longer finds (it falls back to the verb's first), fails.
func TestBuilderEveryOp(t *testing.T) {
	sh := dnd5e.Sheet{
		Slots:     [9]dnd5e.Slot{{Max: 2}},
		Pact:      dnd5e.Slot{Max: 1},
		Spells:    []dnd5e.Spell{{Slug: "bless", Name: "Bless", Level: 1}},
		Inventory: []dnd5e.Item{{Name: "Rope", Qty: 1}},
		Weapons:   []dnd5e.Weapon{{Name: "Dagger", Dice: "1d4"}},
		Modifiers: []dnd5e.CustomMod{{Target: dnd5e.ModAC, Value: 1, Note: "Ring"}},
		Features:  []dnd5e.Feature{{Name: "Darkvision"}},
		// Dead, so Revive is offered too.
		DeathSaves: dnd5e.DeathSaves{Failures: dnd5e.MaxDeathSaves},
	}
	s := SheetView{CanSheet: true, C: dnd5e.Character{Sheet: sh, MaxHP: 10}}
	// Forms that edit one thing need it named.
	subjects := map[string]string{focusSkill: "stealth", focusSave: "con", modEdit: "0"}
	for _, v := range s.builderVerbs() {
		for _, o := range v.Ops {
			var b strings.Builder
			if err := CommandBuilder(s, v.Key, o.Key, subjects[o.Key], BuilderExtra{}).Render(t.Context(), &b); err != nil {
				t.Fatalf("%s/%s: %v", v.Key, o.Key, err)
			}
			if o.Prefix != "" && !strings.Contains(b.String(), `name="prefix" value="`+o.Prefix+`"`) {
				t.Errorf("%s/%s: no prefix %q in %s", v.Key, o.Key, o.Prefix, b.String())
			}
		}
	}
}

// TestModTargetStowedWeapon: a stowed weapon's "to hit" link opens on that
// weapon, not on the first target (Armor class).
func TestModTargetStowedWeapon(t *testing.T) {
	s := SheetView{C: dnd5e.Character{Sheet: dnd5e.Sheet{
		Inventory: []dnd5e.Item{{Name: "Old Blade", Qty: 1, Weapon: &dnd5e.Weapon{Name: "Old Blade"}}},
	}}}
	var b strings.Builder
	if err := CommandBuilder(s, "mods", "mod-add", "attack.Old Blade", BuilderExtra{}).Render(t.Context(), &b); err != nil {
		t.Fatal(err)
	}
	if html := b.String(); !regexp.MustCompile(`<input[^>]*name="subject"[^>]*value="attack\.Old Blade"`).MatchString(html) {
		t.Errorf("stowed weapon target not selected in %s", html)
	}
}
