package web

// The table's side of the campaign starter file: a campaign's books and their
// entries with no character, for a player to start characters at home.

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
)

// campaignStarter downloads the starter file for the GM.
func (a *App) campaignStarter(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	_, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	return a.writeStarter(c, camp)
}

// joinStarter downloads the starter file for a visitor, only while the
// campaign's open join is on.
func (a *App) joinStarter(c *gin.Context) error {
	camp, ok := a.joinCampaign(c)
	if !ok {
		return nil
	}
	if !camp.OpenJoin {
		a.fail(c, http.StatusNotFound, "Not found.")
		return nil
	}
	return a.writeStarter(c, camp)
}

// writeStarter streams the campaign's starter zip. It holds what any player of
// the campaign sees: its books and entries (no monsters), and homebrew that is
// not a character's own.
func (a *App) writeStarter(c *gin.Context, camp sqlc.Campaign) error {
	ctx := c.Request.Context()
	serverID, err := a.Store.GetServerID(ctx)
	if err != nil {
		return err
	}
	var th export.TakeHome
	// No character id: takeHomeCompendium then leaves out every owned homebrew entry.
	if err := takeHomeCompendium(ctx, a.Store.Queries, sqlc.Character{CampaignID: camp.ID}, &th, nil); err != nil {
		return err
	}
	now := time.Now().UTC()
	th.Manifest = export.TakeHomeManifest{
		Format: export.TakeHomeFormat, Kind: export.KindStarter, AppVersion: a.Version, ServerID: serverID,
		Campaign: export.TakeHomeCampaign{ID: camp.ID, Name: camp.Name, System: camp.System, Edition: camp.Edition}, CreatedAt: now,
	}
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", `attachment; filename="`+export.ZipFileName(export.KindStarter, camp.Name, now)+`"`)
	if err := export.WriteTakeHome(c.Writer, th, a.attachmentsDir()); err != nil {
		// Headers are gone; the truncated zip tells the browser it failed.
		a.serverError(c, err)
	}
	return nil
}
