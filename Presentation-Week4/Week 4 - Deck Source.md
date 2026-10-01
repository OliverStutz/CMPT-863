# Week 4 deck source: sections 01 to 12

Generated 2026-09-30 from the section files in `slides/` by `slides/_build-deck-source.py`. Edit the section files, not this file, then regenerate. Verification logs and open items live in the section files. Run sheet, placeholders and figure assets are in `slides/00-Overview.md`.


## Slide 0: Title

**Time:** 0:30 · **Speaker:** Oliver · **Visual:** Figure 1 (PDF page 1, printed p. 1083), the participant panning with a flat hand

**On the slide**
- **User-Defined Gestures for Surface Computing**
- Jacob O. Wobbrock (University of Washington, Information School), Meredith Ringel Morris and Andrew D. Wilson (Microsoft Research)
- CHI 2009, Boston, pp. 1083 to 1092
- CMPT 863 Human-Computer Interaction · Week 4 · [date]
- Presented by Oliver Stutz

**Speaker notes**
- **Say:** Good morning. This week's paper is Wobbrock, Morris and Wilson, "User-Defined Gestures for Surface Computing," CHI 2009. Wobbrock was at the University of Washington; Morris and Wilson at Microsoft Research, which will matter. About twenty-five minutes of talk in twelve short parts, then half an hour of discussion that you will run as much as I do. Section one puts us in 2008.
- **Transition:** Let's go back to 2008.


---

# 01 Background, 2008

**Purpose:** Put the room in 2008, show which touch and gesture devices an ordinary adult could already have used, then give the paper's motivation and the study setup as the base for everything after. **Slides:** 2. **Time:** 2:30. **Speaker:** Oliver. **Previous section:** Title slide. **Next section:** 02 The problem.

## Slide 01.1: 2008, where touch stood

**Time:** 1:15 · **Visual:** horizontal timeline 1985 to 2009, the study marked as a red band ending at August 2008; below it a single row of five small device icons with years. No figure from the paper.

**On the slide**
- 1985: Buxton's multi-touch tablet, Toronto
- 2001 DiamondTouch (MERL), 2002 SmartSkin (Sony)
- Feb 2006: Jeff Han's multi-touch demo at TED
- Jan 2007: iPhone announced; June 2007: on sale, with pinch
- May 2007: Microsoft Surface unveiled; April 2008: in AT&T stores
- **Before Aug 2008: this study** (patent filed 4 Aug 2008)
- April 2009: CHI, Boston; Oct 2009: Windows 7 with multi-touch
- Already in ordinary hands: laptop touchpads (1994), Palm Graffiti (1996), iPod touch wheel (2002), Nintendo DS (2004), Wii (2006)

**Speaker notes**
- **Say:** Let's go back to 2008. Multi-touch was not new in research. Bill Buxton's group in Toronto showed a multi-touch tablet in 1985, and by 2002 there were research tables like DiamondTouch and SmartSkin. What changed was visibility. Jeff Han's TED demo in February 2006 showed pinching photos to resize them. In 2007 the iPhone shipped with pinch-to-zoom, and Microsoft unveiled Surface, its table computer. Surface reached its first stores in April 2008. The paper gives no study date, but Microsoft filed a patent on the method on 4 August 2008, so the study was done by then. Now look at the bottom row. The paper screened out people who had used a tabletop or an iPhone. But an adult in 2008 could have used a touchpad, written Graffiti on a Palm, spun an iPod wheel, or tapped a Nintendo DS. We will come back to that in section 06.
- **Transition:** So that is the world; here is why the authors thought a study was needed, and what they did.
- **Watch out:** Do not say the participants were "touch-naive." The paper says "no participant had used an interactive tabletop, Apple iPhone, or similar" (p. 1085); it did not record touchpad, PDA or game console use. Do not say the study date is in the paper: it is inferred from the patent filing, which the paper does not mention. The Surface "April 2008" date is the first retail deployment (AT&T stores, from 17 April 2008), not consumer sales. FingerWorks (founded 1998, bought by Apple in early 2005) is the bridge from research pads to the iPhone if someone asks where Apple's multi-touch came from; section 03 covers it. Buxton's 1985 tablet was built in 1984 and shown publicly in 1985.
- **Sources:** p. 1085 (screening); CHI 2009 dates, p. 1083 footer. Buxton, "Multi-Touch Systems that I Have Known and Loved," https://www.billbuxton.com/multitouchOverview.html (1985 tablet, DiamondTouch 2001, SmartSkin 2002, Han 2005 to 2006, FingerWorks 1998 and 2005, iPhone pinch 2007, Surface 2007). Han TED: https://en.wikipedia.org/wiki/Jefferson_Han. iPhone: https://en.wikipedia.org/wiki/IPhone_(1st_generation). Surface: https://techcrunch.com/2008/04/01/microsoft-surface-coming-to-att-stores/. Patent: https://patents.google.com/patent/US20100031202A1. Windows 7: https://www.techtarget.com/enterprise-software/definition/Windows-7. Devices: see log.

## Slide 01.2: Why this study, and what it was

**Time:** 1:15 · **Visual:** left column "Why" with four short lines; right column "What" with Figure 1 (participant at the Surface, prompt "Pan, 1 hand," PDF page 1) as a small photo above four setup lines.

**On the slide**
- **Why (the paper says)**
- Tabletops for individuals, groups, multi-display rooms, fun (p. 1083)
- Keyboard and mouse "no longer preferable" (p. 1083)
- Gestures "defined by system designers" (p. 1083)
- Public tables: "the importance of immediate usability is high" (p. 1085)
- **What**
- 20 non-technical adults at a Surface prototype (p. 1085)
- 27 commands: see the effect, perform the cause; 1 and 2 hands; think aloud; no recognizer (p. 1085)
- 1,080 gestures (p. 1086)
- Out: taxonomy, agreement scores, 48 gestures for 22 commands (pp. 1086 to 1088)

**Speaker notes**
- **Say:** The paper gives four reasons. Researchers were building tabletops for single users, for groups, for multi-display rooms and for fun. On a table, keyboard and mouse are, in their words, "no longer preferable." Most surface gestures so far were invented by system designers, possibly chosen, the paper says, out of concern for reliable recognition. And tables may stand in public places, where people must succeed on first contact. So the study turned the usual order around. Twenty non-technical adults sat at a Surface prototype. For each of 27 commands they saw the effect, then made the gesture they thought would cause it, once with one hand and once with two, thinking aloud. Nothing was recognized. That gave 1,080 gestures, a taxonomy, an agreement score per command, and a set of 48 gestures covering 22 commands. We argue the timing matters: gesture vocabularies were being fixed in products right then, and two of the three authors worked at Microsoft Research.
- **Transition:** That leads straight to the problem the paper set out to solve: who gets to decide what a gesture means.
- **Watch out:** "No recognizer" means the system logged touches but did not respond to them (p. 1085); it does not mean participants were free of prior habits (the paper reports desktop habits as a finding, p. 1090). The 48 gestures cover 22 of 27 commands: ten with 1 gesture, four with 2, three with 3, four with 4, one with 5 (p. 1088); the two move commands share gestures, and insert, maximize, task switch and close got none of their own (p. 1087). Section 05 gives the detail of the run; do not go deeper here. The Microsoft point is context, not a claim about motive: say "we argue."
- **Sources:** pp. 1083, 1085, 1086, 1087, 1088; author affiliations, p. 1083.


---

# 02 The problem

**Purpose:** State the problem the paper addresses, why it mattered in 2008 and why it still matters, with analysis rather than summary. **Slides:** 1. **Time:** 1:15. **Speaker:** Oliver. **Previous section:** 01 Background, 2008. **Next section:** 03 Pressure and ownership.

## Slide 02.1: The problem: gestures built for the recognizer

**Time:** 1:15 · **Visual:** quote card at the top with the p. 1083 sentence; below it two columns, "Built from the recognizer outward" and "Users asked before"; tagline across the bottom. No figure from the paper.

