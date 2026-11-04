#!/usr/bin/env bash
# Usage: ./run.sh          the real instance (.env, data/)
#        ./run.sh browser [--clean-playwright]
#                          a test instance for browser checks (data-test/, port 8443);
#                          touch data-test/restart to rebuild and restart it, data-test/stop to quit.
#                          A second start takes over from the first; it quits with its terminal.
#                          data-test/playwright is kept unless --clean-playwright is given.
cd "$(dirname "$0")"
if [ "$1" = browser ]; then
	export DATA_DIR=data-test PORT=8443
	# Its own binary name, so a stray one can be stopped by name without touching the real instance.
	bin=build/gm-session-manager-test$(go env GOEXE)
	mkdir -p "$DATA_DIR"
	rm -f "$DATA_DIR/restart" "$DATA_DIR/stop"
	[ "$2" = --clean-playwright ] && rm -rf "$DATA_DIR/playwright"
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
	# One loop at a time: a newer run takes over (an older one waits for every marker).
	pidfile=$DATA_DIR/run.pid
	old=$(cat "$pidfile" 2>/dev/null)
	echo $$ >"$pidfile"
	for _ in 1 2 3 4 5; do
		[ -n "$old" ] && kill -0 "$old" 2>/dev/null || break
		sleep 1
	done
	stop_server
	trap stop_server EXIT
	# Windows ends the parent (terminal closed, job stopped) without a signal here: watch it.
	parent=$PPID
	[ "$parent" = 1 ] && parent=
	start=1
	while :; do
		if [ -n "$parent" ] && ! kill -0 "$parent" 2>/dev/null || [ "$(cat "$pidfile" 2>/dev/null)" != $$ ]; then
			exit 0
		fi
		if [ -e "$DATA_DIR/stop" ]; then
			rm -f "$DATA_DIR/stop"
			exit 0
		fi
		if [ -e "$DATA_DIR/restart" ]; then
			rm -f "$DATA_DIR/restart"
			start=1
		fi
		if [ -n "$start" ]; then
			start=
			stop_server
			if go build -o "$bin" ./cmd/gm-session-manager; then
				"$bin" &
				pid=$!
			else
				echo "run.sh: build failed; fix it and touch $DATA_DIR/restart"
			fi
		fi
		sleep 1
	done
fi
exec go run ./cmd/gm-session-manager
