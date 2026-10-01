# Opus report C: slides review
Independent review, 2026-09-30, of "Week 4 - Slides.md" against the paper PDF and the course brief. Written without sight of report D.

## Verdict

This is a strong draft and nearly ready. The arithmetic is almost all right: every headline number checks out against the paper (1,080 gestures, 48/22, 31/17, 57.0%, A = 0.32/0.28, r = −.52, 43.5/60.9/19.1%, 25 of 27, 13 of 20, 72%, 44%, 3.4 s vs 2.0 s), and so does every quote. The deck is also more critical than most seminar decks. The problems are in two places. First, the readings of Figure 3 on slide 9 have several wrong values, and the widget list has the wrong commands. Second, and more important, the presenters' own critical spine (slides 13, 18 and 21) rests on readings of the paper that a classmate who read pages 1083–1085 can knock down. The single biggest risk is slide 18 together with the third support of slide 21. The paper's "unrevised behavior" means behaviour not revised by recognizer feedback (p. 1085), not behaviour free of prior experience. The paper never claims an unbiased participant, and its abstract says outright that "desktop idioms strongly influence users' mental models" (p. 1083). Slide 18 also understates the paper's screening claim: p. 1084 says "without prior experience using touch screens," and p. 1085 says "…Apple iPhone, or similar." So the highlight slide currently attacks a claim the paper does not make, and it misses the claim the paper does make, which is the weaker one. The deck can go in as a draft once the blockers below are fixed. All of them are wording fixes, and none needs restructuring. Part A should also be cut by about 3 minutes.

## Blockers (must fix before class)

1. **Slides 18 and 21 (third support): "unrevised behaviour" is misread.** Slide 21 says: "'Unrevised behaviour' assumes a participant without a touch history." The speaker notes for slide 18 say: "The paper's whole design rests on capturing unrevised behaviour from hands that no touch system has trained." This is wrong. The paper defines the term through feedback within the dialogue: "Feedback… causing the user to revise his or her mental model… we sought to remove the gulf of execution… a monologue in which the user's behavior is always acceptable. This enables us to observe users' unrevised behavior" (p. 1085). "Unrevised" means not revised by recognizer feedback during the session. It says nothing about prior experience. The paper also reports legacy influence as a headline finding (abstract, p. 1083; "It's a Windows World," p. 1090), so "an unbiased participant did not exist" is not something the authors denied. **Fix:** make the argument about the screening claim and the method's ceiling, not about "unrevised". For example: "Elicitation measures the participant's existing repertoire. The paper treats that as a finding about users when it is partly a property of the method."

2. **Slide 18: the screening claim is understated.** The slide says: "Screening: no tabletop, no iPhone. **No touch history was recorded.**" The paper says more. p. 1084: "we explicitly recruited non-technical people without prior experience using touch screens (e.g., the Apple iPhone)". p. 1085: "No participant had used an interactive tabletop, Apple iPhone, or similar." Classmates who read p. 1084 will raise this. The accurate quote makes the critique sharper: "no prior experience using touch screens" is implausible for 2008 adults given ATMs, kiosks and point-of-sale terminals, and the paper gives no screening instrument. **Fix:** quote p. 1084 verbatim and attack that claim.

3. **Slide 18: the argument does not hold as stated.**
   - (a) It contradicts itself. The bullet says "Every mass-market item teaches **one finger, one point**", but the table's Wii row says "Whole-hand motion".
   - (b) It conflates touch screens with touch-sensitive and motion devices. Touchpads, the iPod wheel and the Wii are not touch screens, so the table proves less than "had already touched" implies.
   - (c) The table is speculation about a composite "43-year-old". The participants' ages had sd 15.6 (p. 1085), and no one knows what these 20 people owned.
   - (d) The alternative explanation ("the sum of all single-point interfaces") cannot be told apart from the authors' explanation with this data, and it ignores evidence that only the desktop explains: clicking with index and middle fingers "as if… using a two-button mouse", an imaginary taskbar button for task switch, an imaginary top-right button for close, and "Anything I can do that mimics Windows" (p. 1090).
   - (e) It ignores a third explanation the paper itself raises. Single-finger pointing is the default human deictic gesture. Beringer found "most gestures involved pointing… a finding we reinforce here," Voida et al. found "people overwhelming[ly] used finger-pointing," and Epps et al. found the index finger most common (p. 1084).
   - (f) "The final set contains a pinch and a reverse pinch" as evidence of iPhone exposure is weak. Pinch is one of four aliases for enlarge (Figure 4), pulling apart is also a physical stretching action, and only 1 of 5 people used a diagonal reverse pinch (p. 1091).
   - (g) "In 2008" is not in the paper. The paper never dates the study.

   **Fix:** reframe as a question, roughly "Which legacy: mouse, touch screens, or pointing itself?" Keep 3–4 device rows, drop the "every item" sentence, and drop the BMW footnote.

