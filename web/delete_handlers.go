package web

import (
	"encoding/json"
	"io"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/export"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

// maxCharacterFileBytes caps an imported character file.
const maxCharacterFileBytes = 2 << 20

// removeCharacter deletes a character for its player or the GM. It is only
// hidden: its devices are logged out and it leaves combat, and the GM can
// restore it with its notes and images.
func (a *App) removeCharacter(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	if !perm.CanDeleteCharacter(w.viewer, ch.ID) {
		a.fail(c, http.StatusForbidden, "Only the player or the GM deletes a character.")
		return
	}
	var encounters []int64
	err := a.Store.InTx(c.Request.Context(), func(q *sqlc.Queries) error {
		var err error
		if encounters, err = q.ListEncounterIDsByCharacter(c.Request.Context(), nullID(ch.ID)); err != nil {
			return err
		}
		if err := q.SoftDeleteCharacter(c.Request.Context(), ch.ID); err != nil {
			return err
		}
		if err := q.DeleteCharacterSessions(c.Request.Context(), nullID(ch.ID)); err != nil {
			return err
		}
		return q.DeleteCharacterCombatants(c.Request.Context(), nullID(ch.ID))
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.CloseCharacter(ch.ID)
	topics := []hub.Topic{{Kind: hub.TopicCampaign, ID: ch.CampaignID}, {Kind: hub.TopicNotes, ID: ch.CampaignID}}
	for _, e := range encounters {
		topics = append(topics, hub.Topic{Kind: hub.TopicEncounter, ID: e})
	}
	a.Hub.Publish(ch.CampaignID, topics...)
	if w.viewer.IsGM {
		a.render(c, http.StatusOK, views.Flash("Deleted "+ch.Name+". You can restore it below.", false))
		return
	}
	redirect(c, PathHome)
}

// deletedCharacter loads a deleted character for the GM.
func (a *App) deletedCharacter(c *gin.Context) (sqlc.Character, bool) {
	id, ok := a.idParam(c, "character")
	if !ok {
		return sqlc.Character{}, false
	}
	ch, err := a.Store.GetCharacterAny(c.Request.Context(), id)
	return ch, a.found(c, err)
}

func (a *App) restoreCharacter(c *gin.Context) {
	ch, ok := a.deletedCharacter(c)
	if !ok {
		return
	}
	if err := a.Store.RestoreCharacter(c.Request.Context(), ch.ID); err != nil {
		a.serverError(c, err)
		return
	}
	a.changed(c, ch.CampaignID, hub.Topic{Kind: hub.TopicCampaign, ID: ch.CampaignID}, hub.Topic{Kind: hub.TopicNotes, ID: ch.CampaignID})
}

// downloadCharacter is a character's file, in the shape the import reads.
func (a *App) downloadCharacter(c *gin.Context) {
	ch, ok := a.deletedCharacter(c)
	if !ok {
		return
	}
	body, err := export.NewCharacterFile(ch).JSON()
	if err != nil {
		a.serverError(c, err)
		return
	}
	c.Header("Content-Disposition", `attachment; filename="`+path.Base(export.CharacterFileName(ch))+`"`)
	c.Data(http.StatusOK, "application/json", body)
}

// importCharacter adds a character from a downloaded or exported file to
// the campaign, as a new character with no devices.
func (a *App) importCharacter(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), id)
	if !a.found(c, err) {
		return
	}
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxCharacterFileBytes+1<<20)
	fh, err := c.FormFile("file")
	if err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "Choose a character file (.json).")
		return
	}
	f, err := fh.Open()
	if err != nil {
		a.serverError(c, err)
		return
	}
	defer func() { _ = f.Close() }()
	body, err := io.ReadAll(io.LimitReader(f, maxCharacterFileBytes))
	if err != nil {
		a.serverError(c, err)
		return
	}
	cf, err := parseCharacterFile(body)
	if err != nil {
		a.fail(c, http.StatusUnprocessableEntity, "That is not a character file: "+err.Error())
		return
	}
	if _, err := a.Store.CreateCharacterWithSheet(c.Request.Context(), sqlc.CreateCharacterWithSheetParams{
		CampaignID: camp.ID, Name: cf.Name, System: cf.System, Hp: cf.HP, MaxHp: cf.MaxHP, TempHp: cf.TempHP, SheetJson: string(cf.Sheet),
	}); err != nil {
		a.serverError(c, err)
		return
	}
	a.changed(c, camp.ID, hub.Topic{Kind: hub.TopicCampaign, ID: camp.ID})
}

// parseCharacterFile reads and checks a character file: its fields and the
// whole character as a command would leave it.
func parseCharacterFile(body []byte) (export.CharacterFile, error) {
	var cf export.CharacterFile
	if err := json.Unmarshal(body, &cf); err != nil {
		return cf, err
	}
	if err := binding.Validator.ValidateStruct(&cf); err != nil {
		return cf, err
	}
	sheet, err := dnd5e.ParseSheet(string(cf.Sheet))
	if err != nil {
		return cf, err
	}
	ch := dnd5e.Character{Name: cf.Name, HP: int(cf.HP), MaxHP: int(cf.MaxHP), TempHP: int(cf.TempHP), Sheet: sheet}
	if err := ch.Validate(); err != nil {
		return cf, err
	}
	// Stored as parsed, so the row holds exactly what the app understands.
	normalized, err := json.Marshal(sheet)
	cf.Sheet = normalized
	return cf, err
}
