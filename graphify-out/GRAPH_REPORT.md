# Graph Report - gm-session-manager  (2026-10-26)

## Corpus Check
- 231 files · ~447,920 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 49 file(s) not represented in the graph (top: .templ 41, .css 4, (none) 3)

## Summary
- 3278 nodes · 12855 edges · 128 communities (106 shown, 22 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1290 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `79609a14`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- combobox_templ.go
- context.Context
- SearchCompendiumRow
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
- shadcn-templ.go
- Combobox JS
- perm.go
- Select JS
- Floating UI DOM
- Hover Interaction JS
- Character
- popover_templ.go
- Notes Editor JS
- CombatantView
- Anchor Positioning JS
- .Serve
- Project Rules (CLAUDE.md, README)
- tabs_templ.go
- 5e Sheet Model
- monsterBody
- Scroll Lock JS
- Popover JS
- Tabbable Focus JS
- Weapon
- newFixture
- App JS Command Bar
- Dismiss Interaction JS
- ParseWikidot
- ParseAidedd
- Command Palette JS
- Open
- .blockFor
- Dialog JS
- cell
- Tooltip JS
- .serverError
- Queries
- 5e Level-Up and Slots
- testing.T
- IfElse
- Composite List Navigation
- wizardChoice
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- Character
- golangci-lint Config
- select_templ.go
- Markdown Export
- List Navigation JS
- vendorjs Tool
- findItem
- builder_templ.go
- sqlc Config
- Transition Status JS
- Checkbox JS
- Claude Settings
- Switch JS
- Tabs JS
- Sed Hook
- inputgroup_templ.go
- combatCondition
- Portal JS
- Click Interaction JS
- Focus Interaction JS
- TestLiveUpdatesAndKick
- Mark Others (aria-hidden)
- Typeahead JS
- Live View Bindings JS
- nullID
- Internal Backdrop JS
- Lifecycle JS
- Input
- Lint and Release Rules
- Run Script
- d20 Favicon
- Validator Rule
- Go Module
- Entry
- encoding/json.RawMessage
- EquipmentOption
- Import
- Site
- Campaign
- App
- ImportView
- dialog_templ.go
- fromLookup
- command_templ.go
- ArmorData
- tooltip_templ.go
- field_templ.go
- Client
- Topic
- NewClient
- alert_templ.go
- toast_templ.go
- CustomMod
- Parse
- .publishLocked
- Props
- RosterCharacter
- generateSVG
- switch_templ.go
- parseLiveKey
- rowFigure

## God Nodes (most connected - your core abstractions)
1. `App` - 236 edges
2. `CN()` - 110 edges
3. `Command` - 100 edges
4. `Button()` - 99 edges
5. `Sheet` - 93 edges
6. `Entry` - 67 edges
7. `newTestEnv()` - 61 edges
8. `Hub` - 58 edges
9. `Manager` - 53 edges
10. `SheetView` - 48 edges

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

## Communities (128 total, 22 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.08
Nodes (68): TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), looseString, o5eAssignment (+60 more)

### Community 1 - "combobox_templ.go"
Cohesion: 0.11
Nodes (22): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+14 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (60): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Encounter, Session (+52 more)

### Community 3 - "SearchCompendiumRow"
Cohesion: 0.20
Nodes (6): ftsQuery(), Search(), SearchCompendiumRow, SearchCompendiumParams, SearchComboItems(), SearchResults()

