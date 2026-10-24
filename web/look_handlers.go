package web

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/db/sqlc"
	"gm-session-manager/perm"
)

// lookForm is a player's new look for their character: a link or a file
// (read by imageSource), and an optional name.
type lookForm struct {
	Name string `form:"name" validate:"max=200"`
	URL  string `form:"url"  validate:"omitempty,http_url,max=2000"`
}

// lookOps are one owner's look queries, for lookAction.
type lookOps struct {
	clearCurrent func(*sqlc.Queries) error
	setCurrent   func(*sqlc.Queries) (int64, error)
	toggle       func(*sqlc.Queries) error
	unlink       func(*sqlc.Queries) error
}

// lookAction makes a look current, toggles whether players see it, or
// unlinks it.
func (a *App) lookAction(ctx context.Context, action string, ops lookOps) error {
	switch action {
	case "current":
		return a.Store.InTx(ctx, func(q *sqlc.Queries) error {
			if err := ops.clearCurrent(q); err != nil {
				return err
			}
			n, err := ops.setCurrent(q)
			if err == nil && n == 0 {
				err = userError{errors.New("that image is not one of these looks")}
			}
			return err
		})
	case "reveal":
		return ops.toggle(a.Store.Queries)
	case "unlink":
		return ops.unlink(a.Store.Queries)
	}
	return userError{errors.New("unknown action")}
}

// lookFree refuses an image of another campaign, or one the other kind of
// owner wears: an image is an NPC's look or a character's, never both.
func lookFree(ctx context.Context, q *sqlc.Queries, ownerCampaignID, imageCampaignID, imageID int64, forCharacter bool) error {
	if ownerCampaignID != imageCampaignID {
		return userError{errors.New("the image belongs to another campaign")}
	}
	b, err := q.ImageBinding(ctx, imageID)
	switch {
	case err != nil:
		return err
	case forCharacter && b.Npcs > 0:
		return userError{errors.New("that image is an NPC's look")}
	case !forCharacter && b.Characters > 0:
		return userError{errors.New("that image is a character's look")}
	}
	return nil
}

// linkCharacterLook links an image to a character of the same campaign.
func linkCharacterLook(ctx context.Context, q *sqlc.Queries, ch sqlc.Character, imageCampaignID, imageID int64) error {
	if err := lookFree(ctx, q, ch.CampaignID, imageCampaignID, imageID, true); err != nil {
		return err
	}
	return q.LinkCharacterImage(ctx, sqlc.LinkCharacterImageParams{CharacterID: ch.ID, ImageID: imageID})
}

// lookOwner loads :character for its player, who alone manages its looks.
func (a *App) lookOwner(c *gin.Context) (sqlc.Character, bool) {
	w, ch, ok := a.characterAccess(c)
	if ok && !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player changes their character's looks.")
		return ch, false
	}
	return ch, ok
}

// addCharacterLook stores a player's upload or link as a campaign image and
// makes it a look of their character.
func (a *App) addCharacterLook(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxAttachmentBytes+1<<20)
	ch, ok := a.lookOwner(c)
	if !ok {
		return
	}
	var form lookForm
	if !a.bindForm(c, &form) {
		return
	}
	url, attachment, ok := a.imageSource(c, form.URL, sqlc.GetImageRow{CampaignID: ch.CampaignID})
	if !ok {
		return
	}
	name := strings.TrimSpace(form.Name)
	if name == "" {
		name = ch.Name
	}
	ctx := c.Request.Context()
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		img, err := q.CreateImage(ctx, sqlc.CreateImageParams{CampaignID: ch.CampaignID, Name: name, Url: url, AttachmentID: attachment})
		if err != nil {
			return err
		}
		return linkCharacterLook(ctx, q, ch, img.CampaignID, img.ID)
	})
	if a.commandFailed(c, err) {
		return
	}
	a.characterChanged(c, ch)
}

// characterLook makes one of a character's looks current, reveals or hides
// it, or unlinks it (the image stays among the campaign's images).
func (a *App) characterLook(c *gin.Context) {
	ch, ok := a.lookOwner(c)
	if !ok {
		return
	}
	imageID, ok := a.idParam(c, "image")
	if !ok {
		return
	}
	ctx := c.Request.Context()
	key := struct{ CharacterID, ImageID int64 }{ch.ID, imageID}
	err := a.lookAction(ctx, c.Param("action"), lookOps{
		clearCurrent: func(q *sqlc.Queries) error { return q.ClearCharacterCurrentImage(ctx, ch.ID) },
		setCurrent: func(q *sqlc.Queries) (int64, error) {
			return q.SetCharacterCurrentImage(ctx, sqlc.SetCharacterCurrentImageParams(key))
		},
		toggle: func(q *sqlc.Queries) error {
			return q.ToggleCharacterImageRevealed(ctx, sqlc.ToggleCharacterImageRevealedParams(key))
		},
		unlink: func(q *sqlc.Queries) error {
			return q.UnlinkCharacterImage(ctx, sqlc.UnlinkCharacterImageParams(key))
		},
	})
	if a.commandFailed(c, err) {
		return
	}
	a.characterChanged(c, ch)
}

// toLook converts a look row, an NPC's or a character's.
func toLook(id int64, name string, url sql.NullString, attachment sql.NullInt64, revealed, current bool) perm.Look {
	return perm.Look{ImageID: id, Name: name, Version: imageVersion(url, attachment), IsRevealed: revealed, IsCurrent: current}
}
