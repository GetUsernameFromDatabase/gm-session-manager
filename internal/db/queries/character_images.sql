-- LinkCharacterImage adds a look; a character's first look is current and
-- revealed.
-- name: LinkCharacterImage :exec
INSERT INTO character_images (character_id, image_id, is_current, is_revealed)
VALUES (sqlc.arg(character_id), sqlc.arg(image_id),
        NOT EXISTS (SELECT 1 FROM character_images AS c WHERE c.character_id = sqlc.arg(character_id)),
        NOT EXISTS (SELECT 1 FROM character_images AS c WHERE c.character_id = sqlc.arg(character_id)))
ON CONFLICT (character_id, image_id) DO NOTHING;

-- name: UnlinkCharacterImage :exec
DELETE FROM character_images
WHERE character_id = ? AND image_id = ?;

-- name: ToggleCharacterImageRevealed :exec
UPDATE character_images
SET is_revealed = NOT is_revealed
WHERE character_id = ? AND image_id = ?;

-- name: ClearCharacterCurrentImage :exec
UPDATE character_images
SET is_current = 0
WHERE character_id = ? AND is_current;

-- name: SetCharacterCurrentImage :execrows
UPDATE character_images
SET is_current = 1
WHERE character_id = ? AND image_id = ?;

-- name: ListCharacterImages :many
SELECT i.id, i.name, i.url, i.attachment_id, l.is_revealed, l.is_current
FROM character_images AS l
JOIN images AS i ON i.id = l.image_id
WHERE l.character_id = ?
ORDER BY l.is_current DESC, i.name COLLATE NOCASE;

-- name: ListImageCharacters :many
SELECT ch.id, ch.name
FROM character_images AS l
JOIN characters AS ch ON ch.id = l.character_id
WHERE l.image_id = ?
ORDER BY ch.name COLLATE NOCASE;

-- CurrentCharacterLooks is each character's current look in a fight.
-- name: CurrentCharacterLooks :many
SELECT l.character_id, i.id, i.name, i.url, i.attachment_id, l.is_revealed
FROM character_images AS l
JOIN images AS i ON i.id = l.image_id
WHERE l.is_current AND l.character_id IN (
    SELECT cb.character_id FROM combatants AS cb WHERE cb.encounter_id = ? AND cb.character_id IS NOT NULL
);

-- ListImageLookAccess says, for the campaign's images (one image, or every
-- one for 0), who besides the GM may see each through NPCs and characters:
-- every player when it is a revealed look (of a revealed NPC, or of a
-- character) or linked from a revealed field of a revealed NPC, and the
-- character who wears it.
-- name: ListImageLookAccess :many
SELECT
    i.id,
    CAST(EXISTS (
        SELECT 1 FROM entity_images AS l
        JOIN entities AS e ON e.id = l.entity_id
        WHERE l.image_id = i.id AND l.is_revealed AND e.is_revealed
    ) OR EXISTS (
        SELECT 1 FROM character_images AS l
        WHERE l.image_id = i.id AND l.is_revealed
    ) OR EXISTS (
        SELECT 1 FROM entity_field_links AS l
        JOIN entities AS e ON e.id = l.entity_id
        WHERE l.target_image_id = i.id AND l.is_revealed AND e.is_revealed
    ) AS BOOLEAN) AS revealed,
    CAST(COALESCE((SELECT l.character_id FROM character_images AS l WHERE l.image_id = i.id), 0) AS INTEGER) AS owner_character_id
FROM images AS i
WHERE i.campaign_id = sqlc.arg(campaign_id)
  AND (CAST(sqlc.arg(image_id) AS INTEGER) = 0 OR i.id = CAST(sqlc.arg(image_id) AS INTEGER));

-- ImageBinding counts the NPCs and characters wearing an image.
-- name: ImageBinding :one
SELECT
    (SELECT COUNT(*) FROM entity_images AS e WHERE e.image_id = sqlc.arg(image_id)) AS npcs,
    (SELECT COUNT(*) FROM character_images AS c WHERE c.image_id = sqlc.arg(image_id)) AS characters;
