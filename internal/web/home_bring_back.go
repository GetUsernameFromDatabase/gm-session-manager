package web

// The player app's "Save bring-back file": the home copy as a take-home zip of
// kind bring-back, with home ids and the id maps; the table translates them.

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
)

// bringBack builds the bring-back file of a home copy and its file name.
func (a *App) bringBack(ctx context.Context, homeCharacterID int64) (export.TakeHome, string, error) {
	var th export.TakeHome
	hc, err := a.Store.GetHomeCopy(ctx, homeCharacterID)
	if err != nil {
		return th, "", err
	}
	var m export.HomeManifest
	if err := json.Unmarshal([]byte(hc.ManifestJson), &m); err != nil {
		return th, "", err
	}
	ch, err := a.Store.GetCharacter(ctx, homeCharacterID)
	if err != nil {
		return th, "", err
	}
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if th, _, err = buildTakeHome(ctx, q, ch); err != nil {
			return err
		}
		blocks, err := q.ListOwnNoteBlocks(ctx, nullID(ch.ID))
		if err != nil {
			return err
		}
		for _, b := range blocks {
			if b.DocVersion < b.Version {
				return a.openNoteError(ctx, q, ch.ID, b.NoteID)
			}
		}
		return nil
	})
	if err != nil {
		return th, "", err
	}
	now := time.Now().UTC()
	th.Manifest = m.TakeHomeManifest
	th.Manifest.Kind = export.KindBringBack
	th.Manifest.Format = export.TakeHomeFormat
	th.Manifest.AppVersion = a.Version
	th.Manifest.SavedAt = now
	th.Maps = &m.Maps
	return th, export.ZipFileName(export.KindBringBack, ch.Name, now), nil
}

// openNoteError names the note a block with unsnapshotted steps belongs to.
func (a *App) openNoteError(ctx context.Context, q *sqlc.Queries, characterID, noteID int64) error {
	notes, err := q.ListOwnNotes(ctx, nullID(characterID))
	if err != nil {
		return err
	}
	for _, n := range notes {
		if n.ID == noteID {
			return userf("‘%s’ is still open in another tab: close it, then save again.", n.Title)
		}
	}
	return userf("A note is still open in another tab: close it, then save again.")
}

// homeBringBackCheck answers the Save button, an htmx post: a refusal (a note
// still open) shows in its form, else the browser goes to the download.
func (a *App) homeBringBackCheck(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	if _, _, err := a.bringBack(c.Request.Context(), id); err != nil {
		return err
	}
	redirect(c, route.Path(route.AtHomeBringBack, id))
	return nil
}

// homeBringBack downloads the bring-back zip. Nothing is recorded at home:
// the copy stays playable, so a GET is safe.
func (a *App) homeBringBack(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	th, name, err := a.bringBack(c.Request.Context(), id)
	if err != nil {
		return err
	}
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", `attachment; filename="`+name+`"`)
	if err := export.WriteTakeHome(c.Writer, th, a.attachmentsDir()); err != nil {
		// Headers are gone; the truncated zip tells the browser it failed.
		a.serverError(c, err)
	}
	return nil
}
