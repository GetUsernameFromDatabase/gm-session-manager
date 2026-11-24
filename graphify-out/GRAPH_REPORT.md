# Graph Report - gm-session-manager  (2026-11-24)

## Corpus Check
- 380 files · ~652,817 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4811 nodes · 24761 edges · 178 communities (117 shown, 61 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2487 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `304f7997`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- github.com/a-h/templ.Component
- context.Context
- vendorjs/main.go
- createProvider
- Campaign
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
- Queries
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- sourcesEnv
- select_templ.go
- app.js
- tabbable.js
- useDismiss
- Sheet
- Sum
- command.js
- dialog.js
- .editEntity
- ParseWikidot
- tooltip.js
- CompendiumKind
- tabs_templ.go
- Input
- database/sql.NullInt64
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
- empty_templ.go
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
- Values Go and JS share
- checkbox.js
- ParseAidedd
- switch.js
- parseHP
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
- .renderCompendium
- mounted
- Hub
- nullID
- profFocus
- Store
- ftBook
- Pending
- BlockEditor
- TestSourceSwitchKeepsEdits
- deathFix
- capReader
- .features
- time.Time
- build.sh
- condMods
- looseString
- errBodyWriter
- SearchReadableBlocksRow
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
- sourceRow
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
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
- ListSharedWithRow
- .ListSubracesOf
- .ListBlockStepsSince
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
10. `Queries` - 57 edges

## Surprising Connections (you probably didn't know these)
- `Notes` --references--> `linkClick()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/assets/js/wikilink.js
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (178 total, 61 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (55): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+47 more)

### Community 1 - "github.com/a-h/templ.Component"
Cohesion: 0.07
Nodes (79): Input(), Addon(), Button(), Input(), InputGroup(), Kbd(), Scripts(), Separator() (+71 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (64): keepAwake(), MoveSource(), SetSourceOrder(), CompendiumSource, CreateCombatantParams, EntityKind, heldCollected(), notFoundIsNil() (+56 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Campaign"
Cohesion: 0.32
Nodes (5): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (130): Coins, Slot, Signed(), Button(), Variant, Card(), Content(), Description() (+122 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.08
Nodes (50): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestMonsterImageLink(), TestBuilderCompendium() (+42 more)

### Community 9 - "newTestEnv"
Cohesion: 0.03
Nodes (105): ParsePageURL(), TestParsePageURL(), TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight() (+97 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (46): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+38 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.09
Nodes (43): CondOpKind, Effects, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn(), parseInspiration() (+35 more)

### Community 13 - "dialog_templ.go"
Cohesion: 0.11
Nodes (44): ctxKey, ctxState, Props, Dialog(), Close(), Content(), Description(), Dialog() (+36 more)

### Community 14 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (63): ListCharacterNamesRow, GetBlockRow, ListCombatantsRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect() (+55 more)

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
Nodes (43): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Attachment, Campaign, Character, compendiumBack() (+35 more)

### Community 21 - "WizardView"
Cohesion: 0.07
Nodes (56): Props, SpellColumn(), Footer(), Checkbox(), initialChecked(), Label(), initialValue(), Item() (+48 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (19): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages (+11 more)

### Community 23 - "Queries"
Cohesion: 0.06
Nodes (58): Decode(), FiveToolsImages(), bindFiveToolsImages(), Queries, Combatant, CompendiumEntry, Env, ArmorEffect() (+50 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (182): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+174 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.15
Nodes (10): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+2 more)

### Community 28 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 29 - "select_templ.go"
Cohesion: 0.06
Nodes (60): ImageCategory, Defenses, Description(), Field(), Group(), Label(), Input(), damageChoice() (+52 more)

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
Cohesion: 0.03
Nodes (103): EquipmentItem, LevelUpChoice, internal/compendium, Cantrips(), firstTable(), EquipmentOption, LevelSpells(), OptionPrerequisite() (+95 more)

### Community 34 - "Sum"
Cohesion: 0.19
Nodes (4): TestCustomMods(), TestExplainSums(), Sum(), TestWithEffects()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - ".editEntity"
Cohesion: 0.07
Nodes (17): Entity, LinkTarget, Fields, ParseFields(), entityTopics(), saveFieldLinks(), AddBlockLinkParams, AddEntityFieldLinkParams (+9 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.10
Nodes (30): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, dnd5eFeatures(), TestFiveToolsCoreTraitsShapes(), withCoreTraits() (+22 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (47): HomebrewOwner, Site, Source, sourceFormat, TestStoreSourceRollsBackOnAfterSave(), dataFor(), DecodeData(), ValidKind() (+39 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.10
Nodes (35): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+27 more)

### Community 42 - "Input"
Cohesion: 0.15
Nodes (24): ctxState, Chip(), Chips(), ChipsInput(), Combobox(), Content(), Input(), Item() (+16 more)

### Community 43 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (53): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Condition, Encounter, CreateConditionParams, EncounterStatus, TestMalformedEffectSkipped() (+45 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 46 - "CustomMod"
Cohesion: 0.07
Nodes (23): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), UsesLabel(), CustomMod (+15 more)

### Community 47 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 48 - "IfElse"
Cohesion: 0.04
Nodes (96): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, SearchCompendiumRow, Habitat, UI conventions (+88 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (140): Defense, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestIsLoopback(), TestDecode(), EquipmentOptions(), TestEquipmentOptions() (+132 more)

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
Cohesion: 0.11
Nodes (10): IsLoopback(), Identity, RequireAdmin(), Session, Role, deviceKey(), sessionKey(), sameOriginOnly() (+2 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "empty_templ.go"
Cohesion: 0.08
Nodes (52): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListImageCharactersRow (+44 more)

### Community 59 - "Part"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (37): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+29 more)

### Community 62 - "ParseEffects"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.09
Nodes (40): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+32 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Client"
Cohesion: 0.09
Nodes (15): countingHandler, event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch (+7 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - "Values Go and JS share"
Cohesion: 0.14
Nodes (15): Gotchas, internal/web, Notes, Values Go and JS share, TestEffectNoteMaxMatchesTag(), TestPagesRender(), TestNoRouteLiterals(), parseSharedJS() (+7 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.13
Nodes (27): MonsterAction, abilityMod(), addEntry(), ParseAidedd(), setStat(), statLabel(), textLines(), MonsterData (+19 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "parseHP"
Cohesion: 0.19
Nodes (11): Amount, userText, parseCoin(), CritDamage(), Damage(), parseDamage(), parseHP(), parseRevive() (+3 more)

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
Cohesion: 0.08
Nodes (11): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), randomID(), sync() (+3 more)

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
Nodes (20): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+12 more)

### Community 96 - "fields.js"
Cohesion: 0.21
Nodes (13): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+5 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - ".renderCompendium"
Cohesion: 0.12
Nodes (7): TestHomebrew(), ftsQuery(), Search(), CountCompendiumRow, ImportSummary(), CompendiumURL(), CopyKey()

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Hub"
Cohesion: 0.19
Nodes (4): ClientInfo, goneDevice, Hub, TestCloseSessionAndCharacter()

### Community 103 - "nullID"
Cohesion: 0.23
Nodes (19): acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectModRows(), TestEffectModsReadOnly() (+11 more)

### Community 105 - "Store"
Cohesion: 0.05
Nodes (72): fixture, tokens, main(), run(), linkTarget, migration, isCleared(), newFixture() (+64 more)

### Community 107 - "ftBook"
Cohesion: 0.19
Nodes (15): ftBook, each(), ftCastingTime(), ftParts(), ftRecords(), TestFiveToolsCopy(), crProficiency(), ftEntries() (+7 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 111 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 112 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 116 - ".features"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 117 - "time.Time"
Cohesion: 0.05
Nodes (22): ListAttachmentFilesRow, ListImagesRow, Note, NoteKind, fresh(), inCampaign(), FiveToolsView, NewFiveToolsView() (+14 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 127 - "buildReview"
Cohesion: 0.05
Nodes (49): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+41 more)

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
Cohesion: 0.12
Nodes (17): habitatsFrom(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), HomebrewInputFor(), InputForEntry(), actionsText(), bonusText() (+9 more)

### Community 143 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 144 - "CampaignSources"
Cohesion: 0.15
Nodes (19): campaignSourceRow(), CampaignSources(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), CampaignOverview(), CampaignRoster(), deviceRow() (+11 more)

### Community 146 - ".q"
Cohesion: 0.08
Nodes (30): serve(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+22 more)

### Community 148 - "sourceRow"
Cohesion: 0.47
Nodes (5): sourceIssue(), sourceRow(), SourceCopy, SourceIssue, sheetRef

### Community 175 - "planSwitch"
Cohesion: 0.23
Nodes (6): featureSwitch, parseFeature(), featureNames(), SubclassRef, planSwitch(), sameSlug()

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 456 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **61 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Campaign`, `Book`, `perm.go`, `.q`, `Queries`, `Sheet`, `.editEntity`, `CompendiumKind`, `database/sql.NullInt64`, `testing.T`, `Session`, `FiveToolsJob`, `Client`, `.renderCompendium`, `Hub`, `Store`, `Pending`, `time.Time`?**
  _High betweenness centrality (0.049) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13997838759265177 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `newTestEnv` to `compendium_templ.go`, `github.com/a-h/templ.Component`, `.sendFeedback`, `Campaign`, `SheetView`, `ParseSheet`, `dialog_templ.go`, `perm.go`, `CampaignSources`, `.q`, `App`, `WizardView`, `sourceRow`, `github.com/a-h/templ.Attributes`, `sourcesEnv`, `select_templ.go`, `Sheet`, `tabs_templ.go`, `Input`, `database/sql.NullInt64`, `IfElse`, `testing.T`, `empty_templ.go`, `Player`, `CombatantView`, `Values Go and JS share`, `ParseWiki`, `nullID`, `TestSourceSwitchKeepsEdits`, `deathFix`, `time.Time`, `buildReview`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._
- **Should `github.com/a-h/templ.Component` be split into smaller, more focused modules?**
  _Cohesion score 0.06594399277326106 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `Values Go and JS share` to `newTestEnv`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._