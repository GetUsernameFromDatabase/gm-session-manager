# Graph Report - gm-session-manager  (2026-11-26)

## Corpus Check
- 375 files · ~668,703 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4899 nodes · 25602 edges · 164 communities (110 shown, 54 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2555 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `0a52f90b`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Import
- context.Context
- vendorjs/main.go
- createProvider
- export/markdown.go
- SheetView
- combobox.js
- ParseSheet
- Path
- encoding/json.RawMessage
- select.js
- Command
- CompendiumPath
- redirect.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- WizardView
- wikilink.js
- .fail
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- empty_templ.go
- CombatantView
- formSelect
- app.js
- tabbable.js
- useDismiss
- Slugify
- Queries
- command.js
- dialog.js
- time.Time
- ParseWikidot
- tooltip.js
- Character
- tabs_templ.go
- viewerCtx
- composite.js
- Condition
- CustomMod
- Open
- github.com/a-h/templ.Component
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Session
- shadcn-templ Config
- dialog_templ.go
- golangci-lint Config
- TextData
- Sheet
- radiogroup.js
- Player
- ClassInfo
- formErrors
- sqlc Config
- .sourceIssues
- Browser check
- Claude Settings
- Client
- use_transition_status.js
- playwright
- ftPrerequisite
- mirror
- checkbox.js
- ParseAidedd
- switch.js
- Button
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- .compendiumDir
- useTypeahead
- .characterLook
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- CompendiumKind
- mounted
- CharacterFile
- nullID
- profFocus
- cert.go
- .source
- ftBook
- Pending
- BlockEditor
- ListSharedWithRow
- capReader
- build.sh
- looseString
- errBodyWriter
- Before/after pairs
- ftRec
- jsconfig.json
- open5e.go
- testing.M
- perm/CLAUDE.md
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- ImagePage
- Device
- EquipmentOption
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
- planSwitch
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 387 edges
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

## Communities (164 total, 54 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (56): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+48 more)

### Community 1 - "Import"
Cohesion: 0.10
Nodes (37): Source, sourceFormat, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+29 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (135): keepAwake(), ListImageCharactersRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListImageEntitiesRow, GetImageRow, ListImagesRow (+127 more)

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
Nodes (103): Slot, Signed(), Variant, Action(), Card(), Content(), Description(), Header() (+95 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.09
Nodes (43): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeatRaisesCon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass() (+35 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (116): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+108 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (50): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+42 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (55): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+47 more)

### Community 13 - "CompendiumPath"
Cohesion: 0.08
Nodes (43): CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, Empty(), Addon(), InputGroup(), itemDetails(), spellSearch() (+35 more)

### Community 14 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (45): SpellData, ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks() (+37 more)

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
Nodes (22): SetCookie(), writeCookie(), collectHeld(), formInt(), formIntQuery(), isHTMX(), redirect(), homebrewTabURL() (+14 more)

### Community 21 - "WizardView"
Cohesion: 0.08
Nodes (47): Props, Defenses, Footer(), Checkbox(), initialChecked(), Label(), initialValue(), Item() (+39 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - ".fail"
Cohesion: 0.07
Nodes (15): targetColumns(), Character, LinkTarget, NewCharacterFile(), CanChangeSubclass(), CanCharacter(), ImageLook, saveFieldLinks() (+7 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (176): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+168 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "empty_templ.go"
Cohesion: 0.08
Nodes (54): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Content() (+46 more)

### Community 28 - "CombatantView"
Cohesion: 0.07
Nodes (55): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+47 more)

### Community 29 - "formSelect"
Cohesion: 0.05
Nodes (91): ImageCategory, Combobox(), Content(), Input(), Item(), List(), state(), CloseFor() (+83 more)

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
Cohesion: 0.05
Nodes (64): internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice(), TestOptions() (+56 more)

### Community 34 - "Queries"
Cohesion: 0.07
Nodes (43): ClassFeature, Decode(), ClassData, dnd5eFeatures(), ListSubclassesOfRow, Queries, TestSlotMaxima(), ClassFeature (+35 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "time.Time"
Cohesion: 0.04
Nodes (31): IsLoopback(), ClearCookie(), Identity, FromContext(), RequireAdmin(), ListAttachmentFilesRow, browserName(), firstLine() (+23 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (23): block, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Character"
Cohesion: 0.15
Nodes (13): CoreTraits(), Character, checkSkillPick(), classSkillPick(), postedSkills(), legacyBackground(), backgroundRaises(), setupOnce() (+5 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.13
Nodes (32): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+24 more)

### Community 42 - "viewerCtx"
Cohesion: 0.15
Nodes (12): Topic, TopicKind, heldCollected(), notFoundIsNil(), parseLiveKey(), Roster, WithCampaign(), heldTopics (+4 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Condition"
Cohesion: 0.05
Nodes (53): Combatant, Condition, CreateConditionParams, ListCombatantsRow, EffectEnd, combatantRow(), passTurn(), rollFor() (+45 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (24): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), AddFeat(), SaveFeature() (+16 more)

### Community 47 - "Open"
Cohesion: 0.09
Nodes (32): linkTarget, migration, internal/db, apply(), loadMigrations(), Migrate(), Open(), openPool() (+24 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.07
Nodes (58): Props, Variant, Badge(), badgeClasses(), Scripts(), Toast(), Toaster(), toastIcon() (+50 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (160): Defense, TestIsLoopback(), TestDecode(), TestCoreTraits(), TestTokensReadAsLabels(), replaceDir(), TestFiveToolsBooks(), RemoveSourceDir() (+152 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Session"
Cohesion: 0.07
Nodes (30): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+22 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "dialog_templ.go"
Cohesion: 0.10
Nodes (43): ActionProps, DescriptionProps, Props, TitleProps, Variant, ctxKey, ctxState, Props (+35 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "TextData"
Cohesion: 0.17
Nodes (13): TextData, Spellbook(), TestSpellbook(), CountWord(), textMarkdown(), takenFeat(), applyTraitText(), featMasteries() (+5 more)

### Community 59 - "Sheet"
Cohesion: 0.05
Nodes (34): Part, Skill, ItemData, findItem(), parseItem(), TestCustomMods(), TestModifier(), TestResistances() (+26 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.11
Nodes (28): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+20 more)

### Community 62 - "ClassInfo"
Cohesion: 0.07
Nodes (22): LevelUpChoice, armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorGrants() (+14 more)

### Community 63 - "formErrors"
Cohesion: 0.25
Nodes (6): Handlers, keepsBody(), formErrors(), fieldError, pickError, userError

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - ".sourceIssues"
Cohesion: 0.36
Nodes (7): pinnedRefs(), sourceIssue(), featKey(), pinFeat(), SourceCopy, SourceIssue, sheetRef

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Client"
Cohesion: 0.06
Nodes (36): ClientInfo, countingHandler, event, fakeHandler, goneDevice, Handler, message, patchMessage (+28 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.13
Nodes (15): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+7 more)

### Community 72 - "mirror"
Cohesion: 0.13
Nodes (22): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+14 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Button"
Cohesion: 0.07
Nodes (71): ListEntryCopiesRow, Habitat, UI conventions, Button(), Separator(), CampaignOverview(), CampaignRoster(), CharacterPage() (+63 more)

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
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - ".compendiumDir"
Cohesion: 0.13
Nodes (9): ensureCategory(), loadSettings(), saveSettings(), TestSettingsFile(), appState, csvImport, csvImports, fiveToolsJob (+1 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".characterLook"
Cohesion: 0.25
Nodes (3): SetCharacterCurrentImageParams, ToggleCharacterImageRevealedParams, UnlinkCharacterImageParams

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "CompendiumKind"
Cohesion: 0.03
Nodes (64): HomebrewOwner, Site, TestHomebrew(), dataFor(), DecodeData(), ValidKind(), ftsQuery(), HomebrewInput (+56 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 103 - "nullID"
Cohesion: 0.10
Nodes (35): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead(), acBonusForm() (+27 more)

### Community 105 - "cert.go"
Cohesion: 0.10
Nodes (19): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+11 more)

### Community 106 - ".source"
Cohesion: 0.10
Nodes (20): UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceKeyWithSlash() (+12 more)

### Community 107 - "ftBook"
Cohesion: 0.20
Nodes (15): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftParts(), ftRange() (+7 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "ftRec"
Cohesion: 0.27
Nodes (8): ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve(), ftVersions()

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (22): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+14 more)

### Community 139 - "DownloadFiveTools"
Cohesion: 0.11
Nodes (26): EquipmentOptions(), TestEquipmentOptions(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition() (+18 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 143 - "ImagePage"
Cohesion: 0.08
Nodes (38): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Field, Look, attrString(), generateSVG(), getIconContent() (+30 more)

### Community 144 - "Device"
Cohesion: 0.40
Nodes (8): deviceRow(), rosterCard(), devicesWhere(), agoTime(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 146 - "EquipmentOption"
Cohesion: 0.40
Nodes (5): EquipmentItem, EquipmentOption, TestOptionText(), EquipmentChoice, optionText()

### Community 175 - "planSwitch"
Cohesion: 0.13
Nodes (13): featureSwitch, ftRaceAbility(), featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch() (+5 more)

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 456 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **54 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Command`, `perm.go`, `.fail`, `Slugify`, `Queries`, `time.Time`, `Character`, `viewerCtx`, `Condition`, `testing.T`, `TextData`, `Sheet`, `.sourceIssues`, `Client`, `.compendiumDir`, `.characterLook`, `CompendiumKind`, `cert.go`, `.source`, `Pending`?**
  _High betweenness centrality (0.057) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1427035466006834 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Import`, `context.Context`, `export/markdown.go`, `SheetView`, `ParseSheet`, `CompendiumPath`, `ImagePage`, `perm.go`, `Device`, `App`, `WizardView`, `.fail`, `github.com/a-h/templ.Attributes`, `empty_templ.go`, `CombatantView`, `formSelect`, `Slugify`, `time.Time`, `tabs_templ.go`, `Open`, `github.com/a-h/templ.Component`, `testing.T`, `Session`, `dialog_templ.go`, `Player`, `Button`, `ParseWiki`, `.compendiumDir`, `CompendiumKind`, `nullID`, `.source`?**
  _High betweenness centrality (0.037) - this node is a cross-community bridge._
- **Should `Import` be split into smaller, more focused modules?**
  _Cohesion score 0.09986504723346828 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.030) - this node is a cross-community bridge._