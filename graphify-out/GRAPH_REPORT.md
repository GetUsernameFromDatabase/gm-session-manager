# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 248 files · ~499,496 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3606 nodes · 14679 edges · 135 communities (115 shown, 20 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1471 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `ad7423e4`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- .do
- createProvider
- formInt
- SheetView
- combobox.js
- redirect.go
- testing.T
- github.com/gin-gonic/gin.Context
- select.js
- command.go
- .Entry
- monsterBody
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Button
- Look
- wikilink.js
- dialog_templ.go
- popover.js
- wikiText
- Project Rules (CLAUDE.md, README)
- ParseSheet
- who
- Weapon
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- deviceOS
- command.js
- dialog.js
- shadcn-templ.go
- golang.org/x/net/html.Node
- tooltip.js
- Store
- .levelUp
- .wizardCommand
- export/markdown.go
- composite.js
- time.Time
- dialog_templ.go
- .Plan
- tabs_templ.go
- combobox_templ.go
- inputgroup_templ.go
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- field_templ.go
- golangci-lint Config
- NameDesc
- Badge
- radiogroup.js
- vendorjs Tool
- github.com/a-h/templ.Attributes
- Handoff
- sqlc Config
- Spell
- Browser check
- Claude Settings
- github.com/a-h/templ.Component
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
- checkbox.js
- ImportView
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- ParseWikidot
- Lint and Release Rules
- Run Script
- formSelect
- useTypeahead
- ImportView
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Client
- Site
- .ListItemsByCategory
- command_templ.go
- Feature
- Entry
- Import
- dialog_templ.go
- wizardChoice
- tooltip_templ.go
- alert_templ.go
- Props
- radiogroup_templ.go
- Sheet
- Props
- capReader
- build.sh
- spellCounts
- SearchReadableBlocksRow
- errBodyWriter
- NewSheet
- .sendFeedback
- .ListCompendiumBySlugs
- LevelUpView
- importSource
- lanHostname
- ArmorData
- buildReview
- Campaign
- toast_templ.go

## God Nodes (most connected - your core abstractions)
1. `App` - 266 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Sheet` - 90 edges
5. `newTestEnv()` - 88 edges
6. `SheetView` - 52 edges
7. `IfElse()` - 50 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 36 edges

## Surprising Connections (you probably didn't know these)
- `Where things live (creation and level-up)` --references--> `userError`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Where things live (creation and level-up)` --references--> `LevelUpView`  [INFERRED]
  HANDOFF.md → internal/web/views/levelup_templ.go
- `Next up (pick from these)` --references--> `spellChecklist()`  [INFERRED]
  HANDOFF.md → internal/web/views/levelup_templ.go
