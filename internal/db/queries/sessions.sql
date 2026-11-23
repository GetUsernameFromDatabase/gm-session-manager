-- name: CreateSession :one
INSERT INTO sessions (token_hash, role, character_id)
VALUES (?, ?, ?)
RETURNING *;

-- GetSessionByTokenHash skips a session of a deleted character: one issued
-- while the GM deleted it would otherwise leave the device neither player
-- nor visitor.
-- name: GetSessionByTokenHash :one
SELECT * FROM sessions
WHERE token_hash = ?
  AND (character_id IS NULL OR character_id IN (SELECT id FROM characters WHERE deleted_at IS NULL));

-- name: GetSession :one
SELECT * FROM sessions
WHERE id = ?;

-- TouchSession records when a device was last seen, for the GM's device
-- list: at most once a minute, so ordinary requests do not each cost a
-- write.
-- name: TouchSession :exec
UPDATE sessions
SET last_seen_at = CURRENT_TIMESTAMP
WHERE id = ? AND last_seen_at < datetime('now', '-1 minute');

-- LeaveSession records a device's page closing: when, and what it was, for
-- the GM's offline devices. A page left open sends no requests, so its last
-- one may be hours old.
-- name: LeaveSession :exec
UPDATE sessions
SET last_seen_at = CURRENT_TIMESTAMP, user_agent = ?
WHERE id = ?;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = ?;

-- name: ListPlayerSessionsByCampaign :many
SELECT s.id, s.character_id, s.last_seen_at, s.user_agent, s.created_at, c.name AS character_name
FROM sessions AS s
JOIN characters AS c ON c.id = s.character_id
WHERE c.campaign_id = ? AND c.deleted_at IS NULL
ORDER BY c.name, s.created_at;

-- name: ListSessionTokenHashes :many
SELECT token_hash FROM sessions;

-- name: DeleteCharacterSessions :exec
DELETE FROM sessions
WHERE character_id = ?;
