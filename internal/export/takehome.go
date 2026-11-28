package export

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"gm-session-manager/internal/enum"
)

// The take-home zip: a character played on another computer. One file per
// part, written and read here so the format lives in one place. The sheet
// stays opaque JSON; the web layer checks it.
//
//	manifest.json            format, kind, versions, token, campaign, the table's character id
//	maps.json                bring-back only: home id to table id of what came from the table
//	fallback.json            bring-back only, optional: the build-only fallback (sheet, max HP,
//
// level-ups; home ids, not in maps.json)
//
//	character.json           CharacterFile
//	levelups.json            the revert history
//	looks.json               the character's images
//	notes.json               the notes the character owns, with every block both sides may read
//	homebrew.json            the character's own homebrew entries
//	compendium/sources.json  the campaign's sources with their priority
//	compendium/entries.json  their entries (no monsters), and the homebrew the campaign sees
//	files/<name>             files of looks and note blocks
//
// A campaign-starter file holds only manifest.json and the two compendium parts.
const (
	// TakeHomeFormat is the zip's layout version;
	// a reader accepts 1 up to this and refuses a newer one.
	// 2 added the campaign starter kind, TakeHomeCampaign.ID and TakeHomeManifest.MadeAtHome.
	TakeHomeFormat = 2

	manifestFile  = "manifest.json"
	mapsFile      = "maps.json"
	fallbackFile  = "fallback.json"
	characterFile = "character.json"
	levelUpsFile  = "levelups.json"
	looksFile     = "looks.json"
	notesFile     = "notes.json"
	homebrewFile  = "homebrew.json"
	sourcesFile   = "compendium/sources.json"
	entriesFile   = "compendium/entries.json"
	filesPrefix   = "files/"

	// MaxFileBytes caps one attached file when reading; it matches the app's
	// attachment cap.
	MaxFileBytes = 20 << 20
	// maxJSONBytes caps each small JSON part, maxEntriesBytes the compendium's.
	maxJSONBytes    = 16 << 20
	maxEntriesBytes = 64 << 20
)

// Kind says which way a take-home zip travels.
type Kind string

const (
	// KindTakeHome goes from the table to a home computer.
	KindTakeHome Kind = "take-home"
	// KindBringBack goes from a home computer back to the table; its ids are the home's.
	KindBringBack Kind = "bring-back"
	// KindNewCharacter is a character made at home: the bring-back format without a take-out token,
	// maps or fallback.
	KindNewCharacter Kind = "new-character"
	// KindStarter is a campaign's books and entries with no character, for a player to start
	// characters at home.
	KindStarter Kind = "campaign-starter"
)

// TakeHomeManifest says what the file is and where it came from.
type TakeHomeManifest struct {
	Format int  `json:"format"`
	Kind   Kind `json:"kind"`
	// SavedAt is when a bring-back was saved at home; zero for a take-home.
	SavedAt     time.Time        `json:"saved_at,omitzero"`
	AppVersion  string           `json:"app_version"`
	Token       string           `json:"token"`
	ServerID    string           `json:"server_id"`
	Campaign    TakeHomeCampaign `json:"campaign"`
	CharacterID int64            `json:"character_id"`
	CreatedAt   time.Time        `json:"created_at"`
	// MadeAtHome ties a character made at home to its copy at the table; only new-character and
	// take-home files hold it.
	MadeAtHome string `json:"made_at_home,omitempty"`
}

// TakeHomeCampaign is the campaign the character plays in.
type TakeHomeCampaign struct {
	// ID is the table's campaign id; a starter file needs it, the others leave it out.
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name"`
	System  string `json:"system"`
	Edition string `json:"edition"`
}

// TakeHomeLevelUp is one revert snapshot.
type TakeHomeLevelUp struct {
	ID        int64           `json:"id"`
	MaxHP     int64           `json:"max_hp"`
	Sheet     json.RawMessage `json:"sheet"`
	CreatedAt time.Time       `json:"created_at"`
}

