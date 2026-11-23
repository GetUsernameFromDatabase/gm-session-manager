# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 372 files · ~640,863 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4676 nodes · 24155 edges · 152 communities (122 shown, 30 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2414 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `8b1d00ec`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Slugify
- context.Context
- Condition
- createProvider
- tabs_templ.go
- SheetView
- combobox.js
- Queries
- Path
- encoding/json.RawMessage
- select.js
- Command
- ItemData
- Device
- floating_ui_dom.js
- Viewer
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- ParseEffects
- WizardView
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
- viewerCtx
- command.js
- dialog.js
- Session
- ParseWikidot
- tooltip.js
- CompendiumKind
- NameDesc
- config.go
- Session
- composite.js
- buildReview
- Feature
- names
- nullID
- .fail
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- ftItemData
- shadcn-templ Config
- sourcesEnv
- golangci-lint Config
- Import
- Sheet
- radiogroup.js
- vendorjs/main.go
- Input
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- popover_templ.go
- use_transition_status.js
- playwright
- .q
- EquipmentOption
- checkbox.js
- ParseAidedd
- switch.js
- .sourceIssues
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- .editEntity
- useTypeahead
- DownloadFiveTools
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- database/sql.NullString
- mounted
- .sendFeedback
- Player
- profFocus
- CompendiumPath
- time.Time
- ftBook
- NewClient
- BlockEditor
- entityForm
- Input
- csvSource
- ListSharedWithRow
- export/markdown.go
- build.sh
- SearchReadableBlocksRow
- looseString
- errBodyWriter
- State
- CampaignSources
- Before/after pairs
- Hub
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- Topic
- ASIError
- Book
- .Entry
- .attachFile
- sourceUsedError
- Pending
- ftRec
- TestSourceSwitchKeepsEdits
- Button
- Campaign
- InputForEntry
- .source
- toast_templ.go
- Option

