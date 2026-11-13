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
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// withSubclass adds the features of the class's subclass named name (the
// sheet's, or the one being chosen in campaignID) to info; an unknown name
// adds none.
func (a *App) withSubclass(ctx context.Context, info *dnd5e.ClassInfo, sheet dnd5e.Sheet, name string, campaignID int64) error {
	var cl dnd5e.ClassLevel
	if i := slices.IndexFunc(sheet.Classes, func(c dnd5e.ClassLevel) bool { return c.Slug == info.Slug }); i >= 0 {
		cl = sheet.Classes[i]
	}
	if name == "" {
		name, campaignID = cl.Subclass, 0 // the sheet keeps its own, ticked or not
	}
	if name == "" {
		return nil
	}
	slug, pin, err := subclassRef(ctx, a.Store.Queries, info.Slug, cl, name, campaignID)
	if err != nil || slug == "" {
		return err
	}
	row, _, err := entryFor(ctx, a.Store.Queries, enum.KindSubclass, slug, pin)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	if err != nil {
		return err
	}
	return setSubclass(info, row)
}

// subclassRef is the slug and pin of classSlug's subclass with slug or name
// name: the sheet's pin when it is cl's subclass, else found among those
// campaignID allows (0: all; "" for none). A subclass only an unticked
// source has is a userError.
func subclassRef(ctx context.Context, q *sqlc.Queries, classSlug string, cl dnd5e.ClassLevel, name string, campaignID int64) (slug, pin string, err error) {
	if cl.SubclassSlug != "" && (name == cl.SubclassSlug || strings.EqualFold(name, cl.Subclass)) {
		return cl.SubclassSlug, cl.SubclassSource, nil
	}
	subs, err := q.ListSubclassesOf(ctx, sqlc.ListSubclassesOfParams{ClassSlug: classSlug, CampaignID: campaignID})
	if err != nil {
		return "", "", err
	}
	if i := subclassIndex(subs, name); i >= 0 {
		return subs[i].Slug, "", nil
	}
	if campaignID != 0 {
		if slug, _, err = subclassRef(ctx, q, classSlug, cl, name, 0); err == nil && slug != "" {
			err = userError{fmt.Errorf("%s is not in this campaign's sources", name)}
		}
	}
	return "", "", err
}

// subclassIndex finds v in subs by slug (the picker's value: two sources may
// name a subclass alike), else by name (typed commands), or -1.
func subclassIndex(subs []sqlc.ListSubclassesOfRow, v string) int {
	if i := slices.IndexFunc(subs, func(s sqlc.ListSubclassesOfRow) bool { return s.Slug == v }); i >= 0 {
		return i
	}
	return slices.IndexFunc(subs, func(s sqlc.ListSubclassesOfRow) bool { return strings.EqualFold(s.Name, v) })
}

// subclassOptions are the picker's choices; a name two rows share is
// labelled with each one's source.
func subclassOptions(subs []sqlc.ListSubclassesOfRow) []views.Option {
	out := make([]views.Option, 0, len(subs))
	for _, s := range subs {
		name := s.Name
		if slices.ContainsFunc(subs, func(o sqlc.ListSubclassesOfRow) bool { return o.Slug != s.Slug && strings.EqualFold(o.Name, s.Name) }) {
			name += " · " + s.SourceKey
		}
		out = append(out, views.Option{Slug: s.Slug, Name: name})
	}
	return out
}

// setSubclass gives info the features and pin of the subclass row, and its
// spellcasting when the class has none (Eldritch Knight, Arcane Trickster).
func setSubclass(info *dnd5e.ClassInfo, row sqlc.CompendiumEntry) error {
	sub, err := compendium.Decode[compendium.ClassData](enum.KindSubclass, row.DataJson)
	if err != nil {
		return err
	}
	if sub.Caster != "" && sub.Caster != dnd5e.CasterNone && (info.Caster == "" || info.Caster == dnd5e.CasterNone) {
		info.Caster = sub.Caster
	}
	info.SubclassFeatures = classFeatures(sub)
	info.SubclassName, info.SubclassSlug, info.SubclassSource = row.Name, row.Slug, row.SourceKey
	return nil
}

