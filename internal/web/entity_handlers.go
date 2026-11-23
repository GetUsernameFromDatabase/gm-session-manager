package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
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
	// A stale tab of a deleted campaign gets a 404, not a foreign key 500.
	if _, err := a.Store.GetCampaign(c.Request.Context(), id); err != nil {
		return err
	}
	e, err := a.Store.CreateEntity(c.Request.Context(), sqlc.CreateEntityParams{
		CampaignID: id, Kind: enum.EntityNPC, Name: strings.TrimSpace(form.Name), FieldsJson: "{}",
	})
	if err != nil {
		return err
	}
	a.Hub.Publish(id, hub.Topic{Kind: hub.TopicEntities, ID: id})
	c.Redirect(http.StatusSeeOther, route.Path(route.Entity, e.ID))
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
	if ev.EntityLinks, err = a.entityMentions(ctx, w, row.CampaignID, enum.LinkEntity, id); err != nil {
		return err
	}
	l := a.layout(w, ev.E.Name, views.NavEntities, &camp, true)
	if w.isGM() {
		l.Modules = []string{"/assets/js/fields.js"}
	}
	a.render(c, http.StatusOK, views.EntityPage(l, ev))
	return nil
}

// editEntity applies one edit command to an NPC, reading it afresh in the
// transaction, so a stale page on another GM device changes only what it
// sent. change gets the stored row and its fields to modify in place.
func (a *App) editEntity(c *gin.Context, change func(e *sqlc.Entity, fields perm.Fields) error) error {
	id, ok := a.idParam(c, "entity")
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	var row sqlc.Entity
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if row, err = q.GetEntity(ctx, id); err != nil {
			return err
		}
		fields, err := perm.ParseFields(row.FieldsJson)
		if err != nil {
			return err
		}
		if err := change(&row, fields); err != nil {
			return err
		}
		raw, err := json.Marshal(fields)
		if err != nil {
			return err
		}
		if err := q.UpdateEntity(ctx, sqlc.UpdateEntityParams{
			ID: row.ID, Name: row.Name, IsRevealed: row.IsRevealed, FieldsJson: string(raw), CompendiumRef: row.CompendiumRef,
		}); err != nil {
			return err
		}
		return saveFieldLinks(ctx, q, row.ID, fields)
	})
	if err != nil {
		return err
	}
	a.entityChanged(c, row)
	return nil
}

// entityFieldForm names one NPC field; value is read only by the text command.
type entityFieldForm struct {
	Key   string `form:"key"   validate:"required,max=100"`
	Value string `form:"value" validate:"max=100000"`
}

// fieldParam binds the field form, its key trimmed and lowercased.
func (a *App) fieldParam(c *gin.Context) (entityFieldForm, bool) {
	var form entityFieldForm
	if !a.bindForm(c, &form) {
		return form, false
	}
	form.Key = strings.ToLower(strings.TrimSpace(form.Key))
	if form.Key == "" {
		a.fail(c, http.StatusUnprocessableEntity, "Name the field.")
		return form, false
	}
	return form, true
}

// errFieldGone refuses an edit to a field another device removed, rather
// than bringing it back.
var errFieldGone = userError{errors.New("this field was removed on another device")}

func (a *App) renameEntity(c *gin.Context) error {
	var form nameForm
	if !a.bindForm(c, &form) {
		return nil
	}
	return a.editEntity(c, func(e *sqlc.Entity, _ perm.Fields) error {
		e.Name = strings.TrimSpace(form.Name)
		return nil
	})
}

func (a *App) revealEntity(c *gin.Context) error {
	revealed := c.PostForm("revealed") == strTrue
	return a.editEntity(c, func(e *sqlc.Entity, _ perm.Fields) error {
		e.IsRevealed = revealed
		return nil
	})
}

// addEntityField adds an empty, hidden field; an existing one stays as it is.
func (a *App) addEntityField(c *gin.Context) error {
	form, ok := a.fieldParam(c)
	if !ok {
		return nil
	}
	return a.editEntity(c, func(_ *sqlc.Entity, fields perm.Fields) error {
		if _, ok := fields[form.Key]; ok {
			return nil
		}
		if len(fields) >= maxEntityFields {
			return userError{fmt.Errorf("an NPC has at most %d fields", maxEntityFields)}
		}
		fields[form.Key] = perm.Field{}
		return nil
	})
}

func (a *App) setEntityFieldText(c *gin.Context) error {
	form, ok := a.fieldParam(c)
	if !ok {
		return nil
	}
	return a.editEntity(c, func(_ *sqlc.Entity, fields perm.Fields) error {
		f, ok := fields[form.Key]
		if !ok {
			return errFieldGone
		}
		f.Value = strings.TrimSpace(form.Value)
		fields[form.Key] = f
		return nil
	})
}

func (a *App) revealEntityField(c *gin.Context) error {
	form, ok := a.fieldParam(c)
	if !ok {
		return nil
	}
	revealed := c.PostForm("revealed") == strTrue
	return a.editEntity(c, func(_ *sqlc.Entity, fields perm.Fields) error {
		f, ok := fields[form.Key]
		if !ok {
			return errFieldGone
		}
		f.IsRevealed = revealed
		fields[form.Key] = f
		return nil
	})
}

func (a *App) removeEntityField(c *gin.Context) error {
	form, ok := a.fieldParam(c)
	if !ok {
		return nil
	}
	return a.editEntity(c, func(_ *sqlc.Entity, fields perm.Fields) error {
		delete(fields, form.Key)
		return nil
	})
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
	redirect(c, route.Path(route.CampaignEntities, e.CampaignID))
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
		redirect(c, route.Path(route.Image, img.ID))
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
		redirect(c, route.Path(route.Image, img.ID))
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
	redirect(c, route.Path(route.Entity, e.ID))
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
	ref, _ := formInt(c, "entry")
	if ref != 0 {
		entry, err := a.Store.GetCompendiumEntry(c.Request.Context(), ref)
		if err != nil {
			return err
		}
		if entry.Kind != enum.KindMonster {
			a.fail(c, http.StatusUnprocessableEntity, "Link a monster statblock.")
			return nil
		}
	}
	return a.editEntity(c, func(e *sqlc.Entity, _ perm.Fields) error {
		e.CompendiumRef = nullID(ref)
		return nil
	})
}
