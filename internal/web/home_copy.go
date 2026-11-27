package web

// The player app's side of taking a character home: list the characters on
// this computer, open a take-home zip and log the browser in as its character.
//
// shortcut: [[note:…]] and mention tokens in imported notes keep the table's
// ids and do not resolve at home; resolve them through the id maps when the
// bring-back needs them.
// shortcut: 'listed' blocks lose their audience at home (the character becomes
// the only one); keep it when a second character plays at home together.
// shortcut: conditions, effect presets and devices are not taken home.

import (
	"bytes"
	"cmp"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

const (
	// maxTakeHomeBytes caps an uploaded take-home file.
	maxTakeHomeBytes = 512 << 20
	// formReplace is the import form's field a confirmed replace sets.
	formReplace = "replace"
	// formHandoff is the import form's field that takes the file from the phone handoff.
	formHandoff = "handoff"
)

// userf is a userError shown to the player as a sentence.
func userf(format string, args ...any) userError { return userError{fmt.Errorf(format, args...)} }

// notPlayerApp answers 404 where only the GM's app has a page.
func (a *App) notPlayerApp(c *gin.Context) {
	if a.Player {
		a.fail(c, http.StatusNotFound, "Not found.")
	}
}

// homeSources puts the sources added at home on the request's context (the
// player app only), so the views mark them and protect the table's.
func (a *App) homeSources(c *gin.Context) {
	if _, ok := auth.FromContext(c).CharacterID(); !ok {
		return
	}
	keys, err := a.Store.ListHomeSourceKeys(c.Request.Context())
	if err != nil {
		a.serverError(c, err)
		c.Abort()
		return
	}
	c.Request = c.Request.WithContext(views.WithHomeSources(c.Request.Context(), keys))
}

// onlyPlayerApp answers 404 where only the GM's app runs.
func (a *App) onlyPlayerApp(c *gin.Context) {
	if !a.Player {
		a.fail(c, http.StatusNotFound, "Not found.")
		c.Abort()
	}
}

// dropSession ends this device's session, as logout does, and tells the GM's
// roster; no session is no error.
func (a *App) dropSession(c *gin.Context) error {
	sess := auth.FromContext(c).Session
	if sess == nil {
		return nil
	}
	w, err := a.whoami(c)
	if err != nil {
		return err
	}
	if err := a.Auth.Revoke(c.Request.Context(), sess.ID); err != nil {
		return err
	}
	a.Hub.CloseSession(sess.ID)
	a.reqs.forget(sessionKey(*sess))
	if w.character != nil {
		a.Hub.Publish(w.character.CampaignID, hub.Topic{Kind: hub.TopicCampaign, ID: w.character.CampaignID})
	}
	return nil
}

// homePage lists the characters on this computer and the import form.
func (a *App) homePage(c *gin.Context) error {
	if !a.Player {
		a.fail(c, http.StatusNotFound, "Not found.")
		return nil
	}
	ctx := c.Request.Context()
	w, err := a.whoami(c)
	if err != nil {
		return err
	}
	rows, err := a.Store.ListHomeCopies(ctx)
	if err != nil {
		return err
	}
	var copies []views.HomeCopyView
	for _, r := range rows {
		var m export.HomeManifest
		if err := json.Unmarshal([]byte(r.ManifestJson), &m); err != nil {
			return err
		}
		copies = append(copies, views.HomeCopyView{
			CharacterID: r.CharacterID, Name: r.CharacterName, Campaign: r.CampaignName, TakenAt: m.CreatedAt,
			Current: w.character != nil && w.character.ID == r.CharacterID,
		})
	}
	a.render(c, http.StatusOK, views.HomePage(a.layout(w, "Characters", views.NavAtHome, nil, false), copies, a.handoffView()))
	return nil
}

// homePlay switches this browser to another character on this computer, with
// no approval: they are all the player's own.
func (a *App) homePlay(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	if _, err := a.Store.GetHomeCopy(ctx, id); err != nil {
		return err
	}
	ch, err := a.Store.GetCharacter(ctx, id)
	if err != nil {
		return err
	}
	if err := a.dropSession(c); err != nil {
		return err
	}
	if err := a.login(c, ch); err != nil {
		return err
	}
	c.Redirect(http.StatusSeeOther, route.Path(route.Character, ch.ID))
	return nil
}

// savedFile is a take-home file already on disk, waiting for its row.
type savedFile struct {
	path, name, mime string
	size             int64
}

// checkTakeHome refuses a file this app cannot take: a sheet from a newer app,
// a character that does not validate, a link it would not open, an entry whose
// source is missing. It returns the character as the app would store it.
func (a *App) checkTakeHome(file *export.TakeHomeFile) (export.CharacterFile, error) {
	m := file.Manifest
	cf := file.Character
	if m.Token == "" || m.ServerID == "" || m.CharacterID <= 0 || m.Campaign.Name == "" {
		return cf, userError{errors.New("the file has no table or character to come from")}
	}
	if m.Campaign.System != gameSystem || cf.System != gameSystem {
		return cf, userError{errors.New("this app plays " + gameSystem + " only")}
	}
	sheets := []json.RawMessage{cf.Sheet}
	for _, l := range file.LevelUps {
		sheets = append(sheets, l.Sheet)
	}
	for _, s := range sheets {
		switch err := checkSheetFile(s); {
		case errors.Is(err, errNewerSheet):
			return cf, userf("This file comes from a newer app (%s); this player app is %s. Get the matching app from your GM.", m.AppVersion, a.Version)
		case err != nil:
			return cf, userError{err}
		}
	}
	if err := binding.Validator.ValidateStruct(&cf); err != nil {
		return cf, userError{err}
	}
	if err := normalizeCharacterFile(&cf); err != nil {
		return cf, userError{fmt.Errorf("the character is not valid: %w", err)}
	}
	for _, n := range file.Notes {
		for _, b := range n.Blocks {
			if hasUnsafeHref(b.Doc) {
				return cf, userError{errors.New("a note holds a link the app will not open")}
			}
		}
	}
	urls := validator.New()
	for _, l := range file.Looks {
		if l.URL != "" && urls.Var(l.URL, "http_url,max=2000") != nil {
			return cf, userError{fmt.Errorf("the look %q has a link that is not a web address", l.Name)}
		}
	}
	return cf, checkFileSources(file)
}

// checkFileSources checks the compendium part of a take-home or bring-back
// file: sources with a real key, and entries of a listed source (or homebrew,
// which a bring-back never lists among its entries: its homebrew is separate).
// Phase 4 calls it for bring-back files at the table.
func checkFileSources(file *export.TakeHomeFile) error {
	known := map[string]bool{}
	for _, s := range file.Sources {
		if s.SourceKey == "" || s.SourceKey == homebrewSource {
			return userError{errors.New("the file lists a source that is not valid")}
		}
		known[s.SourceKey] = true
	}
	for _, e := range file.Entries {
		if e.SourceKey == homebrewSource && file.Manifest.Kind == export.KindBringBack {
			return userError{errors.New("the file lists homebrew among the compendium entries")}
		}
		if e.SourceKey != homebrewSource && !known[e.SourceKey] {
			return userError{fmt.Errorf("the entry %q comes from a source the file does not list", e.Slug)}
		}
	}
	return nil
}

// homebrewSource is the compendium source key of homebrew entries.
const homebrewSource = "homebrew"

// takeHomeSourcePath is compendium_sources.path of a source the take-home
// wrote; any other path (but homebrew's) means it was added at home. The
// queries keep the literal 'take-home', like 'homebrew'.
const takeHomeSourcePath = "take-home"

// homeImport opens a take-home file: a new home copy, or a replacement of the
// one already here after the player confirmed it.
func (a *App) homeImport(c *gin.Context) error {
	if !a.Player {
		a.fail(c, http.StatusNotFound, "Not found.")
		return nil
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxTakeHomeBytes)
	fromPhone := c.PostForm(formHandoff) == "1"
	var f interface {
		io.ReaderAt
		io.Closer
	}
	var size int64
	if fromPhone {
		hf, n, err := a.handoffReceived()
		if err != nil {
			return err
		}
		f, size = hf, n
	} else {
		fh, err := c.FormFile("file")
		if err != nil {
			return userf("Choose a take-home file (.zip).")
		}
		if f, err = fh.Open(); err != nil {
			return err
		}
		size = fh.Size
	}
	defer func() { _ = f.Close() }()
	file, err := export.ReadTakeHome(f, size, export.KindTakeHome)
	if err != nil {
		return userf("That is not a take-home file: %w", err)
	}
	cf, err := a.checkTakeHome(file)
	if err != nil {
		return err
	}

	ctx := c.Request.Context()
	m := file.Manifest
	// One table per player app: two tables' books and homebrew would share rows.
	// Checked here before any file is saved, and again in the transaction.
	if err := oneTable(ctx, a.Store.Queries, m.ServerID); err != nil {
		return err
	}
	replace := c.PostForm(formReplace) == "1"
	if old, err := a.Store.GetHomeCopyByTable(ctx, sqlc.GetHomeCopyByTableParams{ServerID: m.ServerID, TableCharacterID: m.CharacterID}); err == nil && !replace {
		var om export.HomeManifest
		if err := json.Unmarshal([]byte(old.ManifestJson), &om); err != nil {
			return err
		}
		lost, err := lostHomeSources(ctx, a.Store.Queries, old.CharacterID, file.Sources)
		if err != nil {
			return err
		}
		a.render(c, http.StatusOK, views.HomeReplaceConfirm(views.HomeConfirmView{Name: cf.Name, Old: om.CreatedAt, New: m.CreatedAt, Sources: lost}))
		return nil
	} else if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// Files first: a row needs the file's size and type, and storeFile would
	// open a second connection while the transaction holds the write lock.
	var saved []string
	defer func() {
		for _, p := range saved {
			_ = os.Remove(filepath.Join(a.attachmentsDir(), p))
		}
	}()
	save := func(zipName, display string) (savedFile, error) {
		body, err := file.OpenFile(zipName)
		if err != nil {
			return savedFile{}, userError{err}
		}
		name, mime, size, err := a.saveFile(bytes.NewReader(body))
		if err != nil {
			return savedFile{}, err
		}
		saved = append(saved, name)
		return savedFile{path: name, name: display, mime: mime, size: size}, nil
	}
	lookFiles := make([]savedFile, len(file.Looks))
	for i, l := range file.Looks {
		if l.File == "" {
			continue
		}
		if lookFiles[i], err = save(l.File, cmp.Or(l.FileName, l.Name)); err != nil {
			return err
		}
	}
	blockFiles := map[int64][]savedFile{}
	for _, n := range file.Notes {
		for _, b := range n.Blocks {
			for _, bf := range b.Files {
				sf, err := save(bf.File, bf.Name)
				if err != nil {
					return err
				}
				blockFiles[b.ID] = append(blockFiles[b.ID], sf)
			}
		}
	}

	var newID int64
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		newID, err = importTakeHome(ctx, q, file, cf, lookFiles, blockFiles, replace)
		return err
	})
	if err != nil {
		return err
	}
	saved = nil // rows own the files now
	if fromPhone {
		_ = f.Close() // Windows cannot delete an open file
		a.stopHandoff()
	}

	if _, err := UpgradeSheets(ctx, a.Store); err != nil {
		slog.Warn("sheet upgrade after take-home import", "err", err)
	}
	ch, err := a.Store.GetCharacter(ctx, newID)
	if err != nil {
		return err
	}
	if err := a.dropSession(c); err != nil {
		return err
	}
	if err := a.login(c, ch); err != nil {
		return err
	}
	redirect(c, route.Path(route.Character, ch.ID))
	return nil
}

