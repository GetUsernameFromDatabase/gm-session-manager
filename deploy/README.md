# Running on a server (Debian, Docker)

The app runs in Docker on a Debian box on the LAN. Every night at 04:00 a timer takes a backup, then pulls `main`
and deploys it. New migrations wait for your OK. The GM footer shows the version and the last update.

## Folders

The commands below use this layout; with it, every block copies as is.

| Path | Holds |
| --- | --- |
| `/opt/gm-session-manager/src` | the git checkout, with its `.env` as on the PC |
| `/opt/gm-session-manager/data` | the data and its backups: `DATA_DIR=../data` in `.env` |

> The commands assume the `.env` in the checkout, with `DATA_DIR` unquoted. A relative `DATA_DIR` is read from the
> checkout, as on the PC; an absolute one works too. Keep the data outside the
> checkout: git deletes ignored files on `git clean -x`, and overwrites one if a commit ever adds the same path.
> Another checkout path: change it in the `git clone` and `cd` lines; the units take the checkout's path by themselves
> (step 5).

## Setup

As root, on the server, in one terminal: later steps use what earlier ones set. Needs Docker 23 or newer
(`docker version`).

1. Tools, code and settings:

   ```sh
   apt install git sqlite3
   git clone https://github.com/GetUsernameFromDatabase/gm-session-manager /opt/gm-session-manager/src
   cd /opt/gm-session-manager/src
   cp .env.example .env
   chmod 600 .env
   nano .env   # ADMIN_USERNAME, a strong ADMIN_PASSWORD, DATA_DIR=../data
   ```

   > Every GM device logs in with that admin login: on the server no browser is local. The file stays root's;
   > Docker reads it for the container.

2. Give the server a static DHCP lease in your router.

   > The self-signed certificate names the server's IP. A new IP means every device accepts a new certificate.

3. The data folder:

   ```sh
   DATA_DIR=$(realpath -m "$(grep '^DATA_DIR=' .env | cut -d= -f2-)") && echo "$DATA_DIR"
   install -d -m 750 -o 65532 -g 65532 "$DATA_DIR"
   ```

   > The first line makes `DATA_DIR` a full path, reading a relative one from the checkout as the scripts do. The
   > app runs as user 65532, the image's own non-root user, and owns only this folder. For an account of your own instead, see [Own account](#own-account).

4. Your data from the PC (skip for a fresh start). Stop `task run` on the PC, then in Git Bash there:

   ```sh
   server=root@<server>
   data=/opt/gm-session-manager/data   # the server's data folder
   tar -C data/dev --exclude=./tls -cf - . | ssh "$server" tar -C "$data" -xf -
   ```

   And on the server:

   ```sh
   chown -R 65532:65532 "$DATA_DIR"
   ```

   > `tls/` stays behind: the server makes a certificate for its own IP, and players accept it once.

5. First start, then the nightly timer:

   ```sh
   bash deploy/update.sh --accept-migrations
   bash deploy/systemd-units.sh install
   ```

   > The first run has no deployed version to compare migrations with, hence the flag. The container comes back by
   > itself after a reboot. `systemd-units.sh` writes the timer and its service from `deploy/systemd/` into
   > `/etc/systemd/system` with this checkout's path, and prints where they are.

### Own account

Optional, in the Setup terminal before step 5: run the app as a named service account (no password, no shell, no
home) instead of 65532.

```sh
useradd --system --user-group --no-create-home --shell /usr/sbin/nologin gm-session-manager
chown -R gm-session-manager:gm-session-manager "$DATA_DIR"
id gm-session-manager   # put its uid and gid in GSM_UID and GSM_GID in .env
```

## Day to day

```sh
journalctl -u gsm-update                                  # what the updates did; a stopped one says what to do
cd /opt/gm-session-manager/src && bash deploy/update.sh   # update now (takes a manual backup)
```

### The timer

The installed units are yours: an update never changes them. Change the update time with
`systemctl edit gsm-update.timer`, which keeps your change apart from the unit:

```ini
[Timer]
OnCalendar=
OnCalendar=*-*-* 03:30
```

Or edit the files in `/etc/systemd/system` and run `systemctl daemon-reload`. To take the repo's units after a
commit changed them (`install` shows how yours differ, and changes nothing while they exist):

```sh
cd /opt/gm-session-manager/src
bash deploy/systemd-units.sh install                                        # the differences
bash deploy/systemd-units.sh remove && bash deploy/systemd-units.sh install # the repo's; drop-ins of systemctl edit stay
```

## Backups

In `backups/` in the data folder: the newest 5 `nightly-*` (the timer's), 3 `manual-*` (by hand, and each update
run by hand) and 3 `pre-deploy-*` (taken right before each update). A night with no changes takes no backup, so the
5 nightlies go back 5 days of play, not 5 days.

```sh
cd /opt/gm-session-manager/src && bash deploy/backup.sh manual   # back up now
```

### Rehearse

Optional, before accepting a migration: run the new version on the PC against a copy of the server's data, and see
the migration go through on real data. The server is not touched. On the PC, with the new code checked out:

1. The server's newest backup, of any kind, into `build/`:

   ```sh
   server=root@<server>
   data=/opt/gm-session-manager/data   # the server's data folder
   backup=$(ssh "$server" "ls -t $data/backups/*.tar.gz | head -n 1") && echo "$backup"
   scp "$server:$backup" build/
   ```

2. Its database into `data/test`, then start:

   ```sh
   rm -f data/test/gm-session-manager.db*
   tar -xzf "build/${backup##*/}" -C data/test gm-session-manager.db
   task run:test
   ```

> The new version migrates the copy; check that the pages open. `data/test` is the browser-check instance, so its
> database is replaced. Never use `data/dev` for this.

### Rollback

Back to the state before the last update, as root.

1. Stop the timer, or tonight deploys the same commit again:

   ```sh
   systemctl stop gsm-update.timer
   ```

2. The checkout, and the full paths Compose needs:

   ```sh
   cd /opt/gm-session-manager/src
   export ENV_FILE=$PWD/.env
   export DATA_DIR=$(realpath -m "$(grep '^DATA_DIR=' .env | cut -d= -f2-)") && echo "$DATA_DIR"
   ```

3. The newest pre-deploy backup:

   ```sh
   backup=$(ls -t "$DATA_DIR"/backups/pre-deploy-*.tar.gz | head -n 1) && echo "$backup"
   ```

4. Stop the app and restore the data:

   ```sh
   docker compose -f deploy/compose.yaml --env-file "$ENV_FILE" stop
   rm -f "$DATA_DIR"/gm-session-manager.db-wal "$DATA_DIR"/gm-session-manager.db-shm
   tar -xzf "$backup" -C "$DATA_DIR"
   ```

5. The previous code and image, then start:

   ```sh
   git checkout --detach "$(cat "$DATA_DIR/deployed-commit")"
   docker tag gm-session-manager:prev gm-session-manager:current
   docker compose -f deploy/compose.yaml --env-file "$ENV_FILE" up -d
   ```

> The backup holds the database and the commit it belongs to, and `:prev` is the version that ran before the last
> update. Once `main` is fixed: `systemctl start gsm-update.timer`.
