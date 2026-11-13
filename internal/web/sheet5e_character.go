package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

func (a *App) sheetView(ctx context.Context, v perm.Viewer, row sqlc.Character) (views.SheetView, error) {
	sheet, err := dnd5e.ParseSheet(row.SheetJson)
	if err != nil {
		return views.SheetView{}, err
	}
	conds, err := a.Store.ListConditionsByCharacter(ctx, nullID(row.ID))
	if err != nil {
		return views.SheetView{}, err
	}
	rows, err := a.Store.ListCharacterImages(ctx, row.ID)
	if err != nil {
		return views.SheetView{}, err
	}
	var looks []perm.Look
	for _, l := range rows {
		looks = append(looks, toLook(l.ID, l.Name, l.Url, l.AttachmentID, l.IsRevealed, l.IsCurrent))
	}
	spells, err := a.sheetSpells(ctx, sheet)
	if err != nil {
		return views.SheetView{}, err
	}
	levelUps, err := a.Store.CountLevelUps(ctx, row.ID)
	if err != nil {
		return views.SheetView{}, err
	}
	masteries, err := masteryLimit(ctx, a.Store.Queries, sheet, "")
	if err != nil {
		return views.SheetView{}, err
	}
	canSheet := perm.CanCharacter(v, row.ID, perm.ScopeSheet)
	canSources := perm.CanCharacter(v, row.ID, perm.ScopeSources)
	var issues []views.SourceIssue
	if canSources {
		if issues, err = a.sourceIssues(ctx, sheet, row.CampaignID); err != nil {
			return views.SheetView{}, err
		}
	}
	return views.SheetView{
		ID: row.ID, CampaignID: row.CampaignID, Conditions: conds, Looks: perm.VisibleLooks(v, row.ID, looks), SpellData: spells,
		C:             dnd5e.Character{Name: row.Name, HP: int(row.Hp), MaxHP: int(row.MaxHp), TempHP: int(row.TempHp), Sheet: sheet},
		CanHP:         perm.CanCharacter(v, row.ID, perm.ScopeHP),
		CanConditions: perm.CanCharacter(v, row.ID, perm.ScopeConditions),
		CanSheet:      canSheet,
		CanRevert:     canSheet && levelUps > 0,
		CanSources:    canSources,
		SourceIssues:  issues,
		MasteryLimit:  masteries,
	}, nil
}

// sheetSpells are the compendium details of the sheet's spells, by slug:
// the priority copies in one query, and a pinned copy per spell pinned to
// another source while it exists.
func (a *App) sheetSpells(ctx context.Context, sheet dnd5e.Sheet) (map[string]*compendium.SpellData, error) {
	if len(sheet.Spells) == 0 {
		return nil, nil
	}
	slugs := make([]string, len(sheet.Spells))
	for i, sp := range sheet.Spells {
		slugs[i] = sp.Slug
	}
	rows, err := a.Store.ListCompendiumBySlugs(ctx, sqlc.ListCompendiumBySlugsParams{Kind: enum.KindSpell, Slugs: slugs})
	if err != nil {
		return nil, err
	}
	best := map[string]sqlc.ListCompendiumBySlugsRow{}
	for _, r := range rows {
		best[r.Slug] = r
	}
	out := map[string]*compendium.SpellData{}
	for _, sp := range sheet.Spells {
		if _, seen := out[sp.Slug]; seen {
			continue
		}
		r, ok := best[sp.Slug]
		data := r.DataJson
		if sp.Source != "" && sp.Source != r.SourceKey {
			e, err := a.Store.GetCompendiumEntryBySlugSource(ctx, sqlc.GetCompendiumEntryBySlugSourceParams{Kind: enum.KindSpell, Slug: sp.Slug, SourceKey: sp.Source})
			switch {
			case err == nil:
				data, ok = e.DataJson, true
			case !errors.Is(err, sql.ErrNoRows):
				return nil, err
			}
		}
		if !ok {
			continue
		}
		d, err := compendium.Decode[compendium.SpellData](enum.KindSpell, data)
		if err != nil {
			return nil, err
		}
		out[sp.Slug] = d
	}
	return out, nil
}

// characterAccess loads a character the requester may view.
func (a *App) characterAccess(c *gin.Context) (who, sqlc.Character, bool) {
	id, ok := a.idParam(c, "character")
	if !ok {
		return who{}, sqlc.Character{}, false
	}
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return w, sqlc.Character{}, false
	}
	if !perm.CanCharacter(w.viewer, id, perm.ScopeView) {
		a.fail(c, http.StatusForbidden, "That is not your character.")
		return w, sqlc.Character{}, false
	}
	ch, err := a.Store.GetCharacter(c.Request.Context(), id)
	return w, ch, a.found(c, err)
}

