# Graph Report - gm-session-manager  (2026-11-23)

## Corpus Check
- 366 files · ~638,523 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 56 file(s) not represented in the graph (top: .templ 48, .css 4, (none) 3)

## Summary
- 4656 nodes · 24086 edges · 149 communities (119 shown, 30 thin omitted)
- Extraction: 90% EXTRACTED · 10% INFERRED · 0% AMBIGUOUS · INFERRED: 2407 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `fb64fe8a`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- compendium_templ.go
- Sheet
- context.Context
- Condition
- createProvider
- dialog_templ.go
- SheetView
- combobox.js
- Option
- Path
- encoding/json.RawMessage
- select.js
- Command
- storeSource
- tooltip_templ.go
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
- Store
- Import
- redirect.go
- app.js
- tabbable.js
- useDismiss
- State
- viewerCtx
- command.js
- dialog.js
- NameDesc
- ParseWikidot
- tooltip.js
- CompendiumKind
- TextData
- config.go
- .compendiumDir
- composite.js
- buildReview
- Feature
- NewClient
- nullID
- Topic
- testing.T
- floating_focus_manager.js
- Release Workflow
- useListNavigation
- cert.go
- shadcn-templ Config
- monsterBody
- golangci-lint Config
- Pending
- Sheet
- radiogroup.js
- vendorjs/main.go
- ftItemData
- FiveToolsJob
- sqlc Config
- CombatantView
- Browser check
- Claude Settings
- select_templ.go
- use_transition_status.js
- playwright
- release.sh
- Hub
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
- names
- useTypeahead
- ftRec
- Validator Rule
- Go Module
- fields.js
- live.js
- internal_backdrop.js
- time.Time
- mounted
- .sendFeedback
- Player
- profFocus
- github.com/a-h/templ.Component
- CharacterFile
- ftBook
- testEnv
- BlockEditor
- Look
- EffectFieldsURL
- database/sql.NullInt64
- WeaponKind
- export/markdown.go
- build.sh
- sourcesEnv
- looseString
- errBodyWriter
- capReader
- popover_templ.go
- tabs_templ.go
- Before/after pairs
- coinList
- open5e.go
- testing.M
- db/CLAUDE.md
- perm/CLAUDE.md
- ftItemData
- ASIError
- Book
- .Entry
- sourceUsedError
- Item
- toast_templ.go
- github.com/a-h/templ.Component
- ListSharedWithRow
- SearchReadableBlocksRow
- deathFix
- ParseSheet

## God Nodes (most connected - your core abstractions)
1. `App` - 373 edges
2. `Path()` - 264 edges
3. `newTestEnv()` - 167 edges
4. `Sheet` - 131 edges
5. `Button()` - 123 edges
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
- `internal/compendium` --references--> `Slugify()`  [INFERRED]
  internal/compendium/CLAUDE.md → internal/compendium/entry.go

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Server-side visibility enforcement** — claude_server_side_visibility, claude_permissions_table, claude_note_block_visibility, claude_table_driven_tests [EXTRACTED 1.00]
- **Device identity and admin access model** — claude_session_tokens, claude_admin_rule, claude_identity_rule, claude_sessions_vs_game_sessions, readme_gm_other_device [INFERRED 0.85]
- **Command to live-view patch pipeline** — claude_commands_not_sheets, claude_minimal_ws_patches, claude_live_views, claude_ui_conventions [INFERRED 0.85]

## Communities (149 total, 30 thin omitted)

### Community 0 - "compendium_templ.go"
Cohesion: 0.14
Nodes (54): contextKey, keepAwake(), classRef, ftBookInfo, indexed, o5eAssignment, o5eChild, o5eClass (+46 more)

### Community 1 - "Sheet"
Cohesion: 0.09
Nodes (31): Cantrips(), Slugify(), fiveToolsPage(), AddResistance(), FreeCast(), SwapFeatSpell(), Sheet, Spell (+23 more)

