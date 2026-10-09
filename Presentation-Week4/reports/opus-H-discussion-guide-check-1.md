# Opus H, reviewer 1: check of `slides/16-Discussion-Guide.md`

Run 2026-10-01 under Brief 17. Ground truth: the paper PDF (all ten pages read; Figures 2, 3 and 4 rendered at 170 dpi and read), section files 03, 04, 05, 06, 07, 09, 10, 11, 12, the headings of 13 to 15, and `progress.md` "Verified facts".

**Version note.** The guide changed on disk at 11:56 while this check was running. This report checks the 11:56 version. That edit had already fixed seven items I had found in the earlier text: Epps et al. "had shown" (pitfall 9), "some ... 3 inches above," "because of this paper" in section 0, "traces" in Anchor 3, "the authors themselves" in Anchor 1, "the method its process" in Anchor 2, and "No product shipped" in Anchor 6. Those now count as confirmed. Every OLD block in section 10 matches the 11:56 text exactly, once each.

## 1. Summary

**Citations and quotations against the PDF: 129 checked. 119 confirmed, 5 wrong, 5 unverifiable.**

Wrong:
- the Anchor 3 stem says "the three authors predicted under half." In fact each author matched 43.5% and all three together 60.9%, of the users' set (p. 1090);
- "One participant tapped": the paper says "some participants" (p. 1090);
- two misquotes: "thinking creatively or hopelessly lost" and "allow us to unify";
- "designs around it: one-point gestures," which p. 1091 does not say.

Unverifiable:
- "54 times";
- "coded alone, with no reliability check" (the paper only reports none);
- "Redmond";
- "a device built for groups";
- "never tested" (twice).

**Outside facts: 19 checked.** 15 are faithful to the logs. Four are strengthened or not logged:
- the Anchor 6 stem drops the review's "likely";
- Weiser "disappear" is in no log;
- card H states the pinch causation as a premise;
- "first of its kind."

Two pitfall conflicts remain: 11 (section 2 framing) and 14 (Anchor 3 stem).

**Anchors: 3 strong (1, 2, 4), 3 adequate (3, 5, 6), 0 weak.** Every stem opens on a concrete detail, every anchor has four or five distinct positions, and position 1 always defends the paper.

The plan's main weakness is time. The warm-up, Anchor 6's closing round and its three-minute version do not fit as written.

**Corrections: 3 blockers, 15 should-fix, 16 nice-to-have.**

## 2. Citation table

C = CONFIRMED, W = WRONG, U = UNVERIFIABLE. Confirmed items are grouped by location; each group lists every citation it covers. The exact text found is given for every W and U.

