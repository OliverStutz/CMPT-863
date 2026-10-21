# Fact check of four corrections (2026-10-19)

Independent check against primary sources. Reviewer files (novelty-review-*, design-review-*, the two tallies) were not used as evidence.

Sources read:
- Kempt & Nagel (2022), J Med Ethics 48(4), 222-229. The publisher PDF is behind a Cloudflare check, so the copy came from the Wayback Machine snapshot of https://jme.bmj.com/content/medethics/48/4/222.full.pdf (20250620). Saved outside the project at /tmp/claude-1000/kempt_dl/kempt.pdf.
- Chen, Sun, Liao & Sundar (2026), IJHCS 213, 103824: `references/Chen-2026-when-AI-disagrees.pdf` (already in the project), plus `references/reviews/chen.md`.
- Madanay, O'Donohue & Zikmund-Fisher (2025), JMIR: `Project Literature Review/references/Madanay-2025-patient-reactions-AI-clinician-discrepancies-JMIR-EuropePMC.txt`.
- Adair (2020), Duke Reporters' Lab page, fetched 2026-10-19.
- Kim, Daneshi, Draper & Sreedharan (2026), AIES 9(3), 1514-1525: `Project Literature Review/references/Kim-2026-asymmetric-trust-corrective-AI-AIES.pdf`.

## A. Verdicts

### 1. Burden of proof: CONFIRMED

The argument is Kempt and Nagel's. Their abstract (p. 222) says that when physicians receive an AI-DSS recommendation they cannot confirm or reject "without being able to sufficiently explain why. This shifts the burden of proof towards the physician-in-charge and raises problems regarding the attribution of moral responsibility." For any second opinion, including a human one (p. 224): "A disagreement between initial and second opinion, however, establishes the burden of proof as falling on the physician-in-charge". For the AI case (p. 226): "the burden of proof of the diagnostic process moves from an epistemic equilibrium ... towards the proponent of the initial opinion." Conclusion (p. 227): "It does, however, move the burden of proof towards the physician-in-charge, as they have to incorporate AI-DSS as evidence, but cannot reject the evidence even if they think it is wrong."

Chen et al. do not test the claim. They mention it twice, each time citing someone else:
- Section 2 (p. 4): "disputing or rejecting AI's recommendations places the burden of proof on doctors (Kempt et al., 2023)". That is Kempt, Heilinger & Nagel (2023), AI & Society, not the 2022 paper.
- Section 5.1, discussion (p. 11): "Given that a second opinion from an AI assistant can shift the burden of proof onto doctors (Kempt & Nagel, 2022), the resulting reduction of trust due to increased medical uncertainty could potentially be mitigated by providing explanations for the assistant's disagreement." `references/reviews/chen.md` line 51 quotes this sentence correctly.

So the corrected attribution is right: Kempt and Nagel make the argument and Chen et al. cite it. Wording such as "Kempt and Nagel (2022) argue ..., a point Chen et al. (2026) take up" (Abstract v8, line 61) is accurate.

### 2. Madanay "worst provider": CONFIRMED

Main text and Table 2 (line 64 of the txt file): "Participants were both least likely to recommend (mean 6.47, SE 0.13) and positively rate (mean 6.82, SE 0.11) the radiologist who undercalled AI, which was significantly different from the other conditions (P <.001). There were no significant differences among the radiologist-only, radiologist-AI agreement, and radiologist-overcalls-AI conditions in ... rating of the radiologist (mean 7.45, SE 0.10 vs mean 7.55, SE 0.10 vs mean 7.58, SE =0.10)." Table 2 row: "Rating of radiologist 7.45 (0.10) 7.55 (0.10) 7.58 (0.10) 6.82 (0.11)".

Scale (Table 2 note): "The rating of radiologist measured from 1=worst provider possible to 10=best provider possible."

The paper never says the radiologist was rated "the worst provider". The phrase appears only as the bottom anchor of the scale ("worst provider possible"), which is probably where the old wording came from. A mean of 6.82 on 1 to 10 is lowest of the four conditions and well above the "worst" anchor. "Rated lowest of the four conditions (6.82 against 7.45 to 7.58 on a 1-to-10 scale)" is correct.

### 3. Squash: PARTLY CONFIRMED

