package web

// A character made at home comes to the table as a new-character file, from the
// GM ("Add from file…" on the campaign) or from a visitor on the join page,
// whose file waits as a join request the GM answers by reviewing it. Adding it
// creates the character, its looks, its private notes and homebrew in one
// transaction; for a visitor's file the request is approved last, with the new
// character's id, so the visitor's device logs in as it.

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// newCharacter is a checked new-character file.
type newCharacter struct {
	file *export.TakeHomeFile
	cf   export.CharacterFile
	// tableSources are the books of this table the sheet pins that the file does not list.
	tableSources []export.TakeHomeSource
}

// loadNewCharacter reads the zip at path and checks it against the table: its
// kind, server and content. The review repeats it on every request.
func (a *App) loadNewCharacter(
	ctx context.Context,
	q *sqlc.Queries,
	path string,
) (newCharacter, error) {
	var nc newCharacter
	f, err := os.Open(path) //nolint:gosec // a random name inside our own folder
	if err != nil {
		return nc, userf("The file is gone. Ask for it again.")
	}
	defer func() { _ = f.Close() }()
	st, err := f.Stat()
	if err != nil {
		return nc, err
	}
	file, err := export.ReadTakeHome(f, st.Size(), export.KindNewCharacter)
	if err != nil {
		return nc, userf("That is not a new character file: %w", err)
	}
	m := file.Manifest
	serverID, err := q.GetServerID(ctx)
	if err != nil {
		return nc, err
	}
	if m.ServerID != serverID {
		return nc, userf("This file belongs to another GM's table.")
	}
	if m.Campaign.System != gameSystem || file.Character.System != gameSystem {
		return nc, userf("This table plays %s only.", gameSystem)
	}
	switch nc.cf, err = checkFileContent(file); {
	case errors.Is(err, errNewerSheet):
		return nc, userf(
			"This file comes from a newer player app (%s); this server is %s and needs updating.",
			m.AppVersion,
			a.Version,
		)
	case err != nil:
		return nc, err
	}
	if err := checkFileEntries(file); err != nil {
		return nc, err
	}
	if _, err := checkMadeAtHome(ctx, q, m.MadeAtHome); err != nil {
		return nc, err
	}
	// The file lists only books added at home: the table's own ones it pins
	// come from here, so the review can offer those the campaign lacks.
	// shortcut: only the current sheet is read, not level-up history, and bring-backs
	// skip this (a book the GM switched off since the take-out slips through);
	// reuse this loop if that bites.
	keys, err := sheetSourceKeys(string(nc.cf.Sheet))
	if err != nil {
		return nc, err
	}
	for _, k := range keys {
		if slices.ContainsFunc(
			file.Sources,
			func(s export.TakeHomeSource) bool { return s.SourceKey == k },
		) {
			continue
		}
		src, err := q.GetCompendiumSource(ctx, k)
		if errors.Is(err, sql.ErrNoRows) {
			return nc, userf(
				"%s uses the source %s, which this table does not have.",
				nc.cf.Name,
				k,
			)
		} else if err != nil {
			return nc, err
		}
		nc.tableSources = append(
			nc.tableSources,
			export.TakeHomeSource{SourceKey: k, Edition: src.Edition},
		)
	}
	nc.file = file
	return nc, nil
}

// checkMadeAtHome refuses a file whose made-at-home token a live character
// already holds. takeover says a deleted one holds it: its row goes when the
// new character claims the token.
func checkMadeAtHome(
	ctx context.Context,
	q *sqlc.Queries,
	token string,
) (takeover bool, err error) {
	if token == "" {
		return false, nil
	}
	holder, err := q.GetNewCharacterByToken(ctx, token)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return false, nil
	case err != nil:
		return false, err
	case !holder.DeletedAt.Valid:
		return false, userf("%s was already added from this file.", holder.Name)
	}
	return true, nil
}

