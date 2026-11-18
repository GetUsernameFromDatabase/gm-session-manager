-- Conditions

-- name: ListConditionsByCharacter :many
SELECT * FROM conditions
WHERE character_id = ?
ORDER BY slug, id;

-- name: ListConditionsByEncounter :many
-- Includes the PCs' own conditions (by character), which the tracker shows too.
SELECT cond.* FROM conditions AS cond
JOIN combatants AS cb ON cb.id = cond.combatant_id OR cb.character_id = cond.character_id
WHERE cb.encounter_id = ?
GROUP BY cond.id
ORDER BY cond.slug, cond.id;

-- name: CreateCondition :one
INSERT INTO conditions (character_id, combatant_id, slug, note, expires_round, ends_on, held_by, mods_json, applied_by)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: ListEffectPresets :many
-- One owner's presets (a NULL owner is the GM's) in a campaign.
SELECT * FROM effect_presets
WHERE campaign_id = ? AND owner_character_id IS sqlc.narg(owner)
ORDER BY name COLLATE NOCASE, id;

-- name: GetEffectPreset :one
SELECT * FROM effect_presets WHERE id = ?;

-- name: UpsertEffectPreset :one
-- A same-name save by the same owner overwrites.
INSERT INTO effect_presets (campaign_id, owner_character_id, name, slug, note, ends_on, rounds, mods_json)
VALUES (sqlc.arg(campaign_id), sqlc.narg(owner), sqlc.arg(name), sqlc.arg(slug), sqlc.arg(note), sqlc.arg(ends_on), sqlc.arg(rounds), sqlc.arg(mods_json))
ON CONFLICT (campaign_id, COALESCE(owner_character_id, 0), name) DO UPDATE SET
    slug = excluded.slug, note = excluded.note, ends_on = excluded.ends_on,
    rounds = excluded.rounds, mods_json = excluded.mods_json
RETURNING *;

-- name: DeleteEffectPreset :execrows
DELETE FROM effect_presets WHERE id = ? AND campaign_id = ? AND owner_character_id IS sqlc.narg(owner);

-- name: ListCombatantConditions :many
SELECT * FROM conditions
WHERE combatant_id = ?
ORDER BY slug, id;

-- ListHeldBy are the owners of the effects hanging off the given
-- concentration rows (and, for a combatant, its encounter): who to refresh
-- before those rows are deleted.
-- name: ListHeldBy :many
SELECT cond.character_id, cond.combatant_id, cb.encounter_id
FROM conditions AS cond
LEFT JOIN combatants AS cb ON cb.id = cond.combatant_id
WHERE cond.held_by IN (sqlc.slice(ids));

-- ListRestConditions are an owner's effects that end on this rest: the
-- short ones, and on a long rest the long ones too.
-- name: ListRestConditions :many
SELECT * FROM conditions
WHERE (character_id = sqlc.narg(character_id) OR combatant_id = sqlc.narg(combatant_id))
  AND ends_on IN (sqlc.arg(first), sqlc.arg(second));

-- ListCharacterConcentration and ListEncounterConcentration are the
-- concentration rows an effect may hang off, with whose they are.
-- name: ListCharacterConcentration :many
SELECT cond.id, cond.note, ch.id AS character_id, ch.name AS owner_name
FROM conditions AS cond
JOIN characters AS ch ON ch.id = cond.character_id
WHERE ch.campaign_id = ? AND ch.deleted_at IS NULL AND cond.slug = 'concentration'
ORDER BY ch.name, cond.id;

-- name: ListEncounterConcentration :many
SELECT cond.id, cond.note, cb.id AS combatant_id, cb.name AS owner_name
FROM conditions AS cond
JOIN combatants AS cb ON cb.id = cond.combatant_id
WHERE cb.encounter_id = ? AND cond.slug = 'concentration'
ORDER BY cb.name, cond.id;

-- DeleteRoundConditions ends the round-timed effects of a fight's combatants
-- and their characters, as the fight ends.
-- name: DeleteRoundConditions :exec
DELETE FROM conditions
WHERE expires_round IS NOT NULL
  AND (
    combatant_id IN (SELECT cb.id FROM combatants AS cb WHERE cb.encounter_id = sqlc.arg(encounter_id))
    OR character_id IN (SELECT cb.character_id FROM combatants AS cb WHERE cb.encounter_id = sqlc.arg(encounter_id))
  );

-- ListExpiringConditions is what ExpireTurnConditions is about to delete,
-- so the effects held by those rows can be noted first. Keep the two WHEREs
-- the same.
-- name: ListExpiringConditions :many
SELECT id FROM conditions
WHERE expires_round IS NOT NULL AND expires_round < sqlc.arg(round)
  AND (
    combatant_id = sqlc.arg(combatant_id)
    OR character_id = (SELECT cb.character_id FROM combatants AS cb WHERE cb.id = sqlc.arg(combatant_id))
  );

-- name: GetCondition :one
SELECT * FROM conditions
WHERE id = ?;

-- name: FindCharacterCondition :one
SELECT * FROM conditions
WHERE character_id = ? AND slug = ?
LIMIT 1;

-- name: FindCombatantCondition :one
SELECT * FROM conditions
WHERE combatant_id = ? AND slug = ?
LIMIT 1;

