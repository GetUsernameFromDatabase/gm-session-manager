# Graph Report - gm-session-manager  (2026-11-13)

## Corpus Check
- 314 files · ~593,397 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4286 nodes · 19639 edges · 144 communities (115 shown, 29 thin omitted)
- Extraction: 91% EXTRACTED · 9% INFERRED · 0% AMBIGUOUS · INFERRED: 1820 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `5c64ce45`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- dialog_templ.go
- context.Context
- .do
- createProvider
- Weapon
- SheetView
- combobox.js
- redirect.go
- ImagePage
- Store
- select.js
- Command
- .Entry
- time.Time
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ftItemData
- levelUpForm
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- Spell
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- CommandBuilder
- command.js
- dialog.js
- DownloadFiveTools
- Slugify
- tooltip.js
- SearchCompendiumRow
- ParseWikidot
- Book
- export/markdown.go
- composite.js
- sourcesEnv
- ClassInfo
- github.com/a-h/templ.Component
- Character
- State
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Sheet
- golangci-lint Config
- Open
- capReader
- radiogroup.js
- vendorjs/main.go
- Input
- FiveToolsJob
- sqlc Config
- select_templ.go
- Browser check
- Claude Settings
- .compendiumDir
- use_transition_status.js
- playwright
- release.sh
- Session
- checkbox.js
- ParseAidedd
- switch.js
- IfElse
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- Feature
- Lint and Release Rules
- Run Script
- Client
- useTypeahead
- dialog_templ.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- ImportView
- newLiveFixture
- .featChoices
- Client
- .sendFeedback
- ftBook
- popover_templ.go
- empty_templ.go
- formErrors
- viewerCtx
- config.go
- LinkTarget
- Pending
- Note
- .Router
- TestRunCommandUsesCurrentSheet
- build.sh
- .sourceIssues
- looseString
- errBodyWriter
- Entity
- names
- rowFigure
- Before/after pairs
- SearchReadableBlocksRow
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- ClientInfo
- ListSharedWithRow
- isToggle
- sourceUsedError
- Parse

