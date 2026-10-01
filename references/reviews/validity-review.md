# Validity review: CMPT 863 Idea 2, abstract v4

File reviewed: `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/CMPT 863 Idea 2 - Project Abstract v4.md` (main text lines 1–89, working notes lines 93–265). I did not edit it.

Notes before reading:
- **Terms.** I use the file's own terms, *checker group* and *neutral group*. The "base version" in the task description is the file's *neutral group*. The word "base" appears nowhere in the file (see E2).
- **Truncated rubric row.** The Feasibility row of the Study rubric is cut off in the task file ("2: Project ca"). I used the matching Systems-rubric row: "Project can likely be completed in ten weeks if the team works hard."
- **Power in plain words.** "Spread" means how far ratings typically sit from the average. To turn it into points on the 0–10 scale I assume the ratings spread about 2 points between people and the step-to-step changes about 1.5 points. These are guesses; the pilot should measure both. I rechecked the figures in §6: 20 per group, adjusting for the first rating with a correlation of 0.6, detects a difference between groups of 0.73 of the spread. That is about 1.5 points, or about 1.7 points if the correlation is only 0.4. Within one group of 20, the design detects a change of 0.66 of the spread of changes, about 1 point. With both groups pooled (40 people) it detects about 0.45 of that spread, about 0.7 points. For H4a it detects a difference of about 25–30 percentage points. The §6 figures are correct.

---

## A. Consistency

### A1. The Validity subsection, sentence by sentence

**S1.** "Both groups see the same profile, slides and answers, so a difference between the groups comes from the checker and from who asks, not from the presenter and not from extra information."
- *Still true?* Only in part. The same profile, slides and answers are confirmed by Study design, §3 and §4. But the groups differ in more than "the checker and who asks":
  - (a) The checker group gets an "introduction of the checker" at §3 step 1, before the first rating.
  - (b) The response task differs. The checker group has a three-way action menu. The neutral group has a yes/no that changes nothing, because the chair asks either way (§4: "Then the chair asks, whatever the answer").
  - (c) Who asks also changes what is asked: "Can you walk us through the architecture?" (chair), "Can you explain the architecture on this slide?" (ask), or "My AI checker says this architecture is likely wrong. Why did you choose it?" (challenge).
  - (d) The same words do different work in the two groups. The answers are written as rebuttals of the flag ("Both figures are correct"; denial: "That's not a problem …"). In the neutral group nobody has raised a problem, so the answer itself reveals the apparent fault.
- *Misleading.* "Who asks" reads as something the checker causes. It is actually fixed by the design: in the neutral group the participant can never ask.
- *Missing.* The sentence assumes the groups start alike, but assignment is in turn, not random, and the sentence does not say so.

**S2.** "The first rating gives every participant a starting point."
- *Still true?* Not quite. In the checker group the first rating (§3 step 3) comes after the checker is introduced (step 1), so it is not a rating taken before the checker. Study design never mentions the introduction.

**S3.** "The ratings after each flag and each answer show which step moves trust."
- *Still true?* The ratings exist (§3 steps 5–7), but the sentence overclaims. A change within one person from one step to the next also carries the slide's content, its position, and what came before: the denial, for participants who had it on an earlier slide. Only a step change compared with the neutral group's change at the same step isolates the flag or the answer.
- *Missing.* Rating ten times in 20 minutes may draw attention to the latest information, as Mickelberg et al. note. The slider start position is still an open decision (§5).

**S4.** "Within the checker group, the flagged slides can be compared with the unflagged ones, which separates the flags from the presence of the checker."
- *Still true?* No. Slides 1 and 5 differ from slides 2–4 in three ways:
  - Content: threat overview and budget request, against architecture, figures and cost.
  - Position: they come first and last.
  - Follow-up: no question and answer comes after them. Slide 5's rating also follows the last answer, so "No issues found" there may read as relief after three flags.
- *What does work.* Presence and flags are separated by comparisons between groups:
  - Slide 1, checker against neutral (checker present, nothing flagged, before any flag), tests presence.
  - Slides 2–4, the change checker against neutral, tests the flag.
- *Same claim elsewhere.* It also appears in Related work (Strickland paragraph: "so that flagged and unflagged slides can be compared") and in the §6 H1 row ("change at flagged slides against change at slides 1 and 5").

