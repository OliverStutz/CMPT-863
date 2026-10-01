# 04 The method

**Purpose:** Show how the guessability method works, where it comes from, why each design choice was made, and what it can and cannot deliver. **Slides:** 2. **Time:** 3:30. **Speaker:** Oliver. **Previous section:** 03 Pressure and ownership. **Next section:** 05 How the research was conducted.

## Slide 04.1: The method: a guessability study, from referent to sign

**Time:** 1:40 · **Visual:** left, a lineage strip with four dated stops (1984 Good, 2004 Nielsen, 2005 Wobbrock EdgeWrite, 2009 this paper). Right, a three-step loop drawn as a one-way arrow: "animation of the effect (referent)" → "participant performs a gesture (sign), thinking aloud" → "logged, no response from the system". Optional thumbnail of Figure 1 (PDF p. 1, printed p. 1083), the pan prompt.

**On the slide**
- Lineage: Good et al. 1984, "user-derived" email commands (p. 1084); Nielsen et al. 2004, "a similar approach" (p. 1084); Wobbrock et al. 2005, guessability and agreement score, EdgeWrite (p. 1087)
- Referent = a command's effect, shown as an animation; sign = the gesture performed (McNeill, p. 1085)
- Three choices:
  1. Ask users: "user behavior is rarely so systematic" (p. 1083)
  2. Effect first, plain 2D shapes, nothing from Windows or Mac (p. 1085)
  3. No recognizer: "a monologue in which the user's behavior is always acceptable" (p. 1085)
- Microsoft filed a patent application on this procedure in August 2008

**Speaker notes**
- **Say:** "Section three asked who owned this work. Now, what exactly did they do. We think the method is the paper's lasting contribution. The paper does not claim to have invented it. Good and colleagues in 1984 built a user-derived email command language. Nielsen and colleagues in 2004 describe a similar approach. And Wobbrock's own 2005 EdgeWrite paper introduced the guessability study and the agreement score; the CHI paper says the process was adopted from it. The words come from McNeill: the referent is the effect, the sign is the gesture that causes it. Three design choices. One: ask users, not designers, because, in the paper's words, user behavior is rarely so systematic. Two: show the effect first, with plain shapes and nothing from Windows or the Mac. The meaning is fixed, only the form is open. That answers the bias Epps and colleagues conceded, that their Windows screenshots may have made people emulate the mouse. Three: no recognizer. Nothing fails, so nothing pushes the participant to change. The paper calls this a monologue in which the user's behavior is always acceptable. One link back to section three: Microsoft filed a patent application on exactly this procedure in August 2008." (about 195 words)
- **Transition:** "So the participants produced 1,080 gestures. How do you turn 1,080 gestures into one set?"
- **Watch out:**
  - "Unrevised behavior" (p. 1085) means not revised by recognizer feedback during the session. It does not mean free of prior habits. The paper itself reports the desktop bias as a finding: about 72% of gestures were mouse-like one-point touches or paths, and "I'm a child of the mouse" (p. 1090). The neutral blocks world removed Windows prompts; it did not remove the habits people brought in. Sections 06 and 12 develop this.
  - Epps et al.: the paper says their prompts "may have biased participants to simply emulate the mouse" (p. 1084). Do not say "had pushed."
  - Nielsen et al.: the paper says only that they "describe a similar approach" (p. 1084) and does not say how its own method differs. If challenged on "first," the paper's claim is "the first to employ users, rather than principles, in the development of a gesture set" (p. 1084). Section 02 covers this.
  - "User-derived" is from the title of Good et al.'s paper, "Building a user-derived interface" (ref. 9, p. 1092). The body text says the approach "was used by Good et al. [9] to develop a command-line email interface" (p. 1084).
  - The cost of no feedback, which the paper itself names: participants could not revise earlier gestures (p. 1091); one used more fingers "to ensure that I was pressing," which the authors link to "the lack of feedback from the table" (p. 1090); some gestured above the table although "We instructed participants to touch the table while gesturing" (p. 1090). We argue the no-recognizer choice buys clean data on first instincts at the price of never seeing how people adapt.
  - The 2005 EdgeWrite study measured guessability directly: its abstract reports that users raised the guessability of the EdgeWrite alphabet from 51.0% to 80.1%. The 2009 study borrowed the elicitation and the agreement score but not that measurement step (see slide 04.2).
  - The patent application is US 2010/0031202, inventors Morris, Wobbrock, Wilson, filed 2008-08-04, published 2010-02-04, abandoned 2014-12-09. The paper does not mention it. Do not attribute a motive; section 03 handles interpretation.
