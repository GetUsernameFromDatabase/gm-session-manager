package web

import (
	"context"
	"slices"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/web/views"
)

// booksUsed lists the books the pins come from in first-seen order, with what
// is taken from each, flagged when the campaign does not allow the book or
// (player app) it was added at home; and the campaign's own books, in
// priority order. sources is every source in priority order; allowed is nil
// when the campaign allows all (no campaign).
func booksUsed(pins []dnd5e.Pin, sources []sqlc.CompendiumSource, allowed, homeKeys []string, noCampaign bool) ([]views.BookUse, []string) {
	edition := map[string]string{}
	for _, s := range sources {
		edition[s.SourceKey] = s.Edition
	}
	var used []views.BookUse
	for _, p := range pins {
		i := slices.IndexFunc(used, func(b views.BookUse) bool { return b.Key == p.Source })
		if i < 0 {
			atHome := slices.Contains(homeKeys, p.Source)
			used = append(used, views.BookUse{
				Key: p.Source, Label: views.SourceLabel(p.Source, edition[p.Source]),
				// A home book's own warning says more than "not in this campaign".
				NotAllowed: !atHome && !noCampaign && !slices.Contains(allowed, p.Source),
				AtHome:     atHome,
			})
			i = len(used) - 1
		}
		if u := p.Kind + " " + p.Name; !slices.Contains(used[i].Uses, u) {
			used[i].Uses = append(used[i].Uses, u)
		}
	}
	var campaign []string
	for _, s := range sources {
		// The player app's home campaign may allow home books too: list only the table's.
		if s.SourceKey != compendium.HomebrewSource && !slices.Contains(homeKeys, s.SourceKey) && (noCampaign || slices.Contains(allowed, s.SourceKey)) {
			campaign = append(campaign, views.SourceLabel(s.SourceKey, s.Edition))
		}
	}
	return used, campaign
}

// sheetBooks is booksUsed for a character's sheet (nil when it uses no books).
func (a *App) sheetBooks(ctx context.Context, sheet dnd5e.Sheet, campaignID int64) ([]views.BookUse, []string, error) {
	pins := sheet.Pins()
	if len(pins) == 0 {
		return nil, nil, nil
	}
	sources, err := a.Store.ListCompendiumSources(ctx)
	if err != nil {
		return nil, nil, err
	}
	var allowed, home []string
	if campaignID != 0 {
		if allowed, err = a.Store.ListCampaignSources(ctx, campaignID); err != nil {
			return nil, nil, err
		}
	}
	if a.Player {
		if home, err = a.Store.ListHomeSourceKeys(ctx); err != nil {
			return nil, nil, err
		}
	}
	used, campaign := booksUsed(pins, sources, allowed, home, campaignID == 0)
	return used, campaign, nil
}
