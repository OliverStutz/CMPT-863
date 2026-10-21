# Design review B: structure of the literature review (proposal v1)

Reviewer B, 2026-10-19. Brief: `_design-brief.md`. Worked alone; reviews A and C not read.

## A. Answers

### Q1. Does the proposal meet every requirement in Instructions.md?

**With changes.**

| Requirement (Instructions.md) | Status | Why |
|---|---|---|
| Expand the abstract's related work | met | All 14 abstract sources kept, 15 added. |
| Situate the project within prior work | met | Every section ends with "how we differ"; section 5 states the combination. |
| Scholarly research | met | 21 read scholarly sources plus two benchmarks. |
| Industry, trade, products where appropriate | met, with a caveat | SEC 2023 was read in full. NACD & ISA 2023 was not: only a press release and a secondary blog were seen (results-5), so the "scripted questions" claim rests on second-hand text. Two of the camera-assistant entries (ChatGPT, Gemini Live) are verified only through Neowin and 9to5Google, while the proposal cites "Google 2025". Cite the primary page or name the secondary one. |
| At least 15 sources | met | 29 without optional items. The risk runs the other way (see Q6, Q8). |
| Thematic subsections | met | Five themes plus an opening. |
| Synthesize groups, do not summarise one by one | partly met | Sections 3 and 4 are written as lists of papers, each with its own numbers (nine sources in 300 words for section 3, eleven for section 4). In that space the draft will turn into one paper per sentence, which is the pattern the abstract's appendix already follows. Each section needs one synthesis claim first, with papers cited as support, and numbers kept to two or three per section. |
| Four questions: what exists, what we know, what is unresolved, how we differ | met, with changes | Section 2's "unresolved" (no product reports audience trust) is real but thin. Section 4's "unresolved: repair after a machine accusation" needs hedging: Pezzo et al. (2022) varied a physician's stated reason for setting a computer aid aside, and Kim et al. (2026) found the AI's involvement did not change the penalty. Section 3's "what we know: an AI's contradiction costs the contradicted expert" is contradicted by one of its own sources (Kim et al. 2026; see Q2). |
| Readers see coverage, a meaningful gap, a novel contribution | met | Section 5. The novelty tally's "narrowed claim" (the combination, not any one element) is followed. |
| Under 1,500 words (around 1,000 is typical) | met on paper | 1,290 planned. With 29 sources and APA author-date citations, in-text citations alone add roughly 150 to 200 words if they count. See Q8. |
| Citation format per the proposal guidelines | not verifiable | The proposal assumes "APA as in the abstract". The proposal guidelines are not among the files read. Products and the SEC release need APA web and report formats (results-5 gives them). |

### Q2. Is each bundle coherent?

