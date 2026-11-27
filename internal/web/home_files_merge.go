package web

// What a bring-back (or a new character's file) holds beyond the sheet: looks,
// private notes and homebrew. Each is planned by a pure function over the
// take-out's base, the table now and the file (so the review and Accept plan
// the same way), then applied in Accept's transaction.

import (
	"cmp"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/web/views"
)

// The names of the two sides in the review (badges and conflict columns).
const (
	titleHome  = "Home"
	titleTable = "Table"
)

// Group keys of the rows the review adds after the sheet's.
const (
	groupLooks    = "looks"
	groupNotes    = "notes"
	groupHomebrew = "homebrew"
)

// tableState is what the table holds now of the character's looks, notes and
// homebrew.
type tableState struct {
	Looks       []sqlc.ListCharacterImagesRow
	Notes       []sqlc.Note
	Blocks      []sqlc.NoteBlock
	Attachments map[int64]bool
	Homebrew    []sqlc.CompendiumEntry
}

func loadTableState(ctx context.Context, q *sqlc.Queries, characterID int64) (tableState, error) {
	var t tableState
	var err error
	if t.Looks, err = q.ListCharacterImages(ctx, characterID); err != nil {
		return t, err
	}
	if t.Notes, err = q.ListOwnNotes(ctx, nullID(characterID)); err != nil {
		return t, err
	}
	if t.Blocks, err = q.ListOwnNoteBlocks(ctx, nullID(characterID)); err != nil {
		return t, err
	}
	atts, err := q.ListOwnNoteAttachments(ctx, nullID(characterID))
	if err != nil {
		return t, err
	}
	t.Attachments = map[int64]bool{}
	for _, a := range atts {
		t.Attachments[a.ID] = true
	}
	t.Homebrew, err = q.ListOwnHomebrew(ctx, nullID(characterID))
	return t, err
}

// stamp is the state as text, for the review's was stamp.
func (t tableState) stamp() string {
	var b strings.Builder
	for _, l := range t.Looks {
		fmt.Fprintf(&b, "L%d,%t,%t;", l.ID, l.IsRevealed, l.IsCurrent)
	}
	for _, n := range t.Notes {
		fmt.Fprintf(&b, "N%d,%q;", n.ID, n.Title)
	}
	for _, bl := range t.Blocks {
		fmt.Fprintf(&b, "B%d,%d,%s,%d,%d;", bl.ID, bl.NoteID, bl.Visibility, bl.Version, bl.DocVersion)
	}
	for _, e := range t.Homebrew {
		fmt.Fprintf(&b, "H%d,%s;", e.ID, e.ContentHash)
	}
	return b.String()
}

// merge3 takes the side that changed since base; both changed to different
// values is a conflict (the value returned is then the table's).
func merge3[T comparable](base, table, home T) (v T, conflict bool) {
	switch {
	case home == base || home == table:
		return table, false
	case table == base:
		return home, false
	}
	return table, true
}

// pickedHome is whether the GM's pick for a conflicting group takes the home's
// side (the default).
func pickedHome(pick string) bool { return pick != sideTable }

// sideTable is the pick value of the table's column (dnd5e.SideTable, which
// neutral code does not import).
const sideTable = buildFromTable

// changeGroup is a group's row: Home when only the file changed things, Both
// with two columns when the sides conflict.
func changeGroup(key, label string, lines []string, conflict bool, homeLines, tableLines []string, pick string) views.HomeFileGroup {
	g := views.HomeFileGroup{Key: key, Label: label}
	switch {
	case conflict:
		g.Badge = "Both"
		g.Note = strings.Join(lines, " · ")
		g.Columns = []views.HomeFileColumn{
			{Side: buildFromHome, Title: titleHome, Lines: homeLines, Checked: pickedHome(pick)},
			{Side: sideTable, Title: titleTable, Lines: tableLines, Checked: !pickedHome(pick)},
		}
	case len(lines) > 0:
		g.Badge, g.Lines = titleHome, lines
	}
	return g
}

// --- Looks ---

type lookIn struct {
	// Base is the looks as sent, Table as they are now, File as the home has them.
	Base  []export.TakeHomeBaseLook
	Table []sqlc.ListCharacterImagesRow
	File  []export.TakeHomeLook
	// Images maps a home image id to the table's.
	Images map[int64]int64
	Pick   string
}

// looksPlan is what Accept does to the looks.
type looksPlan struct {
	// New lists the file looks to create (indices of File).
	New []int
	// Unlink lists the table's images no longer worn.
	Unlink []int64
	// Reveal is the new is_revealed of the table's looks that change.
	Reveal map[int64]bool
	// Current is the look that ends up current: a table image id, -(i+1) for
	// New look i, 0 for none.
	Current int64
	Touched bool
	Group   views.HomeFileGroup
}

