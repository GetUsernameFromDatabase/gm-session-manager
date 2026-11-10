# Graph Report - gm-session-manager  (2026-11-10)

## Corpus Check
- 296 files · ~577,326 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4148 nodes · 18019 edges · 250 communities (128 shown, 122 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1839 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e74e9d79`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- .do
- createProvider
- Weapon
- SheetView
- combobox.js
- redirect.go
- Layout
- Import
- select.js
- Command
- .Entry
- ImportView
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Button
- levelUpForm
- wikilink.js
- tabs_templ.go
- popover.js
- github.com/a-h/templ.Component
- Project Rules (CLAUDE.md, README)
- ParseSheet
- github.com/gin-gonic/gin.Context
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- Slugify
- testing.T
- command.js
- dialog.js
- Client
- Queries
- tooltip.js
- .CreateEncounter
- ParseAidedd
- Book
- cell
- composite.js
- Campaign
- Feature
- ClassInfo
- github.com/a-h/templ.Component
- State
- App
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- .Router
- Character
- radiogroup.js
- vendorjs/main.go
- Input
- FiveToolsJob
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- TestDownloadFiveTools
- checkbox.js
- planSwitch
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- fromLookup
- useFocus
- mark_others.js
- .ListEntrySourcesBySlugs
- Lint and Release Rules
- Run Script
- .sendFeedback
- useTypeahead
- toast_templ.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- joinNonEmpty
- mounted
- NewClient
- .levelUp
- Session
- popover_templ.go
- CampaignSources
- ftBook
- TextData
- .encounterView
- formErrors
- export/markdown.go
- Open
- TestRunCommandUsesCurrentSheet
- Pending
- DecodeData
- Topic
- shadcn-templ.go
- build.sh
- combobox_templ.go
- inputgroup_templ.go
- errBodyWriter
- field_templ.go
- buildReview
- .ListItemsByCategory
- Before/after pairs
- select_templ.go
- open5e.go
- ftItemData
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- .ListCompendiumBySlugs
- Site
- .ListSpellsForClass
- command_templ.go
- dialog_templ.go
- Campaign
- ParseWiki
- DBTX
- Character
- CompendiumSource
- .gearFor
- tooltip_templ.go
- .sourceIssues
- alert_templ.go
- radiogroup_templ.go
- Entity
- Props
- Props
- NoteBlock
- keepawake_windows.go
- Client
- subclassPanels
- ListSharedWithRow
- SearchReadableBlocksRow
- lanHostname
- .ListBlockStepsSince
- .ListCompendiumByKind
- .ListCompendiumNames
- .ListSharedEntries
- keepawake_other.go
- ListImageCharactersRow
- CountCompendiumRow
- ListHomebrewByOwnerRow
- ListHomebrewLinksRow
- ListHomebrewRow
- UpsertCompendiumEntryParams
- ListImageEntitiesRow
- AddBlockAudienceParams
- AddBlockLinkParams
- AddEntityFieldLinkParams
- AddHabitatByNameParams
- AddMonsterHabitatParams
- AllowCampaignSourceParams
- BindMonsterImageParams
- BindMonsterImagesParams
- .CreateCampaign
- CreateLevelUpParams
- spellCounts
- DisallowCampaignSourceParams
- EnsureImageCategoryParams
- ExpireTurnConditionsParams
- .GetCompendiumEntryByName
- .GetCompendiumEntryBySlug
- .GetCompendiumEntryBySlugSource
- ImageBindingRow
- ImageLookAccessRow
- InsertBlockStepParams
- LevelUp
- LinkCharacterImageParams
- LinkEntityImageParams
- LinkEntryCampaignParams
- .ListBacklinks
- ListBlockDocsRow
- ListCharacterSheetsNamedRow
- ListCharacterSheetsRow
- ListCompendiumKeysBySourceRow
- .ListEncountersByMonster
- .ListEntityFieldBacklinks
- ListExportBlocksRow
- ListImageAttachmentsRow
- .ListImages
- ListMasteryWeaponsRow
- .ListReadableNotes
- ListSourceRefUsersRow
- ListStepJSONRow
- .ListSubclassesOf
- MonsterHabitat
- .RandomCompendiumEntry
- .RemoveHabitatByName
- .RemoveMonsterHabitat
- .RenameCampaign
- .RenameNote
- .SaveBlockSnapshot
- .SearchCompendium
- .SetBlockPosition
- .SetBlockVersion
- .SetCampaignOpenJoin
- SetCharacterCurrentImageParams
- .SetCombatantHP
- SetCombatantInitiativeParams
- .SetCombatantMarker
- SetCompendiumImageParams
- .SetCompendiumSourcePriority
- SetEntityCurrentImageParams
- ToggleCharacterImageRevealedParams
- ToggleEntityImageRevealedParams
- .TrimBlockSteps
- UnlinkCharacterImageParams
- UnlinkEntityImageParams
- .UnlinkEntryCampaign
- .UpdateBlockAccess
- .UpdateCharacter
- .UpdateCharacterSheet
- UpdateEncounterStateParams
- UpdateEntityParams
- .UpdateHomebrew
- .UpsertCompendiumSource
- .CountSharedByKind
- .ListSharedOrder
- CreateHomebrewParams
- .ListEntryCopies

## God Nodes (most connected - your core abstractions)
1. `App` - 310 edges
2. `newTestEnv()` - 122 edges
3. `Button()` - 120 edges
4. `Sheet` - 118 edges
5. `CN()` - 113 edges
6. `IfElse()` - 61 edges
7. `SheetView` - 57 edges
8. `Command` - 49 edges
9. `Input()` - 47 edges
10. `Slugify()` - 43 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `UI conventions` --references--> `removeButton()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go
- `UI conventions` --references--> `formSelect()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (250 total, 122 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (60): contextKey, classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild, o5eClass (+52 more)

### Community 1 - "CombatantView"
Cohesion: 0.13
Nodes (28): beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), downBadge(), Encounter(), EncounterView (+20 more)

### Community 3 - ".do"
Cohesion: 0.15
Nodes (17): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+9 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.05
Nodes (26): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+18 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (76): SpellData, ArmorEffect(), Signed(), Variant, Action(), Card(), Content(), Description() (+68 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "Layout"
Cohesion: 0.05
Nodes (100): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Content(), Description(), Empty(), Header(), Media() (+92 more)

### Community 10 - "Import"
Cohesion: 0.09
Nodes (39): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+31 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (67): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+59 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (29): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), TestTokensReadAsLabels(), HomebrewInput (+21 more)

### Community 14 - "ImportView"
Cohesion: 0.18
Nodes (13): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows(), importStatusItems() (+5 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (45): Note, GetBlockRow, Entity, Fields, Look, ParseFields(), VisibleEntity(), VisibleLooks() (+37 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Button"
Cohesion: 0.08
Nodes (46): Field, Button(), CloseFor(), Content(), Tooltip(), Trigger(), armorForm(), builderForm() (+38 more)

### Community 21 - "levelUpForm"
Cohesion: 0.07
Nodes (59): EquipmentItem, EquipmentOption, Checkbox(), Label(), Item(), RadioGroup(), FiveToolsPicker(), feedbackKind() (+51 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.13
Nodes (31): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+23 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (78): ctxState, ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer() (+70 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.14
Nodes (28): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+20 more)

### Community 28 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (24): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), TestCampaignExcludesOwnerBlocks(), coinCommand(), commandText() (+16 more)

### Community 29 - "newTestEnv"
Cohesion: 0.05
Nodes (53): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait() (+45 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Slugify"
Cohesion: 0.12
Nodes (27): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, SizeNamed(), catchUpResistances() (+19 more)

### Community 34 - "testing.T"
Cohesion: 0.03
Nodes (82): Defense, serveMirror(), TestDownloadFiveToolsFailureLeavesNothing(), TestFiveToolsBooks(), RemoveSourceDir(), TestRemoveSourceDir(), Parse(), tokenize() (+74 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 38 - "Queries"
Cohesion: 0.11
Nodes (31): dataFor(), DecodeData(), Kind, Queries, CompendiumEntry, ClassFeature, ClassInfo, LevelUpPlan (+23 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - "Book"
Cohesion: 0.16
Nodes (18): bookWriter, mirror, rawBook, internal/compendium, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools() (+10 more)

### Community 43 - "cell"
Cohesion: 0.13
Nodes (18): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+10 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Campaign"
Cohesion: 0.16
Nodes (7): Campaign, TestUserAgent(), userAgentSummary(), joinPath(), visitorOnly(), JoinQRSrc(), CompendiumSwitch

### Community 46 - "Feature"
Cohesion: 0.20
Nodes (7): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, BuilderExtra, usesValue()

### Community 47 - "ClassInfo"
Cohesion: 0.10
Nodes (14): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs() (+6 more)

### Community 48 - "github.com/a-h/templ.Component"
Cohesion: 0.06
Nodes (94): CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, SearchCompendiumRow, UI conventions, Badge(), IfElse(), inFightSuffix() (+86 more)

### Community 49 - "State"
Cohesion: 0.19
Nodes (8): patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff(), liveView

### Community 50 - "App"
Cohesion: 0.06
Nodes (15): compendiumBack(), compendiumPath(), homebrewTabURL(), ImportSummary(), pathCampaign(), viewerCampaign(), firstLine(), mdLinkText() (+7 more)

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
Nodes (14): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+6 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ClassData"
Cohesion: 0.14
Nodes (22): block, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster(), sourceLine() (+14 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - ".Router"
Cohesion: 0.24
Nodes (4): RequireAdmin(), sameOriginOnly(), useValidateTag(), CompendiumRest()

### Community 59 - "Character"
Cohesion: 0.10
Nodes (11): Character, CanChangeSubclass(), CanCharacter(), TestCanChangeSubclass(), linkCharacterLook(), BuilderPickItems(), BuilderPick, pickURL() (+3 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Input"
Cohesion: 0.07
Nodes (76): Defenses, Combobox(), Content(), Item(), List(), Dialog(), Close(), Content() (+68 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.05
Nodes (37): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Condition, CreateCombatantParams, applyCondOps(), saveScreenshot() (+29 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.09
Nodes (32): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps (+24 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "TestDownloadFiveTools"
Cohesion: 0.12
Nodes (20): ClassFeature, ClassData, FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition() (+12 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "planSwitch"
Cohesion: 0.14
Nodes (12): featureSwitch, featureNames(), SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch(), sameSlug() (+4 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.12
Nodes (19): Player(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips() (+11 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - "fromLookup"
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 86 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 100 - "joinNonEmpty"
Cohesion: 0.08
Nodes (38): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+30 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - ".levelUp"
Cohesion: 0.13
Nodes (20): Cantrips(), Slugify(), Spell, spellFromRow(), spellNeed(), classPin(), TestSpellNeed(), featKey() (+12 more)

### Community 104 - "Session"
Cohesion: 0.10
Nodes (19): fixture, tokens, constantTimeEqual(), credentialsMatch(), isCleared(), newFixture(), serve(), TestCredentialsMatch() (+11 more)

### Community 105 - "popover_templ.go"
Cohesion: 0.21
Nodes (11): state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps, HeaderProps (+3 more)

### Community 106 - "CampaignSources"
Cohesion: 0.18
Nodes (16): campaignSourceRow(), CampaignSources(), CampaignSourcesView, sourceGroupButton(), sourceRowTitle(), CampaignOverview(), CampaignRoster(), Roster (+8 more)

### Community 107 - "ftBook"
Cohesion: 0.15
Nodes (22): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftDecode(), ftDuration() (+14 more)

### Community 108 - "TextData"
Cohesion: 0.38
Nodes (6): TextData, textMarkdown(), countWord(), featMasteries(), featText(), TestTextMarkdown()

### Community 109 - ".encounterView"
Cohesion: 0.08
Nodes (21): Combatant, Encounter, Entity, ListCombatantsRow, Character, combatantDefenses(), damageChoice(), initiativeMod() (+13 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 111 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 112 - "Open"
Cohesion: 0.14
Nodes (16): migration, apply(), loadMigrations(), Migrate(), Open(), openPool(), openTemp(), TestAttachmentOutlivesBlock() (+8 more)

### Community 113 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 114 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 115 - "DecodeData"
Cohesion: 0.12
Nodes (22): CoreTraits(), EquipmentOptions(), SkillChoice(), SplitList(), Squash(), TestCoreTraits(), TestEquipmentOptions(), TestSkillChoice() (+14 more)

### Community 116 - "Topic"
Cohesion: 0.13
Nodes (7): event, fakeHandler, Handler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 117 - "shadcn-templ.go"
Cohesion: 0.11
Nodes (17): Props, initialChecked(), Direction(), flattenClasses(), lastTopLevelColon(), MergeAttributes(), Ptr(), twImportantToV3() (+9 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "combobox_templ.go"
Cohesion: 0.11
Nodes (22): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+14 more)

### Community 123 - "inputgroup_templ.go"
Cohesion: 0.17
Nodes (13): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+5 more)

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 126 - "buildReview"
Cohesion: 0.15
Nodes (8): buildReview(), csvSource(), formValues(), parseCSV(), TestGuessMapping(), capReader, csvImport, csvImports

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "select_templ.go"
Cohesion: 0.15
Nodes (13): initialValue(), Align, ContentProps, ctxKey, ctxState, GroupProps, ItemProps, LabelProps (+5 more)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 134 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 137 - "ftPrerequisite"
Cohesion: 0.25
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Site"
Cohesion: 0.14
Nodes (13): Site, Source, fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page, pageSource() (+5 more)

### Community 141 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 142 - "dialog_templ.go"
Cohesion: 0.18
Nodes (11): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+3 more)

### Community 143 - "Campaign"
Cohesion: 0.22
Nodes (7): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), parseCharacterFile()

### Community 144 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 148 - ".gearFor"
Cohesion: 0.40
Nodes (3): ItemData, compendiumWeapon(), TestCompendiumWeapon()

### Community 149 - "tooltip_templ.go"
Cohesion: 0.32
Nodes (7): DisableTrigger(), state(), ContentProps, ctxKey, ctxState, Props, Side

### Community 150 - ".sourceIssues"
Cohesion: 0.53
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 151 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 152 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 154 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 155 - "Props"
Cohesion: 0.60
Nodes (5): attrString(), generateSVG(), getIconContent(), Props, Icon()

### Community 156 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 159 - "Client"
Cohesion: 0.20
Nodes (4): message, IsLoopback(), TestIsLoopback(), Client

### Community 160 - "subclassPanels"
Cohesion: 0.40
Nodes (5): ListSubclassesOfRow, subclassIndex(), subclassOptions(), subclassPanels(), TestSubclassPanels()

### Community 163 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 186 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 384 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **122 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `Weapon`, `Layout`, `perm.go`, `.gearFor`, `.sourceIssues`, `github.com/gin-gonic/gin.Context`, `Slugify`, `Client`, `Queries`, `Campaign`, `State`, `cert.go`, `.Router`, `Character`, `FiveToolsJob`, `database/sql.NullInt64`, `.levelUp`, `TextData`, `.encounterView`, `Open`, `DecodeData`, `Topic`, `buildReview`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1101033241385759 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Weapon` to `compendium_templ.go`, `Slugify`, `testing.T`, `joinNonEmpty`, `Queries`, `.levelUp`, `SheetView`, `planSwitch`, `Command`, `.encounterView`, `Feature`, `ClassInfo`, `TestRunCommandUsesCurrentSheet`, `App`, `DecodeData`, `levelUpForm`, `.sourceIssues`, `ParseSheet`?**
  _High betweenness centrality (0.018) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.13048780487804879 - nodes in this community are weakly interconnected._
- **Why does `Queries` connect `Queries` to `compendium_templ.go`, `database/sql.NullInt64`, `.do`, `Import`, `Site`, `.Entry`, `.encounterView`, `Campaign`, `Open`, `DBTX`, `perm.go`, `App`, `Character`, `github.com/gin-gonic/gin.Context`?**
  _High betweenness centrality (0.014) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.03076923076923077 - nodes in this community are weakly interconnected._