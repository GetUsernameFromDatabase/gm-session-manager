# Graph Report - gm-session-manager  (2026-11-17)

## Corpus Check
- 337 files · ~602,735 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4388 nodes · 21088 edges · 147 communities (116 shown, 31 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2118 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `63c5f840`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Session
- context.Context
- newTestEnv
- createProvider
- Sheet
- SheetView
- combobox.js
- redirect.go
- ClassInfo
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
- IfElse
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
- dialog_templ.go
- github.com/a-h/templ.Component
- Book
- export/markdown.go
- composite.js
- parseCommand
- Feature
- tooltip_templ.go
- cell
- SaveHomebrew
- .do
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Visibility
- golangci-lint Config
- .levelUp
- CompendiumPath
- radiogroup.js
- vendorjs/main.go
- .combatantHP
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- InputForEntry
- use_transition_status.js
- playwright
- release.sh
- Input
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
- Layout
- Lint and Release Rules
- Run Script
- Open
- useTypeahead
- CompendiumKind
- Validator Rule
- Go Module
- LinkTarget
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- .compendiumDir
- newLiveFixture
- profFocus
- State
- Character
- ftBook
- deviceLabel
- Slugify
- viewerCtx
- config.go
- Client
- EquipmentOption
- NewClient
- build.sh
- writeBook
- looseString
- errBodyWriter
- ParseEffects
- WeaponKind
- net/url.Values
- Before/after pairs
- Pending
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- isToggle
- sourceUsedError
- toast_templ.go
- rowFigure
- keepawake_windows.go
- SearchReadableBlocksRow
- TestRunCommandUsesCurrentSheet
- TestLiveUpdatesAndKick

## God Nodes (most connected - your core abstractions)
1. `App` - 358 edges
2. `newTestEnv()` - 142 edges
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
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
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

## Communities (147 total, 31 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (52): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+44 more)

### Community 1 - "Session"
Cohesion: 0.10
Nodes (19): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+11 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (80): keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListSharedWithRow, ListImageEntitiesRow, CompendiumSource, Encounter, Note (+72 more)

### Community 3 - "newTestEnv"
Cohesion: 0.04
Nodes (72): nullID(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestDeleteActiveCombatant(), TestDeleteActiveCombatantLast(), TestDeleteActiveCombatantSkipsOut(), TestDeleteCharacterActiveCombatant() (+64 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Sheet"
Cohesion: 0.05
Nodes (29): DeathSaves, Part, Skill, ItemData, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums() (+21 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (114): ArmorEffect(), Signed(), Field, Button(), Variant, Card(), Content(), Description() (+106 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "ClassInfo"
Cohesion: 0.08
Nodes (32): featureSwitch, Queries, Env, ClassInfo, ClassLevel, featureNames(), SourceRef, SubclassRef (+24 more)

### Community 10 - "Import"
Cohesion: 0.10
Nodes (36): sourceFormat, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+28 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.08
Nodes (60): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+52 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (24): MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, HomebrewInputFor(), abilityKey() (+16 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (58): GetBlockRow, ListCombatantsRow, Entity, Fields, Look, ParseFields(), VisibleEntity(), VisibleLooks() (+50 more)

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

### Community 21 - "IfElse"
Cohesion: 0.07
Nodes (64): Props, ctxState, Props, Description(), Title(), Checkbox(), initialChecked(), Chip() (+56 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (39): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+31 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (25): ClearCookie(), SetCookie(), writeCookie(), FromContext(), homebrewTabURL(), pathCampaign(), entityTopics(), formInt() (+17 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (129): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+121 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.11
Nodes (35): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+27 more)

### Community 28 - "NameDesc"
Cohesion: 0.18
Nodes (19): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+11 more)

### Community 29 - "testing.T"
Cohesion: 0.03
Nodes (89): Defense, TestIsLoopback(), TestDecode(), RemoveSourceDir(), TestRemoveSourceDir(), Heal(), Parse(), tokenize() (+81 more)

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
Nodes (57): Identity, ListHomebrewByOwnerRow, GetImageRow, ListImagesRow, Attachment, Condition, Entity, EntityKind (+49 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (25): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook(), DownloadFiveTools() (+17 more)

### Community 38 - "formErrors"
Cohesion: 0.33
Nodes (4): Handlers, fieldError, pickError, userError

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "dialog_templ.go"
Cohesion: 0.11
Nodes (44): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+36 more)

### Community 41 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (96): ListMonstersByNameRow, Action(), AspectRatio(), Action(), Footer(), Chips(), ChipsInput(), Empty() (+88 more)

### Community 42 - "Book"
Cohesion: 0.13
Nodes (20): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+12 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "parseCommand"
Cohesion: 0.40
Nodes (5): featureCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 46 - "Feature"
Cohesion: 0.17
Nodes (8): FeatureEffects, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 47 - "tooltip_templ.go"
Cohesion: 0.12
Nodes (24): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks(), BuilderPick (+16 more)

### Community 48 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 49 - "SaveHomebrew"
Cohesion: 0.13
Nodes (14): Site, Source, TestStoreSourceRollsBackOnAfterSave(), fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page (+6 more)

### Community 50 - ".do"
Cohesion: 0.13
Nodes (19): testEnv, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm() (+11 more)

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

### Community 56 - "Visibility"
Cohesion: 0.12
Nodes (11): ListAttachmentFilesRow, Visibility, inCampaign(), visibilityLabel(), CreateBlockParams, ListBlockDocsRow, ListExportBlocksRow, ListReadableBlocksParams (+3 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.08
Nodes (26): rules, generated, paths, rules, formatters, enable, exclusions, settings (+18 more)

### Community 58 - ".levelUp"
Cohesion: 0.10
Nodes (22): FreeCast(), SwapFeatSpell(), LevelUpPlan, Spell, Spellbook(), TestSpellbook(), addClassSpells(), spellNeed() (+14 more)

### Community 59 - "CompendiumPath"
Cohesion: 0.06
Nodes (63): SpellData, CountHabitatsRow, CountSharedBySourceRow, classBadges(), classBody(), ComparePage(), CompendiumPage(), CompendiumResults() (+55 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - ".combatantHP"
Cohesion: 0.29
Nodes (6): Combatant, damageChoice(), damageHP(), healHP(), setHP(), hpChange

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.09
Nodes (38): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+30 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "InputForEntry"
Cohesion: 0.20
Nodes (9): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), bindFiveToolsImages(), InputForEntry(), TestParseFiveTools() (+1 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Input"
Cohesion: 0.05
Nodes (81): Defenses, UI conventions, Description(), Field(), Group(), Label(), Input(), Content() (+73 more)

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
Nodes (38): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+30 more)

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

### Community 83 - "Layout"
Cohesion: 0.06
Nodes (61): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Scripts(), Separator(), campaignLink(), PendingRequests(), PendingGroup (+53 more)

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "Open"
Cohesion: 0.09
Nodes (26): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+18 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "CompendiumKind"
Cohesion: 0.03
Nodes (53): HomebrewOwner, TestHomebrew(), ValidKind(), ftsQuery(), HomebrewInput, SaveHomebrew(), Search(), SeedHabitats() (+45 more)

### Community 96 - "LinkTarget"
Cohesion: 0.14
Nodes (8): LinkTarget, AddBlockLinkParams, AddEntityFieldLinkParams, EntityFieldLink, ListBacklinksParams, ListEntityFieldBacklinksParams, NoteLink, snapshotBody

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (50): ftClass, ftCopyRef, ftDistance, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed (+42 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".compendiumDir"
Cohesion: 0.09
Nodes (14): IsLoopback(), RequireAdmin(), FiveToolsBooks(), ImportSummary(), fiveToolsKey(), sameOriginOnly(), useValidateTag(), sourceInUseMsg() (+6 more)

### Community 103 - "newLiveFixture"
Cohesion: 0.12
Nodes (24): Player(), BenchmarkEncounterState(), BenchmarkSheetState(), Gotchas, internal/web, Notes, Websocket and live views, isRandomID() (+16 more)

### Community 105 - "State"
Cohesion: 0.10
Nodes (13): event, fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch, Topic (+5 more)

### Community 106 - "Character"
Cohesion: 0.04
Nodes (32): dataFor(), Decode(), DecodeData(), ListSubclassesOfRow, Character, Character, NewCharacterFile(), CanCharacter() (+24 more)

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (19): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftMod(), ftName() (+11 more)

### Community 108 - "deviceLabel"
Cohesion: 0.50
Nodes (3): deviceLabel(), rosterHP(), Device

### Community 109 - "Slugify"
Cohesion: 0.07
Nodes (42): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), AbilityKeys(), AddFeat() (+34 more)

### Community 111 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "Client"
Cohesion: 0.15
Nodes (3): ClientInfo, Client, Hub

### Community 116 - "EquipmentOption"
Cohesion: 0.29
Nodes (7): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), TestOptionText(), EquipmentChoice, optionText()

### Community 117 - "NewClient"
Cohesion: 0.30
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "writeBook"
Cohesion: 0.33
Nodes (7): loadSettings(), saveSettings(), TestSettingsFile(), TestStaleBooksBanner(), TestStartUpBrokenSettings(), TestStartUpRedownloads(), writeBook()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.21
Nodes (6): Slot, errorText(), TestErrorText(), modTargetLabel(), slotLabel(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.10
Nodes (13): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+5 more)

### Community 126 - "WeaponKind"
Cohesion: 0.40
Nodes (4): Completions(), TestCompletions(), TestWeaponKind(), WeaponKind()

### Community 127 - "net/url.Values"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 143 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 146 - "rowFigure"
Cohesion: 0.50
Nodes (4): TestRowFigure(), crText(), rowFigure(), textOnly()

### Community 150 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.23
Nodes (10): SwitchPlan, PlanSourceSwitch(), TestPlanSourceSwitch(), classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits() (+2 more)

### Community 151 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 399 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **31 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Pending`, `ClassInfo`, `command.go`, `perm.go`, `database/sql.NullInt64`, `Book`, `.do`, `cert.go`, `.levelUp`, `.combatantHP`, `FiveToolsJob`, `Open`, `CompendiumKind`, `LinkTarget`, `.compendiumDir`, `State`, `Character`, `Slugify`, `viewerCtx`, `Client`, `writeBook`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14147967382893875 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Client` to `compendium_templ.go`, `Pending`, `State`, `NewClient`, `App`?**
  _High betweenness centrality (0.014) - this node is a cross-community bridge._
- **Should `Session` be split into smaller, more focused modules?**
  _Cohesion score 0.0989247311827957 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `database/sql.NullInt64`, `encoding/json.RawMessage`, `SheetView`, `ClassInfo`, `Character`, `command.go`, `Slugify`, `Feature`, `errBodyWriter`, `perm.go`, `testing.T`, `TestRunCommandUsesCurrentSheet`, `.levelUp`, `ParseSheet`, `NameDesc`, `ParseEffects`, `WeaponKind`?**
  _High betweenness centrality (0.012) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.010496671786994368 - nodes in this community are weakly interconnected._