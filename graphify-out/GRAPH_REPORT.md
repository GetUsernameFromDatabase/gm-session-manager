# Graph Report - gm-session-manager  (2026-11-03)

## Corpus Check
- 246 files · ~478,416 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3501 nodes · 13918 edges · 125 communities (101 shown, 24 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1374 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `6dc79c0d`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- select_templ.go
- context.Context
- .do
- createProvider
- database/sql.NullInt64
- SheetView
- combobox.js
- redirect.go
- newTestEnv
- .loadLive
- select.js
- command.go
- .Entry
- empty_templ.go
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- who
- .fail
- wikilink.js
- combat_templ.go
- popover.js
- Note
- Project Rules (CLAUDE.md, README)
- time.Time
- Sheet
- github.com/a-h/templ.Component
- app.js
- tabbable.js
- useDismiss
- builderForm
- HomebrewInputFor
- command.js
- dialog.js
- State
- golang.org/x/net/html.Node
- tooltip.js
- Store
- ParseEffects
- Slugify
- export/markdown.go
- composite.js
- .sendFeedback
- github.com/a-h/templ.Attributes
- .Plan
- .mountAssets
- Session
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- toast_templ.go
- golangci-lint Config
- Parse
- coinList
- radiogroup.js
- vendorjs Tool
- IfElse
- .ListBlockStepsSince
- sqlc Config
- RosterCharacter
- Browser check
- Claude Settings
- .ListCompendiumByKind
- use_transition_status.js
- playwright
- release.sh
- CN
- checkbox.js
- CustomMod
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- testing.T
- useFocus
- mark_others.js
- .ListCompendiumBySlugs
- Lint and Release Rules
- Run Script
- .ListSpellsForClass
- useTypeahead
- blockMenu
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- popover_templ.go
- Encounter
- deviceOS
- Entry
- config.go
- message
- wikidotClass
- Client
- radiogroup_templ.go
- Hub
- New
- build.sh
- wikiText
- Pending
- App
- Handoff: open items from the 2026-11-03 review
- SearchReadableBlocksRow
- capReader

## God Nodes (most connected - your core abstractions)
1. `App` - 254 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `newTestEnv()` - 78 edges
5. `Sheet` - 67 edges
6. `SheetView` - 50 edges
7. `IfElse()` - 48 edges
8. `Input()` - 46 edges
9. `Command` - 43 edges
10. `Client` - 35 edges

## Surprising Connections (you probably didn't know these)
- `Next up (pick from these)` --references--> `Options()`  [INFERRED]
  HANDOFF.md → internal/compendium/coretraits.go
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Fingerprint()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/join/cert.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (125 total, 24 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (46): contextKey, keepAwake(), looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes (+38 more)

### Community 1 - "select_templ.go"
Cohesion: 0.05
Nodes (73): ImageCategory, Label(), Content(), Group(), initialValue(), Item(), itemLabel(), Label() (+65 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (65): keepAwake(), ftsQuery(), Search(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow (+57 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (44): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition (+36 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (99): Props, Size, ArmorEffect(), Slot, Signed(), Field, Button(), Variant (+91 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.25
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "newTestEnv"
Cohesion: 0.06
Nodes (47): TestRealSource(), Import(), nullID(), TestBuilderCompendium(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource() (+39 more)

### Community 10 - ".loadLive"
Cohesion: 0.10
Nodes (13): TopicKind, Topic, Fields, ParseFields(), saveFieldLinks(), notFoundIsNil(), parseLiveKey(), toLook() (+5 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.07
Nodes (64): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+56 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (27): ArmorData, HomebrewOwner, MonsterAction, o5eArmor, addEntry(), setStat(), TestHomebrew(), ItemData (+19 more)

### Community 14 - "empty_templ.go"
Cohesion: 0.05
Nodes (70): Props, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps (+62 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.11
Nodes (39): ListCombatantsRow, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock() (+31 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "who"
Cohesion: 0.11
Nodes (6): GetBlockRow, compendiumBack(), HomebrewView, blockAccessLevel, stepsPayload, who

### Community 21 - ".fail"
Cohesion: 0.08
Nodes (17): RequireAdmin(), DecodeData(), Character, ParseSheet(), CanCharacter(), compendiumWeapon(), textMarkdown(), classInfo() (+9 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "combat_templ.go"
Cohesion: 0.10
Nodes (35): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+27 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "Note"
Cohesion: 0.10
Nodes (12): Note, Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks() (+4 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "time.Time"
Cohesion: 0.08
Nodes (15): Identity, CompendiumEntry, CompendiumSource, deviceKey(), keepsBody(), sessionKey(), BlockStep, CreateHomebrewParams (+7 more)

### Community 28 - "Sheet"
Cohesion: 0.07
Nodes (18): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+10 more)

### Community 29 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (99): SpellData, SearchCompendiumRow, Alert(), Description(), Kbd(), Scripts(), Separator(), Spinner() (+91 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "builderForm"
Cohesion: 0.18
Nodes (12): armorForm(), builderForm(), builderURL(), chooseFirst(), CommandBuilder(), compendiumLookup(), modList(), opGroup() (+4 more)

### Community 34 - "HomebrewInputFor"
Cohesion: 0.22
Nodes (10): CoreTraits(), TestCoreTraits(), HomebrewInputFor(), actionsText(), bonusText(), speedText(), TableText(), TraitsText() (+2 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "State"
Cohesion: 0.16
Nodes (7): fakeHandler, patchMessage, stateHandler, viewPatch, State, manyKeys(), liveView

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.08
Nodes (41): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), copySource() (+33 more)

### Community 41 - "ParseEffects"
Cohesion: 0.15
Nodes (17): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+9 more)

### Community 42 - "Slugify"
Cohesion: 0.15
Nodes (22): Cantrips(), OptionLevel(), Options(), TestOptions(), NameDesc, Slugify(), Spell, SizeNamed() (+14 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 46 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (130): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ActionProps, ContentProps (+122 more)

### Community 47 - ".Plan"
Cohesion: 0.14
Nodes (10): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), ClassInfo, LevelUpPlan, hpGain(), raise() (+2 more)

### Community 49 - "Session"
Cohesion: 0.12
Nodes (12): tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), TestCredentialsMatch(), TestIsLoopback(), TestTokens(), Manager (+4 more)

### Community 50 - "cell"
Cohesion: 0.11
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

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
Cohesion: 0.24
Nodes (10): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+2 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Parse"
Cohesion: 0.26
Nodes (5): findItem(), parseArmor(), TestGearCommands(), ensureItem(), Item

### Community 59 - "coinList"
Cohesion: 0.67
Nodes (3): Coins, coinList(), coin

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "IfElse"
Cohesion: 0.09
Nodes (50): CountHabitatsRow, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses() (+42 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "CN"
Cohesion: 0.06
Nodes (86): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+78 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "CustomMod"
Cohesion: 0.15
Nodes (15): FeatureEffects, armorCondition(), columnInts(), ParseEffects(), SaveFeature(), TestParseEffects(), UsesLabel(), CustomMod (+7 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.15
Nodes (14): Player(), newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveVisibility(), TestPlayerCombatantLink(), ImageSrc(), LiveKey() (+6 more)

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
Nodes (45): Completions(), fighter(), TestCommands(), TestCompletions(), TestLanguages(), TestLevelUp(), TestLevelUpSubclassAndPools(), TestParseRawArgs() (+37 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "blockMenu"
Cohesion: 0.13
Nodes (24): ListEntityFieldBacklinksRow, ListBacklinksRow, initialChecked(), Switch(), EntityDetail(), entityForm(), entityLooks(), EntityView (+16 more)

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
Cohesion: 0.12
Nodes (16): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+8 more)

### Community 104 - "popover_templ.go"
Cohesion: 0.16
Nodes (21): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+13 more)

### Community 105 - "Encounter"
Cohesion: 0.10
Nodes (12): Combatant, Entity, Character, entityTopics(), linkEntityLook(), formInt(), lookFree(), Toast() (+4 more)

### Community 106 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 107 - "Entry"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children() (+11 more)

### Community 108 - "config.go"
Cohesion: 0.17
Nodes (12): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+4 more)

### Community 109 - "message"
Cohesion: 0.33
Nodes (4): message, topicIDs(), Diff(), TestDiff()

### Community 111 - "wikidotClass"
Cohesion: 0.12
Nodes (23): block, ClassFeature, ClassData, TextData, withCoreTraits(), cellAt(), classTable(), featuresUnder() (+15 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.20
Nodes (13): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), state() (+5 more)

### Community 115 - "Hub"
Cohesion: 0.18
Nodes (3): ClientInfo, event, Hub

### Community 117 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 123 - "wikiText"
Cohesion: 0.31
Nodes (7): wikiLink(), wikiText(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 124 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 125 - "App"
Cohesion: 0.06
Nodes (20): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Campaign, viewerCampaign(), browserName(), firstLine() (+12 more)

### Community 126 - "Handoff: open items from the 2026-11-03 review"
Cohesion: 0.22
Nodes (8): Bugs fixed since, Character page, Handoff: open items from the 2026-11-03 review, Improvements (bigger, deferred), Next up (pick from these), Notes, Other pages, Test data in data-test

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **24 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `database/sql.NullInt64`, `State`, `Store`, `Encounter`, `.loadLive`, `Slugify`, `config.go`, `perm.go`, `.mountAssets`, `Hub`, `who`, `.fail`, `Note`, `time.Time`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.0954633849011697 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `select_templ.go`, `context.Context`, `SheetView`, `Encounter`, `CustomMod`, `Slugify`, `ParseEffects`, `.Plan`, `testing.T`, `perm.go`, `.fail`, `Parse`, `coinList`?**
  _High betweenness centrality (0.016) - this node is a cross-community bridge._
- **Should `select_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.04673135852911134 - nodes in this community are weakly interconnected._
- **Why does `dial()` connect `Player` to `compendium_templ.go`, `testing.T`, `Client`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011862990810359231 - nodes in this community are weakly interconnected._