| # | Location | Claim in the guide | Status | Text found / note |
|---|---|---|---|---|
| 1 to 9 | Section 0, numbers | Eq. 1 p. 1087; floor 0.05, ceiling 1; 11, 3, 3, 3 gives 0.37 (Eq. 3); 0.32 and 0.28; about 72% (p. 1090); 43.5, 60.9, 19.1 (p. 1090); 25 of 27 (p. 1088); 48 for 22 referents, 57.0% (pp. 1087, 1088); 8 s, 15 s, 3.4 s, 2.0 s (p. 1088) | C (9) | All as printed. "The range for A is [\|Pr\|^-1, 1]" (p. 1087). 10+8+9+16+5 = 48 over 22 referents (p. 1088). |
| 10 | Section 2 | Figure 4 caption, one-finger gestures "could be performed with 1-3 fingers" (p. 1089) | C (quote) | Quote exact. But it is called "the paper's own rule" in the middle of a *grouping* vote; see pitfall 11, section 4. |
| 11 to 13 | Section 2 | several fingers at one point still one-point (p. 1086); "I'm a child of the mouse" (p. 1090); 0.65 and 0.41 | C (3) | "even if the user casually touches with more than one finger at the same point, as our participants often did" (p. 1086). Arithmetic: 0.64+0.01; 0.36+0.04+0.01. |
| 14 to 18 | A1 stem, why, pos. 1 | screening; about 72%; "informed by user behavior" (p. 1083); "where the importance of immediate usability is high" (p. 1085); "desktop idioms strongly influence users' mental models" (p. 1083); "It's a Windows World" (p. 1090) | C (5) | All exact. |
| 19 | A1 pos. 1 | "designs around it: one-point gestures, widgets where agreement is low (p. 1091)" | **W** | p. 1091 recommends widgets beyond the first seven referents and for seven named commands, and says gestures "should not be distinguished by number of fingers." It nowhere recommends one-point gestures as a response to the desktop pull. |
| 20 to 29 | A1 pos. 2 to 4, zoom-in 1 | G and B p. 115 (sec. 10); Enlarge 3 pinch (p. 1089); "obtained from desktop and tabletop systems" (p. 1085); Table 1 (p. 1086); a two-hand gesture for every referent (p. 1086); "an interactive tabletop, Apple iPhone, or similar" (p. 1085); unrevised behavior (p. 1085); "children, Eastern, or uneducated participants would behave differently" (p. 1091); "to avoid bias" (p. 1085); Epps prompt cue (p. 1084) | C (10) | All exact. |
| 30 | A1 zoom-in 4 | "One participant tapped an object to select it, then gestured on the same object" | **W** | "In addition, some participants tapped an object first to select it, then gestured on top of the very same object, negating a key benefit of gestures that couples selection and action [13]" (p. 1090). The quote fragment is exact; the count is wrong. |
| 31 | A2 stem | each group "held identical gestures" (p. 1087) | C | "such that each group held identical gestures" |
| 32 | A2 stem | "The paper made that decision 54 times" | **U** | No page and not in the paper. 54 = 27 referents x 2 hand conditions is the presenters' arithmetic. It also miscounts: within each of the 54 sets, every pair of near-identical proposals needs its own same-or-different decision. |
| 33 to 37 | A2 why, pos. 1, pos. 2 | "adopted from prior work" (p. 1087); formula, two examples, range (p. 1087); both move referents at 1.0, then a cliff and a tail (Fig. 3, p. 1088); r = -0.52 (p. 1087); "identical" never defined (p. 1087) | C (5) | Figure 3 rendered: move a little and move a lot at 1.0, duplicate about 0.52. |
| 38 | A2 pos. 2 | "The authors grouped and coded alone, with no reliability check" | **U** | The paper says "The authors manually classified each gesture" (p. 1086) and "we grouped the gestures" (p. 1087). It reports no second coder and no reliability statistic. It does not say there was none. Pitfall 18 wording is "no inter-rater reliability reported." |
| 39 to 51 | A2 pos. 2 to 4, probes, move | "as our participants often did" (p. 1086); 1 to 3 fingers (p. 1089); "an important next step" (p. 1091); 57.0% (p. 1087); both widget recommendations (p. 1091); "Our user-defined set is useful, therefore, not just for what it contains, but also for what it omits" (p. 1088); eleven of twenty tapped, Figure 4 keeps tap and lasso; floor 0.05; tail about 0.21 task switch to about 0.11 cut (Fig. 3); "the referent with the largest group won the gesture" (p. 1087); N = 1074 (pp. 1087 to 1088); learnability and memorability, Wu et al. (p. 1091, ref. 35); Long et al. (p. 1084); finger count rarely matters (p. 1090) | C (13) | Rendered tail: task switch about 0.21, cut about 0.11, within 0.03. Two notes. "Eleven ... tapped" is inferred from Eq. 3 plus Figure 4 plus the largest-group rule; the paper does not say the eleven tapped. The phrase "the number of fingers rarely matters" is on p. 1088; p. 1090 supports it in substance ("only two said that the number of fingers actually mattered"). |
| 52 | A3 stem | "goes a long way to validating" (p. 1088) | C | Exact. |
| 53 | A3 stem | "The three authors predicted under half of what users produced." | **W** | "each author individually came up with only 43.5% of the user-defined gestures. Even combined, the authors only covered 60.9% of the users' set" (p. 1090). Together the three covered more than half. The denominator is the 48-gesture set, not "what users produced" (1,074 proposals). Conflicts with pitfall 14. |
| 54 to 64 | A3 pos. 1 to 4, zoom-in 1 | four contributions (p. 1084); G and B p. 111; W and K p. 40, p. 41; 43.5 and 60.9 (p. 1090); "a good candidate for deployment", "ease of recognition" (p. 1091); self-ratings (p. 1088); "three experts cannot generate the scope of gestures that 20 participants can" (p. 1090); "from signs to referents" (p. 1091); "wisdom of crowds" (p. 1090); "likely" (p. 859, via 09) | C (11) | All exact. |
| 65 | A3 zoom-in 2 | 'thinking creatively or hopelessly lost' (p. 1090) | **W** | "the authors are either thinking creatively or are hopelessly lost!" (p. 1090). The guide drops "either" and "are". The guide's own log and section 05 have it right. |
| 66, 67 | A3 zoom-in 3, 4 | ratings right after each gesture, before group sizes were known (p. 1088); Olsen's three steps (sec. 10) | C (2) | "Immediately after performing each gesture" (p. 1088); grouping "After all 20 participants had provided gestures" (p. 1087). |
| 68 to 76 | A4 stem, pos. 1 to 3 | "unrevised behavior"; "the vicissitudes of gesture recognition"; "a monologue in which the user's behavior is always acceptable" (p. 1085); Epps "may have biased" (p. 1084, fixed 11:56); "could not change previous gestures"; all referents first; "interactive systems that provide feedback will implicitly remind users to remain in contact with the table" (p. 1091); "to ensure that I was pressing", "the lack of feedback from the table" (p. 1090); two-hand for each referent, random order (p. 1086); 25 of 27, 17 of 48, "consistency and symmetry" (p. 1088) | C (9) | Two notes. The authors hedge: "This may be, at least in part, due to the lack of feedback" (p. 1090), and the guide's "tie to" drops the hedge. "the first attempt is the one that counts (p. 1085)" is the presenters' reading of "immediate usability," with a page attached. |
| 77 | A4 pos. 3 | participants' simplifications "allow us to unify the gesture set" (p. 1090) | **W** | "This allows us to unify the gesture set" (p. 1090). Not word for word. |
| 78 to 84 | A4 probes, threads, traps | "a common gesture was not feasible" (p. 1087); one participant, top hand about 3 inches (p. 1090, fixed 11:56); look down or sideways (p. 1091); reversible gestures though "the study software did not present these referents together" (p. 1090); insert and maximize about 0.58, the highest on the two-hand line, evenly divided, reused (pp. 1087 to 1088); touch instruction under Results (p. 1090); 13 of 20 (p. 1090) | C (7) | Figure 3 rendered: both two-hand peaks about 0.58, enlarge about 0.52; move a little two-hand 0.42. |
| 85 to 103 | A5 stem, pos. 1 to 4, probes | "would behave differently"; "a good candidate for deployment"; symbolic .100 (Fig. 2); check, X, ? (Fig. 4); widgets for accept, reject, help (p. 1091); aliasing (p. 1088); physical 43.9% (p. 1091); Next drawn left to right (Fig. 4); top-right close, bottom taskbar (p. 1090); "Oh, that's the same as enlarge" (p. 1090); Morris quote (p. 1084); Minimize 2 "seating position" (p. 1089); right-handed (p. 1085, Fig. 2 caption); .343, .274, .379, .450 (Fig. 2); "demands for precise positioning should be avoided" (p. 1091); easy gestures longer (p. 1088); 43.2, sd 15.6 (p. 1085); "likely" (p. 859); Furnas (ref. 8, p. 1088) | C (19) | Figure 2 values are printed labels, not chart readings. Next panel: the arrow runs from lower left to upper right. "Largest category" holds within the nature dimension only (object-centric binding is .702). "Smoother" is the authors' "perhaps." The Figure 2 interpretation is a problem; see section 5, Anchor 5. |
| 104 | A5 zoom-out | "a convenience sample of the public in Redmond" | **U** | "All were recruited from the general public" (p. 1085). No recruitment site is reported. Redmond is inferred from the Microsoft address (p. 1083). |
| 105 | A5 stem | "on a device built for groups" | **U** | The paper says tabletops are explored "for use by individuals [29] and groups [17]" (p. 1083). It says nothing about the Surface prototype being built for groups. |
| 106 to 115 | A6 pos. 1 to 5, probes | "an approach to designing tabletop gestures" (p. 1083); finger and four-finger guidelines (p. 1091); 25 of 27 (p. 1088); Table 2 (p. 1086); two of three at MSR (p. 1083); Surface prototype (p. 1085); "the first to employ users, rather than principles" (p. 1084); "goes a long way to validating" (p. 1088); "ease of recognition, consistency, reversibility, and versatility through aliasing" (p. 1091); "not, in fact, 'obvious'", 43.5% (p. 1091) | C (10) | All exact. |
| 116 | A6 pos. 4 and zoom-in 2 | "(p. 1091) was never tested"; "recognition was never tested" | **U** | The paper builds no recognizer and names one as a next step (p. 1091). Whether anyone later tested this set's recognition was not checked (section 11, Watch out). The guide's own Anchor 3 trap says "do not claim later work never did it." |
| 117 to 124 | Cards A to H, "Start here" | all pages and quotes | C (8) | All exact. Morris et al. 2006 is ref. 17 (p. 1092). |
| 125 to 129 | Thread map | Table 1 rated before the study (p. 1085), r = .71 (p. 1088); Malik 1, 2, 3, 5 fingers (p. 1085); Schuler and Namioka (ref. 25, p. 1084), "users are not designers" (p. 1085); 3.4 s, 2.0 s, "smoother or less hasty" (p. 1088); 1.0 for move (Fig. 3) | C (5) | All exact. |

