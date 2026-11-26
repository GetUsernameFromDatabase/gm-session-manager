package web

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// takeHomeAccess loads :character for whoever may take it home. The player
// app has no table to bring it back to, so it answers 404 there.
func (a *App) takeHomeAccess(c *gin.Context) (who, sqlc.Character, bool) {
	if a.Player {
		a.fail(c, http.StatusNotFound, "Not found.")
		return who{}, sqlc.Character{}, false
	}
	w, ch, ok := a.characterAccess(c)
	if ok && !perm.CanTakeHome(w.viewer, ch.ID) {
		a.fail(c, http.StatusForbidden, "That is not your character.")
		return w, ch, false
	}
	return w, ch, ok
}

// takeHomeDialog is the dialog body: the app files on offer and the download.
func (a *App) takeHomeDialog(c *gin.Context) error {
	w, ch, ok := a.takeHomeAccess(c)
	if !ok {
		return nil
	}
	names, err := downloadNames(filepath.Join(a.DataDir, DownloadsDir))
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.TakeHomeBody(views.TakeHomeView{CharacterID: ch.ID, Name: ch.Name, Downloads: names, ByGM: w.isGM()}))
	return nil
}

// takeHome writes the character's take-home zip and records the take-out
// (token and what was sent) as the merge base for the bring-back. The player
// sends everything of theirs; the GM leaves out what only the player may read.
func (a *App) takeHome(c *gin.Context) error {
	w, ch, ok := a.takeHomeAccess(c)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	serverID, err := a.Store.GetServerID(ctx)
	if err != nil {
		return err
	}
	camp, err := a.Store.GetCampaign(ctx, ch.CampaignID)
	if err != nil {
		return err
	}

	var th export.TakeHome
	var rec sqlc.TakeHome
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var base export.TakeHomeBase
		var err error
		if th, base, err = buildTakeHome(ctx, q, w.viewer, ch); err != nil {
			return err
		}
		body, err := json.Marshal(base)
		if err != nil {
			return err
		}
		rec, err = q.CreateTakeHome(ctx, sqlc.CreateTakeHomeParams{CharacterID: ch.ID, Token: rand.Text(), BaseJson: string(body)})
		return err
	})
	if err != nil {
		return err
	}

	sources, err := a.Store.ListTakeHomeSources(ctx, ch.CampaignID)
	if err != nil {
		return err
	}
	for _, s := range sources {
		th.Sources = append(th.Sources, export.TakeHomeSource(s))
	}
	entries, err := a.Store.ListTakeHomeEntries(ctx, sqlc.ListTakeHomeEntriesParams{CharacterID: nullID(ch.ID), CampaignID: ch.CampaignID})
	if err != nil {
		return err
	}
	for _, e := range entries {
		th.Entries = append(th.Entries, export.TakeHomeEntry{
			Kind: e.Kind, Slug: e.Slug, Name: e.Name, SourceKey: e.SourceKey, Data: json.RawMessage(e.DataJson),
			BodyText: e.BodyText, ContentHash: e.ContentHash, Scoped: e.Scoped,
		})
	}
	th.Manifest = export.TakeHomeManifest{
		Format: export.TakeHomeFormat, AppVersion: a.Version, Token: rec.Token, ServerID: serverID,
		Campaign:    export.TakeHomeCampaign{Name: camp.Name, System: camp.System, Edition: camp.Edition},
		CharacterID: ch.ID, CreatedAt: rec.CreatedAt.UTC(),
	}

	a.publish(c, ch.CampaignID, hub.Topic{Kind: hub.TopicCharacter, ID: ch.ID}, hub.Topic{Kind: hub.TopicCampaign, ID: ch.CampaignID})
	c.Header("Content-Type", "application/zip")
	c.Header("Content-Disposition", `attachment; filename="`+export.TakeHomeFileName(ch.Name, rec.CreatedAt)+`"`)
	if err := export.WriteTakeHome(c.Writer, th, a.attachmentsDir()); err != nil {
		// Headers are gone; the truncated zip tells the browser it failed.
		a.serverError(c, err)
	}
	return nil
}