// subclassPanels is the details panel of each subclass row, features by level.
func subclassPanels(rows []sqlc.ListSubclassesOfRow) ([]views.SubclassPanel, error) {
	out := make([]views.SubclassPanel, 0, len(rows))
	for _, r := range rows {
		d, err := compendium.Decode[compendium.ClassData](enum.KindSubclass, r.DataJson)
		if err != nil {
			return nil, err
		}
		p := views.SubclassPanel{Slug: r.Slug, Name: r.Name, Source: r.SourceKey}
		for _, f := range d.Features {
			p.Features = append(p.Features, views.SubclassFeature{Name: f.Name, Desc: f.Desc, Levels: f.Levels})
		}
		slices.SortStableFunc(p.Features, func(a, b views.SubclassFeature) int { return cmp.Compare(a.FirstLevel(), b.FirstLevel()) })
		out = append(out, p)
	}
	return out, nil
}

// subclassForm swaps the subclass of the sheet's class Class for the one
// with slug (the picker's value) or name Subclass.
type subclassForm struct {
	Class    string `form:"class"    validate:"required,max=200"`
	Subclass string `form:"subclass" validate:"required,max=200"`
}

// subclassOwner is the character when the viewer may change its subclass.
func (a *App) subclassOwner(c *gin.Context) (sqlc.Character, bool) {
	w, ch, ok := a.characterAccess(c)
	if ok && !perm.CanChangeSubclass(w.viewer, ch.ID) {
		a.fail(c, http.StatusForbidden, "Only the player changes their subclass.")
		ok = false
	}
	return ch, ok
}

// subclassChangeForm renders the subclass picker for one class into the
// command builder dialog.
func (a *App) subclassChangeForm(c *gin.Context) error {
	ch, ok := a.subclassOwner(c)
	if !ok {
		return nil
	}
	sheet, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		return err
	}
	i := slices.IndexFunc(sheet.Classes, func(cl dnd5e.ClassLevel) bool { return cl.Slug == c.Query("class") })
	if i < 0 || sheet.Classes[i].Subclass == "" {
		a.fail(c, http.StatusNotFound, "That class has no subclass to change.")
		return nil
	}
	cl := sheet.Classes[i]
	subs, err := a.Store.ListSubclassesOf(c.Request.Context(), sqlc.ListSubclassesOfParams{ClassSlug: cl.Slug, CampaignID: ch.CampaignID})
	if err != nil {
		return err
	}
	v := views.SubclassChangeView{
		CharacterID: ch.ID, Class: cl.Slug, ClassName: cl.Name, Level: cl.Level, Subclasses: subclassOptions(subs),
		Current: cl.Subclass, CurrentSlug: cmp.Or(cl.SubclassSlug, compendium.Slugify(cl.Subclass)),
	}
	if v.Panels, err = subclassPanels(subs); err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.SubclassChange(v))
	return nil
}

// changeSubclass resolves the picked name to a compendium copy and runs
// change-subclass with it.
func (a *App) changeSubclass(c *gin.Context) error {
	ch, ok := a.subclassOwner(c)
	if !ok {
		return nil
	}
	var f subclassForm
	if c.ShouldBind(&f) != nil {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a subclass.")
		return nil
	}
	ctx := c.Request.Context()
	subs, err := a.Store.ListSubclassesOf(ctx, sqlc.ListSubclassesOfParams{ClassSlug: f.Class, CampaignID: ch.CampaignID})
	if err != nil {
		return err
	}
	i := subclassIndex(subs, f.Subclass)
	if i < 0 {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a subclass from the list.")
		return nil
	}
	cmd, err := dnd5e.Parse("change-subclass", f.Class, subs[i].Slug, subs[i].SourceKey)
	if err != nil {
		return userError{err}
	}
	if err := a.runCommand(ctx, ch, cmd); err != nil {
		return err
	}
	a.characterChanged(c, ch)
	return nil
}
