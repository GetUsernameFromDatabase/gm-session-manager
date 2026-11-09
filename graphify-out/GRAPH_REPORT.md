# Graph Report - gm-session-manager  (2026-11-09)

## Corpus Check
- 293 files · ~561,145 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 52 file(s) not represented in the graph (top: .templ 44, .css 4, (none) 3)

## Summary
- 4053 nodes · 17489 edges · 234 communities (117 shown, 117 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1777 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `25705215`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- context.Context
- .do
- createProvider
- .blockFor
- SheetView
- combobox.js
- redirect.go
- SetEntityCurrentImageParams
- spellCounts
- select.js
- Command
- .Entry
- Button
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Input
- levelUpForm
- wikilink.js
- tabs_templ.go
- popover.js
- shadcn-templ.go
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- testing.T
- app.js
- tabbable.js
- useDismiss
- App
- Parse
- command.js
- dialog.js
- CN
- select_templ.go
- tooltip.js
- .encounterView
- TestDownloadFiveTools
- buildReview
- cell
- composite.js
- .wizardCommand
- ClassInfo
- State
- Character
- ParseWiki
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- CommandBuilder
- Book
- radiogroup.js
- vendorjs/main.go
- ImportView
- FiveToolsJob
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- Slugify
- checkbox.js
- Session
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .Router
- useFocus
- mark_others.js
- ftDuration
- Lint and Release Rules
- Run Script
- Queries
- useTypeahead
- Feature
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- NewClient
- Site
- .classFile
- .fail
- .whoami
- ftBook
- Store
- SearchReadableBlocksRow
- formErrors
- export/markdown.go
- Topic
- TestSourceSwitchKeepsEdits
- Campaign
- Sheet
- Slugify
- Pending
- build.sh
- Hub
- .levelUp
- errBodyWriter
- PlanSourceSwitch
- .ListItemsByCategory
- RosterCharacter
- Before/after pairs
- alert_templ.go
- open5e.go
- Import
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- Look
- DBTX
- Character
- CompendiumSource
- toast_templ.go
- Entity
- .joinInfo
- NoteBlock
- keepawake_windows.go
- CountHabitatsRow
- lanHostname
- .ListBlockStepsSince
- .ListCompendiumByKind
- .ListCompendiumBySlugs
- .ListCompendiumNames
- .ListEntrySourcesBySlugs
- .ListSpellsForClass
- keepawake_other.go
- ListImageCharactersRow
- CountCompendiumRow
- ListHomebrewByOwnerRow
- ListHomebrewLinksRow
- ListHomebrewRow
- ListSubclassesOfRow
- UpsertCompendiumEntryParams
- ListImageEntitiesRow
- AddBlockAudienceParams
- AddBlockLinkParams
- AddEntityFieldLinkParams
- AddHabitatByNameParams
- AddMonsterHabitatParams
- BindMonsterImageParams
- BindMonsterImagesParams
- .CreateCampaign
- CreateConditionParams
- .CreateEncounter
- CreateHomebrewParams
- CreateLevelUpParams
- EnsureImageCategoryParams
- ExpireTurnConditionsParams
- FindCharacterConditionParams
- FindCombatantConditionParams
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
- ListStepJSONRow
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

## God Nodes (most connected - your core abstractions)
1. `App` - 300 edges
2. `newTestEnv()` - 120 edges
3. `Sheet` - 118 edges
4. `Button()` - 115 edges
5. `CN()` - 113 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 56 edges
8. `Command` - 49 edges
9. `Input()` - 46 edges
10. `Slugify()` - 42 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go
- `internal/compendium` --references--> `Import()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/importer.go
- `Gotchas` --references--> `TestPagesRender()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/pages_test.go
- `UI conventions` --references--> `removeButton()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (234 total, 117 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.10
Nodes (59): contextKey, classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild, o5eClass (+51 more)

### Community 1 - "CombatantView"
Cohesion: 0.09
Nodes (46): Props, Variant, Condition, Defenses, Badge(), badgeClasses(), Header(), Title() (+38 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (21): isCleared(), serve(), TestMiddlewareIdentity(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict() (+13 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - ".blockFor"
Cohesion: 0.16
Nodes (5): GetBlockRow, imageItem(), imageVersion(), blockAccessLevel, stepsPayload

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (98): ArmorEffect(), Signed(), Field, Button(), Variant, Action(), Card(), Content() (+90 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 10 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (69): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+61 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), TestTokensReadAsLabels() (+22 more)

### Community 14 - "Button"
Cohesion: 0.04
Nodes (127): SearchCompendiumRow, ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Scripts(), Separator(), Spinner(), itemDetails() (+119 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.09
Nodes (40): Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew() (+32 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Input"
Cohesion: 0.06
Nodes (90): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+82 more)

### Community 21 - "levelUpForm"
Cohesion: 0.09
Nodes (51): Props, UI conventions, Footer(), Checkbox(), initialChecked(), Label(), initialValue(), Item() (+43 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.09
Nodes (45): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+37 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "shadcn-templ.go"
Cohesion: 0.05
Nodes (40): Props, AspectRatio(), Group(), Kbd(), Description(), state(), triggerID(), initialChecked() (+32 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.13
Nodes (29): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpTwice() (+21 more)

### Community 28 - "App"
Cohesion: 0.08
Nodes (9): entityTopics(), formIntQuery(), isHTMX(), redirect(), setAudience(), securityHeaders(), DisplayPage(), Flash() (+1 more)

### Community 29 - "testing.T"
Cohesion: 0.03
Nodes (110): TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), TestMasteryLimit(), TestParseEffects(), TestParseProfs(), TestDropFeat(), TestSheetPinsSource() (+102 more)

### Community 30 - "app.js"
Cohesion: 0.10
Nodes (18): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+10 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "App"
Cohesion: 0.07
Nodes (10): ftsQuery(), Search(), compendiumBack(), ImportSummary(), viewerCampaign(), fiveToolsKey(), sourceInUseMsg(), sourceUsers() (+2 more)

### Community 34 - "Parse"
Cohesion: 0.07
Nodes (32): Defense, feature(), Parse(), tokenize(), fighter(), TestCommands(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+24 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "CN"
Cohesion: 0.08
Nodes (57): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, ListMonstersByNameRow (+49 more)

### Community 38 - "select_templ.go"
Cohesion: 0.12
Nodes (21): Group(), initialValue(), Item(), itemLabel(), Label(), Separator(), state(), Trigger() (+13 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".encounterView"
Cohesion: 0.11
Nodes (15): Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), damageChoice(), initiativeMod(), outOfFight() (+7 more)

### Community 41 - "TestDownloadFiveTools"
Cohesion: 0.13
Nodes (19): internal/compendium, parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestParseFiveToolsBook(), DownloadFiveTools() (+11 more)

### Community 42 - "buildReview"
Cohesion: 0.13
Nodes (10): saveScreenshot(), buildReview(), csvSource(), formValues(), parseCSV(), TestGuessMapping(), saveSniffed(), capReader (+2 more)

### Community 43 - "cell"
Cohesion: 0.13
Nodes (18): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+10 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 46 - ".wizardCommand"
Cohesion: 0.09
Nodes (30): EquipmentItem, CoreTraits(), EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions() (+22 more)

### Community 47 - "ClassInfo"
Cohesion: 0.11
Nodes (11): LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), applyASI(), hpGain() (+3 more)

### Community 48 - "State"
Cohesion: 0.16
Nodes (10): message, patchMessage, stateHandler, viewPatch, Client, Diff(), State, manyKeys() (+2 more)

### Community 49 - "Character"
Cohesion: 0.20
Nodes (8): Campaign(), CharacterFileName(), fileSlug(), CharacterFile, NewCharacterFile(), writeFile(), TestCampaignExcludesOwnerBlocks(), parseCharacterFile()

### Community 50 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

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
Nodes (23): main(), run(), Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap() (+15 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ClassData"
Cohesion: 0.08
Nodes (42): block, ClassFeature, abilityMod(), ParseAidedd(), statLabel(), textLines(), ClassData, firstInt() (+34 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "CommandBuilder"
Cohesion: 0.18
Nodes (14): armorForm(), BuilderPickItems(), BuilderPicks(), builderRemove(), closeAfter(), CommandBuilder(), compendiumLookup(), BuilderPick (+6 more)

### Community 59 - "Book"
Cohesion: 0.15
Nodes (20): bookWriter, mirror, rawBook, bookFolder(), bookGroup(), downloadBook(), fetchBooks(), fetchJSON() (+12 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "ImportView"
Cohesion: 0.24
Nodes (10): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), ImportReviewPage(), importRows() (+2 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.25
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.09
Nodes (25): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Entity, Note, CreateCombatantParams, applyCondOps() (+17 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (116): Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props (+108 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "Slugify"
Cohesion: 0.11
Nodes (21): Cantrips(), dataFor(), Kind, TextData, Slugify(), textMarkdown(), pinnedRefs(), sourceIssue() (+13 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "Session"
Cohesion: 0.11
Nodes (15): fixture, tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken() (+7 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.11
Nodes (21): Player(), Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips() (+13 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - ".Router"
Cohesion: 0.15
Nodes (8): IsLoopback(), TestIsLoopback(), TestRequireAdmin(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "ftDuration"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 86 - "Queries"
Cohesion: 0.10
Nodes (33): featureSwitch, DecodeData(), Queries, CompendiumEntry, Env, ClassFeature, ClassInfo, LevelUpPlan (+25 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Feature"
Cohesion: 0.14
Nodes (11): FeatureEffects, ItemData, SaveFeature(), UsesLabel(), CustomMod, Feature, Uses, compendiumWeapon() (+3 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (35): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+27 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.31
Nodes (14): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Site"
Cohesion: 0.13
Nodes (14): Site, Source, fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page, pageSource() (+6 more)

### Community 104 - ".classFile"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 105 - ".fail"
Cohesion: 0.13
Nodes (10): Character, CanChangeSubclass(), CanCharacter(), TestCanChangeSubclass(), linkEntityLook(), linkCharacterLook(), lookFree(), nextStep() (+2 more)

### Community 106 - ".whoami"
Cohesion: 0.06
Nodes (25): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), folderName() (+17 more)

### Community 107 - "ftBook"
Cohesion: 0.14
Nodes (22): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+14 more)

### Community 108 - "Store"
Cohesion: 0.13
Nodes (21): migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open() (+13 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 111 - "export/markdown.go"
Cohesion: 0.31
Nodes (9): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+1 more)

### Community 112 - "Topic"
Cohesion: 0.15
Nodes (7): event, fakeHandler, Handler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 113 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.48
Nodes (5): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits()

### Community 114 - "Campaign"
Cohesion: 0.24
Nodes (5): Campaign, TestUserAgent(), userAgentSummary(), visitorOnly(), Roster

### Community 115 - "Sheet"
Cohesion: 0.06
Nodes (23): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+15 more)

### Community 116 - "Slugify"
Cohesion: 0.12
Nodes (26): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, SpellData, SizeNamed() (+18 more)

### Community 117 - "Pending"
Cohesion: 0.29
Nodes (4): PendingState, Pending, PendingKind, viewerCtx

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "Hub"
Cohesion: 0.18
Nodes (3): ClientInfo, Hub, randomID()

### Community 123 - ".levelUp"
Cohesion: 0.10
Nodes (25): Spell, addClassSpells(), spellbook(), spellNeed(), subclassPanels(), classPin(), LevelUpView, Option (+17 more)

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "PlanSourceSwitch"
Cohesion: 0.40
Nodes (6): SwitchFeature, SwitchPlan, PlanSourceSwitch(), names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 127 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "alert_templ.go"
Cohesion: 0.24
Nodes (10): ActionProps, DescriptionProps, Props, TitleProps, Variant, Action(), Alert(), Description() (+2 more)

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 134 - "Import"
Cohesion: 0.11
Nodes (34): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+26 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.25
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), featuresSearchText(), ftFeatureKey()

### Community 138 - "Look"
Cohesion: 0.29
Nodes (10): Entity, Look, removeButton(), currentLook(), EntitiesPage(), EntityListView, lookControls(), LookItem() (+2 more)

### Community 141 - "CompendiumSource"
Cohesion: 0.22
Nodes (6): ListEntryCopiesRow, CompendiumSource, copyLabel(), SourceLabel(), ListEntryCopiesParams, EntryView

### Community 142 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 145 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 149 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 375 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **117 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `.blockFor`, `Button`, `perm.go`, `.joinInfo`, `App`, `CN`, `.encounterView`, `buildReview`, `.wizardCommand`, `State`, `cert.go`, `Book`, `FiveToolsJob`, `database/sql.NullInt64`, `Slugify`, `.Router`, `Queries`, `Feature`, `.fail`, `.whoami`, `Topic`, `Campaign`, `Slugify`, `Hub`, `.levelUp`?**
  _High betweenness centrality (0.057) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10338907469342253 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `Parse`, `encoding/json.RawMessage`, `SheetView`, `.encounterView`, `ParseSheet`, `Slugify`, `Command`, `.wizardCommand`, `ClassInfo`, `TestSourceSwitchKeepsEdits`, `Slugify`, `Queries`, `Feature`, `.levelUp`, `PlanSourceSwitch`?**
  _High betweenness centrality (0.019) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.09307244843997885 - nodes in this community are weakly interconnected._
- **Why does `builderFields()` connect `Input` to `compendium_templ.go`, `CombatantView`, `SheetView`, `Button`, `levelUpForm`, `Feature`, `CommandBuilder`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.029850746268656716 - nodes in this community are weakly interconnected._