// buildTakeHome reads what goes into the file about one character, and the
// base the bring-back is merged against. Only note blocks both the character
// and the requester may read are sent: a GM's file never holds the player's
// private blocks, or their files.
func buildTakeHome(ctx context.Context, q *sqlc.Queries, viewer perm.Viewer, ch sqlc.Character) (export.TakeHome, export.TakeHomeBase, error) {
	th := export.TakeHome{Character: export.NewCharacterFile(ch)}
	base := export.TakeHomeBase{
		Character: export.TakeHomeBaseCharacter{Sheet: json.RawMessage(ch.SheetJson), HP: ch.Hp, MaxHP: ch.MaxHp, TempHP: ch.TempHp},
		Blocks:    map[int64]int64{}, Homebrew: map[int64]string{},
	}

	levelUps, err := q.ListLevelUps(ctx, ch.ID)
	if err != nil {
		return th, base, err
	}
	for _, l := range levelUps {
		th.LevelUps = append(th.LevelUps, export.TakeHomeLevelUp{ID: l.ID, MaxHP: l.MaxHp, Sheet: json.RawMessage(l.SheetJson), CreatedAt: l.CreatedAt.UTC()})
		base.LevelUps = append(base.LevelUps, l.ID)
	}

	looks, err := q.ListCharacterImages(ctx, ch.ID)
	if err != nil {
		return th, base, err
	}
	for _, l := range looks {
		img, err := q.GetImage(ctx, l.ID)
		if err != nil {
			return th, base, err
		}
		look := export.TakeHomeLook{ImageID: l.ID, Name: l.Name, Description: img.Description, URL: l.Url.String, IsRevealed: l.IsRevealed, IsCurrent: l.IsCurrent}
		if l.AttachmentID.Valid {
			att, err := q.GetAttachment(ctx, l.AttachmentID.Int64)
			if err != nil {
				return th, base, err
			}
			look.File, look.FileName, look.AttachmentID = export.ZipFile(att.Path), att.Name, att.ID
		}
		th.Looks = append(th.Looks, look)
		base.Looks = append(base.Looks, export.TakeHomeBaseLook{ImageID: l.ID, IsRevealed: l.IsRevealed, IsCurrent: l.IsCurrent})
	}

	if err := takeHomeNotes(ctx, q, viewer, ch, &th, &base); err != nil {
		return th, base, err
	}

	own, err := q.ListOwnHomebrew(ctx, nullID(ch.ID))
	if err != nil {
		return th, base, err
	}
	for _, e := range own {
		th.Homebrew = append(th.Homebrew, export.TakeHomeEntry{
			ID: e.ID, Kind: e.Kind, Slug: e.Slug, Name: e.Name, SourceKey: e.SourceKey, Data: json.RawMessage(e.DataJson),
			BodyText: e.BodyText, ContentHash: e.ContentHash, Scoped: e.Scoped,
		})
		base.Homebrew[e.ID] = e.ContentHash
	}
	return th, base, nil
}

// takeHomeNotes adds the character's own notes, with the blocks both the
// character and the requester may read: a GM-initiated file leaves the
// player's 'owner' blocks and their files at the table.
func takeHomeNotes(ctx context.Context, q *sqlc.Queries, viewer perm.Viewer, ch sqlc.Character, th *export.TakeHome, base *export.TakeHomeBase) error {
	notes, err := q.ListOwnNotes(ctx, nullID(ch.ID))
	if err != nil {
		return err
	}
	blocks, err := q.ListOwnNoteBlocks(ctx, nullID(ch.ID))
	if err != nil {
		return err
	}
	aud, err := q.ListOwnNoteAudience(ctx, nullID(ch.ID))
	if err != nil {
		return err
	}
	files, err := q.ListOwnNoteAttachments(ctx, nullID(ch.ID))
	if err != nil {
		return err
	}
	audience := map[int64][]int64{}
	for _, r := range aud {
		audience[r.BlockID] = append(audience[r.BlockID], r.CharacterID)
	}
	filesOf := map[int64][]export.TakeHomeBlockFile{}
	for _, f := range files {
		filesOf[f.BlockID.Int64] = append(filesOf[f.BlockID.Int64], export.TakeHomeBlockFile{AttachmentID: f.ID, Name: f.Name, File: export.ZipFile(f.Path)})
	}
	player := perm.Player(ch.ID)
	byNote := map[int64][]sqlc.NoteBlock{}
	for _, b := range blocks {
		byNote[b.NoteID] = append(byNote[b.NoteID], b)
	}
	for _, n := range notes {
		out := export.TakeHomeNote{ID: n.ID, Kind: n.Kind, Title: n.Title, CreatedAt: n.CreatedAt.UTC(), UpdatedAt: n.UpdatedAt.UTC()}
		for _, b := range byNote[n.ID] {
			pb := perm.Block{Visibility: b.Visibility, PlayersCanEdit: b.PlayersCanEdit, NoteOwnerID: ch.ID, Audience: audience[b.ID], AuthorID: b.AuthorCharacterID.Int64}
			if !perm.CanReadBlock(player, pb) || !perm.CanReadBlock(viewer, pb) {
				continue
			}
			out.Blocks = append(out.Blocks, export.TakeHomeBlock{
				ID: b.ID, Position: b.Position, Visibility: b.Visibility, PlayersCanEdit: b.PlayersCanEdit,
				Doc: json.RawMessage(b.DocJson), Version: b.DocVersion, Text: b.Text,
				ByCharacter: b.AuthorCharacterID.Int64 == ch.ID, Files: filesOf[b.ID],
			})
			// doc_json is the last snapshot: steps after it are not in the file, so
			// the base is the snapshot's version and the bring-back sees them as table edits.
			base.Blocks[b.ID] = b.DocVersion
		}
		// A note with blocks the requester may not read is left out whole, title
		// included; an empty one goes only to its own player.
		if len(out.Blocks) == 0 && (len(byNote[n.ID]) > 0 || viewer.IsGM) {
			continue
		}
		th.Notes = append(th.Notes, out)
		base.Notes = append(base.Notes, n.ID)
	}
	return nil
}
