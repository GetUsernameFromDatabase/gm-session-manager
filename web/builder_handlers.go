package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

// maxPicks caps the builder's look-up results and suggestion lists.
const maxPicks = 15

// builderExtra is what a builder form takes from the compendium: a weapon
// or feature to prefill (?entry=, plus ?feature= for a class feature) and
// the names a details field suggests.
func (a *App) builderExtra(c *gin.Context) (views.BuilderExtra, error) {
	ctx := c.Request.Context()
	var x views.BuilderExtra
	switch c.Query("op") {
	case "set-species", "set-background":
		kind := strings.TrimPrefix(c.Query("op"), "set-")
		rows, err := a.Store.ListCompendiumNames(ctx, kind)
		for _, r := range rows {
			x.Suggestions = append(x.Suggestions, r.Name)
		}
		return x, err
	}
	id, err := strconv.ParseInt(c.Query("entry"), 10, 64)
	if err != nil {
		return x, nil
	}
	e, err := a.Store.GetCompendiumEntry(ctx, id)
	if err != nil {
		return x, err
	}
	data, err := compendium.DecodeData(compendium.Kind(e.Kind), e.DataJson)
	if err != nil {
		return x, err
	}
	switch d := data.(type) {
	case *compendium.ItemData:
		w := compendiumWeapon(e.Name, d)
		x.Weapon = &w
	case *compendium.ClassData:
		if i := slices.IndexFunc(d.Features, func(f compendium.ClassFeature) bool { return f.Name == c.Query("feature") }); i >= 0 {
			cf := d.Features[i]
			f := dnd5e.Feature{Name: cf.Name, Description: cf.Desc, Source: e.Name}
			x.FeatureMods = dnd5e.ParseEffects(cf.Name, cf.Desc, d.Table, e.Name).Apply(&f)
			x.Feature = &f
		}
	case *compendium.TextData:
		f := dnd5e.Feature{Name: e.Name, Description: textMarkdown(d), Source: kindSources[e.Kind]}
		x.FeatureMods = dnd5e.ParseEffects(e.Name, f.Description, nil, "").Apply(&f)
		x.Feature = &f
	}
	return x, nil
}

// kindSources name where a feature looked up in the compendium comes from.
var kindSources = map[string]string{string(compendium.KindFeat): "Feat", string(compendium.KindSpecies): "Species", string(compendium.KindBackground): "Background"}

// compendiumWeapon fills a sheet weapon from a compendium item: damage
// "1d8 slashing", finesse weapons use the better of STR and DEX, and
// ammunition (ranged) weapons DEX.
func compendiumWeapon(name string, d *compendium.ItemData) dnd5e.Weapon {
	w := dnd5e.Weapon{Name: name, Ability: "str", Proficient: true, Range: d.Range, Properties: d.Properties}
	w.Dice, w.DamageType, _ = strings.Cut(strings.TrimSpace(d.Damage), " ")
	props := strings.ToLower(d.Properties)
	switch {
	case strings.Contains(props, "finesse"):
		w.Ability = "finesse"
	case strings.Contains(props, "ammunition"):
		w.Ability = "dex"
	}
	return w
}

// textMarkdown is a feat's (or trait's) description with its named parts.
func textMarkdown(d *compendium.TextData) string {
	parts := []string{d.Desc}
	if d.Prerequisite != "" {
		parts = append([]string{"*Prerequisite: " + d.Prerequisite + "*"}, parts...)
	}
	for _, p := range d.Parts {
		parts = append(parts, "**"+p.Name+".** "+p.Desc)
	}
	return strings.TrimSpace(strings.Join(parts, "\n\n"))
}

