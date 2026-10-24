package export

import (
	"archive/zip"
	"bytes"
	"database/sql"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
)

var noLinks = Links{Mention: func(string) string { return "" }, Image: func(string) string { return "" }}

func TestMarkdown(t *testing.T) {
	tests := []struct {
		name, doc, want string
		links           Links
	}{
		{"paragraphs", `{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Hello"}]},{"type":"paragraph","content":[{"type":"text","text":"World"}]}]}`, "Hello\n\nWorld\n", noLinks},
		{"heading and marks", `{"type":"doc","content":[{"type":"heading","attrs":{"level":2},"content":[{"type":"text","text":"Loot"}]},{"type":"paragraph","content":[{"type":"text","text":"a ","marks":[]},{"type":"text","text":"sword","marks":[{"type":"bold"}]}]}]}`, "## Loot\n\na **sword**\n", noLinks},
		{"lists", `{"type":"doc","content":[{"type":"bulletList","content":[{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"one"}]}]},{"type":"listItem","content":[{"type":"paragraph","content":[{"type":"text","text":"two"}]}]}]}]}`, "- one\n- two\n", noLinks},
		{"quote", `{"type":"doc","content":[{"type":"blockquote","content":[{"type":"paragraph","content":[{"type":"text","text":"Read aloud"}]}]}]}`, "> Read aloud\n", noLinks},
		{"escaping", `{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"*not bold* [x]"}]}]}`, `\*not bold\* \[x\]` + "\n", noLinks},
		{"mention resolved", `{"type":"doc","content":[{"type":"paragraph","content":[{"type":"mention","attrs":{"id":"note:3","label":"Inn"}}]}]}`, "[Inn](3-inn.md)\n",
			Links{Mention: func(id string) string { return map[string]string{"note:3": "3-inn.md"}[id] }, Image: noLinks.Image}},
		{"mention unresolved", `{"type":"doc","content":[{"type":"paragraph","content":[{"type":"mention","attrs":{"id":"entity:3","label":"Mira"}}]}]}`, "Mira\n", noLinks},
		{"image dropped when not exported", `{"type":"doc","content":[{"type":"image","attrs":{"src":"/attachments/9","alt":"map"}}]}`, "\n", noLinks},
		{"empty doc", `{"type":"doc","content":[{"type":"paragraph"}]}`, "\n", noLinks},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Markdown(tt.doc, tt.links)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// TestCampaignExcludesOwnerBlocks checks the export rule: players' private
// blocks, and notes left empty without them, never reach the archive.
func TestCampaignExcludesOwnerBlocks(t *testing.T) {
	ctx := t.Context()
	conn, err := db.Open(ctx, filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	q := db.NewStore(conn)

	camp, _ := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "C", System: "dnd5e", Edition: "2024"})
	alice, _ := q.CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: camp.ID, Name: "Alice", System: "dnd5e"})
	owner := sql.NullInt64{Int64: alice.ID, Valid: true}

	gmNote, _ := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: camp.ID, Kind: "game_session", Title: "Session one"})
	private, _ := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: camp.ID, Kind: "generic", Title: "Alice diary", OwnerCharacterID: owner})
	mixed, _ := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: camp.ID, Kind: "generic", Title: "Alice shared", OwnerCharacterID: owner})

	addBlock := func(noteID int64, vis, text string) {
		t.Helper()
		b, err := q.CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: noteID, Visibility: vis})
		if err != nil {
			t.Fatal(err)
		}
		doc := `{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"` + text + `"}]}]}`
		if err := q.SetBlockVersion(ctx, sqlc.SetBlockVersionParams{ID: b.ID, Version: 1}); err != nil {
			t.Fatal(err)
		}
		if _, err := q.SaveBlockSnapshot(ctx, sqlc.SaveBlockSnapshotParams{ID: b.ID, DocJson: doc, DocVersion: 1, Text: text}); err != nil {
			t.Fatal(err)
		}
	}
	addBlock(gmNote.ID, "gm", "GM secret plan")
	addBlock(private.ID, "owner", "Alice private thought")
	addBlock(mixed.ID, "owner", "Alice hidden part")
	addBlock(mixed.ID, "players", "Alice shared part")

	var buf bytes.Buffer
	if err := Campaign(ctx, &buf, q.Queries, camp.ID, t.TempDir()); err != nil {
		t.Fatal(err)
	}
	zr, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		t.Fatal(err)
	}
	var all strings.Builder
	var names []string
	for _, f := range zr.File {
		names = append(names, f.Name)
		rc, _ := f.Open()
		b, _ := io.ReadAll(rc)
		_ = rc.Close()
		all.Write(b)
	}
	text := all.String()
	for _, want := range []string{"GM secret plan", "Alice shared part", "Session one", `"name": "Alice"`} {
		if !strings.Contains(text, want) {
			t.Errorf("archive lacks %q (files %v)", want, names)
		}
	}
	for _, banned := range []string{"Alice private thought", "Alice hidden part", "Alice diary"} {
		if strings.Contains(text, banned) || strings.Contains(strings.Join(names, " "), "diary") {
			t.Errorf("archive leaks %q", banned)
		}
	}
}