func planLooks(in lookIn) looksPlan {
	p := looksPlan{Reveal: map[int64]bool{}}
	table := map[int64]sqlc.ListCharacterImagesRow{}
	var tableCurrent int64
	for _, l := range in.Table {
		table[l.ID] = l
		if l.IsCurrent {
			tableCurrent = l.ID
		}
	}
	base := map[int64]export.TakeHomeBaseLook{}
	var baseCurrent int64
	for _, l := range in.Base {
		base[l.ImageID] = l
		if l.IsCurrent {
			baseCurrent = l.ImageID
		}
	}
	name := func(id int64) string {
		switch {
		case id > 0:
			return table[id].Name
		case id < 0:
			return in.File[-id-1].Name
		}
		return "none"
	}

	var lines []string
	var homeCurrent int64
	inFile := map[int64]bool{}
	for i, l := range in.File {
		inFile[l.ImageID] = true
		key := -int64(i + 1)
		if t, mapped := in.Images[l.ImageID]; !mapped {
			p.New = append(p.New, i)
			lines = append(lines, "New: "+l.Name)
		} else {
			key = t
			tl, linked := table[t]
			if b, inBase := base[t]; linked && inBase {
				if rev, _ := merge3(b.IsRevealed, tl.IsRevealed, l.IsRevealed); rev != tl.IsRevealed {
					p.Reveal[t] = rev
					lines = append(lines, tl.Name+": "+map[bool]string{true: "revealed", false: "hidden"}[rev])
				}
			}
		}
		if l.IsCurrent {
			homeCurrent = key
		}
	}
	for _, home := range slices.Sorted(maps.Keys(in.Images)) {
		if t := in.Images[home]; !inFile[home] {
			if tl, linked := table[t]; linked {
				p.Unlink = append(p.Unlink, t)
				lines = append(lines, "Removed: "+tl.Name)
			}
		}
	}

	cur, conflict := merge3(baseCurrent, tableCurrent, homeCurrent)
	if conflict && !pickedHome(in.Pick) {
		cur = tableCurrent
	} else if conflict {
		cur = homeCurrent
	}
	if _, linked := table[cur]; cur > 0 && !linked {
		cur = tableCurrent
	}
	if slices.Contains(p.Unlink, cur) {
		cur = 0
	}
	p.Current = cur
	p.Touched = len(p.New) > 0 || len(p.Unlink) > 0 || len(p.Reveal) > 0 || cur != tableCurrent
	currentLine := func(id int64) string { return "Current: " + name(id) }
	if conflict {
		p.Group = changeGroup(groupLooks, "Looks", lines, true, []string{currentLine(homeCurrent)}, []string{currentLine(tableCurrent)}, in.Pick)
		return p
	}
	if cur != tableCurrent {
		lines = append(lines, currentLine(cur))
	}
	p.Group = changeGroup(groupLooks, "Looks", lines, false, nil, nil, "")
	return p
}

// applyLooks writes the looks plan: files are the saved files of the New looks,
// in their order (empty for a look that is a link). It returns the new images
// by home id.
func applyLooks(ctx context.Context, q *sqlc.Queries, ch sqlc.Character, p looksPlan, file []export.TakeHomeLook, files []savedFile, used map[string]bool) (map[int64]int64, error) {
	created := map[int64]int64{}
	if !p.Touched {
		return created, nil
	}
	for _, t := range p.Unlink {
		if err := q.UnlinkCharacterImage(ctx, sqlc.UnlinkCharacterImageParams{CharacterID: ch.ID, ImageID: t}); err != nil {
			return nil, err
		}
	}
	byIndex := map[int]int64{}
	for n, i := range p.New {
		l := file[i]
		var att sql.NullInt64
		if l.File != "" {
			if n >= len(files) {
				return nil, errors.New("a look's file was not saved")
			}
			sf := files[n]
			a, err := q.CreateAttachment(ctx, sqlc.CreateAttachmentParams{CampaignID: ch.CampaignID, Path: sf.path, Name: sf.name, Mime: sf.mime, Size: sf.size})
			if err != nil {
				return nil, err
			}
			used[sf.path] = true
			att = nullID(a.ID)
		}
		img, err := q.CreateImage(ctx, sqlc.CreateImageParams{
			CampaignID: ch.CampaignID, Name: cmp.Or(l.Name, ch.Name), Url: sql.NullString{String: l.URL, Valid: l.URL != ""}, AttachmentID: att, Description: l.Description,
		})
		if err != nil {
			return nil, err
		}
		if err := q.LinkCharacterImage(ctx, sqlc.LinkCharacterImageParams{CharacterID: ch.ID, ImageID: img.ID}); err != nil {
			return nil, err
		}
		if err := q.SetCharacterImageFlags(ctx, sqlc.SetCharacterImageFlagsParams{IsRevealed: l.IsRevealed, CharacterID: ch.ID, ImageID: img.ID}); err != nil {
			return nil, err
		}
		created[l.ImageID] = img.ID
		byIndex[i] = img.ID
	}
	for t, rev := range p.Reveal {
		if err := q.SetCharacterImageFlags(ctx, sqlc.SetCharacterImageFlagsParams{IsRevealed: rev, CharacterID: ch.ID, ImageID: t}); err != nil {
			return nil, err
		}
	}
	// The flags above cleared is_current on what they touched, and a first look
	// was made current by linking it: set the one current look last.
	if err := q.ClearCharacterCurrentImage(ctx, ch.ID); err != nil {
		return nil, err
	}
	cur := p.Current
	if cur < 0 {
		cur = byIndex[int(-cur-1)]
	}
	if cur > 0 {
		if _, err := q.SetCharacterCurrentImage(ctx, sqlc.SetCharacterCurrentImageParams{CharacterID: ch.ID, ImageID: cur}); err != nil {
			return nil, err
		}
	}
	return created, nil
}

