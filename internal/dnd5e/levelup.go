package dnd5e

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

// ClassInfo is what level-up needs from a compendium class.
type ClassInfo struct {
	Slug string
	Name string
	// Source is the compendium copy's source key; SubclassSlug and
	// SubclassSource are the subclass copy whose SubclassFeatures these are,
	// SubclassName its name.
	Source         string
	SubclassName   string
	SubclassSlug   string
	SubclassSource string
	HitDie         int
	Caster         string
	Saves          []string
	Features       []ClassFeature
	// SubclassFeatures are the features of the character's subclass in it.
	SubclassFeatures []ClassFeature
	// Table holds class table columns (Cantrips, Prepared Spells, Rages...)
	// by column name, then class level.
	Table map[string]map[int]string
}

// ColumnsAt are the class table values at a class level.
func (c ClassInfo) ColumnsAt(level int) map[string]string {
	out := map[string]string{}
	for col, byLevel := range c.Table {
		if v, ok := byLevel[level]; ok {
			out[col] = v
		}
	}
	return out
}

// MasteryColumn is the class table column counting the weapon kinds a class
// masters (Fighter, Barbarian, Paladin, Ranger, Rogue).
const MasteryColumn = "Weapon Mastery"

// Masteries is how many weapon kinds the class masters at a class level.
func (c ClassInfo) Masteries(level int) int {
	n, _ := strconv.Atoi(strings.TrimSpace(c.Table[MasteryColumn][level]))
	return n
}

// MasteryLimit is how many weapon kinds the character may master: what each
// class allows at its level, added up (a table house rule; the rules leave
// multiclassing open). classes holds the compendium classes by slug; a
// missing one allows none.
func (s Sheet) MasteryLimit(classes map[string]ClassInfo) int {
	n := 0
	for _, cl := range s.Classes {
		if info, ok := classes[cl.Slug]; ok {
			n += info.Masteries(cl.Level)
		}
	}
	return n
}

// ClassFeature is a class feature and the class levels it is gained at.
type ClassFeature struct {
	Name        string
	Description string
	Levels      []int
	// Subclass marks the feature where the subclass is chosen.
	Subclass bool
}

// LevelUpPlan previews the next level in a class, for the helper form.
type LevelUpPlan struct {
	Class      ClassInfo
	NewLevel   int // total character level after
	ClassLevel int // level in this class after
	IsNewClass bool
	// AverageHP is the fixed hit point gain including Constitution.
	AverageHP int
	Features  []ClassFeature
	// Columns are the class table values at the new class level.
	Columns       map[string]string
	ASI           bool
	NeedsSubclass bool
	Slots         [9]int
	Pact          int
	PactLevel     int
	// Warnings are rules the level breaks; they inform but never block.
	Warnings []string
}

// ASIFeature and epicBoon are the feature names that grant a feat or an
// ability score increase, as named in the 2024 class tables. ASIFeature is
// also the 2024 feat of that name.
const (
	ASIFeature = "Ability Score Improvement"
	epicBoon   = "Epic Boon"
)

// Plan computes what gaining a level in class would do.
func (s Sheet) Plan(class ClassInfo) (LevelUpPlan, error) {
	if s.Level() >= 20 {
		return LevelUpPlan{}, errors.New("already level 20")
	}
	p := LevelUpPlan{Class: class, NewLevel: s.Level() + 1}
	i := s.classIndex(class.Slug)
	p.IsNewClass = i < 0
	p.ClassLevel = 1
	if i >= 0 {
		p.ClassLevel = s.Classes[i].Level + 1
	}
	p.AverageHP = hpGain(s, class.HitDie, 0)

	for _, f := range append(slices.Clone(class.Features), class.SubclassFeatures...) {
		if !slices.Contains(f.Levels, p.ClassLevel) {
			continue
		}
		p.Features = append(p.Features, f)
		switch {
		case f.Name == ASIFeature, f.Name == epicBoon:
			p.ASI = true
		case f.Subclass, strings.HasSuffix(f.Name, " Subclass"):
			p.NeedsSubclass = i < 0 || s.Classes[i].Subclass == ""
		}
	}
	p.Columns = class.ColumnsAt(p.ClassLevel)

	next := s.withClassLevel(class, "")
	p.Slots, p.Pact, p.PactLevel = SlotMaxima(next.Classes)
	if p.IsNewClass && len(s.Classes) > 0 {
		for _, name := range append([]string{class.Name}, s.classNames()...) {
			if w := s.multiclassWarning(name); w != "" {
				p.Warnings = append(p.Warnings, w)
			}
		}
	}
	return p, nil
}

