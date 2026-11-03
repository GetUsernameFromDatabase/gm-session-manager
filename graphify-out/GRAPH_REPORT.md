# Graph Report - gm-session-manager  (2026-11-04)

## Corpus Check
- 247 files · ~494,251 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3579 nodes · 14465 edges · 219 communities (115 shown, 104 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1443 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e9f69aca`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- popover_templ.go
- context.Context
- .do
- createProvider
- .encounterView
- SheetView
- combobox.js
- redirect.go
- testing.T
- github.com/gin-gonic/gin.Context
- select.js
- command.go
- .Entry
- DBTX
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- wikiText
- wikilink.js
- CombatantView
- popover.js
- time.Time
- Project Rules (CLAUDE.md, README)
- ParseSheet
- combobox_templ.go
- Page
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- live
- command.js
- dialog.js
- Sheet
- golang.org/x/net/html.Node
- tooltip.js
- Store
- .levelUp
- Slugify
- export/markdown.go
- composite.js
- DecodeData
- github.com/a-h/templ.Component
- .Plan
- IfElse
- Session
- Client
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- LoadOrCreateCert
- shadcn-templ Config
- ImageResults
- golangci-lint Config
- NewSheet
- monsterBody
- radiogroup.js
- vendorjs Tool
- Feature
- New
- sqlc Config
- State
- Browser check
- Claude Settings
- Pending
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
- cell
- Lint and Release Rules
- Run Script
- inputgroup_templ.go
- useTypeahead
- field_templ.go
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- errBodyWriter
- Site
- github.com/a-h/templ.Attributes
- Campaign
- Input
- Entry
- Import
- SearchReadableBlocksRow
- select_templ.go
- ParseWikidot
- Hub
- .wizardCommand
- radiogroup_templ.go
- Props
- Client
- lanHostname
- build.sh
- .ListBlockStepsSince
- .ListItemsByCategory
- .ListSpellsForClass
- command_templ.go
- .sendFeedback
- Condition
- Note
- card_templ.go
- Character
- toast_templ.go
- alert_templ.go
- open5e.go
- .builderPicks
- Entity
- Props
- NoteBlock
- buildReview
- Handler
- switch_templ.go
- keepawake_windows.go
- CountHabitatsRow
- CompendiumSource
- .ListCompendiumByKind
- .ListCompendiumBySlugs
- keepawake_other.go
- ListImageCharactersRow
- CountCompendiumRow
- ListHomebrewByOwnerRow
- ListHomebrewLinksRow
- ListHomebrewRow
- ListImageEntitiesRow
- parseRowSet
- AddBlockAudienceParams
- AddBlockLinkParams
- AddEntityFieldLinkParams
- AddHabitatByNameParams
- AddMonsterHabitatParams
- BindMonsterImagesParams
- .CreateCampaign
- .CreateEncounter
- CreateLevelUpParams
- EnsureImageCategoryParams
- ExpireTurnConditionsParams
- .GetCompendiumEntryByName
- .GetCompendiumEntryBySlug
- ImageBindingRow
- ImageLookAccessRow
- InsertBlockStepParams
- LevelUp
- LinkCharacterImageParams
- LinkEntityImageParams
- LinkEntryCampaignParams
- .ListBacklinks
- ListBlockDocsRow
- ListCharacterSheetsRow
- ListCompendiumKeysBySourceRow
- ListCompendiumNamesRow
- .ListEncountersByMonster
- .ListEntityFieldBacklinks
- ListExportBlocksRow
- ListImageAttachmentsRow
- .ListImages
- .ListReadableNotes
- ListStepJSONRow
- ListSubclassesOfRow
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
- .UpsertCompendiumEntry
- .UpsertCompendiumSource

## God Nodes (most connected - your core abstractions)
1. `App` - 266 edges
2. `CN()` - 112 edges
3. `Button()` - 102 edges
4. `Sheet` - 89 edges
5. `newTestEnv()` - 86 edges
6. `SheetView` - 51 edges
7. `IfElse()` - 50 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 36 edges

## Surprising Connections (you probably didn't know these)
- `Where things live (creation and level-up)` --references--> `userError`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Next up (pick from these)` --references--> `TestLevelUpFeats()`  [INFERRED]
  HANDOFF.md → internal/web/levelup_test.go
- `Where things live (creation and level-up)` --references--> `SpellNeed`  [INFERRED]
  HANDOFF.md → internal/web/views/levelup_templ.go
- `Where things live (creation and level-up)` --references--> `lostPicks()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go
- `Where things live (creation and level-up)` --references--> `wizardLevelUp()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (219 total, 104 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.11
Nodes (33): contextKey, clientMessage, TestMain(), TestMain(), lastTopLevelColon(), twImportantToV3(), twImportantToV4(), isToggle() (+25 more)

### Community 1 - "popover_templ.go"
Cohesion: 0.22
Nodes (12): state(), Trigger(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 2 - "context.Context"
Cohesion: 0.03
Nodes (3): fetch(), Direction(), WithDirection()

### Community 3 - ".do"
Cohesion: 0.17
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - ".encounterView"
Cohesion: 0.12
Nodes (15): Combatant, Encounter, ListCombatantsRow, Character, combatantDefenses(), initiativeMod(), outOfFight(), pcSheet() (+7 more)

### Community 6 - "SheetView"
Cohesion: 0.07
Nodes (87): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+79 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.06
Nodes (59): TestCanCharacter(), nullID(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName(), TestLanguagesForm(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn() (+51 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.10
Nodes (6): coinCommand(), commandText(), formIntQuery(), isHTMX(), securityHeaders(), Flash()

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "command.go"
Cohesion: 0.07
Nodes (63): Amount, CondOpKind, Effects, Env, HelpEntry, parser, condCommand(), conditionSlug() (+55 more)

### Community 13 - ".Entry"
Cohesion: 0.08
Nodes (31): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, ValidKind(), HomebrewInput (+23 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.08
Nodes (40): GetBlockRow, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanCharacter(), CanDeleteCharacter(), CanEditBlock(), CanEditHomebrew() (+32 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "App"
Cohesion: 0.08
Nodes (10): TopicKind, Entity, Topic, compendiumBack(), entityTopics(), linkEntityLook(), redirect(), HomebrewView (+2 more)

### Community 21 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "CombatantView"
Cohesion: 0.12
Nodes (32): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatantState(), downBadge() (+24 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "time.Time"
Cohesion: 0.07
Nodes (18): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), RequireAdmin(), saveFieldLinks(), browserName() (+10 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.16
Nodes (20): TestUpgradeSheet(), ParseSheet(), TestLevelUpFeats(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp(), TestRevertLevelUpInterleaved(), TestUpgradeSheets() (+12 more)

### Community 28 - "combobox_templ.go"
Cohesion: 0.11
Nodes (22): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, EmptyProps, GroupProps (+14 more)

### Community 29 - "Page"
Cohesion: 0.07
Nodes (62): Label(), Scripts(), Spinner(), campaignLink(), CampaignOverview(), PendingGroup, PendingRequests(), CharacterPage() (+54 more)

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
Cohesion: 0.06
Nodes (33): Queries, ListAttachmentFilesRow, GetImageRow, ListImagesRow, Attachment, CompendiumEntry, CreateCombatantParams, applyCondOps() (+25 more)

### Community 34 - "live"
Cohesion: 0.17
Nodes (19): CampaignRoster(), rosterCard(), localTime(), live(), addBlockForm(), authorRef(), blockMenu(), Author (+11 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Sheet"
Cohesion: 0.06
Nodes (24): DeathSaves, Part, Skill, Completions(), TestCustomMods(), TestModifier(), skillByKey(), Sum() (+16 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Store"
Cohesion: 0.07
Nodes (37): fixture, tokens, migration, isCleared(), newFixture(), serve(), TestMiddlewareIdentity(), TestOnlyHashIsStored() (+29 more)

### Community 41 - ".levelUp"
Cohesion: 0.11
Nodes (31): Where things live (creation and level-up), OptionPrerequisite(), Options(), NameDesc, ClassFeature, ClassInfo, LevelUpPlan, classFeatures() (+23 more)

### Community 42 - "Slugify"
Cohesion: 0.11
Nodes (24): Cantrips(), firstTable(), LevelSpells(), TestOptions(), TextData, Slugify(), Spell, SizeNamed() (+16 more)

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "DecodeData"
Cohesion: 0.08
Nodes (14): DecodeData(), Character, Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile(), compendiumWeapon() (+6 more)

### Community 46 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (109): ListMonstersByNameRow, Action(), Alert(), Description(), Title(), AspectRatio(), Footer(), Checkbox() (+101 more)

### Community 47 - ".Plan"
Cohesion: 0.09
Nodes (17): ClassLevel, Defense, LevelUpChoice, Handoff, Improvements (bigger, deferred), Next up (pick from these), Notes, Test data in data-test (all safe to delete) (+9 more)

### Community 48 - "IfElse"
Cohesion: 0.12
Nodes (39): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+31 more)

### Community 49 - "Session"
Cohesion: 0.16
Nodes (8): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), hashToken(), newToken(), Session, CreateSessionParams

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
Cohesion: 0.15
Nodes (16): main(), run(), Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames() (+8 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ImageResults"
Cohesion: 0.07
Nodes (45): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Fields, Look, ParseFields(), initials() (+37 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "NewSheet"
Cohesion: 0.12
Nodes (20): Parse(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestHealingFromZeroResetsDeathSaves(), TestLanguages(), TestLevelUp() (+12 more)

### Community 59 - "monsterBody"
Cohesion: 0.07
Nodes (50): SpellData, SearchCompendiumRow, Badge(), alignmentColor(), classBadges(), classBody(), commaPrefixed(), CompendiumResults() (+42 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Feature"
Cohesion: 0.13
Nodes (12): FeatureEffects, armorCondition(), columnInts(), ParseEffects(), SaveFeature(), TestParseEffects(), UsesLabel(), CustomMod (+4 more)

### Community 63 - "New"
Cohesion: 0.42
Nodes (10): New(), NewClient(), drain(), TestCatchUp(), TestCloseSessionAndCharacter(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+2 more)

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 65 - "State"
Cohesion: 0.13
Nodes (11): patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff(), deviceLabel() (+3 more)

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

### Community 72 - "builder_templ.go"
Cohesion: 0.06
Nodes (80): ctxKey, ctxState, Props, Defenses, damageChoice(), Dialog(), Close(), CloseFor() (+72 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.22
Nodes (11): categoryRow(), columnLabel(), columnOverview(), ImportCategory, ImportView, ImportReview(), ImportReviewPage(), importRows() (+3 more)

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
Cohesion: 0.23
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 81 - "useFocus"
Cohesion: 0.43
Nodes (6): isInsideEnabledTrigger(), matchesFocusVisible(), useFocus(), onBlur(), onFocus(), resetBlockedFocus()

### Community 82 - "mark_others.js"
Cohesion: 0.52
Nodes (6): applyAttributeToOthers(), buildKeepSet(), collectOutsideElements(), correctElements(), markOthers(), unwrapHost()

### Community 83 - "cell"
Cohesion: 0.12
Nodes (19): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+11 more)

### Community 86 - "inputgroup_templ.go"
Cohesion: 0.17
Nodes (13): Props, Size, AddonProps, Align, ButtonProps, ButtonSize, InputProps, Props (+5 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "field_templ.go"
Cohesion: 0.16
Nodes (15): ContentProps, DescriptionProps, ErrorProps, GroupProps, LabelProps, LegendProps, LegendVariant, Orientation (+7 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.16
Nodes (18): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+10 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 103 - "Site"
Cohesion: 0.12
Nodes (17): Site, Source, dataFor(), Kind, fiveToolsPage(), storeSource(), upsert(), FetchPage() (+9 more)

### Community 104 - "github.com/a-h/templ.Attributes"
Cohesion: 0.06
Nodes (38): Props, Props, ContentProps, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+30 more)

### Community 105 - "Campaign"
Cohesion: 0.18
Nodes (7): Campaign, deviceOS(), isPhone(), uaPick(), userAgentSummary(), joinPath(), visitorOnly()

### Community 106 - "Input"
Cohesion: 0.17
Nodes (22): ctxState, Chip(), Combobox(), Content(), Input(), Item(), itemLabel(), List() (+14 more)

### Community 107 - "Entry"
Cohesion: 0.27
Nodes (9): Entry, fixtures, normalizer, byPK(), children(), joinText(), label(), parseFixtures() (+1 more)

### Community 108 - "Import"
Cohesion: 0.16
Nodes (20): copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource() (+12 more)

### Community 110 - "select_templ.go"
Cohesion: 0.19
Nodes (12): initialValue(), Align, ContentProps, ctxKey, GroupProps, ItemProps, LabelProps, Props (+4 more)

### Community 111 - "ParseWikidot"
Cohesion: 0.14
Nodes (20): block, ClassFeature, CoreTraits(), TestCoreTraits(), ClassData, cellAt(), classTable(), featuresUnder() (+12 more)

### Community 112 - "Hub"
Cohesion: 0.21
Nodes (3): ClientInfo, event, Hub

### Community 113 - ".wizardCommand"
Cohesion: 0.14
Nodes (20): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+12 more)

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.31
Nodes (8): initialValue(), RadioGroup(), state(), DirectionAttributes(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "Props"
Cohesion: 0.67
Nodes (4): attrString(), generateSVG(), getIconContent(), Icon()

### Community 116 - "Client"
Cohesion: 0.17
Nodes (4): fakeHandler, message, Client, topicIDs()

### Community 117 - "lanHostname"
Cohesion: 0.67
Nodes (3): lanHostname(), shortHost(), TestShortHost()

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 125 - "command_templ.go"
Cohesion: 0.17
Nodes (12): ctxKey, ctxState, DialogProps, EmptyProps, GroupProps, InputProps, ItemProps, ListProps (+4 more)

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 127 - "Condition"
Cohesion: 0.18
Nodes (5): Condition, conditionList(), CreateConditionParams, FindCharacterConditionParams, FindCombatantConditionParams

### Community 128 - "Note"
Cohesion: 0.27
Nodes (4): Note, setAudience(), CreateNoteParams, blockAccessForm

### Community 129 - "card_templ.go"
Cohesion: 0.25
Nodes (8): ActionProps, ContentProps, DescriptionProps, FooterProps, HeaderProps, Props, Size, TitleProps

### Community 131 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 132 - "alert_templ.go"
Cohesion: 0.38
Nodes (6): ActionProps, DescriptionProps, Props, TitleProps, Variant, variantClasses()

### Community 133 - "open5e.go"
Cohesion: 0.07
Nodes (30): ArmorData, fixture, looseString, o5eArmor, o5eAssignment, o5eChild, o5eClass, o5eCreature (+22 more)

### Community 136 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 137 - "NoteBlock"
Cohesion: 0.40
Nodes (3): CreateBlockParams, ListReadableBlocksParams, NoteBlock

### Community 138 - "buildReview"
Cohesion: 0.17
Nodes (7): buildReview(), csvSource(), formValues(), parseCSV(), capReader, csvImport, csvImports

### Community 140 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **104 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Note`, `context.Context`, `.do`, `.encounterView`, `.builderPicks`, `github.com/gin-gonic/gin.Context`, `buildReview`, `perm.go`, `time.Time`, `database/sql.NullInt64`, `Store`, `.levelUp`, `Slugify`, `DecodeData`, `LoadOrCreateCert`, `State`, `Campaign`, `Hub`, `.wizardCommand`?**
  _High betweenness centrality (0.034) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11037699183832103 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `.encounterView`, `SheetView`, `builder_templ.go`, `.levelUp`, `Slugify`, `command.go`, `DecodeData`, `.Plan`, `.wizardCommand`, `NewSheet`, `ParseSheet`, `Page`, `Feature`?**
  _High betweenness centrality (0.023) - this node is a cross-community bridge._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.031746031746031744 - nodes in this community are weakly interconnected._
- **Why does `capReader` connect `buildReview` to `compendium_templ.go`?**
  _High betweenness centrality (0.008) - this node is a cross-community bridge._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._