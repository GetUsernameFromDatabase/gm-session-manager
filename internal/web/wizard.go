package web

import (
	"cmp"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// wizardSteps are the creation wizard's steps, in order. Class, species,
// background, choices and equipment add to the sheet, so each runs once; afterwards
// the sheet is changed directly.
var wizardSteps = []views.WizardStep{
	{Key: "class", Title: "Class"},
	{Key: "species", Title: "Species"},
	{Key: "abilities", Title: "Ability scores"},
	{Key: "background", Title: "Background"},
	{Key: "choices", Title: "Choices"},
	{Key: "equipment", Title: "Equipment"},
	{Key: "details", Title: "Details"},
	{Key: "review", Title: "Review"},
}

func wizardPath(id int64, step string) string {
	return fmt.Sprintf("/characters/%d/create/%s", id, step)
}

// nextStep is the step after key, or the sheet after the last.
func nextStep(id int64, key string) string {
	i := slices.IndexFunc(wizardSteps, func(s views.WizardStep) bool { return s.Key == key })
	if i < 0 || i+1 >= len(wizardSteps) {
		return fmt.Sprintf("/characters/%d", id)
	}
	return wizardPath(id, wizardSteps[i+1].Key)
}

// wizardAccess loads the character for its player, with the step asked for.
func (a *App) wizardAccess(c *gin.Context) (who, sqlc.Character, string, bool) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return w, ch, "", false
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player sets up their character.")
		return w, ch, "", false
	}
	step := c.Param("step")
	if !slices.ContainsFunc(wizardSteps, func(s views.WizardStep) bool { return s.Key == step }) {
		step = wizardSteps[0].Key
	}
	return w, ch, step, true
}

func (a *App) wizardStart(c *gin.Context) {
	if _, ch, _, ok := a.wizardAccess(c); ok {
		c.Redirect(http.StatusSeeOther, wizardPath(ch.ID, wizardSteps[0].Key))
	}
}

func (a *App) wizardPage(c *gin.Context) {
	w, ch, step, ok := a.wizardAccess(c)
	if !ok {
		return
	}
	v, err := a.wizardView(c.Request.Context(), w, ch, step, c.Query("pick"))
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.renderWizard(c, http.StatusOK, w, ch, v)
}

func (a *App) renderWizard(c *gin.Context, status int, w who, ch sqlc.Character, v views.WizardView) {
	camp, err := a.Store.GetCampaign(c.Request.Context(), ch.CampaignID)
	if !a.found(c, err) {
		return
	}
	a.render(c, status, views.WizardPage(a.layout(w, "Create "+ch.Name, views.NavCharacter, &camp, false), v))
}

// wizardView builds a step's page; pick is the class, species or
// background being looked at.
func (a *App) wizardView(ctx context.Context, w who, ch sqlc.Character, step, pick string) (views.WizardView, error) {
	s, err := a.sheetView(ctx, w.viewer, ch)
	if err != nil {
		return views.WizardView{}, err
	}
	s.CanSheet, s.CanHP, s.CanConditions = false, false, false // the review is read-only
	v := views.WizardView{CharacterID: ch.ID, Step: step, Steps: wizardSteps, Sheet: s, Picked: pick}
	kind := map[string]compendium.Kind{"class": compendium.KindClass, "species": compendium.KindSpecies, "background": compendium.KindBackground}[step]
	if kind == "" {
		switch step {
		case "equipment":
			v.Equipment, err = a.equipmentChoices(ctx, s.C.Sheet)
		case "choices":
			var cr creation
			cr, err = a.creationChoices(ctx, s.C.Sheet, ch.CampaignID)
			v.Choices = &cr.view
		}
		return v, err
	}
	names, err := a.Store.ListCompendiumNames(ctx, string(kind))
	if err != nil {
		return v, err
	}
	for _, n := range names {
		v.Options = append(v.Options, views.Option{Slug: n.Slug, Name: n.Name})
	}
	if pick == "" {
		return v, nil
	}
	e, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(kind), Slug: pick})
	if errors.Is(err, sql.ErrNoRows) {
		return v, nil
	}
	if err != nil {
		return v, err
	}
	data, err := compendium.DecodeData(kind, e.DataJson)
	if err != nil {
		return v, err
	}
	v.Choice = wizardChoice(e, data)
	return v, nil
}

// wizardChoice is what a class, species or background gives, for the page
// to show and the step to apply.
func wizardChoice(e sqlc.CompendiumEntry, data any) *views.WizardChoice {
	ch := &views.WizardChoice{Slug: e.Slug, Name: e.Name}
	switch d := data.(type) {
	case *compendium.ClassData:
		ch.Traits = compendium.CoreTraits(d.Desc)
		ch.SkillCount, ch.Skills = skillChoice(ch.Traits["Skill Proficiencies"])
		for _, f := range d.Features {
			if slices.Contains(f.Levels, 1) && !strings.Contains(strings.ToLower(f.Name), "core") {
				ch.Features = append(ch.Features, views.FeaturePreview{Name: f.Name, Desc: f.Desc, Effects: dnd5e.ParseEffects(f.Name, f.Desc, d.Table, e.Name).Summary()})
			}
		}
	case *compendium.TextData:
		for _, p := range d.Parts {
			switch p.Type {
			case "ability_score":
				ch.Abilities = abilityKeys(p.Desc)
			case "skill_proficiency":
				ch.Skills = skillKeys(compendium.SplitList(p.Desc))
				ch.SkillCount = len(ch.Skills)
			case "":
				// Untyped parts are traits; typed ones (size, speed, tools,
				// equipment, the feat) are applied or shown as facts.
				ch.Features = append(ch.Features, views.FeaturePreview{Name: p.Name, Desc: p.Desc, Effects: dnd5e.ParseEffects(p.Name, p.Desc, nil, "").Summary()})
			}
			if p.Type != "" || p.Name == "Speed" || p.Name == "Size" {
				ch.Facts = append(ch.Facts, p)
			}
		}
	}
	return ch
}