### Community 2 - "context.Context"
Cohesion: 0.01
Nodes (90): keepAwake(), Site, FetchPage(), ParsePageURL(), SiteByKey(), TestParsePageURL(), ListImageCharactersRow, ListHomebrewLinksRow (+82 more)

### Community 3 - "Condition"
Cohesion: 0.05
Nodes (49): Combatant, Condition, CreateConditionParams, ListCombatantsRow, EffectMods(), ValidateEffectMods(), EffectEnd, passTurn() (+41 more)

### Community 4 - "createProvider"
Cohesion: 0.08
Nodes (26): applyDirectionalDamping(), applyLimited(), createProvider(), flushMouseLeave(), handleFocusGuard(), handleWindowBlur(), handleWindowFocus(), provider (+18 more)

### Community 5 - "dialog_templ.go"
Cohesion: 0.11
Nodes (43): ContentProps, ctxKey, ctxState, DescriptionProps, FooterProps, HeaderProps, OverlayProps, PortalProps (+35 more)

### Community 6 - "SheetView"
Cohesion: 0.05
Nodes (108): ArmorEffect(), Slot, Signed(), Button(), Variant, Action(), Card(), Content() (+100 more)

### Community 7 - "combobox.js"
Cohesion: 0.12
Nodes (55): afterSelectionChange(), allContents(), anchorFor(), applyFilter(), byId(), clearSelection(), close(), contentFor() (+47 more)

### Community 8 - "Option"
Cohesion: 0.14
Nodes (16): LevelUpPlan, SpellColumn(), spellChecklist(), TestSpellCounts(), LevelUpView, Option, SpellNeed, SubclassPanel (+8 more)

### Community 9 - "Path"
Cohesion: 0.04
Nodes (113): isCleared(), TestMarkdown(), nullID(), Path(), TestPath(), TestPathPanicsOnWrongIDCount(), activeFight(), TestCombatantNamesFollowSource() (+105 more)

### Community 10 - "encoding/json.RawMessage"
Cohesion: 0.08
Nodes (44): ftClass, ftEntry, ftFeat, ftModOp, ftMonster, ftNamed, ftRace, ftSpell (+36 more)

### Community 11 - "select.js"
Cohesion: 0.10
Nodes (53): allContents(), armCancelOpen(), arrowScrollStep(), clamp(), close(), closeAll(), commitValue(), contentFor() (+45 more)

### Community 12 - "Command"
Cohesion: 0.07
Nodes (54): Amount, CondOpKind, Effects, userText, parseFeature(), parseProf(), parseProfSet(), parseSet() (+46 more)

### Community 13 - "storeSource"
Cohesion: 0.09
Nodes (32): Source, TestStoreSourceRollsBackOnAfterSave(), FiveToolsImages(), parseFiveToolsBook(), parseFiveToolsBookReport(), readBook(), TestFiveToolsCopyTargets(), TestFiveToolsEdition() (+24 more)

### Community 14 - "tooltip_templ.go"
Cohesion: 0.09
Nodes (34): Content(), DisableTrigger(), state(), Tooltip(), Trigger(), BuilderPickItems(), BuilderPicks(), BuilderPick (+26 more)

### Community 15 - "floating_ui_dom.js"
Cohesion: 0.13
Nodes (51): autoUpdate(), frameLoop(), convertOffsetParentRelativeRectToViewportRelativeRect(), getBoundingClientRect(), getClientRectFromClippingAncestor(), getClippingElementAncestors(), getClippingRect(), getComputedStyle$1() (+43 more)

### Community 16 - "Viewer"
Cohesion: 0.06
Nodes (47): ListCharacterNamesRow, GetBlockRow, Visibility, Fields, ParseFields(), VisibleEntity(), VisibleLooks(), BlockVisibilities() (+39 more)

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
Cohesion: 0.06
Nodes (23): featureSwitch, LevelUpChoice, TestSlotMaxima(), armorCondition(), columnInts(), ParseEffects(), parseProfs(), TestParseEffects() (+15 more)

