package web

import (
	"cmp"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// compendiumRoutes are the compendium's pages under one base (g).
func (a *App) compendiumRoutes(g *gin.RouterGroup, gm gin.HandlerFunc) {
	g.GET("", a.compendiumPage)
	g.GET("/search", a.compendiumSearch)
	g.GET("/random", a.compendiumRandom)
	g.GET("/lookup", a.compendiumLookup)
	g.GET("/find/:kind/:slug", a.compendiumFind)
	g.GET("/entries/:entry", a.compendiumEntry)
	g.GET("/compare/:kind/:slug", gm, a.compendiumCompare)
	// Homebrew: the handlers check who may write (perm.CanEditHomebrew).
	g.GET("/entries/:entry/edit", a.homebrewEdit)
	g.POST("/entries/:entry", a.homebrewUpdate)
	g.POST("/entries/:entry/delete", a.homebrewDelete)
	g.POST("/entries/:entry/campaigns", gm, a.homebrewLink)
	g.POST("/entries/:entry/campaigns/:linked/delete", gm, a.homebrewUnlink)
	g.POST("/entries/:entry/image", gm, a.compendiumImage)
	g.GET("/habitats", gm, a.habitatsPage)
	g.POST("/habitats/cell", gm, a.habitatCell)
	g.POST("/habitats/any", gm, a.habitatAny)
	g.POST("/habitats/batch", gm, a.habitatBatch)
	g.POST("/habitats/fill", gm, a.habitatFill)
	g.POST("/habitats/copy", gm, a.habitatCopy)
	g.GET("/homebrew/new", a.homebrewNew)
	g.POST("/homebrew", a.homebrewCreate)
	g.POST("/import", gm, a.compendiumImport)
	g.POST("/folder", gm, a.compendiumFolder)
	g.GET("/5etools", gm, a.fiveToolsList)
	g.POST("/5etools", gm, a.fiveToolsImport)
	g.POST("/5etools/:book/delete", gm, a.fiveToolsRemove)
	g.POST("/sources/:key/move", gm, a.sourceMove)
	g.POST("/sources/:key/delete", gm, a.sourceDelete)
	g.POST(views.CompendiumRest(views.SourceOrderPath), gm, a.sourceOrder)
	g.POST("/page/preview", gm, a.compendiumPagePreview)
	g.POST("/page", gm, a.compendiumPageSave)
}

// Keys for what compendiumCampaign resolved: the path's campaign and the
// requester, so the handler does not look them up again.
const (
	compendiumCampKey = "compendium-campaign"
	compendiumWhoKey  = "compendium-who"
)

// compendiumCampaign opens the compendium as one campaign sees it: any
// campaign for the GM, only their own for a player.
func (a *App) compendiumCampaign(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		c.Abort()
		return
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		c.Abort()
		return
	}
	c.Set(compendiumCampKey, &camp)
	c.Set(compendiumWhoKey, w)
	c.Request = c.Request.WithContext(views.WithCampaign(c.Request.Context(), camp.ID))
}

// compendiumWhole sends page loads of the unscoped compendium to the
// campaign they show, so the address names it: a player's own campaign, and
// the GM's old ?campaign= links. Background requests (htmx) stay put.
func (a *App) compendiumWhole(c *gin.Context) {
	if c.Request.Method != http.MethodGet || isHTMX(c) {
		return
	}
	q := c.Request.URL.Query()
	id, _ := strconv.ParseInt(q.Get(oldCampaignKey), 10, 64)
	q.Del(oldCampaignKey)
	if id == 0 {
		w, err := a.whoami(c)
		if err != nil {
			a.serverError(c, err)
			c.Abort()
			return
		}
		if w.isGM() {
			return
		}
		id = w.campaignID()
	}
	if id == 0 {
		return // a visitor: compendiumReader refuses them
	}
	target := views.CompendiumBase(id) + views.CompendiumRest(c.Request.URL.Path)
	if len(q) > 0 {
		target += "?" + q.Encode()
	}
	c.Redirect(http.StatusSeeOther, target)
	c.Abort()
}

// oldCampaignKey is the query key that named the GM's campaign before the
// campaign moved into the path.
const oldCampaignKey = "campaign"

// compendiumReader allows the GM and any player: the compendium is rules
// text, shared by every campaign. The campaign is the one in the path, or
// the player's (nil: everything, for the GM).
func (a *App) compendiumReader(c *gin.Context) (who, *sqlc.Campaign, bool) {
	if camp := pathCampaign(c); camp != nil {
		return c.MustGet(compendiumWhoKey).(who), camp, true // compendiumCampaign checked access
	}
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return w, nil, false
	}
	if !w.isGM() && w.character == nil {
		a.fail(c, http.StatusForbidden, "Join a campaign first.")
		return w, nil, false
	}
	if w.isGM() {
		return w, nil, true
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), w.campaignID())
	if !a.found(c, err) {
		return w, nil, false
	}
	return w, &camp, true
}

