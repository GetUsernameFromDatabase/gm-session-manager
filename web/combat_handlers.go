package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"

	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/hub"
	"gm-session-manager/web/views"
)

// Encounter statuses (encounters.status).
const (
	statusPreparing = "preparing"
	statusActive    = "active"
	statusEnded     = "ended"
)

func (a *App) combatPage(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	w, camp, ok := a.campaignAccess(c, id)
	if !ok {
		return
	}
	l := a.layout(w, "Combat", views.NavCombat, &camp, true)
	ctx := c.Request.Context()
	v := views.CombatView{CampaignID: id, IsGM: w.isGM()}
	var err error
	want, _ := strconv.ParseInt(c.Query("encounter"), 10, 64)
	if !w.isGM() {
		v.Want = want
		if v.Encounters, v.Current, err = a.playerCombat(ctx, w.viewer, id, want); err != nil {
			a.serverError(c, err)
			return
		}
		a.render(c, http.StatusOK, views.CombatPage(l, v))
		return
	}

	if v.Encounters, err = a.Store.ListEncountersByCampaign(ctx, id); err != nil {
		a.serverError(c, err)
		return
	}
	for _, e := range v.Encounters {
		// The list is ordered active first, so the default is the live fight.
		if e.ID == want || (want == 0 && v.Current == nil) {
			ev, err := a.encounterView(ctx, w.viewer, e)
			if err != nil {
				a.serverError(c, err)
				return
			}
			v.Current = &ev
		}
	}
	a.render(c, http.StatusOK, views.CombatPage(l, v))
}

func (a *App) createEncounter(c *gin.Context) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	var form nameForm
	if !a.bindForm(c, &form) {
		return
	}
	enc, err := a.Store.CreateEncounter(c.Request.Context(), sqlc.CreateEncounterParams{CampaignID: id, Name: strings.TrimSpace(form.Name)})
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(id, hub.Topic{Kind: hub.TopicEncounter, ID: enc.ID})
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/campaigns/%d/combat?encounter=%d", id, enc.ID))
}

func (a *App) encounterParam(c *gin.Context) (sqlc.Encounter, bool) {
	id, ok := a.idParam(c, "encounter")
	if !ok {
		return sqlc.Encounter{}, false
	}
	enc, err := a.Store.GetEncounter(c.Request.Context(), id)
	return enc, a.found(c, err)
}

func (a *App) encounterChanged(c *gin.Context, enc sqlc.Encounter, extra ...hub.Topic) {
	a.encounterNotice(c, enc, "", extra...)
}

func (a *App) encounterNotice(c *gin.Context, enc sqlc.Encounter, msg string, extra ...hub.Topic) {
	a.changedNotice(c, enc.CampaignID, msg, append([]hub.Topic{{Kind: hub.TopicEncounter, ID: enc.ID}}, extra...)...)
}

// addCombatantForm is the add form's options; count applies to monsters.
type addCombatantForm struct {
	Count     int  `form:"count"      validate:"omitempty,min=1,max=20"`
	HPVisible bool `form:"hp_visible"`
}

