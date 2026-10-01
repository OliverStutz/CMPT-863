# Cross-examination of the two slide reviews (Opus C and Opus D)

2026-09-30. Two Claude Opus agents reviewed [[Week 4 - Slides]] and [[Week 4 - Speaker Notes]] against the paper PDF and the course brief, from an identical prompt, without seeing each other. Full reports: `opus-C-slides-review.md` (7 blockers, 14 should-fix, 11 nice-to-have) and `opus-D-slides-review.md` (8 blockers, 19 should-fix, 9 nice-to-have). Where they disagreed on a figure reading, Figure 3 was re-rendered at 220 dpi and read directly.

## Headline

Both reviewers reach the same verdict: the numbers are almost all right and Part B is real critique, so the deck is above "adequate" on the rubric and can go in as the draft submission. Neither considers it ready for class. Both name the same single biggest risk: slide 18 and the thesis on slide 21 attack a claim the paper does not make. The paper's "unrevised behavior" (p. 1085) means behaviour not revised by recognizer feedback during the session, not behaviour free of prior experience. The paper never promises an unbiased participant and reports the desktop bias itself as a finding (abstract p. 1083; "It's a Windows World," p. 1090).

## Votes on the five judgment questions

| Question | C | D | Vote |
|---|---|---|---|
| Q1 Part A too long or summary-like for a room that read the paper | YES | YES | 2 of 2 |
| Q2 Slide 18 holds up as stated | NO | NO | 2 of 2 |
| Q3 Slide 21 thesis defensible from the deck's evidence | PARTLY | PARTLY | 2 of 2 |
| Q4 23 slides in 28 minutes realistic at current density | NO | NO | 2 of 2 |
| Q5 What to cut for 5 minutes | fold 6 into 5; shrink 8 into 9; merge 12 into 11; fold 16 into 15/17; drop slide 22's five small items; trim 14 | merge 11+12; fold 3 into 4; merge 6 into 5 and 7; fold 16 into 15/17; trim lists on 22 and the table on 18 | Agreed on 6→5, 12→11, 16→15/17, trim 22 |

Both estimate about 3,500 words of on-slide text and about 3,500 spoken words before slide 19 is written, giving a 3 to 5 minute overrun. Slide 18's script (about 308 words for 2:00) is the worst offender in both reports.

## Findings raised by both reviewers (2 of 2)

Fix these first. Page numbers are the printed ones.

**The critical spine**
1. **Slides 18, 21: "unrevised behaviour" misread** (C blocker 1, D blocker 1). Reframe: the paper treats prior experience as a finding about users; we argue it is a property of the method, and the source is broader than the desktop. Drop "assumes a participant without a touch history."
2. **Slide 18: screening is quoted incompletely** (C blocker 2; D uses the same passage). p. 1084: "non-technical people without prior experience using touch screens (e.g., the Apple iPhone)"; p. 1085: "interactive tabletop, Apple iPhone, or similar." Quoting this verbatim makes the critique sharper: "no prior touch-screen experience" is the implausible claim for 2008 adults, and no screening instrument is given.
3. **Slide 18: the argument contradicts itself and overreaches** (C blocker 3, D blocker 7). The Wii row ("whole-hand motion") contradicts "every mass-market item teaches one finger, one point." The "sum of all single-point interfaces" cannot be told apart from the desktop explanation with this data, and it ignores desktop-specific think-aloud evidence (two-button mouse clicking, imaginary taskbar, "Anything I can do that mimics Windows," p. 1090) and the plain-pointing explanation the paper cites from Beringer, Voida and Epps (p. 1084). Pinch as evidence of iPhone exposure is weak: one alias of four, and only 5 people used a reverse pinch (p. 1091). Present as a question ("Which legacy: mouse, touch screens, or pointing itself?"), trim the table to four rows, drop the BMW footnote.
4. **Slides 13, 21 against slide 20: "never validated" vs GI 2010** (C blocker 6, D blocker 2). GI 2010 is a preference validation by 22 blind judges. Say "the guessability validation the authors proposed was never run; GI 2010 tested preference instead." D adds that slides 12 and 14 ("tests its own premise") also cut against the thesis.
5. **Slide 13: "the authors' list of four does not include the method" misleads** (C blocker 6b, D should-fix 5). The abstract leads with "We present an approach…" (p. 1083) and the procedure is credited to EdgeWrite 2005 (pp. 1085, 1087). Say the numbered list omits it because they credit it to 2005.