Totals: 129 citations. 119 confirmed, 5 wrong (rows 19, 30, 53, 65, 77), 5 unverifiable (rows 32, 38, 104, 105, 116).

## 3. Outside-facts table

| Fact as the guide states it | Log | Verdict |
|---|---|---|
| Patent application US 2010/0031202: filed 2008-08-04, published 2010, abandoned 2014-12-09, three authors as inventors, on the procedure, not in the paper (A6 pos. 5, section 0) | progress.md; 03 log line 45 | Faithful. "Application," no motive, "facts only" rule kept. |
| GI 2010: 22 blind judges, 81 gestures, preferred larger-group gestures (A3 pos. 1) | progress.md; 04, 11 | Faithful. Position 1 omits "partly because the researchers' gestures were more complex," but position 2 and the traps carry it. Acceptable. |
| Morris et al. 2014: "legacy bias"; production, priming, partners (A1 pos. 2, A4 pos. 4, A6 pos. 4) | 11 log | Faithful. A6 pos. 4 reads "Two of the authors' 2014 article," which is ungrammatical (nice-to-have). |
| Vatavu and Wobbrock 2015: 0.32 becomes 0.284 (A2) | progress.md; 04 | Faithful. |
| DIS 2020: 216 studies; most cited as of 2019 (A6 stem) | 04 (Table 3, p. 859) | Faithful. |
| DIS 2020: "first of its kind" (A3 pos. 3) | 04: "the first application of the guessability method to hand-gesture input" | Slightly broadened. Section 11 uses the same loose phrase. A6 pos. 1 has the exact wording. Nice-to-have. |
| DIS 2020: "216 studies had used its method" (A6 stem); "216 later studies" (section 0) | 04: "N = 216 gesture elicitation studies" | Mild. The method predates this paper (2005), and the logs do not say whether the 216 includes this paper. Nice-to-have. |
| DIS 2020: 20 the norm "likely" because of this paper (section 0, A3 pos. 3, A5, A6 pos. 4) | 09: "likely explained by the fact that Wobbrock et al. [123] employed 20 participants" | Faithful after the 11:56 edit, except the **A6 stem**, "the field adopted its sample size as a norm." That turns 36% "most common" into "a norm" and drops "likely." Should-fix. |
| W and K: "repeated validation" p. 41; artifacts need not come with studies p. 40; 2016 article does not name this paper, 2012 whitepaper does | 10 log | Faithful. |
| G and B: p. 115 and p. 111 quotes | 10 log | Faithful, exact. |
| Olsen: three steps, "empowering new design participants" | 10 Watch out | Faithful. |
| Weiser: menu-bar doubt (A4 threads) | 10 | Faithful. |
| Weiser: "wanted computers that disappear" (A4 zoom-out) | Not in 10 or any log | Not logged. The line is famous, but the brief requires the logs. Should-fix: use the logged "casual use." |
| Japan circle and check; Bulgaria nod (A5 pos. 2) | 06 log | Faithful. |
| Devices: trackpads, Palm Graffiti, ticket machines, Nintendo DS (A1 pos. 3) | 06 log (1994, 1996, 1999, 2004) | Faithful. Undated in the guide, so nothing is strengthened. |
| iPhone shipped June 2007 with pinch (A1 pos. 2) | 11 log | Faithful. |
| "Pinch became universal because Apple shipped it" (A1 pos. 2 and card H) | 11 says "pinch became universal anyway." Its Watch out concedes "the study cannot tell the two apart" | In A1 it is one position among others, so acceptable. Card H states it as the premise of the question. Nice-to-have. |
| BMW 7 Series 2015 (A1 zoom-out) | progress.md | Faithful. |
| Mauney et al. 2010 [UNVERIFIED] (A5 traps) | 06 | Faithful; the tag is kept. |
| EdgeWrite 2005 "measured guessability directly" (A2 pos. 3) | 04 Watch out (abstract only) | Faithful. |
| sd arithmetic, under about 28 or over about 58 (A5 zoom-in) | 06 log (15.6 x sqrt(19/20) = 15.2) | Faithful; labelled "our arithmetic." |
| "Redmond" (A5 zoom-out) | No log | Not logged and not in the paper. Should-fix. |

