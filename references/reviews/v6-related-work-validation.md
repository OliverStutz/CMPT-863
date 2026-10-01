# Validation of the v6 Related work against the full texts

Checked on 2026-09-29. Scope: "## Appendix: Related work" of `CMPT 863 Project Abstract v6.md`, the Aamir sentence in "## Research question", the "## References" list, and the related-work statements about "our" study against the v6 Method. Every verdict rests on the full text in `references/` (pdftotext of the PDFs; Strickland tables read from the page images because the OCR drops them; Aamir from the saved ACM page and its .md; Reif and von Schenk from the Europe PMC Markdown). The older reports in `reviews/` were not used as evidence. Crossref was queried only for reference-list metadata that the full texts lack (Prike accepted manuscript, Strickland scan).

Page numbers: Strickland = journal pages; Prike = accepted-manuscript pages (differ from print); Chen, Greitemeyer, Mickelberg = PDF pages; Aamir, Reif, von Schenk = section names (no pagination in the files).

## Summary

| Paper | Claims checked | Correct | Imprecise | Wrong | Not found |
|---|---|---|---|---|---|
| Strickland (1958) | 11 | 11 | 0 | 0 | 0 |
| Prike et al. (2025) | 17 | 16 | 1 | 0 | 0 |
| Chen et al. (2026) | 14 | 12 | 2 | 0 | 0 |
| Greitemeyer & Sagioglou (2015) | 10 | 8 | 2 | 0 | 0 |
| Mickelberg et al. (2025) | 3 | 3 | 0 | 0 | 0 |
| Aamir et al. (2026), incl. Research question | 4 | 4 | 0 | 0 | 0 |
| von Schenk et al. (2024) | 9 | 9 | 0 | 0 | 0 |
| Reif et al. (2025) | 4 | 4 | 0 | 0 | 0 |
| **Total** | **72** | **67** | **5** | **0** | **0** |

Internal consistency with the Method: 18 statements checked, 13 consistent, 5 inconsistent (2 serious).
Reference list: 8 entries, 7 match; 1 incomplete (Aamir lacks the article number); 2 could not be fully verified from the full text (Prike manuscript, Strickland scan) but Crossref confirms them.

No claim about any paper is wrong. The serious problems are internal: the related work promises two things the Method does not deliver (an end-of-study rating of how true each flag was; a trust rating after the unflagged slides 1 and 5).

---

## 1. Strickland (1958)

| # | Claim | Verdict | Evidence | Note |
|---|---|---|---|---|
| S1 | 40 participants | CORRECT | "Ss were 40 male students at the University of North Carolina" (p. 202) | |
| S2 | male students | CORRECT | same | |
| S3 | played a supervisor | CORRECT | "an experiment on supervision" (p. 202) | |
| S4 | two fictitious workers | CORRECT | "Actually there were no real subordinates" (p. 203) | The supervisors believed the workers were real. |
| S5 | told to check one after nine of ten periods, the other after two | CORRECT | "he would monitor A after work periods 1, 2, 3, 4, 6, 7, 8, 9, and 10, he would monitor B after work periods 5 and 10" (p. 203) | |
| S6 | both produced the same output | CORRECT | "the output totals of his two subordinates were equal, despite occasional fluctuations" (p. 204; Table 1: 52/52, 50/50) | Equal totals, not equal every period. |
| S7 | every check found the work in order | CORRECT | All monitored periods show 5 or more (Table 1, p. 204); Part II: "every production report slip ... showed that the subordinates had completed five examples" (p. 205) | Caveat: in Group II (half the sample) the supervisor later saw that each worker fell below the standard once in an *unchecked* period ("A on trial 5 ... and B on trial 3", p. 204). The checks themselves always found the work in order. |
| S8 | afterwards asked which worker they trusted more | CORRECT | Trust item at "the termination of Part I"; S "could state that he held greater trust in B than A, greater trust in A than B, or equal trust" (p. 206) | "Equal" was an allowed answer. |
| S9 | more named the rarely checked worker | CORRECT | Table 3 (p. 206): B > A 10 + 14 = 24; A > B 1 + 0 = 1; equal 9 + 6 = 15 | Optional precision: "24 of 40 named the rarely checked worker, one the often checked worker, 15 trusted both equally." |
| S10 | the workers never spoke | CORRECT (by design) | Communication was one-way: S read instructions over the sound system, turned on the "no-talking" light (p. 205); the only information from workers was the production slips. Not stated in one sentence, but follows from the procedure. | |
| S11 | the supervisor held power over the person checked | CORRECT | "he could fine 10c of a possible $1.00 wage of a subordinate" (p. 203); hypothesis premise "high power vis-a-vis a subordinate" (p. 202) | |

