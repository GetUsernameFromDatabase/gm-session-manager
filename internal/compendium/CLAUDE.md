# internal/compendium

- This package imports `dnd5e`, so `dnd5e` never imports it back: rule code there takes plain values (`dnd5e.Cantrip`, not `SpellData`) and gets `Slugify` passed in. 5e text helpers (`SplitList`, `Squash`, `CountWords`) live in `dnd5e`.
- A caller that knows the kind's type decodes with `Decode[T]`; `DecodeData` is for code that really wants `any`.
- Imported entries use our types (`entry.go`). Text a player sees holds no source codes (`LA|XPHB`, `SCF`, `{@…}`): map codes to words or drop them. Keep each entry's license and attribution (the 2024 SRD is CC-BY 4.0).
- Parser changes: bump `normalizeVersion` (`importer.go`); it salts the source hash of data folders (5etools, Open5e). Wikidot/aidedd pages are not hashed: re-import them by hand.
- Changes to what a 5etools download writes (bundled records, rewritten links): bump `DownloadVersion` (`fivetools_books.go`). Older books then show the GM a "download again" banner, or download again at start if the GM turned that on.
- Force a re-parse in `data-test`: `update compendium_sources set hash=''`, then `touch data-test/restart`.
- Real data: `GSM_COMPENDIUM_DIR=data/compendium go test ./internal/compendium -run TestRealSource -v`. For 5etools, a temporary `_test.go` calling `FiveToolsBooks` + `DownloadFiveTools` into a scratchpad dir, then `Import`; ask the user to delete it (`rm` is denied).
- 5etools subraces are complete entries with `subrace_of` (the base's slug): the race's traits, those named in `data.overwrite` replaced in place. A race of another book that subraces use is bundled at download (`_copyTargets.race`) and imported under the add-on's source, named "Elf (2014)" only when reprinted under the same name. Pickers leave subraces out; the wizard offers them as variants.
- A subclass shipped for both the PHB and XPHB class keeps the XPHB copy. Subclass features referenced from another book (`…|XGE|8|TCE`) are bundled at download.
- Tags to 2014 records with `reprintedAs` are rewritten to the reprint at download (`bookWriter.write`), keeping their words.
- The trimmed 5etools fixture lives in its own folder, `testdata/fivetools/5etools-xphb/`, so the web tests' `testdata` import skips it.
