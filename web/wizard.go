package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

// wizardSteps are the creation wizard's steps, in order. Class, species,
// background and equipment add to the sheet, so each runs once; afterwards
// the sheet is changed directly.
var wizardSteps = []views.WizardStep{
	{Key: "class", Title: "Class"},
	{Key: "species", Title: "Species"},
	{Key: "abilities", Title: "Ability scores"},
	{Key: "background", Title: "Background"},
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
		if step == "equipment" {
			v.Equipment, err = a.equipmentChoices(ctx, s.C.Sheet)
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

// wizardSave applies a step and moves to the next.
func (a *App) wizardSave(c *gin.Context) {
	w, ch, step, ok := a.wizardAccess(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	cmd, err := a.wizardCommand(c, ch, step)
	if err == nil && cmd != nil {
		err = a.runCommand(ctx, ch.ID, *cmd)
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
		backstory := c.PostForm("backstory")
		cmd := dnd5e.NewCommand("create details", perm.ScopeSheet, func(cc *dnd5e.Character) error {
			cc.Name, cc.Sheet.Backstory = name, backstory
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

// applySpecies sets the species, its speed, and its traits as features
// with what the sheet can automate.
func applySpecies(s *dnd5e.Sheet, name string, d *compendium.TextData) {
	s.Species = name
	for _, p := range d.Parts {
		switch {
		case p.Type == "speed" || p.Name == "Speed":
			if m := reFeet.FindStringSubmatch(p.Desc); m != nil {
				s.Speed, _ = strconv.Atoi(m[1])
			}
		case p.Type == "size" || p.Name == "Size":
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
		f := *feat
		mods := dnd5e.ParseEffects(f.Name, f.Description, nil, "").Apply(&f)
		s.PutFeature(f, mods)
		s.Feats = append(s.Feats, f.Name)
	}
	return nil
}

// gear is what starting equipment items are in the compendium: weapons to
// put under attacks, armor to wear.
type gear struct {
	weapons []dnd5e.Weapon
	armor   *dnd5e.WornArmor
	shield  bool
}

func (a *App) gearFor(ctx context.Context, picked []compendium.EquipmentOption) (gear, error) {
	var g gear
	for _, o := range picked {
		for _, it := range o.Items {
			e, err := a.findItem(ctx, it.Name)
			if errors.Is(err, sql.ErrNoRows) {
				continue
			}
			if err != nil {
				return g, err
			}
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

// findItem looks an equipment line up by name, singular too ("Daggers").
func (a *App) findItem(ctx context.Context, name string) (sqlc.CompendiumEntry, error) {
	slug := compendium.Slugify(name)
	e, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindItem), Slug: slug})
	if errors.Is(err, sql.ErrNoRows) && strings.HasSuffix(slug, "s") {
		return a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindItem), Slug: strings.TrimSuffix(slug, "s")})
	}
	return e, err
}

func applyEquipment(s *dnd5e.Sheet, picked []compendium.EquipmentOption, g gear) {
	for _, o := range picked {
		s.Coins.GP += o.GP
		for _, it := range o.Items {
			if i := slices.IndexFunc(s.Inventory, func(x dnd5e.Item) bool { return strings.EqualFold(x.Name, it.Name) }); i >= 0 {
				s.Inventory[i].Qty += it.Qty
			} else {
				s.Inventory = append(s.Inventory, dnd5e.Item{Name: it.Name, Qty: it.Qty})
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