// pathCampaign is the campaign in the request's path (nil for the whole
// compendium).
func pathCampaign(c *gin.Context) *sqlc.Campaign {
	camp, _ := c.Get(compendiumCampKey)
	p, _ := camp.(*sqlc.Campaign)
	return p
}

// compendiumPath is rest under the request's compendium base.
func compendiumPath(c *gin.Context, rest string) string {
	return views.CompendiumPath(c.Request.Context(), rest)
}

// compendiumLayout is the frame of a compendium page: the campaign's, with
// the GM's switcher to the same page elsewhere.
func (a *App) compendiumLayout(c *gin.Context, w who, title string, camp *sqlc.Campaign) (views.Layout, error) {
	l := a.layout(w, title, views.NavCompendium, camp, false)
	if !w.isGM() {
		return l, nil
	}
	campaigns, err := a.Store.ListCampaigns(c.Request.Context())
	here := views.CompendiumRest(c.Request.URL.Path)
	if c.Request.Method != http.MethodGet {
		here = "" // a form's answer: the switcher opens the list
	} else if c.Request.URL.RawQuery != "" {
		here += "?" + c.Request.URL.RawQuery
	}
	l.Switch = &views.CampaignSwitch{Campaigns: campaigns, Base: views.CompendiumBase, All: "All sources", Here: here}
	return l, err
}

// viewerCampaign is the campaign whose sources and player homebrew the
// viewer sees: the player's, or 0 for the GM, who sees it all.
func viewerCampaign(camp *sqlc.Campaign) int64 {
	if camp == nil {
		return 0
	}
	return camp.ID
}

// entryVisible reports whether a player in campaignID may see entry e.
func (a *App) entryVisible(ctx context.Context, e sqlc.CompendiumEntry, campaignID int64) (bool, error) {
	if !e.Scoped || campaignID == 0 {
		return true, nil
	}
	linked, err := a.Store.ListEntryCampaigns(ctx, e.ID)
	return perm.CanSeeEntry(campaignID, e.Scoped, linked), err
}

// entryBySlug finds an entry by kind and slug that a player in campaignID
// may see; sql.ErrNoRows otherwise.
func (a *App) entryBySlug(ctx context.Context, kind compendium.Kind, slug string, campaignID int64) (sqlc.CompendiumEntry, error) {
	e, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(kind), Slug: slug})
	if err != nil {
		return e, err
	}
	visible, err := a.entryVisible(ctx, e, campaignID)
	if err == nil && !visible {
		err = sql.ErrNoRows
	}
	return e, err
}

// decodeLinked is e's data with its cross-reference tokens resolved for a
// viewer in campaignID.
func (a *App) decodeLinked(ctx context.Context, e sqlc.CompendiumEntry, campaignID int64) (any, error) {
	js, err := a.resolveCompendiumLinks(ctx, e.DataJson, campaignID)
	if err != nil {
		return nil, err
	}
	return compendium.DecodeData(compendium.Kind(e.Kind), js)
}

// compendiumBatch is how many rows one search or scroll step loads; the
// after_* query keys are the browse cursor.
const (
	compendiumBatch     = 100
	compendiumAfterName = "after_name"
	compendiumAfterID   = "after_id"
)

func (a *App) compendiumPage(c *gin.Context) {
	a.renderCompendium(c, views.ImportNotice{Tab: c.Query("import")})
}

