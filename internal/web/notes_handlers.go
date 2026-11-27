package web

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// Limits for collaboration requests, uploads and searches.
const (
	maxCollabBody      = 4 << 20
	maxAttachmentBytes = 20 << 20
	maxReadableNotes   = 500
	maxSearchHits      = 50
	maxLinkSuggestions = 20
	// sniffLen is all http.DetectContentType looks at.
	sniffLen = 512
)

type noteForm struct {
	Title string        `form:"title" validate:"required,max=200"`
	Kind  enum.NoteKind `form:"kind"  validate:"required,oneof=game_session generic"`
}

func (a *App) notesPage(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	l := a.layout(w, "Notes", views.NavNotes, &camp, true)
	ctx := c.Request.Context()
	notes, err := a.noteList(ctx, w.viewer, id)
	if err != nil {
		return err
	}
	v := views.NoteListView{CampaignID: id, Notes: notes, Query: strings.TrimSpace(c.Query("q")), CanCreate: true}
	if v.Query != "" {
		v.Hits, err = a.Store.SearchReadableBlocks(ctx, sqlc.SearchReadableBlocksParams{
			CampaignID: id, Query: v.Query, IsGm: w.isGM(), CharacterID: w.viewer.CharacterID, MaxResults: maxSearchHits,
		})
		if err != nil {
			return err
		}
	}
	a.render(c, http.StatusOK, views.NotesPage(l, v))
	return nil
}

// createNote makes a note with one block: GM-only for the GM, private for
// a player. Visibility can be changed per block afterwards.
func (a *App) createNote(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, _, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	var form noteForm
	if !a.bindForm(c, &form) {
		return nil
	}
	ctx := c.Request.Context()
	owner, vis := int64(0), enum.VisGM
	if !w.isGM() {
		owner, vis = w.viewer.CharacterID, enum.VisOwner
	}
	var note sqlc.Note
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		note, err = q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: id, Kind: form.Kind, Title: strings.TrimSpace(form.Title), OwnerCharacterID: nullID(owner)})
		if err != nil {
			return err
		}
		_, err = q.CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: vis})
		return err
	})
	if err != nil {
		return err
	}
	a.Hub.Publish(id, hub.Topic{Kind: hub.TopicNotes, ID: id})
	c.Redirect(http.StatusSeeOther, route.Path(route.Note, note.ID))
	return nil
}

// noteAccess loads a note in a campaign the requester belongs to. Whether
// they may see it at all is decided by noteView.
func (a *App) noteAccess(c *gin.Context) (who, sqlc.Note, sqlc.Campaign, bool) {
	id, ok := a.idParam(c, "note")
	if !ok {
		return who{}, sqlc.Note{}, sqlc.Campaign{}, false
	}
	note, err := a.Store.GetNote(c.Request.Context(), id)
	if !a.found(c, err) {
		return who{}, sqlc.Note{}, sqlc.Campaign{}, false
	}
	w, camp, ok := a.campaignAccess(c, note.CampaignID)
	return w, note, camp, ok
}

// manageNote is noteAccess for changes: only the note's manager passes.
func (a *App) manageNote(c *gin.Context) (sqlc.Note, bool) {
	w, note, _, ok := a.noteAccess(c)
	if !ok {
		return note, false
	}
	if !perm.CanManageNote(w.viewer, note.OwnerCharacterID.Int64) {
		a.fail(c, http.StatusForbidden, "Only the note's author can change it.")
		return note, false
	}
	locked, err := a.noteLocked(c.Request.Context(), w.viewer, note.ID)
	if err != nil {
		a.serverError(c, err)
		return note, false
	}
	if locked {
		a.fail(c, http.StatusForbidden, "Shared notes are read-only at home.")
		return note, false
	}
	return note, true
}

// noteLocked is perm.LocksShared for a stored note; it asks the database
// only at home.
func (a *App) noteLocked(ctx context.Context, v perm.Viewer, noteID int64) (bool, error) {
	if !v.AtHome {
		return false, nil
	}
	has, err := a.Store.NoteHasSharedBlocks(ctx, noteID)
	return perm.LocksShared(v, has), err
}

func (a *App) notePage(c *gin.Context) error {
	w, note, camp, ok := a.noteAccess(c)
	if !ok {
		return nil
	}
	l := a.layout(w, note.Title, views.NavNotes, &camp, true)
	l.Modules = []string{"/assets/js/notes.js"}
	ctx := c.Request.Context()
	nv, visible, err := a.noteView(ctx, w.viewer, note)
	if err != nil {
		return err
	}
	if !visible {
		a.fail(c, http.StatusNotFound, "Not found.")
		return nil
	}
	nv.Backlinks, err = a.backlinks(ctx, w, note.CampaignID, enum.LinkNote, note.ID)
	if err != nil {
		return err
	}
	if nv.EntityLinks, err = a.entityMentions(ctx, w, note.CampaignID, enum.LinkNote, note.ID); err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.NotePage(l, nv))
	return nil
}

