# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 248 files · ~498,562 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3603 nodes · 14661 edges · 136 communities (109 shown, 27 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1464 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `b715db5b`
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
- Session
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Page
- ImageResults
- wikilink.js
- dialog_templ.go
- popover.js
- wikiText
- Project Rules (CLAUDE.md, README)
- ParseSheet
- .fail
- Sheet
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- deviceOS
- command.js
- dialog.js
- fighter
- golang.org/x/net/html.Node
- tooltip.js
- Open
- .levelUp
- Slugify
- export/markdown.go
- composite.js
- App
- blockMenu
- .Plan
- IfElse
- .whoami
- CN
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- .actions
- golangci-lint Config
- Slugify
- github.com/a-h/templ.Component
- radiogroup.js
- vendorjs Tool
- select_templ.go
- Handoff
- sqlc Config
- .joinInfo
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
- parseLiveKey
- Lint and Release Rules
- Run Script
- .playerGroups
- useTypeahead
- rowFigure
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Topic
- Site
- .ListItemsByCategory
- parseRowSet
- fingerprintPairs
- Entry
- Import
- Client
- wikidotClass
- message
- radiogroup_templ.go
- Hub
- build.sh
- errBodyWriter
- Props
- .sendFeedback
- New
- SearchReadableBlocksRow
- Pending
- ArmorData
- buildReview
- Campaign
- DBTX
- toast_templ.go
- tooltip_templ.go
- Entity
- RosterCharacter
- NoteBlock
- lanHostname
- .ListCompendiumByKind
- .ListSpellsForClass

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
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Next up (pick from these)` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go
- `Where things live (creation and level-up)` --references--> `taughtPicked()`  [INFERRED]
  HANDOFF.md → internal/web/views/levelup_templ.go
- `Next up (pick from these)` --references--> `swapFeatSpell()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Where things live (creation and level-up)` --references--> `optionKind`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (136 total, 27 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (53): contextKey, looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes, o5eFeature (+45 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.15
Nodes (22): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+14 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (81): keepAwake(), keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow (+73 more)

### Community 3 - ".do"
Cohesion: 0.17
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "formInt"
Cohesion: 0.14
Nodes (9): Combatant, Entity, Character, entityTopics(), linkEntityLook(), formInt(), lookFree(), Toast() (+1 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (82): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+74 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (69): Completions(), TestCompletions(), TestDefenses(), TestLanguages(), TestLevelUp(), TestLevelUpSubclassAndPools(), TestArmorAndFeatureUses(), TestExplainSums() (+61 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (19): ClearCookie(), SetCookie(), writeCookie(), ftsQuery(), Search(), Campaign, coinCommand(), commandText() (+11 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.05
Nodes (75): Amount, CondOpKind, Effects, Env, FeatureEffects, HelpEntry, parser, condCommand() (+67 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (31): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+23 more)

### Community 14 - "Session"
Cohesion: 0.10
Nodes (19): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+11 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.08
Nodes (47): Note, ListCombatantsRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+39 more)

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
Cohesion: 0.07
Nodes (57): Label(), Scripts(), Separator(), campaignLink(), PendingGroup, PendingRequests(), CharacterPage(), homebrewAction() (+49 more)

### Community 21 - "ImageResults"
Cohesion: 0.08
Nodes (43): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent() (+35 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "dialog_templ.go"
Cohesion: 0.09
Nodes (40): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+32 more)

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
Cohesion: 0.13
Nodes (28): TestUpgradeSheet(), ParseSheet(), TestLevelUpFeats(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestSpellCounts() (+20 more)

### Community 28 - ".fail"
Cohesion: 0.09
Nodes (9): Character, CanCharacter(), compendiumBack(), linkCharacterLook(), Roster, HomebrewView, nextStep(), wizardPath() (+1 more)

### Community 29 - "Sheet"
Cohesion: 0.07
Nodes (19): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), skillByKey(), Sum() (+11 more)

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
Nodes (40): Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, CompendiumEntry, CompendiumSource, Condition, Encounter (+32 more)

### Community 34 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "fighter"
Cohesion: 0.22
Nodes (7): fighter(), TestCommands(), TestHealingFromZeroResetsDeathSaves(), TestParseRawArgs(), geared(), TestExplainDamage(), TestStowKeepsMods()

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Open"
Cohesion: 0.12
Nodes (19): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+11 more)

### Community 41 - ".levelUp"
Cohesion: 0.09
Nodes (36): Where things live (creation and level-up), Spell, addClassSpells(), classInfo(), errorField(), spellNeed(), TestSpellNeed(), LevelUpView (+28 more)

### Community 42 - "Slugify"
Cohesion: 0.10
Nodes (25): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestCoreTraits() (+17 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "App"
Cohesion: 0.07
Nodes (11): DecodeData(), Attachment, GetBlockRow, compendiumWeapon(), firstLine(), mdLinkText(), saveScreenshot(), saveSniffed() (+3 more)

### Community 46 - "blockMenu"
Cohesion: 0.20
Nodes (16): localTime(), addBlockForm(), authorRef(), blockMenu(), Author, NoteItem, NoteView, moveButton() (+8 more)

### Community 47 - ".Plan"
Cohesion: 0.10
Nodes (16): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+8 more)

### Community 48 - "IfElse"
Cohesion: 0.14
Nodes (36): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+28 more)

### Community 49 - ".whoami"
Cohesion: 0.10
Nodes (12): IsLoopback(), TestIsLoopback(), Identity, FromContext(), RequireAdmin(), browserName(), deviceKey(), keepsBody() (+4 more)

### Community 50 - "CN"
Cohesion: 0.07
Nodes (59): Props, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps (+51 more)

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

### Community 56 - ".actions"
Cohesion: 0.33
Nodes (5): castingKind(), kindOrder(), weaponReach(), TestFreeCastKind(), actionEntry

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Slugify"
Cohesion: 0.12
Nodes (25): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, Slugify() (+17 more)

### Community 59 - "github.com/a-h/templ.Component"
Cohesion: 0.07
Nodes (62): SpellData, CountHabitatsRow, SearchCompendiumRow, Badge(), Item(), alignmentColor(), classBadges(), classBody() (+54 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "select_templ.go"
Cohesion: 0.05
Nodes (52): Defenses, damageChoice(), Content(), Group(), initialValue(), Item(), itemLabel(), Label() (+44 more)

### Community 63 - "Handoff"
Cohesion: 0.20
Nodes (9): Defense, Handoff, Improvements (bigger, deferred), Lessons (read before orchestrating), Next up (pick from these), Notes, Test data in data-test (all safe to delete), Working on the test instance (+1 more)

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
Nodes (116): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+108 more)

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
Nodes (90): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+82 more)

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
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 86 - ".playerGroups"
Cohesion: 0.50
Nodes (3): nameKey(), TestPlayerGroups(), combatantGroup

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.18
Nodes (17): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Topic"
Cohesion: 0.12
Nodes (10): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+2 more)

### Community 103 - "Site"
Cohesion: 0.12
Nodes (17): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+9 more)

### Community 107 - "Entry"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children() (+11 more)

### Community 108 - "Import"
Cohesion: 0.14
Nodes (24): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+16 more)

### Community 111 - "wikidotClass"
Cohesion: 0.14
Nodes (20): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), sourceLine() (+12 more)

### Community 112 - "message"
Cohesion: 0.33
Nodes (4): message, topicIDs(), Diff(), TestDiff()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 117 - "Hub"
Cohesion: 0.23
Nodes (3): ClientInfo, Hub, TestCloseSessionAndCharacter()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 128 - "New"
Cohesion: 0.33
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests(), TestPendingFlow() (+2 more)

### Community 132 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

### Community 134 - "buildReview"
Cohesion: 0.15
Nodes (8): buildReview(), csvSource(), formValues(), parseCSV(), TestGuessMapping(), capReader, csvImport, csvImports

### Community 135 - "Campaign"
Cohesion: 0.28
Nodes (6): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks()

### Community 140 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 141 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 144 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 145 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 148 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 335 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **27 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.do`, `formInt`, `buildReview`, `Campaign`, `github.com/gin-gonic/gin.Context`, `perm.go`, `.fail`, `database/sql.NullInt64`, `Open`, `.levelUp`, `Slugify`, `.whoami`, `CN`, `LoadOrCreateCert`, `Slugify`, `.joinInfo`, `parseLiveKey`, `Topic`, `Hub`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09811372711000964 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `formInt`, `SheetView`, `testing.T`, `Slugify`, `.levelUp`, `command.go`, `App`, `.Plan`, `perm.go`, `Slugify`, `ParseSheet`, `select_templ.go`?**
  _High betweenness centrality (0.022) - this node is a cross-community bridge._
- **Should `popover_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14624505928853754 - nodes in this community are weakly interconnected._
- **Why does `ClassInfo` connect `.Plan` to `compendium_templ.go`, `.levelUp`, `Slugify`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.009253597488891607 - nodes in this community are weakly interconnected._