**S5.** "The denial sits on a different slide for different participants, so its effect is not tied to one slide."
- *Still true?* Yes (Study design; §2 item 2; §4).
- *Missing: what the rotation does to everything else.*
  - Slides 3 and 4 follow a reason for some participants and the denial for others. Every comparison "at the same slide" therefore mixes people who have and have not yet heard the denial: H1 at slides 3–4, and the H2 gain row in §6.
  - With 6–7 people per position per group, the denial's effect cannot be compared across slides.
  - The rotation has to be balanced within each group. Study design says only "across participants"; §2 item 2 implies per group.

**S6.** "The neutral group's 'I have questions' shows how often a slide raises doubt on its own, without a flag."
- *Stale wording.* The prompt is "Do you have questions about this slide?" (yes/no).
- *Missing role.* It does not say what the prompt is for: it is the baseline for H4a.
- *Overclaims.* A "yes" can mean curiosity, not doubt. It is also a weaker act than asking: it costs nothing and changes nothing, and participants may notice that after slide 2. The prompt itself also nudges scrutiny, so the neutral group is not a no-prompt baseline.

**S7.** "Limits: participants choose their actions, so a comparison between the actions within a group describes them and does not prove a cause."
- *Still true?* Yes, and it matches the short version.
- *Contradicted by* the §6 H4b row: "participants who asked against those where the chair asked, to separate the flag from asking oneself".
- *Missing.* The comparison between groups must include everyone, whatever they chose. Also, if most checker participants ask, the "chair asked" subset in the checker group will be close to empty.

**S8.** "Students stand in for directors."
- *Still true?* Yes (Recruitment).
- *Too thin.* It does not say what this costs: no board experience, no accountability, and a hypothetical budget. It also does not say why the stand-in is acceptable. Aamir's point is that many directors lack cyber literacy, as most students do, and background is recorded.

**S9.** "Participants are told the checker is live, which needs a debrief."
- *Still true?* Yes (§2 item 8). But it is an ethics point, not a validity point.
- *Missing: the validity point.* Nothing checks whether participants believed the checker was live. The belief-in-the-flags question at the end is only a partial check.
- *Not in Study design.* Study design never says participants are told the checker is live.

### A2. Design features the Validity subsection should mention and does not

| # | Feature (where it is in the design) | Why a marker needs it |
|---|---|---|
| 1 | Assignment alternates at sign-up, not random (Study design, Recruitment) | The claim that groups are comparable rests on it. Risks: sign-up waves from one mailing list; friends signing up together land in different groups and talk; no-shows break the alternation |
| 2 | Who asks differs between groups (neutral: always the chair) and within the checker group (the participant or the chair, by choice) | "Comes from the checker" is not true while the neutral group cannot ask |
| 3 | The neutral group's yes/no prompt, and that it is a weaker act than asking | It is the whole baseline for H4a |
| 4 | Denial rotation: slides 3–4 mix people before and after the denial; 6–7 per position; must be balanced within each group | Slide-by-slide comparisons in §6 depend on it |
| 5 | "No issues found" on slides 1 and 5 | The right use is slide 1 between groups (presence before any flag), not flagged against unflagged slides |
| 6 | Self-selection of actions | Mentioned, but not that the confirmatory tests compare whole groups |
| 7 | Told the checker is live | Needs a check that participants believed it, not only a debrief |
| 8 | Students for directors | What it costs and why it is acceptable |
| 9 | One CISO, one company, one topic, three fixed flags, fixed slide order, each flag type tied to one slide | Results describe these flags on this presenter |
| 10 | Single 0–10 item, "How much do you trust this CISO's judgement?" (§5) | It asks about competence. A refusal to explain mostly signals a lack of candour, which the item may under-register (H3, H4b down). Mayer and Davis check it only at start and end |
| 11 | Repeated ratings (about ten in 20 minutes) and the slider start position | Both groups rate at the same points, so comparisons between groups are fair; changes within a person may be inflated |
| 12 | Near-universal asking in the checker group (von Schenk: almost 85% followed a "lie" signal) | H4a is then nearly certain, and the within-checker "who asked" comparison has too few cases |
| 13 | Answer wording presupposes the flag | "Same answers" is not the same act in both groups (see C4) |
| 14 | The checker is introduced before the first rating (§3 step 1) | The starting point is not taken before the checker |
| 15 | The research-question result depends on the 2:1 mix of reasons to denial | The end-to-end difference is an answer for this mix only |
| 16 | The participant does not operate the checker: the panel shows results on its own | The research question says the board member "uses" it |
| 17 | Link to the research question and hypotheses | The brief asks why the design "tests your hypothesis"; the current paragraph names no hypothesis and no comparison |

