package web

import (
	"cmp"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// wizardSteps are the creation wizard's steps, in order: what changes the
// most comes first. Every step but details adds to the sheet, so each runs
// once; it can be redone (wizardChange) until the review is saved, then the
// sheet is changed directly.
var wizardSteps = []views.WizardStep{
	{Key: views.StepClass, Title: "Class"},
	{Key: views.StepBackground, Title: "Background"},
	{Key: views.StepSpecies, Title: "Species"},
	{Key: views.StepAbilities, Title: "Ability scores"},
	{Key: views.StepChoices, Title: "Choices"},
	{Key: views.StepEquipment, Title: "Equipment"},
	{Key: views.StepDetails, Title: "Details"},
	{Key: views.StepReview, Title: "Review"},
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

// stepNeeds are the earlier picks a step asks about: saved without them, it
// could never ask for what they bring (a background feat's spells, a class's gear).
var stepNeeds = map[string][]string{
	views.StepAbilities: {views.StepBackground},
	views.StepChoices:   {views.StepClass, views.StepBackground, views.StepSpecies},
	views.StepEquipment: {views.StepClass, views.StepBackground},
}

// missingSteps are the steps step needs that the sheet lacks. A background
// is needed only when the compendium has one to pick.
func (a *App) missingSteps(ctx context.Context, sh dnd5e.Sheet, step string, campaignID int64) ([]views.WizardStep, error) {
	has := map[string]bool{views.StepClass: len(sh.Classes) > 0, views.StepSpecies: sh.Species != "", views.StepBackground: sh.Background != ""}
	if !has[views.StepBackground] && slices.Contains(stepNeeds[step], views.StepBackground) {
		names, err := a.Store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: enum.KindBackground, CampaignID: campaignID})
		if err != nil {
			return nil, err
		}
		has[views.StepBackground] = len(names) == 0
	}
	var out []views.WizardStep
	for _, s := range wizardSteps {
		if slices.Contains(stepNeeds[step], s.Key) && !has[s.Key] {
			out = append(out, s)
		}
	}
	return out, nil
}

// legacyBackground reports a sheet whose background step added the ability
// increases, before they moved to the abilities step.
// shortcut: also true for a finished new sheet that skipped the abilities, which then shows the "set on the sheet" text; drop once no sheet predates the move.
func legacyBackground(sh dnd5e.Sheet) bool {
	_, snap := sh.SetupSnapshots[views.StepBackground]
	return sh.SetUp(views.StepBackground) && !snap && !sh.SetUp(views.StepAbilities)
}

// backgroundChoice is the sheet's background as the wizard shows it, for its
// ability increases; nil without one in the compendium with three abilities.
func (a *App) backgroundChoice(ctx context.Context, sh dnd5e.Sheet) (*views.WizardChoice, error) {
	if sh.Background == "" {
		return nil, nil
	}
	row, _, err := entryFor(ctx, a.Store.Queries, enum.KindBackground, compendium.Slugify(sh.Background), sh.BackgroundSource)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	data, err := compendium.DecodeData(enum.KindBackground, row.DataJson)
	if err != nil {
		return nil, err
	}
	if c := wizardChoice(row, data); len(c.Abilities) == 3 {
		return c, nil
	}
	return nil, nil
}

