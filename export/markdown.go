// Package export writes a campaign archive: notes as Markdown, characters
// as JSON and attachments, never including players' private ('owner')
// blocks.
package export

import (
	"encoding/json"
	"fmt"
	"strings"
)

// pmNode is a ProseMirror JSON node.
type pmNode struct {
	Type    string         `json:"type"`
	Text    string         `json:"text"`
	Attrs   map[string]any `json:"attrs"`
	Marks   []pmMark       `json:"marks"`
	Content []pmNode       `json:"content"`
}

type pmMark struct {
	Type  string         `json:"type"`
	Attrs map[string]any `json:"attrs"`
}

// Links maps references inside a document to Markdown targets. A function
// returning "" renders a mention as plain text and drops an image.
type Links struct {
	// Mention resolves a mention id such as "note:12" or "entity:3".
	Mention func(id string) string
	// Image resolves an image src such as "/attachments/5".
	Image func(src string) string
}

// Markdown converts a ProseMirror document (TipTap StarterKit, mentions and
// images) to Markdown.
func Markdown(docJSON string, links Links) (string, error) {
	var doc pmNode
	if err := json.Unmarshal([]byte(docJSON), &doc); err != nil {
		return "", fmt.Errorf("note block: %w", err)
	}
	var b strings.Builder
	w := writer{b: &b, links: links}
	w.blocks(doc.Content, "")
	return strings.TrimRight(b.String(), "\n") + "\n", nil
}

type writer struct {
	b     *strings.Builder
	links Links
}

// blocks writes block nodes, each line prefixed (for quotes and lists).
func (w writer) blocks(nodes []pmNode, prefix string) {
	for i, n := range nodes {
		if i > 0 {
			w.line(strings.TrimRight(prefix, " "), "")
		}
		w.block(n, prefix)
	}
}

func (w writer) block(n pmNode, prefix string) {
	switch n.Type {
	case "paragraph":
		w.line(prefix, w.inline(n.Content))
	case "heading":
		level := 1
		if l, ok := n.Attrs["level"].(float64); ok {
			level = int(l)
		}
		w.line(prefix, strings.Repeat("#", level)+" "+w.inline(n.Content))
	case "blockquote":
		w.blocks(n.Content, prefix+"> ")
	case "bulletList", "orderedList":
		for i, item := range n.Content {
			marker := "- "
			if n.Type == "orderedList" {
				marker = fmt.Sprintf("%d. ", i+1)
			}
			w.listItem(item, prefix, marker)
		}
	case "codeBlock":
		w.line(prefix, "```")
		for l := range strings.SplitSeq(plain(n.Content), "\n") {
			w.line(prefix, l)
		}
		w.line(prefix, "```")
	case "horizontalRule":
		w.line(prefix, "---")
	case "image":
		w.line(prefix, w.inline([]pmNode{n}))
	default:
		if len(n.Content) > 0 {
			w.blocks(n.Content, prefix)
		}
	}
}

func (w writer) listItem(item pmNode, prefix, marker string) {
	indent := strings.Repeat(" ", len(marker))
	for i, child := range item.Content {
		p := prefix + indent
		if i == 0 {
			// The first line carries the marker; continuation lines indent.
			var sub strings.Builder
			writer{b: &sub, links: w.links}.block(child, "")
			lines := strings.Split(strings.TrimRight(sub.String(), "\n"), "\n")
			for j, l := range lines {
				if j == 0 {
					w.line(prefix, marker+l)
				} else {
					w.line(p, l)
				}
			}
			continue
		}
		w.block(child, p)
	}
}

func (w writer) line(prefix, s string) {
	w.b.WriteString(prefix)
	w.b.WriteString(s)
	w.b.WriteByte('\n')
}

func (w writer) inline(nodes []pmNode) string {
	var b strings.Builder
	for _, n := range nodes {
		switch n.Type {
		case "text":
			b.WriteString(applyMarks(escape(n.Text), n.Marks))
		case "hardBreak":
			b.WriteString("  \n")
		case "mention":
			id, _ := n.Attrs["id"].(string)
			label, _ := n.Attrs["label"].(string)
			if label == "" {
				label = id
			}
			if target := w.links.Mention(id); target != "" {
				fmt.Fprintf(&b, "[%s](%s)", escape(label), target)
			} else {
				b.WriteString(escape(label))
			}
		case "image":
			src, _ := n.Attrs["src"].(string)
			alt, _ := n.Attrs["alt"].(string)
			if target := w.links.Image(src); target != "" {
				fmt.Fprintf(&b, "![%s](%s)", escape(alt), target)
			}
		}
	}
	return b.String()
}

func applyMarks(s string, marks []pmMark) string {
	for _, m := range marks {
		switch m.Type {
		case "bold":
			s = "**" + s + "**"
		case "italic":
			s = "_" + s + "_"
		case "strike":
			s = "~~" + s + "~~"
		case "code":
			s = "`" + s + "`"
		case "link":
			if href, ok := m.Attrs["href"].(string); ok {
				s = "[" + s + "](" + href + ")"
			}
		}
	}
	return s
}

// escape protects Markdown syntax characters in text.
func escape(s string) string {
	r := strings.NewReplacer(`\`, `\\`, "*", `\*`, "_", `\_`, "`", "\\`", "[", `\[`, "]", `\]`, "#", `\#`)
	return r.Replace(s)
}

func plain(nodes []pmNode) string {
	var b strings.Builder
	for _, n := range nodes {
		b.WriteString(n.Text)
		b.WriteString(plain(n.Content))
	}
	return b.String()
}
