# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 248 files · ~497,828 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3599 nodes · 14628 edges · 223 communities (122 shown, 101 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1463 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f25586be`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- .do
- createProvider
- Encounter
- SheetView
- combobox.js
- redirect.go
- testing.T
- github.com/gin-gonic/gin.Context
- select.js
- command.go
- .Entry
- github.com/a-h/templ.Attributes
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Page
- ImageResults
- wikilink.js
- dialog_templ.go
- popover.js
- .serverError
- Project Rules (CLAUDE.md, README)
- ParseSheet
- .fail
- Sheet
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- Campaign
- command.js
- dialog.js
- Input
- golang.org/x/net/html.Node
- tooltip.js
- Store
- .levelUp
- Slugify
- export/markdown.go
- composite.js
- App
- .loadLive
- .Plan
- IfElse
- .whoami
- github.com/a-h/templ.Component
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- levelUpForm
- golangci-lint Config
- Slugify
- monsterBody
- radiogroup.js
- vendorjs Tool
- select_templ.go
- Handoff
- sqlc Config
- .wizardCommand
- Browser check
- Claude Settings
- combobox_templ.go
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
- checkbox.js
- ImportView
- switch.js
- nullID
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- shadcn-templ.go
- Lint and Release Rules
- Run Script
- NameDesc
- useTypeahead
- capReader
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Topic
- Site
- .ListItemsByCategory
- .render
- inputgroup_templ.go
- Entry
- Import
- Client
- field_templ.go
- wikidotClass
- State
- Uses
- radiogroup_templ.go
- dialog_templ.go
- .ListBlockStepsSince
- Hub
- build.sh
- command_templ.go
- alert_templ.go
- errBodyWriter
- Props
- newTestEnv
- Props
- New
- ParseWikidot
- SearchReadableBlocksRow
- Condition
- Pending
- ArmorData
- buildReview
- Campaign
- DBTX
- Character
- empty_templ.go
- HomebrewInputFor
- toast_templ.go
- tooltip_templ.go
- TestLiveUpdatesAndKick
- Entity
- RosterCharacter
- NoteBlock
- CountHabitatsRow
- lanHostname
- .ListCompendiumByKind
- .ListCompendiumBySlugs
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
- ImageBindingRow
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
- .ListEncountersByMonster
- .ListEntityFieldBacklinks
- ListExportBlocksRow
- ListImageAttachmentsRow
- .ListImages
- .ListReadableNotes
- ListStepJSONRow
- ListSubclassesOfRow
- MonsterHabitat
- .RandomCompendiumEntry
- .RemoveHabitatByName
- .RemoveMonsterHabitat
- .RenameCampaign
- .RenameNote
- .SaveBlockSnapshot
- .SearchCompendium
- .SetBlockPosition
- .SetBlockVersion
- .SetCampaignOpenJoin
- SetCharacterCurrentImageParams
- .SetCombatantHP
- SetCombatantInitiativeParams
- .SetCombatantMarker
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
- .UpdateHomebrew
- .UpsertCompendiumEntry
- .UpsertCompendiumSource