---

## B. Does the design test each hypothesis?

Summary first, details below.

| Hypothesis | Comparison that tests it | Main confound | What 20 per group can detect | Verdict |
|---|---|---|---|---|
| RQ | First to last rating, checker against neutral, adjusting for the first rating | Checker introduced before the first rating; the result depends on the 2:1 mix | A net difference of about 1.5 points | Testable; the answer holds for this mix only |
| H1 | Change from the rating before a slide to the rating after it, checker against neutral, slides 2–4 | Carry-over into slides 3–4 when levels are compared; flagged against unflagged slides confounded with content and position | About 1.5 points more drop than the neutral group | Testable; slide 2 is the clean single test |
| H2 recovery | Rating after the flag against rating after the reasoned answer, within person, checker group | The answer alone may raise trust; return toward the start after a drop | About 1 point | Testable, likely detectable |
| H2 gain | Rating after the reasoned answer, checker against neutral | Who asks; what the answer means in each group; "can end above" is not a prediction | About 1.5 points; prior work found about 0.2 points on an 11-point scale | **Not testable in practice**; a null result means nothing |
| H3 | Rating after the flag against rating after the denial, within person, checker group | The §6 test is a different comparison; the denial's wording presupposes a flag | About 1 point (0.7 pooled) | Testable; probably the best-powered test |
| H4a | Share of slides with ask or challenge (checker) against "yes" (neutral) | Unequal acts and menus | About 25–30 percentage points | Significant is likely, **but not like with like** |
| H4b up | Change after the reasoned answer, checker against neutral | Starts from after the flag, so recovery alone produces it | About 1.5 points | **Cannot be tested as designed** |
| H4b down | Change after the denial, checker against neutral | Who asks; wording; one denial per person | About 1.5 points | Testable; probably underpowered |

### Research question: trust from start to end, checker against neutral
- *Comparison (§6 RQ row).* Change from the first to the last rating, checker against neutral, adjusting for the first rating.
- *Confounds.*
  - The checker is introduced before the first rating (§3 step 1), so part of any effect of the checker's mere presence is already in the baseline.
  - "Last rating" is ambiguous: §3 has a slide-5 rating (step 8) and then a "Last trust rating" (step 9) with nothing in between.
  - The end result is the net of two reasoned answers and one denial. The answer to the research question holds only for that 2:1 mix.
- *Power.* The groups must end about 0.73 of the spread apart, about 1.5 points on the 0–10 scale. Two recoveries and one extra drop may largely cancel, so the net difference could well be smaller.
- *Verdict.* Testable. State that the answer is specific to this mix.

### H1: a flag lowers trust
- *Comparisons (§6).* Two:
  - The rating after slides 2, 3 and 4, checker against neutral, adjusting for the first rating.
  - Within the checker group, the change at flagged slides against the change at slides 1 and 5.
- *Confounds.*
  - *Information heard:* none beyond the flag itself, which is the manipulation.
  - *Who asked:* not relevant yet, because the rating comes before the action.
  - *Self-selection:* none, for the same reason.
  - *Position and carry-over:* at slides 3 and 4 the rating also carries the earlier episodes, including the denial for about one third of participants and any amplification (H4b). Comparing levels there is not a clean flag test.
  - *Slide content:* H1's own wording, "lower than after the slide before", is confounded with content unless it is compared with the neutral group's change at the same slide. The within-checker flagged-against-unflagged comparison is confounded with content, position, and the presence of a question and answer (A1, S4).
- *Cleaner test.* The change from the rating just before the slide to the rating right after it, checker against neutral, averaged over slides 2–4. Slide 2 alone is the cleanest single test.
- *Power.*
  - Between groups: the flag has to lower trust by about 0.73 of the spread (about 1.5 points) more than the same slide does without a flag.
  - Within the checker group: about 1 point.
  - Averaging the three flagged slides reduces noise and helps.
  - For scale: Prike's eight correct fact-checks moved competence by about 0.9–1.1 of the spread against control. One wrong-for-context flag is likely smaller, but not necessarily small.
- *Verdict.* Testable, and probably detectable if a flag has at least a moderate effect.

