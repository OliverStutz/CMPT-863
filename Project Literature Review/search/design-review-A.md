# Design review A: structure proposal v1 for the related work

Reviewer A, 2026-10-19. Brief: `_design-brief.md`. Read in the order the brief gives; did not read design-review-B or design-review-C.

## A. Answers

### Q1. Does the proposal meet every requirement in Instructions.md?

**With changes.**

| Requirement | Verdict | Why |
|---|---|---|
| Expand the abstract's related work | met | All 14 abstract sources are in the pool, and 7 read papers, Gupta, 5 trade or product sources and 2 VLM benchmarks are added. |
| Scholarly research plus industry, trade, books, products where appropriate | met | SEC 2023, NACD & ISA 2023, camera assistants, Copilot in Teams and Squash cover regulation, trade and products. Caveat: NACD is known only from a press release and a secondary blog (results-5); it should not carry a claim until the handbook is read. |
| At least 15 sources | met | 29 without the optional ones. |
| Thematic subsections | met | Five themes plus an opening. |
| Synthesize, do not summarize one by one | partly met | The bundles are thematic, but sections 3 and 4 list each paper with its own finding (9 and 11 sources in 300 words each). At that density the text will fall back into one sentence per paper, which is what the abstract appendix already does. The fix is structural (fewer sources per paragraph, grouped citations behind one claim), see Q8 and part B. |
| Each theme answers what exists, what we know, what is unresolved, how we differ | mostly met | All four slots are filled for every section. Section 2's "how we differ" is a design statement ("our flags are drafted with a real VLM"), not a way the project extends research; it should say what no product or benchmark has measured (the speaker's trust). Section 4's third bundle does not answer section 4's question at all (see Q2). |
| Readers see a meaningful gap and a novel contribution | met, with two wording risks | Section 5 lists the elements. Two claims are too broad as written: "none lets the expert answer" is false without "after an AI flag" (Kim et al. 2004 and Ferrin et al. 2007 let the accused answer a human accusation; Pezzo et al. 2022, if used, has a physician's stated reason concerning an aid); and "a wrong AI flag on a correct expert" is true by construction only, since participants never learn the flag is wrong (novelty-tally, Q6 point 5). Word it as an unverified AI flag that the researchers know to be wrong. |
| Under 1,500 words, around 1,000 as a benchmark | met | 1,290 planned. The optional sources would push it to about 1,450, leaving no margin. |
| Citation and reference format as in the proposal guidelines | met in intent | APA as in the abstract is named. Product pages need APA web entries (Microsoft support page as n.d., with retrieval date); "camera assistants as one source" must still be three reference entries. |

Two factual points in the proposal's bundle text to correct before drafting (checked against full texts, see part C):

- Madanay et al. (2025): the radiologist who undercalled the AI was not "rated the worst provider". The rating scale ran from 1 (worst provider possible) to 10 (best); that radiologist got the lowest mean of the four conditions (6.82 against 7.45 to 7.58). Write "rated lowest".
- "Chen et al. 2026's point that AI disagreement shifts the burden of proof onto the expert": Chen et al. take this from Kempt and Nagel (2022), whom they cite for it (Chen full text, introduction and discussion). Either cite Kempt and Nagel (2022) directly (results-4 entry, conceptual paper) or write "Chen et al. (2026), following Kempt and Nagel (2022)". The same fix applies to the abstract's sentence.

### Q2. Is each bundle coherent?

| Section | Coherent | Papers that do not fit |
|---|---|---|
| 1 Boards oversee what they cannot judge | yes | None. Every source speaks to how non-expert boards judge and question a CISO. NACD is the weakest (unread). |
| 2 Checkers are available, and they are often wrong | yes | None. Products and benchmarks both answer "can a listener check a slide live, and how reliable is the check". |
| 3 People follow a wrong flag, and the person flagged pays | yes | All fit the question. Weakest fit: Strickland (1958), whose check found nothing and raised no flag; it fits only as "even a check that finds nothing costs trust", and should be framed that way. Part B's flags (Prike, Madanay, Kim 2026) were correct or right by convention, so the title's "wrong flag" applies to part A only; the title should not promise more. |
| 4 Trust in an expert, and how it is repaired | no | von Schenk et al. (2024) is about accusation behaviour by the tool's holder (H4a), not about trust or its repair. Reif et al. (2025) is an AI-use penalty that justifies a design detail (the CISO never mentions AI); it answers neither "what trust is" nor "how it is repaired". The burden-of-proof point is attributed to the wrong paper (Kempt and Nagel 2022, not Chen). Mayer (1995) and Hendriks (2015) fit the title but do their work earlier: section 1 already reports that boards judge "the credibility of the executives" (Lowry D-9), which is Hendriks's epistemic trustworthiness, and section 1's budget decision is Mayer's risk taking. |
| 5 The gap and this study | yes | Kim et al. 2026, Kim et al. 2004 and Gupta et al. 2026 all serve the gap. |

### Q3. Moves

Four moves.

1. **Mayer et al. (1995): from section 4 to section 1**, because section 1 is where the review explains what a non-expert board actually judges (the presenter's ability, benevolence and integrity) and acts on (the budget as risk taking in the relationship); placed there, the theory answers the instructions' "what theories exist" for the setting instead of arriving after section 3 has already reported trust losses.
2. **Hendriks et al. (2015): from section 4 to section 1**, because epistemic trustworthiness is defined as a layperson's judgement of whether to accept an expert's claim they cannot check, which is exactly Lowry et al.'s (2026) D-9 line and Proudfoot et al.'s (2023) "perceived credibility of the CISO"; it names the outcome variable where the setting first raises it.
3. **von Schenk et al. (2024): from section 4 to section 3 (part A)**, because it shows that people with power to accuse follow a fallible detector into action (85 percent accused when it said "lie"); that is "wrong flags are followed", extended from belief (Horne & Nevo; Lu) and grades (Du) to an act, and it carries H4a next to Du's intention to intervene. It has nothing to say about repair.
4. **Reif et al. (2025): from section 4 to section 3 (part A) as a half sentence on the social cost of acting on an AI, or out of the review** (see Q6), because its only role is to justify why the CISO never mentions AI and why the holder's naming of the checker is logged; that is a method detail.

Not a move but a re-attribution: the burden-of-proof sentence stays in section 4 as the bridge from flag to answer, cited to Kempt and Nagel (2022).

### Q4. Alternative bundlings

- **Alt A: reject.** Section 2 is the only home for the products and trade sources that the instructions name, and merging it into section 3 would make section 3 the longest section with the most heterogeneous evidence (products, benchmarks, three experiments, five person-rating studies).
- **Alt B: reject.** H4a is carried well enough by section 1 (silence from fear of exposure; independent review raises confidence) and section 3 part A after move 3 (von Schenk; Du's intention to intervene), and H4b is exploratory with Kim et al.'s (2026) null against it, so a separate section would spend 150 or more words double-citing four papers for one confirmatory hypothesis.
- **Alt C: reject.** The instructions ask that readers understand the gap "by the end"; an element-by-element gap in the opening states conclusions before the reader has seen the evidence, and folding the gap into section 4 hides the parts of it that come from sections 1 to 3 (no tool in the boardroom, no product evidence on speaker trust, silent experts).

### Q5. Optional sources

| Source | Verdict | Section |
|---|---|---|
| Da Silva & Jensen (2022), CSCW | include | 1. It is the only SIGCHI paper on the CISO and the board, and its "interpreter of a dark art" in a "superior and precarious" role gives the informational-power half of rule 7 a scholarly source rather than only interviews in a security venue. Full text is in `references/`; read it before citing. |
| Shaw, Wild & Colquitt (2003) | only if words allow | 4. One clause: our reasoned answers are justifications, which the meta-analysis finds weaker than excuses; it tempers H2. Full text missing. |
| Sharma, Schoorman & Ballinger (2023) | only if words allow | 4. Gives the partial, complete and over-repair vocabulary for H2's "above the start", and signals coverage of the repair literature in one citation. Full text missing. |
| Pezzo et al. (2022) | include | 4 (and it sharpens 5). It is the only pool source in which an expert's stated reason concerning a decision aid changes how observers rate the expert ("I knew what it would say" helped, "I don't trust it" hurt). Leaving it out makes "none lets the expert answer" overclaim; including it narrows the gap honestly to a live answer to an AI flag. Full text missing; read before citing. |
| Hancock, Naaman & Levy (2020) | include | 5. One clause positions the checker in HCI terms as receiver-side AI-mediated communication, which no other source does and which a CHI-style related work is expected to state. |

### Q6. Surplus sources

- **Reif et al. (2025).** In the proposal it justifies a design choice (the CISO never mentions AI). In a 1,500-word review it would be cited without advancing any of the four questions. Keep it in the method, or keep half a sentence in section 3 if the author wants all abstract sources retained.
- **NACD & ISA (2023)**, conditionally. Its one claim (directors ask scripted questions) is already made with field evidence by Lowry et al. (2026), who describe question lists from professional bodies and show they produce the first question but not the follow-ups. NACD adds an industry voice only if the handbook is actually read; otherwise it is decoration and an unverified citation.
- **Wang et al. (2024) or Bendeck & Stasko (2025)**, marginally. Both support the same sentence. Bendeck does more work (textbook rules known, values misread without the data, which is the "right by the textbook, wrong for this company" profile); CharXiv only adds a headline accuracy. Keep both if the sentence stays short; cut CharXiv first under word pressure.

### Q7. Missing themes

1. **Power and trust.** The novelty rests on the checker sitting with the party that holds power, yet no subsection has a source on what power does to trust; section 1 has only interview statements about formal and informational power. Schilke, Reimann & Cook (2015, PNAS; results-3) show that the more powerful party in an exchange trusts its partner less, through less motivated perception of benevolence. One sentence in section 1 lets the review say why the holder's power matters, and Strickland's (1958) supervisor connects to it.
2. **Listener-held checkers in HCI research.** Section 2 has products and benchmarks but no research system, and Gupta et al. (2026) is held back for the gap. For "what approaches already exist", a CHI reader will expect the wearable checker line: Danry et al. (2020, Wearable Reasoner; results-1) judged heard claims for the wearer and measured agreement with the statements, not trust in the speaker. One clause in section 2, with Gupta in section 5 as decided.

Nothing else is missing for the four questions. Continued influence after a correction is covered without Thorson by Greitemeyer & Sagioglou (2015) and Mickelberg et al. (2025).

### Q8. Word budget

The proposal's total is realistic; its distribution is not. Sections 3 and 4 hold 20 source uses in 600 words (about 30 words per source), which forces one-by-one summary, while the opening's two scenario sentences plus five one-sentence theme previews will not fit in 60 words.

| Section | Proposal | My budget |
|---|---|---|
| 0 Opening (scenario, scope: trust in the human, not in the AI; no theme previews) | 60 | 60 |
| 1 Non-experts who must trust an expert (with Mayer, Hendriks, Da Silva, Schilke) | 260 | 250 |
| 2 Checkers in the listener's hand, and how often they err | 220 | 170 |
| 3 Flags are believed and acted on, and the flagged person pays | 300 | 330 |
| 4 Repairing trust after an accusation | 300 | 330 |
| 5 The gap and this study | 150 | 190 |
| Total | 1,290 | 1,330 |

The 170-word margin is for Shaw or Sharma, or for the drafting overrun that dense sections always produce.

### Q9. Order and author decisions

**Keep the order.** It follows the study's own sequence (setting, tool, flag, answer, gap) and the hypotheses in order (H4a context, H1, H2 and H3), which lets each section's "how we differ" point forward to the next.

Author decisions that cost the review something:

- **Gupta et al. (2026) gap-only.** It is the closest research system to a listener-held checker; keeping it out of section 2 leaves "what approaches exist" answered only by products and benchmarks. Citing it in section 2 as an approach (without measuring trust in the AI, which rule 10 already excludes) would not breach the decision's purpose. If the decision stands, Danry et al. (2020) should take that place (Q7).
- **Politician studies excluded.** Acceptable, but section 3 then has only Loconte et al. (2026) as a counter-result to H1, and that one depends on a visible accuracy display. Say in one clause that results on known politicians are not carried over and why (settled reputations), so the reader does not think the counter-literature was missed.

### Q10. Confidence

**Medium.** The coherence and move judgments rest on full reads of the brief, instructions, abstract, proposal, verified view, novelty tally and candidate-paper tally, plus full-text checks of the claims listed in part C. The results files were read for the entries named here and results-5 in full, not every entry of results-1 to results-4 and results-6.

## B. Revised structure

**0. Opening (60 words).** A board member points a camera assistant at a CISO's slide; the tool flags it; the CISO answers. Scope sentence: the review is about trust between the two humans when one of them holds a checker, not trust in the AI.

**1. Non-experts who must trust an expert (250 words).** SEC 2023; NACD & ISA 2023 (only if read); Lowry et al. 2026; Opdenbusch et al. 2025; Proudfoot et al. 2023; Da Silva & Jensen 2022; Aamir et al. 2026; Hendriks et al. 2015; Mayer et al. 1995; Schilke et al. 2015. Argument: regulation expects non-experts to oversee an expert, so what boards judge is the presenter's epistemic trustworthiness (ability, integrity, benevolence) and what they do with it is the budget, a risk taken on trust; they stay silent from fear of exposure, power is split between the board's formal and the CISO's informational power, and independent review raises directors' confidence; all evidence is interviews, with no tool and no trust measure.

**2. Checkers in the listener's hand, and how often they err (170 words).** Camera assistants (Meta 2024; Apple 2024; Google 2025); Microsoft Copilot in Teams; Squash (Adair 2020); Danry et al. 2020; Bendeck & Stasko 2025; Wang et al. 2024. Argument: listener-held live checking ships today and has a research lineage, but a VLM reading a slide image knows textbook rules and misreads values, and a deployed live checker produced more wrong matches than right ones; neither products nor prototypes have measured what a check does to trust in the speaker.

**3. Flags are believed and acted on, and the flagged person pays (330 words).** Part A: Horne & Nevo 2025; Lu et al. 2025; Du et al. 2026; von Schenk et al. 2024; boundary Loconte et al. 2026 (Reif et al. 2025 half a sentence, optional). Part B: Strickland 1958; Prike et al. 2025; Madanay et al. 2025; Chen et al. 2026; Kim et al. 2026. Argument: an opaque flag is believed by default even when wrong, and it moves acts as well as beliefs (grades, intervention, accusation); a check or an AI contradiction costs the person checked; in every study the person stays silent and trust is rated once, and the flags that cost a person standing were correct, so what an unverified wrong flag does to a correct expert who can answer is open. Carries H1 (as a test in a new setting) and H4a.

**4. Repairing trust after an accusation (330 words).** Kim et al. 2004; Ferrin et al. 2007; Greitemeyer & Sagioglou 2015; Mickelberg et al. 2025; Pezzo et al. 2022; Kempt & Nagel 2022 (burden of proof, as cited in Chen et al. 2026); Shaw et al. 2003 or Sharma et al. 2023 if words allow. Argument: a reply repairs trust when it supplies the signal the listener weighs most, a reply with no information does not, and replies move integrity rather than perceived competence; clearing restores to baseline, rarely above; all accusers were human and most clearing came from a third party; an AI flag is argued to shift the burden of proof onto the expert, but no study has tested whether the expert's own reason meets it. Carries H2, H3 and H4b.

**5. The gap and this study (190 words).** Kim et al. 2026 (closest); Kim et al. 2004 (repair paradigm, human accuser); Gupta et al. 2026 (nearest real-time checker, trust in the AI only); Hancock et al. 2020 (receiver-side AI-mediated communication). Argument: element by element, with the qualifiers "after an AI flag" and "unverified from the participant's side"; the contribution is the combination (listener-held checker with the party in power, plausible flags wrong for the context, the expert's reasoned answer against a bare denial, trust after each flag and answer against a neutral group) in a setting where no tool has been tested.

## C. Evidence log

Files read in full: `../Instructions.md`; `../../CMPT 863 Project Abstract v8.md`; `_review-design-v1.md`; `verified-view.md` (entries 1 to 25 and rules 1 to 22); `novelty-tally.md`; `../Literature Review - Candidate papers.md`; `_master-index.md`; `results-5-board-products.md`.

Results entries read: results-1 (Hancock 2020, Mieczkowski 2021, Navajas Fernández 2026, Danry 2020, Serra-Garcia & Gneezy 2025, de Brito Duarte 2025, Gupta note); results-2 (Hou & Jung 2021); results-3 (Hendriks 2016, Sharma 2023, Shaw 2003, Wegner 1981, Ecker 2022, Tenney 2007, Schilke 2015); results-4 (Pezzo 2022, Kempt & Nagel 2022); results-6 (Aird 2018, Swire 2017, Drolsbach 2024).

Claims checked against full texts (pdftotext):

| Claim in the proposal | Source checked | Result |
|---|---|---|
| Madanay: undercalling radiologist "rated the worst provider" | Madanay 2025, Table 2 and scale note | Wrong wording. Lowest mean rating (6.82 on 1 to 10 where 1 = worst provider possible) against 7.45 to 7.58; agreement 4.01 against 4.55 to 4.63 confirmed. |
| Chen's point that AI disagreement shifts the burden of proof | Chen 2026, introduction and discussion | Chen attribute it to Kempt & Nagel (2022) and Kempt et al. (2023). Re-attribute. |
| SEC: "directors with broad-based skills" | SEC 2023 release | Confirmed ("directors with broad-based skills in risk management and strategy often effectively oversee ... without specific subject matter expertise"). |
| Lowry: "mostly symbolic"; D-9 "all they can do is assess the credibility"; median 40 percent | Lowry 2026 | All three confirmed. |
| Bendeck: about the 16th percentile | Bendeck & Stasko 2025 | Confirmed. |
| Horne & Nevo: "rare failure effect" | Horne & Nevo 2025 | Confirmed (term coined in the discussion). |
| Du: 1.8 points on 10 | Du 2026, table | Confirmed (6.83 to 5.02; 4.26 with highlighting). |
| Lu: people cannot tell right from wrong flags | Lu 2025, conclusion | Confirmed ("people lack the capability to differentiate the correctness of AI predictions"). |
| Loconte: accuracy shown on screen | Loconte 2026 (CHB) | Confirmed ("The accuracy displayed on the screen matched ..."). |
| Kim 2026: N = 157; AI involvement did not change the penalty; trust anchored to the human | Kim 2026, abstract and results | Confirmed. |
| Aamir: directors more confident where independent review is routine | Aamir 2026, discussion | Confirmed in substance; the paper's words are "where external challenges are routine, directors feel more confident in their judgments". |
| Gupta: trust measured in the AI, not the speaker | Gupta 2026 | Confirmed ("measured trust in AI-provided information"). Interviewees also called the wearable inappropriate with doctors and impolite in personal conversation, a social-cost point the gap sentence can use. |
| Proudfoot: "Don't show up the CEO" | Proudfoot 2023 | Confirmed. |
| Da Silva & Jensen: soothsayer, superior and precarious | Da Silva & Jensen 2022 (arXiv) | Confirmed in the abstract. |

Not checked by me: Copilot prompt wording and Squash counts (web pages, verified in results-5 on 2026-10-18; re-check Copilot near submission as the candidate-papers file already says); Wang et al. 2024 numbers (arXiv abstract only in results-5).
