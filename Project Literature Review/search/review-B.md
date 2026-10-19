# Review B: paper selection for the literature review

Reviewer B, 2026-10-18. Judged only from the files listed in `_reviewer-brief.md` (assignment, abstract v7, agent briefs, known-papers list, the six result files, the 75-entry related-papers file). No web search. Where I add outside knowledge it is marked "(outside knowledge)".

Selection: 30 papers, 17 must cite, 11 should cite, 2 optional. All 8 abstract v7 sources are kept. 13 are known candidates or abstract papers, 17 are new from this search.

---

## 1. Proposed thematic subsections

**S1. Trust in an expert, and the cost of being checked.** Laypeople judge an expert on expertise, integrity and benevolence, and checking a person can lower trust in that person even when the check finds nothing (Mayer et al., 1995; Hendriks et al., 2015; Strickland, 1958).

**S2. Boards as non-expert overseers of the CISO.** Board members cannot judge a CISO's technical claims, rarely question them for fear of looking like novices, and read the CISO's competence from how well the CISO explains, while regulation assumes boards will stay non-expert; this is the gap an on-the-spot checker promises to fill (Aamir et al., 2026; Lowry et al., 2026; Opdenbusch et al., 2025; Da Silva & Jensen, 2022; Proudfoot et al., 2023; SEC, 2023).

**S3. AI checkers in the listener's hand: products, reliability and reliance.** Listener-side checking is already on the market and in CHI prototypes, people move toward an AI verdict whether it is right or wrong, vision-language models misread charts and numbers, and holding an algorithm makes people accuse more readily (Microsoft, n.d.; Gupta et al., 2026; Lu et al., 2025; Bendeck & Stasko, 2025; von Schenk et al., 2024; Hancock et al., 2020).

**S4. When an AI or a check says the expert is wrong: what happens to the person.** Corrections and AI disagreement lower ratings of the person flagged, and trusted professionals lose more, but several studies find no loss in the person's standing or find observers resisting a confident AI accusation (Chen et al., 2026; Kim, Daneshi et al., 2026; Madanay et al., 2025; Prike et al., 2025; Dobbs et al., 2026; Reif et al., 2025; counter: Chae & Tewksbury, 2026; Loconte et al., 2026).

**S5. Answering the flag: accounts, denials and what remains.** A reasoned account repairs trust better than no account and exoneration brings a person back to about the starting level but not reliably above it, while a bare denial can help the truly innocent and the original charge leaves an echo (Kim et al., 2004; Ferrin et al., 2007; Greitemeyer & Sagioglou, 2015; Mickelberg et al., 2025; Shaw et al., 2003; counter: Thorson, 2016; Schiff et al., 2025).

The gap and the statement of how the study differs close S5 (see section 4). S3 could be split into "products and reliability" and "reliance and accusation" if the word count allows; I would not split it, because 1,500 words do not fit six subsections with real synthesis.

---

## 2. Selected papers

### S1. Trust in an expert, and the cost of being checked

- **Mayer et al. (1995)**. Anchor, definition. **Must cite.** It is the source of ability, benevolence and integrity, which the abstract's measures section names. Its definition of trust ("irrespective of the ability to monitor or control") also lets the review say that a checker gives the board member exactly that ability to monitor. Colquitt et al. (2007) and Rousseau et al. (1998) say the same thing in less useful ways.
- **Hendriks et al. (2015)**, METI. Measure. **Must cite.** This is the validated scale of laypeople's trust in an expert (expertise, integrity, benevolence). It fits our non-expert-judges-expert setting better than any management scale, and Kim, Daneshi et al. (2026) used it in the closest design.
- **Strickland (1958)**. Anchor, mechanism (H1, design of slides 1 and 5). **Must cite.** This is the founding result that checking alone lowers trust in the person checked, given by a supervisor who holds power over that person. It is in the abstract and justifies the "No issues found" slides.

### S2. Boards as non-expert overseers of the CISO

