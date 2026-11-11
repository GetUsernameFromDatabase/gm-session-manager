# Graph Report - gm-session-manager  (2026-11-11)

## Corpus Check
- 299 files · ~588,917 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4215 nodes · 18474 edges · 142 communities (113 shown, 29 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1900 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `d6336b5b`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- dialog_templ.go
- context.Context
- .do
- createProvider
- Sheet
- SheetView
- combobox.js
- redirect.go
- ImageResults
- Store
- select.js
- Command
- .Entry
- github.com/gin-gonic/gin.Context
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- DecodeData
- levelUpForm
- wikilink.js
- IfElse
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- newTestEnv
- .serverError
- testing.T
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- joinNonEmpty
- command.js
- dialog.js
- Client
- Queries
- tooltip.js
- .fail
- ParseWikidot
- Book
- buildReview
- composite.js
- Parse
- Feature
- IfElse
- .loadLive
- Topic
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- CommandBuilder
- golangci-lint Config
- .changed
- ItemData
- radiogroup.js
- vendorjs/main.go
- CN
- FiveToolsJob
- sqlc Config
- fromLookup
- Browser check
- Claude Settings
- App
- use_transition_status.js
- playwright
- release.sh
- Session
- checkbox.js
- ParseAidedd
- switch.js
- nullID
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- alert_templ.go
- Lint and Release Rules
- Run Script
- .Router
- useTypeahead
- Sheet
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- joinNonEmpty
- mounted
- NewClient
- TestLiveUpdatesAndKick
- .source
- DBTX
- .sendFeedback
- ftBook
- names
- writeBook
- formErrors
- .joinCampaigns
- .CountSharedByKind
- .ListBlockStepsSince
- Pending
- Encounter
- ftItemData
- github.com/a-h/templ.Component
- build.sh
- .sourceIssues
- NoteBlock
- errBodyWriter
- .ListEntrySourcesBySlugs
- .failBack
- parseLiveKey
- Before/after pairs
- .ListSpellsForClass
- open5e.go
- spellCounts
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- .ListSharedOrder
- Site
- IsLoopback
- TestRunCommandUsesCurrentSheet

## God Nodes (most connected - your core abstractions)
1. `App` - 320 edges
2. `newTestEnv()` - 129 edges
3. `Button()` - 122 edges
4. `Sheet` - 118 edges
5. `CN()` - 113 edges
6. `IfElse()` - 62 edges
7. `SheetView` - 57 edges
8. `Command` - 49 edges
9. `Input()` - 47 edges
10. `Slugify()` - 44 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `UI conventions` --references--> `selectItems()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go
- `UI conventions` --references--> `formSelect()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (142 total, 29 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (65): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+57 more)

### Community 1 - "dialog_templ.go"
Cohesion: 0.10
Nodes (35): Defenses, damageChoice(), capitalized(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState() (+27 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (92): keepAwake(), ListImageCharactersRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListSharedWithRow, ListImageEntitiesRow, SearchReadableBlocksRow (+84 more)

### Community 3 - ".do"
Cohesion: 0.15
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Sheet"
Cohesion: 0.07
Nodes (17): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), skillByKey(), Sum() (+9 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (115): ArmorEffect(), Coins, Slot, Signed(), Field, Button(), Variant, Card() (+107 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "ImageResults"
Cohesion: 0.10
Nodes (32): ListEntityFieldBacklinksRow, ListBacklinksRow, attrString(), generateSVG(), getIconContent(), Props, Icon(), EntityDetail() (+24 more)

### Community 10 - "Store"
Cohesion: 0.08
Nodes (44): migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+36 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (70): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+62 more)

### Community 13 - ".Entry"
Cohesion: 0.07
Nodes (33): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), TestTokensReadAsLabels() (+25 more)

### Community 14 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (20): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Campaign, coinCommand(), commandText() (+12 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (46): Note, GetBlockRow, Entity, Fields, Look, ParseFields(), VisibleEntity(), VisibleLooks() (+38 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "DecodeData"
Cohesion: 0.04
Nodes (53): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions() (+45 more)

### Community 21 - "levelUpForm"
Cohesion: 0.08
Nodes (55): LevelUpPlan, Checkbox(), Label(), initialValue(), Item(), RadioGroup(), state(), CompendiumURL() (+47 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "IfElse"
Cohesion: 0.21
Nodes (20): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+12 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (128): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+120 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "newTestEnv"
Cohesion: 0.07
Nodes (58): TestRealSource(), Import(), ParseSheet(), TestBuilderCompendium(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm() (+50 more)

### Community 28 - ".serverError"
Cohesion: 0.08
Nodes (7): isHTMX(), CompendiumRest(), WithCampaign(), LinkItem, SwitcherResults(), blockAccessLevel, who

### Community 29 - "testing.T"
Cohesion: 0.03
Nodes (79): TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), TestLevelUp(), TestLevelUpSubclassAndPools(), TestLevelUpThirdCaster(), TestMasteryLimit(), TestDedupeFeatures() (+71 more)

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
Cohesion: 0.04
Nodes (45): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, CreateCombatantParams, ListCombatantsRow, combatantDefenses() (+37 more)

### Community 34 - "joinNonEmpty"
Cohesion: 0.12
Nodes (24): habitatsFrom(), ftCell(), ftFeatIncrease(), ftParts(), ftTable(), TestFiveToolsCell(), TestFtFeatIncrease(), TestHabitatsFrom() (+16 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Client"
Cohesion: 0.16
Nodes (3): ClientInfo, Client, Hub

### Community 38 - "Queries"
Cohesion: 0.09
Nodes (33): featureSwitch, Queries, Env, ClassInfo, ClassLevel, featureNames(), SourceRef, SubclassRef (+25 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".fail"
Cohesion: 0.14
Nodes (5): Character, CanCharacter(), Flash(), nextStep(), wizardPath()

### Community 41 - "ParseWikidot"
Cohesion: 0.12
Nodes (26): block, ClassFeature, ClassData, dnd5eFeatures(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+18 more)

### Community 42 - "Book"
Cohesion: 0.08
Nodes (41): bookWriter, mirror, rawBook, internal/compendium, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook() (+33 more)

### Community 43 - "buildReview"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Parse"
Cohesion: 0.10
Nodes (21): Defense, feature(), Parse(), tokenize(), fighter(), TestCommands(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+13 more)

### Community 46 - "Feature"
Cohesion: 0.07
Nodes (21): FeatureEffects, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+13 more)

### Community 47 - "IfElse"
Cohesion: 0.06
Nodes (70): ListMonstersByNameRow, SearchCompendiumRow, Cell(), Row(), IfElse(), toJSON(), inFightSuffix(), alignmentColor() (+62 more)

### Community 48 - ".loadLive"
Cohesion: 0.06
Nodes (70): Props, Variant, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props (+62 more)

### Community 49 - "Topic"
Cohesion: 0.09
Nodes (14): event, fakeHandler, Handler, message, patchMessage, stateHandler, TopicKind, viewPatch (+6 more)

### Community 50 - "cell"
Cohesion: 0.14
Nodes (20): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+12 more)

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
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "CommandBuilder"
Cohesion: 0.13
Nodes (17): armorForm(), builderForm(), BuilderPickItems(), BuilderPicks(), chooseFirst(), closeAfter(), CommandBuilder(), compendiumLookup() (+9 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - ".changed"
Cohesion: 0.12
Nodes (10): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks(), parseCharacterFile() (+2 more)

### Community 59 - "ItemData"
Cohesion: 0.07
Nodes (35): ImageCategory, firstValue(), modTargetLabel(), usesItems(), weaponAbilities(), Roster, rosterCard(), conditionItems() (+27 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "CN"
Cohesion: 0.05
Nodes (112): ctxState, ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps (+104 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.31
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

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

### Community 68 - "App"
Cohesion: 0.07
Nodes (16): ftsQuery(), Search(), compendiumBack(), compendiumPath(), ImportSummary(), pathCampaign(), viewerCampaign(), firstLine() (+8 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Session"
Cohesion: 0.10
Nodes (19): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+11 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.22
Nodes (18): abilityMod(), ParseAidedd(), statLabel(), textLines(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+10 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "nullID"
Cohesion: 0.08
Nodes (30): Player(), nullID(), Gotchas, internal/web, Notes, Websocket and live views, TestCombatantNamesFollowSource(), TestEncounterTransitions() (+22 more)

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
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "alert_templ.go"
Cohesion: 0.15
Nodes (16): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+8 more)

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - ".Router"
Cohesion: 0.19
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Sheet"
Cohesion: 0.06
Nodes (52): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, TextData (+44 more)

### Community 100 - "joinNonEmpty"
Cohesion: 0.09
Nodes (29): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+21 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 104 - ".source"
Cohesion: 0.27
Nodes (8): testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceMove(), TestSourceOrder()

### Community 106 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (15): ftBook, ftRec, each(), ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftRecords() (+7 more)

### Community 108 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 109 - "writeBook"
Cohesion: 0.27
Nodes (9): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpBrokenSettings(), TestStartUpRedownloads(), writeBook(), appState (+1 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 114 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 115 - "Encounter"
Cohesion: 0.10
Nodes (11): Combatant, Encounter, Entity, Character, entityTopics(), linkEntityLook(), saveFieldLinks(), formInt() (+3 more)

### Community 116 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 117 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (92): UI conventions, Action(), Description(), Footer(), Header(), Title(), Scripts(), Separator() (+84 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".sourceIssues"
Cohesion: 0.43
Nodes (6): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 123 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 126 - ".failBack"
Cohesion: 0.15
Nodes (7): csvSource(), parseCSV(), TestGuessMapping(), Toast(), capReader, csvImport, csvImports

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (22): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eMastery(), TestUnflattenTable() (+14 more)

### Community 134 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 137 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Site"
Cohesion: 0.16
Nodes (10): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+2 more)

### Community 163 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 388 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **29 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.do`, `github.com/gin-gonic/gin.Context`, `perm.go`, `DecodeData`, `.serverError`, `database/sql.NullInt64`, `Client`, `Queries`, `.fail`, `.loadLive`, `Topic`, `cert.go`, `.changed`, `FiveToolsJob`, `.Router`, `Sheet`, `writeBook`, `.joinCampaigns`, `Encounter`, `.sourceIssues`, `.failBack`, `parseLiveKey`?**
  _High betweenness centrality (0.055) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11009658753382764 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `dialog_templ.go`, `database/sql.NullInt64`, `TestRunCommandUsesCurrentSheet`, `joinNonEmpty`, `Queries`, `SheetView`, `Command`, `Feature`, `Encounter`, `DecodeData`, `levelUpForm`, `.sourceIssues`, `Sheet`, `newTestEnv`, `testing.T`?**
  _High betweenness centrality (0.017) - this node is a cross-community bridge._
- **Should `dialog_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10034013605442177 - nodes in this community are weakly interconnected._
- **Why does `Pending` connect `Pending` to `compendium_templ.go`, `database/sql.NullInt64`, `.loadLive`, `Client`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.008254980079681274 - nodes in this community are weakly interconnected._