// lostHomeSources are the sources added at home that the home copy's campaign
// allows and the new file does not list: a replace deletes that campaign, so
// the new copy leaves them off.
func lostHomeSources(ctx context.Context, q *sqlc.Queries, homeCharacterID int64, listed []export.TakeHomeSource) ([]string, error) {
	ch, err := q.GetCharacterAny(ctx, homeCharacterID)
	if err != nil {
		return nil, err
	}
	allowed, err := q.ListCampaignSources(ctx, ch.CampaignID)
	if err != nil {
		return nil, err
	}
	home, err := q.ListHomeSourceKeys(ctx)
	if err != nil {
		return nil, err
	}
	var lost []string
	for _, k := range allowed {
		if slices.Contains(home, k) && !slices.ContainsFunc(listed, func(s export.TakeHomeSource) bool { return s.SourceKey == k }) {
			lost = append(lost, k)
		}
	}
	return lost, nil
}

// oneTable refuses a file from another table than the characters already here.
func oneTable(ctx context.Context, q *sqlc.Queries, serverID string) error {
	other, err := q.HasOtherServerHomeCopy(ctx, serverID)
	if err != nil {
		return err
	}
	if other {
		return userf("The characters here come from another GM's table. Unzip the player app into a second folder for this GM and open the file there.")
	}
	return nil
}

