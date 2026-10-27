# Graph Report - gm-session-manager  (2026-10-27)

## Corpus Check
- 242 files · ~463,997 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3426 nodes · 13507 edges · 121 communities (101 shown, 20 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1329 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `3aacde2f`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- github.com/a-h/templ.Attributes
- context.Context
- IfElse
- table_templ.go
- database/sql.NullInt64
- Character Sheet View
- select_templ.go
- redirect.go
- .do
- CN
- createProvider
- 5e Command Apply
- .Entry
- Encounter
- Combobox JS
- perm.go
- Select JS
- Floating UI DOM
- Hover Interaction JS
- CompendiumView
- popover_templ.go
- Notes Editor JS
- CombatantView
- Anchor Positioning JS
- radiogroup.js
- Project Rules (CLAUDE.md, README)
- FromContext
- 5e Sheet Model
- github.com/a-h/templ.Component
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- Weapon
- wizardPick
- App JS Command Bar
- Dismiss Interaction JS
- ParseWikidot
- ParseAidedd
- Command Palette JS
- Open
- wikidotClass
- Dialog JS
- export/markdown.go
- Tooltip JS
- github.com/gin-gonic/gin.Context
- .encounterView
- 5e Level-Up and Slots
- Parse
- .Router
- Composite List Navigation
- wizardChoice
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- .loadLive
- golangci-lint Config
- ImportView
- .sendFeedback
- List Navigation JS
- vendorjs Tool
- Item
- parseCommand
- sqlc Config
- Transition Status JS
- Checkbox JS
- Claude Settings
- Switch JS
- Tabs JS
- Sed Hook
- release.sh
- builderFields
- Portal JS
- Click Interaction JS
- Focus Interaction JS
- Player
- Mark Others (aria-hidden)
- Typeahead JS
- Live View Bindings JS
- Import
- Internal Backdrop JS
- Lifecycle JS
- ImagePage
- Lint and Release Rules
- Run Script
- d20 Favicon
- Validator Rule
- Go Module
- CustomMod
- encoding/json.RawMessage
- newTestEnv
- errBodyWriter
- importSource
- deviceOS
- Character
- monsterBody
- Entry
- fromLookup
- open5e.go
- wikiText
- tooltip_templ.go
- radiogroup_templ.go
- Hub
- NewCommand
- build.sh
- switch_templ.go
- who
- SearchReadableBlocksRow
- capReader
- .entityLook

## God Nodes (most connected - your core abstractions)
1. `App` - 246 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Command` - 100 edges
5. `Sheet` - 93 edges
6. `newTestEnv()` - 74 edges
7. `reqLog` - 73 edges
8. `Entry` - 67 edges
9. `Hub` - 58 edges
10. `Manager` - 57 edges

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

## Communities (121 total, 20 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.08
Nodes (64): Props, TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), Entry (+56 more)

### Community 1 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (108): Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props (+100 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (61): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, ListAttachmentFilesRow, Note (+53 more)

### Community 3 - "IfElse"
Cohesion: 0.09
Nodes (45): Look, ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps (+37 more)

### Community 4 - "table_templ.go"
Cohesion: 0.09
Nodes (28): Props, ListMonstersByNameRow, Field, BodyProps, CaptionProps, CellProps, FooterProps, HeaderProps (+20 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (45): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, GetImageRow, ListImagesRow, Attachment, CompendiumSource, Condition (+37 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.08
Nodes (73): ArmorEffect(), Signed(), actionEntry, edit, subtitlePart, Action(), Card(), Content() (+65 more)

### Community 7 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Align, ContentProps, ctxKey, ctxState, GroupProps, ItemProps, LabelProps, Props (+16 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): connListener, peekedConn, newConnListener(), route(), SplitTLS()

### Community 9 - ".do"
Cohesion: 0.12
Nodes (20): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+12 more)

### Community 10 - "CN"
Cohesion: 0.11
Nodes (47): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+39 more)

### Community 11 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.08
Nodes (49): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+41 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (30): setStat(), TestHomebrew(), MonsterData, NameDesc, ValidKind(), HomebrewOwner, HomebrewInput, HomebrewInputFor() (+22 more)

### Community 14 - "Encounter"
Cohesion: 0.14
Nodes (5): Combatant, Encounter, Character, CreateEncounterParams, ListCharacterEncountersParams

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "perm.go"
Cohesion: 0.20
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

### Community 20 - "CompendiumView"
Cohesion: 0.11
Nodes (26): CountHabitatsRow, SearchCompendiumRow, PickerTarget, alignmentColor(), classBadges(), commaPrefixed(), CompendiumResults(), compendiumRow() (+18 more)

### Community 21 - "popover_templ.go"
Cohesion: 0.15
Nodes (21): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+13 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.08
Nodes (43): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Variant, combatantGroup (+35 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "FromContext"
Cohesion: 0.11
Nodes (11): ClearCookie(), writeCookie(), FromContext(), Toast(), browserName(), firstLine(), mdLinkText(), deviceKey() (+3 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (13): DeathSaves, TestModifier(), skillByKey(), Sum(), TestUpgradeSheet(), Coins, Sheet, Slot (+5 more)

### Community 29 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (104): CombatView, PendingView, WS, Button(), Input(), Content(), Description(), Empty() (+96 more)

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
Nodes (4): Part, damageLabel(), Weapon, attackTargets()

### Community 34 - "wizardPick"
Cohesion: 0.26
Nodes (15): FeaturePreview, WizardStep, abilityForm(), backgroundAbilities(), choiceName(), WizardChoice, WizardView, skillName() (+7 more)

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.31
Nodes (11): CoreTraits(), TestCoreTraits(), ParseWikidot(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSpecies(), TestParseWikidotSubclass() (+3 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "Open"
Cohesion: 0.11
Nodes (24): newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open(), openPool() (+16 more)

### Community 41 - "wikidotClass"
Cohesion: 0.18
Nodes (12): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), sourceLine() (+4 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.08
Nodes (25): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+17 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (18): SetCookie(), PendingKind, App, blockAccessForm, entityTopics(), formInt(), formIntQuery(), isHTMX() (+10 more)

### Community 46 - ".encounterView"
Cohesion: 0.27
Nodes (10): ListCombatantsRow, initiativeMod(), outOfFight(), pcSheet(), stepTurn(), combatantKind(), combatantLink(), playerCombatantLink() (+2 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.12
Nodes (15): ClassLevel, TestSlotMaxima(), applyASI(), ClassInfo, LevelUpPlan, hpGain(), raise(), LevelUpChoice (+7 more)

### Community 48 - "Parse"
Cohesion: 0.11
Nodes (20): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+12 more)

### Community 49 - ".Router"
Cohesion: 0.13
Nodes (10): TestRequireAdmin(), TestTokens(), RequireAdmin(), hashToken(), newToken(), tokens, assetETags(), immutableAsset() (+2 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.13
Nodes (21): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), TextData, wikidotText() (+13 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.15
Nodes (16): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+8 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - ".loadLive"
Cohesion: 0.11
Nodes (12): Entity, Fields, ParseFields(), saveFieldLinks(), notFoundIsNil(), parseLiveKey(), liveKey, toLook() (+4 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "ImportView"
Cohesion: 0.09
Nodes (33): ImageCategory, ImagesView, buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping() (+25 more)

### Community 59 - ".sendFeedback"
Cohesion: 0.33
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 63 - "parseCommand"
Cohesion: 0.40
Nodes (5): featureCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

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

### Community 72 - "builderFields"
Cohesion: 0.07
Nodes (68): ItemData, builderOp, builderVerb, coin, profFocus, Combobox(), Content(), Input() (+60 more)

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

### Community 80 - "Import"
Cohesion: 0.21
Nodes (13): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestSearch() (+5 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "ImagePage"
Cohesion: 0.08
Nodes (38): ListEntityFieldBacklinksRow, ListBacklinksRow, ListBacklinksParams, ListEntityFieldBacklinksParams, BlockView, ImageView, attrString(), generateSVG() (+30 more)

### Community 99 - "CustomMod"
Cohesion: 0.18
Nodes (11): findMod(), newMod(), parseMod(), SaveFeature(), UsesLabel(), FeatureEffects, CustomMod, Feature (+3 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.25
Nodes (14): ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags(), ftEntry (+6 more)

### Community 101 - "newTestEnv"
Cohesion: 0.05
Nodes (70): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestRealSource(), Import(), TestCanCharacter(), TestCanManageBlock(), TestCanReadAttachment() (+62 more)

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errBodyWriter, errorText(), TestErrorText()

### Community 103 - "importSource"
Cohesion: 0.43
Nodes (6): importSource(), readSource(), storeSource(), upsert(), parseOpen5e(), Source

### Community 104 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 105 - "Character"
Cohesion: 0.05
Nodes (29): dataFor(), DecodeData(), Kind, Slugify(), wikidotKind(), Character, CompendiumEntry, ParseSheet() (+21 more)

### Community 106 - "monsterBody"
Cohesion: 0.14
Nodes (25): SpellData, classBody(), crText(), entryBody(), EntryDetails(), fact(), itemBody(), levelsText() (+17 more)

### Community 107 - "Entry"
Cohesion: 0.15
Nodes (14): addEntry(), TestUnflattenTable(), fixture, normalizer, actionRank(), children(), decode(), firstOr() (+6 more)

### Community 108 - "fromLookup"
Cohesion: 0.21
Nodes (10): Config, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+2 more)

### Community 110 - "open5e.go"
Cohesion: 0.09
Nodes (26): ArmorData, ItemData, fixtures, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass (+18 more)

### Community 111 - "wikiText"
Cohesion: 0.27
Nodes (8): WikiLink, WikiPart, wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks()

### Community 113 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): ContentProps, ctxKey, ctxState, Props, Side, DisableTrigger(), state()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): ctxKey, ctxState, ItemProps, Props, initialValue(), state()

### Community 115 - "Hub"
Cohesion: 0.06
Nodes (35): IsLoopback(), TestIsLoopback(), ClientInfo, event, fakeHandler, Handler, Client, Hub (+27 more)

### Community 117 - "NewCommand"
Cohesion: 0.29
Nodes (7): EquipmentOption, EquipmentItem, Env, WornArmor, gear, EquipmentChoice, applyEquipment()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): Props, Size, initialChecked()

### Community 125 - "who"
Cohesion: 0.05
Nodes (23): fiveToolsPage(), ftsQuery(), Search(), Site, FetchPage(), Page, pageSource(), PageSourceKey() (+15 more)

### Community 137 - "SearchReadableBlocksRow"
Cohesion: 0.50
Nodes (3): SearchReadableBlocksRow, SearchReadableBlocksParams, NoteListView

### Community 149 - ".entityLook"
Cohesion: 0.22
Nodes (4): SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams, lookOps

## Knowledge Gaps
- **14 isolated node(s):** `CGO_ENABLED`, `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 328 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **20 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `github.com/gin-gonic/gin.Context` to `Auth and Compendium Core`, `context.Context`, `database/sql.NullInt64`, `Character`, `.do`, `Encounter`, `.encounterView`, `perm.go`, `.Router`, `Import`, `Hub`, `.entityLook`, `TLS Cert and Main`, `.loadLive`, `FromContext`, `who`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **What connects `CGO_ENABLED`, `gm-session-manager`, `testEnv` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.08266139657444005 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `Weapon`, `CustomMod`, `Character Sheet View`, `builderFields`, `Character`, `fromLookup`, `Encounter`, `5e Level-Up and Slots`, `Parse`, `.encounterView`, `wizardChoice`, `NewCommand`, `Item`?**
  _High betweenness centrality (0.030) - this node is a cross-community bridge._
- **Should `github.com/a-h/templ.Attributes` be split into smaller, more focused modules?**
  _Cohesion score 0.027503048367429887 - nodes in this community are weakly interconnected._
- **Why does `Command` connect `5e Command Apply` to `Auth and Compendium Core`, `CustomMod`, `builderFields`, `Character`, `Parse`, `perm.go`, `NewCommand`, `Item`, `parseCommand`?**
  _High betweenness centrality (0.020) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01267968112485092 - nodes in this community are weakly interconnected._