# C-08 independent check of 08-Strengths-Limitations.md

Checked 2026-10-19 against the PDF itself (`week7/Week 7 - What Should We Expect from Research through Design_.pdf`, pdftotext per page, plus the Figure 1 render `assets/p9-09.png`), not against the quotation bank. Outside references checked against Crossref, the publisher pages and a fresh web search. The section file was not edited.

Summary: 31 paper quotations or attributions checked, 31 confirmed word for word with the right page. 6 outside references checked (Schön, Pullin, Cross, Koskinen et al., Höök and Löwgren, CHI 2012 timing), 6 confirmed. No blockers. 7 should-fix items, 7 nice-to-have items. No em-dashes or en-dashes in the file (grep counts 0 and 0).

## 1. Quotations, pages and attributions

| Claim in file | Where checked | Verdict |
|---|---|---|
| "prone to 'scientism'" | p. 938, PDF 2 ("insofar as HCI is prone to 'scientism' in its cultural assumptions") | Confirmed |
| R4: "an elephant in the room" | p. 946, PDF 10, acknowledgements | Confirmed. But "in part" is Gaver's framing ("who recognised that this essay addresses, in part, ..."), not R4's phrase. See should-fix 2 |
| "I have some sympathy with these concerns" | p. 937, PDF 1 | Confirmed |
| standards "may be a route towards disciplinary legitimacy within HCI", "helping clarify expectations" | p. 937 | Confirmed |
| "self-policing that would be overly restrictive" | p. 937; the sentence runs on to p. 938 at "value for its ability" | Confirmed ("p. 937 to 938" is acceptable; "overly restrictive" itself sits on p. 937) |
| "depends on the account of science one adopts" | p. 941, PDF 5 | Quote confirmed. The on-slide subject "science" is wrong; Gaver's subject is "whether research through design can be considered a science". See should-fix 1 |
| "tend[s] to think" design is "distinct from the ideals, and often practice, of science" | p. 941 (dash after "science" correctly not reproduced) | Confirmed |
| "It is difficult to call for research to be more scientific if what it means to be scientific is under question (see [29])" | p. 941 | Confirmed. The on-slide splice "when its meaning 'is under question'" is legal but loose; see nice-to-have 1 |
| [29] Sengers 2006, "Must design become 'scientific'?", DIS'06 workshop | p. 946 reference list | Confirmed |
| "a serious form of theoretical contribution" | p. 945 (crosses the column break, text intact) | Confirmed |
| "neither dominates, and neither is subservient" | p. 945 | Confirmed |
| "I am not proposing here that annotated portfolios subsume all other forms of design theory" | p. 945 | Confirmed |
| "productive research programmes" | p. 941 | Confirmed |
| "their ability to inspire new designs" | p. 941 | Confirmed |
| "predict novel facts" | p. 940, Lakatos [19] p. 24 quoted by Gaver | Confirmed. Watch out 1 should say "Lakatos, quoted by Gaver"; see nice-to-have 2 |
| "some actionable metrics for bringing rigor in critique of design research", [9] p824, CHI 2011 panel | p. 937 | Confirmed; attribution to the critics (panel [9]) is correct |
| reviewing, funding, examining absent from the argument | my own per-page regex over PDF 1 to 10 | Confirmed. Only incidental hits: "I review two influential theories" (p. 937), "students" as makers (p. 938), "examination of Rams' philosophy" (p. 945), reviewers in acknowledgements (p. 946). The log missed the p. 937 and p. 945 hits; harmless |
| convergence section spans p. 941 to 943 | PDF 5 to 7 | Confirmed |
| "calls for standardisation, formalisation, overarching theory and the like are misplaced" | p. 942 | Confirmed |
| "most of us working in research through design seem to share a set of common values"; user contact, wide exploration, craft, making as discovery | p. 942 | Confirmed. Gaver's wording is "some contact with the potential audiences for the things we make", so "contact with users" in the Say is a fair paraphrase (nice-to-have 3) |
| "Perhaps we do not all agree with all of these assumptions all of the time"; "seems to be the outlines of a broad consensus" | p. 942 | Confirmed |
| "potentially repressive acts of ontological politics" | p. 943 | Confirmed; the "potentially" hedge is kept (pitfall 6 respected) |
| "political acts to be approached with care" | p. 945, conclusion | Confirmed |
| "reflect on the appropriate ways to pursue our research on its own terms" | p. 945 | Confirmed |
| Figure 1 "Sketch design", caption "(note that annotated portfolios are not defined by their graphic presentation)" | p. 945 and the image | Confirmed. The figure annotates Rams's commercial products for Braun and Vitsoe, notes in the third person, four Rams principles in the centre |
| Footnote 1, ideas developed with John Bowers | p. 944, PDF 8 | Confirmed in the PDF (the log says "bank"; it is right) |
| "consensus can come to look like constraint" | p. 946 | Confirmed |
| References 25 (Pullin 2010) and 28 (Schoen 1983) listed, never cited | p. 946 list; no "[25]" or "[28]" in PDF 1 to 9 | Confirmed |
| Reference slip: [3] is Collins 1994, used for Newton and the moon (p. 939) and "Reprinted in [3]" (p. 946); Carroll and Kellogg [3] (p. 944) missing from list | p. 939, 944, 946 | Confirmed. The file rightly leaves out p. 943 "([17]; c.f. [3])": there Collins 1994, "Why the Social Sciences Won't Become High-Consensus, Rapid-Discovery Science", is the correct source. So pitfall 10 overstates the slip; the file does not repeat that error |
| CHI 2012, 5 to 10 May 2012 | page header, every page | Confirmed |

