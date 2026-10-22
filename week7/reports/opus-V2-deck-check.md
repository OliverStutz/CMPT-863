# V2 deck check: slides2/01-Hook.md to 09-Handover.md

Checked 2026-10-21 against `assets/paper-text.txt` (page breaks at the bare lines 937 to 946) and the v1 files in `slides/`. Files were being edited during the check (01, 05, 09 changed at 17:45); this report reflects the state at 17:47, snapshot in the session scratchpad.

## Summary

- Quotations checked: 120 rows (96 from the paper, 24 from outside works, course files or the CHI 2004 paper).
- Exact or faithful: 113. Inexact: 5 (one misattributed, two dropped inner quotation marks, one singular for plural, one spelling). New and not in v1: 2 (CHI 2004, marked in the v2 log).
- Pages: 96 checked against the paper, 96 right, 0 wrong.
- New facts (in v2 but in neither v1 nor the paper): 8 items; 4 marked, 1 partly marked, 3 unmarked.
- Transitions: 11 of 11 lead into the next slide.
- Dashes: no em-dash or en-dash in any file. Bullets over 25 words: 2.

## Blockers

1. **02-Fight.md, slide 02.1, On the slide, bullet 1.** The quotation "extensible and verifiable" is attributed to the critics. Those exact words are Gaver's own (p. 938, "instead of being extensible and verifiable"). The critics' words, which Gaver quotes from [33] on p. 937, are "a level of extensibility and verifiability."
   - Current: `Give it protocols, guidelines, and theory that is "extensible and verifiable."`
   - Replace: `Give it protocols, guidelines, and theory with "a level of extensibility and verifiability."`

2. **03-Two-Tests.md, slide 03.2, Say.** The script says Gaver names ludic design in the list. He does not: the p. 941 list is Reflective Design, Supple Interfaces and Designing for the Self. Your own Watch out admits the swap, but the Say still says Gaver made the list.
   - Current: `Gaver lines up three design manifestos: Sengers's Reflective Design, Isbister and Höök's Supple Interfaces, and his own ludic design.`
   - Replace: `Gaver lines up three design manifestos: Sengers's Reflective Design, Isbister and Höök's Supple Interfaces, and Zimmerman's Designing for the Self. His own ludic design fits the same pattern.`
   - Also change the on-slide bullet and the card visual to match (see Should fix 6).

3. **05-Rams.md, slide 05.1, Say, last sentence.** It presents our limitation as the paper's own caution. The paper calls Figure 1 a "Sketch design" and says portfolios are "not defined by their graphic presentation." It never says Rams is "not a design research project." Your own If asked says "Gaver does not say."
   - Current: `One caution from the paper itself: this is a sketch of Rams, not a design research project, and Gaver says the form is not defined by how it is laid out.`
   - Replace: `Two cautions. Gaver calls the figure a sketch and says the form is not defined by how it is laid out. And we would add: Rams is a product designer, not a design research project.`

4. **07-Connections.md, slide 07.1, On the slide bullet 3 and Say; also 03-Two-Tests.md, slide 03.2, If asked, third Q.** "Over two hundred follow-up studies" reads as studies built on the 2009 paper. v1 (`slides/09-Connections.md`, Watch out and If asked) explicitly says never to put it that way. The 216 count is the number of elicitation studies in a 2020 review, and the method dates from 2005.
   - Current (slide): `**Wobbrock, Morris and Wilson, 2009:** over two hundred follow-up studies. Fertile, and still science.`
   - Replace: `**Wobbrock, Morris and Wilson, 2009:** a method a 2020 review traces through 216 studies. Fertile, and still science.`
   - Current (Say): `That method spawned over two hundred follow-up studies by 2020.`
   - Replace: `By 2020 a review counted over two hundred studies that used that method, which dates from 2005.`
   - Current (03 If asked): `the gesture-elicitation method behind Wobbrock, Morris and Wilson 2009 spawned over two hundred follow-up studies.`
   - Replace: `the gesture-elicitation method used by Wobbrock, Morris and Wilson 2009 appears in over two hundred studies counted by a 2020 review.`