**On the slide**
- The paper says: gestures "defined by system designers," "chosen out of concern for reliable recognition" (p. 1083)
- Yet the space is vast: "almost anything one can do with one's hands" could be a gesture (p. 1083)
- Built from the recognizer outward:
  - Wu's tabletops: "not elicited from users" (p. 1085)
  - Malik et al.: 1, 2, 3 or 5 fingers, "may be problematic" (p. 1085)
- Asked before: Good 1984, Epps 2006, Nielsen 2004 (p. 1084)
- Claim: "first to employ users, rather than principles" (p. 1084)
- **We argue: users had been asked before; nobody had built a whole surface gesture set from their agreement.**

**Speaker notes**
- **Say:** In the authors' words, most surface gestures were "defined by system designers," producing "somewhat arbitrary gesture sets whose members may be chosen out of concern for reliable recognition." The paper grants "skillful design." We argue the deeper problem is direction. Vocabularies were built from the recognizer outward. Wu's tabletop gestures were tested for usability, but never elicited from users. Malik's system told 1, 2, 3 and 5 fingers apart, which a camera does easily and, the paper suggests, people may not. And almost anything a hand can do is a candidate. Furnas showed in 1987 that two people pick the same word for the same thing less than 20 percent of the time. One designer's guess is a weak bet. Users had been asked before, by Good, Epps, Nielsen. What was missing was a full surface gesture set built from their agreement. And it mattered: public tables need immediate usability, and products were fixing conventions right then.
- **Transition:** If conventions were being fixed in products right then, the next question is who was fixing them, and whether Microsoft was under pressure.
- **Watch out:**
  - Do not say "nobody had asked users." The paper itself lists Good et al. 1984, Nielsen et al. 2004, Epps et al. 2006, Beringer 2002 and Voida et al. 2005 (p. 1084). The "first" claim is narrower: first to build a gesture set from users "rather than principles," where principles refers to Poggi and Wu et al. (refs. 20, 35, p. 1084).
  - Long et al.: the paper says "users are sometimes poor at picking easily differentiable gestures" (p. 1084), not designers. If someone raises it, it cuts both ways: it is the paper's own reason for resolving conflicts by agreement.
  - Malik: the paper hedges, "a feature our current findings suggest may be problematic for users" (p. 1085). Keep "may."
  - Furnas: the paper cites ref. 8 only to justify aliasing (p. 1088, p. 1091). The phrase "vocabulary problem" is the title of ref. 8, not a claim in the paper's text. Linking it to the problem is our reading. The 20 percent figure comes from the Furnas abstract, not from this paper.
  - A sharp classmate may notice that ref. 19, Nielsen et al., is cited both for recognition-driven gesture choice (p. 1083) and as "a similar approach" to elicitation (p. 1084). The paper does not say how its method differs from Nielsen's. Answer honestly: we cannot tell from this paper.
  - Wu et al. got usability studies. The critique is that the gestures were not elicited, not that nobody tested them.
  - "Products were fixing conventions" is our framing, supported in section 01 (timeline) and section 03, not a claim of the paper. The paper's own urgency is public spaces, "where the importance of immediate usability is high" (p. 1085), and "It is therefore timely" (p. 1083).
  - Why it still matters (if asked): every new input channel without an established vocabulary, such as mid-air hand tracking or voice, faces the same choice. Section 11 picks this up. **[classmate]** slot here if a commentary raised the recognizer-first point.
- **Sources:** p. 1083 (designer quote, "despite skillful design," design space, "timely"); p. 1084 ("first to employ users," Good, Nielsen, Epps, Beringer, Voida, Long et al.); p. 1085 (Wu "not elicited," Malik, public spaces); p. 1088 and p. 1091 (aliasing, refs. 8 and 33); p. 1092 (reference years). Furnas et al. 1987 abstract: https://scrapbox.io/UIPedia/George_Furnas:_The_vocabulary_problem_in_human-system_communication (cross-checked by search summary; ACM page https://dl.acm.org/doi/10.1145/32206.32212 returned 403).


---

# 03 Pressure and ownership

**Purpose:** Show who owned multi-touch research and gesture patents around 2005 to 2009, and test the hypothesis that Microsoft was under pressure, keeping documented facts apart from our interpretation. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 02 The problem. **Next section:** 04 The method.

## Slide 03.1: Who owned the gestures?

**Time:** 1:30 · **Visual:** a horizontal timeline, 2005 to 2009, Apple events above the line, Microsoft events below, a vertical marker "study done (before Aug 2008)". Right-hand box split in two: "Documented" and "We argue". No figure from the paper.

**On the slide**
- **Documented.** Apple: buys FingerWorks (2005); Jobs, Jan 2007, on multi-touch: "boy, have we patented it"; patent 7,479,949 granted Jan 2009; Cook, Jan 2009: "we will not stand for having our IP ripped off"
- Microsoft: hires Buxton (Dec 2005); Surface shown May 2007, in stores Apr 2008; Windows 7 touch demo May 2008
- Aug 2008: Microsoft files a patent application on this study's method (inventors Morris, Wobbrock, Wilson); abandoned 2014
- Apple and Microsoft: patent cross-licence since 1997
- **We argue:** a race to own gesture conventions; Apple's public lawsuit threat came only in 2009, after the study

**Speaker notes**
- **Say:** "Around this paper, gestures were becoming property. Apple bought FingerWorks, a multi-touch pad company, in 2005. Launching the iPhone in January 2007, Jobs said of multi-touch, 'and boy, have we patented it.' On Microsoft's side, Bill Buxton joined Microsoft Research in late 2005, Surface was shown in May 2007 and reached AT&T stores in April 2008, and Windows 7 multi-touch was demonstrated in May 2008. Then a fact the paper does not mention: in August 2008 Microsoft filed a patent application on this very method, with the three authors as inventors. It was abandoned in 2014. So was the study done to avoid lawsuits? We found no evidence of that. The paper gives academic reasons and builds on Wobbrock's 2005 guessability work. Apple and Microsoft had held a patent cross-licence since 1997, and Apple's public threats came in January 2009, after the study was finished. What we argue is narrower: Microsoft was racing to establish its own gesture conventions, and to own them, while Apple's were being fixed."
- **Transition:** "Whatever the business context, what Microsoft chose to patent was a procedure, and that procedure is what we think is the most valuable part of the paper. Let us open it up."
- **Watch out:**
  - "So was it defensive?" Answer: we cannot know intent, and we do not claim it. The documents show two things: the method was treated as intellectual property worth filing on (August 2008), and the strongest lawsuit threat in this story, Apple's, was made public only in January 2009, after the study. Between Apple and Microsoft a 1997 cross-licence already existed; in 2012 testimony Apple's Boris Teksler said the Apple-Microsoft cross-licence covers design patents and has an anti-cloning provision (we do not claim that clause dates from 1997). That lowers the lawsuit motive between those two firms in particular. The filing is consistent with offensive, defensive or routine corporate IP practice; we cannot tell which.
  - Do not say the paper hides the patent "on purpose." Say only that the paper does not mention it. Do not claim that a disclosure rule existed or did not exist.
  - Do not say the Jobs patent was filed in response to anything, and do not give its filing date (sources disagree). Do not call the abandoned application a patent: it was never granted.
  - Elan Microelectronics sued Apple over two multi-touch patents in April 2009. Use it only if asked: it shows multi-touch IP was contested, not that Microsoft feared suits.
  - Ownership of the research lineage: DiamondTouch was MERL (2001, ref. 3), SmartSkin Sony (Rekimoto 2002, ref. 21), Han at NYU (FTIR, UIST 2005), Wilson at Microsoft Research (TouchLight 2004, PlayAnywhere 2005, ref. 31). Two of three authors were Microsoft Research staff (p. 1083). Microsoft bought Han's Perceptive Pixel in July 2012, later, context only.
  - Hand-off: section 04 cross-references the patent filing; keep the date and inventor list identical.
- **Sources:** paper pp. 1083, 1084, 1092; URLs in the verification log.


---

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


---

# 05 How the research was conducted

**Purpose:** show how the elicitation study was actually run, then put on screen three results the run produced: the designers' poor prediction of users' gestures, one hand versus two, and finger use. **Slides:** 2. **Time:** 3:00. **Speaker:** Oliver. **Previous section:** 04 The method. **Next section:** 06 Sample and culture.

