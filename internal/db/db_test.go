package db

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"path/filepath"
	"strings"
	"testing"
	"testing/fstest"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
)

func openTemp(t *testing.T) *sql.DB {
	t.Helper()
	conn, err := Open(t.Context(), filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	return conn
}

func TestPragmasOnEveryConnection(t *testing.T) {
	conn := openTemp(t)
	ctx := t.Context()

	// Hold several connections at once so the pool must create new ones.
	const n = 4
	conns := make([]*sql.Conn, n)
	for i := range conns {
		c, err := conn.Conn(ctx)
		if err != nil {
			t.Fatal(err)
		}
		conns[i] = c
	}
	for i, c := range conns {
		var fk int
		var mode string
		if err := c.QueryRowContext(ctx, "PRAGMA foreign_keys").Scan(&fk); err != nil {
			t.Fatal(err)
		}
		if err := c.QueryRowContext(ctx, "PRAGMA journal_mode").Scan(&mode); err != nil {
			t.Fatal(err)
		}
		if fk != 1 {
			t.Errorf("conn %d: foreign_keys = %d, want 1", i, fk)
		}
		if !strings.EqualFold(mode, "wal") {
			t.Errorf("conn %d: journal_mode = %q, want wal", i, mode)
		}
		_ = c.Close()
	}
}

// linkTarget creates and deletes one kind of row that links point at by id.
type linkTarget struct {
	create func() (int64, error)
	remove func(int64) error
}

// linkTargets are those kinds, keyed by their enum.LinkTarget.
func linkTargets(ctx context.Context, q *sqlc.Queries, campaignID int64) map[enum.LinkTarget]linkTarget {
	return map[enum.LinkTarget]linkTarget{
		enum.LinkImage: {
			create: func() (int64, error) {
				img, err := q.CreateImage(ctx, sqlc.CreateImageParams{CampaignID: campaignID, Name: "I", Url: sql.NullString{String: "u", Valid: true}})
				return img.ID, err
			},
			remove: func(n int64) error { return q.DeleteImage(ctx, n) },
		},
		enum.LinkNote: {
			create: func() (int64, error) {
				n, err := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: campaignID, Kind: enum.NoteGeneric, Title: "N"})
				return n.ID, err
			},
			remove: func(n int64) error { return q.DeleteNote(ctx, n) },
		},
		enum.LinkEntity: {
			create: func() (int64, error) {
				e, err := q.CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: campaignID, Kind: enum.EntityNPC, Name: "E", FieldsJson: "{}"})
				return e.ID, err
			},
			remove: func(n int64) error { return q.DeleteEntity(ctx, n) },
		},
	}
}

// targetColumns fills the link target column of kind, as the web handlers do.
func targetColumns(kind enum.LinkTarget, id int64) (note, entity, image sql.NullInt64) {
	v := sql.NullInt64{Int64: id, Valid: true}
	switch kind {
	case enum.LinkNote:
		note = v
	case enum.LinkEntity:
		entity = v
	case enum.LinkImage:
		image = v
	}
	return note, entity, image
}

// TestIDsNeverReused: a deleted image, note or NPC never hands its id to a
// new row, since [[kind:id|label]] tokens in text keep ids.
func TestIDsNeverReused(t *testing.T) {
	ctx := t.Context()
	q := sqlc.New(openTemp(t))
	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "C", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	for name, k := range linkTargets(ctx, q, camp.ID) {
		t.Run(string(name), func(t *testing.T) {
			gone, err := k.create()
			if err != nil {
				t.Fatal(err)
			}
			if err := k.remove(gone); err != nil {
				t.Fatal(err)
			}
			got, err := k.create()
			if err != nil {
				t.Fatal(err)
			}
			if got <= gone {
				t.Errorf("new id %d reuses deleted id %d", got, gone)
			}
		})
	}
}

