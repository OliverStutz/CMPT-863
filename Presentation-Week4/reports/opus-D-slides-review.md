# Opus report D: slides review
Independent review, 2026-09-30, of "Week 4 - Slides.md" against the paper PDF and the course brief. Written without sight of report C.

## Verdict

The deck is well researched. Most numbers match the paper, and Part B has real critique, which puts it above "adequate" on the rubric. It is fine to submit as the draft once the TODOs are filled, but it is not ready for class. Four kinds of problem remain:
- **Factual slips on screen.** A classmate who has read the paper can catch these: the widget list on slide 9, the tail values on slide 9, "what the paper is named after" on slide 16, and the "hold-and-tap" label on slide 10.
- **Procedures the paper never describes.** Two of the critique slides present such procedures as fact: finger count ignored during grouping, and the "losing" referent rescued by re-binding (slides 7 and 17).
- **Contradictions inside the deck.** Slide 12 says "the premise holds," slide 14 says the blocks world "avoids the Windows bias," and slide 20 cites GI 2010. All three cut against the thesis on slide 21.
- **No time margin.** The spoken script is about 3,500 words, plus transitions and an unwritten slide 19, which leaves no buffer.

**The single biggest risk** is the slide 18 to slide 21 argument. It attacks "unrevised behavior" as if the paper had promised a participant with no prior history. The paper defines the term as behaviour not revised by recognizer feedback (p. 1085). The paper also reports the desktop bias itself, as a finding (p. 1090). A sharp classmate or the instructor can call this a straw man, and the deck's thesis rests on it.

## Blockers (must fix before class)

1. **Slides 18 and 21: "unrevised behaviour" is misread.**
   - **Slide says:** Slide 21: "'Unrevised behaviour' assumes a participant without a touch history. In 2008 that person did not exist." Slide 18's title: "an unbiased participant did not exist in 2008."
   - **What is wrong:** The paper never claims an unbiased or history-free participant.
   - **Evidence:** p. 1085 defines unrevised behaviour as behaviour with "the gulf of execution" removed, "a monologue in which the user's behavior is always acceptable." In other words, behaviour not revised by recognizer feedback. On p. 1090 ("It's a Windows World") the authors report prior-paradigm bias themselves: "I'm falling back on the old things that I've learned." The screening on p. 1084 was for "non-technical people without prior experience using touch screens", which is a claim about who the participants were, not a claim that they were unbiased.
   - **Fix:** Reframe as "the paper treats prior experience as a finding (the Windows world). We argue it is a limit of the method, and that the source is broader than the desktop." Retitle slide 18, for example "Whose habits did the study capture?"

2. **Slide 21 against slides 12, 14 and 20: the deck contradicts its own thesis.**
   - **Slide says:** Slide 21: "The gesture set was never validated" and "the paper's own evidence for the method is weaker than its reputation."
   - **What is wrong:** Three other slides cut against this. Slide 12 says the paper tests its own premise, and the notes add "the premise holds." Slide 14 lists "Tests its own premise" as a strength. Slide 20 reports GI 2010, where 22 new participants judged 81 gestures and preferred the ones proposed by larger groups. That is a partial external check of the method, though not the guessability study the paper proposed.
   - **Evidence:** p. 1090 (designer comparison), p. 1091 (validation as next step). GI 2010 as given on slide 20.
   - **Fix:** Narrow slide 21 to "the set's guessability was never tested, and the in-paper evidence for 'popularity = quality' is self-report." Acknowledge GI 2010 as partial support.

3. **Slide 9: the widget list is wrong, and two recommendations are merged into one.**
   - **Slide says:** "The authors' rule: low agreement → give the command an on-screen widget… (menu, help, undo, accept, reject, zoom)."
   - **What is wrong:** The paper's list is "delete, zoom in, zoom out, accept, reject, menu access, and help." It does not include undo, and it does include delete. That list is also a second recommendation, grounded in "enough participants acted on imaginary widgets." It is separate from the low-agreement rule ("referents beyond this point [the first seven] may benefit from an on-screen widget as well as a gesture").
   - **Evidence:** p. 1091, "Implications for User Interfaces."
   - **Fix:** Give two lines. (a) "Beyond the first seven referents: widget as well as gesture." (b) "Imaginary-widget behaviour: consider widgets for delete, zoom in/out, accept, reject, menu, help."

