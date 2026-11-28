package web

// The table's side of bringing a character back: a bring-back zip is uploaded
// (by the player's device, waiting in the GM's popup, or by the GM at once),
// checked, reviewed as a three-way merge of the sheet and accepted or rejected.
// This is the one place a command is not a delta: accepting writes the whole
// merged character, after the GM has seen every group of it.
//
// shortcut: pending files are lost on restart; persist them if reviews ever span restarts.

import (
	"cmp"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

const (
	// HomeFilesDir is the folder under DATA_DIR holding the zips waiting for the GM.
	HomeFilesDir = "home-files"
	// maxHomeFiles bounds the files waiting, like MaxPending does for requests.
	maxHomeFiles = 10
	// bringBackSourcePath is compendium_sources.path of a source a bring-back
	// brought (its entries come from the file, not a download).
	bringBackSourcePath = "bring-back"
	// formFile is the upload field, formWas the stamp the review was made against.
	formFile = "file"
	formWas  = "was"
	// formPage marks a Reject sent from the review page, not the popup.
	formPage = "page"
)

// homeFile is an uploaded file waiting for the GM.
type homeFile struct {
	id            string
	kind          export.Kind
	campaignID    int64
	campaignName  string
	characterID   int64
	characterName string
	byPlayer      bool
	savedAt       time.Time
	added         time.Time
	path          string
	size          int64
	// claimed is set while an Add runs, so a second Add of the file waits its turn.
	claimed bool
}

// homeFiles holds the waiting files in memory, like csvImports; the zips are
// on disk in HomeFilesDir and go with their entry.
type homeFiles struct {
	mu sync.Mutex
	m  map[string]homeFile
}

// put stores f and returns its id. A file of the same kind for the same
// character replaces the older one (a new character's file has none, so it
// never does); beyond maxHomeFiles it is refused.
func (s *homeFiles) put(f homeFile) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m == nil {
		s.m = map[string]homeFile{}
	}
	var oldID string
	for id, old := range s.m {
		if f.characterID != 0 && old.characterID == f.characterID && old.kind == f.kind {
			oldID = id
		}
	}
	if oldID == "" && len(s.m) >= maxHomeFiles {
		return "", userf("Too many files are waiting for the GM. Ask the GM to review them first.")
	}
	if oldID != "" {
		_ = os.Remove(s.m[oldID].path)
		delete(s.m, oldID)
	}
	f.id, f.added = rand.Text(), time.Now()
	s.m[f.id] = f
	return f.id, nil
}

func (s *homeFiles) get(id string, campaignID int64) (homeFile, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	f, ok := s.m[id]
	return f, ok && f.campaignID == campaignID
}

// claim marks a waiting file as being added; false when it is gone or already claimed.
func (s *homeFiles) claim(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	f, ok := s.m[id]
	if !ok || f.claimed {
		return false
	}
	f.claimed = true
	s.m[id] = f
	return true
}

// release undoes claim (nothing once the file is dropped).
func (s *homeFiles) release(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if f, ok := s.m[id]; ok {
		f.claimed = false
		s.m[id] = f
	}
}

// drop forgets a file and deletes its zip.
func (s *homeFiles) drop(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if f, ok := s.m[id]; ok {
		_ = os.Remove(f.path)
		delete(s.m, id)
	}
}

// list is the files in the order they came.
func (s *homeFiles) list() []homeFile {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]homeFile, 0, len(s.m))
	for _, f := range s.m {
		out = append(out, f)
	}
	slices.SortFunc(out, func(a, b homeFile) int { return a.added.Compare(b.added) })
	return out
}

// bringBack is a checked bring-back file with the table's records it belongs to.
type bringBack struct {
	file *export.TakeHomeFile
	// cf is the home character as the table would store it.
	cf   export.CharacterFile
	take sqlc.TakeHome
	ch   sqlc.Character
}

