# Agent briefs for the per-section presentation files

Written 2026-09-30 by the orchestrating session. Each writer agent reads **Part 1 (common rules)** and **its own brief in Part 2**, writes exactly one file in this folder, and returns a short summary. Agents work independently and in parallel; none sees another's output. The orchestrator verifies every file afterwards and writes `00-Overview.md` and `12-Discussion-Bridge.md` itself.

Section order and hand-offs (each brief names the next section so transitions connect):

| File | Section | Slides | Time |
|---|---|---|---|
| 01-Background-2008.md | Background: 2008, the devices, the motivation, the setup | 2 | 2:30 |
| 02-Problem.md | The problem the paper addresses | 1 | 1:15 |
| 03-Pressure-and-Ownership.md | Who owned gesture research and IP; was Microsoft under pressure | 1 | 1:30 |
| 04-Method.md | The method, in depth: the most valuable part | 2 | 3:30 |
| 05-How-Conducted.md | How the research was run; designers' inaccuracy; one vs two hands; fingers | 2 | 3:00 |
| 06-Sample-and-Culture.md | Who was missing: age, ethnicity, culture; the tainted 2008 adult | 1 | 1:30 |
| 07-Findings.md | The most important findings, including the timing results | 2 | 2:30 |
| 08-Technical-Obstacles.md | Technical obstacles of the time | 1 | 1:30 |
| 09-Limitations.md | Limitations of the study | 1 | 1:30 |
| 10-Connections.md | Connections to the course readings and other research | 1 | 1:30 |
| 11-Research-Directions.md | Research directions the paper suggests | 1 | 1:30 |
| 12-Discussion-Bridge.md | Thesis, tensions, hand-over to discussion (orchestrator) | 3 | 3:00 |

Title slide 0:30. Total 25:45 for a 30-minute slot.

---

# Part 1. Common rules for every writer agent

## 1.1 What you are producing

One Markdown file for one section of a 30-minute graduate seminar presentation (CMPT 863, SFU) on Wobbrock, Morris and Wilson, "User-Defined Gestures for Surface Computing," CHI 2009, pp. 1083 to 1092. The class has read the paper. The course brief says: do not summarize section by section; help the class understand the contribution, evaluate it critically, and engage with broader implications. Grading: presentation quality (engaging, polished, insightful), integration of classmates' commentaries, discussion. The mark matters a great deal to the presenter, so accuracy is the first priority and clarity the second.

## 1.2 Ground truth and materials

All paths are under `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/Presentation-Week4/`.

- **The paper (ground truth):** `Week 4 - User-Defined Gestures for Surface Computing.pdf`. Ten pages. PDF page N is printed page 1082 + N. Read it with the Read tool (`pages: "1-5"` then `"6-10"`) or `pdftotext -layout -f N -l N "<pdf>" -`. Tables 1 and 2 are on PDF page 4; the agreement equations (Eq. 1 to 3) and Figure 2 on page 5; Figure 3 on page 6; Figure 4 on page 7; references on page 10. To see a figure, render it: `pdftoppm -r 170 -f 6 -l 6 -png "<pdf>" /tmp/claude-1000/-mnt-RAID-STORAGE-SFU-Projects-CMPT-863-Presentation-Week4/202abb58-4973-4aee-b8ad-95b0d3c365f5/scratchpad/<yourname>` and open the PNG with Read.
- **Verified prior work you may reuse:** `Week 4 - Slides.md` and `Week 4 - Speaker Notes.md` (revision 2, already reviewed by two independent Opus agents and corrected), `Week 4 - Presentation Notes.md` (background and the seven course questions), `progress.md` (section "Verified facts" lists what was checked and where, including a Microsoft patent application on the method), and `reports/slides-review-cross-examination.md` (every error found in the first draft, with the fix). Reuse verified content freely. Do not reintroduce a corrected error.
- **Outside facts:** verify with WebSearch or WebFetch and record the URL. If you cannot verify a fact, keep it only if it is essential and tag it `[UNVERIFIED]` on the slide and in the log. Never invent a citation, a date, a quote or a number.

## 1.3 Known pitfalls (each of these was wrong in the first draft; do not repeat them)

