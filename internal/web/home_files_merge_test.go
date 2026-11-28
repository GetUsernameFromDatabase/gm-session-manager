package web

import (
	"encoding/json"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
)

func tableLook(id int64, name string, revealed, current bool) sqlc.ListCharacterImagesRow {
	return sqlc.ListCharacterImagesRow{ID: id, Name: name, IsRevealed: revealed, IsCurrent: current}
}

func fileLook(home int64, name string, revealed, current bool) export.TakeHomeLook {
	return export.TakeHomeLook{
		ImageID:    home,
		Name:       name,
		URL:        "https://example.com/" + name,
		IsRevealed: revealed,
		IsCurrent:  current,
	}
}

func TestPlanLooks(t *testing.T) {
	base := []export.TakeHomeBaseLook{
		{ImageID: 10, IsRevealed: true, IsCurrent: true},
		{ImageID: 20, IsRevealed: true},
	}
	maps := map[int64]int64{1: 10, 2: 20}
	tests := []struct {
		name  string
		table []sqlc.ListCharacterImagesRow
		file  []export.TakeHomeLook
		maps  map[int64]int64
		pick  string

		wantNew    []int
		wantUnlink []int64
		wantReveal map[int64]bool
		wantNow    int64
		wantBadge  string
		wantLine   string
	}{
		{
			name: "nothing changed",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", true, false),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, true),
				fileLook(2, "B", true, false),
			},
			maps:    maps,
			wantNow: 10,
		},
		{
			name: "new at home",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", true, false),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, true),
				fileLook(2, "B", true, false),
				fileLook(3, "Cloak", true, false),
			},
			maps:      maps,
			wantNew:   []int{2},
			wantNow:   10,
			wantBadge: titleHome,
			wantLine:  "New: Cloak",
		},
		{
			name: "unlinked at home",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", true, false),
			},
			file:       []export.TakeHomeLook{fileLook(1, "A", true, true)},
			maps:       maps,
			wantUnlink: []int64{20},
			wantNow:    10,
			wantBadge:  titleHome,
			wantLine:   "Removed: B",
		},
		{
			name: "never went home stays",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", true, false),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, true),
			},
			maps:    map[int64]int64{1: 10},
			wantNow: 10,
		},
		{
			name:  "unlinked at the table stays gone",
			table: []sqlc.ListCharacterImagesRow{tableLook(10, "A", true, true)},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, true),
				fileLook(2, "B", false, false),
			},
			maps:    maps,
			wantNow: 10,
		},
		{
			name: "hidden at home",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", true, false),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, true),
				fileLook(2, "B", false, false),
			},
			maps: maps,
			wantReveal: map[int64]bool{
				20: false,
			},
			wantNow:   10,
			wantBadge: titleHome,
			wantLine:  "B: hidden",
		},
		{
			name: "hidden at the table only",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", false, false),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, true),
				fileLook(2, "B", true, false),
			},
			maps:    maps,
			wantNow: 10,
		},
		{
			name: "revealed at home",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", false, false),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, true),
				fileLook(2, "B", true, false),
			},
			maps:    maps,
			wantNow: 10,
		},
		{
			name: "current moved at home",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", true, false),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, false),
				fileLook(2, "B", true, true),
			},
			maps:      maps,
			wantNow:   20,
			wantBadge: titleHome,
			wantLine:  "Current: B",
		},
		{
			name: "current moved at the table only",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, false),
				tableLook(20, "B", true, true),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, true),
				fileLook(2, "B", true, false),
			},
			maps:    maps,
			wantNow: 20,
		},
		{
			name: "current conflict starts on home",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, false),
				tableLook(20, "B", true, true),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, false),
				fileLook(2, "B", true, false),
				fileLook(3, "Cloak", true, true),
			},
			maps:      maps,
			wantNew:   []int{2},
			wantNow:   -3,
			wantBadge: "Both",
		},
		{
			name: "current conflict, table picked",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, false),
				tableLook(20, "B", true, true),
			},
			file: []export.TakeHomeLook{
				fileLook(1, "A", true, false),
				fileLook(2, "B", true, false),
				fileLook(3, "Cloak", true, true),
			},
			maps:      maps,
			pick:      sideTable,
			wantNew:   []int{2},
			wantNow:   20,
			wantBadge: "Both",
		},
		{
			name: "the current look unlinked at home",
			table: []sqlc.ListCharacterImagesRow{
				tableLook(10, "A", true, true),
				tableLook(20, "B", true, false),
			},
			file:       []export.TakeHomeLook{fileLook(2, "B", true, false)},
			maps:       maps,
			wantUnlink: []int64{10},
			wantNow:    0,
			wantBadge:  titleHome,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := planLooks(
				lookIn{Base: base, Table: tt.table, File: tt.file, Images: tt.maps, Pick: tt.pick},
			)
			if !slices.Equal(p.New, tt.wantNew) || !slices.Equal(p.Unlink, tt.wantUnlink) ||
				p.Current != tt.wantNow {
				t.Errorf(
					"new %v unlink %v current %d; want %v %v %d",
					p.New,
					p.Unlink,
					p.Current,
					tt.wantNew,
					tt.wantUnlink,
					tt.wantNow,
				)
			}
			if len(p.Reveal) != len(tt.wantReveal) {
				t.Errorf("reveal %v, want %v", p.Reveal, tt.wantReveal)
			}
			for id, v := range tt.wantReveal {
				if got, ok := p.Reveal[id]; !ok || got != v {
					t.Errorf("reveal %v, want %v", p.Reveal, tt.wantReveal)
				}
			}
			if p.Group.Badge != tt.wantBadge ||
				tt.wantLine != "" && !slices.Contains(p.Group.Lines, tt.wantLine) {
				t.Errorf("group %+v, want badge %q and line %q", p.Group, tt.wantBadge, tt.wantLine)
			}
			if p.Touched != (tt.wantBadge != "") {
				t.Errorf("touched %v with badge %q", p.Touched, tt.wantBadge)
			}
		})
	}
}

