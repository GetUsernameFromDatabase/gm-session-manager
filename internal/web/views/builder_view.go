package views

import (
	"encoding/json"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/dnd5e"
	selectcomp "gm-session-manager/internal/web/components/select"

	"github.com/a-h/templ"
)

// builderID is the command builder dialog on the character page.
const builderID = "command-builder"

// builderTarget is where builder forms load into.
const builderTarget = "#builder"

// builderOp is one form of the command builder: the command prefix it sends,
// the sign put before its number, and (by Key) the fields it asks for. The
// server appends the fields to the prefix as raw arguments in a fixed order:
// subject, ability, dice, dtype, prof, n, text.
type builderOp struct {
	Key, Label, Prefix, Sign string
}

// builderVerb is one group of builder forms.
type builderVerb struct {
	Key, Label string
	Ops        []builderOp
}

// builderVerbs lists what the sheet's owner can do, in menu order. Forms
// that could not apply right now (no slots, no items) are left out.
func (s SheetView) builderVerbs() []builderVerb {
	sh := s.C.Sheet
	slots := []builderOp{}
	if sh.MaxSpellLevel() > 0 {
		slots = append(slots,
			builderOp{opSlotUse, "Use slot", "slot use", ""},
			builderOp{opSlotRestore, "Restore slot", "slot restore", ""})
	}
	slots = append(slots, builderOp{opSlotMax, "Set slot count", "slot max", ""})
	if sh.Pact.Max > 0 {
		slots = append(slots,
			builderOp{opPactUse, "Use pact slot", "pact use", ""},
			builderOp{opPactRestore, "Restore pact slot", "pact restore", ""})
	}
	rest := []builderOp{
		{opRestShort, "Short rest", "rest short", ""},
		{opRestLong, "Long rest", "rest long", ""},
		{opHitdieUse, "Spend hit die", "hitdie use", ""},
		{opHitdieRestore, "Regain hit die", "hitdie restore", ""},
		{opInspiration, "Toggle inspiration", "inspiration", ""},
	}
	if dying(s) {
		rest = append(rest,
			builderOp{opDSSuccess, "Death save success", "deathsave success", ""},
			builderOp{opDSFail, "Death save failure", "deathsave fail", ""},
			builderOp{opDSReset, "Reset death saves", "deathsave reset", ""})
	}
	return []builderVerb{
		{verbHP, "Hit points", []builderOp{
			{opHPSet, "Set HP", "hp", "="},
			{opHPTemp, "Temp HP", "temp", ""},
			{opHPMax, "Max HP", "maxhp", "="},
		}},
		{verbConditions, "Conditions", []builderOp{
			{opCondAdd, labelAdd, "condition add", ""},
			{opCondRemove, labelRemove, "condition remove", ""},
			{opExhUp, "More exhaustion", "exhaustion", "+"},
			{opExhDown, "Less exhaustion", "exhaustion", "-"},
		}},
		{verbRest, "Rest & recovery", rest},
		{verbAbilities, "Ability scores", []builderOp{{opAbility, "Set score", "set ability", ""}}},
		{verbProfs, "Proficiencies", []builderOp{
			{focusSkill, "Skill", "prof set skill", ""},
			{focusSave, "Saving throw", "prof set save", ""},
			{opSkillAdd, "Add skill", "prof add skill", ""},
			{opSkillRemove, "Remove skill", "prof remove skill", ""},
			{opExpAdd, "Add expertise", "prof add expertise", ""},
			{opExpRemove, "Remove expertise", "prof remove expertise", ""},
			{opSaveAdd, "Add save", "prof add save", ""},
			{opSaveRemove, "Remove save", "prof remove save", ""},
			{opOtherAdd, "Add other", "prof add other", ""},
			{opOtherRemove, "Remove other", "prof remove other", ""},
			{opLangAdd, "Add language", "prof add language", ""},
			{opResAdd, "Add resistance", "prof add resistance", ""},
			{opImmAdd, "Add immunity", "prof add immunity", ""},
			{opVulnAdd, "Add vulnerability", "prof add vulnerability", ""},
		}},
		{verbDetails, "Details", []builderOp{
			{opSetName, labelName, "set name", ""},
			{opSetSpecies, "Species", "set species", ""},
			{opSetBackground, "Background", "set background", ""},
			{opSetBackstory, "Backstory", "set backstory", ""},
			{opSetAC, "AC bonus", "set ac", ""},
			{opSetSpeed, "Base speed", "set speed", ""},
			{opSetSize, "Size", "set size", ""},
			{opSetVision, "Vision", "set vision", ""},
		}},
		{verbArmor, "Armor", []builderOp{{opArmor, "Armor", "", ""}}},
		{verbWeapons, "Weapons", []builderOp{
			{opWeaponAdd, labelAddOrReplace, "weapon add", ""},
			{opWeaponRemove, labelRemove, "weapon remove", ""},
			{OpMastery, "Weapon Mastery", "", ""},
		}},
		{verbMods, "Custom modifiers", []builderOp{
			{opModAdd, labelAddOrReplace, "mod add", ""},
			{modEdit, "Save", "mod edit", ""},
		}},
		{verbSlots, "Spell slots", slots},
		{verbSpells, "Spells", []builderOp{
			{opSpellAdd, labelAdd, "", ""},
			{opSpellPrepare, "Prepare", "spell prepare", ""},
			{opSpellUnprepare, "Unprepare", "spell unprepare", ""},
			{opSpellRemove, labelRemove, "spell remove", ""},
		}},
		{verbItems, "Inventory", []builderOp{
			{opItemAdd, labelAdd, "item add", ""},
			{opItemEdit, "Change count", "item set", ""},
			{opItemNote, "Note", "item note", ""},
		}},
		{verbCoins, "Coins", []builderOp{{opCoinEdit, "Coins", "", ""}}},
		{verbFeatures, "Features", []builderOp{
			{opFeatureAdd, labelAddOrReplace, "feature add", ""},
			{opFeatureRemove, labelRemove, "feature remove", ""},
		}},
	}
}

