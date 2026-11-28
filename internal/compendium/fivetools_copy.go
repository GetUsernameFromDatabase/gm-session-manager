package compendium

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
)

// ftNotCopied are fields a copy does not inherit (its own pictures and
// page), unless _preserve names them.
var ftNotCopied = []string{
	keyPage,
	keyHasFluff,
	keyHasFluffImages,
	keyHasToken,
	keySrd,
	keySrd52,
	keyBasicRules,
	keyBasicRules2024,
	keyReprintedAs,
	keyOtherSources,
	keyVersions,
}

var errCopy = errors.New("unresolved _copy")

// ftCopyRef is what a _copy names. Class fields count too: a book keeps its
// subclasses once per class edition under the same name and source.
type ftCopyRef struct {
	ftNamed
	ClassName         string `json:"className"`
	ClassSource       string `json:"classSource"`
	ShortName         string `json:"shortName"`
	SubclassShortName string `json:"subclassShortName"`
	SubclassSource    string `json:"subclassSource"`
	Level             any    `json:"level"`
}

func (c ftCopyRef) key() string {
	parts := []string{
		c.Name, c.Source, c.ClassName, c.ClassSource, c.ShortName,
		c.SubclassShortName, c.SubclassSource, fmt.Sprint(c.Level),
	}
	return strings.ToLower(strings.Join(parts, "|"))
}

// ftResolve applies a record's _copy: the target's fields, then its own,
// then _mod.
func ftResolve(index map[string]ftRec, r ftRec, depth int) (ftRec, error) {
	raw, ok := r[keyCopy]
	if !ok {
		return r, nil
	}
	var c struct {
		ftCopyRef
		Mod      map[string]json.RawMessage `json:"_mod"`
		Preserve map[string]bool            `json:"_preserve"`
	}
	if depth > 4 || json.Unmarshal(raw, &c) != nil {
		return nil, errCopy
	}
	// A _copy that names fewer fields than its target falls back to name|source.
	target, ok := index[c.key()]
	if !ok {
		target, ok = index[c.ftNamed.key()]
	}
	if !ok {
		return nil, errCopy
	}
	base, err := ftResolve(index, target, depth+1)
	if err != nil {
		return nil, err
	}
	out := maps.Clone(base)
	for _, k := range ftNotCopied {
		if !c.Preserve[k] {
			delete(out, k)
		}
	}
	for k, v := range r {
		if k != keyCopy {
			out[k] = v
		}
	}
	for _, k := range slices.Sorted(maps.Keys(c.Mod)) {
		if err := ftMod(out, k, c.Mod[k]); err != nil {
			return nil, err
		}
	}
	return out, nil
}

// ftVersions are a record's _versions as records of their own: each copies r
// with its own fields and _mod. An _abstract version is a template, filled
// once per implementation with its _variables ("{{color}}") and fields.
// "Elf; Drow Lineage" is named "Elf (Drow Lineage)", as subraces are.
func ftVersions(r ftRec) ([]ftRec, int) {
	raw, ok := r[keyVersions]
	if !ok {
		return nil, 0
	}
	parent, err := ftDecode[ftNamed](r)
	var vs []ftRec
	if err != nil || json.Unmarshal(raw, &vs) != nil {
		return nil, 1
	}
	var all []ftRec
	skipped := 0
	for _, v := range vs {
		abstract, ok := v[keyAbstract]
		if !ok {
			all = append(all, v)
			continue
		}
		var impls []ftRec
		_ = json.Unmarshal(v[keyImplementations], &impls)
		for _, impl := range impls {
			var vars map[string]string
			_ = json.Unmarshal(impl[keyVariables], &vars)
			body := string(abstract)
			for k, val := range vars {
				quoted, _ := json.Marshal(val) // the value lands inside JSON strings
				body = strings.ReplaceAll(body, "{{"+k+"}}", string(quoted[1:len(quoted)-1]))
			}
			var ver ftRec
			if json.Unmarshal([]byte(body), &ver) != nil {
				skipped++
				continue
			}
			for k, x := range impl {
				if k != keyVariables {
					ver[k] = x
				}
			}
			all = append(all, ver)
		}
	}
	index := map[string]ftRec{parent.key(): r}
	var out []ftRec
	for _, v := range all {
		ref, _ := json.Marshal(
			map[string]json.RawMessage{
				keyMod:    v[keyMod],
				keyName:   r[keyName],
				keySource: r[keySource],
			},
		)
		ver := maps.Clone(v)
		delete(ver, keyMod)
		ver[keyCopy] = ref
		res, err := ftResolve(index, ver, 0)
		if err != nil {
			skipped++
			continue
		}
		var name string
		if json.Unmarshal(res[keyName], &name) == nil {
			if base, version, ok := strings.Cut(name, "; "); ok {
				res[keyName], _ = json.Marshal(base + " (" + version + ")")
			}
		}
		out = append(out, res)
	}
	return out, skipped
}

