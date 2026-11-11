# Graph Report - gm-session-manager  (2026-11-11)

## Corpus Check
- 299 files · ~588,103 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4212 nodes · 18465 edges · 138 communities (112 shown, 26 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1898 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `ba6b7ff3`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- context.Context
- .do
- createProvider
- Sheet
- SheetView
- combobox.js
- redirect.go
- CommandBuilder
- Import
- select.js
- Command
- .Entry
- .whoami
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- dialog_templ.go
- levelUpForm
- wikilink.js
- IfElse
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- Parse
- command.js
- dialog.js
- Client
- Queries
- tooltip.js
- who
- ParseWikidot
- Book
- buildReview
- composite.js
- time.Time
- Feature
- compendium_templ.go
- github.com/a-h/templ.Component
- Topic
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- .wizardCommand
- fivetools_book.go
- radiogroup.js
- vendorjs/main.go
- Input
- FiveToolsJob
- sqlc Config
- fromLookup
- Browser check
- Claude Settings
- .campaignSourcesSet
- use_transition_status.js
- playwright
- release.sh
- Open
- checkbox.js
- ParseAidedd
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- alert_templ.go
- Lint and Release Rules
- Run Script
- RosterCharacter
- useTypeahead
- Sheet
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- joinNonEmpty
- mounted
- NewClient
- planSwitch
- Session
- capReader
- .sendFeedback
- ftBook
- sourcesEnv
- writeBook
- formErrors
- subclassPanels
- parseLiveKey
- spellCounts
- Pending
- Character
- ftItemData
- toast_templ.go
- build.sh
- .sourceIssues
- Props
- errBodyWriter
- SearchReadableBlocksRow
- buildReview
- Before/after pairs
- open5e.go
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- Site
- Client
- TestRunCommandUsesCurrentSheet

## God Nodes (most connected - your core abstractions)
1. `App` - 320 edges
2. `newTestEnv()` - 128 edges
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
- `UI conventions` --references--> `removeButton()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go
- `UI conventions` --references--> `selectItems()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (138 total, 26 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (63): contextKey, keepAwake(), classRef, indexed, looseString, o5eAssignment, o5eChild, o5eClass (+55 more)

### Community 1 - "CombatantView"
Cohesion: 0.07
Nodes (52): Condition, Description(), Header(), state(), Title(), Trigger(), triggerID(), addCombatantForm() (+44 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (93): keepAwake(), ftsQuery(), Search(), ListImageCharactersRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, CompendiumSource (+85 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): main(), isCleared(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Sheet"
Cohesion: 0.06
Nodes (22): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+14 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (118): ArmorEffect(), Slot, Signed(), Field, Look, Button(), Variant, Action() (+110 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "CommandBuilder"
Cohesion: 0.12
Nodes (19): armorForm(), builderForm(), BuilderPickItems(), BuilderPicks(), chooseFirst(), closeAfter(), CommandBuilder(), compendiumLookup() (+11 more)

### Community 10 - "Import"
Cohesion: 0.12
Nodes (32): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+24 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (67): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+59 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (32): MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, habitatsFrom(), ValidKind() (+24 more)

### Community 14 - ".whoami"
Cohesion: 0.05
Nodes (23): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), Session, browserName() (+15 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (45): Note, GetBlockRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+37 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "dialog_templ.go"
Cohesion: 0.21
Nodes (28): ctxKey, ctxState, Props, Dialog(), Close(), CloseFor(), Content(), Description() (+20 more)

### Community 21 - "levelUpForm"
Cohesion: 0.07
Nodes (62): Props, LevelUpPlan, Checkbox(), initialChecked(), attrString(), generateSVG(), getIconContent(), Props (+54 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "IfElse"
Cohesion: 0.13
Nodes (32): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+24 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (174): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+166 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.16
Nodes (26): ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpTwice() (+18 more)

### Community 28 - "App"
Cohesion: 0.06
Nodes (16): coinCommand(), commandText(), homebrewTabURL(), entityTopics(), formInt(), isHTMX(), redirect(), securityHeaders() (+8 more)

### Community 29 - "newTestEnv"
Cohesion: 0.03
Nodes (114): TestFiveToolsCell(), TestFiveToolsItemMastery(), TestFiveToolsRequires(), TestFtFeatIncrease(), TestHabitatsFrom(), fiveToolsClasses(), TestFiveToolsClassCell(), TestFiveToolsClasses() (+106 more)

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
Cohesion: 0.03
Nodes (42): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Attachment, CreateCombatantParams, applyCondOps(), saveScreenshot() (+34 more)

### Community 34 - "Parse"
Cohesion: 0.07
Nodes (35): feature(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+27 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 38 - "Queries"
Cohesion: 0.09
Nodes (30): dataFor(), DecodeData(), Kind, SpellData, Queries, CompendiumEntry, ClassInfo, ClassLevel (+22 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "who"
Cohesion: 0.10
Nodes (13): Campaign, compendiumBack(), compendiumPath(), pathCampaign(), viewerCampaign(), CompendiumRest(), HomebrewView, WithCampaign() (+5 more)

### Community 41 - "ParseWikidot"
Cohesion: 0.11
Nodes (27): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+19 more)

### Community 42 - "Book"
Cohesion: 0.10
Nodes (32): bookWriter, mirror, rawBook, internal/compendium, readBook(), TestFiveToolsCopyTargets(), bookFolder(), bookGroup() (+24 more)

### Community 43 - "buildReview"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "time.Time"
Cohesion: 0.09
Nodes (13): ListAttachmentFilesRow, Entity, linkEntityLook(), inCampaign(), FiveToolsView, NewFiveToolsView(), BlockStep, LevelUp (+5 more)

### Community 46 - "Feature"
Cohesion: 0.19
Nodes (8): FeatureEffects, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 47 - "compendium_templ.go"
Cohesion: 0.06
Nodes (102): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, SearchCompendiumRow, UI conventions, Badge(), Spinner(), IfElse() (+94 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (110): ListMonstersByNameRow, ListEntityFieldBacklinksRow, ListBacklinksRow, Content(), Description(), Empty(), Header(), Media() (+102 more)

### Community 49 - "Topic"
Cohesion: 0.10
Nodes (13): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+5 more)

### Community 50 - "cell"
Cohesion: 0.11
Nodes (21): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+13 more)

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
Nodes (14): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+6 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.10
Nodes (13): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+5 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - ".wizardCommand"
Cohesion: 0.07
Nodes (35): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+27 more)

### Community 59 - "fivetools_book.go"
Cohesion: 0.16
Nodes (21): ftBookInfo, ftDistance, ftRec, FiveToolsImages(), ftAmount(), ftCopyTargets(), ftDecode(), ftDuration() (+13 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.30
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 62 - "Input"
Cohesion: 0.06
Nodes (72): ImageCategory, Defenses, Combobox(), Content(), Item(), List(), Description(), Field() (+64 more)

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

### Community 68 - ".campaignSourcesSet"
Cohesion: 0.17
Nodes (6): CountCompendiumRow, ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), CopyKey()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Open"
Cohesion: 0.13
Nodes (22): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+14 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

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

### Community 80 - "ParseWiki"
Cohesion: 0.25
Nodes (10): wikiLink(), wikiText(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks(), WikiLinks() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "alert_templ.go"
Cohesion: 0.16
Nodes (16): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+8 more)

### Community 86 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Sheet"
Cohesion: 0.06
Nodes (52): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, TextData (+44 more)

### Community 100 - "joinNonEmpty"
Cohesion: 0.08
Nodes (36): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+28 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.28
Nodes (16): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+8 more)

### Community 103 - "planSwitch"
Cohesion: 0.15
Nodes (11): featureSwitch, featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch(), sameSlug() (+3 more)

### Community 104 - "Session"
Cohesion: 0.12
Nodes (16): fixture, tokens, constantTimeEqual(), credentialsMatch(), newFixture(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity() (+8 more)

### Community 106 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 107 - "ftBook"
Cohesion: 0.22
Nodes (12): ftBook, each(), ftCastingTime(), ftParts(), crProficiency(), ftEntries(), ftList(), ftSpells() (+4 more)

### Community 108 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 109 - "writeBook"
Cohesion: 0.27
Nodes (9): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpBrokenSettings(), TestStartUpRedownloads(), writeBook(), appState (+1 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 111 - "subclassPanels"
Cohesion: 0.33
Nodes (5): ListSubclassesOfRow, subclassIndex(), subclassOptions(), subclassPanels(), ListSubclassesOfParams

### Community 113 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 114 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 115 - "Character"
Cohesion: 0.06
Nodes (26): Defense, Character, Combatant, Encounter, ListCombatantsRow, Character, ParseDefenses(), Campaign() (+18 more)

### Community 116 - "ftItemData"
Cohesion: 0.32
Nodes (7): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), armorData(), armorText()

### Community 117 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".sourceIssues"
Cohesion: 0.53
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 123 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 126 - "buildReview"
Cohesion: 0.31
Nodes (5): ensureCategory(), buildReview(), formValues(), csvImport, csvImports

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Site"
Cohesion: 0.12
Nodes (17): HomebrewOwner, Site, Source, TestHomebrew(), fiveToolsPage(), SaveHomebrew(), SeedHabitats(), storeSource() (+9 more)

### Community 159 - "Client"
Cohesion: 0.20
Nodes (4): message, IsLoopback(), TestIsLoopback(), Client

### Community 163 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 388 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **26 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `.whoami`, `perm.go`, `database/sql.NullInt64`, `Client`, `Queries`, `who`, `time.Time`, `github.com/a-h/templ.Component`, `Topic`, `cert.go`, `.wizardCommand`, `FiveToolsJob`, `.campaignSourcesSet`, `Sheet`, `writeBook`, `parseLiveKey`, `Character`, `.sourceIssues`, `buildReview`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10843853820598007 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Parse`, `TestRunCommandUsesCurrentSheet`, `joinNonEmpty`, `Queries`, `SheetView`, `ParseSheet`, `planSwitch`, `Command`, `Feature`, `Character`, `levelUpForm`, `.sourceIssues`, `ParseEffects`, `.wizardCommand`, `Sheet`?**
  _High betweenness centrality (0.016) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.07226107226107226 - nodes in this community are weakly interconnected._
- **Why does `Pending` connect `Pending` to `compendium_templ.go`, `Client`, `time.Time`, `github.com/a-h/templ.Component`, `Client`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.008990318118948824 - nodes in this community are weakly interconnected._