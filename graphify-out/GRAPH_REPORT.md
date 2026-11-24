# Graph Report - gm-session-manager  (2026-11-24)

## Corpus Check
- 371 files · ~654,894 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4824 nodes · 24828 edges · 172 communities (118 shown, 54 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2491 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f79139eb`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- database/sql.NullInt64
- context.Context
- vendorjs/main.go
- createProvider
- DownloadFiveTools
- SheetView
- combobox.js
- spellBoxes
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
- Input
- formSelect
- app.js
- tabbable.js
- useDismiss
- Slugify
- who
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
- NameDesc
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
- ImagePage
- Part
- radiogroup.js
- Player
- ParseEffects
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- State
- use_transition_status.js
- playwright
- ftPrerequisite
- Store
- checkbox.js
- ParseAidedd
- switch.js
- empty_templ.go
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
- Client
- nullID
- profFocus
- cert.go
- .source
- ftBook
- Pending
- BlockEditor
- Entity
- deathFix
- capReader
- Campaign
- InputForEntry
- build.sh
- fakeHandler
- looseString
- errBodyWriter
- TestDeviceAnnouncements
- export/markdown.go
- Before/after pairs
- .sendFeedback
- jsconfig.json
- open5e.go
- testing.M
- TestPagesRender
- perm/CLAUDE.md
- WeaponKind
- ASIError
- Book
- .Entry
- sourceUsedError
- Device
- .q
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
- Queries
- Option
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 376 edges
2. `Path()` - 272 edges
3. `newTestEnv()` - 177 edges
4. `Sheet` - 134 edges
5. `Button()` - 123 edges
6. `CN()` - 114 edges
7. `SheetView` - 79 edges
8. `IfElse()` - 66 edges
9. `Command` - 60 edges
10. `Queries` - 57 edges

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

## Communities (172 total, 54 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (55): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+47 more)

### Community 1 - "database/sql.NullInt64"
Cohesion: 0.02
Nodes (72): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Condition, CreateConditionParams, ListCombatantsRow, EffectEnd (+64 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (76): keepAwake(), Attachment, CompendiumSource, Note, CreateCombatantParams, EntityKind, NoteKind, lanHostname() (+68 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "DownloadFiveTools"
Cohesion: 0.10
Nodes (29): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+21 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (112): UsesLabel(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+104 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "spellBoxes"
Cohesion: 0.67
Nodes (4): spellBoxes(), TestTaughtPicked(), taughtPicked(), taughtRules()

### Community 9 - "newTestEnv"
Cohesion: 0.04
Nodes (116): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount() (+108 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (46): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+38 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (54): Amount, CondOpKind, Effects, userText, ArmorKindNamed(), parseFeature(), parseProf(), parseProfSet() (+46 more)

### Community 13 - "dialog_templ.go"
Cohesion: 0.13
Nodes (37): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+29 more)

### Community 14 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (58): ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+50 more)

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
Nodes (24): SetCookie(), FromContext(), pathCampaign(), collectHeld(), formIntQuery(), isHTMX(), redirect(), homebrewTabURL() (+16 more)

### Community 21 - "WizardView"
Cohesion: 0.09
Nodes (49): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Action(), Alert() (+41 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (19): nodeActions(), openWeb(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages (+11 more)

### Community 23 - "nullID"
Cohesion: 0.07
Nodes (28): Combatant, Encounter, EffectMods(), ValidateEffectMods(), combatantRow(), passTurn(), concentrationIDs(), noteHeld() (+20 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (162): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+154 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "viewerCtx"
Cohesion: 0.25
Nodes (6): notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView, viewerCtx

### Community 28 - "Input"
Cohesion: 0.07
Nodes (56): Uses, Combobox(), Content(), Item(), List(), Field(), Label(), Input() (+48 more)

### Community 29 - "formSelect"
Cohesion: 0.06
Nodes (48): ImageCategory, Defenses, Footer(), Content(), Item(), state(), Trigger(), firstValue() (+40 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Slugify"
Cohesion: 0.05
Nodes (49): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), Slugify(), AbilityKeys(), AddResistance(), FreeCast() (+41 more)

### Community 34 - "who"
Cohesion: 0.08
Nodes (19): HomebrewOwner, HomebrewInput, SaveHomebrew(), SeedHabitats(), targetColumns(), LinkTarget, compendiumBack(), compendiumPath() (+11 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "time.Time"
Cohesion: 0.05
Nodes (26): Identity, Handlers, browserName(), firstLine(), mdLinkText(), TestUserAgent(), userAgentSummary(), ensureCategory() (+18 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (23): block, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (61): TestHomebrew(), CoreTraits(), dataFor(), Decode(), DecodeData(), SpellData, TextData, ValidKind() (+53 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.09
Nodes (39): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+31 more)

### Community 42 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 43 - "Session"
Cohesion: 0.08
Nodes (24): fixture, tokens, constantTimeEqual(), credentialsMatch(), newFixture(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity() (+16 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 46 - "CustomMod"
Cohesion: 0.09
Nodes (19): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), AddFeat(), SaveFeature() (+11 more)

### Community 47 - "NameDesc"
Cohesion: 0.17
Nodes (20): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+12 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.03
Nodes (152): Props, Variant, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, ListSharedWithRow (+144 more)

### Community 49 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (131): Defense, TestIsLoopback(), TestDecode(), TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), TestPageClientRedirect(), TestArmorTrainingGrantsAndRemoval() (+123 more)

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
Cohesion: 0.21
Nodes (6): IsLoopback(), ClearCookie(), writeCookie(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "ImagePage"
Cohesion: 0.07
Nodes (40): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListBacklinksRow, Entity, Field, Look, attrString() (+32 more)

### Community 59 - "Part"
Cohesion: 0.06
Nodes (27): Part, Skill, ItemData, TestCustomMods(), TestModifier(), ArmorEffect(), skillByKey(), Sum() (+19 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.08
Nodes (38): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkImportReview(), BenchmarkSheetState(), testEnv (+30 more)

### Community 62 - "ParseEffects"
Cohesion: 0.11
Nodes (13): LevelUpChoice, TestArmorGrants(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+5 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (43): Content(), Description(), Header(), state(), Title(), Trigger(), triggerID(), effectBits() (+35 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "State"
Cohesion: 0.11
Nodes (12): event, Handler, patchMessage, stateHandler, viewPatch, Topic, TopicKind, Diff() (+4 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 72 - "Store"
Cohesion: 0.06
Nodes (54): linkTarget, migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+46 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "empty_templ.go"
Cohesion: 0.07
Nodes (52): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, SearchReadableBlocksRow (+44 more)

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
Cohesion: 0.24
Nodes (9): ftCopyRef, ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve() (+1 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "NewClient"
Cohesion: 0.14
Nodes (18): isDevice(), New(), NewClient(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+10 more)

### Community 96 - "fields.js"
Cohesion: 0.21
Nodes (13): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+5 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "storeSource"
Cohesion: 0.11
Nodes (20): Site, Source, sourceFormat, TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), formatOf(), Report, ImportProgress() (+12 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Client"
Cohesion: 0.20
Nodes (5): message, Client, Hub, drain(), TestCloseSessionAndCharacter()

### Community 103 - "nullID"
Cohesion: 0.10
Nodes (38): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), acBonusForm(), effectFix, newEffectFix(), notes() (+30 more)

### Community 105 - "cert.go"
Cohesion: 0.10
Nodes (19): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+11 more)

### Community 106 - ".source"
Cohesion: 0.13
Nodes (16): UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceKeyWithSlash() (+8 more)

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (20): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftParts(), ftPrerequisite() (+12 more)

### Community 108 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 111 - "Entity"
Cohesion: 0.15
Nodes (8): Entity, entityTopics(), linkEntityLook(), linkCharacterLook(), lookFree(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams

### Community 112 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 116 - "Campaign"
Cohesion: 0.16
Nodes (7): ListAttachmentFilesRow, Campaign, fresh(), inCampaign(), Roster, CreateCampaignParams, CampaignSwitch

### Community 117 - "InputForEntry"
Cohesion: 0.22
Nodes (8): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "TestDeviceAnnouncements"
Cohesion: 0.33
Nodes (4): ClientInfo, goneDevice, TestFiveToolsClassLevelUp(), TestDeviceAnnouncements()

### Community 127 - "export/markdown.go"
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
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 135 - "TestPagesRender"
Cohesion: 0.40
Nodes (4): Gotchas, internal/web, Notes, TestPagesRender()

### Community 137 - "WeaponKind"
Cohesion: 0.50
Nodes (4): Completions(), parseMastery(), TestWeaponKind(), WeaponKind()

### Community 139 - "Book"
Cohesion: 0.12
Nodes (23): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+15 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 144 - "Device"
Cohesion: 0.39
Nodes (5): devicesWhere(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 146 - ".q"
Cohesion: 0.10
Nodes (28): isCleared(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+20 more)

### Community 175 - "Queries"
Cohesion: 0.07
Nodes (42): ClassFeature, featureSwitch, ClassData, dnd5eFeatures(), bindFiveToolsImages(), Queries, Env, TestSlotMaxima() (+34 more)

### Community 176 - "Option"
Cohesion: 0.10
Nodes (19): LevelUpPlan, SpellColumn(), spellNeed(), TestSpellNeed(), spellChecklist(), TestSpellCounts(), LevelUpView, Option (+11 more)

## Knowledge Gaps
- **40 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+35 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 452 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **54 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `database/sql.NullInt64`, `context.Context`, `Book`, `Command`, `perm.go`, `.q`, `nullID`, `viewerCtx`, `Slugify`, `who`, `time.Time`, `CompendiumKind`, `Queries`, `Option`, `testing.T`, `Session`, `FiveToolsJob`, `State`, `Store`, `Client`, `cert.go`, `Pending`, `Entity`, `Campaign`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _40 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14003926761335098 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `newTestEnv` to `compendium_templ.go`, `.sendFeedback`, `SheetView`, `TestPagesRender`, `dialog_templ.go`, `perm.go`, `.q`, `App`, `WizardView`, `nullID`, `github.com/a-h/templ.Attributes`, `Input`, `formSelect`, `who`, `time.Time`, `CompendiumKind`, `tabs_templ.go`, `Session`, `github.com/a-h/templ.Component`, `testing.T`, `ImagePage`, `Player`, `CombatantView`, `Store`, `empty_templ.go`, `ParseWiki`, `nullID`, `.source`, `Entity`, `deathFix`, `Campaign`, `export/markdown.go`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Should `database/sql.NullInt64` be split into smaller, more focused modules?**
  _Cohesion score 0.022193548387096775 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `newTestEnv`, `Command`, `wikilink.js`, `TestPagesRender`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._