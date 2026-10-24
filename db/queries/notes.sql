-- Read rules for note blocks live in two places that must agree: the
-- readable-block predicate below (for listings) and perm.CanReadBlock (for
-- single checks and websocket fan-out). perm's table-driven tests run both.
--
-- A viewer is (is_gm, character_id). The GM never reads 'owner' blocks.

-- name: CreateNote :one
INSERT INTO notes (campaign_id, kind, title, owner_character_id)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: GetNote :one
SELECT * FROM notes
WHERE id = ?;

-- name: RenameNote :exec
UPDATE notes
SET title = ?, kind = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: TouchNote :exec
UPDATE notes
SET updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteNote :exec
DELETE FROM notes
WHERE id = ?;

-- ListReadableNotes lists notes with at least one block the viewer may read,
-- plus the viewer's own notes and, for the GM, the GM's notes (which may
-- still be empty).
-- name: ListReadableNotes :many
SELECT n.* FROM notes AS n
WHERE n.campaign_id = sqlc.arg(campaign_id)
  AND (
    (CAST(sqlc.arg(is_gm) AS BOOLEAN) AND n.owner_character_id IS NULL)
    OR (NOT CAST(sqlc.arg(is_gm) AS BOOLEAN) AND n.owner_character_id = CAST(sqlc.arg(character_id) AS INTEGER))
    OR EXISTS (
      SELECT 1 FROM note_blocks AS b
      WHERE b.note_id = n.id
        AND (
          (CAST(sqlc.arg(is_gm) AS BOOLEAN) AND b.visibility <> 'owner')
          OR (NOT CAST(sqlc.arg(is_gm) AS BOOLEAN) AND CAST(sqlc.arg(character_id) AS INTEGER) <> 0 AND (
                b.visibility = 'players'
                OR (b.visibility = 'listed' AND EXISTS (
                      SELECT 1 FROM note_block_audience AS a
                      WHERE a.block_id = b.id AND a.character_id = CAST(sqlc.arg(character_id) AS INTEGER)))
                OR (b.visibility = 'owner' AND n.owner_character_id = CAST(sqlc.arg(character_id) AS INTEGER))
          ))
        )
    )
  )
ORDER BY n.updated_at DESC
LIMIT sqlc.arg(max_results);

-- name: ListReadableBlocks :many
SELECT b.* FROM note_blocks AS b
JOIN notes AS n ON n.id = b.note_id
WHERE b.note_id = sqlc.arg(note_id)
  AND (
    (CAST(sqlc.arg(is_gm) AS BOOLEAN) AND b.visibility <> 'owner')
    OR (NOT CAST(sqlc.arg(is_gm) AS BOOLEAN) AND CAST(sqlc.arg(character_id) AS INTEGER) <> 0 AND (
          b.visibility = 'players'
          OR (b.visibility = 'listed' AND EXISTS (
                SELECT 1 FROM note_block_audience AS a
                WHERE a.block_id = b.id AND a.character_id = CAST(sqlc.arg(character_id) AS INTEGER)))
          OR (b.visibility = 'owner' AND n.owner_character_id = CAST(sqlc.arg(character_id) AS INTEGER))
    ))
  )
ORDER BY b.position, b.id;

-- name: SearchReadableBlocks :many
SELECT n.id AS note_id, n.title, b.id AS block_id, b.text
FROM note_blocks AS b
JOIN notes AS n ON n.id = b.note_id
WHERE n.campaign_id = sqlc.arg(campaign_id)
  AND (n.title LIKE sqlc.arg(pattern) OR b.text LIKE sqlc.arg(pattern))
  AND (
    (CAST(sqlc.arg(is_gm) AS BOOLEAN) AND b.visibility <> 'owner')
    OR (NOT CAST(sqlc.arg(is_gm) AS BOOLEAN) AND CAST(sqlc.arg(character_id) AS INTEGER) <> 0 AND (
          b.visibility = 'players'
          OR (b.visibility = 'listed' AND EXISTS (
                SELECT 1 FROM note_block_audience AS a
                WHERE a.block_id = b.id AND a.character_id = CAST(sqlc.arg(character_id) AS INTEGER)))
          OR (b.visibility = 'owner' AND n.owner_character_id = CAST(sqlc.arg(character_id) AS INTEGER))
    ))
  )
ORDER BY n.updated_at DESC
LIMIT sqlc.arg(max_results);

-- name: GetBlock :one
SELECT b.*, n.campaign_id, n.owner_character_id
FROM note_blocks AS b
JOIN notes AS n ON n.id = b.note_id
WHERE b.id = ?;

-- name: ListBlockAudience :many
SELECT character_id FROM note_block_audience
WHERE block_id = ?
ORDER BY character_id;

-- name: ClearBlockAudience :exec
DELETE FROM note_block_audience
WHERE block_id = ?;

