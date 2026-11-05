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

// MoveSource moves a source one place up (wins sooner) or down in the
// priority order. Ties are first renumbered to steps of 10, so the swap
// always changes the order. Moving past either end is a no-op.
func MoveSource(ctx context.Context, store *db.Store, key string, up bool) error {
	return store.InTx(ctx, func(q *sqlc.Queries) error {
		sources, err := q.ListCompendiumSources(ctx)
		if err != nil {
			return err
		}
		i := slices.IndexFunc(sources, func(s sqlc.CompendiumSource) bool { return s.SourceKey == key })
		if i < 0 {
			return ErrNoSource
		}
		j := i + 1
		if up {
			j = i - 1
		}
		if j < 0 || j >= len(sources) {
			return nil
		}
		// The list is already in (priority, key) order, so ties keep it.
		prio := make([]int64, len(sources))
		for n, s := range sources {
			prio[n] = s.Priority
		}
		if len(slices.Compact(slices.Clone(prio))) != len(prio) {
			for n := range prio {
				prio[n] = int64(n+1) * 10
			}
		}
		prio[i], prio[j] = prio[j], prio[i]
		for n, s := range sources {
			if prio[n] == s.Priority {
				continue
			}
			if err := q.SetCompendiumSourcePriority(ctx, sqlc.SetCompendiumSourcePriorityParams{Priority: prio[n], SourceKey: s.SourceKey}); err != nil {
				return err
			}
		}
		return nil
	})
}