// DefaultClass is the slug of the class to level next: the one levelled
// last, else the highest.
func (s Sheet) DefaultClass() string {
	if s.classIndex(s.LastClass) >= 0 {
		return s.LastClass
	}
	if len(s.Classes) == 0 {
		return ""
	}
	return slices.MaxFunc(s.Classes, func(a, b ClassLevel) int { return a.Level - b.Level }).Slug
}

func (s Sheet) classNames() []string {
	out := make([]string, len(s.Classes))
	for i, c := range s.Classes {
		out[i] = c.Name
	}
	return out
}

// multiclassPrereqs are the 2024 multiclassing prerequisites: a score of
// 13 in each listed ability, or in any one of them when anyOf is set.
var multiclassPrereqs = map[string]struct {
	abilities []string
	anyOf     bool
}{
	"barbarian": {[]string{"str"}, false},
	"bard":      {[]string{"cha"}, false},
	"cleric":    {[]string{"wis"}, false},
	"druid":     {[]string{"wis"}, false},
	"fighter":   {[]string{"str", "dex"}, true},
	"monk":      {[]string{"dex", "wis"}, false},
	"paladin":   {[]string{"str", "cha"}, false},
	"ranger":    {[]string{"dex", "wis"}, false},
	"rogue":     {[]string{"dex"}, false},
	"sorcerer":  {[]string{"cha"}, false},
	"warlock":   {[]string{"cha"}, false},
	"wizard":    {[]string{"int"}, false},
}

// multiclassWarning explains an unmet multiclassing prerequisite of the
// named class, or is empty. Unknown (homebrew) classes have none.
func (s Sheet) multiclassWarning(class string) string {
	pre, ok := multiclassPrereqs[strings.ToLower(class)]
	if !ok {
		return ""
	}
	met := 0
	names := make([]string, len(pre.abilities))
	for i, a := range pre.abilities {
		names[i] = AbilityNames[a]
		if s.Score(a) >= 13 {
			met++
		}
	}
	if pre.anyOf && met > 0 || !pre.anyOf && met == len(pre.abilities) {
		return ""
	}
	join := " and "
	if pre.anyOf {
		join = " or "
	}
	return fmt.Sprintf("Multiclassing with %s needs %s 13 or higher.", class, strings.Join(names, join))
}

// LevelUpChoice is what the player picked in the helper.
type LevelUpChoice struct {
	Class ClassInfo
	// HPRoll is the hit die the player rolled at the table; 0 takes the
	// fixed average. Ignored at character level 1 (maximum hit die).
	HPRoll   int
	Subclass string
	// ASI names one ability for +2 or two for +1 each.
	ASI []string
	// Feat is taken instead of an ability score increase.
	Feat     string
	FeatDesc string
	Spells   []Spell
}