4. **Slide 9: the tail values are wrong.**
   - **Slide says:** "The tail, *task switch* to *cut*, sits at 0.10–0.15."
   - **What is wrong:** Read from Figure 3 (one hand), the values are: task switch ≈0.22, zoom out ≈0.21, previous ≈0.19, undo ≈0.19, accept ≈0.17, help ≈0.16, reject ≈0.15, next ≈0.14, cut ≈0.12. Only the last three or four fall in 0.10–0.15.
   - **Evidence:** Figure 3, p. 1088.
   - **Fix:** Write "task switch to cut: about 0.22 down to 0.12."

5. **Slide 16: the paper is not named after guessability.**
   - **Slide says:** "Guessability … It is what the paper is named after."
   - **What is wrong:** The title is "User-Defined Gestures for Surface Computing." Guessability names the methodology, not the paper. Your own speaker notes say "the method is named after."
   - **Evidence:** p. 1083 (title). The method is "a guessability study methodology [33]" (p. 1083). Guessability is also an author keyword.
   - **Fix:** Write "It is what the method is named after."

6. **Slides 7, 11 and 17: grouping procedures the paper does not report are presented as fact.**
   - **Slide says:** Slide 7: "Finger count is ignored." Slide 7 Step 3: the losing command "gets a different gesture or the same gesture bound to a different target." Slide 11 #2: "it also drove the grouping decision." Slide 17 #1: "Finger count was ignored when deciding what counts as 'identical'." Slide 17 #2: "rescuing the loser by re-binding." The notes cite "p. 1087 (grouping)" for this.
   - **What is wrong:** p. 1087 says only "we grouped the gestures within each referent such that each group held identical gestures" and "the referent with the largest group won the gesture." It does not say how finger count was treated in grouping. It does not say what happened to the referent that lost a conflict. The nearest support is indirect:
     - the form coding (p. 1086): a one-point touch stays a one-point touch "even if the user casually touches with more than one finger";
     - the Figure 4 caption (p. 1089): 1–3 fingers allowed;
     - the set's reliance on target-based reuse (p. 1088, 1091).
   - **Why it matters:** Slide 17's first two "decisions" are the presenters' inferences, so their critique stands on unreported procedure. The authors could reply that target is part of a gesture's binding (Table 2), so enlarge and zoom in were never "the same gesture."
   - **Fix:** Rephrase as an inference: "The paper does not say how finger count was handled in grouping. The coding rules and the Figure 4 caption suggest 1–3 fingers were treated as one. If so, …" Drop the word "rescued."

7. **Slide 18: the causal claims are unsupported, and one row contradicts the conclusion.**
   - **Slide says:** "Every mass-market item teaches one finger, one point. The 72% is the sum of all single-point interfaces, not 'desktop habit' alone."
   - **What is wrong, three ways:**
     - **The table contradicts the sentence.** It lists the Wii, a mass-market device, as teaching "whole-hand motion."
     - **The cause cannot be tested.** No touch history was recorded, as the slide itself says. That means the 72% cannot be attributed to ATMs, iPods or Palms any more than to the mouse. Replacing one single-cause story with another is the same error the slide accuses the paper of.
     - **Single-point gestures need no device to explain them.** Pointing and pushing with one finger is ordinary physical behaviour. The paper cites Beringer and Voida et al. finding pointing dominant (p. 1084). Physical-nature gestures are 43.9% of all gestures (Figure 2; p. 1091), and the single-finger drags for move are physical.
   - **On the pinch:** "the final set contains a pinch and a reverse pinch" implies the pinch came from iPhone adverts. The paper notes that only 5 people used a reverse pinch for resizing (p. 1091), which is weak evidence of strong media priming.
   - **Evidence:** p. 1084, p. 1090, p. 1091, Figure 2 (p. 1087).
   - **Fix:** Say "could also come from" rather than "is." Remove the Wii row or address it. Present the media-exposure point as a question, not a finding.

8. **Whole deck: there is no time buffer in practice.**
   - **What is wrong:** The "Say" blocks total about 3,515 words. Adding about 20 transitions takes that to about 3,900 words, before slide 19 is written. At 130–140 wpm that is 28–30 minutes. That counts no builds, no hand-overs between speakers, and no audience reaction.
   - **Slides over their budget:** Slide 18 is 308 words for a 2:00 slot (2:15 or more). Slide 5 is 205 words for 1:30. Slide 20 is 196 words for 1:30. Slide 13 is 188 words for 1:30.
   - **Fix:** Cut about 600 spoken words. See Q5.