- **Aamir et al. (2026)**, CHI. Context, H4a. **Must cite.** This is the SIGCHI source on boardroom cyber governance. It finds directors pushing back more confidently where independent reviews are routine, which the abstract relies on. It is in the abstract.
- **Lowry et al. (2026)**, Management Science. Context, gap. **Must cite.** This is the strongest scholarly evidence that non-expert directors' cyber oversight is "mostly symbolic" and that they do not notice. That makes it the reason a checker appeals to them and the reason they cannot judge its flags.
- **Opdenbusch et al. (2025)**, NDSS. Context, H4a. **Must cite.** Board members do not ask questions because they "fear being exposed as IT novices", and they reduce cyber to budget decisions. This underpins both the budget scenario and the prediction that a checker raises questioning.
- **Da Silva & Jensen (2022)**, CSCW. Context. **Should cite.** This SIGCHI paper names the CISO's interpretive monopoly ("soothsayer") that a checker breaks, and it keeps the board theme anchored in an HCI venue.
- **Proudfoot et al. (2023)**, MIS Quarterly Executive. Context, H2 and H3. **Should cite.** Interviewees say boards judge CISO competence "heavily" by the CISO's ability to communicate. This is the practitioner premise for expecting a reasoned answer to restore trust and a bare denial not to. The full text was read.
- **U.S. Securities and Exchange Commission (2023)**. Industry or trade (regulation). **Should cite.** The SEC dropped the rule on board cyber expertise because directors "often effectively oversee ... without specific subject matter expertise". The regulator assumes the non-expert overseer, so this is the strongest trade or regulatory source in the files. The NACD handbook (not read in full) and the Proofpoint survey (a vendor self-report) are weaker.

### S3. AI checkers in the listener's hand: products, reliability and reliance

