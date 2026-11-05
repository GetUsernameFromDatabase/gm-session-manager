# Graph Report - gm-session-manager  (2026-11-05)

## Corpus Check
- 250 files · ~503,173 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 51 file(s) not represented in the graph (top: .templ 43, .css 4, (none) 3)

## Summary
- 3625 nodes · 14807 edges · 132 communities (108 shown, 24 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 1481 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `fe026ca5`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- FromContext
- context.Context
- .do
- createProvider
- github.com/a-h/templ.Component
- SheetView
- combobox.js
- redirect.go
- testing.T
- github.com/gin-gonic/gin.Context
- select.js
- Command
- .Entry
- rowFigure
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- Page
- .levelUp
- wikilink.js
- EncounterView
- popover.js
- wikiText
- Project Rules (CLAUDE.md, README)
- ParseSheet
- App
- monsterBody
- app.js
- tabbable.js
- useDismiss
- database/sql.NullInt64
- State
- command.js
- dialog.js
- combobox_templ.go
- golang.org/x/net/html.Node
- tooltip.js
- .issueTokens
- NameDesc
- DecodeData
- export/markdown.go
- composite.js
- Encounter
- Client
- .Plan
- tabs_templ.go
- Topic
- New
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- Campaign
- golangci-lint Config
- Slugify
- toast_templ.go
- radiogroup.js
- vendorjs Tool
- select_templ.go
- Pending
- sqlc Config
- .deleteUnused
- Browser check
- Claude Settings
- github.com/a-h/templ.Attributes
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
- ParseWikidot
- Lint and Release Rules
- Run Script
- levelUpForm
- useTypeahead
- ImageResults
- Validator Rule
- Go Module
- live.js
- internal_backdrop.js
- encoding/json.RawMessage
- mounted
- Hub
- Site
- importSource
- Props
- LevelUpView
- Entry
- Import
- .ListSpellsForClass
- .wizardCommand
- RequireAdmin
- Feature
- radiogroup_templ.go
- Weapon
- capReader
- build.sh
- popover_templ.go
- Sheet
- errBodyWriter
- Parse
- .sendFeedback
- .mountAssets
- switch_templ.go
- ArmorData
- Import
- SearchReadableBlocksRow
- .ListCompendiumByKind
- .ListCompendiumBySlugs

## God Nodes (most connected - your core abstractions)
1. `App` - 267 edges
2. `CN()` - 112 edges
3. `Button()` - 105 edges
4. `newTestEnv()` - 91 edges
5. `Sheet` - 90 edges
6. `SheetView` - 52 edges
7. `IfElse()` - 51 edges
8. `Input()` - 46 edges
9. `Command` - 44 edges
10. `formSelect()` - 37 edges

## Surprising Connections (you probably didn't know these)
- `Lessons` --references--> `formSelect()`  [INFERRED]
  HANDOFF.md → internal/web/views/combat_templ.go
- `Creation and level-up pointers` --references--> `spellNeed()`  [INFERRED]
  HANDOFF.md → internal/web/character_handlers.go
- `Creation and level-up pointers` --references--> `WizardView`  [INFERRED]
  HANDOFF.md → internal/web/views/wizard_templ.go
- `Creation and level-up pointers` --references--> `meetsPrerequisite()`  [INFERRED]
  HANDOFF.md → internal/web/wizard.go
- `Creation and level-up pointers` --references--> `errorAt()`  [INFERRED]
  HANDOFF.md → internal/web/wizard_test.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (132 total, 24 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.09
Nodes (65): Props, contextKey, keepAwake(), looseString, o5eAssignment, o5eChild, o5eClass, o5eCreature (+57 more)

### Community 1 - "FromContext"
Cohesion: 0.06
Nodes (20): Identity, Role, ListImagesRow, CompendiumSource, browserName(), firstLine(), mdLinkText(), deviceKey() (+12 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (68): keepAwake(), IsLoopback(), TestIsLoopback(), fetch(), ListImageCharactersRow, CountCompendiumRow, ListHomebrewLinksRow, ListImageEntitiesRow (+60 more)

### Community 3 - ".do"
Cohesion: 0.17
Nodes (16): testEnv, pendingSecret(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder(), TestCommandPermissions(), TestHPQuickForm(), TestJoinBurst() (+8 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "github.com/a-h/templ.Component"
Cohesion: 0.08
Nodes (67): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+59 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (103): ArmorEffect(), Slot, Signed(), Field, Button(), Variant, Action(), Card() (+95 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 9 - "testing.T"
Cohesion: 0.05
Nodes (70): Next up, Verbs(), TestHelpVerbs(), TestCanCharacter(), nullID(), TestBuilderInventory(), TestCompendiumWeapon(), TestItemLookupByName() (+62 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (25): ClearCookie(), SetCookie(), writeCookie(), FromContext(), ftsQuery(), Search(), Campaign, compendiumBack() (+17 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (62): Amount, CondOpKind, Effects, Env, condCommand(), conditionSlug(), Damage(), findItem() (+54 more)

### Community 13 - ".Entry"
Cohesion: 0.09
Nodes (29): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, ValidKind(), HomebrewInput, HomebrewInputFor() (+21 more)

### Community 14 - "rowFigure"
Cohesion: 0.24
Nodes (7): constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), crText(), rowFigure(), textOnly(), TestRowFigure()

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.07
Nodes (43): Note, GetBlockRow, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanDeleteCharacter() (+35 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "Page"
Cohesion: 0.05
Nodes (91): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Description() (+83 more)

### Community 21 - ".levelUp"
Cohesion: 0.10
Nodes (26): ClassFeature, ClassData, ClassFeature, ClassInfo, LevelUpPlan, classFeatures(), classInfo(), errorField() (+18 more)

### Community 22 - "wikilink.js"
Cohesion: 0.07
Nodes (35): editors, extensions, mount(), parts(), setValue(), sync(), toDoc(), toText() (+27 more)

### Community 23 - "EncounterView"
Cohesion: 0.19
Nodes (13): addCombatantForm(), addCombatants(), Encounter(), EncounterView, hpVisibleSwitch(), inFightSuffix(), nameKey(), rollMissingButton() (+5 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "wikiText"
Cohesion: 0.27
Nodes (8): wikiLink(), wikiText(), ImageSrc(), ParseWiki(), TestParseWiki(), WikiLinks(), WikiLink, WikiPart

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "ParseSheet"
Cohesion: 0.15
Nodes (26): TestUpgradeSheet(), ParseSheet(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpInvocations(), TestLevelUpLineage(), TestRevertLevelUp() (+18 more)

### Community 28 - "App"
Cohesion: 0.07
Nodes (12): Character, Entity, CanCharacter(), entityTopics(), linkEntityLook(), linkCharacterLook(), lookFree(), useValidateTag() (+4 more)

### Community 29 - "monsterBody"
Cohesion: 0.14
Nodes (27): SpellData, classBody(), commaPrefixed(), entryBody(), EntryDetails(), entryHead(), EntryPage(), EntryPreview() (+19 more)

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
Cohesion: 0.03
Nodes (44): ListHomebrewByOwnerRow, ListHomebrewRow, Queries, GetImageRow, Attachment, Condition, CreateCombatantParams, applyCondOps() (+36 more)

### Community 34 - "State"
Cohesion: 0.13
Nodes (10): Handler, patchMessage, viewPatch, Diff(), State, manyKeys(), TestDiff(), deviceLabel() (+2 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "combobox_templ.go"
Cohesion: 0.07
Nodes (51): Align, ChipProps, ChipsInputProps, ChipsProps, ContentProps, ctxKey, ctxState, EmptyProps (+43 more)

### Community 38 - "golang.org/x/net/html.Node"
Cohesion: 0.21
Nodes (19): ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr(), byID() (+11 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - ".issueTokens"
Cohesion: 0.18
Nodes (8): tokens, isCleared(), serve(), TestMiddlewareIdentity(), TestRequireAdmin(), TestTokens(), hashToken(), newToken()

### Community 41 - "NameDesc"
Cohesion: 0.21
Nodes (15): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, catchUpResistances(), TestCatchUpResistances() (+7 more)

### Community 42 - "DecodeData"
Cohesion: 0.16
Nodes (6): DecodeData(), CompendiumEntry, compendiumWeapon(), CreateHomebrewParams, GetCompendiumEntryByNameParams, GetCompendiumEntryBySlugParams

### Community 43 - "export/markdown.go"
Cohesion: 0.12
Nodes (20): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+12 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "Encounter"
Cohesion: 0.09
Nodes (13): Combatant, Encounter, ListCombatantsRow, Character, initiativeMod(), outOfFight(), pcSheet(), rollFor() (+5 more)

### Community 46 - "Client"
Cohesion: 0.24
Nodes (3): message, Client, topicIDs()

### Community 47 - ".Plan"
Cohesion: 0.12
Nodes (12): ClassLevel, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), TestParseEffects(), applyASI() (+4 more)

### Community 48 - "tabs_templ.go"
Cohesion: 0.06
Nodes (67): CountHabitatsRow, SearchCompendiumRow, Badge(), Content(), contentID(), IDFromContext(), initialValue(), List() (+59 more)

### Community 49 - "Topic"
Cohesion: 0.15
Nodes (7): event, fakeHandler, stateHandler, TopicKind, Topic, parseLiveKey(), liveKey

### Community 50 - "New"
Cohesion: 0.44
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

### Community 56 - "Campaign"
Cohesion: 0.32
Nodes (5): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

### Community 57 - "golangci-lint Config"
Cohesion: 0.12
Nodes (16): generated, paths, rules, formatters, enable, exclusions, settings, local-prefixes (+8 more)

### Community 58 - "Slugify"
Cohesion: 0.17
Nodes (16): Cantrips(), TextData, Slugify(), Spell, textMarkdown(), addProfs(), featList(), freeCast() (+8 more)

### Community 59 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs Tool"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "select_templ.go"
Cohesion: 0.10
Nodes (28): RadioGroup(), Content(), Group(), initialValue(), Item(), itemLabel(), Label(), Select() (+20 more)

### Community 63 - "Pending"
Cohesion: 0.22
Nodes (5): PendingState, randomID(), Pending, PendingKind, viewerCtx

### Community 64 - "sqlc Config"
Cohesion: 0.18
Nodes (10): go, emit_empty_slices, out, package, sql, engine, gen, queries (+2 more)

### Community 66 - "Browser check"
Cohesion: 0.40
Nodes (4): Browser check, Roles, The test instance, What to check

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "github.com/a-h/templ.Attributes"
Cohesion: 0.03
Nodes (109): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Size, ActionProps (+101 more)

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
Nodes (59): ImageCategory, CloseFor(), Field(), Group(), Label(), Input(), Textarea(), armorForm() (+51 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ImportView"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

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
Cohesion: 0.13
Nodes (19): block, CoreTraits(), TestCoreTraits(), partsText(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+11 more)

### Community 86 - "levelUpForm"
Cohesion: 0.17
Nodes (24): Label(), spellChecklist(), checkField(), abilityForm(), backgroundAbilities(), choiceFieldset(), choiceName(), choiceOptions() (+16 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ImageResults"
Cohesion: 0.08
Nodes (41): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent(), Props (+33 more)

### Community 100 - "encoding/json.RawMessage"
Cohesion: 0.18
Nodes (17): ftEntry, ftMonster, ftSpellcasting, abilityMod(), ftEntries(), ftEntry1(), ftList(), ftSpells() (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 103 - "Site"
Cohesion: 0.13
Nodes (14): Site, dataFor(), Kind, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey() (+6 more)

### Community 104 - "importSource"
Cohesion: 0.43
Nodes (6): Source, importSource(), readSource(), storeSource(), upsert(), parseOpen5e()

### Community 105 - "Props"
Cohesion: 0.70
Nodes (3): Props, Variant, badgeClasses()

### Community 106 - "LevelUpView"
Cohesion: 0.12
Nodes (14): Creation and level-up pointers, Handoff, Lessons, Test data (data-test, all disposable), Test instance, LevelUpView, Option, SpellNeed (+6 more)

### Community 107 - "Entry"
Cohesion: 0.15
Nodes (18): Entry, fixture, fixtures, normalizer, TestUnflattenTable(), actionRank(), byPK(), children() (+10 more)

### Community 108 - "Import"
Cohesion: 0.09
Nodes (37): fixture, main(), run(), migration, newFixture(), TestOnlyHashIsStored(), Manager, NewManager() (+29 more)

### Community 110 - ".wizardCommand"
Cohesion: 0.11
Nodes (22): EquipmentItem, EquipmentOptions(), EquipmentOption, SkillChoice(), SplitList(), Squash(), TestEquipmentOptions(), TestSkillChoice() (+14 more)

### Community 113 - "Feature"
Cohesion: 0.22
Nodes (8): FeatureEffects, UsesLabel(), CustomMod, Feature, Uses, usesValue(), addFeat(), applyBackground()

### Community 114 - "radiogroup_templ.go"
Cohesion: 0.38
Nodes (6): initialValue(), state(), ctxKey, ctxState, ItemProps, Props

### Community 115 - "Weapon"
Cohesion: 0.13
Nodes (10): Part, Skill, TestModifier(), skillByKey(), TestDualWield(), damageLabel(), Weapon, Modifier() (+2 more)

### Community 118 - "build.sh"
Cohesion: 0.70
Nodes (4): build(), CGO_ENABLED, build.sh script, step()

### Community 119 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 123 - "Sheet"
Cohesion: 0.08
Nodes (17): DeathSaves, Completions(), TestCustomMods(), Sum(), TestGearCommands(), Coins, Item, Sheet (+9 more)

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "Parse"
Cohesion: 0.10
Nodes (23): Defense, Parse(), tokenize(), fighter(), TestCommands(), TestCompletions(), TestDefenses(), TestLanguages() (+15 more)

### Community 126 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 132 - "switch_templ.go"
Cohesion: 0.83
Nodes (3): initialChecked(), Props, Size

### Community 133 - "ArmorData"
Cohesion: 0.40
Nodes (5): ArmorData, o5eArmor, ItemData, armorData(), armorText()

### Community 134 - "Import"
Cohesion: 0.15
Nodes (18): copySource(), importOnce(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding(), TestNormalizedShapes(), TestRealSource(), TestSearch() (+10 more)

### Community 136 - "SearchReadableBlocksRow"
Cohesion: 0.50
Nodes (3): SearchReadableBlocksRow, SearchReadableBlocksParams, NoteListView

## Knowledge Gaps
- **14 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `gm-session-manager`, `testEnv` (+9 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 333 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **24 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `FromContext`, `context.Context`, `.do`, `.mountAssets`, `github.com/gin-gonic/gin.Context`, `perm.go`, `Page`, `.levelUp`, `database/sql.NullInt64`, `State`, `DecodeData`, `Encounter`, `Topic`, `Campaign`, `Slugify`, `.deleteUnused`, `Hub`, `LevelUpView`, `Import`, `.wizardCommand`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _14 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.09052903996416663 - nodes in this community are weakly interconnected._
- **Why does `Sheet` connect `Sheet` to `compendium_templ.go`, `github.com/a-h/templ.Component`, `SheetView`, `NameDesc`, `LevelUpView`, `Command`, `Encounter`, `.wizardCommand`, `.Plan`, `Feature`, `Weapon`, `.levelUp`, `Slugify`, `ParseSheet`, `App`, `Parse`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._
- **Should `FromContext` be split into smaller, more focused modules?**
  _Cohesion score 0.06219512195121951 - nodes in this community are weakly interconnected._
- **Should `context.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.011250654107796965 - nodes in this community are weakly interconnected._
- **Should `createProvider` be split into smaller, more focused modules?**
  _Cohesion score 0.07683000604960677 - nodes in this community are weakly interconnected._