// loadBringBack reads the zip at path and checks it against the table: its
// kind, server, character, open take-out and content. The review repeats it on
// every request, and Accept inside its transaction.
func (a *App) loadBringBack(
	ctx context.Context,
	q *sqlc.Queries,
	path string,
	characterID int64,
) (bringBack, error) {
	var bb bringBack
	f, err := os.Open(path) //nolint:gosec // a random name inside our own folder
	if err != nil {
		return bb, userf("The file is gone. Ask for it again.")
	}
	defer func() { _ = f.Close() }()
	st, err := f.Stat()
	if err != nil {
		return bb, err
	}
	file, err := export.ReadTakeHome(f, st.Size(), export.KindBringBack)
	if err != nil {
		return bb, userf("That is not a bring-back file: %w", err)
	}
	m := file.Manifest
	serverID, err := q.GetServerID(ctx)
	if err != nil {
		return bb, err
	}
	if m.ServerID != serverID {
		return bb, userf("This file belongs to another GM's table.")
	}
	if m.CharacterID != characterID {
		return bb, userf("This file is for another character.")
	}
	if bb.ch, err = q.GetCharacter(ctx, characterID); err != nil {
		return bb, err
	}
	if bb.take, err = q.GetTakeHomeByToken(
		ctx,
		m.Token,
	); errors.Is(err, sql.ErrNoRows) ||
		err == nil && bb.take.CharacterID != characterID {
		return bb, userf("This file's take-out is unknown here.")
	} else if err != nil {
		return bb, err
	}
	if bb.take.ReturnedAt.Valid {
		return bb, userf("This file was already brought back.")
	}
	if m.Campaign.System != gameSystem || file.Character.System != gameSystem {
		return bb, userf("This table plays %s only.", gameSystem)
	}
	switch bb.cf, err = checkFileContent(file); {
	case errors.Is(err, errNewerSheet):
		return bb, userf(
			"This file comes from a newer player app (%s); this server is %s and needs updating.",
			m.AppVersion,
			a.Version,
		)
	case err != nil:
		return bb, err
	}
	if err := checkFileEntries(file); err != nil {
		return bb, err
	}
	bb.file = file
	return bb, nil
}

// bringBackDialog is the dialog body that asks for the file.
func (a *App) bringBackDialog(c *gin.Context) error {
	w, ch, ok := a.takeHomeAccess(c)
	if !ok {
		return nil
	}
	a.render(
		c,
		http.StatusOK,
		views.BringBackBody(views.BringBackView{CharacterID: ch.ID, Name: ch.Name, ByGM: w.isGM()}),
	)
	return nil
}

// bringBackUpload takes a bring-back file. The GM goes straight to its review;
// a player's waits in the GM's popup.
func (a *App) bringBackUpload(c *gin.Context) error {
	w, ch, ok := a.takeHomeAccess(c)
	if !ok {
		return nil
	}
	path, size, err := a.receiveHomeFile(c, "a bring-back file")
	if err != nil {
		return err
	}
	kept := false
	defer func() {
		if !kept {
			_ = os.Remove(path)
		}
	}()

	ctx := c.Request.Context()
	bb, err := a.loadBringBack(ctx, a.Store.Queries, path, ch.ID)
	if err != nil {
		return err
	}
	camp, err := a.Store.GetCampaign(ctx, ch.CampaignID)
	if err != nil {
		return err
	}
	f := homeFile{
		kind:          export.KindBringBack,
		campaignID:    ch.CampaignID,
		campaignName:  camp.Name,
		characterID:   ch.ID,
		characterName: ch.Name,
		byPlayer:      !w.isGM(),
		savedAt:       bb.file.Manifest.SavedAt,
		path:          path,
		size:          size,
	}
	id, err := a.homeFiles.put(f)
	if err != nil {
		return err
	}
	kept = true
	a.publish(c, ch.CampaignID, hub.Topic{Kind: hub.TopicPending, ID: ch.CampaignID})
	if w.isGM() {
		redirect(c, route.Path(route.CampaignHomeFile, ch.CampaignID, id))
		return nil
	}
	a.render(c, http.StatusOK, views.BringBackSent(ch.Name))
	return nil
}

// receiveHomeFile saves the request's uploaded zip in HomeFilesDir and returns
// its path and size; what names the file in the message when none was sent.
// The caller removes the zip unless it keeps it.
func (a *App) receiveHomeFile(c *gin.Context, what string) (string, int64, error) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxTakeHomeBytes)
	fh, err := c.FormFile(formFile)
	if err != nil {
		return "", 0, userf("Choose %s (.zip).", what)
	}
	src, err := fh.Open()
	if err != nil {
		return "", 0, err
	}
	defer func() { _ = src.Close() }()

	dir := filepath.Join(a.DataDir, HomeFilesDir)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return "", 0, err
	}
	path := filepath.Join(dir, rand.Text()+".zip")
	//nolint:gosec // a random name inside our own folder
	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return "", 0, err
	}
	_, err = io.Copy(out, src)
	if cerr := out.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		_ = os.Remove(path)
		return "", 0, err
	}
	return path, fh.Size, nil
}

