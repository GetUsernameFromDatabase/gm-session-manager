package web

import (
	"cmp"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// imageForm is the image editor. The file comes from the multipart body;
// a new image needs a link or a file.
type imageForm struct {
	Name        string `form:"name"        validate:"required,max=200"`
	URL         string `form:"url"         validate:"omitempty,http_url,max=2000"`
	Category    string `form:"category"    validate:"max=100"`
	Description string `form:"description" validate:"max=10000"`
	// What the editor opened with; WasSource is the picture's version.
	WasName        string `form:"was_name"`
	WasURL         string `form:"was_url"`
	WasCategory    string `form:"was_category"`
	WasDescription string `form:"was_description"`
	WasSource      string `form:"was_source"`
}

// mergeEditor keeps each stored value the editor did not change, so a page
// opened before another device's save does not undo it. A value both
// changed is a conflict, refused as a userError naming it. newFile says the
// editor uploaded a picture.
func (f *imageForm) mergeEditor(img sqlc.GetImageRow, newFile bool) error {
	var conflicts []string
	// Browsers post every line break as \r\n; a stored value may hold bare \n.
	norm := func(s string) string { return strings.ReplaceAll(strings.TrimSpace(s), "\r\n", "\n") }
	pick := func(label string, posted *string, loaded, stored string) {
		*posted = strings.TrimSpace(*posted)
		switch {
		case norm(*posted) == norm(loaded):
			*posted = stored
		case norm(loaded) != norm(stored):
			conflicts = append(conflicts, label)
		}
	}
	pick("name", &f.Name, f.WasName, img.Name)
	pick("category", &f.Category, f.WasCategory, img.Category)
	pick("description", &f.Description, f.WasDescription, img.Description)
	// Without a new link imageSource keeps the stored picture (or takes the
	// file). The upload tab posts no link at all.
	f.URL = strings.TrimSpace(f.URL)
	newLink := f.URL != "" && f.URL != strings.TrimSpace(f.WasURL)
	if !newLink {
		f.URL = ""
	}
	if (newLink || newFile) && img.ID != 0 && f.WasSource != imageVersion(img.Url, img.AttachmentID) {
		conflicts = append(conflicts, "picture")
	}
	if len(conflicts) > 0 {
		return userError{fmt.Errorf("another device changed the %s since this page opened: reload it to see the change, then edit again", strings.Join(conflicts, " and "))}
	}
	return nil
}

func imageItem(r sqlc.GetImageRow) views.ImageItem {
	return views.ImageItem{ID: r.ID, Name: r.Name, Category: r.Category, Description: r.Description, URL: r.Url.String, Version: imageVersion(r.Url, r.AttachmentID)}
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

func (a *App) imagesPage(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	v := views.ImagesView{CampaignID: id, Query: strings.TrimSpace(c.Query("q"))}
	v.CategoryID = formIntQuery(c, "category")
	v.Imported, _ = strconv.Atoi(c.Query("imported"))
	rows, err := a.Store.ListImages(ctx, sqlc.ListImagesParams{CampaignID: id, Q: v.Query, CategoryID: v.CategoryID})
	if err != nil {
		return err
	}
	gates, err := a.imageGates(ctx, id, 0)
	if err != nil {
		return err
	}
	chars, err := a.Store.ListCharactersByCampaign(ctx, id)
	if err != nil {
		return err
	}
	for _, r := range rows {
		item := imageItem(sqlc.GetImageRow(r))
		item.Seen = len(seenBy(gates[r.ID], chars)) > 0
		v.Images = append(v.Images, item)
	}
	if isHTMX(c) {
		a.render(c, http.StatusOK, views.ImageResults(v))
		return nil
	}
	if v.Categories, err = a.Store.ListImageCategories(ctx, id); err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.ImagesPage(a.layout(w, "Images", views.NavImages, &camp, false), v))
	return nil
}

func (a *App) newImagePage(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	cats, err := a.Store.ListImageCategories(c.Request.Context(), id)
	if err != nil {
		return err
	}
	v := views.ImageView{CampaignID: id, IsGM: true, Categories: cats}
	a.render(c, http.StatusOK, views.ImagePage(a.layout(w, "New image", views.NavImages, &camp, false), v))
	return nil
}

func (a *App) createImage(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	// A stale tab of a deleted campaign gets a 404, not a foreign key 500.
	if _, err := a.Store.GetCampaign(c.Request.Context(), id); err != nil {
		return err
	}
	return a.saveImage(c, sqlc.GetImageRow{CampaignID: id})
}

func (a *App) updateImage(c *gin.Context) error {
	img, ok := a.imageParam(c)
	if !ok {
		return nil
	}
	return a.saveImage(c, img)
}

// saveImage creates (img.ID 0) or updates an image from the editor. A new
// file replaces the link and the old file; a link replaces the old file.
func (a *App) saveImage(c *gin.Context, img sqlc.GetImageRow) error {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxAttachmentBytes+1<<20)
	var form imageForm
	if !a.bindForm(c, &form) {
		return nil
	}
	ctx := c.Request.Context()
	// Whoever wears it as a look shows the new picture and name; read
	// first, so a failure here never answers an error for a saved image.
	topics, err := a.imageTopics(ctx, img)
	if err != nil {
		return err
	}
	upload, ok := a.imageUpload(c, img.CampaignID)
	if !ok {
		return nil
	}
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		if img.ID != 0 {
			// Read again: another device may have saved while the upload came in.
			fresh, err := q.GetImage(ctx, img.ID)
			if err != nil {
				return err
			}
			img = fresh
		}
		if err := form.mergeEditor(img, upload.Valid); err != nil {
			return err
		}
		url, attachment, err := imageSource(form.URL, upload, img)
		if err != nil {
			return err
		}
		category, err := ensureCategory(ctx, q, img.CampaignID, form.Category)
		if err != nil {
			return err
		}
		name, desc := form.Name, form.Description
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
		if upload.Valid { // a refused save keeps no orphan upload
			a.dropUpload(ctx, upload.Int64)
		}
		return err
	}
	a.sweepFiles(c)
	a.Hub.Publish(img.CampaignID, topics...)
	redirect(c, route.Path(route.Image, img.ID))
	return nil
}

