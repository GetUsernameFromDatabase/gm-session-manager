# Graph Report - gm-session-manager  (2026-11-27)

## Corpus Check
- 398 files · ~700,569 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 60 file(s) not represented in the graph (top: .templ 50, .css 4, (none) 3)

## Summary
- 5158 nodes · 27363 edges · 178 communities (120 shown, 58 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2702 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `8d598899`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- nullID
- context.Context
- takehome.go
- createProvider
- export/markdown.go
- Button
- combobox.js
- Path
- Open
- encoding/json.RawMessage
- select.js
- Command
- Import
- database/sql.NullString
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- Input
- wikilink.js
- Book
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Sheet
- CombatantView
- time.Time
- app.js
- tabbable.js
- useDismiss
- ParseArmorTraining
- Sum
- command.js
- dialog.js
- Amount
- ParseWikidot
- tooltip.js
- Layout
- tabs_templ.go
- WizardView
- Sheet
- composite.js
- CampaignSources
- CustomMod
- Item
- CompendiumPath
- spellBoxes
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Character
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- deathFix
- .device
- radiogroup.js
- Player
- WeaponKind
- .info
- sqlc Config
- viewerCtx
- handoff
- Claude Settings
- starterEnv
- use_transition_status.js
- playwright
- ftPrerequisite
- parseMod
- checkbox.js
- ParseAidedd
- switch.js
- run
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
- .sendFeedback
- mounted
- newPlayerEnv
- newEffectFix
- profFocus
- cert.go
- condMods
- ftBook
- SearchReadableBlocksRow
- BlockEditor
- .ListSubracesOf
- hub_test.go
- capReader
- buildReview
- State
- build.sh
- Pending
- looseString
- errBodyWriter
- Topic
- .editEntity
- Before/after pairs
- jsconfig.json
- open5e.go
- testing.M
- github.com/a-h/templ.Component
- combobox_templ.go
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- alert_templ.go
- fakeHandler
- planSwitch
- Client
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
- storeSource
- tooltip_templ.go
- ListSharedWithRow
- Browser check
- player-app.sh
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 421 edges
2. `Path()` - 299 edges
3. `newTestEnv()` - 185 edges
4. `Sheet` - 134 edges
5. `Button()` - 133 edges
6. `CN()` - 114 edges
7. `SheetView` - 80 edges
8. `IfElse()` - 67 edges
9. `Queries` - 64 edges
10. `Command` - 60 edges

## Surprising Connections (you probably didn't know these)
- `Notes` --references--> `linkClick()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/assets/js/wikilink.js
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `SpellData`  [INFERRED]
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

## Communities (178 total, 58 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (56): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+48 more)

### Community 1 - "nullID"
Cohesion: 0.05
Nodes (46): Combatant, Condition, CreateConditionParams, ListCombatantsRow, EffectEnd, combatantRow(), passTurn(), rollFor() (+38 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (144): keepAwake(), HomebrewOwner, sourceFormat, TestHomebrew(), bindFiveToolsImages(), ftsQuery(), SaveHomebrew(), Search() (+136 more)

### Community 3 - "takehome.go"
Cohesion: 0.06
Nodes (52): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+44 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.14
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 6 - "Button"
Cohesion: 0.04
Nodes (140): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+132 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Path"
Cohesion: 0.04
Nodes (125): TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestConditionExpiresOnTurn() (+117 more)

### Community 9 - "Open"
Cohesion: 0.10
Nodes (30): migration, internal/db, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp() (+22 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (50): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+42 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.10
Nodes (41): CondOpKind, Effects, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn(), parseCoin() (+33 more)

### Community 13 - "Import"
Cohesion: 0.09
Nodes (38): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+30 more)

### Community 14 - "database/sql.NullString"
Cohesion: 0.07
Nodes (20): GetImageRow, ensureCategory(), imageItem(), imageSource(), imageVersion(), nullString(), seenBy(), BindMonsterImageParams (+12 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (57): ListCharacterNamesRow, GetBlockRow, Visibility, internal/perm, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect() (+49 more)

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
Nodes (49): Identity, FromContext(), RequireAdmin(), Campaign, CanChangeSubclass(), CanCharacter(), TestCanChangeSubclass(), browsePageSize() (+41 more)

### Community 21 - "Input"
Cohesion: 0.04
Nodes (68): Defenses, Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select() (+60 more)

### Community 22 - "wikilink.js"
Cohesion: 0.17
Nodes (18): nodeActions(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages, LAYOUTS (+10 more)

### Community 23 - "Book"
Cohesion: 0.11
Nodes (26): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+18 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (145): Props, Props, Variant, Props, Size, ActionProps, ContentProps, DescriptionProps (+137 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Sheet"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 28 - "CombatantView"
Cohesion: 0.07
Nodes (54): Description(), Header(), state(), Title(), Trigger(), triggerID(), effectBits(), addCombatantForm() (+46 more)

### Community 29 - "time.Time"
Cohesion: 0.04
Nodes (26): ListAttachmentFilesRow, ListImagesRow, Attachment, browserName(), firstLine(), folderName(), mdLinkText(), saveScreenshot() (+18 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 34 - "Sum"
Cohesion: 0.21
Nodes (3): TestCustomMods(), Sum(), TestWithEffects()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Amount"
Cohesion: 0.22
Nodes (7): Amount, userText, CritDamage(), Damage(), parseRevive(), Set(), TestApplyHPDeath()

### Community 38 - "ParseWikidot"
Cohesion: 0.10
Nodes (30): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, dnd5eFeatures(), TestFiveToolsCoreTraitsShapes(), withCoreTraits() (+22 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Layout"
Cohesion: 0.04
Nodes (80): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ImageCategory, ListBacklinksRow, Entity, Look, Scripts() (+72 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.08
Nodes (48): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+40 more)

### Community 42 - "WizardView"
Cohesion: 0.07
Nodes (59): Props, LevelUpPlan, Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup() (+51 more)

### Community 43 - "Sheet"
Cohesion: 0.03
Nodes (131): EquipmentItem, internal/compendium, Cantrips(), EquipmentOptions(), firstTable(), EquipmentOption, LevelSpells(), OptionPrerequisite() (+123 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "CampaignSources"
Cohesion: 0.09
Nodes (28): campaignSourceRow(), CampaignSources(), CampaignSourcesLive(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), CampaignOverview(), CampaignRoster() (+20 more)

### Community 46 - "CustomMod"
Cohesion: 0.09
Nodes (19): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+11 more)

### Community 47 - "Item"
Cohesion: 0.19
Nodes (9): findItem(), parseArmor(), parseItem(), ArmorEffect(), TestGearCommands(), ensureItem(), Item, WornArmor (+1 more)

### Community 48 - "CompendiumPath"
Cohesion: 0.05
Nodes (56): CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, Habitat, Empty() (+48 more)

### Community 49 - "spellBoxes"
Cohesion: 0.67
Nodes (3): TestTaughtPicked(), taughtPicked(), taughtRules()

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (135): Defense, TestDecode(), ftRaceAbility(), RemoveSourceDir(), TestRemoveSourceDir(), serveOpen5e(), TestDownloadOpen5e(), TestDownloadOpen5eBadName() (+127 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Character"
Cohesion: 0.06
Nodes (25): Character, SizeNamed(), CharacterFile, NewCharacterFile(), HomeIDMaps, ZipFileName(), Handlers, parseCharacterFile() (+17 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.09
Nodes (16): LevelUpChoice, TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+8 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 59 - ".device"
Cohesion: 0.38
Nodes (5): testEnv, TestRequestLog(), TestRequestLogCharacterRemoved(), TestRequestLogReusedID(), TestRequestSent()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.07
Nodes (40): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+32 more)

### Community 62 - "WeaponKind"
Cohesion: 0.50
Nodes (4): Completions(), parseMastery(), TestWeaponKind(), WeaponKind()

### Community 63 - ".info"
Cohesion: 0.50
Nodes (3): ClientInfo, goneDevice, TestFiveToolsClassLevelUp()

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "viewerCtx"
Cohesion: 0.21
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 66 - "handoff"
Cohesion: 0.28
Nodes (4): closeServers(), removeFile(), shutdownServers(), handoff

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "starterEnv"
Cohesion: 0.30
Nodes (10): testEnv, pickerTicks(), serveStarterSets(), starterEnv(), TestImportsTickFromCampaignOnly(), TestPartialDownloadTicks(), TestPickerKeepsJobTicks(), TestSourcesCardStarterSets() (+2 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 72 - "parseMod"
Cohesion: 0.50
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "run"
Cohesion: 0.04
Nodes (50): fixture, tokens, loopbackHostOnly(), main(), parseFlags(), run(), TestLoopbackHostOnly(), TestParseFlags() (+42 more)

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
Nodes (8): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, openWeb(), randomID(), KEEP_IMAGES, LINK_PATHS

### Community 85 - "Run Script"
Cohesion: 0.53
Nodes (8): changes(), GOFLAGS, lock(), restart_server(), run.sh script, stop_server(), superseded(), unlock()

### Community 86 - "mirror"
Cohesion: 0.20
Nodes (12): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), filterSource(), replaceDir(), TestMirrorGetMissingFile() (+4 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".source"
Cohesion: 0.08
Nodes (21): UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceKeyWithSlash() (+13 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.09
Nodes (35): TestEffectFields(), eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack() (+27 more)

### Community 103 - "newEffectFix"
Cohesion: 0.10
Nodes (39): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), TestMonsterAtZeroHPUnaffected(), acBonusForm(), effectFix, newEffectFix() (+31 more)

### Community 105 - "cert.go"
Cohesion: 0.06
Nodes (24): keepAwake(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname() (+16 more)

### Community 106 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

### Community 107 - "ftBook"
Cohesion: 0.15
Nodes (20): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftParts(), ftPrerequisite() (+12 more)

### Community 109 - "BlockEditor"
Cohesion: 0.22
Nodes (5): attachedPicture(), BlockEditor, json(), links(), sync()

### Community 112 - "hub_test.go"
Cohesion: 0.22
Nodes (19): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+11 more)

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 117 - "State"
Cohesion: 0.15
Nodes (8): Handler, message, patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff()

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

### Community 125 - "Topic"
Cohesion: 0.15
Nodes (6): event, stateHandler, Topic, TopicKind, fiveToolsJob, heldTopics

### Community 127 - ".editEntity"
Cohesion: 0.08
Nodes (16): linkTarget, linkTargets(), targetColumns(), Entity, LinkTarget, Fields, ParseFields(), entityTopics() (+8 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 136 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (86): SearchCompendiumRow, SwitchFeature, SwitchPlan, UI conventions, Badge(), Toast(), Toaster(), toastIcon() (+78 more)

### Community 137 - "combobox_templ.go"
Cohesion: 0.05
Nodes (77): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+69 more)

### Community 139 - "DownloadFiveTools"
Cohesion: 0.13
Nodes (23): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook(), DownloadFiveTools() (+15 more)

### Community 140 - ".Entry"
Cohesion: 0.08
Nodes (31): MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels() (+23 more)

### Community 144 - "alert_templ.go"
Cohesion: 0.16
Nodes (17): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+9 more)

### Community 147 - "planSwitch"
Cohesion: 0.23
Nodes (6): featureSwitch, parseFeature(), featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 148 - "Client"
Cohesion: 0.20
Nodes (3): Client, Hub, isDevice()

### Community 175 - "storeSource"
Cohesion: 0.13
Nodes (14): Site, Source, TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page (+6 more)

### Community 176 - "tooltip_templ.go"
Cohesion: 0.11
Nodes (25): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks(), closeAfter() (+17 more)

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 464 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **58 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `nullID`, `context.Context`, `database/sql.NullString`, `perm.go`, `Client`, `Book`, `time.Time`, `Sheet`, `testing.T`, `Character`, `viewerCtx`, `handoff`, `run`, `.source`, `cert.go`, `State`, `Pending`, `Topic`, `.editEntity`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14687100893997446 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `nullID`, `export/markdown.go`, `Button`, `github.com/a-h/templ.Component`, `Open`, `combobox_templ.go`, `Import`, `database/sql.NullString`, `perm.go`, `App`, `Input`, `github.com/a-h/templ.Attributes`, `CombatantView`, `time.Time`, `Layout`, `tabs_templ.go`, `WizardView`, `Sheet`, `CampaignSources`, `tooltip_templ.go`, `CompendiumPath`, `testing.T`, `Character`, `deathFix`, `.device`, `Player`, `starterEnv`, `ParseWiki`, `.source`, `.sendFeedback`, `newPlayerEnv`, `newEffectFix`, `.editEntity`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **Should `nullID` be split into smaller, more focused modules?**
  _Cohesion score 0.052943354313217325 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `Import`, `wikilink.js`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._