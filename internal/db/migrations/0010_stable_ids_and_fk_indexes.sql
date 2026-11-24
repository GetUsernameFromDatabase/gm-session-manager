-- Image, note and NPC ids never come back: links (note_links,
-- entity_field_links) store a bare target id, so a reused id would point an
-- old link, and its backlink, at a new row.
-- AUTOINCREMENT is added by editing the stored schema, as in 0003: a rebuild
-- would drop the tables and cascade away their looks, blocks and links.

-- Creating an AUTOINCREMENT table makes SQLite create sqlite_sequence.
CREATE TABLE seq_boot (id INTEGER PRIMARY KEY AUTOINCREMENT);
DROP TABLE seq_boot;

-- Start past every id already used, including ones only stale links remember.
INSERT INTO sqlite_sequence (name, seq)
SELECT 'images', COALESCE(max(n), 0) FROM (
    SELECT max(id) AS n FROM images
    UNION ALL SELECT max(target_id) FROM note_links WHERE target_type = 'image'
    UNION ALL SELECT max(target_id) FROM entity_field_links WHERE target_type = 'image'
);

INSERT INTO sqlite_sequence (name, seq)
SELECT 'notes', COALESCE(max(n), 0) FROM (
    SELECT max(id) AS n FROM notes
    UNION ALL SELECT max(target_id) FROM note_links WHERE target_type = 'note'
    UNION ALL SELECT max(target_id) FROM entity_field_links WHERE target_type = 'note'
);

INSERT INTO sqlite_sequence (name, seq)
SELECT 'entities', COALESCE(max(n), 0) FROM (
    SELECT max(id) AS n FROM entities
    UNION ALL SELECT max(target_id) FROM note_links WHERE target_type = 'entity'
    UNION ALL SELECT max(target_id) FROM entity_field_links WHERE target_type = 'entity'
);

PRAGMA writable_schema = ON;

UPDATE sqlite_schema
SET sql = replace(sql, 'INTEGER  PRIMARY KEY,', 'INTEGER  PRIMARY KEY AUTOINCREMENT,')
WHERE type = 'table' AND name IN ('images', 'notes', 'entities');

PRAGMA writable_schema = OFF;

-- The search index holds only name and body_text; other updates leave it alone.
DROP TRIGGER compendium_entries_au;

CREATE TRIGGER compendium_entries_au AFTER UPDATE OF name, body_text ON compendium_entries BEGIN
    INSERT INTO compendium_fts (compendium_fts, rowid, name, body_text) VALUES ('delete', old.id, old.name, old.body_text);
    INSERT INTO compendium_fts (rowid, name, body_text) VALUES (new.id, new.name, new.body_text);
END;

-- Every foreign key gets an index, so deleting a parent row (a book's
-- entries, a character, a note) does not scan the child table.
CREATE INDEX entities_compendium_ref ON entities (compendium_ref);
CREATE INDEX encounters_active_combatant_id ON encounters (active_combatant_id);
CREATE INDEX combatants_entity_id ON combatants (entity_id);
CREATE INDEX combatants_monster_ref ON combatants (monster_ref);
CREATE INDEX conditions_applied_by ON conditions (applied_by);
CREATE INDEX notes_owner_character_id ON notes (owner_character_id);
CREATE INDEX note_blocks_author_character_id ON note_blocks (author_character_id);
CREATE INDEX note_block_audience_character_id ON note_block_audience (character_id);
CREATE INDEX note_links_source_note_id ON note_links (source_note_id);
CREATE INDEX attachments_campaign_id ON attachments (campaign_id);
CREATE INDEX images_attachment_id ON images (attachment_id);
CREATE INDEX images_category_id ON images (category_id);
