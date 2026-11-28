// Package web holds the HTTP handlers, templ components and static assets.
package web

import (
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"gm-session-manager/internal/auth"
	"gm-session-manager/internal/db"
	"gm-session-manager/internal/hub"
	"gm-session-manager/internal/route"
)

// UnstampedVersion is main.version in a plain go build (task run); the GM footer shows the commit
// Go recorded instead.
const UnstampedVersion = "dev"

// Deps is what the web layer needs from main.
type Deps struct {
	Auth  *auth.Manager
	Store *db.Store
	// DataDir holds attachments/ and compendium/.
	DataDir string
	Port    int
	// Player is the player app (--player); Version is this binary's version.
	Player  bool
	Version string
	// LANIPs and Fingerprint are shown on the GM's join panel.
	LANIPs      []net.IP
	Fingerprint string
}

// App is the web layer: the router and the websocket hub it renders for.
type App struct {
	Deps
	Hub       *hub.Hub
	imports   csvImports
	homeFiles homeFiles
	reqs      reqLog
	fiveTools fiveToolsJob
	state     appState
	handoffMu sync.Mutex
	handoff   *handoff
	// handoffIPsFn replaces the network detection in tests.
	handoffIPsFn func() []net.IP
}

// New builds the app. The hub renders fragments through the app.
func New(d Deps) *App {
	a := &App{Deps: d}
	a.Hub = hub.New(a)
	return a
}

// Close disconnects all websockets.
func (a *App) Close() {
	a.stopHandoff()
	a.Hub.Close()
}

