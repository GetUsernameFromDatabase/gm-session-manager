// [[links]] for TipTap editors: note blocks and NPC fields.
import { Extension } from "/assets/vendor/tiptap-core.js";
import { Mention } from "/assets/vendor/tiptap-mention.js";
import { Plugin } from "/assets/vendor/pm-state.js";

const campaign = document.body.dataset.campaign;

// Link targets by mention kind; "img:" and the like narrow the [[ picker
// on the server. An "attachment" is a picture uploaded into a note block.
export const LINK_PATHS = { note: "/notes/", entity: "/entities/", image: "/images/", attachment: "/attachments/" };

const linkHref = (kind, id) => (LINK_PATHS[kind] || LINK_PATHS.entity) + id;

// imageSrc is the file an [[image]] or attached picture shows; "" for other links.
const imageSrc = (kind, id) => ({ image: `/images/${id}/file`, attachment: `/attachments/${id}` })[kind] || "";

const isGM = () => "gm" in document.body.dataset;

// LAYOUTS are an image's layouts, as Google Docs names them: [label, float, icon].
const LAYOUTS = [
  ["In line", "inline", "type"],
  ["Wrap text left", "left", "panel-left"],
  ["Wrap text right", "right", "panel-right"],
  ["Break text", "break", "separator-horizontal"],
];

// [[ opens a picker of notes, NPCs and (for the GM) images.
export const WikiLink = Mention.extend({
  // Selectable so a click selects the whole [[link]] (mentionActions).
  selectable: true,
  // Draggable so a press drags it at once.
  draggable: true,
  // An image's layout (LAYOUTS); null is "left", as it was before layouts had names.
  addAttributes() {
    return { ...this.parent?.(), float: { default: null, renderHTML: () => ({}) } };
  },
  // Parse our own <a data-mention> back into a mention on paste; priority beats the Link mark's a[href].
  parseHTML: () => [
    {
      tag: "a[data-mention]",
      priority: 100,
      getAttrs: (el) => ({ id: el.dataset.mention, label: el.textContent, float: el.dataset.float === "left" ? null : el.dataset.float || null }),
    },
  ],
}).configure({
  HTMLAttributes: { class: "wikilink" },
  renderText: ({ node }) => `[[${node.attrs.label}]]`,
  renderHTML: ({ node }) => {
    const [kind, id] = String(node.attrs.id).split(":");
    const attrs = { href: linkHref(kind, id), class: "wikilink", "data-mention": node.attrs.id };
    const src = imageSrc(kind, id);
    if (!src) return ["a", attrs, node.attrs.label];
    attrs["data-float"] = node.attrs.float || "left";
    return ["a", attrs, ["img", { src, alt: node.attrs.label, loading: "lazy" }], ["span", {}, node.attrs.label]];
  },
  suggestion: {
    char: "[[",
    allowSpaces: true,
    // [[ opens the picker right after a word too, as Change… leaves it.
    allowedPrefixes: null,
    items: async ({ query }) => {
      const res = await fetch(`/campaigns/${campaign}/links?q=${encodeURIComponent(query)}`);
      return res.ok ? res.json() : [];
    },
    command: ({ editor, range, props }) => {
      editor
        .chain()
        .focus()
        .insertContentAt(range, [
          // A new [[image]] sits on the right with text on its left.
          { type: "mention", attrs: { id: props.id, label: props.label, float: props.id.startsWith("image:") ? "right" : null } },
          { type: "text", text: " " },
        ])
        .run();
    },
    render: window.gsm.suggestionMenu,
  },
});

// KEEP_IMAGES set to false on a transaction lets it delete images.
export const KEEP_IMAGES = "keepImages";

// isImage tells an [[image]], an attached picture or an older attached picture node.
const isImage = (n) => n.type.name === "image" || (n.type.name === "mention" && !!imageSrc(...String(n.attrs.id).split(":")));

function imagesIn(doc, from = 0, to = doc.content.size) {
  const out = [];
  doc.nodesBetween(from, to, (n) => {
    if (isImage(n)) out.push(n);
  });
  return out;
}

