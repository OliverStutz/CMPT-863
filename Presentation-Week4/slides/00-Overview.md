# 00 Overview: Week 4 presentation, per-section files

Presentation on Wobbrock, Morris and Wilson, "User-Defined Gestures for Surface Computing," CHI 2009, pp. 1083 to 1092. CMPT 863 Human-Computer Interaction, Week 4, presented by Oliver Stutz alone. 30-minute slot for the talk, followed by a 30-minute discussion led by the presenter.

This folder replaces the two long files `Week 4 - Slides.md` and `Week 4 - Speaker Notes.md` (revision 2, kept for reference). Each section is one file holding the slide content and the speaker notes together, so a section can be edited and reviewed on its own. Every file ends with a verification log and open items. The briefs the sections were written from are in `_agent-briefs.md`.

Written 2026-09-30. Eleven sections were drafted by independent Opus agents from identical common rules, section 12 by the orchestrating session, and every file was then read and checked by the orchestrator. Section 10 (connections to the course readings) was additionally cross-checked by a second independent Opus agent; see `../reports/opus-E-connections-crosscheck.md`.

## Deck source for slide generation

Four generated files in the presentation folder, all from `_build-deck-source.py` in this folder; edit the section files and re-run `python3 slides/_build-deck-source.py` from the presentation folder.
- `../Week 4 - Deck Source.md`: sections 01 to 12 behind a title slide, with speaker notes, without verification logs and open items. 19 slides, about 12,800 words.
- `../Week 4 - Deck Source - Presentation only.md`: the same slides with every speaker note removed; on-slide text, timing and visual directions only. About 3,600 words.
- `../Week 4 - Questions Source.md`: question sets 13 to 15 in full (stems, model answers, discrimination notes, probes, class-use plans), for the presenter. 30 questions, about 12,500 words.
- `../Week 4 - Questions - Class only.md`: the 30 question stems with their form tag and nothing else, for projecting or handing out. About 1,000 words. Reviewed for stand-alone readability by an Opus agent (`../reports/opus-G-presentation-only-review.md`): 11 must-fix and 25 should-fix items applied to the section files on 2026-09-30 (dangling section references removed, labels added, ambiguous bullets reworded, slide 12.2 cut to about 100 words); the 11 optional items are listed in the report and not applied.

## Run sheet

| # | Slide | Section file | Time | Speaker |
|---|---|---|---|---|
| 0 | Title | (title slide, see below) | 0:30 | Oliver |
| 1 | 2008, where touch stood | 01-Background-2008.md | 1:15 | Oliver |
| 2 | Why this study, and what it was | 01-Background-2008.md | 1:15 | Oliver |
| 3 | The problem: gestures built for the recognizer | 02-Problem.md | 1:15 | Oliver |
| 4 | Who owned the gestures? | 03-Pressure-and-Ownership.md | 1:30 | Oliver |
| 5 | The method: a guessability study, from referent to sign | 04-Method.md | 1:40 | Oliver |
| 6 | From 1,080 gestures to a set, step by step | 04-Method.md | 1:50 | Oliver |
| 7 | The study as run | 05-How-Conducted.md | 1:15 | Oliver |
| 8 | Three things the run revealed | 05-How-Conducted.md | 1:45 | Oliver |
| 9 | Who was in the room, and who was not | 06-Sample-and-Culture.md | 1:30 | Oliver |
| 10 | What people did | 07-Findings.md | 1:10 | Oliver |
| 11 | Time tells a story | 07-Findings.md | 1:20 | Oliver |
| 12 | What the 2008 table could not see | 08-Technical-Obstacles.md | 1:30 | Oliver |
| 13 | Limitations, theirs and ours | 09-Limitations.md | 1:30 | Oliver |
| 14 | Four readings, three roots | 10-Connections.md | 1:30 | Oliver |
| 15 | What the paper asked for, and what came next | 11-Research-Directions.md | 1:30 | Oliver |
| 16 | Our reading of the paper | 12-Discussion-Bridge.md | 0:45 | Oliver |
| 17 | Tensions worth arguing about | 12-Discussion-Bridge.md | 1:15 | Oliver |
| 18 | Over to you | 12-Discussion-Bridge.md | 1:00 | Oliver |
| | **Total** | | **25:45** | about 4 minutes of buffer |

Oliver presents alone; there is no co-presenter. The course is CMPT 863 Human-Computer Interaction.

**Title slide (not in a section file):** "User-Defined Gestures for Surface Computing." Jacob O. Wobbrock (University of Washington, Information School), Meredith Ringel Morris and Andrew D. Wilson (Microsoft Research). CHI 2009, Boston, pp. 1083 to 1092. CMPT 863 Human-Computer Interaction, Week 4, **[date]**. Presented by Oliver Stutz. Visual: Figure 1 (PDF page 1, printed p. 1083).

