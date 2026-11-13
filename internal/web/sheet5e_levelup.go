package web

import (
	"cmp"
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
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// classInfo converts a compendium class for the level-up helper.
func classInfo(row sqlc.CompendiumEntry) (dnd5e.ClassInfo, error) {
	cd, err := compendium.Decode[compendium.ClassData](enum.KindClass, row.DataJson)
	if err != nil {
		return dnd5e.ClassInfo{}, err
	}
	info := dnd5e.ClassInfo{Slug: row.Slug, Name: row.Name, Source: row.SourceKey, HitDie: cd.HitDie, Caster: cd.Caster, Saves: cd.Saves, Table: cd.Table, Features: classFeatures(cd)}
	if info.HitDie == 0 {
		info.HitDie = 8 // homebrew classes without a hit die
	}
	return info, nil
}

func classFeatures(cd *compendium.ClassData) []dnd5e.ClassFeature {
	out := make([]dnd5e.ClassFeature, 0, len(cd.Features))
	for _, f := range cd.Features {
		out = append(out, dnd5e.ClassFeature{Name: f.Name, Description: f.Desc, Levels: f.Levels, Subclass: f.Subclass})
	}
	return out
}

// levelUpView previews a level in classSlug; with pickDefault it picks the
// class the character would most likely level instead.
func (a *App) levelUpView(ctx context.Context, ch sqlc.Character, classSlug string, pickDefault bool) (views.LevelUpView, error) {
	sheet, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		return views.LevelUpView{}, err
	}
	if pickDefault {
		classSlug = sheet.DefaultClass()
	}
	v := views.LevelUpView{CharacterID: ch.ID, Name: ch.Name, Level: sheet.Level(), Selected: classSlug, Sheet: sheet}
	classes, err := a.Store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: enum.KindClass, CampaignID: ch.CampaignID})
	if err != nil {
		return v, err
	}
	for _, cl := range classes {
		v.Classes = append(v.Classes, views.Option{Slug: cl.Slug, Name: cl.Name})
	}
	if classSlug == "" {
		return v, nil
	}
	info, err := a.levelUpClass(ctx, sheet, classSlug, ch.CampaignID)
	if err != nil {
		return v, err
	}
	if err := a.withSubclass(ctx, &info, sheet, "", 0); err != nil {
		return v, err
	}
	offer, err := a.levelUpOffer(ctx, sheet, info, ch.CampaignID)
	if err != nil {
		return v, err
	}
	if offer.planErr != nil {
		v.Error = offer.planErr.Error()
		return v, nil
	}
	v.Plan = &offer.plan
	subs, err := a.Store.ListSubclassesOf(ctx, sqlc.ListSubclassesOfParams{ClassSlug: classSlug, CampaignID: ch.CampaignID})
	if err != nil {
		return v, err
	}
	v.Subclasses = subclassOptions(subs)
	if v.SubclassPanels, err = subclassPanels(subs); err != nil {
		return v, err
	}
	feats, err := a.Store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: enum.KindFeat, CampaignID: ch.CampaignID})
	if err != nil {
		return v, err
	}
	v.FeatChoices = map[string][]views.ChoiceGroup{}
	for _, f := range feats {
		if !offer.plan.ASI || f.Name == dnd5e.ASIFeature {
			continue // the form's ability selects are the ASI feat
		}
		td, err := a.textEntry(ctx, enum.KindFeat, f.Slug)
		if err != nil {
			return v, err
		}
		if td == nil || !dnd5e.MeetsPrerequisite(td.Prerequisite, sheet, offer.plan.NewLevel, nil, compendium.Slugify) {
			continue
		}
		v.Feats = append(v.Feats, views.Option{Slug: f.Slug, Name: f.Name})
		groups, err := a.featChoices(ctx, "feat_"+f.Slug, f.Name, td, "", sheet, ch.CampaignID)
		if err != nil {
			return v, err
		}
		if len(groups) > 0 {
			v.FeatChoices[f.Slug] = groupViews(groups)
		}
	}
	v.LineageSpells, v.Groups, v.Spells, v.SpellNeed = offer.lineage, groupViews(offer.groups), offer.spells, offer.need
	return v, nil
}

