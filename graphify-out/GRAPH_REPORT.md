# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 372 files · ~643,390 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4690 nodes · 24215 edges · 157 communities (124 shown, 33 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2425 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `25aa24c6`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Slugify
- context.Context
- database/sql.NullInt64
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- Queries
- Path
- encoding/json.RawMessage
- select.js
- scoped
- formSelect
- rosterCard
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- IfElse
- wikilink.js
- github.com/gin-gonic/gin.Context
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- storeSource
- ParseSheet
- app.js
- tabbable.js
- useDismiss
- Client
- popover_templ.go
- command.js
- dialog.js
- Import
- ParseWikidot
- tooltip.js
- CompendiumKind
- .levelUp
- viewerCtx
- Session
- composite.js
- buildReview
- CustomMod
- parseFeature
- nullID
- .fail
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- ftItemData
- shadcn-templ Config
- Item
- golangci-lint Config
- Open
- Weapon
- radiogroup.js
- vendorjs/main.go
- IfElse
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Queries
- use_transition_status.js
- playwright
- testEnv
- blockMenu
- checkbox.js
- ParseAidedd
- switch.js
- deathFix
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- .q
- useTypeahead
- Session
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- parseHP
- mounted
- .sendFeedback
- Player
- profFocus
- CompendiumPath
- redirect.go
- ftBook
- NewClient
- BlockEditor
- .LevelUp
- Input
- capReader
- .features
- time.Time
- build.sh
- config.go
- looseString
- errBodyWriter
- Client
- export/markdown.go
- Before/after pairs
- Hub
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- Command
- ASIError
- Book
- .Entry
- WeaponKind
- sourceUsedError
- Pending
- .deleteUnused
- wizardChoice
- github.com/a-h/templ.Component
- LinkTarget
- InputForEntry
- DBTX
- Note
- sourcesEnv
- fakeHandler
- .sourceIssues
- ftDuration
- Book
- parseRevive