// Builder verb keys: the groups of forms in the menu.
const (
	verbHP         = "hp"
	verbConditions = "conditions"
	verbRest       = "rest"
	verbAbilities  = "abilities"
	verbProfs      = "profs"
	verbDetails    = "details"
	verbArmor      = "armor"
	verbWeapons    = "weapons"
	verbMods       = "mods"
	verbSlots      = "slots"
	verbSpells     = "spells"
	verbItems      = "items"
	verbCoins      = "coins"
	verbFeatures   = "features"
)

// Builder op keys: the forms inside a verb (see also the focus forms and
// OpMastery below).
const (
	opSlotUse        = "slot-use"
	opSlotRestore    = "slot-restore"
	opSlotMax        = "slot-max"
	opPactUse        = "pact-use"
	opPactRestore    = "pact-restore"
	opRestShort      = "rest-short"
	opRestLong       = "rest-long"
	opHitdieUse      = "hitdie-use"
	opHitdieRestore  = "hitdie-restore"
	opInspiration    = "inspiration"
	opDSSuccess      = "ds-success"
	opDSFail         = "ds-fail"
	opDSReset        = "ds-reset"
	opHPSet          = "hp-set"
	opHPTemp         = "hp-temp"
	opHPMax          = "hp-max"
	opCondAdd        = "cond-add"
	opCondRemove     = "cond-remove"
	opExhUp          = "exh-up"
	opExhDown        = "exh-down"
	opAbility        = "ability"
	opSkillAdd       = "skill-add"
	opSkillRemove    = "skill-remove"
	opExpAdd         = "exp-add"
	opExpRemove      = "exp-remove"
	opSaveAdd        = "save-add"
	opSaveRemove     = "save-remove"
	opOtherAdd       = "other-add"
	opOtherRemove    = "other-remove"
	opLangAdd        = "lang-add"
	opResAdd         = "res-add"
	opImmAdd         = "imm-add"
	opVulnAdd        = "vuln-add"
	opSetName        = "set-name"
	opSetSpecies     = "set-species"
	opSetBackground  = "set-background"
	opSetBackstory   = "set-backstory"
	opSetAC          = "set-ac"
	opSetSpeed       = "set-speed"
	opSetSize        = "set-size"
	opSetVision      = "set-vision"
	opArmor          = "armor"
	opWeaponAdd      = "weapon-add"
	opWeaponRemove   = "weapon-remove"
	opModAdd         = "mod-add"
	opSpellAdd       = "spell-add"
	opSpellPrepare   = "spell-prepare"
	opSpellUnprepare = "spell-unprepare"
	opSpellRemove    = "spell-remove"
	opItemAdd        = "item-add"
	opItemEdit       = "item-edit"
	opItemNote       = "item-note"
	opCoinEdit       = "coin-edit"
	opFeatureAdd     = "feature-add"
	opFeatureRemove  = "feature-remove"
)

