# Graph Report - gm-session-manager  (2026-11-13)

## Corpus Check
- 335 files · ~597,514 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4346 nodes · 20884 edges · 147 communities (113 shown, 34 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2077 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `b545c52e`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- newStore
- context.Context
- .do
- createProvider
- Input
- SheetView
- combobox.js
- redirect.go
- net/url.Values
- Store
- select.js
- command.go
- .Entry
- capReader
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
- newTestEnv
- NameDesc
- testing.T
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- ParseWikidot
- command.js
- dialog.js
- NewSheet
- formErrors
- tooltip.js
- tooltip_templ.go
- Feature
- Book
- export/markdown.go
- composite.js
- ListCombatantsRow
- ParseEffects
- CompendiumView
- popover_templ.go
- SaveHomebrew
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Weapon
- golangci-lint Config
- Spell
- github.com/a-h/templ.Component
- radiogroup.js
- vendorjs/main.go
- combobox_templ.go
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- who
- use_transition_status.js
- playwright
- release.sh
- dialog_templ.go
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
- .source
- Lint and Release Rules
- Run Script
- Sheet
- useTypeahead
- CompendiumKind
- Validator Rule
- Go Module
- ImagePage
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- appState
- newLiveFixture
- profFocus
- State
- .wizardView
- ftBook
- deviceLabel
- .sourceIssues
- viewerCtx
- config.go
- Client
- time.Time
- NewClient
- build.sh
- ListSharedWithRow
- looseString
- errBodyWriter
- .ListItemsByCategory
- names
- net/url.Values
- Before/after pairs
- .ListSpellsForClass
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- Pending
- .ListSubracesOf
- isToggle
- sourceUsedError
- Note
- toast_templ.go
- TestRunCommandUsesCurrentSheet
- TestLiveUpdatesAndKick

## God Nodes (most connected - your core abstractions)
1. `App` - 351 edges
2. `newTestEnv()` - 131 edges
3. `Sheet` - 122 edges
4. `Button()` - 122 edges
5. `CN()` - 113 edges
6. `SheetView` - 76 edges
7. `IfElse()` - 63 edges
8. `Command` - 59 edges
9. `CompendiumKind` - 56 edges
10. `ParseSheet()` - 47 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (147 total, 34 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "newStore"
Cohesion: 0.21
Nodes (18): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestSearch() (+10 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (74): keepAwake(), keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Encounter, SearchReadableBlocksRow (+66 more)

### Community 3 - ".do"
Cohesion: 0.12
Nodes (19): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+11 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Input"
Cohesion: 0.05
Nodes (69): ImageCategory, CloseFor(), Description(), Field(), Group(), Label(), Input(), Textarea() (+61 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (83): SpellData, ArmorEffect(), Signed(), Button(), Variant, Action(), Card(), Content() (+75 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "net/url.Values"
Cohesion: 0.11
Nodes (24): LevelUpPlan, SpellColumn(), SpellTotals(), spellNeed(), TestSpellNeed(), withSpells(), spellBoxes(), spellChecklist() (+16 more)

### Community 10 - "Store"
Cohesion: 0.06
Nodes (45): fixture, tokens, main(), run(), migration, isCleared(), newFixture(), serve() (+37 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.05
Nodes (75): Amount, CondOpKind, Effects, featureSwitch, HelpEntry, parser, Completions(), condCommand() (+67 more)

### Community 13 - ".Entry"
Cohesion: 0.07
Nodes (35): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, habitatsFrom(), FiveToolsImages() (+27 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (49): GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+41 more)

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
Nodes (40): Description(), Title(), Checkbox(), Label(), initialValue(), Item(), RadioGroup(), state() (+32 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (19): CanChangeSubclass(), CanCharacter(), homebrewTabURL(), entityTopics(), formIntQuery(), isHTMX(), redirect(), csvSource() (+11 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (148): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Variant (+140 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "newTestEnv"
Cohesion: 0.07
Nodes (59): TestRealSource(), Import(), ParseSheet(), TestConditionExpiresOnTurn(), TestDeleteRestoreImport(), TestRenameDeleteCampaign(), TestRunFiveToolsFinishes(), TestGMHomebrewClassAndWikidot() (+51 more)

### Community 28 - "NameDesc"
Cohesion: 0.12
Nodes (29): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, Slugify() (+21 more)

### Community 29 - "testing.T"
Cohesion: 0.03
Nodes (90): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestDecode(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), RemoveSourceDir() (+82 more)

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
Nodes (46): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Attachment, Condition, Entity, CreateCombatantParams (+38 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.12
Nodes (26): block, ClassFeature, ClassData, withCoreTraits(), pageSource(), PageSourceKey(), cellAt(), classTable() (+18 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "NewSheet"
Cohesion: 0.07
Nodes (34): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestLanguages() (+26 more)

### Community 38 - "formErrors"
Cohesion: 0.18
Nodes (8): Gotchas, Handlers, internal/web, Notes, TestPagesRender(), fieldError, pickError, userError

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "tooltip_templ.go"
Cohesion: 0.10
Nodes (30): Field, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks() (+22 more)

### Community 41 - "Feature"
Cohesion: 0.14
Nodes (7): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 42 - "Book"
Cohesion: 0.07
Nodes (44): bookWriter, mirror, rawBook, internal/compendium, parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+36 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ListCombatantsRow"
Cohesion: 0.16
Nodes (12): Combatant, ListCombatantsRow, rollFor(), stepTurn(), formInt(), combatantKind(), damageHP(), initiativeMod() (+4 more)

### Community 46 - "ParseEffects"
Cohesion: 0.11
Nodes (13): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+5 more)

### Community 47 - "CompendiumView"
Cohesion: 0.10
Nodes (27): CountCompendiumRow, CountHabitatsRow, CountSharedBySourceRow, SearchCompendiumRow, habitatCounts(), kindCounts(), pickerRow(), SearchComboItems() (+19 more)

### Community 48 - "popover_templ.go"
Cohesion: 0.15
Nodes (20): Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition(), conditionList() (+12 more)

### Community 49 - "SaveHomebrew"
Cohesion: 0.18
Nodes (8): Site, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SavePage(), SiteByKey(), TestParsePageURL()

### Community 50 - "cell"
Cohesion: 0.11
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
Cohesion: 0.13
Nodes (12): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+4 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Weapon"
Cohesion: 0.10
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.08
Nodes (26): rules, generated, paths, rules, formatters, enable, exclusions, settings (+18 more)

### Community 58 - "Spell"
Cohesion: 0.13
Nodes (16): TextData, FreeCast(), SwapFeatSpell(), Spell, textMarkdown(), addClassSpells(), applyPicks(), groupViews() (+8 more)

### Community 59 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (150): ListMonstersByNameRow, UI conventions, Badge(), Input(), Content(), Description(), Empty(), Header() (+142 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "combobox_templ.go"
Cohesion: 0.08
Nodes (46): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+38 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.09
Nodes (39): Defenses, damageChoice(), rosterCard(), capitalized(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+31 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "who"
Cohesion: 0.08
Nodes (18): ftsQuery(), Search(), Campaign, compendiumBack(), compendiumPath(), ImportSummary(), pathCampaign(), viewerCampaign() (+10 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "dialog_templ.go"
Cohesion: 0.10
Nodes (49): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+41 more)

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
Cohesion: 0.11
Nodes (38): Title(), Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+30 more)

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

### Community 83 - ".source"
Cohesion: 0.27
Nodes (8): testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceMove(), TestSourceOrder()

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "Sheet"
Cohesion: 0.08
Nodes (18): DeathSaves, TestCustomMods(), Sum(), TestGearCommands(), TestUpgradeSheet(), Coins, Item, Sheet (+10 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "CompendiumKind"
Cohesion: 0.03
Nodes (72): dataFor(), Decode(), DecodeData(), ValidKind(), SeedHabitats(), ListEntryCopiesRow, ListSubclassesOfRow, UpsertCompendiumEntryParams (+64 more)

### Community 96 - "ImagePage"
Cohesion: 0.09
Nodes (33): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent(), Props (+25 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (43): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+35 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "appState"
Cohesion: 0.32
Nodes (7): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpRedownloads(), appState, settings

### Community 103 - "newLiveFixture"
Cohesion: 0.12
Nodes (20): BenchmarkCompendiumList(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views, isRandomID(), shapeHash(), newLiveFixture() (+12 more)

### Community 105 - "State"
Cohesion: 0.09
Nodes (16): event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch, IsLoopback() (+8 more)

### Community 106 - ".wizardView"
Cohesion: 0.09
Nodes (30): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), AbilityKeys(), AddFeat(), SkillKeys() (+22 more)

### Community 107 - "ftBook"
Cohesion: 0.15
Nodes (22): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftCopyTargets(), ftDecode() (+14 more)

### Community 108 - "deviceLabel"
Cohesion: 0.50
Nodes (3): deviceLabel(), rosterHP(), Device

### Community 109 - ".sourceIssues"
Cohesion: 0.43
Nodes (6): pinnedRefs(), sourceIssue(), featKey(), SourceCopy, SourceIssue, sheetRef

### Community 111 - "viewerCtx"
Cohesion: 0.23
Nodes (6): notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView, viewerCtx

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 116 - "time.Time"
Cohesion: 0.04
Nodes (30): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), ListAttachmentFilesRow, Session (+22 more)

### Community 117 - "NewClient"
Cohesion: 0.28
Nodes (16): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+8 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 126 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 127 - "net/url.Values"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (23): Entry, fixture, fixtures, normalizer, o5eProperty, Source, sourceFormat, TestOpen5eMastery() (+15 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.22
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 147 - "Note"
Cohesion: 0.10
Nodes (13): Note, NoteKind, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks() (+5 more)

### Community 149 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 150 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 151 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 398 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **34 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Store`, `Pending`, `command.go`, `perm.go`, `Note`, `NameDesc`, `database/sql.NullInt64`, `ListCombatantsRow`, `CompendiumView`, `cell`, `Spell`, `FiveToolsJob`, `who`, `Sheet`, `CompendiumKind`, `appState`, `.wizardView`, `.sourceIssues`, `viewerCtx`, `Client`, `time.Time`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14090431125131442 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `CombatantView`, `Weapon`, `encoding/json.RawMessage`, `NewSheet`, `SheetView`, `Feature`, `.wizardView`, `net/url.Values`, `command.go`, `ListCombatantsRow`, `ParseEffects`, `.sourceIssues`, `TestRunCommandUsesCurrentSheet`, `CompendiumKind`, `Spell`, `newTestEnv`, `NameDesc`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010641151748189215 - nodes in this community are weakly interconnected._
- **Should `.do` be split into smaller, more focused modules?**
  _Cohesion score 0.11693548387096774 - nodes in this community are weakly interconnected._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._