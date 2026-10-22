# Briefs for version 3 of the Week 7 deck (slides3/)

Written 2026-10-21 by the orchestrating session. Version 3 follows Oliver's outline: scene, what existed before the paper, Gaver's argument, judging it, afterwards, discussion. Twenty-two slides including the title; the talk is slides 1 to 20 (25:30 plus about three minutes of audience moments), slides 21 and 22 open the discussion half. Each slide carries its own speaker notes. At the end a script extracts the slide text alone into `slides-only/`.

## Part 1. Common rules

1. **Material.** Reuse verified material first: version 1 section files `../slides/01-Background.md` to `11-Discussion-Bridge.md` (each has a verification log; its quotation bank and 14 pitfalls are in `../slides/_agent-briefs.md`), version 2 section files `../slides2/01-Hook.md` to `09-Handover.md`, the paper text `../assets/paper-text.txt` (two-column extraction; printed page numbers appear as bare lines "937" to "946"; text before the first "937" line is page 937, text between the "937" line and the "938" line is page 938, and so on). Any outside fact not in those files needs a URL, fetched and read, and the mark NEW in the log. Anything you cannot verify is marked UNVERIFIED in the log and kept off the slide.
2. **Quotations** exact, with the paper's inner single quotes kept, and the printed page in the notes. Ellipses only where marked. On the slide, a pull quote carries only "(Gaver, p. N)"; no other page numbers on slides.
3. **Attribution.** "Gaver says" for his claims, "Gaver quotes" for what he quotes, "we argue" or "ours" for the presenter's. Never state a motive as fact. Keep every hedge (see Part 3, P15). Our labels ("baby science", "judo move", "caption") are flagged as ours in Watch out.
4. **Slides.** The headline is a sentence or a question. At most about 60 words on a slide, pull quote included. Plain words first, Gaver's phrase second. One visual per slide, described precisely.
5. **Say** scripts at or under 130 words per minute of the slide's time, ending with "(about N words)". Written to be read aloud: short sentences, no long appositives.
6. **Transition** is one or two sentences that lead into the next slide's first sentence. Check against the next slide, including across files (the plan in Part 4 gives every slide's opening idea).
7. **No em-dashes or en-dashes** anywhere. Use commas, colons, "to".
8. **Audience moments** (hands on slide 2, the vote on slide 10, hands on slide 12, hands again on slide 20) carry explicit cues in the Say: "[pause]", "[click]", "Hands if yes."
9. **Placeholders.** `**[classmate]**` where a classmate's commentary will be credited (slides 15, 16, 21, 22). Oliver's own project stays out of the deck.
10. After writing, run `python3 slides3/_check-sections.py` from `week7/` and fix every OVER and every dash.

## Part 2. File format (exact; the extraction script depends on it)

```
# 0N Part title

**Purpose:** one or two sentences. **Slides:** N to M. **Time:** total. **Speaker:** Oliver. **Previous:** ... **Next:** ...

Written 2026-10-21 by an Opus writer from `_agent-briefs.md`.

## Slide N: Headline as a sentence

**Time:** m:ss · **Visual:** the layout in one to three sentences (where each element sits, type sizes, colours only as "our colour" or "second colour").
**Images:** file path relative to `week7/` (for example `assets/drift-table-chi2004-fig1.png`), or "none".
**Build:** the click order in one sentence, or "all at once".

**On the slide**
- bullet (bold with ** where the slide is bold)
- ...
- Quote: "exact words" (Gaver, p. N)      <- a pull quote is a bullet beginning with "Quote: "

**Speaker notes**
- **Say:** ... (about N words)
- **Transition:** ...
- **Watch out:**
  - ...
- **If asked:**
  - *Q: ...?* A: ... (page or URL)
  - ...
- **Sources:** pages; verified files; URLs.

## Verification log

| Claim | Source | Status |
|---|---|---|
| ... | ... | OK / OK, reused / NEW, marked / UNVERIFIED |

## Open items
- ...
```