// TakeHomeLook is one look: a link (URL) or a file in the zip.
type TakeHomeLook struct {
	ImageID int64 `json:"image_id"`
	// AttachmentID is the table's attachment of File.
	AttachmentID int64  `json:"attachment_id,omitempty"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	URL          string `json:"url"`
	// File is the zip path of the picture, "" for a link.
	File       string `json:"file"`
	FileName   string `json:"file_name,omitempty"`
	IsRevealed bool   `json:"is_revealed"`
	IsCurrent  bool   `json:"is_current"`
}

// TakeHomeBlockFile is a file attached to a block.
type TakeHomeBlockFile struct {
	AttachmentID int64  `json:"attachment_id"`
	Name         string `json:"name"`
	File         string `json:"file"`
}

// TakeHomeBlock is one note block.
type TakeHomeBlock struct {
	ID             int64               `json:"id"`
	Position       int64               `json:"position"`
	Visibility     enum.Visibility     `json:"visibility"`
	PlayersCanEdit bool                `json:"players_can_edit"`
	Doc            json.RawMessage     `json:"doc"`
	Version        int64               `json:"version"`
	Text           string              `json:"text"`
	ByCharacter    bool                `json:"by_character"`
	Files          []TakeHomeBlockFile `json:"files"`
}

// TakeHomeNote is a note the character owns.
type TakeHomeNote struct {
	ID        int64           `json:"id"`
	Kind      enum.NoteKind   `json:"kind"`
	Title     string          `json:"title"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Blocks    []TakeHomeBlock `json:"blocks"`
}

// TakeHomeEntry is a compendium entry; ID is the table's id, set only for the
// character's own homebrew.
type TakeHomeEntry struct {
	ID          int64               `json:"id,omitempty"`
	Kind        enum.CompendiumKind `json:"kind"`
	Slug        string              `json:"slug"`
	Name        string              `json:"name"`
	SourceKey   string              `json:"source_key"`
	Data        json.RawMessage     `json:"data"`
	BodyText    string              `json:"body_text"`
	ContentHash string              `json:"content_hash"`
	Scoped      bool                `json:"scoped"`
}

// TakeHomeSource is a source the campaign allows.
type TakeHomeSource struct {
	SourceKey   string `json:"source_key"`
	Edition     string `json:"edition"`
	Priority    int64  `json:"priority"`
	Published   string `json:"published"`
	License     string `json:"license"`
	Attribution string `json:"attribution"`
}

// TakeHomeFallback is a bring-back's build-only copy of the character: the
// table uses it if the GM declines a source added at home. Its level-up ids
// are home ids and are not in the maps.
type TakeHomeFallback struct {
	MaxHP    int64             `json:"max_hp"`
	Sheet    json.RawMessage   `json:"sheet"`
	LevelUps []TakeHomeLevelUp `json:"level_ups"`
}

// TakeHome is the whole content of a take-home file.
type TakeHome struct {
	Manifest  TakeHomeManifest
	Character CharacterFile
	LevelUps  []TakeHomeLevelUp
	Looks     []TakeHomeLook
	Notes     []TakeHomeNote
	Homebrew  []TakeHomeEntry
	Sources   []TakeHomeSource
	Entries   []TakeHomeEntry
	// Maps is written (as maps.json) only in a bring-back.
	Maps *HomeIDMaps
	// Fallback is written (as fallback.json) only in a bring-back, when one was prepared.
	Fallback *TakeHomeFallback
}

// TakeHomeBase is what the table recorded when the file was written, the
// merge base of the bring-back.
type TakeHomeBase struct {
	Character TakeHomeBaseCharacter `json:"character"`
	LevelUps  []int64               `json:"level_ups"`
	Looks     []TakeHomeBaseLook    `json:"looks"`
	Notes     []int64               `json:"notes"`
	// Titles maps each note sent to its title; a take-out recorded before the
	// bring-back needed it has none.
	Titles map[int64]string `json:"titles,omitempty"`
	// Blocks maps each block sent to its version.
	Blocks map[int64]int64 `json:"blocks"`
	// Homebrew maps each own entry sent to its content hash.
	Homebrew map[int64]string `json:"homebrew"`
}

