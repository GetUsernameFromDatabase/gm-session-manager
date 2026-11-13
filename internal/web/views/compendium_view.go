package views

import (
	"context"
	"fmt"
	"net/url"
	"slices"
	"strings"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	selectcomp "gm-session-manager/internal/web/components/select"

	"github.com/a-h/templ"
)

// CompendiumView is the compendium search page.
type CompendiumView struct {
	Query   string
	Kind    string
	Results []sqlc.SearchCompendiumRow
	// TextResults are search hits in the text only, listed below Results.
	TextResults []sqlc.SearchCompendiumRow
	Counts      []sqlc.CountCompendiumRow
	// Habitat filters monsters (kind "monster" only); comma-separated.
	Habitat       string
	HabitatCounts []sqlc.CountHabitatsRow
	IsGM          bool
	// Homebrew is every homebrew entry, with the campaigns each player
	// entry is linked to and the campaigns it could be, for the GM.
	Homebrew      []sqlc.ListHomebrewRow
	HomebrewLinks map[int64][]sqlc.ListHomebrewLinksRow
	Campaigns     []sqlc.Campaign
	// MyHomebrew is a player's own entries.
	MyHomebrew []sqlc.ListHomebrewByOwnerRow
	Import     ImportNotice
	// Sources are the imported sources in priority order (GM), with their
	// overlap counts by key.
	Sources      []sqlc.CompendiumSource
	SharedCounts map[string]sqlc.CountSharedBySourceRow
	// Shared browses one source's entries another source also holds (GM),
	// SharedWith those other sources, With narrows to one of them; rows show
	// their CopyOrder (by CopyKey) and open the comparison.
	Shared     string
	SharedWith []sqlc.ListSharedWithRow
	With       string
	CopyOrder  map[string][]string
	More       string // URL of the next batch when browsing; "" at the end
	// Tab is the page tab to open (CompendiumTab*); "" is Browse.
	Tab string
	// Campaign is the campaign in the page's path (a player's is always
	// their own); 0 for the GM's whole compendium.
	Campaign int64
}

// Page tabs of the compendium page.
const (
	CompendiumTabBrowse   = "browse"
	CompendiumTabHomebrew = "homebrew"
	CompendiumTabImport   = "import"
)

// openTab is the page tab to show first: an import message lands on Import.
func (v CompendiumView) openTab() string {
	switch {
	case v.IsGM && (v.Import.Tab != "" || v.Tab == CompendiumTabImport):
		return CompendiumTabImport
	case v.Tab == CompendiumTabHomebrew:
		return CompendiumTabHomebrew
	}
	return CompendiumTabBrowse
}

// Import tabs of the compendium page.
const (
	ImportTabFolder = "folder"
	ImportTabPage   = "page"
)

// ImportNotice is the Import section's state: the open tab and the last
// import's message, shown under that tab.
type ImportNotice struct {
	Tab    string
	Msg    string
	Failed bool
	// Page is the page link the GM pasted, kept after an error.
	Page string
	// Back is the compendium filter (q, kind) to return to after a page import.
	Back string
	// Dir is DATA_DIR/compendium; OnHost is set when the browser runs on
	// the host, which may open it in the file manager.
	Dir    string
	OnHost bool
	// Job is the background 5etools import.
	Job FiveToolsJob
}

// bookIDs is the books' ids, upper case as 5etools writes them.
func bookIDs(books []compendium.Book) string {
	ids := make([]string, len(books))
	for i, b := range books {
		ids[i] = strings.ToUpper(b.ID)
	}
	return strings.Join(ids, ", ")
}

// CompendiumResultsID is the element htmx swaps when searching or changing kind.
const CompendiumResultsID = "compendium-results"

// rowFigure is a result row's right-hand text and its colour: a monster's
// CR, a spell's level, else the category; an item's is tinted by rarity, a
// subclass's (its class) by the class colour.
func rowFigure(r sqlc.SearchCompendiumRow) (text, color string) {
	switch r.Kind {
	case enum.KindMonster:
		if r.Figure == "" {
			return "", ""
		}
		return "CR " + crText(r.Figure), ""
	case enum.KindSpell:
		if r.Figure == "0" {
			return "Cantrip", ""
		}
		return "Level " + r.Figure, ""
	case enum.KindSubclass:
		return r.Classes, textOnly(classColors[r.Classes])
	case enum.KindItem:
		return strings.ReplaceAll(r.Category, "-", " "), rarityColors[strings.ToLower(r.Figure)]
	}
	return strings.ReplaceAll(r.Category, "-", " "), ""
}

