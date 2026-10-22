# Version 3 deck: references, quotations and pitfalls check

Checked 2026-10-21 by an Opus reviewer. Scope: the 22 slides in `slides3/` (title from `_build-deck-source.py`, slides 2 to 22 from `01-Scene.md` to `06-Discussion.md`), against `assets/paper-text.txt` (page rule from the brief) and the logs in `slides/` and `slides2/`. NEW outside facts were re-fetched (Crossref and OpenAlex for Fallman 2003, Crossref and Wikipedia for Cross 1982, the CHI 2004 Drift Table PDF for "about a terabyte"). Nothing in `reports/` was opened. No deck file was edited.

## Summary

- Paper quotations checked: 416 instances (every string in quotation marks traced to the paper, in all fields, repeats counted). Exact: 415. Inexact: 1 (slide 8 Say drops Gaver's inner quotes around 'scientific').
- Pages: 414 right, 2 wrong (one error on slide 3, "a domain of design ... c.f. [7]" given as p. 943, found p. 944; it recurs in the Sources line and the log).
- Misattribution (P5 type): none. Every quotation Gaver takes from others is marked as such (critics p. 937, Lakatos and Popper pp. 939 to 940, Law and Urry p. 943, Stolterman p. 944, Carroll and Kellogg p. 944, Sengers p. 938). The p. 941 fertility sentence is correctly Gaver's own. Note: the ludic design line (p. 938) is Gaver quoting his own 2009 chapter [13]; attributing it to Gaver is correct.
- Outside facts checked: 73. Confirmed: 72. Not confirmed: 1 (title slide Say, Gaver "runs" a studio at Goldsmiths; false since late 2021).
- Pitfall violations: 4 (P1 slide 6; P13 slide 16; P15 slide 10; P18 slide 8). No em-dashes or en-dashes in any slides3 file or the slides-only file.
- Extraction mismatches: 0 (all 22 titles and every On the slide bullet reproduced exactly; all 12 Quote lines carry "(Gaver, p. N)"; no notes, logs or transitions leaked; the only page numbers are the Quote attributions, the title slide's "pp. 937 to 946", and "(p. 945)" inside the slide 1 Layout line).

## Blockers

1. **Wrong page.** `01-Scene.md`, slide 3, If asked.
   Current: `and portfolios making clear "a domain of design, its relevant dimensions" (p. 943, "c.f. [7]")`
   Replace: `and portfolios making clear "a domain of design, its relevant dimensions" (p. 944, "c.f. [7]")`
   Same file, slide 3 Sources. Current: `**Sources:** paper p. 938, p. 940, p. 943 ([7])` Replace: `**Sources:** paper p. 938, p. 940, p. 944 ([7])`
   Same file, log row. Current: `and p. 943 ("c.f. [7]", domain of design, dimensions) | paper p. 940, p. 943 | OK |` Replace: `and p. 944 ("c.f. [7]", domain of design, dimensions) | paper p. 940, p. 944 | OK |`
   (The sentence sits at paper-text lines 900 to 904, after the "943" marker: page 944. The Say's "Gaver cites him for exactly that move" is unaffected.)
2. **False outside fact.** `_build-deck-source.py`, TITLE block, slide 1 Say.
   Current: `William Gaver runs a design research studio at Goldsmiths in London, and in 2012 he answered colleagues who wanted research through design to adopt standards.`
   Replace: `In 2012 William Gaver was at the Interaction Research Studio at Goldsmiths in London, and he answered colleagues who wanted research through design to adopt standards.`
   (`slides/01-Background.md` log: the studio moved to Northumbria University in late 2021.)
3. **Inexact quotation.** `02-Before.md`, slide 8, Say.
   Current: `standards "may adopt, or be interpreted in terms of, inappropriate scientific models."`
   Replace: `standards "may adopt, or be interpreted in terms of, inappropriate 'scientific' models."`
   (p. 938 has inner single quotes. Spoken delivery is unchanged; the Watch out already says to say the word plainly.)

## Should fix

1. **P1, slide 6 Say** (`02-Before.md`). The 2004 Drift Table is credited to "the studio".
   Current: `From the studio's other publications: the Drift Table was lived with for several weeks, and the Prayer Companion stayed with a convent of nuns for ten months.`
   Replace: `From other publications by Gaver and colleagues: the Drift Table was lived with for several weeks, and the Prayer Companion stayed with a convent of nuns for ten months.`
   Same slide, Watch out. Current: `Say "from the studio's other publications" (and for the 2004 table, "Gaver and colleagues", P1).` Replace: `Say "from other publications by Gaver and colleagues" (P1).`
2. **P18, slide 8 Say** (`02-Before.md`). Drops "in part" and overstates agreement.
   Current: `One reviewer, R4, saw it the same way, and said the essay addresses "an elephant in the room": the move towards scientism in HCI.`
   Replace: `One reviewer, R4, wrote that the essay addresses, in part, "an elephant in the room": the move towards scientism in HCI.`
3. **P15, slide 10 Say** (`03-Argument.md`). Drops "from this perspective" and changes the subject: on p. 943 it is our criteria and "even the assumption that shared standards are necessary, possible or desirable" that are the repressive acts.
   Current: `So shared standards are, in his phrase, potentially repressive acts of ontological politics.`
   Replace: `So, from this perspective, he says our criteria, and even the assumption that shared standards are necessary, are potentially repressive acts of ontological politics.`
4. **Presenter reading stated as Gaver's, slide 10 Say** (`03-Argument.md`). Gaver does not pick a reading (p. 938).
   Current: `Gaver leans to B and C.` Replace: `We read Gaver as leaning to B and C.`
5. **P13, slide 16 On the slide** (`04-Judging.md` and the slides-only file). The criterion exists; the judge is unnamed.
   Current: `- 2. "Productive research programmes" has no judge.` Replace: `- 2. "Productive research programmes" names no judge.`
6. **P14 context, slide 15 Say** (`04-Judging.md`). The p. 941 sentence opens "Calls for disciplinary reform ... appeal for"; Gaver is summing up the critics' case, so "himself" overstates.
   Current: `Gaver himself says integration would make assessment fairer and more rigorous, then never returns to it.`
   Replace: `Summing up the critics' case, Gaver writes that greater integration would make assessment fairer and more rigorous, then never returns to it.`
7. **Our reading unmarked, slide 7 Say** (`02-Before.md`; the Watch out marks it ours, the Say does not).
   Current: `Notice that the 2007 lenses already drop validity, so these critics were not asking for laboratory science.`
   Replace: `Notice that the 2007 lenses already drop validity; we read that as a sign these critics were not asking for laboratory science.`
8. **Count, slide 1 Say** (`_build-deck-source.py`). Hands are asked on slides 2, 10, 12 and 20.
   Current: `I will need your hands three times.` Replace: `I will need your hands four times.`
9. **Overstatement, slide 17 Say** (`05-Afterwards.md`). Pierce's abstract "offers the notions of thingly publications".
   Current: `Pierce asked for "thingly publications."` Replace: `Pierce proposed "thingly publications."`

## Nice to have

1. Slide 3 Watch out, log and Open items (`01-Scene.md`) say the Cross 1982 original "was not opened". Version 1 did read it: `slides/08-Strengths-Limitations.md` log ("Cross 1982: three cultures ...", Cross pp. 221 to 222, https://Readings.design/PDF/cross_1982_designerlywaysofknowing.pdf). Cite that row and drop "If challenged, drop the dot".
2. Slide 9 Say (`03-Argument.md`): current `Gaver's first move is a judo move.` Replace: `Gaver's first move is what we call a judo move.` (Label is flagged as ours only in Watch out.)
3. Slide 12 If asked (`03-Argument.md`): current `is environmentally friendly` Replace: `is environmentally-friendly` (Vitsoe's spelling in the v1 log).
4. Slide 20 If asked (`05-Afterwards.md`): current `the paper's claim that what counts as research is contested (p. 941)` Replace: `the paper's point that accounts of science are "unsettled and controversial" (p. 941)` (closer to the page; Gaver writes "how unsettled and controversial accounts of science are").
5. Slides 8 and 14 Say speak Gaver's inner-quoted phrases without the inner quotes ("extensibility and verifiability", "merely"); acceptable for speech, no change needed.
6. Slide 1 Layout line in the slides-only file carries "(p. 945)"; it is layout text, not slide text. Leave or drop.

## Quotation table

Every string in quotation marks that comes from the paper, in file order. "Page given" is the page the slide or notes give for it (Sources lines give pages in list form; checked by hand); "page found" is from `assets/paper-text.txt` with the brief's page rule. Strings that also occur on another page show both. Labels, outside quotations and spoken cues are excluded and covered in the outside-fact table.

| Slide | Field | Text | Page given | Page found | Status |
|---|---|---|---|---|---|
| 2 | On slide | "brought to bear on situations chosen for their topical and theoretical potential" | 937 | 937 | exact (crosses column break) |
| 2 | On slide | "primarily, of artefacts and systems" | 937 | 937 | exact |
| 2 | If asked | "as a long-term practitioner of research through design" | 937 | 937 | exact (sentence-initial capital changed) |
| 2 | Sources | "long-term practitioner" | 937 | 937 | exact |
| 3 | If asked | "as a design discipline [7]" | 938 | 938 | exact |
| 3 | If asked | "potentially independent factors of an artefact" | 940 | 940 | exact |
| 3 | If asked | "c.f. [7]" | 940 | 940 | exact |
| 3 | If asked | "a domain of design, its relevant dimensions" | 943 | 944 | exact; PAGE WRONG |
| 3 | If asked | "c.f. [7]" | 943 | 944 | exact; PAGE WRONG |
| 4 | On slide | "there must be dozens of different 'manifestos' for design, each seeking to establish a ..." | 941 | 941 | exact |
| 4 | On slide | "clear, efficient solutions" | 938 | 938 | exact |
| 4 | On slide | "little agreement" | 942 | 942 | exact |
| 4 | Watch out | "tend to downplay both their theoretical commitments and normative stance" | 938 | 938 | exact |
| 4 | Watch out | "There must be dozens" | 941 | 941 | exact (sentence-initial capital changed) |
| 4 | Watch out | "It is clearly possible to recognise the current state of research through design in thi..." | 942 | 942 | exact |
| 4 | If asked | "reflection on unconscious values embedded in computing and the practices that it suppor..." | 938 | 938 | exact |
| 4 | If asked | "an antidote to assumptions that technology should provide clear, efficient solutions to..." | 938 | 938 | exact |
| 4 | If asked | "designing and evaluating for richer human connections and experiences" | 946 | 946 | exact |
| 4 | If asked | "Making products that help people become the person they desire to be" | 946 | 946 | exact |
| 4 | If asked | "involved in designing for products to be used together" | 938 | 938 | exact |
| 4 | If asked | "suggest certain approaches to design as both as desirable and productive of future prac..." | 938 | 938 | exact |
| 4 | If asked | "their primary function is to build an account of a practice to be pursued in the future" | 938 | 938 | exact |
| 4 | If asked | "their ability to inspire new designs" | 941 | 941 | exact |
| 4 | If asked | "There is a proliferation of research programmes, and little agreement about the values ..." | 942 | 942 | exact |
| 4 | Sources | "ability to inspire new designs" | 941 | 941 | exact |
| 5 | On slide | "a level of extensibility and verifiability" | 937 | 937 | exact |
| 5 | On slide | "develop protocols, descriptions, and guidelines" | 937 | 937 | exact |
| 5 | On slide | "Quality control" | 946 | 946 | exact |
| 5 | On slide | "lacks clear expectations and standards" | 937 | 937 | exact |
| 5 | On slide | "actionable metrics" | 937 | 937 | exact |
| 5 | On slide | "may be a route towards disciplinary legitimacy within HCI" | 937 | 937 | exact |
| 5 | Say | "a level of extensibility and verifiability" | 937 | 937 | exact |
| 5 | Say | "develop protocols, descriptions, and guidelines" | 937 | 937 | exact |
| 5 | Say | "Quality control" | 946 | 946 | exact |
| 5 | Say | "lacks clear expectations and standards" | 937 | 937 | exact |
| 5 | Say | "actionable metrics" | 937 | 937 | exact |
| 5 | Say | "may be a route towards disciplinary legitimacy within HCI" | 937 | 937 | exact |
| 5 | Watch out | "may be a route towards disciplinary legitimacy" | 937 | 937 | exact |
| 5 | Watch out | "A level of extensibility and verifiability" | 937 | 937 | exact (sentence-initial capital changed) |
| 5 | Watch out | "Extensible and verifiable" | 938 | 938 | exact (sentence-initial capital changed) |
| 5 | Watch out | "lacks clear expectations and standards for what constitutes" | 937 | 937 | exact |
| 5 | Watch out | "design research" | 937 | 937 | exact |
| 5 | Watch out | "ultimate particular" | 939, 944 | 939, 944 | exact |
| 5 | If asked | "a proper research methodology that can produce relevant and rigorous theory" | 937 | 937 | exact |
| 5 | If asked | "like any other research approach" | 937 | 937 | exact |
| 5 | If asked | "Paper 37" | 946 | 946 | exact |
| 5 | If asked | "Human-Computer Interaction Institute. Paper 37" | 946 | 946 | exact |
| 5 | Sources | "disciplinary legitimacy" | 937 | 937 | exact |
| 6 | On slide | "sometimes with associated accounts of how these are used in field tests" | 937 | 937 | exact |
| 6 | On slide | "'irresistables'" | 938 | 938 | exact |
| 6 | Watch out | "increasingly includes a variety of methods, conceptual frameworks and theories presente..." | 937 | 937 | exact |
| 6 | Watch out | "'irresistables'" | 938 | 938 | exact |
| 6 | Watch out | "A final genre" | 938 | 938 | exact |
| 6 | Watch out | "is intended as a contribution to meta-theory" | 939 | 939 | exact |
| 6 | Watch out | "more specific, content-oriented theories" | 939 | 939 | exact |
| 6 | Watch out | "Zimmerman & Forlizzi [34]" | 938 | 938 | exact |
| 6 | If asked | "are articulated in general terms and applied to multiple examples, they become recognis..." | 938 | 938 | exact |
| 6 | If asked | "implies a conceptual orientation" | 938 | 938 | exact |
| 6 | If asked | "a rich account of the importance of embodied movement" | 938 | 938 | exact |
| 6 | If asked | "a series of designs they and their students have produced" | 938 | 938 | exact |
| 6 | Sources | "increasingly includes" | 937 | 937 | exact |
| 6 | Sources | "'irresistables'" | 938 | 938 | exact |
| 6 | Sources | "contribution to meta-theory" | 939 | 939 | exact |
| 7 | On slide | "actionable metrics" | 937 | 937 | exact |
| 7 | Watch out | "toward a formalization of a research approach" | 946 | 946 | exact |
| 7 | If asked | "the 'right thing'" | 940 | 940 | exact |
| 7 | If asked | "for the 'extensibility and verifiability' for which some of the HCI design community ha..." | 945 | 945 | exact |
| 8 | On slide | "some sympathy with these concerns" | 937 | 937 | exact |
| 8 | On slide | "might lead to a form of self-policing that would be overly restrictive" | 937 | 937 | exact |
| 8 | On slide | "insofar as HCI is prone to 'scientism'" | 938 | 938 | exact |
| 8 | On slide | "may adopt, or be interpreted in terms of, inappropriate 'scientific' models" | 938 | 938 | exact |
| 8 | On slide | "an elephant in the room" | 946 | 946 | exact |
| 8 | Say | "I have some sympathy with these concerns" | 937 | 937 | exact |
| 8 | Say | "might lead to a form of self-policing that would be overly restrictive" | 937 | 937 | exact |
| 8 | Say | "may adopt, or be interpreted in terms of, inappropriate scientific models" | 938 | 938 | INEXACT: inner quotes around scientific dropped |
| 8 | Say | "an elephant in the room" | 946 | 946 | exact |
| 8 | Watch out | "some sympathy" | 937 | 937 | exact |
| 8 | Watch out | "might lead to" | 937 | 937 | exact |
| 8 | Watch out | "insofar as" | 938 | 938 | exact |
| 8 | Watch out | "may adopt, or be interpreted in terms of" | 938 | 938 | exact |
| 8 | Watch out | "cultural assumptions" | 938 | 938 | exact |
| 8 | Watch out | "an elephant in the room...: the move towards" | 946 | 946 | exact |
| 8 | Watch out | "in HCI" | 946 | 946 | exact |
| 8 | Watch out | "in part" | 946 | 946 | exact |
| 8 | Watch out | "Must design become 'scientific'?" | 946 | 946 | exact |
| 8 | Watch out | "may be a route towards disciplinary legitimacy" | 937 | 937 | exact |
| 8 | If asked | "destabilises the use of 'science' as shorthand for a set of practices and criteria assu..." | 941 | 941 | exact |
| 8 | If asked | "that I value for its ability to continually and creatively challenge status quo thinking" | 938 | 938 | exact (crosses page break) |
| 8 | If asked | "Personally, I tend to think that the nature of design practice and research is distinct..." | 941 | 941 | exact |
| 8 | If asked | "depends on the account of science one adopts" | 941 | 941 | exact |
| 8 | If asked | "might lead to" | 937 | 937 | exact |
| 8 | If asked | "the anonymous R4" | 946 | 946 | exact |
| 8 | Sources | "long-term practitioner" | 937 | 937 | exact |
| 8 | Sources | "some sympathy" | 937 | 937 | exact |
| 8 | Sources | "disciplinary legitimacy" | 937 | 937 | exact |
| 8 | Sources | "value for its ability" | 938 | 938 | exact |
| 8 | Sources | "insofar as" | 938 | 938 | exact |
| 8 | Sources | "may adopt, or be interpreted in terms of" | 938 | 938 | exact |
| 9 | On slide | "implicit sometimes" | 940 | 940 | exact |
| 9 | On slide | "tends to be provisional, contingent and aspirational" | 938 | 938 | exact |
| 9 | On slide | "It is difficult to call for research to be more scientific if what it means to be scien..." | 941 | 941 | exact |
| 9 | Watch out | "to complicate notions of what makes theories 'scientific'" | 938 | 938 | exact |
| 9 | Watch out | "I am not presenting Popper's and Lakatos' accounts as particularly representative of Ph..." | 941 | 941 | exact |
| 9 | Watch out | "the criterion of the scientific status of a theory is its falsifiability, or refutabili..." | 939 | 939 | exact |
| 9 | Watch out | "Research through design is clearly unscientific if Popper's [24] criterion of falsifiab..." | 940 | 940 | exact |
| 9 | Watch out | "there is always an implicit sometimes in statements about how to design successfully" | 940 | 940 | exact |
| 9 | Watch out | "Simply tweaking an unambiguous system to be ambiguous" | 940 | 940 | exact |
| 9 | Watch out | "What Popper would reject as 'merely' confirmatory evidence is, from this point of view,..." | 941 | 941 | exact |
| 9 | Watch out | "predict novel facts" | 940 | 940 | exact |
| 9 | Watch out | "tends to be" | 938 | 938 | exact |
| 9 | Watch out | "it is provisional, contingent, and aspirational" | 941 | 941 | exact |
| 9 | Watch out | "likely to produce theories that are provisional, contingent, and aspirational" | 937 | 937 | exact |
| 9 | Watch out | "tends to be" | 938 | 938 | exact |
| 9 | Watch out | "theories that are sometimes right" | 941 | 941 | exact |
| 9 | If asked | "Personally, I tend to think that the nature of design practice and research is distinct..." | 941 | 941 | exact |
| 9 | If asked | "can be considered a science depends on the account of science one adopts" | 941 | 941 | exact |
| 9 | If asked | "X will sometimes lead to successful outcomes" | 940 | 940 | exact |
| 9 | If asked | "Underspecification makes falsification difficult in practice as well as in principle" | 940 | 940 | exact |
| 9 | If asked | "'hard core'" | 939 | 939 | exact |
| 9 | If asked | "'protective belt'" | 939 | 939 | exact |
| 9 | If asked | "'powerful problem-solving mechanism'" | 939 | 939 | exact |
| 9 | If asked | "predict novel facts" | 940 | 940 | exact |
| 9 | If asked | "ability to inspire new designs" | 941 | 941 | exact |
| 9 | If asked | "the relative productivity of competing programmes" | 941 | 941 | exact |
| 9 | If asked | "as long as that theory can lead to productive research programmes" | 941 | 941 | exact |
| 9 | If asked | "(see [29])" | 941 | 941 | exact |
| 9 | If asked | "any easy scientism in HCI" | 941 | 941 | exact |
| 9 | Sources | "provisional, contingent, and aspirational" | 937 | 937 | exact |
| 9 | Sources | "complicate notions" | 938 | 938 | exact |
| 9 | Sources | "tends to be provisional, contingent and aspirational" | 938 | 938 | exact |
| 9 | Sources | "clearly unscientific if" | 940 | 940 | exact |
| 9 | Sources | "implicit sometimes" | 940 | 940 | exact |
| 9 | Sources | "Simply tweaking" | 940 | 940 | exact |
| 9 | Sources | "predict novel facts" | 940 | 940 | exact |
| 9 | Sources | "Underspecification makes falsification difficult" | 940 | 940 | exact |
| 9 | Sources | "theories that are sometimes right" | 941 | 941 | exact |
| 9 | Sources | "productive research programmes" | 941 | 941 | exact |
| 9 | Sources | "not ... particularly representative" | 937 to 941 | 937 to 941 | exact |
| 9 | Sources | "Personally, I tend to think" | 941 | 941 | exact |
| 9 | Sources | "depends on the account of science one adopts" | 941 | 941 | exact |
| 9 | Sources | "relative productivity" | 941 | 941 | exact |
| 10 | On slide | "must founder" | 942 | 942 | exact |
| 10 | On slide | "most of us" | 942 | 942 | exact |
| 10 | On slide | "a different world" | 943 | 943 | exact |
| 10 | Watch out | "Design as Pre-paradigmatic Research" | 941 | 941 | exact |
| 10 | Watch out | "The Invisibility of Consensus" | 942 | 942 | exact |
| 10 | Watch out | "The Many Worlds of Design" | 943 | 943 | exact |
| 10 | Watch out | "must founder" | 942 | 942 | exact |
| 10 | Watch out | "Valuable contributions may be made by disparate approaches" | 942 | 942 | exact |
| 10 | Watch out | "a tempting conclusion, one seemingly behind recent calls for formalism and standardisat..." | 942 | 942 | exact |
| 10 | Watch out | "most of us" | 942 | 942 | exact |
| 10 | Watch out | "most of us working in research through design seem to share a set of common values" | 942 | 942 | exact |
| 10 | Watch out | "Perhaps we do not all agree with all of these assumptions all of the time" | 942 | 942 | exact |
| 10 | Watch out | "seem to" | 942 | 942 | exact |
| 10 | Watch out | "characterises scientific activity in terms of controversy" | 942 | 942 | exact |
| 10 | Watch out | "may be a function of their presentation in the popular media, particularly with the adv..." | 942 | 942 | exact |
| 10 | Watch out | "anybody seeking to research or develop tablet computers ... or anything at all, for tha..." | 943 | 943 | exact |
| 10 | Watch out | "potentially repressive acts of ontological politics" | 943 | 943 | exact |
| 10 | Watch out | "seems naturally to lead to convergence" | 943 | 943 | exact |
| 10 | Watch out | "implies that the community should work towards greater shared understandings, the other..." | 938 | 938 | exact |
| 10 | If asked | "should be wary of impulses towards convergence and standardisation" | 937 | 937 | exact |
| 10 | If asked | "greater consensus may be emerging in research through design than is sometimes acknowle..." | 945 | 945 | exact |
| 10 | If asked | "may well exhibit just as much uncertainty, interpretation and debate as does research t..." | 942 | 942 | exact |
| 10 | If asked | "are precisely those unlikely to be addressed in the literature" | 942 | 942 | exact |
| 10 | If asked | "overestimate the divergence of research through design" | 942 | 942 | exact |
| 10 | If asked | "many worlds" | 943 | 943 | exact (sentence-initial capital changed) |
| 10 | If asked | "different research practices might be making multiple worlds, and that such worlds migh..." | 943 | 943 | exact |
| 10 | If asked | "a strong presumption that the object of study is a single, unitary world that pre-exist..." | 943 | 943 | exact |
| 10 | If asked | "must founder" | 942 | 942 | exact |
| 10 | If asked | "calls for disciplinary reform" | 941 | 941 | exact (sentence-initial capital changed) |
| 10 | If asked | "'methods, approaches, and outcomes'" | 941 | 941 | exact |
| 10 | If asked | "would make assessment of new contributions fairer and more rigorous" | 941 | 941 | exact |
| 10 | Sources | "wary of impulses" | 937 | 937 | exact |
| 10 | Sources | "dozens of different 'manifestos'" | 941 | 941 | exact |
| 10 | Sources | "Why is there not more convergence" | 941 | 941 | exact |
| 10 | Sources | "'methods, approaches, and outcomes'" | 941 | 941 | exact |
| 10 | Sources | "Design as Pre-paradigmatic Research" | 941 | 941 | exact |
| 10 | Sources | "from the very basics every time" | 942 | 942 | exact |
| 10 | Sources | "must founder" | 942 | 942 | exact |
| 10 | Sources | "Valuable contributions" | 942 | 942 | exact |
| 10 | Sources | "tempting conclusion" | 942 | 942 | exact |
| 10 | Sources | "little agreement about the values" | 942 | 942 | exact |
| 10 | Sources | "most of us" | 942 | 942 | exact |
| 10 | Sources | "just as much uncertainty" | 942 | 942 | exact |
| 10 | Sources | "Invisibility of Consensus" | 942 | 942 | exact |
| 10 | Sources | "single, unitary world" | 943 | 943 | exact |
| 10 | Sources | "seems naturally to lead to convergence" | 943 | 943 | exact |
| 10 | Sources | "potentially repressive acts of ontological politics" | 943 | 943 | exact |
| 10 | Sources | "greater consensus may be emerging" | 945 | 945 | exact |
| 11 | On slide | "the definite facts" | 945 | 945 | exact |
| 11 | On slide | "ultimate particular" | 939, 944 | 939, 944 | exact |
| 11 | On slide | "theory nexus" | 944 | 944 | exact |
| 11 | On slide | "establishes an area" | 944 | 944 | exact |
| 11 | On slide | "If artefacts embody theory, however, they do not encode it" | 944 | 944 | exact |
| 11 | Watch out | "the definite facts of research through design" | 945 | 945 | exact |
| 11 | Watch out | "It is the artefacts we create that are the definite facts of research through design" | 945 | 945 | exact |
| 11 | Watch out | "Theories may be provisional, but designed artefacts (as opposed to demonstrators, for e..." | 944 | 944 | exact |
| 11 | Watch out | "ultimate particular" | 939, 944 | 939, 944 | exact |
| 11 | Watch out | "the same dignity and importance as truth in science" | 944 | 944 | exact |
| 11 | Watch out | "theory nexus" | 944 | 944 | exact |
| 11 | Watch out | "If artefacts embody theory, however, they do not encode it" | 944 | 944 | exact |
| 11 | Watch out | "design theory is best considered a form of annotation, serving to explain and point to ..." | 944 | 944 | exact |
| 11 | Watch out | "a portfolio" | 944 | 944 | exact |
| 11 | Watch out | "establishes an area in that space" | 944 | 944 | exact |
| 11 | Watch out | "the invariances among them" | 944 | 944 | exact |
| 11 | Watch out | "theories predominating, with design examples serving as mere illustrations" | 944 | 944 | exact |
| 11 | Watch out | "is limited to inspiration and annotation" | 945 | 945 | exact |
| 11 | Watch out | "Theoretical writing remains important" | 945 | 945 | exact |
| 11 | If asked | "theory nexus" | 944 | 944 | exact |
| 11 | If asked | "reveal both the issues they think are important, and their beliefs about the right way ..." | 944 | 944 | exact |
| 11 | If asked | "ultimate particular" | 939, 944 | 939, 944 | exact |
| 11 | If asked | "often turn to existing examples of design" | 943 | 943 | exact |
| 11 | If asked | "both because of the provisional nature of theory, and the definite nature of designed a..." | 943 | 943 | exact |
| 11 | If asked | "[27] cited in [31]" | 943 | 943 | exact |
| 11 | If asked | "simultaneous control over multiple functions" | 944 | 944 | exact |
| 11 | If asked | "it is useful to consider specific designs such as a rotary controller for a microwave o..." | 944 | 944 | exact |
| 11 | If asked | "Design space" | 944 | 944 | exact (sentence-initial capital changed) |
| 11 | If asked | "perhaps more accurately creating a design space around themselves (see [12])" | 944 | 944 | exact |
| 11 | If asked | "or even by jumping away from it along understood dimensions" | 944 | 944 | exact |
| 11 | If asked | "Many of the ideas here about annotated portfolios have been developed in collaboration ..." | 944 | 944 | exact |
| 11 | Sources | "turn to existing examples of design" | 943 | 943 | exact |
| 11 | Sources | "provisional nature of theory, and the definite nature" | 943 | 943 | exact |
| 11 | Sources | "[27] cited in [31]" | 943 | 943 | exact |
| 11 | Sources | "ultimate particular" | 939, 944 | 939, 944 | exact |
| 11 | Sources | "the same dignity and importance as truth in science" | 944 | 944 | exact |
| 11 | Sources | "Theories may be provisional" | 944 | 944 | exact |
| 11 | Sources | "theory nexus" | 944 | 944 | exact |
| 11 | Sources | "embody ... do not encode" | 937 to 944 | 937 to 944 | exact |
| 11 | Sources | "best considered a form of annotation" | 944 | 944 | exact |
| 11 | Sources | "theories predominating" | 944 | 944 | exact |
| 11 | Sources | "the invariances among them" | 944 | 944 | exact |
| 11 | Sources | "creating a design space around themselves" | 944 | 944 | exact |
| 11 | Sources | "the definite facts of research through design" | 945 | 945 | exact |
| 11 | Sources | "limited to inspiration and annotation" | 945 | 945 | exact |
| 11 | Sources | "Theoretical writing remains important" | 945 | 945 | exact |
| 12 | On slide | "good design is unobtrusive" | 944 | 944 | exact |
| 12 | On slide | "good design is thorough down to the last detail" | 944 | 944 | exact |
| 12 | On slide | "neither dominates, and neither is subservient" | 945 | 945 | exact |
| 12 | On slide | "One can imagine 'designing in the style of Dieter Rams' without the need for a detailed..." | 944 | 944 | exact |
| 12 | Watch out | "Sketch design for an annotated portfolio of Dieter Rams' designs for Braun and Vitsoe (..." | 945 | 945 | exact |
| 12 | Watch out | "for Braun and Vitsoe" | 944 | 944 | exact |
| 12 | Watch out | "too terse to serve as substantial conceptual contributions by themselves" | 944 | 944 | exact |
| 12 | Watch out | "neither dominates, and neither is subservient" | 945 | 945 | exact |
| 12 | Watch out | "annotate the portfolio, drawing attention to important features and to salient details ..." | 944 | 944 | exact |
| 12 | If asked | "Comparing the variety of stereo equipment, kitchenware and grooming appliances that he ..." | 944 | 944 | exact |
| 12 | Sources | "ten concise principles" | 944 | 944 | exact |
| 12 | Sources | "annotate the portfolio" | 944 | 944 | exact |
| 12 | Sources | "too terse" | 944 | 944 | exact |
| 12 | Sources | "Comparing the variety" | 944 | 944 | exact |
| 12 | Sources | "neither dominates, and neither is subservient" | 945 | 945 | exact |
| 13 | On slide | "The notion of annotated portfolios is not a formal one" | 944 | 944 | exact |
| 13 | On slide | "the issues, values and themes" | 944 | 944 | exact |
| 13 | On slide | "in several respects, the converse" | 945 | 945 | exact |
| 13 | On slide | "'extensibility and verifiability'" | 941, 945 | 941, 945 | exact |
| 13 | Watch out | "The notion of annotated portfolios is not a formal one" | 944 | 944 | exact |
| 13 | Watch out | "a balance is achieved between descriptions of specific, detailed examples of design pra..." | 944 | 944 | exact |
| 13 | Watch out | "may vary depending on material, purpose and audience" | 944 | 944 | exact |
| 13 | Watch out | "emphasises his design principles over the portfolio of products themselves" | 945 | 945 | exact |
| 13 | Watch out | "A more detailed examination of Rams' philosophy might well take the form of an illustra..." | 945 | 945 | exact |
| 13 | Watch out | "annotate the portfolio" | 944 | 944 | exact |
| 13 | Watch out | "Annotated portfolios are, in several respects, the converse of Alexander's [1] design p..." | 945 | 945 | exact |
| 13 | Watch out | "in several respects" | 945 | 945 | exact |
| 13 | Watch out | "abstract regularities from repeated attempts to design for the same domains" | 945 | 945 | exact |
| 13 | Watch out | "maintain the particularity of individual examples" | 945 | 945 | exact |
| 13 | Watch out | "I am not proposing here that annotated portfolios subsume all other forms of design the..." | 945 | 945 | exact |
| 13 | Watch out | "while allowing for the 'extensibility and verifiability' for which some of the HCI desi..." | 945 | 945 | exact |
| 13 | If asked | "in several respects" | 945 | 945 | exact |
| 13 | If asked | "Portfolios can support multiple conceptual perspectives, and similar perspectives can b..." | 945 | 945 | exact |
| 13 | If asked | "are extremely compact and focus more on highlighting features of the artefacts than on ..." | 944 | 944 | exact |
| 13 | Sources | "not a formal one" | 944 | 944 | exact |
| 13 | Sources | "material, purpose and audience" | 944 | 944 | exact |
| 13 | Sources | "extremely compact" | 944 | 944 | exact |
| 13 | Sources | "annotate the portfolio" | 944 | 944 | exact |
| 13 | Sources | "converse of Alexander's [1] design patterns" | 945 | 945 | exact |
| 13 | Sources | "abstract regularities" | 945 | 945 | exact |
| 13 | Sources | "maintain the particularity" | 945 | 945 | exact |
| 13 | Sources | "multiple conceptual perspectives" | 945 | 945 | exact |
| 13 | Sources | "not proposing ... subsume" | 945 | 945 | exact |
| 13 | Sources | "Theoretical writing remains important" | 945 | 945 | exact |
| 13 | Sources | "'extensibility and verifiability'" | 941, 945 | 941, 945 | exact |
| 14 | On slide | "annotated portfolios respect the 'ultimate particular'" | 945 | 945 | exact |
| 14 | Say | "ultimate particular" | 939, 944 | 939, 944 | exact |
| 14 | Say | "extensibility and verifiability" | 937 | 937 | exact |
| 14 | Watch out | "Most fundamentally, annotated portfolios respect the 'ultimate particular' of the desig..." | 945 | 945 | exact |
| 14 | Watch out | "Ultimate particular" | 939, 944 | 939, 944 | exact (sentence-initial capital changed) |
| 14 | Watch out | "Juxtaposing designs with annotations supports appreciation of the conceptual dimensions..." | 945 | 945 | exact |
| 14 | Watch out | "grounds and specifies" | 945 | 945 | exact |
| 14 | Watch out | "extensibility and verifiability" | 937 | 937 | exact |
| 14 | Watch out | "a level of extensibility and verifiability" | 937 | 937 | exact |
| 14 | Watch out | "Portfolios can support multiple conceptual perspectives, and similar perspectives can b..." | 945 | 945 | exact |
| 14 | Watch out | "Most practicing designers, however, do not engage with major theoretical approaches in ..." | 943 | 943 | exact |
| 14 | Watch out | "The notion of annotated portfolios is not a formal one" | 944 | 944 | exact |
| 14 | Watch out | "I am not proposing here that annotated portfolios subsume all other forms of design theory" | 945 | 945 | exact |
| 14 | If asked | "extensibility and verifiability" | 937 | 937 | exact |
| 14 | If asked | "moderate expectations of creating extensible and verifiable theory" | 945 | 945 | exact |
| 14 | If asked | "in several respects, the converse of Alexander's [1] design patterns" | 945 | 945 | exact |
| 14 | If asked | "maintain the particularity of individual examples" | 945 | 945 | exact |
| 14 | Sources | "a level of extensibility and verifiability" | 937 | 937 | exact |
| 14 | Sources | "turn to existing examples of design" | 943 | 943 | exact |
| 14 | Sources | "definite nature of designed artefacts" | 943 | 943 | exact |
| 14 | Sources | "ultimate particular" | 939, 944 | 939, 944 | exact |
| 14 | Sources | "not a formal one" | 944 | 944 | exact |
| 14 | Sources | "Most fundamentally" | 945 | 945 | exact |
| 14 | Sources | "multiple conceptual perspectives" | 945 | 945 | exact |
| 14 | Sources | "converse of Alexander's" | 945 | 945 | exact |
| 14 | Sources | "not proposing here" | 945 | 945 | exact |
| 15 | On slide | "Productive research programmes" | 941 | 941 | exact (sentence-initial capital changed) |
| 15 | On slide | "most of us" | 942 | 942 | exact |
| 15 | Watch out | "are not defined by their graphic presentation" | 945 | 945 | exact |
| 15 | Watch out | "might well take the form of an illustrated essay" | 945 | 945 | exact |
| 15 | Watch out | "the relative productivity of competing programmes as the grounds for scientists' choice..." | 941 | 941 | exact |
| 15 | Watch out | "may nonetheless inspire thriving research programmes" | 945 | 945 | exact |
| 15 | Watch out | "political acts to be approached with care" | 945 | 945 | exact |
| 15 | Watch out | "on its own terms" | 945 | 945 | exact |
| 15 | Watch out | "Greater integration would make assessment of new contributions fairer and more rigorous" | 941 | 941 | exact |
| 15 | Watch out | "actionable metrics" | 937 | 937 | exact |
| 15 | Watch out | "most of us working in research through design seem to share a set of common values" | 942 | 942 | exact |
| 15 | Watch out | "Perhaps we do not all agree with all of these assumptions all of the time" | 942 | 942 | exact |
| 15 | Watch out | "we already share many of the attributes" | 942 | 942 | exact |
| 15 | Watch out | "of a research paradigm" | 943 | 943 | exact |
| 15 | Watch out | "addresses the more specific, content-oriented theories described earlier" | 939 | 939 | exact |
| 15 | Watch out | "can be considered as wide-ranging manifestos" | 939 | 939 | exact |
| 15 | Watch out | "a contribution to meta-theory" | 939 | 939 | exact |
| 15 | Watch out | "overly restrictive" | 937 | 937 | exact |
| 15 | Watch out | "a collection of designs by the same or associated designers" | 944 | 944 | exact |
| 15 | Watch out | "European Research Council's Advanced Investigator Award no. 226528, 'ThirdWaveHCI'" | 946 | 946 | exact |
| 15 | If asked | "the 'right thing'" | 940 | 940 | exact |
| 15 | If asked | "a tempting conclusion, one seemingly behind recent calls for formalism and standardisat..." | 942 | 942 | exact |
| 15 | Sources | "actionable metrics" | 937 | 937 | exact |
| 15 | Sources | "overly restrictive" | 937 | 937 | exact |
| 15 | Sources | "a contribution to meta-theory" | 939 | 939 | exact |
| 15 | Sources | "productive research programmes" | 941 | 941 | exact |
| 15 | Sources | "relative productivity" | 941 | 941 | exact |
| 15 | Sources | "fairer and more rigorous" | 941 | 941 | exact |
| 15 | Sources | "most of us" | 942 | 942 | exact |
| 15 | Sources | "tempting conclusion" | 942 | 942 | exact |
| 15 | Sources | "same or associated designers" | 944 | 944 | exact |
| 15 | Sources | "illustrated essay" | 945 | 945 | exact |
| 15 | Sources | "political acts" | 945 | 945 | exact |
| 15 | Sources | "own terms" | 945 | 945 | exact |
| 15 | Sources | "thriving research programmes" | 945 | 945 | exact |
| 16 | On slide | "Productive research programmes" | 941 | 941 | exact (sentence-initial capital changed) |
| 16 | Watch out | "moderate expectations of creating extensible and verifiable theory" | 945 | 945 | exact |
| 16 | Watch out | "attempts to establish disciplinary norms of process or outcome are political acts to be..." | 945 | 945 | exact |
| 16 | Watch out | "the recognition of annotated portfolios as a way of reflecting and valuing the particul..." | 946 | 946 | exact |
| 16 | Watch out | "consensus can come to look like constraint" | 946 | 946 | exact |
| 16 | Watch out | "what this discussion suggests is the desirability of clearly articulating the methodolo..." | 941 | 941 | exact |
| 16 | Watch out | "most of us working in research through design seem to share a set of common values" | 942 | 942 | exact |
| 16 | Watch out | "Perhaps we do not all agree with all of these assumptions all of the time" | 942 | 942 | exact |
| 16 | Watch out | "productive research programmes" | 941 | 941 | exact |
| 16 | Watch out | "the relative productivity of competing programmes" | 941 | 941 | exact |
| 16 | Watch out | "(note that annotated portfolios are not defined by their graphic presentation)" | 945 | 945 | exact |
| 16 | Watch out | "The notion of annotated portfolios is not a formal one" | 944 | 944 | exact |
| 16 | Watch out | "It is difficult to call for research to be more scientific if what it means to be scien..." | 941 | 941 | exact |
| 16 | Watch out | "insofar as HCI is prone to 'scientism'" | 938 | 938 | exact |
| 16 | Watch out | "in part" | 946 | 946 | exact |
| 16 | Watch out | "even the assumption that shared standards are necessary, possible or desirable, are pot..." | 943 | 943 | exact |
| 16 | If asked | "(see [29])" | 941 | 941 | exact |
| 16 | If asked | "the standards one might want to advocate" | 941 | 941 | exact |
| 16 | If asked | "protocols, descriptions, and guidelines" | 937 | 937 | exact |
| 16 | If asked | "on its own terms" | 945 | 945 | exact |
| 16 | If asked | "its own terms" | 945 | 945 | exact |
| 16 | If asked | "may provide a way to further our ambitions to produce relevant and rigorous theoretical..." | 946 | 946 | exact |
| 16 | If asked | "relevant and rigorous theory" | 937 | 937 | exact |
| 16 | Sources | "protocols, descriptions, and guidelines" | 937 | 937 | exact |
| 16 | Sources | "relevant and rigorous theory" | 937 | 937 | exact |
| 16 | Sources | "prone to 'scientism'" | 938 | 938 | exact |
| 16 | Sources | "clearly articulating" | 941 | 941 | exact |
| 16 | Sources | "productive research programmes" | 941 | 941 | exact |
| 16 | Sources | "relative productivity" | 941 | 941 | exact |
| 16 | Sources | "It is difficult to call" | 941 | 941 | exact |
| 16 | Sources | "most of us" | 942 | 942 | exact |
| 16 | Sources | "potentially repressive acts of ontological politics" | 943 | 943 | exact |
| 16 | Sources | "not a formal one" | 944 | 944 | exact |
| 16 | Sources | "own terms" | 945 | 945 | exact |
| 16 | Sources | "consensus can come to look like constraint" | 946 | 946 | exact |
| 16 | Sources | "relevant and rigorous theoretical work" | 946 | 946 | exact |
| 18 | On slide | "the definite nature of designed artefacts" | 943 | 943 | exact |
| 18 | Watch out | "The definite nature of designed artefacts" | 943 | 943 | exact (sentence-initial capital changed) |
| 18 | Watch out | "both because of the provisional nature of theory, and the definite nature of designed a..." | 943 | 943 | exact |
| 18 | If asked | "multiple, incompatible worlds" | 943 | 943 | exact (sentence-initial capital changed) |
| 18 | If asked | "the definite facts of research through design" | 945 | 945 | exact |
| 18 | Sources | "definite nature of designed artefacts" | 943 | 943 | exact |
| 18 | Sources | "multiple, incompatible worlds" | 943 | 943 | exact (sentence-initial capital changed) |
| 18 | Sources | "definite facts" | 945 | 945 | exact |
| 19 | Watch out | "practice is usually intended to confirm theories, not falsify them" | 940 | 940 | exact |
| 19 | Watch out | "which are simultaneously unambiguous and ecologically valid are inconceivable" | 940 | 940 | exact |
| 19 | If asked | "the ideas and issues that join and differentiate them" | 945 | 945 | exact |
| 19 | If asked | "ultimate particular" | 939, 944 | 939, 944 | exact |
| 19 | If asked | "establishes an area" | 944 | 944 | exact |
| 19 | If asked | "wary of impulses towards convergence and standardisation" | 937 | 937 | exact |
| 20 | Watch out | "a single, unitary world that pre-exists and is independent of its observers" | 943 | 943 | exact |
| 20 | If asked | "able to create multiple new worlds rather than describing a single existing one" | 943 | 943 | exact |
| 21 | On slide | "sometimes right" | 941 | 941 | exact |
| 21 | Watch out | "disciplinary legitimacy" | 937 | 937 | exact |
| 21 | Watch out | "theories that are sometimes right" | 941 | 941 | exact |
| 21 | Watch out | "productive research programmes" | 941 | 941 | exact |
| 21 | Watch out | "a serious form of theoretical contribution" | 945 | 945 | exact (crosses column break) |
| 21 | Watch out | "neither dominates, and neither is subservient" | 945 | 945 | exact |
| 21 | If asked | "may be a route towards disciplinary legitimacy within HCI, helping clarify expectations..." | 937 | 937 | exact |
| 21 | If asked | "a form of self-policing that would be overly restrictive of a form of research that I v..." | 937 to 938 | 937 to 938 | exact (crosses page break) |
| 21 | If asked | "actionable metrics" | 937 | 937 | exact |
| 21 | If asked | "sometimes right" | 941 | 941 | exact |
| 21 | If asked | "is to sometimes create good ones" | 940 | 940 | exact |
| 21 | If asked | "testimony to the fertility of the overarching theory" | 941 | 941 | exact |
| 21 | If asked | "irrefutability is not a virtue of a theory (as people often think) but a vice" | 939 | 939 | exact |
| 21 | If asked | "As artefacts are to theory ... design portfolios are to research programmes" | 944 | 944 | exact |
| 21 | If asked | "It is clearly possible to recognise the current state of research through design in thi..." | 942 | 942 | exact |
| 21 | If asked | "consensus is usually easier to see from a distance than it is to experience from within" | 942 | 942 | exact |
| 21 | If asked | "most of us" | 942 | 942 | exact |
| 21 | If asked | "as a serious form of theoretical contribution" | 945 | 945 | exact (crosses column break) |
| 21 | If asked | "actionable metrics" | 937 | 937 | exact |
| 21 | If asked | "neither dominates, and neither is subservient" | 945 | 945 | exact |
| 21 | If asked | "I am not proposing here that annotated portfolios subsume all other forms of design theory" | 945 | 945 | exact |
| 21 | If asked | "consensus can come to look like constraint" | 946 | 946 | exact |
| 22 | Watch out | "a collection of designs by the same or associated designers" | 944 | 944 | exact |

## Outside-fact table

| Slide | Claim | Verified source | Status |
|---|---|---|---|
| 1 | Gaver PhD UC San Diego 1988; Apple; cultural probes at RCA 1999 (slide 2 If asked) | slides/01-Background.md log (icad.org; research.gold.ac.uk 15061; interactions 1999) | confirmed |
| 1 | Title Say: "William Gaver runs a design research studio at Goldsmiths" (present tense) | slides/01-Background.md log: studio moved to Northumbria late 2021 (research.northumbria.ac.uk/irs/?p=12) | NOT CONFIRMED (false for 2026) |
| 1 | CHI 2012, Austin, session Critical Perspectives on Design, pp. 937 to 946 | paper header; slides/01-Background.md log | confirmed |
| 2 | Frayling 1993, RCA, into / through / for (labels vary: Biggs "in") | slides/01-Background.md log (researchonline.rca.ac.uk 384, 4907; Biggs 2002) | confirmed (secondary) |
| 2 | Frayling's three kinds described (thesis; materials, development, action research; "thorny") | slides/01-Background.md log, depth pass | confirmed (secondary) |
| 2 | Drift Table 2004, Gaver and colleagues at RCA, coffee table, round screen, weight shifts drift | slides2/01-Hook.md log; slides/01-Background.md log | confirmed |
| 2 | "of England and Wales", "about a terabyte of photography of England and Wales" (CHI 2004 p. 886) | slides2/01-Hook.md log row 5; terabyte phrase re-read in https://research.gold.ac.uk/id/eprint/4721/1/p885-gaver.pdf (fetched 2026-10-21) | confirmed |
| 2 | "motivated by curiosity, exploration, and reflection rather than externally-defined tasks"; several weeks in volunteers' homes | slides/01-Background.md log; same PDF re-read | confirmed |
| 2 | Zimmerman et al. 2007 p. 496: "a well known example of critical design in HCI" | slides/01-Background.md log | confirmed |
| 2 | Zimmerman and Forlizzi 2014 p. 175: "has been held up as one of the earliest and best examples of RtD published within HCI" | slides/10-What-Followed.md log | confirmed |
| 3 | Cross 1982, Design Studies 3(4) 221 to 227; design as its own culture beside science and humanities | NEW: api.crossref.org 10.1016/0142-694X(82)90040-0 and Wikipedia re-fetched; also slides/08-Strengths-Limitations.md log row (original PDF, pp. 221 to 222, three cultures) | confirmed |
| 3 | Fallman 2003, "Design-oriented human-computer interaction", CHI 2003, pp. 225 to 232 | NEW: api.crossref.org 10.1145/642611.642652 re-fetched (event CHI03) | confirmed |
| 3 | Fallman abstract quotations ("has emerged as a design-oriented field of research, directed at large ...", "relatively poor", conservative/romantic/pragmatic, sketching, "knowledge-generating Design-oriented Research", "the artifact-generating conduct of Research-oriented Design") | NEW: api.openalex.org/works/doi:10.1145/642611.642652 re-fetched, abstract rebuilt | confirmed, exact |
| 3 | Bødker 2006 NordiCHI, third wave, workplace to homes and everyday life | slides/01-Background.md log | confirmed (abstract) |
| 3 | Zimmerman, Forlizzi, Evenson 2007 CHI: named method, four criteria | slides/01-Background.md log | confirmed |
| 3 | Schön 1983 is ref. [28], never cited in text | slides/08-Strengths-Limitations.md log; paper p. 946 | confirmed |
| 4 | Supple Interfaces (Isbister and Höök 2007), Designing for the Self (Zimmerman 2009) titles | paper p. 946 refs 16, 36 | confirmed |
| 5 | Panel [9] authors Forlizzi, DiSalvo, Bardzell, Koskinen, Wensveen; CHI 2011 | slides/01-Background.md log | confirmed |
| 5 | 2007 p. 499: "there is no agreed upon standard of what research through design means nor what a high quality contribution should be" | slides/01-Background.md log | confirmed |
| 5 | Panel abstract "fundamental differences"; 2010 abstract 12 interviews | slides/01-Background.md log | confirmed |
| 5 | [33] = Artifact 2(1) 2008, pp. 41 to 45, CMU "Paper 37" | slides/01-Background.md log | confirmed |
| 6 | Prayer Companion CHI 2010, cloistered nuns, ten months, meaningful lasting role | slides/01-Background.md log and background item 3 | confirmed |
| 6 | Product Ecology five factors, Roomba household study | slides/04-Theory.md log | confirmed |
| 7 | Four lenses and glosses (process, invention, relevance replacing validity, extensibility), pp. 499 to 500 | slides/01-Background.md background item 2 and log; slides/02-Problem.md log | confirmed |
| 7 | "no expectation that reproducing the process will produce the same results"; "Instead of validity, the benchmark for interaction design research should be relevance" | slides/01-Background.md log; slides/02-Problem.md log | confirmed |
| 7 | 2007 model "true", "how", "real", "the right thing", "preferred state" (pp. 493, 497); "the idea is not new" (p. 499); Xerox, Philips, Apple (pp. 500 to 501) | slides/01-Background.md log | confirmed |
| 7 | 2010 subtitle "toward" (paper) vs "towards" (Crossref) | slides/01-Background.md Watch out (line 64) | confirmed |
| 8 | Markussen, Krogh and Bang 2015 "a policing of the research through design label"; Bardzell 2016 reply | slides/02-Problem.md log | confirmed (secondary) |
| 9 | Popper 1902 to 1994, Vienna; chapter began as 1953 lecture | slides/04-Theory.md log (SEP; Brown 1984) | confirmed |
| 9 | Lakatos 1922 to 1974, LSE | slides/04-Theory.md log (SEP) | confirmed |
| 10 | Kuhn, Structure ch. II p. 13: "made significant contributions", "something less than science" | slides/05-Convergence.md log | confirmed |
| 11 | Carroll and Kellogg, "Artifact as theory-nexus ...", CHI 1989, pp. 7 to 14; multiple psychological claims | slides/06-Annotation.md log | confirmed |
| 11 | Stolterman 2008 p. 59, ultimate particular "in its specific use context" | slides/06-Annotation.md log | confirmed |
| 12 | Figure 1: seven photographs; coffee maker and small portable device in three colours; four principles (2, 4, 5, 8); third-person notes; heating-elements note; TV prototype "never produced"; Flickr CC credits 1 to 4; caption | P7; slides/06-Annotation.md log; assets/rams-fig1-annotated.png opened 2026-10-21 | confirmed |
| 12 | Product list: coffee maker, shelving, radio, lounge chair, portable TV, shaver, small portable device | slides/06-Annotation.md lines 59, 153; image | confirmed |
| 12 | Figure prints "thorough, down to the last detail" with comma | image | confirmed |
| 12 | Vitsoe ten principles, order, late 1970s | slides/06-Annotation.md log (Vitsoe archive Jan 2026) | confirmed (Vitsoe spells "environmentally-friendly") |
| 12 | Figure glosses "'both neutral and restrained'", "'nothing must be arbitrary or left to chance'", "'at best, it is self explanatory'" vs Vitsoe "self-explanatory" | slides/06-Annotation.md log; image | confirmed |
| 13 | A Pattern Language 1977, Alexander, Ishikawa, Silverstein, 253 patterns, photograph and diagram; pattern 159 | slides/06-Annotation.md log | confirmed |
| 13 | Vitsoe site pairs each principle with one product; principle 8 shows a Lubs calculator | slides/06-Annotation.md background item 4 and log | confirmed |
| 14 | Löwgren 2013 "horizontal and vertical grounding" | slides/10-What-Followed.md log | confirmed |
| 15 | CORDIS 226528: 2,439,757 euros, April 2009 to December 2014, Goldsmiths, 50 to 100 prototypes | slides/10-What-Followed.md log; slides/01-Background.md log | confirmed |
| 15 | Gaver 2014 p. 147 "The basic question here is 'does it work?'" | slides/08-Strengths-Limitations.md log | confirmed |
| 15 | Pullin and Schön listed, never cited; Cross absent | slides/08-Strengths-Limitations.md log | confirmed |
| 15 | Gaver and Bowers published two months after CHI (July vs May 2012) | slides/10-What-Followed.md log | confirmed |
| 17 | Gaver and Bowers, Interactions 19(4), July 2012 (no dagger: reference list not checked) | slides/10-What-Followed.md log | confirmed |
| 17 | Bowers DIS June 2012, cites essay | slides/10-What-Followed.md log | confirmed |
| 17 | RTD conference 2013, exhibition and round tables | slides/10-What-Followed.md log | confirmed |
| 17 | DIS Pictorials from 2014 | slides/10-What-Followed.md log | confirmed |
| 17 | Pierce DIS 2014 cites essay; "thingly publications" | slides/10-What-Followed.md log (abstract "offers the notions of thingly publications") | confirmed; Say "asked for" slightly strong |
| 17 | Höök and Löwgren, TOCHI October 2012, cites essay; strong-concept quotations | slides/10-What-Followed.md log | confirmed |
| 17 | Dalsgaard and Dindler CHI 2014, no dagger; abstract quotation | slides/10-What-Followed.md log | confirmed |
| 17 | Zimmerman and Forlizzi 2014 do not cite; "worked to formalize this research approach" (p. 178); five steps (p. 184) | slides/10-What-Followed.md log | confirmed |
| 17 | Bardzell et al. 2016, "a framework for planning and evaluating RtD documentation", cites essay | slides/10-What-Followed.md log | confirmed |
| 17 | Löwgren 2013 cites essay; Bowers "purely indexical ..."; "influencing autonomous drift" | slides/10-What-Followed.md log | confirmed |
| 17 | Ways of Knowing in HCI 2014; Gaver p. 161 "Annotated portfolios may capture best what we can know through design" | slides/10-What-Followed.md log | confirmed |
| 17 | About 1,100 citations (OpenAlex 1,139 on 2026-10-19) | slides/10-What-Followed.md log | confirmed |
| 17 | DIS 2026 Pictorials call wording, 12 pages, dual anonymous review | slides/10-What-Followed.md log | confirmed |
| 18 | CHI 2026 Design subcommittee "alternatives to scientistic ways of knowing"; full sentence | slides/10-What-Followed.md log | confirmed |
| 18 | 18 subcommittees, 6,730 submissions, Design 485, 37.9% vs 38.7% | slides/10-What-Followed.md log | confirmed |
| 18 | "two technical papers committees devoted to design research" (2014, p. 178) | slides/10-What-Followed.md log | confirmed |
| 18 | RTD 2013 BALTIC Gateshead, 2015 Cambridge, 2017 Edinburgh, 2019 Delft and Rotterdam, paused, RTD Roams 2027 York, Melbourne (Naarm), Tokyo, announced 23 July 2026 | slides/10-What-Followed.md log | confirmed |
| 18 | Pictorials at TEI (2020 call) and C&C (2023 call) | slides/10-What-Followed.md log | confirmed |
| 18 | Suh et al. CHI 2024 "hundreds and thousands", "towards rapid convergence on a limited set of ideas"; Davis et al. CHI 2024 title | slides/10-What-Followed.md log | confirmed |
| 19 | Tangible Bits p. 240 "one user interface metaphor ..."; "somewhat weaker"; 1997, fifteen years before 2012 | slides/09-Connections.md log | confirmed |
| 19 | Greenberg and Buxton 2008 p. 113 "confirmatory evidence"; split cure p. 118 | slides/09-Connections.md log | confirmed |
| 19 | Wobbrock et al. method from 2005; 216 studies, DIS 2020 review, p. 855, four authors | slides/09-Connections.md log; P12 | confirmed |
| 19 | Shin et al. 2019: four hours, "ecological validity" p. 5, actuation day 2, "all ... except P2" p. 7, health instruction p. 6 | slides/09-Connections.md log | confirmed |
| 19 | Olsen 2007 "the space between"; Wobbrock and Kientz 2016 p. 44 quotations | slides/09-Connections.md log | confirmed |
| 21 | Zimmerman et al. 2007 p. 500 extensibility "described and documented in a way that the community can leverage the knowledge derived from the work" | slides/11-Discussion-Bridge.md log | confirmed |
| 21 | Course brief "a conversation, not an additional lecture" | slides/11-Discussion-Bridge.md log | confirmed |
| 22 | Rowe wait time, via Allen and Tanner 2002 | slides/11-Discussion-Bridge.md log | confirmed (secondary) |
| 22 | Tangible Bits "can be applied seamlessly" p. 240 | slides/09-Connections.md log | confirmed |
