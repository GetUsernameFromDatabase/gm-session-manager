package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

// Entity kinds (entities.kind). NPCs are the first kind.
const entityNPC = "npc"

// maxEntityFields bounds the editor form.
const maxEntityFields = 100

func (a *App) entitiesPage(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	l := a.layout(w, "NPCs", views.NavEntities, &camp, true)
	v, err := a.entityList(c.Request.Context(), w.viewer, id)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.EntitiesPage(l, v))
}

func (a *App) createEntity(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	var form nameForm
	if !a.bindForm(c, &form) {
		return
	}
	e, err := a.Store.CreateEntity(c.Request.Context(), sqlc.CreateEntityParams{
		CampaignID: id, Kind: entityNPC, Name: strings.TrimSpace(form.Name), FieldsJson: "{}",
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(id, hub.Topic{Kind: hub.TopicEntities, ID: id})
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/entities/%d", e.ID))
}

// entityParam loads :entity for the GM routes.
func (a *App) entityParam(c *gin.Context) (sqlc.Entity, bool) {
	id, ok := a.idParam(c, "entity")
	if !ok {
		return sqlc.Entity{}, false
	}
	e, err := a.Store.GetEntity(c.Request.Context(), id)
	return e, a.found(c, err)
}

func (a *App) entityPage(c *gin.Context) {
	id, ok := a.idParam(c, "entity")
	if !ok {
		return
	}
	ctx := c.Request.Context()
	row, err := a.Store.GetEntity(ctx, id)
	if !a.found(c, err) {
		return
	}
	w, camp, ok := a.campaignAccess(c, row.CampaignID)
	if !ok {
		return
	}
	ev, visible, err := a.entityView(ctx, w.viewer, id)
	if err != nil {
		a.serverError(c, err)
		return
	}
	if !visible {
		a.fail(c, http.StatusNotFound, "Not found.")
		return
	}
	if ev.Backlinks, err = a.backlinks(ctx, w, row.CampaignID, "entity", id); err != nil {
		a.serverError(c, err)
		return
	}
	if ev.EntityLinks, err = a.entityMentions(ctx, w, "entity", id); err != nil {
		a.serverError(c, err)
		return
	}
	l := a.layout(w, ev.E.Name, views.NavEntities, &camp, true)
	if w.isGM() {
		l.Modules = []string{"/assets/js/fields.js"}
	}
	a.render(c, http.StatusOK, views.EntityPage(l, ev))
}

// updateEntity saves the editor form. Rows without a key, and suggestion
// rows left empty, are dropped.
func (a *App) updateEntity(c *gin.Context) {
	e, ok := a.entityParam(c)
	if !ok {
		return
	}
	var form nameForm
	if !a.bindForm(c, &form) {
		return
	}
	count, _ := formInt(c, "count")
	fields := perm.Fields{}
	for i := range min(int(count), maxEntityFields) {
		key := strings.ToLower(strings.TrimSpace(c.PostForm(fmt.Sprintf("key_%d", i))))
		value := strings.TrimSpace(c.PostForm(fmt.Sprintf("value_%d", i)))
		if key == "" || value == "" {
			continue
		}
		fields[key] = perm.Field{Value: value, IsRevealed: c.PostForm(fmt.Sprintf("revealed_%d", i)) == "true"}
	}
	raw, err := json.Marshal(fields)
	if err != nil {
		a.serverError(c, err)
		return
	}
	ctx := c.Request.Context()
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		if err := q.UpdateEntity(ctx, sqlc.UpdateEntityParams{
			ID: e.ID, Name: strings.TrimSpace(form.Name), IsRevealed: c.PostForm("revealed") == "true",
			FieldsJson: string(raw), CompendiumRef: e.CompendiumRef,
		}); err != nil {
			return err
		}
		return saveFieldLinks(ctx, q, e.ID, fields)
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.entityChanged(c, e)
}

// saveFieldLinks replaces an NPC's field [[links]] with those in fields.
func saveFieldLinks(ctx context.Context, q *sqlc.Queries, entityID int64, fields perm.Fields) error {
	if err := q.DeleteEntityFieldLinks(ctx, entityID); err != nil {
		return err
	}
	for key, f := range fields {
		for _, l := range views.WikiLinks(f.Value) {
			if err := q.AddEntityFieldLink(ctx, sqlc.AddEntityFieldLinkParams{
				EntityID: entityID, FieldKey: key, TargetType: l.Kind, TargetID: l.ID, IsRevealed: f.IsRevealed,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *App) entityChanged(c *gin.Context, e sqlc.Entity) {
	a.changed(c, e.CampaignID, hub.Topic{Kind: hub.TopicEntity, ID: e.ID}, hub.Topic{Kind: hub.TopicEntities, ID: e.CampaignID})
}

func (a *App) deleteEntity(c *gin.Context) {
	e, ok := a.entityParam(c)
	if !ok {
		return
	}
	if err := a.Store.DeleteEntity(c.Request.Context(), e.ID); err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(e.CampaignID, hub.Topic{Kind: hub.TopicEntity, ID: e.ID}, hub.Topic{Kind: hub.TopicEntities, ID: e.CampaignID})
	redirect(c, fmt.Sprintf("/campaigns/%d/entities", e.CampaignID))
}

// imageToEntity makes the image a look of the chosen character or NPC, or
// of a new NPC named after the image with its description as a hidden field.
func (a *App) imageToEntity(c *gin.Context) {
	img, ok := a.imageParam(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	if id, ok := formInt(c, "character"); ok {
		ch, err := a.Store.GetCharacter(ctx, id)
		if !a.found(c, err) {
			return
		}
		if a.commandFailed(c, linkCharacterLook(ctx, a.Store.Queries, ch, img.CampaignID, img.ID)) {
			return
		}
		topics, err := a.characterTopics(ctx, ch.ID)
		if err != nil {
			a.serverError(c, err)
			return
		}
		a.Hub.Publish(ch.CampaignID, topics...)
		redirect(c, fmt.Sprintf("/images/%d", img.ID))
		return
	}
	if id, ok := formInt(c, "entity"); ok {
		e, err := a.Store.GetEntity(ctx, id)
		if !a.found(c, err) {
			return
		}
		if a.commandFailed(c, linkEntityLook(ctx, a.Store.Queries, e, img.CampaignID, img.ID)) {
			return
		}
		a.Hub.Publish(e.CampaignID, hub.Topic{Kind: hub.TopicEntity, ID: e.ID})
		redirect(c, fmt.Sprintf("/images/%d", img.ID))
		return
	}
	fields := perm.Fields{}
	if d := strings.TrimSpace(img.Description.String); d != "" {
		fields["description"] = perm.Field{Value: d}
	}
	raw, err := json.Marshal(fields)
	if err != nil {
		a.serverError(c, err)
		return
	}
	var e sqlc.Entity
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if e, err = q.CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: img.CampaignID, Kind: entityNPC, Name: img.Name, FieldsJson: string(raw)}); err != nil {
			return err
		}
		return linkEntityLook(ctx, q, e, img.CampaignID, img.ID)
	})
	if a.commandFailed(c, err) {
		return
	}
	a.Hub.Publish(img.CampaignID, hub.Topic{Kind: hub.TopicEntities, ID: img.CampaignID})
	redirect(c, fmt.Sprintf("/entities/%d", e.ID))
}

// addEntityLook links an image of the campaign to the NPC.
func (a *App) addEntityLook(c *gin.Context) {
	e, ok := a.entityParam(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	id, _ := formInt(c, "image")
	img, err := a.Store.GetImage(ctx, id)
	if !a.found(c, err) {
		return
	}
	if a.commandFailed(c, linkEntityLook(ctx, a.Store.Queries, e, img.CampaignID, img.ID)) {
		return
	}
	a.entityChanged(c, e)
}

// linkEntityLook links an image to an NPC of the same campaign.
func linkEntityLook(ctx context.Context, q *sqlc.Queries, e sqlc.Entity, imageCampaignID, imageID int64) error {
	if err := lookFree(ctx, q, e.CampaignID, imageCampaignID, imageID, false); err != nil {
		return err
	}
	return q.LinkEntityImage(ctx, sqlc.LinkEntityImageParams{EntityID: e.ID, ImageID: imageID})
}

// entityLook makes a look current, toggles whether players see it, or
// unlinks it.
func (a *App) entityLook(c *gin.Context) {
	e, ok := a.entityParam(c)
	if !ok {
		return
	}
	imageID, ok := a.idParam(c, "image")
	if !ok {
		return
	}
	ctx := c.Request.Context()
	key := struct{ EntityID, ImageID int64 }{e.ID, imageID}
	err := a.lookAction(ctx, c.Param("action"), lookOps{
		clearCurrent: func(q *sqlc.Queries) error { return q.ClearEntityCurrentImage(ctx, e.ID) },
		setCurrent: func(q *sqlc.Queries) (int64, error) {
			return q.SetEntityCurrentImage(ctx, sqlc.SetEntityCurrentImageParams(key))
		},
		toggle: func(q *sqlc.Queries) error {
			return q.ToggleEntityImageRevealed(ctx, sqlc.ToggleEntityImageRevealedParams(key))
		},
		unlink: func(q *sqlc.Queries) error { return q.UnlinkEntityImage(ctx, sqlc.UnlinkEntityImageParams(key)) },
	})
	if a.commandFailed(c, err) {
		return
	}
	a.entityChanged(c, e)
}

// linkEntity sets or clears the compendium statblock behind an NPC.
func (a *App) linkEntity(c *gin.Context) {
	e, ok := a.entityParam(c)
	if !ok {
		return
	}
	ref, _ := formInt(c, "entry")
	ctx := c.Request.Context()
	if ref != 0 {
		entry, err := a.Store.GetCompendiumEntry(ctx, ref)
		if !a.found(c, err) {
			return
		}
		if entry.Kind != string(compendium.KindMonster) {
			a.fail(c, http.StatusUnprocessableEntity, "Link a monster statblock.")
			return
		}
	}
	err := a.Store.UpdateEntity(ctx, sqlc.UpdateEntityParams{
		ID: e.ID, Name: e.Name, IsRevealed: e.IsRevealed, FieldsJson: e.FieldsJson, CompendiumRef: nullID(ref),
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.entityChanged(c, e)
}
