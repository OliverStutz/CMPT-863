# Opus E: cross-check of `slides/10-Connections.md`

Written 2026-09-30 by an independent verifier (Brief 12). `10-Connections.md` was not edited.

## Sources opened by the verifier

All six URLs in the file were downloaded and read in full (text extraction plus a visual render where layout mattered).

| Reading | URL opened | What it is | Page mapping used |
|---|---|---|---|
| Wobbrock and Kientz 2016 | https://faculty.washington.edu/wobbrock/pubs/interactions-16.pdf | Published Interactions 23(3) layout, 7 PDF pages | PDF p. 1 = printed 38 ... PDF p. 7 = printed 44 (footers checked) |
| Wobbrock 2012 whitepaper | https://faculty.washington.edu/wobbrock/pubs/Wobbrock-2012.pdf | "Seven Research Contributions in HCI", 6 pages | printed page numbers on the PDF; p. 3 rendered and viewed |
| Greenberg and Buxton 2008 | https://groups.cs.umass.edu/wp-content/uploads/sites/8/2019/01/reading18.pdf | CHI 2008 proceedings layout, 10 pages | PDF p. 1 = printed 111 ... PDF p. 10 = printed 120 (footers checked) |
| Olsen 2007 | http://web.archive.org/web/20120601131116/http://icie.cs.byu.edu:80/ResearchGuidance/SystemsResearchCriteria.pdf | 7-page SIGCHI-format copy, PDF CreationDate 2007-03-27 | PDF page numbers (no proceedings pages) |
| Weiser 1991 | https://www.socs.uoguelph.ca/~qmahmoud/teaching/fall2006/pervasive/computer-21century.pdf | Word reformat, 11 PDF pages, no journal pages | PDF page numbers |
| Weiser 1991 (cross-check) | https://graphics.stanford.edu/courses/cs428-03-spring/Papers/readings/General/Weiser_SciAm91.htm | HTML copy | text identical for every quoted sentence |
| Weiser citation | https://api.crossref.org/works/10.1038/scientificamerican0991-94 | Crossref record | Scientific American 265(3), pp. 94 to 104 |
| Furnas et al. 1987 | https://honnef.co/notes/references/furnasvocabularyproblemhumansystem1987 | BibTeX with abstract | abstract only |
| The paper | `Week 4 - User-Defined Gestures for Surface Computing.pdf` | PDF p. N = printed 1082 + N | |

## Summary

| Status | Count |
|---|---|
| CONFIRMED | 56 |
| WRONG | 9 (4 distinct errors, repeated across table, slide and notes) |
| UNVERIFIABLE | 1 |

The quotes are accurate. Every quoted string from the four readings and from the paper was found verbatim on the page given. The errors are in how quotes are paired: twice the reading's point is set against a paper quote that does not support it (validation of the method vs validation of the set; "early" vs "a good candidate for deployment"). One Weiser equivalence is overstated, and Olsen's three steps are conflated with his list of claims.

Style: no em-dashes or en-dashes in the file. On-slide text is 88 words, under 90. "Say" is 186 words, within the 195-word budget for 1:30. Most judgments carry "we argue". Two in the Say text do not (Olsen "That is this method's claim", "is still open"), and neither does the Weiser equivalence. No motive is attributed to the authors. The Buxton/Microsoft note explicitly says "do not suggest any influence".

## A. Claims for the cross-checker (12 rows)

