package web

// Effects: condition rows that also end on a rest or with a caster's
// concentration and change the numbers on a sheet while they last. Their
// modifiers are shown with the sheet (effectSheet) and never saved into it.

import (
	"cmp"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
)

// effectSep joins an effect's name and note.
const effectSep = " — "

// effectNoteMax is the note's length cap (the validate tag on effectForm.Note).
const effectNoteMax = 90

// effectForm is the posted form that adds a condition or effect to a
// character or combatant. The modifier rows are parallel lists. Name and
// note are capped so that name + effectSep + note fits what a modifier's
// note and source take (dnd5e.CustomMod, 200).
type effectForm struct {
	Slug      string   `form:"slug"       validate:"required"`
	Name      string   `form:"name"       validate:"max=100"`
	Note      string   `form:"note"       validate:"max=90"`
	Spell     string   `form:"spell"      validate:"max=200"`
	Rounds    int      `form:"rounds"     validate:"omitempty,min=1,max=1000"`
	EndsOn    string   `form:"ends_on"    validate:"omitempty,oneof=short_rest long_rest"`
	HeldBy    int64    `form:"held_by"    validate:"omitempty,min=1"`
	Target    int64    `form:"target"     validate:"omitempty,min=1"`
	ModTarget []string `form:"mod_target" validate:"max=10"`
	ModValue  []string `form:"mod_value"  validate:"max=10"`
	// SavePreset also saves the form as the poster's preset (same name overwrites).
	// A checkbox posts "on", which gin cannot read as a bool.
	SavePreset string `form:"save_preset" validate:"omitempty,oneof=on"`
}

// bindEffect reads and validates the posted effect form.
func bindEffect(c *gin.Context) (effectForm, error) {
	var f effectForm
	if err := c.ShouldBind(&f); err != nil {
		return f, userError{fmt.Errorf("check the form: %w", err)}
	}
	return f, nil
}

// effectName is what an effect is called: its note, else its condition.
func effectName(c sqlc.Condition) string { return cmp.Or(c.Note, c.Slug) }

// effectNote joins an effect's name and note into the stored note.
func effectNote(name, note string) string {
	switch {
	case name == "":
		return note
	case note == "":
		return name
	}
	return name + effectSep + note
}

// modRows reads parallel target, value and when lists, as the builder and
// the effect form post them; a row without a target or value is skipped.
func modRows(targets, values, whens []string) ([]dnd5e.CustomMod, error) {
	if len(targets) > 0 && len(values) == 0 {
		return nil, errors.New("give each change a value")
	}
	var mods []dnd5e.CustomMod
	for i, t := range targets {
		t = strings.TrimSpace(t)
		v, _ := strconv.Atoi(strings.TrimSpace(values[min(i, len(values)-1)]))
		if t == "" || v == 0 {
			continue
		}
		if !dnd5e.IsModTarget(t) {
			return nil, fmt.Errorf("cannot change %q", t)
		}
		var when string
		if len(whens) > 0 {
			when = strings.TrimSpace(whens[min(i, len(whens)-1)])
		}
		mods = append(mods, dnd5e.CustomMod{Target: t, Value: v, When: when})
	}
	return mods, nil
}

// effectSheet is sheet with the modifiers of the active effects in conds
// added, for showing numbers. A malformed row is skipped, not fatal.
func effectSheet(sheet dnd5e.Sheet, conds []sqlc.Condition) dnd5e.Sheet {
	var mods []dnd5e.CustomMod
	for _, c := range conds {
		if c.ModsJson == "" || c.ModsJson == noMods {
			continue
		}
		var rows []dnd5e.CustomMod
		if err := json.Unmarshal([]byte(c.ModsJson), &rows); err != nil {
			slog.Warn("skipping malformed effect modifiers", "condition", c.ID, "err", err)
			continue
		}
		mods = append(mods, dnd5e.EffectMods(effectName(c), rows)...)
	}
	return sheet.WithEffects(mods)
}