// skillChoice reads "Choose 2: …" as how many skills and from which keys;
// "any" skills offers them all.
func skillChoice(s string) (int, []string) {
	n, names := compendium.SkillChoice(s)
	if n > 0 && names == nil {
		return n, skillKeys(nil)
	}
	return n, skillKeys(names)
}

// skillKeys matches skill names (spaces from the source's PDF text and all)
// to skill keys; no names gives every skill.
func skillKeys(names []string) []string {
	var keys []string
	for _, sk := range dnd5e.Skills {
		if names == nil || slices.ContainsFunc(names, func(n string) bool { return compendium.Squash(n) == compendium.Squash(sk.Name) }) {
			keys = append(keys, sk.Key)
		}
	}
	return keys
}

func abilityKeys(list string) []string {
	var keys []string
	for _, n := range compendium.SplitList(list) {
		for _, a := range dnd5e.Abilities {
			if strings.EqualFold(dnd5e.AbilityNames[a], n) {
				keys = append(keys, a)
			}
		}
	}
	return keys
}

// equipmentChoices are the class's and the background's starting
// equipment options.
func (a *App) equipmentChoices(ctx context.Context, sh dnd5e.Sheet) ([]views.EquipmentChoice, error) {
	var out []views.EquipmentChoice
	add := func(kind compendium.Kind, slug string, text func(any) string) error {
		if slug == "" {
			return nil
		}
		e, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(kind), Slug: slug})
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
		if err := add(compendium.KindClass, sh.Classes[0].Slug, func(d any) string {
			return compendium.CoreTraits(d.(*compendium.ClassData).Desc)["Starting Equipment"]
		}); err != nil {
			return nil, err
		}
	}
	err := add(compendium.KindBackground, compendium.Slugify(sh.Background), func(d any) string {
		for _, p := range d.(*compendium.TextData).Parts {
			if p.Type == "equipment" {
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
	var cr creation
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
		if err := a.classChoices(ctx, &cr, sh); err != nil {
			return cr, err
		}
	}
	var langs []compendium.NameDesc
	for _, l := range dnd5e.StandardLanguages {
		if !slices.Contains(sh.Languages, l) {
			langs = append(langs, compendium.NameDesc{Name: l})
		}
	}
	cr.groups = append(cr.groups, choiceGroup{views.ChoiceGroup{Key: "language", Title: "Languages", Count: 2, Options: langs}, func(s *dnd5e.Sheet, picked []string) {
		s.Languages = append(s.Languages, picked...)
	}, nil})
	cr.view.Groups = groupViews(cr.groups)
	return cr, nil
}

func groupViews(groups []choiceGroup) []views.ChoiceGroup {
	out := make([]views.ChoiceGroup, len(groups))
	for i, g := range groups {
		out[i] = g.ChoiceGroup
		for o, sub := range g.sub {
			if out[i].Sub == nil {
				out[i].Sub = map[string][]views.ChoiceGroup{}
			}
			out[i].Sub[o] = groupViews(sub)
		}
	}
	return out
}

// readPicks reads each group's picks from the form: up to its count, only
// of its options, each once. Only the picked options' own choices are read.
func readPicks(c *gin.Context, groups []choiceGroup) (map[string][]string, error) {
	picks := map[string][]string{}
	for _, g := range groups {
		picked := slices.Compact(slices.Sorted(slices.Values(c.PostFormArray(g.Key))))
		if len(picked) > g.Count || slices.ContainsFunc(picked, func(p string) bool {
			return !slices.ContainsFunc(g.Options, func(o compendium.NameDesc) bool { return o.Name == p })
		}) {
			return nil, userError{fmt.Errorf("choose up to %d of the listed %s", g.Count, strings.ToLower(g.Title))}
		}
		picks[g.Key] = picked
		for _, p := range picked {
			sub, err := readPicks(c, g.sub[p])
			if err != nil {
				return nil, err
			}
			maps.Copy(picks, sub)
		}
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
		addResistance(s, m[1])
	}
}

func addResistance(s *dnd5e.Sheet, damage string) {
	if d := strings.ToLower(damage); slices.Contains(dnd5e.DamageTypes, d) && !slices.Contains(s.Resistances, d) {
		s.Resistances = append(s.Resistances, d)
	}
}

// speciesChoices asks the species' "Choose a lineage" traits, skill
// choices and Origin feat, and collects the cantrips its other traits grant.
func (a *App) speciesChoices(ctx context.Context, cr *creation, sh dnd5e.Sheet, campaignID int64) error {
	species := sh.Species
	parts, err := a.speciesParts(ctx, species)
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
			// A trait naming this one with "Resistance" (Dragonborn's Damage
			// Resistance) resists the damage type picked here.
			resists := slices.ContainsFunc(parts, func(q compendium.NameDesc) bool {
				return q.Name != p.Name && strings.Contains(q.Desc, p.Name) && strings.Contains(q.Desc, "Resistance")
			})
			g, spells, err := a.lineageChoice(ctx, p, opts, species, resists, campaignID)
			if err != nil {
				return err
			}
			cr.groups = append(cr.groups, g)
			for _, sp := range spells {
				spellSlugs[sp.Slug] = true
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

// speciesParts are the traits of the compendium species of this name; none
// when it has no entry (homebrew typed by hand).
func (a *App) speciesParts(ctx context.Context, species string) ([]compendium.NameDesc, error) {
	if species == "" {
		return nil, nil
	}
	td, err := a.textEntry(ctx, compendium.KindSpecies, compendium.Slugify(species))
	if err != nil || td == nil {
		return nil, err
	}
	return td.Parts, nil
}

// lineageSpells are the spells the sheet's lineage pick teaches by character
// level ("Level 3" and "Level 5" columns) whose free cast it lacks; one
// already known from the class list still owes its free cast.
func (a *App) lineageSpells(ctx context.Context, sh dnd5e.Sheet, level int, campaignID int64) ([]dnd5e.Spell, error) {
	parts, err := a.speciesParts(ctx, sh.Species)
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
				if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return f.Name == freeCast(sp, "").Name }) {
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

// freeCast is a lineage spell's one cast without a slot per Long Rest, as a
// feature so the Actions card counts it. Like other species features it has
// no level (Feature.Level is a class level).
func freeCast(sp dnd5e.Spell, species string) dnd5e.Feature {
	return dnd5e.Feature{
		Name: sp.Name + " (free cast)", Source: species,
		Description: "Cast *" + sp.Name + "* once without a spell slot; you regain this when you finish a Long Rest.",
		Uses:        dnd5e.Uses{Fixed: 1, Reset: "long"},
	}
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

// originFeatChoice offers the compendium's Origin feats (Human Versatile),
// each with its own choices.
func (a *App) originFeatChoice(ctx context.Context, p compendium.NameDesc, sh dnd5e.Sheet, campaignID int64) (choiceGroup, bool, error) {
	names, err := a.Store.ListCompendiumNames(ctx, string(compendium.KindFeat))
	if err != nil {
		return choiceGroup{}, false, err
	}
	key := "trait_" + compendium.Slugify(p.Name)
	var opts []compendium.NameDesc
	sub := map[string][]choiceGroup{}
	for _, n := range names {
		td, err := a.textEntry(ctx, compendium.KindFeat, n.Slug)
		if err != nil {
			return choiceGroup{}, false, err
		}
		if td == nil || td.Category != "Origin" || !meetsPrerequisite(td.Prerequisite, sh, sh.Level()) {
			continue
		}
		opts = append(opts, compendium.NameDesc{Name: n.Name, Desc: textMarkdown(td)})
		groups, err := a.featChoices(ctx, key+"_"+n.Slug, n.Name, td, "", sh, campaignID)
		if err != nil {
			return choiceGroup{}, false, err
		}
		if len(groups) > 0 {
			sub[n.Name] = groups
		}
	}
	if len(opts) == 0 {
		return choiceGroup{}, false, nil
	}
	g := views.ChoiceGroup{Key: key, Title: p.Name + ": an Origin feat", Count: 1, Options: opts}
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
		o := opts[slices.IndexFunc(opts, func(o compendium.NameDesc) bool { return o.Name == picked[0] })]
		addFeat(s, dnd5e.Feature{Name: o.Name, Source: "Feat", Description: o.Desc})
	}, sub}, true, nil
}

// textEntry is the compendium species, background or feat with this slug;
// nil when it has none.
func (a *App) textEntry(ctx context.Context, kind compendium.Kind, slug string) (*compendium.TextData, error) {
	row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(kind), Slug: slug})
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	d, err := compendium.DecodeData(kind, row.DataJson)
	if err != nil {
		return nil, err
	}
	return d.(*compendium.TextData), nil
}

// backgroundFeatChoices are the choices inside the feat the background
// gave ("Magic Initiate (Wizard)").
func (a *App) backgroundFeatChoices(ctx context.Context, sh dnd5e.Sheet, campaignID int64) ([]choiceGroup, error) {
	if sh.Background == "" {
		return nil, nil
	}
	bg, err := a.textEntry(ctx, compendium.KindBackground, compendium.Slugify(sh.Background))
	if err != nil || bg == nil {
		return nil, err
	}
	for _, p := range bg.Parts {
		full := strings.TrimSpace(p.Desc)
		if p.Type != "feat" || !slices.Contains(sh.Feats, full) {
			continue
		}
		name, list := featList(full)
		td, err := a.textEntry(ctx, compendium.KindFeat, compendium.Slugify(name))
		if err != nil || td == nil {
			return nil, err
		}
		return a.featChoices(ctx, "feat_"+compendium.Slugify(name), full, td, list, sh, campaignID)
	}
	return nil, nil
}

// featList splits "Magic Initiate (Wizard)" into the feat and its spell list.
func featList(full string) (string, string) {
	name, list, _ := strings.Cut(full, " (")
	return name, strings.TrimSuffix(list, ")")
}

var (
	rePrereqLevel   = regexp.MustCompile(`^Level (\d+)\+$`)
	rePrereqScore   = regexp.MustCompile(`^(.+) (\d+)\+$`)
	rePrereqFeature = regexp.MustCompile(`^(.+) Feature$`)
)

// meetsPrerequisite reports whether a sheet at this character level meets a
// feat's prerequisite ("Level 4+, Strength or Dexterity 13+", "Fighting Style
// Feature"); every clause must hold.
func meetsPrerequisite(prereq string, sh dnd5e.Sheet, level int) bool {
	for _, clause := range strings.FieldsFunc(prereq, func(r rune) bool { return r == ',' || r == ';' }) {
		clause = strings.TrimSpace(clause)
		var scores []string // "Strength or Dexterity" of "… 13+"
		ms := rePrereqScore.FindStringSubmatch(clause)
		if ms != nil {
			scores = abilityKeys(strings.ReplaceAll(ms[1], " or ", ", "))
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
			if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool { return slices.Contains(names, f.Name) }) {
				return false
			}
		}
		// shortcut: other clauses ("Proficiency with Heavy armor") pass unchecked; add a rule when a source uses one.
	}
	return true
}

