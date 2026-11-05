# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 249 files · ~502,877 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3624 nodes · 14802 edges · 142 communities (113 shown, 29 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1477 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `266d72c1`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- FromContext
- context.Context
- .do
- createProvider
- github.com/a-h/templ.Component
- SheetView
- combobox.js
- redirect.go
- testing.T
- github.com/gin-gonic/gin.Context
- select.js
- command.go
- .Entry
- ParseWikidot
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Page
- spellCounts
- wikilink.js
- CombatantView
- popover.js
- wikiText
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- monsterBody
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- State
- command.js
- dialog.js
- Defenses
- golang.org/x/net/html.Node
- tooltip.js
- Session
- NameDesc
- empty_templ.go
- export/markdown.go
- composite.js
- Queries
- Client
- .Plan
- tabs_templ.go
- DBTX
- New
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- buildReview
- golangci-lint Config
- Slugify
- toast_templ.go
- radiogroup.js
- vendorjs Tool
- select_templ.go
- Pending
- sqlc Config
- .loadBlock
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
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
- levelUpForm
- useTypeahead
- entityForm
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Hub
- Site
- .Serve
- NoteBlock
- Handoff
- Entry
- Import
- .ListSpellsForClass
- Feature
- .fail
- RosterCharacter
- Command
- radiogroup_templ.go
- Sheet
- scoped
- capReader
- build.sh
- popover_templ.go
- Item
- errBodyWriter
- Parse
- .sendFeedback
- Character
- parseHP
- tooltip_templ.go
- Env
- .mountAssets
- switch_templ.go
- open5e.go
- Import
- CountHabitatsRow
- SearchReadableBlocksRow
- userAgentSummary
- .ListBlockStepsSince
- .ListCompendiumByKind
- .ListCompendiumBySlugs
- .ListItemsByCategory

## God Nodes (most connected - your core abstractions)
1. `App` - 267 edges
2. `CN()` - 112 edges
3. `Button()` - 105 edges
4. `newTestEnv()` - 91 edges
5. `Sheet` - 90 edges
6. `SheetView` - 52 edges
7. `IfElse()` - 51 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 37 edges