// addCombatant adds a PC, an NPC entity, or one or more compendium monsters.
func (a *App) addCombatant(c *gin.Context) {
	enc, ok := a.encounterParam(c)
	if !ok {
		return
	}
	refID, ok := formInt(c, "id")
	if !ok {
		a.fail(c, http.StatusUnprocessableEntity, "Choose who to add.")
		return
	}
	var form addCombatantForm
	if !a.bindForm(c, &form) {
		return
	}
	ctx := c.Request.Context()
	p := sqlc.CreateCombatantParams{EncounterID: enc.ID, Ac: 10, HpVisible: form.HPVisible}
	count := 1
	switch c.PostForm("kind") {
	case "pc":
		ch, err := a.Store.GetCharacter(ctx, refID)
		if !a.found(c, err) || !a.sameCampaign(c, ch.CampaignID, enc.CampaignID) {
			return
		}
		sheet, err := dnd5e.ParseSheet(ch.SheetJson)
		if err != nil {
			a.serverError(c, err)
			return
		}
		p.CharacterID, p.Name, p.Ac = nullID(ch.ID), ch.Name, int64(sheet.ArmorClass())
	case "entity":
		e, err := a.Store.GetEntity(ctx, refID)
		if !a.found(c, err) || !a.sameCampaign(c, e.CampaignID, enc.CampaignID) {
			return
		}
		p.EntityID, p.Name = nullID(e.ID), e.Name
		if e.CompendiumRef.Valid {
			if err := a.monsterStats(ctx, e.CompendiumRef.Int64, &p); err != nil {
				a.serverError(c, err)
				return
			}
			p.Name = e.Name
		}
	case "monster":
		if err := a.monsterStats(ctx, refID, &p); errors.Is(err, sql.ErrNoRows) {
			a.fail(c, http.StatusNotFound, "No such monster.")
			return
		} else if err != nil {
			a.serverError(c, err)
			return
		}
		count = max(form.Count, 1)
	default:
		a.fail(c, http.StatusUnprocessableEntity, "Unknown combatant kind.")
		return
	}
	base := p.Name
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		for range count {
			if p.MonsterRef.Valid && !p.EntityID.Valid {
				var err error
				if p.Name, err = uniqueName(ctx, q, enc.ID, base); err != nil {
					return err
				}
			}
			if _, err := q.CreateCombatant(ctx, p); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.encounterChanged(c, enc)
}

func (a *App) sameCampaign(c *gin.Context, got, want int64) bool {
	if got != want {
		a.fail(c, http.StatusUnprocessableEntity, "That belongs to another campaign.")
		return false
	}
	return true
}

// monsterStats fills name, AC and HP from a compendium monster.
func (a *App) monsterStats(ctx context.Context, entryID int64, p *sqlc.CreateCombatantParams) error {
	row, err := a.Store.GetCompendiumEntry(ctx, entryID)
	if err != nil {
		return err
	}
	if row.Kind != string(compendium.KindMonster) {
		return sql.ErrNoRows
	}
	d, err := compendium.DecodeData(compendium.KindMonster, row.DataJson)
	if err != nil {
		return err
	}
	m := d.(*compendium.MonsterData)
	p.MonsterRef, p.Name, p.Ac, p.Hp, p.MaxHp = nullID(row.ID), row.Name, int64(m.AC), int64(m.HP), int64(m.HP)
	return nil
}

// uniqueName numbers repeated monsters: Goblin, Goblin 2, Goblin 3.
func uniqueName(ctx context.Context, q *sqlc.Queries, encounterID int64, name string) (string, error) {
	rows, err := q.ListCombatants(ctx, encounterID)
	if err != nil {
		return "", err
	}
	taken := func(n string) bool {
		return slices.ContainsFunc(rows, func(r sqlc.ListCombatantsRow) bool { return r.Name == n })
	}
	if !taken(name) {
		return name, nil
	}
	for i := 2; ; i++ {
		if n := fmt.Sprintf("%s %d", name, i); !taken(n) {
			return n, nil
		}
	}
}

// pcSheet is a PC row's live sheet; ok is false for other combatants.
func pcSheet(r sqlc.ListCombatantsRow) (dnd5e.Sheet, bool) {
	if !r.SheetJson.Valid {
		return dnd5e.Sheet{}, false
	}
	s, err := dnd5e.ParseSheet(r.SheetJson.String)
	return s, err == nil
}

// initiativeMod is the bonus a combatant adds to its initiative roll: a PC's
// from the sheet, a monster's (or a statted NPC's) from its statblock.
func initiativeMod(r sqlc.ListCombatantsRow) int {
	if s, ok := pcSheet(r); ok {
		return s.Initiative()
	}
	if r.MonsterData.Valid {
		if d, err := compendium.DecodeData(compendium.KindMonster, r.MonsterData.String); err == nil {
			return d.(*compendium.MonsterData).Initiative
		}
	}
	return 0
}

// outOfFight reports a combatant whose turn is skipped: a monster or NPC at
// 0 HP, or a PC who failed three death saves (a PC at 0 HP still rolls them).
func outOfFight(r sqlc.ListCombatantsRow) bool {
	if s, ok := pcSheet(r); ok {
		return s.DeathSaves.Failures >= 3
	}
	return !r.CharacterID.Valid && r.Hp <= 0 && r.MaxHp > 0
}

// stepTurn moves the turn from index i by dir (1 or -1), past combatants out
// of the fight, wrapping into the next or previous round. It returns the new
// index and round and the names it skipped.
func stepTurn(rows []sqlc.ListCombatantsRow, i int, round int64, dir int) (int, int64, []string, error) {
	var skipped []string
	for range rows {
		i += dir
		switch {
		case i >= len(rows):
			i, round = 0, round+1
		case i < 0:
			i, round = len(rows)-1, round-1
		}
		if round < 1 {
			return 0, 0, nil, userError{errors.New("this is the first turn")}
		}
		if !outOfFight(rows[i]) {
			return i, round, skipped, nil
		}
		skipped = append(skipped, rows[i].Name)
	}
	return 0, 0, nil, userError{errors.New("everyone is out of the fight")}
}

// skippedNotice tells the GM whose turns were skipped, and why.
func skippedNotice(names []string) string {
	if len(names) == 0 {
		return ""
	}
	return fmt.Sprintf("Skipped %d (out of the fight): %s", len(names), strings.Join(names, ", "))
}

// encounterAction runs start, next, prev, end and delete.
func (a *App) encounterAction(c *gin.Context) {
	enc, ok := a.encounterParam(c)
	if !ok {
		return
	}
	ctx := c.Request.Context()
	if c.Param("action") == "delete" {
		if err := a.Store.DeleteEncounter(ctx, enc.ID); err != nil {
			a.serverError(c, err)
			return
		}
		a.Hub.Publish(enc.CampaignID, hub.Topic{Kind: hub.TopicEncounter, ID: enc.ID})
		redirect(c, fmt.Sprintf("/campaigns/%d/combat", enc.CampaignID))
		return
	}
	var skipped []string
	err := a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		rows, err := q.ListCombatants(ctx, enc.ID)
		if err != nil {
			return err
		}
		next := sqlc.UpdateEncounterStateParams{ID: enc.ID, Status: enc.Status, Round: enc.Round, ActiveCombatantID: enc.ActiveCombatantID}
		switch action := c.Param("action"); action {
		case "start":
			if len(rows) == 0 {
				return userError{errors.New("add combatants first")}
			}
			next.Status, next.Round, next.ActiveCombatantID = statusActive, 1, nullID(rows[0].ID)
		case "next", "prev":
			if enc.Status != statusActive || len(rows) == 0 {
				return userError{errors.New("the encounter is not running")}
			}
			if slices.ContainsFunc(rows, func(r sqlc.ListCombatantsRow) bool { return !r.Initiative.Valid }) {
				return userError{errors.New("set everyone's initiative first")}
			}
			i := slices.IndexFunc(rows, func(r sqlc.ListCombatantsRow) bool { return r.ID == enc.ActiveCombatantID.Int64 })
			dir := 1
			if action == "prev" {
				dir = -1
			}
			if i, next.Round, skipped, err = stepTurn(rows, i, enc.Round, dir); err != nil {
				return err
			}
			next.ActiveCombatantID = nullID(rows[i].ID)
			// Timed effects end at the start of their creature's turn; going back
			// does not restore them.
			if action == "next" {
				if err := q.ExpireTurnConditions(ctx, sqlc.ExpireTurnConditionsParams{Round: sql.NullInt64{Int64: next.Round, Valid: true}, CombatantID: nullID(rows[i].ID)}); err != nil {
					return err
				}
			}
		case "end":
			next.Status, next.ActiveCombatantID = statusEnded, sql.NullInt64{}
		default:
			return userError{errors.New("unknown action")}
		}
		return q.UpdateEncounterState(ctx, next)
	})
	if a.commandFailed(c, err) {
		return
	}
	a.encounterNotice(c, enc, skippedNotice(skipped), a.pcTopics(ctx, enc.ID)...)
}