1. "Unrevised behavior" (p. 1085) means behaviour not revised by recognizer feedback during the session. The paper never claims participants free of prior habits, and it reports the desktop bias itself as a finding (abstract p. 1083; "It's a Windows World," p. 1090).
2. The agreement score A is not a percentage. Range is 1/|P| to 1, floor about 0.05 with 20 people. Overall A = 0.32 one hand, 0.28 two hands (p. 1087).
3. Figure 3 values (one hand): move a little and move a lot 1.0; duplicate 0.52; rotate, select group, insert about 0.48; select single 0.37; after the first seven nothing above about 0.32; tail runs task switch 0.21, zoom out 0.20, previous 0.19, undo 0.18, accept 0.17, help 0.16, reject 0.15, next 0.13, cut 0.11. Two-hand line: peaks insert 0.58, maximize 0.58, enlarge 0.52; dips pan 0.19, close 0.19, task switch 0.13. Say "about" for chart readings.
4. The widget recommendations on p. 1091 are two separate things: beyond the first seven referents "an on-screen widget as well as a gesture"; and, because participants used imaginary widgets, consider widgets for "delete, zoom in, zoom out, accept, reject, menu access, and help." Undo is not on that list.
5. The 0.37 for select single (one hand, groups 11, 3, 3, 3) is printed in the paper as Eq. 3, p. 1087. The 0.42 for move a little is the two-hand score (Eq. 2).
6. The gesture set was never validated for guessability; validation is "an important next step" (p. 1091). The GI 2010 follow-up tested preference: 22 participants blind to authorship judged 81 gestures and preferred those authored by larger groups, partly because the researchers' gestures were more physically and conceptually complex. Do not shorten that to "users preferred user-authored gestures."
7. Figure 4 labels: "hold and tap" is Select Group1; Move2 is "jump" (hold object, tap destination) and is reused by Delete2, Paste3, Minimize2; Duplicate is "tap source and destination"; Menu is "pull out," a finger of the second hand dragged out from a held object, not from the screen edge; Minimize1 is "drag to bottom of surface." Any one-finger gesture may use 1 to 3 fingers (caption).
8. Long et al.: the paper says "users are sometimes poor at picking easily differentiable gestures" (p. 1084). Not "designers are poor."
9. Epps et al. "may have biased participants to simply emulate the mouse" (p. 1084). Not "had pushed."
10. Referents were "obtained from desktop and tabletop systems" (p. 1085). Several are desktop concepts (cut, paste, minimize, maximize, task switch).
11. Grouping: the paper says only "we grouped the gestures within each referent such that each group held identical gestures" and "the referent with the largest group won the gesture" (p. 1087). It does not say how finger count was treated in grouping, nor what happened to the referent that lost a conflict. Present anything beyond that as inference.
12. Four referents got no dedicated gesture: insert and maximize (users reused dragging and enlarging), task switch and close (users tapped imaginary widgets, "a common gesture was not feasible," p. 1087). 22 of 27 got dedicated gestures with the two move referents combined.
13. Participants: 20 paid volunteers, 9 women, average age 43.2 (sd 15.6), all right-handed, "no participant had used an interactive tabletop, Apple iPhone, or similar" (p. 1085); recruited as "non-technical people without prior experience using touch screens (e.g., the Apple iPhone)" (p. 1084). No age distribution, no ethnicity, no touch history reported.
14. Designer comparison (p. 1090): each author's set matched 43.5% of the user set, all three combined 60.9%, and 19.1% of each author's gestures were never tried by any participant. This measures overlap with the users' set, not quality; the paper concedes "three experts cannot generate the scope of gestures that 20 participants can."
15. The paper gives no study date. Microsoft's patent application on the method (US 2010/0031202, inventors Morris, Wobbrock, Wilson) was filed 2008-08-04 and abandoned 2014-12-09, so the study was complete before August 2008. The paper does not mention the application.
16. The Discussion says "Figure 2 indicates that agreement is low after the first seven referents" (p. 1091); it means Figure 3.
17. Timing (p. 1088): complexity correlates with planning time, r = 0.71; simple referents about 8 s of planning, complex about 15 s. Longer planning went with lower goodness ratings. Articulation time did not affect goodness but affected ease: gestures rated easy took about 3.4 s, those rated difficult about 2.0 s. Goodness ratings: simple referents 5.6, complex 4.9 on the 7-point scale; complexity vs goodness r = -0.59.
18. Statistics: the chi-square tests use N = 1074 gestures from 20 people as independent observations (pp. 1087 to 1088). Grouping and taxonomy coding were done by the authors with no inter-rater reliability reported.
19. 1,080 gestures = 20 x 27 x 2; 6 discarded for confusion, leaving 1,074. Final set: 48 gestures, 31 one-hand (64.6%), 17 two-hand (35.4%); ten referents have 1 gesture, four have 2, three have 3, four have 4, one has 5 (p. 1088). Coverage 57.0% of all proposals, aliases included.
20. Figure 2 reports Form per hand. Right hand: one-point path .450, one-point touch .274 (together the "72% mouse-like" of p. 1090), static pose and path .140, dynamic pose .092, static pose .037, dynamic pose and path .007. Nature: physical .439, abstract .272, metaphorical .189, symbolic .100. Binding: object-centric .702, world-dependent .182, world-independent .089, mixed .027. Flow: discrete .535, continuous .465.

## 1.4 File format (follow exactly)

```
# <NN> <Section title>

**Purpose:** one sentence. **Slides:** n. **Time:** m:ss. **Speaker:** A or B (leave as given in the brief). **Previous section:** <name>. **Next section:** <name>.

## Slide <NN.1>: <Title>

**Time:** m:ss · **Visual:** what to show (figure, table, timeline, quote card), with PDF page if from the paper

**On the slide**
- terse bullets, at most about 90 words of on-slide text in total (reviewers found the first draft too dense)
- every number carries its printed page in brackets, e.g. (p. 1088)

**Speaker notes**
- **Say:** the spoken script, at most 130 words per minute of the slide's time. Spoken register. Label judgments: "the paper says" versus "we argue".
- **Transition:** one sentence that leads into the next slide, or into the next section for the last slide.
- **Watch out:** the traps for this slide, including anything a sharp classmate might challenge and how to answer.
- **Sources:** printed page numbers for paper facts; URLs for outside facts.

(repeat per slide)

## Verification log

| Claim | Where checked | Status |
|---|---|---|
| every number, quote, date and attribution on the slides | p. NNNN or URL | OK / UNVERIFIED |

## Open items
- anything you could not verify, any conflict you found with the pitfalls list or the paper, any suggestion for the orchestrator (for example a cross-reference to another section).
```

