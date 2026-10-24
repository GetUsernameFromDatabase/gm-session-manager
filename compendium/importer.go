package compendium

import (
	"cmp"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"

	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
)

// Report describes the import of one source directory.
type Report struct {
	Source  string
	Entries int
	Removed int
	// Unchanged is true when the files' hash matched the last import.
	Unchanged bool
	Err       error
}

// Import reads every source directory under dir (one per document, e.g.
// DATA_DIR/compendium/srd-2024) and imports those whose files changed since
// the last import. A missing dir is not an error: the compendium is
// optional. One bad source does not stop the others.
func Import(ctx context.Context, store *db.Store, dir string) ([]Report, error) {
	dirs, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("compendium: %w", err)
	}
	var reports []Report
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		r := importSource(ctx, store, d.Name(), filepath.Join(dir, d.Name()))
		reports = append(reports, r)
	}
	if err := store.BindMonsterImages(ctx, MonsterImagePrefix); err != nil {
		return reports, fmt.Errorf("compendium: bind monster images: %w", err)
	}
	return reports, nil
}

func importSource(ctx context.Context, store *db.Store, key, path string) Report {
	r := Report{Source: key}
	if key == HomebrewSource {
		r.Err = fmt.Errorf("%q is reserved for homebrew entries", key)
		return r
	}
	files, hash, err := readSource(path)
	if err != nil {
		r.Err = err
		return r
	}
	prev, err := store.GetCompendiumSource(ctx, key)
	if err == nil && prev.Hash == hash {
		r.Unchanged = true
		return r
	}
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		r.Err = err
		return r
	}

	src, entries, err := parseOpen5e(key, files)
	if err != nil {
		r.Err = fmt.Errorf("%s: %w", key, err)
		return r
	}
	r.Removed, r.Err = storeSource(ctx, store, src, path, hash, entries)
	if r.Err == nil {
		r.Entries = len(entries)
	}
	return r
}

// storeSource writes a source's entries, removes those it no longer has
// and records the source; it returns how many were removed.
func storeSource(ctx context.Context, store *db.Store, src Source, path, hash string, entries []Entry) (int, error) {
	removed := 0
	err := store.InTx(ctx, func(q *sqlc.Queries) error {
		seen := map[string]bool{}
		for _, e := range entries {
			if err := upsert(ctx, q, src, e); err != nil {
				return err
			}
			seen[string(e.Kind)+"/"+e.Slug] = true
		}
		existing, err := q.ListCompendiumKeysBySource(ctx, src.Key)
		if err != nil {
			return err
		}
		for _, row := range existing {
			if !seen[row.Kind+"/"+row.Slug] {
				if err := q.DeleteCompendiumEntry(ctx, row.ID); err != nil {
					return err
				}
				removed++
			}
		}
		return q.UpsertCompendiumSource(ctx, sqlc.UpsertCompendiumSourceParams{
			SourceKey: src.Key, Path: path, Hash: hash,
			Edition: src.Edition, License: src.License, Attribution: src.Attribution,
		})
	})
	return removed, err
}

func upsert(ctx context.Context, q *sqlc.Queries, src Source, e Entry) error {
	data, hash, err := e.encode(src)
	if err != nil {
		return err
	}
	err = q.UpsertCompendiumEntry(ctx, sqlc.UpsertCompendiumEntryParams{
		Kind: string(e.Kind), Slug: e.Slug, Name: e.Name, SourceKey: src.Key,
		DataJson: data, BodyText: e.Text, ContentHash: hash,
	})
	if err != nil {
		return fmt.Errorf("store %s %s: %w", e.Kind, e.Slug, err)
	}
	return nil
}

// normalizeVersion changes whenever the normalized shape gains data, so
// unchanged source files are imported again once to fill it in.
const normalizeVersion = "3"

// readSource reads the .json files of a source directory and hashes their
// names and contents, so re-importing unchanged files is a no-op.
func readSource(dir string) (map[string][]byte, string, error) {
	paths, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		return nil, "", err
	}
	slices.Sort(paths)
	files := map[string][]byte{}
	h := sha256.New()
	h.Write([]byte(normalizeVersion))
	for _, p := range paths {
		body, err := os.ReadFile(p) //nolint:gosec // p comes from globbing the operator's data directory
		if err != nil {
			return nil, "", err
		}
		name := filepath.Base(p)
		files[name] = body
		h.Write([]byte(name))
		h.Write([]byte{0})
		h.Write(body)
	}
	if len(files) == 0 {
		return nil, "", fmt.Errorf("%s: no .json files", dir)
	}
	return files, hex.EncodeToString(h.Sum(nil)), nil
}