| # | Verdict | Evidence |
|---|---|---|
| 1 | **WRONG (pairing)** | The reading quotes are exact: "Methodological research contributions are evaluated on the utility, reproducibility, reliability, and validity of the new method or method enhancement. ... Knowing that a method has these properties requires repeated validation." (Interactions p. 41, PDF p. 4). The paper quote is exact: "An important next step is to validate our user-defined gesture set." (p. 1091). But the next step validates the **gesture set** (the artifact). Wobbrock and Kientz's "repeated validation" concerns the **method**. On the method, the paper does claim validation from one study: "This finding goes a long way to validating this user-driven approach to gesture design." (p. 1088). So "this paper defers validation" is wrong for the method. The accurate point is that the method's validation rests on one study's own ratings, and the set's validation is deferred. |
| 2 | CONFIRMED | Whitepaper p. 3, rendered and viewed. Under "Examples of Methodological Contributions" it lists "Wobbrock, J.O., Morris, M.R. and Wilson, A.D. (2009). User-defined gestures for surface computing ... 1083-1092", next to "Wobbrock, J.O., Aung, H.H., Rothrock, B. and Myers, B.A. (2005). Maximizing the guessability of symbolic input" (= paper ref. 33, p. 1092). Interactions p. 40: "In 2012, the first author posted an unpublished whitepaper ... [3] ... Here we summarize those contribution types". Endnote 3, p. 44, gives Wobbrock-2012.pdf. |
| 3 | CONFIRMED | Seven types are the section heads on Interactions pp. 40 to 42, and the comparison table on p. 43 has a "Wobbrock (2012)" column. Paper p. 1084: "This work contributes the following ... (1) a quantitative and qualitative characterization of user-defined surface gestures, including a taxonomy, (2) a user-defined gesture set, (3) insight into users' mental models ..., and (4) an understanding of implications for surface computing technology and user interface design." No method is listed. The abstract (p. 1083) says "We present an approach to designing tabletop gestures that relies on eliciting gestures from non-technical users". |
| 4 | CONFIRMED | G&B p. 115 (PDF p. 5): "we eliminate ideas too early, we consider far too few ideas at all, we converge on that which we can measure, which is almost always that which we are already familiar with." Paper p. 1090: "In all, about 72% of gestures were mouse-like one-point touches or paths." The analogy is ours. G&B's mechanism is measurement bias, not user habit. The notes label it "We argue", which is fine. |
| 5 | **WRONG (pairing)** | The reading quotes are exact. p. 111 abstract: "If done during early stage design, it can mute creative ideas that do not conform to current interface norms." p. 118: "Examples include a design rationale, a vision of what could be, expected scenarios of use, reflections, case studies, participatory critique, and so on." But the paper-side evidence for an "early candidate set" is "a good candidate for **deployment** in tabletop systems" (p. 1091). That positions the set as ready to ship, not as an early-stage sketch. G&B's leniency is for early design ("Early designs are best considered as sketches", p. 113 region; "usability evaluation has a significant role to play when conditions warrant it", p. 111). The quote therefore weakens the "not automatically a flaw" reading rather than supporting it. |
| 6 | CONFIRMED (quotes) | Olsen pre-print p. 5, "Empowering new design participants": "Tools can also make a contribution by introducing new populations to the UI design process." ... "The last claim is that the resulting UI designs are "better" when the new tools are used by this population. The question of defining "better" is problematic." Paper p. 1084 (non-technical recruits), p. 1083 ("chosen out of concern for reliable recognition [19]"), p. 1090 ("each author individually came up with only 43.5%"), p. 1088 ("goes a long way to validating"). See D4 for the separate error in how the notes describe Olsen's steps. |
| 7 | CONFIRMED | Olsen pre-print p. 1, "Forces for change": "Cameras and touch tables produce inputs that are the size of a hand or finger rather than a point, yet we force such techniques into the standard mouse point model because that is all that our systems support." Paper p. 1090: "it seemed that touches with 1-3 fingers were considered a "single point"" (note the hedge "it seemed"); "about 72%" also p. 1090. |
| 8 | **WRONG (overstated)** | The quotes are exact. Guelph PDF p. 4: "Boards built by Richard Bruce and Scott Elrod at PARC currently measure about 40 by 60 inches and display 1024x768 black-and-white pixels." PDF p. 5: "We have built enough liveboards to permit casual use: they have been placed in ordinary conference rooms and open areas, and no one need sign up or give advance notice before using them." Paper p. 1085: "Microsoft Surface prototype measuring 24" × 18" set at 1024 × 768 resolution"; "interactive tabletops may be used in public spaces, where the importance of immediate usability is high." But Weiser's "no one need sign up" is about booking and access. The paper's "immediate usability" is about use without learning. Calling them "the same requirement" overstates the link. Weiser's own nearby term is "casual use". |
| 9 | CONFIRMED | Guelph PDF p. 4: "wireless electronic "chalk" that can work either in contact with the surface or from a distance." PDF p. 5: "not everyone can reach the top of the board, so a Macintosh-style menu bar may not be a good idea." Paper p. 1091: "tabletop systems may benefit from the ability to look down or sideways at users' hands, rather than just up." Paper p. 1090: "imaginary widgets located at objects' top-right and the screen's bottom". The notes say the paper "recommends" sensing; the paper says "may benefit". Minor, see D7. |
| 10 | CONFIRMED | Furnas abstract: "In every case two people favored the same term with probability <0.20 ... An optimal strategy, unlimited aliasing, is derived". Paper p. 1088: "Aliasing has been shown to dramatically increase input guessability [8,33]". p. 1091: "providing multiple gestures for the same commands, which enhances guessability [8,33]". Ref. [8] is Furnas et al. 1987, CACM 30(11), 964-971 (p. 1092). |
| 11 | CONFIRMED | Paper p. 1085: "we sought to remove the gulf of execution [10] from the dialogue, creating, in essence, a monologue in which the user's behavior is always acceptable." Ref. [10] is Hutchins, Hollan and Norman 1985, "Direct manipulation interfaces" (p. 1092). |
| 12 | CONFIRMED | Paper p. 1084: "Morris et al. [17] offer a classification of cooperative gestures among multiple users at a single interactive table ... These dimensions address groups of users and omit issues relevant to single-user gestures, which we cover here." Figure 4, p. 1089: "Minimize2: Use Move2 to move object to the bottom of the surface (as defined by user's seating position)." Ref. [17] is Morris, Huang, Paepcke, Winograd, CHI '06 (p. 1092). |

