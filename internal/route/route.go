// Package route holds every route path the router registers. Handlers,
// templates and tests build URLs from these and Path, never from a literal
// (TestNoRouteLiterals). It imports only the standard library.
package route

import (
	"fmt"
	"net/url"
	"strings"
)

// Path fills the :name and *name segments of pattern with ids, in order,
// escaped for a URL path. A count that differs from the pattern's params is a
// programming error, so it panics. Same contract as routePath in wikilink.js.
func Path(pattern string, ids ...any) string {
	segs := strings.Split(pattern, "/")
	n := 0
	for i, s := range segs {
		if !strings.HasPrefix(s, ":") && !strings.HasPrefix(s, "*") {
			continue
		}
		if n < len(ids) {
			segs[i] = url.PathEscape(fmt.Sprint(ids[n]))
		}
		n++
	}
	if n != len(ids) {
		panic(fmt.Sprintf("route.Path(%q): %d ids for %d params", pattern, len(ids), n))
	}
	return strings.Join(segs, "/")
}

// Plain roots.
const (
	Home       = "/"
	AdminLogin = "/admin/login"
	AtHome     = "/at-home" // the player app's characters and import
	Campaigns  = "/campaigns"
	Compendium = "/compendium"
	Feedback   = "/feedback"
	Files      = "/files"
	Join       = "/join"
	JoinQR     = "/join.png"
	Live       = "/live"
	Logout     = "/logout"
	WS         = "/ws"

	RequestsComplete = "/requests/complete"
)

// Group roots, with their params.
const (
	AppDownload  = "/downloads/:file"
	AtHomeCopy   = AtHome + "/:character"
	Attachment   = "/attachments/:attachment"
	Block        = "/blocks/:block"
	Campaign     = "/campaigns/:campaign"
	Character    = "/characters/:character"
	Combatant    = "/combatants/:combatant"
	Condition    = "/conditions/:condition"
	Encounter    = "/encounters/:encounter"
	Entity       = "/entities/:entity"
	Image        = "/images/:image"
	JoinCampaign = Join + "/:campaign"
	Note         = "/notes/:note"
	Session      = "/sessions/:session"
)

// Path parts, one per distinct string, appended to a root.
const (
	Access             = "/access"
	Action             = "/:action"
	Attachments        = "/attachments"
	AttachmentDelete   = "/:attachment/delete"
	Blocks             = "/blocks"
	Builder            = "/builder"
	Claim              = "/claim/:character"
	Cmd                = "/cmd"
	Combat             = "/combat"
	Combatants         = "/combatants"
	Complete           = "/complete"
	Copy               = "/copy"
	Create             = "/create"
	CreateStep         = "/create/:step"
	CreateStepChange   = "/create/:step/change"
	Delete             = "/delete"
	DeleteUnused       = "/delete-unused"
	Display            = "/display"
	Doc                = "/doc"
	Download           = "/download"
	Effect             = "/effect"
	EffectSpell        = "/effect/spell"
	EffectPresetsDel   = "/effect/presets/delete"
	Encounters         = "/encounters"
	Entities           = "/entities"
	Export             = "/export"
	Fields             = "/fields"
	FieldsRemove       = "/fields/remove"
	FieldsReveal       = "/fields/reveal"
	FieldsText         = "/fields/text"
	File               = "/file"
	ImageAction        = "/images/:image/:action"
	Images             = "/images"
	ImagesImport       = "/images/import"
	ImagesImportReview = "/images/import/:token" // not named Token: gosec G101
	ImagesNew          = "/images/new"
	ImportCharacter    = "/characters/import"
	Kick               = "/kick"
	Levelup            = "/levelup"
	LevelupRevert      = "/levelup/revert"
	Link               = "/link"
	Links              = "/links"
	Move               = "/move"
	Notes              = "/notes"
	Npc                = "/npc"
	OpenJoin           = "/open-join"
	Picks              = "/picks"
	Play               = "/play"
	Remove             = "/remove"
	Rename             = "/rename"
	RequestDecision    = "/requests/:request/:decision"
	RequestsDecision   = "/requests/all/:decision"
	Restore            = "/restore"
	Reveal             = "/reveal"
	Snapshot           = "/snapshot"
	Sources            = "/sources"
	SourcesPreview     = "/sources/preview"
	Steps              = "/steps"
	Subclass           = "/subclass"
	Switcher           = "/switcher"
	TakeHome           = "/take-home"
	ToImage            = "/to-image"
)

// Compendium parts, appended to Compendium or CampaignCompendium.
const (
	CompendiumSearch              = "/search"
	CompendiumRandom              = "/random"
	CompendiumLookup              = "/lookup"
	CompendiumFind                = "/find/:kind/:slug"
	CompendiumCompare             = "/compare/:kind/:slug"
	CompendiumEntry               = "/entries/:entry"
	CompendiumEntryEdit           = CompendiumEntry + "/edit"
	CompendiumEntryDelete         = CompendiumEntry + Delete
	CompendiumEntryCampaigns      = CompendiumEntry + "/campaigns"
	CompendiumEntryCampaignDelete = CompendiumEntryCampaigns + "/:linked/delete"
	CompendiumEntryImage          = CompendiumEntry + "/image"
	CompendiumHabitats            = "/habitats"
	CompendiumHabitatCell         = "/habitats/cell"
	CompendiumHabitatAny          = "/habitats/any"
	CompendiumHabitatBatch        = "/habitats/batch"
	CompendiumHabitatFill         = "/habitats/fill"
	CompendiumHabitatCopy         = "/habitats/copy"
	CompendiumHomebrewNew         = "/homebrew/new"
	CompendiumHomebrew            = "/homebrew"
	CompendiumImport              = "/import"
	CompendiumImportDismiss       = CompendiumImport + "/dismiss"
	CompendiumFolder              = "/folder"
	CompendiumFiveTools           = "/5etools"
	CompendiumFiveToolsBookDelete = CompendiumFiveTools + "/:book/delete"
	CompendiumFiveToolsAuto       = CompendiumFiveTools + "/auto"
	CompendiumOpen5e              = "/open5e"
	// The source key travels in the form: page sources hold "/" (5etools:mtf/howler).
	CompendiumSourceDelete = "/sources/delete"
	CompendiumSourceOrder  = "/sources/order"
	CompendiumPagePreview  = "/page/preview"
	CompendiumPage         = "/page"
)

