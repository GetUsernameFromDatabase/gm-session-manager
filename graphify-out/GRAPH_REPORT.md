# Graph Report - gm-session-manager  (2026-11-13)

## Corpus Check
- 317 files · ~594,343 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4308 nodes · 19728 edges · 158 communities (128 shown, 30 thin omitted)
- Extraction: 91% EXTRACTED · 9% INFERRED · 0% AMBIGUOUS · INFERRED: 1848 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `1ffe27e8`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- dialog_templ.go
- context.Context
- .do
- createProvider
- Weapon
- SheetView
- combobox.js
- redirect.go
- ImagePage
- Store
- select.js
- command.go
- .Entry
- capReader
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ftItemData
- levelUpForm
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- Slugify
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- testing.T
- command.js
- dialog.js
- DownloadFiveTools
- .levelUp
- tooltip.js
- who
- ParseWikidot
- Book
- export/markdown.go
- composite.js
- Option
- ParseEffects
- compendium_templ.go
- Character
- SaveHomebrew
- ImportView
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Sheet
- golangci-lint Config
- .featChoices
- Props
- radiogroup.js
- vendorjs/main.go
- Input
- FiveToolsJob
- sqlc Config
- select_templ.go
- Browser check
- Claude Settings
- .compendiumDir
- use_transition_status.js
- playwright
- release.sh
- Session
- checkbox.js
- ParseAidedd
- switch.js
- tabs_templ.go
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- Feature
- Lint and Release Rules
- Run Script
- spellCounts
- useTypeahead
- Button
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- empty_templ.go
- newLiveFixture
- .equipmentChoices
- State
- CompendiumEntry
- ftBook
- tooltip_templ.go
- ListCombatantsRow
- Import
- viewerCtx
- config.go
- Client
- Command
- scoped
- Campaign
- NewClient
- build.sh
- .sourceIssues
- looseString
- errBodyWriter
- ClassInfo
- planSwitch
- net/url.Values
- Before/after pairs
- CountSharedByKindRow
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- Pending
- parseHP
- isToggle
- sourceUsedError
- TestRunCommandUsesCurrentSheet
- findItem
- toast_templ.go
- Defenses
- InputForEntry
- writeBook
- WeaponKind
- RosterCharacter
- ftDuration
- parseMod
- IsLoopback
- ListSharedWithRow
- SearchReadableBlocksRow
- ListSharedEntriesRow
- Parse

## God Nodes (most connected - your core abstractions)
1. `App` - 349 edges
2. `newTestEnv()` - 131 edges
3. `Button()` - 122 edges
4. `Sheet` - 121 edges
5. `CN()` - 113 edges
6. `IfElse()` - 62 edges
7. `Command` - 59 edges
8. `SheetView` - 57 edges
9. `CompendiumKind` - 54 edges
10. `ParseSheet()` - 47 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (158 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.13
Nodes (53): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+45 more)