## Should fix

1. **Slide 16 and its speaker note: the 0.37 is printed in the paper.**
   - **Slide says:** "*Select single*: 11 of 20 tapped → A ≈ 0.37." The note says "our estimate from the counts the paper gives, not a number printed in the paper."
   - **What is wrong:** The paper prints it as Eq. 3, with groups 11, 3, 3, 3 for one-hand select single, giving 0.37 (p. 1087). "Tapped" is an inference from Select Single₁ = tap in Figure 4.
   - **Fix:** Cite Eq. 3, p. 1087, and remove the "≈."

2. **Slides 7 and 9: "move a little" appears with two values.**
   - **What is wrong:** Slide 7 says move a little = 0.42. Slide 9 says "Move a little and move a lot: A = 1.0. Everyone dragged." The first is the two-hand score (p. 1087) and the second is the one-hand score (Figure 3, p. 1088). Unlabelled on screen, they look like an error.
   - **Fix:** Label each one "2-hand" or "1-hand" on the slide.

3. **Slide 10: the idiom names do not match Figure 4.**
   - **Slide says:** "*Hold-and-tap* (two hands)… reused for move, duplicate, delete, paste, minimize." Also "Menu: drag in from the screen edge like a drawer."
   - **What is wrong:**
     - In Figure 4, "hold and tap" is the label for **Select Group₁**.
     - The two-hand idiom the slide describes is **Move₂: "jump"** ("Object jumps to index finger location"). Delete₂, Paste₃ and Minimize₂ reuse Move₂.
     - Duplicate is a separate gesture, "tap source and destination."
     - Menu is labelled only "pull out." The drawing shows the pull starting on an object, not at the screen edge, so the "drawer from the edge" reading is not supported.
   - **Evidence:** Figure 4, p. 1089.
   - **Fix:** Rename the idiom "jump (hold source, tap destination)." Write the menu gesture as "pull out."

4. **Slide 17 #4: "minimize became shrink" does not match the set.**
   - **Slide says:** "Maximize became enlarge, minimize became shrink."
   - **What is wrong:** Maximize does reuse enlarge (p. 1087). But in the set, minimize = "drag to bottom of surface" or Move₂ to the bottom (Figure 4). The paper reports only that the *mental models* were similar for shrink and minimize (p. 1090).
   - **Second issue:** "This is why Figure 3's order differs from Table 1" is more causal than the paper. The paper says the orders differ, "thus participants and the authors did not always regard the same referents as complex," and then describes the simplification.
   - **Fix:** Write "users said minimize is like shrink, but the set gives minimize its own gesture," and soften "this is why."

5. **Slide 13: the method claim is misleading.**
   - **Slide says:** "The authors' own list of four does not include the method."
   - **What is wrong:** It is true of the numbered list on p. 1084. But the abstract leads with "We present an approach to designing tabletop gestures that relies on eliciting gestures…" (p. 1083). The authors also explicitly credit the procedure to [33] (pp. 1085, 1087). Leaving the method off the list reads as honest attribution, not an oddity.
   - **Fix:** Write "the numbered list omits the method, because they credit it to EdgeWrite 2005. The abstract still leads with it."

6. **Slides 13 and 20: "nobody shipped it" and "the set went nowhere" cannot be verified and are open to challenge.**
   - **What is wrong:** Several gestures in the set are now universal in products: tap, drag, double-tap to open, pinch and spread to resize or zoom, drag off-screen to dismiss, and pull-out menus. What did not happen is adoption of the *set as a set*. Someone in the room will name a counterexample.
   - **Fix:** Write "no product adopted the set as published. Many of its members are now standard, which cuts both ways."

7. **Slides 15 and 22: the ratings critique is not "circular."**
   - **Slide says:** "'Popularity identifies better gestures' is partly circular" and "Is the validation circular?"
   - **What is wrong:** Participants did not know the group sizes when they rated their gestures (p. 1088), so the association is not circular. The real issues are three:
     - the ratings are self-reports;
     - nobody independent rated the gestures;
     - the authors overstate it ("goes a long way to validating," p. 1088).
   - **Fix:** Write "self-rated, not independently judged," and ask the question in that form.