## 2. Prike, Baker & Ecker (2025)

| # | Claim | Verdict | Evidence | Note |
|---|---|---|---|---|
| P1 | N = 406 | CORRECT | "Final sample size was thus N = 406" (ms. p. 9) | |
| P2 | U.S. adults | IMPRECISE | "U.S.-based Prolific users ... who met the platform's prescreening criterion of identifying as politically moderate"; 40 with strong left or right orientation removed (ms. p. 9) | The screening for political moderates is central to the paper (title: "noncommitted voters"). Suggested: "406 politically moderate U.S. adults". |
| P3 | two unfamiliar, fictional candidates | CORRECT | "two fictional centrist political candidates ... set in Malta ... likely to be unfamiliar to U.S. participants" (ms. pp. 5–6); participants who recognised a candidate excluded (p. 9) | |
| P4 | through a profile and a series of posts | CORRECT | "Politician Profiles" (ms. p. 11); statements "presented as social-media posts" (p. 12) | |
| P5 | nonpartisan fact-checking organisation | CORRECT | "a fictional fact-checking organization ('Fact Checks Malta'), which was introduced as nonpartisan and objective" (ms. p. 8) | |
| P6 | marked as false eight statements | CORRECT | "eight false statements, all associated with one of the candidates" (ms. p. 8); Table 2 fact-checks begin "FALSE." (p. 13) | The fact-checker also confirmed four true statements, two per candidate (ms. p. 8), a useful parallel to "No issues found". |
| P7 | that favoured one candidate | CORRECT | "designed to enhance that candidate's perceived competence and morality or reduce the perceived competence and morality of their opponent" (ms. p. 8) | |
| P8 | rated less competent and less moral than his opponent, compared with control | CORRECT | Fact-check contrasts vs control: competence d = 0.90 and 1.14 (ms. p. 18), morality d = 1.11 and 1.31 (p. 19); "fact-checks reduced perceptions of competence to substantially below control levels" (p. 18) | Analysed as the gap between the two candidates, relative to control. |
| P9 | the control group saw only neutral posts | CORRECT | Control saw only the 16 "generic statements" (ms. p. 8; Table 1, p. 10) | The generic posts positioned the candidates as slightly left or right; "neutral" = no checkable claims. |
| P10 | fewer voted for him | CORRECT | "4.85 times more likely to vote for the non-misinforming candidate ... and 5.42 times" vs control (ms. p. 23) | |
| P11 | Prike used a profile | CORRECT | ms. p. 11 | |
| P12 | 0-to-10 rating | CORRECT | "11-point sliding scale ... '0, Very incompetent' to '10, Very competent'" (ms. p. 12) | Prike rated competence, morality and warmth, not trust. |
| P13 | end-of-study rating of how true each checked claim is | CORRECT (about Prike) | "Belief in the true and false statements was measured ... ('0, Entirely untruthful' to '10, Entirely truthful')" (ms. p. 15), after the vote (p. 16) | See INCONSISTENT I5: v6 has no such measure. |
| P14 | their checker was human | CORRECT (implied) | fact-checks shown "as comments made by a fact-checking group" (ms. p. 7) | Not stated as "human", but nothing suggests automation. |
| P15 | and right | CORRECT | Within the study's fiction, the eight statements were built as false and the fact-checks gave the correct facts (Table 2, p. 13). | |
| P16 | the candidate stayed silent | CORRECT (by design) | No candidate reply to a fact-check in any condition (Table 1, p. 10) | |
| P17 | ratings taken once, at the end | CORRECT | "They then completed the candidate-impression, voting behavior, and statement-belief measures, in that order" after the simulation (ms. p. 16) | |

## 3. Chen, Sun, Liao & Sundar (2026)

