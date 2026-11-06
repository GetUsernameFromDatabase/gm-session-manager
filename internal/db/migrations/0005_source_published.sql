-- Inside one priority (a group, until the GM reorders), the newer source wins.
-- Sources without a date ('') come last. The importer fills the date, so
-- existing sources get it on their next import (normalizeVersion).
-- Keep the ORDER BY in step with the lookups in queries/compendium.sql.
ALTER TABLE compendium_sources ADD COLUMN published TEXT NOT NULL DEFAULT '';

DROP VIEW compendium_top;

CREATE VIEW compendium_top AS
SELECT id FROM (
    SELECT e.id,
        ROW_NUMBER() OVER (
            PARTITION BY e.kind, e.slug
            ORDER BY e.scoped, COALESCE(s.priority, 300), COALESCE(s.published, '') DESC, e.id
        ) AS rn
    FROM compendium_entries AS e
    LEFT JOIN compendium_sources AS s ON s.source_key = e.source_key
)
WHERE rn = 1;