## Slide 05.1: The study as run

**Time:** 1:15 · **Visual:** left, a smaller Figure 1 photo (PDF page 1, p. 1083), or replace it with a simple diagram of one trial (animation plays, participant gestures, nothing responds) since Figure 1 already appears on the title slide and 01.2; right, Table 1 (PDF page 4, p. 1086), cropped to the two columns of referent names and ratings, at least half the slide's height, with the MEAN row (2.70) highlighted; if it cannot be read at that size, show only the five lowest and five highest rows plus MEAN. A small pipeline strip along the bottom: hear and see effect, reset, 1-hand gesture, 2-hand gesture, ratings.

**On the slide**
- 20 paid volunteers, 9 women, mean age 43.2 (sd 15.6), all right-handed, none had used a tabletop or iPhone (p. 1085)
- Microsoft Surface prototype, 24 x 18 inches, 1024 x 768 (p. 1085)
- 27 referents, random order; authors rated complexity 1 to 5, mean 2.70 (p. 1086)
- Per referent: 1-hand and 2-hand gesture, think aloud, preference, two 7-point ratings (pp. 1086, 1088)
- 1,080 gestures, 6 discarded; touch logs, video, notes (p. 1086)
- Baseline: each author's own gestures, made before the study (p. 1090)

**Speaker notes**
- **Say:** "Section 04 gave you the logic. Here is the run. Twenty paid volunteers from the general public, nine women, average age 43, all right-handed. Occupations included a restaurant host, a musician, an author, a steelworker. None had used a tabletop or an iPhone. They stood at a Surface prototype, 24 by 18 inches, at 1024 by 768. For each of 27 commands, in random order, a recorded voice named it, an animation showed the effect, and the scene reset. Then the person made a one-hand gesture and a two-hand gesture, thinking aloud, said which they preferred, and rated each on two seven-point scales: good match, and easy to perform. That gives 20 times 27 times 2, 1,080 gestures; six were dropped for confusion. The table logged every contact with millisecond timestamps, four cameras filmed the hands, two authors took notes. And before any of this, each author designed a gesture set of their own. Hold on to that."
- **Transition:** "That baseline is the first of three things the run revealed."
- **Watch out:**
  - The paper does not say whether the one-hand gesture always came before the two-hand one. It only says "participants performed a 1-hand and a 2-hand gesture" (p. 1086). Do not claim an order.
  - The complexity ratings in Table 1 are the three authors' ratings, made "before participants made gestures" (p. 1085), not participants' ratings.
  - The prototype watched hands "from beneath the table" with its vision system and logged contacts "as ovals having millisecond timestamps" (p. 1086). It did not recognize anything; there was no feedback beyond the animations.
  - Participants were told to touch the table: "We instructed participants to touch the table while gesturing" (p. 1090). Some still gestured above it.
  - Recruitment was "non-technical people without prior experience using touch screens (e.g., the Apple iPhone)" (p. 1084). Section 06 deals with what that screening could and could not exclude.
  - Participants "could not change previous gestures after moving on to subsequent ones" (p. 1091), an acknowledged limitation.
- **Sources:** pp. 1083 (Figure 1), 1084, 1085, 1086, 1088, 1090, 1091.

## Slide 05.2: Three things the run revealed

**Time:** 1:45 · **Visual:** three columns. Column 1: three big numbers 43.5% / 60.9% / 19.1% with the quote card "the authors are either thinking creatively or are hopelessly lost!" (p. 1090). Column 2: Figure 3 (p. 1088) with the 2-hand line highlighted at insert, maximize, enlarge. Column 3: a hand icon with 1 to 3 fingers marked "single point", 4 marked "gray area", 5 or palm marked "something more" (p. 1090).

**On the slide**
- **Designers missed.** Each author matched 43.5% of the user set; all three 60.9%; 19.1% of each author's gestures no participant tried (p. 1090)
- **One hand wins.** Preferred for 25 of 27 referents; tie for insert and maximize; two hands never preferred (p. 1088). Yet 17 of 48 set gestures (35.4%) are 2-hand (p. 1088)
- **Fingers barely matter.** 13 of 20 varied finger count; 2 said it mattered. 1 to 3 fingers: "single point"; 4: "gray area" (p. 1090)

**Speaker notes**
- **Say:** "One. Designers. Each author alone came up with only 43.5 percent of the users' set, all three together 60.9. And 19.1 percent of each author's gestures were never tried by anyone. The authors joke they were 'either thinking creatively or are hopelessly lost.' The fair reading: this measures overlap with the users' set, not quality. The yardstick is the users' set, so designers cannot win. The paper's own conclusion is modest: 'three experts cannot generate the scope of gestures that 20 participants can.'
  Two. Hands. One hand was preferred for 25 of 27 commands; insert and maximize were a tie; two hands never won. Yet a third of the final set is two-handed. The paper says some two-hand gestures 'had good agreement scores and nicely complemented their 1-hand counterparts.' We argue the design matters too: every participant had to make a two-hand gesture every time.
  Three. Fingers. Thirteen of 20 varied their finger count; only two said it mattered. One to three fingers read as a single point, five or a palm as something more, four as a gray area. Hence the guideline: do not distinguish gestures by finger count."
- **Transition:** "All of this came from twenty people of one kind, so next: who was not in the room."
- **Watch out:**
  - Figure 3 readings: the 2-hand line peaks at insert about 0.58, maximize about 0.58, enlarge about 0.52 (p. 1088). Say "about". The twist: insert and maximize, the two referents with the highest 2-hand agreement, are exactly the two with tied hand preference, and neither got a dedicated gesture "because they reused existing gestures" (p. 1088). So of these peaks only enlarge has 2-hand gestures in the set. Do not claim insert or maximize gestures appear in Figure 4.
  - The paper says the authors "hypothesized that the 'wisdom of crowds' would generate a better set than the authors" (p. 1090). "Better" here is judged by the users' own set, which is circular. If challenged: yes, and the paper concedes the narrower point quoted above.
  - Exact wording: "touches with 1-3 fingers were considered a 'single point,' and 5-finger touches or touches with the whole palm were something more. Four fingers, however, constituted a 'gray area'" (p. 1090). Paper spells "gray".
  - Guideline wording: "Gestures should not be distinguished by number of fingers" and "Four fingers should serve as a boundary between a few-finger single-point touch and a whole-hand touch" (p. 1091).
  - Contrast with prior systems: Malik et al. distinguished "1-, 2-, 3-, and 5-finger gestures, a feature our current findings suggest may be problematic for users" (p. 1085). The paper also lists ref. 17, Morris's own cooperative gestures work, among systems "differentiated only on the basis of the number of fingers used" (p. 1090). A nice point of self-correction if asked.
  - Counter-evidence a classmate may raise: two participants who used one-finger touches "felt that the system needed to distinguish among fingers," e.g. a ring-finger tap for a menu (p. 1090). One said he used more fingers "to ensure that I was pressing," which the paper ties to "the lack of feedback from the table" (p. 1090). So finger behaviour was partly shaped by the no-feedback setup.
  - The paper does not say how finger count was treated when grouping identical gestures (p. 1087); the Figure 4 caption lets one-finger gestures be done with 1 to 3 fingers (p. 1089). Present any link to agreement scores as inference.
  - Statistics behind hand preference are not reported per participant; "overall" preference is the paper's word (p. 1088).
- **Sources:** pp. 1085, 1087, 1088, 1089, 1090, 1091.
- **[classmate]** slot: a classmate's commentary on the designer comparison or the finger finding fits after point 1 or point 3.


---

# 06 Sample and culture

**Purpose:** show who was missing from the sample (younger and older people, other ethnicities and cultures), how culture is built into some gestures of the set, and why the 2008 adult arrived already shaped by the mouse and by single-point touch. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 05 How the research was conducted. **Next section:** 07 Findings.

## Slide 06.1: Who was in the room, and who was not

