package compendium

import (
	"slices"
	"testing"

	"gm-session-manager/internal/db"
	"gm-session-manager/internal/db/sqlc"
)

// addCopy stores a spell under a source, creating the source first.
func addCopy(t *testing.T, store *db.Store, key, edition, slug string) {
	t.Helper()
	ctx := t.Context()
	if err := store.UpsertCompendiumSource(ctx, sqlc.UpsertCompendiumSourceParams{SourceKey: key, Path: key, Hash: "h", Edition: edition}); err != nil {
		t.Fatal(err)
	}
	if err := store.UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{
		Kind: "spell", Slug: slug, Name: slug, SourceKey: key, DataJson: `{"level":3}`, ContentHash: key,
	}); err != nil {
		t.Fatal(err)
	}
}

func winningSource(t *testing.T, store *db.Store, slug string) string {
	t.Helper()
	row, err := store.GetCompendiumEntryBySlug(t.Context(), sqlc.GetCompendiumEntryBySlugParams{Kind: "spell", Slug: slug})
	if err != nil {
		t.Fatal(err)
	}
	// The sheets' batch lookup must pick the same copy.
	rows, err := store.ListCompendiumBySlugs(t.Context(), sqlc.ListCompendiumBySlugsParams{Kind: "spell", Slugs: []string{slug}})
	if err != nil || len(rows) != 1 || rows[0].ID != row.ID {
		t.Fatalf("ListCompendiumBySlugs %+v %v, want only entry %d", rows, err, row.ID)
	}
	return row.SourceKey
}

func TestSourcePriority(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	addCopy(t, store, "phb-2014", "2014", "fireball")
	addCopy(t, store, "srd-2024", "2024", "fireball")
	addCopy(t, store, "xphb", "2024", "fireball")

	prios := map[string]int64{}
	sources, err := store.ListCompendiumSources(ctx)
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range sources {
		prios[s.SourceKey] = s.Priority
	}
	for key, want := range map[string]int64{"homebrew": 0, "xphb": 100, "srd-2024": 200, "phb-2014": 300} {
		if prios[key] != want {
			t.Errorf("default priority of %s: got %d, want %d", key, prios[key], want)
		}
	}

	// A re-import keeps a reordered priority.
	if err := store.SetCompendiumSourcePriority(ctx, sqlc.SetCompendiumSourcePriorityParams{Priority: 7, SourceKey: "phb-2014"}); err != nil {
		t.Fatal(err)
	}
	addCopy(t, store, "phb-2014", "2014", "fireball")
	if got := winningSource(t, store, "fireball"); got != "phb-2014" {
		t.Errorf("after pinning phb-2014 first: got %s", got)
	}
	if err := store.SetCompendiumSourcePriority(ctx, sqlc.SetCompendiumSourcePriorityParams{Priority: 300, SourceKey: "phb-2014"}); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		move     string
		up       bool
		wantSlug string
		wantErr  bool
	}{
		{"default winner", "", false, "xphb", false},
		{"srd up past xphb", "srd-2024", true, "srd-2024", false},
		{"srd down again", "srd-2024", false, "xphb", false},
		{"srd up once more", "srd-2024", true, "srd-2024", false},
		{"unknown", "nope", true, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.move != "" {
				if err := MoveSource(ctx, store, tt.move, tt.up); (err != nil) != tt.wantErr {
					t.Fatalf("err %v", err)
				}
			}
			if tt.wantErr {
				return
			}
			if got := winningSource(t, store, "fireball"); got != tt.wantSlug {
				t.Errorf("winner: got %s, want %s", got, tt.wantSlug)
			}
		})
	}

	rows, err := store.ListCompendiumByKind(ctx, sqlc.ListCompendiumByKindParams{Kind: "spell", PageSize: 50})
	if err != nil || len(rows) != 1 || rows[0].SourceKey != "srd-2024" {
		t.Errorf("browse: %+v %v", rows, err)
	}
	copies, err := store.ListEntryCopies(ctx, sqlc.ListEntryCopiesParams{Kind: "spell", Slug: "fireball"})
	if err != nil {
		t.Fatal(err)
	}
	var keys []string
	for _, c := range copies {
		keys = append(keys, c.SourceKey)
	}
	if !slices.Equal(keys, []string{"srd-2024", "xphb", "phb-2014"}) {
		t.Errorf("copies: %v", keys)
	}
}

func TestMoveSourceBreaksTies(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	addCopy(t, store, "aaa", "2014", "x")
	addCopy(t, store, "bbb", "2014", "x") // both 300
	if got := winningSource(t, store, "x"); got != "aaa" {
		t.Errorf("undated tie: %s, want the first key", got)
	}
	// The newer book wins a tie; a re-import fills its date.
	if err := store.UpsertCompendiumSource(ctx, sqlc.UpsertCompendiumSourceParams{SourceKey: "bbb", Path: "bbb", Hash: "h", Edition: "2014", Published: "2020-11-17"}); err != nil {
		t.Fatal(err)
	}
	if got := winningSource(t, store, "x"); got != "bbb" {
		t.Errorf("dated tie: %s, want the newer bbb", got)
	}
	// The GM's order beats the date.
	if err := MoveSource(ctx, store, "aaa", true); err != nil {
		t.Fatal(err)
	}
	if got := winningSource(t, store, "x"); got != "aaa" {
		t.Errorf("winner after move: %s", got)
	}
	// Lists (compendium_top) show the same copy.
	names, err := store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: "spell"})
	if err != nil {
		t.Fatal(err)
	}
	if row, err := store.GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: "spell", Slug: "x", SourceKey: "aaa"}); err != nil || len(names) != 1 || names[0].ID != row.ID {
		t.Errorf("compendium_top: %+v %v", names, err)
	}
}

func TestScopedHomebrewNeverShadows(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	addCopy(t, store, "phb-2014", "2014", "fireball")
	if _, err := store.CreateHomebrew(ctx, sqlc.CreateHomebrewParams{
		Kind: "spell", Slug: "fireball", Name: "Fireball", DataJson: `{}`, ContentHash: "p", Scoped: true,
	}); err != nil {
		t.Fatal(err)
	}
	if got := winningSource(t, store, "fireball"); got != "phb-2014" {
		t.Errorf("player homebrew shadowed the import: %s", got)
	}
	rows, err := store.ListCompendiumByKind(ctx, sqlc.ListCompendiumByKindParams{Kind: "spell", PageSize: 50})
	if err != nil || len(rows) != 1 || rows[0].SourceKey != "phb-2014" {
		t.Errorf("browse: %+v %v", rows, err)
	}
}
