# 00 Overview: Week 7 presentation, version 3 (Oliver's outline, full notes, extractable slide text)

Presentation on William Gaver, "What Should We Expect From Research Through Design?", CHI 2012, pp. 937 to 946. CMPT 863, Week 7, 22 October 2026, Oliver Stutz alone. 30-minute talk, 30-minute discussion.

Version 3 follows the outline Oliver gave on 2026-10-21 after reviewing version 2: introduce the field and the object; how the mess arose historically and how it showed itself; what methods and criteria existed before the paper; the method Gaver proposes; its upsides; critical reflection; what the paper changed; its value in 2026; then connections, the room's own projects, and the discussion. Twenty-two slides including the title. Every slide carries speaker notes, Watch out, If asked and Sources. The slide text alone is extracted mechanically into `slides-only/Week 7 - Slides v3.md`.

Material is reused from version 1 (`../slides/`) and version 2 (`../slides2/`), both cross-checked; the briefs are in `_agent-briefs.md` (rules, file format, 18 pitfalls, per-slide must-haves and sources). Written by three Opus writers in parallel from the briefs, then cross-checked.

## Run sheet

| # | Slide | File | Time |
|---|---|---|---|
| 1 | What Should We Expect From Research Through Design? (title) | (build script) | 0:30 |
| 2 | What is research through design? Here is an object to argue about. | 01-Scene.md | 1:30 |
| 3 | Design got into HCI between 1993 and 2007. | 01-Scene.md | 1:00 |
| 4 | The mess showed itself as dozens of manifestos. | 01-Scene.md | 1:00 |
| 5 | The reaction, 2007 to 2011: a call for standards. | 01-Scene.md | 1:15 |
| 6 | Research through design was being done by building things and writing theory around them. | 02-Before.md | 1:00 |
| 7 | People proposed to judge it by four criteria, then by protocols and metrics. | 02-Before.md | 1:00 |
| 8 | Gaver objects: standards risk self-policing and scientism. | 02-Before.md | 1:15 |
| 9 | Which science? Gaver runs design theory through two tests | 03-Argument.md | 1:30 |
| 10 | Why don't designers converge? Three readings, and you vote | 03-Argument.md | 2:00 |
| 11 | What Gaver proposes: the object is the fact, theory annotates it | 03-Argument.md | 1:30 |
| 12 | Could you design like Dieter Rams? A test in three steps | 03-Argument.md | 2:00 |
| 13 | How to make one, and what it is not | 03-Argument.md | 1:00 |
| 14 | Six things the annotated portfolio gets right | 04-Judging.md | 1:30 |
| 15 | Where we push back | 04-Judging.md | 1:30 |
| 16 | Our verdict, in one sentence | 04-Judging.md | 1:00 |
| 17 | What the paper changed, 2012 to 2016: three lanes | 05-Afterwards.md | 1:30 |
| 18 | What is it worth in 2026? | 05-Afterwards.md | 1:30 |
| 19 | Where this touches what we have read | 05-Afterwards.md | 1:30 |
| 20 | Back to the table. Did you move? | 05-Afterwards.md | 1:15 |
| | **Talk total** | | **26:15** |
| 21 | Five tensions. Which one first? | 06-Discussion.md | 0:45 |
| 22 | The next thirty minutes | 06-Discussion.md | 0:30 |

The slide times are caps at 130 words a minute. The scripts for slides 1 to 20 total about 3,200 words; read at 125 words a minute, with the four hands moments (slides 2, 10, 12, 20) and about forty clicks, the talk runs about 29 minutes. If rehearsal runs long, trim in this order: slide 10 (the three readings), slide 12 (the Rams steps), slide 2 (the definition), slide 17 (the lanes); each file's Open items name the sentences to drop.

## Cross-checks (2026-10-21)

Two independent Opus reviewers, run in parallel after the three writers finished. The references check (`../reports/opus-V3-references-check.md`) tested 416 quotation instances (415 exact, one missing inner quotes, fixed), 416 page references (two wrong, one slip on slide 3, fixed), 73 outside facts (72 confirmed; Cross 1982 was then taken off slide 3), and the extraction (exact). The flow check (`../reports/opus-V3-flow-check.md`) confirmed the deck follows Oliver's outline with every must-have present, and found four blockers, all fixed: slide 3 never said why design entering HCI led to the manifestos (a marked "our reading" now does); the brief's "how was the research conducted" had no home (the slide 8 transition now states it); slide 12's script had four clicks against three in the build; slide 20's counting instructions were only in the notes. Its should-fix list (duplicate transitions, repeated content, read-aloud problems, unmarked opinion, about 250 words of cuts for timing) was applied the same evening.

## Files and scripts

- `_agent-briefs.md`: rules, file format, pitfalls, the 22 slides with must-haves and sources, writer assignments.
- `01-Scene.md` to `06-Discussion.md`: the section files, one per part.
- `_check-sections.py`: dashes, Say words per minute, on-slide density. Run from `week7/`.
- `_build-deck-source.py`: writes `../Week 7 - Deck Source v3.md` (with notes) and `../Week 7 - Deck Source v3 - Presentation only.md`.
- `_extract-slides-only.py`: writes `slides-only/Week 7 - Slides v3.md`, the slide text alone.

## Placeholders

- `**[classmate]**` on slides 15, 16, 21 and 22, filled after the commentaries are posted, credited by name.
- Thesis on slide 16: a candidate for Oliver to accept, sharpen or replace.
- Discussion guide: slide 21's anchors and slide 22's cards are the guide in outline; the full guide with probes stays to be written if Oliver wants one beyond the notes.
