# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 374 files · ~647,011 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4724 nodes · 24458 edges · 158 communities (121 shown, 37 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2451 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `08fcbb15`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- monsterBody
- context.Context
- vendorjs/main.go
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- Condition
- Path
- encoding/json.RawMessage
- select.js
- Command
- ItemData
- rosterCard
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- IfElse
- wikilink.js
- LinkTarget
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- storeSource
- DownloadFiveTools
- app.js
- tabbable.js
- useDismiss
- NewClient
- popover_templ.go
- command.js
- dialog.js
- vendorjs/main.go
- ParseWikidot
- tooltip.js
- CompendiumKind
- Sheet
- viewerCtx
- .Router
- composite.js
- buildReview
- CustomMod
- planSwitch
- alert_templ.go
- deathFix
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- ParseSheet
- shadcn-templ Config
- Item
- golangci-lint Config
- ftRec
- Part
- radiogroup.js
- Player
- ftItemData
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- State
- use_transition_status.js
- playwright
- TestSourceSwitchKeepsEdits
- Defenses
- checkbox.js
- ParseAidedd
- switch.js
- InputForEntry
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
- Handlers
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- Client
- mounted
- .sendFeedback
- nullID
- profFocus
- IfElse
- redirect.go
- ftBook
- CharacterFile
- BlockEditor
- .LevelUp
- github.com/a-h/templ.Component
- capReader
- .features
- time.Time
- build.sh
- Import
- looseString
- errBodyWriter
- ListSharedWithRow
- export/markdown.go
- Before/after pairs
- ParseArmorTraining
- Hub
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- config.go
- ASIError
- Book
- .Entry
- EffectFieldsURL
- sourceUsedError
- Pending
- fakeHandler
- Campaign
- Layout
- TestPagesRender
- parseMod
- condMods
- WeaponKind
- spellCounts
- SearchReadableBlocksRow
- .ListBlockStepsSince
- .ListSpellsForClass
- .ListSubracesOf