// addEffect puts the posted condition or effect on o, in a campaign. v is
// who posts it: a held_by row must be one v may use. expires is the round it
// ends, from the form's rounds.
func (a *App) addEffect(ctx context.Context, v perm.Viewer, o condOwner, campaignID int64, f effectForm, expires sql.NullInt64) error {
	if !dnd5e.IsConditionSlug(f.Slug) {
		return userError{fmt.Errorf("unknown condition %q", f.Slug)}
	}
	name := strings.TrimSpace(f.Name)
	if name == "" && f.Spell != "" {
		var err error
		if name, _, _, _, err = a.spellPrefill(ctx, f.Spell, campaignID); err != nil {
			return err
		}
	}
	note := effectNote(name, strings.TrimSpace(f.Note))
	if f.Slug == dnd5e.TrackCustom && note == "" {
		return userError{errors.New("describe the custom effect")}
	}
	mods, err := modRows(f.ModTarget, f.ModValue, nil) // an effect's change always applies
	if err != nil {
		return userError{err}
	}
	if err := dnd5e.ValidateEffectMods(cmp.Or(note, f.Slug), mods); err != nil {
		return userError{err}
	}
	x := condExtra{Expires: expires, EndsOn: enum.EffectEnd(f.EndsOn), HeldBy: nullID(f.HeldBy), ModsJSON: noMods}
	if len(mods) > 0 {
		raw, err := json.Marshal(dnd5e.EffectMods(cmp.Or(note, f.Slug), mods))
		if err != nil {
			return err
		}
		x.ModsJSON = string(raw)
	}
	if v.IsPlayer() && o.characterID != v.CharacterID {
		x.AppliedBy = nullID(v.CharacterID)
	}
	op := dnd5e.CondOp{Op: dnd5e.CondAdd, Slug: f.Slug, Note: note}
	switch {
	case f.Slug == dnd5e.Exhaustion && (len(mods) > 0 || f.HeldBy != 0 || f.EndsOn != ""):
		return userError{errors.New("exhaustion has levels, not effects")}
	case f.Slug == dnd5e.Exhaustion:
		op = dnd5e.CondOp{Op: dnd5e.CondExhaustion, Slug: f.Slug, Delta: 1}
	case f.Slug == slugConcentration && f.HeldBy != 0:
		return userError{errors.New("concentration cannot be held by another")}
	}
	var preset *sqlc.UpsertEffectPresetParams
	if f.SavePreset != "" {
		if name == "" {
			return userError{errors.New("name the preset")}
		}
		raw, err := json.Marshal(mods) // plain target/value rows, stamped again when applied
		if err != nil {
			return err
		}
		preset = &sqlc.UpsertEffectPresetParams{CampaignID: campaignID, Owner: nullID(v.CharacterID), Name: name, Slug: f.Slug, Note: strings.TrimSpace(f.Note), EndsOn: enum.EffectEnd(f.EndsOn), Rounds: int64(f.Rounds), ModsJson: string(raw)}
	}
	return a.Store.InTx(ctx, func(q *sqlc.Queries) error {
		if f.HeldBy != 0 {
			if err := checkHolder(ctx, q, v, campaignID, f.HeldBy); err != nil {
				return err
			}
		}
		if err := applyCondOps(ctx, q, o, []dnd5e.CondOp{op}, x); err != nil {
			return err
		}
		if preset == nil {
			return nil
		}
		_, err := q.UpsertEffectPreset(ctx, *preset)
		return err
	})
}

// presetNames are the presets of owner in a campaign, for the picker.
func (a *App) presetNames(ctx context.Context, campaignID, owner int64) ([]views.Named, error) {
	rows, err := a.Store.ListEffectPresets(ctx, sqlc.ListEffectPresetsParams{CampaignID: campaignID, Owner: nullID(owner)})
	if err != nil {
		return nil, err
	}
	out := make([]views.Named, 0, len(rows))
	for _, r := range rows {
		out = append(out, views.Named{ID: r.ID, Name: r.Name})
	}
	return out, nil
}