- **Sources:** pp. 1083, 1084, 1085, 1087, 1090, 1091, 1092. Patent: https://patents.google.com/patent/US20100031202A1 (verified in progress.md). EdgeWrite 2005 abstract: https://www.springerprofessional.de/doi/10.1145/1056808.1057043 (DOI 10.1145/1056808.1057043, abstract text via search index).

## Slide 04.2: From 1,080 gestures to a set, step by step

**Time:** 1:50 · **Visual:** a five-step funnel across the top (1,080 → group → score → assign → alias → 48 gestures). Below it, Eq. 1 and Eq. 3 reproduced from p. 1087 (PDF p. 5), with a stacked bar of 20 people split 11 / 3 / 3 / 3 next to Eq. 3.

**On the slide**
- 1,080 gestures, 1,074 kept (p. 1086)
- **Group** identical gestures per referent; "identical" undefined (p. 1087)
- **Score** agreement score A = mean over referents of Σ (group share)² (Eq. 1)
  - Select single, 1 hand: 11, 3, 3, 3 → 0.37 (Eq. 3, p. 1087)
  - Range 1/20 (everyone different) to 1 (everyone the same). Not a percentage
- **Assign** largest group; conflicts go to the larger group (p. 1087)
- **Alias** 48 gestures for 22 referents, covering 57.0% of all gestures proposed (pp. 1087, 1088)
- **Around it** taxonomy (p. 1086); authors' own sets (p. 1090)
- **Limits** no shared gesture for close or task switch, people tapped imaginary buttons (p. 1087); guessability of the set never tested (p. 1091)

