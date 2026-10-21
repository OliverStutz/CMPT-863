# 00 Overview: Week 7 presentation, per-section files

Presentation on William Gaver, "What Should We Expect From Research Through Design?", CHI 2012, Austin, session Critical Perspectives on Design, pp. 937 to 946. CMPT 863 Human-Computer Interaction, Week 7, 22 October 2026, presented by Oliver Stutz alone. 30-minute slot for the talk, followed by a 30-minute discussion led by the presenter. Draft slides due through the course form by 21 October 2026, 11:59 pm; final version before class.

Each section is one file holding the slide content and the speaker notes together, so a section can be edited and reviewed on its own. Every file ends with a verification log and open items. The briefs the sections were written from are in `_agent-briefs.md`; the plan is `../progress.md`.

Written 2026-10-19. Sections 01, 02, 04, 05, 06, 07, 08 and 10 were drafted by independent Opus agents from identical common rules; 09 by an Opus agent from the readings-enrichment report (`../reports/opus-R-readings-enrichment.md`, itself by an independent Opus agent); 03 and 11 by the orchestrating session. Every file was then read and checked by the orchestrator. Section 08 was cross-checked by a second independent Opus agent (`../reports/opus-C08-limitations-check.md`, no blockers, all should-fix items applied); 09 was cross-checked the same way (`../reports/opus-C09-connections-check.md`, one blocker and six should-fix items applied) and the whole deck for consistency (`../reports/opus-CALL-deck-check.md`, 44 cross-file quotations compared, three contradictions and one pitfall slip found and fixed, no transition or timing failures).

## Deck source for slide generation