**Time:** 1:30 · **Visual:** left, a "sample card" with the participant line (p. 1085) and the limitation quote (p. 1091); right, two crops of Figure 4 (PDF page 7, p. 1089): row 3, the Accept "draw check", Reject "draw 'X'" and Help "draw '?'" panels; and the bottom row, right-most panel, Next "draw line across object", with three small culture icons underneath (head nod, ✓ versus ○, thumbs-up). **[classmate]** slot at the bottom.

**On the slide**
- 20 adults, mean age 43.2, sd 15.6 (p. 1085); no range, no results by age
- No ethnicity reported. The authors: "all non-technical literate American adults; undoubtedly, children, Eastern, or uneducated participants would behave differently" (p. 1091)
- Symbols in the set: check, X, ?, a line drawn left to right (Fig. 4, p. 1089)
- Meanings vary: in Japan ○ is correct and ✓ can mark wrong; in Bulgaria a nod means no; thumbs-up insults in Iran. We argue: the set's symbols may not travel
- No iPhone or tabletop users (p. 1085), yet "I’m a child of the mouse" (p. 1090)
- **[classmate]**

**Speaker notes**
- **Say:** Who sat down? Twenty adults, average age 43, standard deviation almost 16. That spread tells us at least one person was under about 28 or over about 58, but the paper gives no range and breaks no result down by age. It reports no ethnicity at all. The authors say it themselves: all were "non-technical literate American adults," and children, Eastern or uneducated people "would behave differently." Now look at the set. Accept is a check mark, reject an X, help a question mark, next a line drawn left to right. In Japan a circle means correct and a check can mark a wrong answer. In Bulgaria a nod means no. We argue the most abstract commands got the most culture-bound gestures. And these adults were not blank: screened for no iPhone or tabletop, yet "I'm a child of the mouse." The paper treats that as a finding about users. We argue it is also a limit of the method. **[classmate]**
- **Transition:** With that lens on who was asked, let us look at what they actually did.
- **Watch out:**
  - The paper does not hide this. The cultural limitation is in its own Limitations section (p. 1091). Our point is the consequence for the set, not that the authors missed it.
  - Age: "at least one person under about 28 or over about 58" follows from the standard deviation (if everyone were within 15 years of 43.2 the sd could not reach 15.6). Do not say "people in their twenties and sixties" as fact. Children are excluded by "adults" (p. 1091).
  - "American" is the authors' word for the sample. Do not equate it with one ethnicity; the paper simply reports none.
  - "Left to right" is how Figure 4 draws Next; the caption says reversible gestures are drawn "in only one direction" (p. 1089). That Next follows reading direction is our inference.
  - The head-nod example is a body gesture, not a surface gesture. Use it only to show that the meaning of a movement is learned. The check mark example is the one that touches the set directly.
  - Thumbs-up is not in the set. Keep it to one clause or drop it if time is short.
  - If asked "did culture actually change gestures?": Mauney, Howarth, Wirtanen and Capra, CHI 2010 Extended Abstracts, ran a gesture elicitation study for touchscreens across several countries. A secondary summary reports mostly similar gestures across countries, with Chinese participants using more symbolic gestures [UNVERIFIED: findings not read in the original]. Say "a later study across countries reportedly found more similarity than difference, except for symbolic gestures," and add that symbolic gestures are exactly where the cultural loading of this set sits.
  - Tainting: the screening named the iPhone, tabletops "or similar" (p. 1085) and recruited people "without prior experience using touch screens (e.g., the Apple iPhone)" (p. 1084). No touch history was recorded. A 2008 adult could already have used a laptop trackpad (from 1994), a Palm Pilot with Graffiti stylus strokes (from 1996), a touch-screen subway ticket machine (New York, 1999) or a Nintendo DS (from November 2004). Every one teaches one finger or one point. Say "could also come from," not "came from." "Unrevised behavior" (p. 1085) means not revised by recognizer feedback; it never promised habit-free participants (pitfall 1).
  - The desktop evidence is the paper's own, under "It's a Windows World" (p. 1090): tapping index and middle finger as on a two-button mouse, imaginary close buttons at objects' top right, an imaginary taskbar at the screen's bottom, "I'm a child of the mouse." About 72% of gestures were mouse-like (p. 1090); section 07 covers that number.
- **Sources:** p. 1084 (recruiting), p. 1085 (participants), p. 1089 (Figure 4 and caption), p. 1090 ("It's a Windows World"), p. 1091 (limitations quote). Japan circle and check: https://www.nippon.com/en/guide-to-japan/cs00032/ and https://en.wikipedia.org/wiki/Check_mark. Bulgaria: https://3seaseurope.com/bulgarians-shake-nod-heads/. Thumbs-up in Iran: https://en.wikipedia.org/wiki/Thumb_signal. Mauney et al. 2010: https://api.semanticscholar.org/graph/v1/paper/DOI:10.1145/1753846.1754095 (existence), https://prezi.com/of9ihhm6_sgd/upa-presentation (secondary summary of findings). Devices: https://en.wikipedia.org/wiki/PowerBook_500_series, https://en.wikipedia.org/wiki/Pilot_1000, https://longreads.com/2015/01/25/how-the-mta-introduced-the-touch-screen-to-new-york-city/, https://www.giantbomb.com/nintendo-ds/3045-52/.


---

# 07 Findings

**Purpose:** Rank the findings that matter most, then show what the timing data say about thinking versus doing. **Slides:** 2. **Time:** 2:30. **Speaker:** Oliver. **Previous section:** 06 Sample and culture. **Next section:** 08 Technical obstacles.

## Slide 07.1: What people did

**Time:** 1:10 · **Visual:** ranked list of eight findings, each with a one-word solidity tag in a coloured chip (Counted, Coded, Observed, Self-rated). Small inset of Figure 2, Form columns only (PDF page 5, p. 1087), with the one-point path and one-point touch segments highlighted.

**On the slide**
1. Desktop "deeply rooted": about 72% mouse-like one-point touches or paths (p. 1090). *Coded*
2. One hand preferred for 25 of 27 referents (p. 1088). *Counted*
3. Agreement falls with complexity, r = -0.52 (p. 1087). *Counted*
4. 43.9% of gestures physical, treating shapes as real objects; complex commands drew metaphors and symbols (p. 1087). *Coded*
5. Finger count rarely matters: 13 of 20 varied it, 2 cared (p. 1090). *Observed*
6. Opposite commands (enlarge, shrink) got mirror-image gestures, though never shown as pairs (p. 1090). *Observed*
7. Off-screen drags, mid-air gestures, imaginary widgets (p. 1090). *Observed*
8. Gestures many people chose got higher "good match" ratings from the people who made them (p. 1088). *Self-rated*

Footer, small type: Tags: Counted = tallied from logs or ratings; Coded = sorted by the authors; Observed = authors' notes and video; Self-rated = participants judged their own gestures

**Speaker notes**
- **Say:** "We ranked the findings by how much they change design. Top: the desktop came along. About 72 percent of gestures were one-point touches or paths, mouse-like, and the paper itself calls the desktop deeply rooted. Second, one hand won for 25 of 27 commands. Third, the harder the command, the less people agreed. Fourth, 44 percent of gestures were physical, treating shapes like objects, and the paper shows simple commands drew physical gestures while complex ones drew metaphors and symbols. Then three observations: finger count rarely mattered, opposite commands got mirror-image gestures although the pairs were never shown together, and people dragged off the screen, gestured in the air, and tapped buttons that did not exist. Last, popular gestures were rated better. The paper reads that as validation; we argue it is the makers rating their own gestures. The chips tell you how firm each finding is."
- **Transition:** "Those are counts. The timing data tell a stranger story."
- **Watch out:**
  - The 72% is the paper's rounded figure (p. 1090). Figure 2 reports form per hand: right hand .450 path plus .274 touch, left hand .379 plus .343, both about 72% (p. 1087). Do not say "72% of participants."
  - 43.9% physical is printed on p. 1091 and as .439 in Figure 2 (p. 1087). Mean complexity per nature category: physical 2.11, abstract 2.99, metaphorical 3.26, symbolic 3.52 (p. 1087).
  - r = -0.52 is over 27 referents (F1,25, p. 1087), with complexity rated by the three authors (Table 1, p. 1086). A classmate may say the authors defined "complex"; answer that the same ratings predict planning time (next slide), so they track something real.
  - The goodness link is a chi-square on N = 1074 gestures treated as independent (p. 1088); 20 people made all of them. "Goes a long way to validating" is the paper's phrase (p. 1088); the set itself was never validated (p. 1091).
  - Imaginary widgets: task switch and close (p. 1087); clicking with two fingers as if on a two-button mouse (p. 1090). Mid-air examples include the beckoning hand for help and "glue" before paste (p. 1090). The paper says participants were instructed to touch the table (p. 1090).
  - Finger numbers: 1 to 3 fingers read as a "single point," 4 a "gray area" (p. 1090). Section 05 covers this in depth; keep it to one line here.