// entityMentions lists the NPCs whose fields link a target, as the viewer
// may see them (revealed links of revealed NPCs for players).
func (a *App) entityMentions(ctx context.Context, w who, campaignID int64, targetType enum.LinkTarget, targetID int64) ([]sqlc.ListEntityFieldBacklinksRow, error) {
	note, entity, image := linkTarget(targetType, targetID)
	return a.Store.ListEntityFieldBacklinks(ctx, sqlc.ListEntityFieldBacklinksParams{
		CampaignID: campaignID, TargetNoteID: note, TargetEntityID: entity, TargetImageID: image, IsGm: w.isGM(),
	})
}

func (a *App) backlinks(ctx context.Context, w who, campaignID int64, targetType enum.LinkTarget, targetID int64) ([]sqlc.ListBacklinksRow, error) {
	note, entity, image := linkTarget(targetType, targetID)
	rows, err := a.Store.ListBacklinks(ctx, sqlc.ListBacklinksParams{
		TargetNoteID: note, TargetEntityID: entity, TargetImageID: image, CampaignID: campaignID, IsGm: w.isGM(), CharacterID: w.viewer.CharacterID,
	})
	if err != nil {
		return nil, err
	}
	// A note does not list itself as a backlink.
	return slices.DeleteFunc(rows, func(r sqlc.ListBacklinksRow) bool { return targetType == enum.LinkNote && r.ID == targetID }), nil
}

func (a *App) renameNote(c *gin.Context) error {
	note, ok := a.manageNote(c)
	if !ok {
		return nil
	}
	var form noteForm
	if !a.bindForm(c, &form) {
		return nil
	}
	if err := a.Store.RenameNote(c.Request.Context(), sqlc.RenameNoteParams{ID: note.ID, Title: strings.TrimSpace(form.Title), Kind: form.Kind}); err != nil {
		return err
	}
	a.changed(c, note.CampaignID, hub.Topic{Kind: hub.TopicNote, ID: note.ID}, hub.Topic{Kind: hub.TopicNotes, ID: note.CampaignID})
	return nil
}

func (a *App) deleteNote(c *gin.Context) error {
	note, ok := a.manageNote(c)
	if !ok {
		return nil
	}
	if err := a.Store.DeleteNote(c.Request.Context(), note.ID); err != nil {
		return err
	}
	a.sweepFiles(c)
	a.Hub.Publish(note.CampaignID, hub.Topic{Kind: hub.TopicNotes, ID: note.CampaignID})
	redirect(c, route.Path(route.CampaignNotes, note.CampaignID))
	return nil
}

// blockAccessForm is the visibility form shared by new and existing blocks.
type blockAccessForm struct {
	Visibility     enum.Visibility `form:"visibility"       validate:"required"`
	PlayersCanEdit bool            `form:"players_can_edit"`
	Audience       []int64         `form:"audience"`
}

// readAccessForm validates the form against what the viewer may set.
func (a *App) readAccessForm(c *gin.Context, w who, note sqlc.Note) (blockAccessForm, bool) {
	var form blockAccessForm
	if !a.bindForm(c, &form) {
		return form, false
	}
	vis := form.Visibility
	if !perm.ValidBlockVisibility(w.viewer, note.OwnerCharacterID.Int64, note.Kind == enum.NoteGameSession, vis) {
		a.fail(c, http.StatusUnprocessableEntity, "You cannot give a block that visibility here.")
		return form, false
	}
	if vis != enum.VisPlayers {
		form.PlayersCanEdit = false
	}
	if vis != enum.VisListed {
		form.Audience = nil
	}
	return form, true
}

func setAudience(ctx context.Context, q *sqlc.Queries, blockID int64, audience []int64) error {
	if err := q.ClearBlockAudience(ctx, blockID); err != nil {
		return err
	}
	for _, id := range audience {
		// AddBlockAudience only inserts characters of the block's campaign.
		if err := q.AddBlockAudience(ctx, sqlc.AddBlockAudienceParams{BlockID: blockID, CharacterID: id}); err != nil {
			return err
		}
	}
	return nil
}