Generated files in the presentation folder, all from `_build-deck-source.py` in this folder; edit the section files and re-run `python3 slides/_build-deck-source.py` from `week7/`.
- `../Week 7 - Deck Source.md`: sections 01 to 11 behind a title slide, with speaker notes, without verification logs and open items.
- `../Week 7 - Deck Source - Presentation only.md`: the same slides with every speaker note removed; on-slide text, timing and visual directions only.
- `../Week 7 - Discussion Guide.md` and `../Week 7 - Discussion Cards.md`: from `12-Discussion-Guide.md` once it exists.
- `../Week 7 - Questions Source.md` and `../Week 7 - Questions - Class only.md`: only if question-set files 13 to 15 are written (Oliver's decision pending).

Mechanical checks: `python3 slides/_check-sections.py` prints, per slide, the on-slide word count, the spoken word count against 130 words per minute, and the dash and [classmate] counts per file.

## Run sheet

| # | Slide | Section file | Time |
|---|---|---|---|
| 0 | Title | (title slide, see below) | 0:30 |
| 1 | Research through design, and whose studio this is | 01-Background.md | 1:30 |
| 2 | 2007 to 2012: the call for standards | 01-Background.md | 1:30 |
| 3 | The problem: standards, scientism, self-policing | 02-Problem.md | 1:15 |
| 4 | The argument on one slide | 03-Argument-Map.md | 1:30 |
| 5 | Five genres of design theory, and the Popper test | 04-Theory.md | 1:30 |
| 6 | Lakatos, and what "provisional, contingent, and aspirational" buys | 04-Theory.md | 1:30 |
| 7 | Why the field does not converge: three readings | 05-Convergence.md | 1:45 |
| 8 | Artefacts as theory nexus, theory as annotation | 06-Annotation.md | 1:30 |
| 9 | Annotated portfolios: Rams, and the converse of patterns | 06-Annotation.md | 1:30 |
| 10 | How the essay argues: method, evidence, reflexivity | 07-Method.md | 1:30 |
| 11 | Strengths | 08-Strengths-Limitations.md | 1:00 |
| 12 | Limitations, and what the essay does not do | 08-Strengths-Limitations.md | 1:30 |
| 13 | Six readings test Gaver | 09-Connections.md | 1:45 |
| 14 | What followed: Gaver's way, the critics' way, or both? | 10-What-Followed.md | 1:30 |
| 15 | Our reading of the paper | 11-Discussion-Bridge.md | 0:45 |
| 16 | Tensions worth arguing about | 11-Discussion-Bridge.md | 1:15 |
| 17 | Over to you | 11-Discussion-Bridge.md | 1:00 |
| | **Total** | | **24:15** |

About 5:45 of buffer in the 30-minute slot. Spoken scripts total about 2,900 words, each within 130 words per minute of its slide.

**Title slide (not in a section file, emitted by the build script):** "What Should We Expect From Research Through Design?" William Gaver, Interaction Research Studio, Goldsmiths, University of London. CHI 2012, Austin, Texas, session Critical Perspectives on Design, pp. 937 to 946. CMPT 863 Human-Computer Interaction, Week 7, 22 October 2026. Presented by Oliver Stutz. Visual: the title question large, with a strip cropped from Figure 1 (p. 945).

## The arc, and how the sections hand over

The rule for every slide: a claim Gaver makes, the evidence he gives, and a test of it labelled "we argue." The course brief forbids a section-by-section retelling; slide 4 puts the whole essay on one page so that slides 5 to 9 can argue with one move at a time.

1. **Background (01)**, two slides: the term (Frayling 1993; Zimmerman, Forlizzi and Evenson 2007), Gaver's definition (p. 937), the author and three studio artefacts from its other publications, the ERC grant; then the 2007 to 2011 timeline of the critics in their own words (p. 937) and reviewer R4's "elephant in the room" (p. 946). Hands over: the critics asked for standards and theory; what does Gaver think is wrong with that?
2. **Problem (02)**: standards as "self-policing" (p. 937), HCI "prone to 'scientism'" (p. 938), what Gaver concedes, why it mattered then and now. Tagline, ours: the fight is over who defines quality, not whether quality matters. Hands over: three moves, on one page.
3. **Argument map (03)**: Move 1 (pp. 939 to 941), Move 2 (pp. 941 to 943), Move 3 (pp. 943 to 945), conclusion (pp. 945 to 946). Hands over: Move 1, by whose test?
4. **Theory (04)**, two slides: the five genres, Popper, the two reasons design theory fails his test (underspecification, "implicit sometimes," generativity, p. 940); then Lakatos, fertility, Gaver's two hedges (p. 941), "theories that are sometimes right," and three tests of ours. Hands over: if theory is only sometimes right, why agree on it?
5. **Convergence (05)**: the three readings in a table with what each implies for standards; the iPad pull quote (p. 943); Gaver's verdict (p. 938); three tests of ours in the Say and a fourth in Watch out. Hands over: if the field will not converge, what should theory do?
6. **Annotation (06)**, two slides: the chain from "ultimate particular" to "theory as annotation" (pp. 943 to 944); then portfolios, Figure 1, Rams, Alexander, the extensibility claim and Gaver's own limit (pp. 944 to 945), with four tests of ours. Hands over: how the essay makes its argument.
7. **Method (07)**: method, evidence and reflexivity in three columns; "judge it as argument, not as a study." Hands over: what does it do well, where does it fall short?
8. **Strengths and limitations (08)**, two slides: five strengths with pages; five limitations as "we argue" with Gaver's best reply to each, and one point against our own case (the "most of us" list, p. 942). Hands over: the course readings sharpen each point.
9. **Connections (09)**: six connections to the earlier course readings, chosen from the 31 in report R, with the alternates listed for Oliver to swap in. Hands over: the field did not stand still after 2012.
10. **What followed (10)**: three lanes, Gaver's way, the critics' way, both, each item verified by URL; three open questions. Hands over: our reading.
11. **Discussion bridge (11)**, three slides: the thesis with three supports; five tensions and three smaller things; classmate quotes and the three-step plan for the half hour.

Every file's **Transition** line was checked against the next file's opening; all hand-overs name the next topic.

## Depth material for the presenter (added 2026-10-20)

Oliver asked that every slide carry enough for him to go deeper when asked. Each section file now has, under every slide, an **If asked** block (six to twelve likely questions from a classmate or the instructor, each answered in two to five sentences with the printed page or a URL in the log), and after the last slide a **Background for the presenter** section (what each cited work actually argues, definitions, dates, numbers, and a few facts beyond the paper). Written by eleven Opus agents from Part 4 of `_agent-briefs.md`, one per section, without changing any existing text; checked for accuracy by two independent Opus reviewers (`../reports/opus-DCHECK-A.md` for 01 to 06: 119 items, no blockers, 9 rewordings applied; `-B.md` for 07 to 11: 112 items, one blocker fixed, 12 rewordings applied). The deck source with notes carries all of it (about 38,000 words); the presentation-only copy drops it.

Roughly 145 questions and 55 background items in all. Highlights a presenter should know exist: Gaver's own map of the essay on p. 938 (03); Kuhn's actual wording on pre-paradigm schools, which makes "must founder" Gaver's stronger phrase (05); Rams's ten principles in full (06); the four papers in the CHI 2012 session and the citation counts, dated (07); the CHI 2012 Design subcommittee members, including SFU faculty (02); the CHI 2026 Design subcommittee numbers and two CHI 2024 papers on generative AI for the open questions (10); the course brief's five discussion principles mapped onto the three steps (11).

## Reference validation (2026-10-20)

Two reports. `../reports/opus-REF-paper-references.md` checks all 36 entries in Gaver's reference list and every in-text citation: 24 entries correct, 11 with errors, 1 unverifiable (Sengers 2006). The errors that matter for the talk: in-text [3] for Newton (p. 939) and "Reprinted in [3]" should be [5]; Carroll and Kellogg (p. 944) is missing from the list; "Zimmerman & Forlizzi [34]" on p. 939 should be [33], since both quoted phrases are in the 2008 paper; "irresistables" should be "irresistibles"; the [33] quotation reads "research in practice" where the source has "research and practice"; reference 25 is Design and Culture 2(3), not 2(1). The rest are spellings (Fallman, Houghton, Wensveen), a missing author (Silverstein), a place (New York), a subtitle ("towards") and Rittel and Webber's title. `../reports/opus-REF-deck-references.md` checks every outside reference our own files cite: 101 works, 159 URLs fetched; 97 correct, one wrong (the SonicFinder issue number, now 4(1) per Crossref, fixed in 01), three partly unverifiable (a repository series label, one venue, one secondary quotation, each marked in its file), and three places where our files repeated a slip of the paper's without flagging it (now flagged in 01 and 03). Inconsistencies between files (Gaver 2014 second-hand in 07 and 08 but quoted in 10, the Kuhn Postscript year, the RTD 2013 venue, two citation-count figures) are resolved. None of this is on a slide; the "watch out" notes carry it as light asides.

## Placeholders still open

- **[classmate]** slots, filled after the commentaries are posted, credited by name: in 02 (notes), 03 (notes), 04 (both slides, notes), 05 (Say), 06 (06.2 Say), 07 (Say and notes), 08 (both slides), 09 (Watch out), 11 (11.2 and 11.3). Section 01 has none by design.
- **Thesis** on slide 11.1: a candidate for Oliver to accept, sharpen or replace (see `11-Discussion-Bridge.md`, open items).
- **Connections** on slide 13: a first selection; the alternates are listed in `09-Connections.md` and in full in report R.
- **Discussion guide** (`12-Discussion-Guide.md`): not yet written; the plan is in `../progress.md`. Slide 17's three-step plan must match it.
- **Question sets 13 to 15**: Oliver's decision pending.

## Verification status

| File | Claims in log | [UNVERIFIED] | Weaker sources, noted in the file |
|---|---|---|---|
| 01 | 24 | none | Frayling's three labels from two secondary sources that differ on "into" versus "in"; Bødker 2006 from its abstract; CORDIS does not name the principal investigator, so the slide names only the grant |
| 02 | 16 | none | |
| 03 | 14 | none | |
| 04 | 28 | none | "Degenerating programme" from the Stanford Encyclopedia of Philosophy, labelled as ours, not Gaver's |
| 05 | 20 | none | |
| 06 | 28 | none | Alexander's pattern photographs from a secondary source (Fincher 2012); Figure 1 contents read off the render |
| 07 | 16 | none | Drift Table from the Goldsmiths repository, named only in the Say as "the studio's other publications" |
| 08 | 42 | none | Cross 1982's content from a secondary source; independently cross-checked (report C08): 31 quotations and 6 outside references confirmed |
| 09 | 28 | none | Written from report R; every quotation re-checked against the reading itself, except Breazeal (scanned PDF, pages from report R, used only in the alternates); Olsen cited by section name from the March 2007 pre-print; independent cross-check in report C09 |
| 10 | 23 | none | ACM Digital Library blocked automated fetches; metadata from Crossref, abstracts from OpenAlex and Semantic Scholar; Ways of Knowing in HCI pages from a book PDF |
| 11 | 22 | none | |

## Corrections made during this round

- The plan and the quotation bank said Figure 1 shows eight products; the render shows seven photographs (two show one product in three colours), and the per-product notes speak of Rams in the third person, so they are the sketch's words, not Rams's. Corrected in the briefs, the plan and section 11; section 06 states both facts.
- Pitfall 10 overstated the reference-list slip: the "c.f. [3]" on p. 943 correctly points to Collins 1994. The slip is limited to p. 939 (Newton), p. 944 (Carroll and Kellogg, missing from the list) and the "Reprinted in [3]" entries on p. 946. Corrected in the briefs, the plan and section 05 (found by the C08 reviewer).
- Brief 07 said Gaver's examples come "largely from the author's own circle"; the paper supports only "two are Gaver's own," since Zimmerman [36] co-wrote the critics' papers. Section 07 says the narrower thing; sections 08 and 11 do not use "own circle."
- Brief 08 listed Höök and Löwgren 2012 as a missing voice; it appeared in October 2012, after CHI, so it is a parallel answer, placed in section 10. Koskinen et al. 2011 appeared around the CHI 2012 deadline, so it is "contemporaneous," off the slide.
- Brief 08's first strength said Gaver names scientism "where the field had not"; he cites Sengers 2006 on it (p. 941), so the slide says only that he names it.
- Section 04 found that Gaver does name Lakatos's criterion, "the relative productivity of competing programmes" (p. 941); the limitation in 08 and the thesis support in 11 are phrased narrowly: the criterion is named, who applies it in design is not.
- Section 05 found the p. 941 sentence in which Gaver accepts "clearly articulating ... the standards one might want to advocate"; the thesis on slide 11.1 is built on it.

## Assets to render for the deck

| Asset | PDF page | Printed page | Used in |
|---|---|---|---|
| Figure 1, the Rams annotated-portfolio sketch, with caption | 9 | 945 | title strip, 06.2 full width |

Rendered at 200 dpi to `../assets/p9-09.png`; crop to the figure and caption. No other figures or tables in the paper. The plain-text extraction of the paper is `../assets/paper-text.txt`.

## Working rules for edits

- Keep each file's format: On the slide, Speaker notes (Say, Transition, Watch out, Sources), Verification log, Open items.
- No em-dashes or en-dashes used as dashes. Check with `python3 slides/_check-sections.py`.
- Every quotation carries its printed page. Outside facts carry a URL in the log.
- "Gaver says" for his claims, "Gaver quotes" for what he quotes, "we argue" for ours. Never a motive stated as fact. Keep his hedges.
- Spoken scripts at or under 130 words per minute of the slide's time; on-slide text about 90 words.
