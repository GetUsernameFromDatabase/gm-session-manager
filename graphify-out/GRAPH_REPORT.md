# Graph Report - gm-session-manager  (2026-11-03)

## Corpus Check
- 245 files · ~466,724 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 63 file(s) not represented in the graph (top: .templ 43, .patch 12, .css 4)

## Summary
- 3455 nodes · 13652 edges · 118 communities (99 shown, 19 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1339 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `66419ef1`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- github.com/a-h/templ.Attributes
- context.Context
- testEnv
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
- Campaign
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
- select_templ.go
- command.js
- dialog.js
- ImportView
- golang.org/x/net/html.Node
- tooltip.js
- Store
- ParseWikidot
- ParseEffects
- export/markdown.go
- composite.js
- .fail
- HomebrewInputFor
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
- toast_templ.go
- .sendFeedback
- radiogroup.js
- vendorjs Tool
- findItem
- .blockFor
- sqlc Config
- .saveReport
- Browser check
- Claude Settings
- Parse
- use_transition_status.js
- playwright
- release.sh
- builderFields
- checkbox.js
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- testing.T
- useFocus
- mark_others.js
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
- Character
- deviceOS
- Entry
- config.go
- open5e.go
- tooltip_templ.go
- radiogroup_templ.go
- Hub
- Verbs
- build.sh
- App
- IfElse
- capReader

## God Nodes (most connected - your core abstractions)
1. `App` - 247 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `newTestEnv()` - 74 edges
5. `Sheet` - 58 edges
6. `SheetView` - 49 edges
7. `IfElse()` - 47 edges
8. `Input()` - 46 edges
9. `Command` - 43 edges
10. `Client` - 35 edges

## Surprising Connections (you probably didn't know these)
- `Bugs / fixes` --references--> `RequireAdmin()`  [INFERRED]
  HANDOFF.md → internal/auth/middleware.go
- `Bugs / fixes` --references--> `ParseEffects()`  [INFERRED]
  HANDOFF.md → internal/dnd5e/effects.go
- `Bugs / fixes` --references--> `applyBackground()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Bugs / fixes` --references--> `applyEquipment()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Bugs / fixes` --references--> `editable()`  [INFERRED]
  HANDOFF.md → internal/web/views/character_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (118 total, 19 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.09
Nodes (52): Props, contextKey, keepAwake(), clientMessage, TestMain(), newValidator(), TestMain(), Direction() (+44 more)

### Community 1 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (165): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Size, ActionProps (+157 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (55): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Note, AddBlockAudienceParams (+47 more)

### Community 3 - "testEnv"
Cohesion: 0.09
Nodes (16): Bugs / fixes, /characters/{id} (player character sheet), Checked and fine (character page), Handoff: discovery review (2026-11-03), Improvements, Improvements, Other pages, Test data left in data-test (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (34): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, Attachment, saveScreenshot(), imageItem(), imageVersion(), nullString() (+26 more)

### Community 6 - "SheetView"
Cohesion: 0.08
Nodes (73): Bugs / fixes, ArmorEffect(), Signed(), Field, Variant, Action(), Card(), Content() (+65 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - ".do"
Cohesion: 0.16
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 10 - "dialog_templ.go"
Cohesion: 0.25
Nodes (25): ctxKey, ctxState, Props, Dialog(), Close(), Content(), Description(), Dialog() (+17 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.10
Nodes (48): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+40 more)

### Community 13 - ".Entry"
Cohesion: 0.12
Nodes (20): MonsterAction, addEntry(), setStat(), MonsterData, NameDesc, ValidKind(), abilityKey(), actionsText() (+12 more)

### Community 14 - "Queries"
Cohesion: 0.08
Nodes (19): Queries, Entity, CreateCombatantParams, applyCondOps(), uniqueName(), entityTopics(), linkEntityLook(), saveFieldLinks() (+11 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (58): SearchReadableBlocksRow, ListCombatantsRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+50 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Campaign"
Cohesion: 0.11
Nodes (8): ListAttachmentFilesRow, Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

### Community 21 - "popover_templ.go"
Cohesion: 0.16
Nodes (21): Condition, Content(), Description(), Header(), state(), Title(), Trigger(), triggerID() (+13 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "combat_templ.go"
Cohesion: 0.13
Nodes (30): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), downBadge() (+22 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "monsterBody"
Cohesion: 0.22
Nodes (19): SpellData, classBody(), entryBody(), EntryDetails(), fact(), itemBody(), levelsText(), monsterBody() (+11 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "time.Time"
Cohesion: 0.11
Nodes (9): ListImagesRow, Campaign, BlockStep, CreateCampaignParams, LevelUp, ListImagesParams, ListPlayerSessionsByCampaignRow, reqEntry (+1 more)

### Community 28 - "5e Sheet Model"
Cohesion: 0.08
Nodes (12): DeathSaves, TestModifier(), TestExplainSums(), Sum(), TestGearCommands(), TestUpgradeSheet(), Coins, Item (+4 more)

### Community 29 - "github.com/a-h/templ.Component"
Cohesion: 0.04
Nodes (138): ListMonstersByNameRow, ListEntityFieldBacklinksRow, ListBacklinksRow, Look, Description(), Button(), Content(), Description() (+130 more)

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
Cohesion: 0.15
Nodes (9): Part, Skill, skillByKey(), TestDualWield(), damageLabel(), Weapon, ProficiencyAt(), attackTargets() (+1 more)

### Community 34 - "select_templ.go"
Cohesion: 0.07
Nodes (48): Label(), Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select() (+40 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "ImportView"
Cohesion: 0.15
Nodes (15): Props, ImageCategory, Checkbox(), initialChecked(), categoryRow(), columnLabel(), columnOverview(), ImportCategory (+7 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.11
Nodes (26): migration, importSource(), readSource(), storeSource(), SavePage(), apply(), loadMigrations(), Migrate() (+18 more)

### Community 41 - "ParseWikidot"
Cohesion: 0.12
Nodes (23): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, pageSource(), PageSourceKey(), cellAt() (+15 more)

### Community 42 - "ParseEffects"
Cohesion: 0.18
Nodes (16): EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice(), armorCondition(), columnInts() (+8 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".fail"
Cohesion: 0.12
Nodes (8): Character, CanCharacter(), Roster, SetCharacterCurrentImageParams, ToggleCharacterImageRevealedParams, UnlinkCharacterImageParams, RosterCharacter, blockAccessForm

### Community 46 - "HomebrewInputFor"
Cohesion: 0.15
Nodes (14): HomebrewOwner, Source, TestHomebrew(), HomebrewInput, HomebrewInputFor(), InputForEntry(), SaveHomebrew(), upsert() (+6 more)

### Community 47 - ".Plan"
Cohesion: 0.14
Nodes (12): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan, hpGain() (+4 more)

### Community 48 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 49 - "Manager"
Cohesion: 0.11
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+10 more)

### Community 50 - "Parse"
Cohesion: 0.20
Nodes (11): findMod(), newMod(), parseMod(), SaveFeature(), IsModTarget(), featureCommand(), weaponCommand(), coinCommand() (+3 more)

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
Nodes (8): TopicKind, Encounter, Topic, notFoundIsNil(), parseLiveKey(), CreateEncounterParams, ListCharacterEncountersParams, liveKey

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 59 - ".sendFeedback"
Cohesion: 0.33
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "findItem"
Cohesion: 0.38
Nodes (3): findItem(), parseArmor(), ensureItem()

### Community 63 - ".blockFor"
Cohesion: 0.22
Nodes (3): GetBlockRow, blockAccessLevel, stepsPayload

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Parse"
Cohesion: 0.13
Nodes (18): Completions(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestLanguages() (+10 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "builderFields"
Cohesion: 0.06
Nodes (65): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, CloseFor(), Description(), Field() (+57 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

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

### Community 80 - "testing.T"
Cohesion: 0.05
Nodes (81): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+73 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 86 - "WornArmor"
Cohesion: 0.18
Nodes (8): EquipmentItem, Env, EquipmentOption, WornArmor, EquipmentChoice, applyEquipment(), TestApplyEquipmentMergesGear(), gear

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Props"
Cohesion: 0.38
Nodes (7): attrString(), generateSVG(), getIconContent(), Props, Icon(), Item(), feedbackKind()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.25
Nodes (14): ftEntry, ftMonster, ftSpellcasting, ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType() (+6 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 103 - "Site"
Cohesion: 0.16
Nodes (10): Site, dataFor(), Kind, fiveToolsPage(), FetchPage(), Page, ParsePageURL(), SiteByKey() (+2 more)

### Community 105 - "Character"
Cohesion: 0.07
Nodes (23): DecodeData(), TextData, Slugify(), Combatant, CompendiumEntry, CompendiumSource, Character, ParseSheet() (+15 more)

### Community 106 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 107 - "Entry"
Cohesion: 0.21
Nodes (12): Entry, fixtures, normalizer, actionRank(), byPK(), children(), formatCR(), intField() (+4 more)

### Community 108 - "config.go"
Cohesion: 0.17
Nodes (12): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+4 more)

### Community 110 - "open5e.go"
Cohesion: 0.08
Nodes (28): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+20 more)

### Community 113 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "Hub"
Cohesion: 0.05
Nodes (34): ClientInfo, event, fakeHandler, Handler, message, patchMessage, PendingState, stateHandler (+26 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 125 - "App"
Cohesion: 0.05
Nodes (33): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), ftsQuery(), Search(), compendiumBack() (+25 more)

### Community 131 - "IfElse"
Cohesion: 0.05
Nodes (73): Props, Variant, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), RadioGroup(), Content() (+65 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 331 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **19 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.saveReport`, `context.Context`, `testEnv`, `database/sql.NullInt64`, `Character`, `config.go`, `.fail`, `Queries`, `perm.go`, `Hub`, `Campaign`, `cert.go`, `.loadLive`, `time.Time`, `.blockFor`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09261893013670798 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `5e Sheet Model` to `compendium_templ.go`, `Weapon`, `select_templ.go`, `Parse`, `SheetView`, `builderFields`, `Character`, `ParseEffects`, `.Plan`, `perm.go`, `WornArmor`, `findItem`?**
  _High betweenness centrality (0.019) - this node is a cross-community bridge._
- **Should `github.com/a-h/templ.Attributes` be split into smaller, more focused modules?**
  _Cohesion score 0.02252443688907777 - nodes in this community are weakly interconnected._
- **Why does `Hub` connect `Hub` to `compendium_templ.go`, `App`?**
  _High betweenness centrality (0.014) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.013087365123564218 - nodes in this community are weakly interconnected._