- **Microsoft (n.d.)**, Copilot in Teams meetings. Product. **Should cite.** Microsoft markets the prompt "Where are some holes in [a meeting participant's] argument?" in a private pane during a live meeting. That is our checker in product form, and it is closer to us than the camera "look and ask" assistants (Meta 2024, Apple 2024). Those can be named in the same sentence without a separate entry.
- **Gupta et al. (2026)**, CHI. Anchor, design precedent, gap. **Must cite.** This is the closest live, listener-worn checker of spoken claims. Users over-relied on it and were misled by wrong flags, but it measured belief in claims (or trust in the AI's information), not trust in the speaker.
- **Lu et al. (2025)**, CHI. Supporting finding, H1. **Should cite.** An AI credibility indicator swayed people more when it disagreed with a human expert than when it agreed, which is the condition our flags create. I chose it over its predecessor Lu et al. (2022), which is content only. Both cannot fit, but the review can mention the 2022 result ("whether right or wrong") in a clause if the 2025 full text confirms it.
- **Bendeck & Stasko (2025)**, IEEE TVCG. Context (VLM reliability). **Should cite.** GPT-4V scored at about the 16th percentile of humans on a visualization literacy test and read values off charts correctly only 13% of the time without the data, yet it knew textbook design rules. That is exactly the profile behind "right by the textbook, wrong for this company", and it comes from a visualization venue. CharXiv, HallusionBench and Quelle & Bovet (2024) are alternatives. I would add Quelle & Bovet only if the review needs the "LLM checkers lean towards calling true claims false without context" point.
- **von Schenk et al. (2024)**, iScience. Supporting finding, H4a. **Must cite.** People who chose to consult a two-in-three-accurate lie detector accused in about 85% of flagged cases, against about one in five without it. This is the clearest evidence that holding an algorithm changes how one treats a person. It is in the abstract.
- **Hancock et al. (2020)**, JCMC. Definition, gap framing. **Optional.** It defines AI-mediated communication as an agent acting for the sender. That lets the review say our checker is receiver-side mediation, a case the definition does not cover. Being conceptual, it is the first to drop if space is short.

### S4. When an AI or a check says the expert is wrong: what happens to the person

- **Chen et al. (2026)**, IJHCS. Anchor, closest design. **Must cite.** An AI second opinion that disagreed with a doctor lowered trust in that doctor. The doctor never answered and brought the AI in himself. It is in the abstract. Its direction of effects needs checking (section 3).
- **Kim, Daneshi, Draper & Sreedharan (2026)**, AIES. Anchor, closest design, measure. **Must cite.** A human advisor's errors lowered METI expertise, integrity and benevolence and global trust. Whether an AI was absent, invoked by the advisor or monitoring and flagging his errors did not change that penalty. Trust in the human is measured with three dimensions while an AI flags him in front of a non-expert, so this is the paper that narrows our novelty claim.
- **Madanay et al. (2025)**, JMIR. Supporting finding, H1. **Must cite.** 1,606 adults agreed less with, and rated as a worse provider, a radiologist who recommended less testing than the AI. A radiologist who recommended more testing paid nothing, so the penalty runs one way: toward the side that looks less cautious than the AI. Our flags play the cautious role.
- **Prike et al. (2025)**, Political Psychology. Supporting finding, H1, design source. **Must cite.** Fact-checks against an unfamiliar candidate lowered ratings of his competence and morality and cut his votes. It is the source of our profile and 0 to 10 rating, and it is in the abstract.
- **Dobbs et al. (2026)**, Scientific Reports. Supporting finding, H1. **Should cite.** Corrections lowered the rated credibility of both a doctor and a senator, and the doctor lost more. The loss was largest when the error was framed as intentional. This supports measuring integrity separately and predicts a large drop for a trusted expert such as a CISO. The full text was read.
- **Chae & Tewksbury (2026)**, New Media & Society. Counter-result to H1. **Must cite.** Across three experiments (N = 3,978), neither AI nor human fact-checks changed feelings toward, or the perceived competence of, the targeted politicians. This is the cleanest null for speaker evaluation, and it is AI-labelled. The likely boundary is well-known targets against our stranger CISO.
- **Loconte et al. (2026)**, CHB. Counter-result to H1 and H4a. **Should cite.** Observers moved away from an AI's verdict the more confidently it said a person was lying. The authors trace this to a default assumption that people tell the truth and to the social cost of accusing. Our flags are a softer accusation, but this predicts some resistance.
- **Reif et al. (2025)**, PNAS. Design element, contrast. **Should cite.** Workers seen using AI were rated lazier and less competent. That is why the CISO never mentions AI, and it frames our inversion: the evaluator, not the expert, uses the AI. It is in the abstract. Schilke & Reimann (2025) and Reis et al. (2025) make the same point and are not needed in addition.

### S5. Answering the flag: accounts, denials and what remains

- **Kim, Ferrin, Cooper & Dirks (2004)**. Anchor, H2 and H3, partial counter-result. **Must cite.** Apology repaired competence violations better and denial repaired integrity violations better. When later evidence showed innocence, denial worked better. Our CISO is innocent by design, so a bare denial may repair more than H3 assumes.
- **Ferrin et al. (2007)**. Supporting finding, H3. **Must cite.** Reticence did worse than the better of apology and denial in every case. A "no, that is fine" with no account shares the weakness of reticence, so this is the best support for H3 in the files.
- **Greitemeyer & Sagioglou (2015)**. Anchor, H2. **Must cite.** A cleared researcher rose back to the level of one never accused but not reliably above it. This is the benchmark H2's "above the start" has to beat. It is in the abstract.
- **Mickelberg et al. (2025)**. Supporting finding, method, H2. **Should cite.** A retraction removed the effect of an allegation on person impressions in two of three experiments, and the authors warn that repeated rating may draw attention to the latest information. That bears on our step-by-step ratings. It is in the abstract.
- **Shaw et al. (2003)**. Supporting finding and nuance, H2. **Optional.** Meta-analytically, explanations improve reactions, but excuses (outside causes) beat justifications (the act was right). Our reasoned answers are justifications, so this is a caution on how much H2 can expect. Bies & Shapiro (1987) is the older single-study alternative.
- **Thorson (2016)**. Counter-result to H2. **Must cite.** Negative information about a candidate kept lowering attitudes after an accepted, immediate correction ("belief echoes"). It is the main reason to expect trust not to return above, or even fully to, the start.
- **Schiff et al. (2025)**, APSR. Counter-result to H3. **Should cite.** Across 15,000 adults, politicians who falsely claimed a text scandal was "fake news" gained more support than those who stayed silent or apologised. An unsupported denial can therefore work, and the review should say why ours may differ: our audience has no partisan loyalty, and the CISO is telling the truth.

---

## 3. Must-read list (full text before writing)

1. **Chen et al. (2026).** The files disagree on the direction. Abstract v7 says *disagreement* made the doctor seem lazier. Agent 4's press-coverage note says *agreement* made the doctor look lazier, while disagreement raised uncertainty and lowered perceived *autonomy*. The 75-entry list says agreement raised credibility and disagreement raised uncertainty and laziness. Check which condition raised laziness, whether "autonomy" is a measure, N = 135, and that the doctor was a chatbot presented as human. The abstract's description of the closest design depends on this.
2. **Kim, Daneshi, Draper & Sreedharan (2026).** Check how the "AI monitors and flags" condition was shown (did participants see a flag before the correction?), whether the advisor ever responded, the cell sizes (157 across 6 cells, about 26 per cell), and whether any condition had the AI flag a *correct* advisor. If no condition had a wrong flag, our "wrong flag" contribution stands. If one did, the gap narrows further.
3. **Chae & Tewksbury (2026).** The null on competence comes from a departmental summary page (Agent 6), not the article. Check that competence was measured on the politicians, which politicians were used (well known or not), and the AI label wording.
4. **Gupta et al. (2026).** Agent 1 took "trust measured in the AI's information, not in the speaker" and the 87.5% accuracy from a search snippet of the arXiv HTML. Confirm, because the gap statement rests on it.
5. **Lu et al. (2025).** The claim "the indicator's influence was larger when it disagreed with a human expert" comes only from the earlier 75-entry list. Confirm the wording and whether the expert was a person shown to participants.
6. **Kim et al. (2004).** Check what the "denial" contained (any reason, or a bare denial), and how the competence violation was framed. Our H3 rewording depends on whether their denial is comparable to our denial without a reason.
7. **Ferrin et al. (2007).** Check the exact wording of the reticence response and whether the competence case had an innocence condition.
8. **Strickland (1958).** The 75-entry list says findings were summarised from secondary sources (Schweitzer et al., 2018; Kramer, 1999). The abstract gives exact numbers (40 male students, 9 of 10 versus 2 checks). Confirm against the original before the review repeats them.
9. **Aamir et al. (2026).** Confirm the exact finding that directors "pushed back more confidently where independent reviews were routine". The abstract leans on it for H4a.
10. **Mickelberg et al. (2025).** Confirm that "two of three experiments" and the note on repeated rating are the authors' words.
11. **Loconte et al. (2026).** Confirm whether participants judged statements or persons, and how "departing from the AI" was measured.
12. **Microsoft Copilot support page.** Re-access it and record the date. The "holes in [a meeting participant's] argument" prompt wording may have changed.

Not on the list because the files report full text or a detailed abstract: Madanay et al. (2025), Dobbs et al. (2026), Proudfoot et al. (2023), Bendeck & Stasko (2025), Greitemeyer & Sagioglou (2015).

---

## 4. Gap and differentiation

Across all six result files, no study gives the *listener* an AI checker that flags a *live expert* in a power-asymmetric professional setting, uses flags that are *wrong*, lets the expert *answer*, and rates trust in that expert *before and after* each step. The three closest designs each miss a different part of this.

- **Kim, Daneshi et al. (2026)** comes nearest on the outcome. A non-expert rated a human advisor on METI while an AI monitored and flagged him. The AI was right (it caught real errors), the advisor never answered, and the student had no power over the advisor.
- **Chen et al. (2026)** has an AI second opinion contradicting a doctor. The doctor himself brought the AI in, never replied, and trust was measured once at the end.
- **Gupta et al. (2026)** puts a real-time checker in the listener's hand. It measures belief in claims, not trust in the speaker.
- **Madanay et al. (2025)** and **Du et al. (2026)** (not selected) add that a disagreeing AI, or a wrong detector report held by an evaluator with power, lowers the evaluation of the person. In neither does the person answer.

No paper in the files closes the gap. Kim, Daneshi et al. (2026) does force a narrower claim. The review cannot say we are the first to measure trust in a human expert, on expertise, integrity and benevolence, while an AI flags that expert in front of a non-expert. The defensible claim is the combination: the checker is held by the more powerful, less expert party; the flags are wrong in context; the expert's answer (reason or bare denial) is the manipulation; and trust is tracked step by step against a group that hears the same answers without the checker. Kim, Daneshi et al. also found that the way the AI was involved did not change the trust penalty. That is a warning that the checker's presence alone (slides 1 and 5) may not move trust, and the review should say so.

---

## 5. Hypothesis check

**H1 (a flag lowers trust).**
- Supported by Chen et al. (2026), Madanay et al. (2025), Prike et al. (2025) and Dobbs et al. (2026). Gupta et al. (2026) and Lu et al. (2025) show that belief follows a wrong AI, and Strickland (1958) that checking itself costs trust.
- Cut against by Chae & Tewksbury (2026), where neither AI nor human checks changed the competence ratings of politicians, and by Loconte et al. (2026), where observers resisted confident AI accusations.
- Suggested rewording: predict the drop mainly in rated competence or expertise, since the flags attack expertise (Hendriks et al., 2015; Dobbs et al., 2026), and keep integrity and benevolence as secondary outcomes.

**H2 (a reasoned answer restores trust, possibly above the start).**
- Supported by Greitemeyer & Sagioglou (2015), Mickelberg et al. (2025) and Proudfoot et al. (2023).
- Cut against by Thorson (2016, belief echoes), by Greitemeyer & Sagioglou's own "not reliably above" result, and by Shaw et al. (2003), where justifications help less than excuses.
- Suggested rewording: "a reasoned answer raises trust above the post-flag rating and back toward the start", with a rise above the start treated as exploratory.

**H3 (a denial without a reason lowers trust further).**
- Supported by Ferrin et al. (2007, reticence worst) and Proudfoot et al. (2023).
- Cut against by Kim et al. (2004), where denial worked best when the accused was in fact innocent, and Schiff et al. (2025), where an unsupported denial raised support.
- The files point to the rewording Agent 3 gives: "a denial without a reason restores less trust than a reasoned answer". "Lower than after the flag" should stay as a secondary, directional test.

**H4 (checker participants question more; the same answer moves trust further).**
- H4a is supported by Opdenbusch et al. (2025, boards stay silent out of fear of looking like novices), Aamir et al. (2026) and von Schenk et al. (2024). Loconte et al. (2026) cuts against it, because the social cost of accusing may hold participants back.
- H4b has no direct support in the files. The burden-of-proof argument (Chen et al., 2026; Kempt & Nagel, 2022, not selected) is conceptual. Kim, Daneshi et al. (2026) found that the form of AI involvement did not change the trust penalty, which cuts against amplification.
- H4b should be stated as exploratory. The review should also note that H4a compares an action (questioning) with a self-report (saying one has questions), which the literature does not treat as the same measure.

---

## Selection table

| Short citation | Subsection | Role | Priority | New or known |
|---|---|---|---|---|
| Mayer et al., 1995 | S1 | Anchor, definition | Must cite | New |
| Hendriks et al., 2015 | S1 | Measure | Must cite | New |
| Strickland, 1958 | S1 | Anchor, mechanism | Must cite | Known (abstract v7) |
| Aamir et al., 2026 | S2 | Context, H4a | Must cite | Known (abstract v7) |
| Lowry et al., 2026 | S2 | Context, gap | Must cite | New |
| Opdenbusch et al., 2025 | S2 | Context, H4a | Must cite | New |
| Da Silva & Jensen, 2022 | S2 | Context | Should cite | New |
| Proudfoot et al., 2023 | S2 | Context, H2 and H3 | Should cite | New |
| U.S. SEC, 2023 | S2 | Industry or trade | Should cite | New |
| Microsoft, n.d. (Copilot in Teams) | S3 | Product | Should cite | New |
| Gupta et al., 2026 | S3 | Anchor, design precedent, gap | Must cite | Known |
| Lu et al., 2025 | S3 | Supporting finding | Should cite | Known |
| Bendeck & Stasko, 2025 | S3 | Context (VLM reliability) | Should cite | New |
| von Schenk et al., 2024 | S3 | Supporting finding, H4a | Must cite | Known (abstract v7) |
| Hancock et al., 2020 | S3 | Definition, gap | Optional | New |
| Chen et al., 2026 | S4 | Anchor, closest design | Must cite | Known (abstract v7) |
| Kim, Daneshi et al., 2026 | S4 | Anchor, closest design, measure | Must cite | New |
| Madanay et al., 2025 | S4 | Supporting finding | Must cite | New |
| Prike et al., 2025 | S4 | Supporting finding, design source | Must cite | Known (abstract v7) |
| Dobbs et al., 2026 | S4 | Supporting finding | Should cite | New |
| Chae & Tewksbury, 2026 | S4 | Counter-result | Must cite | Known |
| Loconte et al., 2026 | S4 | Counter-result | Should cite | New |
| Reif et al., 2025 | S4 | Design element, contrast | Should cite | Known (abstract v7) |
| Kim et al., 2004 | S5 | Anchor, partial counter-result | Must cite | Known |
| Ferrin et al., 2007 | S5 | Supporting finding (H3) | Must cite | Known |
| Greitemeyer & Sagioglou, 2015 | S5 | Anchor (H2) | Must cite | Known (abstract v7) |
| Mickelberg et al., 2025 | S5 | Supporting finding, method | Should cite | Known (abstract v7) |
| Shaw et al., 2003 | S5 | Supporting finding, nuance | Optional | New |
| Thorson, 2016 | S5 | Counter-result | Must cite | New |
| Schiff et al., 2025 | S5 | Counter-result | Should cite | New |
