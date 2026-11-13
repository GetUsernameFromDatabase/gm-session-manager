# Graph Report - gm-session-manager  (2026-11-13)

## Corpus Check
- 335 files · ~596,798 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4342 nodes · 20853 edges · 154 communities (124 shown, 30 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2076 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `21da0244`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Import
- context.Context
- .do
- createProvider
- select_templ.go
- SheetView
- combobox.js
- redirect.go
- .levelUp
- Store
- select.js
- command.go
- .Entry
- capReader
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
- NameDesc
- testing.T
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- ParseWikidot
- command.js
- dialog.js
- Parse
- formErrors
- tooltip.js
- .Router
- Attachment
- Book
- export/markdown.go
- composite.js
- formInt
- ParseEffects
- CompendiumPath
- .fail
- SaveHomebrew
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Sheet
- golangci-lint Config
- Slugify
- Layout
- radiogroup.js
- vendorjs/main.go
- github.com/a-h/templ.Component
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- .compendiumDir
- use_transition_status.js
- playwright
- release.sh
- dialog_templ.go
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
- DownloadFiveTools
- Lint and Release Rules
- Run Script
- .gearFor
- useTypeahead
- CompendiumKind
- Validator Rule
- Go Module
- ImagePage
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- CN
- newLiveFixture
- profFocus
- State
- .wizardView
- ftBook
- deviceLabel
- ftDuration
- viewerCtx
- config.go
- Client
- time.Time
- NewClient
- build.sh
- ListEntryCopiesRow
- looseString
- errBodyWriter
- ClassInfo
- names
- net/url.Values
- Before/after pairs
- Command
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- Pending
- Client
- isToggle
- sourceUsedError
- scoped
- parseHP
- LinkTarget
- sourcesEnv
- Campaign
- findItem
- toast_templ.go
- TestRunCommandUsesCurrentSheet
- TestLiveUpdatesAndKick
- WeaponKind
- spellCounts

## God Nodes (most connected - your core abstractions)
1. `App` - 349 edges
2. `newTestEnv()` - 131 edges
3. `Button()` - 122 edges
4. `Sheet` - 121 edges
5. `CN()` - 113 edges
6. `SheetView` - 76 edges
7. `IfElse()` - 62 edges
8. `Command` - 59 edges
9. `CompendiumKind` - 54 edges
10. `ParseSheet()` - 47 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (154 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "Import"
Cohesion: 0.09
Nodes (40): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+32 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (80): keepAwake(), keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, ListAttachmentFilesRow (+72 more)

### Community 3 - ".do"
Cohesion: 0.15
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "select_templ.go"
Cohesion: 0.04
Nodes (66): ImageCategory, Defenses, Content(), Group(), initialValue(), Item(), itemLabel(), Label() (+58 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (130): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+122 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - ".levelUp"
Cohesion: 0.10
Nodes (22): ListSubclassesOfRow, ClassFeature, LevelUpPlan, Spellbook(), CountWord(), spellNeed(), subclassIndex(), subclassOptions() (+14 more)

### Community 10 - "Store"
Cohesion: 0.08
Nodes (30): fixture, tokens, migration, isCleared(), newFixture(), serve(), TestMiddlewareIdentity(), TestOnlyHashIsStored() (+22 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.12
Nodes (24): CondOpKind, HelpEntry, parser, condCommand(), conditionSlug(), findMod(), CondOp, newMod() (+16 more)

### Community 13 - ".Entry"
Cohesion: 0.06
Nodes (37): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), ValidKind(), FiveToolsImages() (+29 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (50): GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass() (+42 more)

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
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 21 - "levelUpForm"
Cohesion: 0.08
Nodes (52): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Action(), Alert() (+44 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (30): Campaign, compendiumBack(), compendiumPath(), homebrewTabURL(), pathCampaign(), viewerCampaign(), inCampaign(), formIntQuery() (+22 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (118): Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props (+110 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.12
Nodes (33): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage() (+25 more)

### Community 28 - "NameDesc"
Cohesion: 0.18
Nodes (18): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+10 more)

### Community 29 - "testing.T"
Cohesion: 0.03
Nodes (130): Defense, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestDecode(), RemoveSourceDir(), TestRemoveSourceDir(), TestDefenses() (+122 more)

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
Nodes (53): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Combatant, Condition, Encounter, CreateCombatantParams, ListCombatantsRow (+45 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.12
Nodes (25): block, ClassFeature, ClassData, pageSource(), PageSourceKey(), cellAt(), classTable(), featuresUnder() (+17 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Parse"
Cohesion: 0.13
Nodes (17): Heal(), Parse(), tokenize(), fighter(), TestCommands(), TestHealingFromZeroResetsDeathSaves(), TestMasteryCommand(), TestParseRawArgs() (+9 more)

### Community 38 - "formErrors"
Cohesion: 0.29
Nodes (5): Handlers, formErrors(), fieldError, pickError, userError

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".Router"
Cohesion: 0.19
Nodes (6): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag(), CompendiumRest()

### Community 41 - "Attachment"
Cohesion: 0.10
Nodes (9): Attachment, saveScreenshot(), csvSource(), parseCSV(), TestGuessMapping(), saveSniffed(), CreateAttachmentParams, csvImport (+1 more)

### Community 42 - "Book"
Cohesion: 0.09
Nodes (30): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+22 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "formInt"
Cohesion: 0.16
Nodes (9): Entity, entityTopics(), linkEntityLook(), formInt(), linkCharacterLook(), lookFree(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams (+1 more)

### Community 46 - "ParseEffects"
Cohesion: 0.07
Nodes (18): FeatureEffects, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), UsesLabel() (+10 more)

### Community 47 - "CompendiumPath"
Cohesion: 0.05
Nodes (64): CountHabitatsRow, CountSharedBySourceRow, SearchCompendiumRow, classBadges(), classBody(), CompendiumPage(), CompendiumResults(), compendiumRow() (+56 more)

### Community 48 - ".fail"
Cohesion: 0.09
Nodes (9): Character, Character, CanCharacter(), nextStep(), backgroundRaises(), setupOnce(), wizardPath(), Roster (+1 more)

### Community 49 - "SaveHomebrew"
Cohesion: 0.18
Nodes (8): Site, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SavePage(), SiteByKey(), TestParsePageURL()

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

### Community 56 - "Sheet"
Cohesion: 0.05
Nodes (27): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+19 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.08
Nodes (26): rules, generated, paths, rules, formatters, enable, exclusions, settings (+18 more)

### Community 58 - "Slugify"
Cohesion: 0.12
Nodes (24): Cantrips(), Slugify(), AddResistance(), FreeCast(), SwapFeatSpell(), Spell, spellFromRow(), addClassSpells() (+16 more)

### Community 59 - "Layout"
Cohesion: 0.07
Nodes (59): UI conventions, Scripts(), Separator(), Spinner(), campaignLink(), PendingRequests(), PendingGroup, CharacterPage() (+51 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "github.com/a-h/templ.Component"
Cohesion: 0.08
Nodes (69): ctxState, Coins, Chip(), Chips(), ChipsInput(), Combobox(), Content(), Input() (+61 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (35): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+27 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".compendiumDir"
Cohesion: 0.24
Nodes (4): ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "dialog_templ.go"
Cohesion: 0.13
Nodes (41): ctxKey, ctxState, Props, Dialog(), Close(), CloseFor(), Content(), Description() (+33 more)

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
Nodes (39): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+31 more)

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
Cohesion: 0.33
Nodes (8): CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks(), WikiLinks(), WikiLink, WikiPart

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "DownloadFiveTools"
Cohesion: 0.14
Nodes (21): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook(), DownloadFiveTools() (+13 more)

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - ".gearFor"
Cohesion: 0.20
Nodes (9): EquipmentItem, EquipmentOption, ItemData, WornArmor, compendiumWeapon(), TestCompendiumWeapon(), applyEquipment(), EquipmentChoice (+1 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "CompendiumKind"
Cohesion: 0.05
Nodes (27): Decode(), TextData, UpsertCompendiumEntryParams, AddFeat(), CompendiumKind, textMarkdown(), applyBackground(), CopyKey() (+19 more)

### Community 96 - "ImagePage"
Cohesion: 0.08
Nodes (35): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent(), Props (+27 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (45): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+37 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "CN"
Cohesion: 0.06
Nodes (66): Props, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps (+58 more)

### Community 103 - "newLiveFixture"
Cohesion: 0.10
Nodes (25): Player(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas, internal/web, Notes, Websocket and live views (+17 more)

### Community 105 - "State"
Cohesion: 0.11
Nodes (12): event, fakeHandler, Handler, patchMessage, stateHandler, viewPatch, Topic, TopicKind (+4 more)

### Community 106 - ".wizardView"
Cohesion: 0.10
Nodes (26): internal/compendium, CoreTraits(), EquipmentOptions(), SkillChoice(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), dataFor() (+18 more)

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (15): ftBook, ftRec, each(), ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftRecords() (+7 more)

### Community 108 - "deviceLabel"
Cohesion: 0.50
Nodes (3): deviceLabel(), rosterHP(), Device

### Community 109 - "ftDuration"
Cohesion: 0.50
Nodes (5): ftDistance, ftAmount(), ftCastingTime(), ftDuration(), ftRange()

### Community 111 - "viewerCtx"
Cohesion: 0.22
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 116 - "time.Time"
Cohesion: 0.06
Nodes (25): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), ListImagesRow, browserName(), firstLine() (+17 more)

### Community 117 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "ListEntryCopiesRow"
Cohesion: 0.33
Nodes (6): ListEntryCopiesRow, ListEntryCopiesParams, CompareCopy, CompareView, EntryView, HabitatCopy

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ClassInfo"
Cohesion: 0.10
Nodes (29): featureSwitch, Queries, CompendiumEntry, Env, ClassInfo, ClassLevel, featureNames(), SourceRef (+21 more)

### Community 126 - "names"
Cohesion: 0.33
Nodes (7): ftRaceAbility(), SwitchFeature, SwitchPlan, PlanSourceSwitch(), names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 127 - "net/url.Values"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "Command"
Cohesion: 0.22
Nodes (13): Effects, Command, NewCommand(), parseSet(), parseWeapon(), SaveWeapon(), SaveFeature(), SizeNamed() (+5 more)

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (23): Entry, fixture, fixtures, normalizer, o5eProperty, Source, sourceFormat, TestOpen5eMastery() (+15 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 140 - "Client"
Cohesion: 0.20
Nodes (4): message, IsLoopback(), TestIsLoopback(), Client

### Community 143 - "scoped"
Cohesion: 0.24
Nodes (12): init(), parseDeathSave(), parseHitDie(), parseInspiration(), parsePact(), parseRest(), parseShield(), parseSource() (+4 more)

### Community 144 - "parseHP"
Cohesion: 0.30
Nodes (9): Amount, Damage(), parseAmount(), parseCoin(), parseDamage(), parseHP(), parseMaxHP(), parseTemp() (+1 more)

### Community 145 - "LinkTarget"
Cohesion: 0.17
Nodes (8): LinkTarget, AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink, snapshotBody

### Community 146 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 147 - "Campaign"
Cohesion: 0.28
Nodes (6): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks()

### Community 148 - "findItem"
Cohesion: 0.32
Nodes (4): findItem(), parseArmor(), parseItem(), ensureItem()

### Community 149 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 150 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 151 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 152 - "WeaponKind"
Cohesion: 0.40
Nodes (5): Completions(), parseMastery(), TestCompletions(), TestWeaponKind(), WeaponKind()

### Community 153 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 398 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Command`, `.levelUp`, `Pending`, `perm.go`, `Campaign`, `NameDesc`, `database/sql.NullInt64`, `.Router`, `Attachment`, `Book`, `formInt`, `.fail`, `cert.go`, `Sheet`, `Slugify`, `FiveToolsJob`, `.compendiumDir`, `.gearFor`, `CompendiumKind`, `.wizardView`, `viewerCtx`, `Client`, `time.Time`, `ClassInfo`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14113724394042249 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `context.Context`, `SheetView`, `.levelUp`, `findItem`, `TestRunCommandUsesCurrentSheet`, `WeaponKind`, `ParseSheet`, `NameDesc`, `testing.T`, `database/sql.NullInt64`, `ParseEffects`, `.fail`, `Slugify`, `github.com/a-h/templ.Component`, `.gearFor`, `CompendiumKind`, `encoding/json.RawMessage`, `.wizardView`, `ClassInfo`, `names`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `Import` be split into smaller, more focused modules?**
  _Cohesion score 0.09024390243902439 - nodes in this community are weakly interconnected._
- **Why does `serve()` connect `Store` to `compendium_templ.go`, `.do`, `testing.T`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010147492625368732 - nodes in this community are weakly interconnected._