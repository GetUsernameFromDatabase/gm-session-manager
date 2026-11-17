# Graph Report - gm-session-manager  (2026-11-17)

## Corpus Check
- 335 files · ~599,041 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4362 nodes · 20949 edges · 148 communities (115 shown, 33 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2100 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `267c195c`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Session
- context.Context
- newTestEnv
- createProvider
- ImportView
- SheetView
- combobox.js
- redirect.go
- spellCounts
- Import
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
- DownloadFiveTools
- formErrors
- tooltip.js
- shadcn-templ.go
- table_templ.go
- Book
- export/markdown.go
- composite.js
- Encounter
- Feature
- CompendiumPath
- popover_templ.go
- SaveHomebrew
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Weapon
- golangci-lint Config
- .levelUp
- github.com/a-h/templ.Component
- radiogroup.js
- vendorjs/main.go
- Campaign
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- who
- use_transition_status.js
- playwright
- release.sh
- CN
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
- .characterAccess
- Lint and Release Rules
- Run Script
- .Router
- useTypeahead
- CompendiumKind
- Validator Rule
- Go Module
- Client
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- .compendiumDir
- newLiveFixture
- profFocus
- State
- .wizardView
- ftBook
- deviceLabel
- LinkTarget
- viewerCtx
- config.go
- Client
- .whoami
- NewClient
- build.sh
- ListSharedWithRow
- looseString
- errBodyWriter
- CompendiumSource
- names
- net/url.Values
- Before/after pairs
- DBTX
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- buildReview
- ListCompendiumBySlugsParams
- isToggle
- sourceUsedError
- lanHostname
- sortedKeys
- Character
- TestRunCommandUsesCurrentSheet
- TestLiveUpdatesAndKick

## God Nodes (most connected - your core abstractions)
1. `App` - 351 edges
2. `newTestEnv()` - 135 edges
3. `Sheet` - 126 edges
4. `Button()` - 122 edges
5. `CN()` - 113 edges
6. `SheetView` - 76 edges
7. `IfElse()` - 63 edges
8. `Command` - 59 edges
9. `CompendiumKind` - 56 edges
10. `ParseSheet()` - 47 edges

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

## Communities (148 total, 33 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "Session"
Cohesion: 0.08
Nodes (23): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+15 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (85): keepAwake(), keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListImageEntitiesRow, SearchReadableBlocksRow, IsConditionSlug(), EntityKind (+77 more)

### Community 3 - "newTestEnv"
Cohesion: 0.04
Nodes (72): nullID(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestEncounterTransitions(), TestMonsterPortrait(), TestDeleteRestoreImport(), TestRenameDeleteCampaign(), testEnv (+64 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "ImportView"
Cohesion: 0.26
Nodes (9): categoryRow(), columnOverview(), ImportReview(), importRows(), keepSection(), columnLabel(), ImportCategory, ImportView (+1 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (113): ArmorEffect(), Coins, Signed(), Field, Button(), Variant, Action(), Card() (+105 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

### Community 10 - "Import"
Cohesion: 0.06
Nodes (55): migration, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+47 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.06
Nodes (71): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+63 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (27): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor() (+19 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.08
Nodes (41): GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass() (+33 more)

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
Cohesion: 0.10
Nodes (43): Props, Footer(), Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup() (+35 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "App"
Cohesion: 0.06
Nodes (11): entityTopics(), formIntQuery(), isHTMX(), setAudience(), securityHeaders(), joinPath(), DisplayPage(), JoinQRSrc() (+3 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (129): Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props (+121 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.11
Nodes (34): ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage() (+26 more)

### Community 28 - "NameDesc"
Cohesion: 0.08
Nodes (44): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, TextData (+36 more)

### Community 29 - "testing.T"
Cohesion: 0.03
Nodes (107): TestDecode(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), RemoveSourceDir(), TestRemoveSourceDir(), Parse(), fighter() (+99 more)

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
Cohesion: 0.07
Nodes (33): GetImageRow, ListImagesRow, Attachment, Entity, Note, CreateCombatantParams, NoteKind, saveScreenshot() (+25 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.12
Nodes (26): block, ClassFeature, ClassData, dnd5eFeatures(), TestFiveToolsClassLevelUp(), withCoreTraits(), cellAt(), classTable() (+18 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "DownloadFiveTools"
Cohesion: 0.11
Nodes (26): FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook() (+18 more)

### Community 38 - "formErrors"
Cohesion: 0.29
Nodes (5): Handlers, formErrors(), fieldError, pickError, userError

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "shadcn-templ.go"
Cohesion: 0.07
Nodes (27): Props, AspectRatio(), initialChecked(), Switch(), DisableTrigger(), state(), Direction(), DirectionAttributes() (+19 more)

### Community 41 - "table_templ.go"
Cohesion: 0.12
Nodes (21): ListMonstersByNameRow, Body(), Caption(), Cell(), Footer(), Head(), Header(), Row() (+13 more)

### Community 42 - "Book"
Cohesion: 0.18
Nodes (14): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+6 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Encounter"
Cohesion: 0.06
Nodes (32): Defense, Combatant, Condition, Encounter, ListCombatantsRow, Defenses, ParseDefenses(), EncounterStatus (+24 more)

### Community 46 - "Feature"
Cohesion: 0.07
Nodes (20): FeatureEffects, LevelUpChoice, AddFeat(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs() (+12 more)

### Community 47 - "CompendiumPath"
Cohesion: 0.05
Nodes (75): CountHabitatsRow, CountSharedBySourceRow, SearchCompendiumRow, UI conventions, classBadges(), classBody(), ComparePage(), CompendiumResults() (+67 more)

### Community 48 - "popover_templ.go"
Cohesion: 0.15
Nodes (20): Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition(), conditionList() (+12 more)

### Community 49 - "SaveHomebrew"
Cohesion: 0.16
Nodes (10): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+2 more)

### Community 50 - "cell"
Cohesion: 0.12
Nodes (19): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+11 more)

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

### Community 56 - "Weapon"
Cohesion: 0.06
Nodes (19): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+11 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.08
Nodes (26): rules, generated, paths, rules, formatters, enable, exclusions, settings (+18 more)

### Community 58 - ".levelUp"
Cohesion: 0.09
Nodes (27): ListSubclassesOfRow, FreeCast(), SwapFeatSpell(), ClassFeature, LevelUpPlan, Spell, spellFromRow(), addClassSpells() (+19 more)

### Community 59 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (147): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListEntityFieldBacklinksRow (+139 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Campaign"
Cohesion: 0.19
Nodes (6): ListAttachmentFilesRow, Campaign, pathCampaign(), inCampaign(), visitorOnly(), CampaignSwitch

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (38): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+30 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "who"
Cohesion: 0.09
Nodes (12): compendiumBack(), compendiumPath(), homebrewTabURL(), viewerCampaign(), redirect(), CompendiumBase(), CompendiumRest(), CompendiumURL() (+4 more)

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
Cohesion: 0.05
Nodes (107): ActionProps, DescriptionProps, Props, TitleProps, Variant, ContentProps, ctxKey, ctxState (+99 more)

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
Nodes (42): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+34 more)

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

### Community 83 - ".characterAccess"
Cohesion: 0.21
Nodes (4): CanCharacter(), BuilderPickItems(), BuilderPick, pickURL()

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - ".Router"
Cohesion: 0.21
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "CompendiumKind"
Cohesion: 0.04
Nodes (54): HomebrewOwner, featureSwitch, dataFor(), Decode(), DecodeData(), SaveHomebrew(), SeedHabitats(), CountCompendiumRow (+46 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (50): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+42 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".compendiumDir"
Cohesion: 0.10
Nodes (17): FiveToolsBooks(), ImportSummary(), fiveToolsKey(), loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpRedownloads() (+9 more)

### Community 103 - "newLiveFixture"
Cohesion: 0.11
Nodes (24): Player(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas, internal/web, Notes, Websocket and live views (+16 more)

### Community 105 - "State"
Cohesion: 0.09
Nodes (13): event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch, Topic (+5 more)

### Community 106 - ".wizardView"
Cohesion: 0.06
Nodes (36): EquipmentItem, internal/compendium, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), ItemData, SpellData (+28 more)

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (21): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftCopyTargets(), ftDecode() (+13 more)

### Community 108 - "deviceLabel"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 109 - "LinkTarget"
Cohesion: 0.17
Nodes (8): LinkTarget, AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink, snapshotBody

### Community 111 - "viewerCtx"
Cohesion: 0.15
Nodes (10): PendingState, Pending, PendingKind, notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind (+2 more)

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "Client"
Cohesion: 0.21
Nodes (3): ClientInfo, Hub, randomID()

### Community 116 - ".whoami"
Cohesion: 0.08
Nodes (17): IsLoopback(), TestIsLoopback(), ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName() (+9 more)

### Community 117 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "CompendiumSource"
Cohesion: 0.22
Nodes (7): ListEntryCopiesRow, CompendiumSource, ListEntryCopiesParams, CompareCopy, CompareView, EntryView, HabitatCopy

### Community 126 - "names"
Cohesion: 0.40
Nodes (5): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), SourcePreviewView

### Community 127 - "net/url.Values"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (23): Entry, fixture, fixtures, normalizer, o5eProperty, Source, sourceFormat, TestOpen5eMastery() (+15 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "buildReview"
Cohesion: 0.36
Nodes (4): buildReview(), formValues(), csvImport, csvImports

### Community 140 - "ListCompendiumBySlugsParams"
Cohesion: 0.17
Nodes (6): ListCompendiumBySlugsParams, ListCompendiumBySlugsRow, ListEntrySourcesBySlugsParams, ListEntrySourcesBySlugsRow, ListSharedEntriesParams, ListSharedEntriesRow

### Community 143 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 147 - "Character"
Cohesion: 0.08
Nodes (15): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks(), linkEntityLook() (+7 more)

### Community 150 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 151 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 400 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **33 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `newTestEnv`, `buildReview`, `command.go`, `perm.go`, `Character`, `NameDesc`, `database/sql.NullInt64`, `Encounter`, `cert.go`, `Weapon`, `.levelUp`, `Campaign`, `FiveToolsJob`, `who`, `.characterAccess`, `.Router`, `CompendiumKind`, `.compendiumDir`, `State`, `.wizardView`, `viewerCtx`, `Client`, `.whoami`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14113724394042249 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Client` to `compendium_templ.go`, `Client`, `State`, `viewerCtx`, `NewClient`, `App`, `FiveToolsJob`?**
  _High betweenness centrality (0.013) - this node is a cross-community bridge._
- **Should `Session` be split into smaller, more focused modules?**
  _Cohesion score 0.08108108108108109 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Weapon` to `compendium_templ.go`, `encoding/json.RawMessage`, `SheetView`, `.wizardView`, `command.go`, `Encounter`, `Feature`, `TestRunCommandUsesCurrentSheet`, `CompendiumKind`, `.levelUp`, `ParseSheet`, `NameDesc`, `testing.T`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.008856345885634588 - nodes in this community are weakly interconnected._