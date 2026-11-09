package dnd5e

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
	"strings"
	"unicode"
)

// Ref kinds: what a sheet takes from the compendium and pins to a source.
const (
	RefClass      = "class"
	RefSubclass   = "subclass"
	RefSpell      = "spell"
	RefSpecies    = "species"
	RefBackground = "background"
	RefFeat       = "feat"
	RefArmor      = "armor"
)

// RefKinds lists the ref kinds the source command takes.
var RefKinds = []string{RefClass, RefSubclass, RefSpell, RefSpecies, RefBackground, RefFeat, RefArmor}

// SourceRef names a ref on the sheet and the source to pin it to.
type SourceRef struct {
	Kind, Slug, Source string
}

// SwitchPlan is what re-pointing a class or subclass to another copy does
// to the sheet's features, for the dialog to list before confirming.
type SwitchPlan struct {
	Replace, Add, Keep []SwitchFeature
}

// SwitchFeature is one feature in a SwitchPlan. Dropped marks a replaced
// feature the new copy lacks: it is removed.
type SwitchFeature struct {
	Name    string
	Level   int
	Dropped bool
}

// featureSwitch is a SwitchPlan by sheet feature index, for applying it.
type featureSwitch struct {
	class   int
	replace map[int]*ClassFeature // nil drops the feature
	keep    []int
	add     []Feature
}

// PlanSourceSwitch previews re-pointing the class or subclass slug from old
// (nil when that copy is gone) to next. Other kinds change no features.
func PlanSourceSwitch(s Sheet, kind, slug string, old *ClassInfo, next ClassInfo) SwitchPlan {
	var p SwitchPlan
	sw, err := planSwitch(s, kind, slug, old, next)
	if err != nil {
		return p
	}
	for i, to := range sw.replace {
		f := s.Features[i]
		p.Replace = append(p.Replace, SwitchFeature{Name: f.Name, Level: f.Level, Dropped: to == nil})
	}
	slices.SortFunc(p.Replace, func(a, b SwitchFeature) int { return cmp.Or(a.Level-b.Level, strings.Compare(a.Name, b.Name)) })
	for _, i := range sw.keep {
		p.Keep = append(p.Keep, SwitchFeature{Name: s.Features[i].Name, Level: s.Features[i].Level})
	}
	for _, f := range sw.add {
		p.Add = append(p.Add, SwitchFeature{Name: f.Name, Level: f.Level})
	}
	return p
}

func planSwitch(s Sheet, kind, slug string, old *ClassInfo, next ClassInfo) (featureSwitch, error) {
	sw := featureSwitch{class: s.refClass(kind, slug), replace: map[int]*ClassFeature{}}
	if sw.class < 0 {
		return sw, fmt.Errorf("no %s %q on the sheet", kind, slug)
	}
	cl := s.Classes[sw.class]
	// Class and subclass features both carry the class name as Source, so
	// each side skips names the other side's copies list.
	own := func(c ClassInfo) []ClassFeature { return c.Features }
	other := func(c ClassInfo) []ClassFeature { return c.SubclassFeatures }
	if kind == RefSubclass {
		own, other = other, own
	}
	others := featureNames(other(next))
	oldDesc := map[string]string{}
	if old != nil {
		maps.Copy(others, featureNames(other(*old)))
		for _, f := range own(*old) {
			oldDesc[strings.ToLower(f.Name)] = f.Description
		}
	}
	newByName := map[string]ClassFeature{}
	for _, f := range own(next) {
		newByName[strings.ToLower(f.Name)] = f
	}
	onSheet := map[string]bool{}
	for i, f := range s.Features {
		name := strings.ToLower(f.Name)
		onSheet[name] = true
		if f.Source != cl.Name || f.Level == 0 || others[name] {
			continue
		}
		if desc, ok := oldDesc[name]; !ok || desc != f.Description { // edited, or not from the old copy
			sw.keep = append(sw.keep, i)
			continue
		}
		if nf, ok := newByName[name]; ok {
			sw.replace[i] = &nf
		} else {
			sw.replace[i] = nil
		}
	}
	for _, f := range own(next) {
		if onSheet[strings.ToLower(f.Name)] {
			continue
		}
		// Once, at the first level gained, as LevelUp adds them.
		for _, lvl := range slices.Sorted(slices.Values(f.Levels)) {
			if lvl >= 1 && lvl <= cl.Level {
				sw.add = append(sw.add, Feature{Name: f.Name, Source: cl.Name, Level: lvl, Description: f.Description})
				break
			}
		}
	}
	return sw, nil
}