var (
	reFeatCantrips = regexp.MustCompile(`learn (\w+) cantrips of your choice from the (.+?) spell list`)
	reFeatSpell    = regexp.MustCompile(`Choose a level 1 spell from the same list`)
	reFeatProfs    = regexp.MustCompile(`proficiency in any combination of (\w+) skills or tools`)
	reFeatRaise    = regexp.MustCompile(`Increase (?:your (\w+) or (\w+) score|one ability score of your choice) by 1, to a maximum of (\d+)`)
	reToolCost     = regexp.MustCompile(`\s*\(\d+ GP\)$`)
)

// featChoices are the choices a feat's text asks for: Magic Initiate's spell
// list, cantrips, level 1 spell and spellcasting ability, Skilled's skills
// and tools, and an ability increase (Grappler, Epic Boons). name is the
// feature on the sheet; list is the spell list a background names. Lists the
// sheet took with this feat already are left out (Repeatable).
func (a *App) featChoices(ctx context.Context, key, name string, td *compendium.TextData, list string, sh dnd5e.Sheet, campaignID int64) ([]choiceGroup, error) {
	text := td.Desc
	for _, p := range td.Parts {
		text += "\n" + p.Desc
	}
	count := func(w string) int {
		n, _ := strconv.Atoi(cmp.Or(compendium.CountWords[strings.ToLower(w)], w))
		return n
	}
	var out []choiceGroup
	if m := reFeatCantrips.FindStringSubmatch(text); m != nil {
		lists := compendium.SplitList(m[2])
		if list != "" {
			lists = []string{list}
		}
		added := map[string]bool{} // the spells the picks add, for the ability
		perList := map[string][]choiceGroup{}
		var opts []compendium.NameDesc
		for _, l := range lists {
			if list == "" && slices.Contains(sh.Feats, fmt.Sprintf("%s (%s)", name, l)) {
				continue
			}
			full := name
			if list == "" {
				full = fmt.Sprintf("%s (%s)", name, l)
			}
			groups, err := a.featSpellGroups(ctx, key+"_"+compendium.Slugify(l), full, l, count(m[1]), reFeatSpell.MatchString(text), sh, added)
			if err != nil {
				return nil, err
			}
			if len(groups) > 0 {
				perList[l] = groups
				opts = append(opts, compendium.NameDesc{Name: l})
			}
		}
		if list != "" {
			out = append(out, perList[list]...)
		} else if len(opts) > 0 {
			g := views.ChoiceGroup{Key: key + "_list", Title: name + ": spell list", Count: 1, Options: opts}
			out = append(out, choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
				full := fmt.Sprintf("%s (%s)", name, picked[0])
				if i := slices.IndexFunc(s.Features, func(f dnd5e.Feature) bool { return f.Name == name && f.Source == "Feat" }); i >= 0 {
					s.Features[i].Name = full
				}
				if i := slices.Index(s.Feats, name); i >= 0 {
					s.Feats[i] = full
				}
			}, perList})
		}
		if len(out) > 0 && reSpellAbility.MatchString(text) {
			out = append(out, spellAbilityChoice(key+"_ability", name, added, false))
		}
	}
	if m := reFeatRaise.FindStringSubmatch(text); m != nil {
		allowed := dnd5e.Abilities
		if m[1] != "" {
			allowed = abilityKeys(m[1] + ", " + m[2])
		}
		limit, _ := strconv.Atoi(m[3])
		var opts []compendium.NameDesc
		for _, ab := range allowed {
			if sh.Scores[ab] < limit {
				opts = append(opts, compendium.NameDesc{Name: dnd5e.AbilityNames[ab]})
			}
		}
		g := views.ChoiceGroup{Key: key + "_raise", Title: name + ": +1 to an ability score", Count: 1, Options: opts}
		out = append(out, choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
			if ab := abilityKeys(picked[0]); len(ab) == 1 {
				s.Scores[ab[0]]++
			}
		}, nil})
	}
	if m := reFeatProfs.FindStringSubmatch(text); m != nil {
		var opts []compendium.NameDesc
		for _, sk := range dnd5e.Skills {
			if !slices.Contains(sh.SkillProfs, sk.Key) {
				opts = append(opts, compendium.NameDesc{Name: sk.Name})
			}
		}
		tools, err := a.toolNames(ctx, campaignID)
		if err != nil {
			return nil, err
		}
		for _, t := range tools {
			if !slices.Contains(sh.OtherProfs, t) {
				opts = append(opts, compendium.NameDesc{Name: t})
			}
		}
		g := views.ChoiceGroup{Key: key + "_profs", Title: name + ": skills or tools", Count: count(m[1]), Options: opts}
		out = append(out, choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
			tools := slices.DeleteFunc(slices.Clone(picked), func(p string) bool {
				return slices.ContainsFunc(dnd5e.Skills, func(sk dnd5e.Skill) bool { return sk.Name == p })
			})
			addProfs(s, skillKeys(picked), tools...)
		}, nil})
	}
	return out, nil
}