func (a *App) characterPage(c *gin.Context) error {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return nil
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), ch.CampaignID)
	if err != nil {
		return err
	}
	l := a.layout(w, ch.Name, views.NavCharacter, &camp, true)
	s, err := a.sheetView(c.Request.Context(), w.viewer, ch)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.CharacterPage(l, s))
	return nil
}

// commandText builds the command from the sheet's forms: a typed command,
// or the coin and inventory shortcuts.
func commandText(c *gin.Context) string {
	if _, ok := c.GetPostForm("was_gp"); ok {
		return coinCommand(c)
	}
	if name := strings.TrimSpace(c.PostForm("item_name")); name != "" {
		qty := strings.TrimSpace(c.PostForm("item_qty"))
		if qty == "" {
			qty = "1"
		}
		return fmt.Sprintf("item add %s %s", strconv.Quote(name), qty)
	}
	return c.PostForm("cmd")
}

// coinCommand turns the coin form into one command that changes each coin
// by the difference between the amount shown and the amount typed.
func coinCommand(c *gin.Context) string {
	cmd := "coin"
	for _, k := range []string{"pp", "gp", "ep", "sp", "cp"} {
		typed := strings.TrimSpace(c.PostForm("coin_" + k))
		n, err := strconv.Atoi(typed)
		was, werr := strconv.Atoi(c.PostForm("was_" + k))
		if err != nil || werr != nil {
			return "coin " + k + " =" + typed // the parser explains what is wrong
		}
		if n != was {
			cmd += fmt.Sprintf(" %s %+d", k, n-was)
		}
	}
	if cmd == "coin" {
		return "coin gp +0" // nothing changed: succeed without a change
	}
	return cmd
}

// hpForm is the HP card's quick form: an amount and what to do with it.
type hpForm struct {
	Op     string `form:"hp_op"     validate:"oneof=damage heal temp"`
	Amount int    `form:"hp_amount" validate:"min=1"`
	// Type is the damage type, " " (the select's Other) or empty for untyped.
	Type string `form:"hp_type"`
}

// hpCommand turns the HP quick form into damage (of a type), healing or
// temp HP.
func hpCommand(c *gin.Context) (dnd5e.Command, error) {
	var f hpForm
	if err := c.ShouldBind(&f); err != nil {
		return dnd5e.Command{}, errors.New("enter how many hit points, 1 or more")
	}
	line := map[string]string{"damage": "hp -", "heal": "hp +", "temp": "temp "}[f.Op] + strconv.Itoa(f.Amount)
	if kind := strings.TrimSpace(f.Type); f.Op == "damage" && kind != "" {
		return dnd5e.Parse(line, kind) // Parse checks the type
	}
	return dnd5e.Parse(line)
}

// builderArgs are the command builder's fields, in the order the commands
// take them (see views.CommandBuilder). from and fromnote name the thing an
// edit changes.
var builderArgs = []string{argFrom, argFromNote, argSubject, argAbility, argDice, argDamageType, argProf, argN, argText}

// Builder form fields passed to the command as arguments, in order.
const (
	argFrom       = "from"
	argFromNote   = "fromnote"
	argSubject    = "subject"
	argAbility    = "ability"
	argDice       = "dice"
	argDamageType = "dtype"
	argProf       = "prof"
	argN          = "n"
	argText       = "text"
)

// parseCommand reads a command from the builder (a prefix plus raw
// arguments, so names and free text need no quoting) or the other forms.
func parseCommand(c *gin.Context) (dnd5e.Command, error) {
	if _, ok := c.GetPostForm("hp_op"); ok {
		return hpCommand(c)
	}
	prefix, ok := c.GetPostForm("prefix")
	if !ok {
		return dnd5e.Parse(commandText(c))
	}
	switch prefix {
	case "feature add":
		return featureCommand(c)
	case "weapon add":
		return weaponCommand(c)
	}
	var raw []string
	for _, f := range builderArgs {
		v, ok := c.GetPostForm(f)
		// Free text stays as typed; fromnote must match a stored note exactly.
		if f != argText && f != argFromNote {
			v = strings.TrimSpace(v)
		}
		if f == argN && v != "" {
			v = c.PostForm("sign") + v
		}
		// An empty text still counts: it clears a backstory.
		if ok && (v != "" || f == argText) {
			raw = append(raw, v)
		}
	}
	return dnd5e.Parse(prefix, raw...)
}