// wizardChange redoes a step: the character goes back to how it was before
// it, which also undoes every step saved after it.
func (a *App) wizardChange(c *gin.Context) {
	_, ch, step, ok := a.wizardAccess(c)
	if !ok {
		return
	}
	if step != c.Param("step") { // wizardAccess falls back to the first step
		a.fail(c, http.StatusNotFound, "There is no such step.")
		return
	}
	ctx := c.Request.Context()
	err := a.updateCharacter(ctx, ch.ID, func(q *sqlc.Queries, cc *dnd5e.Character) error {
		snap, ok := cc.Sheet.SetupSnapshots[step]
		if !ok {
			return userError{errors.New("this step can no longer be changed here; change it on the sheet")}
		}
		sh, err := dnd5e.ParseSheet(string(snap.Sheet))
		if err != nil {
			return err
		}
		// The steps saved before this one can still be redone.
		for k, v := range cc.Sheet.SetupSnapshots {
			if sh.SetUp(k) {
				if sh.SetupSnapshots == nil {
					sh.SetupSnapshots = map[string]dnd5e.SetupSnapshot{}
				}
				sh.SetupSnapshots[k] = v
			}
		}
		cc.Sheet, cc.MaxHP, cc.HP = sh, snap.MaxHP, snap.HP
		// Level-ups taken since were undone too: Revert would bring back the old sheet.
		return q.DeleteLevelUps(ctx, ch.ID)
	})
	if a.commandFailed(c, err) {
		return
	}
	topics, err := a.characterTopics(ctx, ch.ID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(ch.CampaignID, topics...)
	redirect(c, wizardPath(ch.ID, step))
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
	v, err := a.wizardView(c.Request.Context(), w, ch, step, c.Query("pick"), c.Query("variant"))
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
// background being looked at, variant the subrace of a species picked.
func (a *App) wizardView(ctx context.Context, w who, ch sqlc.Character, step, pick, variant string) (views.WizardView, error) {
	s, err := a.sheetView(ctx, w.viewer, ch)
	if err != nil {
		return views.WizardView{}, err
	}
	s.CanSheet, s.CanHP, s.CanConditions = false, false, false // the review is read-only
	v := views.WizardView{CharacterID: ch.ID, Step: step, Steps: wizardSteps, Sheet: s, Picked: pick}
	if v.Missing, err = a.missingSteps(ctx, s.C.Sheet, step, ch.CampaignID); err != nil {
		return v, err
	}
	if step == views.StepAbilities {
		v.Choice, err = a.backgroundChoice(ctx, s.C.Sheet)
		v.LegacyBackground = legacyBackground(s.C.Sheet)
		return v, err
	}
	kind := map[string]enum.CompendiumKind{views.StepClass: enum.KindClass, views.StepSpecies: enum.KindSpecies, views.StepBackground: enum.KindBackground}[step]
	if kind == "" {
		switch step {
		case views.StepEquipment:
			v.Equipment, err = a.equipmentChoices(ctx, s.C.Sheet)
		case views.StepChoices:
			var cr creation
			cr, err = a.creationChoices(ctx, s.C.Sheet, ch.CampaignID)
			v.Choices = &cr.view
		}
		return v, err
	}
	names, err := a.Store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: kind, CampaignID: ch.CampaignID})
	if err != nil {
		return v, err
	}
	for _, n := range names {
		if n.SubraceOf == "" { // subraces are the variant select's
			v.Options = append(v.Options, views.Option{Slug: n.Slug, Name: n.Name})
		}
	}
	if pick == "" {
		return v, nil
	}
	e, err := a.entryBySlug(ctx, kind, pick, ch.CampaignID)
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
	if step == views.StepSpecies {
		if e, data, err = a.speciesVariant(ctx, &v, e, data, variant, ch.CampaignID); err != nil {
			return v, err
		}
	}
	v.Choice = wizardChoice(e, data)
	switch step {
	case views.StepClass:
		v.Choice.SkillPick = classSkillPick(s.C.Sheet, v.Choice.Skills, v.Choice.SkillCount, nil)
	case views.StepBackground:
		v.Choice.SkillPick, err = a.backgroundSkillPick(ctx, s.C.Sheet, v.Choice.Skills)
	}
	return v, err
}

// speciesVariant offers the picked species' subraces and returns the entry
// to preview and save: the variant when it is one of them. A subrace picked
// (a refused save shown again) shows as its base with itself as the variant.
func (a *App) speciesVariant(ctx context.Context, v *views.WizardView, e sqlc.CompendiumEntry, data any, variant string, campaignID int64) (sqlc.CompendiumEntry, any, error) {
	if td, ok := data.(*compendium.TextData); ok && td.SubraceOf != "" {
		v.Picked, v.Variant = td.SubraceOf, e.Slug
	}
	subs, err := a.Store.ListSubracesOf(ctx, sqlc.ListSubracesOfParams{CampaignID: campaignID, SpeciesSlug: v.Picked})
	if err != nil {
		return e, data, err
	}
	for _, s := range subs {
		v.Variants = append(v.Variants, views.Option{Slug: s.Slug, Name: s.Name + " · " + s.SourceKey})
	}
	if v.Variant != "" || !slices.ContainsFunc(subs, func(s sqlc.ListSubracesOfRow) bool { return s.Slug == variant }) {
		return e, data, nil
	}
	sub, err := a.entryBySlug(ctx, enum.KindSpecies, variant, campaignID)
	if err != nil {
		return e, data, err
	}
	d, err := compendium.DecodeData(enum.KindSpecies, sub.DataJson)
	if err == nil {
		v.Variant = variant
	}
	return sub, d, err
}

