package web

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	qrcode "github.com/skip2/go-qrcode"

	"gm-session-manager/auth"
	"gm-session-manager/db/sqlc"
	"gm-session-manager/dnd5e"
	"gm-session-manager/hub"
	"gm-session-manager/perm"
	"gm-session-manager/web/views"
)

type adminLoginForm struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type nameForm struct {
	Name string `form:"name" validate:"required,max=100"`
}

// playerForm is who a visitor says they are, shown to the GM with a join
// request and kept only in memory.
type playerForm struct {
	Player string `form:"player" validate:"required,max=12"`
}

// home sends everyone where they belong: the GM to the campaigns, a player
// to their character, a visitor to the join flow.
func (a *App) home(c *gin.Context) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return
	}
	switch {
	case w.isGM():
		campaigns, err := a.Store.ListCampaigns(c.Request.Context())
		if err != nil {
			a.serverError(c, err)
			return
		}
		a.render(c, http.StatusOK, views.GMHome(a.layout(w, "Campaigns", views.NavHome, nil, false), campaigns, a.joinInfo(0), a.feedbackPath()))
	case w.character != nil:
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/characters/%d", w.character.ID))
	default:
		c.Redirect(http.StatusSeeOther, PathJoin)
	}
}

// joinURLs are the addresses players can reach this host at, ending in path.
func (a *App) joinURLs(path string) []string {
	var out []string
	for _, ip := range a.LANIPs {
		out = append(out, fmt.Sprintf("https://%s%s", net.JoinHostPort(ip.String(), fmt.Sprint(a.Port)), path))
	}
	return out
}

// joinPath is where players land: the campaign's join page, or the campaign
// list when campaignID is 0.
func joinPath(campaignID int64) string {
	if campaignID == 0 {
		return PathJoin
	}
	return fmt.Sprintf("%s/%d", PathJoin, campaignID)
}

func (a *App) joinInfo(campaignID int64) views.JoinInfo {
	return views.JoinInfo{URLs: a.joinURLs(joinPath(campaignID)), QR: views.JoinQRSrc(campaignID), Fingerprint: a.Fingerprint}
}

// queryCampaign picks the campaign a join QR code leads to (views.JoinQRSrc).
const queryCampaign = "campaign"

// joinQR is the QR code for the first join address, shown on the GM page.
func (a *App) joinQR(c *gin.Context) {
	var campaignID int64
	if v := c.Query(queryCampaign); v != "" {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil || id <= 0 {
			c.Status(http.StatusBadRequest)
			return
		}
		campaignID = id
	}
	urls := a.joinURLs(joinPath(campaignID))
	if len(urls) == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	png, err := qrcode.Encode(urls[0], qrcode.Medium, 256)
	if err != nil {
		a.serverError(c, err)
		return
	}
	c.Header("Cache-Control", "no-cache")
	c.Data(http.StatusOK, "image/png", png)
}

func (a *App) adminLoginForm(c *gin.Context) {
	if auth.FromContext(c).Session != nil {
		c.Redirect(http.StatusSeeOther, PathHome)
		return
	}
	a.render(c, http.StatusOK, views.AdminLogin(views.Layout{Title: "GM login"}, ""))
}

// adminLogin creates an admin session for a device that knows the .env
// credentials. A device that already has a session must log out first: a
// device holds one session, and its role is chosen at login.
func (a *App) adminLogin(c *gin.Context) {
	if auth.FromContext(c).Session != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	var form adminLoginForm
	if err := c.ShouldBind(&form); err != nil {
		a.render(c, http.StatusBadRequest, views.AdminLogin(views.Layout{Title: "GM login"}, "Enter a username and password."))
		return
	}

	token, _, err := a.Auth.LoginAdmin(c.Request.Context(), form.Username, form.Password)
	if errors.Is(err, auth.ErrBadCredentials) {
		a.render(c, http.StatusUnauthorized, views.AdminLogin(views.Layout{Title: "GM login"}, "Wrong username or password."))
		return
	}
	if err != nil {
		a.serverError(c, err)
		return
	}

	auth.SetCookie(c, token)
	c.Redirect(http.StatusSeeOther, PathHome)
}

// logout deletes this device's session row, the same thing a kick does. On
// localhost this is also how a player switches back to admin.
func (a *App) logout(c *gin.Context) {
	if sess := auth.FromContext(c).Session; sess != nil {
		w, err := a.whoami(c)
		if err != nil {
			a.serverError(c, err)
			return
		}
		if err := a.Auth.Revoke(c.Request.Context(), sess.ID); err != nil {
			a.serverError(c, err)
			return
		}
		a.Hub.CloseSession(sess.ID)
		a.reqs.forget(sessionKey(*sess))
		if w.character != nil {
			a.Hub.Publish(w.character.CampaignID, hub.Topic{Kind: hub.TopicCampaign, ID: w.character.CampaignID})
		}
	}
	auth.ClearCookie(c)
	c.Redirect(http.StatusSeeOther, PathHome)
}

