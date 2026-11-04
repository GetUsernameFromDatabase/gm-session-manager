# Graph Report - gm-session-manager  (2026-11-04)

## Corpus Check
- 248 files · ~497,314 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3595 nodes · 14614 edges · 118 communities (98 shown, 20 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1461 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e5d77fb5`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- .do
- createProvider
- Queries
- SheetView
- combobox.js
- redirect.go
- testing.T
- App
- select.js
- command.go
- .Entry
- .Router
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- wikiText
- DBTX
- wikilink.js
- dialog_templ.go
- popover.js
- time.Time
- Project Rules (CLAUDE.md, README)
- ParseSheet
- Character
- toast_templ.go
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- deviceOS
- command.js
- dialog.js
- Sheet
- golang.org/x/net/html.Node
- tooltip.js
- Store
- .levelUp
- Slugify
- export/markdown.go
- composite.js
- App
- github.com/a-h/templ.Component
- .Plan
- IfElse
- Session
- github.com/a-h/templ.Attributes
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- levelUpForm
- golangci-lint Config
- NewSheet
- monsterBody
- radiogroup.js
- vendorjs Tool
- select_templ.go
- Handoff
- sqlc Config
- tooltip_templ.go
- Browser check
- Claude Settings
- NoteBlock
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
- checkbox.js
- ImportView
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- .ListCompendiumBySlugs
- Lint and Release Rules
- Run Script
- useTypeahead
- buildReview
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Site
- .ListItemsByCategory
- who
- Entry
- Import
- EntityView
- ParseWikidot
- Client
- Feature
- radiogroup_templ.go
- .ListBlockStepsSince
- build.sh
- .sendFeedback
- ArmorData

## God Nodes (most connected - your core abstractions)
1. `App` - 266 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Sheet` - 90 edges
5. `newTestEnv()` - 87 edges
6. `SheetView` - 52 edges
7. `IfElse()` - 50 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 37 edges

## Surprising Connections (you probably didn't know these)
- `Where things live (creation and level-up)` --references--> `userError`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up (pick from these)` --references--> `addClassSpells()`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Next up (pick from these)` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go
- `Where things live (creation and level-up)` --references--> `choiceGroup`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (118 total, 20 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.09
Nodes (58): Props, contextKey, looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes (+50 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.08
Nodes (19): Condition, state(), triggerID(), errorText(), TestErrorText(), conditionList(), Align, ContentProps (+11 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (82): keepAwake(), keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow (+74 more)

### Community 3 - ".do"
Cohesion: 0.12
Nodes (20): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Queries"
Cohesion: 0.15
Nodes (8): Queries, Entity, uniqueName(), entityTopics(), linkEntityLook(), linkCharacterLook(), lookFree(), lookOps

### Community 6 - "SheetView"
Cohesion: 0.08
Nodes (71): ArmorEffect(), Signed(), Variant, Action(), Card(), Content(), Description(), Header() (+63 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (65): TestCanCharacter(), nullID(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+57 more)

### Community 10 - "App"
Cohesion: 0.08
Nodes (15): Campaign, Note, coinCommand(), commandText(), hpCommand(), parseCommand(), formIntQuery(), isHTMX() (+7 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.06
Nodes (67): Amount, CondOpKind, Effects, Env, HelpEntry, parser, Completions(), condCommand() (+59 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (27): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+19 more)

### Community 14 - ".Router"
Cohesion: 0.15
Nodes (8): IsLoopback(), TestIsLoopback(), TestRequireAdmin(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.08
Nodes (51): ListCombatantsRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanCharacter() (+43 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "wikiText"
Cohesion: 0.27
Nodes (8): saveFieldLinks(), wikiLink(), wikiText(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.07
Nodes (62): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+54 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "time.Time"
Cohesion: 0.08
Nodes (18): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), folderName() (+10 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.19
Nodes (20): TestUpgradeSheet(), ParseSheet(), TestLevelUpFeats(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestSpellCounts() (+12 more)

### Community 28 - "Character"
Cohesion: 0.11
Nodes (9): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks(), nextStep() (+1 more)

### Community 29 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "database/sql.NullInt64"
Cohesion: 0.07
Nodes (31): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, CompendiumSource, CreateCombatantParams, applyCondOps() (+23 more)

### Community 34 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Sheet"
Cohesion: 0.07
Nodes (19): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+11 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.13
Nodes (24): fixture, migration, newFixture(), TestOnlyHashIsStored(), apply(), loadMigrations(), Migrate(), Open() (+16 more)

### Community 41 - ".levelUp"
Cohesion: 0.09
Nodes (35): Where things live (creation and level-up), ClassFeature, ClassInfo, LevelUpPlan, Spell, addClassSpells(), classFeatures(), classInfo() (+27 more)

### Community 42 - "Slugify"
Cohesion: 0.09
Nodes (35): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, TextData (+27 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.15
Nodes (17): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+9 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "App"
Cohesion: 0.06
Nodes (11): DecodeData(), ftsQuery(), Search(), Combatant, Encounter, Character, compendiumWeapon(), viewerCampaign() (+3 more)

### Community 46 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (115): ListMonstersByNameRow, Description(), Button(), Content(), Description(), Empty(), Header(), Media() (+107 more)

### Community 47 - ".Plan"
Cohesion: 0.17
Nodes (8): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), hpGain(), raise(), CasterLevel(), SlotMaxima()

### Community 48 - "IfElse"
Cohesion: 0.06
Nodes (69): Props, Variant, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), Content(), contentID() (+61 more)

### Community 49 - "Session"
Cohesion: 0.16
Nodes (10): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+2 more)

### Community 50 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (157): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Size, ActionProps (+149 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "LoadOrCreateCert"
Cohesion: 0.10
Nodes (18): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+10 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "levelUpForm"
Cohesion: 0.14
Nodes (26): Props, Checkbox(), initialChecked(), levelUpForm(), slotSummary(), spellChecklist(), checkField(), abilityForm() (+18 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "NewSheet"
Cohesion: 0.08
Nodes (26): Defense, Next up (pick from these), Parse(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+18 more)

### Community 59 - "monsterBody"
Cohesion: 0.13
Nodes (28): SpellData, classBody(), commaPrefixed(), entryBody(), EntryDetails(), entryHead(), EntryPage(), EntryPreview() (+20 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "select_templ.go"
Cohesion: 0.11
Nodes (25): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+17 more)

### Community 63 - "Handoff"
Cohesion: 0.25
Nodes (7): Handoff, Improvements (bigger, deferred), Notes, Test data in data-test (all safe to delete), Working on the test instance, fieldError, userError

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "tooltip_templ.go"
Cohesion: 0.14
Nodes (26): Field, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks() (+18 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "builder_templ.go"
Cohesion: 0.05
Nodes (73): ImageCategory, ListBacklinksRow, CloseFor(), Description(), Field(), Group(), Label(), Input() (+65 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.12
Nodes (19): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.15
Nodes (14): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), TestPlayerCombatantLink(), ImageSrc(), LiveKey() (+6 more)

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
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "buildReview"
Cohesion: 0.17
Nodes (7): buildReview(), csvSource(), formValues(), parseCSV(), capReader, csvImport, csvImports

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.14
Nodes (20): fixture, ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList() (+12 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "Site"
Cohesion: 0.12
Nodes (17): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+9 more)

### Community 105 - "who"
Cohesion: 0.10
Nodes (7): GetBlockRow, compendiumBack(), HomebrewView, LinkItem, SwitcherResults(), blockAccessLevel, who

### Community 107 - "Entry"
Cohesion: 0.16
Nodes (17): Entry, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children(), firstOr() (+9 more)

### Community 108 - "Import"
Cohesion: 0.16
Nodes (20): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+12 more)

### Community 109 - "EntityView"
Cohesion: 0.10
Nodes (29): ListEntityFieldBacklinksRow, Look, attrString(), generateSVG(), getIconContent(), Props, Icon(), Item() (+21 more)

### Community 111 - "ParseWikidot"
Cohesion: 0.11
Nodes (24): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, InputForEntry(), TestMonsterFormRoundTrip(), partsText() (+16 more)

### Community 112 - "Client"
Cohesion: 0.05
Nodes (38): ClientInfo, event, fakeHandler, Handler, message, patchMessage, PendingState, stateHandler (+30 more)

### Community 113 - "Feature"
Cohesion: 0.08
Nodes (27): EquipmentItem, FeatureEffects, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions() (+19 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 334 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **20 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `database/sql.NullInt64`, `.do`, `Queries`, `who`, `App`, `Slugify`, `.levelUp`, `.Router`, `Client`, `perm.go`, `Feature`, `LoadOrCreateCert`, `buildReview`, `time.Time`, `Character`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09454292665971498 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `SheetView`, `.levelUp`, `Slugify`, `command.go`, `App`, `.Plan`, `perm.go`, `Feature`, `NewSheet`, `ParseSheet`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `popover_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.07586206896551724 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `.levelUp` to `compendium_templ.go`, `Feature`, `Slugify`, `.Plan`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009078372787864858 - nodes in this community are weakly interconnected._