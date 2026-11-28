package dnd5e

import (
	"testing"

	"gm-session-manager/internal/perm"
)

func TestApplyHPDeath(t *testing.T) {
	tests := []struct {
		name         string
		hp, max, tmp int
		ds           DeathSaves
		amt          Amount
		wantHP       int
		wantTemp     int
		wantDS       DeathSaves
	}{
		{
			"temp soaks it all",
			0,
			20,
			5,
			DeathSaves{Failures: 1},
			Damage(5),
			0,
			0,
			DeathSaves{Failures: 1},
		},
		{
			"temp soaks part, rest is one failure",
			0,
			20,
			3,
			DeathSaves{},
			Damage(5),
			-2,
			0,
			DeathSaves{Failures: 1},
		},
		{"one failure at 0", 0, 20, 0, DeathSaves{}, Damage(2), -2, 0, DeathSaves{Failures: 1}},
		{
			"crit at 0 is two failures",
			0,
			20,
			0,
			DeathSaves{},
			CritDamage(2),
			-2,
			0,
			DeathSaves{Failures: 2},
		},
		{
			"successes stay",
			0,
			20,
			0,
			DeathSaves{Successes: 2, Failures: 1},
			Damage(2),
			-2,
			0,
			DeathSaves{Successes: 2, Failures: 2},
		},
		{
			"failures cap at 3",
			0,
			20,
			0,
			DeathSaves{Failures: 2},
			CritDamage(2),
			-2,
			0,
			DeathSaves{Failures: 3},
		},
		{
			"stable resets then fails",
			0,
			20,
			0,
			DeathSaves{Successes: 3, Failures: 1},
			Damage(2),
			-2,
			0,
			DeathSaves{Failures: 1},
		},
		{
			"stable crit resets then two",
			0,
			20,
			0,
			DeathSaves{Successes: 3},
			CritDamage(2),
			-2,
			0,
			DeathSaves{Failures: 2},
		},
		{
			"massive at 0",
			0,
			20,
			0,
			DeathSaves{Successes: 1},
			Damage(20),
			-20,
			0,
			DeathSaves{Successes: 1, Failures: 3},
		},
		{
			"massive beats stable reset",
			0,
			20,
			0,
			DeathSaves{Successes: 3},
			Damage(20),
			-20,
			0,
			DeathSaves{Successes: 3, Failures: 3},
		},
		{"drop to 0 is no failure", 5, 20, 0, DeathSaves{}, Damage(5), 0, 0, DeathSaves{}},
		{
			"drop with small overkill is no failure",
			5,
			20,
			0,
			DeathSaves{},
			Damage(10),
			-5,
			0,
			DeathSaves{},
		},
		{"drop crit is no failure", 5, 20, 0, DeathSaves{}, CritDamage(5), 0, 0, DeathSaves{}},
		{
			"drop with overflow at max is dead",
			5,
			20,
			0,
			DeathSaves{},
			Damage(25),
			-20,
			0,
			DeathSaves{Failures: 3},
		},
		{
			"overflow just under max is not",
			5,
			20,
			0,
			DeathSaves{},
			Damage(24),
			-19,
			0,
			DeathSaves{},
		},
		{
			"temp counts before overflow",
			5,
			20,
			10,
			DeathSaves{},
			Damage(35),
			-20,
			0,
			DeathSaves{Failures: 3},
		},
		{
			"stored overkill is not massive",
			-15,
			20,
			0,
			DeathSaves{Failures: 1},
			Damage(2),
			-17,
			0,
			DeathSaves{Failures: 2},
		},
		{
			"no max HP, no massive",
			0,
			0,
			0,
			DeathSaves{},
			Damage(50),
			-50,
			0,
			DeathSaves{Failures: 1},
		},
		{"no max HP, drop is no failure", 5, 0, 0, DeathSaves{}, Damage(50), -45, 0, DeathSaves{}},
		{
			"dead keeps taking damage",
			-5,
			20,
			0,
			DeathSaves{Failures: 3},
			Damage(4),
			-9,
			0,
			DeathSaves{Failures: 3},
		},
		{
			"dead with stable count stays dead",
			-5,
			20,
			0,
			DeathSaves{Successes: 3, Failures: 3},
			Damage(4),
			-9,
			0,
			DeathSaves{Successes: 3, Failures: 3},
		},
		{
			"damage above 0 changes no saves",
			10,
			20,
			0,
			DeathSaves{Failures: 1},
			Damage(3),
			7,
			0,
			DeathSaves{Failures: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := fighter()
			c.HP, c.MaxHP, c.TempHP, c.Sheet.DeathSaves = tt.hp, tt.max, tt.tmp, tt.ds
			if err := c.ApplyHP(tt.amt); err != nil {
				t.Fatal(err)
			}
			if c.HP != tt.wantHP || c.TempHP != tt.wantTemp || c.Sheet.DeathSaves != tt.wantDS {
				t.Errorf(
					"hp %d temp %d saves %+v; want %d %d %+v",
					c.HP,
					c.TempHP,
					c.Sheet.DeathSaves,
					tt.wantHP,
					tt.wantTemp,
					tt.wantDS,
				)
			}
		})
	}
}