// featSpellGroups ask for a feat's cantrips and level 1 spell (always
// prepared, one free cast per Long Rest) from a class's list, leaving out
// spells the sheet has; added collects the picks.
func (a *App) featSpellGroups(ctx context.Context, key, name, class string, cantrips int, spell bool, sh dnd5e.Sheet, added map[string]bool) ([]choiceGroup, error) {
	rows, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: compendium.Slugify(class), MaxLevel: 1})
	if err != nil {
		return nil, err
	}
	var byLevel [2][]compendium.NameDesc
	spells := map[string]dnd5e.Spell{}
	for _, r := range rows {
		if slices.ContainsFunc(sh.Spells, func(have dnd5e.Spell) bool { return have.Slug == r.Slug }) {
			continue
		}
		byLevel[r.Level] = append(byLevel[r.Level], compendium.NameDesc{Name: r.Name})
		spells[r.Name] = dnd5e.Spell{Slug: r.Slug, Name: r.Name, Level: int(r.Level), Feat: name}
	}
	add := func(s *dnd5e.Sheet, picked []string) {
		for _, n := range picked {
			sp := spells[n]
			if sp.Level > 0 {
				sp.Prepared = true
				s.PutFeature(freeCast(sp, name), nil)
			}
			s.AddSpell(sp)
			added[sp.Slug] = true
		}
	}
	var out []choiceGroup
	if cantrips > 0 && len(byLevel[0]) > 0 {
		out = append(out, choiceGroup{views.ChoiceGroup{Key: key + "_cantrips", Title: name + ": cantrips", Count: cantrips, Options: byLevel[0]}, add, nil})
	}
	if spell && len(byLevel[1]) > 0 {
		out = append(out, choiceGroup{views.ChoiceGroup{Key: key + "_spell", Title: name + ": level 1 spell", Count: 1, Options: byLevel[1]}, add, nil})
	}
	return out, nil
}