## B. On the slide (5 bullets)

| Bullet | Verdict | Evidence |
|---|---|---|
| Wobbrock and Kientz | **WRONG (pairing)** | Same as row 1. "repeated validation" (p. 41) is exact. "next step" (p. 1091) is exact but refers to the set, not the method. |
| Greenberg and Buxton | CONFIRMED | p. 115 quote exact. 72% at p. 1090. |
| Olsen | CONFIRMED | "new design participants" appears in quotation marks in Olsen pre-print p. 5. Non-technical recruits at p. 1084. |
| Weiser | **WRONG (overstated)** | 1024x768 is confirmed for both. "same need for 'immediate usability'" makes the paper's term look like Weiser's requirement. See row 8. |
| Roots | CONFIRMED | p. 1088, p. 1085, p. 1084 as in rows 10 to 12. |

## C. Speaker notes

**Say** (10 statements)

| Statement | Verdict | Evidence |
|---|---|---|
| "Wobbrock and Kientz sort research into seven contribution types." | CONFIRMED | Interactions pp. 40 to 42 |
| "We argue the set is an artifact, the findings are empirical, but the core is a method." | CONFIRMED (labelled judgment) | |
| "Wobbrock's own 2012 whitepaper lists this exact paper as a methodological example." | CONFIRMED | whitepaper p. 3 |
| "And methods, they write, are judged on validity, which "requires repeated validation." This paper leaves validation as a next step." | **WRONG** | Three problems. (i) "they" comes right after the single-author whitepaper, so it is ambiguous. (ii) The quote is about all four properties, not validity alone. (iii) The paper's deferred validation is of the set. On the method the paper claims "a long way to validating" (p. 1088). |
| G&B "that which we are already familiar with"; "We argue elicitation carries the same risk: about 72 percent ..." | CONFIRMED | p. 115; p. 1090 |
| "Olsen credits tools that bring "new design participants" into design." | CONFIRMED | Olsen p. 5 |
| "That is this method's claim: non-technical people author the set." | CONFIRMED in substance; **style**: unlabelled analogy | The watch-out itself says applying Olsen to a method "is our analogy". |
| "His last test, that the result is "better," is still open." | CONFIRMED (Olsen: "The last claim is that the resulting UI designs are "better""); **style**: "still open" is our judgment, unlabelled | Olsen p. 5 |
| "Weiser's walk-up boards ran at 1024 by 768, like this Surface, and anyone could use one without signing up." | CONFIRMED | Guelph PDF pp. 4 to 5; p. 1085 |
| "The paper's "immediate usability" in public spaces is that same requirement." | **WRONG (overstated, unlabelled)** | See row 8 |
| Roots sentence: "Furnas explains the aliases, Hutchins, Hollan and Norman explain the missing recognizer, and Morris's group gestures make a single-user study a question." | CONFIRMED | pp. 1088, 1085, 1084 |

