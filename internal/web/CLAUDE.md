# internal/web

Changes to `.templ`, CSS or JS can be checked with the `browser-check` skill. Ask with AskUserQuestion before running it, unless the prompt already says to run checks (e.g. a batch of HANDOFF items). Its cost is turns, so name exact targets and it won't explore: `check: <URL> as GM|player|state: <name> — <selectors to work> → <expected result>`. Before/after pairs (`pair <name>: <element selector> on <URL>, as GM|player`) only when the user asks for them; then run `./run.sh before` before editing, and view the PNGs yourself (the agent doesn't). Don't drive the browser from the main session.

## Values Go and JS share

A value both sides need (limits, link schemes, route patterns, the wikilink pattern, websocket message types, JSON keys) lives in Go and reaches JS only through `/assets/js/shared.js`, an ES module `sharedjs.go` builds at start-up: add it to `sharedExports` and to `TestSharedJS`'s `want`, run `go test ./internal/web -run TestSharedJSTypes -update` (editor types in `.vscode/shared.d.ts`, read through the root `jsconfig.json`), then `import { X } from "./shared.js"`. All our JS files are modules (app.js and live.js too, deferred after the classic htmx scripts), and import by relative paths (`./wikilink.js`, `../vendor/…`), never from `/`: an editor reads `/` as the drive root, and `TestSharedImportsExist` checks only `./` imports and refuses `/` ones. A validator tag cannot name a constant: keep the literal and check it in `TestSharedLimitsMatchTags`. Likewise a field holding a compendium source key (forms here, sheet fields in `dnd5e`) takes `max=512` and goes into `TestSourceKeyLimits`. A route JS calls is a constant in `internal/route` (a group root, or a root plus a part within it), and the router registers it from those constants, in its group (`TestSharedRoutesRegistered`); JS fills it with `routePath` (wikilink.js), Go with `route.Path`. Build every URL with `route.Path(route.X, ids…)`, never a literal string: not as a part appended to a built path (`base + "/rename"`, `base + "reveal"`), not inside an HTML or JS string. Appending a route constant to a built base (`base + route.Rename`) is fine; `TestNoRouteLiterals` enforces it over Go, templ and our JS. The wikilink pattern is shared as source, so it uses only syntax RE2 and JS read alike (`(?<name>…)`, no flags; `TestWikiLinkPatternIsJSSafe`). Values the server already renders into the page (`data-*`, `ws-connect`) stay there.

## Websocket and live views

Send only what changed, as JSON. Often-changing values are bound in the page (`data-bind`) and patched in place. When a view changes wholesale (a whole sheet, new or removed rows, a level-up), send a refetch order and let the page load the view's HTML from `/live`. Never drop shadcn components to save bytes.

A page region that updates in real time is a live view:

