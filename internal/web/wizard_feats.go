package web

import (
	"context"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/dnd5e"
	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/web/views"
)

// originFeatChoice offers the compendium's Origin feats (Human Versatile),
// each with its own choices.
func (a *App) originFeatChoice(ctx context.Context, p compendium.NameDesc, sh dnd5e.Sheet, campaignID int64) (choiceGroup, bool, error) {
	names, err := a.Store.ListCompendiumNames(ctx, sqlc.ListCompendiumNamesParams{Kind: enum.KindFeat, CampaignID: campaignID})
	if err != nil {
		return choiceGroup{}, false, err
	}
	slugs := make([]string, len(names))
	for i, n := range names {
		slugs[i] = n.Slug
	}
	// One query for every feat's priority copy, as entryFor picks unpinned.
	rows, err := a.Store.ListCompendiumBySlugs(ctx, sqlc.ListCompendiumBySlugsParams{Kind: enum.KindFeat, Slugs: slugs})
	if err != nil {
		return choiceGroup{}, false, err
	}
	bySlug := make(map[string]sqlc.ListCompendiumBySlugsRow, len(rows))
	for _, r := range rows {
		bySlug[r.Slug] = r
	}
	key := "trait_" + compendium.Slugify(p.Name)
	var opts []compendium.NameDesc
	sub := map[string][]choiceGroup{}
	sources := map[string]string{}
	for _, n := range names {
		row, ok := bySlug[n.Slug]
		if !ok {
			continue
		}
		td, err := compendium.Decode[compendium.TextData](enum.KindFeat, row.DataJson)
		if err != nil {
			return choiceGroup{}, false, err
		}
		if td.Category != "Origin" || !dnd5e.MeetsPrerequisite(td.Prerequisite, sh, sh.Level(), nil, compendium.Slugify) {
			continue
		}
		sources[n.Name] = row.SourceKey
		opts = append(opts, compendium.NameDesc{Name: n.Name, Desc: textMarkdown(td)})
		groups, err := a.featChoices(ctx, key+"_"+n.Slug, n.Name, td, "", sh, campaignID)
		if err != nil {
			return choiceGroup{}, false, err
		}
		if len(groups) > 0 {
			sub[n.Name] = groups
		}
	}
	if len(opts) == 0 {
		return choiceGroup{}, false, nil
	}
	g := views.ChoiceGroup{Key: key, Title: p.Name + ": an Origin feat", Count: 1, Options: opts}
	return choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
		o := opts[slices.IndexFunc(opts, func(o compendium.NameDesc) bool { return o.Name == picked[0] })]
		dnd5e.AddFeat(s, dnd5e.Feature{Name: o.Name, Source: dnd5e.SourceFeat, Description: o.Desc})
		pinFeat(s, o.Name, sources[o.Name])
	}, sub}, true, nil
}

// pinFeat records the source of the feat's compendium copy under featKey.
func pinFeat(s *dnd5e.Sheet, name, source string) {
	if source == "" {
		return
	}
	if s.FeatSources == nil {
		s.FeatSources = map[string]string{}
	}
	s.FeatSources[featKey(name)] = source
}

// featKey is the key a feat's pin is stored under: the feat's slug without
// a "(choice)" suffix.
func featKey(name string) string {
	base, _, _ := strings.Cut(name, " (")
	return compendium.Slugify(base)
}

// backgroundFeatChoices are the choices inside the feat the background
// gave ("Magic Initiate (Wizard)").
func (a *App) backgroundFeatChoices(ctx context.Context, sh dnd5e.Sheet, campaignID int64) ([]choiceGroup, error) {
	if sh.Background == "" {
		return nil, nil
	}
	bg, _, err := a.textEntrySource(ctx, enum.KindBackground, compendium.Slugify(sh.Background), sh.BackgroundSource)
	if err != nil || bg == nil {
		return nil, err
	}
	for _, p := range bg.Parts {
		full := strings.TrimSpace(p.Desc)
		if p.Type != compendium.PartFeat || !slices.Contains(sh.Feats, full) {
			continue
		}
		name, list := featList(full)
		slug := compendium.Slugify(name)
		td, _, err := a.textEntrySource(ctx, enum.KindFeat, slug, sh.FeatSources[slug])
		if err != nil || td == nil {
			return nil, err
		}
		return a.featChoices(ctx, "feat_"+compendium.Slugify(name), full, td, list, sh, campaignID)
	}
	return nil, nil
}

// featList splits "Magic Initiate (Wizard)" into the feat and its spell list.
func featList(full string) (string, string) {
	name, list, _ := strings.Cut(full, " (")
	return name, strings.TrimSuffix(list, ")")
}

var (
	reFeatCantrips = regexp.MustCompile(`learn (\w+) cantrips of your choice from the (.+?) spell list`)
	reFeatSpell    = regexp.MustCompile(`Choose a level 1 spell from the same list`)
	reFeatProfs    = regexp.MustCompile(`proficiency in any combination of (\w+) skills or tools`)
	reFeatRaise    = regexp.MustCompile(`Increase (?:your (?:((?:\w+, )*\w+) or )?(\w+) score|one ability score of your choice) by 1, to a maximum of (\d+)`)
	reFeatMastery  = regexp.MustCompile(`mastery property of (\w+) kinds? of`)
	reToolCost     = regexp.MustCompile(`\s*\(\d+ GP\)$`)
)

