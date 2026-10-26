// Package web holds the HTTP handlers, templ components and static assets.
package web

import (
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"gm-session-manager/auth"
	"gm-session-manager/db"
	"gm-session-manager/hub"
	"gm-session-manager/web/views"
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
	Hub     *hub.Hub
	imports csvImports
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
	r.Use(a.Auth.Middleware())

	gm := auth.RequireAdmin()

	// Identity and joining.
	r.GET(PathHome, a.home)
	r.GET(PathAdminLogin, a.adminLoginForm)
	r.POST(PathAdminLogin, a.adminLogin)
	r.POST(PathLogout, a.logout)
	r.GET(PathJoin, a.joinCampaigns)
	r.GET("/join/:campaign", a.joinCharacters)
	r.POST("/join/:campaign/create", a.joinCreate)
	r.POST("/join/:campaign/claim/:character", a.joinClaim)
	r.POST("/requests/complete", a.completeRequest)
	r.GET("/join.png", a.joinQR)
	r.GET("/ws", a.websocket)
	r.GET("/live", a.liveHTML)

	// GM table controls.
	r.POST("/campaigns", gm, a.createCampaign)
	r.GET("/campaigns/:campaign", gm, a.campaignOverview)
	r.POST("/campaigns/:campaign/open-join", gm, a.toggleOpenJoin)
	r.POST("/campaigns/:campaign/rename", gm, a.renameCampaign)
	r.POST("/campaigns/:campaign/delete", gm, a.deleteCampaign)
	r.POST("/campaigns/:campaign/requests/all/:decision", gm, a.decideAll)
	r.POST("/campaigns/:campaign/requests/:request/:decision", gm, a.decideRequest)
	r.GET("/campaigns/:campaign/export", gm, a.exportCampaign)
	r.POST("/sessions/:session/kick", gm, a.kick)
	r.POST("/characters/:character/delete", gm, a.deleteCharacter)
	r.POST("/characters/:character/copy", gm, a.copyCharacter)
	r.POST("/characters/:character/restore", gm, a.restoreCharacter)
	r.GET("/characters/:character/download", gm, a.downloadCharacter)
	r.POST("/characters/:character/remove", a.removeCharacter)
	r.POST("/campaigns/:campaign/characters/import", gm, a.importCharacter)
	r.GET(views.PathFiles, gm, a.filesPage)
	r.POST(views.PathFiles+"/delete-unused", gm, a.deleteUnusedFiles)
	r.POST(views.PathFiles+"/:attachment/delete", gm, a.deleteFile)
	r.POST("/attachments/:attachment/to-image", gm, a.attachmentToImage)

	// Characters.
	r.GET("/characters/:character", a.characterPage)
	r.POST("/characters/:character/cmd", a.characterCommand)
	r.GET("/characters/:character/builder", a.commandBuilder)
	r.GET("/characters/:character/picks", a.builderPicks)
	r.GET("/characters/:character/complete", a.commandCompletions)
	r.GET("/characters/:character/create", a.wizardStart)
	r.GET("/characters/:character/create/:step", a.wizardPage)
	r.POST("/characters/:character/create/:step", a.wizardSave)
	r.GET("/characters/:character/levelup", a.levelUpForm)
	r.POST("/characters/:character/levelup", a.levelUp)
	r.POST("/characters/:character/levelup/revert", a.revertLevelUp)
	r.POST("/conditions/:condition/delete", a.deleteCondition)

	// Combat.
	r.GET("/campaigns/:campaign/combat", a.combatPage)
	r.POST("/campaigns/:campaign/encounters", gm, a.createEncounter)
	r.POST("/encounters/:encounter/combatants", gm, a.addCombatant)
	r.POST("/encounters/:encounter/:action", gm, a.encounterAction)
	r.POST("/combatants/:combatant/:action", gm, a.combatantAction)

	// Notes.
	r.GET("/campaigns/:campaign/notes", a.notesPage)
	r.POST("/campaigns/:campaign/notes", a.createNote)
	r.GET("/campaigns/:campaign/links", a.linkSuggestions)
	r.GET("/campaigns/:campaign/switcher", a.switcher)
	r.GET("/notes/:note", a.notePage)
	r.POST("/notes/:note/rename", a.renameNote)
	r.POST("/notes/:note/delete", a.deleteNote)
	r.POST("/notes/:note/blocks", a.addBlock)
	r.GET("/blocks/:block/doc", a.blockDoc)
	r.GET("/blocks/:block/steps", a.blockSteps)
	r.POST("/blocks/:block/steps", a.submitSteps)
	r.POST("/blocks/:block/snapshot", a.blockSnapshot)
	r.POST("/blocks/:block/access", a.blockAccess)
	r.POST("/blocks/:block/move", a.moveBlock)
	r.POST("/blocks/:block/delete", a.deleteBlock)
	r.POST("/blocks/:block/attachments", a.uploadAttachment)
	r.GET("/attachments/:attachment", a.attachment)

	// Entities.
	r.GET("/campaigns/:campaign/entities", a.entitiesPage)
	r.POST("/campaigns/:campaign/entities", gm, a.createEntity)
	r.GET("/entities/:entity", a.entityPage)
	r.POST("/entities/:entity", gm, a.updateEntity)
	r.POST("/entities/:entity/delete", gm, a.deleteEntity)
	r.POST("/entities/:entity/link", gm, a.linkEntity)
	r.POST("/entities/:entity/images", gm, a.addEntityLook)
	r.POST("/entities/:entity/images/:image/:action", gm, a.entityLook)
	r.POST("/characters/:character/images", a.addCharacterLook)
	r.POST("/characters/:character/images/:image/:action", a.characterLook)

	// Images and the display window (GM only).
	r.GET("/campaigns/:campaign/images", gm, a.imagesPage)
	r.POST("/campaigns/:campaign/images", gm, a.createImage)
	r.GET("/campaigns/:campaign/images/new", gm, a.newImagePage)
	r.GET("/campaigns/:campaign/images/import", gm, a.importUploadPage)
	r.POST("/campaigns/:campaign/images/import", gm, a.importUpload)
	r.GET("/campaigns/:campaign/images/import/:token", gm, a.importReviewPage)
	r.POST("/campaigns/:campaign/images/import/:token", gm, a.importCommit)
	r.GET("/campaigns/:campaign/display", gm, a.displayPage)
	r.GET("/images/:image", a.imagePage)
	r.POST("/images/:image", gm, a.updateImage)
	r.POST("/images/:image/delete", gm, a.deleteImage)
	r.POST("/images/:image/npc", gm, a.imageToEntity)
	r.GET("/images/:image/file", a.imageFile)

	// Compendium.
	r.GET("/compendium", a.compendiumPage)
	r.GET("/compendium/search", a.compendiumSearch)
	r.GET("/compendium/random", a.compendiumRandom)
	r.GET("/compendium/lookup", a.compendiumLookup)
	r.GET("/compendium/entries/:entry", a.compendiumEntry)
	// Homebrew: the handlers check who may write (perm.CanEditHomebrew).
	r.GET("/compendium/entries/:entry/edit", a.homebrewEdit)
	r.POST("/compendium/entries/:entry", a.homebrewUpdate)
	r.POST("/compendium/entries/:entry/delete", a.homebrewDelete)
	r.POST("/compendium/entries/:entry/campaigns", gm, a.homebrewLink)
	r.POST("/compendium/entries/:entry/campaigns/:campaign/delete", gm, a.homebrewUnlink)
	r.POST("/compendium/entries/:entry/image", gm, a.compendiumImage)
	r.GET("/compendium/habitats", gm, a.habitatsPage)
	r.POST("/compendium/habitats/cell", gm, a.habitatCell)
	r.POST("/compendium/habitats/any", gm, a.habitatAny)
	r.POST("/compendium/habitats/batch", gm, a.habitatBatch)
	r.GET("/compendium/homebrew/new", a.homebrewNew)
	r.POST("/compendium/homebrew", a.homebrewCreate)
	r.POST("/compendium/import", gm, a.compendiumImport)
	r.POST("/compendium/folder", gm, a.compendiumFolder)
	r.POST("/compendium/page/preview", gm, a.compendiumPagePreview)
	r.POST("/compendium/page", gm, a.compendiumPageSave)
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
