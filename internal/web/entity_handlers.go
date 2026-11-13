package web

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// maxEntityFields bounds the editor form.
const maxEntityFields = 100

func (a *App) entitiesPage(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	l := a.layout(w, "NPCs", views.NavEntities, &camp, true)
	v, err := a.entityList(c.Request.Context(), w.viewer, id)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.EntitiesPage(l, v))
	return nil
}

func (a *App) createEntity(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	var form nameForm
	if !a.bindForm(c, &form) {
		return nil
	}
	e, err := a.Store.CreateEntity(c.Request.Context(), sqlc.CreateEntityParams{
		CampaignID: id, Kind: enum.EntityNPC, Name: strings.TrimSpace(form.Name), FieldsJson: "{}",
	})
	if err != nil {
		return err
	}
	a.Hub.Publish(id, hub.Topic{Kind: hub.TopicEntities, ID: id})
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/entities/%d", e.ID))
	return nil
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

func (a *App) entityPage(c *gin.Context) error {
	id, ok := a.idParam(c, "entity")
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	row, err := a.Store.GetEntity(ctx, id)
	if err != nil {
		return err
	}
	w, camp, ok := a.campaignAccess(c, row.CampaignID)
	if !ok {
		return nil
	}
	ev, visible, err := a.entityView(ctx, w.viewer, id)
	if err != nil {
		return err
	}
	if !visible {
		a.fail(c, http.StatusNotFound, "Not found.")
		return nil
	}
	if ev.Backlinks, err = a.backlinks(ctx, w, row.CampaignID, enum.LinkEntity, id); err != nil {
		return err
	}
	if ev.EntityLinks, err = a.entityMentions(ctx, w, enum.LinkEntity, id); err != nil {
		return err
	}
	l := a.layout(w, ev.E.Name, views.NavEntities, &camp, true)
	if w.isGM() {
		l.Modules = []string{"/assets/js/fields.js"}
	}
	a.render(c, http.StatusOK, views.EntityPage(l, ev))
	return nil
}

// updateEntity saves the editor form. Rows without a key, and suggestion
// rows left empty, are dropped.
func (a *App) updateEntity(c *gin.Context) error {
	e, ok := a.entityParam(c)
	if !ok {
		return nil
	}
	var form nameForm
	if !a.bindForm(c, &form) {
		return nil
	}
	count, _ := formInt(c, "count")
	fields := perm.Fields{}
	for i := range min(int(count), maxEntityFields) {
		key := strings.ToLower(strings.TrimSpace(c.PostForm(fmt.Sprintf("key_%d", i))))
		value := strings.TrimSpace(c.PostForm(fmt.Sprintf("value_%d", i)))
		if key == "" || value == "" {
			continue
		}
		fields[key] = perm.Field{Value: value, IsRevealed: c.PostForm(fmt.Sprintf("revealed_%d", i)) == strTrue}
	}
	raw, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	ctx := c.Request.Context()
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		if err := q.UpdateEntity(ctx, sqlc.UpdateEntityParams{
			ID: e.ID, Name: strings.TrimSpace(form.Name), IsRevealed: c.PostForm("revealed") == strTrue,
			FieldsJson: string(raw), CompendiumRef: e.CompendiumRef,
		}); err != nil {
			return err
		}
		return saveFieldLinks(ctx, q, e.ID, fields)
	})
	if err != nil {
		return err
	}
	// A toast, as the page flash is far above the Save bar.
	a.changedWith(c, e.CampaignID, views.Toast("Saved "+strings.TrimSpace(form.Name)+".", false), entityTopics(e)...)
	return nil
}

