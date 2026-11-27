# Graph Report - gm-session-manager  (2026-11-27)

## Corpus Check
- 408 files · ~729,549 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 61 file(s) not represented in the graph (top: .templ 51, .css 4, (none) 3)

## Summary
- 5374 nodes · 28991 edges · 194 communities (132 shown, 62 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2877 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `fcdce46e`
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
- CampaignSources
- encoding/json.RawMessage
- select.js
- Character
- Path
- Session
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- dialog_templ.go
- wikilink.js
- mirror
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Sheet
- CombatantView
- .fail
- app.js
- tabbable.js
- useDismiss
- CompendiumEntry
- Encounter
- command.js
- dialog.js
- .saveImage
- ParseWikidot
- tooltip.js
- github.com/a-h/templ.Component
- tabs_templ.go
- IfElse
- Slugify
- composite.js
- Topic
- CustomMod
- vendorjs/main.go
- CompendiumPath
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- .wizardView
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- DownloadFiveTools
- .encounterView
- radiogroup.js
- Player
- CompendiumKind
- monsterBody
- sqlc Config
- Client
- handoff
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- planBringBack
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
- TestSourceSwitchKeepsEdits
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- Option
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
- .editEntity
- cert_test.go
- buildReview
- alert_templ.go
- build.sh
- .sendFeedback
- looseString
- errBodyWriter
- Player
- State
- Before/after pairs
- .deleteUnused
- jsconfig.json
- open5e.go
- testing.M
- rosterCard
- Input
- Hub
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- planSwitch
- capReader
- lan.go
- Pending
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
- ftItemData
- .fiveToolsView
- Browser check
- fakeHandler
- player-app.sh
- effectSheet
- ListSharedWithRow
- IMPORT_FILTER_CAT_PREFIX
- toast_templ.go
- names
- .entityPage
- Look
- .device
- sourceRow
- .mountAssets
- loadSettings
- serveSplit

## God Nodes (most connected - your core abstractions)
1. `App` - 442 edges
2. `Path()` - 328 edges
3. `newTestEnv()` - 193 edges
4. `Sheet` - 141 edges
5. `Button()` - 141 edges
6. `CN()` - 114 edges
7. `SheetView` - 80 edges
8. `Queries` - 77 edges
9. `IfElse()` - 67 edges
10. `Command` - 60 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (194 total, 62 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (58): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+50 more)

### Community 1 - "Import"
Cohesion: 0.11
Nodes (18): TestRealSource(), Import(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkImportReview(), Gotchas, internal/web, joinHabitats() (+10 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (149): keepAwake(), SpellData, ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListImageEntitiesRow (+141 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.20
Nodes (21): ReadTakeHome(), rewriteZip(), sampleBringBack(), sampleFallback(), sampleNewCharacter(), sampleTakeHome(), TestBringBackFallbackRoundTrip(), TestBringBackRoundTrip() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.13
Nodes (17): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+9 more)

### Community 6 - "Button"
Cohesion: 0.05
Nodes (121): Slot, Signed(), Button(), Variant, Action(), Card(), Content(), Description() (+113 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.08
Nodes (48): TakeHome, ParseSheet(), testEnv, homeSheet(), homeTableSheet(), TestBringBackAccept(), TestBringBackPlayerUpload(), TestBringBackReject() (+40 more)

### Community 9 - "CampaignSources"
Cohesion: 0.09
Nodes (27): TestSourcesCardStarterSets(), campaignSourceRow(), CampaignSources(), CampaignSourcesLive(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), CampaignRoster() (+19 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (34): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+26 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Character"
Cohesion: 0.03
Nodes (96): HomebrewOwner, HomeManifest, TakeHomeBaseCharacter, TakeHomeBaseLook, TakeHomeBlock, TakeHomeBlockFile, TakeHomeCampaign, TakeHomeFallback (+88 more)

### Community 13 - "Path"
Cohesion: 0.05
Nodes (85): ListEntityFieldBacklinksRow, ListBacklinksRow, Path(), UI conventions, Scripts(), TestCharacterLooks(), TestImageAccess(), TestLiveRefetchStrips() (+77 more)

### Community 14 - "Session"
Cohesion: 0.08
Nodes (28): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+20 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (55): ListCharacterNamesRow, Visibility, internal/perm, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks() (+47 more)

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
Nodes (50): Identity, FromContext(), RequireAdmin(), Campaign, GetBlockRow, compendiumBack(), compendiumPath(), ImportSummary() (+42 more)

### Community 21 - "dialog_templ.go"
Cohesion: 0.08
Nodes (56): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+48 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - "mirror"
Cohesion: 0.20
Nodes (12): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), filterSource(), replaceDir(), TestMirrorGetMissingFile() (+4 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (149): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+141 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Sheet"
Cohesion: 0.04
Nodes (36): Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), ArmorEffect(), skillByKey(), Sum() (+28 more)

### Community 28 - "CombatantView"
Cohesion: 0.12
Nodes (30): beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList(), downBadge(), Encounter() (+22 more)

### Community 29 - ".fail"
Cohesion: 0.09
Nodes (9): Character, CanChangeSubclass(), CanCharacter(), homeBuildSources(), nextStep(), wizardPath(), SetCharacterCurrentImageParams, ToggleCharacterImageRevealedParams (+1 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "CompendiumEntry"
Cohesion: 0.10
Nodes (28): Decode(), CompendiumEntry, TestSlotMaxima(), ClassInfo, ClassLevel, CasterLevel(), SlotMaxima(), Spellbook() (+20 more)

### Community 34 - "Encounter"
Cohesion: 0.12
Nodes (13): Encounter, EncounterStatus, CanAddSources(), TestCanAddSources(), notFoundIsNil(), parseLiveKey(), CreateEncounterParams, ListCharacterEncountersParams (+5 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - ".saveImage"
Cohesion: 0.06
Nodes (22): GetImageRow, Attachment, firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify(), TestFolderName() (+14 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.12
Nodes (24): block, ClassFeature, ClassData, feature(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (116): Props, Variant, Action(), AspectRatio(), Badge(), badgeClasses(), Footer(), ChipsInput() (+108 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.09
Nodes (38): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+30 more)

### Community 42 - "IfElse"
Cohesion: 0.08
Nodes (52): Props, Field, Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup() (+44 more)

### Community 43 - "Slugify"
Cohesion: 0.07
Nodes (52): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, TextData (+44 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Topic"
Cohesion: 0.20
Nodes (6): event, Handler, Topic, TopicKind, heldCollected(), heldTopics

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (24): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+16 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 48 - "CompendiumPath"
Cohesion: 0.06
Nodes (53): CountCompendiumRow, CountEntriesBySourceRow, CountHabitatsRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, browsePageSize(), TestBrowsePageSize() (+45 more)

### Community 49 - "Path"
Cohesion: 0.03
Nodes (95): PageSourceKey(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestDeleteActiveCombatant(), TestDeleteActiveCombatantLast(), TestDeleteActiveCombatantSkipsOut(), TestDeleteCharacterActiveCombatant() (+87 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (148): Defense, TestDecode(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), TestFiveToolsBooks(), RemoveSourceDir(), TestRemoveSourceDir() (+140 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - ".wizardView"
Cohesion: 0.09
Nodes (26): EquipmentItem, internal/compendium, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), DecodeData(), TestFiveToolsClasses() (+18 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.06
Nodes (21): LevelUpChoice, ItemData, armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining() (+13 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "DownloadFiveTools"
Cohesion: 0.14
Nodes (22): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook() (+14 more)

### Community 59 - ".encounterView"
Cohesion: 0.08
Nodes (32): Combatant, ListCombatantsRow, combatantRow(), rollFor(), stepTurn(), formInt(), combatantKind(), combatantLink() (+24 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.11
Nodes (23): Websocket and live views, TestStarterSets(), isRandomID(), shapeHash(), TestLiveAffects(), TestLiveBindsMatchState(), TestLiveHotKeepsHash(), TestLiveModelHashFollowsMarkup() (+15 more)

### Community 62 - "CompendiumKind"
Cohesion: 0.06
Nodes (23): dataFor(), CompendiumKind, CountSharedByKindParams, CountSharedByKindRow, CreateHomebrewParams, GetCompendiumEntryByNameParams, GetCompendiumEntryBySlugParams, GetCompendiumEntryBySlugSourceParams (+15 more)

### Community 63 - "monsterBody"
Cohesion: 0.10
Nodes (34): SearchCompendiumRow, pickerRow(), SearchResults(), TestRowFigure(), commaPrefixed(), compendiumFindURL(), crText(), PickerTarget (+26 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 66 - "handoff"
Cohesion: 0.28
Nodes (4): closeServers(), removeFile(), shutdownServers(), handoff

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.05
Nodes (64): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+56 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.15
Nodes (11): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey(), ftProfTexts(), ftSkillNames() (+3 more)

### Community 72 - "planBringBack"
Cohesion: 0.10
Nodes (28): sheetMap, finish(), GroupDiff, GroupKeys(), GroupLabel(), Merge(), signature(), TakeBuild() (+20 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Open"
Cohesion: 0.07
Nodes (47): linkTarget, migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+39 more)

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
Cohesion: 0.33
Nodes (8): CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks(), WikiLinks(), WikiLink, WikiPart

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
Cohesion: 0.18
Nodes (16): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+8 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "Option"
Cohesion: 0.08
Nodes (24): ListSubclassesOfRow, ClassFeature, LevelUpPlan, SpellColumn(), subclassIndex(), subclassOptions(), subclassPanels(), withSpells() (+16 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.06
Nodes (49): TestEffectFields(), eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack() (+41 more)

### Community 103 - "newEffectFix"
Cohesion: 0.09
Nodes (41): DeathSaves, nullID(), BenchmarkEncounterState(), deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP() (+33 more)

### Community 105 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 106 - "config.go"
Cohesion: 0.22
Nodes (12): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), Player(), lookupMap() (+4 more)

### Community 107 - "ftBook"
Cohesion: 0.13
Nodes (29): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftFeatIncrease(), ftParts() (+21 more)

### Community 108 - "run"
Cohesion: 0.14
Nodes (10): loopbackHostOnly(), main(), parseFlags(), run(), TestLoopbackHostOnly(), TestParseFlags(), openBrowser(), bookFolder() (+2 more)

### Community 112 - ".editEntity"
Cohesion: 0.10
Nodes (10): Entity, entityTopics(), linkEntityLook(), linkCharacterLook(), lookFree(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams (+2 more)

### Community 113 - "cert_test.go"
Cohesion: 0.15
Nodes (10): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+2 more)

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 117 - "alert_templ.go"
Cohesion: 0.14
Nodes (17): ActionProps, DescriptionProps, Props, TitleProps, Variant, Alert(), Description(), Title() (+9 more)

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

### Community 125 - "Player"
Cohesion: 0.16
Nodes (25): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign() (+17 more)

### Community 127 - "State"
Cohesion: 0.16
Nodes (8): patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff(), TestPatch()

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - ".deleteUnused"
Cohesion: 0.36
Nodes (3): ListAttachmentFilesRow, fresh(), inCampaign()

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.09
Nodes (31): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery() (+23 more)

### Community 135 - "rosterCard"
Cohesion: 0.40
Nodes (8): deviceRow(), rosterCard(), devicesWhere(), agoTime(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 136 - "Input"
Cohesion: 0.05
Nodes (87): ImageCategory, Coins, Defenses, Combobox(), Content(), Item(), List(), Description() (+79 more)

### Community 137 - "Hub"
Cohesion: 0.22
Nodes (3): ClientInfo, goneDevice, Hub

### Community 139 - "DownloadFiveTools"
Cohesion: 0.33
Nodes (9): testEnv, pickerTicks(), serveStarterSets(), starterEnv(), TestImportsTickFromCampaignOnly(), TestPartialDownloadTicks(), TestPickerKeepsJobTicks(), TestStarterSetsImport() (+1 more)

### Community 140 - ".Entry"
Cohesion: 0.10
Nodes (23): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInputFor(), abilityKey(), actionsText() (+15 more)

### Community 143 - "planSwitch"
Cohesion: 0.22
Nodes (6): featureSwitch, featureNames(), SubclassRef, PlanSourceSwitch(), planSwitch(), sameSlug()

### Community 147 - "lan.go"
Cohesion: 0.47
Nodes (5): LANAddresses(), PreferDefaultRoute(), preferFirst(), TestPreferFirst(), handoffIPs()

### Community 148 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 175 - "Search"
Cohesion: 0.12
Nodes (18): Site, Source, sourceFormat, fiveToolsPage(), formatOf(), Report, ImportProgress(), importSource() (+10 more)

### Community 176 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 177 - ".fiveToolsView"
Cohesion: 0.39
Nodes (7): TestTickSet(), tickBooks(), tickSet(), FiveToolsView, NewFiveToolsView(), FiveToolsBook, FiveToolsGroup

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 181 - "effectSheet"
Cohesion: 0.17
Nodes (10): Handlers, TestMalformedEffectSkipped(), TestKeepsBodyBringBack(), keepsBody(), formErrors(), effectName(), effectSheet(), fieldError (+2 more)

### Community 184 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 185 - "names"
Cohesion: 0.40
Nodes (5): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), SourcePreviewView

### Community 186 - ".entityPage"
Cohesion: 0.48
Nodes (4): targetColumns(), LinkTarget, saveFieldLinks(), linkTarget()

### Community 188 - "Look"
Cohesion: 0.43
Nodes (7): Look, initials(), lookControls(), looksGrid(), npcPortrait(), currentLook(), LookItem()

### Community 189 - ".device"
Cohesion: 0.38
Nodes (5): testEnv, TestRequestLog(), TestRequestLogCharacterRemoved(), TestRequestLogReusedID(), TestRequestSent()

### Community 190 - "sourceRow"
Cohesion: 0.47
Nodes (5): sourceIssue(), sourceRow(), SourceCopy, SourceIssue, sheetRef

### Community 192 - "loadSettings"
Cohesion: 0.40
Nodes (5): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpRedownloads()

### Community 193 - "serveSplit"
Cohesion: 0.50
Nodes (3): serveSplit(), TestSplitTLS(), TestSplitTLSSilentConn()

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 477 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **62 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.deleteUnused`, `Hub`, `CampaignSources`, `Character`, `Session`, `perm.go`, `Pending`, `Sheet`, `.fail`, `CompendiumEntry`, `Encounter`, `.saveImage`, `Slugify`, `Topic`, `CompendiumPath`, `Path`, `.fiveToolsView`, `.wizardView`, `ParseEffects`, `.entityPage`, `.encounterView`, `.mountAssets`, `handoff`, `Command`, `Book`, `.editEntity`, `cert_test.go`, `State`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14800606420143694 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Import`, `.deleteUnused`, `export/markdown.go`, `Button`, `rosterCard`, `ParseSheet`, `CampaignSources`, `encoding/json.RawMessage`, `DownloadFiveTools`, `Character`, `Input`, `perm.go`, `App`, `dialog_templ.go`, `CombatantView`, `.fail`, `.saveImage`, `github.com/a-h/templ.Component`, `tabs_templ.go`, `IfElse`, `Slugify`, `CompendiumPath`, `Path`, `testing.T`, `.encounterView`, `Look`, `.device`, `Player`, `monsterBody`, `sourceRow`, `ParseWiki`, `TestSourceSwitchKeepsEdits`, `newPlayerEnv`, `newEffectFix`, `.editEntity`, `alert_templ.go`, `.sendFeedback`, `Player`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **Should `Import` be split into smaller, more focused modules?**
  _Cohesion score 0.11052631578947368 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Import`, `Command`, `Path`, `wikilink.js`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._