## The arc, and how the sections hand over

1. **Background (01)** puts the room in 2008: research multi-touch since 1985, Han's demo, the iPhone, Surface in stores April 2008, the study done before the patent filing of 4 August 2008, CHI in April 2009. Then the paper's motivation and the setup in four lines. Hands over: "the problem the paper set out to solve: who gets to decide what a gesture means."
2. **Problem (02)**: gestures built from the recognizer outward; users had been asked before, nobody had built a whole surface set from their agreement. Hands over: who was fixing conventions, and whether Microsoft was under pressure.
3. **Pressure and ownership (03)**: Apple's acquisitions, patents and public threats; Microsoft's Surface and Windows 7 timeline; the patent application on this study's method; the Apple-Microsoft cross-licence. Conclusion labelled "we argue": a race to own conventions, no evidence of lawsuit avoidance. Hands over: what Microsoft chose to patent was a procedure, the most valuable part of the paper.
4. **Method (04)**, two slides: lineage and the three design choices; then group, score, assign, alias, with both worked examples from the paper and what the method cannot do. Hands over: how it was run, and how the authors' own gestures compared.
5. **How conducted (05)**: the run; then designers' inaccuracy (43.5, 60.9, 19.1 percent), one hand versus two, fingers. Hands over: who was not in the room.
6. **Sample and culture (06)**: age spread, no ethnicity, "non-technical literate American adults," the culture-bound marks in the set, the 2008 adult shaped by single-point devices and the mouse. Hands over: what they actually did.
7. **Findings (07)**, two slides: eight ranked findings with solidity tags; then thinking time versus doing time. Hands over: those seconds happened on a 2008 prototype.
8. **Technical obstacles (08)**: 43 pixels per inch, vision sensing from below, no feedback, participants told to touch yet gesturing in the air, physics engine, "ease of recognition" untested. Hands over: the study's limitations.
9. **Limitations (09)**: the authors' own, then ours, with n = 20 sized against later practice. Hands over: the course readings warn about exactly these limits.
10. **Connections (10)**: Weiser 1991, Wobbrock and Kientz 2016, Greenberg and Buxton 2008, Olsen 2007, plus the paper's own references. Hands over: the research agenda.
11. **Research directions (11)**: the authors' agenda in three layers, what followed (GI 2010, Interactions 2014, CHI 2015, DIS 2020), the pinch twist, three open questions. Hands over: our own reading.
12. **Discussion bridge (12)**: the thesis, four tensions and three smaller items, the live elicitation warm-up and the plan for the half hour.

Each file's **Transition** line was checked against the next file's opening; all eleven hand-overs name the next topic.

## Placeholders still open

- **[date]** on the title slide.
- **[classmate]** slots, filled after the commentaries are posted: on the slides of 06, 09, 12.2 and 12.3; in the speaker notes of 02, 05, 11 and 12.1. Credit by name.
- **Discussion plan**: a separate document, not yet written. Section 12.3 states the three-step plan it must detail.

## Verification status

| File | Claims in log | [UNVERIFIED] | Weaker sources, noted in the file |
|---|---|---|---|
| 01 | 30 | none | Synaptics 1995 from a search snippet (touchpad date set to 1994, PowerBook 500, verified in 06) |
| 02 | 17 | none | Furnas 1987 "under 0.20" from a secondary copy of the abstract |
| 03 | 21 | none | Jobs's "boy, have we patented it" from secondary sources; MERL as DiamondTouch's owner in notes only |
| 04 | 39 | none | EdgeWrite 2005 guessability gain checked against the abstract only, notes only |
| 05 | 36 | none | |
| 06 | 20 | 1: the findings of Mauney et al. 2010 (existence verified, findings from a secondary summary), notes only | |
| 07 | 29 | none | |
| 08 | 22 | none | |
| 09 | 18 | none | |
| 10 | 12 cross-checkable rows, 56 claims confirmed by the independent cross-check | none after corrections (Weiser pages now confirmed via Crossref) | Olsen read from a March 2007 author copy, cited by section name; Furnas from the abstract only; 4 quote-pairing errors found by the cross-check and corrected |
| 11 | 17 | none | |
| 12 | 9 | none | "users adapt within minutes" stated as common experience, not a finding |

Four claims from the brief were dropped by their writers because they could not be verified: Palm Pre pinch-to-zoom, the Elan patent numbers, the filing date of Apple's "Jobs patent," and a 1980s date for ATMs and kiosks.

## Corrections to earlier material made during this round

