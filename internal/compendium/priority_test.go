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
	if err := store.UpsertCompendiumSource(
		ctx,
		sqlc.UpsertCompendiumSourceParams{
			SourceKey: key,
			Path:      key,
			Hash:      "h",
			Edition:   edition,
			Priority:  DefaultRank(key, edition),
		},
	); err != nil {
		t.Fatal(err)
	}
	if err := store.UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{
		Kind:        "spell",
		Slug:        slug,
		Name:        slug,
		SourceKey:   key,
		DataJson:    `{"level":3}`,
		ContentHash: key,
	}); err != nil {
		t.Fatal(err)
	}
}

func winningSource(t *testing.T, store *db.Store, slug string) string {
	t.Helper()
	row, err := store.GetCompendiumEntryBySlug(
		t.Context(),
		sqlc.GetCompendiumEntryBySlugParams{Kind: "spell", Slug: slug},
	)
	if err != nil {
		t.Fatal(err)
	}
	// The sheets' batch lookup must pick the same copy.
	rows, err := store.ListCompendiumBySlugs(
		t.Context(),
		sqlc.ListCompendiumBySlugsParams{Kind: "spell", Slugs: []string{slug}},
	)
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
	if err := store.SetCompendiumSourcePriority(
		ctx,
		sqlc.SetCompendiumSourcePriorityParams{Priority: 7, SourceKey: "phb-2014"},
	); err != nil {
		t.Fatal(err)
	}
	addCopy(t, store, "phb-2014", "2014", "fireball")
	if got := winningSource(t, store, "fireball"); got != "phb-2014" {
		t.Errorf("after pinning phb-2014 first: got %s", got)
	}
	if err := store.SetCompendiumSourcePriority(
		ctx,
		sqlc.SetCompendiumSourcePriorityParams{Priority: 300, SourceKey: "phb-2014"},
	); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		order    []string
		wantSlug string
		wantErr  bool
	}{
		{name: "default winner", wantSlug: "xphb"},
		{
			name:     "srd above xphb",
			order:    []string{"homebrew", "srd-2024", "xphb", "phb-2014"},
			wantSlug: "srd-2024",
		},
		{
			name:     "srd below again",
			order:    []string{"homebrew", "xphb", "srd-2024", "phb-2014"},
			wantSlug: "xphb",
		},
		{
			name:     "srd above once more",
			order:    []string{"homebrew", "srd-2024", "xphb", "phb-2014"},
			wantSlug: "srd-2024",
		},
		{name: "unknown", order: []string{"nope"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.order != nil {
				if err := SetSourceOrder(ctx, store, tt.order); (err != nil) != tt.wantErr {
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

	rows, err := store.ListCompendiumByKind(
		ctx,
		sqlc.ListCompendiumByKindParams{Kind: "spell", PageSize: 50},
	)
	if err != nil || len(rows) != 1 || rows[0].SourceKey != "srd-2024" {
		t.Errorf("browse: %+v %v", rows, err)
	}
	// The kind badges count what browsing lists: three copies are one spell.
	if counts, err := store.CountCompendium(
		ctx,
		0,
	); err != nil || len(counts) != 1 ||
		counts[0].N != 1 {
		t.Errorf("counts: %+v %v, want 1 spell", counts, err)
	}
	copies, err := store.ListEntryCopies(
		ctx,
		sqlc.ListEntryCopiesParams{Kind: "spell", Slug: "fireball"},
	)
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

func TestSetSourceOrder(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	addCopy(t, store, "aaa", "2014", "x")
	addCopy(t, store, "bbb", "2024", "x")
	for _, tt := range []struct {
		name    string
		keys    []string
		wantErr bool
		winner  string
	}{
		{name: "aaa first", keys: []string{"aaa", "homebrew", "bbb"}, winner: "aaa"},
		{name: "bbb first", keys: []string{"bbb", "aaa", "homebrew"}, winner: "bbb"},
		{name: "missing key", keys: []string{"aaa", "homebrew"}, wantErr: true, winner: "bbb"},
		{name: "unknown key", keys: []string{"aaa", "homebrew", "nope"}, wantErr: true, winner: "bbb"},
		{name: "repeated key", keys: []string{"aaa", "aaa", "homebrew"}, wantErr: true, winner: "bbb"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetSourceOrder(ctx, store, tt.keys); (err != nil) != tt.wantErr {
				t.Fatalf("err %v", err)
			}
			if got := winningSource(t, store, "x"); got != tt.winner {
				t.Errorf("winner: got %s, want %s", got, tt.winner)
			}
		})
	}
}

func TestSourceOrderBeatsDate(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	addCopy(t, store, "aaa", "2014", "x")
	addCopy(t, store, "bbb", "2014", "x") // both 300
	if got := winningSource(t, store, "x"); got != "aaa" {
		t.Errorf("undated tie: %s, want the first key", got)
	}
	// The newer book wins a tie; a re-import fills its date.
	if err := store.UpsertCompendiumSource(
		ctx,
		sqlc.UpsertCompendiumSourceParams{
			SourceKey: "bbb",
			Path:      "bbb",
			Hash:      "h",
			Edition:   "2014",
			Published: "2020-11-17",
		},
	); err != nil {
		t.Fatal(err)
	}
	if got := winningSource(t, store, "x"); got != "bbb" {
		t.Errorf("dated tie: %s, want the newer bbb", got)
	}
	// The GM's order beats the date.
	if err := SetSourceOrder(ctx, store, []string{"homebrew", "aaa", "bbb"}); err != nil {
		t.Fatal(err)
	}
	if got := winningSource(t, store, "x"); got != "aaa" {
		t.Errorf("winner after reorder: %s", got)
	}
	// Lists (campaign_top) show the same copy.
	names, err := store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: "spell"})
	if err != nil {
		t.Fatal(err)
	}
	if row, err := store.GetCompendiumEntryBySlugSource(
		ctx,
		sqlc.GetCompendiumEntryBySlugSourceParams{Kind: "spell", Slug: "x", SourceKey: "aaa"},
	); err != nil || len(names) != 1 ||
		names[0].ID != row.ID {
		t.Errorf("campaign_top: %+v %v", names, err)
	}
}

// TestCountsHideOtherCampaignsHomebrew: a player's monster linked to no
// campaign counts for the GM (campaign 0), not in another campaign, in both
// the kind and the habitat counts.
func TestCountsHideOtherCampaignsHomebrew(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	row, err := store.CreateHomebrew(
		ctx,
		sqlc.CreateHomebrewParams{
			Kind:        "monster",
			Slug:        "mire-hag",
			Name:        "Mire Hag",
			DataJson:    `{}`,
			ContentHash: "h",
			Scoped:      true,
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.AddMonsterHabitat(
		ctx,
		sqlc.AddMonsterHabitatParams{EntryID: row.ID, Habitat: "Swamp"},
	); err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct {
		name     string
		campaign int64
		want     int64
	}{
		{name: "GM", campaign: 0, want: 1},
		{name: "another campaign", campaign: 7, want: 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var kinds, habitats int64
			counts, err := store.CountCompendium(ctx, tt.campaign)
			if err != nil {
				t.Fatal(err)
			}
			for _, c := range counts {
				kinds += c.N
			}
			hc, err := store.CountHabitats(ctx, tt.campaign)
			if err != nil {
				t.Fatal(err)
			}
			for _, c := range hc {
				habitats += c.N
			}
			if kinds != tt.want || habitats != tt.want {
				t.Errorf("kinds %d, habitats %d, want %d", kinds, habitats, tt.want)
			}
		})
	}
}

func TestScopedHomebrewNeverShadows(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	addCopy(t, store, "phb-2014", "2014", "fireball")
	if _, err := store.CreateHomebrew(ctx, sqlc.CreateHomebrewParams{
		Kind:        "spell",
		Slug:        "fireball",
		Name:        "Fireball",
		DataJson:    `{}`,
		ContentHash: "p",
		Scoped:      true,
	}); err != nil {
		t.Fatal(err)
	}
	if got := winningSource(t, store, "fireball"); got != "phb-2014" {
		t.Errorf("player homebrew shadowed the import: %s", got)
	}
	rows, err := store.ListCompendiumByKind(
		ctx,
		sqlc.ListCompendiumByKindParams{Kind: "spell", PageSize: 50},
	)
	if err != nil || len(rows) != 1 || rows[0].SourceKey != "phb-2014" {
		t.Errorf("browse: %+v %v", rows, err)
	}
}

func sourceKeysInOrder(t *testing.T, store *db.Store) []string {
	t.Helper()
	sources, err := store.ListCompendiumSources(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	keys := make([]string, len(sources))
	for i, s := range sources {
		keys[i] = s.SourceKey
	}
	return keys
}

// importRun stores sources the way one Import run does: the base is read once.
func importRun(t *testing.T, store *db.Store, srcs ...Source) {
	t.Helper()
	base, err := store.MaxSourcePriority(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	for _, s := range srcs {
		if _, err := storeSource(t.Context(), store, base, s, "p", "h", nil, nil); err != nil {
			t.Fatal(err)
		}
	}
}

func TestDefaultRank(t *testing.T) {
	for _, tt := range []struct {
		key, edition string
		want         int64
	}{
		{"homebrew", "", 0},
		{"srd-2024", "2024", 200},
		{"srd52", "", 200},
		{"xphb", "2024", 100},
		{"phb", "2014", 300},
		{"odd", "", 300},
	} {
		if got := DefaultRank(tt.key, tt.edition); got != tt.want {
			t.Errorf("%s/%s: got %d, want %d", tt.key, tt.edition, got, tt.want)
		}
	}
}

func TestNewSourcesAppendBelowTheGMsOrder(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	// A fresh database gives the plain group ranks.
	importRun(
		t,
		store,
		Source{Key: "srd-2024", Edition: "2024"},
		Source{Key: "xphb", Edition: "2024"},
	)
	for _, tt := range []struct {
		key  string
		want int64
	}{{"srd-2024", 200}, {"xphb", 100}} {
		got, err := store.GetCompendiumSource(ctx, tt.key)
		if err != nil || got.Priority != tt.want {
			t.Fatalf("%s: priority %d (%v), want %d", tt.key, got.Priority, err, tt.want)
		}
	}
	if err := SetSourceOrder(ctx, store, []string{"homebrew", "srd-2024", "xphb"}); err != nil {
		t.Fatal(err)
	}
	importRun(t, store,
		Source{Key: "old", Edition: "2014", Published: "2014-01-01"},
		Source{Key: "new2024", Edition: "2024"},
		Source{Key: "newer", Edition: "2014", Published: "2020-01-01"},
	)
	want := []string{"homebrew", "srd-2024", "xphb", "new2024", "newer", "old"}
	if got := sourceKeysInOrder(t, store); !slices.Equal(got, want) {
		t.Errorf("order %v, want %v", got, want)
	}
	// A re-import keeps the dragged priority.
	if err := SetSourceOrder(
		ctx,
		store,
		[]string{"homebrew", "old", "srd-2024", "xphb", "new2024", "newer"},
	); err != nil {
		t.Fatal(err)
	}
	importRun(t, store, Source{Key: "xphb", Edition: "2024"})
	if got := sourceKeysInOrder(t, store); got[2] != "srd-2024" || got[1] != "old" {
		t.Errorf("re-import moved a source: %v", got)
	}
}

func TestSortSources(t *testing.T) {
	store := newStore(t)
	ctx := t.Context()
	importRun(t, store,
		Source{Key: "undated", Edition: "2014"},
		Source{Key: "a-old", Edition: "2014", Published: "2014-01-01"},
		Source{Key: "b-new", Edition: "2014", Published: "2020-01-01"},
		Source{Key: "xphb", Edition: "2024", Published: "2024-09-17"},
		Source{Key: "srd52", Edition: "2024"},
	)
	// Homebrew sits at index 2 and keeps it.
	if err := SetSourceOrder(
		ctx,
		store,
		[]string{"undated", "a-old", "homebrew", "b-new", "srd52", "xphb"},
	); err != nil {
		t.Fatal(err)
	}
	if err := SortSources(ctx, store); err != nil {
		t.Fatal(err)
	}
	want := []string{"xphb", "srd52", "homebrew", "b-new", "a-old", "undated"}
	if got := sourceKeysInOrder(t, store); !slices.Equal(got, want) {
		t.Errorf("order %v, want %v", got, want)
	}
}
