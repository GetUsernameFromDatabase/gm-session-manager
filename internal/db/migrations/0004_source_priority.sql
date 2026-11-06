-- Several sources can hold the same (kind, slug); the source with the lowest
-- priority number wins every lookup and shows once in lists. The GM may
-- reorder sources, so an upsert never rewrites it.
-- The default is by group (homebrew, SRD, 2024 books, others); 0005 orders
-- each group by publication date. Keep in step with UpsertCompendiumSource in queries/compendium.sql.
ALTER TABLE compendium_sources ADD COLUMN priority INTEGER NOT NULL DEFAULT 300;

UPDATE compendium_sources
SET priority = CASE
    WHEN source_key = 'homebrew' THEN 0
    WHEN source_key LIKE 'srd%' THEN 200
    WHEN edition = '2024' THEN 100
    ELSE 300
END;

-- The copy of each (kind, slug) that lists show: an imported or GM entry
-- before player homebrew (scoped), then the lowest source priority.
CREATE VIEW compendium_top AS
SELECT id FROM (
    SELECT e.id,
        ROW_NUMBER() OVER (PARTITION BY e.kind, e.slug ORDER BY e.scoped, COALESCE(s.priority, 300), e.id) AS rn
    FROM compendium_entries AS e
    LEFT JOIN compendium_sources AS s ON s.source_key = e.source_key
)
WHERE rn = 1;
