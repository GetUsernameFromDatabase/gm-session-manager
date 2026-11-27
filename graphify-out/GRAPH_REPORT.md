# Graph Report - gm-session-manager  (2026-11-27)

## Corpus Check
- 413 files · ~747,136 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 61 file(s) not represented in the graph (top: .templ 51, .css 4, (none) 3)

## Summary
- 5502 nodes · 29930 edges · 181 communities (123 shown, 58 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2973 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `c1d4bbae`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Import
- context.Context
- takehome_test.go
- createProvider
- export/markdown.go
- SheetView
- combobox.js
- ParseSheet
- CampaignSources
- encoding/json.RawMessage
- select.js
- Queries
- entityForm
- Open
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- run
- Button
- wikilink.js
- HomeFileReview
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- .SpellAbilityOf
- CombatantView
- TestBringBackAccept
- app.js
- tabbable.js
- useDismiss
- ClassInfo
- Encounter
- command.js
- dialog.js
- FromContext
- ParseWikidot
- tooltip.js
- github.com/a-h/templ.Component
- tabs_templ.go
- WizardView
- Slugify
- composite.js
- Topic
- CustomMod
- vendorjs/main.go
- CompendiumPath
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- .equipmentChoices
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- Book
- .Router
- radiogroup.js
- Player
- CompendiumKind
- Weapon
- sqlc Config
- App
- handoff
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- planBringBack
- checkbox.js
- ParseAidedd
- switch.js
- starterEnv
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
- Slugify
- mounted
- newPlayerEnv
- newEffectFix
- profFocus
- DownloadOpen5e
- ParseArmorTraining
- ftBook
- empty_templ.go
- BlockEditor
- fakeHandler
- cert.go
- buildReview
- build.sh
- .sendFeedback
- looseString
- errBodyWriter
- Player
- State
- Before/after pairs
- jsconfig.json
- open5e.go
- testing.M
- Input
- Client
- ASIError
- .Entry
- sourceUsedError
- Pending
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
- Condition
- Browser check
- Note
- player-app.sh
- effectSheet
- ListSharedWithRow
- IMPORT_FILTER_CAT_PREFIX
- .fiveToolsRemove
- .count
- Item
- subclassPanels
- .sourceIssues
- names
- formErrors

## God Nodes (most connected - your core abstractions)
1. `App` - 453 edges
2. `Path()` - 341 edges
3. `newTestEnv()` - 197 edges
4. `Button()` - 144 edges
5. `Sheet` - 141 edges
6. `CN()` - 114 edges
7. `Queries` - 89 edges
8. `SheetView` - 80 edges
9. `IfElse()` - 67 edges
10. `ParseSheet()` - 64 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
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

## Communities (181 total, 58 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (58): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+50 more)

### Community 1 - "Import"
Cohesion: 0.07
Nodes (46): sourceFormat, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+38 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (140): keepAwake(), ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumSource, Encounter, LevelUp (+132 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.25
Nodes (19): ReadTakeHome(), rewriteZip(), sampleBringBack(), sampleFallback(), sampleNewCharacter(), sampleTakeHome(), TestBringBackFallbackRoundTrip(), TestBringBackRoundTrip() (+11 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.14
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (98): Slot, Signed(), Variant, Action(), Card(), Content(), Description(), Header() (+90 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.09
Nodes (42): TestUpgradeSheet(), ParseSheet(), homeTableSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeatRaisesCon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI() (+34 more)

### Community 9 - "CampaignSources"
Cohesion: 0.10
Nodes (30): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), campaignSourceRow(), CampaignSources(), CampaignSourcesLive() (+22 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (40): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+32 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Queries"
Cohesion: 0.02
Nodes (149): HomeManifest, TakeHomeBaseCharacter, TakeHomeCampaign, TakeHomeFallback, TakeHomeLevelUp, TakeHomeManifest, targetColumns(), ListCharacterImagesRow (+141 more)

### Community 13 - "entityForm"
Cohesion: 0.16
Nodes (15): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListBacklinksRow, Field, EntityDetail(), entityForm(), EntityPage() (+7 more)

### Community 14 - "Open"
Cohesion: 0.04
Nodes (59): fixture, tokens, linkTarget, migration, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared() (+51 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (60): ListCharacterNamesRow, GetBlockRow, Visibility, internal/perm, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect() (+52 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "run"
Cohesion: 0.10
Nodes (22): loopbackHostOnly(), main(), parseFlags(), run(), TestLoopbackHostOnly(), TestParseFlags(), openBrowser(), Config (+14 more)

### Community 21 - "Button"
Cohesion: 0.05
Nodes (88): CountHabitatsRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, UI conventions, Button(), Spinner(), Cell() (+80 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - "HomeFileReview"
Cohesion: 0.15
Nodes (16): campaignLink(), PendingRequests(), PendingGroup, bringBackCard(), HomeFilePage(), HomeFileRejected(), HomeFileReview(), newCharacterCard() (+8 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (173): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+165 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 28 - "CombatantView"
Cohesion: 0.09
Nodes (39): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+31 more)

### Community 29 - "TestBringBackAccept"
Cohesion: 0.33
Nodes (10): TakeHome, testEnv, homeSheet(), TestBringBackAccept(), TestBringBackPlayerUpload(), TestBringBackReject(), TestBringBackSources(), TestBringBackStale() (+2 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "ClassInfo"
Cohesion: 0.06
Nodes (39): featureSwitch, Decode(), SpellData, Env, TestSlotMaxima(), ClassFeature, ClassInfo, LevelUpPlan (+31 more)

### Community 34 - "Encounter"
Cohesion: 0.21
Nodes (8): CanAddSources(), notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "FromContext"
Cohesion: 0.07
Nodes (18): Identity, FromContext(), browserName(), firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify() (+10 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.11
Nodes (27): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt() (+19 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (85): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+77 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.08
Nodes (43): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+35 more)

### Community 42 - "WizardView"
Cohesion: 0.08
Nodes (51): Props, Footer(), Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup() (+43 more)

### Community 43 - "Slugify"
Cohesion: 0.16
Nodes (21): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+13 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Topic"
Cohesion: 0.15
Nodes (8): event, Handler, stateHandler, Topic, TopicKind, heldCollected(), noteTopics(), heldTopics

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (24): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+16 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 48 - "CompendiumPath"
Cohesion: 0.04
Nodes (90): CountEntriesBySourceRow, CountSharedBySourceRow, ListHomebrewLinksRow, SearchCompendiumRow, Empty(), BuilderPickItems(), BuilderPicks(), itemDetails() (+82 more)

### Community 49 - "Path"
Cohesion: 0.03
Nodes (126): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), Gotchas, internal/web, Values Go and JS share (+118 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (127): Defense, TestDecode(), EquipmentOptions(), TestEquipmentOptions(), feature(), TestFiveToolsClasses(), RemoveSourceDir(), TestRemoveSourceDir() (+119 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - ".equipmentChoices"
Cohesion: 0.18
Nodes (11): EquipmentItem, EquipmentOption, ItemData, ArmorEffect(), WornArmor, compendiumWeapon(), TestCompendiumWeapon(), applyEquipment() (+3 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.09
Nodes (18): LevelUpChoice, TestArmorGrants(), AddResistance(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+10 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Book"
Cohesion: 0.08
Nodes (39): rawBook, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants() (+31 more)

### Community 59 - ".Router"
Cohesion: 0.21
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.09
Nodes (32): Player(), TestCanAddSources(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState() (+24 more)

### Community 62 - "CompendiumKind"
Cohesion: 0.03
Nodes (53): Site, dataFor(), DecodeData(), ValidKind(), FetchPage(), Page, pageSource(), PageSourceKey() (+45 more)

### Community 63 - "Weapon"
Cohesion: 0.06
Nodes (25): Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield(), addedRemoved() (+17 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "App"
Cohesion: 0.03
Nodes (47): Campaign, Character, Entity, Fields, ParseFields(), browsePageSize(), compendiumBack(), compendiumPath() (+39 more)

### Community 66 - "handoff"
Cohesion: 0.24
Nodes (4): closeServers(), removeFile(), shutdownServers(), handoff

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.06
Nodes (58): Amount, CondOpKind, Effects, userText, ArmorKindNamed(), parseFeature(), parseProf(), parseProfSet() (+50 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 72 - "planBringBack"
Cohesion: 0.12
Nodes (25): sheetMap, finish(), GroupDiff, GroupKeys(), GroupLabel(), Merge(), signature(), TakeBuild() (+17 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "starterEnv"
Cohesion: 0.30
Nodes (10): testEnv, pickerTicks(), serveStarterSets(), starterEnv(), TestImportsTickFromCampaignOnly(), TestPartialDownloadTicks(), TestPickerKeepsJobTicks(), TestSourcesCardStarterSets() (+2 more)

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
Cohesion: 0.32
Nodes (5): bookWriter, mirror, downloadBook(), filterDoc(), filterSource()

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

### Community 100 - "Slugify"
Cohesion: 0.06
Nodes (49): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), fiveToolsPage(), AbilityKeys() (+41 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.07
Nodes (46): eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack(), TestHandoffExpiresAndReplaces() (+38 more)

### Community 103 - "newEffectFix"
Cohesion: 0.10
Nodes (37): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead() (+29 more)

### Community 105 - "DownloadOpen5e"
Cohesion: 0.25
Nodes (11): fetchJSON(), replaceDir(), DownloadOpen5e(), open5eFiles(), serveOpen5e(), TestDownloadOpen5e(), TestDownloadOpen5eBadName(), TestDownloadOpen5eFailureKeepsOld() (+3 more)

### Community 106 - "ParseArmorTraining"
Cohesion: 0.13
Nodes (11): ftProfTexts(), ftSkillNames(), armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining() (+3 more)

### Community 107 - "ftBook"
Cohesion: 0.12
Nodes (28): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftParts(), ftPrerequisite() (+20 more)

### Community 108 - "empty_templ.go"
Cohesion: 0.06
Nodes (64): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, SearchReadableBlocksRow (+56 more)

### Community 113 - "cert.go"
Cohesion: 0.06
Nodes (24): keepAwake(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname() (+16 more)

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Player"
Cohesion: 0.22
Nodes (21): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+13 more)

### Community 127 - "State"
Cohesion: 0.17
Nodes (11): patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff(), devicesWhere(), deviceLabel() (+3 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.10
Nodes (27): Entry, fixture, fixtures, HomebrewOwner, normalizer, o5eProperty, Source, TestOpen5eHabitats() (+19 more)

### Community 136 - "Input"
Cohesion: 0.04
Nodes (88): ImageCategory, Defenses, Description(), Field(), Label(), Title(), Input(), Content() (+80 more)

### Community 137 - "Client"
Cohesion: 0.16
Nodes (6): ClientInfo, goneDevice, message, Client, Hub, isDevice()

### Community 140 - ".Entry"
Cohesion: 0.08
Nodes (30): MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools() (+22 more)

### Community 148 - "Pending"
Cohesion: 0.14
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 177 - "Condition"
Cohesion: 0.06
Nodes (32): Combatant, Condition, CreateConditionParams, ListCombatantsRow, combatantRow(), rollFor(), stepTurn(), TestMalformedEffectSkipped() (+24 more)

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 179 - "Note"
Cohesion: 0.11
Nodes (10): Note, NoteKind, setAudience(), CreateNoteAtParams, CreateNoteParams, ListReadableNotesParams, RenameNoteParams, blockAccessForm (+2 more)

### Community 181 - "effectSheet"
Cohesion: 0.40
Nodes (5): Handlers, TestKeepsBodyNewCharacter(), TestKeepsBodyBringBack(), keepsBody(), userError

### Community 186 - ".fiveToolsRemove"
Cohesion: 0.07
Nodes (19): FiveTools2024Set(), TestFiveTools2024Set(), ImportSummary(), fiveToolsKey(), TestTickSet(), tickBooks(), tickSet(), csvSource() (+11 more)

### Community 192 - ".count"
Cohesion: 0.20
Nodes (11): testEnv, mentionDoc(), newExtrasEnv(), TestBringBackAcceptExtras(), TestBringBackAcceptFiles(), TestBringBackStaleExtras(), TestNewCharacterFromGM(), TestNewCharacterFromVisitor() (+3 more)

### Community 193 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 204 - "subclassPanels"
Cohesion: 0.18
Nodes (9): ListSubclassesOfRow, TestSubclassPanels(), subclassIndex(), subclassOptions(), subclassPanels(), SubclassPanel, ListSubclassesOfParams, SubclassChangeView (+1 more)

### Community 211 - ".sourceIssues"
Cohesion: 0.31
Nodes (8): pinnedRefs(), sourceIssue(), featKey(), pinFeat(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 212 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 220 - "formErrors"
Cohesion: 0.40
Nodes (3): formErrors(), fieldError, pickError

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 482 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **58 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Import`, `context.Context`, `export/markdown.go`, `SheetView`, `ParseSheet`, `Input`, `CampaignSources`, `Queries`, `entityForm`, `Open`, `perm.go`, `Button`, `HomeFileReview`, `github.com/a-h/templ.Attributes`, `CombatantView`, `TestBringBackAccept`, `github.com/a-h/templ.Component`, `tabs_templ.go`, `WizardView`, `CompendiumPath`, `testing.T`, `.fiveToolsRemove`, `Player`, `.count`, `App`, `starterEnv`, `ParseWiki`, `.source`, `Slugify`, `newPlayerEnv`, `newEffectFix`, `empty_templ.go`, `.sendFeedback`?**
  _High betweenness centrality (0.062) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Import` be split into smaller, more focused modules?**
  _Cohesion score 0.0726950354609929 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `Path` to `testing.T`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.0071780679754389 - nodes in this community are weakly interconnected._
- **Why does `routePath()` connect `wikilink.js` to `Path`, `notes.js`, `BlockEditor`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._