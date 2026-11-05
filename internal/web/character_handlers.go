package web

import (
	"cmp"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/perm"
	"gm-session-manager/internal/web/views"
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
	x, err := a.builderExtra(c, ch.CampaignID)
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
	if a.commandFailed(c, a.runCommand(c.Request.Context(), ch, cmd)) {
		return
	}
	a.characterChanged(c, ch)
}

// runCommand looks up what cmd needs from the compendium and applies it to
// the character in one transaction. Errors the user caused are userErrors.
func (a *App) runCommand(ctx context.Context, ch sqlc.Character, cmd dnd5e.Command) error {
	var env dnd5e.Env
	var err error
	if cmd.SpellSlug != "" {
		if env.Spell, err = a.compendiumSpell(ctx, cmd.SpellSlug, ch.CampaignID); err != nil {
			return userError{err}
		}
	}
	if cmd.ArmorSlug != "" {
		inInventory, err := a.hasArmorItem(ctx, ch.ID, cmd.ArmorSlug)
		if err != nil {
			return err
		}
		if !inInventory {
			if env.Armor, env.IsShield, err = a.compendiumArmor(ctx, cmd.ArmorSlug, ch.CampaignID); err != nil {
				return userError{err}
			}
		}
	}
	return a.updateCharacter(ctx, ch.ID, func(q *sqlc.Queries, dc *dnd5e.Character) error {
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

// userError marks an error caused by the request, shown to the user as is.
type userError struct{ error }

// fieldError is a userError about one form field, shown next to it.
type fieldError struct {
	userError
	field string
}

func (e fieldError) Unwrap() error { return e.userError }

// errorField is the form field err is about, "" for none.
func errorField(err error) string {
	var fe fieldError
	errors.As(err, &fe)
	return fe.field
}

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
func (a *App) compendiumArmor(ctx context.Context, slug string, campaignID int64) (dnd5e.WornArmor, bool, error) {
	row, err := a.entryBySlug(ctx, compendium.KindItem, slug, campaignID)
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

func (a *App) compendiumSpell(ctx context.Context, slug string, campaignID int64) (dnd5e.Spell, error) {
	row, err := a.entryBySlug(ctx, compendium.KindSpell, slug, campaignID)
	if errors.Is(err, sql.ErrNoRows) {
		return dnd5e.Spell{}, userError{fmt.Errorf("no spell %q in the compendium", slug)}
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
	v.FeatChoices = map[string][]views.ChoiceGroup{}
	for _, f := range feats {
		if !plan.ASI || f.Name == dnd5e.ASIFeature {
			continue // the form's ability selects are the ASI feat
		}
		td, err := a.textEntry(ctx, compendium.KindFeat, f.Slug)
		if err != nil {
			return v, err
		}
		if td == nil || !meetsPrerequisite(td.Prerequisite, sheet, plan.NewLevel, nil) {
			continue
		}
		v.Feats = append(v.Feats, views.Option{Slug: f.Slug, Name: f.Name})
		groups, err := a.featChoices(ctx, "feat_"+f.Slug, f.Name, td, "", sheet, ch.CampaignID)
		if err != nil {
			return v, err
		}
		if len(groups) > 0 {
			v.FeatChoices[f.Slug] = groupViews(groups)
		}
	}
	if v.LineageSpells, err = a.lineageSpells(ctx, sheet, plan.NewLevel, ch.CampaignID); err != nil {
		return v, err
	}
	changes, err := a.featSpellChanges(ctx, sheet, v.LineageSpells)
	if err != nil {
		return v, err
	}
	cantrips, err := a.optionCantrips(ctx, sheet, info, plan.ClassLevel, ch.CampaignID)
	if err != nil {
		return v, err
	}
	v.Groups = groupViews(append(classOptionGroups(info, plan.ClassLevel, plan.ClassLevel-1, sheet, cantrips), changes...))
	v.Spells, err = a.classSpellOptions(ctx, classSlug, plan.Slots, plan.PactLevel, sheet, v.LineageSpells)
	v.SpellNeed = spellNeed(info, plan.ClassLevel, sheet, v.Spells)
	return v, err
}

// classSpellOptions are the class's spells the slots (and pact slots) can
// cast that the sheet does not have yet, nor learns in this step (taught:
// lineage spells, which AddSpell would keep over the class pick).
func (a *App) classSpellOptions(ctx context.Context, classSlug string, slots [9]int, pactLevel int, sheet dnd5e.Sheet, taught []dnd5e.Spell) ([]views.Option, error) {
	maxLevel := pactLevel
	for i, n := range slots {
		if n > 0 {
			maxLevel = max(maxLevel, i+1)
		}
	}
	spells, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: classSlug, MaxLevel: int64(maxLevel)})
	if err != nil {
		return nil, err
	}
	known := slices.Concat(sheet.Spells, taught)
	var out []views.Option
	for _, s := range spells {
		if !slices.ContainsFunc(known, func(have dnd5e.Spell) bool { return have.Slug == s.Slug }) {
			out = append(out, views.Option{Slug: s.Slug, Name: s.Name, Level: int(s.Level)})
		}
	}
	return out, nil
}

var (
	bookStart    = regexp.MustCompile(`(?i)starts with (\w+) level 1 \w+ spells`)
	bookPerLevel = regexp.MustCompile(`(?i)level after 1, add (\w+) \w+ spells`)
)

// spellbook is how many level 1+ spells the class's spellbook holds at a
// class level ("starts with six level 1 Wizard spells … add two"), or 0 if
// its features describe none.
func spellbook(info dnd5e.ClassInfo, level int) int {
	count := func(re *regexp.Regexp) int {
		for _, f := range info.Features {
			if m := re.FindStringSubmatch(f.Description); m != nil {
				n, _ := strconv.Atoi(cmp.Or(compendium.CountWords[strings.ToLower(m[1])], m[1]))
				return n
			}
		}
		return 0
	}
	start := count(bookStart)
	if start == 0 {
		return 0
	}
	return start + count(bookPerLevel)*(level-1)
}

// spellNeed is what the class table's cantrip and spell totals (at the class
// level being reached) leave to pick: each total less the class's spells on
// the sheet, at most what is offered. A spellbook's size replaces the spell
// total, exactly: its list is prepared from the book. Lineage and feat spells
// (their own ability or feat) do not count; untagged spells of older sheets
// count while the sheet has this class only.
func spellNeed(info dnd5e.ClassInfo, level int, sheet dnd5e.Sheet, offered []views.Option) views.SpellNeed {
	class, book := info.Slug, spellbook(info, level)
	has := slices.ContainsFunc(sheet.Classes, func(cl dnd5e.ClassLevel) bool { return cl.Slug == class })
	only := has && len(sheet.Classes) == 1
	var have, avail [2]int // cantrips, level 1+
	unsure := false        // untagged spells that may be this class's
	for _, sp := range sheet.Spells {
		if sp.Feat != "" || sp.Ability != "" {
			continue
		}
		if sp.Class == class || sp.Class == "" && only {
			have[min(sp.Level, 1)]++
		}
		unsure = unsure || sp.Class == "" && has && !only
	}
	for _, o := range offered {
		avail[min(o.Level, 1)]++
	}
	// Without knowing which class taught them, the totals can only be upper limits.
	need := views.SpellNeed{Cantrips: -1, Spells: -1, CantripsUpTo: unsure, SpellsUpTo: unsure}
	for col, v := range info.ColumnsAt(level) {
		n, err := strconv.Atoi(v)
		cantrips, spells := views.SpellColumn(col)
		switch {
		case err != nil:
		case cantrips:
			need.Cantrips = min(max(n-have[0], 0), avail[0])
		case spells:
			need.Spells = min(max(n-have[1], 0), avail[1])
			need.SpellsUpTo = unsure || strings.Contains(strings.ToLower(col), "prepared")
		}
	}
	if book > 0 {
		need.Spells, need.SpellsUpTo, need.Book = min(max(book-have[1], 0), avail[1]), unsure, true
	}
	return need
}

// classSpellPicks are the spells ticked in the form for class, refused
// unless offered and as many as need asks.
func (a *App) classSpellPicks(c *gin.Context, class string, offered []views.Option, need views.SpellNeed, campaignID int64) ([]dnd5e.Spell, error) {
	var spells []dnd5e.Spell
	var got [2]int
	for _, slug := range c.PostFormArray("spell") {
		if !slices.ContainsFunc(offered, func(o views.Option) bool { return o.Slug == slug }) || slices.ContainsFunc(spells, func(sp dnd5e.Spell) bool { return sp.Slug == slug }) {
			return nil, fieldError{userError{errors.New("choose spells from the list")}, "spell"}
		}
		sp, err := a.compendiumSpell(c.Request.Context(), slug, campaignID)
		if err != nil {
			return nil, err
		}
		sp.Class = class
		spells = append(spells, sp)
		got[min(sp.Level, 1)]++
	}
	off := func(n, got int, upTo bool) bool { return n >= 0 && (got > n || !upTo && got < n) }
	if off(need.Cantrips, got[0], need.CantripsUpTo) || off(need.Spells, got[1], need.SpellsUpTo) {
		return nil, fieldError{userError{fmt.Errorf("choose %s: you ticked %s and %s", cmp.Or(need.String(), "no spells here"), views.Plural(got[0], "cantrip"), views.Plural(got[1], "level 1+ spell"))}, "spell"}
	}
	return spells, nil
}

// addClassSpells adds the class picks after the lineage's and feats' spells
// of the same step, refusing one those taught already: it would not count as
// a class pick.
func addClassSpells(s *dnd5e.Sheet, spells []dnd5e.Spell) error {
	for _, sp := range spells {
		if slices.ContainsFunc(s.Spells, func(have dnd5e.Spell) bool { return have.Slug == sp.Slug }) {
			return fieldError{userError{fmt.Errorf("%s comes from your species or feat already: choose another spell", sp.Name)}, "spell"}
		}
		s.AddSpell(sp)
	}
	return nil
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
	var featErr error
	if slug := c.PostForm("feat"); slug != "" {
		feat, err := a.Store.GetCompendiumEntryBySlug(ctx, sqlc.GetCompendiumEntryBySlugParams{Kind: string(compendium.KindFeat), Slug: slug})
		switch {
		case errors.Is(err, sql.ErrNoRows): // crafted, or a homebrew feat deleted since the form loaded
			featErr = userError{errors.New("choose a feat from the list")}
		case !a.found(c, err):
			return
		default:
			choice.Feat, choice.FeatDesc = feat.Name, feat.BodyText
		}
	}
	// A plan error (level 20) comes back from LevelUp below.
	plan, planErr := sheet.Plan(choice.Class)
	// The class list does not offer the lineage's spells of this level.
	lineage, err := a.lineageSpells(ctx, sheet, plan.NewLevel, ch.CampaignID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	var spells []dnd5e.Spell
	var spellErr error
	var ue userError
	if planErr == nil { // else LevelUp refuses (level 20) and plan is empty
		offered, err := a.classSpellOptions(ctx, info.Slug, plan.Slots, plan.PactLevel, sheet, lineage)
		if err != nil {
			a.serverError(c, err)
			return
		}
		// A refusal waits for readPicks, so the page re-renders as for any other.
		spells, spellErr = a.classSpellPicks(c, info.Slug, offered, spellNeed(info, plan.ClassLevel, sheet, offered), ch.CampaignID)
		if spellErr != nil && !errors.As(spellErr, &ue) {
			a.serverError(c, spellErr)
			return
		}
	}
	// The class picks go in after the feats' (addClassSpells), so a clash is refused.
	choice.Spells = lineage

	// Spell Change and the feat's own choices apply after LevelUp has added the feat.
	changes, err := a.featSpellChanges(ctx, sheet, lineage)
	if err != nil {
		a.serverError(c, err)
		return
	}
	cantrips, err := a.optionCantrips(ctx, sheet, choice.Class, plan.ClassLevel, ch.CampaignID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	// The cantrips the sheet will know, for checkNewOptions.
	learned, err := a.cantripData(ctx, slices.Concat(sheet.Spells, lineage, spells), ch.CampaignID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	groups := append(classOptionGroups(choice.Class, plan.ClassLevel, plan.ClassLevel-1, sheet, cantrips), changes...)
	if choice.Feat != "" {
		slug := c.PostForm("feat")
		td, err := a.textEntry(ctx, compendium.KindFeat, slug)
		if err != nil {
			a.serverError(c, err)
			return
		}
		if planErr == nil && !meetsPrerequisite(td.Prerequisite, sheet, plan.NewLevel, nil) {
			featErr = userError{fmt.Errorf("%s needs %s", choice.Feat, td.Prerequisite)}
		}
		fg, err := a.featChoices(ctx, "feat_"+slug, choice.Feat, td, "", sheet, ch.CampaignID)
		if err != nil {
			a.serverError(c, err)
			return
		}
		groups = append(groups, fg...)
	}
	picks, err := readPicks(c, groups)
	if featErr != nil {
		err = featErr // ahead of the feat's own choices: it may not be taken at all
	}
	if err == nil {
		err = spellErr
	}
	if err == nil {
		err = a.updateCharacter(ctx, ch.ID, func(q *sqlc.Queries, dc *dnd5e.Character) error {
			before, err := dc.Sheet.JSON()
			if err != nil {
				return err
			}
			maxHP := dc.MaxHP
			if err := dc.LevelUp(choice); err != nil {
				return userError{err}
			}
			applyPicks(&dc.Sheet, groups, picks)
			if err := addClassSpells(&dc.Sheet, spells); err != nil {
				return err
			}
			if err := checkNewOptions(choice.Class, sheet, dc.Sheet, plan.ClassLevel, learned); err != nil {
				return err
			}
			for _, sp := range lineage {
				dc.Sheet.PutFeature(freeCast(sp, sheet.Species), nil)
			}
			return q.CreateLevelUp(ctx, sqlc.CreateLevelUpParams{CharacterID: ch.ID, MaxHp: int64(maxHP), SheetJson: before})
		})
	}
	if errors.As(err, &ue) {
		v, verr := a.levelUpView(ctx, ch, info.Slug, false)
		if verr != nil {
			a.serverError(c, verr)
			return
		}
		v.Error, v.ErrorFor, v.Form = ue.Error(), errorField(err), c.Request.PostForm
		if !v.DrawsField(v.ErrorFor) {
			v.ErrorFor = "" // a crafted post's field is not on the page: show it at the top
		}
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
