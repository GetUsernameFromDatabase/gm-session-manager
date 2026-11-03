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
}

// creationChoices are the picks the species and the level-1 class leave
// open: a lineage, a skill, two languages, invocations and spells.
func (a *App) creationChoices(ctx context.Context, sh dnd5e.Sheet, campaignID int64) (creation, error) {
	var cr creation
	if err := a.speciesChoices(ctx, &cr, sh.Species, campaignID); err != nil {
		return cr, err
	}
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
	}})
	cr.view.Groups = groupViews(cr.groups)
	return cr, nil
}

func groupViews(groups []choiceGroup) []views.ChoiceGroup {
	out := make([]views.ChoiceGroup, len(groups))
	for i, g := range groups {
		out[i] = g.ChoiceGroup
	}
	return out
}

// readPicks reads each group's picks from the form: up to its count, only
// of its options, each once.
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
	}
	return picks, nil
}

func applyPicks(s *dnd5e.Sheet, groups []choiceGroup, picks map[string][]string) {
	for _, g := range groups {
		if len(picks[g.Key]) > 0 {
			g.apply(s, picks[g.Key])
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

// speciesChoices asks the species' "Choose a lineage" traits and skill
// choices, and collects the cantrips its other traits grant.
func (a *App) speciesChoices(ctx context.Context, cr *creation, species string, campaignID int64) error {
	if species == "" {
		return nil
	}
	row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindSpecies), Slug: compendium.Slugify(species)})
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}
	d, err := compendium.DecodeData(compendium.KindSpecies, row.DataJson)
	if err != nil {
		return err
	}
	parts := d.(*compendium.TextData).Parts
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
			g, ok, err := a.originFeatChoice(ctx, p)
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
		cr.groups = append(cr.groups, spellAbilityChoice(abilityFor, spellSlugs))
	}
	return nil
}

// spellAbilityChoice picks Intelligence, Wisdom or Charisma for the spells
// a species grants. It runs after the spells are added.
func spellAbilityChoice(trait string, slugs map[string]bool) choiceGroup {
	var opts []compendium.NameDesc
	for _, ab := range []string{"int", "wis", "cha"} {
		opts = append(opts, compendium.NameDesc{Name: dnd5e.AbilityNames[ab]})
	}
	g := views.ChoiceGroup{Key: "trait_" + compendium.Slugify(trait) + "_ability", Title: "Spellcasting ability for " + trait, Count: 1, Options: opts}
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
		for ab, name := range dnd5e.AbilityNames {
			if name != picked[0] {
				continue
			}
			for i, sp := range s.Spells {
				if slugs[sp.Slug] {
					s.Spells[i].Ability = ab
				}
			}
		}
	}}
}

// originFeatChoice offers the compendium's Origin feats (Human Versatile).
func (a *App) originFeatChoice(ctx context.Context, p compendium.NameDesc) (choiceGroup, bool, error) {
	names, err := a.Store.ListCompendiumNames(ctx, string(compendium.KindFeat))
	if err != nil {
		return choiceGroup{}, false, err
	}
	var opts []compendium.NameDesc
	for _, n := range names {
		row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindFeat), Slug: n.Slug})
		if err != nil {
			return choiceGroup{}, false, err
		}
		d, err := compendium.DecodeData(compendium.KindFeat, row.DataJson)
		if err != nil {
			return choiceGroup{}, false, err
		}
		if td := d.(*compendium.TextData); td.Category == "Origin" {
			opts = append(opts, compendium.NameDesc{Name: row.Name, Desc: textMarkdown(td)})
		}
	}
	if len(opts) == 0 {
		return choiceGroup{}, false, nil
	}
	g := views.ChoiceGroup{Key: "trait_" + compendium.Slugify(p.Name), Title: p.Name + ": an Origin feat", Count: 1, Options: opts}
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
		o := opts[slices.IndexFunc(opts, func(o compendium.NameDesc) bool { return o.Name == picked[0] })]
		addFeat(s, dnd5e.Feature{Name: o.Name, Source: "Feat", Description: o.Desc})
	}}, true, nil
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
	}}, all, nil
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
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) { addProfs(s, skillKeys(picked)) }}, true
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
			}})
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
		row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindFeat), Slug: compendium.Slugify(name)})
		if errors.Is(err, sql.ErrNoRows) {
			return &dnd5e.Feature{Name: strings.TrimSpace(p.Desc), Source: "Feat"}, nil
		}
		if err != nil {
			return nil, err
		}
		data, err := compendium.DecodeData(compendium.KindFeat, row.DataJson)
		if err != nil {
			return nil, err
		}
		return &dnd5e.Feature{Name: strings.TrimSpace(p.Desc), Source: "Feat", Description: textMarkdown(data.(*compendium.TextData))}, nil
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
