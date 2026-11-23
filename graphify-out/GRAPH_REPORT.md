# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 372 files · ~641,994 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4685 nodes · 24192 edges · 151 communities (114 shown, 37 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2419 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f8d77d8b`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Sheet
- context.Context
- database/sql.NullInt64
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- spellCounts
- Path
- encoding/json.RawMessage
- select.js
- Command
- select_templ.go
- rosterCard
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Sum
- IfElse
- wikilink.js
- App
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- .compendiumPageSave
- redirect.go
- app.js
- tabbable.js
- useDismiss
- Client
- popover_templ.go
- command.js
- dialog.js
- Session
- ParseWikidot
- tooltip.js
- CompendiumKind
- parseCommand
- config.go
- Session
- composite.js
- buildReview
- Feature
- parseFeature
- nullID
- Character
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- ftItemData
- shadcn-templ Config
- Item
- golangci-lint Config
- Store
- Weapon
- radiogroup.js
- vendorjs/main.go
- ImageResults
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- Queries
- use_transition_status.js
- playwright
- .q
- ftRec
- checkbox.js
- ParseAidedd
- switch.js
- deathFix
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- Entity
- useTypeahead
- DownloadFiveTools
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- Amount
- mounted
- .sendFeedback
- Player
- profFocus
- CompendiumPath
- toast_templ.go
- ftBook
- NewClient
- BlockEditor
- TestSourceSwitchKeepsEdits
- Input
- csvSource
- .features
- export/markdown.go
- build.sh
- fakeHandler
- looseString
- errBodyWriter
- Client
- formErrors
- Before/after pairs
- Hub
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- condMods
- ASIError
- Book
- .Entry
- WeaponKind
- sourceUsedError
- Pending
- ListSharedWithRow
- SearchReadableBlocksRow
- Layout
- .ListBlockStepsSince
- .ListItemsByCategory
- .ListSpellsForClass
- .ListSubracesOf

## God Nodes (most connected - your core abstractions)
1. `App` - 374 edges
2. `Path()` - 267 edges
3. `newTestEnv()` - 170 edges
4. `Sheet` - 131 edges
5. `Button()` - 124 edges
6. `CN()` - 114 edges
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