8. **Slide 12: "could not predict what users would do" is too strong.**
   - **What is wrong:**
     - Each author matched 43.5% of the user set, and the three combined matched 60.9% (p. 1090). That is substantial overlap, not a failure to predict.
     - The yardstick is the user set itself. Designers' sets are scored by coverage of the users' consensus, so they cannot win by construction.
     - Three people combined are naturally outscored in breadth by twenty, which the paper itself concedes: "three experts cannot generate the scope of gestures that 20 participants can."
   - **Fix:** Write "designers predicted under half individually, and about 60% together. The comparison measures overlap with the users' set, not quality."

9. **Slide 14: "Avoids the Windows bias" contradicts slide 18 and the paper.**
   - **Slide says:** "Neutral blocks world. Avoids the Windows bias that undermined Epps et al."
   - **What is wrong:** The paper itself found that "participants still often thought of the desktop paradigm," and 72% of gestures were mouse-like (p. 1090). Slide 18 argues the same.
   - **Second issue:** Epps et al. said their prompts "may have biased" participants (p. 1084). "Undermined" is stronger than that, and so is slide 5's "had pushed people toward mouse emulation."
   - **Fix:** Write "removes desktop imagery. Desktop habits showed up anyway, which is itself a finding."

10. **Slides 4 and 19: the paper says "users," not "designers," about Long et al.**
    - **Slide says:** "Designers are poor at picking gestures that are easy to tell apart (p. 1084)" and "Long et al.: designers misjudge gesture similarity."
    - **What is wrong:** The paper says "users are sometimes poor at picking easily differentiable gestures" (p. 1084). The participants in Long et al. were users of a design tool. The slide's attribution is defensible, but the quote is wrong and the hedge "sometimes" is lost.
    - **Fix:** Write "Long et al.: people designing gestures are sometimes poor at picking easily differentiable ones."

11. **Slide 15: the source of the commands is misstated.**
    - **Slide says:** "The 27 commands came from desktop systems."
    - **What is wrong:** The paper says "obtained from desktop and tabletop systems [7,17,27,31,34,35]" (p. 1085), and every cited reference is a tabletop or multi-display paper. The defensible point is that several of the commands are desktop concepts (cut, paste, minimize, maximize, task switch).
    - **Fix:** Write "drawn from desktop and tabletop systems. At least five are desktop concepts."

12. **Slide 8: the form shares are for the right hand only.**
    - **Slide says:** "Shares of all 1,080 gestures."
    - **What is wrong:** Figure 2 splits form by hand. The slide's 45/27/14/9/4 are the right-hand column. The left-hand column is 37.9/34.3/17.7/3.7/5.6, which also sums to about 72% one-point.
    - **Two smaller gaps:** "Dynamic pose and path" (0.7%) is missing. In binding, "mixed" (2.7%) is missing, so the binding shares do not sum to 100.
    - **Fix:** Add "right hand" and "+ mixed 3%."

13. **Slide 9: the two-hand line is described incompletely.**
    - **What is wrong:** The two-hand line's highest point is *insert*, at about 0.58. It is above the one-hand line by about 0.10, and the slide omits it. Close dips to about 0.18 and paste to about 0.13, not "near 0.1."
    - **Evidence:** Figure 3, p. 1088.
    - **Fix:** Add insert, and write "close ≈0.18, paste ≈0.13."

14. **Slide 2 and its speaker note: the command counts are off.**
    - **What is wrong:** The note says "Five commands got no gesture of their own." The paper says four: insert, maximize, task switch, close (p. 1087). "22 of 27" is the paper's own count, made after merging the two move referents. In fact 23 referents have dedicated gestures, and 2 more are covered by reuse.
    - **Fix:** Correct the note to four. Optionally add "(the two moves merged)."

15. **Speaker note for slide 12: the AI reviewers are mentioned in the spoken script.**
    - **What is wrong:** The script says "Both Opus reviewers we asked, independently, picked this…" Mentioning AI reviewers in class will distract and invites questions about authorship.
    - **Fix:** Delete the sentence.

