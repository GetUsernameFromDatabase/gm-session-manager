package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/hub"
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
	v.Msg, v.Failed, v.Running = msg, failed, a.fiveTools.get().Running
	return v, nil
}

func (a *App) renderFiveTools(c *gin.Context, msg string, failed bool) {
	v, err := a.fiveToolsView(c.Request.Context(), msg, failed)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.FiveToolsPicker(v))
}

// fiveToolsList is the picker, loaded when its tab is first shown. A plain
// visit (a reload after an import posted here) gets the whole page.
func (a *App) fiveToolsList(c *gin.Context) {
	if !isHTMX(c) {
		c.Redirect(http.StatusSeeOther, compendiumPath(c, "?import="+views.ImportTabFiveTools))
		return
	}
	a.renderFiveTools(c, "", false)
}

// fiveToolsForm is the checked books.
type fiveToolsForm struct {
	Books []string `form:"book" validate:"required,max=2000,dive,max=64"`
}

// fiveToolsImport starts downloading the checked books and importing every
// folder in the background; the import tab shows its progress.
func (a *App) fiveToolsImport(c *gin.Context) {
	ctx := c.Request.Context()
	notice := views.ImportNotice{Tab: views.ImportTabFiveTools}
	var f fiveToolsForm
	if c.ShouldBind(&f) != nil {
		notice.Msg, notice.Failed = "Pick at least one book.", true
		a.renderCompendium(c, notice)
		return
	}
	if a.fiveTools.get().Running { // before the book list's download; start checks again
		notice.Msg, notice.Failed = errImportRunning, true
		a.renderCompendium(c, notice)
		return
	}
	dir := a.compendiumDir()
	all, err := compendium.FiveToolsBooks(ctx, dir)
	if err != nil {
		notice.Msg, notice.Failed = "Could not load the book list: "+err.Error(), true
		a.renderCompendium(c, notice)
		return
	}
	var picked []compendium.Book
	for _, b := range all {
		if slices.ContainsFunc(f.Books, func(id string) bool { return strings.EqualFold(id, b.ID) }) {
			picked = append(picked, b)
		}
	}
	if !a.fiveTools.start(len(picked)) {
		notice.Msg, notice.Failed = errImportRunning, true
		a.renderCompendium(c, notice)
		return
	}
	// The import outlives the request.
	go a.runFiveTools(context.WithoutCancel(ctx), dir, picked)
	c.Redirect(http.StatusSeeOther, compendiumPath(c, "?import="+views.ImportTabFiveTools))
}

// autoRedownloadForm is the GM's "download again automatically" switch.
type autoRedownloadForm struct {
	On bool `form:"on"`
}

// fiveToolsAuto stores whether stale books download again at startup.
func (a *App) fiveToolsAuto(c *gin.Context) {
	var f autoRedownloadForm
	if c.ShouldBind(&f) != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if err := a.setAutoRedownload(f.On); err != nil {
		a.serverError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

// errImportRunning refuses a second import, or a removal under the running one.
const errImportRunning = "An import is running; wait for it to finish."

// fiveToolsRemove deletes a book's entries and source, then its folder.
func (a *App) fiveToolsRemove(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("book")
	if a.fiveTools.get().Running {
		a.renderFiveTools(c, errImportRunning, true)
		return
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
		if err := q.DeleteCompendiumEntriesBySource(ctx, key); err != nil {
			return err
		}
		return q.DeleteCompendiumSource(ctx, key)
	})
	var used *sourceUsedError
	if errors.As(err, &used) {
		a.renderFiveTools(c, sourceInUseMsg(key, used.users), true)
		return
	}
	if err == nil {
		err = compendium.RemoveFiveTools(a.compendiumDir(), id)
	}
	if err == nil {
		err = a.refreshStale()
	}
	if err != nil {
		a.renderFiveTools(c, "Could not remove "+id+": "+err.Error(), true)
		return
	}
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium})
	a.renderFiveTools(c, "Removed "+id+".", false)
}

// sourceMoveForm moves a source one place in the priority order.
type sourceMoveForm struct {
	Dir string `form:"dir" validate:"required,oneof=up down"`
}

// sourceMove reorders the priority list and answers with the card.
func (a *App) sourceMove(c *gin.Context) {
	var f sourceMoveForm
	if c.ShouldBind(&f) != nil {
		a.fail(c, http.StatusBadRequest, "Move a source up or down.")
		return
	}
	ctx := c.Request.Context()
	err := compendium.MoveSource(ctx, a.Store, c.Param("key"), f.Dir == views.SourceUp)
	if errors.Is(err, compendium.ErrNoSource) {
		a.fail(c, http.StatusNotFound, "That source is not imported.")
		return
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.sourcesChanged(c)
}

// sourceOrderForm is every source key in the new order (a drag).
type sourceOrderForm struct {
	Keys []string `form:"key" validate:"required,max=2000,dive,max=200"`
}

// sourceOrder sets the order a drag left and answers with the card.
func (a *App) sourceOrder(c *gin.Context) {
	var f sourceOrderForm
	if c.ShouldBind(&f) != nil {
		a.fail(c, http.StatusBadRequest, "Send every source in its new order.")
		return
	}
	err := compendium.SetSourceOrder(c.Request.Context(), a.Store, f.Keys)
	if errors.Is(err, compendium.ErrNoSource) {
		a.fail(c, http.StatusConflict, "The sources changed meanwhile: reload the page and drag again.")
		return
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.sourcesChanged(c)
}

// sourcesChanged tells every page and answers with the priority card.
func (a *App) sourcesChanged(c *gin.Context) {
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium})
	sources, shared, err := a.sourcePriority(c.Request.Context(), viewerCampaign(pathCampaign(c)))
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.SourcePriority(sources, shared))
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
		rows, err := a.Store.ListCompendiumBySlugs(ctx, sqlc.ListCompendiumBySlugsParams{Kind: kind, Slugs: slugs})
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
