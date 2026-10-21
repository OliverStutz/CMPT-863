# Design review C: structure of the literature review (proposal v1)

Reviewer C, 2026-10-19. Worked alone from `_design-brief.md`; did not open design-review-A.md or design-review-B.md.

## A. Answers

### Q1. Does the proposal meet every requirement in Instructions.md?

**With changes.**

| Requirement | Verdict | Why |
|---|---|---|
| Expand the abstract's related work | met | All 14 abstract sources are carried over and seven read papers are added. |
| Situate the project within prior work | met | Sections 1 to 4 each end on "how we differ", and section 5 states the gap element by element. |
| Scholarly research | met | About 20 peer-reviewed sources across HCI, organisational psychology, medicine and security. |
| Industry, trade, products where appropriate | met, with one weak item | SEC 2023, camera assistants, Copilot in Teams and Squash do real work. NACD and ISA 2023 is cited from a press release and a blog summary only (results-5), so it is the weakest item. Aamir et al. (2026, "Governance Toolkits and Dashboards") lists existing board tools (UK Cyber Security Toolkit for Boards, ASIC key questions, commercial dashboards). These would be a better-grounded product strand for boards than NACD. |
| At least 15 sources | met | 29 without the optional items. The risk is too many, not too few (see Q6 and Q8). |
| Thematic subsections | met | Five themes plus an opening. |
| Synthesis, not one-by-one summaries | at risk | Section 4 bundles about 10 sources into 300 words under three sub-bundles. Section 2 is a list of products followed by two benchmarks. As planned, both will read as catalogues unless each sub-bundle is reduced to one claim with its sources in parentheses. |
| What exists / what we know / unresolved / how we differ | met in the plan, but redundant | Every section repeats "how we differ", and section 5 repeats all of them. That costs about 120 words. Cut each section's "differ" to one sentence that points forward, and let section 5 carry the full argument. |
| Gap and novelty clear by the end | met | Section 5 plus rule 22 wording. |
| Under 1,500 words | met | 1,290 planned. The density per source is the problem (Q8). |
| Citation format as in the abstract (APA) | met | APA 7 as in Abstract v8. Products and the SEC release need APA web and report formats; results-5 already has them. |

Two accuracy fixes the plan should carry into drafting. Neither changes the structure.

- Section 4 credits the burden-of-proof point to Chen et al. (2026). Chen et al. cite it from Kempt and Nagel (2022): "a second opinion from an AI assistant can shift the burden of proof onto doctors (Kempt & Nagel, 2022)" (Chen full text, section 5.1). Cite Kempt and Nagel (2022) as the source (results-4, open access) and Chen et al. as the empirical case.
- Section 2 says Squash produced "more wrong matches than right ones" and that "live checkers produce false flags in practice". The source says 14 relevant against 20 non-relevant fact-checks (Adair 2020). That means irrelevant matches, not false verdicts. Word it as "more irrelevant than relevant pop-ups". Bendeck and Stasko (2025) carry the "wrong" claim better: 13 percent on value lookup without the data, and the 16th percentile overall (checked in the full text).

### Q2. Is each bundle coherent?

| Section | Coherent | Papers that do not fit |
|---|---|---|
| 1 Boards oversee what they cannot judge | **yes** | None. Aamir et al. (2026) fits here as field evidence that independent assurance makes directors challenge more. Note that it is also the H4a anchor. |
| 2 Checkers are available and often wrong | **yes** | None. Wang et al. (2024) and Bendeck and Stasko (2025) answer the same sub-question (see Q6). |
| 3 People follow a wrong flag, and the person flagged pays | **yes**, with one wording problem | All papers fit. But Kim et al. (2026) does not support the section's "what we know" line, "an AI's contradiction costs the contradicted expert". In Kim et al. the advisor's error moved trust, and how the AI was involved changed nothing. The cost came from the exposed error, not from the AI. Madanay et al. (2025) and Chen et al. (2026) show a cost from disagreement alone. State Kim et al. as the boundary ("AI involvement as such did not cost the advisor; the error did"), which is also the null against H4b. |
| 4 Trust in an expert, and how it is repaired | **no** | von Schenk et al. (2024) is about whether the holder of a fallible lie detector accuses. It concerns the observer following an AI accusation (H1 mechanism, H4a act), not repair, and the accused could not answer. Reif et al. (2025) is a disclosed-AI-use penalty and a design justification (the CISO never mentions AI), not repair. The third sub-bundle ("burden of proof and the AI penalty") is a holding pen for leftovers. Only the burden-of-proof point belongs here, because it is the bridge from flag to answer. |
| 5 Gap and this study | **yes** | None. Kim et al. (2026), Kim et al. (2004) and Gupta et al. (2026) serve the gap statement. |

