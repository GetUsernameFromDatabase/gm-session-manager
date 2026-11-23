# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 372 files · ~640,679 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4675 nodes · 24148 edges · 152 communities (121 shown, 31 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2414 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `89dc5b6d`
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
- Queries
- Path
- encoding/json.RawMessage
- select.js
- Command
- ItemData
- Device
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ParseEffects
- WizardView
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
- Session
- ParseWikidot
- tooltip.js
- CompendiumKind
- NameDesc
- config.go
- Session
- composite.js
- buildReview
- Feature
- names
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
- ParseSheet
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- popover_templ.go
- use_transition_status.js
- playwright
- select_templ.go
- EquipmentOption
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
- CharacterFile
- useTypeahead
- DownloadFiveTools
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- database/sql.NullInt64
- mounted
- .sendFeedback
- Player
- profFocus
- CompendiumPath
- .compendiumDir
- ftBook
- NewClient
- BlockEditor
- Look
- Input
- csvSource
- ListSharedWithRow
- export/markdown.go
- build.sh
- SearchReadableBlocksRow
- looseString
- errBodyWriter
- Open
- Attachment
- Before/after pairs
- Hub
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- .features
- ASIError
- Book
- .Entry
- report
- sourceUsedError
- Pending
- internal/web
- TestSourceSwitchKeepsEdits
- github.com/a-h/templ.Component
- Campaign
- saveScreenshot
- folderName
- Option

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

## Communities (152 total, 31 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "Slugify"
Cohesion: 0.09
Nodes (28): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), Slugify(), AbilityKeys(), AddFeat(), FreeCast() (+20 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (85): keepAwake(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListImageEntitiesRow, ListAttachmentFilesRow, CompendiumSource, Encounter (+77 more)

### Community 3 - "Condition"
Cohesion: 0.04
Nodes (56): Defense, Combatant, Condition, CreateConditionParams, ListCombatantsRow, Defenses, ParseDefenses(), EffectEnd (+48 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.10
Nodes (37): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+29 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (131): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+123 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Queries"
Cohesion: 0.09
Nodes (32): featureSwitch, Decode(), Queries, Env, ClassInfo, ClassLevel, WornArmor, featureNames() (+24 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (115): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), Gotchas, activeFight(), TestCombatantNamesFollowSource() (+107 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (47): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+39 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (60): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+52 more)

### Community 13 - "ItemData"
Cohesion: 0.08
Nodes (31): ImageCategory, firstValue(), modTargetLabel(), usesItems(), weaponAbilities(), EffectFieldsForm(), effectModFields(), EffectModRow() (+23 more)

### Community 14 - "Device"
Cohesion: 0.39
Nodes (5): devicesWhere(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.06
Nodes (52): ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+44 more)

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
Cohesion: 0.09
Nodes (15): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 21 - "WizardView"
Cohesion: 0.08
Nodes (53): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, UI conventions, Action() (+45 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (35): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Campaign, pathCampaign(), entityTopics() (+27 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (176): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+168 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.10
Nodes (19): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+11 more)

### Community 28 - ".compendiumPageSave"
Cohesion: 0.12
Nodes (17): HomebrewOwner, Site, Source, TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), SaveHomebrew(), SeedHabitats(), storeSource() (+9 more)

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
Cohesion: 0.20
Nodes (8): notFoundIsNil(), parseLiveKey(), Roster, WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Session"
Cohesion: 0.07
Nodes (25): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+17 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.04
Nodes (39): CoreTraits(), dataFor(), DecodeData(), ValidKind(), UpsertCompendiumEntryParams, CompendiumEntry, CompendiumKind, compendiumBack() (+31 more)

### Community 41 - "NameDesc"
Cohesion: 0.12
Nodes (28): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+20 more)

### Community 42 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 43 - "Session"
Cohesion: 0.24
Nodes (4): IsLoopback(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.13
Nodes (22): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+14 more)

### Community 46 - "Feature"
Cohesion: 0.08
Nodes (23): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), UsesLabel(), CustomMod (+15 more)

### Community 47 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 48 - "nullID"
Cohesion: 0.14
Nodes (27): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead(), acBonusForm() (+19 more)

### Community 49 - "Character"
Cohesion: 0.05
Nodes (25): TextData, Character, Character, NewCharacterFile(), CanCharacter(), linkEntityLook(), linkCharacterLook(), lookFree() (+17 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (107): TestIsLoopback(), TestDecode(), TestCoreTraits(), feature(), TestFiveToolsClasses(), RemoveSourceDir(), TestRemoveSourceDir(), TestPageClientRedirect() (+99 more)

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
Nodes (45): sourceFormat, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+37 more)

### Community 59 - "Sheet"
Cohesion: 0.05
Nodes (26): Part, Skill, AddResistance(), TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+18 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "ParseSheet"
Cohesion: 0.10
Nodes (41): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+33 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.31
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (38): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+30 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "select_templ.go"
Cohesion: 0.11
Nodes (23): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state() (+15 more)

### Community 72 - "EquipmentOption"
Cohesion: 0.29
Nodes (7): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), applyEquipment(), EquipmentChoice, gear

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
Cohesion: 0.43
Nodes (6): pinnedRefs(), sourceIssue(), featKey(), SourceCopy, SourceIssue, sheetRef

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

