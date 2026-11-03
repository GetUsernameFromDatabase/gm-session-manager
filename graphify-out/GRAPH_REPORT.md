# Graph Report - gm-session-manager  (2026-11-03)

## Corpus Check
- 246 files · ~486,253 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3530 nodes · 14088 edges · 134 communities (113 shown, 21 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1400 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `80ad5dc2`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- levelUpForm
- context.Context
- .do
- createProvider
- database/sql.NullInt64
- SheetView
- combobox.js
- redirect.go
- newTestEnv
- App
- select.js
- command.go
- .Entry
- hpDialog
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Part
- wikiText
- wikilink.js
- CombatantView
- popover.js
- Character
- Project Rules (CLAUDE.md, README)
- State
- field_templ.go
- Page
- app.js
- tabbable.js
- useDismiss
- parseCommand
- github.com/a-h/templ.Attributes
- command.js
- dialog.js
- Sheet
- wikidot.go
- tooltip.js
- Store
- Topic
- Slugify
- export/markdown.go
- composite.js
- .sendFeedback
- github.com/a-h/templ.Component
- .Plan
- monsterBody
- Session
- Client
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- toast_templ.go
- golangci-lint Config
- testing.T
- combobox_templ.go
- radiogroup.js
- vendorjs Tool
- IfElse
- New
- sqlc Config
- newStore
- Browser check
- Claude Settings
- Pending
- use_transition_status.js
- playwright
- release.sh
- Input
- checkbox.js
- Feature
- switch.js
- nullID
- tabs.js
- portal.js
- useClick
- deviceOS
- useFocus
- mark_others.js
- .Router
- Lint and Release Rules
- Run Script
- findItem
- useTypeahead
- .device
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- popover_templ.go
- .ListCompendiumByKind
- switch_templ.go
- Entry
- config.go
- inputgroup_templ.go
- ImportView
- roundTrip
- Hub
- Verbs
- radiogroup_templ.go
- command_templ.go
- build.sh
- HomebrewInputFor
- Button
- tooltip_templ.go
- time.Time
- Handoff: open items from the 2026-11-03 review
- TestLiveUpdatesAndKick
- alert_templ.go
- NoteListView
- lanHostname
- Props
- open5e.go
- generateSVG
- capReader
- spellCounts

## God Nodes (most connected - your core abstractions)
1. `App` - 263 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `newTestEnv()` - 82 edges
5. `Sheet` - 80 edges
6. `SheetView` - 50 edges
7. `IfElse()` - 48 edges
8. `Input()` - 46 edges
9. `Command` - 43 edges
10. `Client` - 35 edges

## Surprising Connections (you probably didn't know these)
- `Bugs fixed since` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Bugs fixed since` --references--> `OptionPrerequisite()`  [INFERRED]
  HANDOFF.md → internal/compendium/coretraits.go
- `Next up (pick from these)` --references--> `castingKind()`  [INFERRED]
  HANDOFF.md → internal/web/views/character_templ.go
- `Bugs fixed since` --references--> `meetsPrerequisite()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (134 total, 21 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (42): contextKey, keepAwake(), clientMessage, TestMain(), newValidator(), TestMain(), lastTopLevelColon(), twImportantToV3() (+34 more)

### Community 1 - "levelUpForm"
Cohesion: 0.12
Nodes (35): Label(), homebrewAction(), HomebrewPage(), inlineField(), Notice(), abilitySelect(), LevelUpView, Option (+27 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (68): keepAwake(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow, Encounter, Note (+60 more)

### Community 3 - ".do"
Cohesion: 0.16
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (45): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, GetImageRow, Attachment, Entity, CreateCombatantParams, applyCondOps() (+37 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (85): ArmorEffect(), damageLabel(), Item, Weapon, WornArmor, Signed(), Field, Action() (+77 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "newTestEnv"
Cohesion: 0.07
Nodes (44): Import(), ParseSheet(), TestBuilderCompendium(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestConditionExpiresOnTurn(), TestDeleteRestoreImport() (+36 more)

### Community 10 - "App"
Cohesion: 0.05
Nodes (24): SetCookie(), CanCharacter(), compendiumBack(), viewerCampaign(), entityTopics(), formIntQuery(), isHTMX(), isPhone() (+16 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.10
Nodes (47): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+39 more)

### Community 13 - ".Entry"
Cohesion: 0.17
Nodes (12): HomebrewOwner, ValidKind(), HomebrewInput, SaveHomebrew(), parseActions(), partsText(), FeaturesText(), paragraphs() (+4 more)

### Community 14 - "hpDialog"
Cohesion: 0.18
Nodes (30): Dialog(), Close(), Content(), Description(), Dialog(), Footer(), Header(), Portal() (+22 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (46): GetBlockRow, ListCombatantsRow, Entity, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+38 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Part"
Cohesion: 0.14
Nodes (6): Part, Skill, TestModifier(), skillByKey(), Modifier(), ProficiencyAt()

### Community 21 - "wikiText"
Cohesion: 0.31
Nodes (7): wikiLink(), wikiText(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.10
Nodes (36): Condition, beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), combatCondition(), conditionLabel() (+28 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "Character"
Cohesion: 0.09
Nodes (14): Campaign, Character, Combatant, Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile() (+6 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "State"
Cohesion: 0.13
Nodes (11): fakeHandler, patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff(), deviceLabel() (+3 more)

### Community 28 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 29 - "Page"
Cohesion: 0.07
Nodes (58): MediaProps, MediaVariant, Content(), Description(), Empty(), Header(), Media(), mediaClasses() (+50 more)

### Community 30 - "app.js"
Cohesion: 0.11
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "parseCommand"
Cohesion: 0.25
Nodes (9): findMod(), newMod(), parseMod(), IsModTarget(), featureCommand(), coinCommand(), commandText(), hpCommand() (+1 more)

### Community 34 - "github.com/a-h/templ.Attributes"
Cohesion: 0.05
Nodes (50): Props, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size (+42 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Sheet"
Cohesion: 0.08
Nodes (16): DeathSaves, Completions(), parseItem(), Sum(), Coins, Sheet, Slot, SizeNamed() (+8 more)

### Community 38 - "wikidot.go"
Cohesion: 0.16
Nodes (28): block, ClassFeature, textLines(), ClassData, ParsePlainPage(), attr(), byID(), cellAt() (+20 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.12
Nodes (25): migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open() (+17 more)

### Community 41 - "Topic"
Cohesion: 0.15
Nodes (7): event, Handler, stateHandler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 42 - "Slugify"
Cohesion: 0.05
Nodes (59): EquipmentItem, Bugs fixed since, Cantrips(), EquipmentOptions(), firstTable(), EquipmentOption, LevelSpells(), OptionPrerequisite() (+51 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.06
Nodes (42): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+34 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".sendFeedback"
Cohesion: 0.33
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 46 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (87): ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer(), Checkbox() (+79 more)

### Community 47 - ".Plan"
Cohesion: 0.13
Nodes (11): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan, hpGain() (+3 more)

### Community 48 - "monsterBody"
Cohesion: 0.09
Nodes (45): Badge(), IfElse(), inFightSuffix(), alignmentColor(), classBadges(), classBody(), commaPrefixed(), compendiumRow() (+37 more)

### Community 49 - "Session"
Cohesion: 0.11
Nodes (15): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity() (+7 more)

### Community 50 - "Client"
Cohesion: 0.23
Nodes (3): message, Client, topicIDs()

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

### Community 56 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "testing.T"
Cohesion: 0.06
Nodes (45): Parse(), fighter(), TestCommands(), TestCompletions(), TestCustomMods(), TestHealingFromZeroResetsDeathSaves(), TestLanguages(), TestLevelUp() (+37 more)

### Community 59 - "combobox_templ.go"
Cohesion: 0.09
Nodes (42): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+34 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "IfElse"
Cohesion: 0.08
Nodes (50): CountHabitatsRow, SearchCompendiumRow, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant() (+42 more)

### Community 63 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "newStore"
Cohesion: 0.17
Nodes (17): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+9 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Input"
Cohesion: 0.07
Nodes (61): CloseFor(), Description(), Field(), Group(), Label(), Input(), initialValue(), Select() (+53 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "Feature"
Cohesion: 0.15
Nodes (13): FeatureEffects, parseFeatureSet(), parseUses(), armorCondition(), columnInts(), ParseEffects(), SaveFeature(), TestParseEffects() (+5 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "nullID"
Cohesion: 0.12
Nodes (20): Player(), nullID(), TestCombatantNamesFollowSource(), TestMonsterPortrait(), testEnv, TestAttachmentAccess(), TestAttachmentToImage(), TestFilesCleanup() (+12 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - ".Router"
Cohesion: 0.16
Nodes (7): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 86 - "findItem"
Cohesion: 0.38
Nodes (3): findItem(), parseArmor(), ensureItem()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".device"
Cohesion: 0.47
Nodes (4): testEnv, TestRequestLog(), TestRequestLogReusedID(), TestRequestSent()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.18
Nodes (17): ftEntry, ftMonster, ftSpellcasting, ftEntries(), ftEntry1(), ftList(), ftSpells(), ftType() (+9 more)

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
Cohesion: 0.21
Nodes (11): state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps (+3 more)

### Community 106 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

### Community 107 - "Entry"
Cohesion: 0.27
Nodes (9): Entry, fixtures, normalizer, byPK(), children(), joinText(), label(), parseFixtures() (+1 more)

### Community 108 - "config.go"
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 109 - "inputgroup_templ.go"
Cohesion: 0.16
Nodes (14): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+6 more)

### Community 110 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), ImportReviewPage(), importRows() (+2 more)

### Community 111 - "roundTrip"
Cohesion: 0.24
Nodes (10): CoreTraits(), TestCoreTraits(), InputForEntry(), TestMonsterFormRoundTrip(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSubclass() (+2 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 116 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "HomebrewInputFor"
Cohesion: 0.11
Nodes (21): MonsterAction, abilityMod(), addEntry(), ParseAidedd(), setStat(), statLabel(), MonsterData, firstInt() (+13 more)

### Community 123 - "Button"
Cohesion: 0.09
Nodes (43): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Look, Button(), hpQuickForm(), initials(), selectItems() (+35 more)

### Community 124 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 125 - "time.Time"
Cohesion: 0.06
Nodes (21): ClearCookie(), writeCookie(), Identity, FromContext(), ListAttachmentFilesRow, ListImagesRow, browserName(), firstLine() (+13 more)

### Community 126 - "Handoff: open items from the 2026-11-03 review"
Cohesion: 0.22
Nodes (8): Character page, Handoff: open items from the 2026-11-03 review, Improvements (bigger, deferred), Next up (pick from these), Notes, Other pages, Test data in data-test, castingKind()

### Community 127 - "TestLiveUpdatesAndKick"
Cohesion: 0.32
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 128 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 131 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 132 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 133 - "open5e.go"
Cohesion: 0.07
Nodes (30): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+22 more)

### Community 137 - "generateSVG"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 144 - "spellCounts"
Cohesion: 0.67
Nodes (3): sortedColumns(), spellCounts(), TestSpellCounts()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **21 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `database/sql.NullInt64`, `Topic`, `Slugify`, `Hub`, `perm.go`, `Page`, `.Router`, `LoadOrCreateCert`, `Character`, `State`, `time.Time`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09502822905864731 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `levelUpForm`, `SheetView`, `newTestEnv`, `Feature`, `Slugify`, `.Plan`, `perm.go`, `Part`, `findItem`, `Character`, `testing.T`?**
  _High betweenness centrality (0.024) - this node is a cross-community bridge._
- **Should `levelUpForm` be split into smaller, more focused modules?**
  _Cohesion score 0.12162162162162163 - nodes in this community are weakly interconnected._
- **Why does `Client` connect `Client` to `compendium_templ.go`, `Pending`, `redirect.go`, `Topic`, `Hub`, `perm.go`, `State`, `time.Time`, `New`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011459968602825745 - nodes in this community are weakly interconnected._