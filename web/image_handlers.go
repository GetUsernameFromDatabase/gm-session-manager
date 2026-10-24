package web

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/db/sqlc"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

// linkImage is the note_links target type of images.
const linkImage = "image"

// imageForm is the image editor. The file comes from the multipart body;
// a new image needs a link or a file.
type imageForm struct {
	Name        string `form:"name"        validate:"required,max=200"`
	URL         string `form:"url"         validate:"omitempty,http_url,max=2000"`
	Category    string `form:"category"    validate:"max=100"`
	Description string `form:"description" validate:"max=10000"`
}

func imageItem(r sqlc.GetImageRow) views.ImageItem {
	return views.ImageItem{ID: r.ID, Name: r.Name, Category: r.Category, Description: r.Description.String, URL: r.Url.String, Version: imageVersion(r.Url, r.AttachmentID)}
}

// cacheForever is the Cache-Control of a versioned image file: a new link
// or file changes the version, so the URL itself changes.
const cacheForever = "private, max-age=31536000, immutable"

// imageVersion names an image's source (link or file) in its file URL. It is
// a hash, so it tells players nothing about the link.
func imageVersion(url sql.NullString, attachment sql.NullInt64) string {
	sum := sha256.Sum256(fmt.Appendf(nil, "%d|%s", attachment.Int64, url.String))
	return hex.EncodeToString(sum[:8])
}

func (a *App) imagesPage(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	v := views.ImagesView{CampaignID: id, Query: strings.TrimSpace(c.Query("q"))}
	v.CategoryID = formIntQuery(c, "category")
	v.Imported, _ = strconv.Atoi(c.Query("imported"))
	rows, err := a.Store.ListImages(ctx, sqlc.ListImagesParams{CampaignID: id, Q: v.Query, CategoryID: v.CategoryID})
	if err != nil {
		a.serverError(c, err)
		return
	}
	for _, r := range rows {
		v.Images = append(v.Images, imageItem(sqlc.GetImageRow(r)))
	}
	if isHTMX(c) {
		a.render(c, http.StatusOK, views.ImageResults(v))
		return
	}
	if v.Categories, err = a.Store.ListImageCategories(ctx, id); err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.ImagesPage(a.layout(w, "Images", views.NavImages, &camp, false), v))
}

func (a *App) newImagePage(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	cats, err := a.Store.ListImageCategories(c.Request.Context(), id)
	if err != nil {
		a.serverError(c, err)
		return
	}
	v := views.ImageView{CampaignID: id, IsGM: true, Categories: cats}
	a.render(c, http.StatusOK, views.ImagePage(a.layout(w, "New image", views.NavImages, &camp, false), v))
}

func (a *App) createImage(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	a.saveImage(c, sqlc.GetImageRow{CampaignID: id})
}

func (a *App) updateImage(c *gin.Context) {
	img, ok := a.imageParam(c)
	if !ok {
		return
	}
	a.saveImage(c, img)
}

// saveImage creates (img.ID 0) or updates an image from the editor. A new
// file replaces the link and the old file; a link replaces the old file.
func (a *App) saveImage(c *gin.Context, img sqlc.GetImageRow) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxAttachmentBytes+1<<20)
	var form imageForm
	if !a.bindForm(c, &form) {
		return
	}
	ctx := c.Request.Context()
	url, attachment, ok := a.imageSource(c, form.URL, img)
	if !ok {
		return
	}
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		category, err := ensureCategory(ctx, q, img.CampaignID, form.Category)
		if err != nil {
			return err
		}
		name, desc := strings.TrimSpace(form.Name), nullString(strings.TrimSpace(form.Description))
		if img.ID == 0 {
			created, err := q.CreateImage(ctx, sqlc.CreateImageParams{
				CampaignID: img.CampaignID, Name: name, Url: url, AttachmentID: attachment, CategoryID: category, Description: desc,
			})
			img.ID = created.ID
			return err
		}
		if err := q.UpdateImage(ctx, sqlc.UpdateImageParams{
			ID: img.ID, Name: name, Url: url, AttachmentID: attachment, CategoryID: category, Description: desc,
		}); err != nil {
			return err
		}
		if img.AttachmentID.Valid && img.AttachmentID != attachment {
			if err := q.DeleteAttachment(ctx, img.AttachmentID.Int64); err != nil {
				return err
			}
		}
		return q.DeleteUnusedImageCategories(ctx, img.CampaignID)
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/images/%d", img.ID))
}

