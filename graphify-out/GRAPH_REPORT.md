# Graph Report - gm-session-manager  (2026-11-11)

## Corpus Check
- 296 files · ~584,483 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4184 nodes · 18170 edges · 227 communities (120 shown, 107 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1876 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `3dfd3db5`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- CombatantView
- context.Context
- .do
- createProvider
- Weapon
- SheetView
- combobox.js
- redirect.go
- .loadLive
- Import
- select.js
- Command
- .Entry
- ItemData
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- dialog_templ.go
- levelUpForm
- wikilink.js
- IfElse
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- ParseSheet
- .found
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- testing.T
- command.js
- dialog.js
- Client
- Queries
- tooltip.js
- github.com/gin-gonic/gin.Context
- ParseAidedd
- Book
- buildReview
- composite.js
- App
- Feature
- monsterBody
- Button
- Topic
- cell
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- ftSpell
- TestDownloadFiveTools
- radiogroup.js
- vendorjs/main.go
- Input
- FiveToolsJob
- sqlc Config
- fromLookup
- Browser check
- Claude Settings
- .campaignSourcesSet
- use_transition_status.js
- playwright
- release.sh
- .levelUp
- checkbox.js
- .serverError
- switch.js
- Player
- tabs.js
- portal.js
- useClick
- .wizardCommand
- useFocus
- mark_others.js
- ImagePage
- Lint and Release Rules
- Run Script
- RosterCharacter
- useTypeahead
- Slugify
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- joinNonEmpty
- mounted
- NewClient
- planSwitch
- Session
- State
- .sendFeedback
- ftBook
- InputForEntry
- formInt
- formErrors
- Parse
- Open
- UpsertCompendiumEntryParams
- Pending
- Character
- ftItemData
- toast_templ.go
- build.sh
- .Router
- Character
- errBodyWriter
- .sourceIssues
- buildReview
- CompendiumSource
- Before/after pairs
- PlanSourceSwitch
- open5e.go
- DBTX
- db/CLAUDE.md
- perm/CLAUDE.md
- ftPrerequisite
- Entity
- Site
- keepawake_windows.go
- CountHabitatsRow
- SearchReadableBlocksRow
- ParseWiki
- NoteBlock
- .ListCompendiumNames
- .ListSharedOrder
- lanHostname
- .CountSharedByKind
- .ListBlockStepsSince
- .ListCompendiumByKind
- .ListCompendiumBySlugs
- .ListEntrySourcesBySlugs
- .ListItemsByCategory
- .ListSharedEntries
- .ListSpellsForClass
- .ListSubracesOf
- keepawake_other.go
- Client
- ListImageCharactersRow
- ListSharedWithRow
- CountCompendiumRow
- TestRunCommandUsesCurrentSheet
- ListHomebrewByOwnerRow
- BindMonsterImagesParams
- ListMasteryWeaponsRow
- ListImageEntitiesRow
- AddBlockAudienceParams
- AddEntityFieldLinkParams
- AddMonsterHabitatParams
- AllowCampaignSourceParams
- BindMonsterImageParams
- .CreateCampaign
- .CreateEncounter
- CreateHomebrewParams
- CreateLevelUpParams
- EnsureImageCategoryParams
- ExpireTurnConditionsParams
- .GetCompendiumEntryByName
- .GetCompendiumEntryBySlug
- .GetCompendiumEntryBySlugSource
- ImageBindingRow
- ImageLookAccessRow
- InsertBlockStepParams
- LevelUp
- LinkEntityImageParams
- LinkEntryCampaignParams
- .ListBacklinks
- ListBlockDocsRow
- ListCharacterEncountersParams
- ListCharacterSheetsNamedRow
- ListCharacterSheetsRow
- ListCompendiumKeysBySourceRow
- .ListEncountersByMonster
- .ListEntityFieldBacklinks
- ListImageAttachmentsRow
- .ListImages
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
- .SeedMonsterHabitats
- .SetBlockPosition
- .SetBlockVersion
- .SetCombatantHP
- .SetCombatantMarker
- SetCompendiumImageParams
- SetEntityCurrentImageParams
- ToggleCharacterImageRevealedParams
- ToggleEntityImageRevealedParams
- .TrimBlockSteps
- UnlinkCharacterImageParams
- UnlinkEntityImageParams
- .UpdateBlockAccess
- .UpdateCharacter
- .UpdateCharacterSheet
- UpdateEntityParams
- .UpsertCompendiumSource

## God Nodes (most connected - your core abstractions)
1. `App` - 314 edges
2. `newTestEnv()` - 125 edges
3. `Button()` - 121 edges
4. `Sheet` - 118 edges
5. `CN()` - 113 edges
6. `IfElse()` - 61 edges
7. `SheetView` - 57 edges
8. `Command` - 49 edges
9. `Input()` - 47 edges
10. `Slugify()` - 44 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `internal/compendium` --references--> `DownloadFiveTools()`  [INFERRED]
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

## Communities (227 total, 107 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (63): contextKey, classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild, o5eClass (+55 more)

### Community 1 - "CombatantView"
Cohesion: 0.07
Nodes (61): Input(), Addon(), Button(), Input(), InputGroup(), Content(), Description(), Header() (+53 more)

### Community 2 - "context.Context"
Cohesion: 0.02
Nodes (15): ListHomebrewLinksRow, ListHomebrewRow, uniqueName(), AddBlockLinkParams, AddHabitatByNameParams, DisallowCampaignSourceParams, LinkCharacterImageParams, ListExportBlocksRow (+7 more)

### Community 3 - ".do"
Cohesion: 0.11
Nodes (22): isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), Role, testEnv, pendingSecret(), TestClaimNeedsApproval() (+14 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.05
Nodes (24): DeathSaves, Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey() (+16 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (98): ArmorEffect(), Signed(), Field, Variant, Action(), Card(), Content(), Description() (+90 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - ".loadLive"
Cohesion: 0.07
Nodes (59): Props, Variant, MediaProps, MediaVariant, CountSharedBySourceRow, ListMonstersByNameRow, Encounter, Badge() (+51 more)

### Community 10 - "Import"
Cohesion: 0.13
Nodes (27): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+19 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (65): Amount, CondOpKind, Effects, HelpEntry, parser, Completions(), condCommand(), conditionSlug() (+57 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (25): MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInput, HomebrewInputFor(), abilityKey() (+17 more)

### Community 14 - "ItemData"
Cohesion: 0.16
Nodes (15): ImageCategory, categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), importRows() (+7 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.08
Nodes (44): GetBlockRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass(), CanCharacter() (+36 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "dialog_templ.go"
Cohesion: 0.10
Nodes (46): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+38 more)

### Community 21 - "levelUpForm"
Cohesion: 0.10
Nodes (40): Props, Footer(), Checkbox(), initialChecked(), initialValue(), Item(), RadioGroup(), state() (+32 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "IfElse"
Cohesion: 0.09
Nodes (44): Entity, Look, Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant() (+36 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (180): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+172 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.13
Nodes (30): ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations(), TestLevelUpLineage(), TestLevelUpTwice() (+22 more)

### Community 28 - ".found"
Cohesion: 0.07
Nodes (18): ftsQuery(), Search(), Campaign, compendiumBack(), compendiumPath(), pathCampaign(), viewerCampaign(), isHTMX() (+10 more)

### Community 29 - "newTestEnv"
Cohesion: 0.05
Nodes (61): nullID(), TestBuilderFeat(), TestBuilderInventory(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestMonsterPortrait() (+53 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "database/sql.NullInt64"
Cohesion: 0.05
Nodes (41): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Condition, Note, CreateCombatantParams, ListCombatantsRow, applyCondOps() (+33 more)

### Community 34 - "testing.T"
Cohesion: 0.04
Nodes (69): CoreTraits(), TestCoreTraits(), RemoveSourceDir(), TestRemoveSourceDir(), roundTrip(), TestParseWikidotClass(), TestParseWikidotFeatAndBackground(), TestParseWikidotSpecies() (+61 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Client"
Cohesion: 0.21
Nodes (3): ClientInfo, Hub, randomID()

### Community 38 - "Queries"
Cohesion: 0.07
Nodes (36): LevelUpChoice, Queries, CompendiumEntry, TestSlotMaxima(), applyASI(), ClassFeature, ClassInfo, LevelUpPlan (+28 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (22): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), featureCommand(), coinCommand(), commandText() (+14 more)

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - "Book"
Cohesion: 0.15
Nodes (18): bookWriter, mirror, rawBook, internal/compendium, bookFolder(), bookGroup(), downloadBook(), fetchBooks() (+10 more)

### Community 43 - "buildReview"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "App"
Cohesion: 0.08
Nodes (11): Attachment, firstLine(), mdLinkText(), saveScreenshot(), inCampaign(), parseCSV(), TestGuessMapping(), saveSniffed() (+3 more)

### Community 46 - "Feature"
Cohesion: 0.12
Nodes (17): FeatureEffects, ItemData, armorCondition(), columnInts(), ParseEffects(), parseProfs(), SaveFeature(), UsesLabel() (+9 more)

### Community 47 - "monsterBody"
Cohesion: 0.09
Nodes (38): SearchCompendiumRow, alignmentColor(), classBadges(), classBody(), commaPrefixed(), compendiumFindURL(), compendiumRow(), copyOrder() (+30 more)

### Community 48 - "Button"
Cohesion: 0.06
Nodes (79): UI conventions, Button(), Label(), Scripts(), Separator(), Spinner(), campaignLink(), CampaignOverview() (+71 more)

### Community 49 - "Topic"
Cohesion: 0.14
Nodes (7): event, Handler, stateHandler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 50 - "cell"
Cohesion: 0.13
Nodes (18): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+10 more)

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

### Community 56 - "ClassData"
Cohesion: 0.19
Nodes (14): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+6 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "ftSpell"
Cohesion: 0.67
Nodes (4): ftDistance, ftAmount(), ftDuration(), ftRange()

### Community 59 - "TestDownloadFiveTools"
Cohesion: 0.17
Nodes (19): parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants(), TestParseFiveToolsBook(), DownloadFiveTools() (+11 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Input"
Cohesion: 0.06
Nodes (75): Combobox(), Content(), Input(), Item(), List(), state(), CloseFor(), Description() (+67 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.25
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "fromLookup"
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - ".campaignSourcesSet"
Cohesion: 0.15
Nodes (8): ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers(), FiveToolsView, NewFiveToolsView(), FiveToolsBook, FiveToolsGroup

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - ".levelUp"
Cohesion: 0.10
Nodes (21): ListSubclassesOfRow, Spell, spellbook(), spellFromRow(), spellNeed(), subclassIndex(), subclassOptions(), subclassPanels() (+13 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - ".serverError"
Cohesion: 0.08
Nodes (6): CharacterFile, parseCharacterFile(), setAudience(), Flash(), Toast(), blockAccessForm

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Player"
Cohesion: 0.12
Nodes (18): Gotchas, internal/web, Notes, Websocket and live views, newLiveFixture(), TestLiveBindsMatchState(), TestLiveRefetchStrips(), TestLiveSourcesAlert() (+10 more)

### Community 77 - "tabs.js"
Cohesion: 0.56
Nodes (8): activationDirection(), activeValue(), destroy(), init(), isDisabled(), partsOf(), requestValueChange(), setActiveTab()

### Community 78 - "portal.js"
Cohesion: 0.39
Nodes (6): containerFor(), createOutsideGuards(), onFocus(), remove(), render(), setFocusManagerState()

### Community 79 - "useClick"
Cohesion: 0.39
Nodes (5): useClick(), getNextOpen(), onClick(), onMouseDown(), setOpenWithTouchDelay()

### Community 80 - ".wizardCommand"
Cohesion: 0.06
Nodes (41): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+33 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "ImagePage"
Cohesion: 0.11
Nodes (28): ListEntityFieldBacklinksRow, ListBacklinksRow, attrString(), generateSVG(), getIconContent(), Props, Icon(), EntityDetail() (+20 more)

### Community 86 - "RosterCharacter"
Cohesion: 0.50
Nodes (4): deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "Slugify"
Cohesion: 0.10
Nodes (34): Cantrips(), firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, Slugify() (+26 more)

### Community 100 - "joinNonEmpty"
Cohesion: 0.07
Nodes (42): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+34 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.30
Nodes (15): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+7 more)

### Community 103 - "planSwitch"
Cohesion: 0.23
Nodes (6): featureSwitch, featureNames(), SubclassRef, PlanSourceSwitch(), planSwitch(), sameSlug()

### Community 104 - "Session"
Cohesion: 0.15
Nodes (11): fixture, tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken() (+3 more)

### Community 105 - "State"
Cohesion: 0.15
Nodes (9): fakeHandler, message, patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff() (+1 more)

### Community 106 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (19): ftBook, ftRec, each(), ftCastingTime(), ftCopyTargets(), ftDecode(), ftMod(), ftName() (+11 more)

### Community 108 - "InputForEntry"
Cohesion: 0.22
Nodes (8): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), InputForEntry(), TestParseFiveTools(), TestMonsterFormRoundTrip()

### Community 109 - "formInt"
Cohesion: 0.10
Nodes (9): Combatant, Entity, Character, damageChoice(), entityTopics(), linkEntityLook(), formInt(), linkCharacterLook() (+1 more)

### Community 110 - "formErrors"
Cohesion: 0.33
Nodes (4): formErrors(), fieldError, pickError, userError

### Community 111 - "Parse"
Cohesion: 0.09
Nodes (22): Defense, feature(), TestFiveToolsClasses(), Parse(), tokenize(), fighter(), TestCommands(), TestDefenses() (+14 more)

### Community 112 - "Open"
Cohesion: 0.13
Nodes (22): migration, newFixture(), TestOnlyHashIsStored(), NewManager(), apply(), loadMigrations(), Migrate(), Open() (+14 more)

### Community 114 - "Pending"
Cohesion: 0.25
Nodes (4): PendingState, Pending, PendingKind, viewerCtx

### Community 115 - "Character"
Cohesion: 0.17
Nodes (7): Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), Roster

### Community 116 - "ftItemData"
Cohesion: 0.28
Nodes (8): ArmorData, ftItem, o5eArmor, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData(), armorText()

### Community 117 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - ".Router"
Cohesion: 0.19
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - ".sourceIssues"
Cohesion: 0.31
Nodes (8): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), featKey(), pinFeat(), sheetRef

### Community 126 - "buildReview"
Cohesion: 0.18
Nodes (6): buildReview(), csvSource(), formValues(), capReader, csvImport, csvImports

### Community 127 - "CompendiumSource"
Cohesion: 0.22
Nodes (7): ListEntryCopiesRow, CompendiumSource, ListEntryCopiesParams, CompareCopy, CompareView, EntryView, HabitatCopy

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "PlanSourceSwitch"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "ftPrerequisite"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Site"
Cohesion: 0.09
Nodes (25): HomebrewOwner, Site, Source, TestHomebrew(), fiveToolsPage(), bindFiveToolsImages(), Report, ImportProgress() (+17 more)

### Community 144 - "ParseWiki"
Cohesion: 0.23
Nodes (11): wikiLink(), wikiText(), ImageSrc(), CompendiumRefs(), ParseWiki(), ReplaceCompendiumLinks(), TestParseWiki(), TestReplaceCompendiumLinks() (+3 more)

### Community 145 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 148 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 159 - "Client"
Cohesion: 0.22
Nodes (3): IsLoopback(), TestIsLoopback(), Client

### Community 163 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 387 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **107 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `.do`, `.loadLive`, `perm.go`, `.found`, `database/sql.NullInt64`, `Client`, `Queries`, `github.com/gin-gonic/gin.Context`, `Feature`, `Topic`, `cert.go`, `FiveToolsJob`, `.campaignSourcesSet`, `.levelUp`, `.serverError`, `.wizardCommand`, `Slugify`, `State`, `formInt`, `Pending`, `Character`, `.Router`, `.sourceIssues`, `buildReview`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11082670356278104 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Weapon` to `compendium_templ.go`, `database/sql.NullInt64`, `testing.T`, `TestRunCommandUsesCurrentSheet`, `joinNonEmpty`, `Queries`, `planSwitch`, `.levelUp`, `SheetView`, `Command`, `formInt`, `Feature`, `.wizardCommand`, `Slugify`, `ParseSheet`, `.sourceIssues`?**
  _High betweenness centrality (0.021) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.07027027027027027 - nodes in this community are weakly interconnected._
- **Why does `Pending` connect `Pending` to `compendium_templ.go`, `database/sql.NullInt64`, `Client`, `Client`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.02197802197802198 - nodes in this community are weakly interconnected._