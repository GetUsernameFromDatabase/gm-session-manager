// Collaborative note blocks: one TipTap editor per block, kept in sync with
// ProseMirror's collab protocol. The server is the authority: steps are
// POSTed against a version, broadcast over the websocket, and a 409 hands
// back the steps to rebase on.
import { Editor, Extension } from "/assets/vendor/tiptap-core.js";
import { StarterKit } from "/assets/vendor/tiptap-starter.js";
import { Image } from "/assets/vendor/tiptap-image.js";
import { collab, getVersion, receiveTransaction, sendableSteps } from "/assets/vendor/pm-collab.js";
import { Step } from "/assets/vendor/pm-transform.js";
import { LINK_PATHS, WikiLink, keepCaretBeside, linkClick, mentionActions, showMenu } from "/assets/js/wikilink.js";

const SNAPSHOT_EVERY_MS = 10000;
const campaign = document.body.dataset.campaign;
const editors = new Map(); // block id -> BlockEditor

function randomID() {
  return Math.random().toString(36).slice(2, 10) + Date.now().toString(36);
}

async function json(res) {
  if (!res.ok) throw new Error(`${res.status} ${res.statusText}`);
  return res.json();
}

const Collab = Extension.create({
  name: "collab",
  addOptions() {
    return { version: 0, clientID: "" };
  },
  addProseMirrorPlugins() {
    return [collab({ version: this.options.version, clientID: this.options.clientID })];
  },
});

// links lists the [[mentions]] in a document for backlinks.
function links(doc) {
  const out = [];
  const walk = (n) => {
    if (n.type === "mention" && n.attrs && n.attrs.id) {
      const [type, id] = String(n.attrs.id).split(":");
      const num = Number(id);
      if (LINK_PATHS[type] && num > 0 && !out.some((l) => l.type === type && l.id === num)) {
        out.push({ type, id: num });
      }
    }
    (n.content || []).forEach(walk);
  };
  walk(doc);
  return out;
}

class BlockEditor {
  constructor(el, id) {
    this.el = el;
    this.id = id;
    this.clientID = randomID();
    this.inflight = false;
    this.dirtySince = 0; // version of the last snapshot sent
    this.destroyed = false;
  }

  async start() {
    const data = await json(await fetch(`/blocks/${this.id}/doc`));
    if (this.destroyed) return;
    this.editor?.destroy();
    this.editor = new Editor({
      element: this.el,
      editable: data.editable,
      content: data.doc,
      extensions: [
        StarterKit.configure({ link: { openOnClick: false } }),
        Image,
        WikiLink,
        Collab.configure({ version: data.docVersion, clientID: this.clientID }),
      ],
      onUpdate: () => this.send(),
      onSelectionUpdate: keepCaretBeside,
    });
    this.lastSnapshot = data.docVersion;
    this.receive({ version: data.docVersion, steps: data.steps, clientIDs: data.clientIDs });
    gsm.send({ type: "sub", blocks: [this.id] });
    this.timer = setInterval(() => this.snapshot(), SNAPSHOT_EVERY_MS);
  }

  setEditable(on) {
    if (this.editor && this.editor.isEditable !== on) this.editor.setEditable(on);
  }

  version() {
    return getVersion(this.editor.state);
  }

  // receive applies steps that start at p.version, skipping any already
  // applied and fetching any gap first.
  receive(p) {
    if (!this.editor || !p.steps) return;
    // Changing the DOM under an IME composition (phone keyboards compose whole words) duplicates text; apply after it.
    if (this.editor.view.composing) return this.retryLater();
    const have = this.version();
    if (p.version > have) {
      this.catchUp();
      return;
    }
    const skip = have - p.version;
    const steps = p.steps.slice(skip);
    if (!steps.length) return;
    const schema = this.editor.schema;
    const tr = receiveTransaction(
      this.editor.state,
      steps.map((s) => Step.fromJSON(schema, s)),
      p.clientIDs.slice(skip),
    );
    this.editor.view.dispatch(tr);
    this.send();
  }