// addBlock is open to anyone readAccessForm lets set a visibility: the
// note's manager, or a player in the GM's game session note.
func (a *App) addBlock(c *gin.Context) error {
	w, note, _, ok := a.noteAccess(c)
	if !ok {
		return nil
	}
	form, ok := a.readAccessForm(c, w, note)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		b, err := q.CreateBlock(ctx, sqlc.CreateBlockParams{
			NoteID: note.ID, Visibility: form.Visibility, PlayersCanEdit: form.PlayersCanEdit, AuthorCharacterID: nullID(w.viewer.CharacterID),
		})
		if err != nil {
			return err
		}
		if err := setAudience(ctx, q, b.ID, form.Audience); err != nil {
			return err
		}
		return q.TouchNote(ctx, note.ID)
	})
	if err != nil {
		return err
	}
	a.changed(c, note.CampaignID, hub.Topic{Kind: hub.TopicNote, ID: note.ID}, hub.Topic{Kind: hub.TopicNotes, ID: note.CampaignID})
	return nil
}

// loadBlock reads a block with what the permission rules need.
func (a *App) loadBlock(ctx context.Context, id int64) (sqlc.GetBlockRow, perm.Block, error) {
	b, err := a.Store.GetBlock(ctx, id)
	if err != nil {
		return b, perm.Block{}, err
	}
	audience, err := a.Store.ListBlockAudience(ctx, id)
	if err != nil {
		return b, perm.Block{}, err
	}
	return b, perm.Block{
		Visibility: b.Visibility, PlayersCanEdit: b.PlayersCanEdit,
		NoteOwnerID: b.OwnerCharacterID.Int64, Audience: audience, AuthorID: b.AuthorCharacterID.Int64,
	}, nil
}

// blockAccessLevel is what a block request needs.
type blockAccessLevel int

const (
	needRead blockAccessLevel = iota
	needEdit
	needManage  // delete or re-scope this block
	needArrange // reorder the note's blocks
)

// blockFor loads the :block parameter and checks the requester's rights.
// Every collaboration call goes through here.
func (a *App) blockFor(c *gin.Context, need blockAccessLevel) (who, sqlc.GetBlockRow, perm.Block, bool) {
	id, ok := a.idParam(c, "block")
	if !ok {
		return who{}, sqlc.GetBlockRow{}, perm.Block{}, false
	}
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return w, sqlc.GetBlockRow{}, perm.Block{}, false
	}
	b, pb, err := a.loadBlock(c.Request.Context(), id)
	if !a.found(c, err) {
		return w, b, pb, false
	}
	member := w.isGM() || w.campaignID() == b.CampaignID
	allowed := member && perm.CanReadBlock(w.viewer, pb)
	switch need {
	case needEdit:
		allowed = allowed && perm.CanEditBlock(w.viewer, pb)
	case needManage:
		allowed = member && perm.CanManageBlock(w.viewer, pb)
	case needArrange:
		allowed = member && perm.CanManageNote(w.viewer, pb.NoteOwnerID)
		if allowed {
			locked, err := a.noteLocked(c.Request.Context(), w.viewer, b.NoteID)
			if err != nil {
				a.serverError(c, err)
				return w, b, pb, false
			}
			allowed = !locked
		}
	}
	if !allowed {
		// The same answer as a missing block, so secrets stay invisible.
		a.fail(c, http.StatusNotFound, "Not found.")
		return w, b, pb, false
	}
	return w, b, pb, true
}

func (a *App) blockAccess(c *gin.Context) error {
	w, b, _, ok := a.blockFor(c, needManage)
	if !ok {
		return nil
	}
	note, err := a.Store.GetNote(c.Request.Context(), b.NoteID)
	if err != nil {
		return err
	}
	form, ok := a.readAccessForm(c, w, note)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		if err := q.UpdateBlockAccess(ctx, sqlc.UpdateBlockAccessParams{ID: b.ID, Visibility: form.Visibility, PlayersCanEdit: form.PlayersCanEdit}); err != nil {
			return err
		}
		return setAudience(ctx, q, b.ID, form.Audience)
	})
	if err != nil {
		return err
	}
	a.changed(c, b.CampaignID, hub.Topic{Kind: hub.TopicNote, ID: b.NoteID}, hub.Topic{Kind: hub.TopicNotes, ID: b.CampaignID})
	return nil
}

