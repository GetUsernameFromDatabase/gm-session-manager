# Graph Report - gm-session-manager  (2026-11-17)

## Corpus Check
- 337 files · ~603,742 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4398 nodes · 21115 edges · 143 communities (113 shown, 30 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2122 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `03ca5bf0`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Session
- context.Context
- testing.T
- createProvider
- Sheet
- SheetView
- combobox.js
- redirect.go
- ClassInfo
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
- formSelect
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- .levelUp
- Parse
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- ParseWikidot
- command.js
- dialog.js
- DownloadFiveTools
- formErrors
- tooltip.js
- dialog_templ.go
- CompendiumPath
- Book
- export/markdown.go
- composite.js
- time.Time
- Feature
- Habitat
- cell
- SaveHomebrew
- .do
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- userAgentSummary
- golangci-lint Config
- subclassPanels
- github.com/a-h/templ.Component
- radiogroup.js
- vendorjs/main.go
- Encounter
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Weapon
- use_transition_status.js
- playwright
- release.sh
- ImagePage
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
- Slugify
- Lint and Release Rules
- Run Script
- planSwitch
- useTypeahead
- CompendiumKind
- Validator Rule
- Go Module
- .editEntity
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- TestLiveUpdatesAndKick
- newLiveFixture
- profFocus
- Client
- Campaign
- ftBook
- .sourceIssues
- .featChoices
- sourcesEnv
- config.go
- RosterCharacter
- .equipmentChoices
- credentialsMatch
- build.sh
- .renderCompendium
- looseString
- errBodyWriter
- ParseEffects
- net/url.Values
- Before/after pairs
- viewerCtx
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- isToggle
- sourceUsedError
- ListSharedWithRow
- rowFigure
- TestRunCommandUsesCurrentSheet

## God Nodes (most connected - your core abstractions)
1. `App` - 358 edges
2. `newTestEnv()` - 143 edges
3. `Sheet` - 126 edges
4. `Button()` - 121 edges
5. `CN()` - 113 edges
6. `SheetView` - 76 edges
7. `IfElse()` - 63 edges
8. `Command` - 59 edges
9. `CompendiumKind` - 56 edges
10. `Queries` - 48 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
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

## Communities (143 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (52): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+44 more)

### Community 1 - "Session"
Cohesion: 0.09
Nodes (13): IsLoopback(), TestTokens(), RequireAdmin(), hashToken(), newToken(), Session, Role, assetETags() (+5 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (82): keepAwake(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Note, SearchReadableBlocksRow (+74 more)

### Community 3 - "testing.T"
Cohesion: 0.03
Nodes (115): TestIsLoopback(), TestDecode(), RemoveSourceDir(), TestRemoveSourceDir(), TestMasteryLimit(), TestDropFeat(), TestSheetPinsSource(), TestCanChangeSubclass() (+107 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Sheet"
Cohesion: 0.08
Nodes (12): DeathSaves, TestCustomMods(), Sum(), TestGearCommands(), Coins, Item, Sheet, WornArmor (+4 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (103): ArmorEffect(), Signed(), Variant, Action(), Card(), Content(), Description(), Header() (+95 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "ClassInfo"
Cohesion: 0.11
Nodes (27): Queries, CompendiumEntry, TestSlotMaxima(), ClassInfo, ClassLevel, CasterLevel(), SlotMaxima(), SourceRef (+19 more)

### Community 10 - "Store"
Cohesion: 0.05
Nodes (58): fixture, tokens, sourceFormat, isCleared(), newFixture(), serve(), TestMiddlewareIdentity(), TestOnlyHashIsStored() (+50 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.06
Nodes (68): migration, Amount, CondOpKind, Effects, HelpEntry, parser, apply(), Completions() (+60 more)

### Community 13 - ".Entry"
Cohesion: 0.07
Nodes (33): MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, habitatsFrom(), ValidKind() (+25 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (45): ListCharacterNamesRow, GetBlockRow, Visibility, Entity, Look, VisibleEntity(), VisibleLooks(), BlockVisibilities() (+37 more)

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

### Community 21 - "formSelect"
Cohesion: 0.05
Nodes (79): ImageCategory, UI conventions, Checkbox(), Label(), initialValue(), Item(), RadioGroup(), state() (+71 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (39): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+31 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (39): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Campaign, Character, CanChangeSubclass() (+31 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (202): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+194 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.09
Nodes (44): Import(), TestUpgradeSheet(), ParseSheet(), TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestMonsterImageLink(), TestBuilderCompendium(), TestChangeSubclass() (+36 more)

### Community 28 - ".levelUp"
Cohesion: 0.07
Nodes (40): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), FreeCast(), SwapFeatSpell(), LevelUpPlan (+32 more)

### Community 29 - "Parse"
Cohesion: 0.06
Nodes (38): Heal(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestConIncreaseAddsHPForEarlierLevels(), TestDefenses() (+30 more)

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
Nodes (42): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Attachment, IsConditionSlug(), EncounterStatus, firstLine(), mdLinkText() (+34 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.11
Nodes (27): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+19 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (25): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestImportFiveToolsBook() (+17 more)

### Community 38 - "formErrors"
Cohesion: 0.25
Nodes (6): Handlers, keepsBody(), formErrors(), fieldError, pickError, userError

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "dialog_templ.go"
Cohesion: 0.08
Nodes (73): ctxKey, ctxState, Props, Button(), Dialog(), Close(), CloseFor(), Content() (+65 more)

### Community 41 - "CompendiumPath"
Cohesion: 0.09
Nodes (50): Props, Variant, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props (+42 more)

### Community 42 - "Book"
Cohesion: 0.11
Nodes (25): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+17 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "time.Time"
Cohesion: 0.07
Nodes (18): ListAttachmentFilesRow, ListImagesRow, inCampaign(), ensureCategory(), buildReview(), formValues(), FiveToolsView, NewFiveToolsView() (+10 more)

### Community 46 - "Feature"
Cohesion: 0.15
Nodes (8): FeatureEffects, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 47 - "Habitat"
Cohesion: 0.10
Nodes (29): CountHabitatsRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, Cell(), Row(), toJSON(), campaignSourceRow() (+21 more)

### Community 48 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 49 - "SaveHomebrew"
Cohesion: 0.12
Nodes (16): HomebrewOwner, Site, Source, fiveToolsPage(), SaveHomebrew(), SeedHabitats(), storeSource(), upsert() (+8 more)

### Community 50 - ".do"
Cohesion: 0.14
Nodes (18): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+10 more)

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

### Community 56 - "userAgentSummary"
Cohesion: 0.50
Nodes (3): browserName(), TestUserAgent(), userAgentSummary()

### Community 57 - "golangci-lint Config"
Cohesion: 0.07
Nodes (29): deny, files, rules, generated, paths, rules, formatters, enable (+21 more)

### Community 58 - "subclassPanels"
Cohesion: 0.20
Nodes (8): ListSubclassesOfRow, TestSubclassPanels(), subclassIndex(), subclassPanels(), SubclassPanel, ListSubclassesOfParams, SubclassChangeView, SubclassFeature

### Community 59 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (96): Alert(), Description(), Title(), Kbd(), Scripts(), Separator(), Spinner(), Toast() (+88 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Encounter"
Cohesion: 0.08
Nodes (20): Combatant, Encounter, ListCombatantsRow, passTurn(), rollFor(), stepTurn(), formInt(), nullID() (+12 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.06
Nodes (64): Defense, Condition, Defenses, ParseDefenses(), Description(), Header(), state(), Title() (+56 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Weapon"
Cohesion: 0.13
Nodes (11): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+3 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "ImagePage"
Cohesion: 0.09
Nodes (31): ListEntityFieldBacklinksRow, ListBacklinksRow, Field, initialChecked(), Switch(), removeButton(), initials(), EntityDetail() (+23 more)

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
Nodes (40): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+32 more)

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

### Community 83 - "Slugify"
Cohesion: 0.15
Nodes (21): Cantrips(), Decode(), NameDesc, Slugify(), AddResistance(), catchUpMasteries(), catchUpResistances(), TestCatchUpResistances() (+13 more)

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "planSwitch"
Cohesion: 0.19
Nodes (8): featureSwitch, featureNames(), SubclassRef, SwitchPlan, PlanSourceSwitch(), planSwitch(), sameSlug(), SourcePreviewView

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "CompendiumKind"
Cohesion: 0.05
Nodes (37): dataFor(), DecodeData(), UpsertCompendiumEntryParams, Character, CompendiumKind, viewerCampaign(), checkSkillPick(), classSkillPick() (+29 more)

### Community 96 - ".editEntity"
Cohesion: 0.09
Nodes (12): Entity, Fields, ParseFields(), entityTopics(), linkEntityLook(), saveFieldLinks(), linkCharacterLook(), lookFree() (+4 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (46): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+38 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 103 - "newLiveFixture"
Cohesion: 0.10
Nodes (25): BenchmarkCompendiumList(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas, internal/web, Notes, Websocket and live views (+17 more)

### Community 105 - "Client"
Cohesion: 0.06
Nodes (34): ClientInfo, event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch (+26 more)

### Community 106 - "Campaign"
Cohesion: 0.32
Nodes (5): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (23): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftMod(), ftName() (+15 more)

### Community 108 - ".sourceIssues"
Cohesion: 0.43
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 109 - ".featChoices"
Cohesion: 0.11
Nodes (21): internal/compendium, SkillChoice(), TestSkillChoice(), SpellData, TextData, AbilityKeys(), AddFeat(), SkillKeys() (+13 more)

### Community 111 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 116 - ".equipmentChoices"
Cohesion: 0.15
Nodes (12): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), ItemData, compendiumWeapon(), TestCompendiumWeapon(), applyEquipment() (+4 more)

### Community 117 - "credentialsMatch"
Cohesion: 0.67
Nodes (3): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".renderCompendium"
Cohesion: 0.07
Nodes (22): ftsQuery(), Search(), CountCompendiumRow, SearchCompendiumRow, ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers() (+14 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.21
Nodes (6): Slot, errorText(), TestErrorText(), modTargetLabel(), slotLabel(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.13
Nodes (10): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+2 more)

### Community 127 - "net/url.Values"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "viewerCtx"
Cohesion: 0.15
Nodes (10): PendingState, Pending, PendingKind, notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind (+2 more)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 146 - "rowFigure"
Cohesion: 0.50
Nodes (4): TestRowFigure(), crText(), rowFigure(), textOnly()

### Community 150 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 399 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Session`, `context.Context`, `testing.T`, `viewerCtx`, `Sheet`, `ClassInfo`, `command.go`, `perm.go`, `.levelUp`, `database/sql.NullInt64`, `Book`, `time.Time`, `.do`, `cert.go`, `Encounter`, `FiveToolsJob`, `Slugify`, `CompendiumKind`, `.editEntity`, `Client`, `Campaign`, `.sourceIssues`, `.featChoices`, `.equipmentChoices`, `.renderCompendium`?**
  _High betweenness centrality (0.060) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13866238120746782 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `SheetView`, `ClassInfo`, `command.go`, `TestRunCommandUsesCurrentSheet`, `ParseSheet`, `.levelUp`, `Parse`, `Feature`, `Encounter`, `CombatantView`, `Weapon`, `Slugify`, `planSwitch`, `CompendiumKind`, `encoding/json.RawMessage`, `.sourceIssues`, `.featChoices`, `.equipmentChoices`, `errBodyWriter`, `ParseEffects`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `Session` be split into smaller, more focused modules?**
  _Cohesion score 0.09195402298850575 - nodes in this community are weakly interconnected._
- **Why does `liveFixture` connect `newLiveFixture` to `compendium_templ.go`, `perm.go`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010146704434955398 - nodes in this community are weakly interconnected._