  async catchUp() {
    if (!this.editor) return;
    if (this.editor.view.composing) return this.retryLater();
    const res = await fetch(`/blocks/${this.id}/steps?since=${this.version()}`);
    if (res.status === 410) return this.start(); // trimmed behind a snapshot
    if (res.ok) this.receive(await res.json());
  }

  // retryLater catches up shortly, once an IME composition has ended.
  retryLater() {
    clearTimeout(this.retry);
    this.retry = setTimeout(() => this.catchUp(), 100);
  }

  async send() {
    if (this.inflight || !this.editor || !this.editor.isEditable) return;
    const s = sendableSteps(this.editor.state);
    if (!s) return;
    this.inflight = true;
    try {
      const res = await fetch(`/blocks/${this.id}/steps`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ version: s.version, clientID: String(s.clientID), steps: s.steps.map((x) => x.toJSON()) }),
      });
      this.inflight = false;
      if (res.status === 409) {
        this.receive(await res.json()); // rebases, then sends again
      } else if (res.status === 410) {
        await this.start();
      } else if (res.ok) {
        // Confirmation normally arrives over the websocket; fetch it too in
        // case the socket is down.
        await this.catchUp();
      } else {
        setTimeout(() => this.send(), 2000);
      }
    } catch {
      this.inflight = false;
      setTimeout(() => this.send(), 2000);
    }
  }

  // snapshot stores the whole document so the server can trim old steps
  // and index text and links. Only when everything local is confirmed.
  snapshot(keepalive = false) {
    if (!this.editor || !this.editor.isEditable) return;
    const v = this.version();
    if (v <= this.lastSnapshot || sendableSteps(this.editor.state)) return;
    this.lastSnapshot = v;
    const doc = this.editor.getJSON();
    fetch(`/blocks/${this.id}/snapshot`, {
      method: "POST",
      keepalive,
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ version: v, doc, text: this.editor.getText(), links: links(doc) }),
    }).catch(() => {});
  }

  // upload stores a file on this block; null (after telling the user) on failure.
  async upload(file) {
    const form = new FormData();
    form.append("file", file);
    const res = await fetch(`/blocks/${this.id}/attachments`, { method: "POST", body: form });
    const data = await res.json();
    if (!res.ok) {
      alert(data.error || "Upload failed");
      return null;
    }
    return data;
  }

  async attach(file) {
    const data = await this.upload(file);
    if (!data) return;
    const chain = this.editor.chain().focus();
    if (data.image) chain.setImage({ src: data.url, alt: data.name }).run();
    else chain.insertContent({ type: "text", text: data.name, marks: [{ type: "link", attrs: { href: data.url } }] }).run();
  }

  destroy() {
    this.destroyed = true;
    clearInterval(this.timer);
    clearTimeout(this.retry);
    this.snapshot(true);
    gsm.send({ type: "unsub", blocks: [this.id] });
    this.editor?.destroy();
  }
}

// sync mounts editors for new blocks, drops removed ones and follows
// editability changes after any swap.
function sync() {
  const seen = new Set();
  document.querySelectorAll("[data-editor]").forEach((el) => {
    const id = Number(el.dataset.editor);
    seen.add(id);
    const article = el.closest("[data-block]");
    const editable = article ? article.dataset.editable === "true" : false;
    let ed = editors.get(id);
    if (!ed || ed.el !== el) {
      ed?.destroy();
      ed = new BlockEditor(el, id);
      editors.set(id, ed);
      ed.ready = ed.start().catch((err) => console.error("note block", id, err));
    } else {
      ed.setEditable(editable);
    }
  });
  for (const [id, ed] of editors) {
    if (!seen.has(id)) {
      ed.destroy();
      editors.delete(id);
    }
  }
}

document.addEventListener("gsm:json", (e) => {
  const m = e.detail;
  if (m.type === "steps") editors.get(m.block)?.receive(m);
});

