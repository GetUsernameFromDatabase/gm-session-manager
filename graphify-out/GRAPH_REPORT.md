# Graph Report - gm-session-manager  (2026-10-25)

## Corpus Check
- 215 files · ~407,212 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 47 file(s) not represented in the graph (top: .templ 39, .css 4, (none) 3)

## Summary
- 3217 nodes · 12662 edges · 105 communities (87 shown, 18 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 1409 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `d3b9c44c`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- Vendored shadcn-templ Bundle
- context.Context
- App
- shadcn Component Props
- database/sql.NullInt64
- Character Sheet View
- .loadLive
- Websocket Hub
- select_templ.go
- ImportView
- Character
- 5e Command Apply
- .Entry
- Attachment
- Combobox JS
- Visibility Rules (perm)
- Select JS
- Floating UI DOM
- Hover Interaction JS
- .do
- github.com/a-h/templ.Component
- Notes Editor JS
- combatantRow
- Anchor Positioning JS
- Admin Credentials and Loopback
- Project Rules (CLAUDE.md, README)
- who
- 5e Sheet Model
- parseCommand
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- combobox_templ.go
- 5e Armor and Attack Explain
- App JS Command Bar
- Dismiss Interaction JS
- ArmorData
- ParseAidedd
- Command Palette JS
- DB Open and Migrate
- levelUpForm
- Dialog JS
- Spreadsheet Image Import
- Tooltip JS
- ParseWikidot
- newStore
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
- NewCommand
- Config Loading
- Portal JS
- Click Interaction JS
- Focus Interaction JS
- newLiveFixture
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
- Entry
- buildReview
- testEnv
- lanHostname

## God Nodes (most connected - your core abstractions)
1. `App` - 219 edges
2. `CN()` - 104 edges
3. `Command` - 98 edges
4. `Sheet` - 91 edges
5. `Button()` - 80 edges
6. `Entry` - 65 edges
7. `j()` - 58 edges
8. `Hub` - 57 edges
9. `d()` - 50 edges
10. `Manager` - 49 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `NewCommand()` --calls--> `apply()`  [INFERRED]
  dnd5e/command.go → db/db.go
- `slotLabel()` --references--> `Slot`  [EXTRACTED]
  web/views/live.go → dnd5e/sheet.go
- `TestCommandPermissions()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go
- `TestServerSideStripping()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (105 total, 18 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.08
Nodes (69): TestMain(), contextKey, Identity, Manager, looseString, o5eAssignment, o5eChild, o5eClass (+61 more)

### Community 1 - "Vendored shadcn-templ Bundle"
Cohesion: 0.07
Nodes (152): A(), Ae(), b(), be(), bt(), c(), Ce(), I() (+144 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (63): fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListMonstersByNameRow, ListImageEntitiesRow, Encounter, AddBlockAudienceParams (+55 more)

### Community 3 - "App"
Cohesion: 0.06
Nodes (20): ClearCookie(), SetCookie(), writeCookie(), FromContext(), PendingKind, App, blockAccessForm, coinCommand() (+12 more)

### Community 4 - "shadcn Component Props"
Cohesion: 0.02
Nodes (137): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+129 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (57): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, New(), GetImageRow, ListImagesRow, CompendiumEntry, CompendiumSource (+49 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.08
Nodes (70): ArmorEffect(), Signed(), Field, ctxState, actionEntry, edit, subtitlePart, Action() (+62 more)

### Community 7 - ".loadLive"
Cohesion: 0.05
Nodes (98): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Look, BlockView, CombatView, ImagesView, ImageView (+90 more)

### Community 8 - "Websocket Hub"
Cohesion: 0.06
Nodes (31): event, fakeHandler, Handler, Client, Topic, New(), NewClient(), randomID() (+23 more)

### Community 9 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Align, ContentProps, ctxKey, ctxState, GroupProps, ItemProps, LabelProps, Props (+16 more)

### Community 10 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows(), importStatusItems() (+2 more)

### Community 11 - "Character"
Cohesion: 0.07
Nodes (19): DecodeData(), Character, Combatant, ClassFeature, ParseSheet(), Campaign(), CharacterFileName(), fileSlug() (+11 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.09
Nodes (47): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+39 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (30): addEntry(), setStat(), MonsterData, NameDesc, ValidKind(), HomebrewOwner, HomebrewInput, HomebrewInputFor() (+22 more)

### Community 14 - "Attachment"
Cohesion: 0.19
Nodes (3): Attachment, CreateAttachmentParams, randomFileName()

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "Visibility Rules (perm)"
Cohesion: 0.06
Nodes (47): GetBlockRow, SearchReadableBlocksRow, blockCase, Combatant, Fields, ParseFields(), VisibleEntity(), VisibleLooks() (+39 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - ".do"
Cohesion: 0.19
Nodes (13): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+5 more)

### Community 21 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (64): Props, Variant, CountHabitatsRow, SearchCompendiumRow, PickerTarget, Badge(), badgeClasses(), Checkbox() (+56 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.08
Nodes (24): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+16 more)

### Community 23 - "combatantRow"
Cohesion: 0.24
Nodes (12): CombatantView, addCombatantForm(), addCombatants(), combatantRow(), Encounter(), EncounterView, hpDialog(), hpText() (+4 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "Admin Credentials and Loopback"
Cohesion: 0.08
Nodes (19): constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve(), TestCredentialsMatch(), TestIsLoopback() (+11 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "who"
Cohesion: 0.09
Nodes (9): ftsQuery(), Search(), Campaign, CreateCampaignParams, compendiumBack(), viewerCampaign(), Roster, HomebrewView (+1 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (12): DeathSaves, TestExplainSums(), skillByKey(), Sum(), TestUpgradeSheet(), Character, Sheet, Slot (+4 more)

### Community 29 - "parseCommand"
Cohesion: 0.25
Nodes (9): findMod(), newMod(), parseMod(), SaveWeapon(), IsModTarget(), featureCommand(), weaponCommand(), hpCommand() (+1 more)

### Community 30 - "Scroll Lock JS"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 31 - "Popover JS"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 32 - "Tabbable Focus JS"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 33 - "combobox_templ.go"
Cohesion: 0.05
Nodes (56): Props, Size, Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey (+48 more)

### Community 34 - "5e Armor and Attack Explain"
Cohesion: 0.18
Nodes (6): TestModifier(), Part, damageLabel(), Weapon, Modifier(), attackTargets()

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, ItemData, o5eArmor, armorData(), armorText()

### Community 38 - "ParseAidedd"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "DB Open and Migrate"
Cohesion: 0.13
Nodes (20): apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestFTS5Available(), TestLoadMigrationsRejectsBadNames() (+12 more)

### Community 41 - "levelUpForm"
Cohesion: 0.14
Nodes (26): LevelUpPlan, FeaturePreview, Option, WizardStep, Footer(), Label(), abilitySelect(), LevelUpView (+18 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "Spreadsheet Image Import"
Cohesion: 0.10
Nodes (21): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+13 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - "ParseWikidot"
Cohesion: 0.15
Nodes (20): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+12 more)

### Community 46 - "newStore"
Cohesion: 0.31
Nodes (11): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+3 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.15
Nodes (10): ClassLevel, TestSlotMaxima(), applyASI(), ClassInfo, hpGain(), raise(), LevelUpChoice, Spell (+2 more)

### Community 48 - "5e Command Parser"
Cohesion: 0.13
Nodes (18): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+10 more)

### Community 49 - "Tabs Component"
Cohesion: 0.20
Nodes (23): ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant, sourceField (+15 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "Equipment and Skill Choices"
Cohesion: 0.11
Nodes (27): CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestCoreTraits(), TestEquipmentOptions() (+19 more)

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

### Community 56 - "5e Feature Effects"
Cohesion: 0.26
Nodes (6): SaveFeature(), UsesLabel(), FeatureEffects, CustomMod, Feature, Uses

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
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "5e Gear and Dual Wield"
Cohesion: 0.26
Nodes (5): findItem(), parseArmor(), TestGearCommands(), ensureItem(), Item

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

### Community 71 - "NewCommand"
Cohesion: 0.29
Nodes (4): NewCommand(), Env, WornArmor, gear

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

### Community 76 - "newLiveFixture"
Cohesion: 0.17
Nodes (12): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), liveFixture, LiveKey(), dial() (+4 more)

### Community 77 - "Mark Others (aria-hidden)"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 78 - "Typeahead JS"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 80 - "newTestEnv"
Cohesion: 0.09
Nodes (43): TestRealSource(), Import(), seedGarbage(), TestCanCharacter(), TestVisibleEntity(), TestVisibleLooks(), nullID(), TestBuilderCompendium() (+35 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "builder_templ.go"
Cohesion: 0.07
Nodes (66): Coins, ItemData, builderOp, builderVerb, coin, profFocus, Combobox(), Content() (+58 more)

### Community 99 - "Entry"
Cohesion: 0.14
Nodes (16): TestUnflattenTable(), fixture, fixtures, normalizer, actionRank(), byPK(), children(), decode() (+8 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags() (+7 more)

### Community 103 - "Entry"
Cohesion: 0.11
Nodes (24): Entry, dataFor(), SpellData, TextData, Slugify(), fiveToolsPage(), importSource(), readSource() (+16 more)

### Community 105 - "buildReview"
Cohesion: 0.36
Nodes (4): csvImport, csvImports, buildReview(), formValues()

### Community 106 - "testEnv"
Cohesion: 0.17
Nodes (8): Role, assetETags(), immutableAsset(), useValidateTag(), testEnv, TestAdminLogin(), TestCampaignJoinLinks(), TestLogout()

### Community 119 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 308 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **18 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `context.Context`, `database/sql.NullInt64`, `Entry`, `Websocket Hub`, `buildReview`, `testEnv`, `Character`, `Attachment`, `Visibility Rules (perm)`, `Equipment and Skill Choices`, `TLS Cert and Main`, `who`?**
  _High betweenness centrality (0.031) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.07548770144189991 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `5e Armor and Attack Explain`, `Character Sheet View`, `NewCommand`, `levelUpForm`, `Character`, `5e Level-Up and Slots`, `5e Command Parser`, `builder_templ.go`, `Equipment and Skill Choices`, `5e Feature Effects`, `parseCommand`, `5e Gear and Dual Wield`?**
  _High betweenness centrality (0.030) - this node is a cross-community bridge._
- **Should `Vendored shadcn-templ Bundle` be split into smaller, more focused modules?**
  _Cohesion score 0.07461038961038961 - nodes in this community are weakly interconnected._
- **Why does `Command` connect `5e Command Apply` to `Auth and Compendium Core`, `NewCommand`, `Character`, `5e Command Parser`, `vendorjs Tool`, `Equipment and Skill Choices`, `builder_templ.go`, `5e Feature Effects`, `parseCommand`, `5e Gear and Dual Wield`?**
  _High betweenness centrality (0.030) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01202112804322749 - nodes in this community are weakly interconnected._