## 1.5 Style rules

- **No em-dashes anywhere** (the character "—"). Do not use en-dashes as dashes either. Use a comma, a full stop, or "to" for ranges ("pp. 1083 to 1092", "8 to 15 seconds"). Hyphens in compound words are fine.
- Short sentences. Plain words. Expand an acronym the first time. No rhetorical questions in on-slide text except where the brief asks for a question.
- Quote the paper exactly when you quote; otherwise paraphrase without quotation marks.
- Say "the paper says" for the authors' claims and "we argue" or "we think" for the presenters' judgments. Never attribute a motive to the authors as fact.
- Credit slots: where a classmate's commentary could be inserted, write **[classmate]**.
- Do not write placeholders you could fill. The only allowed placeholders are **[classmate]** and `[UNVERIFIED]`.

## 1.6 Boundaries

- Write only your assigned file. Do not modify any other file.
- Do not spawn sub-agents.
- Scratch files go in `/tmp/claude-1000/-mnt-RAID-STORAGE-SFU-Projects-CMPT-863-Presentation-Week4/202abb58-4973-4aee-b8ad-95b0d3c365f5/scratchpad/`.
- When done, return to the orchestrator a summary of at most 150 words: file written, slide count, number of claims verified, every `[UNVERIFIED]` item, and anything that conflicts with the pitfalls list.

---

# Part 2. Briefs per section

## Brief 01: Background, 2008 (file `01-Background-2008.md`, 2 slides, 2:30, Speaker A, next section: 02 The problem)

Goal: zoom the room into the year of the study, show what touch and gesture devices an ordinary adult could already have used, then give the motivation and the study setup as the foundation for everything after.

**Slide 01.1, "2008: where touch stood."** A timeline from 2005 to 2009 and a short device list. Timeline anchors to verify with URLs: Jeff Han's multi-touch demo (TED, February 2006); iPhone announced January 2007 and shipped June 2007 with pinch-to-zoom; Microsoft Surface announced May 2007; Surface sold to commercial customers from April 2008; the study completed before 2008-08-04 (patent filing, see progress.md); CHI 2009 in April; Windows 7 with multi-touch, October 2009. Device list, each with a verified date and what it taught the hand: ATMs and kiosks; Palm PDAs with Graffiti (1996); laptop touchpads (Synaptics, mid 1990s); iPod touch wheel (2002) and click wheel (2004); Nintendo DS (November 2004); Nintendo Wii (November 2006); FingerWorks multi-touch pads (founded 1998, bought by Apple 2005); research tables DiamondTouch (2001) and SmartSkin (2002); Bill Buxton's 1985 multi-touch tablet. Use Bill Buxton's page "Multi-Touch Systems that I Have Known and Loved" (billbuxton.com) as a primary source for the research lineage and cite it. Mark clearly which dates are verified and give the URL in the log. Drop any device you cannot date.

**Slide 01.2, "Why this study, and what it was."** Motivation from the paper: tabletops spreading to single users, groups, multi-display rooms and entertainment (p. 1083); public-space tables where "the importance of immediate usability is high" (p. 1085); keyboard and mouse "no longer preferable" (p. 1083); gestures until then designed by system designers for reliable recognition (p. 1083). Then the setup in four lines: 20 non-technical adults, a Microsoft Surface prototype, 27 commands, see the effect then perform the gesture, one hand and two hands, think aloud, no recognizer, 1,080 gestures, output taxonomy plus agreement score plus a 48-gesture set for 22 commands. Keep the setup minimal: section 05 covers the detail. End with one sentence on why the moment mattered: gesture vocabularies were being fixed in products at exactly this time, and two of the three authors worked at Microsoft Research.

## Brief 02: The problem (file `02-Problem.md`, 1 slide, 1:15, Speaker A, next section: 03 Pressure and ownership)

Goal: state the problem the paper addresses and why it mattered, with analysis rather than summary.

Content: the stated problem in the authors' words (p. 1083, "defined by system designers ... chosen out of concern for reliable recognition"); the deeper problem, that the field built gesture vocabularies from the recognizer outward with no measurable way to find what people would do (RoomPlanner and Wu et al. "were not elicited from users, although usability studies were conducted," p. 1085; Malik et al. distinguishing 1-, 2-, 3- and 5-finger gestures, which "may be problematic," p. 1085; Long et al., pitfall 8; Furnas's vocabulary problem, ref. 8); why the design space makes it worse ("almost anything one can do with one's hands could be a potential gesture," p. 1083); why it mattered then (public tabletops, products fixing conventions) and why it still matters (any new input channel without a vocabulary). Note honestly that some elicitation existed before: Good et al. 1984, Nielsen et al. 2004 ("a similar approach"), Epps et al. 2006, Beringer, Voida (p. 1084). The paper's claim is to be "the first to employ users, rather than principles, in the development of a gesture set" (p. 1084), which means first complete set from user consensus. Give a tagline that is accurate: not "nobody had asked."

## Brief 03: Pressure and ownership (file `03-Pressure-and-Ownership.md`, 1 slide, 1:30, Speaker A, next section: 04 The method)