## 4. Pitfalls conflicts

| Pitfall | Where | Conflict |
|---|---|---|
| 14 | A3 stem | "The three authors predicted under half" contradicts "all three combined 60.9%." **Blocker.** |
| 11 | Section 2, first bullet | The Figure 4 caption is introduced as "the paper's own rule" during a vote on how to *group*. The caption describes the final set. The paper does not say how finger count was treated in grouping. Anchor 2 pos. 2 handles this correctly ("the paper does not say"); section 2 undercuts it. Should-fix. |
| 18 | A2 pos. 2 | "with no reliability check" should be "no reliability check reported." Should-fix. |
| 9 | A4 pos. 1 | Fixed in the 11:56 version ("conceded ... 'may have biased'"). No conflict now. |
| 6 | A3 pos. 1 | No conflict; the caveat sits in position 2 and the traps. |
| 1, 2, 3, 4, 5, 7, 12, 13, 15, 16, 17, 19, 20 | throughout | No conflict. Trap lines repeat pitfalls 1, 2, 5, 6, 14 and 15 correctly. |

## 5. Per-anchor verdicts

Each anchor was judged on five points:
- the five guidance principles quoted at the top of the guide;
- the requirement that questions "allow some deeper discussions and participants to elaborate properly about the details and express their view points";
- whether the stem invites elaboration;
- whether there are at least three distinct defensible positions, one of them a defence of the paper;
- whether the probes reach both the specific evidence (zoom in) and its significance (zoom out).

**Anchor 1, Whose gestures are these? STRONG.**
- **Stem.** Opens on the room's own data and a paper number, and asks "whose," which cannot be answered yes or no.
- **Positions.** Four genuinely distinct ones: build on it; it locks in the past; it measures the method; it is a snapshot. Position 1 is a fair defence, grounded in p. 1085.
- **Probes.** Zoom-in probes reach a page each. Zoom-out probes reach today's devices.
- **Flaws.** The stem asks two questions at once ("whose habits ... and is that what a designer should build on"), and the three-way hands only answer the first. Zoom-in 4 miscounts "one participant."
- **Single most valuable improvement:** ask only "whose habits is it recording?" for the hands vote. Hold "should a designer build on them?" as the follow-up once the three groups have spoken. This gives the descriptive and the normative question their own turns.

**Anchor 2, What does an agreement score tell a designer? STRONG.**
- **Stem.** Starts from the two numbers the room just produced. This is the best concrete opener in the guide.
- **Positions.** Distinct: sound summary; hidden grouping step; consensus is not guessability; the low end is the useful output. Position 4 is an unusual and well-evidenced defence.
- **Probes.** Precise (0.37, the floor, the conflict rule, N = 1074).
- **Flaws.** The stem's "54 times" is presenters' arithmetic and undercounts the decisions. Position 3 is abstract compared with the others.
- **Single most valuable improvement:** fix the "54 times" sentence (should-fix S2). Then give position 3 a live test: ask one person who missed the warm-up to guess which command the room's winning gesture was for. That shows guessability is a different measurement from agreement.

