package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

func TestAttachmentRefs(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []int64
	}{
		{"attached picture", `{"type":"mention","attrs":{"id":"attachment:7","label":"map"}}`, []int64{7}},
		{"older picture", `{"type":"image","attrs":{"src":"` + route.Path(route.Attachment, 8) + `","alt":"map"}}`, []int64{8}},
		{"file link", `{"type":"link","attrs":{"href":"` + route.Path(route.Attachment, 9) + `"}}`, []int64{9}},
		{"several", `"attachment:1" "` + route.Path(route.Attachment, 2) + `"`, []int64{1, 2}},
		{"other links", `{"id":"image:7"} {"id":"note:8"} "` + route.Path(route.ImageFile, 9) + `"`, nil},
		{"longer number is another file", `"attachment:70"`, []int64{70}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := attachmentRefs(tt.text); !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

// pngHead makes a stored file sniff as a PNG picture.
const pngHead = "\x89PNG\r\n\x1a\n\x00\x00\x00\x0dIHDR"

// attachFile stores a picture on a block, as an upload would.
func (e testEnv) attachFile(t *testing.T, blockID int64, name string) sqlc.Attachment {
	t.Helper()
	att, err := e.app.storeFile(t.Context(), e.campaignID, nullID(blockID), name, strings.NewReader(pngHead))
	if err != nil {
		t.Fatal(err)
	}
	return att
}

func (e testEnv) playersBlock(t *testing.T) int64 {
	t.Helper()
	note, err := e.q().CreateNote(t.Context(), sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "generic", Title: "Maps"})
	if err != nil {
		t.Fatal(err)
	}
	b, err := e.q().CreateBlock(t.Context(), sqlc.CreateBlockParams{NoteID: note.ID, Visibility: "players"})
	if err != nil {
		t.Fatal(err)
	}
	return b.ID
}

// TestDeleteRemovesFiles: deleting what holds a file removes the file at once,
// while a file just written and not yet in a row survives the sweep.
func TestDeleteRemovesFiles(t *testing.T) {
	tests := []struct {
		name   string
		player bool
		route  func(e testEnv, note, block int64) string
	}{
		{"owner block", true, func(_ testEnv, _, block int64) string { return route.Path(route.BlockDelete, block) }},
		{"owner note", true, func(_ testEnv, note, _ int64) string { return route.Path(route.NoteDelete, note) }},
		{"character", false, func(e testEnv, _, _ int64) string { return route.Path(route.CharacterDelete, e.characterID) }},
		{"campaign", false, func(e testEnv, _, _ int64) string { return route.Path(route.CampaignDelete, e.campaignID) }},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newTestEnv(t)
			ctx := t.Context()
			note, err := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: enum.NoteGeneric, Title: "Mine", OwnerCharacterID: nullID(e.characterID)})
			if err != nil {
				t.Fatal(err)
			}
			b, err := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: enum.VisOwner})
			if err != nil {
				t.Fatal(err)
			}
			att := e.attachFile(t, b.ID, "mine.png")
			file := filepath.Join(e.app.attachmentsDir(), att.Path)
			old := time.Now().Add(-time.Hour)
			if err := os.Chtimes(file, old, old); err != nil {
				t.Fatal(err)
			}
			fresh := filepath.Join(e.app.attachmentsDir(), "fresh.png")
			if err := os.WriteFile(fresh, []byte(pngHead), 0o600); err != nil {
				t.Fatal(err)
			}
			role := enum.RoleAdmin
			if tt.player {
				role = enum.RolePlayer
			}
			if w := e.do(t, http.MethodPost, tt.route(e, note.ID, b.ID), e.token(t, role), nil, true); w.Code >= http.StatusMultipleChoices {
				t.Fatalf("delete: %d %s", w.Code, w.Body)
			}
			if _, err := os.Stat(file); !os.IsNotExist(err) {
				t.Errorf("file still there after the delete: %v", err)
			}
			if _, err := os.Stat(fresh); err != nil {
				t.Errorf("fresh file swept: %v", err)
			}
		})
	}
}

