package views

import (
	"context"
	"net/url"
	"slices"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"

	"github.com/a-h/templ"
)

// HabitatGridID is the grid's element, swapped alone on filtering.
const HabitatGridID = "habitat-grid"

// HabitatView is the GM's habitat grid: the monsters whose name contains
// Filter and the habitats each one has.
type HabitatView struct {
	Filter   string
	Monsters []sqlc.ListMonstersByNameRow
	Has      map[int64]map[string]bool
}

// anyPressed styles the Any button pressed while no box in its row is
// unticked, so it follows the GM ticking boxes one by one.
const anyPressed = "[tr:not(:has([data-slot=checkbox][data-unchecked]))_&]:border-primary [tr:not(:has([data-slot=checkbox][data-unchecked]))_&]:bg-primary [tr:not(:has([data-slot=checkbox][data-unchecked]))_&]:text-primary-foreground"

// everyHabitat reports whether has holds every habitat: the monster's habitat is Any.
func everyHabitat(has map[string]bool) bool {
	return !slices.ContainsFunc(compendium.Habitats, func(h string) bool { return !has[h] })
}

// habitatsURL opens the grid filtered to name.
func habitatsURL(ctx context.Context, name string) templ.SafeURL {
	return templ.SafeURL(CompendiumPath(ctx, "/habitats?"+url.Values{"q": {name}}.Encode()))
}