-- name: AddBlockAudience :exec
INSERT OR IGNORE INTO note_block_audience (block_id, character_id)
SELECT sqlc.arg(block_id), c.id FROM characters AS c
JOIN notes AS n ON n.campaign_id = c.campaign_id
JOIN note_blocks AS b ON b.note_id = n.id
WHERE b.id = sqlc.arg(block_id) AND c.id = CAST(sqlc.arg(character_id) AS INTEGER);

-- name: CreateBlock :one
INSERT INTO note_blocks (note_id, position, visibility, players_can_edit, author_character_id)
VALUES (
    sqlc.arg(note_id),
    (SELECT coalesce(max(nb.position), 0) + 1 FROM note_blocks AS nb WHERE nb.note_id = sqlc.arg(note_id)),
    sqlc.arg(visibility),
    sqlc.arg(players_can_edit),
    sqlc.narg(author_character_id)
)
RETURNING *;

-- name: UpdateBlockAccess :exec
UPDATE note_blocks
SET visibility = ?, players_can_edit = ?
WHERE id = ?;

-- name: SetBlockPosition :exec
UPDATE note_blocks
SET position = ?
WHERE id = ?;

-- name: DeleteBlock :exec
DELETE FROM note_blocks
WHERE id = ?;

-- Collaboration (ProseMirror collab, central authority)

-- name: InsertBlockStep :exec
INSERT INTO block_steps (block_id, version, client_id, step_json)
VALUES (?, ?, ?, ?);

-- name: SetBlockVersion :exec
UPDATE note_blocks
SET version = ?
WHERE id = ?;

-- name: ListBlockStepsSince :many
SELECT version, client_id, step_json FROM block_steps
WHERE block_id = ? AND version > ?
ORDER BY version;

-- name: SaveBlockSnapshot :execrows
UPDATE note_blocks
SET doc_json = sqlc.arg(doc_json), doc_version = sqlc.arg(doc_version), text = sqlc.arg(text)
WHERE id = sqlc.arg(id) AND sqlc.arg(doc_version) > doc_version AND sqlc.arg(doc_version) <= version;

-- name: TrimBlockSteps :exec
DELETE FROM block_steps
WHERE block_id = ? AND version <= ?;

-- Links and backlinks

-- name: ClearBlockLinks :exec
DELETE FROM note_links
WHERE source_block_id = ?;

-- name: AddBlockLink :exec
INSERT OR IGNORE INTO note_links (source_note_id, source_block_id, target_type, target_id)
VALUES (?, ?, ?, ?);

-- ListBacklinks returns the notes linking to a target through a block the
-- viewer may read, so a secret block never reveals that it mentions you.
-- name: ListBacklinks :many
SELECT DISTINCT n.id, n.title
FROM note_links AS l
JOIN note_blocks AS b ON b.id = l.source_block_id
JOIN notes AS n ON n.id = l.source_note_id
WHERE l.target_type = sqlc.arg(target_type) AND l.target_id = sqlc.arg(target_id)
  AND n.campaign_id = sqlc.arg(campaign_id)
  AND (
    (CAST(sqlc.arg(is_gm) AS BOOLEAN) AND b.visibility <> 'owner')
    OR (NOT CAST(sqlc.arg(is_gm) AS BOOLEAN) AND CAST(sqlc.arg(character_id) AS INTEGER) <> 0 AND (
          b.visibility = 'players'
          OR (b.visibility = 'listed' AND EXISTS (
                SELECT 1 FROM note_block_audience AS a
                WHERE a.block_id = b.id AND a.character_id = CAST(sqlc.arg(character_id) AS INTEGER)))
          OR (b.visibility = 'owner' AND n.owner_character_id = CAST(sqlc.arg(character_id) AS INTEGER))
    ))
  )
ORDER BY n.title;

-- Attachments

-- name: CreateAttachment :one
INSERT INTO attachments (campaign_id, block_id, path, name, mime, size)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetAttachment :one
SELECT * FROM attachments
WHERE id = ?;

-- name: ListAttachmentPaths :many
SELECT path FROM attachments;

-- Export

-- ListExportBlocks returns every block of a campaign except 'owner' blocks,
-- which never leave the player who wrote them.
-- name: ListExportBlocks :many
SELECT b.id, b.note_id, b.position, b.visibility, b.doc_json
FROM note_blocks AS b
JOIN notes AS n ON n.id = b.note_id
WHERE n.campaign_id = ? AND b.visibility <> 'owner'
ORDER BY b.note_id, b.position, b.id;

-- name: ListExportNotes :many
SELECT * FROM notes
WHERE campaign_id = ?
ORDER BY id;

-- name: ListAttachmentsByCampaign :many
SELECT a.* FROM attachments AS a
LEFT JOIN note_blocks AS b ON b.id = a.block_id
WHERE a.campaign_id = ? AND (b.id IS NULL OR b.visibility <> 'owner')
ORDER BY a.id;
