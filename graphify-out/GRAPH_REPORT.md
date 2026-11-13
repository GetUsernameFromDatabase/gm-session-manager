# Graph Report - gm-session-manager  (2026-11-13)

## Corpus Check
- 317 files · ~595,085 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4313 nodes · 19762 edges · 149 communities (119 shown, 30 thin omitted)
- Extraction: 91% EXTRACTED · 9% INFERRED · 0% AMBIGUOUS · INFERRED: 1851 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `eaa017f6`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- .do
- createProvider
- ImportView
- SheetView
- combobox.js
- redirect.go
- subclassPanels
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
- ParseSheet
- NameDesc
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- Source
- command.js
- dialog.js
- Book
- formErrors
- tooltip.js
- .Router
- Hub
- Book
- export/markdown.go
- composite.js
- .levelUp
- ParseEffects
- compendium_templ.go
- .fail
- SaveHomebrew
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Sheet
- golangci-lint Config
- Slugify
- sourcesEnv
- radiogroup.js
- vendorjs/main.go
- dialog_templ.go
- FiveToolsJob
- sqlc Config
- select_templ.go
- Browser check
- Claude Settings
- .compendiumDir
- use_transition_status.js
- playwright
- release.sh
- Campaign
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
- Feature
- Lint and Release Rules
- Run Script
- spellCounts
- useTypeahead
- selectItems
- Validator Rule
- Go Module
- CharacterFile
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- github.com/a-h/templ.Component
- newLiveFixture
- .equipmentChoices
- State
- CompendiumKind
- ftBook
- tooltip_templ.go
- .combatantHP
- Import
- viewerCtx
- config.go
- Client
- rowFigure
- .ListItemsByCategory
- time.Time
- NewClient
- build.sh
- .sourceIssues
- looseString
- errBodyWriter
- ClassInfo
- planSwitch
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
- TestRunCommandUsesCurrentSheet
- toast_templ.go
- RosterCharacter
- ftDuration
- SearchReadableBlocksRow
- NewSheet

## God Nodes (most connected - your core abstractions)
1. `App` - 349 edges
2. `newTestEnv()` - 131 edges
3. `Button()` - 122 edges
4. `Sheet` - 121 edges
5. `CN()` - 113 edges
6. `IfElse()` - 62 edges
7. `Command` - 59 edges
8. `SheetView` - 57 edges
9. `CompendiumKind` - 54 edges
10. `ParseSheet()` - 47 edges

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