| # | Claim | Verdict | Evidence | Note |
|---|---|---|---|---|
| C1 | N = 135 | CORRECT | "the final sample consisted of 135 participants" (p. 6) | 260 recruited, 125 excluded. |
| C2 | short online therapy chat | CORRECT | "brief, free therapy session" (p. 5); "15-min study" (p. 6) | |
| C3 | at the end of the chat | CORRECT | "At the end of the consultation, the doctor offered to consult an AI assistant" (abstract, p. 1) | |
| C4 | a chatbot presented as a human doctor | CORRECT | "we prompted an LLM-powered chatbot to act as a human doctor" (p. 2); debrief revealed it (p. 5) | |
| C5 | said only that his AI assistant agreed or disagreed | CORRECT | Prompt step 7: "mention that your AI medical system CareBot [agrees/disagrees] with you" (p. 14) | Participants first had to opt in to hear it (p. 5). Wording varied ("CareBot disagreed with me, suggesting a different approach", p. 13) but no reasons were given. |
| C6 | disagreement left participants less sure about their own situation | CORRECT | "AI disagreement increased perceptions of medical uncertainty" (abstract); uncertainty defined as lacking information "to fully understand the true nature of their medical situation" (p. 3) | Caveat: "the negative effect of AI disagreement on perceived medical uncertainty only occurred when patients perceived the LLM-simulated doctor as more humanlike" (p. 11), about half the sample (p. 13). |
| C7 | made the doctor seem lazier | CORRECT | "AI disagreement increased perceived doctor laziness" (Table 2, p. 12) | Same humanlike caveat ("stronger among participants who perceived ... more human-like"). |
| C8 | both went with lower trust in him | CORRECT | Indirect effects on cognitive, affective and behavioural trust (pp. 8–9) | "Went with" fits: mediation, trust measured once. |
| C9 | "Because the doctor never answered" / "the silence" | IMPRECISE (wording) | "the doctor not providing any explanations after the AI's disagreement" (p. 11); the doctor then went on to suggest a full session (p. 5) | Nobody challenged the doctor, so there was nothing to answer; he was not silent, he gave no explanation. Suggested: "Because the doctor gave no explanation, the study cannot separate the disagreement from the missing explanation; the authors suspect the missing explanation ..." |
| C10 | the study cannot separate disagreement from the missing explanation | CORRECT (design inference) | Every disagreement came without an explanation; no condition varied it (prompt, p. 14) | |
| C11 | the authors suspect the silence | CORRECT | "This finding may be due to the doctor not providing any explanations after the AI's disagreement, which could be seen as a lack of effort" (5.2, p. 11) | Said for the laziness path only. |
| C12 | they suggest an explanation, not tested | CORRECT | "could potentially be mitigated by providing explanations for the assistant's disagreement" (5.1, p. 11); example scripts (5.5, p. 13); no such condition was run | |
| C13 | the doctor brought in the AI himself | CORRECT | "the LLM-simulated doctor then asked participants if they would like to hear the opinion of his AI assistant, CareBot" (p. 5) | |
| C14 | "Chen et al. note that an AI's contrary opinion shifts the burden of proof onto the expert" | IMPRECISE (attribution) | "Given that a second opinion from an AI assistant can shift the burden of proof onto doctors (Kempt & Nagel, 2022)" (5.1, p. 11); "disputing or rejecting AI's recommendations places the burden of proof on doctors (Kempt et al., 2023)" (2.8, p. 4) | The passage exists, but it is Chen citing Kempt and Nagel, hedged ("can"), about doctors, and not tested by Chen. Suggested: "Citing Kempt and Nagel (2022), Chen et al. note that a second opinion from an AI can shift the burden of proof onto the doctor." |

Not checked: "Chen et al. (2026) is the closest design" is the authors' judgement.

## 4. Greitemeyer & Sagioglou (2015)

Means on the −5 to +5 scale (Study 1 pp. 5–6; Study 2 pp. 8–9):

| | Before clearing | After clearing |
|---|---|---|
| S1 exoneration (n = 68) | −0.16 | +1.85 |
| S1 uncorrected accusation (46) | −0.30 | −0.28 |
| S1 never accused, no article (71) | +1.54 | +1.63 |
| S2 exoneration (94) | +0.38 | +2.47 |
| S2 uncorrected accusation (78) | +0.15 | +0.27 |
| S2 never accused, same article (89) | +2.35 | +2.27 |
| S2 never accused, no article (97) | +0.97 | +0.96 |

