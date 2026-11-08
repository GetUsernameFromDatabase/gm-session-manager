# internal/compendium

- Normalize source data into our own shape; never leak a source's shape into the app. Keep each entry's license and attribution (the 2024 SRD is CC-BY 4.0).
- Parser changes: bump `normalizeVersion` (`importer.go`); it salts the source hash of data folders (5etools, Open5e). Wikidot/aidedd pages are not hashed: re-import them by hand.
- Force a re-parse in `data-test`: `update compendium_sources set hash=''`, then `touch data-test/restart`.
- Real data: `GSM_COMPENDIUM_DIR=data/compendium go test ./internal/compendium -run TestRealSource -v`. For 5etools, a temporary `_test.go` calling `FiveToolsBooks` + `DownloadFiveTools` into a scratchpad dir, then `Import`; ask the user to delete it (`rm` is denied).
- The trimmed 5etools fixture lives in its own folder, `testdata/fivetools/5etools-xphb/`, so the web tests' `testdata` import skips it.