// wizardChoice is what a class, species or background gives, for the page
// to show and the step to apply.
func wizardChoice(e sqlc.CompendiumEntry, data any) *views.WizardChoice {
	ch := &views.WizardChoice{Slug: e.Slug, Name: e.Name}
	switch d := data.(type) {
	case *compendium.ClassData:
		ch.Traits = compendium.CoreTraits(d.Desc)
		ch.SkillCount, ch.Skills = skillChoice(ch.Traits[compendium.TraitSkillProficiencies])
		for _, f := range d.Features {
			if slices.Contains(f.Levels, 1) && !strings.Contains(strings.ToLower(f.Name), "core") {
				ch.Features = append(ch.Features, views.FeaturePreview{Name: f.Name, Desc: f.Desc, Effects: dnd5e.ParseEffects(f.Name, f.Desc, d.Table, e.Name).Summary()})
			}
		}
	case *compendium.TextData:
		for _, p := range d.Parts {
			switch p.Type {
			case compendium.PartAbilities:
				ch.Abilities = abilityKeys(p.Desc)
			case compendium.PartSkills:
				ch.Skills = skillKeys(compendium.SplitList(p.Desc))
				ch.SkillCount = len(ch.Skills)
			case "":
				// Untyped parts are traits; typed ones (size, speed, tools,
				// equipment, the feat) are applied or shown as facts.
				ch.Features = append(ch.Features, views.FeaturePreview{Name: p.Name, Desc: p.Desc, Effects: dnd5e.ParseEffects(p.Name, p.Desc, nil, "").Summary()})
			}
			if p.Type != "" || p.Name == compendium.LabelSpeed || p.Name == compendium.LabelSize {
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
	return n, skillKeys(names)
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
	check func(before, after dnd5e.Sheet, learned []knownCantrip) error
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
		addResistance(s, m[1])
	}
}

// resistsPick reports whether a trait naming p with "Resistance"
// (Dragonborn's Damage Resistance) resists the damage type picked in p.
func resistsPick(parts []compendium.NameDesc, p compendium.NameDesc) bool {
	return slices.ContainsFunc(parts, func(q compendium.NameDesc) bool {
		return q.Name != p.Name && strings.Contains(q.Desc, p.Name) && strings.Contains(q.Desc, "Resistance")
	})
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
	names, err := a.Store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: enum.KindFeat, CampaignID: campaignID})
	if err != nil {
		return choiceGroup{}, false, err
	}
	slugs := make([]string, len(names))
	for i, n := range names {
		slugs[i] = n.Slug
	}
	// One query for every feat's priority copy, as entryFor picks unpinned.
	rows, err := a.Store.ListCompendiumBySlugs(ctx, sqlc.ListCompendiumBySlugsParams{Kind: enum.KindFeat, Slugs: slugs})
	if err != nil {
		return choiceGroup{}, false, err
	}
	bySlug := make(map[string]sqlc.ListCompendiumBySlugsRow, len(rows))
	for _, r := range rows {
		bySlug[r.Slug] = r
	}
	key := "trait_" + compendium.Slugify(p.Name)
	var opts []compendium.NameDesc
	sub := map[string][]choiceGroup{}
	sources := map[string]string{}
	for _, n := range names {
		row, ok := bySlug[n.Slug]
		if !ok {
			continue
		}
		d, err := compendium.DecodeData(enum.KindFeat, row.DataJson)
		if err != nil {
			return choiceGroup{}, false, err
		}
		td := d.(*compendium.TextData)
		if td.Category != "Origin" || !meetsPrerequisite(td.Prerequisite, sh, sh.Level(), nil) {
			continue
		}
		sources[n.Name] = row.SourceKey
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
		addFeat(s, dnd5e.Feature{Name: o.Name, Source: dnd5e.SourceFeat, Description: o.Desc})
		pinFeat(s, o.Name, sources[o.Name])
	}, sub}, true, nil
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
	d, err := compendium.DecodeData(kind, row.DataJson)
	if err != nil {
		return nil, "", err
	}
	return d.(*compendium.TextData), row.SourceKey, nil
}

// pinFeat records the source of the feat's compendium copy under featKey.
func pinFeat(s *dnd5e.Sheet, name, source string) {
	if source == "" {
		return
	}
	if s.FeatSources == nil {
		s.FeatSources = map[string]string{}
	}
	s.FeatSources[featKey(name)] = source
}

// featKey is the key a feat's pin is stored under: the feat's slug without
// a "(choice)" suffix.
func featKey(name string) string {
	base, _, _ := strings.Cut(name, " (")
	return compendium.Slugify(base)
}

