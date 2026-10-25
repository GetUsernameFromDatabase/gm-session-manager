// NPC field editors. A field's value is the hidden input[data-field-value]
// the form posts: plain text with [[kind:id|label]] tokens
// (views/wikilinks.go reads them). Each row shows it in a small TipTap
// editor that renders [[links]] as notes do, or, with the Raw text switch
// on, as stored in its textarea[data-field-raw]. Both write to the input.
import { Editor } from "/assets/vendor/tiptap-core.js";
import { StarterKit } from "/assets/vendor/tiptap-starter.js";
import { WikiLink, keepCaretBeside, linkClick, touchDrag } from "/assets/js/wikilink.js";

// The same token as wikiLinkPattern in web/views/wikilinks.go.
const TOKEN = /\[\[(note|entity|image):(\d+)\|([^\]|]*)(?:\|(left|right|inline|break))?\]\]/g;

// toDoc turns field text into a document: a paragraph per line, a mention
// per token.
function toDoc(text) {
  const paragraph = (line) => {
    const content = [];
    let last = 0;
    for (const m of line.matchAll(TOKEN)) {
      if (m.index > last) content.push({ type: "text", text: line.slice(last, m.index) });
      content.push({ type: "mention", attrs: { id: `${m[1]}:${m[2]}`, label: m[3].trim(), float: m[4] || null } });
      last = m.index + m[0].length;
    }
    if (last < line.length) content.push({ type: "text", text: line.slice(last) });
    return { type: "paragraph", content };
  };
  return { type: "doc", content: text.split("\n").map(paragraph) };
}

// toText is toDoc's inverse.
function toText(doc) {
  const inline = (n) => {
    if (n.type === "mention") {
      const float = n.attrs.float ? `|${n.attrs.float}` : "";
      return `[[${n.attrs.id}|${String(n.attrs.label).replace(/[\]|]/g, "")}${float}]]`;
    }
    return n.type === "hardBreak" ? "\n" : n.text || "";
  };
  return (doc.content || []).map((p) => (p.content || []).map(inline).join("")).join("\n");
}

// Plain text and links only: formatting would not survive as field text.
const extensions = [
  StarterKit.configure({
    blockquote: false,
    bold: false,
    bulletList: false,
    code: false,
    codeBlock: false,
    heading: false,
    horizontalRule: false,
    italic: false,
    link: false,
    listItem: false,
    listKeymap: false,
    orderedList: false,
    strike: false,
    underline: false,
  }),
  WikiLink,
];

const editors = new Map(); // row -> Editor
const parts = (row) => ({ value: row.querySelector("[data-field-value]"), raw: row.querySelector("[data-field-raw]") });

// setValue stores a row's text in its input and keeps the empty marker.
function setValue(row, text) {
  parts(row).value.value = text;
  row.toggleAttribute("data-empty", text.trim() === "");
}

function mount(row) {
  const box = row.querySelector("[data-field-editor]");
  box.classList.add("field-editor");
  const editor = new Editor({
    element: box,
    content: toDoc(parts(row).value.value),
    extensions,
    onUpdate: ({ editor }) => {
      const text = toText(editor.getJSON());
      parts(row).raw.value = text;
      setValue(row, text);
    },
    onSelectionUpdate: keepCaretBeside,
  });
  editors.set(row, editor);
}

// Rows come and go: added by the add-field picker, replaced by live-view
// refetches. Mount new ones; drop editors whose row left the page.
function sync() {
  document.querySelectorAll("[data-field-key]").forEach((row) => editors.has(row) || !row.querySelector("[data-field-editor]") || mount(row));
  for (const [row, editor] of editors) {
    if (!row.isConnected) {
      editor.destroy();
      editors.delete(row);
    }
  }
}
sync();
new MutationObserver(sync).observe(document.body, { childList: true, subtree: true });

// Raw text edits go to the input; the editor picks them up when shown again.
document.addEventListener("input", (e) => {
  if (e.target.matches?.("[data-field-raw]")) setValue(e.target.closest("[data-field-key]"), e.target.value);
});

// The Raw text switch flips every field of its form between editor and text.
document.addEventListener("change", (e) => {
  if (!e.target.closest?.("[data-raw-switch]") || e.target.type !== "checkbox") return;
  const form = e.target.closest("form");
  form.toggleAttribute("data-raw", e.target.checked);
  if (e.target.checked) return;
  form.querySelectorAll("[data-field-key]").forEach((row) => editors.get(row)?.commands.setContent(toDoc(parts(row).value.value), { emitUpdate: false }));
});

// A click on a [[link]] opens the same menu as in notes; Ctrl/Cmd-click follows it.
document.addEventListener("click", (e) => {
  if (!e.target.closest(".field-editor a[href]")) return;
  const row = e.target.closest("[data-field-key]");
  linkClick(e, editors.get(row));
});

touchDrag((el) => editors.get(el.closest("[data-field-key]")));