// homeFileAccess is the GM's access to one waiting file of the campaign.
func (a *App) homeFileAccess(c *gin.Context) (who, sqlc.Campaign, homeFile, bool) {
	var f homeFile
	campaignID, ok := a.idParam(c, "campaign")
	if !ok {
		return who{}, sqlc.Campaign{}, f, false
	}
	w, camp, ok := a.campaignAccess(c, campaignID)
	if !ok {
		return w, camp, f, false
	}
	if !w.isGM() {
		a.fail(c, http.StatusForbidden, "Only the GM can review a bring-back file.")
		return w, camp, f, false
	}
	if f, ok = a.homeFiles.get(c.Param("file"), campaignID); !ok {
		a.fail(c, http.StatusNotFound, "That file is no longer waiting.")
		return w, camp, f, false
	}
	return w, camp, f, true
}

// homeReview is a bring-back file read, checked and merged for one request.
type homeReview struct {
	// bb is set for a bring-back; a new character's file has only file and cf.
	bb bringBack
	// file is the file, whose home side sheets have the table's spell slugs.
	file     *export.TakeHomeFile
	cf       export.CharacterFile
	plan     bringBackPlan
	reviewed []export.TakeHomeSource
	levelUps []sqlc.LevelUp
	// x is the plan of looks, notes and homebrew.
	x homeExtras
	// was stamps the table the plan was made against.
	was string
}

// reviewHomeFile reviews a waiting file of either kind.
func (a *App) reviewHomeFile(
	ctx context.Context,
	q *sqlc.Queries,
	f homeFile,
	form func(string) string,
) (homeReview, error) {
	if f.kind == export.KindNewCharacter {
		return a.reviewNewCharacter(ctx, q, f, form)
	}
	return a.reviewBringBack(ctx, q, f, form)
}

// reviewBringBack re-reads the file and makes the plan for the picks form reads.
func (a *App) reviewBringBack(
	ctx context.Context,
	q *sqlc.Queries,
	f homeFile,
	form func(string) string,
) (homeReview, error) {
	var rv homeReview
	var err error
	if rv.bb, err = a.loadBringBack(ctx, q, f.path, f.characterID); err != nil {
		return rv, err
	}
	rv.file, rv.cf = rv.bb.file, rv.bb.cf
	var base export.TakeHomeBase
	if err := json.Unmarshal([]byte(rv.bb.take.BaseJson), &base); err != nil {
		return rv, err
	}
	allowed, err := q.ListCampaignSources(ctx, rv.bb.ch.CampaignID)
	if err != nil {
		return rv, err
	}
	if rv.levelUps, err = q.ListLevelUps(ctx, rv.bb.ch.ID); err != nil {
		return rv, err
	}
	if rv.x, err = planExtras(
		ctx,
		q,
		rv.bb.ch.CampaignID,
		rv.bb.ch.ID,
		rv.file,
		base,
		form,
	); err != nil {
		return rv, err
	}
	if err := applyRenames(rv.x.hb.Renames, &rv.cf, rv.file); err != nil {
		return rv, err
	}
	ch, cf := rv.bb.ch, rv.cf
	in := bringBackIn{
		Base: charState{
			Name:   cmp.Or(base.Character.Name, ch.Name),
			HP:     base.Character.HP,
			MaxHP:  base.Character.MaxHP,
			TempHP: base.Character.TempHP,
			Sheet:  string(base.Character.Sheet),
		},
		Table: charState{
			Name:   ch.Name,
			HP:     ch.Hp,
			MaxHP:  ch.MaxHp,
			TempHP: ch.TempHp,
			Sheet:  ch.SheetJson,
		},
		Home: charState{
			Name:   cf.Name,
			HP:     cf.HP,
			MaxHP:  cf.MaxHP,
			TempHP: cf.TempHP,
			Sheet:  string(cf.Sheet),
		},
		Form: form,
	}
	if fb := rv.file.Fallback; fb != nil {
		in.Fallback = &charState{Name: cf.Name, MaxHP: fb.MaxHP, Sheet: string(fb.Sheet)}
	}
	for _, s := range rv.file.Sources {
		if !slices.Contains(allowed, s.SourceKey) {
			rv.reviewed = append(rv.reviewed, s)
			in.Sources = append(in.Sources, reviewSource{Key: s.SourceKey, Label: sourceLabel(s)})
		}
	}
	if rv.plan, err = planBringBack(in); err != nil {
		return rv, err
	}
	rv.plan.Groups = append(rv.plan.Groups, rv.x.groups()...)
	ids := make([]int64, len(rv.levelUps))
	for i, l := range rv.levelUps {
		ids[i] = l.ID
	}
	sum := sha256.Sum256(
		fmt.Appendf(
			nil,
			"%q|%d|%d|%d|%q|%v|%s",
			ch.Name,
			ch.Hp,
			ch.MaxHp,
			ch.TempHp,
			ch.SheetJson,
			ids,
			rv.x.table.stamp(),
		),
	)
	rv.was = hex.EncodeToString(sum[:])
	return rv, nil
}

