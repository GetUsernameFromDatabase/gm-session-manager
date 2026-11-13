package web

import (
	"cmp"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/web/views"
)

// skillChoice reads "Choose 2: …" as how many skills and from which keys;
// "any" skills offers them all.
func skillChoice(s string) (int, []string) {
	n, names := compendium.SkillChoice(s)
	return n, dnd5e.SkillKeys(names)
}

// classSkillPick asks for count of a class's skills, leaving out those the
// sheet has from elsewhere (its background), which it shows known; picked
// are ticked at first. Nil when the class picks no skills.
func classSkillPick(sh dnd5e.Sheet, skills []string, count int, picked []string) *views.SkillPick {
	if count == 0 {
		return nil
	}
	p := &views.SkillPick{Picked: picked}
	for _, k := range skills {
		if slices.Contains(sh.SkillProfs, k) {
			p.Known = append(p.Known, k)
		} else {
			p.Options = append(p.Options, k)
		}
	}
	p.Count = min(count, len(p.Options))
	return p
}

// backgroundSkillPick re-asks the class step's skill pick when the
// background grants (bgSkills) a skill it picked: the class's list less the
// background's skills, the other picks ticked. It is nil otherwise; the
// picks are those ProfSources credits to the class.
func (a *App) backgroundSkillPick(ctx context.Context, sh dnd5e.Sheet, bgSkills []string) (*views.SkillPick, error) {
	if len(sh.Classes) == 0 {
		return nil, nil
	}
	cl := sh.Classes[0]
	var picks, kept []string
	for _, k := range sh.SkillProfs {
		if sh.ProfSources[k] == cl.Name {
			picks = append(picks, k)
			if !slices.Contains(bgSkills, k) {
				kept = append(kept, k)
			}
		}
	}
	if len(kept) == len(picks) {
		return nil, nil
	}
	row, _, err := entryFor(ctx, a.Store.Queries, enum.KindClass, cl.Slug, cl.Source)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	data, err := compendium.DecodeData(enum.KindClass, row.DataJson)
	if err != nil {
		return nil, err
	}
	c := wizardChoice(row, data)
	after := sh
	after.SkillProfs = slices.Concat(slices.DeleteFunc(slices.Clone(sh.SkillProfs), func(k string) bool { return slices.Contains(picks, k) }), bgSkills)
	return classSkillPick(after, c.Skills, len(picks), kept), nil
}

func postedSkills(c *gin.Context) []string {
	return slices.Compact(slices.Sorted(slices.Values(c.PostFormArray("skill"))))
}

// checkSkillPick refuses posted skills that are not p's count of its options.
func checkSkillPick(p *views.SkillPick, skills []string) error {
	need, opts := 0, []string(nil)
	if p != nil {
		need, opts = p.Count, p.Options
	}
	if len(skills) != need || slices.ContainsFunc(skills, func(k string) bool { return !slices.Contains(opts, k) }) {
		return fieldError{userError{errors.New(views.Choose(need, false))}, views.FieldSkill}
	}
	return nil
}

// equipmentChoices are the class's and the background's starting
// equipment options.
func (a *App) equipmentChoices(ctx context.Context, sh dnd5e.Sheet) ([]views.EquipmentChoice, error) {
	var out []views.EquipmentChoice
	add := func(kind enum.CompendiumKind, slug string, text func(any) string) error {
		if slug == "" {
			return nil
		}
		e, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: kind, Slug: slug})
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		if err != nil {
			return err
		}
		data, err := compendium.DecodeData(kind, e.DataJson)
		if err != nil {
			return err
		}
		if opts := compendium.EquipmentOptions(text(data)); len(opts) > 0 {
			out = append(out, views.EquipmentChoice{Key: string(kind), Title: e.Name, Options: opts})
		}
		return nil
	}
	if len(sh.Classes) > 0 {
		if err := add(enum.KindClass, sh.Classes[0].Slug, func(d any) string {
			return compendium.CoreTraits(d.(*compendium.ClassData).Desc)["Starting Equipment"]
		}); err != nil {
			return nil, err
		}
	}
	err := add(enum.KindBackground, compendium.Slugify(sh.Background), func(d any) string {
		for _, p := range d.(*compendium.TextData).Parts {
			if p.Type == compendium.PartEquipment {
				return p.Desc
			}
		}
		return ""
	})
	return out, err
}

