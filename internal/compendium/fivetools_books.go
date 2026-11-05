package compendium

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

// FiveToolsPrefix starts the name of every downloaded book's folder.
const FiveToolsPrefix = "5etools-"

// Book groups, in display order.
const (
	GroupCore       = "core"
	GroupSupplement = "supplement"
	GroupAdventure  = "adventure"
	GroupUA         = "ua"
	GroupThirdParty = "third-party"
)

const (
	bookMaxBody       = 64 << 20
	fiveToolsCacheDir = ".5etools"
	bookFile          = "book.json"
	defaultPublisher  = "Wizards of the Coast"
)

// fiveToolsBase is the data folder of the mirror; tests point it at httptest.
var fiveToolsBase = "https://raw.githubusercontent.com/5etools-mirror-3/5etools-src/main/data/"

var bookIDPattern = regexp.MustCompile(`^[a-z0-9_-]+$`)

// Book is one 5etools source book, normalized.
type Book struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Published string `json:"published"`
	Publisher string `json:"publisher"`
}

type rawBook struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Published string `json:"published"`
	Author    string `json:"author"`
}

func bookGroup(id, group string, adventure bool) string {
	switch {
	case strings.HasPrefix(strings.ToUpper(id), "UA"):
		return GroupUA
	case adventure:
		return GroupAdventure
	case group == "core":
		return GroupCore
	case group == "homebrew" || group == "third-party":
		return GroupThirdParty
	}
	return GroupSupplement
}

func normalizeBooks(raw []rawBook, adventure bool) []Book {
	out := make([]Book, 0, len(raw))
	for _, r := range raw {
		if r.ID == "" {
			continue
		}
		b := Book{
			ID:        strings.ToUpper(r.ID),
			Name:      r.Name,
			Group:     bookGroup(r.ID, r.Group, adventure),
			Published: r.Published,
			Publisher: r.Author,
		}
		// 5etools credits WotC's own books to "Wizards RPG Team".
		if b.Publisher == "" || b.Publisher == "Wizards RPG Team" {
			b.Publisher = defaultPublisher
		}
		out = append(out, b)
	}
	return out
}

func sortBooks(books []Book) {
	order := []string{GroupCore, GroupSupplement, GroupAdventure, GroupUA, GroupThirdParty}
	slices.SortStableFunc(books, func(a, b Book) int {
		if c := slices.Index(order, a.Group) - slices.Index(order, b.Group); c != 0 {
			return c
		}
		if c := strings.Compare(b.Published, a.Published); c != 0 {
			return c
		}
		return strings.Compare(a.Name, b.Name)
	})
}

// fetchJSON decodes one mirror file into v.
func fetchJSON(ctx context.Context, rel string, v any) error {
	body, err := fetch(ctx, fiveToolsBase+rel, bookMaxBody)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("5etools %s: %w", rel, err)
	}
	return nil
}

func fetchBooks(ctx context.Context) ([]Book, error) {
	var books struct {
		Book []rawBook `json:"book"`
	}
	var adventures struct {
		Adventure []rawBook `json:"adventure"`
	}
	if err := fetchJSON(ctx, "books.json", &books); err != nil {
		return nil, err
	}
	if err := fetchJSON(ctx, "adventures.json", &adventures); err != nil {
		return nil, err
	}
	all := append(normalizeBooks(books.Book, false), normalizeBooks(adventures.Adventure, true)...)
	sortBooks(all)
	return all, nil
}

// FiveToolsBooks lists every 5etools book, falling back to the cached list
// when the mirror cannot be reached.
func FiveToolsBooks(ctx context.Context, dir string) ([]Book, error) {
	cache := filepath.Join(dir, fiveToolsCacheDir, "books.json")
	books, err := fetchBooks(ctx)
	if err != nil {
		data, rerr := os.ReadFile(cache) //nolint:gosec // our own cache path
		if rerr != nil {
			return nil, err
		}
		var cached []Book
		if json.Unmarshal(data, &cached) != nil {
			return nil, err
		}
		return cached, nil
	}
	data, err := json.Marshal(books)
	if err == nil && os.MkdirAll(filepath.Dir(cache), 0o750) == nil { //nolint:gosec // dir is the operator's DATA_DIR, never request input
		_ = os.WriteFile(cache, data, 0o600) //nolint:gosec // as above; a failed cache write only costs offline use
	}
	return books, nil
}

// mirror fetches each file once per download run.
type mirror struct {
	ctx   context.Context
	files map[string]json.RawMessage
}

