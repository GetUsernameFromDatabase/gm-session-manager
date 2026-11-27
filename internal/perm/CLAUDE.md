# internal/perm

Every permission and visibility rule lives here; `perm_test.go` mirrors these tables (table-driven).

| Thing | GM | Player |
|---|---|---|
| Own character | n/a | full edit (via commands) |
| Other characters | view; edit HP, conditions and sources | none |
| Monster HP in combat | always | hidden unless the GM toggles `hp_visible` |
| Entities | all fields | revealed entities, revealed fields only |
| Effect presets | only the GM's own (owner none) | only their own character's; never another's, never the GM's |
| Add effect | any character, any target | own character; an ally in the same campaign only held by their own concentration; never another campaign's or a non-PC combatant |
| Remove effect | any | on their own character; on an ally's only the effects they applied (`applied_by`) |
| Take a character home (zip) | any character | own character (`CanTakeHome`); the file holds every block the character can read, `owner` blocks included, whoever asks |

Note blocks:

| Visibility | Who reads | Who edits |
|---|---|---|
| `gm` | GM | GM |
| `players` | GM and every player in the campaign | GM and the note's owner; other players only if `players_can_edit` |
| `listed` | GM and the characters in its audience (empty audience = GM only: fail closed) | GM only |
| `owner` | the owning player only; GM and admin never | the owning player |

At home (`Viewer.AtHome`, the player app) only `owner` blocks of the player's own notes can be edited, deleted or created; `players`, `listed` and `gm` blocks are read-only, even ones the player wrote. New blocks are always `owner`. A note holding any non-`owner` block cannot be renamed, deleted or reordered (`LocksShared`); adding a private block to it stays allowed.

Any player may add a `players` block to the GM's `game_session` notes. A block's author may always edit it, and may delete it or change its access while they can still read it. Reordering stays with the note's manager.