// creation is what the choices step asks and how each pick changes the sheet.
type creation struct {
	view   views.CreationChoices
	groups []choiceGroup
	// cantrips are species cantrips that come without a choice (Thaumaturgy).
	cantrips []dnd5e.Spell
	// taught maps a spell slug to the lineage rules ("<group key>=<option>") that grant it.
	taught map[string][]string
	// check refuses class options the finished step does not qualify for
	// (checkNewOptions); nil without a compendium class.
	check func(before, after dnd5e.Sheet, learned []dnd5e.Cantrip) error
}

// choiceGroup is a creation choice and what picking from it does.
type choiceGroup struct {
	views.ChoiceGroup
	apply func(s *dnd5e.Sheet, picked []string)
	// sub are the choices an option brings when it is picked (an Origin
	// feat's spells), applied after it.
	sub map[string][]choiceGroup
}

// creationChoices are the picks the species, the background's feat and the
// level-1 class leave open: a lineage, a skill, the feat's spells or
// proficiencies, two languages, invocations and spells.
func (a *App) creationChoices(ctx context.Context, sh dnd5e.Sheet, campaignID int64) (creation, error) {
	cr := creation{taught: map[string][]string{}}
	featGroups, err := a.backgroundFeatChoices(ctx, sh, campaignID)
	if err != nil {
		return cr, err
	}
	if err := a.speciesChoices(ctx, &cr, sh, campaignID); err != nil {
		return cr, err
	}
	// After the species' groups, so the feat's ability is the last to set its spells'.
	cr.groups = append(cr.groups, featGroups...)
	if len(sh.Classes) > 0 {
		if err := a.classChoices(ctx, &cr, sh, campaignID); err != nil {
			return cr, err
		}
	}
	var langs []compendium.NameDesc
	for _, l := range dnd5e.StandardLanguages {
		if !slices.Contains(sh.Languages, l) {
			langs = append(langs, compendium.NameDesc{Name: l})
		}
	}
	cr.groups = append(cr.groups, choiceGroup{views.ChoiceGroup{Key: "language", Title: "Languages", Count: 2, Options: langs, Known: sh.Languages}, func(s *dnd5e.Sheet, picked []string) {
		s.Languages = append(s.Languages, picked...)
	}, nil})
	cr.view.Groups = groupViews(cr.groups)
	return cr, nil
}

func groupViews(groups []choiceGroup) []views.ChoiceGroup {
	out := make([]views.ChoiceGroup, len(groups))
	for i, g := range groups {
		out[i] = g.ChoiceGroup
		out[i].Count = min(g.Count, len(g.Options)) // as readPicks needs
		for o, sub := range g.sub {
			if out[i].Sub == nil {
				out[i].Sub = map[string][]views.ChoiceGroup{}
			}
			out[i].Sub[o] = groupViews(sub)
		}
	}
	return out
}

// readPicks reads each group's picks from the form: its count (or all of
// fewer options), only of its options, each once. A group offering "No
// change" may be left empty. Only the picked options' own choices are read.
// Every bad group is reported, joined.
func readPicks(c *gin.Context, groups []choiceGroup) (map[string][]string, error) {
	picks := map[string][]string{}
	var errs []error
	for _, g := range groups {
		picked := slices.Compact(slices.Sorted(slices.Values(c.PostFormArray(g.Key))))
		need := min(g.Count, len(g.Options))
		optional := len(picked) == 0 && slices.ContainsFunc(g.Options, func(o compendium.NameDesc) bool { return o.Name == keepSpells })
		if (len(picked) != need && !optional) || slices.ContainsFunc(picked, func(p string) bool {
			return !slices.ContainsFunc(g.Options, func(o compendium.NameDesc) bool { return o.Name == p })
		}) {
			errs = append(errs, pickError{fieldError{userError{errors.New(views.Choose(need, len(g.Known) > 0))}, g.Key}, g.Title})
			continue
		}
		picks[g.Key] = picked
		for _, p := range picked {
			sub, err := readPicks(c, g.sub[p])
			errs = append(errs, err)
			maps.Copy(picks, sub)
		}
	}
	// A class level and a feat (Weapon Master) may not master one kind twice.
	seen := map[string]bool{}
	for _, k := range slices.Sorted(maps.Keys(picks)) {
		if k != views.FieldMastery && !strings.HasSuffix(k, "_"+views.FieldMastery) {
			continue
		}
		for _, p := range picks[k] {
			if seen[p] {
				errs = append(errs, fieldError{userError{fmt.Errorf("%s is picked twice: choose another weapon", p)}, k})
			}
			seen[p] = true
		}
	}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}
	return picks, nil
}

