package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"

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

// wizardStep is how a step saves.
type wizardStep struct {
	// needs are the earlier picks the step asks about: saved without them, it
	// could never ask for what they bring (a background feat's spells, a class's gear).
	needs []string
	// command is the change the step's form asks for; nil when the step
	// changes nothing (it was applied already, or skipped).
	command func(a *App, c *gin.Context, ch sqlc.Character) (*dnd5e.Command, error)
}

var wizardStepSaves = map[string]wizardStep{
	views.StepClass:      {command: (*App).classStep},
	views.StepBackground: {command: (*App).backgroundStep},
	views.StepSpecies:    {command: (*App).speciesStep},
	views.StepAbilities:  {needs: []string{views.StepBackground}, command: (*App).abilitiesStep},
	views.StepChoices:    {needs: []string{views.StepClass, views.StepBackground, views.StepSpecies}, command: (*App).choicesStep},
	views.StepEquipment:  {needs: []string{views.StepClass, views.StepBackground}, command: (*App).equipmentStep},
	views.StepDetails:    {command: (*App).detailsStep},
	views.StepReview:     {command: (*App).reviewStep},
}

// missingSteps are the steps step needs that the sheet lacks. A background
// is needed only when the compendium has one to pick.
func (a *App) missingSteps(ctx context.Context, sh dnd5e.Sheet, step string, campaignID int64) ([]views.WizardStep, error) {
	has := map[string]bool{views.StepClass: len(sh.Classes) > 0, views.StepSpecies: sh.Species != "", views.StepBackground: sh.Background != ""}
	needs := wizardStepSaves[step].needs
	if !has[views.StepBackground] && slices.Contains(needs, views.StepBackground) {
		names, err := a.Store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: enum.KindBackground, CampaignID: campaignID})
		if err != nil {
			return nil, err
		}
		has[views.StepBackground] = len(names) == 0
	}
	var out []views.WizardStep
	for _, s := range wizardSteps {
		if slices.Contains(needs, s.Key) && !has[s.Key] {
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
	c, err := a.chosenEntry(ctx, enum.KindBackground, sh.Background, sh.BackgroundSource)
	if c == nil || len(c.Abilities) != 3 {
		return nil, err
	}
	return c, nil
}

// chosenEntry is a class, species or background on the sheet as the wizard
// shows it; nil when the compendium lacks it.
func (a *App) chosenEntry(ctx context.Context, kind enum.CompendiumKind, name, pin string) (*views.WizardChoice, error) {
	if name == "" {
		return nil, nil
	}
	row, _, err := entryFor(ctx, a.Store.Queries, kind, compendium.Slugify(name), pin)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	data, err := compendium.DecodeData(kind, row.DataJson)
	if err != nil {
		return nil, err
	}
	return wizardChoice(row, data), nil
}

// doneChoice is what an applied class, species or background step gave: the
// wizard's class is the sheet's first.
func (a *App) doneChoice(ctx context.Context, sh dnd5e.Sheet, kind enum.CompendiumKind) (*views.WizardChoice, error) {
	switch kind {
	case enum.KindClass:
		if len(sh.Classes) == 0 {
			return nil, nil
		}
		return a.chosenEntry(ctx, kind, sh.Classes[0].Slug, sh.Classes[0].Source)
	case enum.KindSpecies:
		return a.chosenEntry(ctx, kind, sh.Species, sh.SpeciesSource)
	}
	return a.chosenEntry(ctx, kind, sh.Background, sh.BackgroundSource)
}

// wizardChange redoes a step: the character goes back to how it was before
// it, which also undoes every step saved after it.
func (a *App) wizardChange(c *gin.Context) error {
	_, ch, step, ok := a.wizardAccess(c)
	if !ok {
		return nil
	}
	if step != c.Param("step") { // wizardAccess falls back to the first step
		a.fail(c, http.StatusNotFound, "There is no such step.")
		return nil
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
	if err != nil {
		return err
	}
	topics, err := a.characterTopics(ctx, ch.ID)
	if err != nil {
		return err
	}
	a.Hub.Publish(ch.CampaignID, topics...)
	redirect(c, wizardPath(ch.ID, step))
	return nil
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

func (a *App) wizardPage(c *gin.Context) error {
	w, ch, step, ok := a.wizardAccess(c)
	if !ok {
		return nil
	}
	v, err := a.wizardView(c.Request.Context(), w, ch, step, c.Query("pick"), c.Query("variant"))
	if err != nil {
		return err
	}
	return a.renderWizard(c, http.StatusOK, w, ch, v)
}

func (a *App) renderWizard(c *gin.Context, status int, w who, ch sqlc.Character, v views.WizardView) error {
	camp, err := a.Store.GetCampaign(c.Request.Context(), ch.CampaignID)
	if err != nil {
		return err
	}
	a.render(c, status, views.WizardPage(a.layout(w, "Create "+ch.Name, views.NavCharacter, &camp, false), v))
	return nil
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
	if s.C.Sheet.SetUp(step) {
		v.Choice, err = a.doneChoice(ctx, s.C.Sheet, kind)
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
				ch.Abilities = dnd5e.AbilityKeys(p.Desc)
			case compendium.PartSkills:
				ch.Skills = dnd5e.SkillKeys(dnd5e.SplitList(p.Desc))
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

// wizardSave applies a step and moves to the next.
func (a *App) wizardSave(c *gin.Context) error {
	w, ch, step, ok := a.wizardAccess(c)
	if !ok {
		return nil
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
			return verr
		}
		v.Error, v.Errors = formErrors(err, v.DrawsField)
		v.Form = c.Request.PostForm
		return a.renderWizard(c, http.StatusUnprocessableEntity, w, ch, v)
	}
	if err != nil {
		return err
	}
	topics, err := a.characterTopics(ctx, ch.ID)
	if err != nil {
		return err
	}
	a.Hub.Publish(ch.CampaignID, topics...)
	c.Redirect(http.StatusSeeOther, nextStep(ch.ID, step))
	return nil
}

// wizardCommand is the change a step's form asks for; nil when the step
// changes nothing (it was applied already, or skipped).
func (a *App) wizardCommand(c *gin.Context, ch sqlc.Character, step string) (*dnd5e.Command, error) {
	ctx := c.Request.Context()
	ws := wizardStepSaves[step]
	if ws.needs != nil {
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil {
			return nil, err
		}
		if missing, err := a.missingSteps(ctx, sheet, step, ch.CampaignID); err != nil || len(missing) > 0 {
			return nil, err // saved now, the step could not ask for what they bring later
		}
	}
	if ws.command == nil {
		return nil, nil
	}
	return ws.command(a, c, ch)
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
