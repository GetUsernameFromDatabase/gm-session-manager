# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 375 files · ~648,277 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 57 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4734 nodes · 24494 edges · 158 communities (124 shown, 34 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2458 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `c0f9d248`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Layout
- context.Context
- vendorjs/main.go
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- database/sql.NullInt64
- Path
- encoding/json.RawMessage
- select.js
- Command
- Input
- tooltip_templ.go
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- IfElse
- wikilink.js
- names
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- cert.go
- Site
- github.com/a-h/templ.Component
- app.js
- tabbable.js
- useDismiss
- NewClient
- Habitat
- command.js
- dialog.js
- Store
- ParseWikidot
- tooltip.js
- CompendiumKind
- Slugify
- viewerCtx
- .Router
- composite.js
- buildReview
- CustomMod
- Queries
- entityForm
- NameDesc
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- .editEntity
- shadcn-templ Config
- Session
- golangci-lint Config
- ItemData
- Sheet
- radiogroup.js
- Player
- DownloadFiveTools
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- State
- use_transition_status.js
- playwright
- Attachment
- Campaign
- checkbox.js
- ParseAidedd
- switch.js
- sourcesEnv
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- .q
- useTypeahead
- ftItemData
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- deathFix
- mounted
- .sendFeedback
- nullID
- profFocus
- CompendiumPath
- redirect.go
- ftBook
- BlockEditor
- .LevelUp
- .compendiumDir
- capReader
- .features
- time.Time
- build.sh
- TestSourceSwitchKeepsEdits
- looseString
- errBodyWriter
- ListSharedWithRow
- export/markdown.go
- Before/after pairs
- ParseArmorTraining
- Hub
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- config.go
- ASIError
- Book
- .Entry
- Topic
- sourceUsedError
- Pending
- rosterCard
- toast_templ.go
- internal/compendium
- subclassPanels
- .sourceIssues
- parseCommand
- ClassData
- EquipmentOption
- SearchReadableBlocksRow
- .SpellAbilityOf
- CharacterFile
- spellCounts

## God Nodes (most connected - your core abstractions)
1. `App` - 374 edges
2. `Path()` - 270 edges
3. `newTestEnv()` - 173 edges
4. `Sheet` - 133 edges
5. `Button()` - 124 edges
6. `CN()` - 114 edges
7. `SheetView` - 78 edges
8. `IfElse()` - 65 edges
9. `Command` - 60 edges
10. `Queries` - 56 edges

