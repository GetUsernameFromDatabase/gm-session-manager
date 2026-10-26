# Graph Report - gm-session-manager  (2026-10-27)

## Corpus Check
- 240 files · ~463,411 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3417 nodes · 13496 edges · 211 communities (109 shown, 102 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1329 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `7db68ea8`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- combobox_templ.go
- context.Context
- IfElse
- github.com/a-h/templ.Attributes
- database/sql.NullInt64
- Character Sheet View
- Button
- select_templ.go
- .do
- CN
- createProvider
- 5e Command Apply
- .Entry
- Queries
- Combobox JS
- perm.go
- Select JS
- Floating UI DOM
- Hover Interaction JS
- Badge
- popover_templ.go
- Notes Editor JS
- CombatantView
- Anchor Positioning JS
- radiogroup.js
- Project Rules (CLAUDE.md, README)
- FromContext
- 5e Sheet Model
- Page
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- Weapon
- Session
- App JS Command Bar
- Dismiss Interaction JS
- ParseWikidot
- ParseAidedd
- Command Palette JS
- Open
- shadcn-templ.go
- Dialog JS
- export/markdown.go
- Tooltip JS
- github.com/gin-gonic/gin.Context
- .ListReadableNotes
- 5e Level-Up and Slots
- Parse
- newFixture
- Composite List Navigation
- wizardChoice
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- .fail
- golangci-lint Config
- ImportView
- .sendFeedback
- List Navigation JS
- vendorjs Tool
- Item
- dialog_templ.go
- sqlc Config
- Transition Status JS
- Checkbox JS
- Claude Settings
- Switch JS
- Tabs JS
- Sed Hook
- inputgroup_templ.go
- builderFields
- Portal JS
- Click Interaction JS
- Focus Interaction JS
- Player
- Mark Others (aria-hidden)
- Typeahead JS
- Live View Bindings JS
- Import
- Internal Backdrop JS
- Lifecycle JS
- Button
- Lint and Release Rules
- Run Script
- d20 Favicon
- Validator Rule
- Go Module
- Parse
- encoding/json.RawMessage
- newTestEnv
- errBodyWriter
- Entry
- .startRequest
- roundTrip
- monsterBody
- NewClient
- fromLookup
- command_templ.go
- open5e.go
- wikiText
- field_templ.go
- redirect.go
- radiogroup_templ.go
- State
- alert_templ.go
- NewCommand
- Client
- CountHabitatsRow
- Encounter
- Props
- Condition
- generateSVG
- Campaign
- .render
- Character
- DBTX
- .publishLocked
- toast_templ.go
- Entity
- rowFigure
- NoteBlock
- ListHomebrewRow
- .SearchCompendium
- SearchReadableBlocksRow
- capReader
- .ListBlockStepsSince
- .ListCompendiumByKind
- lanHostname
- .ListCompendiumBySlugs
- .ListSpellsForClass
- keepawake_other.go
- ListImageCharactersRow
- CountCompendiumRow
- ListHomebrewByOwnerRow
- ListHomebrewLinksRow
- UnlinkEntityImageParams
- ListImageEntitiesRow
- AddBlockAudienceParams
- AddBlockLinkParams
- AddEntityFieldLinkParams
- AddHabitatByNameParams
- AddMonsterHabitatParams
- BindMonsterImagesParams
- .CreateCampaign
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
- ListStepJSONRow
- ListSubclassesOfRow
- MonsterHabitat
- .RandomCompendiumEntry
- .RemoveHabitatByName
- .RemoveMonsterHabitat
- .RenameCampaign
- .RenameNote
- .SaveBlockSnapshot
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
1. `App` - 246 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Command` - 100 edges
5. `Sheet` - 93 edges
6. `newTestEnv()` - 74 edges
7. `reqLog` - 73 edges
8. `Entry` - 67 edges
9. `Hub` - 58 edges
10. `Manager` - 57 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `NewCommand()` --calls--> `apply()`  [INFERRED]
  dnd5e/command.go → db/db.go
- `coinList()` --references--> `Coins`  [EXTRACTED]
  web/views/builder_templ.go → dnd5e/sheet.go
- `TestCommandPermissions()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go
- `TestServerSideStripping()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (211 total, 102 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.10
Nodes (41): TestMain(), contextKey, Identity, Manager, keepAwake(), Config, newValidator(), Store (+33 more)

### Community 1 - "combobox_templ.go"
Cohesion: 0.12
Nodes (20): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+12 more)

### Community 3 - "IfElse"
Cohesion: 0.12
Nodes (40): ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant (+32 more)

### Community 4 - "github.com/a-h/templ.Attributes"
Cohesion: 0.08
Nodes (36): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+28 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.06
Nodes (33): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, CompendiumSource, Note, CreateCombatantParams (+25 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.07
Nodes (84): ArmorEffect(), Slot, Signed(), Field, ContentProps, ctxKey, ctxState, Props (+76 more)

### Community 7 - "Button"
Cohesion: 0.11
Nodes (48): BlockView, Device, NoteListView, RosterCharacter, Button(), Dialog(), Close(), Content() (+40 more)

### Community 8 - "select_templ.go"
Cohesion: 0.09
Nodes (32): Align, ContentProps, ctxKey, ctxState, Props, Size, TriggerProps, Content() (+24 more)

### Community 9 - ".do"
Cohesion: 0.16
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 10 - "CN"
Cohesion: 0.05
Nodes (97): ctxState, ListMonstersByNameRow, Action(), Description(), Title(), AspectRatio(), Footer(), Checkbox() (+89 more)

### Community 11 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.09
Nodes (46): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Heal(), init() (+38 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (31): addEntry(), setStat(), TestHomebrew(), MonsterData, NameDesc, ValidKind(), HomebrewOwner, HomebrewInput (+23 more)

### Community 14 - "Queries"
Cohesion: 0.10
Nodes (19): Queries, Combatant, ClassFeature, ClassInfo, LevelUpPlan, Character, ParseSheet(), applyCondOps() (+11 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (51): GetBlockRow, ListCombatantsRow, blockCase, Combatant, Fields, ParseFields(), VisibleEntity(), VisibleLooks() (+43 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "Badge"
Cohesion: 0.14
Nodes (25): SearchCompendiumRow, PickerTarget, Badge(), alignmentColor(), classBadges(), commaPrefixed(), CompendiumResults(), compendiumRow() (+17 more)

### Community 21 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+3 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.11
Nodes (35): combatantGroup, CombatantView, addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+27 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "FromContext"
Cohesion: 0.09
Nodes (15): ClearCookie(), SetCookie(), writeCookie(), FromContext(), browserName(), firstLine(), folderName(), mdLinkText() (+7 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (16): DeathSaves, TestModifier(), TestExplainSums(), UsesLabel(), skillByKey(), Sum(), FeatureEffects, TestUpgradeSheet() (+8 more)

### Community 29 - "Page"
Cohesion: 0.07
Nodes (58): CombatView, FeaturePreview, Option, PendingView, WizardStep, WS, Label(), Scripts() (+50 more)

### Community 30 - "Scroll Lock JS"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 31 - "Popover JS"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 32 - "Tabbable Focus JS"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 33 - "Weapon"
Cohesion: 0.18
Nodes (6): Part, damageLabel(), Weapon, ProficiencyAt(), attackTargets(), weaponReach()

### Community 34 - "Session"
Cohesion: 0.22
Nodes (5): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), Session, CreateSessionParams

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.19
Nodes (11): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), sourceLine() (+3 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.17
Nodes (22): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+14 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "Open"
Cohesion: 0.12
Nodes (21): apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock(), TestFTS5Available() (+13 more)

### Community 41 - "shadcn-templ.go"
Cohesion: 0.09
Nodes (19): Props, GroupProps, Props, Orientation, Props, Props, Size, directionKey (+11 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.15
Nodes (17): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+9 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (12): Entity, blockAccessForm, blockAccessLevel, coinCommand(), commandText(), entityTopics(), redirect(), csvSource() (+4 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.17
Nodes (8): ClassLevel, TestSlotMaxima(), applyASI(), hpGain(), raise(), LevelUpChoice, CasterLevel(), SlotMaxima()

### Community 48 - "Parse"
Cohesion: 0.12
Nodes (19): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+11 more)

### Community 49 - "newFixture"
Cohesion: 0.07
Nodes (20): IsLoopback(), isCleared(), newFixture(), serve(), TestIsLoopback(), TestMiddlewareIdentity(), TestOnlyHashIsStored(), TestRequireAdmin() (+12 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.17
Nodes (16): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), armorCondition(), columnInts() (+8 more)

### Community 52 - "Release Workflow"
Cohesion: 0.11
Nodes (18): CGO_ENABLED, GH_TOKEN, jobs, release, name, permissions, contents, tags (+10 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.15
Nodes (15): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), TestCertNames(), TestKeyFileIsPrivate() (+7 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - ".fail"
Cohesion: 0.07
Nodes (12): DecodeData(), TextData, Character, SizeNamed(), App, compendiumWeapon(), textMarkdown(), applySpecies() (+4 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "ImportView"
Cohesion: 0.13
Nodes (21): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+13 more)

### Community 59 - ".sendFeedback"
Cohesion: 0.33
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Item"
Cohesion: 0.23
Nodes (7): findItem(), parseWeapon(), SaveWeapon(), TestGearCommands(), ensureItem(), Item, weaponCommand()

### Community 63 - "dialog_templ.go"
Cohesion: 0.18
Nodes (11): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+3 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Transition Status JS"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 66 - "Checkbox JS"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Switch JS"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 69 - "Tabs JS"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 71 - "inputgroup_templ.go"
Cohesion: 0.16
Nodes (14): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+6 more)

### Community 72 - "builderFields"
Cohesion: 0.06
Nodes (69): Feature, ItemData, builderOp, builderVerb, coin, profFocus, Combobox(), Content() (+61 more)

### Community 73 - "Portal JS"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 74 - "Click Interaction JS"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 75 - "Focus Interaction JS"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 76 - "Player"
Cohesion: 0.16
Nodes (13): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), liveFixture, TestPlayerCombatantLink(), LiveKey() (+5 more)

### Community 77 - "Mark Others (aria-hidden)"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 78 - "Typeahead JS"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 80 - "Import"
Cohesion: 0.21
Nodes (14): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+6 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "Button"
Cohesion: 0.10
Nodes (35): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Look, ImagesView, ImageView, initials(), currentLook() (+27 more)

### Community 99 - "Parse"
Cohesion: 0.33
Nodes (7): findMod(), newMod(), parseMod(), IsModTarget(), featureCommand(), hpCommand(), parseCommand()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.25
Nodes (14): ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags(), ftEntry (+6 more)

### Community 101 - "newTestEnv"
Cohesion: 0.06
Nodes (65): Import(), seedGarbage(), TestCanCharacter(), nullID(), TestBuilderCompendium(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName() (+57 more)

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errBodyWriter, errorText(), TestErrorText()

### Community 103 - "Entry"
Cohesion: 0.12
Nodes (22): ArmorData, Entry, dataFor(), ItemData, Kind, fiveToolsPage(), storeSource(), upsert() (+14 more)

### Community 104 - ".startRequest"
Cohesion: 0.17
Nodes (8): PendingKind, deviceOS(), formIntQuery(), isPhone(), uaPick(), userAgentSummary(), joinPath(), visitorOnly()

### Community 105 - "roundTrip"
Cohesion: 0.32
Nodes (8): CoreTraits(), TestCoreTraits(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSubclass(), TestParseWikidotUnknownKind(), wikiPage()

### Community 106 - "monsterBody"
Cohesion: 0.22
Nodes (19): SpellData, classBody(), entryBody(), EntryDetails(), fact(), itemBody(), levelsText(), monsterBody() (+11 more)

### Community 107 - "NewClient"
Cohesion: 0.28
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 108 - "fromLookup"
Cohesion: 0.29
Nodes (8): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc

### Community 109 - "command_templ.go"
Cohesion: 0.16
Nodes (13): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+5 more)

### Community 110 - "open5e.go"
Cohesion: 0.08
Nodes (34): TestUnflattenTable(), fixture, fixtures, looseString, normalizer, o5eArmor, o5eAssignment, o5eChild (+26 more)

### Community 111 - "wikiText"
Cohesion: 0.27
Nodes (8): WikiLink, WikiPart, wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks()

### Community 112 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 113 - "redirect.go"
Cohesion: 0.24
Nodes (5): connListener, peekedConn, newConnListener(), route(), SplitTLS()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): ctxKey, ctxState, ItemProps, Props, initialValue(), state()

### Community 115 - "State"
Cohesion: 0.10
Nodes (13): event, fakeHandler, Handler, Topic, Diff(), State, manyKeys(), TestDiff() (+5 more)

### Community 116 - "alert_templ.go"
Cohesion: 0.36
Nodes (7): ActionProps, DescriptionProps, Props, TitleProps, Variant, Alert(), variantClasses()

### Community 117 - "NewCommand"
Cohesion: 0.13
Nodes (11): EquipmentOption, EquipmentItem, Effects, NewCommand(), SaveFeature(), Env, Spell, WornArmor (+3 more)

### Community 118 - "Client"
Cohesion: 0.23
Nodes (3): Client, topicIDs(), message

### Community 120 - "Encounter"
Cohesion: 0.22
Nodes (5): Encounter, CreateEncounterParams, ListCharacterEncountersParams, encounterLinks(), PlayerCombat()

### Community 121 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 122 - "Condition"
Cohesion: 0.20
Nodes (4): Condition, CreateConditionParams, FindCharacterConditionParams, FindCombatantConditionParams

### Community 123 - "generateSVG"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 124 - "Campaign"
Cohesion: 0.18
Nodes (8): Slugify(), Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

### Community 125 - ".render"
Cohesion: 0.10
Nodes (11): ftsQuery(), Search(), Campaign, compendiumBack(), viewerCampaign(), isHTMX(), HomebrewView, JoinQRSrc() (+3 more)

### Community 129 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Props, ToasterProps, Type, Toast(), Toaster(), toastIcon(), toastListItem()

### Community 132 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 133 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 141 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 326 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **102 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `.fail` to `Auth and Compendium Core`, `database/sql.NullInt64`, `.startRequest`, `.do`, `github.com/gin-gonic/gin.Context`, `Queries`, `perm.go`, `newFixture`, `State`, `TLS Cert and Main`, `FromContext`, `Campaign`, `.render`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.10046362805276218 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `Weapon`, `Parse`, `Character Sheet View`, `builderFields`, `Queries`, `5e Level-Up and Slots`, `Parse`, `perm.go`, `wizardChoice`, `NewCommand`, `.fail`, `Page`, `Item`?**
  _High betweenness centrality (0.025) - this node is a cross-community bridge._
- **Should `combobox_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11904761904761904 - nodes in this community are weakly interconnected._
- **Why does `reqLog` connect `Auth and Compendium Core` to `database/sql.NullInt64`, `errBodyWriter`, `newTestEnv`, `builderFields`, `capReader`, `Client`, `.fail`, `FromContext`?**
  _High betweenness centrality (0.020) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.03571428571428571 - nodes in this community are weakly interconnected._