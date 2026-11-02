package views

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark/v2/extension"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/renderer/html"
)

// Compendium text (SRD and homebrew) is GFM; hard wraps keep single newlines as line breaks.
// The renderer stays safe (no WithUnsafe): raw HTML and javascript: links are dropped.
var (
	mdParser   = parser.New(parser.WithExtensions(extension.GFMParser))
	mdRenderer = html.New(html.WithHardWraps(), html.WithExtensions(extension.GFMHTMLRenderer))
)

// renderMarkdown converts Markdown to sanitized HTML.
func renderMarkdown(s string) (string, error) {
	src := []byte(s)
	var buf bytes.Buffer
	err := mdRenderer.Render(&buf, src, mdParser.Parse(src))
	return buf.String(), err
}

// markdown renders s as styled Markdown.
func markdown(s string) templ.Component {
	return templ.ComponentFunc(func(_ context.Context, w io.Writer) error {
		out, err := renderMarkdown(s)
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, `<div class="markdown">`+out+`</div>`)
		return err
	})
}