// ftsQuery turns user input into an FTS5 query: every word must match as a
// prefix. Quoting each word keeps FTS5 syntax characters from the input
// from being interpreted.
func ftsQuery(input string) string {
	words := strings.FieldsFunc(input, func(r rune) bool {
		return r != '\'' && !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	terms := make([]string, 0, len(words))
	for _, w := range words {
		terms = append(terms, `"`+strings.ReplaceAll(w, `"`, `""`)+`"*`)
	}
	return strings.Join(terms, " ")
}

// Search finds entries matching every word of query, best match first:
// every entry whose name matches, then up to limit others whose text does.
// kind "" searches every kind; habitat "" does not filter by habitat.
func Search(ctx context.Context, q *sqlc.Queries, query, kind, habitat string, campaignID int64, limit int) (names, text []sqlc.SearchCompendiumRow, err error) {
	fq := ftsQuery(query)
	if fq == "" {
		return nil, nil, nil
	}
	byName := "name : (" + fq + ")"
	// LIMIT -1 is no limit in SQLite.
	if names, err = q.SearchCompendium(ctx, sqlc.SearchCompendiumParams{Query: byName, Kind: kind, Habitat: habitat, CampaignID: campaignID, MaxResults: -1}); err != nil {
		return nil, nil, err
	}
	text, err = q.SearchCompendium(ctx, sqlc.SearchCompendiumParams{Query: "(" + fq + ") NOT " + byName, Kind: kind, Habitat: habitat, CampaignID: campaignID, MaxResults: int64(limit)})
	return names, text, err
}

// HomebrewInput is the homebrew editor form. Each kind uses its own
// fields: monsters AC, HP and CR; spells the spell fields; items the item,
// weapon and armor fields.
type HomebrewInput struct {
	Kind  string `form:"kind"  validate:"required"`
	Name  string `form:"name"  validate:"required,max=200"`
	Desc  string `form:"desc"  validate:"max=50000"`
	AC    int    `form:"ac"    validate:"min=0,max=50"`
	HP    int    `form:"hp"    validate:"min=0,max=10000"`
	CR    string `form:"cr"    validate:"max=10"`
	Level int    `form:"level" validate:"min=0,max=9"`

	School        string `form:"school"         validate:"max=50"`
	CastingTime   string `form:"casting_time"   validate:"max=100"`
	Range         string `form:"range"          validate:"max=100"`
	Components    string `form:"components"     validate:"max=300"`
	Duration      string `form:"duration"       validate:"max=100"`
	Concentration bool   `form:"concentration"`
	Ritual        bool   `form:"ritual"`
	HigherLevel   string `form:"higher_level"   validate:"max=5000"`
	Classes       string `form:"classes"        validate:"max=300"`

	Category     string `form:"category"     validate:"max=50"`
	Damage       string `form:"damage"       validate:"max=100"`
	Properties   string `form:"properties"   validate:"max=300"`
	WeaponRange  string `form:"weapon_range" validate:"max=20"`
	Martial      bool   `form:"martial"`
	ArmorType    string `form:"armor_type"   validate:"omitempty,oneof=light medium heavy shield"`
	ArmorBase    int    `form:"armor_base"   validate:"min=0,max=30"`
	ArmorAddDex  bool   `form:"armor_add_dex"`
	ArmorDexCap  int    `form:"armor_dex_cap" validate:"min=0,max=10"`
	ArmorStrReq  int    `form:"armor_str_req" validate:"min=0,max=30"`
	ArmorStealth bool   `form:"armor_stealth"`

	// Classes and subclasses; Features and Table use the text of
	// ParseFeatures and ParseTable.
	HitDie     int      `form:"hit_die"     validate:"omitempty,oneof=6 8 10 12"`
	Caster     string   `form:"caster"      validate:"omitempty,oneof=full half pact none"`
	Saves      []string `form:"saves"       validate:"dive,oneof=str dex con int wis cha"`
	SubclassOf string   `form:"subclass_of" validate:"max=200"`
	Features   string   `form:"features"    validate:"max=200000"`
	Table      string   `form:"table"       validate:"max=20000"`
	Armor      string   `form:"armor"       validate:"max=300"`
	Weapons    string   `form:"weapons"     validate:"max=300"`

	// Species, backgrounds and feats; Traits uses the text of ParseTraits.
	// Skills and Tools are shared with classes.
	Speed        int    `form:"speed"        validate:"min=0,max=200"`
	Size         string `form:"size"         validate:"max=100"`
	Traits       string `form:"traits"       validate:"max=100000"`
	Abilities    string `form:"abilities"    validate:"max=200"`
	Skills       string `form:"skills"       validate:"max=300"`
	Tools        string `form:"tools"        validate:"max=300"`
	Languages    string `form:"languages"    validate:"max=300"`
	Feat         string `form:"feat"         validate:"max=200"`
	Prerequisite string `form:"prerequisite" validate:"max=300"`
}

// HomebrewInputFor fills the editor from a stored entry.
func HomebrewInputFor(kind, name, desc string, data any) HomebrewInput {
	in := HomebrewInput{Kind: kind, Name: name, Desc: desc}
	switch d := data.(type) {
	case *MonsterData:
		in.AC, in.HP, in.CR, in.Desc = d.AC, d.HP, d.CR, d.Desc
	case *SpellData:
		in.Level, in.School, in.CastingTime, in.Range, in.Components, in.Duration = d.Level, d.School, d.CastingTime, d.Range, d.Components, d.Duration
		in.Concentration, in.Ritual, in.HigherLevel, in.Classes, in.Desc = d.Concentration, d.Ritual, d.HigherLevel, strings.Join(d.Classes, ", "), d.Desc
	case *ItemData:
		in.Category, in.Damage, in.Properties, in.WeaponRange, in.Martial, in.Desc = d.Category, d.Damage, d.Properties, d.Range, d.Martial, d.Desc
		if a := d.Armor; a != nil {
			in.ArmorType, in.ArmorBase, in.ArmorAddDex, in.ArmorDexCap, in.ArmorStrReq, in.ArmorStealth = a.Type, a.Base, a.AddDex, a.DexCap, a.StrReq, a.Stealth
		}
	case *ClassData:
		in.HitDie, in.Caster, in.Saves, in.SubclassOf, in.Desc = d.HitDie, d.Caster, d.Saves, d.SubclassOf, withoutTable(d.Desc)
		in.Features, in.Table = FeaturesText(d.Features), TableText(d.Table)
		traits := CoreTraits(d.Desc)
		in.Skills, in.Armor, in.Weapons, in.Tools = traits[classTraits[0]], traits[classTraits[1]], traits[classTraits[2]], traits[classTraits[3]]
	case *TextData:
		in.Desc, in.Prerequisite, in.Traits = d.Desc, d.Prerequisite, TraitsText(d.Parts)
		for _, p := range d.Parts {
			switch p.Type {
			case partSize:
				in.Size = p.Desc
			case partSpeed:
				in.Speed = firstInt(p.Desc)
			case partAbilities:
				in.Abilities = p.Desc
			case partSkills:
				in.Skills = p.Desc
			case partTools:
				in.Tools = p.Desc
			case partLanguages:
				in.Languages = p.Desc
			case partFeat:
				in.Feat = p.Desc
			}
		}
	}
	return in
}

// InputForEntry is the homebrew form for a parsed entry, so it can be
// reviewed and edited before it is saved.
func InputForEntry(e Entry) HomebrewInput {
	data := e.Data
	switch d := data.(type) {
	case ClassData:
		data = &d
	case TextData:
		data = &d
	}
	return HomebrewInputFor(string(e.Kind), e.Name, e.Text, data)
}

// Typed part names (NameDesc.Type) the character wizard applies.
const (
	partSize      = "size"
	partSpeed     = "speed"
	partAbilities = "ability_score"
	partSkills    = "skill_proficiency"
	partTools     = "tool_proficiency"
	partLanguages = "languages"
	partFeat      = "feat"
)

var reFirstInt = regexp.MustCompile(`\d+`)

func firstInt(s string) int {
	n, _ := strconv.Atoi(reFirstInt.FindString(s))
	return n
}

// typedParts are the typed parts the form fills, then the traits.
func (in HomebrewInput) typedParts() []NameDesc {
	var out []NameDesc
	add := func(typ, name, desc string) {
		if desc = strings.TrimSpace(desc); desc != "" {
			out = append(out, NameDesc{Name: name, Desc: desc, Type: typ})
		}
	}
	add(partSize, "Size", in.Size)
	if in.Speed > 0 {
		add(partSpeed, "Speed", fmt.Sprintf("%d feet", in.Speed))
	}
	add(partAbilities, "Ability Scores", in.Abilities)
	add(partSkills, "Skill Proficiencies", in.Skills)
	add(partTools, "Tool Proficiencies", in.Tools)
	add(partLanguages, "Languages", in.Languages)
	add(partFeat, "Feat", in.Feat)
	return append(out, ParseTraits(in.Traits)...)
}

// homebrewSource attributes homebrew to the table that wrote it; its
// compendium_sources row is created in 0001_init.sql.
var homebrewSource = Source{Key: HomebrewSource, Edition: "2024", License: "Homebrew", Attribution: "Homebrew content written for this table."}

// Entry converts the form into a normalized entry of its kind.
func (in HomebrewInput) Entry() (Entry, error) {
	if !ValidKind(in.Kind) {
		return Entry{}, fmt.Errorf("unknown kind %q", in.Kind)
	}
	e := Entry{Kind: Kind(in.Kind), Slug: Slugify(in.Name), Name: strings.TrimSpace(in.Name), Text: in.Desc}
	if e.Slug == "" {
		return Entry{}, errors.New("name needs at least one letter or digit")
	}
	switch e.Kind {
	case KindMonster:
		e.Data = MonsterData{AC: in.AC, HP: in.HP, CR: in.CR, Desc: in.Desc}
	case KindSpell:
		var classes []string
		for _, c := range strings.Split(in.Classes, ",") {
			if c = Slugify(c); c != "" {
				classes = append(classes, c)
			}
		}
		e.Data = SpellData{
			Level: in.Level, School: strings.ToLower(strings.TrimSpace(in.School)), CastingTime: in.CastingTime, Range: in.Range,
			Components: in.Components, Duration: in.Duration, Concentration: in.Concentration, Ritual: in.Ritual,
			Classes: classes, Desc: in.Desc, HigherLevel: in.HigherLevel,
		}
	case KindItem:
		d := ItemData{Category: in.Category, Damage: in.Damage, Properties: in.Properties, Range: in.WeaponRange, Martial: in.Martial, Desc: in.Desc}
		if in.ArmorType != "" {
			d.Armor = &ArmorData{Type: in.ArmorType, Base: in.ArmorBase, AddDex: in.ArmorAddDex, DexCap: in.ArmorDexCap, StrReq: in.ArmorStrReq, Stealth: in.ArmorStealth}
			d.ArmorClass = armorText(*d.Armor)
		}
		e.Data = d
	case KindClass, KindSubclass:
		d := ClassData{
			HitDie: in.HitDie, Caster: cmp.Or(in.Caster, "none"), Saves: in.Saves, Features: ParseFeatures(in.Features), Table: ParseTable(in.Table),
			Desc: withCoreTraits(in.Desc, map[string]string{classTraits[0]: in.Skills, classTraits[1]: in.Armor, classTraits[2]: in.Weapons, classTraits[3]: in.Tools}),
		}
		if e.Kind == KindSubclass {
			d.SubclassOf = Slugify(in.SubclassOf)
		}
		e.Data = d
		e.Text = joinText(in.Desc, FeaturesText(d.Features))
	default:
		d := TextData{Desc: in.Desc, Prerequisite: strings.TrimSpace(in.Prerequisite), Parts: in.typedParts()}
		e.Data = d
		e.Text = joinText(in.Desc, partsText(d.Parts))
	}
	return e, nil
}

// HomebrewOwner is who writes a new homebrew entry: the GM (zero value,
// global) or a player's character, whose entries belong to its campaign.
type HomebrewOwner struct {
	CharacterID, CampaignID int64
}

// SaveHomebrew creates (id 0) or updates a homebrew entry. A player's new
// entry is linked to their campaign and gets a slug of its own, so it never
// collides with another player's entry of the same name.
func SaveHomebrew(ctx context.Context, q *sqlc.Queries, id int64, in HomebrewInput, owner HomebrewOwner) (int64, error) {
	e, err := in.Entry()
	if err != nil {
		return 0, err
	}
	if owner.CharacterID != 0 {
		e.Slug = fmt.Sprintf("%s-%d", e.Slug, owner.CharacterID)
	}
	data, hash, err := e.encode(homebrewSource)
	if err != nil {
		return 0, err
	}
	if id != 0 {
		return id, q.UpdateHomebrew(ctx, sqlc.UpdateHomebrewParams{ID: id, Name: e.Name, DataJson: data, BodyText: e.Text, ContentHash: hash})
	}
	row, err := q.CreateHomebrew(ctx, sqlc.CreateHomebrewParams{
		Kind: string(e.Kind), Slug: e.Slug, Name: e.Name, DataJson: data, BodyText: e.Text, ContentHash: hash,
		OwnerCharacterID: sql.NullInt64{Int64: owner.CharacterID, Valid: owner.CharacterID != 0}, Scoped: owner.CharacterID != 0,
	})
	if err != nil {
		return 0, fmt.Errorf("a homebrew %s named like %q may already exist: %w", e.Kind, e.Name, err)
	}
	if owner.CampaignID != 0 {
		err = q.LinkEntryCampaign(ctx, sqlc.LinkEntryCampaignParams{EntryID: row.ID, CampaignID: owner.CampaignID})
	}
	return row.ID, err
}
