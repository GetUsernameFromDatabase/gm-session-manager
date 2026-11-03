# Graph Report - gm-session-manager  (2026-11-03)

## Corpus Check
- 246 files · ~479,351 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3507 nodes · 13947 edges · 216 communities (115 shown, 101 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1378 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `dfa75917`
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
- .idParam
- select.js
- command.go
- .Entry
- empty_templ.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- .serverError
- .fail
- wikilink.js
- combat_templ.go
- popover.js
- Campaign
- Project Rules (CLAUDE.md, README)
- CompendiumSource
- Sheet
- Button
- app.js
- tabbable.js
- useDismiss
- App
- github.com/a-h/templ.Attributes
- command.js
- dialog.js
- Weapon
- wikidot.go
- tooltip.js
- Store
- ParseEffects
- .wizardCommand
- export/markdown.go
- composite.js
- .sendFeedback
- github.com/a-h/templ.Component
- .Plan
- monsterBody
- Session
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- toast_templ.go
- golangci-lint Config
- Parse
- combobox_templ.go
- radiogroup.js
- vendorjs Tool
- IfElse
- .ListBlockStepsSince
- sqlc Config
- newStore
- Browser check
- Claude Settings
- .ListCompendiumByKind
- use_transition_status.js
- playwright
- release.sh
- Input
- checkbox.js
- CustomMod
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- NewSheet
- useFocus
- mark_others.js
- .ListCompendiumBySlugs
- Lint and Release Rules
- Run Script
- .ListSpellsForClass
- useTypeahead
- .loadLive
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- popover_templ.go
- Encounter
- shadcn-templ.go
- Entry
- config.go
- inputgroup_templ.go
- select_templ.go
- roundTrip
- Client
- ImportView
- radiogroup_templ.go
- field_templ.go
- command_templ.go
- dialog_templ.go
- build.sh
- ParseAidedd
- Condition
- DBTX
- github.com/gin-gonic/gin.Context
- Handoff: open items from the 2026-11-03 review
- Character
- alert_templ.go
- .Router
- Note
- Entity
- Props
- ArmorData
- SearchReadableBlocksRow
- BuilderPicks
- NoteBlock
- generateSVG
- capReader
- switch_templ.go
- keepawake_windows.go
- CountHabitatsRow
- lanHostname
- spellCounts
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
- CreateSessionParams
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
1. `App` - 256 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `newTestEnv()` - 79 edges
5. `Sheet` - 69 edges
6. `SheetView` - 50 edges
7. `IfElse()` - 48 edges
8. `Input()` - 46 edges
9. `Command` - 43 edges
10. `Client` - 35 edges

## Surprising Connections (you probably didn't know these)
- `Next up (pick from these)` --references--> `Options()`  [INFERRED]
  HANDOFF.md → internal/compendium/coretraits.go
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `NewManager()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/auth/manager.go
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (216 total, 101 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (45): contextKey, looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes, o5eFeature (+37 more)

### Community 1 - "levelUpForm"
Cohesion: 0.17
Nodes (24): Label(), abilitySelect(), levelUpForm(), optionItems(), slotSummary(), spellChecklist(), checkField(), abilityForm() (+16 more)

### Community 3 - ".do"
Cohesion: 0.12
Nodes (19): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+11 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.06
Nodes (30): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, Encounter, CreateCombatantParams, uniqueName() (+22 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (81): ArmorEffect(), Slot, Signed(), Field, Action(), Card(), Content(), Description() (+73 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "newTestEnv"
Cohesion: 0.05
Nodes (72): Import(), TestCanCharacter(), nullID(), TestBuilderCompendium(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName(), TestLanguagesForm() (+64 more)

### Community 10 - ".idParam"
Cohesion: 0.12
Nodes (5): isHTMX(), DisplayPage(), LinkItem, SwitcherResults(), who

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.09
Nodes (49): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+41 more)

### Community 13 - ".Entry"
Cohesion: 0.11
Nodes (22): HomebrewOwner, MonsterAction, ValidKind(), HomebrewInput, HomebrewInputFor(), InputForEntry(), SaveHomebrew(), actionsText() (+14 more)

### Community 14 - "empty_templ.go"
Cohesion: 0.29
Nodes (7): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (52): ListCombatantsRow, Entity, Fields, Look, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+44 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - ".serverError"
Cohesion: 0.11
Nodes (5): GetBlockRow, setAudience(), blockAccessForm, blockAccessLevel, stepsPayload

### Community 21 - ".fail"
Cohesion: 0.17
Nodes (3): Character, CanCharacter(), Flash()

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "combat_templ.go"
Cohesion: 0.10
Nodes (39): Badge(), Content(), Header(), Popover(), Title(), beingBadge(), combatantChip(), combatantName() (+31 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "Campaign"
Cohesion: 0.20
Nodes (8): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks(), parseCharacterFile()

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 28 - "Sheet"
Cohesion: 0.07
Nodes (20): DeathSaves, TestCustomMods(), TestExplainSums(), Sum(), TestUpgradeSheet(), Coins, Sheet, WornArmor (+12 more)

### Community 29 - "Button"
Cohesion: 0.06
Nodes (78): ImageCategory, ListBacklinksRow, Button(), Scripts(), Separator(), campaignLink(), CampaignOverview(), CampaignRoster() (+70 more)

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
Nodes (5): compendiumBack(), viewerCampaign(), redirect(), HomebrewView, App

### Community 34 - "github.com/a-h/templ.Attributes"
Cohesion: 0.09
Nodes (30): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+22 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Weapon"
Cohesion: 0.15
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 38 - "wikidot.go"
Cohesion: 0.14
Nodes (30): block, ClassFeature, textLines(), ClassData, ParsePlainPage(), attr(), byID(), cellAt() (+22 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.12
Nodes (25): fixture, migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate() (+17 more)

### Community 41 - "ParseEffects"
Cohesion: 0.19
Nodes (15): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+7 more)

### Community 42 - ".wizardCommand"
Cohesion: 0.09
Nodes (37): Cantrips(), firstTable(), LevelSpells(), OptionLevel(), Options(), TestOptions(), DecodeData(), NameDesc (+29 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.10
Nodes (23): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+15 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".sendFeedback"
Cohesion: 0.33
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 46 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (103): ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer(), Checkbox() (+95 more)

### Community 47 - ".Plan"
Cohesion: 0.12
Nodes (14): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+6 more)

### Community 48 - "monsterBody"
Cohesion: 0.13
Nodes (28): SpellData, alignmentColor(), classBadges(), classBody(), commaPrefixed(), compendiumRow(), crText(), entryBody() (+20 more)

### Community 49 - "Session"
Cohesion: 0.12
Nodes (14): tokens, constantTimeEqual(), credentialsMatch(), isCleared(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity(), TestRequireAdmin() (+6 more)

### Community 50 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

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
Cohesion: 0.14
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Parse"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 59 - "combobox_templ.go"
Cohesion: 0.09
Nodes (23): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+15 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "IfElse"
Cohesion: 0.08
Nodes (52): SearchCompendiumRow, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+44 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "newStore"
Cohesion: 0.14
Nodes (19): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+11 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

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
Nodes (74): Combobox(), Content(), Item(), List(), CloseFor(), Description(), Field(), Group() (+66 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "CustomMod"
Cohesion: 0.16
Nodes (13): FeatureEffects, findMod(), newMod(), parseMod(), SaveFeature(), UsesLabel(), CustomMod, Feature (+5 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.15
Nodes (14): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), TestPlayerCombatantLink(), ImageSrc(), LiveKey() (+6 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "NewSheet"
Cohesion: 0.10
Nodes (21): Completions(), Heal(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestHealingFromZeroResetsDeathSaves() (+13 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".loadLive"
Cohesion: 0.08
Nodes (38): ListEntityFieldBacklinksRow, Spinner(), notFoundIsNil(), encounterLinks(), EncounterList(), initials(), PlayerCombat(), EntityDetail() (+30 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+7 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 103 - "Site"
Cohesion: 0.12
Nodes (16): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+8 more)

### Community 104 - "popover_templ.go"
Cohesion: 0.22
Nodes (12): state(), Trigger(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 105 - "Encounter"
Cohesion: 0.11
Nodes (14): Queries, Combatant, Entity, Character, applyCondOps(), entityTopics(), linkEntityLook(), saveFieldLinks() (+6 more)

### Community 106 - "shadcn-templ.go"
Cohesion: 0.12
Nodes (14): Props, Direction(), flattenClasses(), lastTopLevelColon(), MergeAttributes(), Ptr(), twImportantToV3(), twImportantToV4() (+6 more)

### Community 107 - "Entry"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children() (+12 more)

### Community 108 - "config.go"
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 109 - "inputgroup_templ.go"
Cohesion: 0.16
Nodes (14): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+6 more)

### Community 110 - "select_templ.go"
Cohesion: 0.19
Nodes (14): Content(), initialValue(), Item(), itemLabel(), state(), Trigger(), Value(), Align (+6 more)

### Community 111 - "roundTrip"
Cohesion: 0.32
Nodes (8): CoreTraits(), TestCoreTraits(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSubclass(), TestParseWikidotUnknownKind(), wikiPage()

### Community 112 - "Client"
Cohesion: 0.05
Nodes (37): ClientInfo, event, fakeHandler, Handler, message, patchMessage, PendingState, stateHandler (+29 more)

### Community 113 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows(), importStatusItems() (+2 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 116 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 117 - "dialog_templ.go"
Cohesion: 0.18
Nodes (11): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+3 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "ParseAidedd"
Cohesion: 0.22
Nodes (11): addEntry(), ParseAidedd(), setStat(), statLabel(), MonsterData, firstInt(), abilityKey(), parseBonuses() (+3 more)

### Community 123 - "Condition"
Cohesion: 0.18
Nodes (5): Condition, conditionList(), CreateConditionParams, FindCharacterConditionParams, FindCombatantConditionParams

### Community 125 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (26): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Campaign, coinCommand(), commandText() (+18 more)

### Community 126 - "Handoff: open items from the 2026-11-03 review"
Cohesion: 0.22
Nodes (8): Bugs fixed since, Character page, Handoff: open items from the 2026-11-03 review, Improvements (bigger, deferred), Next up (pick from these), Notes, Other pages, Test data in data-test

### Community 128 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 129 - ".Router"
Cohesion: 0.33
Nodes (3): RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 130 - "Note"
Cohesion: 0.29
Nodes (3): Note, CreateNoteParams, ListReadableNotesParams

### Community 132 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

### Community 135 - "BuilderPicks"
Cohesion: 0.60
Nodes (5): BuilderPickItems(), BuilderPicks(), BuilderPick, pickURL(), TestPickRowActs()

### Community 136 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 137 - "generateSVG"
Cohesion: 0.67
Nodes (4): attrString(), generateSVG(), getIconContent(), Icon()

### Community 139 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

### Community 143 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 144 - "spellCounts"
Cohesion: 0.67
Nodes (3): sortedColumns(), spellCounts(), TestSpellCounts()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **101 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.Router`, `newStore`, `.do`, `database/sql.NullInt64`, `Encounter`, `.idParam`, `.wizardCommand`, `export/markdown.go`, `Client`, `perm.go`, `.serverError`, `.fail`, `LoadOrCreateCert`, `.loadLive`, `Campaign`, `github.com/gin-gonic/gin.Context`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09870500972762646 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `App`, `Weapon`, `SheetView`, `Encounter`, `CustomMod`, `.wizardCommand`, `ParseEffects`, `.Plan`, `NewSheet`, `Parse`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.03508771929824561 - nodes in this community are weakly interconnected._
- **Why does `TestCustomMods()` connect `Sheet` to `compendium_templ.go`, `newTestEnv`, `NewSheet`, `Weapon`?**
  _High betweenness centrality (0.007) - this node is a cross-community bridge._
- **Should `.do` be split into smaller, more focused modules?**
  _Cohesion score 0.12183908045977011 - nodes in this community are weakly interconnected._