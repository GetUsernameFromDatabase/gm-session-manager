-- name: CreateSession :one
INSERT INTO sessions (token_hash, role, character_id)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetSessionByTokenHash :one
SELECT * FROM sessions
WHERE token_hash = ?;

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = ?;

-- TouchSession updates last_seen_at at most once a minute so ordinary
-- requests do not each cost a write.
-- name: TouchSession :exec
UPDATE sessions
SET last_seen_at = CURRENT_TIMESTAMP
WHERE id = ? AND last_seen_at < datetime('now', '-1 minute');

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = ?;

-- name: ListPlayerSessionsByCampaign :many
SELECT s.id, s.character_id, s.last_seen_at, s.created_at, c.name AS character_name
FROM sessions AS s
JOIN characters AS c ON c.id = s.character_id
WHERE c.campaign_id = ? AND c.deleted_at IS NULL
ORDER BY c.name, s.created_at;

-- name: DeleteCharacterSessions :exec
DELETE FROM sessions
WHERE character_id = ?;
