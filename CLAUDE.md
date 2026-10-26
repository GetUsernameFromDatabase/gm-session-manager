# gm-session-manager

A LAN-hosted helper for tabletop RPG game masters and their players. The GM's PC runs a single Go binary; players on the same network join by address or QR code. All data lives on the host in SQLite. GMs manage campaigns, combat, NPCs and the compendium; players manage their characters; everyone shares real-time collaborative notes.

Out of scope: accounts or email, hosting on the internet, S3 or Docker, remote play, in-app dice rolling (physical dice only; the one exception is the GM's initiative roll in combat), maps beyond image attachments, device blocklists, offline mode. Ask before adding any of these.

## Stack

- Go, no CGO; tools (`templ`, `shadcn-templ`) are `tool` deps in `go.mod`. Router: gin.
- SQLite via `modernc.org/sqlite`, WAL mode, `PRAGMA foreign_keys = ON` on every connection. All queries through `sqlc`; the migrations in `db/migrations` are the schema. FTS5 tables live in a migration, with a stub for sqlc in `db/sqlc_stubs.sql`.
- Views: `templ` with [shadcn-templ](https://shadcn-templ.com/llms.txt) components, Tailwind (standalone CLI), `htmx` with the websocket extension.
- Notes: TipTap/ProseMirror with the collab plugin (`web/assets/js/notes.js`). NPC fields reuse TipTap (`fields.js`); both share `wikilink.js`, which renders `[[kind:id|label]]` tokens (parsed server-side in `web/views/wikilinks.go`).
- **No node or npm.** Browser libraries are vendored into `web/assets/vendor` by `cmd/vendorjs`. Everything ships in one binary via `go:embed`.
- HTTPS only, with a self-signed certificate generated in `DATA_DIR/tls` (`join/`).

## Layout

There is **no `internal/` directory. Do not create one.**

```
cmd/gm-session-manager/   main: config, DB, hub, HTTPS server
cmd/vendorjs/             downloads pinned browser JS into web/assets/vendor
config/                   .env loading and fail-fast validation
db/                       migrations/, queries/ (sqlc input), sqlc/ (generated), Open, Migrate, Store.InTx
auth/                     session tokens, cookie, middleware, admin and loopback checks
perm/                     every permission and visibility rule (characters, combatants, note blocks, entities)
hub/                      websocket hub: (epoch, seq) events, catch-up ring, live-view patches, pending join requests
dnd5e/                    5e sheet, commands (Parse/Apply), slots, level-up
compendium/               Open5e importer, search, homebrew
export/                   campaign zip (notes as Markdown, characters as JSON)
join/                     LAN addresses, TLS certificate, QR code
web/                      router, *_handlers.go, render.go (view models), live.go (live views, /live)
web/views/                templ pages and components; live.go holds live-view keys and binds
web/components/, utils/   shadcn-templ CLI output: never edit by hand
web/assets/               js/ (app.js, live.js, notes.js, fields.js, wikilink.js), css/ (globals.css → output.css), vendor/
data/                     DATA_DIR at runtime (database, tls, attachments, compendium); not code
```

## Rules that must not be broken

- **Visibility is enforced on the server.** Never send hidden data to a browser and hide it client-side. Unrevealed entity fields are stripped entirely (keys included). GM and admin queries never return `owner` note blocks.
- **Websocket data is as small as possible.** Send only what changed, as JSON data, never HTML. Values that change often are bound in the page (`data-bind`) and patched in place. When a view changes wholesale (a whole sheet, new or removed rows, a level-up), send a refetch order and let the page load the view's HTML from `/live` over HTTP; do not push the changes. Never drop shadcn components to save bytes.
- **Commands, not saved sheets.** Clients send small commands (`hp -5`, `toggle prone`). The server applies each in a transaction and publishes topics; the hub turns them into per-recipient patches. Commands are deltas, never whole-object overwrites: several devices may hold one character.
- **Session tokens:** 32 random bytes; store only the SHA-256 hash; cookie is `HttpOnly`, `SameSite=Lax`, `Secure`.
- **Admin:** a connection is admin if the real `RemoteAddr` is loopback (never trust `X-Forwarded-For`), or after login with `ADMIN_USERNAME` and `ADMIN_PASSWORD` from `.env`. Refuse to start if either is missing. Constant-time comparison. Never hardcode credentials.
- **Identity:** players have no passwords. One role per device (log out to switch); logging in while holding a session is refused. Logout and kick are the same thing: delete the `sessions` row. Pending join requests live in memory only. A character may have several devices, each approved separately.
- **Naming:** the `sessions` table is the login/device token. Game sessions are `game_sessions` everywhere (notes `kind`, UI, code). Never mix the two.
- **Notes:** each note block is its own collaborative document with its own audience; the server is the authority (stale steps get a 409 and rebase). Trusting client snapshots, text and links is accepted on a LAN of friends.
- **Attachments** are served only through the authorized handler, never as public static files.
- **Compendium:** normalize source data into our own shape; never leak a source's shape into the app. Keep each entry's license and attribution (the 2024 SRD is CC-BY 4.0).

### Permissions

| Thing | GM | Player |
|---|---|---|
| Own character | n/a | full edit (via commands) |
| Other characters | view; edit HP and conditions only | none |
| Monster HP in combat | always | hidden unless the GM toggles `hp_visible` |
| Entities | all fields | revealed entities, revealed fields only |

Note blocks (`perm/perm_test.go` mirrors this table):

| Visibility | Who reads | Who edits |
|---|---|---|
| `gm` | GM | GM |
| `players` | GM and every player in the campaign | GM and the note's owner; other players only if `players_can_edit` |
| `listed` | GM and the characters in its audience (empty audience = GM only: fail closed) | GM only |
| `owner` | the owning player only; GM and admin never | the owning player |

Any player may add a `players` block to the GM's `game_session` notes. A block's author may always edit it, and may delete it or change its access while they can still read it. Reordering stays with the note's manager.

## Live views

A page region that updates in real time is a live view.

1. Give its root `live(kind, id...)` (kinds and helpers in `web/views/live.go`). The page sends its `data-live` keys in the websocket hello.
2. Add the kind to `liveAffects` (which topics change it) and `loadLive` (permission checks, HTML) in `web/live.go`.
3. For values that change often, give the view model `Hot` (adds exactly the keys the markup binds) and `Shape` (the same model with those values fixed) in `web/views/live.go`, and mark the elements with `bind("text=key")`, `bind("aria-pressed=key")`, etc. Everything else is covered by the hash of `Shape`'s markup.
4. Cover it in `TestLiveVisibility` and `TestLiveBindsMatchState` (`web/live_test.go`).

## UI conventions

- Use shadcn-templ components for every control; add missing ones with `go tool shadcn-templ add <name>`. Reuse the helpers: `formSelect`/`selectItems` (combat.templ), `checkField` (notes.templ), `inlineField` (compendium.templ), `removeButton`.
- Show dates with `@localTime(t)` (layout.templ), never `Format` in a template: SQLite stores timestamps in UTC and app.js shows them in the viewer's timezone and locale.
- Style toggled states with `aria-pressed:` or `data-[…]:` variants, so a live update flips one attribute instead of swapping classes.
- Slow forms show a loading indicator: mark the form `data-busy` with class `group/busy` and put `@busyIndicator("…")` inside (layout.templ). app.js sets `aria-busy` while it submits or runs htmx requests, and disables its submit buttons against double submits.
- Previews fetch nothing while their place is not visible (a pane hidden on small screens, a collapsed section): check `checkVisibility()` before loading, as the compendium and import previews do; lists of images use `loading="lazy"`.
- New Tailwind classes need `go generate ./web` (rebuilds `output.css`).
- Never hand-edit `web/components`, `web/utils`, `*_templ.go` or `db/sqlc`.

## Code conventions

- Keep code DRY and easy to extend; prefer a small number of well-named helpers over repetition. Match the surrounding code's naming and comment density: a one-line comment for the *why*, not the *what*.
- Use constants for cookie, context, header and event keys, not string literals.
- Validate request and session structs with `go-playground/validator` struct tags instead of ad hoc per-field checks.
- Table-driven tests for all permission and visibility logic.
- `golangci-lint run` must stay clean (`.golangci.yml`: gofmt, goimports with local prefix `gm-session-manager`, revive, gosec, bodyclose, unparam, …).

## Commands

Generated code and built assets are committed, except `web/assets/vendor` and the shadcn-templ script bundle `web/assets/js/shadcn-templ-*.js` (both gitignored; `web/components/scripts_bundle.go` is committed): run `go run ./cmd/vendorjs` and `go tool shadcn-templ bundle` once after cloning, then `go build` alone works. After changing sources:

| Change | Command |
|---|---|
| `db/migrations`, `db/queries` | `go generate ./db` (sqlc on PATH) |
| `*.templ`, Tailwind classes | `go generate ./web` (script bundle + templ + Tailwind standalone CLI on PATH) |
| shadcn-templ components | `go tool shadcn-templ add <name>`, then `go generate ./web` |
| browser JS versions | `go run ./cmd/vendorjs -check` lists newer releases; edit `cmd/vendorjs/main.go`, then `go run ./cmd/vendorjs` |

Then `go build ./...`, `go test ./...`, `golangci-lint run` and `graphify update .`. Run the app with `./run.sh` (or `go run ./cmd/gm-session-manager`); see README.md for setup.

## graphify

This project has a knowledge graph at graphify-out/ with god nodes, community structure, and cross-file relationships.

Rules:
- When graphify-out/graph.json exists, run `graphify query "<question>"` before your first grep, search or source-file read for a codebase question, and put this rule in every subagent prompt that explores code. Read raw files only for lines the graph or an earlier read already located. Use `graphify path "<A>" "<B>"` for relationships and `graphify explain "<concept>"` for focused concepts. These return a scoped subgraph, usually much smaller than GRAPH_REPORT.md or raw grep output.
- If graphify-out/wiki/index.md exists, use it for broad navigation instead of raw source browsing.
- Read graphify-out/GRAPH_REPORT.md only for broad architecture review or when query/path/explain do not surface enough context.
- After modifying code, run `graphify update .` to keep the graph current (AST-only, no API cost).
- A committed graphify-out/graph.json is rewritten whole on every build, so a large diff to it in a pull request is expected. Do not flag it as churn or suggest skipping the graph rebuild or its hooks.
