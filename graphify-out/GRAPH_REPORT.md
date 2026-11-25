# Graph Report - gm-session-manager  (2026-11-25)

## Corpus Check
- 375 files · ~664,505 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4875 nodes · 25408 edges · 182 communities (123 shown, 59 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2534 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `6ffa108c`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- combobox_templ.go
- context.Context
- vendorjs/main.go
- createProvider
- buildReview
- SheetView
- combobox.js
- ParseSheet
- newTestEnv
- encoding/json.RawMessage
- select.js
- Command
- Input
- redirect.go
- floating_ui_dom.js
- perm.go
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
- popover_templ.go
- CombatantView
- formSelect
- app.js
- tabbable.js
- useDismiss
- Slugify
- mirror
- command.js
- dialog.js
- database/sql.NullInt64
- ParseWikidot
- tooltip.js
- CompendiumKind
- tabs_templ.go
- StarterSets
- Open
- composite.js
- ParseArmorTraining
- CustomMod
- Import
- github.com/a-h/templ.Component
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- .sendFeedback
- shadcn-templ Config
- Item
- golangci-lint Config
- deathFix
- Weapon
- radiogroup.js
- Player
- ParseEffects
- DownloadOpen5e
- sqlc Config
- WeaponKind
- Browser check
- Claude Settings
- State
- use_transition_status.js
- playwright
- ftPrerequisite
- .Router
- checkbox.js
- ParseAidedd
- switch.js
- blockMenu
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- Book
- useTypeahead
- NewClient
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- who
- mounted
- Client
- nullID
- profFocus
- cert.go
- .source
- ftBook
- Pending
- BlockEditor
- names
- InputForEntry
- capReader
- CharacterFile
- sourcesEnv
- build.sh
- spellNeed
- looseString
- errBodyWriter
- Hub
- export/markdown.go
- Before/after pairs
- TestPagesRender
- jsconfig.json
- open5e.go
- testing.M
- Weapon
- perm/CLAUDE.md
- toast_templ.go
- ASIError
- DownloadFiveTools
- .Entry
- SearchReadableBlocksRow
- sourceUsedError
- .ListBlockStepsSince
- Device
- .ListImageLookAccess
- .ListItemsByCategory
- .ListSpellsForClass
- .ListSubracesOf
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- fakeHandler
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
- Env
- Handlers
- .coreTraits
- ftDuration
- parseMod
- condMods
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 380 edges
2. `Path()` - 279 edges
3. `newTestEnv()` - 178 edges
4. `Sheet` - 134 edges
5. `Button()` - 125 edges
6. `CN()` - 114 edges
7. `SheetView` - 79 edges
8. `IfElse()` - 67 edges
9. `Command` - 60 edges
10. `Queries` - 57 edges

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

## Communities (182 total, 59 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (56): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+48 more)

### Community 1 - "combobox_templ.go"
Cohesion: 0.04
Nodes (66): Props, Size, Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey (+58 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (70): keepAwake(), CompendiumSource, Encounter, CreateCombatantParams, EncounterStatus, EntityKind, Topic, TopicKind (+62 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (125): Slot, Signed(), Button(), Variant, Action(), Card(), Content(), Description() (+117 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.11
Nodes (37): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeatRaisesCon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+29 more)

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (94): TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestConditionExpiresOnTurn(), TestDeleteActiveCombatant(), TestDeleteActiveCombatantLast() (+86 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (47): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+39 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (49): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+41 more)

### Community 13 - "Input"
Cohesion: 0.07
Nodes (80): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+72 more)

### Community 14 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
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
Nodes (28): IsLoopback(), ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Campaign, browserName() (+20 more)

### Community 21 - "WizardView"
Cohesion: 0.06
Nodes (63): Props, dnd5eFeatures(), TestFiveToolsClassLevelUp(), ClassFeature, LevelUpPlan, Checkbox(), initialChecked(), Label() (+55 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (16): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+8 more)

### Community 23 - "Condition"
Cohesion: 0.05
Nodes (57): linkTarget, FiveToolsImages(), bindFiveToolsImages(), linkTargets(), Queries, Combatant, Condition, CreateConditionParams (+49 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (157): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+149 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "popover_templ.go"
Cohesion: 0.15
Nodes (21): Description(), Header(), state(), Title(), Trigger(), triggerID(), effectBits(), combatCondition() (+13 more)

### Community 28 - "CombatantView"
Cohesion: 0.09
Nodes (40): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+32 more)

### Community 29 - "formSelect"
Cohesion: 0.07
Nodes (39): ImageCategory, Defenses, Title(), damageChoice(), firstValue(), missingDamageItems(), modTargetLabel(), usesItems() (+31 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Slugify"
Cohesion: 0.04
Nodes (98): internal/compendium, Cantrips(), CoreTraits(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice() (+90 more)

### Community 34 - "mirror"
Cohesion: 0.36
Nodes (4): bookWriter, mirror, downloadBook(), filterDoc()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "database/sql.NullInt64"
Cohesion: 0.02
Nodes (68): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Note, NoteKind, firstLine(), folderName() (+60 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.14
Nodes (23): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (47): dataFor(), DecodeData(), ValidKind(), CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewRow, Character, Character (+39 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.10
Nodes (37): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+29 more)

### Community 42 - "StarterSets"
Cohesion: 0.22
Nodes (7): importBar(), starterSet(), StarterSets(), FiveToolsJob, starterBase(), starterLive(), fiveToolsJob

### Community 43 - "Open"
Cohesion: 0.04
Nodes (59): fixture, tokens, main(), run(), migration, constantTimeEqual(), credentialsMatch(), isCleared() (+51 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 46 - "CustomMod"
Cohesion: 0.11
Nodes (15): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+7 more)

### Community 47 - "Import"
Cohesion: 0.11
Nodes (35): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+27 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.03
Nodes (167): CountHabitatsRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, ListSharedWithRow, SearchCompendiumRow, Habitat, Field (+159 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (128): Defense, TestIsLoopback(), TestDecode(), TestCoreTraits(), TestSkillChoice(), feature(), RemoveSourceDir(), TestRemoveSourceDir() (+120 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Item"
Cohesion: 0.30
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 59 - "Weapon"
Cohesion: 0.12
Nodes (6): Part, Skill, TestModifier(), skillByKey(), Modifier(), ProficiencyAt()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.07
Nodes (39): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+31 more)

### Community 62 - "ParseEffects"
Cohesion: 0.09
Nodes (16): LevelUpChoice, TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+8 more)

### Community 63 - "DownloadOpen5e"
Cohesion: 0.13
Nodes (22): fetchJSON(), replaceDir(), DownloadOpen5e(), open5eFiles(), serveOpen5e(), TestDownloadOpen5e(), TestDownloadOpen5eBadName(), TestDownloadOpen5eFailureKeepsOld() (+14 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "WeaponKind"
Cohesion: 0.40
Nodes (5): Completions(), parseMastery(), TestCompletions(), TestWeaponKind(), WeaponKind()

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "State"
Cohesion: 0.14
Nodes (7): Handler, patchMessage, stateHandler, viewPatch, Diff(), State, TestDiff()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 72 - ".Router"
Cohesion: 0.33
Nodes (3): RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "blockMenu"
Cohesion: 0.08
Nodes (33): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListBacklinksRow, Entity, Look, initials(), lookControls() (+25 more)

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
Cohesion: 0.10
Nodes (13): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), nodeActions(), openWeb() (+5 more)

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - "Book"
Cohesion: 0.12
Nodes (24): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), Book, ImportedFiveTools(), normalizeBooks(), sortBooks() (+16 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "NewClient"
Cohesion: 0.16
Nodes (20): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+12 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "who"
Cohesion: 0.03
Nodes (50): HomebrewOwner, Site, Source, sourceFormat, TestHomebrew(), FiveToolsBooks(), fiveToolsPage(), ftsQuery() (+42 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "nullID"
Cohesion: 0.10
Nodes (40): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), TestMonsterAtZeroHPUnaffected(), acBonusForm(), effectFix, newEffectFix() (+32 more)

### Community 105 - "cert.go"
Cohesion: 0.15
Nodes (10): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+2 more)

### Community 106 - ".source"
Cohesion: 0.12
Nodes (17): ParsePageURL(), TestParsePageURL(), UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove() (+9 more)

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (17): ftBook, ftCopyRef, ftRec, each(), ftCopyTargets(), ftDecode(), ftRecords(), ftRequires() (+9 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 109 - "BlockEditor"
Cohesion: 0.24
Nodes (3): attachedPicture(), BlockEditor, randomID()

### Community 111 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 112 - "InputForEntry"
Cohesion: 0.25
Nodes (7): habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 116 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 117 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "spellNeed"
Cohesion: 0.20
Nodes (9): Spellbook(), SpellColumn(), SpellTotals(), TestSpellbook(), spellNeed(), TestSpellNeed(), TestSpellCounts(), sortedColumns() (+1 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Hub"
Cohesion: 0.23
Nodes (5): ClientInfo, event, goneDevice, Hub, TestCloseSessionAndCharacter()

### Community 127 - "export/markdown.go"
Cohesion: 0.15
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "TestPagesRender"
Cohesion: 0.25
Nodes (7): EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Gotchas, internal/web, TestPagesRender()

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 135 - "Weapon"
Cohesion: 0.13
Nodes (10): ItemData, TestCustomMods(), TestExplainSums(), Sum(), TestDualWield(), damageLabel(), Weapon, compendiumWeapon() (+2 more)

### Community 137 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 139 - "DownloadFiveTools"
Cohesion: 0.10
Nodes (29): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+21 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 144 - "Device"
Cohesion: 0.33
Nodes (10): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), deviceLabel(), rosterHP() (+2 more)

### Community 175 - "Env"
Cohesion: 0.19
Nodes (9): featureSwitch, Env, ArmorEffect(), WornArmor, featureNames(), SubclassRef, planSwitch(), sameSlug() (+1 more)

### Community 176 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 177 - ".coreTraits"
Cohesion: 0.33
Nodes (5): ftProfTexts(), ftSkillNames(), TestFiveToolsCoreTraitsShapes(), withCoreTraits(), ArmorKindNamed()

### Community 178 - "ftDuration"
Cohesion: 0.50
Nodes (5): ftDistance, ftAmount(), ftCastingTime(), ftDuration(), ftRange()

### Community 179 - "parseMod"
Cohesion: 0.50
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 180 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 454 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **59 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Slugify`, `context.Context`, `who`, `database/sql.NullInt64`, `CompendiumKind`, `newTestEnv`, `.Router`, `Open`, `Command`, `Pending`, `StarterSets`, `perm.go`, `Book`, `Condition`, `Hub`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14165782962177448 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `newTestEnv` to `compendium_templ.go`, `combobox_templ.go`, `TestPagesRender`, `SheetView`, `ParseSheet`, `Input`, `perm.go`, `Device`, `App`, `WizardView`, `Condition`, `github.com/a-h/templ.Attributes`, `popover_templ.go`, `CombatantView`, `formSelect`, `Slugify`, `database/sql.NullInt64`, `CompendiumKind`, `tabs_templ.go`, `StarterSets`, `Import`, `github.com/a-h/templ.Component`, `testing.T`, `.sendFeedback`, `deathFix`, `Player`, `DownloadOpen5e`, `blockMenu`, `ParseWiki`, `who`, `nullID`, `.source`, `sourcesEnv`, `export/markdown.go`?**
  _High betweenness centrality (0.047) - this node is a cross-community bridge._
- **Should `combobox_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.043859649122807015 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `newTestEnv`, `TestPagesRender`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._