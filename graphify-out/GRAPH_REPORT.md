# Graph Report - gm-session-manager  (2026-10-26)

## Corpus Check
- 233 files · ~451,068 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 49 file(s) not represented in the graph (top: .templ 41, .css 4, (none) 3)

## Summary
- 3311 nodes · 12973 edges · 122 communities (103 shown, 19 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1298 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `aa1dd92d`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Auth and Compendium Core
- combobox_templ.go
- context.Context
- HomebrewInputFor
- github.com/a-h/templ.Attributes
- database/sql.NullInt64
- Character Sheet View
- live
- EquipmentOption
- .do
- github.com/a-h/templ.Component
- createProvider
- 5e Command Apply
- .Entry
- Encounter
- Combobox JS
- Viewer
- Select JS
- Floating UI DOM
- Hover Interaction JS
- tooltip_templ.go
- popover_templ.go
- Notes Editor JS
- CombatantView
- Anchor Positioning JS
- parseCommand
- Project Rules (CLAUDE.md, README)
- tabs_templ.go
- 5e Sheet Model
- wizardPick
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
- rowFigure
- Dialog JS
- cell
- Tooltip JS
- App
- attachmentRefs
- 5e Level-Up and Slots
- Parse
- Badge
- Composite List Navigation
- wizardChoice
- Release Workflow
- Floating Focus Manager
- TLS Cert and Main
- shadcn-templ Config
- .fail
- golangci-lint Config
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
- Entry
- .websocket
- Queries
- select_templ.go
- fromLookup
- command_templ.go
- ArmorData
- wikiText
- field_templ.go
- redirect.go
- Hub
- alert_templ.go
- toast_templ.go
- CustomMod
- Props
- generateSVG
- switch_templ.go
- empty_templ.go
- SearchReadableBlocksRow
- lanHostname

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
- `slotLabel()` --references--> `Slot`  [EXTRACTED]
  web/views/live.go → dnd5e/sheet.go
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

## Communities (122 total, 19 thin omitted)

### Community 0 - "Auth and Compendium Core"
Cohesion: 0.08
Nodes (66): TestMain(), contextKey, fixture, Identity, Manager, keepAwake(), looseString, o5eAssignment (+58 more)

### Community 1 - "combobox_templ.go"
Cohesion: 0.11
Nodes (22): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+14 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (59): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, CreateCombatantParams, AddBlockAudienceParams (+51 more)

### Community 3 - "HomebrewInputFor"
Cohesion: 0.22
Nodes (10): CoreTraits(), TestCoreTraits(), HomebrewInputFor(), actionsText(), bonusText(), speedText(), TableText(), TraitsText() (+2 more)

### Community 4 - "github.com/a-h/templ.Attributes"
Cohesion: 0.06
Nodes (42): Props, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size (+34 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (42): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumSource, Condition (+34 more)

### Community 6 - "Character Sheet View"
Cohesion: 0.07
Nodes (81): Signed(), Field, actionEntry, edit, profFocus, subtitlePart, Action(), Card() (+73 more)

### Community 7 - "live"
Cohesion: 0.13
Nodes (26): ListEntityFieldBacklinksRow, ListBacklinksRow, ImageView, PendingView, campaignLink(), PendingGroup, PendingRequests(), EntityDetail() (+18 more)

### Community 8 - "EquipmentOption"
Cohesion: 0.25
Nodes (8): EquipmentOption, EquipmentItem, ArmorEffect(), TestExplainDamage(), WornArmor, gear, EquipmentChoice, applyEquipment()

### Community 9 - ".do"
Cohesion: 0.14
Nodes (19): isCleared(), main(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+11 more)

### Community 10 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (122): ListMonstersByNameRow, WS, Action(), Alert(), Description(), Title(), AspectRatio(), Footer() (+114 more)

### Community 11 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 12 - "5e Command Apply"
Cohesion: 0.08
Nodes (47): Amount, condCommand(), conditionSlug(), Damage(), Command, CondOp, Effects, Heal() (+39 more)

### Community 13 - ".Entry"
Cohesion: 0.11
Nodes (23): addEntry(), setStat(), TestHomebrew(), MonsterData, NameDesc, ValidKind(), HomebrewOwner, HomebrewInput (+15 more)

### Community 14 - "Encounter"
Cohesion: 0.20
Nodes (9): Encounter, CreateEncounterParams, ListCharacterEncountersParams, CombatView, CombatPage(), currentID(), encounterLinks(), EncounterList() (+1 more)

### Community 15 - "Combobox JS"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (60): Note, GetBlockRow, ListCombatantsRow, blockCase, Combatant, Fields, ParseFields(), VisibleEntity() (+52 more)

### Community 17 - "Select JS"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 18 - "Floating UI DOM"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 19 - "Hover Interaction JS"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 20 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): ContentProps, ctxKey, ctxState, Props, Side, DisableTrigger(), state()

### Community 21 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps, Props, Side (+3 more)

### Community 22 - "Notes Editor JS"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.12
Nodes (29): combatantGroup, CombatantView, addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+21 more)

### Community 24 - "Anchor Positioning JS"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 25 - "parseCommand"
Cohesion: 0.33
Nodes (6): featureCommand(), weaponCommand(), coinCommand(), commandText(), hpCommand(), parseCommand()

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "tabs_templ.go"
Cohesion: 0.22
Nodes (19): ContentProps, ctxKey, ctxState, ListProps, Orientation, Props, TriggerProps, Variant (+11 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.07
Nodes (17): DeathSaves, TestExplainSums(), skillByKey(), Sum(), TestUpgradeSheet(), Coins, Sheet, Slot (+9 more)

### Community 29 - "wizardPick"
Cohesion: 0.10
Nodes (38): SpellData, FeaturePreview, WizardStep, classBody(), commaPrefixed(), entryBody(), EntryDetails(), fact() (+30 more)

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
Cohesion: 0.10
Nodes (12): constantTimeEqual(), credentialsMatch(), IsLoopback(), TestCredentialsMatch(), TestIsLoopback(), TestTokens(), Role, hashToken() (+4 more)

### Community 35 - "App JS Command Bar"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 36 - "Dismiss Interaction JS"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 37 - "ParseWikidot"
Cohesion: 0.15
Nodes (19): block, TextData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), sourceLine() (+11 more)

### Community 38 - "ParseAidedd"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "Command Palette JS"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 40 - "Open"
Cohesion: 0.13
Nodes (21): apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock(), TestFTS5Available() (+13 more)

### Community 41 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 42 - "Dialog JS"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 43 - "cell"
Cohesion: 0.10
Nodes (25): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+17 more)

