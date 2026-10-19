# Review C: selection for the Project Literature Review

Reviewer C, written 2026-10-18. Judged only from the files listed in `_reviewer-brief.md` (assignment, abstract v7, agent briefs, known papers, results 1 to 6, the 75-entry related-papers file). No web search. Where I add outside knowledge it is marked "(outside the files)".

---

## 1. Proposed thematic subsections

**A. Boards oversee cyber risk without the expertise to judge the CISO, and now have tools that can check slides on the spot.**
Claim: Directors who lack cyber expertise depend on the CISO, rarely question for fear of looking like novices, and judge the CISO's competence by how well the CISO explains; regulation assumes this will stay so, while consumer and meeting products already let a listener point a camera at a slide or ask an AI for "holes" in a colleague's argument (Aamir et al., 2026; Lowry et al., 2026; Opdenbusch et al., 2025; Proudfoot et al., 2023; SEC, 2023; Microsoft, n.d.).

**B. People follow AI flags, including wrong ones, but HCI has measured belief in content, not trust in the person flagged.**
Claim: AI credibility indicators and wearable or live checkers pull non-experts toward the AI's verdict whether it is right or wrong, most of all when the AI contradicts a human expert, and VLMs that read slides without context are often wrong; yet these studies measure belief in claims, and observers can resist an AI that accuses a person (Lu et al., 2025; Gupta et al., 2026; Horne & Nevo, 2025; Bendeck & Stasko, 2025; Loconte et al., 2026).

**C. A check, a correction or an AI disagreement lowers how observers judge the professional checked, though not always.**
Claim: Being checked lowers trust even when nothing is found, and a correction or an AI that disagrees lowers lay ratings of a doctor, advisor, candidate or student; but several studies find corrections change beliefs without moving evaluation of the speaker (Strickland, 1958; Chen et al., 2026; Kim et al., 2026; Madanay et al., 2025; Prike et al., 2025; Du et al., 2026; Reif et al., 2025; Chae & Tewksbury, 2026).

**D. Trust in an expert has three dimensions, and the expert's answer can repair it, partly.**
Claim: Trust rests on perceived ability, benevolence and integrity, which laypeople judge separately in experts; a denial or account can restore trust when the accused is innocent, and exoneration brings a person back to about where they started, but corrected accusations leave an echo and unsupported denials sometimes work better than expected (Mayer et al., 1995; Hendriks et al., 2015; Kim et al., 2004; Ferrin et al., 2007; Sharma et al., 2023; Greitemeyer & Sagioglou, 2015; Mickelberg et al., 2025; Thorson, 2016; Schiff et al., 2025; Pezzo et al., 2022).

**E. A checker changes what the holder does: it lowers the threshold for challenging and shifts the burden of proof.**
Claim: People who consult an algorithm accuse others far more often, and boards that have independent reviews challenge management more confidently, so a checker should raise questioning and make the expert's answer carry more weight (von Schenk et al., 2024; Aamir et al., 2026; Opdenbusch et al., 2025; Chen et al., 2026).

Subsection E is short and could be folded into A and C if the word count is tight; I keep it separate because H4 needs its own evidence and the files show that evidence is thin.

---

## 2. Selected papers (30)

Priorities: 17 must cite, 12 should cite, 1 optional.

### A. Boards, CISOs and products

- **Aamir et al. (2026)**, CHI. Role: anchor, context. **Must.** The only SIGCHI paper on how directors question cyber-risk reports; already carries the research question in the abstract.
- **Lowry et al. (2026)**, Management Science. Role: context, gap. **Must.** The strongest evidence that non-expert directors' oversight is "mostly symbolic" and that they do not see it; it explains why a checker would appeal and why its flags would be hard to judge.
- **Opdenbusch et al. (2025)**, NDSS. Role: context for H4. **Must.** Board members avoid asking the CISO questions because they "fear being exposed as IT novices"; this is the behaviour the checker is expected to change.
- **Proudfoot et al. (2023)**, MIS Quarterly Executive. Role: supporting finding for H2 and H3. **Should.** Interviewees say boards judge CISO competence by how well the CISO communicates, which is the mechanism by which a reason would restore trust and a bare denial would not; stronger evidence base than Hooper and McKissack (2016).
- **U.S. SEC (2023)**, final rule 33-11216. Role: industry or trade source (regulation). **Should.** The regulator dropped the board cyber-expertise disclosure because directors "often effectively oversee ... without specific subject matter expertise"; this fixes the non-expert overseer as the norm. Better than the vendor surveys (Proofpoint, Gartner) because it is primary and not self-report.
- **Microsoft (n.d.), Copilot in Teams meetings**, with Meta (2024) Ray-Ban Meta as a second example. Role: product. **Should.** Microsoft's own suggested prompt "Where are some holes in [a meeting participant's] argument?" is a private, listener-held critique of a colleague in a live meeting, the closest product to our checker; the camera assistants show the point-and-ask interaction is on the market.

