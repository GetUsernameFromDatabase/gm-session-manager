-- name: CreateCampaign :one
INSERT INTO campaigns (name, system, edition)
VALUES (?, ?, ?)
RETURNING *;

-- name: GetCampaign :one
SELECT * FROM campaigns
WHERE id = ?;

-- name: ListCampaigns :many
SELECT * FROM campaigns
ORDER BY name;

-- name: SetCampaignOpenJoin :exec
UPDATE campaigns
SET open_join = ?
WHERE id = ?;