**Anchor 3, What would convince you that the method works? ADEQUATE as written; STRONG once the stem is fixed.**
- **Stem.** The best-designed stem for elaboration ("what would it take").
- **Positions.** Four strong, distinct ones, including a clean defence via Greenberg and Buxton and Wobbrock and Kientz.
- **Move.** The board sort is excellent: it makes students name a standard before judging the paper against it.
- **Flaws.** The stem misstates the designer numbers (blocker B1). A sharp classmate will correct it in the first minute, and the facilitator loses the room's trust on the anchor about evidence. "The guessability test the authors proposed is not in the paper" is also ambiguous: it is in the paper as a proposal.
- **Single most valuable improvement:** blocker B1.

**Anchor 4, Was the monologue the right experiment? STRONG.**
- **Stem.** "What did that choice buy, what did it cost, and how would you have run it" invites design thinking.
- **Positions.** Distinct and well sourced, including the paper's own concessions (p. 1091).
- **Probes.** The insert and maximize probe (about 0.58, yet no gesture) is the sharpest evidence probe in the guide.
- **Flaws.** The Move splits the room in two, which breaks the guide's own rule "Hands in three, never two" (section 6). Position 2 states an inference as fact without "we argue."
- **Single most valuable improvement:** make the seat split three-way, with the middle third building the hybrid (should-fix S15). That turns position 4 from an afterthought into a group's job.

**Anchor 5, Who was not at the table? ADEQUATE.**
- **Stem.** Well aimed ("where does that limitation stop," "which parts of Figure 4 survive").
- **Move.** The Figure 4 round is the best whole-class participation device in the guide.
- **Flaws in the evidence.** Position 4 rests on a misreading. Figure 2's left-hand column covers two-hand gestures only. The right-hand column is the one whose .450 + .274 gives the paper's "about 72% of gestures," so it covers all gestures. The difference is therefore partly a 1-hand against 2-hand difference, and at most says something about right-handers' non-dominant hand, not about left-handed users. The 11:56 edit added "in two-hand gestures," which is still not accurate. Position 2 also states two inferences unlabelled. The zoom-out invents "Redmond." The stem asserts a "device built for groups."
- **Single most valuable improvement:** rebuild position 4 on the sound evidence and label the inference (should-fix S7). The sound evidence is: all right-handed; no age breakdown; the precise-positioning warning; corner drags and drawn marks.

**Anchor 6, Seventeen years on. ADEQUATE.**
- **Positions.** Five distinct ones, including position 5 under the facts-only rule. The zoom-in probes are good: the four contributions against the seven types, and the four Conclusion claims.
- **Flaws.** The stem stacks four premises before the question. Three of them (no product shipped, pinch anyway, the sample-size norm) point toward position 4, so the stem leans before anyone speaks. One premise drops the review's "likely." The closing round cannot fit five minutes, and the three-minute version ("the stem and the closing round only") cannot fit three.
- **Single most valuable improvement:** balance the stem with one fact for the paper to go with the facts against it. Then cap the spoken round at eight to ten voices and have everyone else write their sentence on the back of their card (should-fix S13). For example, "the first application of the guessability method to hand gestures and the most cited."

## 6. Card review

| Card | Real disagreement? | Fair and workable? | Note |
|---|---|---|---|
| A | Yes. The visitor needs the familiar; the designer wants the new. Maps onto A1 pos. 1 against 2. | Yes | Strong. |
| B | Yes. Rigour against shipping. | Yes | The statistician may notice that A is not a statistical test; that is a good thing for them to say. |
| C | Yes. Reviewer against author. | Yes | Strong. The author has p. 1091 "next step" and p. 1088 to work with. |
| D | Yes. Feedback against clean data. | Yes | The participant role rests on one quote; the "Start here" column supplies three. |
| E | Yes. Local marks mean more aliases; the engineer resists. | Workable only with Figure 4 printed on the card or shown on screen | Print Figure 4 at the bottom of card E. |
| F | Yes | Yes | Morris is first author of ref. 17 and second author here, so the "two roles" are one person's two papers. Worth a line on the card; it makes the disagreement vivid. |
| G | Weak. A researcher who used the method and a practitioner building a gesture set can easily agree that "the method" is the credit. | Fair | Give role 2 a stance: a practitioner who says only the concrete guidelines on p. 1091 helped (A6 pos. 2). Nice-to-have N15. |
| H | Yes | Fair | The stem asserts the causal claim as fact; section 11 concedes the study cannot separate market from elicitation. Nice-to-have N7. |

No role is likely to embarrass a student: none asks anyone to speak as themselves about age, culture or ability, and the guide says roles may be dropped. Card A's "70, no smartphone" is a role, not a person in the room. The one practical risk is that eight cards need sixteen students. The guide already says to use A to F with fewer.

## 7. Timing review (section 5)

The group phase adds up to 2 + 6 + 6 + 5 + 1 = 20 minutes; the arithmetic is right. Realism, segment by segment:

