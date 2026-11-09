# Graph Report - gm-session-manager  (2026-11-09)

## Corpus Check
- 296 files · ~568,630 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4092 nodes · 17728 edges · 140 communities (114 shown, 26 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1791 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `43394730`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- context.Context
- .do
- createProvider
- .blockFor
- SheetView
- combobox.js
- redirect.go
- Button
- spellCounts
- select.js
- Command
- .Entry
- Page
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- github.com/a-h/templ.Component
- levelUpForm
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
- App
- testing.T
- command.js
- dialog.js
- monsterBody
- select_templ.go
- tooltip.js
- .encounterView
- ParseAidedd
- .render
- cell
- composite.js
- .sendFeedback
- Feature
- ClassInfo
- State
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
- ftDuration
- Lint and Release Rules
- Run Script
- .levelUp
- useTypeahead
- .gearFor
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- plainTags
- mounted
- NewClient
- Site
- PendingRequests
- .whoami
- ftBook
- SearchReadableBlocksRow
- formErrors
- export/markdown.go
- Topic
- TestSourceSwitchKeepsEdits
- Campaign
- Sheet
- NameDesc
- build.sh
- Hub
- errBodyWriter
- Queries
- Before/after pairs
- open5e.go
- Import
- db/CLAUDE.md
- perm/CLAUDE.md
- ftNamed
- DBTX
- CompendiumSource
- NoteBlock
- Badge
- lanHostname
- .ListBlockStepsSince
- .ListCompendiumBySlugs
- .ListCompendiumNames
- subclassPanels

## God Nodes (most connected - your core abstractions)
1. `App` - 303 edges
2. `newTestEnv()` - 121 edges
3. `Sheet` - 118 edges
4. `Button()` - 116 edges
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
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (140 total, 26 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (63): contextKey, keepAwake(), classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild (+55 more)

### Community 1 - "CombatantView"
Cohesion: 0.09
Nodes (38): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+30 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (83): keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListImageEntitiesRow, AddBlockAudienceParams (+75 more)

### Community 3 - ".do"
Cohesion: 0.15
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - ".blockFor"
Cohesion: 0.22
Nodes (3): GetBlockRow, blockAccessLevel, stepsPayload

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (78): ArmorEffect(), Coins, Slot, Signed(), Variant, Action(), Card(), Content() (+70 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Button"
Cohesion: 0.12
Nodes (41): Props, Size, ctxKey, ctxState, Props, Defenses, damageChoice(), Button() (+33 more)

### Community 10 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (70): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+62 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (31): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), TestTokensReadAsLabels() (+23 more)

### Community 14 - "Page"
Cohesion: 0.04
Nodes (103): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListEntityFieldBacklinksRow (+95 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.09
Nodes (44): Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass(), CanCharacter(), CanDeleteCharacter() (+36 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (97): UI conventions, Checkbox(), Description(), Field(), Group(), Label(), Title(), Input() (+89 more)

### Community 21 - "levelUpForm"
Cohesion: 0.08
Nodes (51): Label(), initialValue(), Item(), RadioGroup(), state(), IfElse(), inFightSuffix(), LevelUpView (+43 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.13
Nodes (31): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+23 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "popover_templ.go"
Cohesion: 0.10
Nodes (31): Entity, Look, Content(), Description(), Header(), Popover(), state(), Title() (+23 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.12
Nodes (31): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage() (+23 more)

### Community 28 - "App"
Cohesion: 0.07
Nodes (14): Character, coinCommand(), commandText(), entityTopics(), linkEntityLook(), redirect(), linkCharacterLook(), lookFree() (+6 more)

### Community 29 - "newTestEnv"
Cohesion: 0.05
Nodes (52): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait() (+44 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "App"
Cohesion: 0.08
Nodes (8): ftsQuery(), Search(), compendiumBack(), ImportSummary(), viewerCampaign(), joinPath(), HomebrewView, App

### Community 34 - "testing.T"
Cohesion: 0.03
Nodes (82): Defense, TestCoreTraits(), readBook(), TestFiveToolsCopyTargets(), feature(), fiveToolsClasses(), TestFiveToolsClasses(), RemoveSourceDir() (+74 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "monsterBody"
Cohesion: 0.12
Nodes (30): classBody(), commaPrefixed(), CompendiumFindURL(), entryBody(), EntryDetails(), entryHead(), EntryPage(), EntryPreview() (+22 more)

### Community 38 - "select_templ.go"
Cohesion: 0.11
Nodes (22): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state() (+14 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".encounterView"
Cohesion: 0.10
Nodes (16): Combatant, Encounter, CreateCombatantParams, ListCombatantsRow, Character, combatantDefenses(), initiativeMod(), outOfFight() (+8 more)

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - ".render"
Cohesion: 0.11
Nodes (9): fiveToolsKey(), isHTMX(), imageItem(), imageVersion(), sourceInUseMsg(), sourceUsers(), LinkItem, SwitcherResults() (+1 more)

### Community 43 - "cell"
Cohesion: 0.14
Nodes (19): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+11 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 46 - "Feature"
Cohesion: 0.10
Nodes (22): FeatureEffects, CoreTraits(), EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+14 more)

### Community 47 - "ClassInfo"
Cohesion: 0.12
Nodes (11): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+3 more)

### Community 48 - "State"
Cohesion: 0.15
Nodes (11): message, patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff(), deviceLabel() (+3 more)

### Community 49 - "Character"
Cohesion: 0.22
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

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

### Community 56 - "ClassData"
Cohesion: 0.14
Nodes (23): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "tooltip_templ.go"
Cohesion: 0.10
Nodes (30): Field, Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks() (+22 more)

### Community 59 - "Book"
Cohesion: 0.08
Nodes (37): bookWriter, ftEntry, ftModOp, ftMonster, ftSpell, ftSpellcasting, mirror, rawBook (+29 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "ImportView"
Cohesion: 0.16
Nodes (16): ImageCategory, importCategories(), TestImportCategories(), categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportRow (+8 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.26
Nodes (4): FiveToolsProgress(), ImportNotice, FiveToolsJob, fiveToolsJob

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.05
Nodes (36): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, Entity, Note, ensureCategory() (+28 more)

### Community 66 - "Browser check"
Cohesion: 0.33
Nodes (5): Browser check, Roles, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (174): Props, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size (+166 more)

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
Cohesion: 0.06
Nodes (49): Cantrips(), firstTable(), LevelSpells(), TestOptions(), dataFor(), DecodeData(), Kind, SpellData (+41 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "Session"
Cohesion: 0.07
Nodes (22): tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), serve(), TestCredentialsMatch(), TestIsLoopback() (+14 more)

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
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "ftDuration"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 86 - ".levelUp"
Cohesion: 0.07
Nodes (40): Options(), dnd5eFeatures(), TestFiveToolsClassLevelUp(), CompendiumEntry, TestSlotMaxima(), ClassFeature, ClassInfo, LevelUpPlan (+32 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".gearFor"
Cohesion: 0.18
Nodes (9): EquipmentItem, EquipmentOption, ItemData, WornArmor, compendiumWeapon(), TestCompendiumWeapon(), EquipmentChoice, applyEquipment() (+1 more)

### Community 100 - "plainTags"
Cohesion: 0.15
Nodes (17): FiveToolsImages(), ftCell(), ftFeatIncrease(), ftTable(), TestFiveToolsCell(), TestFtFeatIncrease(), TestStripTagsTokens(), ftClassCell() (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Site"
Cohesion: 0.16
Nodes (10): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+2 more)

### Community 104 - "PendingRequests"
Cohesion: 0.39
Nodes (4): campaignLink(), PendingGroup, PendingRequests(), PendingView

### Community 106 - ".whoami"
Cohesion: 0.07
Nodes (22): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), folderName() (+14 more)

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (23): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+15 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 111 - "export/markdown.go"
Cohesion: 0.19
Nodes (12): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+4 more)

### Community 112 - "Topic"
Cohesion: 0.13
Nodes (6): event, fakeHandler, Handler, stateHandler, TopicKind, Topic

### Community 113 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.33
Nodes (6): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits()

### Community 114 - "Campaign"
Cohesion: 0.11
Nodes (10): PendingState, Campaign, Pending, PendingKind, formIntQuery(), TestUserAgent(), userAgentSummary(), visitorOnly() (+2 more)

### Community 115 - "Sheet"
Cohesion: 0.07
Nodes (14): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield() (+6 more)

### Community 116 - "NameDesc"
Cohesion: 0.53
Nodes (4): OptionPrerequisite(), NameDesc, isRepeatable(), optionKind

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "Hub"
Cohesion: 0.15
Nodes (4): ClientInfo, Client, Hub, randomID()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Queries"
Cohesion: 0.10
Nodes (17): featureSwitch, Queries, featureNames(), SourceRef, SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch() (+9 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (26): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eMastery(), TestUnflattenTable() (+18 more)

### Community 134 - "Import"
Cohesion: 0.06
Nodes (61): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), copySource() (+53 more)

### Community 137 - "ftNamed"
Cohesion: 0.10
Nodes (20): ArmorData, ftClass, ftClassFeature, ftClassRefs, ftCopyRef, ftFeat, ftItem, ftNamed (+12 more)

### Community 141 - "CompendiumSource"
Cohesion: 0.16
Nodes (11): ListEntryCopiesRow, CompendiumSource, pinnedRefs(), sourceIssue(), copyLabel(), SourceCopy, SourceIssue, SourceLabel() (+3 more)

### Community 145 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 148 - "Badge"
Cohesion: 0.09
Nodes (37): Props, Variant, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), alignmentColor(), classBadges() (+29 more)

### Community 149 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 162 - "subclassPanels"
Cohesion: 0.50
Nodes (3): ListSubclassesOfRow, subclassPanels(), ListSubclassesOfParams

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 382 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **26 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.do`, `.blockFor`, `Import`, `CompendiumSource`, `perm.go`, `App`, `.encounterView`, `.render`, `State`, `FiveToolsJob`, `database/sql.NullInt64`, `Slugify`, `Session`, `.levelUp`, `.gearFor`, `PendingRequests`, `.whoami`, `Campaign`, `Hub`, `Queries`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10646733111849391 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `testing.T`, `SheetView`, `.encounterView`, `Button`, `Slugify`, `ParseSheet`, `Command`, `CompendiumSource`, `Feature`, `ClassInfo`, `TestSourceSwitchKeepsEdits`, `NameDesc`, `levelUpForm`, `.levelUp`, `.gearFor`, `Book`, `Queries`?**
  _High betweenness centrality (0.019) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.09276018099547512 - nodes in this community are weakly interconnected._
- **Why does `Slugify()` connect `Slugify` to `compendium_templ.go`, `plainTags`, `open5e.go`, `monsterBody`, `Site`, `ParseAidedd`, `.render`, `ftBook`, `.Entry`, `CompendiumSource`, `ClassInfo`, `Character`, `NameDesc`, `.levelUp`, `ClassData`, `.gearFor`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.0088379705400982 - nodes in this community are weakly interconnected._