func hbEntry(id int64, slug, name, hash string) export.TakeHomeEntry {
	return export.TakeHomeEntry{
		ID:          id,
		Kind:        enum.KindSpell,
		Slug:        slug,
		Name:        name,
		Data:        json.RawMessage(`{}`),
		ContentHash: hash,
	}
}

func hbRow(hash string) sqlc.CompendiumEntry {
	return sqlc.CompendiumEntry{
		ID:          100,
		Kind:        enum.KindSpell,
		Slug:        "zap",
		Name:        "Zap",
		ContentHash: hash,
	}
}

func TestPlanHomebrew(t *testing.T) {
	maps := map[int64]int64{1: 100}
	base := map[int64]string{100: "h0"}
	tests := []struct {
		name  string
		file  []export.TakeHomeEntry
		table []sqlc.CompendiumEntry
		maps  map[int64]int64
		pick  string
		same  map[string]string // hash -> table slug
		taken []string

		wantUpdate []int64
		wantDelete []int64
		wantCreate []string
		wantRename map[string]string
		wantBadge  string
	}{
		{
			name:       "home changed",
			file:       []export.TakeHomeEntry{hbEntry(1, "zap", "Zap", "h1")},
			table:      []sqlc.CompendiumEntry{hbRow("h0")},
			maps:       maps,
			wantUpdate: []int64{100},
			wantBadge:  titleHome,
		},
		{
			name:  "table changed",
			file:  []export.TakeHomeEntry{hbEntry(1, "zap", "Zap", "h0")},
			table: []sqlc.CompendiumEntry{hbRow("h2")},
			maps:  maps,
		},
		{
			name:  "same change",
			file:  []export.TakeHomeEntry{hbEntry(1, "zap", "Zap", "h1")},
			table: []sqlc.CompendiumEntry{hbRow("h1")},
			maps:  maps,
		},
		{
			name:       "both changed, home",
			file:       []export.TakeHomeEntry{hbEntry(1, "zap", "Zap", "h1")},
			table:      []sqlc.CompendiumEntry{hbRow("h2")},
			maps:       maps,
			wantUpdate: []int64{100},
			wantBadge:  "Both",
		},
		{
			name:      "both changed, table",
			file:      []export.TakeHomeEntry{hbEntry(1, "zap", "Zap", "h1")},
			table:     []sqlc.CompendiumEntry{hbRow("h2")},
			maps:      maps,
			pick:      sideTable,
			wantBadge: "Both",
		},
		{
			name: "gone at the table stays gone",
			file: []export.TakeHomeEntry{hbEntry(1, "zap", "Zap", "h1")},
			maps: maps,
		},
		{
			name:       "deleted at home",
			table:      []sqlc.CompendiumEntry{hbRow("h0")},
			maps:       maps,
			wantDelete: []int64{100},
			wantBadge:  titleHome,
		},
		{
			name:  "deleted at home, table changed it",
			table: []sqlc.CompendiumEntry{hbRow("h2")},
			maps:  maps,
		},
		{
			name:       "new",
			file:       []export.TakeHomeEntry{hbEntry(5, "new-1", "New", "n1")},
			wantCreate: []string{"new-1"},
			wantBadge:  titleHome,
		},
		{
			name:       "new, same as the table's under another slug",
			file:       []export.TakeHomeEntry{hbEntry(5, "new-1", "New", "n1")},
			same:       map[string]string{"n1": "new-9"},
			wantRename: map[string]string{"new-1": "new-9"},
			wantBadge:  titleHome,
		},
		{
			name:  "new, slug taken",
			file:  []export.TakeHomeEntry{hbEntry(5, "new-1", "New", "n1")},
			taken: []string{"new-1", "new-1-2"},
			wantCreate: []string{
				"new-1-3",
			},
			wantRename: map[string]string{"new-1": "new-1-3"},
			wantBadge:  titleHome,
		},
		{
			name: "two new with one slug",
			file: []export.TakeHomeEntry{hbEntry(5, "x", "X", "n1"), hbEntry(6, "x", "X2", "n2")},
			wantCreate: []string{
				"x",
				"x-2",
			},
			wantRename: map[string]string{"x": "x-2"},
			wantBadge:  titleHome,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := planHomebrew(homebrewIn{
				BaseHash: base,
				Maps:     tt.maps,
				File:     tt.file,
				Table:    tt.table,
				Pick:     tt.pick,
				Same:     func(_ enum.CompendiumKind, hash string) (string, bool) { s, ok := tt.same[hash]; return s, ok },
				Taken:    func(_ enum.CompendiumKind, slug string) bool { return slices.Contains(tt.taken, slug) },
			})
			var creates []string
			var ids []int64
			for _, u := range p.Update {
				ids = append(ids, u.ID)
			}
			for _, e := range p.Create {
				creates = append(creates, e.Slug)
			}
			if !slices.Equal(ids, tt.wantUpdate) || !slices.Equal(p.Delete, tt.wantDelete) ||
				!slices.Equal(creates, tt.wantCreate) {
				t.Errorf(
					"update %v delete %v create %v; want %v %v %v",
					ids,
					p.Delete,
					creates,
					tt.wantUpdate,
					tt.wantDelete,
					tt.wantCreate,
				)
			}
			if len(p.Renames) != len(tt.wantRename) {
				t.Errorf("renames %v, want %v", p.Renames, tt.wantRename)
			}
			for k, v := range tt.wantRename {
				if p.Renames[k] != v {
					t.Errorf("renames %v, want %v", p.Renames, tt.wantRename)
				}
			}
			if p.Group.Badge != tt.wantBadge {
				t.Errorf("badge %q, want %q (%+v)", p.Group.Badge, tt.wantBadge, p.Group)
			}
		})
	}
}