func (a *App) moveBlock(c *gin.Context) error {
	w, b, _, ok := a.blockFor(c, needArrange)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		blocks, err := q.ListReadableBlocks(ctx, sqlc.ListReadableBlocksParams{NoteID: b.NoteID, IsGm: w.isGM(), CharacterID: w.viewer.CharacterID})
		if err != nil {
			return err
		}
		i := slices.IndexFunc(blocks, func(x sqlc.NoteBlock) bool { return x.ID == b.ID })
		j := i + 1
		if c.PostForm("dir") == "up" {
			j = i - 1
		}
		if i < 0 || j < 0 || j >= len(blocks) {
			return nil
		}
		// Renumber so equal positions from older rows cannot tie.
		blocks[i], blocks[j] = blocks[j], blocks[i]
		for pos, blk := range blocks {
			if err := q.SetBlockPosition(ctx, sqlc.SetBlockPositionParams{ID: blk.ID, Position: int64(pos + 1)}); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	a.changed(c, b.CampaignID, hub.Topic{Kind: hub.TopicNote, ID: b.NoteID})
	return nil
}

func (a *App) deleteBlock(c *gin.Context) error {
	_, b, _, ok := a.blockFor(c, needManage)
	if !ok {
		return nil
	}
	if err := a.Store.DeleteBlock(c.Request.Context(), b.ID); err != nil {
		return err
	}
	a.sweepFiles(c)
	a.changed(c, b.CampaignID, hub.Topic{Kind: hub.TopicNote, ID: b.NoteID}, hub.Topic{Kind: hub.TopicNotes, ID: b.CampaignID})
	return nil
}

// sweepFiles drops the files a delete just orphaned; a failure is only logged,
// the delete itself went through.
func (a *App) sweepFiles(c *gin.Context) {
	if err := a.CleanAttachments(c.Request.Context()); err != nil {
		slog.Error("clean attachments", "err", err)
	}
}

// stepsPayload is a run of ProseMirror steps starting after Version.
// msgSteps is the websocket message type carrying note steps (MSG.steps in the JS).
const msgSteps = "steps"

type stepsPayload struct {
	Type      string            `json:"type,omitempty"`
	Block     int64             `json:"block,omitempty"`
	Version   int64             `json:"version"`
	Steps     []json.RawMessage `json:"steps"`
	ClientIDs []string          `json:"clientIDs"`
}

// keepSteps is how many steps a snapshot leaves behind it, so a device that
// was offline for a while can still rebase its unsent edits.
const keepSteps = 1000

// stepsSince returns the steps after version, or false if they were
// trimmed (the client must reload the document).
func (a *App) stepsSince(ctx context.Context, q *sqlc.Queries, b sqlc.GetBlockRow, version int64) (stepsPayload, bool, error) {
	if version > b.Version {
		return stepsPayload{}, false, nil
	}
	rows, err := q.ListBlockStepsSince(ctx, sqlc.ListBlockStepsSinceParams{BlockID: b.ID, Version: version})
	if err != nil {
		return stepsPayload{}, false, err
	}
	if (len(rows) == 0 && version != b.Version) || (len(rows) > 0 && rows[0].Version != version+1) {
		return stepsPayload{}, false, nil
	}
	p := stepsPayload{Version: version, Steps: []json.RawMessage{}, ClientIDs: []string{}}
	for _, r := range rows {
		p.Steps = append(p.Steps, json.RawMessage(r.StepJson))
		p.ClientIDs = append(p.ClientIDs, r.ClientID)
	}
	return p, true, nil
}

// blockDoc returns the latest snapshot plus the steps after it, so the
// editor can rebuild the current version.
func (a *App) blockDoc(c *gin.Context) error {
	w, b, pb, ok := a.blockFor(c, needRead)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	var steps stepsPayload
	var complete bool
	// One transaction, so a snapshot landing in between cannot pair the
	// document with steps numbered after another version.
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if b, err = q.GetBlock(ctx, b.ID); err != nil {
			return err
		}
		steps, complete, err = a.stepsSince(ctx, q, b, b.DocVersion)
		return err
	})
	if err != nil {
		return err
	}
	if !complete { // only a damaged database: snapshots never pass the version, trims keep the steps after them
		c.Status(http.StatusGone)
		return nil
	}
	c.JSON(http.StatusOK, gin.H{
		keyDoc: json.RawMessage(b.DocJson), keyDocVersion: b.DocVersion, keySteps: steps.Steps,
		keyClientIDs: steps.ClientIDs, keyEditable: perm.CanEditBlock(w.viewer, pb),
	})
	return nil
}

// blockSteps lets a client that missed pushes catch up.
func (a *App) blockSteps(c *gin.Context) error {
	_, b, _, ok := a.blockFor(c, needRead)
	if !ok {
		return nil
	}
	since, err := strconv.ParseInt(c.Query("since"), 10, 64)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}
	steps, ok, err := a.stepsSince(c.Request.Context(), a.Store.Queries, b, since)
	switch {
	case err != nil:
		return err
	case !ok:
		c.Status(http.StatusGone)
	default:
		c.JSON(http.StatusOK, steps)
	}
	return nil
}