// backgroundFeatChoices are the choices inside the feat the background
// gave ("Magic Initiate (Wizard)").
func (a *App) backgroundFeatChoices(ctx context.Context, sh dnd5e.Sheet, campaignID int64) ([]choiceGroup, error) {
	if sh.Background == "" {
		return nil, nil
	}
	bg, _, err := a.textEntrySource(ctx, enum.KindBackground, compendium.Slugify(sh.Background), sh.BackgroundSource)
	if err != nil || bg == nil {
		return nil, err
	}
	for _, p := range bg.Parts {
		full := strings.TrimSpace(p.Desc)
		if p.Type != compendium.PartFeat || !slices.Contains(sh.Feats, full) {
			continue
		}
		name, list := featList(full)
		slug := compendium.Slugify(name)
		td, _, err := a.textEntrySource(ctx, enum.KindFeat, slug, sh.FeatSources[slug])
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
	rePrereqLevel      = regexp.MustCompile(`^Level (\d+)\+(?: \w+)?$`)
	rePrereqScore      = regexp.MustCompile(`^(.+) (\d+)\+$`)
	rePrereqFeature    = regexp.MustCompile(`^(.+) Feature$`)
	rePrereqInvocation = regexp.MustCompile(`^(.+) Invocation$`)
	rePrereqCantrip    = regexp.MustCompile(`(?i)^an? (.+?) cantrip that deals damage( via an attack roll)?$`)
	reSpellDamage      = regexp.MustCompile(`\d+d\d+ \w+ damage`)
	reSpellRange       = regexp.MustCompile(`(?i)^(\d+)\s*(?:feet|foot|ft\b\.?|(miles?)\b)`)
	reOptionCantrip    = regexp.MustCompile(`(?i)choose one of your known (\w+) cantrips that (deals damage|requires an attack roll)(?: and has a range of (\d+)\+ feet)?`)
)

// knownCantrip is a class cantrip a sheet knows, or may learn this level-up.
type knownCantrip struct {
	Name string
	compendium.SpellData
}

// cantripFits reports whether c is a cantrip of class's list that deals
// damage, by a spell attack if attack, with a range of minRange feet or more
// ("120 feet", "120 ft.", "1 mile"; Touch and Self count as 0).
func cantripFits(c knownCantrip, class string, attack bool, minRange int) bool {
	feet := 0
	if m := reSpellRange.FindStringSubmatch(c.Range); m != nil {
		feet, _ = strconv.Atoi(m[1])
		if m[2] != "" {
			feet *= 5280
		}
	}
	return slices.Contains(c.Classes, compendium.Slugify(class)) && reSpellDamage.MatchString(c.Desc) &&
		(!attack || strings.Contains(c.Desc, "spell attack")) && feet >= minRange
}

// meetsPrerequisite reports whether a sheet meets a feat's or class option's
// prerequisite ("Level 4+, Strength or Dexterity 13+", "Fighting Style
// Feature", "Level 5+ Warlock, Pact of the Blade Invocation", "a Warlock
// Cantrip That Deals Damage"); every clause must hold. level is the character
// level for feats and the class level for class options; cantrips are the
// sheet's class cantrips (cantripData).
func meetsPrerequisite(prereq string, sh dnd5e.Sheet, level int, cantrips []knownCantrip) bool {
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
		} else if m := rePrereqInvocation.FindStringSubmatch(clause); m != nil {
			if !slices.ContainsFunc(sh.Features, func(f dnd5e.Feature) bool {
				kind, name, _ := strings.Cut(f.Name, ": ")
				return strings.HasSuffix(kind, " Invocation") && name == m[1]
			}) {
				return false
			}
		} else if m := rePrereqCantrip.FindStringSubmatch(clause); m != nil {
			if !slices.ContainsFunc(cantrips, func(c knownCantrip) bool { return cantripFits(c, m[1], m[2] != "", 0) }) {
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
	reFeatRaise    = regexp.MustCompile(`Increase (?:your (?:((?:\w+, )*\w+) or )?(\w+) score|one ability score of your choice) by 1, to a maximum of (\d+)`)
	reFeatMastery  = regexp.MustCompile(`mastery property of (\w+) kinds? of`)
	reToolCost     = regexp.MustCompile(`\s*\(\d+ GP\)$`)
)

// featText is a feat's description and parts, for the reFeat patterns.
func featText(td *compendium.TextData) string {
	text := td.Desc
	for _, p := range td.Parts {
		text += "\n" + p.Desc
	}
	return text
}

// countWord reads a count written as a word ("two") or digits.
func countWord(w string) int {
	n, _ := strconv.Atoi(cmp.Or(compendium.CountWords[strings.ToLower(w)], w))
	return n
}

// featMasteries is how many weapon kinds a feat lets the sheet master
// (Weapon Master), read from its text so homebrew feats count too.
func featMasteries(td *compendium.TextData) int {
	if m := reFeatMastery.FindStringSubmatch(featText(td)); m != nil {
		return countWord(m[1])
	}
	return 0
}

// featChoices are the choices a feat's text asks for: Magic Initiate's spell
// list, cantrips, level 1 spell and spellcasting ability, Skilled's skills
// and tools, and an ability increase (Grappler, Epic Boons). name is the
// feature on the sheet; list is the spell list a background names. Lists the
// sheet took with this feat already are left out (Repeatable).
func (a *App) featChoices(ctx context.Context, key, name string, td *compendium.TextData, list string, sh dnd5e.Sheet, campaignID int64) ([]choiceGroup, error) {
	text, count := featText(td), countWord
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
			groups, err := a.featSpellGroups(ctx, key+"_"+compendium.Slugify(l), full, l, count(m[1]), reFeatSpell.MatchString(text), sh, campaignID, added)
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
				if i := slices.IndexFunc(s.Features, func(f dnd5e.Feature) bool { return f.Name == name && f.Source == dnd5e.SourceFeat }); i >= 0 {
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
		if m[2] != "" { // "your X, Y or Z", or a fixed "your X"
			allowed = abilityKeys(strings.TrimPrefix(m[1]+", "+m[2], ", "))
		}
		limit, _ := strconv.Atoi(m[3])
		var opts []compendium.NameDesc
		for _, ab := range allowed {
			if sh.Scores[ab] < limit {
				opts = append(opts, compendium.NameDesc{Name: dnd5e.AbilityNames[ab]})
			}
		}
		// A fixed "your X" already at its maximum asks nothing: the feat is still taken.
		if fixedAtMax := len(opts) == 0 && m[1] == "" && m[2] != ""; !fixedAtMax {
			g := views.ChoiceGroup{Key: key + "_raise", Title: name + ": +1 to an ability score", Count: 1, Options: opts}
			out = append(out, choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
				if ab := abilityKeys(picked[0]); len(ab) == 1 {
					s.Scores[ab[0]]++
				}
			}, nil})
		}
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
			s.AddProfs(name, skillKeys(picked), tools...)
		}, nil})
	}
	if n := featMasteries(td); n > 0 {
		groups, err := a.masteryChoice(ctx, sh, key+"_"+views.FieldMastery, name+": weapon mastery", n, campaignID)
		if err != nil {
			return nil, err
		}
		out = append(out, groups...)
	}
	return out, nil
}