// Router builds the HTTP handler with session middleware on every route.
func (a *App) Router() *gin.Engine {
	useValidateTag()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// Never derive the client address from forwarding headers. Admin
	// detection reads RemoteAddr directly regardless; this keeps ClientIP()
	// honest too.
	if err := r.SetTrustedProxies(nil); err != nil {
		panic(err) // nil is always a valid proxy list
	}
	r.Use(securityHeaders, sameOriginOnly(http.NewCrossOriginProtection()))
	a.mountAssets(r)
	r.Use(a.Auth.Middleware(), a.recordRequest, collectHeld)
	if a.Player {
		r.Use(a.homeSources)
	}

	gm := auth.RequireAdmin(a.gmOnly)

	att := r.Group(route.Attachment)
	block := r.Group(route.Block)
	camp := r.Group(route.Campaign)
	campComp := camp.Group(route.Compendium, a.compendiumCampaign)
	ch := r.Group(route.Character)
	comp := r.Group(route.Compendium, a.compendiumWhole)
	ent := r.Group(route.Entity)
	img := r.Group(route.Image)
	note := r.Group(route.Note)

	// Within each section: sorted by full path, then GET before POST.

	// Joining.
	// The player app has no admin login and no table to join.
	r.GET(route.Home, a.h(a.home))
	r.GET(route.AdminLogin, a.notPlayerApp, a.adminLoginForm)
	r.POST(route.AdminLogin, a.notPlayerApp, a.h(a.adminLogin))
	r.GET(route.AppDownload, a.h(a.appDownload))
	r.GET(route.AtHome, a.h(a.homePage))
	r.POST(route.AtHome, a.h(a.homeImport))
	r.POST(route.AtHomePlay, a.onlyPlayerApp, a.h(a.homePlay))
	r.POST(route.AtHomeMakeCopy, a.onlyPlayerApp, a.h(a.homeMakeCopy))
	r.POST(route.AtHomeFallback, a.onlyPlayerApp, a.h(a.homeFallback))
	r.POST(route.AtHomeNew, a.onlyPlayerApp, a.h(a.homeNew))
	r.POST(route.AtHomeStarterRemove, a.onlyPlayerApp, a.h(a.homeStarterRemove))
	r.POST(route.AtHomeBringBack, a.onlyPlayerApp, a.h(a.homeBringBackCheck))
	r.GET(route.AtHomeBringBack, a.onlyPlayerApp, a.h(a.homeBringBack))
	r.POST(route.AtHomeSendPhone, a.onlyPlayerApp, a.h(a.handoffSend))
	r.POST(route.AtHomeReceive, a.onlyPlayerApp, a.h(a.handoffReceive))
	r.GET(route.AtHomeHandoff, a.onlyPlayerApp, a.h(a.handoffPanel))
	r.POST(route.AtHomeHandoffStop, a.onlyPlayerApp, a.h(a.handoffStopRoute))
	r.POST(route.Feedback, a.h(a.sendFeedback))
	r.GET(route.Join, a.notPlayerApp, a.h(a.joinCampaigns))
	r.GET(route.JoinQR, a.notPlayerApp, a.h(a.joinQR))
	r.GET(route.JoinCampaign, a.notPlayerApp, a.joinCharacters)
	r.POST(route.JoinClaim, a.notPlayerApp, a.h(a.joinClaim))
	r.POST(route.JoinCreate, a.notPlayerApp, a.h(a.joinCreate))
	r.POST(route.JoinNewCharacter, a.notPlayerApp, a.h(a.joinNewCharacter))
	r.GET(route.JoinStarter, a.notPlayerApp, a.h(a.joinStarter))
	r.GET(route.Live, a.h(a.liveHTML))
	r.POST(route.Logout, a.h(a.logout))
	r.POST(route.RequestsComplete, a.h(a.completeRequest))
	r.GET(route.WS, a.websocket)

	// Campaigns.
	r.POST(route.Campaigns, gm, a.h(a.createCampaign))
	camp.GET("", gm, a.h(a.campaignOverview))
	camp.POST(route.Delete, gm, a.h(a.deleteCampaign))
	camp.GET(route.HomeFile, gm, a.h(a.homeFilePage))
	camp.GET(route.NewCharacter, gm, a.h(a.newCharacterDialog))
	camp.POST(route.NewCharacter, gm, a.h(a.newCharacterUpload))
	camp.POST(route.HomeFile+route.Accept, gm, a.h(a.homeFileAccept))
	camp.POST(route.HomeFile+route.Reject, gm, a.h(a.homeFileReject))
	camp.GET(route.EffectSpell, a.h(a.effectSpell))
	camp.POST(route.EffectPresetsDel, a.h(a.deleteEffectPreset))
	camp.GET(route.Export, gm, a.h(a.exportCampaign))
	camp.POST(route.OpenJoin, gm, a.h(a.toggleOpenJoin))
	camp.POST(route.Rename, gm, a.h(a.renameCampaign))
	camp.POST(route.RequestDecision, gm, a.decideRequest)
	camp.POST(route.RequestsDecision, gm, a.decideAll)
	camp.POST(route.Sources, a.sourcesAdmin(gm), a.h(a.campaignSourcesSet))
	camp.GET(route.Starter, gm, a.h(a.campaignStarter))
	camp.GET(route.Switcher, a.h(a.switcher))
	r.POST(route.SessionKick, gm, a.h(a.kick))

	// Characters.
	ch.GET("", a.h(a.characterPage))
	ch.GET(route.BringBack, a.h(a.bringBackDialog))
	ch.POST(route.BringBack, a.h(a.bringBackUpload))
	ch.GET(route.Builder, a.h(a.commandBuilder))
	ch.POST(route.Cmd, a.h(a.characterCommand))
	ch.GET(route.Complete, a.h(a.commandCompletions))
	ch.POST(route.Copy, gm, a.h(a.copyCharacter))
	ch.GET(route.Create, a.wizardStart)
	ch.GET(route.CreateStep, a.h(a.wizardPage))
	ch.POST(route.CreateStep, a.h(a.wizardSave))
	ch.POST(route.CreateStepChange, a.h(a.wizardChange))
	ch.POST(route.Delete, gm, a.h(a.deleteCharacter))
	ch.POST(route.Effect, a.h(a.characterEffect))
	ch.GET(route.Download, gm, a.h(a.downloadCharacter))
	ch.POST(route.Images, a.h(a.addCharacterLook))
	ch.POST(route.ImageAction, a.h(a.characterLook))
	ch.GET(route.Levelup, a.h(a.levelUpForm))
	ch.POST(route.Levelup, a.h(a.levelUp))
	ch.POST(route.LevelupRevert, a.h(a.revertLevelUp))
	ch.GET(route.Picks, a.h(a.builderPicks))
	ch.POST(route.Remove, a.h(a.removeCharacter))
	ch.POST(route.Restore, gm, a.h(a.restoreCharacter))
	ch.GET(route.SourcesPreview, a.h(a.sourcePreview))
	ch.GET(route.Subclass, a.h(a.subclassChangeForm))
	ch.POST(route.Subclass, a.h(a.changeSubclass))
	ch.GET(route.TakeHome, a.h(a.takeHomeDialog))
	ch.POST(route.TakeHome, a.h(a.takeHome))
	r.POST(route.ConditionDelete, a.h(a.deleteCondition))

	// Combat.
	camp.GET(route.Combat, a.h(a.combatPage))
	camp.POST(route.Encounters, gm, a.h(a.createEncounter))
	r.POST(route.CombatantAction, gm, a.h(a.combatantAction))
	r.POST(route.EncounterAction, gm, a.h(a.encounterAction))
	r.POST(route.EncounterCombatants, gm, a.h(a.addCombatant))

	// Notes.
	att.GET("", a.attachment)
	block.POST(route.Access, a.h(a.blockAccess))
	block.POST(route.Attachments, a.uploadAttachment)
	block.POST(route.Delete, a.h(a.deleteBlock))
	block.GET(route.Doc, a.h(a.blockDoc))
	block.POST(route.Move, a.h(a.moveBlock))
	block.POST(route.Snapshot, a.h(a.blockSnapshot))
	block.GET(route.Steps, a.h(a.blockSteps))
	block.POST(route.Steps, a.h(a.submitSteps))
	camp.GET(route.Links, a.h(a.linkSuggestions))
	camp.GET(route.Notes, a.h(a.notesPage))
	camp.POST(route.Notes, a.h(a.createNote))
	note.GET("", a.h(a.notePage))
	note.POST(route.Blocks, a.h(a.addBlock))
	note.POST(route.Delete, a.h(a.deleteNote))
	note.POST(route.Rename, a.h(a.renameNote))

	// NPCs.
	camp.GET(route.Entities, a.h(a.entitiesPage))
	camp.POST(route.Entities, gm, a.h(a.createEntity))
	ent.GET("", a.h(a.entityPage))
	ent.POST(route.Rename, gm, a.h(a.renameEntity))
	ent.POST(route.Reveal, gm, a.h(a.revealEntity))
	ent.POST(route.Fields, gm, a.h(a.addEntityField))
	ent.POST(route.FieldsText, gm, a.h(a.setEntityFieldText))
	ent.POST(route.FieldsReveal, gm, a.h(a.revealEntityField))
	ent.POST(route.FieldsRemove, gm, a.h(a.removeEntityField))
	ent.POST(route.Delete, gm, a.h(a.deleteEntity))
	ent.POST(route.Images, gm, a.h(a.addEntityLook))
	ent.POST(route.ImageAction, gm, a.h(a.entityLook))
	ent.POST(route.Link, gm, a.h(a.linkEntity))

	// Images and the display window.
	att.POST(route.ToImage, gm, a.h(a.attachmentToImage))
	camp.GET(route.Display, gm, a.h(a.displayPage))
	camp.GET(route.Images, gm, a.h(a.imagesPage))
	camp.POST(route.Images, gm, a.h(a.createImage))
	camp.GET(route.ImagesImport, gm, a.importUploadPage)
	camp.POST(route.ImagesImport, gm, a.importUpload)
	camp.GET(route.ImagesImportReview, gm, a.h(a.importReviewPage))
	camp.POST(route.ImagesImportReview, gm, a.h(a.importCommit))
	camp.GET(route.ImagesNew, gm, a.h(a.newImagePage))
	img.GET("", a.h(a.imagePage))
	img.POST("", gm, a.h(a.updateImage))
	img.POST(route.Delete, gm, a.h(a.deleteImage))
	img.GET(route.File, a.h(a.imageFile))
	img.POST(route.Npc, gm, a.h(a.imageToEntity))

	// Files: every campaign's, and each campaign's alone.
	for _, base := range []string{route.Files, route.CampaignFiles} {
		r.GET(base, gm, a.h(a.filesPage))
		r.POST(base+route.DeleteUnused, gm, a.h(a.deleteUnusedFiles))
		r.POST(base+route.AttachmentDelete, gm, a.h(a.deleteFile))
	}

	// Compendium: the whole one, and each campaign's (its sources and
	// homebrew), with the same pages under both.
	a.compendiumRoutes(comp, gm)
	a.compendiumRoutes(campComp, gm)

	r.NoRoute(
		func(c *gin.Context) { a.fail(c, http.StatusNotFound, "There is no page at this address.") },
	)
	return r
}

