package web

import (
	"context"
	"database/sql"
	"errors"

	"gm-session-manager/compendium"
	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
)

// UpgradeSheets catches saved sheets up with what the effects parser has
// learned since (dnd5e.EffectsVersion), against their classes' compendium
// tables. It runs at startup, before any client can hold a sheet, and
// returns how many sheets changed.
func UpgradeSheets(ctx context.Context, store *db.Store) (int, error) {
	changed := 0
	err := store.InTx(ctx, func(q *sqlc.Queries) error {
		rows, err := q.ListCharacterSheets(ctx)
		if err != nil {
			return err
		}
		classes := map[string]*dnd5e.ClassInfo{}
		for _, row := range rows {
			sheet, err := dnd5e.ParseSheet(row.SheetJson)
			if err != nil || sheet.EffectsVersion >= dnd5e.EffectsVersion {
				continue // an unreadable sheet stays as it is
			}
			for _, cl := range sheet.Classes {
				info, err := upgradeClass(ctx, q, classes, cl.Slug)
				if err != nil {
					return err
				}
				if info != nil && sheet.CatchUpUses(*info) {
					changed++
				}
			}
			sheet.EffectsVersion = dnd5e.EffectsVersion
			raw, err := sheet.JSON()
			if err != nil {
				return err
			}
			if err := q.UpdateCharacterSheet(ctx, sqlc.UpdateCharacterSheetParams{ID: row.ID, SheetJson: raw}); err != nil {
				return err
			}
		}
		return nil
	})
	return changed, err
}

// upgradeClass looks a class up once per upgrade; nil when the compendium
// no longer has it.
func upgradeClass(ctx context.Context, q *sqlc.Queries, cache map[string]*dnd5e.ClassInfo, slug string) (*dnd5e.ClassInfo, error) {
	if info, ok := cache[slug]; ok {
		return info, nil
	}
	row, err := q.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindClass), Slug: slug})
	if errors.Is(err, sql.ErrNoRows) {
		cache[slug] = nil
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	info, err := classInfo(row)
	if err != nil {
		return nil, err
	}
	cache[slug] = &info
	return &info, nil
}
