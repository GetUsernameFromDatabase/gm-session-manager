# Graph Report - gm-session-manager  (2026-11-03)

## Corpus Check
- 242 files · ~464,064 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3437 nodes · 13573 edges · 128 communities (109 shown, 19 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1330 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `0de41f7f`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- github.com/a-h/templ.Attributes
- context.Context
- tabs_templ.go
- createProvider
- database/sql.NullInt64
- SheetView
- combobox.js
- redirect.go
- .do
- dialog_templ.go
- select.js
- command.go
- .Entry
- Queries
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- IfElse
- popover_templ.go
- wikilink.js
- combat_templ.go
- popover.js
- monsterBody
- Project Rules (CLAUDE.md, README)
- time.Time
- 5e Sheet Model
- github.com/a-h/templ.Component
- app.js
- tabbable.js
- useDismiss
- Weapon
- levelUpForm
- command.js
- dialog.js
- select_templ.go
- golang.org/x/net/html.Node
- tooltip.js
- Store
- ParseWikidot
- ParseEffects
- export/markdown.go
- composite.js
- github.com/gin-gonic/gin.Context
- combobox_templ.go
- .Plan
- cell
- Manager
- Parse
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- .loadLive
- golangci-lint Config
- Note
- .sendFeedback
- radiogroup.js
- vendorjs Tool
- Item
- Input
- sqlc Config
- .encounterView
- buildReview
- Claude Settings
- fighter
- use_transition_status.js
- Sed Hook
- release.sh
- builderFields
- checkbox.js
- empty_templ.go
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- testing.T
- useFocus
- mark_others.js
- Button
- Lint and Release Rules
- Run Script
- WornArmor
- useTypeahead
- Props
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- shadcn-templ.go
- Character
- deviceOS
- Entry
- config.go
- inputgroup_templ.go
- open5e.go
- wikiText
- field_templ.go
- tooltip_templ.go
- radiogroup_templ.go
- Hub
- Verbs
- command_templ.go
- build.sh
- switch_templ.go
- alert_templ.go
- App
- Props
- Badge
- capReader

## God Nodes (most connected - your core abstractions)
1. `App` - 246 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `newTestEnv()` - 74 edges
5. `Sheet` - 57 edges
6. `SheetView` - 48 edges
7. `Input()` - 46 edges
8. `IfElse()` - 45 edges
9. `Command` - 43 edges
10. `Client` - 35 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `NewManager()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/auth/manager.go
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Open()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/db/db.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (128 total, 19 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (42): contextKey, keepAwake(), clientMessage, TestMain(), SearchReadableBlocksRow, TestMain(), isToggle(), CharacterImage (+34 more)

### Community 1 - "github.com/a-h/templ.Attributes"
Cohesion: 0.11
Nodes (25): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+17 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (67): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Encounter, CreateCombatantParams (+59 more)

### Community 3 - "tabs_templ.go"
Cohesion: 0.22
Nodes (15): IDFromContext(), initialValue(), listVariant(), listVariantClasses(), orientation(), state(), tabsIsActive(), ContentProps (+7 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.04
Nodes (36): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Entity, saveScreenshot() (+28 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (75): ArmorEffect(), Signed(), Field, Action(), Card(), Content(), Description(), Header() (+67 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - ".do"
Cohesion: 0.22
Nodes (13): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+5 more)

### Community 10 - "dialog_templ.go"
Cohesion: 0.18
Nodes (12): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+4 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.11
Nodes (45): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+37 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (31): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, NameDesc, ValidKind(), HomebrewInput (+23 more)

### Community 14 - "Queries"
Cohesion: 0.10
Nodes (14): Queries, Combatant, Character, TestCampaignExcludesOwnerBlocks(), applyCondOps(), uniqueName(), linkEntityLook(), formInt() (+6 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.20
Nodes (25): VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock(), CanManageNote() (+17 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "IfElse"
Cohesion: 0.13
Nodes (34): Entity, Look, Content(), contentID(), List(), Tabs(), Trigger(), triggerID() (+26 more)

### Community 21 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps (+3 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "combat_templ.go"
Cohesion: 0.10
Nodes (50): Condition, Close(), Input(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName() (+42 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "monsterBody"
Cohesion: 0.13
Nodes (28): SpellData, classBody(), commaPrefixed(), entryBody(), EntryDetails(), entryHead(), EntryPage(), EntryPreview() (+20 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "time.Time"
Cohesion: 0.07
Nodes (19): Identity, FromContext(), Campaign, CompendiumSource, Session, browserName(), firstLine(), mdLinkText() (+11 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (12): DeathSaves, Skill, TestExplainSums(), skillByKey(), Sum(), TestUpgradeSheet(), Coins, Sheet (+4 more)

### Community 29 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (134): ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer(), Checkbox() (+126 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Weapon"
Cohesion: 0.18
Nodes (8): TestModifier(), Part, TestDualWield(), damageLabel(), Weapon, Modifier(), ProficiencyAt(), attackTargets()

### Community 34 - "levelUpForm"
Cohesion: 0.16
Nodes (25): Label(), abilitySelect(), LevelUpView, levelUpForm(), LevelUpPage(), optionItems(), slotSummary(), sortedColumns() (+17 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "select_templ.go"
Cohesion: 0.07
Nodes (41): ImageCategory, Content(), initialValue(), Item(), itemLabel(), Select(), state(), Trigger() (+33 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.09
Nodes (36): migration, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+28 more)

### Community 41 - "ParseWikidot"
Cohesion: 0.14
Nodes (20): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, cellAt(), classTable(), featuresUnder() (+12 more)

### Community 42 - "ParseEffects"
Cohesion: 0.14
Nodes (20): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), TextData, armorCondition() (+12 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (14): ClearCookie(), SetCookie(), writeCookie(), CanCharacter(), entityTopics(), isHTMX(), isPhone(), redirect() (+6 more)

### Community 46 - "combobox_templ.go"
Cohesion: 0.11
Nodes (22): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+14 more)

### Community 47 - ".Plan"
Cohesion: 0.14
Nodes (12): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan, hpGain() (+4 more)

### Community 48 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 49 - "Manager"
Cohesion: 0.07
Nodes (24): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+16 more)

### Community 50 - "Parse"
Cohesion: 0.12
Nodes (16): findMod(), NewCommand(), newMod(), Parse(), parseMod(), SaveWeapon(), tokenize(), TestUnknownCommand() (+8 more)

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
Cohesion: 0.11
Nodes (17): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - ".loadLive"
Cohesion: 0.11
Nodes (28): ListEntityFieldBacklinksRow, ListBacklinksRow, notFoundIsNil(), campaignLink(), PendingGroup, PendingRequests(), EntityDetail(), entityForm() (+20 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Note"
Cohesion: 0.13
Nodes (8): Note, Fields, ParseFields(), Toast(), saveFieldLinks(), Toast(), CreateNoteParams, ListReadableNotesParams

### Community 59 - ".sendFeedback"
Cohesion: 0.19
Nodes (14): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+6 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 63 - "Input"
Cohesion: 0.23
Nodes (17): ctxState, Chip(), Combobox(), Content(), Input(), Item(), itemLabel(), List() (+9 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - ".encounterView"
Cohesion: 0.27
Nodes (10): ListCombatantsRow, initiativeMod(), outOfFight(), pcSheet(), stepTurn(), combatantKind(), combatantLink(), playerCombatantLink() (+2 more)

### Community 66 - "buildReview"
Cohesion: 0.27
Nodes (5): ensureCategory(), buildReview(), formValues(), csvImport, csvImports

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "fighter"
Cohesion: 0.20
Nodes (8): fighter(), TestCommands(), TestCustomMods(), TestHealingFromZeroResetsDeathSaves(), TestParseRawArgs(), geared(), TestExplainDamage(), TestStowKeepsMods()

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "builderFields"
Cohesion: 0.06
Nodes (49): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, CloseFor(), Description(), Field() (+41 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "empty_templ.go"
Cohesion: 0.28
Nodes (8): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, mediaClasses()

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.17
Nodes (12): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), LiveKey(), dial(), patchOf() (+4 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "testing.T"
Cohesion: 0.05
Nodes (81): TestRealSource(), Import(), Completions(), TestCompletions(), TestLanguages(), TestLevelUp(), TestLevelUpSubclassAndPools(), TestArmorAndFeatureUses() (+73 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "Button"
Cohesion: 0.17
Nodes (25): Button(), CampaignOverview(), CampaignRoster(), Roster, rosterCard(), ImageItem, imageAction(), imageBindForm() (+17 more)

### Community 86 - "WornArmor"
Cohesion: 0.33
Nodes (6): EquipmentItem, EquipmentOption, WornArmor, EquipmentChoice, applyEquipment(), gear

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Props"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.21
Nodes (16): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+8 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 103 - "Site"
Cohesion: 0.12
Nodes (17): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+9 more)

### Community 104 - "shadcn-templ.go"
Cohesion: 0.12
Nodes (14): Props, Direction(), flattenClasses(), lastTopLevelColon(), MergeAttributes(), Ptr(), twImportantToV3(), twImportantToV4() (+6 more)

### Community 105 - "Character"
Cohesion: 0.06
Nodes (24): DecodeData(), Slugify(), Character, CompendiumEntry, ParseSheet(), Campaign(), CharacterFileName(), fileSlug() (+16 more)

### Community 106 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 107 - "Entry"
Cohesion: 0.28
Nodes (8): Entry, fixtures, normalizer, byPK(), children(), joinText(), parseFixtures(), parseOpen5e()

### Community 108 - "config.go"
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 109 - "inputgroup_templ.go"
Cohesion: 0.16
Nodes (14): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+6 more)

### Community 110 - "open5e.go"
Cohesion: 0.07
Nodes (30): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+22 more)

### Community 111 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 112 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 113 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "Hub"
Cohesion: 0.05
Nodes (35): ClientInfo, event, fakeHandler, Handler, message, patchMessage, PendingState, stateHandler (+27 more)

### Community 117 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

### Community 123 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 125 - "App"
Cohesion: 0.06
Nodes (16): GetBlockRow, compendiumBack(), viewerCampaign(), formIntQuery(), imageItem(), imageVersion(), joinPath(), HomebrewView (+8 more)

### Community 127 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 131 - "Badge"
Cohesion: 0.10
Nodes (31): CountHabitatsRow, SearchCompendiumRow, Badge(), alignmentColor(), classBadges(), CompendiumResults(), compendiumRow(), CompendiumRows() (+23 more)

## Knowledge Gaps
- **12 isolated node(s):** `CGO_ENABLED`, `gm-session-manager`, `testEnv`, `Permissions Table (characters, monster HP, entities)`, `SQLite (modernc, WAL, foreign_keys) via sqlc` (+7 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 325 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **19 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.encounterView`, `context.Context`, `buildReview`, `database/sql.NullInt64`, `time.Time`, `Character`, `github.com/gin-gonic/gin.Context`, `Queries`, `perm.go`, `Manager`, `Hub`, `cert.go`, `.loadLive`, `Note`, `.sendFeedback`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **What connects `CGO_ENABLED`, `gm-session-manager`, `testEnv` to the rest of the system?**
  _12 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10263812842210054 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `compendium_templ.go`, `Weapon`, `.encounterView`, `levelUpForm`, `SheetView`, `builderFields`, `Character`, `ParseEffects`, `Queries`, `.Plan`, `testing.T`, `WornArmor`, `Item`?**
  _High betweenness centrality (0.023) - this node is a cross-community bridge._
- **Should `github.com/a-h/templ.Attributes` be split into smaller, more focused modules?**
  _Cohesion score 0.10591133004926108 - nodes in this community are weakly interconnected._
- **Why does `Client` connect `Hub` to `compendium_templ.go`, `buildReview`, `redirect.go`, `perm.go`, `time.Time`, `App`?**
  _High betweenness centrality (0.018) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011355311355311355 - nodes in this community are weakly interconnected._