// get returns nil, nil for a file the mirror does not have.
func (m *mirror) get(rel string) (json.RawMessage, error) {
	if raw, ok := m.files[rel]; ok {
		return raw, nil
	}
	body, err := fetch(m.ctx, fiveToolsBase+rel, bookMaxBody)
	if err != nil {
		if strings.Contains(err.Error(), "no page at") {
			m.files[rel] = nil
			return nil, nil
		}
		return nil, err
	}
	m.files[rel] = body
	return body, nil
}

// index reads a {"ID": "file.json"} index into a map.
func (m *mirror) index(rel string) (map[string]string, error) {
	raw, err := m.get(rel)
	if err != nil || raw == nil {
		return nil, err
	}
	var idx map[string]string
	if err := json.Unmarshal(raw, &idx); err != nil {
		return nil, fmt.Errorf("5etools %s: %w", rel, err)
	}
	return idx, nil
}

// sourceOf decodes only the source of an entry.
type sourceOf struct {
	Source string `json:"source"`
}

// filterSource keeps the entries of a JSON array whose source is id.
func filterSource(arr json.RawMessage, id string) ([]json.RawMessage, error) {
	var items []json.RawMessage
	if err := json.Unmarshal(arr, &items); err != nil {
		return nil, err
	}
	var kept []json.RawMessage
	for _, it := range items {
		var s sourceOf
		if json.Unmarshal(it, &s) == nil && s.Source == id {
			kept = append(kept, it)
		}
	}
	return kept, nil
}

// filterDoc keeps, per top-level array of a 5etools file (only those in keys
// when given), the entries of source id. A nil result means nothing is left.
func filterDoc(raw json.RawMessage, id string, keys ...string) (map[string][]json.RawMessage, error) {
	if raw == nil {
		return nil, nil
	}
	var doc map[string]json.RawMessage
	if err := json.Unmarshal(raw, &doc); err != nil {
		return nil, err
	}
	out := map[string][]json.RawMessage{}
	for k, v := range doc {
		if len(keys) > 0 && !slices.Contains(keys, k) {
			continue
		}
		kept, err := filterSource(v, id)
		if err != nil {
			continue // _meta and other non-array keys
		}
		if len(kept) > 0 {
			out[k] = kept
		}
	}
	if len(out) == 0 {
		return nil, nil
	}
	return out, nil
}

// bookWriter collects a book's files in a temp folder.
type bookWriter struct{ dir string }

func (w bookWriter) write(name string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(w.dir, name), data, 0o600)
}

// writeDoc writes the filtered rel file as name, unless nothing is left.
// idx, for a per-book file, names the other books' files copies can reach.
// It reports whether it wrote.
func (w bookWriter) writeDoc(m *mirror, id, rel, name string, idx map[string]string, keys ...string) (bool, error) {
	raw, err := m.get(rel)
	if err != nil {
		return false, err
	}
	doc, err := filterDoc(raw, id, keys...)
	if err != nil {
		return false, fmt.Errorf("5etools %s: %w", rel, err)
	}
	if doc == nil {
		return false, nil
	}
	other := func(src string) (json.RawMessage, error) {
		if file, ok := idx[src]; ok {
			return m.get(filepath.ToSlash(filepath.Join(filepath.Dir(rel), file)))
		}
		return nil, nil
	}
	return true, w.writeWithCopies(name, doc, raw, other)
}

// copyTargetsKey holds, per list, the records kept entries _copy from other
// books: ftRecords resolves against them but never imports them.
const copyTargetsKey = "_copyTargets"

// writeWithCopies writes doc plus the copy targets the filter dropped.
func (w bookWriter) writeWithCopies(name string, doc map[string][]json.RawMessage, raw json.RawMessage, other func(string) (json.RawMessage, error)) error {
	targets, err := copyTargets(doc, raw, other)
	if err != nil {
		return fmt.Errorf("5etools %s: %w", name, err)
	}
	if len(targets) == 0 {
		return w.write(name, doc)
	}
	out := map[string]any{copyTargetsKey: targets}
	for k, v := range doc {
		out[k] = v
	}
	return w.write(name, out)
}

// copyIndex maps a list's records by both keys ftResolve looks up.
func copyIndex(recs []json.RawMessage) map[string]json.RawMessage {
	idx := map[string]json.RawMessage{}
	for _, r := range recs {
		var c ftCopyRef
		if json.Unmarshal(r, &c) == nil {
			idx[c.ftNamed.key()] = r
			idx[c.key()] = r
		}
	}
	return idx
}