### Community 21 - "WizardView"
Cohesion: 0.10
Nodes (38): Props, Checkbox(), initialChecked(), Label(), initialValue(), Item(), RadioGroup(), state() (+30 more)

### Community 22 - "wikilink.js"
Cohesion: 0.16
Nodes (20): closeMenu(), imagesIn(), imageSrc(), isGM(), isImage(), KEEP_IMAGES, KeepImages, LAYOUTS (+12 more)

### Community 23 - "App"
Cohesion: 0.04
Nodes (27): SetCookie(), Entity, compendiumBack(), compendiumPath(), pathCampaign(), viewerCampaign(), entityTopics(), linkEntityLook() (+19 more)

### Community 24 - "popover.js"
Cohesion: 0.22
Nodes (29): allContents(), close(), closeAll(), closeNearest(), contentFor(), destroy(), idOf(), init() (+21 more)

### Community 25 - "github.com/a-h/templ.Attributes"
Cohesion: 0.02
Nodes (194): ActionProps, DescriptionProps, Props, TitleProps, Variant, Props, Props, Size (+186 more)

### Community 26 - "Project Rules (CLAUDE.md, README)"
Cohesion: 0.06
Nodes (28): graphify skill trigger, Admin via Loopback or .env Credentials, Collaborative Note Blocks (server authority, 409 rebase), graphify Usage Rules, HTTPS only with self-signed certificate, Device Identity (one role per device, logout = kick), Live Views (live kind, liveAffects, loadLive, Hot/Shape, bind), Note Block Visibility (gm, players, listed, owner) (+20 more)

### Community 27 - "Store"
Cohesion: 0.05
Nodes (47): fixture, tokens, migration, constantTimeEqual(), credentialsMatch(), newFixture(), serve(), TestCredentialsMatch() (+39 more)

### Community 28 - "Import"
Cohesion: 0.07
Nodes (44): sourceFormat, copySource(), importOnce(), newStore(), TestHomebrew(), TestImportIsIdempotent(), TestImportRefusesHomebrewKey(), TestMonsterImageBinding() (+36 more)

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

### Community 33 - "State"
Cohesion: 0.10
Nodes (11): countingHandler, fakeHandler, Handler, message, patchMessage, viewPatch, Client, drain() (+3 more)

### Community 34 - "viewerCtx"
Cohesion: 0.22
Nodes (7): notFoundIsNil(), parseLiveKey(), WithCampaign(), liveKey, liveKind, liveView, viewerCtx

### Community 35 - "command.js"
Cohesion: 0.18
Nodes (24): activateItem(), commandScore(), commandScoreInner(), filterItems(), findNextSibling(), findPreviousSibling(), formatInput(), getSelectedItem() (+16 more)

### Community 36 - "dialog.js"
Cohesion: 0.20
Nodes (23): closeDialog(), createInternalBackdrop(), destroyDialog(), dialogFor(), ensureDialog(), getDialog(), init(), isDialogOpen() (+15 more)

### Community 37 - "NameDesc"
Cohesion: 0.13
Nodes (25): firstTable(), LevelSpells(), OptionPrerequisite(), Options(), TestOptions(), NameDesc, CantripFits(), Cantrip (+17 more)

### Community 38 - "ParseWikidot"
Cohesion: 0.14
Nodes (23): block, ClassFeature, ClassData, cellAt(), classTable(), featuresUnder(), ParseWikidot(), slotCaster() (+15 more)

### Community 39 - "tooltip.js"
Cohesion: 0.26
Nodes (22): arrowOf(), close(), contentFor(), destroy(), emitOpenChange(), init(), joinGroup(), leaveGroup() (+14 more)

### Community 40 - "CompendiumKind"
Cohesion: 0.03
Nodes (78): HomebrewOwner, dataFor(), Decode(), DecodeData(), SpellData, ValidKind(), HomebrewInput, InputForEntry() (+70 more)

