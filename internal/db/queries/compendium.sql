-- name: GetCompendiumSource :one
SELECT * FROM compendium_sources
WHERE source_key = ?;

-- A new source gets a default priority (lower wins; migration 0004 backfills
-- the same way); a re-import keeps the GM's order but refreshes the date.
-- name: UpsertCompendiumSource :exec
INSERT INTO compendium_sources (source_key, path, hash, edition, license, attribution, published, priority)
VALUES (sqlc.arg(source_key), sqlc.arg(path), sqlc.arg(hash), sqlc.arg(edition), sqlc.arg(license), sqlc.arg(attribution), sqlc.arg(published),
    CASE WHEN sqlc.arg(source_key) = 'homebrew' THEN 0
         WHEN sqlc.arg(source_key) LIKE 'srd%' THEN 200
         WHEN sqlc.arg(edition) = '2024' THEN 100
         ELSE 300 END)
ON CONFLICT (source_key) DO UPDATE
SET path = excluded.path, hash = excluded.hash, edition = excluded.edition,
    license = excluded.license, attribution = excluded.attribution,
    published = excluded.published, last_imported = CURRENT_TIMESTAMP;

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

-- Player homebrew never shadows an imported or GM entry of the same slug;
-- among the rest the source with the lowest priority wins, then the newest.
-- name: GetCompendiumEntryBySlug :one
SELECT e.* FROM compendium_entries AS e
LEFT JOIN compendium_sources AS s ON s.source_key = e.source_key
WHERE e.kind = ? AND e.slug = ?
ORDER BY e.scoped, COALESCE(s.priority, 300), COALESCE(s.published, '') DESC, e.id
LIMIT 1;

-- GetCompendiumEntryBySlugSource finds the copy of a slug from one source, as a sheet pins it.
-- name: GetCompendiumEntryBySlugSource :one
SELECT * FROM compendium_entries
WHERE kind = ? AND slug = ? AND source_key = ?
ORDER BY scoped, id
LIMIT 1;

-- ListEntrySourcesBySlugs lists every copy of the slugs a sheet uses, best
-- first per slug (the first row is the copy priority picks); player homebrew
-- not linked to the campaign is skipped.
-- name: ListEntrySourcesBySlugs :many
SELECT e.id, e.slug, e.source_key, COALESCE(s.edition, '') AS edition
FROM compendium_entries AS e
LEFT JOIN compendium_sources AS s ON s.source_key = e.source_key
WHERE e.kind = sqlc.arg(kind) AND e.slug IN (sqlc.slice(slugs))
  AND (e.scoped = 0 OR EXISTS (
    SELECT 1 FROM compendium_campaigns AS cc
    WHERE cc.entry_id = e.id AND cc.campaign_id = sqlc.arg(campaign_id)))
ORDER BY e.slug, e.scoped, COALESCE(s.priority, 300), COALESCE(s.published, '') DESC, e.id;

-- GetCompendiumEntryByName finds an entry by name in any case, as an
-- inventory names its items.
-- name: GetCompendiumEntryByName :one
SELECT e.* FROM compendium_entries AS e
LEFT JOIN compendium_sources AS s ON s.source_key = e.source_key
WHERE e.kind = ? AND lower(e.name) = lower(sqlc.arg(name))
ORDER BY e.scoped, COALESCE(s.priority, 300), COALESCE(s.published, '') DESC, e.id
LIMIT 1;

-- ListEntryCopies is every source holding a (kind, slug), best first, for the
-- entry page's "Also in" list.
-- name: ListEntryCopies :many
SELECT e.id, e.source_key, s.edition, s.attribution, s.priority
FROM compendium_entries AS e
JOIN compendium_sources AS s ON s.source_key = e.source_key
WHERE e.kind = ? AND e.slug = ?
ORDER BY e.scoped, s.priority, s.published DESC, e.id;

-- name: ListCompendiumSources :many
SELECT * FROM compendium_sources
ORDER BY priority, published DESC, source_key;

