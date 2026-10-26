# Graph Report - gm-session-manager  (2026-10-26)

## Corpus Check
- 229 files · ~429,246 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 48 file(s) not represented in the graph (top: .templ 40, .css 4, (none) 3)

## Summary
- 3347 nodes · 13541 edges · 108 communities (86 shown, 22 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 1486 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `dc8dbf65`
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
- dialog_templ.go
- Websocket Hub
- popover_templ.go
- IfElse
- .loadLive
- 5e Command Apply
- .Entry
- .commandFailed
- Combobox JS
- Viewer
- Select JS
- Floating UI DOM
- Hover Interaction JS
- Character
- .blockFor
- Notes Editor JS
- CombatantView
- Anchor Positioning JS
- CompendiumView
- Project Rules (CLAUDE.md, README)
- .fail
- 5e Sheet Model
- .encounterView
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- NewCommand
- 5e Armor and Attack Explain
- App JS Command Bar
- Dismiss Interaction JS
- ParseWikidot
- ParseAidedd
- Command Palette JS
- DB Open and Migrate
- select_templ.go
- Dialog JS
- Spreadsheet Image Import
- Tooltip JS
- .renderCompendium
- fingerprintPairs
- 5e Level-Up and Slots
- 5e Command Parser
- IfElse
- Composite List Navigation
- wizardChoice
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- parseCommand
- golangci-lint Config
- Campaign
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
- Session
- Portal JS
- Click Interaction JS
- Focus Interaction JS
- Player
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
- renderMarkdown
- Import
- Entry
- App
- fromLookup
- open5e.go
- lanHostname

## God Nodes (most connected - your core abstractions)
1. `App` - 234 edges
2. `CN()` - 105 edges
3. `Command` - 100 edges
4. `Sheet` - 93 edges
5. `Button()` - 89 edges
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
- `newLiveFixture()` --calls--> `nullID()`  [INFERRED]
  web/live_test.go → perm/sql_agreement_test.go
- `TestLiveRefetchStrips()` --calls--> `nullID()`  [INFERRED]
  web/live_test.go → perm/sql_agreement_test.go
- `TestPlayerCombatantLink()` --calls--> `nullID()`  [INFERRED]
  web/turn_test.go → perm/sql_agreement_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (108 total, 22 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.08
Nodes (63): Props, TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), Config (+55 more)

### Community 1 - "Vendored shadcn-templ Bundle"
Cohesion: 0.07
Nodes (152): A(), Ae(), b(), be(), bt(), c(), Ce(), I() (+144 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (63): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListMonstersByNameRow, ListImageEntitiesRow, Encounter (+55 more)

### Community 3 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (125): LevelUpPlan, BlockView, CombatView, FeaturePreview, NoteListView, Option, PendingView, WizardStep (+117 more)

### Community 4 - "shadcn Component Props"
Cohesion: 0.02
Nodes (161): Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props (+153 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (43): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumSource, Condition (+35 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.07
Nodes (81): ArmorEffect(), Slot, Signed(), Field, ContentProps, ctxKey, ctxState, Props (+73 more)

### Community 7 - "dialog_templ.go"
Cohesion: 0.10
Nodes (45): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+37 more)

### Community 8 - "Websocket Hub"
Cohesion: 0.06
Nodes (32): ClientInfo, event, fakeHandler, Handler, Client, Hub, Topic, New() (+24 more)

### Community 9 - "popover_templ.go"
Cohesion: 0.15
Nodes (22): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+14 more)

### Community 10 - "IfElse"
Cohesion: 0.12
Nodes (31): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Look, ImagesView, ImageView, IfElse(), characterLooks() (+23 more)

### Community 11 - ".loadLive"
Cohesion: 0.11
Nodes (12): Fields, ParseFields(), RosterCharacter, saveFieldLinks(), notFoundIsNil(), parseLiveKey(), liveKey, toLook() (+4 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.09
Nodes (46): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+38 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (32): addEntry(), setStat(), TestHomebrew(), MonsterData, NameDesc, ValidKind(), HomebrewOwner, HomebrewInput (+24 more)

### Community 14 - ".commandFailed"
Cohesion: 0.11
Nodes (10): Queries, Combatant, Character, uniqueName(), linkEntityLook(), formInt(), linkCharacterLook(), lookFree() (+2 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "Viewer"
Cohesion: 0.21
Nodes (25): blockCase, Combatant, VisibleEntity(), VisibleLooks(), ImageLook, BlockVisibilities(), CanDeleteCharacter(), CanEditBlock() (+17 more)

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
Cohesion: 0.09
Nodes (17): DecodeData(), TextData, Slugify(), CompendiumEntry, ClassFeature, ParseSheet(), SizeNamed(), CreateHomebrewParams (+9 more)

### Community 21 - ".blockFor"
Cohesion: 0.18
Nodes (3): GetBlockRow, blockAccessLevel, stepsPayload

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.08
Nodes (44): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Variant, combatantGroup (+36 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "CompendiumView"
Cohesion: 0.16
Nodes (19): CountHabitatsRow, SearchCompendiumRow, PickerTarget, CompendiumResults(), CompendiumRows(), CompendiumSearch(), CompendiumView, ImportNotice (+11 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - ".fail"
Cohesion: 0.09
Nodes (10): Character, CanCharacter(), SetCharacterCurrentImageParams, ToggleCharacterImageRevealedParams, UnlinkCharacterImageParams, compendiumBack(), HomebrewView, who (+2 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.09
Nodes (11): DeathSaves, TestModifier(), skillByKey(), Sum(), TestUpgradeSheet(), Sheet, IsConditionSlug(), Modifier() (+3 more)

### Community 29 - ".encounterView"
Cohesion: 0.27
Nodes (10): ListCombatantsRow, initiativeMod(), outOfFight(), pcSheet(), stepTurn(), combatantKind(), combatantLink(), playerCombatantLink() (+2 more)

### Community 30 - "Scroll Lock JS"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 31 - "Popover JS"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 32 - "Tabbable Focus JS"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 33 - "NewCommand"
Cohesion: 0.18
Nodes (8): EquipmentOption, EquipmentItem, NewCommand(), Env, WornArmor, gear, EquipmentChoice, applyEquipment()

### Community 34 - "5e Armor and Attack Explain"
Cohesion: 0.18
Nodes (6): Part, TestDualWield(), damageLabel(), Weapon, attackTargets(), weaponReach()

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.14
Nodes (20): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, cellAt(), classTable(), featuresUnder() (+12 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "DB Open and Migrate"
Cohesion: 0.11
Nodes (24): newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open(), openPool() (+16 more)

### Community 41 - "select_templ.go"
Cohesion: 0.07
Nodes (38): Align, ContentProps, ctxKey, ctxState, GroupProps, ItemProps, LabelProps, Props (+30 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "Spreadsheet Image Import"
Cohesion: 0.10
Nodes (25): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+17 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.15
Nodes (10): ClassLevel, TestSlotMaxima(), applyASI(), ClassInfo, hpGain(), raise(), LevelUpChoice, Spell (+2 more)

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
Cohesion: 0.18
Nodes (16): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), armorCondition(), columnInts() (+8 more)

### Community 52 - "Release Workflow"
Cohesion: 0.11
Nodes (18): CGO_ENABLED, GH_TOKEN, jobs, release, name, permissions, contents, tags (+10 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.20
Nodes (13): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+5 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "parseCommand"
Cohesion: 0.18
Nodes (11): findMod(), newMod(), parseMod(), SaveFeature(), IsModTarget(), featureCommand(), weaponCommand(), coinCommand() (+3 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Campaign"
Cohesion: 0.22
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

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
Cohesion: 0.22
Nodes (7): findItem(), parseArmor(), parseWeapon(), SaveWeapon(), TestGearCommands(), ensureItem(), Item

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

### Community 72 - "Session"
Cohesion: 0.07
Nodes (20): constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), serve(), TestCredentialsMatch(), TestIsLoopback(), TestMiddlewareIdentity() (+12 more)

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

### Community 80 - "testing.T"
Cohesion: 0.07
Nodes (59): TestCanCharacter(), TestCanManageBlock(), TestCanReadAttachment(), TestCanReadImage(), TestCanSeeCombatantHP(), TestVisibleEntity(), TestVisibleLooks(), nullID() (+51 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "builder_templ.go"
Cohesion: 0.07
Nodes (53): UsesLabel(), FeatureEffects, Coins, CustomMod, Feature, Uses, ItemData, builderOp (+45 more)

### Community 99 - "Entry"
Cohesion: 0.26
Nodes (9): ArmorData, Entry, ItemData, SpellData, normalizer, armorText(), children(), joinText() (+1 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags(), ftEntry (+7 more)

### Community 102 - "Import"
Cohesion: 0.14
Nodes (22): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+14 more)

### Community 103 - "Entry"
Cohesion: 0.12
Nodes (17): dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), Site, Source, FetchPage() (+9 more)

### Community 105 - "App"
Cohesion: 0.06
Nodes (25): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Campaign, Entity, PendingKind, App (+17 more)

### Community 108 - "fromLookup"
Cohesion: 0.18
Nodes (11): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc (+3 more)

### Community 110 - "open5e.go"
Cohesion: 0.08
Nodes (30): TestUnflattenTable(), fixture, fixtures, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass (+22 more)

### Community 120 - "lanHostname"
Cohesion: 0.30
Nodes (5): attachmentRefs(), TestAttachmentRefs(), lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 319 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **22 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `context.Context`, `database/sql.NullInt64`, `Websocket Hub`, `Session`, `.loadLive`, `fromLookup`, `.renderCompendium`, `.commandFailed`, `Viewer`, `testing.T`, `Character`, `.blockFor`, `TLS Cert and Main`, `Campaign`, `.fail`, `.encounterView`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.07926658840628319 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `NewCommand`, `5e Armor and Attack Explain`, `github.com/a-h/templ.Component`, `Character Sheet View`, `.commandFailed`, `5e Level-Up and Slots`, `5e Command Parser`, `builder_templ.go`, `Character`, `wizardChoice`, `parseCommand`, `.encounterView`, `findItem`?**
  _High betweenness centrality (0.026) - this node is a cross-community bridge._
- **Should `Vendored shadcn-templ Bundle` be split into smaller, more focused modules?**
  _Cohesion score 0.07461038961038961 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Websocket Hub` to `Auth and Compendium Core`, `Viewer`, `App`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011915140947399011 - nodes in this community are weakly interconnected._