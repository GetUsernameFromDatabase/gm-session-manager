# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 249 files · ~502,711 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3619 nodes · 14781 edges · 125 communities (100 shown, 25 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1482 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `1c9977eb`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- .render
- context.Context
- .do
- createProvider
- Character
- SheetView
- combobox.js
- redirect.go
- newTestEnv
- github.com/gin-gonic/gin.Context
- select.js
- command.go
- .Entry
- Topic
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- github.com/a-h/templ.Component
- Hub
- wikilink.js
- dialog_templ.go
- popover.js
- wikiText
- Project Rules (CLAUDE.md, README)
- ParseSheet
- DecodeData
- Weapon
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- deviceOS
- command.js
- dialog.js
- State
- golang.org/x/net/html.Node
- tooltip.js
- Store
- .levelUp
- New
- export/markdown.go
- composite.js
- FromContext
- empty_templ.go
- .Plan
- tabs_templ.go
- Pending
- TestLiveUpdatesAndKick
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- reqLog
- golangci-lint Config
- NameDesc
- .joinInfo
- radiogroup.js
- vendorjs Tool
- select_templ.go
- Handoff
- sqlc Config
- Spell
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
- checkbox.js
- ImportView
- switch.js
- nullID
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- ParseWikidot
- Lint and Release Rules
- Run Script
- formSelect
- useTypeahead
- ImagesPage
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Client
- Site
- RosterCharacter
- IsLoopback
- .ListCompendiumByKind
- Entry
- Import
- .ListSpellsForClass
- .wizardCommand
- fieldError
- radiogroup_templ.go
- Sheet
- capReader
- build.sh
- spellCounts
- SearchReadableBlocksRow
- errBodyWriter
- testing.T
- open5e.go
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 267 edges
2. `CN()` - 112 edges
3. `Button()` - 103 edges
4. `Sheet` - 90 edges
5. `newTestEnv()` - 89 edges
6. `SheetView` - 52 edges
7. `IfElse()` - 51 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 39 edges

## Surprising Connections (you probably didn't know these)
- `Next up (pick from these)` --references--> `ParseDefenses()`  [INFERRED]
  HANDOFF.md → internal/dnd5e/sheet.go
