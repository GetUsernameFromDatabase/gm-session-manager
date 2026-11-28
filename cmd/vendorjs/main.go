// Command vendorjs downloads the browser JavaScript the app serves (htmx, its
// websocket extension, SortableJS, TipTap and ProseMirror) from jsDelivr into
// internal/web/assets/vendor, so the binary embeds it and never needs a CDN or npm at
// runtime. jsDelivr's +esm modules import their dependencies by absolute
// /npm/... paths; those are followed and rewritten to local paths.
//
// Run from the repository root after changing a version below:
//
//	go run ./cmd/vendorjs
//
// To list the pinned packages that have newer releases (downloads nothing):
//
//	go run ./cmd/vendorjs -check
//
// Every file must match its SHA-256 in vendor.sum (committed), so a build never
// ships JS that changed on the CDN unseen. After a version change, or when
// jsDelivr rebuilt its +esm files, look at the difference and accept it:
//
//	go run ./cmd/vendorjs -update
package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"maps"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	cdnBase    = "https://cdn.jsdelivr.net"
	dataBase   = "https://data.jsdelivr.com/v1/packages/npm/"
	outDir     = "internal/web/assets/vendor"
	publicDir  = "/assets/vendor"
	sumFile    = "cmd/vendorjs/vendor.sum"
	scratchDir = "build"
)

// scripts are classic (non-module) files saved as-is under the given name.
var scripts = map[string]string{
	"htmx.min.js": "/npm/htmx.org@2.0.11/dist/htmx.min.js",
	"htmx-ws.js":  "/npm/htmx-ext-ws@2.0.4/dist/ws.min.js",
	"sortable.js": "/npm/sortablejs@1.15.7/Sortable.min.js",
}

// modules are ES module entry points; everything they import is fetched too.
// Each entry also gets a stable alias module (e.g. tiptap-core.js) that
// re-exports it, so app code never names a version.
var modules = map[string]string{
	"tiptap-core.js":    "/npm/@tiptap/core@3.32.0/+esm",
	"tiptap-starter.js": "/npm/@tiptap/starter-kit@3.32.0/+esm",
	"tiptap-mention.js": "/npm/@tiptap/extension-mention@3.32.0/+esm",
	"tiptap-image.js":   "/npm/@tiptap/extension-image@3.32.0/+esm",
	"pm-collab.js":      "/npm/prosemirror-collab@1.3.1/+esm",
	"pm-transform.js":   "/npm/@tiptap/pm@3.32.0/transform/+esm",
	"pm-state.js":       "/npm/@tiptap/pm@3.32.0/state/+esm",
}

// npmImport matches the absolute specifiers jsDelivr writes into +esm files.
var npmImport = regexp.MustCompile(`["'](/npm/[^"'\s]+?/\+esm)["']`)