### Community 41 - "TextData"
Cohesion: 0.12
Nodes (20): internal/compendium, SkillChoice(), TestSkillChoice(), TextData, AbilityKeys(), AddFeat(), SkillKeys(), Spellbook() (+12 more)

### Community 42 - "config.go"
Cohesion: 0.24
Nodes (10): Config, lookupFunc, describe(), fromLookup(), Load(), newValidator(), lookupMap(), TestFromLookup() (+2 more)

### Community 43 - ".compendiumDir"
Cohesion: 0.11
Nodes (10): IsLoopback(), ClearCookie(), writeCookie(), RequireAdmin(), ImportSummary(), fiveToolsKey(), sameOriginOnly(), useValidateTag() (+2 more)

### Community 44 - "composite.js"
Cohesion: 0.18
Nodes (18): findNonDisabledListIndex(), getMaxListIndex(), getMinListIndex(), getOffset(), getStyles(), isElementDisabled(), isListIndexDisabled(), isModifierKeySet() (+10 more)

### Community 45 - "buildReview"
Cohesion: 0.12
Nodes (20): categoryKey(), cell(), columnOverview(), guessMapping(), headerScore(), httpHost(), importCategories(), isHTTPLink() (+12 more)

### Community 46 - "Feature"
Cohesion: 0.08
Nodes (20): FeatureEffects, findMod(), newMod(), parseFeatureSet(), parseMod(), parseUses(), SaveFeature(), UsesLabel() (+12 more)

### Community 47 - "NewClient"
Cohesion: 0.16
Nodes (17): New(), NewClient(), TestApprovedSurvivesDisconnect(), TestCatchUp(), TestCompendiumReachesEveryCampaign(), TestDeviceAnnouncements(), TestDisconnectDropsRequest(), TestDMGetsEveryCampaignsRequests() (+9 more)

### Community 48 - "nullID"
Cohesion: 0.15
Nodes (26): acBonusForm(), effectFix, newEffectFix(), notes(), TestAllyEffects(), TestConcentrationCascade(), TestEffectModRows(), TestEffectModsReadOnly() (+18 more)

### Community 49 - "Topic"
Cohesion: 0.24
Nodes (6): event, stateHandler, Topic, TopicKind, heldCollected(), heldTopics

### Community 50 - "testing.T"
Cohesion: 0.02
Nodes (127): Defense, TestIsLoopback(), TestDecode(), TestTokensReadAsLabels(), feature(), TestFiveToolsClasses(), RemoveSourceDir(), TestRemoveSourceDir() (+119 more)

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
Nodes (18): main(), run(), bookFolder(), CleanDownloads(), TestCleanDownloads(), Fingerprint(), generate(), LoadOrCreateCert() (+10 more)

### Community 55 - "shadcn-templ Config"
Cohesion: 0.12
Nodes (16): aliases, components, utils, iconLibrary, menuAccent, menuColor, rtl, $schema (+8 more)

### Community 56 - "monsterBody"
Cohesion: 0.17
Nodes (23): classBadges(), commaPrefixed(), compendiumFindURL(), levelsText(), parenthesized(), signedMap(), sortedKeys(), speeds() (+15 more)

### Community 57 - "golangci-lint Config"
Cohesion: 0.06
Nodes (31): deny, files, rules, generated, paths, rules, formatters, enable (+23 more)

### Community 58 - "Pending"
Cohesion: 0.24
Nodes (4): PendingState, randomID(), Pending, PendingKind

### Community 59 - "Sheet"
Cohesion: 0.08
Nodes (13): Part, Skill, TestCustomMods(), TestModifier(), skillByKey(), Sum(), TestDualWield(), damageLabel() (+5 more)

### Community 60 - "radiogroup.js"
Cohesion: 0.28
Nodes (10): groupOf(), initGroup(), inputOf(), itemOf(), itemsOf(), requestValueChange(), setupItem(), syncByInput() (+2 more)