4. **Slide 9: the Figure 3 readings and the widget list are wrong.** Values read from Figure 3 rendered at 400 dpi (p. 1088):
   - "The tail, *task switch* to *cut*, sits at 0.10–0.15." In fact, one-hand task switch ≈ 0.22, zoom out ≈ 0.20, previous and undo ≈ 0.19, accept ≈ 0.17, help ≈ 0.16, reject ≈ 0.15, next ≈ 0.14, cut ≈ 0.11. Only the last three are in 0.10–0.15.
   - "Two-hand… dips near 0.1 for *close* and *paste*." In fact, two-hand close ≈ 0.19, and paste ≈ 0.23, which is *above* one-hand. The real two-hand dips are task switch (≈ 0.14), close (≈ 0.19) and menu access (≈ 0.20). Nothing reaches 0.1.
   - "Above one-hand for *select single, pan, enlarge, maximize*" leaves out **insert** (two-hand ≈ 0.58 vs ≈ 0.48), which is one of the two highest two-hand peaks.
   - The widget list "(menu, help, undo, accept, reject, zoom)" is wrong. The paper says "delete, zoom in, zoom out, accept, reject, menu access, and help" (p. 1091): undo is not on it and delete is. The paper also grounds that list in imaginary-widget use ("enough participants acted on imaginary widgets…"), not only in low agreement.

   **Fix:** "tail ≈ 0.11–0.22; two-hand peaks at insert, maximize, enlarge; dips at task switch, close, menu access", and copy the widget list verbatim.

5. **Slide 16: wrong title claim, and wrong sourcing of 0.37.**
   - The slide says guessability "is what the paper is named after". The paper is titled "User-Defined Gestures for Surface Computing". Guessability is the name of the *method* ("a guessability study methodology [33]", p. 1083) and an author keyword.
   - The "11 of 20 → A ≈ 0.37" example is not the presenters' calculation. The paper prints it: Eq. 3, p. 1087, "select single (1-hand)… groups of size 11, 3, 3, and 3… = 0.37". The speaker notes say "our estimate… not a number printed in the paper", which is false.
   - "Tapped" is not stated for the 11-person group; it is only implied by Figure 4's "Select Single1: tap".

   **Fix:** "the paper's own second worked example (Eq. 3, p. 1087)" and "the term the method is named after". Using the paper's own example strengthens the point.

6. **Slides 13 and 21: the thesis supports are contradicted elsewhere in the deck and in the paper.**
   - (a) "The set was never validated" (slides 13, 21) clashes with slide 20, where the same authors' GI 2010 study had 22 blind judges rate 81 gestures. That is a preference validation, even though it is not the guessability reversal proposed on p. 1091.
   - (b) "The authors' own list of four does not include the method" is literally true of the numbered list (p. 1084), but it misleads. The abstract's first claim is "We present an approach to designing tabletop gestures that relies on eliciting gestures…" (p. 1083), and the paper says the result "reflects well on our methodology" (p. 1090). The authors also did not list the method because it was theirs from 2005 (ref. 33, "adopted from prior work", p. 1087). Slide 13 says this itself, which then undercuts the thesis that "the lasting contribution is the method."
   - (c) "Nobody shipped it" and "the set went nowhere" (slide 20) are negative claims nobody verified. Drag, double-tap, pinch and splay are in shipped systems, whether or not that is causal.

   **Fix:** "The guessability validation the authors proposed was never run; GI 2010 tested preference instead." "This paper turned a 2005 four-page method into a genre." Drop "does not include the method" or qualify it with the abstract.

