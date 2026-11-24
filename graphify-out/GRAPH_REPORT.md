# Graph Report - gm-session-manager  (2026-11-24)

## Corpus Check
- 379 files · ~651,238 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4804 nodes · 24734 edges · 175 communities (122 shown, 53 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2479 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `bfa2accb`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Path
- context.Context
- vendorjs/main.go
- createProvider
- EquipmentOption
- SheetView
- combobox.js
- database/sql.NullInt64
- Path
- encoding/json.RawMessage
- select.js
- Command
- Input
- redirect.go
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- WizardView
- wikilink.js
- Condition
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- who
- select_templ.go
- app.js
- tabbable.js
- useDismiss
- Slugify
- monsterBody
- command.js
- dialog.js
- Values Go and JS share
- ParseWikidot
- tooltip.js
- Character
- tabs_templ.go
- Queries
- viewerCtx
- composite.js
- ParseArmorTraining
- CustomMod
- Env
- IfElse
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Sheet
- shadcn-templ Config
- Session
- golangci-lint Config
- popover_templ.go
- Part
- radiogroup.js
- Player
- .data
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- State
- use_transition_status.js
- playwright
- ftPrerequisite
- NameDesc
- checkbox.js
- ParseAidedd
- switch.js
- Client
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
- NewClient
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- .compendiumDir
- mounted
- Hub
- nullID
- profFocus
- Store
- deathFix
- ftBook
- Pending
- BlockEditor
- .LevelUp
- sourcesEnv
- capReader
- .features
- time.Time
- build.sh
- .device
- looseString
- errBodyWriter
- .Router
- export/markdown.go
- Before/after pairs
- .sendFeedback
- jsconfig.json
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- WeaponKind
- ASIError
- Book
- .Entry
- ListSharedWithRow
- sourceUsedError
- .deleteUnused
- alert_templ.go
- tooltip_templ.go
- TestSourceSwitchKeepsEdits
- Props
- spellBoxes
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- ClassInfo
- KEY
- MAX_STEP_BYTES
- MAX_STEPS
- MSG
- ROUTES
- SAFE_SCHEMES
- WIKILINK
- countWriter
- IMPORT_FILTER_ALL
- IMPORT_FILTER_CATEGORY
- IMPORT_FILTER_NONE
- IMPORT_FILTER_STATUS
- IMPORT_ROWS_FIELD
- IMPORT_STATUS_ANY
- IMPORT_STATUS_BROKEN
- IMPORT_STATUS_TICKED
- IMPORT_STATUS_UNTICKED

## God Nodes (most connected - your core abstractions)
1. `App` - 375 edges
2. `Path()` - 274 edges
3. `newTestEnv()` - 176 edges
4. `Sheet` - 134 edges
5. `Button()` - 124 edges
6. `CN()` - 114 edges
7. `SheetView` - 79 edges
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

## Communities (175 total, 53 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "Path"
Cohesion: 0.04
Nodes (116): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ImageCategory, ListBacklinksRow, Entity, Look, Content() (+108 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (105): keepAwake(), ValidKind(), CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewRow, UpsertCompendiumEntryParams, CompendiumSource, Encounter (+97 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "EquipmentOption"
Cohesion: 0.40
Nodes (5): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), EquipmentChoice

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (115): Slot, Signed(), Button(), Variant, Card(), Content(), Description(), Header() (+107 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (39): GetImageRow, ListImagesRow, Attachment, Entity, EntityKind, imageItem(), imageSource(), imageVersion() (+31 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (123): Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight() (+115 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (33): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftNamed, ftRace, ftSpell (+25 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (54): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+46 more)

### Community 13 - "Input"
Cohesion: 0.06
Nodes (87): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+79 more)

### Community 14 - "redirect.go"
Cohesion: 0.27
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.04
Nodes (60): ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+52 more)

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
Nodes (25): ClearCookie(), SetCookie(), writeCookie(), entityTopics(), collectHeld(), formInt(), formIntQuery(), isHTMX() (+17 more)

### Community 21 - "WizardView"
Cohesion: 0.09
Nodes (45): attrString(), generateSVG(), getIconContent(), Props, Icon(), Label(), initialValue(), Item() (+37 more)

### Community 22 - "wikilink.js"
Cohesion: 0.14
Nodes (17): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+9 more)

### Community 23 - "Condition"
Cohesion: 0.05
Nodes (45): Combatant, Condition, CreateConditionParams, ListCombatantsRow, EffectEnd, combatantRow(), rollFor(), stepTurn() (+37 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (142): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+134 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.10
Nodes (19): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+11 more)

### Community 28 - "who"
Cohesion: 0.04
Nodes (31): Site, ftsQuery(), Search(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL() (+23 more)

### Community 29 - "select_templ.go"
Cohesion: 0.06
Nodes (52): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+44 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Slugify"
Cohesion: 0.06
Nodes (45): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), fiveToolsPage(), AbilityKeys() (+37 more)

### Community 34 - "monsterBody"
Cohesion: 0.15
Nodes (24): classBadges(), TestRowFigure(), crText(), levelsText(), parenthesized(), rowFigure(), signedMap(), speeds() (+16 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Values Go and JS share"
Cohesion: 0.20
Nodes (10): Gotchas, internal/web, Values Go and JS share, TestPagesRender(), TestNoRouteLiterals(), parseSharedJS(), TestSharedImportsExist(), TestSharedJS() (+2 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.12
Nodes (26): block, ClassFeature, ClassData, dnd5eFeatures(), TestFiveToolsClassLevelUp(), cellAt(), classTable(), featuresUnder() (+18 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Character"
Cohesion: 0.07
Nodes (24): CoreTraits(), dataFor(), DecodeData(), Character, Character, Campaign(), CharacterFileName(), fileSlug() (+16 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.10
Nodes (37): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+29 more)

### Community 42 - "Queries"
Cohesion: 0.05
Nodes (57): HomebrewOwner, Source, sourceFormat, TestHomebrew(), TestStoreSourceRollsBackOnAfterSave(), Decode(), SpellData, bindFiveToolsImages() (+49 more)

### Community 43 - "viewerCtx"
Cohesion: 0.22
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (22): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+14 more)

### Community 47 - "Env"
Cohesion: 0.13
Nodes (13): featureSwitch, ftRaceAbility(), featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch() (+5 more)

### Community 48 - "IfElse"
Cohesion: 0.05
Nodes (92): Props, CountHabitatsRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, Habitat, UI conventions, Badge() (+84 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (131): Defense, TestIsLoopback(), TestDecode(), TestCoreTraits(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+123 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Sheet"
Cohesion: 0.07
Nodes (21): TestCustomMods(), TestResistances(), TestExplainSums(), ArmorEffect(), Sum(), TestDualWield(), TestGearCommands(), classAbility() (+13 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Session"
Cohesion: 0.09
Nodes (20): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+12 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "popover_templ.go"
Cohesion: 0.13
Nodes (22): Field, Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition() (+14 more)

### Community 59 - "Part"
Cohesion: 0.12
Nodes (8): Part, Skill, TestModifier(), skillByKey(), Weapon, Modifier(), ProficiencyAt(), attackTargets()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (38): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+30 more)

### Community 62 - ".data"
Cohesion: 0.11
Nodes (25): ftMonster, ftSpellcasting, habitatsFrom(), FiveToolsImages(), ftCell(), ftFeatIncrease(), ftTable(), TestFiveToolsCell() (+17 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.11
Nodes (34): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList() (+26 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "State"
Cohesion: 0.10
Nodes (13): event, Handler, message, patchMessage, stateHandler, viewPatch, Topic, TopicKind (+5 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - "NameDesc"
Cohesion: 0.13
Nodes (26): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+18 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Client"
Cohesion: 0.13
Nodes (4): countingHandler, fakeHandler, Client, isDevice()

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
Cohesion: 0.11
Nodes (12): attachedPicture(), ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, nodeActions(), openWeb(), randomID() (+4 more)

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - ".q"
Cohesion: 0.10
Nodes (26): testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+18 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "NewClient"
Cohesion: 0.22
Nodes (19): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+11 more)

### Community 96 - "fields.js"
Cohesion: 0.21
Nodes (13): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+5 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - ".compendiumDir"
Cohesion: 0.17
Nodes (9): FiveToolsBooks(), ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), FiveToolsView, NewFiveToolsView(), FiveToolsBook (+1 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Hub"
Cohesion: 0.23
Nodes (3): ClientInfo, goneDevice, Hub

### Community 103 - "nullID"
Cohesion: 0.11
Nodes (37): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), acBonusForm(), effectFix, newEffectFix(), notes() (+29 more)

### Community 105 - "Store"
Cohesion: 0.07
Nodes (49): migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+41 more)

### Community 106 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (18): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftParts(), ftRecords() (+10 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 109 - "BlockEditor"
Cohesion: 0.24
Nodes (4): BlockEditor, json(), links(), sync()

### Community 111 - ".LevelUp"
Cohesion: 0.11
Nodes (13): LevelUpChoice, TestArmorGrants(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+5 more)

### Community 112 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 116 - ".features"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 117 - "time.Time"
Cohesion: 0.06
Nodes (21): Identity, FromContext(), Handlers, browserName(), firstLine(), mdLinkText(), TestUserAgent(), userAgentSummary() (+13 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - ".device"
Cohesion: 0.38
Nodes (5): testEnv, TestRequestLog(), TestRequestLogCharacterRemoved(), TestRequestLogReusedID(), TestRequestSent()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - ".Router"
Cohesion: 0.25
Nodes (4): IsLoopback(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 127 - "export/markdown.go"
Cohesion: 0.05
Nodes (49): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+41 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery() (+16 more)

### Community 137 - "WeaponKind"
Cohesion: 0.40
Nodes (5): Completions(), parseMastery(), TestCompletions(), TestWeaponKind(), WeaponKind()

### Community 139 - "Book"
Cohesion: 0.12
Nodes (22): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+14 more)

### Community 140 - ".Entry"
Cohesion: 0.10
Nodes (26): MonsterAction, addEntry(), setStat(), MonsterData, TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor(), InputForEntry() (+18 more)

### Community 143 - ".deleteUnused"
Cohesion: 0.43
Nodes (3): ListAttachmentFilesRow, fresh(), inCampaign()

### Community 144 - "alert_templ.go"
Cohesion: 0.11
Nodes (24): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+16 more)

### Community 145 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 146 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 147 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 148 - "spellBoxes"
Cohesion: 0.67
Nodes (4): spellBoxes(), TestTaughtPicked(), taughtPicked(), taughtRules()

### Community 157 - "ClassInfo"
Cohesion: 0.08
Nodes (24): TestSlotMaxima(), ClassInfo, LevelUpPlan, CasterLevel(), SlotMaxima(), Spellbook(), SpellColumn(), SpellTotals() (+16 more)

## Knowledge Gaps
- **40 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+35 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 455 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **53 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `database/sql.NullInt64`, `Book`, `Command`, `.deleteUnused`, `Viewer`, `Condition`, `cert.go`, `who`, `ClassInfo`, `Slugify`, `Character`, `Queries`, `viewerCtx`, `testing.T`, `FiveToolsJob`, `State`, `.q`, `.compendiumDir`, `Hub`, `Store`, `Pending`, `time.Time`, `.Router`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _40 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13919915700737617 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Path`, `context.Context`, `.sendFeedback`, `SheetView`, `database/sql.NullInt64`, `Input`, `.deleteUnused`, `Viewer`, `alert_templ.go`, `TestSourceSwitchKeepsEdits`, `App`, `WizardView`, `who`, `select_templ.go`, `Values Go and JS share`, `Character`, `tabs_templ.go`, `Queries`, `IfElse`, `testing.T`, `popover_templ.go`, `Player`, `CombatantView`, `ParseWiki`, `.q`, `nullID`, `deathFix`, `sourcesEnv`, `.device`, `export/markdown.go`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **Should `Path` be split into smaller, more focused modules?**
  _Cohesion score 0.04011728641876583 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `Values Go and JS share` to `Path`, `testing.T`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._