### Community 61 - "vendorjs/main.go"
Cohesion: 0.28
Nodes (14): canonical(), check(), crawl(), fetch(), fetchURL(), highestVersions(), localPath(), main() (+6 more)

### Community 62 - "ftItemData"
Cohesion: 0.20
Nodes (11): ArmorData, ftItem, o5eArmor, ItemData, ftArmor(), ftItemData(), TestFiveToolsItemMastery(), armorData() (+3 more)

### Community 64 - "sqlc Config"
Cohesion: 0.14
Nodes (13): go, emit_empty_slices, out, overrides, package, column, go_type, sql (+5 more)

### Community 65 - "CombatantView"
Cohesion: 0.11
Nodes (36): addCombatantForm(), addCombatants(), beingBadge(), combatantChip(), combatantName(), combatantRow(), conditionDialog(), conditionList() (+28 more)

### Community 66 - "Browser check"
Cohesion: 0.29
Nodes (6): Browser check, Roles, Saved states, The test instance, What to check, You are alone

### Community 67 - "Claude Settings"
Cohesion: 0.25
Nodes (7): hooks, PreToolUse, permissions, allow, ask, deny, $schema

### Community 68 - "select_templ.go"
Cohesion: 0.07
Nodes (46): ImageCategory, Defenses, damageChoice(), Ptr(), firstValue(), usesItems(), weaponAbilities(), EffectFieldsForm() (+38 more)

### Community 69 - "use_transition_status.js"
Cohesion: 0.40
Nodes (7): animationsFinished(), close(), open(), reset(), resolve(), setOpen(), whenAnimationsFinish()

### Community 70 - "playwright"
Cohesion: 0.50
Nodes (3): npx, playwright, @playwright/mcp

### Community 71 - "release.sh"
Cohesion: 0.60
Nodes (4): die(), GOFLAGS, newer(), release.sh script

### Community 72 - "Hub"
Cohesion: 0.22
Nodes (3): ClientInfo, Hub, TestCloseSessionAndCharacter()

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
Nodes (5): pinnedRefs(), sourceIssue(), SourceCopy, SourceIssue, sheetRef

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

### Community 86 - "names"
Cohesion: 0.33
Nodes (6): ftRaceAbility(), SwitchFeature, SwitchPlan, names(), TestPlanSourceSwitch(), SourcePreviewView

### Community 87 - "useTypeahead"
Cohesion: 0.53
Nodes (4): useTypeahead(), getMatchingIndex(), isItemAvailable(), onKeyDown()

### Community 88 - "ftRec"
Cohesion: 0.21
Nodes (10): ftCopyRef, ftRec, ftCopyTargets(), ftDecode(), TestFiveToolsCopy(), ftMod(), ftName(), ftReplaceText() (+2 more)

### Community 96 - "fields.js"
Cohesion: 0.29
Nodes (12): editors, extensions, focusRow(), mount(), parts(), save(), setValue(), sync() (+4 more)

### Community 98 - "live.js"
Cohesion: 0.40
Nodes (3): bindings(), roots(), within()

### Community 100 - "time.Time"
Cohesion: 0.05
Nodes (31): Identity, FromContext(), ListAttachmentFilesRow, Campaign, Handlers, browserName(), firstLine(), mdLinkText() (+23 more)

### Community 101 - "mounted"
Cohesion: 0.83
Nodes (3): mounted(), register(), sync()

### Community 102 - ".sendFeedback"
Cohesion: 0.36
Nodes (10): bugForm(), testEnv, TestAutoReport(), TestBadConsole(), TestBugReport(), TestFeedbackAccess(), TestFeedbackIndex(), TestFeedbackTooBig() (+2 more)

### Community 103 - "Player"
Cohesion: 0.12
Nodes (24): Player(), TestCanRemoveCondition(), BenchmarkSheetState(), testEnv, Websocket and live views, TestNoteViewAudiences(), isRandomID(), shapeHash() (+16 more)

