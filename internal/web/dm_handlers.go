package web

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/web/views"
)

func (a *App) createCampaign(c *gin.Context) error {
	var form nameForm
	if !a.bindForm(c, &form) {
		return nil
	}
	ctx := c.Request.Context()
	var camp sqlc.Campaign
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		camp, err = q.CreateCampaign(ctx, sqlc.CreateCampaignParams{
			Name: strings.TrimSpace(form.Name), System: gameSystem, Edition: gameEdition,
		})
		if err != nil {
			return err
		}
		return q.AllowAllSources(ctx, camp.ID) // a new campaign starts with every source
	})
	if err != nil {
		return err
	}
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/campaigns/%d", camp.ID))
	return nil
}

func (a *App) campaignOverview(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	l := a.layout(w, camp.Name, views.NavHome, &camp, true)
	r, err := a.roster(c.Request.Context(), id)
	if err != nil {
		return err
	}
	s, err := a.campaignSourcesView(c.Request.Context(), id)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.CampaignOverview(l, r, a.joinInfo(id), s))
	return nil
}

func (a *App) toggleOpenJoin(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	_, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	if err := a.Store.SetCampaignOpenJoin(c.Request.Context(), sqlc.SetCampaignOpenJoinParams{ID: id, OpenJoin: !camp.OpenJoin}); err != nil {
		return err
	}
	a.changed(c, id, hub.Topic{Kind: hub.TopicCampaign, ID: id})
	return nil
}

func (a *App) renameCampaign(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	if _, _, ok := a.campaignAccess(c, id); !ok {
		return nil
	}
	var form nameForm
	if !a.bindForm(c, &form) {
		return nil
	}
	if err := a.Store.RenameCampaign(c.Request.Context(), sqlc.RenameCampaignParams{ID: id, Name: strings.TrimSpace(form.Name)}); err != nil {
		return err
	}
	redirect(c, fmt.Sprintf("/campaigns/%d", id))
	return nil
}

// deleteCampaign removes a campaign and, by cascade, everything in it. Its
// attachment files go at the next start (CleanAttachments).
func (a *App) deleteCampaign(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	if _, _, ok := a.campaignAccess(c, id); !ok {
		return nil
	}
	if err := a.Store.DeleteCampaign(c.Request.Context(), id); err != nil {
		return err
	}
	a.Hub.CloseCampaign(id)
	redirect(c, PathHome)
	return nil
}

// kick deletes a device's session, exactly like logout, and sends that
// device back to the start screen.
func (a *App) kick(c *gin.Context) error {
	id, ok := a.idParam(c, "session")
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	sess, err := a.Store.GetSession(ctx, id)
	if err != nil {
		return err
	}
	if sess.Role != enum.RolePlayer {
		a.fail(c, http.StatusForbidden, "Admin devices log out themselves.")
		return nil
	}
	ch, err := a.Store.GetCharacter(ctx, sess.CharacterID.Int64)
	if err != nil {
		return err
	}
	if err := a.Auth.Revoke(ctx, sess.ID); err != nil {
		return err
	}
	a.Hub.CloseSession(sess.ID)
	a.reqs.forget(sessionKey(sess))
	a.changed(c, ch.CampaignID, hub.Topic{Kind: hub.TopicCampaign, ID: ch.CampaignID})
	return nil
}

func (a *App) deleteCharacter(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	ch, err := a.Store.GetCharacterAny(ctx, id)
	if err != nil {
		return err
	}
	// Sessions, conditions, the player's notes and combat rows cascade.
	if err := a.Store.DeleteCharacter(ctx, id); err != nil {
		return err
	}
	a.Hub.CloseCharacter(id)
	a.changed(c, ch.CampaignID, hub.Topic{Kind: hub.TopicCampaign, ID: ch.CampaignID}, hub.Topic{Kind: hub.TopicNotes, ID: ch.CampaignID})
	return nil
}

// copyCharacter duplicates a character into another campaign. The copy has
// no sessions; its player asks to join it there.
func (a *App) copyCharacter(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	target, ok := formInt(c, "campaign_id")
	if !ok {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a campaign.")
		return nil
	}
	ctx := c.Request.Context()
	ch, err := a.Store.GetCharacter(ctx, id)
	if err != nil {
		return err
	}
	if _, err := a.Store.GetCampaign(ctx, target); err != nil {
		return err
	}
	if _, err := a.Store.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{
		CampaignID: target, Name: ch.Name, System: ch.System, Hp: ch.Hp, MaxHp: ch.MaxHp, TempHp: ch.TempHp, SheetJson: ch.SheetJson,
	}); err != nil {
		return err
	}
	a.Hub.Publish(target, hub.Topic{Kind: hub.TopicCampaign, ID: target})
	a.render(c, http.StatusOK, views.Flash("Copied "+ch.Name+".", false))
	return nil
}

func (a *App) exportCampaign(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	_, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s.zip"`, compendium.Slugify(camp.Name)))
	if err := export.Campaign(c.Request.Context(), c.Writer, a.Store.Queries, id, a.attachmentsDir()); err != nil {
		// Headers are gone; the truncated zip tells the browser it failed.
		a.serverError(c, err)
	}
	return nil
}

func (a *App) attachmentsDir() string { return filepath.Join(a.DataDir, "attachments") }
