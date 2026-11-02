package dnd5e

import (
	"fmt"
	"slices"
	"strings"
)

// Part is one line of a stat's breakdown. Off parts are modifiers whose
// condition does not hold right now; they are shown but not counted.
type Part struct {
	Label string
	Value int
	Off   bool
}

// Sum adds up the parts that count.
func Sum(parts []Part) int {
	n := 0
	for _, p := range parts {
		if !p.Off {
			n += p.Value
		}
	}
	return n
}

// Explain breaks a stat down into what makes it up, so the sheet can show
// why a number is what it is. The stat methods sum these same parts.
// Targets: ac, speed, init, passive, save.<ability>, skill.<key>,
// attack.<weapon name>, damage.<weapon name>, offhand.<weapon name> (its
// damage as a two-weapon fighting bonus attack), spell.attack and spell.dc.
func (s Sheet) Explain(target string) []Part {
	kind, rest, _ := strings.Cut(target, ".")
	switch {
	case target == ModAC:
		return append(s.acBase(), s.modParts(ModAC)...)
	case target == ModSpeed:
		return append(s.speedBase(), s.modParts(ModSpeed)...)
	case target == ModInit:
		return append([]Part{s.abilityPart("dex")}, s.modParts(ModInit)...)
	case target == ModPassive:
		return append(append([]Part{{Label: "Base", Value: 10}}, s.skillParts(skillByKey("perception"))...), s.modParts(ModPassive)...)
	case kind == "save":
		parts := []Part{s.abilityPart(rest)}
		if slices.Contains(s.SaveProfs, rest) {
			parts = append(parts, Part{Label: "Proficiency", Value: s.ProficiencyBonus()})
		}
		return append(parts, s.modParts(target)...)
	case kind == "skill":
		return s.skillParts(skillByKey(rest))
	case kind == ModAttack, kind == ModDamage, kind == "offhand":
		if i := slices.IndexFunc(s.Weapons, func(w Weapon) bool { return w.Name == rest }); i >= 0 {
			switch kind {
			case ModDamage:
				return s.damageParts(s.Weapons[i])
			case "offhand":
				return s.offHandParts(s.Weapons[i])
			}
			return s.attackParts(s.Weapons[i])
		}
	case target == "spell.attack", target == "spell.dc":
		ab := s.SpellAbility()
		if ab == "" {
			return nil
		}
		parts := []Part{s.abilityPart(ab), {Label: "Proficiency", Value: s.ProficiencyBonus()}}
		if rest == "dc" {
			parts = append([]Part{{Label: "Base", Value: 8}}, parts...)
		}
		return parts
	}
	return nil
}

func skillByKey(key string) Skill {
	i := slices.IndexFunc(Skills, func(sk Skill) bool { return sk.Key == key })
	if i < 0 {
		return Skill{}
	}
	return Skills[i]
}

func (s Sheet) abilityPart(ability string) Part {
	return Part{Label: AbilityNames[ability] + " modifier", Value: s.Mod(ability)}
}

// modParts are the custom modifiers on target, those whose condition does
// not hold marked off.
func (s Sheet) modParts(target string) []Part {
	var parts []Part
	for _, m := range s.Modifiers {
		if m.Target != target {
			continue
		}
		label := m.Note
		if w := WhenLabels[m.When]; w != "" {
			label += " (" + w + ")"
		}
		parts = append(parts, Part{Label: label, Value: m.Value, Off: !s.Applies(m)})
	}
	return parts
}

// acBase is the AC before modifiers: worn armor, else the best unarmored
// defense, else 10 + Dexterity; plus a shield and the bonus set by hand.
func (s Sheet) acBase() []Part {
	var parts []Part
	dex := s.Mod("dex")
	if a := s.Armor; a != nil {
		parts = append(parts, Part{Label: a.Name, Value: a.Base})
		if a.AddDex {
			label := "Dexterity modifier"
			if a.DexCap > 0 && dex > a.DexCap {
				dex, label = a.DexCap, fmt.Sprintf("Dexterity modifier (at most %d in %s)", a.DexCap, a.Type)
			}
			parts = append(parts, Part{Label: label, Value: dex})
		}
	} else if f := s.unarmoredDefense(); f != nil {
		parts = append(parts, Part{Label: f.Name, Value: 10}, s.abilityPart("dex"), s.abilityPart(f.Unarmored))
	} else {
		parts = append(parts, Part{Label: "Unarmored", Value: 10}, s.abilityPart("dex"))
	}
	if s.Shield {
		parts = append(parts, Part{Label: "Shield", Value: 2})
	}
	if s.ACBonus != 0 {
		parts = append(parts, Part{Label: "AC bonus (set by hand)", Value: s.ACBonus})
	}
	return parts
}

