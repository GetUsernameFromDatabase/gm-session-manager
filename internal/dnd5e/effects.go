package dnd5e

import (
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/perm"
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
	// Skills (keys) and Tools are fixed proficiencies the text grants; choices
	// ("of your choice", "Insight, Perception, or Survival") are not.
	Skills []string
	Tools  []string
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
	reSkillProf = regexp.MustCompile(`(?i)proficiency in (?:the )?([a-z ,]+?) skills?\b`)
	reToolProf  = regexp.MustCompile(`(?i)proficiency with ([^.;:]+)`)
	reTool      = regexp.MustCompile(`(?i)^(?:(?:the|an?) )?((?:[a-z']+ ){0,3}(?:kit|tools|supplies|set|utensils))$`)
	reChoice    = regexp.MustCompile(`(?i)\b(?:or|one|two|any|type|choice|choose|this)\b`)
	// reTag is a 5etools tag ({@skill Insight|XPHB}) or a compendium link
	// ([[compendium:skill/insight|Insight]]); both keep their label.
	reTag = regexp.MustCompile(`\{@\w+ ([^|}]*)[^}]*\}|\[\[[^\]|]*\|([^\]]*)\]\]`)
)

var abilityByName = map[string]string{
	"strength": Str, "dexterity": Dex, "constitution": Con,
	"intelligence": Int, "wisdom": Wis, "charisma": Cha,
}

var actionKinds = map[string]string{"action": ActionMain, "bonus action": ActionBonus, "reaction": ActionReaction, "magic action": ActionMagic, "magic": ActionMagic}

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
			e.Uses.Reset = ResetShortOne
		case reShortRest.MatchString(desc):
			e.Uses.Reset = ResetShort
		default:
			e.Uses.Reset = ResetLong
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
	e.Skills, e.Tools = parseProfs(desc)
	return e
}

// parseProfs reads fixed grants: "proficiency in the Insight and Medicine
// skills", "proficiency with the herbalism kit". A list with any word that
// is not a skill, or a tool named with "one", "type", "choice"..., is a
// choice the wizard asks for (skillPick, feat picks), not a grant.
func parseProfs(desc string) (skills, tools []string) {
	text := reTag.ReplaceAllString(desc, "$1$2")
	for _, m := range reSkillProf.FindAllStringSubmatch(text, -1) {
		if reChoice.MatchString(m[1]) {
			continue
		}
		var keys []string
		for part := range strings.SplitSeq(strings.ReplaceAll(m[1], " and ", ","), ",") {
			if part = strings.TrimSpace(part); part == "" {
				continue // "History, and Religion"
			}
			i := slices.IndexFunc(Skills, func(sk Skill) bool { return strings.EqualFold(sk.Name, part) })
			if i < 0 {
				keys = nil
				break
			}
			keys = append(keys, Skills[i].Key)
		}
		skills = append(skills, keys...)
	}
	// "heavy armor and smith's tools", "the disguise kit, the forgery kit":
	// each listed tool, not the armor or weapons beside it.
	for _, m := range reToolProf.FindAllStringSubmatch(text, -1) {
		for part := range strings.SplitSeq(strings.ReplaceAll(m[1], " and ", ","), ",") {
			if t := reTool.FindStringSubmatch(strings.TrimSpace(part)); t != nil && !reChoice.MatchString(t[1]) {
				tools = append(tools, t[1])
			}
		}
	}
	return skills, tools
}

// AddProfs adds skill (by key) and other proficiencies granted by origin,
// recording origin for those the sheet did not have yet.
func (s *Sheet) AddProfs(origin string, skills []string, other ...string) {
	note := func(k string) {
		if origin == "" {
			return
		}
		if s.ProfSources == nil {
			s.ProfSources = map[string]string{}
		}
		s.ProfSources[k] = origin
	}
	for _, k := range skills {
		if !slices.Contains(s.SkillProfs, k) {
			s.SkillProfs = append(s.SkillProfs, k)
			note(k)
		}
	}
	for _, o := range other {
		if !slices.ContainsFunc(s.OtherProfs, func(x string) bool { return strings.EqualFold(x, o) }) {
			s.OtherProfs = append(s.OtherProfs, o)
			note(o)
		}
	}
}