// KeepImages stops Backspace, Delete, cut and typing from deleting images
// in note blocks: Backspace or Delete over a selection keeps its images.
// Backspace or Delete on a selected image, the menu's Remove and Change…,
// undo and other people's edits still delete them.
export const KeepImages = Extension.create({
  name: "keepImages",
  addProseMirrorPlugins: () => [
    new Plugin({
      filterTransaction: (tr, state) =>
        !tr.docChanged ||
        tr.getMeta(KEEP_IMAGES) === false ||
        !!tr.getMeta("collab$") ||
        !!tr.getMeta("history$") ||
        imagesIn(tr.doc).length >= imagesIn(state.doc).length,
      props: {
        handleKeyDown(view, e) {
          if (e.key !== "Backspace" && e.key !== "Delete") return false;
          const { state } = view;
          const sel = state.selection;
          if (sel.node && isImage(sel.node)) {
            view.dispatch(state.tr.deleteSelection().setMeta(KEEP_IMAGES, false));
            return true;
          }
          const images = sel.empty ? [] : imagesIn(state.doc, sel.from, sel.to);
          if (!images.length) return false;
          try {
            view.dispatch(state.tr.replaceWith(sel.from, sel.to, images));
          } catch {
            // The images do not fit where the selection was (an older picture among text): delete nothing.
          }
          return true;
        },
      },
    }),
  ],
});

// What the last press hit, so a click beside an [[image]] does not select it.
let pressed = null;
document.addEventListener("pointerdown", (e) => (pressed = e.target), true);
document.addEventListener("keydown", () => (pressed = null), true);

// keepCaretBeside is an onSelectionUpdate handler: clicking beside a tall
// [[image]] can land the caret inside it; keep it after the image instead.
export function keepCaretBeside({ editor }) {
  const sel = editor.state.selection;
  if (sel.node?.type.name === "mention" && pressed && !pressed.closest(".wikilink")) editor.commands.setTextSelection(sel.to);
}

// An [[image]] the viewer may not see (or that is gone) falls back to its name.
document.addEventListener("error", (e) => e.target.matches?.(".wikilink img") && (e.target.hidden = true), true);

let openMenu = null;
function closeMenu() {
  openMenu?.remove();
  openMenu = null;
}
document.addEventListener("click", (e) => openMenu && !openMenu.contains(e.target) && closeMenu());
document.addEventListener("keydown", (e) => e.key === "Escape" && closeMenu());

// menuIcon clones a lucide icon the layout renders into #menu-icons.
const menuIcon = (name) => document.getElementById("menu-icons")?.content.querySelector(`.lucide-${name}`)?.cloneNode(true);

// showMenu lists actions ([label, fn, icon, destructive]) below a viewport
// rect (left, top, bottom), or above it when there is no room below, kept
// inside the viewport's width; a null action is a separator.
export function showMenu(rect, actions) {
  closeMenu();
  openMenu = document.createElement("div");
  openMenu.className = "absolute z-50 flex min-w-40 flex-col rounded-md border bg-popover p-1 text-popover-foreground shadow";
  for (const action of actions) {
    if (!action) {
      openMenu.append(Object.assign(document.createElement("div"), { className: "-mx-1 my-1 h-px bg-border" }));
      continue;
    }
    const [label, fn, icon, destructive] = action;
    const b = document.createElement("button");
    b.type = "button";
    b.className = destructive
      ? "flex items-center gap-2 rounded px-2 py-1.5 text-left text-sm text-destructive hover:bg-destructive/10 dark:hover:bg-destructive/20 [&_svg]:size-4"
      : "flex items-center gap-2 rounded px-2 py-1.5 text-left text-sm hover:bg-accent [&_svg]:size-4 [&_svg]:text-muted-foreground";
    b.append(...[menuIcon(icon), label].filter(Boolean));
    b.addEventListener("mousedown", (e) => {
      e.preventDefault();
      closeMenu();
      fn();
    });
    openMenu.append(b);
  }
  document.body.append(openMenu);
  const { width, height } = openMenu.getBoundingClientRect();
  const left = Math.max(8, Math.min(rect.left, window.innerWidth - width - 8));
  const above = rect.bottom + height + 4 > window.innerHeight && rect.top - height - 4 >= 0;
  openMenu.style.left = `${left + window.scrollX}px`;
  openMenu.style.top = `${(above ? rect.top - height - 4 : rect.bottom + 4) + window.scrollY}px`;
}

// pickFile asks for one file and calls done with it.
export function pickFile(accept, done) {
  const input = document.createElement("input");
  input.type = "file";
  input.accept = accept;
  input.addEventListener("change", () => input.files.length && done(input.files[0]));
  input.click();
}

