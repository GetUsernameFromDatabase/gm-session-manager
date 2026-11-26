package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// fiveToolsKey is the source key of a book's folder.
func fiveToolsKey(id string) string { return compendium.FiveToolsPrefix + strings.ToLower(id) }

// fiveToolsView lists the mirror's books (the cache when offline), plus any
// downloaded book the list no longer names, with each one's import date.
func (a *App) fiveToolsView(ctx context.Context, msg string, failed bool) (views.FiveToolsView, error) {
	dir := a.compendiumDir()
	books, err := compendium.FiveToolsBooks(ctx, dir)
	if err != nil {
		msg, failed = "Could not load the book list: "+err.Error(), true
	}
	have, ierr := compendium.ImportedFiveTools(dir)
	if ierr != nil {
		return views.FiveToolsView{}, ierr
	}
	sources, serr := a.Store.ListCompendiumSources(ctx)
	if serr != nil {
		return views.FiveToolsView{}, serr
	}
	downloaded := map[string]time.Time{}
	for _, b := range have {
		downloaded[strings.ToLower(b.ID)] = time.Time{}
		if !slices.ContainsFunc(books, func(o compendium.Book) bool { return strings.EqualFold(o.ID, b.ID) }) {
			books = append(books, b)
		}
	}
	for _, s := range sources {
		if id, ok := strings.CutPrefix(s.SourceKey, compendium.FiveToolsPrefix); ok {
			if _, ok := downloaded[id]; ok {
				downloaded[id] = s.LastImported
			}
		}
	}
	v := views.NewFiveToolsView(books, downloaded)
	job := a.fiveTools.get()
	v.Msg, v.Failed, v.Running = msg, failed, job.Running
	// The books a running or failed import picked stay ticked.
	tickBooks(&v, func(b compendium.Book) bool {
		return slices.ContainsFunc(job.Books, func(id string) bool { return strings.EqualFold(id, b.ID) })
	})
	return v, nil
}

func (a *App) renderFiveTools(c *gin.Context, msg string, failed bool) error {
	v, err := a.fiveToolsView(c.Request.Context(), msg, failed)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.FiveToolsPicker(v))
	return nil
}

// fiveToolsListForm is the picker's query: a starting set to tick
// (views.FiveToolsSet2024, checked in TestSharedLimitsMatchTags) and the
// books already ticked in the browser, kept ticked.
type fiveToolsListForm struct {
	Set   string   `form:"set"  validate:"omitempty,oneof=2024"`
	Books []string `form:"book" validate:"max=2000,dive,max=64"`
}

// fiveToolsList is the picker, loaded when its tab is first shown; ?set=
// ticks a starting set's books too. A plain visit (a reload after an import
// posted here) gets the whole page.
func (a *App) fiveToolsList(c *gin.Context) error {
	if !isHTMX(c) {
		c.Redirect(http.StatusSeeOther, compendiumPath(c, "?"+views.ImportKey+"="+views.ImportTabFiveTools))
		return nil
	}
	var f fiveToolsListForm
	if c.ShouldBindQuery(&f) != nil {
		a.fail(c, http.StatusBadRequest, "No such set, or too many books.")
		return nil
	}
	v, err := a.fiveToolsView(c.Request.Context(), "", false)
	if err != nil {
		return err
	}
	if f.Set != "" {
		tickSet(&v)
	}
	tickBooks(&v, func(b compendium.Book) bool {
		return slices.ContainsFunc(f.Books, func(id string) bool { return strings.EqualFold(id, b.ID) })
	})
	a.render(c, http.StatusOK, views.FiveToolsPicker(v))
	return nil
}

// tickSet ticks the books of the 2024 set in the picker too.
func tickSet(v *views.FiveToolsView) {
	var all []compendium.Book
	for _, g := range v.Groups {
		for _, b := range g.Books {
			all = append(all, b.Book)
		}
	}
	set := compendium.FiveTools2024Set(all)
	tickBooks(v, func(b compendium.Book) bool { return slices.Contains(set, b) })
}

// tickBooks ticks the picker's books pick holds, keeping earlier ticks.
func tickBooks(v *views.FiveToolsView, pick func(compendium.Book) bool) {
	for _, g := range v.Groups {
		for i := range g.Books {
			g.Books[i].Ticked = g.Books[i].Ticked || pick(g.Books[i].Book)
		}
	}
}

