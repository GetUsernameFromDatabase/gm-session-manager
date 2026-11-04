# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 248 files · ~497,951 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3598 nodes · 14627 edges · 131 communities (110 shown, 21 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1464 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `13e4feaa`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- .do
- createProvider
- Queries
- SheetView
- combobox.js
- redirect.go
- testing.T
- App
- select.js
- command.go
- .Entry
- github.com/a-h/templ.Attributes
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- wikiText
- ImageResults
- wikilink.js
- dialog_templ.go
- popover.js
- .blockFor
- Project Rules (CLAUDE.md, README)
- ParseSheet
- Character
- Sheet
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- parseCommand
- command.js
- dialog.js
- Weapon
- golang.org/x/net/html.Node
- tooltip.js
- Store
- .levelUp
- Slugify
- export/markdown.go
- composite.js
- DecodeData
- Button
- .Plan
- IfElse
- Session
- github.com/a-h/templ.Component
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- levelUpForm
- golangci-lint Config
- NewSheet
- monsterBody
- radiogroup.js
- vendorjs Tool
- select_templ.go
- Handoff
- sqlc Config
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
- Player
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
- tabs_templ.go
- Site
- .ListItemsByCategory
- who
- inputgroup_templ.go
- Entry
- Import
- field_templ.go
- ParseWikidot
- Client
- CustomMod
- radiogroup_templ.go
- dialog_templ.go
- .ListBlockStepsSince
- card_templ.go
- build.sh
- .equipmentChoices
- alert_templ.go
- errBodyWriter
- Props
- .sendFeedback
- Props
- .joinInfo
- switch_templ.go
- SearchReadableBlocksRow
- spellCounts
- ArmorData

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
- `Where things live (creation and level-up)` --references--> `userError`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up (pick from these)` --references--> `addClassSpells()`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Next up (pick from these)` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go
- `Next up (pick from these)` --references--> `swapFeatSpell()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (131 total, 21 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.09
Nodes (54): contextKey, keepAwake(), looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes (+46 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.22
Nodes (12): state(), Trigger(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (66): keepAwake(), ftsQuery(), Search(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow (+58 more)

### Community 3 - ".do"
Cohesion: 0.13
Nodes (19): main(), serve(), testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions() (+11 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Queries"
Cohesion: 0.14
Nodes (12): Queries, Combatant, Character, applyCondOps(), damageChoice(), uniqueName(), linkEntityLook(), formInt() (+4 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (85): ArmorEffect(), Slot, Signed(), Field, Action(), Card(), Content(), Description() (+77 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (70): serveSplit(), TestSplitTLS(), TestSplitTLSSilentConn(), TestCanCharacter(), nullID(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName() (+62 more)

### Community 10 - "App"
Cohesion: 0.05
Nodes (27): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Entity, Toast(), entityTopics() (+19 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.07
Nodes (63): Amount, CondOpKind, Effects, Env, HelpEntry, parser, Completions(), condCommand() (+55 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (28): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+20 more)

### Community 14 - "github.com/a-h/templ.Attributes"
Cohesion: 0.07
Nodes (39): Props, ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps (+31 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (51): Note, ListCombatantsRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanCharacter() (+43 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 21 - "ImageResults"
Cohesion: 0.08
Nodes (55): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Look, Content(), Description(), Empty() (+47 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.10
Nodes (50): Defenses, Header(), Title(), capitalized(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+42 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - ".blockFor"
Cohesion: 0.08
Nodes (9): Attachment, GetBlockRow, firstLine(), mdLinkText(), saveScreenshot(), saveSniffed(), CreateAttachmentParams, blockAccessLevel (+1 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.14
Nodes (24): TestUpgradeSheet(), ParseSheet(), TestLevelUpFeats(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestUpgradeSheets() (+16 more)

### Community 28 - "Character"
Cohesion: 0.13
Nodes (10): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks(), Roster (+2 more)

### Community 29 - "Sheet"
Cohesion: 0.09
Nodes (13): DeathSaves, TestCustomMods(), Sum(), TestGearCommands(), Coins, Item, Sheet, WornArmor (+5 more)

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
Cohesion: 0.03
Nodes (42): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, CompendiumEntry, CompendiumSource, ensureCategory() (+34 more)

### Community 34 - "parseCommand"
Cohesion: 0.29
Nodes (7): coinCommand(), commandText(), hpCommand(), parseCommand(), deviceOS(), uaPick(), userAgentSummary()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Weapon"
Cohesion: 0.13
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.16
Nodes (17): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+9 more)

### Community 41 - ".levelUp"
Cohesion: 0.11
Nodes (33): Where things live (creation and level-up), addClassSpells(), classFeatures(), classInfo(), errorField(), spellbook(), spellNeed(), TestSpellNeed() (+25 more)

### Community 42 - "Slugify"
Cohesion: 0.13
Nodes (25): Cantrips(), SkillChoice(), SplitList(), TestSkillChoice(), TextData, Slugify(), Feature, Spell (+17 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "DecodeData"
Cohesion: 0.08
Nodes (11): DecodeData(), Campaign, Condition, compendiumWeapon(), combatantDefenses(), petrified(), viewerCampaign(), CreateCampaignParams (+3 more)

### Community 46 - "Button"
Cohesion: 0.07
Nodes (79): Button(), Dialog(), Close(), Content(), Description(), Dialog(), Footer(), Header() (+71 more)

### Community 47 - ".Plan"
Cohesion: 0.11
Nodes (15): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+7 more)

### Community 48 - "IfElse"
Cohesion: 0.08
Nodes (46): CountHabitatsRow, SearchCompendiumRow, Badge(), Trigger(), IfElse(), alignmentColor(), classBadges(), CompendiumPage() (+38 more)

### Community 49 - "Session"
Cohesion: 0.07
Nodes (24): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), TestCredentialsMatch() (+16 more)

### Community 50 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (84): ctxState, ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer() (+76 more)

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

### Community 56 - "levelUpForm"
Cohesion: 0.17
Nodes (21): spellChecklist(), abilityForm(), backgroundAbilities(), choiceFieldset(), choiceName(), choiceOptions(), WizardChoice, WizardView (+13 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "NewSheet"
Cohesion: 0.11
Nodes (20): Parse(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves(), TestLanguages(), TestLevelUp() (+12 more)

### Community 59 - "monsterBody"
Cohesion: 0.14
Nodes (25): SpellData, classBody(), commaPrefixed(), entryBody(), EntryDetails(), entryHead(), EntryPreview(), fact() (+17 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.30
Nodes (13): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), publicPath() (+5 more)

### Community 62 - "select_templ.go"
Cohesion: 0.08
Nodes (39): Content(), initialValue(), Item(), itemLabel(), Select(), state(), Trigger(), Value() (+31 more)

### Community 63 - "Handoff"
Cohesion: 0.18
Nodes (10): Defense, Handoff, Improvements (bigger, deferred), Next up (pick from these), Notes, Test data in data-test (all safe to delete), Working on the test instance, ParseDefenses() (+2 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

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
Nodes (55): Combobox(), Content(), Item(), List(), CloseFor(), Description(), Field(), Label() (+47 more)

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
Cohesion: 0.17
Nodes (12): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+4 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "shadcn-templ.go"
Cohesion: 0.10
Nodes (16): Props, Direction(), flattenClasses(), lastTopLevelColon(), MergeAttributes(), Ptr(), twImportantToV3(), twImportantToV4() (+8 more)

### Community 86 - "NameDesc"
Cohesion: 0.21
Nodes (14): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, catchUpResistances(), upgradeClass() (+6 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.18
Nodes (17): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "tabs_templ.go"
Cohesion: 0.20
Nodes (18): Content(), contentID(), IDFromContext(), initialValue(), listVariant(), listVariantClasses(), orientation(), state() (+10 more)

### Community 103 - "Site"
Cohesion: 0.12
Nodes (17): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+9 more)

### Community 105 - "who"
Cohesion: 0.15
Nodes (5): compendiumBack(), HomebrewView, LinkItem, SwitcherResults(), who

### Community 106 - "inputgroup_templ.go"
Cohesion: 0.16
Nodes (14): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+6 more)

### Community 107 - "Entry"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children() (+11 more)

### Community 108 - "Import"
Cohesion: 0.14
Nodes (24): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+16 more)

### Community 110 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 111 - "ParseWikidot"
Cohesion: 0.12
Nodes (21): block, ClassFeature, ClassData, InputForEntry(), TestMonsterFormRoundTrip(), partsText(), cellAt(), classTable() (+13 more)

### Community 112 - "Client"
Cohesion: 0.05
Nodes (37): ClientInfo, event, fakeHandler, Handler, message, patchMessage, PendingState, stateHandler (+29 more)

### Community 113 - "CustomMod"
Cohesion: 0.19
Nodes (5): FeatureEffects, UsesLabel(), CustomMod, Uses, usesValue()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "dialog_templ.go"
Cohesion: 0.18
Nodes (11): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+3 more)

### Community 117 - "card_templ.go"
Cohesion: 0.25
Nodes (8): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".equipmentChoices"
Cohesion: 0.25
Nodes (8): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, Squash(), TestCoreTraits(), TestEquipmentOptions(), EquipmentChoice

### Community 123 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 127 - "Props"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 129 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

### Community 131 - "spellCounts"
Cohesion: 0.67
Nodes (3): sortedColumns(), spellCounts(), TestSpellCounts()

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 334 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **21 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `database/sql.NullInt64`, `context.Context`, `.do`, `.joinInfo`, `Queries`, `who`, `.levelUp`, `Slugify`, `DecodeData`, `Client`, `perm.go`, `Session`, `.get`, `LoadOrCreateCert`, `.equipmentChoices`, `.blockFor`, `Character`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09197080291970802 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Queries`, `Weapon`, `SheetView`, `.levelUp`, `Slugify`, `command.go`, `DecodeData`, `.Plan`, `perm.go`, `CustomMod`, `NameDesc`, `.equipmentChoices`, `NewSheet`, `ParseSheet`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011751601225285437 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `.Plan` to `compendium_templ.go`, `.levelUp`, `NameDesc`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `.do` be split into smaller, more focused modules?**
  _Cohesion score 0.12807881773399016 - nodes in this community are weakly interconnected._