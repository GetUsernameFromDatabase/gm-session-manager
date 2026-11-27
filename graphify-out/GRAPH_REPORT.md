# Graph Report - gm-session-manager  (2026-11-28)

## Corpus Check
- 412 files · ~748,119 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 60 file(s) not represented in the graph (top: .templ 51, .css 4, (none) 3)

## Summary
- 5598 nodes · 30008 edges · 204 communities (135 shown, 69 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2973 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `06ff2f7f`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Search
- context.Context
- takehome_test.go
- createProvider
- export/markdown.go
- SheetView
- combobox.js
- .q
- Import
- encoding/json.RawMessage
- select.js
- Queries
- popover_templ.go
- Session
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- run
- Item
- wikilink.js
- HomeFileReview
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- run
- CombatantView
- TestBringBackAccept
- app.js
- tabbable.js
- useDismiss
- planSwitch
- Encounter
- command.js
- dialog.js
- parseHP
- ParseWikidot
- tooltip.js
- Input
- tabs_templ.go
- IfElse
- Open
- composite.js
- ParseArmorTraining
- CustomMod
- vendorjs/main.go
- Button
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- select_templ.go
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- DownloadFiveTools
- .attachmentsDir
- radiogroup.js
- Player
- .Router
- Sheet
- sqlc Config
- App
- handoff
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- condMods
- checkbox.js
- ParseAidedd
- switch.js
- starterEnv
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- mirror
- useTypeahead
- .source
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- Sheet
- mounted
- newPlayerEnv
- nullID
- profFocus
- cert_test.go
- ImagePage
- ftBook
- CN
- BlockEditor
- State
- WeaponKind
- redirect.go
- buildReview
- vars
- serveSplit
- .sendFeedback
- looseString
- errBodyWriter
- Player
- task
- Before/after pairs
- env
- jsconfig.json
- open5e.go
- testing.M
- SearchReadableBlocksRow
- .ListBlockStepsSince
- Client
- ASIError
- .ListItemsByCategory
- .Entry
- sourceUsedError
- run.sh
- reqLog
- Merge
- Topic
- Pending
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- .ListSpellsForClass
- KEY
- MAX_STEP_BYTES
- MAX_STEPS
- MSG
- ROUTES
- SAFE_SCHEMES
- WIKILINK
- countWriter
- IMPORT_FILTER_ALL
- IMPORT_FILTER_CATEGORY
- IMPORT_FILTER_NONE
- IMPORT_FILTER_STATUS
- IMPORT_ROWS_FIELD
- IMPORT_STATUS_ANY
- IMPORT_STATUS_BROKEN
- IMPORT_STATUS_TICKED
- IMPORT_STATUS_UNTICKED
- .ListSubracesOf
- pause_windows.go
- fakeHandler
- Browser check
- dev:changed
- release.sh
- TestPagesRender
- ListSharedWithRow
- IMPORT_FILTER_CAT_PREFIX
- lan.go
- names
- Book
- dev:build
- BuilderPickItems
- build.sh
- before
- generate:web
- release
- dev:generate
- lint
- experiments
- test
- capReader
- tasks
- github.com/a-h/templ.Component
- lint.sh
- .ListImageLookAccess

## God Nodes (most connected - your core abstractions)
1. `App` - 453 edges
2. `Path()` - 341 edges
3. `newTestEnv()` - 198 edges
4. `Button()` - 144 edges
5. `Sheet` - 141 edges
6. `CN()` - 114 edges
7. `Queries` - 89 edges
8. `SheetView` - 80 edges
9. `IfElse()` - 67 edges
10. `ParseSheet()` - 64 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `internal/compendium` --references--> `storeSource()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (204 total, 69 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (58): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+50 more)

### Community 1 - "Search"
Cohesion: 0.09
Nodes (20): HomebrewOwner, Site, Source, sourceFormat, TestHomebrew(), fiveToolsPage(), ftsQuery(), HomebrewInput (+12 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (171): keepAwake(), bindFiveToolsImages(), ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment (+163 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.07
Nodes (37): writeFile(), TakeHomeFile, readLimited(), ReadTakeHome(), rewriteZip(), sampleBringBack(), sampleFallback(), sampleNewCharacter() (+29 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.14
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (130): UsesLabel(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+122 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - ".q"
Cohesion: 0.08
Nodes (29): testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+21 more)

### Community 9 - "Import"
Cohesion: 0.11
Nodes (35): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+27 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (51): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+43 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Queries"
Cohesion: 0.03
Nodes (143): linkTarget, HomeManifest, TakeHomeBaseCharacter, TakeHomeCampaign, TakeHomeFallback, TakeHomeLevelUp, TakeHomeManifest, linkTargets() (+135 more)

### Community 13 - "popover_templ.go"
Cohesion: 0.12
Nodes (26): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+18 more)

### Community 14 - "Session"
Cohesion: 0.07
Nodes (29): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+21 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.04
Nodes (71): ListCharacterNamesRow, NoteBlock, GetBlockRow, Visibility, internal/perm, Fields, ParseFields(), VisibleEntity() (+63 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "run"
Cohesion: 0.22
Nodes (12): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), Player(), lookupMap() (+4 more)

### Community 21 - "Item"
Cohesion: 0.19
Nodes (9): findItem(), parseArmor(), parseItem(), ArmorEffect(), TestGearCommands(), ensureItem(), Item, WornArmor (+1 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - "HomeFileReview"
Cohesion: 0.10
Nodes (26): campaignLink(), CampaignRoster(), deviceRow(), PendingRequests(), rosterCard(), devicesWhere(), PendingGroup, Roster (+18 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (161): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+153 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "run"
Cohesion: 0.15
Nodes (10): loopbackHostOnly(), main(), parseFlags(), run(), TestLoopbackHostOnly(), TestParseFlags(), openBrowser(), bookFolder() (+2 more)

### Community 28 - "CombatantView"
Cohesion: 0.09
Nodes (39): Defenses, damageChoice(), capitalized(), beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog() (+31 more)

### Community 29 - "TestBringBackAccept"
Cohesion: 0.22
Nodes (15): TakeHome, TestNewCharacterTableSources(), testEnv, homeSheet(), TestBringBackAccept(), TestBringBackPlayerUpload(), TestBringBackReject(), TestBringBackSources() (+7 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "planSwitch"
Cohesion: 0.15
Nodes (8): featureSwitch, parseFeature(), parseFeatureSet(), parseUses(), featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 34 - "Encounter"
Cohesion: 0.20
Nodes (8): CanAddSources(), notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "parseHP"
Cohesion: 0.19
Nodes (11): Amount, userText, parseCoin(), CritDamage(), Damage(), parseDamage(), parseHP(), parseRevive() (+3 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.14
Nodes (23): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Input"
Cohesion: 0.05
Nodes (93): ImageCategory, Coins, Uses, Field, UI conventions, Combobox(), Content(), Item() (+85 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.08
Nodes (44): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+36 more)

### Community 42 - "IfElse"
Cohesion: 0.06
Nodes (68): LevelUpPlan, SpellColumn(), Label(), initialValue(), Item(), RadioGroup(), state(), withSpells() (+60 more)

### Community 43 - "Open"
Cohesion: 0.11
Nodes (26): migration, internal/db, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp() (+18 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (17): FeatureEffects, findMod(), newMod(), parseMod(), CustomMod, IsModTarget(), indexAbilityName(), SpellMods() (+9 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 48 - "Button"
Cohesion: 0.03
Nodes (123): Props, Variant, Props, CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow (+115 more)

### Community 49 - "Path"
Cohesion: 0.03
Nodes (138): TestUpgradeSheet(), ParseSheet(), TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight() (+130 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (159): EquipmentItem, Defense, TestDecode(), EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), TestSkillChoice(), feature() (+151 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "select_templ.go"
Cohesion: 0.11
Nodes (23): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+15 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "DownloadFiveTools"
Cohesion: 0.14
Nodes (23): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestImportFiveToolsBook() (+15 more)

### Community 59 - ".attachmentsDir"
Cohesion: 0.13
Nodes (9): firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify(), TestFolderName(), parseCSV(), saveSniffed() (+1 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.10
Nodes (30): Player(), TestCanAddSources(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState() (+22 more)

### Community 62 - ".Router"
Cohesion: 0.33
Nodes (3): RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 63 - "Sheet"
Cohesion: 0.08
Nodes (15): Part, Skill, ItemData, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+7 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "App"
Cohesion: 0.03
Nodes (42): FromContext(), Campaign, Entity, compendiumBack(), compendiumPath(), ImportSummary(), pathCampaign(), viewerCampaign() (+34 more)

### Community 66 - "handoff"
Cohesion: 0.28
Nodes (4): closeServers(), removeFile(), shutdownServers(), handoff

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.08
Nodes (45): CondOpKind, Effects, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn(), parseInspiration() (+37 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 72 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "starterEnv"
Cohesion: 0.14
Nodes (20): replaceDir(), TestFiveToolsBooks(), DownloadOpen5e(), serveOpen5e(), TestDownloadOpen5e(), TestDownloadOpen5eBadName(), TestDownloadOpen5eFailureKeepsOld(), TestDownloadOpen5eUnknownDoc() (+12 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "ParseWiki"
Cohesion: 0.22
Nodes (10): wikiLink(), wikiText(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks(), WikiLinks() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "notes.js"
Cohesion: 0.11
Nodes (8): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), randomID(), sync()

### Community 86 - "mirror"
Cohesion: 0.23
Nodes (9): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), filterSource(), TestMirrorGetMissingFile(), open5eFiles() (+1 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".source"
Cohesion: 0.12
Nodes (17): ParsePageURL(), TestParsePageURL(), UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove() (+9 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "Sheet"
Cohesion: 0.02
Nodes (142): LevelUpChoice, internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice() (+134 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.07
Nodes (43): eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack(), TestHandoffExpiresAndReplaces() (+35 more)

### Community 103 - "nullID"
Cohesion: 0.10
Nodes (37): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead() (+29 more)

### Community 105 - "cert_test.go"
Cohesion: 0.15
Nodes (10): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+2 more)

### Community 106 - "ImagePage"
Cohesion: 0.13
Nodes (19): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListBacklinksRow, attrString(), generateSVG(), getIconContent(), Props (+11 more)

### Community 107 - "ftBook"
Cohesion: 0.13
Nodes (24): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftCopyTargets(), ftDecode() (+16 more)

### Community 108 - "CN"
Cohesion: 0.06
Nodes (79): MediaProps, MediaVariant, Entity, Look, Content(), Description(), Empty(), Header() (+71 more)

### Community 111 - "State"
Cohesion: 0.20
Nodes (7): patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff()

### Community 112 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 113 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 117 - "vars"
Cohesion: 0.22
Nodes (8): env, CGO_ENABLED, GOFLAGS, vars, APP, EXE, PKG, version

### Community 118 - "serveSplit"
Cohesion: 0.50
Nodes (3): serveSplit(), TestSplitTLS(), TestSplitTLSSilentConn()

### Community 119 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Player"
Cohesion: 0.21
Nodes (21): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+13 more)

### Community 127 - "task"
Cohesion: 0.22
Nodes (11): cmds, desc, env, cmds, desc, env, task, APP (+3 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "env"
Cohesion: 0.11
Nodes (26): BIN, DATA_DIR, MARKERS, MODE, NOREBUILD, PORT, PREPARE, RESTART_ARGS (+18 more)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 137 - "Client"
Cohesion: 0.16
Nodes (6): ClientInfo, goneDevice, message, Client, Hub, isDevice()

### Community 140 - ".Entry"
Cohesion: 0.08
Nodes (29): MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, TestTokensReadAsLabels(), TestFiveToolsCoreTraitsShapes() (+21 more)

### Community 143 - "run.sh"
Cohesion: 0.56
Nodes (8): alive(), lock(), run.sh script, start_server(), stop_server(), superseded(), unchanged(), unlock()

### Community 145 - "Merge"
Cohesion: 0.15
Nodes (19): sheetMap, addedRemoved(), describe(), diffNames(), finish(), GroupDiff, GroupKeys(), GroupLabel() (+11 more)

### Community 147 - "Topic"
Cohesion: 0.12
Nodes (7): event, Handler, Topic, TopicKind, noteTopics(), fiveToolsJob, heldTopics

### Community 148 - "Pending"
Cohesion: 0.15
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 179 - "dev:changed"
Cohesion: 0.40
Nodes (5): cmds, method, sources, exclude, dev:changed

### Community 180 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 181 - "TestPagesRender"
Cohesion: 0.15
Nodes (10): Gotchas, Handlers, internal/web, TestKeepsBodyNewCharacter(), TestKeepsBodyBringBack(), TestPagesRender(), keepsBody(), fieldError (+2 more)

### Community 184 - "lan.go"
Cohesion: 0.47
Nodes (5): LANAddresses(), PreferDefaultRoute(), preferFirst(), TestPreferFirst(), handoffIPs()

### Community 185 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 186 - "Book"
Cohesion: 0.11
Nodes (26): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+18 more)

### Community 188 - "dev:build"
Cohesion: 0.40
Nodes (5): cmds, deps, requires, vars, dev:build

### Community 189 - "BuilderPickItems"
Cohesion: 0.50
Nodes (5): BuilderPickItems(), BuilderPick, pickURL(), addedAtHomeIcon(), SearchComboItems()

### Community 191 - "before"
Cohesion: 0.40
Nodes (5): cmds, desc, vars, before, BEFORE_BIN

### Community 192 - "generate:web"
Cohesion: 0.67
Nodes (3): cmds, desc, generate:web

### Community 193 - "release"
Cohesion: 0.67
Nodes (3): cmds, desc, release

### Community 194 - "dev:generate"
Cohesion: 0.29
Nodes (7): cmd, if, cmds, generates, internal, sources, dev:generate

### Community 195 - "lint"
Cohesion: 0.50
Nodes (4): cmds, desc, silent, lint

### Community 197 - "test"
Cohesion: 0.67
Nodes (3): test, cmds, desc

### Community 200 - "tasks"
Cohesion: 0.12
Nodes (16): cmds, cmds, method, sources, cmds, desc, cmds, internal (+8 more)

### Community 201 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (89): ctxKey, ctxState, Props, Alert(), Description(), Title(), Dialog(), Close() (+81 more)

## Knowledge Gaps
- **37 isolated node(s):** `npx`, `@playwright/mcp`, `gm-session-manager`, `testEnv`, `testEnv` (+32 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 534 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **69 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `Path` to `compendium_templ.go`, `context.Context`, `takehome_test.go`, `export/markdown.go`, `SheetView`, `.q`, `Import`, `Queries`, `popover_templ.go`, `perm.go`, `HomeFileReview`, `github.com/a-h/templ.Attributes`, `CombatantView`, `TestBringBackAccept`, `Input`, `tabs_templ.go`, `IfElse`, `Button`, `testing.T`, `TestPagesRender`, `Player`, `BuilderPickItems`, `App`, `github.com/a-h/templ.Component`, `starterEnv`, `ParseWiki`, `.source`, `Sheet`, `newPlayerEnv`, `nullID`, `ImagePage`, `CN`, `.sendFeedback`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `gm-session-manager` to the rest of the system?**
  _37 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14867663981588033 - nodes in this community are weakly interconnected._
- **Why does `App` connect `App` to `compendium_templ.go`, `Search`, `context.Context`, `.q`, `Client`, `Queries`, `Session`, `perm.go`, `reqLog`, `Topic`, `Pending`, `Encounter`, `testing.T`, `Book`, `.attachmentsDir`, `.Router`, `handoff`, `Command`, `Sheet`, `cert_test.go`, `State`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **Should `Search` be split into smaller, more focused modules?**
  _Cohesion score 0.08620689655172414 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `TestPagesRender`, `wikilink.js`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._