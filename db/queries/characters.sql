-- name: CreateCharacter :one
INSERT INTO characters (campaign_id, name, system)
VALUES (?, ?, ?)
RETURNING *;

-- name: CreateCharacterWithSheet :one
INSERT INTO characters (campaign_id, name, system, hp, max_hp, temp_hp, sheet_json)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- GetCharacter and ListCharactersByCampaign see only live characters; a
-- deleted one is gone everywhere until the GM restores it.
-- name: GetCharacter :one
SELECT * FROM characters
WHERE id = ? AND deleted_at IS NULL;

-- name: ListCharactersByCampaign :many
SELECT * FROM characters
WHERE campaign_id = ? AND deleted_at IS NULL
ORDER BY name;

-- GetCharacterAny also finds a deleted character, for the GM's restore,
-- download and delete forever.
-- name: GetCharacterAny :one
SELECT * FROM characters
WHERE id = ?;

-- name: ListDeletedCharacters :many
SELECT * FROM characters
WHERE campaign_id = ? AND deleted_at IS NOT NULL
ORDER BY deleted_at DESC;

-- name: SoftDeleteCharacter :exec
UPDATE characters SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ? AND deleted_at IS NULL;

-- name: RestoreCharacter :exec
UPDATE characters SET deleted_at = NULL
WHERE id = ?;

-- UpdateCharacter writes the result of one command. It runs inside the
-- command's transaction, after the row was read in the same transaction.
-- name: UpdateCharacter :exec
UPDATE characters
SET name = ?, hp = ?, max_hp = ?, temp_hp = ?, sheet_json = ?
WHERE id = ?;

-- ListCharacterSheets is every character's sheet, deleted ones too, for
-- the startup sheet upgrade.
-- name: ListCharacterSheets :many
SELECT id, sheet_json FROM characters
ORDER BY id;

-- name: UpdateCharacterSheet :exec
UPDATE characters SET sheet_json = ?
WHERE id = ?;

-- name: DeleteCharacter :exec
DELETE FROM characters
WHERE id = ?;

-- Level-up snapshots, newest reverted first.

-- name: CreateLevelUp :exec
INSERT INTO level_ups (character_id, max_hp, sheet_json)
VALUES (?, ?, ?);

-- name: LatestLevelUp :one
SELECT * FROM level_ups
WHERE character_id = ?
ORDER BY id DESC
LIMIT 1;

-- name: CountLevelUps :one
SELECT count(*) FROM level_ups
WHERE character_id = ?;

-- name: DeleteLevelUp :exec
DELETE FROM level_ups
WHERE id = ?;