### Community 44 - "Tooltip JS"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 45 - "App"
Cohesion: 0.06
Nodes (22): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Campaign, PendingKind, App, compendiumBack() (+14 more)

### Community 47 - "5e Level-Up and Slots"
Cohesion: 0.12
Nodes (14): ClassLevel, TestSlotMaxima(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan, hpGain(), raise() (+6 more)

### Community 48 - "Parse"
Cohesion: 0.13
Nodes (18): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+10 more)

### Community 49 - "Badge"
Cohesion: 0.12
Nodes (26): CountHabitatsRow, SearchCompendiumRow, PickerTarget, Badge(), alignmentColor(), classBadges(), CompendiumResults(), compendiumRow() (+18 more)

### Community 50 - "Composite List Navigation"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 51 - "wizardChoice"
Cohesion: 0.18
Nodes (16): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), armorCondition(), columnInts() (+8 more)

### Community 52 - "Release Workflow"
Cohesion: 0.11
Nodes (18): CGO_ENABLED, GH_TOKEN, jobs, release, name, permissions, contents, tags (+10 more)

### Community 53 - "Floating Focus Manager"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 54 - "TLS Cert and Main"
Cohesion: 0.10
Nodes (17): NewManager(), main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - ".fail"
Cohesion: 0.06
Nodes (24): DecodeData(), Slugify(), Character, CompendiumEntry, ParseSheet(), Campaign(), CharacterFileName(), fileSlug() (+16 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 59 - "Markdown Export"
Cohesion: 0.28
Nodes (9): TestMarkdown(), Links, applyMarks(), escape(), Markdown(), plain(), pmMark, pmNode (+1 more)

### Community 60 - "List Navigation JS"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.23
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 62 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 63 - "builder_templ.go"
Cohesion: 0.08
Nodes (64): ctxState, ctxKey, ctxState, Props, sourceField, Chip(), Chips(), Combobox() (+56 more)

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
Cohesion: 0.13
Nodes (19): builderOp, builderVerb, armorForm(), builderForm(), BuilderPickItems(), BuilderPicks(), builderRemove(), builderURL() (+11 more)

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

### Community 80 - "testing.T"
Cohesion: 0.07
Nodes (58): newFixture(), serve(), TestMiddlewareIdentity(), TestOnlyHashIsStored(), TestRequireAdmin(), TestRealSource(), Import(), TestCanCharacter() (+50 more)

### Community 82 - "Lifecycle JS"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 83 - "Button"
Cohesion: 0.10
Nodes (43): Look, FileRow, FilesView, ImagesView, Button(), Item(), IfElse(), CampaignOverview() (+35 more)

### Community 99 - "Entry"
Cohesion: 0.13
Nodes (16): TestUnflattenTable(), fixture, fixtures, normalizer, actionRank(), byPK(), children(), decode() (+8 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType(), ParseFiveTools(), stripTags() (+7 more)

### Community 101 - ".Router"
Cohesion: 0.24
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 102 - "newStore"
Cohesion: 0.21
Nodes (13): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestSearch() (+5 more)

### Community 103 - "Entry"
Cohesion: 0.12
Nodes (22): ClassFeature, Entry, dataFor(), ClassData, Kind, fiveToolsPage(), importSource(), readSource() (+14 more)

### Community 104 - ".websocket"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 105 - "Queries"
Cohesion: 0.06
Nodes (23): ftsQuery(), Search(), Queries, Combatant, Entity, Character, SetCharacterCurrentImageParams, SetEntityCurrentImageParams (+15 more)

### Community 106 - "select_templ.go"
Cohesion: 0.06
Nodes (47): ImageCategory, Align, ContentProps, ctxKey, ctxState, ItemData, Props, Size (+39 more)

### Community 108 - "fromLookup"
Cohesion: 0.29
Nodes (8): describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence(), TestLoadMissingFileIsNotAnError(), lookupFunc

### Community 109 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 110 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, ItemData, o5eArmor, armorData(), armorText()

### Community 111 - "wikiText"
Cohesion: 0.27
Nodes (8): WikiLink, WikiPart, wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks()

### Community 112 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 113 - "redirect.go"
Cohesion: 0.24
Nodes (5): connListener, peekedConn, newConnListener(), route(), SplitTLS()

### Community 115 - "Hub"
Cohesion: 0.06
Nodes (33): ClientInfo, event, fakeHandler, Handler, Client, Hub, Topic, New() (+25 more)

### Community 116 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 117 - "toast_templ.go"
Cohesion: 0.33
Nodes (8): Props, ToasterProps, Type, Toast(), Toaster(), toastIcon(), toastListItem(), Toast()

### Community 118 - "CustomMod"
Cohesion: 0.15
Nodes (14): findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel(), FeatureEffects (+6 more)

### Community 121 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 123 - "generateSVG"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 125 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): Props, Size, initialChecked()

### Community 128 - "empty_templ.go"
Cohesion: 0.28
Nodes (8): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, mediaClasses()

### Community 136 - "SearchReadableBlocksRow"
Cohesion: 0.50
Nodes (3): SearchReadableBlocksRow, SearchReadableBlocksParams, NoteListView

### Community 141 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **13 isolated node(s):** `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc`, `Table-Driven Tests for Permission Logic` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 315 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **19 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Auth and Compendium Core`, `context.Context`, `database/sql.NullInt64`, `.Router`, `newStore`, `Queries`, `.do`, `Viewer`, `Hub`, `TLS Cert and Main`, `.fail`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **What connects `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Auth and Compendium Core` be split into smaller, more focused modules?**
  _Cohesion score 0.08197348640153068 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `Auth and Compendium Core`, `Weapon`, `Character Sheet View`, `EquipmentOption`, `Queries`, `5e Level-Up and Slots`, `Parse`, `Viewer`, `wizardChoice`, `CustomMod`, `.fail`, `Item`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._
- **Should `combobox_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1067193675889328 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Hub` to `Auth and Compendium Core`, `App`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.012348111658456486 - nodes in this community are weakly interconnected._