| # | Claim | Verdict | Evidence | Note |
|---|---|---|---|---|
| G1 | online participants | CORRECT | "took part on Amazon Mechanical Turk" (pp. 3, 7) | |
| G2 | rated a researcher accused of misconduct | CORRECT | "accused of reporting inconclusive data and that the ethics committee of his university demanded the retraction" (p. 4) | Ratings: trustworthy researcher, belief in his findings, further funding. |
| G3 | and again after reading he was cleared | CORRECT | "Sanna and his colleagues were wrongly accused. None of the allegations were confirmed" then the same three questions (p. 4) | |
| G4 | N = 185 and 358 | CORRECT | "One-hundred-and-eighty-five" (p. 3); "Three-hundred-and-fifty-eight" (p. 7) | |
| G5 | accused researcher rated about two points below one never accused | IMPRECISE | Gap 1.70–1.84 in Study 1; 1.97–2.20 against Study 2's same-article control; only 0.59–0.82 against Study 2's no-article control | "About two points" holds only for the first two comparisons, and the next sentence of the abstract brings in the no-article control. Suggested: "rated about two points lower (1.7 to 2.2) than one never accused who was described with the same article". |
| G6 | scale from −5 to +5 | CORRECT | "All items were assessed on a scale from −5 to +5" (p. 4) | |
| G7 | after the clearing he rose by about two points | CORRECT | +2.01 (Study 1), +2.09 (Study 2) | |
| G8 | to that level but not reliably above it | CORRECT | Study 1: exoneration vs control "not significant, p = .536" (p. 6); Study 2: exoneration vs article control "not significant, t(354) = 0.66, p = .508" (p. 9) | |
| G9 | clearly above a never-accused researcher only when that one was described with less information | IMPRECISE (minor) | Study 2: exoneration +2.47 vs no-article control +0.96 (p. 9). Study 1: the control was also described without the article, yet +1.85 vs +1.63 was not reliable (p. 6) | True as stated ("only when"), but less information was not enough by itself: it happened in Study 2 only. Suggested: "He ended clearly above a never-accused researcher only in Study 2, where that researcher was described without the article." |
| G10 | a third party did the clearing | CORRECT | The clearing came from the study text, not from the researcher (p. 4, p. 7) | |

## 5. Mickelberg et al. (2025)

| # | Claim | Verdict | Evidence | Note |
|---|---|---|---|---|
| M1 | a retraction removed the effect of an allegation on how likeable a person seemed in two of three experiments | CORRECT | Exp 1 retraction vs neutral p = .161 (p. 7); Exp 3 p = .516 (p. 12); Exp 2 residual effect p = .003 (p. 10). Abstract: "in the other two experiments the retracted misinformation was fully discounted" | Measure: Reysen Likability Scale. The authors add that equivalence tests were inconclusive ("no conclusive evidence for or against", pp. 9, 14). |
| M2 | the authors note that rating after every step may draw attention to the latest information | CORRECT | "It is possible the dynamic task obscured the continued influence effect by focusing participants' attention on the more recent information [i.e., the retraction]" (p. 2; again p. 15) | Their own Experiments 1 and 3 (rating only at the end) also found full recovery, so they conclude "The nature of the task itself can therefore not satisfactorily explain the observed results" (p. 15). "May" in the abstract is appropriately hedged. |
| M3 | a third party did the clearing | CORRECT | Updates attributed to John's acquaintances; "RETRACTION: John did NOT have an affair ..." (pp. 5–6) | |

## 6. Aamir et al. (2026), Related work and Research question

| # | Claim | Verdict | Evidence | Note |
|---|---|---|---|---|
| A1 | (RW, RQ) directors felt/pushed back more confidently challenging management where independent reviews were routine | CORRECT | "Where external assurance was routine, directors reported greater confidence in challenging management perspectives" (4.1); "It gave directors confidence to challenge assumptions" (6.2) | "Independent reviews" = the paper's "external assurance ... internal audit, red-teaming, or third-party reviews" (6.2). Self-report from 13 interviewees, "directors and C-suite" (section 7). |
| A2 | (RQ) directors lacking cyber literacy and confidence struggle to question cyber-risk reports | CORRECT | "a substantial gap in directors' knowledge and confidence ... limits their ability to interrogate risk reports, challenge assumptions" (Abstract) | |
| A3 | (RQ) often accept management's account | CORRECT | "Where this literacy was missing, boards often defaulted to accepting management narratives" (6.2) | |
| A4 | (RQ) the same finding as A1, in past tense | CORRECT | as A1 | Style only: the RQ sentence mixes "struggle ... accept" with "pushed back". |

## 7. von Schenk et al. (2024)

