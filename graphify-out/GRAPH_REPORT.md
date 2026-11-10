# Graph Report - gm-session-manager  (2026-11-10)

## Corpus Check
- 296 files · ~568,822 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 53 file(s) not represented in the graph (top: .templ 45, .css 4, (none) 3)

## Summary
- 4100 nodes · 17755 edges · 230 communities (116 shown, 114 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1794 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `012de000`
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
- formInt
- Store
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
- .sendFeedback
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- newTestEnv
- app.js
- tabbable.js
- useDismiss
- Sheet
- Parse
- command.js
- dialog.js
- Client
- ftPrerequisite
- tooltip.js
- CN
- ParseAidedd
- Book
- cell
- composite.js
- Item
- Feature
- ClassInfo
- Badge
- Topic
- Queries
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- ClassData
- golangci-lint Config
- FromContext
- Campaign
- radiogroup.js
- vendorjs/main.go
- dialog_templ.go
- FiveToolsJob
- sqlc Config
- database/sql.NullInt64
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
- use_transition_status.js
- playwright
- release.sh
- ImagePage
- checkbox.js
- .subclassChangeForm
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
- popover_templ.go
- useTypeahead
- spellCounts
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- NewClient
- Site
- Session
- .sourceIssues
- .Router
- ftBook
- .ListBlockStepsSince
- .ListCompendiumByKind
- formErrors
- export/markdown.go
- State
- TestRunCommandUsesCurrentSheet
- Pending
- SetEntityCurrentImageParams
- .ListCompendiumBySlugs
- sourcesEnv
- build.sh
- Hub
- Condition
- errBodyWriter
- CampaignSources
- buildReview
- .ListItemsByCategory
- Before/after pairs
- lanHostname
- open5e.go
- Store
- db/CLAUDE.md
- perm/CLAUDE.md
- ftNamed
- .ListSpellsForClass
- .blockFor
- parseCommand
- .CreateCampaign
- DBTX
- Character
- CompendiumSource
- Entity
- ftSpell
- WeaponKind
- select_templ.go
- NoteBlock
- keepawake_windows.go
- sameSlug
- CountHabitatsRow
- SearchReadableBlocksRow
- .ListCompendiumNames
- keepawake_other.go
- ListImageCharactersRow
- CountCompendiumRow
- ListHomebrewByOwnerRow
- ListHomebrewLinksRow
- ListHomebrewRow
- UpsertCompendiumEntryParams
- ListImageEntitiesRow
- Verbs
- AddBlockAudienceParams
- AddBlockLinkParams
- AddHabitatByNameParams
- AddMonsterHabitatParams
- AllowCampaignSourceParams
- BindMonsterImageParams
- BindMonsterImagesParams
- .CreateEncounter
- CreateLevelUpParams
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
- .UpdateHomebrew
- .UpsertCompendiumSource

## God Nodes (most connected - your core abstractions)
1. `App` - 303 edges
2. `newTestEnv()` - 122 edges
3. `Sheet` - 118 edges
4. `Button()` - 117 edges
5. `CN()` - 113 edges
6. `SheetView` - 57 edges
7. `IfElse()` - 56 edges
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
- `UI conventions` --references--> `selectItems()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go
- `UI conventions` --references--> `formSelect()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/views/combat_templ.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (230 total, 114 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (61): contextKey, classRef, ftBookInfo, indexed, looseString, o5eAssignment, o5eChild, o5eClass (+53 more)

### Community 1 - "CombatantView"
Cohesion: 0.09
Nodes (41): Defenses, damageChoice(), capitalized(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName() (+33 more)

### Community 2 - "context.Context"
Cohesion: 0.03
Nodes (4): uniqueName(), AddEntityFieldLinkParams, LinkEntityImageParams, UpdateEntityParams

### Community 3 - ".do"
Cohesion: 0.13
Nodes (18): Role, testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm() (+10 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "Weapon"
Cohesion: 0.09
Nodes (12): Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield(), damageLabel() (+4 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (94): ArmorEffect(), Slot, Signed(), Field, Look, Variant, Action(), Card() (+86 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.25
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "formInt"
Cohesion: 0.17
Nodes (5): Combatant, Character, entityTopics(), formInt(), lookOps

### Community 10 - "Store"
Cohesion: 0.11
Nodes (30): copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes() (+22 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.10
Nodes (46): Amount, CondOpKind, Effects, HelpEntry, parser, condCommand(), conditionSlug(), Damage() (+38 more)

### Community 13 - ".Entry"
Cohesion: 0.10
Nodes (26): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInput, HomebrewInputFor() (+18 more)

### Community 14 - "Button"
Cohesion: 0.07
Nodes (73): UI conventions, Button(), Input(), Input(), Label(), Scripts(), Separator(), Spinner() (+65 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.06
Nodes (54): ListCombatantsRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanChangeSubclass(), CanDeleteCharacter() (+46 more)

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
Cohesion: 0.05
Nodes (85): ctxState, Coins, Chip(), Chips(), ChipsInput(), Combobox(), Content(), Input() (+77 more)

### Community 21 - "levelUpForm"
Cohesion: 0.07
Nodes (57): LevelUpPlan, initialValue(), Item(), RadioGroup(), state(), DirectionAttributes(), abilitySelect(), LevelUpView (+49 more)

### Community 22 - "wikilink.js"
Cohesion: 0.06
Nodes (36): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+28 more)

### Community 23 - "tabs_templ.go"
Cohesion: 0.13
Nodes (32): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+24 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.09
Nodes (41): Import(), TestUpgradeSheet(), ParseSheet(), TestBuilderCompendium(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI() (+33 more)

### Community 28 - "App"
Cohesion: 0.05
Nodes (28): ftsQuery(), Search(), Campaign, Character, CanCharacter(), compendiumBack(), viewerCampaign(), fiveToolsKey() (+20 more)

### Community 29 - "newTestEnv"
Cohesion: 0.04
Nodes (97): TestTokensReadAsLabels(), InputForEntry(), RemoveSourceDir(), TestRemoveSourceDir(), TestMonsterFormRoundTrip(), TestMasteryLimit(), TestDropFeat(), TestSheetPinsSource() (+89 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "Sheet"
Cohesion: 0.04
Nodes (88): EquipmentItem, DeathSaves, Cantrips(), CoreTraits(), EquipmentOptions(), firstTable(), EquipmentOption, LevelSpells() (+80 more)

### Community 34 - "Parse"
Cohesion: 0.06
Nodes (37): feature(), Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves() (+29 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 38 - "ftPrerequisite"
Cohesion: 0.25
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), featuresSearchText(), ftFeatureKey()

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CN"
Cohesion: 0.06
Nodes (65): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, ContentProps, DescriptionProps (+57 more)

### Community 41 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 42 - "Book"
Cohesion: 0.12
Nodes (25): bookWriter, mirror, rawBook, internal/compendium, bookFolder(), bookGroup(), downloadBook(), DownloadFiveTools() (+17 more)

### Community 43 - "cell"
Cohesion: 0.12
Nodes (19): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+11 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Item"
Cohesion: 0.24
Nodes (6): findItem(), parseArmor(), parseItem(), TestGearCommands(), ensureItem(), Item

### Community 46 - "Feature"
Cohesion: 0.14
Nodes (10): FeatureEffects, parseFeatureSet(), parseUses(), SaveFeature(), UsesLabel(), CustomMod, Feature, Uses (+2 more)

### Community 47 - "ClassInfo"
Cohesion: 0.13
Nodes (10): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+2 more)

### Community 48 - "Badge"
Cohesion: 0.06
Nodes (61): Props, Variant, SearchCompendiumRow, Badge(), badgeClasses(), alignmentColor(), classBadges(), classBody() (+53 more)

### Community 49 - "Topic"
Cohesion: 0.15
Nodes (7): event, Handler, stateHandler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 50 - "Queries"
Cohesion: 0.09
Nodes (34): featureSwitch, Queries, CompendiumEntry, Env, TestSlotMaxima(), ClassInfo, ClassLevel, CasterLevel() (+26 more)

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
Cohesion: 0.16
Nodes (9): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestKeyFileIsPrivate(), TestLoadOrCreateCert() (+1 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ClassData"
Cohesion: 0.12
Nodes (26): block, ClassFeature, ClassData, dnd5eFeatures(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+18 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "FromContext"
Cohesion: 0.08
Nodes (19): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), browserName(), firstLine(), folderName() (+11 more)

### Community 59 - "Campaign"
Cohesion: 0.32
Nodes (5): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "dialog_templ.go"
Cohesion: 0.13
Nodes (41): ctxKey, ctxState, Props, Dialog(), Close(), CloseFor(), Content(), Description() (+33 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "database/sql.NullInt64"
Cohesion: 0.05
Nodes (37): ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, Encounter, Entity, Note, CreateCombatantParams (+29 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (129): Props, Props, Size, ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps (+121 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.83
Nodes (3): die(), newer(), release.sh script

### Community 72 - "ImagePage"
Cohesion: 0.16
Nodes (20): ListEntityFieldBacklinksRow, ListBacklinksRow, attrString(), generateSVG(), getIconContent(), Props, Icon(), EntityPage() (+12 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - ".subclassChangeForm"
Cohesion: 0.50
Nodes (4): ListSubclassesOfRow, subclassIndex(), subclassOptions(), subclassPanels()

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
Cohesion: 0.28
Nodes (9): Config, lookupFunc, describe(), fromLookup(), Load(), lookupMap(), TestFromLookup(), TestLoadEnvFileAndPrecedence() (+1 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 86 - "popover_templ.go"
Cohesion: 0.16
Nodes (21): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+13 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "spellCounts"
Cohesion: 0.50
Nodes (4): sortedColumns(), SpellColumn(), spellCounts(), TestSpellCounts()

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (39): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+31 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "NewClient"
Cohesion: 0.25
Nodes (14): message, New(), NewClient(), drain(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest() (+6 more)

### Community 103 - "Site"
Cohesion: 0.15
Nodes (13): Site, Source, fiveToolsPage(), storeSource(), upsert(), FetchPage(), Page, pageSource() (+5 more)

### Community 104 - "Session"
Cohesion: 0.16
Nodes (9): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken(), newToken(), Session (+1 more)

### Community 105 - ".sourceIssues"
Cohesion: 0.53
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 106 - ".Router"
Cohesion: 0.16
Nodes (7): IsLoopback(), TestIsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 107 - "ftBook"
Cohesion: 0.17
Nodes (18): ftBook, ftRec, each(), ftCastingTime(), ftDecode(), ftMod(), ftName(), ftParts() (+10 more)

### Community 110 - "formErrors"
Cohesion: 0.40
Nodes (3): fieldError, pickError, userError

### Community 111 - "export/markdown.go"
Cohesion: 0.15
Nodes (17): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+9 more)

### Community 112 - "State"
Cohesion: 0.16
Nodes (11): patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff(), deviceLabel(), rosterHP() (+3 more)

### Community 113 - "TestRunCommandUsesCurrentSheet"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 114 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 117 - "sourcesEnv"
Cohesion: 0.32
Nodes (6): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestSubclassPickBySlug()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "Hub"
Cohesion: 0.25
Nodes (3): ClientInfo, Hub, TestCloseSessionAndCharacter()

### Community 123 - "Condition"
Cohesion: 0.15
Nodes (8): Defense, Condition, ParseDefenses(), combatantDefenses(), petrified(), CreateConditionParams, FindCharacterConditionParams, FindCombatantConditionParams

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "CampaignSources"
Cohesion: 0.27
Nodes (10): toJSON(), campaignSourceRow(), CampaignSources(), CampaignSourcesView, sourceGroupButton(), sourceRowTitle(), sourceMoveButton(), sourceMoveURL() (+2 more)

### Community 126 - "buildReview"
Cohesion: 0.17
Nodes (7): buildReview(), csvSource(), formValues(), parseCSV(), capReader, csvImport, csvImports

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 134 - "Store"
Cohesion: 0.07
Nodes (44): fixture, tokens, main(), run(), migration, isCleared(), newFixture(), serve() (+36 more)

### Community 137 - "ftNamed"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 139 - ".blockFor"
Cohesion: 0.24
Nodes (3): GetBlockRow, blockAccessLevel, stepsPayload

### Community 140 - "parseCommand"
Cohesion: 0.22
Nodes (10): findMod(), newMod(), parseMod(), IsModTarget(), featureCommand(), weaponCommand(), coinCommand(), commandText() (+2 more)

### Community 144 - "CompendiumSource"
Cohesion: 0.22
Nodes (6): ListEntryCopiesRow, CompendiumSource, copyLabel(), SourceLabel(), ListEntryCopiesParams, EntryView

### Community 146 - "ftSpell"
Cohesion: 0.50
Nodes (5): ftDistance, ftSpell, ftAmount(), ftDuration(), ftRange()

### Community 147 - "WeaponKind"
Cohesion: 0.50
Nodes (4): Completions(), parseMastery(), TestWeaponKind(), WeaponKind()

### Community 148 - "select_templ.go"
Cohesion: 0.05
Nodes (51): ImageCategory, Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select() (+43 more)

### Community 149 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

## Knowledge Gaps
- **15 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 383 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **114 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `.do`, `Store`, `formInt`, `.blockFor`, `perm.go`, `Sheet`, `CN`, `Book`, `Topic`, `Queries`, `FromContext`, `Campaign`, `FiveToolsJob`, `database/sql.NullInt64`, `.subclassChangeForm`, `.sourceIssues`, `.Router`, `State`, `Hub`, `buildReview`?**
  _High betweenness centrality (0.042) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10722408026755853 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `CombatantView`, `Parse`, `encoding/json.RawMessage`, `Weapon`, `SheetView`, `formInt`, `.sourceIssues`, `Item`, `Feature`, `ClassInfo`, `perm.go`, `TestRunCommandUsesCurrentSheet`, `Queries`, `WeaponKind`, `Input`, `levelUpForm`, `ParseSheet`?**
  _High betweenness centrality (0.009) - this node is a cross-community bridge._
- **Should `CombatantView` be split into smaller, more focused modules?**
  _Cohesion score 0.08636363636363636 - nodes in this community are weakly interconnected._
- **Why does `Feature` connect `Feature` to `compendium_templ.go`, `Sheet`, `Weapon`, `SheetView`, `Queries`, `Input`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.02857142857142857 - nodes in this community are weakly interconnected._