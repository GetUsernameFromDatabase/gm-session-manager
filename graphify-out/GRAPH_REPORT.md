# Graph Report - gm-session-manager  (2026-11-06)

## Corpus Check
- 285 files · ~539,572 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 3908 nodes · 16727 edges · 121 communities (102 shown, 19 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1676 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `1d799070`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- FromContext
- context.Context
- .do
- createProvider
- combat_templ.go
- SheetView
- combobox.js
- redirect.go
- Parse
- Feature
- select.js
- Command
- .Entry
- CombatantView
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ImportView
- wizardPick
- wikilink.js
- github.com/a-h/templ.Component
- popover.js
- ImagePage
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- testing.T
- app.js
- tabbable.js
- useDismiss
- DecodeData
- testing.T
- command.js
- dialog.js
- select_templ.go
- golang.org/x/net/html.Node
- tooltip.js
- IfElse
- NameDesc
- capReader
- export/markdown.go
- composite.js
- .levelUp
- NewClient
- ClassInfo
- Client
- Slugify
- ParseWiki
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- popover_templ.go
- golangci-lint Config
- CampaignRoster
- Book
- radiogroup.js
- vendorjs/main.go
- cell
- Pending
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- Input
- checkbox.js
- Campaign
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- ParseWikidot
- Lint and Release Rules
- Run Script
- github.com/a-h/templ.Component
- useTypeahead
- Hub
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Session
- Site
- ftItemData
- Handler
- classEntry
- ftBook
- Open
- State
- .encounterView
- Part
- build.sh
- errBodyWriter
- .sourceIssues
- Topic
- open5e.go
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 286 edges
2. `CN()` - 113 edges
3. `Button()` - 112 edges
4. `Sheet` - 108 edges
5. `newTestEnv()` - 106 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 52 edges
8. `Command` - 47 edges
9. `Input()` - 46 edges
10. `Slugify()` - 41 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `NewManager()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/auth/manager.go
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Open()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/db/db.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (121 total, 19 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.12
Nodes (38): contextKey, classRef, ftBookInfo, indexed, sourceOf, ASIError, clientMessage, TestMain() (+30 more)

### Community 1 - "FromContext"
Cohesion: 0.05
Nodes (24): Identity, FiveToolsBooks(), Campaign, browserName(), firstLine(), mdLinkText(), TestUserAgent(), userAgentSummary() (+16 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (87): keepAwake(), keepAwake(), ListImageCharactersRow, CountCompendiumRow, CountHabitatsRow, ListEntryCopiesRow, ListHomebrewLinksRow, UpsertCompendiumEntryParams (+79 more)

### Community 3 - ".do"
Cohesion: 0.16
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "combat_templ.go"
Cohesion: 0.12
Nodes (40): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+32 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (106): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+98 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.25
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Parse"
Cohesion: 0.08
Nodes (30): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+22 more)

### Community 10 - "Feature"
Cohesion: 0.23
Nodes (7): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (68): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+60 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput (+22 more)

### Community 14 - "CombatantView"
Cohesion: 0.10
Nodes (34): Props, Variant, Badge(), badgeClasses(), addCombatants(), beingBadge(), combatantChip(), combatantName() (+26 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (45): Note, GetBlockRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+37 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ImportView"
Cohesion: 0.19
Nodes (13): ImageCategory, categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows() (+5 more)

### Community 21 - "wizardPick"
Cohesion: 0.09
Nodes (37): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+29 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (118): ListMonstersByNameRow, Checkbox(), Content(), Description(), Empty(), Header(), Media(), Title() (+110 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "ImagePage"
Cohesion: 0.12
Nodes (26): ListEntityFieldBacklinksRow, ListBacklinksRow, Look, initials(), currentLook(), EntityDetail(), entityForm(), EntityPage() (+18 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.15
Nodes (26): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp() (+18 more)

### Community 28 - "App"
Cohesion: 0.05
Nodes (24): ClearCookie(), SetCookie(), writeCookie(), FromContext(), compendiumBack(), ImportSummary(), viewerCampaign(), fiveToolsKey() (+16 more)

### Community 29 - "testing.T"
Cohesion: 0.05
Nodes (57): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait() (+49 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "DecodeData"
Cohesion: 0.07
Nodes (16): DecodeData(), Character, CanCharacter(), coinCommand(), commandText(), hpCommand(), parseCommand(), spellFromRow() (+8 more)

### Community 34 - "testing.T"
Cohesion: 0.05
Nodes (55): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsEdition(), TestParseFiveToolsBook(), RemoveFiveTools(), readDoc(), serveMirror() (+47 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+16 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "IfElse"
Cohesion: 0.10
Nodes (47): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+39 more)

### Community 41 - "NameDesc"
Cohesion: 0.14
Nodes (25): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, SwitchFeature, names() (+17 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".levelUp"
Cohesion: 0.14
Nodes (17): Spell, addClassSpells(), errorField(), spellNeed(), classPin(), TestSpellNeed(), LevelUpView, Option (+9 more)

### Community 46 - "NewClient"
Cohesion: 0.22
Nodes (14): New(), NewClient(), randomID(), drain(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 47 - "ClassInfo"
Cohesion: 0.08
Nodes (23): featureSwitch, LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), ClassInfo (+15 more)

### Community 49 - "Slugify"
Cohesion: 0.08
Nodes (41): EquipmentItem, Cantrips(), CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash() (+33 more)

### Community 50 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "cert.go"
Cohesion: 0.19
Nodes (14): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+6 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "CampaignRoster"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 59 - "Book"
Cohesion: 0.19
Nodes (15): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools(), fetchBooks() (+7 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 63 - "Pending"
Cohesion: 0.29
Nodes (4): PendingState, Pending, PendingKind, viewerCtx

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (51): ftsQuery(), Search(), ListHomebrewByOwnerRow, ListHomebrewRow, Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow (+43 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (167): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+159 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Input"
Cohesion: 0.05
Nodes (75): Combobox(), Content(), Item(), List(), CloseFor(), Description(), Field(), Group() (+67 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "Campaign"
Cohesion: 0.22
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.16
Nodes (13): newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveSourcesAlert(), TestLiveVisibility(), TestLiveMastery(), LiveKey(), dial() (+5 more)

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

### Community 83 - "ParseWikidot"
Cohesion: 0.12
Nodes (24): block, ClassFeature, ClassData, dnd5eFeatures(), pageSource(), PageSourceKey(), cellAt(), classTable() (+16 more)

### Community 86 - "github.com/a-h/templ.Component"
Cohesion: 0.09
Nodes (37): SpellData, SearchCompendiumRow, alignmentColor(), classBadges(), classBody(), commaPrefixed(), compendiumRow(), crText() (+29 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Hub"
Cohesion: 0.25
Nodes (4): ClientInfo, event, Hub, TestCloseSessionAndCharacter()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (44): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+36 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Session"
Cohesion: 0.06
Nodes (26): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+18 more)

### Community 103 - "Site"
Cohesion: 0.14
Nodes (12): Site, dataFor(), Kind, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SavePage() (+4 more)

### Community 104 - "ftItemData"
Cohesion: 0.13
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, ItemData, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 106 - "classEntry"
Cohesion: 0.19
Nodes (14): TestSlotMaxima(), ClassLevel, CasterLevel(), SlotMaxima(), masteryLimit(), setSubclass(), subclassRef(), catchUpCasters() (+6 more)

### Community 107 - "ftBook"
Cohesion: 0.19
Nodes (14): ftBook, ftRec, each(), ftDecode(), ftMod(), ftName(), ftRecords(), ftReplaceText() (+6 more)

### Community 108 - "Open"
Cohesion: 0.15
Nodes (20): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+12 more)

### Community 111 - "State"
Cohesion: 0.14
Nodes (8): fakeHandler, patchMessage, stateHandler, viewPatch, Diff(), State, TestDiff(), liveView

### Community 113 - ".encounterView"
Cohesion: 0.10
Nodes (16): Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), initiativeMod(), outOfFight(), pcSheet() (+8 more)

### Community 115 - "Part"
Cohesion: 0.06
Nodes (23): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+15 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - ".sourceIssues"
Cohesion: 0.53
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 126 - "Topic"
Cohesion: 0.11
Nodes (10): TopicKind, Entity, Topic, entityTopics(), parseLiveKey(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams (+2 more)

### Community 133 - "open5e.go"
Cohesion: 0.08
Nodes (39): Entry, fixture, fixtures, looseString, normalizer, o5eArmor, o5eAssignment, o5eChild (+31 more)

### Community 134 - "Import"
Cohesion: 0.11
Nodes (32): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+24 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 347 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **19 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `database/sql.NullInt64`, `DecodeData`, `FromContext`, `context.Context`, `Session`, `classEntry`, `.levelUp`, `State`, `perm.go`, `Slugify`, `.encounterView`, `cert.go`, `Hub`, `.sourceIssues`, `Topic`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11552346570397112 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Part` to `compendium_templ.go`, `DecodeData`, `testing.T`, `combat_templ.go`, `SheetView`, `Parse`, `classEntry`, `Feature`, `Command`, `.levelUp`, `NameDesc`, `ClassInfo`, `.encounterView`, `Slugify`, `ParseSheet`, `.sourceIssues`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `FromContext` be split into smaller, more focused modules?**
  _Cohesion score 0.05272895467160037 - nodes in this community are weakly interconnected._
- **Why does `Slugify()` connect `Slugify` to `compendium_templ.go`, `context.Context`, `encoding/json.RawMessage`, `open5e.go`, `golang.org/x/net/html.Node`, `Import`, `Site`, `NameDesc`, `Campaign`, `ftBook`, `classEntry`, `.Entry`, `.levelUp`, `ParseWikidot`, `github.com/a-h/templ.Component`, `App`, `.sourceIssues`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010108161258603737 - nodes in this community are weakly interconnected._