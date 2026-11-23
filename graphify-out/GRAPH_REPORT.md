# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 375 files · ~648,754 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4739 nodes · 24511 edges · 148 communities (114 shown, 34 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2461 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `053f7a07`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- ParseSheet
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
- CN
- Import
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- IfElse
- wikilink.js
- names
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- Store
- github.com/a-h/templ.Component
- app.js
- tabbable.js
- useDismiss
- NewClient
- github.com/a-h/templ.Component
- command.js
- dialog.js
- Open
- ParseWikidot
- tooltip.js
- CompendiumKind
- Sheet
- time.Time
- Item
- composite.js
- parseHP
- CustomMod
- Env
- Path
- popover_templ.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- .attachFile
- shadcn-templ Config
- Session
- golangci-lint Config
- ImportView
- Part
- radiogroup.js
- Player
- DownloadFiveTools
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- State
- use_transition_status.js
- playwright
- TestLevelUpPushesTracker
- Handlers
- checkbox.js
- ParseAidedd
- switch.js
- ftSpell
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
- Weapon
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- WeaponKind
- mounted
- .ListItemsByCategory
- nullID
- profFocus
- .ListSubracesOf
- redirect.go
- ftBook
- BlockEditor
- .LevelUp
- capReader
- .features
- database/sql.NullInt64
- build.sh
- TestSourceSwitchKeepsEdits
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
- sourceUsedError
- Pending
- rosterCard
- SearchReadableBlocksRow
- Parse
- CharacterFile
- spellCounts

## God Nodes (most connected - your core abstractions)
1. `App` - 375 edges
2. `Path()` - 271 edges
3. `newTestEnv()` - 174 edges
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