// imageTopics are the live views showing img as a look: its NPCs (their
// pages and the NPC list's portraits) and its PCs.
func (a *App) imageTopics(ctx context.Context, img sqlc.GetImageRow) ([]hub.Topic, error) {
	var topics []hub.Topic
	npcs, err := a.Store.ListImageEntities(ctx, img.ID)
	if err != nil {
		return nil, err
	}
	for _, e := range npcs {
		topics = append(topics, hub.Topic{Kind: hub.TopicEntity, ID: e.ID})
	}
	if len(npcs) > 0 {
		topics = append(topics, hub.Topic{Kind: hub.TopicEntities, ID: img.CampaignID})
	}
	pcs, err := a.Store.ListImageCharacters(ctx, img.ID)
	if err != nil {
		return nil, err
	}
	for _, ch := range pcs {
		more, err := a.characterTopics(ctx, ch.ID)
		if err != nil {
			return nil, err
		}
		topics = append(topics, more...)
	}
	return topics, nil
}

// imageUpload stores the editor's uploaded picture, if any (an invalid ID
// for none).
func (a *App) imageUpload(c *gin.Context, campaignID int64) (sql.NullInt64, bool) {
	fh, err := c.FormFile("file")
	if errors.Is(err, http.ErrMissingFile) {
		return sql.NullInt64{}, true
	}
	if err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "The file could not be read (files are limited to 20 MB).")
		return sql.NullInt64{}, false
	}
	att, ok := a.storeImageFile(c, campaignID, fh)
	return nullID(att.ID), ok
}

// imageSource is the link and file img gets from the editor (zero ID for a
// new image). A new file (upload) replaces the link and the old file; a link
// replaces the old file; neither keeps what img has.
func imageSource(link string, upload sql.NullInt64, img sqlc.GetImageRow) (sql.NullString, sql.NullInt64, error) {
	url, attachment := nullString(strings.TrimSpace(link)), img.AttachmentID
	switch {
	case upload.Valid:
		url, attachment = sql.NullString{}, upload
	case url.Valid:
		attachment = sql.NullInt64{}
	default:
		url = img.Url
	}
	if !url.Valid && !attachment.Valid {
		return url, attachment, userError{errors.New("give a link to the image or upload a file")}
	}
	return url, attachment, nil
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
		a.dropUpload(ctx, att.ID)
		a.fail(c, http.StatusUnprocessableEntity, "Upload a PNG, JPEG, GIF or WebP image.")
		return att, false
	}
	return att, true
}