// pcTopics lists the characters in an encounter, whose sheets show
// conditions that may have expired.
func (a *App) pcTopics(ctx context.Context, encounterID int64) []hub.Topic {
	rows, err := a.Store.ListCombatants(ctx, encounterID)
	if err != nil {
		return nil
	}
	var out []hub.Topic
	for _, r := range rows {
		if r.CharacterID.Valid {
			out = append(out, hub.Topic{Kind: hub.TopicCharacter, ID: r.CharacterID.Int64})
		}
	}
	return out
}

// combatantAction runs initiative, hp, marker, visible, condition and delete.
func (a *App) combatantAction(c *gin.Context) {
	id, ok := a.idParam(c, "combatant")
	if !ok {
		return
	}
	ctx := c.Request.Context()
	cb, err := a.Store.GetCombatant(ctx, id)
	if !a.found(c, err) {
		return
	}
	enc, err := a.Store.GetEncounter(ctx, cb.EncounterID)
	if !a.found(c, err) {
		return
	}
	var extra []hub.Topic
	if cb.CharacterID.Valid {
		extra = append(extra, hub.Topic{Kind: hub.TopicCharacter, ID: cb.CharacterID.Int64})
	}

	switch c.Param("action") {
	case "initiative":
		init := sql.NullInt64{}
		if c.PostForm("roll") != "" {
			init, err = a.rollInitiative(ctx, cb)
		} else if n, ok := formInt(c, "initiative"); ok {
			init = sql.NullInt64{Int64: n, Valid: true}
		}
		if err == nil {
			err = a.Store.SetCombatantInitiative(ctx, sqlc.SetCombatantInitiativeParams{ID: id, Initiative: init})
		}
	case "hp":
		err = a.combatantHP(c, cb)
	case "marker":
		marker := strings.TrimSpace(c.PostForm("marker"))
		if utf8.RuneCountInString(marker) > views.MaxMarker {
			err = userError{fmt.Errorf("keep the marker to %d characters", views.MaxMarker)}
			break
		}
		err = a.Store.SetCombatantMarker(ctx, sqlc.SetCombatantMarkerParams{ID: id, Marker: marker})
	case "visible":
		err = a.Store.SetCombatantHPVisible(ctx, sqlc.SetCombatantHPVisibleParams{ID: id, HpVisible: !cb.HpVisible})
	case "condition":
		err = a.combatantCondition(c, cb, enc)
	case "delete":
		err = a.Store.DeleteCombatant(ctx, id)
	default:
		err = userError{errors.New("unknown action")}
	}
	if a.commandFailed(c, err) {
		return
	}
	a.encounterChanged(c, enc, extra...)
}

