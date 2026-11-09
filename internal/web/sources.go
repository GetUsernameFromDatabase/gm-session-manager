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
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// refCompendiumKind is the compendium kind each sheet ref kind names.
var refCompendiumKind = map[string]compendium.Kind{
	dnd5e.RefClass: compendium.KindClass, dnd5e.RefSubclass: compendium.KindSubclass, dnd5e.RefSpell: compendium.KindSpell,
	dnd5e.RefSpecies: compendium.KindSpecies, dnd5e.RefBackground: compendium.KindBackground,
	dnd5e.RefFeat: compendium.KindFeat, dnd5e.RefArmor: compendium.KindItem,
}

// entryFor is the copy of kind/slug a sheet ref pinned to pin uses: the
// pinned copy while it exists, else the priority copy, with missing set when
// a pin was passed over. sql.ErrNoRows when neither exists. It takes q, not
// the App, so upgrades can run it inside their transaction.
func entryFor(ctx context.Context, q *sqlc.Queries, kind compendium.Kind, slug, pin string) (row sqlc.CompendiumEntry, missing bool, err error) {
	if pin != "" {
		row, err = q.GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: string(kind), Slug: slug, SourceKey: pin})
		if !errors.Is(err, sql.ErrNoRows) {
			return row, false, err
		}
	}
	row, err = q.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(kind), Slug: slug})
	return row, pin != "", err
}

// classEntry is entryFor for a class, as a ClassInfo.
func classEntry(ctx context.Context, q *sqlc.Queries, slug, pin string) (dnd5e.ClassInfo, bool, error) {
	row, missing, err := entryFor(ctx, q, compendium.KindClass, slug, pin)
	if err != nil {
		return dnd5e.ClassInfo{}, missing, err
	}
	info, err := classInfo(row)
	return info, missing, err
}

// classPin is the source the sheet pins its class slug to; "" for an
// unpinned or new class.
func classPin(s dnd5e.Sheet, slug string) string {
	if i := slices.IndexFunc(s.Classes, func(cl dnd5e.ClassLevel) bool { return cl.Slug == slug }); i >= 0 {
		return s.Classes[i].Source
	}
	return ""
}

// sheetRef is a pinned compendium ref on a sheet.
type sheetRef struct{ slug, name, pin string }

// pinnedRefs are the sheet's pinned refs by ref kind; unpinned ones follow
// priority and never warn.
func pinnedRefs(s dnd5e.Sheet) map[string][]sheetRef {
	refs := map[string][]sheetRef{}
	add := func(kind, slug, name, pin string) {
		if slug != "" && pin != "" {
			refs[kind] = append(refs[kind], sheetRef{slug, name, pin})
		}
	}
	for _, cl := range s.Classes {
		add(dnd5e.RefClass, cl.Slug, cl.Name, cl.Source)
		add(dnd5e.RefSubclass, cl.SubclassSlug, cl.Subclass, cl.SubclassSource)
	}
	for _, sp := range s.Spells {
		add(dnd5e.RefSpell, sp.Slug, sp.Name, sp.Source)
	}
	add(dnd5e.RefSpecies, compendium.Slugify(s.Species), s.Species, s.SpeciesSource)
	add(dnd5e.RefBackground, compendium.Slugify(s.Background), s.Background, s.BackgroundSource)
	for _, f := range s.Feats {
		base, _, _ := strings.Cut(f, " (")
		slug := featKey(f)
		add(dnd5e.RefFeat, slug, base, s.FeatSources[slug])
	}
	if s.Armor != nil {
		add(dnd5e.RefArmor, s.Armor.Slug, s.Armor.Name, s.Armor.Source)
	}
	return refs
}

