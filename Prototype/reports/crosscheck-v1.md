# Cross-check v1: architecture and study content against Abstract v8

Written 2026-10-23 by an independent reviewer. Read in full: `CMPT 863 Project Abstract v8.md`, `00 - Original Prompt (2026-10-23).md`, `01 - Study Flow Brief (Abstract v8 mapping).md`, `02 - Architecture.md`, `03 - Study Content v1.md`, `content/study-config.v1.json`. Also checked `CMPT 863 Project Abstract v8 - changes from v7.md` for one claim the content makes about it. The drafts were not edited.

Short names used below: **Arch** = `02 - Architecture.md`, **Content** = `03 - Study Content v1.md`, **JSON** = `content/study-config.v1.json`, **Brief** = `01 - Study Flow Brief`, **Abstract** = Abstract v8.

Findings are numbered 1 to 66 across all sections. Severity is one of **must fix**, **should fix**, **note**.

## What was checked and is consistent with Abstract v8

So that the findings below are read in proportion, these points were checked line by line and hold in both drafts:

- Two groups, labelled `checker` and `neutral` in both drafts. Round-robin allocation at sign-up, server side (Arch 5.5). The denial slide rotates over s2, s3, s4, balanced within each group (Arch 5.5, JSON `denialRotation`).
- Five slides, identical in both groups. "No issues found" on slides 1 and 5, a marking and flag on slides 2, 3 and 4 (JSON `scenes[s1..s5].panel`, Arch 4.2 lint).
- Per flagged slide: rating, action, question, answer, rating (Arch 5.2, JSON scene order).
- Who may ask: checker participants may ask or challenge; the chair asks otherwise and always in the neutral group (Content 4 "Who hears which question", JSON `sN.question.rule`).
- Two reasoned answers and one denial, same answer whoever asks (JSON `sN.answer.rule`).
- The four trust items (trust, competence, benevolence, integrity), the budget decision, both moderators (cybersecurity knowledge, own AI use), placed after the budget decision.
- All three flag labels and details are observations about the slide. None mentions honesty. None carries a confidence or accuracy figure. ("likely" is a verbal hedge and is in the Abstract's own example.)
- All 19 CISO lines in the JSON were read: none mentions AI, a model of AI, a tool or the checker.
- The debrief contains all three Abstract elements: no real issue on any slide, flags fixed in advance, what the study measures.
- Budget amount ($2.4 million) and what depends on it (insurance renewal, largest customer's contract) are identical for both groups and shown on slide 1 and slide 5.

## A. Violations of Abstract v8

**1. Debrief says the flags were drafted with a real AI tool, but they were not.**
- File: Content section 10 and section 4 intro; JSON `scenes[id=debrief].paragraphs.checker[0]`.
- Draft says: "The flags you saw were drafted with a real AI tool and then fixed in advance." Content section 4 says each flag is one "that a vision-language model would plausibly produce", that is, the flags were written by the content author.
- Abstract says: "flags were drafted with a real VLM from images of the slides, then fixed so all participants see the same output."
- Severity: **must fix**. As written, the debrief would tell participants something untrue, and the Method's claim would not hold.
- Fix: render the five slides as images, run a real VLM on each with a fixed prompt, and save model name, date, prompt and raw output (for example `Prototype/content/vlm-drafts/`). Pick and edit the flags from that output, and keep the edits listed. Until then, mark the debrief sentence as a placeholder. Add a step "VLM drafting run" to the content pipeline.

**2. The checker is introduced before the baseline rating, while the Abstract lists it after.**
- File: Arch 5.2 diagram (note on `Setup`: "Checker: checker introduced"), Content section 3 ("shown after the example scene of the setup step"), JSON scene order (`checkerIntro` before `opening` and `rating.T0`).
- Draft says: checker group learns about the AI checker during setup, before the opening and before T0.
- Abstract says (Study design, step 1): "Profile establishing the CISO's experience. First trust rating. Checker group: the checker is introduced." Validity: "we check that the groups start alike on the first rating."
- Severity: **should fix**. The Abstract's order is not fully explicit, and the Brief's flow table also places the introduction in setup, so the drafts followed the Brief. But introducing an AI checker before T0 can shift the checker group's baseline, which is exactly what the Validity check is meant to rule out.
- Fix: add an open point (see F, item 1.2) and a config option `checkerIntroPosition: "setup" | "afterBaseline"`. Suggested default: `afterBaseline` (after T0, before slide 1).

**3. Slide 4's chart already shows the reason, so the denial on slide 4 is not a bare denial.**
- File: Content section 4 slide 4 (a) and (b); JSON `scenes[id=s4].layout[id=s4.chart1]` and `.narration.text`.
- Draft says: the chart draws Phase 2 over months 3 to 7 and Phase 3 over months 5 to 8, and the narration says "You can see all three on the chart." The reasoned answer then says "That is why the bars overlap on the chart."
- Abstract says: "only the presence of a reason differs"; H3 contrasts a denial "with no reason" with a reasoned answer. Every flag is to be plausible.
- Severity: **should fix**. Any participant who looks at the chart can refute the flag alone, and a participant who gets the denial on slide 4 has the reason on screen anyway. Slide 2 has a milder version of the same problem: the diagram and narration put "the main data centre" at head office, which is half of the reason. The denial rotation spreads this over participants, but it adds slide-level noise exactly where H3 is tested.
- Fix: remove "You can see all three on the chart" from the narration. Show the chart without overlapping bars visible at a glance, for example a list of phases with start and end months in small type, so the slide stays correct but the overlap has to be worked out. Or pick a different arithmetic flag whose resolution is not drawn on the slide. Note the remaining effect as a limitation.

**4. Slide 3 narration gives a generic reason before the flag appears.**
- File: Content section 4 slide 3 (b); JSON `scenes[id=s3].narration.text`.
- Draft says: "Our finance team has checked every line, and each one is backed by a written quote from a supplier."
- Abstract says: the denial carries no reason; every flag is plausible but wrong.
- Severity: **should fix**. This sentence pre-empts the price flag in both groups and gives the slide-3 denial an implicit reason ("finance checked it"). It weakens H1 on slide 3 and H3 when the denial lands on slide 3. Slide 2's "This is a simple, proven design" does the same more mildly.
- Fix: delete the slide 3 sentence. Change slide 2's last sentence to "That is the design." or delete it.

**5. The slide 4 reasoned answer is itself slightly wrong.**
- File: Content section 4 slide 4 (c); JSON `scenes[id=s4].answers.reason.text`.
- Draft says: "Phases two and three run at the same time, with different teams."
- Abstract says: all slides are error-free and the CISO is right; the reason explains why the flag is wrong.
- Severity: **should fix**. Phase 2 runs months 3 to 7 and Phase 3 runs months 5 to 8; they overlap for three months, not the whole time. An attentive participant could catch the CISO in an inaccuracy, which would be a real error in a study whose premise is that the CISO is right.
- Fix: "Phase three starts while phase two is still running, with a different team."

**6. Same tone for reason and denial is not protected in the pipeline.**
- File: Arch 4.1 `AudioLine.ssml`, Arch 4.2 lint, Arch 11.
- Draft says: any line may carry its own SSML; `VoiceSpec` has `speakingRate` and `pitch` per voice.
- Abstract says: "gives the reasoned answers and the denial with the same confidence; only the presence of a reason differs."
- Severity: **should fix**. Per-line SSML allows prosody, emphasis or rate differences between a reason and its denial. Also, TTS renders the shared first and last sentences differently depending on what surrounds them, so the "same sentence" will not sound the same.
- Fix: lint rule: CISO answer lines may not contain `<prosody>`, `<emphasis>` or rate changes, and reason and denial of one slide must use the same SSML wrapper. Better: synthesise the shared first sentence, the middle, and the shared last sentence as separate clips and join them, so the shared sentences are acoustically identical in both versions.

**7. Neutral group timeout is recorded as a checker-group action.**
- File: JSON `scenes[id=s2.action.neutral].timeoutChoice` (also s3, s4).
- Draft says: `"timeoutChoice": "none"`, while the neutral options are `yes` and `no`.
- Abstract says: neutral participants "say whether they have questions about the slide"; H4a compares that with checker participants who question the CISO.
- Severity: **should fix**. A neutral timeout must not be counted as "no question" or as a checker-type "say nothing"; H4a needs it as missing.
- Fix: record neutral timeouts as `timedOut: true` with no option id; the chair asks in any case.

## B. Decisions made that should have stayed open

**8. An invented 6-word limit for flag labels, used to change Oliver's own example.**
- File: Content section 4 slide 3 "Note on the label", section 13 point 8.
- Draft says: Oliver's "Firewalls likely too expensive, cheaper ones available" "has 7 words, one over the 6-word limit", and "the brief's limit is 6".
- Authority: neither the Brief, the Abstract nor the prompt sets a word limit. Oliver's wording comes from the prompt.
- Severity: **should fix**.
- Fix: restore Oliver's wording as the default label in Content and JSON (`scenes[id=s3].marking.label`, the challenge question, and `flagReview.items[1].label`), and turn "should labels have a length limit?" into an open point if wanted.

**9. Content calls its checker introduction "the default per the brief", but it is the Brief's alternative.**
- File: Content section 3; JSON `scenes[id=checkerIntro]._open`.
- Draft says: "Default per the brief: shown after the example scene of the setup step, as a short text screen."
- Brief says (OPEN A): default is the example scene that "additionally shows the panel and explains it in one sentence"; "Alternative: a separate text screen."
- Severity: **should fix**. Arch already has `checkerIntro: "exampleScene" | "textScreen"`, so both forms need content.
- Fix: label the text screen as the alternative, and add the one-sentence version for the default (for example spoken or captioned during the example scene: "In the meeting, a slide checker will look at each slide and mark possible issues.").

**10. Content's checker introduction decides the display question (glasses on the slide versus a panel).**
- File: Content section 3; JSON `scenes[id=checkerIntro].text`.
- Draft says: "you wear smart glasses with a slide checker ... it looks at each slide through the camera in your glasses ... shows a short note in the checker panel."
- Authority: Abstract: "a panel shows the camera image of each slide"; prompt: "like we are wearing AR glasses". Arch lists this as FO-2, open.
- Severity: **should fix**.
- Fix: write two intro variants keyed to `checkerDisplay` (`overlay`: glasses wording; `sidePanel`: "a tablet in front of you shows the camera image of each slide"). Leave the choice to Oliver.

**11. Allocation happens before the device gate, so blocked visitors use up group and denial slots.**
- File: Arch 5.5 and 5.2 (`Created --> DeviceGate`), FO-7.
- Draft says: allocation at `POST /sessions`; FO-7 mentions only people who decline consent.
- Brief says: allocation "at session creation" (step 0), round robin "so that the groups stay equal in size" (Abstract).
- Severity: **should fix**. Every phone visitor who is blocked by OPEN R takes a group and a denial slot. With 40 participants and many recruits on LinkedIn on phones, this can unbalance completed groups noticeably.
- Fix: extend FO-7 to cover blocked devices, and propose allocating at `consent_accepted` (after the device gate), not at link opening.

**12. The chair's "routine question" was made slide-specific and points at the flagged element.**
- File: Content section 4 (c) for slides 2 to 4; JSON `scenes[id=sN].routineQuestion.text`.
- Draft says, for example slide 3: "The firewalls are a large line. Could you say a little more about that cost?"
- Abstract says: "the board chair asks a routine question" in both groups.
- Severity: **should fix**. A question aimed at exactly the flagged line is closer to a mild challenge than a routine question, and in the neutral group it draws attention to the very element the checker flags in the other group. This may be the right design (it is needed so one answer fits all questions), but it is a study-design choice and is not listed as open.
- Fix: add as open point (F, item 1.10) with two options: slide-specific (current) or one generic routine question ("Thank you, Daniel. Before we move on, can you tell us why you chose this approach?"). See also finding 50 on answer fit.

**13. Scenario facts were set without being listed for confirmation.**
- File: Content section 1; JSON `scenario`.
- Draft says: company in western Canada, freight and warehousing, 1,400 staff, $480 million revenue, $2.4 million CAD request, insurance renewal in 10 months, CISO with 6 plus 12 years' experience, chair female.
- Authority: the Abstract fixes only that the amount and stakes are identical and shown, and that a profile establishes experience. Content lists gender (point 9) and currency (point 12) but not the amount, the stakes, the sector or the level of experience.
- Severity: **note**. The choices are reasonable. The amount and experience level matter to the trust baseline and the budget decision.
- Fix: one open point "confirm the scenario" (F, item 1.11).

**14. Profile delivery and the CISO's greeting before T0 were decided as "proposals" but not listed.**
- File: Content section 1 ("Proposal for the profile"), section 2 line 5.
- Draft says: the chair reads the profile aloud with a card on screen; the CISO says one greeting line before the baseline rating.
- Severity: **note**. Hearing the CISO's voice before T0 is sensible, but it affects the baseline and should be confirmed.
- Fix: add to the open list (F, item 1.12).

**15. Response scales for the two moderators were chosen without being listed.**
- File: Content section 8 items 6 and 7; JSON `background.items[id=cyberKnowledge]` and `[id=aiUse]`.
- Draft says: cybersecurity knowledge 1 to 5 (0 to 10 noted as alternative in JSON only); AI use six frequency categories.
- Abstract says: both are planned moderators of H1 and H2.
- Severity: **note**. The scales decide how the moderation analysis can be run.
- Fix: list as an open point (F, item 1.14).

**16. Budget and background screens are untimed by default.**
- File: Content section 7 ("Untimed"); JSON `budget.windowMs: null`.
- Authority: the prompt times only the view and the rating; nothing says the budget is untimed.
- Severity: **note**. Untimed is a sensible reading, but the prompt's spirit ("react rather than deliberate") applies to the trust items only; confirm.
- Fix: list (F, item 2.12).

**17. Several application defaults were introduced without being listed.**
- File: Arch 5.3 and 6 (consent at least 3 s, debrief at least 5 s, `audioEndToleranceMs` 1,500, `latencyAllowanceMs` 3,000, duplicate window 60 min, rate limits 5 and 30 per hour), Arch 7.1 ("Started" means before the first `scene_started`).
- Severity: **note**. These are application parameters, not study design, and most are configurable. The "Started" versus "Ongoing" definition is visible to Oliver in the admin list and should be confirmed.
- Fix: list the status definition (F, item 2.13).

## C. Inconsistencies between the architecture and the content

**18. The JSON cannot be loaded by the TypeScript types in Arch section 4.1.**
- File: Arch 4.1 versus the whole JSON.
- Severity: **must fix**. The schema generated from the types would reject the content file at CI, admin upload and server start (Arch 4.1).
- Field-by-field mismatches:

| Path | Arch 4.1 type | JSON | Mismatch |
|---|---|---|---|
| `version` | string, e.g. `"2026-10-23.1"` | `"1.0.0"` | Format differs (note only) |
| `locale` | required `"en"` | missing | Missing |
| `timing.participantQuestionShowMs` | required | `questionTextMs` | Name differs |
| `timing.audioEndToleranceMs`, `latencyAllowanceMs`, `abandonAfterMs` | required | missing | Missing |
| `timing.postAnswerViewMs`, `msPerWordEstimate` | not in type | present | Extra |
| `voices` | keys `ciso`, `chair`, `member` | `ciso`, `chair`, `participant`; extra `provider`, `_open`; `participant.name: null` | Key and fields differ |
| `options` (StudyOptions) | required | missing | Missing |
| `consent` (ConsentSpec) | required | missing | Missing (OPEN M) |
| `setup` (SetupSpec) | required top-level object | scene of `type: "audioTest"` with nested `exampleScene`, plus scene `checkerIntro` | Different place and shape |
| `rating.items[]` | `{ id, label, scale: { min, max, minLabel, maxLabel } }` | top-level `ratingItems[]` with `{ id, text, min, max, anchors: {"0","10"}, meti }`, plus `ratingScreen.instruction` | Name and shape differ; instruction has no home in the type |
| `actions.checker` | array of `{ id: "nothing"|"explain"|"challenge", label }` | object `{ prompt, options: [{ id: "none", text, showsQuestion }], note }` | Shape, id `none` versus `nothing`, `text` versus `label` |
| `actions.neutral.options[].id` | `"noQuestion" | "hasQuestion"` | `"yes" | "no"` | Ids differ |
| `actions.timeoutMapsTo` | required | per scene `timeoutChoice` | Different place |
| top-level `labels`, `colors`, `denialRotation`, `scenario` | not in type | present | Extra |
| `scenes[].kind` | discriminator | `type` | Name differs |
| `scenes[].groups` | not in type | `"both" | "checker" | "neutral"` | Extra; Arch builds group differences in the plan compiler |
| scene kinds | opening, slide, rating, action, question, answer, budget, background, flagReview, debrief | adds `audioTest`, `checkerIntro`; two `action` scenes per slide (one per group) | Extra kinds, duplicated action scenes |
| `AudioLine.id` | required | missing on every line; lines carry `speaker` instead | Missing |
| `AudioLine.audioKey` | `"<sha256>.mp3"`, filled at publish | human key such as `"s2.answer.denial"` | Meaning differs (see finding 19) |
| opening scene | `{ layout, lines, profile, budget }` | `{ introText, profileCard, lines, panel, marking }` | No `layout` or `budget`; `profileCard` versus `profile` |
| `SlideScene.index` | required 1 to 5 | missing | Missing |
| `SlideScene.layout` | `SlideLayout { template, title: TextSeg, blocks: Block[] }` | flat array of `{ id, kind, text }`, plus string `title` | Shape differs |
| number block | `{ value: TextSeg, caption: TextSeg }` | `{ kind: "number", text }` | Shape differs |
| image block | `{ assetKey, alt }` | `{ kind: "image", text: <description> }` | No asset |
| chart block | `{ chartType, data, caption }` | `{ kind: "chart", text: <description> }` | No data |
| `SlideScene.checker` | `{ status: "noIssues", text }` or `{ status: "flagged", markings[] }` | `panel: { groups, state: "noIssue" | "flag", appearAtMs? }` and separate `marking` | Name, values, structure differ |
| `Marking` | `{ id, targetId?, bbox?, appearAtMs?, appearAtMark?, color: hex, shape, label, detail }` | `{ groups, targetIds[], appearAtMs, anchorText, color: "flag", label, detail, _note }` | No `id`, no `shape`, array target, colour token not hex, extra fields |
| `SlideScene.flagged` | `{ routineQuestion, participantQuestions, answers }` nested | the three fields at slide top level | Nesting differs |
| `TextLine` (participant questions) | `{ id, text, audio?: AudioLine }` | `{ text, audioKey, durationMs }` | No `id`, flattened audio |
| rating scene | `{ kind, id, point, after: Id }` | `{ type, id, ratingKey, after: "s2 (flag)", items: "ratingItems", windowMs: "timing.ratingWindowMs" }` | `point` versus `ratingKey`; `after` is not an id; string references in numeric fields |
| action, question, answer | `forSlide` | `slideRef`, plus `rule` strings | Name differs |
| budget | `{ spec: BudgetSpec }` | inline `prompt`, `options`, `confidence`, `windowMs`, `alternatives` | Shape differs |
| flagReview | `{ prompt, scale: ScaleSpec }` | `{ intro, items[{ slideRef, label, options: ["Correct","Incorrect"] }], enabled }` | Numeric versus binary; see finding 25 |
| debrief | `{ text: string }` | `{ paragraphs: { both, checker, neutral, end } }` | Group-specific text has no home in the type |
| `_status`, `_open`, `_note` keys | not in type | many | Schema must allow or strip `_`-prefixed keys |

- Fix: decide which side is the source. Recommended: keep Arch's TS types as the source (they carry the ids the events need), extend them where the content found real needs (group-specific debrief, `postAnswerViewMs`, `noIssues.appearAtMs`, `targetIds[]`, rating instruction, profile card, intro text, budget follow-up), and write JSON v1.1 to that schema. Add a CI test that loads `content/study-config.*.json` with the generated schema.

**19. `audioKey` means two different things, and the content's meaning would leak the denial.**
- File: Arch 4.1 (`audioKey: "<sha256>.mp3"`) and 4.3 ("Audio files are named by content hash, so a URL does not say reason or denial"); JSON every `audioKey`, for example `"s2.answer.denial"`.
- Severity: **must fix**. If the TTS tool or a developer names files after the JSON key, the participant's network tab shows `s2.answer.denial.mp3`.
- Fix: rename the JSON field to `id` (line id) and leave `audioKey` for the hash filled at publish. Lint: `audioKey` must match `^[0-9a-f]{64}\.mp3$` in resolved configs.

**20. View time after the CISO's answer: same open point, opposite defaults.**
- File: Arch FO-6 (default: 10 s everywhere a rating follows audio, using `postAudioViewMs`); Content section 13 point 4 and JSON `timing.postAnswerViewMs: 0`.
- Severity: **should fix**. The server would apply 10 s, the content author assumed 0. Content section 12's time estimate also leaves out the 30 s this adds.
- Fix: one parameter `postAnswerViewMs` in Arch `Timing`, one entry in the merged open list (F, item 1.8), one default.

**21. How "No issues found" is carried and when it appears.**
- File: Arch 4.1 `CheckerPanel` (`{ status: "noIssues"; text }` per slide, no timing); JSON `labels.noIssueLabel` plus `panel.appearAtMs: 8000` on s1 and s5.
- Severity: **should fix**.
- Fix: add `appearAtMs` to the `noIssues` variant; let `text` default to a config-level label so the wording is in one place.

**22. Three different ways to time a marking.**
- File: Arch 4.1 (`appearAtMs` or `appearAtMark`, an SSML `<mark>` resolved by `tools/tts`); JSON `marking.appearAtMs` plus `anchorText`.
- Severity: **should fix**. `anchorText` is not read by any tool in Arch, so after TTS someone must re-time by hand.
- Fix: in the narration SSML, put `<mark name="flag"/>` before the anchor phrase and set `appearAtMark: "flag"`. Keep `anchorText` only as an author note. Check that the chosen voice returns timepoints (Arch 11 step 3 already warns).

**23. One marking with several targets versus one target per marking.**
- File: Arch 4.1 `Marking.targetId` (single); JSON `s2.marking.targetIds` (2 targets), `s4.marking.targetIds` (4 targets).
- Severity: **should fix**. Under the Arch type, slide 4 would need four markings, each with its own `label` and `detail` callout, and four `marking_shown` events.
- Fix: `targetIds: Id[]` on `Marking`, one callout per marking; `marking_shown` keyed by marking id. Give each JSON marking an `id`.

**24. Action ids and timeout handling disagree.**
- File: Arch 4.1 `ActionSpec` (`nothing`, `explain`, `challenge`; `noQuestion`, `hasQuestion`; `timeoutMapsTo: "nothing"`), Arch 5.3 `action_chosen`, Arch 7.1 `answers.actions`; JSON `actions` (`none`; `yes`, `no`) and `timeoutChoice: "none"` on neutral scenes.
- Severity: **should fix**. See also finding 7.
- Fix: one id set; timeouts stored as `timedOut: true` plus, for the checker group only, the mapped action.

**25. Flag review format disagrees.**
- File: Arch 4.1 (`flagReview: { prompt, scale: ScaleSpec }`), Arch 5.3 (`flag_review_submitted { ratings: { [markingId]: n } }`); JSON `flagReview` (binary Correct or Incorrect, keyed by `slideRef`).
- Severity: **should fix**. The Abstract borrows Prike's "rating of how true each checked claim is", which suggests a scale; the Brief's proposal says correct or incorrect. Both are possible, but the two drafts must agree.
- Fix: keep this under OPEN G with both forms; type `FlagReviewItem { markingId, label, response: "binary" | ScaleSpec }`.

**26. Audio check: attempt count and answer ids disagree.**
- File: Arch FO-5 ("two retries, then continue", that is three attempts) and 5.3 (`audio_check_answered { choiceId }`); JSON `audioTest.maxAttempts: 2`, options as plain strings, `correctOption: "lantern"`.
- Severity: **should fix** (small).
- Fix: one number, stated as attempts. Give options ids.

**27. Ratings after slides 1 and 5 are switched in two places.**
- File: Arch 4.1 `options.ratingAfterUnflaggedSlides`; JSON has fixed scenes `rating.T1` and `rating.T5` (marked OPEN B).
- Severity: **should fix**. If Oliver chooses 7 ratings, someone must both flip the option and delete scenes.
- Fix: the plan compiler drops T1 and T5 when the option is false; the JSON keeps them and marks them `optionalBy: "ratingAfterUnflaggedSlides"`.

**28. Group-specific content in the public scene pack.**
- File: Arch 4.3 (public pack contains "opening, rating items, budget and background forms, debrief"); JSON debrief with `checker` and `neutral` paragraphs, `checkerIntro`, `flagReview` with flag labels.
- Severity: **should fix**. A participant who opens the network tab before the end can read the debrief ("None of the five slides had a real issue", "In the other group, a checker flagged slides 2, 3 and 4") and the flag labels. Arch's own rule is that the public pack holds no markings and nothing group-specific.
- Fix: serve debrief, checker introduction and flag review texts in the step descriptor at that step, like the overlay. Keep only truly shared texts in the public pack.

**29. Lint would fail the content: "model" in a CISO line.**
- File: Arch 4.2 (CISO lines may not contain "model"); JSON `scenes[id=s3].answers.reason.text` ("An office model would not cope").
- Severity: **should fix**.
- Fix: reword to "An office firewall would not cope with our warehouses", and narrow the lint to "AI model", "language model".

**30. Opening scene: what the presentation screen shows is undefined.**
- File: Arch 4.1 opening needs `layout` and `budget: BudgetCard`; JSON opening has no layout and puts the amount on slide 1; Content section 2 has only a profile card during line 2.
- Severity: **note**.
- Fix: add a title slide for the opening (for example "Item 4: Protection of our network") and decide whether the amount card appears here as well as on slide 1.

**31. Progress total N is defined differently.**
- File: Brief section 3 ("counts the logged steps"); Arch 5.1 (plan length, which includes post-view steps, `checkerIntro`, `flagReview`). Content does not define it.
- Severity: **note**. Under the defaults both drafts agree on 9 rating points. N will differ by group (checker introduction and flag review), which shows the group in the admin only, which is fine.
- Fix: define N as the number of participant-facing steps in the plan and show it in the admin column header.

**32. Participant question voice key differs.**
- File: Arch 4.1 voice `member`; JSON voice `participant`, `participantQuestions.*.audioKey` without a voice.
- Severity: **note**. Covered by finding 18, listed because OPEN D depends on it.

## D. Technical errors in the architecture

**33. Event `seq` breaks with server-written events and with a reload. (confirmed error)**
- File: Arch 5.3 ("rejects anything that is not exactly `lastSeq + 1`"), 5.4 (server writes synthetic `rating_timeout`), 7.1 (`events/{seq}`, rejected events stored too), 9.2 (`<StepRouter>` owns the counter), 8.1 (`GET /next` returns no seq).
- Problem: if server-written timeouts and stored rejected events take seq numbers, the client's next event is a `seq_gap`. If they do not, they need another document id, which Arch does not give. After a reload the client has no counter at all.
- Severity: **must fix**.
- Fix: keep two sequences: `clientSeq` (checked as now) and a server-side `eventNo` used as the document id for every stored event, including rejected and synthetic ones. Return `expectedSeq` in every response and in `GET /next`.

**34. Lazy timeouts ignore the latency allowance. (likely error)**
- File: Arch 5.3 (`rating_submitted` after `ratingClosesAt + lat` is "late"), 5.4 (lazy evaluation writes a timeout as soon as the window "closed").
- Problem: a rating sent just before the deadline that arrives within `lat` after `ratingClosesAt` finds lazy evaluation has already timed it out and moved to the action step, so it is rejected with `wrong_state`, not stored as on time or late.
- Severity: **should fix**.
- Fix: lazy evaluation treats a window as expired only at `closesAt + lat`. Late submissions are only possible after that and are stored as late side events.

**35. A slow `scene_started` can strand a real participant. (likely error)**
- File: Arch 5.3 (`audio_ended` must satisfy `now >= t0 + d - tol`, `t0` = server receipt of `scene_started`; earlier is rejected).
- Problem: if `scene_started` is delayed (cold start 1 to 2 s, slow network) by more than `tol` (1.5 s) plus the delay of `audio_ended`, a participant who heard the whole clip gets `too_early` and has no way to send it again, because the audio has ended.
- Severity: **should fix**.
- Fix: on `too_early`, return `retryInMs` and have the client resend `audio_ended` after that wait. Also: the client sends `scene_started` and starts playback on the response, not before.

**36. "Continue" after an audio failure cannot work with the timing check.**
- File: Arch 9.3 (stall shows a "Continue" button, `audio_problem` side event) and 5.3 (no early `audio_ended`). `audio_problem` is not in the event table.
- Severity: **should fix**.
- Fix: add `audio_problem` to the event table, and define that it lets the step advance at `t0 + d` at the earliest (the participant waits out the clip with the transcript or a message).

**37. Multi-clip scenes have no timing rule.**
- File: Arch 5.3 (`scene_started { sceneId, audioKey }`, one `d`), 11 step 5 (pauses between lines are "scene timing"); JSON opening has five clips.
- Severity: **should fix**.
- Fix: either concatenate a scene's lines into one clip at publish (simplest, and the pause becomes part of the audio), or add `line_started` and `line_ended` events with per-line checks.

**38. The step descriptor lacks the rating point and the transcript.**
- File: Arch 8.1 `Step`; 5.3 `rating_submitted { point, ... }`; 9.6 (transcript for screen readers).
- Severity: **should fix**. The client must send `point` but is never told it; the public pack has no answer texts, so screen-reader transcripts for answers and questions have no source.
- Fix: add `point?: string` and `transcript?: string` to `Step`, or let the server infer `point` from the current step.

**39. Terraform IAP OAuth brand and client. (likely error)**
- File: Arch 12.1 module `iap` ("OAuth brand and client").
- Problem: the IAP OAuth Admin API behind `google_iap_brand` and `google_iap_client` was deprecated in January 2025 with shutdown announced for March 2026. On 2026-10-23 these resources should be expected to fail. IAP now offers a Google-managed OAuth client, which only admits users in the same Google Workspace or Cloud organisation.
- Severity: **must fix** for the Terraform module.
- Fix: drop both resources. If all admins are in one organisation, enable IAP on the backend service with the Google-managed client. Otherwise create an OAuth client by hand in the console, store its secret in Secret Manager, and pass it to `google_compute_backend_service.iap`. Verify against current docs before building.

**40. Firestore daily backups cannot be kept 14 days. (likely error)**
- File: Arch 12.4 ("daily backup kept 14 days").
- Problem: Firestore scheduled backups allow at most 7 days of retention for daily schedules; longer retention needs a weekly schedule (up to 14 weeks).
- Severity: **should fix**.
- Fix: daily kept 7 days plus weekly kept 14 weeks, or daily 7 days only. PITR (7 days) is already on.

**41. `/healthz` is not routed to the service. (confirmed error, as written)**
- File: Arch 12.1 `network` (URL map: bucket default, `/api/v1/*`, `/admin/*`, `/api/admin/*`), 12.1 `monitoring` (uptime check on `/healthz`), 13 deploy step 5.
- Problem: `/healthz` falls to the bucket default and returns 404; the `run.app` URL is blocked by the ingress setting.
- Severity: **should fix**.
- Fix: serve health at `/api/v1/healthz` (or add a `/healthz` path rule to the API backend).

**42. `/s/*` rewrite to `/index.html`. (likely error if done with a prefix rewrite)**
- File: Arch 12.1 `network`.
- Problem: a `pathPrefixRewrite` replaces only the matched prefix, so `/s/k7m2qp` would become `/index.htmlk7m2qp`.
- Severity: **should fix**.
- Fix: a route rule with `pathTemplateMatch: "/s/**"` and `pathTemplateRewrite: "/index.html"` on the global external Application Load Balancer, or let the SPA live at `/s/index.html` with a 404 fallback. Test with `/s/<code>/blocked` too.

**43. `ci-terraform-plan` cannot take the state lock. (confirmed error with default settings)**
- File: Arch 12.3 ("read on the state bucket").
- Problem: `terraform plan` with the GCS backend writes a lock object by default; read-only access fails.
- Severity: **should fix**.
- Fix: grant `roles/storage.objectUser` on the state prefix, or run plan with `-lock=false` in CI and say so.

**44. Safari may block the preloaded second audio element. (likely error)**
- File: Arch 9.3 (one reused `HTMLAudioElement`, but prefetch on "a second element, swapped in"; the setup click "gives that gesture for the rest of the page life").
- Problem: Chrome's sticky activation covers the whole page; Safari unlocks playback per media element, so a fresh element may refuse `play()` without a new gesture.
- Severity: **should fix**.
- Fix: keep exactly one element and swap `src`, prefetching with `fetch()` into the HTTP cache, or play everything through one Web Audio `AudioContext` unlocked at setup. Include desktop Safari in the Playwright run (WebKit).

**45. Cloud Scheduler needs an invoker identity for the sweeper job. (likely omission)**
- File: Arch 12.3 IAM table.
- Severity: **note**.
- Fix: add a scheduler service account with `roles/run.invoker` on the sweeper job. Also consider dropping the sweeper: the admin list can compute "Abandoned" from `lastEventAt` at read time.

**46. Billing budget needs billing-account rights. (likely omission)**
- File: Arch 12.1 `monitoring` ("billing budget").
- Severity: **note**.
- Fix: `google_billing_budget` needs Billing Account Costs Manager or Administrator on the billing account, which a project-scoped apply account usually lacks. Create the budget by hand or grant that role.

**47. Firestore free tier applies to the `(default)` database only. (note)**
- File: Arch 12.1 `firestore`, 12.5 ("Firestore 0").
- Severity: **note**.
- Fix: create the database with name `(default)`, or budget for paid usage on a named database.

**48. X-Forwarded-For handling. (unsure)**
- File: Arch 6 "Client IP".
- Assessment: for the global external Application Load Balancer, Google documents the appended pair `<client-ip>,<load-balancer-ip>`, so "second to last" is correct. Whether Cloud Run behind a serverless NEG appends anything further was not confirmed.
- Severity: **note**.
- Fix: log the raw header for the first pilot sessions and set `TRUSTED_PROXY_HOPS` from what is observed. The `X-Client-Region: {client_region}` custom header is a valid load balancer variable.

**49. Points checked and found sound.**
- Cookie: HttpOnly, Secure, SameSite=Strict, path `/api/v1`, same origin as the SPA behind one load balancer: works, and the custom `X-Study-Client` header forces a CORS preflight that a cross-site page cannot pass. Sound (`POST /sessions` is unauthenticated anyway).
- Cloud Run ingress "internal and Cloud Load Balancing" plus IAP JWT check on admin routes: correct. IAP cannot protect a backend bucket: correct, so the separate admin app is justified.
- Allocation transaction: one Firestore transaction over the counter document is correct and well within the per-document write rate at 40 participants. The tie rule gives strict alternation, which is round robin. The balance problem is when allocation happens (findings 11, F item 1.6), not the transaction.
- Costs: forwarding rule about USD 18 per month; one idle Cloud Run instance at 1 vCPU and 512 MiB with request-based billing about USD 10; TTS a few cents to a dollar (more if Studio voices are chosen, still small). The USD 20 to 32 range is plausible.
- Severity: **note** (no action).

## E. Content quality

**50. The answers fit a challenge better than a routine question.**
- File: Content section 4 "Answer design"; JSON `answers.*` and `routineQuestion.text`.
- Draft says: every answer opens with a claim ("This design is the right one for us") and ends "We have looked at this carefully, and it stays as planned."
- Problem: after the chair's "could you say a little more about how the warehouses connect?" the denial does not answer the question and the closing line defends against a challenge nobody made. After "Challenge" it fits well. So the same denial sounds more evasive in the neutral group and after "explain" than after "challenge". That difference in fit is a confound for H4b (amplification) and for comparisons between actions.
- Severity: **should fix**.
- Fix: make all three questions ask for a justification ("Daniel, why did you choose this design?", "Could you explain why...", the challenge), so the reason answers all three and the denial fails all three alike. Make the closing line neutral: "That is the plan I recommend."

**51. Flags: plausibility and fit.**
- Slide 2 ("Architecture likely wrong"; "common practice is to connect each site directly to the cloud"): sounds like a VLM applying a branch-office rule. Wrong for this company for the reason stated (systems run in the head-office data centre). Risk: the slide shows a cloud labelled "Internet", so a lay reader may read the flag as "connect each warehouse straight to the internet", which the narration has just framed as unsafe. That makes the flag look silly rather than plausible. Fix: "...connect each site straight to cloud services".
- Slide 3 (price): plausible to a layperson and typical of VLM price comparisons; reason stated in two sentences. "overpriced" sounds closer to a judgement of the buyer than Oliver's "too expensive"; see finding 8.
- Slide 4 ("Numbers do not add up"): the arithmetic error is typical of a VLM, but the chart refutes it at a glance; see finding 3.
- Severity: **should fix** (slide 2 wording); slides 3 and 4 covered elsewhere.

**52. Reason length is uneven.**
- File: JSON `answers.reason` for s2, s3 (two middle sentences) and s4 (three).
- Severity: **note**. The task rule is one or two sentences.
- Fix: after finding 5, slide 4 becomes "Phase three starts while phase two is still running, with a different team. That is why the bars overlap." Two sentences.

**53. Denials are structurally parallel. (check passed)**
- All three denials are exactly the reasoned answer minus its middle; first and last sentences are identical. The only remaining difference is length (about 7 s versus 21 s), which Content point 7 already flags. Severity: **note**. Fix: none beyond finding 6 (acoustically identical shared sentences).

**54. The benevolence item is ambiguous in plain English.**
- File: Content section 5; JSON `ratingItems[id=benevolence]`.
- Draft says: "How responsible is Daniel Hale?"
- Problem: lay readers will read "responsible" as "in charge of" (he is responsible for security), not as the METI sense of acting responsibly toward others. Mayer's benevolence is "wants to do good" for the trustor.
- Severity: **should fix**.
- Fix: "How much does Daniel Hale care about the company's interests?" (anchors "Not at all" to "Very much"), or METI's "considerate". Check against the METI item table (Content point 5).

**55. The CISO's tone is constant. (check passed, one remark)**
- Narrations are calm and confident throughout; no line is defensive except the closing answer sentence (finding 50). Slide 4 ("If we wait another quarter, we miss it") and slide 5 ("I am confident") are firm but within a normal pitch. Severity: **note**.

**56. Technical level is suitable for lay participants, with one exception.**
- Firewalls are explained ("a locked door with a guard"); monitoring and remote access are in plain words. The exception is "route through head office" and "connect each site directly to the cloud" in flag 2 (see finding 51). Severity: **note**.

**57. The opening establishes the CISO's experience well. (check passed)**
- Six years in role, twelve years before at a bank and a hospital group, two attacks stopped, certifications; spoken and shown as a card before T0. Severity: **note**.

**58. Role instruction has one inaccurate sentence.**
- File: Content section 1 "Role instruction"; JSON `opening.introText[2]`.
- Draft says: "After some slides you will rate the presenter."
- Problem: the first rating comes after the opening, before any slide.
- Severity: **note**.
- Fix: "During the meeting you will rate the presenter several times, and at the end you will vote on the budget."

**59. Budget screen does not restate the stakes.**
- File: Content section 7; JSON `budget.prompt`.
- Severity: **note**. The Abstract asks that the amount and what depends on it be shown on the slide, which slides 1 and 5 do. Restating the stakes on the vote screen keeps the decision anchored for both groups.
- Fix: add "Our insurance renewal and our largest customer's contract depend on this protection."

**60. Background intro hints at groups.**
- File: Content section 8 intro; JSON `background.intro`.
- Draft says: "They help us compare groups."
- Severity: **note**. In the checker group the flag review comes after this screen, and the hint could colour it.
- Fix: "They help us describe who took part."

**61. Possible priming about AI or honesty. (check passed, two remarks)**
- The neutral group never hears of AI before the debrief. The checker introduction says "AI tool" and gives no accuracy, as intended. The AI-use item comes after the budget. Remarks: the integrity item "How honest is Daniel Hale?" is asked after every flag; this is the measure, not a flag, and is consistent with the Abstract, but it is the one place honesty is put in front of participants, so keep its wording neutral. The challenge question "My checker flags this slide" names the checker but not AI, which matches the Abstract's note on logging the chosen question. Severity: **note**.

**62. Length exceeds the Abstract's "about 10 minutes".**
- File: Content section 12 and point 2.
- Severity: **note**. The content's own estimate (13 to 15 minutes) leaves out a 10 s view after each answer (finding 20) and the checker introduction, so the real figure is higher.
- Fix: decide with Oliver (F, item 1.7). Cutting each narration to about 110 words saves about 1.5 minutes.

**63. Neutral debrief and group disclosure.**
- File: Content section 10.
- Severity: **note**. Telling each participant their group goes beyond the Abstract's three debrief points but is good practice. Keep.

**64. Slide image and chart are descriptions, not assets.**
- File: JSON `s2.img1`, `s4.chart1`.
- Severity: **note**. The VLM drafting run (finding 1) needs the rendered slides, so these assets come first.

**65. Profile card and narration repeat the attacks; fine.**
- Severity: **note**. "Two attempted break-ins" appears in the profile, slide 1 bullet and narration. It reinforces competence before T0 in both groups equally. No change.

**66. Currency and voice locale do not match.**
- File: JSON `voices.*.languageCode: "en-US"`, `scenario.budgetRequest.currency: "CAD"`, company in Vancouver.
- Severity: **note**.
- Fix: consider `en-CA` voices if available in the chosen family, or keep `en-US` and accept it; part of OPEN L.

## F. Merged list of open points

Sources: Brief OPEN A to U, Arch FO-1 to FO-16, Content points 1 to 12 (written C-1 to C-12 here), and new points from this review (written R-n, with the finding number). Where two lists ask the same question, both labels are given and the overlap is stated.

### 1. Study-design decisions the owner must make

1.1. **How the checker is introduced** [OPEN A]. Example scene with one sentence (Brief default) or a separate text screen (Content's choice, finding 9). What the panel shows during the opening: idle or "No issues found" [OPEN A, Arch `checkerPanelOnOpening`].
1.2. **When the checker is introduced: before or after the baseline rating** [R-1, finding 2]. New.
1.3. **Ratings after slides 1 and 5** [OPEN B]. 9 or 7 rating points.
1.4. **How the participant's question is presented** [OPEN D]. Text bubble or voiced.
1.5. **Budget decision format** [OPEN E].
1.6. **When allocation happens and whether abandoned sessions free their slot** [FO-7, FO-4, R-2 / finding 11]. FO-4 and FO-7 are the same question from two sides (balance of completed groups); this review adds the device gate.
1.7. **Total length versus "about 10 minutes"** [C-2, R-3 / finding 62].
1.8. **View time after an answer and after the opening** [FO-6 = C-4]. Same question in both lists, with opposite defaults (10 s versus 0 s).
1.9. **Background items and their position** [OPEN F]. Oliver's four at the start or all at the end.
1.10. **Routine question: slide-specific or generic, and how answers fit each question type** [R-4 / findings 12, 50]. New.
1.11. **Scenario facts** [R-5 / finding 13; includes C-9 CISO gender and C-12 currency]. Company, sector, amount, stakes, deadline, CISO and chair names, experience level.
1.12. **Profile delivery and the CISO's greeting before T0** [R-6 / finding 14; overlaps C-1 role instruction]. Spoken by the chair with a card, or text only.
1.13. **Rating scales** [OPEN I, C-5, R-7 / finding 54]. 0 to 10 for all four or 7-point METI items; the METI wording of the three items, including the benevolence wording. C-5 is the wording half of OPEN I.
1.14. **Moderator scales** [R-8 / finding 15]. Cybersecurity knowledge 1 to 5 or 0 to 10; AI use categories.
1.15. **End-of-study flag review** [OPEN G]. Include or not; binary or scale (finding 25).
1.16. **Checker display: overlay "AR glasses" or side panel with camera image** [FO-2, R-9 / finding 10]. Content's intro text already assumes glasses.
1.17. **How the flags are produced: real VLM run and provenance** [R-10 / finding 1]. New; needed for the debrief to be true.
1.18. **Panel behaviour on unflagged slides and before a flag** [C-3]. When "No issues found" appears; empty or "Checking..." before a flag. Related to 1.1.
1.19. **Round-robin scope: global or per run** [FO-1].
1.20. **Exclusion rules** [FO-9]. Which flags exclude a session; write into the analysis plan.
1.21. **Repeat participation** [FO-14].
1.22. **Answer length difference (reason about 21 s, denial about 7 s)** [C-7]. Accept and report, or pad.
1.23. **Fixed or shuffled order of the four rating items** [FO-16].
1.24. **Reading of "Slide 1-4" in the prompt** [OPEN T].

### 2. Application behaviour decisions

2.1. **Action step time limit and timeout handling** [OPEN C; finding 7 for the neutral group].
2.2. **Rating timeout: partial or discard** [OPEN H].
2.3. **Attention and bot checks** [OPEN U]. Audio word check and honeypot.
2.4. **Failed audio check: attempts, block or flag** [FO-5 = C-6]. Same question; FO-5 says two retries, Content says two attempts (finding 26).
2.5. **Reload during audio** [FO-3].
2.6. **Audio failure mid-study** [FO-12; finding 36].
2.7. **Captions on screen** [FO-13].
2.8. **Initial slider position** [FO-11].
2.9. **Devices allowed** [OPEN R].
2.10. **Scene visuals: static composite or video** [OPEN K].
2.11. **Elapsed timer start: link opened or consent** [FO-10].
2.12. **Budget and background screens timed or untimed** [R-11 / finding 16].
2.13. **Admin status definitions (Started versus Ongoing)** [R-12 / finding 17].
2.14. **Consent PDF, ethics text, data retention and withdrawal code** [OPEN M, FO-15, C-10]. FO-15 and C-10 both depend on OPEN M; C-10 ("Prefer not to say" and skipping) is the ethics form's call.
2.15. **Debrief request not to discuss the study** [C-11].
2.16. **Run-level config versioning** [OPEN S].

### 3. Infrastructure decisions

3.1. **IP address storage** [OPEN J, FO-8]. FO-8 (ASN) is a sub-question of OPEN J.
3.2. **Admin login: IAP or password** [OPEN N; finding 39 changes what IAP needs].
3.3. **Database: Firestore or Cloud SQL** [OPEN O].
3.4. **Static hosting: Cloud CDN with load balancer (about USD 18 per month) or Firebase Hosting** [OPEN P].
3.5. **GCP project, region, domain, DNS** [OPEN Q].
3.6. **Voices and language** [OPEN L, C-9 voice gender, finding 66 locale]. Listed here because it fixes the TTS setup; the gender part is in 1.11.

### 4. Small wording choices

4.1. **Firewall flag label** [C-8, R-13 / finding 8]. Oliver's wording or "overpriced"; there is no word limit in any source.
4.2. **"No issues found" or a supporting label** [from the v8 change file, noted in Content section 3].
4.3. **Role instruction screen and its wording** [C-1; finding 58].
4.4. **Currency display** [C-12]. "$" versus "CAD $".
4.5. **Slide 2 flag detail wording** [R-14 / finding 51].
4.6. **Background intro line** [R-15 / finding 60].

Count: group 1 has 24 points, group 2 has 16, group 3 has 6, group 4 has 6: **52 merged points** from 49 source items (21 + 16 + 12) plus 15 new review items, after removing duplicates.

## G. Verdict

1. **Architecture**: sound enough to start a vertical slice once Oliver answers the open points, provided the event sequence design (finding 33) and the lazy-timeout latency rule (finding 34) are fixed first; the Terraform IAP module (finding 39) must be redone before phase 4.
2. **Content**: not ready for a first TTS run. The JSON does not load into the architecture's types (finding 18), its `audioKey` would leak the denial (finding 19), and the flags have not been drafted with a real VLM (finding 1). After a schema-aligned v1.1 and the wording fixes in findings 4, 5, 29 and 50, a TTS run is reasonable.
3. **Three most important fixes**: (a) run the real VLM on rendered slides and keep provenance, or the debrief is untrue (finding 1); (b) align the JSON with the TypeScript types and rename `audioKey` (findings 18, 19); (c) fix the event sequence for server-written and rejected events and for reloads (finding 33).

### Severity count

- Must fix: 5 (findings 1, 18, 19, 33, 39).
- Should fix: 34 (findings 2 to 12, 20 to 29, 34 to 38, 40 to 44, 50, 51, 54).
- Note: 27 (findings 13 to 17, 30 to 32, 45 to 49, 52, 53, 55 to 66).