// levelUpClass is classSlug's class for the sheet: its pinned copy, or for
// a new class one campaignID may see (the slug comes from the request).
func (a *App) levelUpClass(ctx context.Context, sheet dnd5e.Sheet, classSlug string, campaignID int64) (dnd5e.ClassInfo, error) {
	if classPin(sheet, classSlug) == "" {
		if _, err := a.entryBySlug(ctx, enum.KindClass, classSlug, campaignID); err != nil {
			return dnd5e.ClassInfo{}, err
		}
	}
	info, _, err := classEntry(ctx, a.Store.Queries, classSlug, classPin(sheet, classSlug))
	return info, err
}

// levelUpOffer is what a level in a class offers besides the form's own
// fields. The page and the save both build it, so they offer the same.
type levelUpOffer struct {
	plan dnd5e.LevelUpPlan
	// planErr refuses the level (level 20): the page says so, LevelUp refuses.
	planErr error
	// lineage are the species lineage's spells of the new level; the class
	// list does not offer them.
	lineage []dnd5e.Spell
	spells  []views.Option // class spells to pick from; none on a planErr
	need    views.SpellNeed
	// groups are class options, weapon masteries and feat spell changes.
	groups []choiceGroup
}

// levelUpOffer builds the offer of a level in info (with the subclass it
// would have) for sheet.
func (a *App) levelUpOffer(ctx context.Context, sheet dnd5e.Sheet, info dnd5e.ClassInfo, campaignID int64) (levelUpOffer, error) {
	var o levelUpOffer
	var err error
	o.plan, o.planErr = sheet.Plan(info)
	if o.lineage, err = a.lineageSpells(ctx, sheet, o.plan.NewLevel, campaignID); err != nil {
		return o, err
	}
	if o.planErr == nil {
		if o.spells, err = a.classSpellOptions(ctx, info.Slug, o.plan.Slots, o.plan.PactLevel, sheet, o.lineage, campaignID); err != nil {
			return o, err
		}
		o.need = spellNeed(info, o.plan.ClassLevel, sheet, o.spells)
	}
	changes, err := a.featSpellChanges(ctx, sheet, o.lineage, campaignID)
	if err != nil {
		return o, err
	}
	cantrips, err := a.optionCantrips(ctx, sheet, info, o.plan.ClassLevel, campaignID)
	if err != nil {
		return o, err
	}
	mastery, err := a.masteryGroups(ctx, sheet, info, o.plan.ClassLevel, campaignID)
	if err != nil {
		return o, err
	}
	o.groups = slices.Concat(classOptionGroups(info, o.plan.ClassLevel, o.plan.ClassLevel-1, sheet, cantrips), mastery, changes)
	return o, nil
}

// classSpellOptions are the class's spells the slots (and pact slots) can
// cast that the sheet does not have yet, nor learns in this step (taught:
// lineage spells, which AddSpell would keep over the class pick).
func (a *App) classSpellOptions(ctx context.Context, classSlug string, slots [9]int, pactLevel int, sheet dnd5e.Sheet, taught []dnd5e.Spell, campaignID int64) ([]views.Option, error) {
	maxLevel := pactLevel
	for i, n := range slots {
		if n > 0 {
			maxLevel = max(maxLevel, i+1)
		}
	}
	spells, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: classSlug, MaxLevel: int64(maxLevel), CampaignID: campaignID})
	if err != nil {
		return nil, err
	}
	known := slices.Concat(sheet.Spells, taught)
	var out []views.Option
	for _, s := range spells {
		if !slices.ContainsFunc(known, func(have dnd5e.Spell) bool { return have.Slug == s.Slug }) {
			out = append(out, views.Option{Slug: s.Slug, Name: s.Name, Level: int(s.Level)})
		}
	}
	return out, nil
}

