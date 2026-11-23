# Graph Report - gm-session-manager  (2026-11-24)

## Corpus Check
- 376 files · ~649,037 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4742 nodes · 24565 edges · 144 communities (114 shown, 30 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2463 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `371288fd`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- spellBoxes
- context.Context
- vendorjs/main.go
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- database/sql.NullInt64
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
- IfElse
- wikilink.js
- Queries
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
- github.com/a-h/templ.Component
- command.js
- dialog.js
- Open
- ParseWikidot
- tooltip.js
- CompendiumKind
- NameDesc
- Slugify
- viewerCtx
- composite.js
- Note
- CustomMod
- Env
- github.com/a-h/templ.Component
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- internal/compendium
- shadcn-templ Config
- Session
- golangci-lint Config
- popover_templ.go
- Sheet
- radiogroup.js
- Player
- joinNonEmpty
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
- .renderCompendium
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
- .deleteUnused
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- .editEntity
- mounted
- sourcesEnv
- nullID
- profFocus
- Item
- deathFix
- ftBook
- .source
- BlockEditor
- .LevelUp
- .sourceIssues
- capReader
- .features
- time.Time
- build.sh
- writeBook
- looseString
- errBodyWriter
- .device
- export/markdown.go
- Before/after pairs
- rowFigure
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ASIError
- Book
- .Entry
- sourceUsedError
- rosterCard
- ClassInfo

## God Nodes (most connected - your core abstractions)
1. `App` - 375 edges
2. `Path()` - 272 edges
3. `newTestEnv()` - 175 edges
4. `Sheet` - 133 edges
5. `Button()` - 124 edges
6. `CN()` - 114 edges
7. `SheetView` - 78 edges
8. `IfElse()` - 65 edges
9. `Command` - 60 edges
10. `Queries` - 56 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (144 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "spellBoxes"
Cohesion: 0.67
Nodes (3): TestTaughtPicked(), taughtPicked(), taughtRules()

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (75): keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, CompendiumSource, Encounter, SearchReadableBlocksRow (+67 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.20
Nodes (18): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+10 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.09
Nodes (39): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+31 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (125): CountSharedBySourceRow, Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+117 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (55): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Condition, CreateConditionParams, Handlers, TestMalformedEffectSkipped() (+47 more)

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (129): Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), Gotchas (+121 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (40): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+32 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (54): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+46 more)

### Community 13 - "dialog_templ.go"
Cohesion: 0.12
Nodes (39): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+31 more)

### Community 14 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (58): ListCharacterNamesRow, ListCombatantsRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanChangeSubclass() (+50 more)

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
Nodes (36): IsLoopback(), ClearCookie(), SetCookie(), writeCookie(), FromContext(), RequireAdmin(), Attachment, GetBlockRow (+28 more)

### Community 21 - "IfElse"
Cohesion: 0.07
Nodes (56): Defenses, UI conventions, Checkbox(), Label(), initialValue(), Item(), RadioGroup(), state() (+48 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "Queries"
Cohesion: 0.06
Nodes (41): HomebrewOwner, TestHomebrew(), HomebrewInput, InputForEntry(), SaveHomebrew(), SeedHabitats(), Queries, Combatant (+33 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (174): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Variant (+166 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.12
Nodes (16): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+8 more)

### Community 28 - "storeSource"
Cohesion: 0.09
Nodes (23): Site, Source, TestStoreSourceRollsBackOnAfterSave(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), TestFiveToolsCopyTargets(), TestFiveToolsEdition() (+15 more)

### Community 29 - "Input"
Cohesion: 0.04
Nodes (103): ImageCategory, Combobox(), Content(), Input(), Item(), List(), state(), Description() (+95 more)

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

### Community 34 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (63): CountHabitatsRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, Cell(), itemDetails(), CompendiumPage(), CompendiumResults() (+55 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Open"
Cohesion: 0.06
Nodes (53): sourceFormat, migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+45 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.15
Nodes (20): block, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster(), sourceLine(), tableText() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (51): CoreTraits(), dataFor(), DecodeData(), ValidKind(), ftsQuery(), Search(), CountCompendiumRow, SearchCompendiumRow (+43 more)

### Community 41 - "NameDesc"
Cohesion: 0.15
Nodes (22): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+14 more)

### Community 42 - "Slugify"
Cohesion: 0.07
Nodes (46): Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), AbilityKeys(), FreeCast(), SkillKeys() (+38 more)

### Community 43 - "viewerCtx"
Cohesion: 0.22
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Note"
Cohesion: 0.14
Nodes (8): Note, NoteKind, setAudience(), CreateNoteParams, ListReadableNotesParams, RenameNoteParams, blockAccessForm, noteForm

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (25): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), AddFeat(), SaveFeature() (+17 more)

