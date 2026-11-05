# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 284 files · ~530,619 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 3857 nodes · 16413 edges · 125 communities (107 shown, 18 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1626 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `a0d64524`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- FromContext
- context.Context
- .do
- createProvider
- CombatantView
- SheetView
- combobox.js
- redirect.go
- Parse
- Feature
- select.js
- Command
- .Entry
- NoteView
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- empty_templ.go
- levelUpForm
- wikilink.js
- github.com/a-h/templ.Component
- popover.js
- wikiText
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- testing.T
- app.js
- tabbable.js
- useDismiss
- CompendiumEntry
- PlanSourceSwitch
- command.js
- dialog.js
- database/sql.NullInt64
- golang.org/x/net/html.Node
- tooltip.js
- LevelUpView
- .levelUp
- capReader
- export/markdown.go
- composite.js
- Character
- Client
- ClassInfo
- tabs_templ.go
- Attachment
- New
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- select_templ.go
- golangci-lint Config
- Slugify
- Book
- radiogroup.js
- vendorjs Tool
- radiogroup_templ.go
- Pending
- sqlc Config
- Encounter
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
- checkbox.js
- cell
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- ParseWikidot
- Lint and Release Rules
- Run Script
- github.com/a-h/templ.Component
- useTypeahead
- NameDesc
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Hub
- Site
- ftItemData
- popover_templ.go
- .wizardCommand
- ftBook
- Import
- .features
- TestSourceSwitchKeepsEdits
- State
- .sourceIssues
- NoteListView
- Weapon
- build.sh
- Sheet
- errBodyWriter
- .sendFeedback
- Session
- open5e.go
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 283 edges
2. `CN()` - 112 edges
3. `Button()` - 110 edges
4. `Sheet` - 103 edges
5. `newTestEnv()` - 100 edges
6. `SheetView` - 56 edges
7. `IfElse()` - 51 edges
8. `Command` - 46 edges
9. `Input()` - 46 edges
10. `Slugify()` - 38 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `NewManager()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/auth/manager.go
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Open()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/db/db.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (125 total, 18 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (55): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+47 more)

### Community 1 - "FromContext"
Cohesion: 0.06
Nodes (21): Identity, FromContext(), browserName(), firstLine(), mdLinkText(), TestUserAgent(), userAgentSummary(), csvSource() (+13 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (71): keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Note, lanHostname(), shortHost() (+63 more)

### Community 3 - ".do"
Cohesion: 0.12
Nodes (20): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "CombatantView"
Cohesion: 0.07
Nodes (76): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+68 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (87): ArmorEffect(), Signed(), Field, Variant, Action(), Card(), Content(), Description() (+79 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Parse"
Cohesion: 0.08
Nodes (28): feature(), Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses() (+20 more)

### Community 10 - "Feature"
Cohesion: 0.18
Nodes (8): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesItems(), usesValue()

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (62): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+54 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput (+22 more)

### Community 14 - "NoteView"
Cohesion: 0.14
Nodes (22): ListEntityFieldBacklinksRow, ListBacklinksRow, initialChecked(), Switch(), EntityDetail(), entityForm(), EntityPage(), EntityView (+14 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (46): GetBlockRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+38 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "empty_templ.go"
Cohesion: 0.14
Nodes (32): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, ListMonstersByNameRow, Content(), Description() (+24 more)

### Community 21 - "levelUpForm"
Cohesion: 0.12
Nodes (30): Props, Footer(), Checkbox(), initialChecked(), abilitySelect(), levelUpForm(), slotSummary(), spellChecklist() (+22 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (108): Look, Description(), Button(), Input(), Addon(), Button(), Input(), InputGroup() (+100 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "wikiText"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.15
Nodes (26): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp() (+18 more)

### Community 28 - "App"
Cohesion: 0.06
Nodes (20): ClearCookie(), SetCookie(), writeCookie(), coinCommand(), commandText(), fiveToolsKey(), formIntQuery(), isHTMX() (+12 more)

### Community 29 - "testing.T"
Cohesion: 0.05
Nodes (79): readBook(), TestParseFiveToolsBook(), fiveToolsClasses(), TestFiveToolsClasses(), TestFiveToolsClassLevelUp(), serveSplit(), TestSplitTLS(), TestSplitTLSSilentConn() (+71 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "CompendiumEntry"
Cohesion: 0.05
Nodes (20): DecodeData(), ftsQuery(), Search(), ListEntryCopiesRow, Campaign, CompendiumEntry, CompendiumSource, spellFromRow() (+12 more)

### Community 34 - "PlanSourceSwitch"
Cohesion: 0.19
Nodes (9): featureSwitch, featureNames(), SourceRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch(), sameSlug() (+1 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (55): Defense, ListHomebrewByOwnerRow, ListHomebrewRow, Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Condition (+47 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "LevelUpView"
Cohesion: 0.27
Nodes (7): LevelUpView, Option, SpellNeed, Plural(), drawsGroup(), ChoiceGroup, CreationChoices

### Community 41 - ".levelUp"
Cohesion: 0.09
Nodes (29): TestSlotMaxima(), ClassFeature, ClassInfo, LevelUpPlan, ClassLevel, CasterLevel(), SlotMaxima(), names() (+21 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Character"
Cohesion: 0.08
Nodes (13): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), CanCharacter(), nextStep() (+5 more)

### Community 46 - "Client"
Cohesion: 0.22
Nodes (3): message, Client, topicIDs()

### Community 47 - "ClassInfo"
Cohesion: 0.23
Nodes (4): LevelUpChoice, applyASI(), hpGain(), raise()

### Community 48 - "tabs_templ.go"
Cohesion: 0.14
Nodes (36): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+28 more)

### Community 49 - "Attachment"
Cohesion: 0.24
Nodes (4): Attachment, saveScreenshot(), saveSniffed(), CreateAttachmentParams

### Community 50 - "New"
Cohesion: 0.44
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

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
Nodes (14): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+6 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "select_templ.go"
Cohesion: 0.07
Nodes (44): ImageCategory, Content(), initialValue(), Item(), itemLabel(), Select(), state(), Trigger() (+36 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Slugify"
Cohesion: 0.13
Nodes (21): Cantrips(), TextData, Slugify(), Spell, textMarkdown(), catchUpResistances(), TestCatchUpResistances(), addResistance() (+13 more)

### Community 59 - "Book"
Cohesion: 0.12
Nodes (26): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools(), fetchBooks() (+18 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "radiogroup_templ.go"
Cohesion: 0.17
Nodes (15): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+7 more)

### Community 63 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Encounter"
Cohesion: 0.08
Nodes (12): Combatant, Encounter, Entity, Character, entityTopics(), formInt(), CreateEncounterParams, ListCharacterEncountersParams (+4 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (162): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+154 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "builder_templ.go"
Cohesion: 0.08
Nodes (44): CloseFor(), Description(), Field(), Group(), Label(), Input(), Textarea(), armorForm() (+36 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "cell"
Cohesion: 0.12
Nodes (23): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+15 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.18
Nodes (13): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveSourcesAlert(), TestLiveVisibility(), LiveKey(), dial() (+5 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - ".get"
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "ParseWikidot"
Cohesion: 0.12
Nodes (24): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, dnd5eFeatures(), pageSource(), PageSourceKey() (+16 more)

### Community 86 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (55): Props, Variant, SpellData, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), alignmentColor() (+47 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "NameDesc"
Cohesion: 0.21
Nodes (16): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, cantripFits(), checkNewOptions() (+8 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.10
Nodes (30): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+22 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "Site"
Cohesion: 0.16
Nodes (10): Site, dataFor(), Kind, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SiteByKey() (+2 more)

### Community 104 - "ftItemData"
Cohesion: 0.15
Nodes (16): ArmorData, ftDistance, ftItem, o5eArmor, ItemData, ftAmount(), ftArmor(), ftDuration() (+8 more)

### Community 105 - "popover_templ.go"
Cohesion: 0.15
Nodes (22): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+14 more)

### Community 106 - ".wizardCommand"
Cohesion: 0.09
Nodes (30): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+22 more)

### Community 107 - "ftBook"
Cohesion: 0.17
Nodes (17): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+9 more)

### Community 108 - "Import"
Cohesion: 0.13
Nodes (23): migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open() (+15 more)

### Community 109 - ".features"
Cohesion: 0.29
Nodes (5): ftClassFeature, ftClassRefs, ftCode(), featuresSearchText(), ftFeatureKey()

### Community 110 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.28
Nodes (6): UpsertCompendiumEntryParams, classCopy(), testEnv, TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 111 - "State"
Cohesion: 0.10
Nodes (13): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+5 more)

### Community 112 - ".sourceIssues"
Cohesion: 0.53
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 113 - "NoteListView"
Cohesion: 0.50
Nodes (3): SearchReadableBlocksRow, SearchReadableBlocksParams, NoteListView

### Community 115 - "Weapon"
Cohesion: 0.12
Nodes (11): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+3 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 123 - "Sheet"
Cohesion: 0.08
Nodes (15): DeathSaves, TestCustomMods(), TestResistances(), Sum(), TestGearCommands(), Coins, Item, Sheet (+7 more)

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 131 - "Session"
Cohesion: 0.08
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), TestCredentialsMatch(), TestIsLoopback(), TestRequireAdmin() (+10 more)

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (23): Entry, fixture, fixtures, normalizer, Source, TestUnflattenTable(), parseFiveToolsBook(), parseFiveToolsBookReport() (+15 more)

### Community 134 - "Import"
Cohesion: 0.10
Nodes (34): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+26 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 345 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **18 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `FromContext`, `context.Context`, `.do`, `Session`, `perm.go`, `CompendiumEntry`, `database/sql.NullInt64`, `LevelUpView`, `.levelUp`, `Character`, `Attachment`, `cert.go`, `Slugify`, `Pending`, `Encounter`, `Hub`, `.wizardCommand`, `State`, `.sourceIssues`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10573264146281196 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `SheetView`, `Parse`, `Feature`, `Command`, `ParseSheet`, `CompendiumEntry`, `PlanSourceSwitch`, `database/sql.NullInt64`, `LevelUpView`, `.levelUp`, `ClassInfo`, `Slugify`, `Encounter`, `NameDesc`, `.wizardCommand`, `TestSourceSwitchKeepsEdits`, `.sourceIssues`, `Weapon`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Should `FromContext` be split into smaller, more focused modules?**
  _Cohesion score 0.06342494714587738 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Hub` to `compendium_templ.go`, `FromContext`, `Client`, `State`, `New`, `App`, `Pending`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011136434269139128 - nodes in this community are weakly interconnected._