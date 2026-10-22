# Progress: CMPT 863 Week 7 presentation

Last updated 2026-10-21. Presentation date: 2026-10-22. Oliver presents alone (confirmed 2026-10-19). This is the only paper for the week. Course: CMPT 863 Human-Computer Interaction. Format: 30 minutes presentation, 30 minutes discussion led by the presenter. Draft slides due through the course form by the day before class, 11:59 pm; final version before class. Presenters do not write a commentary for their own week, and they grade the classmates' commentaries and participation after class.

Rubric as recorded for Week 4: presentation quality 7, integration of classmates' commentaries 4, discussion 7 (18 points).

## Where we stand

The paper has been read in full (text; Figure 1 on PDF page 9 rendered and checked). The presentation now lives in `slides/`, one Markdown file per section, each holding slide content, speaker notes, a verification log and open items. `slides/00-Overview.md` has the run sheet (17 slides plus title, 24:15), the arc with every hand-over, the placeholders, the verification status per file, the corrections made this round and the working rules. The briefs are in `slides/_agent-briefs.md`; the readings-enrichment report with 31 candidate connections is `reports/opus-R-readings-enrichment.md`. Generated deck files are built by `slides/_build-deck-source.py`; mechanical checks by `slides/_check-sections.py`.

**Next step:** Oliver reads the full layout (`Week 7 - Deck Source.md`), accepts or changes the thesis on slide 11.1, picks the connections for slide 13 from the alternates, and answers the questions on sets 13 to 15 and the live exercise. Then the discussion guide (`slides/12-Discussion-Guide.md`) is written and cross-checked, the commentaries fill the [classmate] slots, and the deck is built for the 2026-10-21 draft submission.

## The paper in short

William Gaver, "What Should We Expect From Research Through Design?", CHI 2012, Austin, session Critical Perspectives on Design, pp. 937 to 946. An essay, not a study. Zimmerman, Forlizzi, Stolterman and the 2011 "Quality control" panel had asked research through design to adopt agreed standards, protocols and "extensible and verifiable" theory. Gaver argues in three parts: (1) measured against Popper and Lakatos, design theory is unfalsifiable because theory underspecifies design and design is generative, so theory is "provisional, contingent, and aspirational" (p. 937) and the question whether design is a science is futile; (2) the field's lack of convergence has three readings, pre-paradigmatic, invisible consensus, and "many worlds," and only the first supports standardisation; (3) artefacts are the "definite facts" of the field (p. 945), theory should annotate them, and annotated portfolios of related designs should count as a serious theoretical contribution. Closing line: "consensus can come to look like constraint" (p. 946).

## Page mapping and assets

PDF page N = printed page 936 + N. Printed pages 937 to 946. Figure 1 (the Rams annotated portfolio sketch) is on p. 945, PDF page 9; references on p. 946. Render with `pdftoppm -r 200 -f 9 -l 9 -png "<pdf>" assets/p9` and crop.

Figure 1 shows seven photographs of Rams products for Braun and Vitsoe (coffee maker, shelving, radio, lounge chair, portable TV, a shaver the note calls the S3, a small portable device) each with a compact colour-coded note written about Rams in the third person (so the notes are Gaver's sketch, not Rams's words), and four of Rams's ten principles in the centre: understandable, useful, unobtrusive, thorough. Caption: "note that annotated portfolios are not defined by their graphic presentation."

## What is different from Week 4

1. **An essay, not a study.** No participants, no numbers, one figure. The seven course questions need re-reading: "how was the research conducted" becomes "how does the argument work, and what counts as its evidence." The verification load shifts from numbers to exact quotations and correct attribution: Gaver's own words versus what he quotes from Zimmerman et al., Popper, Lakatos, Kuhn, Latour, Law and Urry, Stolterman.
2. **The risk is a summary.** The class has read it, and the course brief forbids a section-by-section retelling. Rule for every section: state a claim Gaver makes, the evidence he gives for it, and a test of it. Each part of the paper is presented as a move in an argument that the room is invited to accept or reject, not as a chapter.
3. **The paper is about the class's own practice.** Most CMPT 863 projects build something; Oliver's own project is a hypothesis-driven study with conditions and measures, which is the model Gaver sets design against. The discussion can turn on the students' own projects, which Week 4 could not do.
4. **Fewer facts, more positions.** The Week 4 pitfalls list was about numbers. The Week 7 list will be about attributions and about not overstating what Gaver claims (he hedges often; the deck must keep his hedges).

## Proposed arc and run sheet

Shape follows Oliver's Week 4 order: a rundown of the argument first, then the course questions, then the material that feeds the discussion. 17 slides plus title, 24:15 planned, about 5:45 of buffer. Times are a first budget.