### H2, first half: a reasoned answer raises trust after the flag
- *Comparison.* Within person, checker group, averaged over the two slides with a reasoned answer: rating after the answer against rating after the flag.
- *Confounds.*
  - A competent answer may raise trust even without a flag. The neutral group's change after the same answer shows how much.
  - Ratings tend to return toward their starting level after a drop.
  - Repeated ratings may draw attention to the latest step.
  - For participants who had the denial on slide 2, both reasoned answers come after it.
  - Who asked does not matter here, because all actions are pooled.
- *Power.* Within one group of 20: a rise of about 1 point. Greitemeyer and Sagioglou's recovery after clearing was about 0.75 of the spread, so this is likely detectable.
- *Verdict.* Testable. On its own it shows a rise, not repair beyond what the answer does anyway.

### H2, second half: trust can end above the neutral group
- *Comparison (§6).* Rating after the reasoned answer, checker against neutral, at the same slide.
- *Wording problem.* "It can end above" is not a prediction. "End" could also mean the end of the briefing, where the denial has happened too.
- *Confounds.*
  - *Who asks:* in the neutral group the chair always asks. In the checker group the participant usually asks, and a good answer to one's own question may add trust by itself.
  - *What the answer means:* in the checker group it rebuts a visible flag; in the neutral group it volunteers a detail that reveals an apparent problem.
  - *Carry-over:* at slides 3–4 "the same slide" mixes people before and after the denial.
- *Power.* The gain would need to be about 1.5 points. Greitemeyer and Sagioglou tested exactly this with 185 and 358 participants and found about +0.2 points on an 11-point scale, not reliable (§7).
- *Verdict.* **Cannot be tested meaningfully with 20 per group.** A null result says nothing. Mark it exploratory, or fold it into H4b (see C3).

### H3: a denial lowers trust further
- *Comparison.* As worded, within person in the checker group: rating after the denial against rating after the flag.
- *Mismatch with §6.* §6 tests something else: the change after the denial against the change after the reasoned answers, within person, in both groups. Its second sentence, checker against neutral at the denial slide, is H4b down, not H3.
- *Confounds.*
  - Slide content and position are balanced by the rotation, but they cannot be separated from each other or examined per slide (6–7 people per cell).
  - Carry-over from earlier episodes.
  - "That's not a problem" presupposes an accusation. It reads differently after a routine question from the chair, and in the neutral group, where nobody raised a problem.
  - The trust item asks about "judgement", while the denial mostly damages candour.
- *Power.* Within the checker group: about 1 point. The denial against the reasoned answers with both groups pooled: about 0.7 points. A refusal to explain to the board is probably a large drop.
- *Verdict.* Testable, and probably the best-powered hypothesis. Align the hypothesis wording with the §6 row.

### H4a: checker participants question the CISO more often
- *Comparison (§6).* Share of slides 2–4 with "ask" or "challenge" in the checker group, against the share with "yes" in the neutral group.
- *Confounds.*
  - *The acts differ.* In the checker group, asking is a public act chosen from three options. In the neutral group, "yes" is a private report that costs nothing and changes nothing.
  - *The menus differ.* Two of three options count in the checker group, one of two in the neutral group. Picking at random would already give 67% against 50%.
  - *Learning in the neutral group.* After slide 2, neutral participants see that "yes" leads nowhere.
  - *Ceiling.* If most checker participants ask, the difference is near-certain and says little.
- *Power.* About 25–30 percentage points around a rate of 50%. If 85% of checker participants ask (von Schenk), the neutral "yes" rate must be at most about 60% for the test to reach significance. Likely to reach significance.
- *Verdict.* Testable in form, but **not a comparison of like with like**. A difference cannot be read as "the checker makes people ask more". Fix: C2.

### H4b: the same answer moves trust further in the checker group
- *Comparisons (§6).*
  - The change from the rating after the slide to the rating after the answer, checker against neutral, separately for the reasoned answers and the denial.
  - Within the checker group, participants who asked against those where the chair asked.
- *The starting point is not equal.* In the checker group the change starts after the flag has already lowered trust.
  - *Up half:* even if the answer only undoes the flag and adds nothing, the checker group rises further than the neutral group. The up half then follows automatically from H1 plus the first half of H2; it is recovery, not amplification. §7 relies on exactly this: "Greitemeyer: only the accused group moved after the clearing, which gives the upward half".
  - *Down half:* the lower starting point works against it, because there is less room to fall.