// mentionActions shows the menu for a clicked [[link]] in an editable
// editor (mentionMenu). It reports whether the click was on a [[link]].
export function mentionActions(editor, target, e, opts) {
  if (!target.closest("a[data-mention]")) return false;
  const view = editor.view;
  // A click selects a [[link]] first, which may resize it away from the click point.
  const sel = view.state.selection;
  const pos = sel.node?.type.name === "mention" ? sel.from : view.posAtCoords({ left: e.clientX, top: e.clientY })?.inside;
  if (!(pos >= 0) || view.state.doc.nodeAt(pos)?.type.name !== "mention") return false;
  mentionMenu(editor, pos, { left: e.clientX, top: e.clientY, bottom: e.clientY }, opts);
  return true;
}

// mentionMenu shows what can be done to the [[link]] at pos: Open; for
// images Show on display (the GM's) and the other layouts; for attached
// pictures adding them to the campaign's images (the GM) and Replace…
// (given upload, which stores a file on the block); Change… for other
// links; Remove.
export function mentionMenu(editor, pos, point, { upload } = {}) {
  const view = editor.view;
  const node = view.state.doc.nodeAt(pos);
  const chain = () => editor.chain().focus();
  const setAttrs = (attrs) =>
    chain()
      .command(({ tr }) => {
        for (const [k, v] of Object.entries(attrs)) tr.setNodeAttribute(pos, k, v);
        return true;
      })
      .run();
  const [kind, id] = String(node.attrs.id).split(":");
  const src = imageSrc(kind, id);
  const actions = [["Open", () => (window.location.href = linkHref(kind, id)), "external-link"]];
  if (src) {
    if (isGM()) actions.push(["Show on display", () => gsm.showImage(src, node.attrs.label), "cast"]);
    const current = node.attrs.float ?? "left";
    actions.push(null, ...LAYOUTS.filter(([, v]) => v !== current).map(([label, v, icon]) => [label, () => setAttrs({ float: v }), icon]));
  }
  actions.push(null);
  if (kind !== "attachment") {
    // Typing [[ again opens the picker on the old label to choose anew.
    const change = () => chain().setMeta(KEEP_IMAGES, false).deleteRange({ from: pos, to: pos + node.nodeSize }).insertContentAt(pos, `[[${node.attrs.label}`).run();
    actions.push(["Change…", change, "pencil"]);
  } else {
    if (isGM()) actions.push(["Add to campaign images", () => toCampaignImage(id, setAttrs), "image-plus"]);
    if (upload) {
      const replace = (file) => upload(file).then((data) => data?.image && setAttrs({ id: `attachment:${data.id}`, label: data.name }));
      actions.push(["Replace…", () => pickFile("image/*", replace), "pencil"]);
    }
  }
  // Removing only takes the picture out of this note; unused files are deleted from the Files page.
  actions.push(["Remove", () => chain().setMeta(KEEP_IMAGES, false).deleteRange({ from: pos, to: pos + node.nodeSize }).run(), "trash-2", true]);
  // Images can be large, so open at the click rather than below the image.
  showMenu(src ? point : view.nodeDOM(pos).getBoundingClientRect(), actions);
}

// toCampaignImage copies an attached picture into the campaign's images and
// makes this use an [[image]] link to it, keeping its layout.
async function toCampaignImage(id, setAttrs) {
  const res = await fetch(`/attachments/${id}/to-image`, { method: "POST" });
  if (!res.ok) return alert("The picture could not be added to the campaign's images.");
  const img = await res.json();
  setAttrs({ id: `image:${img.id}`, label: img.name });
}

// linkClick handles a click on a link in an editor: Ctrl/Cmd-click follows
// it; in an editable editor a plain click opens its menu (menu, or
// mentionActions by default) instead of following it, since [[links]] are
// non-editable islands the browser would otherwise follow.
export function linkClick(e, editor, menu) {
  const a = e.target.closest("a[href]");
  if (a && (e.ctrlKey || e.metaKey)) {
    e.preventDefault();
    window.location.href = a.getAttribute("href");
  } else if (editor?.isEditable) {
    e.preventDefault();
    (menu || (() => mentionActions(editor, e.target, e)))();
  }
}

const HOLD_MS = 400;
const SLOP_PX = 10;
const EDGE_PX = 56; // a finger this close to the top or bottom scrolls the page
const MAX_SCROLL_PX = 16; // per frame, at the very edge