### Community 105 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (94): Props, Variant, CountHabitatsRow, CountSharedBySourceRow, ListEntryCopiesRow, ListMonstersByNameRow, SearchCompendiumRow, Habitat (+86 more)

### Community 106 - "CharacterFile"
Cohesion: 0.67
Nodes (3): CharacterFile, parseCharacterFile(), normalizeCharacterFile()

### Community 107 - "ftBook"
Cohesion: 0.18
Nodes (15): ftBook, ftDistance, habitatsFrom(), each(), ftAmount(), ftCastingTime(), ftDuration(), ftRange() (+7 more)

### Community 108 - "testEnv"
Cohesion: 0.13
Nodes (13): assetETags(), immutableAsset(), TestEntityFieldLinks(), testEnv, TestAdminLogin(), TestCampaignJoinLinks(), TestLogout(), classCopy() (+5 more)

### Community 111 - "Look"
Cohesion: 0.33
Nodes (9): Entity, Look, initials(), lookControls(), looksGrid(), npcPortrait(), currentLook(), EntityListView (+1 more)

### Community 112 - "EffectFieldsURL"
Cohesion: 0.33
Nodes (4): EffectFieldsURL(), EffectID(), effectQuery(), EffectTarget

### Community 113 - "database/sql.NullInt64"
Cohesion: 0.03
Nodes (48): ListHomebrewByOwnerRow, ListHomebrewRow, GetImageRow, ListImagesRow, Attachment, saveScreenshot(), ensureCategory(), imageItem() (+40 more)

### Community 116 - "WeaponKind"
Cohesion: 0.50
Nodes (4): Completions(), parseMastery(), TestWeaponKind(), WeaponKind()

### Community 117 - "export/markdown.go"
Cohesion: 0.13
Nodes (19): Links, pmMark, pmNode, writer, applyMarks(), escape(), Markdown(), plain() (+11 more)

### Community 118 - "build.sh"
Cohesion: 0.53
Nodes (5): build(), CGO_ENABLED, GOFLAGS, build.sh script, step()

### Community 119 - "sourcesEnv"
Cohesion: 0.28
Nodes (7): testEnv, sourcesEnv(), TestCampaignSourcePicks(), TestCampaignSourcesSet(), TestCompendiumAddresses(), TestSourceOverlap(), TestSubclassPickBySlug()

### Community 123 - "looseString"
Cohesion: 0.50
Nodes (3): looseString, o5eCreature, o5eItem

### Community 124 - "errBodyWriter"
Cohesion: 0.43
Nodes (3): errorText(), TestErrorText(), errBodyWriter

### Community 126 - "popover_templ.go"
Cohesion: 0.19
Nodes (12): Description(), state(), triggerID(), Align, ContentProps, ctxKey, ctxState, DescriptionProps (+4 more)

### Community 127 - "tabs_templ.go"
Cohesion: 0.09
Nodes (40): Content(), contentID(), IDFromContext(), initialValue(), List(), listVariant(), listVariantClasses(), orientation() (+32 more)

### Community 130 - "Before/after pairs"
Cohesion: 0.40
Nodes (4): After, Before/after pairs, Rules, The "before" instance (once per batch)

### Community 132 - "coinList"
Cohesion: 0.67
Nodes (3): Coins, coinList(), coin

### Community 133 - "open5e.go"
Cohesion: 0.13
Nodes (20): Entry, fixture, fixtures, normalizer, o5eProperty, TestOpen5eMastery(), TestUnflattenTable(), actionRank() (+12 more)

### Community 137 - "ftItemData"
Cohesion: 0.27
Nodes (6): ftClassFeature, ftClassRefs, ftCode(), keepReferencedFeatures(), featuresSearchText(), ftFeatureKey()

### Community 139 - "Book"
Cohesion: 0.12
Nodes (24): bookWriter, mirror, rawBook, bookGroup(), downloadBook(), fetchBooks(), fetchJSON(), filterDoc() (+16 more)

