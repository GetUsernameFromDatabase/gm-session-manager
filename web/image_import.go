package web

import (
	"bytes"
	"crypto/rand"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"gm-session-manager/db/sqlc"
	"gm-session-manager/web/views"
)

// Limits for CSV imports held between upload and commit.
const (
	maxImportBytes  = 5 << 20
	maxImportRows   = 20000
	maxImports      = 10
	importLifetime  = time.Hour
	noImportColumn  = -1
	importFileField = "file"
)

// csvImport is an uploaded CSV waiting for the GM to map and pick rows.
type csvImport struct {
	campaignID int64
	header     []string
	rows       [][]string
	expires    time.Time
}

// csvImports holds imports in progress in memory, like join requests.
// shortcut: an import in progress is lost on restart; persist it if imports ever span sessions.
type csvImports struct {
	mu sync.Mutex
	m  map[string]csvImport
}

func (s *csvImports) put(imp csvImport) string {
	token := rand.Text()
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.m == nil {
		s.m = map[string]csvImport{}
	}
	now := time.Now()
	for k, v := range s.m {
		if now.After(v.expires) {
			delete(s.m, k)
		}
	}
	for len(s.m) >= maxImports {
		oldest := ""
		for k, v := range s.m {
			if oldest == "" || v.expires.Before(s.m[oldest].expires) {
				oldest = k
			}
		}
		delete(s.m, oldest)
	}
	imp.expires = now.Add(importLifetime)
	s.m[token] = imp
	return token
}

func (s *csvImports) get(token string, campaignID int64) (csvImport, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	imp, ok := s.m[token]
	return imp, ok && imp.campaignID == campaignID && time.Now().Before(imp.expires)
}

func (s *csvImports) drop(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, token)
}

// parseCSV reads a CSV whose first row is the header. Short rows are
// padded and blank rows dropped.
func parseCSV(r io.Reader) ([]string, [][]string, error) {
	raw, err := io.ReadAll(io.LimitReader(r, maxImportBytes+1))
	if err != nil {
		return nil, nil, err
	}
	if len(raw) > maxImportBytes {
		return nil, nil, errors.New("the CSV is larger than 5 MB")
	}
	cr := csv.NewReader(bytes.NewReader(bytes.TrimPrefix(raw, []byte("\xef\xbb\xbf"))))
	cr.FieldsPerRecord = -1
	cr.LazyQuotes = true
	records, err := cr.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	if len(records) < 2 {
		return nil, nil, fmt.Errorf("the file needs a header row and at least one row")
	}
	header := records[0]
	var rows [][]string
	for _, rec := range records[1:] {
		if strings.TrimSpace(strings.Join(rec, "")) == "" {
			continue
		}
		for len(rec) < len(header) {
			rec = append(rec, "")
		}
		rows = append(rows, rec)
	}
	if len(rows) > maxImportRows {
		return nil, nil, fmt.Errorf("the file has more than %d rows", maxImportRows)
	}
	return header, rows, nil
}

// Header words that suggest each image field.
var (
	nameHints        = []string{"name", "title"}
	linkHints        = []string{"link", "url", "image", "img", "src", "href", "picture", "photo"}
	categoryHints    = []string{"category", "type", "kind", "group", "tag"}
	descriptionHints = []string{"description", "desc", "notes", "details", "about", "summary"}
)

// guessSample bounds how many rows guessMapping looks at.
const guessSample = 200

// guessMapping picks the closest column for each field, never one column
// twice: the link is the column most full of http(s) links, the others go
// by header names. It is only the starting point the GM can change.
func guessMapping(header []string, rows [][]string) views.ImportMapping {
	sample := rows[:min(len(rows), guessSample)]
	used := map[int]bool{}
	best := func(score func(col int) float64) int {
		pick, top := noImportColumn, 0.0
		for i := range header {
			if s := score(i); !used[i] && s > top {
				pick, top = i, s
			}
		}
		if pick != noImportColumn {
			used[pick] = true
		}
		return pick
	}
	firstFree := func() int {
		return best(func(int) float64 { return 1 })
	}
	m := views.ImportMapping{}
	m.Link = best(func(i int) float64 { return 2*share(sample, i, isHTTPLink) + headerScore(header[i], linkHints) })
	m.Name = best(func(i int) float64 {
		return 2*headerScore(header[i], nameHints) + share(sample, i, func(s string) bool { return s != "" && !isHTTPLink(s) })
	})
	m.Category = best(func(i int) float64 { return headerScore(header[i], categoryHints) })
	m.Description = best(func(i int) float64 { return headerScore(header[i], descriptionHints) })
	// Name and link are required, so they always get some column.
	if m.Link == noImportColumn {
		m.Link = firstFree()
	}
	if m.Name == noImportColumn {
		m.Name = firstFree()
	}
	return m
}

