# Graph Report - gm-session-manager  (2026-11-03)

## Corpus Check
- 246 files · ~478,269 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3500 nodes · 13916 edges · 130 communities (109 shown, 21 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1373 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `a76f8c39`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- levelUpForm
- context.Context
- IsLoopback
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
- github.com/a-h/templ.Attributes
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- .fail
- Character
- wikilink.js
- combat_templ.go
- popover.js
- monsterBody
- Project Rules (CLAUDE.md, README)
- time.Time
- Sheet
- Button
- app.js
- tabbable.js
- useDismiss
- Weapon
- ImportView
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
- CN
- .Plan
- select_templ.go
- Session
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- toast_templ.go
- golangci-lint Config
- Parse
- TestLiveUpdatesAndKick
- radiogroup.js
- vendorjs Tool
- IfElse
- App
- sqlc Config
- RosterCharacter
- Browser check
- Claude Settings
- .playerGroups
- use_transition_status.js
- playwright
- release.sh
- builderFields
- checkbox.js
- CustomMod
- switch.js
- nullID
- tabs.js
- portal.js
- useClick
- testing.T
- useFocus
- mark_others.js
- Lint and Release Rules
- Run Script
- useTypeahead
- dialog_templ.go
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
- open5e.go
- wikidotClass
- Client
- radiogroup_templ.go
- Hub
- tooltip_templ.go
- New
- build.sh
- buildReview
- wikiText
- Pending
- App
- Handoff: open items from the 2026-11-03 review
- IfElse
- rowFigure
- SearchReadableBlocksRow
- lanHostname
- spellCounts
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
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `run()` --calls--> `NewManager()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/auth/manager.go
- `run()` --calls--> `Import()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/compendium/importer.go
- `run()` --calls--> `Load()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/config/config.go
- `run()` --calls--> `Open()`  [EXTRACTED]
  cmd/gm-session-manager/main.go → internal/db/db.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (130 total, 21 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.09
Nodes (55): Props, contextKey, keepAwake(), clientMessage, TestMain(), newValidator(), initialChecked(), TestMain() (+47 more)

### Community 1 - "levelUpForm"
Cohesion: 0.16
Nodes (26): Footer(), Label(), abilitySelect(), LevelUpView, Option, levelUpForm(), LevelUpPage(), optionItems() (+18 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (63): keepAwake(), ftsQuery(), Search(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow (+55 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (43): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, GetImageRow, Condition, applyCondOps(), uniqueName(), linkEntityLook() (+35 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (78): ArmorEffect(), Slot, Signed(), Field, Variant, Action(), Card(), Content() (+70 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "newTestEnv"
Cohesion: 0.07
Nodes (44): TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestConditionExpiresOnTurn(), TestDeleteRestoreImport(), TestRenameDeleteCampaign(), testEnv, pendingSecret() (+36 more)

### Community 10 - ".loadLive"
Cohesion: 0.13
Nodes (12): ListCombatantsRow, initiativeMod(), outOfFight(), pcSheet(), stepTurn(), notFoundIsNil(), parseLiveKey(), combatantKind() (+4 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.09
Nodes (50): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+42 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+22 more)

### Community 14 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (105): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+97 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.21
Nodes (25): VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew(), CanManageBlock(), CanManageNote() (+17 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - ".fail"
Cohesion: 0.11
Nodes (7): CanCharacter(), compendiumBack(), HomebrewView, SetCharacterCurrentImageParams, ToggleCharacterImageRevealedParams, UnlinkCharacterImageParams, who

### Community 21 - "Character"
Cohesion: 0.06
Nodes (26): DecodeData(), Character, CompendiumEntry, CompendiumSource, ParseSheet(), Campaign(), CharacterFileName(), fileSlug() (+18 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "combat_templ.go"
Cohesion: 0.10
Nodes (35): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+27 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "monsterBody"
Cohesion: 0.14
Nodes (27): SpellData, classBody(), commaPrefixed(), entryBody(), EntryDetails(), entryHead(), EntryPage(), EntryPreview() (+19 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "time.Time"
Cohesion: 0.06
Nodes (22): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), ListImagesRow, browserName() (+14 more)

### Community 28 - "Sheet"
Cohesion: 0.07
Nodes (17): DeathSaves, TestCustomMods(), TestResistances(), TestExplainSums(), Sum(), TestUpgradeSheet(), Coins, Sheet (+9 more)

### Community 29 - "Button"
Cohesion: 0.04
Nodes (111): Props, Size, ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Look, AspectRatio() (+103 more)

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
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 34 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), ImportReviewPage(), importRows() (+2 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "State"
Cohesion: 0.11
Nodes (11): event, fakeHandler, Handler, patchMessage, stateHandler, TopicKind, viewPatch, Topic (+3 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.24
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.10
Nodes (26): migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open() (+18 more)

### Community 41 - "ParseEffects"
Cohesion: 0.10
Nodes (26): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestCoreTraits() (+18 more)

### Community 42 - "Slugify"
Cohesion: 0.17
Nodes (20): Cantrips(), OptionLevel(), Options(), TestOptions(), NameDesc, TextData, Slugify(), Spell (+12 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - ".sendFeedback"
Cohesion: 0.33
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 46 - "CN"
Cohesion: 0.05
Nodes (77): Align, ContentProps, ctxKey, ctxState, Props, Side, ValueProps, ctxKey (+69 more)

### Community 47 - ".Plan"
Cohesion: 0.13
Nodes (11): ClassLevel, LevelUpChoice, TestSlotMaxima(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan, hpGain() (+3 more)

### Community 48 - "select_templ.go"
Cohesion: 0.11
Nodes (24): Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select(), Separator() (+16 more)

### Community 49 - "Session"
Cohesion: 0.09
Nodes (19): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity() (+11 more)

### Community 50 - "cell"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

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

### Community 56 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Parse"
Cohesion: 0.16
Nodes (11): findItem(), Parse(), tokenize(), TestUnknownCommand(), TestGearCommands(), ensureItem(), Item, coinCommand() (+3 more)

### Community 59 - "TestLiveUpdatesAndKick"
Cohesion: 0.33
Nodes (5): dial(), patchOf(), readUntil(), TestLiveUpdatesAndKick(), wsMessage

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "IfElse"
Cohesion: 0.13
Nodes (38): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+30 more)

### Community 63 - "App"
Cohesion: 0.08
Nodes (12): ListAttachmentFilesRow, Attachment, GetBlockRow, firstLine(), mdLinkText(), saveScreenshot(), saveSniffed(), joinPath() (+4 more)

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

### Community 68 - ".playerGroups"
Cohesion: 0.50
Nodes (3): nameKey(), TestPlayerGroups(), combatantGroup

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
Cohesion: 0.07
Nodes (57): Description(), Field(), Group(), Label(), Input(), Textarea(), armorForm(), builderFields() (+49 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "CustomMod"
Cohesion: 0.13
Nodes (15): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+7 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "nullID"
Cohesion: 0.13
Nodes (19): Player(), nullID(), TestCombatantNamesFollowSource(), TestMonsterPortrait(), testEnv, TestAttachmentAccess(), TestAttachmentToImage(), TestFilesCleanup() (+11 more)

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
Nodes (62): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+54 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "dialog_templ.go"
Cohesion: 0.16
Nodes (34): ctxKey, ctxState, Props, Dialog(), Close(), CloseFor(), Content(), Description() (+26 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.23
Nodes (15): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+7 more)

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
Cohesion: 0.17
Nodes (19): Content(), Description(), Header(), state(), Title(), Trigger(), triggerID(), combatCondition() (+11 more)

### Community 105 - "Encounter"
Cohesion: 0.08
Nodes (12): Combatant, Encounter, Entity, Character, entityTopics(), formInt(), CreateEncounterParams, ListCharacterEncountersParams (+4 more)

### Community 106 - "deviceOS"
Cohesion: 1.00
Nodes (3): deviceOS(), uaPick(), userAgentSummary()

### Community 107 - "Entry"
Cohesion: 0.21
Nodes (12): Entry, fixtures, normalizer, actionRank(), byPK(), children(), formatCR(), intField() (+4 more)

### Community 108 - "config.go"
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 109 - "message"
Cohesion: 0.33
Nodes (4): message, topicIDs(), Diff(), TestDiff()

### Community 110 - "open5e.go"
Cohesion: 0.08
Nodes (28): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+20 more)

### Community 111 - "wikidotClass"
Cohesion: 0.15
Nodes (20): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+12 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 116 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 117 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "buildReview"
Cohesion: 0.27
Nodes (5): buildReview(), csvSource(), formValues(), csvImport, csvImports

### Community 123 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 124 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 125 - "App"
Cohesion: 0.07
Nodes (15): Campaign, Note, viewerCampaign(), formIntQuery(), isHTMX(), isPhone(), redirect(), setAudience() (+7 more)

### Community 126 - "Handoff: open items from the 2026-11-03 review"
Cohesion: 0.25
Nodes (7): Bugs fixed since, Character page, Handoff: open items from the 2026-11-03 review, Improvements (bigger, deferred), Notes, Other pages, Test data in data-test

### Community 131 - "IfElse"
Cohesion: 0.10
Nodes (30): Props, Variant, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), alignmentColor(), classBadges() (+22 more)

### Community 133 - "rowFigure"
Cohesion: 0.50
Nodes (4): crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 135 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 136 - "spellCounts"
Cohesion: 0.67
Nodes (3): sortedColumns(), spellCounts(), TestSpellCounts()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **21 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `database/sql.NullInt64`, `State`, `Store`, `newTestEnv`, `Encounter`, `.loadLive`, `Slugify`, `perm.go`, `Hub`, `.fail`, `Character`, `cert.go`, `buildReview`, `time.Time`, `App`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09100564592766505 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Weapon`, `levelUpForm`, `SheetView`, `Encounter`, `CustomMod`, `Slugify`, `ParseEffects`, `.loadLive`, `.Plan`, `testing.T`, `Character`, `Parse`?**
  _High betweenness centrality (0.016) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011932905549926827 - nodes in this community are weakly interconnected._
- **Why does `dial()` connect `TestLiveUpdatesAndKick` to `compendium_templ.go`, `testing.T`, `Client`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._