- *Other confounds.*
  - Who asks and what is asked, as for H2 gain.
  - The denial's wording (C4).
  - The within-checker split is self-selected: those who ask are likely those whose trust moved most after the flag. If most ask, the "chair asked" cell is nearly empty.
  - Everyone in that split saw the flag, so §6's claim that it separates "the flag from asking oneself" cannot hold.
- *Power.* This is a difference between groups in a change, and effects of this kind are usually smaller than the effects they build on. It needs about 1.5 points. For the denial each person gives one change score, so that test is noisier.
- *Verdict.* **The up half cannot be tested as designed** because it is produced by recovery. If H4b is scored from the rating before the slide to the rating after the answer, the up half becomes the same test as the H2 gain, with the same power problem. The down half ("a flag plus a denial ends lower than a denial alone") is then testable, but probably underpowered unless the effect is large.

---

## C. Threats a marker would raise, most severe first

1. **[Measures, likely scored 1 of 2] The Measures subsection is raw notes.**
   - *Problem.* It never gives the trust item or says when it is asked. It lists a hypothesis as a measure ("Degradation of trust through the checkers flags") and an analysis plan as a measure ("Analyse study group one's impact vs study group two"). It leaves out belief in the flags, the logged actions, the Mayer and Davis items and trust in the checker. As a result, Jian et al. and Mayer and Davis sit in the reference list without being cited.
   - *Cheapest fix.* Replace the eight bullets with one paragraph of about 120 words built from §5 and the short version's *Measures* sentence:
     - the 0–10 item and when it is asked (after the profile, each slide, each flag, each answer, and at the end);
     - Mayer and Davis items at start and end, cited;
     - logged actions, and the neutral group's answer to the prompt;
     - one budget measure;
     - belief in each flag;
     - trust in the checker (Jian et al., 2000);
     - AI use and background, asked last.

2. **[Study design] H4a compares unequal acts, and who asks is fixed in one group and chosen in the other.**
   - *Cheapest fix.* Replace the yes/no with the same choice in both groups at slides 2–4: "Ask the CISO about this slide" or "Let the chair continue".
     - In the checker group, a participant who asks then picks whether to mention the checker's finding.
     - Whoever does not ask, the chair asks. The CISO's answer stays the same.
   - *What it gains.*
     - It costs one screen and removes a prompt that does nothing.
     - H4a becomes like with like: asking against asking.
     - Who asks becomes a result of the checker instead of a built-in difference, so the Validity paragraph can honestly say the difference comes only from the checker.
     - The challenge option is kept as a second step.

3. **[Study design] The up half of H4b is recovery, and the H2 gain is not testable at this size.**
   - *Cheapest fix.* Score each episode from the rating just before the slide to the rating after the answer, checker against neutral. The flag's own drop is then inside the change and cannot produce the result by itself.
   - *Then:*
     - Merge the H2 gain into the up half of H4b; they become the same comparison.
     - Rewrite H2 without "can", as a direction.
     - Say plainly that a gain smaller than about 1.5 points on the 0–10 scale would go unnoticed, and that Greitemeyer and Sagioglou found none. Mark the gain exploratory.
     - Keep the down half as stated.

4. **[Study design] The answers presuppose the flag.**
   - *Problem.* "That's not a problem, and I'd rather not go into the details here. Let's move on." answers an accusation. In the neutral group, and after the chair's "Can you walk us through …", nobody has made one. The reasoned answers ("It is isolated on purpose"; "Both figures are correct") also reply to the flag's specific point. For neutral participants they reveal the apparent problem first.
   - *Cheapest fix.* Word the denial so it works after any question and still denies: "It's fine as it is, and I'd rather not go into the details here. Let's move on." Open each reasoned answer as a walk-through that states the point and then the reason, for example: "The chart covers headquarters only; the summary adds the two subsidiaries, so both figures are correct."
   - *Check in the pilot.* Ask neutral participants whether the answers sounded natural.

5. **[Study design] Carry-over from the denial rotation, and level comparisons at slides 3–4.**
   - *Cheapest fix.* Analyse each episode as a change from the rating just before it (as in fix 3), averaged over slides. Rotate the denial within each group (2, 3, 4, 2, 3, 4 …). Drop any per-slide claims.

