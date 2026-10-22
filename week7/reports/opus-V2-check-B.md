# Version 2 deck check B (independent)

Checked 2026-10-21 against `assets/paper-text.txt` (read in full first), `slides/` v1 files for outside facts, and `Instructions Week 7.md`. Files checked: `slides2/00-Overview.md` and `01-Hook.md` to `09-Handover.md`. No file in `reports/` was opened.

## Summary

- Quotations checked: 141 (118 from the paper, 23 from outside works).
- Exact: 141. Inexact wording: 0. (One quotation is exact but attributed to the wrong speaker; see B1.)
- Pages given for paper quotations: 118 checked, 118 right, 0 wrong. One soft spot: "provisional, contingent, and aspirational" is cited to p. 937 and p. 938; it is exact on p. 937 and p. 941, while p. 938 reads "provisional, contingent and aspirational" (no serial comma).
- Outside facts: all 23 outside quotations and every outside fact traced to the named v1 file and reproduced faithfully, with one wording slip (S17).
- New facts (in neither the paper nor slides/): 7, of which 4 are marked in the logs and 3 are not.
- Blockers: 2. Should fix: 18. Nice to have: 7.

## Blockers

**B1. Misattributed quotation.** File `03-Two-Tests.md`, slide 03.2, On the slide, bullet 2.
- Current: `What Popper would reject as "'merely' confirmatory," Lakatos calls "testimony to the fertility of the overarching theory."`
- Problem: "testimony to the fertility of the overarching theory" is Gaver's wording (p. 941: "is, from this point of view, testimony to the fertility of the overarching theory"), not Lakatos's. The slide puts Gaver's words in Lakatos's mouth.
- Replacement: `What Popper would reject as "'merely' confirmatory" is, seen Lakatos's way, "testimony to the fertility of the overarching theory."`

**B2. False claim in the Say.** File `06-Aftermath.md`, slide 06.1, Say.
- Current: `The man who warned against standardising proposed a form, and the form became a named, citable method with its own conference.`
- Problem: the RTD conference is a conference for research through design in general, not for annotated portfolios; no v1 source says portfolios have their own conference. "Method" also overstates: Gaver calls it "a serious form of theoretical contribution" (p. 945) and "not a formal one" (p. 944). The twist is our reading but is not marked as such.
- Replacement: `Here is the twist, as we read it. The man who warned against standardisation proposed a form, and the form got a name, a citable paper, and venues that exhibit artefacts.`
- Matching slide bullet, current: `**Twist:** the anti-standard proposal became a named method.` Replacement: `**Twist (ours):** the warning against standards produced a named, citable form.`

## Should fix

**S1. Dropped hedge "insofar as".** `02-Fight.md`, slide 02.1, On the slide, bullet 2.
- Current: `**Gaver, 2012:** standards risk "self-policing," and HCI is "prone to 'scientism.'"`
- Paper p. 938: "insofar as HCI is prone to 'scientism' ... attempts to establish standards ... may adopt ... inappropriate 'scientific' models."
- Replacement: `**Gaver, 2012:** standards risk "self-policing," insofar as HCI is "prone to 'scientism.'"`
- Same in Say. Current: `Second, and this is the sharper point, HCI is "prone to scientism," so any standards it adopts may quietly become the standards of experimental science, which he thinks do not fit design.` Replacement: `Second, the sharper point. Insofar as HCI is prone to scientism, as he puts it, its standards may quietly become those of experimental science. He thinks those do not fit design.`

**S2. Unsourced claim stated as the critics' case.** `02-Fight.md`, slide 02.1, Say.
- Current: `Design research had arrived at CHI, but it was being reviewed by people who did not know how to judge it.`
- Not in the paper or in any slides/ file, and not marked. The paper's nearest text is p. 939: HCI "will more easily accept contributions" if research through design has an "agreed upon form of practice, evaluation, and outcome."
- Replacement: `Design research had arrived at CHI, and they argued that HCI would accept it more easily if it had an agreed form of practice, evaluation and outcome.`

