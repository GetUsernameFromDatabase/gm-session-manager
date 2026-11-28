package web

// The player app's "Save bring-back file": the home copy as a take-home zip of
// kind bring-back, with home ids and the id maps; the table translates them.

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"errors"
	"slices"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
)

// homeContent reads what a home character's file holds: the character, and of
// the compendium only what the table lacks (sources added at home that the home
// campaign allows, homebrew is never one, with their entries). A note still
// open in another tab refuses it.
func (a *App) homeContent(ctx context.Context, q *sqlc.Queries, ch sqlc.Character) (export.TakeHome, error) {
	th, _, err := buildTakeHome(ctx, q, ch)
	if err != nil {
		return th, err
	}
	keys, err := q.ListHomeSourceKeys(ctx)
	if err != nil {
		return th, err
	}
	if err := takeHomeCompendium(ctx, q, ch, &th, func(k string) bool { return slices.Contains(keys, k) }); err != nil {
		return th, err
	}
	return th, a.checkOpenNotes(ctx, q, ch.ID)
}

// checkOpenNotes refuses while one of the character's notes has steps after
// its last snapshot (open in another tab): a file or copy would miss them.
func (a *App) checkOpenNotes(ctx context.Context, q *sqlc.Queries, characterID int64) error {
	blocks, err := q.ListOwnNoteBlocks(ctx, nullID(characterID))
	if err != nil {
		return err
	}
	for _, b := range blocks {
		if b.DocVersion < b.Version {
			return a.openNoteError(ctx, q, characterID, b.NoteID)
		}
	}
	return nil
}

// homeHomebrew is homebrew rows as file entries, with their home ids.
func homeHomebrew(rows []sqlc.CompendiumEntry) []export.TakeHomeEntry {
	out := make([]export.TakeHomeEntry, len(rows))
	for i, e := range rows {
		out[i] = export.TakeHomeEntry{
			ID: e.ID, Kind: e.Kind, Slug: e.Slug, Name: e.Name, SourceKey: e.SourceKey, Data: json.RawMessage(e.DataJson),
			BodyText: e.BodyText, ContentHash: e.ContentHash, Scoped: e.Scoped,
		}
	}
	return out
}

// bringBack builds the bring-back file of a home copy and its file name, with
// its fallback build when one is prepared and not deleted.
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
		if th, err = a.homeContent(ctx, q, ch); err != nil {
			return err
		}
		if !hc.FallbackCharacterID.Valid {
			return nil
		}
		fb, err := q.GetCharacter(ctx, hc.FallbackCharacterID.Int64)
		if errors.Is(err, sql.ErrNoRows) {
			return nil // deleted: the file goes without it
		}
		if err != nil {
			return err
		}
		levelUps, err := q.ListLevelUps(ctx, fb.ID)
		if err != nil {
			return err
		}
		th.Fallback = &export.TakeHomeFallback{MaxHP: fb.MaxHp, Sheet: json.RawMessage(fb.SheetJson)}
		for _, l := range levelUps {
			th.Fallback.LevelUps = append(th.Fallback.LevelUps, export.TakeHomeLevelUp{ID: l.ID, MaxHP: l.MaxHp, Sheet: json.RawMessage(l.SheetJson), CreatedAt: l.CreatedAt.UTC()})
		}
		own, err := q.ListOwnHomebrew(ctx, nullID(fb.ID))
		if err != nil {
			return err
		}
		th.Homebrew = append(th.Homebrew, homeHomebrew(own)...)
		return nil
	})
	if err != nil {
		return th, "", err
	}
	now := time.Now().UTC()
	th.Manifest = m.TakeHomeManifest
	// The token did its job when this copy replaced the original; a bring-back must not carry it.
	th.Manifest.MadeAtHome = ""
	th.Manifest.Kind = export.KindBringBack
	th.Manifest.Format = export.TakeHomeFormat
	th.Manifest.AppVersion = a.Version
	th.Manifest.SavedAt = now
	th.Maps = &m.Maps
	return th, export.ZipFileName(export.KindBringBack, ch.Name, now), nil
}

