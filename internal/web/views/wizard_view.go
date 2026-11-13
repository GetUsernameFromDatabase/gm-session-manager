package views

import (
	"fmt"
	"maps"
	"net/url"
	"slices"
	"strings"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/web/utils"

	"github.com/a-h/templ"
)

// WizardStep is one step of character creation.
type WizardStep struct {
	Key, Title string
}

// Wizard step keys, in order. Each also marks the step done in Sheet.SetUp.
const (
	StepClass      = "class"
	StepBackground = "background"
	StepSpecies    = "species"
	StepAbilities  = "abilities"
	StepChoices    = "choices"
	StepEquipment  = "equipment"
	StepDetails    = "details"
	StepReview     = "review"
)

// WizardView is one page of the creation wizard.
type WizardView struct {
	CharacterID int64
	Step        string
	Steps       []WizardStep
	// Sheet is the character as it stands, read-only.
	Sheet SheetView
	// Error is shown at the top; Errors next to their fields.
	Error  string
	Errors FieldErrors
	// Form is the refused form's values, ticked again; nil on a fresh page.
	Form url.Values
	// Options are the compendium's classes, species or backgrounds;
	// Picked is the one being looked at and Choice what it gives.
	Options []Option
	Picked  string
	// Variants are the picked species' subraces; Variant the one looked at.
	Variants  []Option
	Variant   string
	Choice    *WizardChoice
	Equipment []EquipmentChoice
	Choices   *CreationChoices
	// Missing are the earlier steps this one needs before it can be saved.
	Missing []WizardStep
	// LegacyBackground: the background added its increases already (a sheet
	// started before they moved to the abilities step).
	LegacyBackground bool
}

// CreationChoices are the picks a level-1 class and a species leave open.
type CreationChoices struct {
	Groups    []ChoiceGroup
	Spells    []Option
	SpellNeed SpellNeed
	// Columns are the class table at level 1, for the spell totals.
	Columns map[string]string
}

// ChoiceGroup asks for up to Count of Options (a lineage, a skill,
// languages, invocations).
type ChoiceGroup struct {
	Key, Title string
	Count      int
	Options    []compendium.NameDesc
	// Known are had already: shown ticked and disabled, not posted.
	Known []string
	// Sub are the choices an option brings, shown while it is picked.
	Sub map[string][]ChoiceGroup
}

// WizardChoice is what a class, species or background gives.
type WizardChoice struct {
	Slug, Name string
	// Traits are a class's core traits (primary ability, hit die…).
	Traits map[string]string
	// Facts are a species' or background's tagged parts (size, speed,
	// ability scores…), shown as they are.
	Facts []compendium.NameDesc
	// Skills are the skill keys to choose SkillCount of (a class) or that
	// come with it (a background).
	SkillCount int
	Skills     []string
	// Abilities are a background's three abilities for its increases.
	Abilities []string
	Features  []FeaturePreview
	// SkillPick is the class's skill choice: on the class step, and on the
	// background step when the background grants a skill the class step
	// picked, so the player picks another class skill instead.
	SkillPick *SkillPick
}

// SkillPick asks for Count of a class's skill Options. Known are had
// already (from the background), shown ticked and disabled; Picked are
// ticked at first (the class picks a background did not overlap).
type SkillPick struct {
	Count                  int
	Options, Known, Picked []string
}

// FeaturePreview is a feature as the wizard shows it, with what the sheet
// will track for it.
type FeaturePreview struct {
	Name, Desc string
	Effects    []string
}

// EquipmentChoice is a class's or background's starting equipment options.
type EquipmentChoice struct {
	Key, Title string
	Options    []compendium.EquipmentOption
}

func (v WizardView) path(step string) string {
	return fmt.Sprintf("/characters/%d/create/%s", v.CharacterID, step)
}

// done reports whether a step that adds to the sheet was applied.
func (v WizardView) done(step string) bool { return v.Sheet.C.Sheet.SetUp(step) }

// skipped are the steps before the current one that add to the sheet but
// were not applied (details only edits; review is the end).
func (v WizardView) skipped() []WizardStep {
	var out []WizardStep
	for _, st := range v.Steps {
		if st.Key == v.Step {
			break
		}
		if st.Key != StepDetails && !v.done(st.Key) {
			out = append(out, st)
		}
	}
	return out
}

func choiceName(c *WizardChoice) string {
	if c == nil {
		return ""
	}
	return c.Name
}

func skillName(key string) string {
	for _, sk := range dnd5e.Skills {
		if sk.Key == key {
			return sk.Name
		}
	}
	return key
}