One file per part, slides in order. The `## Verification log` and `## Open items` headings appear once, at the end of the file. The check script reads `**Time:**` on the line after the slide heading, `**On the slide**` up to `**Speaker notes**`, and `**Say:**` up to `- **Transition:**`.

## Part 3. Pitfalls (all found during versions 1 and 2; do not repeat them)

- **P1** The Drift Table is not in Gaver's paper. In 2004 Gaver was at the Royal College of Art, not Goldsmiths. Say "Gaver and colleagues", never "the studio", for 2004 work. The aerial photography is "of England and Wales" (CHI 2004 paper, p. 886). Details and the figure: `../slides2/01-Hook.md`.
- **P2** The critics asked for standards and criteria, not for design to become a science. Do not give them a motive beyond what Gaver writes: conformance "may be a route towards disciplinary legitimacy" (p. 937).
- **P3** "extensible and verifiable" is Gaver's paraphrase (p. 938). The critics' words are "a level of extensibility and verifiability" (p. 937, quoting [33]).
- **P4** The manifestos Gaver lists on p. 941 are Reflective Design, Supple Interfaces and Designing for the Self (Zimmerman's). Ludic design is Gaver's own manifesto (p. 938); if added, mark as our addition. Designing for the Self is by one of the critics, which is worth saying.
- **P5** "What Popper would reject as 'merely' confirmatory evidence is, from this point of view, testimony to the fertility of the overarching theory" (p. 941) is Gaver's sentence written from a Lakatosian view. Never "Lakatos says".
- **P6** "must founder" (p. 942) is Gaver's wording, stronger than Kuhn's; say "in Gaver's words". The "most of us" list (p. 942) is hedged twice: "seem to share" and "Perhaps we do not all agree with all of these assumptions all of the time".
- **P7** Figure 1 (p. 945): seven photographs (two show one product in three colours), four of Rams's ten principles (the text names two, p. 944), notes written for Gaver's sketch, products for Braun and Vitsoe, one a prototype never produced, images from Flickr Creative Commons, caption "not defined by their graphic presentation". Ten principles as Vitsoe lists them are in `../slides/06-Annotation.md`.
- **P8** Gaver's word is "annotation"; "caption" is ours. The balance rule is "neither dominates, and neither is subservient" (p. 945). He is "not proposing here that annotated portfolios subsume all other forms of design theory" (p. 945); theory's role is "limited to inspiration and annotation" (p. 945).
- **P9** Slips in the paper: Carroll and Kellogg are cited on p. 944 but missing from the reference list; "[34]" on p. 939 should be [33]; "Reprinted in [3]" should be [5]; the text spells "irresistables" (p. 938). Mention only if asked.
- **P10** Aftermath facts (`../slides/10-What-Followed.md`, verified by URL): Gaver and Bowers 2012 is the same studio, so continuation, not uptake; the RTD conference ran 2013, 2015, 2017, 2019, paused, returns 2027 as "RTD Roams"; "The critics' way" is our lane label (Höök and Löwgren 2012 cite the essay); Zimmerman and Forlizzi 2014 do not cite the essay; the generative AI link (Suh et al. 2024; Davis et al. 2024) is ours; the CHI 2026 Design subcommittee text, "alternatives to scientistic ways of knowing", is a scope statement, not reviewing practice; OpenAlex counted about 1,100 citations in October 2026.
- **P11** Shin et al. 2019 (Week 6): realistic, yes; unambiguous, no (actuation always on day 2; 11 of 12 guessed the purpose, their p. 7). Gaver is "half right", not refuted.
- **P12** Wobbrock, Morris and Wilson 2009 use a method from 2005; a 2020 review counts 216 such studies. Never "follow-up studies".
- **P13** "productive research programmes" (p. 941) is exact; Gaver names Lakatos's criterion, "the relative productivity of competing programmes" (p. 941). The limitation is that no judge is named, not that the criterion is absent.
- **P14** "Funding" and "examining" are our extension of the critics' problem; Gaver says integration "would make assessment of new contributions fairer and more rigorous" (p. 941) and never returns to it.
- **P15** Hedges to keep: "I have some sympathy with these concerns" (p. 937); standards "might lead to a form of self-policing that would be overly restrictive" (p. 937); "insofar as HCI is prone to 'scientism'" (p. 938); "may adopt, or be interpreted in terms of, inappropriate 'scientific' models" (p. 938); theory "tends to be" or "likely to produce theories that are" "provisional, contingent, and aspirational" (serial comma on p. 937 and p. 941; no serial comma on p. 938 and p. 945); "in several respects, the converse of Alexander's [1] design patterns" (p. 945); "seems naturally to lead to convergence" (p. 943); "potentially repressive acts of ontological politics" (p. 943).
- **P16** Gaver's own stance: "Personally, I tend to think that the nature of design practice and research is distinct from the ideals, and often practice, of science" (p. 941), but whether it is a science "depends on the account of science one adopts" (p. 941).
- **P17** The three conclusions: "moderate expectations of creating extensible and verifiable theory" (p. 945); "attempts to establish disciplinary norms of process or outcome are political acts to be approached with care" (p. 945); "the recognition of annotated portfolios as a way of reflecting and valuing the particular nature of design theory and practice" (p. 946); "consensus can come to look like constraint" (p. 946).
- **P18** Reviewer R4: the essay addresses, "in part, 'an elephant in the room...: the move towards "scientism" in HCI'" (p. 946). Gaver cites Sengers 2006 on scientism (p. 941), so he names it in a CHI paper, not first.

## Part 4. The 22 slides, with time, must-haves and sources

Part A, file `01-Scene.md`, slides 2 to 5 (writer W1). Part B, file `02-Before.md`, slides 6 to 8 (writer W1). Part C, file `03-Argument.md`, slides 9 to 13 (writer W2). Part D, file `04-Judging.md`, slides 14 to 16 (writer W3). Part E, file `05-Afterwards.md`, slides 17 to 20 (writer W3). Part F, file `06-Discussion.md`, slides 21 and 22 (writer W3). Slide 1 (title) is emitted by the build script.

| # | Slide | Time | Opens with |
|---|---|---|---|
| 1 | Title | 0:30 | (script) |
| 2 | What research through design is, and an object to argue about | 1:30 | the Drift Table photograph; hands |
| 3 | How design got into HCI, 1993 to 2007 | 1:00 | the timeline |
| 4 | How the mess showed itself: dozens of manifestos | 1:00 | the manifestos |
| 5 | The reaction: the call for standards, 2007 to 2011 | 1:15 | the critics' words |
| 6 | How research through design was being done | 1:00 | artefacts, deployments, five kinds of theory |
| 7 | How people proposed to judge it | 1:00 | the 2007 four criteria |
| 8 | Why Gaver objects: self-policing and scientism | 1:15 | "some sympathy" |
| 9 | Which science? Popper versus Lakatos | 1:30 | the judo move |
| 10 | Why the field does not converge: three readings, and a vote | 2:00 | dozens of manifestos, again |
| 11 | The method Gaver proposes: theory as annotation | 1:30 | the flip |
| 12 | The annotated portfolio, shown: Dieter Rams | 2:00 | principles alone |
| 13 | How to make one, and what it is not | 1:00 | the balance rule |
| 14 | Upsides of his method | 1:30 | respects the particular |
| 15 | Critical reflection | 1:30 | one sketched example |
| 16 | Our verdict | 1:00 | the thesis sentence |
| 17 | What the paper changed, 2012 to 2016 | 1:30 | Gaver and Bowers 2012 |
| 18 | Value today, 2026 | 1:30 | CHI 2026 |
| 19 | Where this touches what we have read | 1:30 | Tangible Bits |
| 20 | Does it apply to your project? | 0:30 | the Drift Table again; hands |
| 21 | Discussion anchors | 0:45 | five tensions |
| 22 | Role cards and the plan for thirty minutes | 0:30 | the plan |

Talk: slides 1 to 20, 25:30. Discussion slides 21 and 22 are shown after the break.

### Slide 2. What research through design is, and an object to argue about (1:30)
Must contain: Frayling 1993's three kinds (into, through, for), Gaver's definition (design practice "brought to bear on situations chosen for their topical and theoretical potential", output "primarily, of artefacts and systems", p. 937), the Drift Table photograph with a one-line description, the question "Is this research?" with Yes / No / Depends what counts, and the fact that Zimmerman, Forlizzi and Evenson 2007 call it "a well known example of critical design in HCI" (p. 496). Audience moment: hands, counted on a card, asked again on slide 20. Sources: `../slides/01-Background.md` slide 01.1 and log; `../slides2/01-Hook.md` (photo `assets/drift-table-chi2004-fig1.png`, P1). About 55 words on the slide.

### Slide 3. How design got into HCI, 1993 to 2007 (1:00)
Must contain: a timeline of four to six dots: Frayling 1993 (Royal College of Art paper naming research through design); Fallman 2003, "Design-oriented human-computer interaction", CHI 2003 (Gaver's [7], which he cites for characterising HCI "as a design discipline", p. 938); Bødker 2006 third-wave HCI (verified from its abstract in v1 01); Zimmerman, Forlizzi and Evenson 2007, CHI, research through design as a named method with four criteria. Optional, if verified by URL during writing: Schön 1983 (reflective practitioner, in Gaver's list [28], never cited) and Cross 1982 "designerly ways of knowing" (absent from the paper). Research pass required for Fallman 2003's content: fetch an abstract or the paper, log the URL, mark NEW. Keep the slide to the dots and one line each. Sources: `../slides/01-Background.md` (Frayling, Bødker, Zimmerman 2007 with URLs), paper p. 938 ([7]).

### Slide 4. How the mess showed itself: dozens of manifestos (1:00)
Must contain: Gaver's observation "there must be dozens of different 'manifestos' for design, each seeking to establish a research programme, and each with its own configuration of goals, issues, methods and realised exemplars" (p. 941); four or five named examples with their one-line claims: Reflective Design (Sengers et al. 2005, "reflection on unconscious values ... can and should be a core principle of technology design", p. 938), ludic design (Gaver, "an antidote to assumptions that technology should provide clear, efficient solutions to practical problems", p. 938), Supple Interfaces (Isbister and Höök 2007), Designing for the Self (Zimmerman 2009), Product Ecology (Forlizzi 2008, a framework); and the symptom in Gaver's words: "little agreement about the values for which we should design, the appropriate methods for doing so, standards for evaluation or agreed forms of output" (p. 942). Sources: paper pp. 938 to 942; `../slides/04-Theory.md` slide 04.1; `../slides/05-Convergence.md`.

### Slide 5. The reaction: the call for standards, 2007 to 2011 (1:15)
Must contain: the timeline 2007, 2008, 2010, 2011 and the critics' words with the right source each: "lacks clear expectations and standards" and "actionable metrics" (CHI 2011 panel "Quality control", Forlizzi et al., [9], p. 937); "develop protocols, descriptions, and guidelines" (Zimmerman, Stolterman and Forlizzi, DIS 2010, [35], p. 937); "a level of extensibility and verifiability" (Zimmerman and Forlizzi 2008, [33], p. 937); and Gaver's reading that this "may be a route towards disciplinary legitimacy" (p. 937, P2). Sources: `../slides/01-Background.md` slide 01.2; `../slides2/02-Fight.md`.

### Slide 6. How research through design was being done (1:00)
Must contain: the practice (artefacts and systems built, "sometimes with associated accounts of how these are used in field tests", p. 937; the Drift Table lived with for weeks, the Prayer Companion for ten months, from `../slides/01-Background.md` background item 3) and the five kinds of theory Gaver inventories (pp. 938 to 939): articulations of practice (Djajadiningrat et al. 2007, embodied movement), borrowed theory (affordances, "irresistables"), manifestos, frameworks (Product Ecology), meta-theory (Zimmerman and Forlizzi 2008, Stolterman 2008, this essay). Sources: `../slides/04-Theory.md` slide 04.1 and notes; paper pp. 937 to 939.

### Slide 7. How people proposed to judge it (1:00)
Must contain: Zimmerman, Forlizzi and Evenson 2007's four criteria, process (rigour and a documented rationale, reproducible "with no expectation of the same result"), invention, relevance (in place of validity), extensibility; the 2010 paper's call for formalisation; the 2011 panel's "actionable metrics". Say that these are the proposals Gaver answers. Sources: `../slides/01-Background.md` background item 2 (pp. 499 to 500 of the 2007 paper, with URL); paper p. 937.

### Slide 8. Why Gaver objects: self-policing and scientism (1:15)
Must contain, in order: "I have some sympathy with these concerns" (p. 937); standards "might lead to a form of self-policing that would be overly restrictive" (p. 937); "insofar as HCI is prone to 'scientism'", standards "may adopt, or be interpreted in terms of, inappropriate 'scientific' models" (p. 938); reviewer R4 (p. 946, P18); the "why you care" line, ours: "Where is your evaluation?" is this fight in four words. Sources: `../slides2/02-Fight.md` slide 02.1; `../slides/02-Problem.md`.

### Slide 9. Which science? Popper versus Lakatos (1:30)
Must contain: the judo move (Gaver never says design is science; he asks which science); Popper's test (falsifiability, p. 939) and why design theory fails it (the "implicit sometimes", p. 940; the ambiguity claim and "Simply tweaking", p. 940); Lakatos's test (fertility; "predict novel facts", p. 940) and why design theory passes it (P5, p. 941); the result: theory that "tends to be" "provisional, contingent, and aspirational", "theories that are sometimes right" (p. 941); the pull quote "It is difficult to call for research to be more scientific if what it means to be scientific is under question" (p. 941). Keep it to one slide by using a two-column scorecard. Sources: `../slides2/02-Fight.md` slide 02.2; `../slides2/03-Two-Tests.md` both slides.

### Slide 10. Why the field does not converge: three readings, and a vote (2:00)
Must contain: the three readings with our labels and Gaver's wording in the notes: A, pre-paradigmatic (Kuhn; "must founder", p. 942, P6); B, science is messier up close (Latour, p. 942) and "most of us ... seem to share" values (p. 942); C, many worlds (Law and Urry, p. 943; the iPad sentence, p. 943; "potentially repressive acts of ontological politics", p. 943); what each implies for standards (only A says standardise; Gaver's verdict, p. 938); the vote: "which is design research in HCI, as you have met it?" with the Fix tags built after the count. Audience moment. Sources: `../slides2/04-Puzzle.md`.

