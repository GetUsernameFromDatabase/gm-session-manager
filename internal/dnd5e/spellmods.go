package dnd5e

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	// "+2 bonus to AC", "−2 penalty to AC and Dexterity saving throws, ..."
	spellBonusRe = regexp.MustCompile(`([+\-−–])(\d+) (?:bonus|penalty) to ([^.,;]+)`)
	spellSpeedRe = regexp.MustCompile(
		`(?i)\bspeed (increases|is increased|is reduced) by (\d+) feet`,
	)
	spellDiceRe = regexp.MustCompile(`(?i)\b(adds|subtracts?) (\d+d\d+)|\+(\d+d\d+) bonus`)
	// "+2 with a level 3–5 spell slot",
	// "spell slot of 4th level or higher, the bonus increases to +2"
	spellSlotRe    = regexp.MustCompile(`\+(\d+) with a level (\d+)`)
	spellSlotOldRe = regexp.MustCompile(
		`spell slot of (\d+)(?:st|nd|rd|th) level or higher[^.]*?\+(\d+)`,
	)
	saveTargetRe = regexp.MustCompile(`(?i)(?:(\w+) )?saving throws`)
	acTargetRe   = regexp.MustCompile(`(?i)\b(?:AC|armor class)\b`)
)

// SpellMods are the stat changes and the dice bonus or penalty ("+1d4",
// "−1d4") a spell's text states, for a cast with a slot of that level. The
// higher-level text may raise a flat bonus. Anything unmatched gives nothing.
// shortcut: only the flat "N bonus/penalty to X" and "Speed increases/is
// reduced by N feet" wordings, add more as other spells need them.
func SpellMods(desc, higher string, spellLevel, slot int) (mods []CustomMod, dice string) {
	slot = max(slot, spellLevel) // a slot below the spell's level is no slot
	add := func(target string, v int) {
		mods = append(mods, CustomMod{Target: target, Value: v, Note: "Spell text"})
	}
	for _, m := range spellBonusRe.FindAllStringSubmatch(desc, -1) {
		n, _ := strconv.Atoi(m[2])
		if m[1] != "+" {
			n = -n
		}
		x := strings.ToLower(m[3])
		if acTargetRe.MatchString(x) {
			add(ModAC, n)
		}
		if strings.Contains(x, "attack rolls") {
			add(ModAttack, n)
		}
		if strings.Contains(x, "damage rolls") {
			add(ModDamage, n)
		}
		if strings.Contains(x, "initiative") {
			add(ModInit, n)
		}
		for _, s := range saveTargetRe.FindAllStringSubmatch(x, -1) {
			if i := indexAbilityName(s[1]); i >= 0 {
				add(ModSave+"."+Abilities[i], n)
				continue
			}
			for _, ab := range Abilities {
				add(ModSave+"."+ab, n)
			}
		}
		for _, sk := range Skills {
			if strings.Contains(x, strings.ToLower(sk.Name)+" checks") {
				add(ModSkill+"."+sk.Key, n)
			}
		}
	}
	if m := spellSpeedRe.FindStringSubmatch(desc); m != nil {
		n, _ := strconv.Atoi(m[2])
		if strings.EqualFold(m[1], "is reduced") {
			n = -n
		}
		add(ModSpeed, n)
	}
	if m := spellDiceRe.FindStringSubmatch(desc); m != nil {
		switch {
		case m[3] != "":
			dice = "+" + m[3]
		case strings.HasPrefix(strings.ToLower(m[1]), "subtract"):
			dice = "−" + m[2]
		default:
			dice = "+" + m[2]
		}
	}
	// The highest threshold at or below the slot wins and replaces every flat value.
	best, bestAt := 0, 0
	pick := func(at, bonus string) {
		a, _ := strconv.Atoi(at)
		if a <= slot && a > bestAt {
			best, _ = strconv.Atoi(bonus)
			bestAt = a
		}
	}
	for _, m := range spellSlotRe.FindAllStringSubmatch(higher, -1) {
		pick(m[2], m[1])
	}
	for _, m := range spellSlotOldRe.FindAllStringSubmatch(higher, -1) {
		pick(m[1], m[2])
	}
	if bestAt > 0 {
		for i := range mods {
			if mods[i].Value < 0 {
				mods[i].Value = -best
			} else {
				mods[i].Value = best
			}
		}
	}
	return mods, dice
}

// indexAbilityName is the index in Abilities of the ability named name (any
// case), or -1.
func indexAbilityName(name string) int {
	for i, ab := range Abilities {
		if strings.EqualFold(AbilityNames[ab], name) {
			return i
		}
	}
	return -1
}
