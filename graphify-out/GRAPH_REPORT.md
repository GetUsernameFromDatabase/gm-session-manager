# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 372 files · ~641,477 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4679 nodes · 24166 edges · 145 communities (115 shown, 30 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2413 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `3b41873d`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Slugify
- context.Context
- Condition
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- ClassInfo
- Path
- encoding/json.RawMessage
- select.js
- Command
- ImportView
- Device
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ParseEffects
- IfElse
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- .compendiumPageSave
- redirect.go
- app.js
- tabbable.js
- useDismiss
- Client
- viewerCtx
- command.js
- dialog.js
- Open
- ParseWikidot
- tooltip.js
- CompendiumKind
- NameDesc
- config.go
- Session
- composite.js
- buildReview
- Feature
- planSwitch
- nullID
- Character
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- ftItemData
- shadcn-templ Config
- sourcesEnv
- golangci-lint Config
- Import
- Sheet
- radiogroup.js
- vendorjs/main.go
- CN
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Queries
- use_transition_status.js
- playwright
- testEnv
- internal/compendium
- checkbox.js
- ParseAidedd
- switch.js
- .sourceIssues
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- .editEntity
- useTypeahead
- DownloadFiveTools
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- buildReview
- mounted
- .sendFeedback
- Player
- profFocus
- CompendiumPath
- time.Time
- ftBook
- NewClient
- BlockEditor
- joinNonEmpty
- Input
- csvSource
- .features
- export/markdown.go
- build.sh
- TestLevelUpPushesTracker
- looseString
- errBodyWriter
- formErrors
- Before/after pairs
- Hub
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ASIError
- Book
- .Entry
- sourceUsedError
- Pending
- github.com/a-h/templ.Component
- InputForEntry
- ParseSheet

