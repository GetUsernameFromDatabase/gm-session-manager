package web

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/compendium"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

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

func (a *App) characterPage(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), ch.CampaignID)
	if !a.found(c, err) {
		return
	}
	l := a.layout(w, ch.Name, views.NavCharacter, &camp, true)
	s, err := a.sheetView(c.Request.Context(), w.viewer, ch)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.CharacterPage(l, s))
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
}

// hpCommand turns the HP quick form into damage, healing or temp HP.
func hpCommand(c *gin.Context) (dnd5e.Command, error) {
	var f hpForm
	if err := c.ShouldBind(&f); err != nil {
		return dnd5e.Command{}, errors.New("enter how many hit points, 1 or more")
	}
	return dnd5e.Parse(map[string]string{"damage": "hp -", "heal": "hp +", "temp": "temp "}[f.Op] + strconv.Itoa(f.Amount))
}

// builderArgs are the command builder's fields, in the order the commands
// take them (see views.CommandBuilder). from and fromnote name the thing an
// edit changes.
var builderArgs = []string{"from", "fromnote", "subject", "ability", "dice", "dtype", "prof", "n", "text"}

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
		if f != "text" && f != "fromnote" {
			v = strings.TrimSpace(v)
		}
		if f == "n" && v != "" {
			v = c.PostForm("sign") + v
		}
		// An empty text still counts: it clears a backstory.
		if ok && (v != "" || f == "text") {
			raw = append(raw, v)
		}
	}
	return dnd5e.Parse(prefix, raw...)
}

// commandBuilder renders one form of the command builder dialog.
func (a *App) commandBuilder(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player changes their character.")
		return
	}
	s, err := a.sheetView(c.Request.Context(), w.viewer, ch)
	if err != nil {
		a.serverError(c, err)
		return
	}
	x, err := a.builderExtra(c)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, http.StatusOK, views.CommandBuilder(s, c.Query("verb"), c.Query("op"), c.Query("subject"), x))
}

// characterCommand applies one command in a transaction and pushes the
// result to everyone who can see the character.
func (a *App) characterCommand(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	cmd, err := parseCommand(c)
	if err != nil {
		a.fail(c, http.StatusUnprocessableEntity, err.Error())
		return
	}
	if !perm.CanCharacter(w.viewer, ch.ID, cmd.Scope) {
		a.fail(c, http.StatusForbidden, "The GM may change HP and conditions only.")
		return
	}
	if a.commandFailed(c, a.runCommand(c.Request.Context(), ch.ID, cmd)) {
		return
	}
	a.characterChanged(c, ch)
}

// runCommand looks up what cmd needs from the compendium and applies it to
// the character in one transaction. Errors the user caused are userErrors.
func (a *App) runCommand(ctx context.Context, characterID int64, cmd dnd5e.Command) error {
	var env dnd5e.Env
	var err error
	if cmd.SpellSlug != "" {
		if env.Spell, err = a.compendiumSpell(ctx, cmd.SpellSlug); err != nil {
			return userError{err}
		}
	}
	if cmd.ArmorSlug != "" {
		inInventory, err := a.hasArmorItem(ctx, characterID, cmd.ArmorSlug)
		if err != nil {
			return err
		}
		if !inInventory {
			if env.Armor, env.IsShield, err = a.compendiumArmor(ctx, cmd.ArmorSlug); err != nil {
				return userError{err}
			}
		}
	}
	return a.updateCharacter(ctx, characterID, func(q *sqlc.Queries, dc *dnd5e.Character) error {
		fx, err := cmd.Apply(dc, env)
		if err != nil {
			return userError{err}
		}
		return applyCondOps(ctx, q, condOwner{characterID: characterID}, fx.Conditions, sql.NullInt64{})
	})
}

// userError marks an error caused by the request, shown to the user as is.
type userError struct{ error }