### B. AI checkers and flags

- **Lu et al. (2025)**, CHI. Role: supporting finding for H1 and H4. **Must.** An AI credibility indicator moved people more when it disagreed with a human expert than when it agreed; this is the AI-versus-expert case in a SIGCHI venue. Chosen over Lu et al. (2022) to avoid two papers on the same point; cite 2022 in passing if space allows.
- **Gupta et al. (2026)**, CHI. Role: closest system precedent, gap. **Must.** A real-time wearable fact-checker of spoken claims with deliberate errors; it measured trust in the AI's information, not in the speaker, which is our gap. Chosen over Danry et al. (2020) because it is recent, CHI, and live.
- **Horne & Nevo (2025)**, Scientific Reports. Role: supporting finding for H1 (wrong flags). **Should.** People kept relying on warning labels when the labels were wrong and lowered trust in correct news; the direct precedent for flags that are all wrong. Chosen over Freeze et al. (2021) and DeVerna et al. (2024), which make the same point with weaker access to the numbers or a content-only outcome.
- **Loconte et al. (2026)**, Computers in Human Behavior. Role: counter-result to H1. **Should.** The more confident an AI was that a person was lying, the more participants departed from it; observers resist confident AI accusations against people, which cuts against a flag lowering trust.
- **Bendeck & Stasko (2025)**, IEEE TVCG. Role: context for the wrong-flag premise. **Optional.** GPT-4V scored near the 16th percentile of humans on a visualization literacy test and read values off charts correctly 13% of the time without the data, while knowing textbook design rules; one citation is enough to ground "right by the textbook, wrong for this company". Chosen over CharXiv and HallusionBench because it is an HCI venue and closest to slides.

### C. What a check does to the person checked

- **Strickland (1958)**, Journal of Personality. Role: anchor (founding psychology source). **Must.** Supervisors trusted the often-checked worker less even when every check found nothing; the origin of the claim that checking itself costs trust and the reason for the "No issues found" slides.
- **Chen et al. (2026)**, IJHCS. Role: closest design. **Must.** An AI second opinion that disagreed with a doctor lowered patients' trust in the doctor; the doctor never answered, which is what our study adds.
- **Kim et al. (2026)**, AIES. Role: closest design, measure, gap. **Must.** An AI that monitored and flagged a human advisor's errors in front of non-expert students; advisor errors lowered METI expertise, integrity and benevolence, and how the AI was involved did not change that loss. The highest-scored new entry and the paper that most narrows our novelty claim.
- **Madanay et al. (2025)**, JMIR. Role: supporting finding for H1. **Must.** With 1,606 adults, a radiologist who recommended less testing than the AI was agreed with less and rated a worse provider; a radiologist who recommended more testing paid no penalty. Shows the cost of AI disagreement is asymmetric. Chosen over Bernstein et al. (2025), which measures liability with an AI that was right.
- **Prike et al. (2025)**, Political Psychology. Role: design source, supporting finding for H1. **Must.** Fact-checks against an unfamiliar candidate lowered ratings of his competence and morality; we borrow the profile and the 0-to-10 rating. Kept from the abstract.
- **Du et al. (2026)**, Frontiers in Psychology. Role: supporting finding for H1. **Should.** 214 teachers graded an identical paper lower when a fictitious AI detector reported 87% AI text; the evaluator with power holds a wrong flag, as our board member does. The nearest analogue outside medicine.
- **Reif et al. (2025)**, PNAS. Role: design element. **Should.** People who get help from AI are judged lazier and less competent; reason the CISO never mentions AI. Kept from the abstract. Schilke and Reimann (2025, 13 experiments, trust in the actor drops after AI disclosure) is the stronger trust-specific alternative if one is swapped.
- **Chae & Tewksbury (2026)**, New Media & Society. Role: counter-result to H1. **Must.** Across three experiments (N = 3,978), neither AI nor human fact-checks changed feelings toward or perceived competence of the politicians checked; the clearest case that a check can move belief but not the person's standing. Chosen over Swire-Thompson et al. (2020) and Nyhan et al. (2020) because it includes an AI checker and measures competence.