**Factual errors on screen**
6. **Slide 9: tail values wrong** (C blocker 4, D blocker 4). One-hand task switch ≈ 0.22 down to cut ≈ 0.11, not 0.10–0.15. Confirmed on the re-render: task switch .21, zoom out .20, previous .19, undo .18, accept .17, help .16, reject .15, next .13, cut .11.
7. **Slide 9: widget list wrong** (C blocker 4, D blocker 3). Paper: "delete, zoom in, zoom out, accept, reject, menu access, and help" (p. 1091). Undo is not on it, delete is. It is also a separate recommendation grounded in imaginary-widget use, distinct from "beyond the first seven referents, widget as well as gesture."
8. **Slide 9: two-hand line misdescribed** (C blocker 4, D should-fix 13). Insert (≈ 0.58) is one of the two highest two-hand points and is missing. Close dips to ≈ 0.19, not 0.1.
9. **Slide 16: "what the paper is named after" is wrong** (C blocker 5, D blocker 5). Guessability names the method (p. 1083). Say "what the method is named after."
10. **Slide 16 and notes: the 0.37 is printed in the paper** (C blocker 5, D should-fix 1). Eq. 3, p. 1087, groups 11, 3, 3, 3 for one-hand select single. The speaker note calling it "our estimate" is false. Cite the equation; it strengthens the slide.
11. **Slide 17 #4: minimize did not become shrink** (C blocker 7, D should-fix 4). Minimize has its own gestures in Figure 4 (drag to bottom; Move2 to bottom). Only the mental models were paired (p. 1090). Maximize did fold into enlarge (p. 1087). Soften "this is why Figure 3's order differs."
12. **Slide 10: idiom labels do not match Figure 4** (C nice-to-have 2, D should-fix 3). "Hold and tap" is the paper's label for Select Group1. The reused two-hand idiom is Move2 "jump," reused by Delete2, Paste3, Minimize2. Duplicate is "tap source and destination."
13. **Slide 4 and 19: Long et al. is misquoted** (C should-fix 1, D should-fix 10). Paper: "users are sometimes poor at picking easily differentiable gestures" (p. 1084). Keep "sometimes."
14. **Slide 15: commands "came from desktop systems"** (C should-fix 7, D should-fix 11). Paper: "desktop and tabletop systems" (p. 1085). Say "several are desktop concepts."
15. **Slide 8: Form shares are right-hand only** (C should-fix 8, D should-fix 12). Figure 2 splits Form by hand. Label "right hand"; add mixed binding (.027).
16. **Speaker note slide 2: four commands, not five, got no gesture** (C should-fix 14, D should-fix 14). Insert, maximize, task switch, close (p. 1087). Strictly 23 referents have gestures with the two moves merged.
17. **Slide 3: "2008 · this study runs" is not in the paper** (C nice-to-have 11, D nice-to-have 1). Say "by 2008."

**Unreported procedures presented as fact**
18. **Slides 7, 11, 17: "finger count ignored in grouping"** (C should-fix 3, D blocker 6). p. 1087 says only "each group held identical gestures." The 1–3 finger equivalence is in the taxonomy coding (p. 1086) and the Figure 4 caption. Phrase as inference: "the paper does not say what counted as identical; the caption implies 1–3 fingers were merged, and that choice raises A."
19. **Slides 7, 17: "the loser is rescued by re-binding"** (C should-fix 4, D blocker 6). The paper says only "the referent with the largest group won" (p. 1087). Label as our reading; drop "rescued."