// visitorOnly refuses join actions from a device that already has a role.
func visitorOnly(c *gin.Context) bool {
	if auth.FromContext(c).Session != nil {
		c.Redirect(http.StatusSeeOther, PathHome)
		return false
	}
	return true
}

func (a *App) joinCampaigns(c *gin.Context) {
	if !visitorOnly(c) {
		return
	}
	campaigns, err := a.Store.ListCampaigns(c.Request.Context())
	if err != nil {
		a.serverError(c, err)
		return
	}
	if len(campaigns) == 1 {
		c.Redirect(http.StatusSeeOther, joinPath(campaigns[0].ID))
		return
	}
	a.render(c, http.StatusOK, views.JoinCampaigns(views.Layout{Title: "Join"}, campaigns))
}

func (a *App) joinCharacters(c *gin.Context) {
	if !visitorOnly(c) {
		return
	}
	camp, ok := a.joinCampaign(c)
	if !ok {
		return
	}
	a.renderJoinCharacters(c, http.StatusOK, camp, "")
}

func (a *App) renderJoinCharacters(c *gin.Context, status int, camp sqlc.Campaign, msg string) {
	chars, err := a.Store.ListCharactersByCampaign(c.Request.Context(), camp.ID)
	if err != nil {
		a.serverError(c, err)
		return
	}
	a.render(c, status, views.JoinCharacters(views.Layout{Title: camp.Name}, camp, chars, msg))
}

func (a *App) joinCampaign(c *gin.Context) (sqlc.Campaign, bool) {
	id, ok := a.idParam(c, "campaign")
	if !ok {
		return sqlc.Campaign{}, false
	}
	camp, err := a.Store.GetCampaign(c.Request.Context(), id)
	return camp, a.found(c, err)
}

// joinCreate creates a character. With open_join the session is created at
// once; otherwise the request waits for the GM like a claim.
func (a *App) joinCreate(c *gin.Context) {
	if !visitorOnly(c) {
		return
	}
	camp, ok := a.joinCampaign(c)
	if !ok {
		return
	}
	var form nameForm
	if err := c.ShouldBind(&form); err != nil {
		a.renderJoinCharacters(c, http.StatusUnprocessableEntity, camp, "Give your character a name (up to 100 characters).")
		return
	}
	name := strings.TrimSpace(form.Name)
	if !camp.OpenJoin {
		a.startRequest(c, camp, hub.PendingCreate, 0, name)
		return
	}
	id, err := a.createAndLogin(c, camp.ID, name)
	if err != nil {
		a.serverError(c, err)
		return
	}
	c.Redirect(http.StatusSeeOther, wizardPath(id, wizardSteps[0].Key))
}

// createAndLogin creates a character and a session for it on this device,
// returning the character's id.
func (a *App) createAndLogin(c *gin.Context, campaignID int64, name string) (int64, error) {
	ctx := c.Request.Context()
	sheet, err := dnd5e.NewSheet().JSON()
	if err != nil {
		return 0, err
	}
	ch, err := a.Store.CreateCharacterWithSheet(ctx, sqlc.CreateCharacterWithSheetParams{
		CampaignID: campaignID, Name: name, System: dnd5e.System, SheetJson: sheet,
	})
	if err != nil {
		return 0, err
	}
	if err := a.login(c, ch); err != nil {
		_ = a.Store.DeleteCharacter(ctx, ch.ID)
		return 0, err
	}
	return ch.ID, nil
}

func (a *App) login(c *gin.Context, ch sqlc.Character) error {
	token, _, err := a.Auth.IssuePlayer(c.Request.Context(), ch.ID)
	if err != nil {
		return err
	}
	auth.SetCookie(c, token)
	a.Hub.Publish(ch.CampaignID, hub.Topic{Kind: hub.TopicCampaign, ID: ch.CampaignID})
	return nil
}

func (a *App) joinClaim(c *gin.Context) {
	if !visitorOnly(c) {
		return
	}
	camp, ok := a.joinCampaign(c)
	if !ok {
		return
	}
	charID, ok := a.idParam(c, "character")
	if !ok {
		return
	}
	ch, err := a.Store.GetCharacter(c.Request.Context(), charID)
	if !a.found(c, err) {
		return
	}
	if ch.CampaignID != camp.ID {
		a.fail(c, http.StatusNotFound, "Not found.")
		return
	}
	a.startRequest(c, camp, hub.PendingClaim, ch.ID, ch.Name)
}

