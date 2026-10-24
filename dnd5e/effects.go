package dnd5e

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/perm"
)

// FeatureEffects is what a feature's rules text says the sheet can automate: the
// action using it takes, how often it can be used, unarmored AC and fixed
// stat changes. Anything else stays in the description.
type FeatureEffects struct {
	Action    string
	Uses      Uses
	Unarmored string
	NoShield  bool
	Mods      []CustomMod
}

var (
	reAction     = regexp.MustCompile(`(?i)\bas an? (bonus action|reaction|magic action|action)\b|\btake (?:a|the) (magic) action\b|\buse your (reaction)\b`)
	reUsesAbil   = regexp.MustCompile(`(?i)a number of times equal to your (strength|dexterity|constitution|intelligence|wisdom|charisma) modifier`)
	reUsesProf   = regexp.MustCompile(`(?i)a number of times equal to your proficiency bonus`)
	reUsesCount  = regexp.MustCompile(`(?i)\byou can use (?:this feature|this trait|it|[a-z' ]{1,40}?) (twice|three times|four times)\b`)
	reUsesOnce   = regexp.MustCompile(`(?i)\bonce you use this (?:feature|trait|ability)\b`)
	reUsesColumn = regexp.MustCompile(`(?i)number of times shown for your [a-z]+ level in the ([a-z' ]+?) column`)
	// reUsesShown is a pool that grows by level: Focus Points, Sorcery Points, Channel Divinity uses.
	reUsesShown = regexp.MustCompile(`(?i)\b(?:points?|uses)\b[^.]*?as shown in the ([a-z' ]+?) column`)
	reShortOne  = regexp.MustCompile(`(?i)regain one (?:of its )?expended uses? when you finish a short rest`)
	reShortRest = regexp.MustCompile(`(?i)finish a short or long rest`)
	reSpeed     = regexp.MustCompile(`(?i)your speed increases by (\d+) feet([^.]*)`)
	reUnarmored = regexp.MustCompile(`(?i)base armor class equals 10 plus your dexterity and (strength|constitution|intelligence|wisdom|charisma) modifiers`)
)

var abilityByName = map[string]string{
	"strength": "str", "dexterity": "dex", "constitution": "con",
	"intelligence": "int", "wisdom": "wis", "charisma": "cha",
}

var actionKinds = map[string]string{"action": "action", "bonus action": "bonus", "reaction": "reaction", "magic action": "magic", "magic": "magic"}

var timesWords = map[string]int{"twice": 2, "three times": 3, "four times": 4}

// ParseEffects reads a feature's rules text. table is its class table
// (column, then class level), used when uses follow a column such as Rages;
// class names that class for those uses.
func ParseEffects(name, desc string, table map[string]map[int]string, class string) FeatureEffects {
	var e FeatureEffects
	if m := reAction.FindStringSubmatch(desc); m != nil {
		e.Action = actionKinds[strings.ToLower(strings.Join(m[1:], ""))]
	}
	switch {
	case reUsesAbil.MatchString(desc):
		e.Uses.Ability = abilityByName[strings.ToLower(reUsesAbil.FindStringSubmatch(desc)[1])]
	case reUsesProf.MatchString(desc):
		e.Uses.Prof = true
	case reUsesColumn.MatchString(desc) || reUsesShown.MatchString(desc):
		m := reUsesColumn.FindStringSubmatch(desc)
		if m == nil {
			m = reUsesShown.FindStringSubmatch(desc)
		}
		e.Uses.ByLevel, e.Uses.Class = columnInts(table, strings.TrimSpace(m[1])), class
	case reUsesCount.MatchString(desc):
		e.Uses.Fixed = timesWords[strings.ToLower(reUsesCount.FindStringSubmatch(desc)[1])]
	case reUsesOnce.MatchString(desc):
		e.Uses.Fixed = 1
	}
	if e.Uses.Limited() {
		switch {
		case reShortOne.MatchString(desc):
			e.Uses.Reset = "short-one"
		case reShortRest.MatchString(desc):
			e.Uses.Reset = "short"
		default:
			e.Uses.Reset = "long"
		}
	}
	if m := reSpeed.FindStringSubmatch(desc); m != nil {
		n, _ := strconv.Atoi(m[1])
		e.Mods = append(e.Mods, CustomMod{Target: ModSpeed, Value: n, Note: name, Source: name, When: armorCondition(m[2])})
	}
	if m := reUnarmored.FindStringSubmatchIndex(desc); m != nil {
		e.Unarmored = abilityByName[strings.ToLower(desc[m[2]:m[3]])]
		sentenceStart := strings.LastIndex(desc[:m[0]], ".") + 1
		e.NoShield = strings.Contains(strings.ToLower(desc[sentenceStart:m[0]]), "wielding a shield")
	}
	return e
}

