package compendium

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

// fakeMirror is a tiny 5etools data folder: two books sharing items, races and a class file.
var fakeMirror = map[string]string{
	"books.json": `{"book":[
		{"name":"Player's Handbook (2024)","id":"XPHB","source":"XPHB","group":"core","published":"2024-09-17","author":"Wizards of the Coast"},
		{"name":"Tasha's","id":"TCE","source":"TCE","group":"supplement","published":"2020-11-17"},
		{"name":"Eberron","id":"ERLW","source":"ERLW","group":"setting","published":"2019-11-19"},
		{"name":"Homebrew Tome","id":"HBT","source":"HBT","group":"homebrew","published":"2021-01-01","author":"Someone"},
		{"name":"UA Subclasses","id":"UASub","source":"UASub","group":"other","published":"2022-02-02"}]}`,
	"adventures.json":             `{"adventure":[{"name":"Lost Mine","id":"LMOP","source":"LMOP","group":"mixed","published":"2014-09-30"}]}`,
	"bestiary/index.json":         `{"XPHB":"bestiary-xphb.json"}`,
	"bestiary/fluff-index.json":   `{}`,
	"bestiary/bestiary-xphb.json": `{"monster":[{"name":"Goblin","source":"XPHB"},{"name":"Other","source":"MM"}]}`,
	"spells/index.json":           `{"XPHB":"spells-xphb.json"}`,
	"spells/fluff-index.json":     `{}`,
	"spells/spells-xphb.json":     `{"spell":[{"name":"Fireball","source":"XPHB"}]}`,
	"spells/sources.json":         `{"XPHB":{"Fireball":{"class":[{"name":"Wizard","source":"XPHB"}]}},"MM":{}}`,
	"items.json":                  `{"item":[{"name":"Wand","source":"XPHB"},{"name":"Cloak","source":"TCE"}],"itemGroup":[]}`,
	"items-base.json":             `{"baseitem":[{"name":"Dagger","source":"PHB"}]}`,
	"races.json":                  `{"race":[{"name":"Elf","source":"XPHB"},{"name":"Tiefling","source":"TCE"}],"subrace":[]}`,
	"class/index.json":            `{"wizard":"class-wizard.json","fighter":"class-fighter.json"}`,
	"class/class-wizard.json": `{"class":[{"name":"Wizard","source":"PHB"}],
		"subclass":[{"name":"Bladesinger","className":"Wizard","classSource":"PHB","source":"TCE"}],
		"classFeature":[],"subclassFeature":[{"name":"Song","className":"Wizard","subclassShortName":"Bladesinging","source":"TCE"}]}`,
	"class/class-fighter.json": `{"class":[{"name":"Fighter","source":"PHB"}]}`,
}

func serveMirror(t *testing.T, files map[string]string, fail bool) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, ok := files[r.URL.Path[1:]]
		switch {
		case fail:
			http.Error(w, "down", http.StatusBadGateway)
		case !ok:
			http.NotFound(w, r)
		default:
			_, _ = w.Write([]byte(body))
		}
	}))
	old := fiveToolsBase
	fiveToolsBase = srv.URL + "/"
	t.Cleanup(func() { srv.Close(); fiveToolsBase = old })
}

func TestFiveToolsBooks(t *testing.T) {
	serveMirror(t, fakeMirror, false)
	dir := t.TempDir()
	books, err := FiveToolsBooks(context.Background(), dir)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, b := range books {
		got = append(got, b.Group+":"+b.ID+":"+b.Publisher)
	}
	want := []string{
		"core:XPHB:Wizards of the Coast",
		"supplement:TCE:Wizards of the Coast",
		"supplement:ERLW:Wizards of the Coast",
		"adventure:LMOP:Wizards of the Coast",
		"ua:UASUB:Wizards of the Coast",
		"third-party:HBT:Someone",
	}
	if !slices.Equal(got, want) {
		t.Fatalf("got %v\nwant %v", got, want)
	}

	// With the mirror down, the cache answers.
	serveMirror(t, nil, true)
	cached, err := FiveToolsBooks(context.Background(), dir)
	if err != nil || !slices.Equal(cached, books) {
		t.Fatalf("cache fallback: %v %v", cached, err)
	}
	if _, err := FiveToolsBooks(context.Background(), t.TempDir()); err == nil {
		t.Fatal("no cache and no mirror must fail")
	}
}

