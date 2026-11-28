#!/usr/bin/env bash
# Takes a backup, then pulls origin/main and deploys it as a Docker image; run by gsm-update.timer (README.md).
#   update.sh [--nightly] [--accept-migrations] [--accept-edited-migrations]
# --nightly (the timer's) takes the nightly backup; a run by hand takes a manual one.
# A new migration, or an edit to a deployed one, pauses the update until you accept it with the flag.
# Every outcome is one line in DATA_DIR/update-status, which the GM footer shows. The .env is the checkout's,
# or ENV_FILE.
# File names come from common.sh.
# -E: the ERR traps also fire for a command failing inside a function (compose, status).
set -Eeuo pipefail

# The whole script is a function called on the last line: it checks out new versions of itself, and bash
# reads a script as it runs.
main() {
	local accept_new=0 accept_edited=0 backup=manual here data image deployed back new short dl
	for arg in "$@"; do
		case $arg in
		--nightly) backup=nightly ;;
		--accept-migrations) accept_new=1 ;;
		--accept-edited-migrations) accept_edited=1 ;;
		*) echo "usage: update.sh [--nightly] [--accept-migrations] [--accept-edited-migrations]" >&2; exit 2 ;;
		esac
	done
	here=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

	die() { echo "update: $*" >&2; exit 1; }
	# shellcheck source=deploy/common.sh
	. "$here/common.sh"
	[ -d "$data" ] || die "$data is not a folder (create it, owned by GSM_UID:GSM_GID, 65532 by default)"
	cd "$REPO"
	lock

	compose() { docker compose -f "$COMPOSE_FILE" --env-file "$ENV_FILE" "$@"; }
	image_exists() { docker image inspect "$image:$1" >/dev/null 2>&1; }
	# Atomic, so the app never reads half a line.
	status() {
		local line
		line="$(date '+%F %H:%M') $*"
		printf '%s\n' "$line" >"$data/.$STATUS_FILE.tmp"
		chmod 644 "$data/.$STATUS_FILE.tmp"
		mv -f "$data/.$STATUS_FILE.tmp" "$data/$STATUS_FILE"
		echo "$line"
	}
	warn() { printf '%s\n' "$@" >&2; }
	# The container is running and the app has logged its listening line.
	healthy() {
		local cid i
		# Migrations run before that line: a short wait would roll back, and so stop, a slow one half done.
		for i in $(seq 180); do
			cid=$(compose ps -q app 2>/dev/null || true)
			# Not a pipe: grep -q quits at the match, and under pipefail the cut-off docker logs fails the test.
			if [ -n "$cid" ] && [ "$(docker inspect -f '{{.State.Running}}' "$cid")" = true ] &&
				grep -q listening <(docker logs "$cid" 2>&1); then
				return 0
			fi
			sleep 1
		done
		return 1
	}
	normalized() { git show "$1:$2" | perl "$NORMALIZER"; }
	# Any other failing command (set -e) still reaches the footer, e.g. a checkout refused over local edits.
	unexpected() { status "stopped: unexpected error, see journalctl -u gsm-update"; }
	trap unexpected ERR

	# compose.yaml names the image it runs (<name>:current); the other tags share the name.
	# shortcut: read before the checkout, so a commit renaming the image fails to start and rolls back; retag by hand then.
	image=$(compose config --images)
	image=${image%:*}

	deployed=$(cat "$data/$DEPLOYED_FILE" 2>/dev/null || true)
	back=${deployed:-$(git rev-parse HEAD)}
	if [ -n "$deployed" ] && [ ! -f "$data/$DB_FILE" ]; then
		status "stopped: $data/$DB_FILE is missing; not updating (restore it from $data/$BACKUPS_DIR)"
		exit 1
	fi
	# Every run: the timer's nightly skips itself when nothing changed; a run by hand always takes a manual one.
	if [ -f "$data/$DB_FILE" ] && ! bash "$BACKUP_SCRIPT" "$backup"; then
		status "stopped: the $backup backup failed, not updating (journalctl -u gsm-update)"
		exit 1
	fi

	if ! git fetch -q origin main; then
		status "stopped: git fetch failed (offline?)"
		exit 1
	fi
	new=$(git rev-parse origin/main)
	short=$(git rev-parse --short "$new")
	if [ "$new" = "$deployed" ] && image_exists current; then
		status "up to date at $short"
		exit 0
	fi
	# A commit that did not start is not retried every night; a new commit on main is.
	if [ "$new" = "$(cat "$data/$FAILED_FILE" 2>/dev/null || true)" ]; then
		status "skipped: $short did not start before, waiting for a new commit (rm $data/$FAILED_FILE to retry)"
		exit 1
	fi

	if [ -z "$deployed" ]; then
		if [ "$accept_new" = 0 ]; then
			status "paused: first install of $short, run $DEPLOY/update.sh --accept-migrations"
			warn "WARNING: first install of $short: the database is created or migrated on the first start." \
				"Move the data in first if it comes from the PC ($README), then run: $DEPLOY/update.sh --accept-migrations"
			exit 1
		fi
	else
		migration_gate
	fi

	git checkout -q --detach "$new"
	echo "update: building $short"
	if ! build_and_check; then
		git checkout -q --detach "$back"
		status "build failed for $short; the running version is untouched (see the journal)"
		exit 1
	fi

	compose stop
	# A failing command (set -e) until the new image is up would leave the app down until the next night.
	trap 'git checkout -q --detach "$back"; compose up -d || true; status "failed while installing $short${deployed:+, restarted ${deployed:0:7}} (journalctl -u gsm-update)"' ERR
	if [ -f "$data/$DB_FILE" ] && ! bash "$BACKUP_SCRIPT" pre-deploy; then
		git checkout -q --detach "$back"
		image_exists current && compose up -d
		status "stopped: backup before the update failed, still on ${deployed:0:7}"
		exit 1
	fi
	# The app serves $dl (the app for taking a character home). Copied here, where a failure restarts the old
	# version; moved in only once the new one is up, as a rolled-back version would refuse the newer app's files.
	dl=$data/$DOWNLOADS_DIR
	rm -rf "$dl.new"
	mkdir "$dl.new"
	cp "$STAGING"/* "$dl.new/"
	chmod -R a+rX "$dl.new"
	image_exists current && docker tag "$image:current" "$image:prev"
	docker tag "$image:new" "$image:current"
	compose up -d || true # a failed start is handled below
	trap unexpected ERR

	if healthy; then
		rm -rf "$dl"
		mv "$dl.new" "$dl"
		printf '%s\n' "$new" >"$data/$DEPLOYED_FILE.tmp"
		mv -f "$data/$DEPLOYED_FILE.tmp" "$data/$DEPLOYED_FILE"
		rm -f "$data/$FAILED_FILE"
		docker image prune -f >/dev/null || true
		status "deployed $short"
		exit 0
	fi

	docker logs --tail 20 "$(compose ps -aq app | head -n 1)" >&2 || true
	git checkout -q --detach "$back"
	rm -rf "$dl.new"
	if [ -z "$deployed" ]; then
		compose stop
		status "failed: the first start of $short did not come up (docker compose logs)"
		exit 1
	fi
	printf '%s\n' "$new" >"$data/$FAILED_FILE"
	docker tag "$image:prev" "$image:current"
	compose up -d || true
	if healthy; then
		status "rolled back to ${deployed:0:7}: $short did not start (docker compose logs)"
	else
		status "rollback failed: $short and ${deployed:0:7} both refuse to start; restore the pre-deploy backup ($README#rollback)"
		warn "A migration of $short probably ran already, and the older version refuses the newer database."
	fi
	exit 1
}

# Stops the update (exit 1) when migrations changed and the matching flag is missing. Uses main's locals.
migration_gate() {
	local changes st f g added=() described=() edited=() list
	changes=$(git diff --name-status --no-renames "$deployed" "$new" -- "$MIGRATIONS/") || {
		status "stopped: cannot compare $deployed with $short (history rewritten?)"
		exit 1
	}
	while IFS=$'\t' read -r st f g; do
		[ -n "$st" ] || continue
		case $st in
		A)
			added+=("${f##*/}")
			# The first line is the migration's describing comment. sed, not head: an early exit fails the pipe.
			described+=("${f##*/}: $(git show "$new:$f" | sed -n 1p)")
			;;
		M)
			if [ "$(normalized "$deployed" "$f")" = "$(normalized "$new" "$f")" ]; then
				echo "${f##*/}: comment-only change, nothing to run"
			else
				edited+=("${f##*/} (modified)")
			fi
			;;
		*) edited+=("${f##*/} (deleted or renamed)") ;;
		esac
	done <<<"$changes"

	if [ "${#edited[@]}" -gt 0 ] && [ "$accept_edited" = 0 ]; then
		status "stopped: a deployed migration was changed (${edited[*]}), see the journal; fix it or run $DEPLOY/update.sh --accept-edited-migrations"
		list=$(printf '  %s\n' "${edited[@]}")
		warn "WARNING: update stopped at ${deployed:0:7} -> $short: a deployed migration was changed" "$list" \
			"The server's database already ran this migration and never runs it again, so the" \
			"change would be skipped silently and the code and database would disagree." \
			"Either: revert the edit in the repo and move the change into a new migration (preferred)," \
			"or:     apply the same change to the server's database by hand (sqlite3 \"$data/$DB_FILE\")," \
			"        then run: $DEPLOY/update.sh --accept-edited-migrations"
		exit 1
	fi
	if [ "${#added[@]}" -gt 0 ] && [ "$accept_new" = 0 ]; then
		status "paused at $short: new migrations (${added[*]}), run $DEPLOY/update.sh --accept-migrations"
		list=$(printf '  %s\n' "${described[@]}")
		warn "WARNING: update paused at ${deployed:0:7} -> $short: new migrations" "$list" \
			"A migration changes the database as soon as the new version starts, and the older" \
			"version refuses to open it afterwards, so migrations are not applied unattended." \
			"Before accepting: read the migration; if it rebuilds tables, rehearse it on a copy" \
			"($README#rehearse). A backup is taken right before it runs." \
			"To apply: $DEPLOY/update.sh --accept-migrations" \
			"Rollback: $README#rollback"
		exit 1
	fi
}

# Builds the image and the release archives of the checked-out commit, and checks the Compose file.
build_and_check() {
	local version
	version=$(git describe --tags --always)
	rm -rf "$STAGING"
	# --pull: the Go image runs GOTOOLCHAIN=local, so a stale cached one refuses a newer go.mod.
	docker build --pull -f "$DOCKERFILE" --build-arg VERSION="$version" --target archives --output type=local,dest="$STAGING" . &&
		docker build --pull -f "$DOCKERFILE" --build-arg VERSION="$version" -t "$image:new" . &&
		compose config -q
}

main "$@"
exit # never read past here: the file changes under a running script when a checkout replaces it