func applyPicks(s *dnd5e.Sheet, groups []choiceGroup, picks map[string][]string) {
	for _, g := range groups {
		if len(picks[g.Key]) > 0 && g.apply != nil {
			g.apply(s, picks[g.Key])
		}
		for _, p := range picks[g.Key] {
			applyPicks(s, g.sub[p], picks)
		}
	}
}

var (
	reSpeedTo      = regexp.MustCompile(`Speed increases to (\d+) feet`)
	reResistance   = regexp.MustCompile(`Resistance to ([A-Z][a-z]+) damage`)
	reSpellAbility = regexp.MustCompile(`Intelligence, Wisdom, or Charisma is your spellcasting ability`)
)

// choiceOptions are a trait's options when it asks to choose one.
func choiceOptions(p compendium.NameDesc) []compendium.NameDesc {
	if !strings.Contains(p.Desc, "Choose") {
		return nil
	}
	return compendium.Options(p.Desc)
}

// applyTraitText applies what a trait's text sets: darkvision, a speed
// increase and damage resistances ("Resistance to Fire damage"; "the
// damage type determined by…" names none).
func applyTraitText(s *dnd5e.Sheet, text string) {
	if strings.Contains(text, "Darkvision") {
		if m := reFeet.FindStringSubmatch(text); m != nil {
			s.Vision = "Darkvision " + m[1] + " ft"
		}
	}
	if m := reSpeedTo.FindStringSubmatch(text); m != nil {
		s.Speed, _ = strconv.Atoi(m[1])
	}
	for _, m := range reResistance.FindAllStringSubmatch(text, -1) {
		dnd5e.AddResistance(s, m[1])
	}
}

// resistsPick reports whether a trait naming p with "Resistance"
// (Dragonborn's Damage Resistance) resists the damage type picked in p.
func resistsPick(parts []compendium.NameDesc, p compendium.NameDesc) bool {
	return slices.ContainsFunc(parts, func(q compendium.NameDesc) bool {
		return q.Name != p.Name && strings.Contains(q.Desc, p.Name) && strings.Contains(q.Desc, "Resistance")
	})
}

// speciesChoices asks the species' "Choose a lineage" traits, skill
// choices and Origin feat, and collects the cantrips its other traits grant.
func (a *App) speciesChoices(ctx context.Context, cr *creation, sh dnd5e.Sheet, campaignID int64) error {
	species := sh.Species
	parts, err := a.speciesParts(ctx, sh)
	if err != nil {
		return err
	}
	var abilityFor string           // the trait whose spells take a picked ability
	spellSlugs := map[string]bool{} // the spells the species can grant
	for _, p := range parts {
		if p.Type != "" {
			continue
		}
		if opts := choiceOptions(p); len(opts) > 0 {
			g, byOption, err := a.lineageChoice(ctx, p, opts, species, resistsPick(parts, p), campaignID)
			if err != nil {
				return err
			}
			cr.groups = append(cr.groups, g)
			for _, name := range slices.Sorted(maps.Keys(byOption)) {
				for _, sp := range byOption[name] {
					spellSlugs[sp.Slug] = true
					cr.taught[sp.Slug] = append(cr.taught[sp.Slug], g.Key+"="+name)
				}
			}
			if reSpellAbility.MatchString(p.Desc) {
				abilityFor = p.Name
			}
			continue
		}
		if strings.Contains(p.Desc, "Origin feat") {
			g, ok, err := a.originFeatChoice(ctx, p, sh, campaignID)
			if err != nil {
				return err
			}
			if ok {
				cr.groups = append(cr.groups, g)
				continue
			}
		}
		if g, ok := skillPick(p); ok {
			cr.groups = append(cr.groups, g)
			continue
		}
		spells, err := a.spellsNamed(ctx, compendium.Cantrips(p.Desc), campaignID)
		if err != nil {
			return err
		}
		cr.cantrips = append(cr.cantrips, spells...)
		for _, sp := range spells {
			spellSlugs[sp.Slug] = true
		}
	}
	if abilityFor != "" && len(spellSlugs) > 0 {
		cr.groups = append(cr.groups, spellAbilityChoice("trait_"+compendium.Slugify(abilityFor)+"_ability", abilityFor, spellSlugs, true))
	}
	return nil
}

// speciesParts are the traits of the sheet's compendium species (its pinned
// copy); none when it has no entry (homebrew typed by hand).
func (a *App) speciesParts(ctx context.Context, sh dnd5e.Sheet) ([]compendium.NameDesc, error) {
	if sh.Species == "" {
		return nil, nil
	}
	td, _, err := a.textEntrySource(ctx, enum.KindSpecies, compendium.Slugify(sh.Species), sh.SpeciesSource)
	if err != nil || td == nil {
		return nil, err
	}
	return td.Parts, nil
}