## Surprising Connections (you probably didn't know these)
- `Lessons` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go
- `Creation and level-up pointers` --references--> `spellNeed()`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up` --references--> `TestRollAllInitiative()`  [INFERRED]
  HANDOFF.md → internal/web/turn_test.go
- `Creation and level-up pointers` --references--> `LevelUpView`  [INFERRED]
  HANDOFF.md → internal/web/views/levelup_templ.go
- `Creation and level-up pointers` --references--> `WizardView`  [INFERRED]
  HANDOFF.md → internal/web/views/wizard_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (142 total, 29 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (43): Props, contextKey, keepAwake(), ASIError, clientMessage, TestMain(), TestMain(), Direction() (+35 more)

### Community 1 - "FromContext"
Cohesion: 0.08
Nodes (17): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), browserName(), firstLine() (+9 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (73): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListImageEntitiesRow (+65 more)

### Community 3 - ".do"
Cohesion: 0.10
Nodes (24): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+16 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "github.com/a-h/templ.Component"
Cohesion: 0.09
Nodes (60): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+52 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (98): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+90 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (74): Next up, Completions(), TestCompletions(), TestLanguages(), TestLevelUp(), TestLevelUpSubclassAndPools(), TestArmorAndFeatureUses(), NewSheet() (+66 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (19): Campaign, coinCommand(), commandText(), compendiumBack(), viewerCampaign(), formIntQuery(), isHTMX(), isPhone() (+11 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.12
Nodes (23): CondOpKind, HelpEntry, parser, condCommand(), conditionSlug(), CondOp, init(), parseConcentrate() (+15 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (26): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInput, HomebrewInputFor() (+18 more)

### Community 14 - "ParseWikidot"
Cohesion: 0.14
Nodes (18): InputForEntry(), TestMonsterFormRoundTrip(), ParseWikidot(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSpecies(), TestParseWikidotSubclass() (+10 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.08
Nodes (44): ListCombatantsRow, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock() (+36 more)

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
Cohesion: 0.09
Nodes (43): Label(), Scripts(), Separator(), campaignLink(), CampaignOverview(), PendingGroup, PendingRequests(), CharacterPage() (+35 more)

### Community 21 - "spellCounts"
Cohesion: 0.67
Nodes (3): sortedColumns(), spellCounts(), TestSpellCounts()

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.10
Nodes (35): Props, Variant, Badge(), badgeClasses(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+27 more)

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
Cohesion: 0.16
Nodes (25): Creation and level-up pointers, ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp() (+17 more)

### Community 28 - "App"
Cohesion: 0.07
Nodes (18): DecodeData(), Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), compendiumWeapon() (+10 more)

### Community 29 - "monsterBody"
Cohesion: 0.10
Nodes (34): SpellData, alignmentColor(), classBadges(), classBody(), commaPrefixed(), compendiumRow(), entryBody(), EntryDetails() (+26 more)

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
Cohesion: 0.06
Nodes (34): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, CompendiumSource, Condition, Note (+26 more)

### Community 34 - "State"
Cohesion: 0.15
Nodes (10): message, patchMessage, stateHandler, viewPatch, topicIDs(), Diff(), State, manyKeys() (+2 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Defenses"
Cohesion: 0.33
Nodes (5): Defense, TestDefenses(), Defenses, ParseDefenses(), damageChoice()

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Session"
Cohesion: 0.16
Nodes (10): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+2 more)

### Community 41 - "NameDesc"
Cohesion: 0.17
Nodes (21): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, catchUpResistances(), TestCatchUpResistances() (+13 more)

### Community 42 - "empty_templ.go"
Cohesion: 0.10
Nodes (44): Props, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps (+36 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Queries"
Cohesion: 0.06
Nodes (25): TopicKind, ftsQuery(), Search(), Queries, Combatant, Encounter, Entity, CreateCombatantParams (+17 more)

### Community 46 - "Client"
Cohesion: 0.18
Nodes (3): fakeHandler, Handler, Client

### Community 47 - ".Plan"
Cohesion: 0.09
Nodes (20): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+12 more)

### Community 48 - "tabs_templ.go"
Cohesion: 0.06
Nodes (72): SearchCompendiumRow, Entity, Look, Content(), contentID(), IDFromContext(), initialValue(), List() (+64 more)

### Community 50 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

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
Cohesion: 0.10
Nodes (18): main(), run(), NewManager(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage() (+10 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "buildReview"
Cohesion: 0.24
Nodes (6): buildReview(), csvSource(), formValues(), parseCSV(), csvImport, csvImports

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Slugify"
Cohesion: 0.09
Nodes (33): Cantrips(), TextData, Slugify(), Spell, textMarkdown(), addClassSpells(), spellNeed(), upgradeSpecies() (+25 more)

### Community 59 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+16 more)

### Community 63 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (142): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Size, ActionProps (+134 more)

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
Cohesion: 0.05
Nodes (80): ImageCategory, ListBacklinksRow, Combobox(), Content(), Item(), List(), CloseFor(), Field() (+72 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.12
Nodes (19): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+11 more)

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
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "ParseWikidot"
Cohesion: 0.19
Nodes (11): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), sourceLine() (+3 more)

### Community 86 - "levelUpForm"
Cohesion: 0.15
Nodes (26): Footer(), abilitySelect(), levelUpForm(), slotSummary(), spellChecklist(), checkField(), abilityForm(), backgroundAbilities() (+18 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "entityForm"
Cohesion: 0.24
Nodes (13): ListEntityFieldBacklinksRow, attrString(), generateSVG(), getIconContent(), Props, Icon(), EntityDetail(), entityForm() (+5 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.19
Nodes (16): ftEntry, ftMonster, ftSpellcasting, ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType() (+8 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Hub"
Cohesion: 0.18
Nodes (3): ClientInfo, event, Hub

### Community 103 - "Site"
Cohesion: 0.12
Nodes (16): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+8 more)

### Community 105 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 106 - "Handoff"
Cohesion: 0.40
Nodes (4): Handoff, Lessons, Test data (data-test, all disposable), Test instance

### Community 107 - "Entry"
Cohesion: 0.30
Nodes (6): Entry, normalizer, children(), joinText(), label(), partsText()

### Community 108 - "Import"
Cohesion: 0.12
Nodes (26): fixture, migration, newFixture(), TestOnlyHashIsStored(), apply(), loadMigrations(), Migrate(), Open() (+18 more)

### Community 110 - "Feature"
Cohesion: 0.10
Nodes (24): EquipmentItem, FeatureEffects, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash() (+16 more)

### Community 111 - ".fail"
Cohesion: 0.12
Nodes (4): CanCharacter(), setAudience(), Flash(), blockAccessForm

### Community 112 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 113 - "Command"
Cohesion: 0.19
Nodes (15): Effects, findMod(), Command, NewCommand(), newMod(), parseMod(), parseWeapon(), SaveWeapon() (+7 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "Sheet"
Cohesion: 0.06
Nodes (21): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+13 more)

### Community 116 - "scoped"
Cohesion: 0.19
Nodes (14): parseCoin(), parseDeathSave(), parseHitDie(), parseInspiration(), parseMaxHP(), parsePact(), parseRest(), parseSet() (+6 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 123 - "Item"
Cohesion: 0.30
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Parse"
Cohesion: 0.18
Nodes (12): Heal(), Parse(), tokenize(), fighter(), TestCommands(), TestHealingFromZeroResetsDeathSaves(), TestParseRawArgs(), TestPrepareCap() (+4 more)

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 128 - "parseHP"
Cohesion: 0.46
Nodes (6): Amount, Damage(), parseAmount(), parseDamage(), parseHP(), Set()

### Community 129 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 130 - "Env"
Cohesion: 0.40
Nodes (5): Env, WornArmor, applyEquipment(), TestApplyEquipmentMergesGear(), gear

### Community 132 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

### Community 133 - "open5e.go"
Cohesion: 0.07
Nodes (35): ArmorData, fixture, fixtures, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass (+27 more)

### Community 134 - "Import"
Cohesion: 0.14
Nodes (22): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+14 more)

### Community 137 - "userAgentSummary"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **29 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `FromContext`, `database/sql.NullInt64`, `.do`, `.loadBlock`, `.mountAssets`, `Hub`, `State`, `github.com/gin-gonic/gin.Context`, `Queries`, `Feature`, `.fail`, `perm.go`, `.Plan`, `cert.go`, `buildReview`, `Slugify`?**
  _High betweenness centrality (0.045) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11076604554865424 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Env`, `Defenses`, `SheetView`, `ParseSheet`, `testing.T`, `NameDesc`, `Queries`, `Feature`, `.Plan`, `perm.go`, `Slugify`, `Item`, `App`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._
- **Should `FromContext` be split into smaller, more focused modules?**
  _Cohesion score 0.07777777777777778 - nodes in this community are weakly interconnected._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009706083390293916 - nodes in this community are weakly interconnected._
- **Should `.do` be split into smaller, more focused modules?**
  _Cohesion score 0.09841269841269841 - nodes in this community are weakly interconnected._