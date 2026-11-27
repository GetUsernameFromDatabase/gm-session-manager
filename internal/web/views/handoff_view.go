package views

import (
	"net/url"
	"time"

	"gm-session-manager/internal/route"
)

// The states of the phone handoff panel.
const (
	HandoffNone      = ""
	HandoffWaiting   = "waiting"
	HandoffReceived  = "received"
	HandoffSent      = "sent"
	HandoffClosed    = "closed"
	HandoffNoNetwork = "no-network"
)

// The query keys the polling panel names the handoff and state it shows in.
const (
	HandoffStateParam = "state"
	HandoffIDParam    = "id"
)

// pollURL asks whether the handoff the page shows has changed.
func (v HandoffView) pollURL() string {
	return route.AtHomeHandoff + "?" + url.Values{HandoffStateParam: {v.State}, HandoffIDParam: {v.ID}}.Encode()
}

// closesAt is the clock time the handoff closes. Formatted here, not by
// app.js: the panel is only in the player app, whose server is the
// viewer's own computer, so its local time is the viewer's.
func (v HandoffView) closesAt() string { return v.Expires.Local().Format("15:04") }

// PhoneCSS is the stylesheet's path under a phone page's own address.
const PhoneCSS = "/output.css"

// handoffRoot is the id of the panel's card, which polling swaps whole.
const handoffRoot = "handoff-panel"

// HandoffView is the phone handoff panel as the PC shows it.
type HandoffView struct {
	// ID tells handoffs apart (the secret, already in the page's addresses).
	ID    string
	State string
	// Send is a bring-back file going to the phone, else a take-home file coming from it.
	Send bool
	// Name is the character's; for a receive it is known once the file came.
	Name string
	// CharacterID is the copy being sent, for "Send again".
	CharacterID int64
	// TakenAt is when the table wrote a received file.
	TakenAt time.Time
	URLs    []string
	// QR is a data: image of the first URL.
	QR      string
	Expires time.Time
}

func (v HandoffView) title() string {
	if v.Send {
		return "Send " + v.Name + "'s bring-back file to your phone"
	}
	return "Receive from your phone"
}

// PhoneView is a page the phone opens.
type PhoneView struct {
	// Path is the page's own address; the stylesheet and the file hang off it.
	Path string
	// Send is the download page, else the upload page.
	Send     bool
	Name     string
	FileName string
	Size     int64
	// Error is why an upload was refused; Done says it was taken.
	Error string
	Done  bool
}

func (v PhoneView) sizeText() string { return fileSize(v.Size) }