func (a *App) renderCompendium(c *gin.Context, notice views.ImportNotice) {
	w, camp, ok := a.compendiumReader(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	v := views.CompendiumView{Query: strings.TrimSpace(c.Query("q")), Kind: c.Query("kind"), IsGM: w.isGM(), Import: notice, Tab: c.Query("tab"), Campaign: viewerCampaign(camp)}
	v.Import.Dir, v.Import.OnHost, v.Import.Job = a.compendiumDir(), w.id.Loopback, a.fiveTools.get()
	// Only monsters have habitats, so the filter applies to them alone.
	if v.Kind == string(compendium.KindMonster) {
		v.Habitat = c.Query("habitat")
	}
	var err error
	// A source's shared entries, from the GM's priority card; a search leaves them.
	if w.isGM() && v.Query == "" {
		v.Shared = c.Query(views.SharedKey)
	}
	if v.Shared != "" {
		if err := a.sharedEntries(c, &v); err != nil {
			a.serverError(c, err)
			return
		}
	} else if v.Query != "" {
		if v.Results, v.TextResults, err = compendium.Search(ctx, a.Store.Queries, v.Query, v.Kind, v.Habitat, viewerCampaign(camp), compendiumBatch); err != nil {
			a.serverError(c, err)
			return
		}
	} else {
		// Browse the kind ("" is every kind) a batch at a time; the last row
		// of a full batch is the cursor for the next.
		afterID, _ := strconv.ParseInt(c.Query(compendiumAfterID), 10, 64)
		rows, err := a.Store.ListCompendiumByKind(ctx, sqlc.ListCompendiumByKindParams{Kind: v.Kind, Habitat: v.Habitat, AfterName: c.Query(compendiumAfterName), AfterID: afterID, CampaignID: viewerCampaign(camp), PageSize: compendiumBatch})
		if err != nil {
			a.serverError(c, err)
			return
		}
		for _, r := range rows {
			v.Results = append(v.Results, sqlc.SearchCompendiumRow(r))
		}
		if len(rows) == compendiumBatch {
			last := rows[len(rows)-1]
			v.More = views.CompendiumURL(url.Values{"kind": {v.Kind}, "habitat": {v.Habitat}, compendiumAfterName: {last.Name}, compendiumAfterID: {strconv.FormatInt(last.ID, 10)}}, v.Campaign)
		}
	}
	if c.Query(compendiumAfterID) != "" {
		a.render(c, http.StatusOK, views.CompendiumRows(v))
		return
	}
	if v.Shared == "" { // the shared list counted its own kinds
		if v.Counts, err = a.Store.CountCompendium(ctx, viewerCampaign(camp)); err != nil {
			a.serverError(c, err)
			return
		}
	}
	if v.HabitatCounts, err = a.Store.CountHabitats(ctx, viewerCampaign(camp)); err != nil {
		a.serverError(c, err)
		return
	}
	if isHTMX(c) && c.GetHeader("HX-Target") == views.CompendiumResultsID {
		a.render(c, http.StatusOK, views.CompendiumSearch(v))
		return
	}
	if err := a.homebrewLists(ctx, w, &v); err != nil {
		a.serverError(c, err)
		return
	}
	if w.isGM() {
		if v.Sources, v.SharedCounts, err = a.sourcePriority(ctx, viewerCampaign(camp)); err != nil {
			a.serverError(c, err)
			return
		}
	}
	l, err := a.compendiumLayout(c, w, "Compendium", camp)
	if err != nil {
		a.serverError(c, err)
		return
	}
	if l.Switch != nil {
		v.Campaigns = l.Switch.Campaigns // the homebrew links pick from the same list
	}
	a.render(c, http.StatusOK, views.CompendiumPage(l, v))
}

// sharedEntries fills v with a batch of v.Shared's entries that another
// source of v.Campaign (v.With, if set) also holds, each with its copy
// order, and the counts for the source and kind chips.
func (a *App) sharedEntries(c *gin.Context, v *views.CompendiumView) error {
	ctx := c.Request.Context()
	v.With = c.Query(views.SharedWithKey)
	afterID, _ := strconv.ParseInt(c.Query(compendiumAfterID), 10, 64)
	rows, err := a.Store.ListSharedEntries(ctx, sqlc.ListSharedEntriesParams{SourceKey: v.Shared, With: v.With, Kind: v.Kind, AfterName: c.Query(compendiumAfterName), AfterID: afterID, CampaignID: v.Campaign, PageSize: compendiumBatch})
	if err != nil {
		return err
	}
	for _, r := range rows {
		v.Results = append(v.Results, sqlc.SearchCompendiumRow(r))
	}
	if len(rows) == compendiumBatch {
		last := rows[len(rows)-1]
		v.More = views.CompendiumURL(url.Values{views.SharedKey: {v.Shared}, views.SharedWithKey: {v.With}, "kind": {v.Kind}, compendiumAfterName: {last.Name}, compendiumAfterID: {strconv.FormatInt(last.ID, 10)}}, v.Campaign)
	}
	if len(rows) > 0 {
		order, err := a.Store.ListSharedOrder(ctx, sqlc.ListSharedOrderParams{SourceKey: v.Shared, CampaignID: v.Campaign})
		if err != nil {
			return err
		}
		v.CopyOrder = map[string][]string{}
		for _, o := range order {
			k := views.CopyKey(o.Kind, o.Slug)
			v.CopyOrder[k] = append(v.CopyOrder[k], o.SourceKey)
		}
	}
	if c.Query(compendiumAfterID) != "" {
		return nil // a later batch: the chips are already on the page
	}
	kinds, err := a.Store.CountSharedByKind(ctx, sqlc.CountSharedByKindParams{SourceKey: v.Shared, With: v.With, CampaignID: v.Campaign})
	for _, k := range kinds {
		v.Counts = append(v.Counts, sqlc.CountCompendiumRow(k))
	}
	if err != nil {
		return err
	}
	v.SharedWith, err = a.Store.ListSharedWith(ctx, sqlc.ListSharedWithParams{SourceKey: v.Shared, CampaignID: v.Campaign})
	return err
}

// homebrewLists fills the page's homebrew: all of it with its campaign
// links for the GM, a player's own entries for them.
func (a *App) homebrewLists(ctx context.Context, w who, v *views.CompendiumView) error {
	var err error
	if w.character != nil {
		v.MyHomebrew, err = a.Store.ListHomebrewByOwner(ctx, nullID(w.character.ID))
		return err
	}
	if v.Homebrew, err = a.Store.ListHomebrew(ctx); err != nil {
		return err
	}
	links, err := a.Store.ListHomebrewLinks(ctx)
	v.HomebrewLinks = map[int64][]sqlc.ListHomebrewLinksRow{}
	for _, l := range links {
		v.HomebrewLinks[l.EntryID] = append(v.HomebrewLinks[l.EntryID], l)
	}
	return err
}

// maxSuggestions keeps a text field's suggestion list short.
const maxSuggestions = 8

// compendiumSearch is the inline picker (or, with format=combobox, an
// itemPicker's options). Only the GM may use it for the
// combat tracker and NPC links; spells for a sheet need sheet rights,
// which the command endpoint checks again.
func (a *App) compendiumSearch(c *gin.Context) {
	w, camp, ok := a.compendiumReader(c)
	if !ok {
		return
	}
	var t views.PickerTarget
	t.EncounterID, _ = strconv.ParseInt(c.Query("encounter"), 10, 64)
	t.EntityID, _ = strconv.ParseInt(c.Query("entity"), 10, 64)
	t.CharacterID, _ = strconv.ParseInt(c.Query("character"), 10, 64)
	if !w.isGM() {
		t.EncounterID, t.EntityID = 0, 0
	}
	if strings.TrimSpace(c.Query("q")) == "" {
		c.Status(http.StatusOK) // cleared box: empty list, so the dropdown hides
		return
	}
	campaignID, err := a.pickerCampaign(c.Request.Context(), t, viewerCampaign(camp))
	if err != nil {
		a.serverError(c, err)
		return
	}
	names, text, err := compendium.Search(c.Request.Context(), a.Store.Queries, c.Query("q"), c.Query("kind"), "", campaignID, 15)
	if err != nil {
		a.serverError(c, err)
		return
	}
	// The picker's entry links open that campaign's compendium.
	c.Request = c.Request.WithContext(views.WithCampaign(c.Request.Context(), campaignID))
	if c.Query("format") == "combobox" {
		all := append(names, text...)
		a.render(c, http.StatusOK, views.SearchComboItems(all[:min(len(all), maxSuggestions)]))
		return
	}
	a.render(c, http.StatusOK, views.SearchResults(names, text, t))
}

// pickerCampaign is the campaign whose sources the GM's combat and NPC
// pickers offer: the encounter's or the NPC's (campaignID when neither).
func (a *App) pickerCampaign(ctx context.Context, t views.PickerTarget, campaignID int64) (int64, error) {
	switch {
	case t.EncounterID != 0:
		enc, err := a.Store.GetEncounter(ctx, t.EncounterID)
		if errors.Is(err, sql.ErrNoRows) {
			return campaignID, nil // the add itself reports the missing encounter
		}
		return enc.CampaignID, err
	case t.EntityID != 0:
		en, err := a.Store.GetEntity(ctx, t.EntityID)
		if errors.Is(err, sql.ErrNoRows) {
			return campaignID, nil
		}
		return en.CampaignID, err
	}
	return campaignID, nil
}

func (a *App) entryParam(c *gin.Context) (sqlc.CompendiumEntry, bool) {
	id, ok := a.idParam(c, "entry")
	if !ok {
		return sqlc.CompendiumEntry{}, false
	}
	e, err := a.Store.GetCompendiumEntry(c.Request.Context(), id)
	return e, a.found(c, err)
}

// compendiumFind redirects to the entry of a kind and slug the viewer would
// see, for links that know a slug but not an id (a weapon's mastery).
func (a *App) compendiumFind(c *gin.Context) {
	_, camp, ok := a.compendiumReader(c)
	if !ok {
		return
	}
	e, err := a.entryBySlug(c.Request.Context(), compendium.Kind(c.Param("kind")), c.Param("slug"), viewerCampaign(camp))
	if !a.found(c, err) {
		return
	}
	c.Redirect(http.StatusSeeOther, compendiumPath(c, fmt.Sprintf("/entries/%d", e.ID)))
}

func (a *App) compendiumEntry(c *gin.Context) {
	w, camp, ok := a.compendiumReader(c)
	if !ok {
		return
	}
	e, ok := a.entryParam(c)
	if !ok {
		return
	}
	if visible, err := a.entryVisible(c.Request.Context(), e, viewerCampaign(camp)); err != nil || !visible {
		a.found(c, cmp.Or(err, sql.ErrNoRows))
		return
	}
	data, err := a.decodeLinked(c.Request.Context(), e, viewerCampaign(camp))
	if err != nil {
		a.serverError(c, err)
		return
	}
	habitats, err := a.Store.ListEntryHabitats(c.Request.Context(), e.ID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	if c.Query("preview") != "" {
		a.render(c, http.StatusOK, views.EntryPreview(e, data, habitats))
		return
	}
	src, err := a.Store.GetCompendiumSource(c.Request.Context(), e.SourceKey)
	if err != nil {
		a.serverError(c, err)
		return
	}
	copies, err := a.entryCopies(c.Request.Context(), e, viewerCampaign(camp))
	if err != nil {
		a.serverError(c, err)
		return
	}
	v := views.EntryView{Entry: e, Source: src, Data: data, IsGM: w.isGM(), Habitats: habitats, Copies: copies, CanEdit: e.IsHomebrew && perm.CanEditHomebrew(w.viewer, e.OwnerCharacterID.Int64)}
	if v.IsGM && e.Kind == string(compendium.KindMonster) && len(habitats) == 0 {
		for _, cp := range copies {
			h, err := a.Store.ListEntryHabitats(c.Request.Context(), cp.ID)
			if err != nil {
				a.serverError(c, err)
				return
			}
			if len(h) > 0 {
				v.HabitatCopies = append(v.HabitatCopies, views.HabitatCopy{Copy: cp, Habitats: h})
			}
		}
	}
	// Set after a page import: back to the list, page tab open for the next page.
	if back, ok := c.GetQuery("back"); ok {
		q, _ := url.ParseQuery(compendiumBack(back))
		q.Set("import", views.ImportTabPage)
		v.Back = views.CompendiumURL(q, viewerCampaign(camp)) + "#compendium-import"
	} else if ref, err := url.Parse(c.Request.Referer()); err == nil && ref.Path == compendiumPath(c, "") {
		// Back to the list as it was: the search keeps its filters in the URL.
		// Only the path and query are kept, so the link never leaves the app.
		v.Back = ref.RequestURI()
	}
	l, err := a.compendiumLayout(c, w, e.Name, camp)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.EntryPage(l, v))
}

// compendiumCompare sets every copy of a kind and slug side by side, in
// priority order, for the GM to choose between sources: only the sources the
// campaign allows, and player homebrew it may see. The first is the one shown.
func (a *App) compendiumCompare(c *gin.Context) {
	w, camp, ok := a.compendiumReader(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	kind, slug := c.Param("kind"), c.Param("slug")
	if !compendium.ValidKind(kind) {
		a.found(c, sql.ErrNoRows)
		return
	}
	rows, err := a.Store.ListEntryCopies(ctx, sqlc.ListEntryCopiesParams{Kind: kind, Slug: slug})
	if err != nil {
		a.serverError(c, err)
		return
	}
	var v views.CompareView
	for _, r := range rows {
		allowed, err := sourceAllowed(ctx, a.Store.Queries, viewerCampaign(camp), r.SourceKey)
		if err != nil {
			a.serverError(c, err)
			return
		}
		if !allowed {
			continue
		}
		e, err := a.Store.GetCompendiumEntry(ctx, r.ID)
		if err != nil {
			a.serverError(c, err)
			return
		}
		visible, err := a.entryVisible(ctx, e, viewerCampaign(camp))
		if err != nil {
			a.serverError(c, err)
			return
		}
		if !visible {
			continue
		}
		data, err := a.decodeLinked(ctx, e, viewerCampaign(camp))
		if err != nil {
			a.serverError(c, err)
			return
		}
		v.Copies = append(v.Copies, views.CompareCopy{Entry: e, Source: r, Data: data})
	}
	if len(v.Copies) == 0 {
		a.found(c, sql.ErrNoRows)
		return
	}
	l, err := a.compendiumLayout(c, w, v.Copies[0].Entry.Name, camp)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.ComparePage(l, v))
}

// compendiumLookup is one entry's body found by kind and slug, as a
// character sheet links them (a spell's details when its row is opened), or
// by name, as an inventory names its items. A name that matches nothing
// gets an empty answer: most items are not in the compendium.
func (a *App) compendiumLookup(c *gin.Context) {
	_, camp, ok := a.compendiumReader(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	kind, name := c.Query("kind"), strings.TrimSpace(c.Query("name"))
	var e sqlc.CompendiumEntry
	var err error
	if name != "" {
		e, err = a.Store.GetCompendiumEntryByName(ctx, sqlc.GetCompendiumEntryByNameParams{Kind: kind, Name: name})
		if err == nil {
			if visible, verr := a.entryVisible(ctx, e, viewerCampaign(camp)); verr != nil || !visible {
				err = cmp.Or(verr, sql.ErrNoRows)
			}
		}
		if errors.Is(err, sql.ErrNoRows) {
			c.Status(http.StatusOK)
			return
		}
	} else {
		e, err = a.entryBySlug(ctx, compendium.Kind(kind), c.Query("slug"), viewerCampaign(camp))
		if errors.Is(err, sql.ErrNoRows) {
			c.String(http.StatusOK, "Not in the compendium.")
			return
		}
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	data, err := a.decodeLinked(ctx, e, viewerCampaign(camp))
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.EntryDetails(data))
}

// compendiumRandom jumps to a random entry of the chosen kind (or any kind).
func (a *App) compendiumRandom(c *gin.Context) {
	_, camp, ok := a.compendiumReader(c)
	if !ok {
		return
	}
	id, err := a.Store.RandomCompendiumEntry(c.Request.Context(), sqlc.RandomCompendiumEntryParams{Kind: c.Query("kind"), Habitat: c.Query("habitat"), CampaignID: viewerCampaign(camp)})
	if !a.found(c, err) {
		return
	}
	redirect(c, compendiumPath(c, fmt.Sprintf("/entries/%d", id)))
}

// monsterImageForm sets a monster's picture link; empty clears it for good.
type monsterImageForm struct {
	ImageURL string `form:"image_url" validate:"omitempty,url,startswith=https://,max=2048"`
}

// compendiumImage lets the GM replace or clear a monster's picture link.
func (a *App) compendiumImage(c *gin.Context) {
	e, ok := a.entryParam(c)
	if !ok {
		return
	}
	var f monsterImageForm
	if err := c.ShouldBind(&f); err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "The picture link must be an https:// URL.")
		return
	}
	n, err := a.Store.SetCompendiumImage(c.Request.Context(), sqlc.SetCompendiumImageParams{ImageUrl: sql.NullString{String: f.ImageURL, Valid: true}, ID: e.ID})
	if err != nil {
		a.serverError(c, err)
		return
	}
	if n == 0 {
		a.fail(c, http.StatusUnprocessableEntity, "Only monsters have pictures.")
		return
	}
	encounters, err := a.Store.ListEncountersByMonster(c.Request.Context(), nullID(e.ID))
	if err != nil {
		a.serverError(c, err)
		return
	}
	for _, enc := range encounters {
		a.Hub.Publish(enc.CampaignID, hub.Topic{Kind: hub.TopicEncounter, ID: enc.ID})
	}
	redirect(c, compendiumPath(c, fmt.Sprintf("/entries/%d", e.ID)))
}

