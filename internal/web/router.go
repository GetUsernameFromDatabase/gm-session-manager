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
	r.GET(PathHome, a.home)
	r.GET(PathAdminLogin, a.adminLoginForm)
	r.POST(PathAdminLogin, a.adminLogin)
	r.POST(views.PathFeedback, a.sendFeedback)
	r.GET(PathJoin, a.joinCampaigns)
	r.GET("/join.png", a.joinQR)
	r.GET("/join/:campaign", a.joinCharacters)
	r.POST("/join/:campaign/claim/:character", a.joinClaim)
	r.POST("/join/:campaign/create", a.joinCreate)
	r.GET("/live", a.liveHTML)
	r.POST(PathLogout, a.logout)
	r.POST("/requests/complete", a.completeRequest)
	r.GET("/ws", a.websocket)

	// Campaigns.
	r.POST("/campaigns", gm, a.createCampaign)
	camp.GET("", gm, a.campaignOverview)
	camp.POST("/characters/import", gm, a.importCharacter)
	camp.POST("/delete", gm, a.deleteCampaign)
	camp.GET("/export", gm, a.exportCampaign)
	camp.POST("/open-join", gm, a.toggleOpenJoin)
	camp.POST("/rename", gm, a.renameCampaign)
	camp.POST("/requests/:request/:decision", gm, a.decideRequest)
	camp.POST("/requests/all/:decision", gm, a.decideAll)
	camp.POST("/sources", gm, a.campaignSourcesSet)
	camp.GET("/switcher", a.switcher)
	r.POST("/sessions/:session/kick", gm, a.kick)

	// Characters.
	ch.GET("", a.characterPage)
	ch.GET("/builder", a.commandBuilder)
	ch.POST("/cmd", a.characterCommand)
	ch.GET("/complete", a.commandCompletions)
	ch.POST("/copy", gm, a.copyCharacter)
	ch.GET("/create", a.wizardStart)
	ch.GET("/create/:step", a.wizardPage)
	ch.POST("/create/:step", a.wizardSave)
	ch.POST("/create/:step/change", a.wizardChange)
	ch.POST("/delete", gm, a.deleteCharacter)
	ch.GET("/download", gm, a.downloadCharacter)
	ch.POST("/images", a.addCharacterLook)
	ch.POST("/images/:image/:action", a.characterLook)
	ch.GET("/levelup", a.levelUpForm)
	ch.POST("/levelup", a.levelUp)
	ch.POST("/levelup/revert", a.revertLevelUp)
	ch.GET("/picks", a.builderPicks)
	ch.POST("/remove", a.removeCharacter)
	ch.POST("/restore", gm, a.restoreCharacter)
	ch.GET("/sources/preview", a.sourcePreview)
	ch.GET("/subclass", a.subclassChangeForm)
	ch.POST("/subclass", a.changeSubclass)
	r.POST("/conditions/:condition/delete", a.deleteCondition)

	// Combat.
	camp.GET("/combat", a.combatPage)
	camp.POST("/encounters", gm, a.createEncounter)
	r.POST("/combatants/:combatant/:action", gm, a.combatantAction)
	r.POST("/encounters/:encounter/:action", gm, a.encounterAction)
	r.POST("/encounters/:encounter/combatants", gm, a.addCombatant)

	// Notes.
	r.GET("/attachments/:attachment", a.attachment)
	block.POST("/access", a.blockAccess)
	block.POST("/attachments", a.uploadAttachment)
	block.POST("/delete", a.deleteBlock)
	block.GET("/doc", a.blockDoc)
	block.POST("/move", a.moveBlock)
	block.POST("/snapshot", a.blockSnapshot)
	block.GET("/steps", a.blockSteps)
	block.POST("/steps", a.submitSteps)
	camp.GET("/links", a.linkSuggestions)
	camp.GET("/notes", a.notesPage)
	camp.POST("/notes", a.createNote)
	note.GET("", a.notePage)
	note.POST("/blocks", a.addBlock)
	note.POST("/delete", a.deleteNote)
	note.POST("/rename", a.renameNote)

	// NPCs.
	camp.GET("/entities", a.entitiesPage)
	camp.POST("/entities", gm, a.createEntity)
	ent.GET("", a.entityPage)
	ent.POST("", gm, a.updateEntity)
	ent.POST("/delete", gm, a.deleteEntity)
	ent.POST("/images", gm, a.addEntityLook)
	ent.POST("/images/:image/:action", gm, a.entityLook)
	ent.POST("/link", gm, a.linkEntity)

	// Images and the display window.
	r.POST("/attachments/:attachment/to-image", gm, a.attachmentToImage)
	camp.GET("/display", gm, a.displayPage)
	camp.GET("/images", gm, a.imagesPage)
	camp.POST("/images", gm, a.createImage)
	camp.GET("/images/import", gm, a.importUploadPage)
	camp.POST("/images/import", gm, a.importUpload)
	camp.GET("/images/import/:token", gm, a.importReviewPage)
	camp.POST("/images/import/:token", gm, a.importCommit)
	camp.GET("/images/new", gm, a.newImagePage)
	img.GET("", a.imagePage)
	img.POST("", gm, a.updateImage)
	img.POST("/delete", gm, a.deleteImage)
	img.GET("/file", a.imageFile)
	img.POST("/npc", gm, a.imageToEntity)

	// Files: every campaign's, and each campaign's alone.
	for _, base := range []string{views.PathFiles, "/campaigns/:campaign/files"} {
		r.GET(base, gm, a.filesPage)
		r.POST(base+"/delete-unused", gm, a.deleteUnusedFiles)
		r.POST(base+"/:attachment/delete", gm, a.deleteFile)
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