## 2. Outside references

| Claim | Where checked | Verdict |
|---|---|---|
| Schön, The Reflective Practitioner: How Professionals Think in Action, Basic Books, New York, 1983 | WorldCat URL in the file returned 403 to my fetch; a fresh search shows the same WorldCat record (OCLC 8709452) and archive.org (https://archive.org/details/reflectivepracti0000scho): Basic Books, New York, 1983 | Confirmed |
| Pullin, "Statement of Practice: Curating and creating design collections, from Social Mobiles to the Museum of Lost Interactions and Six Speaking Chairs", Design and Culture 2(3), 309 to 328, Nov 2010 | Dundee URL in the file, and Crossref DOI 10.2752/175470710X12789399279877 | Confirmed. Gaver's list says 2(1); the file's note on this is correct |
| Cross, "Designerly ways of knowing", Design Studies 3(4), 221 to 227, Oct 1982 | Crossref DOI 10.1016/0142-694X(82)90040-0 | Confirmed. The content claim (design as its own culture, distinct from science and humanities) rests on Wikipedia; it is the standard reading of the paper, but see nice-to-have 5 |
| Koskinen, Zimmerman, Binder, Redström, Wensveen, Design Research Through Practice: From the Lab, Field, and Showroom, 1st ed., Sept 2011 | Elsevier shop URL in the file | Confirmed (authors, title, edition, September 2011). The page does not name the imprint; Morgan Kaufmann is right but not shown on that page |
| Höök and Löwgren, "Strong concepts", ACM TOCHI 19(3), Oct 2012 | Crossref DOI 10.1145/2362364.2362371 | Confirmed. The file is right to treat it as a parallel answer, not an omission |
| Koskinen "may have come too late" | Fresh search: CHI 2012 papers were due in September 2011 (secondary listing, https://hci.rwth-aachen.de/conferences-old, gives 5 Sept 2011) | Plausible and fair as phrased ("may"). It also means the slide should not list Koskinen as "Missing"; see should-fix 4 |

## 3. Strengths (slide 08.1)

| Strength | Fair to the essay? | Rating |
|---|---|---|
| 1. Names scientism (p. 938, 946) | Yes. The file correctly drops the brief's "where the field had not" because Gaver cites Sengers [29]. | Keep |
| 2. Keeps hedges (p. 937, 941) | Yes, but the on-slide fragment makes "science" the subject. | Reword: `Hedges: "I have some sympathy with these concerns" (p. 937); whether RtD is a science "depends on the account of science one adopts" (p. 941)` |
| 3. Constructive: annotated portfolios (p. 945) | Yes; Watch out pairs it with "not proposing ... subsume" (pitfall 9). | Keep |
| 4. Disagreement as evidence (p. 941) | Yes; it is Gaver's own move. | Keep (see nice-to-have 1 for the splice) |
| 5. Balance, "neither dominates" (p. 945) | Yes. "Protects both sides, the maker and the theorist" is labelled as our reading by "we think" at the start. | Keep |

## 4. Limitations (slide 08.2) and Gaver's best reply

| Limitation | Fair? | Is the Watch out reply the strongest available? | Rating |
|---|---|---|---|
| 1. "productive research programmes" has no named judge (p. 941) | Yes, as "we argue". | Good: Lakatos's novel facts (p. 940) and "inspire new designs" (p. 941). Could add the conclusion: theories "may nonetheless inspire thriving research programmes" (p. 945), which shows Gaver treats the test as historical, judged after the fact. | Keep |
| 2. Silent on reviewing and funding; critics asked for "actionable metrics" (p. 937) | Yes; the words are absent and the file states it as fact, not as a claim that Gaver denies criteria. | Not the strongest. Gaver's best reply is on p. 945: annotated portfolios respect the ultimate particular "while allowing for the 'extensibility and verifiability' for which some of the HCI design community have called." That is his direct answer to the critics' demand, and a classmate who has read p. 945 will raise it. Add it first in reply 2. | Keep slide line; fix reply (should-fix 5) |
| 3. Answers "one paradigm" when critics asked for "standards" (p. 941 to 943) | Only partly. The file's own reply concedes Gaver addresses standards on p. 942 and 943. The defensible version is that Gaver himself reads the critics' calls as a wish for convergence: the pre-paradigmatic diagnosis is "a tempting conclusion, one seemingly behind recent calls for formalism and standardisation [9, 35]" (p. 942). | Reply is fair and well sourced. | Reword: on slide `3. Reads calls for standards as calls for convergence: "seemingly behind recent calls" (p. 942)`; in the Say, replace the two sentences with: "Three: Gaver reads the critics' call for standards as a call for one paradigm, 'seemingly behind recent calls for formalism and standardisation.' Standards are a smaller request than a paradigm." (net plus about 9 words; 08.2 stays near 182 of 195) |
| 4. Sketched, not shown: one figure, one designer, "his own principles" (p. 945) | Yes, but "his own" reads as Gaver's on the slide, and it misses the sharper point visible in the figure: the one worked example annotates a commercial product designer (Rams for Braun and Vitsoe), not a research through design portfolio. | Good reply (caption, Bowers footnote, section 10). Add: Gaver says a fuller treatment "might well take the form of an illustrated essay" (p. 945), so he flags it as a sketch himself. | Reword: on slide `4. Sketched, not shown: one figure, of Rams's commercial products, not an RtD portfolio (p. 945)`; in the Say: "Four: the proposal is sketched. One figure, and it annotates Rams's products for Braun, not a research through design project." |
| 5. Missing voices: Schön, Pullin listed and uncited; Cross 1982; Koskinen et al. 2011 | Mixed. Schön and Pullin are fair. Cross supports Gaver, so its absence costs lineage only (the reply says so). Koskinen appeared in Sept 2011, around the CHI 2012 deadline, so listing it as "Missing" on the slide contradicts the file's own reply, which says it "may have come too late". The sharpest item is Pullin: an article on curating design collections, listed and never engaged, in an essay that proposes curated, annotated collections. | Reply good. | Reword: on slide `5. Missing: Pullin on curating design collections and Schön, both listed, never cited (p. 946); Cross 1982 absent`; in the Say: "Five: missing voices. Pullin's article on curating design collections, the closest earlier work to portfolios, and Schön sit in the reference list but are never cited, and Cross is absent." Move Koskinen to the Watch out as "contemporaneous" |
| (a) Many worlds exempts everything (notes only) | Fair as framed; the file correctly calls it the weakest and quotes "potentially" and "approached with care". | Strong. | Keep in notes |
| (b) Who benefits (notes only) | Fair, labelled as our reading, no motive attributed. | The constructed reply ("restriction falls hardest on newcomers") is not in the text; label it "a reply in Gaver's spirit". | Keep in notes (nice-to-have 4) |
| Reference slip (notes only) | Fair and light, correctly kept off the slide. | n/a | Keep |

## 5. Against our own case

Present on the slide (boxed row) and in the Say. Fair: it keeps "most of us" and "Perhaps we do not all agree" (pitfall 5). It can be made stronger and fairer still, because Gaver draws this conclusion himself in the next sentence: "we already share many of the attributes of a research paradigm" (p. 942 to 943). Citing that shows the reading is Gaver's, not a gotcha (nice-to-have 6).

## 6. Say scripts, words per minute

| Slide | Time | Limit (130 per minute) | Count without [classmate] | Count with the token | Verdict |
|---|---|---|---|---|---|
| 08.1 | 1:00 | 130 | 127 | 128 | Pass, but tight: any reword must be word-neutral |
| 08.2 | 1:30 | 195 | 173 | 174 | Pass; room for about 20 words, enough for the rewordings above |

On-slide text: about 81 words (08.1) and 87 words (08.2), under 90.

## Blockers

None. Every quotation is exact and on the stated page; no pitfall is broken in a way a classmate could catch as an error of fact.

## Should-fix

1. 08.1, Hedges bullet: change `science "depends on the account of science one adopts"` to `whether RtD is a science "depends on the account of science one adopts"` (p. 941). As it stands it puts words in Gaver's mouth about science itself.
2. 08.1, Watch out 1: "R4's phrase also says the essay addresses scientism 'in part'" misattributes. Reword: "Gaver himself says the essay addresses R4's elephant only 'in part' (p. 946)."
3. 08.2 item 4: replace "his own principles" (reads as Gaver's) with the wording in section 4: one figure, of Rams's commercial products, not a research through design portfolio.
4. 08.2 item 5: drop "Koskinen et al. 2011" from the slide and the Say, because the file's own Watch out says it may have postdated submission; lead with Pullin as the closest earlier work. Wording in section 4.
5. Watch out 2: add Gaver's strongest reply first: annotated portfolios allow "the 'extensibility and verifiability' for which some of the HCI design community have called" (p. 945). Then keep our narrow claim: he says portfolios allow it, but not how a reviewer would judge it.
6. 08.2 item 3: anchor to p. 942 "seemingly behind recent calls for formalism and standardisation [9, 35]" so the limitation rests on Gaver's own sentence rather than on a framing the reply then retracts. Wording in section 4.
7. Sources line and log: add p. 942 "seemingly behind recent calls" and p. 945 "extensibility and verifiability" once 5 and 6 are made; add the archive.org Schön URL, since the WorldCat URL returns 403 to automated fetches.

## Nice-to-have

1. 08.1 item 4: quote the whole sentence to avoid the splice: "It is difficult to call for research to be more scientific if what it means to be scientific is under question" (p. 941). About 4 more on-slide words.
2. Watch out 1: "programmes that, in Lakatos's words quoted by Gaver, 'predict novel facts' (p. 940)".
3. Say 08.2: "contact with the people we design for" is closer to "potential audiences" than "contact with users".
4. Who-benefits reply: mark it "a reply in Gaver's spirit, not in the text".
5. Cross content: replace the Wikipedia source with Cross 1982 itself or Cross 2001 "Designerly ways of knowing: design discipline versus design science" (Design Issues 17(3)); verify before adding.
6. Against our own case: add Gaver's own conclusion, "we already share many of the attributes of a research paradigm" (p. 942 to 943).
7. Reference-slip aside: note that p. 943 "c.f. [3]" correctly points to Collins (social sciences), so the slip is limited to p. 939, p. 944 and the reprints; tell the orchestrator that pitfall 10 overstates it.

## Single most valuable improvement

Reword limitation 4 to say what Figure 1 actually shows: the only worked example of an annotated portfolio annotates Dieter Rams's commercial products for Braun and Vitsoe, not a research through design project (p. 945). This is checkable by anyone looking at the figure, it is fair (Gaver calls it a "sketch design" and points to an "illustrated essay" as the fuller form), and it ties limitations 2 and 4 together: the essay offers portfolios as a form reviewers could judge, but never shows one as a research contribution, so the reader cannot see how a committee would judge it.