func main() {
	checkOnly := flag.Bool(
		"check",
		false,
		"list pinned packages with newer releases instead of downloading",
	)
	update := flag.Bool(
		"update",
		false,
		"accept downloads that differ from "+sumFile+" and rewrite it",
	)
	flag.Parse()
	do := func(ctx context.Context) error { return run(ctx, *update) }
	if *checkOnly {
		do = check
	}
	if err := do(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, "vendorjs:", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, update bool) error {
	files := map[string][]byte{} // by slash path under outDir
	for name, src := range scripts {
		body, err := fetch(ctx, src)
		if err != nil {
			return err
		}
		files[name] = body
	}

	// jsDelivr resolves each package's dependency ranges separately, so one
	// crawl can pull several copies of prosemirror-model and friends. Mixed
	// copies break ProseMirror's instanceof checks, so every import of a
	// package is pointed at the highest version seen, re-crawling until the
	// choice is stable.
	cache := map[string][]byte{}
	canon := map[string]string{}
	var mods map[string][]byte
	for {
		var err error
		if mods, err = crawl(ctx, cache, canon); err != nil {
			return err
		}
		next := highestVersions(mods)
		if maps.Equal(next, canon) {
			break
		}
		canon = next
	}

	for src, body := range mods {
		rewritten := npmImport.ReplaceAllStringFunc(string(body), func(s string) string {
			quote := s[:1]
			return quote + publicPath(canonical(s[1:len(s)-1], canon)) + quote
		})
		files[localPath(src)] = []byte(rewritten)
	}
	for alias, src := range modules {
		files[alias] = fmt.Appendf(
			nil,
			"export * from %q;\n",
			"./"+localPath(canonical(src, canon)),
		)
	}
	if err := checkSums(files, update); err != nil {
		return err
	}
	if err := install(files); err != nil {
		return err
	}
	fmt.Printf("vendored %d scripts and %d modules into %s\n", len(scripts), len(mods), outDir)
	return nil
}

// checkSums compares the files with vendor.sum ("<sha256>  <path>" lines, as
// sha256sum writes them). With update it writes the file instead; a missing
// vendor.sum is an error, so an uncommitted one never trusts every download.
func checkSums(files map[string][]byte, update bool) error {
	got := map[string]string{}
	var b strings.Builder
	for _, p := range slices.Sorted(maps.Keys(files)) {
		got[p] = fmt.Sprintf("%x", sha256.Sum256(files[p]))
		fmt.Fprintf(&b, "%s  %s\n", got[p], p)
	}
	if update {
		fmt.Printf("wrote %s: review the change and commit it\n", sumFile)
		return os.WriteFile(sumFile, []byte(b.String()), 0o600)
	}
	raw, err := os.ReadFile(sumFile)
	if err != nil {
		return fmt.Errorf("%w (create it with go run ./cmd/vendorjs -update)", err)
	}
	if changed := changedPaths(got, string(raw)); len(changed) > 0 {
		return fmt.Errorf(
			"downloads differ from %s (nothing written): %s\nIf that is expected (a version change, or jsDelivr rebuilt its +esm files), check it and run go run ./cmd/vendorjs -update",
			sumFile,
			strings.Join(changed, ", "),
		)
	}
	return nil
}

// changedPaths returns the sorted paths whose sum in got differs from sumText,
// is new, or is gone; nil when all match.
func changedPaths(got map[string]string, sumText string) []string {
	want := map[string]string{}
	for line := range strings.Lines(sumText) { // git may check it out with CRLF
		if f := strings.Fields(line); len(f) == 2 {
			want[f[1]] = f[0]
		}
	}
	paths := slices.Collect(maps.Keys(got))
	for p := range want {
		if _, ok := got[p]; !ok {
			paths = append(paths, p) // gone from the download
		}
	}
	slices.Sort(paths)
	var changed []string
	for _, p := range paths {
		if got[p] != want[p] {
			changed = append(changed, p)
		}
	}
	return changed
}

// install replaces outDir with files at once:
// a failed write or swap leaves the old folder in place.
func install(files map[string][]byte) error {
	// In the gitignored build/: a leftover from a killed run is never embedded, linted or watched;
	// same volume for the rename.
	if err := os.MkdirAll(scratchDir, 0o750); err != nil {
		return err
	}
	tmp, err := os.MkdirTemp(scratchDir, ".vendor-")
	if err != nil {
		return err
	}
	defer func() { _ = os.RemoveAll(tmp) }() // after the rename there is nothing left to remove
	for p, body := range files {
		if err := write(filepath.Join(tmp, filepath.FromSlash(p)), body); err != nil {
			return err
		}
	}
	// The old folder moves aside rather than being deleted, so a swap that
	// fails (a locked file on Windows) can put it back.
	aside, err := os.MkdirTemp(scratchDir, ".old-vendor-")
	if err != nil {
		return err
	}
	keep := false // set when the old folder could not go back: it is then the only copy
	defer func() {
		if !keep {
			_ = os.RemoveAll(aside)
		}
	}()
	old := filepath.Join(aside, "vendor")
	if err := os.Rename(outDir, old); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	if err := os.Rename(tmp, outDir); err != nil {
		if rerr := os.Rename(old, outDir); rerr != nil && !errors.Is(rerr, fs.ErrNotExist) {
			keep = true
			return errors.Join(err, rerr, fmt.Errorf("the old %s is kept in %s", outDir, old))
		}
		return err
	}
	return nil
}

// check compares each pinned package with its latest release on npm. Only
// the pins are checked: their dependencies follow from them.
func check(ctx context.Context) error {
	pinned := map[string]string{}
	for _, src := range slices.Concat(slices.Collect(maps.Values(scripts)), slices.Collect(maps.Values(modules))) {
		name, version, _ := splitSpec(src)
		pinned[name] = version
	}
	outdated := 0
	for _, name := range slices.Sorted(maps.Keys(pinned)) {
		body, err := fetchURL(ctx, dataBase+name+"/resolved?specifier=latest")
		if err != nil {
			return err
		}
		var latest struct{ Version string }
		if err := json.Unmarshal(body, &latest); err != nil {
			return fmt.Errorf("latest %s: %w", name, err)
		}
		if versionLess(pinned[name], latest.Version) {
			outdated++
			fmt.Printf("%-32s %s -> %s\n", name, pinned[name], latest.Version)
		} else {
			fmt.Printf("%-32s %s (latest)\n", name, pinned[name])
		}
	}
	if outdated > 0 {
		fmt.Println(
			"\nTo update: edit the versions in cmd/vendorjs/main.go (scripts, modules; the @tiptap packages move together),",
		)
		fmt.Println("then run go run ./cmd/vendorjs -update and go test ./internal/web.")
	}
	return nil
}

// crawl fetches the module graph from the entry points, mapping every
// specifier through canon first. cache avoids refetching across rounds.
func crawl(
	ctx context.Context,
	cache map[string][]byte,
	canon map[string]string,
) (map[string][]byte, error) {
	mods := map[string][]byte{}
	queue := slices.Collect(maps.Values(modules))
	for len(queue) > 0 {
		src := canonical(queue[0], canon)
		queue = queue[1:]
		if _, ok := mods[src]; ok {
			continue
		}
		body, ok := cache[src]
		if !ok {
			var err error
			if body, err = fetch(ctx, src); err != nil {
				return nil, err
			}
			cache[src] = body
		}
		mods[src] = body
		for _, m := range npmImport.FindAllStringSubmatch(string(body), -1) {
			queue = append(queue, m[1])
		}
	}
	return mods, nil
}

// splitSpec splits "/npm/@scope/name@1.2.3/sub/+esm" into the package name,
// version and the rest of the path.
func splitSpec(src string) (name, version, rest string) {
	p := strings.TrimPrefix(src, "/npm/")
	scope := ""
	if strings.HasPrefix(p, "@") {
		s, after, _ := strings.Cut(p, "/")
		scope, p = s+"/", after
	}
	pkg, rest, _ := strings.Cut(p, "/")
	name, version, _ = strings.Cut(pkg, "@")
	return scope + name, version, rest
}

func canonical(src string, canon map[string]string) string {
	name, version, rest := splitSpec(src)
	if v, ok := canon[name]; ok {
		version = v
	}
	return "/npm/" + name + "@" + version + "/" + rest
}

func highestVersions(mods map[string][]byte) map[string]string {
	best := map[string]string{}
	for src := range mods {
		name, version, _ := splitSpec(src)
		if cur, ok := best[name]; !ok || versionLess(cur, version) {
			best[name] = version
		}
	}
	return best
}

// versionLess compares dotted numeric versions. Unresolved ranges such as
// "%5E2.0.0" sort below any concrete version.
func versionLess(a, b string) bool {
	pa, pb := versionParts(a), versionParts(b)
	if pa == nil || pb == nil {
		return pa == nil && pb != nil
	}
	return slices.Compare(pa, pb) < 0
}

func versionParts(v string) []int {
	var parts []int
	for f := range strings.SplitSeq(v, ".") {
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil
		}
		parts = append(parts, n)
	}
	return parts
}

// localPath maps "/npm/@tiptap/pm@3.32.0/state/+esm" to
// "@tiptap/pm@3.32.0/state/index.js".
func localPath(src string) string {
	return path.Join(strings.TrimSuffix(strings.TrimPrefix(src, "/npm/"), "/+esm"), "index.js")
}

func publicPath(src string) string {
	return publicDir + "/" + localPath(src)
}

func fetch(ctx context.Context, src string) ([]byte, error) {
	return fetchURL(ctx, cdnBase+src)
}

// fetchTimeout covers one CDN download.
const fetchTimeout = 30 * time.Second

func fetchURL(ctx context.Context, url string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, fetchTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %w", url, err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %s: %s", url, resp.Status)
	}
	return io.ReadAll(resp.Body)
}

func write(p string, body []byte) error {
	if err := os.MkdirAll(filepath.Dir(p), 0o750); err != nil {
		return err
	}
	return os.WriteFile(p, body, 0o600)
}