// deleteEffectPreset removes the poster's own preset (the posted preset id)
// and answers the refreshed picker.
func (a *App) deleteEffectPreset(c *gin.Context) error {
	campaignID, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, _, ok := a.campaignAccess(c, campaignID)
	if !ok {
		return nil
	}
	id, err := strconv.ParseInt(c.PostForm("preset"), 10, 64)
	if err != nil {
		return userError{errors.New("pick a preset")}
	}
	ctx := c.Request.Context()
	p, err := a.Store.GetEffectPreset(ctx, id)
	if err != nil {
		return err
	}
	if p.CampaignID != campaignID {
		return sql.ErrNoRows
	}
	if !perm.CanUsePreset(w.viewer, p.OwnerCharacterID.Int64) {
		a.fail(c, http.StatusForbidden, "That is not your preset.")
		return nil
	}
	n, err := a.Store.DeleteEffectPreset(ctx, sqlc.DeleteEffectPresetParams{ID: id, CampaignID: p.CampaignID, Owner: p.OwnerCharacterID})
	if err != nil {
		return err
	}
	if n == 0 {
		return sql.ErrNoRows
	}
	names, err := a.presetNames(ctx, p.CampaignID, p.OwnerCharacterID.Int64)
	if err != nil {
		return err
	}
	a.render(c, http.StatusOK, views.EffectPresetPicker(campaignID, names))
	return nil
}

// checkHolder refuses a held_by that is not a concentration row of this
// campaign that v may use: the GM any, a player their own character's.
func checkHolder(ctx context.Context, q *sqlc.Queries, v perm.Viewer, campaignID, holderID int64) error {
	bad := userError{errors.New("choose a concentration in this campaign")}
	h, err := q.GetCondition(ctx, holderID)
	if errors.Is(err, sql.ErrNoRows) || err == nil && h.Slug != slugConcentration {
		return bad
	}
	if err != nil {
		return err
	}
	var holderCampaign int64
	if h.CharacterID.Valid {
		ch, err := q.GetCharacter(ctx, h.CharacterID.Int64)
		if errors.Is(err, sql.ErrNoRows) {
			return bad
		}
		if err != nil {
			return err
		}
		if !perm.CanCharacter(v, ch.ID, perm.ScopeConditions) {
			return userError{errors.New("a player may only use their own concentration")}
		}
		holderCampaign = ch.CampaignID
	} else {
		cb, err := q.GetCombatant(ctx, h.CombatantID.Int64)
		if err != nil {
			return err
		}
		enc, err := q.GetEncounter(ctx, cb.EncounterID)
		if err != nil {
			return err
		}
		if !v.IsGM {
			return userError{errors.New("a player may only use their own concentration")}
		}
		holderCampaign = enc.CampaignID
	}
	if holderCampaign != campaignID {
		return bad
	}
	return nil
}