**S3. Motive stated more strongly than Gaver states it.** `02-Fight.md`, slide 02.1, Say.
- Current: `Gaver reads it as a bid for legitimacy: if we have standards, design counts as research like anything else.`
- Paper p. 937: standards "may be a route towards disciplinary legitimacy." Gaver grants a benefit; he does not call it a bid.
- Replacement: `Gaver grants that standards may be a route to legitimacy: with them, design counts as research like any other.`

**S4. Transition overstates and does not meet the next opening.** `02-Fight.md`, slide 02.1, Transition; slide 02.2, first sentence of Say.
- Current transition: `His first move is not to defend design. It is to attack the word "science."` The 02.2 Say then opens `This is an essay.` and reaches "science" only five sentences later. Gaver's own verbs are "complicate" (p. 938) and "problematise" (p. 939).
- Replacement transition: `His first move is not to defend design. It is to question the word "science."`
- Replacement for the first two sentences of the 02.2 Say: `So which science? This is an essay, with no study, no data and no system.`

**S5. "Tests" implies an experiment.** `03-Two-Tests.md`, slide 03.1, On the slide, bullet 2.
- Current: `Gaver tests his own claim: ambiguity in a design makes people engage more deeply.`
- Gaver runs a thought experiment on a claim from Gaver, Beaver and Benford 2003 (p. 940); he does not test it.
- Replacement: `Gaver takes his own 2003 claim: ambiguity in a design makes people engage more deeply.`

**S6. Overstatement of what the manifestos did.** `03-Two-Tests.md`, slide 03.2.
- On the slide, current: `Reflective Design, Supple Interfaces, Designing for the Self: none can be refuted. Each inspires new designs.` Replacement: `Reflective Design, Supple Interfaces, Designing for the Self: none can be refuted. What counts is their power to inspire new designs.`
- Say, current: `None is falsifiable. Each one inspired new designs.` Replacement: `None is falsifiable. What matters, he says, is their ability to inspire new designs.`
- Paper p. 941: "what is at stake is their ability to inspire new designs."

**S7. Dropped hedge, and the adjectives credited to the wrong test.** `03-Two-Tests.md`, slide 03.2.
- On the slide, current: `**Verdict: passes Lakatos.** Design theory is "provisional, contingent, and aspirational": theory that is "sometimes right."` Replacement: `**Verdict: passes Lakatos.** Design theory tends to be "provisional, contingent, and aspirational": theory that is "sometimes right."`
- Say, current: `That gives Gaver his three adjectives: design theory is provisional, contingent and aspirational.` Replacement: `Put the two tests together and Gaver gets his three adjectives: design theory tends to be provisional, contingent and aspirational.`
- Paper: "likely to produce" (p. 937), "tends to be" (p. 938), and on p. 941 the adjectives follow from theory that "tends to underspecify practice and to be generative," that is, from the Popper analysis, not from Lakatos alone.

**S8. Dropped hedge on "most of us", and a repeated lead-in.** `04-Puzzle.md`, slide 04.1, Say.
- Current: `And Gaver adds that most of us already share values: contact with the people we design for, a wide exploration of options, craft, and making as a route to discovery.`
- The file's own Watch out asks for the "seem" hedge (p. 942), and the previous sentence also begins "Gaver adds".
- Replacement: `He also suggests that most of us seem to share values already: contact with the people we design for, a wide exploration of options, craft, and making as a route to discovery.`

**S9. Vote question shifts the object.** `04-Puzzle.md`, slide 04.1, On the slide, last bullet, and Say.
- Current: `**Vote: which is HCI as you have met it?**` and `Which of these is HCI as you have met it?`
- Gaver's three readings are about research through design, not HCI as a whole; a voter thinking of lab studies will pick A for a different field.
- Replacement: `**Vote: which is design research in HCI, as you have met it?**` and `Which of these is design research in HCI, as you have met it?`

**S10. Hedge dropped and a claim added in an answer.** `04-Puzzle.md`, If asked, "Can natural science ever be many worlds?"
- Current: `That presumption, not the method, is what drives convergence.`
- Paper p. 943: "This core belief seems naturally to lead to convergence." "Not the method" is ours.
- Replacement: `That presumption, he says, "seems naturally to lead to convergence."`