// view is the review as the page shows it.
func (rv homeReview) view(f homeFile) views.HomeFileView {
	p := rv.plan
	if f.kind == export.KindNewCharacter {
		return views.HomeFileView{
			CampaignID:   f.campaignID,
			FileID:       f.id,
			Name:         rv.cf.Name,
			NewCharacter: true,
			LevelTo:      p.LevelTo,
			ByGM:         !f.byPlayer,
			Sources:      p.Sources,
			Groups:       p.Groups,
			Block:        p.Block,
			Error:        p.Error,
			Was:          rv.was,
		}
	}
	return views.HomeFileView{
		CampaignID:  f.campaignID,
		FileID:      f.id,
		CharacterID: rv.bb.ch.ID,
		Name:        rv.bb.ch.Name,
		TakenAt:     rv.bb.take.CreatedAt,
		LevelFrom:   p.LevelFrom,
		LevelTo:     p.LevelTo,
		ByGM:        !f.byPlayer,
		Sources:     p.Sources,
		Groups:      p.Groups,
		Removed:     p.Removed,
		Block:       p.Block,
		Error:       p.Error,
		Was:         rv.was,
	}
}

// formField reads a field from the request, the query of a page load or the
// body of a post.
func formField(c *gin.Context) func(string) string {
	return func(k string) string { return cmp.Or(c.PostForm(k), c.Query(k)) }
}

// homeFilePage shows the review; an htmx request (a pick changed) gets the form only.
func (a *App) homeFilePage(c *gin.Context) error {
	w, camp, f, ok := a.homeFileAccess(c)
	if !ok {
		return nil
	}
	rv, err := a.reviewHomeFile(c.Request.Context(), a.Store.Queries, f, formField(c))
	if err != nil {
		return err
	}
	v := rv.view(f)
	if isHTMX(c) {
		a.render(c, http.StatusOK, views.HomeFileReview(v))
		return nil
	}
	a.render(c, http.StatusOK, views.HomeFilePage(a.layout(w, v.PageTitle(), "", &camp, false), v))
	return nil
}