Goal: show who owned gesture research and gesture intellectual property around 2005 to 2009, and test the presenter's hypothesis that Microsoft was under pressure to progress. The presenter also asked whether the study was done to avoid lawsuits. Be scrupulous: facts with URLs on one side, interpretation labelled "we argue" on the other. Never state a motive of the authors as fact.

Facts to verify (each with a URL in the log; drop what you cannot verify):
- FingerWorks: founded 1998 by Wayne Westerman and John Elias (University of Delaware), multi-touch gesture pads, acquired by Apple in 2005.
- Steve Jobs at the iPhone launch, January 2007: multi-touch, "and boy, have we patented it" (verify the wording).
- Apple's US patent 7,479,949 ("Touch screen device, method, and graphical user interface for determining commands by applying heuristics"), granted January 2009, often called the "Jobs patent."
- Tim Cook, Apple earnings call 21 January 2009: "we will go after anyone" who rips off Apple's IP, read as aimed at the Palm Pre (announced January 2009, shipped June 2009, with pinch-to-zoom).
- Elan Microelectronics sued Apple over multi-touch patents in April 2009 (verify date and patent).
- Apple and Microsoft had a patent cross-licence from 1997, and testimony in Apple v. Samsung (August 2012, Boris Teksler) said it covered design patents with an anti-cloning clause.
- Microsoft's own position: Surface announced May 2007 and sold from April 2008; Windows 7 multi-touch demonstrated May 2008 (D6 conference); Bill Buxton joined Microsoft Research in 2005; Microsoft bought Jeff Han's Perceptive Pixel in 2012 (later, context only).
- Microsoft's patent application on this study's method: US 2010/0031202, "User-Defined Gesture Set for Surface Computing," inventors Morris, Wobbrock, Wilson, filed 2008-08-04, published 2010-02-04, abandoned 2014-12-09; claim 1 covers presenting an effect to several users, capturing their surface input by computer vision, and identifying the gesture by correlation. Already verified in progress.md; cite patents.google.com/patent/US20100031202A1.
- Research ownership: MERL DiamondTouch (2001), Sony SmartSkin (2002, Rekimoto), NYU Han (2005 to 2006), Microsoft Research Wilson's TouchLight (2004) and PlayAnywhere (2005). Verify dates from the paper's references (p. 1092) or Buxton's page.

What the slide must conclude, honestly: there is no evidence in the paper or elsewhere that the study was run to avoid lawsuits; the paper gives academic motives and continues Wobbrock's 2005 guessability work; the patent filing shows that Microsoft treated the method as intellectual property worth owning; the Apple-Microsoft cross-licence reduced the lawsuit risk between those two companies in particular; Apple's public threats in January 2009 came after the study was done. "Microsoft was under pressure" is defensible as "Microsoft was racing to establish its own gesture conventions and IP as Apple's were being fixed," and should be labelled "we argue." Include in Watch out how to answer a classmate who asks "so was it defensive?"

## Brief 04: The method (file `04-Method.md`, 2 slides, 3:30, Speaker A, next section: 05 How the research was conducted)

Goal: the presenter considers the method the most valuable part of the paper and wants to dig in. Two slides, the most carefully written in the deck.

**Slide 04.1, "The method: a guessability study, from referent to sign."** Lineage, verified from the paper: Good et al. 1984 built a "user-derived" command-line email interface from users' unconstrained commands (p. 1084); Wobbrock et al. 2005 introduced the guessability study and the agreement score for EdgeWrite unistrokes (ref. 33; "adopted from prior work [33]," p. 1087); McNeill's terms referent and sign (footnote, p. 1085); Nielsen et al. 2004 "describe a similar approach" (p. 1084). Then the three design choices with the paper's reasons: elicit from users because they must guess and "user behavior is rarely so systematic" (p. 1083); show the effect first in a neutral blocks world so meaning is fixed and only form is open, avoiding the bias Epps et al. conceded (pitfall 9); no recognizer, removing the gulf of execution (ref. 10), "a monologue in which the user's behavior is always acceptable," giving "unrevised behavior" in the sense of pitfall 1 (p. 1085). Add one line that Microsoft filed a patent on this procedure in August 2008 (pitfall 15), cross-referencing section 03.

**Slide 04.2, "From 1,080 gestures to a set, step by step."** Grouping into identical gestures (state exactly what the paper says and what it does not, pitfall 11); the agreement score with Eq. 1 written out in words, both worked examples (0.42 two-hand move a little, Eq. 2; 0.37 one-hand select single, Eq. 3), the range 1/|P| to 1 and why it is not a percentage; the conflict rule (largest group wins, p. 1087); aliasing and why ("Aliasing has been shown to dramatically increase input guessability [8,33]," p. 1088), with the 10/4/3/4/1 breakdown giving 48 gestures; coverage 57.0% and what it counts; the taxonomy as the descriptive layer (four dimensions, coded by the authors, Table 2 p. 1086); the designer baseline as a built-in test (p. 1090). Close with what the method cannot do: it assigns only what the largest group did (close and task switch got nothing, pitfall 12), it captures the repertoire people bring in, and guessability of the result is never measured (pitfall 6). Mention in one line why it mattered afterwards: a 2020 review covered 216 gesture elicitation studies (Villarreal-Narvaez, Vanderdonckt, Vatavu, Wobbrock, DIS 2020, verified in progress.md) and Vatavu and Wobbrock (CHI 2015) corrected the agreement measure.