Confirmed: "Squash popped up 14 relevant fact-checks on the screen, up from just six last year." and "This month's State of the Union speech also had 20 non-relevant matches, which means Squash displayed fact-checks that weren't related to what the president said." So the 20 are irrelevant matches (fact-checks that did not relate to what was said), not false verdicts, and the 14 are relevant ones. The page never gives the figure 34; it is the sum 14 + 20, which is fair as long as it is described as 14 relevant plus 20 irrelevant.

Not confirmed as worded: "a human editor was added". The page describes a plan that was not yet in place: "a new interface that will bring humans into the live decision-making. Squash will recommend fact-checks to display, but an editor will make the final judgment." and "Guess, working with our student developer Matt O'Boyle, is building an interface for human editors to control which matches actually pop up on users' screens. The new interface would let them review the fact-check that Squash recommends". Correct wording: the lab decided to add a human editor and was building the interface for one.

### 4. Kim et al. (2026): CONFIRMED

Section 5.2, p. 1520: advisor performance affected every advisor-directed measure, e.g. perceived trustworthiness "F (1, 151) = 13.0, p < .001, η2 = .08", general trust "F (1, 151) = 21.14, p < .001, η2 = .12", reuse "F (1, 151) = 26.29, p < .001, η2 = .15", expertise "F (1, 151) = 21.89", integrity "8.80", benevolence "6.92", "with lower evaluations in error-present conditions".

p. 1521: "No significant main effects of AI intervention modality were observed across outcomes in either the full 3×2 analyses (all F (2, 151) ≤ 1.81, all p ≥ .168, η2 ≤ .023) or the restricted 2×2 analyses ... so H2 was not supported. No significant advisor performance by modality interactions were observed ... so H3 was not supported." Bayes factors favoured excluding modality effects (BFexcl = 3.66 to 6.74) and the interactions. Also: "AI assistant evaluations alone were not significantly affected by advisor performance (all p > .05)."

One design detail settles it. The modality factor includes a no-AI condition, where the advisor works unaided (p. 1515), and every error was corrected before the final schedule. Error-present advisors lost trust whether or not an AI was involved, and descriptively "performance-related declines were largest in advisor-only conditions" (p. 1515). The loss follows the advisor's error, not the AI's contradiction. The paper bounds the claim "an AI's contradiction costs the expert"; it does not support it.

## B. Occurrences

Line numbers are those at the time of the check. "Changed" lists the old and new text.

### 1. Burden of proof attributed to Chen

| File | Line | Wording | Changed? |
|---|---|---|---|
| CMPT 863 Project Abstract v7.md | 61 | "Chen et al. note that an AI's contrary opinion shifts the burden of proof onto the expert." | No (submitted version, report only) |
| CMPT 863 Project Abstract v6.md | 54 | same sentence | No (superseded version, outside the edit scope) |
| CMPT 863 Project Abstract v5.md | 28 | same sentence | No (superseded) |
| CMPT 863 Idea 2 - Project Abstract v4.md | 28 | same sentence | No (superseded) |
| CMPT 863 Idea 2 - Project Abstract v4.md | 114 | "H4b grounded in Chen's burden-of-proof point" | No (superseded) |
| CMPT 863 Idea 2 - Project Abstract v4.md | 218 | "Chen's note ... (their reading of Kempt & Nagel, 2022, not a finding)" | No (superseded; hedged, almost right) |
| Project Literature Review/search/_review-design-v1.md | 66 | "Chen et al. 2026's point that AI disagreement shifts the burden of proof onto the expert" | No (this is the proposal the design reviewers checked, superseded by v2; v2 line 42 is correct) |
| references/reviews/related-work-review-A.md | 96 | "Chen's burden-of-proof remark" | No (reviewer record) |
| references/reviews/related-work-review-C.md | 135, 137, 191 | "Chen's burden-of-proof point"; "Chen et al. note that an AI's contrary opinion shifts the burden of proof" | No (reviewer record; line 101 of the same file gives the Kempt & Nagel citation correctly) |
| Project Literature Review/search/review-C.md | 78, 113 | "Chen et al. already carry the point"; "the burden-of-proof argument in Chen et al. (2026)" | No (selection-reviewer record) |
| Project Literature Review/search/review-B.md | 125 | "(Chen et al., 2026; Kempt & Nagel, 2022, not selected)" | No (reviewer record; partly right) |

