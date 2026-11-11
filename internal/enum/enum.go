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

// EncounterStatus is encounters.status: preparing → active → ended, and
// ended → active on restart.
type EncounterStatus string

// Encounter statuses.
const (
	EncounterPreparing EncounterStatus = "preparing"
	EncounterActive    EncounterStatus = "active"
	EncounterEnded     EncounterStatus = "ended"
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
