package web

// The player app's side of taking a character home: list the characters on
// this computer, open a take-home zip and log the browser in as its character.
//
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
	"unicode/utf8"

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
	rows, err := a.Store.ListHomeCharacters(ctx)
	if err != nil {
		return err
	}
	keys, err := a.Store.ListHomeSourceKeys(ctx)
	if err != nil {
		return err
	}
	byID := map[int64]sqlc.ListHomeCharactersRow{}
	isFallback := map[int64]bool{}
	for _, r := range rows {
		byID[r.ID] = r
		if r.FallbackCharacterID.Valid {
			isFallback[r.FallbackCharacterID.Int64] = true
		}
	}
	var table, onlyHere []views.HomeCopyView
	for _, r := range rows {
		if isFallback[r.ID] {
			continue // listed under its character
		}
		v := views.HomeCopyView{CharacterID: r.ID, CampaignID: r.CampaignID, Name: r.Name, Campaign: r.CampaignName, Current: w.character != nil && w.character.ID == r.ID}
		if r.ManifestJson == "" {
			onlyHere = append(onlyHere, v)
			continue
		}
		var m export.HomeManifest
		if err := json.Unmarshal([]byte(r.ManifestJson), &m); err != nil {
			return err
		}
		v.FromTable, v.TakenAt = true, m.CreatedAt
		if fb, ok := byID[r.FallbackCharacterID.Int64]; r.FallbackCharacterID.Valid && ok {
			v.Fallback = &views.HomeFallbackView{CharacterID: fb.ID, Name: fb.Name, Current: w.character != nil && w.character.ID == fb.ID}
		} else if v.HomeSources, err = homeBuildSources(r.SheetJson, keys); err != nil {
			return err
		}
		table = append(table, v)
	}
	starters, err := a.homeStarters(ctx)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.HomePage(a.layout(w, "Characters", views.NavAtHome, nil, false), table, onlyHere, starters, a.handoffView()))
	return nil
}