// rarityColors tint an item's category; common items stay grey.
var rarityColors = map[string]string{
	"uncommon":  "text-green-600 dark:text-green-400",
	"rare":      "text-sky-600 dark:text-sky-400",
	"very rare": "text-violet-600 dark:text-violet-400",
	"legendary": "text-amber-600 dark:text-amber-400",
	"artifact":  "text-rose-600 dark:text-rose-400",
}

// textOnly keeps the text colours of a badge's classes, dropping its tint.
func textOnly(classes string) string {
	var out []string
	for _, c := range strings.Fields(classes) {
		if strings.HasPrefix(c, "text-") || strings.HasPrefix(c, "dark:text-") {
			out = append(out, c)
		}
	}
	return strings.Join(out, " ")
}

var crFractions = strings.NewReplacer("1/8", "⅛", "1/4", "¼", "1/2", "½")

// crText sets fractional challenge ratings as the book does.
func crText(cr string) string { return crFractions.Replace(cr) }

// chipRow wraps filter chips on desktop; on phones they make one line that
// swipes sideways, keeping the sticky filters short.
const chipRow = "flex flex-wrap gap-1.5 max-md:flex-nowrap max-md:overflow-x-auto max-md:pb-1 max-md:[&>*]:shrink-0"

func habitatCountMap(rows []sqlc.CountHabitatsRow) map[string]int64 {
	m := map[string]int64{}
	for _, r := range rows {
		m[r.Habitat] = r.N
	}
	return m
}

// CompendiumBase is the compendium as campaign sees it (its sources and
// homebrew), or the whole compendium for 0. The campaign is in the path, so
// an address shows the same view to whoever opens it.
func CompendiumBase(campaign int64) string {
	if campaign == 0 {
		return "/compendium"
	}
	return fmt.Sprintf("/campaigns/%d/compendium", campaign)
}

// CompendiumURL is the compendium list of campaign with query q.
func CompendiumURL(q url.Values, campaign int64) string {
	return CompendiumBase(campaign) + "?" + q.Encode()
}

type campaignCtxKey struct{}

// WithCampaign marks ctx as rendering inside campaign, so compendium links
// stay in it.
func WithCampaign(ctx context.Context, campaign int64) context.Context {
	return context.WithValue(ctx, campaignCtxKey{}, campaign)
}

// CompendiumPath is rest (e.g. "/entries/7") under the compendium of the
// campaign ctx renders in.
func CompendiumPath(ctx context.Context, rest string) string {
	id, _ := ctx.Value(campaignCtxKey{}).(int64)
	return CompendiumBase(id) + rest
}

// entryHref is an entry's page in the campaign ctx renders in.
func entryHref(ctx context.Context, id int64) string {
	return CompendiumPath(ctx, fmt.Sprintf("/entries/%d", id))
}

// kindBase is what the kind chips keep: the search, or the shared list and
// its other source. New on each call, since kindHref adds to it.
func (v CompendiumView) kindBase() url.Values {
	q := url.Values{}
	if v.Query != "" {
		q.Set(QueryKey, v.Query)
	}
	if v.Shared != "" {
		q.Set(SharedKey, v.Shared)
	}
	if v.With != "" {
		q.Set(SharedWithKey, v.With)
	}
	return q
}

// CompendiumRest is path without its compendium base ("" for the list), so
// the GM's switcher can open the same page in another campaign.
func CompendiumRest(path string) string {
	if _, rest, ok := strings.Cut(path, "/compendium"); ok {
		return rest
	}
	return ""
}

// habitatHref toggles habitat in the monster filter, in Habitats order.
func habitatHref(query, habitat, current string, campaign int64) string {
	q := url.Values{KindKey: {string(enum.KindMonster)}}
	if query != "" {
		q.Set(QueryKey, query)
	}
	set := kindSet(current)
	var next []string
	for _, h := range compendium.Habitats {
		if (h == habitat) != slices.Contains(set, h) {
			next = append(next, h)
		}
	}
	if len(next) > 0 {
		q.Set(HabitatKey, strings.Join(next, ","))
	}
	return CompendiumURL(q, campaign)
}

