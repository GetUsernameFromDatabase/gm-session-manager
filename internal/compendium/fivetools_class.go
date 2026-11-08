package compendium

import (
	"cmp"
	"encoding/json"
	"maps"
	"slices"
	"strconv"
	"strings"
)

// This file reads the class-<name>.json files of a 5etools book folder:
// classes, subclasses and their features.

type ftClass struct {
	ftNamed
	HD struct {
		Faces int `json:"faces"`
	} `json:"hd"`
	Proficiency       []string          `json:"proficiency"`
	CasterProgression string            `json:"casterProgression"`
	ClassFeatures     []json.RawMessage `json:"classFeatures"`
	TableGroups       []ftTableGroup    `json:"classTableGroups"`
}

type ftSubclass struct {
	ftNamed
	ShortName         string         `json:"shortName"`
	ClassName         string         `json:"className"`
	CasterProgression string         `json:"casterProgression"`
	SubclassFeatures  []string       `json:"subclassFeatures"`
	TableGroups       []ftTableGroup `json:"subclassTableGroups"`
}

type ftTableGroup struct {
	ColLabels            []string            `json:"colLabels"`
	Rows                 [][]json.RawMessage `json:"rows"`
	RowsSpellProgression [][]json.RawMessage `json:"rowsSpellProgression"`
}

type ftClassFeature struct {
	ftNamed
	ClassName         string `json:"className"`
	ClassSource       string `json:"classSource"`
	SubclassShortName string `json:"subclassShortName"`
	SubclassSource    string `json:"subclassSource"`
	Level             int    `json:"level"`
}

// ftCasters maps casterProgression to our caster types. Only subclasses
// (Eldritch Knight, Arcane Trickster) are third casters.
var ftCasters = map[string]string{"full": "full", "1/2": "half", "1/3": "third", "artificer": "half", "pact": "pact"}

// ftClassRefs resolves class and subclass feature references of one file.
type ftClassRefs struct {
	byKey map[string]ftClassFeature // by ftFeatureKey
}

// ftFeatureKey normalizes "Name|Class|ClassSource|Level[|Source]" or, for a
// subclass feature, "Name|Class|ClassSource|Short|SubclassSource|Level[|Source]".
// Empty sources are PHB; the feature's source defaults to its (sub)class's.
func ftFeatureKey(ref string, subclass bool) (key string, level int) {
	n := 5
	if subclass {
		n = 7
	}
	parts := strings.Split(ref, "|")
	parts = append(parts, make([]string, max(n-len(parts), 0))...)[:n]
	parts[2] = cmp.Or(parts[2], "PHB")
	if subclass {
		parts[4] = cmp.Or(parts[4], "PHB")
	}
	parts[n-1] = cmp.Or(parts[n-1], parts[n-3])
	level, _ = strconv.Atoi(parts[n-2])
	return strings.ToLower(strings.Join(parts, "|")), level
}

func (f ftClassFeature) key() string {
	parts := []string{f.Name, f.ClassName, cmp.Or(f.ClassSource, "PHB"), strconv.Itoa(f.Level), f.Source}
	if f.SubclassShortName != "" {
		parts = slices.Insert(parts, 3, f.SubclassShortName, cmp.Or(f.SubclassSource, "PHB"))
	}
	return strings.ToLower(strings.Join(parts, "|"))
}

// text is entries as Markdown with feature references replaced by the
// referenced feature (or its name, when it is not in the file).
func (x ftClassRefs) text(entries []json.RawMessage, depth int) string {
	var v any
	b, _ := json.Marshal(entries)
	_ = json.Unmarshal(b, &v)
	b, _ = json.Marshal(x.expand(v, depth))
	var out []json.RawMessage
	_ = json.Unmarshal(b, &out)
	return ftEntries(out)
}

func (x ftClassRefs) expand(v any, depth int) any {
	switch t := v.(type) {
	case []any:
		for i := range t {
			t[i] = x.expand(t[i], depth)
		}
	case map[string]any:
		var ref string
		var subclass bool
		switch t["type"] {
		case "refClassFeature":
			ref, _ = t["classFeature"].(string)
		case "refSubclassFeature":
			ref, _ = t["subclassFeature"].(string)
			subclass = true
		case "refOptionalfeature":
			ref, _ = t["optionalfeature"].(string)
		default:
			for k, e := range t {
				if k != "type" {
					t[k] = x.expand(e, depth)
				}
			}
			return t
		}
		key, _ := ftFeatureKey(ref, subclass)
		f, ok := x.byKey[key]
		if t["type"] == "refOptionalfeature" || !ok || depth > 3 {
			return ftCode(ref)
		}
		var entries any
		b, _ := json.Marshal(f.Entries)
		_ = json.Unmarshal(b, &entries)
		return map[string]any{"type": "entries", "name": f.Name, "entries": x.expand(entries, depth+1)}
	}
	return v
}

