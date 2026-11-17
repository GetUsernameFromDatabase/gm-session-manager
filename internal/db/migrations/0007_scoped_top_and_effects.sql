-- campaign_top now also applies the scoping: player homebrew (scoped) ranks only
-- in the campaigns it is linked to, so queries no longer repeat that check.
-- Campaign 0 (the GM outside a campaign) still ranks every entry.
-- SearchCompendium repeats this ranking over its hits; change both.
DROP VIEW campaign_top;

CREATE VIEW campaign_top AS
SELECT campaign_id, id FROM (
    SELECT a.campaign_id, e.id,
        ROW_NUMBER() OVER (
            PARTITION BY a.campaign_id, e.kind, e.slug
            ORDER BY e.scoped, COALESCE(s.priority, 300), COALESCE(s.published, '') DESC, e.id
        ) AS rn
    FROM (
        SELECT 0 AS campaign_id, id AS entry_id FROM compendium_entries
        UNION ALL
        SELECT c.id, e.id FROM campaigns AS c
        JOIN compendium_entries AS e
        WHERE (e.source_key = 'homebrew'
               OR EXISTS (SELECT 1 FROM campaign_sources AS cs WHERE cs.campaign_id = c.id AND cs.source_key = e.source_key))
          AND (e.scoped = 0
               OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = c.id))
    ) AS a
    JOIN compendium_entries AS e ON e.id = a.entry_id
    LEFT JOIN compendium_sources AS s ON s.source_key = e.source_key
)
WHERE rn = 1;

-- For the source-allowed joins and the owner's homebrew lookups.
CREATE INDEX compendium_entries_source_key ON compendium_entries (source_key);
CREATE INDEX compendium_entries_owner_character_id ON compendium_entries (owner_character_id);

-- A condition row is also an effect: it can end on a rest, hang off a
-- concentration row (held_by: deleting that row removes it) and carry
-- modifiers (a list of CustomMod) that are shown with the sheet, never saved
-- into sheet_json.
ALTER TABLE conditions ADD COLUMN ends_on TEXT NOT NULL DEFAULT '' CHECK (ends_on IN ('', 'short_rest', 'long_rest'));
ALTER TABLE conditions ADD COLUMN held_by INTEGER REFERENCES conditions (id) ON DELETE CASCADE;
ALTER TABLE conditions ADD COLUMN mods_json TEXT NOT NULL DEFAULT '[]' CHECK (json_valid(mods_json));
CREATE INDEX conditions_held_by ON conditions (held_by);