// featSpellGroups ask for a feat's cantrips and level 1 spell (always
// prepared, one free cast per Long Rest) from a class's list, leaving out
// spells the sheet has; added collects the picks.
func (a *App) featSpellGroups(ctx context.Context, key, name, class string, cantrips int, spell bool, sh dnd5e.Sheet, campaignID int64, added map[string]bool) ([]choiceGroup, error) {
	rows, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: compendium.Slugify(class), MaxLevel: 1, CampaignID: campaignID})
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
		spells[r.Name] = dnd5e.Spell{Slug: r.Slug, Name: r.Name, Source: r.SourceKey, Level: int(r.Level), Feat: name}
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
// level from its list, but not one the sheet has or this step teaches (taught).
func (a *App) featSpellChanges(ctx context.Context, sh dnd5e.Sheet, taught []dnd5e.Spell, campaignID int64) ([]choiceGroup, error) {
	known := slices.Concat(sh.Spells, taught)
	var out []choiceGroup
	for _, feat := range sh.Feats {
		_, list := featList(feat)
		mine := slices.DeleteFunc(slices.Clone(sh.Spells), func(sp dnd5e.Spell) bool { return sp.Feat != feat })
		if list == "" || len(mine) == 0 {
			continue
		}
		rows, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: compendium.Slugify(list), MaxLevel: 1, CampaignID: campaignID})
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
				if int(r.Level) == old.Level && !slices.ContainsFunc(known, func(h dnd5e.Spell) bool { return h.Slug == r.Slug }) {
					with = append(with, compendium.NameDesc{Name: r.Name})
					spells[r.Name] = dnd5e.Spell{Slug: r.Slug, Name: r.Name, Source: r.SourceKey, Level: int(r.Level)}
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
// renames old's free cast. A repl the sheet already has leaves old as it is:
// featSpellChanges offers none, but this keeps a duplicate off the sheet.
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
			addResistance(s, o.Desc)
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
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) { s.AddProfs(p.Name, skillKeys(picked)) }, nil}, true
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
	cr.check = func(before, after dnd5e.Sheet, learned []knownCantrip) error {
		return checkNewOptions(info, before, after, cl.Level, learned)
	}
	return nil
}

