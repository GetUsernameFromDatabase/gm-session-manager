# Graph Report - gm-session-manager  (2026-11-10)

## Corpus Check
- 296 files · ~571,149 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4117 nodes · 17890 edges · 144 communities (107 shown, 37 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1819 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `47dd868d`
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
- combobox_templ.go
- Store
- select.js
- Command
- .Entry
- ImagePage
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Input
- levelUpForm
- wikilink.js
- tabs_templ.go
- popover.js
- .sendFeedback
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- Sheet
- testing.T
- command.js
- dialog.js
- Client
- ParseWiki
- tooltip.js
- Layout
- ParseAidedd
- Book
- cell
- composite.js
- Item
- Feature
- ClassInfo
- github.com/a-h/templ.Component
- State
- .Router
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- time.Time
- Character
- radiogroup.js
- vendorjs/main.go
- github.com/a-h/templ.Component
- FiveToolsJob
- sqlc Config
- buildReview
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- TestDownloadFiveTools
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
- .ListEntrySourcesBySlugs
- Lint and Release Rules
- Run Script
- parseCommand
- useTypeahead
- toast_templ.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- NewClient
- tooltip_templ.go
- Session
- Props
- WeaponKind
- ftBook
- .ListBlockStepsSince
- switch_templ.go
- formErrors
- export/markdown.go
- message
- TestRunCommandUsesCurrentSheet
- .AddPending
- RosterCharacter
- parseLiveKey
- IsLoopback
- build.sh
- Hub
- nullID
- errBodyWriter
- Trigger
- capReader
- .ListItemsByCategory
- Before/after pairs
- .ListCompendiumByKind
- open5e.go
- Open
- db/CLAUDE.md
- perm/CLAUDE.md
- ftNamed
- .ListCompendiumBySlugs
- Site
- .ListSpellsForClass
- Verbs
- SearchReadableBlocksRow
- .ListCompendiumNames

## God Nodes (most connected - your core abstractions)
1. `App` - 307 edges
2. `newTestEnv()` - 122 edges
3. `Button()` - 119 edges
4. `Sheet` - 118 edges
5. `CN()` - 113 edges
6. `IfElse()` - 58 edges
7. `SheetView` - 57 edges
8. `Command` - 49 edges
9. `Input()` - 47 edges
10. `Slugify()` - 43 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `UI conventions` --references--> `selectItems()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (144 total, 37 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (60): Props, contextKey, keepAwake(), classRef, ftBookInfo, indexed, sourceOf, ASIError (+52 more)

### Community 1 - "CombatantView"
Cohesion: 0.12
Nodes (31): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), downBadge() (+23 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (104): keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListHomebrewRow, ListImageEntitiesRow, ListAttachmentFilesRow, GetImageRow (+96 more)

### Community 3 - ".do"
Cohesion: 0.15
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.08
Nodes (14): Part, Skill, ItemData, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+6 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (101): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+93 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "combobox_templ.go"
Cohesion: 0.08
Nodes (45): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+37 more)

### Community 10 - "Store"
Cohesion: 0.12
Nodes (29): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+21 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.10
Nodes (50): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+42 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (26): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInput, HomebrewInputFor() (+18 more)

### Community 14 - "ImagePage"
Cohesion: 0.14
Nodes (25): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, initials(), currentLook(), EntityDetail(), entityForm() (+17 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.10
Nodes (36): PendingState, Pending, PendingKind, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass(), CanDeleteCharacter() (+28 more)

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
Cohesion: 0.04
Nodes (75): ImageCategory, Coins, CloseFor(), Description(), Field(), Group(), Label(), Input() (+67 more)

### Community 21 - "levelUpForm"
Cohesion: 0.06
Nodes (72): Props, EquipmentItem, EquipmentOption, LevelUpPlan, ProficiencyAt(), Checkbox(), initialChecked(), attrString() (+64 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.10
Nodes (41): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+33 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.10
Nodes (40): Import(), TestUpgradeSheet(), ParseSheet(), TestBuilderCompendium(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI() (+32 more)

### Community 28 - "App"
Cohesion: 0.05
Nodes (30): SetCookie(), FromContext(), FetchPage(), Campaign, compendiumBack(), compendiumPath(), homebrewTabURL(), pathCampaign() (+22 more)

### Community 29 - "newTestEnv"
Cohesion: 0.05
Nodes (59): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait() (+51 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Sheet"
Cohesion: 0.03
Nodes (113): DeathSaves, LevelUpChoice, Cantrips(), CoreTraits(), EquipmentOptions(), firstTable(), LevelSpells(), OptionPrerequisite() (+105 more)

### Community 34 - "testing.T"
Cohesion: 0.03
Nodes (88): Defense, TestCoreTraits(), feature(), RemoveSourceDir(), TestRemoveSourceDir(), Parse(), fighter(), TestCommands() (+80 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 38 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Layout"
Cohesion: 0.06
Nodes (85): MediaProps, MediaVariant, CountHabitatsRow, ListHomebrewByOwnerRow, ListMonstersByNameRow, Input(), Content(), Description() (+77 more)

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - "Book"
Cohesion: 0.16
Nodes (18): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+10 more)

### Community 43 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Item"
Cohesion: 0.26
Nodes (5): findItem(), parseArmor(), TestGearCommands(), ensureItem(), Item

### Community 46 - "Feature"
Cohesion: 0.16
Nodes (9): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+1 more)

### Community 47 - "ClassInfo"
Cohesion: 0.10
Nodes (12): TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+4 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (73): SearchCompendiumRow, UI conventions, Badge(), Scripts(), Separator(), removeButton(), alignmentColor(), classBadges() (+65 more)

### Community 49 - "State"
Cohesion: 0.12
Nodes (11): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+3 more)

### Community 50 - ".Router"
Cohesion: 0.20
Nodes (6): ClearCookie(), writeCookie(), RequireAdmin(), sameOriginOnly(), useValidateTag(), CompendiumRest()

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

### Community 56 - "ClassData"
Cohesion: 0.10
Nodes (30): block, ClassFeature, ClassData, TestTokensReadAsLabels(), dnd5eFeatures(), InputForEntry(), TestMonsterFormRoundTrip(), withCoreTraits() (+22 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "time.Time"
Cohesion: 0.06
Nodes (19): Attachment, browserName(), firstLine(), mdLinkText(), saveScreenshot(), TestUserAgent(), userAgentSummary(), saveSniffed() (+11 more)

### Community 59 - "Character"
Cohesion: 0.05
Nodes (28): ftsQuery(), Search(), ListSubclassesOfRow, Character, Completions(), Campaign(), CharacterFileName(), fileSlug() (+20 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "github.com/a-h/templ.Component"
Cohesion: 0.08
Nodes (57): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+49 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "buildReview"
Cohesion: 0.31
Nodes (5): ensureCategory(), buildReview(), formValues(), csvImport, csvImports

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (132): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Size, ActionProps (+124 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "TestDownloadFiveTools"
Cohesion: 0.12
Nodes (20): internal/compendium, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestImportFiveToolsBook() (+12 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "planSwitch"
Cohesion: 0.25
Nodes (5): featureSwitch, featureNames(), SubclassRef, planSwitch(), sameSlug()

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.11
Nodes (21): Player(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips() (+13 more)

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

### Community 86 - "parseCommand"
Cohesion: 0.25
Nodes (9): findMod(), newMod(), parseMod(), IsModTarget(), featureCommand(), coinCommand(), commandText(), hpCommand() (+1 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (47): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+39 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.30
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 104 - "Session"
Cohesion: 0.11
Nodes (19): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+11 more)

### Community 105 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 106 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (15): ftBook, ftRec, each(), ftDecode(), ftMod(), ftName(), ftParts(), ftRecords() (+7 more)

### Community 109 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

### Community 110 - "formErrors"
Cohesion: 0.40
Nodes (3): fieldError, pickError, userError

### Community 111 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 112 - "message"
Cohesion: 0.40
Nodes (3): message, Diff(), TestDiff()

### Community 113 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.33
Nodes (6): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits()

### Community 115 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 123 - "nullID"
Cohesion: 0.06
Nodes (24): Combatant, Condition, Encounter, ListCombatantsRow, Character, applyCondOps(), combatantDefenses(), initiativeMod() (+16 more)

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Trigger"
Cohesion: 0.15
Nodes (23): Field, Content(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks(), BuilderPick, pickURL() (+15 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.08
Nodes (38): Entry, fixture, fixtures, looseString, normalizer, o5eArmor, o5eAssignment, o5eChild (+30 more)

### Community 134 - "Open"
Cohesion: 0.11
Nodes (22): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+14 more)

### Community 137 - "ftNamed"
Cohesion: 0.16
Nodes (11): ArmorData, ftClassFeature, ftClassRefs, ftItem, ftArmor(), ftCode(), ftItemData(), TestFiveToolsItemMastery() (+3 more)

### Community 139 - "Site"
Cohesion: 0.15
Nodes (12): Site, Source, fiveToolsPage(), storeSource(), upsert(), Page, pageSource(), PageSourceKey() (+4 more)

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 385 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **37 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Sheet`, `context.Context`, `.do`, `buildReview`, `Open`, `Layout`, `nullID`, `perm.go`, `State`, `.Router`, `parseLiveKey`, `cert.go`, `Hub`, `time.Time`, `Character`, `FiveToolsJob`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1091091091091091 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `testing.T`, `encoding/json.RawMessage`, `Weapon`, `SheetView`, `ParseSheet`, `Character`, `WeaponKind`, `planSwitch`, `Item`, `Feature`, `ClassInfo`, `TestRunCommandUsesCurrentSheet`, `Input`, `levelUpForm`, `nullID`, `github.com/a-h/templ.Component`?**
  _High betweenness centrality (0.017) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.1226215644820296 - nodes in this community are weakly interconnected._
- **Why does `Feature` connect `Feature` to `compendium_templ.go`, `Sheet`, `Weapon`, `SheetView`, `planSwitch`, `Input`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.00957193657923585 - nodes in this community are weakly interconnected._