**S11. Rams exercise: column three has no hands prompt, and one long sentence.** `05-Rams.md`, slide 05.1, Say.
- Current: `Column three: the products with the principles beside them, each note pointing at a detail. And now? [pause]` Replacement: `Column three: the products with the principles beside them, each note pointing at a detail. Hands if yes now. [pause]`
- Current: `As we read it, here is a form you can extend, by designing the next Rams product, and verify, by checking the note against the object, without pretending the principles are laws.` Replacement: `As we read it, you can extend this form by designing the next Rams product. You can verify it by checking a note against the object. And nobody has to pretend the principles are laws.`

**S12. Dropped hedge on patterns.** `05-Rams.md`, slide 05.2, On the slide, bullet 5.
- Current: `Opposite of a pattern language: patterns drop the example, portfolios keep it.`
- Paper p. 945: "in several respects, the converse". v1 `06-Annotation.md` notes each Alexander pattern carries an archetypal photograph, so "drop the example" is too strong.
- Replacement: `In several respects the converse of a pattern language: patterns abstract across examples, portfolios keep each one particular.`

**S13. "Theory is the caption" against Gaver's balance rule.** `05-Rams.md`, slide 05.2, Watch out (add a line).
- The headline and bullet 2 make theory subordinate; Gaver says "neither dominates, and neither is subservient" (p. 945), and 05.1's Say says "neither part in charge". The label is ours (00-Overview) but the slide does not say so.
- Add to Watch out: `"Caption" is our word. Gaver insists that in a portfolio "neither dominates, and neither is subservient" (p. 945); if a classmate says "caption" makes theory subservient, agree and give his phrase "annotation".`

**S14. Lane label caveat from v1 lost.** `06-Aftermath.md`, slide 06.1, Say and Watch out.
- Current Say: `But the critics' lane kept moving too.` Höök and Löwgren, and Dalsgaard and Dindler, were not among the critics, and Höök and Löwgren 2012 cite the essay. v1 `10-What-Followed.md` carried this caveat; slides2 dropped it.
- Replacement Say: `But the other lane, toward knowledge above single designs, kept moving too.`
- Add to Watch out: `"The critics' way" is our lane label, not a claim that these authors opposed Gaver; Höök and Löwgren 2012 cite the essay.`

**S15. Unmarked opinion on the slide.** `06-Aftermath.md`, slide 06.1, On the slide, Twist bullet. Covered by the replacement in B2.

**S16. Logic of the Shin card.** `07-Connections.md`, slide 07.1, On the slide, bullet 4, and Say.
- Current slide: `**Shin et al., 2019:** a four-hour realistic test, so "inconceivable" was too strong. Yet participants guessed the purpose.` Gaver's word applies to tests "simultaneously unambiguous and ecologically valid" (p. 940); Shin et al. met only the realistic half, so the study does not show "inconceivable" was too strong. v1 said "half right, not refuted."
- Replacement slide: `**Shin et al., 2019:** a four-hour realistic test, but participants guessed the purpose. Realistic, not unambiguous: Gaver half right.`
- Current Say: `That makes Gaver's claim that such tests are inconceivable look too strong.` Replacement: `That achieves the realistic half of what Gaver called inconceivable.`

**S17. Outside quotation reworded.** `01-Hook.md`, If asked, "Who cites it as an exemplar?"
- Current: `Zimmerman and Forlizzi in 2014 call it "one of the earliest and best examples of RtD published within HCI"`
- v1 `10-What-Followed.md` log: the text says the Drift Table "has been held up as one of the earliest and best examples".
- Replacement: `Zimmerman and Forlizzi in 2014 write that it "has been held up as one of the earliest and best examples of RtD published within HCI"`

**S18. Method line covers only part 1 of the essay.** `02-Fight.md`, slide 02.2, On the slide, last bullet.
- Current: `How the essay works: no study, no data. Two borrowed tests.`
- Parts 2 and 3 work differently: three readings of one observation, then one sketched example (v1 `07-Method.md`). See Coverage, row 3.
- Replacement: `How the essay works: no study, no data. Two borrowed tests, three readings, one sketched example.`

## Nice to have