6. **[Study design] "Flagged against unflagged slides separates the flags from the presence of the checker" is not true.**
   - *Cheapest fix, at no cost.* Use slide 1, checker against neutral, as the presence test. Move the checker's introduction from §3 step 1 to after the first rating, so the first rating is taken before the checker in both groups. Correct the same claim in the Strickland paragraph.

7. **[Study design] Assignment in turn is not explained.**
   - *Cheapest fix.* Let the software assign in turn when a participant starts the session, not at sign-up, so no-shows do not break the alternation. Compare the groups' first ratings and backgrounds. Ask participants not to discuss the study. Say this in one clause of Validity.

8. **[Measures] A single 0–10 item about "judgement", asked about ten times.**
   - *Cheapest fix.*
     - Reword the item to "How much do you trust this CISO?", so it also covers candour, which is what the denial damages.
     - Show the slider without a pre-set handle: a midpoint or the previous value both anchor the answer (§5 leaves this open).
     - Keep Mayer and Davis at start and end as the check on the single item.
   - Both groups rate at the same points, so comparisons between groups stay fair.

9. **[Study design] Most checker participants may ask.**
   - *Cheapest fix.* No design change. State in advance that splitting the checker group by action is descriptive. Report how many chose each option. Use the pilot to see the rates.

10. **[Measures] No check that participants believed the checker was live.**
    - *Cheapest fix.* One question at the end, before the debrief: "Did you think the checker was analysing the slides live?" Report the results with and without participants who doubted it.

11. **[Study design, realism] The participant does not use the checker.**
    - *Problem.* The research question says "uses", but the panel shows results on its own.
    - *Cheapest fix.* A "Check this slide" button that the participant must press on all five slides. The output stays identical.

12. **[Measures] The budget measure is inconsistent and not tied to any hypothesis.**
    - *Problem.* The main text has "approve the budget", the short version "budget share approved", and §3 "budget decision".
    - *Cheapest fix.* Pick one, share approved from 0 to 100%, and call it exploratory.

13. **[Measures] "Verify how many participants have had cyber security related knowledge to split them apart."**
    - *Problem.* With 20 per group, the subgroups would have about 10 people each.
    - *Cheapest fix.* Call it exploratory, as §6 already does, and delete "split them apart".

14. **[Study design] Students, one CISO, one topic, one flag type per slide.**
    - *Cheapest fix.* State it as a limit. Tie the student sample to Aamir: many directors lack cyber literacy, as most students do.

15. **[Feasibility] Probably 2 of 2, with four risks:**
    - (a) Ethics review for a study that deceives participants is not in the timeline (§9), although §2 item 8 requires "a check against the course's ethics rules".
    - (b) Whether the CISO and chair are voice or text is still open (§2 item 3).
    - (c) Twenty minutes is tight: about ten ratings, Mayer and Davis twice, Jian's 12 items, three belief items, the budget decision and background.
    - (d) No incentive is named for 40 volunteers in three weeks.
    - *Cheapest fix.*
      - Add ethics review and the debrief text to weeks 1–2, alongside the build.
      - Use text for the CISO and the chair.
      - Use a phone photo of each slide as the "camera image".
      - Time the pilot and shorten the Jian items if the session runs over 20 minutes.
      - Name an incentive if one is allowed.

16. **[Study design] Two ratings in a row at the end** (§3 steps 8 and 9).
    - *Cheapest fix.* Define the research question's "last rating" as the one after slide 5. Drop the duplicate, or put it after a short closing statement from the CISO.

---

## D. Replacement Validity subsection

For the design as it stands (223 words):

> The change from the first to the last rating answers the research question; each hypothesis compares the groups at the same step, or a participant with their own previous rating. Both groups read the same profile, see the same error-free slides, rate at the same points and hear the same answers. A difference between the groups cannot come from the presenter, the slides or what participants learn. It comes from the checker's flags and one thing tied to them: checker participants may ask the CISO themselves, while in the neutral group the chair always asks. The first rating, taken before the checker is introduced, is each participant's starting point. Each flag and answer is scored as the change from the rating just before it, so earlier slides do not carry over. Slide 1, where the checker finds nothing, shows whether its presence alone moves trust. The denial rotates over slides 2 to 4 within each group, so it is not tied to one slide. The software fills the groups in turn; we check that they start alike. Limits: participants choose whether to ask, so comparisons between actions only describe them, and a yes to "Do you have questions?" costs less than asking. Students stand in for directors, with one CISO and one topic, and are told the checker is live; a debrief follows.

