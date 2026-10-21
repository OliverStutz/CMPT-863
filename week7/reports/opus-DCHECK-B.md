# D-CHECK-B: accuracy check of the depth additions, files 07 to 11

Reviewer: independent Opus agent, 2026-10-20. Scope: every "If asked" block and every "Background for the presenter" item in `07-Method.md`, `08-Strengths-Limitations.md`, `09-Connections.md`, `10-What-Followed.md`, `11-Discussion-Bridge.md`. Existing slide, Say, Watch out and Sources text was read for contradiction checks only.

How checked:
- Gaver 2012: `pdftotext` per page of the ground-truth PDF (PDF page N = printed 936 + N), phrase search with hyphenation and quote marks normalised; Figure 1 checked on `assets/p9-09.png`.
- Ways of Knowing in HCI (Olson and Kellogg 2014): book PDF https://people.cs.vt.edu/~gangwang/HCI.pdf, downloaded and read. Gaver's chapter is PDF pp. 152 to 174 (printed 143 to 165). Zimmerman and Forlizzi is printed pp. 167 to 189.
- Course readings: Olsen pre-print (Wayback, 7 pp.), Greenberg and Buxton (UMass PDF), Wobbrock and Kientz (UW PDF), DIS 2020 review (UW PDF), Shin et al. (Wayback KAIST PDF), Tangible Bits and Week 4 paper (local PDFs). Every quoted phrase was searched page by page.
- Other sources: Cross 1982 (Readings.design PDF), Zimmerman, Forlizzi and Evenson 2007 (Berkeley PDF), Löwgren 2013 (DiVA author's version), Koskinen and Krogh 2015 (ijdesign.org), Crossref reference lists and counts, OpenAlex abstracts and counts, Semantic Scholar count, CHI 2012 subcommittee and Session149 pages, CHI 2026 subcommittee page and first-round blog post, RTD Roams and conference-series pages, DIS 2026 Pictorials call, CORDIS 226528, Stanford Encyclopedia "Imre Lakatos", Allen and Tanner 2002 (PMC).
- Dashes: grep for U+2014 and U+2013 across all five files: none found.

Totals: 112 items checked. 99 keep (some with optional nice-to-have notes), 12 reword (should-fix), 1 blocker (also a reword). 0 drop.

---

## 07-Method.md

| # | Slide | Question or item | Where checked | Verdict |
|---|---|---|---|---|
| 1 | 07.1 | What kind of paper, how reviewed | p. 937 ("In this essay, I explore", "The essay has three parts"), p. 938 ("In this essay, my intention"), p. 946 ("reviewers", "the anonymous R4"); chi2012.acm.org/cfp-selecting-subcommittee.shtml (8 subcommittees; Design remit "... critique, and theory") | Reword (minor): the paper says "essay" three times, not twice. See S1 |
| 2 | 07.1 | Session contents | chi2012.acm.org/program/desktop/Session149.html: 8 May 2012, chair Peter Wright, Newcastle; four papers as listed | Keep |
| 3 | 07.1 | Why Popper and Lakatos | p. 939: "can also shed valuable light on design [2]", "sophisticated and well-developed discussions of the nature of theory", demarcation; p. 938; ref [2] p. 946 | Keep. Nice: Gaver names Philosophy of Science "and Science and Technology Studies (STS)" and calls them "a useful foil" (p. 939), which matches the answer's "foil" |
| 4 | 07.1 | Cherry-picking | p. 941 hedge; p. 939, 941 to 942 Kuhn | Keep |
| 5 | 07.1 | What would change his mind | p. 940 "in principle a falsifiable statement", "comparable"; p. 945 "greater consensus may be emerging" | Keep |
| 6 | 07.1 | Conflict of interest | p. 942 "others (including myself)"; p. 937; refs p. 946 | Keep |
| 7 | 07.1 | Did Gaver publish the portfolio | Crossref 10.1145/2212877.2212889; OpenAlex abstract of Bowers 10.1145/2317956.2317968 (nine pieces, Feyerabend); p. 944 footnote 1 | Reword: footnote 1 says the ideas "have been developed in collaboration with John Bowers", not that it credits Bowers with them. See S2 |
| 8 | 07.1 | Did he come back to the science question | Chapter read: p. 147 ("Science is defined by epistemological accountability ... Design, in contrast, works with aesthetic accountability"); footnote 2 p. 145 cites Gaver 2012; Koskinen and Krogh 2015 quote verified on ijdesign.org | Reword: the answer says "We have not read the chapter itself" and relies on secondary summaries; the chapter is available and section 10 already quotes it. Content is correct. See S3 |
| 9 | 07.1 | How influential | api.crossref.org: 880; Semantic Scholar: 1,223 (both re-run 2026-10-20) | Keep. Nice: add OpenAlex 1,139, the figure section 10 uses, so the two files show the same index |
| 10 | BG | The venue | CHI 2012 cfp page: "over 1,800 Papers and Notes", 8 subcommittees, Design chairs DiSalvo, Forlizzi, Wakkary, Harrison; Jeff and Shaowen Bardzell among ACs; refs 9, 10, 33, 34, 35 p. 946 | Keep |
| 11 | BG | The session | Session149 page; OpenAlex Brynjarsdottir abstract ("Drawing on critical sociological theory of modernism"), pp. 947 to 956; Pierce pp. 957 to 966 | Keep. Nice: "analyses persuasive technology as a modernist approach to social problems" is a loose paraphrase; closer: "critiques persuasive sustainability research from 2009 to 2011, drawing on sociological theory of modernism" |
| 12 | BG | Studio methods | Probes interactions 6(1) 21 to 29 (also cited in Gaver 2014 refs); workbooks CHI 2011 pp. 1551 to 1560 = ref [12] p. 946; Datacatchers TEI 2017 pp. 243 to 251, "batch production of 130 of the devices" (OpenAlex abstract) | Keep |
| 13 | BG | Annotated portfolio follow-ups | Bowers abstract: "it is suggested that design research might build its own 'limited rationality' rather than be brought in line with supposed norms" | Reword: "argues that design research should develop" drops Bowers's hedge ("might"). See S4 |
| 14 | BG | The sequel | Chapter p. 147; Koskinen and Krogh 2015 | Reword: "Secondary summaries" is no longer needed and the chapter's own words are stronger. See S5 |

## 08-Strengths-Limitations.md

| # | Slide | Question or item | Where checked | Verdict |
|---|---|---|---|---|
| 15 | 08.1 | Where Gaver uses "scientism" | grep of all ten pages: p. 938 line 7, p. 941 ("subverts any easy scientism in HCI"), p. 945 ("HCI's tendencies towards scientism"), p. 946 (R4) | **Blocker.** The answer says "Twice"; the word appears four times. The answer two items below (item 21) itself quotes the p. 945 use. See B1 |
| 16 | 08.1 | Reviewer comment as weak evidence | p. 946 "in part" | Keep |
| 17 | 08.1 | Who raised scientism before | ref [29] p. 946; p. 941; thanks and refs [2], [30] p. 946 | Keep |
| 18 | 08.1 | Hedges make it irrefutable | p. 941; p. 942 | Keep |
| 19 | 08.1 | Annotated portfolio in one sentence | p. 944 "a balance is achieved between ..."; p. 945 caption; "the converse of Alexander's [1] design patterns" p. 945 | Keep |
| 20 | 08.1 | Constructive versus sketch | p. 944 to 945 (Design Museum, vitsoe.com); footnote 1 p. 944 | Keep |
| 21 | 08.1 | Disagreement as evidence | p. 941; p. 945 "continuing controversies ... towards scientism" | Keep |
| 22 | 08.1 | Balance rule versus "definite facts" | p. 945 both sentences | Keep |
| 23 | 08.2 | Schön | archive.org record, ch. 3 title; p. 942; ref 28 p. 946 | Keep |
| 24 | 08.2 | Cross 1982 | Cross PDF: three cultures, "a concern for 'appropriateness'" (pp. 221 to 222); "We need a 'research programme', in the sense in which Lakatos has described", core "there are designerly ways of knowing" (p. 226) | Keep |
| 25 | 08.2 | Pullin | OpenAlex abstract: "conceived as an entity at the time of creation", "Collections can illuminate themes and imply a landscape of alternatives that transcend any individual piece"; title names Social Mobiles, Six Speaking Chairs | Keep |
| 26 | 08.2 | Koskinen et al. not on slide | Elsevier page (Sept 2011); refs [9], [33 to 35] p. 946 | Reword (consistency): "months before CHI 2012" sits awkwardly with the Watch out's "around the CHI 2012 submission deadline". See S6 |
| 27 | 08.2 | Strong concepts as middle ground | Crossref abstract 10.1145/2362364.2362371 (quotes exact; TOCHI 19(3), Oct 2012); Löwgren 2013 | Keep |
| 28 | 08.2 | Lakatos's judge | p. 941 "the relative productivity of competing programmes as the grounds for scientists' choice amongst them"; p. 945 "thriving" | Keep |
| 29 | 08.2 | Did Gaver answer the reviewing question | Chapter p. 147 ("The basic question here is 'does it work?'"; process questions "are not grounds in and of themselves for judging a design successful or unsuccessful") | Reword: "We have not read the 2014 chapter in full" is now avoidable, and the chapter has a direct, quotable answer. See S7 |
| 30 | 08.2 | "Who benefits" as motive attack | p. 938 ("their students"), p. 937 to 938 | Keep |
| 31 | 08.2 | Against our own case | p. 942 | Keep |
| 32 | 08.2 | Rams figure "not an RtD portfolio" | p. 945 caption; Figure 1 (`p9-09.png`): one item is a "prototype portable TV ... never produced"; notes in third person ("Rams redesigned the unit"); author of notes not stated | Reword: "commercial products, annotated by someone other than the designer" states two things the figure does not support; section 11 (Watch out 11.2 and its open item) explicitly says the paper does not say who wrote the notes. See S8 |
| 33 | BG | Schön 1983 | archive.org record; secondary reviews (labelled) | Keep |
| 34 | BG | Cross 1982 | Cross PDF pp. 221, 226: five aspects verbatim, RCA report, Lakatos programme | Keep |
| 35 | BG | Pullin and Koskinen | OpenAlex abstract: "Several design collections are introduced within which the author was involved, in the context of collections by Hussein Chalayan, Dunne & Raby, the Royal National Institute for Deaf People, and the Museum of Jurassic Technology" | Reword: "drawing on projects with Hussein Chalayan, Dunne and Raby ..." implies Pullin worked with them; the abstract places his collections "in the context of" theirs. See S9 |
| 36 | BG | Strong concepts | Crossref abstract; Löwgren 2013 Crossref | Keep |
| 37 | BG | Sengers 2006 and Gaver 2014 | ref [29]; chapter p. 147; Koskinen and Krogh | Reword (consistency with 10): use the chapter's own words. See S10 |
| 38 | BG | "Who benefits" is ours | p. 937 to 938; p. 946 | Keep |

## 09-Connections.md

| # | Slide | Question or item | Where checked | Verdict |
|---|---|---|---|---|
| 39 | 09.1 | Anachronism | G&B p. 119 "Gaver et. al. are developing methods that probe cultural reactions and technology uptake by niche cultures [11]", ref 11 = Gaver, Dunne, Pacenti 1999; W&K endnotes 4, 5 p. 44; Olsen thanks Saul Greenberg p. 7 | Keep. Nice: "Gaver's studio was in this debate by 2008" overstates a citation; say "Gaver's studio's methods were already cited in this debate by 2008." Also W&K endnote 1 is Olson and Kellogg, Ways of Knowing in HCI (2014), the book that prints Gaver's 2014 chapter: a real, checkable link |
| 40 | 09.1 | Olsen generality a stretch | Olsen pre-print p. 3 "not a design tool but an application"; p. 4 generality passage | Keep |
| 41 | 09.1 | Olsen on science | pre-print p. 6 "like any other science, moves forward"; p. 7 "are not novel but recently have been out of favor" | Keep |
| 42 | 09.1 | Tangible Bits as RtD; optical metaphor | TB p. 240 "somewhat weaker"; Gaver p. 945; W&K p. 41 "Envisionment" with Ishii and Ullmer | Keep |
| 43 | 09.1 | G&B on the other side | G&B p. 118 both quotations | Keep |
| 44 | 09.1 | W&K answer Gaver | W&K p. 41 both quotations; endnotes p. 44 (no Gaver) | Keep |
| 45 | 09.1 | Portfolio among contribution types | W&K p. 40 artifact quotation; Figure 2 p. 44: Essay/Argument 88 submitted, 19 accepted, 21.6%; Theory 142, 25, 17.6%; overall 23.6% | Keep |
| 46 | 09.1 | Wobbrock row a straw man | DIS 2020 p. 863 "rarely applied", legacy bias | Keep |
| 47 | 09.1 | 216 studies meaning | DIS 2020 p. 855 counts; p. 859 "36% of the studies", "likely explained by the fact that Wobbrock et al. [123] employed 20 participants"; Gaver p. 941 exemplar | Keep |
| 48 | 09.1 | Shin refutes "inconceivable" | Shin p. 5 (day 2 only, "unwanted learning effects"); p. 7 "except P2" | Keep |
| 49 | 09.1 | Blaming the setup | Shin p. 6: "We only gave neutral explanations about the study ...: 'The monitor will move to support your health while working'" | Reword (fairness): the authors themselves call that instruction "neutral"; a classmate who read Shin can say so. See S11 |
| 50 | 09.1 | Weakest link | p. 41, p. 113, Gaver p. 940 | Keep |
| 51 | BG | Olsen | pre-print pp. 1 to 7: "articulate criteria" p. 1; usability trap; claim list pp. 3 to 6; "Before all other claims" p. 3; statistical-analysis line p. 3 | Keep |
| 52 | BG | Tangible Bits | TB pp. 235 to 240: metaDESK, activeLENS, passiveLENS, Tangible Geospace (237); Personal Harbor, MIDI (238); SoftBoard, hyperCARDs (239 to 240); quotations p. 240 | Keep |
| 53 | BG | Greenberg and Buxton | G&B p. 112 (a) to (c); p. 113 about 70% and 25%; p. 117 Memex, Sketchpad; p. 118; p. 119 footnote and probes | Keep |
| 54 | BG | Wobbrock and Kientz | W&K pp. 40 to 44: whitepaper, seven types, quoted criteria; Figure 2 shares; "should not be confused" p. 44 | Keep |
| 55 | BG | WMW 2009 and review | Week 4 PDF: 43.2 (p. 1085), 1080 (p. 1086), 57.0% (p. 1087), 48 gestures (p. 1088), 43.5% (p. 1090); DIS 2020 pp. 855, 858, 863, 864 | Keep |
| 56 | BG | Shin et al. | Shin pp. 3 to 7: six students; 1.5 hours; 64.22 s; 1.0 mm/s; 12 participants, mean age 24.55; t = 3.54, p = 0.002; about 66% | Keep. Nice: "four hours on two consecutive days" reads as four hours total; the paper says "four hours on consecutive days (eight hours in total)" (p. 5). Say "four hours on each of two consecutive days" |

## 10-What-Followed.md

| # | Slide | Question or item | Where checked | Verdict |
|---|---|---|---|---|
| 57 | 10.1 | Did later works cite the essay | Crossref reference lists re-run: Höök and Löwgren, Bowers, Pierce, Bardzell et al. cite DOI 10.1145/2207676.2208538; Dalsgaard and Dindler, Frauenberger, Beaudouin-Lafon et al. do not; Löwgren 2013 endnote via DiVA; OpenAlex count 1,139 | Keep |
| 58 | 10.1 | Annotated portfolio in Gaver and Bowers's terms | Löwgren DiVA text: both quotations verbatim; p. 944 | Keep |
| 59 | 10.1 | Löwgren and Bowers disagree | DiVA: "purely indexical and inseparable from their particular artifacts"; "affords abstraction (in the sense of extraction) by others" | Keep |
| 60 | 10.1 | Strong concept | Crossref abstract | Keep |
| 61 | 10.1 | Bridging versus strong | OpenAlex abstract: three constituents, "accountable to both theory and practice", "peepholes", curiosity and engagement | Keep |
| 62 | 10.1 | Z&F answer Gaver | Book PDF: p. 178 "worked to formalize"; five steps all listed on p. 184; p. 175 Drift Table quotation; refs pp. 187 to 189, no Gaver 2012 (Crossref list also lacks it) | Keep |
| 63 | 10.1 | "Repeat" close to Gaver | Book p. 186 quotation; Gaver p. 944 | Keep |
| 64 | 10.1 | Gaver 2014 adds | Chapter p. 147 accountability definitions; "productive indiscipline" pp. 153 and 162, framework warning pp. 162 to 163; Gaver p. 941 | Keep |
| 65 | 10.1 | RTD stopped | researchthroughdesign.org/announcing-rtd-roams/ (23 July 2026; Baltic in Gateshead 2013; Cambridge 2015; Edinburgh 2017; Delft 2019; Covid pause; York, Melbourne, Tokyo 2027) | Keep (the Roams page says only "Delft" for 2019; the file's log cites the 2019 page for "Delft and Rotterdam") |
| 66 | 10.1 | DIS Pictorials | dis.acm.org/2026/pictorials/: 12 pages excluding references, dual anonymous, archival, "same quality of contribution as technical program papers", the images and text sentence verbatim, "introduction at DIS 2014 in Vancouver" | Keep |
| 67 | 10.1 | CHI reviewing moved | CHI 2026 subcommittee page (phrase verbatim); blog post: 6,730 submissions, eighteen subcommittees, Design 485, 184 to RR, 37.9%; overall 38.7% | Keep |
| 68 | 10.1 | Generative AI | OpenAlex abstracts: Luminate "hundreds and thousands of visual and textual outputs", "guiding users towards rapid convergence on a limited set of ideas"; Davis et al. title | Keep |
| 69 | BG | Gaver and Bowers 2012 | Löwgren DiVA (quotes, "influencing autonomous drift", three earlier pieces) | Keep |
| 70 | BG | Bowers 2012 | OpenAlex abstract (all quoted phrases verbatim) | Keep |
| 71 | BG | Löwgren 2013 | DiVA text: Dynabook, experiential qualities, criticism, "generative for design at times", "horizontal and vertical grounding" | Keep |
| 72 | BG | Höök and Löwgren | Crossref abstract; p. 946 "discursiveness and elaboration" | Keep |
| 73 | BG | Dalsgaard and Dindler | OpenAlex abstract; p. 945 | Keep |
| 74 | BG | Zimmerman and Forlizzi 2014 | Book PDF: Drift Table p. 175 and p. 186; Prayer Companion p. 187 (list continues over the page break); p. 178; p. 184 | Reword (page): the exemplar list runs pp. 186 to 187. See S12 |
| 75 | BG | Gaver 2014 | Chapter pp. 147, 161, 162 | Keep |
| 76 | BG | RTD conference | Roams page; conference-series page ("a curated exhibition of design research artefacts", "round-table discussions in 'Rooms of Interest'"); crit | Keep |
| 77 | BG | DIS Pictorials | DIS 2026 call | Keep |
| 78 | BG | Pierce 2014 | OpenAlex abstract (three types, articulations, provocations verbatim); Crossref cites Gaver | Keep |
| 79 | BG | Bardzell et al. 2016 | OpenAlex abstract verbatim; Crossref cites Gaver | Keep |
| 80 | BG | Frauenberger 2019 | OpenAlex abstract: phrase and (a) to (d) verbatim | Keep |
| 81 | BG | Beaudouin-Lafon et al. 2021 | OpenAlex abstract: "actionable principles", three examples, summary tables | Keep |
| 82 | BG | Facts: CHI reviewing | as item 67 | Keep |
| 83 | BG | Facts: generative AI | as item 68 | Keep |
| 84 | BG | Facts: lab budgets | CORDIS 226528: 2,439,757 euros; 1 Apr 2009 to 31 Dec 2014; Goldsmiths; "50 100 batch-produced prototypes" (rendering of a range) | Keep |

## 11-Discussion-Bridge.md

| # | Slide | Question or item | Where checked | Verdict |
|---|---|---|---|---|
| 85 | 11.1 | Support 1: "Instead" context | p. 941 "Instead, what this discussion suggests"; p. 937 [35] | Keep |
| 86 | 11.1 | Support 1: "some sympathy" politeness | p. 937 "uneasy", "helping clarify expectations ..."; p. 946 "may provide a way to further our ambitions to produce relevant and rigorous theoretical work"; p. 937 "relevant and rigorous theory" | Keep |
| 87 | 11.1 | Support 1: describing is not standard-setting | p. 942 to 943; p. 939; p. 938 "often to suggest normative standards" | Keep |
| 88 | 11.1 | Support 2: Lakatos names no judge | SEP "Imre Lakatos" 2.2 (excess empirical content, corroborated, degenerating); p. 941 | Keep |
| 89 | 11.1 | Support 2: not a reviewing manual | p. 941 "Greater integration would make assessment of new contributions fairer and more rigorous" (Gaver summarising the reform calls, as the answer says) | Keep. Nice: add that Gaver's 2014 chapter later names design's test, "does it work?" (Ways of Knowing in HCI, p. 147), still not a procedure |
| 90 | 11.1 | Support 2: critics' proposal | ZFE 2007 p. 499 quotation verbatim; p. 940 "the 'right thing'" | Keep |
| 91 | 11.1 | Support 3: one sketch unfair | p. 945 "illustrated essay"; Figure 1: seven items, "never produced", flickr credits; Crossref July 2012 | Keep |
| 92 | 11.1 | Support 3: form can carry a criterion | p. 945; p. 944 "a balance is achieved", "may vary depending on material, purpose and audience" | Keep |
| 93 | 11.1 | Support 3: extensibility answered | p. 945; chapter p. 161 | Keep |
| 94 | 11.2 | Tension 1 | p. 937 to 938 both quotations, "a long-term practitioner" | Keep |
| 95 | 11.2 | Tension 2 | p. 940 "their concern is to sometimes create good ones"; p. 941; p. 939 Popper; SEP | Keep |
| 96 | 11.2 | Tension 3 | p. 942 three quotations; p. 945; p. 943 | Keep |
| 97 | 11.2 | Tension 4 | p. 945; p. 937; ZFE 2007 p. 500 verbatim | Keep |
| 98 | 11.2 | Tension 5 | p. 943 "gave rise to public opinion itself", "a single, unitary world ...", "may share many assumptions" | Keep |
| 99 | 11.2 | Small 1: manifesto cheap shot | p. 939; p. 938 "describe design practice to illustrate their approach" | Keep |
| 100 | 11.2 | Small 2: Rams annotated his own work | p. 944 "accompanied by short explanations. These annotate the portfolio", "extremely compact ..."; Figure 1 third person | Keep |
| 101 | 11.2 | Small 3: Schön uncited | p. 946 refs 25, 28 | Keep |
| 102 | 11.3 | "Gaver concedes nothing" | p. 943; p. 941 | Keep |
| 103 | 11.3 | Support 3 unfair | section 10 | Keep |
| 104 | 11.3 | Instructor: own position | our position, labelled | Keep |
| 105 | 11.3 | Nobody speaks | Allen and Tanner 2002 (PMC118377): "on average, teachers waited only 1.5 s", "waiting 3 to 5 s ... resulted in dramatic changes in student responses" | Keep |
| 106 | 11.3 | Lopsided vote | facilitation advice | Keep |
| 107 | 11.3 | Quiet after pairs | facilitation advice | Keep |
| 108 | 11.3 | Dominant voices | facilitation advice | Keep |
| 109 | 11.3 | Time short | facilitation advice | Keep |
| 110 | BG | Five discussion principles | course brief wording as logged (not re-opened; quoted phrases match brief Part 1.1 weights) | Keep |
| 111 | BG | Mapping and crediting | labelled "we read"; weights 7, 4, 7 | Keep (two items: mapping, crediting) |
| 112 | BG | Critics' criteria; Lakatos vocabulary | ZFE pp. 499 to 500; SEP 2.2 | Keep (two items) |

---

## Blockers (factual errors a classmate could catch)

**B1. 08, slide 08.1, If asked, "Where exactly does Gaver use the word 'scientism'".** The answer says "Twice." The word appears four times: p. 938, p. 941, p. 945, p. 946. The p. 945 use is quoted in this same block two questions later, so the error is visible from inside the file. Replace the whole answer with:

> A: Four times, and he never defines it. On p. 938 he writes "insofar as HCI is prone to 'scientism' in its cultural assumptions," keeping the word in quotation marks. On p. 941 he says that taking disagreements about science seriously "subverts any easy scientism in HCI more generally, because it destabilises the use of 'science' as shorthand for a set of practices and criteria assumed to be desirable." On p. 945 the conclusion speaks of "HCI's tendencies towards scientism," and on p. 946 it is R4's word, quoted in the acknowledgements. The p. 941 sentence is the nearest thing to a definition: scientism treats "science" as shorthand for practices and criteria assumed to be desirable.

## Should-fix, with exact wording

**S1. 07, If asked Q1.** Replace "It calls itself an essay, twice in the opening pages (p. 937, p. 938)." with "It calls itself an essay three times in the opening pages (p. 937, p. 938)."

**S2. 07, If asked Q7.** Replace "Footnote 1 of this essay credits Bowers with many of the ideas (p. 944)." with "Footnote 1 of this essay says many of its ideas about annotated portfolios "have been developed in collaboration with John Bowers" (p. 944)."

**S3. 07, If asked Q8.** Replace from "As others summarise it" to the end of the answer with: "There he writes that "Science is defined by epistemological accountability," having to defend the basis of its claimed knowledge, while design "works with aesthetic accountability," having to show "that one's design works" (p. 147 of the chapter). He cites the 2012 essay in a footnote (p. 145) and concludes that "Design is not a poor cousin of science" (p. 163). Koskinen and Krogh 2015 trace the distinction back to Gaver, Hooker and Dunne 2001." Then update the log row "Gaver 2014 argument" to "https://people.cs.vt.edu/~gangwang/HCI.pdf, chapter pp. 145, 147, 163: OK" and delete the Open item sentence that the chapter was not read.

**S4. 07, Background, "The annotated portfolio follow-ups".** Replace "and, drawing on Feyerabend, argues that design research should develop its own standards of rationality rather than borrow norms from other fields." with "and, drawing on Feyerabend, suggests that design research "might build its own 'limited rationality'" rather than be brought in line with outside norms for rigour (abstract)."

**S5. 07, Background, "The sequel".** Replace "Secondary summaries: it distinguishes science and design by what each must defend. Science answers for the soundness of its methods; for design, in the 2001 wording Koskinen and Krogh quote, "The question is whether a design 'works,' not whether the methods used to devise them are sound."" with "It distinguishes science and design by what each must defend: science must "explain and defend the basis of one's claimed knowledge," design must show "that one's design works" (p. 147). The idea is older; Koskinen and Krogh 2015 quote Gaver, Hooker and Dunne 2001: "The question is whether a design 'works,' not whether the methods used to devise them are sound.""

**S6. 08, slide 08.2, If asked, Koskinen et al.** Replace "It came out months before CHI 2012, so we call it contemporaneous, not missing." with "It came out in September 2011, around the CHI 2012 submission deadline, so we call it contemporaneous, not missing."

**S7. 08, slide 08.2, If asked, "Did Gaver ever answer the reviewing question himself?"** Replace from "contrasts what science and design each have to defend" to the end with: "contrasts what science and design each have to defend. For design, he says, "The basic question here is 'does it work?'", and questions of process "are not grounds in and of themselves for judging a design successful or unsuccessful" (p. 147 of the chapter). That is a reviewing principle of sorts: judge the design, not the method. It still gives a committee no procedure, so present it as his later direction, not a settled answer to reviewing." Note: inside the double-quoted chapter quotation keep the single marks around 'does it work?'.

**S8. 08, slide 08.2, If asked, "Is it fair to say the Rams figure is 'not an RtD portfolio'?"** Replace "commercial products, annotated by someone other than the designer, showing four of Rams's ten principles." with "Rams's products, one of them a prototype television "never produced," with notes that speak of Rams in the third person; the paper does not say who wrote them. The figure shows four of Rams's ten principles." This also removes a contradiction with section 11 (Watch out 11.2: "the paper does not say who wrote them").

**S9. 08, Background item 3 (Pullin).** Replace "drawing on projects with Hussein Chalayan, Dunne and Raby, the Royal National Institute for Deaf People and the Museum of Jurassic Technology" with "set, in the abstract's words, "in the context of collections by Hussein Chalayan, Dunne & Raby, the Royal National Institute for Deaf People, and the Museum of Jurassic Technology"".

**S10. 08, Background item 5 (Sengers 2006 and Gaver 2014).** Replace "it asks what science and design must each defend. The distinction is older: Koskinen and Krogh 2015 quote Gaver, Hooker and Dunne 2001, "Science is epistemologically accountable," "Design is aesthetically accountable."" with "it asks what science and design must each defend: science, "the basis of one's claimed knowledge"; design, "that one's design works" (p. 147). The distinction is older: Koskinen and Krogh 2015 quote Gaver, Hooker and Dunne 2001, "Science is epistemologically accountable," "Design is aesthetically accountable."" Also change the log row "Gaver 2014 contrasts what science and design must each defend" to cite https://people.cs.vt.edu/~gangwang/HCI.pdf p. 147, OK, and drop the Open item saying the chapter was not read.

**S11. 09, If asked, "You blame the setup for the guessing."** Replace "The instructions themselves pointed to health: "The monitor will move to support your health while working" (p. 6)." with "The instructions themselves pointed to health: "The monitor will move to support your health while working" (p. 6), although the authors call this a "neutral" explanation."

**S12. 10, Background, Zimmerman and Forlizzi 2014.** Replace "(pp. 175, 186; from the studio's other publications, pitfall 12)" with "(pp. 175, 186 to 187; from the studio's other publications, pitfall 12)" and in the log row "Drift Table and Prayer Companion listed as exemplars" change "p. 186" to "pp. 186 to 187".

## Comparison of 07 and 10 on Gaver 2014 (as asked)

Read against the chapter (printed pp. 143 to 165). Nothing either file says is false. Section 10 quotes the chapter directly and correctly: accountability definitions and the meaning of "aesthetic" (p. 147), "Annotated portfolios may capture best what we can know through design" (p. 161), "productive indiscipline" and the warning about imposed frameworks (p. 162). Sections 07 and 08 say the chapter was not read and route the idea through Koskinen and Krogh 2015 and Gaver, Hooker and Dunne 2001. The 2001 attribution is Koskinen and Krogh's; the 2014 chapter itself does not cite Hooker or Dunne 2001 (it cites Gaver, Dunne and Pacenti 1999 only). Section 11 relies on section 10 for p. 161, which is correct. S3, S5, S7 and S10 bring 07 and 08 into line with 10.

## Nice-to-have

- 07 Q3: add that Gaver names "Science and Technology Studies (STS)" beside Philosophy of Science and calls them "a useful foil" (p. 939).
- 07 Q9: add "OpenAlex about 1,140" so 07 and 10 report the same index.
- 07 Background, session: reword the Brynjarsdottir line to "critiques persuasive sustainability research from 2009 to 2011, drawing on sociological theory of modernism."
- 09 Q1: "Gaver's studio was in this debate by 2008" to "Gaver's studio's methods were already cited in this debate by 2008"; add that W&K's endnote 1 is Ways of Knowing in HCI (2014), the volume with Gaver's accountability chapter (W&K p. 44).
- 09 Background, Shin: "four hours on two consecutive days" to "four hours on each of two consecutive days (eight in total)" (p. 5).
- 11 Q (support 2, "reviewing manual"): one added sentence, "His 2014 chapter later names design's test, "does it work?" (Ways of Knowing in HCI, p. 147), which is still not a procedure."
- 11 Open items: the item saying the 11.2 Watch out reads "a critic annotating a designer" is stale; the Watch out now reads "an annotator the paper does not name." Delete the item.
- 10, RTD 2019: the Roams page says "Delft"; if challenged, "Delft and Rotterdam" rests on the 2019 page in the log.

## Single most valuable improvement per file

- **07:** Replace the secondary-source answer on Gaver 2014 (Q8 and Background "The sequel") with the chapter's own p. 147 words (S3, S5).
- **08:** Fix the "scientism" count (B1); the p. 941 sentence also gives the presenter the definition he currently says does not exist.
- **09:** Add the authors' own "neutral" framing to the Shin instructions answer (S11), so the critique cannot be turned back on the presenter.
- **10:** Correct the Zimmerman and Forlizzi exemplar pages (S12); everything else in this file checked out exactly.
- **11:** Add the 2014 "does it work?" line to the support 2 answer (nice-to-have above); it is the strongest reply Gaver's side has, and having it ready makes support 2 look fair.
