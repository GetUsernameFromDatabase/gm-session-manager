package web

import (
	"context"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"

	"gm-session-manager/db/sqlc"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

// attachmentRef is a use of an attachment in a note document: an attached
// picture ("attachment:N") or the src or href of an older picture or file link.
var attachmentRef = regexp.MustCompile(`"attachment:(\d+)"|/attachments/(\d+)"`)

// attachmentRefs collects the attachments the texts use.
func attachmentRefs(texts ...[]string) map[int64]bool {
	used := map[int64]bool{}
	for _, list := range texts {
		for _, s := range list {
			for _, m := range attachmentRef.FindAllStringSubmatch(s, -1) {
				if id, err := strconv.ParseInt(m[1]+m[2], 10, 64); err == nil {
					used[id] = true
				}
			}
		}
	}
	return used
}

// unusedAttachments lists the attachments no note block (saved or in an
// edit not yet saved) and no image uses, newest first.
func (a *App) unusedAttachments(ctx context.Context) ([]sqlc.ListAttachmentFilesRow, error) {
	docs, err := a.Store.ListBlockDocs(ctx)
	if err != nil {
		return nil, err
	}
	steps, err := a.Store.ListStepJSON(ctx)
	if err != nil {
		return nil, err
	}
	images, err := a.Store.ListImageAttachmentIDs(ctx)
	if err != nil {
		return nil, err
	}
	used := attachmentRefs(docs, steps)
	for _, id := range images {
		used[id.Int64] = true
	}
	files, err := a.Store.ListAttachmentFiles(ctx)
	return slices.DeleteFunc(files, func(f sqlc.ListAttachmentFilesRow) bool { return used[f.ID] }), err
}

func (a *App) filesPage(c *gin.Context) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return
	}
	files, err := a.unusedAttachments(c.Request.Context())
	if err != nil {
		a.serverError(c, err)
		return
	}
	v := views.FilesView{}
	for _, f := range files {
		v.Files = append(v.Files, views.FileRow{
			ID: f.ID, Name: f.Name, Campaign: f.CampaignName, Size: f.Size, CreatedAt: f.CreatedAt,
			IsImage: slices.Contains(inlineTypes, f.Mime),
			// The GM never reads a player's own blocks, so not even the name.
			Private: f.BlockVisibility.String == string(perm.VisOwner),
		})
	}
	a.render(c, http.StatusOK, views.FilesPage(a.layout(w, "Files", views.NavFiles, nil, false), v))
}

// deleteFile deletes one attachment, if nothing uses it any more.
func (a *App) deleteFile(c *gin.Context) {
	id, ok := a.idParam(c, "attachment")
	if !ok {
		return
	}
	a.deleteUnused(c, func(f sqlc.ListAttachmentFilesRow) bool { return f.ID == id })
}

func (a *App) deleteUnusedFiles(c *gin.Context) {
	a.deleteUnused(c, func(sqlc.ListAttachmentFilesRow) bool { return true })
}

// deleteUnused deletes the unused attachments that pick chooses, row and
// file, then reloads the Files page. A file used again since the page was
// shown is kept.
func (a *App) deleteUnused(c *gin.Context, pick func(sqlc.ListAttachmentFilesRow) bool) {
	ctx := c.Request.Context()
	files, err := a.unusedAttachments(ctx)
	if err != nil {
		a.serverError(c, err)
		return
	}
	for _, f := range files {
		if !pick(f) {
			continue
		}
		if err := a.Store.DeleteAttachment(ctx, f.ID); err != nil {
			a.serverError(c, err)
			return
		}
		if err := os.Remove(filepath.Join(a.attachmentsDir(), filepath.Base(f.Path))); err != nil && !errors.Is(err, os.ErrNotExist) {
			a.serverError(c, err)
			return
		}
	}
	redirect(c, views.PathFiles)
}
