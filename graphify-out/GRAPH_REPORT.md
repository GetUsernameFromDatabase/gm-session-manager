# Graph Report - gm-session-manager  (2026-11-04)

## Corpus Check
- 247 files · ~495,093 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3586 nodes · 14492 edges · 114 communities (96 shown, 18 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1451 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `5b06b237`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- .do
- createProvider
- .encounterView
- SheetView
- combobox.js
- redirect.go
- testing.T
- App
- select.js
- command.go
- .Entry
- Sheet
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Defenses
- wikiText
- wikilink.js
- CombatantView
- popover.js
- time.Time
- Project Rules (CLAUDE.md, README)
- ParseSheet
- Campaign
- Page
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- deviceOS
- command.js
- dialog.js
- Weapon
- golang.org/x/net/html.Node
- tooltip.js
- Store
- .levelUp
- Slugify
- export/markdown.go
- composite.js
- .levelUp
- github.com/a-h/templ.Component
- .Plan
- IfElse
- Session
- .Router
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- golangci-lint Config
- NewSheet
- monsterBody
- radiogroup.js
- vendorjs Tool
- Feature
- sqlc Config
- Client
- Browser check
- Claude Settings
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
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
- cell
- Lint and Release Rules
- Run Script
- useTypeahead
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- github.com/a-h/templ.Attributes
- .joinInfo
- Entry
- Import
- SearchReadableBlocksRow
- ParseWikidot
- .wizardCommand
- radiogroup_templ.go
- build.sh
- .sendFeedback
- open5e.go
- capReader

## God Nodes (most connected - your core abstractions)
1. `App` - 266 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Sheet` - 89 edges
5. `newTestEnv()` - 86 edges
6. `SheetView` - 51 edges
7. `IfElse()` - 50 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 36 edges

## Surprising Connections (you probably didn't know these)
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Where things live (creation and level-up)` --references--> `choiceGroup`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Where things live (creation and level-up)` --references--> `readPicks()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Where things live (creation and level-up)` --references--> `lostPicks()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go
- `Where things live (creation and level-up)` --references--> `wizardLevelUp()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (114 total, 18 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.09
Nodes (46): contextKey, clientMessage, TestMain(), Toast(), toastIcon(), TestMain(), lastTopLevelColon(), twImportantToV3() (+38 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.17
Nodes (19): Content(), Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition() (+11 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (74): keepAwake(), keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, ListAttachmentFilesRow (+66 more)

### Community 3 - ".do"
Cohesion: 0.17
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - ".encounterView"
Cohesion: 0.18
Nodes (11): Combatant, ListCombatantsRow, Character, combatantDefenses(), initiativeMod(), outOfFight(), pcSheet(), stepTurn() (+3 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (94): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+86 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.06
Nodes (59): TestCanCharacter(), nullID(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+51 more)

### Community 10 - "App"
Cohesion: 0.05
Nodes (29): ClearCookie(), SetCookie(), writeCookie(), ftsQuery(), Search(), Campaign, coinCommand(), commandText() (+21 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.07
Nodes (64): Amount, CondOpKind, Effects, Env, HelpEntry, parser, Completions(), condCommand() (+56 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+22 more)

### Community 14 - "Sheet"
Cohesion: 0.09
Nodes (18): DeathSaves, TestCustomMods(), Sum(), TestGearCommands(), Item, Sheet, WornArmor, catchUpResistances() (+10 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (48): GetBlockRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+40 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Defenses"
Cohesion: 0.13
Nodes (11): Defense, Handoff, Improvements (bigger, deferred), Next up (pick from these), Notes, Test data in data-test (all safe to delete), Working on the test instance, Defenses (+3 more)

### Community 21 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.13
Nodes (29): addCombatantForm(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), downBadge(), Encounter() (+21 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "time.Time"
Cohesion: 0.08
Nodes (15): Identity, FromContext(), browserName(), firstLine(), mdLinkText(), deviceKey(), keepsBody(), Toast() (+7 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.16
Nodes (21): TestUpgradeSheet(), ParseSheet(), TestLevelUpFeats(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestUpgradeSheets() (+13 more)

### Community 28 - "Campaign"
Cohesion: 0.32
Nodes (5): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

### Community 29 - "Page"
Cohesion: 0.06
Nodes (66): Label(), RadioGroup(), Scripts(), Separator(), Textarea(), DirectionAttributes(), campaignLink(), CampaignOverview() (+58 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (48): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, GetImageRow, ListImagesRow, Attachment, Condition, Entity (+40 more)

### Community 34 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Weapon"
Cohesion: 0.13
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.13
Nodes (23): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+15 more)

### Community 41 - ".levelUp"
Cohesion: 0.10
Nodes (32): Where things live (creation and level-up), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, SpellData (+24 more)

### Community 42 - "Slugify"
Cohesion: 0.18
Nodes (15): Cantrips(), TextData, Slugify(), Spell, textMarkdown(), applyPicks(), choiceOptions(), featList() (+7 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".levelUp"
Cohesion: 0.05
Nodes (20): DecodeData(), Character, CompendiumEntry, CompendiumSource, CanCharacter(), compendiumWeapon(), errorField(), petrified() (+12 more)

### Community 46 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (99): ListMonstersByNameRow, ListEntityFieldBacklinksRow, ListBacklinksRow, Look, Checkbox(), Content(), Description(), Empty() (+91 more)

### Community 47 - ".Plan"
Cohesion: 0.10
Nodes (18): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+10 more)

### Community 48 - "IfElse"
Cohesion: 0.12
Nodes (39): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+31 more)

### Community 49 - "Session"
Cohesion: 0.11
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+10 more)

### Community 50 - ".Router"
Cohesion: 0.16
Nodes (7): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "LoadOrCreateCert"
Cohesion: 0.14
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "NewSheet"
Cohesion: 0.12
Nodes (20): Parse(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves(), TestLanguages(), TestLevelUp() (+12 more)

### Community 59 - "monsterBody"
Cohesion: 0.06
Nodes (60): Props, Variant, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), alignmentColor(), classBadges() (+52 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Feature"
Cohesion: 0.17
Nodes (8): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, addFeat(), addProfs(), applyBackground()

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Client"
Cohesion: 0.05
Nodes (37): ClientInfo, event, fakeHandler, Handler, message, patchMessage, PendingState, stateHandler (+29 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

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
Nodes (96): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+88 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "select_templ.go"
Cohesion: 0.06
Nodes (44): ImageCategory, Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select() (+36 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.16
Nodes (13): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), TestPlayerCombatantLink(), LiveKey(), dial() (+5 more)

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

### Community 83 - "cell"
Cohesion: 0.11
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.16
Nodes (18): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+10 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 103 - "Site"
Cohesion: 0.12
Nodes (17): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+9 more)

### Community 104 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (155): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+147 more)

### Community 107 - "Entry"
Cohesion: 0.27
Nodes (9): Entry, fixtures, normalizer, byPK(), children(), joinText(), label(), parseFixtures() (+1 more)

### Community 108 - "Import"
Cohesion: 0.16
Nodes (20): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+12 more)

### Community 109 - "SearchReadableBlocksRow"
Cohesion: 0.50
Nodes (3): SearchReadableBlocksRow, SearchReadableBlocksParams, NoteListView

### Community 111 - "ParseWikidot"
Cohesion: 0.15
Nodes (19): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+11 more)

### Community 113 - ".wizardCommand"
Cohesion: 0.18
Nodes (15): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestCoreTraits() (+7 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 133 - "open5e.go"
Cohesion: 0.07
Nodes (30): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+22 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 334 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **18 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Client`, `context.Context`, `.do`, `database/sql.NullInt64`, `.encounterView`, `.levelUp`, `Slugify`, `.joinInfo`, `.levelUp`, `.Plan`, `perm.go`, `.wizardCommand`, `.Router`, `LoadOrCreateCert`, `time.Time`, `Campaign`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09454585807024357 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `.encounterView`, `Weapon`, `SheetView`, `builder_templ.go`, `.levelUp`, `Slugify`, `command.go`, `.levelUp`, `.Plan`, `.wizardCommand`, `Defenses`, `NewSheet`, `ParseSheet`, `Page`, `Feature`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010866415531729933 - nodes in this community are weakly interconnected._
- **Why does `Handoff` connect `Defenses` to `.levelUp`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._