// keepSpells is Spell Change's option to change nothing.
const keepSpells = "No change"

// featSpellChanges offer each feat's Spell Change on level-up: one spell the
// feat chose ("Magic Initiate (Wizard)") may become another of the same
// level from its list.
func (a *App) featSpellChanges(ctx context.Context, sh dnd5e.Sheet) ([]choiceGroup, error) {
	var out []choiceGroup
	for _, feat := range sh.Feats {
		_, list := featList(feat)
		mine := slices.DeleteFunc(slices.Clone(sh.Spells), func(sp dnd5e.Spell) bool { return sp.Feat != feat })
		if list == "" || len(mine) == 0 {
			continue
		}
		rows, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: compendium.Slugify(list), MaxLevel: 1})
		if err != nil {
			return nil, err
		}
		key := "change_" + compendium.Slugify(feat)
		opts := []compendium.NameDesc{{Name: keepSpells}}
		sub := map[string][]choiceGroup{}
		for _, old := range mine {
			var with []compendium.NameDesc
			spells := map[string]dnd5e.Spell{}
			for _, r := range rows {
				if int(r.Level) == old.Level && !slices.ContainsFunc(sh.Spells, func(h dnd5e.Spell) bool { return h.Slug == r.Slug }) {
					with = append(with, compendium.NameDesc{Name: r.Name})
					spells[r.Name] = dnd5e.Spell{Slug: r.Slug, Name: r.Name, Level: int(r.Level)}
				}
			}
			if len(with) == 0 {
				continue
			}
			opts = append(opts, compendium.NameDesc{Name: old.Name})
			g := views.ChoiceGroup{Key: key + "_" + old.Slug, Title: "Replace " + old.Name + " with", Count: 1, Options: with}
			sub[old.Name] = []choiceGroup{{g, func(s *dnd5e.Sheet, picked []string) { swapFeatSpell(s, old, spells[picked[0]]) }, nil}}
		}
		if len(sub) > 0 {
			out = append(out, choiceGroup{views.ChoiceGroup{Key: key, Title: feat + ": Spell Change", Count: 1, Options: opts}, nil, sub})
		}
	}
	return out, nil
}

// swapFeatSpell puts repl in old's place, prepared and cast as old was, and
// renames old's free cast. A repl the same level-up already taught (ticked in
// the class list) leaves old as it is.
func swapFeatSpell(s *dnd5e.Sheet, old, repl dnd5e.Spell) {
	i := slices.IndexFunc(s.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == old.Slug })
	if i < 0 || slices.ContainsFunc(s.Spells, func(sp dnd5e.Spell) bool { return sp.Slug == repl.Slug }) {
		return
	}
	repl.Prepared, repl.Ability, repl.Feat = s.Spells[i].Prepared, s.Spells[i].Ability, s.Spells[i].Feat
	s.Spells[i] = repl
	if j := slices.IndexFunc(s.Features, func(f dnd5e.Feature) bool { return f.Name == freeCast(old, "").Name }); j >= 0 {
		fc := freeCast(repl, s.Features[j].Source)
		s.Features[j].Name, s.Features[j].Description = fc.Name, fc.Description
	}
}

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
// also returns the spells its options can grant.
func (a *App) lineageChoice(ctx context.Context, p compendium.NameDesc, opts []compendium.NameDesc, species string, resists bool, campaignID int64) (choiceGroup, []dnd5e.Spell, error) {
	cantrips := map[string][]dnd5e.Spell{}
	var all []dnd5e.Spell
	for _, o := range opts {
		spells, err := a.spellsNamed(ctx, compendium.Cantrips(o.Desc), campaignID)
		if err != nil {
			return choiceGroup{}, nil, err
		}
		cantrips[o.Name] = spells
		all = append(all, spells...)
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
			addResistance(s, o.Desc)
		}
	}, nil}, all, nil
}

// skillPick reads "proficiency in the Insight, Perception, or Survival
// skill" or "in one skill of your choice" as a choice of one skill.
func skillPick(p compendium.NameDesc) (choiceGroup, bool) {
	if !strings.Contains(p.Desc, "proficiency in") || !strings.Contains(p.Desc, "skill") {
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
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) { addProfs(s, skillKeys(picked)) }, nil}, true
}

// classChoices offers the class's spells and its "… Options" features
// (Eldritch Invocation Options) as many as its table column gives.
func (a *App) classChoices(ctx context.Context, cr *creation, sh dnd5e.Sheet) error {
	cl := sh.Classes[0]
	row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindClass), Slug: cl.Slug})
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}
	info, err := classInfo(row)
	if err != nil {
		return err
	}
	cr.view.Columns = info.ColumnsAt(cl.Level)
	slots, _, pactLevel := dnd5e.SlotMaxima(sh.Classes)
	if cr.view.Spells, err = a.classSpellOptions(ctx, cl.Slug, slots, pactLevel, sh); err != nil {
		return err
	}
	cr.groups = append(cr.groups, classOptionGroups(info, cl.Level, 0, sh.Features)...)
	return nil
}