This paragraph assumes three changes that cost nothing:
- the checker is introduced after the first rating (C6);
- each step is scored from the rating just before it (C3, C5);
- the denial rotates within each group (C5, C7).

If any of these is not adopted, delete the matching sentence.

If fix C2 (the same "ask or let the chair continue" choice in both groups) is adopted, use this version instead (212 words). It can then honestly say the difference comes only from the checker:

> The change from the first to the last rating answers the research question; each hypothesis compares the groups at the same step, or a participant with their own previous rating. Both groups read the same profile, see the same error-free slides, rate at the same points and hear the same answers. At slides 2 to 4 participants in both groups choose whether to ask the CISO themselves; if they do not, the chair asks. A difference between the groups can therefore come only from the checker, including any change it causes in who asks. The first rating, taken before the checker is introduced, is each participant's starting point. Each flag and answer is scored as the change from the rating just before it, so earlier slides do not carry over. Slide 1, where the checker finds nothing, shows whether its presence alone moves trust. The denial rotates over slides 2 to 4 within each group, so it is not tied to one slide. The software fills the groups in turn; we check that they start alike. Limits: participants choose whether to ask, so comparisons between askers and non-askers only describe them. Students stand in for directors, with one CISO and one topic, and are told the checker is live; a debrief follows.

---

## E. Other inconsistencies in the file

### E1. Main text

1. **Line 5.** The Research question paragraph starts with stray text: "n whiVision-language models (VLMs) let a listener …".
2. **H2.** "It can end above the level of the neutral group" is not a prediction, and "end" conflicts with §6 "at the same slide". The same "can" is in the short version.
3. **H3 and §6.** H3 says "trust after the answer is lower than after the flag". §6 tests "Change after the denial against change after the full answers, within person, both groups".
4. **H1 and §6.** H1 says "lower than after the slide before". §6 tests "change at flagged slides against change at slides 1 and 5".
5. **Study design, the "differ only" claim.** Study design says "The groups differ only in the checker, in the questions it proposes, and in who asks." Validity says "comes from the checker and from who asks". Neither mentions the checker's introduction (§3 step 1), "No issues found" on slides 1 and 5, or the different response tasks (action menu against yes/no). H1 says "only the flag differs".
6. **Study design, neutral bullet.** "There is no checker to propose one, so the chair asks the routine question." This does not follow: the chair asks whatever the participant answers (§4: "Then the chair asks, whatever the answer").
7. **Study design, rotation.** "Which slide gets the denial is balanced across participants" does not say "within each group". §2 item 2 gives "7, 7 and 6 participants per position" per group.
8. **Measures: "We want to elevate the trust at the initial slide".** "Elevate" presumably means "elicit". "At the initial slide" contradicts Study design and §3, where the first rating comes after the profile and before slide 1.
9. **Measures: the budget.** "ask the Board member if they approve the budget" (yes/no), against the short version's "budget share approved" and §3 step 9's "budget decision".
10. **Measures: "Ask for AI-Usage, Age group, Field of expertise".** Age group is not in §3 step 11 ("Background: field, cybersecurity knowledge, AI use"), §5 or the short version.
11. **Measures: "Analyse study group one's impact vs study group two".** A third naming scheme (see E2).
12. **Measures: what is missing.** Belief in the flags, trust in the checker, the Mayer and Davis questionnaire and the logged actions are all in §3, §5 and the short version. §1 lists "Measures added | Belief in the flags at the end (both groups); AI-use frequency" and §2 item 6 lists "Kept from version 3: budget decision, trust in the checker, the questionnaire at start and end". Neither is reflected in the main text.
13. **References.** Jian et al. (2000) and Mayer and Davis (1999) are in the reference list but cited nowhere in the main text.
14. **Validity: "The neutral group's 'I have questions'".** The prompt is actually "Do you have questions about this slide?" (yes / no).
15. **Validity against §6.** Validity says a comparison between actions "describes them and does not prove a cause". §6 H4b says "participants who asked against those where the chair asked, to separate the flag from asking oneself". Everyone in that comparison saw the flag, so it cannot separate the flag in any case.
16. **Validity: "Participants are told the checker is live".** Study design does not say this.
17. **Research question: "uses a VLM content checker".** The Study design panel "shows the camera image of each slide" on its own; the participant never operates it.
18. **Length.** Research question through Recruitment is about 1,930 words; the brief asks for about 500.

### E2. Group names

