package web

import (
	"context"
	"database/sql"
	"errors"
	"slices"
	"strings"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/web/views"
)

// masteryEnv is the compendium weapon with a mastery named name, as the
// compendium spells it ("" when none), and the sheet's mastery limit.
func (a *App) masteryEnv(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet, campaignID int64, name string) (string, int, error) {
	weapons, err := a.masteryWeapons(ctx, campaignID)
	if err != nil {
		return "", 0, err
	}
	limit, err := masteryLimit(ctx, q, sheet, "")
	if i := slices.IndexFunc(weapons, func(w compendium.NameDesc) bool { return strings.EqualFold(w.Name, name) }); i >= 0 {
		return weapons[i].Name, limit, err
	}
	return "", limit, err
}

// masteryWeapons are the compendium weapons with a mastery property a
// player in campaignID may see: Name, and the mastery as Desc.
func (a *App) masteryWeapons(ctx context.Context, campaignID int64) ([]compendium.NameDesc, error) {
	rows, err := a.Store.ListMasteryWeapons(ctx, campaignID)
	out := make([]compendium.NameDesc, len(rows))
	for i, r := range rows {
		out[i] = compendium.NameDesc{Name: r.Name, Desc: r.Mastery}
	}
	return out, err
}

// masteryLimit is how many weapon kinds the sheet's classes and feats let it
// master, leaving out the class skip (a level-up counts that one itself).
func masteryLimit(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet, skip string) (int, error) {
	classes := map[string]dnd5e.ClassInfo{}
	for _, cl := range sheet.Classes {
		if cl.Slug == skip {
			continue
		}
		info, _, err := classEntry(ctx, q, cl.Slug, cl.Source)
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		if err != nil {
			return 0, err
		}
		classes[cl.Slug] = info
	}
	n := sheet.MasteryLimit(classes)
	for _, name := range sheet.Feats {
		key := featKey(name)
		row, _, err := entryFor(ctx, q, enum.KindFeat, key, sheet.FeatSources[key])
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		if err != nil {
			return 0, err
		}
		d, err := compendium.Decode[compendium.TextData](enum.KindFeat, row.DataJson)
		if err != nil {
			return 0, err
		}
		n += featMasteries(d)
	}
	return n, nil
}

// masteryGroups ask for the weapon kinds a level in info (class level
// level) adds to what the other classes let the sheet master, up to what is
// still unchosen.
func (a *App) masteryGroups(ctx context.Context, sh dnd5e.Sheet, info dnd5e.ClassInfo, level int, campaignID int64) ([]choiceGroup, error) {
	others, err := masteryLimit(ctx, a.Store.Queries, sh, info.Slug)
	if err != nil {
		return nil, err
	}
	before, after := others+info.Masteries(level-1), others+info.Masteries(level)
	return a.masteryChoice(ctx, sh, views.FieldMastery, dnd5e.MasteryColumn, min(after-before, after-len(sh.Masteries)), campaignID)
}

// masteryChoice asks for n weapon kinds the sheet does not master yet; none
// when n is 0 or no kind is left.
func (a *App) masteryChoice(ctx context.Context, sh dnd5e.Sheet, key, title string, n int, campaignID int64) ([]choiceGroup, error) {
	if n <= 0 {
		return nil, nil
	}
	weapons, err := a.masteryWeapons(ctx, campaignID)
	if err != nil {
		return nil, err
	}
	var opts []compendium.NameDesc
	for _, w := range weapons {
		if !sh.Masters(w.Name) {
			opts = append(opts, compendium.NameDesc{Name: w.Name, Desc: "Mastery: " + w.Desc})
		}
	}
	if len(opts) == 0 {
		return nil, nil
	}
	return []choiceGroup{{views.ChoiceGroup{Key: key, Title: title, Count: n, Options: opts}, func(s *dnd5e.Sheet, picked []string) {
		for _, p := range picked {
			if !s.Masters(p) { // another device may have chosen it since the page loaded
				s.Masteries = append(s.Masteries, p)
			}
		}
	}, nil}}, nil
}
