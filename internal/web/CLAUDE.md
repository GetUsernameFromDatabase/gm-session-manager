# internal/web

Changes to `.templ`, CSS or JS can be checked with the `browser-check` skill (`check: <pages and controls>`; before/after screenshots: `pair Cnn: …`). Ask with AskUserQuestion before running it, unless the prompt already says to run checks (e.g. a batch of HANDOFF items). Before editing for a UI change, run `./run.sh before` so before shots show the code as it was. When implementing a list of UI changes, give each an ID (Cnn, continuing from `data-test/screenshots/README.md` or the ID HANDOFF.md names) and tag its HANDOFF item. Don't drive the browser from the main session.

## Websocket and live views

Send only what changed, as JSON. Often-changing values are bound in the page (`data-bind`) and patched in place. When a view changes wholesale (a whole sheet, new or removed rows, a level-up), send a refetch order and let the page load the view's HTML from `/live`. Never drop shadcn components to save bytes.

A page region that updates in real time is a live view:

1. Give its root `live(kind, id...)` (`views/live.go`). The page sends its `data-live` keys in the websocket hello.
2. Add the kind to `liveAffects` (which topics change it) and `loadLive` (permission checks, HTML) in `live.go`.
3. For often-changing values, give the view model `Hot` (exactly the keys the markup binds) and `Shape` (the model with those values fixed) in `views/live.go`; mark elements with `bind("text=key")`, `bind("aria-pressed=key")`, etc. The hash of `Shape`'s markup covers everything else.
4. Cover it in `TestLiveVisibility` and `TestLiveBindsMatchState` (`live_test.go`).

## Notes

Each note block is its own collaborative document (TipTap, `assets/js/notes.js`) with its own audience; the server is the authority (stale steps get a 409 and rebase). Trusting client snapshots, text and links is accepted on a LAN of friends. NPC fields reuse TipTap (`fields.js`); both share `wikilink.js` for `[[kind:id|label]]` tokens (parsed in `views/wikilinks.go`).

## UI conventions

- shadcn-templ (shadcn-templ.com, github.com/axadrn/shadcn-templ), not templUI. Use its components for every control, never native `<select>`/`<dialog>`/hand-styled inputs; add missing ones with `go tool shadcn-templ add <name>`. Its Select and popovers portal to `<body>`, so they fail inside a native `<dialog>`: use the shadcn Dialog.
- Reuse the helpers: `formSelect`/`selectItems` (combat.templ), `checkField` (notes.templ), `inlineField` (compendium.templ), `removeButton`.
- Dates: `@localTime(t)` (layout.templ), never `Format` in a template (UTC in SQLite, app.js localizes).
- Addresses: every view state lives in the URL, so a copied address shows the same view to whoever opens it. The compendium's campaign is in the path (`/campaigns/:campaign/compendium/...`; `/compendium` is the whole one): build links with `CompendiumPath(ctx, …)`/`entryHref` in templates (`Page` and the route put the campaign on ctx) and `compendiumPath(c, …)` in handlers, never a literal `/compendium`. Files works the same way (`/campaigns/:campaign/files`, `/files` for all; `FilesBase`); both share `campaignSwitcher` (`Layout.Switch`).
- Toggled states: `aria-pressed:` or `data-[…]:` variants, so a live update flips one attribute.
- Capped checkbox groups: `data-max-checked="N"` on the element holding the boxes (count 1 is a radio group instead). app.js counts its named, enabled, ticked boxes (known ones have no name, taught ones a disabled input) and at N greys the unticked out via the checkbox root's `aria-disabled`, never the input's `disabled`, which `data-hide-for` owns. Give it an `id` and `#<id>-count` shows the live count. Browser help only: the server check stays.
- Slow forms: `data-busy` + class `group/busy` + `@busyIndicator("…")` inside; app.js sets `aria-busy` and disables submit buttons.
- Previews fetch nothing while hidden: check `checkVisibility()` first; image lists use `loading="lazy"`.
- Long lists and notices (books, sources, import summaries) scroll inside a capped box (`scrollBox` in `views/fivetools.go`), so the controls around them stay close together.
- Narrow screens: a grid column needs `min-w-0` before a child's `max-w-full` or truncation works; in a `flex-wrap` row of icon + text + button, give the text `min-w-0 grow basis-48`.
- `sr-only` (or any absolute) text in rows of an `overflow-*-auto` list needs a `relative` row, or the page scrolls far past the body.

## Handlers

A handler with an error path returns `error` and is registered as `a.h(a.x)`: a `userError` is a 422, `sql.ErrNoRows` a 404, anything else a 500. Access helpers that answer themselves (`characterAccess`, `idParam`, `bindForm`, …) stay `ok`-returning; the handler then does `return nil`. Never return a non-nil error after writing a response.

## Gotchas

- `htmx.config.defaultSwapStyle` is `none`: every form that swaps HTML needs `hx-swap`. Never swap a fragment holding a `<form>` into another form.
- templ trims the space before a component call on the next line ("uploaded26 Oct"): write `{ "Updated " }` or keep both on one line. A `{ expr }` ending a line gets a space after it: join adjacent expressions.
- templ 0.3.1100: a spread event handler given a plain string renders blank; wrap it in `templ.JSUnsafeFuncCall(...)`.
- `TestPagesRender` fails on `templ.FailedSanitizationJS`.
