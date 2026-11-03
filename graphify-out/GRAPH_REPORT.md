# Graph Report - gm-session-manager  (2026-11-04)

## Corpus Check
- 246 files · ~491,207 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3564 nodes · 14288 edges · 206 communities (104 shown, 102 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1423 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `534ef1ee`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- levelUpForm
- context.Context
- .do
- createProvider
- database/sql.NullInt64
- SheetView
- combobox.js
- redirect.go
- newTestEnv
- github.com/gin-gonic/gin.Context
- select.js
- command.go
- .Entry
- .serverError
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Part
- wikiText
- wikilink.js
- CombatantView
- popover.js
- Character
- Project Rules (CLAUDE.md, README)
- State
- .encounterView
- Page
- app.js
- tabbable.js
- useDismiss
- App
- Queries
- command.js
- dialog.js
- Sheet
- golang.org/x/net/html.Node
- tooltip.js
- Store
- NameDesc
- Slugify
- export/markdown.go
- composite.js
- .sendFeedback
- github.com/a-h/templ.Attributes
- .Plan
- github.com/a-h/templ.Component
- Session
- Client
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- Feature
- golangci-lint Config
- testing.T
- .homebrewWriter
- radiogroup.js
- vendorjs Tool
- IfElse
- New
- sqlc Config
- Topic
- Browser check
- Claude Settings
- Pending
- use_transition_status.js
- playwright
- release.sh
- Input
- checkbox.js
- .levelUp
- switch.js
- nullID
- tabs.js
- portal.js
- useClick
- builderForm
- useFocus
- mark_others.js
- TestLiveUpdatesAndKick
- Lint and Release Rules
- Run Script
- cell
- useTypeahead
- .levelUpView
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- message
- .ListCompendiumByKind
- ImageBindingRow
- Entry
- config.go
- .ListEncountersByMonster
- select_templ.go
- ParseWikidot
- Hub
- wizardChoice
- radiogroup_templ.go
- .RemoveHabitatByName
- .RenameCampaign
- dialog_templ.go
- build.sh
- .SetCombatantMarker
- .UpdateHomebrew
- time.Time
- Handoff: open items from the 2026-11-03 review
- Condition
- DBTX
- NoteListView
- lanHostname
- Props
- ArmorData
- Character
- capReader
- Entity
- NoteBlock
- .ListBlockStepsSince
- .ListCompendiumBySlugs
- .ListItemsByCategory
- .ListSpellsForClass
- keepawake_other.go
- ListImageCharactersRow
- CountCompendiumRow
- ListHomebrewByOwnerRow
- ListHomebrewLinksRow
- ListHomebrewRow
- ListImageEntitiesRow
- AddBlockAudienceParams
- AddBlockLinkParams
- AddEntityFieldLinkParams
- AddHabitatByNameParams
- AddMonsterHabitatParams
- BindMonsterImagesParams
- .CreateCampaign
- .CreateEncounter
- CreateLevelUpParams
- EnsureImageCategoryParams
- ExpireTurnConditionsParams
- .GetCompendiumEntryByName
- .GetCompendiumEntryBySlug
- ImageLookAccessRow
- InsertBlockStepParams
- LevelUp
- LinkCharacterImageParams
- LinkEntityImageParams
- LinkEntryCampaignParams
- .ListBacklinks
- ListBlockDocsRow
- ListCharacterSheetsRow
- ListCompendiumKeysBySourceRow
- ListCompendiumNamesRow
- .ListEntityFieldBacklinks
- ListExportBlocksRow
- ListImageAttachmentsRow
- .ListImages
- .ListReadableNotes
- ListStepJSONRow
- ListSubclassesOfRow
- MonsterHabitat
- .RandomCompendiumEntry
- .RemoveMonsterHabitat
- .RenameNote
- .SaveBlockSnapshot
- .SearchCompendium
- .SetBlockPosition
- .SetBlockVersion
- .SetCampaignOpenJoin
- SetCharacterCurrentImageParams
- .SetCombatantHP
- SetCombatantInitiativeParams
- SetCompendiumImageParams
- SetEntityCurrentImageParams
- ToggleCharacterImageRevealedParams
- ToggleEntityImageRevealedParams
- .TrimBlockSteps
- UnlinkCharacterImageParams
- UnlinkEntityImageParams
- .UnlinkEntryCampaign
- .UpdateBlockAccess
- .UpdateCharacter
- .UpdateCharacterSheet
- UpdateEncounterStateParams
- UpdateEntityParams
- .UpsertCompendiumEntry
- .UpsertCompendiumSource

## God Nodes (most connected - your core abstractions)
1. `App` - 265 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Sheet` - 88 edges
5. `newTestEnv()` - 84 edges
6. `SheetView` - 51 edges
7. `IfElse()` - 48 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 36 edges

## Surprising Connections (you probably didn't know these)
- `Where things live (creation and level-up)` --references--> `choiceGroup`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Where things live (creation and level-up)` --references--> `readPicks()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Next up (pick from these)` --references--> `ParseDefenses()`  [INFERRED]
  HANDOFF.md → internal/dnd5e/sheet.go
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Next up (pick from these)` --references--> `castingKind()`  [INFERRED]
  HANDOFF.md → internal/web/views/character_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (206 total, 102 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (46): contextKey, keepAwake(), looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes (+38 more)

### Community 1 - "levelUpForm"
Cohesion: 0.25
Nodes (16): abilityForm(), backgroundAbilities(), choiceName(), WizardChoice, WizardView, skillName(), wizardAbilities(), wizardChoices() (+8 more)

### Community 3 - ".do"
Cohesion: 0.17
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.08
Nodes (30): ListAttachmentFilesRow, GetImageRow, ListImagesRow, CompendiumEntry, CompendiumSource, Entity, ListBacklinksRow, ensureCategory() (+22 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (84): Slot, Signed(), Field, Variant, Action(), Card(), Content(), Description() (+76 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "newTestEnv"
Cohesion: 0.06
Nodes (46): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), TestBuilderCompendium(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm() (+38 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (12): CanCharacter(), coinCommand(), commandText(), entityTopics(), formInt(), redirect(), csvSource(), securityHeaders() (+4 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.07
Nodes (66): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+58 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (25): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), firstInt(), HomebrewInput, HomebrewInputFor() (+17 more)

### Community 14 - ".serverError"
Cohesion: 0.08
Nodes (14): Campaign, GetBlockRow, viewerCampaign(), deviceOS(), formIntQuery(), isHTMX(), isPhone(), uaPick() (+6 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.10
Nodes (39): Note, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+31 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Part"
Cohesion: 0.11
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 21 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.08
Nodes (62): Defenses, Content(), Description(), Header(), state(), Title(), Trigger(), triggerID() (+54 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "Character"
Cohesion: 0.11
Nodes (10): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), Roster, nextStep() (+2 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "State"
Cohesion: 0.15
Nodes (9): fakeHandler, patchMessage, viewPatch, State, manyKeys(), deviceLabel(), rosterHP(), Device (+1 more)

### Community 28 - ".encounterView"
Cohesion: 0.11
Nodes (14): Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), damageChoice(), initiativeMod(), outOfFight() (+6 more)

### Community 29 - "Page"
Cohesion: 0.04
Nodes (105): Props, Size, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props (+97 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "App"
Cohesion: 0.07
Nodes (11): Attachment, compendiumWeapon(), firstLine(), mdLinkText(), saveScreenshot(), parseCSV(), TestGuessMapping(), saveSniffed() (+3 more)

### Community 34 - "Queries"
Cohesion: 0.14
Nodes (14): ftsQuery(), Search(), Queries, CreateCombatantParams, TestCampaignExcludesOwnerBlocks(), applyCondOps(), uniqueName(), linkEntityLook() (+6 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Sheet"
Cohesion: 0.07
Nodes (22): DeathSaves, SpellData, TestCustomMods(), TestResistances(), TestExplainSums(), ArmorEffect(), Sum(), geared() (+14 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.22
Nodes (18): ParseAidedd(), statLabel(), textLines(), TestParseAidedd(), ParsePlainPage(), attr(), byID(), collect() (+10 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.07
Nodes (38): HomebrewOwner, migration, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey() (+30 more)

### Community 41 - "NameDesc"
Cohesion: 0.16
Nodes (23): Where things live (creation and level-up), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, catchUpResistances() (+15 more)

### Community 42 - "Slugify"
Cohesion: 0.22
Nodes (11): Cantrips(), Slugify(), Spell, addProfs(), featList(), freeCast(), lineageAbility(), skillPick() (+3 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".sendFeedback"
Cohesion: 0.33
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 46 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (160): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+152 more)

### Community 47 - ".Plan"
Cohesion: 0.16
Nodes (8): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI(), hpGain(), raise()

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (70): CountHabitatsRow, SearchCompendiumRow, Badge(), RadioGroup(), Toast(), Toaster(), toastIcon(), toastListItem() (+62 more)

### Community 49 - "Session"
Cohesion: 0.10
Nodes (19): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+11 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "LoadOrCreateCert"
Cohesion: 0.19
Nodes (14): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+6 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Feature"
Cohesion: 0.21
Nodes (6): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, usesValue()

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "testing.T"
Cohesion: 0.05
Nodes (45): TestCoreTraits(), Completions(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves(), TestLanguages() (+37 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "IfElse"
Cohesion: 0.10
Nodes (45): Look, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+37 more)

### Community 63 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Topic"
Cohesion: 0.18
Nodes (6): event, stateHandler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

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
Cohesion: 0.06
Nodes (77): Combobox(), Content(), Item(), List(), Description(), Field(), Group(), Label() (+69 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - ".levelUp"
Cohesion: 0.15
Nodes (17): DecodeData(), TextData, ClassFeature, ClassInfo, LevelUpPlan, SizeNamed(), textMarkdown(), classFeatures() (+9 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "nullID"
Cohesion: 0.12
Nodes (20): Player(), nullID(), TestCombatantNamesFollowSource(), TestMonsterPortrait(), testEnv, TestAttachmentAccess(), TestAttachmentToImage(), TestFilesCleanup() (+12 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "builderForm"
Cohesion: 0.08
Nodes (30): Props, ListMonstersByNameRow, Checkbox(), initialChecked(), Cell(), Row(), armorForm(), builderForm() (+22 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 86 - "cell"
Cohesion: 0.13
Nodes (22): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+14 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".levelUpView"
Cohesion: 0.20
Nodes (10): ClassLevel, TestSlotMaxima(), CasterLevel(), SlotMaxima(), LevelUpView, Option, ChoiceGroup, CreationChoices (+2 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.18
Nodes (17): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 103 - "Site"
Cohesion: 0.13
Nodes (19): Entry, Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert() (+11 more)

### Community 104 - "message"
Cohesion: 0.33
Nodes (4): message, topicIDs(), Diff(), TestDiff()

### Community 107 - "Entry"
Cohesion: 0.12
Nodes (18): fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children(), decode() (+10 more)

### Community 108 - "config.go"
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 110 - "select_templ.go"
Cohesion: 0.11
Nodes (23): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state() (+15 more)

### Community 111 - "ParseWikidot"
Cohesion: 0.13
Nodes (21): block, ClassFeature, ClassData, InputForEntry(), TestMonsterFormRoundTrip(), withCoreTraits(), cellAt(), classTable() (+13 more)

### Community 113 - "wizardChoice"
Cohesion: 0.18
Nodes (16): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions() (+8 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.20
Nodes (13): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), state() (+5 more)

### Community 117 - "dialog_templ.go"
Cohesion: 0.25
Nodes (24): ctxKey, ctxState, Props, Dialog(), Close(), CloseFor(), Content(), Description() (+16 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 125 - "time.Time"
Cohesion: 0.08
Nodes (16): IsLoopback(), TestIsLoopback(), ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin() (+8 more)

### Community 126 - "Handoff: open items from the 2026-11-03 review"
Cohesion: 0.18
Nodes (10): Defense, Handoff, Improvements (bigger, deferred), Next up (pick from these), Notes, Test data in data-test (all safe to delete), Working on the test instance, ParseDefenses() (+2 more)

### Community 127 - "Condition"
Cohesion: 0.18
Nodes (5): Condition, petrified(), CreateConditionParams, FindCharacterConditionParams, FindCombatantConditionParams

### Community 131 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 132 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

### Community 141 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 334 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **102 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `database/sql.NullInt64`, `github.com/gin-gonic/gin.Context`, `.serverError`, `perm.go`, `Character`, `State`, `.encounterView`, `Page`, `Queries`, `Sheet`, `Store`, `Slugify`, `LoadOrCreateCert`, `.homebrewWriter`, `Topic`, `.levelUp`, `.levelUpView`, `Hub`, `wizardChoice`, `time.Time`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09697059697059697 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `SheetView`, `newTestEnv`, `Slugify`, `NameDesc`, `command.go`, `.levelUp`, `.Plan`, `wizardChoice`, `Part`, `Feature`, `testing.T`, `.levelUpView`, `.encounterView`?**
  _High betweenness centrality (0.023) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.03333333333333333 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `.levelUp` to `compendium_templ.go`, `NameDesc`, `.Plan`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._