# Graph Report - gm-session-manager  (2026-11-18)

## Corpus Check
- 352 files · ~617,285 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4517 nodes · 22642 edges · 150 communities (118 shown, 32 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2276 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `74feab3a`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Session
- context.Context
- Parse
- createProvider
- Slugify
- SheetView
- combobox.js
- redirect.go
- dialog_templ.go
- Import
- select.js
- Command
- .Entry
- capReader
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ftItemData
- WizardView
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- storeSource
- CompendiumKind
- State
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- ParseWikidot
- command.js
- dialog.js
- Book
- EffectMods
- tooltip.js
- Input
- .do
- Book
- export/markdown.go
- composite.js
- EquipmentOption
- Feature
- Open
- nullID
- Topic
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- github.com/a-h/templ.Component
- golangci-lint Config
- NewClient
- CompendiumPath
- radiogroup.js
- vendorjs/main.go
- toast_templ.go
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Weapon
- use_transition_status.js
- playwright
- release.sh
- NameDesc
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
- Client
- Lint and Release Rules
- Run Script
- tooltip_templ.go
- useTypeahead
- Handlers
- Validator Rule
- Go Module
- .sourceIssues
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- viewerCtx
- newLiveFixture
- profFocus
- EffectFieldsURL
- buildReview
- ftBook
- Slugify
- config.go
- .renderCompendium
- Client
- build.sh
- Queries
- looseString
- errBodyWriter
- ParseEffects
- net/url.Values
- Before/after pairs
- Pending
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- sourcesEnv
- popover_templ.go
- sourceUsedError
- ftRec
- spellCounts
- Condition
- Sheet
- .Router
- planSwitch
- TestRunCommandUsesCurrentSheet
- entityForm
- WeaponKind
- Campaign
- rowFigure
- SearchReadableBlocksRow

## God Nodes (most connected - your core abstractions)
1. `App` - 368 edges
2. `newTestEnv()` - 146 edges
3. `Sheet` - 130 edges
4. `Button()` - 122 edges
5. `CN()` - 113 edges
6. `SheetView` - 78 edges
7. `IfElse()` - 65 edges
8. `Command` - 59 edges
9. `CompendiumKind` - 56 edges
10. `Queries` - 53 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (150 total, 32 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "Session"
Cohesion: 0.16
Nodes (11): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+3 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (78): keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Encounter, Note, CreateCombatantParams (+70 more)

### Community 3 - "Parse"
Cohesion: 0.05
Nodes (44): Defense, feature(), Heal(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions() (+36 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Slugify"
Cohesion: 0.14
Nodes (16): LevelUpPlan, spellNeed(), groupViews(), TestClassOptionChoices(), withSpells(), LevelUpView, Option, SpellNeed (+8 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (112): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+104 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "dialog_templ.go"
Cohesion: 0.12
Nodes (37): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+29 more)

### Community 10 - "Import"
Cohesion: 0.11
Nodes (34): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+26 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.08
Nodes (48): Amount, CondOpKind, Effects, parseFeature(), parseProf(), parseProfSet(), parseSet(), parseSpell() (+40 more)

### Community 13 - ".Entry"
Cohesion: 0.07
Nodes (30): MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools() (+22 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (53): ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+45 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ftItemData"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 21 - "WizardView"
Cohesion: 0.08
Nodes (50): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Action(), Alert() (+42 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (39): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+31 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (31): SetCookie(), Identity, FromContext(), pathCampaign(), collectHeld(), formIntQuery(), isHTMX(), redirect() (+23 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (161): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+153 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "storeSource"
Cohesion: 0.11
Nodes (19): HomebrewOwner, Site, Source, sourceFormat, TestHomebrew(), TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), SaveHomebrew() (+11 more)

### Community 28 - "CompendiumKind"
Cohesion: 0.04
Nodes (50): CoreTraits(), dataFor(), DecodeData(), ValidKind(), ftsQuery(), HomebrewInput, Search(), CountCompendiumRow (+42 more)

### Community 29 - "State"
Cohesion: 0.11
Nodes (11): fakeHandler, patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff() (+3 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (49): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Attachment, Entity, entityTopics(), browserName() (+41 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.12
Nodes (25): block, ClassFeature, ClassData, dnd5eFeatures(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+17 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Book"
Cohesion: 0.11
Nodes (27): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestImportFiveToolsBook() (+19 more)

### Community 38 - "EffectMods"
Cohesion: 0.22
Nodes (8): EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Gotchas, internal/web, Notes, TestPagesRender()

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Input"
Cohesion: 0.05
Nodes (87): ImageCategory, ListBacklinksRow, Combobox(), Content(), Input(), Item(), List(), state() (+79 more)

### Community 41 - ".do"
Cohesion: 0.10
Nodes (22): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+14 more)

### Community 42 - "Book"
Cohesion: 0.13
Nodes (20): bookWriter, mirror, rawBook, internal/compendium, SpellData, bookFolder(), bookGroup(), downloadBook() (+12 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "EquipmentOption"
Cohesion: 0.29
Nodes (7): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), TestOptionText(), EquipmentChoice, optionText()

### Community 46 - "Feature"
Cohesion: 0.10
Nodes (15): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+7 more)

### Community 47 - "Open"
Cohesion: 0.08
Nodes (30): fixture, migration, newFixture(), TestOnlyHashIsStored(), apply(), loadMigrations(), Migrate(), Open() (+22 more)

### Community 48 - "nullID"
Cohesion: 0.15
Nodes (26): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), acBonusForm(), newEffectFix(), notes(), TestConcentrationCascade() (+18 more)

### Community 49 - "Topic"
Cohesion: 0.18
Nodes (6): event, Handler, Topic, TopicKind, heldCollected(), heldTopics

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (125): TestIsLoopback(), TestDecode(), TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), Verbs(), TestHelpVerbs(), TestMasteryLimit() (+117 more)

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

### Community 56 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (118): Input(), Content(), Description(), Empty(), Header(), Media(), Title(), Addon() (+110 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.07
Nodes (29): deny, files, rules, generated, paths, rules, formatters, enable (+21 more)

### Community 58 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 59 - "CompendiumPath"
Cohesion: 0.05
Nodes (73): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, Cell(), classBadges(), CompendiumPage() (+65 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.06
Nodes (27): keepAwake(), canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath() (+19 more)

### Community 62 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.12
Nodes (35): Props, Variant, Defenses, Badge(), badgeClasses(), capitalized(), beingBadge(), combatantChip() (+27 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Weapon"
Cohesion: 0.15
Nodes (7): Part, Skill, skillByKey(), TestDualWield(), damageLabel(), Weapon, attackTargets()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "NameDesc"
Cohesion: 0.20
Nodes (18): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+10 more)

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

### Community 86 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 96 - ".sourceIssues"
Cohesion: 0.43
Nodes (6): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.09
Nodes (36): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+28 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 103 - "newLiveFixture"
Cohesion: 0.09
Nodes (32): Player(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views, watchSheets() (+24 more)

### Community 106 - "buildReview"
Cohesion: 0.12
Nodes (23): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+15 more)

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (20): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftParts(), ftPrerequisite() (+12 more)

### Community 108 - "Slugify"
Cohesion: 0.09
Nodes (31): Cantrips(), SkillChoice(), TestSkillChoice(), Slugify(), AbilityKeys(), AddFeat(), AddResistance(), FreeCast() (+23 more)

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - ".renderCompendium"
Cohesion: 0.16
Nodes (9): ImportSummary(), fiveToolsKey(), writeBook(), sourceInUseMsg(), sourceUsers(), FiveToolsView, NewFiveToolsView(), FiveToolsBook (+1 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "Queries"
Cohesion: 0.04
Nodes (63): Decode(), TextData, ListSubclassesOfRow, Queries, Character, Env, ClassInfo, Character (+55 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.11
Nodes (12): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+4 more)

### Community 127 - "net/url.Values"
Cohesion: 0.09
Nodes (37): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+29 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 140 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 141 - "popover_templ.go"
Cohesion: 0.11
Nodes (28): UI conventions, Content(), Description(), Header(), Popover(), state(), Title(), Trigger() (+20 more)

### Community 146 - "ftRec"
Cohesion: 0.27
Nodes (8): ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve(), ftVersions()

### Community 147 - "spellCounts"
Cohesion: 0.67
Nodes (3): TestSpellCounts(), sortedColumns(), spellCounts()

### Community 150 - "Condition"
Cohesion: 0.05
Nodes (47): Combatant, Condition, CreateConditionParams, ListCombatantsRow, IsConditionSlug(), EffectEnd, passTurn(), rollFor() (+39 more)

### Community 156 - "Sheet"
Cohesion: 0.07
Nodes (17): DeathSaves, TestModifier(), Sum(), hpGain(), Coins, Item, Sheet, WornArmor (+9 more)

### Community 165 - ".Router"
Cohesion: 0.19
Nodes (7): IsLoopback(), ClearCookie(), writeCookie(), RequireAdmin(), sameOriginOnly(), useValidateTag(), CompendiumRest()

### Community 166 - "planSwitch"
Cohesion: 0.13
Nodes (13): featureSwitch, ftRaceAbility(), featureNames(), SourceRef, SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch() (+5 more)

### Community 172 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.48
Nodes (5): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits()

### Community 175 - "entityForm"
Cohesion: 0.10
Nodes (28): ListEntityFieldBacklinksRow, Entity, Field, Look, attrString(), generateSVG(), getIconContent(), Props (+20 more)

### Community 181 - "WeaponKind"
Cohesion: 0.50
Nodes (4): Completions(), parseMastery(), TestWeaponKind(), WeaponKind()

### Community 182 - "Campaign"
Cohesion: 0.18
Nodes (5): ListAttachmentFilesRow, Campaign, inCampaign(), CreateCampaignParams, CampaignSwitch

### Community 183 - "rowFigure"
Cohesion: 0.50
Nodes (4): TestRowFigure(), crText(), rowFigure(), textOnly()

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 401 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **32 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Pending`, `Slugify`, `Command`, `Viewer`, `Condition`, `CompendiumKind`, `State`, `database/sql.NullInt64`, `.Router`, `.do`, `Book`, `Open`, `cert.go`, `Campaign`, `FiveToolsJob`, `.sourceIssues`, `viewerCtx`, `Slugify`, `.renderCompendium`, `Client`, `Queries`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14539800995024876 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `.sourceIssues`, `Parse`, `encoding/json.RawMessage`, `Weapon`, `SheetView`, `planSwitch`, `NameDesc`, `Slugify`, `Input`, `Slugify`, `TestRunCommandUsesCurrentSheet`, `Feature`, `nullID`, `WeaponKind`, `Condition`, `Queries`, `CompendiumKind`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01029285099052541 - nodes in this community are weakly interconnected._
- **Why does `SubclassFeature` connect `Slugify` to `compendium_templ.go`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `Parse` be split into smaller, more focused modules?**
  _Cohesion score 0.053877551020408164 - nodes in this community are weakly interconnected._