### Community 1 - "dialog_templ.go"
Cohesion: 0.12
Nodes (31): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), downBadge() (+23 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (78): keepAwake(), ftsQuery(), Search(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource (+70 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.12
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (91): SpellData, ArmorEffect(), Slot, Signed(), Variant, Action(), Card(), Content() (+83 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "ImagePage"
Cohesion: 0.15
Nodes (20): ListEntityFieldBacklinksRow, Entity, Look, initialChecked(), Switch(), initials(), currentLook(), EntityDetail() (+12 more)

### Community 10 - "Store"
Cohesion: 0.08
Nodes (34): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply() (+26 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.13
Nodes (21): CondOpKind, HelpEntry, parser, condCommand(), conditionSlug(), CondOp, parseConcentrate(), parseCondition() (+13 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (23): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInputFor(), abilityKey(), actionsText() (+15 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (48): GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass() (+40 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ftItemData"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 21 - "levelUpForm"
Cohesion: 0.07
Nodes (59): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Action(), Alert() (+51 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "App"
Cohesion: 0.05
Nodes (19): CanCharacter(), homebrewTabURL(), entityTopics(), formInt(), formIntQuery(), isHTMX(), redirect(), csvSource() (+11 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (155): Props, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size (+147 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.14
Nodes (28): ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage() (+20 more)

### Community 28 - "Slugify"
Cohesion: 0.12
Nodes (27): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, Slugify() (+19 more)

### Community 29 - "newTestEnv"
Cohesion: 0.05
Nodes (55): nullID(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestEncounterTransitions(), TestMonsterPortrait(), TestDeleteRestoreImport(), TestRenameDeleteCampaign(), testEnv (+47 more)

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
Nodes (55): SeedHabitats(), Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, Entity (+47 more)

### Community 34 - "testing.T"
Cohesion: 0.04
Nodes (76): TestDecode(), CoreTraits(), TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground() (+68 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "DownloadFiveTools"
Cohesion: 0.11
Nodes (25): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook() (+17 more)

### Community 38 - ".levelUp"
Cohesion: 0.13
Nodes (16): CantripFits(), Cantrip, MeetsPrerequisite(), testCantrip(), TestCantripFits(), TestMeetsPrerequisite(), Handlers, formErrors() (+8 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "who"
Cohesion: 0.08
Nodes (18): HomebrewOwner, HomebrewInput, SaveHomebrew(), LinkTarget, compendiumBack(), pathCampaign(), CompendiumRest(), HomebrewView (+10 more)

### Community 41 - "ParseWikidot"
Cohesion: 0.18
Nodes (15): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+7 more)

### Community 42 - "Book"
Cohesion: 0.11
Nodes (24): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+16 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Option"
Cohesion: 0.16
Nodes (13): LevelUpPlan, subclassOptions(), withSpells(), LevelUpView, Option, SpellNeed, SubclassPanel, splitCantrips() (+5 more)

### Community 46 - "ParseEffects"
Cohesion: 0.13
Nodes (10): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+2 more)

### Community 47 - "compendium_templ.go"
Cohesion: 0.05
Nodes (110): dataFor(), CountCompendiumRow, CountHabitatsRow, ListEntryCopiesRow, ListHomebrewByOwnerRow, ListHomebrewRow, SearchCompendiumRow, CompendiumKind (+102 more)

### Community 48 - "Character"
Cohesion: 0.09
Nodes (17): DecodeData(), Character, Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile() (+9 more)

### Community 49 - "SaveHomebrew"
Cohesion: 0.16
Nodes (10): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+2 more)

### Community 50 - "ImportView"
Cohesion: 0.09
Nodes (33): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+25 more)

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

### Community 56 - "Sheet"
Cohesion: 0.07
Nodes (23): DeathSaves, AddResistance(), TestCustomMods(), TestResistances(), TestExplainSums(), Sum(), TestGearCommands(), TestUpgradeSheet() (+15 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.08
Nodes (26): rules, generated, paths, rules, formatters, enable, exclusions, settings (+18 more)

### Community 58 - ".featChoices"
Cohesion: 0.13
Nodes (18): internal/compendium, SkillChoice(), TestSkillChoice(), TextData, AbilityKeys(), AddFeat(), SkillKeys(), CountWord() (+10 more)

### Community 59 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Input"
Cohesion: 0.05
Nodes (90): ctxKey, ctxState, Props, CustomMod, Dialog(), Close(), CloseFor(), Content() (+82 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "select_templ.go"
Cohesion: 0.10
Nodes (28): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+20 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".compendiumDir"
Cohesion: 0.21
Nodes (4): ImportSummary(), fiveToolsKey(), FiveToolsView, NewFiveToolsView()

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
Cohesion: 0.16
Nodes (11): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+3 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "tabs_templ.go"
Cohesion: 0.12
Nodes (33): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+25 more)

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

### Community 83 - "Feature"
Cohesion: 0.23
Nodes (5): FeatureEffects, UsesLabel(), Feature, Uses, usesValue()

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), sortedColumns(), spellCounts(), TestSpellCounts()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Button"
Cohesion: 0.06
Nodes (72): Props, Size, ImageCategory, ListBacklinksRow, UI conventions, Button(), Variants(), Scripts() (+64 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (38): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+30 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "empty_templ.go"
Cohesion: 0.08
Nodes (51): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Content() (+43 more)

### Community 103 - "newLiveFixture"
Cohesion: 0.10
Nodes (23): Player(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas, internal/web, Notes, Websocket and live views (+15 more)

### Community 104 - ".equipmentChoices"
Cohesion: 0.29
Nodes (7): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), TestOptionText(), EquipmentChoice, optionText()

### Community 105 - "State"
Cohesion: 0.10
Nodes (13): event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch, Topic (+5 more)

### Community 106 - "CompendiumEntry"
Cohesion: 0.07
Nodes (16): Decode(), ListSubclassesOfRow, CompendiumEntry, SourceRef, compendiumPath(), viewerCampaign(), spellFromRow(), sourceAllowed() (+8 more)

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (23): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftMod(), ftName() (+15 more)

### Community 108 - "tooltip_templ.go"
Cohesion: 0.08
Nodes (40): ListMonstersByNameRow, Field, Cell(), Row(), Content(), DisableTrigger(), state(), Tooltip() (+32 more)

### Community 109 - "ListCombatantsRow"
Cohesion: 0.16
Nodes (14): Combatant, ListCombatantsRow, rollFor(), stepTurn(), combatantKind(), combatantDefenses(), damageHP(), healHP() (+6 more)

### Community 110 - "Import"
Cohesion: 0.10
Nodes (33): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+25 more)

### Community 111 - "viewerCtx"
Cohesion: 0.23
Nodes (6): notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView, viewerCtx

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "Client"
Cohesion: 0.16
Nodes (3): ClientInfo, Client, Hub

### Community 114 - "Command"
Cohesion: 0.18
Nodes (15): Effects, Command, Env, NewCommand(), parseSet(), parseWeapon(), SaveWeapon(), SaveFeature() (+7 more)

### Community 115 - "scoped"
Cohesion: 0.20
Nodes (15): init(), parseCoin(), parseDeathSave(), parseHitDie(), parseInspiration(), parseMaxHP(), parsePact(), parseRest() (+7 more)

### Community 116 - "Campaign"
Cohesion: 0.05
Nodes (24): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), Campaign, browserName() (+16 more)

### Community 117 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".sourceIssues"
Cohesion: 0.53
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ClassInfo"
Cohesion: 0.09
Nodes (29): TestSlotMaxima(), ClassFeature, ClassInfo, ClassLevel, CasterLevel(), SlotMaxima(), Spellbook(), SpellTotals() (+21 more)

### Community 126 - "planSwitch"
Cohesion: 0.16
Nodes (11): featureSwitch, ftRaceAbility(), featureNames(), SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch(), sameSlug() (+3 more)

### Community 127 - "net/url.Values"
Cohesion: 0.24
Nodes (13): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+5 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "CountSharedByKindRow"
Cohesion: 0.12
Nodes (8): CountSharedByKindParams, CountSharedByKindRow, ListCompendiumByKindParams, ListCompendiumByKindRow, ListCompendiumBySlugsParams, ListCompendiumBySlugsRow, ListEntrySourcesBySlugsParams, ListEntrySourcesBySlugsRow

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (23): Entry, fixture, fixtures, normalizer, o5eProperty, Source, sourceFormat, TestOpen5eMastery() (+15 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Pending"
Cohesion: 0.20
Nodes (5): PendingState, randomID(), Pending, PendingKind, Waiting

### Community 140 - "parseHP"
Cohesion: 0.33
Nodes (8): Amount, Damage(), Heal(), parseAmount(), parseDamage(), parseHP(), Set(), TestHealingFromZeroResetsDeathSaves()

### Community 143 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 144 - "findItem"
Cohesion: 0.32
Nodes (4): findItem(), parseArmor(), parseItem(), ensureItem()

### Community 145 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 146 - "Defenses"
Cohesion: 0.33
Nodes (5): Defense, TestDefenses(), Defenses, ParseDefenses(), damageChoice()

### Community 147 - "InputForEntry"
Cohesion: 0.29
Nodes (7): habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 148 - "writeBook"
Cohesion: 0.33
Nodes (7): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpBrokenSettings(), TestStartUpRedownloads(), writeBook()

### Community 149 - "WeaponKind"
Cohesion: 0.40
Nodes (5): Completions(), parseMastery(), TestCompletions(), TestWeaponKind(), WeaponKind()

### Community 150 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 151 - "ftDuration"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 152 - "parseMod"
Cohesion: 0.67
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 163 - "Parse"
Cohesion: 0.17
Nodes (13): Parse(), tokenize(), fighter(), TestCommands(), TestMasteryCommand(), TestParseRawArgs(), TestPrepareCap(), TestUnknownCommand() (+5 more)

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 393 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Store`, `Pending`, `perm.go`, `writeBook`, `Slugify`, `database/sql.NullInt64`, `.levelUp`, `who`, `Book`, `Option`, `compendium_templ.go`, `Character`, `Sheet`, `.featChoices`, `FiveToolsJob`, `.compendiumDir`, `State`, `CompendiumEntry`, `ListCombatantsRow`, `viewerCtx`, `Client`, `Command`, `Campaign`, `.sourceIssues`, `ClassInfo`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13232026217144396 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Weapon`, `SheetView`, `TestRunCommandUsesCurrentSheet`, `findItem`, `Defenses`, `WeaponKind`, `ParseSheet`, `Slugify`, `testing.T`, `.levelUp`, `Option`, `ParseEffects`, `Character`, `.featChoices`, `Input`, `Feature`, `encoding/json.RawMessage`, `CompendiumEntry`, `ListCombatantsRow`, `.sourceIssues`, `ClassInfo`, `planSwitch`?**
  _High betweenness centrality (0.012) - this node is a cross-community bridge._
- **Should `dialog_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1226215644820296 - nodes in this community are weakly interconnected._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01033623910336239 - nodes in this community are weakly interconnected._
- **Should `.do` be split into smaller, more focused modules?**
  _Cohesion score 0.10695187165775401 - nodes in this community are weakly interconnected._