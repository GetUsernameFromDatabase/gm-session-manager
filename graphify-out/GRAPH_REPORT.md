# Graph Report - gm-session-manager  (2026-11-26)

## Corpus Check
- 375 files · ~666,640 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4888 nodes · 25475 edges · 174 communities (120 shown, 54 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2545 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `08fdd736`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- .fail
- context.Context
- vendorjs/main.go
- createProvider
- export/markdown.go
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
- time.Time
- ParseWikidot
- tooltip.js
- CompendiumEntry
- tabs_templ.go
- StarterSets
- Session
- composite.js
- database/sql.NullString
- CustomMod
- Store
- IfElse
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- .sendFeedback
- shadcn-templ Config
- .q
- golangci-lint Config
- deathFix
- Sheet
- radiogroup.js
- Player
- ParseEffects
- DownloadOpen5e
- sqlc Config
- .editEntity
- Browser check
- Claude Settings
- State
- use_transition_status.js
- playwright
- ftPrerequisite
- .sourceIssues
- checkbox.js
- ParseAidedd
- switch.js
- Layout
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
- storeSource
- mounted
- Client
- nullID
- profFocus
- cert.go
- TestSourceSwitchKeepsEdits
- ftBook
- Pending
- BlockEditor
- names
- joinNonEmpty
- capReader
- ListSharedWithRow
- sourcesEnv
- build.sh
- .levelUp
- looseString
- errBodyWriter
- CountSharedByKindRow
- ListEntrySourcesBySlugsParams
- Before/after pairs
- Values Go and JS share
- jsconfig.json
- open5e.go
- testing.M
- ftItemData
- perm/CLAUDE.md
- toast_templ.go
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- ImagePage
- viewerCtx
- rosterCard
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
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
- Queries
- Handlers
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 384 edges
2. `Path()` - 278 edges
3. `newTestEnv()` - 178 edges
4. `Sheet` - 134 edges
5. `Button()` - 126 edges
6. `CN()` - 114 edges
7. `SheetView` - 79 edges
8. `IfElse()` - 67 edges
9. `Command` - 60 edges
10. `Queries` - 57 edges

## Surprising Connections (you probably didn't know these)
- `Notes` --references--> `linkClick()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/assets/js/wikilink.js
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (174 total, 54 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (56): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+48 more)

### Community 1 - ".fail"
Cohesion: 0.08
Nodes (11): Character, NewCharacterFile(), CanChangeSubclass(), CanCharacter(), TestCanChangeSubclass(), nextStep(), wizardPath(), Named (+3 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (116): keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Encounter, Note, SearchReadableBlocksRow (+108 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.05
Nodes (49): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+41 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (116): Coins, Slot, Signed(), Button(), Action(), Card(), Content(), Description() (+108 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.08
Nodes (49): Import(), TestUpgradeSheet(), ParseSheet(), TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestMonsterImageLink(), TestBuilderCompendium(), TestChangeSubclass() (+41 more)

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (98): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+90 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (44): ftClass, ftCopyRef, ftFeat, ftModOp, ftNamed, ftRace, ftStartingProfs, ftSubclass (+36 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (54): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+46 more)

### Community 13 - "Input"
Cohesion: 0.07
Nodes (86): Field, Chips(), ChipsInput(), Combobox(), Content(), Item(), List(), state() (+78 more)

### Community 14 - "redirect.go"
Cohesion: 0.25
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (52): ListCharacterNamesRow, GetBlockRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanDeleteCharacter() (+44 more)

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
Nodes (45): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), ftsQuery(), Search(), Campaign (+37 more)

### Community 21 - "WizardView"
Cohesion: 0.10
Nodes (38): Label(), initialValue(), Item(), RadioGroup(), state(), abilitySelect(), levelUpForm(), spellChecklist() (+30 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (16): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+8 more)

### Community 23 - "Condition"
Cohesion: 0.05
Nodes (47): Combatant, Condition, CreateConditionParams, ListCombatantsRow, EffectEnd, combatantRow(), passTurn(), rollFor() (+39 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (170): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+162 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "popover_templ.go"
Cohesion: 0.12
Nodes (29): UI conventions, Content(), Description(), Header(), Popover(), state(), Title(), Trigger() (+21 more)

### Community 28 - "CombatantView"
Cohesion: 0.10
Nodes (38): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+30 more)

### Community 29 - "formSelect"
Cohesion: 0.05
Nodes (60): Props, ImageCategory, Defenses, Checkbox(), initialChecked(), Content(), Group(), initialValue() (+52 more)

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
Cohesion: 0.07
Nodes (47): internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice(), TestOptions() (+39 more)

### Community 34 - "mirror"
Cohesion: 0.36
Nodes (4): bookWriter, mirror, downloadBook(), filterDoc()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "time.Time"
Cohesion: 0.06
Nodes (18): ListAttachmentFilesRow, ListImagesRow, Attachment, Handlers, browserName(), firstLine(), mdLinkText(), fresh() (+10 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.11
Nodes (27): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt() (+19 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumEntry"
Cohesion: 0.09
Nodes (17): DecodeData(), CompendiumEntry, Character, checkSkillPick(), classSkillPick(), postedSkills(), legacyBackground(), backgroundRaises() (+9 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.09
Nodes (42): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+34 more)

### Community 42 - "StarterSets"
Cohesion: 0.19
Nodes (9): FiveToolsProgress(), importBar(), jobNotice(), starterSet(), StarterSets(), FiveToolsJob, starterBase(), starterLive() (+1 more)

### Community 43 - "Session"
Cohesion: 0.06
Nodes (29): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+21 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "database/sql.NullString"
Cohesion: 0.06
Nodes (23): targetColumns(), GetImageRow, LinkTarget, saveFieldLinks(), linkTarget(), ensureCategory(), imageItem(), imageSource() (+15 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (27): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), AddFeat(), SaveFeature() (+19 more)

### Community 47 - "Store"
Cohesion: 0.06
Nodes (52): linkTarget, migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+44 more)

### Community 48 - "IfElse"
Cohesion: 0.04
Nodes (108): dataFor(), CountCompendiumRow, CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewByOwnerRow, ListHomebrewRow (+100 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (113): Defense, TestDecode(), feature(), RemoveSourceDir(), TestRemoveSourceDir(), TestPageClientRedirect(), TestArmorTrainingGrantsAndRemoval(), TestArmorUntrained() (+105 more)

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

### Community 56 - ".q"
Cohesion: 0.10
Nodes (27): testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+19 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 59 - "Sheet"
Cohesion: 0.05
Nodes (33): Part, Skill, Completions(), findItem(), parseItem(), parseMastery(), TestCustomMods(), TestModifier() (+25 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (37): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+29 more)

### Community 62 - "ParseEffects"
Cohesion: 0.06
Nodes (22): LevelUpChoice, armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorGrants() (+14 more)

### Community 63 - "DownloadOpen5e"
Cohesion: 0.12
Nodes (23): fetchJSON(), replaceDir(), TestMirrorGetMissingFile(), DownloadOpen5e(), open5eFiles(), serveOpen5e(), TestDownloadOpen5e(), TestDownloadOpen5eBadName() (+15 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - ".editEntity"
Cohesion: 0.11
Nodes (10): Entity, Fields, ParseFields(), entityTopics(), linkEntityLook(), formInt(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams (+2 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "State"
Cohesion: 0.10
Nodes (13): event, Handler, patchMessage, stateHandler, viewPatch, Topic, TopicKind, Diff() (+5 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.22
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - ".sourceIssues"
Cohesion: 0.31
Nodes (8): pinnedRefs(), sourceIssue(), featKey(), pinFeat(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Layout"
Cohesion: 0.06
Nodes (84): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Content() (+76 more)

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
Nodes (12): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), nodeActions(), openWeb() (+4 more)

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - "Book"
Cohesion: 0.11
Nodes (26): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+18 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "NewClient"
Cohesion: 0.12
Nodes (23): ClientInfo, goneDevice, Hub, isDevice(), New(), NewClient(), TestApprovedSurvivesDisconnect(), TestCatchUp() (+15 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "storeSource"
Cohesion: 0.09
Nodes (24): HomebrewOwner, Site, Source, sourceFormat, TestHomebrew(), TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), formatOf() (+16 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Client"
Cohesion: 0.13
Nodes (5): countingHandler, fakeHandler, message, Client, drain()

### Community 103 - "nullID"
Cohesion: 0.23
Nodes (19): acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectModRows(), TestEffectModsReadOnly() (+11 more)

### Community 105 - "cert.go"
Cohesion: 0.11
Nodes (18): main(), run(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate(), LoadOrCreateCert() (+10 more)

### Community 106 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 107 - "ftBook"
Cohesion: 0.19
Nodes (15): ftBook, ftRec, each(), ftCopyTargets(), ftDecode(), ftRecords(), TestFiveToolsCopy(), ftMod() (+7 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 109 - "BlockEditor"
Cohesion: 0.24
Nodes (3): attachedPicture(), BlockEditor, randomID()

### Community 111 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 112 - "joinNonEmpty"
Cohesion: 0.11
Nodes (21): ftDistance, ftEntry, ftMonster, ftSpell, ftSpellcasting, habitatsFrom(), FiveToolsImages(), ftAmount() (+13 more)

### Community 116 - "ListSharedWithRow"
Cohesion: 0.27
Nodes (4): ListSharedWithRow, ListSharedOrderParams, ListSharedOrderRow, ListSharedWithParams

### Community 117 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - ".levelUp"
Cohesion: 0.06
Nodes (37): TextData, ListSubclassesOfRow, FreeCast(), SwapFeatSpell(), ClassFeature, LevelUpPlan, Spell, Spellbook() (+29 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "Values Go and JS share"
Cohesion: 0.18
Nodes (11): Gotchas, internal/web, Notes, Values Go and JS share, TestPagesRender(), TestNoRouteLiterals(), parseSharedJS(), TestSharedImportsExist() (+3 more)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 135 - "ftItemData"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 137 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 139 - "DownloadFiveTools"
Cohesion: 0.10
Nodes (29): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+21 more)

### Community 140 - ".Entry"
Cohesion: 0.09
Nodes (27): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor() (+19 more)

### Community 143 - "ImagePage"
Cohesion: 0.08
Nodes (37): Props, Size, ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, Variant, Variants() (+29 more)

### Community 144 - "viewerCtx"
Cohesion: 0.22
Nodes (7): notFoundIsNil(), parseLiveKey(), Roster, liveKey, liveKind, liveView, viewerCtx

### Community 145 - "rosterCard"
Cohesion: 0.44
Nodes (7): deviceRow(), rosterCard(), devicesWhere(), agoTime(), deviceLabel(), Device, RosterCharacter

### Community 175 - "Queries"
Cohesion: 0.07
Nodes (39): featureSwitch, Decode(), SpellData, Queries, Env, ClassInfo, ClassLevel, featureNames() (+31 more)

### Community 176 - "Handlers"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 456 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **54 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.fail`, `context.Context`, `Command`, `viewerCtx`, `perm.go`, `Condition`, `Slugify`, `time.Time`, `CompendiumEntry`, `StarterSets`, `Session`, `database/sql.NullString`, `Queries`, `IfElse`, `testing.T`, `.q`, `Sheet`, `.editEntity`, `State`, `.sourceIssues`, `Book`, `NewClient`, `cert.go`, `Pending`, `.levelUp`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14035010620376473 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `newTestEnv` to `compendium_templ.go`, `.fail`, `Values Go and JS share`, `export/markdown.go`, `SheetView`, `ParseSheet`, `Input`, `ImagePage`, `perm.go`, `rosterCard`, `App`, `WizardView`, `Condition`, `github.com/a-h/templ.Attributes`, `popover_templ.go`, `CombatantView`, `formSelect`, `time.Time`, `tabs_templ.go`, `database/sql.NullString`, `Store`, `IfElse`, `testing.T`, `.sendFeedback`, `.q`, `deathFix`, `Player`, `DownloadOpen5e`, `.editEntity`, `Layout`, `ParseWiki`, `nullID`, `TestSourceSwitchKeepsEdits`, `sourcesEnv`, `.levelUp`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **Should `.fail` be split into smaller, more focused modules?**
  _Cohesion score 0.07910014513788098 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `Values Go and JS share` to `newTestEnv`, `testing.T`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._