package dnd5e

import (
	"regexp"
	"slices"
	"strings"
)

// Armor training kinds: the worn armor types (WornArmor.Type) and shields.
const (
	ArmorLight  = "light"
	ArmorMedium = "medium"
	ArmorHeavy  = "heavy"
	ArmorShield = "shield"
)

// ArmorKinds are the armor training kinds, in the order the sheet lists them.
var ArmorKinds = []string{ArmorLight, ArmorMedium, ArmorHeavy, ArmorShield}

// ArmorKindLabels name the kinds as the sheet shows them.
var ArmorKindLabels = map[string]string{ArmorLight: "Light", ArmorMedium: "Medium", ArmorHeavy: "Heavy", ArmorShield: "Shields"}

var (
	// reArmorRun is a run of armor types that ends in "armor": "Light, Medium, and Heavy armor".
	reArmorRun  = regexp.MustCompile(`(?i)((?:\b(?:light|medium|heavy)\b(?:\s*,\s*(?:and\s+|or\s+)?|\s+(?:and|or)\s+|\s+))+)armor\b`)
	reArmorType = regexp.MustCompile(`(?i)light|medium|heavy`)
	reAllArmor  = regexp.MustCompile(`(?i)\ball armor\b`)
	reShields   = regexp.MustCompile(`(?i)\bshields?\b`)
	// reArmorGrant is what a feature's text trains: "training with Light armor and Shields".
	reArmorGrant = regexp.MustCompile(`(?i)\b(?:training|proficiency) with ([^.;:]+)`)
	// reOptions marks a feature offering options (Divine Order), whose grants depend on the pick.
	reOptions = regexp.MustCompile(`(?i)\bone of the following\b`)
)

// ParseArmorTraining reads the armor kinds a proficiency text names, in
// ArmorKinds order: "Light, Medium, and Heavy armor and Shields", "All
// armor, shields", "Light armor". "None" and weapon lists ("light
// crossbows") name none.
func ParseArmorTraining(text string) []string {
	text = reTag.ReplaceAllString(text, "$1$2")
	has := map[string]bool{}
	if reAllArmor.MatchString(text) {
		has[ArmorLight], has[ArmorMedium], has[ArmorHeavy] = true, true, true
	}
	for _, m := range reArmorRun.FindAllStringSubmatch(text, -1) {
		for _, t := range reArmorType.FindAllString(m[1], -1) {
			has[strings.ToLower(t)] = true
		}
	}
	has[ArmorShield] = has[ArmorShield] || reShields.MatchString(text)
	return armorKinds(func(k string) bool { return has[k] })
}

// ArmorTrainingText words kinds as a class's Armor Training trait does,
// for ParseArmorTraining to read back: "Light, Medium, and Heavy armor and
// Shields", "Light armor", "Shields"; "" for none.
func ArmorTrainingText(kinds []string) string {
	var types []string
	for _, k := range armorKinds(func(k string) bool { return k != ArmorShield && slices.Contains(kinds, k) }) {
		types = append(types, ArmorKindLabels[k])
	}
	text := ""
	switch n := len(types); n {
	case 0:
	case 1, 2:
		text = strings.Join(types, " and ") + " armor"
	default:
		text = strings.Join(types[:n-1], ", ") + ", and " + types[n-1] + " armor"
	}
	if !slices.Contains(kinds, ArmorShield) {
		return text
	}
	if text == "" {
		return ArmorKindLabels[ArmorShield]
	}
	return text + " and " + ArmorKindLabels[ArmorShield]
}

// armorKinds are the ArmorKinds has holds, in order; nil for none.
func armorKinds(has func(string) bool) []string {
	var out []string
	for _, k := range ArmorKinds {
		if has(k) {
			out = append(out, k)
		}
	}
	return out
}