### Community 47 - "Env"
Cohesion: 0.14
Nodes (12): featureSwitch, ftRaceAbility(), featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch() (+4 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (145): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Field, Look, Badge(), Content(), Description() (+137 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (133): Defense, TestIsLoopback(), TestDecode(), EquipmentOptions(), TestCoreTraits(), TestEquipmentOptions(), readBook(), TestTokensReadAsLabels() (+125 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "internal/compendium"
Cohesion: 0.22
Nodes (7): internal/compendium, SpellData, Spellbook(), TestSpellbook(), CountWord(), Squash(), TestText()

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Session"
Cohesion: 0.07
Nodes (30): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+22 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 59 - "Sheet"
Cohesion: 0.05
Nodes (28): Part, Skill, AddResistance(), TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+20 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (36): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views (+28 more)

### Community 62 - "joinNonEmpty"
Cohesion: 0.12
Nodes (24): habitatsFrom(), ftCell(), ftFeatIncrease(), ftParts(), ftTable(), TestFiveToolsCell(), TestFtFeatIncrease(), TestHabitatsFrom() (+16 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.12
Nodes (33): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList() (+25 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Client"
Cohesion: 0.05
Nodes (44): ClientInfo, countingHandler, event, fakeHandler, goneDevice, Handler, message, patchMessage (+36 more)

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
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - ".renderCompendium"
Cohesion: 0.27
Nodes (4): ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers()

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

### Community 86 - "testEnv"
Cohesion: 0.13
Nodes (13): assetETags(), immutableAsset(), TestEntityFieldLinks(), testEnv, TestAdminLogin(), TestCampaignJoinLinks(), TestLogout(), classCopy() (+5 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".deleteUnused"
Cohesion: 0.29
Nodes (3): ListAttachmentFilesRow, fresh(), inCampaign()

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - ".editEntity"
Cohesion: 0.07
Nodes (18): Entity, LinkTarget, Fields, ParseFields(), entityTopics(), saveFieldLinks(), AddBlockLinkParams, AddEntityFieldLinkParams (+10 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 103 - "nullID"
Cohesion: 0.11
Nodes (37): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), acBonusForm(), effectFix, newEffectFix(), notes() (+29 more)

### Community 105 - "Item"
Cohesion: 0.13
Nodes (12): EquipmentItem, EquipmentOption, findItem(), parseArmor(), ArmorEffect(), TestGearCommands(), ensureItem(), Item (+4 more)

### Community 106 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 107 - "ftBook"
Cohesion: 0.17
Nodes (15): ftBook, ftRec, each(), ftCopyTargets(), ftDecode(), ftRecords(), TestFiveToolsCopy(), ftMod() (+7 more)

### Community 108 - ".source"
Cohesion: 0.27
Nodes (8): testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceMove(), TestSourceOrder()

### Community 111 - ".LevelUp"
Cohesion: 0.07
Nodes (22): LevelUpChoice, armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorGrants() (+14 more)

### Community 112 - ".sourceIssues"
Cohesion: 0.36
Nodes (6): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 116 - ".features"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 117 - "time.Time"
Cohesion: 0.06
Nodes (19): Identity, browserName(), firstLine(), mdLinkText(), TestUserAgent(), userAgentSummary(), deviceKey(), sessionKey() (+11 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "writeBook"
Cohesion: 0.27
Nodes (9): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpBrokenSettings(), TestStartUpRedownloads(), writeBook(), appState (+1 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - ".device"
Cohesion: 0.38
Nodes (5): testEnv, TestRequestLog(), TestRequestLogCharacterRemoved(), TestRequestLogReusedID(), TestRequestSent()

### Community 127 - "export/markdown.go"
Cohesion: 0.05
Nodes (48): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+40 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "rowFigure"
Cohesion: 0.50
Nodes (4): TestRowFigure(), crText(), rowFigure(), textOnly()

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 139 - "Book"
Cohesion: 0.13
Nodes (22): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), CleanDownloads(), downloadBook(), fetchBooks() (+14 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 144 - "rosterCard"
Cohesion: 0.29
Nodes (10): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), deviceLabel(), rosterHP() (+2 more)

### Community 157 - "ClassInfo"
Cohesion: 0.07
Nodes (39): ClassFeature, Decode(), ClassData, dnd5eFeatures(), ListSubclassesOfRow, TestSlotMaxima(), ClassFeature, ClassInfo (+31 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 419 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `newTestEnv` to `compendium_templ.go`, `tabs_templ.go`, `SheetView`, `database/sql.NullInt64`, `dialog_templ.go`, `Viewer`, `rosterCard`, `App`, `IfElse`, `Queries`, `github.com/a-h/templ.Attributes`, `Input`, `github.com/a-h/templ.Component`, `Open`, `CompendiumKind`, `Slugify`, `github.com/a-h/templ.Component`, `testing.T`, `Session`, `Player`, `CombatantView`, `ParseWiki`, `testEnv`, `.deleteUnused`, `sourcesEnv`, `nullID`, `deathFix`, `.source`, `time.Time`, `.device`, `export/markdown.go`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14049371757395815 - nodes in this community are weakly interconnected._
- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `database/sql.NullInt64`, `Command`, `Viewer`, `Queries`, `cert.go`, `ClassInfo`, `CompendiumKind`, `Slugify`, `viewerCtx`, `Note`, `internal/compendium`, `Sheet`, `FiveToolsJob`, `Client`, `.renderCompendium`, `testEnv`, `.deleteUnused`, `.editEntity`, `.sourceIssues`, `time.Time`, `writeBook`?**
  _High betweenness centrality (0.042) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01004566210045662 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `newTestEnv`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._