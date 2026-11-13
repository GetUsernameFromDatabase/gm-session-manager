# Graph Report - gm-session-manager  (2026-11-13)

## Corpus Check
- 301 files · ~591,301 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4236 nodes · 18627 edges · 141 communities (114 shown, 27 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1903 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `515ecc82`
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
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- github.com/gin-gonic/gin.Context
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- CommandBuilder
- command.js
- dialog.js
- LinkTarget
- Slugify
- tooltip.js
- Character
- ParseWikidot
- Book
- buildReview
- composite.js
- testing.T
- Feature
- github.com/a-h/templ.Component
- CompendiumKind
- State
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- DecodeData
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
- tabs_templ.go
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- Feature
- Lint and Release Rules
- Run Script
- .compendiumWhole
- useTypeahead
- .commandFailed
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- select_templ.go
- newLiveFixture
- .found
- NewClient
- .sendFeedback
- ftBook
- Campaign
- Client
- formErrors
- RosterCharacter
- config.go
- Hub
- Pending
- .encounterView
- writeBook
- .loadLive
- build.sh
- .sourceIssues
- DBTX
- errBodyWriter
- toast_templ.go
- names
- rowFigure
- Before/after pairs
- SearchReadableBlocksRow
- open5e.go
- Queries
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- .ListBlockStepsSince
- Site
- TestRunCommandUsesCurrentSheet

## God Nodes (most connected - your core abstractions)
1. `App` - 320 edges
2. `newTestEnv()` - 130 edges
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
- `UI conventions` --references--> `removeButton()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go
- `UI conventions` --references--> `selectItems()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (141 total, 27 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (61): contextKey, classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild, o5eClass (+53 more)

### Community 1 - "dialog_templ.go"
Cohesion: 0.12
Nodes (35): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+27 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (86): keepAwake(), keepAwake(), ListImageCharactersRow, CountHabitatsRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, Condition (+78 more)

### Community 3 - ".do"
Cohesion: 0.13
Nodes (18): testEnv, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm() (+10 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Sheet"
Cohesion: 0.06
Nodes (20): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+12 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (98): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+90 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "FilesPage"
Cohesion: 0.07
Nodes (44): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent(), Props (+36 more)

### Community 10 - "Store"
Cohesion: 0.10
Nodes (37): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+29 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (68): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+60 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (27): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor() (+19 more)

### Community 14 - "time.Time"
Cohesion: 0.11
Nodes (13): Identity, FromContext(), browserName(), firstLine(), mdLinkText(), formIntQuery(), TestUserAgent(), userAgentSummary() (+5 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (43): GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass() (+35 more)

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
Cohesion: 0.16
Nodes (15): ArmorData, ftDistance, ftItem, o5eArmor, ftAmount(), ftArmor(), ftCastingTime(), ftDuration() (+7 more)

### Community 21 - "levelUpForm"
Cohesion: 0.07
Nodes (60): Props, Description(), Title(), Checkbox(), initialChecked(), Label(), initialValue(), Item() (+52 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "App"
Cohesion: 0.08
Nodes (10): imageItem(), joinPath(), SearchComboItems(), WithCampaign(), DisplayPage(), JoinQRSrc(), LinkItem, SwitcherResults() (+2 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (163): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+155 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.13
Nodes (29): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+21 more)

### Community 28 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (16): ClearCookie(), SetCookie(), writeCookie(), ListSubclassesOfRow, Note, subclassIndex(), subclassOptions(), homebrewTabURL() (+8 more)

### Community 29 - "newTestEnv"
Cohesion: 0.05
Nodes (55): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestEncounterTransitions() (+47 more)

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
Cohesion: 0.05
Nodes (37): GetImageRow, ListImagesRow, Attachment, CreateCombatantParams, NoteKind, saveScreenshot(), nullID(), ensureCategory() (+29 more)

### Community 34 - "CommandBuilder"
Cohesion: 0.06
Nodes (49): Field, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), armorForm(), builderForm() (+41 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "LinkTarget"
Cohesion: 0.15
Nodes (9): LinkTarget, saveFieldLinks(), AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink (+1 more)

### Community 38 - "Slugify"
Cohesion: 0.07
Nodes (46): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, TextData (+38 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Character"
Cohesion: 0.08
Nodes (14): Character, Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks() (+6 more)

### Community 41 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 42 - "Book"
Cohesion: 0.08
Nodes (40): bookWriter, mirror, rawBook, internal/compendium, parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+32 more)

### Community 43 - "buildReview"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "testing.T"
Cohesion: 0.03
Nodes (93): TestCoreTraits(), feature(), RemoveSourceDir(), TestRemoveSourceDir(), Parse(), tokenize(), fighter(), TestCommands() (+85 more)

### Community 46 - "Feature"
Cohesion: 0.14
Nodes (8): LevelUpChoice, TestSlotMaxima(), applyASI(), hpGain(), raise(), ClassLevel, CasterLevel(), SlotMaxima()

### Community 47 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (118): CountSharedBySourceRow, UI conventions, Addon(), Button(), Input(), InputGroup(), Scripts(), Separator() (+110 more)

### Community 48 - "CompendiumKind"
Cohesion: 0.04
Nodes (36): CountCompendiumRow, ListEntryCopiesRow, ListHomebrewByOwnerRow, ListHomebrewRow, UpsertCompendiumEntryParams, CompendiumEntry, CompendiumSource, CompendiumKind (+28 more)

### Community 49 - "State"
Cohesion: 0.11
Nodes (13): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+5 more)

### Community 50 - "cell"
Cohesion: 0.13
Nodes (18): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+10 more)

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

### Community 56 - "DecodeData"
Cohesion: 0.07
Nodes (40): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions() (+32 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.10
Nodes (21): generated, paths, rules, formatters, enable, exclusions, settings, match-constant (+13 more)

### Community 58 - "Open"
Cohesion: 0.16
Nodes (17): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+9 more)

### Community 59 - "ItemData"
Cohesion: 0.19
Nodes (13): ImageCategory, categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows() (+5 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "CN"
Cohesion: 0.06
Nodes (80): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+72 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".campaignSourcesSet"
Cohesion: 0.12
Nodes (12): ftsQuery(), Search(), SearchCompendiumRow, ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), CompendiumURL() (+4 more)

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

### Community 76 - "tabs_templ.go"
Cohesion: 0.11
Nodes (36): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+28 more)

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
Nodes (12): CompendiumBase(), wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki() (+4 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "Feature"
Cohesion: 0.11
Nodes (16): FeatureEffects, ItemData, armorCondition(), columnInts(), ParseEffects(), parseProfs(), SaveFeature(), UsesLabel() (+8 more)

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - ".compendiumWhole"
Cohesion: 0.12
Nodes (8): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag(), CompendiumRest()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".commandFailed"
Cohesion: 0.13
Nodes (8): Combatant, Entity, Character, entityTopics(), linkEntityLook(), formInt(), linkCharacterLook(), lookFree()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (44): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+36 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "select_templ.go"
Cohesion: 0.11
Nodes (22): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state() (+14 more)

### Community 103 - "newLiveFixture"
Cohesion: 0.11
Nodes (21): Player(), BenchmarkEncounterState(), BenchmarkSheetState(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture() (+13 more)

### Community 104 - ".found"
Cohesion: 0.24
Nodes (4): compendiumBack(), compendiumPath(), viewerCampaign(), HomebrewView

### Community 105 - "NewClient"
Cohesion: 0.33
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 106 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 107 - "ftBook"
Cohesion: 0.15
Nodes (18): ftBook, ftRec, each(), ftCopyTargets(), ftDecode(), ftMergeAbility(), ftMod(), ftName() (+10 more)

### Community 108 - "Campaign"
Cohesion: 0.21
Nodes (5): ListAttachmentFilesRow, Campaign, inCampaign(), visitorOnly(), CampaignSwitch

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 111 - "RosterCharacter"
Cohesion: 0.43
Nodes (7): CampaignRoster(), Roster, rosterCard(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 112 - "config.go"
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 114 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 115 - ".encounterView"
Cohesion: 0.10
Nodes (18): Defense, Encounter, ListCombatantsRow, Defenses, ParseDefenses(), EncounterStatus, combatantDefenses(), damageChoice() (+10 more)

### Community 116 - "writeBook"
Cohesion: 0.27
Nodes (9): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpBrokenSettings(), TestStartUpRedownloads(), writeBook(), appState (+1 more)

### Community 117 - ".loadLive"
Cohesion: 0.08
Nodes (50): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Content() (+42 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".sourceIssues"
Cohesion: 0.43
Nodes (6): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 126 - "names"
Cohesion: 0.40
Nodes (5): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), SourcePreviewView

### Community 127 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (18): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+10 more)

### Community 134 - "Queries"
Cohesion: 0.07
Nodes (34): featureSwitch, dnd5eFeatures(), Queries, Env, ClassFeature, ClassInfo, LevelUpPlan, featureNames() (+26 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Site"
Cohesion: 0.12
Nodes (17): HomebrewOwner, Site, Source, TestHomebrew(), fiveToolsPage(), SaveHomebrew(), SeedHabitats(), storeSource() (+9 more)

### Community 163 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 391 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **27 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.do`, `Queries`, `time.Time`, `perm.go`, `github.com/gin-gonic/gin.Context`, `database/sql.NullInt64`, `Slugify`, `Character`, `CompendiumKind`, `State`, `cert.go`, `DecodeData`, `FiveToolsJob`, `Topic`, `.campaignSourcesSet`, `ParseWiki`, `Feature`, `.compendiumWhole`, `.commandFailed`, `.found`, `Campaign`, `Hub`, `.encounterView`, `writeBook`, `.loadLive`, `.sourceIssues`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11320411320411321 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `TestRunCommandUsesCurrentSheet`, `encoding/json.RawMessage`, `SheetView`, `Slugify`, `Queries`, `Command`, `testing.T`, `Feature`, `CompendiumKind`, `.encounterView`, `Feature`, `levelUpForm`, `.sourceIssues`, `.commandFailed`, `DecodeData`, `ParseSheet`?**
  _High betweenness centrality (0.015) - this node is a cross-community bridge._
- **Should `dialog_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11607843137254902 - nodes in this community are weakly interconnected._
- **Why does `ListCombatantsRow` connect `.encounterView` to `compendium_templ.go`, `database/sql.NullInt64`, `.commandFailed`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.008782849239280775 - nodes in this community are weakly interconnected._