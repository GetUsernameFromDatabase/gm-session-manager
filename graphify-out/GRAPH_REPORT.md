# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 375 files · ~648,913 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4740 nodes · 24514 edges · 144 communities (113 shown, 31 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2462 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `acf4e43a`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- spellBoxes
- context.Context
- vendorjs/main.go
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- database/sql.NullInt64
- newTestEnv
- encoding/json.RawMessage
- select.js
- Command
- dialog_templ.go
- Import
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- github.com/a-h/templ.Component
- wikilink.js
- Character
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- storeSource
- Input
- app.js
- tabbable.js
- useDismiss
- Layout
- github.com/a-h/templ.Component
- command.js
- dialog.js
- Open
- ParseWikidot
- tooltip.js
- CompendiumKind
- Slugify
- .levelUp
- viewerCtx
- composite.js
- tooltip_templ.go
- CustomMod
- Env
- CN
- select_templ.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- .levelUpView
- shadcn-templ Config
- Session
- golangci-lint Config
- entityForm
- Sheet
- radiogroup.js
- Player
- DownloadFiveTools
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Client
- use_transition_status.js
- playwright
- Attachment
- Handlers
- checkbox.js
- ParseAidedd
- switch.js
- .renderCompendium
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- .q
- useTypeahead
- ftRecords
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- Entity
- mounted
- .sendFeedback
- nullID
- profFocus
- .equipmentChoices
- deathFix
- ftBook
- TestPagesRender
- BlockEditor
- .LevelUp
- .sourceIssues
- capReader
- .features
- time.Time
- build.sh
- parseCommand
- looseString
- errBodyWriter
- ListSharedWithRow
- export/markdown.go
- Before/after pairs
- rowFigure
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ASIError
- Book
- .Entry
- sourceUsedError
- rosterCard
- ClassInfo

