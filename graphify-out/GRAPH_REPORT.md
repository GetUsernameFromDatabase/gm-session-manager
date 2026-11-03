# Graph Report - gm-session-manager  (2026-11-03)

## Corpus Check
- 246 files · ~472,679 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3461 nodes · 13715 edges · 126 communities (107 shown, 19 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1336 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `cdcff7f1`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- select_templ.go
- context.Context
- testEnv
- createProvider
- database/sql.NullInt64
- SheetView
- combobox.js
- redirect.go
- .do
- dialog_templ.go
- select.js
- command.go
- .Entry
- .encounterView
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ImagesPage
- .fail
- wikilink.js
- combat_templ.go
- popover.js
- findItem
- Project Rules (CLAUDE.md, README)
- time.Time
- 5e Sheet Model
- github.com/a-h/templ.Component
- app.js
- tabbable.js
- useDismiss
- Weapon
- wizardPick
- command.js
- dialog.js
- .loadLive
- golang.org/x/net/html.Node
- tooltip.js
- Store
- ParseWikidot
- ParseEffects
- export/markdown.go
- composite.js
- RosterCharacter
- CN
- .Plan
- cell
- Manager
- Parse
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- lanHostname
- golangci-lint Config
- toast_templ.go
- .sendFeedback
- radiogroup.js
- vendorjs Tool
- spellCounts
- .blockFor
- sqlc Config
- github.com/a-h/templ.Attributes
- Browser check
- Claude Settings
- Parse
- use_transition_status.js
- playwright
- release.sh
- builderFields
- checkbox.js
- CustomMod
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- testing.T
- useFocus
- mark_others.js
- Lint and Release Rules
- Run Script
- useTypeahead
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- Character
- Campaign
- Entry
- config.go
- State
- open5e.go
- Client
- radiogroup_templ.go
- Hub
- Verbs
- New
- build.sh
- buildReview
- wikiText
- Pending
- App
- Handoff: open items from the 2026-11-03 review
- Props
- Topic
- IfElse
- NoteListView
- capReader

## God Nodes (most connected - your core abstractions)
1. `App` - 247 edges
2. `CN()` - 112 edges
3. `Button()` - 103 edges
4. `newTestEnv()` - 75 edges
5. `Sheet` - 58 edges
6. `SheetView` - 50 edges
7. `IfElse()` - 48 edges
8. `Input()` - 46 edges
9. `Command` - 43 edges
10. `Client` - 35 edges

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

## Communities (126 total, 19 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (38): contextKey, keepAwake(), clientMessage, TestMain(), newValidator(), TestMain(), isToggle(), CharacterImage (+30 more)

### Community 1 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+16 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (67): keepAwake(), HomebrewOwner, TestHomebrew(), ftsQuery(), SaveHomebrew(), Search(), fetch(), ListImageCharactersRow (+59 more)

### Community 3 - "testEnv"
Cohesion: 0.16
Nodes (7): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (43): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, GetImageRow, ListImagesRow, Condition, Entity, CreateCombatantParams (+35 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (104): ArmorEffect(), Signed(), Field, Button(), Variant, Action(), Card(), Content() (+96 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - ".do"
Cohesion: 0.16
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 10 - "dialog_templ.go"
Cohesion: 0.15
Nodes (33): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+25 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.10
Nodes (48): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+40 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (26): MonsterAction, addEntry(), setStat(), MonsterData, NameDesc, ValidKind(), HomebrewInput, HomebrewInputFor() (+18 more)

### Community 14 - ".encounterView"
Cohesion: 0.12
Nodes (12): Combatant, ListCombatantsRow, Character, initiativeMod(), outOfFight(), pcSheet(), stepTurn(), uniqueName() (+4 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.17
Nodes (28): VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock(), CanManageNote() (+20 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ImagesPage"
Cohesion: 0.08
Nodes (36): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Look, selectItems(), homebrewLinks(), currentLook() (+28 more)

### Community 21 - ".fail"
Cohesion: 0.10
Nodes (9): Character, CanCharacter(), linkCharacterLook(), Flash(), nextStep(), wizardPath(), SetCharacterCurrentImageParams, ToggleCharacterImageRevealedParams (+1 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "combat_templ.go"
Cohesion: 0.09
Nodes (47): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+39 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "findItem"
Cohesion: 0.38
Nodes (3): findItem(), parseArmor(), ensureItem()

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "time.Time"
Cohesion: 0.07
Nodes (17): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), mdLinkText() (+9 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (14): DeathSaves, TestExplainSums(), Sum(), TestGearCommands(), TestUpgradeSheet(), Coins, Item, Sheet (+6 more)

### Community 29 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (107): SpellData, SearchCompendiumRow, Badge(), Scripts(), Separator(), Spinner(), campaignLink(), CampaignOverview() (+99 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Weapon"
Cohesion: 0.12
Nodes (11): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+3 more)

### Community 34 - "wizardPick"
Cohesion: 0.10
Nodes (32): Label(), categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportRow, ImportView, ImportReview() (+24 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - ".loadLive"
Cohesion: 0.14
Nodes (10): Fields, ParseFields(), saveFieldLinks(), notFoundIsNil(), parseLiveKey(), toLook(), toEntity(), LinkItem (+2 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.09
Nodes (34): migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+26 more)

### Community 41 - "ParseWikidot"
Cohesion: 0.12
Nodes (23): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, InputForEntry(), TestMonsterFormRoundTrip(), withCoreTraits() (+15 more)

### Community 42 - "ParseEffects"
Cohesion: 0.12
Nodes (23): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+15 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 46 - "CN"
Cohesion: 0.06
Nodes (68): Props, Props, MediaProps, MediaVariant, ListMonstersByNameRow, AspectRatio(), Checkbox(), initialChecked() (+60 more)

### Community 47 - ".Plan"
Cohesion: 0.14
Nodes (12): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan, hpGain() (+4 more)

### Community 48 - "cell"
Cohesion: 0.13
Nodes (17): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+9 more)

### Community 49 - "Manager"
Cohesion: 0.11
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+10 more)

### Community 50 - "Parse"
Cohesion: 0.33
Nodes (7): Parse(), tokenize(), TestUnknownCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

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
Cohesion: 0.14
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 59 - ".sendFeedback"
Cohesion: 0.33
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "spellCounts"
Cohesion: 0.67
Nodes (3): sortedColumns(), spellCounts(), TestSpellCounts()

### Community 63 - ".blockFor"
Cohesion: 0.12
Nodes (6): Attachment, GetBlockRow, saveScreenshot(), saveSniffed(), CreateAttachmentParams, stepsPayload

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (113): Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props (+105 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Parse"
Cohesion: 0.12
Nodes (15): Completions(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages(), TestLevelUp(), TestLevelUpSubclassAndPools() (+7 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "builderFields"
Cohesion: 0.06
Nodes (77): Combobox(), Content(), Item(), List(), CloseFor(), Description(), Field(), Group() (+69 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "CustomMod"
Cohesion: 0.16
Nodes (13): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+5 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.17
Nodes (12): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), LiveKey(), dial(), patchOf() (+4 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "testing.T"
Cohesion: 0.05
Nodes (75): Import(), TestCanCharacter(), TestCanManageBlock(), TestCanReadAttachment(), TestCanReadImage(), TestCanSeeCombatantHP(), TestVisibleEntity(), TestVisibleLooks() (+67 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.21
Nodes (16): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+8 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 103 - "Site"
Cohesion: 0.12
Nodes (17): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+9 more)

### Community 105 - "Character"
Cohesion: 0.06
Nodes (25): DecodeData(), Slugify(), CompendiumEntry, CompendiumSource, ParseSheet(), Campaign(), CharacterFileName(), fileSlug() (+17 more)

### Community 106 - "Campaign"
Cohesion: 0.18
Nodes (7): Campaign, deviceOS(), isPhone(), uaPick(), userAgentSummary(), Roster, CreateCampaignParams

### Community 107 - "Entry"
Cohesion: 0.27
Nodes (9): Entry, fixtures, normalizer, byPK(), children(), joinText(), label(), parseFixtures() (+1 more)

### Community 108 - "config.go"
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 109 - "State"
Cohesion: 0.24
Nodes (7): patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff(), liveView

### Community 110 - "open5e.go"
Cohesion: 0.07
Nodes (30): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+22 more)

### Community 112 - "Client"
Cohesion: 0.17
Nodes (4): fakeHandler, message, Client, topicIDs()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.19
Nodes (14): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+6 more)

### Community 117 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "buildReview"
Cohesion: 0.24
Nodes (6): buildReview(), csvSource(), formValues(), parseCSV(), csvImport, csvImports

### Community 123 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 124 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 125 - "App"
Cohesion: 0.06
Nodes (16): compendiumBack(), viewerCampaign(), entityTopics(), formIntQuery(), isHTMX(), redirect(), securityHeaders(), joinPath() (+8 more)

### Community 126 - "Handoff: open items from the 2026-11-03 review"
Cohesion: 0.25
Nodes (7): Bugs still open, Character page, Handoff: open items from the 2026-11-03 review, Improvements, Notes, Other pages, Test data in data-test

### Community 129 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 130 - "Topic"
Cohesion: 0.18
Nodes (5): event, Handler, stateHandler, TopicKind, Topic

### Community 131 - "IfElse"
Cohesion: 0.09
Nodes (48): CountHabitatsRow, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+40 more)

### Community 132 - "NoteListView"
Cohesion: 0.50
Nodes (3): SearchReadableBlocksRow, SearchReadableBlocksParams, NoteListView

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **19 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `testEnv`, `.loadLive`, `database/sql.NullInt64`, `.do`, `Character`, `Campaign`, `State`, `.encounterView`, `perm.go`, `Hub`, `.fail`, `cert.go`, `buildReview`, `time.Time`, `.blockFor`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09840642449338101 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `compendium_templ.go`, `Weapon`, `wizardPick`, `Parse`, `SheetView`, `Character`, `CustomMod`, `ParseEffects`, `.encounterView`, `.Plan`, `findItem`?**
  _High betweenness centrality (0.019) - this node is a cross-community bridge._
- **Should `select_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10541310541310542 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Hub` to `compendium_templ.go`, `Topic`, `Client`, `New`, `Pending`, `App`?**
  _High betweenness centrality (0.014) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011873920552677029 - nodes in this community are weakly interconnected._