package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/web/views"
)

func TestAttachmentRefs(t *testing.T) {
	tests := []struct {
		name string
		text string
		want []int64
	}{
		{"attached picture", `{"type":"mention","attrs":{"id":"attachment:7","label":"map"}}`, []int64{7}},
		{"older picture", `{"type":"image","attrs":{"src":"/attachments/8","alt":"map"}}`, []int64{8}},
		{"file link", `{"type":"link","attrs":{"href":"/attachments/9"}}`, []int64{9}},
		{"several", `"attachment:1" "/attachments/2"`, []int64{1, 2}},
		{"other links", `{"id":"image:7"} {"id":"note:8"} "/images/9/file"`, nil},
		{"longer number is another file", `"attachment:70"`, []int64{70}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := attachmentRefs([]string{tt.text})
			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			for _, id := range tt.want {
				if !got[id] {
					t.Errorf("missing %d in %v", id, got)
				}
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

func TestFilesCleanup(t *testing.T) {
	e := newTestEnv(t)
	block := e.playersBlock(t)
	kept, stray := e.attachFile(t, block, "kept.png"), e.attachFile(t, block, "stray.png")
	// Kept is in an edit not yet in the block's snapshot.
	step := fmt.Sprintf(`{"stepType":"replace","from":1,"to":1,"slice":{"content":[{"type":"mention","attrs":{"id":"attachment:%d"}}]}}`, kept.ID)
	if err := e.q().InsertBlockStep(t.Context(), sqlc.InsertBlockStepParams{BlockID: block, Version: 1, ClientID: "c", StepJson: step}); err != nil {
		t.Fatal(err)
	}
	gm, player := e.token(t, auth.RoleAdmin), e.token(t, auth.RolePlayer)

	if w := e.do(t, http.MethodGet, views.PathFiles, player, nil, false); w.Code == http.StatusOK {
		t.Fatal("a player may not open the Files page")
	}
	page := e.do(t, http.MethodGet, views.PathFiles, gm, nil, false).Body.String()
	if !strings.Contains(page, "stray.png") || strings.Contains(page, "kept.png") {
		t.Errorf("Files page should list only the unused file")
	}

	e.do(t, http.MethodPost, fmt.Sprintf("%s/%d/delete", views.PathFiles, kept.ID), gm, nil, true)
	if _, err := e.q().GetAttachment(t.Context(), kept.ID); err != nil {
		t.Error("a file still in use was deleted")
	}
	e.do(t, http.MethodPost, views.PathFiles+"/delete-unused", gm, nil, true)
	if _, err := e.q().GetAttachment(t.Context(), stray.ID); err == nil {
		t.Error("the unused file's row is still there")
	}
	if _, err := os.Stat(filepath.Join(e.app.attachmentsDir(), stray.Path)); !os.IsNotExist(err) {
		t.Error("the unused file is still on disk")
	}
}

func TestAttachmentToImage(t *testing.T) {
	e := newTestEnv(t)
	att := e.attachFile(t, e.playersBlock(t), "Harbour map.png")
	path := fmt.Sprintf("/attachments/%d/to-image", att.ID)
	if w := e.do(t, http.MethodPost, path, e.token(t, auth.RolePlayer), nil, true); w.Code == http.StatusOK {
		t.Fatal("a player may not add campaign images")
	}
	w := e.do(t, http.MethodPost, path, e.token(t, auth.RoleAdmin), nil, true)
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
	if w := e.do(t, http.MethodGet, fmt.Sprintf("/images/%d/file", img.ID), e.token(t, auth.RolePlayer), nil, false); w.Code != http.StatusOK {
		t.Errorf("player reading the converted image: %d", w.Code)
	}
}
