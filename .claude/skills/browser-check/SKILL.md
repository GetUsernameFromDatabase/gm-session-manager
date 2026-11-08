---
name: browser-check
description: Check UI changes in a real browser with the Playwright MCP server against a test instance of the app, as GM or player, and take before/after screenshot pairs. Use after any change to .templ, CSS or JS, before saying it is done, when a change needs before/after screenshots, and when asked to look at, test or explore pages in the browser.
context: fork
agent: general-purpose
model: haiku
argument-hint: "check: <pages and controls> | pair Cnn: <element> on <page>, as GM|player"
---

# Browser check

You run as your own agent and cannot see the caller's conversation. Your task: $ARGUMENTS

- `check: …`: check those pages and controls.
- `pair Cnn: …`: take a before/after screenshot pair; read `.claude/skills/browser-check/screenshots.md` first and follow it.
- If the task names no page, say so and stop.

## The test instance

- Start it once, as a background Bash job: `./run.sh browser`. It builds `build/gm-session-manager-test.exe`, serves `https://localhost:8443` from `data-test/` (never `data/`), and keeps `data-test/playwright` between runs, so screenshots a handoff refers to survive; `./run.sh browser --clean-playwright` empties it on start. Wait for the `listening` line in the job output. A second start takes over from a running one (it stops the old loop and server), and stopping the job ends its server too.
- After changing sources: run the generate step the change needs (`go generate ./internal/web` for templates and classes), then `touch data-test/restart`. The script rebuilds and restarts; wait for the next `listening` line. A failed build prints `run.sh: build failed`; fix it and touch `restart` again.
- Stop it when done: `touch data-test/stop`. Fallback if a server outlives its loop anyway: `taskkill //IM gm-session-manager-test.exe //F`. Never kill `gm-session-manager.exe`: that is the user's real instance.
- `.env` supplies the admin credentials. For the SRD, copy `data/compendium` into `data-test/`.

## Roles

- GM: `https://localhost:8443` (loopback is always the GM).
- Player: `https://<host LAN IP>:8443` in another tab (the IP is in the startup log line `lan=[...]`). The two hosts keep separate cookies, so both roles work in one browser. You start with no player session; join one yourself:
  1. Open `https://<LAN IP>:8443/join/1`, fill `#player-name`.
  2. An existing character: click the button in `form[action="/join/1/claim/<id>"]`, then approve the device from the GM tab on `/campaigns/1`. A new character: "New character", then Create (no approval).
- Commands and level-ups can be posted from the player tab with `fetch` and header `HX-Request: true`: `/characters/<id>/cmd`, `/characters/<id>/levelup`, `.../levelup/revert`. Select the level-up form by `input[name=class]`, not `form[method=post]` (the feedback form comes first).
- `sqlite3 -readonly data-test/gm-session-manager.db` reads sheets directly.

## What to check

- The changed pages only. Work the controls through `browser_snapshot`. Take at most one screenshot per page per width (390px, 1280px), after the controls are tried, not after every step. Read the console once per page (CSP errors included).
- At 390px, check `document.documentElement.scrollWidth` is 390: anything wider is a horizontal page scroll, a failure.
- After a restart, wait for the page to load before clicking: a click that lands too early is lost without any error.
- The MCP takes over `confirm()` dialogs: click the control, then answer with `browser_handle_dialog` (dismiss destructive ones unless told otherwise).
- Report back in 15 lines or fewer: what you checked, what was wrong, console errors, and the screenshot names.
- **Sample, don't crawl.** When exploring, open one page of each kind, not every entry. No scripted crawls. `browser_run_code_unsafe` only to reach another context's page (`page.context().browser().contexts()`).
- Read screenshots yourself and report what you saw. Save them with a `filename` under `data-test/playwright/` (the MCP output dir), never elsewhere in the repo. Give them descriptive names (`<page>-<width>.png`) and cite them by name in a handoff.
- Test data in `data-test` is shared between sessions: don't delete campaigns or characters you did not create.