### Community 140 - ".Entry"
Cohesion: 0.09
Nodes (25): MonsterAction, addEntry(), setStat(), CoreTraits(), TestCoreTraits(), MonsterData, HomebrewInputFor(), abilityKey() (+17 more)

### Community 144 - "Item"
Cohesion: 0.18
Nodes (10): EquipmentItem, EquipmentOptions(), EquipmentOption, TestEquipmentOptions(), TestGearCommands(), Item, WornArmor, applyEquipment() (+2 more)

### Community 145 - "toast_templ.go"
Cohesion: 0.39
Nodes (7): Toast(), Toaster(), toastIcon(), toastListItem(), Props, ToasterProps, Type

### Community 146 - "github.com/a-h/templ.Component"
Cohesion: 0.05
Nodes (118): ListEntityFieldBacklinksRow, ListBacklinksRow, Field, Combobox(), Content(), Item(), List(), Description() (+110 more)

### Community 151 - "deathFix"
Cohesion: 0.29
Nodes (7): DeathSaves, deathFix(), effectFix, TestDeathClearsConditions(), TestQuickFormCritAtZeroHP(), TestTrackerCritAtZeroHP(), TestTrackerRefusesWhileDead()

### Community 159 - "ParseSheet"
Cohesion: 0.10
Nodes (40): TestUpgradeSheet(), ParseSheet(), TestChangeSubclass(), TestLevelUpEpicBoon(), TestLevelUpFeats(), TestLevelUpFeatWithoutASI(), TestLevelUpGoneSubclass(), TestLevelUpInvocations() (+32 more)

## Knowledge Gaps
- **19 isolated node(s):** `npx`, `@playwright/mcp`, `CGO_ENABLED`, `GOFLAGS`, `gm-session-manager` (+14 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 410 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **30 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `compendium_templ.go`, `Sheet`, `context.Context`, `Condition`, `Book`, `Command`, `Viewer`, `Item`, `State`, `viewerCtx`, `CompendiumKind`, `TextData`, `.compendiumDir`, `testing.T`, `cert.go`, `Pending`, `FiveToolsJob`, `Hub`, `.sourceIssues`, `time.Time`, `testEnv`, `database/sql.NullInt64`?**
  _High betweenness centrality (0.049) - this node is a cross-community bridge._
- **Are the 3 inferred relationships involving `Path()` (e.g. with `TestPath()` and `TestPathPanicsOnWrongIDCount()`) actually correct?**
  _`Path()` has 3 INFERRED edges - model-reasoned connections that need verification._
- **What connects `npx`, `@playwright/mcp`, `CGO_ENABLED` to the rest of the system?**
  _19 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `compendium_templ.go` be split into smaller, more focused modules?**
  _Cohesion score 0.1395869967787616 - nodes in this community are weakly interconnected._
- **Why does `Path()` connect `Path` to `compendium_templ.go`, `context.Context`, `dialog_templ.go`, `SheetView`, `tooltip_templ.go`, `Viewer`, `github.com/a-h/templ.Component`, `WizardView`, `App`, `deathFix`, `github.com/a-h/templ.Attributes`, `Store`, `Import`, `ParseSheet`, `CompendiumKind`, `nullID`, `testing.T`, `monsterBody`, `CombatantView`, `select_templ.go`, `ParseWiki`, `.sendFeedback`, `Player`, `github.com/a-h/templ.Component`, `testEnv`, `Look`, `EffectFieldsURL`, `database/sql.NullInt64`, `export/markdown.go`, `sourcesEnv`, `tabs_templ.go`?**
  _High betweenness centrality (0.047) - this node is a cross-community bridge._
- **Should `Sheet` be split into smaller, more focused modules?**
  _Cohesion score 0.08650937689050212 - nodes in this community are weakly interconnected._
- **Why does `Values Go and JS share` connect `testing.T` to `Path`, `Import`, `Command`, `wikilink.js`?**
  _High betweenness centrality (0.033) - this node is a cross-community bridge._