5. **07-Connections.md, slide 07.1, If asked, "What did Shin et al. do?"** The answer contradicts v1, which says the study was realistic but not unambiguous: actuation always came on day 2, so day and treatment are confounded, and 11 of 12 participants guessed the purpose.
   - Current: `Ecologically valid and controlled at once, which is what Gaver called inconceivable. But nearly every participant guessed the purpose, which is the kind of contamination Gaver's argument predicts.`
   - Replace: `Realistic, yes, which is half of what Gaver called inconceivable. Unambiguous, no: actuation always came on day 2, and 11 of 12 participants guessed the purpose (their p. 7), which is the kind of contamination Gaver's argument predicts.`

## Should fix

1. **02-Fight.md, 02.1, Say.** This states a motive as fact. Gaver says conformance "may be a route towards disciplinary legitimacy" (p. 937).
   - Current: `Their motive was legitimacy.`
   - Replace: `Gaver reads it as a bid for legitimacy.`
2. **02-Fight.md, 02.1, On the slide and Say.** Gaver's hedge ("There is a risk, however, that such standards might lead to a form of self-policing," p. 937) is dropped.
   - Current (slide): `**Gaver, 2012:** that way lies "self-policing," and HCI is "prone to scientism."`
   - Replace: `**Gaver, 2012:** standards risk "self-policing," and HCI is "prone to 'scientism.'"`
   - Current (Say): `First, standards become self-policing: the community restricts itself.`
   - Replace: `First, standards might become self-policing: the community restricts itself.`
3. **02-Fight.md, 02.1, Say.** Gaver's hedge is dropped. He says standards "may adopt, or be interpreted in terms of, inappropriate 'scientific' models" (p. 938).
   - Current: `so any standards it adopts will quietly be the standards of experimental science, and those do not fit design.`
   - Replace: `so any standards it adopts may quietly become the standards of experimental science, which he thinks do not fit design.`
4. **02-Fight.md, 02.1, If asked, second Q.** This overstates p. 937.
   - Current: `He calls their concerns ones he shares in part (p. 937)`
   - Replace: `He says he has "some sympathy with these concerns" (p. 937)`
5. **03-Two-Tests.md, 03.1, On the slide bullet 3 and Say.** Both overstate p. 940, where Gaver says simple tweaking "would be unrepresentative of the ways designers harness untheorised factors."
   - Current (slide): `"Simply tweaking" one is not real design.`
   - Replace: `"Simply tweaking" one misrepresents how designers work.`
   - Current (Say): `Simply tweaking a finished design to add ambiguity would give you something no designer would ever make.`
   - Replace: `Simply tweaking a finished design to add ambiguity would not represent how designers actually work.`
6. **03-Two-Tests.md, 03.2, On the slide bullet 1 and Visual.** To match p. 941 (see Blocker 2), name the manifestos Gaver lists, or label ludic design as ours.
   - Current: `Reflective Design, Supple Interfaces, ludic design: none can be refuted. Each produced new designs.`
   - Replace: `Reflective Design, Supple Interfaces, Designing for the Self: none can be refuted. Each inspires new designs.`
7. **03-Two-Tests.md, 03.2, On the slide bullet 3.** "often useful" is not Gaver's wording, but it sits inside his verdict line.
   - Current: `"provisional, contingent, and aspirational": sometimes right, often useful.`
   - Replace: `"provisional, contingent, and aspirational": theory that is "sometimes right."`
8. **04-Puzzle.md, 04.1, On the slide, card B.** Both hedges are dropped: Latour's science "may well exhibit just as much uncertainty," and "most of us ... seem to share" (p. 942).
   - Current: `**B. Messier up close.** Real science is a brawl too; "most of us" already agree. Fix: nothing.`
   - Replace: `**B. Messier up close.** Science may be just as contested day to day; "most of us" seem to share core values. Fix: nothing.`
