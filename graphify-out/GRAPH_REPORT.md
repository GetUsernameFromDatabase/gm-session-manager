# Graph Report - gm-session-manager  (2026-11-27)

## Corpus Check
- 400 files · ~705,571 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 60 file(s) not represented in the graph (top: .templ 50, .css 4, (none) 3)

## Summary
- 5201 nodes · 27643 edges · 178 communities (123 shown, 55 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2748 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `8f604f52`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- nullID
- context.Context
- takehome_test.go
- createProvider
- export/markdown.go
- Button
- combobox.js
- ParseSheet
- .sourceIssues
- encoding/json.RawMessage
- select.js
- time.Time
- empty_templ.go
- Session
- floating_ui_dom.js
- perm.go
- use_hover.js
- floating_ui_core.js
- scroll_lock.js
- App
- formSelect
- wikilink.js
- Book
- popover.js
- github.com/a-h/templ.Attributes
- Project Rules (CLAUDE.md, README)
- Weapon
- CombatantView
- Attachment
- app.js
- tabbable.js
- useDismiss
- ParseArmorTraining
- NameDesc
- command.js
- dialog.js
- GetImageRow
- ParseWikidot
- tooltip.js
- Path
- tabs_templ.go
- WizardView
- Slugify
- composite.js
- CampaignSources
- CustomMod
- vendorjs/main.go
- SourcePriority
- Path
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- starterEnv
- shadcn-templ Config
- ClassInfo
- golangci-lint Config
- storeSource
- redirect.go
- radiogroup.js
- Player
- Sheet
- Campaign
- sqlc Config
- viewerCtx
- handoff
- Claude Settings
- Command
- use_transition_status.js
- playwright
- ftPrerequisite
- config.go
- checkbox.js
- ParseAidedd
- switch.js
- Store
- tabs.js
- portal.js
- useClick
- ParseWiki
- useFocus
- mark_others.js
- notes.js
- Lint and Release Rules
- Run Script
- mirror
- useTypeahead
- .q
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- ftRec
- mounted
- newPlayerEnv
- newEffectFix
- profFocus
- cert.go
- plainTags
- ftBook
- .sendFeedback
- BlockEditor
- Values Go and JS share
- ftItemData
- .Router
- buildReview
- TestSourceSwitchKeepsEdits
- build.sh
- Pending
- looseString
- errBodyWriter
- State
- EquipmentOption
- Before/after pairs
- RosterCharacter
- jsconfig.json
- open5e.go
- testing.M
- pause_windows.go
- github.com/a-h/templ.Component
- formErrors
- ASIError
- DownloadFiveTools
- .Entry
- sourceUsedError
- planSwitch
- capReader
- Client
- shared.d.ts
- CMD_FIELD
- CONSOLE_LEVELS
- CONSOLE_MSG
- CONSOLE_PAGE
- CONSOLE_SIZE
- KEY
- MAX_STEP_BYTES
- MAX_STEPS
- MSG
- ROUTES
- SAFE_SCHEMES
- WIKILINK
- countWriter
- IMPORT_FILTER_ALL
- IMPORT_FILTER_CATEGORY
- IMPORT_FILTER_NONE
- IMPORT_FILTER_STATUS
- IMPORT_ROWS_FIELD
- IMPORT_STATUS_ANY
- IMPORT_STATUS_BROKEN
- IMPORT_STATUS_TICKED
- IMPORT_STATUS_UNTICKED
- CompendiumKind
- CompendiumView
- Browser check
- player-app.sh
- IMPORT_FILTER_CAT_PREFIX

## God Nodes (most connected - your core abstractions)
1. `App` - 424 edges
2. `Path()` - 303 edges
3. `newTestEnv()` - 186 edges
4. `Sheet` - 134 edges
5. `Button()` - 133 edges
6. `CN()` - 114 edges
7. `SheetView` - 80 edges
8. `Queries` - 67 edges
9. `IfElse()` - 67 edges
10. `Command` - 60 edges

## Surprising Connections (you probably didn't know these)
- `Notes` --references--> `linkClick()`  [INFERRED]
  internal/web/CLAUDE.md → internal/web/assets/js/wikilink.js
- `graphify skill trigger` --semantically_similar_to--> `graphify Usage Rules`  [INFERRED] [semantically similar]
  .claude/CLAUDE.md → CLAUDE.md
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

## Communities (178 total, 55 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.15
Nodes (57): contextKey, classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass, o5eDescribes (+49 more)

### Community 1 - "nullID"
Cohesion: 0.05
Nodes (49): Combatant, Condition, CreateConditionParams, ListCombatantsRow, IsConditionSlug(), EffectEnd, combatantRow(), passTurn() (+41 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (153): keepAwake(), ListImageCharactersRow, CountSharedBySourceRow, ListHomebrewByOwnerRow, ListHomebrewLinksRow, ListHomebrewRow, ListSharedWithRow, ListImageEntitiesRow (+145 more)

### Community 3 - "takehome_test.go"
Cohesion: 0.12
Nodes (22): TakeHomeFile, readLimited(), ReadTakeHome(), rewriteZip(), sampleBringBack(), sampleTakeHome(), TestBringBackRoundTrip(), TestReadTakeHomeKinds() (+14 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "export/markdown.go"
Cohesion: 0.13
Nodes (17): Links, pmMark, pmNode, writer, TestMarkdown(), applyMarks(), escape(), Markdown() (+9 more)

### Community 6 - "Button"
Cohesion: 0.05
Nodes (110): Slot, Signed(), Variant, Action(), Card(), Content(), Description(), Header() (+102 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "ParseSheet"
Cohesion: 0.08
Nodes (51): TestRealSource(), Import(), ParseSheet(), joinHabitats(), TestCompendiumHabitatFilter(), TestEntryBackKeepsFilters(), TestHabitatCopies(), TestHabitatsMatchSchema() (+43 more)

### Community 9 - ".sourceIssues"
Cohesion: 0.43
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.07
Nodes (37): ftClass, ftCopyRef, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace (+29 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "time.Time"
Cohesion: 0.05
Nodes (46): HomeManifest, Kind, TakeHomeBaseCharacter, TakeHomeBaseLook, TakeHomeBlock, TakeHomeBlockFile, TakeHomeCampaign, TakeHomeEntry (+38 more)

### Community 13 - "empty_templ.go"
Cohesion: 0.08
Nodes (47): ContentProps, DescriptionProps, HeaderProps, MediaProps, MediaVariant, Props, TitleProps, Entity (+39 more)

### Community 14 - "Session"
Cohesion: 0.08
Nodes (26): fixture, tokens, constantTimeEqual(), credentialsMatch(), IsLoopback(), isCleared(), newFixture(), serve() (+18 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "perm.go"
Cohesion: 0.05
Nodes (63): ListCharacterNamesRow, GetBlockRow, Visibility, internal/perm, Fields, ParseFields(), VisibleEntity(), VisibleLooks() (+55 more)

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
Cohesion: 0.03
Nodes (38): Identity, FromContext(), Character, Entity, downloadNames(), entityTopics(), browserName(), collectHeld() (+30 more)

### Community 21 - "formSelect"
Cohesion: 0.06
Nodes (47): Defenses, damageChoice(), builderFields(), itemDetails(), subjectField(), firstValue(), missingDamageItems(), modTargetLabel() (+39 more)

### Community 22 - "wikilink.js"
Cohesion: 0.17
Nodes (18): nodeActions(), closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KeepImages, LAYOUTS (+10 more)

### Community 23 - "Book"
Cohesion: 0.11
Nodes (25): rawBook, bookGroup(), fetchBooks(), FiveTools2024Set(), FiveToolsBooks(), Book, ImportedFiveTools(), normalizeBooks() (+17 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (185): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+177 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Weapon"
Cohesion: 0.06
Nodes (25): Part, Skill, TestCustomMods(), TestModifier(), ArmorEffect(), skillByKey(), Sum(), geared() (+17 more)

### Community 28 - "CombatantView"
Cohesion: 0.06
Nodes (71): Description(), Title(), Input(), Addon(), Button(), Input(), InputGroup(), Header() (+63 more)

### Community 29 - "Attachment"
Cohesion: 0.10
Nodes (11): Attachment, firstLine(), folderName(), mdLinkText(), saveScreenshot(), slugify(), TestFolderName(), parseCSV() (+3 more)

### Community 30 - "app.js"
Cohesion: 0.08
Nodes (21): checkLink(), checkLinks(), clearDisplay(), commandChoices(), commandHint(), commandState(), createItem(), filterImport() (+13 more)

### Community 31 - "tabbable.js"
Cohesion: 0.14
Nodes (28): activeElement(), appendCandidates(), appendMatchingElements(), contains(), disableFocusInside(), enableFocusInside(), focusable(), getComposedChildren() (+20 more)

### Community 32 - "useDismiss"
Cohesion: 0.19
Nodes (25): contains(), getTarget(), isLastTraversableNode(), normalizeBubbles(), useDismiss(), addTargetEventListenerOnce(), close(), closeOnEscapeKeyDown() (+17 more)

### Community 33 - "ParseArmorTraining"
Cohesion: 0.16
Nodes (9): armorGrants(), armorKinds(), armorOnly(), armorSourceKey(), ArmorTrainingText(), ParseArmorTraining(), TestArmorTrainingTextReadsBack(), TestParseArmorTraining() (+1 more)

### Community 34 - "NameDesc"
Cohesion: 0.15
Nodes (22): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+14 more)

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "GetImageRow"
Cohesion: 0.11
Nodes (16): GetImageRow, LinkTarget, saveFieldLinks(), linkTarget(), ensureCategory(), imageItem(), imageSource(), imageVersion() (+8 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.13
Nodes (23): block, TestFiveToolsCoreTraitsShapes(), withCoreTraits(), cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "Path"
Cohesion: 0.04
Nodes (147): ListEntryCopiesRow, ListMonstersByNameRow, ListEntityFieldBacklinksRow, ImageCategory, ListBacklinksRow, Habitat, Field, Path() (+139 more)

### Community 41 - "tabs_templ.go"
Cohesion: 0.09
Nodes (40): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+32 more)

### Community 42 - "WizardView"
Cohesion: 0.11
Nodes (33): initialValue(), Item(), RadioGroup(), state(), feedbackKind(), spellChecklist(), splitCantrips(), markdown() (+25 more)

### Community 43 - "Slugify"
Cohesion: 0.07
Nodes (40): internal/compendium, Cantrips(), SkillChoice(), TestSkillChoice(), TextData, Slugify(), fiveToolsPage(), AbilityKeys() (+32 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "CampaignSources"
Cohesion: 0.15
Nodes (16): campaignSourceRow(), CampaignSources(), CampaignSourcesLive(), sourceGroupButton(), CampaignSourcesView, sourceRowTitle(), FiveToolsProgress(), importBar() (+8 more)

### Community 46 - "CustomMod"
Cohesion: 0.07
Nodes (24): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), UsesLabel(), CustomMod (+16 more)

### Community 47 - "vendorjs/main.go"
Cohesion: 0.18
Nodes (20): canonical(), changedPaths(), check(), checkSums(), crawl(), fetch(), fetchURL(), highestVersions() (+12 more)

### Community 48 - "SourcePriority"
Cohesion: 0.11
Nodes (37): ctxState, Props, CountEntriesBySourceRow, Chip(), Chips(), ChipsInput(), Combobox(), Content() (+29 more)

### Community 49 - "Path"
Cohesion: 0.04
Nodes (81): nullID(), activeFight(), TestCombatantNamesFollowSource(), TestConditionExpiresOnTurn(), TestDeleteActiveCombatant(), TestDeleteActiveCombatantLast(), TestDeleteActiveCombatantSkipsOut(), TestDeleteCharacterActiveCombatant() (+73 more)

### Community 50 - "testing.T"
Cohesion: 0.03
Nodes (120): Defense, TestDecode(), EquipmentOptions(), TestCoreTraits(), TestEquipmentOptions(), TestTokensReadAsLabels(), feature(), fiveToolsClasses() (+112 more)

### Community 51 - "floating_focus_manager.js"
Cohesion: 0.23
Nodes (17): addPreviouslyFocusedElement(), clearDisconnectedPreviouslyFocusedElements(), createFocusGuard(), enqueueFocus(), getEventType(), getFirstTabbableElement(), getFloatingFocusElement(), getPreviouslyFocusedElement() (+9 more)

### Community 52 - "Release Workflow"
Cohesion: 0.10
Nodes (19): GH_TOKEN, VERSION, jobs, release, name, permissions, contents, tags (+11 more)

### Community 53 - "useListNavigation"
Cohesion: 0.24
Nodes (12): doSwitch(), isCrossOrientationCloseKey(), isCrossOrientationOpenKey(), isMainOrientationKey(), isMainOrientationToEndKey(), isVirtualClick(), isVirtualPointerEvent(), useListNavigation() (+4 more)

### Community 54 - "starterEnv"
Cohesion: 0.30
Nodes (10): testEnv, pickerTicks(), serveStarterSets(), starterEnv(), TestImportsTickFromCampaignOnly(), TestPartialDownloadTicks(), TestPickerKeepsJobTicks(), TestSourcesCardStarterSets() (+2 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "ClassInfo"
Cohesion: 0.05
Nodes (38): LevelUpChoice, ListSubclassesOfRow, TestArmorGrants(), TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs() (+30 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (32): deny, files, rules, generated, paths, rules, formatters, enable (+24 more)

### Community 58 - "storeSource"
Cohesion: 0.12
Nodes (22): HomebrewOwner, Source, sourceFormat, TestHomebrew(), TestStoreSourceRollsBackOnAfterSave(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport() (+14 more)

### Community 59 - "redirect.go"
Cohesion: 0.26
Nodes (5): newConnListener(), route(), SplitTLS(), connListener, peekedConn

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "Player"
Cohesion: 0.09
Nodes (43): New(), NewClient(), drain(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCloseSessionAndCharacter(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements() (+35 more)

### Community 62 - "Sheet"
Cohesion: 0.07
Nodes (38): ClassFeature, Decode(), ClassData, SpellData, dnd5eFeatures(), Queries, Completions(), parseMastery() (+30 more)

### Community 63 - "Campaign"
Cohesion: 0.13
Nodes (10): ListAttachmentFilesRow, Campaign, attachmentRefs(), fresh(), inCampaign(), TestAttachmentRefs(), Roster, FileUse (+2 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "viewerCtx"
Cohesion: 0.18
Nodes (9): CanAddSources(), TestCanAddSources(), notFoundIsNil(), parseLiveKey(), TestLiveAffects(), liveKey, liveKind, liveView (+1 more)

### Community 66 - "handoff"
Cohesion: 0.28
Nodes (4): closeServers(), removeFile(), shutdownServers(), handoff

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "Command"
Cohesion: 0.06
Nodes (58): Amount, CondOpKind, Effects, userText, ArmorKindNamed(), parseFeature(), parseProf(), parseProfSet() (+50 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "ftPrerequisite"
Cohesion: 0.24
Nodes (7): ftClassFeature, ftClassRefs, ftCode(), ftPrerequisite(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 72 - "config.go"
Cohesion: 0.10
Nodes (22): loopbackHostOnly(), main(), parseFlags(), run(), TestLoopbackHostOnly(), TestParseFlags(), openBrowser(), Config (+14 more)

### Community 73 - "checkbox.js"
Cohesion: 0.31
Nodes (6): forwardClick(), inputOf(), requestCheckedChange(), setChecked(), setup(), sync()

### Community 74 - "ParseAidedd"
Cohesion: 0.20
Nodes (19): abilityMod(), ParseAidedd(), statLabel(), textLines(), firstInt(), TestParseAidedd(), ParsePlainPage(), attr() (+11 more)

### Community 75 - "switch.js"
Cohesion: 0.31
Nodes (5): forwardClick(), inputOf(), requestCheckedChange(), setup(), sync()

### Community 76 - "Store"
Cohesion: 0.06
Nodes (53): linkTarget, migration, copySource(), importOnce(), newStore(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+45 more)

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
Cohesion: 0.10
Nodes (8): ATTACHMENT_PREFIX, ATTACHMENT_SRC, Collab, editors, openWeb(), randomID(), KEEP_IMAGES, LINK_PATHS

### Community 85 - "Run Script"
Cohesion: 0.53
Nodes (8): changes(), GOFLAGS, lock(), restart_server(), run.sh script, stop_server(), superseded(), unlock()

### Community 86 - "mirror"
Cohesion: 0.36
Nodes (4): bookWriter, mirror, downloadBook(), filterDoc()

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - ".q"
Cohesion: 0.09
Nodes (29): testEnv, testEnv, mustJSON(), pendingSecret(), TestBlockDocGone(), TestClaimNeedsApproval(), TestCollabConflict(), TestCommandBuilder() (+21 more)

### Community 96 - "fields.js"
Cohesion: 0.18
Nodes (14): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+6 more)

### Community 98 - "live.js"
Cohesion: 0.33
Nodes (4): bindings(), reload(), roots(), within()

### Community 100 - "ftRec"
Cohesion: 0.27
Nodes (8): ftRec, ftCopyTargets(), ftDecode(), ftMod(), ftName(), ftReplaceText(), ftResolve(), ftVersions()

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - "newPlayerEnv"
Cohesion: 0.08
Nodes (39): TestEffectFields(), eventually(), testEnv, phoneGet(), phoneUpload(), shortLinger(), TestAtHomeNoteRules(), TestBringBack() (+31 more)

### Community 103 - "newEffectFix"
Cohesion: 0.10
Nodes (36): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestMonsterAtZeroHPUnaffected(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead() (+28 more)

### Community 105 - "cert.go"
Cohesion: 0.10
Nodes (18): Fingerprint(), generate(), LoadOrCreateCert(), seedCert(), seedGarbage(), TestCertNames(), TestCertUsableHostname(), TestKeyFileIsPrivate() (+10 more)

### Community 106 - "plainTags"
Cohesion: 0.15
Nodes (17): habitatsFrom(), ftCell(), ftFeatIncrease(), ftTable(), TestFiveToolsCell(), TestFtFeatIncrease(), TestHabitatsFrom(), TestStripTagsTokens() (+9 more)

### Community 107 - "ftBook"
Cohesion: 0.16
Nodes (19): ftBook, ftDistance, each(), ftAmount(), ftCastingTime(), ftDuration(), ftParts(), ftRange() (+11 more)

### Community 108 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 109 - "BlockEditor"
Cohesion: 0.22
Nodes (5): attachedPicture(), BlockEditor, json(), links(), sync()

### Community 111 - "Values Go and JS share"
Cohesion: 0.14
Nodes (15): Gotchas, internal/web, Notes, Values Go and JS share, TestEffectNoteMaxMatchesTag(), TestPagesRender(), TestNoRouteLiterals(), parseSharedJS() (+7 more)

### Community 112 - "ftItemData"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 113 - ".Router"
Cohesion: 0.21
Nodes (5): RequireAdmin(), assetETags(), immutableAsset(), sameOriginOnly(), useValidateTag()

### Community 116 - "buildReview"
Cohesion: 0.12
Nodes (24): buildReview(), categoryKey(), cell(), columnOverview(), formValues(), guessMapping(), headerScore(), httpHost() (+16 more)

### Community 117 - "TestSourceSwitchKeepsEdits"
Cohesion: 0.29
Nodes (7): UpsertCompendiumEntryParams, classCopy(), testEnv, TestRunCommandUsesCurrentSheet(), TestSourceCommandPermissions(), TestSourceSwitchKeepsEdits(), TestSourceWarningAfterImport()

### Community 118 - "build.sh"
Cohesion: 0.21
Nodes (11): build(), CGO_ENABLED, GOFLAGS, build.sh script, step(), lint.sh script, sources(), die() (+3 more)

### Community 119 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 125 - "State"
Cohesion: 0.07
Nodes (16): countingHandler, event, fakeHandler, Handler, patchMessage, stateHandler, viewPatch, Topic (+8 more)

### Community 127 - "EquipmentOption"
Cohesion: 0.40
Nodes (5): EquipmentItem, EquipmentOption, TestOptionText(), EquipmentChoice, optionText()

### Community 130 - "Before/after pairs"
Cohesion: 0.33
Nodes (5): After, Before/after pairs, Rules, Taking a pair: one script, The "before" instance (once per batch)

### Community 131 - "RosterCharacter"
Cohesion: 0.39
Nodes (5): devicesWhere(), deviceLabel(), rosterHP(), Device, RosterCharacter

### Community 132 - "jsconfig.json"
Cohesion: 0.33
Nodes (5): compilerOptions, module, target, exclude, include

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eHabitats(), TestOpen5eMastery(), TestUnflattenTable() (+12 more)

### Community 136 - "github.com/a-h/templ.Component"
Cohesion: 0.07
Nodes (69): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+61 more)

### Community 137 - "formErrors"
Cohesion: 0.40
Nodes (3): formErrors(), fieldError, pickError

### Community 139 - "DownloadFiveTools"
Cohesion: 0.12
Nodes (25): readBook(), DownloadFiveTools(), fetchJSON(), RemoveFiveTools(), replaceDir(), readDoc(), serveMirror(), TestDownloadFiveTools() (+17 more)

### Community 140 - ".Entry"
Cohesion: 0.11
Nodes (22): MonsterAction, addEntry(), setStat(), MonsterData, HomebrewInputFor(), abilityKey(), actionsText(), bonusText() (+14 more)

### Community 143 - "planSwitch"
Cohesion: 0.14
Nodes (12): featureSwitch, featureNames(), SourceRef, SubclassRef, SwitchFeature, SwitchPlan, PlanSourceSwitch(), planSwitch() (+4 more)

### Community 148 - "Client"
Cohesion: 0.14
Nodes (6): ClientInfo, goneDevice, message, Client, Hub, isDevice()

### Community 175 - "CompendiumKind"
Cohesion: 0.03
Nodes (68): Site, CoreTraits(), dataFor(), DecodeData(), ValidKind(), HomebrewInput, InputForEntry(), FetchPage() (+60 more)

### Community 176 - "CompendiumView"
Cohesion: 0.05
Nodes (55): Props, Variant, CountHabitatsRow, SearchCompendiumRow, Badge(), badgeClasses(), BuilderPickItems(), BuilderPicks() (+47 more)

### Community 178 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

## Knowledge Gaps
- **41 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+36 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 468 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **55 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Path()` connect `Path` to `compendium_templ.go`, `nullID`, `takehome_test.go`, `export/markdown.go`, `Button`, `ParseSheet`, `github.com/a-h/templ.Component`, `time.Time`, `empty_templ.go`, `perm.go`, `App`, `formSelect`, `github.com/a-h/templ.Attributes`, `CombatantView`, `GetImageRow`, `tabs_templ.go`, `WizardView`, `Slugify`, `CampaignSources`, `CompendiumKind`, `CompendiumView`, `Path`, `testing.T`, `SourcePriority`, `starterEnv`, `Player`, `Campaign`, `Store`, `ParseWiki`, `.q`, `newPlayerEnv`, `newEffectFix`, `.sendFeedback`, `Values Go and JS share`, `TestSourceSwitchKeepsEdits`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _41 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1477271132443546 - nodes in this community are weakly interconnected._
- **Why does `App` connect `App` to `compendium_templ.go`, `nullID`, `context.Context`, `.sourceIssues`, `encoding/json.RawMessage`, `time.Time`, `perm.go`, `Client`, `Book`, `Attachment`, `GetImageRow`, `Slugify`, `CompendiumKind`, `CompendiumView`, `ClassInfo`, `Sheet`, `Campaign`, `viewerCtx`, `handoff`, `Command`, `Store`, `.q`, `cert.go`, `.Router`, `Pending`, `State`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **Should `nullID` be split into smaller, more focused modules?**
  _Cohesion score 0.05213089802130898 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `Values Go and JS share` to `Path`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.045) - this node is a cross-community bridge._