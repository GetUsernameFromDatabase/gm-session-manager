# Graph Report - gm-session-manager  (2026-10-27)

## Corpus Check
- 241 files · ~463,646 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3422 nodes · 13501 edges · 203 communities (108 shown, 95 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1329 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `17d7d512`
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
- CampaignRoster
- ImportView
- .do
- CN
- createProvider
- 5e Command Apply
- .Entry
- Encounter
- Combobox JS
- perm.go
- Select JS
- Floating UI DOM
- Hover Interaction JS
- CompendiumView
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
- .Router
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
- CustomMod
- encoding/json.RawMessage
- newTestEnv
- errBodyWriter
- Entry
- .startRequest
- .fail
- monsterBody
- Entry
- fromLookup
- command_templ.go
- open5e.go
- wikiText
- card_templ.go
- tooltip_templ.go
- radiogroup_templ.go
- Hub
- alert_templ.go
- NewCommand
- build.sh
- switch_templ.go
- .CreateEncounter
- Props
- Condition
- generateSVG
- Campaign
- who
- Character
- DBTX
- CompendiumSource
- toast_templ.go
- Entity
- NoteBlock
- SearchReadableBlocksRow
- capReader
- .ListBlockStepsSince
- .ListCompendiumByKind
- lanHostname
- .ListCompendiumBySlugs
- .ListSpellsForClass
- keepawake_other.go
- ListImageCharactersRow
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
- `slotLabel()` --references--> `Slot`  [EXTRACTED]
  web/views/live.go → dnd5e/sheet.go
- `coinList()` --references--> `Coins`  [EXTRACTED]
  web/views/builder_templ.go → dnd5e/sheet.go
- `TestCommandPermissions()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (203 total, 95 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.10
Nodes (51): TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), Store, Command (+43 more)

### Community 1 - "combobox_templ.go"
Cohesion: 0.10
Nodes (28): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+20 more)

### Community 3 - "IfElse"
Cohesion: 0.09
Nodes (45): Look, ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps (+37 more)

### Community 4 - "github.com/a-h/templ.Attributes"
Cohesion: 0.08
Nodes (35): Props, ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant (+27 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.06
Nodes (35): Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, Entity, CreateCombatantParams (+27 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.08
Nodes (73): ArmorEffect(), Signed(), Field, actionEntry, edit, subtitlePart, Action(), Card() (+65 more)

### Community 7 - "CampaignRoster"
Cohesion: 0.33
Nodes (9): Device, RosterCharacter, CampaignOverview(), CampaignRoster(), Roster, rosterCard(), JoinInfo, deviceLabel() (+1 more)

### Community 8 - "ImportView"
Cohesion: 0.19
Nodes (13): ImageCategory, categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows() (+5 more)

### Community 9 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+13 more)

### Community 10 - "CN"
Cohesion: 0.04
Nodes (124): ListMonstersByNameRow, Align, ContentProps, ctxKey, ctxState, Props, Size, TriggerProps (+116 more)

### Community 11 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.09
Nodes (34): Amount, condCommand(), conditionSlug(), Damage(), Heal(), init(), parseAmount(), parseCoin() (+26 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (30): addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, NameDesc, ValidKind(), HomebrewInput (+22 more)

### Community 14 - "Encounter"
Cohesion: 0.18
Nodes (6): Combatant, Encounter, Character, ListCharacterEncountersParams, uniqueName(), formInt()

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (49): GetBlockRow, ListCombatantsRow, blockCase, Combatant, Fields, ParseFields(), VisibleEntity(), VisibleLooks() (+41 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "CompendiumView"
Cohesion: 0.08
Nodes (27): CountCompendiumRow, CountHabitatsRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, SearchCompendiumRow, SearchCompendiumParams, PickerTarget (+19 more)

### Community 21 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+3 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.09
Nodes (45): combatantGroup, CombatantView, CombatView, Badge(), Header(), Title(), addCombatantForm(), addCombatants() (+37 more)

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
Cohesion: 0.10
Nodes (12): ClearCookie(), SetCookie(), writeCookie(), FromContext(), browserName(), reqEntry, deviceKey(), keepsBody() (+4 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.10
Nodes (11): DeathSaves, skillByKey(), Sum(), TestUpgradeSheet(), Coins, Sheet, Slot, Uses (+3 more)

### Community 29 - "Page"
Cohesion: 0.08
Nodes (57): LevelUpPlan, FeaturePreview, Option, PendingView, WizardStep, WS, Label(), Scripts() (+49 more)

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
Nodes (6): TestModifier(), Part, damageLabel(), Weapon, Modifier(), attackTargets()

### Community 34 - "Session"
Cohesion: 0.13
Nodes (9): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), hashToken(), newToken(), tokens, Session (+1 more)

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.19
Nodes (16): block, withCoreTraits(), featuresUnder(), ParseWikidot(), sourceLine(), tableText(), roundTrip(), TestParseWikidotClass() (+8 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.20
Nodes (22): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+14 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "Open"
Cohesion: 0.10
Nodes (28): newFixture(), TestOnlyHashIsStored(), NewManager(), main(), run(), apply(), loadMigrations(), Migrate() (+20 more)

### Community 41 - "shadcn-templ.go"
Cohesion: 0.10
Nodes (16): Props, GroupProps, Props, Orientation, Props, directionKey, Direction(), flattenClasses() (+8 more)

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+12 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.09
Nodes (7): entityTopics(), formIntQuery(), isHTMX(), redirect(), csvSource(), securityHeaders(), DisplayPage()

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.15
Nodes (10): ClassLevel, TestSlotMaxima(), applyASI(), ClassInfo, hpGain(), raise(), LevelUpChoice, CasterLevel() (+2 more)

### Community 48 - "Parse"
Cohesion: 0.12
Nodes (19): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+11 more)

### Community 49 - ".Router"
Cohesion: 0.22
Nodes (6): TestRequireAdmin(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.18
Nodes (16): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), armorCondition(), columnInts() (+8 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.13
Nodes (12): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+4 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - ".fail"
Cohesion: 0.06
Nodes (17): DecodeData(), TextData, Slugify(), ftsQuery(), Search(), SizeNamed(), App, compendiumWeapon() (+9 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "ImportView"
Cohesion: 0.11
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

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
Cohesion: 0.22
Nodes (7): findItem(), parseArmor(), parseItem(), parseWeapon(), TestGearCommands(), ensureItem(), Item

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
Cohesion: 0.08
Nodes (56): ItemData, builderOp, builderVerb, coin, profFocus, CloseFor(), Field(), Label() (+48 more)

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
Cohesion: 0.18
Nodes (15): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestSearch() (+7 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "Button"
Cohesion: 0.07
Nodes (67): ListEntityFieldBacklinksRow, ListBacklinksRow, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props (+59 more)

### Community 99 - "CustomMod"
Cohesion: 0.13
Nodes (11): findMod(), newMod(), parseFeature(), parseFeatureSet(), parseMod(), parseUses(), UsesLabel(), FeatureEffects (+3 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.21
Nodes (16): ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags(), ftEntry (+8 more)

### Community 101 - "newTestEnv"
Cohesion: 0.06
Nodes (69): TestRealSource(), Import(), Verbs(), TestHelpVerbs(), TestCanCharacter(), nullID(), TestBuilderCompendium(), TestBuilderInventory() (+61 more)

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errBodyWriter, errorText(), TestErrorText()

### Community 103 - "Entry"
Cohesion: 0.11
Nodes (19): TestHomebrew(), dataFor(), Kind, fiveToolsPage(), HomebrewOwner, SaveHomebrew(), storeSource(), upsert() (+11 more)

### Community 104 - ".startRequest"
Cohesion: 0.13
Nodes (11): Campaign, Note, PendingKind, CreateNoteParams, blockAccessForm, deviceOS(), isPhone(), uaPick() (+3 more)

### Community 105 - ".fail"
Cohesion: 0.11
Nodes (7): Character, ClassFeature, ParseSheet(), CanCharacter(), classFeatures(), classInfo(), Flash()

### Community 106 - "monsterBody"
Cohesion: 0.11
Nodes (32): SpellData, alignmentColor(), classBadges(), classBody(), commaPrefixed(), compendiumRow(), entryBody(), EntryDetails() (+24 more)

### Community 107 - "Entry"
Cohesion: 0.24
Nodes (10): ArmorData, ClassFeature, Entry, ClassData, ItemData, normalizer, armorText(), children() (+2 more)

### Community 108 - "fromLookup"
Cohesion: 0.22
Nodes (10): Config, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+2 more)

### Community 109 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 110 - "open5e.go"
Cohesion: 0.08
Nodes (30): TestUnflattenTable(), fixture, fixtures, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass (+22 more)

### Community 111 - "wikiText"
Cohesion: 0.27
Nodes (8): WikiLink, WikiPart, wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks()

### Community 112 - "card_templ.go"
Cohesion: 0.25
Nodes (8): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps

### Community 113 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): ContentProps, ctxKey, ctxState, Props, Side, DisableTrigger(), state()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): ctxKey, ctxState, ItemProps, Props, initialValue(), state()

### Community 115 - "Hub"
Cohesion: 0.05
Nodes (36): IsLoopback(), TestIsLoopback(), ClientInfo, event, fakeHandler, Handler, Client, Hub (+28 more)

### Community 116 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 117 - "NewCommand"
Cohesion: 0.15
Nodes (11): EquipmentOption, EquipmentItem, CondOp, Effects, CondOpKind, Env, Spell, WornArmor (+3 more)

### Community 118 - "build.sh"
Cohesion: 0.67
Nodes (3): build(), CGO_ENABLED, build.sh script

### Community 119 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): Props, Size, initialChecked()

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
Cohesion: 0.22
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

### Community 125 - "who"
Cohesion: 0.15
Nodes (5): compendiumBack(), HomebrewView, LinkItem, SwitcherResults(), who

### Community 129 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Props, ToasterProps, Type, Toast(), Toaster(), toastIcon(), toastListItem()

### Community 133 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 141 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **14 isolated node(s):** `CGO_ENABLED`, `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 328 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **95 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `.fail` to `Auth and Compendium Core`, `database/sql.NullInt64`, `.startRequest`, `.fail`, `Open`, `.do`, `github.com/gin-gonic/gin.Context`, `Encounter`, `perm.go`, `.Router`, `Import`, `Hub`, `FromContext`, `who`?**
  _High betweenness centrality (0.045) - this node is a cross-community bridge._
- **What connects `CGO_ENABLED`, `gm-session-manager`, `testEnv` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.10144350423096067 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `Weapon`, `CustomMod`, `Character Sheet View`, `builderFields`, `.fail`, `5e Command Apply`, `fromLookup`, `Encounter`, `5e Level-Up and Slots`, `Parse`, `perm.go`, `wizardChoice`, `NewCommand`, `.fail`, `Page`, `Item`?**
  _High betweenness centrality (0.028) - this node is a cross-community bridge._
- **Should `combobox_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09659090909090909 - nodes in this community are weakly interconnected._
- **Why does `reqLog` connect `Auth and Compendium Core` to `database/sql.NullInt64`, `errBodyWriter`, `newTestEnv`, `builderFields`, `capReader`, `Hub`, `.fail`, `FromContext`?**
  _High betweenness centrality (0.028) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.03333333333333333 - nodes in this community are weakly interconnected._