// importTakeHome writes one take-home file as a home campaign with its
// character, in one transaction, and returns the new character's id. With
// replace, the copy of the same table character is deleted first, with its own
// homebrew (a character's deletion only unowns it).
func importTakeHome(ctx context.Context, q *sqlc.Queries, file *export.TakeHomeFile, cf export.CharacterFile, lookFiles []savedFile, blockFiles map[int64][]savedFile, replace bool) (int64, error) {
	m := file.Manifest
	if err := oneTable(ctx, q, m.ServerID); err != nil {
		return 0, err
	}
	if old, err := q.GetHomeCopyByTable(ctx, sqlc.GetHomeCopyByTableParams{ServerID: m.ServerID, TableCharacterID: m.CharacterID}); err == nil {
		if !replace {
			return 0, userError{errors.New("this character was opened on this computer meanwhile, try again")}
		}
		oldChar, err := q.GetCharacterAny(ctx, old.CharacterID)
		if err != nil {
			return 0, err
		}
		if err := q.DeleteHomebrewByOwner(ctx, sqlc.DeleteHomebrewByOwnerParams{OwnerCharacterID: nullID(old.CharacterID), CampaignID: oldChar.CampaignID}); err != nil {
			return 0, err
		}
		if err := q.DeleteCampaign(ctx, oldChar.CampaignID); err != nil {
			return 0, err
		}
	} else if !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}

	// The compendium, shared by every home copy.
	for _, s := range file.Sources {
		if err := q.UpsertTakeHomeSource(ctx, sqlc.UpsertTakeHomeSourceParams{
			SourceKey: s.SourceKey, Edition: s.Edition, License: s.License, Attribution: s.Attribution, Published: s.Published, Priority: s.Priority,
		}); err != nil {
			return 0, err
		}
	}
	// Sources added at home keep their place below the table's.
	if err := q.SinkHomeSources(ctx); err != nil {
		return 0, err
	}
	homeKeys, err := q.ListHomeSourceKeys(ctx)
	if err != nil {
		return 0, err
	}
	atHome := make(map[string]bool, len(homeKeys))
	for _, k := range homeKeys {
		atHome[k] = true
	}
	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: m.Campaign.Name, System: m.Campaign.System, Edition: m.Campaign.Edition})
	if err != nil {
		return 0, err
	}
	if err := q.SetCampaignOpenJoin(ctx, sqlc.SetCampaignOpenJoinParams{OpenJoin: false, ID: camp.ID}); err != nil {
		return 0, err
	}
	for _, s := range file.Sources {
		if _, err := q.AllowCampaignSource(ctx, sqlc.AllowCampaignSourceParams{CampaignID: camp.ID, SourceKey: s.SourceKey}); err != nil {
			return 0, err
		}
	}

	ch, err := q.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{
		CampaignID: camp.ID, Name: cf.Name, System: cf.System, Hp: cf.HP, MaxHp: cf.MaxHP, TempHp: cf.TempHP, SheetJson: string(cf.Sheet),
	})
	if err != nil {
		return 0, err
	}
	maps := export.HomeIDMaps{
		LevelUps: map[int64]int64{}, Images: map[int64]int64{}, Attachments: map[int64]int64{},
		Notes: map[int64]int64{}, Blocks: map[int64]int64{}, Homebrew: map[int64]int64{},
	}

	// Entries: other sources' as they are; homebrew scoped to this campaign.
	for _, e := range file.Entries {
		if atHome[e.SourceKey] {
			continue // the home download is that source's copy
		}
		if e.SourceKey != homebrewSource {
			if err := q.UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{
				Kind: e.Kind, Slug: e.Slug, Name: e.Name, SourceKey: e.SourceKey, DataJson: string(e.Data), BodyText: e.BodyText, ContentHash: e.ContentHash,
			}); err != nil {
				return 0, err
			}
			continue
		}
		id, err := q.UpsertSharedHomebrew(ctx, sqlc.UpsertSharedHomebrewParams{
			Kind: e.Kind, Slug: e.Slug, Name: e.Name, DataJson: string(e.Data), BodyText: e.BodyText, ContentHash: e.ContentHash,
		})
		if err != nil {
			return 0, err
		}
		if err := q.LinkEntryCampaign(ctx, sqlc.LinkEntryCampaignParams{EntryID: id, CampaignID: camp.ID}); err != nil {
			return 0, err
		}
	}
	for _, e := range file.Homebrew {
		id, err := q.UpsertHomeHomebrew(ctx, sqlc.UpsertHomeHomebrewParams{
			Kind: e.Kind, Slug: e.Slug, Name: e.Name, DataJson: string(e.Data), BodyText: e.BodyText, ContentHash: e.ContentHash, OwnerCharacterID: nullID(ch.ID),
		})
		if err != nil {
			return 0, err
		}
		if err := q.LinkEntryCampaign(ctx, sqlc.LinkEntryCampaignParams{EntryID: id, CampaignID: camp.ID}); err != nil {
			return 0, err
		}
		maps.Homebrew[id] = e.ID
	}

	for _, l := range file.LevelUps {
		id, err := q.CreateLevelUpAt(ctx, sqlc.CreateLevelUpAtParams{CharacterID: ch.ID, MaxHp: l.MaxHP, SheetJson: string(l.Sheet), CreatedAt: l.CreatedAt})
		if err != nil {
			return 0, err
		}
		maps.LevelUps[id] = l.ID
	}

	if err := importLooks(ctx, q, file.Looks, lookFiles, ch, camp.ID, &maps); err != nil {
		return 0, err
	}
	if err := importNotes(ctx, q, file.Notes, blockFiles, ch, camp.ID, &maps); err != nil {
		return 0, err
	}

	manifest, err := json.Marshal(export.HomeManifest{TakeHomeManifest: m, Maps: maps})
	if err != nil {
		return 0, err
	}
	err = q.CreateHomeCopy(ctx, sqlc.CreateHomeCopyParams{CharacterID: ch.ID, ServerID: m.ServerID, TableCharacterID: m.CharacterID, ManifestJson: string(manifest)})
	return ch.ID, err
}