// Limits of submitBody.Steps; tags cannot name constants (TestSharedLimitsMatchTags), shared.js sends them to the browser.
const (
	maxSteps     = 500
	maxStepBytes = 262144
)

type submitBody struct {
	Version  int64             `json:"version"  validate:"min=0"`
	ClientID string            `json:"clientID" validate:"required,max=64"`
	Steps    []json.RawMessage `json:"steps"    validate:"required,min=1,max=500,dive,max=262144"`
}

// submitSteps is the collab authority: steps based on the current version
// are appended and broadcast; anything else gets 409 with the steps the
// client is missing, so it can rebase and retry.
func (a *App) submitSteps(c *gin.Context) error {
	_, b, _, ok := a.blockFor(c, needEdit)
	if !ok {
		return nil
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxCollabBody)
	var body submitBody
	if c.ShouldBindJSON(&body) != nil {
		// notes.js shows it; the limits are all a working editor can break.
		c.JSON(http.StatusBadRequest, gin.H{keyError: errTooLarge})
		return nil
	}
	// A raw step skips TipTap's link validation, so a script link is refused here.
	for _, s := range body.Steps {
		if msg := badStep(s); msg != "" {
			c.JSON(http.StatusUnprocessableEntity, gin.H{keyError: msg})
			return nil
		}
	}
	ctx := c.Request.Context()
	var conflict *stepsPayload
	var gone bool
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		cur, err := q.GetBlock(ctx, b.ID)
		if err != nil {
			return err
		}
		if body.Version != cur.Version {
			p, ok, err := a.stepsSince(ctx, q, cur, body.Version)
			if err != nil {
				return err
			}
			conflict, gone = &p, !ok
			return nil
		}
		for i, step := range body.Steps {
			if err := q.InsertBlockStep(ctx, sqlc.InsertBlockStepParams{
				BlockID: b.ID, Version: body.Version + int64(i) + 1, ClientID: body.ClientID, StepJson: string(step),
			}); err != nil {
				return err
			}
		}
		return q.SetBlockVersion(ctx, sqlc.SetBlockVersionParams{ID: b.ID, Version: body.Version + int64(len(body.Steps))})
	})
	switch {
	case err != nil:
		return err
	case gone:
		c.Status(http.StatusGone)
		return nil
	case conflict != nil:
		c.JSON(http.StatusConflict, conflict)
		return nil
	}

	ids := make([]string, len(body.Steps))
	for i := range ids {
		ids[i] = body.ClientID
	}
	a.broadcastSteps(ctx, b, stepsPayload{Type: msgSteps, Block: b.ID, Version: body.Version, Steps: body.Steps, ClientIDs: ids})
	c.JSON(http.StatusOK, gin.H{keyVersion: body.Version + int64(len(body.Steps))})
	return nil
}

// broadcastSteps sends steps to subscribed clients that may still read the
// block. The rule is checked against the block as it is now, so a client
// that lost access stops receiving immediately.
func (a *App) broadcastSteps(ctx context.Context, b sqlc.GetBlockRow, p stepsPayload) {
	_, pb, err := a.loadBlock(ctx, b.ID)
	if err != nil {
		slog.Error("broadcast steps", "block", b.ID, "err", err)
		return
	}
	raw, err := json.Marshal(p)
	if err != nil {
		slog.Error("broadcast steps", "block", b.ID, "err", err)
		return
	}
	a.Hub.SendRaw(b.CampaignID, raw, func(cl *hub.Client) bool {
		return cl.Subscribed(b.ID) && perm.CanReadBlock(cl.Viewer, pb)
	})
}

const (
	errScriptLink = "links must not run scripts"
	errNoStepType = "a step needs a stepType"
	errTooLarge   = "This edit is too large to save."
)

// safeSchemes run no script when followed: TipTap Link's defaults, which it
// autolinks, so refusing one would stall the block's saving.
var safeSchemes = []string{"http", "https", "ftp", "ftps", "mailto", "tel", "callto", "sms", "cid", "xmpp"}

