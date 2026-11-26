# Graph Report - gm-session-manager  (2026-11-27)

## Corpus Check
- 382 files · ~670,613 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 58 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4922 nodes · 25671 edges · 185 communities (121 shown, 64 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2564 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `928d5da3`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Button
- context.Context
- vendorjs/main.go
- createProvider
- export/markdown.go
- SheetView
- combobox.js
- Path
- Store
- encoding/json.RawMessage
- select.js
- Command
- viewerCtx
- redirect.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- Input
- wikilink.js
- dialog_templ.go
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Sum
- CombatantView
- Condition
- app.js
- tabbable.js
- useDismiss
- Sheet
- github.com/a-h/templ.Component
- command.js
- dialog.js
- Handlers
- ParseWikidot
- tooltip.js
- Client
- tabs_templ.go
- WizardView
- State
- composite.js
- time.Time
- CustomMod
- run
- CompendiumView
- config.go
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- Open
- shadcn-templ Config
- ParseEffects
- golangci-lint Config
- Entity
- Weapon
- radiogroup.js
- Player
- nullID
- ParseArmorTraining
- sqlc Config
- monsterBody
- Session
- Claude Settings
- starterEnv
- use_transition_status.js
- playwright
- ftPrerequisite
- Book
- checkbox.js
- ParseAidedd
- switch.js
- Hub
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- mirror
- useTypeahead
- .source
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- CompendiumKind
- mounted
- pause_windows.go
- nullID
- profFocus
- cert.go
- .q
- ftBook
- .sendFeedback
- BlockEditor
- NewClient
- popover_templ.go
- capReader
- Amount
- ListSharedWithRow
- build.sh
- Pending
- looseString
- errBodyWriter
- sourcesEnv
- fakeHandler
- Before/after pairs
- ParseFiveTools
- jsconfig.json
- open5e.go
- testing.M
- .renderCompendium
- perm/CLAUDE.md
- names
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- parseCommand
- CampaignOverview
- condMods
- WeaponKind
- CharacterFile
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- spellBoxes
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
- .Router
- SearchReadableBlocksRow
- .ListBlockStepsSince
- .ListImageLookAccess
- player-app.sh
- .ListItemsByCategory
- .ListSpellsForClass
- IMPORT_FILTER_CAT_PREFIX
- .ListSubracesOf

## God Nodes (most connected - your core abstractions)
1. `App` - 388 edges
2. `Path()` - 278 edges
3. `newTestEnv()` - 179 edges
4. `Sheet` - 134 edges
5. `Button()` - 126 edges
6. `CN()` - 114 edges
7. `SheetView` - 79 edges
8. `IfElse()` - 67 edges
9. `Command` - 60 edges
10. `Queries` - 57 edges

## Surprising Connections (you probably didn't know these)
- `What to check` --references--> `check()`  [INFERRED]
  .claude/skills/browser-check/SKILL.md → cmd/vendorjs/main.go
- `Notes` --references--> `linkClick()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/assets/js/wikilink.js
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (185 total, 64 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (55): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+47 more)

### Community 1 - "Button"
Cohesion: 0.05
Nodes (92): ListEntryCopiesRow, ListMonstersByNameRow, Habitat, UI conventions, Description(), Button(), Checkbox(), Label() (+84 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (119): keepAwake(), linkTarget, linkTargets(), targetColumns(), ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow (+111 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.12
Nodes (26): Browser check, Roles, Saved states, The test instance, What to check, You are alone, canonical(), changedPaths() (+18 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.05
Nodes (47): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+39 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (112): Slot, Signed(), Variant, Action(), Card(), Content(), Description(), Header() (+104 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Path"
Cohesion: 0.04
Nodes (125): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount() (+117 more)

### Community 9 - "Store"
Cohesion: 0.15
Nodes (27): Source, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+19 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (53): ftClass, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+45 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.09
Nodes (44): CondOpKind, Effects, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn(), parseArmor() (+36 more)

### Community 13 - "viewerCtx"
Cohesion: 0.15
Nodes (11): Topic, TopicKind, heldCollected(), notFoundIsNil(), parseLiveKey(), WithCampaign(), heldTopics, liveKey (+3 more)

### Community 14 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (44): ListCharacterNamesRow, GetBlockRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanChangeSubclass() (+36 more)

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
Nodes (37): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Campaign, LinkTarget, downloadNames(), browserName() (+29 more)

### Community 21 - "Input"
Cohesion: 0.05
Nodes (80): ListImageCharactersRow, ListImageEntitiesRow, ImageCategory, ListBacklinksRow, Coins, Defenses, Description(), Field() (+72 more)

### Community 22 - "wikilink.js"
Cohesion: 0.17
Nodes (18): nodeActions(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages, LAYOUTS (+10 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.08
Nodes (52): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+44 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (198): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+190 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Sum"
Cohesion: 0.11
Nodes (10): findItem(), TestCustomMods(), TestExplainSums(), ArmorEffect(), Sum(), TestGearCommands(), ensureItem(), Item (+2 more)

### Community 28 - "CombatantView"
Cohesion: 0.09
Nodes (42): Header(), Title(), effectBits(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName() (+34 more)

### Community 29 - "Condition"
Cohesion: 0.07
Nodes (39): Combatant, Condition, CreateConditionParams, ListCombatantsRow, EffectEnd, combatantRow(), passTurn(), rollFor() (+31 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Sheet"
Cohesion: 0.03
Nodes (115): sourceFormat, LevelUpChoice, internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options() (+107 more)

### Community 34 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (96): ListEntityFieldBacklinksRow, Entity, Field, Look, Content(), Description(), Empty(), Header() (+88 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 38 - "ParseWikidot"
Cohesion: 0.14
Nodes (23): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.13
Nodes (32): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+24 more)

### Community 42 - "WizardView"
Cohesion: 0.08
Nodes (48): SpellColumn(), initialValue(), Item(), RadioGroup(), state(), abilitySelect(), levelUpForm(), spellChecklist() (+40 more)

### Community 43 - "State"
Cohesion: 0.13
Nodes (8): Handler, patchMessage, stateHandler, viewPatch, Diff(), State, TestDiff(), rosterHP()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "time.Time"
Cohesion: 0.07
Nodes (14): ListImagesRow, Attachment, firstLine(), folderName(), mdLinkText(), BlockStep, CreateAttachmentParams, LevelUp (+6 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (23): FeatureEffects, findMod(), newMod(), parseFeature(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature() (+15 more)

### Community 47 - "run"
Cohesion: 0.17
Nodes (9): main(), run(), openBrowser(), bookFolder(), CleanDownloads(), TestCleanDownloads(), LANAddresses(), New() (+1 more)

### Community 48 - "CompendiumView"
Cohesion: 0.06
Nodes (54): Props, Variant, CountEntriesBySourceRow, CountHabitatsRow, CountSharedBySourceRow, ListHomebrewLinksRow, SearchCompendiumRow, Badge() (+46 more)

### Community 49 - "config.go"
Cohesion: 0.18
Nodes (14): parseFlags(), TestParseFlags(), Config, lookupFunc, describe(), fromLookup(), Load(), newValidator() (+6 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (147): Defense, TestIsLoopback(), TestDecode(), TestCoreTraits(), TestSkillChoice(), TestTokensReadAsLabels(), feature(), RemoveSourceDir() (+139 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "Open"
Cohesion: 0.11
Nodes (28): migration, internal/db, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp() (+20 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ParseEffects"
Cohesion: 0.10
Nodes (15): TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+7 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Entity"
Cohesion: 0.11
Nodes (11): Entity, Fields, ParseFields(), entityTopics(), linkEntityLook(), saveFieldLinks(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams (+3 more)

### Community 59 - "Weapon"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.09
Nodes (34): Player(), TestCanRemoveCondition(), BenchmarkEncounterState(), BenchmarkSheetState(), Websocket and live views, TestStarterSets(), TestNoteViewAudiences(), isRandomID() (+26 more)

### Community 62 - "nullID"
Cohesion: 0.09
Nodes (27): EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), nullID(), Gotchas, internal/web, Notes (+19 more)

### Community 63 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "monsterBody"
Cohesion: 0.19
Nodes (21): commaPrefixed(), compendiumFindURL(), levelsText(), parenthesized(), signedMap(), speeds(), classBody(), entryBody() (+13 more)

### Community 66 - "Session"
Cohesion: 0.09
Nodes (24): fixture, tokens, loopbackHostOnly(), TestLoopbackHostOnly(), constantTimeEqual(), credentialsMatch(), newFixture(), serve() (+16 more)

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "starterEnv"
Cohesion: 0.24
Nodes (12): TestFiveToolsBooks(), UseTestServer(), testEnv, pickerTicks(), serveStarterSets(), starterEnv(), TestImportsTickFromCampaignOnly(), TestPartialDownloadTicks() (+4 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 72 - "Book"
Cohesion: 0.11
Nodes (26): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+18 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Hub"
Cohesion: 0.23
Nodes (5): ClientInfo, event, goneDevice, Hub, TestCloseSessionAndCharacter()

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
Cohesion: 0.10
Nodes (8): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, openWeb(), randomID(), KEEP_IMAGES, LINK_PATHS

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - "mirror"
Cohesion: 0.20
Nodes (12): bookWriter, mirror, downloadBook(), fetchJSON(), filterDoc(), filterSource(), replaceDir(), TestMirrorGetMissingFile() (+4 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".source"
Cohesion: 0.12
Nodes (17): ParsePageURL(), TestParsePageURL(), UpsertCompendiumEntryParams, testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove() (+9 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "CompendiumKind"
Cohesion: 0.04
Nodes (42): CoreTraits(), dataFor(), DecodeData(), ValidKind(), Character, Character, CompendiumKind, Campaign() (+34 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "nullID"
Cohesion: 0.13
Nodes (29): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead() (+21 more)

### Community 105 - "cert.go"
Cohesion: 0.15
Nodes (10): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+2 more)

### Community 106 - ".q"
Cohesion: 0.08
Nodes (31): isCleared(), testEnv, testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict() (+23 more)

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (16): ftBook, ftCopyRef, ftRec, each(), ftCopyTargets(), ftDecode(), ftRecords(), TestFiveToolsCopy() (+8 more)

### Community 108 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 109 - "BlockEditor"
Cohesion: 0.22
Nodes (5): attachedPicture(), BlockEditor, json(), links(), sync()

### Community 111 - "NewClient"
Cohesion: 0.16
Nodes (20): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+12 more)

### Community 112 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 116 - "Amount"
Cohesion: 0.22
Nodes (7): Amount, userText, CritDamage(), Damage(), parseRevive(), Set(), TestApplyHPDeath()

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "ParseFiveTools"
Cohesion: 0.33
Nodes (5): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), ParseFiveTools(), TestParseFiveTools()

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 135 - ".renderCompendium"
Cohesion: 0.06
Nodes (29): HomebrewOwner, Site, TestHomebrew(), fiveToolsPage(), ftsQuery(), HomebrewInput, InputForEntry(), SaveHomebrew() (+21 more)

### Community 137 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 139 - "DownloadFiveTools"
Cohesion: 0.10
Nodes (28): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+20 more)

### Community 140 - ".Entry"
Cohesion: 0.09
Nodes (26): MonsterAction, addEntry(), setStat(), MonsterData, ftSkillNames(), TestFiveToolsCoreTraitsShapes(), HomebrewInputFor(), abilityKey() (+18 more)

### Community 143 - "parseCommand"
Cohesion: 0.40
Nodes (5): coinCommand(), commandText(), hpCommand(), parseCommand(), damageArgs()

### Community 144 - "CampaignOverview"
Cohesion: 0.44
Nodes (7): deviceRow(), rosterCard(), devicesWhere(), agoTime(), deviceLabel(), Device, RosterCharacter

### Community 145 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

### Community 147 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 148 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 157 - "spellBoxes"
Cohesion: 0.67
Nodes (4): spellBoxes(), TestTaughtPicked(), taughtPicked(), taughtRules()

### Community 175 - "planSwitch"
Cohesion: 0.23
Nodes (6): featureSwitch, featureNames(), SourceRef, SubclassRef, planSwitch(), sameSlug()

### Community 176 - ".Router"
Cohesion: 0.23
Nodes (4): IsLoopback(), RequireAdmin(), sameOriginOnly(), useValidateTag()

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 458 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **64 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Button`, `context.Context`, `.renderCompendium`, `Command`, `viewerCtx`, `perm.go`, `Condition`, `Sheet`, `State`, `time.Time`, `run`, `.Router`, `testing.T`, `Entity`, `Book`, `Hub`, `CompendiumKind`, `.q`, `Pending`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1451905082715772 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Button`, `context.Context`, `export/markdown.go`, `SheetView`, `.renderCompendium`, `perm.go`, `CampaignOverview`, `App`, `Input`, `dialog_templ.go`, `github.com/a-h/templ.Attributes`, `CombatantView`, `Condition`, `Sheet`, `github.com/a-h/templ.Component`, `tabs_templ.go`, `WizardView`, `CompendiumView`, `testing.T`, `Entity`, `Player`, `nullID`, `monsterBody`, `starterEnv`, `ParseWiki`, `.source`, `CompendiumKind`, `nullID`, `.q`, `.sendFeedback`, `sourcesEnv`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **Should `Button` be split into smaller, more focused modules?**
  _Cohesion score 0.04854045037531276 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `nullID`, `wikilink.js`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._