// Form labels that several builder forms share.
const (
	labelAdd          = "Add"
	labelRemove       = "Remove"
	labelAddOrReplace = "Add or replace"
)

// pickBuilder finds the chosen verb and op, defaulting to the first of each.
func (s SheetView) pickBuilder(verb, op string) ([]builderVerb, builderVerb, builderOp) {
	verbs := s.builderVerbs()
	v := verbs[max(slices.IndexFunc(verbs, func(b builderVerb) bool { return b.Key == verb }), 0)]
	return verbs, v, v.Ops[max(slices.IndexFunc(v.Ops, func(b builderOp) bool { return b.Key == op }), 0)]
}

// builderURL opens a builder form, with subject preselected when not empty.
func builderURL(s SheetView, verb, op, subject string) string {
	return fmt.Sprintf("/characters/%d/builder?verb=%s&op=%s&subject=%s", s.ID, verb, op, url.QueryEscape(subject))
}

// builderLoad makes an element load a builder form into the dialog.
func builderLoad(s SheetView, verb, op, subject string) templ.Attributes {
	return templ.Attributes{attrHxGet: builderURL(s, verb, op, subject), attrHxTarget: builderTarget, attrHxSwap: swapInner}
}

// BuilderExtra is what a builder form takes from the compendium: a weapon
// or feature picked in its look-up, and names a field suggests.
type BuilderExtra struct {
	Weapon  *dnd5e.Weapon
	Feature *dnd5e.Feature
	// FeatureMods are Feature's stat changes, read from its text.
	FeatureMods []dnd5e.CustomMod
	// FeatSource is the source key of the feat Feature was picked from;
	// saving it adds the feat to the sheet's feats, pinned.
	FeatSource  string
	Suggestions []string
	// Masteries are the compendium weapons with a mastery property (as
	// Desc), for OpMastery.
	Masteries []compendium.NameDesc
}

// OpMastery is the builder form that chooses the weapon kinds mastered.
const OpMastery = "mastery"

// BuilderPick is one result of the builder's compendium look-up. Feature
// names a class feature inside the class entry EntryID.
type BuilderPick struct {
	EntryID                     int64
	Slug, Feature, Name, Detail string
}

// Focus forms are reached by tapping one thing on the sheet, so they have no
// chip. modEdit's subject is the modifier's index on the sheet; the others'
// is a skill or ability key.
const (
	modEdit    = "mod-edit"
	focusSkill = "skill"
	focusSave  = "save"
)

// chips are the verb's forms that get a chip; proficiencies show only the
// forms of cur's tab.
func (v builderVerb) chips(cur builderOp) []builderOp {
	return slices.DeleteFunc(slices.Clone(v.Ops), func(o builderOp) bool {
		return slices.Contains([]string{modEdit, focusSkill, focusSave}, o.Key) || (v.Key == verbProfs && opGroup(o.Key) != opGroup(cur.Key))
	})
}

// opGroup is what a form's key changes: "skill" for skill, skill-add and
// skill-remove.
func opGroup(key string) string {
	g, _, _ := strings.Cut(key, "-")
	return g
}

