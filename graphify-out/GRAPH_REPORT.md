# Graph Report - gm-session-manager  (2026-11-24)

## Corpus Check
- 378 files · ~649,931 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4785 nodes · 24649 edges · 165 communities (116 shown, 49 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2473 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `62022856`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Path
- context.Context
- vendorjs/main.go
- createProvider
- DownloadFiveTools
- SheetView
- combobox.js
- condMods
- newTestEnv
- encoding/json.RawMessage
- select.js
- Command
- dialog_templ.go
- redirect.go
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- github.com/a-h/templ.Component
- wikilink.js
- Condition
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- storeSource
- Input
- app.js
- tabbable.js
- useDismiss
- toast_templ.go
- monsterBody
- command.js
- dialog.js
- Import
- ParseWikidot
- tooltip.js
- CompendiumKind
- buildReview
- Sheet
- viewerCtx
- composite.js
- ParseArmorTraining
- CustomMod
- Env
- IfElse
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Weapon
- shadcn-templ Config
- Session
- golangci-lint Config
- popover_templ.go
- Part
- radiogroup.js
- Player
- InputForEntry
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Client
- use_transition_status.js
- playwright
- ftPrerequisite
- Handlers
- checkbox.js
- ParseAidedd
- switch.js
- Item
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- testEnv
- useTypeahead
- ftRec
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- LinkTarget
- mounted
- EffectMods
- nullID
- profFocus
- Store
- deathFix
- ftBook
- ImportView
- BlockEditor
- .LevelUp
- radiogroup_templ.go
- capReader
- .features
- time.Time
- build.sh
- .device
- looseString
- errBodyWriter
- .Router
- export/markdown.go
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
- SearchReadableBlocksRow
- rosterCard
- .ListBlockStepsSince
- .ListItemsByCategory
- .ListSpellsForClass
- .ListSubracesOf
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

## God Nodes (most connected - your core abstractions)
1. `App` - 375 edges
2. `Path()` - 272 edges
3. `newTestEnv()` - 175 edges
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

## Communities (165 total, 49 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "Path"
Cohesion: 0.04
Nodes (108): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListImageCharactersRow (+100 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (118): keepAwake(), bindFiveToolsImages(), ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, CompendiumSource (+110 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.20
Nodes (18): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+10 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "DownloadFiveTools"
Cohesion: 0.08
Nodes (33): ClassFeature, EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), ClassData, parseFiveToolsBook(), parseFiveToolsBookReport() (+25 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (108): Slot, Signed(), Variant, Action(), Card(), Content(), Description(), Header() (+100 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (85): nullID(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestDeleteActiveCombatant(), TestDeleteActiveCombatantLast(), TestDeleteActiveCombatantSkipsOut(), TestDeleteCharacterActiveCombatant() (+77 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (45): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+37 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.08
Nodes (46): Amount, CondOpKind, Effects, userText, parseFeature(), parseSet(), parseSpell(), parseCoin() (+38 more)

### Community 13 - "dialog_templ.go"
Cohesion: 0.13
Nodes (40): ctxKey, ctxState, Props, Button(), Dialog(), Close(), CloseFor(), Content() (+32 more)

### Community 14 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.06
Nodes (45): ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+37 more)

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
Nodes (36): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Campaign, compendiumBack(), compendiumPath(), ImportSummary() (+28 more)

### Community 21 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (117): SwitchFeature, SwitchPlan, Field, Footer(), Checkbox(), Description(), Group(), Title() (+109 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (19): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages (+11 more)

### Community 23 - "Condition"
Cohesion: 0.05
Nodes (50): Defense, Combatant, Condition, CreateConditionParams, ListCombatantsRow, Defenses, ParseDefenses(), EffectEnd (+42 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (179): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+171 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.10
Nodes (19): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+11 more)

### Community 28 - "storeSource"
Cohesion: 0.09
Nodes (23): HomebrewOwner, Site, Source, sourceFormat, TestHomebrew(), fiveToolsPage(), ftsQuery(), HomebrewInput (+15 more)

### Community 29 - "Input"
Cohesion: 0.08
Nodes (51): Coins, Combobox(), Content(), Item(), List(), Field(), Label(), builderFields() (+43 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 34 - "monsterBody"
Cohesion: 0.10
Nodes (35): SearchCompendiumRow, classBadges(), compendiumRow(), copyOrder(), homebrewBadge(), pickerRow(), rowBadges(), SearchComboItems() (+27 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Import"
Cohesion: 0.07
Nodes (44): Import(), Gotchas, internal/web, Notes, TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestMonsterImageLink(), TestPagesRender() (+36 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.15
Nodes (21): block, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster(), sourceLine(), tableText() (+13 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (51): CoreTraits(), dataFor(), DecodeData(), ValidKind(), CountCompendiumRow, UpsertCompendiumEntryParams, Character, Character (+43 more)

### Community 41 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 42 - "Sheet"
Cohesion: 0.03
Nodes (118): LevelUpChoice, internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice() (+110 more)

### Community 43 - "viewerCtx"
Cohesion: 0.23
Nodes (6): notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView, viewerCtx

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseArmorTraining"
Cohesion: 0.13
Nodes (13): armorGrants(), ArmorKindNamed(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack() (+5 more)

### Community 46 - "CustomMod"
Cohesion: 0.10
Nodes (16): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+8 more)

### Community 47 - "Env"
Cohesion: 0.25
Nodes (5): featureSwitch, featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 48 - "IfElse"
Cohesion: 0.04
Nodes (92): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, Habitat, Cell(), Row() (+84 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (162): TestIsLoopback(), TestDecode(), TestCoreTraits(), TestSkillChoice(), RemoveSourceDir(), TestRemoveSourceDir(), TestPageClientRedirect(), TestArmorTrainingGrantsAndRemoval() (+154 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Weapon"
Cohesion: 0.14
Nodes (10): ItemData, TestCustomMods(), Sum(), TestDualWield(), TestUpgradeSheet(), damageLabel(), Weapon, compendiumWeapon() (+2 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Session"
Cohesion: 0.09
Nodes (21): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+13 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "popover_templ.go"
Cohesion: 0.13
Nodes (27): UI conventions, Content(), Description(), Header(), Popover(), state(), Title(), Trigger() (+19 more)

### Community 59 - "Part"
Cohesion: 0.12
Nodes (6): Part, Skill, TestModifier(), skillByKey(), Modifier(), ProficiencyAt()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (36): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views (+28 more)

### Community 62 - "InputForEntry"
Cohesion: 0.22
Nodes (8): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (37): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+29 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Client"
Cohesion: 0.05
Nodes (43): ClientInfo, countingHandler, event, fakeHandler, goneDevice, Handler, message, patchMessage (+35 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Item"
Cohesion: 0.20
Nodes (9): findItem(), parseArmor(), parseItem(), ArmorEffect(), TestGearCommands(), ensureItem(), Item, WornArmor (+1 more)

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

### Community 86 - "testEnv"
Cohesion: 0.13
Nodes (12): assetETags(), immutableAsset(), TestEntityFieldLinks(), testEnv, TestAdminLogin(), TestCampaignJoinLinks(), TestLogout(), classCopy() (+4 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ftRec"
Cohesion: 0.24
Nodes (9): ftCopyRef, ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve() (+1 more)

### Community 96 - "fields.js"
Cohesion: 0.21
Nodes (13): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+5 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "LinkTarget"
Cohesion: 0.17
Nodes (8): LinkTarget, AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink, snapshotBody

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "EffectMods"
Cohesion: 0.25
Nodes (8): findMod(), newMod(), parseMod(), IsModTarget(), EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods()

### Community 103 - "nullID"
Cohesion: 0.20
Nodes (21): TestMonsterAtZeroHPUnaffected(), acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectFields() (+13 more)

### Community 105 - "Store"
Cohesion: 0.06
Nodes (54): migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+46 more)

### Community 106 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 107 - "ftBook"
Cohesion: 0.19
Nodes (15): ftBook, each(), ftCastingTime(), ftParts(), ftRecords(), TestFiveToolsCopy(), crProficiency(), ftEntries() (+7 more)

### Community 108 - "ImportView"
Cohesion: 0.19
Nodes (12): ImageCategory, categoryRow(), columnOverview(), ImportReview(), ImportReviewPage(), importRows(), keepSection(), columnLabel() (+4 more)

### Community 111 - ".LevelUp"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 112 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 116 - ".features"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 117 - "time.Time"
Cohesion: 0.05
Nodes (23): Attachment, browserName(), firstLine(), mdLinkText(), saveScreenshot(), TestUserAgent(), userAgentSummary(), ensureCategory() (+15 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - ".device"
Cohesion: 0.38
Nodes (5): testEnv, TestRequestLog(), TestRequestLogCharacterRemoved(), TestRequestLogReusedID(), TestRequestSent()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - ".Router"
Cohesion: 0.25
Nodes (4): IsLoopback(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 127 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - ".sendFeedback"
Cohesion: 0.24
Nodes (14): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+6 more)

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
Cohesion: 0.13
Nodes (21): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+13 more)

### Community 140 - ".Entry"
Cohesion: 0.10
Nodes (24): MonsterAction, addEntry(), setStat(), MonsterData, TestFiveToolsCoreTraitsShapes(), HomebrewInputFor(), abilityKey(), actionsText() (+16 more)

### Community 144 - "rosterCard"
Cohesion: 0.36
Nodes (9): deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), deviceLabel(), rosterHP(), Device (+1 more)

### Community 157 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

## Knowledge Gaps
- **31 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+26 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 445 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **49 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `Path` to `compendium_templ.go`, `context.Context`, `.sendFeedback`, `SheetView`, `newTestEnv`, `dialog_templ.go`, `Viewer`, `rosterCard`, `App`, `github.com/a-h/templ.Component`, `Condition`, `github.com/a-h/templ.Attributes`, `storeSource`, `Input`, `monsterBody`, `Import`, `CompendiumKind`, `buildReview`, `Sheet`, `IfElse`, `testing.T`, `popover_templ.go`, `Player`, `CombatantView`, `ParseWiki`, `testEnv`, `nullID`, `Store`, `deathFix`, `ImportView`, `time.Time`, `.device`, `export/markdown.go`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _31 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1405702392066237 - nodes in this community are weakly interconnected._
- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Client`, `CompendiumKind`, `Sheet`, `viewerCtx`, `Book`, `Viewer`, `testing.T`, `time.Time`, `testEnv`, `Condition`, `cert.go`, `.Router`, `FiveToolsJob`?**
  _High betweenness centrality (0.056) - this node is a cross-community bridge._
- **Should `Path` be split into smaller, more focused modules?**
  _Cohesion score 0.035240202275600506 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `Import`, `wikilink.js`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._