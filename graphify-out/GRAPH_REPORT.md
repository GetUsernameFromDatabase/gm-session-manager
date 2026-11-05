# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 278 files · ~521,685 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3791 nodes · 15824 edges · 129 communities (108 shown, 21 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1585 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f4985a8e`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- FromContext
- context.Context
- .do
- createProvider
- CombatantView
- SheetView
- combobox.js
- redirect.go
- testing.T
- dialog_templ.go
- select.js
- Command
- .Entry
- ImportView
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- empty_templ.go
- levelUpForm
- wikilink.js
- github.com/a-h/templ.Component
- popover.js
- wikiText
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- DecodeData
- Topic
- command.js
- dialog.js
- database/sql.NullInt64
- golang.org/x/net/html.Node
- tooltip.js
- open5e.go
- .levelUp
- capReader
- export/markdown.go
- composite.js
- .render
- Client
- ClassInfo
- tabs_templ.go
- Attachment
- New
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- select_templ.go
- golangci-lint Config
- Slugify
- Book
- radiogroup.js
- vendorjs Tool
- radiogroup_templ.go
- Pending
- sqlc Config
- Queries
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- builder_templ.go
- checkbox.js
- cell
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
- github.com/a-h/templ.Component
- useTypeahead
- NameDesc
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Hub
- Site
- ftItemData
- popover_templ.go
- .wizardCommand
- ftBook
- Import
- ParseWikidot
- .fiveToolsImport
- State
- toast_templ.go
- .publishLocked
- RosterCharacter
- Weapon
- .joinInfo
- .importUpload
- build.sh
- spellCounts
- Sheet
- errBodyWriter
- .sendFeedback
- Session
- open5e.go
- Import

