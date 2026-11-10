# Graph Report - gm-session-manager  (2026-11-10)

## Corpus Check
- 296 files · ~580,211 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4163 nodes · 18092 edges · 141 communities (111 shown, 30 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1860 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `27639bc2`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- context.Context
- .do
- createProvider
- Weapon
- SheetView
- combobox.js
- redirect.go
- Button
- Import
- select.js
- Command
- .Entry
- select_templ.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- CampaignSourcesView
- levelUpForm
- wikilink.js
- tabs_templ.go
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- Item
- testing.T
- command.js
- dialog.js
- Client
- Queries
- tooltip.js
- .whoami
- ParseAidedd
- Book
- buildReview
- composite.js
- who
- Feature
- ClassInfo
- github.com/a-h/templ.Component
- State
- NameDesc
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- ftSpell
- TestDownloadFiveTools
- radiogroup.js
- vendorjs/main.go
- Input
- FiveToolsJob
- sqlc Config
- time.Time
- Browser check
- Claude Settings
- ImportView
- use_transition_status.js
- playwright
- release.sh
- Parse
- checkbox.js
- planSwitch
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .sendFeedback
- useFocus
- mark_others.js
- WeaponKind
- Lint and Release Rules
- Run Script
- EquipmentOption
- useTypeahead
- toast_templ.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- joinNonEmpty
- mounted
- NewClient
- Sheet
- Session
- popover_templ.go
- .sourceIssues
- ftBook
- Props
- database/sql.NullInt64
- formErrors
- rowFigure
- Open
- TestClassOptionChoices
- Pending
- IsLoopback
- .ListCompendiumNames
- CN
- build.sh
- .ListEntrySourcesBySlugs
- errBodyWriter
- capReader
- Before/after pairs
- open5e.go
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- .ListCompendiumBySlugs
- Site
- .ListSpellsForClass
- ParseWiki
- Client
- ListSharedWithRow
- .ListCompendiumByKind
- .ListSharedOrder

## God Nodes (most connected - your core abstractions)
1. `App` - 313 edges
2. `newTestEnv()` - 124 edges
3. `Button()` - 121 edges
4. `Sheet` - 118 edges
5. `CN()` - 113 edges
6. `IfElse()` - 61 edges
7. `SheetView` - 57 edges
8. `Command` - 49 edges
9. `Input()` - 47 edges
10. `Queries` - 44 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `UI conventions` --references--> `removeButton()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (141 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (66): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+58 more)

### Community 1 - "CombatantView"
Cohesion: 0.09
Nodes (34): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+26 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (89): keepAwake(), ftsQuery(), Search(), SeedHabitats(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, SearchCompendiumRow (+81 more)

### Community 3 - ".do"
Cohesion: 0.15
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.10
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (85): ArmorEffect(), Signed(), Field, Card(), Content(), TriggerFor(), Content(), state() (+77 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Button"
Cohesion: 0.05
Nodes (93): ListEntityFieldBacklinksRow, ListBacklinksRow, UI conventions, Button(), Scripts(), Separator(), campaignLink(), CampaignOverview() (+85 more)

### Community 10 - "Import"
Cohesion: 0.09
Nodes (37): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+29 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.09
Nodes (57): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+49 more)

### Community 13 - ".Entry"
Cohesion: 0.07
Nodes (34): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), ValidKind(), TestHabitatsFrom() (+26 more)

### Community 14 - "select_templ.go"
Cohesion: 0.11
Nodes (22): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state() (+14 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (58): GetBlockRow, ListCombatantsRow, Entity, Fields, Look, ParseFields(), VisibleEntity(), VisibleLooks() (+50 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "CampaignSourcesView"
Cohesion: 0.33
Nodes (5): CampaignSourcesView, sourceGroupButton(), sourceRowTitle(), CampaignSource, sourceGroup

### Community 21 - "levelUpForm"
Cohesion: 0.08
Nodes (54): Props, Footer(), Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup() (+46 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.13
Nodes (32): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+24 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (127): Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props (+119 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.13
Nodes (30): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage() (+22 more)

### Community 28 - "App"
Cohesion: 0.05
Nodes (27): TopicKind, SetCookie(), Topic, coinCommand(), commandText(), hpCommand(), parseCommand(), homebrewTabURL() (+19 more)

### Community 29 - "newTestEnv"
Cohesion: 0.05
Nodes (53): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait() (+45 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Item"
Cohesion: 0.15
Nodes (7): findItem(), parseArmor(), TestCustomMods(), Sum(), TestGearCommands(), ensureItem(), Item

### Community 34 - "testing.T"
Cohesion: 0.04
Nodes (66): RemoveSourceDir(), TestRemoveSourceDir(), TestCompletions(), TestLevelUp(), TestLevelUpSubclassAndPools(), TestLevelUpThirdCaster(), TestMasteryLimit(), TestDedupeFeatures() (+58 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Client"
Cohesion: 0.20
Nodes (3): ClientInfo, event, Hub

### Community 38 - "Queries"
Cohesion: 0.04
Nodes (66): dataFor(), DecodeData(), Kind, TextData, ListEntryCopiesRow, ListSubclassesOfRow, Queries, Character (+58 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".whoami"
Cohesion: 0.05
Nodes (19): ClearCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), ListAttachmentFilesRow, Attachment, browserName() (+11 more)

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - "Book"
Cohesion: 0.19
Nodes (15): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+7 more)

### Community 43 - "buildReview"
Cohesion: 0.06
Nodes (46): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+38 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "who"
Cohesion: 0.07
Nodes (16): CountCompendiumRow, Campaign, compendiumBack(), compendiumPath(), pathCampaign(), viewerCampaign(), imageItem(), imageVersion() (+8 more)

### Community 46 - "Feature"
Cohesion: 0.19
Nodes (7): FeatureEffects, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, usesValue()

### Community 47 - "ClassInfo"
Cohesion: 0.09
Nodes (16): LevelUpChoice, dnd5eFeatures(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+8 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (113): CountHabitatsRow, ListMonstersByNameRow, Badge(), Variant, Action(), Description(), Header(), Title() (+105 more)

### Community 49 - "State"
Cohesion: 0.14
Nodes (9): Handler, patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff() (+1 more)

### Community 50 - "NameDesc"
Cohesion: 0.18
Nodes (19): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, SpellData, optionsGrid() (+11 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "cert.go"
Cohesion: 0.11
Nodes (23): main(), run(), Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap() (+15 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ClassData"
Cohesion: 0.12
Nodes (24): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, cellAt(), classTable(), featuresUnder() (+16 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "ftSpell"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 59 - "TestDownloadFiveTools"
Cohesion: 0.12
Nodes (21): internal/compendium, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestImportFiveToolsBook() (+13 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Input"
Cohesion: 0.05
Nodes (99): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+91 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.25
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "time.Time"
Cohesion: 0.11
Nodes (13): ListImagesRow, FiveToolsView, NewFiveToolsView(), BlockStep, LevelUp, ListImagesParams, ListPlayerSessionsByCampaignRow, FiveToolsBook (+5 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "ImportView"
Cohesion: 0.18
Nodes (14): ImageCategory, categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows() (+6 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Parse"
Cohesion: 0.09
Nodes (22): Defense, feature(), Parse(), tokenize(), fighter(), TestCommands(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+14 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "planSwitch"
Cohesion: 0.25
Nodes (5): featureSwitch, featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.11
Nodes (20): Player(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips() (+12 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "WeaponKind"
Cohesion: 0.40
Nodes (5): Completions(), parseMastery(), TestWeaponKind(), WeaponKind(), catchUpMasteries()

### Community 86 - "EquipmentOption"
Cohesion: 0.25
Nodes (8): EquipmentItem, EquipmentOption, WornArmor, TestOptionText(), EquipmentChoice, optionText(), applyEquipment(), gear

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 100 - "joinNonEmpty"
Cohesion: 0.08
Nodes (34): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+26 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Sheet"
Cohesion: 0.06
Nodes (50): DeathSaves, Cantrips(), EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+42 more)

### Community 104 - "Session"
Cohesion: 0.11
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+10 more)

### Community 105 - "popover_templ.go"
Cohesion: 0.16
Nodes (22): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+14 more)

### Community 106 - ".sourceIssues"
Cohesion: 0.43
Nodes (6): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (23): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+15 more)

### Community 108 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 109 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (40): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Combatant, Condition, CreateCombatantParams, Character, applyCondOps() (+32 more)

### Community 110 - "formErrors"
Cohesion: 0.40
Nodes (3): fieldError, pickError, userError

### Community 111 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 112 - "Open"
Cohesion: 0.11
Nodes (22): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+14 more)

### Community 113 - "TestClassOptionChoices"
Cohesion: 0.15
Nodes (14): UpsertCompendiumEntryParams, SwitchFeature, SwitchPlan, PlanSourceSwitch(), names(), TestPlanSourceSwitch(), classCopy(), testEnv (+6 more)

### Community 114 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 117 - "CN"
Cohesion: 0.06
Nodes (66): Props, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps (+58 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.12
Nodes (16): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ItemData, ftArmor(), ftCode() (+8 more)

### Community 139 - "Site"
Cohesion: 0.13
Nodes (14): Site, Source, fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page, pageSource() (+6 more)

### Community 144 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 159 - "Client"
Cohesion: 0.19
Nodes (3): fakeHandler, message, Client

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 386 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `time.Time`, `context.Context`, `.do`, `Client`, `Queries`, `Sheet`, `.whoami`, `Import`, `.sourceIssues`, `who`, `database/sql.NullInt64`, `perm.go`, `Open`, `State`, `cert.go`, `FiveToolsJob`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09951257475219137 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `context.Context`, `Weapon`, `SheetView`, `perm.go`, `levelUpForm`, `ParseSheet`, `Item`, `testing.T`, `Queries`, `Feature`, `ClassInfo`, `github.com/a-h/templ.Component`, `NameDesc`, `planSwitch`, `WeaponKind`, `EquipmentOption`, `joinNonEmpty`, `.sourceIssues`, `database/sql.NullInt64`, `TestClassOptionChoices`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.08599290780141844 - nodes in this community are weakly interconnected._
- **Why does `Page` connect `Site` to `compendium_templ.go`, `who`, `open5e.go`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009475792033183151 - nodes in this community are weakly interconnected._