// importLooks makes each look an image of the home campaign, linked to the
// character with the flags it had at the table.
func importLooks(ctx context.Context, q *sqlc.Queries, looks []export.TakeHomeLook, files []savedFile, ch sqlc.Character, campaignID int64, maps *export.HomeIDMaps) error {
	homeIDs := make([]int64, len(looks))
	for i, l := range looks {
		var att sql.NullInt64
		if l.File != "" {
			a, err := q.CreateAttachment(ctx, sqlc.CreateAttachmentParams{CampaignID: campaignID, Path: files[i].path, Name: files[i].name, Mime: files[i].mime, Size: files[i].size})
			if err != nil {
				return err
			}
			att = nullID(a.ID)
			maps.Attachments[a.ID] = l.AttachmentID
		}
		img, err := q.CreateImage(ctx, sqlc.CreateImageParams{
			CampaignID: campaignID, Name: cmp.Or(l.Name, ch.Name), Url: sql.NullString{String: l.URL, Valid: l.URL != ""}, AttachmentID: att, Description: l.Description,
		})
		if err != nil {
			return err
		}
		maps.Images[img.ID] = l.ImageID
		homeIDs[i] = img.ID
		if err := q.LinkCharacterImage(ctx, sqlc.LinkCharacterImageParams{CharacterID: ch.ID, ImageID: img.ID}); err != nil {
			return err
		}
	}
	// LinkCharacterImage made the first look current: set the real flags.
	if err := q.ClearCharacterCurrentImage(ctx, ch.ID); err != nil {
		return err
	}
	for i, l := range looks {
		if err := q.SetCharacterImageFlags(ctx, sqlc.SetCharacterImageFlagsParams{
			IsRevealed: l.IsRevealed, IsCurrent: l.IsCurrent, CharacterID: ch.ID, ImageID: homeIDs[i],
		}); err != nil {
			return err
		}
	}
	return nil
}