// TakeHomeBaseCharacter is the character as sent.
type TakeHomeBaseCharacter struct {
	// Name is empty in a take-out recorded before the review needed it.
	Name   string          `json:"name"`
	Sheet  json.RawMessage `json:"sheet"`
	HP     int64           `json:"hp"`
	MaxHP  int64           `json:"max_hp"`
	TempHP int64           `json:"temp_hp"`
}

// TakeHomeBaseLook is a look's flags as sent.
type TakeHomeBaseLook struct {
	ImageID    int64 `json:"image_id"`
	IsRevealed bool  `json:"is_revealed"`
	IsCurrent  bool  `json:"is_current"`
}

// ZipFile is the path a file of an attachment has in the zip.
func ZipFile(attachmentPath string) string { return filesPrefix + attachmentPath }

// WriteTakeHome writes t as a zip, copying each referenced file from
// attachmentsDir. A file missing on disk is dropped (a look with no link then
// goes too) and listed in files/MISSING.txt, so one lost file does not stop
// the trip.
func WriteTakeHome(w io.Writer, t TakeHome, attachmentsDir string) error {
	zw := zip.NewWriter(w)
	missing := DropMissing(&t, attachmentsDir)

	parts := []struct {
		name string
		v    any
	}{
		{
			manifestFile,
			t.Manifest,
		},
		{sourcesFile, nonNil(t.Sources)},
		{entriesFile, nonNil(t.Entries)},
	}
	if t.Manifest.Kind != KindStarter {
		char, err := t.Character.JSON()
		if err != nil {
			return err
		}
		if err := writeFile(zw, characterFile, char); err != nil {
			return err
		}
		parts = append(parts, []struct {
			name string
			v    any
		}{{levelUpsFile, nonNil(t.LevelUps)}, {looksFile, nonNil(t.Looks)}, {notesFile, nonNil(t.Notes)}, {homebrewFile, nonNil(t.Homebrew)}}...)
	}
	if t.Maps != nil {
		parts = append(parts, struct {
			name string
			v    any
		}{mapsFile, t.Maps})
	}
	if t.Fallback != nil {
		t.Fallback.LevelUps = nonNil(t.Fallback.LevelUps)
		parts = append(parts, struct {
			name string
			v    any
		}{fallbackFile, t.Fallback})
	}
	for _, p := range parts {
		body, err := json.MarshalIndent(p.v, "", "  ")
		if err != nil {
			return err
		}
		if err := writeFile(zw, p.name, body); err != nil {
			return err
		}
	}

	written := map[string]bool{}
	copyFile := func(file string) error {
		if file == "" || written[file] {
			return nil
		}
		written[file] = true
		//nolint:gosec // see DropMissing
		body, err := os.ReadFile(filepath.Join(attachmentsDir, path.Base(file)))
		if err != nil {
			return err
		}
		return writeFile(zw, file, body)
	}
	for _, l := range t.Looks {
		if err := copyFile(l.File); err != nil {
			return err
		}
	}
	for _, n := range t.Notes {
		for _, b := range n.Blocks {
			for _, f := range b.Files {
				if err := copyFile(f.File); err != nil {
					return err
				}
			}
		}
	}
	if len(missing) > 0 {
		list := "Files that were not on disk, so not taken home:\n" + strings.Join(
			missing,
			"\n",
		) + "\n"
		if err := writeFile(zw, filesPrefix+"MISSING.txt", []byte(list)); err != nil {
			return err
		}
	}
	return zw.Close()
}

