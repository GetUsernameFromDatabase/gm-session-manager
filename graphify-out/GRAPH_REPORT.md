# Graph Report - gm-session-manager  (2026-11-06)

## Corpus Check
- 286 files · ~541,680 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 3919 nodes · 16796 edges · 127 communities (107 shown, 20 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1689 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `db83a48f`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- combobox_templ.go
- context.Context
- .do
- createProvider
- dialog_templ.go
- SheetView
- combobox.js
- redirect.go
- parseFiveToolsBook
- Feature
- select.js
- Command
- .Entry
- Input
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ImagePage
- wizardPick
- wikilink.js
- IfElse
- popover.js
- time.Time
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- testing.T
- app.js
- tabbable.js
- useDismiss
- .levelUp
- Parse
- command.js
- dialog.js
- FromContext
- golang.org/x/net/html.Node
- tooltip.js
- CombatantView
- .wizardCommand
- capReader
- export/markdown.go
- composite.js
- inputgroup_templ.go
- Slugify
- ParseEffects
- State
- Queries
- ParseWiki
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- .renderCompendium
- golangci-lint Config
- Sheet
- Book
- radiogroup.js
- vendorjs/main.go
- cell
- BuilderPicks
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- CN
- checkbox.js
- radiogroup_templ.go
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
- spellCounts
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- .data
- mounted
- NewClient
- Site
- ftItemData
- Spell
- Hub
- ftBook
- Store
- LevelUpView
- Session
- Client
- .sendFeedback
- TestSourceSwitchKeepsEdits
- .combatantHP
- Weapon
- NameDesc
- .loadBlock
- build.sh
- ftItemData
- RosterCharacter
- errBodyWriter
- open5e.go
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 287 edges
2. `CN()` - 113 edges
3. `Button()` - 112 edges
4. `Sheet` - 109 edges
5. `newTestEnv()` - 109 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 52 edges
8. `Command` - 47 edges
9. `Input()` - 46 edges
10. `Slugify()` - 41 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Fingerprint()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/join/cert.go
- `run()` --calls--> `LoadOrCreateCert()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/join/cert.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (127 total, 20 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.12
Nodes (38): contextKey, classRef, ftBookInfo, indexed, sourceOf, ASIError, clientMessage, TestMain() (+30 more)

### Community 1 - "combobox_templ.go"
Cohesion: 0.08
Nodes (45): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+37 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (81): keepAwake(), keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Encounter, SearchReadableBlocksRow (+73 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "dialog_templ.go"
Cohesion: 0.11
Nodes (45): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+37 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (104): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+96 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "parseFiveToolsBook"
Cohesion: 0.13
Nodes (18): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestParseFiveToolsBook(), DownloadFiveTools(), RemoveFiveTools() (+10 more)

### Community 10 - "Feature"
Cohesion: 0.15
Nodes (12): FeatureEffects, armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), UsesLabel(), CustomMod, Feature (+4 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.05
Nodes (76): Amount, CondOpKind, Effects, featureSwitch, HelpEntry, parser, Completions(), condCommand() (+68 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput (+22 more)

### Community 14 - "Input"
Cohesion: 0.04
Nodes (101): ImageCategory, Footer(), CloseFor(), Field(), Group(), Label(), Input(), Label() (+93 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (53): Note, ListCombatantsRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanCharacter() (+45 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ImagePage"
Cohesion: 0.09
Nodes (34): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, initials(), currentLook(), EntityDetail(), EntityPage() (+26 more)

### Community 21 - "wizardPick"
Cohesion: 0.16
Nodes (24): spellChecklist(), taughtPicked(), taughtRules(), TestTaughtPicked(), abilityForm(), backgroundAbilities(), choiceFieldset(), choiceName() (+16 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "IfElse"
Cohesion: 0.09
Nodes (46): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+38 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "time.Time"
Cohesion: 0.06
Nodes (21): PendingState, ListEntryCopiesRow, CompendiumSource, Pending, PendingKind, buildReview(), csvSource(), formValues() (+13 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.19
Nodes (22): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp() (+14 more)

### Community 28 - "App"
Cohesion: 0.06
Nodes (18): Campaign, entityTopics(), formInt(), formIntQuery(), isHTMX(), redirect(), imageItem(), imageVersion() (+10 more)

### Community 29 - "testing.T"
Cohesion: 0.05
Nodes (80): TestMasteryLimit(), TestDropFeat(), TestCanCharacter(), nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm() (+72 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - ".levelUp"
Cohesion: 0.13
Nodes (22): firstTable(), LevelSpells(), Options(), TestOptions(), errorField(), spellbook(), spellNeed(), classPin() (+14 more)

### Community 34 - "Parse"
Cohesion: 0.07
Nodes (33): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestDefenses(), TestHealingFromZeroResetsDeathSaves(), TestLanguages() (+25 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "FromContext"
Cohesion: 0.10
Nodes (14): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), mdLinkText() (+6 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CombatantView"
Cohesion: 0.09
Nodes (37): Defenses, damageChoice(), capitalized(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState() (+29 more)

### Community 41 - ".wizardCommand"
Cohesion: 0.11
Nodes (26): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestCoreTraits() (+18 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "inputgroup_templ.go"
Cohesion: 0.12
Nodes (24): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+16 more)

### Community 46 - "Slugify"
Cohesion: 0.14
Nodes (18): TextData, Slugify(), textMarkdown(), upgradeSpecies(), pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue (+10 more)

### Community 47 - "ParseEffects"
Cohesion: 0.14
Nodes (9): LevelUpChoice, TestSlotMaxima(), applyASI(), LevelUpPlan, hpGain(), raise(), ClassLevel, CasterLevel() (+1 more)

### Community 48 - "State"
Cohesion: 0.10
Nodes (13): event, fakeHandler, Handler, message, patchMessage, stateHandler, TopicKind, viewPatch (+5 more)

### Community 49 - "Queries"
Cohesion: 0.05
Nodes (36): DecodeData(), Queries, Character, CompendiumEntry, Entity, ClassInfo, PlanSourceSwitch(), Campaign() (+28 more)

### Community 50 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

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

### Community 56 - ".renderCompendium"
Cohesion: 0.10
Nodes (9): ftsQuery(), Search(), SearchCompendiumRow, compendiumBack(), ImportSummary(), viewerCampaign(), fiveToolsKey(), SearchComboItems() (+1 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Sheet"
Cohesion: 0.08
Nodes (15): DeathSaves, TestCustomMods(), TestResistances(), TestExplainSums(), Sum(), TestGearCommands(), Coins, Item (+7 more)

### Community 59 - "Book"
Cohesion: 0.08
Nodes (36): bookWriter, ftClass, ftClassFeature, ftCopyRef, ftEntry, ftFeat, ftModOp, ftNamed (+28 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 63 - "BuilderPicks"
Cohesion: 0.38
Nodes (7): BuilderPickItems(), BuilderPicks(), builderRemove(), closeAfter(), BuilderPick, pickURL(), TestPickRowActs()

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (45): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, CreateCombatantParams (+37 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (130): Props, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size (+122 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "CN"
Cohesion: 0.08
Nodes (56): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, MediaProps, MediaVariant (+48 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "radiogroup_templ.go"
Cohesion: 0.19
Nodes (14): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+6 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.18
Nodes (12): newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveSourcesAlert(), TestLiveVisibility(), LiveKey(), dial(), patchOf() (+4 more)

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
Cohesion: 0.13
Nodes (21): block, ClassFeature, ClassData, dnd5eFeatures(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+13 more)

### Community 86 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (77): Props, Variant, SpellData, CountHabitatsRow, Badge(), badgeClasses(), Toast(), Toaster() (+69 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 100 - ".data"
Cohesion: 0.14
Nodes (22): ftMonster, ftSpellcasting, ftCell(), ftPrerequisite(), ftTable(), TestFiveToolsCell(), TestStripTagsTokens(), ftClassCell() (+14 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.22
Nodes (14): New(), NewClient(), randomID(), drain(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Site"
Cohesion: 0.11
Nodes (19): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+11 more)

### Community 104 - "ftItemData"
Cohesion: 0.39
Nodes (4): ftClassRefs, ftCode(), featuresSearchText(), ftFeatureKey()

### Community 105 - "Spell"
Cohesion: 0.19
Nodes (13): Cantrips(), Spell, catchUpResistances(), TestCatchUpResistances(), addResistance(), applyTraitText(), choiceOptions(), freeCast() (+5 more)

### Community 106 - "Hub"
Cohesion: 0.23
Nodes (3): ClientInfo, Hub, TestCloseSessionAndCharacter()

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (15): ftBook, ftRec, each(), ftDecode(), ftMod(), ftName(), ftParts(), ftRecords() (+7 more)

### Community 108 - "Store"
Cohesion: 0.10
Nodes (32): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply() (+24 more)

### Community 109 - "LevelUpView"
Cohesion: 0.26
Nodes (8): LevelUpView, Option, SpellNeed, Plural(), drawsGroup(), ChoiceGroup, CreationChoices, withSpells()

### Community 110 - "Session"
Cohesion: 0.09
Nodes (17): tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), TestCredentialsMatch(), TestIsLoopback(), TestRequireAdmin(), TestTokens() (+9 more)

### Community 112 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 113 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.28
Nodes (6): UpsertCompendiumEntryParams, classCopy(), testEnv, TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 115 - "Weapon"
Cohesion: 0.15
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 116 - "NameDesc"
Cohesion: 0.53
Nodes (4): OptionPrerequisite(), NameDesc, isRepeatable(), optionKind

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "ftItemData"
Cohesion: 0.16
Nodes (15): ArmorData, ftDistance, ftItem, ItemData, ftAmount(), ftArmor(), ftCastingTime(), ftDuration() (+7 more)

### Community 123 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 133 - "open5e.go"
Cohesion: 0.08
Nodes (36): Entry, fixture, fixtures, looseString, normalizer, o5eArmor, o5eAssignment, o5eChild (+28 more)

### Community 134 - "Import"
Cohesion: 0.12
Nodes (29): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+21 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 347 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **20 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `perm.go`, `time.Time`, `.levelUp`, `FromContext`, `.wizardCommand`, `Slugify`, `State`, `Queries`, `.renderCompendium`, `database/sql.NullInt64`, `Spell`, `Hub`, `Store`, `LevelUpView`, `Session`, `.combatantHP`, `.loadBlock`?**
  _High betweenness centrality (0.062) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.12495389416938572 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `SheetView`, `Feature`, `Command`, `perm.go`, `ParseSheet`, `.levelUp`, `Parse`, `CombatantView`, `.wizardCommand`, `Slugify`, `ParseEffects`, `Queries`, `Spell`, `LevelUpView`, `TestSourceSwitchKeepsEdits`, `.combatantHP`, `Weapon`, `NameDesc`?**
  _High betweenness centrality (0.023) - this node is a cross-community bridge._
- **Should `combobox_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.08244897959183674 - nodes in this community are weakly interconnected._
- **Why does `Campaign()` connect `Queries` to `compendium_templ.go`, `context.Context`, `export/markdown.go`, `Store`, `App`?**
  _High betweenness centrality (0.014) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009893036824081512 - nodes in this community are weakly interconnected._