// sourceIssues are the sheet's pinned refs whose copy is gone or outranked,
// with one query per ref kind.
func (a *App) sourceIssues(ctx context.Context, s dnd5e.Sheet, campaignID int64) ([]views.SourceIssue, error) {
	refs := pinnedRefs(s)
	var out []views.SourceIssue
	for _, kind := range dnd5e.RefKinds {
		list := refs[kind]
		if len(list) == 0 {
			continue
		}
		slugs := make([]string, len(list))
		for i, r := range list {
			slugs[i] = r.slug
		}
		rows, err := a.Store.ListEntrySourcesBySlugs(ctx, sqlc.ListEntrySourcesBySlugsParams{Kind: string(refCompendiumKind[kind]), Slugs: slugs, CampaignID: campaignID})
		if err != nil {
			return nil, err
		}
		copies := map[string][]views.SourceCopy{}
		for _, r := range rows {
			copies[r.Slug] = append(copies[r.Slug], views.SourceCopy{Key: r.SourceKey, Label: views.SourceLabel(r.SourceKey, r.Edition)})
		}
		seen := map[string]bool{}
		for _, r := range list {
			if seen[r.slug] {
				continue // a feat taken twice, a spell from two classes
			}
			seen[r.slug] = true
			if is, ok := sourceIssue(kind, r, copies[r.slug]); ok {
				out = append(out, is)
			}
		}
	}
	return out, nil
}

// sourceIssue reports r when its pin is not among copies (best first) or a
// copy outranks it.
func sourceIssue(kind string, r sheetRef, copies []views.SourceCopy) (views.SourceIssue, bool) {
	is := views.SourceIssue{Kind: kind, Slug: r.slug, Name: r.name, Pinned: r.pin, Copies: copies}
	if len(copies) > 0 {
		is.Best = copies[0].Key
	}
	is.Missing = !slices.ContainsFunc(copies, func(c views.SourceCopy) bool { return c.Key == r.pin })
	return is, is.Missing || is.Best != r.pin
}

// sourceEnv resolves what the source command for ref needs: the copy at
// ref.Source and, for a class or subclass, the pinned copy it replaces; name
// is that copy's name.
func (a *App) sourceEnv(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet, ref dnd5e.SourceRef, campaignID int64) (env dnd5e.Env, name string, err error) {
	row, err := a.visibleCopy(ctx, q, ref, campaignID)
	if err != nil {
		return env, "", err
	}
	switch ref.Kind {
	case dnd5e.RefSpell:
		env.SourceSpell, err = spellFromRow(row)
	case dnd5e.RefClass:
		env.SourceOld, env.SourceNew, err = a.classSwitch(ctx, q, sheet, row)
	case dnd5e.RefSubclass:
		env.SourceOld, env.SourceNew, err = subclassSwitch(ctx, q, sheet, row)
	}
	return env, row.Name, err
}

// visibleCopy is ref's entry at ref.Source when campaignID may see it; a
// userError otherwise.
func (a *App) visibleCopy(ctx context.Context, q *sqlc.Queries, ref dnd5e.SourceRef, campaignID int64) (sqlc.CompendiumEntry, error) {
	row, err := q.GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: string(refCompendiumKind[ref.Kind]), Slug: ref.Slug, SourceKey: ref.Source})
	if err == nil {
		var visible bool
		if visible, err = a.entryVisible(ctx, row, campaignID); err == nil && !visible {
			err = sql.ErrNoRows
		}
	}
	if errors.Is(err, sql.ErrNoRows) {
		return row, userError{fmt.Errorf("%s has no %s %q", ref.Source, ref.Kind, ref.Slug)}
	}
	return row, err
}

// subclassChangeEnv is the sheet's class ref.Class with its current
// subclass (nil when that copy is gone) and with ref's, which must be a
// subclass of that class the campaign sees.
func (a *App) subclassChangeEnv(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet, ref dnd5e.SubclassRef, campaignID int64) (*dnd5e.ClassInfo, dnd5e.ClassInfo, error) {
	i := slices.IndexFunc(sheet.Classes, func(cl dnd5e.ClassLevel) bool { return cl.Slug == ref.Class })
	if i < 0 {
		return nil, dnd5e.ClassInfo{}, nil // ChangeSubclass refuses it
	}
	row, err := a.visibleCopy(ctx, q, dnd5e.SourceRef{Kind: dnd5e.RefSubclass, Slug: ref.Slug, Source: ref.Source}, campaignID)
	if err != nil {
		return nil, dnd5e.ClassInfo{}, err
	}
	d, err := compendium.DecodeData(compendium.KindSubclass, row.DataJson)
	if err != nil {
		return nil, dnd5e.ClassInfo{}, err
	}
	if d.(*compendium.ClassData).SubclassOf != ref.Class {
		return nil, dnd5e.ClassInfo{}, userError{fmt.Errorf("%s is not a %s subclass", row.Name, sheet.Classes[i].Name)}
	}
	cl := sheet.Classes[i]
	oldSlug, oldPin, err := subclassRef(ctx, q, cl.Slug, cl, cl.Subclass)
	if err != nil {
		return nil, dnd5e.ClassInfo{}, err
	}
	return subclassPair(ctx, q, cl, oldSlug, oldPin, row)
}

