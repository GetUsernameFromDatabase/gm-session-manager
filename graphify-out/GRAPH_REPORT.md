# Graph Report - gm-session-manager  (2026-11-24)

## Corpus Check
- 380 files · ~654,790 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4823 nodes · 24812 edges · 182 communities (122 shown, 60 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2489 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `ca1a3e59`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- toast_templ.go
- context.Context
- vendorjs/main.go
- createProvider
- DownloadFiveTools
- SheetView
- combobox.js
- ParseSheet
- newTestEnv
- encoding/json.RawMessage
- select.js
- Command
- dialog_templ.go
- redirect.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- WizardView
- wikilink.js
- nullID
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- viewerCtx
- sourcesEnv
- IfElse
- app.js
- tabbable.js
- useDismiss
- Sheet
- Sum
- command.js
- dialog.js
- time.Time
- ParseWikidot
- tooltip.js
- CompendiumKind
- tabs_templ.go
- tooltip_templ.go
- Session
- composite.js
- ParseArmorTraining
- CustomMod
- names
- github.com/a-h/templ.Component
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Item
- shadcn-templ Config
- Session
- golangci-lint Config
- popover_templ.go
- Part
- radiogroup.js
- Player
- ParseEffects
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Client
- use_transition_status.js
- playwright
- ftPrerequisite
- newStore
- checkbox.js
- ParseAidedd
- switch.js
- Visibility
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- ftRec
- useTypeahead
- NewClient
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- storeSource
- mounted
- Hub
- nullID
- profFocus
- cert.go
- ftBook
- Pending
- BlockEditor
- export/markdown.go
- deathFix
- capReader
- report
- .importParam
- build.sh
- .saveReport
- looseString
- errBodyWriter
- Topic
- buildReview
- Before/after pairs
- .sendFeedback
- jsconfig.json
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- WeaponKind
- ASIError
- Book
- .Entry
- .ListItemsByCategory
- sourceUsedError
- Handlers
- CampaignSources
- .ListSpellsForClass
- .q
- importCategories
- io.Reader
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- parseMod
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
- planSwitch
- spellCounts
- .ListImageLookAccess
- ListSharedWithRow
- .ListSubracesOf
- .ListBlockStepsSince
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 376 edges
2. `Path()` - 272 edges
3. `newTestEnv()` - 177 edges
4. `Sheet` - 134 edges
5. `Button()` - 123 edges
6. `CN()` - 114 edges
7. `SheetView` - 79 edges
8. `IfElse()` - 65 edges
9. `Command` - 60 edges
10. `Queries` - 57 edges

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

## Communities (182 total, 60 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.13
Nodes (59): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+51 more)

### Community 1 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (121): keepAwake(), MoveSource(), SetSourceOrder(), ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow (+113 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (24): EquipmentOptions(), TestEquipmentOptions(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants() (+16 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (87): Coins, Slot, Signed(), Variant, Action(), Card(), Content(), Description() (+79 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.08
Nodes (48): Import(), TestUpgradeSheet(), ParseSheet(), TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestMonsterImageLink(), TestBuilderCompendium(), TestChangeSubclass() (+40 more)

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (106): ParsePageURL(), TestParsePageURL(), TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), Gotchas (+98 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (46): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+38 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.08
Nodes (45): Amount, CondOpKind, Effects, userText, parseSet(), parseSpell(), parseCoin(), parseInspiration() (+37 more)

### Community 13 - "dialog_templ.go"
Cohesion: 0.12
Nodes (43): ctxKey, ctxState, Props, Button(), Dialog(), Close(), CloseFor(), Content() (+35 more)

### Community 14 - "redirect.go"
Cohesion: 0.25
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (46): ListCharacterNamesRow, GetBlockRow, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanChangeSubclass(), CanCharacter() (+38 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "App"
Cohesion: 0.04
Nodes (36): SetCookie(), writeCookie(), FromContext(), ftsQuery(), Search(), Attachment, Campaign, compendiumBack() (+28 more)

### Community 21 - "WizardView"
Cohesion: 0.08
Nodes (43): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+35 more)

### Community 22 - "wikilink.js"
Cohesion: 0.14
Nodes (17): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+9 more)