**Overclaims in the critique**
20. **"Circular" is the wrong word** (C should-fix 5, D should-fix 7). Participants did not know group sizes when rating (p. 1088). The real issues: self-rated fit, no independent judges, and the authors' "goes a long way to validating" overstates it. Both reviewers also flag the missing statistics critique: all χ² tests use N = 1,074 gestures from 20 people as independent observations (pp. 1087–1088).
21. **Slide 12: "could not predict what users would do" is too strong** (C should-fix 6, D should-fix 8). Designers are scored against the users' set by construction; 60.9% combined is substantial overlap; the paper itself says three experts cannot match the scope of 20 participants. Say "predicted under half individually, about 60% together; the comparison measures overlap, not quality."
22. **Slides 5, 14: Epps et al. overstated; "avoids the Windows bias" contradicted** (C should-fix 2, D should-fix 9). Epps "may have biased" (p. 1084); the desktop got in anyway (p. 1090). Make it a strength-and-limitation pair.
23. **Slide 15: "two authors with a commercial stake" is loaded** (C should-fix 7, D should-fix 19). Say "two authors at Microsoft Research, run on their own prototype; the paper does not discuss this."
24. **Slides 13, 20: "nobody shipped it," "the set went nowhere" unverifiable** (C blocker 6c, D should-fix 6). Tap, drag, double-tap, pinch, spread, drag-off-screen are universal. Say "no product adopted the set as published; many of its members are now standard, which cuts both ways." Slide 20's pinch irony weakens: pinch was elicited and is in the set.
25. **Slide 20: "named by the authors" mixes sources** (C should-fix 9; D fact table). Reverse study and recognizer are the next steps; the rest are limitations and technology implications (p. 1091).

**Structure, timing, redundancy**
26. **Part A follows the paper's section order** (C should-fix 10, D Q1). Slides 6, 8, 10 are recap; 3, 5, 7 earn their time.
27. **Redundancy** (C should-fix 12, D should-fix 17): designer numbers on 11 and 12; grouping on 11, 15, 17; self-rating on 11, 15, 21, 22; floor 0.05 on 7, 15, 16; "never validated" on 13, 15, 16, 21; 3.4 s on 11 and 22; 44% physical on 8, 19, 22; Figure 2/3 slip on 9 and 15. Each point once, at its strongest place.
28. **Slide 19 is the only rubric question unanswered** (C should-fix 13, D should-fix 18). Needs explicit links to weeks 1–3 readings, each with a "so what."
29. **Slide 6 "Control" is a mislabel** (C nice-to-have 1, D nice-to-have 3). "Designer baseline." Both also note the paper does not state the one-hand-then-two-hand order (p. 1086).
30. **Slides 7 and 9: move a little carries two unlabeled values** (C nice-to-have 3, D should-fix 2). Label 0.42 as two-hand and 1.0 as one-hand on the slides.
31. **Missing fairness point** (C nice-to-have 4, D should-fix 16): the authors' complexity ratings correlate r = .71 with planning time (8 s vs 15 s, p. 1088), which partly answers "complexity was rated by the authors." Say so on slide 11.
32. **Speaker note slide 12 mentions the Opus reviewers** (C should-fix 6, D should-fix 15). Delete.

## Findings raised by one reviewer (1 of 2)

C only:
- Slide 4 tagline "Nobody had asked what people would actually do" is contradicted by the paper's own related work (Good, Nielsen, Epps, Beringer, Voida, p. 1084). Say "few had asked, and never for a full multi-touch set." Reasonable; adopt.
- Slide 17: the quote is "Oh, that's the same as enlarge," said when prompted to zoom in (p. 1090), not "Zoom in, that's the same as enlarge." Adopt.
- Speaker note slide 18: "the same group, Morris and colleagues" is loose; the 2014 paper shares Morris and Wobbrock, not Wilson. Adopt.
- Slide 15: the paper's own words are "children, Eastern, or uneducated participants" (p. 1091), which is itself discussable. Optional.
- Coverage denominator is unstated (1,074?); slide 22 #4 "users adapt within minutes" has no source. Optional.

