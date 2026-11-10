# Graph Report - gm-session-manager  (2026-11-10)

## Corpus Check
- 296 files · ~580,368 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4164 nodes · 18093 edges · 150 communities (123 shown, 27 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1860 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `b356bb30`
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
- .loadLive
- Import
- select.js
- Command
- .Entry
- ImportView
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- tooltip_templ.go
- levelUpForm
- wikilink.js
- tabs_templ.go
- popover.js
- github.com/a-h/templ.Component
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- testing.T
- command.js
- dialog.js
- Client
- Queries
- tooltip.js
- github.com/gin-gonic/gin.Context
- ParseAidedd
- Book
- buildReview
- composite.js
- .serverError
- Feature
- ClassInfo
- Button
- State
- cell
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
- fromLookup
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- dialog_templ.go
- checkbox.js
- .fail
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .wizardCommand
- useFocus
- mark_others.js
- ImagePage
- Lint and Release Rules
- Run Script
- CampaignSources
- useTypeahead
- shadcn-templ.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- joinNonEmpty
- mounted
- NewClient
- Sheet
- Session
- parseLiveKey
- combobox_templ.go
- ftBook
- spellCounts
- .encounterView
- formErrors
- Parse
- Open
- TestClassOptionChoices
- Pending
- inputgroup_templ.go
- field_templ.go
- select_templ.go
- build.sh
- .Router
- popover_templ.go
- errBodyWriter
- .home
- buildReview
- CompendiumSource
- Before/after pairs
- empty_templ.go
- open5e.go
- DBTX
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- toast_templ.go
- Site
- alert_templ.go
- radiogroup_templ.go
- Props
- Props
- ParseWiki
- NoteBlock
- .ListCompendiumNames
- .ListSharedOrder
- Client
- ListSharedWithRow

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
- `UI conventions` --references--> `selectItems()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (150 total, 27 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (66): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+58 more)

### Community 1 - "CombatantView"
Cohesion: 0.08
Nodes (55): Defenses, Dialog(), Close(), Content(), Description(), Dialog(), Footer(), Header() (+47 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (98): keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListImageEntitiesRow, SearchReadableBlocksRow (+90 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (22): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+14 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.06
Nodes (23): DeathSaves, Part, Skill, SpellData, TestCustomMods(), skillByKey(), Sum(), TestDualWield() (+15 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (78): CountSharedBySourceRow, Condition, ArmorEffect(), Signed(), Variant, Action(), Card(), Content() (+70 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.25
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - ".loadLive"
Cohesion: 0.06
Nodes (69): CountHabitatsRow, Badge(), Content(), Description(), Empty(), Header(), Media(), Title() (+61 more)

### Community 10 - "Import"
Cohesion: 0.10
Nodes (35): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+27 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (66): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+58 more)

### Community 13 - ".Entry"
Cohesion: 0.07
Nodes (35): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, habitatsFrom(), ValidKind() (+27 more)

### Community 14 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows(), importStatusItems() (+2 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.08
Nodes (42): Note, GetBlockRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+34 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "tooltip_templ.go"
Cohesion: 0.11
Nodes (31): SearchCompendiumRow, Field, Content(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks() (+23 more)

### Community 21 - "levelUpForm"
Cohesion: 0.09
Nodes (46): RadioGroup(), namedParts(), abilitySelect(), LevelUpView, Option, SpellNeed, SubclassPanel, joinInts() (+38 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.20
Nodes (18): Content(), contentID(), IDFromContext(), initialValue(), listVariant(), listVariantClasses(), orientation(), state() (+10 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (86): ctxState, ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer() (+78 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.15
Nodes (27): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage() (+19 more)

### Community 28 - "App"
Cohesion: 0.06
Nodes (24): ftsQuery(), Search(), Campaign, compendiumBack(), compendiumPath(), ImportSummary(), pathCampaign(), viewerCampaign() (+16 more)

### Community 29 - "newTestEnv"
Cohesion: 0.05
Nodes (63): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait() (+55 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "database/sql.NullInt64"
Cohesion: 0.06
Nodes (32): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CreateCombatantParams, saveScreenshot(), nullID(), ensureCategory() (+24 more)

### Community 34 - "testing.T"
Cohesion: 0.03
Nodes (84): CoreTraits(), TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), pageSource(), PageSourceKey(), roundTrip(), TestParseWikidotClass() (+76 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Client"
Cohesion: 0.21
Nodes (3): ClientInfo, Hub, randomID()

### Community 38 - "Queries"
Cohesion: 0.07
Nodes (37): dataFor(), DecodeData(), Kind, ListSubclassesOfRow, Queries, CompendiumEntry, TestSlotMaxima(), ClassLevel (+29 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (23): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), featureCommand(), coinCommand(), commandText() (+15 more)

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - "Book"
Cohesion: 0.18
Nodes (16): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+8 more)

### Community 43 - "buildReview"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".serverError"
Cohesion: 0.07
Nodes (13): Entity, Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks() (+5 more)

### Community 46 - "Feature"
Cohesion: 0.19
Nodes (8): FeatureEffects, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 47 - "ClassInfo"
Cohesion: 0.08
Nodes (20): featureSwitch, LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+12 more)

### Community 48 - "Button"
Cohesion: 0.05
Nodes (102): UI conventions, Button(), Label(), Scripts(), Separator(), List(), Tabs(), Trigger() (+94 more)

### Community 49 - "State"
Cohesion: 0.11
Nodes (13): event, Handler, message, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+5 more)

### Community 50 - "cell"
Cohesion: 0.12
Nodes (19): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+11 more)

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
Cohesion: 0.12
Nodes (15): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+7 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ClassData"
Cohesion: 0.15
Nodes (18): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+10 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "ftSpell"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 59 - "TestDownloadFiveTools"
Cohesion: 0.13
Nodes (19): internal/compendium, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestParseFiveToolsBook() (+11 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Input"
Cohesion: 0.05
Nodes (73): Coins, CloseFor(), Description(), Field(), Group(), Label(), Input(), Content() (+65 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "fromLookup"
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.07
Nodes (38): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+30 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "dialog_templ.go"
Cohesion: 0.18
Nodes (11): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+3 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - ".fail"
Cohesion: 0.11
Nodes (6): Character, CanCharacter(), linkEntityLook(), linkCharacterLook(), lookFree(), nextStep()

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

### Community 80 - ".wizardCommand"
Cohesion: 0.08
Nodes (31): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+23 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "ImagePage"
Cohesion: 0.11
Nodes (29): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Look, initials(), currentLook(), EntityDetail(), entityForm() (+21 more)

### Community 86 - "CampaignSources"
Cohesion: 0.17
Nodes (17): campaignSourceRow(), CampaignSources(), CampaignSourcesView, sourceGroupButton(), sourceRowTitle(), CampaignOverview(), CampaignRoster(), Roster (+9 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "shadcn-templ.go"
Cohesion: 0.09
Nodes (19): Props, initialChecked(), Direction(), flattenClasses(), lastTopLevelColon(), MergeAttributes(), Ptr(), twImportantToV3() (+11 more)

### Community 100 - "joinNonEmpty"
Cohesion: 0.08
Nodes (38): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+30 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Sheet"
Cohesion: 0.07
Nodes (51): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, TextData (+43 more)

### Community 104 - "Session"
Cohesion: 0.12
Nodes (14): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+6 more)

### Community 106 - "combobox_templ.go"
Cohesion: 0.11
Nodes (22): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+14 more)

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (17): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+9 more)

### Community 108 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 109 - ".encounterView"
Cohesion: 0.10
Nodes (19): Defense, Combatant, Encounter, ListCombatantsRow, Character, ParseDefenses(), combatantDefenses(), damageChoice() (+11 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 111 - "Parse"
Cohesion: 0.09
Nodes (23): feature(), Parse(), tokenize(), fighter(), TestCommands(), TestHealingFromZeroResetsDeathSaves(), TestMasteryCommand(), TestParseRawArgs() (+15 more)

### Community 112 - "Open"
Cohesion: 0.14
Nodes (21): fixture, migration, newFixture(), TestOnlyHashIsStored(), apply(), loadMigrations(), Migrate(), Open() (+13 more)

### Community 113 - "TestClassOptionChoices"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 114 - "Pending"
Cohesion: 0.29
Nodes (4): PendingState, Pending, PendingKind, viewerCtx

### Community 115 - "inputgroup_templ.go"
Cohesion: 0.17
Nodes (13): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+5 more)

### Community 116 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 117 - "select_templ.go"
Cohesion: 0.15
Nodes (13): initialValue(), Align, ContentProps, ctxKey, ctxState, GroupProps, ItemProps, LabelProps (+5 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".Router"
Cohesion: 0.19
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 123 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps (+3 more)

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - ".home"
Cohesion: 0.20
Nodes (4): firstLine(), mdLinkText(), joinPath(), JoinQRSrc()

### Community 126 - "buildReview"
Cohesion: 0.17
Nodes (7): buildReview(), csvSource(), formValues(), parseCSV(), capReader, csvImport, csvImports

### Community 127 - "CompendiumSource"
Cohesion: 0.22
Nodes (7): ListEntryCopiesRow, CompendiumSource, ListEntryCopiesParams, CompareCopy, CompareView, EntryView, HabitatCopy

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "empty_templ.go"
Cohesion: 0.28
Nodes (8): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, mediaClasses()

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (22): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eMastery(), TestUnflattenTable() (+14 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.14
Nodes (13): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+5 more)

### Community 138 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 139 - "Site"
Cohesion: 0.18
Nodes (8): Site, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SavePage(), SiteByKey(), TestParsePageURL()

### Community 140 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 141 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 142 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 143 - "Props"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 144 - "ParseWiki"
Cohesion: 0.21
Nodes (12): saveFieldLinks(), wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki() (+4 more)

### Community 145 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 159 - "Client"
Cohesion: 0.16
Nodes (4): fakeHandler, IsLoopback(), TestIsLoopback(), Client

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 387 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **27 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Weapon`, `.loadLive`, `perm.go`, `database/sql.NullInt64`, `Client`, `Queries`, `github.com/gin-gonic/gin.Context`, `.serverError`, `State`, `cert.go`, `FiveToolsJob`, `.fail`, `.wizardCommand`, `Sheet`, `parseLiveKey`, `.encounterView`, `.Router`, `.home`, `buildReview`?**
  _High betweenness centrality (0.045) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10433563416738567 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Weapon` to `compendium_templ.go`, `testing.T`, `joinNonEmpty`, `Queries`, `Sheet`, `SheetView`, `Command`, `.encounterView`, `Feature`, `ClassInfo`, `.wizardCommand`, `TestClassOptionChoices`, `levelUpForm`, `ParseSheet`, `Input`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.08397271952259165 - nodes in this community are weakly interconnected._
- **Why does `Page` connect `Site` to `compendium_templ.go`, `testing.T`, `App`, `open5e.go`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.0077440792993720255 - nodes in this community are weakly interconnected._