// unsafeHref reports whether the decoded JSON holds an "href" at any depth
// whose scheme is not in safeSchemes. Browsers drop whitespace and control
// characters inside a URL, so they are removed before the scheme is read.
func unsafeHref(v any) bool {
	switch v := v.(type) {
	case map[string]any:
		for k, x := range v {
			if s, ok := x.(string); ok && k == "href" {
				clean := strings.ToLower(strings.Map(func(r rune) rune {
					if r <= ' ' || r == 0x7f {
						return -1
					}
					return r
				}, s))
				// A scheme is what precedes the first ':' before any '/', '?' or '#'.
				if i := strings.IndexAny(clean, ":/?#"); i >= 0 && clean[i] == ':' && !slices.Contains(safeSchemes, clean[:i]) {
					return true
				}
			} else if unsafeHref(x) {
				return true
			}
		}
	case []any:
		return slices.ContainsFunc(v, unsafeHref)
	}
	return false
}

// hasUnsafeHref is unsafeHref for a raw JSON value; invalid JSON counts as unsafe.
func hasUnsafeHref(raw json.RawMessage) bool {
	var v any
	return json.Unmarshal(raw, &v) != nil || unsafeHref(v)
}

// badStep returns why a raw collab step is refused, or "": it must be a JSON
// object with a string stepType and no script link.
func badStep(raw json.RawMessage) string {
	var v any
	if json.Unmarshal(raw, &v) != nil {
		return errScriptLink
	}
	m, _ := v.(map[string]any)
	if _, ok := m["stepType"].(string); !ok {
		return errNoStepType
	}
	if unsafeHref(v) {
		return errScriptLink
	}
	return ""
}

type snapshotBody struct {
	Version int64           `json:"version" validate:"min=1"`
	Doc     json.RawMessage `json:"doc"     validate:"required"`
	Text    string          `json:"text"    validate:"max=1000000"`
	Links   []struct {
		Type enum.LinkTarget `json:"type" validate:"oneof=note entity image"`
		ID   int64           `json:"id"   validate:"min=1"`
	} `json:"links" validate:"max=1000,dive"`
}

// blockSnapshot stores a client's full document if it is newer than the
// stored one, trims the steps behind it and records its links. Trusting
// the client here is accepted for a LAN of friends.
func (a *App) blockSnapshot(c *gin.Context) error {
	w, b, _, ok := a.blockFor(c, needEdit)
	if !ok {
		return nil
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxCollabBody)
	var body snapshotBody
	if err := c.ShouldBindJSON(&body); err != nil || !json.Valid(body.Doc) {
		c.Status(http.StatusBadRequest)
		return nil
	}
	if hasUnsafeHref(body.Doc) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{keyError: errScriptLink})
		return nil
	}
	ctx := c.Request.Context()
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		n, err := q.SaveBlockSnapshot(ctx, sqlc.SaveBlockSnapshotParams{ID: b.ID, DocJson: string(body.Doc), DocVersion: body.Version, Text: body.Text})
		if err != nil || n == 0 {
			return err
		}
		if err := q.TrimBlockSteps(ctx, sqlc.TrimBlockStepsParams{BlockID: b.ID, Version: body.Version - keepSteps}); err != nil {
			return err
		}
		// An image link lets the block's readers see the image, so only the
		// GM adds them; a player's snapshot keeps the ones already there.
		images, err := q.ListBlockImageLinks(ctx, b.ID)
		if err != nil {
			return err
		}
		if err := q.ClearBlockLinks(ctx, b.ID); err != nil {
			return err
		}
		for _, l := range body.Links {
			if l.Type == enum.LinkImage && !w.isGM() && !slices.Contains(images, l.ID) {
				continue
			}
			note, entity, image := linkTarget(l.Type, l.ID)
			if err := q.AddBlockLink(ctx, sqlc.AddBlockLinkParams{SourceNoteID: b.NoteID, SourceBlockID: b.ID, TargetNoteID: note, TargetEntityID: entity, TargetImageID: image}); err != nil {
				return err
			}
		}
		return q.TouchNote(ctx, b.NoteID)
	})
	if err != nil {
		return err
	}
	c.Status(http.StatusNoContent)
	return nil
}

// inlineTypes are served inline; everything else downloads.
var inlineTypes = []string{"image/png", "image/jpeg", "image/gif", "image/webp"}

func (a *App) uploadAttachment(c *gin.Context) {
	_, b, _, ok := a.blockFor(c, needEdit)
	if !ok {
		return
	}
	att, ok := a.receiveUpload(c, b.CampaignID, nullID(b.ID))
	if !ok {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		keyID: att.ID, keyURL: route.Path(route.Attachment, att.ID), keyName: att.Name, keyMime: att.Mime,
		keyImage: slices.Contains(inlineTypes, att.Mime),
	})
}

