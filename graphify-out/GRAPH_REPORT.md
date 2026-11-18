# Graph Report - gm-session-manager  (2026-11-18)

## Corpus Check
- 354 files · ~620,952 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4541 nodes · 22839 edges · 139 communities (111 shown, 28 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2305 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `3d084273`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Session
- context.Context
- SearchReadableBlocksRow
- createProvider
- github.com/a-h/templ.Component
- SheetView
- combobox.js
- redirect.go
- Input
- encoding/json.RawMessage
- select.js
- Command
- .Entry
- Parse
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Option
- WizardView
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- .editEntity
- CompendiumKind
- Client
- app.js
- tabbable.js
- useDismiss
- .campaignSourcesSet
- ParseWikidot
- command.js
- dialog.js
- Book
- NameDesc
- tooltip.js
- select_templ.go
- .do
- Book
- export/markdown.go
- composite.js
- monsterBody
- Feature
- Store
- nullID
- Note
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Layout
- golangci-lint Config
- EffectMods
- CompendiumPath
- radiogroup.js
- vendorjs/main.go
- toast_templ.go
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- .sendFeedback
- use_transition_status.js
- playwright
- release.sh
- TestRunCommandUsesCurrentSheet
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
- EquipmentOption
- Lint and Release Rules
- Run Script
- useTypeahead
- .fail
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- ftRec
- mounted
- time.Time
- newLiveFixture
- profFocus
- EffectFieldsURL
- buildReview
- ftBook
- Slugify
- Condition
- CampaignSources
- config.go
- Look
- build.sh
- Queries
- looseString
- errBodyWriter
- ParseEffects
- appState
- Before/after pairs
- viewerCtx
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- popover_templ.go
- sourceUsedError
- database/sql.NullInt64
- Weapon
- planSwitch

## God Nodes (most connected - your core abstractions)
1. `App` - 368 edges
2. `newTestEnv()` - 146 edges
3. `Sheet` - 130 edges
4. `Button()` - 122 edges
5. `CN()` - 113 edges
6. `SheetView` - 78 edges
7. `IfElse()` - 65 edges
8. `Command` - 60 edges
9. `CompendiumKind` - 56 edges
10. `Queries` - 53 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (139 total, 28 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (57): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+49 more)

### Community 1 - "Session"
Cohesion: 0.10
Nodes (16): tokens, constantTimeEqual(), credentialsMatch(), isCleared(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity(), TestRequireAdmin() (+8 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (86): keepAwake(), keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, CompendiumSource (+78 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "github.com/a-h/templ.Component"
Cohesion: 0.07
Nodes (75): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+67 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (103): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+95 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Input"
Cohesion: 0.06
Nodes (70): ImageCategory, Coins, Combobox(), Content(), Item(), List(), Description(), Field() (+62 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (50): ftClass, ftEntry, ftFeat, ftItem, ftModOp, ftMonster, ftNamed, ftRace (+42 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (53): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+45 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (29): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor() (+21 more)

### Community 14 - "Parse"
Cohesion: 0.09
Nodes (29): feature(), Heal(), Parse(), tokenize(), TestApplyHPMonsterUnaffected(), TestDeadRefusals(), TestDeathSaveFailKills(), TestHPCritParsing() (+21 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (54): ListCharacterNamesRow, GetBlockRow, ListCombatantsRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+46 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Option"
Cohesion: 0.11
Nodes (20): dnd5eFeatures(), TestFiveToolsClassLevelUp(), ClassFeature, LevelUpPlan, spellNeed(), TestSpellNeed(), groupViews(), TestClassOptionChoices() (+12 more)

### Community 21 - "WizardView"
Cohesion: 0.07
Nodes (55): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+47 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (39): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+31 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (41): SetCookie(), Identity, FromContext(), ListAttachmentFilesRow, Attachment, Campaign, compendiumBack(), compendiumPath() (+33 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (134): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+126 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - ".editEntity"
Cohesion: 0.11
Nodes (11): Entity, Fields, ParseFields(), entityTopics(), linkEntityLook(), saveFieldLinks(), lookFree(), SetEntityCurrentImageParams (+3 more)

### Community 28 - "CompendiumKind"
Cohesion: 0.05
Nodes (32): Site, dataFor(), DecodeData(), ItemData, fiveToolsPage(), SiteByKey(), UpsertCompendiumEntryParams, CompendiumEntry (+24 more)

### Community 29 - "Client"
Cohesion: 0.06
Nodes (35): ClientInfo, event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch (+27 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - ".campaignSourcesSet"
Cohesion: 0.08
Nodes (14): IsLoopback(), TestIsLoopback(), ClearCookie(), writeCookie(), ImportSummary(), fiveToolsKey(), assetETags(), immutableAsset() (+6 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.14
Nodes (23): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Book"
Cohesion: 0.13
Nodes (24): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestImportFiveToolsBook(), TestParseFiveToolsBook() (+16 more)

### Community 38 - "NameDesc"
Cohesion: 0.17
Nodes (19): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "select_templ.go"
Cohesion: 0.07
Nodes (37): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+29 more)

### Community 41 - ".do"
Cohesion: 0.13
Nodes (20): main(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm() (+12 more)

### Community 42 - "Book"
Cohesion: 0.17
Nodes (16): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+8 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "monsterBody"
Cohesion: 0.08
Nodes (39): SpellData, SearchCompendiumRow, classBadges(), compendiumRow(), copyOrder(), pickerRow(), rowBadges(), SearchComboItems() (+31 more)

### Community 46 - "Feature"
Cohesion: 0.11
Nodes (14): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+6 more)

### Community 47 - "Store"
Cohesion: 0.04
Nodes (79): fixture, main(), run(), HomebrewOwner, Source, sourceFormat, migration, newFixture() (+71 more)

### Community 48 - "nullID"
Cohesion: 0.07
Nodes (48): DeathSaves, nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), deathFix(), effectFix, TestDeathClearsConditions() (+40 more)

### Community 49 - "Note"
Cohesion: 0.14
Nodes (8): Note, NoteKind, setAudience(), CreateNoteParams, ListReadableNotesParams, RenameNoteParams, blockAccessForm, noteForm

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (157): TestDecode(), TestCoreTraits(), Import(), RemoveSourceDir(), TestRemoveSourceDir(), Completions(), Verbs(), TestCompletions() (+149 more)

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

### Community 56 - "Layout"
Cohesion: 0.07
Nodes (58): ListEntryCopiesRow, UI conventions, Scripts(), Separator(), campaignLink(), PendingRequests(), PendingGroup, CharacterPage() (+50 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.07
Nodes (29): deny, files, rules, generated, paths, rules, formatters, enable (+21 more)

### Community 58 - "EffectMods"
Cohesion: 0.22
Nodes (8): EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Gotchas, internal/web, Notes, TestPagesRender()

### Community 59 - "CompendiumPath"
Cohesion: 0.06
Nodes (72): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, CountHabitatsRow (+64 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 62 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.09
Nodes (39): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+31 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".sendFeedback"
Cohesion: 0.24
Nodes (14): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+6 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

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
Cohesion: 0.10
Nodes (38): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+30 more)

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
Cohesion: 0.18
Nodes (12): CompendiumBase(), wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki() (+4 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "EquipmentOption"
Cohesion: 0.29
Nodes (7): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), TestOptionText(), EquipmentChoice, optionText()

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".fail"
Cohesion: 0.05
Nodes (34): CoreTraits(), ListSubclassesOfRow, Character, Character, ParseSheet(), CharacterFile, NewCharacterFile(), CanChangeSubclass() (+26 more)

### Community 100 - "ftRec"
Cohesion: 0.21
Nodes (10): ftCopyRef, ftRec, ftCopyTargets(), ftDecode(), TestFiveToolsCopy(), ftMod(), ftName(), ftReplaceText() (+2 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "time.Time"
Cohesion: 0.06
Nodes (20): ListImagesRow, Handlers, csvSource(), keepsBody(), FiveToolsView, NewFiveToolsView(), BlockStep, LevelUp (+12 more)

### Community 103 - "newLiveFixture"
Cohesion: 0.13
Nodes (21): BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views, isRandomID(), shapeHash() (+13 more)

### Community 105 - "EffectFieldsURL"
Cohesion: 0.33
Nodes (4): EffectFieldsURL(), EffectID(), effectQuery(), EffectTarget

### Community 106 - "buildReview"
Cohesion: 0.12
Nodes (23): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+15 more)

### Community 107 - "ftBook"
Cohesion: 0.22
Nodes (10): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftRange(), ftRecords() (+2 more)

### Community 108 - "Slugify"
Cohesion: 0.06
Nodes (47): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), AbilityKeys(), AddFeat() (+39 more)

### Community 109 - "Condition"
Cohesion: 0.06
Nodes (37): Defense, Combatant, CreateCombatantParams, Defenses, ParseDefenses(), passTurn(), uniqueName(), concentrationIDs() (+29 more)

### Community 111 - "CampaignSources"
Cohesion: 0.10
Nodes (29): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), campaignSourceRow(), CampaignSources(), sourceGroupButton() (+21 more)

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "Look"
Cohesion: 0.17
Nodes (15): Entity, Field, Look, removeButton(), initials(), fieldRow(), lookControls(), looksGrid() (+7 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "Queries"
Cohesion: 0.10
Nodes (30): Decode(), Queries, TestSlotMaxima(), ClassInfo, ClassLevel, CasterLevel(), SlotMaxima(), SourceRef (+22 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.13
Nodes (10): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+2 more)

### Community 127 - "appState"
Cohesion: 0.32
Nodes (7): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpRedownloads(), appState, settings

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "viewerCtx"
Cohesion: 0.14
Nodes (10): PendingState, Pending, PendingKind, notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind (+2 more)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.14
Nodes (13): ArmorData, ftClassFeature, ftClassRefs, o5eArmor, ftArmor(), ftCode(), ftItemData(), TestFiveToolsItemMastery() (+5 more)

### Community 141 - "popover_templ.go"
Cohesion: 0.12
Nodes (21): Description(), Header(), state(), Title(), Trigger(), triggerID(), modTargetLabel(), effectBits() (+13 more)

### Community 150 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (47): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Condition, CreateConditionParams, EffectEnd, TestMalformedEffectSkipped(), ensureCategory() (+39 more)

### Community 156 - "Weapon"
Cohesion: 0.05
Nodes (28): Part, Skill, findItem(), parseArmor(), parseItem(), TestCustomMods(), TestModifier(), TestResistances() (+20 more)

### Community 166 - "planSwitch"
Cohesion: 0.14
Nodes (12): featureSwitch, ftRaceAbility(), featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch() (+4 more)

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 403 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **28 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `viewerCtx`, `Command`, `Viewer`, `Option`, `database/sql.NullInt64`, `.editEntity`, `CompendiumKind`, `Client`, `Weapon`, `.campaignSourcesSet`, `NameDesc`, `.do`, `Store`, `Note`, `FiveToolsJob`, `ParseWiki`, `.fail`, `time.Time`, `Slugify`, `Condition`, `Queries`, `appState`?**
  _High betweenness centrality (0.068) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1407221244868895 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Weapon` to `compendium_templ.go`, `context.Context`, `SheetView`, `Input`, `encoding/json.RawMessage`, `Command`, `Viewer`, `Option`, `database/sql.NullInt64`, `CompendiumKind`, `NameDesc`, `planSwitch`, `Feature`, `nullID`, `testing.T`, `TestRunCommandUsesCurrentSheet`, `.fail`, `Slugify`, `Condition`, `Queries`, `ParseEffects`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Should `Session` be split into smaller, more focused modules?**
  _Cohesion score 0.10317460317460317 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `Queries` to `compendium_templ.go`, `planSwitch`, `NameDesc`, `Command`, `Slugify`, `Option`, `Weapon`, `ParseEffects`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.00939766168417937 - nodes in this community are weakly interconnected._