func featureNames(fs []ClassFeature) map[string]bool {
	out := map[string]bool{}
	for _, f := range fs {
		out[strings.ToLower(f.Name)] = true
	}
	return out
}

// refClass finds the class a class or subclass ref names, or -1.
func (s Sheet) refClass(kind, slug string) int {
	switch kind {
	case RefClass:
		return s.classIndex(slug)
	case RefSubclass:
		// Unpinned sheets know the subclass by name only.
		return slices.IndexFunc(s.Classes, func(c ClassLevel) bool {
			return c.SubclassSlug == slug || c.SubclassSlug == "" && c.Subclass != "" && sameSlug(c.Subclass, slug)
		})
	}
	return -1
}

// sameSlug compares a name with a compendium slug by letters and digits.
func sameSlug(name, slug string) bool {
	alnum := func(s string) string {
		return strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				return unicode.ToLower(r)
			}
			return -1
		}, s)
	}
	return alnum(name) == alnum(slug)
}

// SwitchSource re-points ref to ref.Source. A class or subclass replaces
// its untouched features with env.SourceNew's, adds those of levels taken
// and keeps edited ones; a spell takes env.SourceSpell's name and level;
// the other kinds only change the pin.
func (c *Character) SwitchSource(ref SourceRef, env Env) error {
	s := &c.Sheet
	switch ref.Kind {
	case RefClass, RefSubclass:
		next := env.SourceNew
		if ref.Kind == RefClass && next.Slug != ref.Slug || ref.Kind == RefSubclass && next.SubclassSlug != ref.Slug {
			return fmt.Errorf("unknown %s %q", ref.Kind, ref.Slug)
		}
		sw, err := planSwitch(*s, ref.Kind, ref.Slug, env.SourceOld, next)
		if err != nil {
			return err
		}
		s.applySwitch(sw, next)
		cl := &s.Classes[sw.class]
		if ref.Kind == RefSubclass {
			cl.SubclassSlug, cl.SubclassSource = ref.Slug, ref.Source
			if cl.Caster == "" || cl.Caster == CasterNone {
				cl.Caster = cmp.Or(next.Caster, CasterNone)
				s.RecomputeSlots()
			}
			return nil
		}
		cl.HitDie, cl.Caster, cl.Source = next.HitDie, cmp.Or(next.Caster, CasterNone), ref.Source
		s.RecomputeSlots()
	case RefSpell:
		i := slices.IndexFunc(s.Spells, func(sp Spell) bool { return sp.Slug == ref.Slug })
		if i < 0 || env.SourceSpell.Slug != ref.Slug {
			return fmt.Errorf("spell %q is not on the sheet", ref.Slug)
		}
		sp := &s.Spells[i]
		sp.Name, sp.Level, sp.Source = env.SourceSpell.Name, env.SourceSpell.Level, ref.Source
	case RefSpecies, RefBackground:
		pin, have := &s.SpeciesSource, s.Species
		if ref.Kind == RefBackground {
			pin, have = &s.BackgroundSource, s.Background
		}
		if have == "" {
			return fmt.Errorf("no %s on the sheet", ref.Kind)
		}
		*pin = ref.Source
	case RefFeat:
		_, pinned := s.FeatSources[ref.Slug]
		if !pinned && !slices.ContainsFunc(s.Feats, func(n string) bool { return sameSlug(n, ref.Slug) }) {
			return fmt.Errorf("no feat %q on the sheet", ref.Slug)
		}
		if s.FeatSources == nil {
			s.FeatSources = map[string]string{}
		}
		s.FeatSources[ref.Slug] = ref.Source
	case RefArmor:
		if s.Armor == nil || s.Armor.Slug != ref.Slug {
			return fmt.Errorf("armor %q is not worn", ref.Slug)
		}
		s.Armor.Source = ref.Source
	default:
		return fmt.Errorf("unknown ref kind %q", ref.Kind)
	}
	return nil
}

