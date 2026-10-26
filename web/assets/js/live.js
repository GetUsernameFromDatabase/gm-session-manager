// Live views. Each [data-live] root on the page is a view the hub keeps
// current: it sends only the values that changed, as JSON. Elements bind
// keys with data-bind="target=key ..." (target: text, value, show, or an
// attribute name). A change no bound element can show (new markup, a list
// row added or removed) refetches the view's HTML from /live instead.
(function () {
  "use strict";
  const gsm = window.gsm;

  function roots(view) {
    return document.querySelectorAll(`[data-live="${CSS.escape(view)}"]`);
  }

  // watching lists the live views on this page, for the hello message.
  gsm.watching = () => [...new Set([...document.querySelectorAll("[data-live]")].map((el) => el.dataset.live))];

  // bindings maps each key of a view to the elements showing it.
  function bindings(view) {
    const map = new Map();
    for (const root of roots(view)) {
      for (const el of [root, ...root.querySelectorAll("[data-bind]")]) {
        for (const spec of (el.dataset.bind || "").split(" ")) {
          const i = spec.indexOf("=");
          if (i < 0) continue;
          const key = spec.slice(i + 1);
          if (!map.has(key)) map.set(key, []);
          map.get(key).push({ el, target: spec.slice(0, i) });
        }
      }
    }
    return map;
  }

  function show(el, target, value) {
    switch (target) {
      case "text":
        el.textContent = value;
        break;
      case "value":
        if (document.activeElement !== el) el.value = value; // never under the cursor
        break;
      case "show":
        el.hidden = !value;
        break;
      default:
        el.setAttribute(target, String(value));
    }
  }

  // refetch reloads a view's fragments (out-of-band swaps). A change that
  // lands while a refetch is in flight fetches again afterwards.
  const inflight = new Map(); // view -> fetch again
  gsm.refetch = (view) => {
    if (inflight.has(view)) {
      inflight.set(view, true);
      return;
    }
    inflight.set(view, false);
    htmx.ajax("GET", "/live?v=" + encodeURIComponent(view), { source: document.body, swap: "none" }).finally(() => {
      const again = inflight.get(view);
      inflight.delete(view);
      if (again) gsm.refetch(view);
    });
  };

  // A refetched view replaces its markup; scrolled lists inside it keep
  // their place (data-keep-scroll, matched by id). The old target is still
  // in the document here, and the microtask runs after the swap, before paint.
  document.addEventListener("htmx:oobBeforeSwap", (e) => {
    for (const el of e.detail.target.querySelectorAll("[data-keep-scroll][id]")) {
      const { id, scrollTop } = el;
      queueMicrotask(() => {
        const now = document.getElementById(id);
        if (now) now.scrollTop = scrollTop;
      });
    }
  });

  gsm.applyPatch = (m) => {
    const st = document.getElementById("ws-state");
    if (st) {
      st.dataset.epoch = m.epoch;
      st.dataset.seq = m.seq;
    }
    for (const p of m.views || []) {
      const set = p.set || {};
      const map = p.refetch ? null : bindings(p.v);
      // A key nothing shows yet, or a shown key going away, means the
      // markup itself changed.
      if (!map || Object.keys(set).some((k) => !map.has(k)) || (p.del || []).some((k) => map.has(k))) {
        gsm.refetch(p.v);
        continue;
      }
      for (const [k, v] of Object.entries(set)) {
        for (const b of map.get(k)) show(b.el, b.target, v);
      }
    }
  };

  // Without a websocket the actor's own changes would not show: refetch
  // the page's views after each successful change instead.
  document.addEventListener("htmx:afterRequest", (e) => {
    const verb = e.detail.requestConfig && e.detail.requestConfig.verb;
    if (!gsm.socket && e.detail.successful && verb && verb !== "get") gsm.watching().forEach(gsm.refetch);
  });
})();