// homebrewWriter is who may write homebrew: the GM, or a player (spells
// and items only, for their campaign).
func (a *App) homebrewWriter(c *gin.Context) (who, bool) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return w, false
	}
	if !w.isGM() && w.character == nil {
		a.fail(c, http.StatusForbidden, "Join a campaign first.")
		return w, false
	}
	return w, true
}

// playerKinds are the homebrew kinds players write.
var playerKinds = []string{string(compendium.KindSpell), string(compendium.KindItem)}

// homebrewEntry loads a homebrew entry the viewer may change.
func (a *App) homebrewEntry(c *gin.Context) (who, sqlc.CompendiumEntry, bool) {
	w, ok := a.homebrewWriter(c)
	if !ok {
		return w, sqlc.CompendiumEntry{}, false
	}
	e, ok := a.entryParam(c)
	if !ok {
		return w, e, false
	}
	switch {
	case !e.IsHomebrew:
		a.fail(c, http.StatusForbidden, "Imported entries are edited by re-importing their source.")
	case !perm.CanEditHomebrew(w.viewer, e.OwnerCharacterID.Int64):
		a.fail(c, http.StatusForbidden, "Only the GM and the player who wrote it may change this entry.")
	default:
		return w, e, true
	}
	return w, e, false
}

func (a *App) renderHomebrew(c *gin.Context, status int, w who, v views.HomebrewView) {
	v.IsGM = w.isGM()
	// No switcher: homebrew is the same in every campaign.
	a.render(c, status, views.HomebrewPage(a.layout(w, "Homebrew", views.NavCompendium, pathCampaign(c), false), v))
}