// newLooks is the file's looks that Accept creates (TakeHome for saveHomeFiles).
func (p looksPlan) newLooks(file []export.TakeHomeLook) []export.TakeHomeLook {
	out := make([]export.TakeHomeLook, len(p.New))
	for n, i := range p.New {
		out[n] = file[i]
	}
	return out
}

// --- Notes ---

type noteIn struct {
	// BaseTitles is the notes' titles as sent (nil in an older take-out).
	BaseTitles map[int64]string
	File       []export.TakeHomeNote
	// Notes and Blocks map a home id to the table's.
	Notes, Blocks map[int64]int64
	Table         []sqlc.Note
	TableBlocks   []sqlc.NoteBlock
	// OnlyOwner (a new character's file) leaves out the blocks that were not
	// private: they are the table's shared ones copied home, which the table has.
	OnlyOwner bool
}

// blockWrite is one file block written into an existing table note: a new
// block (TableBlock 0) or the replacement of the table's.
type blockWrite struct {
	NoteID     int64
	TableBlock int64
	Block      export.TakeHomeBlock
}

type noteRename struct {
	NoteID int64
	Kind   enum.NoteKind
	Title  string
}

// notesPlan is what Accept does to the character's notes; only private
// ('owner') blocks and the notes made at home are touched.
type notesPlan struct {
	Replace, Add []blockWrite
	DeleteBlocks []int64
	DeleteNotes  []int64
	Rename       []noteRename
	// NewNotes are the file's notes that Accept creates, with their blocks forced private.
	NewNotes []export.TakeHomeNote
	// Touched are the table's notes that change or go.
	Touched               []int64
	Changed, New, Deleted int
}

// ownerBlocks makes every block private and the character's own.
func ownerBlocks(blocks []export.TakeHomeBlock) []export.TakeHomeBlock {
	out := make([]export.TakeHomeBlock, len(blocks))
	for i, b := range blocks {
		b.Visibility, b.PlayersCanEdit, b.ByCharacter = enum.VisOwner, false, true
		out[i] = b
	}
	return out
}

func planNotes(in noteIn) notesPlan {
	var p notesPlan
	tableNote := map[int64]sqlc.Note{}
	for _, n := range in.Table {
		tableNote[n.ID] = n
	}
	tableBlock := map[int64]sqlc.NoteBlock{}
	blocksOf := map[int64][]sqlc.NoteBlock{}
	for _, b := range in.TableBlocks {
		tableBlock[b.ID] = b
		blocksOf[b.NoteID] = append(blocksOf[b.NoteID], b)
	}
	shared := func(noteID int64) bool {
		return slices.ContainsFunc(blocksOf[noteID], func(b sqlc.NoteBlock) bool { return b.Visibility != enum.VisOwner })
	}
	inFile := map[int64]bool{}
	for _, n := range in.File {
		inFile[n.ID] = true
		t, mapped := in.Notes[n.ID]
		if !mapped {
			blocks := n.Blocks
			if in.OnlyOwner {
				blocks = slices.DeleteFunc(slices.Clone(blocks), func(b export.TakeHomeBlock) bool { return b.Visibility != enum.VisOwner })
			}
			if len(blocks) > 0 {
				nn := n
				nn.Blocks = ownerBlocks(blocks)
				p.NewNotes = append(p.NewNotes, nn)
				p.New++
			}
			continue
		}
		tn, ok := tableNote[t]
		if !ok {
			continue // deleted at the table, it stays deleted
		}
		changed := false
		fileBlock := map[int64]bool{}
		for _, b := range n.Blocks {
			fileBlock[b.ID] = true
			tb, known := in.Blocks[b.ID]
			if !known {
				nb := ownerBlocks([]export.TakeHomeBlock{b})[0]
				p.Add = append(p.Add, blockWrite{NoteID: t, Block: nb})
				changed = true
				continue
			}
			// Home wins over a table edit; only a private block the character owns is replaced.
			if cur, exists := tableBlock[tb]; b.Version > 0 && exists && cur.Visibility == enum.VisOwner {
				p.Replace = append(p.Replace, blockWrite{NoteID: cur.NoteID, TableBlock: tb, Block: b})
				changed = true
			}
		}
		for _, home := range slices.Sorted(maps.Keys(in.Blocks)) {
			tb := in.Blocks[home]
			if cur, exists := tableBlock[tb]; !fileBlock[home] && exists && cur.NoteID == t && cur.Visibility == enum.VisOwner {
				p.DeleteBlocks = append(p.DeleteBlocks, tb)
				changed = true
			}
		}
		if bt, sent := in.BaseTitles[t]; sent && n.Title != bt && n.Title != tn.Title && !shared(t) {
			p.Rename = append(p.Rename, noteRename{NoteID: t, Kind: tn.Kind, Title: n.Title})
			changed = true
		}
		if changed {
			p.Changed++
			p.Touched = append(p.Touched, t)
		}
	}
	sent := map[int64]bool{}
	for _, tb := range in.Blocks {
		sent[tb] = true
	}
	for _, home := range slices.Sorted(maps.Keys(in.Notes)) {
		t := in.Notes[home]
		if _, ok := tableNote[t]; !ok || inFile[home] {
			continue
		}
		var private []int64
		for _, b := range blocksOf[t] {
			// A block written at the table after the take-out never went home: it stays.
			if b.Visibility == enum.VisOwner && sent[b.ID] {
				private = append(private, b.ID)
			}
		}
		left := len(blocksOf[t]) - len(private)
		if len(private) == 0 && left > 0 {
			continue
		}
		p.DeleteBlocks = append(p.DeleteBlocks, private...)
		p.Touched = append(p.Touched, t)
		if left == 0 {
			p.DeleteNotes = append(p.DeleteNotes, t)
			p.Deleted++
		} else {
			p.Changed++
		}
	}
	return p
}