**Transition**: CONFIRMED. The authors list validation and a recognizer as next steps (p. 1091).

**Watch out**

| Item | Verdict | Evidence |
|---|---|---|
| Interactions article (pp. 38 to 44) does not name the paper | CONFIRMED | Full-text search for "gesture", "surface", "1083" found no hit. Printed pages 38 to 44 checked on the footers. |
| Whitepaper title "Seven Research Contributions in HCI", p. 3, next to Wobbrock et al. 2005 (ref. 33) | CONFIRMED | whitepaper p. 1 title; p. 3; paper ref. 33 p. 1092 |
| 2016 article summarizes the whitepaper (p. 40, endnote 3, p. 44) | CONFIRMED | |
| Wobbrock first author of both | CONFIRMED | p. 1083; Interactions p. 39 |
| Paper's four contributions, method not among them; abstract "We present an approach ..." | CONFIRMED (2 items) | p. 1084; p. 1083 |
| "Wobbrock and Kientz say papers often make more than one type" | CONFIRMED in substance | Interactions p. 44: of 2,316 CHI 2016 submissions, "705 indicated more than one (30.4 percent)". Whitepaper p. 1: "some research articles make more than one type of contribution." This is not a direct statement that papers "often" do so. A precise wording is in D6. |
| G&B (a): missing validation of "an early, candidate set" not automatically a flaw | **WRONG (pairing)** | See row 5 |
| G&B (b): convergence on the familiar; "I'm a child of the mouse" | CONFIRMED | p. 115; p. 1090 |
| G&B "can mute creative ideas that do not conform to current interface norms" (p. 111) | CONFIRMED | |
| G&B "participatory critique" (p. 118) | CONFIRMED | |
| Buxton's affiliation Microsoft Research (p. 111) | CONFIRMED | "Bill Buxton, Principle Researcher, Microsoft Research, Redmond" |
| Olsen's criteria are for systems and tools; applying them to a method is an analogy | CONFIRMED | The Olsen abstract, p. 1, is about "new UI systems work" |
| Olsen's "three steps ... show the new tool works for it, the last being "better" designs" | **WRONG (conflation)** | Olsen p. 5 has three parts: "describe some population ... show why they should be involved"; "establish why existing tools are not acceptable"; "Lastly one must demonstrate that the new tools are accessible, easier or more effective for this desired population." After these come a separate graded list of claims: "The simplest criterion ... could not be used", "A second claim ... easier", "The last claim is that the resulting UI designs are "better"". So "better" is the last claim in that second list, not the third step. |
| Step 1 p. 1084; step 2 p. 1083 and 43.5% p. 1090; self-ratings "goes a long way to validating" p. 1088 | CONFIRMED (3 items) | |
| Olsen p. 1 touch-table quote; "single point" p. 1090 | CONFIRMED (2 items) | |
| Weiser boards about 40 by 60 inches, vertical; Surface 24 by 18 (p. 1085) | CONFIRMED (2 items) | "not everyone can reach the top of the board" (PDF p. 5) implies vertical. |
| Surface size between pad and board scales | CONFIRMED | Weiser: "inch-scale ... foot-scale ... yard-scale" (Stanford HTML). 24 by 18 inches is still nearer foot-scale, but "between" is defensible. |
| "Weiser does not discuss tabletop gestures" | CONFIRMED | Weiser's gestures are pen gestures at a liveboard ("gestures with it towards her liveboard") and VR. There is no tabletop. |
| Chalk "either in contact with the surface or from a distance" | CONFIRMED | Guelph PDF p. 4 |
| "the paper recommends sensing above the table (p. 1091)" | CONFIRMED (wording soft) | The paper says "may benefit from the ability to look down or sideways at users' hands". |
| Menu bar quote; imaginary widgets at top-right and screen bottom (p. 1090) | CONFIRMED (2 items) | |
| Weiser copy has no journal pages | CONFIRMED | Guelph PDF is a Word reformat |
| Olsen copy is "the March 2007 pre-print (7 pages)" | 7 pages CONFIRMED; "pre-print" **UNVERIFIABLE** | The PDF carries no pre-print label. March 2007 comes only from the PDF metadata (CreationDate 2007-03-27). Its status relative to the UIST camera-ready cannot be checked from the copy. |
| UIST proceedings pp. 251 to 258 | CONFIRMED | Interactions endnote 5, p. 44; G&B ref. 23, p. 120 |