// DropMissing removes what refers to a file missing from attachmentsDir (a look
// with no link goes too) and returns the missing names.
func DropMissing(t *TakeHome, attachmentsDir string) []string {
	var missing []string
	have := func(file string) bool {
		if file == "" {
			return false
		}
		_, err := os.Stat(
			filepath.Join(attachmentsDir, path.Base(file)),
		) //nolint:gosec // our random file name, kept inside attachmentsDir by Base
		if err != nil {
			missing = append(missing, path.Base(file))
		}
		return err == nil
	}

	var looks []TakeHomeLook
	for _, l := range t.Looks {
		if l.File != "" && !have(l.File) {
			if l.URL == "" {
				continue
			}
			l.File, l.FileName = "", ""
		}
		looks = append(looks, l)
	}
	t.Looks = looks
	notes := make([]TakeHomeNote, len(t.Notes))
	for i, n := range t.Notes {
		blocks := make([]TakeHomeBlock, len(n.Blocks))
		for j, b := range n.Blocks {
			var files []TakeHomeBlockFile
			for _, f := range b.Files {
				if have(f.File) {
					files = append(files, f)
				}
			}
			b.Files = files
			blocks[j] = b
		}
		n.Blocks = blocks
		notes[i] = n
	}
	t.Notes = notes
	return missing
}

// nonNil makes a nil slice an empty JSON list, never null.
func nonNil[T any](s []T) []T {
	if s == nil {
		return []T{}
	}
	return s
}

// TakeHomeFile is a parsed take-home zip.
type TakeHomeFile struct {
	TakeHome
	zr *zip.Reader
}

// ReadTakeHome parses a take-home zip of the kind want. It refuses a newer
// format, the other kind, unknown keys in the envelope, oversize parts and
// files the JSON names but the zip lacks. The sheets inside are not looked at.
func ReadTakeHome(r io.ReaderAt, size int64, want Kind) (*TakeHomeFile, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return nil, fmt.Errorf("it is not a zip file (%w)", err)
	}
	f := &TakeHomeFile{zr: zr}

	var head struct {
		Format     int    `json:"format"`
		AppVersion string `json:"app_version"`
	}
	if err := f.readJSON(manifestFile, maxJSONBytes, &head, false); err != nil {
		return nil, err
	}
	if head.Format < 1 || head.Format > TakeHomeFormat {
		return nil, fmt.Errorf(
			"this take-home file (from app %s) has format %d; this app reads formats 1 to %d",
			head.AppVersion,
			head.Format,
			TakeHomeFormat,
		)
	}
	if err := f.readJSON(manifestFile, maxJSONBytes, &f.Manifest, true); err != nil {
		return nil, err
	}
	if k := f.Manifest.Kind; k != want {
		switch {
		case k == KindStarter && want == KindTakeHome:
			return nil, errors.New("this is a campaign starter file: open it in the player app")
		case k == KindStarter:
			return nil, errors.New(
				"this is a campaign starter file: open it in the player app, not at the table",
			)
		case want == KindStarter:
			return nil, fmt.Errorf("this is a %s file, not a campaign starter file", k)
		case want == KindTakeHome && k == KindBringBack:
			return nil, errors.New("this is a bring-back file: bring it to the table")
		case want == KindTakeHome && k == KindNewCharacter:
			return nil, errors.New("this is a new character file: give it to your GM at the table")
		case k == KindTakeHome && (want == KindBringBack || want == KindNewCharacter):
			return nil, errors.New("this is a take-home file: open it in the player app")
		case (k == KindBringBack || k == KindNewCharacter) && (want == KindBringBack || want == KindNewCharacter):
			return nil, fmt.Errorf("this is a %s file, not a %s file", k, want)
		}
		// Callers say "not a take-home file" first; this says why and what to do.
		return nil, fmt.Errorf(
			"it comes from app %s and is not a take-home, bring-back or new character file. Take the character home again from the table",
			f.Manifest.AppVersion,
		)
	}
	if want == KindNewCharacter && (f.Manifest.Token != "" || f.Manifest.CharacterID != 0) {
		return nil, errors.New("a new-character file must not name a table character or take-out")
	}
	if f.Manifest.MadeAtHome != "" && want != KindNewCharacter && want != KindTakeHome {
		return nil, fmt.Errorf("a %s file must not name a character made at home", want)
	}
	if want == KindStarter {
		return f.readStarter()
	}
	if has := f.find(mapsFile) != nil; has != (want == KindBringBack) {
		if has {
			return nil, fmt.Errorf("a %s file must not hold %s", want, mapsFile)
		}
		return nil, errors.New("the file lacks " + mapsFile)
	}
	hasFallback := f.find(fallbackFile) != nil
	if hasFallback && want != KindBringBack {
		return nil, fmt.Errorf("a %s file must not hold %s", want, fallbackFile)
	}
	if hasFallback {
		f.Fallback = &TakeHomeFallback{}
		if err := f.readJSON(fallbackFile, maxJSONBytes, f.Fallback, true); err != nil {
			return nil, err
		}
	}
	if want == KindBringBack {
		f.Maps = &HomeIDMaps{}
		if err := f.readJSON(mapsFile, maxJSONBytes, f.Maps, true); err != nil {
			return nil, err
		}
	}
	for _, p := range []struct {
		name  string
		limit int64
		v     any
	}{
		{characterFile, maxJSONBytes, &f.Character}, {levelUpsFile, maxJSONBytes, &f.LevelUps},
		{looksFile, maxJSONBytes, &f.Looks}, {notesFile, maxJSONBytes, &f.Notes},
		{homebrewFile, maxJSONBytes, &f.Homebrew}, {sourcesFile, maxJSONBytes, &f.Sources},
		{entriesFile, maxEntriesBytes, &f.Entries},
	} {
		if err := f.readJSON(p.name, p.limit, p.v, true); err != nil {
			return nil, err
		}
	}
	return f, f.check()
}