// group is the notes row: counts only, since the GM must not read a player's
// private notes (a new character's file shows how many notes it brings).
func (p notesPlan) group(forNewCharacter bool) views.HomeFileGroup {
	g := views.HomeFileGroup{Key: groupNotes, Label: "Notes"}
	if forNewCharacter {
		if p.New > 0 {
			g.Badge, g.Lines = titleHome, []string{fmt.Sprintf("%d %s", p.New, map[bool]string{true: "note", false: "notes"}[p.New == 1])}
		}
		return g
	}
	var parts []string
	for _, c := range []struct {
		n     int
		label string
	}{{p.Changed, "changed"}, {p.New, "new"}, {p.Deleted, "deleted"}} {
		if c.n > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", c.n, c.label))
		}
	}
	if len(parts) > 0 {
		g.Badge, g.Lines = titleHome, []string{strings.Join(parts, " · ")}
	}
	return g
}

// needFile reports whether a note picture of the file has no live copy at the
// table, so its file must be saved.
type needFile func(homeAttachmentID int64) bool

// attachmentNeeded is needFile for a table state: a picture the table never
// had, or whose copy was deleted.
func attachmentNeeded(attachments map[int64]int64, table map[int64]bool) needFile {
	return func(id int64) bool {
		t, mapped := attachments[id]
		return !mapped || !table[t]
	}
}

// filesToSave are the block's files that need saving.
func filesToSave(b export.TakeHomeBlock, need needFile) []export.TakeHomeBlockFile {
	return slices.DeleteFunc(slices.Clone(b.Files), func(f export.TakeHomeBlockFile) bool { return !need(f.AttachmentID) })
}

// fileNotes is a holder note with the blocks Accept writes, reduced to the
// files that need saving, for saveHomeFiles.
func (p notesPlan) fileNotes(need needFile) []export.TakeHomeNote {
	var blocks []export.TakeHomeBlock
	add := func(b export.TakeHomeBlock) {
		if b.Files = filesToSave(b, need); len(b.Files) > 0 {
			blocks = append(blocks, b)
		}
	}
	for _, w := range slices.Concat(p.Replace, p.Add) {
		add(w.Block)
	}
	for _, n := range p.NewNotes {
		for _, b := range n.Blocks {
			add(b)
		}
	}
	return []export.TakeHomeNote{{Blocks: blocks}}
}

// writtenBlock is a block Accept wrote whose document still needs its ids rewritten.
type writtenBlock struct {
	id, noteID int64
	doc        json.RawMessage
	text       string
	replace    bool
}