// homePlay switches this browser to another character on this computer, with
// no approval: they are all the player's own.
func (a *App) homePlay(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	ch, err := a.Store.GetCharacter(c.Request.Context(), id)
	if err != nil {
		return err
	}
	if err := a.loginAs(c, ch); err != nil {
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
	cf, err := checkFileContent(file)
	if errors.Is(err, errNewerSheet) {
		return cf, userf("This file comes from a newer app (%s); this player app is %s. Get the matching app from your GM.", m.AppVersion, a.Version)
	}
	return cf, err
}

// checkFileContent is checkTakeHome's part that looks at what the file holds
// (the table's check of a bring-back runs it too): sheets (errNewerSheet when a
// newer app wrote one, for the caller to word), the character, notes, looks and
// sources. It stores the character's and the fallback's sheets as parsed.
func checkFileContent(file *export.TakeHomeFile) (export.CharacterFile, error) {
	cf := file.Character
	for _, s := range fileSheets(file) {
		switch err := checkSheetFile(s); {
		case errors.Is(err, errNewerSheet):
			return cf, err
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
	if fb := file.Fallback; fb != nil {
		fcf := export.CharacterFile{Name: cf.Name, System: cf.System, MaxHP: fb.MaxHP, Sheet: fb.Sheet}
		if err := normalizeCharacterFile(&fcf); err != nil {
			return cf, userError{fmt.Errorf("the fallback build is not valid: %w", err)}
		}
		fb.Sheet = fcf.Sheet
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

// fileSheets are the sheets a file holds: the character's, its level-ups', and
// the fallback's with its level-ups'. Phase 4 checks a bring-back with it too.
func fileSheets(file *export.TakeHomeFile) []json.RawMessage {
	sheets := []json.RawMessage{file.Character.Sheet}
	for _, l := range file.LevelUps {
		sheets = append(sheets, l.Sheet)
	}
	if fb := file.Fallback; fb != nil {
		sheets = append(sheets, fb.Sheet)
		for _, l := range fb.LevelUps {
			sheets = append(sheets, l.Sheet)
		}
	}
	return sheets
}

// checkFileSources checks the compendium part of a take-home or bring-back
// file: sources with a real key, and entries of a listed source (or homebrew,
// which a bring-back or new character's file never lists among its entries: its
// homebrew is separate). The table calls it for both.
func checkFileSources(file *export.TakeHomeFile) error {
	known := map[string]bool{}
	for _, s := range file.Sources {
		if s.SourceKey == "" || s.SourceKey == homebrewSource {
			return userError{errors.New("the file lists a source that is not valid")}
		}
		known[s.SourceKey] = true
	}
	for _, e := range file.Entries {
		if k := file.Manifest.Kind; e.SourceKey == homebrewSource && k != export.KindTakeHome && k != export.KindStarter {
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
	kind, err := export.ManifestKind(f, size)
	if err != nil {
		return userf("That is not a take-home file: %w", err)
	}
	if kind == export.KindStarter {
		return a.homeStarterImport(c, f, size, fromPhone)
	}
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
	old, err := a.Store.GetHomeCopyByTable(ctx, sqlc.GetHomeCopyByTableParams{ServerID: m.ServerID, TableCharacterID: m.CharacterID})
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	copyHere := err == nil
	original, err := liveMadeAtHome(ctx, a.Store.Queries, m.MadeAtHome)
	if err != nil {
		return err
	}
	if (copyHere || original != nil) && !replace {
		v := views.HomeConfirmView{Name: cf.Name, New: m.CreatedAt}
		if copyHere {
			var om export.HomeManifest
			if err := json.Unmarshal([]byte(old.ManifestJson), &om); err != nil {
				return err
			}
			v.CopyHere, v.Old = true, om.CreatedAt
			if v.Sources, err = lostHomeSources(ctx, a.Store.Queries, old.CharacterID, file.Sources); err != nil {
				return err
			}
			if old.FallbackCharacterID.Valid {
				if fb, err := a.Store.GetCharacter(ctx, old.FallbackCharacterID.Int64); err == nil {
					v.Fallback = fb.Name
				}
			}
		}
		if original != nil {
			v.MadeAtHome = original.Name
			lost, err := lostHomeSources(ctx, a.Store.Queries, original.CharacterID, file.Sources)
			if err != nil {
				return err
			}
			v.Sources = slices.Compact(slices.Sorted(slices.Values(slices.Concat(v.Sources, lost))))
		}
		a.render(c, http.StatusOK, views.HomeReplaceConfirm(v))
		return nil
	}

	// Files first: a row needs the file's size and type, and storeFile would
	// open a second connection while the transaction holds the write lock.
	var saved []string
	defer func() {
		for _, p := range saved {
			_ = os.Remove(filepath.Join(a.attachmentsDir(), p))
		}
	}()
	lookFiles, blockFiles, err := a.saveHomeFiles(file.TakeHome, func(zipName string) ([]byte, error) {
		body, err := file.OpenFile(zipName)
		if err != nil {
			return nil, userError{err}
		}
		return body, nil
	}, &saved)
	if err != nil {
		return err
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

// saveHomeFiles writes the files of th's looks and note blocks (read by open
// from their zip name) to the attachments folder as new files. Their names go
// on *saved, which the caller removes unless rows own them.
func (a *App) saveHomeFiles(th export.TakeHome, open func(zipName string) ([]byte, error), saved *[]string) ([]savedFile, map[int64][]savedFile, error) {
	save := func(zipName, display string) (savedFile, error) {
		body, err := open(zipName)
		if err != nil {
			return savedFile{}, err
		}
		name, mime, size, err := a.saveFile(bytes.NewReader(body))
		if err != nil {
			return savedFile{}, err
		}
		*saved = append(*saved, name)
		return savedFile{path: name, name: display, mime: mime, size: size}, nil
	}
	lookFiles := make([]savedFile, len(th.Looks))
	for i, l := range th.Looks {
		if l.File == "" {
			continue
		}
		var err error
		if lookFiles[i], err = save(l.File, cmp.Or(l.FileName, l.Name)); err != nil {
			return nil, nil, err
		}
	}
	blockFiles := map[int64][]savedFile{}
	for _, n := range th.Notes {
		for _, b := range n.Blocks {
			for _, bf := range b.Files {
				sf, err := save(bf.File, bf.Name)
				if err != nil {
					return nil, nil, err
				}
				blockFiles[b.ID] = append(blockFiles[b.ID], sf)
			}
		}
	}
	return lookFiles, blockFiles, nil
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
// replace, the copy of the same table character and the made-at-home character
// the file stands in for are deleted first, with their own homebrew (a
// character's deletion only unowns it).
func importTakeHome(ctx context.Context, q *sqlc.Queries, file *export.TakeHomeFile, cf export.CharacterFile, lookFiles []savedFile, blockFiles map[int64][]savedFile, replace bool) (int64, error) {
	m := file.Manifest
	if err := oneTable(ctx, q, m.ServerID); err != nil {
		return 0, err
	}
	if old, err := q.GetHomeCopyByTable(ctx, sqlc.GetHomeCopyByTableParams{ServerID: m.ServerID, TableCharacterID: m.CharacterID}); err == nil {
		if !replace {
			return 0, userError{errors.New("this character was opened on this computer meanwhile, try again")}
		}
		if err := deleteHomeCharacter(ctx, q, old.CharacterID, old.FallbackCharacterID); err != nil {
			return 0, err
		}
	} else if !errors.Is(err, sql.ErrNoRows) {
		return 0, err
	}
	// The made-at-home original this file stands in for goes the same way.
	original, err := liveMadeAtHome(ctx, q, m.MadeAtHome)
	if err != nil {
		return 0, err
	}
	if original != nil {
		if !replace {
			return 0, userError{errors.New("this character was made at home, confirm the replacement and try again")}
		}
		if err := deleteHomeCharacter(ctx, q, original.CharacterID, sql.NullInt64{}); err != nil {
			return 0, err
		}
	}

	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: m.Campaign.Name, System: m.Campaign.System, Edition: m.Campaign.Edition})
	if err != nil {
		return 0, err
	}
	if err := q.SetCampaignOpenJoin(ctx, sqlc.SetCampaignOpenJoinParams{OpenJoin: false, ID: camp.ID}); err != nil {
		return 0, err
	}
	if err := importHomeCompendium(ctx, q, file, camp.ID); err != nil {
		return 0, err
	}

	maps := newHomeIDMaps()
	ch, err := createHomeCharacter(ctx, q, camp.ID, cf, file.TakeHome, lookFiles, blockFiles, &maps, homeMentions)
	if err != nil {
		return 0, err
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

	manifest, err := json.Marshal(export.HomeManifest{TakeHomeManifest: m, Maps: maps})
	if err != nil {
		return 0, err
	}
	err = q.CreateHomeCopy(ctx, sqlc.CreateHomeCopyParams{CharacterID: ch.ID, ServerID: m.ServerID, TableCharacterID: m.CharacterID, ManifestJson: string(manifest)})
	return ch.ID, err
}

func newHomeIDMaps() export.HomeIDMaps {
	return export.HomeIDMaps{
		LevelUps: map[int64]int64{}, Images: map[int64]int64{}, Attachments: map[int64]int64{},
		Notes: map[int64]int64{}, Blocks: map[int64]int64{}, Homebrew: map[int64]int64{},
	}
}

// createHomeCharacter writes cf as a character of the campaign with th's
// level-ups, looks and notes (their files already saved), recording the home
// ids in maps. mentions (nil: none) rewrites the notes' mention ids once every
// note exists.
func createHomeCharacter(ctx context.Context, q *sqlc.Queries, campaignID int64, cf export.CharacterFile, th export.TakeHome, lookFiles []savedFile, blockFiles map[int64][]savedFile, maps *export.HomeIDMaps, mentions mentionMaker) (sqlc.Character, error) {
	ch, err := q.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{
		CampaignID: campaignID, Name: cf.Name, System: cf.System, Hp: cf.HP, MaxHp: cf.MaxHP, TempHp: cf.TempHP, SheetJson: string(cf.Sheet),
	})
	if err != nil {
		return ch, err
	}
	for _, l := range th.LevelUps {
		id, err := q.CreateLevelUpAt(ctx, sqlc.CreateLevelUpAtParams{CharacterID: ch.ID, MaxHp: l.MaxHP, SheetJson: string(l.Sheet), CreatedAt: l.CreatedAt})
		if err != nil {
			return ch, err
		}
		maps.LevelUps[id] = l.ID
	}
	if err := importLooks(ctx, q, th.Looks, lookFiles, ch, campaignID, maps); err != nil {
		return ch, err
	}
	return ch, importNotes(ctx, q, th.Notes, blockFiles, ch, campaignID, maps, mentions)
}

// cloneHomeCampaign makes a home campaign with the same sources and homebrew
// links as another, for a copy or new character that must outlive a replace
// of the original's campaign.
func cloneHomeCampaign(ctx context.Context, q *sqlc.Queries, fromID int64) (int64, error) {
	from, err := q.GetCampaign(ctx, fromID)
	if err != nil {
		return 0, err
	}
	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: from.Name, System: from.System, Edition: from.Edition})
	if err != nil {
		return 0, err
	}
	if err := q.SetCampaignOpenJoin(ctx, sqlc.SetCampaignOpenJoinParams{OpenJoin: false, ID: camp.ID}); err != nil {
		return 0, err
	}
	if err := q.CopyCampaignSources(ctx, sqlc.CopyCampaignSourcesParams{ToCampaignID: camp.ID, FromCampaignID: fromID}); err != nil {
		return 0, err
	}
	return camp.ID, q.CopyCampaignEntryLinks(ctx, sqlc.CopyCampaignEntryLinksParams{ToCampaignID: camp.ID, FromCampaignID: fromID})
}

// maxCharacterName is the longest character name (the name form's max=100).
const maxCharacterName = 100

// homeName is base with suffix, cut rune-wise so it fits a character name.
func homeName(base, suffix string) string {
	r := []rune(base)
	if keep := maxCharacterName - utf8.RuneCountInString(suffix); len(r) > keep {
		r = r[:keep]
	}
	return string(r) + suffix
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
// mentions (nil: none) rewrites note and image mentions once every note exists.
func importNotes(ctx context.Context, q *sqlc.Queries, notes []export.TakeHomeNote, files map[int64][]savedFile, ch sqlc.Character, campaignID int64, maps *export.HomeIDMaps, mentions mentionMaker) error {
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
	if len(byTable) == 0 && mentions == nil {
		return nil
	}
	var mention mentionIDs
	if mentions != nil {
		mention = mentions(maps)
	}
	for blockID, d := range docs {
		doc, err := rewriteDocAttachments(d, byTable)
		if err != nil {
			return err
		}
		if mention != nil {
			if doc, err = rewriteDocMentions(doc, mention); err != nil {
				return err
			}
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

// deleteHomeCharacter removes a character here with its own homebrew (and its
// fallback's) and the home campaign it lives in; a character's deletion alone
// would only unown the homebrew.
func deleteHomeCharacter(ctx context.Context, q *sqlc.Queries, characterID int64, fallbackID sql.NullInt64) error {
	ch, err := q.GetCharacterAny(ctx, characterID)
	if err != nil {
		return err
	}
	if err := q.DeleteHomebrewByOwner(ctx, sqlc.DeleteHomebrewByOwnerParams{OwnerCharacterID: nullID(characterID), CampaignID: ch.CampaignID}); err != nil {
		return err
	}
	if fallbackID.Valid {
		if err := q.DeleteHomebrewByOwner(ctx, sqlc.DeleteHomebrewByOwnerParams{OwnerCharacterID: fallbackID, CampaignID: ch.CampaignID}); err != nil {
			return err
		}
	}
	return q.DeleteCampaign(ctx, ch.CampaignID)
}

// liveMadeAtHome is the character on this computer that a file's made-at-home
// token names, nil for none (or a deleted one).
func liveMadeAtHome(ctx context.Context, q *sqlc.Queries, token string) (*sqlc.GetNewCharacterByTokenRow, error) {
	if token == "" {
		return nil, nil //nolint:nilnil // nil means no such character
	}
	h, err := q.GetNewCharacterByToken(ctx, token)
	if errors.Is(err, sql.ErrNoRows) || err == nil && h.DeletedAt.Valid {
		return nil, nil //nolint:nilnil // nil means no such character
	}
	return &h, err
}

// importHomeCompendium writes the compendium part of a take-home or starter
// file for the home campaign: its sources (ticked for the campaign) and
// entries, homebrew linked to the campaign only. The compendium is shared by
// every home campaign.
func importHomeCompendium(ctx context.Context, q *sqlc.Queries, file *export.TakeHomeFile, campaignID int64) error {
	for _, s := range file.Sources {
		if err := q.UpsertTakeHomeSource(ctx, sqlc.UpsertTakeHomeSourceParams{
			SourceKey: s.SourceKey, Edition: s.Edition, License: s.License, Attribution: s.Attribution, Published: s.Published, Priority: s.Priority,
		}); err != nil {
			return err
		}
	}
	// Sources added at home keep their place below the table's.
	if err := q.SinkHomeSources(ctx); err != nil {
		return err
	}
	homeKeys, err := q.ListHomeSourceKeys(ctx)
	if err != nil {
		return err
	}
	for _, s := range file.Sources {
		if _, err := q.AllowCampaignSource(ctx, sqlc.AllowCampaignSourceParams{CampaignID: campaignID, SourceKey: s.SourceKey}); err != nil {
			return err
		}
	}
	// Entries: other sources' as they are; homebrew scoped to this campaign.
	for _, e := range file.Entries {
		if slices.Contains(homeKeys, e.SourceKey) {
			continue // the home download is that source's copy
		}
		if e.SourceKey != homebrewSource {
			if err := q.UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{
				Kind: e.Kind, Slug: e.Slug, Name: e.Name, SourceKey: e.SourceKey, DataJson: string(e.Data), BodyText: e.BodyText, ContentHash: e.ContentHash,
			}); err != nil {
				return err
			}
			continue
		}
		id, err := q.UpsertSharedHomebrew(ctx, sqlc.UpsertSharedHomebrewParams{
			Kind: e.Kind, Slug: e.Slug, Name: e.Name, DataJson: string(e.Data), BodyText: e.BodyText, ContentHash: e.ContentHash,
		})
		if err != nil {
			return err
		}
		if err := q.LinkEntryCampaign(ctx, sqlc.LinkEntryCampaignParams{EntryID: id, CampaignID: campaignID}); err != nil {
			return err
		}
	}
	return nil
}
