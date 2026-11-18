package web

import (
	"errors"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

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
	kind := enum.CompendiumKind(c.DefaultQuery(views.KindKey, string(enum.KindMonster)))
	if !w.isGM() && !slices.Contains(compendium.PlayerKinds, kind) {
		kind = enum.KindSpell
	}
	a.renderHomebrew(c, http.StatusOK, w, views.HomebrewView{Input: compendium.HomebrewInput{Kind: kind}, AddToSheet: c.Query("add") != ""})
}

func (a *App) homebrewEdit(c *gin.Context) error {
	w, e, ok := a.homebrewEntry(c)
	if !ok {
		return nil
	}
	data, err := compendium.DecodeData(e.Kind, e.DataJson)
	if err != nil {
		return err
	}
	a.renderHomebrew(c, http.StatusOK, w, views.HomebrewView{ID: e.ID, Input: compendium.HomebrewInputFor(e.Kind, e.Name, e.BodyText, data)})
	return nil
}

func (a *App) homebrewCreate(c *gin.Context) error {
	w, ok := a.homebrewWriter(c)
	if !ok {
		return nil
	}
	return a.saveHomebrew(c, w, 0)
}

func (a *App) homebrewUpdate(c *gin.Context) error {
	w, e, ok := a.homebrewEntry(c)
	if !ok {
		return nil
	}
	return a.saveHomebrew(c, w, e.ID)
}

func (a *App) saveHomebrew(c *gin.Context, w who, id int64) error {
	ctx := c.Request.Context()
	var in compendium.HomebrewInput
	err := c.ShouldBind(&in)
	var owner compendium.HomebrewOwner
	if w.character != nil {
		owner = compendium.HomebrewOwner{CharacterID: w.character.ID, CampaignID: w.character.CampaignID}
		if err == nil && !slices.Contains(compendium.PlayerKinds, in.Kind) {
			err = errors.New("players write homebrew spells and items")
		}
	}
	if err == nil {
		id, err = compendium.SaveHomebrew(ctx, a.Store.Queries, id, in, owner)
	}
	add := c.PostForm("add") != ""
	if err != nil {
		a.renderHomebrew(c, http.StatusUnprocessableEntity, w, views.HomebrewView{ID: id, Input: in, Error: err.Error(), AddToSheet: add})
		return nil
	}
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium}) // a homebrew copy may outrank a pin
	if add && w.character != nil {
		if err := a.addHomebrewToSheet(ctx, *w.character, id); err != nil {
			return err
		}
		topics, err := a.characterTopics(ctx, w.character.ID)
		if err != nil {
			return err
		}
		a.Hub.Publish(w.character.CampaignID, topics...)
		c.Redirect(http.StatusSeeOther, route.Path(route.Character, w.character.ID))
		return nil
	}
	c.Redirect(http.StatusSeeOther, compendiumPath(c, route.Path(route.CompendiumEntry, id)))
	return nil
}

func (a *App) homebrewDelete(c *gin.Context) error {
	_, e, ok := a.homebrewEntry(c)
	if !ok {
		return nil
	}
	if err := a.Store.DeleteHomebrew(c.Request.Context(), e.ID); err != nil {
		return err
	}
	a.Hub.Publish(0, hub.Topic{Kind: hub.TopicCompendium})
	c.Redirect(http.StatusSeeOther, homebrewTabURL(c))
	return nil
}

// homebrewTabURL is the request's compendium with its Homebrew tab open.
func homebrewTabURL(c *gin.Context) string {
	return compendiumPath(c, "?tab="+views.CompendiumTabHomebrew)
}

// homebrewLink shows a player's entry in another campaign too.
func (a *App) homebrewLink(c *gin.Context) error {
	e, ok := a.entryParam(c)
	if !ok {
		return nil
	}
	campaignID, ok := formInt(c, "campaign_id")
	if !ok {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a campaign.")
		return nil
	}
	if err := a.Store.LinkEntryCampaign(c.Request.Context(), sqlc.LinkEntryCampaignParams{EntryID: e.ID, CampaignID: campaignID}); err != nil {
		return err
	}
	redirect(c, homebrewTabURL(c))
	return nil
}

// homebrewUnlink hides a player's entry from one campaign.
func (a *App) homebrewUnlink(c *gin.Context) error {
	e, ok := a.entryParam(c)
	if !ok {
		return nil
	}
	campaignID, ok := a.idParam(c, "linked")
	if !ok {
		return nil
	}
	if err := a.Store.UnlinkEntryCampaign(c.Request.Context(), sqlc.UnlinkEntryCampaignParams{EntryID: e.ID, CampaignID: campaignID}); err != nil {
		return err
	}
	redirect(c, homebrewTabURL(c))
	return nil
}