// homeFileAccept writes the merged character, and closes the take-out. The
// review is made again inside the transaction: the page may be stale.
func (a *App) homeFileAccept(c *gin.Context) error {
	_, camp, f, ok := a.homeFileAccess(c)
	if !ok {
		return nil
	}
	if f.kind == export.KindNewCharacter {
		return a.homeFileAdd(c, camp, f)
	}
	ctx := c.Request.Context()
	form := formField(c)

	// Files first, from the review the page saw: a row needs the file's size and
	// type, and storeFile would open a second connection while the transaction
	// holds the write lock. What the transaction does not use is removed again.
	var saved []string
	used, done := map[string]bool{}, false
	defer a.removeUnusedFiles(&saved, used, &done)
	var lookFiles []savedFile
	var blockFiles map[int64][]savedFile
	pre, err := a.reviewBringBack(ctx, a.Store.Queries, f, form)
	if err != nil {
		return err
	}
	preSaved := pre.was == c.PostForm(formWas) && pre.plan.Block == "" && pre.plan.Error == ""
	if preSaved {
		m := mapsOf(pre.file)
		if lookFiles, blockFiles, err = a.saveZipFiles(
			f,
			pre.x.saveContent(pre.file, m.Attachments),
			&saved,
		); err != nil {
			return err
		}
	}

	var rv homeReview
	var res extrasResult
	stale, allowedAny := false, false
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if rv, err = a.reviewBringBack(ctx, q, f, form); err != nil {
			return err
		}
		switch p := rv.plan; {
		case rv.was != c.PostForm(formWas):
			stale = true
			return nil
		case p.Block != "":
			return userError{errors.New(p.Block)}
		case p.Error != "":
			return userError{errors.New(p.Error)}
		case !preSaved:
			stale = true // the table changed between the two reviews
			return nil
		}
		for i, s := range rv.reviewed {
			if rv.plan.Sources[i].Choice != views.SourceAllow {
				continue
			}
			allowedAny = true
			if err := allowSource(ctx, q, camp.ID, s, rv.file.Entries); err != nil {
				return err
			}
		}
		if err := writeMerged(ctx, q, rv.bb.ch.ID, rv.plan.Merged); err != nil {
			return err
		}
		if err := replaceLevelUps(ctx, q, rv.bb.ch.ID, rv.plan.BuildFrom, rv.file); err != nil {
			return err
		}
		if res, err = applyExtras(ctx, q, rv, lookFiles, blockFiles, used); err != nil {
			return err
		}
		n, err := q.ReturnTakeHome(ctx, rv.bb.take.ID)
		if err != nil {
			return err
		}
		if n == 0 {
			return userf("This file was already brought back.")
		}
		return q.ReturnOlderTakeHomes(
			ctx,
			sqlc.ReturnOlderTakeHomesParams{CharacterID: rv.bb.ch.ID, ID: rv.bb.take.ID},
		)
	})
	if err != nil {
		return err
	}
	if stale {
		v := rv.view(f)
		v.Stale = rv.bb.ch.Name + " changed at the table meanwhile; review again."
		a.render(c, http.StatusOK, views.HomeFileReview(v))
		return nil
	}
	done = true

	a.homeFiles.drop(f.id)
	if _, err := UpgradeSheets(ctx, a.Store); err != nil {
		slog.Warn("sheet upgrade after bring-back", "err", err)
	}
	topics, err := a.characterTopics(ctx, rv.bb.ch.ID)
	if err != nil {
		return err
	}
	topics = append(
		topics,
		hub.Topic{Kind: hub.TopicCampaign, ID: camp.ID},
		hub.Topic{Kind: hub.TopicPending, ID: camp.ID},
	)
	if allowedAny || res.compendium {
		topics = append(topics, hub.Topic{Kind: hub.TopicCompendium})
	}
	topics = append(topics, noteTopics(camp.ID, res.notes)...)
	a.publish(c, camp.ID, topics...)
	a.render(c, http.StatusOK, views.HomeFileDone(rv.bb.ch.Name, rv.bb.ch.ID))
	return nil
}

// allowSource ticks a source for the campaign; one the table does not have yet
// is first created from the file, with its entries.
func allowSource(
	ctx context.Context,
	q *sqlc.Queries,
	campaignID int64,
	s export.TakeHomeSource,
	entries []export.TakeHomeEntry,
) error {
	if _, err := q.GetCompendiumSource(ctx, s.SourceKey); errors.Is(err, sql.ErrNoRows) {
		prio, err := appendedSourcePriority(ctx, q, s.SourceKey, s.Edition)
		if err != nil {
			return err
		}
		if err := q.UpsertCompendiumSource(ctx, sqlc.UpsertCompendiumSourceParams{
			SourceKey:   s.SourceKey,
			Path:        bringBackSourcePath,
			Edition:     s.Edition,
			License:     s.License,
			Attribution: s.Attribution,
			Published:   s.Published,
			Priority:    prio,
		}); err != nil {
			return err
		}
		for _, e := range entries {
			if e.SourceKey != s.SourceKey {
				continue
			}
			if err := q.UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{
				Kind:        e.Kind,
				Slug:        e.Slug,
				Name:        e.Name,
				SourceKey:   e.SourceKey,
				DataJson:    string(e.Data),
				BodyText:    e.BodyText,
				ContentHash: e.ContentHash,
			}); err != nil {
				return err
			}
		}
	} else if err != nil {
		return err
	}
	n, err := q.AllowCampaignSource(
		ctx,
		sqlc.AllowCampaignSourceParams{CampaignID: campaignID, SourceKey: s.SourceKey},
	)
	if err != nil {
		return err
	}
	if n == 0 {
		return userf("The source %q could not be added.", s.SourceKey)
	}
	return nil
}

