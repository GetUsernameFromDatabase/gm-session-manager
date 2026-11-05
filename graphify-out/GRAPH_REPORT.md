# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 249 files · ~502,827 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3625 nodes · 14799 edges · 124 communities (102 shown, 22 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1476 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `661d6bf1`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- github.com/gin-gonic/gin.Context
- context.Context
- .do
- createProvider
- Character
- SheetView
- combobox.js
- redirect.go
- testing.T
- .serverError
- select.js
- command.go
- .Entry
- CompendiumView
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- github.com/a-h/templ.Component
- spellChecklist
- wikilink.js
- dialog_templ.go
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
- Encounter
- golang.org/x/net/html.Node
- tooltip.js
- Session
- .levelUp
- NameDesc
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
- LoadOrCreateCert
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
- formSelect
- useTypeahead
- ImageView
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
- .equipmentChoices
- .joinCampaigns
- RosterCharacter
- radiogroup_templ.go
- Sheet
- capReader
- build.sh
- errBodyWriter
- testing.T
- open5e.go
- Import

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
- `Creation and level-up pointers` --references--> `WizardView`  [INFERRED]
  HANDOFF.md → internal/web/views/wizard_templ.go
- `Creation and level-up pointers` --references--> `errorAt()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go
- `Creation and level-up pointers` --references--> `lostPicks()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go
- `Creation and level-up pointers` --references--> `userError`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (124 total, 22 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.12
Nodes (29): contextKey, keepAwake(), ASIError, clientMessage, TestMain(), TestMain(), isToggle(), CharacterImage (+21 more)

### Community 1 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (24): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), Campaign, coinCommand() (+16 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (83): keepAwake(), fetch(), ListImageCharactersRow, ListImageEntitiesRow, SearchReadableBlocksRow, AddBlockAudienceParams, AddBlockLinkParams, AddEntityFieldLinkParams (+75 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Character"
Cohesion: 0.11
Nodes (11): Character, Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile() (+3 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (108): Condition, ArmorEffect(), Signed(), Field, Button(), Variant, Action(), Card() (+100 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.07
Nodes (49): nullID(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait(), TestDeleteRestoreImport() (+41 more)

### Community 10 - ".serverError"
Cohesion: 0.11
Nodes (5): DisplayPage(), LinkItem, SwitcherResults(), blockAccessLevel, who

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.06
Nodes (67): Amount, CondOpKind, Effects, Env, HelpEntry, parser, Completions(), condCommand() (+59 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+22 more)

### Community 14 - "CompendiumView"
Cohesion: 0.08
Nodes (23): CountCompendiumRow, CountHabitatsRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, SearchCompendiumRow, CompendiumResults(), CompendiumRows() (+15 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.09
Nodes (41): Note, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanCharacter(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew() (+33 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (120): ListMonstersByNameRow, Description(), Content(), Description(), Empty(), Header(), Media(), Title() (+112 more)

### Community 21 - "spellChecklist"
Cohesion: 0.13
Nodes (16): spellTotals(), sortedColumns(), spellChecklist(), SpellColumn(), spellCounts(), TestSpellCounts(), backgroundAbilities(), choiceFieldset() (+8 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.07
Nodes (67): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+59 more)

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
Cohesion: 0.15
Nodes (26): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp() (+18 more)

### Community 28 - "App"
Cohesion: 0.07
Nodes (10): DecodeData(), ftsQuery(), Search(), CreateCombatantParams, compendiumWeapon(), viewerCampaign(), isHTMX(), useValidateTag() (+2 more)

### Community 29 - "monsterBody"
Cohesion: 0.10
Nodes (31): Props, Variant, SpellData, Badge(), badgeClasses(), alignmentColor(), classBadges(), classBody() (+23 more)

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
Nodes (30): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, CompendiumSource, firstLine(), mdLinkText() (+22 more)

### Community 34 - "State"
Cohesion: 0.15
Nodes (10): message, patchMessage, stateHandler, viewPatch, topicIDs(), Diff(), State, manyKeys() (+2 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Encounter"
Cohesion: 0.11
Nodes (16): Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), damageChoice(), initiativeMod(), outOfFight() (+8 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Session"
Cohesion: 0.16
Nodes (10): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+2 more)

### Community 41 - ".levelUp"
Cohesion: 0.11
Nodes (28): Creation and level-up pointers, firstTable(), LevelSpells(), Options(), TestOptions(), classFeatures(), classInfo(), errorField() (+20 more)

### Community 42 - "NameDesc"
Cohesion: 0.27
Nodes (10): OptionPrerequisite(), NameDesc, catchUpResistances(), TestCatchUpResistances(), upgradeClass(), UpgradeSheets(), upgradeSpecies(), isRepeatable() (+2 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Queries"
Cohesion: 0.10
Nodes (19): TopicKind, Queries, Entity, TestCampaignExcludesOwnerBlocks(), Topic, Fields, ParseFields(), applyCondOps() (+11 more)

### Community 46 - "Client"
Cohesion: 0.18
Nodes (3): fakeHandler, Handler, Client

### Community 47 - ".Plan"
Cohesion: 0.10
Nodes (17): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+9 more)

### Community 48 - "tabs_templ.go"
Cohesion: 0.11
Nodes (41): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+33 more)

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

### Community 54 - "LoadOrCreateCert"
Cohesion: 0.27
Nodes (9): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+1 more)

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
Nodes (37): Cantrips(), SkillChoice(), SplitList(), TestSkillChoice(), TextData, Slugify(), Spell, SizeNamed() (+29 more)

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
Cohesion: 0.07
Nodes (46): ImageCategory, Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select() (+38 more)

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
Cohesion: 0.02
Nodes (161): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+153 more)

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
Nodes (79): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, Combobox(), Content(), Item() (+71 more)

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
Cohesion: 0.15
Nodes (20): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+12 more)

### Community 86 - "formSelect"
Cohesion: 0.25
Nodes (16): Props, Checkbox(), initialChecked(), Label(), checkField(), abilityForm(), WizardView, skillName() (+8 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ImageView"
Cohesion: 0.18
Nodes (17): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, initials(), currentLook(), EntityDetail(), entityLooks() (+9 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+7 more)

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
Cohesion: 0.33
Nodes (5): Handoff, Lessons, Next up, Test data (data-test, all disposable), Test instance

### Community 107 - "Entry"
Cohesion: 0.21
Nodes (12): Entry, fixtures, normalizer, actionRank(), byPK(), children(), formatCR(), intField() (+4 more)

### Community 108 - "Import"
Cohesion: 0.09
Nodes (32): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply() (+24 more)

### Community 110 - ".equipmentChoices"
Cohesion: 0.25
Nodes (8): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, Squash(), TestCoreTraits(), TestEquipmentOptions(), EquipmentChoice

### Community 112 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.17
Nodes (15): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+7 more)

### Community 115 - "Sheet"
Cohesion: 0.07
Nodes (20): DeathSaves, Part, Skill, TestCustomMods(), skillByKey(), Sum(), TestDualWield(), TestGearCommands() (+12 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "testing.T"
Cohesion: 0.04
Nodes (56): Defense, fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves(), TestLanguages(), TestLevelUp() (+48 more)

### Community 133 - "open5e.go"
Cohesion: 0.08
Nodes (28): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+20 more)

### Community 134 - "Import"
Cohesion: 0.16
Nodes (20): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+12 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 334 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **22 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `github.com/gin-gonic/gin.Context`, `database/sql.NullInt64`, `.do`, `.loadBlock`, `Encounter`, `Hub`, `Character`, `State`, `.levelUp`, `.serverError`, `Import`, `Queries`, `.equipmentChoices`, `.joinCampaigns`, `perm.go`, `buildReview`, `Slugify`?**
  _High betweenness centrality (0.042) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11830869997599533 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Encounter`, `SheetView`, `builder_templ.go`, `.levelUp`, `NameDesc`, `command.go`, `.equipmentChoices`, `.Plan`, `Slugify`, `ParseSheet`, `App`, `testing.T`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.058018018018018015 - nodes in this community are weakly interconnected._
- **Why does `WizardView` connect `formSelect` to `compendium_templ.go`, `Character`, `SheetView`, `.levelUp`, `.serverError`, `.equipmentChoices`, `spellChecklist`, `ParseSheet`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.008877150508051795 - nodes in this community are weakly interconnected._