**Speaker notes**
- **Say:** "Twenty people, 27 referents, one and two hands: 1,080 gestures, six discarded. Step one: for each referent the authors sorted the gestures into groups of identical gestures. That is all the paper says. Step two, the agreement score. Take each group's share of the proposals, square it, add up, then average over referents. The paper's own example: one-hand select single. Eleven people made the same gesture, a tap according to Figure 4, and three groups of three did something else. That gives 0.37. Squaring rewards one dominant group. Two warnings. It is not a percentage: 0.37 does not mean 37 percent agreed. And it never reaches zero: twenty different gestures still score 0.05. Step three: the largest group wins the referent, and if two referents claim one gesture, the referent with the larger group wins. Step four: aliases. A referent may keep several gestures, because aliasing raises guessability. The result: 48 gestures for 22 referents, covering 57 percent of all proposals. Around this sit a taxonomy and a built-in test against the authors' own designs, which section five covers. What the method cannot do: where no group forms, as for close and task switch, it assigns nothing. And it never measures guessability itself. Still, a 2020 review counted 216 elicitation studies, and in 2015 Vatavu and Wobbrock corrected the score." (about 205 words)
- **Transition:** "That is the method on paper. Section five shows how it was run, and how the three authors' own gestures compared."
- **Watch out:**
  - **Eq. 1 in words** (p. 1087): for each referent r, Pr is the set of proposals and each Pi a group of identical proposals; sum (|Pi| / |Pr|)² over the groups, then sum over all referents and divide by |R|, the number of referents. Eq. 2 and Eq. 3 are single-referent cases.
  - **Both worked examples are in the paper.** Eq. 2: move a little, **two hands**, groups 12, 3, 3, 2 → 0.42. Eq. 3: select single, **one hand**, groups 11, 3, 3, 3 → 0.37. Arithmetic: (144 + 9 + 9 + 4) / 400 = 0.415; (121 + 27) / 400 = 0.37. Do not call 0.42 the one-hand score: one-hand move a little is 1.0 in Figure 3 (p. 1088).
  - **The range** is printed as [|Pr|⁻¹, 1] (p. 1087). With 20 proposals the floor is 1/20 = 0.05: every proposal "agrees with itself." Overall A is 0.32 for one hand and 0.28 for two hands (p. 1087). If asked "so 32% agreed?": no, A is a sum of squared shares. Vatavu and Wobbrock's corrected agreement rate runs from 0 to 1 and puts this paper's one-hand 0.32 at 0.284 (CHI 2015, PDF pp. 3 and 7).
  - **Grouping:** the paper says only "we grouped the gestures within each referent such that each group held identical gestures" (p. 1087). It does not say how finger count was treated. Our inference: the taxonomy rule that touching "with more than one finger at the same point" still counts as one-point (p. 1086) and the Figure 4 caption ("could be performed with 1-3 fingers," p. 1089) suggest finger count was merged, which would raise A. Label this "we think." Grouping was done by the authors and no inter-rater reliability is reported.
  - **Conflicts:** the paper does not say what happened to the referent that lost a conflict. Do not claim it was "rescued." The final set reuses gestures by target: the same four enlarge gestures mean zoom in on the background and open on a container (p. 1088).
  - **Aliases:** the paper does not state the threshold for keeping a second, third or fourth group. Breakdown (p. 1088): ten referents with 1 gesture, four with 2, three with 3, four with 4, one with 5, so 10 + 8 + 9 + 16 + 5 = 48. Of these 31 (64.6%) one-hand, 17 (35.4%) two-hand. Reason given: "Aliasing has been shown to dramatically increase input guessability [8,33]" (p. 1088); ref. 8 is Furnas et al., the vocabulary problem.
  - **22 referents:** 22 of 27 got dedicated gestures, with the two move referents combined. Insert and maximize reused dragging and enlarging; task switch and close got nothing because people tapped imaginary widgets and "a common gesture was not feasible" (p. 1087).
  - **Coverage 57.0%** is "of all gestures proposed" (p. 1087). The paper does not state the denominator (1,074 is likely). It counts proposals matched by any gesture in the set, aliases included, so it is not the same thing as agreement or guessability. The other 43% is not accommodated.
  - **Popularity and goodness:** gestures in larger groups got higher self-rated goodness, which the authors say "goes a long way to validating this user-driven approach" (p. 1088). The ratings were made by the gesture's own author; the chi-square uses N = 1074 gestures from 20 people as independent observations. Sections 07 and 09 cover this.
  - **Validation:** "An important next step is to validate our user-defined gesture set" (p. 1091): show unlabeled clips to 20 new people and see if they guess the commands, then build a recognizer. Neither was done in this paper. The GI 2010 follow-up tested preference, not guessability: 22 participants blind to authorship judged 81 gestures and preferred those authored by larger groups, partly because the researchers' gestures were more physically and conceptually complex. Do not shorten this to "users preferred user-authored gestures."
  - **Designer baseline** belongs to section 05 (43.5%, 60.9%, 19.1%, p. 1090). Do not give the numbers here.
  - **Taxonomy:** four dimensions, form, nature, binding and flow, "manually classified" by the authors (p. 1086, Table 2). It describes gestures; it does not choose them.
  - **Influence:** the DIS 2020 review counts N = 216 gesture elicitation studies, calls this paper "the first application of the guessability method to hand-gesture input," and lists it first among the most cited, 1,070 Google Scholar citations as of 15 August 2019 (Table 3, p. 859).
- **Sources:** pp. 1086, 1087, 1088, 1089, 1090, 1091. Vatavu and Wobbrock, "Formalizing Agreement Analysis for Elicitation Studies," CHI 2015: https://faculty.washington.edu/wobbrock/pubs/chi-15.01.pdf. Villarreal-Narvaez, Vanderdonckt, Vatavu, Wobbrock, DIS 2020: https://faculty.washington.edu/wobbrock/pubs/dis-20.01.pdf. GI 2010: verified in progress.md (graphicsinterface.org).

## Verification log