type ftModOp struct {
	Mode    string          `json:"mode"`
	Replace json.RawMessage `json:"replace"`
	With    string          `json:"with"`
	Flags   string          `json:"flags"`
	Items   json.RawMessage `json:"items"`
	Index   int             `json:"index"`
	Names   json.RawMessage `json:"names"`
}

// reJSGroup is a JavaScript replacement group ($1), ${1} in Go.
var reJSGroup = regexp.MustCompile(`\$(\d)`)

// ftMod applies one key's _mod operations; "*" is every field. Unknown
// modes are an error, so the record is skipped rather than wrong.
func ftMod(rec ftRec, key string, raw json.RawMessage) error {
	var s string
	if json.Unmarshal(raw, &s) == nil {
		if s != keyRemove {
			return errCopy
		}
		delete(rec, key)
		return nil
	}
	var ops []ftModOp
	if json.Unmarshal(raw, &ops) != nil {
		var op ftModOp
		if err := json.Unmarshal(raw, &op); err != nil {
			return err
		}
		ops = []ftModOp{op}
	}
	for _, op := range ops {
		if op.Mode == modReplaceTxt {
			var pat string
			if json.Unmarshal(op.Replace, &pat) != nil || strings.Trim(op.Flags, "gi") != "" {
				return errCopy
			}
			if strings.Contains(op.Flags, "i") {
				pat = "(?i)" + pat
			}
			re, err := regexp.Compile(pat)
			if err != nil {
				return errCopy
			}
			with := reJSGroup.ReplaceAllString(op.With, "$${$1}")
			for k, v := range rec {
				if (key == keyModAll && k != keyName && k != keySource && !strings.HasPrefix(k, "_")) ||
					k == key {
					var val any
					_ = json.Unmarshal(v, &val)
					rec[k], _ = json.Marshal(ftReplaceText(val, re, with))
				}
			}
			continue
		}
		var arr, items []any
		_ = json.Unmarshal(rec[key], &arr)
		if json.Unmarshal(op.Items, &items) != nil && op.Items != nil {
			var one any
			_ = json.Unmarshal(op.Items, &one)
			items = []any{one}
		}
		switch op.Mode {
		case modAppendArr:
			arr = append(arr, items...)
		case modPrependArr:
			arr = append(items, arr...)
		case modInsertArr:
			arr = slices.Insert(arr, min(max(op.Index, 0), len(arr)), items...)
		case modRemoveArr:
			var names []string
			if json.Unmarshal(op.Names, &names) != nil {
				var one string
				_ = json.Unmarshal(op.Names, &one)
				names = []string{one}
			}
			for _, it := range items {
				if s, ok := it.(string); ok {
					names = append(names, s)
				}
			}
			arr = slices.DeleteFunc(
				arr,
				func(v any) bool { return slices.Contains(names, ftName(v)) },
			)
		case modReplaceArr:
			var name string
			_ = json.Unmarshal(op.Replace, &name)
			i := slices.IndexFunc(arr, func(v any) bool { return ftName(v) == name })
			if i < 0 {
				return errCopy
			}
			arr = slices.Replace(arr, i, i+1, items...)
		default:
			return errCopy
		}
		rec[key], _ = json.Marshal(arr)
	}
	return nil
}

// ftName is a list element's name: the string itself or its "name".
func ftName(v any) string {
	if s, ok := v.(string); ok {
		return s
	}
	if m, ok := v.(map[string]any); ok {
		s, _ := m[keyName].(string)
		return s
	}
	return ""
}

// ftReplaceText runs re over every string in v.
func ftReplaceText(v any, re *regexp.Regexp, with string) any {
	switch x := v.(type) {
	case string:
		return re.ReplaceAllString(x, with)
	case []any:
		for i := range x {
			x[i] = ftReplaceText(x[i], re, with)
		}
	case map[string]any:
		for k := range x {
			if k != keyType {
				x[k] = ftReplaceText(x[k], re, with)
			}
		}
	}
	return v
}