### Community 4 - "github.com/a-h/templ.Attributes"
Cohesion: 0.07
Nodes (39): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+31 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (42): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, CompendiumSource, Condition, Note (+34 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.08
Nodes (68): ArmorEffect(), Signed(), Field, actionEntry, edit, subtitlePart, Action(), Card() (+60 more)

### Community 7 - "Button"
Cohesion: 0.05
Nodes (87): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, CombatView, ImagesView, ImageView, PendingView, WS (+79 more)

### Community 8 - "State"
Cohesion: 0.16
Nodes (8): fakeHandler, Diff(), State, manyKeys(), TestDiff(), patchMessage, viewPatch, liveView

### Community 9 - ".do"
Cohesion: 0.08
Nodes (37): TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestConditionExpiresOnTurn(), TestDeleteRestoreImport(), TestRenameDeleteCampaign(), testEnv, pendingSecret() (+29 more)

### Community 10 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (98): ctxState, ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer() (+90 more)

### Community 11 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.07
Nodes (51): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+43 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (29): addEntry(), setStat(), TestHomebrew(), MonsterData, NameDesc, ValidKind(), HomebrewOwner, HomebrewInput (+21 more)

### Community 14 - "shadcn-templ.go"
Cohesion: 0.12
Nodes (14): Props, Orientation, Props, directionKey, Direction(), flattenClasses(), lastTopLevelColon(), MergeAttributes() (+6 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (50): ListCombatantsRow, blockCase, Combatant, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), ImageLook (+42 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "Character"
Cohesion: 0.08
Nodes (17): DecodeData(), Slugify(), CompendiumEntry, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile() (+9 more)

### Community 21 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+3 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.12
Nodes (31): combatantGroup, CombatantView, Badge(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName() (+23 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "tabs_templ.go"
Cohesion: 0.20
Nodes (18): ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant (+10 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.09
Nodes (13): DeathSaves, TestExplainSums(), skillByKey(), Sum(), TestGearCommands(), Coins, Item, Sheet (+5 more)

### Community 29 - "monsterBody"
Cohesion: 0.17
Nodes (23): SpellData, alignmentColor(), classBadges(), classBody(), commaPrefixed(), compendiumRow(), entryBody(), EntryDetails() (+15 more)

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

### Community 34 - "newFixture"
Cohesion: 0.07
Nodes (20): constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity(), TestOnlyHashIsStored() (+12 more)

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.12
Nodes (23): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, InputForEntry(), TestMonsterFormRoundTrip(), cellAt() (+15 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "Open"
Cohesion: 0.36
Nodes (8): apply(), loadMigrations(), Migrate(), Open(), openPool(), TestLoadMigrationsRejectsBadNames(), TestMigrateIsIdempotent(), migration

### Community 41 - ".blockFor"
Cohesion: 0.10
Nodes (6): Attachment, GetBlockRow, CreateAttachmentParams, blockAccessLevel, saveSniffed(), stepsPayload

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - ".serverError"
Cohesion: 0.09
Nodes (9): compendiumBack(), formIntQuery(), isHTMX(), joinPath(), HomebrewView, JoinQRSrc(), LinkItem, SwitcherResults() (+1 more)

### Community 46 - "Queries"
Cohesion: 0.12
Nodes (13): Queries, Combatant, Character, applyCondOps(), initiativeMod(), uniqueName(), condOwner, linkEntityLook() (+5 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.14
Nodes (12): ClassLevel, TestSlotMaxima(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan, hpGain(), raise() (+4 more)

### Community 48 - "testing.T"
Cohesion: 0.06
Nodes (42): openTemp(), TestAttachmentOutlivesBlock(), TestFTS5Available(), TestPragmasOnEveryConnection(), TestSessionConstraints(), New(), NewStore(), countCampaigns() (+34 more)

### Community 49 - "IfElse"
Cohesion: 0.15
Nodes (25): CountHabitatsRow, sourceField, List(), Tabs(), Trigger(), IfElse(), profTabs(), characterLooks() (+17 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.13
Nodes (22): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), TextData, armorCondition() (+14 more)

### Community 52 - "Release Workflow"
Cohesion: 0.11
Nodes (18): CGO_ENABLED, GH_TOKEN, jobs, release, name, permissions, contents, tags (+10 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.19
Nodes (14): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+6 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Character"
Cohesion: 0.11
Nodes (7): Character, CanCharacter(), SetCharacterCurrentImageParams, ToggleCharacterImageRevealedParams, UnlinkCharacterImageParams, nextStep(), wizardPath()

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "select_templ.go"
Cohesion: 0.09
Nodes (41): Align, ContentProps, ctxKey, ctxState, Props, Size, TriggerProps, FeaturePreview (+33 more)

### Community 59 - "Markdown Export"
Cohesion: 0.28
Nodes (9): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+1 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.25
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "findItem"
Cohesion: 0.38
Nodes (3): findItem(), parseArmor(), ensureItem()

### Community 63 - "builder_templ.go"
Cohesion: 0.08
Nodes (56): ItemData, builderOp, builderVerb, coin, profFocus, CloseFor(), Description(), Field() (+48 more)

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

### Community 72 - "combatCondition"
Cohesion: 0.16
Nodes (19): Look, Content(), Header(), Popover(), Title(), Trigger(), DirectionAttributes(), whyPopover() (+11 more)

### Community 73 - "Portal JS"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 74 - "Click Interaction JS"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 75 - "Focus Interaction JS"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 76 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 77 - "Mark Others (aria-hidden)"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 78 - "Typeahead JS"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 80 - "nullID"
Cohesion: 0.11
Nodes (22): Player(), nullID(), TestCombatantNamesFollowSource(), TestMonsterPortrait(), testEnv, TestAttachmentAccess(), TestAttachmentToImage(), TestFilesCleanup() (+14 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "Input"
Cohesion: 0.15
Nodes (32): BlockView, NoteListView, Dialog(), Close(), Content(), Description(), Dialog(), Footer() (+24 more)

### Community 99 - "Entry"
Cohesion: 0.13
Nodes (17): TestUnflattenTable(), fixture, fixtures, normalizer, actionRank(), byPK(), children(), decode() (+9 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.19
Nodes (16): ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags(), ftEntry (+8 more)

### Community 101 - "EquipmentOption"
Cohesion: 0.29
Nodes (7): EquipmentOption, EquipmentItem, Env, WornArmor, gear, EquipmentChoice, applyEquipment()

### Community 102 - "Import"
Cohesion: 0.10
Nodes (28): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+20 more)

### Community 103 - "Site"
Cohesion: 0.13
Nodes (20): Entry, dataFor(), Kind, fiveToolsPage(), importSource(), readSource(), storeSource(), upsert() (+12 more)

### Community 104 - "Campaign"
Cohesion: 0.15
Nodes (9): Campaign, PendingKind, CreateCampaignParams, viewerCampaign(), deviceOS(), isPhone(), uaPick(), userAgentSummary() (+1 more)

### Community 105 - "App"
Cohesion: 0.07
Nodes (16): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Entity, App, blockAccessForm, entityTopics() (+8 more)

### Community 106 - "ImportView"
Cohesion: 0.22
Nodes (11): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), ImportReviewPage(), importRows() (+3 more)

### Community 107 - "dialog_templ.go"
Cohesion: 0.18
Nodes (12): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+4 more)

### Community 108 - "fromLookup"
Cohesion: 0.29
Nodes (8): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc

### Community 109 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 110 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, ItemData, o5eArmor, armorData(), armorText()

### Community 111 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): ContentProps, ctxKey, ctxState, Props, Side, DisableTrigger(), state()

### Community 112 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 113 - "Client"
Cohesion: 0.23
Nodes (3): Client, topicIDs(), message

### Community 114 - "Topic"
Cohesion: 0.18
Nodes (5): event, Handler, Topic, stateHandler, TopicKind

### Community 115 - "NewClient"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 116 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 117 - "toast_templ.go"
Cohesion: 0.33
Nodes (8): Props, ToasterProps, Type, Toast(), Toaster(), toastIcon(), toastListItem(), Toast()

### Community 118 - "CustomMod"
Cohesion: 0.16
Nodes (10): findMod(), newMod(), parseMod(), SaveFeature(), UsesLabel(), FeatureEffects, CustomMod, Feature (+2 more)

### Community 119 - "Parse"
Cohesion: 0.33
Nodes (7): Parse(), tokenize(), TestUnknownCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 121 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 122 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): Device, RosterCharacter, deviceLabel(), rosterHP()

### Community 123 - "generateSVG"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 125 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): Props, Size, initialChecked()

### Community 127 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 314 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **22 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `context.Context`, `SearchCompendiumRow`, `database/sql.NullInt64`, `Button`, `State`, `.do`, `perm.go`, `Character`, `newFixture`, `.blockFor`, `.serverError`, `Queries`, `wizardChoice`, `TLS Cert and Main`, `Character`, `Import`, `Campaign`, `parseLiveKey`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.07861990950226244 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `Weapon`, `EquipmentOption`, `Import`, `Character Sheet View`, `5e Command Apply`, `Queries`, `5e Level-Up and Slots`, `testing.T`, `perm.go`, `wizardChoice`, `Character`, `CustomMod`, `select_templ.go`, `findItem`, `builder_templ.go`?**
  _High betweenness centrality (0.033) - this node is a cross-community bridge._
- **Should `combobox_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1067193675889328 - nodes in this community are weakly interconnected._
- **Why does `Command` connect `5e Command Apply` to `Auth and Compendium Core`, `EquipmentOption`, `testing.T`, `perm.go`, `wizardChoice`, `CustomMod`, `Parse`, `Character`, `vendorjs Tool`, `findItem`, `builder_templ.go`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.012352576039726879 - nodes in this community are weakly interconnected._