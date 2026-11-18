# Graph Report - gm-session-manager  (2026-11-18)

## Corpus Check
- 360 files · ~627,742 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4587 nodes · 23456 edges · 154 communities (120 shown, 34 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2336 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `df7787a8`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- Condition
- createProvider
- Layout
- SheetView
- combobox.js
- redirect.go
- viewerCtx
- encoding/json.RawMessage
- select.js
- Command
- tooltip_templ.go
- .do
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- newTestEnv
- WizardView
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Session
- Button
- NewClient
- app.js
- tabbable.js
- useDismiss
- Character
- ParseWikidot
- command.js
- dialog.js
- CompendiumPath
- Habitat
- tooltip.js
- Sheet
- Client
- Book
- Import
- composite.js
- buildReview
- Feature
- db/db.go
- nullID
- CompendiumKind
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- parseHP
- golangci-lint Config
- sameSlug
- Item
- radiogroup.js
- Look
- capReader
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- github.com/a-h/templ.Component
- use_transition_status.js
- playwright
- release.sh
- .Router
- checkbox.js
- ParseAidedd
- switch.js
- storeSource
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- sourcesEnv
- useTypeahead
- vendorjs/main.go
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- database/sql.NullInt64
- mounted
- Handlers
- newLiveFixture
- profFocus
- InputForEntry
- Campaign
- ftBook
- TestSourceSwitchKeepsEdits
- BlockEditor
- WeaponKind
- config.go
- ItemData
- State
- export/markdown.go
- build.sh
- Pending
- looseString
- errBodyWriter
- ParseEffects
- DownloadFiveTools
- tabs_templ.go
- Before/after pairs
- parseMod
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- ASIError
- appState
- .Entry
- EffectMods
- sourceUsedError
- ListSharedWithRow
- .sendFeedback
- EffectFieldsURL
- TestTaughtPicked
- .ListSpellsForClass
- RosterCharacter
- rowFigure
- keepawake_windows.go
- SearchReadableBlocksRow
- .ListSubracesOf
- Weapon

## God Nodes (most connected - your core abstractions)
1. `App` - 369 edges
2. `Path()` - 248 edges
3. `newTestEnv()` - 156 edges
4. `Sheet` - 130 edges
5. `Button()` - 122 edges
6. `CN()` - 113 edges
7. `SheetView` - 78 edges
8. `IfElse()` - 65 edges
9. `Command` - 60 edges
10. `CompendiumKind` - 56 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (154 total, 34 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (53): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+45 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.14
Nodes (25): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+17 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (113): keepAwake(), HomebrewOwner, sourceFormat, featureSwitch, Decode(), bindFiveToolsImages(), HomebrewInput, SaveHomebrew() (+105 more)

### Community 3 - "Condition"
Cohesion: 0.05
Nodes (46): Combatant, Condition, CreateConditionParams, ListCombatantsRow, Character, EffectEnd, passTurn(), rollFor() (+38 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Layout"
Cohesion: 0.05
Nodes (52): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, Scripts(), Separator(), Toast(), Toaster(), toastIcon() (+44 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (84): ArmorEffect(), Coins, Slot, Signed(), Variant, Action(), Card(), Content() (+76 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "viewerCtx"
Cohesion: 0.17
Nodes (10): Topic, TopicKind, heldCollected(), notFoundIsNil(), parseLiveKey(), heldTopics, liveKey, liveKind (+2 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (45): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+37 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.10
Nodes (37): CondOpKind, Effects, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn(), parseCoin() (+29 more)

### Community 13 - "tooltip_templ.go"
Cohesion: 0.09
Nodes (34): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks(), closeAfter() (+26 more)

### Community 14 - ".do"
Cohesion: 0.12
Nodes (22): testEnv, mustJSON(), pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm() (+14 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.04
Nodes (56): ListCharacterNamesRow, GetBlockRow, ListBacklinksRow, LinkTarget, Visibility, Entity, Fields, ParseFields() (+48 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "newTestEnv"
Cohesion: 0.04
Nodes (106): TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestConditionExpiresOnTurn() (+98 more)

### Community 21 - "WizardView"
Cohesion: 0.06
Nodes (65): LevelUpPlan, UI conventions, Alert(), Description(), Checkbox(), Empty(), Label(), initialValue() (+57 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "App"
Cohesion: 0.03
Nodes (44): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), CanChangeSubclass(), CanCharacter(), compendiumBack() (+36 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (153): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+145 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Session"
Cohesion: 0.08
Nodes (23): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+15 more)

### Community 28 - "Button"
Cohesion: 0.09
Nodes (56): ctxKey, ctxState, Props, SwitchFeature, SwitchPlan, Button(), Dialog(), Close() (+48 more)

### Community 29 - "NewClient"
Cohesion: 0.28
Nodes (15): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+7 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Character"
Cohesion: 0.07
Nodes (30): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, DecodeData(), Character, WornArmor, Campaign() (+22 more)

### Community 34 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "CompendiumPath"
Cohesion: 0.06
Nodes (70): Props, Variant, ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props (+62 more)

### Community 38 - "Habitat"
Cohesion: 0.06
Nodes (50): CountHabitatsRow, ListEntryCopiesRow, ListMonstersByNameRow, Habitat, builderForm(), builderRemove(), chooseFirst(), opGroup() (+42 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Sheet"
Cohesion: 0.05
Nodes (71): internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), SkillChoice(), TestOptions() (+63 more)

### Community 41 - "Client"
Cohesion: 0.20
Nodes (3): ClientInfo, event, Hub

### Community 42 - "Book"
Cohesion: 0.12
Nodes (23): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+15 more)

### Community 43 - "Import"
Cohesion: 0.10
Nodes (38): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+30 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.11
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

### Community 46 - "Feature"
Cohesion: 0.12
Nodes (10): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+2 more)

### Community 47 - "db/db.go"
Cohesion: 0.12
Nodes (24): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+16 more)

### Community 48 - "nullID"
Cohesion: 0.09
Nodes (39): DeathSaves, nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), deathFix(), effectFix, TestDeathClearsConditions() (+31 more)

### Community 49 - "CompendiumKind"
Cohesion: 0.04
Nodes (36): TestHomebrew(), dataFor(), ValidKind(), ftsQuery(), Search(), CountCompendiumRow, ListHomebrewByOwnerRow, ListHomebrewRow (+28 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (111): Defense, TestDecode(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), ftRaceAbility(), RemoveSourceDir(), TestRemoveSourceDir() (+103 more)

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

### Community 56 - "parseHP"
Cohesion: 0.12
Nodes (20): Amount, userText, CritDamage(), Damage(), Heal(), parseDamage(), parseHP(), parseMaxHP() (+12 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (31): deny, files, rules, generated, paths, rules, formatters, enable (+23 more)

### Community 59 - "Item"
Cohesion: 0.12
Nodes (9): findItem(), parseArmor(), parseItem(), TestCustomMods(), TestExplainSums(), Sum(), TestGearCommands(), ensureItem() (+1 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Look"
Cohesion: 0.43
Nodes (7): Look, initials(), lookControls(), looksGrid(), npcPortrait(), currentLook(), LookItem()

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.11
Nodes (38): Defenses, capitalized(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow() (+30 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Component"
Cohesion: 0.08
Nodes (70): ctxState, Field, Chip(), Chips(), ChipsInput(), Combobox(), Content(), Input() (+62 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.60
Nodes (4): die(), GOFLAGS, newer(), release.sh script

### Community 72 - ".Router"
Cohesion: 0.15
Nodes (8): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), TestHandlerErrors(), sameOriginOnly(), useValidateTag()

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "storeSource"
Cohesion: 0.16
Nodes (10): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+2 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "ParseWiki"
Cohesion: 0.22
Nodes (10): wikiLink(), wikiText(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks(), WikiLinks() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "notes.js"
Cohesion: 0.15
Nodes (11): attachedPicture(), ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, json(), links(), nodeActions() (+3 more)

### Community 85 - "Run Script"
Cohesion: 0.60
Nodes (5): changes(), GOFLAGS, restart_server(), run.sh script, stop_server()

### Community 86 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "database/sql.NullInt64"
Cohesion: 0.02
Nodes (60): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Encounter, Entity, Note, EncounterStatus (+52 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "Handlers"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 103 - "newLiveFixture"
Cohesion: 0.11
Nodes (26): Player(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Websocket and live views, TestNoteViewAudiences() (+18 more)

### Community 105 - "InputForEntry"
Cohesion: 0.22
Nodes (8): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 106 - "Campaign"
Cohesion: 0.25
Nodes (4): Campaign, Roster, CreateCampaignParams, CampaignSwitch

### Community 107 - "ftBook"
Cohesion: 0.13
Nodes (23): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftCopyTargets(), ftDecode() (+15 more)

### Community 108 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 111 - "WeaponKind"
Cohesion: 0.40
Nodes (5): Completions(), parseMastery(), TestCompletions(), TestWeaponKind(), WeaponKind()

### Community 112 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 113 - "ItemData"
Cohesion: 0.05
Nodes (56): ImageCategory, Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select() (+48 more)

### Community 116 - "State"
Cohesion: 0.09
Nodes (11): fakeHandler, Handler, message, patchMessage, stateHandler, viewPatch, Client, Diff() (+3 more)

### Community 117 - "export/markdown.go"
Cohesion: 0.13
Nodes (19): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+11 more)

### Community 118 - "build.sh"
Cohesion: 0.53
Nodes (5): build(), CGO_ENABLED, GOFLAGS, build.sh script, step()

### Community 119 - "Pending"
Cohesion: 0.22
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "ParseEffects"
Cohesion: 0.12
Nodes (10): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+2 more)

### Community 126 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (25): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook(), DownloadFiveTools() (+17 more)

### Community 127 - "tabs_templ.go"
Cohesion: 0.12
Nodes (36): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+28 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "parseMod"
Cohesion: 0.50
Nodes (4): findMod(), newMod(), parseMod(), IsModTarget()

### Community 133 - "open5e.go"
Cohesion: 0.11
Nodes (24): Entry, fixture, fixtures, normalizer, o5eProperty, Source, TestOpen5eMastery(), TestStoreSourceRollsBackOnAfterSave() (+16 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.14
Nodes (14): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ftArmor(), ftCode(), ftItemData() (+6 more)

### Community 139 - "appState"
Cohesion: 0.47
Nodes (5): loadSettings(), saveSettings(), TestSettingsFile(), appState, settings

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 141 - "EffectMods"
Cohesion: 0.25
Nodes (7): EffectMods(), TestValidateEffectMods(), TestWithEffects(), ValidateEffectMods(), Gotchas, internal/web, TestPagesRender()

### Community 144 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 145 - "EffectFieldsURL"
Cohesion: 0.33
Nodes (4): EffectFieldsURL(), EffectID(), effectQuery(), EffectTarget

### Community 146 - "TestTaughtPicked"
Cohesion: 0.67
Nodes (3): TestTaughtPicked(), taughtPicked(), taughtRules()

### Community 149 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 150 - "rowFigure"
Cohesion: 0.50
Nodes (4): TestRowFigure(), crText(), rowFigure(), textOnly()

### Community 156 - "Weapon"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 410 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **34 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Condition`, `viewerCtx`, `appState`, `Command`, `.do`, `Viewer`, `Character`, `Sheet`, `Client`, `Book`, `CompendiumKind`, `testing.T`, `cert.go`, `FiveToolsJob`, `.Router`, `database/sql.NullInt64`, `Campaign`, `State`, `Pending`?**
  _High betweenness centrality (0.056) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14431486880466474 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `newTestEnv` to `compendium_templ.go`, `popover_templ.go`, `Condition`, `Layout`, `SheetView`, `EffectMods`, `.do`, `tooltip_templ.go`, `.sendFeedback`, `Viewer`, `EffectFieldsURL`, `WizardView`, `App`, `github.com/a-h/templ.Attributes`, `Session`, `Button`, `Character`, `CompendiumPath`, `Habitat`, `Sheet`, `Import`, `nullID`, `CompendiumKind`, `testing.T`, `Look`, `CombatantView`, `github.com/a-h/templ.Component`, `ParseWiki`, `sourcesEnv`, `database/sql.NullInt64`, `newLiveFixture`, `TestSourceSwitchKeepsEdits`, `ItemData`, `export/markdown.go`, `tabs_templ.go`?**
  _High betweenness centrality (0.038) - this node is a cross-community bridge._
- **Should `popover_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14153846153846153 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `newTestEnv`, `Command`, `EffectMods`, `wikilink.js`?**
  _High betweenness centrality (0.024) - this node is a cross-community bridge._