// kindHref toggles kind in the filter, on top of q (a fresh copy, see kindBase).
func kindHref(q url.Values, kind enum.CompendiumKind, current string, campaign int64) string {
	set := kindSet(current)
	var next []string
	for _, k := range compendium.Kinds {
		// Kinds order keeps URLs stable however the badges were clicked.
		if (k == kind) != slices.Contains(set, string(k)) {
			next = append(next, string(k))
		}
	}
	if len(next) > 0 {
		q.Set(KindKey, strings.Join(next, ","))
	}
	return CompendiumURL(q, campaign)
}

// kindSet splits the comma-separated kind filter; "" is every kind.
func kindSet(kind string) []string {
	return strings.FieldsFunc(kind, func(r rune) bool { return r == ',' })
}

// classColors are literal so Tailwind finds them; other classes stay grey.
var classColors = map[string]string{
	"barbarian": tintRed,
	"bard":      tintPink,
	"cleric":    tintYellow,
	"druid":     tintGreen,
	"fighter":   tintStone,
	"monk":      tintCyan,
	"paladin":   tintAmber,
	"ranger":    tintEmerald,
	"rogue":     tintSlate,
	"sorcerer":  tintOrange,
	"warlock":   tintPurple,
	"wizard":    tintBlue,
}

// sizeColors run cool to hot from tiny to gargantuan.
var sizeColors = map[string]string{
	"tiny":       tintSky,
	"small":      tintTeal,
	"medium":     tintGreen,
	"large":      tintAmber,
	"huge":       tintOrange,
	"gargantuan": tintRed,
}

// The tint badges, literal so Tailwind finds them.
const (
	tintRed     = "bg-red-500/15 text-red-700 dark:text-red-300"
	tintPink    = "bg-pink-500/15 text-pink-700 dark:text-pink-300"
	tintYellow  = "bg-yellow-500/15 text-yellow-700 dark:text-yellow-300"
	tintGreen   = "bg-green-500/15 text-green-700 dark:text-green-300"
	tintStone   = "bg-stone-500/15 text-stone-700 dark:text-stone-300"
	tintCyan    = "bg-cyan-500/15 text-cyan-700 dark:text-cyan-300"
	tintAmber   = "bg-amber-500/15 text-amber-700 dark:text-amber-300"
	tintEmerald = "bg-emerald-500/15 text-emerald-700 dark:text-emerald-300"
	tintSlate   = "bg-slate-500/15 text-slate-700 dark:text-slate-300"
	tintOrange  = "bg-orange-500/15 text-orange-700 dark:text-orange-300"
	tintPurple  = "bg-purple-500/15 text-purple-700 dark:text-purple-300"
	tintBlue    = "bg-blue-500/15 text-blue-700 dark:text-blue-300"
	tintSky     = "bg-sky-500/15 text-sky-700 dark:text-sky-300"
	tintTeal    = "bg-teal-500/15 text-teal-700 dark:text-teal-300"
)

// alignmentColor codes the good/evil axis; unaligned and "any" stay grey.
func alignmentColor(a string) string {
	a = strings.ToLower(a)
	switch {
	case strings.Contains(a, "good"):
		return tintBlue
	case strings.Contains(a, "evil"):
		return tintRed
	case strings.Contains(a, "neutral"):
		return tintAmber
	}
	return ""
}

// kindColors tint each kind's icon; literal so Tailwind finds them.
var kindColors = map[enum.CompendiumKind]string{
	enum.KindMonster:    "text-rose-600 dark:text-rose-400",
	enum.KindSpell:      "text-violet-600 dark:text-violet-400",
	enum.KindItem:       "text-amber-600 dark:text-amber-400",
	enum.KindClass:      "text-blue-600 dark:text-blue-400",
	enum.KindSubclass:   "text-blue-600 dark:text-blue-400",
	enum.KindSpecies:    "text-emerald-600 dark:text-emerald-400",
	enum.KindBackground: "text-stone-600 dark:text-stone-400",
	enum.KindFeat:       "text-yellow-600 dark:text-yellow-400",
	enum.KindCondition:  "text-orange-600 dark:text-orange-400",
	enum.KindRule:       "text-slate-600 dark:text-slate-400",
}

