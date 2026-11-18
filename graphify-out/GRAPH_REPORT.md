# Graph Report - gm-session-manager  (2026-11-18)

## Corpus Check
- 354 files · ~621,646 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4547 nodes · 22858 edges · 148 communities (114 shown, 34 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2308 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `a93cc7c7`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- github.com/a-h/templ.Component
- createProvider
- github.com/a-h/templ.Component
- SheetView
- combobox.js
- redirect.go
- Input
- encoding/json.RawMessage
- select.js
- Command
- Site
- sourcesEnv
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Weapon
- WizardView
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Session
- Sheet
- State
- app.js
- tabbable.js
- useDismiss
- CompendiumKind
- ParseWikidot
- command.js
- dialog.js
- Book
- Import
- tooltip.js
- Hub
- .do
- Book
- Handlers
- composite.js
- ParseSheet
- Feature
- Store
- nullID
- .sendFeedback
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Layout
- golangci-lint Config
- Client
- Item
- radiogroup.js
- storeSource
- capReader
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- .Router
- use_transition_status.js
- playwright
- release.sh
- parseHP
- checkbox.js
- ParseAidedd
- switch.js
- tabs_templ.go
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- keepawake_windows.go
- Lint and Release Rules
- Run Script
- NewClient
- useTypeahead
- vendorjs/main.go
- Validator Rule
- Go Module
- .builderPicks
- live.js
- internal_backdrop.js
- LinkTarget
- mounted
- .compendiumDir
- newLiveFixture
- profFocus
- EffectFieldsURL
- buildReview
- ftBook
- ftRec
- Condition
- WeaponKind
- config.go
- github.com/a-h/templ.Component
- spellCounts
- export/markdown.go
- build.sh
- planSwitch
- looseString
- errBodyWriter
- ParseEffects
- .ListBlockStepsSince
- .ListItemsByCategory
- Before/after pairs
- viewerCtx
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftItemData
- ASIError
- .ListSpellsForClass
- .Entry
- .ListSubracesOf
- sourceUsedError
- Parse
- RosterCharacter
- Pending
- time.Time
- Weapon

## God Nodes (most connected - your core abstractions)
1. `App` - 368 edges
2. `newTestEnv()` - 148 edges
3. `Sheet` - 130 edges
4. `Button()` - 122 edges
5. `CN()` - 113 edges
6. `SheetView` - 78 edges
7. `IfElse()` - 65 edges
8. `Command` - 60 edges
9. `CompendiumKind` - 56 edges
10. `Queries` - 53 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
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

## Communities (148 total, 34 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.12
Nodes (24): Content(), Description(), Header(), state(), Title(), Trigger(), triggerID(), modTargetLabel() (+16 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (107): keepAwake(), ListImageCharactersRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListSharedWithRow, ListImageEntitiesRow, ListImagesRow (+99 more)

### Community 3 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (87): Field, Group(), Textarea(), Scripts(), Separator(), Spinner(), Textarea(), Toast() (+79 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "github.com/a-h/templ.Component"
Cohesion: 0.21
Nodes (28): ctxKey, ctxState, Props, Dialog(), Close(), Content(), Description(), Dialog() (+20 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (123): ArmorEffect(), Coins, Slot, Signed(), Button(), Variant, Action(), Card() (+115 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Input"
Cohesion: 0.05
Nodes (77): Defenses, UI conventions, Combobox(), Content(), Item(), List(), Description(), Field() (+69 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (39): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+31 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.08
Nodes (42): CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet(), parseSpell() (+34 more)

### Community 13 - "Site"
Cohesion: 0.18
Nodes (8): Site, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SavePage(), SiteByKey(), TestParsePageURL()

### Community 14 - "sourcesEnv"
Cohesion: 0.28
Nodes (7): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.06
Nodes (46): ListCharacterNamesRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass() (+38 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Weapon"
Cohesion: 0.17
Nodes (10): ItemData, TestCustomMods(), TestExplainSums(), Sum(), TestDualWield(), damageLabel(), Weapon, compendiumWeapon() (+2 more)

### Community 21 - "WizardView"
Cohesion: 0.06
Nodes (62): Props, LevelUpPlan, Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup() (+54 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (39): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+31 more)

### Community 23 - "App"
Cohesion: 0.03
Nodes (43): SetCookie(), writeCookie(), Identity, FromContext(), GetImageRow, Campaign, Character, Entity (+35 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (191): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+183 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Session"
Cohesion: 0.11
Nodes (15): tokens, constantTimeEqual(), credentialsMatch(), isCleared(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity(), TestRequireAdmin() (+7 more)

### Community 28 - "Sheet"
Cohesion: 0.04
Nodes (93): internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice(), TestOptions() (+85 more)

### Community 29 - "State"
Cohesion: 0.09
Nodes (14): event, fakeHandler, Handler, patchMessage, stateHandler, viewPatch, Topic, TopicKind (+6 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "CompendiumKind"
Cohesion: 0.04
Nodes (52): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), dataFor(), DecodeData(), ValidKind() (+44 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Book"
Cohesion: 0.14
Nodes (19): readBook(), DownloadFiveTools(), RemoveFiveTools(), readDoc(), serveMirror(), TestDownloadFiveTools(), TestDownloadFiveToolsFailureLeavesNothing(), TestDownloadFiveToolsReplacesFolder() (+11 more)

### Community 38 - "Import"
Cohesion: 0.07
Nodes (41): sourceFormat, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+33 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 41 - ".do"
Cohesion: 0.12
Nodes (21): main(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm() (+13 more)

### Community 42 - "Book"
Cohesion: 0.11
Nodes (25): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+17 more)

### Community 43 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseSheet"
Cohesion: 0.13
Nodes (30): ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage() (+22 more)

### Community 46 - "Feature"
Cohesion: 0.10
Nodes (15): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+7 more)

### Community 47 - "Store"
Cohesion: 0.08
Nodes (33): fixture, migration, newFixture(), TestOnlyHashIsStored(), apply(), loadMigrations(), Migrate(), Open() (+25 more)

### Community 48 - "nullID"
Cohesion: 0.10
Nodes (34): DeathSaves, nullID(), deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP() (+26 more)

### Community 49 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (132): TestIsLoopback(), TestDecode(), TestCoreTraits(), TestSkillChoice(), RemoveSourceDir(), TestRemoveSourceDir(), Verbs(), TestHelpVerbs() (+124 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "cert.go"
Cohesion: 0.10
Nodes (18): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+10 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Layout"
Cohesion: 0.06
Nodes (70): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListMonstersByNameRow, SearchCompendiumRow, Habitat, Cell(), Row() (+62 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.07
Nodes (29): deny, files, rules, generated, paths, rules, formatters, enable (+21 more)

### Community 59 - "Item"
Cohesion: 0.31
Nodes (4): findItem(), TestGearCommands(), ensureItem(), Item

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "storeSource"
Cohesion: 0.14
Nodes (17): HomebrewOwner, Source, TestStoreSourceRollsBackOnAfterSave(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), TestFiveToolsCopyTargets(), TestFiveToolsEdition() (+9 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (35): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+27 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".Router"
Cohesion: 0.23
Nodes (5): IsLoopback(), ClearCookie(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "parseHP"
Cohesion: 0.29
Nodes (8): Amount, CritDamage(), Damage(), parseDamage(), parseHP(), Set(), parseAmount(), TestApplyHPDeath()

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "tabs_templ.go"
Cohesion: 0.09
Nodes (39): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+31 more)

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
Cohesion: 0.20
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "NewClient"
Cohesion: 0.28
Nodes (15): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+7 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 96 - ".builderPicks"
Cohesion: 0.24
Nodes (7): Completions(), BuilderPickItems(), BuilderPicks(), closeAfter(), BuilderPick, pickURL(), TestPickRowActs()

### Community 98 - "live.js"
Cohesion: 0.50
Nodes (3): bindings(), roots(), within()

### Community 100 - "LinkTarget"
Cohesion: 0.17
Nodes (8): LinkTarget, AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink, snapshotBody

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".compendiumDir"
Cohesion: 0.21
Nodes (5): ImportSummary(), fiveToolsKey(), sourceUsers(), FiveToolsView, NewFiveToolsView()

### Community 103 - "newLiveFixture"
Cohesion: 0.15
Nodes (21): Player(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views, isRandomID(), shapeHash(), newLiveFixture() (+13 more)

### Community 105 - "EffectFieldsURL"
Cohesion: 0.33
Nodes (4): EffectFieldsURL(), EffectID(), effectQuery(), EffectTarget

### Community 106 - "buildReview"
Cohesion: 0.09
Nodes (24): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+16 more)

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (16): ftBook, each(), ftCastingTime(), ftParts(), ftPrerequisite(), ftRecords(), TestFiveToolsCopy(), crProficiency() (+8 more)

### Community 108 - "ftRec"
Cohesion: 0.27
Nodes (8): ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve(), ftVersions()

### Community 109 - "Condition"
Cohesion: 0.04
Nodes (66): Queries, Combatant, Condition, CreateConditionParams, ListCombatantsRow, EffectMods(), EndsConcentration(), TestEndsConcentration() (+58 more)

### Community 111 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (84): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListEntityFieldBacklinksRow (+76 more)

### Community 116 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

### Community 117 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "planSwitch"
Cohesion: 0.24
Nodes (6): featureSwitch, Env, featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.10
Nodes (13): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+5 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "viewerCtx"
Cohesion: 0.22
Nodes (7): notFoundIsNil(), parseLiveKey(), TestLiveAffects(), liveKey, liveKind, liveView, viewerCtx

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (23): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable() (+15 more)

### Community 137 - "ftItemData"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 140 - ".Entry"
Cohesion: 0.08
Nodes (30): MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools() (+22 more)

### Community 146 - "Parse"
Cohesion: 0.05
Nodes (49): Defense, ftRaceAbility(), Heal(), Parse(), tokenize(), TestApplyHPMonsterUnaffected(), TestDeadRefusals(), TestDeathSaveFailKills() (+41 more)

### Community 148 - "RosterCharacter"
Cohesion: 0.48
Nodes (6): rosterCard(), Roster, deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 149 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 150 - "time.Time"
Cohesion: 0.06
Nodes (20): ListAttachmentFilesRow, Attachment, browserName(), firstLine(), mdLinkText(), saveScreenshot(), inCampaign(), TestUserAgent() (+12 more)

### Community 156 - "Weapon"
Cohesion: 0.11
Nodes (6): Part, Skill, TestModifier(), skillByKey(), Modifier(), ProficiencyAt()

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 402 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **34 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `viewerCtx`, `Viewer`, `Pending`, `time.Time`, `Sheet`, `State`, `CompendiumKind`, `Hub`, `.do`, `Book`, `Store`, `testing.T`, `cert.go`, `FiveToolsJob`, `.Router`, `.builderPicks`, `.compendiumDir`, `Condition`?**
  _High betweenness centrality (0.056) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1464365353888294 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `SheetView`, `Input`, `encoding/json.RawMessage`, `Parse`, `Weapon`, `WizardView`, `Weapon`, `CompendiumKind`, `ParseSheet`, `Feature`, `nullID`, `testing.T`, `Item`, `.builderPicks`, `Condition`, `WeaponKind`, `planSwitch`, `ParseEffects`?**
  _High betweenness centrality (0.014) - this node is a cross-community bridge._
- **Should `popover_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1164021164021164 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `Sheet` to `compendium_templ.go`, `WizardView`, `ParseEffects`, `planSwitch`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009306490123724767 - nodes in this community are weakly interconnected._