# Graph Report - gm-session-manager  (2026-11-08)

## Corpus Check
- 286 files · ~548,091 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 3949 nodes · 17005 edges · 132 communities (111 shown, 21 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1722 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `85ef29f0`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- context.Context
- .do
- createProvider
- popover_templ.go
- SheetView
- combobox.js
- redirect.go
- parseFiveToolsBook
- Feature
- select.js
- Command
- .Entry
- Page
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- .Router
- levelUpForm
- wikilink.js
- tabs_templ.go
- popover.js
- FromContext
- Project Rules (CLAUDE.md, README)
- ParseSheet
- github.com/gin-gonic/gin.Context
- testing.T
- app.js
- tabbable.js
- useDismiss
- Queries
- Parse
- command.js
- dialog.js
- parseLiveKey
- golang.org/x/net/html.Node
- tooltip.js
- github.com/a-h/templ.Component
- DecodeData
- capReader
- export/markdown.go
- composite.js
- Client
- .levelUp
- ParseEffects
- State
- App
- ParseWiki
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- .render
- golangci-lint Config
- .sourceIssues
- Book
- radiogroup.js
- vendorjs/main.go
- cell
- Button
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- CN
- checkbox.js
- select_templ.go
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- roundTrip
- Lint and Release Rules
- Run Script
- github.com/a-h/templ.Component
- useTypeahead
- spellCounts
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- NewClient
- Site
- ftItemData
- .ListSpellsForClass
- .renderCompendium
- ftBook
- Store
- ImportView
- Campaign
- Campaign
- .sendFeedback
- TestSourceSwitchKeepsEdits
- .encounterView
- Weapon
- Slugify
- Pending
- build.sh
- ftItemData
- DBTX
- errBodyWriter
- buildReview
- toast_templ.go
- RosterCharacter
- NoteBlock
- ftDuration
- open5e.go
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 289 edges
2. `Sheet` - 116 edges
3. `Button()` - 113 edges
4. `newTestEnv()` - 113 edges
5. `CN()` - 113 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 55 edges
8. `Command` - 48 edges
9. `Input()` - 46 edges
10. `Slugify()` - 42 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Fingerprint()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/join/cert.go
- `run()` --calls--> `LoadOrCreateCert()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/join/cert.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (132 total, 21 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (58): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+50 more)

### Community 1 - "CombatantView"
Cohesion: 0.09
Nodes (39): Props, Variant, Badge(), badgeClasses(), Header(), Title(), beingBadge(), combatantChip() (+31 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (87): keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListImageEntitiesRow, SearchReadableBlocksRow (+79 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (80): ArmorEffect(), Coins, Slot, Signed(), Look, Variant, Action(), Card() (+72 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "parseFiveToolsBook"
Cohesion: 0.12
Nodes (21): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestImportFiveToolsBook(), TestParseFiveToolsBook() (+13 more)

### Community 10 - "Feature"
Cohesion: 0.07
Nodes (25): FeatureEffects, featureSwitch, ItemData, armorCondition(), columnInts(), ParseEffects(), UsesLabel(), CustomMod (+17 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (71): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+63 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput (+22 more)

### Community 14 - "Page"
Cohesion: 0.05
Nodes (67): ListEntityFieldBacklinksRow, ListBacklinksRow, attrString(), generateSVG(), getIconContent(), Props, Icon(), Scripts() (+59 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.10
Nodes (37): GetBlockRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock() (+29 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - ".Router"
Cohesion: 0.15
Nodes (8): IsLoopback(), TestIsLoopback(), TestRequireAdmin(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 21 - "levelUpForm"
Cohesion: 0.08
Nodes (53): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+45 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.09
Nodes (42): CountHabitatsRow, SearchCompendiumRow, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant() (+34 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "FromContext"
Cohesion: 0.06
Nodes (28): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), folderName() (+20 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.14
Nodes (28): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpTwice() (+20 more)

### Community 28 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (11): entityTopics(), isHTMX(), redirect(), csvSource(), setAudience(), securityHeaders(), visitorOnly(), DisplayPage() (+3 more)

### Community 29 - "testing.T"
Cohesion: 0.04
Nodes (84): TestMasteryLimit(), TestParseEffects(), TestDropFeat(), TestCanCharacter(), nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName() (+76 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Queries"
Cohesion: 0.10
Nodes (28): Queries, CompendiumEntry, Env, ClassFeature, ClassInfo, LevelUpPlan, classFeatures(), classInfo() (+20 more)

### Community 34 - "Parse"
Cohesion: 0.08
Nodes (29): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+21 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.10
Nodes (34): block, ClassFeature, abilityMod(), ParseAidedd(), statLabel(), textLines(), ClassData, dnd5eFeatures() (+26 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (121): ImageCategory, Defenses, damageChoice(), Combobox(), Content(), Item(), List(), Dialog() (+113 more)

### Community 41 - "DecodeData"
Cohesion: 0.08
Nodes (32): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+24 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.31
Nodes (9): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+1 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Client"
Cohesion: 0.16
Nodes (3): ClientInfo, Client, Hub

### Community 46 - ".levelUp"
Cohesion: 0.10
Nodes (22): Spell, addClassSpells(), errorField(), spellFromRow(), spellNeed(), classPin(), LevelUpView, Option (+14 more)

### Community 47 - "ParseEffects"
Cohesion: 0.16
Nodes (8): LevelUpChoice, TestSlotMaxima(), applyASI(), hpGain(), raise(), ClassLevel, CasterLevel(), SlotMaxima()

### Community 48 - "State"
Cohesion: 0.10
Nodes (14): event, fakeHandler, Handler, message, patchMessage, stateHandler, TopicKind, viewPatch (+6 more)

### Community 49 - "App"
Cohesion: 0.11
Nodes (9): Character, CanCharacter(), linkEntityLook(), linkCharacterLook(), lookFree(), Roster, nextStep(), wizardPath() (+1 more)

### Community 50 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

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
Cohesion: 0.13
Nodes (12): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+4 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - ".render"
Cohesion: 0.09
Nodes (8): compendiumBack(), formIntQuery(), joinPath(), HomebrewView, JoinQRSrc(), LinkItem, SwitcherResults(), who

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - ".sourceIssues"
Cohesion: 0.15
Nodes (12): ListEntryCopiesRow, CompendiumSource, pinnedRefs(), sourceIssue(), copyLabel(), SourceCopy, SourceIssue, SourceLabel() (+4 more)

### Community 59 - "Book"
Cohesion: 0.19
Nodes (15): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+7 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "cell"
Cohesion: 0.13
Nodes (18): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+10 more)

### Community 63 - "Button"
Cohesion: 0.06
Nodes (56): Props, Size, Field, Button(), Variants(), CloseFor(), Content(), DisableTrigger() (+48 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.06
Nodes (36): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, Entity, Note, CreateCombatantParams (+28 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (118): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+110 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "CN"
Cohesion: 0.05
Nodes (85): Props, Props, MediaProps, MediaVariant, ListMonstersByNameRow, Entity, AspectRatio(), Checkbox() (+77 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "select_templ.go"
Cohesion: 0.11
Nodes (23): Content(), Group(), initialValue(), itemLabel(), Label(), Select(), Separator(), state() (+15 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.16
Nodes (13): newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveSourcesAlert(), TestLiveVisibility(), TestLiveMastery(), LiveKey(), dial() (+5 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - ".get"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "roundTrip"
Cohesion: 0.19
Nodes (13): CoreTraits(), TestCoreTraits(), pageSource(), PageSourceKey(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSpecies() (+5 more)

### Community 86 - "github.com/a-h/templ.Component"
Cohesion: 0.07
Nodes (45): TestModifier(), Modifier(), alignmentColor(), classBadges(), classBody(), commaPrefixed(), CompendiumFindURL(), compendiumRow() (+37 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "spellCounts"
Cohesion: 0.24
Nodes (7): lanHostname(), shortHost(), TestShortHost(), sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (39): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+31 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.30
Nodes (15): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+7 more)

### Community 103 - "Site"
Cohesion: 0.18
Nodes (8): Site, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SavePage(), SiteByKey(), TestParsePageURL()

### Community 104 - "ftItemData"
Cohesion: 0.25
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), featuresSearchText(), ftFeatureKey()

### Community 106 - ".renderCompendium"
Cohesion: 0.16
Nodes (7): FiveToolsBooks(), ImportSummary(), fiveToolsKey(), FiveToolsView, NewFiveToolsView(), FiveToolsBook, FiveToolsGroup

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (17): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+9 more)

### Community 108 - "Store"
Cohesion: 0.07
Nodes (40): fixture, tokens, main(), run(), migration, constantTimeEqual(), credentialsMatch(), newFixture() (+32 more)

### Community 109 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), ImportReviewPage(), importRows() (+2 more)

### Community 111 - "Campaign"
Cohesion: 0.18
Nodes (8): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks(), parseCharacterFile()

### Community 112 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 113 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 114 - ".encounterView"
Cohesion: 0.12
Nodes (15): Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), initiativeMod(), outOfFight(), pcSheet() (+7 more)

### Community 115 - "Weapon"
Cohesion: 0.08
Nodes (14): DeathSaves, Part, Skill, TestCustomMods(), skillByKey(), Sum(), TestDualWield(), damageLabel() (+6 more)

### Community 116 - "Slugify"
Cohesion: 0.12
Nodes (29): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, Slugify() (+21 more)

### Community 117 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "buildReview"
Cohesion: 0.36
Nodes (4): buildReview(), formValues(), csvImport, csvImports

### Community 126 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 127 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 128 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 129 - "ftDuration"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (22): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eMastery(), TestUnflattenTable() (+14 more)

### Community 134 - "Import"
Cohesion: 0.12
Nodes (28): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+20 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 350 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **21 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.do`, `Feature`, `Page`, `perm.go`, `.Router`, `FromContext`, `github.com/gin-gonic/gin.Context`, `Queries`, `parseLiveKey`, `DecodeData`, `Client`, `.levelUp`, `State`, `.render`, `.sourceIssues`, `database/sql.NullInt64`, `CN`, `.renderCompendium`, `Store`, `Campaign`, `Campaign`, `.encounterView`, `Slugify`, `buildReview`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1095733534757925 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Weapon` to `compendium_templ.go`, `Queries`, `Parse`, `encoding/json.RawMessage`, `SheetView`, `github.com/a-h/templ.Component`, `DecodeData`, `Feature`, `Command`, `.levelUp`, `ParseEffects`, `TestSourceSwitchKeepsEdits`, `.encounterView`, `Slugify`, `.sourceIssues`, `ParseSheet`?**
  _High betweenness centrality (0.017) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.09294199860237597 - nodes in this community are weakly interconnected._
- **Why does `Entry` connect `open5e.go` to `compendium_templ.go`, `encoding/json.RawMessage`, `golang.org/x/net/html.Node`, `Site`, `DecodeData`, `parseFiveToolsBook`, `ftBook`, `.Entry`, `roundTrip`?**
  _High betweenness centrality (0.016) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.008722476332305074 - nodes in this community are weakly interconnected._