func (a *App) commandFailed(c *gin.Context, err error) bool {
	var ue userError
	switch {
	case err == nil:
		return false
	case errors.As(err, &ue):
		a.fail(c, http.StatusUnprocessableEntity, ue.Error())
	default:
		a.serverError(c, err)
	}
	return true
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

// hasArmorItem reports whether the character carries armor (or a shield)
// named name, which it then wears without a compendium lookup.
func (a *App) hasArmorItem(ctx context.Context, characterID int64, name string) (bool, error) {
	row, err := a.Store.GetCharacter(ctx, characterID)
	if err != nil {
		return false, err
	}
	sheet, err := dnd5e.ParseSheet(row.SheetJson)
	if err != nil {
		return false, err
	}
	return sheet.ArmorItem(name) >= 0, nil
}

// compendiumArmor finds armor (or a shield) in the compendium by slug.
func (a *App) compendiumArmor(ctx context.Context, slug string) (dnd5e.WornArmor, bool, error) {
	row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindItem), Slug: slug})
	if errors.Is(err, sql.ErrNoRows) {
		return dnd5e.WornArmor{}, false, fmt.Errorf("no item %q in the compendium", slug)
	}
	if err != nil {
		return dnd5e.WornArmor{}, false, err
	}
	d, err := compendium.DecodeData(compendium.KindItem, row.DataJson)
	if err != nil {
		return dnd5e.WornArmor{}, false, err
	}
	ar := d.(*compendium.ItemData).Armor
	switch {
	case ar == nil:
		return dnd5e.WornArmor{}, false, fmt.Errorf("%s is not armor", row.Name)
	case ar.Type == compendium.ArmorShield:
		return dnd5e.WornArmor{}, true, nil
	}
	return dnd5e.WornArmor{Slug: row.Slug, Name: row.Name, Type: ar.Type, Base: ar.Base, AddDex: ar.AddDex, DexCap: ar.DexCap, StrReq: ar.StrReq, Stealth: ar.Stealth}, false, nil
}

func (a *App) compendiumSpell(ctx context.Context, slug string) (dnd5e.Spell, error) {
	row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindSpell), Slug: slug})
	if errors.Is(err, sql.ErrNoRows) {
		return dnd5e.Spell{}, fmt.Errorf("no spell %q in the compendium", slug)
	}
	if err != nil {
		return dnd5e.Spell{}, err
	}
	d, err := compendium.DecodeData(compendium.KindSpell, row.DataJson)
	if err != nil {
		return dnd5e.Spell{}, err
	}
	return dnd5e.Spell{Slug: row.Slug, Name: row.Name, Level: d.(*compendium.SpellData).Level}, nil
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
			level = min(max(level+op.Delta, 0), 6)
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
func (a *App) deleteCondition(c *gin.Context) {
	id, ok := a.idParam(c, "condition")
	if !ok {
		return
	}
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return
	}
	ctx := c.Request.Context()
	cond, err := a.Store.GetCondition(ctx, id)
	if !a.found(c, err) {
		return
	}
	if cond.CharacterID.Valid {
		if !perm.CanCharacter(w.viewer, cond.CharacterID.Int64, perm.ScopeConditions) {
			a.fail(c, http.StatusForbidden, "Not your character.")
			return
		}
		ch, err := a.Store.GetCharacter(ctx, cond.CharacterID.Int64)
		if !a.found(c, err) {
			return
		}
		if err := a.Store.DeleteCondition(ctx, id); err != nil {
			a.serverError(c, err)
			return
		}
		a.characterChanged(c, ch)
		return
	}
	if !w.isGM() {
		a.fail(c, http.StatusForbidden, "Only the GM changes monsters.")
		return
	}
	cb, err := a.Store.GetCombatant(ctx, cond.CombatantID.Int64)
	if !a.found(c, err) {
		return
	}
	enc, err := a.Store.GetEncounter(ctx, cb.EncounterID)
	if !a.found(c, err) {
		return
	}
	if err := a.Store.DeleteCondition(ctx, id); err != nil {
		a.serverError(c, err)
		return
	}
	a.changed(c, enc.CampaignID, hub.Topic{Kind: hub.TopicEncounter, ID: enc.ID})
}

// classInfo converts a compendium class for the level-up helper.
func classInfo(row sqlc.CompendiumEntry) (dnd5e.ClassInfo, error) {
	d, err := compendium.DecodeData(compendium.KindClass, row.DataJson)
	if err != nil {
		return dnd5e.ClassInfo{}, err
	}
	cd := d.(*compendium.ClassData)
	info := dnd5e.ClassInfo{Slug: row.Slug, Name: row.Name, HitDie: cd.HitDie, Caster: cd.Caster, Saves: cd.Saves, Table: cd.Table, Features: classFeatures(cd)}
	if info.HitDie == 0 {
		info.HitDie = 8 // homebrew classes without a hit die
	}
	return info, nil
}