- `Where things live (creation and level-up)` --references--> `userError`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Handoff` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go
- `Lessons (read before orchestrating)` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (125 total, 25 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (44): contextKey, keepAwake(), ASIError, clientMessage, TestMain(), TestMain(), lastTopLevelColon(), twImportantToV3() (+36 more)

### Community 1 - ".render"
Cohesion: 0.10
Nodes (12): Campaign, compendiumBack(), viewerCampaign(), isHTMX(), imageItem(), imageVersion(), Roster, HomebrewView (+4 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (70): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Encounter, Note (+62 more)

### Community 3 - ".do"
Cohesion: 0.12
Nodes (19): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+11 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Character"
Cohesion: 0.07
Nodes (23): Queries, Character, Combatant, Character, Campaign(), CharacterFileName(), fileSlug(), CharacterFile (+15 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (90): ArmorEffect(), Signed(), Field, Button(), Variant, Action(), Card(), Content() (+82 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "newTestEnv"
Cohesion: 0.09
Nodes (34): TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestConditionExpiresOnTurn(), TestDeleteRestoreImport(), TestRenameDeleteCampaign(), bugForm(), testEnv (+26 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (14): SetCookie(), coinCommand(), commandText(), entityTopics(), formInt(), isPhone(), redirect(), csvSource() (+6 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.05
Nodes (74): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+66 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (28): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInput, HomebrewInputFor() (+20 more)

### Community 14 - "Topic"
Cohesion: 0.14
Nodes (6): event, fakeHandler, Handler, stateHandler, TopicKind, Topic

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (45): GetBlockRow, ListCombatantsRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanCharacter() (+37 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (109): SpellData, SearchCompendiumRow, Alert(), Description(), Label(), Scripts(), Separator(), Spinner() (+101 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.11
Nodes (32): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+24 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.16
Nodes (25): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp() (+17 more)

### Community 28 - "DecodeData"
Cohesion: 0.15
Nodes (6): DecodeData(), CompendiumEntry, compendiumWeapon(), CreateHomebrewParams, GetCompendiumEntryByNameParams, GetCompendiumEntryBySlugParams

### Community 29 - "Weapon"
Cohesion: 0.13
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

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
Nodes (44): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumSource, Condition (+36 more)

### Community 34 - "deviceOS"
Cohesion: 0.50
Nodes (4): deviceOS(), formIntQuery(), uaPick(), userAgentSummary()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "State"
Cohesion: 0.24
Nodes (7): patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff(), liveView

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.11
Nodes (15): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity() (+7 more)

### Community 41 - ".levelUp"
Cohesion: 0.10
Nodes (33): Where things live (creation and level-up), ClassFeature, ClassInfo, LevelUpPlan, addClassSpells(), classFeatures(), classInfo(), errorField() (+25 more)

### Community 42 - "New"
Cohesion: 0.44
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "FromContext"
Cohesion: 0.09
Nodes (15): ClearCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), Toast(), browserName(), firstLine() (+7 more)

### Community 46 - "empty_templ.go"
Cohesion: 0.06
Nodes (65): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListEntityFieldBacklinksRow (+57 more)

### Community 47 - ".Plan"
Cohesion: 0.17
Nodes (8): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), hpGain(), raise(), CasterLevel(), SlotMaxima()

### Community 48 - "tabs_templ.go"
Cohesion: 0.09
Nodes (45): CountHabitatsRow, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+37 more)

### Community 49 - "Pending"
Cohesion: 0.25
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 50 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

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
Cohesion: 0.11
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "NameDesc"
Cohesion: 0.17
Nodes (16): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, catchUpResistances(), TestCatchUpResistances() (+8 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "select_templ.go"
Cohesion: 0.10
Nodes (24): RadioGroup(), Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator() (+16 more)

### Community 63 - "Handoff"
Cohesion: 0.22
Nodes (8): Handoff, Improvements (bigger, deferred), Lessons (read before orchestrating), Notes, Test data in data-test (all safe to delete), Working on the test instance, readPicks(), TestReadPicks()

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Spell"
Cohesion: 0.17
Nodes (15): Cantrips(), TextData, Slugify(), Spell, textMarkdown(), addProfs(), choiceOptions(), featList() (+7 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (170): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+162 more)

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
Cohesion: 0.04
Nodes (115): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+107 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.09
Nodes (30): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+22 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "nullID"
Cohesion: 0.12
Nodes (20): Player(), nullID(), TestCombatantNamesFollowSource(), TestMonsterPortrait(), testEnv, TestAttachmentAccess(), TestAttachmentToImage(), TestFilesCleanup() (+12 more)

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
Cohesion: 0.21
Nodes (12): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), sourceLine() (+4 more)

### Community 86 - "formSelect"
Cohesion: 0.10
Nodes (36): Footer(), attrString(), generateSVG(), getIconContent(), Props, Icon(), Item(), selectItems() (+28 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ImagesPage"
Cohesion: 0.18
Nodes (20): ImageCategory, categoryFilterItems(), categoryList(), displayControls(), DisplayURL(), ImageItem, imageAction(), imageBindForm() (+12 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): ftEntry, ftMonster, ftSpellcasting, ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType() (+7 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Client"
Cohesion: 0.22
Nodes (3): message, Client, topicIDs()

### Community 103 - "Site"
Cohesion: 0.12
Nodes (16): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+8 more)

### Community 104 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 107 - "Entry"
Cohesion: 0.27
Nodes (9): Entry, fixtures, normalizer, byPK(), children(), joinText(), label(), parseFixtures() (+1 more)

### Community 108 - "Import"
Cohesion: 0.12
Nodes (25): migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open() (+17 more)

### Community 110 - ".wizardCommand"
Cohesion: 0.08
Nodes (29): EquipmentItem, FeatureEffects, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions() (+21 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "Sheet"
Cohesion: 0.09
Nodes (13): DeathSaves, TestCustomMods(), TestResistances(), Sum(), TestGearCommands(), Coins, Item, Sheet (+5 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "testing.T"
Cohesion: 0.05
Nodes (55): Defense, CoreTraits(), TestCoreTraits(), InputForEntry(), TestMonsterFormRoundTrip(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground() (+47 more)

### Community 133 - "open5e.go"
Cohesion: 0.07
Nodes (30): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+22 more)

### Community 134 - "Import"
Cohesion: 0.11
Nodes (26): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+18 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 335 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **25 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `github.com/gin-gonic/gin.Context` to `compendium_templ.go`, `.render`, `context.Context`, `.do`, `Character`, `Import`, `perm.go`, `Hub`, `DecodeData`, `database/sql.NullInt64`, `deviceOS`, `State`, `.levelUp`, `FromContext`, `LoadOrCreateCert`, `reqLog`, `.joinInfo`, `Spell`, `.wizardCommand`?**
  _High betweenness centrality (0.033) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10553410553410554 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Spell`, `Character`, `SheetView`, `builder_templ.go`, `.levelUp`, `command.go`, `.wizardCommand`, `.Plan`, `perm.go`, `testing.T`, `NameDesc`, `ParseSheet`, `DecodeData`, `Weapon`?**
  _High betweenness centrality (0.025) - this node is a cross-community bridge._
- **Should `.render` be split into smaller, more focused modules?**
  _Cohesion score 0.10256410256410256 - nodes in this community are weakly interconnected._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011256281407035176 - nodes in this community are weakly interconnected._
- **Should `.do` be split into smaller, more focused modules?**
  _Cohesion score 0.12413793103448276 - nodes in this community are weakly interconnected._