// spellNeed is what the class table's cantrip and spell totals (at the class
// level being reached) leave to pick: each total less the class's spells on
// the sheet, at most what is offered. A spellbook's size replaces the spell
// total, exactly: its list is prepared from the book. Lineage and feat spells
// (their own ability or feat) do not count; untagged spells of older sheets
// count while the sheet has this class only.
func spellNeed(info dnd5e.ClassInfo, level int, sheet dnd5e.Sheet, offered []views.Option) views.SpellNeed {
	class, book := info.Slug, dnd5e.Spellbook(info, level)
	has := slices.ContainsFunc(sheet.Classes, func(cl dnd5e.ClassLevel) bool { return cl.Slug == class })
	only := has && len(sheet.Classes) == 1
	var have, avail [2]int // cantrips, level 1+
	unsure := false        // untagged spells that may be this class's
	for _, sp := range sheet.Spells {
		if sp.Feat != "" || sp.Ability != "" {
			continue
		}
		if sp.Class == class || sp.Class == "" && only {
			have[min(sp.Level, 1)]++
		}
		unsure = unsure || sp.Class == "" && has && !only
	}
	for _, o := range offered {
		avail[min(o.Level, 1)]++
	}
	// Without knowing which class taught them, the totals can only be upper limits.
	need := views.SpellNeed{Cantrips: -1, Spells: -1, CantripsUpTo: unsure, SpellsUpTo: unsure}
	cantrips, spells, prepared := dnd5e.SpellTotals(info, level)
	if cantrips >= 0 {
		need.Cantrips = min(max(cantrips-have[0], 0), avail[0])
	}
	if spells >= 0 {
		need.Spells = min(max(spells-have[1], 0), avail[1])
		need.SpellsUpTo = unsure || prepared
	}
	if book > 0 {
		need.Spells, need.SpellsUpTo, need.Book = min(max(book-have[1], 0), avail[1]), unsure, true
		// Only a book known to be short earlier gets the hint, never a new one (creation, a new class).
		need.BookShort = level > 1 && !unsure && have[1] < dnd5e.Spellbook(info, level-1)
		need.BookHave, need.BookSize = have[1], book
	}
	return need
}

// classSpellPicks are the spells ticked in the form for class, refused
// unless offered and as many as need asks.
func (a *App) classSpellPicks(c *gin.Context, class string, offered []views.Option, need views.SpellNeed, campaignID int64) ([]dnd5e.Spell, error) {
	var spells []dnd5e.Spell
	var got [2]int
	for _, slug := range c.PostFormArray("spell") {
		if !slices.ContainsFunc(offered, func(o views.Option) bool { return o.Slug == slug }) || slices.ContainsFunc(spells, func(sp dnd5e.Spell) bool { return sp.Slug == slug }) {
			return nil, fieldError{userError{errors.New("choose spells from the list")}, views.FieldSpell}
		}
		sp, err := a.compendiumSpell(c.Request.Context(), slug, campaignID)
		if err != nil {
			return nil, err
		}
		sp.Class = class
		spells = append(spells, sp)
		got[min(sp.Level, 1)]++
	}
	off := func(n, got int, upTo bool) bool { return n >= 0 && (got > n || !upTo && got < n) }
	if off(need.Cantrips, got[0], need.CantripsUpTo) || off(need.Spells, got[1], need.SpellsUpTo) {
		return nil, fieldError{userError{fmt.Errorf("choose %s: you ticked %s and %s", cmp.Or(need.String(), "no spells here"), views.Plural(got[0], "cantrip"), views.Plural(got[1], "level 1+ spell"))}, views.FieldSpell}
	}
	return spells, nil
}

// addClassSpells adds the class picks after the lineage's and feats' spells
// of the same step, refusing one those taught already: it would not count as
// a class pick.
func addClassSpells(s *dnd5e.Sheet, spells []dnd5e.Spell) error {
	for _, sp := range spells {
		if slices.ContainsFunc(s.Spells, func(have dnd5e.Spell) bool { return have.Slug == sp.Slug }) {
			return fieldError{userError{fmt.Errorf("%s comes from your species or feat already: choose another spell", sp.Name)}, views.FieldSpell}
		}
		s.AddSpell(sp)
	}
	return nil
}