// importNotes recreates the character's notes and blocks with fresh
// collaborative state (version 0), their files attached and each picture in a
// document pointed at its new file. A block may show another block's picture
// (copied between notes), so every block is rewritten with the map of all files.
func importNotes(ctx context.Context, q *sqlc.Queries, notes []export.TakeHomeNote, files map[int64][]savedFile, ch sqlc.Character, campaignID int64, maps *export.HomeIDMaps) error {
	byTable := map[int64]int64{}
	docs := map[int64]json.RawMessage{} // home block id → the document as the table had it
	for _, n := range notes {
		noteID, err := q.CreateNoteAt(ctx, sqlc.CreateNoteAtParams{
			CampaignID: campaignID, Kind: n.Kind, Title: n.Title, OwnerCharacterID: nullID(ch.ID), CreatedAt: n.CreatedAt, UpdatedAt: n.UpdatedAt,
		})
		if err != nil {
			return err
		}
		maps.Notes[noteID] = n.ID
		for _, b := range n.Blocks {
			var author sql.NullInt64
			if b.ByCharacter {
				author = nullID(ch.ID)
			}
			blockID, err := q.CreateTakeHomeBlock(ctx, sqlc.CreateTakeHomeBlockParams{
				NoteID: noteID, Position: b.Position, Visibility: b.Visibility, PlayersCanEdit: b.PlayersCanEdit,
				DocJson: string(b.Doc), Text: b.Text, AuthorCharacterID: author,
			})
			if err != nil {
				return err
			}
			maps.Blocks[blockID] = b.ID
			docs[blockID] = b.Doc
			if b.Visibility == enum.VisListed {
				if err := q.AddBlockAudience(ctx, sqlc.AddBlockAudienceParams{BlockID: blockID, CharacterID: ch.ID}); err != nil {
					return err
				}
			}
			for i, bf := range b.Files {
				sf := files[b.ID][i]
				a, err := q.CreateAttachment(ctx, sqlc.CreateAttachmentParams{CampaignID: campaignID, BlockID: nullID(blockID), Path: sf.path, Name: sf.name, Mime: sf.mime, Size: sf.size})
				if err != nil {
					return err
				}
				byTable[bf.AttachmentID] = a.ID
				maps.Attachments[a.ID] = bf.AttachmentID
			}
		}
	}
	if len(byTable) == 0 {
		return nil
	}
	for blockID, d := range docs {
		doc, err := rewriteDocAttachments(d, byTable)
		if err != nil {
			return err
		}
		if err := q.SetTakeHomeBlockDoc(ctx, sqlc.SetTakeHomeBlockDocParams{DocJson: string(doc), ID: blockID}); err != nil {
			return err
		}
	}
	return nil
}