9. **04-Puzzle.md, 04.1, Say.** Three places overstate the paper.
   - Current: `everyone argues from first principles, nothing accumulates` / Replace: `everyone argues from first principles, and findings do not add up` (p. 942: "little sense of priority or order," "Progress is slow").
   - Current: `Latour's point: real science is full of controversy day to day; the unity you see is hindsight.` / Replace: `Latour's point: real science is full of controversy day to day. Gaver adds that the unity we see may come from popular accounts and hindsight.` The hindsight point is Gaver's, and it is hedged ("may be a function of," p. 942).
   - Current: `Science converges because there is one world to be right about.` / Replace: `Science presumes one world to be right about, and that, Gaver says, seems to drive convergence.` (p. 943: "strong presumption," "seems naturally to lead").
10. **05-Rams.md, 05.1, Say.** The hedge is dropped. Gaver says the principles "may seem too terse to serve as substantial conceptual contributions by themselves" (p. 944).
    - Current: `They are too terse to be a theory, and Gaver says so himself.`
    - Replace: `Gaver grants they may seem too terse to stand on their own.`
11. **05-Rams.md, 05.1, Say.** Our gloss runs on from "his answer" as if it were Gaver's. The Watch out marks it as ours; the Say should too.
    - Current: `Here is a form you can extend, by designing the next Rams product,`
    - Replace: `As we read it, here is a form you can extend, by designing the next Rams product,`
12. **06-Aftermath.md, 06.1, Say.** Your own Watch out says the two words are "close, not identical."
    - Current: `That is a scope statement, not reviewing practice, but it is his word.`
    - Replace: `That is a scope statement, not reviewing practice, but it is close to his word.`
13. **06-Aftermath.md, 06.1, Say.** v1 says the framework "speaks to" the 2010 call.
    - Current: `which is the 2010 call for protocols, answered.`
    - Replace: `which speaks to the 2010 call for protocols.`
14. **06-Aftermath.md, 06.1, If asked, last Q.** v1 (`slides/01-Background.md`, Watch out) says to use "the grant that funded this essay," not a phrase that names the holder. The paper's term is "Advanced Investigator Award."
    - Current: `Gaver's studio had a European Research Council Advanced grant (p. 946).`
    - Replace: `The essay was funded by a European Research Council Advanced Investigator Award (p. 946).`
15. **07-Connections.md, 07.1, Say.** v1 and the v2 Watch out both say "for design theory."
    - Current: `Gaver drops refutation entirely.`
    - Replace: `Gaver drops refutation as the test for design theory.`
16. **08-Verdict.md, 08.1, Say.** This contradicts the file's own Watch out, which notes Gaver cites Sengers 2006, so he was not first.
    - Current: `He names the elephant when others talked around it.`
    - Replace: `He names the elephant in a CHI paper, following Sengers.`
17. **08-Verdict.md, Watch out, "Gaver's best replies".** "Should be slow" is not in the paper. Gaver says norms are "political acts to be approached with care" (p. 945).
    - Current: `that standard-setting is "ontological politics" and should be slow (p. 943, p. 945)`
    - Replace: `that standard-setting is "ontological politics" and should be "approached with care" (p. 943, p. 945)`
18. **08-Verdict.md, Sources.** This lists our inference as paper content.
    - Current: `"neither dominates"; Rams not RtD)`
    - Replace: `"neither dominates"; Figure 1 caption, from which we infer Rams is not RtD)`
19. **01-Hook.md, 01.1, Say.** This contradicts the updated Watch out ("Say 'Gaver and colleagues,' not 'the studio'").
    - Current: `and the studio wrote up what people did with it.`
    - Replace: `and the team wrote up what people did with it.`
20. **02-Fight.md, 02.2, If asked, Popper.** "Then London" is NEW and leaves out New Zealand (1937 to 1945). v1 has only "born in Vienna."
    - Current: `Karl Popper (1902 to 1994), Vienna then London,`
    - Replace: `Karl Popper (1902 to 1994), born in Vienna,`