// commandBuilder renders one form of the command builder dialog.
func (a *App) commandBuilder(c *gin.Context) error {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return nil
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player changes their character.")
		return nil
	}
	s, err := a.sheetView(c.Request.Context(), w.viewer, ch)
	if err != nil {
		return err
	}
	x, err := a.builderExtra(c, ch.CampaignID)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.CommandBuilder(s, c.Query("verb"), c.Query("op"), c.Query(argSubject), x))
	return nil
}

// characterCommand applies one command in a transaction and pushes the
// result to everyone who can see the character.
func (a *App) characterCommand(c *gin.Context) error {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return nil
	}
	cmd, err := parseCommand(c)
	if err != nil {
		return userError{err}
	}
	if !perm.CanCharacter(w.viewer, ch.ID, cmd.Scope) {
		a.fail(c, http.StatusForbidden, "The GM may change HP and conditions only.")
		return nil
	}
	if err := a.runCommand(c.Request.Context(), ch, cmd); err != nil {
		return err
	}
	a.characterChanged(c, ch)
	return nil
}

// runCommand looks up what cmd needs from the compendium and applies it to
// the character in one transaction. Errors the user caused are userErrors.
func (a *App) runCommand(ctx context.Context, ch sqlc.Character, cmd dnd5e.Command) error {
	return a.updateCharacter(ctx, ch.ID, func(q *sqlc.Queries, dc *dnd5e.Character) error {
		// Built from the sheet this transaction writes, not ch's earlier copy.
		env, err := a.commandEnv(ctx, q, dc.Sheet, ch.CampaignID, cmd)
		if err != nil {
			return err
		}
		conds, err := q.ListConditionsByCharacter(ctx, sql.NullInt64{Int64: ch.ID, Valid: true})
		if err != nil {
			return err
		}
		env.Petrified = petrified(conds)
		fx, err := cmd.Apply(dc, env)
		var ue userError
		if errors.As(err, &ue) {
			return err // keeps a fieldError's field
		}
		if err != nil {
			return userError{err}
		}
		return applyCondOps(ctx, q, condOwner{characterID: ch.ID}, fx.Conditions, sql.NullInt64{})
	})
}

// commandEnv looks up what cmd needs from the compendium for sheet.
func (a *App) commandEnv(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet, campaignID int64, cmd dnd5e.Command) (dnd5e.Env, error) {
	var env dnd5e.Env
	var err error
	if cmd.SourceRef.Kind != "" {
		if env, _, err = a.sourceEnv(ctx, q, sheet, cmd.SourceRef, campaignID); err != nil {
			return env, err
		}
	}
	if cmd.SubclassChange.Class != "" {
		if env.SourceOld, env.SourceNew, err = a.subclassChangeEnv(ctx, q, sheet, cmd.SubclassChange, campaignID); err != nil {
			return env, err
		}
	}
	if cmd.SpellSlug != "" {
		if env.Spell, err = a.compendiumSpell(ctx, cmd.SpellSlug, campaignID); err != nil {
			return env, userError{err}
		}
	}
	if cmd.PrepareCaps {
		if env.PreparedCaps, err = preparedCaps(ctx, q, sheet); err != nil {
			return env, err
		}
	}
	if cmd.MasteryWeapon != "" {
		if env.MasteryWeapon, env.MasteryLimit, err = a.masteryEnv(ctx, q, sheet, campaignID, cmd.MasteryWeapon); err != nil {
			return env, err
		}
	}
	// Armor the character carries is worn without a compendium lookup.
	if cmd.ArmorSlug != "" && sheet.ArmorItem(cmd.ArmorSlug) < 0 {
		if env.Armor, env.IsShield, err = a.compendiumArmor(ctx, cmd.ArmorSlug, campaignID); err != nil {
			return env, userError{err}
		}
	}
	return env, nil
}

// preparedCaps are the prepared-spell totals of the sheet's classes at
// their levels, by class slug; a class the compendium lacks has no cap.
func preparedCaps(ctx context.Context, q *sqlc.Queries, sheet dnd5e.Sheet) (map[string]int, error) {
	caps := map[string]int{}
	for _, cl := range sheet.Classes {
		info, _, err := classEntry(ctx, q, cl.Slug, cl.Source)
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		if err != nil {
			return nil, err
		}
		if _, n, prepared := dnd5e.SpellTotals(info, cl.Level); prepared && n >= 0 {
			caps[cl.Slug] = n
		}
	}
	return caps, nil
}

// userError marks an error caused by the request, shown to the user as is.
type userError struct{ error }

// fieldError is a userError about one form field, shown next to it.
type fieldError struct {
	userError
	field string
}

