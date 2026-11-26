# internal/db

- `migrations/0001_schema.sql` is the whole schema; a change is a new numbered file, never an edit to a released one. WAL mode and `PRAGMA foreign_keys = ON` on every connection.
- A migration that edits `sqlite_schema` (`PRAGMA writable_schema`) relies on `apply()`'s `schema_version` bump, so a later `ALTER TABLE` on the same connection sees the edited text.
- All queries through `sqlc` (`queries/` → `sqlc/`, generated). sqlc makes every comment since the previous query (blank lines and `/* */` included) the next query's Go doc: no file headers or section labels in `queries/`; a note about a group of queries goes on the first of them, worded "here and below". FTS5 tables live in a migration, with a stub for sqlc in `sqlc_stubs.sql`.
- Prefer separate simple queries, split in Go or the view model, over one query whose flags change its shape. An optional filter as `CAST(sqlc.arg(x) AS INTEGER) = 0 OR …` is fine (sqlc needs the CAST to type the arg).
- Which compendium copy a campaign sees, player homebrew scoping included, lives in the `campaign_top` view: join it rather than repeating the check in a query.
- A new CHECK enum needs its type in `internal/enum` plus a column override in `sqlc.yaml`.
- A multi-statement `ExecContext` binds each statement's `?` from the first argument again: one statement per bound value.
- sqlc column `overrides` ignore nullability: an overridden column read through a LEFT JOIN needs `COALESCE(col, '')`, or a NULL fails to scan.