## Brief 05: How the research was conducted (file `05-How-Conducted.md`, 2 slides, 3:00, Speaker B, next section: 06 Sample and culture)

Goal: show how the study was run and highlight three results the presenter wants on screen: the researchers' inaccuracy in predicting gestures, one hand versus two hands, and finger use.

**Slide 05.1, "The study as run."** Participants (pitfall 13, with occupations: restaurant host, musician, author, steelworker, public affairs consultant, p. 1085); apparatus (Microsoft Surface prototype, 24 x 18 inches, 1024 x 768, custom software, recorded voice plus animation, reset, p. 1085); the 27 referents and the authors' complexity ratings (Table 1, 1 to 5, mean 2.70); the trial (random order, one-hand and two-hand gesture, think aloud, preference, two 7-point ratings "good match" and "easy to perform," p. 1086 and p. 1088); data (1,080 gestures, 6 discarded, touch logs as ovals with millisecond timestamps, four video angles, two authors' notes, p. 1086); the designer baseline made before the study (p. 1090). State that the paper does not say whether one-hand always preceded two-hand.

**Slide 05.2, "Three things the run revealed."** (1) Designers' inaccuracy: the three numbers, the authors' joke "thinking creatively or ... hopelessly lost" (quote exactly from p. 1090), and the fair reading (pitfall 14). (2) One hand versus two: preferred for 25 of 27 referents, evenly divided for insert and maximize, two hands never preferred overall; yet 17 of 48 (35.4%) set gestures are two-handed because every participant had to produce a two-hand gesture; "some 2-hand gestures had good agreement scores and nicely complemented their 1-hand counterparts" (p. 1088); which ones (Figure 3: insert, maximize, enlarge). (3) Fingers: 13 of 20 varied finger count, only 2 said it mattered; 1 to 3 fingers read as a single point, 5 fingers or a palm as something more, 4 as a grey area (p. 1090, verify exact wording); design implication: do not distinguish gestures by finger count, contrast with Malik et al. (p. 1085). Verify each quote.

## Brief 06: Sample and culture (file `06-Sample-and-Culture.md`, 1 slide, 1:30, Speaker A, next section: 07 Findings)

Goal: one slide on who was missing from the sample and why it matters: under-representation of younger people and of other ethnicities and cultures, the cultural loading of the gestures in the set, and the "tainting" of the average 2008 adult by prior devices and the mouse.

