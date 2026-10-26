package web

import (
	"context"
	"errors"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/web/views"
)

// habitatCellForm is one box of the grid. The oneof lists are
// compendium.Habitats (TestHabitatsMatchSchema keeps them in step).
type habitatCellForm struct {
	EntryID int64  `form:"entry_id" validate:"required,gt=0"`
	Habitat string `form:"habitat" validate:"required,oneof=Arctic Coastal Desert Forest Grassland Hill Mountain Planar Swamp Underdark Underwater Urban"`
	On      bool   `form:"on"`
}

// habitatBatchForm adds or removes a habitat for every monster whose name
// contains Filter.
type habitatBatchForm struct {
	Filter  string `form:"q" validate:"max=200"`
	Habitat string `form:"habitat" validate:"required,oneof=Arctic Coastal Desert Forest Grassland Hill Mountain Planar Swamp Underdark Underwater Urban"`
	On      bool   `form:"on"`
}

// habitatAnyForm toggles every habitat of one monster.
type habitatAnyForm struct {
	EntryID int64 `form:"entry_id" validate:"required,gt=0"`
}

// errNotMonster rolls back habitatAny for an entry that is not a monster.
var errNotMonster = errors.New("not a monster")

// habitatView reads the grid for the monsters whose name contains filter.
func (a *App) habitatView(ctx context.Context, filter string) (views.HabitatView, error) {
	v := views.HabitatView{Filter: filter, Has: map[int64]map[string]bool{}}
	var err error
	if v.Monsters, err = a.Store.ListMonstersByName(ctx, filter); err != nil {
		return v, err
	}
	rows, err := a.Store.ListMonsterHabitats(ctx)
	for _, r := range rows {
		if v.Has[r.EntryID] == nil {
			v.Has[r.EntryID] = map[string]bool{}
		}
		v.Has[r.EntryID][r.Habitat] = true
	}
	return v, err
}

// habitatsPage is the GM's monsters × habitats grid; typing in the filter
// swaps the grid alone.
func (a *App) habitatsPage(c *gin.Context) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return
	}
	v, err := a.habitatView(c.Request.Context(), strings.TrimSpace(c.Query("q")))
	if err != nil {
		a.serverError(c, err)
		return
	}
	if isHTMX(c) && c.GetHeader("HX-Target") == views.HabitatGridID {
		a.render(c, http.StatusOK, views.HabitatGrid(v))
		return
	}
	a.render(c, http.StatusOK, views.HabitatsPage(a.layout(w, "Habitats", views.NavCompendium, nil, false), v))
}

// habitatCell saves one box of the grid.
func (a *App) habitatCell(c *gin.Context) {
	var f habitatCellForm
	if err := c.ShouldBind(&f); err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "Pick a monster and a known habitat.")
		return
	}
	ctx := c.Request.Context()
	if !f.On {
		if err := a.Store.RemoveMonsterHabitat(ctx, sqlc.RemoveMonsterHabitatParams{EntryID: f.EntryID, Habitat: f.Habitat}); err != nil {
			a.serverError(c, err)
			return
		}
		a.render(c, http.StatusOK, views.Flash("", false))
		return
	}
	n, err := a.Store.AddMonsterHabitat(ctx, sqlc.AddMonsterHabitatParams{EntryID: f.EntryID, Habitat: f.Habitat})
	if err != nil {
		a.serverError(c, err)
		return
	}
	if n == 0 {
		a.fail(c, http.StatusUnprocessableEntity, "Only monsters have habitats.")
		return
	}
	a.render(c, http.StatusOK, views.Flash("", false))
}

// habitatAny gives a monster every habitat ("Habitat: Any"), or clears them
// all when it has every one, then redraws its row.
func (a *App) habitatAny(c *gin.Context) {
	var f habitatAnyForm
	if err := c.ShouldBind(&f); err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "Pick a monster.")
		return
	}
	ctx := c.Request.Context()
	has := map[string]bool{}
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		current, err := q.ListEntryHabitats(ctx, f.EntryID)
		if err != nil {
			return err
		}
		every := !slices.ContainsFunc(compendium.Habitats, func(h string) bool { return !slices.Contains(current, h) })
		for _, h := range compendium.Habitats {
			if every {
				if err := q.RemoveMonsterHabitat(ctx, sqlc.RemoveMonsterHabitatParams{EntryID: f.EntryID, Habitat: h}); err != nil {
					return err
				}
				continue
			}
			n, err := q.AddMonsterHabitat(ctx, sqlc.AddMonsterHabitatParams{EntryID: f.EntryID, Habitat: h})
			if err != nil {
				return err
			}
			if n == 0 {
				return errNotMonster
			}
			has[h] = true
		}
		return nil
	})
	if errors.Is(err, errNotMonster) {
		a.fail(c, http.StatusUnprocessableEntity, "Only monsters have habitats.")
		return
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	e, err := a.Store.GetCompendiumEntry(ctx, f.EntryID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.HabitatRow(sqlc.ListMonstersByNameRow{ID: e.ID, Name: e.Name}, has))
}

// habitatBatch adds or removes a habitat for every match, then redraws the grid.
func (a *App) habitatBatch(c *gin.Context) {
	var f habitatBatchForm
	if err := c.ShouldBind(&f); err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "Pick a known habitat.")
		return
	}
	ctx := c.Request.Context()
	f.Filter = strings.TrimSpace(f.Filter)
	var err error
	if f.On {
		_, err = a.Store.AddHabitatByName(ctx, sqlc.AddHabitatByNameParams{Habitat: f.Habitat, Filter: f.Filter})
	} else {
		_, err = a.Store.RemoveHabitatByName(ctx, sqlc.RemoveHabitatByNameParams{Habitat: f.Habitat, Filter: f.Filter})
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	v, err := a.habitatView(ctx, f.Filter)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.HabitatGrid(v))
}