- `Where things live (creation and level-up)` --references--> `taughtPicked()`  [INFERRED]
  HANDOFF.md → internal/web/views/levelup_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (135 total, 20 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (51): contextKey, keepAwake(), looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes (+43 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps (+3 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (60): keepAwake(), ftsQuery(), Search(), fetch(), CountCompendiumRow, ListHomebrewLinksRow, ListAttachmentFilesRow, Encounter (+52 more)

### Community 3 - ".do"
Cohesion: 0.16
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "formInt"
Cohesion: 0.07
Nodes (18): Queries, Combatant, Entity, Character, uniqueName(), entityTopics(), linkEntityLook(), formInt() (+10 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (80): ArmorEffect(), Signed(), Field, Action(), Card(), Content(), Description(), Header() (+72 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (66): TestModifier(), Modifier(), TestCanCharacter(), nullID(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName(), TestLanguagesForm() (+58 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (20): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Campaign, viewerCampaign(), formIntQuery(), isHTMX() (+12 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.06
Nodes (67): Amount, CondOpKind, Effects, Env, HelpEntry, parser, Completions(), condCommand() (+59 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+22 more)

### Community 14 - "monsterBody"
Cohesion: 0.14
Nodes (25): SpellData, classBody(), commaPrefixed(), entryBody(), EntryDetails(), EntryPage(), fact(), inHabitatOrder() (+17 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (56): GetBlockRow, ListCombatantsRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+48 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Button"
Cohesion: 0.06
Nodes (82): Button(), Content(), Description(), Empty(), Header(), Media(), mediaClasses(), Title() (+74 more)

### Community 21 - "Look"
Cohesion: 0.33
Nodes (9): Look, initials(), removeButton(), currentLook(), lookControls(), LookItem(), looksGrid(), looksSection() (+1 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.12
Nodes (33): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), downBadge() (+25 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.17
Nodes (22): ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestSpellCounts() (+14 more)

### Community 28 - "who"
Cohesion: 0.06
Nodes (16): DecodeData(), Character, CompendiumEntry, CompendiumSource, compendiumWeapon(), compendiumBack(), Roster, HomebrewView (+8 more)

### Community 29 - "Weapon"
Cohesion: 0.15
Nodes (8): Part, Skill, skillByKey(), TestDualWield(), damageLabel(), Weapon, ProficiencyAt(), attackTargets()

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
Cohesion: 0.04
Nodes (38): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Attachment, Condition, CreateCombatantParams, applyCondOps(), petrified() (+30 more)

### Community 34 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "shadcn-templ.go"
Cohesion: 0.11
Nodes (17): Props, initialChecked(), Direction(), flattenClasses(), lastTopLevelColon(), MergeAttributes(), Ptr(), twImportantToV3() (+9 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.07
Nodes (25): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+17 more)

### Community 41 - ".levelUp"
Cohesion: 0.15
Nodes (23): Where things live (creation and level-up), ClassFeature, ClassInfo, LevelUpPlan, addClassSpells(), classFeatures(), classInfo(), errorField() (+15 more)

### Community 42 - ".wizardCommand"
Cohesion: 0.23
Nodes (11): TextData, Feature, SizeNamed(), textMarkdown(), addFeat(), addProfs(), applyBackground(), applyPicks() (+3 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "time.Time"
Cohesion: 0.05
Nodes (25): Identity, ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListImagesRow, Session, ListBacklinksRow, browserName() (+17 more)

### Community 46 - "dialog_templ.go"
Cohesion: 0.14
Nodes (34): Defenses, damageChoice(), Dialog(), Close(), Content(), Description(), Dialog(), Footer() (+26 more)

### Community 47 - ".Plan"
Cohesion: 0.12
Nodes (12): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+4 more)

### Community 48 - "tabs_templ.go"
Cohesion: 0.20
Nodes (18): Content(), contentID(), IDFromContext(), initialValue(), listVariant(), listVariantClasses(), orientation(), state() (+10 more)

### Community 49 - "combobox_templ.go"
Cohesion: 0.11
Nodes (22): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+14 more)

### Community 50 - "inputgroup_templ.go"
Cohesion: 0.16
Nodes (14): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+6 more)

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

### Community 56 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "NameDesc"
Cohesion: 0.24
Nodes (12): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, catchUpResistances(), UpgradeSheets() (+4 more)

### Community 59 - "Badge"
Cohesion: 0.10
Nodes (33): CountHabitatsRow, SearchCompendiumRow, Badge(), Spinner(), alignmentColor(), classBadges(), CompendiumPage(), CompendiumResults() (+25 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "github.com/a-h/templ.Attributes"
Cohesion: 0.07
Nodes (37): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+29 more)

### Community 63 - "Handoff"
Cohesion: 0.22
Nodes (8): Handoff, Improvements (bigger, deferred), Lessons (read before orchestrating), Notes, Test data in data-test (all safe to delete), Working on the test instance, readPicks(), TestReadPicks()

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Spell"
Cohesion: 0.23
Nodes (12): Cantrips(), Slugify(), Spell, choiceOptions(), featList(), freeCast(), lineageAbility(), skillPick() (+4 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (113): ctxState, ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer() (+105 more)

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
Cohesion: 0.07
Nodes (64): CloseFor(), Description(), Field(), Group(), Label(), Input(), Textarea(), armorForm() (+56 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.16
Nodes (13): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), TestPlayerCombatantLink(), LiveKey(), dial() (+5 more)

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

### Community 83 - "ParseWikidot"
Cohesion: 0.11
Nodes (24): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, partsText(), pageSource(), PageSourceKey() (+16 more)

### Community 86 - "formSelect"
Cohesion: 0.10
Nodes (39): Label(), rosterCard(), formSelect(), selectItems(), homebrewAction(), homebrewLinks(), HomebrewPage(), kindSelect() (+31 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ImportView"
Cohesion: 0.19
Nodes (13): ImageCategory, categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), ImportReviewPage() (+5 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.18
Nodes (17): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Client"
Cohesion: 0.05
Nodes (36): ClientInfo, event, fakeHandler, Handler, message, patchMessage, PendingState, stateHandler (+28 more)

### Community 103 - "Site"
Cohesion: 0.15
Nodes (11): Site, dataFor(), Kind, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SavePage() (+3 more)

### Community 105 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 106 - "Feature"
Cohesion: 0.24
Nodes (4): FeatureEffects, UsesLabel(), CustomMod, Uses

### Community 107 - "Entry"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children() (+11 more)

### Community 108 - "Import"
Cohesion: 0.08
Nodes (40): migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+32 more)

### Community 109 - "dialog_templ.go"
Cohesion: 0.18
Nodes (11): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+3 more)

### Community 110 - "wizardChoice"
Cohesion: 0.21
Nodes (13): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+5 more)

### Community 111 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 112 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 113 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "Sheet"
Cohesion: 0.09
Nodes (16): DeathSaves, TestCustomMods(), Sum(), TestGearCommands(), TestUpgradeSheet(), Coins, Item, Sheet (+8 more)

### Community 116 - "Props"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "NewSheet"
Cohesion: 0.08
Nodes (27): Defense, Next up (pick from these), Parse(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+19 more)

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 128 - "LevelUpView"
Cohesion: 0.27
Nodes (7): LevelUpView, Option, SpellNeed, Plural(), drawsGroup(), ChoiceGroup, CreationChoices

### Community 129 - "importSource"
Cohesion: 0.47
Nodes (5): Source, importSource(), readSource(), storeSource(), upsert()

### Community 130 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

### Community 134 - "buildReview"
Cohesion: 0.31
Nodes (5): ensureCategory(), buildReview(), formValues(), csvImport, csvImports

### Community 135 - "Campaign"
Cohesion: 0.28
Nodes (6): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks()

### Community 140 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 334 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **20 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `github.com/gin-gonic/gin.Context` to `compendium_templ.go`, `Spell`, `context.Context`, `LevelUpView`, `database/sql.NullInt64`, `formInt`, `Client`, `Campaign`, `Store`, `.levelUp`, `buildReview`, `.wizardCommand`, `time.Time`, `wizardChoice`, `perm.go`, `LoadOrCreateCert`, `who`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10113772630224328 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `LevelUpView`, `formInt`, `SheetView`, `command.go`, `perm.go`, `ParseSheet`, `who`, `Weapon`, `.levelUp`, `.wizardCommand`, `dialog_templ.go`, `.Plan`, `NameDesc`, `Spell`, `builder_templ.go`, `Feature`, `wizardChoice`, `NewSheet`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.012338425381903642 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `.levelUp` to `compendium_templ.go`, `NameDesc`, `.Plan`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._