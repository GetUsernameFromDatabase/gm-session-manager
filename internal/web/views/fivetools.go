package views

import (
	"net/url"
	"strings"
	"time"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
)

// ImportTabFiveTools is the import tab that picks 5etools books.
const ImportTabFiveTools = "fivetools"

// FiveToolsPath is the picker's endpoint.
const FiveToolsPath = "/compendium/5etools"

// Element ids and form keys of the 5etools picker and the priority card.
const (
	FiveToolsID    = "fivetools-books"
	FiveToolsField = "book"
	SourcesID      = "source-priority"
	SourceDirField = "dir"
	SourceUp       = "up"
	SourceDown     = "down"
)

// FiveToolsView is the book picker: books by group, with the import date of
// each downloaded one.
type FiveToolsView struct {
	Groups []FiveToolsGroup
	Msg    string
	Failed bool
}

// FiveToolsGroup is the books of one compendium.Group*.
type FiveToolsGroup struct {
	Title string
	Books []FiveToolsBook
}

// FiveToolsBook is a book and, when downloaded, when it was last imported.
type FiveToolsBook struct {
	compendium.Book
	Downloaded bool
	Imported   time.Time // zero until the first import
}

var fiveToolsGroupTitles = map[string]string{
	compendium.GroupCore:       "Core rules",
	compendium.GroupSupplement: "Supplements",
	compendium.GroupAdventure:  "Adventures",
	compendium.GroupUA:         "Unearthed Arcana",
	compendium.GroupThirdParty: "Third party",
}

// NewFiveToolsView groups books in display order. downloaded maps a book id
// (lowercase) to its source's last import (zero if not imported yet).
func NewFiveToolsView(books []compendium.Book, downloaded map[string]time.Time) FiveToolsView {
	var v FiveToolsView
	for _, g := range []string{compendium.GroupCore, compendium.GroupSupplement, compendium.GroupAdventure, compendium.GroupUA, compendium.GroupThirdParty} {
		var list []FiveToolsBook
		for _, b := range books {
			if b.Group == g {
				at, ok := downloaded[strings.ToLower(b.ID)]
				list = append(list, FiveToolsBook{Book: b, Downloaded: ok, Imported: at})
			}
		}
		if len(list) > 0 {
			v.Groups = append(v.Groups, FiveToolsGroup{Title: fiveToolsGroupTitles[g], Books: list})
		}
	}
	return v
}

// copyLabel names an entry's copy in a source for the "Also in" line.
func copyLabel(c sqlc.ListEntryCopiesRow) string { return SourceLabel(c.SourceKey, c.Edition) }

// sourceMoveURL is where a source's priority arrows post.
func sourceMoveURL(key string) string { return "/compendium/sources/" + url.PathEscape(key) + "/move" }
