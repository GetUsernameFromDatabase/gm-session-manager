-- An attachment outlives the block it was uploaded to, since other notes may
-- still show it; the admin Files page deletes it once nothing uses it. With
-- no block it is the GM's, so files of 'owner' blocks still go with them. Only the foreign key action changes, so the
-- stored schema is edited in place (sqlite.org/lang_altertable.html): a
-- rebuild would drop the table and null every images.attachment_id.
PRAGMA writable_schema = ON;

UPDATE sqlite_schema
SET sql = replace(sql, 'REFERENCES note_blocks (id) ON DELETE CASCADE', 'REFERENCES note_blocks (id) ON DELETE SET NULL')
WHERE type = 'table' AND name = 'attachments';

PRAGMA writable_schema = OFF;

-- A player's own block keeps its files private for good: with no block they
-- would become the GM's, who must never read them.
CREATE TRIGGER note_blocks_owner_delete BEFORE DELETE ON note_blocks
WHEN old.visibility = 'owner'
BEGIN
    DELETE FROM attachments WHERE block_id = old.id;
END;
