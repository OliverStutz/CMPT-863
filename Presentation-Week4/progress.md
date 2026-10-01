# Progress: CMPT 863 Week 4 presentation

Last updated 2026-09-30. Presentation date: not recorded yet. Oliver presents alone (no co-presenter, confirmed 2026-09-30). Course: CMPT 863 Human-Computer Interaction. Format: 30 minutes presentation, 30 minutes discussion.

## Where we stand

The presentation now lives in `slides/`, one Markdown file per section, each holding slide content, speaker notes, a verification log and open items. `slides/00-Overview.md` has the run sheet (18 slides plus title, 25:45, about 4 minutes of buffer), the arc with every hand-over, the placeholders, the verification status per file and the asset list. The briefs the sections were written from are in `slides/_agent-briefs.md`. The two long files `Week 4 - Slides.md` and `Week 4 - Speaker Notes.md` (revision 2) are superseded and kept for reference. [[Week 4 - Presentation Notes]] remains the background document; its section 6 is now answered by `slides/10-Connections.md`.

Section order, decided by Oliver on 2026-09-30: background 2008 (devices, motivation, setup), problem, pressure and ownership (patents, Microsoft's position), the method in depth (the most valuable part), how the research was conducted (designers' inaccuracy, one vs two hands, fingers), sample and culture (age, ethnicity, culture-bound marks, the tainted 2008 adult), findings (with the timing results), technical obstacles, limitations, connections to the course readings, research directions, discussion bridge.

**Next step:** apply the cross-check corrections to `slides/10-Connections.md` if any, then fill the title-slide placeholders, read the commentaries for the [classmate] slots, write the discussion plan, render the figures and build the deck.

## The paper in short

Wobbrock, Morris and Wilson, "User-Defined Gestures for Surface Computing," CHI 2009, pp. 1083–1092. Twenty non-technical adults saw the effect of 27 commands on a Microsoft Surface table and performed the gesture they thought would cause it, one-handed and two-handed, with no recognizer and thinking aloud. Output: a taxonomy, an agreement score per command, and a 48-gesture set for 22 commands. Headline findings: finger count rarely matters, one hand is preferred, 72% of gestures were mouse-like, agreement falls with command complexity, and the three authors' own designs matched under half of what users produced.

## Files

| File | What it is | Status |
|---|---|---|
| Week 4 - User-Defined Gestures for Surface Computing.pdf | The paper, 10 pages | Input |
| [[Week 4 - Presentation Notes]] | Background + seven course questions, about 3,900 words | Current. Section 6 empty |
| reports/opus-A-problem-and-contribution.md | Independent Opus report on questions 1 and 2 | Input. One claim in it is wrong, see verified facts |
| reports/opus-B-problem-and-contribution.md | Second independent Opus report on questions 1 and 2 | Input |
| [[Week 4 - Slides]] | Slide-by-slide content, run sheet with timing and speaker split, asset list, placeholders | Current, revision 2 after the Opus reviews. 22 slides. Slide 18 and slide 22 commentaries open |
| [[Week 4 - Speaker Notes]] | What to say per slide, transitions, pitfalls, page sources | Current, revision 2. About 3,100 spoken words |
| reports/opus-C-slides-review.md | Independent Opus review of the slides against the PDF, 7 blockers / 14 should-fix / 11 nice-to-have | Input, 2026-09-30 |
| reports/opus-D-slides-review.md | Second independent Opus review, 8 blockers / 19 should-fix / 9 nice-to-have | Input, 2026-09-30 |
| reports/slides-review-cross-examination.md | Where C and D agree (2 of 2), disagree (1 of 2), and the fix list in priority order | Current, 2026-09-30 |
| slides/00-Overview.md | Run sheet, arc and hand-overs, placeholders, verification status, assets | Current, 2026-09-30 |
| slides/01 to 12 | One file per section: slide content, speaker notes, verification log, open items | Current, 2026-09-30. Written by eleven Opus agents plus the orchestrator, each read and checked by the orchestrator |
| slides/_agent-briefs.md | Common rules, 20-item pitfalls list, file format, one brief per section | Reference |
| reports/opus-E-connections-crosscheck.md | Independent Opus cross-check of slides/10-Connections.md | See "What was done," step 8 |
| Week 4 - Deck Source.md | Sections 01 to 12 concatenated behind a title slide, logs and open items stripped; the input for building the deck. Regenerate with `python3 slides/_build-deck-source.py` | Derived, 2026-09-30 |
| Week 4 - Deck Source - Presentation only.md | Same slides without speaker notes; on-slide text and visual directions only. Same script | Derived, 2026-09-30 |
| Week 4 - Questions Source.md | Question sets 13 to 15 in full, presenter version. Same script | Derived, 2026-09-30 |
| Week 4 - Questions - Class only.md | The 30 question stems only, class version. Same script | Derived, 2026-09-30 |
| progress.md | This file | |
| Claude memory: week4-wobbrock-2009-figures | Numbers read off Figures 2–4 and Tables 1–2 | Reference for future sessions |

