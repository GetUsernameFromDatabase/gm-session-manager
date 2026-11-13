package views

import (
	"context"
	"io"

	"gm-session-manager/internal/compendium"
	"gm-session-manager/internal/db/sqlc"

	"github.com/a-h/templ"
)

// HTML attribute names that Go helpers put in templ.Attributes.
const (
	attrAriaLabel       = "aria-label"
	attrAriaPressed     = "aria-pressed"
	attrAriaDescribedby = "aria-describedby"
	attrAriaInvalid     = "aria-invalid"
	attrInputmode       = "inputmode"
	attrMin             = "min"
	attrHxGet           = "hx-get"
	attrHxTarget        = "hx-target"
	attrHxSwap          = "hx-swap"
)

// inputmodeNumeric brings up the number pad.
const inputmodeNumeric = "numeric"

// swapInner is the hx-swap that replaces the target's contents.
const swapInner = "innerHTML"

// Values of boolean-like ARIA attributes.
const (
	ariaTrue  = "true"
	ariaFalse = "false"
)

// Link targets and labels that more than one view repeats.
const (
	targetBlank = "_blank"
	labelNone   = "None"
	labelOther  = "Other"
	labelName   = "Name"
	valNone     = "none"
)

// Nav keys for Layout.Active.
const (
	NavHome       = "home"
	NavCharacter  = "character"
	NavCombat     = "combat"
	NavNotes      = "notes"
	NavEntities   = "entities"
	NavCompendium = "compendium"
	NavImages     = "images"
	NavFiles      = "files"
)

// Layout is the frame around every page.
type Layout struct {
	Title string
	// IsGM shows the GM navigation.
	IsGM bool
	// Campaign is the campaign the page belongs to (0 for none).
	CampaignID   int64
	CampaignName string
	// CharacterID is the player's character, for the player navigation.
	CharacterID int64
	Active      string
	// HasSession shows the logout button.
	HasSession bool
	// WS connects the page to the hub; nil for pages without live updates.
	WS *WS
	// Modules are extra ES module scripts (the notes editor).
	Modules []string
	// Requests are the join requests waiting for the GM, in every campaign.
	Requests []PendingGroup
	// Switch is the GM's campaign switcher (compendium, files); nil elsewhere.
	Switch *CampaignSwitch
	// StaleBooks are the GM's 5etools books to download again (banner);
	// AutoRedownload is the banner's switch.
	StaleBooks     []compendium.Book
	AutoRedownload bool
}

// CampaignSwitch opens the same page in another campaign, or in all of them.
type CampaignSwitch struct {
	Campaigns []sqlc.Campaign
	// Base is the page's base address in a campaign (0: all of them).
	Base func(campaign int64) string
	// All names the page of every campaign.
	All string
	// Here is the page's path after its base, with its query.
	Here string
}

// WS is the websocket endpoint and the hub position the page was read at.
type WS struct {
	URL   string
	Epoch string
	Seq   uint64
}

func (l Layout) title() string {
	if l.Title == "" {
		return "GM Session Manager"
	}
	return l.Title + " · GM Session Manager"
}

// Page renders its body inside l's campaign, so compendium links stay in it.
func Page(l Layout) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if l.CampaignID != 0 {
			ctx = WithCampaign(ctx, l.CampaignID)
		}
		return page(l).Render(ctx, w)
	})
}