16. **Missing critique that a grad seminar will expect.**
    - **Non-independence in the statistics.** All the χ² and logistic tests use N = 1074 gestures from only 20 people, each contributing about 54 gestures (pp. 1087–1088). That treats repeated measures as independent observations (pseudo-replication) and inflates significance.
    - **Your own caveat is half-answered by the paper.** The authors' complexity ratings correlate r = .71 with participants' planning time, and simple referents took about 8 s of planning against about 15 s for complex ones (p. 1088). That is behavioural evidence that the ratings are valid. It partly answers slide 11 #5's "complexity was rated by the authors themselves." Say so, to be fair.
    - **Design implications are missing.** They are one of the authors' four claimed contributions, yet the deck covers almost none of them. Examples: "Only 2 of 14 participants for 2-hand enlarge resized along the diagonal" (no precise positioning), hit-testing at object corners, four fingers as the boundary, and select-then-act negating "a key benefit of gestures" (pp. 1090–1091).
    - **Fix:** Add the statistics point to slide 15. Add the r = .71 caveat to slide 11. Add one implications line to slide 9 or 10.

17. **Redundancy across slides.**
    - **11 #1 and 12:** the same three numbers twice.
    - **15 "worth adding", 16 and 17:** the 0.05 floor, the 43% uncovered, self-rating and grouping each appear twice.
    - **21:** restates 15, 16 and 18.
    - **22:** its smaller items repeat 11 (3.4 s) and 19 (physics).
    - **Fix:** See Q5.

18. **Slide 19: course connections are still a TODO.**
    - **What is wrong:** One of the seven course questions is "How does it connect to other research or topics we have discussed?" At the moment it is answered only by in-paper references. A list of the paper's own references reads as summary.
    - **Fix:** Write at least two explicit links to weeks 1–3 readings, each with a "so what."

19. **Slide 15: "Two authors with a commercial stake" sounds like an accusation.**
    - **What is wrong:** The paper shows only an affiliation (p. 1083), and the study ran on a Microsoft prototype (p. 1085).
    - **Fix:** Write "two authors at Microsoft Research, run on their own prototype; the paper does not discuss this."

## Nice to have

1. **Slide 3: the study date cannot be verified.** "2008 · this study runs" is not in the paper, which gives no study dates. With a CHI 2009 submission deadline in September 2008, the study could also predate 2008. Slide 18's title depends on this date. Fix: write "by 2008."
2. **Slide 18 table: weak rows.**
   - "ATMs, kiosks, POS, GPS – 1980s" is wrong for GPS: consumer touch-screen GPS units date from the mid-2000s.
   - The BMW footnote does not support the argument.
   - DS and Wii ownership among adults averaging 43 is assumed.
   - Fix: trim the table to 4 rows.
3. **Slide 6: a mislabel and an unstated order.** The "Control" row is not a control condition; call it "Designer baseline." "One-hand gesture, then two-hand gesture" (also on slide 2): the paper does not state the order or whether it was counterbalanced (p. 1086).
4. **Slide 17 #3: the Morris et al. 2014 link is stretched.** "Production" there means asking for several proposals per referent. Forcing a two-hand variant is related but not the same. Fix: write "related to."
5. **Slide 13: Nielsen et al. is missing.** The paper itself says Nielsen et al. [19] "describe a similar approach" (p. 1084). That bears on "first to employ users" more directly than Good et al.
6. **Slide 11 #6: wrong cross-reference.** It points to "(slide 16)," but the self-rating critique is on slide 15.
7. **Slide 15: the floor is not always 0.05.** With 6 discarded gestures, some referents have |Pr| = 19, so their floor is 1/19 ≈ 0.053. Write "≈0.05."
8. **Commentary markers are missing on the slides.** The notes say slides 15, 18, 22 and 23 carry commentary hooks, but only slide 23 has a placeholder. Add small "[classmate]" markers so the integration of commentaries (4 points) is visible.
9. **Speaker note for slide 13 overclaims.** "216 gesture elicitation studies descended from this one" is not what the review says: it reviewed 216 studies, not all of which derive from this paper. The slide's own wording is fine.

## Fact check table