- **Sources:** pp. 1086, 1087, 1088, 1090, 1091.

## Slide 07.2: Time tells a story: finding the gesture costs more than making it

**Time:** 1:20 · **Visual:** one horizontal time axis from 0 to 16 s with two bars. Top bar "thinking" (planning time) with marks at about 8 s (simple commands) and about 15 s (complex). Bottom bar "doing" (articulation time) spanning about 2 to 3.4 s. A small inset at the right: "rated easy 3.4 s / rated difficult 2.0 s." All numbers from p. 1088 (PDF page 6).

**On the slide**
- Thinking: about 8 s before touching for simple commands, about 15 s for complex ones; r = 0.71 with complexity (p. 1088)
- Doing: about 2 to 3.4 s, and complexity made no difference to it (p. 1088)
- Longer thinking went with lower "good match" ratings: good gestures came to mind fast (p. 1088)
- We argue: on a walk-up table the cost is in finding the gesture, not making it. That is what guessability is for
- Speed is not ease: gestures rated easy took about 3.4 s, rated difficult about 2.0 s (p. 1088). Comfortable gestures were slow and smooth, awkward ones quick and precise (we argue)
- Caveat: averages only, and the tests count 1,074 gestures as if 1,074 people had made them; 20 did

**Speaker notes**
- **Say:** "Three things about time, all from one page. First, finding the gesture is the expensive part. People paused about eight seconds before touching for simple commands and about fifteen for complex ones, while the gesture itself took two to three seconds whatever the command. We argue that is the number a walk-up table lives or dies by: the cost is in guessing, not in doing, which is exactly what this method targets. Second, a slow start was a bad sign. Longer planning went with lower fit ratings, and the paper concludes good gestures were the ones most quickly apparent. Third, a small surprise worth one sentence: speed is not ease. Gestures people rated easy took longer, about 3.4 seconds, than gestures rated difficult, about 2. The comfortable gestures were the slow, smooth ones, drags and pull-aparts; the awkward ones were quick and precise, corner taps and drawn marks. A designer who optimised for brevity would pick the gestures users dislike. The paper only says smoother or less hasty; the precision reading is ours."
- **Transition:** "Those fifteen seconds of thinking happened on a 2008 prototype, which brings us to what the technology of the time could and could not do."
- **Watch out:**
  - Keep the two difficulties apart. Command complexity (the authors' 1 to 5 rating of the animation's effect) drove thinking time and did not affect doing time. The "easy" and "difficult" in the 3.4 and 2.0 figures are the participants' ratings of their own gesture on the "easy to perform" scale. A hard animation did not produce a fast gesture; the paper reports that complexity "did not correlate significantly with gesture articulation time" (p. 1088).
  - Exact wording to quote if asked: "good gestures were those most quickly apparent to participants" and "perhaps because they were smoother or less hasty" (p. 1088). The second is the authors' hedge, not a tested result.
  - Planning time is "the more time participants took to begin articulating their gesture" (p. 1088). The paper does not define articulation time or say where simple and complex are cut on the 1 to 5 scale; 8 s and 15 s are rounded values with no spread reported.
  - The thinking and doing bars split the data differently (by complexity, by ease rating), and the paper gives no single overall mean for either. Say "several times longer," not a precise factor.
  - Planning time did not affect ease, and complexity did not correlate with ease (p. 1088). Only articulation time is linked to ease; only planning time and complexity are linked to goodness. Complexity against goodness r = -0.59, means 5.6 against 4.9 on the 7-point scale.
  - Touch-event counts: easy-rated gestures had about twice as many contact events (p. 1088). The paper itself notes "long lived gestures will have more touch events," so we argue this is the same duration measured a second way, not independent support. It does fit the reading that the easy gestures were long drags.
  - Statistics: r = 0.71 and r = -0.59 are over 27 referents (F1,25). The time-on-rating effects are chi-square tests on N = 1074 gestures as if independent, though 20 people produced them. Response times are usually right-skewed and the paper reports no distribution.
  - If a classmate asks whether the easy gestures were simply the continuous ones: plausible, but the paper does not break timing down by flow or form; label it "we argue."
- **Sources:** p. 1088 for every number on this slide; Table 1 (p. 1086) for the complexity scale.


---

# 08 Technical Obstacles

**Purpose:** Show what the 2008 table could and could not sense, and how those limits shaped the study and its recommendations. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 07 Findings. **Next section:** 09 Limitations.

## Slide 08.1: What the 2008 table could not see

**Time:** 1:30 · **Visual:** left, a scale bar comparing pixel density, about 43 ppi (2008 Surface prototype) against 264 ppi (current iPad Pro); right, a side view of a table with three zones labelled "sensed: contact from below," "not sensed: above the table," "not sensed: beyond the edge." Optional inset: Figure 1 photo, PDF page 1 (p. 1083).

**On the slide**
- Display: 24" x 18" at 1024 x 768 (p. 1085), about 43 pixels per inch. A current iPad Pro: 264 ppi.
- Sensing: a vision system "from beneath the table"; contacts logged as ovals with millisecond timestamps (p. 1086). Cameras recorded the session separately (p. 1086).
- No recognizer, no touch feedback (p. 1085, p. 1090).
- Participants were told to touch the table (p. 1090). Some still gestured in the air and beyond the screen edge (p. 1090).
- A physics engine (Wilson et al. 2008): supported by 43.9% physical gestures, but "will probably not suffice" alone (p. 1091).
- Set claims "ease of recognition," never tested (p. 1091).

**Speaker notes**
- **Say:** The table in this study was coarse by today's standards. Twenty-four by eighteen inches at ten twenty-four by seven sixty-eight is about forty-three pixels per inch. A current iPad Pro has two hundred sixty-four, roughly six times finer in each direction. The paper sensed touch with a vision system watching the hands from beneath the table and logged each contact as an oval with a timestamp. The four video cameras were a separate recording of the session. There was no recognizer and, as one participant showed by adding fingers "to ensure that I was pressing," no feedback that a touch had landed. The paper says participants were instructed to touch the table. Even so, some beckoned palm-up, lifted objects with two hands, or dragged things off the screen edge. The authors read this as a sensing gap: they suggest looking down or sideways at hands and adding a low-resolution sensing border. They also argue the many physical gestures support a physics engine, but not alone. We argue the "ease of recognition" claim is a hypothesis, since no recognizer was built.
- **Transition:** A table that could not see the air, and a set that was never run through a recognizer, lead straight into the study's limitations.
- **Watch out:**
  - The touch instruction is not in the Procedure section (pp. 1085 to 1086). It is one sentence under "Acting above the Table" on p. 1090: "We instructed participants to touch the table while gesturing." Quote it exactly; the paper does not say how or when the instruction was given.
  - The paper never names resolution as a problem. Linking 43 ppi to the advice that "demands for precise positioning should be avoided" (p. 1091) is our inference, and the paper's reason for that advice is user behaviour (12 of 14 resized sideways), not pixels.
  - Do not say the Surface used infrared cameras or projectors underneath: the paper says only "vision system" "from beneath the table" (p. 1086) and "the table's projected screen" (p. 1090). Extra hardware detail would need an outside source.
  - 43.9% physical matches the Figure 2 nature value .439 (pitfall 20).
  - If asked whether the iPad comparison is fair: it is a hand-held screen, not a table; the point is density only.