// canAddEffect is perm.CanAddEffect for a target character: the campaign rule
// compares it with the player's own character.
func (a *App) canAddEffect(ctx context.Context, v perm.Viewer, target sqlc.Character) (bool, error) {
	same := false
	if v.IsPlayer() {
		own, err := a.Store.GetCharacter(ctx, v.CharacterID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		same = err == nil && own.CampaignID == target.CampaignID
	}
	return perm.CanAddEffect(v, target.ID, same), nil
}

// characterEffect adds a condition or effect to a character's sheet: the GM
// on anyone, a player on themselves or, held by their own concentration, on
// an ally in their campaign (the posted target, else the path's character).
func (a *App) characterEffect(c *gin.Context) error {
	id, ok := a.idParam(c, "character")
	if !ok {
		return nil
	}
	w, err := a.whoami(c)
	if err != nil {
		return err
	}
	f, err := bindEffect(c)
	if err != nil {
		return err
	}
	ctx := c.Request.Context()
	ch, err := a.Store.GetCharacter(ctx, cmp.Or(f.Target, id))
	if err != nil {
		return err
	}
	allowed, err := a.canAddEffect(ctx, w.viewer, ch)
	if err != nil {
		return err
	}
	if !allowed {
		a.fail(c, http.StatusForbidden, "That is not your character.")
		return nil
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeConditions) && f.HeldBy == 0 {
		return userError{errors.New("on an ally, hold the effect with your concentration")}
	}
	expires, err := a.characterRounds(ctx, ch.ID, f.Rounds)
	if err != nil {
		return err
	}
	if err := a.addEffect(ctx, w.viewer, condOwner{characterID: ch.ID}, ch.CampaignID, f, expires); err != nil {
		return err
	}
	a.characterChanged(c, ch)
	return nil
}

// characterRounds is the round an effect of rounds rounds ends in the
// running fight the character is in. Rounds only count in a fight.
func (a *App) characterRounds(ctx context.Context, characterID int64, rounds int) (sql.NullInt64, error) {
	if rounds == 0 {
		return sql.NullInt64{}, nil
	}
	enc, ok, err := a.characterFight(ctx, characterID)
	if err != nil {
		return sql.NullInt64{}, err
	}
	if ok {
		return roundsEnd(enc.Round, rounds), nil
	}
	return sql.NullInt64{}, userError{errors.New("rounds count only in a fight that is running; use a rest instead")}
}

// roundsEnd is the last round of an effect of rounds rounds started in round.
func roundsEnd(round int64, rounds int) sql.NullInt64 {
	return sql.NullInt64{Int64: max(round, 1) + int64(rounds) - 1, Valid: true}
}

// effectSpell answers the effect form's fields for a character (?character=)
// or combatant (?combatant=), filled in from the picked spell (?spell=, a
// name or slug; none leaves them blank).
func (a *App) effectSpell(c *gin.Context) error {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return nil
	}
	w, _, ok := a.campaignAccess(c, id)
	if !ok {
		return nil
	}
	ctx := c.Request.Context()
	characterID, _ := strconv.ParseInt(c.Query("character"), 10, 64)
	combatantID, _ := strconv.ParseInt(c.Query("combatant"), 10, 64)
	f := views.EffectFields{ID: views.EffectID(characterID, combatantID), URL: views.EffectFieldsURL(id, characterID, combatantID)}
	switch {
	case characterID != 0:
		ch, err := a.Store.GetCharacter(ctx, characterID)
		var allowed bool
		if err == nil {
			allowed, err = a.canAddEffect(ctx, w.viewer, ch)
		}
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return err
		}
		if err != nil || ch.CampaignID != id || !allowed {
			a.fail(c, http.StatusForbidden, "That is not your character.")
			return nil
		}
		// An ally's sheet stays with its player: no weapon names reach this one.
		f.NeedsHolder = !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeConditions)
		if !f.NeedsHolder {
			if f.Sheet, err = dnd5e.ParseSheet(ch.SheetJson); err != nil {
				return err
			}
		}
		if w.viewer.IsPlayer() && !f.NeedsHolder {
			if f.Targets, err = a.allyTargets(ctx, ch); err != nil {
				return err
			}
			f.Target, _ = strconv.ParseInt(c.Query("target"), 10, 64)
		}
		_, f.InFight, err = a.characterFight(ctx, ch.ID)
		if err != nil {
			return err
		}
	case combatantID != 0 && w.isGM():
		cb, err := a.Store.GetCombatant(ctx, combatantID)
		if err != nil {
			return err
		}
		enc, err := a.Store.GetEncounter(ctx, cb.EncounterID)
		if err != nil || enc.CampaignID != id {
			a.fail(c, http.StatusForbidden, "That is not in this campaign.")
			return nil
		}
		f.InFight = enc.Status == enum.EncounterActive
	default:
		a.fail(c, http.StatusForbidden, "Pick a character or combatant to add the effect to.")
		return nil
	}
	if c.Query("presets") == "1" {
		names, err := a.presetNames(ctx, id, w.viewer.CharacterID)
		if err != nil {
			return err
		}
		a.render(c, http.StatusOK, views.EffectPresetPicker(id, names))
		return nil
	}
	if row := c.Query("row"); row != "" {
		i, err := strconv.Atoi(row)
		if err != nil || i < 1 || i >= dnd5e.MaxEffectMods {
			return userError{fmt.Errorf("an effect takes at most %d changes", dnd5e.MaxEffectMods)}
		}
		a.render(c, http.StatusOK, views.EffectModRow(f, i))
		return nil
	}
	var err error
	if f.HeldBy, err = a.heldByOptions(ctx, w.viewer, id, characterID, combatantID); err != nil {
		return err
	}
	if presetID, _ := strconv.ParseInt(strings.TrimSpace(c.Query("preset")), 10, 64); presetID != 0 {
		p, err := a.Store.GetEffectPreset(ctx, presetID)
		if err != nil {
			return err
		}
		if p.CampaignID != id || !perm.CanUsePreset(w.viewer, p.OwnerCharacterID.Int64) {
			a.fail(c, http.StatusForbidden, "That is not your preset.")
			return nil
		}
		f.Name, f.Note, f.Slug, f.EndsOn, f.Rounds = p.Name, p.Note, p.Slug, p.EndsOn, int(p.Rounds)
		if err := json.Unmarshal([]byte(p.ModsJson), &f.Mods); err != nil {
			return err
		}
		if !f.InFight {
			f.Rounds = 0 // rounds count only in a running fight
		}
	} else if spell := strings.TrimSpace(c.Query("spell")); spell != "" {
		var d compendium.SpellData
		if f.Name, f.Note, f.Rounds, d, err = a.spellPrefill(ctx, spell, id); err != nil {
			return err
		}
		if !f.InFight {
			f.Rounds = 0 // rounds count only in a running fight
		}
		var hints []string
		if d.Concentration {
			hints = append(hints, "Concentration — pick the caster below")
		}
		slot, _ := strconv.Atoi(c.Query("slot"))
		f.Slot = min(max(slot, d.Level), 9)
		var dice string
		f.Mods, dice = dnd5e.SpellMods(d.Desc, d.HigherLevel, d.Level, f.Slot)
		if note := f.Note + ", " + dice; dice != "" && len([]rune(note)) <= effectNoteMax {
			f.Note = note
		}
		if len(f.Mods) > 0 || dice != "" {
			hints = append(hints, "From the spell text — check before adding")
		}
		if d.Level > 0 && d.HigherLevel != "" && len(f.Mods) > 0 {
			f.SpellLevel = d.Level // the slot can change a parsed bonus: offer it
		}
		f.Hint = strings.Join(hints, ". ")
	}
	a.render(c, http.StatusOK, views.EffectFieldsForm(f))
	return nil
}