- The paper does not print "5 of 20 used a one-hand reverse pinch." It says "only 1 of 5 used a diagonal 'reverse pinch'" (p. 1091). Revision 2 had the stronger wording; no section file uses it.
- The paper does instruct participants to touch the table: "We instructed participants to touch the table while gesturing" (p. 1090, under Results, not Procedure). Section 08 quotes it.
- The limitation sentence reads "non-technical literate American adults" (p. 1091); the paper states nationality, not ethnicity.
- Figure 1 is on PDF page 1 (printed p. 1083).
- Touchpad date aligned to 1994 across 01 and 06.
- CHI 2015 and DIS 2020 appear in both 04 (one sentence, as influence of the method) and 11 (as what followed). Kept in both on purpose; trim 04 if time runs short.

## Question sets (files 13 to 15)

Three files of questions for the discussion half and for the participation grading, written 2026-09-30 by three Opus agents from Part 3 of `_agent-briefs.md` and verified by a fourth (`../reports/opus-F-questions-check.md`).

| File | Tests | Count | Status |
|---|---|---|---|
| 13-Questions-Paper.md | Did they read the paper? Answerable from the paper, not from the slides | 12 | verified: 12 of 12 answers confirmed; Q13.5 and Q13.6 re-aimed at their paper-only parts; Q13.12 borderline |
| 14-Questions-Presentation.md | Did they listen? Answerable from what the sections say or argue, not from the paper | 10 | verified: 10 of 10 answers confirmed; Q14.7 re-aimed at the Olsen and floor-or-ceiling points; Q14.2 and Q14.4 borderline |
| 15-Questions-Synthesis.md | Both, with more than one defensible answer; doubles as discussion material, mapped to the tensions on slide 12.2 | 8 | verified: all 8 discriminate; five side-arguments corrected (attribution, what 43.5% measures, physical share, command versus gesture complexity, Figure 3 tail) |

Each question carries its form (quick check, short answer, pair discussion), what it tests, a model answer with page or slide, why it discriminates, likely wrong answers, and a follow-up probe. Each file ends with a plan for when to use its questions in the 30-minute discussion.

**How the three sets combine into the discussion half (30 minutes plus the buffer):**
1. **Right after the talk, 2 min, written:** Q14.1 to Q14.4 on a slip or a show of hands. A quick read on who followed the talk.
2. **Warm-up, 5 min:** the live elicitation from slide 12.3, then Q13.2 and Q13.3 as a show of hands, which tie to Figure 3 and to "same as enlarge." Note any smartphone-era gestures the room produces; they feed Q15.2.
3. **Pairs, 5 min:** one card per pair from set 15 with sides pre-assigned (T1 Q15.2; T2 Q15.2 or Q15.6; T3 Q15.3 or Q15.1; T4 Q15.8). Each pair writes one piece of evidence from the paper with a page and one from the talk with a section.
4. **Group, 20 min:** set 15's order: two opposed pairs report on Q15.2; Q15.5 as a show of hands; Q15.4 to bring in the course readings; Q15.7 only if time allows, facts and interpretation kept apart; close on Q15.8. For a lull: Q14.5 restarts discussion, Q14.6 moves it to design, Q13.6 to Q13.9 probe the paper's details.
5. **Grading:** per contribution note the student, the question, and three marks: cited the paper with a page, cited the talk with a section, acknowledged the other side's strongest point. Paper-only details show reading; section-specific points show listening.

This combination is the discussion plan; a separate document is no longer needed unless Oliver wants one for submission.

## Assets to render for the deck

| Asset | PDF page | Printed page | Used in |
|---|---|---|---|
| Figure 1, participant panning | 1 | 1083 | title, 01.2, 05.1 |
| Table 1, complexity ratings | 4 | 1086 | 05.1 |
| Table 2, taxonomy | 4 | 1086 | 04.2 (mention) |
| Eq. 1 to 3, Figure 2 | 5 | 1087 | 04.2, 07.1 |
| Figure 3, agreement per referent | 6 | 1088 | 05.2 |
| Figure 4, the gesture set | 7 | 1089 | 06.1, 11.1 |

Render with `pdftoppm -r 200 -f <page> -l <page> -png "../Week 4 - User-Defined Gestures for Surface Computing.pdf" assets/p<page>` and crop.

## Working rules for edits

- Keep each file's format: On the slide, Speaker notes (Say, Transition, Watch out, Sources), Verification log, Open items.
- No em-dashes. Check with `grep -c $'\xe2\x80\x94' *.md` (the byte sequence of the em-dash, so this file itself stays clean).
- Every number on a slide carries its printed page. Outside facts carry a URL in the log.
- "The paper says" for the authors' claims, "we argue" for ours. Never a motive stated as fact.
- Spoken scripts at or under 130 words per minute of the slide's time.