### Community 23 - "nullID"
Cohesion: 0.06
Nodes (41): Combatant, ListCombatantsRow, Character, combatantRow(), passTurn(), rollFor(), stepTurn(), concentrationIDs() (+33 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (176): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+168 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 28 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 29 - "IfElse"
Cohesion: 0.04
Nodes (116): ListImageCharactersRow, ListImageEntitiesRow, ImageCategory, Field, BenchmarkImportReview(), UI conventions, Checkbox(), Combobox() (+108 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Sheet"
Cohesion: 0.04
Nodes (107): LevelUpChoice, internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice() (+99 more)

### Community 34 - "Sum"
Cohesion: 0.19
Nodes (4): TestCustomMods(), TestExplainSums(), Sum(), TestWithEffects()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "time.Time"
Cohesion: 0.04
Nodes (36): ListEntityFieldBacklinksRow, Entity, ListBacklinksRow, LinkTarget, Entity, Fields, ParseFields(), entityTopics() (+28 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.11
Nodes (27): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+19 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (57): EquipmentItem, CoreTraits(), EquipmentOption, dataFor(), DecodeData(), ValidKind(), CountCompendiumRow, UpsertCompendiumEntryParams (+49 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.09
Nodes (40): Separator(), Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+32 more)

### Community 42 - "tooltip_templ.go"
Cohesion: 0.08
Nodes (40): SearchCompendiumRow, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks() (+32 more)

### Community 43 - "Session"
Cohesion: 0.15
Nodes (12): fixture, tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken() (+4 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseArmorTraining"
Cohesion: 0.13
Nodes (13): armorGrants(), ArmorKindNamed(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack() (+5 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (19): FeatureEffects, parseFeature(), parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature (+11 more)

### Community 47 - "names"
Cohesion: 0.40
Nodes (5): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), SourcePreviewView

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.03
Nodes (174): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, SearchReadableBlocksRow, Habitat, Content() (+166 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (127): TestIsLoopback(), TestDecode(), TestCoreTraits(), TestSkillChoice(), RemoveSourceDir(), TestRemoveSourceDir(), TestPageClientRedirect(), TestArmorTrainingGrantsAndRemoval() (+119 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Item"
Cohesion: 0.19
Nodes (9): findItem(), parseArmor(), parseItem(), ArmorEffect(), TestGearCommands(), ensureItem(), Item, WornArmor (+1 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Session"
Cohesion: 0.16
Nodes (8): IsLoopback(), ClearCookie(), Identity, RequireAdmin(), deviceKey(), sessionKey(), sameOriginOnly(), useValidateTag()

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "popover_templ.go"
Cohesion: 0.09
Nodes (33): Look, Content(), Description(), Header(), Popover(), state(), Title(), Trigger() (+25 more)

### Community 59 - "Part"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.07
Nodes (41): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Handlers (+33 more)

### Community 62 - "ParseEffects"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

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

### Community 68 - "Client"
Cohesion: 0.11
Nodes (10): countingHandler, fakeHandler, Handler, message, patchMessage, viewPatch, Client, Diff() (+2 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.13
Nodes (15): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+7 more)

### Community 72 - "newStore"
Cohesion: 0.20
Nodes (18): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+10 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Visibility"
Cohesion: 0.18
Nodes (12): Visibility, addBlockForm(), NoteBlocks(), NotePage(), visibilityFields(), NoteView, visibilityLabel(), ListBlockDocsRow (+4 more)

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
Nodes (12): attachedPicture(), ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, nodeActions(), openWeb(), randomID() (+4 more)

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - "ftRec"
Cohesion: 0.27
Nodes (8): ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve(), ftVersions()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "NewClient"
Cohesion: 0.18
Nodes (19): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDisconnectDropsRequest() (+11 more)

### Community 96 - "fields.js"
Cohesion: 0.21
Nodes (13): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+5 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "storeSource"
Cohesion: 0.08
Nodes (23): HomebrewOwner, Site, Source, sourceFormat, TestHomebrew(), fiveToolsPage(), formatOf(), Report (+15 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Hub"
Cohesion: 0.18
Nodes (6): ClientInfo, goneDevice, TestFiveToolsClassLevelUp(), Hub, TestCloseSessionAndCharacter(), TestDeviceAnnouncements()

### Community 103 - "nullID"
Cohesion: 0.09
Nodes (36): Defense, Condition, CreateConditionParams, TestDefenses(), Defenses, ParseDefenses(), acBonusForm(), effectFix (+28 more)

### Community 105 - "cert.go"
Cohesion: 0.06
Nodes (46): main(), run(), linkTarget, migration, newFixture(), TestOnlyHashIsStored(), TestResolveDeletedCharacter(), TestResolveTouchesStaleSession() (+38 more)

### Community 107 - "ftBook"
Cohesion: 0.19
Nodes (15): ftBook, each(), ftCastingTime(), ftParts(), ftRecords(), TestFiveToolsCopy(), crProficiency(), ftEntries() (+7 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 109 - "BlockEditor"
Cohesion: 0.24
Nodes (4): BlockEditor, json(), links(), sync()

### Community 111 - "export/markdown.go"
Cohesion: 0.35
Nodes (8): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain()

### Community 112 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 116 - "report"
Cohesion: 0.26
Nodes (8): mdCell(), oneLine(), parseConsole(), TestMarkdownCell(), consoleEntry, consoleLog, feedbackForm, report

### Community 117 - ".importParam"
Cohesion: 0.39
Nodes (3): ensureCategory(), csvImport, csvImports

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - ".saveReport"
Cohesion: 0.27
Nodes (6): firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify(), TestFolderName()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Topic"
Cohesion: 0.28
Nodes (5): event, stateHandler, Topic, TopicKind, heldTopics

### Community 127 - "buildReview"
Cohesion: 0.17
Nodes (16): buildReview(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost(), isHTTPLink() (+8 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery() (+16 more)

### Community 137 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 139 - "Book"
Cohesion: 0.12
Nodes (22): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+14 more)

### Community 140 - ".Entry"
Cohesion: 0.08
Nodes (31): MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels() (+23 more)

### Community 143 - "Handlers"
Cohesion: 0.40
Nodes (3): fieldError, pickError, userError

### Community 144 - "CampaignSources"
Cohesion: 0.23
Nodes (13): CampaignOverview(), CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, JoinInfo, agoTime() (+5 more)

### Community 146 - ".q"
Cohesion: 0.06
Nodes (38): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone() (+30 more)

### Community 147 - "importCategories"
Cohesion: 0.22
Nodes (9): categoryKey(), importCategories(), keepValues(), TestImportCategories(), TestKeepValues(), ImportRow, ImportValue, rowFilterCategory() (+1 more)

### Community 148 - "io.Reader"
Cohesion: 0.33
Nodes (4): parseCSV(), TestGuessMapping(), saveSniffed(), hookReader

### Community 157 - "parseMod"
Cohesion: 0.50
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 175 - "planSwitch"
Cohesion: 0.25
Nodes (5): featureSwitch, featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 176 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

## Knowledge Gaps
- **40 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+35 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 454 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **60 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Book`, `perm.go`, `.q`, `nullID`, `viewerCtx`, `Sheet`, `time.Time`, `CompendiumKind`, `testing.T`, `Session`, `FiveToolsJob`, `Client`, `storeSource`, `Hub`, `Pending`, `.importParam`, `.saveReport`?**
  _High betweenness centrality (0.049) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _40 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1333275678168872 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `newTestEnv` to `compendium_templ.go`, `context.Context`, `.sendFeedback`, `SheetView`, `ParseSheet`, `dialog_templ.go`, `perm.go`, `CampaignSources`, `.q`, `App`, `WizardView`, `nullID`, `github.com/a-h/templ.Attributes`, `sourcesEnv`, `IfElse`, `Sheet`, `time.Time`, `CompendiumKind`, `tabs_templ.go`, `tooltip_templ.go`, `github.com/a-h/templ.Component`, `testing.T`, `popover_templ.go`, `Player`, `CombatantView`, `Visibility`, `ParseWiki`, `nullID`, `export/markdown.go`, `deathFix`, `.importParam`?**
  _High betweenness centrality (0.038) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009090705034679356 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `newTestEnv`, `Command`, `Player`, `wikilink.js`?**
  _High betweenness centrality (0.026) - this node is a cross-community bridge._