// homebrewNew is the empty editor. From a character sheet (?add=1 with a
// kind), the new entry is added to the character once saved.
func (a *App) homebrewNew(c *gin.Context) {
	w, ok := a.homebrewWriter(c)
	if !ok {
		return
	}
	kind := c.DefaultQuery("kind", string(compendium.KindMonster))
	if !w.isGM() && !slices.Contains(playerKinds, kind) {
		kind = string(compendium.KindSpell)
	}
	a.renderHomebrew(c, http.StatusOK, w, views.HomebrewView{Input: compendium.HomebrewInput{Kind: kind}, AddToSheet: c.Query("add") != ""})
}

func (a *App) homebrewEdit(c *gin.Context) {
	w, e, ok := a.homebrewEntry(c)
	if !ok {
		return
	}
	data, err := compendium.DecodeData(compendium.Kind(e.Kind), e.DataJson)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.renderHomebrew(c, http.StatusOK, w, views.HomebrewView{ID: e.ID, Input: compendium.HomebrewInputFor(e.Kind, e.Name, e.BodyText, data)})
}

func (a *App) homebrewCreate(c *gin.Context) {
	w, ok := a.homebrewWriter(c)
	if !ok {
		return
	}
	a.saveHomebrew(c, w, 0)
}

func (a *App) homebrewUpdate(c *gin.Context) {
	w, e, ok := a.homebrewEntry(c)
	if !ok {
		return
	}
	a.saveHomebrew(c, w, e.ID)
}

