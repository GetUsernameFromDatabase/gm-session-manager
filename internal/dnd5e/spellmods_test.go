package dnd5e

import (
	"fmt"
	"slices"
	"testing"
)

// Texts are the srd-2024 spell entries, verbatim.
const (
	shieldOfFaith = "A shimmering field surrounds a creature of your choice within range, granting it a +2 bonus to AC for the duration."
	haste         = "Choose a willing creature that you can see within range. Until the spell ends, the target's Speed is doubled, it gains a +2 bonus to Armor Class, it has Advantage on Dexterity saving throws, and it gains an additional action on each of its turns. That action can be used to take only the Attack (one attack only), Dash, Disengage, Hide, or Utilize action. When the spell ends, the target is Incapacitated and has a Speed of 0 until the end of its next turn, as a wave of lethargy washes over it."
	slow          = "You alter time around up to six creatures of your choice in a 40-foot Cube within range. Each target must succeed on a Wisdom saving throw or be affected by this spell for the duration. An affected target's Speed is halved, it takes a −2 penalty to AC and Dexterity saving throws, and it can't take Reactions. On its turns, it can take either an action or a Bonus Action, not both, and it can make only one attack if it takes the Attack action. If it casts a spell with a Somatic component, there is a 25 percent chance the spell fails as a result of the target making the spell's gestures too slowly. An affected target repeats the save at the end of each of its turns, ending the spell on itself on a success."
	longstrider   = "You touch a creature. The target's Speed increases by 10 feet until the spell ends."
	bless         = "You bless up to three creatures within range. Whenever a target makes an attack roll or a saving throw before the spell ends, the target adds 1d4 to the attack roll or save."
	bane          = "Up to three creatures of your choice that you can see within range must each make a Charisma saving throw. Whenever a target that fails this save makes an attack roll or a saving throw before the spell ends, the target must subtract 1d4 from the attack roll or save."
	barkskin      = "You touch a willing creature. Until the spell ends, the target's skin assumes a bark-like appearance, and the target has an Armor Class of 17 if its AC is lower than that."
	magicWeapon   = "You touch a nonmagical weapon. Until the spell ends, that weapon becomes a magic weapon with a +1 bonus to attack rolls and damage rolls. The spell ends early if you cast it again."
	magicWeaponHL = "The bonus increases to +2 with a level 3–5 spell slot. The bonus increases to +3 with a level 6+ spell slot."
	// The SRD 5.1 wording of the same higher-level rule.
	magicWeapon2014HL = "When you cast this spell using a spell slot of 4th level or higher, the bonus increases to +2. When you use a spell slot of 6th level or higher, the bonus increases to +3."
)

func modPairs(mods []CustomMod) []string {
	var out []string
	for _, m := range mods {
		out = append(out, fmt.Sprintf("%s=%d", m.Target, m.Value))
	}
	slices.Sort(out)
	return out
}

func checkMods(t *testing.T, mods []CustomMod, want []string) {
	t.Helper()
	slices.Sort(want)
	if got := modPairs(mods); !slices.Equal(got, want) {
		t.Errorf("mods = %v, want %v", got, want)
	}
}

func TestSpellMods(t *testing.T) {
	tests := []struct {
		name  string
		desc  string
		level int
		want  []string
		dice  string
	}{
		{"Shield of Faith", shieldOfFaith, 1, []string{"ac=2"}, ""},
		{"Haste", haste, 3, []string{"ac=2"}, ""},
		{"Slow", slow, 3, []string{"ac=-2", "save.dex=-2"}, ""},
		{"Longstrider", longstrider, 1, []string{"speed=10"}, ""},
		{"Bless", bless, 1, nil, "+1d4"},
		{"Bane", bane, 1, nil, "−1d4"},
		{"Barkskin", barkskin, 2, nil, ""},
		{"nothing", "", 1, nil, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mods, dice := SpellMods(tt.desc, "", tt.level, tt.level)
			checkMods(t, mods, tt.want)
			if dice != tt.dice {
				t.Errorf("dice = %q, want %q", dice, tt.dice)
			}
		})
	}
}

// Slot cases are their own loop: the table above has no slot column.
func TestSpellModsSlot(t *testing.T) {
	tests := []struct {
		name, higher string
		slot, bonus  int
	}{
		{"2024 slot 2", magicWeaponHL, 2, 1},
		{"2024 slot 4", magicWeaponHL, 4, 2},
		{"2024 slot 6", magicWeaponHL, 6, 3},
		{"2014 slot 3", magicWeapon2014HL, 3, 1},
		{"2014 slot 5", magicWeapon2014HL, 5, 2},
		{"2014 slot 7", magicWeapon2014HL, 7, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mods, _ := SpellMods(magicWeapon, tt.higher, 2, tt.slot)
			checkMods(t, mods, []string{fmt.Sprintf("attack=%d", tt.bonus), fmt.Sprintf("damage=%d", tt.bonus)})
		})
	}
}
