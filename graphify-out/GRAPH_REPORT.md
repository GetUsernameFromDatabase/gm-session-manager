# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 249 files · ~502,052 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3619 nodes · 14773 edges · 126 communities (102 shown, 24 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1467 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f0dd84da`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- .render
- context.Context
- .do
- createProvider
- Campaign
- SheetView
- combobox.js
- redirect.go
- testing.T
- github.com/gin-gonic/gin.Context
- select.js
- command.go
- .Entry
- popover_templ.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- github.com/a-h/templ.Component
- ImportView
- wikilink.js
- dialog_templ.go
- popover.js
- wikiText
- Project Rules (CLAUDE.md, README)
- ParseSheet
- .fail
- monsterBody
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- Campaign
- command.js
- dialog.js
- formInt
- golang.org/x/net/html.Node
- tooltip.js
- Store
- Spell
- .sendFeedback
- export/markdown.go
- composite.js
- FromContext
- empty_templ.go
- .Plan
- tabs_templ.go
- DBTX
- Character
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
- tooltip_templ.go
- sqlc Config
- setStat
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
- ImagesPage
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Client
- Site
- Entity
- NoteBlock
- .ListCompendiumByKind
- Entry
- Import
- .ListSpellsForClass
- .featChoices
- switch_templ.go
- credentialsMatch
- .ListBlockStepsSince
- radiogroup_templ.go
- Sheet
- capReader
- build.sh
- SearchReadableBlocksRow
- errBodyWriter
- Parse
- ArmorData
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 267 edges
2. `CN()` - 112 edges
3. `Button()` - 104 edges
4. `Sheet` - 90 edges
5. `newTestEnv()` - 90 edges
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
- `Creation and level-up pointers` --references--> `meetsPrerequisite()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Creation and level-up pointers` --references--> `errorAt()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go
- `Creation and level-up pointers` --references--> `lostPicks()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (126 total, 24 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (51): contextKey, keepAwake(), looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes (+43 more)

### Community 1 - ".render"
Cohesion: 0.08
Nodes (12): ftsQuery(), Search(), compendiumBack(), viewerCampaign(), formIntQuery(), isHTMX(), joinPath(), HomebrewView (+4 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (80): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, CountHabitatsRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow (+72 more)

### Community 3 - ".do"
Cohesion: 0.17
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Campaign"
Cohesion: 0.20
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

### Community 6 - "SheetView"
Cohesion: 0.08
Nodes (81): ArmorEffect(), Signed(), Field, Button(), Variant, Action(), Card(), Content() (+73 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (71): TestCanCharacter(), nullID(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+63 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (14): TopicKind, SetCookie(), writeCookie(), Entity, Topic, uniqueName(), entityTopics(), linkEntityLook() (+6 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.06
Nodes (68): Amount, CondOpKind, Effects, Env, HelpEntry, parser, Completions(), condCommand() (+60 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (24): HomebrewOwner, MonsterAction, TestHomebrew(), ValidKind(), HomebrewInput, HomebrewInputFor(), InputForEntry(), SaveHomebrew() (+16 more)

### Community 14 - "popover_templ.go"
Cohesion: 0.15
Nodes (19): Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition(), conditionLabel() (+11 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (50): Encounter, GetBlockRow, ListCombatantsRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+42 more)

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
Nodes (101): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+93 more)

### Community 21 - "ImportView"
Cohesion: 0.27
Nodes (9): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), ImportReviewPage(), importRows() (+1 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.11
Nodes (38): Defenses, capitalized(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+30 more)

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
Cohesion: 0.19
Nodes (22): ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved() (+14 more)

### Community 28 - ".fail"
Cohesion: 0.07
Nodes (15): DecodeData(), Character, CompendiumEntry, CanCharacter(), compendiumWeapon(), classFeatures(), classInfo(), errorField() (+7 more)

### Community 29 - "monsterBody"
Cohesion: 0.12
Nodes (28): SpellData, TestModifier(), Modifier(), classBody(), commaPrefixed(), crText(), entryBody(), EntryDetails() (+20 more)

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
Cohesion: 0.05
Nodes (42): Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, Note, CreateCombatantParams (+34 more)

### Community 34 - "Campaign"
Cohesion: 0.23
Nodes (7): Campaign, deviceOS(), isPhone(), uaPick(), userAgentSummary(), visitorOnly(), Roster

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "formInt"
Cohesion: 0.30
Nodes (4): Combatant, Character, damageChoice(), formInt()

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.10
Nodes (16): fixture, tokens, IsLoopback(), isCleared(), newFixture(), serve(), TestIsLoopback(), TestMiddlewareIdentity() (+8 more)

### Community 41 - "Spell"
Cohesion: 0.08
Nodes (26): Creation and level-up pointers, Handoff, Lessons, Next up, Test data (data-test, all disposable), Test instance, Spell, addClassSpells() (+18 more)

### Community 42 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "FromContext"
Cohesion: 0.07
Nodes (21): ClearCookie(), Identity, Role, FromContext(), RequireAdmin(), CompendiumSource, Session, browserName() (+13 more)

### Community 46 - "empty_templ.go"
Cohesion: 0.12
Nodes (34): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Content() (+26 more)

### Community 47 - ".Plan"
Cohesion: 0.09
Nodes (19): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+11 more)

### Community 48 - "tabs_templ.go"
Cohesion: 0.06
Nodes (69): Props, Variant, SearchCompendiumRow, Badge(), badgeClasses(), Content(), contentID(), IDFromContext() (+61 more)

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

### Community 56 - "buildReview"
Cohesion: 0.36
Nodes (4): buildReview(), formValues(), csvImport, csvImports

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Slugify"
Cohesion: 0.10
Nodes (36): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, TextData (+28 more)

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
Cohesion: 0.10
Nodes (27): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+19 more)

### Community 63 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "setStat"
Cohesion: 0.33
Nodes (7): addEntry(), setStat(), MonsterData, abilityKey(), parseBonuses(), parseSpeed(), TestMonsterFormParsers()

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (164): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+156 more)

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
Nodes (64): Combobox(), Content(), Item(), List(), CloseFor(), Field(), Label(), Input() (+56 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.13
Nodes (18): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+10 more)

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
Cohesion: 0.16
Nodes (18): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), sourceLine() (+10 more)

### Community 86 - "formSelect"
Cohesion: 0.17
Nodes (24): Label(), spellChecklist(), checkField(), abilityForm(), backgroundAbilities(), choiceFieldset(), choiceName(), choiceOptions() (+16 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ImagesPage"
Cohesion: 0.09
Nodes (39): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Look, initials(), currentLook(), EntityDetail() (+31 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+7 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Client"
Cohesion: 0.05
Nodes (35): ClientInfo, event, fakeHandler, Handler, message, patchMessage, PendingState, stateHandler (+27 more)

### Community 103 - "Site"
Cohesion: 0.12
Nodes (17): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+9 more)

### Community 105 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 107 - "Entry"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children() (+11 more)

### Community 108 - "Import"
Cohesion: 0.12
Nodes (19): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+11 more)

### Community 110 - ".featChoices"
Cohesion: 0.16
Nodes (16): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestCoreTraits() (+8 more)

### Community 111 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

### Community 112 - "credentialsMatch"
Cohesion: 0.67
Nodes (3): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "Sheet"
Cohesion: 0.05
Nodes (30): DeathSaves, FeatureEffects, Part, Skill, TestCustomMods(), UsesLabel(), skillByKey(), Sum() (+22 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Parse"
Cohesion: 0.09
Nodes (25): Defense, Parse(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves(), TestLanguages() (+17 more)

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

### Community 134 - "Import"
Cohesion: 0.14
Nodes (24): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+16 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **24 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `github.com/gin-gonic/gin.Context` to `compendium_templ.go`, `database/sql.NullInt64`, `.render`, `.do`, `Campaign`, `formInt`, `Client`, `Campaign`, `Spell`, `Import`, `FromContext`, `.featChoices`, `.Plan`, `perm.go`, `LoadOrCreateCert`, `buildReview`, `Slugify`, `.fail`?**
  _High betweenness centrality (0.033) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10295106899476471 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `formInt`, `SheetView`, `Spell`, `command.go`, `.featChoices`, `.Plan`, `perm.go`, `Slugify`, `ParseSheet`, `.fail`, `Parse`?**
  _High betweenness centrality (0.025) - this node is a cross-community bridge._
- **Should `.render` be split into smaller, more focused modules?**
  _Cohesion score 0.08367254635911352 - nodes in this community are weakly interconnected._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009259259259259259 - nodes in this community are weakly interconnected._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._