# Graph Report - gm-session-manager  (2026-11-10)

## Corpus Check
- 296 files · ~567,934 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4093 nodes · 17729 edges · 141 communities (115 shown, 26 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1791 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e7834757`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- context.Context
- .do
- createProvider
- who
- SheetView
- combobox.js
- redirect.go
- inputgroup_templ.go
- spellChecklist
- select.js
- Command
- .Entry
- github.com/a-h/templ.Component
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Input
- WizardView
- wikilink.js
- tabs_templ.go
- popover.js
- popover_templ.go
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- DecodeData
- Parse
- command.js
- dialog.js
- WizardChoice
- empty_templ.go
- tooltip.js
- .encounterView
- ParseAidedd
- .renderCompendium
- cell
- composite.js
- .sendFeedback
- Feature
- ClassInfo
- deviceLabel
- Character
- ParseWiki
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- tooltip_templ.go
- Book
- radiogroup.js
- vendorjs/main.go
- ImportView
- FiveToolsJob
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- Slugify
- checkbox.js
- Session
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- fromLookup
- useFocus
- mark_others.js
- choiceOptions
- Lint and Release Rules
- Run Script
- .levelUp
- useTypeahead
- report
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- NewClient
- Site
- .Router
- entityForm
- .whoami
- ftBook
- radiogroup_templ.go
- SearchReadableBlocksRow
- formErrors
- export/markdown.go
- Topic
- TestSourceSwitchKeepsEdits
- Pending
- Sheet
- toast_templ.go
- sourcesEnv
- build.sh
- Hub
- parseCommand
- errBodyWriter
- Queries
- capReader
- .ListCompendiumByKind
- Before/after pairs
- open5e.go
- Import
- db/CLAUDE.md
- perm/CLAUDE.md
- ftNamed
- .sourceIssues
- Badge
- .ListBlockStepsSince
- .ListCompendiumNames

## God Nodes (most connected - your core abstractions)
1. `App` - 303 edges
2. `newTestEnv()` - 121 edges
3. `Sheet` - 118 edges
4. `Button()` - 117 edges
5. `CN()` - 113 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 56 edges
8. `Command` - 49 edges
9. `Input()` - 47 edges
10. `Queries` - 43 edges

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

## Communities (141 total, 26 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (62): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+54 more)

### Community 1 - "CombatantView"
Cohesion: 0.08
Nodes (46): Defenses, capitalized(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+38 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (81): keepAwake(), ftsQuery(), Search(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource (+73 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (22): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+14 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "who"
Cohesion: 0.12
Nodes (6): GetBlockRow, compendiumBack(), HomebrewView, blockAccessLevel, stepsPayload, who

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (104): ArmorEffect(), Signed(), Field, Button(), Variant, Action(), Card(), Content() (+96 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "inputgroup_templ.go"
Cohesion: 0.12
Nodes (24): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+16 more)

### Community 10 - "spellChecklist"
Cohesion: 0.13
Nodes (16): Option, SpellNeed, SubclassPanel, Plural(), sortedColumns(), spellChecklist(), SpellColumn(), spellCounts() (+8 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.08
Nodes (61): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+53 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (31): HomebrewOwner, MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, ValidKind() (+23 more)

### Community 14 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (108): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+100 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.08
Nodes (47): Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass(), CanCharacter() (+39 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Input"
Cohesion: 0.06
Nodes (67): Combobox(), Content(), Input(), Item(), List(), state(), CloseFor(), Description() (+59 more)

### Community 21 - "WizardView"
Cohesion: 0.31
Nodes (15): TestOptionText(), abilityForm(), WizardStep, WizardView, optionText(), skillName(), stepLinks(), wizardAbilities() (+7 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.12
Nodes (40): Look, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+32 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "popover_templ.go"
Cohesion: 0.15
Nodes (19): Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition(), conditionLabel() (+11 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.11
Nodes (36): Import(), TestUpgradeSheet(), ParseSheet(), TestBuilderCompendium(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI() (+28 more)

### Community 28 - "App"
Cohesion: 0.06
Nodes (19): ClearCookie(), SetCookie(), writeCookie(), entityTopics(), formIntQuery(), isHTMX(), redirect(), setAudience() (+11 more)

### Community 29 - "newTestEnv"
Cohesion: 0.04
Nodes (86): RemoveSourceDir(), TestRemoveSourceDir(), TestMasteryLimit(), TestDropFeat(), TestSheetPinsSource(), TestCanCharacter(), nullID(), TestBuilderFeat() (+78 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "DecodeData"
Cohesion: 0.06
Nodes (23): DecodeData(), ListEntryCopiesRow, ListSubclassesOfRow, Campaign, CompendiumEntry, spellFromRow(), subclassPanels(), viewerCampaign() (+15 more)

### Community 34 - "Parse"
Cohesion: 0.06
Nodes (41): Defense, feature(), TestFiveToolsClasses(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions() (+33 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "WizardChoice"
Cohesion: 0.29
Nodes (7): choiceName(), Choose(), SkillPick, WizardChoice, wizardFacts(), checkSkillPick(), FeaturePreview

### Community 38 - "empty_templ.go"
Cohesion: 0.10
Nodes (43): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListMonstersByNameRow (+35 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".encounterView"
Cohesion: 0.09
Nodes (15): Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), damageChoice(), initiativeMod(), outOfFight() (+7 more)

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - ".renderCompendium"
Cohesion: 0.18
Nodes (7): fiveToolsKey(), sourceInUseMsg(), sourceUsers(), FiveToolsView, NewFiveToolsView(), FiveToolsBook, FiveToolsGroup

### Community 43 - "cell"
Cohesion: 0.11
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 46 - "Feature"
Cohesion: 0.20
Nodes (7): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 47 - "ClassInfo"
Cohesion: 0.10
Nodes (14): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+6 more)

### Community 48 - "deviceLabel"
Cohesion: 0.50
Nodes (3): deviceLabel(), rosterHP(), Device

### Community 49 - "Character"
Cohesion: 0.21
Nodes (8): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), Roster, RosterCharacter

### Community 50 - "ParseWiki"
Cohesion: 0.21
Nodes (12): saveFieldLinks(), wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki() (+4 more)

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

### Community 56 - "ClassData"
Cohesion: 0.12
Nodes (26): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+18 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 59 - "Book"
Cohesion: 0.12
Nodes (25): bookWriter, mirror, rawBook, internal/compendium, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools() (+17 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "ImportView"
Cohesion: 0.11
Nodes (24): ImageCategory, categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows() (+16 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.31
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (45): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, CreateCombatantParams (+37 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (159): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+151 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Slugify"
Cohesion: 0.08
Nodes (41): EquipmentItem, Cantrips(), EquipmentOptions(), firstTable(), EquipmentOption, LevelSpells(), Options(), SkillChoice() (+33 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "Session"
Cohesion: 0.18
Nodes (9): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+1 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.11
Nodes (20): Player(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips() (+12 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "fromLookup"
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "choiceOptions"
Cohesion: 0.22
Nodes (10): Item(), LevelUpView, backgroundAbilities(), choiceFieldset(), choiceOptions(), drawsGroup(), ChoiceGroup, FieldErrors (+2 more)

### Community 86 - ".levelUp"
Cohesion: 0.11
Nodes (24): OptionPrerequisite(), Spell, addClassSpells(), spellNeed(), classPin(), TestSpellNeed(), applyPicks(), cantripFits() (+16 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "report"
Cohesion: 0.18
Nodes (11): folderName(), mdCell(), oneLine(), parseConsole(), slugify(), TestFolderName(), TestMarkdownCell(), consoleEntry (+3 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (41): ftClass, ftCopyRef, ftEntry, ftFeat, ftItem, ftModOp, ftMonster, ftNamed (+33 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Site"
Cohesion: 0.13
Nodes (14): Site, dataFor(), Kind, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey() (+6 more)

### Community 104 - ".Router"
Cohesion: 0.16
Nodes (7): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 105 - "entityForm"
Cohesion: 0.31
Nodes (9): ListBacklinksRow, initialChecked(), Switch(), EntityDetail(), entityForm(), EntityView, OrderedFieldKeys(), Props (+1 more)

### Community 106 - ".whoami"
Cohesion: 0.07
Nodes (20): Identity, FromContext(), Entity, linkEntityLook(), browserName(), firstLine(), mdLinkText(), TestUserAgent() (+12 more)

### Community 107 - "ftBook"
Cohesion: 0.15
Nodes (22): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftDecode(), ftDuration() (+14 more)

### Community 108 - "radiogroup_templ.go"
Cohesion: 0.36
Nodes (7): initialValue(), RadioGroup(), state(), ctxKey, ctxState, ItemProps, Props

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 111 - "export/markdown.go"
Cohesion: 0.31
Nodes (9): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+1 more)

### Community 112 - "Topic"
Cohesion: 0.10
Nodes (14): event, fakeHandler, Handler, message, patchMessage, stateHandler, TopicKind, viewPatch (+6 more)

### Community 113 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 114 - "Pending"
Cohesion: 0.20
Nodes (5): PendingState, randomID(), Pending, PendingKind, Waiting

### Community 115 - "Sheet"
Cohesion: 0.05
Nodes (27): DeathSaves, Part, Skill, Completions(), TestCustomMods(), TestModifier(), TestWeaponKind(), skillByKey() (+19 more)

### Community 116 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 117 - "sourcesEnv"
Cohesion: 0.33
Nodes (5): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "Hub"
Cohesion: 0.16
Nodes (3): ClientInfo, Client, Hub

### Community 123 - "parseCommand"
Cohesion: 0.40
Nodes (5): featureCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Queries"
Cohesion: 0.09
Nodes (30): featureSwitch, Queries, Env, ClassInfo, ClassLevel, featureNames(), SourceRef, SubclassRef (+22 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.10
Nodes (30): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eMastery(), TestUnflattenTable() (+22 more)

### Community 134 - "Import"
Cohesion: 0.06
Nodes (60): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), copySource() (+52 more)

### Community 137 - "ftNamed"
Cohesion: 0.11
Nodes (16): ArmorData, ftClassFeature, ftClassRefs, o5eArmor, ItemData, ftArmor(), ftCode(), ftItemData() (+8 more)

### Community 141 - ".sourceIssues"
Cohesion: 0.36
Nodes (7): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, featKey(), pinFeat(), sheetRef

### Community 148 - "Badge"
Cohesion: 0.05
Nodes (65): Props, Variant, SpellData, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), alignmentColor() (+57 more)

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 383 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **26 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `who`, `Import`, `.sourceIssues`, `perm.go`, `DecodeData`, `.encounterView`, `.renderCompendium`, `Character`, `FiveToolsJob`, `database/sql.NullInt64`, `Slugify`, `.levelUp`, `.Router`, `.whoami`, `Topic`, `Pending`, `Hub`, `Queries`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10769679692936185 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `DecodeData`, `Parse`, `encoding/json.RawMessage`, `SheetView`, `.encounterView`, `Slugify`, `Command`, `.sourceIssues`, `Feature`, `ClassInfo`, `TestSourceSwitchKeepsEdits`, `choiceOptions`, `.levelUp`, `ParseSheet`, `Queries`?**
  _High betweenness centrality (0.019) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.08305084745762711 - nodes in this community are weakly interconnected._
- **Why does `Slugify()` connect `Slugify` to `compendium_templ.go`, `DecodeData`, `encoding/json.RawMessage`, `open5e.go`, `Import`, `Site`, `ParseAidedd`, `ftBook`, `.Entry`, `.sourceIssues`, `Character`, `Badge`, `.levelUp`, `ClassData`, `App`, `Queries`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01015531660692951 - nodes in this community are weakly interconnected._