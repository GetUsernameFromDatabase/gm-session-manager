package views

import (
	"slices"
	"strings"

	"gm-session-manager/internal/db/sqlc"
	"gm-session-manager/internal/enum"
	selectcomp "gm-session-manager/internal/web/components/select"

	"github.com/a-h/templ"
)

// NoteListView is the notes index.
type NoteListView struct {
	CampaignID int64
	Notes      []NoteItem
	Query      string
	Hits       []sqlc.SearchReadableBlocksRow
	// CanCreate is true for the GM and players.
	CanCreate bool
}

// sectionHeading is the quiet heading over a list on the notes pages.
const sectionHeading = "mb-2 text-sm font-medium text-muted-foreground"

func notesOfKind(notes []NoteItem, sessions bool) []NoteItem {
	var out []NoteItem
	for _, n := range notes {
		if (n.Kind == enum.NoteGameSession) == sessions {
			out = append(out, n)
		}
	}
	return out
}

// closeDialog closes the dialog around this element.
const closeDialog = "this.closest('[data-slot=dialog-content]').querySelector('[data-slot=dialog-close]').click()"

// closeDialogJS closes the dialog around an htmx request's element once the
// server accepts it.
var closeDialogJS = closeDialogAfter("")

// closeDialogAfter is closeDialogJS running reset first, to clear fields.
// A ComponentScript, since templ blanks a spread event handler given a string.
func closeDialogAfter(reset string) templ.ComponentScript {
	return templ.JSUnsafeFuncCall("if (event.detail.successful) { " + reset + closeDialog + " }")
}

// BlockView is one note block as a viewer sees it. The document itself is
// loaded by the editor from /blocks/{id}/doc.
type BlockView struct {
	ID             int64
	Visibility     enum.Visibility
	PlayersCanEdit bool
	Audience       []int64
	// ForYou is set for non-GM readers of a secret, who are always in its audience.
	ForYou  bool
	CanEdit bool
	// CanManage is true for the note's manager and the block's author.
	CanManage bool
	Author    Author
}

// Author is who wrote a note or block: a character's name, or "GM". Href
// links to the sheet when the viewer may open it.
type Author struct {
	Name string
	Href string
}

// NoteItem is one row of the notes index.
type NoteItem struct {
	sqlc.Note
	Author Author
	// CanManage is true for the note's manager.
	CanManage bool
}

// NoteView is a note page.
type NoteView struct {
	Note         sqlc.Note
	Blocks       []BlockView
	CanManage    bool
	Visibilities []enum.Visibility
	// Characters are the campaign's characters, for 'listed' audiences.
	Characters  []sqlc.Character
	Backlinks   []sqlc.ListBacklinksRow
	EntityLinks []sqlc.ListEntityFieldBacklinksRow
}

func (v NoteView) characterNames(ids []int64) string {
	var names []string
	for _, c := range v.Characters {
		if slices.Contains(ids, c.ID) {
			names = append(names, c.Name)
		}
	}
	if len(names) == 0 {
		return "nobody yet"
	}
	return strings.Join(names, ", ")
}

func visibilityLabel(vis enum.Visibility) string {
	switch vis {
	case enum.VisGM:
		return "GM only"
	case enum.VisPlayers:
		return "Everyone"
	case enum.VisListed:
		return "Secret"
	case enum.VisOwner:
		return "Private"
	}
	return string(vis)
}

// visRail colours a block's left rail and its access label by visibility:
// the colour is how a reader tells who else can read a passage.
var visRail = map[enum.Visibility]struct{ rail, label string }{
	enum.VisGM:      {"border-l-rose-500", "text-rose-700 dark:text-rose-300"},
	enum.VisPlayers: {"border-l-sky-500", "text-sky-700 dark:text-sky-300"},
	enum.VisListed:  {"border-l-amber-500", "text-amber-700 dark:text-amber-300"},
	enum.VisOwner:   {"border-l-violet-500", "text-violet-700 dark:text-violet-300"},
}

// buttonLike styles a label as an outline button.
const buttonLike = "inline-flex h-8 cursor-pointer items-center gap-1.5 rounded-md border bg-background px-3 text-sm font-medium shadow-xs hover:bg-accent dark:border-input dark:bg-input/30 dark:hover:bg-input/50 [&_svg]:size-4"

// visToggleJS shows each [data-for-vis] option only while its visibility is
// chosen. It runs on the visibility fields' change events.
const visToggleJS = `const vis = this.querySelector("[name=visibility]").value; for (const f of this.querySelectorAll("[data-for-vis]")) f.hidden = f.dataset.forVis !== vis`

var noteKindItems = []selectcomp.ItemData{{Value: string(enum.NoteGeneric), Label: "Note"}, {Value: string(enum.NoteGameSession), Label: "Game session"}}

// LinkItem is a note, entity or image a viewer may link to or jump to.
type LinkItem struct {
	ID    string `json:"id"` // "note:12", "entity:3" or "image:5"
	Label string `json:"label"`
	Kind  string `json:"kind"`
	URL   string `json:"url"`
}
