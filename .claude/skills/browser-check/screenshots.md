# Before/after pairs

For `pair Cnn: <element> on <page>, as GM|player`, one pair or several separated by `;` (a batch). "Before" is the last commit (HEAD) on port 8444; "after" is the working tree on 8443. `<scratch>` is the scratchpad directory from your environment; `<run>` is the first pair's ID (`C05`).

## Rules

- One change per pair: `data-test/changes/Cnn-before.png` and `Cnn-after.png`, never several changes in one image. "After" is the whole working tree: when it holds other items too, keep their changes out of the crop or name them in the report.
- The same element, crop, state and viewport in both. Screenshot the element (`target`), not the full page. Add a 390px pair only when the change is about phones.
- Changes you can't see (aria names): a before/after accessibility snapshot `.yml` of just that element instead.
- Element shots taller than the viewport get the sticky header over their top: use a taller viewport or a `clip` below the header.
- Add a row to `data-test/changes/CHANGES.md` (ID, item, files, what changed, screenshots). Change IDs never go in commit messages.

## The "before" instance (once per batch)

Do this before you start or restart the 8443 instance: it upgrades sheets on load, so a copy taken after it shows already-fixed data.

1. Pristine database: if `<scratch>/pristine/gm-session-manager.db` is missing, `sqlite3 data-test/gm-session-manager.db ".backup '<scratch>/pristine/gm-session-manager.db'"`. Use the existing one otherwise (the caller may have taken it earlier).
2. A new folder `<scratch>/data-old-<run>` (never reuse an old one: stale `-wal`/`-shm` files corrupt the copy). Copy `data-test/compendium`, `tls` and `attachments` into it, and copy only the pristine `.db` file (no `-wal`/`-shm`). If the working tree adds a migration, also run `pragma user_version=<old>; update compendium_sources set hash='';` on the copy.
3. `git worktree add --detach <scratch>/head HEAD`, copy `internal/web/assets/vendor` and `internal/web/assets/js/shadcn-templ-*.js` into it, then in it `go build -o <scratch>/gm-old.exe ./cmd/gm-session-manager`.
4. From the repo root (the binary reads `.env` there), run it in the background with `DATA_DIR=<scratch>/data-old-<run> PORT=8444`.
5. Then start the 8443 instance. Ports share the host's cookie: use one browser context per port, and log a player in there in its own context (from a saved state in `SKILL.md` when one fits; apply its scene to each instance separately, checking `ready` against that instance's DB).

## After

Once all pairs of the batch are taken: stop the 8444 job, `git worktree remove --force <scratch>/head`, stop the test instance (`touch data-test/stop`), close the browser (`browser_close`). Report per pair the file names and what differs, in 15 lines or fewer per pair.
