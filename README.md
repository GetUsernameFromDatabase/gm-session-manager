<!-- omit in toc -->
# gm-session-manager

A LAN host for a tabletop RPG table. The GM's PC runs one binary; players
join from phones on the same network by address or QR code. All data stays
on the host in SQLite. GMs manage campaigns, combat, NPCs and the
compendium; players manage their characters; everyone shares real-time
collaborative notes. See `CLAUDE.md` for the architecture, rules and
conventions.

<!-- USER SECTION START: written by the developer. Claude and other agents: never edit, move or delete anything between these two markers. -->
My first project made with claude
<!-- USER SECTION END -->

- [Quick start (from a release)](#quick-start-from-a-release)
  - [First visit on each device](#first-visit-on-each-device)
  - [Windows firewall](#windows-firewall)
  - [macOS](#macos)
  - [GM on another device](#gm-on-another-device)
- [Configuration](#configuration)
- [Compendium (2024 SRD)](#compendium-2024-srd)
- [Development](#development)
  - [Prerequisites](#prerequisites)
  - [Run from source](#run-from-source)
  - [After changing sources](#after-changing-sources)
  - [Knowledge graph (graphify)](#knowledge-graph-graphify)
- [Releasing on GitHub](#releasing-on-github)
  - [Each release](#each-release)
  - [Manual fallback](#manual-fallback)

## Quick start (from a release)

1. Download the archive for your system from the
   [Releases](../../releases) page and unpack it into its own folder.
2. Copy `.env.example` to `.env` in that folder and set `ADMIN_USERNAME`
   and `ADMIN_PASSWORD` (the server refuses to start without them).
3. Run `gm-session-manager` (`gm-session-manager.exe` on Windows) from that
   folder. `.env` and `data/` are looked up in the current directory.
4. Open `https://localhost:8080` on the host PC. Localhost is always the GM.
   The page shows the join address and a QR code for players.

### First visit on each device

The host uses a self-signed certificate (in `DATA_DIR/tls`, renewed when a
network address changes or it nears expiry). Each browser shows a warning
once per certificate: choose "Advanced" and continue. The GM page shows the
certificate's SHA-256 fingerprint if you want to compare it with what the
browser shows.

### Windows firewall

The first time the host listens, Windows asks whether to allow
`gm-session-manager.exe` on networks. Allow **Private networks** (your home
Wi-Fi). If you dismissed the prompt, players cannot connect: open
"Windows Defender Firewall → Allow an app through firewall" and tick
Private for gm-session-manager, or run once as administrator:

```bat
netsh advfirewall firewall add rule name="gm-session-manager" dir=in action=allow protocol=TCP localport=8080 profile=private
```

`go run` builds a new temporary executable each time, so Windows may ask
again; a built binary keeps its permission.

### macOS

Release binaries are not signed. On first run, right-click the binary and
choose "Open", or run `xattr -d com.apple.quarantine gm-session-manager`.

### GM on another device

Log in at `/admin/login` with the `.env` credentials. A device holds one
role: to switch between player and GM, log out and log in again.

## Configuration

Settings come from `.env` in the working directory; real environment
variables override it.

| Variable | Default | Notes |
| --- | --- | --- |
| `ADMIN_USERNAME` | — | required |
| `ADMIN_PASSWORD` | — | required |
| `PORT` | `8080` | HTTPS port |
| `DATA_DIR` | `data` | everything the app stores |

`DATA_DIR` holds the database (`gm-session-manager.db`), the certificate
(`tls/`), uploaded files (`attachments/`) and compendium sources
(`compendium/`). Back up this folder to keep your campaigns; stop the
server first, or copy the `-wal` and `-shm` files with the database.

## Compendium (2024 SRD)

The importer reads data folders from `DATA_DIR/compendium/<source>/`. For
the 2024 SRD, download the JSON files of
[open5e-api `data/v2/wizards-of-the-coast/srd-2024`](https://github.com/open5e/open5e-api/tree/staging/data/v2/wizards-of-the-coast/srd-2024)
into `data/compendium/srd-2024/`. They are imported at start-up and from the
Compendium page ("Import now"); unchanged folders are skipped. The SRD 5.2
is licensed CC-BY-4.0 and every entry shows its attribution.

Homebrew is written in the Compendium page and is shared by all campaigns.

## Development

### Prerequisites

- [Go](https://go.dev/dl/) at the version in `go.mod`. That is all you need
  to build and run: generated code and built assets are committed, except the
  browser libraries and the component script bundle: run
  `go run ./cmd/vendorjs` and `go tool shadcn-templ bundle` once after cloning.
- Only for changing sources:
  - [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) on PATH
    (migrations and queries),
  - the [Tailwind standalone CLI](https://tailwindcss.com/blog/standalone-cli)
    on PATH as `tailwindcss` (templates and classes),
  - [golangci-lint](https://golangci-lint.run/welcome/install/).
  - `templ` and `shadcn-templ` are Go tool dependencies (`go tool …`), no
    install needed.
- No node or npm. Browser libraries are vendored by `cmd/vendorjs`.

### Run from source

```sh
cp .env.example .env    # then set ADMIN_USERNAME and ADMIN_PASSWORD
./run.sh                # or: go run ./cmd/gm-session-manager
```

### After changing sources

| Change | Command |
| --- | --- |
| `db/migrations`, `db/queries` | `go generate ./db` (sqlc) |
| `*.templ`, Tailwind classes | `go generate ./web` (script bundle + templ + Tailwind) |
| shadcn-templ components | `go tool shadcn-templ add <name>`, then `go generate ./web` |
| browser JS versions | `go run ./cmd/vendorjs -check` lists newer releases; edit `cmd/vendorjs/main.go`, then `go run ./cmd/vendorjs` |

Then `go build ./...`, `go test ./...` and `golangci-lint run`.

`GSM_COMPENDIUM_DIR=data/compendium go test ./compendium -run TestRealSource -v`
checks the importer against a real data download.

### Knowledge graph (graphify)

Claude Code uses [graphify](https://github.com/Graphify-Labs/graphify) to
query the codebase (see `CLAUDE.md`). The graph is committed
(`graphify-out/graph.json`, `GRAPH_REPORT.md`, `graph.html`); the rest of
`graphify-out/` is a local cache. Optional; not needed to build or run.

1. Install [uv](https://docs.astral.sh/uv/), then graphify at the version in
   `.claude/skills/graphify/.graphify_version`:

   ```sh
   uv tool install graphifyy==<version>   # e.g. graphifyy==0.9.95
   ```

2. After changing code, refresh the graph and commit it with your change
   (AST only, no API cost):

   ```sh
   graphify update .
   ```

3. For a full rebuild (including the Markdown docs), run `/graphify .` in
   Claude Code.

To update graphify:

1. Upgrade the tool: `uv tool upgrade graphifyy`.
2. Refresh the project's skill, `CLAUDE.md` section and hooks:

   ```sh
   graphify claude install --project
   ```

3. Review `git diff`: it rewrites `.claude/skills/graphify/` (including
   `.graphify_version`), the `## graphify` section of `CLAUDE.md` and the
   hooks in `.claude/settings.json`. A locally edited `SKILL.md` is kept as
   `SKILL.md.bak`; delete it once reviewed.
4. Rebuild the graph with the new version (`graphify update .`) and commit
   everything together.

See the [graphify README](https://github.com/Graphify-Labs/graphify) for
everything else (queries, hooks, exports).

## Releasing on GitHub

`build.sh` is the one build process, for you and CI. It fetches the vendored
JS, builds the shadcn-templ script bundle (both are gitignored), runs the
tests and stamps the version (`git describe`, or `$VERSION`) into the binary:

```sh
./build.sh           # this machine → dist/gm-session-manager[.exe]
./build.sh release   # every target → dist/ archives + SHA256SUMS (needs zip)
```

Pushing a `v*` tag runs `.github/workflows/release.yml`: it runs
`golangci-lint`, then `build.sh release` (Windows, Linux and macOS binaries,
no CGO, so every target builds from one runner, each packed with
`README.md` and `.env.example`), and publishes a GitHub release with
generated notes. `.env` and `data/` are git-ignored and never shipped.

### Each release

1. If sources changed, regenerate (`go generate ./db ./web`) and commit.
2. Check: `go build ./...`, `go test ./...`, `golangci-lint run`.
3. Tag and push:

   ```sh
   git tag v0.1.0
   git push origin v0.1.0
   ```

4. Watch the run on the repository's Actions tab (`gh run watch`); the
   release appears under Releases when it finishes.

### Manual fallback

Build and upload it yourself:

```sh
VERSION=v0.1.0 ./build.sh release
gh release create v0.1.0 dist/* --generate-notes
```