## What was done on 2026-09-29

1. **Paper read in full**, text and rendered figures. Figure and table values recorded in memory. Pages 2, 3, 8, 9, 10 are text only.
2. **Two Opus agents** answered questions 1 and 2 independently from the same brief. Both ranked the method above the gesture set, both picked the designers-versus-users comparison as the one finding to remember, and both read the "first to employ users" claim as narrower than it sounds. Reports saved in `reports/`.
3. **Checked a disagreement between the agents.** The paper prints the range of A as [|Pr|^-1, 1], which is correct. Agent A's claim of a misprint came from text extraction flattening the exponent. The Figure 2 versus Figure 3 mislabel on p. 1091 is real.
4. **Notes file written**: background first, then the seven questions, section 6 left empty.
5. **Checked the notes against the course goal** (understand the contribution, evaluate critically, engage with broader implications). Understanding: strong. Critical evaluation: concentrated in sections 2 and 5, section 4 is a flat list. Broader implications: thin. No discussion-lead material.
6. **Legacy-bias highlight added** from Oliver's insight that participants in 2008 already had touch experience. The claim that cars had gesture control in 2008 was checked and corrected: production in-car gesture control began with the BMW 7 Series in 2015.
7. **Pre-2007 gesture device history added** to section 5, with a table of mass-market devices, niche multi-touch products and the research lineage.
8. **Ten further key points identified** for helping the class understand the paper. Not yet in the notes. Listed below.

## What was done on 2026-09-30