// lineageSpells are the spells the sheet's lineage pick teaches by character
// level ("Level 3" and "Level 5" columns) whose free cast it lacks; one
// already known from the class list still owes its free cast.
func (a *App) lineageSpells(ctx context.Context, sh dnd5e.Sheet, level int, campaignID int64) ([]dnd5e.Spell, error) {
	parts, err := a.speciesParts(ctx, sh)
	if err != nil {
		return nil, err
	}
	var out []dnd5e.Spell
	for _, p := range parts {
		for _, o := range choiceOptions(p) {
			if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool {
				return f.Source == sh.Species && f.Name == fmt.Sprintf("%s (%s)", p.Name, o.Name)
			}) {
				continue
			}
			byLevel := compendium.LevelSpells(p.Desc)[o.Name]
			var names []string
			for _, l := range slices.Sorted(maps.Keys(byLevel)) {
				if l <= level {
					names = append(names, byLevel[l])
				}
			}
			spells, err := a.spellsNamed(ctx, names, campaignID)
			if err != nil {
				return nil, err
			}
			ability := cmp.Or(sh.SpeciesAbility, lineageAbility(sh, o.Desc))
			for _, sp := range spells {
				if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == dnd5e.FreeCast(sp, "").Name }) {
					sp.Prepared, sp.Ability = true, ability
					out = append(out, sp)
				}
			}
		}
	}
	return out, nil
}

// lineageAbility is the ability picked for a lineage option's spells, read
// from its level-1 cantrip, for sheets made before SpeciesAbility.
func lineageAbility(sh dnd5e.Sheet, optionDesc string) string {
	for _, name := range compendium.Cantrips(optionDesc) {
		if i := slices.IndexFunc(sh.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == compendium.Slugify(name) }); i >= 0 {
			return sh.Spells[i].Ability
		}
	}
	return ""
}

// spellAbilityChoice picks Intelligence, Wisdom or Charisma for the spells
// a species (species) or a feat grants; slugs may fill up until it runs,
// after the spells are added.
func spellAbilityChoice(key, from string, slugs map[string]bool, species bool) choiceGroup {
	var opts []compendium.NameDesc
	for _, ab := range []string{"int", "wis", "cha"} {
		opts = append(opts, compendium.NameDesc{Name: dnd5e.AbilityNames[ab]})
	}
	g := views.ChoiceGroup{Key: key, Title: "Spellcasting ability for " + from, Count: 1, Options: opts}
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
		for ab, name := range dnd5e.AbilityNames {
			if name != picked[0] {
				continue
			}
			if species {
				s.SpeciesAbility = ab
			}
			for i, sp := range s.Spells {
				if slugs[sp.Slug] {
					s.Spells[i].Ability = ab
				}
			}
		}
	}, nil}
}

// textEntry is the compendium species, background or feat with this slug;
// nil when it has none.
func (a *App) textEntry(ctx context.Context, kind enum.CompendiumKind, slug string) (*compendium.TextData, error) {
	td, _, err := a.textEntrySource(ctx, kind, slug, "")
	return td, err
}

// textEntrySource is textEntry for a ref pinned to pin (see entryFor), with
// the entry's source key, for pinning.
func (a *App) textEntrySource(ctx context.Context, kind enum.CompendiumKind, slug, pin string) (*compendium.TextData, string, error) {
	row, _, err := entryFor(ctx, a.Store.Queries, kind, slug, pin)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, "", nil
	}
	if err != nil {
		return nil, "", err
	}
	d, err := compendium.Decode[compendium.TextData](kind, row.DataJson)
	if err != nil {
		return nil, "", err
	}
	return d, row.SourceKey, nil
}

var reOptionCantrip = regexp.MustCompile(`(?i)choose one of your known (\w+) cantrips that (deals damage|requires an attack roll)(?: and has a range of (\d+)\+ feet)?`)

// toolNames are the compendium's tools, without their cost ("Smith's Tools").
// Tools a proficiency applies to say which ability they use ("Ability:
// Strength. Utilize: …"); kits filed with them (Climber's Kit) do not.
func (a *App) toolNames(ctx context.Context, campaignID int64) ([]string, error) {
	rows, err := a.Store.ListItemsByCategory(ctx, sqlc.ListItemsByCategoryParams{Category: "tools", CampaignID: campaignID})
	if err != nil {
		return nil, err
	}
	var out []string
	for _, r := range rows {
		if strings.HasPrefix(r.Description, "Ability:") {
			out = append(out, reToolCost.ReplaceAllString(r.Name, ""))
		}
	}
	return out, nil
}