func (a *App) saveHomebrew(c *gin.Context, w who, id int64) {
	ctx := c.Request.Context()
	var in compendium.HomebrewInput
	err := c.ShouldBind(&in)
	var owner compendium.HomebrewOwner
	if w.character != nil {
		owner = compendium.HomebrewOwner{CharacterID: w.character.ID, CampaignID: w.character.CampaignID}
		if err == nil && !slices.Contains(playerKinds, in.Kind) {
			err = errors.New("players write homebrew spells and items")
		}
	}
	if err == nil {
		id, err = compendium.SaveHomebrew(ctx, a.Store.Queries, id, in, owner)
	}
	add := c.PostForm("add") != ""
	if err != nil {
		a.renderHomebrew(c, http.StatusUnprocessableEntity, w, views.HomebrewView{ID: id, Input: in, Error: err.Error(), AddToSheet: add})
		return
	}
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium}) // a homebrew copy may outrank a pin
	if add && w.character != nil {
		if err := a.addHomebrewToSheet(ctx, *w.character, id); a.commandFailed(c, err) {
			return
		}
		topics, err := a.characterTopics(ctx, w.character.ID)
		if err != nil {
			a.serverError(c, err)
			return
		}
		a.Hub.Publish(w.character.CampaignID, topics...)
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/characters/%d", w.character.ID))
		return
	}
	c.Redirect(http.StatusSeeOther, compendiumPath(c, fmt.Sprintf("/entries/%d", id)))
}

