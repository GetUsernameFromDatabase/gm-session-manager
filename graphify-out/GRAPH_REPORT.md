# Graph Report - gm-session-manager  (2026-10-26)

## Corpus Check
- 231 files · ~449,916 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 49 file(s) not represented in the graph (top: .templ 41, .css 4, (none) 3)

## Summary
- 3290 nodes · 12912 edges · 212 communities (110 shown, 102 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1295 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `01941cae`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- combobox_templ.go
- context.Context
- .renderCompendium
- github.com/a-h/templ.Attributes
- database/sql.NullInt64
- Character Sheet View
- Button
- State
- .do
- github.com/a-h/templ.Component
- createProvider
- 5e Command Apply
- .Entry
- Encounter
- Combobox JS
- perm.go
- Select JS
- Floating UI DOM
- Hover Interaction JS
- Campaign
- popover_templ.go
- Notes Editor JS
- CombatantView
- Anchor Positioning JS
- IsLoopback
- Project Rules (CLAUDE.md, README)
- tabs_templ.go
- 5e Sheet Model
- Badge
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
- who
- Dialog JS
- cell
- Tooltip JS
- .serverError
- App
- 5e Level-Up and Slots
- Parse
- CountHabitatsRow
- Composite List Navigation
- wizardChoice
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- .fail
- golangci-lint Config
- select_templ.go
- Markdown Export
- List Navigation JS
- vendorjs Tool
- Item
- builder_templ.go
- sqlc Config
- Transition Status JS
- Checkbox JS
- Claude Settings
- Switch JS
- Tabs JS
- Sed Hook
- inputgroup_templ.go
- CommandBuilder
- Portal JS
- Click Interaction JS
- Focus Interaction JS
- Player
- Mark Others (aria-hidden)
- Typeahead JS
- Live View Bindings JS
- testing.T
- Internal Backdrop JS
- Lifecycle JS
- Button
- Lint and Release Rules
- Run Script
- d20 Favicon
- Validator Rule
- Go Module
- Entry
- encoding/json.RawMessage
- .Router
- newStore
- Site
- .websocket
- github.com/gin-gonic/gin.Context
- ImportView
- Note
- fromLookup
- command_templ.go
- open5e.go
- wikiText
- field_templ.go
- Client
- Topic
- NewClient
- alert_templ.go
- toast_templ.go
- CustomMod
- Condition
- card_templ.go
- Props
- RosterCharacter
- generateSVG
- switch_templ.go
- parseLiveKey
- Character
- empty_templ.go
- DBTX
- importSource
- .importUpload
- Entity
- NoteBlock
- .playerGroups
- SearchReadableBlocksRow
- .ListBlockStepsSince
- .ListCompendiumByKind
- .ListCompendiumBySlugs
- .ListSpellsForClass
- lanHostname
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
1. `App` - 237 edges
2. `CN()` - 110 edges
3. `Command` - 100 edges
4. `Button()` - 100 edges
5. `Sheet` - 93 edges
6. `Entry` - 67 edges
7. `newTestEnv()` - 62 edges
8. `Hub` - 58 edges
9. `Manager` - 53 edges
10. `SheetView` - 48 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `NewCommand()` --calls--> `apply()`  [INFERRED]
  dnd5e/command.go → db/db.go
- `TestCommandPermissions()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go
- `TestServerSideStripping()` --calls--> `nullID()`  [INFERRED]
  web/flows_test.go → perm/sql_agreement_test.go
- `newLiveFixture()` --calls--> `nullID()`  [INFERRED]
  web/live_test.go → perm/sql_agreement_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (212 total, 102 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.10
Nodes (46): TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), Config, newValidator() (+38 more)

### Community 1 - "combobox_templ.go"
Cohesion: 0.09
Nodes (23): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+15 more)

### Community 2 - "context.Context"
Cohesion: 0.03
Nodes (3): fetch(), Direction(), WithDirection()

### Community 4 - "github.com/a-h/templ.Attributes"
Cohesion: 0.06
Nodes (36): Props, Props, ContentProps, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+28 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.07
Nodes (31): Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, CompendiumSource, CreateCombatantParams (+23 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.07
Nodes (82): ArmorEffect(), Slot, Signed(), ctxState, actionEntry, edit, profFocus, subtitlePart (+74 more)

### Community 7 - "Button"
Cohesion: 0.08
Nodes (57): NoteListView, PendingView, WS, Content(), Description(), Empty(), Header(), Media() (+49 more)

### Community 8 - "State"
Cohesion: 0.19
Nodes (7): Diff(), State, TestDiff(), patchMessage, stateHandler, viewPatch, liveView

### Community 9 - ".do"
Cohesion: 0.12
Nodes (19): isCleared(), serve(), TestMiddlewareIdentity(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+11 more)

### Community 10 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (105): ListMonstersByNameRow, ctxKey, ctxState, Props, Align, ContentProps, ctxKey, ctxState (+97 more)

### Community 11 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.08
Nodes (50): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+42 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (30): addEntry(), setStat(), MonsterData, NameDesc, ValidKind(), HomebrewOwner, HomebrewInput, HomebrewInputFor() (+22 more)

### Community 14 - "Encounter"
Cohesion: 0.15
Nodes (11): Encounter, ListCombatantsRow, ListCharacterEncountersParams, initiativeMod(), outOfFight(), pcSheet(), stepTurn(), uniqueName() (+3 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "perm.go"
Cohesion: 0.13
Nodes (34): blockCase, Combatant, Look, VisibleEntity(), VisibleLooks(), ImageLook, BlockVisibilities(), CanDeleteCharacter() (+26 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "Campaign"
Cohesion: 0.22
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

### Community 21 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+3 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.13
Nodes (30): CombatantView, CombatView, beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), CombatPage() (+22 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "tabs_templ.go"
Cohesion: 0.09
Nodes (45): ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant (+37 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (12): DeathSaves, skillByKey(), Sum(), TestUpgradeSheet(), Coins, Sheet, IsConditionSlug(), ProficiencyAt() (+4 more)

### Community 29 - "Badge"
Cohesion: 0.07
Nodes (46): SearchCompendiumRow, SearchCompendiumParams, PickerTarget, Badge(), alignmentColor(), classBadges(), classBody(), commaPrefixed() (+38 more)

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
Cohesion: 0.17
Nodes (7): TestModifier(), Part, TestDualWield(), damageLabel(), Weapon, Modifier(), attackTargets()

### Community 34 - "Session"
Cohesion: 0.12
Nodes (10): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Role, hashToken(), newToken(), tokens (+2 more)

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.11
Nodes (26): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, withCoreTraits(), pageSource(), PageSourceKey() (+18 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "Open"
Cohesion: 0.11
Nodes (24): newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open(), openPool() (+16 more)

### Community 41 - "who"
Cohesion: 0.10
Nodes (8): GetBlockRow, blockAccessLevel, compendiumBack(), stepsPayload, HomebrewView, LinkItem, SwitcherResults(), who

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "cell"
Cohesion: 0.13
Nodes (22): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+14 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - ".serverError"
Cohesion: 0.11
Nodes (4): isHTMX(), joinPath(), JoinQRSrc(), Flash()

### Community 46 - "App"
Cohesion: 0.10
Nodes (8): Combatant, Character, Fields, ParseFields(), App, saveFieldLinks(), saveSniffed(), toEntity()

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.13
Nodes (12): ClassLevel, TestSlotMaxima(), applyASI(), ClassInfo, hpGain(), raise(), LevelUpChoice, Spell (+4 more)

### Community 48 - "Parse"
Cohesion: 0.13
Nodes (18): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+10 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.12
Nodes (22): EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), EquipmentItem (+14 more)

### Community 52 - "Release Workflow"
Cohesion: 0.11
Nodes (18): CGO_ENABLED, GH_TOKEN, jobs, release, name, permissions, contents, tags (+10 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.13
Nodes (13): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+5 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - ".fail"
Cohesion: 0.10
Nodes (15): DecodeData(), TextData, Slugify(), Character, ClassFeature, ParseSheet(), SizeNamed(), CanCharacter() (+7 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "select_templ.go"
Cohesion: 0.12
Nodes (31): LevelUpPlan, Field, FeaturePreview, Option, WizardStep, Label(), homebrewAction(), HomebrewPage() (+23 more)

### Community 59 - "Markdown Export"
Cohesion: 0.28
Nodes (9): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+1 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 63 - "builder_templ.go"
Cohesion: 0.08
Nodes (56): ItemData, sourceField, Combobox(), Content(), Input(), Item(), List(), state() (+48 more)

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

### Community 72 - "CommandBuilder"
Cohesion: 0.17
Nodes (14): builderOp, builderVerb, armorForm(), BuilderPickItems(), BuilderPicks(), closeAfter(), CommandBuilder(), compendiumLookup() (+6 more)

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
Nodes (13): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), liveFixture, ImageSrc(), LiveKey() (+5 more)

### Community 77 - "Mark Others (aria-hidden)"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 78 - "Typeahead JS"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 80 - "testing.T"
Cohesion: 0.07
Nodes (60): TestRealSource(), Import(), TestCanCharacter(), nullID(), TestBuilderCompendium(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName() (+52 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "Button"
Cohesion: 0.08
Nodes (52): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, BlockView, ImagesView, ImageView, Button(), IfElse() (+44 more)

### Community 99 - "Entry"
Cohesion: 0.24
Nodes (10): ArmorData, Entry, ItemData, SpellData, fixtures, normalizer, armorText(), byPK() (+2 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.21
Nodes (16): abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags() (+8 more)

### Community 101 - ".Router"
Cohesion: 0.21
Nodes (6): TestRequireAdmin(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 102 - "newStore"
Cohesion: 0.24
Nodes (13): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+5 more)

### Community 103 - "Site"
Cohesion: 0.18
Nodes (9): dataFor(), Kind, fiveToolsPage(), Site, FetchPage(), Page, ParsePageURL(), SiteByKey() (+1 more)

### Community 104 - ".websocket"
Cohesion: 0.50
Nodes (4): deviceOS(), formIntQuery(), uaPick(), userAgentSummary()

### Community 105 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (18): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Entity, PendingKind, entityTopics(), linkEntityLook() (+10 more)

### Community 106 - "ImportView"
Cohesion: 0.20
Nodes (12): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportColumn, ImportValue, ImportView, ImportReview() (+4 more)

### Community 107 - "Note"
Cohesion: 0.21
Nodes (4): Note, CreateNoteParams, blockAccessForm, setAudience()

### Community 108 - "fromLookup"
Cohesion: 0.29
Nodes (8): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc

### Community 109 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 110 - "open5e.go"
Cohesion: 0.08
Nodes (27): TestUnflattenTable(), fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+19 more)

### Community 111 - "wikiText"
Cohesion: 0.31
Nodes (7): WikiLink, WikiPart, wikiLink(), wikiText(), ParseWiki(), TestParseWiki(), WikiLinks()

### Community 112 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 113 - "Client"
Cohesion: 0.21
Nodes (3): Client, topicIDs(), message

### Community 114 - "Topic"
Cohesion: 0.17
Nodes (5): event, fakeHandler, Handler, Topic, TopicKind

### Community 115 - "NewClient"
Cohesion: 0.19
Nodes (12): New(), NewClient(), randomID(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+4 more)

### Community 116 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 117 - "toast_templ.go"
Cohesion: 0.33
Nodes (8): Props, ToasterProps, Type, Toast(), Toaster(), toastIcon(), toastListItem(), Toast()

### Community 118 - "CustomMod"
Cohesion: 0.13
Nodes (17): findMod(), newMod(), parseMod(), SaveFeature(), UsesLabel(), FeatureEffects, CustomMod, Feature (+9 more)

### Community 119 - "Condition"
Cohesion: 0.20
Nodes (4): Condition, CreateConditionParams, FindCharacterConditionParams, FindCombatantConditionParams

### Community 120 - "card_templ.go"
Cohesion: 0.25
Nodes (8): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps

### Community 121 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 122 - "RosterCharacter"
Cohesion: 0.48
Nodes (6): Device, RosterCharacter, Roster, rosterCard(), deviceLabel(), rosterHP()

### Community 123 - "generateSVG"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 125 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): Props, Size, initialChecked()

### Community 128 - "empty_templ.go"
Cohesion: 0.28
Nodes (8): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, mediaClasses()

### Community 130 - "importSource"
Cohesion: 0.36
Nodes (7): importSource(), readSource(), storeSource(), upsert(), parseFixtures(), parseOpen5e(), Source

### Community 131 - ".importUpload"
Cohesion: 0.33
Nodes (3): csvImport, csvImports, csvSource()

### Community 133 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 134 - ".playerGroups"
Cohesion: 0.50
Nodes (3): combatantGroup, nameKey(), TestPlayerGroups()

### Community 141 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 312 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **102 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `.renderCompendium`, `.importUpload`, `database/sql.NullInt64`, `Button`, `State`, `.do`, `Encounter`, `perm.go`, `Campaign`, `who`, `.serverError`, `TLS Cert and Main`, `.fail`, `.Router`, `.websocket`, `github.com/gin-gonic/gin.Context`, `Note`, `parseLiveKey`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.0990352830061391 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `Weapon`, `Character Sheet View`, `App`, `5e Level-Up and Slots`, `Parse`, `Encounter`, `wizardChoice`, `CustomMod`, `.fail`, `select_templ.go`, `Item`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._
- **Should `combobox_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09401709401709402 - nodes in this community are weakly interconnected._
- **Why does `Command` connect `5e Command Apply` to `Auth and Compendium Core`, `Parse`, `perm.go`, `CustomMod`, `.fail`, `Item`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.03225806451612903 - nodes in this community are weakly interconnected._