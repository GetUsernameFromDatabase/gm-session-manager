#!/usr/bin/env bash
# task lint [-- --all]: golangci-lint, gopls check on Go files changed against HEAD (--all: every Go
# file, e.g. after a gopls upgrade), then stale Tailwind classes (tailwindcss canonicalize).
# Task runs it in its own shell (source), so it also works from PowerShell: no find, sort or bash-only set flags.
# Reports every problem before failing.
set -uo pipefail # a tool failing mid-pipe (tailwindcss missing) must fail the lint, not read as "nothing found"
cd "$(git rev-parse --show-toplevel)"

status=0
golangci-lint run || status=1

# gopls has analyzers golangci-lint cannot run (writestring). Each run starts cold (~1.5s,
# ~19s for every package), so by default only changed files. Skips what .golangci.yml skips.
X=(':!*_templ.go' ':!internal/db/sqlc' ':!internal/web/components' ':!internal/web/utils')
if [ "${1:-}" = --all ]; then
	gofiles=$(git ls-files -co --exclude-standard -- '*.go' "${X[@]}")
	gofiles=$(for f in $gofiles; do [ -f "$f" ] && echo "$f"; done) # a deleted, uncommitted file is still listed
else
	gofiles=$(git diff --name-only --diff-filter=d HEAD -- '*.go' "${X[@]}"; git ls-files -o --exclude-standard -- '*.go' "${X[@]}")
fi
if [ -n "$gofiles" ]; then
	found=$(gopls check $gofiles 2>&1) || { echo "lint: gopls check did not run" >&2; status=1; }
	if [ -n "$found" ]; then
		echo "$found"
		status=1
	fi
fi

# Every quoted word in hand-written web sources is a class candidate; tailwindcss returns
# non-classes unchanged. components/, utils/ and the script bundle belong to the shadcn-templ CLI.
Y=(':!internal/web/components' ':!internal/web/utils' ':!internal/web/assets/vendor' ':!*_templ.go' ':!*_test.go' ':!*shadcn-templ-*.js')
# '…' only in JS: in templ and Go text, two apostrophes on a line would swallow a class string.
stale=$({ git grep --untracked -ohE "\"[^\"]*\"|\`[^\`]*\`" -- 'internal/web/*.templ' 'internal/web/*.go' 'internal/web/*.js' "${Y[@]}" &&
	git grep --untracked -ohE "'[^']*'" -- 'internal/web/*.js' "${Y[@]}"; } |
	tr -d "\"'\`" | tr -s ' \t' '\n' |
	tailwindcss canonicalize --stream --css internal/web/assets/css/globals.css --format jsonl |
	{ grep -F '"changed":true' || [ $? = 1 ]; }) || # any key order; grep's 1 is "none found"
	{ echo "lint: the Tailwind class check did not run" >&2; status=1; }
if [ -n "$stale" ]; then
	echo "rename these Tailwind classes (input -> output):"
	echo "$stale"
	status=1
fi
exit $status