// lineageChoice picks one of a trait's options: the feature is named after
// it, starts with its text, and its cantrips, speed, darkvision and
// resistance apply (resists: the option names a damage type to resist). It
// also returns the spells each option can grant, by option name.
func (a *App) lineageChoice(ctx context.Context, p compendium.NameDesc, opts []compendium.NameDesc, species string, resists bool, campaignID int64) (choiceGroup, map[string][]dnd5e.Spell, error) {
	cantrips := map[string][]dnd5e.Spell{}
	for _, o := range opts {
		spells, err := a.spellsNamed(ctx, compendium.Cantrips(o.Desc), campaignID)
		if err != nil {
			return choiceGroup{}, nil, err
		}
		cantrips[o.Name] = spells
	}
	g := views.ChoiceGroup{Key: "trait_" + compendium.Slugify(p.Name), Title: p.Name, Count: 1, Options: opts}
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
		o := opts[slices.IndexFunc(opts, func(o compendium.NameDesc) bool { return o.Name == picked[0] })]
		if i := slices.IndexFunc(s.Features, func(f dnd5e.Feature) bool { return f.Name == p.Name && f.Source == species }); i >= 0 {
			name := fmt.Sprintf("%s (%s)", p.Name, o.Name)
			for j, m := range s.Modifiers {
				if m.Source == p.Name {
					s.Modifiers[j].Source, s.Modifiers[j].Note = name, name
				}
			}
			s.Features[i].Name, s.Features[i].Description = name, "**"+o.Name+".** "+o.Desc+"\n\n"+s.Features[i].Description
		}
		for _, sp := range cantrips[o.Name] {
			s.AddSpell(sp)
		}
		applyTraitText(s, o.Desc)
		if resists {
			dnd5e.AddResistance(s, o.Desc)
		}
	}, nil}, cantrips, nil
}

// skillPick reads "proficiency in the Insight, Perception, or Survival
// skill" or "in one skill of your choice" as a choice of one skill.
func skillPick(p compendium.NameDesc) (choiceGroup, bool) {
	if !strings.Contains(p.Desc, "proficiency in") || !strings.Contains(p.Desc, "skill") ||
		len(dnd5e.ParseEffects(p.Name, p.Desc, nil, "").Skills) > 0 { // a fixed grant ("the Insight and Medicine skills") is no choice
		return choiceGroup{}, false
	}
	var opts []compendium.NameDesc
	for _, sk := range dnd5e.Skills {
		if strings.Contains(p.Desc, "of your choice") || strings.Contains(p.Desc, sk.Name) {
			opts = append(opts, compendium.NameDesc{Name: sk.Name})
		}
	}
	if len(opts) < 2 {
		return choiceGroup{}, false
	}
	g := views.ChoiceGroup{Key: "trait_" + compendium.Slugify(p.Name), Title: p.Name, Count: 1, Options: opts}
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) { s.AddProfs(p.Name, dnd5e.SkillKeys(picked)) }, nil}, true
}

// classChoices offers the class's spells and its "… Options" features
// (Eldritch Invocation Options) as many as its table column gives.
func (a *App) classChoices(ctx context.Context, cr *creation, sh dnd5e.Sheet, campaignID int64) error {
	cl := sh.Classes[0]
	info, _, err := classEntry(ctx, a.Store.Queries, cl.Slug, cl.Source)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}
	cr.view.Columns = info.ColumnsAt(cl.Level)
	slots, _, pactLevel := dnd5e.SlotMaxima(sh.Classes)
	if cr.view.Spells, err = a.classSpellOptions(ctx, cl.Slug, slots, pactLevel, sh, cr.cantrips, campaignID); err != nil {
		return err
	}
	for i, sp := range cr.view.Spells {
		cr.view.Spells[i].TaughtBy = cr.taught[sp.Slug]
	}
	cr.view.SpellNeed = spellNeed(info, cl.Level, sh, cr.view.Spells)
	cantrips, err := a.optionCantrips(ctx, sh, info, cl.Level, campaignID)
	if err != nil {
		return err
	}
	mastery, err := a.masteryGroups(ctx, sh, info, cl.Level, campaignID)
	if err != nil {
		return err
	}
	cr.groups = slices.Concat(cr.groups, classOptionGroups(info, cl.Level, 0, sh, cantrips), mastery)
	cr.check = func(before, after dnd5e.Sheet, learned []dnd5e.Cantrip) error {
		return checkNewOptions(info, before, after, cl.Level, learned)
	}
	return nil
}

