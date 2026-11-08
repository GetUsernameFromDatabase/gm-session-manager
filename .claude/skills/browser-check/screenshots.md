# Before/after pairs

For `pair Cnn: <element> on <page>, as GM|player`. "Before" is the last commit (HEAD) on port 8444; "after" is the working tree on 8443. `<scratch>` is the scratchpad directory from your environment.

## Rules

- One change per pair: `data-test/changes/Cnn-before.png` and `Cnn-after.png`, never several changes in one image.
- The same element, crop, state and viewport in both. Screenshot the element (`target`), not the full page. Add a 390px pair only when the change is about phones.
- Changes you can't see (aria names): a before/after accessibility snapshot `.yml` of just that element instead.
- Element shots taller than the viewport get the sticky header over their top: use a taller viewport or a `clip` below the header.
- Add a row to `data-test/changes/CHANGES.md` (ID, item, files, what changed, screenshots). Change IDs never go in commit messages.

## The "before" instance (once per batch)

1. Copy `data-test/compendium`, `tls` and `attachments` into `<scratch>/data-old`, then `sqlite3 data-test/gm-session-manager.db ".backup '<scratch>/data-old/gm-session-manager.db'"`. If the working tree adds a migration, also run `pragma user_version=<old>; update compendium_sources set hash='';` on the copy.
2. `git worktree add --detach <scratch>/head HEAD`, copy `internal/web/assets/vendor` and `internal/web/assets/js/shadcn-templ-*.js` into it, then in it `go build -o <scratch>/gm-old.exe ./cmd/gm-session-manager`.
3. From the repo root, with `.env` loaded, run it in the background with `DATA_DIR=<scratch>/data-old PORT=8444`.
4. Ports share the host's cookie: use one browser context per port, and join a player there in its own context.

## After

Stop the 8444 job, `git worktree remove --force <scratch>/head`, stop the test instance (`touch data-test/stop`). Report the file names and what differs, in 15 lines or fewer.
