package views

import "gm-session-manager/internal/db/sqlc"

// Device is one logged-in device of a character.
type Device struct {
	SessionID int64
	UserAgent string // empty when not connected right now
	Online    bool
}

// RosterCharacter is a character with its devices.
type RosterCharacter struct {
	Character sqlc.Character
	Devices   []Device
}

// Roster is the GM's live view of the table.
type Roster struct {
	Campaign   sqlc.Campaign
	Characters []RosterCharacter
	// OtherCampaigns are copy targets.
	OtherCampaigns []sqlc.Campaign
	// Deleted are characters their player or the GM deleted; the GM may
	// restore them.
	Deleted []sqlc.Character
}

// PendingView is a join request on the GM's screen.
type PendingView struct {
	ID         string
	CampaignID int64
	Name       string
	Claim      bool
	Player     string
	Device     string
	Phone      bool
	Host       string
}

// PendingGroup is one campaign's waiting requests, oldest first.
type PendingGroup struct {
	CampaignID   int64
	CampaignName string
	Requests     []PendingView
}
