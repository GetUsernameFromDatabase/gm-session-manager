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

// linkTarget creates and deletes one kind of row that links point at by bare id.
type linkTarget struct {
	create func() (int64, error)
	remove func(int64) error
}

// linkTargets are those kinds, keyed by their link target_type.
func linkTargets(ctx context.Context, q *sqlc.Queries, campaignID int64) map[string]linkTarget {
	return map[string]linkTarget{
		"image": {
			create: func() (int64, error) {
				img, err := q.CreateImage(ctx, sqlc.CreateImageParams{CampaignID: campaignID, Name: "I", Url: sql.NullString{String: "u", Valid: true}})
				return img.ID, err
			},
			remove: func(id int64) error { return q.DeleteImage(ctx, id) },
		},
		"note": {
			create: func() (int64, error) {
				n, err := q.CreateNote(ctx, sqlc.CreateNoteParams{CampaignID: campaignID, Kind: enum.NoteGeneric, Title: "N"})
				return n.ID, err
			},
			remove: func(id int64) error { return q.DeleteNote(ctx, id) },
		},
		"entity": {
			create: func() (int64, error) {
				e, err := q.CreateEntity(ctx, sqlc.CreateEntityParams{CampaignID: campaignID, Kind: enum.EntityNPC, Name: "E", FieldsJson: "{}"})
				return e.ID, err
			},
			remove: func(id int64) error { return q.DeleteEntity(ctx, id) },
		},
	}
}

// TestIDsNeverReused: a deleted image, note or NPC never hands its id to a
// new row, since links store bare ids.
func TestIDsNeverReused(t *testing.T) {
	ctx := t.Context()
	q := sqlc.New(openTemp(t))
	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "C", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	for name, k := range linkTargets(ctx, q, camp.ID) {
		t.Run(name, func(t *testing.T) {
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

// TestStaleLinksNeverMatch: a link left by a delete before migration 10
// never points at an image, note or NPC created after it.
func TestStaleLinksNeverMatch(t *testing.T) {
	ctx := t.Context()
	path := filepath.Join(t.TempDir(), "test.db")
	all, err := loadMigrations(migrationsFS)
	if err != nil {
		t.Fatal(err)
	}
	old, err := openPool(ctx, path)
	if err != nil {
		t.Fatal(err)
	}
	for _, m := range all[:9] {
		if err := apply(ctx, old, m); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := old.ExecContext(ctx, `
		INSERT INTO campaigns (id, name, system, edition) VALUES (1, 'C', 'dnd5e', '2024');
		INSERT INTO entities (id, campaign_id, name) VALUES (1, 1, 'E');
		INSERT INTO notes (id, campaign_id, kind, title) VALUES (1, 1, 'generic', 'N');
		INSERT INTO note_blocks (id, note_id, visibility) VALUES (1, 1, 'players')`); err != nil {
		t.Fatal(err)
	}
	// Stale target ids per link table. The highest comes from note_links for
	// images and NPCs and from entity_field_links for notes, so both branches are checked.
	stale := map[string]map[string]int64{
		"image":  {"entity_field_links": 99, "note_links": 150},
		"note":   {"entity_field_links": 160, "note_links": 120},
		"entity": {"entity_field_links": 110, "note_links": 170},
	}
	insert := map[string]string{
		"entity_field_links": `INSERT INTO entity_field_links (entity_id, field_key, target_type, target_id) VALUES (1, 'f', ?, ?)`,
		"note_links":         `INSERT INTO note_links (source_note_id, source_block_id, target_type, target_id) VALUES (1, 1, ?, ?)`,
	}
	for kind, byTable := range stale {
		for table, id := range byTable {
			// One statement per Exec: a multi-statement Exec rebinds from the first argument.
			if _, err := old.ExecContext(ctx, insert[table], kind, id); err != nil {
				t.Fatal(err)
			}
		}
	}
	_ = old.Close()

	conn, err := Open(ctx, path)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	for kind, k := range linkTargets(ctx, sqlc.New(conn), 1) {
		highest := max(stale[kind]["entity_field_links"], stale[kind]["note_links"])
		id, err := k.create()
		if err != nil {
			t.Fatal(err)
		}
		if id <= highest {
			t.Errorf("%s id %d can match a stale link to %d", kind, id, highest)
		}
	}
}

// TestAlterAfterSchemaEdit: a migration that edits the stored schema (0003,
// 0010) leaves a later ALTER TABLE on the same connection working, as the
// next migration's would.
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
	for _, table := range []string{"images", "notes", "entities"} {
		if _, err := conn.ExecContext(ctx, "ALTER TABLE "+table+" ADD COLUMN probe INTEGER NOT NULL DEFAULT 0"); err != nil {
			t.Fatalf("alter %s: %v", table, err)
		}
		var n int
		if err := conn.QueryRowContext(ctx, "SELECT count(*) FROM "+table).Scan(&n); err != nil {
			t.Fatalf("%s unreadable after ALTER: %v", table, err)
		}
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