// applyNotes writes the notes plan and returns the notes (table ids) it touched.
// Documents are rewritten once every note and picture exists: pictures through
// the attachment map, mentions through the notes and looks of the file.
func applyNotes(ctx context.Context, q *sqlc.Queries, ch sqlc.Character, p notesPlan, m export.HomeIDMaps, newImages map[int64]int64, blockFiles map[int64][]savedFile, need needFile, used map[string]bool) ([]int64, error) {
	touched := slices.Clone(p.Touched)
	for _, id := range p.DeleteBlocks {
		if err := q.DeleteBlock(ctx, id); err != nil {
			return nil, err
		}
	}
	for _, id := range p.DeleteNotes {
		if err := q.DeleteNote(ctx, id); err != nil {
			return nil, err
		}
	}
	for _, r := range p.Rename {
		if err := q.RenameNote(ctx, sqlc.RenameNoteParams{Title: r.Title, Kind: r.Kind, ID: r.NoteID}); err != nil {
			return nil, err
		}
	}

	attachments := maps.Clone(m.Attachments)
	if attachments == nil {
		attachments = map[int64]int64{}
	}
	notes := maps.Clone(m.Notes)
	if notes == nil {
		notes = map[int64]int64{}
	}
	var written []writtenBlock
	addFiles := func(blockID int64, b export.TakeHomeBlock) error {
		files := blockFiles[b.ID]
		for i, bf := range filesToSave(b, need) {
			if i >= len(files) {
				return errors.New("a note picture was not saved")
			}
			sf := files[i]
			a, err := q.CreateAttachment(ctx, sqlc.CreateAttachmentParams{CampaignID: ch.CampaignID, BlockID: nullID(blockID), Path: sf.path, Name: sf.name, Mime: sf.mime, Size: sf.size})
			if err != nil {
				return err
			}
			used[sf.path] = true
			attachments[bf.AttachmentID] = a.ID
		}
		return nil
	}
	addBlock := func(noteID int64, b export.TakeHomeBlock) error {
		id, err := q.CreateTakeHomeBlock(ctx, sqlc.CreateTakeHomeBlockParams{
			NoteID: noteID, Position: b.Position, Visibility: enum.VisOwner, DocJson: string(b.Doc), Text: b.Text, AuthorCharacterID: nullID(ch.ID),
		})
		if err != nil {
			return err
		}
		written = append(written, writtenBlock{id: id, noteID: noteID, doc: b.Doc, text: b.Text})
		return addFiles(id, b)
	}
	for _, n := range p.NewNotes {
		id, err := q.CreateNoteAt(ctx, sqlc.CreateNoteAtParams{
			CampaignID: ch.CampaignID, Kind: n.Kind, Title: n.Title, OwnerCharacterID: nullID(ch.ID), CreatedAt: n.CreatedAt, UpdatedAt: n.UpdatedAt,
		})
		if err != nil {
			return nil, err
		}
		notes[n.ID] = id
		touched = append(touched, id)
		for _, b := range n.Blocks {
			if err := addBlock(id, b); err != nil {
				return nil, err
			}
		}
	}
	for _, w := range p.Add {
		if err := addBlock(w.NoteID, w.Block); err != nil {
			return nil, err
		}
	}
	for _, w := range p.Replace {
		written = append(written, writtenBlock{id: w.TableBlock, noteID: w.NoteID, doc: w.Block.Doc, text: w.Block.Text, replace: true})
		if err := addFiles(w.TableBlock, w.Block); err != nil {
			return nil, err
		}
	}

	images := maps.Clone(m.Images)
	if images == nil {
		images = map[int64]int64{}
	}
	maps.Copy(images, newImages)
	mention := fromHomeMentions(notes, images)
	looks, err := q.ListCharacterImages(ctx, ch.ID)
	if err != nil {
		return nil, err
	}
	lookIDs := make([]int64, len(looks))
	for i, l := range looks {
		lookIDs[i] = l.ID
	}
	for _, w := range written {
		doc, err := rewriteDocAttachments(w.doc, attachments)
		if err != nil {
			return nil, err
		}
		if doc, err = rewriteDocMentions(doc, mention); err != nil {
			return nil, err
		}
		if w.replace {
			if err := q.ReplaceBlockDoc(ctx, sqlc.ReplaceBlockDocParams{DocJson: string(doc), Text: w.text, ID: w.id}); err != nil {
				return nil, err
			}
			if err := q.DeleteBlockSteps(ctx, w.id); err != nil {
				return nil, err
			}
		} else if err := q.SetTakeHomeBlockDoc(ctx, sqlc.SetTakeHomeBlockDocParams{DocJson: string(doc), ID: w.id}); err != nil {
			return nil, err
		}
		if err := writeBlockLinks(ctx, q, ch.CampaignID, w.noteID, w.id, doc, lookIDs); err != nil {
			return nil, err
		}
	}
	return touched, nil
}

// --- Mentions and links ---

// mentionIDs gives the id a note or image mention should have after a move
// between the table and a home computer.
type mentionIDs func(kind enum.LinkTarget, id int64) int64

// mentionMaker builds a mentionIDs from the id maps of a import (home id to
// table id).
type mentionMaker func(m *export.HomeIDMaps) mentionIDs

// invert swaps a map's keys and values.
func invert(m map[int64]int64) map[int64]int64 {
	out := make(map[int64]int64, len(m))
	for k, v := range m {
		out[v] = k
	}
	return out
}

// homeMentions are the ids for a document taken home: a note or image that went
// home gets its home id, any other the table's id negated (an id that resolves
// to nothing at home, which the editor's links skip).
func homeMentions(m *export.HomeIDMaps) mentionIDs {
	notes, images := invert(m.Notes), invert(m.Images)
	return func(kind enum.LinkTarget, id int64) int64 {
		if id <= 0 {
			return id
		}
		target := images
		if kind == enum.LinkNote {
			target = notes
		}
		if home, ok := target[id]; ok {
			return home
		}
		return -id
	}
}

// fromHomeMentions are the ids for a document that came back: a negated id is
// the table's again, a home id goes through notes and images (home id to table
// id) or becomes 0, which links nowhere.
func fromHomeMentions(notes, images map[int64]int64) mentionIDs {
	return func(kind enum.LinkTarget, id int64) int64 {
		if id < 0 {
			return -id
		}
		if kind == enum.LinkNote {
			return notes[id]
		}
		return images[id]
	}
}