func TestFilesCleanup(t *testing.T) {
	e := newTestEnv(t)
	block := e.playersBlock(t)
	kept, stray := e.attachFile(t, block, "kept.png"), e.attachFile(t, block, "stray.png")
	// Kept is in an edit not yet in the block's snapshot.
	step := fmt.Sprintf(`{"stepType":"replace","from":1,"to":1,"slice":{"content":[{"type":"mention","attrs":{"id":"attachment:%d"}}]}}`, kept.ID)
	if err := e.q().InsertBlockStep(t.Context(), sqlc.InsertBlockStepParams{BlockID: block, Version: 1, ClientID: "c", StepJson: step}); err != nil {
		t.Fatal(err)
	}
	gm, player := e.token(t, enum.RoleAdmin), e.token(t, enum.RolePlayer)

	if w := e.do(t, http.MethodGet, route.Files, player, nil, false); w.Code == http.StatusOK {
		t.Fatal("a player may not open the Files page")
	}
	uses, err := e.app.attachmentUses(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	if got := uses[kept.ID]; len(got) != 1 || got[0].Label != "Maps" || len(uses[stray.ID]) != 0 {
		t.Errorf("uses of kept %v, of stray %v", got, uses[stray.ID])
	}
	page := e.do(t, http.MethodGet, route.Files, gm, nil, false).Body.String()
	if !strings.Contains(page, "stray.png") || !strings.Contains(page, "kept.png") || !strings.Contains(page, ">Maps<") {
		t.Errorf("Files page should list every file and the note using it")
	}

	e.do(t, http.MethodPost, route.Path(route.FilesAttachmentDelete, kept.ID), gm, nil, true)
	if _, err := e.q().GetAttachment(t.Context(), kept.ID); err != nil {
		t.Error("a file still in use was deleted")
	}
	e.do(t, http.MethodPost, route.FilesDeleteUnused, gm, nil, true)
	if _, err := e.q().GetAttachment(t.Context(), stray.ID); err == nil {
		t.Error("the unused file's row is still there")
	}
	if _, err := os.Stat(filepath.Join(e.app.attachmentsDir(), stray.Path)); !os.IsNotExist(err) {
		t.Error("the unused file is still on disk")
	}
}

// TestCampaignFiles: a campaign's Files page lists and cleans up only that
// campaign's files; /files lists every campaign's.
func TestCampaignFiles(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	other, err := e.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Other", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	mine := e.attachFile(t, 0, "mine.png")
	theirs, err := e.app.storeFile(ctx, other.ID, nullID(0), "theirs.png", strings.NewReader(pngHead))
	if err != nil {
		t.Fatal(err)
	}
	gm, player := e.token(t, enum.RoleAdmin), e.token(t, enum.RolePlayer)
	base := views.FilesBase(e.campaignID)

	page := e.do(t, http.MethodGet, base, gm, nil, false).Body.String()
	if !strings.Contains(page, "mine.png") || strings.Contains(page, "theirs.png") {
		t.Error("a campaign's Files page should list its own files only")
	}
	if page := e.do(t, http.MethodGet, route.Files, gm, nil, false).Body.String(); !strings.Contains(page, "mine.png") || !strings.Contains(page, "theirs.png") {
		t.Error("the all-campaigns files page should list every campaign's files")
	}
	if w := e.do(t, http.MethodGet, base, player, nil, false); w.Code == http.StatusOK {
		t.Error("a player may not open a campaign's Files page")
	}
	if w := e.do(t, http.MethodGet, views.FilesBase(other.ID+100), gm, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("unknown campaign: status %d, want 404", w.Code)
	}

	e.do(t, http.MethodPost, fmt.Sprintf("%s/%d/delete", base, theirs.ID), gm, nil, true)
	e.do(t, http.MethodPost, base+route.DeleteUnused, gm, nil, true)
	if _, err := e.q().GetAttachment(ctx, mine.ID); err == nil {
		t.Error("the campaign's unused file is still there")
	}
	if _, err := e.q().GetAttachment(ctx, theirs.ID); err != nil {
		t.Error("another campaign's file was deleted from this campaign's page")
	}
}

func TestAttachmentToImage(t *testing.T) {
	e := newTestEnv(t)
	att := e.attachFile(t, e.playersBlock(t), "Harbour map.png")
	path := route.Path(route.AttachmentToImage, att.ID)
	if w := e.do(t, http.MethodPost, path, e.token(t, enum.RolePlayer), nil, true); w.Code == http.StatusOK {
		t.Fatal("a player may not add campaign images")
	}
	w := e.do(t, http.MethodPost, path, e.token(t, enum.RoleAdmin), nil, true)
	if w.Code != http.StatusOK {
		t.Fatalf("status %d: %s", w.Code, w.Body)
	}
	var got struct {
		ID   int64
		Name string
	}
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	img, err := e.q().GetImage(t.Context(), got.ID)
	if err != nil || img.Name != "Harbour map" || got.Name != img.Name || img.CampaignID != e.campaignID {
		t.Fatalf("image %+v (%v), reply %+v", img, err, got)
	}
	// The image has its own copy: other notes may still show the attachment.
	if !img.AttachmentID.Valid || img.AttachmentID.Int64 == att.ID {
		t.Errorf("image file %v, want a copy of %d", img.AttachmentID, att.ID)
	}
	if _, err := e.q().GetAttachment(t.Context(), att.ID); err != nil {
		t.Error("the note's attachment is gone")
	}
	// Players of the block see the new image before the editor's next snapshot links it.
	if w := e.do(t, http.MethodGet, route.Path(route.ImageFile, img.ID), e.token(t, enum.RolePlayer), nil, false); w.Code != http.StatusOK {
		t.Errorf("player reading the converted image: %d", w.Code)
	}
}

// TestUploadAttachmentURL: the url an upload answers with is the route that
// serves the file (notes.js recognises attachments by ROUTES.attachment).
func TestUploadAttachmentURL(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	w := e.doFile(t, route.Path(route.BlockAttachments, e.playersBlock(t)), gm, nil, "file", "map.png", pngHead)
	if w.Code != http.StatusOK {
		t.Fatalf("upload: %d %s", w.Code, w.Body)
	}
	var got struct {
		ID  int64
		URL string
	}
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
		t.Fatal(err)
	}
	want := route.Path(route.Attachment, got.ID)
	if got.URL != want {
		t.Errorf("url %q, want %q", got.URL, want)
	}
	if w := e.do(t, http.MethodGet, want, gm, nil, false); w.Code != http.StatusOK {
		t.Errorf("GET %s: %d", want, w.Code)
	}
}