// builderPicks is the builder's compendium look-up: weapons for an attack,
// armor to wear, or (as featurePicker's options) feats and class features.
// Each weapon or feature pick reloads the form filled in from it.
func (a *App) builderPicks(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player changes their character.")
		return
	}
	q := strings.TrimSpace(c.Query("q"))
	if q == "" {
		c.Status(http.StatusOK)
		return
	}
	var picks []views.BuilderPick
	var err error
	switch c.Query("pick") {
	case "weapon", "armor":
		picks, err = a.itemPicks(c.Request.Context(), q, ch.CampaignID, c.Query("pick") == "armor")
	default:
		picks, err = a.featurePicks(c.Request.Context(), q, ch.CampaignID)
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	if c.Query("format") == "combobox" {
		a.render(c, http.StatusOK, views.BuilderPickItems(ch.ID, picks))
		return
	}
	a.render(c, http.StatusOK, views.BuilderPicks(ch.ID, c.Query("pick"), picks))
}

// itemPicks are compendium weapons, or armor and shields.
func (a *App) itemPicks(ctx context.Context, q string, campaignID int64, armor bool) ([]views.BuilderPick, error) {
	rows, err := a.searchRows(ctx, q, string(compendium.KindItem), campaignID)
	if err != nil {
		return nil, err
	}
	var picks []views.BuilderPick
	for _, r := range rows {
		if len(picks) == maxPicks {
			break
		}
		d, err := a.entryData(ctx, r.ID)
		if err != nil {
			return nil, err
		}
		it := d.(*compendium.ItemData)
		switch {
		case armor && it.Armor != nil:
			detail := it.Armor.Type + " · AC " + it.ArmorClass
			if it.Armor.StrReq > 0 {
				detail += fmt.Sprintf(" · Str %d", it.Armor.StrReq)
			}
			picks = append(picks, views.BuilderPick{EntryID: r.ID, Slug: r.Slug, Name: r.Name, Detail: detail})
		case !armor && it.Damage != "":
			picks = append(picks, views.BuilderPick{EntryID: r.ID, Slug: r.Slug, Name: r.Name, Detail: strings.Trim(it.Damage+" · "+it.Properties, " ·")})
		}
	}
	return picks, nil
}

// entryData is a compendium entry's decoded data.
func (a *App) entryData(ctx context.Context, id int64) (any, error) {
	e, err := a.Store.GetCompendiumEntry(ctx, id)
	if err != nil {
		return nil, err
	}
	return compendium.DecodeData(compendium.Kind(e.Kind), e.DataJson)
}

// featurePicks are feats and class features named like q. Class features
// live inside their class, so classes whose text matches are opened up.
func (a *App) featurePicks(ctx context.Context, q string, campaignID int64) ([]views.BuilderPick, error) {
	feats, err := a.searchRows(ctx, q, string(compendium.KindFeat), campaignID)
	if err != nil {
		return nil, err
	}
	var picks []views.BuilderPick
	for _, r := range feats {
		picks = append(picks, views.BuilderPick{EntryID: r.ID, Name: r.Name, Detail: "Feat"})
	}
	classes, err := a.searchRows(ctx, q, string(compendium.KindClass), campaignID)
	if err != nil {
		return nil, err
	}
	words := strings.Fields(strings.ToLower(q))
	for _, r := range classes {
		data, err := a.entryData(ctx, r.ID)
		if err != nil {
			return nil, err
		}
		for _, f := range data.(*compendium.ClassData).Features {
			name := strings.ToLower(f.Name)
			if !slices.ContainsFunc(words, func(w string) bool { return !strings.Contains(name, w) }) && len(f.Levels) > 0 {
				picks = append(picks, views.BuilderPick{EntryID: r.ID, Feature: f.Name, Name: f.Name, Detail: r.Name + " " + strconv.Itoa(f.Levels[0])})
			}
		}
	}
	return picks[:min(len(picks), maxPicks)], nil
}

