# Graph Report - gm-session-manager  (2026-11-24)

## Corpus Check
- 379 files · ~651,481 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4805 nodes · 24732 edges · 184 communities (124 shown, 60 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2479 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `71f23648`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- empty_templ.go
- context.Context
- vendorjs/main.go
- createProvider
- Parse
- SheetView
- combobox.js
- database/sql.NullInt64
- newTestEnv
- encoding/json.RawMessage
- select.js
- Command
- github.com/a-h/templ.Component
- redirect.go
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- WizardView
- wikilink.js
- Queries
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- Search
- select_templ.go
- app.js
- tabbable.js
- useDismiss
- Sheet
- monsterBody
- command.js
- dialog.js
- TestPagesRender
- ParseWikidot
- tooltip.js
- CompendiumKind
- tabs_templ.go
- DownloadFiveTools
- export/markdown.go
- composite.js
- ParseArmorTraining
- CustomMod
- names
- IfElse
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
- parseHP
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- State
- use_transition_status.js
- playwright
- ftPrerequisite
- report
- checkbox.js
- ParseAidedd
- switch.js
- Client
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- TestLevelUpPushesTracker
- useTypeahead
- NewClient
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- internal/compendium
- mounted
- Hub
- nullID
- profFocus
- Store
- Campaign
- ftBook
- Pending
- BlockEditor
- .LevelUp
- importCategories
- capReader
- .features
- time.Time
- build.sh
- io.Reader
- looseString
- errBodyWriter
- fakeHandler
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
- ListSharedWithRow
- sourceUsedError
- Handlers
- rosterCard
- .coreTraits
- TestSourceSwitchKeepsEdits
- Props
- ParseSheet
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- spellCounts
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
- parseFeature
- parseMod
- parseCommand
- SearchReadableBlocksRow
- consoleEntry
- .ListBlockStepsSince
- .ListItemsByCategory
- .ListSpellsForClass
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 375 edges
2. `Path()` - 270 edges
3. `newTestEnv()` - 176 edges
4. `Sheet` - 134 edges
5. `Button()` - 124 edges
6. `CN()` - 114 edges
7. `SheetView` - 79 edges
8. `IfElse()` - 65 edges
9. `Command` - 60 edges
10. `Queries` - 56 edges

## Surprising Connections (you probably didn't know these)
- `Notes` --references--> `linkClick()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/assets/js/wikilink.js
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
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

## Communities (184 total, 60 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (55): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+47 more)

### Community 1 - "empty_templ.go"
Cohesion: 0.05
Nodes (73): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListImageCharactersRow (+65 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (85): keepAwake(), CompendiumSource, Encounter, Note, CreateCombatantParams, EncounterStatus, EntityKind, LinkTarget (+77 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Parse"
Cohesion: 0.07
Nodes (36): Defense, feature(), TestFiveToolsClasses(), TestArmorTrainingGrantsAndRemoval(), Heal(), Parse(), tokenize(), TestApplyHPMonsterUnaffected() (+28 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (116): CountSharedBySourceRow, UsesLabel(), Slot, Signed(), Variant, Action(), Card(), Content() (+108 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (46): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Attachment, GetBlockRow, CreateConditionParams, EffectEnd, firstLine() (+38 more)

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (84): isCleared(), PageSourceKey(), ParsePageURL(), TestParsePageURL(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestEncounterTransitions(), TestMonsterPortrait() (+76 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (54): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+46 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.10
Nodes (38): CondOpKind, Effects, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn(), parseCoin() (+30 more)

### Community 13 - "github.com/a-h/templ.Component"
Cohesion: 0.07
Nodes (69): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+61 more)

### Community 14 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (56): ListCharacterNamesRow, Character, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanChangeSubclass() (+48 more)

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
Cohesion: 0.03
Nodes (42): IsLoopback(), ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), Campaign (+34 more)

### Community 21 - "WizardView"
Cohesion: 0.09
Nodes (47): Label(), initialValue(), Item(), RadioGroup(), state(), IfElse(), abilitySelect(), levelUpForm() (+39 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (19): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages (+11 more)

### Community 23 - "Queries"
Cohesion: 0.04
Nodes (61): FiveToolsImages(), bindFiveToolsImages(), Queries, Combatant, Condition, Entity, ListCombatantsRow, Topic (+53 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (169): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+161 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.15
Nodes (10): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+2 more)

### Community 28 - "Search"
Cohesion: 0.22
Nodes (5): TestHomebrew(), ftsQuery(), Search(), Completions(), TestCampaignSourceVisibility()

### Community 29 - "select_templ.go"
Cohesion: 0.04
Nodes (101): ImageCategory, Coins, Uses, Combobox(), Content(), Item(), List(), Description() (+93 more)

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
Nodes (106): featureSwitch, LevelUpChoice, Cantrips(), CoreTraits(), firstTable(), LevelSpells(), OptionPrerequisite(), Options() (+98 more)

### Community 34 - "monsterBody"
Cohesion: 0.14
Nodes (25): TestRowFigure(), commaPrefixed(), compendiumFindURL(), crText(), levelsText(), parenthesized(), rowFigure(), signedMap() (+17 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "TestPagesRender"
Cohesion: 0.18
Nodes (9): TestUpgradeSheet(), EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Gotchas, internal/web, Notes (+1 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.14
Nodes (23): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (56): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), dataFor(), Decode(), DecodeData(), SpellData (+48 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.09
Nodes (38): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+30 more)

### Community 42 - "DownloadFiveTools"
Cohesion: 0.15
Nodes (21): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook(), DownloadFiveTools() (+13 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.31
Nodes (9): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+1 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 46 - "CustomMod"
Cohesion: 0.09
Nodes (13): FeatureEffects, SaveFeature(), CustomMod, indexAbilityName(), SpellMods(), checkMods(), modPairs(), TestSpellMods() (+5 more)

### Community 47 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 48 - "IfElse"
Cohesion: 0.04
Nodes (110): CountHabitatsRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, SearchCompendiumRow, Habitat, Field, UI conventions (+102 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (120): TestIsLoopback(), TestDecode(), TestCoreTraits(), TestSkillChoice(), RemoveSourceDir(), TestRemoveSourceDir(), TestPageClientRedirect(), TestArmorUntrained() (+112 more)

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
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Session"
Cohesion: 0.08
Nodes (18): tokens, constantTimeEqual(), credentialsMatch(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity(), TestRequireAdmin(), TestTokens() (+10 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "popover_templ.go"
Cohesion: 0.20
Nodes (13): Description(), state(), Trigger(), triggerID(), Align, ContentProps, ctxKey, ctxState (+5 more)

### Community 59 - "Part"
Cohesion: 0.08
Nodes (15): Part, Skill, ItemData, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+7 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.12
Nodes (25): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+17 more)

### Community 62 - "parseHP"
Cohesion: 0.21
Nodes (9): Amount, userText, CritDamage(), Damage(), parseDamage(), parseHP(), parseRevive(), Set() (+1 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.06
Nodes (60): Defenses, Path(), Header(), Title(), TestCharacterLooks(), damageChoice(), BuilderPickItems(), BuilderPicks() (+52 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "State"
Cohesion: 0.13
Nodes (7): Handler, patchMessage, stateHandler, viewPatch, Diff(), State, TestDiff()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - "report"
Cohesion: 0.24
Nodes (8): folderName(), mdCell(), oneLine(), slugify(), TestFolderName(), TestMarkdownCell(), feedbackForm, report

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Client"
Cohesion: 0.20
Nodes (3): message, Client, isDevice()

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
Cohesion: 0.33
Nodes (8): CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks(), WikiLinks(), WikiLink, WikiPart

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "notes.js"
Cohesion: 0.08
Nodes (11): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), randomID(), sync() (+3 more)

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - "TestLevelUpPushesTracker"
Cohesion: 0.31
Nodes (8): TestLevelUpPushesTracker(), dial(), patchOf(), readUntil(), rosterRefetch(), TestDeviceSeenOnLeave(), TestLiveUpdatesAndKick(), wsMessage

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "NewClient"
Cohesion: 0.20
Nodes (18): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+10 more)

### Community 96 - "fields.js"
Cohesion: 0.21
Nodes (13): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+5 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "internal/compendium"
Cohesion: 0.07
Nodes (23): HomebrewOwner, Site, Source, sourceFormat, internal/compendium, TestStoreSourceRollsBackOnAfterSave(), FiveToolsBooks(), fiveToolsPage() (+15 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Hub"
Cohesion: 0.23
Nodes (6): ClientInfo, event, goneDevice, Hub, TestCloseSessionAndCharacter(), TestDeviceAnnouncements()

### Community 103 - "nullID"
Cohesion: 0.10
Nodes (39): DeathSaves, nullID(), deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP() (+31 more)

### Community 105 - "Store"
Cohesion: 0.05
Nodes (63): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), TestResolveDeletedCharacter(), TestResolveTouchesStaleSession() (+55 more)

### Community 106 - "Campaign"
Cohesion: 0.28
Nodes (6): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks()

### Community 107 - "ftBook"
Cohesion: 0.13
Nodes (21): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftCopyTargets(), ftDecode() (+13 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 111 - ".LevelUp"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 112 - "importCategories"
Cohesion: 0.22
Nodes (9): categoryKey(), importCategories(), keepValues(), TestImportCategories(), TestKeepValues(), ImportRow, ImportValue, rowFilterCategory() (+1 more)

### Community 116 - ".features"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 117 - "time.Time"
Cohesion: 0.08
Nodes (14): ListAttachmentFilesRow, ListImagesRow, browserName(), fresh(), TestUserAgent(), userAgentSummary(), BlockStep, LevelUp (+6 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "io.Reader"
Cohesion: 0.29
Nodes (5): saveScreenshot(), parseCSV(), TestGuessMapping(), saveSniffed(), hookReader

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

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
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 137 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 139 - "Book"
Cohesion: 0.13
Nodes (21): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+13 more)

### Community 140 - ".Entry"
Cohesion: 0.10
Nodes (26): MonsterAction, addEntry(), setStat(), MonsterData, TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor(), InputForEntry() (+18 more)

### Community 143 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 144 - "rosterCard"
Cohesion: 0.33
Nodes (10): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), deviceLabel(), rosterHP() (+2 more)

### Community 145 - ".coreTraits"
Cohesion: 0.33
Nodes (5): ftProfTexts(), ftSkillNames(), TestFiveToolsCoreTraitsShapes(), withCoreTraits(), ArmorKindNamed()

### Community 146 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 147 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 148 - "ParseSheet"
Cohesion: 0.07
Nodes (53): TestRealSource(), Import(), ParseSheet(), joinHabitats(), TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestHabitatCopies(), TestHabitatsMatchSchema() (+45 more)

### Community 157 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

### Community 175 - "parseFeature"
Cohesion: 0.33
Nodes (4): parseFeature(), parseFeatureSet(), parseUses(), sameSlug()

### Community 176 - "parseMod"
Cohesion: 0.50
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 177 - "parseCommand"
Cohesion: 0.40
Nodes (5): coinCommand(), commandText(), hpCommand(), parseCommand(), damageArgs()

### Community 179 - "consoleEntry"
Cohesion: 0.67
Nodes (3): parseConsole(), consoleEntry, consoleLog

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 456 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **60 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `CombatantView` to `compendium_templ.go`, `empty_templ.go`, `.sendFeedback`, `SheetView`, `database/sql.NullInt64`, `newTestEnv`, `github.com/a-h/templ.Component`, `Viewer`, `rosterCard`, `TestSourceSwitchKeepsEdits`, `ParseSheet`, `App`, `WizardView`, `Queries`, `github.com/a-h/templ.Attributes`, `Search`, `select_templ.go`, `Sheet`, `monsterBody`, `TestPagesRender`, `tabs_templ.go`, `export/markdown.go`, `IfElse`, `testing.T`, `Session`, `Player`, `ParseWiki`, `TestLevelUpPushesTracker`, `nullID`, `Campaign`, `time.Time`?**
  _High betweenness centrality (0.045) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13993272757712738 - nodes in this community are weakly interconnected._
- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `database/sql.NullInt64`, `newTestEnv`, `Book`, `Viewer`, `Queries`, `Search`, `Sheet`, `CompendiumKind`, `testing.T`, `FiveToolsJob`, `State`, `internal/compendium`, `Hub`, `Store`, `Campaign`, `Pending`, `time.Time`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Should `empty_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.04746835443037975 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `CombatantView`, `Command`, `TestPagesRender`, `wikilink.js`?**
  _High betweenness centrality (0.034) - this node is a cross-community bridge._