| Section | Coherent? | Papers that do not fit, and why |
|---|---|---|
| 1 Boards oversee what they cannot judge | **yes** | All six answer "how do non-expert boards oversee an expert CISO". Aamir et al. (2026) also serves H4a, which is fine. |
| 2 Checkers are available and often wrong | **yes** | Products and benchmarks answer "do such tools exist and how reliable are they". Squash is a political-speech checker, but it is cited for checker error, not for any result about politicians, so it does not conflict with the exclusion. Weakness: no HCI research system of a listener-held checker (see Q7). |
| 3 People follow a wrong flag, and the person flagged pays | **no** | (a) **Kim et al. (2026)** is filed as support for "the flagged person's standing falls", but its result is the opposite on that point: advisor errors lowered trust, and how the AI was involved changed nothing (no main effect of modality, Bayes factors favouring exclusion, full text section 5). It is the boundary result for H1 and the null for H4b, and belongs next to Loconte et al. (2026) as a counter-result. (b) **Strickland (1958)** answers a different question: his checks found nothing and there was no flag, so the paper is about the act of checking, not about a flag being believed. It fits section 3 only as a separate point on the checker's presence, which the design tests at slide 1. Keep it, but label it so. |
| 4 Trust in an expert, and how it is repaired | **no** | (a) **von Schenk et al. (2024)** is about what the tool's holder does (accusation rate rises when a fallible detector says "lie"). The accused could not reply, so nothing is repaired. It is a "people follow a fallible flag" result for the holder's actions, so it belongs in section 3. (b) **Reif et al. (2025)** is a penalty for the AI user. It justifies a design choice (the CISO never mentions AI) and answers none of section 4's questions. (c) **The burden-of-proof point attributed to Chen et al. (2026)** is Chen et al. citing Kempt & Nagel (2022) (Chen full text, line 382: "disputing or rejecting AI's recommendations places the burden of proof on doctors (Kempt et al., ..."). The section should cite the primary source. |
| 5 Gap | **yes** | Kim et al. (2026), Kim et al. (2004) and Gupta et al. (2026) each mark one edge of the gap. |

### Q3. Moves

1. **von Schenk et al. (2024): from section 4 to section 3 (part A)**, because it shows people acting on a fallible tool's verdict against a person (accusations rose to almost 85 percent when the detector said "lie"). That is the H1 and H4a mechanism, not repair. Together with Du et al. (2026), where intention to intervene rose under the flag, it lets section 3 cover both belief and action, which is what H4a needs.
2. **Kim et al. (2026): from section 3 part B (the flagged person pays) to section 3's counter-result and boundary paragraph, next to Loconte et al. (2026)**, because its result is that the AI's involvement did not change the trust penalty. Section 5 still names it as the closest design. Part B's claim should then be qualified: the cost appears when the AI disagrees with a human the observer cannot check (Madanay et al. 2025; Chen et al. 2026), not when it confirms an error that happened anyway.
3. **Reif et al. (2025): from section 4 to one clause in section 3 part B (AI involvement colours how a person is judged), or out of the review into the method** (see Q6). In section 4 it does no work.

Not a move, a replacement: in section 4, cite Kempt & Nagel (2022) for the burden of proof (see Q7). Chen et al. (2026) stays in section 3.

### Q4. Alternative bundlings

- **Alt A (merge 2 into 3): reject.** Section 2 is where the review meets the instructions' request for industry and product sources. Merged into section 3, the products shrink to a clause, and section 3, already the most crowded, gains a third part.
- **Alt B (separate H4 section): reject.** After move 1, section 3 carries the holder's action (von Schenk et al. 2024; Du et al. 2026) and section 1 carries the boardroom silence (Opdenbusch et al. 2025; Aamir et al. 2026). A separate section would double-cite four papers and cost about 150 words for H4a, which is a secondary hypothesis, and H4b, which is exploratory.
- **Alt C (fold 5 into 4, gap into the opening): reject.** The instructions say readers should understand the gap "by the end of the review". A closing synthesis is what CHI related-work sections do, and section 4 is not the right home for a gap that also involves sections 1 to 3.

### Q5. Optional sources

| Source | Verdict | Section | Why |
|---|---|---|---|
| Da Silva & Jensen (2022) | only if words allow | 1 | A CSCW anchor for the CISO's interpretive power and for "cyber sophistry", which backs Lowry et al.'s filtering. But Opdenbusch et al. already give the "dark art" quote, and the paper is unread. |
| Shaw, Wild & Colquitt (2003) | only if words allow | 4 | It is a useful caution for H2: our reasoned answers are justifications, which help less than excuses. It is also unread and needs an SFU fetch. |
| Sharma, Schoorman & Ballinger (2023) | include (one clause, after reading) | 4 | It gives H2 its vocabulary: partial, complete and over-repair, measured against a pre-violation baseline. That is exactly why "above the start" is exploratory. |
| Pezzo et al. (2022) | include (one sentence, after reading; open access) | 4 | It is the only pool study in which a professional's own stated reason about a computer aid changed lay ratings: "I knew what it would say" helped and "I don't trust it" hurt. That is the nearest machine-context precedent for H2 versus H3, and without it the section 4 claim that repair after a machine is unstudied is overstated. |
| Hancock, Naaman & Levy (2020) | only if words allow | 2 or 5 | One sentence would frame the checker as receiver-side AI-mediated communication, a useful HCI positioning. It does no evidential work. |

### Q6. Surplus sources

- **Reif et al. (2025):** a design rationale (why the CISO never mentions AI), not evidence on any of the five themes. Cite it in the method, not the review, or as one clause after move 3.
- **NACD & ISA (2023):** unread (press release only), and Lowry et al. (2026) already make the same point with field evidence: non-experts ask questions off professional bodies' lists and cannot follow up. Keep it only if the handbook is actually read, and then only as the industry form of that norm.
- **Wang et al. (2024, CharXiv):** borderline. Bendeck & Stasko (2025) already make the point on chart images, and theirs is a visualization paper whose full text is in hand. CharXiv adds one realistic-chart error rate and costs little, so keep it only if section 2 has room.
- **Mickelberg et al. (2025):** not surplus, but its unique contribution is a method caution (rating at every step draws attention to the latest information). One clause, not a sentence of results.

### Q7. Missing themes

1. **HCI research systems for listener-held, real-time checking (what exists).** Section 2 has products and benchmarks but no research prototype, and Gupta et al. (2026) is held for the gap by the author's decision. For an HCI course this is the most visible hole. **Danry et al. (2020), Wearable Reasoner** (results-1, Augmented Humans, open access) is a wearable that judges heard claims. It measured agreement with statements, not trust in the speaker, so it supports the gap. It was verified from the abstract only, so read it before citing.
2. **The burden of proof when an AI contradicts an expert (the mechanism for H2, H3, H4b).** Currently carried by a second-hand sentence in Chen et al. (2026). **Kempt & Nagel (2022)** (results-4, Journal of Medical Ethics, open access) is the primary source, and the replies in the same issue dispute the asymmetry. That dispute is an honest "unresolved" for section 4.
3. **Power and trust as a mechanism (optional).** The review asserts that the holder's power matters, but its only sources for power are interviews and Strickland. **Schilke, Reimann & Cook (2015)** (results-3, PNAS, open access: more power, less trust in the partner) would give section 1 or 3 one experimental clause. Only if words allow.
4. **Hedged accusation and the denial that repeats the charge (optional, H3 exploratory part).** **Wegner et al. (1981)**: innuendo lowers impressions whatever the source's credibility, and a denial headline does too, more weakly. This speaks to whether a bare denial can make trust fall further, which section 4 lists as unresolved. It was verified from an abstract extract only, so cite it only if read.

### Q8. Word budget

The proposal's 1,290 is realistic in total but misallocated. Section 2 gets 220 words for one claim, while sections 3 and 4 get 300 each for nine to eleven sources and four questions. Repeating "how we differ" in every section and again in section 5 also wastes about 80 words. My budget, with in-text citations counted:

| Section | Words |
|---|---|
| 0 Opening | 60 |
| 1 Boards oversee what they cannot judge | 230 |
| 2 Checkers exist, and they err | 170 |
| 3 Wrong flags are believed and acted on | 360 |
| 4 Trust in an expert and its repair | 360 |
| 5 Gap and contribution | 150 |
| **Total** | **1,330** |

That leaves about 170 words for Sharma et al. (2023), Pezzo et al. (2022), Kempt & Nagel (2022) and Danry et al. (2020), which take one clause or sentence each. Sections 1 to 4 should end with one sentence of difference, and section 5 should state the combination once instead of repeating the lists.

### Q9. Order and author decisions

**Order: keep.** It runs setting, tool, the flag's effect (H1, H4a), the answer's effect (H2, H3), gap. That is the order of the briefing and of the hypotheses.

Author decisions, stated plainly:

- **The politician exclusion together with declining Dobbs et al. (2026) leaves the exclusion without evidence.** Dobbs found that a correction cut a doctor's credibility more than a politician's. That is the one result in the pool showing that politician findings understate the effect on experts, which is why they do not transfer. Without it the exclusion rests on the author's judgement alone. And Prike et al. (2025), a study of candidates, remains, which an informed reader may see as inconsistent. At minimum, section 3 needs one uncited scope sentence: results on known politicians are set aside because settled reputations do not carry over, and Prike's candidates were fictional and unfamiliar.
- **Excluding Chae & Tewksbury (2026)** removes the only direct null for H1 on competence ratings, which all three selection reviewers chose. The harm is reduced if section 3 keeps Kim et al. (2026) and Loconte et al. (2026) as visible counter-results (move 2).
- **Gupta gap-only** is harmless if Danry et al. (2020) fills section 2's research-system slot. Without that, section 2 has no HCI system at all.

### Q10. Confidence

**Medium.** The structural calls (moves 1 to 3, the Kim 2026 misreading, the second-hand burden-of-proof source) are checked against full texts and I hold them with high confidence. The verdicts on the optional and missing sources rest on results-file entries verified from abstracts only.

## B. Revised structure

**0. Opening (60).** The scenario in two sentences, then one line per theme.

**1. Boards oversee what they cannot judge (230).** SEC (2023); Lowry et al. (2026); Opdenbusch et al. (2025); Proudfoot et al. (2023); Aamir et al. (2026). Optional: Da Silva & Jensen (2022); Schilke et al. (2015); NACD & ISA (2023) only if read.
Argument: regulators assume non-expert directors, who hold formal power, so they judge the presenter's credibility rather than the substance and stay silent for fear of exposure. An independent check raises willingness to challenge, but all of this comes from interviews, with no tool and no trust measure.

**2. Checkers exist in the listener's hand, and they err (170).** Camera assistants (Meta 2024; Apple 2024; Google 2025, primary pages); Microsoft Copilot in Teams; Squash (Adair 2020); Danry et al. (2020); Bendeck & Stasko (2025); Wang et al. (2024) if room. Optional: Hancock et al. (2020).
Argument: any listener can now point a model at a speaker's material, and the models misread charts while knowing the textbook rules. Neither products nor prototypes report what a check does to trust in the speaker.

**3. Wrong flags are believed and acted on, and the flagged person pays (360).**
- Belief and action: Horne & Nevo (2025); Lu et al. (2025); Du et al. (2026); von Schenk et al. (2024).
- Standing of the person: Strickland (1958, checking alone); Prike et al. (2025); Madanay et al. (2025); Chen et al. (2026); Reif et al. (2025) as one clause if kept.
- Counter-results and boundary: Kim et al. (2026, AI involvement did not change the penalty); Loconte et al. (2026, resistance when accuracy and confidence are shown), plus the uncited scope sentence on politician studies.

Argument: opaque flags are believed and acted on by evaluators with power, and an AI's disagreement costs a human the observer cannot check. Where the AI only confirms a real error, or shows its own uncertainty, the cost vanishes or reverses. In every study the expert is silent and rated once. This section frames H1 as a replication in a new setting.

**4. Trust in an expert, and how a reply repairs it (360).** Mayer et al. (1995); Hendriks et al. (2015); Kim et al. (2004); Ferrin et al. (2007); Greitemeyer & Sagioglou (2015); Mickelberg et al. (2025, one clause); Sharma et al. (2023); Kempt & Nagel (2022); Pezzo et al. (2022). Optional: Shaw et al. (2003).
Argument: trust in an expert rests on competence, integrity and benevolence. A reply repairs trust only when it supplies what the accusation removed; bare replies do not, and clearing returns trust to baseline, rarely above. An AI contradiction is argued to put the burden of proof on the expert, and an expert's stated reason about a computer aid changes lay ratings. No study has the expert answer a machine's flag and rate trust before and after. Carries H2, H3 and H4b.

**5. Gap and contribution (150).** Kim et al. (2026); Kim et al. (2004); Gupta et al. (2026).
Argument: the elements exist separately (wrong flags, AI disagreement penalties, reply-based repair, boards' dependence). The combination does not: a checker held by the powerful non-expert, flags wrong for the company, the expert's own answer, and trust rated at each step against a neutral group.

## C. Evidence log

**Files read in the brief's order:** `../Instructions.md`; `../../CMPT 863 Project Abstract v8.md`; `_review-design-v1.md`; `verified-view.md` (all 25 entries and 22 rules); `results-5-board-products.md` (whole); entries in results-1 to results-4 and results-6 for Pezzo 2022, Shaw 2003, Sharma 2023, Hancock 2020, Mieczkowski 2021, Kempt & Nagel 2022, Danry 2020, 2023, 2025, Hendriks 2016, Wegner 1981, Schilke 2015, Drolsbach 2024, Serra-Garcia & Gneezy 2025, de Brito Duarte 2025, Navajas Fernández 2026, Hou & Jung 2021; `_master-index.md`; `novelty-tally.md`; `../Literature Review - Candidate papers.md`.

**Claims checked against full texts (pdftotext):**

| Claim in the proposal | Source checked | Result |
|---|---|---|
| Kim 2026: N = 157; AI involvement did not change the penalty | Kim-2026 AIES, pp. results 5.2 to 5.5 | Confirmed: N = 157; no main effect of modality, all p ≥ .168, Bayes factors favour exclusion; advisor errors drove trust. This supports move 2: it is a null for the AI's role, not support for "the flagged person pays". |
| Chen 2026: AI disagreement shifts the burden of proof onto the expert | Chen-2026 (abstract-phase references), line 382 and line 1193 | Chen states it citing Kempt et al. The primary source is Kempt & Nagel (2022). |
| SEC 2023: "directors with broad-based skills" suffice | SEC 33-11216 | Confirmed: "directors with broad-based skills in risk management and strategy often effectively oversee ... without specific subject matter expertise". |
| Lowry 2026: "all they can do is assess the credibility of the executives in front of them"; median 40 percent filtering | Lowry Management Science | Both confirmed (D-9 quote; CISO survey median 40%). |
| Proudfoot 2023: "always on audition", 10 to 15 minutes, "don't show up the CEO" | Proudfoot MISQE | All three confirmed. |
| Aamir 2026: directors challenge more confidently where independent reviews are routine | Aamir 2026 (markdown full text) | Confirmed: "Where external assurance was routine, directors reported greater confidence in challenging management perspectives". The paper's word is "external assurance", not "independent review". |
| Bendeck & Stasko 2025: about the 16th percentile; cannot read values without data | Bendeck TVCG | Confirmed (16th percentile; 13% accuracy without data). |
| Horne & Nevo 2025: rare failure effect | Horne & Nevo Sci Rep | Phrase confirmed. |
| Du 2026: 1.8-point drop on 10 for an identical paper | Du Front Psych, section 4.3 | Means 6.83, 6.50, 5.02, 4.26 (groups A to D). The 1.8 contrasts an 87 percent report with a 7 percent report, not a report with no report. Word it that way in the review. |
| Loconte 2026: accuracy shown on screen | Loconte CHB (published) | Confirmed: 54 vs 89 percent accuracy displayed, matching the model's real hit rate. |
| Gupta 2026: rated trust in the AI, never the speaker | Gupta arXiv | Measures are "trust in AI-provided information" and "Trust in AI"; no speaker rating found in a search of the text. |

**Not checked against full text:** NACD & ISA 2023 (not available), the product pages, CharXiv (no PDF in references), and all optional and missing-theme sources, which come from results-file entries only.