## God Nodes (most connected - your core abstractions)
1. `App` - 374 edges
2. `Path()` - 265 edges
3. `newTestEnv()` - 168 edges
4. `Sheet` - 131 edges
5. `Button()` - 124 edges
6. `CN()` - 114 edges
7. `SheetView` - 78 edges
8. `IfElse()` - 65 edges
9. `Command` - 60 edges
10. `CompendiumKind` - 56 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (145 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "Slugify"
Cohesion: 0.08
Nodes (32): Cantrips(), TextData, Slugify(), AddFeat(), AddResistance(), FreeCast(), SkillKeys(), SwapFeatSpell() (+24 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (80): keepAwake(), Manager, ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Note (+72 more)

### Community 3 - "Condition"
Cohesion: 0.03
Nodes (62): ListHomebrewByOwnerRow, ListHomebrewRow, ListImagesRow, Condition, Encounter, CreateConditionParams, ListCombatantsRow, EffectEnd (+54 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.09
Nodes (38): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+30 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (129): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+121 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ClassInfo"
Cohesion: 0.10
Nodes (28): Decode(), TestSlotMaxima(), ClassInfo, ClassLevel, CasterLevel(), SlotMaxima(), SpellColumn(), SpellTotals() (+20 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (109): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+101 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (32): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+24 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (61): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+53 more)

### Community 13 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnOverview(), ImportReview(), importRows(), importStep(), keepSection(), columnLabel(), ImportCategory (+2 more)

### Community 14 - "Device"
Cohesion: 0.39
Nodes (5): devicesWhere(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.06
Nodes (49): ListCharacterNamesRow, GetBlockRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanDeleteCharacter() (+41 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ParseEffects"
Cohesion: 0.07
Nodes (19): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+11 more)

### Community 21 - "IfElse"
Cohesion: 0.07
Nodes (61): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+53 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (39): SetCookie(), writeCookie(), GetImageRow, compendiumBack(), compendiumPath(), ImportSummary(), pathCampaign(), viewerCampaign() (+31 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (135): Props, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size (+127 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.11
Nodes (17): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+9 more)

### Community 28 - ".compendiumPageSave"
Cohesion: 0.13
Nodes (14): Site, Source, TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page (+6 more)

### Community 29 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Client"
Cohesion: 0.09
Nodes (17): countingHandler, event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch (+9 more)

### Community 34 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Open"
Cohesion: 0.07
Nodes (35): fixture, tokens, migration, isCleared(), newFixture(), serve(), TestMiddlewareIdentity(), TestOnlyHashIsStored() (+27 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.04
Nodes (45): CoreTraits(), dataFor(), DecodeData(), ValidKind(), CountCompendiumRow, UpsertCompendiumEntryParams, CompendiumEntry, CompendiumKind (+37 more)

### Community 41 - "NameDesc"
Cohesion: 0.12
Nodes (29): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+21 more)

### Community 42 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 43 - "Session"
Cohesion: 0.23
Nodes (5): IsLoopback(), ClearCookie(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 46 - "Feature"
Cohesion: 0.08
Nodes (23): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), UsesLabel(), CustomMod (+15 more)

### Community 47 - "planSwitch"
Cohesion: 0.15
Nodes (11): featureSwitch, featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch(), sameSlug() (+3 more)

### Community 48 - "nullID"
Cohesion: 0.14
Nodes (27): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead(), acBonusForm() (+19 more)

### Community 49 - "Character"
Cohesion: 0.06
Nodes (19): ListSubclassesOfRow, Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), CanChangeSubclass() (+11 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (116): Defense, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestIsLoopback(), TestDecode(), TestCoreTraits(), RemoveSourceDir() (+108 more)

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

### Community 56 - "sourcesEnv"
Cohesion: 0.28
Nodes (7): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Import"
Cohesion: 0.07
Nodes (46): sourceFormat, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+38 more)

### Community 59 - "Sheet"
Cohesion: 0.05
Nodes (28): Part, Skill, findItem(), parseItem(), TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums() (+20 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.20
Nodes (16): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+8 more)

### Community 62 - "CN"
Cohesion: 0.05
Nodes (88): Props, Size, Props, ctxKey, ctxState, DialogProps, EmptyProps, GroupProps (+80 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (36): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+28 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Queries"
Cohesion: 0.07
Nodes (30): HomebrewOwner, HomebrewInput, SaveHomebrew(), SeedHabitats(), Queries, Combatant, Character, passTurn() (+22 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "testEnv"
Cohesion: 0.11
Nodes (14): testEnv, assetETags(), immutableAsset(), TestEntityFieldLinks(), testEnv, TestAdminLogin(), TestCampaignJoinLinks(), TestLogout() (+6 more)

### Community 72 - "internal/compendium"
Cohesion: 0.09
Nodes (20): EquipmentItem, internal/compendium, EquipmentOptions(), EquipmentOption, SkillChoice(), TestEquipmentOptions(), TestSkillChoice(), SpellData (+12 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - ".sourceIssues"
Cohesion: 0.36
Nodes (6): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

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
Cohesion: 0.15
Nodes (11): attachedPicture(), ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), nodeActions() (+3 more)

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - ".editEntity"
Cohesion: 0.07
Nodes (20): Entity, LinkTarget, Fields, ParseFields(), entityTopics(), linkEntityLook(), saveFieldLinks(), lookFree() (+12 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (25): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestImportFiveToolsBook() (+17 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "buildReview"
Cohesion: 0.31
Nodes (5): ensureCategory(), buildReview(), formValues(), csvImport, csvImports

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 103 - "Player"
Cohesion: 0.15
Nodes (22): Player(), TestCanRemoveCondition(), BenchmarkEncounterState(), BenchmarkSheetState(), Websocket and live views, isRandomID(), shapeHash(), newLiveFixture() (+14 more)

### Community 105 - "CompendiumPath"
Cohesion: 0.05
Nodes (77): CountHabitatsRow, ListEntryCopiesRow, SearchCompendiumRow, Habitat, itemDetails(), classBadges(), CompendiumPage(), CompendiumResults() (+69 more)

### Community 106 - "time.Time"
Cohesion: 0.05
Nodes (28): Identity, FromContext(), ListAttachmentFilesRow, Attachment, Campaign, browserName(), firstLine(), mdLinkText() (+20 more)

### Community 107 - "ftBook"
Cohesion: 0.17
Nodes (15): ftBook, ftRec, each(), ftCopyTargets(), ftDecode(), ftRecords(), TestFiveToolsCopy(), ftMod() (+7 more)

### Community 108 - "NewClient"
Cohesion: 0.21
Nodes (19): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign() (+11 more)

### Community 111 - "joinNonEmpty"
Cohesion: 0.13
Nodes (25): ftDistance, ftAmount(), ftCastingTime(), ftCell(), ftDuration(), ftParts(), ftPrerequisite(), ftRange() (+17 more)

### Community 112 - "Input"
Cohesion: 0.05
Nodes (104): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+96 more)

### Community 116 - ".features"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 117 - "export/markdown.go"
Cohesion: 0.13
Nodes (19): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+11 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "TestLevelUpPushesTracker"
Cohesion: 0.31
Nodes (8): TestLevelUpPushesTracker(), dial(), patchOf(), readUntil(), rosterRefetch(), TestDeviceSeenOnLeave(), TestLiveUpdatesAndKick(), wsMessage

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 127 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "Hub"
Cohesion: 0.20
Nodes (3): ClientInfo, goneDevice, Hub

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (18): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+10 more)

### Community 139 - "Book"
Cohesion: 0.12
Nodes (24): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+16 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 143 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 146 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (116): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Field, Look, UI conventions, Kbd() (+108 more)

### Community 148 - "InputForEntry"
Cohesion: 0.29
Nodes (7): habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 159 - "ParseSheet"
Cohesion: 0.10
Nodes (39): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+31 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 413 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Slugify`, `context.Context`, `Condition`, `Hub`, `ClassInfo`, `Book`, `Command`, `Pending`, `Viewer`, `Client`, `viewerCtx`, `CompendiumKind`, `Session`, `Character`, `testing.T`, `FiveToolsJob`, `Queries`, `testEnv`, `internal/compendium`, `.sourceIssues`, `.editEntity`, `buildReview`, `time.Time`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1397743454923869 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Slugify`, `Condition`, `tabs_templ.go`, `SheetView`, `ImportView`, `Viewer`, `github.com/a-h/templ.Component`, `IfElse`, `App`, `ParseSheet`, `CompendiumKind`, `nullID`, `Character`, `testing.T`, `sourcesEnv`, `Import`, `CN`, `CombatantView`, `Queries`, `testEnv`, `ParseWiki`, `.editEntity`, `buildReview`, `.sendFeedback`, `Player`, `CompendiumPath`, `time.Time`, `Input`, `export/markdown.go`, `TestLevelUpPushesTracker`?**
  _High betweenness centrality (0.038) - this node is a cross-community bridge._
- **Should `Slugify` be split into smaller, more focused modules?**
  _Cohesion score 0.07740112994350283 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Import`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._