// allyTargets are the characters a player may add an effect to from their own
// sheet: themselves first, then the others of the campaign.
func (a *App) allyTargets(ctx context.Context, own sqlc.Character) ([]views.Named, error) {
	rows, err := a.Store.ListCharacterNames(ctx, own.CampaignID)
	if err != nil {
		return nil, err
	}
	out := []views.Named{{ID: own.ID, Name: own.Name}}
	for _, r := range rows {
		if r.ID != own.ID {
			out = append(out, views.Named{ID: r.ID, Name: r.Name})
		}
	}
	return out, nil
}

// characterFight is the running fight a character is in, if any.
func (a *App) characterFight(ctx context.Context, characterID int64) (sqlc.Encounter, bool, error) {
	ids, err := a.Store.ListEncounterIDsByCharacter(ctx, nullID(characterID))
	if err != nil {
		return sqlc.Encounter{}, false, err
	}
	for _, id := range ids {
		enc, err := a.Store.GetEncounter(ctx, id)
		if err != nil {
			return enc, false, err
		}
		if enc.Status == enum.EncounterActive {
			return enc, true, nil
		}
	}
	return sqlc.Encounter{}, false, nil
}

// heldLabel is what a concentration row is called when offered as a holder.
func heldLabel(owner, note string) string {
	label := cmp.Or(note, slugConcentration)
	if owner == "" {
		return label
	}
	return owner + ": " + label
}

// heldLabels name what holds the effects among conds, by concentration row
// id, for showing. A combatant's name is the GM's: a player sees the spell.
func (a *App) heldLabels(ctx context.Context, v perm.Viewer, conds []sqlc.Condition) (map[int64]string, error) {
	var out map[int64]string
	for _, c := range conds {
		if !c.HeldBy.Valid || out[c.HeldBy.Int64] != "" {
			continue
		}
		h, err := a.Store.GetCondition(ctx, c.HeldBy.Int64)
		if errors.Is(err, sql.ErrNoRows) {
			continue
		}
		if err != nil {
			return nil, err
		}
		var owner string
		switch {
		case h.CharacterID.Valid:
			ch, err := a.Store.GetCharacter(ctx, h.CharacterID.Int64)
			if err != nil {
				return nil, err
			}
			owner = ch.Name
		case v.IsGM:
			cb, err := a.Store.GetCombatant(ctx, h.CombatantID.Int64)
			if err != nil {
				return nil, err
			}
			owner = cb.Name
		}
		if out == nil {
			out = map[int64]string{}
		}
		out[h.ID] = heldLabel(owner, h.Note)
	}
	return out, nil
}

