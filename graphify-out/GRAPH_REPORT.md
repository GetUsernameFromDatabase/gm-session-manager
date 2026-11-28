# Graph Report - gm-session-manager  (2026-11-28)

## Corpus Check
- 443 files · ~773,515 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 65 file(s) not represented in the graph (top: .templ 51, (none) 4, .css 4)

## Summary
- 5771 nodes · 31012 edges · 204 communities (139 shown, 65 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 3060 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `210d7308`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- db/db.go
- context.Context
- takehome_test.go
- createProvider
- report
- SheetView
- combobox.js
- Weapon
- DownloadFiveTools
- encoding/json.RawMessage
- select.js
- select_templ.go
- ParseArmorTraining
- Import
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- entityForm
- CompendiumPath
- wikilink.js
- CompendiumKind
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- popover_templ.go
- CombatantView
- TestBringBackAccept
- app.js
- tabbable.js
- useDismiss
- Player
- Encounter
- command.js
- dialog.js
- handoff
- ParseWikidot
- tooltip.js
- Input
- tabs_templ.go
- WizardView
- Hub
- composite.js
- .sendFeedback
- Feature
- vendorjs/main.go
- github.com/a-h/templ.Component
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Sheet
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- ftRec
- CampaignSources
- radiogroup.js
- Topic
- .Router
- Part
- sqlc Config
- App
- parseHP
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- Condition
- checkbox.js
- ParseAidedd
- switch.js
- app
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- TestTakeHomeDownload
- homeImportForm
- useTypeahead
- .source
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- Client
- mounted
- newPlayerEnv
- Session
- profFocus
- cert.go
- NewClient
- ftBook
- Book
- BlockEditor
- planBringBack
- mirror
- starterEnv
- buildReview
- vars
- ListSharedWithRow
- .q
- looseString
- errBodyWriter
- Pending
- task
- Before/after pairs
- env
- jsconfig.json
- open5e.go
- testing.M
- Queries
- alert_templ.go
- parseMod
- ASIError
- .ListBlockStepsSince
- .Entry
- sourceUsedError
- run.sh
- .saveReport
- .ListImageLookAccess
- .ListItemsByCategory
- main
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
- deathFix
- Browser check
- dev:changed
- release.sh
- .ListSubracesOf
- systemd-units.sh
- IMPORT_FILTER_CAT_PREFIX
- Env
- internal/web
- dev:build
- build.sh
- before
- release
- dev:generate
- lint
- experiments
- test
- capReader
- tasks
- Input
- lint.sh
- WeaponKind
- run:before
- dev:migrations
- run
- run:test-player

## God Nodes (most connected - your core abstractions)
1. `App` - 463 edges
2. `Path()` - 347 edges
3. `newTestEnv()` - 204 edges
4. `Button()` - 147 edges
5. `Sheet` - 142 edges
6. `CN()` - 114 edges
7. `Queries` - 97 edges
8. `SheetView` - 83 edges
9. `IfElse()` - 68 edges
10. `ParseSheet()` - 64 edges

## Surprising Connections (you probably didn't know these)
- `The timer` --references--> `install()`  [INFERRED]
  deploy/README.md → cmd/vendorjs/main.go
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
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

## Communities (204 total, 65 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (59): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+51 more)

### Community 1 - "db/db.go"
Cohesion: 0.09
Nodes (31): linkTarget, migration, internal/db, apply(), loadMigrations(), Migrate(), Open(), openPool() (+23 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (161): keepAwake(), sortedKeys(), SortSources(), undated(), writeOrder(), ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow (+153 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.15
Nodes (26): writeFile(), readLimited(), ReadTakeHome(), plainTakeHome(), rewriteZip(), sampleBringBack(), sampleFallback(), sampleNewCharacter() (+18 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "report"
Cohesion: 0.14
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (146): Slot, Signed(), Button(), Variant, Action(), Card(), Content(), Description() (+138 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Weapon"
Cohesion: 0.13
Nodes (7): findItem(), parseArmor(), TestCustomMods(), Sum(), TestGearCommands(), ensureItem(), Item

### Community 9 - "DownloadFiveTools"
Cohesion: 0.11
Nodes (26): habitatsFrom(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants() (+18 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (49): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+41 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "select_templ.go"
Cohesion: 0.12
Nodes (21): Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state(), Trigger() (+13 more)

### Community 13 - "ParseArmorTraining"
Cohesion: 0.13
Nodes (13): armorGrants(), ArmorKindNamed(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack() (+5 more)

### Community 14 - "Import"
Cohesion: 0.16
Nodes (27): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestSearch() (+19 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.04
Nodes (64): ListCharacterNamesRow, GetBlockRow, Visibility, internal/perm, Entity, Look, VisibleEntity(), VisibleLooks() (+56 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "entityForm"
Cohesion: 0.08
Nodes (33): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ImageCategory, ListBacklinksRow, Field, attrString(), generateSVG() (+25 more)

### Community 21 - "CompendiumPath"
Cohesion: 0.05
Nodes (59): CountCompendiumRow, CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, Habitat (+51 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - "CompendiumKind"
Cohesion: 0.03
Nodes (56): Site, dataFor(), DecodeData(), ValidKind(), fiveToolsPage(), ftsQuery(), Search(), FetchPage() (+48 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (194): Props, Props, Variant, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+186 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 28 - "CombatantView"
Cohesion: 0.11
Nodes (33): beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList(), downBadge(), Encounter() (+25 more)

### Community 29 - "TestBringBackAccept"
Cohesion: 0.13
Nodes (23): TakeHome, testEnv, mentionDoc(), newExtrasEnv(), TestBringBackAcceptExtras(), TestBringBackAcceptFiles(), TestBringBackStaleExtras(), TestNewCharacterFromGM() (+15 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (23): checkLink(), checkLinks(), clearDisplay(), clearSourceSelection(), commandChoices(), commandHint(), commandState(), createItem() (+15 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Player"
Cohesion: 0.05
Nodes (70): Player(), TestCanAddSources(), TestCanRemoveCondition(), nullID(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views (+62 more)

### Community 34 - "Encounter"
Cohesion: 0.20
Nodes (8): CanAddSources(), notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "handoff"
Cohesion: 0.23
Nodes (6): IsLoopback(), TestIsLoopback(), closeServers(), removeFile(), shutdownServers(), handoff

### Community 38 - "ParseWikidot"
Cohesion: 0.12
Nodes (25): block, ClassFeature, ClassData, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+17 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Input"
Cohesion: 0.05
Nodes (89): Coins, Defenses, Combobox(), Content(), Item(), List(), Description(), Field() (+81 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.08
Nodes (42): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+34 more)

### Community 42 - "WizardView"
Cohesion: 0.06
Nodes (68): LevelUpPlan, SpellColumn(), Checkbox(), Label(), initialValue(), Item(), RadioGroup(), state() (+60 more)

### Community 43 - "Hub"
Cohesion: 0.27
Nodes (3): ClientInfo, goneDevice, Hub

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 46 - "Feature"
Cohesion: 0.10
Nodes (16): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+8 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.11
Nodes (29): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+21 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (125): SwitchFeature, SwitchPlan, UI conventions, Badge(), Scripts(), Separator(), Spinner(), Toast() (+117 more)

### Community 49 - "Path"
Cohesion: 0.03
Nodes (143): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount() (+135 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (152): TestNamesMatchGo(), Defense, TestDecode(), EquipmentOptions(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), feature() (+144 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Sheet"
Cohesion: 0.04
Nodes (102): EquipmentItem, internal/compendium, Cantrips(), CoreTraits(), firstTable(), EquipmentOption, LevelSpells(), OptionPrerequisite() (+94 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.09
Nodes (16): LevelUpChoice, TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+8 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (35): deny, files, rules, generated, paths, rules, formatters, enable (+27 more)

### Community 58 - "ftRec"
Cohesion: 0.27
Nodes (8): ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve(), ftVersions()

### Community 59 - "CampaignSources"
Cohesion: 0.09
Nodes (27): campaignSourceRow(), CampaignSources(), CampaignSourcesLive(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), CampaignOverview(), CampaignRoster() (+19 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Topic"
Cohesion: 0.10
Nodes (14): event, Handler, message, patchMessage, stateHandler, viewPatch, Topic, TopicKind (+6 more)

### Community 62 - ".Router"
Cohesion: 0.29
Nodes (4): RequireAdmin(), TestHandlerErrors(), sameOriginOnly(), useValidateTag()

### Community 63 - "Part"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "App"
Cohesion: 0.03
Nodes (41): FromContext(), Campaign, browsePageSize(), compendiumBack(), compendiumPath(), ImportSummary(), pathCampaign(), viewerCampaign() (+33 more)

### Community 66 - "parseHP"
Cohesion: 0.22
Nodes (7): Amount, userText, CritDamage(), Damage(), parseRevive(), Set(), TestApplyHPDeath()

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.10
Nodes (40): CondOpKind, Effects, parseFeature(), parseSet(), parseSpell(), parseCoin(), parseInspiration(), parseItem() (+32 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 72 - "Condition"
Cohesion: 0.06
Nodes (42): Combatant, Condition, CreateConditionParams, ListCombatantsRow, combatantRow(), passTurn(), rollFor(), stepTurn() (+34 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "app"
Cohesion: 0.11
Nodes (17): env_file, environment, image, logging, network_mode, restart, user, volumes (+9 more)

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

### Community 85 - "TestTakeHomeDownload"
Cohesion: 0.27
Nodes (7): TestTakeHomeCarriesMadeAtHomeToken(), testEnv, hasBlock(), readZip(), TestTakeHomeDialog(), TestTakeHomeDownload(), takeHomeSeed

### Community 86 - "homeImportForm"
Cohesion: 0.38
Nodes (5): Props, Size, Variants(), homeImportForm(), importConfirmID()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".source"
Cohesion: 0.11
Nodes (19): DefaultRank(), TestDefaultRank(), ParsePageURL(), TestParsePageURL(), UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly() (+11 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "Client"
Cohesion: 0.18
Nodes (4): countingHandler, fakeHandler, Client, isDevice()

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.07
Nodes (50): eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack(), TestHandoffExpiresAndReplaces() (+42 more)

### Community 103 - "Session"
Cohesion: 0.04
Nodes (51): fixture, tokens, loopbackHostOnly(), main(), parseFlags(), run(), TestLoopbackHostOnly(), TestParseFlags() (+43 more)

### Community 105 - "cert.go"
Cohesion: 0.06
Nodes (24): keepAwake(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname() (+16 more)

### Community 106 - "NewClient"
Cohesion: 0.21
Nodes (21): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+13 more)

### Community 107 - "ftBook"
Cohesion: 0.19
Nodes (16): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftParts(), ftPrerequisite() (+8 more)

### Community 108 - "Book"
Cohesion: 0.10
Nodes (28): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+20 more)

### Community 111 - "planBringBack"
Cohesion: 0.11
Nodes (27): sheetMap, addedRemoved(), describe(), diffNames(), finish(), GroupDiff, GroupKeys(), GroupLabel() (+19 more)

### Community 112 - "mirror"
Cohesion: 0.16
Nodes (16): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), replaceDir(), TestMirrorGetMissingFile(), DownloadOpen5e() (+8 more)

### Community 113 - "starterEnv"
Cohesion: 0.30
Nodes (10): testEnv, pickerTicks(), serveStarterSets(), starterEnv(), TestImportsTickFromCampaignOnly(), TestPartialDownloadTicks(), TestPickerKeepsJobTicks(), TestSourcesCardStarterSets() (+2 more)

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 117 - "vars"
Cohesion: 0.22
Nodes (8): env, CGO_ENABLED, GOFLAGS, vars, APP, EXE, PKG, version

### Community 119 - ".q"
Cohesion: 0.08
Nodes (30): TestSourceDeleteMany(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+22 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Pending"
Cohesion: 0.14
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 127 - "task"
Cohesion: 0.22
Nodes (11): cmds, desc, env, cmds, desc, env, task, APP (+3 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "env"
Cohesion: 0.26
Nodes (15): BIN, DATA_DIR, MARKERS, MODE, NOREBUILD, PORT, PREPARE, RESTART_ARGS (+7 more)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eHabitats(), TestOpen5eMastery() (+16 more)

### Community 135 - "Queries"
Cohesion: 0.02
Nodes (166): HomebrewOwner, sourceFormat, HomeManifest, TakeHomeBaseCharacter, TakeHomeCampaign, TakeHomeFallback, TakeHomeLevelUp, TakeHomeManifest (+158 more)

### Community 136 - "alert_templ.go"
Cohesion: 0.09
Nodes (25): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+17 more)

### Community 137 - "parseMod"
Cohesion: 0.50
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 140 - ".Entry"
Cohesion: 0.10
Nodes (26): MonsterAction, addEntry(), setStat(), MonsterData, TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor(), InputForEntry() (+18 more)

### Community 143 - "run.sh"
Cohesion: 0.56
Nodes (8): alive(), lock(), run.sh script, start_server(), stop_server(), superseded(), unchanged(), unlock()

### Community 144 - ".saveReport"
Cohesion: 0.23
Nodes (6): firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify(), TestFolderName()

### Community 148 - "main"
Cohesion: 0.16
Nodes (14): die(), backup.sh script, DATA_DIR, lock(), common.sh script, build_and_check(), compose(), healthy() (+6 more)

### Community 175 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 179 - "dev:changed"
Cohesion: 0.40
Nodes (5): cmds, method, sources, exclude, dev:changed

### Community 180 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 182 - "systemd-units.sh"
Cohesion: 0.83
Nodes (3): die(), render(), systemd-units.sh script

### Community 185 - "Env"
Cohesion: 0.16
Nodes (6): featureSwitch, Pin, featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 186 - "internal/web"
Cohesion: 0.12
Nodes (14): EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Gotchas, Handlers, internal/web, TestKeepsBodyNewCharacter() (+6 more)

### Community 188 - "dev:build"
Cohesion: 0.40
Nodes (5): cmds, deps, requires, vars, dev:build

### Community 191 - "before"
Cohesion: 0.40
Nodes (5): cmds, desc, vars, before, BEFORE_BIN

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
Cohesion: 0.11
Nodes (18): cmds, cmds, desc, cmds, desc, cmds, internal, cmds (+10 more)

### Community 201 - "Input"
Cohesion: 0.08
Nodes (56): ctxKey, ctxState, Props, SearchReadableBlocksRow, Dialog(), Close(), CloseFor(), Content() (+48 more)

### Community 204 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 205 - "run:before"
Cohesion: 0.27
Nodes (6): cmds, desc, cmds, desc, run:before, run:test

### Community 209 - "dev:migrations"
Cohesion: 0.50
Nodes (4): cmds, method, sources, dev:migrations

### Community 210 - "run"
Cohesion: 0.67
Nodes (3): cmds, desc, run

### Community 211 - "run:test-player"
Cohesion: 0.67
Nodes (3): cmds, desc, run:test-player

## Knowledge Gaps
- **38 isolated node(s):** `npx`, `@playwright/mcp`, `DATA_DIR`, `gm-session-manager`, `testEnv` (+33 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 565 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **65 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Queries`, `perm.go`, `.saveReport`, `CompendiumPath`, `CompendiumKind`, `Encounter`, `handoff`, `Hub`, `testing.T`, `Sheet`, `CampaignSources`, `Topic`, `.Router`, `Condition`, `Session`, `cert.go`, `Book`, `.q`, `Pending`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `DATA_DIR` to the rest of the system?**
  _38 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `db/db.go` be split into smaller, more focused modules?**
  _Cohesion score 0.08677098150782361 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `db/db.go`, `context.Context`, `report`, `SheetView`, `Queries`, `alert_templ.go`, `perm.go`, `entityForm`, `CompendiumPath`, `CompendiumKind`, `github.com/a-h/templ.Attributes`, `CombatantView`, `TestBringBackAccept`, `Player`, `Input`, `tabs_templ.go`, `WizardView`, `.sendFeedback`, `deathFix`, `github.com/a-h/templ.Component`, `testing.T`, `Sheet`, `internal/web`, `CampaignSources`, `App`, `Condition`, `Input`, `ParseWiki`, `TestTakeHomeDownload`, `.source`, `newPlayerEnv`, `starterEnv`, `.q`?**
  _High betweenness centrality (0.042) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.006913834093683737 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `internal/web`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.023) - this node is a cross-community bridge._