// newCharacterFile builds the new-character file of a character that only
// exists on this computer (a test copy or one made here), and its file name.
// Its homebrew is whatever the home campaign holds, so the table matches it by
// content hash.
func (a *App) newCharacterFile(ctx context.Context, id int64) (export.TakeHome, string, error) {
	var th export.TakeHome
	ch, err := a.Store.GetCharacter(ctx, id)
	if err != nil {
		return th, "", err
	}
	if _, err := a.Store.GetHomeCopy(ctx, id); err == nil {
		return th, "", userf("%s goes back with a bring-back file, not as a new character.", ch.Name)
	} else if !errors.Is(err, sql.ErrNoRows) {
		return th, "", err
	}
	serverID, err := a.Store.GetHomeServerID(ctx)
	if errors.Is(err, sql.ErrNoRows) {
		return th, "", userf("This computer does not know its table yet: open a take-home file first.")
	}
	if err != nil {
		return th, "", err
	}
	now := time.Now().UTC()
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if th, err = a.homeContent(ctx, q, ch); err != nil {
			return err
		}
		camp, err := q.GetCampaign(ctx, ch.CampaignID)
		if err != nil {
			return err
		}
		th.Manifest.Campaign = export.TakeHomeCampaign{Name: camp.Name, System: camp.System, Edition: camp.Edition}
		rows, err := q.ListCampaignHomeHomebrew(ctx, ch.CampaignID)
		if err != nil {
			return err
		}
		th.Homebrew = homeHomebrew(rows)
		// The token ties this character to its table copy; the same one goes in every save.
		th.Manifest.MadeAtHome, err = q.GetNewCharacterToken(ctx, ch.ID)
		if errors.Is(err, sql.ErrNoRows) {
			th.Manifest.MadeAtHome = rand.Text()
			err = q.CreateNewCharacterToken(ctx, sqlc.CreateNewCharacterTokenParams{CharacterID: ch.ID, Token: th.Manifest.MadeAtHome})
		}
		return err
	})
	if err != nil {
		return th, "", err
	}
	th.Manifest.Format, th.Manifest.Kind, th.Manifest.AppVersion = export.TakeHomeFormat, export.KindNewCharacter, a.Version
	th.Manifest.ServerID, th.Manifest.SavedAt, th.Manifest.CreatedAt = serverID, now, now
	return th, export.ZipFileName(export.KindNewCharacter, ch.Name, now), nil
}

// homeFile builds the file a character on this computer goes back in: a
// bring-back for a table character, a new-character file for one that only
// exists here; a fallback goes inside its character's file.
func (a *App) homeFile(ctx context.Context, id int64) (export.TakeHome, string, error) {
	_, err := a.Store.GetHomeCopy(ctx, id)
	switch {
	case err == nil:
		return a.bringBack(ctx, id)
	case !errors.Is(err, sql.ErrNoRows):
		return export.TakeHome{}, "", err
	}
	owner, err := a.Store.GetFallbackOwner(ctx, nullID(id))
	if err == nil {
		return export.TakeHome{}, "", userf("A fallback build goes back inside %s's file.", owner.Name)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return export.TakeHome{}, "", err
	}
	return a.newCharacterFile(ctx, id)
}

// openNoteError names the note a block with unsnapshotted steps belongs to.
func (a *App) openNoteError(ctx context.Context, q *sqlc.Queries, characterID, noteID int64) error {
	notes, err := q.ListOwnNotes(ctx, nullID(characterID))
	if err != nil {
		return err
	}
	for _, n := range notes {
		if n.ID == noteID {
			return userf("‘%s’ is still open in another tab: close it, then try again.", n.Title)
		}
	}
	return userf("A note is still open in another tab: close it, then try again.")
}

// homeBringBackCheck answers the Save button, an htmx post: a refusal (a note
// still open) shows in its form, else the browser goes to the download.
func (a *App) homeBringBackCheck(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	// The route serves both file kinds (bring-back and new character).
	if _, _, err := a.homeFile(c.Request.Context(), id); err != nil {
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
	th, name, err := a.homeFile(c.Request.Context(), id)
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
