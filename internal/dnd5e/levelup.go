package dnd5e

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

// ClassInfo is what level-up needs from a compendium class.
type ClassInfo struct {
	Slug     string
	Name     string
	HitDie   int
	Caster   string
	Saves    []string
	Features []ClassFeature
	// SubclassFeatures are the features of the character's subclass in it.
	SubclassFeatures []ClassFeature
	// Table holds class table columns (Cantrips, Prepared Spells, Rages...)
	// by column name, then class level.
	Table map[string]map[int]string
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

// asiFeature and epicBoon are the feature names that grant a feat or an
// ability score increase, as named in the 2024 class tables.
const (
	asiFeature = "Ability Score Improvement"
	epicBoon   = "Epic Boon"
)

// Plan computes what gaining a level in class would do.
func (s Sheet) Plan(class ClassInfo) (LevelUpPlan, error) {
	if s.Level() >= 20 {
		return LevelUpPlan{}, errors.New("already level 20")
	}
	p := LevelUpPlan{Class: class, NewLevel: s.Level() + 1, Columns: map[string]string{}}
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
		case f.Name == asiFeature, f.Name == epicBoon:
			p.ASI = true
		case f.Subclass, strings.HasSuffix(f.Name, " Subclass"):
			p.NeedsSubclass = i < 0 || s.Classes[i].Subclass == ""
		}
	}
	for col, byLevel := range class.Table {
		if v, ok := byLevel[p.ClassLevel]; ok {
			p.Columns[col] = v
		}
	}

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
			return err
		}
	} else if len(ch.ASI) > 0 || ch.Feat != "" {
		return errors.New("this level grants no ability score improvement")
	}
	s.CatchUpUses(ch.Class)
	for _, f := range plan.Features {
		feat := Feature{Name: f.Name, Source: ch.Class.Name, Level: plan.ClassLevel, Description: f.Description}
		mods := ParseEffects(f.Name, f.Description, ch.Class.Table, ch.Class.Name).Apply(&feat)
		s.Features = append(s.Features, feat)
		s.Modifiers = append(s.Modifiers, mods...)
	}
	for _, sp := range ch.Spells {
		if !slices.ContainsFunc(s.Spells, func(x Spell) bool { return x.Slug == sp.Slug }) {
			s.Spells = append(s.Spells, sp)
		}
	}

	if err := next.Validate(); err != nil {
		return err
	}
	*c = next
	return nil
}

// EffectsVersion goes up whenever ParseEffects learns to read more uses, so
// saved sheets catch up once at startup.
const EffectsVersion = 1

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

func (s Sheet) classIndex(slug string) int {
	return slices.IndexFunc(s.Classes, func(c ClassLevel) bool { return c.Slug == slug })
}

// withClassLevel returns a copy of s with one more level in class.
func (s Sheet) withClassLevel(class ClassInfo, subclass string) Sheet {
	s.Classes = slices.Clone(s.Classes)
	if i := s.classIndex(class.Slug); i >= 0 {
		s.Classes[i].Level++
		if subclass != "" {
			s.Classes[i].Subclass = subclass
		}
		return s
	}
	caster := class.Caster
	if caster == "" {
		caster = CasterNone
	}
	s.Classes = append(s.Classes, ClassLevel{
		Slug: class.Slug, Name: class.Name, Subclass: subclass, Level: 1,
		HitDie: class.HitDie, Caster: caster,
	})
	return s
}