## God Nodes (most connected - your core abstractions)
1. `App` - 374 edges
2. `Path()` - 265 edges
3. `newTestEnv()` - 168 edges
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
- `internal/compendium` --references--> `SpellData`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `DecodeData()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go
- `internal/compendium` --references--> `FiveToolsBooks()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/fivetools_books.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (152 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "Slugify"
Cohesion: 0.09
Nodes (33): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), AbilityKeys(), AddFeat() (+25 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (75): keepAwake(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewLinksRow, ListImageEntitiesRow, CompendiumSource, Note, EntityKind (+67 more)

### Community 3 - "Condition"
Cohesion: 0.03
Nodes (68): Defense, Combatant, Condition, Encounter, CreateCombatantParams, CreateConditionParams, ListCombatantsRow, Defenses (+60 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "tabs_templ.go"
Cohesion: 0.14
Nodes (24): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+16 more)

### Community 6 - "SheetView"
Cohesion: 0.06
Nodes (95): ArmorEffect(), Slot, Signed(), Variant, Action(), Card(), Content(), Description() (+87 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Queries"
Cohesion: 0.06
Nodes (42): featureSwitch, Decode(), ListSubclassesOfRow, Queries, Env, TestSlotMaxima(), ClassInfo, ClassLevel (+34 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (115): Import(), TestUpgradeSheet(), ParseSheet(), TestMarkdown(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight() (+107 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.06
Nodes (47): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+39 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.06
Nodes (57): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+49 more)

### Community 13 - "ItemData"
Cohesion: 0.04
Nodes (65): ImageCategory, Title(), Content(), Group(), initialValue(), Item(), itemLabel(), Label() (+57 more)

### Community 14 - "Device"
Cohesion: 0.40
Nodes (8): deviceRow(), rosterCard(), devicesWhere(), agoTime(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.05
Nodes (53): ListCharacterNamesRow, GetBlockRow, Visibility, VisibleEntity(), VisibleLooks(), BlockVisibilities(), CanAddEffect(), CanDeleteCharacter() (+45 more)

### Community 17 - "use_hover.js"
Cohesion: 0.09
Nodes (41): applySafePolygonPointerEventsMutation(), clearSafePolygonPointerEventsMutation(), createHoverInteraction(), emitClosed(), getDelay(), getRestMs(), hasIntersectingEdge(), isClickLikeOpenEvent() (+33 more)

### Community 18 - "floating_ui_core.js"
Cohesion: 0.13
Nodes (37): arrowMiddleware(), getLogicalSide(), hideMiddleware(), isRtlAt(), roundByDPR(), useAnchorPositioning(), getOffsetData(), middleware() (+29 more)

### Community 19 - "scroll_lock.js"
Cohesion: 0.11
Nodes (15): anchoredPopupScrollLock(), AnimationFrame, getViewportScroller(), hasInsetScrollbars(), isOverflowElement(), isPageScrollLocked(), preventScrollInsetScrollbars(), cleanup() (+7 more)

### Community 20 - "ParseEffects"
Cohesion: 0.11
Nodes (12): LevelUpChoice, armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects(), TestParseProfs(), applyASI() (+4 more)

### Community 21 - "WizardView"
Cohesion: 0.25
Nodes (20): Label(), TestOptionText(), abilityForm(), backgroundAbilities(), stepLinks(), wizardAbilities(), wizardChoiceCard(), wizardChoices() (+12 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (42): ClearCookie(), SetCookie(), writeCookie(), Identity, FromContext(), Campaign, compendiumBack(), compendiumPath() (+34 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (162): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+154 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "cert.go"
Cohesion: 0.11
Nodes (18): main(), run(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate(), LoadOrCreateCert() (+10 more)

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
Cohesion: 0.14
Nodes (4): countingHandler, fakeHandler, message, Client

### Community 34 - "viewerCtx"
Cohesion: 0.23
Nodes (7): notFoundIsNil(), parseLiveKey(), Roster, liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "Session"
Cohesion: 0.10
Nodes (18): fixture, tokens, constantTimeEqual(), credentialsMatch(), newFixture(), serve(), TestCredentialsMatch(), TestMiddlewareIdentity() (+10 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (24): block, ClassFeature, ClassData, withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot() (+16 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.04
Nodes (45): CoreTraits(), dataFor(), DecodeData(), ValidKind(), ListHomebrewByOwnerRow, ListHomebrewRow, UpsertCompendiumEntryParams, CompendiumEntry (+37 more)

### Community 41 - "NameDesc"
Cohesion: 0.11
Nodes (29): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, AddResistance(), CantripFits() (+21 more)

### Community 42 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 43 - "Session"
Cohesion: 0.16
Nodes (6): IsLoopback(), RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.14
Nodes (20): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+12 more)

### Community 46 - "Feature"
Cohesion: 0.08
Nodes (26): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+18 more)

### Community 47 - "names"
Cohesion: 0.33
Nodes (7): SwitchFeature, SwitchPlan, PlanSourceSwitch(), names(), TestPlanSourceSwitch(), TestClassOptionChoices(), SourcePreviewView

### Community 48 - "nullID"
Cohesion: 0.14
Nodes (29): DeathSaves, nullID(), deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead() (+21 more)

### Community 49 - ".fail"
Cohesion: 0.07
Nodes (17): Attachment, Character, CanChangeSubclass(), CanCharacter(), saveScreenshot(), parseCSV(), TestGuessMapping(), linkCharacterLook() (+9 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (116): TestIsLoopback(), TestDecode(), TestCoreTraits(), feature(), TestFiveToolsClasses(), RemoveSourceDir(), TestRemoveSourceDir(), TestPageClientRedirect() (+108 more)

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
Cohesion: 0.11
Nodes (17): ArmorData, ftClassFeature, ftClassRefs, ftItem, o5eArmor, ItemData, ftArmor(), ftCode() (+9 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "sourcesEnv"
Cohesion: 0.28
Nodes (7): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "Import"
Cohesion: 0.07
Nodes (53): sourceFormat, migration, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey() (+45 more)

### Community 59 - "Sheet"
Cohesion: 0.06
Nodes (25): Part, Skill, TestCustomMods(), TestModifier(), TestResistances(), TestExplainSums(), skillByKey(), Sum() (+17 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "Input"
Cohesion: 0.08
Nodes (53): ctxState, Chip(), Chips(), Combobox(), Content(), Input(), Item(), itemLabel() (+45 more)

### Community 63 - "FiveToolsJob"
Cohesion: 0.31
Nodes (3): ImportNotice, FiveToolsJob, fiveToolsJob

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.09
Nodes (40): Props, Variant, Badge(), badgeClasses(), addCombatantForm(), addCombatants(), beingBadge(), combatantChip() (+32 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "popover_templ.go"
Cohesion: 0.15
Nodes (24): Content(), Description(), Header(), Popover(), state(), Title(), Trigger(), triggerID() (+16 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - ".q"
Cohesion: 0.10
Nodes (28): isCleared(), testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+20 more)

### Community 72 - "EquipmentOption"
Cohesion: 0.33
Nodes (6): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), applyEquipment(), EquipmentChoice

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - ".sourceIssues"
Cohesion: 0.43
Nodes (6): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, SourceLabel(), sheetRef

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

### Community 86 - ".editEntity"
Cohesion: 0.07
Nodes (19): Entity, LinkTarget, Fields, ParseFields(), entityTopics(), linkEntityLook(), saveFieldLinks(), lookFree() (+11 more)

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "DownloadFiveTools"
Cohesion: 0.14
Nodes (23): Source, TestStoreSourceRollsBackOnAfterSave(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition(), TestFiveToolsMagicVariants() (+15 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "database/sql.NullString"
Cohesion: 0.07
Nodes (17): GetImageRow, ensureCategory(), imageItem(), imageVersion(), nullString(), BindMonsterImageParams, CompendiumFt, CreateImageParams (+9 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 103 - "Player"
Cohesion: 0.07
Nodes (38): Player(), TestCanRemoveCondition(), BenchmarkCompendiumList(), BenchmarkCompendiumSearch(), BenchmarkEncounterState(), BenchmarkSheetState(), testEnv, Gotchas (+30 more)

### Community 105 - "CompendiumPath"
Cohesion: 0.05
Nodes (59): SpellData, CountCompendiumRow, CountHabitatsRow, SearchCompendiumRow, classBadges(), CompendiumResults(), compendiumRow(), CompendiumRows() (+51 more)

### Community 106 - "time.Time"
Cohesion: 0.05
Nodes (24): ListAttachmentFilesRow, ListImagesRow, Handlers, browserName(), firstLine(), mdLinkText(), inCampaign(), TestUserAgent() (+16 more)

### Community 107 - "ftBook"
Cohesion: 0.20
Nodes (15): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftParts(), ftRange() (+7 more)

### Community 108 - "NewClient"
Cohesion: 0.23
Nodes (18): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+10 more)

### Community 111 - "entityForm"
Cohesion: 0.07
Nodes (42): ListEntityFieldBacklinksRow, ListBacklinksRow, Entity, Look, attrString(), generateSVG(), getIconContent(), Props (+34 more)

### Community 112 - "Input"
Cohesion: 0.06
Nodes (85): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+77 more)

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

### Community 125 - "State"
Cohesion: 0.18
Nodes (7): patchMessage, stateHandler, viewPatch, Diff(), State, manyKeys(), TestDiff()

### Community 127 - "CampaignSources"
Cohesion: 0.32
Nodes (7): campaignSourceRow(), CampaignSources(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), CampaignSource, sourceGroup

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "Hub"
Cohesion: 0.22
Nodes (4): ClientInfo, goneDevice, Hub, isDevice()

### Community 133 - "open5e.go"
Cohesion: 0.14
Nodes (19): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+11 more)

### Community 137 - "Topic"
Cohesion: 0.18
Nodes (6): event, Handler, Topic, TopicKind, heldCollected(), heldTopics

### Community 139 - "Book"
Cohesion: 0.13
Nodes (21): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+13 more)

### Community 140 - ".Entry"
Cohesion: 0.10
Nodes (26): HomebrewOwner, MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInput, HomebrewInputFor(), SaveHomebrew() (+18 more)

### Community 141 - ".attachFile"
Cohesion: 0.24
Nodes (9): testEnv, TestAttachmentAccess(), TestAttachmentFileMissing(), TestAttachmentToImage(), TestCampaignFiles(), TestDeleteRemovesFiles(), TestFilesCleanup(), TestFileUsePrivacy() (+1 more)

### Community 143 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 144 - "ftRec"
Cohesion: 0.27
Nodes (8): ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve(), ftVersions()

### Community 145 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.39
Nodes (6): classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 146 - "Button"
Cohesion: 0.06
Nodes (107): ListEntryCopiesRow, ListMonstersByNameRow, Habitat, UI conventions, Button(), Content(), Description(), Empty() (+99 more)

### Community 147 - "Campaign"
Cohesion: 0.32
Nodes (5): Campaign(), CharacterFileName(), fileSlug(), NewCharacterFile(), writeFile()

### Community 148 - "InputForEntry"
Cohesion: 0.20
Nodes (9): habitatsFrom(), FiveToolsImages(), TestHabitatsFrom(), TestTokensReadAsLabels(), ParseFiveTools(), bindFiveToolsImages(), InputForEntry(), TestParseFiveTools() (+1 more)

### Community 149 - ".source"
Cohesion: 0.27
Nodes (8): testEnv, TestEntryAlsoInAndCrossRefs(), TestFiveToolsGMOnly(), TestFiveToolsImportRunning(), TestFiveToolsRemove(), TestSourceDeleteUsed(), TestSourceMove(), TestSourceOrder()

### Community 150 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 159 - "Option"
Cohesion: 0.09
Nodes (23): dnd5eFeatures(), TestFiveToolsClassLevelUp(), ClassFeature, LevelUpPlan, SpellColumn(), spellBoxes(), spellChecklist(), TestSpellCounts() (+15 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 413 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Slugify`, `context.Context`, `Condition`, `Hub`, `Queries`, `Book`, `Command`, `Pending`, `Viewer`, `Campaign`, `cert.go`, `viewerCtx`, `CompendiumKind`, `NameDesc`, `Session`, `.fail`, `testing.T`, `Sheet`, `FiveToolsJob`, `.q`, `.sourceIssues`, `.editEntity`, `database/sql.NullString`, `CompendiumPath`, `time.Time`, `State`?**
  _High betweenness centrality (0.069) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.13903885576577 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `Condition`, `SheetView`, `.attachFile`, `Device`, `ItemData`, `Viewer`, `TestSourceSwitchKeepsEdits`, `Button`, `Campaign`, `.source`, `WizardView`, `App`, `github.com/a-h/templ.Attributes`, `CompendiumKind`, `nullID`, `.fail`, `testing.T`, `sourcesEnv`, `Import`, `Input`, `CombatantView`, `popover_templ.go`, `.q`, `ParseWiki`, `.editEntity`, `database/sql.NullString`, `.sendFeedback`, `Player`, `CompendiumPath`, `time.Time`, `entityForm`, `Input`, `export/markdown.go`, `CampaignSources`?**
  _High betweenness centrality (0.042) - this node is a cross-community bridge._
- **Should `Slugify` be split into smaller, more focused modules?**
  _Cohesion score 0.08705882352941176 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Command`, `wikilink.js`, `Player`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._