// fiveToolsForm is the checked books, or a starting set (views.FiveToolsSet2024,
// checked in TestSharedLimitsMatchTags).
type fiveToolsForm struct {
	Books []string `form:"book" validate:"required_without=Set,max=2000,dive,max=64"`
	Set   string   `form:"set"  validate:"omitempty,oneof=2024"`
}

// fiveToolsImport starts downloading the checked books (or the set's) and
// importing every folder in the background. Ticked books come back to the
// import tab; a set, posted from the empty compendium, to the Browse page.
// Both show the progress.
func (a *App) fiveToolsImport(c *gin.Context) error {
	ctx := c.Request.Context()
	notice := views.ImportNotice{Tab: views.ImportTabFiveTools}
	var f fiveToolsForm
	if c.ShouldBind(&f) != nil {
		notice.Msg, notice.Failed = "Pick at least one book.", true
		return a.renderCompendium(c, notice)
	}
	if a.fiveTools.get().Running { // before the book list's download; start checks again
		notice.Msg, notice.Failed = errImportRunning, true
		return a.renderCompendium(c, notice)
	}
	dir := a.compendiumDir()
	all, err := compendium.FiveToolsBooks(ctx, dir)
	if err != nil {
		notice.Msg, notice.Failed = "Could not load the book list: "+err.Error(), true
		return a.renderCompendium(c, notice)
	}
	var picked []compendium.Book
	for _, b := range all {
		if slices.ContainsFunc(f.Books, func(id string) bool { return strings.EqualFold(id, b.ID) }) {
			picked = append(picked, b)
		}
	}
	back := fiveToolsBack(c) + "?" + views.ImportKey + "=" + views.ImportTabFiveTools
	if f.Set != "" {
		picked, back = compendium.FiveTools2024Set(all), fiveToolsBack(c)
		if len(picked) == 0 {
			notice.Msg, notice.Failed = "The 5etools mirror lists no 2024 books.", true
			return a.renderCompendium(c, notice)
		}
	}
	if len(picked) == 0 { // e.g. the stale-books banner's books have left the mirror
		notice.Msg, notice.Failed = "None of those books is on the 5etools mirror any more; remove them instead.", true
		return a.renderCompendium(c, notice)
	}
	if !a.fiveTools.start(len(picked)) {
		notice.Msg, notice.Failed = errImportRunning, true
		return a.renderCompendium(c, notice)
	}
	ids := make([]string, len(picked))
	for i, b := range picked {
		ids[i] = b.ID
	}
	a.fiveTools.update(func(j *views.FiveToolsJob) { j.Books = ids })
	// The import outlives the request; posted from a campaign's compendium
	// (its path, never the form or Referer), the books are ticked there.
	go a.runFiveTools(context.WithoutCancel(ctx), dir, picked, viewerCampaign(pathCampaign(c)))
	c.Redirect(http.StatusSeeOther, back)
	return nil
}

// open5eForm names the SRD (compendium.SRD2024Key or SRD2014Key) and, from
// the Data folder tab, that tab to come back to (views.ImportTabFolder);
// both checked in TestSharedLimitsMatchTags.
type open5eForm struct {
	Doc    string `form:"doc"    validate:"required,oneof=srd-2024 srd-2014"`
	Import string `form:"import" validate:"omitempty,oneof=folder"`
}

// open5eImport starts downloading an SRD from Open5e and importing every
// folder in the background; the page posted from shows its progress.
func (a *App) open5eImport(c *gin.Context) error {
	var f open5eForm
	if c.ShouldBind(&f) != nil {
		return a.renderCompendium(c, views.ImportNotice{Tab: views.ImportTabFolder, Msg: "Pick an SRD.", Failed: true})
	}
	if !a.fiveTools.start(0) { // the file count comes with the listing
		tab := views.ImportTabFiveTools
		if f.Import != "" { // posted from the Data folder tab: say it there
			tab = views.ImportTabFolder
		}
		return a.renderCompendium(c, views.ImportNotice{Tab: tab, Msg: errImportRunning, Failed: true})
	}
	// The import outlives the request; a campaign's compendium gets the SRD ticked.
	go a.runOpen5e(context.WithoutCancel(c.Request.Context()), a.compendiumDir(), f.Doc, viewerCampaign(pathCampaign(c)))
	back := fiveToolsBack(c)
	if f.Import != "" { // the Data folder tab, whose progress sits above the tabs
		back += "?" + views.ImportKey + "=" + views.ImportTabFolder
	}
	c.Redirect(http.StatusSeeOther, back)
	return nil
}

