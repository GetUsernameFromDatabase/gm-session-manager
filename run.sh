#!/usr/bin/env bash
# Usage: ./run.sh          the real instance (.env and its DATA_DIR); rebuilds and restarts when sources
#                          change, keeping the old server if the build fails. A changed migration pauses
#                          reloading. touch build/restart to force a restart (and resume), build/stop to quit.
#        ./run.sh browser [--clean-playwright]
#                          a test instance for browser checks (data-test/, port 8443); restarts only on
#                          touch data-test/restart, quits on data-test/stop.
#                          data-test/playwright is kept unless --clean-playwright is given.
#        ./run.sh before [--head]
#                          snapshot for "before" screenshots, taken before editing: builds the working tree
#                          (--head: HEAD) and copies the data-test DB to data-test/.before/pristine.db.
#        ./run.sh browser-before
#                          serves that snapshot on port 8444 from a fresh copy in data-test/.before/data;
#                          touch data-test/.before/stop to quit.
# A second start of a mode takes over from the first; each quits with its terminal.
cd "$(dirname "$0")"
pkg=./cmd/gm-session-manager
exe=$(go env GOEXE)
# Unique binary names, so a stray one can be stopped by name without touching any other instance.
before_bin=build/gm-session-manager-before$exe
mkdir -p build

case $1 in
'')
	bin=build/gm-session-manager-dev$exe markers=build watch=1
	;;
browser)
	export DATA_DIR=data-test PORT=8443
	bin=build/gm-session-manager-test$exe markers=data-test
	[ "$2" = --clean-playwright ] && rm -rf "$DATA_DIR/playwright"
	;;
before)
	mkdir -p data-test/.before
	if [ "$2" = --head ]; then
		wt=build/head
		git worktree remove --force "$wt" 2>/dev/null
		git worktree add --detach "$wt" HEAD || exit 1
		trap 'git worktree remove --force "$wt"' EXIT
		# The worktree lacks the gitignored browser assets.
		cp -r internal/web/assets/vendor "$wt/internal/web/assets/" &&
			cp internal/web/assets/js/shadcn-templ-*.js "$wt/internal/web/assets/js/" &&
			(cd "$wt" && go build -o "../../$before_bin" "$pkg") || exit 1
	else
		go build -o "$before_bin" "$pkg" || exit 1
	fi
	rm -f data-test/.before/pristine.db
	if [ -f data-test/gm-session-manager.db ]; then
		sqlite3 data-test/gm-session-manager.db ".backup 'data-test/.before/pristine.db'" || exit 1
	fi
	echo "run.sh: before snapshot ready; serve it with ./run.sh browser-before"
	exit 0
	;;
browser-before)
	[ -f "$before_bin" ] || { echo "run.sh: no $before_bin; run ./run.sh before first"; exit 1; }
	export DATA_DIR=data-test/.before/data PORT=8444
	bin=$before_bin markers=data-test/.before norebuild=1
	;;
*)
	sed -n '/^# /p;/^cd /q' "$0"
	exit 1
	;;
esac

rm -f "$markers/restart" "$markers/stop"
pid=
stop_server() {
	# On Windows, kill does not always end a native process; a stray from a killed loop goes too.
	if [ -n "$WINDIR" ]; then
		taskkill //F //IM "${bin##*/}" >/dev/null 2>&1
	elif [ -n "$pid" ]; then
		kill "$pid" 2>/dev/null
	fi
	[ -n "$pid" ] && wait "$pid" 2>/dev/null
	pid=
}
# Build first and swap only on success, so a broken build leaves the old server running.
restart_server() {
	new=${bin%"$exe"}-new$exe
	if [ -z "$norebuild" ] && ! go build -o "$new" "$pkg"; then
		echo "run.sh: build failed${pid:+; the old server keeps running}. Fix it (or touch $markers/restart)"
		return
	fi
	stop_server
	[ -z "$norebuild" ] && mv -f "$new" "$bin"
	"$bin" &
	pid=$!
}
# One loop at a time: a newer run takes over (an older one waits for every marker).
pidfile=$markers/run.pid
old=$(cat "$pidfile" 2>/dev/null)
echo $$ >"$pidfile"
for _ in 1 2 3 4 5; do
	[ -n "$old" ] && kill -0 "$old" 2>/dev/null || break
	sleep 1
done
stop_server
trap stop_server EXIT

if [ "$1" = browser-before ]; then
	# A fresh copy each start: a stale -wal/-shm next to an old DB corrupts it.
	rm -rf "$DATA_DIR" && mkdir -p "$DATA_DIR"
	for d in compendium tls attachments; do
		[ -d "data-test/$d" ] && cp -r "data-test/$d" "$DATA_DIR/"
	done
	[ -f data-test/.before/pristine.db ] && cp data-test/.before/pristine.db "$DATA_DIR/gm-session-manager.db"
fi

stamp=$markers/.stamp
[ -n "$watch" ] && touch "$stamp"
# Sources changed since the last build. Generated files are left out: the watcher regenerates them
# whenever their sources change, and its own output must not count as a change.
changes() {
	find cmd internal -type f -newer "$stamp" \( -name '*.go' -o -name '*.templ' -o -name '*.sql' -o -name '*.css' -o -name '*.js' \) \
		-not -path 'internal/web/assets/vendor/*' -not -name '*_templ.go' -not -name output.css
}

# Windows ends the parent (terminal closed, job stopped) without a signal here: watch it.
parent=$PPID
[ "$parent" = 1 ] && parent=
start=1 paused=
while :; do
	if [ -n "$parent" ] && ! kill -0 "$parent" 2>/dev/null || [ "$(cat "$pidfile" 2>/dev/null)" != $$ ]; then
		exit 0
	fi
	if [ -e "$markers/stop" ]; then
		rm -f "$markers/stop"
		exit 0
	fi
	if [ -e "$markers/restart" ]; then
		rm -f "$markers/restart"
		start=1
	fi
	changed=
	[ -n "$watch" ] && changed=$(changes)
	if [ -n "$changed" ] && [ -z "$paused" ] && [ -z "$start" ]; then
		# A migration from half-finished work would change the real DB: wait for the user.
		if grep -q '^internal/db/migrations/' <<<"$changed"; then
			paused=1
			echo "run.sh: migrations changed; not restarting. touch $markers/restart when ready"
		else
			start=1
		fi
	fi
	if [ -n "$start" ]; then
		start= paused=
		if [ -n "$watch" ]; then
			# Before generating and building, so an edit saved meanwhile counts as the next change.
			touch "$stamp"
			# Not templ fmt: the watcher must never rewrite a file someone is editing.
			grep -qE '\.templ$|globals\.css$' <<<"$changed" && go generate -run 'templ generate|tailwindcss' ./internal/web
		fi
		restart_server
	fi
	sleep 2
done