**Sources line**: all six URLs open and contain the stated document. CONFIRMED (6).

**Open items**: the Weiser journal page range, marked `[UNVERIFIED]`, can now be verified. Crossref DOI 10.1038/scientificamerican0991-94 gives Scientific American 265(3), pp. 94 to 104. CONFIRMED.

## D. Corrections (exact replacements, priority order)

### D1 (high). Method validation vs set validation (row 1, slide bullet 1, Say)

Slide, line 10:
- Old: `- **Wobbrock and Kientz 2016:** methods need "repeated validation" (p. 41). Here validation is a "next step" (p. 1091).`
- New: `- **Wobbrock and Kientz 2016:** methods need "repeated validation" (p. 41). Here one study (p. 1088); set validation a "next step" (p. 1091).`

Say, line 17:
- Old: `And methods, they write, are judged on validity, which "requires repeated validation." This paper leaves validation as a next step.`
- New: `And methods, Wobbrock and Kientz write, need "repeated validation." This paper has one study, whose own ratings the paper says go "a long way to validating" the approach.`

Claims table, row 1, Claim cell:
- Old: `Methodological contributions are judged on validity, which needs repeated validation; this paper defers validation`
- New: `Methodological contributions need repeated validation of utility, reproducibility, reliability and validity; this paper validates the method with one study's own goodness ratings (p. 1088) and defers validating the set (p. 1091)`

Claims table, row 1, Paper page cell:
- Old: `p. 1091 ("An important next step is to validate our user-defined gesture set.")`
- New: `p. 1088 ("This finding goes a long way to validating this user-driven approach to gesture design."); p. 1091 ("An important next step is to validate our user-defined gesture set.")`

### D2 (high). "Early candidate" is not supported by "candidate for deployment" (row 5, Watch out G&B)

Watch out, line 22:
- Old: `(a) under their argument, the missing validation of an early, candidate set (the paper calls it "a good candidate for deployment," p. 1091) is not automatically a flaw;`
- New: `(a) under their argument, skipping usability evaluation of an early-stage design is not automatically a flaw, but the paper presents the set as "a good candidate for deployment" (p. 1091), which is not early-stage, so we argue their excuse only partly applies;`

Claims table, row 5, Claim cell:
- Old: `Missing validation of an early candidate set is not automatically a flaw under Greenberg and Buxton`
- New: `Greenberg and Buxton excuse skipped evaluation of early-stage designs; the paper presents its set as ready for deployment, so we argue the excuse only partly applies`

### D3 (medium). Weiser "same requirement" overstated (row 8, slide bullet 4, Say)

Slide, line 13:
- Old: `- **Weiser 1991:** walk-up boards at 1024x768; same Surface resolution, same need for "immediate usability" (p. 1085).`
- New: `- **Weiser 1991:** boards for "casual use," 1024x768, like this Surface, meant for public "immediate usability" (p. 1085).`

(With D1 and D3 applied, on-slide text is about 91 words. To get back under 90, also replace `Here non-technical users author the set (p. 1084).` with `Non-technical users author the set (p. 1084).`)