// commandCompletions lists what the command bar's Tab offers for a
// placeholder: names from the sheet, or spell slugs from the compendium.
func (a *App) commandCompletions(c *gin.Context) {
	_, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	words := strings.Fields(c.Query("words"))
	q := strings.ToLower(c.Query("q"))
	out := []string{}
	if len(words) >= 2 && strings.EqualFold(words[0], "spell") && strings.EqualFold(words[1], "add") {
		if q != "" {
			rows, err := a.searchRows(c.Request.Context(), q, string(compendium.KindSpell), ch.CampaignID)
			if err != nil {
				a.serverError(c, err)
				return
			}
			for _, r := range rows {
				out = append(out, r.Slug)
			}
		}
	} else {
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil {
			a.serverError(c, err)
			return
		}
		for _, v := range dnd5e.Completions(sheet, words, c.Query("slot")) {
			if strings.HasPrefix(strings.ToLower(v), q) {
				out = append(out, v)
			}
		}
	}
	c.JSON(http.StatusOK, out[:min(len(out), maxPicks)])
}

// searchRows is a compendium search of one kind, name matches first.
func (a *App) searchRows(ctx context.Context, q, kind string, campaignID int64) ([]sqlc.SearchCompendiumRow, error) {
	names, text, err := compendium.Search(ctx, a.Store.Queries, q, kind, "", campaignID, maxPicks)
	return append(names, text...), err
}

// featureCommand saves a feature from the builder with what the sheet
// tracks for it: action, uses, unarmored AC and stat change rows.
func featureCommand(c *gin.Context) (dnd5e.Command, error) {
	f := dnd5e.Feature{
		Name: strings.TrimSpace(c.PostForm("subject")), Source: c.PostForm("source"), Description: c.PostForm("text"),
		Action: strings.TrimSpace(c.PostForm("action")), Unarmored: strings.TrimSpace(c.PostForm("unarmored")), NoShield: c.PostForm("no_shield") != "",
	}
	if f.Name == "" {
		return dnd5e.Command{}, errors.New("name the feature")
	}
	f.Uses.Reset = c.PostForm("reset")
	switch u := strings.TrimSpace(c.PostForm("uses")); u {
	case "":
		f.Uses.Reset = ""
	case "n":
		f.Uses.Fixed, _ = strconv.Atoi(c.PostForm("uses_n"))
	case "prof":
		f.Uses.Prof = true
	case "table":
		if err := json.Unmarshal([]byte(c.PostForm("uses_table")), &f.Uses.ByLevel); err != nil {
			return dnd5e.Command{}, err
		}
		f.Uses.Class = c.PostForm("uses_class")
	default:
		f.Uses.Ability = u
	}
	targets, values, whens := c.PostFormArray("mod_target"), c.PostFormArray("mod_value"), c.PostFormArray("mod_when")
	var mods []dnd5e.CustomMod
	for i, t := range targets {
		t = strings.TrimSpace(t)
		v, _ := strconv.Atoi(values[min(i, len(values)-1)])
		if t == "" || v == 0 {
			continue
		}
		if !dnd5e.IsModTarget(t) {
			return dnd5e.Command{}, fmt.Errorf("cannot change %q", t)
		}
		mods = append(mods, dnd5e.CustomMod{Target: t, Value: v, When: strings.TrimSpace(whens[min(i, len(whens)-1)])})
	}
	return dnd5e.SaveFeature(f, mods), nil
}

// weaponCommand saves a weapon from the builder, range and properties too.
func weaponCommand(c *gin.Context) (dnd5e.Command, error) {
	name := strings.TrimSpace(c.PostForm("subject"))
	if name == "" {
		return dnd5e.Command{}, errors.New("name the weapon")
	}
	return dnd5e.SaveWeapon(dnd5e.Weapon{
		Name: name, Ability: strings.ToLower(c.PostForm("ability")), Dice: strings.ToLower(strings.TrimSpace(c.PostForm("dice"))),
		DamageType: strings.TrimSpace(c.PostForm("dtype")), Proficient: c.PostForm("prof") != "",
		Range: strings.TrimSpace(c.PostForm("range")), Properties: strings.TrimSpace(c.PostForm("properties")),
	}), nil
}