// cantripData are the compendium entries of these spells' class cantrips
// (not a species' or a feat's), for "a Warlock Cantrip That Deals Damage". A
// cantrip learned from one class's list (Spell.Class) counts for that class
// only; older sheets without it count for every list it is on.
func (a *App) cantripData(ctx context.Context, spells []dnd5e.Spell, campaignID int64) ([]knownCantrip, error) {
	var out []knownCantrip
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
		d, err := compendium.DecodeData(enum.KindSpell, row.DataJson)
		if err != nil {
			return nil, err
		}
		c := knownCantrip{Name: sp.Name, SpellData: *d.(*compendium.SpellData)}
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
func (a *App) optionCantrips(ctx context.Context, sh dnd5e.Sheet, info dnd5e.ClassInfo, classLevel int, campaignID int64) ([]knownCantrip, error) {
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
func classOptionGroups(info dnd5e.ClassInfo, level, before int, sh dnd5e.Sheet, cantrips []knownCantrip) []choiceGroup {
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
	cantrips []knownCantrip
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
			if meetsPrerequisite(compendium.OptionPrerequisite(o.Desc), with, k.level, k.cantrips) && (!taken || isRepeatable(o)) && (!asks || len(cs) > 0) {
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
		if cantripFits(c, m[1], strings.EqualFold(m[2], "requires an attack roll"), minRange) &&
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
			return strings.HasPrefix(f.Name, kind+": ") && meetsPrerequisite(p, sh, k.level, k.cantrips) && !meetsPrerequisite(p, without, k.level, k.cantrips)
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
func checkNewOptions(info dnd5e.ClassInfo, before, after dnd5e.Sheet, level int, cantrips []knownCantrip) error {
	for _, f := range after.Features {
		kind, name, _ := strings.Cut(f.Name, ": ")
		if f.Source != info.Name || slices.ContainsFunc(before.Features, func(b dnd5e.Feature) bool { return b.Name == f.Name }) ||
			!slices.ContainsFunc(info.Features, func(cf dnd5e.ClassFeature) bool { return cf.Name == kind+" Options" }) {
			continue
		}
		base, cantrip, _ := strings.Cut(name, " (")
		if p := compendium.OptionPrerequisite(f.Description); !meetsPrerequisite(p, after, level, cantrips) {
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
		v, verr := a.wizardView(ctx, w, ch, step, c.PostForm("pick"), "")
		if verr != nil {
			a.serverError(c, verr)
			return
		}
		v.Error, v.Errors = formErrors(err, v.DrawsField)
		v.Form = c.Request.PostForm
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
	if stepNeeds[step] != nil {
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil {
			return nil, err
		}
		if missing, err := a.missingSteps(ctx, sheet, step, ch.CampaignID); err != nil || len(missing) > 0 {
			return nil, err // saved now, the step could not ask for what they bring later
		}
	}
	once := func(apply func(s *dnd5e.Sheet, cc *dnd5e.Character) error) *dnd5e.Command {
		cmd := dnd5e.NewCommand("create "+step, perm.ScopeSheet, func(cc *dnd5e.Character) error {
			if cc.Sheet.SetUp(step) {
				return nil
			}
			before := cc.Sheet
			before.SetupSnapshots = nil
			raw, err := json.Marshal(before)
			if err != nil {
				return err
			}
			snap := dnd5e.SetupSnapshot{Sheet: raw, MaxHP: cc.MaxHP, HP: cc.HP}
			if err := apply(&cc.Sheet, cc); err != nil {
				return err
			}
			cc.Sheet.SetupDone = append(cc.Sheet.SetupDone, step)
			if cc.Sheet.SetupSnapshots == nil {
				cc.Sheet.SetupSnapshots = map[string]dnd5e.SetupSnapshot{}
			}
			cc.Sheet.SetupSnapshots[step] = snap
			return nil
		})
		return &cmd
	}
	switch step {
	case views.StepClass:
		if pick == "" {
			return nil, nil
		}
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil || sheet.SetUp(step) {
			return nil, err // chosen already: the skills would be ignored
		}
		row, err := a.wizardPick(ctx, enum.KindClass, pick, ch.CampaignID)
		if err != nil {
			return nil, err
		}
		info, err := classInfo(row)
		if err != nil {
			return nil, err
		}
		data, _ := compendium.DecodeData(enum.KindClass, row.DataJson)
		choice := wizardChoice(row, data)
		skills := postedSkills(c)
		if err := checkSkillPick(classSkillPick(sheet, choice.Skills, choice.SkillCount, nil), skills); err != nil {
			return nil, err
		}
		return once(func(s *dnd5e.Sheet, cc *dnd5e.Character) error {
			if s.Level() > 0 {
				return nil
			}
			if err := cc.LevelUp(dnd5e.LevelUpChoice{Class: info}); err != nil {
				return err
			}
			cc.Sheet.AddProfs(info.Name, skills, trainedIn(choice.Traits)...)
			return nil
		}), nil
	case views.StepSpecies, views.StepBackground:
		if pick == "" {
			return nil, nil
		}
		kind := enum.CompendiumKind(step)
		row, err := a.wizardPick(ctx, kind, pick, ch.CampaignID)
		if err != nil {
			return nil, err
		}
		data, err := compendium.DecodeData(kind, row.DataJson)
		if err != nil {
			return nil, err
		}
		td := data.(*compendium.TextData)
		if step == views.StepSpecies {
			return once(func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
				applySpecies(s, row.Name, td)
				s.SpeciesSource = row.SourceKey
				return nil
			}), nil
		}
		feat, featSource, err := a.originFeat(ctx, td)
		if err != nil {
			return nil, err
		}
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil {
			return nil, err
		}
		pickAgain, err := a.backgroundSkillPick(ctx, sheet, wizardChoice(row, data).Skills)
		if err != nil {
			return nil, err
		}
		skills := postedSkills(c)
		if err := checkSkillPick(pickAgain, skills); err != nil {
			return nil, err
		}
		return once(func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
			if pickAgain != nil { // the class's picks are made again, with the background's known
				class := s.Classes[0].Name
				s.SkillProfs = slices.DeleteFunc(s.SkillProfs, func(k string) bool { return s.ProfSources[k] == class })
				maps.DeleteFunc(s.ProfSources, func(k, from string) bool { return from == class && !slices.Contains(s.OtherProfs, k) })
			}
			applyBackground(s, row.Name, td, feat)
			if pickAgain != nil {
				s.AddProfs(s.Classes[0].Name, skills)
			}
			s.BackgroundSource = row.SourceKey
			if feat != nil {
				pinFeat(s, feat.Name, featSource)
			}
			return nil
		}), nil
	case views.StepAbilities:
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil || legacyBackground(sheet) {
			return nil, err
		}
		var f abilitiesForm
		var errs []error
		if err := c.ShouldBind(&f); err != nil {
			var ve validator.ValidationErrors
			if !errors.As(err, &ve) {
				return nil, userError{errors.New("ability scores must be 3 to 20")}
			}
			for _, e := range ve {
				ab := strings.ToLower(e.StructField())
				errs = append(errs, fieldError{userError{fmt.Errorf("%s must be 3 to 20", dnd5e.AbilityNames[ab])}, ab})
			}
		}
		scores := map[string]int{"str": f.Str, "dex": f.Dex, "con": f.Con, "int": f.Int, "wis": f.Wis, "cha": f.Cha}
		bg, err := a.backgroundChoice(ctx, sheet)
		if err != nil {
			return nil, err
		}
		raises := map[string]int{}
		if bg != nil {
			if raises, err = backgroundRaises(c, bg.Abilities); err != nil {
				errs = append(errs, err)
			}
		}
		for _, ab := range dnd5e.Abilities {
			if scores[ab]+raises[ab] > 20 {
				errs = append(errs, fieldError{userError{fmt.Errorf("%s cannot go above 20", dnd5e.AbilityNames[ab])}, ab})
			}
		}
		if err := errors.Join(errs...); err != nil {
			return nil, err
		}
		return once(func(s *dnd5e.Sheet, cc *dnd5e.Character) error {
			before := s.Mod("con")
			for ab, n := range scores {
				s.Scores[ab] = n + raises[ab]
			}
			// Hit points already gained follow the new Constitution.
			delta := (s.Mod("con") - before) * s.Level()
			cc.MaxHP, cc.HP = max(cc.MaxHP+delta, 0), max(cc.HP+delta, 0)
			return nil
		}), nil
	case views.StepChoices:
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil {
			return nil, err
		}
		cr, err := a.creationChoices(ctx, sheet, ch.CampaignID)
		if err != nil {
			return nil, err
		}
		// Refused together, so every field shows its own error.
		picks, pickErr := readPicks(c, cr.groups)
		spells, err := a.classSpellPicks(c, sheet.Classes[0].Slug, cr.view.Spells, cr.view.SpellNeed, ch.CampaignID)
		if err != nil && !errors.As(err, new(userError)) {
			return nil, err
		}
		if err := errors.Join(pickErr, err); err != nil {
			return nil, err
		}
		learned, err := a.cantripData(ctx, append(slices.Clone(sheet.Spells), spells...), ch.CampaignID)
		if err != nil {
			return nil, err
		}
		return once(func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
			before := *s
			before.Features = slices.Clone(s.Features)
			// The species' spells come before the picks, which set their
			// ability; the class's after, so they keep the class's.
			for _, sp := range cr.cantrips {
				s.AddSpell(sp)
			}
			applyPicks(s, cr.groups, picks)
			if err := addClassSpells(s, spells); err != nil {
				return err
			}
			if cr.check != nil {
				return cr.check(before, *s, learned)
			}
			return nil
		}), nil
	case views.StepEquipment:
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
	case views.StepDetails:
		var f detailsForm
		var errs []error
		if err := c.ShouldBind(&f); err != nil {
			var ve validator.ValidationErrors
			if !errors.As(err, &ve) {
				return nil, userError{errors.New("keep the name to 100 characters, the vision to 200 and the backstory to 20000")}
			}
			for _, e := range ve {
				errs = append(errs, fieldError{userError{fmt.Errorf("keep it to %s characters", e.Param())}, strings.ToLower(e.StructField())})
			}
		}
		name := strings.TrimSpace(f.Name)
		if name == "" {
			errs = append(errs, fieldError{userError{errors.New("give your character a name")}, views.FieldName})
		}
		backstory, vision := f.Backstory, strings.TrimSpace(f.Vision)
		size := dnd5e.SizeNamed(f.Size)
		if size == "" {
			errs = append(errs, fieldError{userError{errors.New("pick a size")}, views.FieldSize})
		}
		if err := errors.Join(errs...); err != nil {
			return nil, err
		}
		cmd := dnd5e.NewCommand("create details", perm.ScopeSheet, func(cc *dnd5e.Character) error {
			cc.Name, cc.Sheet.Backstory, cc.Sheet.Size, cc.Sheet.Vision = name, backstory, size, vision
			return nil
		})
		return &cmd, nil
	case views.StepReview:
		cmd := dnd5e.NewCommand("create review", perm.ScopeSheet, func(cc *dnd5e.Character) error {
			if !cc.Sheet.SetUp(step) {
				cc.Sheet.SetupDone = append(cc.Sheet.SetupDone, step)
			}
			cc.Sheet.SetupSnapshots = nil // finished: steps change on the sheet from now on
			return nil
		})
		return &cmd, nil
	}
	return nil, nil
}

// wizardPick is the posted pick's entry, refused as a user error when the
// character's campaign cannot see it.
func (a *App) wizardPick(ctx context.Context, kind enum.CompendiumKind, pick string, campaignID int64) (sqlc.CompendiumEntry, error) {
	row, err := a.entryBySlug(ctx, kind, pick, campaignID)
	if errors.Is(err, sql.ErrNoRows) {
		err = userError{fmt.Errorf("that %s is not in the compendium", kind)}
	}
	return row, err
}

// trainedIn are a class's armor, weapon and tool proficiencies, as other
// proficiencies (tool choices are left to the player).
func trainedIn(traits map[string]string) []string {
	var out []string
	for _, k := range []string{compendium.TraitArmorTraining, compendium.TraitWeaponProficiencies, compendium.TraitToolProficiencies} {
		if v := traits[k]; v != "" && !strings.HasPrefix(strings.ToLower(v), "choose") {
			out = append(out, v)
		}
	}
	return out
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
		case p.Type == compendium.PartSpeed || p.Name == compendium.LabelSpeed:
			if m := reFeet.FindStringSubmatch(p.Desc); m != nil {
				s.Speed, _ = strconv.Atoi(m[1])
			}
		case p.Type == compendium.PartSize || p.Name == compendium.LabelSize:
			// "Medium or Small" names no single size: the player picks it in Details.
			if size := dnd5e.SizeNamed(strings.TrimSuffix(strings.TrimSpace(p.Desc), ".")); size != "" {
				s.Size = size
			}
		default:
			f := dnd5e.Feature{Name: p.Name, Source: name, Description: p.Desc}
			e := dnd5e.ParseEffects(p.Name, p.Desc, nil, "")
			s.PutFeature(f, e.Apply(&f))
			s.AddProfs(p.Name, e.Skills, e.Tools...)
		}
	}
}

// backgroundRaises reads the background's ability increases: +2 and +1 to
// two of its abilities, or +1 to all three.
func backgroundRaises(c *gin.Context, allowed []string) (map[string]int, error) {
	errRaises := fieldError{userError{errors.New("choose two different abilities of this background for +2 and +1")}, views.FieldMode}
	var f backgroundForm
	if err := c.ShouldBind(&f); err != nil {
		return nil, errRaises
	}
	raises := map[string]int{}
	if f.Mode == "three" {
		for _, a := range allowed {
			raises[a] = 1
		}
		return raises, nil
	}
	if f.Plus2 == "" && f.Plus1 == "" {
		return raises, nil
	}
	if !slices.Contains(allowed, f.Plus2) || !slices.Contains(allowed, f.Plus1) || f.Plus1 == f.Plus2 {
		return nil, errRaises
	}
	raises[f.Plus2], raises[f.Plus1] = 2, 1
	return raises, nil
}

// backgroundForm is the background step's ability increases.
type backgroundForm struct {
	Mode  string `form:"mode"  validate:"omitempty,oneof=two three"`
	Plus2 string `form:"plus2" validate:"omitempty,oneof=str dex con int wis cha"`
	Plus1 string `form:"plus1" validate:"omitempty,oneof=str dex con int wis cha"`
}

// abilitiesForm is the abilities step's base scores.
type abilitiesForm struct {
	Str int `form:"str" validate:"min=3,max=20"`
	Dex int `form:"dex" validate:"min=3,max=20"`
	Con int `form:"con" validate:"min=3,max=20"`
	Int int `form:"int" validate:"min=3,max=20"`
	Wis int `form:"wis" validate:"min=3,max=20"`
	Cha int `form:"cha" validate:"min=3,max=20"`
}

// detailsForm is the details step; the limits are the sheet's.
type detailsForm struct {
	Name      string `form:"name"      validate:"max=100"`
	Backstory string `form:"backstory" validate:"max=20000"`
	Vision    string `form:"vision"    validate:"max=200"`
	Size      string `form:"size"      validate:"max=20"`
}

// originFeat is the background's feat from the compendium, if it is there.
// The second result is the feat's source key ("" when it has no entry).
func (a *App) originFeat(ctx context.Context, d *compendium.TextData) (*dnd5e.Feature, string, error) {
	for _, p := range d.Parts {
		if p.Type != compendium.PartFeat {
			continue
		}
		name, _, _ := strings.Cut(p.Desc, " (")
		td, source, err := a.textEntrySource(ctx, enum.KindFeat, compendium.Slugify(name), "")
		if err != nil {
			return nil, "", err
		}
		f := &dnd5e.Feature{Name: strings.TrimSpace(p.Desc), Source: dnd5e.SourceFeat}
		if td != nil {
			f.Description = textMarkdown(td)
		}
		return f, source, nil
	}
	return nil, "", nil
}

func applyBackground(s *dnd5e.Sheet, name string, d *compendium.TextData, feat *dnd5e.Feature) {
	s.Background = name
	for _, p := range d.Parts {
		switch p.Type {
		case compendium.PartSkills:
			s.AddProfs(name, skillKeys(compendium.SplitList(p.Desc)))
		case compendium.PartTools:
			s.AddProfs(name, nil, p.Desc)
		}
	}
	if feat != nil {
		addFeat(s, *feat)
	}
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
			d, err := compendium.DecodeData(enum.KindItem, e.DataJson)
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
				g.armor = &dnd5e.WornArmor{Slug: e.Slug, Name: e.Name, Source: e.SourceKey, Type: ar.Type, Base: ar.Base, AddDex: ar.AddDex, DexCap: ar.DexCap, StrReq: ar.StrReq, Stealth: ar.Stealth}
			}
		}
	}
	return g, nil
}

// findItem looks an equipment line up by name, singular too ("Daggers",
// "Pouches").
func (a *App) findItem(ctx context.Context, name string) (sqlc.CompendiumEntry, error) {
	slug := compendium.Slugify(name)
	e, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: enum.KindItem, Slug: slug})
	for _, suffix := range []string{"s", "es"} {
		if !errors.Is(err, sql.ErrNoRows) || !strings.HasSuffix(slug, suffix) {
			break
		}
		e, err = a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: enum.KindItem, Slug: strings.TrimSuffix(slug, suffix)})
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
