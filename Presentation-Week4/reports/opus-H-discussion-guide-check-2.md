# Opus H-2: independent check of `slides/16-Discussion-Guide.md`

Reviewer 2 of 2, run 2026-10-01 from Brief 17. **Version note:** I read file 16 first at its 2026-10-01 morning state; it was edited at 11:56 while this check ran (eight lines changed, apparently from reviewer 1's corrections). Every finding below was re-checked against the 11:56 version, and every OLD block matches that version exactly. Items the 11:56 edit already fixed are listed under "Already fixed" at the end of section 10 and are not counted as corrections. Sources read: the paper in full (pdftotext plus Figures 2, 3 and 4 rendered at 150 to 170 dpi), file 16 in full, Part 1 of `_agent-briefs.md`, the verification logs of sections 03, 04, 06, 09, 10, 11, section 12 in full, the question headings of 13 to 15, and the "Verified facts" table in `progress.md`. I did not see reviewer 1's report.

## 1. Summary

The guide is accurate on almost every number and page, keeps to the pitfalls list on the points that went wrong before (A not a percentage, Eq. 3, widget lists, GI 2010 wording, "unrevised behavior", patent as an application), and works as a discussion plan: open stems, at least one position defending the paper in every anchor, probes at both altitudes.

**Citations (106 checked, 11:56 version):** 93 confirmed, 11 wrong or overstated, 2 unverifiable. The wrong ones that matter: two misquotes ("thinking creatively or hopelessly lost"; "allow us to unify"), "the three authors predicted under half" (combined they matched 60.9%, pitfall 14), and a page slip (p. 1090 for "rarely matters", which is p. 1088). The rest are hedges dropped or inferences pinned to a page. The 11:56 edit already fixed Epps "had shown" (pitfall 9) and the 3-inch generalization.

**Outside facts (22):** 18 confirmed as logged, 2 stated more strongly than the logs (pinch became universal "because Apple shipped it"; the Anchor 6 stem's "the field adopted its sample size as a norm" without "likely"), 1 not in any log (Weiser "disappear"), 1 with no source at all ("Redmond").

**Anchors:** 3 strong (1, 3, 6), 3 adequate (2, 4, 5), 0 weak.

**Cards:** 5 produce real disagreement, 1 adequate (F), 2 weak (E, G). None embarrassing.

**Timing:** the totals add up, but the everyone-speaks rounds in Anchors 5 and 6 do not fit a class of more than about ten, and Anchor 6's three-minute version can't be used as written.

**Corrections:** 3 blockers, 14 should-fix, 13 nice-to-have.

## 2. Citation table

Status: C = confirmed, W = wrong or overstated against the page, U = unverifiable from the paper. "L" is the line in file 16.

| # | L | Guide says | Paper text found | Status |
|---|---|---|---|---|
| 1 | 24 | Square each share, add (Eq. 1, p. 1087) | Eq. 1, p. 1087 | C |
| 2 | 24 | Floor 0.05, ceiling 1 | "The range for A is [\|Pr\|^-1, 1]" p. 1087 | C |
| 3 | 24 | Select single 11, 3, 3, 3 gives 0.37 (Eq. 3) | Eq. 3, p. 1087, "select single (1-hand)" | C |
| 4 | 24 | 0.32 one hand, 0.28 two | "A1H=0.32 and A2H=0.28" p. 1087 | C |
| 5 | 24 | About 72% mouse-like (p. 1090) | "about 72% of gestures were mouse-like one-point touches or paths" | C |
| 6 | 24 | 43.5%, 60.9%, 19.1% (p. 1090) | all three on p. 1090 | C |
| 7 | 24 | One hand preferred 25 of 27 (p. 1088) | p. 1088 | C |
| 8 | 24 | 48 gestures for 22 referents, 57.0% (pp. 1087, 1088) | 10+4+3+4+1 = 22 referents, 48 gestures p. 1088; "covers 57.0%" p. 1087 | C |
| 9 | 24 | 8 s and 15 s; 3.4 s and 2.0 s (p. 1088) | p. 1088 | C |
| 10 | 42 | Warm-up score by Eq. 1 | p. 1087 | C |
| 11 | 44 | "could be performed with 1-3 fingers" (p. 1089), framed as "The paper's own rule" for the grouping vote | Caption: "Gestures depicted as using one finger could be performed with 1-3 fingers." The caption describes the final set; the paper never says how grouping treated finger count (pitfall 11) | W (framing) |
| 12 | 44 | Several fingers at one point still one-point touch (p. 1086) | "still considered a one-point touch or path even if the user casually touches with more than one finger at the same point" | C |
| 13 | 44 | 0.65 and 0.41 illustration | 0.64 + 0.01; 0.36 + 0.04 + 0.01 | C (own arithmetic, labelled) |
| 14 | 45 | "I'm a child of the mouse" (p. 1090) | p. 1090 | C |
| 15 | 56 | "screened for never having used an iPhone or a tabletop" (no page) | "No participant had used an interactive tabletop, Apple iPhone, or similar" p. 1085 | C (page missing) |
| 16 | 57 | "informed by user behavior" (p. 1083) | "better gesture sets informed by user behavior" p. 1083 | C |
| 17 | 59 | "where the importance of immediate usability is high" (p. 1085) | p. 1085 | C |
| 18 | 59 | "desktop idioms strongly influence users' mental models" (p. 1083) | p. 1083 | C |
| 19 | 59 | "It's a Windows World" (p. 1090) | heading p. 1090 | C |
| 20 | 59 | "designs around it: one-point gestures, widgets where agreement is low (p. 1091)" | Widgets: p. 1091. The paper nowhere recommends one-point gestures as a response to the desktop pull | W (half unsupported) |
| 21 | 60 | Pinch in the set, Enlarge 3 (p. 1089) | "Enlarge (Shrink)3: pinch" | C |
| 22 | 61 | "obtained from desktop and tabletop systems" (p. 1085) | p. 1085 | C |
| 23 | 61 | Table 1 includes cut, paste, minimize, maximize, task switch (p. 1086) | Table 1 | C |
| 24 | 61 | Two-hand gesture for every referent (p. 1086) | "participants performed a 1-hand and a 2-hand gesture" | C |
| 25 | 61 | "an interactive tabletop, Apple iPhone, or similar" (p. 1085) | p. 1085 | C |
| 26 | 61 | "Unrevised behavior" means unrevised by feedback (p. 1085) | p. 1085 context | C |
| 27 | 62 | "children, Eastern, or uneducated participants would behave differently" (p. 1091) | p. 1091 | C |
| 28 | 64 | "to avoid bias" (p. 1085); Epps prompt cue (p. 1084) | "To avoid bias [5], no elements specific to Windows or the Macintosh were shown." | C |
| 29 | 67 | "a key benefit of gestures that couples selection and action" (p. 1090) | "negating a key benefit of gestures that couples selection and action [13]" | C |
| 30 | 79 | "held identical gestures" | p. 1087 | C |
| 31 | 79 | "The paper made that decision 54 times" (no label) | 27 referents x 2 hand conditions; the paper gives no count | C (arithmetic, unlabelled) |
| 32 | 80 | 11:56 version: the process was "adopted from prior work" (p. 1087) | "(This process was adopted from prior work [33].)" | C (fixed at 11:56; the earlier wording was garbled) |
| 33 | 82 | Formula, two worked examples, range (p. 1087) | p. 1087 | C |
| 34 | 82 | Move referents 1.0, cliff, tail (Figure 3, p. 1088) | Figure 3 rendered | C |
| 35 | 82 | "complexity predicts agreement, r = -0.52 (p. 1087)" | "correlated significantly and inversely ... (r=-.52 ...)" | C ("predicts" is loose) |
| 36 | 83 | "Identical" never defined (p. 1087) | correct | C |
| 37 | 83 | "grouped and coded alone, with no reliability check (pp. 1086 to 1087)" | "The authors manually classified each gesture" p. 1086; "we grouped" p. 1087. No reliability is *reported*; the paper does not say none was done | W (overstated; pitfall 18 says "reported") |
| 38 | 83 | "as our participants often did" (p. 1086) | p. 1086 | C |
| 39 | 83 | Figure 4, 1 to 3 fingers (p. 1089) | caption | C |
| 40 | 84 | "an important next step" (p. 1091); coverage 57.0% (p. 1087) | both | C |
| 41 | 85 | Widget recommendations (p. 1091) | "an on-screen widget as well as a gesture"; "delete, zoom in, zoom out, accept, reject, menu access, and help" | C |
| 42 | 85 | "Our user-defined set is useful, therefore, not just for what it contains, but also for what it omits" (p. 1088) | p. 1088, top of left column. Its "therefore" refers to the four referents without gestures (p. 1087), not to the low tail generally | C (context slightly widened) |
| 43 | 87 | "Select single: eleven of twenty tapped" | Paper gives group sizes 11, 3, 3, 3 only. That the 11 tapped is inferred from Figure 4 (section 04 says "a tap according to Figure 4") | U (inference unlabelled) |
| 44 | 87 | Figure 4 keeps tap and lasso (p. 1089) | Select Single1 tap, Select Single2 lasso | C |
| 45 | 88 | Tail about 0.21 task switch to about 0.11 cut | rendered: task switch about 0.21, cut about 0.11 | C |
| 46 | 89 | "the referent with the largest group won the gesture" (p. 1087) | p. 1087 | C |
| 47 | 90 | Chi-square on 1,074 (pp. 1087 to 1088) | χ2(3,N=1074) p. 1087; four tests p. 1088 | C |
| 48 | 92 | "learnability and memorability" (p. 1091, Wu et al.) | "Gesture reuse is important to increase learnability and memorability [35]"; [35] is Wu et al. | C |
| 49 | 92 | Long et al., p. 1084 | p. 1084 | C |
| 50 | 94 | "The paper says finger count rarely matters, p. 1090" | Wording is p. 1088 ("the number of fingers rarely matters"); p. 1090 says only 2 of 13 said it mattered | W (page) |
| 51 | 97 | Goodness ratings not circular, ratings given before group sizes known | "Immediately after performing each gesture" p. 1088 | C |
| 52 | 102 | "goes a long way to validating" | p. 1088 | C |
| 53 | 102 | "The three authors predicted under half of what users produced." | "each author individually came up with only 43.5% ... Even combined, the authors only covered 60.9%" p. 1090. Together they matched more than half | W (pitfall 14) |
| 54 | 102 | "The guessability test the authors proposed is not in the paper." | It is in the paper (p. 1091) as a proposal; it was not run | W (wording) |
| 55 | 105 | Four contributions (p. 1084) | p. 1084 | C |
| 56 | 105 | 43.5%, 60.9% (p. 1090) | p. 1090 | C |
| 57 | 106 | "a good candidate for deployment" with "ease of recognition" (p. 1091) | Conclusion p. 1091 | C |
| 58 | 106 | Inventor rated own gesture (p. 1088); N = 1,074 treated as independent (p. 1088) | p. 1088 | C |
| 59 | 106 | "three experts cannot generate the scope of gestures that 20 participants can" (p. 1090) | exact | C |
| 60 | 108 | Unlabeled clips, 20 new people, "from signs to referents" (p. 1091) | exact | C |
| 61 | 110 | "wisdom of crowds" would generate a better set than the authors (p. 1090) | "it was hypothesized that the "wisdom of crowds" would generate a better set than the authors" | C |
| 62 | 111 | 19.1%: "thinking creatively or hopelessly lost" (p. 1090) | "the authors are either thinking creatively or are hopelessly lost!" The quote drops "are" | W (misquote) |
| 63 | 112 | Ratings right after each gesture (p. 1088) | p. 1088 | C |
| 64 | 125 | "unrevised behavior" | p. 1085 | C |
| 65 | 128 | "the vicissitudes of gesture recognition"; "a monologue in which the user's behavior is always acceptable" (p. 1085) | both exact | C |
| 66 | 128 | 11:56 version: Epps et al. "conceded that their Windows prompts 'may have biased' people toward the mouse" (p. 1084) | "acknowledged that their Windows-based prompts may have biased participants to simply emulate the mouse" | C (fixed at 11:56; the earlier "had shown" broke pitfall 9) |
| 67 | 128 | "For walk-up public use, the first attempt is the one that counts (p. 1085)" | p. 1085 says only "the importance of immediate usability is high" | W (paraphrase stronger than page) |
| 68 | 129 | "could not change previous gestures"; all referents first; "interactive systems that provide feedback will implicitly remind users to remain in contact with the table" (p. 1091) | all exact | C |
| 69 | 129 | "to ensure that I was pressing"; "lack of feedback from the table" (p. 1090), "which the authors tie to" | "This may be, at least in part, due to the lack of feedback" | C (hedge dropped) |
| 70 | 130 | Two-hand for every referent; 25 of 27; 17 of 48; random order; "consistency and symmetry" | pp. 1086, 1088 | C |
| 71 | 130 | "allow us to unify the gesture set" (p. 1090) | "This allows us to unify the gesture set" | W (misquote) |
| 72 | 133 | "a common gesture was not feasible" (p. 1087) | exact | C |
| 73 | 134 | 11:56 version: "one still crossed her hands with the top hand about 3 inches above it (p. 1090)" | One participant made an X with her hands "such that the top hand was about 3" off the table's surface" | C (fixed at 11:56; the earlier "some still gestured" generalized one person) |
| 74 | 134 | Sense above the table (p. 1091) | "look down or sideways at users' hands"; "off-table sensing" | C |
| 75 | 135 | Reversible gestures though never shown as pairs (p. 1090) | "even though the study software did not present these referents together" | C |
| 76 | 136 | Insert and maximize two-hand about 0.58 (Figure 3); tied preference; reused gestures (pp. 1087 to 1088) | rendered: both about 0.58; p. 1088 text | C |
| 77 | 141 | 13 of 20 varied fingers (p. 1090) | exact | C |
| 78 | 143 | Touch instruction under Results (p. 1090) | "Acting above the Table," under Mental Model Observations | C |
| 79 | 148 | "would behave differently"; literate American adults; "tested one at a time" | p. 1091; one-at-a-time implied by "two authors observed each session" p. 1086 | C |
| 80 | 149 | "a good candidate for deployment" (p. 1091) | exact | C |
| 81 | 151 | Symbolic .100 (Figure 2, p. 1087) | printed label .100 | C |
| 82 | 151 | Check, X, ? marks (Figure 4, p. 1089); widgets for accept, reject, help (p. 1091) | both | C |
| 83 | 151 | "aliasing lets a locale add its own mark (p. 1088)" | p. 1088 says aliasing increases guessability [8,33]; nothing about locales | W (inference attributed to the page) |
| 84 | 151 | Physical 43.9% largest category (p. 1091) | "the large number of physical gestures (43.9%)" p. 1091; .439 is the largest Nature bar | C |
| 85 | 152 | Next drawn left to right (Figure 4); reading direction labelled inference | Figure 4 arrow lower left to upper right; caption: reversible gestures "depicted in only one direction" | C |
| 86 | 152 | Close top right, taskbar bottom (p. 1090) | p. 1090 (also p. 1087) | C |
| 87 | 152 | "Oh, that's the same as enlarge" (p. 1090) "is a desktop education speaking" | quote exact; the reading is the presenters' and is unlabelled | C (quote) |
| 88 | 153 | Morris et al. "address groups of users and omit issues relevant to single-user gestures, which we cover here" (p. 1084) | exact | C |
| 89 | 153 | Minimize 2 "to the bottom of the surface (as defined by user's seating position)" (p. 1089) | exact | C |
| 90 | 154 | All right-handed (p. 1085; Figure 2 caption p. 1087) | both | C |
| 91 | 154 | "in two-hand gestures" (added at 11:56) left .343 vs right .274; left .379 vs right .450, "about ... chart readings" | Printed labels in Figure 2, not readings. Left column covers only two-hand gestures (caption) | C (values); see S9 for the inference |
| 92 | 154 | Rotate corner drag (Figure 4); "demands for precise positioning should be avoided" (p. 1091) | both | C |
| 93 | 154 | "gestures rated easy were the longer, smoother ones (p. 1088)" | "gestures that took longer to perform were generally rated as easier, perhaps because they were smoother or less hasty" | W (hedge dropped) |
| 94 | 157 | Mean 43.2, sd 15.6 (p. 1085); someone under about 28 or over about 58 | p. 1085; 15.6 x sqrt(19/20) = 15.2 | C |
| 95 | 162 | Furnas behind aliasing (p. 1088) | [8] Furnas et al. cited for aliasing p. 1088 | C |
| 96 | 174 | "an approach to designing tabletop gestures" (p. 1083) | "We present an approach to designing tabletop gestures" | C |
| 97 | 175 | "Gestures should not be distinguished by number of fingers"; four-finger boundary; precise positioning; widgets (p. 1091) | all p. 1091 | C |
| 98 | 176 | Table 2 taxonomy (p. 1086) | p. 1086 | C |
| 99 | 178 | Two of three at Microsoft Research (p. 1083); Surface prototype (p. 1085) | both | C |
| 100 | 178 | "the first to employ users, rather than principles" (p. 1084) | "ours is the first to employ users, rather than principles, in the development of a gesture set" | C |
| 101 | 181 | "ease of recognition, consistency, reversibility, and versatility through aliasing" (p. 1091) | exact | C |
| 102 | 182 | 'not, in fact, "obvious"' because each author matched 43.5% (p. 1091) | "the gestures are not, in fact, "obvious", for example, ... each author proposed only 43.5%" | C |
| 103 | 200 to 207 | Card start-here pages: p. 1085 "immediate usability"; p. 1090 "child of the mouse"; Eq. 3; p. 1088 "goes a long way to validating"; p. 1085 "monologue"; p. 1090 "lack of feedback"; p. 1091 "could not change previous gestures"; p. 1091 "children, Eastern, or uneducated"; p. 1084 "omit issues..."; Figure 4 Minimize 2; p. 1091 "good candidate for deployment"; Enlarge 3 pinch; p. 1091 "1 of 5" | all found on the pages given | C |
| 104 | 261 | Complexity rated before the study (p. 1085); r = .71 (p. 1088) | "The three authors independently rated each referent's conceptual complexity before participants made gestures" p. 1085; r=.71 p. 1088 | C |
| 105 | 260, 262, 264 | Malik et al. 1, 2, 3, 5 fingers (p. 1085); "smoother or less hasty" (p. 1088); Schuler and Namioka [25] (p. 1084); "users are not designers" (p. 1085) | all found | C |
| 106 | 161 | "a convenience sample of the public in Redmond" | The paper gives no study location; participants "recruited from the general public" (p. 1085) | U |

Totals for the 11:56 version: C 93, W 11 (rows 11, 20, 37, 50, 53, 54, 62, 67, 71, 83, 93; the dropped hedge in row 69 is counted as C), U 2 (rows 43, 106). Against the morning version the count was C 90, W 14 (rows 32, 66, 73 also wrong). Rows 103 to 105 bundle 19 card and thread-map cites, all confirmed; counted as three rows.

## 3. Outside-facts table

| Fact as stated in 16 | Log or source | Verdict |
|---|---|---|
| Patent application US 2010/0031202 filed 2008-08-04, published 2010, abandoned 2014-12-09, three authors as inventors, "on the procedure," not in the paper (L24, L178) | progress.md verified facts; 03 log ("claim 1 on eliciting surface input from several users shown an effect") | Confirmed, not strengthened. "Application" and "attribute no motive" kept |
| GI 2010: 22 blind judges, 81 gestures, preferred larger-group gestures (L105) | progress.md; 04 and 11 logs | Confirmed. Position 1 omits "partly because the researchers' gestures were more complex"; position 2 has it. Acceptable (pitfall 6 not violated), nice-to-have N8 |
| Morris et al. 2014 "named the problem 'legacy bias'"; production, priming, partners (L60, L131) | 11 log; progress.md | Confirmed |
| Morris et al. 2014 as the authors' own work: L60 now "two of the three authors"; L177 now "Two of the authors' 2014 article" | 11 log says "Morris, Wobbrock and colleagues" | Confirmed at 11:56 (the morning version said "the authors themselves" and "the authors' own"). L177 is ungrammatical, nice-to-have N13 |
| Vatavu and Wobbrock 2015: 0.32 becomes 0.284 (L93) | 04 log (CHI 2015 PDF p. 7) | Confirmed |
| DIS 2020: 216 studies; first application to hand gestures; most cited (1,070 citations as of 2019) (L107, L171, L174) | 04 log (pp. 855, 859); progress.md | Confirmed |
| DIS 2020: 20 the most common sample size "likely" because of this paper | 09 log: "likely explained by the fact that Wobbrock et al. [123] employed 20 participants" (p. 859) | Kept at L24, L107 (both restored at 11:56), L159 and L177. Still dropped in the Anchor 6 stem, L171 ("the field adopted its sample size as a norm"). Should-fix S4. L107's 'the review "likely" traces' puts the quote marks on a word the review uses differently, nice-to-have N14 |
| Wobbrock and Kientz: seven types; artifacts need not come with studies (p. 40); "repeated validation" (p. 41); 2016 article does not name this paper | 10 log | Confirmed |
| Wobbrock 2012 whitepaper lists this paper as a methodological example | 10 log (p. 3) | Confirmed |
| Greenberg and Buxton p. 111 "can mute creative ideas that do not conform to current interface norms"; p. 115 "converge on that which we can measure, which is almost always that which we are already familiar with" | 10 log | Confirmed, both exact |
| Olsen: "empowering new design participants," three steps | 10 log (pre-print, section heading) | Confirmed. The guide's step 3 "show the new tool is more effective" compresses "accessible, easier or more effective"; fine |
| Weiser: menu bar doubt on a board (L141) | 10 log | Confirmed |
| "Weiser wanted computers that disappear (section 10)" (L138) | Not in section 10 or its log | Not in the logs. Widely known (the article's opening sentence), but the guide cites section 10 for it. Nice-to-have N10 |
| Japan circle and check; Bulgaria nod means no (L152) | 06 log | Confirmed, not strengthened ("a check can mark wrong") |
| Trackpads, Palm Graffiti, ticket machines, Nintendo DS (L61) | 06 log (1994, 1996, 1999, 2004) | Confirmed |
| Mauney et al. 2010 [UNVERIFIED] (L166) | 06 log | Correctly kept as unverified |
| iPhone shipped June 2007 with pinch (L60) | 11 log | Confirmed |
| "Pinch ... became universal because Apple shipped it ..., not because anyone elicited it (section 11)" (L60); card H "Pinch became universal because Apple shipped it" (L207) | 11 says "pinch became universal anyway" and, under Watch out, that "the study cannot tell the two apart." No source for the causal claim | Strengthened. A presenters' argument stated as fact and credited to section 11. Should-fix S3 |
| "No product we could find shipped its gesture set" (L171, stem, 11:56 version) | 11 log: "Judgment, labelled as such" | Confirmed; hedged at 11:56 (the morning version stated it as fact) |
| BMW 7 Series 2015 (L70) | progress.md | Confirmed |
| EdgeWrite 2005 measured guessability (L84) | 04 log (51.0% to 80.1%, abstract only) | Confirmed |
| "a convenience sample of the public in Redmond" (L161) | No log, no paper text | Unsupported. Should-fix S6 |

## 4. Pitfalls conflicts

| Pitfall | Where in 16 | Conflict |
|---|---|---|
| 9 (Epps "may have biased", not "had pushed") | Anchor 4, position 1 | The morning version had "Epps et al. had shown that prompts bias people," the same error in a new form. Fixed in the 11:56 version; no conflict now |
| 14 (each 43.5%, combined 60.9%) | Anchor 3 stem: "The three authors predicted under half of what users produced." | Yes. Combined they matched 60.9%. Blocker B2 |
| 11 (the paper does not say how finger count was treated in grouping) | Section 2, first bullet: "The paper's own rule, in the Figure 4 caption ..." used to settle the live grouping vote | Mild. The caption is about the final set, not the grouping. Anchor 2 position 2 handles it correctly ("we argue ... the paper does not say"). Should-fix S8 |
| 18 (no inter-rater reliability *reported*) | Anchor 2, position 2: "with no reliability check" | Mild overstatement. Should-fix S1 |
| 17 (articulation time and ease) | Anchor 5, position 4: "the longer, smoother ones" | Mild. The paper only says "perhaps" smoother. Should-fix S10 |
| 13 (participants; no ethnicity) | Anchor 5 traps handle "American" correctly. "Redmond" (L161) adds a location the paper does not give | Not a pitfall item, but the same family. Should-fix S6 |
| 1, 2, 3, 4, 5, 6, 7, 8, 10, 12, 15, 16, 19, 20 | Throughout | No conflict. Traps in each anchor cite the right pitfall numbers |

## 5. Per-anchor verdicts

The five principles: (P1) move between significance and specific evidence; (P2) more than one defensible answer; (P3) pair or small-group work; (P4) follow ideas, not a list; (P5) spread participation. Plus Oliver's requirement: room to elaborate on details and express views.

### Anchor 1. Whose gestures are these? **Strong**

- **Stem:** open ("whose habits," "is that what a designer should build on"). It grounds itself in the room's own warm-up data, which is the best opener in the guide. Nobody can answer it yes or no.
- **Positions:** four that differ in kind: build on it (defends the paper with its own words, pp. 1083, 1085, 1090, 1091); it locks in the past; it measures the method; a snapshot. Each can be argued at length.
- **Probes:** zoom-in uses specific evidence (the "to avoid bias" sentence, the desktop referents, select-then-gesture on p. 1090). Zoom-out reaches design philosophy and today's devices.
- **Weakness:** the three-way hands options ("mostly about the people," "mostly about the procedure," "badly posed") cover only the first half of the stem. The second half, which holds the paper's defence, gets no vote.
- **Most valuable improvement:** make the hands vote match the three positions people can take on the second half: "build on it," "design against it," "rerun it per population and say so." Then the paper's defence has a corner of the room from the first minute.

### Anchor 2. What does an agreement score tell a designer? **Adequate**

- **Stem:** good. It opens on the room's two numbers, but it asks two things at once ("what does it tell," "where would you stop trusting it").
- **Positions:** four that differ in kind, two of which defend the paper (1 and 4). Position 4, "the low end is the most useful output," is the most original reading in the guide.
- **Risk:** this is the anchor most likely to turn into a quiz on the formula. Both zoom-out probes stay inside the metric (alternative criteria; the 2015 correction). Neither asks what the score means for research practice beyond this paper. Students without a statistics background may say nothing.
- **Most valuable improvement:** add one zoom-out probe on significance that anyone can answer. For example: "Think of a number in your own research that hides a judgment call like this grouping. What would you have to report next to it for a reader to trust it?" This ties the anchor to the course goal of broader implications and gives non-quantitative students a way in.

### Anchor 3. What would convince you the method works? **Strong**

- **Stem:** the best-shaped question in the guide. "What would it take" asks people to build a standard before they judge the paper against it.
- **Positions:** four, clearly different. Position 1 is a real, well-sourced defence (Greenberg and Buxton, Wobbrock and Kientz, the designer comparison, GI 2010).
- **Probes:** Olsen's three steps and the "which column did the paper fill" board sort tie the specific evidence to the course readings.
- **Weakness:** the stem has a factual error ("the three authors predicted under half"; still present at 11:56). It also stacks three premises against the paper before asking, which tilts the room before position 1 is heard.
- **Most valuable improvement:** fix the stem (blocker B1), and add one premise in the paper's favour. For example: "Each author alone matched under half of what users produced." That fact cuts both ways and leaves the question open.

### Anchor 4. Was the monologue the right experiment? **Adequate**

- **Stem:** excellent ("what did that choice buy, what did it cost, how would you have run it"). It invites design reasoning, not a verdict.
- **Positions:** four that differ in kind. Position 1 defends the paper in its own words.
- **Probes:** concrete (imaginary widgets, reversible gestures, insert and maximize at about 0.58).
- **Weaknesses:** the move is a two-way seat split, which breaks the guide's own rule "Hands in three, never two" (section 6). Position 3 misquotes p. 1090 (B3). Position 1 says "the first attempt is the one that counts (p. 1085)," which is stronger than the page (N4). The pitfall 9 slip and the 3-inch overstatement were fixed at 11:56. The insert and maximize probe is the strongest evidence probe in the guide and deserves a better place than fourth.
- **Most valuable improvement:** split the room in three: first instinct, learned behaviour, and "design the two-phase study." Then the hybrid comes from a group, not from a single late voice.

### Anchor 5. Who was not at the table, and how far does that reach? **Adequate**

- **Stem:** three questions in one breath ("Where does that limitation stop? Which parts of Figure 4 would survive ..., and which would not?"), plus a list of premises. Too much to hold before speaking.
- **Positions:** strong and varied, with a fair defence in position 1 (symbolic gestures are .100 of proposals and already hedged by widgets). Position 3 (groups) is the most original idea in the guide.
- **Probes:** concrete.
- **Weaknesses:** the move, "Everyone speaks once in under two minutes," only works for about eight people. Position 4's left-hand argument rests on a Figure 2 comparison that mixes up hand and role (S9).
- **Most valuable improvement:** cut the stem to one question with Figure 4 on screen ("Which of these would you not ship in a different room, and why?"). Cap the round at eight to ten voices and take the panels nobody has named yet.

### Anchor 6. Seventeen years on, what did this paper give us? **Strong**

- **Stem:** open. Five positions: method, findings, shared language, influence ahead of evidence, and the industry context with a careful facts-only rule. At least two defend the paper.
- **Probes:** the "four contributions against seven types" zoom-in and the "reviewer in 2009 versus today" zoom-out are both excellent.
- **Weakness:** the stem sets out four premises that all lead to position 4, and one drops the review's "likely" ("the field adopted its sample size as a norm"). "No product we could find" is now properly hedged. The closer asks what to change, which does not answer the stem's own question.
- **Most valuable improvement:** make the stem neutral: restore "likely" (S4), and add one fact on the paper's side, such as the four-finger rule or the widget advice, which practitioners can use. The closing sentence then answers a fair question.

## 6. Card review

| Card | Real disagreement? | Fair and workable? | Note |
|---|---|---|---|
| A | Yes. Familiar for the walk-up user against new for the designer | Yes. "70, no smartphone" is fine; consider "has never used a smartphone" without the age, so nobody plays a caricature of an older person | Good opener for Anchor 1 |
| B | Yes. Rigour against a deadline | Workable. The statistician role may intimidate; the card gives Eq. 3 and the floor, which is enough | Good |
| C | Yes | Workable. Add "argue from the paper's text, not from motives" so the author role does not drift into speculation about Microsoft | Good |
| D | Yes, and very concrete | Yes. Both roles come straight from the paper's text | Best card |
| E | **Weak.** The question asks "which gestures travel," but the recognizer engineer's stake is aliases, not travel. The two roles talk past each other | Workable only if reframed | S15: ask whether a Tokyo edition should swap the marks or add local aliases. Then the engineer has something to resist |
| F | Adequate. Needs some knowledge of Morris et al. 2006, which the class has not read; p. 1084 and Minimize 2 are enough to start | Fair | Point out to the pair that Morris is an author of both papers; it makes the role fun |
| G | **Weak.** A researcher and a practitioner can easily agree the method mattered. Neither role is given a stance | Fair but flat | S16: give stances (the researcher credits the method; the practitioner needs findings she can ship) |
| H | Yes | Workable. The premise "because Apple shipped it" is the presenters' argument presented as fact | S3: make the premise factual |

None of the roles is likely to embarrass a student. The note that "a student may abandon the role once they have a view" is good and should be said aloud when the cards go out.

## 7. Timing review

- **Totals:** 5 + 5 + 2 + 6 + 6 + 5 + 1 = 30 minutes, with about 4 minutes of buffer from the talk (25:45 planned for a 30-minute slot). Consistent with 12.3 and the brief.
- **Warm-up, 5 min:** tight. Two referents, a show of gestures, tallying 20-odd undo gestures by eye, a grouping vote and two computations of A is closer to 6 or 7 minutes. Suggestions: hand out the pair cards *during* the warm-up, and compute A both ways only for move a little, where the two-finger variant actually appears. For undo, just count groups.
- **Pairs, 5 min:** feasible if the cards are already on the desks. Asking for "a page from the paper or a section from the talk" for each line will take most pairs the whole five minutes; accept "a figure or a section" as enough.
- **Harvest, 2 min:** feasible for three pairs, not four.
- **Anchor 1, 6 min, and Anchor 3, 6 min:** feasible. Anchor 3's 90-second sentence completion plus five answers plus the board sort uses about 4 minutes, which leaves 2 for argument. Tight but workable.
- **Anchor 6, 5 min:** not feasible as written for more than about 12 students ("each person, one sentence ... no repeats," plus the stem, the tally and rereading 12.1). At 15 seconds a voice, 20 students take 5 minutes for the round alone.
- **Three-minute versions:**
  - Anchors 1, 2, 4: usable as written.
  - Anchor 3: usable only if the sentence completion is cut to 60 seconds.
  - Anchor 5: usable with the four-panel round.
  - Anchor 6: "the stem and the closing round only" is **not usable** with a full class (S14).
- **Missing from the clock:** the [classmate] hooks. Reading a commentary sentence and letting its author speak first costs about 45 seconds an anchor. That fits inside the 6-minute slots but not inside a three-minute version. Say which hook to drop when an anchor is shortened.

## 8. Missing questions

1. **Is the set "conflict-free" because users agreed, or because the authors overloaded gestures by target?** The same splay means enlarge on an object, zoom in on the background and open on a container (p. 1088). The paper says simplified mental models "allow[s] us to ... disambiguate the effects of gestures based on where they occur" (p. 1090) and that "hit-testing within objects will be necessary" (p. 1091). Conflicts were settled by "the referent with the largest group won the gesture" (p. 1087). Zoom in: which conflicts were resolved by target and which by group size? Zoom out: is overloading by context good for learnability, as the paper says with "learnability and memorability [35]" (p. 1091), or a source of errors on a shared table? Several positions are defensible. It also draws on progress.md "Ten further points," item 2, which the guide does not use.
2. **Can you code a gesture's meaning without asking the person who made it?** The paper says of metaphorical gestures that "the gesture itself usually is not enough to reveal its metaphorical nature; the answer lies in the user's mental model" (p. 1086). Yet the authors "manually classified each gesture" (p. 1086), with no reliability reported, and Figure 2's Nature split (physical .439, abstract .272, metaphorical .189, symbolic .100) drives the physics-engine implication (p. 1091). This opens the taxonomy contribution, which only Anchor 6 position 3 touches, and connects to qualitative coding in the students' own research.
3. **Which of the technology implications follow from the data?** "Seven participants ... expected intervening objects to move out of the way," and "four participants 'threw' an object off-screen" (p. 1091). From this and 43.9% physical gestures the paper argues for a physics engine [32] and for off-table and beyond-edge sensing, while conceding "a physics engine alone will probably not suffice" (p. 1091). Zoom in: are 7 and 4 of 20 enough to steer an architecture? Zoom out: what should a designer take from elicitation data to an engineer? Ref. 32 is Wilson's own UIST 2008 work (p. 1092). That fact may be stated; a motive may not, so the facts-only rule from Anchor 6 position 5 applies. This also gives section 08 (technical obstacles), which no anchor currently uses, a place in the discussion.

## 9. Style findings

- **Dashes:** 0 em-dashes (U+2014) and 0 en-dashes (U+2013) in file 16.
- **Labels:** generally careful. "We argue," "our inference," "our judgment" and "ours" appear about 20 times, and the traps sections repeat the labels. These claims are unlabelled or attributed to the paper or a section without support:
  - "eleven of twenty tapped" (S2)
  - "aliasing lets a locale add its own mark (p. 1088)" (S11)
  - "Oh, that's the same as enlarge ... is a desktop education speaking" (N7)
  - "designs around it: one-point gestures" (N6)
  - "the first attempt is the one that counts (p. 1085)" (N4)
  - pinch "because Apple shipped it" (S3)
  - "the field adopted its sample size as a norm" in the Anchor 6 stem, without "likely" (S4)
- **Motive:** none stated as fact. Anchor 6 position 5 and its trap are well judged.
- **Grammar:** "Two of the authors' 2014 article named legacy bias" (L177, from the 11:56 edit) does not parse (N13).
- **Internal consistency:**
  - "Hands in three, never two" (section 6) is broken by Anchor 4's two-way seat split and by Anchor 6's "Accept? ... Accept?" probe. The second is a thought experiment, so it is fine.
  - Section 5 says "Q14.6 moves to design," and section 8 labels Q14.6 "timing." Same question, different labels (N9).
  - Section 8 lists "part of Q15.5" under Anchor 2, but Q15.5 is the sample-size question, which section 8 also maps to Anchor 5 (N11).
- **Tone:** "the joke lands" (L94) predicts the room. "Expect a laugh," or nothing, reads better (N1).

## 10. Corrections

Each OLD block is copied exactly from `slides/16-Discussion-Guide.md`.

### Blockers

**B1. Pitfall 14 (Anchor 3, Ask)**

OLD:
```
The three authors predicted under half of what users produced.
```
NEW:
```
Each author, working alone, came up with under half of the users' set; all three together, 60.9 percent.
```

**B2. Misquote (Anchor 3, zoom-in)**

OLD:
```
'thinking creatively or hopelessly lost' (p. 1090)
```
NEW:
```
'either thinking creatively or are hopelessly lost!' (p. 1090)
```

**B3. Misquote (Anchor 4, position 3)**

OLD:
```
participants' simplifications "allow us to unify the gesture set" (p. 1090)
```
NEW:
```
of participants' simplified mental models the paper says "This allows us to unify the gesture set" (p. 1090)
```

### Should-fix

**S1. Pitfall 18 (Anchor 2, position 2)**

OLD:
```
The authors grouped and coded alone, with no reliability check (pp. 1086 to 1087; section 09).
```
NEW:
```
The authors grouped and coded the gestures themselves, and no reliability check is reported (pp. 1086 to 1087; section 09).
```

**S2. Unlabelled inference (Anchor 2, zoom-in)**

OLD:
```
"Select single: eleven of twenty tapped; nine did something else in three groups of three; A = 0.37 (Eq. 3, p. 1087).
```
NEW:
```
"Select single: the largest group was eleven of twenty, by Figure 4 the tap (our reading); nine did something else in three groups of three; A = 0.37 (Eq. 3, p. 1087).
```

**S3a. Causal claim stated as fact (Anchor 1, position 2)**

OLD:
```
but it became universal because Apple shipped it in June 2007, not because anyone elicited it (section 11).
```
NEW:
```
but it was already shipping on the iPhone from June 2007, and we argue the market, not elicitation, made it universal; the study cannot tell the two apart (section 11).
```

**S3b. Same claim (card H)**

OLD:
```
| H | Pinch became universal because Apple shipped it. Does that make elicitation pointless, or prove it right? |
```
NEW:
```
| H | Pinch shipped on the iPhone in 2007, and it is also in this study's set. Does that make elicitation pointless, or prove it right? |
```

**S4. Dropped hedge in the Anchor 6 stem**

OLD:
```
And the field adopted its sample size as a norm.
```
NEW:
```
And twenty participants became the field's most common sample, likely because this paper used twenty.
```

**S6. Unsupported location (Anchor 5, zoom-out)**

OLD:
```
Is a convenience sample of the public in Redmond a stand-in for 'the public'?"
```
NEW:
```
Is a paid sample 'recruited from the general public' (p. 1085) a stand-in for 'the public'?"
```

**S7. Wording (Anchor 3, Ask)**

OLD:
```
The guessability test the authors proposed is not in the paper.
```
NEW:
```
The guessability test the authors proposed was not run in this paper.
```

**S8. Pitfall 11 (section 2, first bullet)**

OLD:
```
The paper's own rule, in the Figure 4 caption, is that a gesture drawn with one finger "could be performed with 1-3 fingers" (p. 1089);
```
NEW:
```
The paper never says how grouping treated finger count (pitfall 11). Two hints: the Figure 4 caption says a gesture drawn with one finger "could be performed with 1-3 fingers" (p. 1089);
```

**S9. Figure 2 inference and "chart readings" (Anchor 5, position 4)**

OLD:
```
in two-hand gestures the two hands differ in Figure 2: the left hand made more one-point touches and fewer one-point paths than the right (about .343 against .274 and .379 against .450, chart readings).
```
NEW:
```
in two-hand gestures the two hands differ in Figure 2: the left hand made more one-point touches and fewer one-point paths than the right (.343 against .274 and .379 against .450, printed in the figure). The non-dominant hand often holds while the other acts, so we read this as a difference of role as much as of hand.
```

**S10. Hedge dropped (Anchor 5, position 4)**

OLD:
```
found that gestures rated easy were the longer, smoother ones (p. 1088; the precision reading is ours, 07.2)
```
NEW:
```
found that gestures rated easy took longer, "perhaps because they were smoother or less hasty" (p. 1088; the precision reading is ours, 07.2)
```

**S11. Inference attributed to a page (Anchor 5, position 1)**

OLD:
```
and aliasing lets a locale add its own mark (p. 1088).
```
NEW:
```
and the set already uses aliasing (p. 1088), so we argue a locale could add its own mark.
```

**S13. Two-way split breaks the guide's own rule (Anchor 4, Move)**

OLD:
```
split the room by seat. Left half argues "first instinct is the product," right half "learned behaviour is the product," 60 seconds with a neighbour, two voices each. Then ask someone who has not spoken to propose the hybrid.
```
NEW:
```
split the room by seat into thirds. Left argues "first instinct is the product," right "learned behaviour is the product," middle "design the two-phase study." 60 seconds with a neighbour, two voices from each third, the middle last.
```

**S14a. Timing (Anchor 6, three-minute version)**

OLD:
```
- **Three-minute version:** the stem and the closing round only.
```
NEW:
```
- **Three-minute version:** the stem, then six to eight volunteers give their one sentence; everyone else writes theirs on the back of the pair card, collected for grading.
```

**S14b. Timing (section 5, Anchor 6 row)**

OLD:
```
| 24:00 to 29:00 | Anchor 6 | The closing round, one sentence each, no repeats. Reread 12.1's sentence as a question |
```
NEW:
```
| 24:00 to 29:00 | Anchor 6 | The closing round: one sentence each from as many as fit (about twelve), no repeats; the rest write theirs on the pair card. Reread 12.1's sentence as a question |
```

**S14c. Timing (Anchor 5, Move)**

OLD:
```
Everyone speaks once in under two minutes.
```
NEW:
```
Eight to ten voices in about two minutes, each taking a panel nobody has named yet.
```

**S15. Card E roles do not meet the question**

OLD:
```
| E | Which gestures in Figure 4 travel? |
```
NEW:
```
| E | Should a Tokyo edition of Figure 4 swap the marks, or keep them and add local marks as aliases? |
```

**S16. Card G has no built-in disagreement**

OLD:
```
| A researcher who ran one of the 216 later studies | A practitioner designing an AR gesture set this year |
```
NEW:
```
| A researcher who ran one of the 216 later studies and credits the method | A practitioner designing an AR gesture set this year who wants findings she can ship |
```

### Nice-to-have

**N1. Page and tone (Anchor 2, Move)**

OLD:
```
(The paper says finger count rarely matters, p. 1090; the joke lands.)
```
NEW:
```
(The paper says "the number of fingers rarely matters," p. 1088; expect a laugh.)
```

**N3. Label the count (Anchor 2, Ask)**

OLD:
```
The paper made that decision 54 times for 1,074 gestures
```
NEW:
```
By our count the paper made that decision 54 times, once per referent and hand, for 1,074 gestures,
```

**N4. Paraphrase stronger than the page (Anchor 4, position 1)**

OLD:
```
For walk-up public use, the first attempt is the one that counts (p. 1085).
```
NEW:
```
Public tables need "immediate usability" (p. 1085); we argue the first attempt is the one that counts there.
```

**N5. Hedge dropped (Anchor 4, position 2)**

OLD:
```
which the authors tie to "the lack of feedback from the table" (p. 1090)
```
NEW:
```
which the authors say "may be, at least in part, due to the lack of feedback from the table" (p. 1090)
```

**N6. Unsupported half-claim (Anchor 1, position 1)**

OLD:
```
and designs around it: one-point gestures, widgets where agreement is low (p. 1091).
```
NEW:
```
and designs around it: widgets where agreement is low (p. 1091).
```

**N7. Label the reading (Anchor 5, position 2)**

OLD:
```
"Oh, that's the same as enlarge" (p. 1090) is a desktop education speaking.
```
NEW:
```
We read "Oh, that's the same as enlarge" (p. 1090) as a desktop education speaking.
```

**N8. GI 2010 caveat (Anchor 3, position 1)**

OLD:
```
who preferred gestures authored by larger groups (sections 04, 11).
```
NEW:
```
who preferred gestures authored by larger groups, partly because the researchers' gestures were more complex (sections 04, 11).
```

**N9. Label consistency (section 5)**

OLD:
```
Q14.6 moves to design;
```
NEW:
```
Q14.6 (timing) moves to design;
```

**N10. Source for Weiser (Anchor 4, zoom-out)**

OLD:
```
Weiser wanted computers that disappear (section 10).
```
NEW:
```
Weiser wanted computers that disappear (Weiser 1991, opening sentence; not in section 10's log, so check it before quoting).
```

**N11. Mapping (section 8, Anchor 2 row)**

OLD:
```
| 2 Agreement score | part of Q15.5 |
```
NEW:
```
| 2 Agreement score | none (Q15.5 sits with Anchor 5) |
```

**N12. Correlation, not prediction (Anchor 2, position 1)**

OLD:
```
complexity predicts agreement, r = -0.52 (p. 1087)
```
NEW:
```
complexity correlates inversely with agreement, r = -0.52 (p. 1087)
```

**N13. Grammar (Anchor 6, position 4; introduced at 11:56)**

OLD:
```
Two of the authors' 2014 article named legacy bias as a limit of the method (section 11).
```
NEW:
```
A 2014 article by Morris, Wobbrock and colleagues, two of this paper's authors, named legacy bias as a limit of the method (section 11).
```

**N14. Quote marks on the wrong word (Anchor 3, position 3)**

OLD:
```
which the review "likely" traces to this paper (p. 859)
```
NEW:
```
which the review says is "likely explained" by this paper's 20 (p. 859)
```

### Already fixed in the 11:56 version (not counted)

- Epps et al. "had shown that prompts bias people" (pitfall 9), Anchor 4 position 1: now "conceded ... 'may have biased'". Confirmed against p. 1084.
- "some still gestured with a hand 3 inches above it," Anchor 4 zoom-in: now "one still crossed her hands." Confirmed against p. 1090.
- DIS 2020 "likely" restored in Numbers to have ready (L24) and Anchor 3 position 3 (L107).
- "No product shipped its gesture set" in the Anchor 6 stem: now "No product we could find."
- "the authors themselves" (L60) and "The authors' own 2014 article" (L177): now "two of the three authors" and "Two of the authors'" (see N13).
- "The paper calls the method its process" (L80): now 'the process was "adopted from prior work"'. Confirmed against p. 1087.
- Figure 2 left-hand comparison: "in two-hand gestures" added (S9 still applies for the role caveat and "chart readings").

**Counts (11:56 version):** 3 blockers (B1 to B3), 14 should-fix (S1 to S4, S6 to S11, S13 to S16; S5 and S12 were fixed at 11:56; S3 and S14 have several replacement blocks), 13 nice-to-have (N1, N3 to N14; N2 was fixed at 11:56).
