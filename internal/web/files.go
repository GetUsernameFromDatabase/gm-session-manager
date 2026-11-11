package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/web/views"
)

// attachmentRef is a use of an attachment in a note document: an attached
// picture ("attachment:N") or the src or href of an older picture or file link.
var attachmentRef = regexp.MustCompile(`"attachment:(\d+)"|/attachments/(\d+)"`)

// attachmentRefs lists the attachments a text uses.
func attachmentRefs(s string) []int64 {
	var ids []int64
	for _, m := range attachmentRef.FindAllStringSubmatch(s, -1) {
		if id, err := strconv.ParseInt(m[1]+m[2], 10, 64); err == nil {
			ids = append(ids, id)
		}
	}
	return ids
}

// attachmentUses maps each attachment to the notes (saved or in an edit not
// yet saved) and images that use it, each once.
func (a *App) attachmentUses(ctx context.Context) (map[int64][]views.FileUse, error) {
	blocks, err := a.Store.ListBlockDocs(ctx)
	if err != nil {
		return nil, err
	}
	steps, err := a.Store.ListStepJSON(ctx)
	if err != nil {
		return nil, err
	}
	images, err := a.Store.ListImageAttachments(ctx)
	if err != nil {
		return nil, err
	}
	type use struct {
		id int64
		u  views.FileUse
	}
	uses := map[int64][]views.FileUse{}
	seen := map[use]bool{}
	add := func(id int64, u views.FileUse) {
		if k := (use{id, u}); !seen[k] {
			seen[k] = true
			uses[id] = append(uses[id], u)
		}
	}
	byBlock := map[int64]views.FileUse{}
	for _, b := range blocks {
		u := views.FileUse{Kind: views.FileUseNote, URL: fmt.Sprintf("/notes/%d", b.NoteID), Label: b.Title}
		if b.Visibility == enum.VisOwner {
			// The GM never reads a player's own blocks, so not even the note's title.
			u = views.FileUse{Kind: views.FileUseNote, Private: true}
		}
		byBlock[b.ID] = u
		for _, id := range attachmentRefs(b.DocJson) {
			add(id, u)
		}
	}
	for _, s := range steps {
		for _, id := range attachmentRefs(s.StepJson) {
			add(id, byBlock[s.BlockID])
		}
	}
	for _, img := range images {
		add(img.AttachmentID.Int64, views.FileUse{Kind: views.FileUseImage, URL: fmt.Sprintf("/images/%d", img.ID), Label: img.Name})
	}
	return uses, nil
}

// unusedAttachments lists the attachments nothing uses, newest first.
func (a *App) unusedAttachments(ctx context.Context) ([]sqlc.ListAttachmentFilesRow, error) {
	uses, err := a.attachmentUses(ctx)
	if err != nil {
		return nil, err
	}
	files, err := a.Store.ListAttachmentFiles(ctx)
	return slices.DeleteFunc(files, func(f sqlc.ListAttachmentFilesRow) bool { return len(uses[f.ID]) > 0 }), err
}

// filesCampaign is the campaign in the path, nil for every campaign's files.
func (a *App) filesCampaign(c *gin.Context) (*sqlc.Campaign, bool) {
	if c.Param("campaign") == "" {
		return nil, true
	}
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil, false
	}
	_, camp, ok := a.campaignAccess(c, id)
	return &camp, ok
}

// inCampaign reports whether a file belongs on the page of camp (nil: all).
func inCampaign(camp *sqlc.Campaign, f sqlc.ListAttachmentFilesRow) bool {
	return camp == nil || f.CampaignID == camp.ID
}

func (a *App) filesPage(c *gin.Context) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return
	}
	camp, ok := a.filesCampaign(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	uses, err := a.attachmentUses(ctx)
	if err != nil {
		a.serverError(c, err)
		return
	}
	files, err := a.Store.ListAttachmentFiles(ctx)
	if err != nil {
		a.serverError(c, err)
		return
	}
	campaigns, err := a.Store.ListCampaigns(ctx)
	if err != nil {
		a.serverError(c, err)
		return
	}
	l := a.layout(w, "Files", views.NavFiles, camp, false)
	l.Switch = &views.CampaignSwitch{Campaigns: campaigns, Base: views.FilesBase, All: "All campaigns"}
	v := views.FilesView{Tab: c.Query("tab")}
	for _, f := range files {
		if !inCampaign(camp, f) {
			continue
		}
		v.Files = append(v.Files, views.FileRow{
			ID: f.ID, Name: f.Name, CampaignID: f.CampaignID, Campaign: f.CampaignName, Size: f.Size, CreatedAt: f.CreatedAt,
			IsImage: slices.Contains(inlineTypes, f.Mime),
			// The GM never reads a player's own blocks, so not even the name.
			Private: f.BlockVisibility == enum.VisOwner,
			Uses:    uses[f.ID],
		})
	}
	a.render(c, http.StatusOK, views.FilesPage(l, v))
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

// deleteUnused deletes the unused attachments of the path's campaign that
// pick chooses, row and file, then reloads the Files page. A file used again
// since the page was shown is kept.
func (a *App) deleteUnused(c *gin.Context, pick func(sqlc.ListAttachmentFilesRow) bool) {
	camp, ok := a.filesCampaign(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	files, err := a.unusedAttachments(ctx)
	if err != nil {
		a.serverError(c, err)
		return
	}
	for _, f := range files {
		if !inCampaign(camp, f) || !pick(f) {
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
	// Back to the tab the delete came from; Delete unused only shows in Unused.
	tab := views.FilesTabUnused
	if c.PostForm("tab") == views.FilesTabAll {
		tab = views.FilesTabAll
	}
	redirect(c, views.FilesBase(viewerCampaign(camp))+"?tab="+tab)
}