func classFeatures(cd *compendium.ClassData) []dnd5e.ClassFeature {
	out := make([]dnd5e.ClassFeature, 0, len(cd.Features))
	for _, f := range cd.Features {
		out = append(out, dnd5e.ClassFeature{Name: f.Name, Description: f.Desc, Levels: f.Levels, Subclass: f.Subclass})
	}
	return out
}

// withSubclass adds the features of the class's subclass named name (the
// sheet's, or the one being chosen) to info; an unknown name adds none.
func (a *App) withSubclass(ctx context.Context, info *dnd5e.ClassInfo, sheet dnd5e.Sheet, name string) error {
	if name == "" {
		for _, cl := range sheet.Classes {
			if cl.Slug == info.Slug {
				name = cl.Subclass
			}
		}
	}
	if name == "" {
		return nil
	}
	subs, err := a.Store.ListSubclassesOf(ctx, info.Slug)
	if err != nil {
		return err
	}
	i := slices.IndexFunc(subs, func(s sqlc.ListSubclassesOfRow) bool { return strings.EqualFold(s.Name, name) })
	if i < 0 {
		return nil
	}
	row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindSubclass), Slug: subs[i].Slug})
	if err != nil {
		return err
	}
	d, err := compendium.DecodeData(compendium.KindSubclass, row.DataJson)
	if err != nil {
		return err
	}
	info.SubclassFeatures = classFeatures(d.(*compendium.ClassData))
	return nil
}

// levelUpView previews a level in classSlug; with pickDefault it picks the
// class the character would most likely level instead.
func (a *App) levelUpView(ctx context.Context, ch sqlc.Character, classSlug string, pickDefault bool) (views.LevelUpView, error) {
	sheet, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		return views.LevelUpView{}, err
	}
	if pickDefault {
		classSlug = sheet.DefaultClass()
	}
	v := views.LevelUpView{CharacterID: ch.ID, Name: ch.Name, Level: sheet.Level(), Selected: classSlug, Sheet: sheet}
	classes, err := a.Store.ListCompendiumNames(ctx, string(compendium.KindClass))
	if err != nil {
		return v, err
	}
	for _, cl := range classes {
		v.Classes = append(v.Classes, views.Option{Slug: cl.Slug, Name: cl.Name})
	}
	if classSlug == "" {
		return v, nil
	}
	row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindClass), Slug: classSlug})
	if err != nil {
		return v, err
	}
	info, err := classInfo(row)
	if err != nil {
		return v, err
	}
	if err := a.withSubclass(ctx, &info, sheet, ""); err != nil {
		return v, err
	}
	plan, err := sheet.Plan(info)
	if err != nil {
		v.Error = err.Error()
		return v, nil
	}
	v.Plan = &plan
	subs, err := a.Store.ListSubclassesOf(ctx, classSlug)
	if err != nil {
		return v, err
	}
	for _, s := range subs {
		v.Subclasses = append(v.Subclasses, views.Option{Slug: s.Slug, Name: s.Name})
	}
	feats, err := a.Store.ListCompendiumNames(ctx, string(compendium.KindFeat))
	if err != nil {
		return v, err
	}
	for _, f := range feats {
		v.Feats = append(v.Feats, views.Option{Slug: f.Slug, Name: f.Name})
	}
	maxLevel := 0
	for i, n := range plan.Slots {
		if n > 0 {
			maxLevel = i + 1
		}
	}
	maxLevel = max(maxLevel, plan.PactLevel)
	spells, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: classSlug, MaxLevel: int64(maxLevel)})
	if err != nil {
		return v, err
	}
	for _, s := range spells {
		known := false
		for _, have := range sheet.Spells {
			known = known || have.Slug == s.Slug
		}
		if !known {
			v.Spells = append(v.Spells, views.Option{Slug: s.Slug, Name: s.Name, Level: int(s.Level)})
		}
	}
	return v, nil
}

