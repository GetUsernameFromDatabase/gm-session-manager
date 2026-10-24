# Graph Report - gm-session-manager  (2026-10-24)

## Corpus Check
- 210 files · ~399,373 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 46 file(s) not represented in the graph (top: .templ 39, .css 4, (none) 2)

## Summary
- 3166 nodes · 12372 edges · 99 communities (81 shown, 18 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 1367 edges (avg confidence: 0.85)
- Token cost: 100,265 input · 6,500 output

## Community Hubs (Navigation)
- Auth and Compendium Core
- Vendored shadcn-templ Bundle
- sqlc Row Models
- HTTP Handlers and Cookies
- shadcn Component Props
- sqlc Queries and Attachments
- Character Sheet View
- Page Templates and Layout
- Websocket Hub
- Character Builder View
- Command and Table Components
- Character Wizard Handlers
- 5e Command Apply
- Open5e Importer Types
- Image Import and Backlinks
- Combobox JS
- Visibility Rules (perm)
- Select JS
- Floating UI DOM
- Hover Interaction JS
- Permission and Image Tests
- Badges and Compendium Search View
- Notes Editor JS
- Combat and Notes Live Views
- Anchor Positioning JS
- Admin Credentials and Loopback
- Project Rules (CLAUDE.md, README)
- Compendium Search and Campaigns
- 5e Sheet Model
- Entity and Character Images
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- Button and Input Group
- 5e Armor and Attack Explain
- App JS Command Bar
- Dismiss Interaction JS
- Web Flow Tests
- Wikidot Class Parser
- Command Palette JS
- DB Open and Migrate
- Wizard and Level-Up Views
- Dialog JS
- Spreadsheet Image Import
- Tooltip JS
- Wikidot Core Traits
- Compendium Import Tests
- 5e Level-Up and Slots
- 5e Command Parser
- Tabs Component
- Composite List Navigation
- Equipment and Skill Choices
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- 5e Feature Effects
- golangci-lint Config
- Popover Component
- Markdown Export
- List Navigation JS
- vendorjs Tool
- 5e Gear and Dual Wield
- Wiki Links
- sqlc Config
- Transition Status JS
- Checkbox JS
- Claude Settings
- Switch JS
- Tabs JS
- Sed Hook
- Equipment Options
- Config Loading
- Portal JS
- Click Interaction JS
- Focus Interaction JS
- Websocket Tests
- Mark Others (aria-hidden)
- Typeahead JS
- Live View Bindings JS
- Character Command Parsing
- Internal Backdrop JS
- Lifecycle JS
- LAN Hostname Helpers
- Lint and Release Rules
- Run Script
- d20 Favicon
- Validator Rule
- Go Module

## God Nodes (most connected - your core abstractions)
1. `App` - 219 edges
2. `CN()` - 104 edges
3. `Command` - 98 edges
4. `Sheet` - 90 edges
5. `Button()` - 80 edges
6. `Entry` - 62 edges
7. `j()` - 58 edges
8. `Hub` - 57 edges
9. `d()` - 50 edges
10. `Manager` - 49 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `NewCommand()` --calls--> `apply()`  [INFERRED]
  dnd5e/command.go → db/db.go
- `TestServerSideStripping()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go
- `Identity` --references--> `Session`  [EXTRACTED]
  auth/identity.go → db/sqlc/models.go
- `who` --references--> `Identity`  [EXTRACTED]
  web/helpers.go → auth/identity.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]

## Communities (99 total, 18 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.07
Nodes (72): TestMain(), contextKey, fixture, Identity, Manager, Entry, looseString, o5eAssignment (+64 more)

### Community 1 - "Vendored shadcn-templ Bundle"
Cohesion: 0.07
Nodes (152): A(), Ae(), b(), be(), bt(), c(), Ce(), I() (+144 more)

### Community 2 - "sqlc Row Models"
Cohesion: 0.01
Nodes (57): ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Encounter, Note, AddBlockAudienceParams, AddBlockLinkParams (+49 more)

### Community 3 - "HTTP Handlers and Cookies"
Cohesion: 0.06
Nodes (20): ClearCookie(), SetCookie(), writeCookie(), FromContext(), PendingKind, App, blockAccessForm, formIntQuery() (+12 more)

### Community 4 - "shadcn Component Props"
Cohesion: 0.03
Nodes (113): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+105 more)

### Community 5 - "sqlc Queries and Attachments"
Cohesion: 0.03
Nodes (49): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, CompendiumSource (+41 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.07
Nodes (82): SpellData, Slot, Signed(), Field, ContentProps, ctxKey, ctxState, Props (+74 more)

### Community 7 - "Page Templates and Layout"
Cohesion: 0.06
Nodes (78): CombatView, PendingView, WS, Alert(), Description(), Kbd(), Label(), Scripts() (+70 more)

### Community 8 - "Websocket Hub"
Cohesion: 0.06
Nodes (34): ClientInfo, event, fakeHandler, Handler, Client, Hub, Topic, New() (+26 more)

### Community 9 - "Character Builder View"
Cohesion: 0.06
Nodes (68): ItemData, Props, builderOp, builderVerb, sourceField, Combobox(), Content(), Item() (+60 more)

### Community 10 - "Command and Table Components"
Cohesion: 0.05
Nodes (67): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+59 more)

### Community 11 - "Character Wizard Handlers"
Cohesion: 0.06
Nodes (23): dataFor(), DecodeData(), Slugify(), Kind, Character, ParseSheet(), Campaign(), CharacterFileName() (+15 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.08
Nodes (52): Amount, condCommand(), conditionSlug(), Damage(), findMod(), Command, CondOp, Effects (+44 more)

### Community 13 - "Open5e Importer Types"
Cohesion: 0.06
Nodes (42): ArmorData, TestUnflattenTable(), MonsterData, NameDesc, ValidKind(), fixture, fixtures, firstInt() (+34 more)

### Community 14 - "Image Import and Backlinks"
Cohesion: 0.06
Nodes (46): Props, ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, ctxState, Props, Size, ImagesView (+38 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "Visibility Rules (perm)"
Cohesion: 0.10
Nodes (35): GetBlockRow, blockCase, Combatant, Fields, Look, ParseFields(), VisibleEntity(), VisibleLooks() (+27 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "Permission and Image Tests"
Cohesion: 0.08
Nodes (46): TestRealSource(), Import(), TestCanCharacter(), TestCanManageBlock(), TestCanReadAttachment(), TestCanReadImage(), TestCanSeeCombatantHP(), TestVisibleEntity() (+38 more)

### Community 21 - "Badges and Compendium Search View"
Cohesion: 0.08
Nodes (38): Props, Variant, CountHabitatsRow, SearchCompendiumRow, Device, PickerTarget, RosterCharacter, Badge() (+30 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.08
Nodes (24): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+16 more)

### Community 23 - "Combat and Notes Live Views"
Cohesion: 0.10
Nodes (32): BlockView, CombatantView, NoteListView, notFoundIsNil(), addCombatantForm(), addCombatants(), combatantRow(), Encounter() (+24 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "Admin Credentials and Loopback"
Cohesion: 0.06
Nodes (20): constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), serve(), TestCredentialsMatch(), TestIsLoopback(), TestMiddlewareIdentity() (+12 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Compendium Search and Campaigns"
Cohesion: 0.11
Nodes (10): ftsQuery(), Search(), Campaign, CreateCampaignParams, compendiumBack(), viewerCampaign(), HomebrewView, LinkItem (+2 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (13): DeathSaves, skillByKey(), Sum(), TestUpgradeSheet(), Character, Coins, Sheet, IsConditionSlug() (+5 more)

### Community 29 - "Entity and Character Images"
Cohesion: 0.09
Nodes (13): Combatant, Entity, SetCharacterCurrentImageParams, SetEntityCurrentImageParams, ToggleCharacterImageRevealedParams, ToggleEntityImageRevealedParams, UnlinkCharacterImageParams, UnlinkEntityImageParams (+5 more)

### Community 30 - "Scroll Lock JS"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 31 - "Popover JS"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 32 - "Tabbable Focus JS"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 33 - "Button and Input Group"
Cohesion: 0.11
Nodes (25): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+17 more)

### Community 34 - "5e Armor and Attack Explain"
Cohesion: 0.14
Nodes (10): ItemData, TestModifier(), Part, TestDualWield(), damageLabel(), Weapon, Modifier(), compendiumWeapon() (+2 more)

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "Web Flow Tests"
Cohesion: 0.15
Nodes (17): main(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 38 - "Wikidot Class Parser"
Cohesion: 0.13
Nodes (23): block, ClassFeature, ClassData, attr(), byID(), cellAt(), classTable(), collect() (+15 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "DB Open and Migrate"
Cohesion: 0.12
Nodes (23): newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open(), openPool() (+15 more)

### Community 41 - "Wizard and Level-Up Views"
Cohesion: 0.15
Nodes (23): LevelUpPlan, FeaturePreview, Option, WizardStep, Footer(), abilitySelect(), LevelUpView, levelUpForm() (+15 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "Spreadsheet Image Import"
Cohesion: 0.10
Nodes (22): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+14 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - "Wikidot Core Traits"
Cohesion: 0.14
Nodes (21): CoreTraits(), TestCoreTraits(), storeSource(), upsert(), Source, FetchWikidot(), fetchWikidot(), ParseWikidot() (+13 more)

### Community 46 - "Compendium Import Tests"
Cohesion: 0.14
Nodes (18): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+10 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.15
Nodes (10): ClassLevel, TestSlotMaxima(), applyASI(), ClassInfo, hpGain(), raise(), LevelUpChoice, Spell (+2 more)

### Community 48 - "5e Command Parser"
Cohesion: 0.13
Nodes (18): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestHealingFromZeroResetsDeathSaves() (+10 more)

### Community 49 - "Tabs Component"
Cohesion: 0.22
Nodes (19): ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant (+11 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "Equipment and Skill Choices"
Cohesion: 0.15
Nodes (19): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), TextData, wikidotText() (+11 more)

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

### Community 56 - "5e Feature Effects"
Cohesion: 0.20
Nodes (10): parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), FeatureEffects, CustomMod, Feature, Uses (+2 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Popover Component"
Cohesion: 0.16
Nodes (16): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+8 more)

### Community 59 - "Markdown Export"
Cohesion: 0.28
Nodes (9): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+1 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.30
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 62 - "5e Gear and Dual Wield"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 63 - "Wiki Links"
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
Cohesion: 0.22
Nodes (8): hooks, PreToolUse, permissions, allow, ask, defaultMode, deny, $schema

### Community 68 - "Switch JS"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 69 - "Tabs JS"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 71 - "Equipment Options"
Cohesion: 0.25
Nodes (8): EquipmentOption, EquipmentItem, ArmorEffect(), TestExplainDamage(), WornArmor, gear, EquipmentChoice, applyEquipment()

### Community 72 - "Config Loading"
Cohesion: 0.29
Nodes (8): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc

### Community 73 - "Portal JS"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 74 - "Click Interaction JS"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 75 - "Focus Interaction JS"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 76 - "Websocket Tests"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 77 - "Mark Others (aria-hidden)"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 78 - "Typeahead JS"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 80 - "Character Command Parsing"
Cohesion: 0.50
Nodes (4): coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "LAN Hostname Helpers"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `graphify skill trigger`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Permissions Table (characters, monster HP, entities)` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 305 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **18 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Command` connect `5e Command Apply` to `Auth and Compendium Core`, `Character Wizard Handlers`, `5e Command Parser`, `Visibility Rules (perm)`, `Character Command Parsing`, `5e Feature Effects`, `5e Gear and Dual Wield`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `graphify skill trigger` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.0708895033990091 - nodes in this community are weakly interconnected._
- **Why does `App` connect `HTTP Handlers and Cookies` to `Auth and Compendium Core`, `sqlc Queries and Attachments`, `Web Flow Tests`, `Websocket Hub`, `Character Wizard Handlers`, `Compendium Import Tests`, `Visibility Rules (perm)`, `TLS Cert and Main`, `Combat and Notes Live Views`, `Admin Credentials and Loopback`, `Compendium Search and Campaigns`, `Entity and Character Images`?**
  _High betweenness centrality (0.031) - this node is a cross-community bridge._
- **Should `Vendored shadcn-templ Bundle` be split into smaller, more focused modules?**
  _Cohesion score 0.07461038961038961 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `5e Armor and Attack Explain`, `Character Sheet View`, `Equipment Options`, `Wizard and Level-Up Views`, `Character Wizard Handlers`, `5e Command Apply`, `5e Level-Up and Slots`, `5e Command Parser`, `Equipment and Skill Choices`, `5e Feature Effects`, `5e Gear and Dual Wield`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `sqlc Row Models` be split into smaller, more focused modules?**
  _Cohesion score 0.013435560423512231 - nodes in this community are weakly interconnected._