// cantripData are the compendium entries of these spells' class cantrips
// (not a species' or a feat's), for "a Warlock Cantrip That Deals Damage". A
// cantrip learned from one class's list (Spell.Class) counts for that class
// only; older sheets without it count for every list it is on.
func (a *App) cantripData(ctx context.Context, spells []dnd5e.Spell, campaignID int64) ([]dnd5e.Cantrip, error) {
	var out []dnd5e.Cantrip
	for _, sp := range spells {
		if sp.Level != 0 || sp.Feat != "" || sp.Ability != "" {
			continue
		}
		row, err := a.entryBySlug(ctx, enum.KindSpell, sp.Slug, campaignID)
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		if err != nil {
			return nil, err
		}
		sd, err := compendium.Decode[compendium.SpellData](enum.KindSpell, row.DataJson)
		if err != nil {
			return nil, err
		}
		c := dnd5e.Cantrip{Name: sp.Name, Range: sd.Range, Desc: sd.Desc, Classes: sd.Classes}
		if sp.Class != "" {
			c.Classes = []string{sp.Class}
		}
		out = append(out, c)
	}
	return out, nil
}

// optionCantrips are the cantrips reaching classLevel in info may count for
// its options: the sheet's, and the class list's when the Cantrips column
// grows (one may be learned in the same step, creation included).
func (a *App) optionCantrips(ctx context.Context, sh dnd5e.Sheet, info dnd5e.ClassInfo, classLevel int, campaignID int64) ([]dnd5e.Cantrip, error) {
	could := sh
	could.Spells = slices.Clone(sh.Spells)
	now, _ := strconv.Atoi(info.ColumnsAt(classLevel)["Cantrips"])
	was, _ := strconv.Atoi(info.ColumnsAt(classLevel - 1)["Cantrips"])
	if now > was {
		rows, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: info.Slug, MaxLevel: 0, CampaignID: campaignID})
		if err != nil {
			return nil, err
		}
		for _, r := range rows {
			could.AddSpell(dnd5e.Spell{Slug: r.Slug, Name: r.Name, Class: info.Slug})
		}
	}
	return a.cantripData(ctx, could.Spells, campaignID)
}

// classOptionGroups offer a class's "… Options" features (Eldritch
// Invocation Options) as many times as the matching table column grows
// from class level before to level. Options whose prerequisite the sheet
// does not meet, or already taken and not repeatable, are left out; options
// picked together count for each other (checkNewOptions checks the result).
// An option that applies to one of your cantrips asks which. On a level-up
// (before > 0) whose feature text allows it, one taken option may be
// replaced, unless another taken option needs it.
func classOptionGroups(info dnd5e.ClassInfo, level, before int, sh dnd5e.Sheet, cantrips []dnd5e.Cantrip) []choiceGroup {
	now, was := info.ColumnsAt(level), info.ColumnsAt(before)
	var out []choiceGroup
	for _, f := range info.Features {
		kind, ok := strings.CutSuffix(f.Name, " Options")
		if !ok {
			continue
		}
		o := optionKind{info: info, kind: kind, level: level, all: compendium.Options(f.Description), cantrips: cantrips}
		opts := o.eligible(sh)
		for _, col := range slices.Sorted(maps.Keys(now)) {
			n, _ := strconv.Atoi(now[col])
			m, _ := strconv.Atoi(was[col])
			if n <= m || len(opts) == 0 || !strings.HasPrefix(f.Name, strings.TrimSuffix(col, "s")) {
				continue
			}
			out = append(out, o.group("class_"+compendium.Slugify(kind), col, n-m, opts, sh, nil))
		}
		if g, ok := o.swap(before, sh); ok {
			out = append(out, g)
		}
	}
	return out
}

// optionKind is one "… Options" feature of a class (Eldritch Invocation) at
// the class level being reached.
type optionKind struct {
	info     dnd5e.ClassInfo
	kind     string
	level    int
	all      []compendium.NameDesc
	cantrips []dnd5e.Cantrip
}

// eligible are the options sh may take: prerequisite met, not taken unless
// Repeatable, a cantrip left to choose if they ask for one. Options eligible
// together count for each other's prerequisites ("Pact of the Blade
// Invocation"), until no more open up.
func (k optionKind) eligible(sh dnd5e.Sheet) []compendium.NameDesc {
	var opts []compendium.NameDesc
	for grown := true; grown; {
		with := sh
		with.Features = slices.Clone(sh.Features)
		for _, o := range opts {
			with.Features = append(with.Features, dnd5e.Feature{Name: k.kind + ": " + o.Name})
		}
		var next []compendium.NameDesc
		for _, o := range k.all {
			taken := slices.ContainsFunc(sh.Features, func(h dnd5e.Feature) bool { return h.Name == k.kind+": "+o.Name })
			cs, asks := k.cantripChoices(sh, o)
			if dnd5e.MeetsPrerequisite(compendium.OptionPrerequisite(o.Desc), with, k.level, k.cantrips, compendium.Slugify) && (!taken || isRepeatable(o)) && (!asks || len(cs) > 0) {
				next = append(next, o)
			}
		}
		grown, opts = len(next) > len(opts), next
	}
	return opts
}