1. Give its root `live(kind, id...)` (`views/live.go`). The page sends its `data-live` keys in the websocket hello.
2. Add an entry to `liveKinds` in `live.go`: the topic kinds that change it (`own` for its id, `any` for any id) and a `load…Live` method (permission checks, HTML). A topic that must cross campaigns also goes in `everyCampaign` or `gmEverywhere` (`hub/hub.go`).
3. For often-changing values, give the view model `Hot` (exactly the keys the markup binds) and `Shape` (the model with those values fixed) in `views/live.go` (the sheet's in `views/character_live.go`); mark elements with `bind("text=key")`, `bind("aria-pressed=key")`, etc. The hash of `Shape`'s markup covers everything else. A heavy view (sheet, GM tracker) sets `liveView.model` instead: `Shape` is hashed as JSON, skipping the render, so its markup must read nothing outside the model, and the model should hold only what the page shows (rows offered by name are `views.Named`, not whole `sqlc` rows), or an unseen change refetches.
4. Cover it in `TestLiveVisibility` and `TestLiveBindsMatchState` (`live_test.go`); a `model` view also in `modelViews` (`TestLiveHotKeepsHash`, `TestLiveModelHashFollowsMarkup`) and `TestLiveModelsEncodeWhole`.

A region with an `id` and `data-reload-on="<view key>"` follows a live view: each refetch of that view reloads it (live.js `reload`) from `data-reload-url`, or the page address with its tab's query. It is answered out of band when the region is itself an `hx-swap-oob` root (Browse), swapped in place otherwise (the 5etools picker). A form sends its fields, and `data-reload-rows="<param>"` sends how many `[data-row]` it shows; the reload waits while focus is inside. The marker view `compendium:<campaign>` (`CompendiumLive`, a nonce model) exists only to drive Browse; keep a marker outside the region it drives, or the region's reload carries it and loops.

The htmx websocket extension retries forever on an abnormal close (1006, which includes an HTTP error at the upgrade) and stops on a normal one (1000). To turn a page's socket away, accept it, send a control message and close normally (`Hub.ServeGone`), not a 401.

## Notes

Each note block is its own collaborative document (TipTap, `assets/js/notes.js`) with its own audience; the server is the authority (stale steps get a 409 and rebase). Trusting client snapshots, text and links is accepted on a LAN of friends, except link schemes: raw steps skip TipTap's link check, so `submitSteps` and `blockSnapshot` refuse an `href` outside `safeSchemes` and `linkClick` follows only those. NPC fields reuse TipTap (`fields.js`); both share `wikilink.js` for `[[kind:id|label]]` tokens (parsed in `views/wikilinks.go`).

## UI conventions

- shadcn-templ (shadcn-templ.com, github.com/axadrn/shadcn-templ), not templUI. Use its components for every control, never native `<select>`/`<dialog>`/hand-styled inputs; add missing ones with `go tool shadcn-templ add <name>`. Its Select and popovers portal to `<body>`, so they fail inside a native `<dialog>`: use the shadcn Dialog.
- A `.templ` holds only imports and `templ` components: its types, consts, vars and helper funcs go in `<name>_view.go` beside it (templ copies them into `_templ.go`; `TestTemplHasNoGoDecls`).
- Reuse the helpers: `formSelect` (combat.templ)/`selectItems` (combat_view.go), `checkField` (notes.templ), `inlineField` (compendium.templ), `removeButton`.
- Dates: `@localTime(t)` (layout.templ), never `Format` in a template (UTC in SQLite, app.js localizes).
- Addresses: every view state lives in the URL, so a copied address shows the same view to whoever opens it. The compendium's campaign is in the path (`/campaigns/:campaign/compendium/...`; `/compendium` is the whole one): build links with `CompendiumPath(ctx, …)`/`entryHref` in templates (`Page` and the route put the campaign on ctx) and `compendiumPath(c, …)` in handlers, never a literal `/compendium`. Files works the same way (`/campaigns/:campaign/files`, `/files` for all; `FilesBase`); both share `campaignSwitcher` (`Layout.Switch`).
- Toggled states: `aria-pressed:` or `data-[…]:` variants, so a live update flips one attribute.
- Capped checkbox groups: `data-max-checked="N"` on the element holding the boxes (count 1 is a radio group instead). app.js counts its named, enabled, ticked boxes (known ones have no name, taught ones a disabled input) and at N greys the unticked out via the checkbox root's `aria-disabled`, never the input's `disabled`, which `data-hide-for` owns. Give it an `id` and `#<id>-count` shows the live count. Browser help only: the server check stays.
- Command boxes (one `cmd` input sending delta commands): `form[data-command=url]` with an `id`, no `hx-post`/`reset()`. app.js empties the box on send, sends in order, ignores an empty Enter and restores a refused command. A success `reset()` would wipe the next command typed meanwhile.
- Slow forms: `data-busy` + class `group/busy` + `@busyIndicator("…")` inside; app.js sets `aria-busy` and disables submit buttons.
- Previews fetch nothing while hidden: check `checkVisibility()` first; image lists use `loading="lazy"`.
- Long lists and notices (books, sources, import summaries) scroll inside a capped box (`scrollBox` in `views/import_view.go`), so the controls around them stay close together.
- Narrow screens: a grid column needs `min-w-0` before a child's `max-w-full` or truncation works; in a `flex-wrap` row of icon + text + button, give the text `min-w-0 grow basis-48`.
- `sr-only` (or any absolute) text in rows of an `overflow-*-auto` list needs a `relative` row, or the page scrolls far past the body.

## Handlers

A handler with an error path returns `error` and is registered as `a.h(a.x)`: a `userError` is a 422, `sql.ErrNoRows` a 404, anything else a 500. Access helpers that answer themselves (`characterAccess`, `idParam`, `bindForm`, …) stay `ok`-returning; the handler then does `return nil`. Never return a non-nil error after writing a response.

- A command reads the row it changes inside its own transaction, never from a read made before it. A button whose meaning depends on what the page showed (combat Next/Previous) posts that state, kept current by the live view, and the server refuses a mismatch, so a repeated or stale press does nothing.
- A refused plain form post shows a bare error page and loses what was typed: editors post through htmx (`hx-encoding="multipart/form-data"` for files) and answer with `redirect(c, …)`, not `c.Redirect`.
- The request log keeps POST bodies except where `keepsBody` (`reqlog.go`, matched on the route pattern) says no. A new route whose body can carry a player's private text (note titles, blocks) goes there, or that text reaches the GM's reports.
- Browsers post every line break as `\r\n`, hidden inputs included; Go's CSV reader stores `\n`. Normalize before comparing a posted value with a stored one.
- A hidden input's `defaultValue` always equals its `value`, so live.js `edited` and any browser-side "what changed" check cannot see edits in controls backed by one (shadcn select/combobox, NPC field values): keep the loaded value yourself (`data-saved`, `was_*`).
- An effect (`conditions.held_by`) dies with its concentration row by FK cascade, and its owner is not in the request: before any delete that can cascade, call `noteHeld(ctx, q, ids...)` in the same transaction (`held.go`) and publish with `a.changed`/`a.publish`, which add the collected topics. Effect modifiers are merged only for display (`effectSheet`: `shownSheetView`, tracker rows), never into the stored sheet.
- In tests, loopback is the GM: a visitor or player request needs a LAN `RemoteAddr` (`lanAddr`).

## Gotchas

- htmx inherits `hx-include` (and `hx-target`, `hx-swap`) to child elements that make their own requests: a wrapper whose include names something a child loads later needs `hx-disinherit="hx-include"`, or htmx logs "returned no matches".
- shadcn buttons start `border border-transparent`: a visible (e.g. dashed) border on one needs its own colour class.
- A shadcn checkbox posts `"on"`, which gin cannot bind to a `bool`: bind a `string` with `validate:"omitempty,oneof=on"`.
- `htmx.config.defaultSwapStyle` is `none`: every form that swaps HTML needs `hx-swap`. Never swap a fragment holding a `<form>` into another form.
- templ trims the space before a component call on the next line ("uploaded26 Oct"): write `{ "Updated " }` or keep both on one line. A `{ expr }` ending a line gets a space after it: join adjacent expressions.
- templ 0.3.1100: a spread event handler given a plain string renders blank; wrap it in `templ.JSUnsafeFuncCall(...)`.
- templ 0.3.1100: `continue` in a template loop renders as the text "continue"; wrap the body in an `if`. `templ fmt` panics (nil pointer in `imports.containsImport`) on an unused import: remove it.
- NPC fields save themselves (`fields.js`, `field-save`), and live.js holds a refetch while the caret is in a rich-text editor, so the GM's own save never wipes the field they moved on to.
- `TestPagesRender` fails on `templ.FailedSanitizationJS`.