7. **Slide 17 #4: wrong fact and a misquote.**
   - "Minimize became shrink." In the final set, minimize got its own gestures: "Minimize1: drag to bottom of surface" and "Minimize2: use Move2…" (Figure 4, p. 1089). The paper only says that *mental models* paired shrink and minimize (p. 1090). Maximize did fold into enlarging (p. 1087).
   - "'Zoom in, that's the same as enlarge.'" The quote is "Oh, that's the same as enlarge," said "upon being prompted to zoom in" (p. 1090).
   - "This is why Figure 3's order differs from Table 1" overstates what the paper does. The paper only puts the two observations side by side (p. 1090).

## Should fix

1. **Slide 4: the Long et al. attribution, and "Nobody had asked."**
   - The slide says "Designers are poor at picking gestures that are easy to tell apart (p. 1084)". The paper says "users are sometimes poor at picking easily differentiable gestures" (p. 1084). Long's participants were users of a design tool, so "designers" is defensible from the original paper but not from this one. Keep "sometimes" either way.
   - The tagline "Nobody had asked what people would actually do" is contradicted by the paper's own related work: Good 1984, Nielsen 2004, Epps 2006, Beringer, Voida, Mignot (p. 1084), and by slide 13's novelty box.
   - **Fix:** "users, sometimes"; "Few had asked, and never for a full multi-touch set."

2. **Slides 5 and 14: Epps et al. is overstated.** "Epps et al.'s Windows screenshots had pushed people toward mouse emulation" and "Avoids the Windows bias that undermined Epps et al." The paper says Epps "acknowledged that their Windows-based prompts **may** have biased participants" (p. 1084). Slide 14 saying the bias was "avoided" also contradicts slides 11 and 18 (72% mouse-like anyway, "It's a Windows World"). **Fix:** "removes Windows imagery; the desktop got in anyway (p. 1090)". That makes a better strength-and-limitation pair.

3. **Slides 7, 11, 17: "finger count is ignored" in grouping is not in the paper.** The grouping description says only "such that each group held identical gestures" (p. 1087). The 1–3-finger equivalence appears in the taxonomy coding (p. 1086) and in the Figure 4 caption, not as a grouping rule. Slide 17 #1 builds its argument on this, and the speaker notes cite "p. 1087 (grouping)". **Fix:** turn it into a sharper critique: "The paper never says what counted as 'identical'. The Figure 4 caption implies 1–3 fingers were merged, and that choice raises A." That is stronger and true.

4. **Slides 7 (step 3) and 17 #2: the conflict-resolution rule is extended beyond the paper.** The paper's rule is only "the referent with the largest group won the gesture" (p. 1087). Rescue by re-binding to a different target is the presenters' reconstruction from the reuse and disambiguation passages (p. 1088, 1090, 1091). **Fix:** label it "our reading" on slide 7, or move it to slide 17 only.

5. **Slides 11 #6, 15, 21, 22: "circular" is overstated.** Participants rated their own gesture without knowing how many others had made it, so popularity and rating are independently measured: one comes from other people's choices, the other from one's own rating. The real weaknesses are different:
   - (a) self-rated "good match" is not other people's guessability;
   - (b) the χ² tests use N = 1,074 gestures as if they were independent, although they are nested in 20 people (pp. 1087–1088). That is the critique a statistics-minded instructor would make, and it is missing from the deck.

   **Fix:** replace "circular" with "self-rated fit, with non-independent observations".

6. **Slide 12 overclaims.**
   - "Could not predict what users would do" and "the paper testing its own premise" (notes: "the premise holds"). The test scores designers against a benchmark the users define, so designers lose by construction. The sizes of the author sets are not reported, so 43.5% cannot be read as a hit rate. 19.1% "never tried" is not shown to be worse; the authors' own line "thinking creatively or hopelessly lost" leaves it open.
   - Speaker notes, slide 12: "Both Opus reviewers we asked, independently, picked this…" should not be said in class.

   **Fix:** "Three expert designers covered at most 61% of what 20 novices produced."

7. **Slide 15: three wording problems.**
   - "The 27 commands came from desktop systems." The paper says "from desktop and tabletop systems" (p. 1085). Say "many".
   - "Two authors with a commercial stake" is loaded and does not say how a stake would bias *these* results. Say "two authors at Microsoft Research, on a Microsoft device".
   - "Literate American adults only; children and other cultures". The paper's own words are "children, Eastern, or uneducated participants would behave differently" (p. 1091). The original phrasing is itself worth discussing.