// tableMentions is homeMentions' counterpart for importNotes on the table: the
// import's maps go table id to home id.
func tableMentions(m *export.HomeIDMaps) mentionIDs {
	return fromHomeMentions(invert(m.Notes), invert(m.Images))
}

// rewriteDocMentions gives every note and image mention of a ProseMirror
// document the id rewrite returns (the other kinds are left alone: attachments
// go through rewriteDocAttachments, entities keep their ids).
func rewriteDocMentions(doc json.RawMessage, rewrite mentionIDs) (json.RawMessage, error) {
	var v any
	if err := json.Unmarshal(doc, &v); err != nil {
		return nil, err
	}
	changed := false
	var walk func(any)
	walk = func(n any) {
		switch n := n.(type) {
		case map[string]any:
			if attrs, ok := n["attrs"].(map[string]any); ok && n["type"] == mentionNode {
				if s, ok := attrs["id"].(string); ok {
					kind, rest, _ := strings.Cut(s, ":")
					id, err := strconv.ParseInt(rest, 10, 64)
					if k := enum.LinkTarget(kind); err == nil && (k == enum.LinkNote || k == enum.LinkImage) {
						if now := rewrite(k, id); now != id {
							attrs["id"] = fmt.Sprintf("%s:%d", k, now)
							changed = true
						}
					}
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
	if !changed {
		return doc, nil
	}
	return json.Marshal(v)
}

// mentionNode is the ProseMirror node type of a [[mention]].
const mentionNode = "mention"

// docLink is a link a document holds, as the editor's snapshot lists it.
type docLink struct {
	Type enum.LinkTarget
	ID   int64
}

// docLinks lists the note, entity and image mentions of a document, once each:
// a port of links() in notes.js (attached pictures are not link targets, and
// ids of 0 or less point nowhere).
func docLinks(doc json.RawMessage) []docLink {
	var v any
	if json.Unmarshal(doc, &v) != nil {
		return nil
	}
	var out []docLink
	var walk func(any)
	walk = func(n any) {
		switch n := n.(type) {
		case map[string]any:
			if attrs, ok := n["attrs"].(map[string]any); ok && n["type"] == mentionNode {
				if s, ok := attrs["id"].(string); ok {
					parts := strings.Split(s, ":")
					if len(parts) >= 2 {
						k := enum.LinkTarget(parts[0])
						id, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
						l := docLink{Type: k, ID: id}
						if (k == enum.LinkNote || k == enum.LinkEntity || k == enum.LinkImage) && err == nil && id > 0 && !slices.Contains(out, l) {
							out = append(out, l)
						}
					}
				}
			}
			walk(n["content"])
		case []any:
			for _, c := range n {
				walk(c)
			}
		}
	}
	walk(v)
	return out
}

// writeBlockLinks rebuilds a block's links from its document, as a player's
// snapshot would: an image link stays only if the block had it or the image is
// one of the character's looks. A target outside the campaign is skipped.
func writeBlockLinks(ctx context.Context, q *sqlc.Queries, campaignID, noteID, blockID int64, doc json.RawMessage, lookIDs []int64) error {
	images, err := q.ListBlockImageLinks(ctx, blockID)
	if err != nil {
		return err
	}
	if err := q.ClearBlockLinks(ctx, blockID); err != nil {
		return err
	}
	for _, l := range docLinks(doc) {
		if l.Type == enum.LinkImage && !slices.Contains(images, l.ID) && !slices.Contains(lookIDs, l.ID) {
			continue
		}
		if ok, err := linkInCampaign(ctx, q, campaignID, l); err != nil {
			return err
		} else if !ok {
			continue
		}
		note, entity, image := linkTarget(l.Type, l.ID)
		if err := q.AddBlockLink(ctx, sqlc.AddBlockLinkParams{SourceNoteID: noteID, SourceBlockID: blockID, TargetNoteID: note, TargetEntityID: entity, TargetImageID: image}); err != nil {
			return err
		}
	}
	return nil
}

// linkInCampaign reports whether a link's target exists in the campaign.
func linkInCampaign(ctx context.Context, q *sqlc.Queries, campaignID int64, l docLink) (bool, error) {
	var got int64
	var err error
	switch l.Type {
	case enum.LinkNote:
		var n sqlc.Note
		n, err = q.GetNote(ctx, l.ID)
		got = n.CampaignID
	case enum.LinkEntity:
		var e sqlc.Entity
		e, err = q.GetEntity(ctx, l.ID)
		got = e.CampaignID
	case enum.LinkImage:
		var i sqlc.GetImageRow
		i, err = q.GetImage(ctx, l.ID)
		got = i.CampaignID
	}
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return err == nil && got == campaignID, err
}

// --- Homebrew ---

type homebrewIn struct {
	// BaseHash is the content hash of each own entry as sent, by table id.
	BaseHash map[int64]string
	// Maps maps a home entry id to the table's.
	Maps  map[int64]int64
	File  []export.TakeHomeEntry
	Table []sqlc.CompendiumEntry
	Pick  string
	// Same finds a homebrew entry the campaign sees with this content (its slug).
	Same func(kind enum.CompendiumKind, hash string) (slug string, ok bool)
	// Taken says whether a homebrew entry of the kind has the slug.
	Taken func(kind enum.CompendiumKind, slug string) bool
}

type homebrewUpdate struct {
	ID    int64
	Entry export.TakeHomeEntry
}

// homebrewPlan is what Accept does to the character's homebrew.
type homebrewPlan struct {
	Update []homebrewUpdate
	Delete []int64
	// Create are the entries to insert, with the slug they will have.
	Create []export.TakeHomeEntry
	// Renames are the homebrew spells whose slug differs at the table (old to new).
	Renames map[string]string
	Group   views.HomeFileGroup
}

func (p homebrewPlan) touched() bool { return len(p.Update)+len(p.Delete)+len(p.Create) > 0 }

func planHomebrew(in homebrewIn) homebrewPlan {
	p := homebrewPlan{Renames: map[string]string{}}
	table := map[int64]sqlc.CompendiumEntry{}
	for _, e := range in.Table {
		table[e.ID] = e
	}
	var lines, conflictHome, conflictTable []string
	var conflicts []homebrewUpdate
	inFile := map[int64]bool{}
	taken := map[string]bool{} // slugs this plan gives
	slugKey := func(k enum.CompendiumKind, s string) string { return string(k) + "/" + s }
	for _, e := range in.File {
		inFile[e.ID] = true
		if t, mapped := in.Maps[e.ID]; mapped {
			row, ok := table[t]
			base, sent := in.BaseHash[t]
			switch {
			case !ok || !sent || e.ContentHash == base || e.ContentHash == row.ContentHash:
				// gone at the table, unchanged at home, or the same change on both sides
			case row.ContentHash == base:
				p.Update = append(p.Update, homebrewUpdate{ID: t, Entry: e})
				lines = append(lines, "Changed: "+e.Name)
			default:
				conflicts = append(conflicts, homebrewUpdate{ID: t, Entry: e})
				conflictHome = append(conflictHome, "Changed: "+e.Name)
				conflictTable = append(conflictTable, "Changed: "+row.Name)
			}
			continue
		}
		if slug, ok := in.Same(e.Kind, e.ContentHash); ok {
			lines = append(lines, "Same as the table's: "+e.Name)
			if slug != e.Slug && e.Kind == enum.KindSpell {
				p.Renames[e.Slug] = slug
			}
			continue
		}
		slug := e.Slug
		for n := 2; in.Taken(e.Kind, slug) || taken[slugKey(e.Kind, slug)]; n++ {
			slug = fmt.Sprintf("%s-%d", e.Slug, n)
		}
		taken[slugKey(e.Kind, slug)] = true
		if slug != e.Slug && e.Kind == enum.KindSpell {
			p.Renames[e.Slug] = slug
		}
		ne := e
		ne.Slug = slug
		p.Create = append(p.Create, ne)
		lines = append(lines, "New: "+e.Name)
	}
	for _, home := range slices.Sorted(maps.Keys(in.Maps)) {
		t := in.Maps[home]
		row, ok := table[t]
		if base, sent := in.BaseHash[t]; ok && sent && !inFile[home] && row.ContentHash == base {
			p.Delete = append(p.Delete, t)
			lines = append(lines, "Deleted: "+row.Name)
		}
	}
	if len(conflicts) > 0 && pickedHome(in.Pick) {
		p.Update = append(p.Update, conflicts...)
	}
	p.Group = changeGroup(groupHomebrew, "Homebrew", lines, len(conflicts) > 0, conflictHome, conflictTable, in.Pick)
	return p
}

func applyHomebrew(ctx context.Context, q *sqlc.Queries, ch sqlc.Character, p homebrewPlan) error {
	for _, u := range p.Update {
		e := u.Entry
		if err := q.UpdateHomebrew(ctx, sqlc.UpdateHomebrewParams{Name: e.Name, DataJson: string(e.Data), BodyText: e.BodyText, ContentHash: e.ContentHash, ID: u.ID}); err != nil {
			return err
		}
	}
	for _, id := range p.Delete {
		if err := q.DeleteHomebrew(ctx, id); err != nil {
			return err
		}
	}
	for _, e := range p.Create {
		row, err := q.CreateHomebrew(ctx, sqlc.CreateHomebrewParams{
			Kind: e.Kind, Slug: e.Slug, Name: e.Name, DataJson: string(e.Data), BodyText: e.BodyText, ContentHash: e.ContentHash,
			OwnerCharacterID: nullID(ch.ID), Scoped: true,
		})
		if err != nil {
			return err
		}
		if err := q.LinkEntryCampaign(ctx, sqlc.LinkEntryCampaignParams{EntryID: row.ID, CampaignID: ch.CampaignID}); err != nil {
			return err
		}
	}
	return nil
}

// --- The plan together ---

// mapsOf is the id maps of a bring-back (empty for a new character's file).
func mapsOf(file *export.TakeHomeFile) export.HomeIDMaps {
	if file.Maps == nil {
		return export.HomeIDMaps{}
	}
	return *file.Maps
}

// homeExtras is the plan of everything beyond the sheet.
type homeExtras struct {
	table   tableState
	looks   looksPlan
	notes   notesPlan
	hb      homebrewPlan
	newChar bool
}

// planExtras plans the looks, notes and homebrew of file for the campaign. A
// new character's file (characterID 0) has no table state, base or maps, so all
// of it is new.
func planExtras(ctx context.Context, q *sqlc.Queries, campaignID, characterID int64, file *export.TakeHomeFile, base export.TakeHomeBase, form func(string) string) (homeExtras, error) {
	x := homeExtras{newChar: characterID == 0}
	var err error
	if !x.newChar {
		if x.table, err = loadTableState(ctx, q, characterID); err != nil {
			return x, err
		}
	}
	m := mapsOf(file)
	x.looks = planLooks(lookIn{Base: base.Looks, Table: x.table.Looks, File: file.Looks, Images: m.Images, Pick: form(views.PickParam(groupLooks))})
	x.notes = planNotes(noteIn{BaseTitles: base.Titles, File: file.Notes, Notes: m.Notes, Blocks: m.Blocks, Table: x.table.Notes, TableBlocks: x.table.Blocks, OnlyOwner: x.newChar})

	var lookupErr error
	x.hb = planHomebrew(homebrewIn{
		BaseHash: base.Homebrew, Maps: m.Homebrew, File: file.Homebrew, Table: x.table.Homebrew, Pick: form(views.PickParam(groupHomebrew)),
		Same: func(kind enum.CompendiumKind, hash string) (string, bool) {
			row, err := q.FindCampaignHomebrewByHash(ctx, sqlc.FindCampaignHomebrewByHashParams{Kind: kind, ContentHash: hash, CampaignID: campaignID})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				lookupErr = err
			}
			return row.Slug, err == nil
		},
		Taken: func(kind enum.CompendiumKind, slug string) bool {
			_, err := q.GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: kind, Slug: slug, SourceKey: homebrewSource})
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				lookupErr = err
			}
			return err == nil
		},
	})
	return x, lookupErr
}