// LevelUp applies one level. The structured parts are automated (HP, class
// levels, saving throws for a first class, slots, ability increases, feats,
// spells); class features are added as descriptions. Everything stays
// editable afterwards with commands.
func (c *Character) LevelUp(ch LevelUpChoice) error {
	plan, err := c.Sheet.Plan(ch.Class)
	if err != nil {
		return err
	}
	next := c.clone()
	s := &next.Sheet

	if ch.HPRoll < 0 || (ch.HPRoll > ch.Class.HitDie && ch.Class.HitDie > 0) {
		return fmt.Errorf("hit die roll must be 1-%d", ch.Class.HitDie)
	}
	gain := hpGain(*s, ch.Class.HitDie, ch.HPRoll)
	next.MaxHP += gain
	next.HP += gain

	if s.Level() == 0 {
		s.SaveProfs = slices.Clone(ch.Class.Saves)
	}
	*s = s.withClassLevel(ch.Class, ch.Subclass)
	s.LastClass = ch.Class.Slug
	s.RecomputeSlots()

	if plan.ASI {
		if err := applyASI(s, ch); err != nil {
			return ASIError{err}
		}
	} else if len(ch.ASI) > 0 || ch.Feat != "" {
		return errors.New("this level grants no ability score improvement")
	}
	s.CatchUpUses(ch.Class)
	for _, f := range plan.Features {
		feat := Feature{Name: f.Name, Source: ch.Class.Name, Level: plan.ClassLevel, Description: f.Description}
		s.AddFeature(feat, ParseEffects(f.Name, f.Description, ch.Class.Table, ch.Class.Name))
	}
	s.ScaleSpeeds(ch.Class, plan.ClassLevel)
	for _, sp := range ch.Spells {
		s.AddSpell(sp)
	}

	if err := next.Validate(); err != nil {
		return err
	}
	*c = next
	return nil
}

// EffectsVersion goes up whenever saved sheets need catching up once at
// startup: 1 when ParseEffects learned more uses, 2 for species resistances,
// 3 for third casters and weapon masteries, 4 for features added once per
// level, scaling speed bonuses and proficiencies granted by features.
const EffectsVersion = 4

// reNumber is the number in a class table cell ("+15 ft.").
var reNumber = regexp.MustCompile(`[+-]?\d+`)

// ScaleSpeeds sets the speed bonus of the class's features to the class
// table column of the feature's name at the class level (Unarmored Movement:
// "This bonus increases ... as shown on the Monk Features table"). It
// reports whether anything changed.
func (s *Sheet) ScaleSpeeds(class ClassInfo, level int) bool {
	changed := false
	for i, m := range s.Modifiers {
		if m.Target != ModSpeed || s.featureIndex(m.Source, class.Name) < 0 {
			continue
		}
		for col, byLevel := range class.Table {
			if !strings.EqualFold(col, m.Source) {
				continue
			}
			if n, err := strconv.Atoi(reNumber.FindString(byLevel[level])); err == nil && n != m.Value {
				s.Modifiers[i].Value, changed = n, true
			}
		}
	}
	return changed
}

// DedupeFeatures collapses the copies older level-ups added of a feature
// listed at several levels (same name and source) into the first, and
// their stacked modifiers into one. It reports whether anything changed.
func (s *Sheet) DedupeFeatures() bool {
	type key struct{ name, source string }
	seen, dup := map[key]bool{}, map[string]bool{}
	features := s.Features[:0]
	for _, f := range s.Features {
		k := key{strings.ToLower(f.Name), f.Source}
		if seen[k] {
			dup[f.Name] = true
			continue
		}
		seen[k] = true
		features = append(features, f)
	}
	if len(dup) == 0 {
		return false
	}
	s.Features = features
	kept := map[CustomMod]bool{}
	s.Modifiers = slices.DeleteFunc(s.Modifiers, func(m CustomMod) bool {
		if !dup[m.Source] {
			return false
		}
		m.Value = 0 // copies of one bonus differ by nothing else
		drop := kept[m]
		kept[m] = true
		return drop
	})
	return true
}

// CatchUpUses gives class features without uses the uses their text names,
// for features added before the parser knew them (Focus Points). Uses already
// set, by hand or not, stay. It reports whether anything changed.
func (s *Sheet) CatchUpUses(class ClassInfo) bool {
	changed := false
	for i, f := range s.Features {
		if f.Source == class.Name && !f.Uses.Limited() {
			if u := ParseEffects(f.Name, f.Description, class.Table, class.Name).Uses; u.Limited() {
				s.Features[i].Uses = u
				changed = true
			}
		}
	}
	return changed
}

