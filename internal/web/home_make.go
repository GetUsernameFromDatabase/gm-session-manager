package web

// The player app's ways to make another character on this computer: a test
// copy, a fallback build of a table character, a new character.

import (
	"database/sql"
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/export"
	"gm-session-manager/internal/route"
)

// Name suffixes of what is made here.
const (
	copySuffix     = " (copy)"
	fallbackSuffix = " (fallback)"
)

// homeNewForm is the "New character" dialog: a name, and the table character
// whose campaign and sources the new one starts with.
type homeNewForm struct {
	Name string `form:"name" validate:"required,max=100"`
	From int64  `form:"from" validate:"required,min=1"`
}

// loginAs ends this device's session and logs it in as ch.
func (a *App) loginAs(c *gin.Context, ch sqlc.Character) error {
	if err := a.dropSession(c); err != nil {
		return err
	}
	return a.login(c, ch)
}

// homeMakeCopy makes a test copy: everything copied into a campaign of its own
// (cloned from the original's), with no table link.
func (a *App) homeMakeCopy(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	ch, err := a.Store.GetCharacter(ctx, id)
	if err != nil {
		return err
	}
	if _, err := a.Store.GetFallbackOwner(ctx, nullID(id)); err == nil {
		return userf("A fallback is not copied; copy its character.")
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}
	var th export.TakeHome
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		var err error
		if th, _, err = buildTakeHome(ctx, q, ch); err != nil {
			return err
		}
		return a.checkOpenNotes(ctx, q, ch.ID)
	})
	if err != nil {
		return err
	}
	export.DropMissing(&th, a.attachmentsDir())

	// Files first, as in homeImport: the copy gets files of its own.
	var saved []string
	defer func() {
		for _, p := range saved {
			_ = os.Remove(filepath.Join(a.attachmentsDir(), p))
		}
	}()
	lookFiles, blockFiles, err := a.saveHomeFiles(th, func(zipName string) ([]byte, error) {
		return os.ReadFile(filepath.Join(a.attachmentsDir(), path.Base(zipName))) //nolint:gosec // our random file name, kept inside attachmentsDir by Base
	}, &saved)
	if err != nil {
		return err
	}
	var cp sqlc.Character
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		campaignID, err := cloneHomeCampaign(ctx, q, ch.CampaignID)
		if err != nil {
			return err
		}
		cf := th.Character
		cf.Name = homeName(cf.Name, copySuffix)
		// th.Homebrew stays with the original; the cloned links already show it.
		maps := newHomeIDMaps()
		cp, err = createHomeCharacter(ctx, q, campaignID, cf, th, lookFiles, blockFiles, &maps, nil)
		return err
	})
	if err != nil {
		return err
	}
	saved = nil // rows own the files now
	if err := a.loginAs(c, cp); err != nil {
		return err
	}
	redirect(c, route.Path(route.Character, cp.ID))
	return nil
}

// homeFallback prepares a fallback build of a table character: a separate
// character in the same home campaign with the sheet and level-up history only.
func (a *App) homeFallback(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	var fb sqlc.Character
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		hc, err := q.GetHomeCopy(ctx, id)
		if err != nil {
			return err
		}
		ch, err := q.GetCharacter(ctx, id)
		if err != nil {
			return err
		}
		if hc.FallbackCharacterID.Valid {
			if _, err := q.GetCharacter(ctx, hc.FallbackCharacterID.Int64); err == nil {
				return userf("%s already has a fallback build.", ch.Name)
			} else if !errors.Is(err, sql.ErrNoRows) {
				return err
			}
		}
		levelUps, err := q.ListLevelUps(ctx, id)
		if err != nil {
			return err
		}
		fb, err = q.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{
			CampaignID: ch.CampaignID, Name: homeName(ch.Name, fallbackSuffix), System: ch.System,
			Hp: ch.Hp, MaxHp: ch.MaxHp, TempHp: ch.TempHp, SheetJson: ch.SheetJson,
		})
		if err != nil {
			return err
		}
		for _, l := range levelUps {
			if _, err := q.CreateLevelUpAt(ctx, sqlc.CreateLevelUpAtParams{CharacterID: fb.ID, MaxHp: l.MaxHp, SheetJson: l.SheetJson, CreatedAt: l.CreatedAt}); err != nil {
				return err
			}
		}
		return q.SetHomeCopyFallback(ctx, sqlc.SetHomeCopyFallbackParams{FallbackCharacterID: nullID(fb.ID), CharacterID: id})
	})
	if err != nil {
		return err
	}
	if err := a.loginAs(c, fb); err != nil {
		return err
	}
	redirect(c, route.Path(route.Character, fb.ID))
	return nil
}

// homeNew makes a blank character that starts in the wizard, in a campaign
// cloned from a table character's.
func (a *App) homeNew(c *gin.Context) error {
	var form homeNewForm
	if err := c.ShouldBind(&form); err != nil {
		return userf("Give the character a name (up to 100 characters) and choose a table character.")
	}
	name := strings.TrimSpace(form.Name)
	if name == "" {
		return userf("Give the character a name (up to 100 characters) and choose a table character.")
	}
	ctx := c.Request.Context()
	if _, err := a.Store.GetHomeCopy(ctx, form.From); err != nil {
		return err
	}
	from, err := a.Store.GetCharacter(ctx, form.From)
	if err != nil {
		return err
	}
	sheet, err := newSheetJSON()
	if err != nil {
		return err
	}
	var ch sqlc.Character
	err = a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		campaignID, err := cloneHomeCampaign(ctx, q, from.CampaignID)
		if err != nil {
			return err
		}
		ch, err = q.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{CampaignID: campaignID, Name: name, System: gameSystem, SheetJson: sheet})
		return err
	})
	if err != nil {
		return err
	}
	if err := a.loginAs(c, ch); err != nil {
		return err
	}
	redirect(c, wizardPath(ch.ID, wizardSteps[0].Key))
	return nil
}
