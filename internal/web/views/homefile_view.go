package views

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"

	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/components/dialog"
)

// What the GM chooses for a source added at home.
const (
	SourceAllow   = "allow"
	SourceDecline = "decline"
)

// homeFileReviewID is the review form, which the page swaps whole.
const homeFileReviewID = "home-file-review"

// PickParam is the query key holding the side picked for a conflicting group.
func PickParam(group string) string { return "pick_" + group }

// SourceParam is the query key holding the choice for a source.
func SourceParam(key string) string { return "source:" + key }

// HomeFileColumn is one side of a conflicting group.
type HomeFileColumn struct {
	Side, Title string
	Lines       []string
	Checked     bool
}

// HomeFileGroup is one row of the review's Changes list.
type HomeFileGroup struct {
	Key, Label string
	// Badge says which side changed ("" when nothing did).
	Badge string
	// Lines are what changed, for a group that is not in conflict.
	Lines []string
	// Note explains where the value comes from when that is not the changed side.
	Note string
	// Columns are the two sides of a conflict, the GM picks one.
	Columns []HomeFileColumn
}

// HomeFileSource is a source the file uses that the campaign does not allow:
// added at home, or (a new character) one of the table's.
type HomeFileSource struct {
	Key, Label string
	// Uses is "Used: …" or "Not used".
	Uses string
	// Effect is what declining does ("" for nothing).
	Effect string
	// Choice is SourceAllow, SourceDecline or "" (not chosen yet).
	Choice string
}

// HomeFileView is the review of one bring-back file.
type HomeFileView struct {
	CampaignID  int64
	FileID      string
	CharacterID int64
	Name        string
	TakenAt     time.Time
	// LevelFrom and LevelTo are the level at take-out and after the merge.
	LevelFrom, LevelTo int
	ByGM               bool
	Sources            []HomeFileSource
	Groups             []HomeFileGroup
	// Removed names what a declined source takes off the sheet.
	Removed []string
	// Block is why Accept is disabled ("" when it is not).
	Block string
	// Error is a merge that does not make a valid character; Stale a change at
	// the table since the page was shown. Both disable nothing but show on top.
	Error, Stale string
	// Was stamps the table character the review was made against.
	Was string
	// NewCharacter marks a character made at home that Accept adds to the
	// campaign (LevelTo is its level; there is no take-out).
	NewCharacter bool
}

// PageTitle is the review page's title.
func (v HomeFileView) PageTitle() string {
	if v.NewCharacter {
		return "Add " + v.Name
	}
	return "Bring " + v.Name + " back"
}

func (v HomeFileView) acceptLabel() string {
	if v.NewCharacter {
		return "Add to campaign"
	}
	return "Accept bring-back"
}

// URL is the review's own address, which carries the picks.
func (v HomeFileView) URL() string {
	return route.Path(route.CampaignHomeFile, v.CampaignID, v.FileID)
}

func (v HomeFileView) acceptURL() string {
	return route.Path(route.CampaignHomeFileAccept, v.CampaignID, v.FileID)
}

// RejectURL is where Reject posts.
func RejectURL(campaignID int64, fileID string) string {
	return route.Path(route.CampaignHomeFileReject, campaignID, fileID)
}

func (v HomeFileView) rejectURL() string { return RejectURL(v.CampaignID, v.FileID) }

// meta is the muted line under the name.
func (v HomeFileView) meta() string {
	if v.NewCharacter {
		s := "Made at home · level " + strconv.Itoa(v.LevelTo)
		if v.ByGM {
			return s + " · uploaded by the GM"
		}
		return s + " · sent from the player's device"
	}
	s := ""
	if v.LevelFrom != v.LevelTo {
		s += " · level " + strconv.Itoa(v.LevelFrom) + " → " + strconv.Itoa(v.LevelTo)
	}
	if v.ByGM {
		return s + " · uploaded by the GM"
	}
	return s + " · sent from the player's device"
}

// BringBackNotice is a bring-back file a player sent, waiting in the GM's popup.
type BringBackNotice struct {
	ID         string
	CampaignID int64
	Name       string
	SavedAt    time.Time
	// MadeAtHome is a new character's file, not a bring-back.
	MadeAtHome bool
}

// NewCharacterNotice is a visitor's new-character file in the GM's popup, with
// who sent it; Review opens it, Deny drops it.
type NewCharacterNotice struct {
	FileID     string
	CampaignID int64
	Name       string
	Player     string
	Device     string
	Phone      bool
	Host       string
}

// newCharacterLoad turns a button into the trigger that loads the "add from
// file" dialog body into the take-home dialog.
func newCharacterLoad(campaignID int64) templ.Attributes {
	return with(dialog.TriggerFor(takeHomeID), templ.Attributes{
		attrHxGet: route.Path(
			route.CampaignNewCharacter,
			campaignID,
		),
		attrHxTarget: takeHomeTarget,
		attrHxSwap:   swapInner,
	})
}

// BringBackView is the dialog that takes a bring-back file.
type BringBackView struct {
	CharacterID int64
	Name        string
	ByGM        bool
}

// bringBackLoad turns a button into the trigger that loads the bring-back
// dialog body into the take-home dialog.
func bringBackLoad(characterID int64) templ.Attributes {
	return with(dialog.TriggerFor(takeHomeID), templ.Attributes{
		attrHxGet: route.Path(
			route.CharacterBringBack,
			characterID,
		),
		attrHxTarget: takeHomeTarget,
		attrHxSwap:   swapInner,
	})
}

// acceptVals is the Accept button's extra field: the stamp of the table
// character the review was made against (not a form field, so the picks'
// address does not carry it).
func (v HomeFileView) acceptVals() string {
	b, _ := json.Marshal(map[string]string{"was": v.Was}) // a string map always marshals
	return string(b)
}

// checkedSide is the side a conflicting group's radio starts on.
func checkedSide(g HomeFileGroup) string {
	for _, c := range g.Columns {
		if c.Checked {
			return c.Side
		}
	}
	return ""
}

func joinNames(names []string) string { return strings.Join(names, ", ") }