// headerScore is 1 for a header that is one of hints, 0.5 for one that
// contains a hint ("Image Link"), else 0.
func headerScore(header string, hints []string) float64 {
	h := strings.ToLower(strings.TrimSpace(header))
	if slices.Contains(hints, h) {
		return 1
	}
	if slices.ContainsFunc(hints, func(hint string) bool { return strings.Contains(h, hint) }) {
		return 0.5
	}
	return 0
}

// share is the fraction of rows whose value in col satisfies ok.
func share(rows [][]string, col int, ok func(string) bool) float64 {
	if len(rows) == 0 {
		return 0
	}
	n := 0
	for _, row := range rows {
		if ok(cell(row, col)) {
			n++
		}
	}
	return float64(n) / float64(len(rows))
}

// httpHost is the host of an absolute http(s) link, or "" for anything else.
func httpHost(s string) string {
	u, err := url.Parse(s)
	if err != nil || (u.Scheme != "http" && u.Scheme != "https") {
		return ""
	}
	return u.Host
}

func isHTTPLink(s string) bool { return httpHost(s) != "" }

// cell is row's value in column col, or "" for no column.
func cell(row []string, col int) string {
	if col < 0 || col >= len(row) {
		return ""
	}
	return strings.TrimSpace(row[col])
}

// categoryKey says which spellings start out as one category: those that
// differ only in case and spacing. The GM can still split or merge any.
func categoryKey(s string) string { return strings.ToLower(strings.Join(strings.Fields(s), " ")) }

// reviewRows maps each CSV row through m, marking rows that cannot be imported.
func reviewRows(rows [][]string, m views.ImportMapping) []views.ImportRow {
	out := make([]views.ImportRow, 0, len(rows))
	for i, row := range rows {
		r := views.ImportRow{Index: i, Name: cell(row, m.Name), Link: cell(row, m.Link), Category: cell(row, m.Category), Description: cell(row, m.Description)}
		r.Host = httpHost(r.Link)
		switch {
		case r.Name == "":
			r.Problem = "no name"
		case len(r.Name) > 200:
			r.Problem = "name too long"
		case r.Host == "":
			r.Problem = "not an http(s) link"
		}
		out = append(out, r)
	}
	return out
}

// importCategories lists each distinct category value of rows on its own,
// so nothing is merged unless the GM gives two values the same target.
// A value's default target is an existing category with the same
// categoryKey, else the normalized (lowercase) categoryKey itself.
func importCategories(rows []views.ImportRow, existing []sqlc.ImageCategory) []views.ImportCategory {
	counts := map[string]int{}
	defaults := map[string]string{}
	for _, r := range rows {
		if r.Category != "" {
			counts[r.Category]++
			defaults[categoryKey(r.Category)] = categoryKey(r.Category)
		}
	}
	// An existing category wins, so the import never adds a near-duplicate.
	for _, c := range existing {
		defaults[categoryKey(c.Name)] = c.Name
	}
	cats := make([]views.ImportCategory, 0, len(counts))
	for s, n := range counts {
		cats = append(cats, views.ImportCategory{Source: s, Count: n, Target: defaults[categoryKey(s)]})
	}
	// Sorted by key so the spellings of one category sit together.
	slices.SortFunc(cats, func(a, b views.ImportCategory) int {
		if c := strings.Compare(categoryKey(a.Source), categoryKey(b.Source)); c != 0 {
			return c
		}
		return strings.Compare(a.Source, b.Source)
	})
	return cats
}

// Limits of the column overview and the keep filter.
const (
	overviewSamples = 3
	maxKeepValues   = 100
)

// columnOverview describes each CSV column: distinct values and a few samples.
func columnOverview(header []string, rows [][]string) []views.ImportColumn {
	cols := make([]views.ImportColumn, len(header))
	for i, h := range header {
		seen := map[string]bool{}
		col := views.ImportColumn{Header: h}
		for _, row := range rows {
			v := cell(row, i)
			if seen[v] {
				continue
			}
			seen[v] = true
			if v != "" && len(col.Samples) < overviewSamples {
				col.Samples = append(col.Samples, v)
			}
		}
		col.Distinct = len(seen)
		cols[i] = col
	}
	return cols
}