### D. Trust in experts and its repair

- **Mayer et al. (1995)**, Academy of Management Review. Role: definition. **Must.** Founding definition of trust and the ability, benevolence and integrity dimensions the study will ask about.
- **Hendriks et al. (2015)**, PLOS ONE. Role: measure. **Must.** METI: laypeople judge experts on expertise, integrity and benevolence, and all three move when source information is manipulated; it is the scale Kim et al. (2026) used, so our items can follow it.
- **Kim et al. (2004)**, Journal of Applied Psychology. Role: anchor for H2 and H3, counter-result to H3. **Must.** Denial repaired trust better than apology when later evidence showed innocence; our CISO is innocent by design, so this predicts a denial may help rather than hurt.
- **Ferrin et al. (2007)**, Journal of Applied Psychology. Role: supporting finding for H3. **Should.** Reticence did worse than the better of apology or denial; a denial with no reason shares reticence's lack of information and is the better support for H3 than Kim et al. (2004).
- **Sharma et al. (2023)**, Journal of Management. Role: definition and measure for H2. **Should.** Defines repair as any rise above the post-violation level and complete repair as a return to the pre-violation level; gives H2 the vocabulary of partial, complete and over-repair.
- **Greitemeyer & Sagioglou (2015)**, PLOS ONE. Role: supporting finding for H2. **Must.** A cleared researcher rose back to the level of one never accused but not reliably above it; the main evidence on how far a clearing goes. Kept from the abstract.
- **Mickelberg et al. (2025)**, PLOS ONE. Role: supporting finding for H2, method caution. **Should.** A retraction removed an allegation's effect on likeability in two of three experiments, and repeated ratings may focus attention on the latest information. Kept from the abstract.
- **Thorson (2016)**, Political Communication. Role: counter-result to H2. **Must.** Negative information kept lowering attitudes toward a candidate after a correction was accepted, even an immediate one; the main reason to expect H2's "above the start" to fail.
- **Schiff et al. (2025)**, American Political Science Review. Role: counter-result to H3. **Should.** Over 15,000 adults: politicians who falsely claimed a text story was misinformation gained support, more than by silence or apology; an unsupported denial can work.
- **Pezzo et al. (2022)**, Medical Decision Making. Role: supporting finding for H2 and H3. **Should.** A physician's stated reason for setting a diagnostic aid aside changed competence ratings: "I don't trust it" lowered them, "I already knew what it would say" raised them; the closest professional analogue to reason versus bare denial.

### E. The checker changes the holder's behaviour

- **von Schenk et al. (2024)**, iScience. Role: supporting finding for H4. **Must.** Unaided, about one in five accused; among those who chose a lie detector that said "lie", almost 85% accused. Kept from the abstract.
- Aamir et al. (2026), Opdenbusch et al. (2025) and Chen et al. (2026) also serve this subsection (listed above).

**Considered and left out**, with the reason: Kempt and Nagel (2022) name the burden-of-proof mechanism but are conceptual and Chen et al. already carry the point; Dobbs et al. (2026) (a doctor loses more credibility than a senator after correction) is strong but overlaps with Kim et al. (2026); Navajas Fernández et al. (2026, CHI) is the best AI-mediated communication entry if a marker wants that strand named, and could replace Reif et al.; Hendriks et al. (2016) is a close analogue (external criticism lowered expertise, self-disclosure kept integrity) but rests on secondary summaries; Hou and Jung (2021) on framing; Da Silva and Jensen (2022, CSCW) on the CISO as interpreter; Klingbeil et al. (2024) and Dratsch et al. (2023) on automation bias; Wegner et al. (1981) with Letourneau and Gawronski (2024) on innuendo, which would sharpen H1 if space allows.

