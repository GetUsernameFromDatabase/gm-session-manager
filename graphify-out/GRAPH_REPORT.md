# Graph Report - gm-session-manager  (2026-10-26)

## Corpus Check
- 228 files · ~426,663 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 48 file(s) not represented in the graph (top: .templ 40, .css 4, (none) 3)

## Summary
- 3340 nodes · 13475 edges · 109 communities (90 shown, 19 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 1479 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `a0c8edc9`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- Vendored shadcn-templ Bundle
- context.Context
- github.com/a-h/templ.Component
- shadcn Component Props
- database/sql.NullInt64
- Character Sheet View
- combobox_templ.go
- Websocket Hub
- popover_templ.go
- select_templ.go
- ImagePage
- 5e Command Apply
- .Entry
- .loadLive
- Combobox JS
- Viewer
- Select JS
- Floating UI DOM
- Hover Interaction JS
- Character
- tooltip_templ.go
- Notes Editor JS
- CombatantView
- Anchor Positioning JS
- CompendiumView
- Project Rules (CLAUDE.md, README)
- App
- 5e Sheet Model
- .do
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- WornArmor
- 5e Armor and Attack Explain
- App JS Command Bar
- Dismiss Interaction JS
- ParseWikidot
- ParseAidedd
- Command Palette JS
- DB Open and Migrate
- ImportView
- Dialog JS
- Spreadsheet Image Import
- Tooltip JS
- .renderCompendium
- testEnv
- 5e Level-Up and Slots
- 5e Command Parser
- IfElse
- Composite List Navigation
- wizardChoice
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- parseMod
- golangci-lint Config
- .startRequest
- Markdown Export
- List Navigation JS
- vendorjs Tool
- findItem
- dialog_templ.go
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
- Player
- Mark Others (aria-hidden)
- Typeahead JS
- Live View Bindings JS
- newTestEnv
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
- monsterBody
- Import
- Entry
- SlotMaxima
- App
- fromLookup
- Entry
- lanHostname

## God Nodes (most connected - your core abstractions)
1. `App` - 234 edges
2. `CN()` - 105 edges
3. `Command` - 100 edges
4. `Sheet` - 93 edges
5. `Button()` - 87 edges
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
- `TestCommandPermissions()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go
- `TestServerSideStripping()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go
- `newLiveFixture()` --calls--> `nullID()`  [INFERRED]
  web/live_test.go → perm/sql_agreement_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (109 total, 19 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.08
Nodes (65): TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), looseString, o5eAssignment (+57 more)

### Community 1 - "Vendored shadcn-templ Bundle"
Cohesion: 0.07
Nodes (152): A(), Ae(), b(), be(), bt(), c(), Ce(), I() (+144 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (68): keepAwake(), fetch(), ListImageCharactersRow, ListMonstersByNameRow, ListImageEntitiesRow, Condition, Note, Session (+60 more)

### Community 3 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (130): Look, BlockView, FeaturePreview, ImagesView, Option, PendingView, WizardStep, WS (+122 more)

### Community 4 - "shadcn Component Props"
Cohesion: 0.02
Nodes (150): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+142 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (40): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Attachment, Combatant, CompendiumSource, Character (+32 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.09
Nodes (59): UsesLabel(), ArmorEffect(), Slot, Signed(), actionEntry, edit, profFocus, subtitlePart (+51 more)

### Community 7 - "combobox_templ.go"
Cohesion: 0.08
Nodes (32): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+24 more)

### Community 8 - "Websocket Hub"
Cohesion: 0.06
Nodes (34): IsLoopback(), TestIsLoopback(), ClientInfo, event, fakeHandler, Handler, Client, Hub (+26 more)

### Community 9 - "popover_templ.go"
Cohesion: 0.16
Nodes (16): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+8 more)

### Community 10 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Align, ContentProps, ctxKey, ctxState, GroupProps, ItemProps, LabelProps, Props (+16 more)

### Community 11 - "ImagePage"
Cohesion: 0.21
Nodes (14): ListEntityFieldBacklinksRow, ListBacklinksRow, Props, Size, ImageView, initialChecked(), Switch(), EntityDetail() (+6 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.08
Nodes (46): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+38 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (33): addEntry(), setStat(), TestHomebrew(), CoreTraits(), TestCoreTraits(), MonsterData, NameDesc, ValidKind() (+25 more)

### Community 14 - ".loadLive"
Cohesion: 0.06
Nodes (21): Queries, Encounter, Entity, CreateCombatantParams, CreateEncounterParams, ListCharacterEncountersParams, applyCondOps(), uniqueName() (+13 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "Viewer"
Cohesion: 0.07
Nodes (49): ListCombatantsRow, blockCase, Combatant, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), ImageLook (+41 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "Character"
Cohesion: 0.07
Nodes (19): DecodeData(), Slugify(), Character, CompendiumEntry, Campaign(), CharacterFileName(), fileSlug(), CharacterFile (+11 more)

### Community 21 - "tooltip_templ.go"
Cohesion: 0.14
Nodes (26): Field, ContentProps, ctxKey, ctxState, Props, Side, PickerTarget, Content() (+18 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.10
Nodes (37): Props, Variant, combatantGroup, CombatantView, CombatView, Badge(), badgeClasses(), addCombatantForm() (+29 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "CompendiumView"
Cohesion: 0.16
Nodes (15): CountCompendiumRow, CountHabitatsRow, ListHomebrewLinksRow, CompendiumResults(), CompendiumRows(), CompendiumSearch(), CompendiumView, ImportNotice (+7 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "App"
Cohesion: 0.08
Nodes (11): FromContext(), ListAttachmentFilesRow, GetBlockRow, App, blockAccessLevel, compendiumBack(), parseCSV(), saveSniffed() (+3 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (14): parseSet(), DeathSaves, TestModifier(), skillByKey(), Sum(), TestUpgradeSheet(), Sheet, Modifier() (+6 more)

### Community 29 - ".do"
Cohesion: 0.17
Nodes (15): serve(), main(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+7 more)

### Community 30 - "Scroll Lock JS"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 31 - "Popover JS"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 32 - "Tabbable Focus JS"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 33 - "WornArmor"
Cohesion: 0.50
Nodes (4): Env, WornArmor, gear, applyEquipment()

### Community 34 - "5e Armor and Attack Explain"
Cohesion: 0.19
Nodes (5): Part, damageLabel(), Weapon, attackTargets(), weaponReach()

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.15
Nodes (19): block, TextData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), sourceLine() (+11 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "DB Open and Migrate"
Cohesion: 0.11
Nodes (24): newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open(), openPool() (+16 more)

### Community 41 - "ImportView"
Cohesion: 0.22
Nodes (11): ImageCategory, categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows() (+3 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "Spreadsheet Image Import"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - ".renderCompendium"
Cohesion: 0.11
Nodes (9): ftsQuery(), Search(), SearchCompendiumRow, Campaign, CreateCampaignParams, SearchCompendiumParams, viewerCampaign(), SearchComboItems() (+1 more)

### Community 46 - "testEnv"
Cohesion: 0.14
Nodes (10): TestRequireAdmin(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag(), testEnv, TestAdminLogin() (+2 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.15
Nodes (12): armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan (+4 more)

### Community 48 - "5e Command Parser"
Cohesion: 0.11
Nodes (20): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+12 more)

### Community 49 - "IfElse"
Cohesion: 0.20
Nodes (21): ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant (+13 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.19
Nodes (15): EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), EquipmentItem (+7 more)

### Community 52 - "Release Workflow"
Cohesion: 0.11
Nodes (18): CGO_ENABLED, GH_TOKEN, jobs, release, name, permissions, contents, tags (+10 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.13
Nodes (12): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), TestCertNames(), TestKeyFileIsPrivate() (+4 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "parseMod"
Cohesion: 0.67
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - ".startRequest"
Cohesion: 0.09
Nodes (13): ClearCookie(), SetCookie(), writeCookie(), PendingKind, deviceOS(), isPhone(), uaPick(), userAgentSummary() (+5 more)

### Community 59 - "Markdown Export"
Cohesion: 0.28
Nodes (9): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+1 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.23
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 62 - "findItem"
Cohesion: 0.19
Nodes (9): findItem(), parseArmor(), parseItem(), parseWeapon(), SaveWeapon(), TestGearCommands(), ensureItem(), Item (+1 more)

### Community 63 - "dialog_templ.go"
Cohesion: 0.18
Nodes (25): ctxKey, ctxState, Props, NoteListView, Dialog(), Close(), Content(), Description() (+17 more)

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
Cohesion: 0.22
Nodes (8): hooks, PreToolUse, permissions, allow, ask, defaultMode, deny, $schema

### Community 68 - "Switch JS"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 69 - "Tabs JS"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 72 - "TestMiddlewareIdentity"
Cohesion: 0.14
Nodes (10): constantTimeEqual(), credentialsMatch(), isCleared(), TestCredentialsMatch(), TestMiddlewareIdentity(), TestTokens(), Role, hashToken() (+2 more)

### Community 73 - "Portal JS"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 74 - "Click Interaction JS"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 75 - "Focus Interaction JS"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 76 - "Player"
Cohesion: 0.17
Nodes (12): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), liveFixture, LiveKey(), dial() (+4 more)

### Community 77 - "Mark Others (aria-hidden)"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 78 - "Typeahead JS"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 80 - "newTestEnv"
Cohesion: 0.07
Nodes (53): Import(), seedGarbage(), TestCanCharacter(), nullID(), TestBuilderCompendium(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName() (+45 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "builder_templ.go"
Cohesion: 0.06
Nodes (70): SaveFeature(), FeatureEffects, Coins, CustomMod, Feature, Uses, ItemData, builderOp (+62 more)

### Community 99 - "Entry"
Cohesion: 0.14
Nodes (16): TestUnflattenTable(), fixture, fixtures, normalizer, actionRank(), byPK(), children(), decode() (+8 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags() (+7 more)

### Community 101 - "monsterBody"
Cohesion: 0.20
Nodes (17): classBadges(), classBody(), commaPrefixed(), entryBody(), EntryDetails(), itemBody(), levelsText(), monsterBody() (+9 more)

### Community 102 - "Import"
Cohesion: 0.19
Nodes (14): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+6 more)

### Community 103 - "Entry"
Cohesion: 0.14
Nodes (13): dataFor(), Kind, fiveToolsPage(), Site, FetchPage(), Page, pageSource(), PageSourceKey() (+5 more)

### Community 104 - "SlotMaxima"
Cohesion: 0.50
Nodes (4): ClassLevel, TestSlotMaxima(), CasterLevel(), SlotMaxima()

### Community 105 - "App"
Cohesion: 0.06
Nodes (19): ParseSheet(), CanCharacter(), blockAccessForm, featureCommand(), classFeatures(), classInfo(), coinCommand(), commandText() (+11 more)

### Community 108 - "fromLookup"
Cohesion: 0.29
Nodes (8): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc

### Community 110 - "Entry"
Cohesion: 0.20
Nodes (15): ArmorData, ClassFeature, Entry, ClassData, ItemData, SpellData, importSource(), readSource() (+7 more)

### Community 120 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 320 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **19 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `database/sql.NullInt64`, `Import`, `Websocket Hub`, `App`, `.renderCompendium`, `.loadLive`, `testEnv`, `Viewer`, `Character`, `TLS Cert and Main`, `.startRequest`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.08227311280746395 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `WornArmor`, `5e Armor and Attack Explain`, `github.com/a-h/templ.Component`, `database/sql.NullInt64`, `Character Sheet View`, `SlotMaxima`, `App`, `5e Command Apply`, `5e Level-Up and Slots`, `5e Command Parser`, `Viewer`, `builder_templ.go`, `Character`, `wizardChoice`, `parseMod`, `findItem`?**
  _High betweenness centrality (0.031) - this node is a cross-community bridge._
- **Should `Vendored shadcn-templ Bundle` be split into smaller, more focused modules?**
  _Cohesion score 0.07461038961038961 - nodes in this community are weakly interconnected._
- **Why does `Command` connect `5e Command Apply` to `Auth and Compendium Core`, `WornArmor`, `App`, `5e Command Parser`, `Viewer`, `builder_templ.go`, `Character`, `parseMod`, `5e Sheet Model`, `findItem`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011671029668411867 - nodes in this community are weakly interconnected._