-- name: SetCompendiumSourcePriority :exec
UPDATE compendium_sources SET priority = ?
WHERE source_key = ?;

-- Removing a book deletes its entries first (the FTS and habitat rows go
-- with them; combatants and entities that referenced one keep their row with
-- the link cleared). Homebrew is not a book.
-- name: DeleteCompendiumEntriesBySource :exec
DELETE FROM compendium_entries
WHERE source_key = ? AND source_key <> 'homebrew';

-- name: DeleteCompendiumSource :exec
DELETE FROM compendium_sources
WHERE source_key = ? AND source_key <> 'homebrew';

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
WHERE e.id IN (SELECT id FROM compendium_top)
  AND (CAST(sqlc.arg(kind) AS TEXT) = '' OR instr(',' || CAST(sqlc.arg(kind) AS TEXT) || ',', ',' || e.kind || ',') > 0)
  AND (e.name > CAST(sqlc.arg(after_name) AS TEXT) OR (e.name = CAST(sqlc.arg(after_name) AS TEXT) AND e.id > sqlc.arg(after_id)))
  AND (CAST(sqlc.arg(habitat) AS TEXT) = '' OR EXISTS (SELECT 1 FROM monster_habitats AS h WHERE h.entry_id = e.id AND instr(',' || CAST(sqlc.arg(habitat) AS TEXT) || ',', ',' || h.habitat || ',') > 0))
  AND (e.scoped = 0 OR CAST(sqlc.arg(campaign_id) AS INTEGER) = 0
       OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = CAST(sqlc.arg(campaign_id) AS INTEGER)))
ORDER BY e.name, e.id
LIMIT sqlc.arg(page_size);

-- name: RandomCompendiumEntry :one
SELECT e.id FROM compendium_entries AS e
WHERE e.id IN (SELECT id FROM compendium_top)
  AND (CAST(sqlc.arg(kind) AS TEXT) = '' OR instr(',' || CAST(sqlc.arg(kind) AS TEXT) || ',', ',' || e.kind || ',') > 0)
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
  AND e.id IN (SELECT id FROM compendium_top)
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

-- ListSpellsForClass lists the spells a player in campaign_id may see (0: all).
-- name: ListSpellsForClass :many
SELECT e.id, e.slug, e.name, e.source_key, CAST(json_extract(e.data_json, '$.level') AS INTEGER) AS level
FROM compendium_entries AS e
WHERE e.kind = 'spell' AND e.id IN (SELECT id FROM compendium_top)
  AND CAST(json_extract(e.data_json, '$.level') AS INTEGER) <= CAST(sqlc.arg(max_level) AS INTEGER)
  AND EXISTS (SELECT 1 FROM json_each(e.data_json, '$.classes') AS c WHERE c.value = CAST(sqlc.arg(class_slug) AS TEXT))
  AND (e.scoped = 0 OR CAST(sqlc.arg(campaign_id) AS INTEGER) = 0
       OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = CAST(sqlc.arg(campaign_id) AS INTEGER)))
ORDER BY level, e.name;

-- name: ListItemsByCategory :many
SELECT e.name, CAST(COALESCE(json_extract(e.data_json, '$.desc'), '') AS TEXT) AS description
FROM compendium_entries AS e
WHERE e.kind = 'item' AND e.id IN (SELECT id FROM compendium_top)
  AND json_extract(e.data_json, '$.category') = CAST(sqlc.arg(category) AS TEXT)
  AND (e.scoped = 0 OR CAST(sqlc.arg(campaign_id) AS INTEGER) = 0
       OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = CAST(sqlc.arg(campaign_id) AS INTEGER)))
ORDER BY e.name;