// classSwitch is the sheet's pinned copy of row's class (nil when gone) and
// row's copy, each with the subclass: the new one from row's source when it
// has it, else the old one's.
func (a *App) classSwitch(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet, row sqlc.CompendiumEntry) (*dnd5e.ClassInfo, dnd5e.ClassInfo, error) {
	next, err := classInfo(row)
	i := slices.IndexFunc(sheet.Classes, func(cl dnd5e.ClassLevel) bool { return cl.Slug == row.Slug })
	if err != nil || i < 0 {
		return nil, next, err // SwitchSource refuses a class the sheet lacks
	}
	cl := sheet.Classes[i]
	var old *dnd5e.ClassInfo
	info, missing, err := classEntry(ctx, q, cl.Slug, cl.Source)
	switch {
	case err == nil && !missing:
		if err := a.withSubclass(ctx, &info, sheet, ""); err != nil {
			return nil, next, err
		}
		old = &info
	case err != nil && !errors.Is(err, sql.ErrNoRows):
		return nil, next, err
	}
	if cl.Subclass == "" {
		return old, next, nil
	}
	subSlug, _, err := subclassRef(ctx, q, cl.Slug, cl, cl.Subclass)
	if err != nil || subSlug == "" {
		return old, next, err
	}
	sub, err := q.GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: string(compendium.KindSubclass), Slug: subSlug, SourceKey: row.SourceKey})
	switch {
	case err == nil:
		err = setSubclass(&next, sub)
	case errors.Is(err, sql.ErrNoRows):
		err = nil
		if old != nil {
			next.SubclassFeatures, next.SubclassSlug, next.SubclassSource = old.SubclassFeatures, old.SubclassSlug, old.SubclassSource
		}
	}
	return old, next, err
}

// subclassSwitch is the sheet's class with its pinned copy of row's
// subclass (nil when gone) and with row's.
func subclassSwitch(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet, row sqlc.CompendiumEntry) (*dnd5e.ClassInfo, dnd5e.ClassInfo, error) {
	i := slices.IndexFunc(sheet.Classes, func(cl dnd5e.ClassLevel) bool {
		return cl.SubclassSlug == row.Slug || cl.SubclassSlug == "" && cl.Subclass != "" && compendium.Slugify(cl.Subclass) == row.Slug
	})
	if i < 0 {
		return nil, dnd5e.ClassInfo{}, nil // SwitchSource refuses it
	}
	return subclassPair(ctx, q, sheet.Classes[i], row.Slug, sheet.Classes[i].SubclassSource, row)
}

// subclassPair is cl's class with the subclass oldSlug at oldPin (nil when
// that copy is gone or oldSlug is "") and with row's.
func subclassPair(ctx context.Context, q *sqlc.Queries, cl dnd5e.ClassLevel, oldSlug, oldPin string, row sqlc.CompendiumEntry) (*dnd5e.ClassInfo, dnd5e.ClassInfo, error) {
	base, _, err := classEntry(ctx, q, cl.Slug, cl.Source)
	if errors.Is(err, sql.ErrNoRows) {
		base, err = dnd5e.ClassInfo{Slug: cl.Slug, Name: cl.Name}, nil
	}
	if err != nil {
		return nil, base, err
	}
	next := base
	if err := setSubclass(&next, row); err != nil {
		return nil, next, err
	}
	if oldSlug == "" {
		return nil, next, nil
	}
	oldRow, missing, err := entryFor(ctx, q, compendium.KindSubclass, oldSlug, oldPin)
	if errors.Is(err, sql.ErrNoRows) || err == nil && missing {
		return nil, next, nil
	}
	if err != nil {
		return nil, next, err
	}
	old := base
	return &old, next, setSubclass(&old, oldRow)
}

