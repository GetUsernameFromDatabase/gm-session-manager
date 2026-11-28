package web

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// textString finds quoted and backtick strings in a .templ or .js file;
// jsComment is a // comment in JS (not the // of "https://").
var (
	textString = regexp.MustCompile("\"(?:[^\"\\\\\\n]|\\\\.)*\"|'(?:[^'\\\\\\n]|\\\\.)*'|`[^`]*`")
	jsComment  = regexp.MustCompile(`(?m)(^|[ \t])//.*$`)
)

// TestNoRouteLiterals: every route path comes from internal/route, never a
// string literal, in Go, templ and our JS files (comments are not strings),
// also inside a longer string (an HTML attribute, a Markdown link).
func TestNoRouteLiterals(t *testing.T) {
	// roots are first segments (/notes); parts the static segments after
	// them (/rename), which a literal appends to a built path.
	roots, parts := map[string]bool{}, map[string]bool{}
	for _, r := range newTestEnv(t).router.Routes() {
		segs := strings.Split(strings.TrimPrefix(r.Path, "/"), "/")
		if segs[0] == "" || segs[0] == "assets" ||
			strings.Contains(segs[0], ".") { // /join.png is under /join
			continue
		}
		roots[segs[0]] = true
		for _, s := range segs[1:] {
			if !strings.HasPrefix(s, ":") && !strings.HasPrefix(s, "*") {
				parts[s] = true
			}
		}
	}
	// startsWith reports whether rest begins with segment seg: "notes" and
	// "notes/1" do, "notesy" does not.
	startsWith := func(rest, seg string) bool {
		after, ok := strings.CutPrefix(rest, seg)
		return ok &&
			(after == "" || !strings.ContainsRune("_-", rune(after[0])) && !unicode.IsLetter(rune(after[0])) && !unicode.IsDigit(rune(after[0])))
	}
	// routeIn is the route segment a path in lit starts with ("" for none): a
	// root at the start of lit or after a quote, =, ( or space
	// ("../attachments/x" and "https://host/notes" are not route paths), or a
	// part at the start of lit.
	routeIn := func(lit string) string {
		for i := range len(lit) {
			if lit[i] != '/' || i > 0 && !strings.ContainsRune("\"'=( ", rune(lit[i-1])) {
				continue
			}
			for root := range roots {
				if startsWith(lit[i+1:], root) {
					return root
				}
			}
			if i > 0 {
				continue
			}
			for part := range parts {
				if startsWith(lit[1:], part) {
					return part
				}
			}
		}
		return ""
	}

	skipDirs := []string{"web/components", "web/utils", "web/assets/vendor", "db/sqlc", "route"}
	skipFiles := []string{"route_literals_test.go"}
	found := map[string][]string{}
	add := func(file string, line int, lit string) {
		if seg := routeIn(lit); seg != "" {
			found[seg] = append(found[seg], fmt.Sprintf("%s:%d: %s", file, line, lit))
		}
	}
	err := filepath.WalkDir("..", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel("..", p)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if d.IsDir() {
			if slices.Contains(skipDirs, rel) {
				return filepath.SkipDir
			}
			return nil
		}
		// shadcn-templ-*.js is the generated component bundle.
		if slices.Contains(skipFiles, d.Name()) || strings.HasSuffix(p, "_templ.go") ||
			strings.HasPrefix(d.Name(), "shadcn-templ-") {
			return nil
		}
		ext := filepath.Ext(p)
		if ext != ".go" && ext != ".templ" && ext != ".js" {
			return nil
		}
		src, err := os.ReadFile(p)
		if err != nil {
			return err
		}
		if ext == ".js" {
			src = jsComment.ReplaceAll(src, nil) // keeps the line breaks
		}
		if ext != ".go" {
			for _, m := range textString.FindAllIndex(src, -1) {
				lit := string(src[m[0]+1 : m[1]-1])
				add(rel, 1+strings.Count(string(src[:m[0]]), "\n"), lit)
			}
			return nil
		}
		fset := token.NewFileSet()
		file := fset.AddFile(p, -1, len(src))
		var s scanner.Scanner
		s.Init(file, src, nil, 0)
		for {
			pos, tok, lit := s.Scan()
			if tok == token.EOF {
				return nil
			}
			if tok == token.STRING {
				if v, err := strconv.Unquote(lit); err == nil {
					add(rel, file.Line(pos), v)
				}
			}
		}
	})
	if err != nil {
		t.Fatal(err)
	}

	for _, root := range slices.Sorted(maps.Keys(found)) {
		t.Logf("/%s: %d literals", root, len(found[root]))
		for _, l := range found[root] {
			t.Errorf("route literal, use internal/route: %s", l)
		}
	}
}