// rollInitiative rolls d20 plus the combatant's initiative bonus: the one
// roll the app makes, so the GM can set up a fight quickly.
func (a *App) rollInitiative(ctx context.Context, cb sqlc.Combatant) (sql.NullInt64, error) {
	rows, err := a.Store.ListCombatants(ctx, cb.EncounterID)
	if err != nil {
		return sql.NullInt64{}, err
	}
	i := slices.IndexFunc(rows, func(r sqlc.ListCombatantsRow) bool { return r.ID == cb.ID })
	if i < 0 {
		return sql.NullInt64{}, sql.ErrNoRows
	}
	return sql.NullInt64{Int64: int64(rand.IntN(20) + 1 + initiativeMod(rows[i])), Valid: true}, nil //nolint:gosec // a game die, not a secret
}

// combatantHP applies damage or healing, or sets HP, temporary HP and a
// monster's max HP. A PC's HP lives on the character, so the same rules as
// the sheet's hp command apply.
func (a *App) combatantHP(c *gin.Context, cb sqlc.Combatant) error {
	ctx := c.Request.Context()
	if c.PostForm("op") == "set" {
		return a.combatantSetHP(c, cb)
	}
	n, ok := formInt(c, "amount")
	if !ok || n < 0 || n > views.MaxHPInput {
		return userError{fmt.Errorf("enter an amount from 0 to %d", views.MaxHPInput)}
	}
	amount := dnd5e.Damage(int(n))
	if c.PostForm("op") == "heal" {
		amount = dnd5e.Heal(int(n))
	}
	if cb.CharacterID.Valid {
		return a.updateCharacter(ctx, cb.CharacterID.Int64, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
			dc.ApplyHP(amount)
			return nil
		})
	}
	m := dnd5e.Character{HP: int(cb.Hp), MaxHP: int(cb.MaxHp), TempHP: int(cb.TempHp)}
	m.ApplyHP(amount)
	return a.Store.SetCombatantHP(ctx, sqlc.SetCombatantHPParams{ID: cb.ID, Hp: int64(m.HP), MaxHp: cb.MaxHp, TempHp: int64(m.TempHP)})
}

