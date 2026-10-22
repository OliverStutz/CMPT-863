# Check: slides-only write-down against the v2 section slides

File checked: `slides2/slides-only/Week 7 - Slides v2.md` (21 slides). Sources: `slides2/01-Hook.md` to `slides2/09-Handover.md` (12 slides), mapping in `slides2/00-Overview.md` ("Slides-only layout: 21 slides"), title slide from `slides2/_build-deck-source.py`, v1 `slides/01-Background.md` (01.2) and `slides/03-Argument-Map.md`, paper text `assets/paper-text.txt`. Checked 2026-10-21.

## Summary

- Source bullets checked: 54 (12 slides). Reproduced exactly: 43. Reworded: 11 (6 of them sanctioned by the mapping's split or its listed additions; 5 not). Missing: 0.
- Additions checked: 20. Sourced: 14 (all quotations confirmed in the paper text on the stated page; v1 lines confirmed). Unsourced against the mapping list: 6 (all traceable to a section file's Say or If asked, none factually wrong).
- Layout issues: 5 (none reveals content too early; slide 10 tags and slide 13 definition are built correctly).
- Leaks: 3 (one page citation in body text, one file-management header, label words in Text).
- Dashes: none. Count, numbering and mapping titles: agree, with two abbreviated titles in the mapping.

## Blockers

Unsourced additions. Each is traceable to the section notes but is not on a source slide and not in the mapping's list of additions. Either cut (replacement given) or keep and add to the mapping paragraph (text given at the end of this section).

1. Slide 2. Current: `- Credit: Gaver et al., "The Drift Table," CHI 2004`. Source 01.1 Visual says "Nothing else." Replacement: delete the line, or keep it and list it (verified in `slides/01-Background.md` line 112).
2. Slide 11. Current: `- Two of Rams's ten principles.` Not on 05.1 (from the 05.1 Say, "wrote ten principles ... two of the principles"). Replacement: delete, or list it.
3. Slides 11, 12, 13. Current: `- Hands if yes.` (three times). Not on 05.1 (from the 05.1 Say). Replacement: keep, and list it; it is what makes the build on slide 13 work.
4. Slide 13. Current: `- Products and principles side by side, each note pointing at a detail.` Not on 05.1 (from the 05.1 Say). Replacement: delete, or list it.
5. Slide 15. Current: `- Not a replacement for theory. Theory's role: "inspiration and annotation."` The mapping lists only "inspiration and annotation"; "Not a replacement for theory" comes from the 05.2 If asked ("not proposing ... subsume," p. 945). Replacement: `- Theory's role: "inspiration and annotation."`, or list the first sentence.
6. Slide 17. Current: `- **Open 1:** CHI 2026's Design subcommittee invites "alternatives to scientistic ways of knowing." Has reviewing moved toward Gaver?` The question is not on 06.1 (the Say says only "a scope statement, not reviewing practice"). Replacement: `- **Open 1:** CHI 2026's Design subcommittee invites "alternatives to scientistic ways of knowing."`, or list the question.

If kept, add to the mapping paragraph in `00-Overview.md` after "the definition line on slide 13 (p. 945, in the 05.1 notes);": `"Hands if yes" on slides 11 to 13, "Two of Rams's ten principles" on slide 11 and the "side by side" line on slide 13 (05.1 Say); the photo credit on slide 2 (v1 01-Background.md); "Not a replacement for theory" on slide 15 (05.2 If asked, p. 945); the reviewing question on slide 17 (06.1 Say);`

Leak:

7. Slide 13. Current: `- **This is an annotated portfolio:** related designs, plus short notes on what joins them. "neither dominates, and neither is subservient." (Gaver, p. 945)`. Page citation in body text, not a Quote line. Replacement: `- **This is an annotated portfolio:** related designs, plus short notes on what joins them.` and a new line `**Quote:** "neither dominates, and neither is subservient." (Gaver, p. 945)`.

No missing content and no wrong build order.

## Should fix

1. Slide 15, reworded source bullet (05.2 bullet 4, not sanctioned). Source: `One design is a point. A portfolio is an area: what varied, what never did.` Current: `One design is a point in design space. A portfolio marks out an area: what varied, what never did.` Either revert to the source, or update 05.2 to match (the new wording is closer to p. 944, "a point in design space ... establishes an area").
2. Slide 17, reworded source bullet (06.1 bullet 5). Source: `Generative AI makes examples cheap. What is one object worth now?` Current: `**Open 2:** generative AI makes design examples cheap. What is one object worth now?` Replacement: `- **Open 2:** generative AI makes examples cheap. What is one object worth now?`
3. Slide 9, reworded source bullet (04.1 bullet 2). Source: `No shared paradigm yet; it "must founder."` Current: `No shared paradigm yet; such research "must founder." (Kuhn)`. Acceptable (matches the Say), but not listed; revert to "it" or note it.
4. Slide 3, the 02.1 critics' bullet is replaced by three quote cards adding "actionable metrics" and "develop protocols, descriptions, and guidelines" with year tags. All verified (paper p. 937; v1 01.2), but the mapping lists only the timeline. Add "the three quote cards on slide 3 (p. 937, v1 01.2)" to the mapping paragraph. Also, `**Quote cards:**` is a fifth field no other slide has; for a uniform format, fold the three cards into Text as bullets 5 to 7.
5. Slide 4, line 2. Current: `But standards "might lead to a form of self-policing that would be overly restrictive."` The paper goes on, "overly restrictive of a form of research" (p. 937). Replacement: `But standards "might lead to a form of self-policing that would be overly restrictive ..."`
6. Slide 4, line 3. Current: `its standards may follow "inappropriate 'scientific' models."` In the paper the subject is "attempts to establish standards for research through design" (p. 938), not HCI's standards. Replacement: `And "insofar as HCI is prone to 'scientism'," standards may adopt "inappropriate 'scientific' models."`
7. Slide 19, row 2. Current: `2. Disciplinary norms "are political acts to be approached with care"`. The paper says "attempts to establish disciplinary norms of process or outcome are political acts" (p. 945). Replacement: `2. Attempts to establish disciplinary norms "are political acts to be approached with care"`
8. Slide 10. Text puts each card name and its Fix tag in one bullet, but Build reveals the tags separately. For a slides tool, split into card bullets (`**A. Baby science.**`, `**B. Messier up close.**`, `**C. Makes worlds, not maps.**`) and tag bullets (`Fix: standardise.`, `Fix: nothing.`, `Fix: let it multiply.`).
9. Slide 2. Layout adds a description and credit under the title; the source Visual says the headline question only, "Nothing else." Also the title already ends "Is this research?" and the second bullet repeats it. Decide one: drop `- **Is this research?**` from Text (the title carries it), or shorten the title to "A coffee table that drifts."
10. Label words in Text that would be pasted onto the slide: `Credit:` (slide 2), `Caption:` (slide 13), `Footer:` (slide 18), and the bracketed tags `(instance)` etc. (slide 18). Move the labels to Layout and keep only the text that appears.
11. File header (line 3) is file-management text ("paths relative to `week7/`", "Nothing here is a speaker note"). Keep it if the file is never pasted whole; otherwise cut to `Gaver, "What Should We Expect From Research Through Design?", CHI 2012. 21 slides.`
12. Bold dropped: slide 20 `[classmate]` and slide 21 `[classmate] · [classmate] · [classmate]` are bold in 08.1 and 09.1. Restore `**[classmate]**`.
13. Slides over about 60 words of Text and Quote: 3 (about 80), 4 (64), 6 (70), 7 (63), 8 (70), 9 (86), 18 (87), 20 (73); 5 and 21 at about 61. Two ideas on one slide: slide 8 (Lakatos verdict plus our pushback, which slide 20 row 1 repeats), slide 4 (Gaver's reply plus "Why you care"). Slide 8 bullet 4 could move to slide 20 since row 1 already says it.

## Nice to have

1. Mapping table titles for rows 12 and 13 are abbreviated ("... from the products alone?"); row 1 says "Title" while the slide title is the paper title. Write them out in full.
2. Slide 1 image is `assets/rams-fig1-annotated.png` cropped; the build script's title slide uses `assets/p9-09.png`. Both exist; pick one and use it in both.
3. Slide 1 Text carries `pp. 937 to 946`. It is the bibliographic range from the source title slide, not a citation; fine to keep, but it is the only page reference outside Quote lines and slide 6.
4. Slide 8 Layout says "inspires designs" ticked; 03.2 Visual says "spawned designs." Either is fine; align the section file.
5. Slide 4: `"an elephant in the room."` adds a closing period inside the quote that 02.1 lacks; the paper has "an elephant in the room...:" (p. 946). Use `"an elephant in the room"` without the period.
6. Slide 3 timeline wording differs slightly from v1 01.2 ("research through design as a named HCI method" vs "RtD as a method"; "design artifacts and theory construction" vs "artifacts and theory"; "toward" vs v1 "towards"). "toward" matches the paper's reference list; "named" is ours. All fine.
7. Slide 20 Build does not place `[classmate]`. Add "classmate credit with the row it belongs to."
8. Slide 21 Text has a nested numbered list under a bullet (as in 09.1). Most tools accept it; if not, make the three steps top-level bullets.
9. `00-Overview.md` "Placeholders still open" still uses the 12-slide numbering (Drift Table "slides 1 and 12", thesis "slide 11", Figure 1 crops "slide 7", connections "slide 10", plan "slide 12"). Add the 21-slide numbers to avoid confusion.

## Completeness table

| Source slide | Bullets | Found on (slides-only) | Status |
|---|---|---|---|
| Title (build script) | 5 | 1 (title line as heading) | exact |
| 01.1 | 3 | 2 | 3 exact; credit line added (Blocker 1) |
| 02.1 | 4 | b1: 3 (quote cards); b2: 4 (three quoted lines); b3, b4: 4 | b1 reworded (Should fix 4); b2 reworded as mapped, quotes confirmed pp. 937, 938 (Should fix 5, 6); b3 exact (period added); b4 exact |
| 02.2 | 5 | b1 to b3: 5; b4: 5 Quote line; b5: 6 | 4 exact; b5 expanded into the move map as mapped (v1 03.1 pages confirmed; "no data, no system" from 02.2 Say) |
| 03.1 | 4 | 7 | 4 exact |
| 03.2 | 4 | 8 | 4 exact |
| 04.1 | 5 | b1 to b4: 9 (cards), Fix tags: 10; b5: 10 title | b1, b3 exact; b2 "it" to "such research" (Should fix 3); b4 replaced by Law and Urry line plus iPad Quote as mapped (p. 943 confirmed); b5 now the slide 10 title |
| 05.1 | 5 | b1: titles 11 to 13; b2: 11; b3: 12; b4: 13 title; b5: 13 | b2, b3, b4 exact; b1 split into titles as mapped; b5 replaced by the definition as mapped (p. 945 confirmed; Blocker 7) |
| 05.2 | 5 | b1 to b3: 14; b4, b5: 15 | 4 exact; b4 reworded (Should fix 1) |
| 06.1 | 5 | b1 to b3: 16; b4, b5: 17 | 4 exact; b5 split and reworded (Blocker 6, Should fix 2) |
| 07.1 | 5 | 18 | 5 exact (tags appended, "Footer:" label added) |
| 08.1 | 6 | 20 | 5 exact; [classmate] bold dropped |
| 09.1 | 3 | 21 | 2 exact; [classmate] bold dropped |

## Passed

- No source bullet missing; every one sits on the slide the mapping names.
- Quotations confirmed in the paper text on the cited page: "some sympathy" (937), "self-policing" (937), "prone to 'scientism'" and "inappropriate 'scientific' models" (938), "It is difficult to call ..." (941), iPad sentence (943), "neither dominates, and neither is subservient" (945), "inspiration and annotation" (945), "moderate expectations ..." (945), "political acts to be approached with care" (945), "reflecting and valuing ..." (946), "consensus can come to look like constraint" (946), R4 elephant (946).
- v1 sources confirmed: timeline lines (slides/01-Background.md 01.2), move page ranges (slides/03-Argument-Map.md 03.1).
- "Only A supports the critics" (slide 10) is named in the mapping table and is in the 04.1 Say; Gaver's own reading at p. 938.
- Builds: slide 10 tags and bottom line stay hidden until after the vote; slide 13 definition comes on click after the figure and the hands call; all other builds follow Text order.
- Image paths all exist in `assets/` (drift-table-chi2004-fig1.png, rams-fig1-annotated.png, rams-fig1-products-only.png).
- No speaker notes, times, verification text or "we note" asides. Page citations only in Quote lines, on slide 6, slide 1's bibliographic line, and slide 13 (Blocker 7).
- Every slide has heading, Layout, Images, Build, Text in the same order (slide 3 adds a Quote cards field).
- Header "21 slides", "N of 21" numbering 1 to 21, and the 21-row mapping table agree; titles 2 to 11 and 14 to 21 match the mapping exactly.
- No em-dashes or en-dashes anywhere in the file.