// fiveToolsBack is the compendium page the picker was opened from (the global
// one or a campaign's), since its form always posts to the global path, or
// the campaign page whose Sources card posted a starting set. Only the
// Referer's path is read, so the redirect stays on this site.
func fiveToolsBack(c *gin.Context) string {
	ref, err := url.Parse(c.Request.Referer())
	if err != nil {
		return compendiumPath(c, "")
	}
	if ref.Path == route.Compendium {
		return views.CompendiumBase(0)
	}
	p, ok := strings.CutPrefix(ref.Path, strings.TrimSuffix(route.Campaign, ":campaign"))
	idStr, rest, _ := strings.Cut(p, "/")
	if id, err := strconv.ParseInt(idStr, 10, 64); ok && err == nil && id > 0 {
		switch rest {
		case strings.TrimPrefix(route.Compendium, "/"):
			return views.CompendiumBase(id)
		case "":
			return route.Path(route.Campaign, id)
		}
	}
	return compendiumPath(c, "")
}

// autoRedownloadForm is the GM's "download again automatically" switch.
type autoRedownloadForm struct {
	On bool `form:"on"`
}

// fiveToolsAuto stores whether stale books download again at startup.
func (a *App) fiveToolsAuto(c *gin.Context) error {
	var f autoRedownloadForm
	if c.ShouldBind(&f) != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	if err := a.setAutoRedownload(f.On); err != nil {
		return err
	}
	c.Status(http.StatusNoContent)
	return nil
}

// errImportRunning refuses a second import, or a removal under the running one.
const errImportRunning = "An import is running; wait for it to finish."

// importDismiss clears the last import's summary from every GM page, unless
// it is no longer the one the page showed.
func (a *App) importDismiss(c *gin.Context) {
	if a.fiveTools.get().Running {
		a.fail(c, http.StatusConflict, errImportRunning)
		return
	}
	was := strings.ReplaceAll(c.PostForm(views.DismissWasField), "\r\n", "\n")
	if !a.fiveTools.dismiss(was) {
		a.fail(c, http.StatusConflict, "The import changed meanwhile; dismiss the summary shown now.")
		return
	}
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicImport})
	c.Status(http.StatusNoContent)
}

// fiveToolsRemove deletes a book's entries and source, then its folder.
func (a *App) fiveToolsRemove(c *gin.Context) error {
	ctx := c.Request.Context()
	id := c.Param("book")
	if a.fiveTools.get().Running {
		return a.renderFiveTools(c, errImportRunning, true)
	}
	key := fiveToolsKey(id)
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		users, err := sourceUsers(ctx, q, 0, key)
		if err != nil {
			return err
		}
		if len(users[key]) > 0 {
			return &sourceUsedError{users[key]}
		}
		return q.DeleteCompendiumSource(ctx, key)
	})
	if used, ok := errors.AsType[*sourceUsedError](err); ok {
		return a.renderFiveTools(c, sourceInUseMsg(key, used.users), true)
	}
	if err == nil {
		err = compendium.RemoveFiveTools(a.compendiumDir(), id)
	}
	if err == nil {
		err = a.refreshStale()
	}
	if err != nil {
		return a.renderFiveTools(c, "Could not remove "+id+": "+err.Error(), true)
	}
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium})
	return a.renderFiveTools(c, "Removed "+id+".", false)
}

// sourceOrderForm is every source key in the new order (a drag or a typed
// position).
type sourceOrderForm struct {
	Keys []string `form:"key" validate:"required,max=2000,dive,max=512"`
}

// sourceOrder sets the order a drag left and answers with the card.
func (a *App) sourceOrder(c *gin.Context) error {
	var f sourceOrderForm
	if c.ShouldBind(&f) != nil {
		a.fail(c, http.StatusBadRequest, "Send every source in its new order.")
		return nil
	}
	err := compendium.SetSourceOrder(c.Request.Context(), a.Store, f.Keys)
	if errors.Is(err, compendium.ErrNoSource) {
		a.fail(c, http.StatusConflict, "The sources changed meanwhile: reload the page and drag again.")
		return nil
	}
	if err != nil {
		return err
	}
	a.sourcesChanged(c)
	return nil
}