## God Nodes (most connected - your core abstractions)
1. `App` - 374 edges
2. `Path()` - 268 edges
3. `newTestEnv()` - 170 edges
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
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (157 total, 33 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "Slugify"
Cohesion: 0.10
Nodes (34): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, Slugify() (+26 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (84): keepAwake(), keepAwake(), bindFiveToolsImages(), ListImageCharactersRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, SearchReadableBlocksRow (+76 more)

### Community 3 - "database/sql.NullInt64"
Cohesion: 0.04
Nodes (37): GetImageRow, ListImagesRow, Attachment, ensureCategory(), imageItem(), imageVersion(), nullString(), parseCSV() (+29 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.09
Nodes (44): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+36 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (122): UsesLabel(), ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card() (+114 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Queries"
Cohesion: 0.08
Nodes (38): featureSwitch, Decode(), Queries, CompendiumEntry, Env, ClassInfo, ClassLevel, WornArmor (+30 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (86): TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+78 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (46): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+38 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "scoped"
Cohesion: 0.17
Nodes (17): parseProf(), parseProfSet(), parseSpell(), setIn(), parseInspiration(), parseShield(), parseDeathSave(), init() (+9 more)

### Community 13 - "formSelect"
Cohesion: 0.05
Nodes (52): ImageCategory, Defenses, firstValue(), modTargetLabel(), weaponAbilities(), EffectFieldsForm(), effectModFields(), EffectModRow() (+44 more)

### Community 14 - "rosterCard"
Cohesion: 0.29
Nodes (10): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), deviceLabel(), rosterHP() (+2 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.06
Nodes (50): ListCharacterNamesRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanChangeSubclass(), CanCharacter() (+42 more)

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
Nodes (22): compendiumBack(), compendiumPath(), ImportSummary(), viewerCampaign(), fiveToolsKey(), TestImportSummary(), formIntQuery(), isHTMX() (+14 more)

### Community 21 - "IfElse"
Cohesion: 0.11
Nodes (36): Props, Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup(), state() (+28 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.05
Nodes (31): ClearCookie(), SetCookie(), writeCookie(), Entity, CharacterFile, NewCharacterFile(), Fields, ParseFields() (+23 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (142): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+134 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.10
Nodes (19): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+11 more)

### Community 28 - "storeSource"
Cohesion: 0.11
Nodes (20): Site, Source, sourceFormat, TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), formatOf(), Report, ImportProgress() (+12 more)

### Community 29 - "ParseSheet"
Cohesion: 0.09
Nodes (42): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+34 more)

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
Nodes (14): event, Handler, message, patchMessage, stateHandler, viewPatch, Topic, TopicKind (+6 more)

### Community 34 - "popover_templ.go"
Cohesion: 0.14
Nodes (22): Field, Header(), state(), Title(), Trigger(), triggerID(), effectBits(), combatCondition() (+14 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Import"
Cohesion: 0.07
Nodes (42): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+34 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.11
Nodes (27): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+19 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.04
Nodes (29): dataFor(), DecodeData(), ListHomebrewByOwnerRow, ListHomebrewRow, UpsertCompendiumEntryParams, CompendiumKind, choiceName(), WizardChoice (+21 more)

### Community 41 - ".levelUp"
Cohesion: 0.06
Nodes (40): TextData, ListSubclassesOfRow, FreeCast(), SwapFeatSpell(), LevelUpPlan, Spell, TestDurationRounds(), textMarkdown() (+32 more)

### Community 42 - "viewerCtx"
Cohesion: 0.23
Nodes (6): notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView, viewerCtx

### Community 43 - "Session"
Cohesion: 0.14
Nodes (8): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), TestHandlerErrors(), sameOriginOnly(), useValidateTag()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.12
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

### Community 46 - "CustomMod"
Cohesion: 0.11
Nodes (18): FeatureEffects, findMod(), newMod(), parseMod(), SaveFeature(), CustomMod, IsModTarget(), indexAbilityName() (+10 more)

### Community 47 - "parseFeature"
Cohesion: 0.67
Nodes (3): parseFeature(), parseFeatureSet(), parseUses()

### Community 48 - "nullID"
Cohesion: 0.23
Nodes (19): acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectModRows(), TestEffectModsReadOnly() (+11 more)

### Community 49 - ".fail"
Cohesion: 0.07
Nodes (11): Campaign, Character, GetBlockRow, pathCampaign(), linkCharacterLook(), CompendiumRest(), HomebrewView, CampaignSwitch (+3 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (119): Defense, TestDecode(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants() (+111 more)

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
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Open"
Cohesion: 0.09
Nodes (29): migration, newFixture(), TestOnlyHashIsStored(), TestResolveDeletedCharacter(), TestResolveTouchesStaleSession(), apply(), loadMigrations(), Migrate() (+21 more)

### Community 59 - "Weapon"
Cohesion: 0.06
Nodes (21): Part, Skill, AddFeat(), TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+13 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "IfElse"
Cohesion: 0.10
Nodes (32): SpellData, classBadges(), compendiumRow(), copyOrder(), rowBadges(), TestRowFigure(), alignmentColor(), commaPrefixed() (+24 more)

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

### Community 68 - "Queries"
Cohesion: 0.04
Nodes (54): Combatant, Condition, Encounter, CreateConditionParams, ListCombatantsRow, EffectEnd, EncounterStatus, combatantRow() (+46 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "testEnv"
Cohesion: 0.18
Nodes (10): TestEntityFieldLinks(), testEnv, TestAdminLogin(), TestCampaignJoinLinks(), TestLogout(), classCopy(), testEnv, TestRunCommandUsesCurrentSheet() (+2 more)

### Community 72 - "blockMenu"
Cohesion: 0.08
Nodes (42): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent(), Props (+34 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

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

### Community 86 - ".q"
Cohesion: 0.10
Nodes (26): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone() (+18 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Session"
Cohesion: 0.15
Nodes (12): fixture, tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken() (+4 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "parseHP"
Cohesion: 0.22
Nodes (11): Amount, parseCoin(), CritDamage(), Damage(), parseDamage(), parseHP(), parseMaxHP(), parseTemp() (+3 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.24
Nodes (14): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+6 more)

### Community 103 - "Player"
Cohesion: 0.09
Nodes (35): Player(), TestCanRemoveCondition(), BenchmarkSheetState(), testEnv, Websocket and live views, watchSheets(), TestNoteViewAudiences(), isRandomID() (+27 more)

### Community 105 - "CompendiumPath"
Cohesion: 0.05
Nodes (86): CountCompendiumRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListMonstersByNameRow, SearchCompendiumRow, CompendiumSource, Habitat (+78 more)

### Community 106 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (16): ftBook, ftCopyRef, ftRec, each(), ftCopyTargets(), ftDecode(), ftRecords(), TestFiveToolsCopy() (+8 more)

### Community 108 - "NewClient"
Cohesion: 0.22
Nodes (19): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+11 more)

### Community 111 - ".LevelUp"
Cohesion: 0.09
Nodes (15): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 112 - "Input"
Cohesion: 0.07
Nodes (78): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+70 more)

### Community 116 - ".features"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 117 - "time.Time"
Cohesion: 0.06
Nodes (29): Identity, FromContext(), Handlers, browserName(), firstLine(), folderName(), mdCell(), mdLinkText() (+21 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 127 - "export/markdown.go"
Cohesion: 0.35
Nodes (8): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain()

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "Hub"
Cohesion: 0.24
Nodes (4): ClientInfo, goneDevice, TestFiveToolsClassLevelUp(), Hub

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "Command"
Cohesion: 0.14
Nodes (20): CondOpKind, Effects, parseSet(), parseWeapon(), SaveWeapon(), Command, CondOp, condCommand() (+12 more)

### Community 139 - "Book"
Cohesion: 0.27
Nodes (7): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), TestMirrorGetMissingFile(), fetch()

### Community 140 - ".Entry"
Cohesion: 0.09
Nodes (27): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInput, HomebrewInputFor() (+19 more)

### Community 141 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 143 - "Pending"
Cohesion: 0.27
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 144 - ".deleteUnused"
Cohesion: 0.23
Nodes (5): ListAttachmentFilesRow, attachmentRefs(), fresh(), inCampaign(), FileUse

### Community 145 - "wizardChoice"
Cohesion: 0.06
Nodes (35): EquipmentItem, internal/compendium, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), TestCoreTraits(), TestEquipmentOptions() (+27 more)

### Community 146 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (112): Action(), AspectRatio(), Footer(), ChipsInput(), Empty(), Group(), Label(), Separator() (+104 more)

### Community 147 - "LinkTarget"
Cohesion: 0.17
Nodes (8): LinkTarget, AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink, snapshotBody

### Community 148 - "InputForEntry"
Cohesion: 0.22
Nodes (8): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 150 - "Note"
Cohesion: 0.25
Nodes (5): Note, NoteKind, CreateNoteParams, RenameNoteParams, noteForm

### Community 151 - "sourcesEnv"
Cohesion: 0.28
Nodes (7): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 153 - ".sourceIssues"
Cohesion: 0.53
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 154 - "ftDuration"
Cohesion: 0.50
Nodes (5): ftDistance, ftAmount(), ftCastingTime(), ftDuration(), ftRange()

### Community 155 - "Book"
Cohesion: 0.17
Nodes (17): rawBook, bookGroup(), fetchBooks(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks(), sortBooks() (+9 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 414 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **33 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Slugify`, `context.Context`, `database/sql.NullInt64`, `Hub`, `Queries`, `Command`, `formSelect`, `Viewer`, `.deleteUnused`, `wizardChoice`, `github.com/gin-gonic/gin.Context`, `.sourceIssues`, `cert.go`, `Book`, `Client`, `CompendiumKind`, `.levelUp`, `viewerCtx`, `Session`, `.fail`, `testing.T`, `FiveToolsJob`, `Queries`, `testEnv`, `CompendiumPath`, `time.Time`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14492683242683244 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `database/sql.NullInt64`, `tabs_templ.go`, `SheetView`, `formSelect`, `rosterCard`, `Viewer`, `.deleteUnused`, `github.com/a-h/templ.Component`, `App`, `IfElse`, `github.com/gin-gonic/gin.Context`, `sourcesEnv`, `ParseSheet`, `popover_templ.go`, `Import`, `.levelUp`, `nullID`, `testing.T`, `Open`, `IfElse`, `CombatantView`, `Queries`, `testEnv`, `blockMenu`, `deathFix`, `ParseWiki`, `.q`, `.sendFeedback`, `Player`, `CompendiumPath`, `Input`, `export/markdown.go`?**
  _High betweenness centrality (0.047) - this node is a cross-community bridge._
- **Should `Slugify` be split into smaller, more focused modules?**
  _Cohesion score 0.10241545893719807 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `scoped`, `Import`, `wikilink.js`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._