// TestLinksFollowTargets: a link row lives only as long as its target, in
// both link tables, and a link to a missing target is skipped, not refused.
func TestLinksFollowTargets(t *testing.T) {
	ctx := t.Context()
	conn := openTemp(t)
	q := sqlc.New(conn)
	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "C", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	note, err := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: camp.ID, Kind: enum.NoteGeneric, Title: "Source"})
	if err != nil {
		t.Fatal(err)
	}
	block, err := q.CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: enum.VisPlayers})
	if err != nil {
		t.Fatal(err)
	}
	npc, err := q.CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: camp.ID, Kind: enum.EntityNPC, Name: "Source", FieldsJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}
	link := func(kind enum.LinkTarget, id int64) {
		t.Helper()
		n, e, i := targetColumns(kind, id)
		if err := q.AddBlockLink(ctx, sqlc.AddBlockLinkParams{SourceNoteID: note.ID, SourceBlockID: block.ID, TargetNoteID: n, TargetEntityID: e, TargetImageID: i}); err != nil {
			t.Fatalf("note link: %v", err)
		}
		if err := q.AddEntityFieldLink(ctx, sqlc.AddEntityFieldLinkParams{EntityID: npc.ID, FieldKey: "f", TargetNoteID: n, TargetEntityID: e, TargetImageID: i}); err != nil {
			t.Fatalf("field link: %v", err)
		}
	}
	count := func(table string) int {
		t.Helper()
		var n int
		if err := conn.QueryRowContext(ctx, "SELECT count(*) FROM "+table).Scan(&n); err != nil {
			t.Fatal(err)
		}
		return n
	}
	tables := []string{"note_links", "entity_field_links"}
	for kind, k := range linkTargets(ctx, q, camp.ID) {
		t.Run(string(kind), func(t *testing.T) {
			id, err := k.create()
			if err != nil {
				t.Fatal(err)
			}
			link(kind, id)
			link(kind, id) // a repeat is ignored
			for _, table := range tables {
				if got := count(table); got != 1 {
					t.Errorf("%s: %d rows after linking, want 1", table, got)
				}
			}
			if err := k.remove(id); err != nil {
				t.Fatal(err)
			}
			for _, table := range tables {
				if got := count(table); got != 0 {
					t.Errorf("%s: %d rows survive the target's delete", table, got)
				}
			}
			link(kind, id)
			for _, table := range tables {
				if got := count(table); got != 0 {
					t.Errorf("%s: %d rows stored for a missing target", table, got)
				}
			}
		})
	}
}

// TestAlterAfterSchemaEdit: a migration that edits the stored schema
// (writable_schema) leaves a later ALTER TABLE on the same connection
// working, as the next migration's would.
func TestAlterAfterSchemaEdit(t *testing.T) {
	ctx := t.Context()
	conn, err := openPool(ctx, filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	conn.SetMaxOpenConns(1) // the ALTER must run on the connection that migrated
	if err := Migrate(ctx, conn); err != nil {
		t.Fatal(err)
	}
	// Shifts the stored text before every later column, as a real edit would.
	edit := migration{version: 2, name: "0002_edit.sql", sql: `
		PRAGMA writable_schema = ON;
		UPDATE sqlite_schema
		SET sql = replace(sql, 'REFERENCES note_blocks (id) ON DELETE SET NULL', 'REFERENCES note_blocks (id)  ON  DELETE  SET  NULL')
		WHERE type = 'table' AND name = 'attachments';
		PRAGMA writable_schema = OFF;`}
	if err := apply(ctx, conn, edit); err != nil {
		t.Fatal(err)
	}
	var stored string
	if err := conn.QueryRowContext(ctx, "SELECT sql FROM sqlite_schema WHERE name = 'attachments'").Scan(&stored); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stored, "(id)  ON  DELETE") {
		t.Fatal("the edit matched nothing, so the ALTER below tests nothing")
	}
	if _, err := conn.ExecContext(ctx, "ALTER TABLE attachments ADD COLUMN probe INTEGER NOT NULL DEFAULT 0"); err != nil {
		t.Fatalf("alter: %v", err)
	}
	var n int
	if err := conn.QueryRowContext(ctx, "SELECT count(*) FROM attachments").Scan(&n); err != nil {
		t.Fatalf("unreadable after ALTER: %v", err)
	}
}

func TestMigrateIsIdempotent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "test.db")
	ctx := t.Context()
	all, err := loadMigrations(migrationsFS)
	if err != nil {
		t.Fatal(err)
	}
	for range 2 {
		conn, err := Open(ctx, path)
		if err != nil {
			t.Fatal(err)
		}
		var v int
		if err := conn.QueryRowContext(ctx, "PRAGMA user_version").Scan(&v); err != nil {
			t.Fatal(err)
		}
		if v != len(all) {
			t.Errorf("user_version = %d, want %d", v, len(all))
		}
		_ = conn.Close()
	}
}

func TestLoadMigrationsRejectsBadNames(t *testing.T) {
	tests := []struct {
		name  string
		files []string
	}{
		{"gap", []string{"migrations/0001_a.sql", "migrations/0003_c.sql"}},
		{"duplicate", []string{"migrations/0001_a.sql", "migrations/0001_b.sql"}},
		{"no number", []string{"migrations/init.sql"}},
		{"starts at two", []string{"migrations/0002_a.sql"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsys := fstest.MapFS{}
			for _, f := range tt.files {
				fsys[f] = &fstest.MapFile{Data: []byte("SELECT 1;")}
			}
			if _, err := loadMigrations(fsys); err == nil {
				t.Error("expected error")
			}
		})
	}
}

