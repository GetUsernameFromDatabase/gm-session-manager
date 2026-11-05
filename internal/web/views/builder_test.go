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
