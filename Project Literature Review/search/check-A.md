# Check A of "Literature Review - Draft v1.md" (2026-10-19)

Checker A, working alone from `_check-brief.md`. check-B.md and check-C.md were not read. The draft and all other files were left unchanged.

## A. Verdict

**Minor fixes.** None of the 52 cited claims is wrong or unsupported. Six are slightly overstated (Bendeck's 13 percent, Du's red highlighting, Opdenbusch's "too shy", Aamir's "pushed back", Kim et al.'s "monitored", and Sharma's "most studies they review"), and every one can be fixed in the same number of words. The draft uses all four corrected wordings, meets every Instructions.md requirement at 1,472 body words, and keeps every author rule. What remains is wording, plus a few places where a CHI reader meets undefined terms (H2, H3, VLM, METI, CISO) or a source whose relevance is not stated (Schilke, Pezzo).

## B. Claim table

V = verified, W = wrong, U = unsupported, O = overstated. "Full text" means the PDF or text in `references/` (or `../references/` for the Abstract v8 sources), read with pdftotext.

| # | Draft sentence (short) | Source checked | Class | Correction |
|---|---|---|---|---|
| 1 | SEC "chose not to require disclosure of board cyber expertise", quote "directors with broad-based skills in risk management and strategy often effectively oversee management's efforts without specific subject matter expertise" | SEC final rule full text, p. 85 ("we are not adopting proposed Item 407(j)") | V | Quote is exact. |
| 2 | Lowry: non-expert oversight "mostly symbolic" even when doing what experts did | Lowry full text, abstract p. 783 | V | |
| 3 | Lowry: one expert director said for most boards "all they can do is assess the credibility of the executives that are put in front of them" | Lowry full text, p. 797 to 798 (D-9 expert; also the epigraph, "Former NASDAQ Board Director") | V | |
| 4 | Proudfoot: competence view "heavily determined by the CISO's ability to effectively communicate" | Proudfoot full text, p. 240 | V | |
| 5 | "Directors are 'too shy to ask questions since they do not want to get embarrassed for their missing knowledge'" (Opdenbusch) | Opdenbusch full text, section IV-B4: "Some participants (CISO4; EX6; NED3,4) pointed out that NEDs lack ... Hence, they would be too shy ..." | O | Four interviewees said it; the draft states it as a general fact about directors. Write "Four interviewees said directors are 'too shy ...'". |
| 6 | "most interviewees named the budget as the board's only cyber decision" (Opdenbusch) | Opdenbusch full text, p. 5: "Most participants (NED1,3,4; EX1,3-6; CISO1-3) explained that simply setting the right level of budget ... would be the only (abstracted) decision" | V | 11 of 18. |
| 7 | "Directors pushed back more confidently where independent reviews were routine" (Aamir) | Aamir full text, Findings: "Where external assurance was routine, directors reported greater confidence in challenging management perspectives"; "It gave directors confidence to challenge assumptions" | O | Self-reported confidence, not observed pushing back. Use the Abstract v8 appendix wording: "felt more confident challenging management". |
| 8 | Opdenbusch: "the CISOs are the ones in power", because they decide what reaches the board | Opdenbusch full text, contribution III (p. 2) and Discussion ("CISOs ... can decide what they want to report to the board") | V | |
| 9 | Lowry "trace the deference to low expertise instead" | Lowry full text, p. 785: "CISOs are not considered to be powerful relative to board members ... hence, we find that low expertise is a novel reason board members rely more heavily on and defer to management" | V | So the stated disagreement is real and fairly put. |
| 10 | Schilke: party with better alternatives trusted its partner less | Schilke full text, pp. 12950 to 12953 (four studies) | V | |
| 11 | Schilke: from the same starting level, trust fell further during a negotiation | Schilke full text, p. 12951 (both 2.59 before; 2.12 against 2.46 after; drop 0.48 against 0.15) | V | |
| 12 | "Their power was having alternatives, not authority" | Schilke full text, p. 12950 (power as structural dependence, alternatives) | V | The consequence for our study is not said; see readability and fix 8. |
| 13 | "The board evidence is from interviews, with no tool and no trust measure" | Lowry, Opdenbusch, Proudfoot, Aamir full texts | V | Lowry adds a CISO survey and a proxy count; "mostly interviews" would be exact, but the point stands. |
| 14 | Ray-Ban Meta glasses answer questions about what the wearer is looking at (Meta, 2024) | results-5, camera assistants entry (Meta newsroom, 23 April 2024) | V | |
| 15 | Apple Visual Intelligence lets users ask for "an explanation of a diagram from class notes" | results-5, same entry | V | Strictly, users ask ChatGPT through Visual Intelligence; acceptable. |
| 16 | Copilot in Teams suggests the private prompt "Where are some holes in [a meeting participant's] argument?" | results-5, Microsoft entry (private pane) | V | |
| 17 | Danry: audio glasses right on 63.5 percent of statements running live | Danry full text, p. 8, section 5.1 | V | |
| 18 | Squash: 14 relevant and 20 irrelevant fact-checks; lab began building an interface for a human editor to approve each pop-up | results-5, Adair entry; fact-check-2026-10-19 item 3 | V | Corrected wording used. |
| 19 | GPT-4V at about the 16th percentile of humans on a visualization literacy test | Bendeck full text, p. 3 | V | |
| 20 | GPT-4V "showed some knowledge of design best practice" | Bendeck full text, abstract and p. 5 | V | |
| 21 | "answered only 13 percent of questions that needed a value read off a chart when it lacked the data" | Bendeck full text, p. 4 and Table 5: 13 percent is the "non-visual lookup" subset; all lookup (single value retrieval) questions were 23 percent without data (93 percent with) | O | The draft gives the lower subset figure for the wider category. Write "answered only 23 percent of questions asking for a single value when it lacked the data". |
| 22 | "None of this work reports what a check does to trust in the speaker; Danry et al. rated recorded statements, not a person" | Danry full text, p. 7 (pre-recorded arguments from a loudspeaker; ratings of agreement, reasonableness, donation) | V | |
| 23 | Horne and Nevo: wrong AI warning on a true headline cut trust by 0.30 on a scale from -1 to +1 | Horne full text, Table 2 (-0.298), label "Warning: An AI Tool Says This Story is False" | V | |
| 24 | Lu: AI credibility indicator had a larger effect on accuracy than an expert's vote | Lu full text, abstract and section 6 (exploratory regressions) | V | |
| 25 | Lu quote "lack the capability to differentiate correct and incorrect AI predictions" | Lu full text, p. 9 (the abstract-level variant reads "the correctness of AI predictions"; the draft quotes the body text exactly) | V | |
| 26 | Du: 214 teachers graded an identical paper 1.8 points lower on a 10-point scale after a fictitious high AI-detection rate | Du full text, abstract and section 4.3 (6.83 against 5.02 without highlighting) | V | 1.8 is computed from the means; the paper does not print it. Fine. |
| 27 | "red highlighting made the report seem more credible and the drop larger" | Du full text, sections 3.4.5, 4.3 and 4.4: red had a main effect on the 1-to-10 rating (6.50 and 4.26), but the warning by highlighting interaction for that rating was not significant (p = .074); the "report perception" scale measures perceived informational value and influence on judgement | O | "Drop larger" implies the interaction the paper did not find for the 1-to-10 rating, and "credible" is a loose name for the scale. Write "red highlighting gave the report more weight and lowered the grade further". |
| 28 | von Schenk: people who chose to consult a lie-detection algorithm accused in almost 85 percent of cases when it said "lie" | `../references/vonSchenk-2024-lie-detection-algorithms-accusation.md` (84.91 percent, Choice treatment, requesters); Abstract v8 | V | |
| 29 | Strickland: more supervisors trusted a rarely checked worker than an often checked one | Strickland full text (`../references/`), Table 3 and Hypothesis 2; Abstract v8 | V | |
| 30 | Prike: a human fact-check lowered an unfamiliar candidate's ratings of competence and morality | Prike full text (`../references/`), abstract and results | V | Their fact-checks were correct (the candidate did spread misinformation); the draft does not say otherwise, but see readability. |
| 31 | Madanay: 1,606 adults; agreed less with the undercalling radiologist, rated lowest of the four conditions; overcalling one paid no penalty | Madanay full text, Table 2; fact-check item 2 | V | Corrected wording ("lowest of the four conditions") used. |
| 32 | Chen: AI disagreement made a doctor seem lazier and lowered trust in him | Chen full text (`../references/`), abstract | V | The effect on trust runs through laziness and uncertainty; the doctor was an LLM chatbot presented as human. Acceptable at this length. |
| 33 | Reif: employees getting help from AI rated lazier and less competent than those helped elsewhere | `../references/Reif-2025-social-evaluation-penalty-AI.md`, Study 2 | V | |
| 34 | D. Kim et al. (2026): advisor's error lowered trust, how the AI was involved changed nothing, so the loss is tied to the error | Kim 2026 full text, pp. 1520 to 1521; fact-check item 4 | V | Corrected reading used. |
| 35 | Loconte: people resisted confident AI accusations of lying when accuracy and confidence were on screen | Loconte published full text; verified-view entry 17 | V | |
| 36 | Mayer: trust as willingness to be vulnerable, based on ability, benevolence, integrity; trust leads to risk taking | Mayer full text, pp. 712, 720, 726 | V | |
| 37 | METI built on the same factors for laypeople judging an expert (Hendriks) | Hendriks full text (cites Mayer et al., ref. 31) | V | |
| 38 | P. H. Kim et al. (2004): after a hearsay competence accusation, apology repaired trust better than a denial, and that denial came with an account | Kim 2004 full text; verified-view entry 2 | V | |
| 39 | Ferrin: a reply that neither confirmed nor denied did no better than the weaker of the other two | Ferrin full text; verified-view entry 10 | V | |
| 40 | Greitemeyer and Sagioglou: clearing by a third party restored the researcher to the never-accused level, not reliably above | Greitemeyer full text (`../references/`), Study 1 discussion | V | |
| 41 | Mickelberg: rating after every step may draw attention to the latest information | Mickelberg full text (`../references/`), p. 3 (recency effect with repeated ratings, citing Stewart) | V | |
| 42 | Sharma: return to the start is complete repair; a rise beyond it is "further development of trust" | Sharma full text, "What Are the Limits of Trust Repair?" | V | Corrected term used ("over-repair" gone). |
| 43 | "Most studies they review find that repair falls short of the start" | Sharma full text, introduction: "most of the studies have shown that repair tactics do not return the trust levels to the pre-transgression levels (Kim, 2018)" | O | It is a statement they cite from Kim (2018), not a finding of their own review. Write "Most repair studies, they note, find that trust falls short of the start". |
| 44 | Sharma: denial works less well after competence than after integrity violations | Sharma full text, "Denial" | V | |
| 45 | Kempt and Nagel: when an AI contradicts the expert in charge, the burden of proof moves to the expert, who cannot settle it by reasons; a second human should settle it | Kempt and Nagel full text, pp. 222, 224, 226, 227; fact-check item 1 | V | Corrected attribution used. |
| 46 | "a point Chen et al. (2026) take up" | Chen full text, p. 11; fact-check item 1 | V | |
| 47 | Jongsma and Sand reply that agreement carries the same burden | read-danry-kempt-dasilva.md (accepted manuscript read in full) | V | |
| 48 | Pezzo: physician who had not consulted an aid before a bad outcome rated higher when he believed he knew what it would say than when he did not trust it, never below no-aid | read-pezzo-schilke-sharma.md (abstract only) | V | Matches the cautious abstract-level clause; no numbers, not called "the only study". |
| 49 | D. Kim et al. (2026) closest design: "an AI monitored a human advisor in front of the student", METI, N = 157 | Kim 2026 full text, pp. 1515 to 1517 (three modalities: advisor only, advisor-invoked AI, automatic monitoring; N = 157; METI items adapted from Hendriks) | O | Only one of three modalities monitored; in another the advisor invoked the AI and in the third there was no AI (fact-check note C6). Write "an AI checked a human advisor's work". |
| 50 | Kim 2026: AI always right, advisor never replied, student held neither the tool nor any power | Kim 2026 full text; verified-view entry 1 | V | |
| 51 | Gupta: nearest real-time wearable checker, measured belief and trust in the AI, never rated the speaker | Gupta full text (arXiv), abstract, measures table, section 6.3 | V | |
| 52 | Sharma: only 8 of 94 empirical repair studies used any within-subject design; a repair claim needs three ratings | Sharma full text, "Measurement Issues and Recommendations" | V | |