## Communities (151 total, 37 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "Sheet"
Cohesion: 0.03
Nodes (121): featureSwitch, LevelUpChoice, internal/compendium, Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options() (+113 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (68): keepAwake(), FiveToolsBooks(), ftsQuery(), Search(), CompendiumSource, CreateCombatantParams, Completions(), EntityKind (+60 more)

### Community 3 - "database/sql.NullInt64"
Cohesion: 0.02
Nodes (56): ListHomebrewByOwnerRow, ListHomebrewRow, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Note, EncounterStatus (+48 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.10
Nodes (37): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+29 more)

### Community 6 - "SheetView"
Cohesion: 0.04
Nodes (124): ArmorEffect(), Coins, Slot, Signed(), Button(), Variant, Action(), Card() (+116 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

### Community 9 - "Path"
Cohesion: 0.04
Nodes (116): Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight() (+108 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (43): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+35 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.09
Nodes (43): CondOpKind, Effects, userText, parseProf(), parseProfSet(), parseSet(), parseSpell(), setIn() (+35 more)

### Community 13 - "select_templ.go"
Cohesion: 0.05
Nodes (55): ImageCategory, Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Separator() (+47 more)

### Community 14 - "rosterCard"
Cohesion: 0.17
Nodes (19): CampaignRoster(), deviceRow(), rosterCard(), devicesWhere(), Roster, agoTime(), localTime(), deviceLabel() (+11 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (59): ListCharacterNamesRow, Encounter, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+51 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Sum"
Cohesion: 0.19
Nodes (4): TestCustomMods(), TestExplainSums(), Sum(), TestWithEffects()

### Community 21 - "IfElse"
Cohesion: 0.08
Nodes (48): Props, Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup(), state() (+40 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "App"
Cohesion: 0.03
Nodes (48): ClearCookie(), SetCookie(), writeCookie(), FromContext(), Campaign, GetBlockRow, compendiumBack(), compendiumPath() (+40 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (177): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+169 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.10
Nodes (19): main(), run(), NewManager(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate() (+11 more)

### Community 28 - ".compendiumPageSave"
Cohesion: 0.16
Nodes (10): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+2 more)

### Community 29 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Client"
Cohesion: 0.11
Nodes (12): event, Handler, patchMessage, stateHandler, viewPatch, Topic, TopicKind, Diff() (+4 more)

### Community 34 - "popover_templ.go"
Cohesion: 0.13
Nodes (27): UI conventions, Content(), Description(), Header(), Popover(), state(), Title(), Trigger() (+19 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Session"
Cohesion: 0.10
Nodes (20): fixture, tokens, constantTimeEqual(), credentialsMatch(), newFixture(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity() (+12 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.04
Nodes (47): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, dataFor(), DecodeData(), ValidKind(), CountCompendiumRow (+39 more)

### Community 41 - "parseCommand"
Cohesion: 0.14
Nodes (14): findMod(), newMod(), parseMod(), IsModTarget(), EffectMods(), TestValidateEffectMods(), ValidateEffectMods(), featureCommand() (+6 more)

### Community 42 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 43 - "Session"
Cohesion: 0.24
Nodes (4): IsLoopback(), RequireAdmin(), sameOriginOnly(), useValidateTag()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 46 - "Feature"
Cohesion: 0.10
Nodes (16): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+8 more)

### Community 48 - "nullID"
Cohesion: 0.11
Nodes (37): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), acBonusForm(), effectFix, newEffectFix(), notes() (+29 more)

### Community 49 - "Character"
Cohesion: 0.08
Nodes (13): Character, Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), formErrors() (+5 more)

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (140): Defense, TestIsLoopback(), TestDecode(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice(), RemoveSourceDir(), TestRemoveSourceDir() (+132 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Store"
Cohesion: 0.06
Nodes (53): migration, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+45 more)

### Community 59 - "Weapon"
Cohesion: 0.10
Nodes (13): Part, Skill, ItemData, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon (+5 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "ImageResults"
Cohesion: 0.07
Nodes (44): ListImageCharactersRow, ListEntityFieldBacklinksRow, ListImageEntitiesRow, ListBacklinksRow, Entity, Look, attrString(), generateSVG() (+36 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.08
Nodes (46): Defenses, damageChoice(), capitalized(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName() (+38 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Queries"
Cohesion: 0.04
Nodes (55): HomebrewOwner, Source, sourceFormat, bindFiveToolsImages(), SaveHomebrew(), SeedHabitats(), storeSource(), upsert() (+47 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - ".q"
Cohesion: 0.10
Nodes (28): isCleared(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+20 more)

### Community 72 - "ftRec"
Cohesion: 0.24
Nodes (9): ftCopyRef, ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve() (+1 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

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

### Community 86 - "Entity"
Cohesion: 0.17
Nodes (6): Entity, entityTopics(), SetEntityCurrentImageParams, ToggleEntityImageRevealedParams, UnlinkEntityImageParams, lookOps

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (24): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook(), DownloadFiveTools() (+16 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "Amount"
Cohesion: 0.32
Nodes (5): Amount, CritDamage(), Damage(), Set(), TestApplyHPDeath()

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.24
Nodes (14): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+6 more)

### Community 103 - "Player"
Cohesion: 0.07
Nodes (39): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas (+31 more)

### Community 105 - "CompendiumPath"
Cohesion: 0.04
Nodes (102): Props, Variant, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListHomebrewLinksRow, ListMonstersByNameRow, SearchCompendiumRow (+94 more)

### Community 106 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (18): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftRange(), ftRecords() (+10 more)

### Community 108 - "NewClient"
Cohesion: 0.23
Nodes (18): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+10 more)

### Community 111 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 112 - "Input"
Cohesion: 0.06
Nodes (102): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+94 more)

### Community 116 - ".features"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 117 - "export/markdown.go"
Cohesion: 0.13
Nodes (19): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+11 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Client"
Cohesion: 0.19
Nodes (3): message, Client, isDevice()

### Community 127 - "formErrors"
Cohesion: 0.29
Nodes (5): Handlers, keepsBody(), fieldError, pickError, userError

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "Hub"
Cohesion: 0.22
Nodes (3): ClientInfo, goneDevice, Hub

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "condMods"
Cohesion: 0.40
Nodes (4): isToggle(), condMods(), effectChip(), hasEffect()

### Community 139 - "Book"
Cohesion: 0.11
Nodes (24): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+16 more)

### Community 140 - ".Entry"
Cohesion: 0.08
Nodes (31): MonsterAction, addEntry(), setStat(), MonsterData, habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels() (+23 more)

### Community 141 - "WeaponKind"
Cohesion: 0.67
Nodes (3): parseMastery(), TestWeaponKind(), WeaponKind()

### Community 143 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 146 - "Layout"
Cohesion: 0.11
Nodes (37): Scripts(), Separator(), Spinner(), campaignLink(), PendingRequests(), PendingGroup, CharacterPage(), HomebrewPage() (+29 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 414 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **37 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Sheet`, `context.Context`, `database/sql.NullInt64`, `Hub`, `Book`, `Command`, `Pending`, `Viewer`, `cert.go`, `CompendiumKind`, `Session`, `Character`, `testing.T`, `Store`, `FiveToolsJob`, `Queries`, `.q`, `Entity`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1398056429275621 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Sheet`, `database/sql.NullInt64`, `tabs_templ.go`, `SheetView`, `select_templ.go`, `rosterCard`, `Viewer`, `Layout`, `IfElse`, `App`, `github.com/a-h/templ.Attributes`, `popover_templ.go`, `nullID`, `Character`, `testing.T`, `ImageResults`, `CombatantView`, `Queries`, `.q`, `deathFix`, `ParseWiki`, `Entity`, `.sendFeedback`, `Player`, `CompendiumPath`, `TestSourceSwitchKeepsEdits`, `Input`, `export/markdown.go`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **Should `Sheet` be split into smaller, more focused modules?**
  _Cohesion score 0.02885113436237439 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `wikilink.js`, `Player`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._