var validateTagOnce sync.Once

// useValidateTag makes gin's binding read `validate` struct tags, the same
// tag every other struct in the project uses, instead of gin's `binding`.
func useValidateTag() {
	validateTagOnce.Do(func() {
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if !ok {
			// Without this, `validate` tags would be silently ignored.
			panic("gin binding validator is not go-playground/validator")
		}
		v.SetTagName("validate")
	})
}

// sameOriginOnly refuses state-changing requests sent by other sites. The
// host's browser is admin without a cookie, so SameSite cannot protect it.
func sameOriginOnly(cop *http.CrossOriginProtection) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := cop.Check(c.Request); err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

// contentSecurityPolicy: scripts need 'unsafe-inline' (about 40 inline handlers)
// and 'unsafe-eval' (htmx runs hx-on through new Function); nonces cannot cover
// handler attributes. Styles need 'unsafe-inline' (shadcn style= attributes,
// <style> injected by htmx and TipTap). Images may come from https hosts
// (compendium portraits) and blob: (upload previews). The websocket and fetch
// calls are same-origin.
const contentSecurityPolicy = "default-src 'self'; " +
	"script-src 'self' 'unsafe-inline' 'unsafe-eval'; " +
	"style-src 'self' 'unsafe-inline'; " +
	"img-src 'self' data: blob: https:; " +
	"connect-src 'self'; " +
	"font-src 'self'; " +
	"object-src 'none'; " +
	"base-uri 'self'; " +
	"form-action 'self'; " +
	"frame-ancestors 'none'"

// securityHeaders applies to every response. Pages only load scripts and
// styles from this host; attachments get stricter headers of their own.
func securityHeaders(c *gin.Context) {
	h := c.Writer.Header()
	h.Set("Content-Security-Policy", contentSecurityPolicy)
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("Referrer-Policy", "same-origin")
	h.Set("X-Frame-Options", "DENY")
	c.Next()
}