// CampaignCompendium is the compendium as one campaign sees it.
const CampaignCompendium = Campaign + Compendium

// Full patterns: the routes a caller or test names.
const (
	AtHomePlay = AtHomeCopy + Play

	// The player app's bring-back file and the phone handoff.
	AtHomeBringBack   = AtHomeCopy + "/bring-back"
	AtHomeSendPhone   = AtHomeCopy + "/send-to-phone"
	AtHomeMakeCopy    = AtHomeCopy + Copy
	AtHomeFallback    = AtHomeCopy + "/fallback"
	AtHomeNew         = AtHome + "/new"
	AtHomeReceive     = AtHome + "/receive"
	AtHomeHandoff     = AtHome + "/handoff"
	AtHomeHandoffStop = AtHomeHandoff + "/stop"

	AttachmentToImage = Attachment + ToImage

	BlockAccess      = Block + Access
	BlockAttachments = Block + Attachments
	BlockDelete      = Block + Delete
	BlockDoc         = Block + Doc
	BlockMove        = Block + Move
	BlockSnapshot    = Block + Snapshot
	BlockSteps       = Block + Steps

	CampaignCombat              = Campaign + Combat
	CampaignDelete              = Campaign + Delete
	CampaignDisplay             = Campaign + Display
	CampaignEffectSpell         = Campaign + EffectSpell
	CampaignEffectPresetsDelete = Campaign + EffectPresetsDel
	CampaignEncounters          = Campaign + Encounters
	CampaignEntities            = Campaign + Entities
	CampaignExport              = Campaign + Export
	CampaignFiles               = Campaign + Files
	CampaignImages              = Campaign + Images
	CampaignImagesImport        = Campaign + ImagesImport
	CampaignImagesImportReview  = Campaign + ImagesImportReview
	CampaignImagesNew           = Campaign + ImagesNew
	CampaignImportCharacter     = Campaign + ImportCharacter
	CampaignLinks               = Campaign + Links
	CampaignNotes               = Campaign + Notes
	CampaignOpenJoin            = Campaign + OpenJoin
	CampaignRename              = Campaign + Rename
	CampaignRequestDecision     = Campaign + RequestDecision
	CampaignRequestsDecision    = Campaign + RequestsDecision
	CampaignSources             = Campaign + Sources
	CampaignSwitcher            = Campaign + Switcher

	CharacterBuilder        = Character + Builder
	CharacterCmd            = Character + Cmd
	CharacterComplete       = Character + Complete
	CharacterCopy           = Character + Copy
	CharacterCreate         = Character + Create
	CharacterCreateStep     = Character + CreateStep
	CharacterCreateChange   = Character + CreateStepChange
	CharacterDelete         = Character + Delete
	CharacterDownload       = Character + Download
	CharacterEffect         = Character + Effect
	CharacterImages         = Character + Images
	CharacterImageAction    = Character + ImageAction
	CharacterLevelup        = Character + Levelup
	CharacterLevelupRevert  = Character + LevelupRevert
	CharacterPicks          = Character + Picks
	CharacterRemove         = Character + Remove
	CharacterRestore        = Character + Restore
	CharacterSourcesPreview = Character + SourcesPreview
	CharacterSubclass       = Character + Subclass
	CharacterTakeHome       = Character + TakeHome

	CombatantAction = Combatant + Action

	ConditionDelete = Condition + Delete

	EncounterAction     = Encounter + Action
	EncounterCombatants = Encounter + Combatants

	EntityDelete       = Entity + Delete
	EntityFields       = Entity + Fields
	EntityFieldsRemove = Entity + FieldsRemove
	EntityFieldsReveal = Entity + FieldsReveal
	EntityFieldsText   = Entity + FieldsText
	EntityImages       = Entity + Images
	EntityImageAction  = Entity + ImageAction
	EntityLink         = Entity + Link
	EntityRename       = Entity + Rename
	EntityReveal       = Entity + Reveal

	FilesAttachmentDelete         = Files + AttachmentDelete
	FilesDeleteUnused             = Files + DeleteUnused
	CampaignFilesAttachmentDelete = CampaignFiles + AttachmentDelete
	CampaignFilesDeleteUnused     = CampaignFiles + DeleteUnused

	ImageDelete = Image + Delete
	ImageFile   = Image + File
	ImageNpc    = Image + Npc

	JoinClaim  = JoinCampaign + Claim
	JoinCreate = JoinCampaign + Create

	NoteBlocks = Note + Blocks
	NoteDelete = Note + Delete
	NoteRename = Note + Rename

	SessionKick = Session + Kick
)