Say, line 17:
- Old: `Weiser's walk-up boards ran at 1024 by 768, like this Surface, and anyone could use one without signing up. The paper's "immediate usability" in public spaces is that same requirement.`
- New: `Weiser's boards ran at 1024 by 768, like this Surface, built for "casual use" with no sign-up. We argue the paper's "immediate usability" in public spaces goes further: no learning.`

Claims table, row 8, Claim cell:
- Old: `Weiser's walk-up boards share the Surface's 1024x768 and its immediate-usability requirement`
- New: `Weiser's boards share the Surface's 1024x768 and were built for walk-up "casual use"; linking that to the paper's "immediate usability" is our analogy`

Claims table, row 8, Reading cell: before `(copy PDF pp. 4 to 5; no journal page)`, insert `"We have built enough liveboards to permit casual use"`.

### D4 (medium). Olsen's three steps conflated with his list of claims (Watch out, line 23)

- Old: `Olsen's three steps for that claim: describe the population and why it should be involved; show why existing tools do not serve it; show the new tool works for it, the last being "better" designs, where "The question of defining 'better' is problematic" (pre-print p. 5).`
- New: `Olsen's three steps for that claim: describe the population and why it should be involved; establish why existing tools are not acceptable for it; demonstrate the new tools are "accessible, easier or more effective" for it. He then grades the claims, the last being that the resulting designs are "better," where "The question of defining 'better' is problematic" (pre-print p. 5).`
- Same line, old: `and only self-ratings for step 3 ("goes a long way to validating," p. 1088).`
- New: `and only self-ratings for step 3 and the "better" claim ("goes a long way to validating," p. 1088).`

### D5 (low, style). Unlabelled judgments in Say (line 17)

- Old: `That is this method's claim: non-technical people author the set. His last test, that the result is "better," is still open.`
- New: `We argue that is this method's claim: non-technical people author the set. His last test, "better" designs, we think is still open.`

(With D1, D3 and D5 applied, Say is 195 words, exactly the 130 wpm × 1.5 budget. Update `(about 185 words)` to `(about 195 words)`.)

### D6 (low). "Often make more than one type" (Watch out, line 21)

- Old: `the honest answer is that Wobbrock and Kientz say papers often make more than one type and our mapping is a judgment.`
- New: `the honest answer is that Wobbrock's whitepaper says "some research articles make more than one type of contribution" (p. 1), 30.4 percent of CHI 2016 submissions ticked more than one (Wobbrock and Kientz, p. 44), and our mapping is a judgment.`

### D7 (low). "Recommends" sensing (Watch out, line 24)

- Old: `and the paper recommends sensing above the table (p. 1091);`
- New: `and the paper says tables "may benefit from the ability to look down or sideways at users' hands" (p. 1091);`

### D8 (low). Weiser journal pages now verifiable (Open items, line 69)

- Old: `- Weiser page numbers: the original is Scientific American 265(3), September 1991, pp. 94 to 104 per common citation, but no copy I opened shows printed pages, so no page is put on the slide. The journal page range itself is `[UNVERIFIED]` and is used nowhere on the slide.`
- New: `- Weiser page numbers: Scientific American 265(3), September 1991, pp. 94 to 104 (Crossref, DOI 10.1038/scientificamerican0991-94). No copy opened shows printed pages, so quotes carry no page.`

### D9 (low). Olsen "pre-print" (Watch out, line 25)

- Old: `The Olsen copy is the March 2007 pre-print (7 pages);`
- New: `The Olsen copy is a 7-page author copy dated March 2007 (PDF metadata), not the proceedings layout;`

## E. Optional strengthening (not corrections)

- G&B (b) has a closer quote for the "2008 habits" point than p. 115: "If done without regard to how cultures adopt technology over time, then today's reluctant reactions by users will forestall tomorrow's eager acceptance." (p. 111). It is about users' current reactions, which is nearer to elicitation than the measurement-bias sentence.
- Wobbrock and Kientz p. 40 cite Greenberg and Buxton ([4]) for artifacts that "sometimes should not be" accompanied by empirical studies. The file already logs this. It is a neat bridge between two readings if a classmate asks.