| # | Slide | Section file | Time |
|---|---|---|---|
| 0 | Title | title slide | 0:30 |
| 1 | What research through design is, and whose studio this is | 01-Background | 1:30 |
| 2 | 2007 to 2012: the call for standards | 01-Background | 1:30 |
| 3 | The problem: standards, scientism, self-policing | 02-Problem | 1:15 |
| 4 | The argument on one slide | 03-Argument-Map | 1:30 |
| 5 | Move 1a: five genres of design theory, and the Popper test | 04-Theory | 1:30 |
| 6 | Move 1b: Lakatos, and what "provisional, contingent, aspirational" buys | 04-Theory | 1:30 |
| 7 | Move 2: three readings of why the field does not converge | 05-Convergence | 1:45 |
| 8 | Move 3a: artefacts as theory nexus, theory as annotation | 06-Annotation | 1:30 |
| 9 | Move 3b: annotated portfolios, Rams, and pattern languages | 06-Annotation | 1:30 |
| 10 | How the essay argues: its method, its evidence, its reflexivity | 07-Method | 1:30 |
| 11 | Strengths | 08-Strengths-Limitations | 1:00 |
| 12 | Limitations, and what the essay does not do | 08-Strengths-Limitations | 1:30 |
| 13 | Connections to the course readings | 09-Connections | 1:45 |
| 14 | What followed, 2012 to 2026 | 10-What-Followed | 1:30 |
| 15 | Our reading of the paper | 11-Discussion-Bridge | 0:45 |
| 16 | Tensions worth arguing about | 11-Discussion-Bridge | 1:15 |
| 17 | Over to you | 11-Discussion-Bridge | 1:00 |
| | **Total** | | **24:15** |

Discussion guide: `12-Discussion-Guide.md`, same pattern as Week 4's file 16.

## Section by section: what each must do, and the critical angle

### 01 Background (2 slides)

Goal: put the room in 2012 and explain who is arguing with whom. Not a history of design research.

- Slide 1: research through design in one line (Frayling 1993 "research through design" [11]; Zimmerman, Forlizzi and Evenson CHI 2007 [34] as the paper that brought the term into CHI). What Gaver's own studio makes: the Interaction Research Studio at Goldsmiths and its artefacts (Drift Table, History Tablecloth, Prayer Companion, Indoor Weather Stations; verify each with a URL, drop what cannot be verified). The point: the essay is written by a practitioner with a portfolio, and the essay's proposal is the form his own work takes. Gaver's funding, ERC "ThirdWaveHCI" (p. 946 acknowledgements), names the camp.
- Slide 2: the call for standards. The quotations Gaver answers: "lacks clear expectations and standards for what constitutes 'good' design research" and "actionable metrics" (Forlizzi et al. 2011 panel [9], quoted p. 937); "develop protocols, descriptions, and guidelines," "like any other research approach," "a proper research methodology that can produce relevant and rigorous theory" (Zimmerman, Stolterman and Forlizzi DIS 2010 [35], quoted p. 937); "a level of extensibility and verifiability" (Zimmerman and Forlizzi 2008 [33], quoted p. 937). The CHI 2012 session title, Critical Perspectives on Design. Reviewer R4's phrase, "an elephant in the room ...: the move towards 'scientism' in HCI" (p. 946).

### 02 The problem (1 slide)

Goal: state the problem Gaver addresses and why it mattered, as analysis. The surface problem: what should the community expect from research through design? The deeper problem as Gaver frames it: the proposed cure (standards) risks "self-policing that would be overly restrictive" (p. 937) and may import "inappropriate 'scientific' models" because HCI is "prone to 'scientism'" (p. 938). Note what Gaver concedes: "I have some sympathy with these concerns" (p. 937) and that standards "may be a route towards disciplinary legitimacy" (p. 937). Why it mattered then: design practitioners "increasingly integrated within the HCI research community" (p. 937). Why it matters now: every student here will be reviewed by someone who holds one of these two views.

### 03 The argument map (1 slide)

Goal: the whole essay on one slide so nobody needs a retelling afterwards. Three moves, each with its premise, its conclusion and the page. Move 1: Popper and Lakatos, p. 939 to 941, conclusion "provisional, contingent and aspirational." Move 2: three readings of non-convergence, p. 941 to 943, conclusion "proliferation rather than agreement." Move 3: theory as annotation, p. 943 to 945, conclusion "annotated portfolios as a serious form of theoretical contribution." Then the one-line conclusion and the three "I suggest" sentences of p. 945 to 946.

### 04 Theory: the Popper test and the Lakatos rescue (2 slides)

Goal: present Move 1 as a test the paper runs on itself, and then test the test.

- Slide 5: the five genres of conceptual work Gaver catalogues (articulations of practice, e.g. Djajadiningrat [6]; borrowed theory, e.g. affordances, "irresistables" [23]; manifestos, e.g. Reflective Design [30], ludic design [13]; frameworks, e.g. Product Ecology [10]; meta-theory, e.g. Zimmerman and Forlizzi, Stolterman, and this essay itself, p. 938 to 939). Then Popper: falsifiability as the criterion (p. 939). Why design theory fails it: theory underspecifies design (wicked problems [26]; many decisions; design changes its own context [21], p. 940), so claims take the form "designing for X can lead to successful outcomes" with an "implicit sometimes" (p. 940), which no failure can refute; and tests that are "simultaneously unambiguous and ecologically valid are inconceivable" (p. 940, the ambiguity example). Design is generative: "theories that are sometimes right" (p. 941).
- Slide 6: Lakatos: research programmes with a hard core, a protective belt, judged by whether they "predict novel facts" (p. 940). Under this account the examples that accompany manifestos are "testimony to the fertility of the overarching theory" (p. 941). Gaver's hedge: "I am not presenting Popper's and Lakatos' accounts as particularly representative of Philosophy of Science" (p. 941); his conclusion that whether design is a science "depends on the account of science one adopts" (p. 941).
- Critical angle (labelled "we argue"): (a) the "implicit sometimes" structure is not unique to design; medicine, education and management make claims of the same form and still built evidence standards, so unfalsifiability does not by itself exempt a field from standards; (b) Lakatos's productivity is itself a standard, and Lakatos also has degenerating programmes; Gaver never says how fertility is judged or by whom; (c) the ambiguity example shows that controlled comparison fails, not that evaluation fails; Gaver's own studio reports field deployments, which is evidence of a kind he does not discuss here.

