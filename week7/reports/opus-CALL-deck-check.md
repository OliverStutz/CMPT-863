# C-ALL deck-wide consistency check, sections 00 to 11

Independent Opus reviewer, 2026-10-19. Read: `_agent-briefs.md` Part 1, Part 2 (for transitions and [classmate] requirements) and Brief C-ALL; `00-Overview.md`; `01` to `11`. Disputed quotations settled against the PDF (pdftotext per page, printed page = PDF page + 936). No section file was edited.

## Summary

| Check | Result |
|---|---|
| 1. Transitions name the next topic and match its opening | 16 of 16 hand-overs pass (14 inside the talk, plus 11.1 to 11.2 and 11.2 to 11.3). None fails. |
| 2. Quotations in more than one file | 44 compared (table below). No wording conflict that changes meaning; 4 small variances (one missing "and" inside quotation marks in the 04.2 title, a capital C in 03, a page range in 08, inner quote marks dropped in 11 notes). |
| 3. Contradictions between files | 3 substantive, 4 minor (listed below). |
| 4. Pitfalls | 1 numbered pitfall broken (pitfall 5, hedge dropped in 11.1); 1 style rule 1.5 break (unlabelled judgment "a manifesto" on 11.2). No other pitfall broken. |
| 5. Dashes | `_check-sections.py`: em=0 en=0 in every file. No other dash characters (U+2012, U+2015, U+2212, " -- ") found by grep. |
| 6. Timing | 0 Say scripts over the limit. 3 at zero or near-zero margin (09.1 227/227, 11.1 96/97, 08.1 128/130 with an unfilled [classmate] inside the Say). On-slide text above about 90 words on 6 slides (03.1 110, 11.2 115, 05.1 104, 10.1 101, 11.1 98, 01.2 97). |
| 7. [classmate] slots | All slots the briefs ask for are present (02, 04.2, 05, 06.2, 07, 08.1, 08.2). Extra slots in 03, 04.1, 09, 11. 00-Overview's list omits 09. |
| 8. Run sheet | Times and total (24:15) match. One title mismatch: slide 13. One count error: "29" connections in report R; the report has 31. |
| 9. Thesis on 11.1 | Supported by 04, 06, 08 and 10, but its second support says Gaver refuses "metrics and convergence"; sections 05 and 08 quote him refusing standardisation (p. 942) and calling shared standards "potentially repressive" (p. 943). 11 does not acknowledge the p. 943 sentence. Fix 1. |

### Output of `python3 slides/_check-sections.py`

```
01-Background.md: em=0 en=0 classmate=0
   01.1 1:30  on-slide  93w  say 174w / limit 195   Research through design, and whose studio this is
   01.2 1:30  on-slide  97w  say 168w / limit 195   2007 to 2012: the call for standards
02-Problem.md: em=0 en=0 classmate=1
   02.1 1:15  on-slide  87w  say 154w / limit 162   The problem: standards, scientism, self-policing
03-Argument-Map.md: em=0 en=0 classmate=1
   03.1 1:30  on-slide 110w  say 187w / limit 195   The argument on one slide
04-Theory.md: em=0 en=0 classmate=2
   04.1 1:30  on-slide  95w  say 177w / limit 195   Five genres of design theory, and the Popper test
   04.2 1:30  on-slide  84w  say 190w / limit 195   Lakatos, and what "provisional, contingent, aspira
05-Convergence.md: em=0 en=0 classmate=1
   05.1 1:45  on-slide 104w  say 202w / limit 227   Why the field does not converge: three readings
06-Annotation.md: em=0 en=0 classmate=1
   06.1 1:30  on-slide  80w  say 176w / limit 195   Artefacts as theory nexus, theory as annotation
   06.2 1:30  on-slide  78w  say 153w / limit 195   Annotated portfolios: Rams, and the converse of pa
07-Method.md: em=0 en=0 classmate=2
   07.1 1:30  on-slide  84w  say 181w / limit 195   How the essay argues: method, evidence, reflexivit
08-Strengths-Limitations.md: em=0 en=0 classmate=4
   08.1 1:00  on-slide  83w  say 128w / limit 130   Strengths
   08.2 1:30  on-slide  85w  say 191w / limit 195   Limitations, and what the essay does not do
09-Connections.md: em=0 en=0 classmate=1
   09.1 1:45  on-slide  92w  say 227w / limit 227   Six readings test Gaver
10-What-Followed.md: em=0 en=0 classmate=0
   10.1 1:30  on-slide 101w  say 188w / limit 195   What followed: Gaver's way, the critics' way, or b
11-Discussion-Bridge.md: em=0 en=0 classmate=11
   11.1 0:45  on-slide  98w  say  96w / limit 97   Our reading of the paper
   11.2 1:15  on-slide 115w  say 152w / limit 162   Tensions worth arguing about
   11.3 1:00  on-slide  88w  say 117w / limit 130   Over to you

Sections total: 23:45, 2861 spoken words
```

