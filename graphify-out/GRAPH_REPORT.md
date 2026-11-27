# Graph Report - gm-session-manager  (2026-11-27)

## Corpus Check
- 403 files · ~714,647 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 60 file(s) not represented in the graph (top: .templ 50, .css 4, (none) 3)

## Summary
- 5260 nodes · 27953 edges · 185 communities (121 shown, 64 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2801 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `d2e360b7`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Import
- context.Context
- takehome_test.go
- createProvider
- export/markdown.go
- Button
- combobox.js
- ParseSheet
- tooltip_templ.go
- encoding/json.RawMessage
- select.js
- Queries
- live
- Session
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- ItemData
- wikilink.js
- mirror
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Weapon
- CombatantView
- time.Time
- app.js
- tabbable.js
- useDismiss
- ParseArmorTraining
- popover_templ.go
- command.js
- dialog.js
- .editEntity
- ParseWikidot
- tooltip.js
- github.com/a-h/templ.Component
- tabs_templ.go
- WizardView
- Sheet
- composite.js
- StarterSets
- CustomMod
- vendorjs/main.go
- CompendiumView
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Item
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- DownloadFiveTools
- parseHP
- radiogroup.js
- Player
- WeaponKind
- Weapon
- sqlc Config
- viewerCtx
- handoff
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- toast_templ.go
- checkbox.js
- ParseAidedd
- switch.js
- Open
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- Book
- useTypeahead
- testEnv
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- ftRec
- mounted
- newPlayerEnv
- newEffectFix
- profFocus
- redirect.go
- config.go
- ftBook
- run
- BlockEditor
- SearchReadableBlocksRow
- .ListBlockStepsSince
- cert_test.go
- buildReview
- .ListImageLookAccess
- build.sh
- .sendFeedback
- looseString
- errBodyWriter
- Client
- .ListItemsByCategory
- Before/after pairs
- CampaignOverview
- jsconfig.json
- open5e.go
- testing.M
- .ListSubracesOf
- Input
- TestLevelUpPushesTracker
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- planSwitch
- capReader
- sourcesEnv
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- pause_windows.go
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
- Search
- Habitat
- Defenses
- Browser check
- .coreTraits
- player-app.sh
- fieldError
- ListSharedWithRow
- IMPORT_FILTER_CAT_PREFIX
- .ListSpellsForClass

## God Nodes (most connected - your core abstractions)
1. `App` - 433 edges
2. `Path()` - 312 edges
3. `newTestEnv()` - 186 edges
4. `Button()` - 136 edges
5. `Sheet` - 134 edges
6. `CN()` - 114 edges
7. `SheetView` - 80 edges
8. `Queries` - 71 edges
9. `IfElse()` - 67 edges
10. `Command` - 60 edges

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

## Communities (185 total, 64 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (58): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+50 more)

### Community 1 - "Import"
Cohesion: 0.11
Nodes (34): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+26 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (164): keepAwake(), CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment (+156 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.13
Nodes (27): TakeHomeFile, readLimited(), ReadTakeHome(), rewriteZip(), sampleBringBack(), sampleFallback(), sampleNewCharacter(), sampleTakeHome() (+19 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.15
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 6 - "Button"
Cohesion: 0.04
Nodes (121): Coins, Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+113 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.09
Nodes (42): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeatRaisesCon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass() (+34 more)

### Community 9 - "tooltip_templ.go"
Cohesion: 0.08
Nodes (32): ListImageCharactersRow, ListImageEntitiesRow, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), campaignSourceRow() (+24 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (50): ftClass, ftEntry, ftFeat, ftItem, ftModOp, ftMonster, ftNamed, ftRace (+42 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Queries"
Cohesion: 0.04
Nodes (82): HomeManifest, TakeHomeBaseCharacter, TakeHomeBaseLook, TakeHomeBlock, TakeHomeBlockFile, TakeHomeCampaign, TakeHomeFallback, TakeHomeLevelUp (+74 more)

### Community 13 - "live"
Cohesion: 0.11
Nodes (28): ListEntityFieldBacklinksRow, ListBacklinksRow, initialChecked(), Switch(), CompendiumLive(), EntityDetail(), entityForm(), EntityView (+20 more)

### Community 14 - "Session"
Cohesion: 0.09
Nodes (24): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+16 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (51): ListCharacterNamesRow, internal/perm, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanChangeSubclass(), CanCharacter() (+43 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "App"
Cohesion: 0.03
Nodes (43): Identity, FromContext(), Campaign, browsePageSize(), compendiumBack(), compendiumPath(), ImportSummary(), pathCampaign() (+35 more)

### Community 21 - "ItemData"
Cohesion: 0.10
Nodes (29): ImageCategory, firstValue(), missingDamageItems(), usesItems(), weaponAbilities(), EffectFieldsForm(), effectModFields(), EffectModRow() (+21 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - "mirror"
Cohesion: 0.22
Nodes (10): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), replaceDir(), DownloadOpen5e(), open5eFiles() (+2 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (146): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+138 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Weapon"
Cohesion: 0.12
Nodes (6): Part, Skill, TestModifier(), skillByKey(), Modifier(), ProficiencyAt()

### Community 28 - "CombatantView"
Cohesion: 0.09
Nodes (39): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+31 more)

### Community 29 - "time.Time"
Cohesion: 0.05
Nodes (26): FiveTools2024Set(), browserName(), firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify(), TestFolderName() (+18 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 34 - "popover_templ.go"
Cohesion: 0.10
Nodes (29): Entity, Field, Look, Header(), state(), Title(), Trigger(), triggerID() (+21 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - ".editEntity"
Cohesion: 0.10
Nodes (15): linkTarget, linkTargets(), targetColumns(), Entity, LinkTarget, Fields, ParseFields(), entityTopics() (+7 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.12
Nodes (26): block, ClassFeature, ClassData, dnd5eFeatures(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+18 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (155): ListMonstersByNameRow, SwitchFeature, SwitchPlan, UI conventions, Action(), Alert(), Description(), Title() (+147 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.09
Nodes (45): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+37 more)

### Community 42 - "WizardView"
Cohesion: 0.06
Nodes (65): LevelUpPlan, SpellColumn(), ItemData, attrString(), generateSVG(), getIconContent(), Props, Icon() (+57 more)

### Community 43 - "Sheet"
Cohesion: 0.02
Nodes (140): LevelUpChoice, internal/compendium, Cantrips(), CoreTraits(), firstTable(), LevelSpells(), OptionPrerequisite(), Options() (+132 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "StarterSets"
Cohesion: 0.16
Nodes (11): ImportNotice, FiveToolsProgress(), importBar(), jobNotice(), offSources(), starterSet(), StarterSets(), FiveToolsJob (+3 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (24): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+16 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 48 - "CompendiumView"
Cohesion: 0.06
Nodes (53): CountEntriesBySourceRow, CountSharedBySourceRow, ListHomebrewLinksRow, SearchCompendiumRow, CompendiumSource, BuilderPickItems(), BuilderPicks(), BuilderPick (+45 more)

### Community 49 - "Path"
Cohesion: 0.04
Nodes (113): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), Gotchas, activeFight(), TestCombatantNamesFollowSource() (+105 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (137): TestDecode(), TestCoreTraits(), TestSkillChoice(), ftRaceAbility(), TestFiveTools2024Set(), feature(), RemoveSourceDir(), TestRemoveSourceDir() (+129 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Item"
Cohesion: 0.22
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

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
Cohesion: 0.10
Nodes (29): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+21 more)

### Community 59 - "parseHP"
Cohesion: 0.22
Nodes (11): Amount, parseCoin(), CritDamage(), Damage(), parseDamage(), parseHP(), parseMaxHP(), parseTemp() (+3 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (36): Player(), TestCanAddSources(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState() (+28 more)

### Community 62 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 63 - "Weapon"
Cohesion: 0.13
Nodes (10): ItemData, TestCustomMods(), Sum(), TestDualWield(), damageLabel(), Weapon, TestWithEffects(), compendiumWeapon() (+2 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "viewerCtx"
Cohesion: 0.13
Nodes (15): Topic, TopicKind, CanAddSources(), heldCollected(), notFoundIsNil(), parseLiveKey(), campaignLink(), PendingRequests() (+7 more)

### Community 66 - "handoff"
Cohesion: 0.08
Nodes (17): IsLoopback(), TestIsLoopback(), RequireAdmin(), generate(), LANAddresses(), PreferDefaultRoute(), preferFirst(), TestPreferFirst() (+9 more)

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.08
Nodes (42): CondOpKind, Effects, userText, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn() (+34 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (13): ArmorData, ftClassFeature, ftClassRefs, o5eArmor, ftArmor(), ftCode(), ftItemData(), TestFiveToolsItemMastery() (+5 more)

### Community 72 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.24
Nodes (17): ParseAidedd(), statLabel(), textLines(), TestParseAidedd(), ParsePlainPage(), attr(), byID(), collect() (+9 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Open"
Cohesion: 0.11
Nodes (28): migration, internal/db, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp() (+20 more)

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

### Community 85 - "Run Script"
Cohesion: 0.53
Nodes (8): changes(), GOFLAGS, lock(), restart_server(), run.sh script, stop_server(), superseded(), unlock()

### Community 86 - "Book"
Cohesion: 0.17
Nodes (17): rawBook, bookGroup(), fetchBooks(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks(), readBooks() (+9 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "testEnv"
Cohesion: 0.08
Nodes (23): ParsePageURL(), TestParsePageURL(), UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove() (+15 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "ftRec"
Cohesion: 0.21
Nodes (10): ftCopyRef, ftRec, ftCopyTargets(), ftDecode(), TestFiveToolsCopy(), ftMod(), ftName(), ftReplaceText() (+2 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.08
Nodes (41): eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack(), TestHandoffExpiresAndReplaces() (+33 more)

### Community 103 - "newEffectFix"
Cohesion: 0.10
Nodes (36): Condition, DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead() (+28 more)

### Community 105 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 106 - "config.go"
Cohesion: 0.22
Nodes (12): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), Player(), lookupMap() (+4 more)

### Community 107 - "ftBook"
Cohesion: 0.21
Nodes (12): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftPrerequisite(), ftRange() (+4 more)

### Community 108 - "run"
Cohesion: 0.14
Nodes (10): loopbackHostOnly(), main(), parseFlags(), run(), TestLoopbackHostOnly(), TestParseFlags(), openBrowser(), bookFolder() (+2 more)

### Community 113 - "cert_test.go"
Cohesion: 0.29
Nodes (9): Fingerprint(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+1 more)

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Client"
Cohesion: 0.05
Nodes (42): ClientInfo, countingHandler, event, fakeHandler, goneDevice, Handler, message, patchMessage (+34 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "CampaignOverview"
Cohesion: 0.14
Nodes (18): CampaignOverview(), CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, JoinInfo, agoTime() (+10 more)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 136 - "Input"
Cohesion: 0.07
Nodes (81): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+73 more)

### Community 137 - "TestLevelUpPushesTracker"
Cohesion: 0.31
Nodes (8): TestLevelUpPushesTracker(), dial(), patchOf(), readUntil(), rosterRefetch(), TestDeviceSeenOnLeave(), TestLiveUpdatesAndKick(), wsMessage

### Community 139 - "DownloadFiveTools"
Cohesion: 0.24
Nodes (12): TestFiveToolsBooks(), UseTestServer(), testEnv, pickerTicks(), serveStarterSets(), starterEnv(), TestImportsTickFromCampaignOnly(), TestPartialDownloadTicks() (+4 more)

### Community 140 - ".Entry"
Cohesion: 0.09
Nodes (27): MonsterAction, addEntry(), setStat(), MonsterData, TestTokensReadAsLabels(), firstInt(), HomebrewInput, HomebrewInputFor() (+19 more)

### Community 143 - "planSwitch"
Cohesion: 0.23
Nodes (6): featureSwitch, parseFeature(), featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 147 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 175 - "Search"
Cohesion: 0.09
Nodes (19): HomebrewOwner, Site, Source, sourceFormat, TestHomebrew(), fiveToolsPage(), ftsQuery(), SaveHomebrew() (+11 more)

### Community 176 - "Habitat"
Cohesion: 0.08
Nodes (38): CountHabitatsRow, ListEntryCopiesRow, Habitat, classBadges(), commaPrefixed(), compendiumFindURL(), habitatCountMap(), inHabitatOrder() (+30 more)

### Community 177 - "Defenses"
Cohesion: 0.33
Nodes (5): Defense, TestDefenses(), Defenses, ParseDefenses(), damageChoice()

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 179 - ".coreTraits"
Cohesion: 0.33
Nodes (5): ftProfTexts(), ftSkillNames(), TestFiveToolsCoreTraitsShapes(), withCoreTraits(), ArmorKindNamed()

### Community 181 - "fieldError"
Cohesion: 0.40
Nodes (3): fieldError, pickError, userError

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 474 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **64 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Import`, `context.Context`, `CampaignOverview`, `export/markdown.go`, `Button`, `ParseSheet`, `TestLevelUpPushesTracker`, `Input`, `DownloadFiveTools`, `Queries`, `live`, `tooltip_templ.go`, `perm.go`, `sourcesEnv`, `App`, `ItemData`, `CombatantView`, `popover_templ.go`, `.editEntity`, `github.com/a-h/templ.Component`, `tabs_templ.go`, `WizardView`, `Sheet`, `StarterSets`, `CompendiumView`, `Habitat`, `testing.T`, `Player`, `viewerCtx`, `ParseWiki`, `testEnv`, `newPlayerEnv`, `newEffectFix`, `.sendFeedback`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14487444123959373 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `Player`, `wikilink.js`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **Should `Import` be split into smaller, more focused modules?**
  _Cohesion score 0.11229946524064172 - nodes in this community are weakly interconnected._
- **Why does `routePath()` connect `wikilink.js` to `testing.T`, `notes.js`, `BlockEditor`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._