package compendium

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strings"
	"time"

	"golang.org/x/net/html"

	"gm-session-manager/db"
	"gm-session-manager/db/sqlc"
)

// The GM imports single pages of a few known sites by link. The link's
// host picks the parser, and only these sites' fixed hosts are fetched.
// Each page is its own source, so re-importing it updates it in place and
// it never collides with an SRD entry of the same slug. The sites' own
// licenses do not cover the book text they copy, so none is claimed.

// Site is a web site the GM may import single pages from.
type Site struct {
	Key, Name, Home string
	// Imports says what kinds of page it has.
	Imports string
	// Note says how the site is read, linking NoteURL.
	Note, NoteURL string

	hosts []string
	// page is the page id of a link, check validates an id and gives its
	// kind ("" when no parser knows it: the GM picks it in the preview),
	// link is the id's page for people.
	page  func(u *url.URL) string
	check func(page string) (Kind, error)
	link  func(page string) string
	// load fetches and parses a page: the book it names, its picture.
	load    func(ctx context.Context, page string) (book, image string, e Entry, err error)
	edition func(page string) string
}

// Sites are the sites pages can be imported from.
var Sites = []Site{wikidotSite, aideddSite, fiveToolsSite}

// Page is one fetched page under review.
type Page struct {
	Site        Site
	ID          string
	Book, Image string
	Entry       Entry
}

// URL is the page for people.
func (p Page) URL() string { return p.Site.link(p.ID) }

// Kind is the kind of entry the page holds, "" when the GM picks it.
func (p Page) Kind() Kind {
	k, _ := p.Site.check(p.ID)
	return k
}

// Host is the site's address as people type it.
func (s Site) Host() string { return s.hosts[0] }

// errNotASite sends the GM to the list of sites under the import form.
var errNotASite = errors.New("paste a link from one of the sites listed below")

// SiteByKey is the site with key, if any.
func SiteByKey(key string) (Site, bool) {
	i := slices.IndexFunc(Sites, func(s Site) bool { return s.Key == key })
	if i < 0 {
		return Site{}, false
	}
	return Sites[i], true
}

// ParsePageURL finds the site and page id of a link; the scheme may be
// left out. Any page of a listed site is accepted: the preview shows what
// came of it.
func ParsePageURL(raw string) (Site, string, error) {
	raw = strings.TrimSpace(raw)
	if !strings.Contains(raw, "://") {
		raw = "https://" + raw
	}
	u, err := url.Parse(raw)
	if err != nil || (u.Scheme != "https" && u.Scheme != "http") || u.Port() != "" || u.User != nil {
		return Site{}, "", errNotASite
	}
	for _, s := range Sites {
		if !slices.Contains(s.hosts, strings.ToLower(u.Hostname())) {
			continue
		}
		page := s.page(u)
		if _, err := s.check(page); err != nil {
			return Site{}, "", err
		}
		return s, page, nil
	}
	return Site{}, "", errNotASite
}

// FetchPage fetches and parses one page without saving it, for the GM to
// review.
func FetchPage(ctx context.Context, raw string) (Page, error) {
	s, page, err := ParsePageURL(raw)
	if err != nil {
		return Page{}, err
	}
	book, image, e, err := s.load(ctx, page)
	return Page{Site: s, ID: page, Book: book, Image: image, Entry: e}, err
}

// PageSourceKey is the source key of a site's page.
func PageSourceKey(site, page string) string { return site + ":" + page }

// SavePage stores a reviewed page entry as the page's own source, with its
// picture when it has one, and returns the source's key. The slug comes
// from the page id (or the name, when the GM picked the kind), so saving
// the same page again replaces its entry.
func SavePage(ctx context.Context, store *db.Store, p Page) (string, error) {
	kind, err := p.Site.check(p.ID)
	if err != nil {
		return "", err
	}
	if kind != "" && p.Entry.Kind != kind {
		return "", fmt.Errorf("%s is a %s page, not a %s", p.ID, kind, p.Entry.Kind)
	}
	if p.Image != "" && !strings.HasPrefix(p.Image, "https://") {
		return "", errors.New("the picture link must start with https://")
	}
	e := p.Entry
	e.Slug = Slugify(p.ID[strings.LastIndexAny(p.ID, ":/")+1:])
	if kind == "" {
		e.Slug = Slugify(e.Name)
	}
	src := pageSource(p)
	body, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(body)
	if _, err = storeSource(ctx, store, src, p.URL(), hex.EncodeToString(sum[:]), []Entry{e}); err != nil || p.Image == "" {
		return src.Key, err
	}
	rows, err := store.ListCompendiumKeysBySource(ctx, src.Key)
	if err != nil || len(rows) == 0 {
		return src.Key, err
	}
	_, err = store.SetCompendiumImage(ctx, sqlc.SetCompendiumImageParams{ImageUrl: sql.NullString{String: p.Image, Valid: true}, ID: rows[0].ID})
	return src.Key, err
}

func pageSource(p Page) Source {
	return Source{
		Key: PageSourceKey(p.Site.Key, p.ID), Edition: p.Site.edition(p.ID), License: "Unverified (third-party site copy)",
		Attribution: strings.TrimSpace(strings.TrimSpace(p.Book) + " via " + p.URL()),
	}
}

func always(edition string) func(string) string { return func(string) string { return edition } }

// ParsePlainPage reads a page no parser knows: its <h1> (else its title)
// is the name and the text around the heading the description. The GM
// picks the kind and tidies it up in the preview.
func ParsePlainPage(body []byte) (Entry, error) {
	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return Entry{}, err
	}
	root, name := doc, ""
	if h1 := find(doc, func(n *html.Node) bool { return n.Data == "h1" }); h1 != nil && h1.Parent != nil {
		root, name = h1.Parent, textOf(h1)
	}
	if name == "" {
		name = textOf(find(doc, func(n *html.Node) bool { return n.Data == "title" }))
	}
	if name == "" {
		return Entry{}, errors.New("the page has no title")
	}
	desc := strings.Join(slices.DeleteFunc(textLines(root), func(l string) bool { return l == name }), "\n\n")
	return Entry{Name: name, Data: TextData{Desc: desc}, Text: desc}, nil
}

// pageClient never follows a redirect off the host it was asked for.
var pageClient = &http.Client{
	Timeout: 20 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		if !strings.EqualFold(req.URL.Hostname(), via[0].URL.Hostname()) {
			return errors.New("redirect off " + via[0].URL.Hostname())
		}
		return nil
	},
}

// fetch gets a page of at most maxBody bytes.
func fetch(ctx context.Context, link string, maxBody int64) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)
	if err != nil {
		return nil, err
	}
	res, err := pageClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %w", link, err)
	}
	defer func() { _ = res.Body.Close() }()
	switch res.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return nil, fmt.Errorf("no page at %s", link)
	default:
		return nil, fmt.Errorf("fetch %s: %s", link, res.Status)
	}
	body, err := io.ReadAll(io.LimitReader(res.Body, maxBody+1))
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %w", link, err)
	}
	if int64(len(body)) > maxBody {
		return nil, fmt.Errorf("fetch %s: page too large", link)
	}
	return body, nil
}
