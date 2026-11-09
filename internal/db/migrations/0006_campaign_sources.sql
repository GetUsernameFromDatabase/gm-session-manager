-- The sources a campaign allows (C03): pickers and lists in the campaign show
-- only these, plus homebrew, which is always allowed and never stored. A new
-- campaign gets every source; a source imported later starts off everywhere.
CREATE TABLE campaign_sources (
    campaign_id INTEGER NOT NULL REFERENCES campaigns (id) ON DELETE CASCADE,
    source_key  TEXT    NOT NULL,
    PRIMARY KEY (campaign_id, source_key)
);

-- Existing campaigns keep everything they could see.
INSERT INTO campaign_sources (campaign_id, source_key)
SELECT c.id, s.source_key FROM campaigns AS c, compendium_sources AS s
WHERE s.source_key <> 'homebrew';

-- campaign_top replaces compendium_top: the copy of each (kind, slug) that
-- lists show in a campaign, ranked over the sources it allows (plus
-- homebrew), so an unticked book's entry falls back to the next allowed
-- copy. Campaign 0 ranks every entry (the GM outside a campaign). Filter on
-- campaign_id: SQLite pushes it into the ranking, so only that campaign's
-- rows are ranked.
DROP VIEW compendium_top;

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
        WHERE e.source_key = 'homebrew'
           OR EXISTS (SELECT 1 FROM campaign_sources AS cs WHERE cs.campaign_id = c.id AND cs.source_key = e.source_key)
    ) AS a
    JOIN compendium_entries AS e ON e.id = a.entry_id
    LEFT JOIN compendium_sources AS s ON s.source_key = e.source_key
)
WHERE rn = 1;