// copyTargets finds, per list of doc, the records its entries _copy that are
// not in doc: in raw, else in other(source), another book's file (nil if
// none). It follows copies of copies as deep as ftResolve does.
func copyTargets(doc map[string][]json.RawMessage, raw json.RawMessage, other func(string) (json.RawMessage, error)) (map[string][]json.RawMessage, error) {
	pools := map[string]map[string]json.RawMessage{} // by source file ("" is raw) and list
	pool := func(src, list string) (map[string]json.RawMessage, error) {
		if p, ok := pools[src+"|"+list]; ok {
			return p, nil
		}
		body := raw
		if src != "" {
			if other == nil {
				return nil, nil
			}
			var err error
			if body, err = other(src); err != nil {
				return nil, err
			}
		}
		var all map[string]json.RawMessage
		var recs []json.RawMessage
		if body != nil {
			if err := json.Unmarshal(body, &all); err != nil {
				return nil, err
			}
			_ = json.Unmarshal(all[list], &recs) // _meta and other non-array keys
		}
		pools[src+"|"+list] = copyIndex(recs)
		return pools[src+"|"+list], nil
	}
	out := map[string][]json.RawMessage{}
	for list, recs := range doc {
		have := copyIndex(recs)
		queue := recs
		for depth := 0; depth <= 4 && len(queue) > 0; depth++ {
			var next []json.RawMessage
			for _, r := range queue {
				var rec struct {
					Copy *ftCopyRef `json:"_copy"`
				}
				if json.Unmarshal(r, &rec) != nil || rec.Copy == nil || have[rec.Copy.key()] != nil || have[rec.Copy.ftNamed.key()] != nil {
					continue
				}
				for _, src := range []string{"", rec.Copy.Source} {
					p, err := pool(src, list)
					if err != nil {
						return nil, err
					}
					t, ok := p[rec.Copy.key()]
					if !ok {
						t, ok = p[rec.Copy.ftNamed.key()]
					}
					if !ok {
						continue
					}
					maps.Copy(have, copyIndex([]json.RawMessage{t}))
					out[list] = append(out[list], t)
					next = append(next, t)
					break
				}
			}
			queue = next
		}
	}
	return out, nil
}

// indexed is a per-book file named by an index.
type indexed struct{ index, name string }

var (
	perBookFiles = []indexed{
		{"bestiary/index.json", "bestiary.json"},
		{"bestiary/fluff-index.json", "fluff-bestiary.json"},
		{"spells/index.json", "spells.json"},
		{"spells/fluff-index.json", "fluff-spells.json"},
	}
	sharedFiles = []indexed{
		{"races.json", "races.json"},
		{"fluff-races.json", "fluff-races.json"},
		{"backgrounds.json", "backgrounds.json"},
		{"fluff-backgrounds.json", "fluff-backgrounds.json"},
		{"feats.json", "feats.json"},
		{"conditionsdiseases.json", "conditionsdiseases.json"},
		{"variantrules.json", "variantrules.json"},
		{"fluff-items.json", "fluff-items.json"},
	}
	classKeys = []string{"class", "subclass", "classFeature", "subclassFeature"}
)

func (m *mirror) writeBook(w bookWriter, b Book) error {
	for _, f := range perBookFiles {
		idx, err := m.index(f.index)
		if err != nil {
			return err
		}
		file, ok := idx[b.ID]
		if !ok {
			continue
		}
		if _, err := w.writeDoc(m, b.ID, filepath.ToSlash(filepath.Join(filepath.Dir(f.index), file)), f.name, idx); err != nil {
			return err
		}
	}
	if err := m.writeSpellClasses(w, b); err != nil {
		return err
	}
	hasItems, err := w.writeDoc(m, b.ID, "items.json", "items.json", nil)
	if err != nil {
		return err
	}
	// Magic items resolve their base item from the full file.
	base, err := m.get("items-base.json")
	if err != nil {
		return err
	}
	ownBase, err := filterDoc(base, b.ID)
	if err != nil {
		return fmt.Errorf("5etools items-base.json: %w", err)
	}
	if base != nil && (hasItems || ownBase != nil) {
		if err := os.WriteFile(filepath.Join(w.dir, "items-base.json"), base, 0o600); err != nil {
			return err
		}
	}
	for _, f := range sharedFiles {
		if _, err := w.writeDoc(m, b.ID, f.index, f.name, nil); err != nil {
			return err
		}
	}
	return m.writeClasses(w, b)
}

func (m *mirror) writeSpellClasses(w bookWriter, b Book) error {
	raw, err := m.get("spells/sources.json")
	if err != nil || raw == nil {
		return err
	}
	var all map[string]json.RawMessage
	if err := json.Unmarshal(raw, &all); err != nil {
		return fmt.Errorf("5etools spells/sources.json: %w", err)
	}
	if v, ok := all[b.ID]; ok {
		return w.write("spell-classes.json", map[string]json.RawMessage{b.ID: v})
	}
	return nil
}