// claimMadeAtHome stores the file's token for the character just added.
func claimMadeAtHome(ctx context.Context, q *sqlc.Queries, token string, characterID int64) error {
	takeover, err := checkMadeAtHome(ctx, q, token)
	if err != nil || token == "" {
		return err
	}
	if takeover {
		if err := q.DeleteNewCharacterToken(ctx, token); err != nil {
			return err
		}
	}
	return q.CreateNewCharacterToken(
		ctx,
		sqlc.CreateNewCharacterTokenParams{CharacterID: characterID, Token: token},
	)
}

// newCharacterDialog is the dialog body that asks the GM for the file.
func (a *App) newCharacterDialog(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), id)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.NewCharacterBody(camp.ID, camp.Name))
	return nil
}

// storeHomeFile saves an uploaded new-character file, checks it and puts it
// waiting for the GM's review.
func (a *App) storeHomeFile(c *gin.Context, camp sqlc.Campaign, byPlayer bool) (string, error) {
	path, size, err := a.receiveHomeFile(c, "a new character file")
	if err != nil {
		return "", err
	}
	kept := false
	defer func() {
		if !kept {
			_ = os.Remove(path)
		}
	}()
	nc, err := a.loadNewCharacter(c.Request.Context(), a.Store.Queries, path)
	if err != nil {
		return "", err
	}
	id, err := a.homeFiles.put(homeFile{
		kind:          export.KindNewCharacter,
		campaignID:    camp.ID,
		campaignName:  camp.Name,
		characterName: nc.cf.Name,
		byPlayer:      byPlayer,
		savedAt:       nc.file.Manifest.SavedAt,
		path:          path,
		size:          size,
	})
	kept = err == nil
	return id, err
}

// newCharacterUpload takes the GM's new-character file straight to its review;
// a character file (.json, a download or an export) is added at once.
func (a *App) newCharacterUpload(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), id)
	if err != nil {
		return err
	}
	// The cap comes before the first read of the form.
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxTakeHomeBytes)
	fh, err := c.FormFile(formFile)
	if err != nil {
		return userf("Choose a character file (.zip or .json).")
	}
	if strings.EqualFold(filepath.Ext(fh.Filename), ".json") {
		if fh.Size > maxCharacterFileBytes {
			return userf("This character file is too large.")
		}
		return a.importCharacter(c, camp, fh)
	}
	fileID, err := a.storeHomeFile(c, camp, false)
	if err != nil {
		return err
	}
	a.publish(c, camp.ID, hub.Topic{Kind: hub.TopicPending, ID: camp.ID})
	redirect(c, route.Path(route.CampaignHomeFile, camp.ID, fileID))
	return nil
}

