package perm

import (
	"database/sql"
	"path/filepath"
	"slices"
	"testing"

	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
)

// TestSQLAgreesWithCanReadBlock runs blockCases through the readable-block
// predicate copied into the listing, search and backlink queries, so the SQL
// and Go read rules cannot drift apart.
func TestSQLAgreesWithCanReadBlock(t *testing.T) {
	ctx := t.Context()
	conn, err := db.Open(ctx, filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	q := db.NewStore(conn)

	camp, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: "C", System: "dnd5e", Edition: "2024"})
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Alice", "Bob"} { // ids 1 and 2, matching alice and bob
		if _, err := q.CreateCharacter(ctx, sqlc.CreateCharacterParams{CampaignID: camp.ID, Name: name, System: "dnd5e"}); err != nil {
			t.Fatal(err)
		}
	}

	const targetEntity = 99
	for _, tt := range blockCases {
		t.Run(tt.name, func(t *testing.T) {
			note, err := q.CreateNote(ctx, sqlc.CreateNoteParams{
				CampaignID: camp.ID, Kind: "generic", Title: "needle " + tt.name,
				OwnerCharacterID: nullID(tt.block.NoteOwnerID),
			})
			if err != nil {
				t.Fatal(err)
			}
			block, err := q.CreateBlock(ctx, sqlc.CreateBlockParams{
				NoteID: note.ID, Visibility: string(tt.block.Visibility), PlayersCanEdit: tt.block.PlayersCanEdit,
				AuthorCharacterID: nullID(tt.block.AuthorID),
			})
			if len(tt.readers) == 0 && tt.block.Visibility != VisListed {
				// Blocks nobody may read must be refused by the schema itself.
				if err == nil {
					t.Fatalf("schema accepted an unreadable %q block", tt.block.Visibility)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			for _, id := range tt.block.Audience {
				if err := q.AddBlockAudience(ctx, sqlc.AddBlockAudienceParams{BlockID: block.ID, CharacterID: id}); err != nil {
					t.Fatal(err)
				}
			}
			if err := q.AddBlockLink(ctx, sqlc.AddBlockLinkParams{SourceNoteID: note.ID, SourceBlockID: block.ID, TargetType: "entity", TargetID: targetEntity}); err != nil {
				t.Fatal(err)
			}

			for _, who := range viewerOf {
				v := viewers[who]
				want := CanReadBlock(v, tt.block)
				if want != slices.Contains(tt.readers, who) {
					t.Fatalf("%s: table and CanReadBlock disagree", who)
				}
				isGM, char := v.IsGM, v.CharacterID

				blocks, err := q.ListReadableBlocks(ctx, sqlc.ListReadableBlocksParams{NoteID: note.ID, IsGm: isGM, CharacterID: char})
				if err != nil {
					t.Fatal(err)
				}
				if got := len(blocks) == 1; got != want {
					t.Errorf("%s: ListReadableBlocks = %v, want %v", who, got, want)
				}

				hits, err := q.SearchReadableBlocks(ctx, sqlc.SearchReadableBlocksParams{
					CampaignID: camp.ID, Pattern: "%needle " + tt.name + "%", IsGm: isGM, CharacterID: char, MaxResults: 10,
				})
				if err != nil {
					t.Fatal(err)
				}
				if got := len(hits) == 1; got != want {
					t.Errorf("%s: SearchReadableBlocks = %v, want %v", who, got, want)
				}

				back, err := q.ListBacklinks(ctx, sqlc.ListBacklinksParams{
					TargetType: "entity", TargetID: targetEntity, CampaignID: camp.ID, IsGm: isGM, CharacterID: char,
				})
				if err != nil {
					t.Fatal(err)
				}
				if got := slices.ContainsFunc(back, func(r sqlc.ListBacklinksRow) bool { return r.ID == note.ID }); got != want {
					t.Errorf("%s: ListBacklinks = %v, want %v", who, got, want)
				}

				notes, err := q.ListReadableNotes(ctx, sqlc.ListReadableNotesParams{CampaignID: camp.ID, IsGm: isGM, CharacterID: char, MaxResults: 1000})
				if err != nil {
					t.Fatal(err)
				}
				wantNote := want || CanManageNote(v, tt.block.NoteOwnerID)
				if got := slices.ContainsFunc(notes, func(n sqlc.Note) bool { return n.ID == note.ID }); got != wantNote {
					t.Errorf("%s: ListReadableNotes = %v, want %v", who, got, wantNote)
				}
			}
		})
	}
}

func nullID(id int64) sql.NullInt64 {
	return sql.NullInt64{Int64: id, Valid: id != 0}
}

// TestSQLAgreesWithCanSeeEntry checks the entry visibility predicate in the
// compendium search against CanSeeEntry, for each campaign link state.
func TestSQLAgreesWithCanSeeEntry(t *testing.T) {
	ctx := t.Context()
	conn, err := db.Open(ctx, filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	q := db.NewStore(conn)
	var camps []int64
	for _, name := range []string{"A", "B"} {
		c, err := q.CreateCampaign(ctx, sqlc.CreateCampaignParams{Name: name, System: "dnd5e", Edition: "2024"})
		if err != nil {
			t.Fatal(err)
		}
		camps = append(camps, c.ID)
	}
	for i, tt := range []struct {
		scoped bool
		linked []int64
	}{{false, nil}, {true, nil}, {true, camps[:1]}, {true, camps}} {
		e, err := q.CreateHomebrew(ctx, sqlc.CreateHomebrewParams{
			Kind: "spell", Slug: "needle-" + string(rune('a'+i)), Name: "Needle", DataJson: "{}", BodyText: "", ContentHash: "x", Scoped: tt.scoped,
		})
		if err != nil {
			t.Fatal(err)
		}
		for _, c := range tt.linked {
			if err := q.LinkEntryCampaign(ctx, sqlc.LinkEntryCampaignParams{EntryID: e.ID, CampaignID: c}); err != nil {
				t.Fatal(err)
			}
		}
		for _, campaign := range camps {
			rows, err := q.SearchCompendium(ctx, sqlc.SearchCompendiumParams{Query: `"needle"`, CampaignID: campaign, MaxResults: -1})
			if err != nil {
				t.Fatal(err)
			}
			inSQL := slices.ContainsFunc(rows, func(r sqlc.SearchCompendiumRow) bool { return r.ID == e.ID })
			if want := CanSeeEntry(campaign, tt.scoped, tt.linked); inSQL != want {
				t.Errorf("entry %d (scoped %v, linked %v) in campaign %d: SQL %v, perm %v", i, tt.scoped, tt.linked, campaign, inSQL, want)
			}
		}
	}
}
