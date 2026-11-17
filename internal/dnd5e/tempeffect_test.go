package dnd5e

import (
	"strings"
	"testing"
)

func TestEndsConcentration(t *testing.T) {
	for slug, want := range map[string]bool{
		Incapacitated: true, Paralyzed: true, Petrified: true, Stunned: true, Unconscious: true,
		Prone: false, Poisoned: false, Exhaustion: false, TrackConcentration: false, TrackCustom: false,
	} {
		if got := EndsConcentration(slug); got != want {
			t.Errorf("EndsConcentration(%q) = %v, want %v", slug, got, want)
		}
	}
}

// TestWithEffects checks the copy shows the effects, names them in the
// breakdown, and leaves the stored sheet's modifiers alone.
func TestWithEffects(t *testing.T) {
	stored := Sheet{Scores: map[string]int{Dex: 10}, Modifiers: make([]CustomMod, 1, 8)}
	stored.Modifiers[0] = CustomMod{Target: ModAC, Value: 1, Note: "Ring"}
	base := stored.ArmorClass()

	shown := stored.WithEffects(EffectMods("Shield of Faith", []CustomMod{{Target: ModAC, Value: 2}}))
	if got := shown.ArmorClass(); got != base+2 {
		t.Errorf("AC with the effect %d, want %d", got, base+2)
	}
	if len(stored.Modifiers) != 1 || stored.ArmorClass() != base {
		t.Errorf("stored sheet changed: %+v", stored.Modifiers)
	}
	if stored.Modifiers[:2][1] != (CustomMod{}) {
		t.Error("the copy was appended into the stored sheet's backing array")
	}
	named := false
	for _, p := range shown.Explain(ModAC) {
		named = named || strings.Contains(p.Label, "Shield of Faith")
	}
	if !named {
		t.Errorf("breakdown does not name the effect: %+v", shown.Explain(ModAC))
	}
	if same := stored.WithEffects(nil); same.ArmorClass() != base {
		t.Error("no effects must leave the sheet as it is")
	}
}

func TestValidateEffectMods(t *testing.T) {
	many := make([]CustomMod, MaxEffectMods+1)
	for i := range many {
		many[i] = CustomMod{Target: ModAC, Value: 1}
	}
	tests := []struct {
		name string
		mods []CustomMod
		ok   bool
	}{
		{"none", nil, true},
		{"ac", []CustomMod{{Target: ModAC, Value: 2}}, true},
		{"save", []CustomMod{{Target: "save.wis", Value: -1}}, true},
		{"conditional", []CustomMod{{Target: ModAC, Value: 1, When: WhenUnarmored}}, true},
		{"unknown target", []CustomMod{{Target: "flying", Value: 1}}, false},
		{"zero", []CustomMod{{Target: ModAC}}, false},
		{"too big", []CustomMod{{Target: ModAC, Value: 51}}, false},
		{"too small", []CustomMod{{Target: ModAC, Value: -51}}, false},
		{"bad when", []CustomMod{{Target: ModAC, Value: 1, When: "never"}}, false},
		{"too many", many, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateEffectMods("Bless", tt.mods); (err == nil) != tt.ok {
				t.Errorf("err = %v, want ok %v", err, tt.ok)
			}
		})
	}
	if err := ValidateEffectMods(strings.Repeat("x", 201), []CustomMod{{Target: ModAC, Value: 1}}); err == nil {
		t.Error("a name over 200 characters must be refused")
	}
}

// TestRestEndsEffects checks which rest emits which end op, with the
// long rest's exhaustion op kept.
func TestRestEndsEffects(t *testing.T) {
	tests := []struct {
		line string
		want []CondOp
	}{
		{"rest short", []CondOp{{Op: CondEndRest, Rest: ResetShort}}},
		{"rest long", []CondOp{{Op: CondExhaustion, Slug: Exhaustion, Delta: -1}, {Op: CondEndRest, Rest: ResetLong}}},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			cmd, err := Parse(tt.line)
			if err != nil {
				t.Fatal(err)
			}
			c := fighter()
			fx, err := cmd.Apply(&c, Env{})
			if err != nil {
				t.Fatal(err)
			}
			if len(fx.Conditions) != len(tt.want) {
				t.Fatalf("ops %+v, want %+v", fx.Conditions, tt.want)
			}
			for i, op := range fx.Conditions {
				if op != tt.want[i] {
					t.Errorf("op %d = %+v, want %+v", i, op, tt.want[i])
				}
			}
		})
	}
}