## Surprising Connections (you probably didn't know these)
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
- `TestScopedEntryInEveryCompendiumQuery()` --calls--> `check()`  [INFERRED]
  internal/perm/sql_agreement_test.go → cmd/vendorjs/main.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (158 total, 34 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "Layout"
Cohesion: 0.08
Nodes (46): UI conventions, Scripts(), Separator(), Spinner(), campaignLink(), PendingRequests(), PendingGroup, CharacterPage() (+38 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (83): keepAwake(), ListImageCharactersRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Note, CreateCombatantParams, EntityKind (+75 more)

### Community 3 - "vendorjs/main.go"
Cohesion: 0.20
Nodes (18): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+10 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.11
Nodes (38): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+30 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (106): ArmorEffect(), Coins, Slot, Signed(), Button(), Variant, Action(), Card() (+98 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (74): ListHomebrewByOwnerRow, ListHomebrewRow, Combatant, Condition, Encounter, CreateConditionParams, ListCombatantsRow, Character (+66 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (126): TestRealSource(), Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount() (+118 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (50): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+42 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (53): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+45 more)

### Community 13 - "Input"
Cohesion: 0.07
Nodes (57): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+49 more)

### Community 14 - "tooltip_templ.go"
Cohesion: 0.09
Nodes (34): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks(), BuilderPick (+26 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (56): ListCharacterNamesRow, GetBlockRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanChangeSubclass() (+48 more)

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
Cohesion: 0.04
Nodes (26): SetCookie(), Identity, FromContext(), pathCampaign(), collectHeld(), formIntQuery(), isHTMX(), redirect() (+18 more)

### Community 21 - "IfElse"
Cohesion: 0.07
Nodes (63): Props, Checkbox(), initialChecked(), ItemData, Label(), initialValue(), Item(), RadioGroup() (+55 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (144): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+136 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.14
Nodes (11): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+3 more)

### Community 28 - "Site"
Cohesion: 0.15
Nodes (12): Site, fiveToolsPage(), FetchPage(), Page, pageSource(), PageSourceKey(), ParsePageURL(), SavePage() (+4 more)

### Community 29 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (112): ListMonstersByNameRow, Field, Combobox(), Content(), Empty(), Group(), Input(), Item() (+104 more)

### Community 30 - "app.js"
Cohesion: 0.09
Nodes (20): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+12 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "NewClient"
Cohesion: 0.19
Nodes (20): isDevice(), New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+12 more)

### Community 34 - "Habitat"
Cohesion: 0.09
Nodes (44): CountHabitatsRow, ListEntryCopiesRow, Habitat, classBadges(), TestRowFigure(), commaPrefixed(), compendiumFindURL(), crText() (+36 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Store"
Cohesion: 0.05
Nodes (64): main(), run(), Source, sourceFormat, migration, NewManager(), copySource(), importOnce() (+56 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.12
Nodes (25): block, CoreTraits(), TestCoreTraits(), TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt(), classTable(), featuresUnder() (+17 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (48): dataFor(), DecodeData(), ValidKind(), ftsQuery(), Search(), CountCompendiumRow, SearchCompendiumRow, UpsertCompendiumEntryParams (+40 more)

### Community 41 - "Slugify"
Cohesion: 0.06
Nodes (46): Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), AbilityKeys(), AddFeat(), AddResistance() (+38 more)

### Community 42 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 43 - ".Router"
Cohesion: 0.15
Nodes (8): IsLoopback(), ClearCookie(), writeCookie(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.10
Nodes (26): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+18 more)

### Community 46 - "CustomMod"
Cohesion: 0.08
Nodes (21): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+13 more)

### Community 47 - "Queries"
Cohesion: 0.09
Nodes (33): featureSwitch, Decode(), Queries, Env, ClassInfo, ClassLevel, WornArmor, featureNames() (+25 more)

### Community 48 - "entityForm"
Cohesion: 0.09
Nodes (36): ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent() (+28 more)

### Community 49 - "NameDesc"
Cohesion: 0.12
Nodes (29): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+21 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (113): Defense, newFixture(), serve(), TestIsLoopback(), TestMiddlewareIdentity(), TestOnlyHashIsStored(), TestRequireAdmin(), TestResolveDeletedCharacter() (+105 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - ".editEntity"
Cohesion: 0.10
Nodes (11): Entity, Fields, ParseFields(), entityTopics(), linkEntityLook(), saveFieldLinks(), formInt(), SetEntityCurrentImageParams (+3 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "Session"
Cohesion: 0.15
Nodes (12): fixture, tokens, constantTimeEqual(), credentialsMatch(), TestCredentialsMatch(), TestTokens(), Manager, hashToken() (+4 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "ItemData"
Cohesion: 0.10
Nodes (25): firstValue(), modTargetLabel(), usesItems(), weaponAbilities(), EffectFieldsForm(), effectModFields(), EffectModRow(), effectHeldItems() (+17 more)

### Community 59 - "Sheet"
Cohesion: 0.06
Nodes (21): Part, Skill, findItem(), parseItem(), TestCustomMods(), TestModifier(), TestExplainSums(), skillByKey() (+13 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.07
Nodes (40): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas (+32 more)

### Community 62 - "DownloadFiveTools"
Cohesion: 0.11
Nodes (26): EquipmentOptions(), TestEquipmentOptions(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition() (+18 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.10
Nodes (40): Header(), Title(), effectBits(), beingBadge(), combatantChip(), combatantName(), combatantRow(), combatCondition() (+32 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "State"
Cohesion: 0.11
Nodes (10): countingHandler, fakeHandler, Handler, message, patchMessage, viewPatch, Client, Diff() (+2 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "Attachment"
Cohesion: 0.14
Nodes (6): Attachment, firstLine(), mdLinkText(), saveScreenshot(), saveSniffed(), CreateAttachmentParams

### Community 72 - "Campaign"
Cohesion: 0.16
Nodes (7): ListAttachmentFilesRow, Campaign, fresh(), inCampaign(), Roster, CreateCampaignParams, CampaignSwitch

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "sourcesEnv"
Cohesion: 0.24
Nodes (8): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCampaignSourceVisibility(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

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

### Community 86 - ".q"
Cohesion: 0.10
Nodes (28): isCleared(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+20 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ftItemData"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 103 - "nullID"
Cohesion: 0.11
Nodes (37): nullID(), TestCombatantNamesFollowSource(), TestEncounterTransitions(), TestMonsterPortrait(), acBonusForm(), effectFix, newEffectFix(), notes() (+29 more)

### Community 105 - "CompendiumPath"
Cohesion: 0.08
Nodes (55): Props, Variant, MediaProps, MediaVariant, CountSharedBySourceRow, Badge(), badgeClasses(), Content() (+47 more)

### Community 106 - "redirect.go"
Cohesion: 0.24
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 107 - "ftBook"
Cohesion: 0.13
Nodes (23): ftBook, ftDistance, ftRec, each(), ftAmount(), ftCastingTime(), ftCopyTargets(), ftDecode() (+15 more)

### Community 111 - ".LevelUp"
Cohesion: 0.09
Nodes (16): LevelUpChoice, TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+8 more)

### Community 112 - ".compendiumDir"
Cohesion: 0.27
Nodes (4): ImportSummary(), fiveToolsKey(), sourceInUseMsg(), sourceUsers()

### Community 116 - ".features"
Cohesion: 0.24
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 117 - "time.Time"
Cohesion: 0.04
Nodes (35): GetImageRow, ListImagesRow, Handlers, browserName(), TestUserAgent(), userAgentSummary(), ensureCategory(), imageItem() (+27 more)

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 127 - "export/markdown.go"
Cohesion: 0.13
Nodes (19): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+11 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 131 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 132 - "Hub"
Cohesion: 0.19
Nodes (4): ClientInfo, goneDevice, Hub, TestCloseSessionAndCharacter()

### Community 133 - "open5e.go"
Cohesion: 0.12
Nodes (22): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+14 more)

### Community 137 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 139 - "Book"
Cohesion: 0.12
Nodes (24): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+16 more)

### Community 140 - ".Entry"
Cohesion: 0.08
Nodes (30): HomebrewOwner, MonsterAction, addEntry(), setStat(), TestHomebrew(), MonsterData, TestTokensReadAsLabels(), HomebrewInput (+22 more)

### Community 141 - "Topic"
Cohesion: 0.24
Nodes (6): event, stateHandler, Topic, TopicKind, heldCollected(), heldTopics

### Community 143 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 144 - "rosterCard"
Cohesion: 0.40
Nodes (8): deviceRow(), rosterCard(), devicesWhere(), agoTime(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 145 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 146 - "internal/compendium"
Cohesion: 0.29
Nodes (4): internal/compendium, SpellData, TestDurationRounds(), durationRounds()

### Community 147 - "subclassPanels"
Cohesion: 0.29
Nodes (6): ListSubclassesOfRow, TestSubclassPanels(), subclassIndex(), subclassOptions(), subclassPanels(), ListSubclassesOfParams

### Community 148 - ".sourceIssues"
Cohesion: 0.43
Nodes (6): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

### Community 151 - "parseCommand"
Cohesion: 0.33
Nodes (6): featureCommand(), coinCommand(), commandText(), hpCommand(), parseCommand(), damageArgs()

### Community 152 - "ClassData"
Cohesion: 0.50
Nodes (5): ClassFeature, ClassData, dnd5eFeatures(), ClassFeature, classFeatures()

### Community 153 - "EquipmentOption"
Cohesion: 0.40
Nodes (5): EquipmentItem, EquipmentOption, applyEquipment(), EquipmentChoice, gear

### Community 156 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 157 - "spellCounts"
Cohesion: 0.50
Nodes (4): SpellColumn(), TestSpellCounts(), sortedColumns(), spellCounts()

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 418 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **34 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Layout`, `tabs_templ.go`, `SheetView`, `database/sql.NullInt64`, `Input`, `tooltip_templ.go`, `Viewer`, `rosterCard`, `App`, `IfElse`, `Site`, `github.com/a-h/templ.Component`, `Habitat`, `Store`, `CompendiumKind`, `Slugify`, `entityForm`, `testing.T`, `.editEntity`, `ItemData`, `Player`, `CombatantView`, `Attachment`, `Campaign`, `sourcesEnv`, `ParseWiki`, `.q`, `deathFix`, `.sendFeedback`, `nullID`, `CompendiumPath`, `time.Time`, `TestSourceSwitchKeepsEdits`, `export/markdown.go`?**
  _High betweenness centrality (0.049) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.14059114825131538 - nodes in this community are weakly interconnected._
- **Why does `App` connect `App` to `compendium_templ.go`, `context.Context`, `Hub`, `database/sql.NullInt64`, `Book`, `Command`, `Pending`, `Viewer`, `internal/compendium`, `.sourceIssues`, `Site`, `Store`, `CompendiumKind`, `Slugify`, `viewerCtx`, `.Router`, `Queries`, `testing.T`, `.editEntity`, `Sheet`, `FiveToolsJob`, `State`, `Attachment`, `Campaign`, `.q`, `.compendiumDir`, `time.Time`?**
  _High betweenness centrality (0.037) - this node is a cross-community bridge._
- **Should `Layout` be split into smaller, more focused modules?**
  _Cohesion score 0.08244680851063829 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `Player`, `wikilink.js`?**
  _High betweenness centrality (0.025) - this node is a cross-community bridge._