func fileBlock(id, version int64, vis enum.Visibility) export.TakeHomeBlock {
	return export.TakeHomeBlock{
		ID:         id,
		Version:    version,
		Visibility: vis,
		Doc:        json.RawMessage(`{}`),
	}
}

func tblBlock(id int64, vis enum.Visibility) sqlc.NoteBlock {
	return sqlc.NoteBlock{ID: id, NoteID: 100, Visibility: vis}
}

func TestPlanNotes(t *testing.T) {
	notes := map[int64]int64{1: 100}
	blocks := map[int64]int64{11: 1100, 12: 1200}
	table := []sqlc.Note{{ID: 100, Kind: enum.NoteGeneric, Title: "Diary"}}
	both := []sqlc.NoteBlock{tblBlock(1100, enum.VisOwner), tblBlock(1200, enum.VisPlayers)}
	note := func(title string, bl ...export.TakeHomeBlock) []export.TakeHomeNote {
		return []export.TakeHomeNote{{ID: 1, Kind: enum.NoteGeneric, Title: title, Blocks: bl}}
	}
	tests := []struct {
		name    string
		file    []export.TakeHomeNote
		notes   map[int64]int64
		table   []sqlc.Note
		blocks  []sqlc.NoteBlock
		titles  map[int64]string
		changed [3]int // changed, new, deleted

		wantReplace []int64
		wantAdd     int
		wantDelete  []int64
		wantNotes   []int64
		wantRename  string
		check       func(*testing.T, notesPlan)
	}{
		{
			name: "edited private block replaces",
			file: note(
				"Diary",
				fileBlock(11, 2, enum.VisOwner),
				fileBlock(12, 0, enum.VisPlayers),
			),
			notes:  notes,
			table:  table,
			blocks: both,
			titles: map[int64]string{
				100: "Diary",
			},
			changed:     [3]int{1, 0, 0},
			wantReplace: []int64{1100},
		},
		{
			name: "edited shared block is left",
			file: note(
				"Diary",
				fileBlock(11, 0, enum.VisOwner),
				fileBlock(12, 3, enum.VisPlayers),
			),
			notes:  notes,
			table:  table,
			blocks: both,
			titles: map[int64]string{100: "Diary"},
		},
		{
			name: "new block is private",
			file: note(
				"Diary",
				fileBlock(11, 0, enum.VisOwner),
				fileBlock(12, 0, enum.VisPlayers),
				fileBlock(13, 0, enum.VisPlayers),
			),
			notes:   notes,
			table:   table,
			blocks:  both,
			titles:  map[int64]string{100: "Diary"},
			changed: [3]int{1, 0, 0},
			wantAdd: 1,
			check: func(t *testing.T, p notesPlan) {
				if b := p.Add[0].Block; b.Visibility != enum.VisOwner || !b.ByCharacter ||
					p.Add[0].NoteID != 100 {
					t.Errorf("added block %+v in note %d", b, p.Add[0].NoteID)
				}
			},
		},
		{
			name:   "deleted private block",
			file:   note("Diary", fileBlock(12, 0, enum.VisPlayers)),
			notes:  notes,
			table:  table,
			blocks: both,
			titles: map[int64]string{
				100: "Diary",
			},
			changed:    [3]int{1, 0, 0},
			wantDelete: []int64{1100},
		},
		{
			name: "new note, empty one skipped", file: []export.TakeHomeNote{
				{
					ID:    7,
					Kind:  enum.NoteGeneric,
					Title: "Fresh",
					Blocks: []export.TakeHomeBlock{
						fileBlock(21, 0, enum.VisPlayers),
						fileBlock(22, 0, enum.VisGM),
					},
				},
				{ID: 8, Kind: enum.NoteGeneric, Title: "Empty"},
			}, notes: map[int64]int64{}, changed: [3]int{0, 1, 0},
			check: func(t *testing.T, p notesPlan) {
				if len(p.NewNotes) != 1 || len(p.NewNotes[0].Blocks) != 2 {
					t.Fatalf("new notes %+v", p.NewNotes)
				}
				for _, b := range p.NewNotes[0].Blocks {
					if b.Visibility != enum.VisOwner || !b.ByCharacter {
						t.Errorf("block %+v is not private", b)
					}
				}
			},
		},
		{
			name:       "note deleted at home keeps shared blocks",
			notes:      notes,
			table:      table,
			blocks:     both,
			changed:    [3]int{1, 0, 0},
			wantDelete: []int64{1100},
			wantNotes:  nil,
		},
		{
			name:       "note deleted at home, nothing else in it",
			notes:      notes,
			table:      table,
			blocks:     []sqlc.NoteBlock{tblBlock(1100, enum.VisOwner)},
			changed:    [3]int{0, 0, 1},
			wantDelete: []int64{1100},
			wantNotes:  []int64{100},
		},
		{
			name:  "note deleted at home keeps a private block written at the table since",
			notes: notes,
			table: table,
			blocks: []sqlc.NoteBlock{
				tblBlock(1100, enum.VisOwner),
				tblBlock(1300, enum.VisOwner),
			},
			changed:    [3]int{1, 0, 0},
			wantDelete: []int64{1100},
		},
		{
			name:  "note deleted at the table stays deleted",
			file:  note("Diary", fileBlock(11, 2, enum.VisOwner)),
			notes: notes,
		},
		{
			name:       "title changed at home",
			file:       note("Journal", fileBlock(11, 0, enum.VisOwner)),
			notes:      notes,
			table:      table,
			blocks:     []sqlc.NoteBlock{tblBlock(1100, enum.VisOwner)},
			titles:     map[int64]string{100: "Diary"},
			changed:    [3]int{1, 0, 0},
			wantRename: "Journal",
		},
		{
			name: "title changed, shared block in the note",
			file: note(
				"Journal",
				fileBlock(11, 0, enum.VisOwner),
				fileBlock(12, 0, enum.VisPlayers),
			),
			notes:  notes,
			table:  table,
			blocks: both,
			titles: map[int64]string{100: "Diary"},
		},
		{
			name:   "title changed, older take-out has no titles",
			file:   note("Journal", fileBlock(11, 0, enum.VisOwner)),
			notes:  notes,
			table:  table,
			blocks: []sqlc.NoteBlock{tblBlock(1100, enum.VisOwner)},
		},
		{
			name:  "title renamed at the table too",
			file:  note("Journal", fileBlock(11, 0, enum.VisOwner)),
			notes: notes,
			table: []sqlc.Note{{ID: 100, Kind: enum.NoteGeneric, Title: "Journal"}},
			blocks: []sqlc.NoteBlock{
				tblBlock(1100, enum.VisOwner),
			},
			titles: map[int64]string{100: "Diary"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := planNotes(
				noteIn{
					BaseTitles:  tt.titles,
					File:        tt.file,
					Notes:       tt.notes,
					Blocks:      blocks,
					Table:       tt.table,
					TableBlocks: tt.blocks,
				},
			)
			var replaced []int64
			for _, w := range p.Replace {
				replaced = append(replaced, w.TableBlock)
			}
			if !slices.Equal(replaced, tt.wantReplace) || len(p.Add) != tt.wantAdd ||
				!slices.Equal(p.DeleteBlocks, tt.wantDelete) {
				t.Errorf(
					"replace %v add %d delete %v; want %v %d %v",
					replaced,
					len(p.Add),
					p.DeleteBlocks,
					tt.wantReplace,
					tt.wantAdd,
					tt.wantDelete,
				)
			}
			if tt.wantNotes != nil && !slices.Equal(p.DeleteNotes, tt.wantNotes) ||
				tt.wantNotes == nil && len(p.DeleteNotes) != 0 {
				t.Errorf("delete notes %v, want %v", p.DeleteNotes, tt.wantNotes)
			}
			if got := [3]int{p.Changed, p.New, p.Deleted}; got != tt.changed {
				t.Errorf("counts %v, want %v", got, tt.changed)
			}
			if tt.wantRename == "" && len(p.Rename) != 0 ||
				tt.wantRename != "" && (len(p.Rename) != 1 || p.Rename[0].Title != tt.wantRename) {
				t.Errorf("rename %+v, want %q", p.Rename, tt.wantRename)
			}
			if tt.check != nil {
				tt.check(t, p)
			}
		})
	}
}

