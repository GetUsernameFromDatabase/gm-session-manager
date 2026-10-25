# Graph Report - gm-session-manager  (2026-10-26)

## Corpus Check
- 227 files · ~423,496 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 48 file(s) not represented in the graph (top: .templ 40, .css 4, (none) 3)

## Summary
- 3328 nodes · 13401 edges · 107 communities (87 shown, 20 thin omitted)
- Extraction: 89% EXTRACTED · 11% INFERRED · 0% AMBIGUOUS · INFERRED: 1463 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `36bea5cd`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- Vendored shadcn-templ Bundle
- context.Context
- .do
- shadcn Component Props
- database/sql.NullInt64
- Character Sheet View
- github.com/a-h/templ.Component
- Websocket Hub
- popover_templ.go
- select_templ.go
- who
- 5e Command Apply
- .Entry
- .loadLive
- Combobox JS
- perm.go
- Select JS
- Floating UI DOM
- Hover Interaction JS
- monsterBody
- Campaign
- Notes Editor JS
- CombatantView
- Anchor Positioning JS
- TestMiddlewareIdentity
- Project Rules (CLAUDE.md, README)
- parseCommand
- 5e Sheet Model
- ImportView
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- .entityLook
- 5e Armor and Attack Explain
- App JS Command Bar
- Dismiss Interaction JS
- ParseWikidot
- ParseAidedd
- Command Palette JS
- DB Open and Migrate
- Dialog JS
- Spreadsheet Image Import
- Tooltip JS
- testEnv
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
- Campaign
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
- SearchReadableBlocksRow
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
- newStore
- Site
- App
- EquipmentOption
- .builderPicks
- generateSVG
- open5e.go
- lanHostname

## God Nodes (most connected - your core abstractions)
1. `App` - 230 edges
2. `CN()` - 105 edges
3. `Command` - 99 edges
4. `Sheet` - 93 edges
5. `Button()` - 85 edges
6. `Entry` - 66 edges
7. `j()` - 58 edges
8. `Hub` - 57 edges
9. `newTestEnv()` - 57 edges
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

## Communities (107 total, 20 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.09
Nodes (49): TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), Store, ClientInfo (+41 more)

### Community 1 - "Vendored shadcn-templ Bundle"
Cohesion: 0.07
Nodes (152): A(), Ae(), b(), be(), bt(), c(), Ce(), I() (+144 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (64): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumEntry, CompendiumSource (+56 more)

### Community 3 - ".do"
Cohesion: 0.19
Nodes (13): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+5 more)

### Community 4 - "shadcn Component Props"
Cohesion: 0.02
Nodes (194): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+186 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (54): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, New(), ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment (+46 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.06
Nodes (87): ArmorEffect(), Signed(), Field, ContentProps, ctxKey, ctxState, Props, Side (+79 more)

### Community 7 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (109): ListMonstersByNameRow, BlockView, CombatView, FeaturePreview, NoteListView, PendingView, WizardStep, WS (+101 more)

### Community 8 - "Websocket Hub"
Cohesion: 0.05
Nodes (33): IsLoopback(), TestIsLoopback(), event, fakeHandler, Handler, Client, Topic, New() (+25 more)

### Community 9 - "popover_templ.go"
Cohesion: 0.16
Nodes (16): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+8 more)

### Community 10 - "select_templ.go"
Cohesion: 0.11
Nodes (23): Align, ContentProps, ctxKey, ctxState, GroupProps, ItemProps, LabelProps, Props (+15 more)

