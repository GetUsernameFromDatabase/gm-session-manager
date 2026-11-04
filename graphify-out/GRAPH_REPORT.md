# Graph Report - gm-session-manager  (2026-11-04)

## Corpus Check
- 248 files · ~496,082 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3593 nodes · 14572 edges · 129 communities (106 shown, 23 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1459 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `354d5282`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- .do
- createProvider
- .encounterView
- SheetView
- combobox.js
- redirect.go
- testing.T
- github.com/gin-gonic/gin.Context
- select.js
- command.go
- .Entry
- Sheet
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Defenses
- wikiText
- wikilink.js
- dialog_templ.go
- popover.js
- time.Time
- Project Rules (CLAUDE.md, README)
- ParseSheet
- Campaign
- ImportView
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- Campaign
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
- who
- github.com/a-h/templ.Component
- .Plan
- IfElse
- Session
- github.com/a-h/templ.Attributes
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- ImageResults
- golangci-lint Config
- NewSheet
- Badge
- radiogroup.js
- vendorjs Tool
- levelUpForm
- Feature
- sqlc Config
- Hub
- Browser check
- Claude Settings
- .blockFor
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
- checkbox.js
- ImportView
- switch.js
- TestLiveUpdatesAndKick
- tabs.js
- portal.js
- useClick
- .get
- useFocus
- mark_others.js
- roundTrip
- Lint and Release Rules
- Run Script
- .builderPicks
- useTypeahead
- .loadLive
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- .ListItemsByCategory
- App
- NameDesc
- Entry
- Import
- blockMenu
- Client
- ParseWikidot
- State
- .wizardCommand
- radiogroup_templ.go
- New
- .Router
- Pending
- build.sh
- .sendFeedback
- RosterCharacter
- SearchReadableBlocksRow
- open5e.go
- .ListCompendiumByKind
- .ListCompendiumBySlugs
- capReader

## God Nodes (most connected - your core abstractions)
1. `App` - 266 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Sheet` - 89 edges
5. `newTestEnv()` - 86 edges
6. `SheetView` - 52 edges
7. `IfElse()` - 50 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 37 edges

## Surprising Connections (you probably didn't know these)
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Next up (pick from these)` --references--> `TestSpellNeed()`  [INFERRED]
  HANDOFF.md → internal/web/spells_test.go
- `Next up (pick from these)` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go
- `Where things live (creation and level-up)` --references--> `choiceGroup`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Where things live (creation and level-up)` --references--> `readPicks()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (129 total, 23 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (43): contextKey, clientMessage, TestMain(), newValidator(), Toast(), toastIcon(), TestMain(), lastTopLevelColon() (+35 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.14
Nodes (23): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+15 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (66): keepAwake(), keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Note (+58 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - ".encounterView"
Cohesion: 0.07
Nodes (20): Combatant, Condition, Encounter, ListCombatantsRow, Character, combatantDefenses(), initiativeMod(), outOfFight() (+12 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (85): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+77 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.06
Nodes (65): TestCanCharacter(), TestCanManageBlock(), TestCanReadAttachment(), TestCanSeeCombatantHP(), nullID(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName() (+57 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (20): ClearCookie(), SetCookie(), writeCookie(), Entity, CanCharacter(), coinCommand(), commandText(), hpCommand() (+12 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.07
Nodes (63): Amount, CondOpKind, Effects, Env, HelpEntry, parser, Completions(), condCommand() (+55 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (28): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+20 more)

### Community 14 - "Sheet"
Cohesion: 0.09
Nodes (13): DeathSaves, TestCustomMods(), Sum(), TestGearCommands(), Coins, Item, Sheet, WornArmor (+5 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.12
Nodes (36): VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock(), CanManageNote() (+28 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Defenses"
Cohesion: 0.20
Nodes (8): Handoff, Improvements (bigger, deferred), Next up (pick from these), Notes, Test data in data-test (all safe to delete), Working on the test instance, ParseDefenses(), TestSpellbook()

### Community 21 - "wikiText"
Cohesion: 0.31
Nodes (7): wikiLink(), wikiText(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.06
Nodes (66): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+58 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "time.Time"
Cohesion: 0.08
Nodes (18): Identity, FromContext(), CompendiumSource, browserName(), buildReview(), formValues(), deviceKey(), keepsBody() (+10 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.16
Nodes (20): TestUpgradeSheet(), ParseSheet(), TestLevelUpFeats(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestUpgradeSheets() (+12 more)

### Community 28 - "Campaign"
Cohesion: 0.28
Nodes (6): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks()

### Community 29 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows(), importStatusItems() (+2 more)

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
Nodes (42): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, GetImageRow, ListImagesRow, Attachment, CreateCombatantParams, applyCondOps() (+34 more)

### Community 34 - "Campaign"
Cohesion: 0.13
Nodes (11): Campaign, deviceOS(), formIntQuery(), isPhone(), lanHostname(), shortHost(), uaPick(), userAgentSummary() (+3 more)

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
Cohesion: 0.10
Nodes (30): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply() (+22 more)

### Community 41 - ".levelUp"
Cohesion: 0.09
Nodes (36): Where things live (creation and level-up), Options(), ClassFeature, ClassInfo, LevelUpPlan, classFeatures(), classInfo(), errorField() (+28 more)

### Community 42 - "Slugify"
Cohesion: 0.16
Nodes (18): Cantrips(), firstTable(), LevelSpells(), TestOptions(), TextData, Slugify(), Spell, textMarkdown() (+10 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "who"
Cohesion: 0.07
Nodes (13): DecodeData(), Character, CompendiumEntry, compendiumWeapon(), compendiumBack(), viewerCampaign(), HomebrewView, nextStep() (+5 more)

### Community 46 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (103): ctxState, ListMonstersByNameRow, AspectRatio(), Checkbox(), Chip(), Chips(), ChipsInput(), Combobox() (+95 more)

### Community 47 - ".Plan"
Cohesion: 0.17
Nodes (8): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), hpGain(), raise(), CasterLevel(), SlotMaxima()

### Community 48 - "IfElse"
Cohesion: 0.07
Nodes (72): Input(), Scripts(), Spinner(), Content(), contentID(), IDFromContext(), initialValue(), List() (+64 more)

### Community 49 - "Session"
Cohesion: 0.18
Nodes (9): tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken() (+1 more)

### Community 50 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (125): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+117 more)

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
Cohesion: 0.13
Nodes (12): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+4 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ImageResults"
Cohesion: 0.07
Nodes (47): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent() (+39 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "NewSheet"
Cohesion: 0.12
Nodes (20): Parse(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves(), TestLanguages(), TestLevelUp() (+12 more)

### Community 59 - "Badge"
Cohesion: 0.06
Nodes (57): Props, Variant, SpellData, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), alignmentColor() (+49 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "levelUpForm"
Cohesion: 0.17
Nodes (23): Label(), spellChecklist(), checkField(), abilityForm(), backgroundAbilities(), choiceFieldset(), choiceName(), choiceOptions() (+15 more)

### Community 63 - "Feature"
Cohesion: 0.15
Nodes (12): FeatureEffects, armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), UsesLabel(), CustomMod, Feature (+4 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".blockFor"
Cohesion: 0.22
Nodes (3): GetBlockRow, blockAccessLevel, stepsPayload

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
Cohesion: 0.04
Nodes (75): Description(), Field(), Label(), Content(), initialValue(), Item(), itemLabel(), Select() (+67 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "TestLiveUpdatesAndKick"
Cohesion: 0.32
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

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

### Community 83 - "roundTrip"
Cohesion: 0.24
Nodes (11): CoreTraits(), TestCoreTraits(), InputForEntry(), TestMonsterFormRoundTrip(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSpecies() (+3 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".loadLive"
Cohesion: 0.16
Nodes (8): Fields, ParseFields(), saveFieldLinks(), notFoundIsNil(), parseLiveKey(), toEntity(), Roster, liveKey

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.18
Nodes (17): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 103 - "Site"
Cohesion: 0.12
Nodes (16): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+8 more)

### Community 105 - "App"
Cohesion: 0.07
Nodes (9): ListAttachmentFilesRow, imageItem(), joinPath(), DisplayPage(), JoinQRSrc(), Flash(), LinkItem, SwitcherResults() (+1 more)

### Community 106 - "NameDesc"
Cohesion: 0.22
Nodes (12): OptionPrerequisite(), NameDesc, catchUpResistances(), TestCatchUpResistances(), UpgradeSheets(), upgradeSpecies(), addResistance(), applyTraitText() (+4 more)

### Community 107 - "Entry"
Cohesion: 0.21
Nodes (12): Entry, fixtures, normalizer, actionRank(), byPK(), children(), formatCR(), intField() (+4 more)

### Community 108 - "Import"
Cohesion: 0.16
Nodes (20): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+12 more)

### Community 109 - "blockMenu"
Cohesion: 0.20
Nodes (16): localTime(), addBlockForm(), authorRef(), blockMenu(), Author, NoteItem, NoteView, moveButton() (+8 more)

### Community 110 - "Client"
Cohesion: 0.23
Nodes (3): message, Client, topicIDs()

### Community 111 - "ParseWikidot"
Cohesion: 0.17
Nodes (13): block, ClassFeature, ClassData, partsText(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+5 more)

### Community 112 - "State"
Cohesion: 0.10
Nodes (13): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+5 more)

### Community 113 - ".wizardCommand"
Cohesion: 0.12
Nodes (21): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+13 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 116 - ".Router"
Cohesion: 0.16
Nodes (7): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 117 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 129 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 133 - "open5e.go"
Cohesion: 0.08
Nodes (28): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+20 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 334 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **23 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `.encounterView`, `github.com/gin-gonic/gin.Context`, `perm.go`, `time.Time`, `Campaign`, `database/sql.NullInt64`, `Campaign`, `Store`, `.levelUp`, `Slugify`, `who`, `Hub`, `.blockFor`, `.builderPicks`, `.loadLive`, `State`, `.wizardCommand`, `.Router`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10259357500736811 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Weapon`, `.encounterView`, `SheetView`, `.levelUp`, `Slugify`, `NameDesc`, `command.go`, `who`, `.Plan`, `.wizardCommand`, `dialog_templ.go`, `NewSheet`, `ParseSheet`, `Feature`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `popover_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13768115942028986 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `.levelUp` to `compendium_templ.go`, `NameDesc`, `Feature`, `.Plan`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011692126909518213 - nodes in this community are weakly interconnected._