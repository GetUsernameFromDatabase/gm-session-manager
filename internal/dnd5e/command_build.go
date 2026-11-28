package dnd5e

import (
	"cmp"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/perm"
)

func parseSet(args []string) (Command, error) {
	const usage = "set name|species|background|backstory|vision <text> | set ac|speed <n> | set size <size> | set ability <str..cha> <n>"
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	field, value := strings.ToLower(args[0]), strings.Join(args[1:], " ")
	switch field {
	case "size":
		size := SizeNamed(value)
		if size == "" {
			return Command{}, fmt.Errorf("size must be one of %s", strings.Join(Sizes, ", "))
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Size = size
			return nil
		}), nil
	case "vision":
		if strings.EqualFold(value, "none") {
			value = ""
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Vision = value
			return nil
		}), nil
	case "name", "species", "background", "backstory":
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			switch field {
			case "name":
				c.Name = value
			case "species":
				c.Sheet.Species = value
			case "backstory":
				c.Sheet.Backstory = value
			default:
				c.Sheet.Background = value
			}
			return nil
		}), nil
	case "ac", "speed":
		n, err := strconv.Atoi(value)
		if err != nil {
			return Command{}, fmt.Errorf("%s must be a number", field)
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if field == "ac" {
				c.Sheet.ACBonus = n
			} else {
				c.Sheet.Speed = n
			}
			return nil
		}), nil
	case "ability":
		if len(args) != 3 {
			return Command{}, fmt.Errorf("usage: %s", usage)
		}
		ab := strings.ToLower(args[1])
		n, err := strconv.Atoi(args[2])
		if err != nil || !slices.Contains(Abilities, ab) {
			return Command{}, errors.New("usage: set ability str 16")
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			c.Sheet.Scores[ab] = n // Apply moves hit points with Constitution
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

// Proficiency levels for "prof set".
const (
	ProfNone       = "none"
	ProfProficient = "proficient"
	ProfExpertise  = "expertise"
)

// What "prof add|remove" changes, besides ModSkill and ModSave.
const (
	profKindExpertise     = "expertise"
	profKindOther         = "other"
	profKindLanguage      = "language"
	profKindArmor         = "armor"
	profKindResistance    = "resistance"
	profKindImmunity      = "immunity"
	profKindVulnerability = "vulnerability"
)

// ProfLevel is a skill's proficiency level.
func (s Sheet) ProfLevel(skill string) string {
	switch {
	case slices.Contains(s.Expertise, skill):
		return ProfExpertise
	case slices.Contains(s.SkillProfs, skill):
		return ProfProficient
	}
	return ProfNone
}

// setIn adds value to or removes it from list.
func setIn(list *[]string, value string, on bool) {
	i := slices.IndexFunc(*list, func(x string) bool { return strings.EqualFold(x, value) })
	switch {
	case on && i < 0:
		*list = append(*list, value)
	case !on && i >= 0:
		*list = slices.Delete(*list, i, i+1)
	}
}

// parseProfSet sets a skill or save to a level in one step, so going from
// none to expertise is one command: prof set skill|save <key> <level>.
func parseProfSet(args []string) (Command, error) {
	const usage = "prof set skill <skill> none|proficient|expertise | prof set save <ability> none|proficient"
	if len(args) != 4 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	kind, key, level := strings.ToLower(args[1]), strings.ToLower(args[2]), strings.ToLower(args[3])
	switch {
	case kind == ModSkill && slices.ContainsFunc(Skills, func(sk Skill) bool { return sk.Key == key }) &&
		slices.Contains([]string{ProfNone, ProfProficient, ProfExpertise}, level):
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			setIn(&c.Sheet.SkillProfs, key, level != ProfNone)
			setIn(&c.Sheet.Expertise, key, level == ProfExpertise)
			if level == ProfNone {
				delete(c.Sheet.ProfSources, key)
			}
			return nil
		}), nil
	case kind == ModSave && slices.Contains(Abilities, key) && (level == ProfNone || level == ProfProficient):
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			setIn(&c.Sheet.SaveProfs, key, level == ProfProficient)
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

func parseProf(args []string) (Command, error) {
	const usage = "prof add|remove skill|save|expertise|other|language|armor|resistance|immunity|vulnerability <value> | prof set skill|save <key> <level>"
	if len(args) < 3 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op := strings.ToLower(args[0])
	kind := strings.ToLower(args[1])
	value := strings.Join(args[2:], " ")
	if op == verbSet {
		return parseProfSet(args)
	}
	if op != verbAdd && op != verbRemove {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	var pick func(s *Sheet) *[]string
	switch kind {
	case ModSkill, profKindExpertise:
		value = strings.ToLower(value)
		if !slices.ContainsFunc(Skills, func(sk Skill) bool { return sk.Key == value }) {
			return Command{}, fmt.Errorf("unknown skill %q", value)
		}
		pick = func(s *Sheet) *[]string { return &s.SkillProfs }
		if kind == profKindExpertise {
			pick = func(s *Sheet) *[]string { return &s.Expertise }
		}
	case ModSave:
		value = strings.ToLower(value)
		if !slices.Contains(Abilities, value) {
			return Command{}, fmt.Errorf("unknown ability %q", value)
		}
		pick = func(s *Sheet) *[]string { return &s.SaveProfs }
	case profKindOther:
		pick = func(s *Sheet) *[]string { return &s.OtherProfs }
	case profKindLanguage:
		pick = func(s *Sheet) *[]string { return &s.Languages }
	case profKindArmor:
		if value = ArmorKindNamed(value); value == "" {
			return Command{}, fmt.Errorf(
				"armor training is light, medium, heavy or shield, not %q",
				strings.Join(args[2:], " "),
			)
		}
		pick = func(s *Sheet) *[]string { return &s.ArmorTraining }
	case profKindResistance, profKindImmunity, profKindVulnerability:
		value = strings.ToLower(value)
		if !slices.Contains(DamageTypes, value) {
			return Command{}, fmt.Errorf("unknown damage type %q", value)
		}
		pick = map[string]func(s *Sheet) *[]string{
			profKindResistance:    func(s *Sheet) *[]string { return &s.Resistances },
			profKindImmunity:      func(s *Sheet) *[]string { return &s.Immunities },
			profKindVulnerability: func(s *Sheet) *[]string { return &s.Vulnerabilities },
		}[kind]
	default:
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
		setIn(pick(&c.Sheet), value, op == verbAdd)
		c.Sheet.sortArmorTraining()
		if op == verbRemove { // removed by hand: no longer the feature's
			switch kind {
			case ModSkill, profKindOther:
				delete(c.Sheet.ProfSources, value)
			case profKindArmor:
				delete(c.Sheet.ProfSources, armorSourceKey(value))
			}
		}
		return nil
	}), nil
}

func parseSpell(args []string) (Command, error) {
	const usage = "spell add|remove|prepare|unprepare <slug>"
	if err := want(args, 2, usage); err != nil {
		return Command{}, err
	}
	op, slug := strings.ToLower(args[0]), strings.ToLower(args[1])
	find := func(c *Character) int {
		return slices.IndexFunc(c.Sheet.Spells, func(s Spell) bool { return s.Slug == slug })
	}
	switch op {
	case verbAdd:
		cmd := scoped(perm.ScopeSheet, func(c *Character, env Env, _ *Effects) error {
			if env.Spell.Slug != slug {
				return fmt.Errorf("unknown spell %q", slug)
			}
			if find(c) < 0 {
				c.Sheet.Spells = append(c.Sheet.Spells, env.Spell)
			}
			return nil
		})
		cmd.SpellSlug = slug
		return cmd, nil
	case verbRemove:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := find(c); i >= 0 {
				c.Sheet.Spells = slices.Delete(c.Sheet.Spells, i, i+1)
			}
			return nil
		}), nil
	case verbPrepare, verbUnprepare:
		cmd := scoped(perm.ScopeSheet, func(c *Character, env Env, _ *Effects) error {
			i := find(c)
			if i < 0 {
				return fmt.Errorf("spell %q is not on the sheet", slug)
			}
			if op == verbPrepare && !c.Sheet.Spells[i].Prepared {
				if err := c.Sheet.checkPrepareCap(c.Sheet.Spells[i], env.PreparedCaps); err != nil {
					return err
				}
			}
			c.Sheet.Spells[i].Prepared = op == verbPrepare
			return nil
		})
		cmd.PrepareCaps = op == verbPrepare
		return cmd, nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

// checkPrepareCap refuses preparing sp when its class has prepared as many
// spells as caps allows. Cantrips, lineage and feat spells are never capped.
func (s Sheet) checkPrepareCap(sp Spell, caps map[string]int) error {
	counted := func(sp Spell) bool { return sp.Level > 0 && sp.Feat == "" && sp.Ability == "" }
	// shortcut: untagged spells on a multiclass sheet are not capped,
	// tag them by class if that matters.
	classOf := func(sp Spell) string {
		if sp.Class == "" && len(s.Classes) == 1 {
			return s.Classes[0].Slug
		}
		return sp.Class
	}
	class := classOf(sp)
	limit, ok := caps[class]
	if !counted(sp) || class == "" || !ok {
		return nil
	}
	n := 0
	for _, o := range s.Spells {
		if o.Prepared && counted(o) && classOf(o) == class {
			n++
		}
	}
	if n < limit {
		return nil
	}
	name := class
	if i := slices.IndexFunc(
		s.Classes,
		func(cl ClassLevel) bool { return cl.Slug == class },
	); i >= 0 {
		name = s.Classes[i].Name
	}
	noun := "spells"
	if limit == 1 {
		noun = "spell"
	}
	return fmt.Errorf("%s prepares %d %s at this level: unprepare one first", name, limit, noun)
}

func parseFeature(args []string) (Command, error) {
	const usage = `feature add "Name" [description] | feature set "Name" key=value… | feature remove "Name"`
	if len(args) < 2 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, name := strings.ToLower(args[0]), args[1]
	find := func(c *Character) int {
		return slices.IndexFunc(
			c.Sheet.Features,
			func(f Feature) bool { return strings.EqualFold(f.Name, name) },
		)
	}
	switch op {
	case verbAdd:
		desc := strings.Join(args[2:], " ")
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := find(c); i >= 0 {
				c.Sheet.Features[i].Description = desc
			} else {
				c.Sheet.Features = append(c.Sheet.Features, Feature{Name: name, Description: desc})
			}
			return nil
		}), nil
	case verbRemove:
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := find(c); i >= 0 {
				name := c.Sheet.Features[i].Name
				c.Sheet.Features = slices.Delete(c.Sheet.Features, i, i+1)
				c.Sheet.Modifiers = slices.DeleteFunc(
					c.Sheet.Modifiers,
					func(m CustomMod) bool { return m.Source == name },
				)
				c.Sheet.dropFeat(name)
			}
			return nil
		}), nil
	case verbSet:
		set, err := parseFeatureSet(args[2:])
		if err != nil {
			return Command{}, err
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := find(c)
			if i < 0 {
				return fmt.Errorf("no feature %q", name)
			}
			set(&c.Sheet.Features[i])
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

// dropFeat takes the feat named name off the sheet with its feature, and
// its pin once no copy of it ("Magic Initiate (Wizard)") is left.
func (s *Sheet) dropFeat(name string) {
	before := len(s.Feats)
	s.Feats = slices.DeleteFunc(s.Feats, func(n string) bool { return strings.EqualFold(n, name) })
	base, _, _ := strings.Cut(name, " (")
	if len(s.Feats) == before ||
		slices.ContainsFunc(
			s.Feats,
			func(n string) bool { b, _, _ := strings.Cut(n, " ("); return sameSlug(b, base) },
		) {
		return
	}
	maps.DeleteFunc(s.FeatSources, func(k, _ string) bool { return sameSlug(base, k) })
}

// parseFeatureSet reads key=value settings of a feature: action
// (none|action|bonus|reaction|magic), uses (none|n|ability|prof) and reset.
func parseFeatureSet(args []string) (func(f *Feature), error) {
	if len(args) == 0 {
		return nil, errors.New("usage: feature set \"Name\" action=bonus uses=2 reset=long")
	}
	var steps []func(f *Feature)
	for _, a := range args {
		k, v, ok := strings.Cut(strings.ToLower(a), "=")
		switch {
		case !ok:
			return nil, fmt.Errorf("%q is not key=value", a)
		case k == "action" && (v == "none" || ActionLabels[v] != ""):
			steps = append(steps, func(f *Feature) { f.Action = strings.TrimPrefix(v, "none") })
		case k == "uses":
			u, err := parseUses(v)
			if err != nil {
				return nil, err
			}
			steps = append(steps, func(f *Feature) {
				u.Reset = cmp.Or(f.Uses.Reset, ResetLong)
				f.Uses = u
			})
		case k == "reset" && slices.Contains([]string{ResetShort, ResetLong, ResetShortOne}, v):
			steps = append(steps, func(f *Feature) { f.Uses.Reset = v })
		default:
			return nil, fmt.Errorf("unknown setting %q", a)
		}
	}
	return func(f *Feature) {
		for _, s := range steps {
			s(f)
		}
	}, nil
}

// parseUses reads a uses setting: none, a count, an ability key or prof.
func parseUses(v string) (Uses, error) {
	switch {
	case v == "none":
		return Uses{}, nil
	case v == verbProf:
		return Uses{Prof: true}, nil
	case slices.Contains(Abilities, v):
		return Uses{Ability: v}, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil || n < 1 || n > 99 {
		return Uses{}, errors.New("uses must be none, a number, an ability (cha) or prof")
	}
	return Uses{Fixed: n}, nil
}

func parseMod(args []string) (Command, error) {
	const usage = "mod add <target> <+n|-n> <note> | mod remove <target> <note> | mod edit <target> <note> <new target> <+n|-n> <new note>"
	if len(args) < 3 {
		return Command{}, fmt.Errorf("usage: %s", usage)
	}
	op, target := strings.ToLower(args[0]), args[1]
	if !IsModTarget(target) {
		return Command{}, fmt.Errorf("unknown modifier target %q", target)
	}
	switch op {
	case verbAdd:
		if len(args) < 4 {
			return Command{}, fmt.Errorf("usage: %s (the note says where it comes from)", usage)
		}
		m, err := newMod(target, args[2], args[3:])
		if err != nil {
			return Command{}, err
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			if i := findMod(c, m.Target, m.Note); i >= 0 {
				c.Sheet.Modifiers[i] = m
			} else {
				c.Sheet.Modifiers = append(c.Sheet.Modifiers, m)
			}
			return nil
		}), nil
	case verbEdit:
		if len(args) < 6 {
			return Command{}, fmt.Errorf("usage: %s", usage)
		}
		note := args[2]
		m, err := newMod(args[3], args[4], args[5:])
		if err != nil {
			return Command{}, err
		}
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := findMod(c, target, note)
			if i < 0 {
				return fmt.Errorf("no modifier %q on %s", note, target)
			}
			if j := findMod(c, m.Target, m.Note); j >= 0 && j != i {
				return fmt.Errorf("there is already a modifier %q on %s", m.Note, m.Target)
			}
			c.Sheet.Modifiers[i] = m
			return nil
		}), nil
	case verbRemove:
		note := strings.Join(args[2:], " ")
		return scoped(perm.ScopeSheet, func(c *Character, _ Env, _ *Effects) error {
			i := findMod(c, target, note)
			if i < 0 {
				return fmt.Errorf("no modifier %q on %s", note, target)
			}
			c.Sheet.Modifiers = slices.Delete(c.Sheet.Modifiers, i, i+1)
			return nil
		}), nil
	}
	return Command{}, fmt.Errorf("usage: %s", usage)
}

func newMod(target, value string, note []string) (CustomMod, error) {
	if !IsModTarget(target) {
		return CustomMod{}, fmt.Errorf("unknown modifier target %q", target)
	}
	n, err := strconv.Atoi(value)
	if err != nil || n == 0 {
		return CustomMod{}, errors.New("modifier must be a non-zero number such as +1 or -2")
	}
	return CustomMod{Target: target, Value: n, Note: strings.Join(note, " ")}, nil
}

func findMod(c *Character, target, note string) int {
	return slices.IndexFunc(c.Sheet.Modifiers, func(m CustomMod) bool {
		return m.Target == target && strings.EqualFold(m.Note, note)
	})
}