// replaceLevelUps gives the character the revert history of the build it now has:
// the home's, the fallback's, or (the table's build) the one it already has.
func replaceLevelUps(
	ctx context.Context,
	q *sqlc.Queries,
	characterID int64,
	buildFrom string,
	file *export.TakeHomeFile,
) error {
	var from []export.TakeHomeLevelUp
	switch buildFrom {
	case buildFromHome:
		from = file.LevelUps
	case buildFromFallback:
		from = file.Fallback.LevelUps
	default:
		return nil
	}
	if err := q.DeleteLevelUps(ctx, characterID); err != nil {
		return err
	}
	for _, l := range from {
		if _, err := q.CreateLevelUpAt(
			ctx,
			sqlc.CreateLevelUpAtParams{
				CharacterID: characterID,
				MaxHp:       l.MaxHP,
				SheetJson:   string(l.Sheet),
				CreatedAt:   l.CreatedAt,
			},
		); err != nil {
			return err
		}
	}
	return nil
}

// saveZipFiles saves the attached files th names, read from the waiting zip f:
// opened again, since the review that planned them has closed its reader.
func (a *App) saveZipFiles(
	f homeFile,
	th export.TakeHome,
	saved *[]string,
) ([]savedFile, map[int64][]savedFile, error) {
	zf, err := os.Open(f.path) //nolint:gosec // a random name inside our own folder
	if err != nil {
		return nil, nil, userf("The file is gone. Ask for it again.")
	}
	defer func() { _ = zf.Close() }()
	st, err := zf.Stat()
	if err != nil {
		return nil, nil, err
	}
	file, err := export.ReadTakeHome(zf, st.Size(), f.kind)
	if err != nil {
		return nil, nil, userError{err}
	}
	return a.saveHomeFiles(th, func(zipName string) ([]byte, error) {
		body, err := file.OpenFile(zipName)
		if err != nil {
			return nil, userError{err}
		}
		return body, nil
	}, saved)
}

// removeUnusedFiles deletes the files saved before a transaction that its rows
// do not own: all of them unless it committed (done), then the unused ones.
func (a *App) removeUnusedFiles(saved *[]string, used map[string]bool, done *bool) {
	for _, p := range *saved {
		if !*done || !used[p] {
			_ = os.Remove(filepath.Join(a.attachmentsDir(), p))
		}
	}
}

// noteTopics are the topics of notes that changed in a campaign, and its list.
func noteTopics(campaignID int64, noteIDs []int64) []hub.Topic {
	if len(noteIDs) == 0 {
		return nil
	}
	topics := []hub.Topic{{Kind: hub.TopicNotes, ID: campaignID}}
	for _, id := range noteIDs {
		topics = append(topics, hub.Topic{Kind: hub.TopicNote, ID: id})
	}
	return topics
}

// homeFileReject drops the file; a bring-back leaves the take-out open, a
// visitor's new character is denied.
func (a *App) homeFileReject(c *gin.Context) error {
	_, camp, f, ok := a.homeFileAccess(c)
	if !ok {
		return nil
	}
	a.homeFiles.drop(f.id)
	newChar := f.kind == export.KindNewCharacter
	if newChar {
		a.Hub.DenyFile(camp.ID, f.id)
	}
	a.publish(c, camp.ID, hub.Topic{Kind: hub.TopicPending, ID: camp.ID})
	if c.PostForm(formPage) != "" {
		if newChar {
			redirect(c, route.Path(route.Campaign, camp.ID))
		} else {
			redirect(c, route.Path(route.Character, f.characterID))
		}
		return nil
	}
	note := "Rejected. The take-out stays open."
	if newChar {
		note = "Rejected."
	}
	a.render(c, http.StatusOK, views.HomeFileRejected(a.pendingViews(), note))
	return nil
}