// featText is a feat's description and parts, for the reFeat patterns.
func featText(td *compendium.TextData) string {
	text := td.Desc
	for _, p := range td.Parts {
		text += "\n" + p.Desc
	}
	return text
}

// featMasteries is how many weapon kinds a feat lets the sheet master
// (Weapon Master), read from its text so homebrew feats count too.
func featMasteries(td *compendium.TextData) int {
	if m := reFeatMastery.FindStringSubmatch(featText(td)); m != nil {
		return dnd5e.CountWord(m[1])
	}
	return 0
}

// featChoices are the choices a feat's text asks for: Magic Initiate's spell
// list, cantrips, level 1 spell and spellcasting ability, Skilled's skills
// and tools, and an ability increase (Grappler, Epic Boons). name is the
// feature on the sheet; list is the spell list a background names. Lists the
// sheet took with this feat already are left out (Repeatable).
func (a *App) featChoices(ctx context.Context, key, name string, td *compendium.TextData, list string, sh dnd5e.Sheet, campaignID int64) ([]choiceGroup, error) {
	text, count := featText(td), dnd5e.CountWord
	var out []choiceGroup
	if m := reFeatCantrips.FindStringSubmatch(text); m != nil {
		lists := dnd5e.SplitList(m[2])
		if list != "" {
			lists = []string{list}
		}
		added := map[string]bool{} // the spells the picks add, for the ability
		perList := map[string][]choiceGroup{}
		var opts []compendium.NameDesc
		for _, l := range lists {
			if list == "" && slices.Contains(sh.Feats, fmt.Sprintf("%s (%s)", name, l)) {
				continue
			}
			full := name
			if list == "" {
				full = fmt.Sprintf("%s (%s)", name, l)
			}
			groups, err := a.featSpellGroups(ctx, key+"_"+compendium.Slugify(l), full, l, count(m[1]), reFeatSpell.MatchString(text), sh, campaignID, added)
			if err != nil {
				return nil, err
			}
			if len(groups) > 0 {
				perList[l] = groups
				opts = append(opts, compendium.NameDesc{Name: l})
			}
		}
		if list != "" {
			out = append(out, perList[list]...)
		} else if len(opts) > 0 {
			g := views.ChoiceGroup{Key: key + "_list", Title: name + ": spell list", Count: 1, Options: opts}
			out = append(out, choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
				full := fmt.Sprintf("%s (%s)", name, picked[0])
				if i := slices.IndexFunc(s.Features, func(f dnd5e.Feature) bool { return f.Name == name && f.Source == dnd5e.SourceFeat }); i >= 0 {
					s.Features[i].Name = full
				}
				if i := slices.Index(s.Feats, name); i >= 0 {
					s.Feats[i] = full
				}
			}, perList})
		}
		if len(out) > 0 && reSpellAbility.MatchString(text) {
			out = append(out, spellAbilityChoice(key+"_ability", name, added, false))
		}
	}
	if m := reFeatRaise.FindStringSubmatch(text); m != nil {
		allowed := dnd5e.Abilities
		if m[2] != "" { // "your X, Y or Z", or a fixed "your X"
			allowed = dnd5e.AbilityKeys(strings.TrimPrefix(m[1]+", "+m[2], ", "))
		}
		limit, _ := strconv.Atoi(m[3])
		var opts []compendium.NameDesc
		for _, ab := range allowed {
			if sh.Scores[ab] < limit {
				opts = append(opts, compendium.NameDesc{Name: dnd5e.AbilityNames[ab]})
			}
		}
		// A fixed "your X" already at its maximum asks nothing: the feat is still taken.
		if fixedAtMax := len(opts) == 0 && m[1] == "" && m[2] != ""; !fixedAtMax {
			g := views.ChoiceGroup{Key: key + "_raise", Title: name + ": +1 to an ability score", Count: 1, Options: opts}
			out = append(out, choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
				if ab := dnd5e.AbilityKeys(picked[0]); len(ab) == 1 {
					s.Scores[ab[0]]++
				}
			}, nil})
		}
	}
	if m := reFeatProfs.FindStringSubmatch(text); m != nil {
		var opts []compendium.NameDesc
		for _, sk := range dnd5e.Skills {
			if !slices.Contains(sh.SkillProfs, sk.Key) {
				opts = append(opts, compendium.NameDesc{Name: sk.Name})
			}
		}
		tools, err := a.toolNames(ctx, campaignID)
		if err != nil {
			return nil, err
		}
		for _, t := range tools {
			if !slices.Contains(sh.OtherProfs, t) {
				opts = append(opts, compendium.NameDesc{Name: t})
			}
		}
		g := views.ChoiceGroup{Key: key + "_profs", Title: name + ": skills or tools", Count: count(m[1]), Options: opts}
		out = append(out, choiceGroup{g, func(s *dnd5e.Sheet, picked []string) {
			tools := slices.DeleteFunc(slices.Clone(picked), func(p string) bool {
				return slices.ContainsFunc(dnd5e.Skills, func(sk dnd5e.Skill) bool { return sk.Name == p })
			})
			s.AddProfs(name, dnd5e.SkillKeys(picked), tools...)
		}, nil})
	}
	if n := featMasteries(td); n > 0 {
		groups, err := a.masteryChoice(ctx, sh, key+"_"+views.FieldMastery, name+": weapon mastery", n, campaignID)
		if err != nil {
			return nil, err
		}
		out = append(out, groups...)
	}
	return out, nil
}

