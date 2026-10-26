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

	"gm-session-manager/auth"
	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

// compendiumReader allows the GM and any player: the compendium is rules
// text, shared by every campaign.
func (a *App) compendiumReader(c *gin.Context) (who, *sqlc.Campaign, bool) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return w, nil, false
	}
	if !w.isGM() && w.character == nil {
		a.fail(c, http.StatusForbidden, "Join a campaign first.")
		return w, nil, false
	}
	if w.character == nil {
		return w, nil, true
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), w.character.CampaignID)
	if !a.found(c, err) {
		return w, nil, false
	}
	return w, &camp, true
}

// viewerCampaign is the campaign whose player homebrew the viewer sees: the
// player's, or 0 for the GM, who sees it all.
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
	v := views.CompendiumView{Query: strings.TrimSpace(c.Query("q")), Kind: c.Query("kind"), IsGM: w.isGM(), Import: notice, Tab: c.Query("tab")}
	v.Import.Dir, v.Import.OnHost = a.compendiumDir(), w.id.Loopback
	// Only monsters have habitats, so the filter applies to them alone.
	if v.Kind == string(compendium.KindMonster) {
		v.Habitat = c.Query("habitat")
	}
	var err error
	if v.Query != "" {
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
			v.More = "/compendium?" + url.Values{"kind": {v.Kind}, "habitat": {v.Habitat}, compendiumAfterName: {last.Name}, compendiumAfterID: {strconv.FormatInt(last.ID, 10)}}.Encode()
		}
	}
	if c.Query(compendiumAfterID) != "" {
		a.render(c, http.StatusOK, views.CompendiumRows(v))
		return
	}
	if v.Counts, err = a.Store.CountCompendium(ctx); err != nil {
		a.serverError(c, err)
		return
	}
	if v.HabitatCounts, err = a.Store.CountHabitats(ctx); err != nil {
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
	a.render(c, http.StatusOK, views.CompendiumPage(a.layout(w, "Compendium", views.NavCompendium, camp, false), v))
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
	if err == nil {
		v.Campaigns, err = a.Store.ListCampaigns(ctx)
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
	names, text, err := compendium.Search(c.Request.Context(), a.Store.Queries, c.Query("q"), c.Query("kind"), "", viewerCampaign(camp), 15)
	if err != nil {
		a.serverError(c, err)
		return
	}
	if c.Query("format") == "combobox" {
		all := append(names, text...)
		a.render(c, http.StatusOK, views.SearchComboItems(all[:min(len(all), maxSuggestions)]))
		return
	}
	a.render(c, http.StatusOK, views.SearchResults(names, text, t))
}

func (a *App) entryParam(c *gin.Context) (sqlc.CompendiumEntry, bool) {
	id, ok := a.idParam(c, "entry")
	if !ok {
		return sqlc.CompendiumEntry{}, false
	}
	e, err := a.Store.GetCompendiumEntry(c.Request.Context(), id)
	return e, a.found(c, err)
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
	data, err := compendium.DecodeData(compendium.Kind(e.Kind), e.DataJson)
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
	v := views.EntryView{Entry: e, Source: src, Data: data, IsGM: w.isGM(), Habitats: habitats, CanEdit: e.IsHomebrew && perm.CanEditHomebrew(w.viewer, e.OwnerCharacterID.Int64)}
	// Set after a page import: back to the list, page tab open for the next page.
	if back, ok := c.GetQuery("back"); ok {
		q, _ := url.ParseQuery(compendiumBack(back))
		q.Set("import", views.ImportTabPage)
		v.Back = "/compendium?" + q.Encode() + "#compendium-import"
	} else if ref, err := url.Parse(c.Request.Referer()); err == nil && ref.Path == "/compendium" {
		// Back to the list as it was: the search keeps its filters in the URL.
		// Only the path and query are kept, so the link never leaves the app.
		v.Back = ref.RequestURI()
	}
	a.render(c, http.StatusOK, views.EntryPage(a.layout(w, e.Name, views.NavCompendium, camp, false), v))
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
	data, err := compendium.DecodeData(compendium.Kind(e.Kind), e.DataJson)
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
	redirect(c, fmt.Sprintf("/compendium/entries/%d", id))
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
	redirect(c, fmt.Sprintf("/compendium/entries/%d", e.ID))
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
	a.render(c, status, views.HomebrewPage(a.layout(w, "Homebrew", views.NavCompendium, nil, false), v))
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
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/compendium/entries/%d", id))
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
	c.Redirect(http.StatusSeeOther, homebrewTabURL)
}

// homebrewTabURL is the compendium with its Homebrew tab open.
const homebrewTabURL = "/compendium?tab=" + views.CompendiumTabHomebrew

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
	redirect(c, homebrewTabURL)
}

// homebrewUnlink hides a player's entry from one campaign.
func (a *App) homebrewUnlink(c *gin.Context) {
	e, ok := a.entryParam(c)
	if !ok {
		return
	}
	campaignID, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	if err := a.Store.UnlinkEntryCampaign(c.Request.Context(), sqlc.UnlinkEntryCampaignParams{EntryID: e.ID, CampaignID: campaignID}); err != nil {
		a.serverError(c, err)
		return
	}
	redirect(c, homebrewTabURL)
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
	if ref, err := url.Parse(c.Request.Referer()); err == nil && ref.Path == "/compendium" {
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
	redirect(c, fmt.Sprintf("/compendium/entries/%d?", rows[0].ID)+url.Values{"back": {back}}.Encode())
}

// ImportSummary describes import results, one line per source, and
// reports whether any source failed.
func ImportSummary(reports []compendium.Report) (string, bool) {
	if len(reports) == 0 {
		return "No source folders found in the compendium folder.", true
	}
	var lines []string
	failed := false
	for _, r := range reports {
		switch {
		case r.Err != nil:
			lines = append(lines, r.Source+": failed: "+r.Err.Error())
			failed = true
		case r.Unchanged:
			lines = append(lines, r.Source+": unchanged")
		default:
			lines = append(lines, fmt.Sprintf("%s: %d entries imported, %d removed", r.Source, r.Entries, r.Removed))
		}
	}
	return strings.Join(lines, "\n"), failed
}
