package views

import (
	"fmt"
	"slices"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/web/components/combobox"
	selectcomp "gm-session-manager/internal/web/components/select"

	"github.com/a-h/templ"
)

// ImageItem is one image as a page shows it. Players get only ID, Name
// and Src; the handlers leave the rest empty for them.
type ImageItem struct {
	ID          int64
	Name        string
	Category    string
	Description string
	// URL is the external link (empty for uploads), for the edit form.
	URL string
	// Version names the link or file, so Src changes with it and browsers
	// may cache it for good.
	Version string
}

// Src is where browsers load the image: the authorized handler, which
// serves an upload or redirects to the external link.
func (i ImageItem) Src() string {
	if i.Version == "" {
		return ImageSrc(i.ID)
	}
	return ImageSrc(i.ID) + "?v=" + i.Version
}

// ImageSrc is the authorized URL of image id.
func ImageSrc(id int64) string { return fmt.Sprintf("/images/%d/file", id) }

// DisplayURL is the display window of a campaign.
func DisplayURL(campaignID int64) string { return fmt.Sprintf("/campaigns/%d/display", campaignID) }

// JoinQRSrc is the join QR code image for a campaign, or for the campaign
// list when campaignID is 0. The query name matches web's queryCampaign.
func JoinQRSrc(campaignID int64) string {
	if campaignID == 0 {
		return "/join.png"
	}
	return fmt.Sprintf("/join.png?campaign=%d", campaignID)
}

// ImagesView is the GM's image picker.
type ImagesView struct {
	CampaignID int64
	Query      string
	CategoryID int64
	Categories []sqlc.ImageCategory
	Images     []ImageItem
	// Imported is how many images an import just added (shown once).
	Imported int
}

// ImageResultsID is the element htmx swaps when searching or filtering.
const ImageResultsID = "image-results"

func categoryFilterItems(cats []sqlc.ImageCategory) []selectcomp.ItemData {
	items := []selectcomp.ItemData{{Value: "0", Label: "All categories"}}
	return append(items, selectItems(cats, func(c sqlc.ImageCategory) string { return fmt.Sprint(c.ID) }, func(c sqlc.ImageCategory) string { return c.Name })...)
}

// showPressed marks a show button while its image is on the display.
const showPressed = "aria-pressed:bg-primary aria-pressed:text-primary-foreground"

// showAttrs make a button put src on the display window (app.js).
func showAttrs(src, name string) templ.Attributes {
	return templ.Attributes{"data-show-image": "", "data-src": src, "data-name": name, attrAriaPressed: ariaFalse}
}

// ImageView is one image page: the GM's editor or a player's view.
type ImageView struct {
	CampaignID int64
	Image      ImageItem
	IsGM       bool
	Categories []sqlc.ImageCategory
	Backlinks  []sqlc.ListBacklinksRow
	// EntityLinks are the NPCs whose fields link the image.
	EntityLinks []sqlc.ListEntityFieldBacklinksRow
	// NPCs or Characters wear the image as a look (never both); the choices
	// are who else it may be given to.
	NPCs             []sqlc.ListImageEntitiesRow
	NPCChoices       []sqlc.Entity
	Characters       []sqlc.ListImageCharactersRow
	CharacterChoices []sqlc.Character
}

func imageAction(v ImageView) string {
	if v.Image.ID == 0 {
		return fmt.Sprintf("/campaigns/%d/images", v.CampaignID)
	}
	return fmt.Sprintf("/images/%d", v.Image.ID)
}

// comboEmptyHidden hides a combobox popup with nothing to show (no items, no
// create button), which would otherwise be a bare ring around the input.
// invisible, not hidden, so its close transition still ends.
const comboEmptyHidden = "data-empty:has-[[data-combobox-create][hidden]]:invisible"

// sourceField is a choice between a link and an upload, as tabs; app.js
// submits only the active tab's input.
type sourceField struct {
	ID, Label, URL, Placeholder, Accept, FileHint string
	// Upload opens the upload tab.
	Upload bool
}

// comboItems are a combobox's labels for names, plus value when it has no
// item of its own (a custom species, say) so the input still shows it.
func comboItems(value string, names []string) []combobox.ItemData {
	items := make([]combobox.ItemData, 0, len(names)+1)
	for _, n := range names {
		items = append(items, combobox.ItemData{Value: n, Label: n})
	}
	if value != "" && !slices.Contains(names, value) {
		items = append(items, combobox.ItemData{Value: value, Label: value})
	}
	return items
}