func (a *App) levelUpForm(c *gin.Context) error {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return nil
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player levels up their character.")
		return nil
	}
	ctx := c.Request.Context()
	slug, picked := c.GetQuery("class")
	v, err := a.levelUpView(ctx, ch, slug, !picked)
	if errors.Is(err, sql.ErrNoRows) {
		v.Error = "That class is not in the compendium."
	} else if err != nil {
		return err
	}
	camp, err := a.Store.GetCampaign(ctx, ch.CampaignID)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.LevelUpPage(a.layout(w, "Level up", views.NavCharacter, &camp, false), v))
	return nil
}

func (a *App) levelUp(c *gin.Context) error {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return nil
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player levels up their character.")
		return nil
	}
	ctx := c.Request.Context()
	sheet, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		return err
	}
	info, err := a.levelUpClass(ctx, sheet, c.PostForm("class"), ch.CampaignID)
	if err != nil {
		return err
	}
	choice := dnd5e.LevelUpChoice{Class: info, Subclass: strings.TrimSpace(c.PostForm("subclass"))}
	// The form's level; a second submit (a double click, another device) finds it taken.
	level, checkLevel := formInt(c, "level")
	if err := a.withSubclass(ctx, &choice.Class, sheet, choice.Subclass, ch.CampaignID); err != nil {
		return err
	}
	if choice.Subclass != "" && choice.Class.SubclassName == "" {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a subclass from the list.") // gone since the form loaded
		return nil
	}
	choice.Subclass = choice.Class.SubclassName // the sheet keeps the name, not the picked slug
	if roll, ok := formInt(c, "hp_roll"); ok {
		choice.HPRoll = int(roll)
	}
	for _, f := range []string{"asi1", "asi2"} {
		if v := c.PostForm(f); v != "" {
			choice.ASI = append(choice.ASI, v)
		}
	}
	var featErr error
	var featSource string
	if slug := c.PostForm("feat"); slug != "" {
		feat, err := a.entryBySlug(ctx, enum.KindFeat, slug, ch.CampaignID)
		switch {
		case errors.Is(err, sql.ErrNoRows): // crafted, another campaign's, or a homebrew feat deleted since the form loaded
			featErr = fieldError{userError{errors.New("choose a feat from the list")}, views.FieldASI}
		case err != nil:
			return err
		default:
			choice.Feat, choice.FeatDesc = feat.Name, feat.BodyText
			featSource = feat.SourceKey
		}
	}
	// A plan error (level 20) comes back from LevelUp below.
	offer, err := a.levelUpOffer(ctx, sheet, choice.Class, ch.CampaignID)
	if err != nil {
		return err
	}
	plan, planErr, lineage := offer.plan, offer.planErr, offer.lineage
	var spells []dnd5e.Spell
	var spellErr error
	var ue userError
	if planErr == nil { // else LevelUp refuses (level 20) and plan is empty
		// A refusal waits for readPicks, so the page re-renders as for any other.
		spells, spellErr = a.classSpellPicks(c, info.Slug, offer.spells, offer.need, ch.CampaignID)
		if spellErr != nil && !errors.As(spellErr, &ue) {
			return spellErr
		}
	}
	// The class picks go in after the feats' (addClassSpells), so a clash is refused.
	choice.Spells = lineage

	// The cantrips the sheet will know, for checkNewOptions.
	learned, err := a.cantripData(ctx, slices.Concat(sheet.Spells, lineage, spells), ch.CampaignID)
	if err != nil {
		return err
	}
	// Spell Change and the feat's own choices apply after LevelUp has added the feat.
	groups := offer.groups
	notFeat := len(groups)
	if choice.Feat != "" {
		slug := c.PostForm("feat")
		td, err := a.textEntry(ctx, enum.KindFeat, slug)
		if err != nil {
			return err
		}
		if planErr == nil && !dnd5e.MeetsPrerequisite(td.Prerequisite, sheet, plan.NewLevel, nil, compendium.Slugify) {
			featErr = fieldError{userError{fmt.Errorf("%s needs %s", choice.Feat, td.Prerequisite)}, views.FieldASI}
		}
		fg, err := a.featChoices(ctx, "feat_"+slug, choice.Feat, td, "", sheet, ch.CampaignID)
		if err != nil {
			return err
		}
		groups = append(groups, fg...)
	}
	read := groups
	if featErr != nil {
		read = groups[:notFeat] // the feat's own choices are moot: it may not be taken at all
	}
	picks, err := readPicks(c, read)
	err = errors.Join(featErr, err, spellErr)
	if err == nil {
		err = a.updateCharacter(ctx, ch.ID, func(q *sqlc.Queries, dc *dnd5e.Character) error {
			before, err := dc.Sheet.JSON()
			if err != nil {
				return err
			}
			if checkLevel && int64(dc.Sheet.Level()+1) != level {
				return userError{fmt.Errorf("level %d was already taken", level)}
			}
			maxHP := dc.MaxHP
			if err := dc.LevelUp(choice); err != nil {
				var ae dnd5e.ASIError
				if errors.As(err, &ae) {
					return fieldError{userError{err}, views.FieldASI}
				}
				return userError{err}
			}
			if choice.Feat != "" {
				pinFeat(&dc.Sheet, choice.Feat, featSource)
			}
			applyPicks(&dc.Sheet, groups, picks)
			if err := addClassSpells(&dc.Sheet, spells); err != nil {
				return err
			}
			if err := checkNewOptions(choice.Class, sheet, dc.Sheet, plan.ClassLevel, learned); err != nil {
				return err
			}
			for _, sp := range lineage {
				dc.Sheet.PutFeature(dnd5e.FreeCast(sp, sheet.Species), nil)
			}
			return q.CreateLevelUp(ctx, sqlc.CreateLevelUpParams{CharacterID: ch.ID, MaxHp: int64(maxHP), SheetJson: before})
		})
	}
	if errors.As(err, &ue) {
		v, verr := a.levelUpView(ctx, ch, info.Slug, false)
		if verr != nil {
			return verr
		}
		v.Error, v.Errors = formErrors(err, v.DrawsField)
		v.Form = c.Request.PostForm
		camp, _ := a.Store.GetCampaign(ctx, ch.CampaignID)
		a.render(c, http.StatusUnprocessableEntity, views.LevelUpPage(a.layout(w, "Level up", views.NavCharacter, &camp, false), v))
		return nil
	}
	if err != nil {
		return err
	}
	a.Hub.Publish(ch.CampaignID, hub.Topic{Kind: hub.TopicCharacter, ID: ch.ID})
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/characters/%d", ch.ID))
	return nil
}

// revertLevelUp puts back the sheet and max HP from before the latest
// level-up. Current HP is kept, up to that max.
func (a *App) revertLevelUp(c *gin.Context) error {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return nil
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player reverts a level-up.")
		return nil
	}
	ctx := c.Request.Context()
	err := a.updateCharacter(ctx, ch.ID, func(q *sqlc.Queries, dc *dnd5e.Character) error {
		snap, err := q.LatestLevelUp(ctx, ch.ID)
		if errors.Is(err, sql.ErrNoRows) {
			return userError{errors.New("there is no level-up to revert")}
		}
		if err != nil {
			return err
		}
		if dc.Sheet, err = dnd5e.ParseSheet(snap.SheetJson); err != nil {
			return err
		}
		dc.MaxHP = int(snap.MaxHp)
		dc.HP = min(dc.HP, dc.MaxHP)
		return q.DeleteLevelUp(ctx, snap.ID)
	})
	if err != nil {
		return err
	}
	a.characterChanged(c, ch)
	return nil
}