// features resolves feature references in order; features of one name
// gained at several levels (Ability Score Improvement) become one.
func (x ftClassRefs) features(refs []string, subclass bool, gainsSubclass map[string]bool) []ClassFeature {
	var out []ClassFeature
	at := map[string]int{}
	for _, ref := range refs {
		key, level := ftFeatureKey(ref, subclass)
		f, ok := x.byKey[key]
		name, desc := f.Name, ""
		if ok {
			desc = x.text(f.Entries, 0)
		} else {
			name = ftCode(ref)
		}
		i, seen := at[name]
		if !seen {
			i, at[name] = len(out), len(out)
			out = append(out, ClassFeature{Name: name, Desc: desc})
		}
		out[i].Levels = append(out[i].Levels, level)
		out[i].Subclass = out[i].Subclass || gainsSubclass[ref]
	}
	return out
}

func (p *ftBook) classes() error {
	for _, file := range slices.Sorted(maps.Keys(p.files)) {
		if !strings.HasPrefix(file, "class-") {
			continue
		}
		if err := p.classFile(file); err != nil {
			return err
		}
	}
	return nil
}

func (p *ftBook) classFile(file string) error {
	x := ftClassRefs{byKey: map[string]ftClassFeature{}}
	for _, list := range []string{"classFeature", "subclassFeature"} {
		if err := each(p, file, list, func(f ftClassFeature) { x.byKey[f.key()] = f }); err != nil {
			return err
		}
	}
	// The file keeps other books' classes as parents of this book's subclasses.
	if err := each(p, file, "class", func(c ftClass) {
		if !strings.EqualFold(c.Source, p.id) {
			return
		}
		var refs []string
		gains := map[string]bool{}
		for _, raw := range c.ClassFeatures {
			var ref string
			var o struct {
				ClassFeature        string `json:"classFeature"`
				GainSubclassFeature bool   `json:"gainSubclassFeature"`
			}
			if json.Unmarshal(raw, &ref) != nil && json.Unmarshal(raw, &o) == nil {
				ref, gains[o.ClassFeature] = o.ClassFeature, o.GainSubclassFeature
			}
			if ref == "" {
				continue
			}
			refs = append(refs, ref)
		}
		d := ClassData{
			HitDie: c.HD.Faces, Caster: cmp.Or(ftCasters[c.CasterProgression], "none"), Saves: c.Proficiency,
			Features: x.features(refs, false, gains), Table: ftClassTable(c.TableGroups),
		}
		p.add(KindClass, c.Name, d, featuresSearchText(d.Features))
	}); err != nil {
		return err
	}
	return each(p, file, "subclass", func(s ftSubclass) {
		if !strings.EqualFold(s.Source, p.id) {
			return
		}
		d := ClassData{
			Caster: cmp.Or(ftCasters[s.CasterProgression], "none"), SubclassOf: Slugify(s.ClassName),
			Features: x.features(s.SubclassFeatures, true, nil), Table: ftClassTable(s.TableGroups),
		}
		p.add(KindSubclass, s.Name, d, featuresSearchText(d.Features))
	})
}

func featuresSearchText(fs []ClassFeature) string {
	var text []string
	for _, f := range fs {
		text = append(text, f.Name+". "+f.Desc)
	}
	return strings.Join(text, "\n")
}

// ftClassTable is a class table by column and class level; spell slot
// columns are "Slots " and the 5etools label ("Slots 1"). Empty cells
// and zero slots are left out.
func ftClassTable(groups []ftTableGroup) map[string]map[int]string {
	out := map[string]map[int]string{}
	for _, g := range groups {
		rows, prefix := g.Rows, ""
		if g.RowsSpellProgression != nil {
			rows, prefix = g.RowsSpellProgression, "Slots "
		}
		for i, row := range rows {
			for j, cell := range row {
				v := ftClassCell(cell)
				if j >= len(g.ColLabels) || v == "" || v == "—" || v == "-" || (prefix != "" && v == "0") {
					continue
				}
				col := prefix + plainTags(g.ColLabels[j])
				if out[col] == nil {
					out[col] = map[int]string{}
				}
				out[col][i+1] = v
			}
		}
	}
	return out
}

// ftClassCell is a class table cell as people read it: 2, "+2", "1d6",
// "+10 ft.".
func ftClassCell(raw json.RawMessage) string {
	var n float64
	if json.Unmarshal(raw, &n) == nil {
		return strconv.FormatFloat(n, 'f', -1, 64)
	}
	var o struct {
		Type   string  `json:"type"`
		Value  float64 `json:"value"`
		ToRoll []struct {
			Number int `json:"number"`
			Faces  int `json:"faces"`
		} `json:"toRoll"`
	}
	_ = json.Unmarshal(raw, &o)
	val := strconv.FormatFloat(o.Value, 'f', -1, 64)
	if o.Value >= 0 {
		val = "+" + val
	}
	switch o.Type {
	case "bonus":
		return val
	case "bonusSpeed":
		return val + " ft."
	case "dice":
		var dice []string
		for _, d := range o.ToRoll {
			dice = append(dice, strconv.Itoa(d.Number)+"d"+strconv.Itoa(d.Faces))
		}
		return strings.Join(dice, " + ")
	}
	return plainText(ftCell(raw))
}