// profTabGroups are the proficiency tabs by opGroup; each opens its add form.
var profTabGroups = []struct{ Key, Label string }{
	{"skill", "Skills"}, {"exp", "Expertise"}, {"save", "Saves"}, {"lang", "Languages"}, {"res", "Resistances"},
	{"imm", "Immunities"}, {"vuln", "Vulnerabilities"}, {"other", labelOther},
}

// damageList is what a damage defense op edits: the command's noun, the
// add field's label and the sheet's list.
func (s SheetView) damageList(key string) (noun, label string, list []string) {
	switch key {
	case opImmAdd:
		return "immunity", "Add an immunity", s.C.Sheet.Immunities
	case opVulnAdd:
		return "vulnerability", "Add a vulnerability", s.C.Sheet.Vulnerabilities
	}
	return "resistance", "Add a resistance", s.C.Sheet.Resistances
}

// pressedClass styles a toggle button by its aria-pressed state.
const pressedClass = "aria-pressed:border-primary aria-pressed:bg-primary aria-pressed:text-primary-foreground aria-pressed:hover:bg-primary/80"

// profFocus is one skill or save as its focus form shows it.
type profFocus struct {
	Name, Level, ModTarget string
	Bonus                  int
	Levels                 []string
}

func (s SheetView) profFocus(key, subject string) (profFocus, bool) {
	sh := s.C.Sheet
	if key == focusSave {
		name, ok := dnd5e.AbilityNames[subject]
		level := dnd5e.ProfNone
		if slices.Contains(sh.SaveProfs, subject) {
			level = dnd5e.ProfProficient
		}
		return profFocus{name + " save", level, "save." + subject, sh.SaveBonus(subject), []string{dnd5e.ProfNone, dnd5e.ProfProficient}}, ok
	}
	i := slices.IndexFunc(dnd5e.Skills, func(sk dnd5e.Skill) bool { return sk.Key == subject })
	if i < 0 {
		return profFocus{}, false
	}
	sk := dnd5e.Skills[i]
	return profFocus{sk.Name, sh.ProfLevel(sk.Key), "skill." + sk.Key, sh.SkillBonus(sk), []string{dnd5e.ProfNone, dnd5e.ProfProficient, dnd5e.ProfExpertise}}, true
}

var profLabels = map[string]string{dnd5e.ProfNone: "Not proficient", dnd5e.ProfProficient: "Proficient", dnd5e.ProfExpertise: "Expertise"}

// removeVals are the fields of the button that removes what an edit form
// shows, or nil when the subject does not exist yet.
func (s SheetView) removeVals(key, subject string) map[string]string {
	switch {
	case key == opWeaponAdd && s.weapon(subject) != nil:
		return map[string]string{fieldPrefix: "weapon remove", fieldSubject: subject}
	case key == opFeatureAdd && s.feature(subject) != nil:
		return map[string]string{fieldPrefix: "feature remove", fieldSubject: subject}
	}
	return nil
}

func (s SheetView) weapon(name string) *dnd5e.Weapon {
	i := slices.IndexFunc(s.C.Sheet.Weapons, func(w dnd5e.Weapon) bool { return w.Name == name })
	if i < 0 {
		return nil
	}
	return &s.C.Sheet.Weapons[i]
}

func (s SheetView) feature(name string) *dnd5e.Feature {
	i := slices.IndexFunc(s.C.Sheet.Features, func(f dnd5e.Feature) bool { return f.Name == name })
	if i < 0 {
		return nil
	}
	return &s.C.Sheet.Features[i]
}

// modAt is the modifier at index subject, if it is still there.
func (s SheetView) modAt(subject string) (dnd5e.CustomMod, bool) {
	i, err := strconv.Atoi(subject)
	if err != nil || i < 0 || i >= len(s.C.Sheet.Modifiers) {
		return dnd5e.CustomMod{}, false
	}
	return s.C.Sheet.Modifiers[i], true
}

// stayOpen are forms (op key to verb) used several times in a row: after a
// submit they reload with the sheet's new choices instead of closing the dialog.
var stayOpen = map[string]string{opFeatureRemove: verbFeatures, opLangAdd: verbProfs, opResAdd: verbProfs, opImmAdd: verbProfs, opVulnAdd: verbProfs}