// joinNewCharacter takes a visitor's new-character file. It always becomes a
// request the GM answers, whether or not the campaign is open to join.
func (a *App) joinNewCharacter(c *gin.Context) error {
	if !visitorOnly(c) {
		return nil
	}
	camp, ok := a.joinCampaign(c)
	if !ok {
		return nil
	}
	// Binding the name reads the whole multipart body: cap it first.
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxTakeHomeBytes)
	var form playerForm
	if err := c.ShouldBind(&form); err != nil {
		if tooBig := new(http.MaxBytesError); errors.As(err, &tooBig) {
			a.renderJoinCharacters(
				c,
				http.StatusRequestEntityTooLarge,
				camp,
				"That file is too large for a character file.",
			)
			return nil
		}
		a.renderJoinCharacters(
			c,
			http.StatusUnprocessableEntity,
			camp,
			"Tell the GM your name (up to 12 characters).",
		)
		return nil
	}
	fileID, err := a.storeHomeFile(c, camp, true)
	if ue, ok := errors.AsType[userError](err); ok {
		a.renderJoinCharacters(c, http.StatusUnprocessableEntity, camp, ue.Error())
		return nil
	} else if err != nil {
		return err
	}
	f, _ := a.homeFiles.get(fileID, camp.ID)
	ua := useragent.Parse(c.Request.UserAgent())
	addr, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		addr = c.Request.RemoteAddr
	}
	p, ok := a.Hub.AddPending(hub.Pending{
		Addr:         addr,
		CampaignID:   camp.ID,
		CampaignName: camp.Name,
		Kind:         hub.PendingNewCharacter,
		Name:         f.characterName,
		FileID:       fileID,
		Player:       strings.TrimSpace(form.Player),
		Device:       userAgentSummary(ua),
		Phone:        ua.Mobile,
		Host:         lanHostname(c.Request.Context(), c.Request.RemoteAddr),
	})
	if !ok {
		a.homeFiles.drop(fileID)
		a.renderJoinCharacters(
			c,
			http.StatusTooManyRequests,
			camp,
			"Too many people are waiting to join. Ask the GM to answer the requests first.",
		)
		return nil
	}
	epoch, seq := a.Hub.Position()
	l := views.Layout{
		Title: "Waiting",
		WS:    &views.WS{URL: route.WS + "?pending=" + p.Secret, Epoch: epoch, Seq: seq},
	}
	a.render(
		c,
		http.StatusOK,
		views.JoinWaiting(l, views.Waiting{Secret: p.Secret, Name: p.Name, State: stateWaiting}),
	)
	return nil
}

// reviewNewCharacter reads, checks and plans a new character's file.
func (a *App) reviewNewCharacter(
	ctx context.Context,
	q *sqlc.Queries,
	f homeFile,
	form func(string) string,
) (homeReview, error) {
	var rv homeReview
	nc, err := a.loadNewCharacter(ctx, q, f.path)
	if err != nil {
		return rv, err
	}
	rv.file, rv.cf = nc.file, nc.cf
	allowed, err := q.ListCampaignSources(ctx, f.campaignID)
	if err != nil {
		return rv, err
	}
	var sources []reviewSource
	for _, s := range slices.Concat(rv.file.Sources, nc.tableSources) {
		if !slices.Contains(allowed, s.SourceKey) {
			rv.reviewed = append(rv.reviewed, s)
			sources = append(sources, reviewSource{Key: s.SourceKey, Label: sourceLabel(s)})
		}
	}
	if rv.x, err = planExtras(
		ctx,
		q,
		f.campaignID,
		0,
		rv.file,
		export.TakeHomeBase{},
		form,
	); err != nil {
		return rv, err
	}
	if err := applyRenames(rv.x.hb.Renames, &rv.cf, rv.file); err != nil {
		return rv, err
	}
	if rv.plan.Sources, rv.plan.Block, err = planNewSources(
		string(rv.cf.Sheet),
		rv.cf.Name,
		sources,
		form,
	); err != nil {
		return rv, err
	}
	if rv.plan.LevelTo, err = sheetLevel(string(rv.cf.Sheet)); err != nil {
		return rv, err
	}
	rv.plan.Groups = rv.x.groups()
	// There is no character at the table to stamp: what the GM was shown is the stamp.
	sum := sha256.Sum256(fmt.Appendf(nil, "%v", rv.plan.Groups))
	rv.was = hex.EncodeToString(sum[:])
	return rv, nil
}

// addContent is the file as it is added: its notes reduced to the private
// blocks, empty notes left out.
func (rv homeReview) addContent() export.TakeHome {
	th := rv.file.TakeHome
	th.Notes = rv.x.notes.NewNotes
	return th
}

