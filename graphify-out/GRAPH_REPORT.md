# Graph Report - gm-session-manager  (2026-11-28)

## Corpus Check
- 437 files · ~753,623 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 65 file(s) not represented in the graph (top: .templ 51, (none) 4, .css 4)

## Summary
- 5656 nodes · 30195 edges · 199 communities (136 shown, 63 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2983 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `aed3185a`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Sheet
- context.Context
- takehome_test.go
- createProvider
- report
- SheetView
- combobox.js
- CompendiumKind
- Book
- encoding/json.RawMessage
- select.js
- Queries
- CampaignSources
- Session
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Item
- github.com/a-h/templ.Component
- wikilink.js
- parseHP
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Sum
- CombatantView
- TestBringBackAccept
- app.js
- tabbable.js
- useDismiss
- TestNewCharacterFromGM
- Encounter
- command.js
- dialog.js
- nullID
- ParseWikidot
- tooltip.js
- builderFields
- tabs_templ.go
- IfElse
- Client
- composite.js
- popover_templ.go
- Feature
- vendorjs/main.go
- CompendiumPath
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- DownloadFiveTools
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- ParseArmorTraining
- config.go
- radiogroup.js
- Store
- .Router
- Part
- sqlc Config
- App
- handoff
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- .combatantHP
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
- Client
- redirect.go
- useTypeahead
- TestSourceSwitchKeepsEdits
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- who
- mounted
- newPlayerEnv
- ParseSheet
- profFocus
- cert.go
- NewClient
- ftBook
- run
- BlockEditor
- Topic
- pause_windows.go
- normalize_test.go
- buildReview
- vars
- net.IP
- .sendFeedback
- looseString
- errBodyWriter
- Pending
- task
- Before/after pairs
- env
- jsconfig.json
- open5e.go
- testing.M
- appState
- names
- ASIError
- time.Time
- HomebrewInputFor
- sourceUsedError
- run.sh
- WeaponKind
- planBringBack
- main
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- Running on a server (Debian, Docker)
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
- .ListBlockStepsSince
- .ListImageLookAccess
- Browser check
- dev:changed
- release.sh
- internal/web
- IMPORT_FILTER_CAT_PREFIX
- planSwitch
- .ListItemsByCategory
- .ListSpellsForClass
- dev:build
- .ListSubracesOf
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

## God Nodes (most connected - your core abstractions)
1. `App` - 454 edges
2. `Path()` - 341 edges
3. `newTestEnv()` - 199 edges
4. `Button()` - 145 edges
5. `Sheet` - 141 edges
6. `CN()` - 114 edges
7. `Queries` - 89 edges
8. `SheetView` - 80 edges
9. `IfElse()` - 68 edges
10. `ParseSheet()` - 64 edges

## Surprising Connections (you probably didn't know these)
- `Notes` --references--> `linkClick()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/assets/js/wikilink.js
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
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

## Communities (199 total, 63 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (58): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+50 more)

### Community 1 - "Sheet"
Cohesion: 0.04
Nodes (106): LevelUpChoice, internal/compendium, Cantrips(), CoreTraits(), EquipmentOptions(), firstTable(), LevelSpells(), OptionPrerequisite() (+98 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (155): keepAwake(), bindFiveToolsImages(), ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Condition (+147 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.17
Nodes (21): readLimited(), ReadTakeHome(), rewriteZip(), sampleBringBack(), sampleFallback(), sampleNewCharacter(), sampleTakeHome(), TestBringBackFallbackRoundTrip() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "report"
Cohesion: 0.12
Nodes (19): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+11 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (102): UsesLabel(), Slot, Signed(), Variant, Action(), Card(), Content(), Description() (+94 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "CompendiumKind"
Cohesion: 0.04
Nodes (51): EquipmentItem, EquipmentOption, dataFor(), Decode(), DecodeData(), SpellData, ValidKind(), CountCompendiumRow (+43 more)

### Community 9 - "Book"
Cohesion: 0.10
Nodes (29): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+21 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (35): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+27 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Queries"
Cohesion: 0.03
Nodes (128): HomeManifest, TakeHomeBaseCharacter, TakeHomeCampaign, TakeHomeFallback, TakeHomeLevelUp, TakeHomeManifest, ListCharacterImagesRow, Queries (+120 more)

### Community 13 - "CampaignSources"
Cohesion: 0.08
Nodes (30): campaignSourceRow(), CampaignSources(), CampaignSourcesLive(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), CampaignOverview(), CampaignRoster() (+22 more)

### Community 14 - "Session"
Cohesion: 0.05
Nodes (53): fixture, tokens, loopbackHostOnly(), TestLoopbackHostOnly(), constantTimeEqual(), credentialsMatch(), isCleared(), newFixture() (+45 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.04
Nodes (70): ListCharacterNamesRow, ListCombatantsRow, Visibility, internal/perm, Fields, ParseFields(), VisibleEntity(), VisibleLooks() (+62 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 21 - "github.com/a-h/templ.Component"
Cohesion: 0.03
Nodes (168): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListImageCharactersRow (+160 more)

### Community 22 - "wikilink.js"
Cohesion: 0.17
Nodes (18): nodeActions(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages, LAYOUTS (+10 more)

### Community 23 - "parseHP"
Cohesion: 0.16
Nodes (13): Amount, userText, CritDamage(), Damage(), Heal(), parseDamage(), parseHP(), parseRevive() (+5 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (187): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+179 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Sum"
Cohesion: 0.15
Nodes (8): TestCustomMods(), TestExplainSums(), ArmorEffect(), Sum(), TestExplainDamage(), WornArmor, TestWithEffects(), armorDetail()

### Community 28 - "CombatantView"
Cohesion: 0.10
Nodes (38): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+30 more)

### Community 29 - "TestBringBackAccept"
Cohesion: 0.26
Nodes (13): TakeHome, TestNewCharacterTableSources(), testEnv, homeSheet(), TestBringBackAccept(), TestBringBackPlayerUpload(), TestBringBackReject(), TestBringBackSources() (+5 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (23): checkLink(), checkLinks(), clearDisplay(), clearSourceSelection(), commandChoices(), commandHint(), commandState(), createItem() (+15 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "TestNewCharacterFromGM"
Cohesion: 0.20
Nodes (11): testEnv, mentionDoc(), newExtrasEnv(), TestBringBackAcceptExtras(), TestBringBackAcceptFiles(), TestBringBackStaleExtras(), TestNewCharacterFromGM(), TestNewCharacterFromVisitor() (+3 more)

### Community 34 - "Encounter"
Cohesion: 0.20
Nodes (8): CanAddSources(), notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "nullID"
Cohesion: 0.06
Nodes (65): DeathSaves, Player(), TestCanAddSources(), TestCanRemoveCondition(), nullID(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv (+57 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (23): block, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "builderFields"
Cohesion: 0.05
Nodes (75): ImageCategory, Coins, Defenses, Uses, Description(), Field(), Group(), Label() (+67 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.10
Nodes (41): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+33 more)

### Community 42 - "IfElse"
Cohesion: 0.05
Nodes (71): Props, SpellColumn(), Alert(), Description(), Title(), Checkbox(), initialChecked(), Label() (+63 more)

### Community 43 - "Client"
Cohesion: 0.22
Nodes (4): ClientInfo, goneDevice, Hub, isDevice()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "popover_templ.go"
Cohesion: 0.13
Nodes (27): UI conventions, Content(), Description(), Header(), Popover(), state(), Title(), Trigger() (+19 more)

### Community 46 - "Feature"
Cohesion: 0.08
Nodes (20): FeatureEffects, findMod(), newMod(), parseMod(), CustomMod, IsModTarget(), indexAbilityName(), SpellMods() (+12 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 48 - "CompendiumPath"
Cohesion: 0.03
Nodes (125): CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, ListSharedWithRow, SearchCompendiumRow (+117 more)

### Community 49 - "Path"
Cohesion: 0.03
Nodes (133): ParsePageURL(), TestParsePageURL(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestCombatantNamesFollowSource() (+125 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (171): TestNamesMatchGo(), Defense, TestDecode(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), replaceDir(), TestFiveToolsBooks() (+163 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (25): ClassFeature, ClassData, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition() (+17 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.08
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 59 - "config.go"
Cohesion: 0.22
Nodes (12): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), Player(), lookupMap() (+4 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Store"
Cohesion: 0.08
Nodes (31): migration, internal/db, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp() (+23 more)

### Community 62 - ".Router"
Cohesion: 0.33
Nodes (3): RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 63 - "Part"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "App"
Cohesion: 0.03
Nodes (32): FromContext(), Campaign, Entity, downloadNames(), entityTopics(), linkEntityLook(), inCampaign(), collectHeld() (+24 more)

### Community 66 - "handoff"
Cohesion: 0.23
Nodes (6): IsLoopback(), TestIsLoopback(), closeServers(), removeFile(), shutdownServers(), handoff

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.08
Nodes (46): CondOpKind, Effects, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn(), parseCoin() (+38 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 72 - ".combatantHP"
Cohesion: 0.32
Nodes (7): Combatant, combatantRow(), damageHP(), healHP(), monsterHP(), setHP(), hpChange

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.13
Nodes (26): abilityMod(), addEntry(), ParseAidedd(), setStat(), statLabel(), textLines(), MonsterData, firstInt() (+18 more)

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
Cohesion: 0.10
Nodes (8): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, openWeb(), randomID(), KEEP_IMAGES, LINK_PATHS

### Community 85 - "Client"
Cohesion: 0.09
Nodes (12): countingHandler, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch, Client (+4 more)

### Community 86 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.33
Nodes (6): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits()

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "who"
Cohesion: 0.04
Nodes (50): HomebrewOwner, Site, Source, sourceFormat, linkTarget, TestHomebrew(), fiveToolsPage(), ftsQuery() (+42 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.07
Nodes (48): TestEffectFields(), eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack() (+40 more)

### Community 103 - "ParseSheet"
Cohesion: 0.07
Nodes (54): Import(), TestUpgradeSheet(), ParseSheet(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkImportReview(), homeTableSheet(), TestCompendiumHabitatFilter() (+46 more)

### Community 105 - "cert.go"
Cohesion: 0.15
Nodes (10): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+2 more)

### Community 106 - "NewClient"
Cohesion: 0.21
Nodes (21): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+13 more)

### Community 107 - "ftBook"
Cohesion: 0.11
Nodes (34): ftBook, ftRec, each(), ftCastingTime(), ftCell(), ftCopyTargets(), ftDecode(), ftFeatIncrease() (+26 more)

### Community 108 - "run"
Cohesion: 0.25
Nodes (7): main(), parseFlags(), run(), TestParseFlags(), bookFolder(), CleanDownloads(), TestCleanDownloads()

### Community 109 - "BlockEditor"
Cohesion: 0.22
Nodes (5): attachedPicture(), BlockEditor, json(), links(), sync()

### Community 111 - "Topic"
Cohesion: 0.29
Nodes (5): event, Topic, TopicKind, noteTopics(), heldTopics

### Community 113 - "normalize_test.go"
Cohesion: 0.33
Nodes (3): openBrowser(), normalize(), TestNormalize()

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (27): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+19 more)

### Community 117 - "vars"
Cohesion: 0.22
Nodes (8): env, CGO_ENABLED, GOFLAGS, vars, APP, EXE, PKG, version

### Community 118 - "net.IP"
Cohesion: 0.52
Nodes (5): LANAddresses(), PreferDefaultRoute(), preferFirst(), TestPreferFirst(), handoffIPs()

### Community 119 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Pending"
Cohesion: 0.15
Nodes (4): PendingState, randomID(), Pending, PendingKind

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
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery() (+16 more)

### Community 135 - "appState"
Cohesion: 0.47
Nodes (5): loadSettings(), saveSettings(), TestSettingsFile(), appState, settings

### Community 136 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 139 - "time.Time"
Cohesion: 0.06
Nodes (18): Attachment, browserName(), firstLine(), mdLinkText(), saveScreenshot(), TestUserAgent(), userAgentSummary(), saveSniffed() (+10 more)

### Community 140 - "HomebrewInputFor"
Cohesion: 0.11
Nodes (18): MonsterAction, habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), HomebrewInputFor(), InputForEntry(), actionsText() (+10 more)

### Community 143 - "run.sh"
Cohesion: 0.56
Nodes (8): alive(), lock(), run.sh script, start_server(), stop_server(), superseded(), unchanged(), unlock()

### Community 144 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 145 - "planBringBack"
Cohesion: 0.11
Nodes (27): sheetMap, addedRemoved(), describe(), diffNames(), finish(), GroupDiff, GroupKeys(), GroupLabel() (+19 more)

### Community 148 - "main"
Cohesion: 0.16
Nodes (14): die(), backup.sh script, DATA_DIR, lock(), common.sh script, build_and_check(), compose(), healthy() (+6 more)

### Community 157 - "Running on a server (Debian, Docker)"
Cohesion: 0.22
Nodes (8): Backups, Day to day, Folders, Own account, Rehearse, Rollback, Running on a server (Debian, Docker), Setup

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 179 - "dev:changed"
Cohesion: 0.40
Nodes (5): cmds, method, sources, exclude, dev:changed

### Community 180 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 181 - "internal/web"
Cohesion: 0.14
Nodes (11): Gotchas, Handlers, internal/web, Notes, TestKeepsBodyNewCharacter(), TestKeepsBodyBringBack(), TestPagesRender(), keepsBody() (+3 more)

### Community 184 - "planSwitch"
Cohesion: 0.15
Nodes (8): featureSwitch, parseFeature(), parseFeatureSet(), parseUses(), featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 188 - "dev:build"
Cohesion: 0.40
Nodes (5): cmds, deps, requires, vars, dev:build

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
Cohesion: 0.09
Nodes (53): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+45 more)

## Knowledge Gaps
- **38 isolated node(s):** `npx`, `@playwright/mcp`, `DATA_DIR`, `gm-session-manager`, `testEnv` (+33 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 556 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **63 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Sheet`, `context.Context`, `appState`, `CompendiumKind`, `time.Time`, `Queries`, `CampaignSources`, `Session`, `perm.go`, `Encounter`, `Client`, `CompendiumPath`, `Path`, `testing.T`, `.Router`, `handoff`, `.combatantHP`, `Client`, `who`, `Topic`, `net.IP`, `Pending`?**
  _High betweenness centrality (0.029) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `DATA_DIR` to the rest of the system?**
  _38 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Sheet` be split into smaller, more focused modules?**
  _Cohesion score 0.03627022467602178 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Sheet`, `context.Context`, `report`, `SheetView`, `Queries`, `CampaignSources`, `perm.go`, `github.com/a-h/templ.Component`, `github.com/a-h/templ.Attributes`, `CombatantView`, `TestBringBackAccept`, `TestNewCharacterFromGM`, `nullID`, `builderFields`, `tabs_templ.go`, `IfElse`, `popover_templ.go`, `CompendiumPath`, `testing.T`, `internal/web`, `Store`, `App`, `github.com/a-h/templ.Component`, `ParseWiki`, `TestSourceSwitchKeepsEdits`, `who`, `newPlayerEnv`, `ParseSheet`, `.sendFeedback`?**
  _High betweenness centrality (0.028) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.007301247771836007 - nodes in this community are weakly interconnected._
- **Why does `Websocket and live views` connect `nullID` to `builderFields`, `live.js`, `github.com/a-h/templ.Component`, `internal/web`?**
  _High betweenness centrality (0.019) - this node is a cross-community bridge._