// dropUpload deletes a just-stored upload nothing uses, row and file: it is
// too fresh for the sweep's upload guard.
func (a *App) dropUpload(ctx context.Context, id int64) {
	att, err := a.Store.GetAttachment(ctx, id)
	if err == nil {
		err = a.Store.DeleteAttachment(ctx, id)
	}
	if err == nil {
		err = os.Remove(filepath.Join(a.attachmentsDir(), att.Path))
	}
	if err != nil {
		slog.Error("remove refused upload", "err", err)
	}
}

// attachmentToImage copies a picture attached to a note block into its
// campaign's images (the GM's). The attachment stays, since other notes may
// still show it; the Files page deletes it once nothing does.
func (a *App) attachmentToImage(c *gin.Context) error {
	att, ok := a.readableAttachment(c)
	if !ok {
		return nil
	}
	if !slices.Contains(inlineTypes, att.Mime) {
		a.fail(c, http.StatusUnprocessableEntity, "Only pictures can become campaign images.")
		return nil
	}
	src, err := os.Open(filepath.Join(a.attachmentsDir(), filepath.Base(att.Path)))
	if err != nil {
		return err
	}
	defer func() { _ = src.Close() }()
	ctx := c.Request.Context()
	file, err := a.storeFile(ctx, att.CampaignID, sql.NullInt64{}, att.Name, src)
	if err != nil {
		return err
	}
	name := cmp.Or(strings.TrimSuffix(att.Name, filepath.Ext(att.Name)), att.Name, "Picture")
	img, err := a.Store.CreateImage(ctx, sqlc.CreateImageParams{CampaignID: att.CampaignID, Name: name, AttachmentID: nullID(file.ID)})
	if err != nil {
		return err
	}
	// The block links the image now, not at its next snapshot, so its readers can load it as soon as the link shows.
	if att.BlockID.Valid {
		b, _, err := a.loadBlock(ctx, att.BlockID.Int64)
		if err == nil {
			err = a.Store.AddBlockLink(ctx, sqlc.AddBlockLinkParams{SourceNoteID: b.NoteID, SourceBlockID: b.ID, TargetImageID: nullID(img.ID)})
		}
		if err != nil {
			return err
		}
	}
	c.JSON(http.StatusOK, gin.H{keyID: img.ID, keyName: img.Name})
	return nil
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

func (a *App) deleteImage(c *gin.Context) error {
	img, ok := a.imageParam(c)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	// Whoever wears it as a look loses it.
	topics, err := a.imageTopics(ctx, img)
	if err != nil {
		return err
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
		return err
	}
	a.sweepFiles(c)
	a.Hub.Publish(img.CampaignID, topics...)
	redirect(c, route.Path(route.CampaignImages, img.CampaignID))
	return nil
}

// imageGate is what perm.CanReadImage needs to know about one image.
type imageGate struct {
	look   perm.ImageLook
	blocks []perm.Block
}

// imageGates loads the gates of a campaign's images (one image, or every
// one for imageID 0) in three queries, however many blocks link them.
func (a *App) imageGates(ctx context.Context, campaignID, imageID int64) (map[int64]imageGate, error) {
	looks, err := a.Store.ListImageLookAccess(ctx, sqlc.ListImageLookAccessParams{CampaignID: campaignID, ImageID: imageID})
	if err != nil {
		return nil, err
	}
	links, err := a.Store.ListImageLinkBlocks(ctx, sqlc.ListImageLinkBlocksParams{CampaignID: campaignID, ImageID: imageID})
	if err != nil {
		return nil, err
	}
	audience, err := a.Store.ListImageLinkAudience(ctx, sqlc.ListImageLinkAudienceParams{CampaignID: campaignID, ImageID: imageID})
	if err != nil {
		return nil, err
	}
	gates := make(map[int64]imageGate, len(looks))
	for _, l := range looks {
		gates[l.ID] = imageGate{look: perm.ImageLook{Revealed: l.Revealed, Owner: l.OwnerCharacterID}}
	}
	members := map[int64][]int64{}
	for _, m := range audience {
		members[m.BlockID] = append(members[m.BlockID], m.CharacterID)
	}
	for _, l := range links {
		g := gates[l.ImageID]
		// The read rule needs only these three; edit fields stay empty.
		g.blocks = append(g.blocks, perm.Block{Visibility: l.Visibility, NoteOwnerID: l.OwnerCharacterID.Int64, Audience: members[l.BlockID]})
		gates[l.ImageID] = g
	}
	return gates, nil
}

// seenBy lists the characters who can see an image through its gate, for
// the GM: a player's own blocks are left out, as the GM never learns of them.
func seenBy(g imageGate, chars []sqlc.Character) []sqlc.Character {
	blocks := slices.DeleteFunc(slices.Clone(g.blocks), func(b perm.Block) bool { return b.Visibility == enum.VisOwner })
	return slices.DeleteFunc(slices.Clone(chars), func(ch sqlc.Character) bool {
		return !perm.CanReadImage(perm.Player(ch.ID), blocks, g.look)
	})
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
	var g imageGate
	if w.campaignID() == img.CampaignID {
		gates, err := a.imageGates(c.Request.Context(), img.CampaignID, img.ID)
		if err != nil {
			a.serverError(c, err)
			return w, img, false
		}
		g = gates[img.ID]
	}
	if !perm.CanReadImage(w.viewer, g.blocks, g.look) {
		a.fail(c, http.StatusNotFound, "Not found.")
		return w, img, false
	}
	return w, img, true
}

func (a *App) imagePage(c *gin.Context) error {
	w, img, ok := a.readableImage(c)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	camp, err := a.Store.GetCampaign(ctx, img.CampaignID)
	if err != nil {
		return err
	}
	v := views.ImageView{CampaignID: img.CampaignID, IsGM: w.isGM()}
	if !v.IsGM {
		// Players get the picture and its name, nothing else.
		v.Image = views.ImageItem{ID: img.ID, Name: img.Name, Version: imageVersion(img.Url, img.AttachmentID)}
		a.render(c, http.StatusOK, views.ImagePage(a.layout(w, img.Name, views.NavNotes, &camp, false), v))
		return nil
	}
	v.Image = imageItem(img)
	if v.Categories, err = a.Store.ListImageCategories(ctx, img.CampaignID); err != nil {
		return err
	}
	if v.Backlinks, err = a.backlinks(ctx, w, img.CampaignID, enum.LinkImage, img.ID); err != nil {
		return err
	}
	if v.EntityLinks, err = a.entityMentions(ctx, w, img.CampaignID, enum.LinkImage, img.ID); err != nil {
		return err
	}
	if v.NPCs, err = a.Store.ListImageEntities(ctx, img.ID); err != nil {
		return err
	}
	all, err := a.Store.ListEntitiesByCampaign(ctx, img.CampaignID)
	if err != nil {
		return err
	}
	v.NPCChoices = slices.DeleteFunc(all, func(e sqlc.Entity) bool {
		return slices.ContainsFunc(v.NPCs, func(n sqlc.ListImageEntitiesRow) bool { return n.ID == e.ID })
	})
	if v.Characters, err = a.Store.ListImageCharacters(ctx, img.ID); err != nil {
		return err
	}
	if v.CharacterChoices, err = a.Store.ListCharactersByCampaign(ctx, img.CampaignID); err != nil {
		return err
	}
	gates, err := a.imageGates(ctx, img.CampaignID, img.ID)
	if err != nil {
		return err
	}
	seen := seenBy(gates[img.ID], v.CharacterChoices)
	for _, ch := range seen {
		v.SeenBy = append(v.SeenBy, ch.Name)
	}
	v.SeenByAll = len(seen) > 0 && len(seen) == len(v.CharacterChoices)
	a.render(c, http.StatusOK, views.ImagePage(a.layout(w, img.Name, views.NavImages, &camp, false), v))
	return nil
}

// imageFile serves an uploaded image, or redirects to the external link,
// for anyone who may see the image.
func (a *App) imageFile(c *gin.Context) error {
	_, img, ok := a.readableImage(c)
	if !ok {
		return nil
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
		return nil
	}
	att, err := a.Store.GetAttachment(c.Request.Context(), img.AttachmentID.Int64)
	if err != nil {
		return err
	}
	cache := cacheAttachment
	if current {
		cache = cacheForever
	}
	a.serveAttachment(c, att, cache)
	return nil
}

func (a *App) displayPage(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	if _, err := a.Store.GetCampaign(c.Request.Context(), id); err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.DisplayPage(id))
	return nil
}
