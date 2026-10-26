package web

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/auth"
	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/export"
	"gm-session-manager/hub"
	"gm-session-manager/web/views"
)

func (a *App) createCampaign(c *gin.Context) {
	var form nameForm
	if !a.bindForm(c, &form) {
		return
	}
	camp, err := a.Store.CreateCampaign(c.Request.Context(), sqlc.CreateCampaignParams{
		Name: strings.TrimSpace(form.Name), System: dnd5e.System, Edition: dnd5e.Edition,
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/campaigns/%d", camp.ID))
}

func (a *App) campaignOverview(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	l := a.layout(w, camp.Name, views.NavHome, &camp, true)
	r, err := a.roster(c.Request.Context(), id)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.CampaignOverview(l, r, a.joinInfo(id)))
}

func (a *App) toggleOpenJoin(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	_, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	if err := a.Store.SetCampaignOpenJoin(c.Request.Context(), sqlc.SetCampaignOpenJoinParams{ID: id, OpenJoin: !camp.OpenJoin}); err != nil {
		a.serverError(c, err)
		return
	}
	a.changed(c, id, hub.Topic{Kind: hub.TopicCampaign, ID: id})
}

func (a *App) renameCampaign(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	if _, _, ok := a.campaignAccess(c, id); !ok {
		return
	}
	var form nameForm
	if !a.bindForm(c, &form) {
		return
	}
	if err := a.Store.RenameCampaign(c.Request.Context(), sqlc.RenameCampaignParams{ID: id, Name: strings.TrimSpace(form.Name)}); err != nil {
		a.serverError(c, err)
		return
	}
	redirect(c, fmt.Sprintf("/campaigns/%d", id))
}

// deleteCampaign removes a campaign and, by cascade, everything in it. Its
// attachment files go at the next start (CleanAttachments).
func (a *App) deleteCampaign(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	if _, _, ok := a.campaignAccess(c, id); !ok {
		return
	}
	if err := a.Store.DeleteCampaign(c.Request.Context(), id); err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.CloseCampaign(id)
	redirect(c, PathHome)
}

// kick deletes a device's session, exactly like logout, and sends that
// device back to the start screen.
func (a *App) kick(c *gin.Context) {
	id, ok := a.idParam(c, "session")
	if !ok {
		return
	}
	ctx := c.Request.Context()
	sess, err := a.Store.GetSession(ctx, id)
	if !a.found(c, err) {
		return
	}
	if auth.Role(sess.Role) != auth.RolePlayer {
		a.fail(c, http.StatusForbidden, "Admin devices log out themselves.")
		return
	}
	ch, err := a.Store.GetCharacter(ctx, sess.CharacterID.Int64)
	if !a.found(c, err) {
		return
	}
	if err := a.Auth.Revoke(ctx, sess.ID); err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.CloseSession(sess.ID)
	a.changed(c, ch.CampaignID, hub.Topic{Kind: hub.TopicCampaign, ID: ch.CampaignID})
}

func (a *App) deleteCharacter(c *gin.Context) {
	id, ok := a.idParam(c, "character")
	if !ok {
		return
	}
	ctx := c.Request.Context()
	ch, err := a.Store.GetCharacterAny(ctx, id)
	if !a.found(c, err) {
		return
	}
	// Sessions, conditions, the player's notes and combat rows cascade.
	if err := a.Store.DeleteCharacter(ctx, id); err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.CloseCharacter(id)
	a.changed(c, ch.CampaignID, hub.Topic{Kind: hub.TopicCampaign, ID: ch.CampaignID}, hub.Topic{Kind: hub.TopicNotes, ID: ch.CampaignID})
}

// copyCharacter duplicates a character into another campaign. The copy has
// no sessions; its player asks to join it there.
func (a *App) copyCharacter(c *gin.Context) {
	id, ok := a.idParam(c, "character")
	if !ok {
		return
	}
	target, ok := formInt(c, "campaign_id")
	if !ok {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a campaign.")
		return
	}
	ctx := c.Request.Context()
	ch, err := a.Store.GetCharacter(ctx, id)
	if !a.found(c, err) {
		return
	}
	if _, err := a.Store.GetCampaign(ctx, target); !a.found(c, err) {
		return
	}
	if _, err := a.Store.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{
		CampaignID: target, Name: ch.Name, System: ch.System, Hp: ch.Hp, MaxHp: ch.MaxHp, TempHp: ch.TempHp, SheetJson: ch.SheetJson,
	}); err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(target, hub.Topic{Kind: hub.TopicCampaign, ID: target})
	a.render(c, http.StatusOK, views.Flash("Copied "+ch.Name+".", false))
}

func (a *App) exportCampaign(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	_, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.zip"`, compendium.Slugify(camp.Name)))
	if err := export.Campaign(c.Request.Context(), c.Writer, a.Store.Queries, id, a.attachmentsDir()); err != nil {
		// Headers are gone; the truncated zip tells the browser it failed.
		a.serverError(c, err)
	}
}

func (a *App) attachmentsDir() string { return filepath.Join(a.DataDir, "attachments") }
