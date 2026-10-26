# Graph Report - gm-session-manager  (2026-10-26)

## Corpus Check
- 231 files · ~443,216 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 49 file(s) not represented in the graph (top: .templ 41, .css 4, (none) 3)

## Summary
- 3468 nodes · 13973 edges · 110 communities (91 shown, 19 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 1523 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e219481f`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- shadcn-templ-3310c53c2176f836.js
- context.Context
- CompendiumView
- github.com/a-h/templ.Attributes
- database/sql.NullInt64
- Character Sheet View
- github.com/a-h/templ.Component
- Hub
- .do
- combobox_templ.go
- createProvider
- 5e Command Apply
- .Entry
- IfElse
- Combobox JS
- perm.go
- Select JS
- Floating UI DOM
- Hover Interaction JS
- who
- popover_templ.go
- Notes Editor JS
- combat_templ.go
- Anchor Positioning JS
- CN
- Project Rules (CLAUDE.md, README)
- .loadLive
- 5e Sheet Model
- monsterBody
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- Weapon
- Session
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
- deviceOS
- Queries
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
- Encounter
- Markdown Export
- List Navigation JS
- vendorjs Tool
- Item
- wikiText
- sqlc Config
- Transition Status JS
- Checkbox JS
- Claude Settings
- Switch JS
- Tabs JS
- Sed Hook
- SearchReadableBlocksRow
- .blockFor
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
- EquipmentOption
- Import
- Site
- toast_templ.go
- App
- Slugify
- fromLookup
- Entry
- lanHostname

## God Nodes (most connected - your core abstractions)
1. `App` - 235 edges
2. `CN()` - 107 edges
3. `Command` - 100 edges
4. `Button()` - 95 edges
5. `Sheet` - 93 edges
6. `j()` - 83 edges
7. `Entry` - 66 edges
8. `newTestEnv()` - 60 edges
9. `Hub` - 58 edges
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

## Communities (110 total, 19 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.08
Nodes (68): TestMain(), contextKey, Identity, Manager, keepAwake(), looseString, o5eAssignment, o5eChild (+60 more)

### Community 1 - "shadcn-templ-3310c53c2176f836.js"
Cohesion: 0.06
Nodes (163): A(), Ae(), b(), be(), bt(), c(), Ce(), I() (+155 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (63): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListMonstersByNameRow, ListImageEntitiesRow, Note (+55 more)

### Community 3 - "CompendiumView"
Cohesion: 0.13
Nodes (24): CountHabitatsRow, SearchCompendiumRow, PickerTarget, alignmentColor(), CompendiumPage(), CompendiumResults(), compendiumRow(), CompendiumRows() (+16 more)

### Community 4 - "github.com/a-h/templ.Attributes"
Cohesion: 0.04
Nodes (72): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+64 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (39): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Attachment, CompendiumSource, Condition, BlockStep (+31 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.07
Nodes (74): ArmorEffect(), Signed(), Field, ContentProps, ctxKey, ctxState, Props, Side (+66 more)

### Community 7 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (130): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant (+122 more)

### Community 8 - "Hub"
Cohesion: 0.06
Nodes (33): ClientInfo, event, fakeHandler, Handler, Client, Hub, Topic, New() (+25 more)

### Community 9 - ".do"
Cohesion: 0.16
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 10 - "combobox_templ.go"
Cohesion: 0.05
Nodes (55): Props, Size, Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey (+47 more)

### Community 11 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.08
Nodes (49): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+41 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (30): addEntry(), setStat(), MonsterData, NameDesc, ValidKind(), HomebrewOwner, HomebrewInput, HomebrewInputFor() (+22 more)

### Community 14 - "IfElse"
Cohesion: 0.16
Nodes (20): Look, Props, Size, initialChecked(), Switch(), IfElse(), characterLooks(), hpQuickForm() (+12 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "perm.go"
Cohesion: 0.15
Nodes (31): blockCase, Combatant, VisibleEntity(), VisibleLooks(), ImageLook, BlockVisibilities(), CanDeleteCharacter(), CanEditBlock() (+23 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "who"
Cohesion: 0.06
Nodes (22): DecodeData(), TextData, Character, CompendiumEntry, ClassFeature, ParseSheet(), CanCharacter(), CreateHomebrewParams (+14 more)

### Community 21 - "popover_templ.go"
Cohesion: 0.14
Nodes (20): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+12 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "combat_templ.go"
Cohesion: 0.11
Nodes (34): Props, Variant, combatantGroup, CombatantView, Badge(), badgeClasses(), addCombatantForm(), addCombatants() (+26 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "CN"
Cohesion: 0.09
Nodes (52): ctxKey, ctxState, DialogProps, ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps (+44 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - ".loadLive"
Cohesion: 0.07
Nodes (20): Encounter, Fields, ParseFields(), CreateEncounterParams, ListCharacterEncountersParams, CombatView, saveFieldLinks(), notFoundIsNil() (+12 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (14): DeathSaves, TestModifier(), Sum(), TestUpgradeSheet(), Coins, Sheet, Slot, IsConditionSlug() (+6 more)

### Community 29 - "monsterBody"
Cohesion: 0.18
Nodes (18): SpellData, classBadges(), classBody(), commaPrefixed(), entryBody(), EntryDetails(), itemBody(), levelsText() (+10 more)

### Community 30 - "Scroll Lock JS"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 31 - "Popover JS"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 32 - "Tabbable Focus JS"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 33 - "Weapon"
Cohesion: 0.19
Nodes (6): Part, skillByKey(), damageLabel(), Weapon, Skill, attackTargets()

### Community 34 - "Session"
Cohesion: 0.06
Nodes (25): constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve(), TestCredentialsMatch(), TestIsLoopback() (+17 more)

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.12
Nodes (24): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, withCoreTraits(), pageSource(), PageSourceKey() (+16 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "DB Open and Migrate"
Cohesion: 0.13
Nodes (21): apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock(), TestFTS5Available() (+13 more)

### Community 41 - "select_templ.go"
Cohesion: 0.06
Nodes (58): LevelUpPlan, Align, ContentProps, ctxKey, ctxState, GroupProps, ItemData, ItemProps (+50 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "cell"
Cohesion: 0.09
Nodes (29): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+21 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 46 - "Queries"
Cohesion: 0.13
Nodes (13): Queries, Combatant, Character, applyCondOps(), uniqueName(), condOwner, linkEntityLook(), formInt() (+5 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.13
Nodes (13): ClassLevel, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), ClassInfo (+5 more)

### Community 48 - "5e Command Parser"
Cohesion: 0.12
Nodes (19): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+11 more)

### Community 49 - "tabs_templ.go"
Cohesion: 0.20
Nodes (22): ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant (+14 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.26
Nodes (12): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), abilityKeys(), addProfs() (+4 more)

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

### Community 56 - "CustomMod"
Cohesion: 0.15
Nodes (13): findMod(), newMod(), parseMod(), SaveFeature(), UsesLabel(), FeatureEffects, CustomMod, Feature (+5 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Encounter"
Cohesion: 0.23
Nodes (10): ListCombatantsRow, initiativeMod(), outOfFight(), pcSheet(), stepTurn(), combatantKind(), combatantLink(), playerCombatantLink() (+2 more)

### Community 59 - "Markdown Export"
Cohesion: 0.28
Nodes (9): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+1 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

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

### Community 72 - ".blockFor"
Cohesion: 0.24
Nodes (3): GetBlockRow, blockAccessLevel, stepsPayload

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
Cohesion: 0.07
Nodes (54): Import(), TestCanCharacter(), nullID(), TestBuilderCompendium(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName(), TestLanguagesForm() (+46 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "builder_templ.go"
Cohesion: 0.07
Nodes (66): builderOp, builderVerb, coin, profFocus, sourceField, Combobox(), Content(), Item() (+58 more)

### Community 99 - "Entry"
Cohesion: 0.14
Nodes (16): TestUnflattenTable(), fixture, fixtures, normalizer, actionRank(), byPK(), children(), decode() (+8 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags() (+7 more)

### Community 101 - "EquipmentOption"
Cohesion: 0.20
Nodes (8): EquipmentOption, EquipmentItem, Env, Spell, WornArmor, gear, EquipmentChoice, applyEquipment()

### Community 102 - "Import"
Cohesion: 0.19
Nodes (15): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+7 more)

### Community 103 - "Site"
Cohesion: 0.16
Nodes (10): dataFor(), Kind, fiveToolsPage(), Site, FetchPage(), Page, ParsePageURL(), SiteByKey() (+2 more)

### Community 104 - "toast_templ.go"
Cohesion: 0.33
Nodes (8): Props, ToasterProps, Type, Toast(), Toaster(), toastIcon(), toastListItem(), Toast()

### Community 105 - "App"
Cohesion: 0.06
Nodes (29): ClearCookie(), SetCookie(), writeCookie(), FromContext(), ftsQuery(), Search(), Campaign, Entity (+21 more)

### Community 107 - "Slugify"
Cohesion: 0.07
Nodes (12): Slugify(), SavePage(), ListAttachmentFilesRow, Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile() (+4 more)

### Community 108 - "fromLookup"
Cohesion: 0.29
Nodes (8): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc

### Community 110 - "Entry"
Cohesion: 0.31
Nodes (10): ArmorData, Entry, ItemData, storeSource(), upsert(), o5eArmor, armorData(), armorText() (+2 more)

### Community 120 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 331 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **19 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `context.Context`, `Session`, `database/sql.NullInt64`, `Hub`, `.blockFor`, `.do`, `Slugify`, `Queries`, `perm.go`, `who`, `TLS Cert and Main`, `.loadLive`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.07578392621870883 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `Weapon`, `EquipmentOption`, `Character Sheet View`, `select_templ.go`, `Queries`, `5e Level-Up and Slots`, `5e Command Parser`, `builder_templ.go`, `who`, `wizardChoice`, `CustomMod`, `Encounter`, `Item`?**
  _High betweenness centrality (0.034) - this node is a cross-community bridge._
- **Should `shadcn-templ-3310c53c2176f836.js` be split into smaller, more focused modules?**
  _Cohesion score 0.05754269873915765 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Hub` to `Auth and Compendium Core`, `perm.go`, `App`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011927144535840189 - nodes in this community are weakly interconnected._