// spellPrefill is what a compendium spell fills into the effect form: its
// name, a clock duration as the note (a clock never auto-expires), the rounds
// a duration in rounds or minutes lasts in a fight, and whether it needs
// concentration (in d, with the text SpellMods reads). spell is its name or slug.
func (a *App) spellPrefill(ctx context.Context, spell string, campaignID int64) (name, note string, rounds int, d compendium.SpellData, err error) {
	slug := compendium.Slugify(spell)
	row, err := a.entryBySlug(ctx, enum.KindSpell, slug, campaignID)
	if errors.Is(err, sql.ErrNoRows) {
		return "", "", 0, d, userError{fmt.Errorf("no spell %q in the compendium", slug)}
	}
	if err != nil {
		return "", "", 0, d, err
	}
	dp, err := compendium.Decode[compendium.SpellData](enum.KindSpell, row.DataJson)
	if err != nil {
		return "", "", 0, d, err
	}
	d = *dp
	// The note keeps the duration even with rounds: outside a fight it is all there is.
	return row.Name, d.Duration, durationRounds(d.Duration), d, nil
}

// durationRounds is a spell duration in rounds: 10 a minute. Only "N
// minute(s)" or "N round(s)", after an optional "Concentration, " and "up to "
// (5etools keeps concentration as a flag), counts; anything else (an hour,
// until dispelled) is 0.
func durationRounds(duration string) int {
	d := strings.ToLower(strings.TrimSpace(duration))
	d = strings.TrimPrefix(d, "concentration, ")
	d = strings.TrimSpace(strings.TrimPrefix(d, "up to "))
	n, unit, ok := strings.Cut(d, " ")
	count, err := strconv.Atoi(n)
	if !ok || err != nil || count <= 0 {
		return 0
	}
	switch unit {
	case "minute", "minutes":
		return count * 10
	case "round", "rounds":
		return count
	}
	return 0
}

// heldByOptions are the concentration rows v may pick as what holds an
// effect on a character (characterID) or a combatant (combatantID): the GM
// any character's in the campaign and those of the target's fights, a player
// their own. The label is "Caster: spell".
func (a *App) heldByOptions(ctx context.Context, v perm.Viewer, campaignID, characterID, combatantID int64) ([]views.Named, error) {
	label := heldLabel
	var out []views.Named
	if !v.IsGM {
		if v.CharacterID == 0 {
			return nil, nil
		}
		ch, err := a.Store.GetCharacter(ctx, v.CharacterID)
		if err != nil || ch.CampaignID != campaignID {
			return nil, notFoundIsNil(err) // no row means no options
		}
		cond, err := a.Store.FindCharacterCondition(ctx, sqlc.FindCharacterConditionParams{CharacterID: nullID(ch.ID), Slug: slugConcentration})
		if err != nil {
			return nil, notFoundIsNil(err)
		}
		return []views.Named{{ID: cond.ID, Name: label(ch.Name, cond.Note)}}, nil
	}
	chars, err := a.Store.ListCharacterConcentration(ctx, campaignID)
	if err != nil {
		return nil, err
	}
	for _, r := range chars {
		out = append(out, views.Named{ID: r.ID, Name: label(r.OwnerName, r.Note)})
	}
	var encounters []int64
	if combatantID != 0 {
		cb, err := a.Store.GetCombatant(ctx, combatantID)
		if err != nil {
			return nil, err
		}
		encounters = []int64{cb.EncounterID}
	} else if encounters, err = a.Store.ListEncounterIDsByCharacter(ctx, nullID(characterID)); err != nil {
		return nil, err
	}
	for _, e := range encounters {
		rows, err := a.Store.ListEncounterConcentration(ctx, e)
		if err != nil {
			return nil, err
		}
		for _, r := range rows {
			out = append(out, views.Named{ID: r.ID, Name: label(r.OwnerName, r.Note)})
		}
	}
	return out, nil
}