21. **Bullet length.** Two on-slide bullets are over about 25 words.
    - 01-Hook.md, 01.1, bullet 1 (34 words). Current: `The Drift Table, 2004: a coffee table with a porthole that shows aerial photographs. Lean on one side and the landscape drifts that way. No task, no manual. People lived with it for weeks.` / Replace: `The Drift Table, 2004: aerial photographs drift in a porthole as you shift weight. No task, no manual. Lived with for weeks.`
    - 02-Fight.md, 02.1, bullet 1 (32 words): move the names in parentheses into the Say, leaving `**The critics, 2007 to 2011:** design research "lacks clear expectations and standards." Give it protocols, guidelines, and theory with "a level of extensibility and verifiability."`

## Nice to have

- 05-Rams.md, 05.2, On the slide: `the artefact is the "definite fact."` The paper's phrase is plural ("the definite facts," p. 945). Use `the artefacts are the "definite facts."`
- 03-Two-Tests.md, 03.1, If asked: "irresistibles" is in quotation marks, but the paper's text spells it 'irresistables' (p. 938); only the reference title has "irresistibles." Drop the quotation marks or add [sic].
- 03-Two-Tests.md, 03.2: "merely confirmatory" drops Gaver's inner quotation marks ("'merely' confirmatory," p. 941). 02.1 slide: "prone to scientism" likewise (fixed by Should fix 2).
- 03-Two-Tests.md, 03.2, Say: `merely confirmatory, worthless as evidence` goes further than Popper's "merely." Use `merely confirmatory, weak as evidence`.
- 02-Fight.md, 02.2, title: `"You want us scientific? Which science?"` is in quotation marks but is our line. Mark it as ours or drop the marks.
- 02-Fight.md, 02.2, If asked: "Kuhn comes in on the next question, convergence (p. 941 to 942)." Kuhn is also quoted on p. 939 (puzzle-solving). Say "mainly on the next question."
- 02-Fight.md, 02.1, If asked: the 2007 process criterion drops v1's caveat "with no expectation of the same result."
- 02-Fight.md, 02.2, If asked: "Hungarian" for Lakatos is NEW (true, harmless). Add it to the log.
- 05-Rams.md, 05.1, Say: "for decades" is NEW (v1 has only shelving 1957, Vitsoe 1959). Add it to the log or cut it.
- 06-Aftermath.md, 06.1, On the slide: "CHI 2026 Design track"; v1 and the notes say "subcommittee."
- 06-Aftermath.md, Say: "the definite, costly, particular object." "Costly" is ours; fine, but it is the premise of the AI question, so own it ("we would add, costly").
- 09-Handover.md, Watch out: "cited the talk with a slide." `progress.md` says "with a section." 09 Say: "one sentence from each of you." `progress.md` caps it at eight to ten voices.
- 08-Verdict.md, Say: "Most critiques of rigour do not." This is an unsupported generalisation; mark it as ours or cut it.
- Say scripts that would be unclear read aloud:
  - 03.1: "His logical point is about the bare form 'sometimes.'" Say "about the claim with a bare 'sometimes' in it."
  - 06.1: four author pairs in two sentences (Höök and Löwgren, Dalsgaard and Dindler, Bardzell and colleagues, Löwgren). Put the concept names first and the names second, or let the slide carry the names.
  - 08.1: "Your own list of what most of us share, contact with the people we design for, wide exploration, craft, making as discovery, already works as criteria." This is one long appositive. Split it: "Your own list of what most of us share: contact with users, wide exploration, craft, making as discovery. That list already works as criteria."
- 06.1 and 07.1 are flagged SLIDE DENSE by `_check-sections.py` (76 and 75 words).

## Transitions

