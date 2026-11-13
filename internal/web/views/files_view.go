package views

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

// PathFiles is the admin page of every attachment and where it is used.
const PathFiles = "/files"

// FilesBase is the Files page of campaign (0: every campaign).
func FilesBase(campaign int64) string {
	if campaign == 0 {
		return PathFiles
	}
	return fmt.Sprintf("/campaigns/%d/files", campaign)
}

// Tabs of the Files page.
const (
	FilesTabAll    = "all"
	FilesTabUnused = "unused"
)

// FilesView is the admin's list of attachments, newest first: the layout's
// campaign's, or every campaign's.
type FilesView struct {
	Files []FileRow
	// Tab is the tab to open (FilesTab*); "" is All.
	Tab string
}

// FileRow is one attachment. A private one (from a player's own block)
// shows neither its name nor its picture.
type FileRow struct {
	ID         int64
	Name       string
	CampaignID int64
	Campaign   string
	Size       int64
	CreatedAt  time.Time
	IsImage    bool
	Private    bool
	Uses       []FileUse
}

// Kinds of FileUse.
const (
	FileUseNote  = "note"
	FileUseImage = "image"
)

// FileUse is a note or image that uses a file. A private use is a player's
// own block, shown without its note's title or a link.
type FileUse struct {
	Kind    string
	URL     string
	Label   string
	Private bool
}

func unusedFiles(files []FileRow) []FileRow {
	var idle []FileRow
	for _, f := range files {
		if len(f.Uses) == 0 {
			idle = append(idle, f)
		}
	}
	return idle
}

// filesByCampaign groups the files by campaign, in the order each first appears.
func filesByCampaign(files []FileRow) [][]FileRow {
	var groups [][]FileRow
	at := map[int64]int{}
	for _, f := range files {
		i, ok := at[f.CampaignID]
		if !ok {
			i = len(groups)
			at[f.CampaignID] = i
			groups = append(groups, nil)
		}
		groups[i] = append(groups[i], f)
	}
	return groups
}

// fileExt is a file name's extension in capitals, as "PDF".
func fileExt(name string) string {
	return strings.ToUpper(strings.TrimPrefix(filepath.Ext(name), "."))
}

func totalSize(files []FileRow) int64 {
	var n int64
	for _, f := range files {
		n += f.Size
	}
	return n
}

// fileSize is a byte count in KB or MB.
func fileSize(n int64) string {
	if n < 1<<20 {
		return fmt.Sprintf("%.0f KB", float64(n)/(1<<10))
	}
	return fmt.Sprintf("%.1f MB", float64(n)/(1<<20))
}