## God Nodes (most connected - your core abstractions)
1. `App` - 276 edges
2. `CN()` - 112 edges
3. `Button()` - 107 edges
4. `newTestEnv()` - 95 edges
5. `Sheet` - 90 edges
6. `SheetView` - 52 edges
7. `IfElse()` - 51 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 36 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Fingerprint()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/join/cert.go
- `run()` --calls--> `LoadOrCreateCert()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/join/cert.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (129 total, 21 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.12
Nodes (38): contextKey, keepAwake(), classRef, ftBookInfo, indexed, sourceOf, ASIError, clientMessage (+30 more)

### Community 1 - "FromContext"
Cohesion: 0.07
Nodes (19): ClearCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), mdLinkText(), formIntQuery() (+11 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (78): keepAwake(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, SearchReadableBlocksRow, CreateCombatantParams (+70 more)

### Community 3 - ".do"
Cohesion: 0.12
Nodes (20): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+12 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "CombatantView"
Cohesion: 0.09
Nodes (38): Encounter, notFoundIsNil(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), currentID() (+30 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (88): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+80 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (53): Defense, feature(), fiveToolsClasses(), TestFiveToolsClasses(), TestFiveToolsClassLevelUp(), fighter(), TestCommands(), TestCompletions() (+45 more)

### Community 10 - "dialog_templ.go"
Cohesion: 0.24
Nodes (26): ctxKey, ctxState, Props, Dialog(), Close(), CloseFor(), Content(), Description() (+18 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (71): Amount, CondOpKind, Effects, Env, HelpEntry, parser, Completions(), condCommand() (+63 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (23): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInputFor(), abilityKey(), actionsText() (+15 more)

### Community 14 - "ImportView"
Cohesion: 0.09
Nodes (29): Props, ImageCategory, Checkbox(), initialChecked(), categoryRow(), columnLabel(), columnOverview(), ImportCategory (+21 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (41): Note, GetBlockRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+33 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "empty_templ.go"
Cohesion: 0.13
Nodes (34): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListMonstersByNameRow (+26 more)

### Community 21 - "levelUpForm"
Cohesion: 0.12
Nodes (30): Footer(), selectItems(), homebrewLinks(), abilitySelect(), levelUpForm(), optionItems(), slotSummary(), spellChecklist() (+22 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (97): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, Description(), Kbd(), Label(), Content() (+89 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "wikiText"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.18
Nodes (23): ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved() (+15 more)

### Community 28 - "App"
Cohesion: 0.06
Nodes (16): SetCookie(), CanCharacter(), coinCommand(), commandText(), entityTopics(), formInt(), nullID(), redirect() (+8 more)

### Community 29 - "newTestEnv"
Cohesion: 0.07
Nodes (40): nullID(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait(), TestDeleteRestoreImport() (+32 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "DecodeData"
Cohesion: 0.07
Nodes (14): DecodeData(), ListEntryCopiesRow, Character, CompendiumEntry, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile() (+6 more)

### Community 34 - "Topic"
Cohesion: 0.15
Nodes (7): event, Handler, stateHandler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (46): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, ListImagesRow, Combatant, Condition, ListCombatantsRow, Character (+38 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.19
Nodes (20): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+12 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "open5e.go"
Cohesion: 0.08
Nodes (28): fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature, o5eDescribes (+20 more)

### Community 41 - ".levelUp"
Cohesion: 0.10
Nodes (32): Options(), ClassFeature, ClassInfo, LevelUpPlan, classFeatures(), classInfo(), errorField(), spellbook() (+24 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".render"
Cohesion: 0.10
Nodes (13): GetImageRow, Campaign, compendiumBack(), viewerCampaign(), isHTMX(), imageItem(), imageVersion(), Roster (+5 more)

### Community 47 - "ClassInfo"
Cohesion: 0.17
Nodes (8): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), hpGain(), raise(), CasterLevel(), SlotMaxima()

### Community 48 - "tabs_templ.go"
Cohesion: 0.13
Nodes (38): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+30 more)

### Community 49 - "Attachment"
Cohesion: 0.20
Nodes (6): Attachment, saveScreenshot(), parseCSV(), TestGuessMapping(), saveSniffed(), CreateAttachmentParams

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
Cohesion: 0.13
Nodes (12): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+4 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "select_templ.go"
Cohesion: 0.11
Nodes (23): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+15 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Slugify"
Cohesion: 0.15
Nodes (17): Cantrips(), firstTable(), LevelSpells(), TestOptions(), TextData, Slugify(), Spell, textMarkdown() (+9 more)

### Community 59 - "Book"
Cohesion: 0.12
Nodes (26): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools(), fetchBooks() (+18 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "radiogroup_templ.go"
Cohesion: 0.17
Nodes (15): attrString(), generateSVG(), getIconContent(), Props, Icon(), initialValue(), Item(), RadioGroup() (+7 more)

### Community 63 - "Pending"
Cohesion: 0.25
Nodes (4): PendingState, Pending, PendingKind, viewerCtx

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "Queries"
Cohesion: 0.16
Nodes (13): HomebrewOwner, HomebrewInput, SaveHomebrew(), Queries, Entity, applyCondOps(), uniqueName(), linkEntityLook() (+5 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (177): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+169 more)

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
Nodes (73): FeatureEffects, UsesLabel(), Coins, CustomMod, Defenses, Feature, Uses, Description() (+65 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "cell"
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
Cohesion: 0.18
Nodes (12): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+4 more)

### Community 86 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (56): Props, Variant, SpellData, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), alignmentColor() (+48 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "NameDesc"
Cohesion: 0.30
Nodes (9): OptionPrerequisite(), NameDesc, catchUpResistances(), TestCatchUpResistances(), UpgradeSheets(), upgradeSpecies(), isRepeatable(), resistsPick() (+1 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.11
Nodes (29): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+21 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "Site"
Cohesion: 0.14
Nodes (13): Site, dataFor(), Kind, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey() (+5 more)

### Community 104 - "ftItemData"
Cohesion: 0.13
Nodes (16): ArmorData, ftClassFeature, ftClassRefs, ftDistance, ftItem, ftAmount(), ftArmor(), ftCode() (+8 more)

### Community 105 - "popover_templ.go"
Cohesion: 0.20
Nodes (13): Description(), state(), Trigger(), triggerID(), Align, ContentProps, ctxKey, ctxState (+5 more)

### Community 106 - ".wizardCommand"
Cohesion: 0.08
Nodes (33): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestCoreTraits() (+25 more)

### Community 107 - "ftBook"
Cohesion: 0.17
Nodes (17): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+9 more)

### Community 108 - "Import"
Cohesion: 0.10
Nodes (32): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply() (+24 more)

### Community 109 - "ParseWikidot"
Cohesion: 0.27
Nodes (12): TestTokensReadAsLabels(), InputForEntry(), TestMonsterFormRoundTrip(), ParseWikidot(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSpecies() (+4 more)

### Community 110 - ".fiveToolsImport"
Cohesion: 0.24
Nodes (4): ImportSummary(), fiveToolsKey(), FiveToolsView, NewFiveToolsView()

### Community 111 - "State"
Cohesion: 0.14
Nodes (10): fakeHandler, message, patchMessage, viewPatch, topicIDs(), Diff(), State, manyKeys() (+2 more)

### Community 112 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 114 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 115 - "Weapon"
Cohesion: 0.11
Nodes (14): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+6 more)

### Community 117 - ".importUpload"
Cohesion: 0.24
Nodes (6): ensureCategory(), nullString(), buildReview(), formValues(), csvImport, csvImports

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 123 - "Sheet"
Cohesion: 0.09
Nodes (9): DeathSaves, TestCustomMods(), Sum(), TestUpgradeSheet(), Sheet, slotMaxima(), addResistance(), applyTraitText() (+1 more)

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 131 - "Session"
Cohesion: 0.09
Nodes (17): tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), TestCredentialsMatch(), TestIsLoopback(), TestRequireAdmin(), TestTokens() (+9 more)

### Community 133 - "open5e.go"
Cohesion: 0.20
Nodes (13): Entry, fixtures, normalizer, Source, parseFiveToolsBook(), parseFiveToolsBookReport(), TestFiveToolsEdition(), storeSource() (+5 more)

### Community 134 - "Import"
Cohesion: 0.10
Nodes (33): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+25 more)

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 343 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **21 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `FromContext`, `context.Context`, `.do`, `Session`, `CombatantView`, `perm.go`, `DecodeData`, `Topic`, `database/sql.NullInt64`, `.levelUp`, `.render`, `Attachment`, `Slugify`, `Pending`, `Hub`, `.wizardCommand`, `Import`, `.fiveToolsImport`, `State`, `.joinInfo`, `.importUpload`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11913302529182879 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `DecodeData`, `database/sql.NullInt64`, `SheetView`, `builder_templ.go`, `testing.T`, `.wizardCommand`, `.levelUp`, `Command`, `ClassInfo`, `Weapon`, `NameDesc`, `Slugify`, `ParseSheet`?**
  _High betweenness centrality (0.014) - this node is a cross-community bridge._
- **Should `FromContext` be split into smaller, more focused modules?**
  _Cohesion score 0.07017543859649122 - nodes in this community are weakly interconnected._
- **Why does `Command` connect `Command` to `DecodeData`, `.wizardCommand`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.01029113067027759 - nodes in this community are weakly interconnected._