// imageSource reads an image form's file or link for img (zero ID for a new
// image). A new file replaces the link and the old file; a link replaces
// the old file; neither keeps what img has.
func (a *App) imageSource(c *gin.Context, link string, img sqlc.GetImageRow) (sql.NullString, sql.NullInt64, bool) {
	url, attachment := nullString(strings.TrimSpace(link)), img.AttachmentID
	if fh, err := c.FormFile("file"); err == nil {
		att, ok := a.storeImageFile(c, img.CampaignID, fh)
		if !ok {
			return url, attachment, false
		}
		url, attachment = sql.NullString{}, nullID(att.ID)
	} else if !errors.Is(err, http.ErrMissingFile) {
		a.fail(c, http.StatusUnprocessableEntity, "The file could not be read (files are limited to 20 MB).")
		return url, attachment, false
	} else if url.Valid {
		attachment = sql.NullInt64{}
	} else {
		url = img.Url
	}
	if !url.Valid && !attachment.Valid {
		a.fail(c, http.StatusUnprocessableEntity, "Give a link to the image or upload a file.")
		return url, attachment, false
	}
	return url, attachment, true
}

// storeImageFile saves an uploaded image with no block, so only the GM and
// the image handlers can reach it. Anything but a picture is refused.
func (a *App) storeImageFile(c *gin.Context, campaignID int64, fh *multipart.FileHeader) (sqlc.Attachment, bool) {
	if fh.Size > maxAttachmentBytes {
		a.fail(c, http.StatusRequestEntityTooLarge, "Files are limited to 20 MB.")
		return sqlc.Attachment{}, false
	}
	ctx := c.Request.Context()
	att, err := a.storeUpload(ctx, campaignID, sql.NullInt64{}, fh)
	if err != nil {
		a.serverError(c, err)
		return att, false
	}
	if !slices.Contains(inlineTypes, att.Mime) {
		// The file is removed at the next startup with its row gone.
		if err := a.Store.DeleteAttachment(ctx, att.ID); err != nil {
			a.serverError(c, err)
			return att, false
		}
		a.fail(c, http.StatusUnprocessableEntity, "Upload a PNG, JPEG, GIF or WebP image.")
		return att, false
	}
	return att, true
}

// ensureCategory returns the category named name exactly, creating it,
// or no category for an empty name.
func ensureCategory(ctx context.Context, q *sqlc.Queries, campaignID int64, name string) (sql.NullInt64, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return sql.NullInt64{}, nil
	}
	id, err := q.EnsureImageCategory(ctx, sqlc.EnsureImageCategoryParams{CampaignID: campaignID, Name: name})
	return nullID(id), err
}

func nullString(s string) sql.NullString {
	return sql.NullString{String: s, Valid: s != ""}
}

// imageParam loads :image for the GM routes.
func (a *App) imageParam(c *gin.Context) (sqlc.GetImageRow, bool) {
	id, ok := a.idParam(c, "image")
	if !ok {
		return sqlc.GetImageRow{}, false
	}
	img, err := a.Store.GetImage(c.Request.Context(), id)
	return img, a.found(c, err)
}

