# Graph Report - gm-session-manager  (2026-11-28)

## Corpus Check
- 412 files · ~749,254 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 60 file(s) not represented in the graph (top: .templ 51, .css 4, (none) 3)

## Summary
- 5602 nodes · 30025 edges · 195 communities (133 shown, 62 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2979 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `6e048ef3`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Slugify
- context.Context
- takehome_test.go
- createProvider
- export/markdown.go
- SheetView
- combobox.js
- ClassInfo
- Visibility
- encoding/json.RawMessage
- select.js
- Queries
- popover_templ.go
- Session
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- run
- Item
- wikilink.js
- HomeFileReview
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Store
- CombatantView
- TestBringBackAccept
- app.js
- tabbable.js
- useDismiss
- planSwitch
- Encounter
- command.js
- dialog.js
- nullID
- ParseWikidot
- tooltip.js
- select_templ.go
- tabs_templ.go
- IfElse
- Entity
- composite.js
- rosterCard
- CustomMod
- vendorjs/main.go
- CompendiumPath
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- ftItemData
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- Spell
- time.Time
- radiogroup.js
- InputForEntry
- .Router
- Sheet
- sqlc Config
- App
- handoff
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- .combatantHP
- checkbox.js
- ParseAidedd
- switch.js
- .q
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- NameDesc
- mirror
- useTypeahead
- .source
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- CompendiumKind
- mounted
- newPlayerEnv
- nullID
- profFocus
- cert_test.go
- NewClient
- ftBook
- github.com/a-h/templ.Component
- BlockEditor
- ftMod
- Topic
- redirect.go
- buildReview
- vars
- storeSource
- .sendFeedback
- looseString
- errBodyWriter
- Pending
- task
- Before/after pairs
- env
- jsconfig.json
- open5e.go
- testing.M
- SearchReadableBlocksRow
- fakeHandler
- Client
- ASIError
- .Entry
- sourceUsedError
- run.sh
- spellCounts
- Merge
- busyIndicator
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
- pause_windows.go
- Browser check
- dev:changed
- release.sh
- TestPagesRender
- IMPORT_FILTER_CAT_PREFIX
- Book
- dev:build
- build.sh
- before
- generate:web
- release
- dev:generate
- lint
- experiments
- test
- capReader
- tasks
- Button
- lint.sh

## God Nodes (most connected - your core abstractions)
1. `App` - 453 edges
2. `Path()` - 341 edges
3. `newTestEnv()` - 199 edges
4. `Button()` - 145 edges
5. `Sheet` - 141 edges
6. `CN()` - 114 edges
7. `Queries` - 89 edges
8. `SheetView` - 80 edges
9. `IfElse()` - 68 edges
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

## Communities (195 total, 62 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (58): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+50 more)

### Community 1 - "Slugify"
Cohesion: 0.06
Nodes (50): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), fiveToolsPage(), AbilityKeys() (+42 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (151): keepAwake(), ListCharacterImagesRow, ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Attachment, Condition, Encounter (+143 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.18
Nodes (21): readLimited(), ReadTakeHome(), rewriteZip(), sampleBringBack(), sampleFallback(), sampleNewCharacter(), sampleTakeHome(), TestBringBackFallbackRoundTrip() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.14
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (106): Signed(), Variant, Action(), Card(), Content(), Description(), Header(), Title() (+98 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ClassInfo"
Cohesion: 0.07
Nodes (34): Decode(), Completions(), parseMastery(), TestResistances(), TestSlotMaxima(), TestWeaponKind(), ClassInfo, classAbility() (+26 more)

### Community 9 - "Visibility"
Cohesion: 0.06
Nodes (23): ListAttachmentFilesRow, GetBlockRow, Visibility, BlockVisibilities(), TestAtHomeBlockVisibilities(), TestBlockVisibilities(), ValidBlockVisibility(), fresh() (+15 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (43): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+35 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Queries"
Cohesion: 0.03
Nodes (133): HomeManifest, TakeHomeBaseCharacter, TakeHomeCampaign, TakeHomeFallback, TakeHomeLevelUp, TakeHomeManifest, bindFiveToolsImages(), targetColumns() (+125 more)

### Community 13 - "popover_templ.go"
Cohesion: 0.11
Nodes (33): Entity, Look, UI conventions, Content(), Description(), Header(), Popover(), state() (+25 more)

### Community 14 - "Session"
Cohesion: 0.08
Nodes (26): fixture, tokens, loopbackHostOnly(), TestLoopbackHostOnly(), constantTimeEqual(), credentialsMatch(), isCleared(), newFixture() (+18 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.04
Nodes (70): ListCharacterNamesRow, ListCombatantsRow, internal/perm, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), CanAddEffect() (+62 more)

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
Cohesion: 0.11
Nodes (20): main(), parseFlags(), run(), TestParseFlags(), openBrowser(), Config, lookupFunc, bookFolder() (+12 more)

### Community 21 - "Item"
Cohesion: 0.14
Nodes (13): EquipmentItem, EquipmentOption, findItem(), parseArmor(), ArmorEffect(), TestGearCommands(), ensureItem(), Item (+5 more)

### Community 22 - "wikilink.js"
Cohesion: 0.13
Nodes (22): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES (+14 more)

### Community 23 - "HomeFileReview"
Cohesion: 0.06
Nodes (44): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+36 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (152): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+144 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Store"
Cohesion: 0.06
Nodes (56): linkTarget, migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+48 more)

### Community 28 - "CombatantView"
Cohesion: 0.10
Nodes (37): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+29 more)

### Community 29 - "TestBringBackAccept"
Cohesion: 0.12
Nodes (24): TakeHome, testEnv, mentionDoc(), newExtrasEnv(), TestBringBackAcceptExtras(), TestBringBackAcceptFiles(), TestBringBackStaleExtras(), TestNewCharacterFromGM() (+16 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (23): checkLink(), checkLinks(), clearDisplay(), clearSourceSelection(), commandChoices(), commandHint(), commandState(), createItem() (+15 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "planSwitch"
Cohesion: 0.11
Nodes (13): featureSwitch, ftRaceAbility(), featureNames(), SourceRef, SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch() (+5 more)

### Community 34 - "Encounter"
Cohesion: 0.20
Nodes (8): CanAddSources(), notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "nullID"
Cohesion: 0.08
Nodes (47): DeathSaves, nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), deathFix(), effectFix, TestDeathClearsConditions() (+39 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.11
Nodes (27): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt() (+19 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "select_templ.go"
Cohesion: 0.05
Nodes (61): ImageCategory, Defenses, Content(), Group(), initialValue(), Item(), itemLabel(), Label() (+53 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.08
Nodes (44): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+36 more)

### Community 42 - "IfElse"
Cohesion: 0.11
Nodes (38): Label(), IfElse(), inFightSuffix(), StaleBooksBanner(), switchItem(), abilitySelect(), levelUpForm(), spellChecklist() (+30 more)

### Community 43 - "Entity"
Cohesion: 0.18
Nodes (6): Entity, entityTopics(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams, lookOps

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "rosterCard"
Cohesion: 0.29
Nodes (10): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), deviceLabel(), rosterHP() (+2 more)

### Community 46 - "CustomMod"
Cohesion: 0.07
Nodes (26): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), UsesLabel(), CustomMod (+18 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 48 - "CompendiumPath"
Cohesion: 0.03
Nodes (126): SpellData, CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, ListSharedWithRow (+118 more)

### Community 49 - "Path"
Cohesion: 0.03
Nodes (133): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount() (+125 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (178): Defense, TestDecode(), EquipmentOptions(), TestEquipmentOptions(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+170 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "ftItemData"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.07
Nodes (22): LevelUpChoice, armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorGrants() (+14 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Spell"
Cohesion: 0.19
Nodes (12): ClassFeature, LevelUpPlan, LevelUpView, Option, SpellNeed, SubclassPanel, drawsGroup(), ChoiceGroup (+4 more)

### Community 59 - "time.Time"
Cohesion: 0.06
Nodes (18): ListImagesRow, browserName(), firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify(), TestFolderName() (+10 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "InputForEntry"
Cohesion: 0.22
Nodes (8): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 62 - ".Router"
Cohesion: 0.24
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 63 - "Sheet"
Cohesion: 0.09
Nodes (13): Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield(), damageLabel() (+5 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "App"
Cohesion: 0.03
Nodes (45): Identity, FromContext(), ftsQuery(), Search(), Campaign, browsePageSize(), compendiumBack(), compendiumPath() (+37 more)

### Community 66 - "handoff"
Cohesion: 0.20
Nodes (6): IsLoopback(), TestIsLoopback(), closeServers(), removeFile(), shutdownServers(), handoff

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.07
Nodes (56): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+48 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - ".combatantHP"
Cohesion: 0.32
Nodes (7): Combatant, combatantRow(), damageHP(), healHP(), monsterHP(), setHP(), hpChange

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - ".q"
Cohesion: 0.10
Nodes (27): testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+19 more)

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

### Community 85 - "NameDesc"
Cohesion: 0.17
Nodes (21): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+13 more)

### Community 86 - "mirror"
Cohesion: 0.22
Nodes (10): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), replaceDir(), DownloadOpen5e(), open5eFiles() (+2 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".source"
Cohesion: 0.11
Nodes (18): ParsePageURL(), TestParsePageURL(), UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove() (+10 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "CompendiumKind"
Cohesion: 0.03
Nodes (48): dataFor(), DecodeData(), ValidKind(), CountCompendiumRow, ListSubclassesOfRow, CompendiumEntry, CompendiumKind, formErrors() (+40 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.07
Nodes (44): eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack(), TestHandoffExpiresAndReplaces() (+36 more)

### Community 103 - "nullID"
Cohesion: 0.07
Nodes (40): Player(), TestCanAddSources(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState() (+32 more)

### Community 105 - "cert_test.go"
Cohesion: 0.13
Nodes (19): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+11 more)

### Community 106 - "NewClient"
Cohesion: 0.19
Nodes (22): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+14 more)

### Community 107 - "ftBook"
Cohesion: 0.15
Nodes (20): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftParts(), ftRecords() (+12 more)

### Community 108 - "github.com/a-h/templ.Component"
Cohesion: 0.03
Nodes (144): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListImageCharactersRow (+136 more)

### Community 111 - "ftMod"
Cohesion: 0.50
Nodes (3): ftMod(), ftName(), ftReplaceText()

### Community 112 - "Topic"
Cohesion: 0.11
Nodes (13): event, Handler, patchMessage, stateHandler, viewPatch, Topic, TopicKind, Diff() (+5 more)

### Community 113 - "redirect.go"
Cohesion: 0.29
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 117 - "vars"
Cohesion: 0.22
Nodes (8): env, CGO_ENABLED, GOFLAGS, vars, APP, EXE, PKG, version

### Community 118 - "storeSource"
Cohesion: 0.12
Nodes (18): HomebrewOwner, Site, Source, sourceFormat, TestHomebrew(), TestStoreSourceRollsBackOnAfterSave(), HomebrewInput, SaveHomebrew() (+10 more)

### Community 119 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.21
Nodes (6): Slot, errorText(), TestErrorText(), modTargetLabel(), slotLabel(), errBodyWriter

### Community 125 - "Pending"
Cohesion: 0.14
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 127 - "task"
Cohesion: 0.22
Nodes (11): cmds, desc, env, cmds, desc, env, task, APP (+3 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "env"
Cohesion: 0.11
Nodes (26): BIN, DATA_DIR, MARKERS, MODE, NOREBUILD, PORT, PREPARE, RESTART_ARGS (+18 more)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery() (+16 more)

### Community 137 - "Client"
Cohesion: 0.15
Nodes (6): ClientInfo, goneDevice, message, Client, Hub, isDevice()

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 143 - "run.sh"
Cohesion: 0.56
Nodes (8): alive(), lock(), run.sh script, start_server(), stop_server(), superseded(), unchanged(), unlock()

### Community 144 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

### Community 145 - "Merge"
Cohesion: 0.08
Nodes (32): sheetMap, addedRemoved(), describe(), diffNames(), finish(), GroupDiff, GroupKeys(), GroupLabel() (+24 more)

### Community 147 - "busyIndicator"
Cohesion: 0.18
Nodes (10): ImportNotice, FiveToolsProgress(), importBar(), jobNotice(), starterSet(), StarterSets(), FiveToolsJob, starterBase() (+2 more)

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 179 - "dev:changed"
Cohesion: 0.40
Nodes (5): cmds, method, sources, exclude, dev:changed

### Community 180 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 181 - "TestPagesRender"
Cohesion: 0.18
Nodes (8): Handlers, internal/web, TestKeepsBodyNewCharacter(), TestKeepsBodyBringBack(), keepsBody(), fieldError, pickError, userError

### Community 186 - "Book"
Cohesion: 0.11
Nodes (26): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+18 more)

### Community 188 - "dev:build"
Cohesion: 0.40
Nodes (5): cmds, deps, requires, vars, dev:build

### Community 191 - "before"
Cohesion: 0.40
Nodes (5): cmds, desc, vars, before, BEFORE_BIN

### Community 192 - "generate:web"
Cohesion: 0.67
Nodes (3): cmds, desc, generate:web

### Community 193 - "release"
Cohesion: 0.67
Nodes (3): cmds, desc, release

### Community 194 - "dev:generate"
Cohesion: 0.29
Nodes (7): cmd, if, cmds, generates, internal, sources, dev:generate

### Community 195 - "lint"
Cohesion: 0.50
Nodes (4): cmds, desc, silent, lint

### Community 197 - "test"
Cohesion: 0.67
Nodes (3): test, cmds, desc

### Community 200 - "tasks"
Cohesion: 0.12
Nodes (16): cmds, cmds, method, sources, cmds, desc, cmds, internal (+8 more)

### Community 201 - "Button"
Cohesion: 0.06
Nodes (94): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+86 more)

## Knowledge Gaps
- **37 isolated node(s):** `npx`, `@playwright/mcp`, `gm-session-manager`, `testEnv`, `testEnv` (+32 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 534 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **62 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `Path` to `compendium_templ.go`, `export/markdown.go`, `SheetView`, `Visibility`, `Queries`, `popover_templ.go`, `perm.go`, `HomeFileReview`, `github.com/a-h/templ.Attributes`, `Store`, `CombatantView`, `TestBringBackAccept`, `nullID`, `select_templ.go`, `tabs_templ.go`, `IfElse`, `rosterCard`, `CompendiumPath`, `testing.T`, `time.Time`, `App`, `Button`, `.q`, `ParseWiki`, `.source`, `CompendiumKind`, `newPlayerEnv`, `nullID`, `github.com/a-h/templ.Component`, `.sendFeedback`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `gm-session-manager` to the rest of the system?**
  _37 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1434638223275681 - nodes in this community are weakly interconnected._
- **Why does `App` connect `App` to `compendium_templ.go`, `Slugify`, `context.Context`, `ClassInfo`, `Client`, `Visibility`, `Queries`, `perm.go`, `busyIndicator`, `Encounter`, `Entity`, `testing.T`, `Book`, `time.Time`, `.Router`, `handoff`, `Command`, `.combatantHP`, `.q`, `CompendiumKind`, `cert_test.go`, `NewClient`, `Topic`, `Pending`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **Should `Slugify` be split into smaller, more focused modules?**
  _Cohesion score 0.055130784708249496 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `TestPagesRender`, `wikilink.js`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._