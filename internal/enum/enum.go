// Package enum holds the values SQL CHECK constraints allow, typed. sqlc
// overrides (sqlc.yaml) give the generated rows these types, so
// it must import nothing from the app.
package enum

// Role is sessions.role.
type Role string

// Session roles. The GM is an admin; there is no separate GM role.
const (
	RolePlayer Role = "player"
	RoleAdmin  Role = "admin"
)

// CompendiumKind is compendium_entries.kind.
type CompendiumKind string

// Compendium entry kinds.
const (
	KindMonster    CompendiumKind = "monster"
	KindSpell      CompendiumKind = "spell"
	KindItem       CompendiumKind = "item"
	KindClass      CompendiumKind = "class"
	KindSubclass   CompendiumKind = "subclass"
	KindSpecies    CompendiumKind = "species"
	KindBackground CompendiumKind = "background"
	KindFeat       CompendiumKind = "feat"
	KindCondition  CompendiumKind = "condition"
	KindRule       CompendiumKind = "rule"
)

// Habitat is monster_habitats.habitat.
type Habitat string

// Monster habitats. Any is allowed by the CHECK but filters nothing.
const (
	HabitatAny        Habitat = "Any"
	HabitatArctic     Habitat = "Arctic"
	HabitatCoastal    Habitat = "Coastal"
	HabitatDesert     Habitat = "Desert"
	HabitatForest     Habitat = "Forest"
	HabitatGrassland  Habitat = "Grassland"
	HabitatHill       Habitat = "Hill"
	HabitatMountain   Habitat = "Mountain"
	HabitatPlanar     Habitat = "Planar"
	HabitatSwamp      Habitat = "Swamp"
	HabitatUnderdark  Habitat = "Underdark"
	HabitatUnderwater Habitat = "Underwater"
	HabitatUrban      Habitat = "Urban"
)

// EncounterStatus is encounters.status: preparing → active → ended, and
// ended → active on restart.
type EncounterStatus string

// Encounter statuses.
const (
	EncounterPreparing EncounterStatus = "preparing"
	EncounterActive    EncounterStatus = "active"
	EncounterEnded     EncounterStatus = "ended"
)

// EffectEnd is conditions.ends_on: the rest that removes an effect.
type EffectEnd string

// Effect end triggers; none means only a person or the round count ends it.
const (
	EffectEndNone      EffectEnd = ""
	EffectEndShortRest EffectEnd = "short_rest"
	EffectEndLongRest  EffectEnd = "long_rest"
)

// NoteKind is notes.kind.
type NoteKind string

// Note kinds.
const (
	NoteGameSession NoteKind = "game_session"
	NoteGeneric     NoteKind = "generic"
)

// EntityKind is entities.kind.
type EntityKind string

// Entity kinds.
const (
	EntityNPC EntityKind = "npc"
)

// Visibility is note_blocks.visibility; the rules for each are in perm.
type Visibility string

// Block visibilities.
const (
	VisGM      Visibility = "gm"
	VisPlayers Visibility = "players"
	VisListed  Visibility = "listed"
	VisOwner   Visibility = "owner"
)

// LinkTarget is the target_type of note_links and entity_field_links.
type LinkTarget string

// Link targets.
const (
	LinkNote   LinkTarget = "note"
	LinkEntity LinkTarget = "entity"
	LinkImage  LinkTarget = "image"
)
