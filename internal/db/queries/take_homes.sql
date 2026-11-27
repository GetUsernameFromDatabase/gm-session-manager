-- This database's identity, written into take-home files.
-- name: GetServerID :one
SELECT server_id FROM server_identity;

-- name: CreateTakeHome :one
INSERT INTO take_homes (character_id, token, base_json)
VALUES (?, ?, ?)
RETURNING *;

-- name: LatestOpenTakeHome :one
SELECT * FROM take_homes
WHERE character_id = ? AND returned_at IS NULL
ORDER BY id DESC
LIMIT 1;

-- The newest open take-home of each character in a campaign.
-- name: ListOpenTakeHomesByCampaign :many
SELECT t.character_id, t.created_at
FROM take_homes AS t
JOIN characters AS c ON c.id = t.character_id
WHERE c.campaign_id = ? AND t.returned_at IS NULL
  AND t.id = (SELECT max(x.id) FROM take_homes AS x WHERE x.character_id = t.character_id AND x.returned_at IS NULL);

-- name: ListLevelUps :many
SELECT * FROM level_ups
WHERE character_id = ?
ORDER BY id;

-- name: ListOwnNotes :many
SELECT * FROM notes
WHERE owner_character_id = ?
ORDER BY id;

-- name: ListOwnNoteBlocks :many
SELECT b.* FROM note_blocks AS b
JOIN notes AS n ON n.id = b.note_id
WHERE n.owner_character_id = ?
ORDER BY b.note_id, b.position, b.id;

-- name: ListOwnNoteAudience :many
SELECT a.block_id, a.character_id FROM note_block_audience AS a
JOIN note_blocks AS b ON b.id = a.block_id
JOIN notes AS n ON n.id = b.note_id
WHERE n.owner_character_id = ?
ORDER BY a.block_id, a.character_id;

-- name: ListOwnNoteAttachments :many
SELECT a.* FROM attachments AS a
JOIN note_blocks AS b ON b.id = a.block_id
JOIN notes AS n ON n.id = b.note_id
WHERE n.owner_character_id = ?
ORDER BY a.id;

-- name: ListOwnHomebrew :many
SELECT * FROM compendium_entries
WHERE source_key = 'homebrew' AND owner_character_id = ?
ORDER BY id;

-- The campaign's rules content for a take-home: every copy of the entries of
-- its ticked sources (a source pin may name a copy that is not the top one),
-- and the homebrew it sees, except the character's own and monsters.
-- name: ListTakeHomeEntries :many
SELECT e.kind, e.slug, e.name, e.source_key, e.data_json, e.body_text, e.content_hash, e.scoped
FROM compendium_entries AS e
WHERE e.kind <> 'monster'
  AND (e.owner_character_id IS NULL OR e.owner_character_id <> sqlc.arg(character_id))
  AND (EXISTS (SELECT 1 FROM campaign_sources AS cs WHERE cs.campaign_id = sqlc.arg(campaign_id) AND cs.source_key = e.source_key)
       OR (e.source_key = 'homebrew'
           AND (e.scoped = 0
                OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = sqlc.arg(campaign_id)))))
ORDER BY e.id;

-- name: ListTakeHomeSources :many
SELECT s.source_key, s.edition, s.priority, s.published, s.license, s.attribution
FROM compendium_sources AS s
JOIN campaign_sources AS cs ON cs.source_key = s.source_key
WHERE cs.campaign_id = ?
ORDER BY s.source_key;

-- name: CreateLevelUpAt :one
INSERT INTO level_ups (character_id, max_hp, sheet_json, created_at)
VALUES (?, ?, ?, ?)
RETURNING id;

-- A character's own homebrew, now seen only in the home campaigns it is linked
-- to. A player app holds one table's characters, so a slug already here is the
-- same entry: the import updates it and sets its owner.
-- name: UpsertHomeHomebrew :one
INSERT INTO compendium_entries (
    kind, slug, name, source_key, data_json, body_text, content_hash, owner_character_id, scoped
) VALUES (?, ?, ?, 'homebrew', ?, ?, ?, ?, 1)
ON CONFLICT (kind, slug, source_key) DO UPDATE
SET name = excluded.name, data_json = excluded.data_json, body_text = excluded.body_text,
    content_hash = excluded.content_hash, owner_character_id = excluded.owner_character_id,
    scoped = 1, imported_at = CURRENT_TIMESTAMP
RETURNING id;