| Slide | Claim | Paper says (printed page) | Status |
|---|---|---|---|
| 1 | Authors, affiliations, CHI 2009 Boston, pp. 1083–1092 | Header and author block (1083), page numbers 1083–1092 | OK |
| 2, 6 | 20 participants, Surface prototype, 27 commands | 1085 | OK |
| 2, 6 | 1,080 gestures, 6 discarded | 20 × 27 × 2 = 1080; 6 discarded (1086) | OK |
| 2, 6 | Filmed from four angles | "videotaped from four angles" (1086) | OK |
| 2, 6 | One hand, then two | "performed a 1-hand and a 2-hand gesture"; order not stated (1086) | UNVERIFIABLE |
| 2, 10 | 48 gestures, 22 of 27 commands | 48 gestures (1088); "Twenty-two of 27 referents… two move referents were combined" (1087) | OK |
| 3 | Study ran in 2008 | No study date given | UNVERIFIABLE |
| 3 | "importance of immediate usability is high" | 1085 | OK |
| 3 | Two of three authors at MSR | 1083 | OK |
| 4 | Quote "To date… reliable recognition" | 1083 (ellipsis drops the citation and "who personally employ them…"; acceptable) | OK |
| 4 | Malik 1-, 2-, 3-, 5-finger | 1085 | OK |
| 4 | Long et al.: "Designers are poor at picking…" | "users are sometimes poor at picking easily differentiable gestures" (1084) | WRONG (word and hedge) |
| 4 | "almost anything one can do with one's hands could be a potential gesture" | 1083 | OK |
| 5 | "principled, logical fashion" / "rarely so systematic" | 1083 | OK |
| 5 | Epps' screenshots "had pushed" people to emulate the mouse | "may have biased participants to simply emulate the mouse" (1084) | OVERSTATED |
| 5 | "monologue…always acceptable", "unrevised behavior" | 1085 | OK |
| 5 | Procedure and score from EdgeWrite ref. 33 | 1085, 1087 "(This process was adopted from prior work [33].)" | OK |
| 5 | Referent/sign from McNeill | 1085 and footnote 1 | OK |
| 6 | 9 women, mean age 43.2 (sd 15.6), right-handed | 1085 | OK |
| 6 | No tabletop or iPhone use | "No participant had used an interactive tabletop, Apple iPhone, or similar" (1085) | OK ("or similar" dropped) |
| 6 | Occupations | 1085 | OK |
| 6 | 24 × 18 in, 1024 × 768, recorded voice + animation, reset and wait | 1085–1086 | OK |
| 6 | Complexity rated by the three authors beforehand, 1–5, mean 2.70 | Table 1 (1086), 1085 | OK |
| 6 | Two 7-point ratings: good match, easy | 1086, 1088 | OK |
| 6 | Ovals with millisecond timestamps; two authors took notes | 1086 | OK |
| 6 | Authors designed their own sets beforehand | 1090 | OK |
| 7 | Finger count ignored in grouping | Not stated; only indirect support (1086 form coding, Figure 4 caption) | UNVERIFIABLE (presented as fact) |
| 7 | 0.42 worked example, groups 12/3/3/2 | Eq. 2 (1087); this is the 2-hand score | OK (label it 2-hand) |
| 7 | Everyone identical → 1.0; all different → 0.05 | Range [\|Pr\|⁻¹, 1] (1087) | OK |
| 7 | Conflict: larger group wins | 1087 | OK |
| 7 | The loser gets another gesture or is re-bound to a target | Not stated | UNVERIFIABLE |
| 7 | 57.0% coverage, conflict-free | 1087 | OK |
| 8 | Form 45 / 27 / 14 / 9 / 4 | Right hand .450/.274/.140/.092/.037; left hand .379/.343/.177/.037/.056 (Figure 2, 1087) | OK for right hand; label missing |
| 8 | 72% mouse-like | "about 72%" (1090) | OK |
| 8 | Nature 44 / 27 / 19 / 10 | .439/.272/.189/.100 (Figure 2); 43.9% (1091) | OK |
| 8 | Binding 70 / 18 / 9 | .702/.182/.089, plus mixed .027 (Figure 2) | OK (mixed omitted) |
| 8 | Flow 54 / 46 | .535/.465 (Figure 2) | OK |
| 8 | Simple → physical, complex → metaphorical/symbolic | 1087 | OK |
| 9 | Move a little/a lot at 1.0 | Figure 3, 1-hand (1088) | OK (1-hand only) |
| 9 | Duplicate, rotate, select group, insert, select single ≈0.35–0.55 | ≈.52, .48, .48, .48, .37 (Figure 3) | OK |
| 9 | After seven, at or below about 0.3 | Shrink ≈.31, then lower (Figure 3) | OK |
| 9 | Tail task switch → cut at 0.10–0.15 | ≈.22 down to ≈.12 (Figure 3) | WRONG |
| 9 | Two-hand above one-hand for select single, pan, enlarge, maximize | True (≈.46, .28, .51, .58); insert ≈.58 omitted (Figure 3) | OK (incomplete) |
| 9 | Two-hand dips near 0.1 for close, paste | Close ≈.18, paste ≈.13 (Figure 3) | PARTLY WRONG |
| 9 | A = 0.32 / 0.28 | 1087 | OK |
| 9 | r = −.52 | r = −.52, F1,25 = 9.51 (1087) | OK |
| 9 | Widget list: menu, help, undo, accept, reject, zoom | "delete, zoom in, zoom out, accept, reject, menu access, and help" (1091) | WRONG |
| 9, 15 | Discussion cites Figure 2 for the cutoff after seven | "Figure 2 indicates that agreement is low after the first seven referents" (1091) | OK (genuine slip in the paper) |
| 10 | 31 one-hand, 17 two-hand (35%) | 31 (64.6%), 17 (35.4%) (1088) | OK |
| 10 | Ten referents with one gesture, rest 2–5 | 10/4/3/4/1 (1088) | OK |
| 10 | Spread on object/background/folder = enlarge/zoom/open | 1088, Figure 4 | OK |
| 10 | Drag off-screen: delete; dialog: reject | Delete₁; Reject₂,₃ (Figure 4) | OK |
| 10 | "Hold-and-tap" reused for move, duplicate, delete, paste, minimize | "Hold and tap" = Select Group₁; the reused idiom is Move₂ "jump"; Duplicate = "tap source and destination" (Figure 4) | WRONG label |
| 10 | ✓ accept, ✗ reject, ? help, scribble undo | Figure 4 ("scratch out") | OK |
| 10 | Menu: drag in from the screen edge like a drawer | "Menu: pull out", drawn on an object (Figure 4) | UNVERIFIABLE / likely wrong |
| 10 | Insert, maximize reuse; close and task switch → imaginary widgets | 1087, 1090 | OK |
| 10 | 1–3 fingers per Figure 4 caption | 1089 | OK |
| 11, 12 | 43.5% / 60.9% / 19.1% | 1090 | OK |
| 11 | 13 of 20 varied fingers, 2 said it mattered; 1–3 point, palm more, 4 grey | 1090 | OK |
| 11 | One hand preferred for 25 of 27, two hands never | 1088 | OK |
| 11 | "I'm a child of the mouse" | 1090 | OK |
| 11 | Popular gestures rated better | χ² = 34.10 (1088) | OK |
| 11 | Reversible gestures although pairs never shown together | 1090 | OK |
| 11, 22 | Easy ≈3.4 s, hard ≈2.0 s | 1088 | OK |
| 12 | "thinking creatively" / "hopelessly lost" | 1090 | OK |
| 13 | Four contributions as listed | 1084 | OK |
| 13 | "an important next step" | 1091 | OK |
| 13 | Own list does not include the method | True of the list (1084); the abstract leads with "an approach" (1083) | MISLEADING |
| 13 | "first to employ users, rather than principles" | 1084 | OK |
| 13 | Good et al. 1984 typed commands; Epps 2006 Windows images | 1084 | OK |
| 13 | 216 studies, DIS 2020 | External (briefed as verified) | OK (not re-checked) |
| 14 | "useful not just for what it contains, but also for what it omits" | 1088 | OK |
| 14 | Neutral world "avoids" the Windows bias | Paper finds desktop bias persisted (1090) | CONTRADICTED |
| 15 | Acknowledged limits: revision, context, culture, literate American adults | 1091 | OK |
| 15 | Commands "came from desktop systems" | "desktop and tabletop systems" (1085) | PARTLY WRONG |
| 15 | No inter-rater reliability | None reported (1086, 1087) | OK |
| 15 | "ease of recognition" claimed | Conclusion (1091) | OK |
| 15 | Floor 0.05 | 1/\|Pr\|; ≈0.053 where one gesture was discarded | OK (approximately) |
| 16 | Select single 11 of 20 → 0.37 | Printed as Eq. 3, groups 11/3/3/3 (1087) | OK (it is the paper's number, not the presenters') |
| 16 | Guessability is "what the paper is named after" | Title is "User-Defined Gestures…" (1083) | WRONG |
| 17 | Two-hand gestures are 35% of the set | 1088 | OK |
| 17 | "Zoom in, that's the same as enlarge" | "Oh, that's the same as enlarge" (1090) | OK |
| 17 | Minimize became shrink | Minimize = drag to bottom (Figure 4); only mental model similar (1090) | WRONG |
| 17 | Figure 3's order differs from Table 1 | 1090 | OK (the causal "this is why" is overstated) |
| 18 | Section title "It's a Windows World" | 1090 | OK |
| 18 | Screening: no tabletop, no iPhone | "…or similar" (1085) | OK (partial) |
| 18 | Pinch and reverse pinch in the set | Enlarge (Shrink)₃ pinch (Figure 4); 5 used a reverse pinch (1091) | OK |
| 18 | Device dates | External; GPS "1980s" doubtful | UNVERIFIABLE / partly doubtful |
| 19 | Furnas ref. 8, Hutchins/Hollan/Norman ref. 10, Morris ref. 17, Wilson UIST 2008 ref. 32 | 1092 | OK |
| 19, 22 | 44% physical as support for the physics engine | "43.9%… the idea of using a physics engine [32]… has support" (1091) | OK |
| 20 | Authors' next steps: reverse study, recognizer | 1091 | OK |
| 20 | "Let participants see all commands first" | Stated as a limitation ("perhaps users would have performed differently…"), not as a next step (1091) | OK (loosely) |
| 20 | Children, other cultures, context | "worthy of investigation" (1091) | OK |
| 20 | Sense above the table / beyond the edge | 1091 | OK |
| 22 | Minimize "as defined by the user's seating position" | Figure 4 caption, Minimize₂ (1089) | OK |
| 22 | People gestured in the air and off-screen | 1090 | OK |