- **Sources:** pp. 1085, 1086, 1090, 1091, 1092 (ref. 32); https://www.apple.com/ipad-pro/specs/


---

# 09 Limitations

**Purpose:** Separate the limitations the authors state from the ones we add, and size up what a sample of 20 can carry. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 08 Technical obstacles. **Next section:** 10 Connections.

## Slide 09.1: Limitations, theirs and ours

**Time:** 1:30 · **Visual:** two columns, "The authors say (p. 1091)" left, "We add" right; a small footer card: "n = 20. Across 216 later elicitation studies the average was 25, and 20 the most common choice (Villarreal-Narvaez et al., DIS 2020, p. 859)"; **[classmate]** slot under the right column

**On the slide**

*The authors say (p. 1091)*
- No revising earlier gestures
- No application, organizational or cultural context
- "children, Eastern, or uneducated participants would behave differently"
- Set not validated; no recognizer

*We add*
- n = 20 (p. 1085): enough to show strong agreement, too few to compare groups such as age or culture
- Authors grouped and coded; no inter-rater reliability (pp. 1086 to 1087)
- The statistics count 1,074 gestures as if 1,074 different people had made them. Twenty did, so the p-values are too confident (p. 1088)
- Gestures judged only by the people who made them, yet "goes a long way to validating" (p. 1088)
- Five of the 27 commands are desktop ideas: cut, paste, minimize, maximize, task switch (p. 1085). Desktop questions invite desktop gestures
- The set matches 57% of what participants did; the other 43% of their gestures get nothing (p. 1087)
- "ease of recognition" without a recognizer (p. 1091)
- Two of three authors at Microsoft, a Microsoft prototype; the patent application goes unmentioned
- **[classmate]**

**Speaker notes**
- **Say:** The paper says a lot about its own limits, on page 1091. Participants could not go back and change a gesture. There was no application, organizational or cultural context. In their words, "children, Eastern, or uneducated participants would behave differently." And the set is not yet validated, with no recognizer built. We covered the sample in sections 05 and 06. Here we add size: twenty people. That is enough to see strong consensus, like the moves, but too few to split by age or culture. We argue a few more. The authors grouped and coded every gesture themselves, with no reliability check. The statistics treat the 1,074 gestures as if 1,074 different people had made them. Twenty did, so every p-value comes out more confident than it should. Participants rated their own gestures, yet the paper says this "goes a long way to validating" the approach. Five of the 27 commands are desktop ideas, cut, paste, minimize, maximize and task switch, and asking desktop questions invites desktop gestures. And the set matches 57 percent of what participants did, so the other 43 percent of their gestures get nothing. The paper claims "ease of recognition" with no recognizer. And two authors were at Microsoft Research, testing their own prototype, while Microsoft filed a patent application on the method that the paper does not mention. **[classmate]**
- **Transition:** Many of these limits are exactly what the course readings warn about, so let us connect the paper to them.
- **Watch out:**
  - Keep the tone fair. The authors flag the biggest limits themselves and call validation "an important next step" (p. 1091). Do not say the set "failed" validation; it was never tested for guessability (pitfall 6).
  - The exact limitation sentence is "Our participants were all non-technical literate American adults; undoubtedly, children, Eastern, or uneducated participants would behave differently" (p. 1091). The word "American" matters for section 06; the paper reports no ethnicity.
  - Self-ratings are not circular: participants did not know group sizes when they rated. The weakness is that nobody independent judged the gestures. The GI 2010 follow-up added independent judges (see section 11).
  - Sample size, if challenged: the Villarreal-Narvaez et al. review of 216 elicitation studies (DIS 2020, p. 859) found an average of 25 participants (SD 4), and 20 was the most frequent choice (36% of studies), which the review says is "likely explained by the fact that Wobbrock et al. [123] employed 20 participants." So 20 became the norm partly because of this paper. Norm is not the same as sufficient.
  - Desktop concepts among the 27 referents: cut, paste, minimize, maximize, task switch (Table 1, p. 1086). The paper says the referents were "obtained from desktop and tabletop systems" (p. 1085).
  - 57.0% coverage counts every proposal matched by any gesture in the set, aliases included (p. 1087); 43% is the remainder, not a measured failure rate.
  - Patent: US 2010/0031202, inventors Morris, Wobbrock, Wilson, filed 2008-08-04, abandoned 2014-12-09. Say "the paper does not mention it," never that it was hidden or that it drove the study (see section 03).
  - Minor slip, only if asked: the Discussion says "Figure 2 indicates that agreement is low after the first seven referents" (p. 1091); it means Figure 3.
- **Sources:** p. 1083 (affiliations), p. 1085 (20 participants, referents from desktop and tabletop systems), p. 1086 (Table 1; "The authors manually classified each gesture"), p. 1087 (grouping, largest group wins, coverage 57.0%, chi-square N = 1074), p. 1088 (chi-square tests, "goes a long way to validating"), p. 1091 (limitations, next steps, "ease of recognition," Figure 2 slip). Outside: https://faculty.washington.edu/wobbrock/pubs/dis-20.01.pdf (DIS 2020, p. 859); https://patents.google.com/patent/US20100031202A1 (via progress.md).


---

# 10 Connections to other research

**Purpose:** Place the paper against the four course readings from weeks 1 to 3 and against three of its own references, each with a "so what" for how we judge it. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 09 Limitations. **Next section:** 11 Research directions.

## Slide 10.1: Four readings, three roots

**Time:** 1:30 · **Visual:** a two-column card grid. Left column: four cards, one per course reading (author, year, one short quote). Right column: the matching point in the paper with its page. Bottom strip: three small "roots" chips (Furnas, Hutchins et al., Morris). No figure from the paper.

**On the slide**
- **Wobbrock and Kientz 2016:** methods need "repeated validation" (p. 41). Here one study (p. 1088); set validation a "next step" (p. 1091).
- **Greenberg and Buxton 2008:** early evaluation converges on "that which we are already familiar with" (p. 115). Elicitation too: about 72% mouse-like (p. 1090).
- **Olsen 2007:** "new design participants." Non-technical users author the set (p. 1084).
- **Weiser 1991:** wall boards at 1024 x 768 for walk-up "casual use." This Surface: same resolution, for public tables where "immediate usability" matters (p. 1085).
- **Roots:** Furnas et al., aliasing (p. 1088); Hutchins et al., gulf of execution (p. 1085); Morris et al., gestures for groups, yet this study tested one person at a time (p. 1084).

