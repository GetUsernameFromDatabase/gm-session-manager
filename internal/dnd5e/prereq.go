package dnd5e

import (
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	rePrereqLevel      = regexp.MustCompile(`^Level (\d+)\+(?: \w+)?$`)
	rePrereqScore      = regexp.MustCompile(`^(.+) (\d+)\+$`)
	rePrereqFeature    = regexp.MustCompile(`^(.+) Feature$`)
	rePrereqInvocation = regexp.MustCompile(`^(.+) Invocation$`)
	rePrereqCantrip    = regexp.MustCompile(`(?i)^an? (.+?) cantrip that deals damage( via an attack roll)?$`)
	reSpellDamage      = regexp.MustCompile(`\d+d\d+ \w+ damage`)
	reSpellRange       = regexp.MustCompile(`(?i)^(\d+)\s*(?:feet|foot|ft\b\.?|(miles?)\b)`)
)

// Cantrip is a class cantrip a sheet knows, or may learn this level-up:
// the parts of its compendium entry prerequisites read.
type Cantrip struct {
	Name, Range, Desc string
	// Classes are the slugs of the spell lists it counts for.
	Classes []string
}

// CantripFits reports whether c is a cantrip of classSlug's list that deals
// damage, by a spell attack if attack, with a range of minRange feet or more
// ("120 feet", "120 ft.", "1 mile"; Touch and Self count as 0).
func CantripFits(c Cantrip, classSlug string, attack bool, minRange int) bool {
	feet := 0
	if m := reSpellRange.FindStringSubmatch(c.Range); m != nil {
		feet, _ = strconv.Atoi(m[1])
		if m[2] != "" {
			feet *= 5280
		}
	}
	return slices.Contains(c.Classes, classSlug) && reSpellDamage.MatchString(c.Desc) &&
		(!attack || strings.Contains(c.Desc, "spell attack")) && feet >= minRange
}

// MeetsPrerequisite reports whether a sheet meets a feat's or class option's
// prerequisite ("Level 4+, Strength or Dexterity 13+", "Fighting Style
// Feature", "Level 5+ Warlock, Pact of the Blade Invocation", "a Warlock
// Cantrip That Deals Damage"); every clause must hold. level is the character
// level for feats and the class level for class options; cantrips are the
// sheet's class cantrips. slug turns a class name into its compendium slug.
func MeetsPrerequisite(prereq string, sh Sheet, level int, cantrips []Cantrip, slug func(string) string) bool {
	for _, clause := range strings.FieldsFunc(prereq, func(r rune) bool { return r == ',' || r == ';' }) {
		clause = strings.TrimSpace(clause)
		var scores []string // "Strength or Dexterity" of "… 13+"
		ms := rePrereqScore.FindStringSubmatch(clause)
		if ms != nil {
			scores = AbilityKeys(strings.ReplaceAll(ms[1], " or ", ", "))
		}
		if m := rePrereqLevel.FindStringSubmatch(clause); m != nil {
			if n, _ := strconv.Atoi(m[1]); level < n {
				return false
			}
		} else if len(scores) > 0 {
			n, _ := strconv.Atoi(ms[2])
			if !slices.ContainsFunc(scores, func(k string) bool { return sh.Score(k) >= n }) {
				return false
			}
		} else if m := rePrereqFeature.FindStringSubmatch(clause); m != nil {
			names := strings.Split(m[1], " or ")
			if !slices.ContainsFunc(sh.Features, func(f Feature) bool { return slices.Contains(names, f.Name) }) {
				return false
			}
		} else if m := rePrereqInvocation.FindStringSubmatch(clause); m != nil {
			if !slices.ContainsFunc(sh.Features, func(f Feature) bool {
				kind, name, _ := strings.Cut(f.Name, ": ")
				return strings.HasSuffix(kind, " Invocation") && name == m[1]
			}) {
				return false
			}
		} else if m := rePrereqCantrip.FindStringSubmatch(clause); m != nil {
			if !slices.ContainsFunc(cantrips, func(c Cantrip) bool { return CantripFits(c, slug(m[1]), m[2] != "", 0) }) {
				return false
			}
		}
		// shortcut: other clauses ("Proficiency with Heavy armor") pass unchecked; add a rule when a source uses one.
	}
	return true
}
