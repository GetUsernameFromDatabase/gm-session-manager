package dnd5e

import (
	"slices"
	"strings"
)

// SkillKeys matches skill names (spaces from the source's PDF text and all)
// to skill keys; no names gives every skill.
func SkillKeys(names []string) []string {
	var keys []string
	for _, sk := range Skills {
		if names == nil || slices.ContainsFunc(names, func(n string) bool { return Squash(n) == Squash(sk.Name) }) {
			keys = append(keys, sk.Key)
		}
	}
	return keys
}

// AddResistance adds a damage type the sheet resists, once.
func AddResistance(s *Sheet, damage string) {
	if d := strings.ToLower(damage); slices.Contains(DamageTypes, d) && !slices.Contains(s.Resistances, d) {
		s.Resistances = append(s.Resistances, d)
	}
}

// AddFeat puts a feat on the sheet with what the sheet can automate.
func AddFeat(s *Sheet, f Feature) {
	mods := ParseEffects(f.Name, f.Description, nil, "").Apply(&f)
	s.PutFeature(f, mods)
	s.Feats = append(s.Feats, f.Name)
}

// FreeCast is a lineage spell's one cast without a slot per Long Rest, as a
// feature so the Actions card counts it. Like other species features it has
// no level (Feature.Level is a class level).
func FreeCast(sp Spell, species string) Feature {
	return Feature{
		Name: sp.Name + " (free cast)", Source: species,
		Description: "Cast *" + sp.Name + "* once without a spell slot; you regain this when you finish a Long Rest.",
		Uses:        Uses{Fixed: 1, Reset: "long"},
	}
}

// SwapFeatSpell puts repl in old's place, prepared and cast as old was, and
// renames old's free cast. A repl the sheet already has leaves old as it is:
// the picker offers none, but this keeps a duplicate off the sheet.
func SwapFeatSpell(s *Sheet, old, repl Spell) {
	i := slices.IndexFunc(s.Spells, func(sp Spell) bool { return sp.Slug == old.Slug })
	if i < 0 || slices.ContainsFunc(s.Spells, func(sp Spell) bool { return sp.Slug == repl.Slug }) {
		return
	}
	repl.Prepared, repl.Ability, repl.Feat = s.Spells[i].Prepared, s.Spells[i].Ability, s.Spells[i].Feat
	s.Spells[i] = repl
	if j := slices.IndexFunc(s.Features, func(f Feature) bool { return f.Name == FreeCast(old, "").Name }); j >= 0 {
		fc := FreeCast(repl, s.Features[j].Source)
		s.Features[j].Name, s.Features[j].Description = fc.Name, fc.Description
	}
}

// AbilityKeys matches a list of ability names ("Strength, Dexterity, or
// Constitution") to ability keys.
func AbilityKeys(list string) []string {
	var keys []string
	for _, n := range SplitList(list) {
		for _, a := range Abilities {
			if strings.EqualFold(AbilityNames[a], n) {
				keys = append(keys, a)
			}
		}
	}
	return keys
}