// startRequest records a pending request and shows the waiting page, whose
// websocket keeps the request alive.
func (a *App) startRequest(c *gin.Context, camp sqlc.Campaign, kind hub.PendingKind, characterID int64, name string) {
	var form playerForm
	if err := c.ShouldBind(&form); err != nil {
		a.renderJoinCharacters(c, http.StatusUnprocessableEntity, camp, "Tell the GM your name (up to 12 characters).")
		return
	}
	ua := c.Request.UserAgent()
	p := a.Hub.AddPending(hub.Pending{
		CampaignID: camp.ID, CampaignName: camp.Name, Kind: kind, CharacterID: characterID, Name: name,
		Player: strings.TrimSpace(form.Player), OS: deviceOS(ua), Phone: isPhone(ua), Host: lanHostname(c.Request.Context(), c.Request.RemoteAddr),
	})
	epoch, seq := a.Hub.Position()
	l := views.Layout{Title: "Waiting", WS: &views.WS{URL: "/ws?pending=" + p.Secret, Epoch: epoch, Seq: seq}}
	a.render(c, http.StatusOK, views.JoinWaiting(l, views.Waiting{Secret: p.Secret, Name: p.Name, State: "waiting"}))
}

// completeRequest redeems an approved request: the session is created and
// the cookie set on the device that asked.
func (a *App) completeRequest(c *gin.Context) {
	if auth.FromContext(c).Session != nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	p, ok := a.Hub.TakeApproved(c.PostForm("secret"))
	if !ok {
		a.fail(c, http.StatusNotFound, "That request is no longer waiting. Please ask again.")
		return
	}
	var err error
	next := PathHome
	if p.Kind == hub.PendingCreate {
		var id int64
		id, err = a.createAndLogin(c, p.CampaignID, p.Name)
		next = wizardPath(id, wizardSteps[0].Key)
	} else {
		ch, gerr := a.Store.GetCharacter(c.Request.Context(), p.CharacterID)
		if !a.found(c, gerr) { // deleted while the request waited
			return
		}
		err = a.login(c, ch)
	}
	if err != nil {
		a.serverError(c, err)
		return
	}
	redirect(c, next)
}

// decideRequest is the GM's approve or deny.
func (a *App) decideRequest(c *gin.Context) {
	campID, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	approve, ok := a.decision(c)
	if !ok {
		return
	}
	if !a.Hub.Decide(campID, c.Param("request"), approve) {
		a.fail(c, http.StatusConflict, "That request is no longer waiting.")
		return
	}
	a.render(c, http.StatusOK, views.PendingRequests(a.pendingViews()))
}

// decision reads the approve or deny in the URL, answering 404 otherwise.
func (a *App) decision(c *gin.Context) (approve, ok bool) {
	switch c.Param("decision") {
	case "approve":
		return true, true
	case "deny":
		return false, true
	}
	a.fail(c, http.StatusNotFound, "Not found.")
	return false, false
}

// idsForm is the requests a campaign's approve-all or deny-all bar showed.
type idsForm struct {
	IDs []string `form:"id" validate:"required,dive,required"`
}

// decideAll approves or denies the requests the GM saw; any already decided
// or gone are skipped, and newer ones keep waiting.
func (a *App) decideAll(c *gin.Context) {
	campID, ok := a.idParam(c, "campaign")
	if !ok {
		return
	}
	approve, ok := a.decision(c)
	if !ok {
		return
	}
	var form idsForm
	if !a.bindForm(c, &form) {
		return
	}
	for _, id := range form.IDs {
		a.Hub.Decide(campID, id, approve)
	}
	a.render(c, http.StatusOK, views.PendingRequests(a.pendingViews()))
}

// websocket connects a page to the hub. The GM names the campaign; a
// player's campaign is their character's; a visitor needs a pending
// request's secret.
func (a *App) websocket(c *gin.Context) {
	w, err := a.whoami(c)
	if err != nil {
		a.serverError(c, err)
		return
	}
	ua := userAgentSummary(c.Request.UserAgent())
	var sessionID int64
	if w.id.Session != nil {
		sessionID = w.id.Session.ID
	}
	var client *hub.Client
	switch {
	case w.isGM():
		// Pages outside a campaign still get join requests.
		client = hub.NewClient(formIntQuery(c, "campaign"), w.viewer, sessionID, ua)
	case w.character != nil:
		client = hub.NewClient(w.character.CampaignID, w.viewer, sessionID, ua)
	default:
		p, ok := a.Hub.PendingBySecret(c.Query("pending"))
		if !ok || w.id.Session != nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		client = hub.NewClient(p.CampaignID, perm.Viewer{}, 0, ua)
		client.Pending = p
	}
	a.Hub.Serve(c.Writer, c.Request, client)
}
