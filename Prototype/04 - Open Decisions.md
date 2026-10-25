# Open decisions for the prototype

This file lists every open point that must be decided before the architecture and the study content can be finalised.
It merges the brief's OPEN A to U, the architecture's FO-1 to FO-16, the content draft's points (C-1 to C-12) and new points from `reports/crosscheck-v1.md` (R-1 to R-15). Finding numbers refer to that report.
Abstract v8 decides the study design. Your prompt of 2026-10-23 decides the application.
To answer, fill the "Answer:" line under each point. Write "default" to accept the recommendation. Nothing here is assumed until you answer.
Tags: `BLOCKS BUILD` means the vertical slice cannot start without it. `BLOCKS CONTENT` means the text-to-speech run or the slides cannot be finalised without it. `CAN WAIT` means a later answer changes only parameters or later phases.

Status 2026-10-24: all 53 points answered. Flags selected from the VLM run (see 4.1). Code location: the Prototype folder is the monorepo root described in 02 section 3; the study documents stay at its root. Build of the vertical slice started 2026-10-24.

## 1. Study-design decisions

### 1.1 How the checker is introduced [OPEN A] `BLOCKS CONTENT`
**Question:** How does the checker group learn about the checker, and what does the panel show during the opening?
**Options:**
- A: during the example scene, the panel appears and one spoken or written sentence explains it (brief default).
- B: a separate text screen after the example scene (content draft).
- Panel during the opening: idle (title only) or "No issues found".
**Recommendation:** A, with an idle panel during the opening. A is the brief's default, and the content draft used B while calling it the default (finding 9).
**Changes if decided later:** the example scene's script and audio, the checker introduction text, and the `checkerIntro` option in the config.
**Answer:** Decided (Oliver, 2026-10-23): no separate checker screen. The checker group reads the same role instruction as the neutral group plus one neutral sentence, for example "A checker on your tablet reads each slide and marks anything it finds." The panel shows "No issues found" on unflagged slides as Abstract v8 says. The challenge option wording stays. Follow-up 2026-10-24: with the role screen dropped (4.3), the sentence sits on the last setup screen ("The meeting is about to start"), shown to the checker group only; the neutral group sees the same screen without it.

