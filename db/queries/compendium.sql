-- name: GetCompendiumSource :one
SELECT * FROM compendium_sources
WHERE source_key = ?;

-- name: UpsertCompendiumSource :exec
INSERT INTO compendium_sources (source_key, path, hash, edition, license, attribution)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT (source_key) DO UPDATE
SET path = excluded.path, hash = excluded.hash, edition = excluded.edition,
    license = excluded.license, attribution = excluded.attribution,
    last_imported = CURRENT_TIMESTAMP;

-- UpsertCompendiumEntry keeps the row id across re-imports so combatants and
-- entities that reference it stay linked. Unchanged content is not rewritten.
-- name: UpsertCompendiumEntry :exec
INSERT INTO compendium_entries (
    kind, slug, name, source_key, data_json, body_text, is_homebrew, content_hash
) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
ON CONFLICT (kind, slug, source_key) DO UPDATE
SET name = excluded.name, data_json = excluded.data_json,
    body_text = excluded.body_text, content_hash = excluded.content_hash,
    imported_at = CURRENT_TIMESTAMP
WHERE compendium_entries.content_hash <> excluded.content_hash;

-- name: ListCompendiumKeysBySource :many
SELECT id, kind, slug FROM compendium_entries
WHERE source_key = ?;

-- name: DeleteCompendiumEntry :exec
DELETE FROM compendium_entries
WHERE id = ?;

-- name: GetCompendiumEntry :one
SELECT * FROM compendium_entries
WHERE id = ?;

-- Player homebrew never shadows an imported or GM entry of the same slug.
-- name: GetCompendiumEntryBySlug :one
SELECT * FROM compendium_entries
WHERE kind = ? AND slug = ?
ORDER BY scoped, is_homebrew DESC
LIMIT 1;

-- GetCompendiumEntryByName finds an entry by name in any case, as an
-- inventory names its items.
-- name: GetCompendiumEntryByName :one
SELECT * FROM compendium_entries
WHERE kind = ? AND lower(name) = lower(sqlc.arg(name))
ORDER BY scoped, is_homebrew DESC
LIMIT 1;

-- Browse, random and search show player homebrew (scoped) only in the
-- campaigns it is linked to; campaign_id 0 is the GM, who sees it all
-- (perm.CanSeeEntry).

-- ListCompendiumByKind reads one kind, or every kind when kind is '', after
-- the (name, id) cursor; names repeat, so id breaks ties.
-- name: ListCompendiumByKind :many
SELECT e.id, e.kind, e.slug, e.name, e.source_key, e.is_homebrew,
    CAST(COALESCE((SELECT group_concat(c.value, ' ') FROM json_each(e.data_json, '$.classes') AS c), json_extract(e.data_json, '$.subclass_of'), '') AS TEXT) AS classes,
    CAST(COALESCE(json_extract(e.data_json, '$.category'), '') AS TEXT) AS category,
    CAST(COALESCE(json_extract(e.data_json, '$.size'), '') AS TEXT) AS size,
    CAST(COALESCE(json_extract(e.data_json, '$.alignment'), '') AS TEXT) AS alignment,
    CAST(COALESCE(CASE e.kind WHEN 'monster' THEN json_extract(e.data_json, '$.cr')
        WHEN 'spell' THEN json_extract(e.data_json, '$.level')
        WHEN 'item' THEN json_extract(e.data_json, '$.rarity') END, '') AS TEXT) AS figure
FROM compendium_entries AS e
WHERE (CAST(sqlc.arg(kind) AS TEXT) = '' OR instr(',' || CAST(sqlc.arg(kind) AS TEXT) || ',', ',' || e.kind || ',') > 0)
  AND (e.name > CAST(sqlc.arg(after_name) AS TEXT) OR (e.name = CAST(sqlc.arg(after_name) AS TEXT) AND e.id > sqlc.arg(after_id)))
  AND (CAST(sqlc.arg(habitat) AS TEXT) = '' OR EXISTS (SELECT 1 FROM monster_habitats AS h WHERE h.entry_id = e.id AND instr(',' || CAST(sqlc.arg(habitat) AS TEXT) || ',', ',' || h.habitat || ',') > 0))
  AND (e.scoped = 0 OR CAST(sqlc.arg(campaign_id) AS INTEGER) = 0
       OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = CAST(sqlc.arg(campaign_id) AS INTEGER)))
ORDER BY e.name, e.id
LIMIT sqlc.arg(page_size);

