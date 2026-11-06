# Graph Report - gm-session-manager  (2026-11-06)

## Corpus Check
- 286 files · ~540,430 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 3911 nodes · 16739 edges · 128 communities (104 shown, 24 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1680 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `1b4d3371`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- ImportView
- context.Context
- .do
- createProvider
- CommandBuilder
- SheetView
- combobox.js
- redirect.go
- parseFiveToolsBook
- Feature
- select.js
- Command
- .Entry
- Input
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- NoteView
- wizardPick
- wikilink.js
- IfElse
- popover.js
- Props
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- testing.T
- app.js
- tabbable.js
- useDismiss
- Character
- Parse
- command.js
- dialog.js
- select_templ.go
- golang.org/x/net/html.Node
- tooltip.js
- CombatantView
- parseCommand
- capReader
- export/markdown.go
- composite.js
- WeaponKind
- popover_templ.go
- ParseEffects
- Client
- Sheet
- ParseWiki
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- CompendiumEntry
- golangci-lint Config
- Sum
- Book
- radiogroup.js
- vendorjs/main.go
- cell
- BuilderPicks
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- empty_templ.go
- checkbox.js
- radiogroup_templ.go
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- ParseWikidot
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
- .ListCompendiumByKind
- Site
- ftItemData
- .ListItemsByCategory
- Verbs
- ftBook
- Store
- .Router
- .sendFeedback
- TestSourceSwitchKeepsEdits
- .loadLive
- Weapon
- PlanSourceSwitch
- build.sh
- ftItemData
- RosterCharacter
- errBodyWriter
- NoteListView
- .ListBlockStepsSince
- .ListCompendiumBySlugs
- .ListEntrySourcesBySlugs
- open5e.go
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 286 edges
2. `CN()` - 113 edges
3. `Button()` - 112 edges
4. `Sheet` - 108 edges
5. `newTestEnv()` - 107 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 52 edges
8. `Command` - 47 edges
9. `Input()` - 46 edges
10. `Slugify()` - 41 edges

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

## Communities (128 total, 24 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (58): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+50 more)

### Community 1 - "ImportView"
Cohesion: 0.20
Nodes (12): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportColumn, ImportValue, ImportView, ImportReview() (+4 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (69): keepAwake(), ftsQuery(), Search(), Note, GetBlockRow, CreateCombatantParams, Fields, ParseFields() (+61 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (22): isCleared(), serve(), TestMiddlewareIdentity(), Identity, Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+14 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "CommandBuilder"
Cohesion: 0.27
Nodes (9): armorForm(), CommandBuilder(), compendiumLookup(), modList(), opGroup(), profTabs(), spellSearch(), builderOp (+1 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (106): ArmorEffect(), Signed(), Field, Button(), Variant, Action(), Card(), Content() (+98 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.25
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "parseFiveToolsBook"
Cohesion: 0.24
Nodes (10): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestParseFiveToolsBook(), feature(), fiveToolsClasses() (+2 more)

### Community 10 - "Feature"
Cohesion: 0.13
Nodes (15): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+7 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.10
Nodes (48): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+40 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (31): HomebrewOwner, MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, ValidKind() (+23 more)

### Community 14 - "Input"
Cohesion: 0.06
Nodes (90): ctxKey, ctxState, Props, Coins, Defenses, damageChoice(), Combobox(), Content() (+82 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.11
Nodes (34): PendingState, Pending, PendingKind, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock() (+26 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "NoteView"
Cohesion: 0.10
Nodes (28): ListEntityFieldBacklinksRow, ListBacklinksRow, initialChecked(), Switch(), EntityDetail(), entityForm(), EntityPage(), EntityView (+20 more)

### Community 21 - "wizardPick"
Cohesion: 0.15
Nodes (24): Footer(), abilitySelect(), levelUpForm(), slotSummary(), spellChecklist(), abilityForm(), backgroundAbilities(), choiceFieldset() (+16 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "IfElse"
Cohesion: 0.06
Nodes (77): Label(), Scripts(), Spinner(), Content(), contentID(), IDFromContext(), initialValue(), List() (+69 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.15
Nodes (26): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp() (+18 more)

### Community 28 - "App"
Cohesion: 0.05
Nodes (29): ClearCookie(), SetCookie(), writeCookie(), FromContext(), CanCharacter(), ImportSummary(), entityTopics(), fiveToolsKey() (+21 more)

### Community 29 - "testing.T"
Cohesion: 0.04
Nodes (84): TestMasteryLimit(), TestDropFeat(), TestCanCharacter(), nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm() (+76 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Character"
Cohesion: 0.13
Nodes (10): Character, Entity, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), linkEntityLook() (+2 more)

### Community 34 - "Parse"
Cohesion: 0.06
Nodes (38): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+30 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "select_templ.go"
Cohesion: 0.11
Nodes (22): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state() (+14 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CombatantView"
Cohesion: 0.12
Nodes (32): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), downBadge() (+24 more)

### Community 41 - "parseCommand"
Cohesion: 0.40
Nodes (5): weaponCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 46 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 47 - "ParseEffects"
Cohesion: 0.11
Nodes (13): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), LevelUpPlan (+5 more)

### Community 48 - "Client"
Cohesion: 0.07
Nodes (30): ClientInfo, event, fakeHandler, Handler, message, patchMessage, stateHandler, TopicKind (+22 more)

### Community 49 - "Sheet"
Cohesion: 0.04
Nodes (96): EquipmentItem, DeathSaves, Cantrips(), EquipmentOptions(), firstTable(), EquipmentOption, LevelSpells(), OptionPrerequisite() (+88 more)

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
Cohesion: 0.18
Nodes (12): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+4 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "CompendiumEntry"
Cohesion: 0.06
Nodes (20): ListEntryCopiesRow, Campaign, CompendiumEntry, CompendiumSource, compendiumBack(), viewerCampaign(), sourceIssue(), Roster (+12 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Sum"
Cohesion: 0.13
Nodes (7): findItem(), parseArmor(), TestCustomMods(), Sum(), TestGearCommands(), ensureItem(), Item

### Community 59 - "Book"
Cohesion: 0.11
Nodes (27): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools(), fetchBooks() (+19 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "cell"
Cohesion: 0.13
Nodes (22): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+14 more)

### Community 63 - "BuilderPicks"
Cohesion: 0.83
Nodes (4): BuilderPickItems(), BuilderPicks(), BuilderPick, pickURL()

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.02
Nodes (57): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, applyCondOps() (+49 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (174): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+166 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "empty_templ.go"
Cohesion: 0.06
Nodes (61): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListImageCharactersRow (+53 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "radiogroup_templ.go"
Cohesion: 0.17
Nodes (15): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+7 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.17
Nodes (14): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveSourcesAlert(), TestLiveVisibility(), TestLiveMastery(), LiveKey() (+6 more)

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
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "ParseWikidot"
Cohesion: 0.13
Nodes (23): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+15 more)

### Community 86 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (80): CountCompendiumRow, CountHabitatsRow, ListHomebrewLinksRow, SearchCompendiumRow, Badge(), Toast(), Toaster(), toastIcon() (+72 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (37): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+29 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "Site"
Cohesion: 0.13
Nodes (14): Site, dataFor(), Kind, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey() (+6 more)

### Community 104 - "ftItemData"
Cohesion: 0.26
Nodes (5): ftClassFeature, ftClassRefs, ftCode(), featuresSearchText(), ftFeatureKey()

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (18): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+10 more)

### Community 108 - "Store"
Cohesion: 0.07
Nodes (41): fixture, tokens, main(), run(), migration, constantTimeEqual(), credentialsMatch(), newFixture() (+33 more)

### Community 110 - ".Router"
Cohesion: 0.15
Nodes (8): IsLoopback(), TestIsLoopback(), TestRequireAdmin(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 112 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 113 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.28
Nodes (6): UpsertCompendiumEntryParams, classCopy(), testEnv, TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 114 - ".loadLive"
Cohesion: 0.07
Nodes (24): SpellData, Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), initiativeMod(), outOfFight() (+16 more)

### Community 115 - "Weapon"
Cohesion: 0.13
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 116 - "PlanSourceSwitch"
Cohesion: 0.27
Nodes (5): featureSwitch, featureNames(), SourceRef, planSwitch(), sameSlug()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "ftItemData"
Cohesion: 0.14
Nodes (17): ArmorData, ftDistance, ftItem, o5eArmor, ItemData, ftAmount(), ftArmor(), ftDuration() (+9 more)

### Community 123 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (21): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eMastery(), TestUnflattenTable() (+13 more)

### Community 134 - "Import"
Cohesion: 0.12
Nodes (29): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+21 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 347 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **24 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Character`, `database/sql.NullInt64`, `.do`, `context.Context`, `Store`, `.Router`, `Client`, `Sheet`, `.loadLive`, `perm.go`, `CompendiumEntry`, `Book`?**
  _High betweenness centrality (0.062) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09811378488842262 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Parse`, `SheetView`, `Feature`, `WeaponKind`, `Input`, `ParseEffects`, `TestSourceSwitchKeepsEdits`, `.loadLive`, `Weapon`, `PlanSourceSwitch`, `CompendiumEntry`, `Sum`, `ParseSheet`?**
  _High betweenness centrality (0.024) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01152652578690705 - nodes in this community are weakly interconnected._
- **Why does `Entry` connect `open5e.go` to `compendium_templ.go`, `encoding/json.RawMessage`, `golang.org/x/net/html.Node`, `Site`, `parseFiveToolsBook`, `ftBook`, `.Entry`, `ParseWikidot`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `.do` be split into smaller, more focused modules?**
  _Cohesion score 0.11174242424242424 - nodes in this community are weakly interconnected._