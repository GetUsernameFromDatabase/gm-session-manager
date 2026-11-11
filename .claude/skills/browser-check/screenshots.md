# Before/after pairs

For `pair Cnn: <element> on <page>, as GM|player`, one pair or several separated by `;` (a batch). "Before" is the snapshot `./run.sh before` took before the edits, on port 8444; "after" is the working tree on 8443. `<run>` is `<date>-<first ID>-<last ID>` (`2026-11-11-C05-C07`), the run folder `data-test/screenshots/<run>/`.

## Rules

- One change per pair: `data-test/screenshots/<run>/Cnn-before.png` and `Cnn-after.png` (MCP `filename: ../screenshots/<run>/Cnn-before.png`), never several changes in one image. "After" is the whole working tree: when it holds other items too, keep their changes out of the crop or name them in the report.
- The same element, crop, state and viewport in both. Screenshot the element (`target`), not the full page. Add a 390px pair only when the change is about phones.
- Changes you can't see (aria names): a before/after accessibility snapshot `.yml` of just that element instead.
- Element shots taller than the viewport get the sticky header over their top: use a taller viewport or a `clip` below the header.
- In `data-test/screenshots/<run>/README.md`, per pair: a `## Cnn <item>` heading, one line on what changed and the files, then a table `| before | after |` with `![](Cnn-before.png)` and `![](Cnn-after.png)`. Add `- [<run> <topic>](<run>/README.md)` at the top of `data-test/screenshots/README.md` (newest first). Change IDs never go in commit messages.

## The "before" instance (once per batch)

1. If `build/gm-session-manager-before.exe` is missing (nobody ran `./run.sh before` before editing), run `./run.sh before --head`: it builds HEAD and copies the current test DB. Say so in the report: "before" is then HEAD, with data the 8443 instance may already have upgraded.
2. From the repo root (the binary reads `.env` there), start `./run.sh browser-before` as a background job and wait for its `listening` line (port 8444). It serves a fresh copy of the snapshot from `data-test/.before/data` on each start.
3. Then start the 8443 instance. Ports share the host's cookie: use one browser context per port, and log a player in there in its own context (from a saved state in `SKILL.md` when one fits; apply its scene to each instance separately, checking `ready` against that instance's DB).

## After

Once all pairs of the batch are taken: stop the 8444 instance (`touch data-test/.before/stop`), stop the test instance (`touch data-test/stop`), close the browser (`browser_close`). Report per pair the file names and what differs, in 15 lines or fewer per pair.
