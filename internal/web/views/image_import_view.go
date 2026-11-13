package views

import (
	"fmt"
	"slices"
	"strings"

	"gm-session-manager/internal/db/sqlc"
	selectcomp "gm-session-manager/internal/web/components/select"
)

// Form fields of the import review.
const (
	ImportColName        = "col_name"
	ImportColLink        = "col_link"
	ImportColCategory    = "col_category"
	ImportColDescription = "col_description"
	ImportCatSource      = "cat_source"
	ImportCatTarget      = "cat_target"
	// ImportRowsField holds the ticked rows, encoded by app.js (see
	// parseRowSet). Its presence marks a reviewed form.
	ImportRowsField      = "rows"
	ImportFilterQ        = "filter_q"
	ImportFilterCategory = "filter_category"
	ImportFilterStatus   = "filter_status"
	ImportKeepCol        = "keep_col"
	// ImportKeepVal values carry ImportKeepPrefix, so an empty CSV value is
	// still a real checkbox value.
	ImportKeepVal    = "keep_val"
	ImportKeepPrefix = "="
)

// Values of the rows' category filter besides a category name.
const (
	importFilterAll  = "*"
	importFilterNone = "-"
)

// ImportMapping is the CSV column of each image field (-1 for none).
type ImportMapping struct {
	Name, Link, Category, Description int
}

// ImportRow is one CSV row as it would be imported.
type ImportRow struct {
	Index       int
	Name        string
	Link        string
	Host        string
	Category    string
	Description string
	// Target is the category the row goes into ("" for none).
	Target string
	Ticked bool
	// Problem says why the row cannot be imported ("" when it can).
	Problem string
}

// ImportCategory is one category value of the CSV and the category it
// becomes ("" for none). Values with the same target are merged.
type ImportCategory struct {
	Source string
	Count  int
	Target string
}

// ImportColumn describes one CSV column for the overview.
type ImportColumn struct {
	Header   string
	Distinct int
	Samples  []string
}

// ImportValue is one value of the keep-filter column.
type ImportValue struct {
	Value string
	Count int
	Kept  bool
}

// ImportView is the review step of a CSV import.
type ImportView struct {
	CampaignID int64
	Token      string
	Header     []string
	Map        ImportMapping
	Rows       []ImportRow
	Cats       []ImportCategory
	Categories []sqlc.ImageCategory
	// The rows filter, kept across refreshes.
	FilterQ        string
	FilterCategory string
	FilterStatus   string
	// Total is the number of rows in the CSV; Rows holds the kept ones.
	Total   int
	Columns []ImportColumn
	// KeepCol is the column filtering the rows (-1 for none).
	KeepCol       int
	KeepValues    []ImportValue
	TooManyValues bool
}

func columnLabel(i int, header string) string {
	if header == "" {
		return fmt.Sprintf("Column %d", i+1)
	}
	return header
}

// roleOf names what column i is used for, if anything.
func (v ImportView) roleOf(i int) string {
	switch i {
	case v.Map.Name:
		return labelName
	case v.Map.Link:
		return "Link"
	case v.Map.Category:
		return "Category"
	case v.Map.Description:
		return "Description"
	}
	return ""
}

func (v ImportView) filtering() bool {
	return slices.ContainsFunc(v.KeepValues, func(x ImportValue) bool { return x.Kept })
}

func (v ImportView) url() string {
	return fmt.Sprintf("/campaigns/%d/images/import/%s", v.CampaignID, v.Token)
}

func (v ImportView) columnItems(optional bool) []selectcomp.ItemData {
	var items []selectcomp.ItemData
	if optional {
		items = append(items, selectcomp.ItemData{Value: "-1", Label: labelNone})
	}
	for i, h := range v.Header {
		items = append(items, selectcomp.ItemData{Value: fmt.Sprint(i), Label: columnLabel(i, h)})
	}
	return items
}

// filterItems filter the rows by the categories they will go into.
func (v ImportView) filterItems() []selectcomp.ItemData {
	items := []selectcomp.ItemData{{Value: importFilterAll, Label: "All categories"}, {Value: importFilterNone, Label: "No category"}}
	var seen []string
	for _, c := range v.Cats {
		if c.Target != "" && !slices.Contains(seen, c.Target) {
			seen = append(seen, c.Target)
		}
	}
	slices.Sort(seen)
	return append(items, selectItems(seen, func(s string) string { return s }, func(s string) string { return s })...)
}

// catGroups groups the category values by target, so values merged into
// one category sit together. Groups sort by their category name (a value
// with no category by its own name), values within a group by name.
func (v ImportView) catGroups() [][]ImportCategory {
	byTarget := map[string][]ImportCategory{}
	var groups [][]ImportCategory
	for _, c := range v.Cats {
		if c.Target == "" {
			groups = append(groups, []ImportCategory{c})
			continue
		}
		byTarget[c.Target] = append(byTarget[c.Target], c)
	}
	for _, g := range byTarget {
		slices.SortFunc(g, func(a, b ImportCategory) int {
			return strings.Compare(strings.ToLower(a.Source), strings.ToLower(b.Source))
		})
		groups = append(groups, g)
	}
	sortKey := func(g []ImportCategory) string {
		if g[0].Target == "" {
			return strings.ToLower(g[0].Source)
		}
		return strings.ToLower(g[0].Target)
	}
	slices.SortFunc(groups, func(a, b []ImportCategory) int {
		if c := strings.Compare(sortKey(a), sortKey(b)); c != 0 {
			return c
		}
		return strings.Compare(a[0].Source, b[0].Source)
	})
	return groups
}

// importStatusItems filter rows by tick and by link-check result; app.js
// applies them.
func importStatusItems() []selectcomp.ItemData {
	return []selectcomp.ItemData{
		{Value: "any", Label: "Any status"},
		{Value: "ticked", Label: "Ticked"},
		{Value: "unticked", Label: "Unticked"},
		{Value: "broken", Label: "Didn't load"},
	}
}

// filterValue is the chosen filter, or all when it no longer exists.
func (v ImportView) filterValue() string {
	for _, it := range v.filterItems() {
		if it.Value == v.FilterCategory {
			return it.Value
		}
	}
	return importFilterAll
}

func rowFilterCategory(r ImportRow) string {
	if r.Target == "" {
		return importFilterNone
	}
	return r.Target
}