| From | Transition line | Next slide opens with | Fit |
|---|---|---|---|
| 01.1 | "So who is fighting whom?" | 02.1 the two sides | OK |
| 02.1 | "...attack the word 'science.'" | 02.2 "which science?" | OK |
| 02.2 | "Test one. Could you prove a design rule wrong?" | 03.1 Popper test | OK |
| 03.1 | "...what is it good for? Test two." | 03.2 Lakatos | OK |
| 03.2 | "Now the puzzle the critics built their case on." | 04.1 "the observation the critics start from" | OK |
| 04.1 | "...what is theory for? Try this." | 05.1 "so try this" | OK |
| 05.1 | "What this quietly flips is the order..." | 05.2 "Normally in HCI the theory comes first" | OK |
| 05.2 | "That was 2012. Did he win?" | 06.1 "Partly." | OK |
| 06.1 | "Before our verdict, four places..." | 07.1 four links | OK |
| 07.1 | "Our verdict, in three sentences..." | 08.1 three pushbacks | OK |
| 08.1 | "Which brings us back to the table." | 09.1 "Back to the table." | OK |

Slide cross-references (slide 10 = 07.1, slide 11 = 08.1, slide 12 = 09.1) agree with the 12-slide numbering in 00-Overview.

## Audience moments

- **01.1 hands.** Clear enough to run. The Say gives the three calls, and Watch out says to count roughly and has follow-ups for "no" and "yes." Missing: what to do with "depends" voters, and where to record the split for slide 12. Add: "Call each option once, one hand per person; jot the three counts on a card."
- **04.1 vote.** Clear. The script gives "Hands up for A. For B. For C.", and Watch out covers a skewed room. Two gaps. First, the vote comes after about 230 words of exposition in a 2:30 slot, so any follow-up runs over; say which follow-up to drop if time is short. Second, the footer question ("which is HCI as you have met it?") is never spoken. Add it just before "Hands up for A."
- **05.1 Rams columns.** Not yet clear enough. The Visual says to reveal the columns one at a time, but the Say answers its own questions ("Probably not." "You would get the surface, maybe."), so the room never responds. Add explicit cues: "[click: column 1] Could you design in his style from these two sentences? Hands if yes. [pause] [click: column 2] ... Hands if yes. [click: column 3] And now?" Then let the scripted answers follow the hands.

## Passed

Status: E = exact; I = inexact (see the sections above). Pages are printed pages. "v1" means the claim was found in the named v1 file and reproduced faithfully.