// sourcesChanged tells every page and answers with the priority card.
func (a *App) sourcesChanged(c *gin.Context) {
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium})
	sources, shared, counts, err := a.sourcesCard(c.Request.Context(), viewerCampaign(pathCampaign(c)))
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.SourcePriority(sources, shared, counts))
}

// sourcesCard is what the priority card of campaignID's compendium shows.
func (a *App) sourcesCard(ctx context.Context, campaignID int64) ([]sqlc.CompendiumSource, map[string]sqlc.CountSharedBySourceRow, map[string]sqlc.CountEntriesBySourceRow, error) {
	sources, shared, err := a.sourcePriority(ctx, campaignID)
	if err != nil {
		return nil, nil, nil, err
	}
	counts, err := a.entryCounts(ctx, campaignID)
	return sources, shared, counts, err
}

// entryCounts is, by source key, each source's entries and how many of them
// campaignID's lists show.
func (a *App) entryCounts(ctx context.Context, campaignID int64) (map[string]sqlc.CountEntriesBySourceRow, error) {
	rows, err := a.Store.CountEntriesBySource(ctx, campaignID)
	counts := map[string]sqlc.CountEntriesBySourceRow{}
	for _, r := range rows {
		counts[r.SourceKey] = r
	}
	return counts, err
}

// sourcePriority is the priority card's sources and, by source key, how many
// of each one's entries another source campaignID allows also holds.
func (a *App) sourcePriority(ctx context.Context, campaignID int64) ([]sqlc.CompendiumSource, map[string]sqlc.CountSharedBySourceRow, error) {
	sources, err := a.Store.ListCompendiumSources(ctx)
	if err != nil {
		return nil, nil, err
	}
	rows, err := a.Store.CountSharedBySource(ctx, campaignID)
	shared := map[string]sqlc.CountSharedBySourceRow{}
	for _, r := range rows {
		shared[r.SourceKey] = r
	}
	return sources, shared, err
}

// entryCopies are the other sources holding e's kind and slug, without
// player homebrew the viewer may not see.
func (a *App) entryCopies(ctx context.Context, e sqlc.CompendiumEntry, campaignID int64) ([]sqlc.ListEntryCopiesRow, error) {
	rows, err := a.Store.ListEntryCopies(ctx, sqlc.ListEntryCopiesParams{Kind: e.Kind, Slug: e.Slug})
	if err != nil {
		return nil, err
	}
	var out []sqlc.ListEntryCopiesRow
	for _, r := range rows {
		if r.ID == e.ID {
			continue
		}
		if r.SourceKey == compendium.HomebrewSource && campaignID != 0 {
			copyEntry, err := a.Store.GetCompendiumEntry(ctx, r.ID)
			if err != nil {
				return nil, err
			}
			ok, err := a.entryVisible(ctx, copyEntry, campaignID)
			if err != nil {
				return nil, err
			}
			if !ok {
				continue
			}
		}
		out = append(out, r)
	}
	return out, nil
}

// resolveCompendiumLinks turns the [[compendium:kind/slug|label]] tokens in
// an entry's JSON into Markdown links to the top-priority copy the viewer
// may see, or the bare label. One query per kind; import order never matters.
func (a *App) resolveCompendiumLinks(ctx context.Context, dataJSON string, campaignID int64) (string, error) {
	refs := views.CompendiumRefs(dataJSON)
	if len(refs) == 0 {
		return dataJSON, nil
	}
	ids := map[string]int64{}
	for kind, slugs := range refs {
		rows, err := a.Store.ListCompendiumBySlugs(ctx, sqlc.ListCompendiumBySlugsParams{Kind: enum.CompendiumKind(kind), Slugs: slugs})
		if err != nil {
			return "", err
		}
		for _, r := range rows {
			if r.Scoped && campaignID != 0 { // only player homebrew can be hidden
				e, err := a.Store.GetCompendiumEntry(ctx, r.ID)
				if err != nil {
					return "", err
				}
				ok, err := a.entryVisible(ctx, e, campaignID)
				if err != nil {
					return "", err
				}
				if !ok {
					continue
				}
			}
			ids[fmt.Sprintf("%s:%s/%s", views.LinkCompendium, kind, r.Slug)] = r.ID
		}
	}
	return views.ReplaceCompendiumLinks(dataJSON, ids, views.CompendiumBase(campaignID)), nil
}
