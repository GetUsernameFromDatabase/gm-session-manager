package views

import (
	"cmp"
	"context"
	"fmt"
	"net/url"
	"slices"
	"strings"
	"time"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/utils"

	"github.com/a-h/templ"
)

// ImportTabFiveTools is the import tab that picks 5etools books.
const ImportTabFiveTools = "fivetools"

// ImportKey is the query (and form) key naming the import tab to open.
const ImportKey = "import"

// FiveToolsPath is the picker's endpoint.
const FiveToolsPath = route.Compendium + route.CompendiumFiveTools

// Element ids and form keys of the 5etools picker and the priority card.
const (
	FiveToolsID    = "fivetools-books"
	FiveToolsField = "book"
	SourcesID      = "source-priority"
	// SourceOrderPath takes every source key (SourceKeyField) in the new order.
	SourceOrderPath = route.Compendium + route.CompendiumSourceOrder
	SourceKeyField  = "key"
	// FiveToolsSetField names a starting set (FiveToolsSet2024) instead of books.
	FiveToolsSetField = "set"
	FiveToolsSet2024  = "2024"
	// Open5eDocField names the SRD an Open5e download fetches (compendium.SRD*Key).
	Open5eDocField = "doc"
	// DismissWasField is the summary the dismiss X was shown, so a stale tab
	// cannot clear a newer one.
	DismissWasField = "was"
	// StarterSetsID is the element id of StarterSets.
	StarterSetsID = "starter-sets"
)

// StarterLiveCompendium is the empty compendium's StarterSets' live id
// (LiveImport).
const StarterLiveCompendium = 1

// starterLive makes the empty compendium's StarterSets live; card is the
// Sources card's campaign (0 in the empty compendium), whose card is live
// itself.
func starterLive(card int64) templ.Attributes {
	if card > 0 {
		return nil
	}
	return live(LiveImport, StarterLiveCompendium)
}

// starterBase is the compendium the set forms post to: the card's campaign,
// so the import ticks the sets there, or the one ctx renders in.
func starterBase(ctx context.Context, card int64) string {
	if card > 0 {
		return CompendiumBase(card)
	}
	return CompendiumPath(ctx, "")
}

// scrollBox caps a long list or notice so the controls around it stay close.
const scrollBox = "max-h-80 overflow-y-auto"

// SourceGotoField names the priority card's "Go to source" box (app.js).
const SourceGotoField = "goto-source"

// sourcePositions numbers the priority list as the GM types it: homebrew
// is 0 wherever it sits, the others 1..n in order.
func sourcePositions(sources []sqlc.CompendiumSource) []int {
	out := make([]int, len(sources))
	n := 0
	for i, s := range sources {
		if s.SourceKey != compendium.HomebrewSource {
			n++
			out[i] = n
		}
	}
	return out
}

// emptySourceKeys are the sources holding no entries that the trash button
// would also offer to delete.
func emptySourceKeys(
	ctx context.Context,
	sources []sqlc.CompendiumSource,
	counts map[string]sqlc.CountEntriesBySourceRow,
) []string {
	var keys []string
	for _, s := range sources {
		if s.SourceKey != compendium.HomebrewSource && !fromTable(ctx, s.SourceKey) &&
			counts[s.SourceKey].Total == 0 {
			keys = append(keys, s.SourceKey)
		}
	}
	return keys
}

// entryCountText is a source's count chip: shown/total, or the total when
// lists show every entry.
func entryCountText(c sqlc.CountEntriesBySourceRow) string {
	if c.Visible < c.Total {
		return fmt.Sprintf("%d/%d", c.Visible, c.Total)
	}
	return fmt.Sprint(c.Total)
}

// entryCountTitle explains the chip; hasShared adds the overlap with other
// sources (s).
func entryCountTitle(
	c sqlc.CountEntriesBySourceRow,
	s sqlc.CountSharedBySourceRow,
	hasShared bool,
) string {
	t := fmt.Sprintf("%d %s", c.Total, utils.IfElse(c.Total == 1, "entry", "entries"))
	switch {
	case c.Total == 0:
		t += "."
	case c.Visible < c.Total:
		t += fmt.Sprintf(", %d shown.", c.Visible)
	default:
		t += ", all shown."
	}
	if hasShared {
		t += fmt.Sprintf(" %d shared with other sources, %d of those win.", s.Shared, s.Shown)
	}
	return t
}

// Shown/total ratios below which the count turns a warmer colour
// (entryCountClass).
const (
	countAmberBelow  = 0.5
	countOrangeBelow = 0.35
	countRedBelow    = 0.2
)

// Colours of the count (entryCountClass); half shown or more stays muted.
const (
	countAmber     = "text-amber-700 dark:text-amber-300"
	countOrange    = "text-orange-700 dark:text-orange-300"
	countRed       = "text-red-600 dark:text-red-400"
	countNoneShown = "text-destructive"
)