// readStarter reads the rest of a starter: the compendium parts, and nothing else.
func (f *TakeHomeFile) readStarter() (*TakeHomeFile, error) {
	m := f.Manifest
	if m.Campaign.ID == 0 || m.ServerID == "" || m.Token != "" || m.CharacterID != 0 {
		return nil, errors.New(
			"a campaign starter file must name its table and campaign, and no character",
		)
	}
	for _, zf := range f.zr.File {
		if zf.Name != manifestFile && zf.Name != sourcesFile && zf.Name != entriesFile {
			return nil, fmt.Errorf("a campaign starter file must not hold %s", zf.Name)
		}
	}
	if err := f.readJSON(sourcesFile, maxJSONBytes, &f.Sources, true); err != nil {
		return nil, err
	}
	if err := f.readJSON(entriesFile, maxEntriesBytes, &f.Entries, true); err != nil {
		return nil, err
	}
	return f, f.check()
}

// ManifestKind reads only the manifest's kind, so a caller can pick the kind to read.
func ManifestKind(r io.ReaderAt, size int64) (Kind, error) {
	zr, err := zip.NewReader(r, size)
	if err != nil {
		return "", fmt.Errorf("it is not a zip file (%w)", err)
	}
	f := &TakeHomeFile{zr: zr}
	var head struct {
		Kind Kind `json:"kind"`
	}
	if err := f.readJSON(manifestFile, maxJSONBytes, &head, false); err != nil {
		return "", err
	}
	return head.Kind, nil
}

