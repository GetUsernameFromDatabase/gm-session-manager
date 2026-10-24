package compendium

import (
	"database/sql"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
)

func newStore(t *testing.T) *db.Store {
	t.Helper()
	conn, err := db.Open(t.Context(), filepath.Join(t.TempDir(), "test.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = conn.Close() })
	return db.NewStore(conn)
}

// copySource copies testdata/srd-mini to dir/srd-2024 so tests can edit it.
func copySource(t *testing.T, dir string) string {
	t.Helper()
	dst := filepath.Join(dir, "srd-2024")
	if err := os.MkdirAll(dst, 0o750); err != nil {
		t.Fatal(err)
	}
	paths, _ := filepath.Glob("testdata/srd-mini/*.json")
	for _, p := range paths {
		b, err := os.ReadFile(p) //nolint:gosec // test fixture
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dst, filepath.Base(p)), b, 0o600); err != nil {
			t.Fatal(err)
		}
	}
	return dst
}

func importOnce(t *testing.T, store *db.Store, dir string) Report {
	t.Helper()
	reports, err := Import(t.Context(), store, dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(reports) != 1 || reports[0].Err != nil {
		t.Fatalf("reports %+v", reports)
	}
	return reports[0]
}

func TestImportIsIdempotent(t *testing.T) {
	store := newStore(t)
	dir := t.TempDir()
	src := copySource(t, dir)

	r := importOnce(t, store, dir)
	if r.Unchanged || r.Entries != 5 {
		t.Fatalf("first import %+v", r)
	}
	goblin, err := store.GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	if err != nil {
		t.Fatal(err)
	}
	source, err := store.GetCompendiumSource(t.Context(), goblin.SourceKey)
	if err != nil || source.License != "CC-BY-4.0" || source.Edition != "2024" || source.Attribution == "" {
		t.Errorf("source metadata %+v %v", source, err)
	}

	if r := importOnce(t, store, dir); !r.Unchanged {
		t.Errorf("second import %+v", r)
	}

	// Dropping a file removes its entries; the goblin keeps its id.
	if err := os.Remove(filepath.Join(src, "Spell.json")); err != nil {
		t.Fatal(err)
	}
	if r := importOnce(t, store, dir); r.Unchanged || r.Removed != 1 {
		t.Errorf("third import %+v", r)
	}
	again, err := store.GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "monster", Slug: "goblin-warrior"})
	if err != nil || again.ID != goblin.ID {
		t.Errorf("goblin id changed: %v %v", again.ID, err)
	}
}

func TestNormalizedShapes(t *testing.T) {
	store := newStore(t)
	dir := t.TempDir()
	copySource(t, dir)
	importOnce(t, store, dir)
	ctx := t.Context()

	get := func(kind Kind, slug string) any {
		t.Helper()
		row, err := store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(kind), Slug: slug})
		if err != nil {
			t.Fatalf("%s %s: %v", kind, slug, err)
		}
		d, err := DecodeData(kind, row.DataJson)
		if err != nil {
			t.Fatal(err)
		}
		return d
	}

	m := get(KindMonster, "goblin-warrior").(*MonsterData)
	if m.AC != 15 || m.HP != 10 || m.CR != "1/4" || m.Speed["walk"] != 30 || m.Abilities["dex"] != 15 || m.Skills["stealth"] != 6 {
		t.Errorf("monster %+v", m)
	}
	if len(m.Actions) != 2 || m.Actions[0].Name != "Scimitar" || m.Actions[1].Type != "bonus_action" {
		t.Errorf("actions %+v", m.Actions)
	}
	if m.Senses != "Darkvision 60 ft., Passive Perception 9" {
		t.Errorf("senses %q", m.Senses)
	}

	s := get(KindSpell, "fire-bolt").(*SpellData)
	if s.Level != 0 || s.Components != "V, S" || len(s.Classes) != 2 || s.Classes[1] != "wizard" {
		t.Errorf("spell %+v", s)
	}

	c := get(KindClass, "wizard").(*ClassData)
	if c.HitDie != 6 || c.Caster != "full" || c.Table["Cantrips"][4] != "4" || len(c.Features) != 2 {
		t.Errorf("class %+v", c)
	}
	if c.Features[1].Name != "Ability Score Improvement" || len(c.Features[1].Levels) != 2 {
		t.Errorf("features %+v", c.Features)
	}
	sub := get(KindSubclass, "wizard_evoker").(*ClassData)
	if sub.SubclassOf != "wizard" {
		t.Errorf("subclass %+v", sub)
	}
}

