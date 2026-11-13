package dnd5e

import (
	"strings"
	"testing"
)

func testCantrip(class, rng, desc string) Cantrip {
	return Cantrip{Name: "Test", Classes: []string{class}, Range: rng, Desc: desc}
}

// TestCantripFits: an invocation's cantrip must be on the list and deal
// damage, by an attack roll or with a range when it says so.
func TestCantripFits(t *testing.T) {
	const attackDesc, saveDesc = "Make a melee spell attack. The target takes 1d8 Necrotic damage.", "Wisdom saving throw or take 1d8 Necrotic damage."
	for _, tc := range []struct {
		name     string
		c        Cantrip
		attack   bool
		minRange int
		want     bool
	}{
		{"damage", testCantrip("warlock", "60 feet", saveDesc), false, 0, true},
		{"no damage", testCantrip("warlock", "30 feet", "A spectral hand appears."), false, 0, false},
		{"other list", testCantrip("wizard", "60 feet", saveDesc), false, 0, false},
		{"attack roll", testCantrip("warlock", "Touch", attackDesc), true, 0, true},
		{"saving throw is no attack", testCantrip("warlock", "60 feet", saveDesc), true, 0, false},
		{"touch is under 10 feet", testCantrip("warlock", "Touch", attackDesc), false, 10, false},
		{"self cone is under 10 feet", testCantrip("warlock", "Self (15-foot cone)", saveDesc), false, 10, false},
		{"120 feet", testCantrip("warlock", "120 feet", attackDesc), false, 10, true},
		{"120 ft.", testCantrip("warlock", "120 ft.", attackDesc), false, 10, true},
		{"1 mile", testCantrip("warlock", "1 mile", attackDesc), false, 10, true},
		{"5 feet", testCantrip("warlock", "5 feet", attackDesc), false, 10, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CantripFits(tc.c, "warlock", tc.attack, tc.minRange); got != tc.want {
				t.Errorf("CantripFits = %v, want %v", got, tc.want)
			}
		})
	}
}

// TestMeetsPrerequisite checks each clause of a feat's or invocation's
// prerequisite must hold.
func TestMeetsPrerequisite(t *testing.T) {
	style := []Feature{{Name: "Fighting Style"}}
	blade := []Feature{{Name: "Eldritch Invocation: Pact of the Blade"}}
	blast := []Cantrip{testCantrip("warlock", "120 feet", "Make a ranged spell attack. The target takes 1d10 Force damage.")}
	toll := []Cantrip{testCantrip("warlock", "60 feet", "The target must succeed on a Wisdom saving throw or take 1d8 Necrotic damage.")}
	hand := []Cantrip{testCantrip("warlock", "30 feet", "A spectral hand appears.")}
	frost := []Cantrip{testCantrip("wizard", "60 feet", "Make a ranged spell attack. The target takes 1d8 Cold damage.")}
	const damage, attack = "Level 2+ Warlock, a Warlock Cantrip That Deals Damage", "Level 2+ Warlock, a Warlock Cantrip That Deals Damage via an Attack Roll"
	for _, tc := range []struct {
		name     string
		prereq   string
		level    int
		str, dex int
		features []Feature
		cantrips []Cantrip
		want     bool
	}{
		{"none", "", 1, 10, 10, nil, nil, true},
		{"level too low", "Level 4+", 3, 10, 10, nil, nil, false},
		{"level reached", "Level 4+", 4, 10, 10, nil, nil, true},
		{"Epic Boon at 18", "Level 19+", 18, 10, 10, nil, nil, false},
		{"Epic Boon at 19", "Level 19+", 19, 10, 10, nil, nil, true},
		{"either score", "Level 4+, Strength or Dexterity 13+", 4, 12, 14, nil, nil, true},
		{"neither score", "Level 4+, Strength or Dexterity 13+", 4, 12, 12, nil, nil, false},
		{"feature missing", "Fighting Style Feature", 4, 10, 10, nil, nil, false},
		{"feature held", "Fighting Style Feature", 1, 10, 10, style, nil, true},
		{"one of two features", "Spellcasting or Pact Magic Feature", 4, 10, 10, []Feature{{Name: "Pact Magic"}}, nil, true},
		{"level met, feature not", "Level 19+, Spellcasting Feature", 19, 10, 10, style, nil, false},
		{"unknown clause passes", "Proficiency with Heavy armor", 1, 10, 10, nil, nil, true},
		{"class level too low", "Level 5+ Warlock", 4, 10, 10, nil, nil, false},
		{"class level reached", "Level 5+ Warlock", 5, 10, 10, nil, nil, true},
		{"invocation missing", "Level 5+ Warlock, Pact of the Blade Invocation", 5, 10, 10, style, nil, false},
		{"invocation held", "Level 5+ Warlock, Pact of the Blade Invocation", 5, 10, 10, blade, nil, true},
		{"no cantrip", damage, 2, 10, 10, nil, nil, false},
		{"cantrip without damage", damage, 2, 10, 10, nil, hand, false},
		{"damaging cantrip of another list", damage, 2, 10, 10, nil, frost, false},
		{"damaging cantrip", damage, 2, 10, 10, nil, blast, true},
		{"saving throw cantrip", damage, 2, 10, 10, nil, toll, true},
		{"saving throw is no attack roll", attack, 2, 10, 10, nil, toll, false},
		{"attack roll cantrip", attack, 2, 10, 10, nil, blast, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			sh := Sheet{Scores: map[string]int{"str": tc.str, "dex": tc.dex}, Features: tc.features}
			// strings.ToLower slugs the one-word class names used here.
			if got := MeetsPrerequisite(tc.prereq, sh, tc.level, tc.cantrips, strings.ToLower); got != tc.want {
				t.Errorf("MeetsPrerequisite(%q) = %v, want %v", tc.prereq, got, tc.want)
			}
		})
	}
}