// cantripChoices are the cantrips an option's "Choose one of your known
// Warlock cantrips that deals damage" may name, less those sh took it for
// already; asks reports whether its text asks.
func (k optionKind) cantripChoices(sh dnd5e.Sheet, o compendium.NameDesc) (out []compendium.NameDesc, asks bool) {
	m := reOptionCantrip.FindStringSubmatch(o.Desc)
	if m == nil {
		return nil, false
	}
	minRange, _ := strconv.Atoi(m[3])
	for _, c := range k.cantrips {
		if dnd5e.CantripFits(c, compendium.Slugify(m[1]), strings.EqualFold(m[2], "requires an attack roll"), minRange) &&
			!slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == k.name(o, c.Name) }) {
			out = append(out, compendium.NameDesc{Name: c.Name})
		}
	}
	return out, true
}

// name is the feature an option becomes: "Eldritch Invocation: Agonizing
// Blast (Eldritch Blast)" with a cantrip.
func (k optionKind) name(o compendium.NameDesc, cantrip string) string {
	if cantrip != "" {
		return fmt.Sprintf("%s: %s (%s)", k.kind, o.Name, cantrip)
	}
	return k.kind + ": " + o.Name
}

// add puts option o on s as a feature, for cantrip if it asks for one.
func (k optionKind) add(s *dnd5e.Sheet, o compendium.NameDesc, cantrip string) {
	name := k.name(o, cantrip)
	// A Repeatable option taken again is a new feature, not a replacement.
	for n := 2; slices.ContainsFunc(s.Features, func(h dnd5e.Feature) bool { return strings.EqualFold(h.Name, name) }); n++ {
		name = fmt.Sprintf("%s (%d)", k.name(o, cantrip), n)
	}
	feat := dnd5e.Feature{Name: name, Source: k.info.Name, Level: k.level, Description: o.Desc}
	mods := dnd5e.ParseEffects(feat.Name, feat.Description, k.info.Table, k.info.Name).Apply(&feat)
	s.PutFeature(feat, mods)
}

// group asks for count of opts. An option that asks for a cantrip is added
// by its own "Agonizing Blast: cantrip" choice below it. first runs before
// each pick and may skip it (a swap removes the replaced option).
func (k optionKind) group(key, title string, count int, opts []compendium.NameDesc, sh dnd5e.Sheet, first func(*dnd5e.Sheet, compendium.NameDesc) bool) choiceGroup {
	sub := map[string][]choiceGroup{}
	for _, o := range opts {
		if cs, asks := k.cantripChoices(sh, o); asks {
			g := views.ChoiceGroup{Key: key + "_" + compendium.Slugify(o.Name), Title: o.Name + ": cantrip", Count: 1, Options: cs}
			sub[o.Name] = []choiceGroup{{g, func(s *dnd5e.Sheet, picked []string) { k.add(s, o, picked[0]) }, nil}}
		}
	}
	return choiceGroup{views.ChoiceGroup{Key: key, Title: title, Count: count, Options: opts}, func(s *dnd5e.Sheet, picked []string) {
		for _, name := range picked {
			o := opts[slices.IndexFunc(opts, func(o compendium.NameDesc) bool { return o.Name == name })]
			if first != nil && !first(s, o) {
				continue
			}
			if _, ok := sub[o.Name]; !ok {
				k.add(s, o, "")
			}
		}
	}, sub}
}