8. **Slide 8: Form shares are per hand, and "mouse-like" counts bimanual gestures.**
   - Figure 2 splits Form by hand (caption, p. 1087). The slide's Form figures are the right-hand column, not "all 1,080 gestures". The left-hand column, which exists only for two-hand gestures, is .379 + .343 ≈ 72% too.
   - So "72% mouse-like" (p. 1090) counts each index finger of a two-handed pull-apart as a "one-point path". The label is a taxonomy convention, not evidence of mouse habit. The deck tags this as "Robust number" (slide 11 #4), and this weakens that tag. It should also feed slide 18.
   - **Fix:** say this in one line on slide 8 or slide 18.

9. **Slide 20: the authors' next steps are mislabelled, and the irony is weaker than stated.**
   - "Named by the authors (p. 1091)" mixes three sources. The real next steps are the reverse study and the recognizer. "See all commands first" and "context, children, other cultures" are limitations judged "beyond the scope". Sensing above the table and beyond the edge come from "Implications for Surface Technology".
   - "Pinch-to-zoom became universal because Apple shipped it, not because anyone elicited it": pinch *was* elicited and is in the set (Figure 4). Elicitation and the market converged, which is a more interesting point.

10. **Part A is the paper's section order in miniature.** Intro (2–4), method (5–7), taxonomy (8), agreement (9), set (10) follow the paper's sections, at about 10 min, a third of the talk. Slides 6, 8 and 10 are pure recap for a room that has read the paper. See Q1 and Q5.

11. **Density and timing (see Q4).** On-slide text is about 3,500 words. Slides 18 (247 words), 11 (240), 20 (211), 22 (207), 6 (197), 13 (177), 5 (173), 15 (168), 17 (166) cannot be read while someone talks over them. The spoken script for slide 18 is 307 words in 2:00 (154 wpm), and slide 10 runs at 157 wpm. Neither allows for two-speaker handovers.

12. **Redundancy across slides.**
   - Designer comparison: 11 #1 and 12.
   - Grouping and IRR: 11 #2, 15, 17 #1.
   - Forced two hands: 11 #3, 17 #3.
   - Self-rating: 11 #6, 15, 21, 22 #3.
   - "Not a percentage / floor 0.05": 7, 15, 16.
   - "Never validated": 13, 15, 16, 21.
   - 3.4 s vs 2.0 s: 11, 22.
   - 44% physical: 8, 19, 22.
   - Figure 2/3 slip: 9, 15.
   - Smartphone-vocabulary question: 18, 20, 21.

   Each point should be made once, at its strongest place.

13. **Slide 19 (rubric item "connections to topics we have discussed") is empty.** It currently connects only to the paper's own references. This is a rubric question, not decoration, and it is the only one of the seven with no answer yet.

14. **Speaker notes have factual slips.**
   - Slide 2: "Five commands got no gesture of their own." It is four: insert, maximize, task switch, close (p. 1087). 22 of 27 counts the two moves as one.
   - Slide 3: "sold it from 2008, which is when this study ran." The paper gives no date.
   - Slide 16: the 0.37 note (see Blocker 5).
   - Slide 18: "the same group, Morris and colleagues". The 2014 paper shares Morris and Wobbrock, not Wilson.

## Nice to have

1. **Slide 6:** "Control" is the wrong word. The authors' own sets are a designer comparison, not a control condition (p. 1090). Rename it.
2. **Slide 10, hold-and-tap:** Figure 4 names "hold and tap" for **Select Group1**, which is missing from the slide's list. Move2 is "jump", and duplicate is "tap source and destination". Add select group, and say "Move2 (jump)" is what delete, paste and minimize reuse.
3. **Slide 7:** label the 0.42 as the two-hand score on the slide itself. One-hand move a little is 1.0 on Figure 3, and the room will notice.
4. **Slide 9:** say that r = −.52 is complexity against agreement (p. 1087). Consider adding the companion results: planning time r = .71 (8 s vs 15 s) and goodness r = −.59 (p. 1088). Both strengthen "complexity drives everything" and add a caveat, since all three use the authors' own complexity ratings.
5. **Slide 8:** Binding omits "mixed dependencies" (.027), and Flow .465 is rounded to 46. This is trivial, but the slide claims "shares of all".
6. **Slide 11 #2:** the paper says "5-finger touches or touches with the whole palm were something more" (p. 1090). Add "5 fingers".
7. **Slide 18:** cut the BMW footnote. It answers a question no one asked.
8. **Slide 22 #4:** "Real users adapt to feedback within minutes" has no source. Cite one or soften it.
9. **Missing methodological points** a critical reader may raise:
   - the one-hand gesture always came before the two-hand one, so the order was fixed (p. 1086);
   - how one-hand and two-hand groups were combined into a single set is not explained;
   - the coverage denominator is unstated (1,074?);
   - ordinal Likert data are analysed with χ² / logistic models over gestures (p. 1088).

   One line on slide 15 would show statistical literacy.
10. **Slide 2:** "covering 22 of the 27 commands" uses the paper's phrase. Strictly, 23 referents received gestures, with move a little and move a lot merged. Worth getting right because the speaker notes already get it wrong.
11. **Slide 3:** mark the study year as inferred ("2008, likely").

## Fact check table

| Slide | Claim | Paper says (printed page) | Status |
|---|---|---|---|
| 1 | Authors and affiliations; CHI 2009 Boston pp. 1083–1092 | Header and footer (1083–1092) | OK |
| 2 | 20 non-technical adults, Surface prototype, 27 commands | p. 1083, 1085 | OK |
| 2 | One hand then two, think aloud, no recognizer | p. 1085–1086 | OK |
| 2 | 1,080 gestures, logged, four video angles | p. 1086 (6 discarded) | OK |
| 2, 7, 10 | 48 gestures covering 22 of 27 commands | p. 1087–1088 ("Twenty-two of 27… two move referents were combined"; 48 gestures) | OK (strictly 23 referents; see NH 10) |
| 3 | "the importance of immediate usability is high" in public settings | p. 1085 | OK |
| 3 | Two of three authors at MSR; Surface prototype | p. 1083, 1085 | OK |
| 3 | Study ran in 2008 | Not stated | UNVERIFIABLE |
| 3 | Han 2005–06, iPhone and Surface announced 2007, Surface sold 2008, Windows 7 Oct 2009 | General knowledge; consistent | OK (outside paper) |
| 4 | Quote "To date, most surface gestures… reliable recognition." | p. 1083 (ellipsis drops "who personally employ them…" and "Despite skillful design") | OK (fair ellipsis) |
| 4 | Malik et al. 1-, 2-, 3-, 5-finger | p. 1085 | OK |
| 4 | Long et al.: "Designers are poor at picking…" | "users are sometimes poor at picking easily differentiable gestures" (p. 1084) | WRONG (wording) |
| 4 | "almost anything one can do with one's hands could be a potential gesture" | p. 1083 | OK |
| 4 | Furnas ref. 8, vocabulary problem | Ref. 8; cited for aliasing (p. 1088, 1091) | OK (presenters' framing) |
| 4 | "Nobody had asked what people would actually do" | Related work lists prior elicitation (p. 1084) | WRONG (overclaim) |
| 5 | "in a principled, logical fashion" / "user behavior is rarely so systematic" | p. 1083 | OK |
| 5 | Epps et al. Windows prompts "had pushed" people to mouse emulation | "may have biased" (p. 1084) | WRONG (overstated) |
| 5 | "monologue… always acceptable", "unrevised behavior" | p. 1085 | OK (quotes); slides 18 and 21 misread the meaning |
| 5 | Procedure and score from EdgeWrite 2005, ref. 33 | p. 1083, 1085, 1087 | OK |
| 5 | "Referent" and "sign" from McNeill | p. 1085 and footnote 1 | OK |
| 6 | 20 paid, 9 women, mean age 43.2 (sd 15.6), all right-handed | p. 1085 | OK |
| 6 | No tabletop or iPhone use | "interactive tabletop, Apple iPhone, or similar" (p. 1085); "without prior experience using touch screens" (p. 1084) | OK but incomplete (matters for slide 18) |
| 6 | Occupations | p. 1085 (also public affairs consultant) | OK |
| 6 | 24 × 18 in, 1024 × 768; voice and animation, reset, wait | p. 1085–1086 | OK |
| 6 | Commands from desktop and tabletop; authors rated complexity 1–5; mean 2.70 | p. 1085, Table 1 p. 1086 | OK |
| 6 | Random order; 1H then 2H; preference; two 7-point scales | p. 1086, 1088 | OK |
| 6 | 20×27×2 = 1,080, 6 discarded; ovals with ms timestamps; four angles; two authors took notes | p. 1086 | OK |
| 6 | Authors designed their own sets before the study ("Control") | p. 1090 | OK (label inaccurate) |
| 7 | Finger count ignored in grouping | Not stated in the grouping description (p. 1087); 1–3 fingers in taxonomy (p. 1086) and Figure 4 caption | UNVERIFIABLE |
| 7 | A = Σ(group/20)², averaged over referents | Eq. 1, p. 1087 | OK |
| 7 | Move a little, two hands: 12, 3, 3, 2 → 0.42 | Eq. 2, p. 1087 | OK |
| 7 | All identical → 1.0; all different → 0.05 | Range [1/\|Pr\|, 1], p. 1087 | OK |
| 7 | Largest group wins; loser re-bound to a different target | Only "the referent with the largest group won" (p. 1087) | PARTLY (re-binding is the presenters' inference) |
| 7 | 57.0% coverage, conflict-free | p. 1087 | OK |
| 8 | Four dimensions and categories | Table 2, p. 1086 | OK (mixed dependencies omitted) |
| 8 | Form 45 + 27 = 72; 14 / 9 / 4 | Figure 2 R-hand .450, .274, .140, .092, .037 (p. 1087) | OK as R-hand; WRONG to call it "all 1,080" |
| 8 | Nature 44 / 27 / 19 / 10 | .439 / .272 / .189 / .100 | OK |
| 8 | Binding 70 / 18 / 9 | .702 / .182 / .089 (+ .027 mixed) | OK |
| 8 | Flow 54 / 46 | .535 / .465 | OK |
| 8 | Simple → physical; complex → metaphorical or symbolic | p. 1087 | OK |
| 9 | Move a little and move a lot at 1.0 (1H) | Figure 3 | OK |
| 9 | "Everyone dragged" | A = 1.0 means one group; the gesture is not named (Figure 4 Move1: drag) | UNVERIFIABLE (plausible) |
| 9 | Duplicate, rotate, select group, insert, select single ≈ 0.35–0.55 | ≈ .52, .48, .48, .48, .37 | OK |
| 9 | After first seven, at or below about 0.3 | Shrink ≈ .31, then falling | OK |
| 9 | Tail task switch → cut at 0.10–0.15 | ≈ .22 → .11 | WRONG |
| 9 | 2H above 1H for select single, pan, enlarge, maximize | Also insert (≈ .58 vs .48), and slightly zoom in/out, previous, undo | WRONG (incomplete) |
| 9 | 2H dips near 0.1 for close and paste | Close ≈ .19, paste ≈ .23 (above 1H); task switch ≈ .14; menu access ≈ .20 | WRONG |
| 9 | A1H = 0.32, A2H = 0.28 | p. 1087 | OK |
| 9 | r = −.52 | p. 1087 (F1,25 = 9.51, p < .01) | OK |
| 9 | Widgets for menu, help, undo, accept, reject, zoom | "delete, zoom in, zoom out, accept, reject, menu access, and help" (p. 1091) | WRONG (undo in, delete out) |
| 9 | Discussion cites "Figure 2" for the cutoff after seven | p. 1091 | OK (the slip is real) |
| 10 | 31 one-hand, 17 two-hand (35%) | 31 (64.6%), 17 (35.4%), p. 1088 | OK |
| 10 | Ten commands have one gesture, rest have 2–5 | 10×1, 4×2, 3×3, 4×4, 1×5, p. 1088 | OK |
| 10 | Spread fingers: object = enlarge, background = zoom in, folder = open | p. 1088, 1091; Figure 4 | OK |
| 10 | Drag off-screen: object = delete, dialog = reject | Figure 4 Delete1, Reject2/3 | OK |
| 10 | Hold-and-tap reused for move, duplicate, delete, paste, minimize | Figure 4: "hold and tap" names Select Group1; Move2 "jump" reused by Delete2, Paste3, Minimize2; Duplicate "tap source and destination" | PARTLY |
| 10 | ✓ accept, ✗ reject, ? help, scribble undo | Figure 4 (Undo: scratch out) | OK |
| 10 | Menu: pull out like a drawer | Figure 4 "Menu: pull out" | OK |
| 10 | Insert and maximize reuse drag and enlarge; close and task switch → imaginary widgets | p. 1087 | OK |
| 10 | 1–3 fingers | Figure 4 caption | OK |
| 11, 12 | 43.5% each, 60.9% combined, 19.1% never tried | p. 1090 | OK |
| 11 | 13 of 20 varied fingers; only 2 said it mattered; 1–3 a point, palm more, four grey | p. 1090 (also "5-finger") | OK |
| 11 | One hand preferred for 25 of 27; two hands never preferred | p. 1088 | OK |
| 11 | 72% mouse-like; "I'm a child of the mouse" | p. 1090 | OK |
| 11 | Popular gestures rated better | p. 1088 (χ² N = 1,074) | OK |
| 11, 22 | Easy ≈ 3.4 s, hard ≈ 2.0 s | p. 1088 (articulation time) | OK |
| 11 | Reversible gestures although pairs never shown together; off-screen; above table | p. 1090 | OK |
| 12 | "thinking creatively" / "hopelessly lost" | p. 1090 | OK |
| 13 | Four contributions as listed | p. 1084 | OK |
| 13 | List of four does not include the method | True of the list; abstract leads with "We present an approach" (p. 1083) | MISLEADING |
| 13 | "an important next step" | p. 1091 | OK |
| 13 | Set never validated | Not in this paper (p. 1091); GI 2010 tested preference (slide 20) | PARTLY |
| 13 | 216 studies, Villarreal-Narvaez et al. DIS 2020 | Outside paper; matches the published title | OK (outside paper) |
| 13 | Good 1984 typed commands; Epps 2006 Windows images | p. 1084 | OK |
| 13 | "first to employ users, rather than principles" | p. 1083–1084 (sentence spans the page break) | OK |
| 14 | "useful not just for what it contains, but also for what it omits" | p. 1088 | OK |
| 14 | "Avoids the Windows bias that undermined Epps" | Epps "may have biased"; desktop thinking persisted (p. 1090) | WRONG (overclaim) |
| 15 | Authors' limitations: no revision, no context, literate American adults, not validated | p. 1091 | OK |
| 15 | No inter-rater reliability reported | None reported (p. 1086–1087) | OK |
| 15 | 27 commands "came from desktop systems" | "desktop and tabletop systems" (p. 1085) | WRONG (partly) |
| 15 | Claims "ease of recognition" | Conclusion, p. 1091 | OK |
| 15 | 43% not accommodated | 100 − 57.0 | OK |
| 15 | Figure 2 cited where Figure 3 is meant | p. 1091 | OK |
| 16 | Select single 11 of 20 → 0.37, 9 others | Eq. 3, p. 1087: 11, 3, 3, 3 → 0.37 (printed by the paper) | OK (but the notes wrongly call it "our estimate"; "tapped" not stated) |
| 16 | Vatavu and Wobbrock CHI 2015 corrected agreement rate | Outside paper | OK (outside paper) |
| 16 | Coverage = share of all proposals matched | p. 1087 | OK |
| 16 | Guessability is "what the paper is named after" | Title is "User-Defined Gestures for Surface Computing"; guessability is the method's name (p. 1083) | WRONG |
| 17 | Two-hand gestures forced for all commands | p. 1086 | OK |
| 17 | "Zoom in, that's the same as enlarge." | "Oh, that's the same as enlarge." upon being prompted to zoom in (p. 1090) | WRONG (misquote) |
| 17 | Maximize became enlarge | p. 1087 | OK |
| 17 | Minimize became shrink | Minimize = drag to bottom / Move2 (Figure 4) | WRONG |
| 17 | Figure 3 order ≠ Table 1 ranking | p. 1090 | OK ("this is why" overclaims) |
| 17 | Morris et al. 2014 formalized production | Outside paper; "production" is one of their techniques | OK (loose) |
| 18 | Section title "It's a Windows World" | p. 1090 | OK |
| 18 | Screening: no tabletop, no iPhone | Also "or similar" (p. 1085) and "without prior experience using touch screens" (p. 1084) | WRONG (understated) |
| 18 | Final set has pinch and reverse pinch | Figure 4 Enlarge/Shrink3 "pinch" (reversible); p. 1091 "reverse pinch" 1 of 5 | OK (inference weak) |
| 18 | Device dates (Palm 1996, iPod touch wheel 2002, DS 2004, Wii 2006, FingerWorks / Apple 2005, BMW 7 Series 2015) | Outside paper; consistent with general knowledge | OK (outside paper) |
| 18 | "Every mass-market item teaches one finger, one point" | Contradicted by the slide's own Wii row | WRONG (internal) |
| 19 | Hutchins, Hollan and Norman ref. 10; Morris cooperative gestures ref. 17; Wilson physics UIST 2008 ref. 32; 44% physical as support | p. 1084, 1085, 1091 (43.9%) | OK |
| 19 | Long et al.: designers misjudge similarity | "users are sometimes poor…" (p. 1084) | WRONG (wording) |
| 20 | Reverse study; vision-based recognizer | p. 1091 | OK |
| 20 | See all commands first; context, children, cultures; sense above and beyond the table | Limitations and Implications sections, not next steps (p. 1091) | PARTLY |
| 20 | GI 2010, CHI 2015, Interactions 2014, DIS 2020 | Outside paper (verified by presenters) | OK (outside paper) |
| 22 | Minimize "as defined by the user's seating position" | Figure 4 Minimize2 | OK |
| — | "only 2 of 14 resized along the diagonal" (notes only, not on slides) | p. 1091 "Only 2 of 14 participants for 2-hand enlarge…" | OK |

## Answers to the five judgment questions

**Q1. Is Part A too long or too summary-like? YES.** Slides 2–10 follow the paper's section order and take about a third of the talk. Slides 6 (method table), 8 (taxonomy shares) and 10 (the set) are recap for a room that has read the paper. Slides 3, 5 and 7 earn their time because later critique depends on them; the rest could shrink to about 6 minutes.

**Q2. Does slide 18 hold up as stated? NO.** It misreads "unrevised behavior" and understates the paper's own screening claim ("without prior experience using touch screens", p. 1084). Its "every item teaches one point" line is contradicted by its own Wii row, and it ignores both the desktop-specific think-aloud evidence and the plain pointing explanation the paper cites. There is a good argument underneath: elicitation measures the repertoire the participant brings in. The slide needs to make that argument against the right claim.

**Q3. Is the slide 21 thesis defensible from the deck's evidence? PARTLY.** "The lasting contribution is the method" is defensible only as "this paper made the 2005 method a genre". Slide 13 itself says the method was borrowed. Of the three supports, one is contradicted by slide 20 (GI 2010), one is overstated ("circular"), and one misreads the paper ("unrevised"). "Its reputation" is never described in the deck.

**Q4. Is 23 slides in 28 minutes realistic at this density? NO.** The slide count is fine (about 73 s each), but about 3,500 words of on-slide text is not. Nine slides carry 165–250 words. The script for slide 18 runs at 154 wpm, the TODO slides 19 and 23 have no spoken text yet, and two-speaker handovers are not costed. Expect a 3–5 minute overrun into the discussion half.

**Q5. If the deck had to lose 5 minutes, what goes?**
- Fold slide 6 into slide 5 as three lines: about −1:00.
- Cut slide 8 to a single "72% single-point, 44% physical" callout on slide 9: about −0:45.
- Merge slide 12 into slide 11, with finding 1 as the big-number visual: about −1:00.
- Merge slide 16's agreement card into slide 15 or 17, and move the guessability point to slide 21: about −0:45.
- Drop slide 22's "five smaller things", all of which repeat earlier slides: about −0:45.
- Trim slide 14 to three tiles: about −0:30.

Every cut removes recap or duplication, not critique.

## Strengths of the deck

- Numbers are overwhelmingly accurate and page-referenced. The Figure 2 shares, the 0.42 example, 43.5/60.9/19.1 and all quotes check out exactly.
- Slide 16 ("three numbers that are not the same") is the best slide. Agreement ≠ percentage, coverage rewards size, guessability unmeasured: exactly the kind of insight the "mastery" band asks for, and it lines up with the CHI 2015 correction.
- Slide 13's "novel vs borrowed" is honest and well sourced (EdgeWrite, Good 1984, Epps).
- Slide 17's framing, "design decisions that shaped the result", is the right analytical move. It needs factual tightening, not rethinking.
- The legacy-bias theme (slides 18, 20, 22) leads into a strong discussion and a live-elicitation warm-up that tests the paper in the room.
- The speaker notes' "Watch out" lines already catch several traps: the 0.42 hand count, the [1/|Pr|, 1] exponent, the Figure 2/3 slip, not claiming validation.