// TestAttachmentFileMissing: a row whose file is gone from disk is a 404,
// not a server error that files a report.
func TestAttachmentFileMissing(t *testing.T) {
	e := newTestEnv(t)
	gm := e.token(t, enum.RoleAdmin)
	att := e.attachFile(t, e.playersBlock(t), "gone.png")
	if err := os.Remove(filepath.Join(e.app.attachmentsDir(), att.Path)); err != nil {
		t.Fatal(err)
	}
	if w := e.do(t, http.MethodGet, route.Path(route.Attachment, att.ID), gm, nil, false); w.Code != http.StatusNotFound {
		t.Errorf("GET a missing file: %d, want 404", w.Code)
	}
}

// TestAttachmentAccess: a file is readable exactly when its block is, only
// in its campaign, and a file with no block is the GM's.
func TestAttachmentAccess(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	player := func(campaignID int64, name string) string {
		t.Helper()
		ch, err := e.q().CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: campaignID, Name: name, System: "dnd5e", SheetJson: "{}"})
		if err != nil {
			t.Fatal(err)
		}
		tok, _, err := e.mgr.IssuePlayer(ctx, ch.ID)
		if err != nil {
			t.Fatal(err)
		}
		return tok
	}
	other, err := e.q().CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Other", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	viewers := map[string]string{
		"gm": e.token(t, enum.RoleAdmin), "alice": e.token(t, enum.RolePlayer),
		"bob": player(e.campaignID, "Bob"), "carol": player(other.ID, "Carol"),
	}
	block := func(owner int64, vis enum.Visibility) int64 {
		t.Helper()
		note, err := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: enum.NoteGeneric, Title: string(vis), OwnerCharacterID: nullID(owner)})
		if err != nil {
			t.Fatal(err)
		}
		b, err := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: vis})
		if err != nil {
			t.Fatal(err)
		}
		return b.ID
	}
	listed := block(0, "listed")
	if err := e.q().AddBlockAudience(ctx, sqlc.AddBlockAudienceParams{BlockID: listed, CharacterID: e.characterID}); err != nil {
		t.Fatal(err)
	}
	gone := block(0, "players")
	orphan := e.attachFile(t, gone, "orphan.png")
	if err := e.q().DeleteBlock(ctx, gone); err != nil {
		t.Fatal(err)
	}
	// The orphan's missing block must still scan into the typed visibility.
	if _, err := e.q().ListAttachmentFiles(ctx); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		att     sqlc.Attachment
		readers []string
	}{
		{"gm block", e.attachFile(t, block(0, "gm"), "a.png"), []string{"gm"}},
		{"listed for alice", e.attachFile(t, listed, "b.png"), []string{"gm", "alice"}},
		{"players block", e.attachFile(t, block(0, "players"), "c.png"), []string{"gm", "alice", "bob"}},
		{"alice's own block", e.attachFile(t, block(e.characterID, "owner"), "d.png"), []string{"alice"}},
		{"block deleted", orphan, []string{"gm"}},
	}
	for _, tt := range tests {
		for who, tok := range viewers {
			want := http.StatusNotFound
			if slices.Contains(tt.readers, who) {
				want = http.StatusOK
			}
			if w := e.do(t, http.MethodGet, route.Path(route.Attachment, tt.att.ID), tok, nil, false); w.Code != want {
				t.Errorf("%s, %s: %d, want %d", tt.name, who, w.Code, want)
			}
		}
	}
}

