// [[links]] for TipTap editors: note blocks and NPC fields.
import { Mention } from "/assets/vendor/tiptap-mention.js";

const campaign = document.body.dataset.campaign;

// Link targets by mention kind; "img:" and the like narrow the [[ picker
// on the server.
export const LINK_PATHS = { note: "/notes/", entity: "/entities/", image: "/images/" };

// [[ opens a picker of notes, NPCs and (for the GM) images.
export const WikiLink = Mention.extend({
  // Selectable so a click selects the whole [[link]] (mentionActions).
  selectable: true,
  // An [[image]] floats left (null), right or inline; text wraps around a floated one.
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
    const href = (LINK_PATHS[kind] || LINK_PATHS.entity) + id;
    const attrs = { href, class: "wikilink", "data-mention": node.attrs.id };
    if (kind !== "image") return ["a", attrs, node.attrs.label];
    attrs["data-float"] = node.attrs.float || "left";
    return ["a", attrs, ["img", { src: `/images/${id}/file`, alt: node.attrs.label, loading: "lazy" }], ["span", {}, node.attrs.label]];
  },
  suggestion: {
    char: "[[",
    allowSpaces: true,
    items: async ({ query }) => {
      const res = await fetch(`/campaigns/${campaign}/links?q=${encodeURIComponent(query)}`);
      return res.ok ? res.json() : [];
    },
    command: ({ editor, range, props }) => {
      editor
        .chain()
        .focus()
        .insertContentAt(range, [
          { type: "mention", attrs: { id: props.id, label: props.label } },
          { type: "text", text: " " },
        ])
        .run();
    },
    render: window.gsm.suggestionMenu,
  },
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
// rect (left, bottom); a null action is a separator.
export function showMenu(rect, actions) {
  closeMenu();
  openMenu = document.createElement("div");
  openMenu.className = "absolute z-50 flex min-w-40 flex-col rounded-md border bg-popover p-1 text-popover-foreground shadow";
  openMenu.style.left = `${rect.left + window.scrollX}px`;
  openMenu.style.top = `${rect.bottom + window.scrollY + 4}px`;
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
}

// mentionActions shows the menu for a clicked [[link]] in an editable
// editor: Open and Show on display (the GM's images), an image's float,
// then Change… and Remove. It reports whether the click was on a [[link]].
export function mentionActions(editor, target, e) {
  const mention = target.closest("a[data-mention]");
  if (!mention) return false;
  const view = editor.view;
  // A click selects a [[link]] first, which may resize it away from the click point.
  const sel = view.state.selection;
  const pos = sel.node?.type.name === "mention" ? sel.from : view.posAtCoords({ left: e.clientX, top: e.clientY })?.inside;
  const node = pos >= 0 ? view.state.doc.nodeAt(pos) : null;
  if (!node || node.type.name !== "mention") return false;
  const chain = () => editor.chain().focus();
  const [kind, id] = String(node.attrs.id).split(":");
  const actions = [["Open", () => (window.location.href = mention.getAttribute("href")), "external-link"]];
  if (kind === "image") {
    if (document.body.dataset.gm) actions.push(["Show on display", () => gsm.showImage(`/images/${id}/file`, node.attrs.label), "cast"]);
    const floats = [["Float left", null, "panel-left"], ["Float right", "right", "panel-right"], ["Inline", "inline", "type"]];
    const setFloat = (v) => chain().command(({ tr }) => !!tr.setNodeAttribute(pos, "float", v)).run();
    actions.push(null, ...floats.filter(([, v]) => v !== (node.attrs.float ?? null)).map(([label, v, icon]) => [label, () => setFloat(v), icon]));
  }
  actions.push(
    null,
    // Typing [[ again opens the picker on the old label to choose anew.
    ["Change…", () => chain().deleteRange({ from: pos, to: pos + node.nodeSize }).insertContentAt(pos, `[[${node.attrs.label}`).run(), "pencil"],
    ["Remove", () => chain().deleteRange({ from: pos, to: pos + node.nodeSize }).run(), "trash-2", true],
  );
  // Images can be large, so open at the click rather than below the image.
  showMenu(kind === "image" ? { left: e.clientX, bottom: e.clientY } : mention.getBoundingClientRect(), actions);
  return true;
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
