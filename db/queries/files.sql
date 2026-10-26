-- The admin Files page: every attachment and where it is used. Uses are
-- found in Go (web/files.go) from the texts below.

-- ListAttachmentFiles lists every attachment with its campaign and, for one
-- in a note block, that block's visibility.
-- name: ListAttachmentFiles :many
SELECT a.id, a.campaign_id, a.block_id, a.path, a.name, a.mime, a.size, a.created_at,
       c.name AS campaign_name, b.visibility AS block_visibility
FROM attachments AS a
JOIN campaigns AS c ON c.id = a.campaign_id
LEFT JOIN note_blocks AS b ON b.id = a.block_id
ORDER BY a.created_at DESC, a.id DESC;

-- name: ListBlockDocs :many
SELECT b.id, b.note_id, b.visibility, n.title, b.doc_json
FROM note_blocks AS b
JOIN notes AS n ON n.id = b.note_id
ORDER BY n.title, b.position;

-- ListStepJSON covers edits not yet in a block's snapshot.
-- name: ListStepJSON :many
SELECT block_id, step_json FROM block_steps;

-- name: ListImageAttachments :many
SELECT id, name, attachment_id FROM images
WHERE attachment_id IS NOT NULL
ORDER BY name;