### Community 11 - "who"
Cohesion: 0.06
Nodes (24): DecodeData(), TextData, Slugify(), Character, Combatant, ParseSheet(), SizeNamed(), CanCharacter() (+16 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.08
Nodes (49): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+41 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (33): addEntry(), setStat(), TestHomebrew(), CoreTraits(), TestCoreTraits(), MonsterData, NameDesc, ValidKind() (+25 more)

### Community 14 - ".loadLive"
Cohesion: 0.07
Nodes (20): Encounter, ListCombatantsRow, CreateEncounterParams, ListCharacterEncountersParams, initiativeMod(), outOfFight(), pcSheet(), skippedNotice() (+12 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "perm.go"
Cohesion: 0.16
Nodes (27): blockCase, Combatant, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), ImageLook, BlockVisibilities() (+19 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "monsterBody"
Cohesion: 0.18
Nodes (18): SpellData, classBadges(), classBody(), commaPrefixed(), entryBody(), EntryDetails(), itemBody(), levelsText() (+10 more)

### Community 21 - "Campaign"
Cohesion: 0.22
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.12
Nodes (30): Props, Variant, combatantGroup, CombatantView, Badge(), badgeClasses(), beingBadge(), combatantChip() (+22 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "TestMiddlewareIdentity"
Cohesion: 0.11
Nodes (14): constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity(), TestOnlyHashIsStored() (+6 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "parseCommand"
Cohesion: 0.33
Nodes (6): featureCommand(), weaponCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 28 - "5e Sheet Model"
Cohesion: 0.09
Nodes (12): DeathSaves, TestExplainSums(), skillByKey(), Sum(), Character, Coins, Sheet, Slot (+4 more)

### Community 29 - "ImportView"
Cohesion: 0.07
Nodes (40): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, ListBacklinksParams, Props, Size, ImagesView, ImageView (+32 more)

### Community 30 - "Scroll Lock JS"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 31 - "Popover JS"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 32 - "Tabbable Focus JS"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 33 - ".entityLook"
Cohesion: 0.22
Nodes (4): SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams, lookOps

### Community 34 - "5e Armor and Attack Explain"
Cohesion: 0.17
Nodes (7): TestModifier(), Part, damageLabel(), Weapon, Modifier(), attackTargets(), weaponReach()

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.13
Nodes (22): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+14 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "DB Open and Migrate"
Cohesion: 0.14
Nodes (19): apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestFTS5Available(), TestLoadMigrationsRejectsBadNames() (+11 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "Spreadsheet Image Import"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 46 - "testEnv"
Cohesion: 0.17
Nodes (8): Role, assetETags(), immutableAsset(), useValidateTag(), testEnv, TestAdminLogin(), TestCampaignJoinLinks(), TestLogout()

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.14
Nodes (12): ClassLevel, TestSlotMaxima(), applyASI(), ClassInfo, LevelUpPlan, hpGain(), raise(), LevelUpChoice (+4 more)

### Community 48 - "5e Command Parser"
Cohesion: 0.13
Nodes (18): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+10 more)

### Community 49 - "Tabs Component"
Cohesion: 0.08
Nodes (51): CountHabitatsRow, SearchCompendiumRow, Look, ContentProps, ctxKey, ctxState, ListProps, Orientation (+43 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "Equipment and Skill Choices"
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

### Community 56 - "5e Feature Effects"
Cohesion: 0.18
Nodes (11): findMod(), newMod(), parseMod(), SaveFeature(), UsesLabel(), FeatureEffects, CustomMod, Feature (+3 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Campaign"
Cohesion: 0.09
Nodes (12): ClearCookie(), SetCookie(), writeCookie(), RequireAdmin(), Campaign, PendingKind, CreateCampaignParams, deviceOS() (+4 more)

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
Cohesion: 0.19
Nodes (6): findItem(), parseArmor(), TestGearCommands(), TestUpgradeSheet(), ensureItem(), Item

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

### Community 72 - "Config Loading"
Cohesion: 0.21
Nodes (10): Config, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+2 more)

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
Cohesion: 0.06
Nodes (59): Import(), TestCanCharacter(), TestCanManageBlock(), TestCanReadAttachment(), TestCanReadImage(), TestCanSeeCombatantHP(), TestVisibleEntity(), TestVisibleLooks() (+51 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "builder_templ.go"
Cohesion: 0.07
Nodes (67): ItemData, builderOp, builderVerb, coin, profFocus, sourceField, Close(), CloseFor() (+59 more)

### Community 99 - "Entry"
Cohesion: 0.23
Nodes (11): Entry, fixtures, storeSource(), upsert(), normalizer, byPK(), children(), joinText() (+3 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.21
Nodes (16): abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags() (+8 more)

### Community 102 - "newStore"
Cohesion: 0.24
Nodes (13): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+5 more)

### Community 103 - "Site"
Cohesion: 0.14
Nodes (13): dataFor(), fiveToolsPage(), Kind, Site, FetchPage(), Page, pageSource(), PageSourceKey() (+5 more)

### Community 105 - "App"
Cohesion: 0.06
Nodes (16): FromContext(), App, blockAccessForm, blockAccessLevel, compendiumBack(), viewerCampaign(), formIntQuery(), isHTMX() (+8 more)

### Community 106 - "EquipmentOption"
Cohesion: 0.25
Nodes (8): EquipmentOption, EquipmentItem, Env, Spell, WornArmor, gear, EquipmentChoice, applyEquipment()

### Community 109 - "generateSVG"
Cohesion: 0.60
Nodes (5): Props, attrString(), generateSVG(), getIconContent(), Icon()

### Community 110 - "open5e.go"
Cohesion: 0.07
Nodes (30): ArmorData, TestUnflattenTable(), ItemData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild (+22 more)

### Community 119 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 318 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **20 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `.entityLook`, `context.Context`, `database/sql.NullInt64`, `Websocket Hub`, `who`, `.builderPicks`, `.loadLive`, `testEnv`, `perm.go`, `TLS Cert and Main`, `Campaign`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.09137289753464337 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `5e Armor and Attack Explain`, `Character Sheet View`, `Config Loading`, `EquipmentOption`, `who`, `.loadLive`, `5e Level-Up and Slots`, `5e Command Parser`, `builder_templ.go`, `Equipment and Skill Choices`, `5e Feature Effects`, `5e Gear and Dual Wield`?**
  _High betweenness centrality (0.026) - this node is a cross-community bridge._
- **Should `Vendored shadcn-templ Bundle` be split into smaller, more focused modules?**
  _Cohesion score 0.07461038961038961 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Auth and Compendium Core` to `Websocket Hub`, `perm.go`, `App`?**
  _High betweenness centrality (0.024) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.012805222522126672 - nodes in this community are weakly interconnected._