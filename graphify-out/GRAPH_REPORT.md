# Graph Report - gm-session-manager  (2026-11-27)

## Corpus Check
- 392 files · ~689,585 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 59 file(s) not represented in the graph (top: .templ 49, .css 4, (none) 3)

## Summary
- 5078 nodes · 26711 edges · 181 communities (125 shown, 56 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2639 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `7ff868d8`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CampaignSources
- context.Context
- vendorjs/main.go
- createProvider
- export/markdown.go
- SheetView
- combobox.js
- Path
- combobox_templ.go
- encoding/json.RawMessage
- select.js
- Command
- viewerCtx
- redirect.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- Input
- wikilink.js
- dialog_templ.go
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Sheet
- CombatantView
- Condition
- app.js
- tabbable.js
- useDismiss
- Slugify
- Layout
- command.js
- dialog.js
- CharacterFile
- ParseWikidot
- tooltip.js
- .wizardView
- tabs_templ.go
- WizardView
- Client
- composite.js
- time.Time
- CustomMod
- ParseSheet
- github.com/a-h/templ.Component
- Option
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Open
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- Queries
- ImagePage
- radiogroup.js
- Player
- Character
- .renderCompendium
- sqlc Config
- .sourceIssues
- Session
- Claude Settings
- starterEnv
- use_transition_status.js
- playwright
- ftPrerequisite
- Book
- checkbox.js
- ParseAidedd
- switch.js
- config.go
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- mirror
- useTypeahead
- .source
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- .levelUp
- mounted
- net/http/httptest.ResponseRecorder
- newEffectFix
- profFocus
- cert.go
- .Router
- ftBook
- .editEntity
- BlockEditor
- Hub
- popover_templ.go
- capReader
- buildReview
- Attachment
- build.sh
- Pending
- looseString
- errBodyWriter
- takehome.go
- storeSource
- Before/after pairs
- deathFix
- jsconfig.json
- open5e.go
- testing.M
- CompendiumKind
- TakeHomeFile
- names
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- takehome_test.go
- .deleteUnused
- sourcesEnv
- Campaign
- ftItemData
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- .importParam
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
- pause_windows.go
- importLooks
- Browser check
- player-app.sh
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 398 edges
2. `Path()` - 291 edges
3. `newTestEnv()` - 183 edges
4. `Sheet` - 134 edges
5. `Button()` - 130 edges
6. `CN()` - 114 edges
7. `SheetView` - 80 edges
8. `IfElse()` - 67 edges
9. `Queries` - 62 edges
10. `Command` - 60 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (181 total, 56 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (56): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+48 more)

### Community 1 - "CampaignSources"
Cohesion: 0.08
Nodes (35): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks(), BuilderPick (+27 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (140): keepAwake(), targetColumns(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, GetImageRow (+132 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.15
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (100): Coins, Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+92 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Path"
Cohesion: 0.04
Nodes (114): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), Gotchas, internal/web, activeFight() (+106 more)

### Community 9 - "combobox_templ.go"
Cohesion: 0.05
Nodes (85): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+77 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (48): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+40 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (56): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+48 more)

### Community 13 - "viewerCtx"
Cohesion: 0.23
Nodes (6): notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView, viewerCtx

### Community 14 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.04
Nodes (68): ListCharacterNamesRow, GetBlockRow, IsConditionSlug(), EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Visibility (+60 more)

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
Nodes (31): FromContext(), Campaign, LinkTarget, compendiumBack(), compendiumPath(), pathCampaign(), viewerCampaign(), downloadNames() (+23 more)

### Community 21 - "Input"
Cohesion: 0.04
Nodes (66): ImageCategory, Defenses, Title(), Content(), Group(), initialValue(), Item(), itemLabel() (+58 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.08
Nodes (55): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+47 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (161): Props, Props, Variant, Props, Size, ActionProps, ContentProps, DescriptionProps (+153 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Sheet"
Cohesion: 0.05
Nodes (30): Part, Skill, ItemData, Completions(), parseMastery(), TestCustomMods(), TestModifier(), TestResistances() (+22 more)

### Community 28 - "CombatantView"
Cohesion: 0.11
Nodes (36): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList() (+28 more)

### Community 29 - "Condition"
Cohesion: 0.04
Nodes (49): Combatant, Condition, Encounter, ListCombatantsRow, EncounterStatus, combatantRow(), passTurn(), rollFor() (+41 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Slugify"
Cohesion: 0.07
Nodes (43): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, Slugify() (+35 more)

### Community 34 - "Layout"
Cohesion: 0.09
Nodes (40): Scripts(), Separator(), Toast(), Toaster(), toastIcon(), toastListItem(), campaignLink(), PendingRequests() (+32 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "CharacterFile"
Cohesion: 0.22
Nodes (8): CharacterFile, parseCharacterFile(), userf(), checkSheetFile(), normalizeCharacterFile(), fieldError, pickError, userError

### Community 38 - "ParseWikidot"
Cohesion: 0.14
Nodes (23): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".wizardView"
Cohesion: 0.07
Nodes (32): EquipmentItem, internal/compendium, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), TestCoreTraits(), TestEquipmentOptions() (+24 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.10
Nodes (35): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+27 more)

### Community 42 - "WizardView"
Cohesion: 0.09
Nodes (52): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+44 more)

### Community 43 - "Client"
Cohesion: 0.08
Nodes (17): countingHandler, event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch (+9 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "time.Time"
Cohesion: 0.08
Nodes (16): folderName(), slugify(), TestFolderName(), devicesWhere(), deviceLabel(), rosterHP(), BlockStep, CreateLevelUpAtParams (+8 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (23): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), AddFeat(), SaveFeature() (+15 more)

### Community 47 - "ParseSheet"
Cohesion: 0.08
Nodes (51): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestMonsterImageLink(), TestBuilderCompendium() (+43 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (110): CountHabitatsRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, UI conventions, Badge(), Spinner(), Cell() (+102 more)

### Community 49 - "Option"
Cohesion: 0.08
Nodes (24): ListSubclassesOfRow, SpellColumn(), formErrors(), spellNeed(), TestSubclassPanels(), TestSpellNeed(), subclassIndex(), subclassOptions() (+16 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (116): Defense, TestDecode(), feature(), RemoveSourceDir(), TestRemoveSourceDir(), TestPageClientRedirect(), TestArmorTrainingGrantsAndRemoval(), TestArmorUntrained() (+108 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Open"
Cohesion: 0.07
Nodes (48): linkTarget, migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+40 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.07
Nodes (22): LevelUpChoice, armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorGrants() (+14 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Queries"
Cohesion: 0.08
Nodes (35): featureSwitch, Decode(), Queries, CompendiumEntry, Env, TestSlotMaxima(), ClassFeature, ClassInfo (+27 more)

### Community 59 - "ImagePage"
Cohesion: 0.12
Nodes (27): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, initials(), EntityDetail(), entityForm(), EntityPage() (+19 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (40): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+32 more)

### Community 62 - "Character"
Cohesion: 0.13
Nodes (11): Character, Character, NewCharacterFile(), checkSkillPick(), postedSkills(), backgroundRaises(), setupOnce(), buildTakeHome() (+3 more)

### Community 63 - ".renderCompendium"
Cohesion: 0.15
Nodes (5): ImportSummary(), fiveToolsBack(), fiveToolsKey(), sourceInUseMsg(), sourceUsers()

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - ".sourceIssues"
Cohesion: 0.31
Nodes (8): pinnedRefs(), sourceIssue(), featKey(), pinFeat(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 66 - "Session"
Cohesion: 0.08
Nodes (28): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+20 more)

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "starterEnv"
Cohesion: 0.14
Nodes (20): replaceDir(), TestFiveToolsBooks(), DownloadOpen5e(), serveOpen5e(), TestDownloadOpen5e(), TestDownloadOpen5eBadName(), TestDownloadOpen5eFailureKeepsOld(), TestDownloadOpen5eUnknownDoc() (+12 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.24
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - "Book"
Cohesion: 0.11
Nodes (26): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+18 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "config.go"
Cohesion: 0.14
Nodes (16): loopbackHostOnly(), parseFlags(), TestLoopbackHostOnly(), TestParseFlags(), Config, lookupFunc, describe(), fromLookup() (+8 more)

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
Nodes (8): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), randomID(), sync()

### Community 85 - "Run Script"
Cohesion: 0.53
Nodes (8): changes(), GOFLAGS, lock(), restart_server(), run.sh script, stop_server(), superseded(), unlock()

### Community 86 - "mirror"
Cohesion: 0.25
Nodes (8): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), TestMirrorGetMissingFile(), open5eFiles(), fetch()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".source"
Cohesion: 0.10
Nodes (20): UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceKeyWithSlash() (+12 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - ".levelUp"
Cohesion: 0.14
Nodes (15): TextData, FreeCast(), SwapFeatSpell(), Spell, textMarkdown(), spellFromRow(), addClassSpells(), takenFeat() (+7 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "net/http/httptest.ResponseRecorder"
Cohesion: 0.10
Nodes (29): TestEffectFields(), bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex() (+21 more)

### Community 103 - "newEffectFix"
Cohesion: 0.23
Nodes (19): acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectModRows(), TestEffectModsReadOnly() (+11 more)

### Community 105 - "cert.go"
Cohesion: 0.09
Nodes (17): main(), run(), openBrowser(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+9 more)

### Community 106 - ".Router"
Cohesion: 0.21
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (24): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftCopyTargets(), ftDecode() (+16 more)

### Community 108 - ".editEntity"
Cohesion: 0.15
Nodes (9): Entity, Fields, ParseFields(), entityTopics(), linkEntityLook(), saveFieldLinks(), linkCharacterLook(), lookFree() (+1 more)

### Community 111 - "Hub"
Cohesion: 0.10
Nodes (23): ClientInfo, goneDevice, Hub, isDevice(), New(), NewClient(), TestApprovedSurvivesDisconnect(), TestCatchUp() (+15 more)

### Community 112 - "popover_templ.go"
Cohesion: 0.13
Nodes (23): Field, Description(), Header(), state(), Title(), Trigger(), triggerID(), effectBits() (+15 more)

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 117 - "Attachment"
Cohesion: 0.06
Nodes (16): Identity, Attachment, Handlers, browserName(), firstLine(), mdLinkText(), saveScreenshot(), TestUserAgent() (+8 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "takehome.go"
Cohesion: 0.23
Nodes (14): HomeManifest, TakeHomeBaseCharacter, TakeHomeBaseLook, TakeHomeBlock, TakeHomeBlockFile, TakeHomeCampaign, TakeHomeEntry, TakeHomeLevelUp (+6 more)

### Community 127 - "storeSource"
Cohesion: 0.17
Nodes (14): Source, sourceFormat, TestStoreSourceRollsBackOnAfterSave(), formatOf(), Report, ImportProgress(), importSource(), readSource() (+6 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.10
Nodes (27): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+19 more)

### Community 135 - "CompendiumKind"
Cohesion: 0.03
Nodes (53): HomebrewOwner, Site, TestHomebrew(), ValidKind(), ftsQuery(), HomebrewInput, SaveHomebrew(), Search() (+45 more)

### Community 136 - "TakeHomeFile"
Cohesion: 0.30
Nodes (4): TakeHomeFile, readLimited(), ReadTakeHome(), hasBlock()

### Community 137 - "names"
Cohesion: 0.33
Nodes (7): ftRaceAbility(), SwitchFeature, SwitchPlan, PlanSourceSwitch(), names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 139 - "DownloadFiveTools"
Cohesion: 0.14
Nodes (22): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook() (+14 more)

### Community 140 - ".Entry"
Cohesion: 0.10
Nodes (24): MonsterAction, addEntry(), setStat(), MonsterData, TestFiveToolsCoreTraitsShapes(), HomebrewInputFor(), abilityKey(), actionsText() (+16 more)

### Community 143 - "takehome_test.go"
Cohesion: 0.43
Nodes (7): rewriteZip(), sampleTakeHome(), TestReadTakeHomeRefusals(), TestTakeHomeMissingFileOnDisk(), TestTakeHomeRoundTrip(), writeZip(), ZipFile()

### Community 144 - ".deleteUnused"
Cohesion: 0.27
Nodes (5): ListAttachmentFilesRow, attachmentRefs(), fresh(), inCampaign(), FileUse

### Community 145 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 147 - "Campaign"
Cohesion: 0.28
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), writeFile(), nonNil(), TakeHomeFileName(), WriteTakeHome()

### Community 148 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 157 - ".importParam"
Cohesion: 0.39
Nodes (3): ensureCategory(), csvImport, csvImports

### Community 176 - "importLooks"
Cohesion: 0.53
Nodes (6): HomeIDMaps, cmpOr(), importLooks(), importNotes(), importTakeHome(), savedFile

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 461 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **56 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14678433759383241 - nodes in this community are weakly interconnected._
- **Why does `routePath()` connect `wikilink.js` to `testing.T`, `notes.js`, `BlockEditor`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Should `CampaignSources` be split into smaller, more focused modules?**
  _Cohesion score 0.07653061224489796 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `CampaignSources`, `context.Context`, `deathFix`, `export/markdown.go`, `SheetView`, `CompendiumKind`, `combobox_templ.go`, `perm.go`, `sourcesEnv`, `.deleteUnused`, `Campaign`, `App`, `Input`, `dialog_templ.go`, `github.com/a-h/templ.Attributes`, `CombatantView`, `Condition`, `.importParam`, `Layout`, `tabs_templ.go`, `WizardView`, `ParseSheet`, `github.com/a-h/templ.Component`, `testing.T`, `ImagePage`, `Player`, `.renderCompendium`, `starterEnv`, `ParseWiki`, `.source`, `.levelUp`, `net/http/httptest.ResponseRecorder`, `newEffectFix`, `.editEntity`, `popover_templ.go`, `Attachment`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._