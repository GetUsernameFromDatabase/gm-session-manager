# Graph Report - gm-session-manager  (2026-11-09)

## Corpus Check
- 292 files · ~553,275 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 3996 nodes · 17252 edges · 131 communities (107 shown, 24 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1746 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `6c0d16ff`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Input
- context.Context
- .do
- createProvider
- .layout
- SheetView
- combobox.js
- redirect.go
- parseFiveToolsBook
- spellCounts
- select.js
- Command
- .Entry
- Button
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
- popover_templ.go
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- testing.T
- app.js
- tabbable.js
- useDismiss
- CompendiumEntry
- Parse
- command.js
- dialog.js
- empty_templ.go
- Sheet
- tooltip.js
- .encounterView
- ImagePage
- capReader
- export/markdown.go
- composite.js
- Client
- .wizardCommand
- ClassInfo
- State
- Character
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
- FiveToolsJob
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- Slugify
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
- IsLoopback
- Lint and Release Rules
- Run Script
- DecodeData
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
- .classFile
- .joinInfo
- time.Time
- ftBook
- Store
- SearchReadableBlocksRow
- formErrors
- .sendFeedback
- TestSourceSwitchKeepsEdits
- Weapon
- Slugify
- Pending
- build.sh
- Hub
- .levelUp
- errBodyWriter
- TestClassOptionChoices
- toast_templ.go
- RosterCharacter
- Before/after pairs
- open5e.go
- Import
- db/CLAUDE.md
- perm/CLAUDE.md

## God Nodes (most connected - your core abstractions)
1. `App` - 293 edges
2. `Sheet` - 116 edges
3. `newTestEnv()` - 116 edges
4. `Button()` - 113 edges
5. `CN()` - 113 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 56 edges
8. `Command` - 48 edges
9. `Input()` - 46 edges
10. `Slugify()` - 42 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `run()` --calls--> `NewManager()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/auth/manager.go
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (131 total, 24 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (57): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+49 more)

### Community 1 - "Input"
Cohesion: 0.05
Nodes (97): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+89 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (77): keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, lanHostname(), shortHost() (+69 more)

### Community 3 - ".do"
Cohesion: 0.16
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - ".layout"
Cohesion: 0.12
Nodes (8): Campaign, formIntQuery(), Roster, HomebrewView, LinkItem, SwitcherResults(), CreateCampaignParams, who

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (106): ArmorEffect(), Slot, Signed(), Action(), Card(), Content(), Description(), Header() (+98 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "parseFiveToolsBook"
Cohesion: 0.22
Nodes (10): internal/compendium, DownloadFiveTools(), FiveToolsBooks(), RemoveFiveTools(), readDoc(), serveMirror(), TestDownloadFiveTools(), TestDownloadFiveToolsFailureLeavesNothing() (+2 more)

### Community 10 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (68): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+60 more)

### Community 13 - ".Entry"
Cohesion: 0.07
Nodes (34): HomebrewOwner, MonsterAction, abilityMod(), addEntry(), ParseAidedd(), setStat(), statLabel(), CoreTraits() (+26 more)

### Community 14 - "Button"
Cohesion: 0.05
Nodes (96): Props, ImageCategory, UI conventions, Button(), Checkbox(), initialChecked(), Label(), Scripts() (+88 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (44): Note, GetBlockRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+36 more)

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
Cohesion: 0.07
Nodes (43): Field, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), armorForm(), builderForm() (+35 more)

### Community 21 - "levelUpForm"
Cohesion: 0.15
Nodes (28): initialValue(), Item(), RadioGroup(), state(), spellChecklist(), abilityForm(), backgroundAbilities(), choiceFieldset() (+20 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.05
Nodes (74): Props, Variant, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), Content(), contentID() (+66 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "popover_templ.go"
Cohesion: 0.15
Nodes (19): Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition(), conditionLabel() (+11 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.13
Nodes (29): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpTwice() (+21 more)

### Community 28 - "App"
Cohesion: 0.06
Nodes (22): ClearCookie(), SetCookie(), writeCookie(), FromContext(), CanCharacter(), coinCommand(), commandText(), hpCommand() (+14 more)

### Community 29 - "testing.T"
Cohesion: 0.04
Nodes (87): readBook(), TestFiveToolsCopyTargets(), fiveToolsClasses(), TestFiveToolsClasses(), TestFiveToolsClassLevelUp(), TestMasteryLimit(), TestDropFeat(), TestCanCharacter() (+79 more)

### Community 30 - "app.js"
Cohesion: 0.10
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "CompendiumEntry"
Cohesion: 0.06
Nodes (15): ListEntryCopiesRow, CompendiumEntry, spellFromRow(), compendiumBack(), ImportSummary(), viewerCampaign(), fiveToolsKey(), copyLabel() (+7 more)

### Community 34 - "Parse"
Cohesion: 0.07
Nodes (31): Defense, feature(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses() (+23 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "empty_templ.go"
Cohesion: 0.06
Nodes (61): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListMonstersByNameRow (+53 more)

### Community 38 - "Sheet"
Cohesion: 0.09
Nodes (16): DeathSaves, TestCustomMods(), Sum(), TestGearCommands(), Coins, Item, Sheet, WornArmor (+8 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".encounterView"
Cohesion: 0.09
Nodes (20): Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), damageChoice(), initiativeMod(), outOfFight() (+12 more)

### Community 41 - "ImagePage"
Cohesion: 0.07
Nodes (43): Props, Size, ButtonProps, ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, Type (+35 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 46 - ".wizardCommand"
Cohesion: 0.08
Nodes (30): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+22 more)

### Community 47 - "ClassInfo"
Cohesion: 0.08
Nodes (27): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), ClassInfo (+19 more)

### Community 48 - "State"
Cohesion: 0.11
Nodes (13): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+5 more)

### Community 49 - "Character"
Cohesion: 0.23
Nodes (6): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

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
Cohesion: 0.11
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ClassData"
Cohesion: 0.11
Nodes (40): block, ClassFeature, textLines(), ClassData, dnd5eFeatures(), ParsePlainPage(), attr(), byID() (+32 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - ".sourceIssues"
Cohesion: 0.36
Nodes (7): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, featKey(), pinFeat(), sheetRef

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
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (50): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, ListAttachmentFilesRow, GetImageRow, Attachment, Condition, Entity (+42 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (169): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+161 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Slugify"
Cohesion: 0.19
Nodes (14): Cantrips(), Slugify(), Spell, addProfs(), choiceOptions(), featList(), freeCast(), groupViews() (+6 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "Session"
Cohesion: 0.11
Nodes (17): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+9 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.12
Nodes (19): Player(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips() (+11 more)

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
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "IsLoopback"
Cohesion: 0.22
Nodes (5): IsLoopback(), TestIsLoopback(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 86 - "DecodeData"
Cohesion: 0.15
Nodes (15): dataFor(), DecodeData(), Kind, SpellData, TextData, TestTokensReadAsLabels(), InputForEntry(), TestMonsterFormRoundTrip() (+7 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Feature"
Cohesion: 0.22
Nodes (8): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, usesValue(), addFeat(), applyBackground()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (40): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+32 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.33
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Site"
Cohesion: 0.16
Nodes (10): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+2 more)

### Community 104 - ".classFile"
Cohesion: 0.14
Nodes (13): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+5 more)

### Community 106 - "time.Time"
Cohesion: 0.05
Nodes (26): Identity, ListImagesRow, Session, browserName(), firstLine(), mdLinkText(), TestUserAgent(), userAgentSummary() (+18 more)

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (23): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftDecode(), ftDuration() (+15 more)

### Community 108 - "Store"
Cohesion: 0.11
Nodes (22): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+14 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 112 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 113 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 115 - "Weapon"
Cohesion: 0.10
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 116 - "Slugify"
Cohesion: 0.21
Nodes (16): OptionPrerequisite(), NameDesc, catchUpResistances(), TestCatchUpResistances(), optionsGrid(), cantripFits(), checkNewOptions(), classOptionGroups() (+8 more)

### Community 117 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 123 - ".levelUp"
Cohesion: 0.16
Nodes (12): spellbook(), spellNeed(), classPin(), TestSpellbook(), TestSpellNeed(), LevelUpView, Option, SpellNeed (+4 more)

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "TestClassOptionChoices"
Cohesion: 0.14
Nodes (15): featureSwitch, firstTable(), LevelSpells(), Options(), TestOptions(), featureNames(), SwitchFeature, SwitchPlan (+7 more)

### Community 126 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 127 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (26): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eMastery(), TestUnflattenTable() (+18 more)

### Community 134 - "Import"
Cohesion: 0.10
Nodes (37): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+29 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 367 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **24 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `.layout`, `perm.go`, `CompendiumEntry`, `.encounterView`, `.wizardCommand`, `ClassInfo`, `State`, `Character`, `cert.go`, `.sourceIssues`, `FiveToolsJob`, `database/sql.NullInt64`, `Slugify`, `IsLoopback`, `DecodeData`, `.joinInfo`, `time.Time`, `Store`, `Hub`, `.levelUp`?**
  _High betweenness centrality (0.057) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11254599059477108 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `SheetView`, `Command`, `ParseSheet`, `CompendiumEntry`, `Parse`, `.encounterView`, `.wizardCommand`, `ClassInfo`, `.sourceIssues`, `Slugify`, `DecodeData`, `Feature`, `encoding/json.RawMessage`, `TestSourceSwitchKeepsEdits`, `Weapon`, `Slugify`, `.levelUp`, `TestClassOptionChoices`?**
  _High betweenness centrality (0.018) - this node is a cross-community bridge._
- **Should `Input` be split into smaller, more focused modules?**
  _Cohesion score 0.050458003415618694 - nodes in this community are weakly interconnected._
- **Why does `Entry` connect `open5e.go` to `compendium_templ.go`, `encoding/json.RawMessage`, `Site`, `ftBook`, `.Entry`, `DecodeData`, `ClassData`, `testing.T`?**
  _High betweenness centrality (0.016) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010245994132250056 - nodes in this community are weakly interconnected._