## God Nodes (most connected - your core abstractions)
1. `App` - 375 edges
2. `Path()` - 271 edges
3. `newTestEnv()` - 175 edges
4. `Sheet` - 133 edges
5. `Button()` - 124 edges
6. `CN()` - 114 edges
7. `SheetView` - 78 edges
8. `IfElse()` - 65 edges
9. `Command` - 60 edges
10. `Queries` - 56 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (144 total, 31 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "spellBoxes"
Cohesion: 0.67
Nodes (4): spellBoxes(), TestTaughtPicked(), taughtPicked(), taughtRules()

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (79): keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Note, SearchReadableBlocksRow, EntityKind (+71 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.20
Nodes (18): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+10 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.10
Nodes (35): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+27 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (105): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+97 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (79): HomebrewOwner, SaveHomebrew(), SeedHabitats(), Queries, Combatant, Condition, Encounter, CreateCombatantParams (+71 more)

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (112): TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestConditionExpiresOnTurn() (+104 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (33): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+25 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (57): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+49 more)

### Community 13 - "dialog_templ.go"
Cohesion: 0.10
Nodes (42): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+34 more)

### Community 14 - "Import"
Cohesion: 0.12
Nodes (30): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+22 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (58): ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+50 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "App"
Cohesion: 0.04
Nodes (26): SetCookie(), FromContext(), Campaign, compendiumPath(), pathCampaign(), collectHeld(), formInt(), formIntQuery() (+18 more)

### Community 21 - "github.com/a-h/templ.Component"
Cohesion: 0.08
Nodes (67): ctxState, Field, Alert(), Description(), Title(), Checkbox(), Chip(), Chips() (+59 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "Character"
Cohesion: 0.05
Nodes (21): Character, Character, NewCharacterFile(), CanCharacter(), TestCanCharacter(), checkSkillPick(), classSkillPick(), postedSkills() (+13 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (147): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+139 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.05
Nodes (34): main(), run(), Config, lookupFunc, NewManager(), describe(), fromLookup(), Load() (+26 more)

### Community 28 - "storeSource"
Cohesion: 0.11
Nodes (20): Site, Source, sourceFormat, TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), formatOf(), Report, ImportProgress() (+12 more)

### Community 29 - "Input"
Cohesion: 0.06
Nodes (79): ImageCategory, UI conventions, CloseFor(), Description(), Field(), Group(), Label(), Input() (+71 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Layout"
Cohesion: 0.08
Nodes (44): Scripts(), Separator(), Toast(), Toaster(), toastIcon(), toastListItem(), campaignLink(), PendingRequests() (+36 more)

### Community 34 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (88): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, Input(), Addon(), Button() (+80 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Open"
Cohesion: 0.06
Nodes (41): fixture, tokens, migration, newFixture(), serve(), TestMiddlewareIdentity(), TestOnlyHashIsStored(), TestRequireAdmin() (+33 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (23): block, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.04
Nodes (41): dataFor(), DecodeData(), ValidKind(), HomebrewInput, InputForEntry(), CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewRow (+33 more)

### Community 41 - "Slugify"
Cohesion: 0.08
Nodes (44): internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice(), TestOptions() (+36 more)

### Community 42 - ".levelUp"
Cohesion: 0.10
Nodes (22): FreeCast(), SwapFeatSpell(), LevelUpPlan, Spell, spellFromRow(), addClassSpells(), spellNeed(), TestSpellNeed() (+14 more)

### Community 43 - "viewerCtx"
Cohesion: 0.12
Nodes (15): TestHomebrew(), ftsQuery(), Search(), SearchCompendiumRow, notFoundIsNil(), parseLiveKey(), TestLiveAffects(), Roster (+7 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "tooltip_templ.go"
Cohesion: 0.10
Nodes (29): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks(), BuilderPick (+21 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (22): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), AddFeat(), SaveFeature() (+14 more)

### Community 47 - "Env"
Cohesion: 0.14
Nodes (12): featureSwitch, ftRaceAbility(), featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch() (+4 more)

### Community 48 - "CN"
Cohesion: 0.07
Nodes (60): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Entity (+52 more)

### Community 49 - "select_templ.go"
Cohesion: 0.11
Nodes (23): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state() (+15 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (125): Defense, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestIsLoopback(), TestDecode(), TestTokensReadAsLabels(), feature() (+117 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - ".levelUpView"
Cohesion: 0.11
Nodes (18): TextData, ListSubclassesOfRow, Spellbook(), TestSpellbook(), CountWord(), textMarkdown(), takenFeat(), TestSubclassPanels() (+10 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Session"
Cohesion: 0.10
Nodes (13): IsLoopback(), ClearCookie(), writeCookie(), Identity, Manager, RequireAdmin(), Session, Role (+5 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "entityForm"
Cohesion: 0.16
Nodes (20): ListEntityFieldBacklinksRow, ListBacklinksRow, attrString(), generateSVG(), getIconContent(), Props, Icon(), EntityDetail() (+12 more)

### Community 59 - "Sheet"
Cohesion: 0.05
Nodes (25): Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey(), Sum() (+17 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (36): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views (+28 more)

### Community 62 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (24): habitatsFrom(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants() (+16 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.09
Nodes (49): Props, Variant, Defenses, Badge(), badgeClasses(), Header(), Title(), capitalized() (+41 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Client"
Cohesion: 0.05
Nodes (45): ClientInfo, countingHandler, event, fakeHandler, goneDevice, Handler, message, patchMessage (+37 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "Attachment"
Cohesion: 0.13
Nodes (4): Attachment, firstLine(), mdLinkText(), CreateAttachmentParams

### Community 72 - "Handlers"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - ".renderCompendium"
Cohesion: 0.22
Nodes (4): ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers()

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "ParseWiki"
Cohesion: 0.22
Nodes (10): wikiLink(), wikiText(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks(), WikiLinks() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "notes.js"
Cohesion: 0.15
Nodes (11): attachedPicture(), ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), nodeActions() (+3 more)

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - ".q"
Cohesion: 0.08
Nodes (33): isCleared(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+25 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ftRecords"
Cohesion: 0.24
Nodes (10): ftRec, ftCopyTargets(), ftDecode(), ftRecords(), TestFiveToolsCopy(), ftMod(), ftName(), ftReplaceText() (+2 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "Entity"
Cohesion: 0.18
Nodes (6): Entity, entityTopics(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams, lookOps

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 103 - "nullID"
Cohesion: 0.11
Nodes (37): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), acBonusForm(), effectFix, newEffectFix(), notes() (+29 more)

### Community 105 - ".equipmentChoices"
Cohesion: 0.24
Nodes (9): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, TestCoreTraits(), TestEquipmentOptions(), TestFiveToolsClasses(), applyEquipment() (+1 more)

### Community 106 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 107 - "ftBook"
Cohesion: 0.13
Nodes (27): ftBook, each(), ftCastingTime(), ftCell(), ftFeatIncrease(), ftParts(), ftPrerequisite(), ftTable() (+19 more)

### Community 108 - "TestPagesRender"
Cohesion: 0.25
Nodes (7): EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Gotchas, internal/web, TestPagesRender()

### Community 111 - ".LevelUp"
Cohesion: 0.07
Nodes (22): LevelUpChoice, armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorGrants() (+14 more)

### Community 112 - ".sourceIssues"
Cohesion: 0.36
Nodes (7): pinnedRefs(), sourceIssue(), featKey(), pinFeat(), SourceCopy, SourceIssue, sheetRef

### Community 116 - ".features"
Cohesion: 0.11
Nodes (17): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ItemData, ftArmor(), ftCode() (+9 more)

### Community 117 - "time.Time"
Cohesion: 0.04
Nodes (35): GetImageRow, ListImagesRow, Handlers, browserName(), TestUserAgent(), userAgentSummary(), imageItem(), imageSource() (+27 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "parseCommand"
Cohesion: 0.29
Nodes (7): featureCommand(), coinCommand(), commandText(), hpCommand(), parseCommand(), damageArgs(), modRows()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 127 - "export/markdown.go"
Cohesion: 0.05
Nodes (48): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+40 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "rowFigure"
Cohesion: 0.50
Nodes (4): TestRowFigure(), crText(), rowFigure(), textOnly()

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery() (+16 more)

### Community 139 - "Book"
Cohesion: 0.10
Nodes (27): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), CleanDownloads(), downloadBook(), fetchBooks() (+19 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 144 - "rosterCard"
Cohesion: 0.39
Nodes (5): devicesWhere(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 157 - "ClassInfo"
Cohesion: 0.09
Nodes (33): ClassFeature, Decode(), ClassData, SpellData, dnd5eFeatures(), TestSlotMaxima(), ClassFeature, ClassInfo (+25 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 419 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **31 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `newTestEnv` to `compendium_templ.go`, `tabs_templ.go`, `SheetView`, `dialog_templ.go`, `Import`, `Viewer`, `App`, `github.com/a-h/templ.Component`, `Character`, `Input`, `Layout`, `github.com/a-h/templ.Component`, `Open`, `CompendiumKind`, `.levelUp`, `tooltip_templ.go`, `CN`, `testing.T`, `entityForm`, `Player`, `CombatantView`, `Attachment`, `ParseWiki`, `.q`, `.sendFeedback`, `nullID`, `deathFix`, `TestPagesRender`, `time.Time`, `export/markdown.go`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14025614291716593 - nodes in this community are weakly interconnected._
- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `database/sql.NullInt64`, `Book`, `Command`, `Viewer`, `Character`, `cert.go`, `ClassInfo`, `CompendiumKind`, `Slugify`, `.levelUp`, `viewerCtx`, `testing.T`, `.levelUpView`, `Session`, `Sheet`, `FiveToolsJob`, `Client`, `Attachment`, `.renderCompendium`, `.q`, `Entity`, `.equipmentChoices`, `.sourceIssues`, `time.Time`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010319780317122863 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `newTestEnv`, `TestPagesRender`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.025) - this node is a cross-community bridge._