// TestFileUsePrivacy: the Files page names the notes using a file, except a
// player's own block, which shows neither its note's title nor a link.
func TestFileUsePrivacy(t *testing.T) {
	e := newTestEnv(t)
	ctx := t.Context()
	tests := []struct {
		vis   enum.Visibility
		owner int64
		want  views.FileUse
	}{
		{"players", 0, views.FileUse{Kind: views.FileUseNote, Label: "Secret plans"}},
		{"gm", 0, views.FileUse{Kind: views.FileUseNote, Label: "Secret plans"}},
		{"owner", e.characterID, views.FileUse{Kind: views.FileUseNote, Private: true}},
	}
	for _, tt := range tests {
		t.Run(string(tt.vis), func(t *testing.T) {
			note, err := e.q().CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: e.campaignID, Kind: "generic", Title: "Secret plans", OwnerCharacterID: nullID(tt.owner)})
			if err != nil {
				t.Fatal(err)
			}
			b, err := e.q().CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: tt.vis})
			if err != nil {
				t.Fatal(err)
			}
			att := e.attachFile(t, b.ID, "plan.png")
			step := fmt.Sprintf(`{"stepType":"replace","from":1,"to":1,"slice":{"content":[{"type":"mention","attrs":{"id":"attachment:%d"}}]}}`, att.ID)
			if err := e.q().InsertBlockStep(ctx, sqlc.InsertBlockStepParams{BlockID: b.ID, Version: 1, ClientID: "c", StepJson: step}); err != nil {
				t.Fatal(err)
			}
			uses, err := e.app.attachmentUses(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if tt.want.Label != "" {
				tt.want.URL = route.Path(route.Note, note.ID)
			}
			if got := uses[att.ID]; len(got) != 1 || got[0] != tt.want {
				t.Errorf("uses %v, want %v", got, tt.want)
			}
		})
	}
}
