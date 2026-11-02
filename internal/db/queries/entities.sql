-- name: CreateEntity :one
INSERT INTO entities (campaign_id, kind, name, is_revealed, fields_json, compendium_ref)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetEntity :one
SELECT * FROM entities
WHERE id = ?;

-- name: ListEntitiesByCampaign :many
SELECT * FROM entities
WHERE campaign_id = ?
ORDER BY name;

-- ListRevealedEntities is the player listing. Field stripping still happens
-- in Go (perm.StripEntity) before anything is rendered.
-- name: ListRevealedEntities :many
SELECT * FROM entities
WHERE campaign_id = ? AND is_revealed = 1
ORDER BY name;

-- name: UpdateEntity :exec
UPDATE entities
SET name = ?, is_revealed = ?, fields_json = ?, compendium_ref = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;

-- name: DeleteEntity :exec
DELETE FROM entities
WHERE id = ?;

-- LinkEntityImage adds a look; the NPC's first look becomes its current one.
-- name: LinkEntityImage :exec
INSERT INTO entity_images (entity_id, image_id, is_current)
VALUES (sqlc.arg(entity_id), sqlc.arg(image_id),
        NOT EXISTS (SELECT 1 FROM entity_images AS c WHERE c.entity_id = sqlc.arg(entity_id) AND c.is_current))
ON CONFLICT (entity_id, image_id) DO NOTHING;

-- name: UnlinkEntityImage :exec
DELETE FROM entity_images
WHERE entity_id = ? AND image_id = ?;

-- name: ToggleEntityImageRevealed :exec
UPDATE entity_images
SET is_revealed = NOT is_revealed
WHERE entity_id = ? AND image_id = ?;

-- name: ClearEntityCurrentImage :exec
UPDATE entity_images
SET is_current = 0
WHERE entity_id = ? AND is_current;

-- name: SetEntityCurrentImage :execrows
UPDATE entity_images
SET is_current = 1
WHERE entity_id = ? AND image_id = ?;

-- name: ListEntityImages :many
SELECT i.id, i.name, i.url, i.attachment_id, l.is_revealed, l.is_current
FROM entity_images AS l
JOIN images AS i ON i.id = l.image_id
WHERE l.entity_id = ?
ORDER BY l.is_current DESC, i.name COLLATE NOCASE;

-- name: ListImageEntities :many
SELECT e.id, e.name
FROM entity_images AS l
JOIN entities AS e ON e.id = l.entity_id
WHERE l.image_id = ?
ORDER BY e.name COLLATE NOCASE;

-- name: DeleteEntityFieldLinks :exec
DELETE FROM entity_field_links
WHERE entity_id = ?;

-- name: AddEntityFieldLink :exec
INSERT INTO entity_field_links (entity_id, field_key, target_type, target_id, is_revealed)
VALUES (?, ?, ?, ?, ?)
ON CONFLICT DO NOTHING;

-- ListEntityFieldBacklinks lists the NPCs whose fields link a target: all
-- for the GM, for players only revealed links of revealed NPCs.
-- name: ListEntityFieldBacklinks :many
SELECT DISTINCT e.id, e.name
FROM entity_field_links AS l
JOIN entities AS e ON e.id = l.entity_id
WHERE l.target_type = sqlc.arg(target_type) AND l.target_id = sqlc.arg(target_id)
  AND (CAST(sqlc.arg(is_gm) AS BOOLEAN) OR (l.is_revealed AND e.is_revealed))
ORDER BY e.name COLLATE NOCASE;

-- CurrentEntityLooks is each NPC's current look in a fight.
-- name: CurrentEntityLooks :many
SELECT l.entity_id, e.is_revealed AS entity_revealed, i.id, i.name, i.url, i.attachment_id, l.is_revealed
FROM entity_images AS l
JOIN images AS i ON i.id = l.image_id
JOIN entities AS e ON e.id = l.entity_id
WHERE l.is_current AND l.entity_id IN (
    SELECT cb.entity_id FROM combatants AS cb WHERE cb.encounter_id = ? AND cb.entity_id IS NOT NULL
);

-- CurrentCampaignEntityLooks is each NPC's current look, for the NPC list.
-- name: CurrentCampaignEntityLooks :many
SELECT l.entity_id, i.id, i.name, i.url, i.attachment_id, l.is_revealed
FROM entity_images AS l
JOIN images AS i ON i.id = l.image_id
JOIN entities AS e ON e.id = l.entity_id
WHERE l.is_current AND e.campaign_id = ?;
