package web

// The player app's side of the campaign starter file: opening one makes (or
// refreshes) a home campaign of the table's books, and a new character starts
// from it. A character made from a starter has a campaign of its own.

import (
	"cmp"
	"context"
	"database/sql"
	"errors"
	"io"
	"slices"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// homeStarterImport opens a starter file from the upload or the phone handoff.
func (a *App) homeStarterImport(c *gin.Context, f io.ReaderAt, size int64, fromPhone bool) error {
	file, err := export.ReadTakeHome(f, size, export.KindStarter)
	if err != nil {
		return userf("That is not a campaign starter file: %w", err)
	}
	if file.Manifest.Campaign.Name == "" {
		return userf("The campaign starter file has no campaign name.")
	}
	if file.Manifest.Campaign.System != gameSystem {
		return userf("This app plays %s only.", gameSystem)
	}
	if err := checkFileSources(file); err != nil {
		return err
	}
	ctx := c.Request.Context()
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error { return importStarter(ctx, q, file, time.Now().UTC()) })
	if err != nil {
		return err
	}
	if fromPhone {
		if closer, ok := f.(io.Closer); ok {
			_ = closer.Close() // Windows cannot delete an open file
		}
		a.stopHandoff()
	}
	redirect(c, route.AtHome)
	return nil
}

// importStarter writes a starter as a home campaign, or refreshes the one made
// from the same table campaign in place: name, books and entries follow the
// file, and the homebrew it links is the file's, none left over.
func importStarter(ctx context.Context, q *sqlc.Queries, file *export.TakeHomeFile, now time.Time) error {
	m := file.Manifest
	// One table per player app, checked again inside the transaction.
	if err := oneTable(ctx, q, m.ServerID); err != nil {
		return err
	}
	old, err := q.GetHomeStarterByTable(ctx, sqlc.GetHomeStarterByTableParams{ServerID: m.ServerID, TableCampaignID: m.Campaign.ID})
	switch {
	case err == nil:
		if err := q.RenameCampaign(ctx, sqlc.RenameCampaignParams{Name: m.Campaign.Name, ID: old.CampaignID}); err != nil {
			return err
		}
		if err := q.ClearCampaignSources(ctx, old.CampaignID); err != nil {
			return err
		}
		if err := q.ClearCampaignEntryLinks(ctx, old.CampaignID); err != nil {
			return err
		}
		if err := importHomeCompendium(ctx, q, file, old.CampaignID); err != nil {
			return err
		}
		if err := q.SetHomeStarterSavedAt(ctx, sqlc.SetHomeStarterSavedAtParams{SavedAt: now, CampaignID: old.CampaignID}); err != nil {
			return err
		}
	case errors.Is(err, sql.ErrNoRows):
		camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: m.Campaign.Name, System: m.Campaign.System, Edition: m.Campaign.Edition})
		if err != nil {
			return err
		}
		if err := q.SetCampaignOpenJoin(ctx, sqlc.SetCampaignOpenJoinParams{OpenJoin: false, ID: camp.ID}); err != nil {
			return err
		}
		if err := importHomeCompendium(ctx, q, file, camp.ID); err != nil {
			return err
		}
		if err := q.CreateHomeStarter(ctx, sqlc.CreateHomeStarterParams{CampaignID: camp.ID, ServerID: m.ServerID, TableCampaignID: m.Campaign.ID, SavedAt: now}); err != nil {
			return err
		}
	default:
		return err
	}
	// A refresh may have dropped homebrew the file no longer lists.
	return q.DeleteUnlinkedSharedHomebrew(ctx)
}

// homeStarterRemove deletes a starter campaign; characters made from it have
// campaigns of their own and stay.
func (a *App) homeStarterRemove(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		if _, err := q.GetHomeStarter(ctx, id); err != nil {
			return err
		}
		if err := q.DeleteCampaign(ctx, id); err != nil {
			return err
		}
		return q.DeleteUnlinkedSharedHomebrew(ctx)
	})
	if err != nil {
		return err
	}
	redirect(c, route.AtHome)
	return nil
}

// homeStarters are the starters for the At home page, each with its books in
// priority order.
func (a *App) homeStarters(ctx context.Context) ([]views.HomeStarterView, error) {
	rows, err := a.Store.ListHomeStarters(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]views.HomeStarterView, 0, len(rows))
	for _, r := range rows {
		sources, err := a.Store.ListTakeHomeSources(ctx, r.CampaignID)
		if err != nil {
			return nil, err
		}
		slices.SortStableFunc(sources, func(x, y sqlc.ListTakeHomeSourcesRow) int { return cmp.Compare(x.Priority, y.Priority) })
		v := views.HomeStarterView{CampaignID: r.CampaignID, Name: r.Name, SavedAt: r.SavedAt}
		for _, s := range sources {
			v.Books = append(v.Books, views.SourceLabel(s.SourceKey, s.Edition))
		}
		out = append(out, v)
	}
	return out, nil
}
