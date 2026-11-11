# Graph Report - gm-session-manager  (2026-11-11)

## Corpus Check
- 300 files · ~590,621 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4228 nodes · 18511 edges · 133 communities (109 shown, 24 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1899 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `1c17ea57`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- dialog_templ.go
- context.Context
- .do
- createProvider
- Sheet
- SheetView
- combobox.js
- redirect.go
- FilesPage
- Store
- select.js
- Command
- .Entry
- time.Time
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ftItemData
- levelUpForm
- wikilink.js
- Visibility
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- testing.T
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- Viewer
- command.js
- dialog.js
- LinkTarget
- Sheet
- tooltip.js
- Character
- ParseWikidot
- Book
- buildReview
- composite.js
- Parse
- Feature
- IfElse
- Item
- Client
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- sourcesEnv
- golangci-lint Config
- Open
- ItemData
- radiogroup.js
- vendorjs/main.go
- CN
- FiveToolsJob
- sqlc Config
- Topic
- Browser check
- Claude Settings
- .campaignSourcesSet
- use_transition_status.js
- playwright
- release.sh
- Session
- checkbox.js
- ParseAidedd
- switch.js
- WeaponKind
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- Feature
- Lint and Release Rules
- Run Script
- .Router
- useTypeahead
- parseCommand
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- joinNonEmpty
- mounted
- parseMod
- newLiveFixture
- DownloadFiveTools
- sameSlug
- .sendFeedback
- ftBook
- Verbs
- formErrors
- RosterCharacter
- capReader
- Pending
- .encounterView
- github.com/a-h/templ.Component
- build.sh
- .sourceIssues
- errBodyWriter
- buildReview
- Before/after pairs
- open5e.go
- spellCounts
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- Site
- TestRunCommandUsesCurrentSheet

## God Nodes (most connected - your core abstractions)
1. `App` - 320 edges
2. `newTestEnv()` - 129 edges
3. `Button()` - 122 edges
4. `Sheet` - 118 edges
5. `CN()` - 113 edges
6. `IfElse()` - 62 edges
7. `SheetView` - 57 edges
8. `CompendiumKind` - 53 edges
9. `Command` - 49 edges
10. `Input()` - 47 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `UI conventions` --references--> `selectItems()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go
- `UI conventions` --references--> `entryHref()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/compendium_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (133 total, 24 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (60): contextKey, classRef, indexed, looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature (+52 more)

### Community 1 - "dialog_templ.go"
Cohesion: 0.08
Nodes (47): Description(), Header(), state(), Title(), Trigger(), triggerID(), beingBadge(), combatantChip() (+39 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (107): keepAwake(), keepAwake(), SavePage(), ListImageCharactersRow, ListHomebrewLinksRow, ListSharedWithRow, UpsertCompendiumEntryParams, ListImageEntitiesRow (+99 more)

### Community 3 - ".do"
Cohesion: 0.15
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Sheet"
Cohesion: 0.09
Nodes (12): Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield(), damageLabel() (+4 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (110): ArmorEffect(), Coins, Slot, Signed(), Field, Variant, Action(), Card() (+102 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "FilesPage"
Cohesion: 0.07
Nodes (48): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent(), Props (+40 more)

### Community 10 - "Store"
Cohesion: 0.08
Nodes (43): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+35 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.10
Nodes (49): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+41 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (29): MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, ValidKind(), TestTokensReadAsLabels() (+21 more)

### Community 14 - "time.Time"
Cohesion: 0.07
Nodes (16): Identity, ListAttachmentFilesRow, browserName(), firstLine(), mdLinkText(), inCampaign(), TestUserAgent(), userAgentSummary() (+8 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.23
Nodes (17): CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock(), CanManageNote(), CanReadAttachment(), CanReadBlock(), CanReadImage() (+9 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ftItemData"
Cohesion: 0.24
Nodes (9): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+1 more)

### Community 21 - "levelUpForm"
Cohesion: 0.07
Nodes (63): Props, UI conventions, Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup() (+55 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "Visibility"
Cohesion: 0.14
Nodes (12): Visibility, BlockVisibilities(), TestBlockVisibilities(), ValidBlockVisibility(), visibilityLabel(), CreateBlockParams, ListBlockDocsRow, ListExportBlocksRow (+4 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (191): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Variant (+183 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.12
Nodes (31): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage() (+23 more)

### Community 28 - "App"
Cohesion: 0.04
Nodes (28): SetCookie(), FromContext(), Campaign, CanCharacter(), compendiumBack(), compendiumPath(), homebrewTabURL(), pathCampaign() (+20 more)

### Community 29 - "testing.T"
Cohesion: 0.03
Nodes (127): RemoveSourceDir(), TestRemoveSourceDir(), TestCompletions(), TestLevelUp(), TestLevelUpSubclassAndPools(), TestLevelUpThirdCaster(), TestMasteryLimit(), TestDedupeFeatures() (+119 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (44): GetImageRow, ListImagesRow, Attachment, Condition, Entity, GetBlockRow, CreateCombatantParams, EntityKind (+36 more)

### Community 34 - "Viewer"
Cohesion: 0.18
Nodes (11): VisibleEntity(), VisibleLooks(), CanChangeSubclass(), CanSeeCombatantHP(), Viewer, TestCanChangeSubclass(), TestCanSeeCombatantHP(), TestVisibleEntity() (+3 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "LinkTarget"
Cohesion: 0.13
Nodes (11): LinkTarget, Fields, ParseFields(), saveFieldLinks(), AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams (+3 more)

### Community 38 - "Sheet"
Cohesion: 0.03
Nodes (131): EquipmentItem, DeathSaves, featureSwitch, SetupSnapshot, Cantrips(), EquipmentOptions(), firstTable(), EquipmentOption (+123 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Character"
Cohesion: 0.11
Nodes (12): Character, Combatant, Character, Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile() (+4 more)

### Community 41 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 42 - "Book"
Cohesion: 0.10
Nodes (31): bookWriter, mirror, rawBook, internal/compendium, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools() (+23 more)

### Community 43 - "buildReview"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Parse"
Cohesion: 0.10
Nodes (21): feature(), TestFiveToolsClasses(), Parse(), tokenize(), fighter(), TestCommands(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+13 more)

### Community 46 - "Feature"
Cohesion: 0.14
Nodes (7): LevelUpChoice, TestSlotMaxima(), applyASI(), hpGain(), raise(), CasterLevel(), SlotMaxima()

### Community 47 - "IfElse"
Cohesion: 0.04
Nodes (124): dataFor(), CountCompendiumRow, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewByOwnerRow, ListHomebrewRow, SearchCompendiumRow (+116 more)

### Community 48 - "Item"
Cohesion: 0.26
Nodes (5): findItem(), parseArmor(), TestGearCommands(), ensureItem(), Item

### Community 49 - "Client"
Cohesion: 0.06
Nodes (36): ClientInfo, event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch (+28 more)

### Community 50 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

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
Cohesion: 0.09
Nodes (23): main(), run(), Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap() (+15 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 57 - "golangci-lint Config"
Cohesion: 0.10
Nodes (21): generated, paths, rules, formatters, enable, exclusions, settings, match-constant (+13 more)

### Community 58 - "Open"
Cohesion: 0.14
Nodes (16): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+8 more)

### Community 59 - "ItemData"
Cohesion: 0.10
Nodes (24): ImageCategory, firstValue(), modTargetLabel(), usesItems(), weaponAbilities(), conditionItems(), selectItems(), kindSelect() (+16 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "CN"
Cohesion: 0.05
Nodes (95): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+87 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "Topic"
Cohesion: 0.33
Nodes (4): TopicKind, Topic, parseLiveKey(), liveKey

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".campaignSourcesSet"
Cohesion: 0.21
Nodes (6): ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), FiveToolsView, NewFiveToolsView()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Session"
Cohesion: 0.10
Nodes (19): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+11 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "WeaponKind"
Cohesion: 0.50
Nodes (4): Completions(), parseMastery(), TestWeaponKind(), WeaponKind()

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
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "Feature"
Cohesion: 0.14
Nodes (10): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+2 more)

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - ".Router"
Cohesion: 0.22
Nodes (5): ClearCookie(), writeCookie(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "parseCommand"
Cohesion: 0.40
Nodes (5): featureCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 100 - "joinNonEmpty"
Cohesion: 0.09
Nodes (38): ftBookInfo, ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster (+30 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "parseMod"
Cohesion: 0.67
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 103 - "newLiveFixture"
Cohesion: 0.12
Nodes (17): Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveSourcesAlert() (+9 more)

### Community 104 - "DownloadFiveTools"
Cohesion: 0.13
Nodes (19): HomebrewOwner, Source, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition() (+11 more)

### Community 106 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 107 - "ftBook"
Cohesion: 0.12
Nodes (30): ftBook, habitatsFrom(), each(), ftCastingTime(), ftCell(), ftParts(), ftRecords(), ftTable() (+22 more)

### Community 110 - "formErrors"
Cohesion: 0.40
Nodes (3): fieldError, pickError, userError

### Community 111 - "RosterCharacter"
Cohesion: 0.36
Nodes (7): CampaignRoster(), Roster, rosterCard(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 114 - "Pending"
Cohesion: 0.29
Nodes (4): PendingState, Pending, PendingKind, viewerCtx

### Community 115 - ".encounterView"
Cohesion: 0.19
Nodes (13): Defense, ListCombatantsRow, Defenses, ParseDefenses(), combatantDefenses(), damageChoice(), initiativeMod(), outOfFight() (+5 more)

### Community 117 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (106): MediaProps, MediaVariant, ListMonstersByNameRow, Content(), Description(), Empty(), Header(), Media() (+98 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".sourceIssues"
Cohesion: 0.60
Nodes (4): sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 126 - "buildReview"
Cohesion: 0.31
Nodes (5): ensureCategory(), buildReview(), formValues(), csvImport, csvImports

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 134 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 137 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Site"
Cohesion: 0.17
Nodes (9): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SiteByKey() (+1 more)

### Community 163 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 390 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **24 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `time.Time`, `perm.go`, `testing.T`, `database/sql.NullInt64`, `Viewer`, `Sheet`, `Character`, `IfElse`, `Client`, `cert.go`, `Open`, `FiveToolsJob`, `Topic`, `.campaignSourcesSet`, `.Router`, `RosterCharacter`, `.encounterView`, `github.com/a-h/templ.Component`, `.sourceIssues`, `buildReview`?**
  _High betweenness centrality (0.045) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10816816682004893 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `TestRunCommandUsesCurrentSheet`, `Sheet`, `SheetView`, `Character`, `WeaponKind`, `Feature`, `Item`, `Feature`, `.encounterView`, `levelUpForm`, `.sourceIssues`, `ParseSheet`, `testing.T`?**
  _High betweenness centrality (0.016) - this node is a cross-community bridge._
- **Should `dialog_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.07704918032786885 - nodes in this community are weakly interconnected._
- **Why does `feature()` connect `Parse` to `compendium_templ.go`, `ParseWikidot`, `Book`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.008363605643321788 - nodes in this community are weakly interconnected._