// Choose asks for n picks, as legends and errors say it: "choose one",
// "choose 2", "choose 2 more" (more: on top of ones already had).
func Choose(n int, more bool) string {
	s := "choose " + map[int]string{0: valNone, 1: "one"}[n]
	if n > 1 {
		s += fmt.Sprint(n)
	}
	if more && n > 0 {
		s += " more"
	}
	return s
}

// FieldErrors are a refused form's errors, by the field they are about.
type FieldErrors map[string]string

// Error keys shared by handlers, templates and tests.
const (
	FieldName      = "name"
	FieldSize      = "size"
	FieldVision    = "vision"
	FieldBackstory = "backstory"
	FieldMode      = "mode"
	FieldSkill     = "skill"
	FieldSpell     = "spell"
	FieldASI       = "asi"
	// FieldMastery is the Weapon Mastery choice group's key.
	FieldMastery = "weapon_mastery"
)

func (v WizardView) fieldErr() FieldErrors { return v.Errors }

// DrawsField reports whether the page shows an error at key; the handler shows
// one for a field it does not draw (a crafted post) at the top instead.
func (v WizardView) DrawsField(key string) bool {
	switch key {
	case FieldName, FieldSize, FieldVision, FieldBackstory:
		return v.Step == StepDetails
	case dnd5e.Str, dnd5e.Dex, dnd5e.Con, dnd5e.Int, dnd5e.Wis, dnd5e.Cha:
		return v.Step == StepAbilities
	case FieldMode:
		return v.Step == StepAbilities && v.Choice != nil && len(v.Choice.Abilities) == 3
	case FieldSkill:
		return v.Choice != nil && (v.Step == StepClass && v.Choice.SkillCount > 0 || v.Step == StepBackground && v.Choice.SkillPick != nil)
	case FieldSpell:
		return v.Choices != nil && len(v.Choices.Spells) > 0
	}
	return v.Choices != nil && drawsGroup(v.Choices.Groups, key)
}

// drawsGroup mirrors choiceFieldset: a group's key, then its options' own groups.
func drawsGroup(groups []ChoiceGroup, key string) bool {
	for _, g := range groups {
		if g.Key == key {
			return true
		}
		for _, o := range g.Options {
			if drawsGroup(g.Sub[o.Name], key) {
				return true
			}
		}
	}
	return false
}

// describedBy links a control group to the error at(key) shows; no attribute
// when key has no error.
func (fe FieldErrors) describedBy(key string) templ.Attributes {
	if fe[key] == "" {
		return nil
	}
	return templ.Attributes{attrAriaDescribedby: "error-" + key}
}

// radioAttrs labels a single-pick radio group, and marks it invalid when
// it holds an error (aria-invalid is valid on a radiogroup, not a fieldset).
// The fieldset links the error already (describedBy), so this does not.
func (fe FieldErrors) radioAttrs(key, label string) templ.Attributes {
	a := templ.Attributes{attrAriaLabel: label}
	if fe[key] != "" {
		a[attrAriaInvalid] = ariaTrue
	}
	return a
}

// optionsGrid lays a long list of terse options (weapon masteries, skills)
// out in columns; other lists get short, as feat texts need the width.
func optionsGrid(opts []compendium.NameDesc, short string) string {
	if len(opts) > 8 && !slices.ContainsFunc(opts, func(o compendium.NameDesc) bool { return len(o.Desc) > 40 }) {
		return "grid grid-cols-2 sm:grid-cols-3"
	}
	return short
}

// optionText is an equipment option as one line: "(A) 2 Daggers, Robe and 8 GP".
func optionText(o compendium.EquipmentOption) string {
	var items []string
	for _, it := range o.Items {
		items = append(items, strings.TrimSpace(utils.IfElse(it.Qty > 1, fmt.Sprint(it.Qty), "")+" "+it.Name))
	}
	if o.GP > 0 {
		items = append(items, fmt.Sprint(o.GP, " GP"))
	}
	if n := len(items); n > 1 {
		items = append(items[:n-2], items[n-2]+" and "+items[n-1])
	}
	return "(" + o.Letter + ") " + strings.Join(items, ", ")
}

// inputAttrs are an input's attributes, linked to its error and marked
// invalid when key has one.
func (fe FieldErrors) inputAttrs(key string, a templ.Attributes) templ.Attributes {
	if fe[key] == "" {
		return a
	}
	out := templ.Attributes{attrAriaInvalid: ariaTrue}
	maps.Copy(out, a)
	maps.Copy(out, fe.describedBy(key))
	return out
}