-- Homebrew the table's GM or another player wrote, seen in this home campaign.
-- Home copies share the rows, so an existing one keeps its owner and scope: the
-- caller only links it.
-- name: UpsertSharedHomebrew :one
INSERT INTO compendium_entries (
    kind, slug, name, source_key, data_json, body_text, content_hash, scoped
) VALUES (?, ?, ?, 'homebrew', ?, ?, ?, 1)
ON CONFLICT (kind, slug, source_key) DO UPDATE
SET name = excluded.name, data_json = excluded.data_json, body_text = excluded.body_text,
    content_hash = excluded.content_hash, imported_at = CURRENT_TIMESTAMP
RETURNING id;

-- A replaced home copy's own homebrew, except what another home campaign still
-- sees: that stays (its owner is cleared with the character, and the new copy's
-- import sets it again). Run before the home campaign is deleted.
-- name: DeleteHomebrewByOwner :exec
DELETE FROM compendium_entries
WHERE source_key = 'homebrew' AND owner_character_id = sqlc.arg(owner_character_id)
  AND NOT EXISTS (SELECT 1 FROM compendium_campaigns AS l
                  WHERE l.entry_id = compendium_entries.id AND l.campaign_id <> sqlc.arg(campaign_id));

-- name: CreateNoteAt :one
INSERT INTO notes (campaign_id, kind, title, owner_character_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id;

-- A block taken home starts a fresh collaborative document (version 0).
-- name: CreateTakeHomeBlock :one
INSERT INTO note_blocks (note_id, position, visibility, players_can_edit, doc_json, text, author_character_id)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING id;

-- name: CreateHomeCopy :exec
INSERT INTO home_copies (character_id, server_id, table_character_id, manifest_json)
VALUES (?, ?, ?, ?);

-- name: GetHomeCopy :one
SELECT * FROM home_copies
WHERE character_id = ?;

-- name: GetHomeCopyByTable :one
SELECT * FROM home_copies
WHERE server_id = ? AND table_character_id = ?;

-- name: HasOtherServerHomeCopy :one
SELECT EXISTS (SELECT 1 FROM home_copies WHERE server_id <> ?);

-- name: ListHomeCopies :many
SELECT h.character_id, ch.campaign_id, ch.name AS character_name, c.name AS campaign_name, h.manifest_json
FROM home_copies AS h
JOIN characters AS ch ON ch.id = h.character_id
JOIN campaigns AS c ON c.id = ch.campaign_id
ORDER BY ch.name COLLATE NOCASE, h.character_id;

-- SetCharacterImageFlags restores a look's flags as the file had them; clear
-- the current look first (LinkCharacterImage makes the first one current).
-- name: SetCharacterImageFlags :exec
UPDATE character_images
SET is_revealed = ?, is_current = ?
WHERE character_id = ? AND image_id = ?;

-- name: SetTakeHomeBlockDoc :exec
UPDATE note_blocks SET doc_json = ?
WHERE id = ?;

-- ListHomeSourceKeys lists the sources added at home: not the take-home's
-- marker path, and not homebrew (its path is '').
-- name: ListHomeSourceKeys :many
SELECT source_key FROM compendium_sources
WHERE path <> 'take-home' AND source_key <> 'homebrew'
ORDER BY source_key;

-- UpsertTakeHomeSource writes a source the take-home file lists. A source that
-- was added at home stays the home's: its row is kept as it is.
-- name: UpsertTakeHomeSource :exec
INSERT INTO compendium_sources (source_key, path, hash, edition, license, attribution, published, priority)
VALUES (sqlc.arg(source_key), 'take-home', '', sqlc.arg(edition), sqlc.arg(license), sqlc.arg(attribution), sqlc.arg(published), sqlc.arg(priority))
ON CONFLICT (source_key) DO UPDATE
SET edition = excluded.edition, license = excluded.license, attribution = excluded.attribution,
    published = excluded.published, priority = excluded.priority, last_imported = CURRENT_TIMESTAMP
WHERE compendium_sources.path = 'take-home';

-- SinkHomeSources: in the player app, sources added at home rank below every
-- source the take-home wrote, keeping their order among themselves. All move by
-- the same amount, and only when the lowest-numbered one is at or above the
-- table's highest number (so a second run does nothing).
-- name: SinkHomeSources :exec
UPDATE compendium_sources
SET priority = priority + 1 + (SELECT COALESCE(MAX(priority), 0) FROM compendium_sources WHERE path = 'take-home')
WHERE path <> 'take-home' AND source_key <> 'homebrew'
  AND (SELECT MIN(priority) FROM compendium_sources WHERE path <> 'take-home' AND source_key <> 'homebrew')
      <= (SELECT COALESCE(MAX(priority), 0) FROM compendium_sources WHERE path = 'take-home');