// swap is "Replace one Eldritch Invocation": No change, or a taken option no
// other taken option needs, then one it may be replaced with.
func (k optionKind) swap(before int, sh dnd5e.Sheet) (choiceGroup, bool) {
	kind := k.kind
	if before == 0 || !slices.ContainsFunc(k.info.Features, func(f dnd5e.ClassFeature) bool {
		return f.Name == kind+"s" && strings.Contains(f.Description, "replace one of your")
	}) {
		return choiceGroup{}, false
	}
	key := "replace_" + compendium.Slugify(kind)
	opts := []compendium.NameDesc{{Name: keepSpells}}
	sub := map[string][]choiceGroup{}
	for _, old := range sh.Features {
		if !strings.HasPrefix(old.Name, kind+": ") {
			continue
		}
		without := sh
		without.Features = slices.DeleteFunc(slices.Clone(sh.Features), func(f dnd5e.Feature) bool { return f.Name == old.Name })
		if slices.ContainsFunc(without.Features, func(f dnd5e.Feature) bool {
			p := compendium.OptionPrerequisite(f.Description)
			// Only one that holds now and fails without old needs it.
			return strings.HasPrefix(f.Name, kind+": ") && dnd5e.MeetsPrerequisite(p, sh, k.level, k.cantrips, compendium.Slugify) &&
				!dnd5e.MeetsPrerequisite(p, without, k.level, k.cantrips, compendium.Slugify)
		}) {
			continue // another option needs it
		}
		with := slices.DeleteFunc(k.eligible(without), func(o compendium.NameDesc) bool {
			return old.Name == kind+": "+o.Name || strings.HasPrefix(old.Name, kind+": "+o.Name+" (") // "… (2)", "… (Eldritch Blast)"
		})
		if len(with) == 0 {
			continue
		}
		label := strings.TrimPrefix(old.Name, kind+": ")
		opts = append(opts, compendium.NameDesc{Name: label})
		sub[label] = []choiceGroup{k.group(key+"_"+compendium.Slugify(label), "Replace "+label+" with", 1, with, without, func(s *dnd5e.Sheet, o compendium.NameDesc) bool {
			// Taken again by the same level-up and not Repeatable: keep old.
			if !isRepeatable(o) && slices.ContainsFunc(s.Features, func(f dnd5e.Feature) bool { return f.Name == kind+": "+o.Name }) {
				return false
			}
			s.Features = slices.DeleteFunc(s.Features, func(f dnd5e.Feature) bool { return f.Name == old.Name })
			s.Modifiers = slices.DeleteFunc(s.Modifiers, func(m dnd5e.CustomMod) bool { return m.Source == old.Name })
			return true
		})}
	}
	if len(sub) == 0 {
		return choiceGroup{}, false
	}
	return choiceGroup{views.ChoiceGroup{Key: key, Title: "Replace one " + kind, Count: 1, Options: opts}, nil, sub}, true
}

// checkNewOptions refuses a level-up whose new class options (after, not in
// before) miss their prerequisite on the result, or name a cantrip it does
// not know: picks made together were offered for each other.
func checkNewOptions(info dnd5e.ClassInfo, before, after dnd5e.Sheet, level int, cantrips []dnd5e.Cantrip) error {
	for _, f := range after.Features {
		kind, name, _ := strings.Cut(f.Name, ": ")
		if f.Source != info.Name || slices.ContainsFunc(before.Features, func(b dnd5e.Feature) bool { return b.Name == f.Name }) ||
			!slices.ContainsFunc(info.Features, func(cf dnd5e.ClassFeature) bool { return cf.Name == kind+" Options" }) {
			continue
		}
		base, cantrip, _ := strings.Cut(name, " (")
		if p := compendium.OptionPrerequisite(f.Description); !dnd5e.MeetsPrerequisite(p, after, level, cantrips, compendium.Slugify) {
			return userError{fmt.Errorf("%s needs %s", base, p)}
		}
		k := optionKind{kind: kind, cantrips: cantrips}
		o := compendium.NameDesc{Name: base, Desc: f.Description}
		if cs, asks := k.cantripChoices(before, o); asks {
			cantrip = strings.TrimSuffix(cantrip, ")")
			if !slices.ContainsFunc(cs, func(c compendium.NameDesc) bool { return c.Name == cantrip }) {
				return fieldError{userError{fmt.Errorf("%s needs %s among your cantrips: tick it under spells", o.Name, cantrip)}, views.FieldSpell}
			}
		}
	}
	return nil
}

// isRepeatable reports whether a class option may be taken more than once.
func isRepeatable(o compendium.NameDesc) bool { return strings.Contains(o.Desc, "**Repeatable.**") }

// spellsNamed are the compendium spells of these names; ones it lacks are
// left out.
func (a *App) spellsNamed(ctx context.Context, names []string, campaignID int64) ([]dnd5e.Spell, error) {
	var out []dnd5e.Spell
	for _, n := range names {
		sp, err := a.compendiumSpell(ctx, compendium.Slugify(n), campaignID)
		var ue userError
		if errors.As(err, &ue) {
			continue
		}
		if err != nil {
			return nil, err
		}
		out = append(out, sp)
	}
	return out, nil
}
