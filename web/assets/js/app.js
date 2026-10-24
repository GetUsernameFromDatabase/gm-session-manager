// Page glue for htmx and the websocket. Loaded as a classic script after
// htmx, before any page content.
(function () {
  "use strict";

  // Responses are out-of-band fragments (flash messages, refetched live
  // views); nothing swaps into the element that made the request unless it
  // asks with hx-swap.
  htmx.config.defaultSwapStyle = "none";
  // An invalid form says why instead of silently not sending.
  htmx.config.reportValidityOfForms = true;
  // 422 carries a flash message to show; other errors are not swapped.
  htmx.config.responseHandling = [
    { code: "204", swap: false },
    { code: "[23]..", swap: true },
    { code: "422", swap: true, error: false },
    { code: "[45]..", swap: false, error: true },
  ];

  // Theme is per device: the stored choice, else the system preference.
  // Applied here, in <head>, so the page never paints in the wrong theme.
  const THEME_KEY = "gsm-theme";
  let stored = null;
  try {
    stored = localStorage.getItem(THEME_KEY);
  } catch (_) {}
  document.documentElement.classList.toggle(
    "dark",
    stored ? stored === "dark" : matchMedia("(prefers-color-scheme: dark)").matches,
  );

  const gsm = (window.gsm = {
    socket: null,

    toggleTheme() {
      const dark = document.documentElement.classList.toggle("dark");
      try {
        localStorage.setItem(THEME_KEY, dark ? "dark" : "light");
      } catch (_) {}
    },

    // send writes a JSON message to the hub if connected.
    send(msg) {
      if (gsm.socket) gsm.socket.send(JSON.stringify(msg));
    },

    openSwitcher() {
      const d = document.getElementById("switcher");
      if (!d) return;
      d.showModal();
      // The dialog has no padding, so a click whose target is the dialog itself landed on the backdrop.
      d.onclick = (e) => { if (e.target === d) d.close(); };
      const input = d.querySelector("input");
      input.value = "";
      input.focus();
      htmx.trigger(input, "search");
    },
  });

  // The display window shows the image the GM picked. The pick lives in
  // this browser only: localStorage, followed through the storage event.
  // shortcut: windows in other browsers or devices never see it; use a hub topic if that is wanted.
  const DISPLAY_KEY = "gsm-display-";
  const displayKey = () => DISPLAY_KEY + document.body.dataset.campaign;

  function shownImage() {
    try {
      return JSON.parse(localStorage.getItem(displayKey()));
    } catch (_) {
      return null;
    }
  }

  function syncDisplay() {
    const shown = shownImage();
    const img = document.getElementById("display-image");
    if (img) {
      img.hidden = !shown;
      if (shown) {
        img.src = shown.src;
        img.alt = shown.name;
      } else img.removeAttribute("src");
    }
    document.querySelectorAll("[data-show-image]").forEach((b) => {
      b.setAttribute("aria-pressed", String(!!shown && b.dataset.src === shown.src));
    });
  }

  Object.assign(gsm, {
    // showImage puts an image on the display window, or clears it when that
    // image is already shown.
    showImage(src, name) {
      const shown = shownImage();
      try {
        if (shown && shown.src === src) localStorage.removeItem(displayKey());
        else localStorage.setItem(displayKey(), JSON.stringify({ src, name }));
      } catch (_) {}
      syncDisplay();
    },

    clearDisplay() {
      try {
        localStorage.removeItem(displayKey());
      } catch (_) {}
      syncDisplay();
    },

    // openDisplay opens the display as its own window, or focuses it.
    openDisplay(href) {
      const win = window.open(href, "gsm-display", "popup,width=1280,height=720");
      if (win) win.focus();
      return !!win;
    },
  });

  window.addEventListener("storage", (e) => {
    if (e.key === displayKey()) syncDisplay();
  });

  // filterImport hides import rows outside the name filter, the chosen
  // category ("*" for all; rows with no category carry "-") and the chosen
  // status (ticked, unticked, or a link that didn't load).
  function filterImport(bar) {
    if (!bar) return;
    const text = bar.querySelector("[data-import-q]").value.trim().toLowerCase();
    const sel = bar.querySelector("input[name=filter_category]");
    const category = sel ? sel.value : "*";
    const status = bar.querySelector("input[name=filter_status]")?.value || "any";
    const hasStatus = (tr) => {
      const box = tr.querySelector("input[type=checkbox]");
      if (status === "ticked") return box.checked;
      if (status === "unticked") return !box.checked;
      if (status === "broken") return linkResults.get(tr.dataset.previewSrc) === "broken";
      return true;
    };
    document.querySelectorAll("[data-import-row]").forEach((tr) => {
      const inCategory = category === "*" || tr.dataset.category === category;
      tr.hidden = !inCategory || !tr.dataset.name.toLowerCase().includes(text) || !hasStatus(tr);
    });
    const hint = document.getElementById("status-hint");
    if (hint) hint.hidden = !(status === "broken" && !linkResults.size);
  }

  // encodeImportRows writes the ticked import rows for parseRowSet: runs as
  // "a-b" and single rows as "a", or after "!" the unticked rows when that
  // is shorter. Rows that cannot be imported never break a run.
  function encodeImportRows(form) {
    const rows = [...form.querySelectorAll("[data-import-row]")].map((tr) => {
      const box = tr.querySelector("input[type=checkbox]");
      return { i: Number(tr.dataset.index), ticked: box.disabled ? null : box.checked };
    });
    const runs = (want) => {
      const out = [];
      let start = -1;
      let end = -1;
      const flush = () => {
        if (start >= 0) out.push(start === end ? `${start}` : `${start}-${end}`);
        start = -1;
      };
      for (const r of rows) {
        if (r.ticked === null) continue;
        if (r.ticked !== want) flush();
        else {
          if (start < 0) start = r.i;
          end = r.i;
        }
      }
      flush();
      return out.join(",");
    };
    const ticked = runs(true);
    const unticked = "!" + runs(false);
    return unticked.length < ticked.length ? unticked : ticked;
  }

  const IMPORT_ROWS = "rows";
  document.addEventListener("htmx:configRequest", (e) => {
    const form = document.getElementById("import-review");
    if (form && form.contains(e.detail.elt)) e.detail.parameters[IMPORT_ROWS] = encodeImportRows(form);
  });
  document.addEventListener("submit", (e) => {
    if (e.target.id === "import-review") e.target.elements[IMPORT_ROWS].value = encodeImportRows(e.target);
  });

  // Link check: each distinct link loads as an image, exactly as the picker
  // and display will load it, so 404s, hotlink blocks, expired links and
  // non-images all fail. The download stops as soon as the image's size is
  // known (its first few KB); an error response is a few bytes anyway.
  const LINK_CHECKS_AT_ONCE = 8;
  const LINK_TIMEOUT_MS = 10000;
  const LINK_POLL_MS = 50;
  const LINK_LABELS = { checking: "checking…", ok: "✓ loads", broken: "didn't load" };
  const linkResults = new Map(); // link -> checking | ok | broken

  function checkLink(link) {
    return new Promise((resolve) => {
      const img = new Image();
      const started = Date.now();
      let poll = 0;
      const finish = (state) => {
        clearInterval(poll);
        img.onload = img.onerror = null;
        img.src = ""; // cancels the rest of the download
        resolve(state);
      };
      img.onload = () => finish("ok");
      img.onerror = () => finish("broken");
      poll = setInterval(() => {
        if (img.naturalWidth > 0) finish("ok");
        else if (Date.now() - started > LINK_TIMEOUT_MS) finish("broken");
      }, LINK_POLL_MS);
      img.src = link;
    });
  }

  // tickedLinks are the links of the ticked rows that can be imported.
  function tickedLinks() {
    const links = new Set();
    document.querySelectorAll("[data-import-row][data-preview-src]").forEach((tr) => {
      if (tr.querySelector("input[type=checkbox]").checked) links.add(tr.dataset.previewSrc);
    });
    return links;
  }

  async function checkLinks() {
    const queue = [...tickedLinks()].filter((l) => !linkResults.has(l));
    queue.forEach((l) => linkResults.set(l, "checking"));
    queueSummary();
    const worker = async () => {
      for (let link = queue.shift(); link; link = queue.shift()) {
        linkResults.set(link, await checkLink(link));
        queueSummary();
      }
    };
    await Promise.all(Array.from({ length: LINK_CHECKS_AT_ONCE }, worker));
    // A "Didn't load" filter can only be applied once results are in.
    filterImport(document.querySelector("[data-import-filter]"));
  }

  // paintLinkResults shows each row's result, the progress line and the
  // untick button; rows re-rendered by htmx get their results back here.
  function paintLinkResults() {
    if (!linkResults.size) return;
    document.querySelectorAll("[data-import-row][data-preview-src]").forEach((tr) => {
      const state = linkResults.get(tr.dataset.previewSrc);
      const cell = tr.querySelector("[data-link-status]");
      if (!cell) return;
      if (state) cell.dataset.state = state;
      else delete cell.dataset.state;
      cell.textContent = state ? LINK_LABELS[state] : "";
    });
    const links = tickedLinks();
    const states = [...links].map((l) => linkResults.get(l)).filter(Boolean);
    const done = states.filter((s) => s !== "checking").length;
    const broken = states.filter((s) => s === "broken").length;
    const box = document.getElementById("link-check");
    if (box) {
      box.hidden = false;
      document.getElementById("link-check-status").textContent =
        `Checked ${done} of ${links.size} links` + (broken ? ` · ${broken} didn't load` : "");
    }
    document.querySelector("[data-untick-broken]")?.classList.toggle("hidden", !broken);
  }

  // summarizeImport counts the ticked rows, in total and per category they
  // will go into, next to the import button.
  function summarizeImport() {
    const out = document.getElementById("import-summary");
    if (!out) return;
    const counts = new Map();
    let total = 0;
    let failed = 0;
    document.querySelectorAll("[data-import-row]").forEach((tr) => {
      const box = tr.querySelector("input[type=checkbox]");
      if (box.disabled || !box.checked) return;
      total++;
      counts.set(tr.dataset.category, (counts.get(tr.dataset.category) || 0) + 1);
      if (linkResults.get(tr.dataset.previewSrc) === "broken") failed++;
    });
    const parts = [...counts]
      .sort(([a], [b]) => (a === "-") - (b === "-") || a.localeCompare(b))
      .map(([c, n]) => `${c === "-" ? "No category" : c} ${n}`);
    out.textContent =
      `${total} ${total === 1 ? "row" : "rows"} to import` +
      (parts.length ? ": " + parts.join(" · ") : "") +
      (failed ? ` · ${failed} didn't load` : "");
  }

  // Ticking many rows or checking many links fires an update each; repaint
  // the link results and the summary once a frame.
  let summaryQueued = false;
  function queueSummary() {
    if (summaryQueued) return;
    summaryQueued = true;
    requestAnimationFrame(() => {
      summaryQueued = false;
      paintLinkResults();
      summarizeImport();
    });
  }

  ["input", "change"].forEach((type) =>
    document.addEventListener(type, (e) => {
      const bar = e.target.closest && e.target.closest("[data-import-filter]");
      if (bar) filterImport(bar);
      if (e.target.closest && e.target.closest("[data-import-row]")) queueSummary();
    }),
  );

  document.addEventListener("click", (e) => {
    const show = e.target.closest("[data-show-image]");
    if (show) gsm.showImage(show.dataset.src, show.dataset.name);
    if (e.target.closest("[data-clear-display]")) gsm.clearDisplay();
    const open = e.target.closest("a[data-open-display]");
    if (open && !e.ctrlKey && !e.metaKey && gsm.openDisplay(open.href)) e.preventDefault();
    // Tick or untick the import rows the filter shows.
    const check = e.target.closest("[data-check-shown]");
    if (check) {
      const want = check.dataset.checkShown;
      document.querySelectorAll("[data-import-row]:not([hidden]) [role=checkbox]:not([data-disabled])").forEach((box) => {
        if (box.getAttribute("aria-checked") !== want) box.click();
      });
    }
    const copy = e.target.closest("[data-copy]");
    if (copy) {
      navigator.clipboard.writeText(copy.dataset.copy).then(() => {
        copy.dataset.copied = "";
        setTimeout(() => delete copy.dataset.copied, 1500);
      });
    }
    if (e.target.closest("[data-check-links]")) checkLinks();
    if (e.target.closest("[data-untick-broken]")) {
      document.querySelectorAll("[data-import-row][data-preview-src]").forEach((tr) => {
        const box = tr.querySelector("[role=checkbox]");
        if (linkResults.get(tr.dataset.previewSrc) === "broken" && box.getAttribute("aria-checked") === "true") box.click();
      });
    }
  });

  // Hovering an import row previews its image in the one preview pane. A
  // hidden pane (small screens) fetches nothing.
  document.addEventListener("mouseover", (e) => {
    const tr = e.target.closest && e.target.closest("[data-preview-src]");
    const img = document.getElementById("import-preview");
    if (!tr || !img || !img.parentElement.checkVisibility() || img.getAttribute("src") === tr.dataset.previewSrc) return;
    img.src = tr.dataset.previewSrc;
    img.alt = tr.dataset.name;
    img.hidden = false;
    document.querySelector("[data-import-hint]")?.remove();
  });

  // A joining visitor's name goes with each join form and is remembered on
  // this device only.
  const PLAYER_KEY = "gsm-player";
  document.addEventListener("DOMContentLoaded", () => {
    const input = document.getElementById("player-name");
    if (!input) return;
    try {
      input.value = localStorage.getItem(PLAYER_KEY) || "";
    } catch (_) {}
  });
  // The native bubble fades fast, so an invalid name also stays marked.
  const markPlayer = (input, invalid) => {
    if (invalid) input.setAttribute("aria-invalid", "true");
    else input.removeAttribute("aria-invalid");
    document.getElementById("player-name-error").hidden = !invalid;
  };
  document.addEventListener("input", (e) => {
    if (e.target.id === "player-name" && e.target.checkValidity()) markPlayer(e.target, false);
  });
  document.addEventListener("submit", (e) => {
    const input = document.getElementById("player-name");
    if (!input || !e.target.matches("form[data-player]")) return;
    if (!input.reportValidity()) {
      markPlayer(input, true);
      input.focus();
      e.preventDefault();
      return;
    }
    const name = input.value.trim();
    e.target.elements.player.value = name;
    try {
      localStorage.setItem(PLAYER_KEY, name);
    } catch (_) {}
  });

  // Forms marked data-busy show their busy indicator (aria-busy) while they
  // submit or while one of their htmx requests runs. A submitting form also
  // disables its submit buttons, so a second click cannot submit twice.
  const busyForm = (el) => el && el.closest && el.closest("form[data-busy]");
  document.addEventListener("submit", (e) => {
    const form = busyForm(e.target);
    if (!form || e.defaultPrevented) return;
    form.setAttribute("aria-busy", "true");
    // Disabled after this tick, so the clicked button still submits its value.
    setTimeout(() => form.querySelectorAll("button[type=submit]").forEach((b) => (b.disabled = true)));
  });
  document.addEventListener("htmx:beforeRequest", (e) => busyForm(e.detail.elt)?.setAttribute("aria-busy", "true"));
  document.addEventListener("htmx:afterRequest", () => {
    // The form itself may have been swapped out; clear whichever is on the page.
    document.querySelectorAll("form[data-busy][aria-busy]").forEach((f) => f.removeAttribute("aria-busy"));
  });
  // A form's 422 flash shows inside that form, where the user is looking
  // (the page flash sits behind an open dialog).
  document.addEventListener("htmx:beforeRequest", (e) => {
    e.detail.elt.closest("form")?.querySelector("[data-form-error]")?.remove();
  });
  document.addEventListener("htmx:beforeSwap", (e) => {
    const form = e.detail.elt.closest("form");
    if (e.detail.xhr.status !== 422 || !form) return;
    const flash = new DOMParser().parseFromString(e.detail.serverResponse, "text/html").querySelector("#flash[hx-swap-oob]");
    if (!flash) return;
    const p = document.createElement("p");
    p.dataset.formError = "";
    p.setAttribute("role", "alert");
    p.className = "basis-full text-sm text-destructive";
    p.textContent = flash.textContent.trim();
    form.append(p);
    e.detail.shouldSwap = false;
  });

  // The command palette puts the picked example in the command bar, ready to
  // adjust; "/" jumps to the bar from anywhere but a text field.
  document.addEventListener("command-select", (e) => {
    const root = e.target.closest("[data-command-help]");
    const input = root && document.getElementById(root.dataset.commandHelp);
    if (!input) return;
    input.value = e.target.dataset.example;
    commandHint(input);
    root.querySelector("[data-slot=dialog-close]").click();
    setTimeout(() => input.focus());
  });

  // While a command is typed, the hint under the bar shows the matching
  // forms (from the palette) with the part being typed in bold; Run is off
  // until the first word is a known command.
  // A slot's alternatives: "-n | +n | =n" or "add|remove"; "<a|b>" is one.
  const SLOT_ALT = /<[^>]*>|[^|\s]+/g;
  // The fixed part an alternative starts with: "+" of "+n", none for "<name>".
  const altLiteral = (a) => (/^"?</.test(a) ? "" : /^\W?n$/.test(a) ? a.slice(0, -1) : a);
  function commandState(input) {
    const verbs = input.dataset.verbs.split(" ");
    const tokens = input.value.match(/"[^"]*"?|\S+/g) || [];
    const at = /(^|\s)$/.test(input.value) ? tokens.length : tokens.length - 1;
    const verb = (tokens[0] || "").toLowerCase();
    const all = [...document.querySelectorAll("[data-command-help] [data-syntax]")].map((el) => {
      // Spaced alternatives fill one slot, unless they start another form ("use … | restore …").
      const words = [];
      const raw = el.dataset.syntax.split(" ");
      for (let i = 0; i < raw.length; i++) {
        if (raw[i] === "|" && words.length && !verbs.includes(raw[i + 1])) words[words.length - 1] += " | " + raw[++i];
        else words.push(raw[i]);
      }
      const forVerbs = el.dataset.syntax.split(" | ").map((p) => p.split(" ")[0]).filter((w) => verbs.includes(w));
      return { words, forVerbs, desc: el.dataset.desc };
    });
    let found = all.filter((h) => h.forVerbs.some((v) => (at === 0 ? v.startsWith(verb) : v === verb)));
    // A second word that is a plain choice ("use|restore") narrows the forms.
    const second = (tokens[1] || "").toLowerCase();
    if (at >= 1 && second) {
      const narrowed = found.filter((h) => {
        const opts = (h.words[1] || "").split("|");
        if (!opts.every((o) => /^[a-z]+$/.test(o))) return false;
        return opts.some((o) => (at === 1 ? o.startsWith(second) : o === second));
      });
      if (narrowed.length) found = narrowed;
    }
    return { verbs, tokens, at, verb, found };
  }
  const strongText = (text) => Object.assign(document.createElement("strong"), { className: "text-foreground", textContent: text });
  function commandHint(input) {
    const out = document.getElementById("command-hint");
    const { verbs, tokens, at, verb, found } = commandState(input);
    input.form.querySelector("[data-command-run]").disabled = !verbs.includes(verb);
    out.replaceChildren();
    if (!tokens.length) return;
    if (!found.length) {
      out.textContent = "Unknown command. Browse Commands for the list.";
      return;
    }
    const typed = (tokens[at] || "").toLowerCase();
    for (const h of found.slice(0, 5)) {
      const line = document.createElement("div");
      const bold = Math.min(at, h.words.length - 1);
      const code = document.createElement("code");
      h.words.forEach((w, i) => {
        if (i) code.append(" ");
        if (i !== bold) return code.append(w);
        // Bold the alternative being typed ("+n" for "hp +"), else the whole slot.
        const alts = w.match(SLOT_ALT) || [];
        const pick =
          typed &&
          (alts.find((a) => altLiteral(a) && (typed.startsWith(altLiteral(a)) || altLiteral(a).startsWith(typed))) ||
            alts.find((a) => !altLiteral(a)));
        if (!pick) return code.append(strongText(w));
        w.split(/(<[^>]*>|[^|\s]+)/).forEach((p) => code.append(p === pick ? strongText(p) : p));
      });
      line.append(code, " — " + h.desc);
      out.append(line);
    }
  }

  // Tab completes the word at the end of the bar (or starts the next one),
  // and again cycles through the choices (Shift+Tab back); Ctrl+Space lists
  // them in the hint first.
  // Choices are command words, the syntax's keywords and, for placeholders,
  // names from the server.
  async function commandChoices(input) {
    const { verbs, tokens, at, found } = commandState(input);
    const typed = at < tokens.length ? tokens[at] : "";
    const prefix = typed.replace(/^"/, "").toLowerCase();
    const list = [];
    if (at === 0) list.push(...verbs);
    for (const slot of new Set(at ? found.map((h) => h.words[at]).filter(Boolean) : [])) {
      for (const a of slot.match(SLOT_ALT) || []) {
        if (/^<[a-z]+(\|[a-z]+)+>$/.test(a)) list.push(...a.slice(1, -1).split("|"));
        else if (/^[a-z][a-z-]*$/.test(a) && a !== "n") list.push(a);
        else if (a === "±n") list.push("+", "-");
        else if (/^\Wn$/.test(a)) list.push(altLiteral(a)); // the sign of "-n", "+n", "=n"
        else if (/^"?<[^>]*>"?$/.test(a)) {
          const q = new URLSearchParams({ words: tokens.slice(0, at).join(" "), slot: a, q: prefix });
          const res = await fetch(input.dataset.complete + "?" + q);
          const names = res.ok ? await res.json() : [];
          list.push(...names.map((n) => (a.includes('"') && n.includes(" ") ? `"${n}"` : n)));
        }
      }
    }
    const choices = [...new Set(list)].filter((c) => c.replace(/^"/, "").toLowerCase().startsWith(prefix));
    return { base: input.value.slice(0, input.value.length - typed.length), choices };
  }
  let cycle = null; // { input, base, choices, i, value } while Tab cycles
  function showChoices() {
    const out = document.getElementById("command-hint");
    if (!cycle.choices.length) {
      // Keep the forms in view; the note is a warning, not an error.
      commandHint(cycle.input);
      return out.append(Object.assign(document.createElement("div"), { className: "text-amber-600 dark:text-amber-500", textContent: "No choices here." }));
    }
    out.replaceChildren();
    const line = document.createElement("div");
    // Every command word fits; names and slugs can run long.
    const shown = cycle.base ? 12 : cycle.choices.length;
    cycle.choices.slice(0, shown).forEach((c, i) => {
      if (i) line.append(" ");
      const code = document.createElement("code");
      code.append(i === cycle.i ? strongText(c) : c);
      line.append(code);
    });
    if (cycle.choices.length > shown) line.append(" …");
    out.append(line);
  }
  document.addEventListener("keydown", async (e) => {
    const input = e.target;
    if (input.id !== "command-input") return;
    if (e.key === "Escape" && cycle) {
      cycle = null;
      return commandHint(input);
    }
    const tab = e.key === "Tab" && !e.ctrlKey && !e.altKey && !e.metaKey;
    const list = e.key === " " && e.ctrlKey;
    if (!tab && !list) return;
    if (tab && cycle && cycle.input === input && cycle.value === input.value && cycle.choices.length) {
      e.preventDefault();
      const n = cycle.choices.length;
      cycle.i = e.shiftKey ? (cycle.i - 1 + n) % n : (cycle.i + 1) % n;
      input.value = cycle.value = cycle.base + cycle.choices[cycle.i];
      commandHint(input);
      return showChoices();
    }
    // Tab on an empty bar, and Shift+Tab outside a cycle, still move focus.
    if (tab && (!input.value.trim() || e.shiftKey)) return;
    e.preventDefault();
    const value = input.value;
    const { base, choices } = await commandChoices(input);
    if (input.value !== value) return;
    cycle = { input, base, choices, i: -1, value };
    if (tab && choices.length === 1) {
      cycle = null;
      // A sign waits for its number ("hp -").
      input.value = base + choices[0] + (/^\W$/.test(choices[0]) ? "" : " ");
      return commandHint(input);
    }
    if (tab && choices.length) {
      cycle.i = 0;
      input.value = cycle.value = base + choices[0];
      commandHint(input);
    }
    showChoices();
  });
  document.addEventListener("input", (e) => {
    if (e.target.id !== "command-input") return;
    cycle = null;
    commandHint(e.target);
  });
  document.addEventListener("reset", (e) => {
    const input = e.target.querySelector && e.target.querySelector("#command-input");
    if (input) setTimeout(() => commandHint(input));
  });
  htmx.onLoad((el) => {
    const input = el.querySelector("#command-input");
    if (input) commandHint(input);
  });

  document.addEventListener("keydown", (e) => {
    const input = document.getElementById("command-input");
    if (e.key !== "/" || !input || e.ctrlKey || e.metaKey || e.altKey) return;
    if (e.target.closest("input, textarea, select, [contenteditable]")) return;
    e.preventDefault();
    input.focus();
  });

  // The homebrew editor shows only the chosen kind's fields; hidden fields
  // are disabled so they submit nothing.
  function showKindFields(form) {
    const kind = form.elements.kind?.value;
    form.querySelectorAll("[data-kinds]").forEach((el) => {
      const off = !el.dataset.kinds.split(" ").includes(kind);
      el.hidden = off;
      el.querySelectorAll("input, textarea").forEach((f) => (f.disabled = off));
    });
  }
  document.addEventListener("change", (e) => {
    const form = e.target.closest && e.target.closest("form[data-homebrew]");
    if (form) showKindFields(form);
  });
  htmx.onLoad(() => document.querySelectorAll("form[data-homebrew]").forEach(showKindFields));

  // The creation wizard's ability scores show their point-buy cost (27
  // points for scores 8 to 15), or why they are not a point buy.
  const POINT_COSTS = { 8: 0, 9: 1, 10: 2, 11: 3, 12: 4, 13: 5, 14: 7, 15: 9 };
  function pointCost(form) {
    const out = form.querySelector("[data-point-cost]");
    if (!out) return;
    let total = 0;
    for (const f of form.querySelectorAll("input[type=number]")) {
      const cost = POINT_COSTS[f.value];
      if (cost === undefined) {
        out.textContent = "Not a point buy (scores outside 8–15).";
        return;
      }
      total += cost;
    }
    out.textContent = `Point buy: ${total} of 27 points` + (total > 27 ? " (too many)" : "");
  }
  document.addEventListener("input", (e) => {
    const form = e.target.closest && e.target.closest("form[data-point-buy]");
    if (form) pointCost(form);
  });
  htmx.onLoad(() => document.querySelectorAll("form[data-point-buy]").forEach(pointCost));

  // A page restored by Back must not stay busy.
  window.addEventListener("pageshow", () => {
    document.querySelectorAll("form[data-busy]").forEach((f) => {
      f.removeAttribute("aria-busy");
      f.querySelectorAll("button[type=submit]").forEach((b) => (b.disabled = false));
    });
  });

  htmx.onLoad(() => {
    syncDisplay();
    // Refreshed import rows keep the filter the GM had set.
    filterImport(document.querySelector("[data-import-filter]"));
    queueSummary();
  });

  // Times arrive in UTC; show them in this browser's timezone and locale.
  const localTime = new Intl.DateTimeFormat(undefined, { dateStyle: "medium", timeStyle: "short" });
  htmx.onLoad((el) =>
    el.querySelectorAll("time[datetime]").forEach((t) => (t.textContent = localTime.format(new Date(t.dateTime)))),
  );

  // On every (re)connect, tell the hub where this page is and which live
  // views it shows, so it can send what was missed, or ask for a reload
  // after a restart.
  document.addEventListener("htmx:wsOpen", (e) => {
    gsm.socket = e.detail.socketWrapper;
    const st = document.getElementById("ws-state");
    gsm.send({ type: "hello", epoch: st.dataset.epoch, seq: Number(st.dataset.seq), watch: gsm.watching() });
    document.body.classList.remove("ws-down");
    document.dispatchEvent(new CustomEvent("gsm:wsopen"));
  });

  document.addEventListener("htmx:wsClose", () => {
    gsm.socket = null;
    document.body.classList.add("ws-down");
  });

  // The hub only sends JSON: live view patches, a reload (after a kick or
  // when catch-up is impossible), and note collaboration for the editors.
  document.addEventListener("htmx:wsBeforeMessage", (e) => {
    e.preventDefault();
    const m = JSON.parse(e.detail.message);
    if (m.type === "patch") gsm.applyPatch(m);
    else if (m.type === "reload") window.location.reload();
    else document.dispatchEvent(new CustomEvent("gsm:json", { detail: m }));
  });

  // A link-or-upload choice submits only the active tab's input, so a file
  // picked in a tab the GM left does not win over the link.
  document.addEventListener(
    "formdata",
    (e) => e.target.querySelectorAll('[data-source-field] [data-slot="tabs-content"][data-hidden] [name]').forEach((el) => e.formData.delete(el.name)),
    true,
  );

  // The image editor previews the active tab's link or file before it is
  // saved, falling back to the saved image.
  let previewObjectURL;
  function syncImagePreview() {
    const img = document.querySelector("[data-image-preview]");
    const panel = document.querySelector('[data-source-field] [data-slot="tabs-content"]:not([data-hidden])');
    if (!img || !panel) return;
    if (previewObjectURL) URL.revokeObjectURL(previewObjectURL);
    previewObjectURL = null;
    const file = panel.querySelector("input[type=file]")?.files[0];
    const link = panel.querySelector("input[type=url]");
    let src = img.dataset.imagePreview;
    if (file) src = previewObjectURL = URL.createObjectURL(file);
    else if (link?.value && link.checkValidity()) src = link.value;
    if (src) img.src = src;
    img.hidden = !src;
    document.querySelector("[data-image-preview-empty]").hidden = !!src;
  }
  ["input", "change"].forEach((type) =>
    document.addEventListener(type, (e) => {
      if (e.target.closest?.("[data-source-field]")) syncImagePreview();
    }),
  );
  // On window, so tabs.js has switched the panel first.
  window.addEventListener("click", (e) => {
    if (e.target.closest?.('[data-source-field] [data-slot="tabs-trigger"]')) syncImagePreview();
  });

  // Create-as-you-type comboboxes (image category, NPC field keys) take new
  // names: a create button shows while the text names no item and adds it as
  // an item, which it selects. Closing the list reverts the text to the
  // selection, so a name must become an item.
  const createButton = (input) => document.querySelector(`[data-combobox-create="${input.id}"]`);
  // The item named name (any case), if the list has one.
  function createItem(btn, name) {
    const items = btn.parentElement.querySelectorAll('[data-slot="combobox-item"]');
    return [...items].find((i) => i.getAttribute("data-templ-value").toLowerCase() === name.toLowerCase());
  }
  function syncCreate(input) {
    const btn = createButton(input);
    if (!btn) return;
    const name = input.value.trim();
    btn.hidden = name === "" || !!createItem(btn, name);
    btn.querySelector("span").textContent = `Create “${name}”`;
  }
  // pickCreate selects the item named by the input's text, creating it
  // first if it is new.
  function pickCreate(input) {
    const btn = createButton(input);
    const name = input.value.trim();
    if (!btn || name === "") return;
    let item = createItem(btn, name);
    if (!item) {
      const tpl = btn.parentElement.querySelector("template[data-combobox-create-item]");
      item = tpl.content.firstElementChild.cloneNode(true);
      item.id = `${input.id}-new-${tpl.parentElement.children.length}`;
      item.setAttribute("data-templ-value", name);
      item.setAttribute("data-templ-label", name);
      item.prepend(name);
      tpl.before(item);
    }
    btn.hidden = true;
    item.click();
  }
  ["input", "focusin"].forEach((type) =>
    document.addEventListener(type, (e) => {
      if (e.target.matches?.("[data-combobox-create-input]")) syncCreate(e.target);
    }),
  );
  document.addEventListener("click", (e) => {
    const btn = e.target.closest("[data-combobox-create]");
    if (btn) pickCreate(document.getElementById(btn.dataset.comboboxCreate));
  });
  // Item pickers load their options as the text is typed; once they arrive,
  // filter them against the text and update the create button.
  document.addEventListener("htmx:afterSwap", (e) => {
    const list = e.target;
    if (list.matches?.('[data-slot="combobox-list"]')) {
      document.querySelector(`[hx-target="#${CSS.escape(list.id)}"]`)?.dispatchEvent(new Event("input", { bubbles: true }));
    }
  });

  // Feature picker: picking a compendium feature (mouse or Enter, which
  // selects without a click) reloads the builder form filled in from it.
  // The open list is portaled out of the picker, so its items are looked up
  // in the document.
  document.addEventListener("change", (e) => {
    const picker = e.target.closest?.("[data-builder-pick]");
    if (!picker || e.target.getAttribute("aria-hidden") !== "true") return;
    const url = document.querySelector("[data-builder-load][data-selected]")?.dataset.builderLoad;
    if (url) htmx.ajax("GET", url, { target: "#builder", swap: "innerHTML" });
  });

  // NPC fields: picking a key in the add-field combobox adds a row for it
  // (or goes to the existing one); the picked item leaves the list. Rows are
  // numbered from the form's count, which Save reads.
  document.addEventListener("change", (e) => {
    const picker = e.target.closest?.("[data-add-field]");
    const key = e.target.value?.trim().toLowerCase();
    if (!picker || e.target.getAttribute("aria-hidden") !== "true" || !key) return;
    const form = picker.closest("form");
    const rows = form.querySelector("[data-field-rows]");
    let row = [...rows.children].find((r) => r.dataset.fieldKey === key);
    if (!row) {
      const count = form.elements.count;
      const i = Number(count.value);
      row = form.querySelector("template[data-field-row]").content.firstElementChild.cloneNode(true);
      row.dataset.fieldKey = key;
      row.querySelector("[data-field-label]").textContent = key;
      row.querySelectorAll("[name]").forEach((el) => (el.name = el.name.replace(/_\d+$/, `_${i}`)));
      row.querySelector('input[type="hidden"][name^="key_"]').value = key;
      rows.append(row);
      count.value = i + 1;
    }
    picker.querySelector(`[data-slot="combobox-item"][data-templ-value="${CSS.escape(key)}"]`)?.remove();
    e.target.value = "";
    picker.querySelector("[data-combobox-create-input]").value = "";
    row.scrollIntoView({ block: "nearest" });
    // fields.js mounts the row's editor first (raw mode shows the textarea).
    requestAnimationFrame(() => {
      const raw = form.hasAttribute("data-raw");
      row.querySelector(raw ? "[data-field-raw]" : ".field-editor .tiptap")?.focus();
    });
  });
  document.addEventListener("click", (e) => e.target.closest?.("[data-remove-field]")?.closest("[data-field-key]").remove());

  // The [[ picker of notes, NPCs and images to link: TipTap suggestion
  // render hooks for the editors in notes and NPC fields (wikilink.js).
  function suggestionMenu() {
    let el = null;
    let items = [];
    let index = 0;
    let command = null;

    const draw = () => {
      el.replaceChildren();
      if (!items.length) {
        const empty = document.createElement("div");
        empty.className = "px-2 py-1 text-sm text-muted-foreground";
        empty.textContent = "No matches";
        el.append(empty);
      }
      items.forEach((it, i) => {
        const b = document.createElement("button");
        b.type = "button";
        b.className = "block w-full rounded px-2 py-1 text-left text-sm" + (i === index ? " bg-accent" : "");
        b.className += " flex items-center justify-between gap-2";
        const name = document.createElement("span");
        name.className = "truncate";
        name.textContent = it.label;
        const kind = document.createElement("span");
        kind.className = "shrink-0 rounded border px-1 text-xs text-muted-foreground";
        kind.textContent = it.kind;
        b.append(name, kind);
        b.addEventListener("mousedown", (e) => {
          e.preventDefault();
          command(it);
        });
        el.append(b);
      });
    };
    const place = (rect) => {
      if (!rect || !el) return;
      el.style.left = `${rect.left + window.scrollX}px`;
      el.style.top = `${rect.bottom + window.scrollY + 4}px`;
    };
    const update = (p) => {
      items = p.items;
      command = p.command;
      index = 0;
      draw();
      place(p.clientRect && p.clientRect());
    };

    return {
      onStart: (p) => {
        el = document.createElement("div");
        el.className = "absolute z-50 w-64 rounded-md border bg-popover p-1 text-popover-foreground shadow";
        document.body.append(el);
        update(p);
      },
      onUpdate: update,
      onKeyDown: ({ event }) => {
        const n = Math.max(items.length, 1);
        if (event.key === "ArrowDown") index = (index + 1) % n;
        else if (event.key === "ArrowUp") index = (index - 1 + n) % n;
        else if (event.key === "Enter") {
          if (items[index]) command(items[index]);
          return true;
        } else if (event.key === "Escape") {
          el?.remove();
          el = null;
          return true;
        } else return false;
        draw();
        return true;
      },
      onExit: () => {
        el?.remove();
        el = null;
      },
    };
  }
  gsm.suggestionMenu = suggestionMenu;

  document.addEventListener("keydown", (e) => {
    // Runs before combobox.js, which picks a highlighted item on Enter;
    // without one, Enter picks the typed name instead of submitting the form.
    if (e.key === "Enter" && e.target.matches?.("[data-combobox-create-input]") && !document.querySelector('[data-slot="combobox-item"][data-highlighted]')) {
      e.preventDefault();
      pickCreate(e.target);
    }
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === "k") {
      e.preventDefault();
      gsm.openSwitcher();
    }
  });
})();