// ArmorEffect says what armor makes of AC with a Dexterity modifier of
// dex: "AC 16, no Dex", "AC 11 + Dex (+3)" or "AC 14 + Dex, at most +2 (+2)".
func ArmorEffect(a WornArmor, dex int) string {
	switch {
	case !a.AddDex:
		return fmt.Sprintf("AC %d, no Dex", a.Base)
	case a.DexCap > 0:
		return fmt.Sprintf("AC %d + Dex, at most +%d (%s)", a.Base, a.DexCap, Signed(min(dex, a.DexCap)))
	}
	return fmt.Sprintf("AC %d + Dex (%s)", a.Base, Signed(dex))
}

// unarmoredDefense is the feature giving the best unarmored AC that
// applies now, or nil.
func (s Sheet) unarmoredDefense() *Feature {
	var best *Feature
	for i, f := range s.Features {
		if f.Unarmored == "" || f.NoShield && s.Shield {
			continue
		}
		if best == nil || s.Mod(f.Unarmored) > s.Mod(best.Unarmored) {
			best = &s.Features[i]
		}
	}
	return best
}

// ArmorTooHeavy reports whether the worn armor's Strength requirement is
// not met, which costs 10 feet of speed.
func (s Sheet) ArmorTooHeavy() bool {
	return s.Armor != nil && s.Armor.StrReq > 0 && s.Score("str") < s.Armor.StrReq
}

// StealthDisadvantage reports whether the worn armor hampers Stealth.
func (s Sheet) StealthDisadvantage() bool { return s.Armor != nil && s.Armor.Stealth }

// ArmorUntrained names the worn armor types (or shield) the character has
// no armor training for, judged from its other proficiencies.
func (s Sheet) ArmorUntrained() []string {
	trained := strings.ToLower(strings.Join(s.OtherProfs, " "))
	var out []string
	if s.Armor != nil && !strings.Contains(trained, s.Armor.Type) {
		out = append(out, s.Armor.Type+" armor")
	}
	if s.Shield && !strings.Contains(trained, "shield") {
		out = append(out, "shields")
	}
	return out
}

func (s Sheet) speedBase() []Part {
	parts := []Part{{Label: "Base speed", Value: s.Speed}}
	if s.ArmorTooHeavy() {
		parts = append(parts, Part{Label: fmt.Sprintf("%s needs Strength %d", s.Armor.Name, s.Armor.StrReq), Value: -10})
	}
	return parts
}

func (s Sheet) skillParts(sk Skill) []Part {
	if sk.Key == "" {
		return nil
	}
	parts := []Part{s.abilityPart(sk.Ability)}
	switch {
	case slices.Contains(s.Expertise, sk.Key):
		parts = append(parts, Part{Label: "Expertise (twice proficiency)", Value: 2 * s.ProficiencyBonus()})
	case slices.Contains(s.SkillProfs, sk.Key):
		parts = append(parts, Part{Label: "Proficiency", Value: s.ProficiencyBonus()})
	}
	return append(parts, s.modParts("skill."+sk.Key)...)
}

// weaponAbilityPart is the ability modifier a weapon attacks and damages with.
func (s Sheet) weaponAbilityPart(w Weapon) Part {
	label := AbilityNames[w.Ability] + " modifier"
	if w.Ability == "finesse" {
		label = "Finesse: better of Strength and Dexterity"
	}
	return Part{Label: label, Value: s.AttackMod(w)}
}

func (s Sheet) attackParts(w Weapon) []Part {
	parts := []Part{s.weaponAbilityPart(w)}
	if w.Proficient {
		parts = append(parts, Part{Label: "Proficiency", Value: s.ProficiencyBonus()})
	}
	return append(append(parts, s.modParts(ModAttack)...), s.modParts(ModAttack+"."+w.Name)...)
}

func (s Sheet) damageParts(w Weapon) []Part {
	parts := []Part{s.weaponAbilityPart(w)}
	return append(append(parts, s.modParts(ModDamage)...), s.modParts(ModDamage+"."+w.Name)...)
}

// offHandParts are a light weapon's damage parts as the bonus attack of the
// Light property: no ability modifier unless it is negative, or the
// Two-Weapon Fighting style adds it.
func (s Sheet) offHandParts(w Weapon) []Part {
	parts := s.damageParts(w)
	if parts[0].Value > 0 && !s.hasFeature("Two-Weapon Fighting") {
		parts[0].Off = true
		parts[0].Label += " (not on the off-hand attack)"
	}
	return parts
}

func (s Sheet) hasFeature(name string) bool {
	return slices.ContainsFunc(s.Features, func(f Feature) bool { return strings.EqualFold(f.Name, name) })
}

// LightWeapons are the ready weapons with the Light property.
func (s Sheet) LightWeapons() []Weapon {
	return slices.DeleteFunc(slices.Clone(s.Weapons), func(w Weapon) bool {
		return !strings.Contains(strings.ToLower(w.Properties), "light")
	})
}

// CanDualWield reports whether the character holds two light weapons, two
// of one kind counting, for the Light property's bonus attack.
func (s Sheet) CanDualWield() bool {
	n := 0
	for _, w := range s.LightWeapons() {
		if i := findItem(s.Inventory, w.Name); i >= 0 {
			n += max(s.Inventory[i].Qty, 1)
		} else {
			n++
		}
	}
	return n >= 2
}
