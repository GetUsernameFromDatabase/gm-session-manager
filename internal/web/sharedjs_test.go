package web

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"path"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"testing"

	"gm-session-manager/internal/web/views"
)

// sharedExportLine is one "export const NAME = JSON;" line of shared.js.
var sharedExportLine = regexp.MustCompile(`^export const ([A-Z_0-9]+) = (.*);$`)

func parseSharedJS(t *testing.T, body string) map[string]string {
	t.Helper()
	got := map[string]string{}
	for line := range strings.SplitSeq(strings.TrimSpace(body), "\n") {
		m := sharedExportLine.FindStringSubmatch(line)
		if m == nil {
			t.Fatalf("not an export line: %q", line)
		}
		got[m[1]] = m[2]
	}
	return got
}

// tagHas fails unless the validate tag of v's field holds the rule (e.g. "max=500").
func tagHas(t *testing.T, v any, field, rule string) {
	t.Helper()
	f, ok := reflect.TypeOf(v).FieldByName(field)
	if !ok {
		t.Fatalf("%T has no field %s", v, field)
	}
	if !slices.Contains(strings.Split(f.Tag.Get("validate"), ","), rule) {
		t.Errorf("%T.%s tag %q lacks %q", v, field, f.Tag.Get("validate"), rule)
	}
}

// TestSharedJS: shared.js is served revalidated and holds every Go value the browser copies.
func TestSharedJS(t *testing.T) {
	e := newTestEnv(t)
	get := func(ifNoneMatch string) *httptest.ResponseRecorder {
		req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, "/assets/js/shared.js", nil)
		req.RemoteAddr = lanAddr
		if ifNoneMatch != "" {
			req.Header.Set("If-None-Match", ifNoneMatch)
		}
		w := httptest.NewRecorder()
		e.router.ServeHTTP(w, req)
		return w
	}
	w := get("")
	etag := w.Header().Get("ETag")
	if w.Code != http.StatusOK || etag == "" || w.Header().Get("Cache-Control") != "no-cache" || !strings.HasPrefix(w.Header().Get("Content-Type"), "text/javascript") {
		t.Fatalf("%d etag %q cache %q type %q", w.Code, etag, w.Header().Get("Cache-Control"), w.Header().Get("Content-Type"))
	}
	if w := get(etag); w.Code != http.StatusNotModified {
		t.Errorf("revalidated: %d", w.Code)
	}

	got := parseSharedJS(t, w.Body.String())
	want := map[string]any{
		"SAFE_SCHEMES":   safeSchemes,
		"MAX_STEPS":      maxSteps,
		"MAX_STEP_BYTES": maxStepBytes,
		"CONSOLE_SIZE":   consoleSize,
		"CONSOLE_MSG":    consoleMsg,
		"CONSOLE_PAGE":   consolePage,
		"CONSOLE_LEVELS": consoleLevels,
		"ROUTES":         sharedRoutes,
		"WIKILINK":       views.WikiLinkPattern,
		"MSG":            sharedMsg,
		"KEY":            sharedKeys,
	}
	for name, v := range want {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		if got[name] != string(b) {
			t.Errorf("%s = %s, want %s", name, got[name], b)
		}
	}
	if len(got) != len(want) {
		t.Errorf("shared.js exports %d names, test knows %d", len(got), len(want))
	}
}

// TestSharedRoutesRegistered: a route the browser builds URLs from must exist
// in the router, so a renamed route cannot leave the JS pointing nowhere.
func TestSharedRoutesRegistered(t *testing.T) {
	registered := map[string]bool{}
	for _, r := range newTestEnv(t).router.Routes() {
		registered[r.Path] = true
	}
	for name, route := range sharedRoutes {
		if !registered[route] {
			t.Errorf("ROUTES.%s = %q is not registered in the router", name, route)
		}
	}
}

var (
	importLine   = regexp.MustCompile(`(?m)^import \{([^}]*)\} from "/assets/js/([\w.-]+\.js)"`)
	exportedName = `export (?:const|function|class|let) %s\b`
)

// TestSharedImportsExist: a name imported from our own module but not
// exported fails the module load and kills the page.
func TestSharedImportsExist(t *testing.T) {
	shared := parseSharedJS(t, string(sharedJS))
	files, err := fs.Glob(assetsFS, "assets/js/*.js")
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if strings.HasPrefix(path.Base(f), "shadcn-templ-") {
			continue
		}
		src, err := fs.ReadFile(assetsFS, f)
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range importLine.FindAllStringSubmatch(string(src), -1) {
			var from []byte // shared.js is built in Go, not embedded
			if m[2] != "shared.js" {
				if from, err = fs.ReadFile(assetsFS, "assets/js/"+m[2]); err != nil {
					t.Errorf("%s imports missing file %s", f, m[2])
					continue
				}
			}
			for spec := range strings.SplitSeq(m[1], ",") {
				name, _, _ := strings.Cut(strings.TrimSpace(spec), " as ")
				if name == "" {
					continue
				}
				if m[2] == "shared.js" {
					if _, ok := shared[name]; !ok {
						t.Errorf("%s imports %s, not exported by shared.js", f, name)
					}
				} else if !regexp.MustCompile(fmt.Sprintf(exportedName, regexp.QuoteMeta(name))).Match(from) {
					t.Errorf("%s imports %s, not exported by %s", f, name, m[2])
				}
			}
		}
	}
}

// TestSharedLimitsMatchTags: validator tags cannot name constants, so the
// constants the browser gets must equal the literals in the tags.
func TestSharedLimitsMatchTags(t *testing.T) {
	tagHas(t, submitBody{}, "Steps", fmt.Sprintf("max=%d", maxSteps))
	tagHas(t, submitBody{}, "Steps", fmt.Sprintf("max=%d", maxStepBytes))
	tagHas(t, consoleLog{}, "Entries", fmt.Sprintf("max=%d", consoleSize))
	tagHas(t, consoleEntry{}, "Msg", fmt.Sprintf("max=%d", consoleMsg))
	tagHas(t, consoleEntry{}, "Page", fmt.Sprintf("max=%d", consolePage))
	tagHas(t, feedbackForm{}, "Page", fmt.Sprintf("max=%d", consolePage))
	tagHas(t, consoleEntry{}, "Level", "oneof="+strings.Join(consoleLevels, " ")+" uncaught")
}