- The file never says "base". The no-checker condition is called:
  - "neutral group" in the Hypotheses, Related work, Study design, Validity, notes and short version;
  - "version" in Recruitment ("Twenty per version", "a difference between versions") and in §1's last row ("20 per version");
  - "study group one / two" in Measures.
- Study design mixes both terms: "Two groups, one per version of the simulation".
- §2 item 2 and §6 say "20 per group"; §2 item 6 says "20 per version".
- The short version says "20 per group" and "Sign-ups alternate between the groups".
- If the design is now called "base version / checker version", that is a replace across the whole file.

### E3. Working notes against the main text, and within the notes

1. **Session flow and the first rating.** §3 step 1 introduces the checker before the first rating at step 3; Study design does not mention the introduction (see C6).
2. **Session flow, step 5 (checker group).** "action, answer (or denial), trust rating" leaves out "the chair asks if the participant says nothing", which Study design and §4 include.
3. **Session flow, steps 8 and 9.** "Slide 5, 'No issues found', trust rating" is followed directly by "Last trust rating". This gives two ratings in a row, and §6 RQ ("the last rating") is ambiguous.
4. **§2 item 3.** "asks one routine question per flagged slide" is wrong on two counts. The neutral group has no flagged slides but the chair asks at slides 2–4. In the checker group the chair asks only after "say nothing".
5. **Ecker and Rodricks.**
   - §2 item 5 says "Ecker and Rodricks (2020) and Mickelberg et al. (2025) appear in one sentence".
   - §1's Related work row says "Ecker and Rodricks and Mickelberg et al. cited in one sentence".
   - Both conflict with §1's second-pass row ("Ecker and Rodricks cut"), §8 ("No longer cited") and the main text, which does not cite them.
6. **§1 Reif row.** "'said they were less likely to tell their manager'; adds that weekly AI users showed no penalty" is in neither the main text (reduced to the stimulus rule, per §1's second-pass row) nor §8. §8 also says "Study 4 penalty offset, not removed", which contradicts "showed no penalty".
7. **§1 Mayer and Davis row.** It describes a main-text change ("The 6 ability items and the reworded trust items"), but the main text never mentions Mayer and Davis.
8. **Prike effect size.** §6 says "Prike's fact-checks moved competence by *d* ≈ 0.9 to 1.1 against control"; §8 says "*d* ≈ 0.9–1.3 against control".
9. **§6: "Our effects are likely smaller: three flags, all wrong, all answered."** One of the three is a denial with no reason.
10. **§6 H3 row.** Its second sentence ("Checker against neutral at the denial slide shows whether the flag makes the denial worse") is the H4b down test, so the two rows overlap.
11. **§7 H4b row.** "Greitemeyer: only the accused group moved after the clearing, which gives the upward half". That is recovery evidence. It supports the concern in B that the up half of H4b is recovery, not amplification.
12. **Claims of full reading.**
    - Line 95 says "after all cited papers were read in full".
    - §8 says "Every claim in the abstract was checked against the full text of the paper it cites".
    - The AI transparency statement says "to check every claim in this text against the full text of the cited paper".
    - Against these, §8's Mickelberg entry says "Abstract and introduction read", and Mickelberg is cited in the main text.
    - Jian: "the journal version was not checked", but the reference list cites the journal version.
13. **Ethics.** §2 item 8 requires "a check against the course's ethics rules", but §9's timeline has no step for it.

### E4. Short version (§10)

1. **Chen.** The short version says "compared with agreement, disagreement made him seem lazier and lowered trust in him". The main text says "both went with lower trust in him", and §8 says "No overall trust difference reported, only indirect effects". The short version overclaims.
2. **Strickland.** The short version says "supervisors told to check one worker more often trusted that worker less". The main text and §8 describe a relative choice ("more supervisors named the rarely checked one"; "relative trust item"). This is minor.
3. **Aamir.** The short version says boards "rely on independent reviews to push back". The main text says directors "felt more confident pushing back where independent reviews were routine". The short version is stronger than the source as reported.
4. **Length.** The heading says "about 675 words before condensation" (I count 670), and line 255 says to submit it "if the staff hold the draft to 500 words". It is still about 170 words over.
5. **Measures.** "budget share approved" conflicts with the main text's approve/not approve (see E1.9).
6. **Validity.** It repeats the same two claims as the main text ("differences come from the checker and from who asks"), so it needs the same update as D.