func (a *App) deleteImage(c *gin.Context) {
	img, ok := a.imageParam(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	// Whoever wears it as a look loses it.
	var topics []hub.Topic
	npcs, err := a.Store.ListImageEntities(ctx, img.ID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	for _, e := range npcs {
		topics = append(topics, hub.Topic{Kind: hub.TopicEntity, ID: e.ID})
	}
	pcs, err := a.Store.ListImageCharacters(ctx, img.ID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	for _, ch := range pcs {
		more, err := a.characterTopics(ctx, ch.ID)
		if err != nil {
			a.serverError(c, err)
			return
		}
		topics = append(topics, more...)
	}
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		// The image goes first: its CHECK needs the link or the file.
		if err := q.DeleteImage(ctx, img.ID); err != nil {
			return err
		}
		if img.AttachmentID.Valid {
			if err := q.DeleteAttachment(ctx, img.AttachmentID.Int64); err != nil {
				return err
			}
		}
		return q.DeleteUnusedImageCategories(ctx, img.CampaignID)
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(img.CampaignID, topics...)
	redirect(c, fmt.Sprintf("/campaigns/%d/images", img.CampaignID))
}

// readableImage loads :image for anyone who may see it: the GM, or a
// player through a note block they can read that links it, or as a look
// (see perm.CanReadImage).
func (a *App) readableImage(c *gin.Context) (who, sqlc.GetImageRow, bool) {
	img, ok := a.imageParam(c)
	if !ok {
		return who{}, img, false
	}
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return w, img, false
	}
	if w.isGM() {
		return w, img, true
	}
	ctx := c.Request.Context()
	var blocks []perm.Block
	var look perm.ImageLook
	if w.campaignID() == img.CampaignID {
		access, err := a.Store.ImageLookAccess(ctx, img.ID)
		if err != nil {
			a.serverError(c, err)
			return w, img, false
		}
		look = perm.ImageLook{Revealed: access.Revealed, Owner: access.OwnerCharacterID}
		ids, err := a.Store.ListImageLinkBlocks(ctx, img.ID)
		if err != nil {
			a.serverError(c, err)
			return w, img, false
		}
		for _, id := range ids {
			_, pb, err := a.loadBlock(ctx, id)
			if err != nil {
				a.serverError(c, err)
				return w, img, false
			}
			blocks = append(blocks, pb)
		}
	}
	if !perm.CanReadImage(w.viewer, blocks, look) {
		a.fail(c, http.StatusNotFound, "Not found.")
		return w, img, false
	}
	return w, img, true
}

func (a *App) imagePage(c *gin.Context) {
	w, img, ok := a.readableImage(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	camp, err := a.Store.GetCampaign(ctx, img.CampaignID)
	if !a.found(c, err) {
		return
	}
	v := views.ImageView{CampaignID: img.CampaignID, IsGM: w.isGM()}
	if !v.IsGM {
		// Players get the picture and its name, nothing else.
		v.Image = views.ImageItem{ID: img.ID, Name: img.Name, Version: imageVersion(img.Url, img.AttachmentID)}
		a.render(c, http.StatusOK, views.ImagePage(a.layout(w, img.Name, views.NavNotes, &camp, false), v))
		return
	}
	v.Image = imageItem(img)
	if v.Categories, err = a.Store.ListImageCategories(ctx, img.CampaignID); err != nil {
		a.serverError(c, err)
		return
	}
	if v.Backlinks, err = a.backlinks(ctx, w, img.CampaignID, linkImage, img.ID); err != nil {
		a.serverError(c, err)
		return
	}
	if v.EntityLinks, err = a.entityMentions(ctx, w, linkImage, img.ID); err != nil {
		a.serverError(c, err)
		return
	}
	if v.NPCs, err = a.Store.ListImageEntities(ctx, img.ID); err != nil {
		a.serverError(c, err)
		return
	}
	all, err := a.Store.ListEntitiesByCampaign(ctx, img.CampaignID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	v.NPCChoices = slices.DeleteFunc(all, func(e sqlc.Entity) bool {
		return slices.ContainsFunc(v.NPCs, func(n sqlc.ListImageEntitiesRow) bool { return n.ID == e.ID })
	})
	if v.Characters, err = a.Store.ListImageCharacters(ctx, img.ID); err != nil {
		a.serverError(c, err)
		return
	}
	if v.CharacterChoices, err = a.Store.ListCharactersByCampaign(ctx, img.CampaignID); err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.ImagePage(a.layout(w, img.Name, views.NavImages, &camp, false), v))
}

// imageFile serves an uploaded image, or redirects to the external link,
// for anyone who may see the image.
func (a *App) imageFile(c *gin.Context) {
	_, img, ok := a.readableImage(c)
	if !ok {
		return
	}
	// Only the current version is cached for good; without it the GM may
	// change the link or file behind the same URL.
	current := c.Query("v") != "" && c.Query("v") == imageVersion(img.Url, img.AttachmentID)
	if !img.AttachmentID.Valid {
		cache := "no-store"
		if current {
			cache = cacheForever
		}
		c.Header("Cache-Control", cache)
		c.Redirect(http.StatusFound, img.Url.String)
		return
	}
	att, err := a.Store.GetAttachment(c.Request.Context(), img.AttachmentID.Int64)
	if !a.found(c, err) {
		return
	}
	cache := cacheAttachment
	if current {
		cache = cacheForever
	}
	a.serveAttachment(c, att, cache)
}

func (a *App) displayPage(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	if _, err := a.Store.GetCampaign(c.Request.Context(), id); !a.found(c, err) {
		return
	}
	a.render(c, http.StatusOK, views.DisplayPage(id))
}
