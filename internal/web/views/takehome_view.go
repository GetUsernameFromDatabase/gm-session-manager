package views

import (
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"

	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/components/dialog"
)

// takeHomeID is the take-home dialog, one per page.
const takeHomeID = "take-home-dialog"

// takeHomeTarget is where the dialog's body loads into.
const takeHomeTarget = "#take-home"

// takeHomeLoad turns a button into the dialog's trigger, loading its body.
func takeHomeLoad(characterID int64) templ.Attributes {
	return with(dialog.TriggerFor(takeHomeID), templ.Attributes{
		attrHxGet: route.Path(route.CharacterTakeHome, characterID), attrHxTarget: takeHomeTarget, attrHxSwap: swapInner,
	})
}

// TakeHomeView is the dialog body: the app files on offer and who asks.
type TakeHomeView struct {
	CharacterID int64
	Name        string
	// Downloads are the file names in the downloads folder.
	Downloads []string
	// ByGM is true when the GM takes a player's character: the dialog asks them
	// to pass the file, private notes included, on unopened.
	ByGM bool
}

func (v TakeHomeView) noApp() string {
	if v.ByGM {
		return "Put the app files in DATA_DIR/downloads to offer them here."
	}
	return "Ask your GM for the app."
}

// importConfirmID is the container of the replace question: the import form
// and the handoff panel's each have one.
func importConfirmID(fromPhone bool) string {
	if fromPhone {
		return "handoff-import-confirm"
	}
	return "home-import-confirm"
}

// HomeCopyView is one character on this computer (the player app).
type HomeCopyView struct {
	CharacterID int64
	// CampaignID is the home campaign the character lives in.
	CampaignID int64
	Name       string
	Campaign   string
	// FromTable is true for a table character (it has a home copy); false for a copy or new character.
	FromTable bool
	// TakenAt is when the table wrote the file (zero when not FromTable).
	TakenAt time.Time
	// Current is the character this browser plays now.
	Current bool
	// Fallback is the prepared fallback build of a table character, nil for none.
	Fallback *HomeFallbackView
	// HomeSources are the sources added here that the build uses, for a table
	// character with no fallback yet: a fallback is offered when there are any.
	HomeSources []string
}

// newCharacterID is the New character dialog.
const newCharacterID = "new-character-dialog"

func (c HomeCopyView) label() string { return c.Name + " · " + c.Campaign }

// fallbackWarning says why a fallback is worth preparing.
func (c HomeCopyView) fallbackWarning() string {
	it := "it"
	if len(c.HomeSources) > 1 {
		it = "them"
	}
	return c.Name + " uses " + sourceList(c.HomeSources) + ", added here. If the GM declines " + it + ", " + c.Name + " can't come back with this level."
}

// HomeFallbackView is a table character's fallback build.
type HomeFallbackView struct {
	CharacterID int64
	Name        string
	Current     bool
}

// sourceList names sources in a sentence: "A", "A and B", "A, B and C".
func sourceList(keys []string) string {
	if len(keys) < 2 {
		return strings.Join(keys, "")
	}
	return strings.Join(keys[:len(keys)-1], ", ") + " and " + keys[len(keys)-1]
}

// HomeConfirmView is the question asked before a copy here is replaced.
type HomeConfirmView struct {
	Name     string
	Old, New time.Time
	// CopyHere is true when a copy of the same table character is here (Old is
	// when it was taken home).
	CopyHere bool
	// MadeAtHome names the character made here that this file replaces ("" for none).
	MadeAtHome string
	// Sources are those added here that the copy allows and the new file does
	// not list: the new copy leaves them off.
	Sources []string
	// Fallback is the name of the copy's fallback build, which goes too ("" for none).
	Fallback string
}

// title says what the question is about.
func (v HomeConfirmView) title() string {
	if v.CopyHere {
		return v.Name + " is already on this computer"
	}
	return v.Name + " replaces a character made here"
}

// buttonLabel names what Replace removes.
func (v HomeConfirmView) buttonLabel() string {
	if v.CopyHere {
		return "Replace " + v.Name
	}
	return "Replace " + v.MadeAtHome
}

// message explains what replacing does, by which file is newer.
func (v HomeConfirmView) message() string {
	const day = "2 Jan"
	var msg string
	if v.CopyHere {
		msg = "The copy here was taken home " + v.Old.Local().Format(day) + "; this file is from " + v.New.Local().Format(day) + ". Replacing it drops the changes made here since " + v.Old.Local().Format(day) + "."
		if v.New.Before(v.Old) {
			msg = "This file (" + v.New.Local().Format(day) + ") is older than the copy here (" + v.Old.Local().Format(day) + "). Replacing it drops the changes made here."
		}
	}
	if v.MadeAtHome != "" {
		msg += " Replaces " + v.MadeAtHome + ", made at home. Changes made at home since you sent them are lost."
	}
	msg = strings.TrimSpace(msg)
	if len(v.Sources) > 0 {
		msg += " The new copy also leaves off the sources added here (" + strings.Join(v.Sources, ", ") + "): add them again from its compendium's Import tab."
	}
	if v.Fallback != "" {
		msg += " Its fallback build (" + v.Fallback + ") goes too."
	}
	return msg
}

// HomeStarterView is a campaign starter on this computer.
type HomeStarterView struct {
	CampaignID int64
	Name       string
	// Books are the labels of its books, in priority order.
	Books   []string
	SavedAt time.Time
}

func (s HomeStarterView) removeConfirm() string {
	return "Remove the " + s.Name + " starter? Characters made from it stay."
}

func (s HomeStarterView) dialogID() string {
	return newCharacterID + "-" + strconv.FormatInt(s.CampaignID, 10)
}

func (s HomeStarterView) options() []homeNewOption {
	return []homeNewOption{{CampaignID: s.CampaignID, Label: s.Name + " (starter)"}}
}

// homeNewOption is a campaign a new character can start in.
type homeNewOption struct {
	CampaignID int64
	Label      string
}

func (o homeNewOption) value() string { return strconv.FormatInt(o.CampaignID, 10) }
func (o homeNewOption) label() string { return o.Label }

// newCharacterOptions lists the campaigns a new character can start in: each
// table character's (as before), then each starter.
func newCharacterOptions(table []HomeCopyView, starters []HomeStarterView) []homeNewOption {
	var opts []homeNewOption
	for _, c := range table {
		opts = append(opts, homeNewOption{CampaignID: c.CampaignID, Label: c.label()})
	}
	for _, s := range starters {
		opts = append(opts, s.options()...)
	}
	return opts
}