Already correct: Abstract v8 line 61; v8 change log item 11; Candidate papers line 37; results-4 lines 82-83; trust-topic-map line 196; _review-design-v2 line 42; references/reviews/chen.md line 51; v6-related-work-validation line 83 (flags it as IMPRECISE).

### 2. Madanay "worst provider"

| File | Line | Old | New | Changed? |
|---|---|---|---|---|
| Project Literature Review/search/reduced-view.md | 36 | "and that radiologist was rated the worst provider." | "and that radiologist was rated lowest of the four conditions (6.82 against 7.45 to 7.58 on a 1-to-10 scale)." | Yes |
| Project Literature Review/search/results-4-professionals.md | 19 | "and rated them as the worst provider." | "and rated them lowest of the four conditions (6.82 against 7.45 to 7.58 on a 1-to-10 scale)." | Yes |
| Project Literature Review/search/verified-view.md | 35 | "and was rated the worst provider." | "and was rated lowest of the four conditions (6.82 against 7.45 to 7.58 on a 1-to-10 scale)." | Yes |
| Project Literature Review/search/_review-design-v1.md | 56 | "was rated the worst provider" | (none) | No (reviewed proposal, superseded) |

Already correct: verified-view line 39 (the correction note), Abstract v8 line 55, v8 change log item 11. Abstract v7 does not contain the phrase.

### 3. Squash

| File | Line | Old | New | Changed? |
|---|---|---|---|---|
| Project Literature Review/Literature Review - Candidate papers.md | 39 | "Duke's Squash live checker showed more wrong matches than right ones in 2020." | "Duke's Squash live checker showed 20 irrelevant matches among its 34 pop-ups at the 2020 State of the Union (fact-checks unrelated to what was said, not false verdicts)." | Yes |
| Project Literature Review/search/results-5-board-products.md | 176 | "The lab responded by adding a human editor who approves each pop-up." | "The lab responded by building an interface for a human editor to approve each pop-up before it appears." | Yes |
| Project Literature Review/search/results-5-board-products.md | 177 | "A fully automated live checker produced more wrong matches than right ones, and its makers added a human in the loop." | "A fully automated live checker showed more irrelevant matches than relevant ones (20 of 34 pop-ups; unrelated fact-checks, not false verdicts), and its makers began adding a human in the loop." | Yes |
| Project Literature Review/search/_review-design-v2.md | 23 | "20 of 34 pop-ups were irrelevant matches, and a human editor was added." | "20 of 34 pop-ups were irrelevant matches, and the lab began building an interface for a human editor to approve each pop-up." | Yes |
| Project Literature Review/search/_review-design-v1.md | 47 | "a live checker that produced more wrong matches than right ones ..., after which a human editor was added" | (none) | No (reviewed proposal, superseded) |
| Project Literature Review/search/_review-design-v1.md | 49 | "live checkers produce false flags in practice" | (none) | No (same reason) |

Not changed and left as is: results-5 line 176 "14 relevant fact-checks and 20 non-relevant ones" is the source's own wording and is correct. Abstract v7 and v8 do not mention Squash.

### 4. Kim et al. (2026) cited as showing that AI contradiction lowers trust

| File | Line | Old | New | Changed? |
|---|---|---|---|---|
| CMPT 863 Project Abstract v8.md | 5 | "Prior studies show that an AI's disagreement lowers trust in a human expert when the expert stays silent (Kim et al., 2026; Madanay et al., 2025; Chen et al., 2026)" | "... (Madanay et al., 2025; Chen et al., 2026)" | Yes. The old wording was still in v8. Kim et al. are still cited in line 55 (closest design), so the reference list is unchanged. |
| CMPT 863 Project Abstract v8 - changes from v7.md | 5 | "prior work shows the AI-disagreement penalty with a silent expert (Kim 2026, Madanay 2025, Chen 2026)" | "... (Madanay 2025, Chen 2026)" | Yes, plus a new item 12 recording this correction and pointing to this file |
| Project Literature Review/search/results-6-factchecking.md | 141 | "Supports H1 when the AI's flag is right." | "Bounds H1 rather than supporting it: the advisor's error lowered trust and how the AI was involved changed nothing, so it does not show that an AI's flag or contradiction as such costs the expert." | Yes |
| Project Literature Review/Literature Review - Candidate papers.md | 62 | "Supporting: Madanay 2025, Du 2026, Dobbs 2026, Horne & Nevo 2025, Kim et al. 2026, Prike 2025 (known)." | "Supporting: Madanay 2025, Du 2026, Dobbs 2026, Horne & Nevo 2025, Prike 2025 (known). Kim et al. 2026 bounds it: the advisor's error lowered trust and how the AI was involved changed nothing." | Yes |
| Project Literature Review/search/review-A.md | 11, 87 | "AI disagreement or an AI-detected error lowers lay ratings of an expert (Kim et al., 2026; ...)"; "H1 ... Supported by Kim et al. (2026)" | (none) | No (selection-reviewer record) |
| Project Literature Review/search/review-B.md | 17 | Kim, Daneshi et al. (2026) grouped with studies where "AI disagreement lower[s] ratings of the person flagged" | (none) | No (reviewer record) |
| Project Literature Review/search/review-C.md | 16, 107 | "an AI that disagrees lowers lay ratings of ... advisor (... Kim et al., 2026 ...)"; "H1 ... Supported by ... Kim et al. (2026)" | (none) | No (reviewer record) |

