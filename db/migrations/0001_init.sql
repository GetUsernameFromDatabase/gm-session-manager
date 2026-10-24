-- The whole schema. Login/device sessions live in sessions; game sessions are
-- a separate concept (notes of kind 'game_session').

CREATE TABLE campaigns (
    id         INTEGER  PRIMARY KEY,
    name       TEXT     NOT NULL CHECK (name <> ''),
    system     TEXT     NOT NULL,
    edition    TEXT     NOT NULL,
    open_join  BOOLEAN  NOT NULL DEFAULT 1 CHECK (open_join IN (0, 1)),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE characters (
    id          INTEGER  PRIMARY KEY,
    campaign_id INTEGER  NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    name        TEXT     NOT NULL CHECK (name <> ''),
    system      TEXT     NOT NULL,
    hp          INTEGER  NOT NULL DEFAULT 0,
    max_hp      INTEGER  NOT NULL DEFAULT 0,
    temp_hp     INTEGER  NOT NULL DEFAULT 0,
    sheet_json  TEXT     NOT NULL DEFAULT '{}' CHECK (json_valid(sheet_json)),
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Deleting a character only hides it (deleted_at set), so the GM can
    -- restore it with its notes and images; deleting it forever removes the row.
    deleted_at  DATETIME
);

CREATE INDEX characters_campaign_id ON characters (campaign_id);

-- One row per logged-in device. Only the SHA-256 of the cookie token is stored.
-- No unique index on character_id: several devices may hold one character.
CREATE TABLE sessions (
    id           INTEGER  PRIMARY KEY,
    token_hash   BLOB     NOT NULL UNIQUE CHECK (length(token_hash) = 32),
    role         TEXT     NOT NULL CHECK (role IN ('player', 'admin')),
    character_id INTEGER  REFERENCES characters (id) ON DELETE CASCADE,
    last_seen_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at   DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK ((role = 'player') = (character_id IS NOT NULL))
);

CREATE INDEX sessions_character_id ON sessions (character_id);

-- The compendium. Imported rules content and homebrew share one table;
-- homebrew uses source_key 'homebrew' so its slugs never collide with
-- imported ones.

CREATE TABLE compendium_sources (
    source_key    TEXT     PRIMARY KEY,
    path          TEXT     NOT NULL,
    hash          TEXT     NOT NULL,
    edition       TEXT     NOT NULL DEFAULT '',
    license       TEXT     NOT NULL DEFAULT '',
    attribution   TEXT     NOT NULL DEFAULT '',
    last_imported DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Matches homebrewSource in compendium/importer.go.
INSERT INTO compendium_sources (source_key, path, hash, edition, license, attribution)
VALUES ('homebrew', '', '', '2024', 'Homebrew', 'Homebrew content written for this table.');

CREATE TABLE compendium_entries (
    id           INTEGER  PRIMARY KEY,
    kind         TEXT     NOT NULL CHECK (kind IN (
                     'monster', 'spell', 'item', 'class', 'subclass', 'species',
                     'background', 'feat', 'condition', 'rule')),
    slug         TEXT     NOT NULL CHECK (slug <> ''),
    name         TEXT     NOT NULL CHECK (name <> ''),
    source_key   TEXT     NOT NULL CHECK (source_key <> ''),
    -- Our normalized shape for the kind (compendium/entry.go), never the raw
    -- shape of a data source.
    data_json    TEXT     NOT NULL DEFAULT '{}' CHECK (json_valid(data_json)),
    -- Plain text indexed for search alongside the name.
    body_text    TEXT     NOT NULL DEFAULT '',
    -- A monster's picture is an external link only; nothing is downloaded.
    -- NULL: never bound (the importer guesses one). '': the GM cleared it.
    image_url    TEXT,
    is_homebrew  BOOLEAN  NOT NULL DEFAULT 0 CHECK (is_homebrew IN (0, 1)),
    content_hash TEXT     NOT NULL,
    imported_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- Player homebrew: written by a character and seen only in the campaigns it
    -- is linked to (scoped). Imported entries and the GM's homebrew stay global.
    owner_character_id INTEGER REFERENCES characters (id) ON DELETE SET NULL,
    scoped       BOOLEAN  NOT NULL DEFAULT 0 CHECK (scoped IN (0, 1)),
    UNIQUE (kind, slug, source_key),
    CHECK ((is_homebrew = 1) = (source_key = 'homebrew'))
);

CREATE INDEX compendium_entries_kind_name ON compendium_entries (kind, name);

CREATE TABLE compendium_campaigns (
    entry_id    INTEGER NOT NULL REFERENCES compendium_entries (id) ON DELETE CASCADE,
    campaign_id INTEGER NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    PRIMARY KEY (entry_id, campaign_id)
);

CREATE INDEX compendium_campaigns_campaign_id ON compendium_campaigns (campaign_id);

-- A monster's habitats, typed in by the GM (sources carry none). Its own
-- table so re-imports, which keep entry ids, never touch it.
CREATE TABLE monster_habitats (
    entry_id INTEGER NOT NULL REFERENCES compendium_entries (id) ON DELETE CASCADE,
    habitat  TEXT    NOT NULL CHECK (habitat IN ('Any', 'Arctic', 'Coastal', 'Desert',
                 'Forest', 'Grassland', 'Hill', 'Mountain', 'Planar', 'Swamp',
                 'Underdark', 'Underwater', 'Urban')),
    PRIMARY KEY (entry_id, habitat)
);

-- External-content FTS5 index over name and body_text, kept in sync by the
-- triggers below.
CREATE VIRTUAL TABLE compendium_fts USING fts5 (
    name,
    body_text,
    content = 'compendium_entries',
    content_rowid = 'id',
    tokenize = 'unicode61 remove_diacritics 2'
);

CREATE TRIGGER compendium_entries_ai AFTER INSERT ON compendium_entries BEGIN
    INSERT INTO compendium_fts (rowid, name, body_text) VALUES (new.id, new.name, new.body_text);
END;

CREATE TRIGGER compendium_entries_ad AFTER DELETE ON compendium_entries BEGIN
    INSERT INTO compendium_fts (compendium_fts, rowid, name, body_text) VALUES ('delete', old.id, old.name, old.body_text);
END;

CREATE TRIGGER compendium_entries_au AFTER UPDATE ON compendium_entries BEGIN
    INSERT INTO compendium_fts (compendium_fts, rowid, name, body_text) VALUES ('delete', old.id, old.name, old.body_text);
    INSERT INTO compendium_fts (rowid, name, body_text) VALUES (new.id, new.name, new.body_text);
END;

-- Entities (NPCs first). fields_json maps a field key to
-- {"value": ..., "is_revealed": bool}; unrevealed fields never reach players.
CREATE TABLE entities (
    id             INTEGER  PRIMARY KEY,
    campaign_id    INTEGER  NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    kind           TEXT     NOT NULL DEFAULT 'npc' CHECK (kind IN ('npc')),
    name           TEXT     NOT NULL CHECK (name <> ''),
    is_revealed    BOOLEAN  NOT NULL DEFAULT 0 CHECK (is_revealed IN (0, 1)),
    fields_json    TEXT     NOT NULL DEFAULT '{}' CHECK (json_valid(fields_json)),
    compendium_ref INTEGER  REFERENCES compendium_entries (id) ON DELETE SET NULL,
    created_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX entities_campaign_id ON entities (campaign_id, name);

CREATE TABLE encounters (
    id                  INTEGER  PRIMARY KEY,
    campaign_id         INTEGER  NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    name                TEXT     NOT NULL CHECK (name <> ''),
    status              TEXT     NOT NULL DEFAULT 'preparing' CHECK (status IN ('preparing', 'active', 'ended')),
    round               INTEGER  NOT NULL DEFAULT 0 CHECK (round >= 0),
    active_combatant_id INTEGER  REFERENCES combatants (id) ON DELETE SET NULL,
    created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX encounters_campaign_id ON encounters (campaign_id);

-- A PC combatant references its character, which stays the single source of
-- truth for HP; the hp columns here are used for monsters and NPCs only.
-- marker ties a combatant to its piece on the physical table ("mini 5"); only
-- the GM sees it.
CREATE TABLE combatants (
    id           INTEGER PRIMARY KEY,
    encounter_id INTEGER NOT NULL REFERENCES encounters (id) ON DELETE CASCADE,
    character_id INTEGER REFERENCES characters (id) ON DELETE CASCADE,
    entity_id    INTEGER REFERENCES entities (id) ON DELETE SET NULL,
    monster_ref  INTEGER REFERENCES compendium_entries (id) ON DELETE SET NULL,
    name         TEXT    NOT NULL CHECK (name <> ''),
    initiative   INTEGER,
    hp           INTEGER NOT NULL DEFAULT 0,
    max_hp       INTEGER NOT NULL DEFAULT 0,
    temp_hp      INTEGER NOT NULL DEFAULT 0,
    ac           INTEGER NOT NULL DEFAULT 10,
    hp_visible   BOOLEAN NOT NULL DEFAULT 0 CHECK (hp_visible IN (0, 1)),
    sort_order   INTEGER NOT NULL DEFAULT 0,
    marker       TEXT    NOT NULL DEFAULT ''
);

CREATE INDEX combatants_encounter_id ON combatants (encounter_id, sort_order);
CREATE INDEX combatants_character_id ON combatants (character_id);

-- A condition belongs to exactly one of a character or a combatant. slug is a
-- 2024 condition, 'concentration', 'exhaustion' (level in note) or 'custom'.
-- expires_round is compared with the round of the encounter being advanced.
CREATE TABLE conditions (
    id            INTEGER  PRIMARY KEY,
    character_id  INTEGER  REFERENCES characters (id) ON DELETE CASCADE,
    combatant_id  INTEGER  REFERENCES combatants (id) ON DELETE CASCADE,
    slug          TEXT     NOT NULL CHECK (slug <> ''),
    note          TEXT     NOT NULL DEFAULT '',
    expires_round INTEGER,
    created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK ((character_id IS NULL) <> (combatant_id IS NULL))
);

CREATE INDEX conditions_character_id ON conditions (character_id);
CREATE INDEX conditions_combatant_id ON conditions (combatant_id);

-- Notes. A note is an ordered list of blocks; each block is its own
-- collaborative document with its own audience.
CREATE TABLE notes (
    id                 INTEGER  PRIMARY KEY,
    campaign_id        INTEGER  NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    kind               TEXT     NOT NULL CHECK (kind IN ('game_session', 'generic')),
    title              TEXT     NOT NULL CHECK (title <> ''),
    -- Set for notes a player created; NULL for the GM's notes.
    owner_character_id INTEGER  REFERENCES characters (id) ON DELETE CASCADE,
    created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX notes_campaign_id ON notes (campaign_id, updated_at);

CREATE TABLE note_blocks (
    id                  INTEGER PRIMARY KEY,
    note_id             INTEGER NOT NULL REFERENCES notes (id) ON DELETE CASCADE,
    position            INTEGER NOT NULL DEFAULT 0,
    visibility          TEXT    NOT NULL CHECK (visibility IN ('gm', 'players', 'listed', 'owner')),
    players_can_edit    BOOLEAN NOT NULL DEFAULT 0 CHECK (players_can_edit IN (0, 1)),
    -- ProseMirror JSON snapshot at doc_version; version is the latest step.
    doc_json            TEXT    NOT NULL DEFAULT '{"type":"doc","content":[{"type":"paragraph"}]}' CHECK (json_valid(doc_json)),
    doc_version         INTEGER NOT NULL DEFAULT 0,
    version             INTEGER NOT NULL DEFAULT 0,
    -- Plain text from the client's last snapshot, for search.
    text                TEXT    NOT NULL DEFAULT '',
    -- Players add blocks to the GM's game session notes and manage the blocks
    -- they wrote. NULL for blocks the GM wrote. A deleted character's blocks
    -- stay in the log and fall back to the note's manager.
    author_character_id INTEGER REFERENCES characters (id) ON DELETE SET NULL,
    CHECK (doc_version <= version)
);

CREATE INDEX note_blocks_note_id ON note_blocks (note_id, position);

-- 'owner' blocks belong to a player's note. Refuse them anywhere else so the
-- owner rule can never resolve to nobody, or to the GM.
CREATE TRIGGER note_blocks_owner_insert BEFORE INSERT ON note_blocks
WHEN new.visibility = 'owner'
     AND (SELECT owner_character_id FROM notes WHERE id = new.note_id) IS NULL
BEGIN
    SELECT RAISE(ABORT, 'owner blocks need a player-owned note');
END;

CREATE TRIGGER note_blocks_owner_update BEFORE UPDATE OF visibility ON note_blocks
WHEN new.visibility = 'owner'
     AND (SELECT owner_character_id FROM notes WHERE id = new.note_id) IS NULL
BEGIN
    SELECT RAISE(ABORT, 'owner blocks need a player-owned note');
END;

CREATE TABLE note_block_audience (
    block_id     INTEGER NOT NULL REFERENCES note_blocks (id) ON DELETE CASCADE,
    character_id INTEGER NOT NULL REFERENCES characters (id) ON DELETE CASCADE,
    PRIMARY KEY (block_id, character_id)
);

CREATE TABLE block_steps (
    block_id   INTEGER  NOT NULL REFERENCES note_blocks (id) ON DELETE CASCADE,
    version    INTEGER  NOT NULL,
    client_id  TEXT     NOT NULL,
    step_json  TEXT     NOT NULL CHECK (json_valid(step_json)),
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (block_id, version)
);

-- Links come from block snapshots, so each is tied to the block holding it
-- and backlinks can be filtered by what the viewer may read.
CREATE TABLE note_links (
    source_note_id  INTEGER NOT NULL REFERENCES notes (id) ON DELETE CASCADE,
    source_block_id INTEGER NOT NULL REFERENCES note_blocks (id) ON DELETE CASCADE,
    target_type     TEXT    NOT NULL CHECK (target_type IN ('note', 'entity', 'image')),
    target_id       INTEGER NOT NULL,
    PRIMARY KEY (source_block_id, target_type, target_id)
);

CREATE INDEX note_links_target ON note_links (target_type, target_id);

-- Files live under DATA_DIR/attachments with random names (path). A file
-- attached to a block is readable exactly when that block is.
CREATE TABLE attachments (
    id          INTEGER  PRIMARY KEY,
    campaign_id INTEGER  NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    block_id    INTEGER  REFERENCES note_blocks (id) ON DELETE CASCADE,
    path        TEXT     NOT NULL UNIQUE,
    name        TEXT     NOT NULL,
    mime        TEXT     NOT NULL,
    size        INTEGER  NOT NULL,
    created_at  DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX attachments_block_id ON attachments (block_id);

-- Images the GM shows on a display window. Names are unique per campaign
-- as typed: whether "Person" and "person" are one category is the GM's call.
CREATE TABLE image_categories (
    id          INTEGER PRIMARY KEY,
    campaign_id INTEGER NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    name        TEXT    NOT NULL CHECK (name <> ''),
    UNIQUE (campaign_id, name)
);

-- An image is an external URL or an uploaded file (an attachment with no
-- block, which only the image handlers serve to players).
CREATE TABLE images (
    id            INTEGER  PRIMARY KEY,
    campaign_id   INTEGER  NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    name          TEXT     NOT NULL CHECK (name <> ''),
    url           TEXT,
    attachment_id INTEGER  REFERENCES attachments (id) ON DELETE SET NULL,
    category_id   INTEGER  REFERENCES image_categories (id) ON DELETE SET NULL,
    description   TEXT,
    created_at    DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CHECK (url IS NOT NULL OR attachment_id IS NOT NULL)
);

CREATE INDEX images_campaign_id ON images (campaign_id, name);

-- An NPC's looks (outfits, disguises): images linked to it, each revealed to
-- players on its own. At most one is the current look, its portrait.
CREATE TABLE entity_images (
    entity_id   INTEGER NOT NULL REFERENCES entities (id) ON DELETE CASCADE,
    image_id    INTEGER NOT NULL REFERENCES images (id) ON DELETE CASCADE,
    is_revealed BOOLEAN NOT NULL DEFAULT 0 CHECK (is_revealed IN (0, 1)),
    is_current  BOOLEAN NOT NULL DEFAULT 0 CHECK (is_current IN (0, 1)),
    PRIMARY KEY (entity_id, image_id)
);

CREATE INDEX entity_images_image ON entity_images (image_id);

CREATE UNIQUE INDEX entity_images_current ON entity_images (entity_id) WHERE is_current;

-- A character's looks, like an NPC's: the owner reveals each one to the
-- campaign; the first is current and revealed.
CREATE TABLE character_images (
    character_id INTEGER NOT NULL REFERENCES characters (id) ON DELETE CASCADE,
    image_id     INTEGER NOT NULL REFERENCES images (id) ON DELETE CASCADE,
    is_revealed  BOOLEAN NOT NULL DEFAULT 0 CHECK (is_revealed IN (0, 1)),
    is_current   BOOLEAN NOT NULL DEFAULT 0 CHECK (is_current IN (0, 1)),
    PRIMARY KEY (character_id, image_id)
);

CREATE INDEX character_images_image ON character_images (image_id);

CREATE UNIQUE INDEX character_images_current ON character_images (character_id) WHERE is_current;

-- An image is an NPC's look or a character's, never both.
CREATE TRIGGER character_images_exclusive BEFORE INSERT ON character_images
WHEN EXISTS (SELECT 1 FROM entity_images WHERE image_id = NEW.image_id)
BEGIN
    SELECT RAISE(ABORT, 'image is an NPC look');
END;

CREATE TRIGGER entity_images_exclusive BEFORE INSERT ON entity_images
WHEN EXISTS (SELECT 1 FROM character_images WHERE image_id = NEW.image_id)
BEGIN
    SELECT RAISE(ABORT, 'image is a character look');
END;

-- [[links]] in NPC fields, rebuilt on every save: they feed the targets'
-- linked mentions, and a revealed field of a revealed NPC lets players open
-- the images it links.
CREATE TABLE entity_field_links (
    entity_id   INTEGER NOT NULL REFERENCES entities (id) ON DELETE CASCADE,
    field_key   TEXT    NOT NULL,
    target_type TEXT    NOT NULL CHECK (target_type IN ('note', 'entity', 'image')),
    target_id   INTEGER NOT NULL,
    is_revealed BOOLEAN NOT NULL DEFAULT 0 CHECK (is_revealed IN (0, 1)),
    PRIMARY KEY (entity_id, field_key, target_type, target_id)
);

CREATE INDEX entity_field_links_target ON entity_field_links (target_type, target_id);
