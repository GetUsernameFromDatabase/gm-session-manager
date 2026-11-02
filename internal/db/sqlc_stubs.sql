-- Read by sqlc only, never applied to a database (it is outside migrations/).
-- sqlc parses the FTS5 virtual table in 0001_init.sql but does not know
-- its hidden columns: the table-name MATCH operand, rowid and rank. Declaring
-- them here lets the search query stay in sqlc.
ALTER TABLE compendium_fts ADD COLUMN compendium_fts TEXT;
ALTER TABLE compendium_fts ADD COLUMN rowid INTEGER;
ALTER TABLE compendium_fts ADD COLUMN rank REAL;