### Slide 11. The method Gaver proposes: theory as annotation (1:30)
Must contain: the flip (usual order, theory leads and the system illustrates; Gaver's order, artefacts are "the definite facts of research through design", p. 945); "ultimate particular" (Stolterman, p. 944); "theory nexus" (Carroll and Kellogg, p. 944, P9); "If artefacts embody theory, however, they do not encode it" (p. 944); theory as annotation, "to explain and point to features of 'ultimate particulars'" (p. 944); from one to many: "a point in design space", a portfolio "establishes an area", "the invariances among them" (p. 944). Sources: `../slides2/05-Rams.md` slide 05.2; `../slides/06-Annotation.md` slide 06.1.

### Slide 12. The annotated portfolio, shown: Dieter Rams (2:00)
Must contain: the three-step test as a build on one slide: principles alone ("Good design is unobtrusive", "Good design is thorough down to the last detail", p. 944), products alone (`assets/rams-fig1-products-only.png`), both (`assets/rams-fig1-annotated.png`); "Hands if yes" at each step; the definition after the third: related designs plus short notes on what joins them, "neither dominates, and neither is subservient" (p. 945); "one can imagine 'designing in the style of Dieter Rams' without the need for a detailed formal theory" (p. 944). Audience moment with [click] and [pause] cues. Sources: `../slides2/05-Rams.md` slide 05.1 (P7).

### Slide 13. How to make one, and what it is not (1:00)
Must contain: "The notion of annotated portfolios is not a formal one" and what defines it, the balance "between descriptions of specific, detailed examples of design practice, and articulations of the issues, values and themes" (p. 944); presentation "may vary depending on material, purpose and audience" (p. 944); the Vitsoe site and an illustrated essay as other forms (p. 945); who annotates (open; Rams annotated Rams); what it is not: "in several respects, the converse of Alexander's [1] design patterns" (p. 945), not a replacement for theoretical writing (p. 945), the claim to "'extensibility and verifiability'" (p. 945). Sources: paper pp. 944 to 945; `../slides/06-Annotation.md` slide 06.2 and background.

### Slide 14. Upsides of his method (1:30)
Must contain five or six upsides, each anchored: respects "the 'ultimate particular' of the designed artefact" (p. 945); keeps provisional theory honest by "yoking them to particular design manifestations" (p. 945); allows "'extensibility and verifiability'" in his sense (p. 945); "can support multiple conceptual perspectives" (p. 945); matches how designers already work, they "turn to existing examples of design" (p. 943); costs no new method, it is a form of presentation (ours, say so). Sources: paper pp. 943 to 945; `../slides/08-Strengths-Limitations.md` slide 08.1.

### Slide 15. Critical reflection (1:30)
Must contain the limitations as "we argue" with Gaver's best reply in the notes: one sketched example, and it is Rams's commercial work for Braun and Vitsoe, not research through design (p. 945); "productive research programmes" is a standard with no judge named (p. 941, P13); silent on reviewing, funding and examining (P14); the "most of us" list already works as criteria, standards without metrics (p. 942); by its own taxonomy a manifesto with one example, though Gaver calls it "a contribution to meta-theory" (p. 939); who annotates, and do portfolios favour funded labs (the essay's ERC Advanced Investigator Award, p. 946). `**[classmate]**` slot. Sources: `../slides2/08-Verdict.md`; `../slides/08-Strengths-Limitations.md` slide 08.2.

### Slide 16. Our verdict (1:00)
Must contain: the thesis, a candidate for Oliver: "Gaver wins the argument against scientism. He concedes more to his critics than his conclusion admits, and he never says how a field without shared criteria decides what to publish, fund and examine." Three supports: he accepts "clearly articulating ... the standards one might want to advocate" (p. 941) and a shared practice (p. 942); "productive research programmes" has no judge; the portfolio is a form, not a criterion ("not defined by their graphic presentation", p. 945). `**[classmate]**` slot. Sources: `../slides/11-Discussion-Bridge.md` slide 11.1; `../slides2/08-Verdict.md`.

### Slide 17. What the paper changed, 2012 to 2016 (1:30)
Must contain the three lanes with years (P10): Gaver's way (Gaver and Bowers 2012 "Annotated Portfolios", Interactions; Bowers DIS 2012; RTD conference 2013, exhibition format; DIS Pictorials 2014; Pierce DIS 2014), the critics' way (strong concepts, Höök and Löwgren 2012; bridging concepts, Dalsgaard and Dindler 2014; Zimmerman and Forlizzi 2014 "worked to formalize"; Bardzell et al. 2016 documentation framework), both (Löwgren 2013 intermediate-level knowledge; Ways of Knowing in HCI 2014 printing both chapters); which of these cite the essay; about 1,100 citations (OpenAlex, October 2026). Mark the lane labels as ours. Sources: `../slides/10-What-Followed.md` (all verified by URL); `../slides2/06-Aftermath.md`.

### Slide 18. Value today, 2026 (1:30)
Must contain: the CHI 2026 Design subcommittee invites work exploring "alternatives to scientistic ways of knowing" (485 submissions, 37.9% to revise and resubmit, a scope statement not reviewing practice); RTD returns in 2027 as RTD Roams (York, Melbourne, Tokyo); Pictorials at TEI and Creativity and Cognition; generative AI: Suh et al. CHI 2024 (models produce "hundreds and thousands" of outputs, interfaces push "towards rapid convergence on a limited set of ideas") and Davis et al. CHI 2024 (graphical design-space interfaces beat text prompts for ideation), with the link to Gaver's "definite nature of designed artefacts" (p. 943) marked as our question; and the question whether portfolios favour funded labs. Sources: `../slides/10-What-Followed.md` background "Facts for the three open questions"; `../slides2/06-Aftermath.md`.

### Slide 19. Where this touches what we have read (1:30)
Must contain four cards: Tangible Bits 1997 (Ishii and Ullmer, "one user interface metaphor", three platforms, their p. 240; an annotated portfolio before the name, ours); Greenberg and Buxton 2008 ("confirmatory evidence" p. 113; split cure p. 118; Gaver drops refutation as the test for design theory); Wobbrock, Morris and Wilson 2009 (P12; fertile and still science); Shin et al. 2019 (P11). Footer: Gaver cites none of these; every link is ours. Alternates in the notes: Olsen 2007; Wobbrock and Kientz 2016. Sources: `../slides2/07-Connections.md`; `../slides/09-Connections.md`.

### Slide 20. Does it apply to your project? (0:30)
Must contain: the Drift Table photograph small, the question "Is the Drift Table research?" with Yes / No / Depends what counts, "Did you move?", and one line turning it on the room: "Name the one world your project assumes." (from the many-worlds reading, p. 943). Audience moment: hands again, compared with slide 2. Transition into the break and the discussion slides. Sources: `../slides2/09-Handover.md`.

### Slide 21. Discussion anchors (0:45)
Must contain the five tensions as open questions, each with more than one defensible answer, with the page in the notes: whose interest do standards serve (p. 937); is theory that is "sometimes right" worth publishing, and who counts fertility (p. 941); which reading of non-convergence fits design research in HCI (pp. 941 to 943); would you accept an annotated portfolio as a CHI paper, and what would you ask the authors to add (p. 945); fourteen years on, who won. `**[classmate]**` credits beside the ones classmates raised. The notes carry zoom-in and zoom-out probes for each, in the style of `../slides/11-Discussion-Bridge.md` slide 11.2 and the discussion plan in `../progress.md`.

### Slide 22. Role cards and the plan for thirty minutes (0:30)
Must contain: the three steps (hands and two voices from each corner; pairs, five minutes, one card each with a question and two roles; a one-sentence round, eight to ten voices, "One thing I now expect from research through design"); the five candidate role cards from `../progress.md` (Gaver versus Zimmerman on standards; a reviewer versus an author on a portfolio submission; a student versus a studio head on who benefits from no criteria; Popper versus Lakatos on a claim from Tangible Bits; an artist versus a scientist on many worlds); `**[classmate]**` strip. The notes carry the grading pattern (cited the paper with a page, cited the talk with a slide, acknowledged the other side's strongest point, applied the argument to their own project) and the optional six-minute live annotated-portfolio exercise (five course artefacts: Weiser's pads, the Surface gesture set, metaDESK, Zooids, Kismet).

## Part 5. Writer assignments

- **W1:** `01-Scene.md` (slides 2 to 5) and `02-Before.md` (slides 6 to 8). Research pass for slide 3 (Fallman 2003; optional Schön 1983 and Cross 1982), with URLs logged.
- **W2:** `03-Argument.md` (slides 9 to 13).
- **W3:** `04-Judging.md` (slides 14 to 16), `05-Afterwards.md` (slides 17 to 20), `06-Discussion.md` (slides 21 and 22).

Each writer: read Parts 1 to 4 of this file, the version 1 and 2 files named for your slides, and the paper text for every page you cite. Write the files in the exact format of Part 2. Run the check script and fix overruns. Reply with: slides written, Say word counts against limits, any NEW or UNVERIFIED items, and anything in the brief you could not do.
