package views

import (
	"slices"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/perm"
)

// SuggestedFields are the NPC fields offered when creating or editing an
// entity. They are a frontend recommendation only: any key may be stored.
var SuggestedFields = []string{
	fieldRole, fieldDescription, fieldAppearance, fieldPersonality, fieldMotivation,
	fieldLocation, fieldFaction, fieldRelationship, fieldStatBlock, fieldGMNotes,
}

// The suggested NPC field keys; fieldRole is also shown under the name.
const (
	fieldRole         = "role"
	fieldDescription  = "description"
	fieldAppearance   = "appearance"
	fieldPersonality  = "personality"
	fieldMotivation   = "motivation"
	fieldLocation     = "location"
	fieldFaction      = "faction"
	fieldRelationship = "relationship to the party"
	fieldStatBlock    = "stat block"
	fieldGMNotes      = "gm notes"
)

// OrderedFieldKeys lists suggested keys first, in order, then the rest
// alphabetically.
func OrderedFieldKeys(f perm.Fields) []string {
	var keys []string
	for _, k := range SuggestedFields {
		if _, ok := f[k]; ok {
			keys = append(keys, k)
		}
	}
	var rest []string
	for k := range f {
		if !slices.Contains(SuggestedFields, k) {
			rest = append(rest, k)
		}
	}
	slices.Sort(rest)
	return append(keys, rest...)
}

// EntityListView is the NPC index.
type EntityListView struct {
	CampaignID int64
	Entities   []perm.Entity
	IsGM       bool
}

const newEntityID = "new-npc"

// currentLook is the look shown as the portrait.
func currentLook(looks []perm.Look) (perm.Look, bool) {
	i := slices.IndexFunc(looks, func(l perm.Look) bool { return l.IsCurrent })
	if i < 0 {
		return perm.Look{}, false
	}
	return looks[i], true
}

// EntityView is one entity page.
type EntityView struct {
	CampaignID int64
	E          perm.Entity
	IsGM       bool
	// Compendium is the linked statblock (GM only).
	Compendium *sqlc.CompendiumEntry
	// ActiveEncounterID lets the GM add the NPC to the running fight.
	ActiveEncounterID int64
	// ImageChoices are the campaign's images not yet one of its looks (GM only).
	ImageChoices []sqlc.ListImagesRow
	// Readable are the [[link]] targets in the fields a player may open
	// (keys as WikiLink.Key); the others show as their label.
	Readable    map[string]bool
	Backlinks   []sqlc.ListBacklinksRow
	EntityLinks []sqlc.ListEntityFieldBacklinksRow
}

func (v EntityView) canOpen(l WikiLink) bool {
	return l.Kind != LinkCompendium && (v.IsGM || v.Readable[l.Key()])
}

// entityGrid is the page's two columns: looks, then fields or mentions.
const entityGrid = "grid gap-6 md:grid-cols-[2fr_3fr] md:gap-8"

// LookItem is a look as an image to load.
func LookItem(img perm.Look) ImageItem {
	return ImageItem{ID: img.ImageID, Name: img.Name, Version: img.Version}
}

func missingSuggestions(keys []string) []string {
	var out []string
	for _, k := range SuggestedFields {
		if !slices.Contains(keys, k) {
			out = append(out, k)
		}
	}
	return out
}
