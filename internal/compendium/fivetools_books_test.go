package compendium

import (
	"context"
	"encoding/json"
	"errors"
	"maps"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
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
	"bestiary/index.json":         `{"XPHB":"bestiary-xphb.json","MM":"bestiary-mm.json"}`,
	"bestiary/fluff-index.json":   `{}`,
	"bestiary/bestiary-xphb.json": `{"monster":[{"name":"Goblin","source":"XPHB"},{"name":"Other","source":"MM"},{"name":"Dire Wolf","source":"XPHB","_copy":{"name":"Wolf","source":"MM"}}]}`,
	"bestiary/bestiary-mm.json":   `{"monster":[{"name":"Wolf","source":"MM","_copy":{"name":"Pup","source":"MM"}},{"name":"Pup","source":"MM","size":["S"]},{"name":"Bear","source":"MM"}]}`,
	"spells/index.json":           `{"XPHB":"spells-xphb.json"}`,
	"spells/fluff-index.json":     `{}`,
	"spells/spells-xphb.json":     `{"spell":[{"name":"Fireball","source":"XPHB"}]}`,
	"spells/sources.json":         `{"XPHB":{"Fireball":{"class":[{"name":"Wizard","source":"XPHB"}]}},"MM":{}}`,
	"items.json":                  `{"item":[{"name":"Wand","source":"XPHB"},{"name":"Cloak","source":"TCE","rarity":"rare"},{"name":"Wand Cloak","source":"XPHB","_copy":{"name":"Cloak","source":"TCE"}}],"itemGroup":[]}`,
	"items-base.json":             `{"baseitem":[{"name":"Dagger","source":"PHB"}]}`,
	"magicvariants.json":          `{"magicvariant":[{"name":"+1 Weapon","requires":[{"weapon":true}],"inherits":{"source":"XPHB"}},{"name":"Vorpal Sword","inherits":{"source":"DMG"}}]}`,
	"races.json":                  `{"race":[{"name":"Elf","source":"XPHB"},{"name":"Tiefling","source":"TCE"},{"name":"Half-Elf","source":"PHB"}],"subrace":[{"name":"Variant; Drow Descent","source":"XPHB","raceName":"Half-Elf","raceSource":"PHB"}]}`,
	"class/index.json":            `{"wizard":"class-wizard.json","fighter":"class-fighter.json","x/../../evil":"class-evil.json"}`,
	"class/class-evil.json":       `{"class":[{"name":"Evil","source":"TCE"}]}`,
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

func TestMirrorGetMissingFile(t *testing.T) {
	serveMirror(t, fakeMirror, false)
	m := &mirror{ctx: context.Background(), files: map[string]json.RawMessage{}}
	raw, err := m.get("missing.json")
	if err != nil || raw != nil {
		t.Fatalf("get(missing) = %q, %v; want nil, nil", raw, err)
	}
	_, err = fetch(context.Background(), fiveToolsBase+"missing.json", bookMaxBody)
	if !errors.Is(err, ErrNoPage) || !strings.Contains(err.Error(), "no page at "+fiveToolsBase) {
		t.Errorf("fetch 404 = %v; want ErrNoPage with the link", err)
	}
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
	if err := DownloadFiveTools(context.Background(), dir, books, nil); err != nil {
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
		{"xphb", []string{"bestiary.json", "book.json", "items-base.json", "items.json", "magicvariants.json", "races.json", "spell-classes.json", "spells.json"}},
		{"tce", []string{"book.json", "class-wizard.json", "items-base.json", "items.json", "races.json"}},
	}
	for _, tc := range tests {
		if got := ls(tc.id); !slices.Equal(got, tc.want) {
			t.Errorf("%s: got %v want %v", tc.id, got, tc.want)
		}
	}
	// Each book is stamped with DownloadVersion, so none is stale.
	if stale, err := StaleFiveTools(dir); err != nil || len(stale) != 0 {
		t.Errorf("stale after download: %v %v", stale, err)
	}
	if items := readDoc(t, filepath.Join(dir, "5etools-tce", "items.json"))["item"]; len(items) != 1 || items[0]["name"] != "Cloak" {
		t.Errorf("items not filtered: %v", items)
	}
	if base := readDoc(t, filepath.Join(dir, "5etools-tce", "items-base.json"))["baseitem"]; len(base) != 1 {
		t.Errorf("items-base must stay whole: %v", base)
	}
	// A variant belongs to the book named in inherits.source.
	if mv := readDoc(t, filepath.Join(dir, "5etools-xphb", "magicvariants.json"))["magicvariant"]; len(mv) != 1 || mv[0]["name"] != "+1 Weapon" {
		t.Errorf("variants not filtered on inherits.source: %v", mv)
	}
	// Subclass-only book keeps the parent class.
	wiz := readDoc(t, filepath.Join(dir, "5etools-tce", "class-wizard.json"))
	if len(wiz["class"]) != 1 || len(wiz["subclass"]) != 1 || len(wiz["subclassFeature"]) != 1 {
		t.Errorf("class file: %v", wiz)
	}
	// Copies of other books' records carry their targets, copies of copies too.
	for _, tc := range []struct {
		file, list string
		want       []string
	}{
		{"bestiary.json", "monster", []string{"Wolf", "Pup"}},
		{"items.json", "item", []string{"Cloak"}},
		{"races.json", "race", []string{"Half-Elf"}}, // a subrace's race
	} {
		data, err := os.ReadFile(filepath.Join(dir, "5etools-xphb", tc.file))
		if err != nil {
			t.Fatal(err)
		}
		var doc map[string]map[string][]ftNamed
		_ = json.Unmarshal(data, &doc)
		var got []string
		for _, r := range doc[copyTargetsKey][tc.list] {
			got = append(got, r.Name)
		}
		if !slices.Equal(got, tc.want) {
			t.Errorf("%s targets %v, want %v", tc.file, got, tc.want)
		}
	}
	_, entries, skipped, err := parseFiveToolsBookReport("k", readBook(t, filepath.Join(dir, "5etools-xphb")))
	if err != nil || skipped != 0 {
		t.Fatalf("parse: skipped %d, %v", skipped, err)
	}
	got := map[string]Entry{}
	for _, e := range entries {
		got[string(e.Kind)+"/"+e.Slug] = e
	}
	for _, tc := range []struct {
		key  string
		want bool
	}{
		{"monster/dire-wolf", true}, {"item/wand-cloak", true},
		{"monster/wolf", false}, {"monster/pup", false}, {"monster/bear", false}, {"item/cloak", false},
	} {
		if _, ok := got[tc.key]; ok != tc.want {
			t.Errorf("%s imported %v, want %v", tc.key, ok, tc.want)
		}
	}
	if d := got["monster/dire-wolf"].Data.(MonsterData); d.Size == "" {
		t.Errorf("dire wolf did not inherit the pup's size: %+v", d)
	}
	if d := got["item/wand-cloak"].Data.(ItemData); d.Rarity == "" {
		t.Errorf("wand cloak did not inherit the cloak's rarity: %+v", d)
	}
	// A class key that escapes the book folder is skipped.
	if _, err := os.Stat(filepath.Join(dir, "evil.json")); !os.IsNotExist(err) {
		t.Errorf("bad class key wrote outside the book: %v", err)
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

// A subclass shipped for both the 2014 and the 2024 class keeps the 2024 copy
// in either order, and a feature referenced from another book comes bundled.
// TestStaleFiveTools: books written before DownloadVersion (or unstamped,
// from before the stamp) are listed, current ones are not.
func TestStaleFiveTools(t *testing.T) {
	dir := t.TempDir()
	var want []string
	for _, tc := range []struct {
		id    string
		ver   int
		stale bool
	}{{"xge", 0, true}, {"tce", DownloadVersion - 1, true}, {"xphb", DownloadVersion, false}} {
		if tc.stale {
			want = append(want, tc.id)
		}
		body, _ := json.Marshal(Book{ID: tc.id, Version: tc.ver})
		if err := os.MkdirAll(filepath.Join(dir, FiveToolsPrefix+tc.id), 0o750); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, FiveToolsPrefix+tc.id, bookFile), body, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	stale, err := StaleFiveTools(dir)
	if err != nil {
		t.Fatal(err)
	}
	var ids []string
	for _, b := range stale {
		ids = append(ids, b.ID)
	}
	slices.Sort(ids)
	if slices.Sort(want); !slices.Equal(ids, want) {
		t.Errorf("stale: %v, want %v", ids, want)
	}
}

func TestDownloadFiveToolsSubclassTwinsAndForeignFeatures(t *testing.T) {
	phb := `{"name":"Forge Domain","shortName":"Forge","className":"Cleric","classSource":"PHB","source":"XGE","subclassFeatures":["Forge Rite|Cleric|PHB|Forge|XGE|1"]}`
	xphb := `{"name":"Forge Domain","shortName":"Forge","className":"Cleric","classSource":"XPHB","source":"XGE","subclassFeatures":["Forge Rite|Cleric|XPHB|Forge|XGE|3","Blessed Strikes|Cleric|XPHB|Forge|XGE|8|TCE"]}`
	features := `"subclassFeature":[
		{"name":"Forge Rite","className":"Cleric","classSource":"PHB","subclassShortName":"Forge","subclassSource":"XGE","level":1,"source":"XGE","entries":["Early rite."]},
		{"name":"Forge Rite","className":"Cleric","classSource":"XPHB","subclassShortName":"Forge","subclassSource":"XGE","level":3,"source":"XGE","entries":["Late rite."]},
		{"name":"Blessed Strikes","className":"Cleric","classSource":"XPHB","subclassShortName":"Forge","subclassSource":"XGE","level":8,"source":"TCE","entries":["Strikes deal extra damage."]}]`
	for _, order := range [][]string{{phb, xphb}, {xphb, phb}} {
		files := maps.Clone(fakeMirror)
		files["class/index.json"] = `{"cleric":"class-cleric.json"}`
		files["class/class-cleric.json"] = `{"class":[{"name":"Cleric","source":"PHB"}],"subclass":[` + strings.Join(order, ",") + `],` + features + `}`
		serveMirror(t, files, false)
		dir := t.TempDir()
		if err := DownloadFiveTools(context.Background(), dir, []Book{{ID: "XGE", Name: "Xanathar", Group: GroupSupplement}}, nil); err != nil {
			t.Fatal(err)
		}
		doc := readDoc(t, filepath.Join(dir, "5etools-xge", "class-cleric.json"))
		if len(doc["subclassFeature"]) != 3 {
			t.Fatalf("bundled features: %v", doc["subclassFeature"])
		}
		_, entries, err := parseFiveToolsBook("5etools-xge", readBook(t, filepath.Join(dir, "5etools-xge")))
		if err != nil {
			t.Fatal(err)
		}
		var subclasses []ClassData
		for _, e := range entries {
			if e.Kind == KindSubclass {
				subclasses = append(subclasses, e.Data.(ClassData))
			}
			if e.Slug == "blessed-strikes" || e.Kind == KindClass {
				t.Errorf("bundled record became an entry: %s/%s", e.Kind, e.Slug)
			}
		}
		if len(subclasses) != 1 {
			t.Fatalf("%d subclasses, want 1", len(subclasses))
		}
		rite, _ := feature(subclasses[0], "Forge Rite")
		strikes, _ := feature(subclasses[0], "Blessed Strikes")
		if !slices.Equal(rite.Levels, []int{3}) || !strings.Contains(rite.Desc, "Late rite.") {
			t.Errorf("Forge Rite: %+v, want the level 3 copy", rite)
		}
		if !strings.Contains(strikes.Desc, "extra damage") {
			t.Errorf("Blessed Strikes has no text: %+v", strikes)
		}
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
	if err := DownloadFiveTools(context.Background(), dir, []Book{{ID: "XPHB"}}, nil); err == nil {
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

// TestDownloadLinksReprints: a tag to a 2014 record with a reprint points at
// the reprint, keeping its words; other tags are left alone.
func TestDownloadLinksReprints(t *testing.T) {
	files := maps.Clone(fakeMirror)
	files["bestiary/bestiary-mm.json"] = `{"monster":[{"name":"Goblin","source":"MM","reprintedAs":["Goblin Warrior|XMM"]},{"name":"Bear","source":"MM"}]}`
	files["bestiary/bestiary-xphb.json"] = `{"monster":[{"name":"Goblin Warrior","source":"XPHB","entries":["{@creature goblin|mm} {@creature Goblin|MM|goblins} {@creature bear|mm} {@creature ghost|zzz}"]}]}`
	files["spells/spells-xphb.json"] = `{"spell":[{"name":"Fireball","source":"XPHB","entries":["{@spell branding smite} {@spell fireball|xphb}"]}]}`
	files["items.json"] = `{"item":[{"name":"Abacus","source":"PHB","reprintedAs":[{"uid":"Abacus|XPHB"}]},{"name":"Wand","source":"XPHB","entries":["{@item abacus|phb}"]}]}`
	serveMirror(t, files, false)
	dir := t.TempDir()
	if err := DownloadFiveTools(context.Background(), dir, []Book{{ID: "XPHB", Name: "PHB", Group: GroupCore}}, nil); err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ file, list, want string }{
		{"bestiary.json", "monster", "{@creature Goblin Warrior|XMM|goblin} {@creature Goblin Warrior|XMM|goblins} {@creature bear|mm} {@creature ghost|zzz}"},
		{"spells.json", "spell", "{@spell branding smite} {@spell fireball|xphb}"},
		{"items.json", "item", "{@item Abacus|XPHB|abacus}"},
	} {
		recs := readDoc(t, filepath.Join(dir, "5etools-xphb", tc.file))[tc.list]
		got, _ := recs[len(recs)-1]["entries"].([]any)
		if len(got) != 1 || got[0] != tc.want {
			t.Errorf("%s: got %v, want %q", tc.file, got, tc.want)
		}
	}
}
