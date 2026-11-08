# Graph Report - gm-session-manager  (2026-11-08)

## Corpus Check
- 291 files · ~548,625 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 3967 nodes · 17028 edges · 137 communities (113 shown, 24 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1733 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `a846dd34`
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
- ClassInfo
- select.js
- Command
- .Entry
- github.com/a-h/templ.Component
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- CommandBuilder
- levelUpForm
- wikilink.js
- tabs_templ.go
- popover.js
- Campaign
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
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
- dialog_templ.go
- .wizardCommand
- capReader
- export/markdown.go
- composite.js
- Client
- .levelUp
- UpgradeSheets
- State
- App
- ParseWiki
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- .sourceIssues
- Book
- radiogroup.js
- vendorjs/main.go
- cell
- tooltip_templ.go
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- Button
- checkbox.js
- Session
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
- Feature
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- NewClient
- Site
- ftItemData
- Attachment
- time.Time
- ftBook
- Store
- ItemData
- Kind
- Campaign
- .sendFeedback
- TestSourceSwitchKeepsEdits
- formInt
- Weapon
- Slugify
- Pending
- build.sh
- Hub
- LevelUpView
- errBodyWriter
- TestClassOptionChoices
- toast_templ.go
- RosterCharacter
- .loadBlock
- ftDuration
- Before/after pairs
- message
- .ListItemsByCategory
- open5e.go
- Import
- db/CLAUDE.md
- perm/CLAUDE.md

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

## Communities (137 total, 24 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.12
Nodes (42): contextKey, keepAwake(), classRef, ftBookInfo, indexed, sourceOf, ASIError, clientMessage (+34 more)

### Community 1 - "CombatantView"
Cohesion: 0.12
Nodes (31): Props, Variant, Defenses, damageChoice(), Badge(), badgeClasses(), capitalized(), beingBadge() (+23 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (70): keepAwake(), ftsQuery(), Search(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, SearchReadableBlocksRow (+62 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (22): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+14 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "popover_templ.go"
Cohesion: 0.11
Nodes (29): Look, Content(), Description(), Header(), Popover(), state(), Title(), Trigger() (+21 more)

### Community 6 - "SheetView"
Cohesion: 0.08
Nodes (68): ArmorEffect(), Slot, Signed(), Variant, Action(), Card(), Content(), Description() (+60 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "parseFiveToolsBook"
Cohesion: 0.11
Nodes (20): internal/compendium, parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestParseFiveToolsBook(), DownloadFiveTools() (+12 more)

### Community 10 - "ClassInfo"
Cohesion: 0.07
Nodes (21): featureSwitch, LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), ClassInfo (+13 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (66): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+58 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (27): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+19 more)

### Community 14 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (129): ListMonstersByNameRow, ListEntityFieldBacklinksRow, ListBacklinksRow, UI conventions, Checkbox(), Content(), Description(), Empty() (+121 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (48): Note, ListCombatantsRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+40 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "CommandBuilder"
Cohesion: 0.13
Nodes (17): armorForm(), builderForm(), builderURL(), chooseFirst(), CommandBuilder(), compendiumLookup(), masteryCount(), masteryForm() (+9 more)

### Community 21 - "levelUpForm"
Cohesion: 0.09
Nodes (44): EquipmentItem, EquipmentOption, Footer(), initialValue(), Item(), RadioGroup(), state(), abilitySelect() (+36 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.08
Nodes (52): CountHabitatsRow, Title(), Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant() (+44 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "Campaign"
Cohesion: 0.06
Nodes (22): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), Campaign, browserName() (+14 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.14
Nodes (28): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpTwice() (+20 more)

### Community 28 - "App"
Cohesion: 0.06
Nodes (16): compendiumBack(), ImportSummary(), viewerCampaign(), fiveToolsKey(), isHTMX(), redirect(), csvSource(), securityHeaders() (+8 more)

### Community 29 - "testing.T"
Cohesion: 0.05
Nodes (77): TestMasteryLimit(), TestDropFeat(), TestCanCharacter(), nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm() (+69 more)

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
Cohesion: 0.07
Nodes (20): DecodeData(), ListEntryCopiesRow, Queries, CompendiumEntry, CompendiumSource, masteryLimit(), setSubclass(), spellFromRow() (+12 more)

### Community 34 - "Parse"
Cohesion: 0.09
Nodes (27): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+19 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "dialog_templ.go"
Cohesion: 0.08
Nodes (64): ctxKey, ctxState, Props, Content(), Item(), List(), optionID(), Dialog() (+56 more)

### Community 41 - ".wizardCommand"
Cohesion: 0.11
Nodes (24): CoreTraits(), EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice() (+16 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.09
Nodes (23): Links, pmMark, pmNode, writer, IsLoopback(), TestIsLoopback(), TestMarkdown(), applyMarks() (+15 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 46 - ".levelUp"
Cohesion: 0.12
Nodes (24): Cantrips(), Slugify(), Spell, addClassSpells(), errorField(), spellNeed(), classEntry(), classPin() (+16 more)

### Community 47 - "UpgradeSheets"
Cohesion: 0.25
Nodes (8): TestSlotMaxima(), ClassLevel, CasterLevel(), SlotMaxima(), catchUpMasteries(), upgradeClass(), UpgradeSheets(), upgradeSpecies()

### Community 48 - "State"
Cohesion: 0.12
Nodes (11): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+3 more)

### Community 49 - "App"
Cohesion: 0.09
Nodes (14): Character, CanCharacter(), coinCommand(), commandText(), hpCommand(), parseCommand(), linkCharacterLook(), Roster (+6 more)

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

### Community 56 - "ClassData"
Cohesion: 0.14
Nodes (16): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+8 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - ".sourceIssues"
Cohesion: 0.53
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 59 - "Book"
Cohesion: 0.17
Nodes (17): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+9 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "cell"
Cohesion: 0.13
Nodes (22): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+14 more)

### Community 63 - "tooltip_templ.go"
Cohesion: 0.11
Nodes (32): SearchCompendiumRow, Field, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems() (+24 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (44): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Condition, CreateCombatantParams, applyCondOps(), combatantDefenses(), petrified() (+36 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (192): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+184 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Button"
Cohesion: 0.09
Nodes (32): Props, Size, Encounter, Button(), Variants(), notFoundIsNil(), campaignLink(), CampaignRoster() (+24 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "Session"
Cohesion: 0.18
Nodes (9): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+1 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.12
Nodes (18): Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveSourcesAlert() (+10 more)

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
Cohesion: 0.24
Nodes (14): TestTokensReadAsLabels(), InputForEntry(), TestMonsterFormRoundTrip(), ParseWikidot(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSpecies() (+6 more)

### Community 86 - "github.com/a-h/templ.Component"
Cohesion: 0.10
Nodes (34): SpellData, alignmentColor(), classBadges(), classBody(), commaPrefixed(), compendiumRow(), crText(), entryBody() (+26 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Feature"
Cohesion: 0.19
Nodes (8): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesItems(), usesValue()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (40): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+32 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.30
Nodes (15): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+7 more)

### Community 103 - "Site"
Cohesion: 0.14
Nodes (13): Site, Source, fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page, pageSource() (+5 more)

### Community 104 - "ftItemData"
Cohesion: 0.13
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, ItemData, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 105 - "Attachment"
Cohesion: 0.20
Nodes (6): Attachment, saveScreenshot(), parseCSV(), TestGuessMapping(), saveSniffed(), CreateAttachmentParams

### Community 106 - "time.Time"
Cohesion: 0.10
Nodes (11): FiveToolsBooks(), ListAttachmentFilesRow, ListImagesRow, FiveToolsView, NewFiveToolsView(), BlockStep, LevelUp, ListImagesParams (+3 more)

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (19): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+11 more)

### Community 108 - "Store"
Cohesion: 0.09
Nodes (32): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply() (+24 more)

### Community 109 - "ItemData"
Cohesion: 0.11
Nodes (24): ImageCategory, firstValue(), modTargetLabel(), weaponAbilities(), conditionItems(), selectItems(), kindSelect(), masteryItems() (+16 more)

### Community 110 - "Kind"
Cohesion: 0.20
Nodes (11): dataFor(), Kind, TextData, wikidotKind(), textMarkdown(), CompendiumFindURL(), MasteryLinks(), countWord() (+3 more)

### Community 111 - "Campaign"
Cohesion: 0.22
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks(), setAudience()

### Community 112 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 113 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 114 - "formInt"
Cohesion: 0.10
Nodes (13): Combatant, Entity, Character, entityTopics(), linkEntityLook(), saveFieldLinks(), formInt(), lookFree() (+5 more)

### Community 115 - "Weapon"
Cohesion: 0.06
Nodes (24): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+16 more)

### Community 116 - "Slugify"
Cohesion: 0.16
Nodes (21): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, classInfo(), catchUpResistances() (+13 more)

### Community 117 - "Pending"
Cohesion: 0.17
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 123 - "LevelUpView"
Cohesion: 0.27
Nodes (7): LevelUpView, Option, SpellNeed, Plural(), drawsGroup(), ChoiceGroup, CreationChoices

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "TestClassOptionChoices"
Cohesion: 0.25
Nodes (8): SwitchFeature, SwitchPlan, PlanSourceSwitch(), names(), switchFighter(), TestPlanSourceSwitch(), TestClassOptionChoices(), SourcePreviewView

### Community 126 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 127 - "RosterCharacter"
Cohesion: 0.50
Nodes (3): deviceLabel(), rosterHP(), Device

### Community 129 - "ftDuration"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "message"
Cohesion: 0.40
Nodes (3): message, Diff(), TestDiff()

### Community 133 - "open5e.go"
Cohesion: 0.08
Nodes (37): Entry, fixture, fixtures, looseString, normalizer, o5eArmor, o5eAssignment, o5eChild (+29 more)

### Community 134 - "Import"
Cohesion: 0.13
Nodes (27): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+19 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 362 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **24 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.loadBlock`, `context.Context`, `.do`, `perm.go`, `Campaign`, `Queries`, `parseLiveKey`, `.wizardCommand`, `.levelUp`, `State`, `App`, `.sourceIssues`, `database/sql.NullInt64`, `Button`, `Attachment`, `time.Time`, `Store`, `Kind`, `Campaign`, `formInt`, `Hub`, `LevelUpView`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.12157988246549133 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Weapon` to `compendium_templ.go`, `CombatantView`, `SheetView`, `ClassInfo`, `Command`, `perm.go`, `ParseSheet`, `Queries`, `Parse`, `.wizardCommand`, `.levelUp`, `UpgradeSheets`, `.sourceIssues`, `Feature`, `encoding/json.RawMessage`, `TestSourceSwitchKeepsEdits`, `formInt`, `Slugify`, `LevelUpView`, `TestClassOptionChoices`?**
  _High betweenness centrality (0.020) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.11605937921727395 - nodes in this community are weakly interconnected._
- **Why does `Entry` connect `open5e.go` to `compendium_templ.go`, `encoding/json.RawMessage`, `golang.org/x/net/html.Node`, `Site`, `parseFiveToolsBook`, `ftBook`, `.Entry`, `Kind`, `roundTrip`?**
  _High betweenness centrality (0.016) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011084636900445458 - nodes in this community are weakly interconnected._