| # | Claim | Verdict | Evidence | Note |
|---|---|---|---|---|
| V1 | 2,040 people | CORRECT | "We recruited 2,040 participants (='judges')" (Judgment study) | |
| V2 | each judged whether a stranger's written statement was a lie | CORRECT | "each of the participants who acted as a judge read one statement and decided whether this statement was true or false" (Judgment study) | Statements about weekend plans, written by other Prolific participants. |
| V3 | knowing an accusation would cost the writer money | CORRECT | "the author of the statement would only earn a bonus with certainty when not being accused of lying" (Judgment study) | |
| V4 | without help, about one in five accused | CORRECT | "In the Baseline treatment ... the accusation rate was 19.22%" (Lie accusations across treatments) | |
| V5 | those who chose to consult an algorithm | CORRECT | Choice treatment: "around 31.37% of participants made that choice", at £0.05 | |
| V6 | right about two times in three | CORRECT | "The algorithm reached 66.86% accuracy" (Lie-detection algorithm) | |
| V7 | mostly followed it | CORRECT | 88.13% of requesters followed the prediction | |
| V8 | when it said "lie", almost 85% accused | CORRECT | "84.91% when the prediction is 'lie'" | |
| V9 | the accused could not answer | CORRECT (by design) | Statements were collected in January 2022 and judged in May 2023 (Procedural details); no reply stage exists | Not stated in words; follows from the design. |

## 8. Reif, Larrick & Soll (2025)

| # | Claim | Verdict | Evidence | Note |
|---|---|---|---|---|
| R1 | employees described as getting help from generative AI | CORRECT | Study 2: a lawyer "sometimes asks generative AI to summarize information" (Study 2) | One of four studies (Study 2, 1,215 recruited). |
| R2 | rated lazier | CORRECT | laziness M = 2.50 vs 2.16, P < 0.001 (Study 2) | |
| R3 | and less competent | CORRECT | "rated as less competent, less diligent, less independent, and less self-assured" (Study 2) | |
| R4 | than employees who got the same help from another source | CORRECT | "the help provided in the AI help and Non-AI help conditions was described identically and only the source varied" (Study 2) | Supports logging own AI use: in Study 4 the laziness penalty depended on the rater's own AI use (Study 4). |

---

## 9. Internal consistency: Related work vs the v6 Method

