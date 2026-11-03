---
name: browser-check
description: Check UI changes in a real browser with the Playwright MCP server against a test instance of the app. Use after any change to .templ, CSS or JS, before saying it is done, and when asked to look at, test or explore pages in the browser.
---

# Browser check

## The test instance

- Start it once, as a background Bash job: `./run.sh browser`. It builds `build/gm-session-manager-test.exe`, serves `https://localhost:8443` from `data-test/` (never `data/`), and keeps `data-test/playwright` between runs, so screenshots a handoff refers to survive; `./run.sh browser --clean-playwright` empties it on start. Wait for the `listening` line in the job output.
- After changing sources: run the generate step the change needs (`go generate ./internal/web` for templates and classes), then `touch data-test/restart`. The script rebuilds and restarts; wait for the next `listening` line. A failed build prints `run.sh: build failed`; fix it and touch `restart` again.
- Stop it when done: `touch data-test/stop`. Fallback if the job was killed: `taskkill //IM gm-session-manager-test.exe //F`. Never kill `gm-session-manager.exe`: that is the user's real instance.
- `.env` supplies the admin credentials. For the SRD, copy `data/compendium` into `data-test/`.

## Roles

- GM: `https://localhost:8443` (loopback is always the GM).
- Player: `https://<host LAN IP>:8443` in another tab (the IP is in the startup log line `lan=[...]`). It goes through the real join flow; approve the device from the GM tab. The two hosts keep separate cookies, so both roles work in one browser.

## What to check

- The changed pages only, at 390px and 1280px wide. Try the new or changed controls and check the console for errors (CSP included).
- **Sample, don't crawl.** When exploring, open one page of each kind, not every entry. No scripted crawls and no `browser_run_code_unsafe` unless the user asks.
- Read screenshots yourself and report what you saw. Save them with a `filename` under `data-test/playwright/` (the MCP output dir), never elsewhere in the repo. Give them descriptive names (`<page>-<width>.png`) and cite them by name in a handoff.
- Test data in `data-test` is shared between sessions: don't delete campaigns or characters you did not create.
