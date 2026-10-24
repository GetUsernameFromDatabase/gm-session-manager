package export

import (
	"archive/zip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
)

// Campaign writes a zip of one campaign to w:
//
//	notes/<id>-<title>.md       every block except players' 'owner' blocks
//	characters/<id>-<name>.json the character rows, sheet included
//	attachments/<file>          files attached to exported blocks
//
// Notes left with no exportable block (a player's private note) are
// skipped entirely, title included.
func Campaign(ctx context.Context, w io.Writer, q *sqlc.Queries, campaignID int64, attachmentsDir string) error {
	notes, err := q.ListExportNotes(ctx, campaignID)
	if err != nil {
		return err
	}
	blocks, err := q.ListExportBlocks(ctx, campaignID)
	if err != nil {
		return err
	}
	chars, err := q.ListCharactersByCampaign(ctx, campaignID)
	if err != nil {
		return err
	}
	files, err := q.ListAttachmentsByCampaign(ctx, campaignID)
	if err != nil {
		return err
	}

	byNote := map[int64][]sqlc.ListExportBlocksRow{}
	for _, b := range blocks {
		byNote[b.NoteID] = append(byNote[b.NoteID], b)
	}
	noteFile := map[int64]string{}
	for _, n := range notes {
		if len(byNote[n.ID]) > 0 || !n.OwnerCharacterID.Valid {
			noteFile[n.ID] = fmt.Sprintf("%d-%s.md", n.ID, fileSlug(n.Title))
		}
	}
	attachmentFile := map[string]string{}
	for _, a := range files {
		attachmentFile[fmt.Sprintf("/attachments/%d", a.ID)] = a.Path
	}
	links := Links{
		Mention: func(id string) string {
			kind, rawID, _ := strings.Cut(id, ":")
			n, _ := strconv.ParseInt(rawID, 10, 64)
			if kind == "note" {
				return noteFile[n]
			}
			return ""
		},
		Image: func(src string) string {
			if p, ok := attachmentFile[src]; ok {
				return "../attachments/" + p
			}
			return ""
		},
	}

	zw := zip.NewWriter(w)
	for _, n := range notes {
		name, ok := noteFile[n.ID]
		if !ok {
			continue
		}
		var b strings.Builder
		fmt.Fprintf(&b, "# %s\n\n", n.Title)
		if n.Kind == "game_session" {
			b.WriteString("_Game session_\n\n")
		}
		for _, blk := range byNote[n.ID] {
			md, err := Markdown(blk.DocJson, links)
			if err != nil {
				return err
			}
			fmt.Fprintf(&b, "<!-- block %d, visibility: %s -->\n\n%s\n", blk.ID, blk.Visibility, md)
		}
		if err := writeFile(zw, "notes/"+name, []byte(b.String())); err != nil {
			return err
		}
	}

	for _, c := range chars {
		body, err := NewCharacterFile(c).JSON()
		if err != nil {
			return err
		}
		if err := writeFile(zw, CharacterFileName(c), body); err != nil {
			return err
		}
	}

	for _, a := range files {
		body, err := os.ReadFile(filepath.Join(attachmentsDir, filepath.Base(a.Path))) //nolint:gosec // our random file name, kept inside attachmentsDir by Base
		if err != nil {
			return fmt.Errorf("attachment %d: %w", a.ID, err)
		}
		if err := writeFile(zw, "attachments/"+a.Path, body); err != nil {
			return err
		}
	}
	return zw.Close()
}

func writeFile(zw *zip.Writer, name string, body []byte) error {
	f, err := zw.Create(name)
	if err != nil {
		return err
	}
	_, err = f.Write(body)
	return err
}

func fileSlug(s string) string {
	if slug := compendium.Slugify(s); slug != "" {
		return slug
	}
	return "untitled"
}

// CharacterFile is a character as the campaign export and the GM's
// download write it, and as the GM's import reads it back.
type CharacterFile struct {
	ID     int64           `json:"id"`
	Name   string          `json:"name"   validate:"required,max=200"`
	System string          `json:"system" validate:"required"`
	HP     int64           `json:"hp"`
	MaxHP  int64           `json:"max_hp"`
	TempHP int64           `json:"temp_hp"`
	Sheet  json.RawMessage `json:"sheet"  validate:"required"`
}

// NewCharacterFile is c as a file.
func NewCharacterFile(c sqlc.Character) CharacterFile {
	return CharacterFile{c.ID, c.Name, c.System, c.Hp, c.MaxHp, c.TempHp, json.RawMessage(c.SheetJson)}
}

// JSON is the file's indented body.
func (f CharacterFile) JSON() ([]byte, error) {
	return json.MarshalIndent(f, "", "  ")
}

// CharacterFileName is the file's name: id and name, as in the export.
func CharacterFileName(c sqlc.Character) string {
	return fmt.Sprintf("characters/%d-%s.json", c.ID, fileSlug(c.Name))
}