- **Warm-up, 0:00 to 5:00.** Not realistic as written. The plan has two referents, each with an effect shown, a performance and a live tally of about 20 gestures, then a grouping vote, A computed twice, circling, and noting names. Tallying one referent live takes 1.5 to 2 minutes. Expect 7 to 8 minutes. Fix: take move a little by show of hands only (the expected result is one large group), and spend the full tally and the grouping vote on undo or help (should-fix S14).
- **Pairs, 5:00 to 10:00.** Tight but workable. Handing out cards and reading them costs about a minute, and writing two lines with a page or section each is demanding in four. Accept "page or section or figure" and say at the start that one line is enough if time runs out.
- **Harvest, 10:00 to 12:00.** Three or four pairs reading one line each fits. Adding the board threads makes it tight; write the threads during the pair phase instead.
- **Anchor 1, 12:00 to 18:00.** Fits: stem, three-way hands, one voice per group, two or three exchanges.
- **Anchor 3, 18:00 to 24:00.** Fits only just. The 90-second sentence completion, five answers and the board sort take about 4 minutes, leaving 2 for argument. Acceptable.
- **Anchor 6, 24:00 to 29:00.** Does not fit. With 16 to 20 students, one sentence each with no repeats is 4 to 5 minutes on its own, before the stem and the tally. Cap at 8 to 10 voices and collect the rest in writing (should-fix S13).
- **Close, 29:00 to 30:00.** Fine.

**Three-minute versions:**

| Anchor | Usable? | Note |
|---|---|---|
| A1 | Yes | |
| A2 | Yes | |
| A3 | Barely | The 90-second completion plus the board sort fill the three minutes; drop the board sort and ask for three sentences aloud. |
| A4 | Yes | |
| A5 | Yes | Four panels only. |
| A6 | No | "The stem and the closing round only" cannot fit; see S13. |

**Buffer.** The "about 4 minutes" (25:45 talk in a 30-minute slot) exists only if the talk runs to time. Treat it as contingency, not as plan.

## 8. Missing questions (up to three)

1. **Should the recognizer bend to the user, or the user to the recognizer?** The guide touches recognition only through "ease of recognition." The paper sets up a real tension:
   - designer gestures are "chosen out of concern for reliable recognition [19]" (p. 1083);
   - Long et al. found "users are sometimes poor at picking easily differentiable gestures" (p. 1084);
   - the set has 48 gestures for 22 referents, with aliases (p. 1088);
   - physical gestures (43.9%) support "a physics engine ... rather than a traditional recognizer": seven participants expected objects to move out of the way and four "threw" objects, yet "a physics engine alone will probably not suffice" (p. 1091);
   - hit-testing within objects and special corners (p. 1091);
   - no recognizer was built (p. 1091).

   Positions: user first; recognizer first; physics engine as the third way. It fits section 08 and A6 pos. 2.
2. **Whose idea of "complex"?** Every main correlation uses the authors' complexity ratings:
   - complexity against agreement, r = -.52 (p. 1087);
   - complexity against the nature categories, 2.11 to 3.52 (p. 1087);
   - complexity against planning time, r = .71 (p. 1088);
   - complexity against goodness, r = -.59 (p. 1088).

   The three authors rated complexity before the study (p. 1085; Table 1, with an SD of 1.53 for duplicate, p. 1086), and the paper concedes "participants and the authors did not always regard the same referents as 'complex'" (p. 1090). This is a sharp question about where designer judgment re-enters a user-defined method, and it sits naturally inside Anchor 2 or 3.
3. **Can you code a gesture's meaning from its form?** The taxonomy, which A6 pos. 3 offers as a lasting gift, has a nature dimension that "usually is not enough to reveal its metaphorical nature; the answer lies in the user's mental model" (p. 1086). Several questions follow:
   - the authors coded it manually (p. 1086) and resolved multi-finger cases by debriefing (p. 1086);
   - the set's breakdown differs from all proposals by 6.7 percentage points on average (p. 1088).

   Can another lab apply this taxonomy without think-aloud data? Does that limit its value as "a shared language"? This adds a specific-evidence probe to Anchor 6, which is currently the anchor most at the level of significance.

## 9. Style findings

- **Dashes.** No em-dashes and no en-dashes (byte search, 0 hits). Ranges use "to."
- **"The paper says" against "we argue."** Mostly well kept. Unlabelled inferences that read as fact:
  - A1 pos. 1, "designs around it: one-point gestures";
  - A4 pos. 2, "So part of the finger finding ... is an artefact of the setup. People adapt";
  - A5 pos. 1, "aliasing lets a locale add its own mark";
  - A5 pos. 2, "is a desktop education speaking" and "what counted as consensus is cultural too";
  - A5 pos. 4, the left-handed inference from Figure 2.

  Positions are argument briefs, so lighter labelling is defensible, but each of these sits beside a page number and will be read as the paper's.
- **Motives.** None stated as fact. A6 pos. 5 and its traps are careful ("motive is not").
- **Rhetorical questions.** Acceptable here; this is a discussion guide.
- **Small inconsistencies.**
  - The section 1 table says "Anchors 4 and 5 have a 60-second neighbour split." Only Anchor 4 has one; Anchor 3 has the pair sentence completion, and Anchor 5 has a round.
  - The purpose line says "about 40 probes"; there are 35 zoom-in and zoom-out probes.
  - A5 threads send "aliases and recognizers" to Anchor 2, position 3, which is about guessability, not aliases.
- **Grammar.** A6 pos. 4: "Two of the authors' 2014 article named" is ungrammatical. A3 pos. 3: 'which the review "likely" traces' is awkward.

## 10. Corrections (exact replacements, priority order)

### Blockers (3)

**B1. Anchor 3 stem (pitfall 14; also removes the proposal ambiguity).**