| # | File | Quotation | Page claimed | Page found | Status |
|---|---|---|---|---|---|
| 1 | 01 | "a well known example of critical design in HCI" | Z 2007 p. 496 | v1 01 | E |
| 2 | 01 | "one of the earliest and best examples of RtD published within HCI" | Z&F 2014 p. 175 | v1 10 | E |
| 3 | 01 | "motivated by curiosity, exploration, and reflection rather than externally-defined tasks" | CHI 2004 | v1 01 | E |
| 4 | 01 | "about a terabyte of photography of England and Wales" | CHI 2004 p. 886 | NEW, v2 log | not in v1 |
| 5 | 01 | "apparent height, direction and speed" | CHI 2004 p. 886 | NEW, v2 log | not in v1 |
| 6 | 01 | "the definite facts of research through design" | 945 | 945 | E |
| 7 | 02 | "lacks clear expectations and standards" | 937 | 937 | E |
| 8 | 02 | "extensible and verifiable" (as the critics' words) | 937 | 938, Gaver's | I (Blocker 1) |
| 9 | 02 | "self-policing" | 937 | 937 | E |
| 10 | 02 | "prone to scientism" / "prone to 'scientism'" | 938 | 938 | I (inner marks dropped on slide) |
| 11 | 02 | "an elephant in the room" | 946 | 946 | E |
| 12 | 02 | "actionable metrics" | 937 | 937 | E |
| 13 | 02 | "some sympathy" | 937 | 937 | E |
| 14 | 02 | "develop protocols, descriptions, and guidelines" | 937 | 937 | E |
| 15 | 02 | "a level of extensibility and verifiability" | 937 | 937 | E |
| 16 | 02 | R4: "the anonymous R4, who recognised ... 'an elephant in the room...: the move towards "scientism" in HCI'" | 946 | 946 | E |
| 17 | 02 | "Must design become 'scientific'?" | 941 | 941 (cite), 946 (ref) | E |
| 18 | 02 | "It is difficult to call for research to be more scientific if what it means to be scientific is under question." | 941 | 941 | E |
| 19 | 02 | "not ... particularly representative of Philosophy of Science" | 941 | 941 | E |
| 20 | 02 | "complicate notions of what makes theories 'scientific'" | 938 | 938 | E |
| 21 | 02 | "on its own terms" | 945 | 945 | E |
| 22 | 02 | "the criterion of the scientific status of a theory is its falsifiability" | 939 | 939 | E |
| 23 | 02 | "hard core," "protective belt" | 939 to 940 | 939 | E |
| 24 | 02 | "predict novel facts" | 939 to 940 | 940 | E |
| 25 | 02 | 2007 criteria: process, invention, relevance, extensibility | Z 2007 pp. 499 to 500 | v1 01 | v1 (caveat dropped, Nice to have) |
| 26 | 03 | "X will sometimes lead to successful outcomes" | 940 | 940 | E |
| 27 | 03 | "there is always an implicit sometimes in statements about how to design successfully" | 940 | 940 | E |
| 28 | 03 | "Simply tweaking an unambiguous system to be ambiguous" | 940 | 940 | E |
| 29 | 03 | "tests which are simultaneously unambiguous and ecologically valid are inconceivable" | 940 | 940 | E |
| 30 | 03 | "in principle a falsifiable statement" | 940 | 940 | E |
| 31 | 03 | "theory is underspecified by design" | 944 | 944 | E |
| 32 | 03 | "potentially unflattering" | 940 | 940 | E |
| 33 | 03 | "reveals two important characteristics" | 940 | 940 | E |
| 34 | 03 | "Underspecification makes falsification difficult in practice as well as in principle" | 940 | 940 | E |
| 35 | 03 | "is widely criticised both in Philosophy of Science and by practicing scientists" | 940 | 940 | E |
| 36 | 03 | "is considered the fault of the scientist using the theory, not the fault of the theory itself" | 939 | 939 | E |
| 37 | 03 | "irresistibles" | 938 | 938, spelled 'irresistables' | I (spelling) |
| 38 | 03 | "merely confirmatory" | 941 | 941, "'merely' confirmatory" | I (inner marks dropped) |
| 39 | 03 | "testimony to the fertility of the overarching theory" | 941 | 941 | E |
| 40 | 03 | "What Popper would reject as 'merely' confirmatory evidence is ... testimony ..." | 941 | 941 | E |
| 41 | 03 | "provisional, contingent, and aspirational" | 937, 938 | 937 (also 938 without serial comma, 941) | E |
| 42 | 03 | "'Reflective Design' [30] or 'Supple Interfaces' [16] or 'Designing for the Self' [36]" | 941 | 941 | E |
| 43 | 03 | "theories that are sometimes right" | 941 | 941 | E |
| 44 | 03 | "depends on the account of science one adopts" | 941 | 941 | E |
| 45 | 03 | "relative productivity of competing programmes" | 941 | 941 | E |
| 46 | 03 | "novel facts" | 940 | 940 | E |
| 47 | 03 | Stolterman: "can be prepared-for-action but not guided-in-action by detailed prescriptive procedures" | 941 | 941 | E |
| 48 | 03 | Ambiguity paper: three kinds; CHI 2003 pp. 233 to 240 | outside | v1 04 | v1 |
| 49 | 03 | "degenerating" from Stanford Encyclopedia | outside | v1 04 | v1 |
| 50 | 04 | "dozens of different 'manifestos'" | 941 | 941 | E |
| 51 | 04 | "Why is there not more convergence" | 941 | 941 | E |
| 52 | 04 | "pre-paradigmatic" | 941 | 941 (heading) | E |
| 53 | 04 | "must founder" | 942 | 942 | E |
| 54 | 04 | "most of us ... seem to share" | 942 | 942 | E |
| 55 | 04 | "Perhaps we do not all agree with all of these assumptions all of the time." | 942 | 942 | E |
| 56 | 04 | "there is little sense of priority or order to the findings that are amassed" | 942 | 942 | E |
| 57 | 04 | "may well exhibit just as much uncertainty, interpretation and debate as does research through design" | 942 | 942 | E |
| 58 | 04 | "performative" | 943 | 943 | E |
| 59 | 04 | "anybody seeking to research or develop tablet computers ... is designing for a different world, one in which the iPad exists" | 943 | 943 | E |
| 60 | 04 | "potentially repressive acts of ontological politics" | 943 | 943 | E |
| 61 | 04 | Law and Urry: "different research practices might be making multiple worlds, and that such worlds might be equally valid, equally true, but simply unlike one another" | 943 | 943 | E |
| 62 | 04 | "a strong presumption that the object of study is a single, unitary world that pre-exists and is independent of its observers" | 943 | 943 | E |
| 63 | 04 | "the other two imply either that there is no problem to solve or that progress in this area will be marked by proliferation rather than agreement" | 938 | 938 | E |
| 64 | 04 | "wary of impulses towards convergence and standardisation" | 937 | 937 | E |
| 65 | 05 | "in the style of Dieter Rams" | 944 | 944 | E |
| 66 | 05 | "Good design is unobtrusive." | 944 | 944 | E |
| 67 | 05 | "Good design is thorough down to the last detail." | 944 | 944 | E |
| 68 | 05 | "ten concise principles" | 944 | 944 | E |
| 69 | 05 | "one can imagine 'designing in the style of Dieter Rams' without the need for a detailed formal theory" | 944 | 944 | E |
| 70 | 05 | "too terse to serve as substantial conceptual contributions by themselves" | 944 | 944 | E |
| 71 | 05 | "not a formal one"; "not how materials are presented" | 944 | 944 | E |
| 72 | 05 | "Sketch design for an annotated portfolio" | 945 | 945 | E |
| 73 | 05 | "annotated portfolios are not defined by their graphic presentation" | 945 | 945 | E |
| 74 | 05 | "neither dominates, and neither is subservient" | 945 | 945 | E |
| 75 | 05 | "for which some of the HCI design community have called" | 945 | 945 | E |
| 76 | 05 | "emphasises his design principles over the portfolio of products themselves" | 945 | 945 | E |
| 77 | 05 | "definite fact" (slide 05.2) | 945 | 945, "definite facts" | I (number) |
| 78 | 05 | "Theories may be provisional, but designed artefacts (as opposed to demonstrators, for example) are not" | 944 | 944 | E |
| 79 | 05 | "If artefacts embody theory, however, they do not encode it" | 944 | 944 | E |
| 80 | 05 | "limited to inspiration and annotation" | 945 | 945 | E |
| 81 | 05 | "the converse of Alexander's [1] design patterns" | 945 | 945 | E |
| 82 | 05 | "abstract regularities from repeated attempts"; "maintain the particularity of individual examples" | 945 | 945 | E |
| 83 | 05 | "point in design space ... an area in that space" | 944 | 944 | E |
| 84 | 05 | "the invariances among them" | 944 | 944 | E |
| 85 | 05 | "theory nexus" | 944 | 944 | E |
| 86 | 05 | "ultimate particular" | 944 | 944 | E |
| 87 | 05 | "the same dignity and importance as truth in science" | 944 | 944 | E |
| 88 | 05 | "I am not proposing here that annotated portfolios subsume ... particularly for an interdisciplinary audience" | 945 | 945 | E |
| 89 | 05 | "it is useful to consider specific designs such as a rotary controller for a microwave oven" | 944 | 944 | E |
| 90 | 05 | "turn to existing examples" | 943 | 943 | E |
| 91 | 05 | Ten principles, Figure 1 contents (seven photos, principles 2, 4, 5, 8), prototype TV, Lubs calculator | outside | v1 06 | v1 |
| 92 | 05 | Alexander "a million times over" | outside | v1 06 | v1 |
| 93 | 05 | Carroll and Kellogg missing; in-text [3] is Collins | outside | v1 06 | v1 |
| 94 | 06 | "scientism" | 938 | 938 | E |
| 95 | 06 | "definite nature of designed artefacts" | 943 | 943 | E |
| 96 | 06 | "a serious form of theoretical contribution" | 945 | 945 | E |
| 97 | 06 | "discursiveness and elaboration" | 946 | 946 | E |
| 98 | 06 | "wary of impulses" | 937 | 937 | E |
| 99 | 06 | "intermediate-level knowledge" (Löwgren 2013) | outside | v1 10 | v1 |
| 100 | 06 | "alternatives to scientistic ways of knowing"; 18 subcommittees, 6,730, 485, 37.9%, 38.7% | outside | v1 10 | v1 |
| 101 | 06 | "worked to formalize" p. 178; five steps p. 184 | outside | v1 10 | v1 |
| 102 | 06 | "contestable, defensible, and substantive"; social navigation, seamfulness | outside | v1 10 | v1 |
| 103 | 06 | "may capture best what we can know through design" | outside | v1 10 | v1 |
| 104 | 06 | Suh et al. "hundreds and thousands"; "towards rapid convergence on a limited set of ideas" | outside | v1 10 | v1 |
| 105 | 06 | RTD 2013 to 2019 cities, RTD Roams 2027; Pictorials 2014, TEI, C&C; citers list; OpenAlex about 1,100 | outside | v1 10 | v1 |
| 106 | 07 | "inconceivable" | 940 | 940 | E |
| 107 | 07 | "establishes an area" | 944 | 944 | E |
| 108 | 07 | "wary" (Gaver) | 937 | 937 | E |
| 109 | 07 | Ishii and Ullmer "one user interface metaphor" p. 240 | outside | v1 09 | v1 |
| 110 | 07 | Greenberg and Buxton "confirmatory evidence" p. 113; split cure p. 118 | outside | v1 09 | v1 |
| 111 | 07 | Shin et al. "ecological validity" p. 5; guessed purpose p. 7; four hours | outside | v1 09 | v1 (but see Blocker 5) |
| 112 | 07 | Olsen "the space between"; Wobbrock and Kientz "wary," "embrace giving definition" p. 44 | outside | v1 09 | v1 |
| 113 | 08 | "clearly articulating ... the standards one might want to advocate" | 941 | 941 | E |
| 114 | 08 | "productive research programmes" | 941 | 941 | E |
| 115 | 08 | "political acts to be approached with care" | 945 | 945 | E |
| 116 | 08 | "relative productivity of competing programmes as the grounds for scientists' choice amongst them" | 941 | 941 | E |
| 117 | 08 | "ontological politics" | 943 | 943 | E |
| 118 | 08 | Pullin [25] and Schön [28] listed, never cited in text | 946 | 946; no in-text [25] or [28] | E |
| 119 | 09 | "One thing I now expect from research through design." | progress.md | progress.md | E |
| 120 | 09 | Role cards, optional exercise, anchors | progress.md | progress.md | E |

Repeated uses of the same quotation in 08 ("some sympathy," "depends on the account," "neither dominates," "It is difficult to call," "Perhaps we do not all agree," "on its own terms," "actionable metrics") were checked and match the rows above.

## New facts (not in v1 or the paper)

| File | Fact | Marked in v2? |
|---|---|---|
| 01 | Gaver at the Royal College of Art in 2004; nine authors; "England and Wales"; terabyte; "apparent height, direction and speed" (CHI 2004 p. 886); Figure 1 photo | Yes, log row 5 (CHI 2004 paper read 2026-10-21) |
| 01 | Co-authors at KTH Stockholm, Lancaster and UCL | Partly: in Watch out, not in the log row |
| 02 | Gaver's HCI work "about twenty years" | Yes, log |
| 02 | Popper "Vienna then London" | No (and incomplete; Should fix 20) |
| 02 | Lakatos "Hungarian" | No |
| 05 | Rams worked for Braun and Vitsoe "for decades" | No |
| 05 | "accidents of 1960s manufacturing" | Yes, Watch out and Open items, as ours |
| 06 | Generative AI makes "hundreds of design variants in an afternoon"; portfolios "cost more than a study" | Yes, as ours (Watch out; "We raise it") |
