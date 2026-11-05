# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 286 files · ~538,027 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 3900 nodes · 16710 edges · 129 communities (105 shown, 24 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1663 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `a5c4b766`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- FromContext
- context.Context
- .do
- createProvider
- CombatantView
- SheetView
- combobox.js
- redirect.go
- Parse
- CustomMod
- select.js
- Command
- .Entry
- NoteView
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ImportView
- levelUpForm
- wikilink.js
- github.com/a-h/templ.Component
- popover.js
- tooltip_templ.go
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- testing.T
- app.js
- tabbable.js
- useDismiss
- ParseSheet
- TestClassOptionChoices
- command.js
- dialog.js
- .importCommit
- golang.org/x/net/html.Node
- tooltip.js
- dialog_templ.go
- Sheet
- capReader
- export/markdown.go
- composite.js
- Queries
- Client
- ClassInfo
- Client
- Env
- parseHP
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- .encounterView
- golangci-lint Config
- CampaignRoster
- Book
- radiogroup.js
- vendorjs Tool
- radiogroup_templ.go
- Pending
- sqlc Config
- Condition
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
- checkbox.js
- cell
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
- toast_templ.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- IsLoopback
- Site
- ftItemData
- spellCounts
- parseMastery
- ftBook
- Import
- .ListBlockStepsSince
- TestSourceSwitchKeepsEdits
- State
- CompendiumEntry
- NoteListView
- .ListItemsByCategory
- Part
- .ListSpellsForClass
- .ListCompendiumByKind
- build.sh
- .ListCompendiumBySlugs
- Item
- errBodyWriter
- Verbs
- Manager
- open5e.go
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 287 edges
2. `Button()` - 112 edges
3. `CN()` - 112 edges
4. `Sheet` - 106 edges
5. `newTestEnv()` - 103 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 52 edges
8. `Command` - 47 edges
9. `Input()` - 46 edges
10. `Slugify()` - 40 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Open()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/db/db.go
- `run()` --calls--> `NewStore()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/db/store.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (129 total, 24 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (49): contextKey, keepAwake(), classRef, ftBookInfo, indexed, sourceOf, ASIError, clientMessage (+41 more)

### Community 1 - "FromContext"
Cohesion: 0.04
Nodes (28): Identity, FromContext(), RequireAdmin(), Attachment, Campaign, Session, browserName(), firstLine() (+20 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (93): keepAwake(), ftsQuery(), Search(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow (+85 more)

### Community 3 - ".do"
Cohesion: 0.16
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "CombatantView"
Cohesion: 0.11
Nodes (34): Header(), Title(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), combatCondition() (+26 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (80): ArmorEffect(), TestDualWield(), damageLabel(), Coins, Weapon, Signed(), Variant, Action() (+72 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Parse"
Cohesion: 0.07
Nodes (30): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestDefenses() (+22 more)

### Community 10 - "CustomMod"
Cohesion: 0.11
Nodes (16): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+8 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.13
Nodes (40): CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Command, CondOp (+32 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (27): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+19 more)

### Community 14 - "NoteView"
Cohesion: 0.09
Nodes (33): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, initials(), removeButton(), currentLook(), EntityDetail() (+25 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.15
Nodes (27): BlockVisibilities(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock(), CanManageNote(), CanReadAttachment(), CanReadBlock() (+19 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ImportView"
Cohesion: 0.18
Nodes (14): Props, ImageCategory, Checkbox(), initialChecked(), categoryRow(), columnLabel(), columnOverview(), ImportCategory (+6 more)

### Community 21 - "levelUpForm"
Cohesion: 0.13
Nodes (30): Footer(), Label(), abilitySelect(), levelUpForm(), slotSummary(), spellChecklist(), taughtPicked(), taughtRules() (+22 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (106): ListMonstersByNameRow, Description(), Input(), Content(), Description(), Empty(), Header(), Media() (+98 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "tooltip_templ.go"
Cohesion: 0.10
Nodes (33): SearchCompendiumRow, Field, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems() (+25 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.18
Nodes (19): TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestPrepareCapCommand() (+11 more)

### Community 28 - "App"
Cohesion: 0.05
Nodes (30): ClearCookie(), SetCookie(), writeCookie(), GetImageRow, CanCharacter(), coinCommand(), commandText(), compendiumBack() (+22 more)

### Community 29 - "testing.T"
Cohesion: 0.04
Nodes (96): TestMasteryLimit(), TestCanCharacter(), nullID(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+88 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "ParseSheet"
Cohesion: 0.11
Nodes (14): Character, Completions(), ParseSheet(), Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile() (+6 more)

### Community 34 - "TestClassOptionChoices"
Cohesion: 0.12
Nodes (16): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsEdition(), TestParseFiveToolsBook(), feature(), fiveToolsClasses(), TestFiveToolsClasses() (+8 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - ".importCommit"
Cohesion: 0.33
Nodes (4): ensureCategory(), nullString(), csvImport, csvImports

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "dialog_templ.go"
Cohesion: 0.11
Nodes (48): Props, Size, ctxKey, ctxState, Props, Button(), Variants(), Dialog() (+40 more)

### Community 41 - "Sheet"
Cohesion: 0.04
Nodes (101): EquipmentItem, DeathSaves, LevelUpChoice, Cantrips(), EquipmentOptions(), firstTable(), EquipmentOption, LevelSpells() (+93 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Queries"
Cohesion: 0.06
Nodes (24): TopicKind, Queries, Entity, GetBlockRow, Topic, Fields, ParseFields(), applyCondOps() (+16 more)

### Community 46 - "Client"
Cohesion: 0.18
Nodes (15): ClientInfo, event, Hub, New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter() (+7 more)

### Community 47 - "ClassInfo"
Cohesion: 0.12
Nodes (10): TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), hpGain(), raise() (+2 more)

### Community 48 - "Client"
Cohesion: 0.14
Nodes (4): fakeHandler, message, Client, topicIDs()

### Community 49 - "Env"
Cohesion: 0.19
Nodes (6): featureSwitch, Env, featureNames(), SourceRef, planSwitch(), sameSlug()

### Community 50 - "parseHP"
Cohesion: 0.39
Nodes (7): Amount, Damage(), Heal(), parseAmount(), parseDamage(), parseHP(), Set()

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
Cohesion: 0.10
Nodes (18): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+10 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - ".encounterView"
Cohesion: 0.32
Nodes (7): VisibleEntity(), VisibleLooks(), TestVisibleEntity(), TestVisibleLooks(), combatantKind(), combatantLink(), playerCombatantLink()

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "CampaignRoster"
Cohesion: 0.43
Nodes (7): CampaignRoster(), Roster, rosterCard(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 59 - "Book"
Cohesion: 0.12
Nodes (25): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools(), fetchBooks() (+17 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "radiogroup_templ.go"
Cohesion: 0.19
Nodes (14): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+6 more)

### Community 63 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Condition"
Cohesion: 0.11
Nodes (15): Combatant, Condition, ListCombatantsRow, combatantDefenses(), damageChoice(), initiativeMod(), outOfFight(), pcSheet() (+7 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (178): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+170 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "builder_templ.go"
Cohesion: 0.05
Nodes (80): Defenses, Combobox(), Content(), Item(), List(), Description(), Field(), Group() (+72 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "cell"
Cohesion: 0.11
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

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

### Community 83 - "ParseWikidot"
Cohesion: 0.10
Nodes (28): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, TestTokensReadAsLabels(), dnd5eFeatures(), TestFiveToolsClassLevelUp() (+20 more)

### Community 86 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (91): Props, Variant, CountHabitatsRow, Badge(), badgeClasses(), Content(), contentID(), IDFromContext() (+83 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.09
Nodes (34): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+26 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "IsLoopback"
Cohesion: 0.29
Nodes (3): IsLoopback(), TestIsLoopback(), sameOriginOnly()

### Community 103 - "Site"
Cohesion: 0.14
Nodes (13): Site, Source, fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page, pageSource() (+5 more)

### Community 104 - "ftItemData"
Cohesion: 0.10
Nodes (20): ArmorData, ftClassFeature, ftClassRefs, ftDistance, ftItem, ItemData, ftAmount(), ftArmor() (+12 more)

### Community 105 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 106 - "parseMastery"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 107 - "ftBook"
Cohesion: 0.17
Nodes (18): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+10 more)

### Community 108 - "Import"
Cohesion: 0.11
Nodes (20): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+12 more)

### Community 110 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.28
Nodes (6): UpsertCompendiumEntryParams, classCopy(), testEnv, TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 111 - "State"
Cohesion: 0.14
Nodes (9): Handler, patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff() (+1 more)

### Community 112 - "CompendiumEntry"
Cohesion: 0.10
Nodes (15): ListEntryCopiesRow, CompendiumEntry, pinnedRefs(), sourceIssue(), copyLabel(), SourceCopy, SourceIssue, SourceLabel() (+7 more)

### Community 115 - "Part"
Cohesion: 0.13
Nodes (6): Part, Skill, TestModifier(), skillByKey(), Modifier(), ProficiencyAt()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 123 - "Item"
Cohesion: 0.26
Nodes (5): findItem(), parseArmor(), TestGearCommands(), ensureItem(), Item

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 131 - "Manager"
Cohesion: 0.12
Nodes (16): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+8 more)

### Community 133 - "open5e.go"
Cohesion: 0.08
Nodes (36): Entry, fixture, fixtures, looseString, normalizer, o5eArmor, o5eAssignment, o5eChild (+28 more)

### Community 134 - "Import"
Cohesion: 0.12
Nodes (32): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+24 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 347 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **24 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `ParseSheet`, `context.Context`, `.do`, `FromContext`, `Condition`, `.importCommit`, `Sheet`, `Import`, `Queries`, `Client`, `State`, `perm.go`, `CompendiumEntry`, `cert.go`, `github.com/a-h/templ.Component`, `.encounterView`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11273512963904422 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `ParseSheet`, `Condition`, `SheetView`, `Parse`, `CustomMod`, `parseMastery`, `TestSourceSwitchKeepsEdits`, `ClassInfo`, `CompendiumEntry`, `Env`, `Part`, `Item`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `FromContext` be split into smaller, more focused modules?**
  _Cohesion score 0.04480874316939891 - nodes in this community are weakly interconnected._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010079656862745099 - nodes in this community are weakly interconnected._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._