---

## 3. Must-read list (full text before writing)

1. **Chen et al. (2026).** The files conflict on the central result. The abstract v7 and the earlier list say *disagreement* made the doctor seem lazier; Agent 4's press-coverage note says *agreement* made the doctor look lazier and disagreement raised perceived uncertainty and lowered perceived autonomy. Check which condition produced "lazy", N = 135, and that the doctor was a chatbot role-playing a human.
2. **Kim et al. (2026), AIES.** Check whether the AI-monitoring condition without advisor errors changed trust at all (it bears on our "No issues found" slides and on whether the checker's presence alone moves trust), whether the AI's corrections were always correct, and the exact effect sizes per METI dimension. This paper decides how narrow our novelty claim must be.
3. **Kim et al. (2004).** Listed as known but "not yet read" before this search. Check that the denial was a bare denial or included reasons, and the size of the denial advantage when innocence was shown; this decides H3's direction.
4. **Ferrin et al. (2007).** Abstract only. Check how reticence was worded and whether competence-violation results hold in both studies.
5. **Gupta et al. (2026).** The note that trust was measured in the AI's information, not the speaker, comes from a search snippet of the arXiv HTML; the gap statement depends on it.
6. **Chae & Tewksbury (2026).** The competence null comes from an abstract on a university news page; check the competence measure and whether targets were well-known politicians (which would limit it as a counter-result).
7. **Lu et al. (2025).** The "larger influence when the AI disagrees with an expert" reading is from the earlier list; check that "expert" is a labelled human expert opinion and what outcome was measured.
8. **Strickland (1958).** The earlier list says findings were taken from secondary sources because the original is paywalled; the abstract v7 gives specific numbers (40 students, 9 of 10 versus 2 checks), so confirm they come from the original.
9. **Madanay et al. (2025).** Check which direction of discrepancy maps onto our flags: our flags say the CISO's slide is wrong, not that more caution is needed, so the asymmetry may not transfer.
10. **Schiff et al. (2025).** Check that the "fake news" claims were unsupported and that the effect was against text, not video; our CISO's denial is truthful, theirs false.
11. **Opdenbusch et al. (2025)** and **Lowry et al. (2026).** Abstract-level only; confirm the "IT novices" quote and the "mostly symbolic" wording before quoting.
12. **Microsoft (n.d.).** The suggested prompt about "holes in [a meeting participant's] argument" should be re-checked on the live support page close to submission, since product pages change (outside the files: such pages are revised often).

---

## 4. Gap and differentiation

Across all six result files, the checker literature in HCI measures belief in content or reliance on the AI, and the trust literature measures trust in a person after a human accusation or correction; the two meet in only a handful of studies. The closest is Kim et al. (2026), where an AI flagged a human advisor's errors in front of non-expert students and METI trust in the advisor was measured, but the errors were real, the AI sat in the advisor's own workflow, and the advisor never answered. Chen et al. (2026) had an AI disagree with a doctor in front of a patient, but the doctor brought in the AI himself and stayed silent, and trust was measured once. Madanay et al. (2025) and Du et al. (2026) show lay or powerful evaluators marking down a professional or a student after an AI disagreement or a wrong AI report, but without any reply from the person judged and with ratings or grades rather than trust tracked over time. Gupta et al. (2026) put a live checker in the listener's hand but measured belief in claims, not trust in the speaker. No paper in the files closes the gap, but Kim et al. (2026) forces a narrower claim: we cannot say we are the first to measure trust in a human expert whom an AI checks. The defensible claim is that this is the first study in which the non-expert who holds power also holds the checker, every flag is wrong for the context, the expert answers with a reason or a bare denial, and trust is rated before and after both the flag and the answer, against a group that hears the same answers without a checker.

---

## 5. Hypothesis check

**H1 (a flag lowers trust).** Supported by Strickland (1958), Chen et al. (2026), Kim et al. (2026), Madanay et al. (2025), Prike et al. (2025), Du et al. (2026), Lu et al. (2025) and Horne and Nevo (2025). Cut against by Chae and Tewksbury (2026), where checks did not move perceived competence, and Loconte et al. (2026), where observers resisted confident AI accusations; the files suggest stating which dimension is expected to drop (expertise, per Hendriks et al. 2015 and the competence framing of our flags) rather than trust only overall.

**H2 (a reason restores trust, possibly above the start).** Supported for recovery by Greitemeyer and Sagioglou (2015), Mickelberg et al. (2025), Proudfoot et al. (2023) and Pezzo et al. (2022). Cut against for "above the start" by Thorson (2016) and by Greitemeyer and Sagioglou's own result (back to level, not reliably above); reword H2 in Sharma et al.'s (2023) terms as repair toward the starting rating, with a rise above it as an exploratory question.

**H3 (a denial without a reason restores less or lowers further).** Supported by Ferrin et al. (2007) and Pezzo et al. (2022). Cut against by Kim et al. (2004), where denial worked when the accused was innocent, and by Schiff et al. (2025), where an unsupported denial raised support; the files point to rewording H3 as "a denial without a reason restores less trust than a reasoned answer" and dropping "lower than after the flag" as the main prediction.

**H4 (checker participants question more; the same answer moves trust further).** H4a is supported by von Schenk et al. (2024), Opdenbusch et al. (2025) and Aamir et al. (2026); note that the comparison is between choosing to ask the CISO and saying one has questions, which are not the same act. H4b (amplification) has no direct support in the selected papers beyond the burden-of-proof argument in Chen et al. (2026), and Loconte et al. (2026) and Chae and Tewksbury (2026) suggest the flag may weigh less than expected, so H4b should be framed as exploratory.

---

## Selection table

| Short citation | Subsection | Role | Priority | New or known |
|---|---|---|---|---|
| Aamir et al., 2026 | A, E | Anchor, context | Must | Known (abstract) |
| Lowry et al., 2026 | A | Context, gap | Must | New |
| Opdenbusch et al., 2025 | A, E | Context, H4 finding | Must | New |
| Proudfoot et al., 2023 | A | Supporting finding (H2, H3) | Should | New |
| U.S. SEC, 2023 | A | Industry or trade (regulation) | Should | New |
| Microsoft, n.d. (Copilot in Teams), with Meta, 2024 | A | Product | Should | New |
| Lu et al., 2025 | B | Supporting finding (H1, H4) | Must | Known |
| Gupta et al., 2026 | B | Design precedent, gap | Must | Known |
| Horne & Nevo, 2025 | B | Supporting finding (H1, wrong flags) | Should | New |
| Loconte et al., 2026 | B | Counter-result (H1) | Should | New |
| Bendeck & Stasko, 2025 | B | Context (VLM reliability) | Optional | New |
| Strickland, 1958 | C | Anchor | Must | Known (abstract) |
| Chen et al., 2026 | C, E | Closest design | Must | Known (abstract) |
| Kim et al., 2026 (AIES) | C | Closest design, measure, gap | Must | New |
| Madanay et al., 2025 | C | Supporting finding (H1) | Must | New |
| Prike et al., 2025 | C | Design source, supporting finding (H1) | Must | Known (abstract) |
| Du et al., 2026 | C | Supporting finding (H1) | Should | New |
| Reif et al., 2025 | C | Design element | Should | Known (abstract) |
| Chae & Tewksbury, 2026 | C | Counter-result (H1) | Must | Known |
| Mayer et al., 1995 | D | Definition | Must | New |
| Hendriks et al., 2015 | D | Measure | Must | New |
| Kim et al., 2004 | D | Anchor, counter-result (H3) | Must | Known |
| Ferrin et al., 2007 | D | Supporting finding (H3) | Should | Known |
| Sharma et al., 2023 | D | Definition, measure (H2) | Should | New |
| Greitemeyer & Sagioglou, 2015 | D | Supporting finding (H2) | Must | Known (abstract) |
| Mickelberg et al., 2025 | D | Supporting finding (H2), method caution | Should | Known (abstract) |
| Thorson, 2016 | D | Counter-result (H2) | Must | New |
| Schiff et al., 2025 | D | Counter-result (H3) | Should | New |
| Pezzo et al., 2022 | D | Supporting finding (H2, H3) | Should | New |
| von Schenk et al., 2024 | E | Supporting finding (H4) | Must | Known (abstract) |