Content:
- Age: mean 43.2, sd 15.6, no distribution, no children (the limitations mention children, p. 1091). Say what can be inferred (a spread that wide implies people in their twenties and sixties) and what cannot. No results are broken down by age.
- Ethnicity and culture: the paper reports no ethnicity. Find the exact wording of the cultural limitation on p. 1091 (it mentions "children, Eastern, or uneducated participants"). Quote it.
- Culturally loaded gestures in the set: check mark for accept, X for reject, question mark for help, a line drawn across the object for next and previous (reading direction), the pull-out menu (Figure 4). Verify with sources and give URLs: in Bulgaria a head nod means no and a head shake means yes (the presenter's example; verify, and name the countries carefully); the check mark is not universally "correct" (in Japan a circle marks correct and in some countries a check can mark wrong; verify before using); the thumbs-up or the OK sign is offensive in some regions (verify one example with a source). Use two or three verified examples, no more.
- Tainting: reuse the verified material from revision 2 slide 17 ("Whose habits did the study capture?"): the screening quotes (pitfall 13), "It's a Windows World" evidence (p. 1090: two-button-mouse clicking, imaginary close buttons top right, imaginary taskbar, "I'm a child of the mouse"), the devices a 2008 adult had used (ATMs, Palm, touchpads, DS), and the fact that no touch history was recorded. The point: the paper treats prior experience as a finding about users; we argue it is also a limit of the method, and the sample's age and culture decide which priors show up.
- Later work on culture: Mauney, Howarth, Wirtanen and Capra, "Cultural Similarities and Differences in User-Defined Gestures for Touchscreen User Interfaces," CHI 2010 Extended Abstracts, a study across several countries. Verify it exists and what it found before citing; tag `[UNVERIFIED]` otherwise.
Keep the on-slide text within 100 words; the notes carry the detail. Include a **[classmate]** slot.

## Brief 07: Findings (file `07-Findings.md`, 2 slides, 2:30, Speaker B, next section: 08 Technical obstacles)

Goal: the most important findings, ranked, with the timing results the presenter specifically wants: that gestures rated difficult were executed faster, that thinking time grows with complexity, and that thinking time dwarfs execution time.

**Slide 07.1, "What people did."** Ranked, each with a page: 72% one-point gestures (pitfall 20, note per hand); finger count rarely matters; one hand preferred; physical 44% and the link simple referents to physical gestures, complex to metaphorical and symbolic (p. 1087); dichotomous referents got reversible gestures even though they were never shown together (p. 1090, examples shrink/enlarge, previous/next, zoom in/out); people gestured off-screen and above the table and tapped imaginary widgets (p. 1090); agreement falls with complexity r = -0.52 (p. 1087); popular gestures were rated better by their makers (p. 1088, self-rating caveat); the desktop is "deeply rooted" (p. 1090). Give each a one-word solidity tag.

**Slide 07.2, "Time tells a story."** All from p. 1088, verify each: planning time rises with complexity (r = 0.71), about 8 s for simple referents and about 15 s for complex; longer planning went with lower goodness ratings ("good gestures were those most quickly apparent"); conceptual complexity correlates with goodness r = -0.59, simple 5.6 vs complex 4.9; articulation time did not affect goodness but did affect ease: gestures rated easy took about 3.4 s and those rated difficult about 2.0 s; the authors' explanation ("perhaps because they were smoother or less hasty"); touch-event counts as the authors' corroboration and why that is the same measurement twice. Make the presenter's points explicit: thinking time (8 to 15 s) is several times the execution time (2 to 3.4 s), and the counterintuitive direction has a plausible reading (long continuous drags are easy and slow; quick taps and marks at precise spots are hard and fast; difficulty in motor terms is precision, not duration). Caveats: means of skewed data; chi-square on N = 1074 non-independent gestures (pitfall 18).

## Brief 08: Technical obstacles (file `08-Technical-Obstacles.md`, 1 slide, 1:30, Speaker B, next section: 09 Limitations)

Goal: what the technology of the time could and could not do, and how it shaped the study.

Content to verify:
- Display: 24 x 18 inches at 1024 x 768 (p. 1085). Compute the pixel density (about 43 pixels per inch) and compare with one verified modern figure (for example a current iPad or a 27-inch 4K monitor), with a URL. Say what coarse resolution meant for small targets and precise positioning.
- Sensing: what the paper says about how the prototype sensed touch (find the exact wording on p. 1085; the Surface used cameras below the display; do not add details the paper does not give unless verified with a URL). Note that the system logged touches as ovals with timestamps (p. 1086).
- Physics engine: ref. 32 (verify the title and venue on p. 1092; it should be Wilson et al., "Bringing physics to the surface," UIST 2008). The paper's argument: 43.9% physical gestures support a physics engine; "Seven participants ... expected intervening objects to move out of the way"; "Four participants 'threw' an object" (verify on p. 1091); but "a physics engine alone will probably not suffice as an adequate recognizer for all surface gestures" (p. 1091).
- Recording: four video cameras, think-aloud notes by two authors (p. 1086). Keep the session recording distinct from the touch sensing.
- Gestures the table could not sense: participants gestured above the table and beyond the screen edge (find the exact sentences on p. 1090 and p. 1091, including what the authors recommend: sensing beyond the display edge and above the table). The presenter believes participants were clearly instructed to touch the surface. Check the procedure on p. 1085 to p. 1086 for any instruction about touching and report exactly what the paper says. If it says nothing, say so on the slide ("the paper does not say whether touching was required") rather than asserting an instruction.
- No recognizer was built, yet the paper claims the set has "ease of recognition" (verify wording on p. 1091).

## Brief 09: Limitations (file `09-Limitations.md`, 1 slide, 1:30, Speaker B, next section: 10 Connections)

Goal: the clear limitations of the study, referring back to section 06 for culture and age and to section 05 for the sample, and adding the small sample size.

Content: the authors' acknowledged limitations (p. 1091: no revising earlier gestures, no application or organizational or cultural context, the "children, Eastern, or uneducated" sentence, set not validated, recognizer not built); the sample size, 20 people (say what a sample of 20 can and cannot support; if you can verify a typical participant count in later gesture elicitation studies from the DIS 2020 review by Villarreal-Narvaez et al., cite it, otherwise leave it out); the limitations the presenters add (pitfall 18 on coding and statistics; self-ratings with no independent judges and the authors' "goes a long way to validating" on p. 1088; referents from desktop and tabletop systems with several desktop concepts; 43% of proposals not accommodated; "ease of recognition" claimed without a recognizer; two authors at Microsoft Research on their own prototype and the undisclosed patent application, pitfall 15; the Figure 2 for Figure 3 slip, pitfall 16, as a minor note). On-slide text within 100 words; the notes carry the detail. Cross-reference sections 05 and 06 explicitly. Include a **[classmate]** slot.

## Brief 10: Connections to other research (file `10-Connections.md`, 1 slide, 1:30, Speaker A, next section: 11 Research directions)

Goal: connect the paper to the course's first three weeks of readings and to the research it builds on. This file will be independently cross-checked by a second agent, so every connection must carry a quote or page reference from the reading and from the paper.

The course readings (fetch each; use WebSearch to find a PDF or HTML copy, then WebFetch; record the URL you used):
1. Mark Weiser, "The Computer for the 21st Century," Scientific American, 1991.
2. Jacob O. Wobbrock and Julie A. Kientz, "Research Contributions in Human-Computer Interaction," Interactions, 2016.
3. Saul Greenberg and Bill Buxton, "Usability Evaluation Considered Harmful (Some of the Time)," CHI 2008.
4. Dan R. Olsen Jr., "Evaluating User Interface Systems Research," UIST 2007.