// rebuildLinks writes the links of every note block of a new character.
func rebuildLinks(ctx context.Context, q *sqlc.Queries, ch sqlc.Character) error {
	blocks, err := q.ListOwnNoteBlocks(ctx, nullID(ch.ID))
	if err != nil {
		return err
	}
	looks, err := q.ListCharacterImages(ctx, ch.ID)
	if err != nil {
		return err
	}
	lookIDs := make([]int64, len(looks))
	for i, l := range looks {
		lookIDs[i] = l.ID
	}
	for _, b := range blocks {
		if err := writeBlockLinks(
			ctx,
			q,
			ch.CampaignID,
			b.NoteID,
			b.ID,
			json.RawMessage(b.DocJson),
			lookIDs,
		); err != nil {
			return err
		}
	}
	return nil
}

// homeFileAdd adds a new character's file to the campaign. The review is made
// again inside the transaction: the page may be stale.
func (a *App) homeFileAdd(c *gin.Context, camp sqlc.Campaign, f homeFile) error {
	// The visitor is approved after the commit, so the request cannot stop a second Add.
	if !a.homeFiles.claim(f.id) {
		return userf("This file is being added already.")
	}
	defer a.homeFiles.release(f.id)
	ctx := c.Request.Context()
	form := formField(c)

	var saved []string
	used, done := map[string]bool{}, false
	defer a.removeUnusedFiles(&saved, used, &done)
	var lookFiles []savedFile
	var blockFiles map[int64][]savedFile
	pre, err := a.reviewNewCharacter(ctx, a.Store.Queries, f, form)
	if err != nil {
		return err
	}
	preSaved := pre.was == c.PostForm(formWas) && pre.plan.Block == "" && pre.plan.Error == ""
	if preSaved {
		if lookFiles, blockFiles, err = a.saveZipFiles(f, pre.addContent(), &saved); err != nil {
			return err
		}
	}

	var rv homeReview
	var ch sqlc.Character
	stale, allowedAny := false, false
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if rv, err = a.reviewNewCharacter(ctx, q, f, form); err != nil {
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
		maps := newHomeIDMaps()
		if ch, err = createHomeCharacter(
			ctx,
			q,
			camp.ID,
			rv.cf,
			rv.addContent(),
			lookFiles,
			blockFiles,
			&maps,
			tableMentions,
		); err != nil {
			return err
		}
		if err := applyHomebrew(ctx, q, ch, rv.x.hb); err != nil {
			return err
		}
		if err := rebuildLinks(ctx, q, ch); err != nil {
			return err
		}
		if err := claimMadeAtHome(ctx, q, rv.file.Manifest.MadeAtHome, ch.ID); err != nil {
			return err
		}
		if f.byPlayer && !a.Hub.FileRequestWaiting(camp.ID, f.id) {
			return userf("The player left; ask them to send it again.")
		}
		return nil
	})
	if err != nil {
		return err
	}
	if stale {
		v := rv.view(f)
		v.Stale = "The table changed meanwhile; review again."
		a.render(c, http.StatusOK, views.HomeFileReview(v))
		return nil
	}
	done = true
	for _, p := range saved {
		used[p] = true // rows own every file the file's looks and private notes brought
	}
	// Approved only once committed, or the visitor could redeem it before the
	// character exists. One who left in between can claim it from the join page.
	if f.byPlayer && !a.Hub.ApproveNewCharacter(camp.ID, f.id, ch.ID) {
		slog.Info("new character added after its visitor left", "character", ch.ID)
	}

	a.homeFiles.drop(f.id)
	if _, err := UpgradeSheets(ctx, a.Store); err != nil {
		slog.Warn("sheet upgrade after adding a new character", "err", err)
	}
	topics := []hub.Topic{
		{Kind: hub.TopicCampaign, ID: camp.ID},
		{Kind: hub.TopicPending, ID: camp.ID},
		{Kind: hub.TopicNotes, ID: camp.ID},
	}
	if allowedAny || rv.x.hb.touched() {
		topics = append(topics, hub.Topic{Kind: hub.TopicCompendium})
	}
	a.publish(c, camp.ID, topics...)
	a.render(c, http.StatusOK, views.HomeFileAdded(ch.Name, ch.ID))
	return nil
}