// combatantSetHP sets the fields the HP dialog sends: only those the GM
// changed, so a form opened before someone else's change does not undo it.
func (a *App) combatantSetHP(c *gin.Context, cb sqlc.Combatant) error {
	var vals [3]*int
	for i, f := range []struct {
		name, label string
		lo          int
	}{{"hp", "HP", 0}, {"max_hp", "max HP", 1}, {"temp_hp", "temporary HP", 0}} {
		if strings.TrimSpace(c.PostForm(f.name)) == "" {
			continue
		}
		n, ok := formInt(c, f.name)
		if !ok || n < int64(f.lo) || n > views.MaxHPInput {
			return userError{fmt.Errorf("%s must be a whole number from %d to %d", f.label, f.lo, views.MaxHPInput)}
		}
		v := int(n)
		vals[i] = &v
	}
	hp, maxHP, temp := vals[0], vals[1], vals[2]
	if hp == nil && maxHP == nil && temp == nil {
		return userError{errors.New("change a value to set")}
	}
	set := func(dc *dnd5e.Character) {
		if maxHP != nil {
			if dc.MaxHP == 0 && hp == nil { // no stats yet: start at full
				dc.HP = *maxHP
			}
			dc.MaxHP = *maxHP
			dc.HP = min(dc.HP, dc.MaxHP)
		}
		if hp != nil {
			dc.ApplyHP(dnd5e.Set(*hp))
		}
		if temp != nil {
			dc.TempHP = *temp
		}
	}
	if cb.CharacterID.Valid {
		if maxHP != nil {
			return userError{errors.New("a character's max HP is set on their sheet")}
		}
		return a.updateCharacter(c.Request.Context(), cb.CharacterID.Int64, func(_ *sqlc.Queries, dc *dnd5e.Character) error {
			set(dc)
			return nil
		})
	}
	m := dnd5e.Character{HP: int(cb.Hp), MaxHP: int(cb.MaxHp), TempHP: int(cb.TempHp)}
	set(&m)
	return a.Store.SetCombatantHP(c.Request.Context(), sqlc.SetCombatantHPParams{ID: cb.ID, Hp: int64(m.HP), MaxHp: int64(m.MaxHP), TempHp: int64(m.TempHP)})
}

func (a *App) combatantCondition(c *gin.Context, cb sqlc.Combatant, enc sqlc.Encounter) error {
	slug := c.PostForm("slug")
	if !dnd5e.IsConditionSlug(slug) {
		return userError{fmt.Errorf("unknown condition %q", slug)}
	}
	note := strings.TrimSpace(c.PostForm("note"))
	if slug == dnd5e.TrackCustom && note == "" {
		return userError{errors.New("describe the custom effect")}
	}
	var expires sql.NullInt64
	if rounds, ok := formInt(c, "rounds"); ok && rounds > 0 {
		expires = sql.NullInt64{Int64: max(enc.Round, 1) + rounds - 1, Valid: true}
	}
	owner := condOwner{combatantID: cb.ID}
	if cb.CharacterID.Valid {
		owner = condOwner{characterID: cb.CharacterID.Int64}
	}
	op := dnd5e.CondOp{Op: dnd5e.CondAdd, Slug: slug, Note: note}
	if slug == "exhaustion" {
		op = dnd5e.CondOp{Op: dnd5e.CondExhaustion, Slug: slug, Delta: 1}
	}
	ctx := c.Request.Context()
	return a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		return applyCondOps(ctx, q, owner, []dnd5e.CondOp{op}, expires)
	})
}