// entryCountClass colours the count by the share lists show: none from half
// up (a few entries losing to another source is normal), amber to red below.
func entryCountClass(c sqlc.CountEntriesBySourceRow) string {
	if c.Total == 0 {
		return ""
	}
	r := float64(c.Visible) / float64(c.Total)
	switch {
	case c.Visible == 0:
		return countNoneShown
	case r < countRedBelow:
		return countRed
	case r < countOrangeBelow:
		return countOrange
	case r < countAmberBelow:
		return countAmber
	}
	return ""
}

// sortableSources is true with two or more sources besides homebrew.
func sortableSources(sources []sqlc.CompendiumSource) bool {
	return len(
		slices.DeleteFunc(
			slices.Clone(sources),
			func(s sqlc.CompendiumSource) bool { return s.SourceKey == compendium.HomebrewSource },
		),
	) >= 2
}

// minLostPercent keeps a tiny loss (1 of 400) visible as a sliver of the pie.
const minLostPercent = 8

// lostPercent is the share of a source's entries lists do not show, as a pie
// fill; false when none are lost.
func lostPercent(c sqlc.CountEntriesBySourceRow) (int, bool) {
	if c.Total <= 0 || c.Visible >= c.Total {
		return 0, false
	}
	return max(minLostPercent, int(100*(c.Total-c.Visible)/c.Total)), true
}

// lostDotClass colours the pie: the count's tier colour, amber from the first lost entry.
func lostDotClass(c sqlc.CountEntriesBySourceRow) string {
	return cmp.Or(entryCountClass(c), countAmber)
}

// lostDotStyle sets the pie's fill (--lost, read by the dot's gradient).
func lostDotStyle(pct int) templ.SafeCSS {
	return templ.SafeCSS(fmt.Sprintf("--lost:%d%%;", pct))
}

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
	// Books are the picked 5etools book ids while Running and after a
	// failed end, so the reloaded picker keeps the GM's ticks.
	Books []string
}

func (j FiveToolsJob) label() string {
	switch {
	case j.Total == 0: // the Open5e listing is not in yet
		return "Starting…"
	case j.Importing:
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
	// Ticked pre-ticks a book a starting set holds (?set=).
	Ticked   bool
	Imported time.Time // zero until the first import
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

// SharedKey is the browse filter to one source's entries another source
// also holds (CompendiumView.Shared); SharedWithKey narrows it to one other
// source (CompendiumView.With); SharedOnlyKey set to SharedOnlyOn flips it
// to the entries no other source holds (CompendiumView.Only).
const (
	SharedKey     = "shared"
	SharedWithKey = "with"
	SharedOnlyKey = "only"
	SharedOnlyOn  = "1"
)

// Compendium list query keys: the search, kind and habitat filters, the open
// tab, the browse cursor (the after_* keys) and how many rows a reload keeps.
const (
	QueryKey     = "q"
	KindKey      = "kind"
	HabitatKey   = "habitat"
	TabKey       = "tab"
	AfterNameKey = "after_name"
	AfterIDKey   = "after_id"
	RowsKey      = "rows"
)

// sharedHref browses source's shared entries in the campaign ctx renders in.
func sharedHref(ctx context.Context, source string) templ.SafeURL {
	return templ.SafeURL(CompendiumPath(ctx, "?"+url.Values{SharedKey: {source}}.Encode()))
}

// sharedOnlyHref browses the entries only source holds in the campaign ctx
// renders in.
func sharedOnlyHref(ctx context.Context, source string) templ.SafeURL {
	return templ.SafeURL(
		CompendiumPath(
			ctx,
			"?"+url.Values{SharedKey: {source}, SharedOnlyKey: {SharedOnlyOn}}.Encode(),
		),
	)
}

// sharedWithHref narrows the shared list to entries with also holds ("" for
// any); the kind filter stays.
func sharedWithHref(v CompendiumView, with string) templ.SafeURL {
	q := url.Values{}
	if with != "" {
		q.Set(SharedWithKey, with)
	}
	return sharedListHref(v, q)
}

// sharedListHref is v's shared list with q's filters and the kind filter.
func sharedListHref(v CompendiumView, q url.Values) templ.SafeURL {
	q.Set(SharedKey, v.Shared)
	if v.Kind != "" {
		q.Set(KindKey, v.Kind)
	}
	return templ.SafeURL(CompendiumURL(q, v.Campaign))
}

// onlyChipHref lists the entries only v.Shared holds; the kind filter stays.
func onlyChipHref(v CompendiumView) templ.SafeURL {
	return sharedListHref(v, url.Values{SharedOnlyKey: {SharedOnlyOn}})
}

// onlyCounts are how many of v.Shared's entries no other source holds, and
// how many of those lists show, from the priority card's counts.
func (v CompendiumView) onlyCounts() (total, shown int64) {
	e, s := v.EntryCounts[v.Shared], v.SharedCounts[v.Shared]
	return e.Total - s.Shared, e.Visible - s.Shown
}

// CopyKey keys CompendiumView.CopyOrder by an entry's kind and slug.
func CopyKey(kind enum.CompendiumKind, slug string) string { return string(kind) + "/" + slug }

// compareHref sets every copy of a kind and slug side by side.
func compareHref(ctx context.Context, kind enum.CompendiumKind, slug string) string {
	return CompendiumPath(ctx, route.Path(route.CompendiumCompare, kind, slug))
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
