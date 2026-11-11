# internal/db

- The migrations in `migrations/` are the schema. WAL mode and `PRAGMA foreign_keys = ON` on every connection.
- All queries through `sqlc` (`queries/` → `sqlc/`, generated). FTS5 tables live in a migration, with a stub for sqlc in `sqlc_stubs.sql`.
- Prefer separate simple queries, split in Go or the view model, over one clever query with CAST/subquery flags.
- sqlc column `overrides` ignore nullability: an overridden column read through a LEFT JOIN needs `COALESCE(col, '')`, or a NULL fails to scan.