func (m *mirror) writeClasses(w bookWriter, b Book) error {
	idx, err := m.index("class/index.json")
	if err != nil {
		return err
	}
	for _, name := range slices.Sorted(maps.Keys(idx)) {
		raw, err := m.get("class/" + idx[name])
		if err != nil {
			return err
		}
		doc, err := filterDoc(raw, b.ID, classKeys...)
		if err != nil {
			return fmt.Errorf("5etools class/%s: %w", idx[name], err)
		}
		if doc == nil {
			continue
		}
		if err := m.keepClassParents(raw, doc); err != nil {
			return err
		}
		if err := w.writeWithCopies("class-"+name+".json", doc, raw, nil); err != nil {
			return err
		}
	}
	return nil
}

type classRef struct {
	Name      string `json:"name"`
	ClassName string `json:"className"`
}

// keepClassParents adds the class entries the kept subclasses and subclass
// features belong to, whatever their own source.
func (m *mirror) keepClassParents(raw json.RawMessage, doc map[string][]json.RawMessage) error {
	parents := map[string]bool{}
	for _, k := range []string{"subclass", "subclassFeature"} {
		for _, e := range doc[k] {
			var r classRef
			if json.Unmarshal(e, &r) == nil && r.ClassName != "" {
				parents[r.ClassName] = true
			}
		}
	}
	if len(parents) == 0 {
		return nil
	}
	var all struct {
		Class []json.RawMessage `json:"class"`
	}
	if err := json.Unmarshal(raw, &all); err != nil {
		return err
	}
	have := map[string]bool{}
	for _, e := range doc["class"] {
		var r classRef
		if json.Unmarshal(e, &r) == nil {
			have[r.Name] = true
		}
	}
	for _, e := range all.Class {
		var r classRef
		if json.Unmarshal(e, &r) == nil && parents[r.Name] && !have[r.Name] {
			doc["class"] = append(doc["class"], e)
		}
	}
	return nil
}

func bookFolder(id string) string { return FiveToolsPrefix + strings.ToLower(id) }

// DownloadFiveTools writes a folder under dir for each book. A failed book
// leaves no folder behind.
func DownloadFiveTools(ctx context.Context, dir string, books []Book) error {
	if err := os.MkdirAll(dir, 0o750); err != nil { //nolint:gosec // dir is the operator's DATA_DIR, never request input
		return err
	}
	m := &mirror{ctx: ctx, files: map[string]json.RawMessage{}}
	for _, b := range books {
		if !bookIDPattern.MatchString(strings.ToLower(b.ID)) {
			return fmt.Errorf("bad 5etools book id %q", b.ID)
		}
		if err := downloadBook(m, dir, b); err != nil {
			return fmt.Errorf("5etools %s: %w", b.ID, err)
		}
	}
	return nil
}

func downloadBook(m *mirror, dir string, b Book) error {
	tmp, err := os.MkdirTemp(dir, ".tmp-5etools-")
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(tmp) }() //nolint:gosec // a temp dir under DATA_DIR; no-op after the rename
	w := bookWriter{tmp}
	if err := w.write(bookFile, b); err != nil {
		return err
	}
	if err := m.writeBook(w, b); err != nil {
		return err
	}
	final := filepath.Join(dir, bookFolder(b.ID))
	// final is DATA_DIR plus a book id DownloadFiveTools checked.
	if err := os.RemoveAll(final); err != nil { //nolint:gosec // see above
		return err
	}
	return os.Rename(tmp, final) //nolint:gosec // see above
}

// RemoveFiveTools deletes a book's folder.
func RemoveFiveTools(dir, id string) error {
	id = strings.ToLower(id)
	if !bookIDPattern.MatchString(id) {
		return errors.New("bad 5etools book id")
	}
	return os.RemoveAll(filepath.Join(dir, FiveToolsPrefix+id))
}

// ImportedFiveTools lists the books downloaded into dir.
func ImportedFiveTools(dir string) ([]Book, error) {
	files, err := filepath.Glob(filepath.Join(dir, FiveToolsPrefix+"*", bookFile))
	if err != nil {
		return nil, err
	}
	var books []Book
	for _, f := range files {
		data, err := os.ReadFile(f) //nolint:gosec // glob under our own dir
		if err != nil {
			continue
		}
		var b Book
		if json.Unmarshal(data, &b) == nil && b.ID != "" {
			books = append(books, b)
		}
	}
	sortBooks(books)
	return books, nil
}