1. `01-Hook.md` Say: `Nobody in this fight thinks the table is bad work.` Only Zimmerman's group is shown praising it. Replace with `The critics themselves hold the table up as good work.`
2. `08-Verdict.md` slide bullet 2 and Say: "funding" and "examining" are our extension (v1 log: these words are absent from the essay). Anchor it in the paper: add to Watch out `Gaver himself says integration "would make assessment of new contributions fairer and more rigorous" (p. 941) and never returns to it.`
3. `03-Two-Tests.md` Say: `His own ludic design fits the same pattern.` Mark it: `We would add his own ludic design, which fits the same pattern.`
4. `03-Two-Tests.md` Say: `His 2003 paper claimed` becomes `His 2003 paper with Beaver and Benford claimed`.
5. `03-Two-Tests.md` Say: `His logical point is about the claim with a bare "sometimes" in it.` is awkward aloud. Replace with `His logical point only covers a bare "sometimes" claim.`
6. `04-Puzzle.md` Visual line says the iPad sentence is "the only quotation", but the slide also quotes "must founder" and "most of us". Align the Visual line.
7. `03-Two-Tests.md` Watch out and Sources: cite "provisional, contingent, and aspirational" as p. 937 and p. 941 (exact), with p. 938 as the variant without the serial comma. Also `02-Fight.md` slide 02.2: Gaver does say where he leans ("Personally, I tend to think that the nature of design practice and research is distinct from the ideals, and often practice, of science", p. 941); worth one line in If asked.

## Coverage of the brief