// check refuses values the database would refuse and files the zip lacks.
func (f *TakeHomeFile) check() error {
	kinds := map[enum.CompendiumKind]bool{}
	for _, k := range []enum.CompendiumKind{enum.KindSpell, enum.KindItem, enum.KindClass, enum.KindSubclass,
		enum.KindSpecies, enum.KindBackground, enum.KindFeat, enum.KindCondition, enum.KindRule} {
		kinds[k] = true
	}
	for _, e := range append(append([]TakeHomeEntry{}, f.Homebrew...), f.Entries...) {
		if !kinds[e.Kind] || e.Slug == "" || e.Name == "" || e.SourceKey == "" ||
			!json.Valid(e.Data) {
			return fmt.Errorf("compendium entry %q is not valid", e.Slug)
		}
	}
	for _, l := range f.Looks {
		if l.File == "" && l.URL == "" {
			return fmt.Errorf("look %q has no picture", l.Name)
		}
		if err := f.needFile(l.File); err != nil {
			return err
		}
	}
	for _, n := range f.Notes {
		if n.Kind != enum.NoteGameSession && n.Kind != enum.NoteGeneric || n.Title == "" {
			return fmt.Errorf("note %d is not valid", n.ID)
		}
		for _, b := range n.Blocks {
			if b.Visibility != enum.VisGM && b.Visibility != enum.VisPlayers &&
				b.Visibility != enum.VisListed &&
				b.Visibility != enum.VisOwner ||
				!json.Valid(b.Doc) {
				return fmt.Errorf("block %d is not valid", b.ID)
			}
			for _, bf := range b.Files {
				if err := f.needFile(bf.File); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (f *TakeHomeFile) needFile(name string) error {
	if name == "" {
		return nil
	}
	if !strings.HasPrefix(name, filesPrefix) || path.Clean(name) != name || f.find(name) == nil {
		return fmt.Errorf("the file lacks %q, which it refers to", name)
	}
	return nil
}

func (f *TakeHomeFile) find(name string) *zip.File {
	for _, zf := range f.zr.File {
		if zf.Name == name {
			return zf
		}
	}
	return nil
}

// OpenFile reads one attached file named by the JSON, up to MaxFileBytes.
func (f *TakeHomeFile) OpenFile(name string) ([]byte, error) {
	if name == "" {
		return nil, errors.New("no file named")
	}
	if err := f.needFile(name); err != nil {
		return nil, err
	}
	return readLimited(f.find(name), MaxFileBytes)
}

func (f *TakeHomeFile) readJSON(name string, limit int64, v any, strict bool) error {
	zf := f.find(name)
	if zf == nil {
		return fmt.Errorf("the file lacks %s", name)
	}
	body, err := readLimited(zf, limit)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	if strict {
		dec.DisallowUnknownFields()
	}
	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	return nil
}

// readLimited reads a whole zip member, failing past limit however big its
// header claims it is.
func readLimited(zf *zip.File, limit int64) ([]byte, error) {
	if zf.UncompressedSize64 > uint64(limit) { //nolint:gosec // limit is a positive constant
		return nil, fmt.Errorf("%s is too large", zf.Name)
	}
	rc, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer func() { _ = rc.Close() }()
	body, err := io.ReadAll(io.LimitReader(rc, limit+1))
	if err != nil {
		return nil, fmt.Errorf("%s: %w", zf.Name, err)
	}
	if int64(len(body)) > limit {
		return nil, fmt.Errorf("%s is too large", zf.Name)
	}
	return body, nil
}

// HomeManifest is what the player app keeps of an imported file: the manifest
// plus the home id of everything created from it, by the table's id, so a
// bring-back can name what changed.
type HomeManifest struct {
	TakeHomeManifest
	Maps HomeIDMaps `json:"maps"`
}

// HomeIDMaps map the id of a home row (the key) to the table's id it was made
// from (the value).
type HomeIDMaps struct {
	LevelUps    map[int64]int64 `json:"level_ups"`
	Images      map[int64]int64 `json:"images"`
	Attachments map[int64]int64 `json:"attachments"`
	Notes       map[int64]int64 `json:"notes"`
	Blocks      map[int64]int64 `json:"blocks"`
	Homebrew    map[int64]int64 `json:"homebrew"`
}

// ZipFileName is the download's name: "<name>-<kind>-<date>.zip".
func ZipFileName(kind Kind, characterName string, at time.Time) string {
	return fmt.Sprintf("%s-%s-%s.zip", fileSlug(characterName), kind, at.Format("2006-01-02"))
}