// receiveUpload stores the request's "file" as an attachment, answering
// with a JSON error itself when it is missing or too large.
func (a *App) receiveUpload(c *gin.Context, campaignID int64, blockID sql.NullInt64) (sqlc.Attachment, bool) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxAttachmentBytes+1<<20)
	fh, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{keyError: "file missing or larger than 20 MB"})
		return sqlc.Attachment{}, false
	}
	if fh.Size > maxAttachmentBytes {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{keyError: "files are limited to 20 MB"})
		return sqlc.Attachment{}, false
	}
	att, err := a.storeUpload(c.Request.Context(), campaignID, blockID, fh)
	if err != nil {
		a.serverError(c, err)
		return att, false
	}
	return att, true
}

// storeUpload saves an uploaded file under a random name and records it.
func (a *App) storeUpload(ctx context.Context, campaignID int64, blockID sql.NullInt64, fh *multipart.FileHeader) (sqlc.Attachment, error) {
	src, err := fh.Open()
	if err != nil {
		return sqlc.Attachment{}, err
	}
	defer func() { _ = src.Close() }()
	return a.storeFile(ctx, campaignID, blockID, filepath.Base(fh.Filename), src)
}

// storeFile saves src under a random name and records it as fileName.
func (a *App) storeFile(ctx context.Context, campaignID int64, blockID sql.NullInt64, fileName string, src io.Reader) (sqlc.Attachment, error) {
	name, mime, size, err := a.saveFile(src)
	if err != nil {
		return sqlc.Attachment{}, err
	}
	att, err := a.Store.CreateAttachment(ctx, sqlc.CreateAttachmentParams{
		CampaignID: campaignID, BlockID: blockID, Path: name, Name: fileName, Mime: mime, Size: size,
	})
	if err != nil {
		_ = os.Remove(filepath.Join(a.attachmentsDir(), name))
	}
	return att, err
}

// saveFile writes src under a random name in the attachments folder, for a
// row the caller adds; a failed write leaves nothing behind.
func (a *App) saveFile(src io.Reader) (name, mime string, size int64, err error) {
	dir := a.attachmentsDir()
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return "", "", 0, err
	}
	name = rand.Text()
	if mime, size, err = saveSniffed(filepath.Join(dir, name), src); err != nil {
		_ = os.Remove(filepath.Join(dir, name))
	}
	return name, mime, size, err
}

// saveSniffed writes src to path and returns the content type sniffed from
// its first bytes (never the client's claim) and its size.
func saveSniffed(path string, src io.Reader) (string, int64, error) {
	head := make([]byte, sniffLen)
	n, err := io.ReadFull(src, head)
	if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) && !errors.Is(err, io.EOF) {
		return "", 0, err
	}
	head = head[:n]
	mime, _, _ := strings.Cut(http.DetectContentType(head), ";")
	dst, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600) //nolint:gosec // path is our random name in our directory
	if err != nil {
		return "", 0, err
	}
	written, err := io.Copy(dst, io.MultiReader(strings.NewReader(string(head)), src))
	if cerr := dst.Close(); err == nil {
		err = cerr
	}
	return mime, written, err
}

// attachment serves a file to viewers who may read its block.
func (a *App) attachment(c *gin.Context) {
	if att, ok := a.readableAttachment(c); ok {
		a.serveAttachment(c, att, cacheAttachment)
	}
}

// readableAttachment loads :attachment for viewers who may read its block,
// answering 404 itself otherwise.
func (a *App) readableAttachment(c *gin.Context) (sqlc.Attachment, bool) {
	id, ok := a.idParam(c, "attachment")
	if !ok {
		return sqlc.Attachment{}, false
	}
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return sqlc.Attachment{}, false
	}
	ctx := c.Request.Context()
	att, err := a.Store.GetAttachment(ctx, id)
	if !a.found(c, err) {
		return att, false
	}
	var block *perm.Block
	if att.BlockID.Valid {
		_, pb, err := a.loadBlock(ctx, att.BlockID.Int64)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			a.serverError(c, err)
			return att, false
		}
		block = &pb
	}
	if (!w.isGM() && w.campaignID() != att.CampaignID) || !perm.CanReadAttachment(w.viewer, block) {
		a.fail(c, http.StatusNotFound, "Not found.")
		return att, false
	}
	return att, true
}

// serveAttachment sends a file the caller has already authorized, sandboxed
// so an uploaded page or script can never run as this site.
// cacheAttachment is the Cache-Control of an attachment URL that does not
// change with its file.
const cacheAttachment = "private, max-age=3600"