### 1.2 Checker introduced before or after the baseline rating [R-1] `CAN WAIT`
**Question:** Does the checker group learn about the checker before the first trust rating (T0) or after it?
**Options:**
- Before T0, during setup (both drafts and the brief's flow table).
- After T0, just before slide 1.
**Recommendation:** after T0. Abstract v8 step 1 lists "First trust rating. Checker group: the checker is introduced" in that order, and its Validity section checks that the groups start alike on the first rating (finding 2).
**Changes if decided later:** the step order in the plan compiler and one option in the config. No text changes.
**Answer:** Resolved by 1.1: the one sentence sits in the role instruction, which both groups read before the board intro and the baseline rating. No further introduction after T0 (Oliver, 2026-10-23).

### 1.3 Ratings after slides 1 and 5 [OPEN B] `CAN WAIT`
**Question:** Should participants rate the CISO after slide 1 and slide 5 as well, giving 9 ratings instead of 7?
**Options:**
- 9 ratings (T0, T1, T2a, T2b, T3a, T3b, T4a, T4b, T5).
- 7 ratings (no T1 and T5).
**Recommendation:** 9 ratings. Abstract v8's Validity section needs slide 1 to show "whether the checker's presence alone moves trust", which needs a rating after slide 1. Your prompt also rates after every slide.
**Changes if decided later:** one option in the config (`ratingAfterUnflaggedSlides`). Study length changes by about 30 s.
**Answer:** Every presentation slide is followed by a rating, with a baseline rating after the board intro (Oliver, 2026-10-23). Exact count depends on the slide-count clarification under 1.24.

### 1.4 How the participant's question is presented [OPEN D] `BLOCKS CONTENT`
**Question:** When a checker participant asks or challenges, is the question shown as text or spoken by a voice?
**Options:**
- Text in a speech bubble for about 3.5 s, then the CISO answers (brief default).
- Spoken by a third "board member" voice.
**Recommendation:** text bubble. It keeps the participant's question as their own and needs no extra voice.
**Changes if decided later:** if voiced, six more audio lines and a third voice must be generated.
**Answer:** Spoken by a third board-member text-to-speech voice; the chosen question is voiced before the CISO answers. Adds one voice and six clips. (Oliver, 2026-10-24)

### 1.5 Budget decision format [OPEN E] `CAN WAIT`
**Question:** How does the participant vote on the budget?
**Options:**
- Approve full / approve a reduced amount / defer / reject, plus a 0 to 10 "how sure" item (brief default).
- Approve / reject only.
- Approve / reject plus the 0 to 10 item.
- One slider from $0 to $2.4 million.
**Recommendation:** the brief's default. It shows whether a change in trust reaches the decision even when nobody rejects outright.
**Changes if decided later:** the budget screen and the export columns only.
**Answer:** Approve or reject, plus "How sure are you?" 0 to 10. (Oliver, 2026-10-24)

### 1.6 When the group is assigned, and whether abandoned sessions free their slot [FO-7, FO-4, R-2] `BLOCKS BUILD`
**Question:** Is the group (and the denial slide) assigned when the link is opened, or later? Does a session that is abandoned give its slot to the next person? FO-4 and FO-7 are the same question from two sides: balance of completed groups.
**Options:**
- Assign when the link is opened (brief, step 0). Abandoned slots are not freed.
- Assign at consent, after the device check. Abandoned slots are not freed.
- Assign at consent. Slots of sessions abandoned before slide 1 are freed.
**Recommendation:** assign at consent, do not free slots. Abstract v8 wants groups "equal in size" through round robin, and assigning at link opening lets every blocked phone visitor and every decliner use up a slot (finding 11).
**Changes if decided later:** session creation and the allocation transaction, which are written in the first phase.
**Answer:** At consent, after the device check; abandoned slots are not freed. (Oliver, 2026-10-24)

### 1.7 Total length versus "about 10 minutes" [C-2, R-3] `BLOCKS CONTENT`
**Question:** The current script runs about 13 to 15 minutes or more against Abstract v8's "about 10 minutes". What should give?
**Options:**
- Shorten each narration to about 110 words.
- Drop the ratings after slides 1 and 5 (see 1.3).
- Keep the script and update the Abstract's estimate.
**Recommendation:** shorten the narrations. 1.3 keeps the slide 1 rating, which the Abstract needs, and the Abstract states the length to participants (finding 62).
**Changes if decided later:** all five narration scripts, their audio, and the marking times.
**Answer:** Shorten each CISO narration to about 110 words; keep every rating point. (Oliver, 2026-10-24)

### 1.8 View time after an answer and after the opening [FO-6, C-4] `CAN WAIT`
**Question:** Your prompt gives 10 s to look at the slide after the audio. Does that pause also apply after the opening and after each CISO answer, before the rating? The two drafts have opposite defaults (10 s and 0 s).
**Options:**
- 10 s everywhere a rating follows audio (architecture).
- 10 s after slides and the opening, 0 s after answers (content).
- A different value after answers.
**Recommendation:** 10 s everywhere, so every rating follows the same rhythm and the denial (shorter audio) is not also followed by a shorter pause (finding 20).
**Changes if decided later:** one timing parameter. Study length changes by up to 30 s.
**Answer:** 10 s everywhere a rating follows audio: after the opening, after each slide, after each answer. (Oliver, 2026-10-24)

### 1.9 Background questions: which and where [OPEN F] `CAN WAIT`
**Question:** Your prompt asks field, country, age and job level right after consent. Abstract v8 asks education, age, field, cybersecurity knowledge and AI use after the budget decision. Which items, and where?
**Options:**
- All seven items after the budget decision (brief default).
- Your four after consent, the other three at the end.
**Recommendation:** all after the budget decision. Abstract v8 places background questions there, and asking about cybersecurity or AI first could prime participants.
**Changes if decided later:** step order and one option in the config.
**Answer:** All seven at the end, after the budget decision (Oliver, 2026-10-23).

### 1.10 The chair's routine question and how the answers fit it [R-4] `BLOCKS CONTENT`
**Question:** Should the chair's routine question be specific to each slide (now: it points at the flagged item) or the same generic question each time? And should the answers be reworded so the denial fails every kind of question alike?
**Options:**
- Keep slide-specific "say a little more" questions and the current answers.
- Slide-specific questions that ask "why" ("Daniel, why did you choose this design?"), and a neutral closing line ("That is the plan I recommend.").
- One generic question for all three slides.
**Recommendation:** the second option. Abstract v8 calls it a "routine question" and wants the same answer whoever asks; at present the denial sounds more evasive after the chair's question than after a challenge, which confounds H4b (findings 12, 50).
**Changes if decided later:** three chair questions and six CISO answers, with their audio.
**Answer:** Slide-specific "why" questions with the neutral closing line, as in content v1.1 and v1.2. (Oliver, 2026-10-24)

### 1.11 Scenario facts [R-5, C-9, C-12] `BLOCKS CONTENT`
**Question:** Do you accept the scenario the content draft set (company, sector, amount, stakes, names, experience)? See the table at the end of this file.
**Options:**
- Keep as drafted.
- Change specific rows in the table.
**Recommendation:** keep as drafted unless a row conflicts with how you want the CISO to come across. Abstract v8 only fixes that the amount and stakes are identical and shown, and that the profile establishes experience, which the draft does.
**Changes if decided later:** slide texts, all narrations, the opening and the profile card.
**Answer:** Keep the scenario, but the company name must be neutral and not country-specific, and no country-specific details (no "Canadian", no named country). Rename Kestrel Bay Logistics and strip country references in content v1.3. (Oliver, 2026-10-24)

### 1.12 How the profile is delivered, and the CISO's greeting before T0 [R-6, C-1] `BLOCKS CONTENT`
**Question:** Is the CISO profile read aloud by the chair with a card on screen, or shown as text only? Does the CISO say one greeting line before the first rating?
**Options:**
- Chair reads it, card on screen, CISO greets once (content draft).
- Text card only, no greeting.
- Chair reads it, no CISO greeting.
**Recommendation:** the content draft. Every participant receives the profile before T0, and the baseline rating is of a person whose voice they have heard.
**Changes if decided later:** the opening script and its audio.
**Answer:** Chair reads the profile aloud with a card on screen; the CISO says one greeting line before T0. (Oliver, 2026-10-24)

### 1.13 Rating scales and item wording [OPEN I, C-5, R-7] `CAN WAIT`
**Question:** Are all four items on 0 to 10, or are competence, benevolence and integrity on 7-point METI scales? Which words are used for the three items?
**Options:**
- All four on 0 to 10 (brief default).
- Trust 0 to 10, the three items on 7 points.
- Benevolence worded "How responsible is Daniel Hale?" (draft) or "How much does Daniel Hale care about the company's interests?".
**Recommendation:** all four on 0 to 10, with the benevolence item reworded. Abstract v8 uses a 0-to-10 rating and items "worded after" METI; "responsible" reads as "in charge of" in plain English (finding 54). Check the wording against the METI item table.
**Changes if decided later:** rating labels and scales in the config. No audio.
**Answer:** Trust stays 0 to 10. The three items use METI's 7-point antonym format (for example incompetent 1 to 7 competent, dishonest 1 to 7 honest) with neutral wording; no question form such as "How honest is Daniel?", which can suggest he is not. (Oliver, 2026-10-24)

### 1.14 Scales for the two moderators [R-8] `CAN WAIT`
**Question:** How are self-rated cybersecurity knowledge and own AI use measured?
**Options:**
- Knowledge 1 to 5 (None to Expert); AI use six frequency steps (Never to Every day) (draft).
- Knowledge 0 to 10.
- Other wording.
**Recommendation:** the draft. Abstract v8 names both as planned moderators of H1 and H2 but not their scales; short labelled scales are quick for lay participants.
**Changes if decided later:** two items on the background screen.
**Answer:** Cybersecurity knowledge 1 to 5 (None, Basic, Some, Good, Expert); own AI use six steps (Never to Every day). (Oliver, 2026-10-24)

### 1.15 End-of-study flag review [OPEN G] `CAN WAIT`
**Question:** Should checker participants rate each flag before the debrief, and how?
**Options:**
- Include, checker group only, Correct / Incorrect (brief default).
- Include with Correct / Incorrect / Not sure.
- Include as a scale of how true each flag was (closer to Prike et al.).
- Leave out.
**Recommendation:** include, checker group only, with a short scale. Abstract v8's appendix borrows Prike's end-of-study rating of how true each checked claim is, which is a scale (finding 25).
**Changes if decided later:** one screen and its export columns.
**Answer:** Leave out. No flag review screen; the debrief follows the background questions. Remove the flagReview scene from the content. (Oliver, 2026-10-24)

### 1.16 Checker display: on the slide or in a side panel [FO-2, R-9] `BLOCKS BUILD`
**Question:** Does the checker mark the slide itself, as if through AR glasses, or show the camera image of the slide in a panel beside it?
**Options:**
- Overlay on the slide inside a glasses or phone-like frame (architecture default, your prompt).
- Side panel showing the camera image with the marking (Abstract v8 wording).
**Recommendation:** overlay. Your prompt decides the application's look and asks for it; the Abstract's "panel shows the camera image" is met by framing the overlay as the camera view. The checker introduction text must then match (finding 10).
**Changes if decided later:** the boardroom scene component built in the first phase, and the checker introduction text.
**Answer:** The checker shows as highlighting on the slide while the CISO talks; the neutral group sees the identical simulation without highlighting (Oliver, 2026-10-23).

### 1.17 Producing the flags with a real VLM [R-10] `BLOCKS CONTENT`
**Question:** The current flags were written by hand. Abstract v8 says they were "drafted with a real VLM from images of the slides". How are they produced?
**Options:**
- Render the five slides, run a real VLM with a fixed prompt, keep model, date, prompt and raw output, then select and edit flags from it.
- Keep the hand-written flags and change the Abstract and debrief wording.
**Recommendation:** the real VLM run. Abstract v8 settles this, and without it the debrief sentence "drafted with a real AI tool" would be untrue (finding 1).
**Changes if decided later:** flag labels, details, the challenge questions and the CISO's reasons may all change.
**Answer:** Real VLM run, using a Claude Opus agent as the vision-language model: render the slides to images, one fixed prompt per slide, store model, date, prompt and raw output under content/vlm-drafts/, then select and edit flags and log the edits. (Oliver, 2026-10-24)

### 1.18 Panel on unflagged slides and before a flag [C-3] `CAN WAIT`
**Question:** When does "No issues found" appear on slides 1 and 5, and what does the panel show before it or before a flag appears?
**Options:**
- Label appears at 8 s; before that the panel shows its title only (draft).
- A "Checking..." state before the label or flag.
- Label shown from the start.
**Recommendation:** the draft. It treats flagged and unflagged slides alike until the checker "speaks".
**Changes if decided later:** two timing values and one label in the config.
**Answer:** Panel is present from the first CISO slide; "No issues found" on unflagged slides (Oliver, 2026-10-23). Follow-up decided by Claude on Oliver's instruction (2026-10-24): before any marking the panel shows only its title, nothing else; on unflagged slides the "No issues found" label appears at the configured mark, default 8 s into the narration.

### 1.19 Round robin across all links or per link [FO-1] `CAN WAIT`
**Question:** Does group assignment alternate across all recruitment links together, or separately per link?
**Options:**
- Global (architecture default).
- Per run.
**Recommendation:** global. Abstract v8 wants the groups equal in size overall.
**Changes if decided later:** one option in the config.
**Answer:** Across all links (global round robin). (Oliver, 2026-10-24)

### 1.20 Exclusion rules [FO-9] `CAN WAIT`
**Question:** Which flags (duplicate, honeypot, wrong audio word, many rating timeouts, tab hidden during audio) exclude a session from analysis?
**Options:**
- The system only flags; you write the rules into the analysis plan before data collection (architecture default).
- Some flags end the session automatically.
**Recommendation:** only flag, rules written in advance.
**Changes if decided later:** nothing in the application.
**Answer:** Flag only; exclusion rules go into the analysis plan before data collection. (Oliver, 2026-10-24)

### 1.21 Repeat participation [FO-14] `CAN WAIT`
**Question:** May one person take the study more than once?
**Options:**
- No; asked in the welcome text, duplicates flagged (architecture default).
- Yes.
**Recommendation:** no. A second run would expose the design and break the group comparison.
**Changes if decided later:** one sentence in the welcome text.
**Answer:** Record IP address and browser and mark duplicates in the admin view; do not block. (Oliver, 2026-10-24)

### 1.22 Reason longer than denial [C-7] `BLOCKS CONTENT`
**Question:** A reasoned answer is about 21 s, the denial about 7 s. Accept the difference or pad the denial?
**Options:**
- Accept and report it as a limitation.
- Add silence after the denial so the times match.
**Recommendation:** accept. Abstract v8 says "only the presence of a reason differs", and length follows from the reason; padding with silence would itself sound odd.
**Changes if decided later:** the three denial clips.
**Answer:** First try to write each denial at the same length as its reasoned answer while still carrying no reason; if that cannot be made credible, pad the denial with silence so the time before the rating matches. (Oliver, 2026-10-24)

### 1.23 Order of the four rating items [FO-16] `CAN WAIT`
**Question:** Are the four items always in the same order, or shuffled?
**Options:**
- Fixed, trust first (architecture default).
- Shuffled per participant.
**Recommendation:** fixed, so participants answer quickly, as Abstract v8 wants.
**Changes if decided later:** one option in the rating screen.
**Answer:** Fixed order, trust first. (Oliver, 2026-10-24)

### 1.24 Reading of "Slide 1-4" in your prompt [OPEN T] `BLOCKS CONTENT`
**Question:** The brief reads your "Slide 1-4 with the Questions ahead of the start" as: opening scene, then five CISO slides, and the "questions ahead" are the baseline rating. Is that right?
**Options:**
- Yes, as the brief reads it.
- Something else (please describe).
**Recommendation:** yes. Abstract v8 settles five slides with flags on 2 to 4.
**Changes if decided later:** the number and order of slides.
**Answer:** Does not matter. Keep Abstract v8 (intro, then five CISO slides) as the default, but the number of slides, which slides carry a flag and the questions must be configurable so slides can be added after testing (Oliver, 2026-10-23).

### 1.25 Does the participant need to know the checker is AI? [C-22, new 2026-10-23] `BLOCKS CONTENT`
**Question:** The decided checker sentence ("During this meeting, a checker reads each slide and marks anything it finds.") never says the checker is an AI or a vision-language model. Abstract v8's research question, the debrief and the own-AI-use moderator all assume participants experience the flags as coming from an AI.
**Options:**
- Keep the sentence as decided; participants learn it was an AI at the debrief only.
- Add two words: "an AI checker" (or "an automated checker").
- Keep the sentence, and ask at the end of the study what participants thought the checker was (a manipulation check).
**Recommendation:** "an automated checker" plus the end-of-study manipulation check. The hypotheses rest on an AI's contradiction (Kempt and Nagel 2022; Chen et al. 2026), so participants should know the source is a machine without being told it may be wrong.
**Changes if decided later:** one sentence in the role instruction and one optional end-of-study item; no code change.
**Answer:** Keep "a checker" as decided; participants learn it was an AI at the debrief. (Oliver, 2026-10-24)

## 2. Application behaviour decisions

### 2.1 Time limit on the action choice [OPEN C] `CAN WAIT`
**Question:** How long do participants have to choose an action (or say whether they have a question), and what happens on timeout?
**Options:**
- 15 s; on timeout the chair asks (brief default). Neutral timeouts recorded as missing, not "No".
- Untimed.
**Recommendation:** 15 s, with timeouts stored as missing in both groups so H4a is not biased (finding 7).
**Changes if decided later:** one timing parameter.
**Answer:** 15 s; on timeout the chair asks the routine question; the timeout is stored as missing in both groups. (Oliver, 2026-10-24)

### 2.2 Rating timeout [OPEN H] `CAN WAIT`
**Question:** If 15 s pass before all four items are answered, keep what was given or discard the rating?
**Options:**
- Keep partial, mark the rest missing (brief default).
- Discard the whole rating.
**Recommendation:** keep partial; nothing is lost and the analysis can still drop it.
**Changes if decided later:** one option.
**Answer:** Keep partial answers, mark the missing items as timed out. (Oliver, 2026-10-24)

### 2.3 Attention and bot checks [OPEN U] `CAN WAIT`
**Question:** Beyond IP and browser, include an audio word check and a hidden honeypot field?
**Options:**
- Both (brief default).
- Audio check only.
- Neither.
**Recommendation:** both; they cost little and help spot careless sessions.
**Changes if decided later:** the setup screen and one hidden field.
**Answer:** Both the audio word check and the honeypot field; failures flag, never block. (Oliver, 2026-10-24)

### 2.4 Failed audio check [FO-5, C-6] `CAN WAIT`
**Question:** How many attempts, and does a failure block or only flag? The drafts disagree: two retries (three attempts) versus two attempts (finding 26).
**Options:**
- Two attempts, then continue with a flag.
- Three attempts, then continue with a flag.
- Block after the last attempt.
**Recommendation:** two attempts, then continue with a flag.
**Changes if decided later:** one parameter.
**Answer:** One answer attempt. The test clip in the setup step may be replayed as often as the participant likes (the only replayable clip). The answer, right or wrong, is recorded and shown on the session record in the admin view; never block. (Oliver, 2026-10-24)

### 2.5 Reload during audio [FO-3] `CAN WAIT`
**Question:** If a participant reloads mid-clip, what happens, and is the session kept?
**Options:**
- Skip to the silent view of the slide, flag "reloaded", keep (architecture default).
- Replay the clip.
- End the session.
**Recommendation:** the default; your prompt forbids replay.
**Changes if decided later:** resume logic in phase 2.
**Answer:** Replay the interrupted clip from the start after a reload; log the reload on the session record. (Oliver, 2026-10-24)

### 2.6 Audio failure mid-study [FO-12] `CAN WAIT`
**Question:** If a clip fails or stalls, may the participant continue?
**Options:**
- Show a message, let them continue after the clip's length, flag the session (architecture default, adjusted per finding 36).
- End the session.
**Recommendation:** continue after the clip's length, flagged.
**Changes if decided later:** client and server handling in phase 2.
**Answer:** Continue after the clip length has passed; session flagged "audio problem" and kept. (Oliver, 2026-10-24)

### 2.7 Captions [FO-13] `CAN WAIT`
**Question:** Should the spoken words appear as captions?
**Options:**
- No captions; transcript for screen readers only (architecture default).
- Captions on screen.
**Recommendation:** no captions, so the experience stays a spoken briefing.
**Changes if decided later:** one display option.
**Answer:** No captions; transcript for screen readers only. (Oliver, 2026-10-24)

### 2.8 Starting position of the rating sliders [FO-11] `BLOCKS BUILD`
**Question:** Do sliders start with no handle (untouched is recorded as missing) or at a value such as 5?
**Options:**
- No handle until clicked (architecture default).
- Start at 5.
**Recommendation:** no handle; a slider at 5 makes an untouched item look like an answer.
**Changes if decided later:** the rating component, built in the first phase.
**Answer:** No default position at all. The 0 to 10 scale is shown as clickable numbers with nothing preselected, so no handle or starting value can taint the choice; untouched is recorded as missing. (Oliver, 2026-10-24)

### 2.9 Devices allowed [OPEN R] `CAN WAIT`
**Question:** Desktop only, or also tablets and phones?
**Options:**
- Desktop and landscape tablet, at least 1,024 px wide; phones see a "please use a larger screen" page (brief default).
- Phones allowed.
**Recommendation:** the default; the boardroom scene with slide text needs space.
**Changes if decided later:** device gate values.
**Answer:** Desktop and landscape tablet, at least 1,024 px wide; phones see a larger-screen page. (Oliver, 2026-10-24)

### 2.10 Scene visuals [OPEN K] `BLOCKS BUILD`
**Question:** Is a still boardroom picture with a CISO portrait that lights up while speaking enough, or do you want video (for example an avatar)?
**Options:**
- Still composite with light animation (brief default).
- Generated video.
**Recommendation:** still composite now; video can be added later.
**Changes if decided later:** the scene component built in the first phase, and the asset pipeline.
**Answer:** Still composite with light animation (Oliver, 2026-10-23).

### 2.11 Start of the elapsed timer in the admin view [FO-10] `CAN WAIT`
**Question:** Does the admin timer start when the link is opened or at consent?
**Options:**
- At link opening, consent time also exported (architecture default).
- At consent.
**Recommendation:** the default.
**Changes if decided later:** one admin calculation.
**Answer:** Elapsed timer starts at consent. Link-opening time is still stored and exported. (Oliver, 2026-10-24)

### 2.12 Untimed screens [R-11] `CAN WAIT`
**Question:** Should the budget vote and background questions have a time limit?
**Options:**
- Untimed (draft).
- Timed, with a value you set.
**Recommendation:** untimed. Your prompt times the view and the ratings only.
**Changes if decided later:** one parameter per screen.
**Answer:** Timed: 30 s each for the budget vote and the background questions screen. On timeout keep partial answers, mark the rest missing, move on, as for ratings. (Oliver, 2026-10-24)
**Revised:** The background questions screen has no time limit; 30 s is not enough to fill in the information. The budget vote keeps its 30 s. Config: `timing.backgroundWindowMs: 0` (absent means `formWindowMs`, as before). The abandon sweep (30 min without an event) still ends a session left open on that screen; its ratings and budget decision are complete by then and usable. (Oliver, 2026-10-25)

### 2.13 Admin status meanings [R-12] `CAN WAIT`
**Question:** Your prompt lists Started, Ongoing, Finished. The architecture defines Started as "link opened, no slide audio yet", Ongoing as "after the first audio", plus Abandoned and Blocked. Is that what you want?
**Options:**
- As defined.
- Started = consent given; Ongoing = after the first slide.
**Recommendation:** as defined.
**Changes if decided later:** the status mapping in the admin view.
**Answer:** Started begins when the audio check is confirmed (before that a session is Opened). Ongoing after the first scene audio, Finished, Abandoned after 30 min without events, Blocked by the device gate, as proposed. (Oliver, 2026-10-24)

### 2.14 Consent, ethics text, data keeping and withdrawal [OPEN M, FO-15, C-10] `CAN WAIT`
**Question:** Who supplies the consent PDF and ethics wording? How long is data kept? Can participants withdraw after finishing? Should every background item allow "Prefer not to say"? FO-15 and C-10 both depend on the ethics form.
**Options:**
- Placeholder now; SFU ethics form later; show a session code at the debrief for withdrawal; "Prefer not to say" on every item (drafts).
- Other, per the ethics approval.
**Recommendation:** the drafts' placeholders until the ethics form arrives.
**Changes if decided later:** the consent PDF, debrief end text and background options. Blocks launch, not the build.
**Answer:** Placeholders until the SFU ethics form arrives: placeholder consent PDF, session code on the debrief for withdrawal, "Prefer not to say" on every background item, retention per the ethics approval. (Oliver, 2026-10-24)

### 2.15 Ask participants not to discuss the study [C-11] `CAN WAIT`
**Question:** Should the debrief ask participants not to tell others what the study is about?
**Options:**
- Yes, one sentence.
- No.
**Recommendation:** yes; recruitment runs through personal networks where people talk.
**Changes if decided later:** one sentence in the debrief.
**Answer:** Yes, one sentence asking participants not to discuss the study while recruitment runs. (Oliver, 2026-10-24)

### 2.16 Content versions per link [OPEN S] `CAN WAIT`
**Question:** Each link points at one fixed version of the content, so results stay comparable. Agree?
**Options:**
- Yes, versioned config file, locked once used (brief default).
- No versioning.
**Recommendation:** yes.
**Changes if decided later:** config loading and the admin config tab.
**Answer:** Yes: each link is pinned to one content version, locked once used. (Oliver, 2026-10-24)

## 3. Infrastructure decisions

### 3.1 How the IP address is stored [OPEN J, FO-8] `BLOCKS BUILD`
**Question:** Store the raw IP address, or only a keyed hash plus country? FO-8 asks whether the network provider (ASN) is also needed.
**Options:**
- Keyed hash and country; raw IP only in logs for 7 days; no ASN (brief and architecture default).
- Raw IP stored (needs ethics wording).
- Add ASN through a separate database.
**Recommendation:** hash and country, no ASN. It still detects duplicates and keeps less personal data.
**Changes if decided later:** session creation code written in the first phase, and the ethics text.
**Answer:** Store the raw IP address. Needs a line in the consent and ethics text. (Oliver, 2026-10-24)

### 3.2 Admin login [OPEN N] `CAN WAIT`
**Question:** Google login through Identity-Aware Proxy with a list of allowed accounts, or a shared password?
**Options:**
- IAP with an allowlist (brief default).
- Shared admin password.
**Recommendation:** IAP. Note that the architecture's Terraform for IAP must be redone either way (finding 39), and the Google-managed login only admits accounts in one organisation; say which Google accounts need access.
**Changes if decided later:** the Terraform module and admin auth check in phase 4.
**Answer:** Shared admin password, stored in Secret Manager and checked by the Go service; no IAP. (Oliver, 2026-10-24)

### 3.3 Database [OPEN O] `BLOCKS BUILD`
**Question:** Firestore or Cloud SQL Postgres?
**Options:**
- Firestore (brief default; free at this size).
- Cloud SQL Postgres (easier SQL export; about CAD 15 per month minimum).
**Recommendation:** Firestore.
**Changes if decided later:** the storage layer, which the first phase writes.
**Answer:** Firestore. (Oliver, 2026-10-24)

### 3.4 Static hosting and CDN [OPEN P] `CAN WAIT`
**Question:** Your prompt asks for Google Cloud CDN, which needs a load balancer at about USD 18 per month. Accept the cost, or use Firebase Hosting or Cloud Run instead?
**Options:**
- Cloud CDN with the load balancer (your prompt).
- Firebase Hosting.
- Serve everything from Cloud Run.
**Recommendation:** Cloud CDN, as your prompt asks; the load balancer also carries IAP and the country header.
**Changes if decided later:** Terraform in phase 4.
**Answer:** Everything from Cloud Run: the Go service serves the app, the slide assets and the audio. No Cloud CDN, no load balancer. Country comes from a GeoIP lookup or is dropped; raw IP is stored anyway (3.1). (Oliver, 2026-10-24)

### 3.5 GCP project, region, domain [OPEN Q] `CAN WAIT`
**Question:** Which GCP project, region and domain, and who runs DNS?
**Options:**
- Region northamerica-northeast1 (Montreal) or us-west1.
- Domain and DNS provider: please name them.
**Recommendation:** northamerica-northeast1, to keep data in Canada. The Firestore location cannot be changed later.
**Changes if decided later:** Terraform variables; must be set before the first cloud deploy.
**Answer:** Region: the one closest to Vancouver, as a Terraform variable so it can change. Firestore location is fixed once created, so a change later means a new database. Default region northamerica-northeast1 (Montreal), as a Terraform variable. GCP project id and domain still to be named. (Oliver, 2026-10-24)

### 3.6 Voices and language [OPEN L, C-9] `BLOCKS CONTENT`
**Question:** Which Google text-to-speech voices for the CISO (male, per the Abstract's "himself") and the chair (clearly different), and which accent?
**Options:**
- English (US) voices of one family (draft).
- English (Canada) or (UK) voices if the family offers them.
**Recommendation:** pick by listening to two or three candidates; prefer a family that supports timing marks, so flags can be timed to words (finding 22).
**Changes if decided later:** all audio is regenerated.
**Answer:** US English. Shortlist two or three Google voices per role (CISO male, chair clearly different, board member for the voiced question) from a family with word timing marks, generate samples, Oliver picks by listening. (Oliver, 2026-10-24)

## 4. Small wording choices

### 4.1 Firewall flag label [C-8, R-13] `BLOCKS CONTENT`
**Question:** Use your wording "Firewalls likely too expensive, cheaper ones available" or the draft's "Firewalls likely overpriced, cheaper ones available"?
**Options:**
- Your wording.
- The draft's wording.
**Recommendation:** your wording. No source sets the 6-word limit the draft cited (finding 8). The VLM run (1.17) may change it anyway.
**Changes if decided later:** the flag label, the challenge question and the flag review.
**Answer:** The label wording is secondary. The deck must tell the CISO's case for funding in stages: what is going on, what is the risk and impact, how we solve it, how much it costs, then the decision. Restructure the five slides to that arc in content v1.3; flag labels come from the VLM run (1.17). (Oliver, 2026-10-24) Follow-up 2026-10-24: Oliver delegated the selection to Claude. Chosen from the run: slide 2 "$8 million does not follow from shown figures", slide 3 "Head office is a single point of failure", slide 4 "Firewall cost per site looks high"; see content/vlm-drafts/v1.3/selection.md.

### 4.2 "No issues found" or a supporting label [v8 change file] `CAN WAIT`
**Question:** Keep "No issues found" on slides 1 and 5, or use a supporting label?
**Options:**
- "No issues found" (Abstract v8).
- A supporting label (please give wording).
**Recommendation:** "No issues found"; Abstract v8 uses it.
**Changes if decided later:** one label in the config.
**Answer:** "No issues found". (Oliver, 2026-10-24)

### 4.3 Role instruction screen [C-1] `CAN WAIT`
**Question:** Keep the three-line screen telling participants they sit on the board and will vote on the budget?
**Options:**
- Keep, with the second sentence fixed to "During the meeting you will rate the presenter several times" (finding 58).
- Drop it.
**Recommendation:** keep, fixed; Abstract v8 says the participant sits on the board and decides the budget, and this tells them.
**Changes if decided later:** one text screen.
**Answer:** Drop the role instruction screen. The checker sentence from 1.1 needs another place; see 1.1 follow-up. (Oliver, 2026-10-24)

### 4.4 Currency display [C-12] `BLOCKS CONTENT`
**Question:** Show amounts as "$" or as "CAD $"?
**Options:**
- "$" (draft).
- "CAD $" on the slides.
**Recommendation:** "$"; the company is Canadian and the exact currency does not affect trust.
**Changes if decided later:** slide texts.
**Answer:** Plain "$". (Oliver, 2026-10-24)

### 4.5 Slide 2 flag detail [R-14] `BLOCKS CONTENT`
**Question:** The detail says "connect each site directly to the cloud". The slide shows a cloud labelled Internet, so readers may read it as "connect to the internet". Reword?
**Options:**
- Keep.
- "...connect each site straight to cloud services."
**Recommendation:** reword, so the flag stays plausible (finding 51).
**Changes if decided later:** one flag detail.
**Answer:** Resolved by the rewrite: build the whole scenario so that slide, flag and answer match; details can change later. (Oliver, 2026-10-24)

### 4.6 Background screen intro [R-15] `CAN WAIT`
**Question:** The intro says "They help us compare groups", which hints that groups exist. Reword?
**Options:**
- Keep.
- "They help us describe who took part."
**Recommendation:** reword (finding 60).
**Changes if decided later:** one sentence.
**Answer:** Reword to state plainly that, as part of the study, the following data is collected; no mention of groups. (Oliver, 2026-10-24)

## Facts set by the drafts that Oliver should confirm or change

| Fact | Current value (source) | Keep / change to |
|---|---|---|
| Company name | Kestrel Bay Logistics, fictional (Content 1) | |
| Sector and location | Freight and warehousing, western Canada, head office and data centre near Vancouver (Content 1) | |
| Size | 1,400 staff, about $480 million revenue a year (Content 1) | |
| Number of sites | 13: 12 warehouses and head office (Content 1, slide 2) | |
| Budget amount | $2.4 million (Content 1, slides 1, 3, 5) | |
| Cost breakdown | Firewalls $780,000; monitoring 2 years $960,000; remote access $310,000; installation and training $350,000 (slide 3) | |
| What depends on it | Cyber insurance renewal, and the largest customer's new contract (slide 1) | |
| Deadline | Insurance renewal in 10 months; work ready in 8 months (slide 4) | |
| Phase timing | Phase 1 months 1 to 2, Phase 2 months 3 to 7, Phase 3 months 5 to 8 (slide 4) | |
| Cost of one stopped day | About $1.2 million in lost shipments (slide 4) | |
| CISO name and pronoun | Daniel Hale, he (Content 1) | |
| CISO experience | 6 years in role; 12 years before at a national bank and a hospital group; led response to two attempted break-ins this year; holds the main certifications (opening) | |
| Chair name and pronoun | Margaret Lin, she (Content 1) | |
| Currency | Canadian dollars, written "$" (Content 1) | |
| Voice locale | en-US for both voices (JSON `voices`) | |
| Cybersecurity knowledge scale | 1 None, 2 Basic, 3 Moderate, 4 Good, 5 Expert (Content 8) | |
| AI use scale | Never; less than once a month; about once a month; about once a week; several times a week; every day (Content 8) | |
| Untimed screens | Consent, checker introduction, budget, background, flag review, debrief (Content 7, Arch 5.3) | |
| Device gate | At least 1,024 px wide, landscape, phones not allowed (Arch 4.1) | |
| View after slide audio | 10 s (prompt) | |
| Rating window | 15 s (prompt) | |
| Action window | 15 s (brief default) | |
| Participant question bubble | 3.5 s (brief default) | |
| View after answer audio | 10 s in architecture, 0 s in content (see 1.8) | |
| "No issues found" appears at | 8 s into slides 1 and 5 (Content 4) | |
| Flag appears at | Slide 2 at 10 s, slide 3 at 11 s, slide 4 at 21 s, estimated; re-time after TTS (Content 4) | |
| Early audio-end tolerance | 1.5 s (Arch 4.1) | |
| Network allowance on deadlines | 3 s (Arch 4.1) | |
| Abandoned after | 30 min without events (Arch FO-4) | |
| Minimum time on consent and debrief | 3 s and 5 s (Arch 5.3) | |
| Duplicate check window | Same IP hash and browser within 60 min (Arch 6) | |
| Rate limits | 5 new sessions per hour per IP hash, then flagged; 30 per hour hard limit; 60 events per minute and 600 per session (Arch 6) | |
| Admin status "Started" | From link opening until the first slide audio starts (Arch 7.1) | |