1. **Course brief reviewed** (30 min presentation + 30 min discussion; rubric: presentation quality 7, integration of commentaries 4, discussion 7). Oliver's requested shape: a rundown of the paper first, then the key questions, then the interesting parts that feed the discussion.
2. **Slides file written**, 23 slides in three parts, 28 minutes with 2 minutes buffer. Suggested speaker split A/B with slides 18 and 19 fixed to Oliver. Placeholders: date, co-presenter, slide 19, slide 23 commentaries, figure renders.
3. **Speaker notes written**, one block per slide. Three "never say" rules at the top: set not validated, agreement score not invented here, A is not a percentage.
4. **Three follow-up citations verified** by web search (see verified facts). The GI 2010 finding was reworded: the notes had "users preferred user-authored gestures"; the paper's finding is that blind judges preferred gestures authored by larger groups, partly because the researchers' gestures were more complex. Notes section 7 corrected accordingly.
5. **Live elicitation proposed** as the discussion warm-up (slide 23): show two commands, everyone gestures on the desk, tally groups, compute A in the room. Discussion plan itself not yet written.
6. **Two independent Opus reviews of the slides** (same brief, each checked every number against the PDF, neither saw the other). Reports in `reports/opus-C-slides-review.md` and `reports/opus-D-slides-review.md`. Cross-examination with vote counts in `reports/slides-review-cross-examination.md`. Both agree: numbers mostly right, Part B is real critique, but slide 18 and the slide 21 thesis attack a claim the paper does not make ("unrevised behavior" means not revised by recognizer feedback, p. 1085, not "free of prior experience"), slide 9's Figure 3 readings and widget list are wrong, and the deck runs 3–5 minutes over. Corrected in step 7.
7. **Review fixes applied** to both files (revision 2). Old slides 11 and 12 merged; slide 18 rewritten as "Whose habits did the study capture?" with the verbatim screening quotes, three candidate explanations for the 72% and the pinch as an open question; slide 9 corrected from a 220 dpi re-render of Figure 3 (tail .21 to .11; two-hand peaks insert, maximize, enlarge; dips pan, close, task switch; two separate widget recommendations with the paper's list); slide 16 cites Eq. 3 for the 0.37; grouping and re-binding on slides 7 and 16 hedged as inference; minimize fix; Figure 4 idiom labels fixed (jump, hold and tap, menu "pull out" from a held object, confirmed on the rendered figure); "circular" replaced by self-rated with no independent judges; N = 1,074 non-independence and the r = .71 fairness caveat added; thesis narrowed; all wording items from the cross-examination applied. Now 22 slides, 25:45 planned, about 3,100 spoken words, every block within its slot at 100–137 words per minute. [classmate] slots on slides 14, 17, 21.

8. **Restructure into per-section files** at Oliver's request. Briefs written (`slides/_agent-briefs.md`), eleven Opus writer agents run in parallel, section 12 written by the orchestrator from revision 2, every file read and checked, consistency edits made (touchpad date 1994 in 01 and 06; "eleven people made the same gesture, a tap according to Figure 4" in 04; Figure 1 on PDF page 1 in 08; thesis script trimmed in 12). Mechanical checks: zero em-dashes in all 13 files; every spoken script between 107 and 135 words per minute; 2,975 spoken words for 25:15 of section slides. Section 10 cross-checked by a second Opus agent (report in `reports/opus-E-connections-crosscheck.md`). New facts established this round are in the verified-facts table.

9. **Question sets written** at Oliver's request, as `slides/13-Questions-Paper.md` (12 questions, answerable from the paper only), `slides/14-Questions-Presentation.md` (10, answerable from the talk only) and `slides/15-Questions-Synthesis.md` (8, needing both, mapped to the tensions on slide 12.2, with pair cards, a timed group order and a participation rubric). Briefs 13 to 16 appended to `slides/_agent-briefs.md`; three Opus writers in parallel, then an Opus verifier (`reports/opus-F-questions-check.md`). The combined use of the three sets, written into `slides/00-Overview.md`, is the discussion plan.
10. **Deck source built**: `Week 4 - Deck Source.md`, sections 01 to 12 behind a title slide, regenerated by `slides/_build-deck-source.py`.

## Decisions made by Oliver

| Decision | Date |
|---|---|
| Background section comes first and holds only the minimum needed to follow the paper | 2026-09-29 |
| Notes are structured by the seven course questions | 2026-09-29 |
| Section 6 (connections to course topics) is Oliver's to write | 2026-09-29 |
| Highlight how hard unbiased elicitation is in an age of touch displays | 2026-09-29 |
| Retain the pre-2007 device history in the notes | 2026-09-29 |
| Two independent Opus agents with the same brief for questions 1 and 2 | 2026-09-29 |
| Presentation order: rundown of the paper, then the key questions, then interesting parts usable as discussion topics | 2026-09-30 |
| Slide content and speaker text in two separate files (superseded the same day) | 2026-09-30 |
| One Markdown file per section, slide content and notes together, so each question can be worked on alone | 2026-09-30 |
| Section order: background, problem, pressure and ownership, method, how conducted, sample and culture, findings, technical obstacles, limitations, connections, research directions, discussion bridge | 2026-09-30 |
| No em-dashes anywhere; every fact verified against the paper or a URL; unverifiable claims dropped or tagged | 2026-09-30 |
| Connections section written by an Opus agent and cross-checked by a second, independent Opus agent | 2026-09-30 |

## Findings not yet in the notes

### Ten further points for understanding the paper

Points 1 to 5 explain how the paper works and belong in sections 2 to 4. Points 6 to 10 open discussion and belong in a discussion-lead section.

1. **Agreement, coverage and guessability are three different numbers.** A rewards one dominant group (select single: 11 of 20 tapped, A about 0.37, 9 people did something else). Coverage of 57% counts proposals matched by any gesture in the final set, aliases included. Guessability, the chance a new user guesses right, is never measured.
2. **The set is conflict-free only because the target is part of the gesture.** Spread fingers means enlarge on an object, zoom on the background, open on a folder. Drag off-screen means delete for an object, reject for a dialog. Conflicts were resolved by letting the larger group win, then the loser was rescued by binding the same gesture to a different target. A design move, not a finding.
3. **The grouping decision feeds back into the score.** Finger count was ignored when deciding which gestures are "identical," on the grounds that users did not care. That merges groups and raises agreement. No inter-rater reliability is reported for the grouping or the taxonomy coding.
4. **Two-hand gestures were forced.** Every participant had to produce one for every command although one hand was preferred for 25 of 27. The hold-and-tap idiom came from this, and two-hand gestures are 35% of the final set. Morris et al. 2014 later formalized forced production as a bias-reduction technique.
5. **Simplified mental models drive the reuse principle.** "Zoom in, that's the same as enlarge." Maximize became enlarge, minimize became shrink. The set's one-gesture-per-family rule came from users collapsing categories. This is why Figure 3's order differs from Table 1's complexity ranking.
6. **Symbolic gestures import writing conventions.** Check mark, X, question mark, caret are Western marks. The most abstract commands got the most culture-bound gestures.
7. **Single-user elicitation on a device built for groups.** Morris wrote the cooperative gestures paper the study cites. Minimize is "drag to the bottom of the surface as defined by the user's seating position," which differs for each person at a table.
8. **The physics-engine tie-in.** The 44% physical gestures are read as support for Wilson's own physics-based Surface work (ref. 32, UIST 2008). Same lab, previous year.
9. **Elicitation as floor, not ceiling.** 19.1% of each author's gestures were never tried by any participant. Does elicitation find the most guessable gesture or the best one?
10. **Longer gestures were rated easier.** About 3.4 s for gestures rated easy against 2.0 s for those rated hard. A memorable case of self-report behaving unexpectedly.

### Discussion-lead candidates from the Opus reports

Present in the notes in scattered form (sections 5 and 7), not yet framed as prompts.

1. Should the most popular gesture win? 72% mouse-like, "I'm a child of the mouse." Does consensus from novices lock in legacy habits?
2. Did the commands shape the answers? The 27 referents came from desktop systems (cut, paste, minimize, task switch).
3. Is the validation circular? "Popularity identifies better gestures" rests on participants rating gestures they had just invented.
4. Is unrevised behaviour the right design target when real users adapt to feedback within minutes?

### Gaps from the course-goal check

- No thesis line. Suggested: the lasting contribution is the method, and the paper's own evidence for the method is weaker than its reputation.
- Section 4 lists about twenty findings without ranking or saying which are robust.
- Section 3 could name the design choices to scrutinize (authors coded their own data; grouping drives everything downstream). The screening criterion paragraph is already there.
- Broader implications to add: the method spread to hundreds of studies but the set went nowhere, and pinch-to-zoom became universal because Apple shipped it, not because anyone elicited it. The legacy-bias finding became a limitation of the method itself (Morris et al. 2014).

## Verified facts

| Claim | Result | Source |
|---|---|---|
| Range of A printed as [|Pr|^-1, 1] | Correct in the paper; the "misprint" in report A is an extraction artifact | Page 1087 zoomed at 300 dpi |
| "Figure 2 indicates that agreement is low after the first seven referents" (p. 1091) | Should be Figure 3 | Paper |
| First production car with gesture control | BMW 7 Series, 2015 | bmwblog.com, digitaltrends.com |
| Samsung SCH-S310 air-writing phone | Announced January 2005, Korea only | photographyblog.com, gearlive.com |
| FingerWorks multi-touch gesture pads, acquired by Apple | Founded 1998, sold to Apple 2005 | dealroom.co, appleinsider.com |
| Morris et al., "Reducing Legacy Bias in Gesture Elicitation Studies" | Interactions, May–June 2014, exists as cited | interactions.acm.org |
| Morris, Wobbrock, Wilson, "Understanding Users' Preferences for Surface Gestures," GI 2010, pp. 261–268 | Exists. 22 participants blind to authorship judged 81 gestures; preferred gestures authored by larger groups, partly because researchers' gestures were more complex. The earlier "preferred user-authored" wording was too strong | graphicsinterface.org, faculty.washington.edu/wobbrock, microsoft.com/research, 2026-09-30 |
| Vatavu and Wobbrock, "Formalizing Agreement Analysis for Elicitation Studies," CHI 2015 | Exists. Agreement, disagreement and coagreement rates, significance test, toolkit, Honorable Mention | faculty.washington.edu/wobbrock/pubs/chi-15.01.pdf, 2026-09-30 |
| Villarreal-Narvaez, Vanderdonckt, Vatavu, Wobbrock, "A Systematic Review of Gesture Elicitation Studies: What Can We Learn from 216 Studies?", DIS 2020 | Exists. Supports "hundreds of elicitation studies" | research.dial.uclouvain.be, faculty.washington.edu/wobbrock/pubs/dis-20.01.pdf, 2026-09-30 |
| Microsoft patented the method | US patent application 2010/0031202 A1, "User-Defined Gesture Set for Surface Computing," inventors Morris, Wobbrock, Wilson, applicant Microsoft, filed 2008-08-04, published 2010-02-04, **abandoned 2014-12-09**. Claim 1 covers the elicitation procedure (show an effect to several users, capture their surface input by computer vision, pick the gesture by correlation). The paper does not mention it. Also dates the study: completed before August 2008 | patents.google.com/patent/US20100031202A1, patents.justia.com/patent/20100031202, 2026-09-30 |
| Related Microsoft Surface gesture patent | US 10,409,381 B2, "Gestures, interactions, and common ground in a surface computing environment," 13 inventors incl. Morris and Wilson, priority 2008-12-15, granted 2019. Not derived from this study | patents.google.com/patent/US10409381, 2026-09-30 |
| Study conducted to avoid lawsuits? | No evidence either way. The paper states academic motives; the patent filing shows Microsoft treated the method as IP. See Oliver's question of 2026-09-30 | |
| Participants were told to touch the table | "We instructed participants to touch the table while gesturing." Under Results ("Acting above the Table"), not in the Procedure | p. 1090 |
| Nationality of the sample | "Our participants were all non-technical literate American adults; undoubtedly, children, Eastern, or uneducated participants would behave differently." No ethnicity reported | p. 1091 |
| Reverse pinch count | "only 1 of 5 used a diagonal 'reverse pinch' to resize along the diagonal, while 4 of 5 resized in other orientations." The paper does not say "5 of 20" or "one hand"; revision 2 overstated this | p. 1091 |
| Ref. 32 | Wilson, Izadi, Hilliges, Garcia-Mendoza, Kirk, "Bringing physics to the surface," UIST 2008, pp. 67 to 76 | p. 1092 |
| Pixel density of the prototype | 1024 x 768 on 24 x 18 inches is about 43 ppi; current iPad Pro 264 ppi | p. 1085; apple.com/ipad-pro/specs |
| Later elicitation studies' sample sizes | DIS 2020 review: average 25 participants (sd 4); 20 the most common (36% of studies), "likely" following this paper. The review calls this paper the first hand-gesture elicitation study and the most cited (1,070 citations as of 2019) | faculty.washington.edu/wobbrock/pubs/dis-20.01.pdf, p. 859 |
| Vatavu and Wobbrock 2015 rescaling | Their corrected agreement rate puts this paper's one-hand 0.32 at 0.284 | faculty.washington.edu/wobbrock/pubs/chi-15.01.pdf |
| Cross-cultural gesture examples | Bulgaria: a head nod means no. Japan: a circle marks correct, a check can mark wrong. Iran: thumbs-up is an insult | URLs in slides/06-Sample-and-Culture.md |
| Device dates | PowerBook 500 trackpad May 1994; Palm Pilot 1996; iPod touch wheel 2002; Nintendo DS Nov 2004; Wii Nov 2006; New York subway touch-screen ticket machines 1999. ATMs and kiosks: no date verified, dropped | URLs in slides/01 and 06 |
| Tim Cook, January 2009 | Reported words: "we will not stand for having our IP ripped off" and "if they do, we'll go after them." He did not name the Palm Pre | URLs in slides/03 |
| Wobbrock's 2012 whitepaper | "Seven Research Contributions in HCI," p. 3, lists this paper as a methodological contribution example; the 2016 Interactions article does not name it | faculty.washington.edu/wobbrock/pubs/Wobbrock-2012.pdf |
| Dropped as unverifiable this round | Palm Pre pinch-to-zoom; Elan patent numbers; filing date of Apple's "Jobs patent"; whether the 1997 Apple-Microsoft deal itself contained the anti-cloning clause; Mauney et al. 2010 findings (existence verified, findings only from a secondary summary) | |
| Device dates in the pre-2007 table other than the three above | From general knowledge | |

## To do

- [x] Cross-check corrections applied to `slides/10-Connections.md` (2026-09-30: 56 confirmed, 4 pairing errors fixed, 5 wording items)
- [ ] Title slide: presentation date (no co-presenter; all slides are Oliver's)
- [ ] Read classmates' commentaries; fill the [classmate] slots (slides of 06, 09, 12.2, 12.3; notes of 02, 05, 11, 12.1)
- [x] Discussion plan: the combined use of question sets 13 to 15 in `slides/00-Overview.md` (2026-09-30). Write a separate submission document only if required
- [x] Corrections from `reports/opus-F-questions-check.md` applied to the question files (2026-09-30: 25 answers confirmed, 5 partly wrong side-arguments fixed, 4 stems re-aimed for discrimination)
- [x] Readability corrections from `reports/opus-G-presentation-only-review.md` applied (36 of 47; the 11 optional items are not), deck sources regenerated (2026-09-30)
- [ ] Render the figure assets listed in `slides/00-Overview.md` and build the deck from `Week 4 - Deck Source.md`
- [ ] Optional: read Mauney et al. 2010 and the Furnas 1987 abstract through SFU library access to upgrade two secondary sources
- [ ] Submit draft slides the day before class by 11:59 PM, final before class
- [ ] Commit the Week 4 files to git