-- name: DeleteCondition :exec
DELETE FROM conditions
WHERE id = ?;

-- ExpireTurnConditions removes the timed effects on a combatant (or on its
-- character) whose last round has passed, as its turn starts.
-- name: ExpireTurnConditions :exec
DELETE FROM conditions
WHERE expires_round IS NOT NULL AND expires_round < sqlc.arg(round)
  AND (
    combatant_id = sqlc.arg(combatant_id)
    OR character_id = (SELECT cb.character_id FROM combatants AS cb WHERE cb.id = sqlc.arg(combatant_id))
  );

-- Encounters

-- name: CreateEncounter :one
INSERT INTO encounters (campaign_id, name)
VALUES (?, ?)
RETURNING *;

-- name: GetEncounter :one
SELECT * FROM encounters
WHERE id = ?;

-- name: ListEncountersByCampaign :many
SELECT * FROM encounters
WHERE campaign_id = ?
ORDER BY CASE status WHEN 'active' THEN 0 WHEN 'preparing' THEN 1 ELSE 2 END, created_at DESC;

-- name: ListCharacterEncounters :many
SELECT e.* FROM encounters AS e
WHERE e.campaign_id = ? AND e.status IN ('active', 'ended')
  AND EXISTS (SELECT 1 FROM combatants AS cb WHERE cb.encounter_id = e.id AND cb.character_id = ?)
ORDER BY CASE e.status WHEN 'active' THEN 0 ELSE 1 END, e.created_at DESC;

-- name: GetActiveEncounter :one
SELECT * FROM encounters
WHERE campaign_id = ? AND status = 'active'
ORDER BY created_at DESC
LIMIT 1;

-- name: UpdateEncounterState :exec
UPDATE encounters
SET status = ?, round = ?, active_combatant_id = ?
WHERE id = ?;

-- name: DeleteEncounter :exec
DELETE FROM encounters
WHERE id = ?;

-- Combatants

-- name: CreateCombatant :one
INSERT INTO combatants (
    encounter_id, character_id, entity_id, monster_ref, name, initiative,
    hp, max_hp, ac, hp_visible, sort_order
) VALUES (
    sqlc.arg(encounter_id), sqlc.narg(character_id), sqlc.narg(entity_id), sqlc.narg(monster_ref),
    sqlc.arg(name), sqlc.narg(initiative), sqlc.arg(hp), sqlc.arg(max_hp), sqlc.arg(ac), sqlc.arg(hp_visible),
    (SELECT coalesce(max(o.sort_order), 0) + 1 FROM combatants AS o WHERE o.encounter_id = sqlc.arg(encounter_id))
)
RETURNING *;

-- name: GetCombatant :one
SELECT * FROM combatants
WHERE id = ?;

-- ListCombatants returns the turn order. PC rows carry the character's live
-- name and HP, NPC rows their entity's name, so renames show up in fights.
-- name: ListCombatants :many
SELECT cb.id, cb.encounter_id, cb.character_id, cb.entity_id, cb.monster_ref,
       CAST(coalesce(ch.name, en.name, cb.name) AS TEXT) AS name,
       cb.initiative, cb.ac, cb.hp_visible, cb.sort_order, cb.marker,
       CAST(coalesce(ch.hp, cb.hp) AS INTEGER)           AS hp,
       CAST(coalesce(ch.max_hp, cb.max_hp) AS INTEGER)   AS max_hp,
       CAST(coalesce(ch.temp_hp, cb.temp_hp) AS INTEGER) AS temp_hp,
       m.image_url AS monster_image, m.data_json AS monster_data,
       ch.sheet_json, en.is_revealed AS entity_revealed
FROM combatants AS cb
LEFT JOIN characters AS ch ON ch.id = cb.character_id
LEFT JOIN entities AS en ON en.id = cb.entity_id
LEFT JOIN compendium_entries AS m ON m.id = cb.monster_ref
WHERE cb.encounter_id = ?
ORDER BY cb.initiative IS NULL, cb.initiative DESC, cb.sort_order;

-- name: ListEncounterIDsByCharacter :many
SELECT DISTINCT e.id FROM encounters AS e
JOIN combatants AS cb ON cb.encounter_id = e.id
WHERE cb.character_id = ? AND e.status <> 'ended';

-- ListEncountersByMonster finds the open fights that show a monster's picture.
-- name: ListEncountersByMonster :many
SELECT DISTINCT e.id, e.campaign_id FROM encounters AS e
JOIN combatants AS cb ON cb.encounter_id = e.id
WHERE cb.monster_ref = ? AND e.status <> 'ended';

-- name: SetCombatantInitiative :exec
UPDATE combatants
SET initiative = ?
WHERE id = ?;

-- name: SetCombatantHP :exec
UPDATE combatants
SET hp = ?, max_hp = ?, temp_hp = ?
WHERE id = ?;

-- name: ToggleCombatantHPVisible :exec
UPDATE combatants
SET hp_visible = NOT hp_visible
WHERE id = ?;

-- name: SetCombatantMarker :exec
UPDATE combatants
SET marker = ?
WHERE id = ?;

-- name: DeleteCharacterCombatants :exec
DELETE FROM combatants
WHERE character_id = ?;

-- name: DeleteCombatant :exec
DELETE FROM combatants
WHERE id = ?;