For each reading, one concrete connection with a "so what" for this paper, grounded in a quote or page from the reading and a page from the paper. Leads to test, not to copy: Wobbrock and Kientz list contribution types (empirical, artifact, methodological, theoretical, dataset, survey, opinion); classify this paper's four listed contributions (p. 1084) and the method by their types, and note that Wobbrock co-authored both. Greenberg and Buxton argue against premature usability evaluation of early designs; consider how that reframes the critique that the gesture set was never validated, and whether elicitation is the kind of early-stage method they call for. Olsen's evaluation criteria for systems research (importance, problem not previously solved, generality, reducing solution viscosity, empowering new design participants, power in combination, scale); apply two or three to the elicitation method, especially "empowering new design participants." Weiser's vision of invisible, ubiquitous computing with tabs, pads and boards; a tabletop that needs a guessable gesture language is a board-sized device, and the paper's "immediate usability" in public spaces (p. 1085) is a Weiser-style requirement; check what Weiser actually says before claiming it.

Then two or three connections inside the paper's own references, each with a page: Furnas et al., the vocabulary problem (ref. 8) and aliasing (p. 1088); Hutchins, Hollan and Norman's gulf of execution (ref. 10, p. 1085); Good et al. 1984 (p. 1084); Wobbrock et al. 2005 EdgeWrite (ref. 33); Morris's cooperative gestures (ref. 17) against a single-user study of a group device; Wilson's physics work (ref. 32, p. 1091).

Add at the end a section `## Claims for the cross-checker`, a table with one row per connection: the claim, the exact quote or page in the reading, the page in the paper, and the URL used. The cross-checker will verify every row.

## Brief 11: Research directions (file `11-Research-Directions.md`, 1 slide, 1:30, Speaker B, next section: 12 Discussion bridge)

Goal: the research directions the paper suggests, what actually followed, and what remains open.

Content: the authors' agenda from p. 1091 in three layers, labelled as such: next steps (reverse the study by showing gesture videos and asking people to guess the command; build a vision-based recognizer and measure recognition rates), limitations left for later (let participants see all commands first; add application context; study children and other cultures), technology implications (sense above the table and beyond the display edge). Verify each on p. 1091. What followed, reusing the verified wording in progress.md: Morris, Wobbrock and Wilson, GI 2010 (pitfall 6 wording); Morris et al., "Reducing Legacy Bias in Gesture Elicitation Studies," Interactions 2014 (several gestures per command, priming with the device's capabilities, eliciting in pairs); Vatavu and Wobbrock, CHI 2015 (agreement rate, disagreement and co-agreement rates, significance test, toolkit); Villarreal-Narvaez et al., DIS 2020 (216 studies). The twist: no product adopted the set as published, yet drag, double-tap, pinch and spread became universal; Apple shipped pinch-to-zoom in 2007 and this study elicited a pinch a year later from people recruited for never having used a touch screen (Figure 4; 5 of 20 used a one-hand reverse pinch to resize, p. 1091). Open questions: mid-air, AR and VR hand tracking, wearables, voice and LLM command phrasing; whether elicitation can find anything new once everyone carries a smartphone vocabulary. Close with three concrete research questions a student could pursue, each tied to a gap in this paper (for example: record prior device exposure and analyse it as a covariate; run the reverse guessability test the authors proposed; compare elicitation with and without feedback; replicate across cultures).

## Brief 12: Cross-check of the connections file (report `reports/opus-E-connections-crosscheck.md`; run only after Brief 10 is complete)

You are an independent verifier. Read `slides/10-Connections.md`. Fetch the four course readings yourself (do not trust the URLs in the file without opening them) and read the paper. For every row in "Claims for the cross-checker" and every statement on the slide and in the notes, mark CONFIRMED, WRONG or UNVERIFIABLE with the evidence (quote and page or URL). Check the paper-side page numbers too. Check style: no em-dashes, on-slide text under about 90 words, labels "the paper says" versus "we argue," no motive attributed as fact. List every correction as an exact replacement (old text, new text). Do not edit `10-Connections.md`. Return a summary under 200 words: counts of CONFIRMED / WRONG / UNVERIFIABLE, and the corrections in priority order.

---

# Part 3. Question sets (added 2026-09-30)

Three further files hold questions for the class: one set that tests whether people read the paper, one that tests whether they listened to the presentation, and one that needs both and doubles as discussion material. Each is written by its own agent; a fourth agent verifies all three. Part 1 applies in full (ground truth, pitfalls, style, boundaries), with the file format replaced by the one below.

## 3.1 Common rules for the question writers

- **Read first:** the whole paper PDF, and all section files `01-Background-2008.md` to `12-Discussion-Bridge.md` in this folder. The section files are what the class will have heard; the paper is what they should have read. Do not use `Week 4 - Slides.md` or `Week 4 - Speaker Notes.md` (superseded).
- **Meaningful, not trivia.** A good question tests whether someone understood the contribution, the method, the evidence or the implications. Avoid questions whose only point is a number nobody would remember (no "how many women took part"). A number is fair game when it carries meaning (for example, why 48 gestures for 22 commands, or why 0.37 is not 37 percent).
- **Discriminating by design.** Each question states what it discriminates. In the paper set, every question must be answerable from the paper and **not** from the section files alone; check this against the section files. In the presentation set, every question must be answerable from what the sections say or argue and **not** from the paper alone. In the synthesis set, every question needs both.
- **Verified answers.** Every model answer carries a printed page from the paper, or a section file and slide number, or both. Quote exactly when quoting. No answer may contradict the pitfalls list in Part 1.
- **Mix of forms.** Include some quick checks (answerable in a sentence, usable as a show of hands or a one-minute written answer) and some that need a paragraph. Tag each.
- **No em-dashes.** Plain words. Short sentences.

