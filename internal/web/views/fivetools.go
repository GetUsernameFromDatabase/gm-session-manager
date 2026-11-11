package views

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"

	"github.com/a-h/templ"
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
	// SourceOrderPath takes every source key (SourceKeyField) in the new order.
	SourceOrderPath = "/compendium/sources/order"
	SourceKeyField  = "key"
)

// scrollBox caps a long list or notice so the controls around it stay close.
const scrollBox = "max-h-80 overflow-y-auto"

// FiveToolsView is the book picker: books by group, with the import date of
// each downloaded one. Running disables it while an import runs.
type FiveToolsView struct {
	Groups  []FiveToolsGroup
	Msg     string
	Failed  bool
	Running bool
}

// FiveToolsJob is the GM's background 5etools import: its progress while
// Running, then the summary of the last one (Msg).
type FiveToolsJob struct {
	Running bool
	// Importing is past the downloads; Done of Total books, or sources.
	Importing   bool
	Done, Total int
	Current     string
	Msg         string
	Failed      bool
}

func (j FiveToolsJob) label() string {
	if j.Importing {
		return fmt.Sprintf("Importing %d of %d: %s", j.Done+1, j.Total, j.Current)
	}
	return fmt.Sprintf("Downloading %d of %d: %s", j.Done+1, j.Total, j.Current)
}

func (j FiveToolsJob) bar() string {
	return fmt.Sprintf("width: %d%%", j.Done*100/max(j.Total, 1))
}

// Hot adds the running import's label and bar.
func (j FiveToolsJob) Hot(st hub.State) {
	if j.Running {
		st["imp.label"], st["imp.bar"] = j.label(), j.bar()
	}
}

// Shape is the job with its hot values fixed.
func (j FiveToolsJob) Shape() FiveToolsJob {
	j.Importing, j.Done, j.Total, j.Current = false, 0, 0, ""
	return j
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

// sourceURL is a source's priority action (move, delete) in the campaign ctx
// renders in, whose overlap counts the answer shows.
func sourceURL(ctx context.Context, key, action string) string {
	return CompendiumPath(ctx, "/sources/"+url.PathEscape(key)+"/"+action)
}

// SharedKey is the browse filter to one source's entries another source
// also holds (CompendiumView.Shared); SharedWithKey narrows it to one other
// source (CompendiumView.With).
const (
	SharedKey     = "shared"
	SharedWithKey = "with"
)

// sharedHref browses source's shared entries in the campaign ctx renders in.
func sharedHref(ctx context.Context, source string) templ.SafeURL {
	return templ.SafeURL(CompendiumPath(ctx, "?"+url.Values{SharedKey: {source}}.Encode()))
}

// sharedWithHref narrows the shared list to entries with also holds ("" for
// any); the kind filter stays.
func sharedWithHref(v CompendiumView, with string) templ.SafeURL {
	q := url.Values{SharedKey: {v.Shared}}
	if with != "" {
		q.Set(SharedWithKey, with)
	}
	if v.Kind != "" {
		q.Set("kind", v.Kind)
	}
	return templ.SafeURL(CompendiumURL(q, v.Campaign))
}

// CopyKey keys CompendiumView.CopyOrder by an entry's kind and slug.
func CopyKey(kind enum.CompendiumKind, slug string) string { return string(kind) + "/" + slug }

// compareHref sets every copy of a kind and slug side by side.
func compareHref(ctx context.Context, kind enum.CompendiumKind, slug string) string {
	return CompendiumPath(ctx, "/compare/"+url.PathEscape(string(kind))+"/"+url.PathEscape(slug))
}

// CompareView is every copy of one entry the campaign sees, best first.
type CompareView struct {
	Copies []CompareCopy
}

// CompareCopy is one source's copy, decoded for its stat block.
type CompareCopy struct {
	Entry  sqlc.CompendiumEntry
	Source sqlc.ListEntryCopiesRow
	Data   any
}