// addHomebrewToSheet puts a new homebrew spell or item on the character.
func (a *App) addHomebrewToSheet(ctx context.Context, ch sqlc.Character, id int64) error {
	e, err := a.Store.GetCompendiumEntry(ctx, id)
	if err != nil {
		return err
	}
	line := fmt.Sprintf("item add %s", strconv.Quote(e.Name))
	if e.Kind == string(compendium.KindSpell) {
		line = "spell add " + e.Slug
	}
	cmd, err := dnd5e.Parse(line)
	if err != nil {
		return err
	}
	return a.runCommand(ctx, ch, cmd)
}

func (a *App) homebrewDelete(c *gin.Context) {
	_, e, ok := a.homebrewEntry(c)
	if !ok {
		return
	}
	if err := a.Store.DeleteHomebrew(c.Request.Context(), e.ID); err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium})
	c.Redirect(http.StatusSeeOther, homebrewTabURL(c))
}

// homebrewTabURL is the request's compendium with its Homebrew tab open.
func homebrewTabURL(c *gin.Context) string {
	return compendiumPath(c, "?tab="+views.CompendiumTabHomebrew)
}

// homebrewLink shows a player's entry in another campaign too.
func (a *App) homebrewLink(c *gin.Context) {
	e, ok := a.entryParam(c)
	if !ok {
		return
	}
	campaignID, ok := formInt(c, "campaign_id")
	if !ok {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a campaign.")
		return
	}
	if err := a.Store.LinkEntryCampaign(c.Request.Context(), sqlc.LinkEntryCampaignParams{EntryID: e.ID, CampaignID: campaignID}); err != nil {
		a.serverError(c, err)
		return
	}
	redirect(c, homebrewTabURL(c))
}

// homebrewUnlink hides a player's entry from one campaign.
func (a *App) homebrewUnlink(c *gin.Context) {
	e, ok := a.entryParam(c)
	if !ok {
		return
	}
	campaignID, ok := a.idParam(c, "linked")
	if !ok {
		return
	}
	if err := a.Store.UnlinkEntryCampaign(c.Request.Context(), sqlc.UnlinkEntryCampaignParams{EntryID: e.ID, CampaignID: campaignID}); err != nil {
		a.serverError(c, err)
		return
	}
	redirect(c, homebrewTabURL(c))
}

// compendiumDir is the absolute DATA_DIR/compendium, shown to the GM.
func (a *App) compendiumDir() string {
	dir := filepath.Join(a.DataDir, "compendium")
	if abs, err := filepath.Abs(dir); err == nil {
		return abs
	}
	return dir
}

func (a *App) compendiumImport(c *gin.Context) {
	reports, err := compendium.Import(c.Request.Context(), a.Store, a.compendiumDir())
	if err != nil {
		a.serverError(c, err)
		return
	}
	msg, failed := ImportSummary(reports)
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium}) // global: no campaign
	a.renderCompendium(c, views.ImportNotice{Tab: views.ImportTabFolder, Msg: msg, Failed: failed})
}

