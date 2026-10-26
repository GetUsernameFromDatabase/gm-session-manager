# Graph Report - gm-session-manager  (2026-10-26)

## Corpus Check
- 229 files · ~434,431 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 48 file(s) not represented in the graph (top: .templ 40, .css 4, (none) 3)

## Summary
- 3358 nodes · 13625 edges · 107 communities (86 shown, 21 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 1488 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `454ade56`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- Vendored shadcn-templ Bundle
- context.Context
- github.com/a-h/templ.Component
- github.com/a-h/templ.Attributes
- database/sql.NullInt64
- Character Sheet View
- levelUpForm
- Websocket Hub
- .do
- table_templ.go
- ImportView
- 5e Command Apply
- .Entry
- Props
- Combobox JS
- perm.go
- Select JS
- Floating UI DOM
- Hover Interaction JS
- DecodeData
- popover_templ.go
- Notes Editor JS
- combat_templ.go
- Anchor Positioning JS
- CompendiumView
- Project Rules (CLAUDE.md, README)
- who
- 5e Sheet Model
- setStat
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- .importUpload
- ClearCookie
- App JS Command Bar
- Dismiss Interaction JS
- ParseWikidot
- ParseAidedd
- Command Palette JS
- DB Open and Migrate
- select_templ.go
- Dialog JS
- cell
- Tooltip JS
- Campaign
- attachmentRefs
- 5e Level-Up and Slots
- 5e Command Parser
- tabs_templ.go
- Composite List Navigation
- wizardChoice
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- CustomMod
- golangci-lint Config
- renderMarkdown
- Markdown Export
- List Navigation JS
- vendorjs Tool
- findItem
- wikiText
- sqlc Config
- Transition Status JS
- Checkbox JS
- Claude Settings
- Switch JS
- Tabs JS
- Sed Hook
- SearchReadableBlocksRow
- TestMiddlewareIdentity
- Portal JS
- Click Interaction JS
- Focus Interaction JS
- nullID
- Mark Others (aria-hidden)
- Typeahead JS
- Live View Bindings JS
- testing.T
- Internal Backdrop JS
- Lifecycle JS
- builder_templ.go
- Lint and Release Rules
- Run Script
- d20 Favicon
- Validator Rule
- Go Module
- Entry
- encoding/json.RawMessage
- Import
- Entry
- App
- fromLookup
- ArmorData
- lanHostname

## God Nodes (most connected - your core abstractions)
1. `App` - 234 edges
2. `CN()` - 106 edges
3. `Command` - 100 edges
4. `Sheet` - 93 edges
5. `Button()` - 93 edges
6. `Entry` - 66 edges
7. `newTestEnv()` - 60 edges
8. `Hub` - 58 edges
9. `j()` - 58 edges
10. `Manager` - 53 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `NewCommand()` --calls--> `apply()`  [INFERRED]
  dnd5e/command.go → db/db.go
- `slotLabel()` --references--> `Slot`  [EXTRACTED]
  web/views/live.go → dnd5e/sheet.go
- `coinList()` --references--> `Coins`  [EXTRACTED]
  web/views/builder_templ.go → dnd5e/sheet.go
- `TestCommandPermissions()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (107 total, 21 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.07
Nodes (70): TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), looseString, o5eAssignment (+62 more)

### Community 1 - "Vendored shadcn-templ Bundle"
Cohesion: 0.07
Nodes (152): A(), Ae(), b(), be(), bt(), c(), Ce(), I() (+144 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (67): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Encounter, Note (+59 more)

### Community 3 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (128): ListMonstersByNameRow, ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Look, BlockView, CombatView, ImagesView (+120 more)

### Community 4 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (180): Props, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size (+172 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (54): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry (+46 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.08
Nodes (68): ArmorEffect(), Signed(), Field, ctxState, actionEntry, edit, subtitlePart, Action() (+60 more)

### Community 7 - "levelUpForm"
Cohesion: 0.08
Nodes (42): SpellData, LevelUpPlan, FeaturePreview, Option, WizardStep, Footer(), classBadges(), classBody() (+34 more)

### Community 8 - "Websocket Hub"
Cohesion: 0.06
Nodes (33): IsLoopback(), TestIsLoopback(), ClientInfo, event, fakeHandler, Handler, Client, Hub (+25 more)

### Community 9 - ".do"
Cohesion: 0.16
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 10 - "table_templ.go"
Cohesion: 0.14
Nodes (20): BodyProps, CaptionProps, CellProps, FooterProps, HeaderProps, HeadProps, Props, RowProps (+12 more)

### Community 11 - "ImportView"
Cohesion: 0.12
Nodes (20): Props, Device, RosterCharacter, Checkbox(), initialChecked(), Roster, rosterCard(), categoryRow() (+12 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.07
Nodes (51): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+43 more)

### Community 13 - ".Entry"
Cohesion: 0.11
Nodes (23): NameDesc, ValidKind(), HomebrewOwner, HomebrewInput, HomebrewInputFor(), InputForEntry(), SaveHomebrew(), actionsText() (+15 more)

### Community 14 - "Props"
Cohesion: 0.36
Nodes (6): Props, Size, ButtonProps, Type, Variant, Variants()

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (48): GetBlockRow, ListCombatantsRow, blockCase, Combatant, Fields, ParseFields(), VisibleEntity(), VisibleLooks() (+40 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "DecodeData"
Cohesion: 0.06
Nodes (26): DecodeData(), TextData, Slugify(), Character, Combatant, ClassFeature, Character, ParseSheet() (+18 more)

### Community 21 - "popover_templ.go"
Cohesion: 0.15
Nodes (22): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+14 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "combat_templ.go"
Cohesion: 0.08
Nodes (47): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Variant, combatantGroup (+39 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "CompendiumView"
Cohesion: 0.15
Nodes (21): CountHabitatsRow, SearchCompendiumRow, PickerTarget, alignmentColor(), CompendiumResults(), compendiumRow(), CompendiumRows(), CompendiumSearch() (+13 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "who"
Cohesion: 0.10
Nodes (8): ftsQuery(), Search(), compendiumBack(), viewerCampaign(), HomebrewView, LinkItem, SwitcherResults(), who

### Community 28 - "5e Sheet Model"
Cohesion: 0.07
Nodes (18): DeathSaves, TestModifier(), Part, skillByKey(), Sum(), TestDualWield(), TestUpgradeSheet(), damageLabel() (+10 more)

### Community 29 - "setStat"
Cohesion: 0.29
Nodes (8): addEntry(), setStat(), MonsterData, abilityKey(), parseBonuses(), parseSpeed(), TestMonsterFormParsers(), joinNonEmpty()

### Community 30 - "Scroll Lock JS"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 31 - "Popover JS"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 32 - "Tabbable Focus JS"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 33 - ".importUpload"
Cohesion: 0.29
Nodes (4): csvImport, csvImports, csvSource(), parseCSV()

### Community 34 - "ClearCookie"
Cohesion: 0.40
Nodes (3): ClearCookie(), SetCookie(), writeCookie()

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.16
Nodes (18): block, CoreTraits(), TestCoreTraits(), withCoreTraits(), featuresUnder(), ParseWikidot(), sourceLine(), tableText() (+10 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "DB Open and Migrate"
Cohesion: 0.11
Nodes (25): newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open(), openPool() (+17 more)

### Community 41 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Align, ContentProps, ctxKey, ctxState, GroupProps, ItemProps, LabelProps, Props (+16 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "cell"
Cohesion: 0.12
Nodes (23): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+15 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - "Campaign"
Cohesion: 0.19
Nodes (7): Campaign, PendingKind, CreateCampaignParams, deviceOS(), isPhone(), uaPick(), userAgentSummary()

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.12
Nodes (14): ClassLevel, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), ClassInfo (+6 more)

### Community 48 - "5e Command Parser"
Cohesion: 0.13
Nodes (18): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+10 more)

### Community 49 - "tabs_templ.go"
Cohesion: 0.22
Nodes (19): ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant (+11 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.15
Nodes (18): EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), EquipmentItem (+10 more)

### Community 52 - "Release Workflow"
Cohesion: 0.11
Nodes (18): CGO_ENABLED, GH_TOKEN, jobs, release, name, permissions, contents, tags (+10 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.18
Nodes (14): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+6 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "CustomMod"
Cohesion: 0.18
Nodes (11): findMod(), newMod(), parseMod(), SaveFeature(), UsesLabel(), FeatureEffects, CustomMod, Feature (+3 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 59 - "Markdown Export"
Cohesion: 0.28
Nodes (9): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+1 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "findItem"
Cohesion: 0.29
Nodes (5): findItem(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 63 - "wikiText"
Cohesion: 0.27
Nodes (8): WikiLink, WikiPart, wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks()

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Transition Status JS"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 66 - "Checkbox JS"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Switch JS"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 69 - "Tabs JS"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 72 - "TestMiddlewareIdentity"
Cohesion: 0.08
Nodes (17): constantTimeEqual(), credentialsMatch(), isCleared(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity(), TestRequireAdmin(), TestTokens() (+9 more)

### Community 73 - "Portal JS"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 74 - "Click Interaction JS"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 75 - "Focus Interaction JS"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 76 - "nullID"
Cohesion: 0.17
Nodes (12): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), liveFixture, LiveKey(), dial() (+4 more)

### Community 77 - "Mark Others (aria-hidden)"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 78 - "Typeahead JS"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 80 - "testing.T"
Cohesion: 0.08
Nodes (53): TestRealSource(), Import(), TestCanCharacter(), nullID(), TestBuilderCompendium(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName() (+45 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "builder_templ.go"
Cohesion: 0.07
Nodes (62): ItemData, builderOp, builderVerb, coin, profFocus, Close(), CloseFor(), Description() (+54 more)

### Community 99 - "Entry"
Cohesion: 0.14
Nodes (16): TestUnflattenTable(), fixture, fixtures, normalizer, actionRank(), byPK(), children(), decode() (+8 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (14): abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags() (+6 more)

### Community 102 - "Import"
Cohesion: 0.23
Nodes (13): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+5 more)

### Community 103 - "Entry"
Cohesion: 0.12
Nodes (22): ClassFeature, Entry, dataFor(), ClassData, Kind, fiveToolsPage(), storeSource(), upsert() (+14 more)

### Community 105 - "App"
Cohesion: 0.06
Nodes (19): FromContext(), Entity, App, blockAccessForm, coinCommand(), commandText(), hpCommand(), parseCommand() (+11 more)

### Community 108 - "fromLookup"
Cohesion: 0.29
Nodes (8): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc

### Community 110 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, ItemData, o5eArmor, armorData(), armorText()

### Community 120 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 319 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **21 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `.importUpload`, `context.Context`, `github.com/a-h/templ.Component`, `database/sql.NullInt64`, `Websocket Hub`, `TestMiddlewareIdentity`, `.do`, `Campaign`, `perm.go`, `DecodeData`, `TLS Cert and Main`, `who`?**
  _High betweenness centrality (0.056) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.0749593873288466 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `Character Sheet View`, `levelUpForm`, `5e Level-Up and Slots`, `5e Command Parser`, `perm.go`, `wizardChoice`, `DecodeData`, `builder_templ.go`, `CustomMod`, `findItem`?**
  _High betweenness centrality (0.026) - this node is a cross-community bridge._
- **Should `Vendored shadcn-templ Bundle` be split into smaller, more focused modules?**
  _Cohesion score 0.07461038961038961 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Websocket Hub` to `Auth and Compendium Core`, `App`?**
  _High betweenness centrality (0.023) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011698093050584904 - nodes in this community are weakly interconnected._