## Communities (149 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.13
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.11
Nodes (20): Condition, Description(), Header(), state(), Title(), triggerID(), combatantDefenses(), petrified() (+12 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (79): keepAwake(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, CompendiumSource, Encounter (+71 more)

### Community 3 - ".do"
Cohesion: 0.12
Nodes (20): main(), isCleared(), testEnv, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows(), importStatusItems() (+2 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (116): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+108 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "subclassPanels"
Cohesion: 0.18
Nodes (9): ListSubclassesOfRow, TestSubclassPanels(), subclassIndex(), subclassOptions(), subclassPanels(), SubclassPanel, ListSubclassesOfParams, SubclassChangeView (+1 more)

### Community 10 - "Store"
Cohesion: 0.06
Nodes (41): fixture, tokens, migration, constantTimeEqual(), credentialsMatch(), newFixture(), serve(), TestCredentialsMatch() (+33 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.06
Nodes (73): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+65 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (23): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+15 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (55): GetBlockRow, ListCombatantsRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+47 more)

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
Nodes (10): ArmorData, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText() (+2 more)

### Community 21 - "levelUpForm"
Cohesion: 0.09
Nodes (47): Props, Footer(), Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup() (+39 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (26): SetCookie(), writeCookie(), FromContext(), compendiumBack(), compendiumPath(), homebrewTabURL(), pathCampaign(), entityTopics() (+18 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (174): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+166 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.11
Nodes (34): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+26 more)

### Community 28 - "NameDesc"
Cohesion: 0.20
Nodes (18): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+10 more)

### Community 29 - "newTestEnv"
Cohesion: 0.04
Nodes (101): TestDecode(), TestTokensReadAsLabels(), RemoveSourceDir(), TestRemoveSourceDir(), TestMonsterFormRoundTrip(), TestMasteryLimit(), TestDropFeat(), TestSheetPinsSource() (+93 more)

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
Nodes (42): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Attachment, Entity, CreateCombatantParams, EntityKind (+34 more)

### Community 34 - "Source"
Cohesion: 0.14
Nodes (18): Source, sourceFormat, CoreTraits(), TestCoreTraits(), storeSource(), upsert(), pageSource(), PageSourceKey() (+10 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Book"
Cohesion: 0.07
Nodes (44): rawBook, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants() (+36 more)

### Community 38 - "formErrors"
Cohesion: 0.29
Nodes (5): Handlers, formErrors(), fieldError, pickError, userError

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".Router"
Cohesion: 0.24
Nodes (4): ClearCookie(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 42 - "Book"
Cohesion: 0.25
Nodes (8): bookWriter, mirror, bookFolder(), downloadBook(), fetchJSON(), filterDoc(), TestMirrorGetMissingFile(), fetch()

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".levelUp"
Cohesion: 0.10
Nodes (22): dnd5eFeatures(), TestFiveToolsClassLevelUp(), FreeCast(), SwapFeatSpell(), ClassFeature, LevelUpPlan, Spell, spellFromRow() (+14 more)

### Community 46 - "ParseEffects"
Cohesion: 0.11
Nodes (12): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+4 more)

### Community 47 - "compendium_templ.go"
Cohesion: 0.05
Nodes (82): Props, Variant, SpellData, CountHabitatsRow, UI conventions, Badge(), badgeClasses(), IfElse() (+74 more)

### Community 48 - ".fail"
Cohesion: 0.06
Nodes (18): dataFor(), DecodeData(), Character, Character, CanCharacter(), checkSkillPick(), classSkillPick(), postedSkills() (+10 more)

### Community 49 - "SaveHomebrew"
Cohesion: 0.11
Nodes (14): HomebrewOwner, Site, TestHomebrew(), fiveToolsPage(), HomebrewInput, InputForEntry(), SaveHomebrew(), SeedHabitats() (+6 more)

### Community 50 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

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
Cohesion: 0.11
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Sheet"
Cohesion: 0.05
Nodes (28): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+20 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.08
Nodes (26): rules, generated, paths, rules, formatters, enable, exclusions, settings (+18 more)

### Community 58 - "Slugify"
Cohesion: 0.07
Nodes (38): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), AbilityKeys(), AddFeat() (+30 more)

### Community 59 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 62 - "dialog_templ.go"
Cohesion: 0.08
Nodes (60): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+52 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "select_templ.go"
Cohesion: 0.05
Nodes (63): Description(), Field(), Label(), Content(), Group(), initialValue(), Item(), itemLabel() (+55 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".compendiumDir"
Cohesion: 0.19
Nodes (6): ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), FiveToolsView, NewFiveToolsView()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Campaign"
Cohesion: 0.28
Nodes (6): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks()

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.10
Nodes (33): block, ClassFeature, abilityMod(), ParseAidedd(), statLabel(), textLines(), ClassData, firstInt() (+25 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "tabs_templ.go"
Cohesion: 0.09
Nodes (39): Title(), Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+31 more)

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
Cohesion: 0.18
Nodes (8): FeatureEffects, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), sortedColumns(), spellCounts(), TestSpellCounts()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "selectItems"
Cohesion: 0.07
Nodes (45): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent() (+37 more)

### Community 96 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (47): ftClass, ftCopyRef, ftEntry, ftFeat, ftItem, ftModOp, ftMonster, ftNamed (+39 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (104): ListMonstersByNameRow, Content(), Description(), Empty(), Header(), Media(), Title(), Input() (+96 more)

### Community 103 - "newLiveFixture"
Cohesion: 0.11
Nodes (23): Player(), BenchmarkEncounterState(), BenchmarkSheetState(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture() (+15 more)

### Community 104 - ".equipmentChoices"
Cohesion: 0.29
Nodes (7): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), TestOptionText(), EquipmentChoice, optionText()

### Community 105 - "State"
Cohesion: 0.09
Nodes (16): event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch, Topic (+8 more)

### Community 106 - "CompendiumKind"
Cohesion: 0.04
Nodes (44): ValidKind(), ftsQuery(), Search(), CountCompendiumRow, ListEntryCopiesRow, SearchCompendiumRow, UpsertCompendiumEntryParams, CompendiumEntry (+36 more)

### Community 107 - "ftBook"
Cohesion: 0.17
Nodes (18): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftMod(), ftName() (+10 more)

### Community 108 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 109 - ".combatantHP"
Cohesion: 0.36
Nodes (4): Combatant, damageHP(), setHP(), hpChange

### Community 110 - "Import"
Cohesion: 0.11
Nodes (32): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+24 more)

### Community 111 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "Client"
Cohesion: 0.22
Nodes (3): IsLoopback(), TestIsLoopback(), Client

### Community 114 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 116 - "time.Time"
Cohesion: 0.05
Nodes (22): Identity, ListAttachmentFilesRow, Campaign, browserName(), firstLine(), mdLinkText(), inCampaign(), TestUserAgent() (+14 more)

### Community 117 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".sourceIssues"
Cohesion: 0.36
Nodes (7): pinnedRefs(), sourceIssue(), featKey(), pinFeat(), SourceCopy, SourceIssue, sheetRef

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.31
Nodes (4): errorText(), TestErrorText(), namesOnly(), errBodyWriter

### Community 125 - "ClassInfo"
Cohesion: 0.13
Nodes (24): Decode(), Queries, ClassInfo, ClassLevel, SourceRef, SpellTotals(), uniqueName(), preparedCaps() (+16 more)

### Community 126 - "planSwitch"
Cohesion: 0.14
Nodes (12): featureSwitch, ftRaceAbility(), featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch() (+4 more)

### Community 127 - "net/url.Values"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Pending"
Cohesion: 0.17
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 143 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 145 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 150 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 151 - "ftDuration"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 163 - "NewSheet"
Cohesion: 0.08
Nodes (30): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestDefenses(), TestLanguages(), TestLevelUp() (+22 more)

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 392 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Store`, `Pending`, `command.go`, `perm.go`, `database/sql.NullInt64`, `Book`, `.Router`, `Hub`, `.levelUp`, `.fail`, `SaveHomebrew`, `cert.go`, `Sheet`, `Slugify`, `FiveToolsJob`, `.compendiumDir`, `Campaign`, `State`, `CompendiumKind`, `.combatantHP`, `viewerCtx`, `time.Time`, `.sourceIssues`, `ClassInfo`?**
  _High betweenness centrality (0.060) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13348837209302325 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `NewSheet`, `encoding/json.RawMessage`, `planSwitch`, `SheetView`, `command.go`, `.levelUp`, `TestRunCommandUsesCurrentSheet`, `.fail`, `perm.go`, `Feature`, `.sourceIssues`, `Slugify`, `ParseSheet`, `NameDesc`, `ClassInfo`, `dialog_templ.go`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `popover_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10869565217391304 - nodes in this community are weakly interconnected._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01069332753749185 - nodes in this community are weakly interconnected._
- **Should `.do` be split into smaller, more focused modules?**
  _Cohesion score 0.12043010752688173 - nodes in this community are weakly interconnected._