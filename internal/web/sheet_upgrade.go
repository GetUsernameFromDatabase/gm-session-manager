package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"slices"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
)

// UpgradeSheets catches saved sheets up with what the app has learned since
// (dnd5e.EffectsVersion), against their classes' and species' compendium
// entries. It runs at startup, before any client can hold a sheet, and
// returns how many sheets changed.
func UpgradeSheets(ctx context.Context, store *db.Store) (int, error) {
	changed := 0
	err := store.InTx(ctx, func(q *sqlc.Queries) error {
		rows, err := q.ListCharacterSheets(ctx)
		if err != nil {
			return err
		}
		classes := map[string]*dnd5e.ClassInfo{}
		items := map[string]string{}
		for _, row := range rows {
			sheet, err := dnd5e.ParseSheet(row.SheetJson)
			if err != nil || sheet.EffectsVersion >= dnd5e.EffectsVersion {
				continue // an unreadable sheet stays as it is
			}
			for _, cl := range sheet.Classes {
				info, err := upgradeClass(ctx, q, classes, cl)
				if err != nil {
					return err
				}
				if info != nil && sheet.CatchUpUses(*info) {
					changed++
				}
			}
			// Version 2: sheets made before creation set resistances. One with
			// some already was made since, or set by hand.
			if sheet.EffectsVersion < 2 && len(sheet.Resistances) == 0 {
				parts, err := upgradeSpecies(ctx, q, sheet)
				if err != nil {
					return err
				}
				if catchUpResistances(&sheet, parts) {
					changed++
				}
			}
			// Version 3: subclasses that cast before third casters, weapons added
			// before mastery. main runs the compendium import first, so both
			// read re-parsed entries.
			if sheet.EffectsVersion < 3 {
				casters, err := catchUpCasters(ctx, q, &sheet)
				if err != nil {
					return err
				}
				masteries, err := catchUpMasteries(ctx, q, items, &sheet)
				if err != nil {
					return err
				}
				if casters || masteries {
					changed++
				}
			}
			// Version 4: features older level-ups added once per level, speed
			// bonuses that scale with a table column, profs features grant.
			if sheet.EffectsVersion < 4 {
				fixed := sheet.DedupeFeatures()
				for _, cl := range sheet.Classes {
					info, err := upgradeClass(ctx, q, classes, cl)
					if err != nil {
						return err
					}
					if info != nil && sheet.ScaleSpeeds(*info, cl.Level) {
						fixed = true
					}
				}
				if sheet.CatchUpProfs() || fixed {
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

// upgradeSpecies is the sheet's species' traits; none when the compendium
// no longer has it.
func upgradeSpecies(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet) ([]compendium.NameDesc, error) {
	if sheet.Species == "" {
		return nil, nil
	}
	row, _, err := entryFor(ctx, q, compendium.KindSpecies, compendium.Slugify(sheet.Species), sheet.SpeciesSource)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	d, err := compendium.DecodeData(compendium.KindSpecies, row.DataJson)
	if err != nil {
		return nil, err
	}
	return d.(*compendium.TextData).Parts, nil
}

// catchUpResistances adds the resistances creation would have: those the
// species' traits name, and those of the lineage the sheet picked (its
// feature is "<trait> (<option>)", as lineageSpells reads it).
func catchUpResistances(s *dnd5e.Sheet, parts []compendium.NameDesc) bool {
	before := len(s.Resistances)
	addAll := func(text string) {
		for _, m := range reResistance.FindAllStringSubmatch(text, -1) {
			addResistance(s, m[1])
		}
	}
	for _, p := range parts {
		if p.Type != "" {
			continue
		}
		opts := choiceOptions(p)
		if len(opts) == 0 {
			addAll(p.Desc)
		}
		for _, o := range opts {
			if !slices.ContainsFunc(s.Features, func(f dnd5e.Feature) bool {
				return f.Source == s.Species && f.Name == fmt.Sprintf("%s (%s)", p.Name, o.Name)
			}) {
				continue
			}
			addAll(o.Desc)
			if resistsPick(parts, p) {
				addResistance(s, o.Desc)
			}
		}
	}
	return len(s.Resistances) > before
}

// catchUpCasters gives a class without spellcasting its subclass's, as
// setSubclass does now (Eldritch Knight), and its slots.
func catchUpCasters(ctx context.Context, q *sqlc.Queries, s *dnd5e.Sheet) (bool, error) {
	changed := false
	for i, cl := range s.Classes {
		if cl.Subclass == "" || (cl.Caster != "" && cl.Caster != dnd5e.CasterNone) {
			continue
		}
		slug, pin, err := subclassRef(ctx, q, cl.Slug, cl, cl.Subclass)
		if err != nil {
			return false, err
		}
		if slug == "" {
			continue
		}
		row, _, err := entryFor(ctx, q, compendium.KindSubclass, slug, pin)
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		if err != nil {
			return false, err
		}
		info := dnd5e.ClassInfo{Caster: cl.Caster}
		if err := setSubclass(&info, row); err != nil {
			return false, err
		}
		if info.Caster != cl.Caster {
			s.Classes[i].Caster, changed = info.Caster, true
		}
	}
	if changed {
		s.RecomputeSlots()
	}
	return changed, nil
}

// catchUpMasteries gives weapons without a mastery their compendium
// weapon's, looking each kind up once per upgrade (items).
func catchUpMasteries(ctx context.Context, q *sqlc.Queries, items map[string]string, s *dnd5e.Sheet) (bool, error) {
	changed := false
	for i, w := range s.Weapons {
		if w.Mastery != "" {
			continue
		}
		slug := compendium.Slugify(dnd5e.WeaponKind(w.Name))
		mastery, ok := items[slug]
		if !ok {
			row, _, err := entryFor(ctx, q, compendium.KindItem, slug, "")
			switch {
			case errors.Is(err, sql.ErrNoRows):
			case err != nil:
				return false, err
			default:
				d, err := compendium.DecodeData(compendium.KindItem, row.DataJson)
				if err != nil {
					return false, err
				}
				mastery = d.(*compendium.ItemData).Mastery
			}
			items[slug] = mastery
		}
		if mastery != "" {
			s.Weapons[i].Mastery, changed = mastery, true
		}
	}
	return changed, nil
}

// upgradeClass looks a class's pinned copy up once per upgrade; nil when
// the compendium no longer has it.
func upgradeClass(ctx context.Context, q *sqlc.Queries, cache map[string]*dnd5e.ClassInfo, cl dnd5e.ClassLevel) (*dnd5e.ClassInfo, error) {
	key := cl.Slug + "|" + cl.Source
	if info, ok := cache[key]; ok {
		return info, nil
	}
	info, _, err := classEntry(ctx, q, cl.Slug, cl.Source)
	if errors.Is(err, sql.ErrNoRows) {
		cache[key] = nil
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	cache[key] = &info
	return &info, nil
}