// The builder form fields (and hx-vals keys) holding the command prefix and
// what it applies to.
const (
	fieldPrefix  = "prefix"
	fieldSubject = "subject"
)

// builderDoneJS closes the dialog (or reloads a stayOpen form) after the
// form's own submit, not after requests from fields inside it.
const builderDoneJS = `if (event.detail.elt !== this || !event.detail.successful) return; if (this.dataset.reload) htmx.ajax("GET", this.dataset.reload, {target: "#builder", swap: "innerHTML"}); else this.querySelector("[data-slot=dialog-close]").click()`

// submitLabel names the form's button: Save for forms that set a value,
// whose labels name the field rather than an action.
func (o builderOp) submitLabel() string {
	if strings.HasPrefix(o.Prefix, "set ") {
		return "Save"
	}
	return o.Label
}

// builderChangeJS sets the number field's limit (0 is none, as in
// numberAttrs) and current value for the chosen subject, from the form's
// data-max and data-value maps.
const builderChangeJS = `if (event.target.name !== "subject" || !this.elements.n) return; const n = this.elements.n, k = event.target.value, m = JSON.parse(this.dataset.max), v = JSON.parse(this.dataset.value); if (m[k] > 0) n.max = m[k]; else n.removeAttribute("max"); if (k in v) n.value = v[k]`