OLD:
```
The three authors predicted under half of what users produced. The guessability test the authors proposed is not in the paper.
```
NEW:
```
Each author, designing alone, came up with only 43.5 percent of the users' final set; all three together, 60.9 percent. The guessability test the authors proposed was not run in this paper.
```

**B2. Anchor 1, zoom-in 4 (count wrong).**

OLD:
```
"One participant tapped an object to select it, then gestured on the same object,
```
NEW:
```
"Some participants tapped an object to select it, then gestured on the same object,
```

**B3. Anchor 3, zoom-in 2 (misquote).**

OLD:
```
19.1% of each author's gestures were never tried by any participant: 'thinking creatively or hopelessly lost' (p. 1090).
```
NEW:
```
19.1% of each author's gestures were never tried by any participant: the authors are 'either thinking creatively or are hopelessly lost!' (p. 1090).
```

### Should-fix (15)

**S1. Section 2, first bullet (pitfall 11).**

OLD:
```
The paper's own rule, in the Figure 4 caption, is that a gesture drawn with one finger "could be performed with 1-3 fingers" (p. 1089);
```
NEW:
```
The paper does not say how it grouped different finger counts (p. 1087; pitfall 11). What it does say is that in the final set a gesture drawn with one finger "could be performed with 1-3 fingers" (Figure 4 caption, p. 1089);
```

**S2. Anchor 2 stem ("54 times").**

OLD:
```
The paper made that decision 54 times for 1,074 gestures and tells us only that each group 'held identical gestures.'
```
NEW:
```
The paper had to make decisions like that across 1,074 gestures, sorted within 27 referents for one hand and for two, and tells us only that each group 'held identical gestures.'
```

**S3. Anchor 2, position 2 (pitfall 18).**

OLD:
```
The authors grouped and coded alone, with no reliability check (pp. 1086 to 1087; section 09).
```
NEW:
```
The authors grouped and coded the gestures themselves, and the paper reports no second coder and no reliability check (pp. 1086 to 1087; section 09).
```

**S4. Anchor 4, position 3 (misquote).**

OLD:
```
participants' simplifications "allow us to unify the gesture set" (p. 1090)
```
NEW:
```
the paper says participants' simplified mental models mean "This allows us to unify the gesture set" (p. 1090)
```

**S5. Anchor 4, position 2 (hedge restored, inference labelled).**

OLD:
```
which the authors tie to "the lack of feedback from the table" (p. 1090). So part of the finger finding, and perhaps the above-table and off-screen behaviour, is an artefact of the setup. People adapt; the data say nothing about the second use.
```
NEW:
```
which the authors say "may be, at least in part, due to the lack of feedback from the table" (p. 1090). We argue part of the finger finding, and perhaps the above-table and off-screen behaviour, is an artefact of the setup. People adapt (common experience, not a finding of the paper); the data say nothing about the second use.
```

**S6. Anchor 1, position 1 (attribution wrong).**

OLD:
```
and designs around it: one-point gestures, widgets where agreement is low (p. 1091).
```
NEW:
```
and designs for it: a widget as well as a gesture where agreement is low, and widgets for commands where participants tapped imaginary ones (p. 1091).
```

**S7. Anchor 5, position 4 (Figure 2 misread; inference labelled).**

OLD:
```
and in two-hand gestures the two hands differ in Figure 2: the left hand made more one-point touches and fewer one-point paths than the right (about .343 against .274 and .379 against .450, chart readings).
```
NEW:
```
and the paper reports no result by handedness or age. Figure 2's left-hand column covers only two-hand gestures (caption, p. 1087), so its difference from the right hand (.343 against .274 one-point touches, printed labels) describes right-handers' other hand, not left-handed users; that left-handers would differ is our inference.
```

**S8. Anchor 5, position 2 (inferences labelled).**

OLD:
```
"Oh, that's the same as enlarge" (p. 1090) is a desktop education speaking. And the agreement score was computed inside one culture, so what counted as consensus is cultural too.
```
NEW:
```
We argue "Oh, that's the same as enlarge" (p. 1090) may be a desktop education speaking; the paper reads it as a simplifying assumption. And the agreement score was computed inside one culture, so, we argue, what counted as consensus is cultural too.
```

**S9. Anchor 5, zoom-out 1 (Redmond not in the paper).**

OLD:
```
Is a convenience sample of the public in Redmond a stand-in for 'the public'?
```
NEW:
```
Are twenty paid volunteers 'recruited from the general public' (p. 1085), recruitment site not reported, a stand-in for 'the public'?
```

**S10. Anchor 6, position 4 and zoom-in 2 ("never tested"; two replacements).**

OLD:
```
(p. 1091) was never tested.
```
NEW:
```
(p. 1091) was not tested in the paper; no recognizer was built.
```
OLD:
```
aliasing is in the set; recognition was never tested.)
```
NEW:
```
aliasing is in the set; recognition was not tested in the paper.)
```

**S11. Anchor 6 stem ("likely" dropped).**

OLD:
```
And the field adopted its sample size as a norm.
```
NEW:
```
And 20 participants became the field's most common sample size, the review says likely because this paper used 20.
```

**S12. Anchor 4, zoom-out 1 (Weiser claim not in any log).**

OLD:
```
Weiser wanted computers that disappear (section 10). Does a gesture you have to guess count as disappearing?
```
NEW:
```
Weiser wanted boards for walk-up "casual use," with no sign-up (section 10). Does a gesture you have to guess count as casual use?
```

