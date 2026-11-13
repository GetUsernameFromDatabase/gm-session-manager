# gm-session-manager

A LAN-hosted helper for tabletop RPG game masters and their players. The GM's PC runs a single Go binary; players on the same network join by address or QR code. All data lives on the host in SQLite; app-wide settings live in `DATA_DIR/settings.json` (no migration for a switch).

Out of scope: accounts or email, hosting on the internet, S3 or Docker, remote play, in-app dice rolling (physical dice only; the one exception is the GM's initiative roll in combat), maps beyond image attachments, device blocklists, offline mode. Ask before adding any of these.

## Stack

- Go, no CGO, gin; tools (`templ`, `shadcn-templ`) are `tool` deps in `go.mod`.
- SQLite (`modernc.org/sqlite`) through `sqlc`. Views: `templ` + [shadcn-templ](https://shadcn-templ.com/llms.txt) + Tailwind standalone CLI + `htmx` websocket extension. Notes: TipTap collab.
- **No node or npm.** Browser libraries are vendored by `cmd/vendorjs`; everything ships in one binary via `go:embed`.
- HTTPS only, self-signed certificate in `DATA_DIR/tls`.
- Layout: [README.md](README.md#project-layout).

## Rules that must not be broken

- **Visibility is enforced on the server.** Never send hidden data to a browser and hide it client-side. Unrevealed entity fields are stripped entirely (keys included). GM and admin queries never return `owner` note blocks. Permission tables: `internal/perm/CLAUDE.md`.
- **Websocket data is as small as possible:** changed values as JSON, never HTML; wholesale changes send a refetch order (`internal/web/CLAUDE.md`).
- **Commands, not saved sheets.** Clients send small delta commands (`hp -5`, `toggle prone`), never whole-object overwrites: several devices may hold one character. The server applies each in a transaction and publishes topics.
- **Session tokens:** 32 random bytes; store only the SHA-256 hash; cookie is `HttpOnly`, `SameSite=Lax`, `Secure`.
- **Admin:** real `RemoteAddr` is loopback (never trust `X-Forwarded-For`), or login with `ADMIN_USERNAME`/`ADMIN_PASSWORD` from `.env` (refuse to start if missing; constant-time compare; never hardcode).
- **Identity:** players have no passwords. One role per device; logging in while holding a session is refused. Logout and kick both delete the `sessions` row. Pending join requests live in memory only. A character may have several devices, each approved separately.
- **Naming:** the `sessions` table is the login/device token; game sessions are `game_sessions` everywhere. Never mix the two.
- **Attachments** are served only through the authorized handler, never as public static files.

## Code conventions

- DRY, a few well-named helpers; match the surrounding naming and comment density (one-line *why* comments).
- Constants for cookie, context, header and event keys. Validate structs with `go-playground/validator` tags.
- Values a SQL CHECK lists are typed in `internal/enum` and wired through sqlc `overrides`; goconst guards other repeated strings. A family of values (one list, map or switch domain, numbers too) is all constants or none.
- Table-driven tests for all permission and visibility logic. Positional test-table rows break when a field is added; add a separate loop instead.
- `golangci-lint run` stays clean. Formatters always win: keep what `templ fmt` and `golangci-lint fmt` change, even in untouched files.
- Never hand-edit `internal/web/components`, `internal/web/utils`, `*_templ.go` or `internal/db/sqlc`.
- Durable lessons go in the nearest folder's `CLAUDE.md`; workflows in a skill; `HANDOFF.md` holds only current state and open items.

## Game systems

D&D 5e (2024) is the only system, fenced in rather than abstracted:

- **5e:** `internal/dnd5e` (rules), `internal/web/sheet5e_*.go` (wizard, level-up, builder, sources, sheet commands) and in `internal/web/views` the sheet, wizard, builder, level-up, sources and compendium templates, `character_live.go` and `sources.go`. The exceptions in the depguard rule `neutral` are the full list.
- **Neutral:** everything else (combat, sessions, campaigns, notes, entities, images, files, hub, perm, export). It reaches 5e only through `internal/web/sheet5e.go` and `internal/web/views/sheet5e.go`, plain-typed functions. depguard (`.golangci.yml`, rule `neutral`) fails the lint on any other `dnd5e` import, generated `*_templ.go` included; a new 5e file must be added to its exceptions.
- **A second system** starts as its own `sheet<sys>_*.go` beside `sheet5e`; only then is the shared interface extracted from the two. Existing hooks: `campaigns.system` and `characters.system`.
- **Largest open question:** `internal/compendium` is 5e-shaped (statblocks, spells, classes; the `compendium_entries.kind` CHECK), and is exempt from the fence.

## Commands

After cloning: `go run ./cmd/vendorjs` and `go tool shadcn-templ bundle` (both outputs gitignored). After changing sources:

| Change | Command |
|---|---|
| migrations, queries | `go generate ./internal/db` (sqlc on PATH) |
| `*.templ`, Tailwind classes, shadcn-templ components (`go tool shadcn-templ add <name>`) | `go generate ./internal/web` |
| Go formatting | `golangci-lint fmt` |
| browser JS versions | `go run ./cmd/vendorjs -check`, edit `cmd/vendorjs/main.go`, `go run ./cmd/vendorjs` |

Then `go build ./...`, `go test ./...`, `golangci-lint run`, `graphify update .`. Full build: `./build.sh`; release: `./release.sh <version>`; run: `./run.sh`.

## graphify

This project has a knowledge graph at graphify-out/ with god nodes, community structure, and cross-file relationships.

Rules:
- When graphify-out/graph.json exists, run `graphify query "<question>"` before your first grep, search or source-file read for a codebase question, and put this rule in every subagent prompt that explores code. Read raw files only for lines the graph or an earlier read already located. Use `graphify path "<A>" "<B>"` for relationships and `graphify explain "<concept>"` for focused concepts. These return a scoped subgraph, usually much smaller than GRAPH_REPORT.md or raw grep output.
- If graphify-out/wiki/index.md exists, use it for broad navigation instead of raw source browsing.
- Read graphify-out/GRAPH_REPORT.md only for broad architecture review or when query/path/explain do not surface enough context.
- After modifying code, run `graphify update .` to keep the graph current (AST-only, no API cost).
- A committed graphify-out/graph.json is rewritten whole on every build, so a large diff to it in a pull request is expected. Do not flag it as churn or suggest skipping the graph rebuild or its hooks.
