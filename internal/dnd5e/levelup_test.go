package dnd5e

import "testing"

func TestRevertSheet(t *testing.T) {
	snap := Sheet{
		Classes:     []ClassLevel{{Slug: "wizard", Name: "Wizard", Level: 2}},
		Scores:      map[string]int{"str": 8},
		Feats:       []string{"alert"},
		Coins:       Coins{GP: 5},
		Inventory:   []Item{{Name: "Rope", Qty: 1}},
		HitDiceUsed: 0,
	}
	snap.Slots[0] = Slot{Max: 2, Used: 0}
	snap.Pact = Slot{Max: 1}
	snap.Features = []Feature{
		{Name: "Rage", Source: "Barbarian", Uses: Uses{Fixed: 2, Used: 2}},
		{Name: "Second Wind", Source: "Fighter", Uses: Uses{Fixed: 3, Used: 1}},
	}
	cur := Sheet{
		Classes:     []ClassLevel{{Slug: "wizard", Name: "Wizard", Level: 3}},
		Scores:      map[string]int{"str": 10},
		Feats:       []string{"alert", "tough"},
		Coins:       Coins{GP: 90},
		Inventory:   []Item{{Name: "Torch", Qty: 4}},
		Weapons:     []Weapon{{Name: "Dagger"}},
		Armor:       &WornArmor{},
		Shield:      true,
		Backstory:   "now",
		Inspiration: true,
		DeathSaves:  DeathSaves{Failures: 1},
		HitDiceUsed: 3,
	}
	cur.Slots[0] = Slot{Max: 4, Used: 3}
	cur.Slots[2] = Slot{Max: 2, Used: 1}
	cur.Pact = Slot{Max: 2, Used: 2}
	cur.Features = []Feature{
		{Name: "Rage", Source: "Barbarian", Uses: Uses{Fixed: 3, Used: 0}}, // a long rest since
		{Name: "Second Wind", Source: "Fighter", Uses: Uses{Fixed: 4, Used: 4}},
	}

	got := RevertSheet(snap, cur)
	cases := []struct {
		name  string
		check func(t *testing.T)
	}{
		{"gear and play state kept", func(t *testing.T) {
			if got.Coins.GP != 90 || len(got.Inventory) != 1 || got.Inventory[0].Name != "Torch" ||
				len(got.Weapons) != 1 || got.Armor == nil || !got.Shield || got.Backstory != "now" ||
				!got.Inspiration || got.DeathSaves.Failures != 1 {
				t.Errorf("not kept from current: %+v", got)
			}
		}},
		{"classes, scores and feats restored", func(t *testing.T) {
			if got.Level() != 2 || got.Scores["str"] != 8 || len(got.Feats) != 1 {
				t.Errorf("not restored from snapshot: %+v", got)
			}
		}},
		{"slot used clamped to restored max", func(t *testing.T) {
			if got.Slots[0] != (Slot{Max: 2, Used: 2}) {
				t.Errorf("level-1 slot %+v, want max 2 used 2", got.Slots[0])
			}
		}},
		{"slot only the current sheet had is empty", func(t *testing.T) {
			if got.Slots[2] != (Slot{}) {
				t.Errorf("level-3 slot %+v, want zero", got.Slots[2])
			}
		}},
		{"pact used clamped", func(t *testing.T) {
			if got.Pact != (Slot{Max: 1, Used: 1}) {
				t.Errorf("pact %+v, want max 1 used 1", got.Pact)
			}
		}},
		{"feature uses kept, clamped to restored max", func(t *testing.T) {
			if got.Features[0].Uses.Used != 0 || got.Features[1].Uses.Used != 3 {
				t.Errorf("feature uses %+v %+v, want rage 0/2 and second wind 3/3", got.Features[0].Uses, got.Features[1].Uses)
			}
			if snap.Features[0].Uses.Used != 2 {
				t.Error("snapshot features changed")
			}
		}},
		{"hit dice used clamped to snapshot level", func(t *testing.T) {
			if got.HitDiceUsed != 2 {
				t.Errorf("hit dice used %d, want 2", got.HitDiceUsed)
			}
		}},
		{"result does not alias current", func(t *testing.T) {
			got.Inventory[0].Name = "changed"
			if cur.Inventory[0].Name != "Torch" {
				t.Error("inventory shares memory with current")
			}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, tc.check)
	}
}