// keepValues lists the distinct values of column col with their row counts,
// spellings of one value side by side, marking the kept ones. A column with
// too many values to tick through returns none and true.
func keepValues(rows [][]string, col int, kept map[string]bool) ([]views.ImportValue, bool) {
	counts := map[string]int{}
	for _, row := range rows {
		counts[cell(row, col)]++
	}
	if len(counts) > maxKeepValues {
		return nil, true
	}
	vals := make([]views.ImportValue, 0, len(counts))
	for v, n := range counts {
		vals = append(vals, views.ImportValue{Value: v, Count: n, Kept: kept[v]})
	}
	slices.SortFunc(vals, func(a, b views.ImportValue) int {
		if c := strings.Compare(categoryKey(a.Value), categoryKey(b.Value)); c != 0 {
			return c
		}
		return strings.Compare(a.Value, b.Value)
	})
	return vals, false
}

// formValues reads a repeated field from the POST form, else the query.
func formValues(c *gin.Context, name string) ([]string, bool) {
	if vs, ok := c.GetPostFormArray(name); ok {
		return vs, true
	}
	return c.GetQueryArray(name)
}

// parseRowSet reads the ticked rows of n as app.js sends them: runs and
// single rows ("0-78,80,82"), or after "!" the rows that are not ticked,
// whichever was shorter. Runs may span rows that cannot be imported.
func parseRowSet(s string, n int) []bool {
	set := make([]bool, n)
	invert := strings.HasPrefix(s, "!")
	for part := range strings.SplitSeq(strings.TrimPrefix(s, "!"), ",") {
		from, till, isRange := strings.Cut(part, "-")
		if !isRange {
			till = from
		}
		a, errA := strconv.Atoi(from)
		b, errB := strconv.Atoi(till)
		if errA != nil || errB != nil || a < 0 || a > b {
			continue
		}
		for i := a; i <= min(b, n-1); i++ {
			set[i] = true
		}
	}
	if invert {
		for i := range set {
			set[i] = !set[i]
		}
	}
	return set
}

// buildReview applies the review form so far (columns, category targets,
// ticked rows and the row filter) to the CSV. A first visit has no form:
// the guessed columns, the default targets and every importable row.
func buildReview(c *gin.Context, imp csvImport, v *views.ImportView) {
	v.Header = imp.header
	v.Map = guessMapping(imp.header, imp.rows)
	read := func(dst *int, name string) {
		vs, _ := formValues(c, name)
		if len(vs) == 0 {
			return
		}
		if n, err := strconv.Atoi(vs[0]); err == nil && n >= noImportColumn && n < len(imp.header) {
			*dst = n
		}
	}
	read(&v.Map.Name, views.ImportColName)
	read(&v.Map.Link, views.ImportColLink)
	read(&v.Map.Category, views.ImportColCategory)
	read(&v.Map.Description, views.ImportColDescription)

	// Keep only rows whose filter column holds one of the ticked values (all
	// rows while none is ticked). Rows keep their CSV index.
	v.Total, v.KeepCol = len(imp.rows), noImportColumn
	read(&v.KeepCol, views.ImportKeepCol)
	v.Columns = columnOverview(imp.header, imp.rows)
	kept := map[string]bool{}
	if vals, _ := formValues(c, views.ImportKeepVal); v.KeepCol != noImportColumn {
		for _, val := range vals {
			if s, ok := strings.CutPrefix(val, views.ImportKeepPrefix); ok {
				kept[s] = true
			}
		}
		v.KeepValues, v.TooManyValues = keepValues(imp.rows, v.KeepCol, kept)
	}
	v.Rows = slices.DeleteFunc(reviewRows(imp.rows, v.Map), func(r views.ImportRow) bool {
		return len(kept) > 0 && !v.TooManyValues && !kept[cell(imp.rows[r.Index], v.KeepCol)]
	})
	v.Cats = importCategories(v.Rows, v.Categories)
	sources, _ := formValues(c, views.ImportCatSource)
	targets, _ := formValues(c, views.ImportCatTarget)
	chosen := map[string]string{}
	for i := range min(len(sources), len(targets)) {
		chosen[sources[i]] = strings.TrimSpace(targets[i])
	}
	target := map[string]string{}
	for i, cat := range v.Cats {
		if t, ok := chosen[cat.Source]; ok {
			v.Cats[i].Target = t
		}
		target[cat.Source] = v.Cats[i].Target
	}

	// Once the form has been seen, only the rows it ticks stay ticked.
	encoded, reviewed := formValues(c, views.ImportRowsField)
	var picked []bool
	if reviewed {
		picked = parseRowSet(encoded[0], len(imp.rows))
	}
	for i := range v.Rows {
		r := &v.Rows[i]
		r.Target = target[r.Category]
		r.Ticked = r.Problem == "" && (!reviewed || picked[r.Index])
	}
	if q, _ := formValues(c, views.ImportFilterQ); len(q) > 0 {
		v.FilterQ = q[0]
	}
	if f, _ := formValues(c, views.ImportFilterCategory); len(f) > 0 {
		v.FilterCategory = f[0]
	}
	if f, _ := formValues(c, views.ImportFilterStatus); len(f) > 0 {
		v.FilterStatus = f[0]
	}
}

