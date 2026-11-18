# Graph Report - gm-session-manager  (2026-11-18)

## Corpus Check
- 354 files · ~622,464 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4549 nodes · 22880 edges · 138 communities (111 shown, 27 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2316 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `4f7aff8a`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- entityForm
- createProvider
- github.com/a-h/templ.Component
- SheetView
- combobox.js
- redirect.go
- Input
- encoding/json.RawMessage
- select.js
- Command
- database/sql.NullInt64
- sourcesEnv
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- writeBook
- WizardView
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Session
- Sheet
- Client
- app.js
- tabbable.js
- useDismiss
- CompendiumKind
- ParseWikidot
- command.js
- dialog.js
- Book
- Import
- tooltip.js
- EquipmentOption
- .do
- mirror
- .sourceIssues
- composite.js
- .fail
- Feature
- Store
- nullID
- .sendFeedback
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Layout
- golangci-lint Config
- names
- Item
- radiogroup.js
- InputForEntry
- capReader
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- .compendiumDir
- use_transition_status.js
- playwright
- release.sh
- SkillChoice
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
- ListSharedWithRow
- Lint and Release Rules
- Run Script
- coinList
- useTypeahead
- vendorjs/main.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- .editEntity
- mounted
- newLiveFixture
- profFocus
- EffectFieldsURL
- buildReview
- ftBook
- Condition
- config.go
- github.com/a-h/templ.Component
- export/markdown.go
- build.sh
- Queries
- looseString
- errBodyWriter
- ParseEffects
- Before/after pairs
- viewerCtx
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftItemData
- ASIError
- .Entry
- sourceUsedError
- Parse
- CampaignSources
- time.Time
- Weapon

## God Nodes (most connected - your core abstractions)
1. `App` - 369 edges
2. `newTestEnv()` - 149 edges
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

## Communities (138 total, 27 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.15
Nodes (24): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+16 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (77): keepAwake(), keepAwake(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListImageEntitiesRow, ListAttachmentFilesRow, CompendiumSource (+69 more)

### Community 3 - "entityForm"
Cohesion: 0.12
Nodes (25): ListEntityFieldBacklinksRow, Entity, Field, Look, attrString(), generateSVG(), getIconContent(), Props (+17 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "github.com/a-h/templ.Component"
Cohesion: 0.13
Nodes (39): ctxKey, ctxState, Props, Dialog(), Close(), Content(), Description(), Dialog() (+31 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (108): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+100 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Input"
Cohesion: 0.06
Nodes (71): ImageCategory, ListBacklinksRow, Defenses, UI conventions, Description(), Field(), Group(), Label() (+63 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (40): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+32 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (59): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+51 more)

### Community 13 - "database/sql.NullInt64"
Cohesion: 0.04
Nodes (37): ListHomebrewByOwnerRow, GetImageRow, Attachment, firstLine(), mdLinkText(), saveScreenshot(), ensureCategory(), imageItem() (+29 more)

### Community 14 - "sourcesEnv"
Cohesion: 0.28
Nodes (7): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (46): ListCharacterNamesRow, Note, GetBlockRow, NoteKind, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities() (+38 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "writeBook"
Cohesion: 0.27
Nodes (9): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpBrokenSettings(), TestStartUpRedownloads(), writeBook(), appState (+1 more)

### Community 21 - "WizardView"
Cohesion: 0.05
Nodes (76): ActionProps, DescriptionProps, Props, TitleProps, Variant, LevelUpPlan, Action(), Alert() (+68 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (39): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+31 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (36): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Campaign, compendiumBack(), compendiumPath() (+28 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (197): Props, Props, Variant, Props, Size, ActionProps, ContentProps, DescriptionProps (+189 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Session"
Cohesion: 0.10
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+10 more)

### Community 28 - "Sheet"
Cohesion: 0.05
Nodes (63): internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc (+55 more)

### Community 29 - "Client"
Cohesion: 0.06
Nodes (34): ClientInfo, event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch (+26 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "CompendiumKind"
Cohesion: 0.04
Nodes (46): Site, CoreTraits(), dataFor(), DecodeData(), ValidKind(), HomebrewInput, FetchPage(), Page (+38 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Book"
Cohesion: 0.10
Nodes (32): rawBook, parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook() (+24 more)

### Community 38 - "Import"
Cohesion: 0.09
Nodes (40): sourceFormat, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+32 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "EquipmentOption"
Cohesion: 0.25
Nodes (8): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), applyEquipment(), TestApplyEquipmentMergesGear(), EquipmentChoice, gear

### Community 41 - ".do"
Cohesion: 0.07
Nodes (30): main(), testEnv, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+22 more)

### Community 42 - "mirror"
Cohesion: 0.29
Nodes (6): bookWriter, mirror, bookFolder(), downloadBook(), filterDoc(), filterSource()

### Community 43 - ".sourceIssues"
Cohesion: 0.36
Nodes (7): pinnedRefs(), sourceIssue(), featKey(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".fail"
Cohesion: 0.07
Nodes (26): ListSubclassesOfRow, Character, Character, ParseSheet(), CharacterFile, NewCharacterFile(), CanChangeSubclass(), CanCharacter() (+18 more)

### Community 46 - "Feature"
Cohesion: 0.10
Nodes (14): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), UsesLabel(), CustomMod (+6 more)

### Community 47 - "Store"
Cohesion: 0.11
Nodes (25): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+17 more)

### Community 48 - "nullID"
Cohesion: 0.11
Nodes (33): DeathSaves, nullID(), deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP() (+25 more)

### Community 49 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (164): TestIsLoopback(), TestDecode(), TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), Completions(), Verbs(), TestCompletions() (+156 more)

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
Cohesion: 0.10
Nodes (18): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+10 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Layout"
Cohesion: 0.04
Nodes (118): CountCompendiumRow, CountHabitatsRow, ListEntryCopiesRow, ListHomebrewRow, ListMonstersByNameRow, SearchCompendiumRow, Habitat, Badge() (+110 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.07
Nodes (29): deny, files, rules, generated, paths, rules, formatters, enable (+21 more)

### Community 58 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 59 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "InputForEntry"
Cohesion: 0.20
Nodes (9): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), bindFiveToolsImages(), InputForEntry(), TestParseFiveTools() (+1 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.29
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

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

### Community 68 - ".compendiumDir"
Cohesion: 0.13
Nodes (8): IsLoopback(), RequireAdmin(), ImportSummary(), fiveToolsKey(), sameOriginOnly(), useValidateTag(), sourceInUseMsg(), sourceUsers()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "SkillChoice"
Cohesion: 0.67
Nodes (3): SkillChoice(), TestSkillChoice(), skillChoice()

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
Cohesion: 0.09
Nodes (41): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+33 more)

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

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "coinList"
Cohesion: 0.67
Nodes (3): Coins, coinList(), coin

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 98 - "live.js"
Cohesion: 0.50
Nodes (3): bindings(), roots(), within()

### Community 100 - ".editEntity"
Cohesion: 0.10
Nodes (15): Entity, LinkTarget, Fields, ParseFields(), entityTopics(), saveFieldLinks(), toEntity(), AddBlockLinkParams (+7 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "newLiveFixture"
Cohesion: 0.12
Nodes (26): Player(), BenchmarkEncounterState(), BenchmarkSheetState(), Gotchas, internal/web, Notes, Websocket and live views, isRandomID() (+18 more)

### Community 105 - "EffectFieldsURL"
Cohesion: 0.33
Nodes (4): EffectFieldsURL(), EffectID(), effectQuery(), EffectTarget

### Community 106 - "buildReview"
Cohesion: 0.10
Nodes (23): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+15 more)

### Community 107 - "ftBook"
Cohesion: 0.15
Nodes (19): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftParts(), ftRecords() (+11 more)

### Community 109 - "Condition"
Cohesion: 0.04
Nodes (61): Combatant, Condition, CreateCombatantParams, CreateConditionParams, ListCombatantsRow, EffectMods(), TestValidateEffectMods(), TestWithEffects() (+53 more)

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "github.com/a-h/templ.Component"
Cohesion: 0.07
Nodes (67): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Content() (+59 more)

### Community 117 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "Queries"
Cohesion: 0.08
Nodes (35): featureSwitch, Decode(), Queries, Env, ClassInfo, ClassLevel, WornArmor, featureNames() (+27 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.10
Nodes (13): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+5 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "viewerCtx"
Cohesion: 0.16
Nodes (9): PendingState, Pending, PendingKind, notFoundIsNil(), parseLiveKey(), liveKey, liveKind, liveView (+1 more)

### Community 133 - "open5e.go"
Cohesion: 0.09
Nodes (30): Entry, fixture, fixtures, ftDistance, HomebrewOwner, normalizer, o5eProperty, Source (+22 more)

### Community 137 - "ftItemData"
Cohesion: 0.11
Nodes (17): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ItemData, ftArmor(), ftCode() (+9 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 146 - "Parse"
Cohesion: 0.06
Nodes (37): Defense, dnd5eFeatures(), feature(), fiveToolsClasses(), TestFiveToolsClasses(), TestFiveToolsClassLevelUp(), Heal(), Parse() (+29 more)

### Community 148 - "CampaignSources"
Cohesion: 0.16
Nodes (16): campaignSourceRow(), CampaignSources(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), CampaignOverview(), CampaignRoster(), rosterCard() (+8 more)

### Community 150 - "time.Time"
Cohesion: 0.06
Nodes (23): ListImagesRow, Handlers, browserName(), TestUserAgent(), userAgentSummary(), csvSource(), parseCSV(), keepsBody() (+15 more)

### Community 156 - "Weapon"
Cohesion: 0.07
Nodes (15): Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey(), Sum() (+7 more)

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 402 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **27 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `viewerCtx`, `Command`, `database/sql.NullInt64`, `Viewer`, `writeBook`, `time.Time`, `Sheet`, `Client`, `CompendiumKind`, `.do`, `.sourceIssues`, `.fail`, `cert.go`, `Layout`, `FiveToolsJob`, `.compendiumDir`, `.editEntity`, `Condition`, `Queries`?**
  _High betweenness centrality (0.057) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14602642648932643 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `SheetView`, `Input`, `encoding/json.RawMessage`, `Command`, `WizardView`, `Weapon`, `CompendiumKind`, `EquipmentOption`, `.do`, `.sourceIssues`, `.fail`, `Feature`, `nullID`, `testing.T`, `Item`, `coinList`, `Condition`, `Queries`, `ParseEffects`?**
  _High betweenness centrality (0.013) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010169762972453556 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `Queries` to `compendium_templ.go`, `WizardView`, `Sheet`, `ParseEffects`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `entityForm` be split into smaller, more focused modules?**
  _Cohesion score 0.12 - nodes in this community are weakly interconnected._