func toJSON(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

// featureToEdit is the feature the form starts from: the one on the sheet,
// one picked in the look-up, or a new one; with its stat changes. A feature
// with nothing automated yet gets suggestions read from its text.
func (s SheetView) featureToEdit(subject string, x BuilderExtra) (dnd5e.Feature, []dnd5e.CustomMod, bool) {
	if f := s.feature(subject); f != nil {
		var mods []dnd5e.CustomMod
		for _, m := range s.C.Sheet.Modifiers {
			if m.Source == f.Name {
				mods = append(mods, m)
			}
		}
		if f.Action == "" && !f.Uses.Limited() && f.Unarmored == "" && len(mods) == 0 && f.Description != "" {
			fe := *f
			return fe, dnd5e.ParseEffects(f.Name, f.Description, nil, "").Apply(&fe), true
		}
		return *f, mods, false
	}
	if x.Feature != nil {
		return *x.Feature, x.FeatureMods, true
	}
	return dnd5e.Feature{}, nil, false
}

// chooseFirst moves the item with value subject to the front, where the
// select and the number field take their starting choice from.
func chooseFirst(items []selectcomp.ItemData, subject string) []selectcomp.ItemData {
	i := slices.IndexFunc(items, func(it selectcomp.ItemData) bool { return it.Value == subject })
	if i <= 0 {
		return items
	}
	return slices.Concat(items[i:i+1], items[:i], items[i+1:])
}

func firstValue(items []selectcomp.ItemData) string {
	if len(items) == 0 {
		return ""
	}
	return items[0].Value
}

func subjectLabel(key string) string {
	switch {
	case strings.HasPrefix(key, "cond"):
		return "Condition"
	case strings.HasPrefix(key, "slot"):
		return "Slot level"
	case strings.HasPrefix(key, "item"):
		return "Item"
	case strings.HasPrefix(key, "save"):
		return "Saving throw"
	case strings.HasPrefix(key, "spell"):
		return "Spell"
	case strings.HasPrefix(key, "feature"):
		return "Feature"
	case strings.HasPrefix(key, "weapon"):
		return "Weapon"
	case strings.HasPrefix(key, "other"):
		return "Proficiency"
	}
	return "Skill"
}

func numberAttrs(lo, hi int, focus bool) templ.Attributes {
	a := templ.Attributes{attrMin: fmt.Sprint(lo), attrInputmode: inputmodeNumeric}
	if hi > 0 {
		a["max"] = fmt.Sprint(hi)
	}
	if focus {
		a["autofocus"] = true
	}
	return a
}

func weaponAbilities() []selectcomp.ItemData {
	items := []selectcomp.ItemData{{Value: dnd5e.Finesse, Label: "Finesse (better of STR and DEX)"}}
	for _, a := range dnd5e.Abilities {
		items = append(items, selectcomp.ItemData{Value: a, Label: dnd5e.AbilityNames[a]})
	}
	return items
}

var actionItems = []selectcomp.ItemData{{Value: " ", Label: "Nothing special"}, {Value: dnd5e.ActionMain, Label: "An action"}, {Value: dnd5e.ActionBonus, Label: "A bonus action"}, {Value: dnd5e.ActionReaction, Label: "A reaction"}, {Value: dnd5e.ActionMagic, Label: "A Magic action"}}

var sizeItems = selectItems(dnd5e.Sizes, func(s string) string { return s }, func(s string) string { return s })

var resetItems = []selectcomp.ItemData{{Value: dnd5e.ResetLong, Label: "A long rest"}, {Value: dnd5e.ResetShort, Label: "A short or long rest"}, {Value: dnd5e.ResetShortOne, Label: "One on a short rest, all on a long rest"}}

var whenItems = []selectcomp.ItemData{{Value: " ", Label: "Always"}, {Value: dnd5e.WhenNoHeavyArmor, Label: "Not in heavy armor"}, {Value: dnd5e.WhenUnarmored, Label: "Wearing no armor"}, {Value: dnd5e.WhenUnarmoredNoShield, Label: "No armor and no shield"}}

var unarmoredItems = append([]selectcomp.ItemData{{Value: " ", Label: "Armor or base AC as usual"}},
	selectItems(dnd5e.Abilities, func(a string) string { return a }, func(a string) string { return "10 + Dexterity + " + dnd5e.AbilityNames[a] })...)

// usesItems are the choices of how often a feature can be used; the class
// table choice shows only for a feature that has one.
func usesItems(u dnd5e.Uses) []selectcomp.ItemData {
	items := []selectcomp.ItemData{{Value: usesUnlimited, Label: "Unlimited"}, {Value: usesNumber, Label: "A number"}, {Value: usesProf, Label: "Proficiency bonus"}}
	for _, a := range dnd5e.Abilities {
		items = append(items, selectcomp.ItemData{Value: a, Label: dnd5e.AbilityNames[a] + " modifier"})
	}
	if len(u.ByLevel) > 0 {
		items = append(items, selectcomp.ItemData{Value: usesTable, Label: "By " + u.Class + " level (class table)"})
	}
	return items
}

// How often a feature can be used: the values of the uses select, which
// Ability keys join.
const (
	usesUnlimited = " "
	usesNumber    = "n"
	usesProf      = "prof"
	usesTable     = "table"
)

func usesValue(u dnd5e.Uses) string {
	switch {
	case u.Fixed > 0:
		return usesNumber
	case u.Ability != "":
		return u.Ability
	case u.Prof:
		return usesProf
	case len(u.ByLevel) > 0:
		return usesTable
	}
	return usesUnlimited
}

// modTargetItems are the modifier targets, for a select.
func (s SheetView) modTargetItems() []selectcomp.ItemData {
	var items []selectcomp.ItemData
	for _, t := range s.modTargets() {
		items = append(items, selectcomp.ItemData{Value: t, Label: modTargetLabel(t)})
	}
	return items
}

// pickURL reloads the builder form filled in from a look-up pick.
func pickURL(characterID int64, verb, op string, p BuilderPick) string {
	return fmt.Sprintf("/characters/%d/builder?verb=%s&op=%s&entry=%d&feature=%s", characterID, verb, op, p.EntryID, url.QueryEscape(p.Feature))
}

// closeAfter makes a builder button close the dialog once its request
// succeeds, since the sheet behind it shows the result.
func closeAfter(a templ.Attributes) templ.Attributes {
	// templ blanks a spread event handler unless it is a ComponentScript.
	a["hx-on::after-request"] = templ.JSUnsafeFuncCall("if (event.detail.successful) document.querySelector('#" + builderID + " [data-slot=dialog-close]')?.click()")
	return a
}

// masteryCount is "2/3 chosen": weapon kinds mastered of those allowed.
func masteryCount(s SheetView) string {
	return fmt.Sprintf("%d/%d chosen", len(s.C.Sheet.Masteries), s.MasteryLimit)
}

// reloadAfter makes a builder button load the form at url again after a
// successful request, instead of closing the dialog.
func reloadAfter(a templ.Attributes, url string) templ.Attributes {
	a["hx-on::after-request"] = templ.JSUnsafeFuncCall("if (event.detail.successful) htmx.ajax('GET', '" + url + "', {target: '#builder', swap: 'innerHTML'})")
	return a
}

// detail is the current value a details form edits.
func (s SheetView) detail(key string) string {
	switch key {
	case opSetName:
		return s.C.Name
	case opSetSpecies:
		return s.C.Sheet.Species
	}
	return s.C.Sheet.Background
}

// builderSubjects lists what a form can pick (nil when it picks nothing,
// empty when there is nothing to pick right now), with the most its number
// may be and the number to start from for each pick.
func (s SheetView) builderSubjects(key string) ([]selectcomp.ItemData, map[string]int, map[string]int) {
	sh := s.C.Sheet
	items := []selectcomp.ItemData{}
	maxes, values := map[string]int{}, map[string]int{}
	add := func(value, label string) { items = append(items, selectcomp.ItemData{Value: value, Label: label}) }
	skills := func(keep func(sk dnd5e.Skill) bool) {
		for _, sk := range dnd5e.Skills {
			if keep(sk) {
				add(sk.Key, sk.Name)
			}
		}
	}
	abilities := func(keep func(a string) bool) {
		for _, a := range dnd5e.Abilities {
			if keep(a) {
				add(a, dnd5e.AbilityNames[a])
			}
		}
	}
	has := func(list []string) func(string) bool { return func(x string) bool { return slices.Contains(list, x) } }
	switch key {
	case opCondAdd:
		for _, slug := range append(slices.Clone(dnd5e.Conditions), dnd5e.TrackConcentration) {
			if _, on := s.condition(slug); !on && slug != dnd5e.Exhaustion {
				add(slug, slug)
			}
		}
		add(dnd5e.TrackCustom, "custom effect")
	case opCondRemove:
		for _, c := range s.Conditions {
			if c.Slug != dnd5e.Exhaustion && !slices.ContainsFunc(items, func(it selectcomp.ItemData) bool { return it.Value == c.Slug }) {
				add(c.Slug, strings.TrimSpace(c.Slug+" "+c.Note))
			}
		}
	case opExhUp, opExhDown:
		level := 0
		if c, on := s.condition(dnd5e.Exhaustion); on {
			level, _ = strconv.Atoi(c.Note)
		}
		maxes[""] = 6 - level
		if key == opExhUp {
			return nil, maxes, values
		}
		maxes[""] = level
		if level > 0 {
			add(dnd5e.Exhaustion, dnd5e.Exhaustion)
		}
	case opAbility:
		abilities(func(string) bool { return true })
		for _, a := range dnd5e.Abilities {
			values[a] = sh.Scores[a]
		}
	case opSkillAdd:
		skills(func(sk dnd5e.Skill) bool { return !has(sh.SkillProfs)(sk.Key) })
	case opSkillRemove:
		skills(func(sk dnd5e.Skill) bool { return has(sh.SkillProfs)(sk.Key) })
	case opExpAdd:
		skills(func(sk dnd5e.Skill) bool { return has(sh.SkillProfs)(sk.Key) && !has(sh.Expertise)(sk.Key) })
	case opExpRemove:
		skills(func(sk dnd5e.Skill) bool { return has(sh.Expertise)(sk.Key) })
	case opSaveAdd:
		abilities(func(a string) bool { return !has(sh.SaveProfs)(a) })
	case opSaveRemove:
		abilities(has(sh.SaveProfs))
	case opOtherRemove:
		for _, p := range sh.OtherProfs {
			add(p, p)
		}
	case opSlotUse, opSlotRestore, opSlotMax:
		for i, sl := range sh.Slots {
			lvl := fmt.Sprint(i + 1)
			if key == opSlotMax || key == opSlotUse && sl.Used < sl.Max || key == opSlotRestore && sl.Used > 0 {
				add(lvl, fmt.Sprintf("Level %s (%s)", lvl, slotLabel(sl)))
			}
			values[lvl] = sl.Max
		}
	case opPactUse, opPactRestore:
		if key == opPactUse && sh.Pact.Used < sh.Pact.Max || key == opPactRestore && sh.Pact.Used > 0 {
			add("", "Pact slot ("+slotLabel(sh.Pact)+")")
		}
	case opItemEdit, opItemNote:
		for _, it := range sh.Inventory {
			add(it.Name, it.Name)
			values[it.Name] = it.Qty
		}
	case opSpellPrepare, opSpellUnprepare, opSpellRemove:
		for _, sp := range sh.Spells {
			if key == opSpellRemove || sp.Level > 0 && sp.Prepared == (key == opSpellUnprepare) {
				add(sp.Slug, sp.Name)
			}
		}
	case opFeatureRemove:
		for _, f := range sh.Features {
			add(f.Name, f.Name)
		}
	case opWeaponRemove:
		for _, w := range sh.Weapons {
			add(w.Name, w.Name)
		}
	case opModAdd:
		for _, t := range s.modTargets() {
			add(t, modTargetLabel(t))
		}
	default:
		return nil, maxes, values
	}
	return items, maxes, values
}

type coin struct {
	n    int
	name string
}

func coinList(c dnd5e.Coins) []coin {
	return []coin{{c.PP, "pp"}, {c.GP, "gp"}, {c.EP, "ep"}, {c.SP, "sp"}, {c.CP, "cp"}}
}

// modTargets lists every value a custom modifier can change.
func (s SheetView) modTargets() []string {
	out := []string{dnd5e.ModAC, dnd5e.ModSpeed, dnd5e.ModInit, dnd5e.ModPassive, dnd5e.ModAttack, dnd5e.ModDamage}
	for _, w := range s.C.Sheet.Weapons {
		out = append(out, dnd5e.ModAttack+"."+w.Name, dnd5e.ModDamage+"."+w.Name)
	}
	// Stowed weapons' rows link here too.
	for _, it := range s.C.Sheet.Inventory {
		if it.Weapon != nil && !slices.Contains(out, dnd5e.ModAttack+"."+it.Weapon.Name) {
			out = append(out, dnd5e.ModAttack+"."+it.Weapon.Name, dnd5e.ModDamage+"."+it.Weapon.Name)
		}
	}
	out = append(out, dnd5e.Abilities...)
	for _, a := range dnd5e.Abilities {
		out = append(out, "save."+a)
	}
	for _, sk := range dnd5e.Skills {
		out = append(out, "skill."+sk.Key)
	}
	// A modifier on a renamed or merged weapon keeps its target selectable.
	for _, m := range s.C.Sheet.Modifiers {
		if m.Target != "" && !slices.Contains(out, m.Target) {
			out = append(out, m.Target)
		}
	}
	return out
}

// modTargetLabel names a modifier target for people.
func modTargetLabel(t string) string {
	if name, ok := dnd5e.AbilityNames[t]; ok {
		return name + " score"
	}
	kind, rest, _ := strings.Cut(t, ".")
	switch kind {
	case "save":
		return dnd5e.AbilityNames[rest] + " save"
	case "skill":
		for _, sk := range dnd5e.Skills {
			if sk.Key == rest {
				return sk.Name
			}
		}
	case dnd5e.ModAttack:
		if rest != "" {
			return rest + " to hit"
		}
		return "All attacks to hit"
	case dnd5e.ModDamage:
		if rest != "" {
			return rest + " damage"
		}
		return "All weapon damage"
	}
	return map[string]string{dnd5e.ModAC: "Armor class", dnd5e.ModSpeed: statSpeed, dnd5e.ModInit: "Initiative", dnd5e.ModPassive: "Passive perception"}[t]
}