// armorGrants are the armor kinds a feature's or feat's text trains ("You
// gain training with Medium armor and Shields"). A feature offering options
// grants none: the pick decides (Divine Order's Protector).
// shortcut: the picked option's training is added by hand; read it from the pick when options are applied.
func armorGrants(desc string) []string {
	if reOptions.MatchString(desc) {
		return nil
	}
	var all []string
	for _, m := range reArmorGrant.FindAllStringSubmatch(reTag.ReplaceAllString(desc, "$1$2"), -1) {
		all = append(all, ParseArmorTraining(m[1])...)
	}
	return armorKinds(func(k string) bool { return slices.Contains(all, k) })
}

// armorSourceKey is the ProfSources key of an armor training kind, apart
// from skill keys and other proficiencies.
func armorSourceKey(kind string) string { return "armor:" + kind }

// ArmorKindNamed is the kind its key or label names, in any case ("shield",
// "Shields"); "" for none.
func ArmorKindNamed(name string) string {
	name = strings.TrimSpace(name)
	for _, k := range ArmorKinds {
		if strings.EqualFold(name, k) || strings.EqualFold(name, ArmorKindLabels[k]) {
			return k
		}
	}
	return ""
}

// TrainedIn reports whether the character has training with an armor kind.
func (s Sheet) TrainedIn(kind string) bool { return slices.Contains(s.ArmorTraining, kind) }

// AddArmorTraining gives the kinds, recording origin for those the sheet
// did not have yet, as AddProfs does.
func (s *Sheet) AddArmorTraining(origin string, kinds ...string) {
	for _, k := range kinds {
		if s.TrainedIn(k) {
			continue
		}
		s.ArmorTraining = append(s.ArmorTraining, k)
		if origin != "" {
			if s.ProfSources == nil {
				s.ProfSources = map[string]string{}
			}
			s.ProfSources[armorSourceKey(k)] = origin
		}
	}
	s.sortArmorTraining()
}

func (s *Sheet) sortArmorTraining() { s.ArmorTraining = armorKinds(s.TrainedIn) }

// ArmorTrainingLabels are the kinds the character is trained in, each with
// what granted it when known: "Heavy (Fighter)".
func (s Sheet) ArmorTrainingLabels() []string {
	out := make([]string, len(s.ArmorTraining))
	for i, k := range s.ArmorTraining {
		out[i] = s.ProfLabel(ArmorKindLabels[k], armorSourceKey(k))
	}
	return out
}

// CatchUpArmorTraining is for sheets made before ArmorTraining: it moves
// the training kept as text (moveArmorTraining) and gives each class after
// the first its multiclass training. It reports whether anything changed.
func (s *Sheet) CatchUpArmorTraining() bool {
	n := len(s.ArmorTraining)
	moved := s.moveArmorTraining()
	for _, cl := range s.Classes[min(1, len(s.Classes)):] {
		s.AddArmorTraining(cl.Name, multiclassArmor(cl.Name)...)
	}
	return moved || len(s.ArmorTraining) > n
}

// moveArmorTraining moves the armor entries of OtherProfs (class training
// kept as text, "None" included) into ArmorTraining, with their origins. An
// entry naming other things too ("Light armor, light crossbows") stays and
// gives its armor kinds as well. It reports whether anything changed.
func (s *Sheet) moveArmorTraining() bool {
	moved := false
	s.OtherProfs = slices.DeleteFunc(s.OtherProfs, func(p string) bool {
		kinds := ParseArmorTraining(p)
		none := strings.EqualFold(strings.TrimSpace(p), "none")
		if len(kinds) == 0 && !none {
			return false
		}
		moved = true
		s.AddArmorTraining(s.ProfSources[p], kinds...)
		only := none || armorOnly(p)
		if only {
			delete(s.ProfSources, p)
		}
		return only
	})
	return moved
}

// reNotArmor is what an armor training text may hold besides armor kinds.
var reNotArmor = regexp.MustCompile(`(?i)\b(?:light|medium|heavy|all|armor|shields?|and|or|training|none)\b|[\s,.]`)

// armorOnly reports whether a proficiency text names armor training and
// nothing else.
func armorOnly(p string) bool { return reNotArmor.ReplaceAllString(p, "") == "" }
