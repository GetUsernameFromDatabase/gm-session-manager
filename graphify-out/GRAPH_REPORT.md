# Graph Report - gm-session-manager  (2026-11-18)

## Corpus Check
- 354 files · ~622,565 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4554 nodes · 22885 edges · 145 communities (114 shown, 31 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2316 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `cab8c8da`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- NameDesc
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- redirect.go
- Input
- encoding/json.RawMessage
- select.js
- Command
- database/sql.NullInt64
- monsterBody
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- newTestEnv
- WizardView
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Session
- Slugify
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
- Option
- .do
- mirror
- NewClient
- composite.js
- ParseSheet
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
- CompendiumPath
- golangci-lint Config
- planSwitch
- Item
- radiogroup.js
- dialog_templ.go
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
- Hub
- checkbox.js
- ParseAidedd
- switch.js
- Site
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- ListSharedWithRow
- Lint and Release Rules
- Run Script
- ftItemData
- useTypeahead
- vendorjs/main.go
- Validator Rule
- Go Module
- ftRec
- live.js
- internal_backdrop.js
- .editEntity
- mounted
- Pending
- newLiveFixture
- profFocus
- EffectFieldsURL
- buildReview
- ftBook
- internal/web
- Queries
- Campaign
- config.go
- github.com/a-h/templ.Component
- toast_templ.go
- export/markdown.go
- build.sh
- Queries
- looseString
- errBodyWriter
- ParseEffects
- rowFigure
- Before/after pairs
- viewerCtx
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- keepawake_windows.go
- .Entry
- sourceUsedError
- RosterCharacter
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

## Communities (145 total, 31 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.16
Nodes (14): Description(), Header(), state(), Title(), triggerID(), Align, ContentProps, ctxKey (+6 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (77): keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Encounter, Note, SearchReadableBlocksRow (+69 more)

### Community 3 - "NameDesc"
Cohesion: 0.17
Nodes (19): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+11 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.09
Nodes (40): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+32 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (125): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+117 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Input"
Cohesion: 0.06
Nodes (70): ImageCategory, Defenses, Footer(), Description(), Field(), Label(), Input(), Textarea() (+62 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (39): ftClass, ftEntry, ftFeat, ftItem, ftModOp, ftMonster, ftNamed, ftRace (+31 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (58): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+50 more)

### Community 13 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (61): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Attachment, Condition, CreateConditionParams, ListCombatantsRow, EffectEnd (+53 more)

### Community 14 - "monsterBody"
Cohesion: 0.13
Nodes (27): SpellData, classBadges(), rowBadges(), alignmentColor(), commaPrefixed(), compendiumFindURL(), levelsText(), parenthesized() (+19 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (48): ListCharacterNamesRow, GetBlockRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock() (+40 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "newTestEnv"
Cohesion: 0.04
Nodes (68): activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestDeleteActiveCombatant(), TestDeleteActiveCombatantLast(), TestDeleteActiveCombatantSkipsOut(), TestDeleteCharacterActiveCombatant(), TestEncounterTransitions() (+60 more)

### Community 21 - "WizardView"
Cohesion: 0.07
Nodes (62): Props, Props, Description(), Checkbox(), initialChecked(), Combobox(), Content(), ItemData (+54 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (39): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+31 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (35): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), CanChangeSubclass(), CanCharacter(), compendiumBack() (+27 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (180): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+172 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Session"
Cohesion: 0.06
Nodes (27): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+19 more)

### Community 28 - "Slugify"
Cohesion: 0.08
Nodes (38): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), AbilityKeys(), AddFeat() (+30 more)

### Community 29 - "Client"
Cohesion: 0.10
Nodes (16): event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch, Client (+8 more)

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
Nodes (49): dataFor(), DecodeData(), ValidKind(), CountCompendiumRow, SearchCompendiumRow, UpsertCompendiumEntryParams, Character, CompendiumEntry (+41 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.14
Nodes (22): block, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster(), sourceLine() (+14 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Book"
Cohesion: 0.12
Nodes (24): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook() (+16 more)

### Community 38 - "Import"
Cohesion: 0.07
Nodes (47): Source, sourceFormat, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+39 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Option"
Cohesion: 0.14
Nodes (16): LevelUpPlan, Spellbook(), spellNeed(), subclassOptions(), groupViews(), withSpells(), LevelUpView, Option (+8 more)

### Community 41 - ".do"
Cohesion: 0.09
Nodes (27): main(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm() (+19 more)

### Community 42 - "mirror"
Cohesion: 0.14
Nodes (19): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+11 more)

### Community 43 - "NewClient"
Cohesion: 0.26
Nodes (17): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+9 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "ParseSheet"
Cohesion: 0.12
Nodes (31): ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpPageAtTwenty() (+23 more)

### Community 46 - "Feature"
Cohesion: 0.09
Nodes (19): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+11 more)

### Community 47 - "Store"
Cohesion: 0.15
Nodes (20): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+12 more)

### Community 48 - "nullID"
Cohesion: 0.10
Nodes (35): DeathSaves, nullID(), deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP() (+27 more)

### Community 49 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (109): Defense, TestDecode(), TestEquipmentOptions(), RemoveSourceDir(), TestRemoveSourceDir(), Completions(), Parse(), tokenize() (+101 more)

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
Cohesion: 0.11
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "CompendiumPath"
Cohesion: 0.07
Nodes (48): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, Cell(), BuilderPickItems(), BuilderPicks() (+40 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (31): deny, files, rules, generated, paths, rules, formatters, enable (+23 more)

### Community 58 - "planSwitch"
Cohesion: 0.15
Nodes (11): featureSwitch, featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch(), sameSlug() (+3 more)

### Community 59 - "Item"
Cohesion: 0.13
Nodes (14): EquipmentItem, EquipmentOptions(), EquipmentOption, findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem() (+6 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "dialog_templ.go"
Cohesion: 0.13
Nodes (37): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+29 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.23
Nodes (4): FiveToolsJob, appState, fiveToolsJob, settings

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (36): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+28 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".compendiumDir"
Cohesion: 0.20
Nodes (5): ImportSummary(), fiveToolsKey(), writeBook(), sourceInUseMsg(), sourceUsers()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.60
Nodes (4): die(), GOFLAGS, newer(), release.sh script

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Site"
Cohesion: 0.15
Nodes (11): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+3 more)

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
Cohesion: 0.20
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - "ftItemData"
Cohesion: 0.20
Nodes (10): ArmorData, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText() (+2 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 96 - "ftRec"
Cohesion: 0.24
Nodes (9): ftCopyRef, ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve() (+1 more)

### Community 98 - "live.js"
Cohesion: 0.50
Nodes (3): bindings(), roots(), within()

### Community 100 - ".editEntity"
Cohesion: 0.07
Nodes (19): Entity, LinkTarget, Fields, ParseFields(), entityTopics(), linkEntityLook(), saveFieldLinks(), AddBlockLinkParams (+11 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 103 - "newLiveFixture"
Cohesion: 0.15
Nodes (19): BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views, isRandomID(), shapeHash(), newLiveFixture(), TestLiveAffects() (+11 more)

### Community 106 - "buildReview"
Cohesion: 0.09
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 107 - "ftBook"
Cohesion: 0.19
Nodes (15): ftBook, each(), ftCastingTime(), ftParts(), ftRecords(), TestFiveToolsCopy(), crProficiency(), ftEntries() (+7 more)

### Community 108 - "internal/web"
Cohesion: 0.18
Nodes (8): Handlers, internal/web, Notes, keepsBody(), formErrors(), fieldError, pickError, userError

### Community 109 - "Queries"
Cohesion: 0.05
Nodes (41): HomebrewOwner, TestHomebrew(), Decode(), ftsQuery(), SaveHomebrew(), Search(), SeedHabitats(), ListSubclassesOfRow (+33 more)

### Community 111 - "Campaign"
Cohesion: 0.29
Nodes (6): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks()

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (142): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Field, Look, UI conventions, Content(), Description() (+134 more)

### Community 116 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 117 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 118 - "build.sh"
Cohesion: 0.53
Nodes (5): build(), CGO_ENABLED, GOFLAGS, build.sh script, step()

### Community 119 - "Queries"
Cohesion: 0.10
Nodes (29): ClassFeature, ClassData, dnd5eFeatures(), TestSlotMaxima(), ClassFeature, ClassInfo, ClassLevel, CasterLevel() (+21 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.13
Nodes (10): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+2 more)

### Community 127 - "rowFigure"
Cohesion: 0.50
Nodes (4): TestRowFigure(), crText(), rowFigure(), textOnly()

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (23): Entry, fixture, fixtures, ftDistance, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable() (+15 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 140 - ".Entry"
Cohesion: 0.07
Nodes (32): MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, habitatsFrom(), TestHabitatsFrom() (+24 more)

### Community 148 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 150 - "time.Time"
Cohesion: 0.05
Nodes (23): FiveToolsBooks(), ListAttachmentFilesRow, ListImagesRow, Campaign, browserName(), firstLine(), mdLinkText(), inCampaign() (+15 more)

### Community 156 - "Weapon"
Cohesion: 0.05
Nodes (28): Part, Skill, parseMastery(), TestCustomMods(), TestModifier(), TestResistances(), TestWeaponKind(), TestExplainSums() (+20 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 406 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **31 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `NameDesc`, `viewerCtx`, `Command`, `database/sql.NullInt64`, `Viewer`, `time.Time`, `Session`, `Slugify`, `Weapon`, `Client`, `CompendiumKind`, `Option`, `.do`, `cert.go`, `Item`, `FiveToolsJob`, `.compendiumDir`, `Hub`, `.editEntity`, `Pending`, `buildReview`, `Queries`, `Queries`?**
  _High betweenness centrality (0.056) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14654188541428897 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Weapon` to `compendium_templ.go`, `NameDesc`, `SheetView`, `Input`, `encoding/json.RawMessage`, `database/sql.NullInt64`, `Slugify`, `CompendiumKind`, `Option`, `.do`, `ParseSheet`, `Feature`, `nullID`, `testing.T`, `planSwitch`, `Item`, `Queries`, `Queries`, `ParseEffects`?**
  _High betweenness centrality (0.014) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010411332991978153 - nodes in this community are weakly interconnected._
- **Why does `fiveToolsJob` connect `FiveToolsJob` to `compendium_templ.go`, `App`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._