Note: the script counts the literal token "**[classmate]**" as one word. Every Say that holds a slot will grow when a name and a phrase are filled in; 08.1 (128/130) and 08.2 (191/195) have almost no room for that.

## Findings per file

### 00-Overview.md
- Line 36, run sheet slide 13: title "Connections to the course readings"; the section file's slide title is "Six readings test Gaver" (09, line 5). Mismatch.
- Line 59: "chosen from the 29 in report R". Report R has 31 connections (W1 to W4, K1 to K4, G1 to G4, O1 to O3, M1 to M3, T1 to T4, Z1 to Z3, B1 to B3, S1 to S3); 09 shows 6 and lists 25 alternates, which sums to 31.
- Line 67, [classmate] list: omits 09 (slot in 09 Watch out, line 29). Otherwise matches.
- Line 43: "about 2,900 words" is fine (script: 2,861).
- Line 29: run-sheet title for slide 6 carries the same quotation-mark issue as 04.2 (see 04).
- Line 91 and the 11 file call Figure 1 "Gaver's sketch"; 06 says the paper does not say who wrote the notes. See contradiction C3.
- Run-sheet times: every slide time matches its file; sections sum to 23:45, plus title 0:30, total 24:15. Correct.

### 01-Background.md
- Line 19, Say: "William Gaver runs the Interaction Research Studio." The paper gives only the affiliation (p. 937), and 01's own Watch out (line 27) says CORDIS does not name him as principal investigator. "Runs" is an unverified outside fact. Replace with "works at".
- Line 43 and line 46: R4 quotation consistent with 07, 08, 11 and the PDF.
- On-slide 97 words on 01.2 (near the limit; acceptable, the quote cards could lose "needs" fragment if crowded).
- Transition 01.1 to 01.2 and 01.2 to 02: pass ("the problem in section 02"; 02 opens on the title question and Gaver's target).

### 02-Problem.md
- Line 104: "might lead to a form of self-policing that would be overly restrictive" (p. 937). Matches the PDF (PDF p. 1, the clause is complete on p. 937). 08 gives the same words as "p. 937 to 938". Both defensible since the sentence crosses the page; see table.
- Transition: pass (three moves; 03 opens "Here is the whole essay on one slide").
- Pitfalls 2, 6, 13, 14 respected.

### 03-Argument-Map.md
- Line 171, on-slide: "Consensus can come to look like constraint" (p. 946) with a capital C; the paper reads "From this perspective, consensus can come to look like constraint." Use lower case or the full sentence.
- On-slide 110 words, the second highest in the deck. A cut is offered in the fix list.
- Line 168: "By Popper's test ... By Lakatos's test": consistent with 04 and pitfall 1.
- Transition: pass (04 opens "Gaver's first move starts with an inventory").

### 04-Theory.md
- Line 34 (and 00 line 29): the slide title puts "provisional, contingent, aspirational" in quotation marks, but the paper never prints the triple without "and" (p. 937 and p. 941 have "provisional, contingent, and aspirational"; p. 938 and p. 945 "provisional, contingent and aspirational"). Inside quotation marks this is a misquotation, small but on a title slide.
- Line 55: the narrow phrasing of the "productive" point (criterion named, who applies it is not) matches 08 limitation 1 and 11.1 support 2. Consistent.
- Line 60 to 61: reference slip described the same way as in 05, 06 and 08 (Newton [3] should be [5]; "Reprinted in [3]" should be [5]). Consistent.
- Transition 04.1 to 04.2 and 04.2 to 05: pass.

### 05-Convergence.md
- Line 22, Say: "most of us value user contact". The paper says "some contact with the potential audiences for the things we make" (p. 942); 08 says "contact with the people we design for". Not a contradiction, but the two scripts should use one wording, and 08's is closer to the paper.
- Line 29, Watch out point (d): the best reply from Gaver's side cites only p. 941 ("'methods, approaches, and outcomes'"). 08's Watch out (line 47) gives two stronger replies (p. 942 "calls for standardisation ... are misplaced"; p. 943 shared standards "potentially repressive") and softens the point. 05 keeps it unsoftened ("Our answer: agreement on methods is narrower than a shared paradigm"). See contradiction C2.
- Line 65, open items: "keep the Say at about 225 words"; the script counts 202. Stale.
- On-slide 104 words. The file itself suggests dropping the opening question line first.
- Transition: pass (06 opens "Move three is the constructive one"; the slide title names annotation).

### 06-Annotation.md
- Line 48 (a): "someone other than Rams wrote them; the paper does not say who." 11 (lines 25, 46, 50, 60) and 00 (line 91) call it "Gaver's sketch" or "a critic annotating a designer". See contradiction C3.
- Line 38 and 50: seven products, four of ten principles. Consistent with 08 and 11; pitfall 8 respected.
- Transitions 06.1 to 06.2 and 06.2 to 07: pass.

### 07-Method.md
- Line 16 and line 29: "two are Gaver's own [13, 14]", "partly from his own work". 08 and 11 do not use "own circle" (grep confirms). Consistent.
- Line 30: "manifesto" correctly flagged as our inference. 11.2 does not keep that label (see 11).
- Transition: pass (08 opens "Five strengths").

### 08-Strengths-Limitations.md
- Line 35 (slide item 3) and line 42 (Say): "Gaver reads the call for standards as a call for one paradigm." The file's own Watch out (line 47) says "soften ours": Gaver addresses standards directly on p. 942 and p. 943. The Say states the unsoftened version. See contradiction C2.
- Line 42, Say: "it annotates Rams's products for Braun". The caption reads "Rams' designs for Braun and Vitsoe" (p. 945), and the figure includes the Vitsoe shelving system. Factual slip.
- Line 18, Say: 128/130 words with an unfilled [classmate] in the Say. Filling it will push the slide over.
- Line 110, open items: "08.2 174 words"; the script counts 191. Stale.
- Line 50: the "against our own case" item is present and fair, with hedges (pitfall 5 respected here).
- Transitions 08.1 to 08.2 and 08.2 to 09: pass.

### 09-Connections.md
- Line 12, on-slide: "Same diagnosis (p. 940), split cure (p. 118)". The second page is Greenberg and Buxton's, placed right after a Gaver page; a reader may take it as Gaver's. Tag it.
- Line 18, Say: 227/227 words, no margin.
- Line 28: states that 02's Greenberg and Buxton and Olsen sentences differ from 09's and do not conflict. Confirmed: 02 quotes p. 111 and the Olsen abstract; 09 quotes p. 113, p. 118 and the "Generality" section.
- Transition: pass (10 opens "So did the field go Gaver's way?").

### 10-What-Followed.md
- Line 19, Say: "We argue that is Gaver's "serious form" (p. 945)". A page number in a spoken script; drop "(p. 945)" from the Say or keep it as a slide reference only.
- On-slide 101 words.
- Line 70 cross-reference to 11 (Löwgren 2013) matches 11 line 26.
- Transition: "So here is our reading." Names the next topic; 11.1 opens "One sentence from us." Pass, though thin.

### 11-Discussion-Bridge.md
- Line 13, on-slide, and line 18, Say: "What he refuses is metrics and convergence (p. 937, p. 946)." Gaver never addresses metrics directly; he quotes the critics' "actionable metrics" and does not answer it. What he does refuse, in his own words, is standardisation: "wary of impulses towards convergence and standardisation" (p. 937), "calls for standardisation, formalisation, overarching theory and the like are misplaced" (p. 942, quoted in 08), and "even the assumption that shared standards are necessary, possible or desirable" is among the "potentially repressive acts of ontological politics" (p. 943, quoted in 02, 05, 08). 11's own Watch out (line 21) says he opposes "convergence and standardisation", so the slide and the notes disagree. See contradiction C1.
- Line 13 and line 18: "a practice 'most of us' share" and "most of us share a practice" drop "seem to" and "Perhaps we do not all agree" (p. 942). Pitfall 5 asks to keep both hedges; 05 and 08 keep them.
- Line 21: "he opposes ... 'disciplinary consolidation' (p. 945)". The paper says the essay "explored some of the issues that complicate calls for disciplinary consolidation" (p. 945). "Opposes" is stronger than "complicate".
- Line 23: "'Metrics and convergence' are what he refuses" repeats C1.
- Line 45 and line 50: "A manifesto with one example" and "The essay is a manifesto with one example" are stated as fact. 07 (line 30) and 11's own Watch out (line 59) say this is our inference. Rule 1.5: label judgments.
- Line 46, 50, 60: "in Gaver's sketch"; see C3.
- On-slide 98 words on a 0:45 slide and 115 words on 11.2: the two densest slides in the deck.
- Line 18: Say 96/97 words.
- Transitions 11.1 to 11.2 and 11.2 to 11.3: pass.
- Thesis support check (item 9): support 1 rests on 05 line 30 (p. 941 sentence, verified in context: it follows "Instead," after the point that only one programme dominates a field); support 2 matches 04 line 55 and 08 limitation 1; support 3 matches 06 tests (a) and (c) and 10 (Löwgren). 08's strongest replies (portfolios allow "extensibility and verifiability", p. 945; the community judges fertility over time) are acknowledged in 11 lines 24 and 26. The one undercut not acknowledged is p. 942 and p. 943 on standards (C1).

## Contradictions

Substantive:
- **C1.** 11.1 says Gaver refuses "metrics and convergence"; 02, 05 and 08 quote him on standardisation (p. 942) and shared standards (p. 943); 11's own Watch out says "convergence and standardisation". The thesis needs to name what he refuses in his own words and face p. 943.
- **C2.** Standards versus paradigm: 08.2 slide and Say state it strongly ("reads the call for standards as a call for one paradigm"); 08's Watch out softens it; 05's Watch out keeps it with a weaker Gaver reply. The deck should give one version.
- **C3.** Who wrote Figure 1's notes: 06 says the paper does not say; 11 and 00 say "Gaver's sketch" and "a critic annotating a designer".

Minor:
- **C4.** 08 Say "Rams's products for Braun" versus the caption "Braun and Vitsoe" (p. 945), which 06 renders correctly.
- **C5.** 05 "user contact" versus 08 "contact with the people we design for" (p. 942: "some contact with the potential audiences").
- **C6.** 00 run-sheet title for slide 13 versus 09's slide title.
- **C7.** 00 "29 in report R" versus 31.

Checked and consistent across files: the reference-list slip (04, 05, 06, 08: Newton and the reprints should be [5]; Carroll and Kellogg missing; "c.f. [3]" on p. 943 correct); the Rams figure (seven products, four of ten principles, text names two); the "own circle" point (only 07 addresses it, narrowly); Lakatos's criterion (04, 08, 09, 11 all say: named as "relative productivity", who applies it in design is not); pitfall 1 phrasing (03, 04, 07, 08); "potentially repressive" (02, 05, 08, 11).

## Cross-file quotation table

Every quotation from the Gaver paper used in more than one file, with the wording and page in each. "Same" means identical wording and page. PDF-checked items are marked.

| # | Quotation | Files: wording and page | Verdict |
|---|---|---|---|
| 1 | "provisional, contingent, and aspirational" | 03 slide p. 937; 04 notes p. 941 (also 937, 938); 01 notes, no page; 04 title "provisional, contingent, aspirational" | PDF: p. 937 and 941 have "and" with comma; title drops "and" inside quotation marks. Fix 7 |
| 2 | "depends on the account of science one adopts" | 03, 04, 07, 08 slide, all p. 941 | Same |
| 3 | "extensibility and verifiability" | 01 "a level of extensibility and verifiability" p. 937 ([33]); 02 p. 937; 04 "limit the potential 'extensibility and verifiability'" p. 941; 06, 09 "'extensibility and verifiability'" p. 945; 08, 11 "'extensibility and verifiability' for which some of the HCI design community have called" p. 945; 11 notes line 26 without inner marks | Three real occurrences on three pages; all correct. 11 line 26 should keep the inner single marks |
| 4 | "actionable metrics" | 01 "some actionable metrics for bringing rigor in critique" p. 937; 02, 08, 11 p. 937 | Same (01 Say shortened, logged) |
| 5 | "lacks clear expectations and standards" | 01 p. 937; 02 "clear expectations and standards" p. 937 | Same |
| 6 | "protocols, descriptions, and guidelines" | 01, 02, 11, all p. 937 | Same |
| 7 | "like any other research approach" | 01, 02, p. 937 | Same |
| 8 | "self-policing ... overly restrictive" | 02 slide "might lead to a form of self-policing that would be overly restrictive" p. 937; 08 "a form of self-policing that would be overly restrictive" p. 937 to 938; 11 "self-policing" p. 937 | PDF: words all on p. 937, sentence ends on p. 938. Acceptable; prefer p. 937 for the fragment |
| 9 | "prone to 'scientism'" | 02, 08, 11, p. 938 | Same |
| 10 | R4 "an elephant in the room...: the move towards 'scientism' in HCI" | 01 full, p. 946; 07, 08 slide "an elephant in the room" p. 946; 11 log full, p. 946 | Same |
| 11 | "I have some sympathy with these concerns" | 02 "some sympathy"; 07, 08, 11 full; all p. 937 | Same |
| 12 | "a route towards disciplinary legitimacy" | 01, 02, 08, 11, p. 937 | Same |
| 13 | "potentially repressive acts of ontological politics" | 02, 05, 08 p. 943; 11 Q5 cites p. 943 for many worlds | PDF-checked. Same |
| 14 | "political acts to be approached with care" | 03, 05, 08, p. 945 | Same |
| 15 | "consensus can come to look like constraint" | 03 slide "Consensus can ..." p. 946; 08, 11 p. 946 | Capital C in 03 only |
| 16 | "theories that are sometimes right" | 04 p. 941 (log: sentence starts p. 940); 11 p. 941 | Same |
| 17 | "productive research programmes" | 03 log p. 940 to 941; 04, 08, 09, 11 p. 941 | Same (03 log range harmless) |
| 18 | "relative productivity of competing programmes" | 04, 11, p. 941 | PDF-checked. Same |
| 19 | "testimony to the fertility of the overarching theory" | 04, 09, p. 941 | Same |
| 20 | "predict novel facts" | 04, 08, 09, p. 940 | PDF-checked. Same |
| 21 | "in principle a falsifiable statement" | 04, 07, p. 940 | Same |
| 22 | "comparable ambiguous and unambiguous systems" | 04, 07, p. 940 | Same |
| 23 | "must founder" | 05 p. 942; 08 log p. 941 to 943 (section span) and p. 942 | Same |
| 24 | "seemingly behind recent calls for formalism and standardisation" | 05, 08, p. 942 | Same |
| 25 | "most of us" (shared-values list) | 05, 08, 11, p. 942 | Same page; 11 drops "seem to" (pitfall 5) |
| 26 | "Perhaps we do not all agree with all of these assumptions all of the time" | 05, 08, p. 942 | Same |
| 27 | "the desirability of clearly articulating ... the standards one might want to advocate for them" | 05, 11, p. 941 | PDF-checked in context. Same |
| 28 | "no problem to solve" / "proliferation rather than agreement" | 03, 05, p. 938 | Same |
| 29 | "a serious form of theoretical contribution" | 03, 06, 08, 10, p. 945 | Same |
| 30 | "neither dominates, and neither is subservient" | 06, 08, 11, p. 945 | Same |
| 31 | "not defined by their graphic presentation" | 06, 08, 11, p. 945 (caption) | Same |
| 32 | "The notion of annotated portfolios is not a formal one" | 06, 11, p. 944 | Same |
| 33 | "I am not proposing here that annotated portfolios subsume all other forms of design theory" | 06, 07, 08, p. 945 | Same |
| 34 | "a contribution to meta-theory" | 04, 07, 11, p. 939 | PDF-checked. Same |
| 35 | "can be considered as wide-ranging manifestos" | 07, 11, p. 939 | PDF-checked. Same |
| 36 | "It is difficult to call for research to be more scientific if what it means to be scientific is under question" | 08, 11, p. 941 | Same |
| 37 | "wary of impulses towards convergence and standardisation" | 09, 10, 11, p. 937 | Same |
| 38 | "establishes an area in that space" | 06 "an area in that space"; 09 "establishes an area"; both p. 944 | PDF-checked. Same |
| 39 | "ultimate particular(s)" | 01 p. 939, 944; 04 p. 939; 06 p. 944 | Same (both pages correct) |
| 40 | "a collection of designs by the same or associated designers" | 09 log, 11, p. 944 | PDF-checked. Same |
| 41 | "both because of the provisional nature of theory, and the definite nature of designed artefacts" | 03, 06, p. 943 | Same |
| 42 | "Sketch design" (caption) | 08, 11, p. 945 | PDF-checked. Same |
| 43 | the iPad sentence | 05 p. 943 (dashed clause replaced by "..."); 11 cites p. 943 | Same |
| 44 | "Annotated portfolios may capture best what we can know through design" (Gaver 2014, outside source) | 10, 11, Ways of Knowing in HCI p. 161 | Same |

Course-reading quotations: Greenberg and Buxton and Olsen appear in 02 and 09 with different sentences (G&B p. 111 versus p. 113 and p. 118; Olsen abstract versus "Generality"). No overlap, no conflict.

## Fix list, in priority order

1. **11.1 thesis support (C1, item 9).**
   - Slide, line 13: replace "What he refuses is metrics and convergence (p. 937, p. 946)." with "What he refuses is standardisation and convergence (p. 937, p. 942)."
   - Same bullet: replace "and a practice "most of us" share (p. 942)" with "and a practice "most of us ... seem to share" (p. 942)".
   - Say, line 18: replace "on page 942, that most of us share a practice. What he refuses is metrics and convergence." with "on page 942, that most of us seem to share a practice. What he refuses is standardisation and convergence." This adds 2 words (98/97), so also cut "Our discussion starts there." from the end of the Say (94/97).
   - Watch out, line 23: replace the bullet with: "Standardisation and convergence are what he refuses, in his words: "wary of impulses towards convergence and standardisation" (p. 937) and "calls for standardisation, formalisation, overarching theory and the like are misplaced" (p. 942). "Actionable metrics" is the critics' phrase (p. 937); the essay never answers it, which is our third point."
   - Add a Watch out bullet: "A classmate may quote p. 943 against us: "even the assumption that shared standards are necessary, possible or desirable" is among the "potentially repressive acts of ontological politics." Concede it. Page 941 accepts articulated standards and page 943 doubts shared ones; that gap is what "concedes more than his conclusion admits" names, and the conclusion settles on "political acts to be approached with care" (p. 945)."
   - Watch out, line 21: replace "and "disciplinary consolidation" (p. 945)" with "and says his essay explores issues that "complicate calls for disciplinary consolidation" (p. 945)".

2. **08.2 limitation 3 and 05 point (d) (C2).**
   - 08 slide, line 35: replace "Reads calls for standards as calls for convergence" with "Mostly reads calls for standards as calls for convergence".
   - 08 Say, line 42: replace "Three: Gaver reads the call for standards as a call for one paradigm," with "Three: Gaver mostly reads the call for standards as a call for one paradigm," (Say 192/195).
   - 05 Watch out, line 29: after "on his framing the convergence question does address them." insert "He also answers standards directly, once in each of the other two readings: "calls for standardisation, formalisation, overarching theory and the like are misplaced" (p. 942) and shared standards may be "potentially repressive" (p. 943); section 08 softens its version of this point the same way." and replace "Our answer: agreement on methods is narrower than a shared paradigm." with "Our narrow answer: his strongest reading, Kuhn's, answers the paradigm question; the standards question gets a sentence each."

3. **Figure 1 authorship (C3).**
   - 11.2 slide, line 46: replace "Rams's principles on Rams's products, in Gaver's sketch (p. 944 to 945)." with "Rams's principles on Rams's products, in the essay's sketch (p. 944 to 945)."
   - 11.2 Say, line 50: replace "in Gaver's sketch" with "in the essay's sketch".
   - 11.1 Watch out, line 25: replace "Our point: the essay shows a critic annotating a designer with the designer's own principles" with "Our point: the essay shows an annotator it does not name annotating a designer with the designer's own principles".
   - 11.2 Watch out, line 60: replace "the figure is Gaver's "Sketch design" and its notes speak of Rams in the third person" with "the figure is a "Sketch design" whose notes speak of Rams in the third person; the paper does not say who wrote them".
   - 00 line 91: replace "so they are Gaver's sketch, not Rams's words" with "so they are the sketch's words, not Rams's".

4. **11.2 unlabelled judgment (rule 1.5).**
   - Slide, line 45: replace "A manifesto with one example (p. 939, p. 945)." with "We argue: by its own taxonomy, a manifesto with one example (p. 939, p. 945)."
   - Say, line 50: replace "The essay is a manifesto with one example." with "By its own taxonomy, we would call the essay a manifesto with one example." (Say 158/162.)

5. **08 Braun (C4).** Say, line 42: replace "it annotates Rams's products for Braun" with "it annotates Rams's products for Braun and Vitsoe".

6. **00-Overview (C6, C7, slot list).**
   - Line 36: replace "Connections to the course readings" with "Six readings test Gaver".
   - Line 59: replace "chosen from the 29 in report R" with "chosen from the 31 in report R".
   - Line 67: replace "08 (both slides), 11 (11.2 and 11.3)" with "08 (both slides), 09 (Watch out), 11 (11.2 and 11.3)".

7. **04.2 title quotation.** In 04 line 34 and 00 line 29, replace `what "provisional, contingent, aspirational" buys` with `what "provisional, contingent, and aspirational" buys`.

8. **Timing margins with [classmate] fills.**
   - 08.1 Say, line 18: cut "These strengths set the bar for the limitations." (8 words) so the [classmate] credit has room (120/130 before the fill).
   - 09.1 Say, line 18: cut "Six of them change how we judge him." (8 words) to leave a margin (219/227).

9. **On-slide density (about 90 words).**
   - 11.2: move the "Three smaller things" block (heading and three bullets, about 25 words) to the Watch out; the Say already speaks them. Brings the slide to about 90.
   - 03.1, line 168: replace "By Popper's test, design theory is unfalsifiable: theory underspecifies design, and design is generative. By Lakatos's test, no flaw if the theory is fertile." with "By Popper's test, unfalsifiable: theory underspecifies design; design is generative. By Lakatos's test, no flaw if fertile." (saves 7 words); line 171, replace ""Consensus can come to look like constraint" (p. 946)" with ""consensus can come to look like constraint" (p. 946)".
   - 05.1: drop the opening line ""Why is there not more convergence in the field?" (p. 941)" as the file's own open item suggests (saves 10 words).
   - 10.1, line 12: replace "Zimmerman and Forlizzi 2014, "formalize this research approach" (p. 178)" with "Zimmerman and Forlizzi 2014" (saves 5 words; the quote stays in the log).

10. **01.1 unverified role.** Say, line 19: replace "William Gaver runs the Interaction Research Studio at Goldsmiths." with "William Gaver works at the Interaction Research Studio at Goldsmiths."

11. **05 and 08 wording of the shared values (C5).** 05 Say, line 22: replace "value user contact" with "value contact with the people we design for" (Say 207/227).

12. **Small items.**
   - 09 slide, line 12: replace "split cure (p. 118)" with "split cure (G&B p. 118)".
   - 10 Say, line 19: replace "Gaver's "serious form" (p. 945) surviving" with "Gaver's "serious form" surviving".
   - 11 Watch out, line 26: replace "allow "extensibility and verifiability" (p. 945)" with "allow "the 'extensibility and verifiability'" (p. 945)".
   - Stale counts: 05 line 65, replace "about 225 words" with "about 202 words"; 08 line 110, replace "08.2 174 words" with "08.2 191 words".