func (e fieldError) Unwrap() error { return e.userError }

// pickError is a choice group's fieldError, named by its group's title when
// the page does not draw the group and shows it at the top.
type pickError struct {
	fieldError
	title string
}

func (e pickError) Unwrap() error { return e.fieldError }

// formErrors splits err (one error or several joined) into the field errors
// the page draws, the first per field, and the rest as top lines.
func formErrors(err error, draws func(string) bool) (string, views.FieldErrors) {
	var top []string
	fields := views.FieldErrors{}
	var walk func(error)
	walk = func(err error) {
		var joined interface{ Unwrap() []error }
		var pe pickError
		var fe fieldError
		var ue userError
		switch {
		case errors.As(err, &joined):
			for _, x := range joined.Unwrap() {
				walk(x)
			}
		case errors.As(err, &pe) && !draws(pe.field):
			top = append(top, pe.title+": "+pe.Error())
		case errors.As(err, &fe):
			if !draws(fe.field) {
				top = append(top, fe.Error()) // a crafted post's field is not on the page
			} else if fields[fe.field] == "" {
				fields[fe.field] = fe.Error()
			}
		case errors.As(err, &ue):
			top = append(top, ue.Error())
		}
	}
	walk(err)
	return strings.Join(top, "\n"), fields
}

// updateCharacter reads, changes and writes a character in one transaction.
func (a *App) updateCharacter(ctx context.Context, id int64, fn func(q *sqlc.Queries, dc *dnd5e.Character) error) error {
	return a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		row, err := q.GetCharacter(ctx, id)
		if err != nil {
			return err
		}
		sheet, err := dnd5e.ParseSheet(row.SheetJson)
		if err != nil {
			return err
		}
		dc := dnd5e.Character{Name: row.Name, HP: int(row.Hp), MaxHP: int(row.MaxHp), TempHP: int(row.TempHp), Sheet: sheet}
		if err := fn(q, &dc); err != nil {
			return err
		}
		raw, err := dc.Sheet.JSON()
		if err != nil {
			return err
		}
		return q.UpdateCharacter(ctx, sqlc.UpdateCharacterParams{
			ID: id, Name: dc.Name, Hp: int64(dc.HP), MaxHp: int64(dc.MaxHP), TempHp: int64(dc.TempHP), SheetJson: raw,
		})
	})
}

