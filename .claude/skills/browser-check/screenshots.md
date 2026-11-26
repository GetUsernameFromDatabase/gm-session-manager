# Before/after pairs

For `pair <name>: <element> on <page>, as GM|player`, one pair or several separated by `;` (a batch). "Before" is the snapshot `./run.sh before` took before the edits, on port 8444; "after" is the working tree on 8443. The run folder is `data/test/screenshots/<date>-<topic>/`.

## Rules

- One change per pair: `<name>-before.png` and `<name>-after.png` in the run folder, never several changes in one image. "After" is the whole working tree: when it holds other changes too, keep them out of the crop or name them in the report.
- The same element, crop, state and viewport in both. Screenshot the element (`locator.screenshot`), not the full page. Add a 390px pair only when the change is about phones.
- Changes you can't see (aria names): a before/after `ariaSnapshot()` of just that element, saved as `.yml`, instead.
- Element shots taller than the viewport get the sticky header over their top: use a taller viewport or a `clip` below the header.
- Don't Read the PNGs: the caller views them. No README.

## The "before" instance (once per batch)

1. If `build/gm-session-manager-before.exe` is missing (nobody ran `./run.sh before` before editing), run `./run.sh before --head`: it builds HEAD and copies the current test DB. Say so in the report: "before" is then HEAD, with data the 8443 instance may already have upgraded.
2. From the repo root (the binary reads `.env` there), start `./run.sh browser-before` as a background job and wait for its `listening` line (port 8444). It serves a fresh copy of the snapshot from `data/test/.before/data` on each start.
3. Then start the 8443 instance. A saved state's scene (`SKILL.md`) is applied to each instance separately, checking `ready` against that instance's DB.

## Taking a pair: one script

Both ports share the host's cookies, so each port gets its own browser context. One `browser_run_code_unsafe` takes the whole pair (or the whole batch):

```js
async (page) => {
  const browser = page.context().browser(), out = {};
  for (const [port, side] of [[8444, 'before'], [8443, 'after']]) {
    const ctx = await browser.newContext({ ignoreHTTPSErrors: true, viewport: { width: 1280, height: 900 } });
    // player: await ctx.addCookies(<state cookies>) and use the LAN host instead of localhost
    const p = await ctx.newPage();
    await p.goto(`https://localhost:${port}/campaigns/1`);
    // the same steps on both sides
    await p.locator('<element>').screenshot({ path: `data/test/screenshots/<run>/<name>-${side}.png` });
    out[side] = await p.locator('<element>').innerText();
    await ctx.close();
  }
  return out;
}
```

Paths in a script are relative to the repo root (the MCP's working folder); never an absolute path. Fix a failing script and run it again; step through by hand only to find out why it fails. If its steps change data, check `ready` on each instance (or restart 8444 for a fresh copy) before running it again.

## After

Once all pairs of the batch are taken: stop the 8444 instance (`touch data/test/.before/stop`), stop the test instance (`touch data/test/stop`), close the browser (`browser_close`). Report per pair the file paths and what the script returned, in 5 lines or fewer per pair.