// classOptionGroups offer a class's "… Options" features (Eldritch
// Invocation Options) as many times as the matching table column grows
// from class level before to level. Options above level, or already taken
// and not repeatable, are left out.
func classOptionGroups(info dnd5e.ClassInfo, level, before int, have []dnd5e.Feature) []choiceGroup {
	now, was := info.ColumnsAt(level), info.ColumnsAt(before)
	var out []choiceGroup
	for _, f := range info.Features {
		kind, ok := strings.CutSuffix(f.Name, " Options")
		if !ok {
			continue
		}
		var opts []compendium.NameDesc
		for _, o := range compendium.Options(f.Description) {
			taken := slices.ContainsFunc(have, func(h dnd5e.Feature) bool { return h.Name == kind+": "+o.Name })
			if compendium.OptionLevel(o.Desc) <= level && (!taken || strings.Contains(o.Desc, "**Repeatable.**")) {
				opts = append(opts, o)
			}
		}
		for _, col := range slices.Sorted(maps.Keys(now)) {
			n, _ := strconv.Atoi(now[col])
			m, _ := strconv.Atoi(was[col])
			if n <= m || len(opts) == 0 || !strings.HasPrefix(f.Name, strings.TrimSuffix(col, "s")) {
				continue
			}
			g := views.ChoiceGroup{Key: "class_" + compendium.Slugify(kind), Title: col, Count: n - m, Options: opts}
			out = append(out, choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
				for _, name := range picked {
					o := opts[slices.IndexFunc(opts, func(o compendium.NameDesc) bool { return o.Name == name })]
					name := kind + ": " + o.Name
					// A Repeatable option taken again is a new feature, not a replacement.
					for n := 2; slices.ContainsFunc(s.Features, func(h dnd5e.Feature) bool { return strings.EqualFold(h.Name, name) }); n++ {
						name = fmt.Sprintf("%s: %s (%d)", kind, o.Name, n)
					}
					feat := dnd5e.Feature{Name: name, Source: info.Name, Level: level, Description: o.Desc}
					mods := dnd5e.ParseEffects(feat.Name, feat.Description, info.Table, info.Name).Apply(&feat)
					s.PutFeature(feat, mods)
				}
			}, nil})
		}
	}
	return out
}

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

