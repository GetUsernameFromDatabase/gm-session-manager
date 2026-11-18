-- Saved effect presets: the add-effect form minus its target and holder. The
-- GM's have no owner; a player's belong to their character. Private to the
-- creator, per campaign; the unique index lets a same-name save overwrite.
CREATE TABLE effect_presets (
    id                 INTEGER PRIMARY KEY,
    campaign_id        INTEGER NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    owner_character_id INTEGER REFERENCES characters (id) ON DELETE CASCADE,
    name               TEXT NOT NULL,
    slug               TEXT NOT NULL,
    note               TEXT NOT NULL DEFAULT '',
    ends_on            TEXT NOT NULL DEFAULT '' CHECK (ends_on IN ('', 'short_rest', 'long_rest')),
    rounds             INTEGER NOT NULL DEFAULT 0,
    mods_json          TEXT NOT NULL DEFAULT '[]' CHECK (json_valid(mods_json))
);

CREATE UNIQUE INDEX effect_presets_owner_name ON effect_presets (campaign_id, COALESCE(owner_character_id, 0), name);
CREATE INDEX effect_presets_owner_character_id ON effect_presets (owner_character_id);

-- The character whose player applied an effect to someone else, who may then
-- remove it; NULL for the GM's and for a player's own.
ALTER TABLE conditions ADD COLUMN applied_by INTEGER REFERENCES characters (id) ON DELETE SET NULL;