// After a reconnect, resubscribe and fetch whatever was missed.
document.addEventListener("gsm:wsopen", () => {
  if (!editors.size) return;
  gsm.send({ type: "sub", blocks: [...editors.keys()] });
  editors.forEach((ed) => ed.catchUp());
});

document.addEventListener("change", (e) => {
  const input = e.target.closest("input[data-attach]");
  if (!input || !input.files.length) return;
  editors.get(Number(input.dataset.attach))?.attach(input.files[0]);
  input.value = "";
});

// pickFile asks for one file and calls done with it.
function pickFile(accept, done) {
  const input = document.createElement("input");
  input.type = "file";
  input.accept = accept;
  input.addEventListener("change", () => input.files.length && done(input.files[0]));
  input.click();
}

// nodeActions offers what can be done to the [[link]], attached image or
// file link that was clicked in an editable block.
function nodeActions(ed, target, e) {
  const view = ed.editor.view;
  const hit = view.posAtCoords({ left: e.clientX, top: e.clientY });
  if (!hit) return;
  const chain = () => ed.editor.chain().focus();
  if (mentionActions(ed.editor, target, e)) return;
  const pos = hit.inside;
  const node = pos >= 0 ? view.state.doc.nodeAt(pos) : null;
  const removeNode = () => chain().deleteRange({ from: pos, to: pos + node.nodeSize }).run();
  if (target.closest("img") && node && node.type.name === "image") {
    // Images can be large, so open at the click rather than below the image.
    showMenu({ left: e.clientX, bottom: e.clientY }, [
      ["Open", () => window.open(node.attrs.src, "_blank")],
      [
        "Replace…",
        () =>
          pickFile("image/*", async (file) => {
            const data = await ed.upload(file);
            if (data) ed.editor.view.dispatch(ed.editor.state.tr.setNodeMarkup(pos, null, { ...node.attrs, src: data.url, alt: data.name }));
          }),
      ],
      ["Remove", removeNode],
    ]);
    return;
  }
  const link = target.closest("a[href]");
  if (!link) return;
  const href = link.getAttribute("href");
  // Select the whole link so Replace and Remove act on all of its text.
  const select = () => chain().setTextSelection(hit.pos).extendMarkRange("link");
  const actions = [
    ["Open", () => window.open(href, "_blank")],
    ["Remove", () => select().deleteSelection().run()],
  ];
  if (href.startsWith("/attachments/")) {
    actions.splice(1, 0, [
      "Replace…",
      () =>
        pickFile("", async (file) => {
          const data = await ed.upload(file);
          if (data) select().insertContent({ type: "text", text: data.name, marks: [{ type: "link", attrs: { href: data.url } }] }).run();
        }),
    ]);
  }
  showMenu(link.getBoundingClientRect(), actions);
}

// In an editable block a click on a link or attachment opens its menu;
// Ctrl/Cmd-click follows the link.
document.addEventListener("click", (e) => {
  if (!e.target.closest(".note-editor a[href], .note-editor img")) return;
  const ed = [...editors.values()].find((x) => x.el.contains(e.target));
  linkClick(e, ed?.editor, () => nodeActions(ed, e.target, e));
});

document.addEventListener("click", (e) => {
  // A tap or click outside an open access popover closes it.
  document.querySelectorAll("details[data-popover][open]").forEach((d) => {
    if (!d.contains(e.target)) d.open = false;
  });
});


document.addEventListener("visibilitychange", () => {
  if (document.visibilityState === "hidden") editors.forEach((ed) => ed.snapshot(true));
});

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
sync();

// Editors fill in after load, so the browser's own restore on Back hits a short page; restore once they have.
history.scrollRestoration = "manual";
window.addEventListener("pagehide", () => history.replaceState({ ...history.state, scrollY: window.scrollY }, ""));
const savedY = history.state?.scrollY;
if (savedY) Promise.allSettled([...editors.values()].map((ed) => ed.ready)).then(() => window.scrollTo(0, savedY));
