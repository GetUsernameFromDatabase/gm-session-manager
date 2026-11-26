package views

import (
	"slices"
	"time"

	"gm-session-manager/internal/db/sqlc"
)

// Device is one logged-in device of a character.
type Device struct {
	SessionID int64
	// UserAgent is what the device is (OS · browser): live while online, as it
	// last left when offline; empty for one that never left since it was saved.
	UserAgent string
	Online    bool
	// LastSeen is the device's last request or disconnect; zero while online.
	LastSeen time.Time
}

// RosterCharacter is a character with its devices.
type RosterCharacter struct {
	Character sqlc.Character
	Devices   []Device
	// TakenHome is when the newest open take-home file was written (zero: none).
	TakenHome time.Time
}

// Online and Offline split the devices: the card folds the offline ones away.
func (rc RosterCharacter) Online() []Device { return devicesWhere(rc.Devices, true) }

// Offline is the devices not connected now, which the GM may want to kick.
func (rc RosterCharacter) Offline() []Device { return devicesWhere(rc.Devices, false) }

func devicesWhere(ds []Device, online bool) []Device {
	return slices.DeleteFunc(slices.Clone(ds), func(d Device) bool { return d.Online != online })
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
