# internal/perm

Every permission and visibility rule lives here; `perm_test.go` mirrors these tables (table-driven).

| Thing | GM | Player |
|---|---|---|
| Own character | n/a | full edit (via commands) |
| Other characters | view; edit HP, conditions and sources | none |
| Monster HP in combat | always | hidden unless the GM toggles `hp_visible` |
| Entities | all fields | revealed entities, revealed fields only |

Note blocks:

| Visibility | Who reads | Who edits |
|---|---|---|
| `gm` | GM | GM |
| `players` | GM and every player in the campaign | GM and the note's owner; other players only if `players_can_edit` |
| `listed` | GM and the characters in its audience (empty audience = GM only: fail closed) | GM only |
| `owner` | the owning player only; GM and admin never | the owning player |

Any player may add a `players` block to the GM's `game_session` notes. A block's author may always edit it, and may delete it or change its access while they can still read it. Reordering stays with the note's manager.