// applySwitch carries out sw with next's features.
func (s *Sheet) applySwitch(sw featureSwitch, next ClassInfo) {
	var drop []int
	for i, to := range sw.replace {
		if to == nil {
			drop = append(drop, i)
			continue
		}
		// By index: an older sheet may still hold a feature once per level.
		old := s.Features[i]
		f := Feature{Name: old.Name, Source: old.Source, Level: old.Level, Description: to.Description}
		e := ParseEffects(to.Name, to.Description, next.Table, next.Name)
		s.putFeature(i, f, e.Apply(&f))
		s.RemoveProfsFrom(old.Name)
		s.AddProfs(old.Name, e.Skills, e.Tools...)
	}
	slices.Sort(drop)
	for _, i := range slices.Backward(drop) {
		name := s.Features[i].Name
		s.Features = slices.Delete(s.Features, i, i+1)
		if !slices.ContainsFunc(s.Features, func(f Feature) bool { return f.Name == name }) {
			s.Modifiers = slices.DeleteFunc(s.Modifiers, func(m CustomMod) bool { return m.Source == name })
			s.RemoveProfsFrom(name)
		}
	}
	for _, f := range sw.add {
		s.AddFeature(f, ParseEffects(f.Name, f.Description, next.Table, next.Name))
	}
	s.ScaleSpeeds(next, s.Classes[sw.class].Level)
}

// SubclassRef names a class on the sheet and the subclass copy to give it.
type SubclassRef struct {
	Class, Slug, Source string
}

// ChangeSubclass swaps ref.Class's subclass for env.SourceNew's: untouched
// features of the current one (env.SourceOld) leave with the proficiencies
// they gave, the new one's arrive for every level taken. Spells and option
// picks the old one gave are not tracked, so they stay.
func (c *Character) ChangeSubclass(ref SubclassRef, env Env) error {
	s := &c.Sheet
	i := s.classIndex(ref.Class)
	if i < 0 || s.Classes[i].Subclass == "" {
		return fmt.Errorf("no subclass of %q to change", ref.Class)
	}
	next := env.SourceNew
	if next.Slug != ref.Class || next.SubclassSlug != ref.Slug || next.SubclassName == "" {
		return fmt.Errorf("unknown subclass %q", ref.Slug)
	}
	cl := s.Classes[i]
	// An unpinned subclass is found by its name.
	sw, err := planSwitch(*s, RefSubclass, cmp.Or(cl.SubclassSlug, cl.Subclass), env.SourceOld, next)
	if err != nil {
		return err
	}
	if sw.class != i {
		return fmt.Errorf("no subclass of %q to change", ref.Class)
	}
	s.applySwitch(sw, next)
	s.Classes[i].Subclass, s.Classes[i].SubclassSlug, s.Classes[i].SubclassSource = next.SubclassName, ref.Slug, ref.Source
	s.Classes[i].Caster = cmp.Or(next.Caster, CasterNone)
	s.RecomputeSlots()
	return nil
}

// PinsSource reports whether the sheet pins any compendium ref to the source.
func (s Sheet) PinsSource(key string) bool {
	if s.SpeciesSource == key || s.BackgroundSource == key || s.Armor != nil && s.Armor.Source == key {
		return true
	}
	return slices.ContainsFunc(s.Classes, func(c ClassLevel) bool { return c.Source == key || c.SubclassSource == key }) ||
		slices.ContainsFunc(s.Spells, func(sp Spell) bool { return sp.Source == key }) ||
		slices.Contains(slices.Collect(maps.Values(s.FeatSources)), key)
}