-- name: RandomCompendiumEntry :one
SELECT e.id FROM compendium_entries AS e
WHERE (CAST(sqlc.arg(kind) AS TEXT) = '' OR instr(',' || CAST(sqlc.arg(kind) AS TEXT) || ',', ',' || e.kind || ',') > 0)
  AND (CAST(sqlc.arg(habitat) AS TEXT) = '' OR EXISTS (SELECT 1 FROM monster_habitats AS h WHERE h.entry_id = e.id AND instr(',' || CAST(sqlc.arg(habitat) AS TEXT) || ',', ',' || h.habitat || ',') > 0))
  AND (e.scoped = 0 OR CAST(sqlc.arg(campaign_id) AS INTEGER) = 0
       OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = CAST(sqlc.arg(campaign_id) AS INTEGER)))
ORDER BY random()
LIMIT 1;

-- ListHomebrew is every homebrew entry with its author's name (empty for the
-- GM's), for the GM.
-- name: ListHomebrew :many
SELECT e.id, e.kind, e.slug, e.name, e.scoped, e.owner_character_id,
    CAST(COALESCE(c.name, '') AS TEXT) AS owner_name
FROM compendium_entries AS e
LEFT JOIN characters AS c ON c.id = e.owner_character_id
WHERE e.is_homebrew = 1
ORDER BY e.scoped, e.kind, e.name;

-- name: ListHomebrewByOwner :many
SELECT id, kind, slug, name FROM compendium_entries
WHERE is_homebrew = 1 AND owner_character_id = ?
ORDER BY kind, name;

-- name: ListHomebrewLinks :many
SELECT l.entry_id, l.campaign_id, c.name AS campaign_name
FROM compendium_campaigns AS l
JOIN campaigns AS c ON c.id = l.campaign_id
ORDER BY c.name;

-- name: ListEntryCampaigns :many
SELECT campaign_id FROM compendium_campaigns
WHERE entry_id = ?;

-- name: LinkEntryCampaign :exec
INSERT INTO compendium_campaigns (entry_id, campaign_id)
VALUES (?, ?)
ON CONFLICT DO NOTHING;

-- name: UnlinkEntryCampaign :exec
DELETE FROM compendium_campaigns
WHERE entry_id = ? AND campaign_id = ?;

-- name: SearchCompendium :many
SELECT e.id, e.kind, e.slug, e.name, e.source_key, e.is_homebrew,
    CAST(COALESCE((SELECT group_concat(c.value, ' ') FROM json_each(e.data_json, '$.classes') AS c), json_extract(e.data_json, '$.subclass_of'), '') AS TEXT) AS classes,
    CAST(COALESCE(json_extract(e.data_json, '$.category'), '') AS TEXT) AS category,
    CAST(COALESCE(json_extract(e.data_json, '$.size'), '') AS TEXT) AS size,
    CAST(COALESCE(json_extract(e.data_json, '$.alignment'), '') AS TEXT) AS alignment,
    CAST(COALESCE(CASE e.kind WHEN 'monster' THEN json_extract(e.data_json, '$.cr')
        WHEN 'spell' THEN json_extract(e.data_json, '$.level')
        WHEN 'item' THEN json_extract(e.data_json, '$.rarity') END, '') AS TEXT) AS figure
FROM compendium_fts
JOIN compendium_entries AS e ON e.id = compendium_fts.rowid
WHERE compendium_fts MATCH CAST(sqlc.arg(query) AS TEXT)
  AND (CAST(sqlc.arg(kind) AS TEXT) = '' OR instr(',' || CAST(sqlc.arg(kind) AS TEXT) || ',', ',' || e.kind || ',') > 0)
  AND (CAST(sqlc.arg(habitat) AS TEXT) = '' OR EXISTS (SELECT 1 FROM monster_habitats AS h WHERE h.entry_id = e.id AND instr(',' || CAST(sqlc.arg(habitat) AS TEXT) || ',', ',' || h.habitat || ',') > 0))
  AND (e.scoped = 0 OR CAST(sqlc.arg(campaign_id) AS INTEGER) = 0
       OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = CAST(sqlc.arg(campaign_id) AS INTEGER)))
ORDER BY compendium_fts.rank
LIMIT sqlc.arg(max_results);

-- CreateHomebrew writes a GM entry (no owner, global) or a player's (owner
-- set, scoped to the campaigns it is linked to).
-- name: CreateHomebrew :one
INSERT INTO compendium_entries (
    kind, slug, name, source_key, data_json, body_text, is_homebrew, content_hash, owner_character_id, scoped
) VALUES (?, ?, ?, 'homebrew', ?, ?, 1, ?, ?, ?)
RETURNING *;

