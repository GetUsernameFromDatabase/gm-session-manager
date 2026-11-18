package web

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"gm-session-manager/internal/enum"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/route"
	"gm-session-manager/internal/web/views"
)

// sharedExports are the Go values browser JS needs too, served as the ES
// module /assets/js/shared.js. Adding an export is one line here.
var sharedExports = []struct {
	name  string
	value any
}{
	{"SAFE_SCHEMES", safeSchemes},
	{"MAX_STEPS", maxSteps},
	{"MAX_STEP_BYTES", maxStepBytes},
	{"CONSOLE_SIZE", consoleSize},
	{"CONSOLE_MSG", consoleMsg},
	{"CONSOLE_PAGE", consolePage},
	{"CONSOLE_LEVELS", consoleLevels},
	{"ROUTES", sharedRoutes},
	{"WIKILINK", views.WikiLinkPattern},
	{"MSG", sharedMsg},
	{"KEY", sharedKeys},
}

// sharedMsg are the websocket message types, as {value: value}.
var sharedMsg = map[string]string{
	hub.MsgHello:  hub.MsgHello,
	hub.MsgSub:    hub.MsgSub,
	hub.MsgUnsub:  hub.MsgUnsub,
	hub.MsgPatch:  hub.MsgPatch,
	hub.MsgReload: hub.MsgReload,
	hub.MsgGone:   hub.MsgGone,
	msgSteps:      msgSteps,
}

// sharedKeys (name = key) are the JSON keys notes.js reads from the /doc and attachment
// replies; struct-tag keys (steps payload) stay plain in the JS.
var sharedKeys = map[string]string{
	keyDoc:        keyDoc,
	keyDocVersion: keyDocVersion,
	keySteps:      keySteps,
	keyClientIDs:  keyClientIDs,
	keyEditable:   keyEditable,
	keyID:         keyID,
	keyName:       keyName,
	keyURL:        keyURL,
	keyImage:      keyImage,
	keyError:      keyError,
}

// sharedRoutes are the route patterns browser JS fills with ids (routePath in
// wikilink.js); a link kind's page is under its kind (LINK_PATHS).
var sharedRoutes = map[string]string{
	"attachment":            route.Attachment,
	"attachmentToImage":     route.AttachmentToImage,
	"blockAttachments":      route.BlockAttachments,
	"blockDoc":              route.BlockDoc,
	"blockSnapshot":         route.BlockSnapshot,
	"blockSteps":            route.BlockSteps,
	"compendium":            route.Compendium,
	string(enum.LinkEntity): route.Entity,
	string(enum.LinkImage):  route.Image,
	"imageFile":             route.ImageFile,
	"links":                 route.CampaignLinks,
	"live":                  route.Live,
	string(enum.LinkNote):   route.Note,
}

// sharedJS is built once at start-up; sharedJSETag validates it like assetETags does.
var sharedJS, sharedJSETag = buildSharedJS()

func buildSharedJS() ([]byte, string) {
	var b []byte
	for _, e := range sharedExports {
		v, err := json.Marshal(e.value)
		if err != nil {
			panic(err) // plain strings, numbers and slices always encode
		}
		b = fmt.Appendf(b, "export const %s = %s;\n", e.name, v)
	}
	h := sha256.Sum256(b)
	return b, fmt.Sprintf(`"%x"`, h[:8])
}