**Speaker notes**
- **Say:** Four of our readings each change how we judge this paper. Wobbrock and Kientz sort research into seven contribution types. We argue the set is an artifact, the findings are empirical, but the core is a method. Wobbrock's own 2012 whitepaper lists this exact paper as a methodological example. And methods, Wobbrock and Kientz write, need "repeated validation." This paper has one study, whose own ratings the paper says go "a long way to validating" the approach. Greenberg and Buxton warn that early evaluation converges on "that which we are already familiar with." We argue elicitation carries the same risk: about 72 percent of gestures were mouse-like. Olsen credits tools that bring "new design participants" into design. We argue that is this method's claim: non-technical people author the set. His last test, "better" designs, we think is still open. Weiser's boards ran at 1024 by 768, like this Surface, built for "casual use" with no sign-up. We argue the paper's "immediate usability" in public spaces goes further: no learning. Inside its own references: Furnas explains the aliases, Hutchins, Hollan and Norman explain the missing recognizer, and Morris's group gestures make a single-user study a question. (about 195 words)
- **Transition:** Both Wobbrock and Kientz's "repeated validation" and Olsen's open "better" test point at what the authors themselves list as next, so here are the research directions.
- **Watch out:**
  - **Which Wobbrock paper lists this one?** The published Interactions article (2016, pp. 38 to 44) does *not* name this paper. Wobbrock's 2012 whitepaper "Seven Research Contributions in HCI" (p. 3, "Examples of Methodological Contributions") does, next to Wobbrock et al. 2005 (the guessability paper, ref. 33). The 2016 article says it summarizes that whitepaper (p. 40, endnote 3, p. 44). Say "his whitepaper," not "the reading."
  - **The paper's own contribution list (p. 1084) does not name the method.** It lists (1) a characterization of gestures with a taxonomy, (2) the gesture set, (3) insight into mental models, (4) implications for technology and design. The method appears in the abstract ("We present an approach to designing tabletop gestures," p. 1083). Our mapping, labelled as our argument: (1) empirical plus theoretical (the taxonomy is a framework), (2) artifact (a design), (3) empirical, (4) design implications drawn from empirical data. If a classmate disagrees, the honest answer is that Wobbrock's whitepaper says "some research articles make more than one type of contribution" (p. 1), 30.4 percent of CHI 2016 submissions ticked more than one (Wobbrock and Kientz, p. 44), and our mapping is a judgment.
  - **Greenberg and Buxton are about usability evaluation, not elicitation.** Do not claim they discuss elicitation. Our reframing: (a) under their argument, skipping usability evaluation of an early-stage design is not automatically a flaw, but the paper presents the set as "a good candidate for deployment" (p. 1091), which is not early-stage, so we argue their excuse only partly applies; (b) but their core warning, convergence on the familiar, applies to elicitation because asking 2008 users captures 2008 habits ("I'm a child of the mouse," p. 1090). They also say early evaluation "can mute creative ideas that do not conform to current interface norms" (p. 111) and list "participatory critique" among alternative ways to validate (p. 118). Fact for context only: Buxton's affiliation on the paper is Microsoft Research (p. 111), the same lab as Morris and Wilson; do not suggest any influence.
  - **Olsen's criteria are for UI systems and toolkits, not methods.** Applying "empowering new design participants" to a method is our analogy. Olsen's three steps for that claim: describe the population and why it should be involved; establish why existing tools are not acceptable for it; demonstrate the new tools are "accessible, easier or more effective" for it. He then grades the claims, the last being that the resulting designs are "better," where "The question of defining 'better' is problematic" (pre-print p. 5). The paper covers step 1 (non-technical users, p. 1084), arguably step 2 (designer gestures "chosen out of concern for reliable recognition," p. 1083; each author matched only 43.5%, p. 1090), and only self-ratings for step 3 and the "better" claim ("goes a long way to validating," p. 1088). A second, sharper Olsen link if asked: he wrote that touch tables produce hand-sized input, "yet we force such techniques into the standard mouse point model" (pre-print p. 1). This paper found the users doing the same: 1 to 3 fingers counted as a "single point" (p. 1090).
  - **Weiser's boards are vertical whiteboards about 40 by 60 inches; the Surface was a 24 by 18 inch table** (p. 1085). Size falls between his pad and board scales. The match is the shared 1024x768 resolution and the walk-up, no-sign-up use, not the form factor. Weiser does not discuss tabletop gestures. Two further Weiser links if time allows: his "chalk" worked "either in contact with the surface or from a distance," and the paper says tables "may benefit from the ability to look down or sideways at users' hands" (p. 1091); he warned a "Macintosh-style menu bar may not be a good idea" on a board, while participants tapped imaginary widgets at the top-right and screen bottom (p. 1090).
  - **Page numbers for two readings:** the Weiser copy used (a reformatted Word version) carries no Scientific American page numbers, so quote the sentence, not a page. The Olsen copy is a 7-page author copy dated March 2007 (PDF metadata), not the proceedings layout; the UIST proceedings version is pp. 251 to 258, so give section names ("Forces for change," "Empowering new design participants").
  - **[classmate]** slot: if a classmate's commentary raised contribution type or the lack of validation, credit it on the Wobbrock and Kientz or Greenberg and Buxton line.
- **Sources:** Paper pp. 1083, 1084, 1085, 1088, 1089, 1090, 1091. Readings: Wobbrock and Kientz, Interactions 23(3), 2016, https://faculty.washington.edu/wobbrock/pubs/interactions-16.pdf ; Wobbrock 2012 whitepaper, https://faculty.washington.edu/wobbrock/pubs/Wobbrock-2012.pdf ; Greenberg and Buxton, CHI 2008, https://groups.cs.umass.edu/wp-content/uploads/sites/8/2019/01/reading18.pdf ; Olsen, UIST 2007 pre-print, http://web.archive.org/web/20120601131116/http://icie.cs.byu.edu:80/ResearchGuidance/SystemsResearchCriteria.pdf ; Weiser, Scientific American 1991, https://www.socs.uoguelph.ca/~qmahmoud/teaching/fall2006/pervasive/computer-21century.pdf (cross-check copy: https://graphics.stanford.edu/courses/cs428-03-spring/Papers/readings/General/Weiser_SciAm91.htm ); Furnas et al. abstract, https://honnef.co/notes/references/furnasvocabularyproblemhumansystem1987


---

# 11 Research directions

**Purpose:** show the research agenda the paper sets itself, what actually followed, and which questions are still open for a student today. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 10 Connections to other research. **Next section:** 12 Discussion bridge.

## Slide 11.1: What the paper asked for, and what came next

**Time:** 1:30 · **Visual:** two columns, "The authors' agenda (p. 1091)" on the left and "What followed" on the right; a strip along the bottom with the Figure 4 "pinch" panel (Enlarge (Shrink)3, p. 1089) and three open questions