| Brief item | Slides | Judgement |
|---|---|---|
| Problem and why it matters | 01.1, 02.1 | Sufficient. The object, both sides in their words, and a "why you care" line. |
| Main idea or contribution | 02.2, 03.2, 05.1, 05.2 | Sufficient, spread across slides. Gaver's closing stance ("wary of impulses towards convergence and standardisation", p. 937; "consensus can come to look like constraint", p. 946) is only in notes. |
| How the research was conducted | 02.2 (bullet and Say) | Thin. "Two borrowed tests" describes part 1 only; see S18. |
| Most important takeaways | 03.1, 03.2, 04.1, 05.2 verdict lines; 08.1 | Partly. Each part has a verdict line, but no slide gathers Gaver's three conclusions (moderate expectations; norms are political acts; theory as annotation, p. 945). Suggest one line in the 06.1 or 08.1 Say: `Gaver's three conclusions: expect provisional theory, treat standard-setting as politics, and annotate portfolios.` |
| Strengths and limitations | 08.1; also 03.1 "We argue", 03.2 pushback, 05.1 Rams caution | Sufficient for 30 minutes. Strengths get one line; limitations get three pushbacks with Gaver's best replies in the notes. |
| Connections to course readings | 07.1 | Sufficient. Four readings, each link marked as ours. |
| New questions or directions | 06.1 (two open questions), 08.1, 09.1 | Sufficient. |

## Audience moments

- **01.1, hands.** Clear. Watch out gives one hand per person, three counts on a card, and follow-ups for each answer. Add `[pause]` after `Depends what counts.` in the Say so the presenter waits.
- **04.1, A/B/C vote.** Mostly clear: the Say calls A, B, C in turn and Watch out covers counting and lopsided results. Two issues: the question names HCI rather than design research (S9), and at 277 spoken words in 2:30 the vote itself relies on the overall buffer; the Open item already names what to cut.
- **05.1, three Rams columns.** Workable with the [click] and [pause] cues, but the Say answers columns one and two itself ("Probably not."; "You would get the surface, maybe."). Column three has no hands prompt, so the room cannot see its own shift (S11). The Open item says to let the room answer first; move that into Watch out so it is read at the podium.

## Transitions

| From | Transition line | Next opening | Fit |
|---|---|---|---|
| 01.1 | So who is fighting whom? | 02.1: "Between 2007 and 2011 a group of design researchers ..." | Good |
| 02.1 | His first move is not to defend design. It is to attack the word "science." | 02.2: "This is an essay. There is no study ..." | Weak: "science" arrives five sentences later; "attack" overstates (S4) |
| 02.2 | Test one. Could you prove a design rule wrong? | 03.1: "Take any design rule." | Good |
| 03.1 | If you cannot refute it, what is it good for? Test two. | 03.2: "Lakatos asks a different question." | Good |
| 03.2 | Two tests, two verdicts. Now the puzzle the critics built their case on. | 04.1: "Here is the observation the critics start from." | Good (Gaver: "seemingly behind recent calls", p. 942) |
| 04.1 | So if the field will not converge, what is theory for? Try this. | 05.1: "Gaver's proposal is best met as an experience, so try this." | Fits; "try this" said twice in a row, drop it from one |
| 05.1 | What this quietly flips is the order of theory and object in HCI. | 05.2: "Normally in HCI the theory comes first ..." | Good |
| 05.2 | That was 2012. Did he win? | 06.1: "Partly." | Good |
| 06.1 | Before our verdict, four places where this paper touches what we already read. | 07.1: "Gaver cites none of our course readings ..." | Good |
| 07.1 | Our verdict, in three sentences we would say to his face. | 08.1: "Here is our reading." | Good |
| 08.1 | Which brings us back to the table. | 09.1: "Back to the table." | Fits; repetitive, acceptable |

## Style

- Em-dashes and en-dashes: none in files 00 to 09 (checker and grep).
- On-slide bullets over about 25 words: none.
- Slides over about 70 on-slide words: 04.1 (76), 06.1 (76), 07.1 (79); 08.1 is at 70. The overview accepts the first three.
- Hard to read aloud: 05.1 "As we read it, here is a form ..." (S11); 02.1 "Second, and this is the sharper point, ..." (S1); 04.1 two consecutive "Gaver adds" (S8); 03.1 "the claim with a bare 'sometimes' in it" (nice to have 5).

## New facts (not in the paper or slides/)

| Item | File | Marked in log |
|---|---|---|
| Drift Table details: "about a terabyte of photography of England and Wales", "apparent height, direction and speed", nine authors, Gaver at the RCA in 2004, Bowers at KTH | 01 | Yes (checked against CHI 2004 paper) |
| Lakatos "Hungarian" | 02 | Yes (NEW) |
| "Accidents of 1960s manufacturing" | 05 | Yes (flagged as ours) |
| Discussion plan, role cards, grading from `progress.md` | 09 | Yes |
| "reviewed by people who did not know how to judge it" | 02 | No (S2) |
| Portfolios became a "method with its own conference" | 06 | No (B2) |
| "Nobody in this fight thinks the table is bad work" | 01 | No (nice to have 1) |

## Passed

Status "Exact" means wording matches `paper-text.txt` (ellipses as marked) or the named v1 file.

| File | Quotation | Page claimed | Page found | Status |
|---|---|---|---|---|
| 01 | "a well known example of critical design in HCI" | ZFE 2007 p. 496 | v1 01 | Exact |
| 01 | "one of the earliest and best examples of RtD published within HCI" | ZF 2014 p. 175 | v1 10 | Exact fragment; verb wrong (S17) |
| 01 | "motivated by curiosity, exploration, and reflection rather than externally-defined tasks" | CHI 2004 | v1 01 | Exact |
| 01 | "about a terabyte of photography of England and Wales" | CHI 2004 p. 886 | new, logged | Not in v1 |
| 01 | "apparent height, direction and speed" | CHI 2004 p. 886 | new, logged | Not in v1 |
| 02 | "lacks clear expectations and standards" | 937 | 937 | Exact |
| 02 | "actionable metrics" | 937 | 937 | Exact |
| 02 | "develop protocols, descriptions, and guidelines" | 937 | 937 | Exact |
| 02 | "a level of extensibility and verifiability" | 937 | 937 | Exact |
| 02 | "some sympathy" / "some sympathy with these concerns" | 937 | 937 | Exact |
| 02 | "self-policing" | 937 | 937 | Exact |
| 02 | "prone to 'scientism'" | 938 | 938 | Exact; hedge dropped (S1) |
| 02 | "an elephant in the room" | 946 | 946 | Exact |
| 02 | "the anonymous R4, who recognised ... 'scientism' in HCI'" | 946 | 946 | Exact |
| 02 | "It is difficult to call for research to be more scientific if what it means to be scientific is under question." | 941 | 941 | Exact |
| 02 | "complicate notions of what makes theories 'scientific'" | 938 | 938 | Exact |
| 02 | "the criterion of the scientific status of a theory is its falsifiability" | 939 | 939 | Exact |
| 02 | "particularly representative of Philosophy of Science" / "not ... particularly representative" | 941 | 941 | Exact |
| 02 | "on its own terms" | 945 | 945 | Exact |
| 02 | "hard core", "protective belt" | 939 to 940 | 939 | Exact |
| 02 | "predict novel facts" | 939 to 940 | 940 | Exact |
| 02 | "Must design become 'scientific'?" | 941 (citation) | 941 cite, 946 ref | Exact |
| 02 | "Quality control"; "Science: Conjectures and Refutations"; dates of Popper and Lakatos | ref list; v1 04 | 946; v1 04 | Exact |
| 03 | "X will sometimes lead to successful outcomes" | 940 | 940 | Exact |
| 03 | "there is always an implicit sometimes ..." | 940 | 940 | Exact |
| 03 | "Simply tweaking an unambiguous system to be ambiguous" | 940 | 940 | Exact |
| 03 | "tests which are simultaneously unambiguous and ecologically valid are inconceivable" | 940 | 940 | Exact |
| 03 | "in principle a falsifiable statement" | 940 | 940 | Exact |
| 03 | "theory is underspecified by design" | 944 | 944 | Exact |
| 03 | "potentially unflattering" mirror | 940 | 940 | Exact |
| 03 | "reveals two important characteristics" | 940 | 940 | Exact |
| 03 | "Underspecification makes falsification difficult in practice as well as in principle" | 940 | 940 | Exact |
| 03 | "is widely criticised both in Philosophy of Science and by practicing scientists" | 940 | 940 | Exact |
| 03 | "is considered the fault of the scientist using the theory, not the fault of the theory itself" | 939 | 939 | Exact |
| 03 | "irresistables" | 938 | 938 | Exact |
| 03 | "'Reflective Design' [30] or 'Supple Interfaces' [16] or 'Designing for the Self' [36]" | 941 | 941 | Exact |
| 03 | "'merely' confirmatory" | 941 | 941 | Exact |
| 03 | "testimony to the fertility of the overarching theory" | 941 | 941 | Exact; misattributed on slide (B1) |
| 03 | "provisional, contingent, and aspirational" | 937, 938 | 937, 941 (938 without serial comma) | Exact |
| 03 | "theories that are sometimes right" / "sometimes right" | 941 | 941 | Exact |
| 03 | "depends on the account of science one adopts" | 941 | 941 | Exact |
| 03 | "relative productivity of competing programmes" | 941 | 941 | Exact |
| 03 | "can be prepared-for-action but not guided-in-action by detailed prescriptive procedures" | 941 | 941 | Exact (extraction hyphen) |
| 03 | 2005 method, 216 studies in a 2020 review; ambiguity paper's three kinds, pp. 233 to 240; "degenerating" from SEP | v1 09, v1 04 | v1 09, v1 04 | Faithful |
| 04 | "must founder" | 942 | 942 | Exact |
| 04 | "most of us" / "most of us ... seem to share" | 942 | 942 | Exact |
| 04 | "Perhaps we do not all agree with all of these assumptions all of the time." | 942 | 942 | Exact |
| 04 | "for a different world" / "anybody seeking ... one in which the iPad exists." | 943 | 943 | Exact |
| 04 | "potentially repressive acts of ontological politics" | 943 | 943 | Exact |
| 04 | "the other two imply either that there is no problem to solve or ... rather than agreement." | 938 | 938 | Exact |
| 04 | "there is little sense of priority or order to the findings that are amassed." | 941 to 942 | 942 | Exact |
| 04 | "may well exhibit just as much uncertainty, interpretation and debate as does research through design" | 942 | 942 | Exact |
| 04 | Law and Urry "different research practices might be making multiple worlds ..." | 943 | 943 | Exact |
| 04 | "wary of impulses towards convergence and standardisation" | 937 | 937 | Exact |
| 04 | "a strong presumption that the object of study is a single, unitary world ..." | 943 | 943 | Exact |
| 04 | "dozens of different 'manifestos'"; "Why is there not more convergence"; "three possibilities" | 941 | 941 | Exact |
| 04 | "pre-paradigmatic"; "performative" | 941; 943 | 941; 943 | Exact |
| 05 | "in the style of Dieter Rams"; "one can imagine 'designing in the style of Dieter Rams' without the need for a detailed formal theory" | 944 | 944 | Exact |
| 05 | "Good design is unobtrusive." "Good design is thorough down to the last detail." | 944 | 944 | Exact (capitalised) |
| 05 | "ten concise principles"; "for Braun and Vitsoe" | 944 | 944 | Exact |
| 05 | "too terse to serve as substantial conceptual contributions by themselves" | 944 | 944 | Exact |
| 05 | "neither dominates, and neither is subservient" | 945 | 945 | Exact |
| 05 | "annotated portfolios are not defined by their graphic presentation"; "Sketch design for an annotated portfolio" | 945 | 945 | Exact |
| 05 | "not a formal one"; "not how materials are presented" | 944 | 944 | Exact |
| 05 | "extensibility and verifiability" ... "for which some of the HCI design community have called" | 945 | 945 | Exact |
| 05 | "emphasises his design principles over the portfolio of products themselves" | 945 | 945 | Exact |
| 05 | "the definite facts of research through design" | 945 | 945 | Exact |
| 05 | "Theories may be provisional, but designed artefacts (as opposed to demonstrators, for example) are not" | 944 | 944 | Exact |
| 05 | "theory nexus"; "ultimate particular"; "the same dignity and importance as truth in science" | 944 | 944 | Exact |
| 05 | "If artefacts embody theory, however, they do not encode it" | 944 | 944 | Exact |
| 05 | "limited to inspiration and annotation" | 945 | 945 | Exact |
| 05 | "the converse of Alexander's [1] design patterns"; "abstract regularities from repeated attempts"; "maintain the particularity of individual examples" | 945 | 945 | Exact |
| 05 | "point in design space ... an area in that space"; "the invariances among them" | 944 | 944 | Exact |
| 05 | "I am not proposing here that annotated portfolios subsume ... particularly for an interdisciplinary audience" | 945 | 945 | Exact |
| 05 | "it is useful to consider specific designs such as a rotary controller for a microwave oven" | 944 | 944 | Exact |
| 05 | "turn to existing examples" | 943 | 943 | Exact |
| 05 | Figure 1: seven photographs, four principles (2, 4, 5, 8), prototype TV; Vitsoe's ten; Lubs calculator; "a million times over" | v1 06 | v1 06 | Faithful |
| 06 | "scientism"; "wary of impulses"; "definite nature of designed artefacts"; "a serious form of theoretical contribution"; "discursiveness and elaboration" | 938; 937; 943; 945; 946 | same | Exact |
| 06 | "Annotated Portfolios"; "intermediate-level knowledge"; "alternatives to scientistic ways of knowing"; "worked to formalize" (p. 178); "hundreds and thousands"; "towards rapid convergence on a limited set of ideas"; "contestable, defensible, and substantive"; "may capture best what we can know through design"; "RTD Roams"; CHI 2026 counts; OpenAlex about 1,100 | v1 10 | v1 10 | Faithful |
| 07 | "inconceivable"; "establishes an area"; "neither dominates"; "wary" | 940; 944; 945; 937 | same | Exact |
| 07 | "one user interface metaphor" (p. 240); "confirmatory evidence" (p. 113); "ecological validity" (p. 5); "the space between"; "embrace giving definition" (p. 44); 216 studies; 11 of 12 guessed | v1 09 | v1 09 | Faithful |
| 08 | "clearly articulating ... the standards one might want to advocate"; "productive research programmes"; "the relative productivity of competing programmes as the grounds for scientists' choice amongst them" | 941 | 941 | Exact |
| 08 | "political acts to be approached with care"; "on its own terms"; "a serious form" | 945 | 945 | Exact |
| 08 | "ontological politics" | 943 | 943 | Exact |
| 08 | "actionable metrics"; "some sympathy" | 937 | 937 | Exact |
| 08 | "depends on the account of science one adopts"; "It is difficult to call ..." | 941 | 941 | Exact |
| 08 | "Perhaps we do not all agree ..."; "most of us" | 942 | 942 | Exact |
| 08 | Pullin and Schön listed, never cited | 946 | 946 (no in-text [25] or [28]) | Confirmed |