**Totals:** 52 claims; verified 46, wrong 0, unsupported 0, overstated 6.

**The four corrected wordings (fact-check-2026-10-19.md):** burden of proof given to Kempt and Nagel, with Chen taking it up (claim 45 to 46): used. Madanay "lowest of the four conditions", not "worst provider" (31): used. Squash "began building an interface" for a human editor, 14 relevant and 20 irrelevant (18): used. Kim 2026 as a bound, not as support for an AI penalty (34): used. None of the old wordings appears in the draft.

## C. Compliance (items 2 to 5)

| Item | Requirement | Status | Evidence |
|---|---|---|---|
| 2 | Under 1,500 body words | Met (narrowly) | My count without title, headings, references and writer's notes: **1,472** (opening 47, boards 278, checkers 226, wrong flags 360, trust and repair 375, gap 186). This matches the writer's count. Only 28 words of headroom, so fixes must stay word-neutral overall (see D). |
| 2 | At least 15 sources | Met | 34 references, all cited: 27 scholarly, 5 products or industry (Meta, Apple, Microsoft, Adair, plus the SEC rule as regulation), Bendeck as a benchmark paper. |
| 2 | Thematic subsections | Met | Five themed sections plus an opening. |
| 2 | Synthesis, not one-by-one summary | Mostly met | Sections 1 and 4 (Kempt, Jongsma, own hypothesis) synthesise and state a disagreement. Section 3 paragraph 1 and section 4 paragraph 2 still run one study per sentence (see F). |
| 2 | Each theme: what exists, what we know, what is unresolved, how we differ | Met, with one weak spot | 1: regulation and practice / deference and fear / interviews only / a check in the member's hands. 2: products / often wrong / no trust in the speaker / VLM flags wrong for the company. 3: stated / stated / "the expert stays silent and trust is rated once" / implicit only (the CISO answers is never said in this section). 4: stated / stated / "the accuser is human" / own hypothesis. |
| 2 | Gap and contribution clear by the end | Met for the gap; contribution partly met | The gap paragraph lists what the study combines, but does not say in one sentence what would be learned (whether a correct expert's reason can win back what a wrong machine flag costs). |
| 2 | APA as in Abstract v8 | Met | The 15 entries shared with Abstract v8 are identical character for character. In-text forms are APA 7; "D. Kim" and "P. H. Kim" initials are correct APA 7 for same-surname first authors. Two small points: Gupta et al. (2026) has no page range while Aamir and Lu do (the ACM reference says 20 pages, so "(pp. 1–20)"); Mayer's DOI 10.2307/258792 is the JSTOR DOI (the publisher DOI is 10.5465/amr.1995.9508080335), kept as in Abstract v8. |
| 3 | Structure fidelity to `_review-design-v2.md` | Met, with harmless deviations | See list below. |
| 4 | No em-dashes | Met | Zero U+2014 in the file; en dashes only in page ranges and in the Jongsma title. |
| 4 | Plain English | Mostly met | Three acronyms are never expanded in the body: CISO (opening), VLM (section 2), METI (section 4). "H2" and "H3" are used without being defined. |
| 4 | Politician studies not evidence, with the one-sentence statement | Met | "Results on well-known politicians, whose reputations are settled, are not carried over." Prike is used, which v2 allows (unfamiliar, fictional candidates). |
| 4 | Trust in the AI not an outcome | Met | Trust in the AI appears only as what Gupta measured. |
| 4 | Gupta only in the gap | Met | One mention, in the gap. |
| 4 | Dobbs not cited | Met | |
| 4 | Kim 2026 closest and bounding the AI-penalty claim | Met | "Two results set limits ... their data tie the loss to the error, not to the AI's contradiction" and "D. Kim et al. (2026) is the closest design". |
| 5 | Every in-text citation has an entry and every entry is cited | Met | 34 cited, 34 listed; no orphans either way. |
| 5 | Entry details match the verified sources | Met | Checked against the PDFs: Aamir (CHI 2026, 22 pages, DOI), Bendeck (TVCG 31(1) 1105–1115 from results-6; author version carries the DOI), Chen (IJHCS 213, 103824), Danry (12 pages), Du (17, 1889402), Ferrin (92(4) 893–908), Greitemeyer (e0126316), Hendriks (e0139309), Horne (15, 13896), Kempt (48(4) 222–229), Kim 2004 (89(1) 104–118), Kim 2026 (AIES 2026, N = 157), Loconte (182, 109019), Lowry (72(2) 783–804), Lu (CHI 2025, 19 pages), Madanay (e68823), Mayer (20(3) 709–734), Mickelberg (e0322045), Opdenbusch (NDSS DOI), Proudfoot (22(4), DOI), Schilke (112(42) 12950–12955). Others match the read files or results entries. One spelling to confirm: Gupta's third author is printed "Rajendran" once and "Ranjendran" once in the arXiv PDF; the draft uses "Rajendran" as in verified-view. Check against the ACM DL record. |

**Structure deviations from v2 (item 3), and whether they harm the review:**

1. Word budget: 1,472 against 1,330 planned (boards +38, checkers +56, flags +20, repair +25, gap +16). Inside the limit, so no harm, but there is no slack.
2. Section 1: Lowry's "they believe they ask the right questions", the 40 percent filtering estimate, and Proudfoot's "always on audition", "10 to 15 minutes" and "don't show up the CEO" are dropped. Da Silva and Jensen (optional) dropped. Little harm. "Don't show up the CEO" was the best support for the norm against challenging in the room (H4a), and section 1 now carries H4a only through "too shy" and Aamir.
3. Section 1, "how we differ": v2 named trust, the act of asking (H4a) and the budget decision. The draft names only "measure trust in the CISO". Minor harm, because H4a has no stated motivation anywhere in the review.
4. Schilke reworded as the read file asked (alternatives, not authority; not a variable). No harm to fidelity, but the sentence that says what it means for a powerful listener is missing (fix 8).
5. Section 2: Google (Gemini Live) dropped because the source was secondary; Danry reworded as the read file asked; Wang (optional) dropped. No harm.
6. Section 3: Horne's below-chance result and the "rare failure effect" dropped; von Schenk's "the accused could not answer" folded into the closing "In all of these studies the expert stays silent". No harm.
7. Section 4: Mayer's propensity and "integrity weighs most early" dropped. Kim 2004's "evidence of innocence raised every measure" dropped. Ferrin's "no reply moved perceived competence" dropped. That last point is the one real loss: v2's "What we know" says "replies move integrity, not perceived competence", and our competence item is the one the flags target. Moderate harm to how H2 reads; one clause would restore it if words allow.
8. Section 4: Pezzo reduced to the abstract-level clause, and the gap no longer cites it as a reply paradigm. This follows read-pezzo-schilke-sharma.md and is correct. Shaw (optional) dropped as the read file advised. No harm.
9. Jongsma and Sand (2022) added. It gives section 4 a stated disagreement. Improves the review.
10. Gap: v2's element-by-element "none ..." list is compressed into "No study combines these pieces" plus the Kim, Gupta and Sharma sentences. "None lets the expert answer a flag on content the audience is watching" is now said only for Kim ("the advisor never replied"). Slight harm: the strongest novelty claim is implicit. Hancock (optional) dropped.

## D. Ranked fix list

Each fix gives the old text, then the new text, most serious first, with the word change. Fixes 1 to 7 together add about 4 words (total about 1,476). Fixes 8 and 9 add about 25 more and should be adopted only together with the cut in fix 10.

1. Bendeck figure (claim 21; -2 words)
   - Old: "yet answered only 13 percent of questions that needed a value read off a chart when it lacked the data"
   - New: "yet answered only 23 percent of questions asking for a single value when it lacked the data"

2. Du red highlighting (claim 27; 0 words)
   - Old: "and red highlighting made the report seem more credible and the drop larger"
   - New: "and red highlighting gave the report more weight and lowered the grade further"

3. Kim 2026 modality (claim 49; +1 word)
   - Old: "an AI monitored a human advisor in front of the student who depended on him"
   - New: "an AI checked a human advisor's work in front of the student who depended on him"

4. Opdenbusch attribution (claim 5; +3 words)
   - Old: "Directors are \"too shy to ask questions"
   - New: "Four interviewees said directors are \"too shy to ask questions"

5. Aamir wording (claim 7; +1 word)
   - Old: "Directors pushed back more confidently where independent reviews were routine (Aamir et al., 2026)."
   - New: "Directors felt more confident challenging management where independent reviews were routine (Aamir et al., 2026)."

6. Sharma attribution (claim 43; +1 word)
   - Old: "Most studies they review find that repair falls short of the start"
   - New: "Most repair studies, they note, find that trust falls short of the start"

7. Undefined terms for a standalone review (plain English; about 0 words net)
   - Old: "Hence H2 treats a rise above the first rating as exploratory, and H3 expects a bare denial to repair less."
   - New: "Hence we treat a rise above the first rating as exploratory and expect a bare denial to repair less."
   - Old: "Our flags are drafted with a real VLM and are all wrong for the company."
   - New: "Our flags are drafted with a real vision-language model and are all wrong for the company."
   - Old: "A board member points a camera assistant at the CISO's slide"
   - New: "A board member points a camera assistant at the chief information security officer's (CISO's) slide"
   - Old: "Our three items come from the METI, built on the same factors"
   - New: "Our three trust items come from the Muenster Epistemic Trustworthiness Inventory, built on the same factors"
   - (The four replacements add about 9 words; fix 1 and the H2/H3 change save 3. If that is too many, keep the CISO and VLM changes and drop the METI one.)

8. Schilke's relevance (readability; +11 words; adopt with fix 10)
   - Old: "Their power was having alternatives, not authority."
   - New: "Their power was having alternatives, not authority, so a powerful listener's trust may simply fall faster."

9. Pezzo sentence left hanging at the end of section 4 (readability; +9 words; adopt with fix 10)
   - Old: "A physician who had not consulted an available diagnostic aid"
   - New: "The only evidence on an expert's stated reason about a machine is indirect: a physician who had not consulted an available diagnostic aid"
   - (If "only" cannot be checked on the abstract alone, write "The nearest evidence".)

10. Compensating cut, needed only if fixes 8 and 9 are adopted (-17 words). This is the writer's own first choice.
    - Old: "Ray-Ban Meta glasses answer questions about what the wearer is looking at (Meta, 2024), Apple's Visual Intelligence lets users ask for \"an explanation of a diagram from class notes\" (Apple, 2024), and Copilot"
    - New: "Ray-Ban Meta glasses answer questions about what the wearer is looking at (Meta, 2024), and Copilot"
    - Then delete the Apple reference entry (33 sources remain).

11. Gupta reference (APA consistency; no body words)
    - Old: "In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems*. Association for Computing Machinery. https://doi.org/10.1145/3772318.3791679"
    - New: "In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–20). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791679"
    - Also confirm "Rajendran" against the ACM DL record (the arXiv PDF prints both "Rajendran" and "Ranjendran").

12. Optional, if words allow after the above: restore Ferrin's competence point (deviation 7; +12 words)
    - Old: "A reply that neither confirmed nor denied did no better than the weaker of the other two (Ferrin et al., 2007)."
    - New: "A reply that neither confirmed nor denied did no better than the weaker of the other two, and no reply moved perceived competence (Ferrin et al., 2007)."

## E. Evidence log

Brief and rules: `search/_check-brief.md`, `Instructions.md`, `search/_review-design-v2.md`, `search/fact-check-2026-10-19.md`, `CMPT 863 Project Abstract v8.md` (full, including references).

Checked records: `search/verified-view.md` (entries 1 to 5, 8, 10, 11, 13, 15 to 18), `search/read-pezzo-schilke-sharma.md` (all), `search/read-danry-kempt-dasilva.md` (all), `search/results-5-board-products.md` (camera assistants, Microsoft, Adair entries), reference lines in `search/results-1`, `results-4`, `results-6` (Bendeck, Danry, Horne, Loconte, Lowry, Opdenbusch, Proudfoot, Du).

Full texts read with pdftotext (passages located by search, then read in context):
- SEC 2023 final rule, pp. 82 to 85 (Item 407(j) not adopted; the quoted sentence).
- Lowry et al. 2026: abstract p. 783, introduction p. 784, contribution p. 785 (CISOs not powerful; low expertise as reason for deference), pp. 797 to 798 (D-9 quote), filtering section and Table 3.
- Proudfoot et al. 2023: p. 240 (communication quote).
- Opdenbusch et al. 2025: abstract, contributions, IV-A (budget, "Most participants ..."), IV-B4 ("too shy", four participants), Discussion (power imbalance).
- Aamir et al. 2026: Findings ("Where external assurance was routine ..."), Discussion ("Challenge" paragraph), ACM reference (22 pages).
- Schilke et al. 2015: pp. 12950 to 12953 (Study 1 means, trust drop).
- Danry et al. 2020: pp. 7 to 8 (procedure, 63.5 percent).
- Bendeck and Stasko 2025: abstract, p. 3 (16th percentile), p. 4 (Tables 4 and 5, 13 and 23 percent), p. 5 (best practice).
- Horne and Nevo 2025: method (label text, -1 to +1), Table 2, results text.
- Lu et al. 2025: abstract, section 6 (AI against expert), p. 9 (quote), conclusion.
- Du et al. 2026: abstract, sections 3.4, 4.1 to 4.4, Tables 1 and 2.
- Kim, Daneshi, Draper and Sreedharan 2026: abstract, conditions, hypotheses, sample (N = 157), measures.
- Kempt and Nagel 2022: header and pages 222 to 229 (pagination check); content as in the fact check.
- Kim et al. 2004, Ferrin et al. 2007: header pages (volume, issue, pages, DOI).
- Hendriks et al. 2015: Mayer citation; article number and DOI.
- Mayer et al. 1995: JSTOR header (20(3), 709–734).
- Sharma et al. 2023 (web text): introduction (Kim 2018 sentence), limits of repair (definition, "further development of trust"), denial section, measurement section (8 of 94).
- Gupta et al. 2026 (arXiv): abstract, measures, section 6.3, ACM reference (20 pages, author spelling).
- Loconte et al. 2026 (published): header (182, 109019).
- Madanay et al. 2025 (Europe PMC text): header (e68823); results as in the fact check.
- From `../references/`: Chen et al. 2026 (abstract; header 213, 103824), Greitemeyer and Sagioglou 2015 (Study 1 discussion), Mickelberg et al. 2025 (p. 3, recency), Prike et al. 2025 (abstract, competence and morality results), Strickland 1958 (Table 3, Hypothesis 2), von Schenk et al. 2024 markdown (Choice treatment, 84.91 percent), Reif et al. 2025 markdown (Study 2).

Not read in full text: Pezzo et al. 2022 (abstract only, via the read file), Jongsma and Sand 2022 (via the read file), the product pages and the Adair page (via results-5 and the fact check).

## F. Readability for a CHI reviewer

**Does it read as one argument?** Mostly yes. The section titles form a chain (boards cannot judge, checkers exist and are wrong, wrong flags are believed and cost the person, replies can repair, nobody has tested the combination), and each section ends by pointing at what our study does. The power disagreement in section 1 and the Kempt and Jongsma exchange in section 4 are real synthesis.

**Where it still lists papers:**
- Section 3, paragraph 1: Horne, Lu, Du, von Schenk, one per sentence. A topic sentence that sorts them would help, for example label believed (Horne, Lu) against acted on (Du, von Schenk). The draft's "Flags also move people with power to act" does this halfway.
- Section 3, paragraph 2: Strickland, Prike, Madanay, Chen, Reif in sequence. The paragraph does not say that Prike's and Kim's checkers were right while ours are wrong, which is the contrast the reader needs.
- Section 4, paragraph 2: Kim 2004, Ferrin, Greitemeyer, Mickelberg, Sharma, one per sentence. Mickelberg in particular is a design note, not part of the argument.
- Section 2, the product sentence: three products in one list, which is acceptable for products.

**Where it is vague:**
- The opening calls "the gap where they meet" a fifth strand of work; it is not a strand.
- "H2" and "H3" (section 4), "our three items" and "both our groups" appear before the reader knows the design; the groups are first explained in the last sentence of the review.
- CISO, VLM and METI are never expanded.
- Schilke: the reader is not told what the result means for a powerful listener (fix 8).
- Pezzo: the last sentence of section 4 has no link to what comes before it (fix 9).
- "A flag can thus be right by the textbook and wrong for this company" follows from Bendeck only in part. Bendeck shows right on rules and wrong on values; "wrong for this company" is our design premise, not their finding. Consider "A flag can thus cite a sound rule and still misread the slide."
- The close says what the study combines but not what it will tell us. One clause, such as "to test whether a correct expert's reason wins back what a wrong flag costs", would state the contribution, if words allow.
