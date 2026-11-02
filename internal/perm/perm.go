// Package perm holds every permission and visibility rule as pure functions,
// so handlers, the websocket fan-out and tests all ask the same question the
// same way. Nothing here touches the database.
package perm

import "slices"

// Viewer is who is looking: the GM (admin) or the player of one character.
// The zero value is an anonymous visitor and may read nothing.
type Viewer struct {
	IsGM        bool
	CharacterID int64 // 0 unless a player
}

// GM is the viewer for admin requests.
var GM = Viewer{IsGM: true}

// Player is the viewer for a player session bound to characterID.
func Player(characterID int64) Viewer { return Viewer{CharacterID: characterID} }

// IsPlayer reports whether the viewer is bound to a character.
func (v Viewer) IsPlayer() bool { return !v.IsGM && v.CharacterID != 0 }

func (v Viewer) owns(characterID int64) bool {
	return v.IsPlayer() && characterID != 0 && v.CharacterID == characterID
}

// Visibility is note_blocks.visibility.
type Visibility string

// Block visibilities.
const (
	VisGM      Visibility = "gm"
	VisPlayers Visibility = "players"
	VisListed  Visibility = "listed"
	VisOwner   Visibility = "owner"
)

// Block is what the block rules need to know about a note block.
type Block struct {
	Visibility     Visibility
	PlayersCanEdit bool
	// NoteOwnerID is the note's owner_character_id (0 for the GM's notes).
	NoteOwnerID int64
	// Audience lists the characters of a 'listed' block.
	Audience []int64
	// AuthorID is the character who wrote the block (0 for the GM).
	AuthorID int64
}

// CanReadBlock is the read rule for note blocks. It must agree with the
// readable-block predicate in db/queries/notes.sql.
func CanReadBlock(v Viewer, b Block) bool {
	switch b.Visibility {
	case VisGM:
		return v.IsGM
	case VisPlayers:
		return v.IsGM || v.IsPlayer()
	case VisListed:
		// An empty audience means GM only: fail closed.
		return v.IsGM || (v.IsPlayer() && slices.Contains(b.Audience, v.CharacterID))
	case VisOwner:
		// The GM and admin can never read a player's own blocks.
		return v.owns(b.NoteOwnerID)
	}
	return false
}

// CanEditBlock is the write rule for note blocks (submitting steps).
// The author may always edit a block they can still read.
func CanEditBlock(v Viewer, b Block) bool {
	if v.owns(b.AuthorID) && CanReadBlock(v, b) {
		return true
	}
	switch b.Visibility {
	case VisGM, VisListed:
		return v.IsGM
	case VisPlayers:
		return v.IsGM || (v.IsPlayer() && b.PlayersCanEdit) || v.owns(b.NoteOwnerID)
	case VisOwner:
		return v.owns(b.NoteOwnerID)
	}
	return false
}

// CanManageBlock reports whether the viewer may delete or re-scope a block:
// the note's manager, or the block's author while they can still read it.
func CanManageBlock(v Viewer, b Block) bool {
	return CanManageNote(v, b.NoteOwnerID) || (v.owns(b.AuthorID) && CanReadBlock(v, b))
}

// CanManageNote reports whether the viewer may rename or delete a note and
// add, move, delete or re-scope its blocks: the GM for the GM's notes, the
// owning player for a player's note.
func CanManageNote(v Viewer, noteOwnerID int64) bool {
	if noteOwnerID == 0 {
		return v.IsGM
	}
	return v.owns(noteOwnerID)
}

// BlockVisibilities lists the visibilities the viewer may give a block in a
// note: any in a note they manage, and 'players' for a player writing in the
// GM's game session note. Players write private blocks or blocks for the
// table; only the GM keeps secrets from players.
func BlockVisibilities(v Viewer, noteOwnerID int64, gameSession bool) []Visibility {
	switch {
	case !CanManageNote(v, noteOwnerID):
		if v.IsPlayer() && noteOwnerID == 0 && gameSession {
			return []Visibility{VisPlayers}
		}
		return nil
	case noteOwnerID == 0:
		return []Visibility{VisGM, VisPlayers, VisListed}
	default:
		return []Visibility{VisOwner, VisPlayers}
	}
}

// ValidBlockVisibility reports whether vis is one the viewer may set.
func ValidBlockVisibility(v Viewer, noteOwnerID int64, gameSession bool, vis Visibility) bool {
	return slices.Contains(BlockVisibilities(v, noteOwnerID, gameSession), vis)
}

// CharacterScope groups character commands by who may run them.
type CharacterScope int

// Character scopes, from least to most privileged.
const (
	ScopeView CharacterScope = iota
	ScopeHP
	ScopeConditions
	ScopeSheet
)

// CanCharacter applies the permissions table: a player has full control of
// their own character and nothing else; the GM views every character and
// edits only HP and conditions.
func CanCharacter(v Viewer, characterID int64, scope CharacterScope) bool {
	if v.owns(characterID) {
		return true
	}
	return v.IsGM && scope <= ScopeConditions
}

// CanDeleteCharacter reports whether v may delete a character: its player
// or the GM. Deleting only hides it; restoring and deleting it forever are
// the GM's alone.
func CanDeleteCharacter(v Viewer, characterID int64) bool {
	return v.IsGM || v.owns(characterID)
}

// CanEditHomebrew reports whether v may change or delete a homebrew entry:
// the GM any, a player only the entries their character wrote.
func CanEditHomebrew(v Viewer, ownerCharacterID int64) bool {
	return v.IsGM || v.owns(ownerCharacterID)
}

// CanSeeEntry reports whether a compendium entry shows in a campaign:
// imported and GM entries everywhere, player homebrew (scoped) only in the
// campaigns it is linked to. The GM's homebrew list shows them all.
func CanSeeEntry(campaignID int64, scoped bool, linked []int64) bool {
	return !scoped || slices.Contains(linked, campaignID)
}

// Combatant is what the HP rule needs to know about a combatant.
type Combatant struct {
	CharacterID int64 // 0 for monsters and NPCs
	HPVisible   bool
}

// CanSeeCombatantHP reports whether v sees a combatant's HP: the GM always,
// a player for their own character, or when the GM toggled hp_visible.
func CanSeeCombatantHP(v Viewer, c Combatant) bool {
	return v.IsGM || v.owns(c.CharacterID) || (v.IsPlayer() && c.HPVisible)
}

// CanReadAttachment reports whether v may download a file. A file attached
// to a block is readable exactly when the block is; a file with no block is
// the GM's.
func CanReadAttachment(v Viewer, block *Block) bool {
	if block == nil {
		return v.IsGM
	}
	return CanReadBlock(v, *block)
}

// ImageLook is what image access needs to know about an image worn as a look.
type ImageLook struct {
	Revealed bool  // a revealed look of a revealed NPC, or of a character
	Owner    int64 // the character wearing it, 0 if none
}

// CanReadImage reports whether v may see an image. Images are the GM's; a
// player sees one through a block they can read that links it, as a
// revealed look, or as a look of their own character.
func CanReadImage(v Viewer, linkingBlocks []Block, look ImageLook) bool {
	return v.IsGM || (v.IsPlayer() && look.Revealed) || v.owns(look.Owner) ||
		slices.ContainsFunc(linkingBlocks, func(b Block) bool { return CanReadBlock(v, b) })
}