// featSpellGroups ask for a feat's cantrips and level 1 spell (always
// prepared, one free cast per Long Rest) from a class's list, leaving out
// spells the sheet has; added collects the picks.
func (a *App) featSpellGroups(ctx context.Context, key, name, class string, cantrips int, spell bool, sh dnd5e.Sheet, campaignID int64, added map[string]bool) ([]choiceGroup, error) {
	rows, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: compendium.Slugify(class), MaxLevel: 1, CampaignID: campaignID})
	if err != nil {
		return nil, err
	}
	var byLevel [2][]compendium.NameDesc
	spells := map[string]dnd5e.Spell{}
	for _, r := range rows {
		if slices.ContainsFunc(sh.Spells, func(have dnd5e.Spell) bool { return have.Slug == r.Slug }) {
			continue
		}
		byLevel[r.Level] = append(byLevel[r.Level], compendium.NameDesc{Name: r.Name})
		spells[r.Name] = dnd5e.Spell{Slug: r.Slug, Name: r.Name, Source: r.SourceKey, Level: int(r.Level), Feat: name}
	}
	add := func(s *dnd5e.Sheet, picked []string) {
		for _, n := range picked {
			sp := spells[n]
			if sp.Level > 0 {
				sp.Prepared = true
				s.PutFeature(dnd5e.FreeCast(sp, name), nil)
			}
			s.AddSpell(sp)
			added[sp.Slug] = true
		}
	}
	var out []choiceGroup
	if cantrips > 0 && len(byLevel[0]) > 0 {
		out = append(out, choiceGroup{views.ChoiceGroup{Key: key + "_cantrips", Title: name + ": cantrips", Count: cantrips, Options: byLevel[0]}, add, nil})
	}
	if spell && len(byLevel[1]) > 0 {
		out = append(out, choiceGroup{views.ChoiceGroup{Key: key + "_spell", Title: name + ": level 1 spell", Count: 1, Options: byLevel[1]}, add, nil})
	}
	return out, nil
}

// keepSpells is Spell Change's option to change nothing.
const keepSpells = "No change"

// featSpellChanges offer each feat's Spell Change on level-up: one spell the
// feat chose ("Magic Initiate (Wizard)") may become another of the same
// level from its list, but not one the sheet has or this step teaches (taught).
func (a *App) featSpellChanges(ctx context.Context, sh dnd5e.Sheet, taught []dnd5e.Spell, campaignID int64) ([]choiceGroup, error) {
	known := slices.Concat(sh.Spells, taught)
	var out []choiceGroup
	for _, feat := range sh.Feats {
		_, list := featList(feat)
		mine := slices.DeleteFunc(slices.Clone(sh.Spells), func(sp dnd5e.Spell) bool { return sp.Feat != feat })
		if list == "" || len(mine) == 0 {
			continue
		}
		rows, err := a.Store.ListSpellsForClass(ctx, sqlc.ListSpellsForClassParams{ClassSlug: compendium.Slugify(list), MaxLevel: 1, CampaignID: campaignID})
		if err != nil {
			return nil, err
		}
		key := "change_" + compendium.Slugify(feat)
		opts := []compendium.NameDesc{{Name: keepSpells}}
		sub := map[string][]choiceGroup{}
		for _, old := range mine {
			var with []compendium.NameDesc
			spells := map[string]dnd5e.Spell{}
			for _, r := range rows {
				if int(r.Level) == old.Level && !slices.ContainsFunc(known, func(h dnd5e.Spell) bool { return h.Slug == r.Slug }) {
					with = append(with, compendium.NameDesc{Name: r.Name})
					spells[r.Name] = dnd5e.Spell{Slug: r.Slug, Name: r.Name, Source: r.SourceKey, Level: int(r.Level)}
				}
			}
			if len(with) == 0 {
				continue
			}
			opts = append(opts, compendium.NameDesc{Name: old.Name})
			g := views.ChoiceGroup{Key: key + "_" + old.Slug, Title: "Replace " + old.Name + " with", Count: 1, Options: with}
			sub[old.Name] = []choiceGroup{{g, func(s *dnd5e.Sheet, picked []string) { dnd5e.SwapFeatSpell(s, old, spells[picked[0]]) }, nil}}
		}
		if len(sub) > 0 {
			out = append(out, choiceGroup{views.ChoiceGroup{Key: key, Title: feat + ": Spell Change", Count: 1, Options: opts}, nil, sub})
		}
	}
	return out, nil
}