### Community 86 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "DownloadFiveTools"
Cohesion: 0.16
Nodes (20): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook() (+12 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (52): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Entity, Note, EntityKind, NoteKind (+44 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 103 - "Player"
Cohesion: 0.09
Nodes (32): Player(), TestCanRemoveCondition(), BenchmarkSheetState(), testEnv, Websocket and live views, TestNoteViewAudiences(), isRandomID(), shapeHash() (+24 more)

### Community 105 - "CompendiumPath"
Cohesion: 0.05
Nodes (77): SpellData, CountCompendiumRow, CountHabitatsRow, ListEntryCopiesRow, SearchCompendiumRow, Habitat, itemDetails(), classBadges() (+69 more)

### Community 106 - ".compendiumDir"
Cohesion: 0.18
Nodes (8): ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), FiveToolsView, NewFiveToolsView(), FiveToolsBook, FiveToolsGroup

### Community 107 - "ftBook"
Cohesion: 0.13
Nodes (23): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftCopyTargets(), ftDecode() (+15 more)

### Community 108 - "NewClient"
Cohesion: 0.21
Nodes (19): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign() (+11 more)

### Community 111 - "Look"
Cohesion: 0.33
Nodes (9): Entity, Look, initials(), lookControls(), looksGrid(), npcPortrait(), currentLook(), EntityListView (+1 more)

### Community 112 - "Input"
Cohesion: 0.06
Nodes (84): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+76 more)

### Community 117 - "export/markdown.go"
Cohesion: 0.35
Nodes (8): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain()

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Open"
Cohesion: 0.14
Nodes (21): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+13 more)

### Community 127 - "Attachment"
Cohesion: 0.16
Nodes (4): Attachment, firstLine(), mdLinkText(), CreateAttachmentParams

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "Hub"
Cohesion: 0.20
Nodes (3): ClientInfo, goneDevice, Hub

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - ".features"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Book"
Cohesion: 0.11
Nodes (25): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+17 more)

### Community 140 - ".Entry"
Cohesion: 0.10
Nodes (26): MonsterAction, addEntry(), setStat(), MonsterData, TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor(), InputForEntry() (+18 more)

### Community 141 - "report"
Cohesion: 0.24
Nodes (8): mdCell(), oneLine(), parseConsole(), TestMarkdownCell(), consoleEntry, consoleLog, feedbackForm, report

### Community 143 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 144 - "internal/web"
Cohesion: 0.22
Nodes (6): Handlers, internal/web, keepsBody(), fieldError, pickError, userError

### Community 145 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 146 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (121): ListEntityFieldBacklinksRow, ListBacklinksRow, Field, Content(), Description(), Empty(), Header(), Media() (+113 more)

### Community 147 - "Campaign"
Cohesion: 0.40
Nodes (4): Campaign(), CharacterFileName(), fileSlug(), writeFile()

### Community 149 - "saveScreenshot"
Cohesion: 0.40
Nodes (4): saveScreenshot(), parseCSV(), TestGuessMapping(), saveSniffed()

### Community 150 - "folderName"
Cohesion: 0.67
Nodes (3): folderName(), slugify(), TestFolderName()

### Community 159 - "Option"
Cohesion: 0.07
Nodes (28): dnd5eFeatures(), TestFiveToolsClassLevelUp(), ListSubclassesOfRow, ClassFeature, LevelUpPlan, Spellbook(), SpellColumn(), TestSpellbook() (+20 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 413 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **31 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Slugify`, `context.Context`, `Condition`, `Hub`, `Queries`, `Book`, `Pending`, `Viewer`, `cert.go`, `Client`, `viewerCtx`, `Session`, `CompendiumKind`, `Session`, `Character`, `testing.T`, `Sheet`, `FiveToolsJob`, `.sourceIssues`, `database/sql.NullInt64`, `CompendiumPath`, `.compendiumDir`, `Attachment`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13971175062203653 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `context.Context`, `Condition`, `tabs_templ.go`, `SheetView`, `ItemData`, `Viewer`, `TestSourceSwitchKeepsEdits`, `github.com/a-h/templ.Component`, `Campaign`, `WizardView`, `App`, `github.com/a-h/templ.Attributes`, `Session`, `CompendiumKind`, `nullID`, `Character`, `testing.T`, `sourcesEnv`, `Import`, `ParseSheet`, `CombatantView`, `ParseWiki`, `database/sql.NullInt64`, `.sendFeedback`, `Player`, `CompendiumPath`, `Look`, `Input`, `export/markdown.go`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **Should `Slugify` be split into smaller, more focused modules?**
  _Cohesion score 0.08953900709219859 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `internal/web`, `Path`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._