-- ListMasteryWeapons are the weapons with a mastery property a player in
-- campaign_id may see (0: all), for choosing Weapon Mastery. Only mundane
-- weapons: a kind is chosen, and magic variants (rarity set) are not kinds.
-- name: ListMasteryWeapons :many
SELECT e.name, CAST(json_extract(e.data_json, '$.mastery') AS TEXT) AS mastery
FROM compendium_entries AS e
WHERE e.kind = 'item' AND e.id IN (SELECT id FROM compendium_top)
  AND COALESCE(json_extract(e.data_json, '$.mastery'), '') != ''
  AND COALESCE(json_extract(e.data_json, '$.rarity'), '') = ''
  AND (e.scoped = 0 OR CAST(sqlc.arg(campaign_id) AS INTEGER) = 0
       OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = CAST(sqlc.arg(campaign_id) AS INTEGER)))
ORDER BY e.name;

-- name: ListSubclassesOf :many
SELECT id, slug, name FROM compendium_entries
WHERE kind = 'subclass' AND id IN (SELECT id FROM compendium_top) AND json_extract(data_json, '$.subclass_of') = CAST(sqlc.arg(class_slug) AS TEXT)
ORDER BY name;

-- ListCompendiumBySlugs reads the entries a sheet links to, in one query: the
-- copy compendium_top would pick, found through the (kind, slug) index
-- instead of ranking the whole compendium on every sheet render.
-- name: ListCompendiumBySlugs :many
SELECT e.id, e.slug, e.name, e.data_json, e.source_key, e.scoped FROM compendium_entries AS e
LEFT JOIN compendium_sources AS s ON s.source_key = e.source_key
WHERE e.kind = sqlc.arg(kind) AND e.slug IN (sqlc.slice(slugs))
  AND NOT EXISTS (
    SELECT 1 FROM compendium_entries AS b
    LEFT JOIN compendium_sources AS bs ON bs.source_key = b.source_key
    WHERE b.kind = e.kind AND b.slug = e.slug
      -- The dates swap sides: a newer b ranks before e (published DESC).
      AND (b.scoped, COALESCE(bs.priority, 300), COALESCE(s.published, ''), b.id)
        < (e.scoped, COALESCE(s.priority, 300), COALESCE(bs.published, ''), e.id))
ORDER BY e.slug;

-- ListCompendiumNames is what pickers offer: species versions are left out,
-- since picking the base and then its lineage grants more. Only entries a
-- player in campaign_id may see (0: all).
-- name: ListCompendiumNames :many
SELECT e.id, e.kind, e.slug, e.name FROM compendium_entries AS e
WHERE e.kind = sqlc.arg(kind) AND e.id IN (SELECT id FROM compendium_top)
  AND json_extract(e.data_json, '$.version_of') IS NULL
  AND (e.scoped = 0 OR CAST(sqlc.arg(campaign_id) AS INTEGER) = 0
       OR EXISTS (SELECT 1 FROM compendium_campaigns AS l WHERE l.entry_id = e.id AND l.campaign_id = CAST(sqlc.arg(campaign_id) AS INTEGER)))
ORDER BY e.name;

-- BindMonsterImages guesses a link for a data folder's monsters that never
-- had one; a link the GM set or cleared ('') is kept.
-- name: BindMonsterImages :exec
UPDATE compendium_entries SET image_url = sqlc.arg(prefix) || slug || '.png'
WHERE kind = 'monster' AND image_url IS NULL AND source_key = sqlc.arg(source_key);

-- BindMonsterImage sets the picture a data source names for one of its
-- monsters; a link the GM set or cleared ('') is kept.
-- name: BindMonsterImage :exec
UPDATE compendium_entries SET image_url = sqlc.arg(image_url)
WHERE kind = 'monster' AND image_url IS NULL AND source_key = sqlc.arg(source_key) AND slug = sqlc.arg(slug);

-- name: SetCompendiumImage :execrows
UPDATE compendium_entries SET image_url = ?
WHERE id = ? AND kind = 'monster';

-- The habitat grid filters monsters by name only (not body text); instr
-- needs no escaping of % and _.
-- name: ListMonstersByName :many
SELECT id, name FROM compendium_entries
WHERE kind = 'monster' AND id IN (SELECT id FROM compendium_top)
  AND instr(lower(name), lower(CAST(sqlc.arg(filter) AS TEXT))) > 0
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