### Q3. Moves

1. **von Schenk et al. (2024): from section 4 to section 3**, because it shows that an observer given a fallible AI accuser follows it (85 percent accused when it said "lie") against a person who cannot answer. That is section 3's question. It also pairs with Loconte et al. (2026), its stated counter-result and already in section 3: opaque tool, people follow; accuracy and confidence on screen, people resist. Its rise in accusations then carries H4a's "the tool changes what the holder does" next to Du et al.'s (2026) rise in intention to intervene.
2. **Reif et al. (2025): from section 4 to section 3**, because it belongs to the cluster "AI near a professional changes how the professional is judged" (Reif: using AI costs; Madanay and Chen: AI disagreement costs; Kim 2026: AI involvement as such does not). One clause there also justifies why the CISO never mentions AI. Left in section 4 it does no repair work.

No other moves. Strickland (1958) stays in section 3 as the root claim that checking alone lowers trust, which also grounds the "No issues found" slides. Mayer et al. (1995) and Hendriks et al. (2015) stay at the head of section 4, compressed to about 40 words.

Two additions (not moves) follow from Q7: Kempt and Nagel (2022) to section 4, and Pezzo et al. (2022) to section 4.

### Q4. Alternative bundlings

- **Alt A: reject.** Instructions.md asks explicitly for existing products and technologies. A visible products-and-reliability section is where a grader sees that requirement met. Shrink it to about 170 words instead of merging it.
- **Alt B: reject.** With the two moves in Q3, H4a is carried by section 1 (silence from fear of exposure, Aamir's independent assurance) and section 3 (von Schenk, Du on intervening). H4b is exploratory and has almost no literature, so a separate section would cost about 150 words mostly in double citations.
- **Alt C: reject.** Instructions.md wants readers to understand the gap "by the end", so the element-by-element gap belongs after the evidence. Opening with it previews claims the reader cannot yet judge. The redundancy Alt C targets is better fixed by cutting each section's "how we differ" to one sentence.

### Q5. Optional sources

| Source | Verdict | Section |
|---|---|---|
| Da Silva and Jensen (2022), CSCW | only if words allow | 1. One clause: the CISO's precarious interpreter role motivates self-serving framing ("cyber sophistry", checked in the full text). Its use overlaps Lowry et al.'s filtering and Opdenbusch et al.'s "dark art" quote, but it is a SIGCHI venue in a strand that otherwise has none. |
| Shaw, Wild and Colquitt (2003) | only if words allow | 4. Explanations versus none is the H3 contrast, and its "justifications help less than excuses" qualifies H2, because our reasoned answers are justifications. Unread and not yet fetched, so read it before citing. |
| Sharma, Schoorman and Ballinger (2023) | include | 4. Ten words give H2 its vocabulary: partial repair, complete repair (back to the start) and over-repair (above it, our exploratory claim). Must be read first. |
| Pezzo et al. (2022) | include | 4. It is the only evidence in the pool that an expert's stated reason about an AI changes how observers rate the expert ("I knew what it would say" raised competence ratings; "I don't trust it" lowered them). That is the nearest machine-side analogue to the CISO's answer and partly fills section 4's "unresolved: repair after a machine accusation". Open access. |
| Hancock, Naaman and Levy (2020) | only if words allow | 2 or 5. One clause positions the checker as receiver-side AI-mediated communication, the established HCI frame. It is useful for a CHI-style reader but does no evidential work. |

### Q6. Surplus sources

- **Wang et al. (2024), CharXiv.** It duplicates Bendeck and Stasko (2025) on "VLMs misread charts". Bendeck and Stasko is the better fit: it was checked in the full text, and it shows the exact split the flags rely on (knows design guidelines, cannot retrieve values without the data). Keep one benchmark; Wang only if a second number is wanted.
- **NACD and ISA (2023).** Unread, known from a press release and a blog. Its one job (the norm is that directors ask scripted questions) is already done by Lowry et al. (2026), whose directors ask questions "off a professional body's list", and by Opdenbusch et al. (2025). Replace it with Aamir et al.'s list of board toolkits, or drop it.
- **Mickelberg et al. (2025), borderline.** Its result overlaps Greitemeyer and Sagioglou (2015). Its unique point is a measurement caution (rating at every step draws attention to the latest information), which is method more than literature. Keep it as half a sentence or not at all.
- Reif et al. (2025) is surplus if left in section 4. It earns its place only after the move in Q3.

### Q7. Missing themes

1. **Why the holder's power should matter for trust.** "The board member holds the power" is one of the novelty elements (E2 in the novelty tally). Section 1 supports it only with interview descriptions of formal and informational power. No source predicts what power does to trust. Schilke, Reimann and Cook (2015, PNAS, results-3, open access) do: more powerful parties trusted their exchange partner less, through lower perceived benevolence. One sentence in section 1 or 5 turns E2 from a setting detail into a reason. It would also connect to Mayer et al.'s benevolence item.
2. **An expert answering or overriding a machine.** Section 3 ends with experts who stay silent, and section 4 covers human accusers. The bridge, an expert who sets an AI's verdict aside and gives a reason, has no source, so "how we differ" in section 4 overclaims a total absence. Pezzo and Pezzo (2006, results-4: a doctor who defied the aid lost its protection) and Pezzo et al. (2022, see Q5) fill it. Kempt and Nagel (2022) gives the normative claim that the burden of proof shifts.

Not missing: the "theories" in "approaches, technologies, or theories" are covered by Mayer et al., Hendriks et al., trust repair and the burden of proof. A judge-advisor frame (Bonaccio & Dalal 2006; Hou & Jung 2021) would be nice to have, not needed.

### Q8. Word budget

The plan is not realistic as written. Section 4 has about 10 sources in 300 words (about 30 words each, which forces a list), and the five "how we differ" blocks repeat section 5. The total is fine. The split should move words from sections 1 and 2 to 3 and 4, where the hypotheses live, and cut the repeated "differ" sentences.

| Section | Words |
|---|---|
| 0 Opening | 60 |
| 1 Boards oversee what they cannot judge | 240 |
| 2 Checkers exist and misread slides | 170 |
| 3 Flags are believed, followed, and cost the person flagged | 340 |
| 4 Trust in an expert and its repair | 360 |
| 5 Gap and this study | 170 |
| **Total** | **1,340** |

That leaves about 160 words of margin under 1,500 for Sharma et al. (2023) or Pezzo et al. (2022) and for the drafting overrun that always happens.

### Q9. Order, and author decisions

**Keep the order.** It follows the study's own sequence (setting, tool, flag, answer, gap), so each section hands the next its question.

Decisions:

- **Gupta gap-only.** The decision is respected here, but it costs something. Gupta et al. (2026) is the one HCI system that is a listener-held, real-time checker, so it is naturally "what exists" in section 2. Its interviews also report a human-to-human finding: participants called using it with a trusted doctor unnecessary, called it impolite in personal conversation, and expected "distrust between friends" (full text, Section 5). Citing that in section 2 measures nothing about trust in the AI, so it does not breach rule 10.
- **Politician exclusion.** Harmless if the review says in one sentence that results on politicians with settled reputations are not carried over (verified-view entry 12). Without that sentence, section 3 reads as one-sided, with Loconte as the only counter-result.
- **The other decisions** (trust in the AI not an outcome, Dobbs declined, Kim 2026 closest) do not harm the structure.

### Q10. Confidence

**Medium.** The structural calls (Q2 to Q4, Q8) are high. The optional and missing-source calls rest partly on search-entry summaries of unread papers (Shaw, Sharma, Pezzo, Schilke).

## B. Revised structure

**0. Opening (60).** The scenario in two sentences, then one sentence naming the five themes.

**1. Boards oversee what they cannot judge (240).**
- Bundle: SEC (2023); Lowry et al. (2026); Opdenbusch et al. (2025); Proudfoot et al. (2023); Aamir et al. (2026), including its board toolkits in place of NACD; Schilke et al. (2015) for power and trust. Optional: Da Silva and Jensen (2022).
- Argument: regulation asks non-experts to oversee an expert. They therefore judge the presenter's credibility, stay silent for fear of exposure, and challenge more only when an independent review backs them; power over the expert is itself a reason to trust less.

**2. Checkers exist and misread slides (170).**
- Bundle: camera assistants (Meta 2024; Apple 2024; Google 2025); Copilot in Teams; Squash (Adair 2020); Bendeck and Stasko (2025); Wang et al. (2024) optional. Optional: Hancock et al. (2020).
- Argument: any listener can now check a presenter live. The models know textbook rules but misread values without context, live checkers surface irrelevant matches, and none of them reports what a check does to the speaker.

**3. Flags are believed, followed, and cost the person flagged (340).**
- Bundle, followed: Horne and Nevo (2025); Lu et al. (2025); Du et al. (2026); von Schenk et al. (2024), moved; boundary from Loconte et al. (2026).
- Bundle, cost: Strickland (1958); Prike et al. (2025); Madanay et al. (2025); Chen et al. (2026); Reif et al. (2025), moved; Kim et al. (2026) as the boundary (the error costs, AI involvement as such does not).
- Argument: opaque flags are believed even when wrong, and they move the holder to act. Disagreement from an AI lowers the expert's standing, but in every study the expert is silent and trust is rated once. Sets up H1, H4a and the size of H1.

**4. Trust in an expert and its repair (360).**
- Bundle: Mayer et al. (1995); Hendriks et al. (2015); Kim et al. (2004); Ferrin et al. (2007); Greitemeyer and Sagioglou (2015); Mickelberg et al. (2025), short; Sharma et al. (2023); Kempt and Nagel (2022) with Chen et al. (2026) for the burden of proof; Pezzo et al. (2022) (and optionally Pezzo & Pezzo 2006) for an expert's reason about an AI. Optional: Shaw et al. (2003).
- Argument: trust in an expert has competence, integrity and benevolence parts. A reply repairs only when it supplies what the listener lacks. Recovery after an accusation usually reaches the start, not above it. An AI's disagreement places the burden of proof on the expert. The only evidence that an expert's reason about an AI matters comes from vignettes about outcomes, not live answers to a flag. Sets up H2, H3 and H4b.

**5. Gap and this study (170).**
- Bundle: Kim et al. (2026); Kim et al. (2004); Gupta et al. (2026); one sentence on politician studies not carried over.
- Argument: the combination is untested. It has wrong flags on a correct expert, a checker held by the non-expert with power, the expert's answer, trust rated at each step, and a boardroom setting. End with the contribution sentence.

Source count: about 30 with the additions (Kempt & Nagel 2022, Pezzo et al. 2022, Schilke et al. 2015, Sharma et al. 2023), and about 28 after dropping NACD and Wang.

## C. Evidence log

**Files read:**
- `../Instructions.md`
- `../../CMPT 863 Project Abstract v8.md`
- `_review-design-v1.md`
- `verified-view.md` (all entries and the 22 rules)
- `novelty-tally.md`
- `../Literature Review - Candidate papers.md`
- `_master-index.md`
- Entries pulled from the results files: Hancock 2020, Mieczkowski 2021, Navajas Fernández 2026, Danry 2020, Serra-Garcia and Gneezy 2025, Sundar and Kim 2019, de Brito Duarte 2025 (results-1); Hou and Jung 2021, Bonaccio and Dalal 2006 (results-2); Sharma 2023, Shaw 2003, Hendriks 2016, Schilke 2015 (results-3); Pezzo and Pezzo 2006, Pezzo 2022, Kempt and Nagel 2022 (results-4); Da Silva and Jensen 2022, SEC 2023, NACD 2023, camera assistants, Copilot, Adair 2020, Wang 2024 (results-5); Drolsbach 2024 (results-6); the Gupta note in results-1.

**Claims checked against full texts (pdftotext):**
- **Chen et al. (2026), section 5.1:** the burden-of-proof sentence cites Kempt and Nagel (2022). It is not Chen's own finding.
- **Kim et al. (2026):** the AI was always correct (Limitations), and the authors call for varying AI uncertainty or failure. The user "does not directly interact with the AI". This is consistent with the proposal, and it supports the Q2 caveat for section 3.
- **Bendeck and Stasko (2025):** the 16th percentile of humans, and 13 percent on non-visual lookup without the data. This supports the section 2 claim.
- **Gupta et al. (2026):** trust was measured in the AI-provided information. The interviews report the hospital quote ("I trust the doctor's profession"), "impolite in personal conversations" and "distrust between friends".
- **SEC (2023):** the quote "directors with broad-based skills in risk management and strategy" is confirmed.
- **Aamir et al. (2026):** "Where external assurance was routine, directors reported greater confidence in challenging management perspectives" is confirmed. Its "Governance Toolkits and Dashboards" subsection lists existing board tools.
- **Da Silva and Jensen (2022):** "cyber sophistry" is confirmed.
- **Adair (2020), from the results-5 entry:** 14 relevant against 20 non-relevant pop-ups. These are irrelevant matches, not false verdicts.