func TestSessionConstraints(t *testing.T) {
	conn := openTemp(t)
	ctx := t.Context()
	q := sqlc.New(conn)

	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "Test", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	if !camp.OpenJoin {
		t.Error("open_join should default to true")
	}
	char, err := q.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: camp.ID, Name: "Alice", System: "dnd5e", SheetJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}

	hash := func(s string) []byte { h := sha256.Sum256([]byte(s)); return h[:] }
	charID := sql.NullInt64{Int64: char.ID, Valid: true}

	tests := []struct {
		name    string
		params  sqlc.CreateSessionParams
		wantErr bool
	}{
		{"player with character", sqlc.CreateSessionParams{TokenHash: hash("a"), Role: "player", CharacterID: charID}, false},
		{"second device same character", sqlc.CreateSessionParams{TokenHash: hash("b"), Role: "player", CharacterID: charID}, false},
		{"admin without character", sqlc.CreateSessionParams{TokenHash: hash("c"), Role: "admin"}, false},
		{"player without character", sqlc.CreateSessionParams{TokenHash: hash("d"), Role: "player"}, true},
		{"admin with character", sqlc.CreateSessionParams{TokenHash: hash("e"), Role: "admin", CharacterID: charID}, true},
		{"unknown role", sqlc.CreateSessionParams{TokenHash: hash("f"), Role: "gm"}, true},
		{"duplicate token hash", sqlc.CreateSessionParams{TokenHash: hash("a"), Role: "admin"}, true},
		{"short token hash", sqlc.CreateSessionParams{TokenHash: []byte("short"), Role: "admin"}, true},
		{"missing character", sqlc.CreateSessionParams{TokenHash: hash("g"), Role: "player", CharacterID: sql.NullInt64{Int64: 9999, Valid: true}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := q.CreateSession(ctx, tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	t.Run("deleting character removes its sessions", func(t *testing.T) {
		if err := q.DeleteCharacter(ctx, char.ID); err != nil {
			t.Fatal(err)
		}
		var n int
		if err := conn.QueryRowContext(ctx, "SELECT count(*) FROM sessions WHERE role = 'player'").Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n != 0 {
			t.Errorf("%d player sessions survived character delete", n)
		}
	})
}

// The milestone 2 compendium search depends on FTS5 being compiled into the driver.
func TestFTS5Available(t *testing.T) {
	conn := openTemp(t)
	if _, err := conn.ExecContext(t.Context(), "CREATE VIRTUAL TABLE temp.fts_probe USING fts5(body)"); err != nil {
		t.Fatalf("fts5 not available: %v", err)
	}
}

// TestAttachmentOutlivesBlock: deleting the block a file was uploaded to
// keeps the file (as the GM's) and the image made from it.
func TestAttachmentOutlivesBlock(t *testing.T) {
	ctx := t.Context()
	q := sqlc.New(openTemp(t))
	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "C", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	note, err := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: camp.ID, Kind: "generic", Title: "N"})
	if err != nil {
		t.Fatal(err)
	}
	block, err := q.CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: note.ID, Visibility: "gm"})
	if err != nil {
		t.Fatal(err)
	}
	att, err := q.CreateAttachment(ctx, sqlc.CreateAttachmentParams{
		CampaignID: camp.ID, BlockID: sql.NullInt64{Int64: block.ID, Valid: true}, Path: "f", Name: "map.png", Mime: "image/png", Size: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	img, err := q.CreateImage(ctx, sqlc.CreateImageParams{CampaignID: camp.ID, Name: "Map", AttachmentID: sql.NullInt64{Int64: att.ID, Valid: true}})
	if err != nil {
		t.Fatal(err)
	}
	if err := q.DeleteBlock(ctx, block.ID); err != nil {
		t.Fatal(err)
	}
	got, err := q.GetAttachment(ctx, att.ID)
	if err != nil {
		t.Fatalf("attachment deleted with its block: %v", err)
	}
	if got.BlockID.Valid {
		t.Errorf("block_id = %d, want NULL", got.BlockID.Int64)
	}
	ids, err := q.ListImageAttachments(ctx)
	if err != nil || len(ids) != 1 || ids[0].AttachmentID.Int64 != att.ID {
		t.Errorf("image %d lost its attachment: %v %v", img.ID, ids, err)
	}

	// A private block's file goes with it, even through a cascade.
	ch, err := q.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: camp.ID, Name: "A", System: "dnd5e", SheetJson: "{}"})
	if err != nil {
		t.Fatal(err)
	}
	own, err := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: camp.ID, Kind: "generic", Title: "Mine", OwnerCharacterID: sql.NullInt64{Int64: ch.ID, Valid: true}})
	if err != nil {
		t.Fatal(err)
	}
	ownBlock, err := q.CreateBlock(ctx, sqlc.CreateBlockParams{NoteID: own.ID, Visibility: "owner"})
	if err != nil {
		t.Fatal(err)
	}
	private, err := q.CreateAttachment(ctx, sqlc.CreateAttachmentParams{
		CampaignID: camp.ID, BlockID: sql.NullInt64{Int64: ownBlock.ID, Valid: true}, Path: "p", Name: "diary.png", Mime: "image/png", Size: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := q.DeleteCharacter(ctx, ch.ID); err != nil {
		t.Fatal(err)
	}
	if _, err := q.GetAttachment(ctx, private.ID); !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("private attachment survived its block: %v", err)
	}
}