| # | Related-work statement | Method (v6) | Verdict |
|---|---|---|---|
| I1 | "the checker also shows 'No issues found' on slides 1 and 5" (Strickland para.) | Step 2: "'No issues found' on slides 1 and 5" | CONSISTENT |
| I2 | "... so that flagged and unflagged slides can be compared" (Strickland para.) | Step 3: "After each of slides 2–4: trust rating." No rating after slide 1 or slide 5; step 4 is "Budget decision and background questions". Validity also says "Slide 1, with no issue reported, shows whether the checker's presence alone moves trust." | **INCONSISTENT**. With no rating after slides 1 and 5, trust after an unflagged slide is never measured, so the comparison cannot be made. Fix the Method (add a trust rating after slides 1 and 5), or drop the comparison from the related work and Validity. |
| I3 | "We borrow the profile" (Prike para.) | Step 1: "Profile establishing the CISO's experience." | CONSISTENT |
| I4 | "the 0-to-10 rating" | Measures: "Trust in the CISO, 0 to 10" | CONSISTENT |
| I5 | "and the end-of-study rating of how true each checked claim is" | Not in the Method. Step 4 and Measures list budget decision, education, age, field, cyber knowledge, own AI use; no rating of the flags | **INCONSISTENT**. Either add to step 4 / Measures ("At the end, participants rate how true each flag was, 0 to 10"), or cut the phrase: "We borrow the profile and the 0-to-10 rating." |
| I6 | "ours is an AI tool, wrong for this company, the CISO answers" | "Every flag is wrong for this company, and each full answer explains why" | CONSISTENT |
| I7 | "trust is rated after every step" (Prike para.) | Seven ratings: at the start, then after each of slides 2–4 twice (after the flag, after the answer). None after slides 1 and 5 or after the budget decision. Measures says "at each step above", which is equally loose | **INCONSISTENT** (overstated). Suggested: "trust is rated at the start and twice at each flagged slide, after the flag and after the answer" (or "after every slide and every answer" if ratings after slides 1 and 5 are added). |
| I8 | "rate trust after the flag and after the answer" (Chen para.) | Step 3: trust rating, then answer, then "Second trust rating" | CONSISTENT |
| I9 | "H2 tests an answer with the company's reason, H3 a denial" | H2, H3 | CONSISTENT |
| I10 | "our neutral group hears the same answers" (Greitemeyer para.) | "The CISO's answer is the same whoever asks" | CONSISTENT (see observation 3 below) |
| I11 | "both our groups rate at the same steps" (Mickelberg) | Step 3 ratings apply to both groups | CONSISTENT |
| I12 | "our CISO answers for himself" | Step 3 | CONSISTENT |
| I13 | "the same answer should move trust further than in the neutral group (H4b)" | H4b | CONSISTENT |
| I14 | "so the CISO never mentions AI" (Reif) | Not stated anywhere in the Method | **INCONSISTENT** (missing). Add to the Method, e.g. in step 3: "The CISO's answers never mention AI." |
| I15 | "we log which question participants choose" | Measures: "The action chosen at each flag" (say nothing, ask to explain, challenge with the finding) | **INCONSISTENT** (wording, minor). "Saying nothing" is not a question. Suggested: "we log which action participants choose at each flag". |
| I16 | "and how often they use AI themselves" | Measures: "own AI use" | CONSISTENT (the Method could say "how often") |
| I17 | "our checker claims faults that are not there" | "Every flag is wrong for this company"; debrief: "no slide had a real issue" | CONSISTENT (the RQ's "a flag can be right by the textbook" is compatible) |
| I18 | "our board member holds power over the person checked" | RQ: "the board holds the power and the budget" | CONSISTENT |

### Observations on the Method that bear on these statements (not counted above)

1. H1 says trust after a flagged slide is "lower than after the slide before". For slide 2 the slide before is slide 1, which has no rating (same gap as I2).
2. The research question asks about change "from the start to the end of the briefing", and Validity uses "the change from the first to the last rating", but the last rating comes before slide 5.
3. The neutral group hears "the same" answers, yet each full answer "explains why" a flag is wrong. The answer text has to make sense to someone who saw no flag; worth stating how.
4. H4a compares checker participants with "neutral participants [who] say they have questions", but the Measures list only "the action chosen at each flag"; the neutral group's yes/no answer should be listed as a measure.

---

## 10. Reference list

| Entry | Authors | Year | Title | Venue, volume(issue), pages / article | DOI | Verdict |
|---|---|---|---|---|---|---|
| Aamir et al. | Tooba Aamir, Georgia Psaroulis, Marthie Grobler, Helge Janicke: match | 2026 (published 13 April 2026): match | match | CHI '26 Proceedings, ACM: match; ACM page gives "Article No.: 1599, Pages 1–22"; the reference has pp. 1–22 but **no article number** | 10.1145/3772318.3791142: match | INCOMPLETE. Use "(Article 1599, pp. 1–22)". |
| Chen et al. | match | 2026: match | match | *Int. J. Human–Computer Studies* 213, 103824: match | 10.1016/j.ijhcs.2026.103824: match | OK |
| Greitemeyer & Sagioglou | match | 2015: match | match | *PLoS ONE* 10(5): e0126316: match | 10.1371/journal.pone.0126316: match | OK |
| Mickelberg et al. | Amy J. Mickelberg, Bradley Walker, Ullrich K. H. Ecker, Piers D. L. Howe, Andrew Perfors, Nicolas Fay: match | 2025: match | match | *PLoS One* 20(5): e0322045: match | 10.1371/journal.pone.0322045: match | OK (the 2025 PDF styles the journal "PLOS One"; "PLOS ONE" is acceptable) |
| Prike et al. | Toby Prike, Joseph Baker, Ullrich K. H. Ecker: match | not in manuscript | match | Manuscript says only "In press at Political Psychology"; volume, issue, pages not in the file | not in the file | CANNOT VERIFY FROM FULL TEXT. Crossref confirms *Political Psychology* 46(5), 1081–1100, DOI 10.1111/pops.13059 (online November 2024, issue 2025; citing 2025 is fine). |
| Reif et al. | Jessica A. Reif, Richard P. Larrick, Jack B. Soll: match | 2025: match | match | *PNAS* 122(19), e2426766122: match | 10.1073/pnas.2426766122: match | OK (Crossref agrees) |
| Strickland | Lloyd H. Strickland: match | not on scan | "Surveillance and trust": match | Scan shows pages 200–215 (match); journal, volume, issue not printed on the scanned pages | not on scan | PARTLY VERIFIABLE. Crossref confirms *Journal of Personality* 26(2), 200–215, DOI 10.1111/j.1467-6494.1958.tb01580.x. |
| von Schenk et al. | Alicia von Schenk, Victor Klockmann, Jean-François Bonnefon, Iyad Rahwan, Nils Köbis: match | 2024: match | match | *iScience* 27(7), 110201: match | 10.1016/j.isci.2024.110201: match | OK (Crossref agrees) |

All eight papers cited in the related work appear in the list, and the list has no extra entries.
