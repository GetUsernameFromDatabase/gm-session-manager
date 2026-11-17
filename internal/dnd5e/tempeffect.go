package dnd5e

import (
	"errors"
	"fmt"
	"slices"
)

// MaxEffectMods caps the modifier rows of one effect.
const MaxEffectMods = 10

// Incapacitating are the conditions that end the concentration of whoever
// gets one.
var Incapacitating = []string{Incapacitated, Paralyzed, Petrified, Stunned, Unconscious}

// EndsConcentration reports whether gaining the condition slug ends the
// holder's concentration.
func EndsConcentration(slug string) bool { return slices.Contains(Incapacitating, slug) }

// WithEffects is a copy of s with the modifiers of active effects added, each
// sourced from its effect, for showing numbers. The copy has its own slice,
// so it never reaches the stored sheet through s.
func (s Sheet) WithEffects(mods []CustomMod) Sheet {
	if len(mods) == 0 {
		return s
	}
	s.Modifiers = append(slices.Clone(s.Modifiers), mods...)
	return s
}

// EffectMods are the rows of an effect named name, as sheet modifiers: the
// note and source are the name, which the explain breakdown shows.
func EffectMods(name string, rows []CustomMod) []CustomMod {
	out := make([]CustomMod, len(rows))
	for i, m := range rows {
		m.Note, m.Source = name, name
		out[i] = m
	}
	return out
}

// ValidateEffectMods checks an effect's modifier rows by the rules of a
// modifier on a sheet: a known target and a non-zero value.
func ValidateEffectMods(name string, rows []CustomMod) error {
	if len(rows) > MaxEffectMods {
		return fmt.Errorf("an effect takes at most %d changes", MaxEffectMods)
	}
	for _, m := range EffectMods(name, rows) {
		if !IsModTarget(m.Target) {
			return fmt.Errorf("cannot change %q", m.Target)
		}
		if m.Value == 0 {
			return errors.New("a change must be a non-zero number such as +1 or -2")
		}
		if err := validate.Struct(m); err != nil {
			return fmt.Errorf("change to %s: %w", m.Target, err)
		}
	}
	return nil
}