func (a *App) importUploadPage(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	a.render(c, http.StatusOK, views.ImportUploadPage(a.layout(w, "Import images", views.NavImages, &camp, false), id))
}

func (a *App) importUpload(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxImportBytes+1<<20)
	f, err := csvSource(c)
	if err != nil {
		a.fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	defer func() { _ = f.Close() }()
	header, rows, err := parseCSV(f)
	if err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "Could not read the CSV: "+err.Error())
		return
	}
	token := a.imports.put(csvImport{campaignID: id, header: header, rows: rows})
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/campaigns/%d/images/import/%s", id, token))
}

// csvImportForm is the CSV link, used when no file is uploaded.
type csvImportForm struct {
	URL string `form:"url" validate:"required,http_url,max=2000"`
}

// csvClient fetches CSV links; the timeout keeps a dead link from hanging the request.
var csvClient = &http.Client{Timeout: 20 * time.Second}

// csvSource opens the uploaded CSV file, or else downloads the CSV link.
func csvSource(c *gin.Context) (io.ReadCloser, error) {
	if fh, err := c.FormFile(importFileField); err == nil {
		if fh.Size > maxImportBytes {
			return nil, errors.New("the CSV is larger than 5 MB")
		}
		return fh.Open()
	}
	var form csvImportForm
	if err := c.ShouldBind(&form); err != nil {
		return nil, errors.New("choose a CSV file or give an http(s) link to one")
	}
	req, err := http.NewRequestWithContext(c.Request.Context(), http.MethodGet, strings.TrimSpace(form.URL), nil)
	if err != nil {
		return nil, errors.New("that link is not valid")
	}
	res, err := csvClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not download the CSV: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		_ = res.Body.Close()
		return nil, fmt.Errorf("could not download the CSV: %s", res.Status)
	}
	return res.Body, nil
}

// importParam loads the import in progress for :campaign and :token.
func (a *App) importParam(c *gin.Context) (int64, string, csvImport, bool) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return 0, "", csvImport{}, false
	}
	token := c.Param("token")
	imp, ok := a.imports.get(token, id)
	if !ok {
		a.failBack(c, http.StatusNotFound, "This import has expired. Upload the file again.", fmt.Sprintf("/campaigns/%d/images/import", id), "Start the import again")
	}
	return id, token, imp, ok
}

func (a *App) importReviewPage(c *gin.Context) {
	id, token, imp, ok := a.importParam(c)
	if !ok {
		return
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	cats, err := a.Store.ListImageCategories(c.Request.Context(), id)
	if err != nil {
		a.serverError(c, err)
		return
	}
	v := views.ImportView{CampaignID: id, Token: token, Categories: cats}
	buildReview(c, imp, &v)
	if isHTMX(c) {
		a.render(c, http.StatusOK, views.ImportReview(v))
		return
	}
	a.render(c, http.StatusOK, views.ImportReviewPage(a.layout(w, "Import images", views.NavImages, &camp, false), v))
}

// importCommit creates the ticked rows, each in the category its value
// was mapped to (none for an empty target).
func (a *App) importCommit(c *gin.Context) {
	id, token, imp, ok := a.importParam(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	cats, err := a.Store.ListImageCategories(ctx, id)
	if err != nil {
		a.serverError(c, err)
		return
	}
	v := views.ImportView{Categories: cats}
	buildReview(c, imp, &v)
	n := 0
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		for _, r := range v.Rows {
			if !r.Ticked {
				continue
			}
			category, err := ensureCategory(ctx, q, id, r.Target)
			if err != nil {
				return err
			}
			if _, err := q.CreateImage(ctx, sqlc.CreateImageParams{
				CampaignID: id, Name: r.Name, Url: nullString(r.Link), CategoryID: category, Description: nullString(r.Description),
			}); err != nil {
				return err
			}
			n++
		}
		return nil
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.imports.drop(token)
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/campaigns/%d/images?imported=%d", id, n))
}
