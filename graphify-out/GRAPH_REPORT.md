# Graph Report - gm-session-manager  (2026-11-18)

## Corpus Check
- 343 files · ~617,057 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4508 nodes · 21962 edges · 148 communities (118 shown, 30 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2207 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f3042b75`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Session
- context.Context
- Parse
- createProvider
- Slugify
- SheetView
- combobox.js
- redirect.go
- EffectFieldsURL
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
- WizardView
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- CompendiumKind
- .render
- app.js
- tabbable.js
- useDismiss
- time.Time
- ParseWikidot
- command.js
- dialog.js
- Book
- formErrors
- tooltip.js
- Input
- .do
- mirror
- export/markdown.go
- composite.js
- .equipmentChoices
- Feature
- Open
- nullID
- ImportView
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Button
- golangci-lint Config
- NewClient
- github.com/a-h/templ.Component
- radiogroup.js
- vendorjs/main.go
- database/sql.NullInt64
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Sheet
- use_transition_status.js
- playwright
- release.sh
- NameDesc
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
- Client
- Lint and Release Rules
- Run Script
- select_templ.go
- useTypeahead
- storeSource
- Validator Rule
- Go Module
- Character
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- viewerCtx
- newLiveFixture
- profFocus
- State
- buildReview
- ftBook
- CampaignSources
- Habitat
- GetImageRow
- config.go
- .compendiumDir
- Client
- Topic
- build.sh
- Queries
- looseString
- errBodyWriter
- ParseEffects
- names
- net/url.Values
- Before/after pairs
- Pending
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- Note
- sourcesEnv
- toast_templ.go
- sourceUsedError
- keepawake_windows.go
- ListSharedWithRow
- ftDuration
- ftMod
- spellCounts

## God Nodes (most connected - your core abstractions)
1. `App` - 368 edges
2. `newTestEnv()` - 146 edges
3. `Sheet` - 130 edges
4. `Button()` - 122 edges
5. `CN()` - 113 edges
6. `SheetView` - 78 edges
7. `IfElse()` - 65 edges
8. `Command` - 59 edges
9. `CompendiumKind` - 56 edges
10. `Queries` - 53 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
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

## Communities (148 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (51): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+43 more)

### Community 1 - "Session"
Cohesion: 0.11
Nodes (17): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), TestCredentialsMatch(), TestMiddlewareIdentity() (+9 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (86): keepAwake(), ListImageCharactersRow, CountCompendiumRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Encounter (+78 more)

### Community 3 - "Parse"
Cohesion: 0.06
Nodes (37): Defense, Parse(), fighter(), TestCommands(), TestConIncreaseAddsHPForEarlierLevels(), TestDefenses(), TestHealingFromBelowZeroStartsAtZero(), TestHealingFromZeroResetsDeathSaves() (+29 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Slugify"
Cohesion: 0.07
Nodes (38): Cantrips(), Slugify(), ListSubclassesOfRow, FreeCast(), SwapFeatSpell(), Spell, spellFromRow(), addClassSpells() (+30 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (114): ArmorEffect(), Signed(), Button(), Variant, Content(), TriggerFor(), Content(), Description() (+106 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "EffectFieldsURL"
Cohesion: 0.33
Nodes (4): EffectFieldsURL(), EffectID(), effectQuery(), EffectTarget

### Community 10 - "Import"
Cohesion: 0.09
Nodes (42): HomebrewOwner, Source, sourceFormat, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent() (+34 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.06
Nodes (66): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+58 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (28): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput, HomebrewInputFor() (+20 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (46): ListCharacterNamesRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass(), CanDeleteCharacter(), CanEditBlock() (+38 more)

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

### Community 21 - "WizardView"
Cohesion: 0.08
Nodes (46): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+38 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (39): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+31 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (25): GetBlockRow, CanCharacter(), homebrewTabURL(), entityTopics(), collectHeld(), formInt(), formIntQuery(), redirect() (+17 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (188): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+180 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.15
Nodes (27): ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpPageAtTwenty() (+19 more)

### Community 28 - "CompendiumKind"
Cohesion: 0.04
Nodes (52): internal/compendium, SkillChoice(), dataFor(), Decode(), DecodeData(), SpellData, TextData, UpsertCompendiumEntryParams (+44 more)

### Community 29 - ".render"
Cohesion: 0.06
Nodes (19): ftsQuery(), Search(), ListAttachmentFilesRow, Campaign, compendiumBack(), compendiumPath(), pathCampaign(), viewerCampaign() (+11 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "time.Time"
Cohesion: 0.04
Nodes (38): IsLoopback(), ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), Attachment (+30 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.14
Nodes (22): block, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster(), sourceLine() (+14 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Book"
Cohesion: 0.06
Nodes (49): ClassFeature, rawBook, ClassData, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets() (+41 more)

### Community 38 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Input"
Cohesion: 0.05
Nodes (102): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+94 more)

### Community 41 - ".do"
Cohesion: 0.10
Nodes (23): main(), serve(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+15 more)

### Community 42 - "mirror"
Cohesion: 0.32
Nodes (5): bookWriter, mirror, bookFolder(), downloadBook(), filterDoc()

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".equipmentChoices"
Cohesion: 0.15
Nodes (12): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, ItemData, WornArmor, compendiumWeapon(), TestCompendiumWeapon() (+4 more)

### Community 46 - "Feature"
Cohesion: 0.13
Nodes (12): FeatureEffects, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, EffectMods(), TestValidateEffectMods() (+4 more)

### Community 47 - "Open"
Cohesion: 0.14
Nodes (22): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+14 more)

### Community 48 - "nullID"
Cohesion: 0.16
Nodes (25): nullID(), acBonusForm(), newEffectFix(), notes(), TestConcentrationCascade(), TestEffectFields(), TestEffectModRows(), TestEffectModsReadOnly() (+17 more)

### Community 49 - "ImportView"
Cohesion: 0.16
Nodes (15): ImageCategory, categoryRow(), columnOverview(), ImportReview(), importRows(), importStep(), keepSection(), mappingField() (+7 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (126): TestIsLoopback(), TestDecode(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), RemoveSourceDir(), TestRemoveSourceDir(), TestMasteryLimit() (+118 more)

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
Cohesion: 0.10
Nodes (18): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+10 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Button"
Cohesion: 0.05
Nodes (73): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Fields, Look, Scripts(), Separator(), campaignLink() (+65 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.07
Nodes (29): deny, files, rules, generated, paths, rules, formatters, enable (+21 more)

### Community 58 - "NewClient"
Cohesion: 0.26
Nodes (17): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+9 more)

### Community 59 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (129): SearchCompendiumRow, Field, UI conventions, Action(), Card(), Description(), Footer(), Header() (+121 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 62 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (74): ListHomebrewByOwnerRow, ListHomebrewRow, ListImagesRow, Combatant, Condition, Entity, CreateConditionParams, ListCombatantsRow (+66 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.31
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (36): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+28 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Sheet"
Cohesion: 0.05
Nodes (32): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestExplainSums(), skillByKey(), Sum() (+24 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "NameDesc"
Cohesion: 0.12
Nodes (27): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, AddResistance(), CantripFits() (+19 more)

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
Cohesion: 0.12
Nodes (28): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+20 more)

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

### Community 83 - "Client"
Cohesion: 0.16
Nodes (3): fakeHandler, message, Client

### Community 85 - "Run Script"
Cohesion: 0.80
Nodes (4): changes(), restart_server(), run.sh script, stop_server()

### Community 86 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+16 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "storeSource"
Cohesion: 0.16
Nodes (10): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+2 more)

### Community 96 - "Character"
Cohesion: 0.18
Nodes (7): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), Roster

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (48): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+40 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 103 - "newLiveFixture"
Cohesion: 0.08
Nodes (31): BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas, internal/web, Notes (+23 more)

### Community 105 - "State"
Cohesion: 0.14
Nodes (11): patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff(), deviceLabel() (+3 more)

### Community 106 - "buildReview"
Cohesion: 0.11
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

### Community 107 - "ftBook"
Cohesion: 0.21
Nodes (15): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftParts(), ftRecords() (+7 more)

### Community 108 - "CampaignSources"
Cohesion: 0.14
Nodes (21): toJSON(), campaignSourceRow(), CampaignSources(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), CampaignOverview(), filesGrid() (+13 more)

### Community 109 - "Habitat"
Cohesion: 0.14
Nodes (12): CountHabitatsRow, ListEntryCopiesRow, Habitat, habitatCountMap(), everyHabitat(), ListEntryCopiesParams, MonsterHabitat, RemoveMonsterHabitatParams (+4 more)

### Community 111 - "GetImageRow"
Cohesion: 0.16
Nodes (10): GetImageRow, ensureCategory(), imageItem(), imageVersion(), nullString(), TestImageEditorKeepsOtherChanges(), TestImageFileCache(), csvImport (+2 more)

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - ".compendiumDir"
Cohesion: 0.24
Nodes (4): ImportSummary(), fiveToolsKey(), writeBook(), sourceUsers()

### Community 117 - "Topic"
Cohesion: 0.18
Nodes (6): event, Handler, Topic, TopicKind, heldCollected(), heldTopics

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "Queries"
Cohesion: 0.08
Nodes (35): featureSwitch, Queries, Env, TestSlotMaxima(), ClassFeature, ClassInfo, LevelUpPlan, ClassLevel (+27 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.13
Nodes (9): LevelUpChoice, AddFeat(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), applyASI(), hpGain() (+1 more)

### Community 126 - "names"
Cohesion: 0.40
Nodes (5): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), SourcePreviewView

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
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Note"
Cohesion: 0.20
Nodes (6): Note, NoteKind, CreateNoteParams, ListReadableNotesParams, RenameNoteParams, noteForm

### Community 140 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 141 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 145 - "ftDuration"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 146 - "ftMod"
Cohesion: 0.50
Nodes (3): ftMod(), ftName(), ftReplaceText()

### Community 147 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

## Knowledge Gaps
- **16 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+11 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 401 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Pending`, `Slugify`, `command.go`, `perm.go`, `CompendiumKind`, `.render`, `time.Time`, `Book`, `.do`, `.equipmentChoices`, `cert.go`, `database/sql.NullInt64`, `FiveToolsJob`, `Sheet`, `NameDesc`, `Character`, `viewerCtx`, `State`, `GetImageRow`, `.compendiumDir`, `Client`, `Queries`?**
  _High betweenness centrality (0.055) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _16 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.144318597004019 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Parse`, `encoding/json.RawMessage`, `Slugify`, `SheetView`, `newLiveFixture`, `NameDesc`, `Input`, `command.go`, `.equipmentChoices`, `Feature`, `nullID`, `Queries`, `ParseSheet`, `CompendiumKind`, `ParseEffects`, `database/sql.NullInt64`?**
  _High betweenness centrality (0.013) - this node is a cross-community bridge._
- **Should `Session` be split into smaller, more focused modules?**
  _Cohesion score 0.10846560846560846 - nodes in this community are weakly interconnected._
- **Why does `serveSplit()` connect `cert.go` to `redirect.go`, `testing.T`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009777578743095984 - nodes in this community are weakly interconnected._