func (a *App) serveAttachment(c *gin.Context, att sqlc.Attachment, cacheControl string) {
	f, err := os.Open(filepath.Join(a.attachmentsDir(), filepath.Base(att.Path)))
	if errors.Is(err, os.ErrNotExist) { // the data folder lost it: nothing a report can fix
		a.fail(c, http.StatusNotFound, "Not found.")
		return
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	defer func() { _ = f.Close() }()
	disposition := "attachment"
	if slices.Contains(inlineTypes, att.Mime) {
		disposition = "inline"
	}
	h := c.Writer.Header()
	h.Set("Content-Type", att.Mime)
	h.Set("Content-Disposition", fmt.Sprintf("%s; filename=%q", disposition, att.Name))
	h.Set("Content-Security-Policy", "default-src 'none'; sandbox")
	h.Set("Cache-Control", cacheControl)
	http.ServeContent(c.Writer, c.Request, "", att.CreatedAt, f)
}

// freshFile is how long a new file is spared: storeFile writes it before its
// row, and the note edit that uses it arrives after the row.
const freshFile = time.Minute

// CleanAttachments removes files no attachments row refers to any more
// (their campaign or block was deleted). Run at startup and after deletes.
func (a *App) CleanAttachments(ctx context.Context) error {
	entries, err := os.ReadDir(a.attachmentsDir())
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	paths, err := a.Store.ListAttachmentPaths(ctx)
	if err != nil {
		return err
	}
	known := make(map[string]bool, len(paths))
	for _, p := range paths {
		known[p] = true
	}
	// One file that cannot go (open on Windows, or taken by a concurrent
	// sweep) must not keep the rest.
	var errs error
	for _, e := range entries {
		if e.IsDir() || known[e.Name()] {
			continue
		}
		if info, err := e.Info(); err == nil && time.Since(info.ModTime()) > freshFile {
			if err := os.Remove(filepath.Join(a.attachmentsDir(), e.Name())); err != nil && !errors.Is(err, os.ErrNotExist) {
				errs = errors.Join(errs, err)
			}
		}
	}
	return errs
}

// linkKinds maps the prefixes that narrow [[ links to the kind they pick.
var linkKinds = map[string]string{"note": string(enum.LinkNote), "npc": string(enum.EntityNPC), "img": string(enum.LinkImage), "image": string(enum.LinkImage)}

// linkItems lists notes, entities and (for the GM) images the viewer may
// see whose title contains q, for [[ links and the quick switcher. A
// "kind:" prefix such as "img:" narrows the list to that kind.
func (a *App) linkItems(ctx context.Context, w who, campaignID int64, q string) ([]views.LinkItem, error) {
	q = strings.ToLower(strings.TrimSpace(q))
	kind := ""
	if prefix, rest, ok := strings.Cut(q, ":"); ok && linkKinds[prefix] != "" {
		kind, q = linkKinds[prefix], strings.TrimSpace(rest)
	}
	want := func(k, title string) bool {
		return (kind == "" || kind == k) && strings.Contains(strings.ToLower(title), q)
	}
	items := []views.LinkItem{}
	notes, err := a.readableNotes(ctx, w.viewer, campaignID)
	if err != nil {
		return nil, err
	}
	for _, n := range notes {
		if note := string(enum.LinkNote); want(note, n.Title) {
			items = append(items, views.LinkItem{ID: fmt.Sprintf("%s:%d", note, n.ID), Label: n.Title, Kind: note, URL: route.Path(route.Note, n.ID)})
		}
	}
	ents, err := a.entityList(ctx, w.viewer, campaignID)
	if err != nil {
		return nil, err
	}
	for _, e := range ents.Entities {
		if npc := string(enum.EntityNPC); want(npc, e.Name) {
			items = append(items, views.LinkItem{ID: fmt.Sprintf("entity:%d", e.ID), Label: e.Name, Kind: npc, URL: route.Path(route.Entity, e.ID)})
		}
	}
	if image := string(enum.LinkImage); w.isGM() && (kind == "" || kind == image) {
		imgs, err := a.Store.ListImages(ctx, sqlc.ListImagesParams{CampaignID: campaignID, Q: q})
		if err != nil {
			return nil, err
		}
		for _, img := range imgs {
			items = append(items, views.LinkItem{ID: fmt.Sprintf("%s:%d", image, img.ID), Label: img.Name, Kind: image, URL: route.Path(route.Image, img.ID)})
		}
	}
	return items[:min(len(items), maxLinkSuggestions)], nil
}

func (a *App) linkSuggestions(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, _, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	items, err := a.linkItems(c.Request.Context(), w, id, c.Query("q"))
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, items)
	return nil
}

func (a *App) switcher(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, _, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	items, err := a.linkItems(c.Request.Context(), w, id, c.Query("q"))
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.SwitcherResults(items))
	return nil
}