// PickerTarget says what the result buttons of a compendium picker do.
type PickerTarget struct {
	EncounterID int64
	CharacterID int64
	EntityID    int64
}

// pickerLinkAttrs opens entries in a new tab from pickers embedded in
// other pages; the compendium's own dropdown navigates in place.
func pickerLinkAttrs(t PickerTarget) templ.Attributes {
	if t == (PickerTarget{}) {
		return nil
	}
	return templ.Attributes{"target": targetBlank}
}

// EntryView is one compendium entry page.
type EntryView struct {
	// CanEdit is set for homebrew the viewer may change.
	CanEdit bool
	Entry   sqlc.CompendiumEntry
	Source  sqlc.CompendiumSource
	Data    any
	IsGM    bool
	// Habitats are a monster's, seeded from its source and edited by the GM.
	Habitats []string
	// HabitatCopies are the copies the GM may take habitats from while the
	// monster has none.
	HabitatCopies []HabitatCopy
	// Copies are the entry's other sources the viewer may see.
	Copies []sqlc.ListEntryCopiesRow
	// Back links to the compendium list the entry was opened from, filters
	// kept, or that a page import started from.
	Back string
}

// HabitatCopy is another source's copy of a monster and its habitats.
type HabitatCopy struct {
	Copy     sqlc.ListEntryCopiesRow
	Habitats []string
}

// inHabitatOrder joins habitats in compendium.Habitats order; every one of
// them reads "Any", as in the Monster Manual.
func inHabitatOrder(habitats []string) string {
	var out []string
	for _, h := range compendium.Habitats {
		if slices.Contains(habitats, h) {
			out = append(out, h)
		}
	}
	if len(out) == len(compendium.Habitats) {
		return "Any"
	}
	return strings.Join(out, ", ")
}

func signedMap(m map[string]int, order []string) string {
	var parts []string
	for _, k := range order {
		if v, ok := m[k]; ok {
			parts = append(parts, k+" "+dnd5e.Signed(v))
		}
	}
	return strings.Join(parts, ", ")
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func speeds(m map[string]int) string {
	var parts []string
	for _, k := range sortedKeys(m) {
		parts = append(parts, fmt.Sprintf("%s %d ft.", k, m[k]))
	}
	return strings.Join(parts, ", ")
}

const factsGrid = "grid grid-cols-2 gap-x-4 gap-y-1.5"

func parenthesized(s string) string {
	if s == "" {
		return ""
	}
	return "(" + s + ")"
}

func commaPrefixed(s string) string {
	if s == "" {
		return ""
	}
	return ", " + s
}

// compendiumFindURL opens the entry of kind and slug the viewer would see.
func compendiumFindURL(ctx context.Context, kind enum.CompendiumKind, slug string) string {
	return CompendiumPath(ctx, fmt.Sprintf("/find/%s/%s", kind, url.PathEscape(slug)))
}

func levelsText(levels []int) string {
	parts := make([]string, len(levels))
	for i, l := range levels {
		parts[i] = fmt.Sprint(l)
	}
	return strings.Join(parts, ", ")
}

// HomebrewView is the homebrew editor.
type HomebrewView struct {
	ID    int64
	Input compendium.HomebrewInput
	Error string
	IsGM  bool
	// AddToSheet adds the new entry to the player's character on save.
	AddToSheet bool
	// Page is the site page under review (its kind is fixed; ID "" when
	// none); Replaces is the entry an earlier import of it made.
	Page     compendium.Page
	Replaces int64
	// Back is the compendium filter the import started from.
	Back string
}

// masteryItems offer no mastery, then each weapon mastery property, and
// current when it is none of them ("Nick, Vex"), so saving keeps it.
func masteryItems(current string) []selectcomp.ItemData {
	names := compendium.WeaponMasteries
	if current != "" && !slices.Contains(names, current) {
		names = append(slices.Clone(names), current)
	}
	return append([]selectcomp.ItemData{{Value: "", Label: labelNone}}, selectItems(names, func(m string) string { return m }, func(m string) string { return m })...)
}

const homebrewFormID = "homebrew-form"

func homebrewAction(id int64) string {
	if id == 0 {
		return "/homebrew"
	}
	return fmt.Sprintf("/entries/%d", id)
}
