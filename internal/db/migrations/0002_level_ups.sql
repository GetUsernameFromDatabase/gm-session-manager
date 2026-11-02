-- A character's sheet and hit points as they were before each level-up, so
-- level-ups can be reverted, newest first.
CREATE TABLE level_ups (
    id           INTEGER  PRIMARY KEY,
    character_id INTEGER  NOT NULL REFERENCES characters (id) ON DELETE CASCADE,
    max_hp       INTEGER  NOT NULL,
    sheet_json   TEXT     NOT NULL CHECK (json_valid(sheet_json)),
    created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX level_ups_character_id ON level_ups (character_id);
