# Sourced by update.sh and backup.sh (not run): every file and folder name they use, the .env reading and the lock.
# Rename a file here and both scripts follow. The image name is compose.yaml's (update.sh reads it from there).
# Needs here (the deploy/ folder) and die. Sets ENV_FILE and DATA_DIR (absolute, exported for compose.yaml) and data
# (the same folder; it may not exist yet).

# Shared with the Go app: TestNamesMatchGo (names_test.go) fails when these and the Go constants differ.
DB_FILE=gm-session-manager.db
STATUS_FILE=update-status
DOWNLOADS_DIR=downloads

# In DATA_DIR, the scripts' own.
BACKUPS_DIR=backups
DEPLOYED_FILE=deployed-commit
FAILED_FILE=failed-commit

# In the repo. DEPLOY is this folder's name, for the commands that messages print.
REPO=$(dirname "$here")
DEPLOY=${here##*/}
COMPOSE_FILE=$here/compose.yaml
DOCKERFILE=$here/Dockerfile
NORMALIZER=$here/sql-normalize.pl
BACKUP_SCRIPT=$here/backup.sh
README=$DEPLOY/README.md
MIGRATIONS=internal/db/migrations
STAGING=$REPO/build/staging

# In the checkout, as on the PC.
ENV_FILE=${ENV_FILE:-$REPO/.env}
LOCK_FILE=${LOCK_FILE:-/run/lock/gsm-update.lock}

[ -f "$ENV_FILE" ] || die "$ENV_FILE not found (cp .env.example .env in the checkout)"
ENV_FILE=$(realpath "$ENV_FILE")
export ENV_FILE

# A relative DATA_DIR is read from the checkout, as the app on the PC reads it from the project folder. Compose gets
# the full path (the shell's value beats --env-file): it would read a relative one from deploy/.
data=$(grep -E '^DATA_DIR=' "$ENV_FILE" | tail -n 1 | cut -d= -f2- || true)
data=${data%[\"\']} data=${data#[\"\']}
[ -n "$data" ] || die "set DATA_DIR in $ENV_FILE"
case $data in /*) ;; *) data=$REPO/$data ;; esac
data=$(realpath -m "$data")
export DATA_DIR=$data

# One update or backup at a time; a second one waits. update.sh sets GSM_LOCKED for the backups it runs.
lock() {
	exec 9>"$LOCK_FILE"
	flock -w 1800 9 || die "another update or backup is still running"
	export GSM_LOCKED=1
}