func TestSearch(t *testing.T) {
	store := newStore(t)
	dir := t.TempDir()
	copySource(t, dir)
	importOnce(t, store, dir)

	tests := []struct {
		query, kind         string
		wantNames, wantText []string
	}{
		{"gob", "", []string{"Goblin Warrior"}, nil},
		{"goblin", "spell", nil, nil},
		{"fire", "spell", []string{"Fire Bolt"}, nil},
		{"mote fire", "", nil, []string{"Fire Bolt"}},
		{"crawl", "condition", nil, []string{"Prone"}},
		{`" OR * NEAR(`, "", nil, nil},
		{"", "", nil, nil},
	}
	rowNames := func(rows []sqlc.SearchCompendiumRow) (got []string) {
		for _, r := range rows {
			got = append(got, r.Name)
		}
		return got
	}
	for _, tt := range tests {
		t.Run(tt.query+"/"+tt.kind, func(t *testing.T) {
			names, text, err := Search(t.Context(), store.Queries, tt.query, tt.kind, "", 0, 10)
			if err != nil {
				t.Fatal(err)
			}
			if got := rowNames(names); !slices.Equal(got, tt.wantNames) {
				t.Errorf("names: got %v, want %v", got, tt.wantNames)
			}
			if got := rowNames(text); !slices.Equal(got, tt.wantText) {
				t.Errorf("text: got %v, want %v", got, tt.wantText)
			}
		})
	}
}

func TestHomebrew(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	in := HomebrewInput{Kind: "monster", Name: "Bog Wight", Desc: "A drowned soul.", AC: 13, HP: 45, CR: "3"}
	id, err := SaveHomebrew(ctx, store.Queries, 0, in, HomebrewOwner{})
	if err != nil {
		t.Fatal(err)
	}
	row, err := store.GetCompendiumEntry(ctx, id)
	if err != nil || !row.IsHomebrew || row.SourceKey != HomebrewSource || row.Slug != "bog-wight" {
		t.Fatalf("row %+v %v", row, err)
	}
	if _, err := SaveHomebrew(ctx, store.Queries, 0, in, HomebrewOwner{}); err == nil {
		t.Error("duplicate homebrew accepted")
	}
	in.Desc = "A drowned soul that hates light."
	if _, err := SaveHomebrew(ctx, store.Queries, id, in, HomebrewOwner{}); err != nil {
		t.Fatal(err)
	}
	_, text, err := Search(ctx, store.Queries, "light", "monster", "", 0, 5)
	if err != nil || len(text) != 1 {
		t.Errorf("search after update: %v %v", text, err)
	}
	if _, err := (HomebrewInput{Kind: "vehicle", Name: "Cart"}).Entry(); err == nil {
		t.Error("unknown kind accepted")
	}
}

func TestImportRefusesHomebrewKey(t *testing.T) {
	store := newStore(t)
	dir := t.TempDir()
	src := copySource(t, dir)
	if err := os.Rename(src, filepath.Join(dir, HomebrewSource)); err != nil {
		t.Fatal(err)
	}
	reports, err := Import(t.Context(), store, dir)
	if err != nil || len(reports) != 1 || reports[0].Err == nil {
		t.Errorf("reports %+v %v", reports, err)
	}
}

// TestRealSource imports a full Open5e document directory when
// GSM_COMPENDIUM_DIR points at a folder containing one (for example
// data/compendium), to check the parser against the real data.
func TestRealSource(t *testing.T) {
	dir := os.Getenv("GSM_COMPENDIUM_DIR")
	if dir == "" {
		t.Skip("set GSM_COMPENDIUM_DIR to run")
	}
	store := newStore(t)
	reports, err := Import(t.Context(), store, dir)
	if err != nil {
		t.Fatal(err)
	}
	for _, r := range reports {
		if r.Err != nil {
			t.Errorf("%s: %v", r.Source, r.Err)
		}
		t.Logf("%s: %d entries", r.Source, r.Entries)
	}
	counts, err := store.CountCompendium(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range counts {
		t.Logf("%s: %d", c.Kind, c.N)
	}
}

func TestMonsterImageBinding(t *testing.T) {
	store := newStore(t)
	dir := t.TempDir()
	src := copySource(t, dir)
	importOnce(t, store, dir)
	ctx := t.Context()

	get := func(kind Kind, slug string) sqlc.CompendiumEntry {
		t.Helper()
		row, err := store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(kind), Slug: slug})
		if err != nil {
			t.Fatal(err)
		}
		return row
	}
	goblin := get(KindMonster, "goblin-warrior")
	if goblin.ImageUrl.String != MonsterImagePrefix+"goblin-warrior.png" {
		t.Fatalf("bound %q", goblin.ImageUrl.String)
	}
	if spell := get(KindSpell, "fire-bolt"); spell.ImageUrl.Valid {
		t.Errorf("spell bound %q", spell.ImageUrl.String)
	}

	// The GM's link, and a cleared one, survive a changed re-import.
	for _, want := range []string{"https://example.com/goblin.png", ""} {
		if _, err := store.SetCompendiumImage(ctx, sqlc.SetCompendiumImageParams{ImageUrl: sql.NullString{String: want, Valid: true}, ID: goblin.ID}); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(filepath.Join(src, "Spell.json")); err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
		importOnce(t, store, dir)
		if got := get(KindMonster, "goblin-warrior").ImageUrl; !got.Valid || got.String != want {
			t.Errorf("after re-import %+v, want %q", got, want)
		}
		copySource(t, dir)
	}
}
