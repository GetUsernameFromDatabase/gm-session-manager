-- Conditions

-- name: ListConditionsByCharacter :many
SELECT * FROM conditions
WHERE character_id = ?
ORDER BY slug, id;

-- name: ListConditionsByEncounter :many
SELECT cond.* FROM conditions AS cond
JOIN combatants AS cb ON cb.id = cond.combatant_id
WHERE cb.encounter_id = ?
ORDER BY cond.slug, cond.id;

-- name: CreateCondition :one
INSERT INTO conditions (character_id, combatant_id, slug, note, expires_round)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

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