## God Nodes (most connected - your core abstractions)
1. `App` - 346 edges
2. `newTestEnv()` - 130 edges
3. `Button()` - 122 edges
4. `Sheet` - 121 edges
5. `CN()` - 113 edges
6. `IfElse()` - 62 edges
7. `Command` - 59 edges
8. `SheetView` - 57 edges
9. `CompendiumKind` - 54 edges
10. `Input()` - 47 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `UI conventions` --references--> `removeButton()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (144 total, 29 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "dialog_templ.go"
Cohesion: 0.08
Nodes (44): Props, Variant, Defenses, damageChoice(), Badge(), badgeClasses(), Label(), capitalized() (+36 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (114): keepAwake(), Site, Decode(), ValidKind(), SeedHabitats(), FetchPage(), ParsePageURL(), SiteByKey() (+106 more)

### Community 3 - ".do"
Cohesion: 0.15
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.12
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (91): Slot, Signed(), Button(), Variant, Action(), Card(), Content(), Description() (+83 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "ImagePage"
Cohesion: 0.10
Nodes (34): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, attrString(), generateSVG(), getIconContent(), Props, Icon() (+26 more)

### Community 10 - "Store"
Cohesion: 0.11
Nodes (35): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+27 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (71): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+63 more)

### Community 13 - ".Entry"
Cohesion: 0.07
Nodes (33): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, habitatsFrom(), TestHabitatsFrom() (+25 more)

### Community 14 - "time.Time"
Cohesion: 0.05
Nodes (26): ListAttachmentFilesRow, ListImagesRow, Attachment, browserName(), firstLine(), folderName(), mdLinkText(), saveScreenshot() (+18 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (49): GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass() (+41 more)

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

### Community 21 - "levelUpForm"
Cohesion: 0.12
Nodes (37): initialValue(), Item(), RadioGroup(), state(), feedbackKind(), spellChecklist(), splitCantrips(), abilityForm() (+29 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (28): SetCookie(), Identity, FromContext(), Campaign, compendiumBack(), compendiumPath(), homebrewTabURL(), pathCampaign() (+20 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (169): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+161 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.13
Nodes (31): ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpTwice() (+23 more)

### Community 28 - "Spell"
Cohesion: 0.10
Nodes (24): FreeCast(), SwapFeatSpell(), LevelUpPlan, Spell, addClassSpells(), spellNeed(), TestSpellNeed(), LevelUpView (+16 more)

### Community 29 - "newTestEnv"
Cohesion: 0.04
Nodes (103): TestDecode(), TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), TestMasteryLimit(), TestDropFeat(), TestSheetPinsSource(), TestCanCharacter() (+95 more)

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
Nodes (45): GetImageRow, Combatant, Condition, Encounter, CreateCombatantParams, ListCombatantsRow, EncounterStatus, applyCondOps() (+37 more)

### Community 34 - "CommandBuilder"
Cohesion: 0.11
Nodes (26): Field, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), campaignSourceRow(), CampaignSourcesView (+18 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "DownloadFiveTools"
Cohesion: 0.11
Nodes (24): internal/compendium, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants() (+16 more)

### Community 38 - "Slugify"
Cohesion: 0.11
Nodes (31): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, Slugify() (+23 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "SearchCompendiumRow"
Cohesion: 0.15
Nodes (8): ftsQuery(), Search(), CountCompendiumRow, SearchCompendiumRow, CompendiumBase(), CompendiumURL(), CopyKey(), SearchCompendiumParams

### Community 41 - "ParseWikidot"
Cohesion: 0.11
Nodes (28): block, ClassFeature, ClassData, dnd5eFeatures(), TestFiveToolsClassLevelUp(), withCoreTraits(), cellAt(), classTable() (+20 more)

### Community 42 - "Book"
Cohesion: 0.12
Nodes (22): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+14 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.15
Nodes (17): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+9 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 46 - "ClassInfo"
Cohesion: 0.05
Nodes (33): featureSwitch, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+25 more)

### Community 47 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (115): SpellData, UI conventions, Footer(), Checkbox(), Input(), Kbd(), Label(), Scripts() (+107 more)

### Community 48 - "Character"
Cohesion: 0.05
Nodes (30): CoreTraits(), dataFor(), DecodeData(), ListSubclassesOfRow, Character, Character, Campaign(), CharacterFileName() (+22 more)

### Community 49 - "State"
Cohesion: 0.12
Nodes (12): Handler, patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff(), TestPatch() (+4 more)

### Community 50 - "cell"
Cohesion: 0.12
Nodes (23): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+15 more)

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

### Community 56 - "Sheet"
Cohesion: 0.08
Nodes (13): DeathSaves, ArmorEffect(), Sum(), Coins, Item, Sheet, WornArmor, classPin() (+5 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.10
Nodes (21): generated, paths, rules, formatters, enable, exclusions, settings, match-constant (+13 more)

### Community 58 - "Open"
Cohesion: 0.12
Nodes (19): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+11 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Input"
Cohesion: 0.06
Nodes (49): CloseFor(), Description(), Field(), Group(), armorForm(), builderFields(), builderForm(), BuilderPickItems() (+41 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "select_templ.go"
Cohesion: 0.11
Nodes (25): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+17 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".compendiumDir"
Cohesion: 0.18
Nodes (7): FiveToolsBooks(), ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), FiveToolsView, NewFiveToolsView()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Session"
Cohesion: 0.10
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+10 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "IfElse"
Cohesion: 0.06
Nodes (67): CountHabitatsRow, CountSharedBySourceRow, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant() (+59 more)

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
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "Feature"
Cohesion: 0.19
Nodes (8): FeatureEffects, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "Client"
Cohesion: 0.11
Nodes (9): event, fakeHandler, message, stateHandler, IsLoopback(), TestIsLoopback(), Client, Topic (+1 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "dialog_templ.go"
Cohesion: 0.12
Nodes (42): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+34 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (38): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+30 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows(), importStatusItems() (+2 more)

### Community 103 - "newLiveFixture"
Cohesion: 0.09
Nodes (24): Player(), BenchmarkCompendiumList(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas, internal/web, Notes (+16 more)

### Community 104 - ".featChoices"
Cohesion: 0.12
Nodes (21): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), TestEquipmentOptions(), TestSkillChoice(), TextData, AbilityKeys() (+13 more)

### Community 105 - "Client"
Cohesion: 0.18
Nodes (14): Hub, New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+6 more)

### Community 106 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (23): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftMod(), ftName() (+15 more)

### Community 108 - "popover_templ.go"
Cohesion: 0.15
Nodes (19): Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition(), conditionLabel() (+11 more)

### Community 109 - "empty_templ.go"
Cohesion: 0.07
Nodes (54): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListMonstersByNameRow (+46 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): Handlers, fieldError, pickError, userError

### Community 111 - "viewerCtx"
Cohesion: 0.24
Nodes (6): notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView, viewerCtx

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "LinkTarget"
Cohesion: 0.17
Nodes (8): LinkTarget, AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink, snapshotBody

### Community 114 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 115 - "Note"
Cohesion: 0.20
Nodes (6): Note, NoteKind, CreateNoteParams, ListReadableNotesParams, RenameNoteParams, noteForm

### Community 116 - ".Router"
Cohesion: 0.25
Nodes (5): ClearCookie(), writeCookie(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 117 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".sourceIssues"
Cohesion: 0.27
Nodes (9): masteryLimit(), pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), featKey(), pinFeat() (+1 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Entity"
Cohesion: 0.13
Nodes (9): Entity, EntityKind, entityTopics(), linkEntityLook(), CreateEntityParams, SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams (+1 more)

### Community 126 - "names"
Cohesion: 0.40
Nodes (5): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), SourcePreviewView

### Community 127 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.09
Nodes (30): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, Source, TestOpen5eMastery() (+22 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 163 - "Parse"
Cohesion: 0.06
Nodes (38): Defense, Parse(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+30 more)

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 391 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **29 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Command`, `time.Time`, `perm.go`, `Spell`, `newTestEnv`, `database/sql.NullInt64`, `Slugify`, `SearchCompendiumRow`, `Book`, `Character`, `State`, `cert.go`, `Sheet`, `Open`, `FiveToolsJob`, `.compendiumDir`, `.featChoices`, `Client`, `viewerCtx`, `Pending`, `.Router`, `.sourceIssues`, `Entity`?**
  _High betweenness centrality (0.047) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13506428942670554 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `dialog_templ.go`, `database/sql.NullInt64`, `Parse`, `encoding/json.RawMessage`, `Weapon`, `Slugify`, `SheetView`, `.featChoices`, `context.Context`, `Command`, `ClassInfo`, `Character`, `Feature`, `TestRunCommandUsesCurrentSheet`, `.sourceIssues`, `ParseSheet`, `Spell`?**
  _High betweenness centrality (0.015) - this node is a cross-community bridge._
- **Should `dialog_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.08090957165520889 - nodes in this community are weakly interconnected._
- **Why does `Client` connect `Client` to `compendium_templ.go`, `redirect.go`, `Client`, `time.Time`, `perm.go`, `State`, `Pending`, `FiveToolsJob`?**
  _High betweenness centrality (0.013) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010382081686429512 - nodes in this community are weakly interconnected._