## Answers to the five judgment questions

**Q1. Is Part A too long or too summary-like?** YES. Slides 2–10 follow the paper's own section order: intro, related work, method, agreement, taxonomy, results, set. They take about 1,330 spoken words, roughly 9.5 minutes, a third of the talk, retelling content the room has read. Slides 3, 6 and 8 could each shrink to a single line, and slides 6 and 7 could merge.

**Q2. Does slide 18's argument hold up as stated?** NO. The underlying point holds and makes a good discussion topic: elicitation captures whatever repertoire people bring, as Morris et al. 2014 later named. But the slide attacks a claim the paper never makes (an "unbiased participant"). Its own table contradicts "every mass-market item teaches one point" (the Wii row). And it replaces one untested cause of the 72% with another.

**Q3. Is the slide 21 thesis defensible from the evidence in the deck?** PARTLY. "The lasting contribution is the method" is well supported. "Evidence weaker than its reputation" rests on three supports:
- The first is contradicted by the deck's own slide 20 (GI 2010).
- The second mislabels self-report as circularity.
- The third misreads "unrevised behaviour."

"Reputation" is also never shown on any slide.

**Q4. Is 23 slides in 28 minutes realistic at the current density?** NO. Many slides carry 180–310 words of on-screen text: 11, 13, 17, 18, 20 and 22, plus slide 6's table. The spoken script already fills 28–30 minutes before slide 19 is written. No one can read those slides while listening.

**Q5. Which slides to cut or merge to save 5 minutes?**
- Merge 11 and 12: keep 12 as the visual and cut the table to 4 rows (saves ~1:00).
- Fold slide 3 into one line on slide 4 (saves ~0:45).
- Merge slide 6 into slides 5 and 7 (saves ~1:00).
- Fold slide 16 into slide 15 or 17, and remove the repeated items from slide 15's "worth adding" column (saves ~1:15).
- Cut slide 22's "five smaller things" to two, and trim slide 18's table to four rows (saves ~1:00).

## Strengths of the deck

- Most numbers are correct, and page references are given throughout. The fact base is unusually careful for a student deck.
- Part B is organised by question rather than by paper section, which is what the brief asks for.
- Slide 16, agreement vs coverage vs guessability, is the deck's best idea. It is clear, correct and teachable. Fix only the "named after" slip.
- Slide 13 separates the contributions by kind and flags the borrowed procedure (EdgeWrite, Good et al.). That is honest novelty analysis.
- Slide 17's framing, "design decisions that shaped the result," is a strong critical lens once the unreported procedures are hedged.
- The live elicitation warm-up on slide 23 is an excellent bridge into the discussion.
- Catching the Figure 2 / Figure 3 slip on p. 1091 shows close reading.