// RemoveProfsFrom removes the proficiencies ProfSources says origin added
// (a feature dropped by a subclass or source change); others stay.
func (s *Sheet) RemoveProfsFrom(origin string) {
	for k, from := range s.ProfSources {
		if from != origin {
			continue
		}
		s.SkillProfs = slices.DeleteFunc(s.SkillProfs, func(x string) bool { return x == k })
		s.Expertise = slices.DeleteFunc(s.Expertise, func(x string) bool { return x == k })
		s.OtherProfs = slices.DeleteFunc(s.OtherProfs, func(x string) bool { return x == k })
		delete(s.ProfSources, k)
	}
}

// armorCondition reads "while you aren't wearing Heavy armor" and the like.
func armorCondition(text string) string {
	t := strings.ToLower(text)
	switch {
	case strings.Contains(t, "heavy armor"):
		return WhenNoHeavyArmor
	case strings.Contains(t, "armor") && strings.Contains(t, "shield"):
		return WhenUnarmoredNoShield
	case strings.Contains(t, "armor"):
		return WhenUnarmored
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
var ActionLabels = map[string]string{ActionMain: "Action", ActionBonus: "Bonus Action", ActionReaction: "Reaction", ActionMagic: "Magic action"}

// Feature.Action values.
const (
	ActionMain     = "action"
	ActionBonus    = "bonus"
	ActionReaction = "reaction"
	ActionMagic    = "magic"
)

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
	return n + map[string]string{ResetShort: " per short rest", ResetLong: " per long rest", ResetShortOne: " per long rest (one back on a short rest)"}[u.Reset]
}

// SaveFeature adds the feature, or replaces the one of the same name
// (keeping where it came from and the uses spent), and replaces its stat
// changes with mods. A feat from the compendium (featKey set, as pinFeat
// keys it) joins the sheet's feats too, pinned to featSource.
func SaveFeature(f Feature, mods []CustomMod, featKey, featSource string) Command {
	return NewCommand("feature add "+strconv.Quote(f.Name), perm.ScopeSheet, func(c *Character) error {
		c.Sheet.PutFeature(f, mods)
		if featKey == "" {
			return nil
		}
		if !slices.ContainsFunc(c.Sheet.Feats, func(n string) bool { return strings.EqualFold(n, f.Name) }) {
			c.Sheet.Feats = append(c.Sheet.Feats, f.Name)
		}
		if c.Sheet.FeatSources == nil {
			c.Sheet.FeatSources = map[string]string{}
		}
		c.Sheet.FeatSources[featKey] = featSource
		return nil
	})
}

// PutFeature is SaveFeature's change, for code that builds sheets.
func (s *Sheet) PutFeature(f Feature, mods []CustomMod) {
	s.putFeature(slices.IndexFunc(s.Features, func(x Feature) bool { return strings.EqualFold(x.Name, f.Name) }), f, mods)
}

// AddFeature puts a gained class, subclass or species feature and its
// effects on the sheet. One of the same name and source already there (a
// feature listed at several levels) is updated, keeping the level first
// gained and the uses spent; its modifiers replace the old ones, never
// stack. The fixed proficiencies it grants are recorded as the feature's.
func (s *Sheet) AddFeature(f Feature, e FeatureEffects) {
	mods := e.Apply(&f)
	i := s.featureIndex(f.Name, f.Source)
	if i >= 0 {
		f.Level = cmp.Or(s.Features[i].Level, f.Level)
	}
	s.putFeature(i, f, mods)
	s.AddProfs(f.Name, e.Skills, e.Tools...)
}

func (s Sheet) featureIndex(name, source string) int {
	return slices.IndexFunc(s.Features, func(x Feature) bool { return strings.EqualFold(x.Name, name) && x.Source == source })
}

// putFeature replaces the feature at i (appends when i < 0) and its mods.
func (s *Sheet) putFeature(i int, f Feature, mods []CustomMod) {
	if i >= 0 {
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
