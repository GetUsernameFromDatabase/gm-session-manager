# Graph Report - gm-session-manager  (2026-11-28)

## Corpus Check
- 443 files · ~766,498 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 65 file(s) not represented in the graph (top: .templ 51, (none) 4, .css 4)

## Summary
- 5767 nodes · 31006 edges · 213 communities (152 shown, 61 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 3060 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `ea2a6b02`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Slugify
- context.Context
- takehome_test.go
- createProvider
- report
- SheetView
- combobox.js
- .wizardView
- Book
- encoding/json.RawMessage
- select.js
- Queries
- CampaignSources
- Store
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- blockMenu
- CompendiumPath
- wikilink.js
- who
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- CN
- CombatantView
- TestBringBackAccept
- app.js
- tabbable.js
- useDismiss
- Player
- Encounter
- command.js
- dialog.js
- newEffectFix
- ParseWikidot
- tooltip.js
- ItemData
- tabs_templ.go
- WizardView
- Client
- composite.js
- ParseSheet
- Feature
- vendorjs/main.go
- github.com/a-h/templ.Component
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Spell
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- config.go
- rosterCard
- radiogroup.js
- Client
- .Router
- Part
- sqlc Config
- App
- handoff
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- Condition
- checkbox.js
- ParseAidedd
- switch.js
- app
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Topic
- TestNewCharacterFromGM
- useTypeahead
- .source
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- Site
- mounted
- newPlayerEnv
- Session
- profFocus
- connListener
- NewClient
- ftBook
- NameDesc
- BlockEditor
- Merge
- mirror
- starterEnv
- buildReview
- vars
- ftItemData
- testEnv
- looseString
- errBodyWriter
- Pending
- task
- Before/after pairs
- env
- jsconfig.json
- open5e.go
- testing.M
- Queries
- alert_templ.go
- cert.go
- ASIError
- .attachmentsDir
- .Entry
- sourceUsedError
- run.sh
- time.Time
- normalize_test.go
- cert_test.go
- main
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- State
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
- deathFix
- InputForEntry
- Browser check
- dev:changed
- release.sh
- run
- systemd-units.sh
- IMPORT_FILTER_CAT_PREFIX
- ClassInfo
- sameSlug
- internal/web
- dev:build
- .sourceIssues
- build.sh
- before
- pause_windows.go
- release
- dev:generate
- lint
- experiments
- test
- capReader
- net.IP
- tasks
- Input
- lint.sh
- fakeHandler
- WeaponKind
- run:before
- ListSubclassesOfRow
- SourcePreview
- formErrors
- dev:migrations
- run
- run:test-player

## God Nodes (most connected - your core abstractions)
1. `App` - 463 edges
2. `Path()` - 347 edges
3. `newTestEnv()` - 204 edges
4. `Button()` - 147 edges
5. `Sheet` - 142 edges
6. `CN()` - 114 edges
7. `Queries` - 97 edges
8. `SheetView` - 83 edges
9. `IfElse()` - 68 edges
10. `ParseSheet()` - 64 edges

## Surprising Connections (you probably didn't know these)
- `The timer` --references--> `install()`  [INFERRED]
  deploy/README.md → cmd/vendorjs/main.go
- `Notes` --references--> `linkClick()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/assets/js/wikilink.js
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
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

## Communities (213 total, 61 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (59): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+51 more)

### Community 1 - "Slugify"
Cohesion: 0.08
Nodes (33): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), AbilityKeys(), AddFeat() (+25 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (149): keepAwake(), ValidKind(), ListCharacterImagesRow, CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewRow, CompendiumEntry, Encounter (+141 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.16
Nodes (26): writeFile(), ManifestKind(), readLimited(), ReadTakeHome(), plainTakeHome(), rewriteZip(), sampleBringBack(), sampleFallback() (+18 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "report"
Cohesion: 0.14
Nodes (16): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+8 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (117): Props, Size, Slot, Signed(), Button(), Variant, Variants(), Action() (+109 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - ".wizardView"
Cohesion: 0.11
Nodes (24): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, TestCoreTraits(), TestEquipmentOptions(), dataFor(), DecodeData() (+16 more)

### Community 9 - "Book"
Cohesion: 0.09
Nodes (35): rawBook, parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook() (+27 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (34): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+26 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Queries"
Cohesion: 0.04
Nodes (87): HomeManifest, TakeHomeBaseCharacter, TakeHomeCampaign, TakeHomeFallback, TakeHomeLevelUp, TakeHomeManifest, Character, GetNewCharacterByTokenRow (+79 more)

### Community 13 - "CampaignSources"
Cohesion: 0.08
Nodes (31): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), campaignSourceRow(), CampaignSources(), CampaignSourcesLive() (+23 more)

### Community 14 - "Store"
Cohesion: 0.06
Nodes (63): linkTarget, migration, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey() (+55 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (68): ListCharacterNamesRow, Visibility, internal/perm, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+60 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "blockMenu"
Cohesion: 0.09
Nodes (28): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListBacklinksRow, SearchReadableBlocksRow, CampaignRoster(), ctxCampaign(), EntityView (+20 more)

### Community 21 - "CompendiumPath"
Cohesion: 0.05
Nodes (61): CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, ListSharedWithRow, CompendiumSource (+53 more)

### Community 22 - "wikilink.js"
Cohesion: 0.17
Nodes (18): nodeActions(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages, LAYOUTS (+10 more)

### Community 23 - "who"
Cohesion: 0.05
Nodes (33): FiveTools2024Set(), TestFiveTools2024Set(), Campaign, GetBlockRow, browsePageSize(), compendiumBack(), compendiumPath(), ImportSummary() (+25 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (141): Props, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size (+133 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "CN"
Cohesion: 0.05
Nodes (89): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+81 more)

### Community 28 - "CombatantView"
Cohesion: 0.11
Nodes (35): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList() (+27 more)

### Community 29 - "TestBringBackAccept"
Cohesion: 0.14
Nodes (23): TakeHome, Character, TestNewCharacterTableSources(), testEnv, homeSheet(), TestBringBackAccept(), TestBringBackPlayerUpload(), TestBringBackReject() (+15 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (23): checkLink(), checkLinks(), clearDisplay(), clearSourceSelection(), commandChoices(), commandHint(), commandState(), createItem() (+15 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Player"
Cohesion: 0.07
Nodes (40): Player(), TestCanAddSources(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState() (+32 more)

### Community 34 - "Encounter"
Cohesion: 0.21
Nodes (8): CanAddSources(), notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "newEffectFix"
Cohesion: 0.20
Nodes (21): TestMonsterAtZeroHPUnaffected(), acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectFields() (+13 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (23): block, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "ItemData"
Cohesion: 0.08
Nodes (37): Props, ImageCategory, Defenses, Checkbox(), initialChecked(), chooseFirst(), firstValue(), missingDamageItems() (+29 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.07
Nodes (52): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+44 more)

### Community 42 - "WizardView"
Cohesion: 0.09
Nodes (46): Footer(), Label(), BuilderPickItems(), BuilderPicks(), BuilderPick, pickURL(), addedAtHome(), addedAtHomeIcon() (+38 more)

### Community 43 - "Client"
Cohesion: 0.27
Nodes (4): ClientInfo, goneDevice, TestFiveToolsClassLevelUp(), Hub

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseSheet"
Cohesion: 0.08
Nodes (50): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), homeTableSheet(), TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestMonsterImageLink() (+42 more)

### Community 46 - "Feature"
Cohesion: 0.08
Nodes (23): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), UsesLabel(), CustomMod (+15 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.11
Nodes (29): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+21 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (113): SearchCompendiumRow, UI conventions, Kbd(), Scripts(), Separator(), Toast(), Toaster(), toastIcon() (+105 more)

### Community 49 - "Path"
Cohesion: 0.03
Nodes (126): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), Gotchas, activeFight(), TestCombatantNamesFollowSource() (+118 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (142): TestNamesMatchGo(), Defense, TestDecode(), TestFiveToolsBooks(), feature(), RemoveSourceDir(), TestRemoveSourceDir(), serveOpen5e() (+134 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Spell"
Cohesion: 0.08
Nodes (28): FreeCast(), SwapFeatSpell(), LevelUpPlan, Spell, spellFromRow(), addClassSpells(), spellNeed(), TestSubclassPanels() (+20 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.07
Nodes (22): LevelUpChoice, armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorGrants() (+14 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "config.go"
Cohesion: 0.22
Nodes (12): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), Player(), lookupMap() (+4 more)

### Community 59 - "rosterCard"
Cohesion: 0.40
Nodes (8): deviceRow(), rosterCard(), devicesWhere(), agoTime(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Client"
Cohesion: 0.23
Nodes (3): message, Client, isDevice()

### Community 62 - ".Router"
Cohesion: 0.33
Nodes (3): RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 63 - "Part"
Cohesion: 0.05
Nodes (36): Part, Skill, findItem(), parseArmor(), parseItem(), TestCustomMods(), TestModifier(), TestResistances() (+28 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "App"
Cohesion: 0.03
Nodes (38): FromContext(), Entity, passTurn(), downloadNames(), entityTopics(), linkEntityLook(), collectHeld(), concentrationIDs() (+30 more)

### Community 66 - "handoff"
Cohesion: 0.28
Nodes (4): closeServers(), removeFile(), shutdownServers(), handoff

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.06
Nodes (56): Amount, CondOpKind, Effects, userText, ArmorKindNamed(), parseFeature(), parseProf(), parseProfSet() (+48 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - "Condition"
Cohesion: 0.07
Nodes (36): Combatant, Condition, CreateConditionParams, ListCombatantsRow, EffectEnd, combatantRow(), rollFor(), stepTurn() (+28 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "app"
Cohesion: 0.11
Nodes (17): env_file, environment, image, logging, network_mode, restart, user, volumes (+9 more)

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

### Community 85 - "Topic"
Cohesion: 0.15
Nodes (7): event, Handler, Topic, TopicKind, heldCollected(), noteTopics(), heldTopics

### Community 86 - "TestNewCharacterFromGM"
Cohesion: 0.20
Nodes (11): testEnv, mentionDoc(), newExtrasEnv(), TestBringBackAcceptExtras(), TestBringBackAcceptFiles(), TestBringBackStaleExtras(), TestNewCharacterFromGM(), TestNewCharacterFromVisitor() (+3 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".source"
Cohesion: 0.13
Nodes (16): UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceKeyWithSlash() (+8 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "Site"
Cohesion: 0.11
Nodes (20): Site, Source, sourceFormat, TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), formatOf(), Report, ImportProgress() (+12 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.07
Nodes (49): eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack(), TestHandoffExpiresAndReplaces() (+41 more)

### Community 103 - "Session"
Cohesion: 0.07
Nodes (30): fixture, tokens, loopbackHostOnly(), TestLoopbackHostOnly(), constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared() (+22 more)

### Community 105 - "connListener"
Cohesion: 0.25
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 106 - "NewClient"
Cohesion: 0.21
Nodes (21): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+13 more)

### Community 107 - "ftBook"
Cohesion: 0.11
Nodes (34): ftBook, ftRec, each(), ftCastingTime(), ftCell(), ftCopyTargets(), ftDecode(), ftFeatIncrease() (+26 more)

### Community 108 - "NameDesc"
Cohesion: 0.14
Nodes (23): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+15 more)

### Community 109 - "BlockEditor"
Cohesion: 0.22
Nodes (5): attachedPicture(), BlockEditor, json(), links(), sync()

### Community 111 - "Merge"
Cohesion: 0.16
Nodes (18): sheetMap, finish(), GroupDiff, GroupKeys(), GroupLabel(), Merge(), signature(), TakeBuild() (+10 more)

### Community 112 - "mirror"
Cohesion: 0.21
Nodes (11): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), replaceDir(), TestMirrorGetMissingFile(), DownloadOpen5e() (+3 more)

### Community 113 - "starterEnv"
Cohesion: 0.30
Nodes (10): testEnv, pickerTicks(), serveStarterSets(), starterEnv(), TestImportsTickFromCampaignOnly(), TestPartialDownloadTicks(), TestPickerKeepsJobTicks(), TestSourcesCardStarterSets() (+2 more)

### Community 116 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 117 - "vars"
Cohesion: 0.22
Nodes (8): env, CGO_ENABLED, GOFLAGS, vars, APP, EXE, PKG, version

### Community 118 - "ftItemData"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 119 - "testEnv"
Cohesion: 0.18
Nodes (9): TestSourceDeleteMany(), TestPlayerBlocksInGameSession(), assetETags(), immutableAsset(), TestEntityFieldLinks(), testEnv, TestAdminLogin(), TestCampaignJoinLinks() (+1 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Pending"
Cohesion: 0.15
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 127 - "task"
Cohesion: 0.22
Nodes (11): cmds, desc, env, cmds, desc, env, task, APP (+3 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "env"
Cohesion: 0.26
Nodes (15): BIN, DATA_DIR, MARKERS, MODE, NOREBUILD, PORT, PREPARE, RESTART_ARGS (+7 more)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery() (+16 more)

### Community 135 - "Queries"
Cohesion: 0.04
Nodes (68): HomebrewOwner, bindFiveToolsImages(), HomebrewInput, SaveHomebrew(), SeedHabitats(), targetColumns(), Queries, LinkTarget (+60 more)

### Community 136 - "alert_templ.go"
Cohesion: 0.05
Nodes (52): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Variant, Action() (+44 more)

### Community 137 - "cert.go"
Cohesion: 0.15
Nodes (4): generate(), serveSplit(), TestSplitTLS(), TestSplitTLSSilentConn()

### Community 139 - ".attachmentsDir"
Cohesion: 0.07
Nodes (14): ListAttachmentFilesRow, Attachment, firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify(), TestFolderName() (+6 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 143 - "run.sh"
Cohesion: 0.56
Nodes (8): alive(), lock(), run.sh script, start_server(), stop_server(), superseded(), unchanged(), unlock()

### Community 144 - "time.Time"
Cohesion: 0.03
Nodes (38): Identity, GetImageRow, ListImagesRow, ListHomeStartersRow, browserName(), formIntQuery(), TestUserAgent(), userAgentSummary() (+30 more)

### Community 145 - "normalize_test.go"
Cohesion: 0.33
Nodes (3): openBrowser(), normalize(), TestNormalize()

### Community 147 - "cert_test.go"
Cohesion: 0.29
Nodes (9): Fingerprint(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+1 more)

### Community 148 - "main"
Cohesion: 0.16
Nodes (14): die(), backup.sh script, DATA_DIR, lock(), common.sh script, build_and_check(), compose(), healthy() (+6 more)

### Community 157 - "State"
Cohesion: 0.20
Nodes (7): patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff()

### Community 175 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 177 - "InputForEntry"
Cohesion: 0.22
Nodes (8): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 179 - "dev:changed"
Cohesion: 0.40
Nodes (5): cmds, method, sources, exclude, dev:changed

### Community 180 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 181 - "run"
Cohesion: 0.25
Nodes (7): main(), parseFlags(), run(), TestParseFlags(), bookFolder(), CleanDownloads(), TestCleanDownloads()

### Community 182 - "systemd-units.sh"
Cohesion: 0.83
Nodes (3): die(), render(), systemd-units.sh script

### Community 184 - "ClassInfo"
Cohesion: 0.08
Nodes (37): ClassFeature, featureSwitch, Decode(), ClassData, SpellData, dnd5eFeatures(), TestSlotMaxima(), ClassFeature (+29 more)

### Community 186 - "internal/web"
Cohesion: 0.25
Nodes (7): Handlers, internal/web, Notes, TestKeepsBodyNewCharacter(), TestKeepsBodyBringBack(), keepsBody(), userError

### Community 188 - "dev:build"
Cohesion: 0.40
Nodes (5): cmds, deps, requires, vars, dev:build

### Community 189 - ".sourceIssues"
Cohesion: 0.36
Nodes (7): pinnedRefs(), sourceIssue(), featKey(), pinFeat(), SourceCopy, SourceIssue, sheetRef

### Community 191 - "before"
Cohesion: 0.40
Nodes (5): cmds, desc, vars, before, BEFORE_BIN

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

### Community 199 - "net.IP"
Cohesion: 0.52
Nodes (5): LANAddresses(), PreferDefaultRoute(), preferFirst(), TestPreferFirst(), handoffIPs()

### Community 200 - "tasks"
Cohesion: 0.11
Nodes (18): cmds, cmds, desc, cmds, desc, cmds, internal, cmds (+10 more)

### Community 201 - "Input"
Cohesion: 0.06
Nodes (94): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+86 more)

### Community 204 - "WeaponKind"
Cohesion: 0.40
Nodes (5): Completions(), parseMastery(), TestCompletions(), TestWeaponKind(), WeaponKind()

### Community 205 - "run:before"
Cohesion: 0.27
Nodes (6): cmds, desc, cmds, desc, run:before, run:test

### Community 206 - "ListSubclassesOfRow"
Cohesion: 0.40
Nodes (4): ListSubclassesOfRow, subclassIndex(), subclassOptions(), ListSubclassesOfParams

### Community 207 - "SourcePreview"
Cohesion: 0.50
Nodes (5): SwitchFeature, SwitchPlan, SourcePreview(), switchList(), SourcePreviewView

### Community 208 - "formErrors"
Cohesion: 0.40
Nodes (3): formErrors(), fieldError, pickError

### Community 209 - "dev:migrations"
Cohesion: 0.50
Nodes (4): cmds, method, sources, dev:migrations

### Community 210 - "run"
Cohesion: 0.67
Nodes (3): cmds, desc, run

### Community 211 - "run:test-player"
Cohesion: 0.67
Nodes (3): cmds, desc, run:test-player

## Knowledge Gaps
- **38 isolated node(s):** `npx`, `@playwright/mcp`, `DATA_DIR`, `gm-session-manager`, `testEnv` (+33 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 563 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **61 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Slugify`, `context.Context`, `Queries`, `.wizardView`, `Book`, `.attachmentsDir`, `Queries`, `CampaignSources`, `perm.go`, `time.Time`, `who`, `State`, `Encounter`, `Client`, `testing.T`, `Spell`, `ClassInfo`, `.sourceIssues`, `.Router`, `Part`, `handoff`, `Command`, `net.IP`, `Condition`, `Topic`, `Session`, `NameDesc`, `testEnv`, `Pending`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `DATA_DIR` to the rest of the system?**
  _38 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Slugify` be split into smaller, more focused modules?**
  _Cohesion score 0.08405797101449275 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `report`, `SheetView`, `Queries`, `alert_templ.go`, `Queries`, `CampaignSources`, `Store`, `perm.go`, `time.Time`, `blockMenu`, `CompendiumPath`, `who`, `CN`, `CombatantView`, `TestBringBackAccept`, `Player`, `newEffectFix`, `ItemData`, `tabs_templ.go`, `WizardView`, `ParseSheet`, `deathFix`, `github.com/a-h/templ.Component`, `testing.T`, `Spell`, `rosterCard`, `App`, `Condition`, `Input`, `SourcePreview`, `ParseWiki`, `TestNewCharacterFromGM`, `.source`, `newPlayerEnv`, `starterEnv`, `testEnv`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.0068328937904579235 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `internal/web`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.034) - this node is a cross-community bridge._