// rewriteDocAttachments points a ProseMirror document's pictures at the new
// attachments (table id to home id): an image's src and a picture mention's id.
// It walks the JSON, so text that only looks like an address stays as typed.
func rewriteDocAttachments(doc json.RawMessage, ids map[int64]int64) (json.RawMessage, error) {
	var v any
	if err := json.Unmarshal(doc, &v); err != nil {
		return nil, err
	}
	oldSrc, oldID := map[string]string{}, map[string]string{}
	for table, home := range ids {
		oldSrc[route.Path(route.Attachment, table)] = route.Path(route.Attachment, home)
		oldID[fmt.Sprintf("%s%d", attachmentMention, table)] = fmt.Sprintf("%s%d", attachmentMention, home)
	}
	var walk func(any)
	walk = func(n any) {
		switch n := n.(type) {
		case map[string]any:
			if attrs, ok := n["attrs"].(map[string]any); ok {
				if s, ok := attrs["src"].(string); ok && oldSrc[s] != "" {
					attrs["src"] = oldSrc[s]
				}
				if s, ok := attrs["id"].(string); ok && oldID[s] != "" {
					attrs["id"] = oldID[s]
				}
			}
			if content, ok := n["content"]; ok {
				walk(content)
			}
		case []any:
			for _, c := range n {
				walk(c)
			}
		}
	}
	walk(v)
	return json.Marshal(v)
}

// attachmentMention starts the id of a picture mention node.
const attachmentMention = "attachment:"