// CatchUpProfs gives the proficiencies class, subclass and species features
// grant, for features added before the parser knew them. It reports whether
// anything changed.
func (s *Sheet) CatchUpProfs() bool {
	before := len(s.SkillProfs) + len(s.OtherProfs)
	for _, f := range s.Features {
		if f.Source == s.Species || slices.Contains(s.classNames(), f.Source) {
			e := ParseEffects(f.Name, f.Description, nil, "")
			s.AddProfs(f.Name, e.Skills, e.Tools...)
		}
	}
	return len(s.SkillProfs)+len(s.OtherProfs) > before
}

// ASIError is a refusal about the ability score increase or feat pick, so the
// page can show it at that field.
type ASIError struct{ error }

func (e ASIError) Unwrap() error { return e.error }

func applyASI(s *Sheet, ch LevelUpChoice) error {
	switch {
	case ch.Feat != "" && len(ch.ASI) > 0:
		return errors.New("choose a feat or an ability score increase, not both")
	case ch.Feat != "":
		s.Feats = append(s.Feats, ch.Feat)
		s.Features = append(s.Features, Feature{Name: ch.Feat, Source: "Feat", Description: ch.FeatDesc})
	case len(ch.ASI) == 1:
		return raise(s, ch.ASI[0], 2)
	case len(ch.ASI) == 2 && ch.ASI[0] != ch.ASI[1]:
		if err := raise(s, ch.ASI[0], 1); err != nil {
			return err
		}
		return raise(s, ch.ASI[1], 1)
	case len(ch.ASI) == 2:
		return raise(s, ch.ASI[0], 2)
	}
	return nil
}

// raise increases an ability score, never above 20.
func raise(s *Sheet, ability string, by int) error {
	if !slices.Contains(Abilities, ability) {
		return fmt.Errorf("unknown ability %q", ability)
	}
	if s.Scores[ability]+by > 20 {
		return fmt.Errorf("%s cannot go above 20", AbilityNames[ability])
	}
	s.Scores[ability] += by
	return nil
}

// hpGain is the hit point increase for a level: the full hit die at
// character level 1, otherwise the roll (or the fixed average when roll is
// 0), plus the Constitution modifier, minimum 1.
func hpGain(s Sheet, hitDie, roll int) int {
	base := hitDie/2 + 1
	switch {
	case s.Level() == 0:
		base = hitDie
	case roll > 0:
		base = roll
	}
	return max(base+s.Mod("con"), 1)
}

// AddSpell adds sp unless the sheet has it already.
func (s *Sheet) AddSpell(sp Spell) {
	if !slices.ContainsFunc(s.Spells, func(x Spell) bool { return x.Slug == sp.Slug }) {
		s.Spells = append(s.Spells, sp)
	}
}

func (s Sheet) classIndex(slug string) int {
	return slices.IndexFunc(s.Classes, func(c ClassLevel) bool { return c.Slug == slug })
}

// withClassLevel returns a copy of s with one more level in class.
func (s Sheet) withClassLevel(class ClassInfo, subclass string) Sheet {
	s.Classes = slices.Clone(s.Classes)
	if i := s.classIndex(class.Slug); i >= 0 {
		s.Classes[i].Level++
		if s.Classes[i].Source == "" { // an unpinned class pins the copy it levels with
			s.Classes[i].Source = class.Source
		}
		if subclass != "" {
			s.Classes[i].Subclass = subclass
		}
		if cl := &s.Classes[i]; (cl.Caster == "" || cl.Caster == CasterNone) && class.Caster != "" { // a subclass brought spellcasting
			cl.Caster = class.Caster
		}
		if s.Classes[i].Subclass != "" && s.Classes[i].SubclassSlug == "" {
			s.Classes[i].SubclassSlug, s.Classes[i].SubclassSource = class.SubclassSlug, class.SubclassSource
		}
		return s
	}
	caster := class.Caster
	if caster == "" {
		caster = CasterNone
	}
	cl := ClassLevel{
		Slug: class.Slug, Name: class.Name, Subclass: subclass, Level: 1,
		HitDie: class.HitDie, Caster: caster, Source: class.Source,
	}
	if subclass != "" {
		cl.SubclassSlug, cl.SubclassSource = class.SubclassSlug, class.SubclassSource
	}
	s.Classes = append(s.Classes, cl)
	return s
}
