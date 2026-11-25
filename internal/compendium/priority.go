package compendium

import (
	"context"
	"fmt"
	"slices"

	"gm-session-manager/internal/db"
	"gm-session-manager/internal/db/sqlc"
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
		for i, k := range keys {
			if err := q.SetCompendiumSourcePriority(ctx, sqlc.SetCompendiumSourcePriorityParams{Priority: int64(i+1) * 10, SourceKey: k}); err != nil {
				return err
			}
		}
		return nil
	})
}