## Communities (148 total, 34 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "ParseSheet"
Cohesion: 0.10
Nodes (41): ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeatRaisesCon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+33 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (89): keepAwake(), TestHomebrew(), ftsQuery(), Search(), Encounter, CreateCombatantParams, EncounterStatus, EntityKind (+81 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.20
Nodes (18): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+10 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.09
Nodes (42): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+34 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (123): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+115 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "database/sql.NullInt64"
Cohesion: 0.04
Nodes (61): Combatant, Condition, CreateConditionParams, ListCombatantsRow, Character, EffectEnd, combatantRow(), passTurn() (+53 more)

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (72): activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestDeleteActiveCombatant(), TestDeleteActiveCombatantLast(), TestDeleteActiveCombatantSkipsOut(), TestDeleteCharacterActiveCombatant(), TestEncounterTransitions() (+64 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (49): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+41 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.09
Nodes (41): CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet(), parseSpell() (+33 more)

### Community 13 - "CN"
Cohesion: 0.05
Nodes (74): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ContentProps, ctxKey (+66 more)

### Community 14 - "Import"
Cohesion: 0.10
Nodes (31): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+23 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.04
Nodes (58): ListCharacterNamesRow, Note, GetBlockRow, NoteKind, Visibility, Fields, ParseFields(), VisibleEntity() (+50 more)

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
Nodes (39): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Campaign, Entity, CanChangeSubclass() (+31 more)

### Community 21 - "IfElse"
Cohesion: 0.05
Nodes (91): Props, Field, UI conventions, Checkbox(), initialChecked(), Label(), initialValue(), Item() (+83 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (125): Props, Variant, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps (+117 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.13
Nodes (14): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+6 more)

### Community 28 - "Store"
Cohesion: 0.06
Nodes (30): HomebrewOwner, Site, Source, sourceFormat, FiveToolsBooks(), fiveToolsPage(), formatOf(), Report (+22 more)

### Community 29 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (93): Coins, Defenses, Combobox(), Content(), Item(), List(), Description(), Field() (+85 more)

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
Cohesion: 0.15
Nodes (20): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+12 more)

### Community 34 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (114): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, SearchCompendiumRow, CompendiumSource, Habitat (+106 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Open"
Cohesion: 0.07
Nodes (37): fixture, tokens, migration, newFixture(), serve(), TestMiddlewareIdentity(), TestOnlyHashIsStored(), TestRequireAdmin() (+29 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.09
Nodes (32): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, dnd5eFeatures(), TestFiveToolsCoreTraitsShapes(), withCoreTraits() (+24 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.04
Nodes (42): dataFor(), DecodeData(), ValidKind(), CountCompendiumRow, UpsertCompendiumEntryParams, Character, CompendiumEntry, CompendiumKind (+34 more)

### Community 41 - "Sheet"
Cohesion: 0.04
Nodes (108): LevelUpChoice, internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice() (+100 more)

### Community 42 - "time.Time"
Cohesion: 0.17
Nodes (8): FiveToolsView, NewFiveToolsView(), BlockStep, LevelUp, FiveToolsBook, FiveToolsGroup, reqEntry, reqLog

### Community 43 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "parseHP"
Cohesion: 0.26
Nodes (9): Amount, CritDamage(), Damage(), parseDamage(), parseHP(), Set(), parseAmount(), TestApplyHPDeath() (+1 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (22): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+14 more)

### Community 47 - "Env"
Cohesion: 0.24
Nodes (6): featureSwitch, Env, featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 48 - "Path"
Cohesion: 0.04
Nodes (108): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListImageCharactersRow (+100 more)

### Community 49 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (100): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestIsLoopback(), TestDecode(), TestSkillChoice(), RemoveSourceDir(), TestRemoveSourceDir() (+92 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - ".attachFile"
Cohesion: 0.24
Nodes (9): testEnv, TestAttachmentAccess(), TestAttachmentFileMissing(), TestAttachmentToImage(), TestCampaignFiles(), TestDeleteRemovesFiles(), TestFilesCleanup(), TestFileUsePrivacy() (+1 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Session"
Cohesion: 0.11
Nodes (11): IsLoopback(), Manager, RequireAdmin(), Session, Role, sessionKey(), New(), sameOriginOnly() (+3 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnOverview(), ImportReview(), importRows(), importStep(), keepSection(), columnLabel(), ImportCategory (+2 more)

### Community 59 - "Part"
Cohesion: 0.12
Nodes (6): Part, Skill, TestModifier(), skillByKey(), Modifier(), ProficiencyAt()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.11
Nodes (25): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views (+17 more)

### Community 62 - "DownloadFiveTools"
Cohesion: 0.10
Nodes (29): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+21 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.12
Nodes (33): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList() (+25 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "State"
Cohesion: 0.08
Nodes (15): countingHandler, event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch (+7 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "TestLevelUpPushesTracker"
Cohesion: 0.31
Nodes (8): TestLevelUpPushesTracker(), dial(), patchOf(), readUntil(), rosterRefetch(), TestDeviceSeenOnLeave(), TestLiveUpdatesAndKick(), wsMessage

### Community 72 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "ftSpell"
Cohesion: 0.50
Nodes (5): ftDistance, ftSpell, ftAmount(), ftDuration(), ftRange()

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
Cohesion: 0.33
Nodes (8): CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks(), WikiLinks(), WikiLink, WikiPart

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
Nodes (30): isCleared(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+22 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Weapon"
Cohesion: 0.15
Nodes (10): ItemData, TestCustomMods(), TestExplainSums(), Sum(), TestDualWield(), damageLabel(), Weapon, compendiumWeapon() (+2 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "nullID"
Cohesion: 0.13
Nodes (32): DeathSaves, nullID(), deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP() (+24 more)

### Community 106 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (17): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftPrerequisite(), ftRecords() (+9 more)

### Community 111 - ".LevelUp"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 116 - ".features"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 117 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (48): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, firstLine(), mdLinkText() (+40 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 127 - "export/markdown.go"
Cohesion: 0.05
Nodes (49): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+41 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 132 - "Hub"
Cohesion: 0.22
Nodes (4): ClientInfo, goneDevice, Hub, TestCloseSessionAndCharacter()

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 137 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 139 - "Book"
Cohesion: 0.12
Nodes (22): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), CleanDownloads(), downloadBook(), fetchBooks() (+14 more)

### Community 140 - ".Entry"
Cohesion: 0.08
Nodes (30): MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools() (+22 more)

### Community 143 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 144 - "rosterCard"
Cohesion: 0.29
Nodes (10): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), deviceLabel(), rosterHP() (+2 more)

### Community 155 - "Parse"
Cohesion: 0.05
Nodes (48): Defense, feature(), TestArmorTrainingGrantsAndRemoval(), TestCatchUpArmorTraining(), Heal(), Parse(), tokenize(), TestDeadRefusals() (+40 more)

### Community 156 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 157 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 419 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **34 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Hub`, `State`, `CompendiumKind`, `database/sql.NullInt64`, `Sheet`, `Book`, `Command`, `time.Time`, `Pending`, `Viewer`, `testing.T`, `database/sql.NullInt64`, `.q`, `Session`, `Store`, `FiveToolsJob`?**
  _High betweenness centrality (0.045) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14099350046425255 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `ParseSheet`, `context.Context`, `tabs_templ.go`, `SheetView`, `database/sql.NullInt64`, `newTestEnv`, `CN`, `Import`, `Viewer`, `rosterCard`, `App`, `IfElse`, `github.com/a-h/templ.Attributes`, `Parse`, `github.com/a-h/templ.Component`, `github.com/a-h/templ.Component`, `Open`, `Sheet`, `testing.T`, `.attachFile`, `ImportView`, `Player`, `CombatantView`, `TestLevelUpPushesTracker`, `ParseWiki`, `.q`, `nullID`, `database/sql.NullInt64`, `TestSourceSwitchKeepsEdits`, `export/markdown.go`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **Should `ParseSheet` be split into smaller, more focused modules?**
  _Cohesion score 0.09745293466223699 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `Import`, `wikilink.js`?**
  _High betweenness centrality (0.025) - this node is a cross-community bridge._