Already correct: Candidate papers lines 49 and 168; reduced-view entry 1; verified-view entry 1; _review-design-v1 line 56; _review-design-v2 line 33. _review-design-v2 line 34 ("an AI's contradiction costs the contradicted expert") does not cite Kim and comes straight after the Kim bound in line 33. It was left alone, but it should stay tied to Madanay and Chen when the section is drafted.

Totals: 33 occurrences of the old wording found (14 for the burden of proof, 4 for Madanay, 6 for Squash, 9 for Kim), counting each line once. 11 were changed: 9 in Project Literature Review notes and 2 in Abstract v8 and its change log (v8 line 5, change-log line 5). A new item 12 was also added to the change log. The other 22 are in the submitted v7, the superseded abstracts v4 to v6, the reviewed proposal _review-design-v1.md, and reviewer records (review-A/B/C.md, references/reviews/related-work-review-A/C.md). None of those were edited.

## C. Other points noticed about these four sources

1. Chen et al. cite two different Kempt papers for the burden of proof: Kempt et al. (2023, AI & Society) on p. 4, and Kempt & Nagel (2022) on p. 11. Citing Kempt & Nagel (2022) is right for the p. 11 sentence.
2. Kempt & Nagel's burden shift is not specific to AI. It holds for any disagreeing second opinion, human included (p. 224). What is specific to AI is that the disagreement "cannot [be] overturn[ed] ... in any meaningful way based on an exchange of reasons" because the AI is "not receptive for diagnostic reasons" (p. 226), and the physician "cannot reject the evidence even if they think it is wrong" (p. 227). Their remedy is a human second opinion, not the physician's reason. Abstract v8 line 61 ("a reason meets that burden and a denial fails it") is the project's own inference, and Kempt & Nagel's text cuts somewhat against it. It is worth saying so, together with Braun's objection reported in the same paper (p. 226: automation "will not shift the burden of proof") and the Jongsma & Sand reply in the same issue.
3. Chen et al. (p. 2) describe Kempt & Nagel (2022) as having "revealed that consulting an AI's opinion on medical decisions is perceived as more responsible". Kempt & Nagel is a conceptual ethics paper with no perception data, so do not carry that sentence over.
4. Madanay: the provider rating and the recommend-to-others item were exploratory ("In exploratory analyses"). The primary outcome was agreement on a 1-to-6 scale (4.01 against 4.55 to 4.63). Cell sizes were 413, 401, 408 and 384. The "Open" item in verified-view line 39 ("check whether the published version reports the three outcome measures separately by condition") can be closed: Table 2 reports all three by condition.
5. Squash entry (results-5 lines 175-176): the ClaimBuster and Elasticsearch details are on the page. Two things are not on the cited page and must come from another Reporters' Lab post or need a source: the sentence "A UX test found viewers liked the idea of real-time fact-checks but were often confused by the pop-up design", and the list of fact-check publishers (Washington Post, FactCheck.org, PolitiFact). The page also says only that a public demo was planned "this summer" if the interface worked.
6. Kim: the AI was not always a monitor. In proactive conditions the advisor invoked it, and only in reactive conditions did it intervene on its own (p. 1515). Abstract v8 line 55 ("an AI monitored a human advisor") is a slight simplification. The "4 to 15 percent of variance" in reduced-view and results-6 matches the reported η2 of .04 to .15.