// compendiumFolder opens DATA_DIR/compendium in the host's file manager. Only
// a browser on the host itself may ask: elsewhere it would open on a screen
// the GM is not looking at.
func (a *App) compendiumFolder(c *gin.Context) {
	if !auth.IsLoopback(c.Request) {
		a.fail(c, http.StatusForbidden, "The folder opens only from a browser on the host computer.")
		return
	}
	dir := a.compendiumDir()
	if err := os.MkdirAll(dir, 0o750); err != nil {
		a.serverError(c, err)
		return
	}
	opener := "xdg-open"
	switch runtime.GOOS {
	case "windows":
		opener = "explorer"
	case "darwin":
		opener = "open"
	}
	// explorer exits non-zero even when it opened the folder, so only a
	// failure to start counts; the file manager outlives the request.
	ctx := context.WithoutCancel(c.Request.Context())
	if err := exec.CommandContext(ctx, opener, dir).Start(); err != nil { //nolint:gosec // dir is the configured data directory, not request input
		a.serverError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// pageForm is a link to a page of one of compendium.Sites to preview.
type pageForm struct {
	Page string `form:"page" validate:"required,max=500"`
}

// compendiumPagePreview fetches one site page and shows it in the
// homebrew editor for review; nothing is saved yet.
func (a *App) compendiumPagePreview(c *gin.Context) {
	w, ok := a.homebrewWriter(c)
	if !ok {
		return
	}
	// The search keeps the page URL in step, so the referrer is the list
	// as filtered now; after an error the page carries it in a field.
	back := c.PostForm("back")
	if ref, err := url.Parse(c.Request.Referer()); err == nil && ref.Path == compendiumPath(c, "") {
		back = ref.RawQuery
	}
	back = compendiumBack(back)
	var f pageForm
	if err := c.ShouldBind(&f); err != nil {
		a.renderCompendium(c, views.ImportNotice{Tab: views.ImportTabPage, Back: back, Msg: "Paste a link to a page of one of the sites below.", Failed: true})
		return
	}
	p, err := compendium.FetchPage(c.Request.Context(), f.Page)
	if err != nil {
		a.renderCompendium(c, views.ImportNotice{Tab: views.ImportTabPage, Back: back, Page: f.Page, Msg: "Import failed: " + err.Error(), Failed: true})
		return
	}
	a.renderPagePreview(c, http.StatusOK, w, views.HomebrewView{Input: compendium.InputForEntry(p.Entry), Page: p, Back: back})
}

// compendiumBack keeps only the list filter (q, kind) of a query string, so
// a back link always stays on the compendium page.
func compendiumBack(query string) string {
	in, _ := url.ParseQuery(query)
	out := url.Values{}
	for _, k := range []string{"q", "kind"} {
		if v := in.Get(k); v != "" {
			out.Set(k, v)
		}
	}
	return out.Encode()
}

// renderPagePreview shows a page's review form, linking the entry an
// earlier import of the page made, which saving replaces.
func (a *App) renderPagePreview(c *gin.Context, status int, w who, v views.HomebrewView) {
	rows, err := a.Store.ListCompendiumKeysBySource(c.Request.Context(), compendium.PageSourceKey(v.Page.Site.Key, v.Page.ID))
	if err != nil {
		a.serverError(c, err)
		return
	}
	if len(rows) > 0 {
		v.Replaces = rows[0].ID
	}
	a.renderHomebrew(c, status, w, v)
}

// compendiumPageSave saves a reviewed site page and opens its entry.
func (a *App) compendiumPageSave(c *gin.Context) {
	w, ok := a.homebrewWriter(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	site, ok := compendium.SiteByKey(c.PostForm("site"))
	if !ok {
		c.String(http.StatusBadRequest, "unknown site")
		return
	}
	var in compendium.HomebrewInput
	err := c.ShouldBind(&in)
	p := compendium.Page{Site: site, ID: c.PostForm("page"), Book: c.PostForm("book"), Image: strings.TrimSpace(c.PostForm("image_url"))}
	back := compendiumBack(c.PostForm("back"))
	var key string
	if err == nil {
		if p.Entry, err = in.Entry(); err == nil {
			key, err = compendium.SavePage(ctx, a.Store, p)
		}
	}
	if err != nil {
		a.renderPagePreview(c, http.StatusUnprocessableEntity, w, views.HomebrewView{Input: in, Page: p, Back: back, Error: err.Error()})
		return
	}
	rows, err := a.Store.ListCompendiumKeysBySource(ctx, key)
	if err != nil || len(rows) == 0 {
		a.serverError(c, cmp.Or(err, errors.New("imported page has no entry")))
		return
	}
	redirect(c, compendiumPath(c, fmt.Sprintf("/entries/%d?", rows[0].ID))+url.Values{"back": {back}}.Encode())
}

// ImportSummary describes import results, one line per source, and
// reports whether any source failed.
func ImportSummary(reports []compendium.Report) (string, bool) {
	if len(reports) == 0 {
		return "No source folders found in the compendium folder.", true
	}
	var lines []string
	failed := false
	entries, removed := 0, 0
	for _, r := range reports {
		switch {
		case r.Err != nil:
			lines = append(lines, r.Source+": failed: "+r.Err.Error())
			failed = true
		case r.Unchanged:
			lines = append(lines, r.Source+": unchanged")
		default:
			lines = append(lines, fmt.Sprintf("%s: %d entries imported, %d removed", r.Source, r.Entries, r.Removed))
			entries, removed = entries+r.Entries, removed+r.Removed
		}
	}
	if len(reports) > 1 {
		lines = append([]string{fmt.Sprintf("%d sources: %d entries imported, %d removed in all", len(reports), entries, removed)}, lines...)
	}
	return strings.Join(lines, "\n"), failed
}