// TestPlanNotesNewCharacter: a new character's file brings its private blocks
// only (the shared ones are the table's own, copied home), and a note left
// empty is skipped.
func TestPlanNotesNewCharacter(t *testing.T) {
	p := planNotes(noteIn{OnlyOwner: true, File: []export.TakeHomeNote{
		{
			ID:    1,
			Kind:  enum.NoteGeneric,
			Title: "Mixed",
			Blocks: []export.TakeHomeBlock{
				fileBlock(1, 0, enum.VisOwner),
				fileBlock(2, 0, enum.VisPlayers),
			},
		},
		{
			ID:     2,
			Kind:   enum.NoteGeneric,
			Title:  "Shared only",
			Blocks: []export.TakeHomeBlock{fileBlock(3, 0, enum.VisPlayers)},
		},
	}})
	if len(p.NewNotes) != 1 || len(p.NewNotes[0].Blocks) != 1 || p.NewNotes[0].Blocks[0].ID != 1 ||
		p.New != 1 {
		t.Errorf("new notes %+v", p.NewNotes)
	}
}

// TestNotesGroupShowsCounts: the review row holds numbers, never a title or text.
func TestNotesGroupShowsCounts(t *testing.T) {
	p := planNotes(noteIn{
		BaseTitles: map[int64]string{100: "Secret title"},
		File: []export.TakeHomeNote{
			{
				ID:     1,
				Kind:   enum.NoteGeneric,
				Title:  "Renamed secret",
				Blocks: []export.TakeHomeBlock{fileBlock(11, 2, enum.VisOwner)},
			},
			{
				ID:     7,
				Kind:   enum.NoteGeneric,
				Title:  "Fresh secret",
				Blocks: []export.TakeHomeBlock{fileBlock(21, 0, enum.VisOwner)},
			},
		},
		Notes:  map[int64]int64{1: 100},
		Blocks: map[int64]int64{11: 1100},
		Table: []sqlc.Note{
			{ID: 100, Title: "Secret title"},
		},
		TableBlocks: []sqlc.NoteBlock{tblBlock(1100, enum.VisOwner)},
	})
	g := p.group(false)
	if len(g.Lines) != 1 || g.Lines[0] != "1 changed · 1 new" {
		t.Errorf("lines %q", g.Lines)
	}
	if got := p.group(true).Lines; len(got) != 1 || got[0] != "1 note" {
		t.Errorf("new character lines %q", got)
	}
}