D only:
- Slide 10: Figure 4's menu gesture is labelled only "pull out" and the drawing starts on an object, not at the screen edge. **Verified on the rendered figure 2026-09-30: D is right.** One hand holds the object, a finger of the other drags out from it. Adopted.
- Slide 18: consumer touch-screen GPS units date from the mid-2000s, so "1980s" is wrong for GPS. Adopt or drop the row.
- Slide 17 #3: Morris et al. 2014 "production" means several proposals per referent; forcing a two-hand variant is related, not the same. Say "related to."
- Slide 13: Nielsen et al. ("a similar approach," p. 1084) is on the notes but not the slide, and bears on "first to employ users" more directly than Good. Adopt.
- Slide 11 #6 cross-reference should point to slide 15, not 16. Adopt.
- Floor is ≈ 0.05: with 6 discards some referents have |Pr| = 19. Trivial.
- Commentary markers exist only on slide 23; add visible "[classmate]" slots on 15, 18, 22 so the integration points are visible to the grader. Adopt.
- Speaker note slide 13: "216 studies descended from this one" overclaims; the review covered 216 studies. Adopt.

## Disagreement resolved by re-reading Figure 3

C read two-hand *paste* at ≈ 0.23 (above one-hand); D read it at ≈ 0.13. The 220 dpi render shows the ≈ 0.13 dip at *task switch*, with paste at ≈ 0.22, slightly above the one-hand line. **C is right.** Two-hand dips are close (≈ 0.19) and task switch (≈ 0.13). The render also shows two-hand *pan* at ≈ 0.19, below one-hand (≈ 0.25), so slide 9's "above one-hand for select single, pan, enlarge, maximize" should become "peaks at insert, maximize, enlarge; roughly equal at select single." Neither reviewer caught the pan error; it came from the original figure note, now corrected.

## What both reviewers say to keep

- The fact base: every headline number and quote checks out (1,080; 48/22; 31/17; 57.0%; 0.32/0.28; r = −.52; 0.42; 43.5/60.9/19.1; 25 of 27; 13 of 20; 72%; 44%; 3.4 s vs 2.0 s).
- Slide 16, agreement vs coverage vs guessability: "the deck's best idea" (D), "the best slide" (C).
- Slide 13's novel-vs-borrowed box.
- Slide 17's framing of design decisions, once hedged.
- The live elicitation warm-up on slide 23.
- Catching the Figure 2/3 slip and the exponent in the range of A.

## Recommended fix order (applied 2026-09-30, revision 2 of both files)

1. Rewrite slides 18 and 21 and their scripts (items 1–5). Retitle 18, e.g. "Whose habits did the study capture?"
2. Correct slide 9 (items 6–8 and the pan/paste reading).
3. Correct slide 16 and its note (items 9–10), slide 17 #4 (item 11), slide 10 idioms (item 12).
4. Hedge the grouping and re-binding claims on slides 7, 11, 17 (items 18–19).
5. Wording pass: items 13–15, 20–25, plus the C-only and D-only adoptions above.
6. Add the N = 1,074 non-independence point to slide 15, the r = .71 caveat to slide 11, one design-implications line to slide 10.
7. Cut for time: fold 6 into 5, 12 into 11, 16 into 15/17; trim slide 22 to two small items; trim slides 14 and 18. Target about 2,900 spoken words.
8. Speaker-note fixes: four not five (slide 2), remove the Opus mention (12), 2014 authorship (18), "216 studies reviewed" (13).
9. Fill slide 19; add commentary slots on 15, 18, 22.
