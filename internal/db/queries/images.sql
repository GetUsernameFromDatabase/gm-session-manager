-- name: CreateImage :one
INSERT INTO images (campaign_id, name, url, attachment_id, category_id, description)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: GetImage :one
SELECT i.*, COALESCE(c.name, '') AS category
FROM images AS i
LEFT JOIN image_categories AS c ON c.id = i.category_id
WHERE i.id = ?;

-- ListImages filters by name (empty q matches all; instr, unlike LIKE, has
-- no wildcards) and category (0 for any).
-- name: ListImages :many
SELECT i.*, COALESCE(c.name, '') AS category
FROM images AS i
LEFT JOIN image_categories AS c ON c.id = i.category_id
WHERE i.campaign_id = sqlc.arg(campaign_id)
  AND instr(lower(i.name), lower(CAST(sqlc.arg(q) AS TEXT))) > 0
  AND (CAST(sqlc.arg(category_id) AS INTEGER) = 0 OR i.category_id = CAST(sqlc.arg(category_id) AS INTEGER))
ORDER BY i.name COLLATE NOCASE;

-- name: UpdateImage :exec
UPDATE images
SET name = ?, url = ?, attachment_id = ?, category_id = ?, description = ?
WHERE id = ?;

-- name: DeleteImage :exec
DELETE FROM images
WHERE id = ?;

-- name: ListImageCategories :many
SELECT * FROM image_categories
WHERE campaign_id = ?
ORDER BY name COLLATE NOCASE;

-- EnsureImageCategory returns the category with exactly this name,
-- creating it if needed.
-- name: EnsureImageCategory :one
INSERT INTO image_categories (campaign_id, name)
VALUES (?, ?)
ON CONFLICT (campaign_id, name) DO UPDATE SET name = excluded.name
RETURNING id;

-- name: DeleteUnusedImageCategories :exec
DELETE FROM image_categories
WHERE image_categories.campaign_id = ? AND image_categories.id NOT IN (SELECT i.category_id FROM images AS i WHERE i.category_id IS NOT NULL);

-- ListImageLinkBlocks lists the note blocks linking the campaign's images
-- (one image, or every one for 0) with what the read rule needs.
-- name: ListImageLinkBlocks :many
SELECT DISTINCT l.target_id AS image_id, b.id AS block_id, b.visibility, n.owner_character_id
FROM note_links AS l
JOIN note_blocks AS b ON b.id = l.source_block_id
JOIN notes AS n ON n.id = b.note_id
WHERE l.target_type = 'image' AND n.campaign_id = sqlc.arg(campaign_id)
  AND (CAST(sqlc.arg(image_id) AS INTEGER) = 0 OR l.target_id = CAST(sqlc.arg(image_id) AS INTEGER));

-- ListImageLinkAudience lists the audiences of the blocks ListImageLinkBlocks returns.
-- name: ListImageLinkAudience :many
SELECT DISTINCT a.block_id, a.character_id
FROM note_links AS l
JOIN note_blocks AS b ON b.id = l.source_block_id
JOIN notes AS n ON n.id = b.note_id
JOIN note_block_audience AS a ON a.block_id = b.id
WHERE l.target_type = 'image' AND n.campaign_id = sqlc.arg(campaign_id)
  AND (CAST(sqlc.arg(image_id) AS INTEGER) = 0 OR l.target_id = CAST(sqlc.arg(image_id) AS INTEGER));

-- name: DeleteAttachment :exec
DELETE FROM attachments
WHERE id = ?;