// characterChanged pushes a character and the fights it is in.
func (a *App) characterChanged(c *gin.Context, ch sqlc.Character) {
	topics, err := a.characterTopics(c.Request.Context(), ch.ID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.changed(c, ch.CampaignID, topics...)
}

// characterTopics are what a change to a character touches: the character
// and the fights it is in.
func (a *App) characterTopics(ctx context.Context, id int64) ([]hub.Topic, error) {
	topics := []hub.Topic{{Kind: hub.TopicCharacter, ID: id}}
	encounters, err := a.Store.ListEncounterIDsByCharacter(ctx, nullID(id))
	for _, e := range encounters {
		topics = append(topics, hub.Topic{Kind: hub.TopicEncounter, ID: e})
	}
	return topics, err
}

// compendiumArmor finds armor (or a shield) in the compendium by slug.
func (a *App) compendiumArmor(ctx context.Context, slug string, campaignID int64) (dnd5e.WornArmor, bool, error) {
	row, err := a.entryBySlug(ctx, enum.KindItem, slug, campaignID)
	if errors.Is(err, sql.ErrNoRows) {
		return dnd5e.WornArmor{}, false, fmt.Errorf("no item %q in the compendium", slug)
	}
	if err != nil {
		return dnd5e.WornArmor{}, false, err
	}
	d, err := compendium.Decode[compendium.ItemData](enum.KindItem, row.DataJson)
	if err != nil {
		return dnd5e.WornArmor{}, false, err
	}
	ar := d.Armor
	switch {
	case ar == nil:
		return dnd5e.WornArmor{}, false, fmt.Errorf("%s is not armor", row.Name)
	case ar.Type == compendium.ArmorShield:
		return dnd5e.WornArmor{}, true, nil
	}
	return dnd5e.WornArmor{Slug: row.Slug, Name: row.Name, Source: row.SourceKey, Type: ar.Type, Base: ar.Base, AddDex: ar.AddDex, DexCap: ar.DexCap, StrReq: ar.StrReq, Stealth: ar.Stealth}, false, nil
}

func (a *App) compendiumSpell(ctx context.Context, slug string, campaignID int64) (dnd5e.Spell, error) {
	row, err := a.entryBySlug(ctx, enum.KindSpell, slug, campaignID)
	if errors.Is(err, sql.ErrNoRows) {
		return dnd5e.Spell{}, userError{fmt.Errorf("no spell %q in the compendium", slug)}
	}
	if err != nil {
		return dnd5e.Spell{}, err
	}
	return spellFromRow(row)
}

// spellFromRow is a compendium spell as the sheet stores it.
func spellFromRow(row sqlc.CompendiumEntry) (dnd5e.Spell, error) {
	d, err := compendium.Decode[compendium.SpellData](enum.KindSpell, row.DataJson)
	if err != nil {
		return dnd5e.Spell{}, err
	}
	return dnd5e.Spell{Slug: row.Slug, Name: row.Name, Source: row.SourceKey, Level: d.Level}, nil
}

// condOwner is the character or combatant a condition belongs to.
type condOwner struct {
	characterID, combatantID int64
}

func (o condOwner) find(ctx context.Context, q *sqlc.Queries, slug string) (sqlc.Condition, error) {
	if o.characterID != 0 {
		return q.FindCharacterCondition(ctx, sqlc.FindCharacterConditionParams{CharacterID: nullID(o.characterID), Slug: slug})
	}
	return q.FindCombatantCondition(ctx, sqlc.FindCombatantConditionParams{CombatantID: nullID(o.combatantID), Slug: slug})
}

func (o condOwner) add(ctx context.Context, q *sqlc.Queries, slug, note string, expires sql.NullInt64) error {
	_, err := q.CreateCondition(ctx, sqlc.CreateConditionParams{
		CharacterID: nullID(o.characterID), CombatantID: nullID(o.combatantID), Slug: slug, Note: note, ExpiresRound: expires,
	})
	return err
}

// applyCondOps applies condition changes. Each 2024 condition and
// concentration is held at most once; custom effects may stack.
func applyCondOps(ctx context.Context, q *sqlc.Queries, o condOwner, ops []dnd5e.CondOp, expires sql.NullInt64) error {
	for _, op := range ops {
		existing, err := o.find(ctx, q, op.Slug)
		has := err == nil
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		remove := func() error {
			if has {
				return q.DeleteCondition(ctx, existing.ID)
			}
			return nil
		}
		switch op.Op {
		case dnd5e.CondToggle:
			if has {
				err = remove()
			} else {
				err = o.add(ctx, q, op.Slug, op.Note, expires)
			}
		case dnd5e.CondAdd:
			if op.Slug != dnd5e.TrackCustom {
				if err = remove(); err != nil {
					return err
				}
			}
			err = o.add(ctx, q, op.Slug, op.Note, expires)
		case dnd5e.CondRemove:
			err = remove()
		case dnd5e.CondExhaustion:
			level := 0
			if has {
				level, _ = strconv.Atoi(existing.Note)
			}
			level = min(max(level+op.Delta, 0), dnd5e.MaxExhaustion)
			if err = remove(); err == nil && level > 0 {
				err = o.add(ctx, q, op.Slug, strconv.Itoa(level), sql.NullInt64{})
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// deleteCondition removes one condition: the GM anywhere, a player on
// their own character.
func (a *App) deleteCondition(c *gin.Context) error {
	id, ok := a.idParam(c, "condition")
	if !ok {
		return nil
	}
	w, err := a.whoami(c)
	if err != nil {
		return err
	}
	ctx := c.Request.Context()
	cond, err := a.Store.GetCondition(ctx, id)
	if err != nil {
		return err
	}
	if cond.CharacterID.Valid {
		if !perm.CanCharacter(w.viewer, cond.CharacterID.Int64, perm.ScopeConditions) {
			a.fail(c, http.StatusForbidden, "Not your character.")
			return nil
		}
		ch, err := a.Store.GetCharacter(ctx, cond.CharacterID.Int64)
		if err != nil {
			return err
		}
		if err := a.Store.DeleteCondition(ctx, id); err != nil {
			return err
		}
		a.characterChanged(c, ch)
		return nil
	}
	if !w.isGM() {
		a.fail(c, http.StatusForbidden, "Only the GM changes monsters.")
		return nil
	}
	cb, err := a.Store.GetCombatant(ctx, cond.CombatantID.Int64)
	if err != nil {
		return err
	}
	enc, err := a.Store.GetEncounter(ctx, cb.EncounterID)
	if err != nil {
		return err
	}
	if err := a.Store.DeleteCondition(ctx, id); err != nil {
		return err
	}
	a.changed(c, enc.CampaignID, hub.Topic{Kind: hub.TopicEncounter, ID: enc.ID})
	return nil
}
