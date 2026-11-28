-- Playing at home: taking a character home (a zip played in the player app,
-- brought back later), campaign starter files and characters made at home.

-- This database's random identity, written into take-home files so a player app
-- tells tables apart. In the database, not settings.json: it qualifies the
-- character ids, so it must travel and restore with them. One row only.
CREATE TABLE server_identity (
    only                INTEGER  PRIMARY KEY CHECK (only = 1),
    server_id           TEXT     NOT NULL CHECK (server_id <> '')
);

INSERT INTO server_identity (only, server_id) VALUES (1, lower(hex(randomblob(16))));

-- GM app: the merge base of an open trip, what the character looked like when
-- the file was written. returned_at is set when the bring-back is accepted.
CREATE TABLE take_homes (
    id                  INTEGER  PRIMARY KEY,
    character_id        INTEGER  NOT NULL REFERENCES characters (id) ON DELETE CASCADE,
    token               TEXT     NOT NULL UNIQUE CHECK (token <> ''),
    base_json           TEXT     NOT NULL CHECK (json_valid(base_json)),
    created_at          DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    returned_at         DATETIME
);

CREATE INDEX take_homes_character_id ON take_homes (character_id);

-- Player app: maps a home character to the table's ids, so a bring-back can name
-- what changed. manifest_json is the file's manifest plus the id maps.
CREATE TABLE home_copies (
    character_id        INTEGER  PRIMARY KEY REFERENCES characters (id) ON DELETE CASCADE,
    server_id           TEXT     NOT NULL CHECK (server_id <> ''),
    table_character_id  INTEGER  NOT NULL,
    manifest_json       TEXT     NOT NULL CHECK (json_valid(manifest_json)),
    -- The build-only copy carried inside the bring-back file.
    fallback_character_id INTEGER REFERENCES characters (id) ON DELETE SET NULL,
    UNIQUE (server_id, table_character_id),
    UNIQUE (fallback_character_id)
);

-- Player app: a campaign opened from a starter file. The campaign is the
-- player's own copy of the table's books; a second opening of the same table
-- campaign refreshes it in place.
CREATE TABLE home_starters (
    campaign_id         INTEGER  PRIMARY KEY REFERENCES campaigns (id) ON DELETE CASCADE,
    server_id           TEXT     NOT NULL CHECK (server_id <> ''),
    table_campaign_id   INTEGER  NOT NULL,
    saved_at            DATETIME NOT NULL,
    UNIQUE (server_id, table_campaign_id)
);

-- Both apps: at home the character made here, at the table the character added
-- from that file, so a later take-home can replace the made-at-home original.
CREATE TABLE new_character_tokens (
    character_id        INTEGER  PRIMARY KEY REFERENCES characters (id) ON DELETE CASCADE,
    token               TEXT     NOT NULL UNIQUE CHECK (token <> '')
);
