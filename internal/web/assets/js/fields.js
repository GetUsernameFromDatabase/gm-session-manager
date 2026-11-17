// NPC field editors. A field's value is the hidden input[data-field-value]
// its row form posts: plain text with [[kind:id|label]] tokens
// (views/wikilinks.go reads them). Each row shows it in a small TipTap
// editor that renders [[links]] as notes do, or, with the Raw text switch
// on, as stored in its textarea[data-field-raw]. Both write to the input,
// and each field saves on its own: a second after typing stops, and when
// it loses focus.
import { Editor } from "/assets/vendor/tiptap-core.js";
import { StarterKit } from "/assets/vendor/tiptap-starter.js";
import { WikiLink, keepCaretBeside, linkClick, touchDrag } from "/assets/js/wikilink.js";

// The same token as wikiLinkPattern in internal/web/views/wikilinks.go.
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

const SAVE_DELAY = 1000;
const timers = new Map(); // row -> pending save

// save posts a row's text if it changed since its last save (the row form
// posts on field-save).
function save(row) {
  clearTimeout(timers.get(row));
  timers.delete(row);
  const text = parts(row).value.value;
  if (text === row.dataset.saved) return;
  row.dataset.saved = text;
  htmx.trigger(row, "field-save");
}

// A failed save is sent again at the next pause or blur.
document.addEventListener("htmx:afterRequest", (e) => {
  const row = e.detail.elt.closest?.("[data-field-key]");
  if (row && !e.detail.successful) delete row.dataset.saved;
});

// setValue stores a row's text in its input and saves it once typing stops.
function setValue(row, text) {
  parts(row).value.value = text;
  clearTimeout(timers.get(row));
  timers.set(row, setTimeout(() => save(row), SAVE_DELAY));
}

// focusKey is a field just added: its row is focused once the live view
// brings it.
let focusKey = null;
function focusRow(row) {
  row.scrollIntoView({ block: "nearest" });
  if (row.closest("[data-fields]").hasAttribute("data-raw")) parts(row).raw.focus();
  else editors.get(row)?.commands.focus("end");
}

function mount(row) {
  const box = row.querySelector("[data-field-editor]");
  box.classList.add("field-editor");
  row.dataset.saved = parts(row).value.value;
  const editor = new Editor({
    element: box,
    content: toDoc(parts(row).value.value),
    extensions,
    onUpdate: ({ editor }) => {
      const text = toText(editor.getJSON());
      parts(row).raw.value = text;
      setValue(row, text);
    },
    onBlur: () => save(row),
    onSelectionUpdate: keepCaretBeside,
  });
  editors.set(row, editor);
  if (row.dataset.fieldKey === focusKey) {
    focusKey = null;
    focusRow(row);
  }
}

// Rows come and go with live-view refetches. Mount new ones; drop editors
// whose row left the page.
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
// Editors change the DOM on every keystroke; sync at most once a frame.
let queued = false;
new MutationObserver(() => {
  if (queued) return;
  queued = true;
  requestAnimationFrame(() => {
    queued = false;
    sync();
  });
}).observe(document.body, { childList: true, subtree: true });

// Raw text edits go to the input; the editor picks them up when shown again.
document.addEventListener("input", (e) => {
  if (e.target.matches?.("[data-field-raw]")) setValue(e.target.closest("[data-field-key]"), e.target.value);
});
document.addEventListener("focusout", (e) => {
  if (e.target.matches?.("[data-field-raw]")) save(e.target.closest("[data-field-key]"));
});

// The Raw text switch flips every field between editor and text.
document.addEventListener("change", (e) => {
  if (!e.target.closest?.("[data-raw-switch]") || e.target.type !== "checkbox") return;
  const fields = e.target.closest("[data-fields]");
  fields.toggleAttribute("data-raw", e.target.checked);
  if (e.target.checked) return;
  fields.querySelectorAll("[data-field-key]").forEach((row) => editors.get(row)?.commands.setContent(toDoc(parts(row).value.value), { emitUpdate: false }));
});

// Picking a key in the add-field combobox adds the field, or goes to the
// row that has it; the picked item leaves the list.
document.addEventListener("change", (e) => {
  const picker = e.target.closest?.("[data-add-field]");
  const key = e.target.value?.trim().toLowerCase();
  if (!picker || e.target.getAttribute("aria-hidden") !== "true" || !key) return;
  picker.querySelector(`[data-slot="combobox-item"][data-templ-value="${CSS.escape(key)}"]`)?.remove();
  e.target.value = "";
  picker.querySelector("[data-combobox-create-input]").value = "";
  const row = document.querySelector(`[data-field-key="${CSS.escape(key)}"]`);
  if (row) return focusRow(row);
  focusKey = key;
  htmx.ajax("POST", picker.dataset.addField, { values: { key }, source: picker, swap: "none" });
});

// A click on a [[link]] opens the same menu as in notes; Ctrl/Cmd-click follows it.
document.addEventListener("click", (e) => {
  if (!e.target.closest(".field-editor a[href]")) return;
  const row = e.target.closest("[data-field-key]");
  linkClick(e, editors.get(row));
});

touchDrag((el) => editors.get(el.closest("[data-field-key]")));
