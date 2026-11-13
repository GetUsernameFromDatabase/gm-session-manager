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
	"gm-session-manager/internal/web/views"
)

// Route paths used in more than one place.
const (
	PathHome       = "/"
	PathAdminLogin = "/admin/login"
	PathLogout     = "/logout"
	PathJoin       = "/join"
)

// Deps is what the web layer needs from main.
type Deps struct {
	Auth  *auth.Manager
	Store *db.Store
	// DataDir holds attachments/ and compendium/.
	DataDir string
	Port    int
	// LANIPs and Fingerprint are shown on the GM's join panel.
	LANIPs      []net.IP
	Fingerprint string
}

// App is the web layer: the router and the websocket hub it renders for.
type App struct {
	Deps
	Hub       *hub.Hub
	imports   csvImports
	reqs      reqLog
	fiveTools fiveToolsJob
	state     appState
}

// New builds the app. The hub renders fragments through the app.
func New(d Deps) *App {
	a := &App{Deps: d}
	a.Hub = hub.New(a)
	return a
}

// Close disconnects all websockets.
func (a *App) Close() { a.Hub.Close() }

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
	r.Use(a.Auth.Middleware(), a.recordRequest)

	gm := auth.RequireAdmin(a.gmOnly)

	block := r.Group("/blocks/:block")
	camp := r.Group("/campaigns/:campaign")
	campComp := camp.Group("/compendium", a.compendiumCampaign)
	ch := r.Group("/characters/:character")
	comp := r.Group("/compendium", a.compendiumWhole)
	ent := r.Group("/entities/:entity")
	img := r.Group("/images/:image")
	note := r.Group("/notes/:note")

	// Within each section: sorted by full path, then GET before POST.

	// Joining.
	r.GET(PathHome, a.h(a.home))
	r.GET(PathAdminLogin, a.adminLoginForm)
	r.POST(PathAdminLogin, a.h(a.adminLogin))
	r.POST(views.PathFeedback, a.h(a.sendFeedback))
	r.GET(PathJoin, a.h(a.joinCampaigns))
	r.GET("/join.png", a.h(a.joinQR))
	r.GET("/join/:campaign", a.joinCharacters)
	r.POST("/join/:campaign/claim/:character", a.h(a.joinClaim))
	r.POST("/join/:campaign/create", a.h(a.joinCreate))
	r.GET("/live", a.h(a.liveHTML))
	r.POST(PathLogout, a.h(a.logout))
	r.POST("/requests/complete", a.h(a.completeRequest))
	r.GET("/ws", a.websocket)

	// Campaigns.
	r.POST("/campaigns", gm, a.h(a.createCampaign))
	camp.GET("", gm, a.h(a.campaignOverview))
	camp.POST("/characters/import", gm, a.h(a.importCharacter))
	camp.POST("/delete", gm, a.h(a.deleteCampaign))
	camp.GET("/export", gm, a.h(a.exportCampaign))
	camp.POST("/open-join", gm, a.h(a.toggleOpenJoin))
	camp.POST("/rename", gm, a.h(a.renameCampaign))
	camp.POST("/requests/:request/:decision", gm, a.decideRequest)
	camp.POST("/requests/all/:decision", gm, a.decideAll)
	camp.POST("/sources", gm, a.h(a.campaignSourcesSet))
	camp.GET("/switcher", a.h(a.switcher))
	r.POST("/sessions/:session/kick", gm, a.h(a.kick))

	// Characters.
	ch.GET("", a.h(a.characterPage))
	ch.GET("/builder", a.h(a.commandBuilder))
	ch.POST("/cmd", a.h(a.characterCommand))
	ch.GET("/complete", a.h(a.commandCompletions))
	ch.POST("/copy", gm, a.h(a.copyCharacter))
	ch.GET("/create", a.wizardStart)
	ch.GET("/create/:step", a.h(a.wizardPage))
	ch.POST("/create/:step", a.h(a.wizardSave))
	ch.POST("/create/:step/change", a.h(a.wizardChange))
	ch.POST("/delete", gm, a.h(a.deleteCharacter))
	ch.GET("/download", gm, a.h(a.downloadCharacter))
	ch.POST("/images", a.h(a.addCharacterLook))
	ch.POST("/images/:image/:action", a.h(a.characterLook))
	ch.GET("/levelup", a.h(a.levelUpForm))
	ch.POST("/levelup", a.h(a.levelUp))
	ch.POST("/levelup/revert", a.h(a.revertLevelUp))
	ch.GET("/picks", a.h(a.builderPicks))
	ch.POST("/remove", a.h(a.removeCharacter))
	ch.POST("/restore", gm, a.h(a.restoreCharacter))
	ch.GET("/sources/preview", a.h(a.sourcePreview))
	ch.GET("/subclass", a.h(a.subclassChangeForm))
	ch.POST("/subclass", a.h(a.changeSubclass))
	r.POST("/conditions/:condition/delete", a.h(a.deleteCondition))

	// Combat.
	camp.GET("/combat", a.h(a.combatPage))
	camp.POST("/encounters", gm, a.h(a.createEncounter))
	r.POST("/combatants/:combatant/:action", gm, a.h(a.combatantAction))
	r.POST("/encounters/:encounter/:action", gm, a.h(a.encounterAction))
	r.POST("/encounters/:encounter/combatants", gm, a.h(a.addCombatant))

	// Notes.
	r.GET("/attachments/:attachment", a.attachment)
	block.POST("/access", a.h(a.blockAccess))
	block.POST("/attachments", a.uploadAttachment)
	block.POST("/delete", a.h(a.deleteBlock))
	block.GET("/doc", a.h(a.blockDoc))
	block.POST("/move", a.h(a.moveBlock))
	block.POST("/snapshot", a.h(a.blockSnapshot))
	block.GET("/steps", a.h(a.blockSteps))
	block.POST("/steps", a.h(a.submitSteps))
	camp.GET("/links", a.h(a.linkSuggestions))
	camp.GET("/notes", a.h(a.notesPage))
	camp.POST("/notes", a.h(a.createNote))
	note.GET("", a.h(a.notePage))
	note.POST("/blocks", a.h(a.addBlock))
	note.POST("/delete", a.h(a.deleteNote))
	note.POST("/rename", a.h(a.renameNote))

	// NPCs.
	camp.GET("/entities", a.h(a.entitiesPage))
	camp.POST("/entities", gm, a.h(a.createEntity))
	ent.GET("", a.h(a.entityPage))
	ent.POST("", gm, a.h(a.updateEntity))
	ent.POST("/delete", gm, a.h(a.deleteEntity))
	ent.POST("/images", gm, a.h(a.addEntityLook))
	ent.POST("/images/:image/:action", gm, a.h(a.entityLook))
	ent.POST("/link", gm, a.h(a.linkEntity))

	// Images and the display window.
	r.POST("/attachments/:attachment/to-image", gm, a.h(a.attachmentToImage))
	camp.GET("/display", gm, a.h(a.displayPage))
	camp.GET("/images", gm, a.h(a.imagesPage))
	camp.POST("/images", gm, a.h(a.createImage))
	camp.GET("/images/import", gm, a.importUploadPage)
	camp.POST("/images/import", gm, a.importUpload)
	camp.GET("/images/import/:token", gm, a.h(a.importReviewPage))
	camp.POST("/images/import/:token", gm, a.h(a.importCommit))
	camp.GET("/images/new", gm, a.h(a.newImagePage))
	img.GET("", a.h(a.imagePage))
	img.POST("", gm, a.h(a.updateImage))
	img.POST("/delete", gm, a.h(a.deleteImage))
	img.GET("/file", a.h(a.imageFile))
	img.POST("/npc", gm, a.h(a.imageToEntity))

	// Files: every campaign's, and each campaign's alone.
	for _, base := range []string{views.PathFiles, "/campaigns/:campaign/files"} {
		r.GET(base, gm, a.h(a.filesPage))
		r.POST(base+"/delete-unused", gm, a.h(a.deleteUnusedFiles))
		r.POST(base+"/:attachment/delete", gm, a.h(a.deleteFile))
	}

	// Compendium: the whole one, and each campaign's (its sources and
	// homebrew), with the same pages under both.
	a.compendiumRoutes(comp, gm)
	a.compendiumRoutes(campComp, gm)

	r.NoRoute(func(c *gin.Context) { a.fail(c, http.StatusNotFound, "There is no page at this address.") })
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

// securityHeaders applies to every response. Pages only load scripts and
// styles from this host; attachments get stricter headers of their own.
func securityHeaders(c *gin.Context) {
	h := c.Writer.Header()
	h.Set("X-Content-Type-Options", "nosniff")
	h.Set("Referrer-Policy", "same-origin")
	h.Set("X-Frame-Options", "DENY")
	c.Next()
}