**On the slide**
- **Authors' agenda (p. 1091).** Next: reverse the study (clips to commands); build a vision-based recognizer. Later: see all commands first; application context; "children, Eastern, or uneducated participants." Technology: sense above the table and beyond the edge.
- **What followed: the method became a standard**
  - **2020, Villarreal-Narvaez, Vanderdonckt, Vatavu and Wobbrock:** a review of 216 studies that used this method. It calls this paper the first of its kind and the most cited, and notes that 20 participants became the norm because of it
  - **2010, the same three authors:** 22 new people judged the gestures without knowing who made them. They preferred the gestures of larger groups. That tests liking, not guessing
  - **2014, Morris, Wobbrock and colleagues:** named the old-habits problem "legacy bias" and proposed fixes: several gestures per command, show the device's abilities first, work in pairs
  - **2015, Vatavu and Wobbrock:** corrected the agreement score (this paper's 0.32 becomes 0.28) and added a significance test
- **Twist (we argue).** No product adopted the set as published; pinch became universal anyway (Fig. 4, p. 1089).
- **Open.** Prior exposure as a variable? The reverse test? Other cultures, mid-air, VR?

**Speaker notes**
- **Say:** The paper ends with an agenda on page 1091. Next steps: reverse the study, show unlabelled clips and ask people to guess the command, then build a recognizer. Left for later: let people see all commands first, add application context, study children and other cultures. Technology: sense above the table and beyond the edge.
  What followed is the strongest evidence that the method is the lasting contribution. By 2020 a review by Villarreal-Narvaez, Vanderdonckt, Vatavu and Wobbrock had counted 216 studies using it, called this paper the first of its kind and the most cited, and found that twenty participants became the norm because this paper used twenty. Three refinements: in 2010 the same authors had 22 new people judge the gestures blind, which tests liking, not guessing; in 2014 Morris, Wobbrock and colleagues named legacy bias and proposed fixes; in 2015 Vatavu and Wobbrock corrected the agreement score.
  The twist: no product adopted the set as published, yet pinch is everywhere. Apple shipped it in 2007, and this study still produced it from people screened for no touch-screen experience. We argue three questions are open for you. **[classmate]**
- **Transition:** Those open questions lead straight into our own reading of the paper, and the tensions we want you to argue about.
- **Watch out:**
  - Do not say the reverse guessability study or the recognizer were later built. We did not check, and the GI 2010 paper tested preference, not guessability (pitfall 6).
  - Do not shorten GI 2010 to "users preferred user-authored gestures." The finding is about group size, and partly about the researchers' gestures being more physically and conceptually complex.
  - The "Eastern" wording is the authors' own (p. 1091); quote it, do not paraphrase it as "non-Western."
  - The pinch numbers: the paper says "only 1 of 5 used a diagonal 'reverse pinch' to resize along the diagonal, while 4 of 5 resized in other orientations" (p. 1091). It does not say 5 of 20, nor one hand. Say "five participants who made a reverse pinch" if asked.
  - "A year later" is approximate: iPhone shipped June 2007; the study was complete before the patent filing of 2008-08-04 (pitfall 15); the paper gives no study date.
  - If a classmate says the market, not elicitation, made pinch universal: agree that the study cannot tell the two apart. Screening removed iPhone owners, not exposure to iPhone adverts (pitfall 13, section 06).
  - The three open questions, if time allows or in discussion, each tied to a gap in the paper:
    1. Record participants' prior device exposure and analyse it as a covariate. Gap: no touch history was recorded, only the screening criterion (p. 1084, p. 1085).
    2. Run the reverse guessability test the authors proposed, with the user-defined set against the designers' sets. Gap: "An important next step is to validate our user-defined gesture set" (p. 1091).
    3. Replicate across cultures and in new modalities (mid-air, augmented and virtual reality hand tracking, wearables, voice phrasing), with legacy-bias controls. Gap: all participants were "non-technical literate American adults" (p. 1091).
- **Sources:** p. 1091 (agenda, limitations, sensing implications, pinch counts); p. 1089 (Figure 4, Enlarge (Shrink)3: pinch); p. 1084 and p. 1085 (screening). GI 2010: graphicsinterface.org, faculty.washington.edu/wobbrock (progress.md, verified 2026-09-30). Interactions 2014: https://interactions.acm.org/archive/view/may-june-2014/reducing-legacy-bias-in-gesture-elicitation-studies and https://faculty.washington.edu/wobbrock/pubs/interactions-14.pdf. CHI 2015: faculty.washington.edu/wobbrock/pubs/chi-15.01.pdf. DIS 2020: faculty.washington.edu/wobbrock/pubs/dis-20.01.pdf. iPhone dates: https://en.wikipedia.org/wiki/IPhone_(1st_generation).


---

# 12 Discussion bridge: our reading, the tensions, the hand-over

**Purpose:** close the presentation with the presenters' own reading of the paper, the tensions the class will argue about, and the plan for the 30-minute discussion. **Slides:** 3. **Time:** 3:00. **Speaker:** Oliver (sole presenter). **Previous section:** 11 Research directions. **Next:** the discussion itself (plan in a separate document, to be written).

Source material: revision 2 of `Week 4 - Slides.md`, slides 20 to 22, which passed two independent Opus reviews after correction. Written by the orchestrator, 2026-09-30.

## Slide 12.1: Our reading of the paper

**Time:** 0:45 · **Visual:** one sentence large, three supports beneath

**On the slide**
- **The lasting contribution is the method. The evidence in this paper for the method is thinner than its influence.**
- The guessability test the authors proposed was never run (p. 1091). The 2010 follow-up tested preference, which is partial support.
- "Popularity identifies better gestures" rests on participants rating their own gestures (p. 1088). No independent judges.
- The study measured the repertoire participants brought in. The paper calls that a finding about users. We argue it is also a property of the method.

**Speaker notes**
- **Say:** One sentence from us: the lasting contribution is the method, and the evidence in this paper for the method is thinner than its influence. Three reasons. The guessability test the authors proposed was never run; the 2010 study tested preference, which is partial support. The claim that popularity identifies better gestures rests on people rating their own ideas. And the study measured what participants brought in with them; the paper calls that a finding about users, and it is also a property of the method. Not a bad paper. A paper whose influence outran its evidence, which is worth arguing about.
- **Transition:** So here is what we want to argue about.
- **Watch out:** Say "thinner than its influence," not "weaker than its reputation"; the deck never describes a reputation. Do not say "unrevised behaviour assumes no prior habits"; the paper's term means not revised by recognizer feedback (p. 1085). If challenged that the designer comparison and GI 2010 are evidence for the method, agree: they are, and they are the strongest the paper and its follow-up offer; neither measures guessability.
- **Sources:** p. 1085 (unrevised behavior), p. 1088 (goodness ratings, "goes a long way to validating"), p. 1090 (designer comparison), p. 1091 (validation as next step); Morris, Wobbrock and Wilson, GI 2010, pp. 261 to 268.

## Slide 12.2: Tensions worth arguing about

**Time:** 1:15 · **Visual:** four large questions, three smaller items beneath, with **[classmate]** credits beside the ones classmates raised

**On the slide**

Four questions with more than one defensible answer
1. **Should the most popular gesture win?** About 72% of gestures were mouse-like (p. 1090). Does novice consensus lock in old habits?
2. **Did the commands shape the answers?** Cut, paste, minimize, maximize and task switch are desktop concepts (p. 1085).
3. **Does popularity identify better gestures, or ones people enjoyed inventing?** Only the gestures' own makers rated them (p. 1088).
4. **Is behavior without feedback the right design target?** Participants gestured in the air and off-screen (p. 1090).

Three smaller things we found interesting
- **Most guessable, or best?** 19.1% of each author's gestures were never tried by any participant (p. 1090).
- **Symbolic gestures import writing.** Check, X and question mark are Western marks (Figure 4, p. 1089).
- **Single users on a group device.** Minimize moves to the bottom "as defined by user's seating position" (Figure 4), different for each person at the table.

**Speaker notes**
- **Say:** Four questions with more than one defensible answer. Should the most popular gesture win, when 72 percent were mouse-like and consensus from novices may lock in old habits? Did the commands shape the answers, with five desktop concepts on the list? Does popularity identify better gestures, or gestures people enjoyed inventing, and what would independent validation look like? And is behaviour without feedback the right target, when real users adapt to a system within minutes? Then three smaller things. Nineteen percent of the designers' gestures were never tried by anyone: is elicitation a floor or a ceiling? The check mark, the X and the question mark are Western writing, and the paper itself says Eastern participants would differ. And minimize means drag toward your own edge of the table, a different direction for every person sitting at a group device. **[classmate]**
- **Transition:** Before we open it up, here is what you wrote.
- **Watch out:** Read these as invitations, not conclusions. Credit each classmate who raised one. Question 3 must not be called "circular": participants did not know the group sizes when they rated (p. 1088). The "adapt within minutes" claim has no source in the paper; say it as common experience, not as a finding.
- **Sources:** p. 1085 (referents), p. 1088 (ratings), p. 1090 (mouse quotes, off-screen gestures, 19.1%), p. 1091 ("children, Eastern, or uneducated"), Figure 4 p. 1089 (marks, Minimize2 caption).

## Slide 12.3: Over to you

**Time:** 1:00 · **Visual:** three or four short quotes from classmates' commentaries with names, each paired with the tension it speaks to (filled after the commentaries are posted), then the three-step plan

**On the slide**
- **What you wrote.** **[classmate]** **[classmate]** **[classmate]**
- **How the next 30 minutes run**
  1. **Warm-up, 5 min: a live elicitation.** We show the effect of two commands, one the paper's participants agreed on and one they did not (Figure 3). On three, everyone performs the gesture they would use, on the desk. We count the groups and compute the agreement score together.
  2. **Pairs, 5 min:** pick one of the four questions and take opposite sides.
  3. **Full group, 20 min:** we follow what comes up.

**Speaker notes**
- **Say:** **[classmate]** (name three or four themes from the commentaries, credit each person, and say which tension each feeds). Here is how the next half hour runs. First a five-minute live elicitation: we show you the effect of two commands, one from the top of Figure 3 and one from the tail. On three, everyone performs the gesture they would use, on the desk. We count the groups and compute the agreement score together, so you see the method work and see what it misses. Then five minutes in pairs: pick one of the four questions and take opposite sides. Then twenty minutes together, and we follow whatever is interesting.
- **Transition:** none; the discussion starts.
- **Watch out:** Choose "move a little" and "undo" or "help." Expect near-total agreement on move and a spread on undo, including smartphone-era gestures such as shake, which shows legacy bias in the room. Tally on the board, not on a slide. Keep the agreement arithmetic ready: square each group's share, add. The full discussion plan is a separate document still to be written.
- **Sources:** Figure 3 p. 1088 for the top and tail commands; Eq. 1 p. 1087 for the live score.