func (a *App) levelUpForm(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player levels up their character.")
		return
	}
	ctx := c.Request.Context()
	slug, picked := c.GetQuery("class")
	v, err := a.levelUpView(ctx, ch, slug, !picked)
	if errors.Is(err, sql.ErrNoRows) {
		v.Error = "That class is not in the compendium."
	} else if err != nil {
		a.serverError(c, err)
		return
	}
	camp, err := a.Store.GetCampaign(ctx, ch.CampaignID)
	if !a.found(c, err) {
		return
	}
	a.render(c, http.StatusOK, views.LevelUpPage(a.layout(w, "Level up", views.NavCharacter, &camp, false), v))
}

func (a *App) levelUp(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player levels up their character.")
		return
	}
	ctx := c.Request.Context()
	row, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindClass), Slug: c.PostForm("class")})
	if !a.found(c, err) {
		return
	}
	info, err := classInfo(row)
	if err != nil {
		a.serverError(c, err)
		return
	}
	sheet, err := dnd5e.ParseSheet(ch.SheetJson)
	if err != nil {
		a.serverError(c, err)
		return
	}
	choice := dnd5e.LevelUpChoice{Class: info, Subclass: strings.TrimSpace(c.PostForm("subclass"))}
	if err := a.withSubclass(ctx, &choice.Class, sheet, choice.Subclass); err != nil {
		a.serverError(c, err)
		return
	}
	if roll, ok := formInt(c, "hp_roll"); ok {
		choice.HPRoll = int(roll)
	}
	for _, f := range []string{"asi1", "asi2"} {
		if v := c.PostForm(f); v != "" {
			choice.ASI = append(choice.ASI, v)
		}
	}
	if slug := c.PostForm("feat"); slug != "" {
		feat, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindFeat), Slug: slug})
		if !a.found(c, err) {
			return
		}
		choice.Feat, choice.FeatDesc = feat.Name, feat.BodyText
	}
	for _, slug := range c.PostFormArray("spell") {
		sp, err := a.compendiumSpell(ctx, slug)
		if err != nil {
			a.fail(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		choice.Spells = append(choice.Spells, sp)
	}

	err = a.updateCharacter(ctx, ch.ID, func(q *sqlc.Queries, dc *dnd5e.Character) error {
		before, err := dc.Sheet.JSON()
		if err != nil {
			return err
		}
		maxHP := dc.MaxHP
		if err := dc.LevelUp(choice); err != nil {
			return userError{err}
		}
		return q.CreateLevelUp(ctx, sqlc.CreateLevelUpParams{CharacterID: ch.ID, MaxHp: int64(maxHP), SheetJson: before})
	})
	var ue userError
	if errors.As(err, &ue) {
		v, verr := a.levelUpView(ctx, ch, info.Slug, false)
		if verr != nil {
			a.serverError(c, verr)
			return
		}
		v.Error = ue.Error()
		camp, _ := a.Store.GetCampaign(ctx, ch.CampaignID)
		a.render(c, http.StatusUnprocessableEntity, views.LevelUpPage(a.layout(w, "Level up", views.NavCharacter, &camp, false), v))
		return
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.Hub.Publish(ch.CampaignID, hub.Topic{Kind: hub.TopicCharacter, ID: ch.ID})
	c.Redirect(http.StatusSeeOther, fmt.Sprintf("/characters/%d", ch.ID))
}

// revertLevelUp puts back the sheet and max HP from before the latest
// level-up. Current HP is kept, up to that max.
func (a *App) revertLevelUp(c *gin.Context) {
	w, ch, ok := a.characterAccess(c)
	if !ok {
		return
	}
	if !perm.CanCharacter(w.viewer, ch.ID, perm.ScopeSheet) {
		a.fail(c, http.StatusForbidden, "Only the player reverts a level-up.")
		return
	}
	ctx := c.Request.Context()
	err := a.updateCharacter(ctx, ch.ID, func(q *sqlc.Queries, dc *dnd5e.Character) error {
		snap, err := q.LatestLevelUp(ctx, ch.ID)
		if errors.Is(err, sql.ErrNoRows) {
			return userError{errors.New("there is no level-up to revert")}
		}
		if err != nil {
			return err
		}
		if dc.Sheet, err = dnd5e.ParseSheet(snap.SheetJson); err != nil {
			return err
		}
		dc.MaxHP = int(snap.MaxHp)
		dc.HP = min(dc.HP, dc.MaxHP)
		return q.DeleteLevelUp(ctx, snap.ID)
	})
	if a.commandFailed(c, err) {
		return
	}
	a.characterChanged(c, ch)
}
