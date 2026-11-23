#!/usr/bin/env bash
# Usage: ./lint.sh         golangci-lint, gopls check on Go files changed against HEAD,
#                          then stale Tailwind classes (tailwindcss canonicalize).
#        ./lint.sh --all   gopls check on every Go file, e.g. after a gopls upgrade.
# Run by .githooks/pre-commit and host builds (build.sh). Reports every problem before failing.
set -uo pipefail # a tool failing mid-pipe (tailwindcss missing) must fail the lint, not read as "nothing found"
cd "$(dirname "$0")"

status=0
golangci-lint run || status=1

# gopls has analyzers golangci-lint cannot run (writestring). Each run starts cold (~1.5s,
# ~19s for every package), so by default only changed files. Skips what .golangci.yml skips.
if [ "${1:-}" = --all ]; then
	gofiles=$(git ls-files -co --exclude-standard -- '*.go')
else
	gofiles=$(git diff --name-only HEAD -- '*.go'; git ls-files -o --exclude-standard -- '*.go')
fi
gofiles=$(grep -vE '_templ\.go$|^internal/(db/sqlc|web/components|web/utils)/' <<<"$gofiles" |
	while read -r f; do [ -f "$f" ] && echo "$f"; done) # a deleted file is still listed as changed
if [ -n "$gofiles" ]; then
	found=$(xargs gopls check <<<"$gofiles" 2>&1) || { echo "lint.sh: gopls check did not run" >&2; status=1; }
	if [ -n "$found" ]; then
		echo "$found"
		status=1
	fi
fi

# Every quoted word in hand-written web sources is a class candidate; tailwindcss returns
# non-classes unchanged. components/, utils/ and the script bundle belong to the shadcn-templ CLI.
web=internal/web
sources() { # find args naming the file types
	find "$web" \( -path "$web/components" -o -path "$web/utils" -o -path "$web/assets/vendor" \) -prune -o \
		-type f "$@" ! -name '*_templ.go' ! -name '*_test.go' ! -name 'shadcn-templ-*.js' -print0
}
# '…' only in JS: in templ and Go text, two apostrophes on a line would swallow a class string.
stale=$({ sources \( -name '*.templ' -o -name '*.go' -o -name '*.js' \) | xargs -0 grep -ohE "\"[^\"]*\"|\`[^\`]*\`" &&
	sources -name '*.js' | xargs -0 grep -ohE "'[^']*'"; } | tr -d "\"'\`" | tr -s ' \t' '\n' | sort -u |
	tailwindcss canonicalize --stream --css "$web/assets/css/globals.css" --format jsonl |
	{ grep -F '"changed":true' || [ $? = 1 ]; }) || # any key order; grep's 1 is "none found"
	{ echo "lint.sh: the Tailwind class check did not run" >&2; status=1; }
if [ -n "$stale" ]; then
	echo "rename these Tailwind classes (input -> output):"
	echo "$stale"
	status=1
fi
exit $status