// wizardSave applies a step and moves to the next.
func (a *App) wizardSave(c *gin.Context) {
	w, ch, step, ok := a.wizardAccess(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	cmd, err := a.wizardCommand(c, ch, step)
	if err == nil && cmd != nil {
		err = a.runCommand(ctx, ch, *cmd)
	}
	var ue userError
	if errors.As(err, &ue) {
		v, verr := a.wizardView(ctx, w, ch, step, c.PostForm("pick"))
		if verr != nil {
			a.serverError(c, verr)
			return
		}
		v.Error = ue.Error()
		a.renderWizard(c, http.StatusUnprocessableEntity, w, ch, v)
		return
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	topics, err := a.characterTopics(ctx, ch.ID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(ch.CampaignID, topics...)
	c.Redirect(http.StatusSeeOther, nextStep(ch.ID, step))
}

// wizardCommand is the change a step's form asks for; nil when the step
// changes nothing (it was applied already, or skipped).
func (a *App) wizardCommand(c *gin.Context, ch sqlc.Character, step string) (*dnd5e.Command, error) {
	ctx := c.Request.Context()
	pick := c.PostForm("pick")
	once := func(apply func(s *dnd5e.Sheet, cc *dnd5e.Character) error) *dnd5e.Command {
		cmd := dnd5e.NewCommand("create "+step, perm.ScopeSheet, func(cc *dnd5e.Character) error {
			if cc.Sheet.SetUp(step) {
				return nil
			}
			if err := apply(&cc.Sheet, cc); err != nil {
				return err
			}
			cc.Sheet.SetupDone = append(cc.Sheet.SetupDone, step)
			return nil
		})
		return &cmd
	}
	switch step {
	case "class":
		if pick == "" {
			return nil, nil
		}
		row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindClass), Slug: pick})
		if err != nil {
			return nil, err
		}
		info, err := classInfo(row)
		if err != nil {
			return nil, err
		}
		data, _ := compendium.DecodeData(compendium.KindClass, row.DataJson)
		choice := wizardChoice(row, data)
		skills := c.PostFormArray("skill")
		if len(skills) > choice.SkillCount || slices.ContainsFunc(skills, func(k string) bool { return !slices.Contains(choice.Skills, k) }) {
			return nil, userError{fmt.Errorf("choose up to %d of the listed skills", choice.SkillCount)}
		}
		return once(func(s *dnd5e.Sheet, cc *dnd5e.Character) error {
			if s.Level() > 0 {
				return nil
			}
			if err := cc.LevelUp(dnd5e.LevelUpChoice{Class: info}); err != nil {
				return err
			}
			addProfs(&cc.Sheet, skills, trainedIn(choice.Traits)...)
			return nil
		}), nil
	case "species", "background":
		if pick == "" {
			return nil, nil
		}
		kind := compendium.Kind(step)
		row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(kind), Slug: pick})
		if err != nil {
			return nil, err
		}
		data, err := compendium.DecodeData(kind, row.DataJson)
		if err != nil {
			return nil, err
		}
		td := data.(*compendium.TextData)
		if step == "species" {
			return once(func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
				applySpecies(s, row.Name, td)
				return nil
			}), nil
		}
		raises, err := backgroundRaises(c, wizardChoice(row, data).Abilities)
		if err != nil {
			return nil, err
		}
		feat, err := a.originFeat(ctx, td)
		if err != nil {
			return nil, err
		}
		return once(func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
			return applyBackground(s, row.Name, td, raises, feat)
		}), nil
	case "abilities":
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil || sheet.SetUp("background") {
			return nil, err // base scores would undo the background's increases
		}
		scores := map[string]int{}
		for _, ab := range dnd5e.Abilities {
			n, err := strconv.Atoi(c.PostForm(ab))
			if err != nil || n < 3 || n > 20 {
				return nil, userError{fmt.Errorf("%s must be 3 to 20", dnd5e.AbilityNames[ab])}
			}
			scores[ab] = n
		}
		cmd := dnd5e.NewCommand("create abilities", perm.ScopeSheet, func(cc *dnd5e.Character) error {
			before := cc.Sheet.Mod("con")
			for ab, n := range scores {
				cc.Sheet.Scores[ab] = n
			}
			// Hit points already gained follow the new Constitution.
			delta := (cc.Sheet.Mod("con") - before) * cc.Sheet.Level()
			cc.MaxHP, cc.HP = max(cc.MaxHP+delta, 0), max(cc.HP+delta, 0)
			return nil
		})
		return &cmd, nil
	case "choices":
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil || len(sheet.Classes) == 0 || sheet.Species == "" {
			return nil, err // saved now, the step could not ask for them later
		}
		cr, err := a.creationChoices(ctx, sheet, ch.CampaignID)
		if err != nil {
			return nil, err
		}
		picks, err := readPicks(c, cr.groups)
		if err != nil {
			return nil, err
		}
		var spells []dnd5e.Spell
		for _, slug := range c.PostFormArray("spell") {
			if !slices.ContainsFunc(cr.view.Spells, func(o views.Option) bool { return o.Slug == slug }) {
				return nil, userError{errors.New("choose spells from the list")}
			}
			sp, err := a.compendiumSpell(ctx, slug, ch.CampaignID)
			if err != nil {
				return nil, err
			}
			spells = append(spells, sp)
		}
		return once(func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
			// The species' spells come before the picks, which set their
			// ability; the class's after, so they keep the class's.
			for _, sp := range cr.cantrips {
				s.AddSpell(sp)
			}
			applyPicks(s, cr.groups, picks)
			for _, sp := range spells {
				s.AddSpell(sp)
			}
			return nil
		}), nil
	case "equipment":
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil {
			return nil, err
		}
		choices, err := a.equipmentChoices(ctx, sheet)
		if err != nil {
			return nil, err
		}
		var picked []compendium.EquipmentOption
		for _, ec := range choices {
			letter := c.PostForm("option_" + ec.Key)
			if i := slices.IndexFunc(ec.Options, func(o compendium.EquipmentOption) bool { return o.Letter == letter }); i >= 0 {
				picked = append(picked, ec.Options[i])
			}
		}
		gear, err := a.gearFor(ctx, picked)
		if err != nil {
			return nil, err
		}
		return once(func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
			applyEquipment(s, picked, gear)
			return nil
		}), nil
	case "details":
		name := strings.TrimSpace(c.PostForm("name"))
		if name == "" {
			return nil, userError{errors.New("give your character a name")}
		}
		backstory, vision := c.PostForm("backstory"), strings.TrimSpace(c.PostForm("vision"))
		size := dnd5e.SizeNamed(c.PostForm("size"))
		if size == "" {
			return nil, userError{errors.New("pick a size")}
		}
		cmd := dnd5e.NewCommand("create details", perm.ScopeSheet, func(cc *dnd5e.Character) error {
			cc.Name, cc.Sheet.Backstory, cc.Sheet.Size, cc.Sheet.Vision = name, backstory, size, vision
			return nil
		})
		return &cmd, nil
	case "review":
		return once(func(*dnd5e.Sheet, *dnd5e.Character) error { return nil }), nil
	}
	return nil, nil
}

// trainedIn are a class's armor, weapon and tool proficiencies, as other
// proficiencies (tool choices are left to the player).
func trainedIn(traits map[string]string) []string {
	var out []string
	for _, k := range []string{"Armor Training", "Weapon Proficiencies", "Tool Proficiencies"} {
		if v := traits[k]; v != "" && !strings.HasPrefix(strings.ToLower(v), "choose") {
			out = append(out, v)
		}
	}
	return out
}

func addProfs(s *dnd5e.Sheet, skills []string, other ...string) {
	for _, k := range skills {
		if !slices.Contains(s.SkillProfs, k) {
			s.SkillProfs = append(s.SkillProfs, k)
		}
	}
	for _, o := range other {
		if !slices.Contains(s.OtherProfs, o) {
			s.OtherProfs = append(s.OtherProfs, o)
		}
	}
}

var reFeet = regexp.MustCompile(`(\d+)\s*feet`)

// applySpecies sets the species, its speed, size and darkvision, and its
// traits as features with what the sheet can automate.
func applySpecies(s *dnd5e.Sheet, name string, d *compendium.TextData) {
	s.Species = name
	for _, p := range d.Parts {
		// A choice's text holds every option's; the picked one applies in Choices.
		if p.Type == "" && choiceOptions(p) == nil {
			applyTraitText(s, p.Name+". "+p.Desc)
		}
		switch {
		case p.Type == "speed" || p.Name == "Speed":
			if m := reFeet.FindStringSubmatch(p.Desc); m != nil {
				s.Speed, _ = strconv.Atoi(m[1])
			}
		case p.Type == "size" || p.Name == "Size":
			// "Medium or Small" names no single size: the player picks it in Details.
			if size := dnd5e.SizeNamed(strings.TrimSuffix(strings.TrimSpace(p.Desc), ".")); size != "" {
				s.Size = size
			}
		default:
			f := dnd5e.Feature{Name: p.Name, Source: name, Description: p.Desc}
			mods := dnd5e.ParseEffects(p.Name, p.Desc, nil, "").Apply(&f)
			s.PutFeature(f, mods)
		}
	}
}