| Claim | Where checked | Status |
|---|---|---|
| Good et al. used the referent-sign approach for a command-line email interface | p. 1084 | OK |
| "user-derived" (title of ref. 9, "Building a user-derived interface," CACM 1984) | p. 1092 | OK |
| Nielsen et al. [19] "describe a similar approach"; ref. 19 dated 2004 | pp. 1084, 1092 | OK |
| Wobbrock et al. 2005 (ref. 33) used the approach for EdgeWrite unistrokes | pp. 1084, 1092 | OK |
| Grouping and agreement "adopted from prior work [33]" | p. 1087 | OK |
| 2005 paper introduced the agreement rate formula (sum of squared shares) | Vatavu and Wobbrock CHI 2015, PDF p. 2, citing [37] p. 1871 | OK |
| GES introduced 2005 as a "guessability method"; this paper the first application to hand gestures | DIS 2020, p. 855 | OK |
| EdgeWrite guessability raised from 51.0% to 80.1% | ACM abstract of DOI 10.1145/1056808.1057043, via search index; full text not opened | OK (abstract only) |
| Referent and sign from McNeill (footnote 1) | p. 1085 | OK |
| "user behavior is rarely so systematic" | p. 1083 | OK |
| Blocks world of 2D shapes; "To avoid bias [5], no elements specific to Windows or the Macintosh were shown" | p. 1085 | OK |
| Epps et al. "may have biased participants to simply emulate the mouse" | p. 1084 | OK |
| "remove the gulf of execution [10]"; "a monologue in which the user's behavior is always acceptable"; "unrevised behavior" | p. 1085 | OK |
| About 72% mouse-like one-point gestures; "I'm a child of the mouse" | p. 1090 | OK |
| "We instructed participants to touch the table while gesturing" | p. 1090 | OK |
| "to ensure that I was pressing"; "lack of feedback from the table" | p. 1090 | OK |
| Users could not change previous gestures | p. 1091 | OK |
| Patent application US 2010/0031202, filed 2008-08-04, abandoned 2014-12-09 | progress.md, patents.google.com/patent/US20100031202A1 | OK (verified earlier, reused) |
| 1,080 = 20 × 27 × 2; 6 discarded | p. 1086 | OK |
| "each group held identical gestures"; identical not defined | p. 1087 | OK |
| Eq. 1 structure and symbols | p. 1087, rendered PDF p. 5 | OK |
| Eq. 2: move a little, 2-hand, 12, 3, 3, 2 → 0.42 (arithmetic 0.415) | p. 1087 | OK |
| Eq. 3: select single, 1-hand, 11, 3, 3, 3 → 0.37 | p. 1087 | OK |
| Range [\|Pr\|⁻¹, 1] | p. 1087 | OK |
| A1H = 0.32, A2H = 0.28 | p. 1087 | OK |
| One-hand move a little = 1.0 | Figure 3, p. 1088 | OK (chart reading) |
| Corrected AR in [0, 1]; A min 1/\|P\|; 20 distinct proposals give A = .05; this paper's one-hand .320 → AR .284 | CHI 2015 PDF pp. 3, 7 | OK |
| "the referent with the largest group won the gesture" | p. 1087 | OK |
| Covers 57.0% of all gestures proposed | p. 1087 | OK |
| 22 of 27 dedicated, move referents combined; insert, maximize, task switch, close unassigned; "a common gesture was not feasible" | p. 1087 | OK |
| Aliasing quote and 10/4/3/4/1 breakdown; 48 gestures; 31 (64.6%) / 17 (35.4%) | p. 1088 | OK |
| Ref. 8 is Furnas et al., the vocabulary problem | p. 1092 | OK |
| Enlarge gestures reused for zoom in and open | p. 1088 | OK |
| Taxonomy four dimensions, manually classified by authors; one-point rule for several fingers | p. 1086 | OK |
| Figure 4 caption "1-3 fingers" | p. 1089 | OK |
| "goes a long way to validating"; chi-square N = 1074 | p. 1088 | OK |
| "An important next step is to validate"; reverse study with 20 new participants; recognizer | p. 1091 | OK |
| GI 2010: 22 blind participants, 81 gestures, preferred larger-group gestures | progress.md (graphicsinterface.org) | OK (verified earlier, reused) |
| DIS 2020: N = 216 studies; this paper most cited, 1,070 citations as of 2019-08-15 | DIS 2020 pp. 855, 859 (Table 3) | OK |

## Open items
- No `[UNVERIFIED]` items. The EdgeWrite 51.0% to 80.1% figure was checked against the published abstract only, and it appears in the notes, not on the slide.
- No conflict with the pitfalls list. One addition: Eq. 1 in the paper normalizes by |Pr| per referent, so with 6 discarded gestures some referents had 19 proposals, not 20. "Floor about 0.05" stays correct.
- Cross-references: section 03 (patent interpretation), section 05 (designer baseline numbers, so they are not repeated here), sections 07 and 09 (self-rated goodness and chi-square independence), section 11 (GI 2010, CHI 2015, DIS 2020 in more detail). The orchestrator may want to keep the Vatavu and Wobbrock and DIS 2020 mentions in only one of 04 and 11.
- Speaker-note word counts: 04.1 about 195 words for 1:40 (limit 216); 04.2 about 205 words for 1:50 (limit 238). On-slide text about 95 words per slide counting page references, about 75 without.