### 05 Convergence: three readings (1 slide)

Goal: present Move 2 as a choice the room has to make.

- The three readings, each with its source and its implication: pre-paradigmatic (Kuhn [17], p. 941 to 942; implies standardise); invisible consensus (Latour [20]; the list of what "most of us" share, user contact, exploring a wide design space, craft, "the practice of making is a route to discovery," p. 942; implies there is no problem); many worlds (Law and Urry [21], Osborne and Rose [22]; arts, architecture, fashion as fields that are "cumulative in the way a conversation is," p. 942 to 943; the iPad example, p. 943; standards as "potentially repressive acts of ontological politics," p. 943; implies proliferation is progress).
- Critical angle: (a) the three readings are not exclusive, and the "invisible consensus" list is itself a draft of the shared description Zimmerman et al. asked for; writing down what "most of us" already share is not obviously repressive; (b) Gaver's "Pre-paradigmatic research, in contrast, must founder" (p. 942) is stronger than Kuhn, who grants such fields "valuable contributions" (Gaver's own words, p. 942); (c) the many-worlds argument proves too much: if every criterion is ontological politics, so is the call to value proliferation; Gaver softens this in the conclusion to "political acts to be approached with care" (p. 945); (d) the opponents asked for standards, not for a single paradigm; the section answers the second, not the first.

### 06 Annotation: artefacts, theory and portfolios (2 slides)

Goal: present Move 3 as the paper's constructive proposal and ask whether it does what Gaver claims for it.

- Slide 8: designers use examples more than theory; artefacts are "ultimate particulars" (Stolterman [31], p. 944) with "the same dignity and importance as truth in science" (quoted p. 944); an artefact is a "theory nexus" (Carroll and Kellogg, p. 944; see the reference-list note below); artefacts "embody theory" but "do not encode it" (p. 944); so theory's role is annotation, "serving to explain and point to features of 'ultimate particulars', the truths of design" (p. 944). Theory is underspecified by design as well as the reverse (p. 944).
- Slide 9: annotated portfolios. One design is a point in design space; a portfolio "establishes an area" and makes visible "the dimensions along which a designer's choices may range and the invariances among them" (p. 944). "As artefacts are to theory ... design portfolios are to research programmes" (p. 944). Rams and his ten principles as the worked example (Figure 1, p. 945); the three presentations of Rams's work (the sketch, vitsoe.com, an illustrated essay) and the defining property, "mutually informative juxtaposition ... neither dominates, and neither is subservient" (p. 945). "The converse of Alexander's design patterns" (p. 945). The claim that portfolios allow "the 'extensibility and verifiability'" the opponents asked for (p. 945). Footnote 1: developed with John Bowers.
- Critical angle: (a) the worked example is a product designer annotating his own work with his own principles; who annotates, and can someone else annotate your portfolio, is left open, yet that is what "verifiability" would need; (b) Alexander's patterns were also built from photographed examples, so the distance from pattern languages is smaller than "converse" suggests; (c) the paper's own test case is a sketch with a caption saying the form does not matter, so the proposal is a stance, not a format; (d) Gaver says portfolios "allow for" extensibility and verifiability without showing how.

### 07 How the essay argues (1 slide)

Goal: answer "how was the research conducted" for a paper with no study. Method: argument by contrast with two accounts of science chosen "to complicate notions of what makes theories 'scientific'" (p. 938), plus one sketched example. Evidence: quotations from the opponents; examples of design theory drawn mostly from the author's own circle (Gaver [13, 14], Sengers [30], Isbister and Höök [16], Djajadiningrat [6], Zimmerman [36]); one figure. Reflexivity: by the paper's own taxonomy this essay is a manifesto and meta-theory ("this essay, which itself is intended as a contribution to meta-theory," p. 939), and an essay arguing that theory should annotate examples contains exactly one example. The peer-review trace: thanks to reviewer R4 (p. 946). State plainly that this is normal for a CHI essay and that the session title announces it.

### 08 Strengths and limitations (2 slides)

Strengths (slide 11): names the scientism issue openly; keeps its hedges; offers a constructive form rather than only critique; treats disagreement about science as evidence rather than noise ("It is difficult to call for research to be more scientific if what it means to be scientific is under question," p. 941); insists that theory and artefact should stand in balance, which protects both.

Limitations (slide 12), all labelled "we argue": the standard of "productivity" is undefined; there is no account of how a field without shared criteria reviews papers, funds projects or examines students; the convergence section answers "one paradigm" when the opponents asked for "standards"; the many-worlds argument exempts any practice from any criterion; the proposal is sketched, not demonstrated; missing voices: Schön's reflective practitioner is listed as reference 28 but never cited in the text, Pullin on curating design collections is listed as reference 25 and never cited, Cross's "designerly ways of knowing" is absent, and Koskinen et al.'s 2011 book "Design Research Through Practice" (lab, field, showroom), a contemporaneous attempt at exactly the kind of account Gaver wants, is not mentioned (verify publication date and content before use). Who benefits: an established studio with a recognisable portfolio loses little from the absence of criteria; a student or a newcomer who must get published and examined needs legible criteria; "consensus can come to look like constraint" (p. 946) reads differently from a PhD seat.