// backgroundRaises reads the background's ability increases: +2 and +1 to
// two of its abilities, or +1 to all three.
func backgroundRaises(c *gin.Context, allowed []string) (map[string]int, error) {
	raises := map[string]int{}
	if c.PostForm("mode") == "three" {
		for _, a := range allowed {
			raises[a] = 1
		}
		return raises, nil
	}
	plus2, plus1 := c.PostForm("plus2"), c.PostForm("plus1")
	if plus2 == "" && plus1 == "" {
		return raises, nil
	}
	if !slices.Contains(allowed, plus2) || !slices.Contains(allowed, plus1) || plus1 == plus2 {
		return nil, userError{errors.New("choose two different abilities of this background for +2 and +1")}
	}
	raises[plus2], raises[plus1] = 2, 1
	return raises, nil
}

// originFeat is the background's feat from the compendium, if it is there.
func (a *App) originFeat(ctx context.Context, d *compendium.TextData) (*dnd5e.Feature, error) {
	for _, p := range d.Parts {
		if p.Type != "feat" {
			continue
		}
		name, _, _ := strings.Cut(p.Desc, " (")
		td, err := a.textEntry(ctx, compendium.KindFeat, compendium.Slugify(name))
		if err != nil {
			return nil, err
		}
		f := &dnd5e.Feature{Name: strings.TrimSpace(p.Desc), Source: "Feat"}
		if td != nil {
			f.Description = textMarkdown(td)
		}
		return f, nil
	}
	return nil, nil
}

func applyBackground(s *dnd5e.Sheet, name string, d *compendium.TextData, raises map[string]int, feat *dnd5e.Feature) error {
	s.Background = name
	for ab, n := range raises {
		if s.Scores[ab]+n > 20 {
			return userError{fmt.Errorf("%s cannot go above 20", dnd5e.AbilityNames[ab])}
		}
		s.Scores[ab] += n
	}
	for _, p := range d.Parts {
		switch p.Type {
		case "skill_proficiency":
			addProfs(s, skillKeys(compendium.SplitList(p.Desc)))
		case "tool_proficiency":
			addProfs(s, nil, p.Desc)
		}
	}
	if feat != nil {
		addFeat(s, *feat)
	}
	return nil
}

// addFeat puts a feat on the sheet with what the sheet can automate.
func addFeat(s *dnd5e.Sheet, f dnd5e.Feature) {
	mods := dnd5e.ParseEffects(f.Name, f.Description, nil, "").Apply(&f)
	s.PutFeature(f, mods)
	s.Feats = append(s.Feats, f.Name)
}

// gear is what starting equipment items are in the compendium: weapons to
// put under attacks, armor to wear.
type gear struct {
	weapons []dnd5e.Weapon
	armor   *dnd5e.WornArmor
	shield  bool
	// names maps an equipment line ("Daggers") to its compendium entry
	// ("Dagger"), so the inventory and the weapon share one item.
	names map[string]string
}

func (a *App) gearFor(ctx context.Context, picked []compendium.EquipmentOption) (gear, error) {
	g := gear{names: map[string]string{}}
	for _, o := range picked {
		for _, it := range o.Items {
			e, err := a.findItem(ctx, it.Name)
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			if err != nil {
				return g, err
			}
			g.names[it.Name] = e.Name
			d, err := compendium.DecodeData(compendium.KindItem, e.DataJson)
			if err != nil {
				return g, err
			}
			item := d.(*compendium.ItemData)
			switch {
			case item.Damage != "":
				g.weapons = append(g.weapons, compendiumWeapon(e.Name, item))
			case item.Armor != nil && item.Armor.Type == compendium.ArmorShield:
				g.shield = true
			case item.Armor != nil && g.armor == nil:
				ar := item.Armor
				g.armor = &dnd5e.WornArmor{Slug: e.Slug, Name: e.Name, Type: ar.Type, Base: ar.Base, AddDex: ar.AddDex, DexCap: ar.DexCap, StrReq: ar.StrReq, Stealth: ar.Stealth}
			}
		}
	}
	return g, nil
}

// findItem looks an equipment line up by name, singular too ("Daggers",
// "Pouches").
func (a *App) findItem(ctx context.Context, name string) (sqlc.CompendiumEntry, error) {
	slug := compendium.Slugify(name)
	e, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindItem), Slug: slug})
	for _, suffix := range []string{"s", "es"} {
		if !errors.Is(err, sql.ErrNoRows) || !strings.HasSuffix(slug, suffix) {
			break
		}
		e, err = a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindItem), Slug: strings.TrimSuffix(slug, suffix)})
	}
	return e, err
}

func applyEquipment(s *dnd5e.Sheet, picked []compendium.EquipmentOption, g gear) {
	for _, o := range picked {
		s.Coins.GP += o.GP
		for _, it := range o.Items {
			name := cmp.Or(g.names[it.Name], it.Name)
			if i := slices.IndexFunc(s.Inventory, func(x dnd5e.Item) bool { return strings.EqualFold(x.Name, name) }); i >= 0 {
				s.Inventory[i].Qty += it.Qty
			} else {
				s.Inventory = append(s.Inventory, dnd5e.Item{Name: name, Qty: it.Qty})
			}
		}
	}
	for _, w := range g.weapons {
		if !slices.ContainsFunc(s.Weapons, func(x dnd5e.Weapon) bool { return x.Name == w.Name }) {
			s.Weapons = append(s.Weapons, w)
		}
	}
	if s.Armor == nil && g.armor != nil {
		s.Armor = g.armor
	}
	s.Shield = s.Shield || g.shield
}