func TestDocLinks(t *testing.T) {
	mention := func(id string) string { return `{"type":"mention","attrs":{"id":"` + id + `"}}` }
	doc := func(inner ...string) json.RawMessage {
		return json.RawMessage(
			`{"type":"doc","content":[{"type":"paragraph","content":[` + strings.Join(
				inner,
				",",
			) + `]}]}`,
		)
	}
	tests := []struct {
		name string
		doc  json.RawMessage
		want []docLink
	}{
		{
			"note entity image",
			doc(mention("note:3"), mention("entity:4"), mention("image:5")),
			[]docLink{{enum.LinkNote, 3}, {enum.LinkEntity, 4}, {enum.LinkImage, 5}},
		},
		{
			"once each",
			doc(mention("note:3"), mention("note:3"), mention("note:4")),
			[]docLink{{enum.LinkNote, 3}, {enum.LinkNote, 4}},
		},
		{"attachments are not link targets", doc(mention("attachment:3")), nil},
		{"unknown kinds", doc(mention("monster:3"), mention("3")), nil},
		{
			"zero and negative point nowhere",
			doc(mention("note:0"), mention("note:-4"), mention("note:x")),
			nil,
		},
		{"a third part is ignored", doc(mention("note:3:extra")), []docLink{{enum.LinkNote, 3}}},
		{"not a mention", json.RawMessage(`{"type":"text","attrs":{"id":"note:3"}}`), nil},
		{"id that is not a string", json.RawMessage(`{"type":"mention","attrs":{"id":3}}`), nil},
		{"not json", json.RawMessage(`[`), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := docLinks(tt.doc); !slices.Equal(got, tt.want) {
				t.Errorf("links %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRewriteDocMentions(t *testing.T) {
	doc := json.RawMessage(`{"type":"doc","content":[{"type":"paragraph","content":[` +
		`{"type":"mention","attrs":{"id":"note:5"}},{"type":"mention","attrs":{"id":"image:6"}},` +
		`{"type":"mention","attrs":{"id":"entity:7"}},{"type":"mention","attrs":{"id":"attachment:8"}},{"type":"mention","attrs":{"id":"note:-9"}}]}]}`)
	m := &export.HomeIDMaps{Notes: map[int64]int64{50: 5}, Images: map[int64]int64{60: 6}}
	tests := []struct {
		name    string
		rewrite mentionIDs
		want    []string
		absent  []string
	}{
		{
			name:    "taken home: what went home gets its home id, the rest is negated",
			rewrite: homeMentions(m),
			want: []string{
				`"note:50"`,
				`"image:60"`,
				`"entity:7"`,
				`"attachment:8"`,
				`"note:-9"`,
			},
		},
		{
			name:    "brought back: negated is the table's again, a home id goes through the maps, an unknown one is 0",
			rewrite: fromHomeMentions(map[int64]int64{5: 105}, map[int64]int64{6: 106}),
			want: []string{
				`"note:105"`,
				`"image:106"`,
				`"entity:7"`,
				`"attachment:8"`,
				`"note:9"`,
			},
			absent: []string{`"note:-9"`},
		},
		{
			name:    "unresolvable",
			rewrite: fromHomeMentions(nil, nil),
			want:    []string{`"note:0"`, `"image:0"`, `"note:9"`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rewriteDocMentions(doc, tt.rewrite)
			if err != nil {
				t.Fatal(err)
			}
			for _, w := range tt.want {
				if !strings.Contains(string(got), w) {
					t.Errorf("%s lacks %s", got, w)
				}
			}
			for _, w := range tt.absent {
				if strings.Contains(string(got), w) {
					t.Errorf("%s still has %s", got, w)
				}
			}
		})
	}
	// Nothing to change leaves the document as it was.
	plain := json.RawMessage(`{"type":"doc","content":[]}`)
	if got, err := rewriteDocMentions(
		plain,
		homeMentions(m),
	); err != nil ||
		string(got) != string(plain) {
		t.Errorf("plain document %s (%v)", got, err)
	}
}