## 3.2 File format for question sets

```
# <NN> Questions: <set title>

**Purpose:** one sentence. **Count:** n questions. **Discriminates:** what this set tells apart. **Use:** when in the session these are meant to be used (see 3.3).

## Q<NN>.<k>: <the question, as it would be asked>

- **Form:** quick check (sentence) / short answer (paragraph) / pair discussion
- **Tests:** recall / comprehension / analysis / evaluation, and what exactly
- **Model answer:** the answer with page (paper) or file and slide (presentation)
- **Why it discriminates:** what someone who skipped the paper (or the talk) would say instead
- **Wrong answers that reveal:** one or two likely wrong answers and what each shows
- **Follow-up probe:** one question to deepen a correct answer

(repeat)

## How to use in class
(a short plan: which questions open the discussion, which to hold for a lull, which to use in pairs, how answers feed the participation grading the presenters must submit)

## Verification log
| Question | Answer checked against | Status |

## Open items
```

## 3.3 Where the questions fit

The discussion half of the session runs 30 minutes: a 5-minute live elicitation warm-up, 5 minutes in pairs, 20 minutes as a group (see `12-Discussion-Bridge.md`, slide 12.3). The presenters also grade classmates' commentaries and participation afterwards. Questions may serve as warm-up checks, as pair prompts, as probes when discussion stalls, or as a quiet written check. Say in "How to use" which is which.

## Brief 13: Did they read the paper? (file `13-Questions-Paper.md`, 12 questions)

Twelve questions answerable from the paper and not from the presentation. Spread: four quick checks, five short answers, three analysis questions. Cover at least: the three design choices and the paper's reasons (p. 1083, 1085); what the two rating scales were and what each predicted (p. 1088); why insert and maximize have no dedicated gesture and why close and task switch have none (p. 1087); what "A Land Beyond the Screen" and "Acting above the Table" report (p. 1090); the authors' own designer comparison and how the paper reads it (p. 1090); what the paper claims as its four contributions (p. 1084); the Figure 2 for Figure 3 slip (p. 1091); what the taxonomy's four dimensions are and what the Form split per hand means (p. 1086, 1087); the two separate widget recommendations (p. 1091); the implications for surface technology (p. 1091); what the limitations section admits (p. 1091). Every model answer carries a printed page. For each question, state why the section files do not give it away (name the section if it comes close).

## Brief 14: Did they listen to the presentation? (file `14-Questions-Presentation.md`, 10 questions)

Ten questions answerable only from what the presentation said or argued, not from the paper. Spread: four quick checks, four short answers, two analysis questions. Draw on what the sections add to the paper: the patent application on the method and its dates (03, 04); why the study must have been done before August 2008 (01, 03); the corrected reading of "unrevised behavior" and why slide 12.1 says the evidence is thinner than the influence (04, 12); the three explanations for the 72 percent that the data cannot separate (06); the culture-bound marks and the verified examples (06); the pixel density comparison and what the table could not sense (08); the thinking-versus-doing reading of the timing data and our motor explanation (07); the DIS 2020 figures on sample size and influence (09, 04); the connections to Weiser, Wobbrock and Kientz, Greenberg and Buxton, Olsen (10); the sample-size argument in 09; the "designers measured overlap, not quality" reading (05). Every model answer carries a section file and slide number, and a paper page where the presentation built on one. For each question, state why the paper alone does not give it away.

## Brief 15: Both, and worth arguing (file `15-Questions-Synthesis.md`, 8 questions)

Eight questions that need the paper and the presentation together and have more than one defensible answer. Each should set a claim of the paper against a reading from the presentation, or ask the student to weigh evidence, for example: whether the designer comparison supports the method or only shows that twenty people cover more than three; whether elicitation measures users or measures the method; whether the 2010 preference study answers the validation gap; whether Greenberg and Buxton excuse the missing validation; whether the sample of 20 is a weakness or the field's norm; whether culture-bound marks disqualify the set or only its symbolic part; whether the patent filing changes how the paper should be read; whether the "thinking versus doing" result changes what a walk-up table should optimise. For each: the question, the strongest answer on each side with pages and sections, what a weak answer looks like, two follow-up probes, and a note on which tension in `12-Discussion-Bridge.md` it feeds. Add a "How to use in class" plan that assigns these to the pair phase and the group phase.

## Brief 16: Verify the question sets (report `../reports/opus-F-questions-check.md`; run after Briefs 13 to 15 are complete)

Independent verifier. Read the paper PDF, all section files 01 to 12, and the three question files. For every question: (a) is the model answer correct, with the page or slide it cites (CONFIRMED / WRONG / UNVERIFIABLE, with evidence); (b) does it discriminate as claimed (for set 13, could a student answer it from the section files alone; for set 14, from the paper alone); (c) does it contradict any item in the Part 1 pitfalls list; (d) is it meaningful rather than trivia. Check style: no em-dashes. List every correction as an exact replacement (old text, new text). Do not edit the question files. Return a summary under 200 words with counts per file and corrections in priority order.
