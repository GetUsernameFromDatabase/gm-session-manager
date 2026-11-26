# Graph Report - gm-session-manager  (2026-11-26)

## Corpus Check
- 375 files · ~668,871 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4900 nodes · 25604 edges · 180 communities (120 shown, 60 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2556 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f5aad862`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Import
- context.Context
- vendorjs/main.go
- createProvider
- export/markdown.go
- SheetView
- combobox.js
- ParseSheet
- Path
- encoding/json.RawMessage
- select.js
- Command
- SourcePriority
- redirect.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- WizardView
- wikilink.js
- .renderCompendium
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ImportView
- CombatantView
- formSelect
- app.js
- tabbable.js
- useDismiss
- Slugify
- popover_templ.go
- command.js
- dialog.js
- time.Time
- ParseWikidot
- tooltip.js
- NewClient
- tabs_templ.go
- viewerCtx
- Hub
- composite.js
- Queries
- CustomMod
- Open
- monsterBody
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Session
- shadcn-templ Config
- ParseArmorTraining
- golangci-lint Config
- Item
- Weapon
- radiogroup.js
- Player
- ClassInfo
- formErrors
- sqlc Config
- sourceRow
- testEnv
- Claude Settings
- Client
- use_transition_status.js
- playwright
- ftPrerequisite
- mirror
- checkbox.js
- ParseAidedd
- switch.js
- github.com/a-h/templ.Component
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
- Values Go and JS share
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- CompendiumKind
- mounted
- CharacterFile
- nullID
- profFocus
- cert.go
- .source
- ftBook
- Pending
- BlockEditor
- ListSharedWithRow
- StarterSets
- capReader
- Sum
- Amount
- build.sh
- deathFix
- looseString
- errBodyWriter
- sourcesEnv
- ftItemData
- Before/after pairs
- toast_templ.go
- jsconfig.json
- open5e.go
- testing.M
- fakeHandler
- perm/CLAUDE.md
- names
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- ImagePage
- Device
- parseMod
- condMods
- WeaponKind
- .ListBlockStepsSince
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- .ListImageLookAccess
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
- planSwitch
- .ListItemsByCategory
- .ListSpellsForClass
- .ListSubracesOf
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 387 edges
2. `Path()` - 278 edges
3. `newTestEnv()` - 178 edges
4. `Sheet` - 134 edges
5. `Button()` - 126 edges
6. `CN()` - 114 edges
7. `SheetView` - 79 edges
8. `IfElse()` - 67 edges
9. `Command` - 60 edges
10. `Queries` - 57 edges

## Surprising Connections (you probably didn't know these)
- `What to check` --references--> `check()`  [INFERRED]
  .claude/skills/browser-check/SKILL.md → cmd/vendorjs/main.go
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (180 total, 60 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (56): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+48 more)

### Community 1 - "Import"
Cohesion: 0.10
Nodes (37): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+29 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (127): keepAwake(), ftsQuery(), Search(), targetColumns(), ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, CompendiumSource (+119 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.12
Nodes (26): Browser check, Roles, Saved states, The test instance, What to check, You are alone, canonical(), changedPaths() (+18 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.05
Nodes (49): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+41 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (120): Slot, Signed(), Button(), Variant, Action(), Card(), Content(), Description() (+112 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.09
Nodes (43): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeatRaisesCon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass() (+35 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (106): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+98 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (52): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+44 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.10
Nodes (39): CondOpKind, Effects, parseFeature(), parseSet(), parseSpell(), parseCoin(), parseInspiration(), parseShield() (+31 more)

### Community 13 - "SourcePriority"
Cohesion: 0.19
Nodes (16): CountEntriesBySourceRow, CountSharedBySourceRow, Empty(), sourceCountLink(), SourcePriority(), SourcePriorityLive(), entryCountClass(), entryCountText() (+8 more)

### Community 14 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (66): ListCharacterNamesRow, ListCombatantsRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+58 more)

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
Nodes (35): ClearCookie(), SetCookie(), writeCookie(), FromContext(), RequireAdmin(), Campaign, LinkTarget, pathCampaign() (+27 more)

### Community 21 - "WizardView"
Cohesion: 0.08
Nodes (45): SpellColumn(), attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item() (+37 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - ".renderCompendium"
Cohesion: 0.11
Nodes (11): CountCompendiumRow, browsePageSize(), ImportSummary(), fiveToolsBack(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), CompendiumBase() (+3 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (181): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Variant (+173 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ImportView"
Cohesion: 0.16
Nodes (14): ImageCategory, categoryRow(), columnOverview(), ImportReview(), importRows(), importStep(), keepSection(), columnLabel() (+6 more)

### Community 28 - "CombatantView"
Cohesion: 0.11
Nodes (34): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList() (+26 more)

### Community 29 - "formSelect"
Cohesion: 0.04
Nodes (113): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+105 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Slugify"
Cohesion: 0.04
Nodes (105): EquipmentItem, LevelUpChoice, internal/compendium, Cantrips(), EquipmentOptions(), firstTable(), EquipmentOption, LevelSpells() (+97 more)

### Community 34 - "popover_templ.go"
Cohesion: 0.12
Nodes (22): Description(), Header(), state(), Title(), Trigger(), triggerID(), modTargetLabel(), effectBits() (+14 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "time.Time"
Cohesion: 0.04
Nodes (26): IsLoopback(), Identity, ListAttachmentFilesRow, ListImagesRow, Attachment, Session, Role, browserName() (+18 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.10
Nodes (30): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, dnd5eFeatures(), TestFiveToolsCoreTraitsShapes(), withCoreTraits() (+22 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "NewClient"
Cohesion: 0.18
Nodes (19): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDisconnectDropsRequest() (+11 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.09
Nodes (42): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+34 more)

### Community 42 - "viewerCtx"
Cohesion: 0.21
Nodes (6): notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView, viewerCtx

### Community 43 - "Hub"
Cohesion: 0.16
Nodes (7): ClientInfo, event, goneDevice, TestFiveToolsClassLevelUp(), Hub, TestCloseSessionAndCharacter(), TestDeviceAnnouncements()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Queries"
Cohesion: 0.05
Nodes (49): sourceFormat, FiveToolsImages(), bindFiveToolsImages(), SeedHabitats(), Queries, Combatant, Entity, Character (+41 more)

### Community 46 - "CustomMod"
Cohesion: 0.10
Nodes (15): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+7 more)

### Community 47 - "Open"
Cohesion: 0.10
Nodes (28): linkTarget, migration, internal/db, apply(), loadMigrations(), Migrate(), Open(), openPool() (+20 more)

### Community 48 - "monsterBody"
Cohesion: 0.15
Nodes (24): TestRowFigure(), commaPrefixed(), compendiumFindURL(), crText(), levelsText(), parenthesized(), rowFigure(), signedMap() (+16 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (150): Defense, TestIsLoopback(), TestDecode(), TestEquipmentOptions(), TestSkillChoice(), TestFiveToolsBooks(), RemoveSourceDir(), TestRemoveSourceDir() (+142 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Session"
Cohesion: 0.09
Nodes (27): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+19 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseArmorTraining"
Cohesion: 0.13
Nodes (12): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+4 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Item"
Cohesion: 0.19
Nodes (9): findItem(), parseArmor(), parseItem(), ArmorEffect(), TestGearCommands(), ensureItem(), Item, WornArmor (+1 more)

### Community 59 - "Weapon"
Cohesion: 0.10
Nodes (14): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+6 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.10
Nodes (33): Player(), TestCanRemoveCondition(), BenchmarkEncounterState(), BenchmarkSheetState(), Websocket and live views, watchSheets(), TestStarterSets(), isRandomID() (+25 more)

### Community 62 - "ClassInfo"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 63 - "formErrors"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "sourceRow"
Cohesion: 0.47
Nodes (5): sourceIssue(), sourceRow(), SourceCopy, SourceIssue, sheetRef

### Community 66 - "testEnv"
Cohesion: 0.16
Nodes (8): testEnv, assetETags(), immutableAsset(), TestEntityFieldLinks(), testEnv, TestAdminLogin(), TestCampaignJoinLinks(), TestLogout()

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Client"
Cohesion: 0.11
Nodes (9): Handler, message, patchMessage, stateHandler, viewPatch, Client, Diff(), State (+1 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.24
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - "mirror"
Cohesion: 0.21
Nodes (11): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), replaceDir(), TestMirrorGetMissingFile(), DownloadOpen5e() (+3 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "github.com/a-h/templ.Component"
Cohesion: 0.03
Nodes (167): CountHabitatsRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, SearchCompendiumRow, SearchReadableBlocksRow, Habitat, UI conventions (+159 more)

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
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - "Book"
Cohesion: 0.11
Nodes (26): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+18 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Values Go and JS share"
Cohesion: 0.15
Nodes (14): Gotchas, internal/web, Values Go and JS share, TestEffectNoteMaxMatchesTag(), TestPagesRender(), TestNoRouteLiterals(), parseSharedJS(), tagHas() (+6 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "CompendiumKind"
Cohesion: 0.03
Nodes (56): Site, dataFor(), DecodeData(), ValidKind(), fiveToolsPage(), FetchPage(), Page, pageSource() (+48 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 103 - "nullID"
Cohesion: 0.20
Nodes (21): TestMonsterAtZeroHPUnaffected(), acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectFields() (+13 more)

### Community 105 - "cert.go"
Cohesion: 0.10
Nodes (19): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+11 more)

### Community 106 - ".source"
Cohesion: 0.12
Nodes (17): ParsePageURL(), TestParsePageURL(), UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove() (+9 more)

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (20): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftParts(), ftPrerequisite() (+12 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 112 - "StarterSets"
Cohesion: 0.23
Nodes (7): FiveToolsProgress(), importBar(), jobNotice(), StarterSets(), FiveToolsJob, starterBase(), fiveToolsJob

### Community 116 - "Sum"
Cohesion: 0.19
Nodes (4): TestCustomMods(), TestExplainSums(), Sum(), TestWithEffects()

### Community 117 - "Amount"
Cohesion: 0.22
Nodes (7): Amount, userText, CritDamage(), Damage(), parseRevive(), Set(), TestApplyHPDeath()

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 127 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eHabitats(), TestOpen5eMastery() (+16 more)

### Community 137 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 139 - "DownloadFiveTools"
Cohesion: 0.16
Nodes (20): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook(), DownloadFiveTools() (+12 more)

### Community 140 - ".Entry"
Cohesion: 0.07
Nodes (33): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, habitatsFrom(), TestHabitatsFrom() (+25 more)

### Community 143 - "ImagePage"
Cohesion: 0.08
Nodes (39): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListBacklinksRow, Entity, Look, initials(), EntityDetail() (+31 more)

### Community 144 - "Device"
Cohesion: 0.22
Nodes (12): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), localTime(), Plural() (+4 more)

### Community 145 - "parseMod"
Cohesion: 0.50
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 146 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

### Community 147 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 175 - "planSwitch"
Cohesion: 0.25
Nodes (5): featureSwitch, featureNames(), SubclassRef, planSwitch(), sameSlug()

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 456 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **60 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Slugify`, `testEnv`, `context.Context`, `CompendiumKind`, `time.Time`, `cert.go`, `viewerCtx`, `Hub`, `Pending`, `Queries`, `perm.go`, `StarterSets`, `testing.T`, `Book`, `.renderCompendium`?**
  _High betweenness centrality (0.057) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14559682741500923 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Import`, `export/markdown.go`, `SheetView`, `ParseSheet`, `ImagePage`, `perm.go`, `Device`, `App`, `WizardView`, `.renderCompendium`, `ImportView`, `CombatantView`, `formSelect`, `Slugify`, `popover_templ.go`, `time.Time`, `tabs_templ.go`, `Queries`, `monsterBody`, `testing.T`, `Session`, `Player`, `sourceRow`, `testEnv`, `github.com/a-h/templ.Component`, `ParseWiki`, `Values Go and JS share`, `CompendiumKind`, `nullID`, `.source`, `deathFix`, `sourcesEnv`?**
  _High betweenness centrality (0.036) - this node is a cross-community bridge._
- **Should `Import` be split into smaller, more focused modules?**
  _Cohesion score 0.09815078236130868 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `Values Go and JS share` to `Path`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.030) - this node is a cross-community bridge._