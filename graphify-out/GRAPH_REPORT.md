# Graph Report - gm-session-manager  (2026-11-10)

## Corpus Check
- 296 files · ~577,929 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4152 nodes · 18044 edges · 141 communities (110 shown, 31 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1847 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `65c3ce11`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- context.Context
- .do
- createProvider
- Weapon
- SheetView
- combobox.js
- redirect.go
- Layout
- Import
- select.js
- Command
- .Entry
- select_templ.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Button
- levelUpForm
- wikilink.js
- IfElse
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- Item
- testing.T
- command.js
- dialog.js
- Client
- Queries
- tooltip.js
- Sum
- ParseAidedd
- Book
- cell
- composite.js
- Campaign
- Feature
- ClassInfo
- monsterBody
- State
- parseHP
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- ftSpell
- Character
- radiogroup.js
- vendorjs/main.go
- Input
- FiveToolsJob
- sqlc Config
- time.Time
- Browser check
- Claude Settings
- parseCommand
- use_transition_status.js
- playwright
- release.sh
- Parse
- checkbox.js
- planSwitch
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- fromLookup
- useFocus
- mark_others.js
- WeaponKind
- Lint and Release Rules
- Run Script
- useTypeahead
- toast_templ.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- joinNonEmpty
- mounted
- NewClient
- Sheet
- Session
- popover_templ.go
- RosterCharacter
- ftBook
- .encounterView
- formErrors
- export/markdown.go
- Open
- TestClassOptionChoices
- Pending
- Topic
- Layout
- build.sh
- errBodyWriter
- capReader
- .ListItemsByCategory
- Before/after pairs
- open5e.go
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- .ListCompendiumBySlugs
- Site
- .ListSpellsForClass
- dialog_templ.go
- Campaign
- ParseWiki
- radiogroup_templ.go
- Client
- ListSharedWithRow
- .ListCompendiumByKind
- .ListSharedOrder

## God Nodes (most connected - your core abstractions)
1. `App` - 311 edges
2. `newTestEnv()` - 123 edges
3. `Button()` - 120 edges
4. `Sheet` - 118 edges
5. `CN()` - 113 edges
6. `IfElse()` - 61 edges
7. `SheetView` - 57 edges
8. `Command` - 49 edges
9. `Input()` - 47 edges
10. `Slugify()` - 43 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `UI conventions` --references--> `removeButton()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (141 total, 31 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (68): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+60 more)

### Community 1 - "CombatantView"
Cohesion: 0.10
Nodes (36): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+28 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (113): keepAwake(), ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumSource (+105 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (87): ArmorEffect(), Signed(), Variant, Action(), Card(), Content(), Description(), Header() (+79 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Layout"
Cohesion: 0.05
Nodes (83): MediaProps, MediaVariant, ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListBacklinksRow, SearchReadableBlocksRow, Content() (+75 more)

### Community 10 - "Import"
Cohesion: 0.12
Nodes (29): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+21 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.11
Nodes (42): CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Command, CondOp (+34 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (29): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput (+21 more)

### Community 14 - "select_templ.go"
Cohesion: 0.06
Nodes (50): ImageCategory, Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select() (+42 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.14
Nodes (32): VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock() (+24 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Button"
Cohesion: 0.06
Nodes (64): Coins, Field, Button(), CloseFor(), Group(), Content(), DisableTrigger(), state() (+56 more)

### Community 21 - "levelUpForm"
Cohesion: 0.05
Nodes (59): spellbook(), spellNeed(), spellTotals(), folderName(), mdCell(), oneLine(), parseConsole(), slugify() (+51 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "IfElse"
Cohesion: 0.05
Nodes (72): CountHabitatsRow, CountSharedBySourceRow, ListHomebrewLinksRow, Entity, Look, Content(), contentID(), IDFromContext() (+64 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (152): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+144 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.11
Nodes (34): internal/compendium, Import(), TestUpgradeSheet(), ParseSheet(), TestBuilderCompendium(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats() (+26 more)

### Community 28 - "App"
Cohesion: 0.04
Nodes (37): ftsQuery(), Search(), Entity, GetBlockRow, compendiumBack(), compendiumPath(), homebrewTabURL(), pathCampaign() (+29 more)

### Community 29 - "newTestEnv"
Cohesion: 0.04
Nodes (70): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait() (+62 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 34 - "testing.T"
Cohesion: 0.04
Nodes (73): TestCoreTraits(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestImportFiveToolsBook() (+65 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 38 - "Queries"
Cohesion: 0.07
Nodes (34): dataFor(), DecodeData(), Kind, ListEntryCopiesRow, Queries, CompendiumEntry, ClassInfo, ClassLevel (+26 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Sum"
Cohesion: 0.21
Nodes (3): TestCustomMods(), TestExplainSums(), Sum()

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - "Book"
Cohesion: 0.17
Nodes (17): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+9 more)

### Community 43 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Campaign"
Cohesion: 0.17
Nodes (7): Campaign, TestUserAgent(), userAgentSummary(), visitorOnly(), Roster, CreateCampaignParams, CampaignSwitch

### Community 46 - "Feature"
Cohesion: 0.13
Nodes (13): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+5 more)

### Community 47 - "ClassInfo"
Cohesion: 0.10
Nodes (13): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+5 more)

### Community 48 - "monsterBody"
Cohesion: 0.10
Nodes (35): SearchCompendiumRow, alignmentColor(), classBadges(), classBody(), commaPrefixed(), crText(), entryBody(), EntryDetails() (+27 more)

### Community 49 - "State"
Cohesion: 0.15
Nodes (9): Handler, message, patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff() (+1 more)

### Community 50 - "parseHP"
Cohesion: 0.39
Nodes (7): Amount, Damage(), Heal(), parseAmount(), parseDamage(), parseHP(), Set()

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

### Community 56 - "ClassData"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "ftSpell"
Cohesion: 0.40
Nodes (6): ftDistance, ftSpell, ftAmount(), ftCastingTime(), ftDuration(), ftRange()

### Community 59 - "Character"
Cohesion: 0.07
Nodes (14): ListSubclassesOfRow, Character, Completions(), CanCharacter(), subclassIndex(), subclassOptions(), subclassPanels(), nextStep() (+6 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Input"
Cohesion: 0.09
Nodes (46): ctxState, Chip(), Chips(), Combobox(), Content(), Input(), Item(), itemLabel() (+38 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "time.Time"
Cohesion: 0.07
Nodes (22): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), mdLinkText() (+14 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "parseCommand"
Cohesion: 0.40
Nodes (5): featureCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Parse"
Cohesion: 0.06
Nodes (36): Defense, feature(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses() (+28 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "planSwitch"
Cohesion: 0.24
Nodes (6): featureSwitch, Env, featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.12
Nodes (19): Player(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips() (+11 more)

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
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 100 - "joinNonEmpty"
Cohesion: 0.07
Nodes (41): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+33 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Sheet"
Cohesion: 0.04
Nodes (88): EquipmentItem, DeathSaves, Cantrips(), CoreTraits(), EquipmentOptions(), firstTable(), EquipmentOption, LevelSpells() (+80 more)

### Community 104 - "Session"
Cohesion: 0.08
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), TestCredentialsMatch(), TestIsLoopback(), TestRequireAdmin() (+10 more)

### Community 105 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 106 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 107 - "ftBook"
Cohesion: 0.19
Nodes (16): ftBook, ftRec, each(), ftDecode(), ftMod(), ftName(), ftParts(), ftRecords() (+8 more)

### Community 109 - ".encounterView"
Cohesion: 0.11
Nodes (17): CountCompendiumRow, Combatant, ListCombatantsRow, Character, combatantDefenses(), damageChoice(), initiativeMod(), outOfFight() (+9 more)

### Community 110 - "formErrors"
Cohesion: 0.40
Nodes (3): fieldError, pickError, userError

### Community 111 - "export/markdown.go"
Cohesion: 0.31
Nodes (9): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+1 more)

### Community 112 - "Open"
Cohesion: 0.13
Nodes (23): migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open() (+15 more)

### Community 113 - "TestClassOptionChoices"
Cohesion: 0.15
Nodes (14): UpsertCompendiumEntryParams, SwitchFeature, SwitchPlan, PlanSourceSwitch(), names(), TestPlanSourceSwitch(), classCopy(), testEnv (+6 more)

### Community 114 - "Pending"
Cohesion: 0.18
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 116 - "Topic"
Cohesion: 0.14
Nodes (7): event, fakeHandler, stateHandler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 117 - "Layout"
Cohesion: 0.07
Nodes (63): Props, UI conventions, Checkbox(), initialChecked(), Label(), Scripts(), Separator(), ComparePage() (+55 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.14
Nodes (13): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+5 more)

### Community 139 - "Site"
Cohesion: 0.14
Nodes (13): Site, Source, fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page, pageSource() (+5 more)

### Community 142 - "dialog_templ.go"
Cohesion: 0.13
Nodes (39): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+31 more)

### Community 143 - "Campaign"
Cohesion: 0.32
Nodes (5): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

### Community 144 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 152 - "radiogroup_templ.go"
Cohesion: 0.18
Nodes (15): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+7 more)

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 384 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **31 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Layout`, `Import`, `Campaign`, `perm.go`, `Client`, `Queries`, `Campaign`, `State`, `cert.go`, `Character`, `FiveToolsJob`, `time.Time`, `Sheet`, `Session`, `.encounterView`, `Topic`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1027442464499372 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Item`, `joinNonEmpty`, `Weapon`, `Queries`, `SheetView`, `Sum`, `Parse`, `ParseSheet`, `planSwitch`, `.encounterView`, `Feature`, `ClassInfo`, `TestClassOptionChoices`, `WeaponKind`, `Button`, `levelUpForm`, `Character`?**
  _High betweenness centrality (0.019) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.10196078431372549 - nodes in this community are weakly interconnected._
- **Why does `Queries` connect `Queries` to `compendium_templ.go`, `context.Context`, `.do`, `Import`, `Site`, `.Entry`, `Campaign`, `Open`, `Character`, `App`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009639109881240632 - nodes in this community are weakly interconnected._