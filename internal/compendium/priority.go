package compendium

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strings"

	"gm-session-manager/internal/db"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
)

// ErrNoSource means the key names no imported source.
var ErrNoSource = fmt.Errorf("unknown source")

// SetSourceOrder puts the sources in keys' order (priorities 10, 20, …). keys
// must name every source once: a stale page that missed an import or
// removal is refused with ErrNoSource rather than half applied.
func SetSourceOrder(ctx context.Context, store *db.Store, keys []string) error {
	return store.InTx(ctx, func(q *sqlc.Queries) error {
		sources, err := q.ListCompendiumSources(ctx)
		if err != nil {
			return err
		}
		have := make([]string, len(sources))
		for i, s := range sources {
			have[i] = s.SourceKey
		}
		if !slices.Equal(slices.Sorted(slices.Values(have)), slices.Sorted(slices.Values(keys))) {
			return ErrNoSource
		}
		return writeOrder(ctx, q, keys)
	})
}

// writeOrder numbers keys 10, 20, … in order.
func writeOrder(ctx context.Context, q *sqlc.Queries, keys []string) error {
	for i, k := range keys {
		if err := q.SetCompendiumSourcePriority(ctx, sqlc.SetCompendiumSourcePriorityParams{Priority: int64(i+1) * 10, SourceKey: k}); err != nil {
			return err
		}
	}
	return nil
}

// Group ranks of DefaultRank (lower wins); homebrew is 0 and never moves.
const (
	rank2024 = 100
	rankSRD  = 200
	rankRest = 300
)

// DefaultRank is a source's group rank: the 2024 books, then the SRDs, then the rest.
func DefaultRank(key, edition string) int64 {
	switch {
	case key == HomebrewSource:
		return 0
	case strings.HasPrefix(key, "srd"):
		return rankSRD
	case edition == dnd5e.Edition:
		return rank2024
	}
	return rankRest
}

// SortSources orders every source by the default rule: group rank, newest
// first (undated last), then key. Homebrew keeps its place in the list.
func SortSources(ctx context.Context, store *db.Store) error {
	return store.InTx(ctx, func(q *sqlc.Queries) error {
		sources, err := q.ListCompendiumSources(ctx)
		if err != nil {
			return err
		}
		keys := sortedKeys(sources)
		return writeOrder(ctx, q, keys)
	})
}

func undated(published string) int {
	if published == "" {
		return 1
	}
	return 0
}

// sortedKeys is the sources' keys in default order, homebrew at its current index.
func sortedKeys(sources []sqlc.CompendiumSource) []string {
	home := slices.IndexFunc(sources, func(s sqlc.CompendiumSource) bool { return s.SourceKey == HomebrewSource })
	rest := slices.DeleteFunc(slices.Clone(sources), func(s sqlc.CompendiumSource) bool { return s.SourceKey == HomebrewSource })
	slices.SortFunc(rest, func(a, b sqlc.CompendiumSource) int {
		return cmp.Or(
			cmp.Compare(DefaultRank(a.SourceKey, a.Edition), DefaultRank(b.SourceKey, b.Edition)),
			// '' sorts last, otherwise newest first.
			cmp.Compare(undated(a.Published), undated(b.Published)),
			cmp.Compare(b.Published, a.Published),
			cmp.Compare(a.SourceKey, b.SourceKey),
		)
	})
	keys := make([]string, len(rest))
	for i, s := range rest {
		keys[i] = s.SourceKey
	}
	if home >= 0 {
		keys = slices.Insert(keys, min(home, len(keys)), HomebrewSource)
	}
	return keys
}