// armorCondition reads "while you aren't wearing Heavy armor" and the like.
func armorCondition(text string) string {
	t := strings.ToLower(text)
	switch {
	case strings.Contains(t, "heavy armor"):
		return "no-heavy-armor"
	case strings.Contains(t, "armor") && strings.Contains(t, "shield"):
		return "unarmored-no-shield"
	case strings.Contains(t, "armor"):
		return "unarmored"
	}
	return ""
}

// columnInts reads a class table column of counts ("2", "Unlimited" is
// skipped) by class level.
func columnInts(table map[string]map[int]string, col string) map[int]int {
	for name, byLevel := range table {
		if !strings.EqualFold(name, col) {
			continue
		}
		out := map[int]int{}
		for level, v := range byLevel {
			if n, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
				out[level] = n
			}
		}
		return out
	}
	return nil
}

// Apply puts the effects on a feature and returns its modifiers.
func (e FeatureEffects) Apply(f *Feature) []CustomMod {
	f.Action, f.Uses, f.Unarmored, f.NoShield = e.Action, e.Uses, e.Unarmored, e.NoShield
	return e.Mods
}

// Summary describes the effects in a few words each, for previews.
func (e FeatureEffects) Summary() []string {
	var out []string
	if e.Action != "" {
		out = append(out, ActionLabels[e.Action])
	}
	if l := UsesLabel(e.Uses); l != "" {
		out = append(out, l)
	}
	if e.Unarmored != "" {
		out = append(out, "AC 10 + Dex + "+AbilityNames[e.Unarmored][:3]+" without armor")
	}
	for _, m := range e.Mods {
		out = append(out, fmt.Sprintf("%s %s %s", m.Target, Signed(m.Value), WhenLabels[m.When]))
	}
	return out
}

// ActionLabels name a feature's action kinds.
var ActionLabels = map[string]string{"action": "Action", "bonus": "Bonus Action", "reaction": "Reaction", "magic": "Magic action"}

// UsesLabel says how often something can be used, e.g. "Charisma
// modifier uses per long rest".
func UsesLabel(u Uses) string {
	var n string
	switch {
	case u.Fixed > 0:
		n = fmt.Sprintf("%d uses", u.Fixed)
		if u.Fixed == 1 {
			n = "1 use"
		}
	case u.Ability != "":
		n = AbilityNames[u.Ability] + " modifier uses"
	case u.Prof:
		n = "Proficiency bonus uses"
	case len(u.ByLevel) > 0:
		n = "uses by " + u.Class + " level"
	default:
		return ""
	}
	return n + map[string]string{"short": " per short rest", "long": " per long rest", "short-one": " per long rest (one back on a short rest)"}[u.Reset]
}

// SaveFeature adds the feature, or replaces the one of the same name
// (keeping where it came from and the uses spent), and replaces its stat
// changes with mods.
func SaveFeature(f Feature, mods []CustomMod) Command {
	return NewCommand("feature add "+strconv.Quote(f.Name), perm.ScopeSheet, func(c *Character) error {
		c.Sheet.PutFeature(f, mods)
		return nil
	})
}

// PutFeature is SaveFeature's change, for code that builds sheets.
func (s *Sheet) PutFeature(f Feature, mods []CustomMod) {
	if i := slices.IndexFunc(s.Features, func(x Feature) bool { return strings.EqualFold(x.Name, f.Name) }); i >= 0 {
		old := s.Features[i]
		f.Name, f.Source, f.Level = old.Name, cmp.Or(f.Source, old.Source), cmp.Or(f.Level, old.Level)
		f.Uses.Used = old.Uses.Used
		s.Features[i] = f
	} else {
		s.Features = append(s.Features, f)
	}
	mods = slices.Clone(mods)
	for i := range mods {
		mods[i].Note, mods[i].Source = f.Name, f.Name
	}
	s.Modifiers = append(slices.DeleteFunc(s.Modifiers, func(m CustomMod) bool { return m.Source == f.Name }), mods...)
}