**S13. Anchor 6 closing round and its three-minute version (timing; two replacements).**

OLD:
```
**Move:** the closing round. Each person, one sentence: "the one thing I would change if I ran this study today." No repeats.
```
NEW:
```
**Move:** the closing round. Eight to ten people aloud, one sentence each: "the one thing I would change if I ran this study today." No repeats. Everyone else writes their sentence on the back of their card, collected at the door.
```
OLD:
```
**Three-minute version:** the stem and the closing round only.
```
NEW:
```
**Three-minute version:** the stem and five volunteers' sentences; the rest written on the cards.
```

**S14. Section 5, warm-up row (timing).**

OLD:
```
Move a little, then undo or help. Tally on the board. One grouping vote; A computed both ways.
```
NEW:
```
Move a little by show of hands only (expect one large group), then undo or help with a full tally on the board. One grouping vote on the undo tally; A computed both ways.
```

**S15. Anchor 4 Move (two-way split breaks the guide's own "hands in three" rule).**

OLD:
```
split the room by seat. Left half argues "first instinct is the product," right half "learned behaviour is the product," 60 seconds with a neighbour, two voices each. Then ask someone who has not spoken to propose the hybrid.
```
NEW:
```
split the room by seat in three. Left third argues "first instinct is the product," right third "learned behaviour is the product," middle third designs the hybrid study; 60 seconds with a neighbour, one or two voices from each third, the hybrid last and from someone who has not spoken.
```

### Nice-to-have (16)

**N1. Anchor 2 Move.**

OLD:
```
(The paper says finger count rarely matters, p. 1090; the joke lands.)
```
NEW:
```
(The paper says the number of fingers rarely matters, pp. 1088, 1090; the joke lands.)
```

**N2. Anchor 2, zoom-in 1.**

OLD:
```
Select single: eleven of twenty tapped;
```
NEW:
```
Select single: eleven of twenty made one gesture, a tap by Figure 4;
```

**N3a. Anchor 5, position 1.**

OLD:
```
Physical gestures were the largest category, 43.9% (p. 1091)
```
NEW:
```
Physical gestures were the largest nature category, 43.9% (p. 1091)
```

**N3b. Anchor 5, position 1.**

OLD:
```
and aliasing lets a locale add its own mark (p. 1088)
```
NEW:
```
and aliasing (p. 1088) would, we argue, let a locale add its own mark
```

**N4. Anchor 5, position 4.**

OLD:
```
found that gestures rated easy were the longer, smoother ones
```
NEW:
```
found that gestures rated easy took longer, "perhaps because they were smoother or less hasty"
```

**N5. Anchor 5 threads.**

OLD:
```
"aliases and recognizers" to Anchor 2, position 3, and section 08
```
NEW:
```
"aliases and recognizers" to the aliasing zoom-out above and section 08
```

**N6. Anchor 6, position 4 (grammar).**

OLD:
```
Two of the authors' 2014 article named legacy bias
```
NEW:
```
A 2014 article by two of the authors and colleagues named legacy bias
```

**N7. Card H.**

OLD:
```
Pinch became universal because Apple shipped it. Does that make elicitation pointless, or prove it right?
```
NEW:
```
Pinch is in this set, and it became universal after Apple shipped it. Does that make elicitation pointless, or prove it right?
```

**N8. Anchor 3, position 3.**

OLD:
```
called this paper the first of its kind and the most cited
```
NEW:
```
called this paper the first application of the guessability method to hand gestures and the most cited
```

**N9. Anchor 6 stem.**

OLD:
```
by 2020, 216 studies had used its method.
```
NEW:
```
by 2020, a review counted 216 gesture elicitation studies.
```

**N10. Purpose line.**

OLD:
```
about 40 probes
```
NEW:
```
35 probes
```

**N11. Anchor 4, position 1.**

OLD:
```
For walk-up public use, the first attempt is the one that counts (p. 1085).
```
NEW:
```
The paper ties this to public tables, where "the importance of immediate usability is high" (p. 1085); we read that as: the first attempt is the one that counts.
```

**N12. Anchor 5 stem.**

OLD:
```
tested one at a time, on a device built for groups.
```
NEW:
```
tested one at a time, on a table that groups can share.
```

**N13. Section 1 table.**

OLD:
```
Anchors 4 and 5 have a 60-second neighbour split inside the group phase
```
NEW:
```
Anchor 3 has a 90-second pair sentence completion and Anchor 4 a 60-second neighbour split inside the group phase
```

**N14. Anchor 3, position 3 (grammar).**

OLD:
```
which the review "likely" traces to this paper (p. 859)
```
NEW:
```
which the review says is "likely explained" by this paper's choice of 20 (p. 859)
```

**N15. Card G, role 2.**

OLD:
```
| A practitioner designing an AR gesture set this year |
```
NEW:
```
| A practitioner designing an AR gesture set this year, who says only the guidelines on p. 1091 helped |
```

**N16. Anchor 1 stem (single most valuable improvement for A1).** Keep the stem's text. In the Move, change the order:
- ask "whose habits is it recording?" for the three-way hands;
- after one voice per group, ask "and should a designer build on them?"

No text replacement is proposed. This is a facilitation change for the presenter to decide.