## God Nodes (most connected - your core abstractions)
1. `App` - 266 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Sheet` - 90 edges
5. `newTestEnv()` - 87 edges
6. `SheetView` - 52 edges
7. `IfElse()` - 50 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 37 edges

## Surprising Connections (you probably didn't know these)
- `Next up (pick from these)` --references--> `ParseDefenses()`  [INFERRED]
  HANDOFF.md → internal/dnd5e/sheet.go
- `Where things live (creation and level-up)` --references--> `addClassSpells()`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Next up (pick from these)` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go
- `Where things live (creation and level-up)` --references--> `taughtPicked()`  [INFERRED]
  HANDOFF.md → internal/web/views/levelup_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (223 total, 101 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (44): contextKey, looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes, o5eFeature (+36 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps (+3 more)

### Community 3 - ".do"
Cohesion: 0.17
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Encounter"
Cohesion: 0.18
Nodes (5): Combatant, Encounter, Character, formInt(), ListCharacterEncountersParams

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (83): ArmorEffect(), Signed(), Field, Button(), Variant, Card(), Content(), Description() (+75 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (51): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestCoreTraits(), Completions(), fighter(), TestCommands(), TestCompletions() (+43 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.10
Nodes (13): ClearCookie(), SetCookie(), writeCookie(), Entity, coinCommand(), commandText(), entityTopics(), saveFieldLinks() (+5 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.07
Nodes (66): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+58 more)

### Community 13 - ".Entry"
Cohesion: 0.13
Nodes (19): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+11 more)

### Community 14 - "github.com/a-h/templ.Attributes"
Cohesion: 0.09
Nodes (30): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+22 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (58): Defense, ListCombatantsRow, TestDefenses(), Defenses, ParseDefenses(), Entity, Fields, ParseFields() (+50 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Page"
Cohesion: 0.10
Nodes (38): Scripts(), Separator(), Spinner(), campaignLink(), CampaignOverview(), PendingGroup, PendingRequests(), CharacterPage() (+30 more)

### Community 21 - "ImageResults"
Cohesion: 0.09
Nodes (39): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Look, initials(), selectItems(), homebrewLinks(), currentLook() (+31 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.12
Nodes (33): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), conditionList() (+25 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - ".serverError"
Cohesion: 0.12
Nodes (5): GetBlockRow, setAudience(), blockAccessForm, blockAccessLevel, stepsPayload

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.13
Nodes (28): TestUpgradeSheet(), ParseSheet(), TestLevelUpFeats(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestSpellCounts() (+20 more)

### Community 28 - ".fail"
Cohesion: 0.12
Nodes (6): Character, CanCharacter(), uniqueName(), DisplayPage(), nextStep(), wizardPath()

### Community 29 - "Sheet"
Cohesion: 0.07
Nodes (21): DeathSaves, Part, Skill, TestCustomMods(), TestResistances(), skillByKey(), Sum(), TestDualWield() (+13 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "database/sql.NullInt64"
Cohesion: 0.05
Nodes (40): Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, CompendiumSource, Note (+32 more)

### Community 34 - "Campaign"
Cohesion: 0.23
Nodes (7): Campaign, deviceOS(), formIntQuery(), isPhone(), uaPick(), userAgentSummary(), visitorOnly()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Input"
Cohesion: 0.14
Nodes (28): ctxState, Chip(), Chips(), Combobox(), Content(), Input(), Item(), itemLabel() (+20 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.08
Nodes (33): fixture, tokens, migration, isCleared(), newFixture(), serve(), TestMiddlewareIdentity(), TestOnlyHashIsStored() (+25 more)

### Community 41 - ".levelUp"
Cohesion: 0.10
Nodes (33): Where things live (creation and level-up), Options(), ClassFeature, ClassInfo, LevelUpPlan, classFeatures(), classInfo(), errorField() (+25 more)

### Community 42 - "Slugify"
Cohesion: 0.14
Nodes (21): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions() (+13 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "App"
Cohesion: 0.07
Nodes (8): ftsQuery(), Search(), viewerCampaign(), firstLine(), mdLinkText(), joinPath(), JoinQRSrc(), App

### Community 46 - ".loadLive"
Cohesion: 0.12
Nodes (39): Dialog(), Close(), Content(), Description(), Dialog(), Footer(), Header(), Portal() (+31 more)

### Community 47 - ".Plan"
Cohesion: 0.13
Nodes (11): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), applyASI(), hpGain() (+3 more)

### Community 48 - "IfElse"
Cohesion: 0.07
Nodes (64): SearchCompendiumRow, Badge(), Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant() (+56 more)

### Community 49 - ".whoami"
Cohesion: 0.07
Nodes (17): Identity, Role, FromContext(), RequireAdmin(), hashToken(), Session, browserName(), assetETags() (+9 more)

### Community 50 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (84): ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Action(), Footer() (+76 more)

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
Cohesion: 0.11
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "levelUpForm"
Cohesion: 0.17
Nodes (22): Label(), checkField(), abilityForm(), backgroundAbilities(), choiceFieldset(), choiceName(), choiceOptions(), WizardChoice (+14 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Slugify"
Cohesion: 0.15
Nodes (15): Next up (pick from these), Cantrips(), firstTable(), LevelSpells(), TestOptions(), Slugify(), Spell, addResistance() (+7 more)

### Community 59 - "monsterBody"
Cohesion: 0.22
Nodes (19): classBody(), commaPrefixed(), entryBody(), EntryDetails(), fact(), itemBody(), levelsText(), monsterBody() (+11 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "select_templ.go"
Cohesion: 0.10
Nodes (29): Content(), initialValue(), Item(), itemLabel(), Select(), state(), Trigger(), Value() (+21 more)

### Community 63 - "Handoff"
Cohesion: 0.29
Nodes (6): Handoff, Improvements (bigger, deferred), Lessons (read before orchestrating), Notes, Test data in data-test (all safe to delete), Working on the test instance

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - ".wizardCommand"
Cohesion: 0.11
Nodes (19): dataFor(), DecodeData(), Kind, SpellData, wikidotKind(), SizeNamed(), compendiumWeapon(), TestCompendiumWeapon() (+11 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "combobox_templ.go"
Cohesion: 0.11
Nodes (22): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+14 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "builder_templ.go"
Cohesion: 0.06
Nodes (66): CustomMod, CloseFor(), Description(), Field(), Group(), Label(), Input(), Textarea() (+58 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.13
Nodes (18): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+10 more)

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

### Community 80 - ".get"
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "shadcn-templ.go"
Cohesion: 0.11
Nodes (17): Props, initialChecked(), Direction(), flattenClasses(), lastTopLevelColon(), MergeAttributes(), Ptr(), twImportantToV3() (+9 more)

### Community 86 - "NameDesc"
Cohesion: 0.53
Nodes (4): OptionPrerequisite(), NameDesc, isRepeatable(), optionKind

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.18
Nodes (17): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Topic"
Cohesion: 0.14
Nodes (7): event, Handler, stateHandler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 103 - "Site"
Cohesion: 0.14
Nodes (17): Entry, Site, Source, fiveToolsPage(), importSource(), readSource(), storeSource(), upsert() (+9 more)

### Community 105 - ".render"
Cohesion: 0.10
Nodes (6): compendiumBack(), Roster, HomebrewView, LinkItem, SwitcherResults(), who

### Community 106 - "inputgroup_templ.go"
Cohesion: 0.17
Nodes (13): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+5 more)

### Community 107 - "Entry"
Cohesion: 0.12
Nodes (18): fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children(), decode() (+10 more)

### Community 108 - "Import"
Cohesion: 0.18
Nodes (18): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+10 more)

### Community 109 - "Client"
Cohesion: 0.19
Nodes (3): IsLoopback(), TestIsLoopback(), Client

### Community 110 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 111 - "wikidotClass"
Cohesion: 0.19
Nodes (11): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), sourceLine() (+3 more)

### Community 112 - "State"
Cohesion: 0.14
Nodes (10): fakeHandler, message, patchMessage, viewPatch, topicIDs(), Diff(), State, manyKeys() (+2 more)

### Community 113 - "Uses"
Cohesion: 0.40
Nodes (3): FeatureEffects, UsesLabel(), Uses

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "dialog_templ.go"
Cohesion: 0.18
Nodes (12): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+4 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 123 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 126 - "newTestEnv"
Cohesion: 0.08
Nodes (37): TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestConditionExpiresOnTurn(), TestDeleteRestoreImport(), TestRenameDeleteCampaign(), bugForm(), testEnv (+29 more)

### Community 127 - "Props"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 128 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 129 - "ParseWikidot"
Cohesion: 0.31
Nodes (11): InputForEntry(), TestMonsterFormRoundTrip(), ParseWikidot(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSpecies(), TestParseWikidotSubclass() (+3 more)

### Community 131 - "Condition"
Cohesion: 0.18
Nodes (5): Condition, petrified(), CreateConditionParams, FindCharacterConditionParams, FindCombatantConditionParams

### Community 132 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

### Community 134 - "buildReview"
Cohesion: 0.33
Nodes (4): buildReview(), formValues(), csvImport, csvImports

### Community 135 - "Campaign"
Cohesion: 0.28
Nodes (6): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks()

### Community 138 - "empty_templ.go"
Cohesion: 0.29
Nodes (7): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps

### Community 139 - "HomebrewInputFor"
Cohesion: 0.29
Nodes (8): HomebrewInputFor(), actionsText(), bonusText(), speedText(), TableText(), TraitsText(), withoutTable(), TestHomebrewTextRoundTrip()

### Community 140 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 141 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 142 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 144 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 145 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 148 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 335 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **101 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Encounter`, `buildReview`, `Campaign`, `github.com/gin-gonic/gin.Context`, `perm.go`, `.serverError`, `.fail`, `database/sql.NullInt64`, `Campaign`, `.levelUp`, `Slugify`, `.loadLive`, `.whoami`, `LoadOrCreateCert`, `Slugify`, `.wizardCommand`, `Topic`, `.render`, `State`, `Hub`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10817049263264403 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `.wizardCommand`, `Encounter`, `SheetView`, `builder_templ.go`, `testing.T`, `Slugify`, `.levelUp`, `command.go`, `.Plan`, `perm.go`, `NameDesc`, `Slugify`, `ParseSheet`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.031746031746031744 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `.levelUp` to `compendium_templ.go`, `NameDesc`, `.Plan`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._