// sourcePreviewForm asks what switching a ref to another source does;
// kind is one of dnd5e.RefKinds.
type sourcePreviewForm struct {
	Kind   string `form:"kind"   validate:"oneof=class subclass spell species background feat armor"`
	Slug   string `form:"slug"   validate:"required,max=200"`
	Source string `form:"source" validate:"required,max=200"`
}

// sourcePreview lists what the source command would replace, add and keep,
// with the command Confirm posts.
func (a *App) sourcePreview(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSources) {
		a.fail(c, http.StatusForbidden, "Only the player and the GM switch sources.")
		return
	}
	var f sourcePreviewForm
	if c.ShouldBindQuery(&f) != nil {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a source for the entry.")
		return
	}
	// Parsed as the command endpoint will, so Confirm posts a valid line.
	line := fmt.Sprintf("source %s %s %s", f.Kind, f.Slug, f.Source)
	cmd, err := dnd5e.Parse(line)
	if err != nil {
		a.fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	sheet, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		a.serverError(c, err)
		return
	}
	ctx := c.Request.Context()
	ref := cmd.SourceRef
	env, name, err := a.sourceEnv(ctx, a.Store.Queries, sheet, ref, ch.CampaignID)
	if a.commandFailed(c, err) {
		return
	}
	src, err := a.Store.GetCompendiumSource(ctx, ref.Source)
	if err != nil && !errors.Is(err, sql.ErrNoRows) { // homebrew may have no source row
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.SourcePreview(views.SourcePreviewView{
		CharacterID: ch.ID, Kind: ref.Kind, Slug: ref.Slug, Source: ref.Source, Name: name,
		SourceLabel: views.SourceLabel(ref.Source, src.Edition), Command: line,
		Plan: dnd5e.PlanSourceSwitch(sheet, ref.Kind, ref.Slug, env.SourceOld, env.SourceNew),
	}))
}

// sourceUsers names the characters whose sheet pins the source, and the
// entities and combatants linked to one of its entries. C03 reuses it for
// the per-campaign untick rule.
func sourceUsers(ctx context.Context, q *sqlc.Queries, key string) ([]string, error) {
	rows, err := q.ListCharacterSheetsNamed(ctx)
	if err != nil {
		return nil, err
	}
	var users []string
	for _, r := range rows {
		if sheet, err := dnd5e.ParseSheet(r.SheetJson); err == nil && sheet.PinsSource(key) {
			users = append(users, r.Name)
		}
	}
	linked, err := q.ListSourceRefUsers(ctx, key)
	if err != nil {
		return nil, err
	}
	return append(users, linked...), nil
}

// sourceInUseMsg is the refusal naming who uses the source.
func sourceInUseMsg(key string, users []string) string {
	return key + " is used by " + strings.Join(users, ", ") + ". Switch their sources first."
}

// sourceDelete removes a source and its entries unless something uses it.
func (a *App) sourceDelete(c *gin.Context) {
	ctx := c.Request.Context()
	key := c.Param("key")
	if key == compendium.HomebrewSource {
		a.fail(c, http.StatusBadRequest, "Homebrew cannot be deleted.")
		return
	}
	if a.fiveTools.get().Running {
		a.fail(c, http.StatusConflict, errImportRunning)
		return
	}
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		users, err := sourceUsers(ctx, q, key)
		if err != nil {
			return err
		}
		if len(users) > 0 {
			return &sourceUsedError{users}
		}
		if err := q.DeleteCompendiumEntriesBySource(ctx, key); err != nil {
			return err
		}
		return q.DeleteCompendiumSource(ctx, key)
	})
	var used *sourceUsedError
	if errors.As(err, &used) {
		a.fail(c, http.StatusConflict, sourceInUseMsg(key, used.users))
		return
	}
	if err == nil {
		err = compendium.RemoveSourceDir(a.compendiumDir(), key)
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.sourcesChanged(c)
}

// sourceUsedError carries the users out of the delete transaction.
type sourceUsedError struct{ users []string }

func (e *sourceUsedError) Error() string { return "source in use" }