// saveFieldLinks replaces an NPC's field [[links]] with those in fields.
func saveFieldLinks(ctx context.Context, q *sqlc.Queries, entityID int64, fields perm.Fields) error {
	if err := q.DeleteEntityFieldLinks(ctx, entityID); err != nil {
		return err
	}
	for key, f := range fields {
		for _, l := range views.WikiLinks(f.Value) {
			if err := q.AddEntityFieldLink(ctx, sqlc.AddEntityFieldLinkParams{
				EntityID: entityID, FieldKey: key, TargetType: enum.LinkTarget(l.Kind), TargetID: l.ID, IsRevealed: f.IsRevealed,
			}); err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *App) entityChanged(c *gin.Context, e sqlc.Entity) {
	a.changed(c, e.CampaignID, entityTopics(e)...)
}

// entityTopics are the live views an NPC change touches: its page and the list.
func entityTopics(e sqlc.Entity) []hub.Topic {
	return []hub.Topic{{Kind: hub.TopicEntity, ID: e.ID}, {Kind: hub.TopicEntities, ID: e.CampaignID}}
}

func (a *App) deleteEntity(c *gin.Context) error {
	e, ok := a.entityParam(c)
	if !ok {
		return nil
	}
	if err := a.Store.DeleteEntity(c.Request.Context(), e.ID); err != nil {
		return err
	}
	a.Hub.Publish(e.CampaignID, entityTopics(e)...)
	redirect(c, fmt.Sprintf("/campaigns/%d/entities", e.CampaignID))
	return nil
}

// imageToEntity makes the image a look of the chosen character or NPC, or
// of a new NPC named after the image with its description as a hidden field.
func (a *App) imageToEntity(c *gin.Context) error {
	img, ok := a.imageParam(c)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	if id, ok := formInt(c, "character"); ok {
		ch, err := a.Store.GetCharacter(ctx, id)
		if err != nil {
			return err
		}
		if err := linkCharacterLook(ctx, a.Store.Queries, ch, img.CampaignID, img.ID); err != nil {
			return err
		}
		topics, err := a.characterTopics(ctx, ch.ID)
		if err != nil {
			return err
		}
		a.Hub.Publish(ch.CampaignID, topics...)
		redirect(c, fmt.Sprintf("/images/%d", img.ID))
		return nil
	}
	if id, ok := formInt(c, "entity"); ok {
		e, err := a.Store.GetEntity(ctx, id)
		if err != nil {
			return err
		}
		if err := linkEntityLook(ctx, a.Store.Queries, e, img.CampaignID, img.ID); err != nil {
			return err
		}
		a.Hub.Publish(e.CampaignID, hub.Topic{Kind: hub.TopicEntity, ID: e.ID})
		redirect(c, fmt.Sprintf("/images/%d", img.ID))
		return nil
	}
	fields := perm.Fields{}
	if d := strings.TrimSpace(img.Description.String); d != "" {
		fields["description"] = perm.Field{Value: d}
	}
	raw, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	var e sqlc.Entity
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if e, err = q.CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: img.CampaignID, Kind: enum.EntityNPC, Name: img.Name, FieldsJson: string(raw)}); err != nil {
			return err
		}
		return linkEntityLook(ctx, q, e, img.CampaignID, img.ID)
	})
	if err != nil {
		return err
	}
	a.Hub.Publish(img.CampaignID, hub.Topic{Kind: hub.TopicEntities, ID: img.CampaignID})
	redirect(c, fmt.Sprintf("/entities/%d", e.ID))
	return nil
}

// addEntityLook links an image of the campaign to the NPC.
func (a *App) addEntityLook(c *gin.Context) error {
	e, ok := a.entityParam(c)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	id, _ := formInt(c, "image")
	img, err := a.Store.GetImage(ctx, id)
	if err != nil {
		return err
	}
	if err := linkEntityLook(ctx, a.Store.Queries, e, img.CampaignID, img.ID); err != nil {
		return err
	}
	a.entityChanged(c, e)
	return nil
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
func (a *App) entityLook(c *gin.Context) error {
	e, ok := a.entityParam(c)
	if !ok {
		return nil
	}
	imageID, ok := a.idParam(c, "image")
	if !ok {
		return nil
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
	if err != nil {
		return err
	}
	a.entityChanged(c, e)
	return nil
}

// linkEntity sets or clears the compendium statblock behind an NPC.
func (a *App) linkEntity(c *gin.Context) error {
	e, ok := a.entityParam(c)
	if !ok {
		return nil
	}
	ref, _ := formInt(c, "entry")
	ctx := c.Request.Context()
	if ref != 0 {
		entry, err := a.Store.GetCompendiumEntry(ctx, ref)
		if err != nil {
			return err
		}
		if entry.Kind != enum.KindMonster {
			a.fail(c, http.StatusUnprocessableEntity, "Link a monster statblock.")
			return nil
		}
	}
	err := a.Store.UpdateEntity(ctx, sqlc.UpdateEntityParams{
		ID: e.ID, Name: e.Name, IsRevealed: e.IsRevealed, FieldsJson: e.FieldsJson, CompendiumRef: nullID(ref),
	})
	if err != nil {
		return err
	}
	a.entityChanged(c, e)
	return nil
}