// groups are the review's rows for the extras; a new character's file shows
// only the ones that bring something.
func (x homeExtras) groups() []views.HomeFileGroup {
	all := []views.HomeFileGroup{x.looks.Group, x.notes.group(x.newChar), x.hb.Group}
	if x.newChar {
		all = slices.DeleteFunc(all, func(g views.HomeFileGroup) bool { return g.Badge == "" })
	}
	return all
}

// needFile is the test for a note picture that must be saved.
func (x homeExtras) needFile(attachments map[int64]int64) needFile {
	return attachmentNeeded(attachments, x.table.Attachments)
}

// saveContent is the part of file whose files Accept must save before its
// transaction: the new looks and the pictures of the blocks it writes.
func (x homeExtras) saveContent(file *export.TakeHomeFile, attachments map[int64]int64) export.TakeHome {
	return export.TakeHome{Looks: x.looks.newLooks(file.Looks), Notes: x.notes.fileNotes(x.needFile(attachments))}
}

// extrasResult is what applying the extras changed, for publishing.
type extrasResult struct {
	notes      []int64
	compendium bool
}

// applyExtras writes the looks, then the notes (which link to the looks), then
// the homebrew of a bring-back.
func applyExtras(ctx context.Context, q *sqlc.Queries, rv homeReview, lookFiles []savedFile, blockFiles map[int64][]savedFile, used map[string]bool) (extrasResult, error) {
	var res extrasResult
	ch, m := rv.bb.ch, mapsOf(rv.file)
	images, err := applyLooks(ctx, q, ch, rv.x.looks, rv.file.Looks, lookFiles, used)
	if err != nil {
		return res, err
	}
	if res.notes, err = applyNotes(ctx, q, ch, rv.x.notes, m, images, blockFiles, rv.x.needFile(m.Attachments), used); err != nil {
		return res, err
	}
	res.compendium = rv.x.hb.touched()
	return res, applyHomebrew(ctx, q, ch, rv.x.hb)
}

// applyRenames gives the home side's sheets the table's slugs for the spells
// the homebrew plan renamed: the character's, the level-ups' and the fallback's.
func applyRenames(renames map[string]string, cf *export.CharacterFile, file *export.TakeHomeFile) error {
	if len(renames) == 0 {
		return nil
	}
	rename := func(sheet *json.RawMessage) error {
		s, err := renameSpells(string(*sheet), renames)
		*sheet = json.RawMessage(s)
		return err
	}
	if err := rename(&cf.Sheet); err != nil {
		return err
	}
	for i := range file.LevelUps {
		if err := rename(&file.LevelUps[i].Sheet); err != nil {
			return err
		}
	}
	if fb := file.Fallback; fb != nil {
		if err := rename(&fb.Sheet); err != nil {
			return err
		}
		for i := range fb.LevelUps {
			if err := rename(&fb.LevelUps[i].Sheet); err != nil {
				return err
			}
		}
	}
	return nil
}