// A monster or NPC combatant is a Character with an empty sheet; it takes
// damage at 0 HP and heals from it as before.
func TestApplyHPMonsterUnaffected(t *testing.T) {
	m := Character{HP: 0, MaxHP: 20}
	if err := m.ApplyHP(Damage(5)); err != nil || m.HP != -5 {
		t.Fatalf("damage: hp %d err %v", m.HP, err)
	}
	if err := m.ApplyHP(Heal(4)); err != nil || m.HP != 4 {
		t.Fatalf("heal: hp %d err %v", m.HP, err)
	}
}

func TestDeadRefusals(t *testing.T) {
	for _, line := range []string{"hp +3", "hp +0", "hp =5", "temp 5", "temp +2", "hitdie use", "hitdie restore", "rest short", "rest long"} {
		t.Run(line, func(t *testing.T) {
			c := fighter()
			c.HP = -3
			c.Sheet.DeathSaves.Failures = MaxDeathSaves
			cmd, err := Parse(line)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := cmd.Apply(&c, Env{}); err == nil || err.Error() != "Dead. Use revive." {
				t.Errorf("err = %v", err)
			}
			if c.HP != -3 {
				t.Errorf("hp changed to %d", c.HP)
			}
		})
	}
	// Everything else still works while dead.
	for _, line := range []string{"hp -3", "maxhp +2", "item add Rope", "deathsave success", "deathsave reset", "revive"} {
		t.Run("allowed "+line, func(t *testing.T) {
			c := fighter()
			c.HP = -3
			c.Sheet.DeathSaves.Failures = MaxDeathSaves
			cmd, err := Parse(line)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := cmd.Apply(&c, Env{}); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestDeathSaveFailKills(t *testing.T) {
	c := fighter()
	c.HP = 0
	c.Sheet.DeathSaves.Failures = 2
	cmd, _ := Parse("deathsave fail")
	if _, err := cmd.Apply(&c, Env{}); err != nil || !c.Sheet.Dead() {
		t.Fatalf("dead %v err %v", c.Sheet.Dead(), err)
	}
	cmd, _ = Parse("deathsave reset")
	if _, err := cmd.Apply(&c, Env{}); err != nil || c.Sheet.Dead() || c.HP != 0 {
		t.Errorf("reset: dead %v hp %d err %v", c.Sheet.Dead(), c.HP, err)
	}
}

func TestRevive(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		dead    bool
		wantHP  int
		wantErr string
	}{
		{"default 1", "revive", true, 1, ""},
		{"n", "revive 6", true, 6, ""},
		{"capped at max", "revive 99", true, 28, ""},
		{"not dead", "revive", false, 0, "Not dead."},
		{"zero", "revive 0", true, 0, "revive: n must be a positive whole number"},
		{"negative", "revive -2", true, 0, "revive: n must be a positive whole number"},
		{"text", "revive x", true, 0, "revive: n must be a positive whole number"},
		{"too many", "revive 1 2", true, 0, "revive: usage: revive | revive 5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := fighter()
			c.HP = -4
			if tt.dead {
				c.Sheet.DeathSaves = DeathSaves{Successes: 1, Failures: MaxDeathSaves}
			}
			cmd, err := Parse(tt.line)
			if err == nil {
				if cmd.Scope != perm.ScopeSheet {
					t.Errorf("scope %v", cmd.Scope)
				}
				_, err = cmd.Apply(&c, Env{})
			}
			if tt.wantErr != "" {
				if err == nil || err.Error() != tt.wantErr {
					t.Fatalf("err = %v, want %q", err, tt.wantErr)
				}
				return
			}
			if err != nil || c.HP != tt.wantHP || c.Sheet.DeathSaves != (DeathSaves{}) {
				t.Errorf("hp %d saves %+v err %v", c.HP, c.Sheet.DeathSaves, err)
			}
		})
	}
}

func TestHPCritParsing(t *testing.T) {
	tests := []struct {
		line      string
		wantFails int
		wantErr   bool
	}{
		{"hp -3 crit", 2, false},
		{"hp -3 fire crit", 2, false},
		{"hp -3 CRIT", 2, false},
		{"hp -3 fire", 1, false},
		{"hp -3", 1, false},
		{"hp +3 crit", 0, true},
		{"hp =3 crit", 0, true},
		{"hp -3 crit fire", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			c := fighter()
			c.HP, c.TempHP = 0, 0
			cmd, err := Parse(tt.line)
			if err == nil {
				_, err = cmd.Apply(&c, Env{})
			}
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && c.Sheet.DeathSaves.Failures != tt.wantFails {
				t.Errorf("failures %d, want %d", c.Sheet.DeathSaves.Failures, tt.wantFails)
			}
		})
	}
}