-- name: UpdateHomebrew :exec
UPDATE compendium_entries
SET name = ?, data_json = ?, body_text = ?, content_hash = ?, imported_at = CURRENT_TIMESTAMP
WHERE id = ? AND is_homebrew = 1;

-- name: DeleteHomebrew :exec
DELETE FROM compendium_entries
WHERE id = ? AND is_homebrew = 1;

-- name: CountCompendium :many
SELECT kind, count(*) AS n FROM compendium_entries
GROUP BY kind
ORDER BY kind;

-- name: ListSpellsForClass :many
SELECT id, slug, name, CAST(json_extract(data_json, '$.level') AS INTEGER) AS level
FROM compendium_entries
WHERE kind = 'spell'
  AND CAST(json_extract(data_json, '$.level') AS INTEGER) <= CAST(sqlc.arg(max_level) AS INTEGER)
  AND EXISTS (SELECT 1 FROM json_each(compendium_entries.data_json, '$.classes') AS c WHERE c.value = CAST(sqlc.arg(class_slug) AS TEXT))
ORDER BY level, name;

-- name: ListSubclassesOf :many
SELECT id, slug, name FROM compendium_entries
WHERE kind = 'subclass' AND json_extract(data_json, '$.subclass_of') = CAST(sqlc.arg(class_slug) AS TEXT)
ORDER BY name;

-- ListCompendiumBySlugs reads the entries a sheet links to, in one query.
-- name: ListCompendiumBySlugs :many
SELECT id, slug, name, data_json FROM compendium_entries
WHERE kind = sqlc.arg(kind) AND slug IN (sqlc.slice(slugs))
ORDER BY scoped, is_homebrew DESC;

-- name: ListCompendiumNames :many
SELECT id, kind, slug, name FROM compendium_entries
WHERE kind = ?
ORDER BY name;

-- BindMonsterImages guesses a link for a data folder's monsters that never
-- had one; a link the GM set or cleared ('') is kept.
-- name: BindMonsterImages :exec
UPDATE compendium_entries SET image_url = sqlc.arg(prefix) || slug || '.png'
WHERE kind = 'monster' AND image_url IS NULL AND source_key = sqlc.arg(source_key);

-- name: SetCompendiumImage :execrows
UPDATE compendium_entries SET image_url = ?
WHERE id = ? AND kind = 'monster';

-- The habitat grid filters monsters by name only (not body text); instr
-- needs no escaping of % and _.
-- name: ListMonstersByName :many
SELECT id, name FROM compendium_entries
WHERE kind = 'monster' AND instr(lower(name), lower(CAST(sqlc.arg(filter) AS TEXT))) > 0
ORDER BY name, id;

-- name: ListMonsterHabitats :many
SELECT entry_id, habitat FROM monster_habitats;

-- name: CountHabitats :many
SELECT habitat, count(*) AS n FROM monster_habitats
GROUP BY habitat;

-- name: ListEntryHabitats :many
SELECT habitat FROM monster_habitats
WHERE entry_id = ?;

-- AddMonsterHabitat touches one row even when the habitat is already set,
-- so 0 rows means the entry is not a monster.
-- name: AddMonsterHabitat :execrows
INSERT INTO monster_habitats (entry_id, habitat)
SELECT id, CAST(sqlc.arg(habitat) AS TEXT) FROM compendium_entries
WHERE id = sqlc.arg(entry_id) AND kind = 'monster'
ON CONFLICT DO UPDATE SET habitat = excluded.habitat;

-- name: RemoveMonsterHabitat :exec
DELETE FROM monster_habitats
WHERE entry_id = ? AND habitat = ?;

-- name: AddHabitatByName :execrows
INSERT INTO monster_habitats (entry_id, habitat)
SELECT id, CAST(sqlc.arg(habitat) AS TEXT) FROM compendium_entries
WHERE kind = 'monster' AND instr(lower(name), lower(CAST(sqlc.arg(filter) AS TEXT))) > 0
ON CONFLICT DO NOTHING;

-- name: RemoveHabitatByName :execrows
DELETE FROM monster_habitats
WHERE habitat = CAST(sqlc.arg(habitat) AS TEXT) AND entry_id IN (
    SELECT id FROM compendium_entries
    WHERE kind = 'monster' AND instr(lower(name), lower(CAST(sqlc.arg(filter) AS TEXT))) > 0);