// touchDrag lets a finger move an image in an editable editor, since phones
// have no drag and drop there: hold it still for a moment, then drag; a
// caret shows where it lands, and the page scrolls near the screen's edges.
// editorOf finds the editor holding an element.
export function touchDrag(editorOf) {
  let hold = null; // {x, y, timer} while waiting out the hold
  let drag = null; // {editor, pos, node, link, caret, drop, x, y, frame} while dragging
  // place moves the caret to the drop spot under the finger.
  const place = () => {
    const view = drag.editor.view;
    drag.drop = view.posAtCoords({ left: drag.x, top: drag.y })?.pos ?? null;
    if (drag.drop === null) return drag.caret.classList.add("hidden");
    const c = view.coordsAtPos(drag.drop);
    Object.assign(drag.caret.style, { left: `${c.left + window.scrollX}px`, top: `${c.top + window.scrollY}px`, height: `${c.bottom - c.top}px` });
    drag.caret.classList.remove("hidden");
  };
  const autoScroll = () => {
    if (!drag) return;
    const fromBottom = window.innerHeight - drag.y;
    const depth = drag.y < EDGE_PX ? drag.y - EDGE_PX : fromBottom < EDGE_PX ? EDGE_PX - fromBottom : 0;
    if (depth) {
      window.scrollBy(0, (depth / EDGE_PX) * MAX_SCROLL_PX);
      place();
    }
    drag.frame = requestAnimationFrame(autoScroll);
  };
  const editableImage = (el) => {
    const img = el.closest?.(".tiptap .wikilink img");
    const editor = img && editorOf(img);
    return editor?.isEditable ? [img, editor] : [];
  };
  const stopHold = () => {
    clearTimeout(hold?.timer);
    hold = null;
  };
  const stopDrag = () => {
    if (!drag) return;
    cancelAnimationFrame(drag.frame);
    drag.caret.remove();
    delete drag.link.dataset.dragging;
    drag = null;
  };
  const start = (editor, img, x, y) => {
    hold = null;
    const view = editor.view;
    const link = img.closest(".wikilink");
    let pos = view.posAtDOM(link, 0);
    if (view.state.doc.nodeAt(pos)?.type.name !== "mention") pos -= 1;
    const node = view.state.doc.nodeAt(pos);
    if (node?.type.name !== "mention") return;
    const caret = Object.assign(document.createElement("div"), { className: "pointer-events-none absolute z-50 hidden w-0.5 bg-primary" });
    document.body.append(caret);
    link.dataset.dragging = "";
    navigator.vibrate?.(20);
    drag = { editor, pos, node, link, caret, drop: null, x, y };
    autoScroll();
  };
  // The phone's own long-press menu (open in new tab, save image) would cover the drag.
  document.addEventListener("contextmenu", (e) => (hold || drag) && editableImage(e.target).length && e.preventDefault());
  document.addEventListener(
    "touchstart",
    (e) => {
      const [img, editor] = editableImage(e.target);
      if (!img || e.touches.length !== 1) return;
      const t = e.touches[0];
      hold = { x: t.clientX, y: t.clientY, timer: setTimeout(() => start(editor, img, t.clientX, t.clientY), HOLD_MS) };
    },
    { passive: true },
  );
  document.addEventListener(
    "touchmove",
    (e) => {
      const t = e.touches[0];
      if (hold && Math.hypot(t.clientX - hold.x, t.clientY - hold.y) > SLOP_PX) stopHold();
      if (!drag) return;
      e.preventDefault(); // no scrolling by hand while dragging; autoScroll does it
      drag.x = t.clientX;
      drag.y = t.clientY;
      place();
    },
    { passive: false },
  );
  document.addEventListener(
    "touchend",
    (e) => {
      stopHold();
      if (!drag) return;
      e.preventDefault(); // no click, so the menu stays shut
      const { editor, pos, node, drop } = drag;
      stopDrag();
      if (drop === null || (drop >= pos && drop <= pos + node.nodeSize)) return;
      // One transaction deletes and reinserts, so KeepImages sees no image lost.
      const tr = editor.state.tr.delete(pos, pos + node.nodeSize);
      const at = tr.mapping.map(drop);
      if (tr.doc.resolve(at).parent.inlineContent) editor.view.dispatch(tr.insert(at, node));
    },
    { passive: false },
  );
  document.addEventListener("touchcancel", () => {
    stopHold();
    stopDrag();
  });
}