Reference-list note (checked 2026-10-19 against the PDF text): in-text citation [3] is used for Newton and the moon (p. 939), for "Reprinted in [3]" after Popper, Lakatos and Kuhn 1977 (p. 946) and for "Carroll and Kellogg [3]" (p. 944), but reference 3 in the list is Collins 1994; the "c.f. [3]" on p. 943 correctly points to Collins. The Newton point and the three reprints belong to Curd and Cover, reference 5; Carroll and Kellogg ("Artifact as theory-nexus," CHI 1989, verify) is missing from the list altogether. A copy-editing slip, not a flaw in the argument; use it only as a light "watch out" or as a gentle aside about rigour, never as a gotcha.

### 09 Connections (1 slide, cross-checked by a second agent)

Goal: place the essay against the course readings and let each change how we judge it. Candidates, all to be verified against the course copies before use (Week 4's section 10 logs hold URLs for the first four):

- **Wobbrock and Kientz 2016**, seven contribution types: this essay is an opinion contribution; artefact contributions "do not have to be" accompanied by studies and "sometimes should not be" (their words, citing Greenberg and Buxton; verify page); methodological contributions need "repeated validation" (p. 41), which is the demand Gaver resists. Where would an annotated portfolio sit among the seven types?
- **Greenberg and Buxton 2008**: the same target (evaluation as dogma, convergence "on that which we are already familiar with," p. 115) and a milder cure (they still propose other forms of validation, including critique and case studies, p. 118). Gaver goes further: not a different validation, but theory as annotation.
- **Olsen 2007**: explicit criteria for systems research. A counterexample to Gaver's worry (standards for a non-science that are not scientistic), or an instance of what he fears? Both readings defensible.
- **Weiser 1991**: tabs, pads and boards as a portfolio with a vision as annotation; ubicomp as a manifesto that "created a world" in Gaver's iPad sense (p. 943).
- **Wobbrock, Morris and Wilson 2009** (Week 4): the empirical model Gaver contrasts design with. Yet by Lakatos's test it founded a programme: about 216 later elicitation studies (DIS 2020 figure from the Week 4 logs; verify before reuse). An empirical paper passes Gaver's own fertility test.
- **Ishii and Ullmer 1997, Tangible Bits** (Week 5): a vision statement plus metaDESK, transBOARD and ambientROOM. Arguably the clearest annotated portfolio in the course. Check how the paper names its contribution.
- **Le Goc et al. 2016, Zooids** (Week 5): artefact plus design space plus open hardware. Would Gaver count it as research through design? What annotation does its design-space section do?
- **Breazeal 2003, Toward Sociable Robots** (Week 6): Kismet as an ultimate particular and a theory nexus.
- **Oliver's own project**: a two-condition study of trust in a CISO briefing. By Gaver's account it is the model research through design is not. A personal hook for the discussion, if Oliver wants it.

### 10 What followed, 2012 to 2026 (1 slide)

Goal: did the field go Gaver's way, Zimmerman's way, or both? Candidates, each to be verified with a URL before use and dropped if not: Gaver and Bowers, "Annotated Portfolios," Interactions 2012; Bowers, "The logic of annotated portfolios," DIS 2012; Löwgren, "Annotated portfolios and other forms of intermediate-level knowledge," Interactions 2013; Höök and Löwgren, "Strong concepts," TOCHI 2012 (the direct alternative: intermediate knowledge that does generalise); Zimmerman and Forlizzi, "Research Through Design in HCI," 2014 chapter; Dalsgaard and Dindler, "Between theory and practice: bridging concepts," CHI 2014; the Research Through Design (RTD) conference, first held 2013, where artefacts are exhibited; DIS Pictorials from 2014 as a publication format for visual argument; Pierce, "On the presentation and production of design research artifacts in HCI," DIS 2014; Bardzell, Bardzell, Dalsgaard, Gross and Halskov, "Documenting the research through design process," DIS 2016; Frauenberger, "Entanglement HCI," TOCHI 2019; Beaudouin-Lafon, Bødker and Mackay, "Generative theories of interaction," TOCHI 2021. Three open questions to close on: whether CHI reviewing moved toward or away from Gaver; whether generative AI tools, which make "many worlds" cheap, strengthen or weaken the annotated-portfolio idea; and whether the portfolio form favours labs with resources.

### 11 Discussion bridge (3 slides)

- Slide 15, our reading, candidate thesis for Oliver to accept or change: **Gaver wins the argument against scientism and loses the argument against standards. The annotated portfolio is a form, not a criterion, and the essay never says how a field that refuses shared criteria decides what to publish, fund and examine.** Three supports: productivity undefined (p. 941); the opponents asked for standards, the essay answers "one paradigm" (p. 941 to 943); the worked example is self-annotated and sketched (p. 945).
- Slide 16, tensions with more than one defensible answer: (1) are standards repressive or enabling, and for whom; (2) is "sometimes right" theory worth publishing, and what would count as evidence of its fertility; (3) which of the three convergence readings describes HCI as this class experiences it; (4) would you accept an annotated portfolio as a CHI paper; (5) does the many-worlds argument apply to your own project. Smaller items: the uncited Schön; the self-annotating Rams; the one-example essay.
- Slide 17, over to you: classmate quotes with names, then the plan for the half hour.

## Discussion plan (to be written as `slides/12-Discussion-Guide.md` in the Week 4 pattern)

Anchor candidates, each to get an open stem, why it matters, three to five positions with one defending the paper, zoom-in probes with pages, zoom-out probes, a move, threads, a three-minute version, traps and a [classmate] hook:

1. **Whose interest do standards serve?** Positions: Gaver (self-policing, scientism); Zimmerman et al. (legitimacy, cumulative work); a PhD student who needs legible criteria; a reviewer who must say yes or no to a portfolio.
2. **Is "sometimes right" theory worth having?** What evidence of fertility would you accept, and who counts it?
3. **Which reading of non-convergence fits this course?** Three-way hands: pre-paradigmatic, invisible consensus, many worlds. Then: does the list of what "most of us" share (p. 942) describe your own practice?
4. **Would you accept an annotated portfolio as a CHI paper?** Reviewer role-play with a real candidate (Tangible Bits or Zooids). What would you ask the authors to add?
5. **Does the many-worlds argument hold for your project?** Name the single world your project assumes.
6. **Fourteen years on, who won?** RTD conference and Pictorials versus the growth of quantitative studies at CHI (facts to verify in section 10).
7. **Reserve: is this essay an annotated portfolio with one item?** Does it meet its own standard?

Activities:
- **Warm-up (5 min):** three-way hands on "Is your own course project research through design? Yes, no, depends what counts." Two or three one-sentence justifications from the "depends" corner. These are the room's own data for Anchors 3 and 5.
- **Pairs (5 min), role cards:** one question and two roles per card, as in Week 4. Candidate cards: Gaver versus Zimmerman on standards; a reviewer versus an author on a portfolio submission; a student versus a studio head on who benefits from no criteria; Popper versus Lakatos on a claim from Tangible Bits; an artist versus a scientist on "many worlds."
- **Live annotated portfolio (optional, 6 min):** five course artefacts on one slide (Weiser's pads, the Surface gesture set, metaDESK, Zooids, Kismet). Pairs write one dimension on which they vary and one invariant. Harvest. This tests whether annotation across different designers' work does anything, since Gaver's portfolios come from one designer or studio.
- **Popper test (optional, 3 min):** take one claim from a course paper and try to state it falsifiably in 90 seconds. Discuss why it is hard and whether that matters.
- **Close:** one-sentence round, capped at eight to ten voices: "One thing I now expect from research through design."

Grading on the day: the Week 4 marks pattern (cited the paper with a page, cited the talk with a section, acknowledged the other side's strongest point) with one addition for this paper: applied the argument to their own project.

## Working method (same as Week 4)

- Folder: `week7/` holds the PDF, the instructions, this file, `slides/` (00-Overview, 01 to 11 section files, 12-Discussion-Guide, `_agent-briefs.md`, `_build-deck-source.py` adapted from Week 4), `reports/` (agent reviews), `assets/` (Figure 1 render). Generated files: `Week 7 - Deck Source.md`, `Week 7 - Deck Source - Presentation only.md`, `Week 7 - Discussion Guide.md`, `Week 7 - Discussion Cards.md`. `Presentation.md` (empty, 2026-10-18) is Oliver's; leave it.
- One Markdown file per section, slide content and speaker notes together, verification log and open items at the end; the Week 4 file format in `Presentation-Week4/slides/_agent-briefs.md` section 1.4.
- Writers: one Opus agent per section in parallel from a saved briefs file; the orchestrator writes 00, 03 and 11 and checks every file. Sections 09 (connections) and 12 (discussion guide) get an independent Opus cross-check; 08 (limitations) should too, since every limitation is an argument that can be overstated.
- No em-dashes anywhere (check with `grep -c $'\xe2\x80\x94' *.md`). Every quotation exact, with its printed page. "Gaver says" for his claims, "Gaver quotes" for what he quotes from others, "we argue" for ours. Keep his hedges.
- Spoken scripts at or under 130 words per minute of slide time. On-slide text about 90 words at most.
- [classmate] slots in sections 02, 05, 06, 08, 11 and in every anchor, filled after the commentaries are posted; credit by name.

## Pitfalls list for the briefs (first draft)

1. Gaver does not say design is unscientific; he says it depends on the account of science (p. 941). Do not flatten it.
2. The opponents asked for standards, protocols and extensible theory (p. 937). They did not ask for design to become a science. Quote them, do not paraphrase them into a stronger claim.
3. "Provisional, contingent, and aspirational" (p. 937, 941) are Gaver's words; "extensible and verifiable" are Zimmerman and Forlizzi's (p. 937). Never swap the attributions.
4. Lakatos's test is novel facts and productivity (p. 940), not confirmation. Gaver's point is that confirmations show fertility, not truth (p. 941).
5. The "invisible consensus" list (p. 942) is hedged: "most of us," "perhaps we do not all agree." Keep the hedges.
6. "Many worlds" is Law and Urry's phrase (p. 943), applied by Gaver to design. The iPad example is Gaver's (p. 943).
7. "Ultimate particular" is Stolterman's (p. 944); "theory nexus" is Carroll and Kellogg's (p. 944); "ontological politics" is Law and Urry's (p. 943).
8. The Rams principles in Figure 1 are four of ten; the paper names two in the text ("unobtrusive," "thorough down to the last detail," p. 944).
9. The paper proposes annotated portfolios as one serious form, not the only one: "I am not proposing here that annotated portfolios subsume all other forms of design theory" (p. 945).
10. The reference list has the [3] and [5] slip and the missing Carroll and Kellogg entry (see section 08). Cite Carroll and Kellogg by name and year only after verifying the 1989 paper.
11. Reference 28 (Schön) and reference 25 (Pullin) are never cited in the text. Say "listed but not cited," not "ignored."
12. Session title and conference: Critical Perspectives on Design, CHI 2012, Austin, 5 to 10 May. Pages 937 to 946.

## Open questions for Oliver

1. Presentation date: answered, 2026-10-22.
2. Only paper: answered, yes.
3. Confirm the readings to connect to: Weeks 1 to 3 (Weiser 1991; Wobbrock and Kientz 2016; Greenberg and Buxton 2008; Olsen 2007), Week 4 (Wobbrock, Morris and Wilson 2009), Week 5 (Zooids; Tangible Bits), Week 6 (Toward Sociable Robots; Slow Robots for Unobtrusive Posture Correction). Anything else, for example Zimmerman, Forlizzi and Evenson 2007, which would change section 01?
4. The candidate thesis on slide 15: accept, sharpen or replace.
5. Own project in the discussion: answered, keep it out.
6. Question sets 13 to 15 again (paper-only, talk-only, synthesis), or only the discussion guide? Week 4 ended with the guide as the plan and the sets as probes and grading aids.
7. The live annotated-portfolio exercise: in or out? It costs about six minutes of the discussion.

## To do

- [ ] Oliver answers the open questions
- [ ] Write `slides/_agent-briefs.md` (common rules, pitfalls, file format, one brief per section)
- [ ] Render Figure 1 to `assets/`
- [ ] Run the writers in parallel; orchestrator writes 00, 03, 11
- [ ] Verify every file; cross-check 08, 09
- [ ] Adapt `_build-deck-source.py`, build the deck source
- [ ] Write `12-Discussion-Guide.md`; two independent cross-checks
- [ ] Read the commentaries, fill the [classmate] slots
- [ ] Build the deck, submit the draft by the day before, 11:59 pm
- [ ] Commit

## What was done on 2026-10-19 (evening)

1. **Briefs written** (`slides/_agent-briefs.md`, about 9,300 words): common rules, a quotation bank of every passage to be quoted with its printed page, 14 pitfalls, the file format, style rules, one brief per section, a brief for the readings report, and three cross-check briefs.
2. **Eight Opus writers run in parallel** for sections 01, 02, 04, 05, 06, 07, 08, 10, plus a ninth Opus agent for the readings report (31 connections across the nine earlier course readings, every quotation checked against the reading; both Week 6 papers read in full, Breazeal from page images). The orchestrator wrote 03 and 11, read and checked every file, and adapted the build script.
3. **Section 09** written by an Opus agent from the readings report, with the alternates listed for Oliver to choose from.
4. **Cross-checks:** 08 by an independent Opus reviewer (31 quotations and 6 outside references confirmed, no blockers, 7 should-fix and 5 nice-to-have applied); 09 by another (20 of 20 quotations exact, alternates complete; one blocker, a false date claim in the transition, and six should-fix applied); the whole deck by a third (44 cross-file quotations compared, no conflicts of meaning; three contradictions, one pitfall slip and a misquoted slide title fixed; no transition, dash or timing failures).
5. **Corrections to the plan found by the writers and reviewers:** Figure 1 has seven photographs, not eight, and its notes are Gaver's sketch; the "c.f. [3]" on p. 943 is correct, so the reference slip is narrower than pitfall 10 said; Gaver names Lakatos's criterion ("relative productivity," p. 941); Gaver accepts articulated standards (p. 941); Höök and Löwgren 2012 postdates CHI 2012; Gaver cites Sengers 2006 on scientism; "own circle" is not supported. All applied; listed in `slides/00-Overview.md`.
6. **Mechanical checks** pass: zero em-dashes or en-dashes in every file; every Say script within 130 words per minute of its slide.

## What was done on 2026-10-20

1. **Depth pass at Oliver's request** ("enough information for every slide ... additional insight so I can extend if someone asks"): Part 4 added to `slides/_agent-briefs.md` with must-cover items per section; eleven Opus agents in parallel added an "If asked" block under every slide and a "Background for the presenter" section to every section file without changing existing text; two independent Opus reviewers checked the additions (`reports/opus-DCHECK-A.md`, `-B.md`); fixes applied. The build script now drops the presenter material from the presentation-only copy.
2. **Reference validation at Oliver's request** ("validate all references which we have, and make note of the ones which are wrong"): `reports/opus-REF-paper-references.md` (Gaver's 36 entries and every in-text number; 11 entries wrong, 1 unverifiable, plus in-text slips, listed in `slides/00-Overview.md`) and `reports/opus-REF-deck-references.md` (every outside reference in our files). The newly confirmed slip that touches the deck, [34] for [33] on p. 939, is recorded in sections 01, 03, 04 and in pitfall 10.
3. **Results:** depth checks found no factual error on any slide; in the added material, one blocker (a miscount of Gaver's uses of "scientism") and 21 rewordings, all applied. The deck-reference check found 97 of 101 outside references correct, one wrong (SonicFinder issue number, fixed) and three repeats of the paper's own slips, now flagged.
4. Small corrections from the depth agents: Kuhn's text checked, so the "must founder" line in 05 is now stated as Gaver's stronger wording; "commercial products" in 08 changed to "products for Braun and Vitsoe" (one item in Figure 1 was never produced); the annotator of Figure 1 is unnamed (11).

## What was done on 2026-10-21: version 2 of the deck (`slides2/`)

Oliver found version 1 complete but dull ("utterly boring while it might have importance") and asked for a `slides2/` folder with a simpler, more appealing way to carry the same information. Written by the orchestrating session from the paper text and the verified v1 files; every fact and quotation reused from v1 or the paper, with the source file named in each log.

1. **Structure.** Twelve slides instead of seventeen, 22:00 of talk. One question runs the whole talk: "Is this research?", asked about the Drift Table on slide 1 (hands), answered by the critics' own citation of it as an exemplar, and asked again on slide 12. Three audience moments: hands (slide 1), a three-way vote on the three readings of non-convergence (slide 6), and the Rams test, principles only versus products only versus both (slide 7). Headlines are sentences; pages moved to the notes; under about 70 words per slide. Run sheet and rationale in `slides2/00-Overview.md`.
2. **Files.** `slides2/01-Hook.md` to `09-Handover.md`, same format as v1 (On the slide, Say, Transition, Watch out, If asked, Sources, Verification log, Open items). `slides2/_build-deck-source.py` writes `Week 7 - Deck Source v2.md` and `Week 7 - Deck Source v2 - Presentation only.md`; `slides2/_check-sections.py` checks dashes, words per minute and slide density. All Say scripts within 130 words per minute; zero dashes.
3. **Assets.** `assets/drift-table-chi2004-fig1.png`: Figure 1 of the CHI 2004 Drift Table paper, cropped from the author's version in the Goldsmiths repository (https://research.gold.ac.uk/id/eprint/4721/1/p885-gaver.pdf). `assets/rams-fig1-annotated.png` and `assets/rams-fig1-products-only.png`: Figure 1 of Gaver 2012 as printed and with the notes masked, for the three columns of slide 7.
4. **Corrections found while sourcing the photo.** In 2004 Gaver was at the Royal College of Art, not Goldsmiths (the CHI 2004 paper's affiliation line), so v2 says "Gaver and colleagues," never "the studio," for the Drift Table. The aerial photography is "of England and Wales." Both recorded in `slides2/01-Hook.md`.
5. **Cross-check.** One independent Opus reviewer checked every quotation, page, paraphrase, attribution, transition and audience instruction in `slides2/` against the paper text and the v1 files: `reports/opus-V2-deck-check.md`. Result: 120 quotations checked, 96 pages all right, all 11 transitions fit. Five blockers ("extensible and verifiable" given to the critics on slide 2 when it is Gaver's own phrase, p. 938; the manifesto list on slide 5 naming ludic design where p. 941 names Designing for the Self; a caution on slide 7 attributed to the paper that is ours; "follow-up studies" for the 216 elicitation studies of a 2020 review on slide 10; an overstated Shin et al. note), 21 should-fix items (mostly dropped hedges such as "might," "may," "seem to," and one motive stated as fact) and the nice-to-have items were all applied on 2026-10-21, including explicit click-and-hands cues for the Rams columns on slide 7. Checks re-run: zero dashes, every Say within 130 words per minute, 22:00 total.

6. **Second review round, at Oliver's request ("spawn an OPUS agent or two to validate the slides ... swap the order of your instructions").** Two further independent Opus reviewers, neither shown the first report, with the same checks in opposite order: A read the deck as the audience first and checked the paper last (`reports/opus-V2-check-A.md`); B read the paper in full and listed every hedge before opening a slide (`reports/opus-V2-check-B.md`). Each checked 141 quotations, all pages right. Agreement on every blocker: the slide 5 sentence credited to Lakatos (Gaver's own words from a Lakatosian view); the slide 9 twist ("a named method with its own conference," when the RTD conference is for research through design generally); an unsourced claim on slide 2 that design research was reviewed by people who could not judge it (v1 had said never to give the critics a motive); the Shin et al. card on slide 10 (the logic of "inconceivable" was backwards: Shin et al. managed the realistic half, not the unambiguous half); the method line on slide 3 (covered only part one of the essay); and the missing hands call on the third Rams column. A alone caught "Productive programmes" for "productive research programmes" and "funding" put in the critics' mouths; B alone caught the dropped "insofar as," "tends to be" and "in several respects," and that the vote on slide 6 should ask about design research in HCI, not HCI. Applied the same evening: all blockers, all 39 should-fix items and most nice-to-have items from both, plus Gaver's own three conclusions spoken on slide 11, Gaver's "Personally, I tend to think" line (p. 941) in the notes, and slide 7 at 3:00. Checks re-run: zero dashes, every Say within its limit, 22:30 of sections plus the title.

7. **Slides-only file, 21 slides (Oliver's request: "write down the slide information very precisely ... no distractions ... like 20 slides in total").** `slides2/slides-only/Week 7 - Slides v2.md`: for each slide the title, the body text bullet by bullet, pull quotes with attribution, layout, image file and click order; no notes, logs or times. The twelve section slides became twenty plus the title by splitting every slide that carried two ideas: the fight into the critics' timeline (slide 3) and Gaver's reply (4); the judo move into the move (5) and a map of the essay's three moves with page ranges (6); the puzzle into three diagnoses (9) and the vote with the Fix tags built after the count (10); the Rams test into three question slides, principles only, products only, both (11 to 13); the flip into object-as-fact (14) and point-to-portfolio (15); the aftermath into the scoreboard (16) and the twist with two open questions (17); the verdict into Gaver's own three conclusions (19) and ours (20). Mapping to the section slides, with the time split (23:00 total), is in `slides2/00-Overview.md`. Two independent Opus checks the same evening: the write-down against the section files (`reports/opus-SO-writedown-check.md`: 54 source bullets, none missing, 20 additions all traced; seven note-derived lines were unlisted in the mapping, now listed) and every quotation and reference (`reports/opus-SO-references-check.md`: 33 paper quotations exact, 8 pages right, 26 outside references confirmed, no blockers; three attribution placements tightened: "must founder" marked as Gaver's words after Kuhn, "most of us" as Gaver's not Latour's, "caption" as our word for his "annotation"). Also applied: the slide 19 quotation subject ("attempts to establish disciplinary norms"), an ellipsis on the slide 4 self-policing quotation, image credits on slides 1, 2 and 12, the vote slide split into card and tag bullets, and label words moved out of the slide text.

8. **Version 3, `slides3/` (Oliver's outline, full notes, extractable slide text).** After reviewing v2 Oliver said he "could just leave the slides alone and leave the room" and gave a new outline: intro; how the mess arose historically and how it showed itself; the methods prior to the paper; the method Gaver proposes, its upsides, critical reflection; what the paper changed; its value in 2026. The orchestrator completed it to 22 slides (scene 2 to 5; before the paper 6 to 8; Gaver's argument 9 to 13; judging 14 to 16; afterwards 17 to 20; discussion 21 and 22), wrote `slides3/_agent-briefs.md` (rules, exact file format, 18 pitfalls consolidated from v1 and v2, per-slide must-haves and sources), `00-Overview.md`, and three scripts (`_check-sections.py`, `_build-deck-source.py` emitting `Week 7 - Deck Source v3.md` and the presentation-only copy, `_extract-slides-only.py` emitting `slides3/slides-only/Week 7 - Slides v3.md`, the slide text alone). Three Opus writers in parallel wrote the six section files (slides 2 to 8 with a research pass on Fallman 2003; 9 to 13; 14 to 22). Two Opus cross-checks in parallel: references (416 quotation instances, 415 exact; 73 outside facts, 72 confirmed; three blockers: a wrong page on slide 3, a stale "runs the studio at Goldsmiths" in the title script, missing inner quotes on slide 8) and flow against Oliver's outline and the brief (every must-have present; four blockers: no causal line between design entering HCI and the manifestos, no home for "how was the research conducted", a click mismatch in the Rams test, slide 20's counting only in the notes; plus duplicate transitions, repeated content and about 250 words of timing cuts). All applied; Cross 1982 taken off slide 3; slide 20 lengthened to 1:15. Final: 22 slides, talk 26:15 of caps, about 3,200 spoken words, zero dashes, every script within its limit. Reports: `reports/opus-V3-references-check.md`, `reports/opus-V3-flow-check.md`.

Still open for v3 (and v2, v1): the thesis (v3 slide 16), the [classmate] slots (slides 11 and 12), the discussion guide (`slides/12-Discussion-Guide.md`, serves both decks), and the draft submission by 2026-10-21, 11:59 pm. Oliver decides which deck to submit.

## Decisions made by Oliver

| Decision | Date |
|---|---|
| Oliver presents alone; no co-presenter | 2026-10-19 |
| Only one paper this week; presentation on 2026-10-22, draft due 2026-10-21 11:59 pm | 2026-10-19 |
| 17 slides is the right size to start | 2026-10-19 |
| The full layout is written out now, at high quality, so that reading it once is enough and iteration is the exception | 2026-10-19 |
| An Opus agent mines the earlier course readings for candidate connections so Oliver can choose which to use (report R) | 2026-10-19 |
| Oliver's own CISO trust study stays out of this presentation; it belongs to another week | 2026-10-19 |

## 9. Question day, 2026-10-22 (morning before the talk)
- Oliver's questions on the deck, the paper and its history answered through briefed Opus agents (follow-ups to the same agent) and checked against `assets/paper-text.txt`; full record in `Week 7 - Question Log 2026-10-22.md` (34 entries, corrections, open decisions).
- New files: `slides3/extra-3b-Waves.md` (optional three-waves slide, 1:15, checks pass); `assets/frayling-1993.pdf` and `-text.txt` (Frayling's original, OCR); `assets/harrison-tatar-sengers-three-paradigms-text.txt`; `slides3/Role Cards v3.md` (role-card pack, written by an agent, committed separately when done).
- Open: slide 10 labels (three options offered), slide 12 coffee-maker opener, two Say wording fixes (slides 5 and 10), Bowers dagger on slide 17. None applied; see the log.
