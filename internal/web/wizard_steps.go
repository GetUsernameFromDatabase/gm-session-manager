package web

import (
	"cmp"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
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

// setupOnce is a creation step's command: it applies the step once and
// snapshots the sheet before it, so the step can be redone (wizardChange).
func setupOnce(step string, apply func(s *dnd5e.Sheet, cc *dnd5e.Character) error) *dnd5e.Command {
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

func (a *App) classStep(c *gin.Context, ch sqlc.Character) (*dnd5e.Command, error) {
	pick := c.PostForm("pick")
	if pick == "" {
		return nil, nil
	}
	sheet, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil || sheet.SetUp(views.StepClass) {
		return nil, err // chosen already: the skills would be ignored
	}
	row, err := a.wizardPick(c.Request.Context(), enum.KindClass, pick, ch.CampaignID)
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
	return setupOnce(views.StepClass, func(s *dnd5e.Sheet, cc *dnd5e.Character) error {
		if s.Level() > 0 {
			return nil
		}
		if err := cc.LevelUp(dnd5e.LevelUpChoice{Class: info}); err != nil {
			return err
		}
		cc.Sheet.AddProfs(info.Name, skills, trainedIn(choice.Traits)...)
		return nil
	}), nil
}

// textPick is the posted species or background (a step and kind both), nil
// without one.
func (a *App) textPick(c *gin.Context, ch sqlc.Character, kind enum.CompendiumKind) (sqlc.CompendiumEntry, *compendium.TextData, error) {
	pick := c.PostForm("pick")
	if pick == "" {
		return sqlc.CompendiumEntry{}, nil, nil
	}
	row, err := a.wizardPick(c.Request.Context(), kind, pick, ch.CampaignID)
	if err != nil {
		return row, nil, err
	}
	data, err := compendium.Decode[compendium.TextData](kind, row.DataJson)
	if err != nil {
		return row, nil, err
	}
	return row, data, nil
}

func (a *App) speciesStep(c *gin.Context, ch sqlc.Character) (*dnd5e.Command, error) {
	row, td, err := a.textPick(c, ch, enum.KindSpecies)
	if err != nil || td == nil {
		return nil, err
	}
	return setupOnce(views.StepSpecies, func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
		applySpecies(s, row.Name, td)
		s.SpeciesSource = row.SourceKey
		return nil
	}), nil
}

func (a *App) backgroundStep(c *gin.Context, ch sqlc.Character) (*dnd5e.Command, error) {
	ctx := c.Request.Context()
	row, td, err := a.textPick(c, ch, enum.KindBackground)
	if err != nil || td == nil {
		return nil, err
	}
	feat, featSource, err := a.originFeat(ctx, td)
	if err != nil {
		return nil, err
	}
	sheet, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		return nil, err
	}
	pickAgain, err := a.backgroundSkillPick(ctx, sheet, wizardChoice(row, td).Skills)
	if err != nil {
		return nil, err
	}
	skills := postedSkills(c)
	if err := checkSkillPick(pickAgain, skills); err != nil {
		return nil, err
	}
	return setupOnce(views.StepBackground, func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
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
}

func (a *App) abilitiesStep(c *gin.Context, ch sqlc.Character) (*dnd5e.Command, error) {
	ctx := c.Request.Context()
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
		if scores[ab]+raises[ab] > dnd5e.AbilityCap {
			errs = append(errs, fieldError{userError{fmt.Errorf("%s cannot go above 20", dnd5e.AbilityNames[ab])}, ab})
		}
	}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}
	return setupOnce(views.StepAbilities, func(s *dnd5e.Sheet, cc *dnd5e.Character) error {
		before := s.Mod("con")
		for ab, n := range scores {
			s.Scores[ab] = n + raises[ab]
		}
		// Hit points already gained follow the new Constitution.
		delta := (s.Mod("con") - before) * s.Level()
		cc.MaxHP, cc.HP = max(cc.MaxHP+delta, 0), max(cc.HP+delta, 0)
		return nil
	}), nil
}

func (a *App) choicesStep(c *gin.Context, ch sqlc.Character) (*dnd5e.Command, error) {
	ctx := c.Request.Context()
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
	return setupOnce(views.StepChoices, func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
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
}

func (a *App) equipmentStep(c *gin.Context, ch sqlc.Character) (*dnd5e.Command, error) {
	ctx := c.Request.Context()
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
	return setupOnce(views.StepEquipment, func(s *dnd5e.Sheet, _ *dnd5e.Character) error {
		applyEquipment(s, picked, gear)
		return nil
	}), nil
}

func (a *App) detailsStep(c *gin.Context, _ sqlc.Character) (*dnd5e.Command, error) {
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
}

func (a *App) reviewStep(*gin.Context, sqlc.Character) (*dnd5e.Command, error) {
	cmd := dnd5e.NewCommand("create review", perm.ScopeSheet, func(cc *dnd5e.Character) error {
		if !cc.Sheet.SetUp(views.StepReview) {
			cc.Sheet.SetupDone = append(cc.Sheet.SetupDone, views.StepReview)
		}
		cc.Sheet.SetupSnapshots = nil // finished: steps change on the sheet from now on
		return nil
	})
	return &cmd, nil
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
			s.AddProfs(name, dnd5e.SkillKeys(dnd5e.SplitList(p.Desc)))
		case compendium.PartTools:
			s.AddProfs(name, nil, p.Desc)
		}
	}
	if feat != nil {
		dnd5e.AddFeat(s, *feat)
	}
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
			item, err := compendium.Decode[compendium.ItemData](enum.KindItem, e.DataJson)
			if err != nil {
				return g, err
			}
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