## God Nodes (most connected - your core abstractions)
1. `App` - 374 edges
2. `Path()` - 270 edges
3. `newTestEnv()` - 172 edges
4. `Sheet` - 132 edges
5. `Button()` - 124 edges
6. `CN()` - 114 edges
7. `SheetView` - 78 edges
8. `IfElse()` - 65 edges
9. `Command` - 60 edges
10. `CompendiumKind` - 56 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (158 total, 37 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "monsterBody"
Cohesion: 0.08
Nodes (39): SearchCompendiumRow, classBadges(), pickerRow(), rowBadges(), SearchComboItems(), SearchResults(), textMatchesHeader(), TestRowFigure() (+31 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (111): keepAwake(), TestHomebrew(), bindFiveToolsImages(), ftsQuery(), Search(), ListHomebrewByOwnerRow, ListHomebrewRow, ListImagesRow (+103 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.09
Nodes (42): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+34 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (125): ArmorEffect(), Coins, Slot, Signed(), Button(), Variant, Action(), Card() (+117 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Condition"
Cohesion: 0.05
Nodes (46): Condition, CreateConditionParams, ListCombatantsRow, Character, EffectEnd, passTurn(), rollFor(), stepTurn() (+38 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (89): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+81 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (39): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+31 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (51): Amount, CondOpKind, Effects, userText, parseFeature(), parseSet(), parseSpell(), parseCoin() (+43 more)

### Community 13 - "ItemData"
Cohesion: 0.05
Nodes (58): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+50 more)

### Community 14 - "rosterCard"
Cohesion: 0.36
Nodes (9): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), deviceLabel(), Device (+1 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (56): ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+48 more)

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
Cohesion: 0.03
Nodes (49): ClearCookie(), SetCookie(), writeCookie(), FromContext(), FiveToolsBooks(), GetImageRow, Campaign, Character (+41 more)

### Community 21 - "IfElse"
Cohesion: 0.09
Nodes (47): Props, Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup(), state() (+39 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "LinkTarget"
Cohesion: 0.17
Nodes (8): LinkTarget, AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink, snapshotBody

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (157): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+149 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.07
Nodes (28): fixture, tokens, main(), run(), constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens() (+20 more)

### Community 28 - "storeSource"
Cohesion: 0.10
Nodes (19): HomebrewOwner, Site, Source, sourceFormat, fiveToolsPage(), HomebrewInput, SaveHomebrew(), SeedHabitats() (+11 more)

### Community 29 - "DownloadFiveTools"
Cohesion: 0.15
Nodes (21): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestImportFiveToolsBook(), TestParseFiveToolsBook() (+13 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "NewClient"
Cohesion: 0.22
Nodes (19): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+11 more)

### Community 34 - "popover_templ.go"
Cohesion: 0.14
Nodes (26): UI conventions, Content(), Description(), Header(), Popover(), state(), Title(), Trigger() (+18 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "vendorjs/main.go"
Cohesion: 0.10
Nodes (27): migration, newFixture(), TestOnlyHashIsStored(), TestResolveDeletedCharacter(), TestResolveTouchesStaleSession(), apply(), loadMigrations(), Migrate() (+19 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.10
Nodes (30): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, dnd5eFeatures(), TestFiveToolsCoreTraitsShapes(), withCoreTraits() (+22 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.04
Nodes (49): EquipmentItem, EquipmentOptions(), EquipmentOption, dataFor(), DecodeData(), ValidKind(), CountCompendiumRow, UpsertCompendiumEntryParams (+41 more)

### Community 41 - "Sheet"
Cohesion: 0.03
Nodes (110): LevelUpChoice, internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice() (+102 more)

### Community 42 - "viewerCtx"
Cohesion: 0.22
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 43 - ".Router"
Cohesion: 0.23
Nodes (4): IsLoopback(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 46 - "CustomMod"
Cohesion: 0.10
Nodes (16): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+8 more)

### Community 47 - "planSwitch"
Cohesion: 0.15
Nodes (11): featureSwitch, featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch(), sameSlug() (+3 more)

### Community 48 - "alert_templ.go"
Cohesion: 0.05
Nodes (53): ActionProps, DescriptionProps, Props, TitleProps, Variant, ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow (+45 more)

### Community 49 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (130): Defense, TestIsLoopback(), TestDecode(), TestEquipmentOptions(), TestSkillChoice(), feature(), RemoveSourceDir(), TestRemoveSourceDir() (+122 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "ParseSheet"
Cohesion: 0.10
Nodes (40): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeatRaisesCon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass() (+32 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Item"
Cohesion: 0.12
Nodes (8): findItem(), parseArmor(), TestCustomMods(), TestExplainSums(), Sum(), TestGearCommands(), ensureItem(), Item

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "ftRec"
Cohesion: 0.24
Nodes (9): ftCopyRef, ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve() (+1 more)

### Community 59 - "Part"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (37): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views (+29 more)

### Community 62 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.09
Nodes (39): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+31 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "State"
Cohesion: 0.09
Nodes (14): event, Handler, message, patchMessage, stateHandler, viewPatch, Topic, TopicKind (+6 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 72 - "Defenses"
Cohesion: 0.43
Nodes (5): Defenses, capitalized(), hpQuickForm(), petrifiedNote(), damageTypeItems()

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "InputForEntry"
Cohesion: 0.29
Nodes (7): habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

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
Cohesion: 0.07
Nodes (33): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone() (+25 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 103 - "nullID"
Cohesion: 0.20
Nodes (21): TestMonsterAtZeroHPUnaffected(), acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectFields() (+13 more)

### Community 105 - "IfElse"
Cohesion: 0.05
Nodes (85): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, CountHabitatsRow (+77 more)

### Community 106 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 107 - "ftBook"
Cohesion: 0.17
Nodes (20): ftBook, each(), ftCastingTime(), ftParts(), ftPrerequisite(), ftRecords(), TestFiveToolsCopy(), TestStripTagsTokens() (+12 more)

### Community 108 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 111 - ".LevelUp"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 112 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (102): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+94 more)

### Community 116 - ".features"
Cohesion: 0.24
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 117 - "time.Time"
Cohesion: 0.05
Nodes (27): Identity, ListAttachmentFilesRow, Session, Role, browserName(), firstLine(), mdLinkText(), fresh() (+19 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "Import"
Cohesion: 0.11
Nodes (35): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+27 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 127 - "export/markdown.go"
Cohesion: 0.13
Nodes (19): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+11 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "ParseArmorTraining"
Cohesion: 0.13
Nodes (13): armorGrants(), ArmorKindNamed(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack() (+5 more)

### Community 132 - "Hub"
Cohesion: 0.24
Nodes (4): ClientInfo, goneDevice, TestFiveToolsClassLevelUp(), Hub

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (23): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable() (+15 more)

### Community 137 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 139 - "Book"
Cohesion: 0.14
Nodes (19): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+11 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 141 - "EffectFieldsURL"
Cohesion: 0.33
Nodes (4): EffectFieldsURL(), EffectID(), effectQuery(), EffectTarget

### Community 143 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 145 - "Campaign"
Cohesion: 0.32
Nodes (5): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

### Community 146 - "Layout"
Cohesion: 0.15
Nodes (29): Spinner(), CharacterPage(), EntitiesPage(), HomebrewPage(), ImportReviewPage(), ImportUploadPage(), ImagesPage(), importButton() (+21 more)

### Community 147 - "TestPagesRender"
Cohesion: 0.25
Nodes (7): EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Gotchas, internal/web, TestPagesRender()

### Community 148 - "parseMod"
Cohesion: 0.50
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 151 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

### Community 152 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 153 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 418 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **37 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Hub`, `Condition`, `Book`, `Pending`, `Viewer`, `Campaign`, `cert.go`, `storeSource`, `CompendiumKind`, `Sheet`, `viewerCtx`, `.Router`, `testing.T`, `FiveToolsJob`, `State`, `.q`, `time.Time`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14339294192119006 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `monsterBody`, `tabs_templ.go`, `SheetView`, `Condition`, `ItemData`, `rosterCard`, `EffectFieldsURL`, `Viewer`, `Campaign`, `Layout`, `TestPagesRender`, `App`, `IfElse`, `github.com/a-h/templ.Attributes`, `popover_templ.go`, `CompendiumKind`, `Sheet`, `alert_templ.go`, `deathFix`, `testing.T`, `ParseSheet`, `Player`, `CombatantView`, `TestSourceSwitchKeepsEdits`, `ParseWiki`, `.q`, `.sendFeedback`, `nullID`, `IfElse`, `github.com/a-h/templ.Component`, `time.Time`, `Import`, `export/markdown.go`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **Should `monsterBody` be split into smaller, more focused modules?**
  _Cohesion score 0.08205128205128205 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `TestPagesRender`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.025) - this node is a cross-community bridge._