func readDoc(t *testing.T, path string) map[string][]map[string]any {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var doc map[string][]map[string]any
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatal(err)
	}
	return doc
}

func TestDownloadFiveTools(t *testing.T) {
	serveMirror(t, fakeMirror, false)
	dir := t.TempDir()
	books := []Book{{ID: "XPHB", Name: "PHB", Group: GroupCore}, {ID: "TCE", Name: "Tasha", Group: GroupSupplement}}
	if err := DownloadFiveTools(context.Background(), dir, books); err != nil {
		t.Fatal(err)
	}
	ls := func(id string) []string {
		entries, err := os.ReadDir(filepath.Join(dir, FiveToolsPrefix+id))
		if err != nil {
			t.Fatal(err)
		}
		var names []string
		for _, e := range entries {
			names = append(names, e.Name())
		}
		return names
	}
	tests := []struct {
		id   string
		want []string
	}{
		{"xphb", []string{"bestiary.json", "book.json", "items-base.json", "items.json", "races.json", "spell-classes.json", "spells.json"}},
		{"tce", []string{"book.json", "class-wizard.json", "items-base.json", "items.json", "races.json"}},
	}
	for _, tc := range tests {
		if got := ls(tc.id); !slices.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.id, got, tc.want)
		}
	}
	if items := readDoc(t, filepath.Join(dir, "5etools-tce", "items.json"))["item"]; len(items) != 1 || items[0]["name"] != "Cloak" {
		t.Errorf("items not filtered: %v", items)
	}
	if base := readDoc(t, filepath.Join(dir, "5etools-tce", "items-base.json"))["baseitem"]; len(base) != 1 {
		t.Errorf("items-base must stay whole: %v", base)
	}
	// Subclass-only book keeps the parent class.
	wiz := readDoc(t, filepath.Join(dir, "5etools-tce", "class-wizard.json"))
	if len(wiz["class"]) != 1 || len(wiz["subclass"]) != 1 || len(wiz["subclassFeature"]) != 1 {
		t.Errorf("class file: %v", wiz)
	}
	// No temp folders left, and the book can be listed and removed.
	if matches, _ := filepath.Glob(filepath.Join(dir, ".tmp-*")); len(matches) != 0 {
		t.Errorf("temp left: %v", matches)
	}
	imported, err := ImportedFiveTools(dir)
	if err != nil || len(imported) != 2 || imported[0].ID != "XPHB" {
		t.Fatalf("imported: %v %v", imported, err)
	}
	if err := RemoveFiveTools(dir, "XPHB"); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "5etools-xphb")); !os.IsNotExist(err) {
		t.Error("folder not removed")
	}
}

func TestDownloadFiveToolsFailureLeavesNothing(t *testing.T) {
	files := map[string]string{}
	for k, v := range fakeMirror {
		files[k] = v
	}
	files["items.json"] = "not json"
	serveMirror(t, files, false)
	dir := t.TempDir()
	if err := DownloadFiveTools(context.Background(), dir, []Book{{ID: "XPHB"}}); err == nil {
		t.Fatal("want error")
	}
	if entries, _ := os.ReadDir(dir); len(entries) != 0 {
		t.Errorf("left behind: %v", entries)
	}
}

func TestRemoveFiveToolsRejectsBadIDs(t *testing.T) {
	dir := t.TempDir()
	for _, id := range []string{"", "..", "../x", "a/b", `a\b`, "a b"} {
		if err := RemoveFiveTools(dir, id); err == nil {
			t.Errorf("%q accepted", id)
		}
	}
}
