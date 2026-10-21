# Reduced view: the ten papers to read first

Written 2026-10-18. These ten were chosen by all three selection reviewers (`review-A/B/C.md`) and none is cited in abstract v7 yet. Order is reading order. "Summary from" says what the summary rests on; "Full text" says whether the PDF is in `../references/` or still needs the SFU library. Each entry ends with what to check when reading, taken from the reviewers' must-read notes.

Rubric scores (out of 25) are from the search files: trust in a human measured / an AI checks the human / power or high-stakes setting / direct use / confidence.

---

## 1. Kim, D., Daneshi, R., Draper, B., & Sreedharan, S. (2026). Asymmetric trust effects of corrective AI in expert advisory workflows under epistemic dependence. *Proceedings of the AAAI/ACM Conference on AI, Ethics, and Society*, *9*(3), 1514–1525. https://doi.org/10.1609/aies.v9i3.43244

- **Venue and year:** AIES 2026 (AAAI/ACM). Score 24/25.
- **Summary from:** full text, checked twice (search agent and a second read of the PDF).
- **Full text:** in `../references/`.
- **Design:** 157 Prolific participants played students receiving course-planning advice from a human advisor. 3 x 2 between subjects: the advisor made an error or not; the AI was absent, invoked by the advisor, or monitored automatically and flagged the error. Every workflow ended with a correct schedule, and the AI was always right. Trust in the advisor was measured with the METI scale (expertise, integrity, benevolence), an entity trust scale, a 100-point rating and intention to reuse the advisor.
- **Finding:** advisor errors lowered every trust measure (4 to 15 percent of variance). How the AI was involved changed nothing, and a visible AI correction did not raise trust in the AI or the human-AI team. Trust stayed anchored to the human advisor.
- **Why it matters:** the closest published design. It means our novelty claim is not "first to measure trust in an expert whom an AI checks" but the combination of wrong flags, a checker held by the more powerful non-expert, the expert's answer as the repair step, and trust rated at each step. Their result that the AI's involvement did not change the penalty cuts against H4b.
- **Check when reading:** cell sizes (157 over six cells); whether monitoring without errors moved trust at all (bears on our "No issues found" slides); whether the advisor ever replied.

## 2. Kim, P. H., Ferrin, D. L., Cooper, C. D., & Dirks, K. T. (2004). Removing the shadow of suspicion: The effects of apology versus denial for repairing competence- versus integrity-based trust violations. *Journal of Applied Psychology*, *89*(1), 104–118. https://doi.org/10.1037/0021-9010.89.1.104

- **Venue and year:** J. Applied Psychology, 2004. Score 15/25.
- **Summary from:** full text (published PDF, added by Oliver 2026-10-18; the three check items below were read from it).
- **Full text:** in `../references/` as `Kim-2004-apology-vs-denial-JAP.pdf`.
- **Design:** two studies with participants acting as a hiring manager watching a videotaped interview of a job applicant at an accounting firm. Study 1 (N = 200: 63 undergraduates, 137 graduate business students), 2 x 2 between subjects: the alleged past violation concerned competence (filing an incorrect tax return through lack of knowledge) or integrity (filing it intentionally), and the candidate apologised or denied. Study 2 (N = 444: 320 undergraduates, 124 graduate students) added a third factor: after the interview, a memo from the company's HR department reported that the State Board of Public Accountancy had investigated and ruled the candidate guilty or innocent. Measured perceived competence and integrity (trusting beliefs) and willingness to risk, hiring, and job responsibilities (trusting intentions).
- **Finding:** Study 1: apology repaired trust better after the competence violation and denial after the integrity violation, on perceived integrity and on the intentions; the interaction did not appear for perceived competence. Study 2: evidence of innocence raised all five measures; beyond that, a response that did not fit the later evidence hurt: an apology followed by evidence of innocence did worse than a denial followed by innocence, and a denial followed by guilt did worse than an apology followed by guilt.
- **What the full text settles:** the denial was not bare. The candidate "denied responsibility for the trust violation, attributed the allegation to bad office politics at the previous firm, and stated that the firm would not have to worry about her committing such an infraction if she were hired" (p. 108). So their denial carried an account and an assurance, and the evidence of innocence came from a third party (a state board ruling), not from the accused. Our "denial with no reason" is closer to reticence (Ferrin et al., 2007) than to this denial.
- **Why it matters:** our flags are competence-type and our CISO is innocent by design. This paper says an innocent party's denial with an account repairs trust, and that an admission that later proves unnecessary costs trust. It supports rewording H3 as "a denial without a reason restores less trust than a reasoned answer", and it supplies the measures (competence, integrity, willingness to risk) and the videotaped-applicant design.
- **Still to check:** nothing on the design. For the review, pull the Study 2 cell means from Table 6 (p. 113) if a number is needed.

## 3. Madanay, F. L., O'Donohue, L. S., & Zikmund-Fisher, B. J. (2025). Patient reactions to artificial intelligence–clinician discrepancies: Web-based randomized experiment. *Journal of Medical Internet Research*, *27*, Article e68823. https://doi.org/10.2196/68823

- **Venue and year:** JMIR, 2025. Score 24/25.
- **Summary from:** publisher abstract (Europe PMC).
- **Full text:** in `../references/` as plain text from the Europe PMC XML (no figures).
- **Design:** 1,606 US adults imagined a lung-cancer screening CT and were randomised to a result from a radiologist only, AI and radiologist in agreement, a radiologist who recommended more testing than the AI, or one who recommended less. Measured agreement with the recommendation, likelihood of recommending the radiologist, and how good a provider the radiologist is.
- **Finding:** when the radiologist recommended less than the AI, agreement fell to a mean of 4.01 against 4.55 to 4.63 in the other three conditions, and that radiologist was rated lowest of the four conditions (6.82 against 7.45 to 7.58 on a 1-to-10 scale). A radiologist who recommended more than the AI paid no penalty.
- **Why it matters:** the strongest support for H1 from a large sample, and a warning that the penalty is asymmetric: it depends on whether the AI looks more cautious than the expert. Our flags play the cautious role.
- **Check when reading:** which direction of discrepancy maps onto our flags; our flags say the slide is wrong, not that more caution is needed.

## 4. Hendriks, F., Kienhues, D., & Bromme, R. (2015). Measuring laypeople's trust in experts in a digital age: The Muenster Epistemic Trustworthiness Inventory (METI). *PLOS ONE*, *10*(10), e0139309. https://doi.org/10.1371/journal.pone.0139309

- **Venue and year:** PLOS ONE, 2015. Score 19/25.
- **Summary from:** publisher abstract.
- **Full text:** in `../references/`.
- **Design:** scale development. Exploratory factor analysis (n = 237) and confirmatory factor analysis (n = 345) of laypeople's ratings of experts met online, then an experiment (n = 137) that varied what was said about the source. Fourteen bipolar adjective pairs.
- **Finding:** laypeople's judgements of an expert form three factors, expertise, integrity and benevolence, and all three shift when the source description changes.
- **Why it matters:** this is the measure for the three-part trust question (competence, benevolence, integrity) in a non-expert-judges-expert setting. Kim et al. (2026) used it. Its expertise factor is what "Architecture likely wrong" attacks.
- **Check when reading:** the 14 item pairs, so the three ratings in our study can be taken from them; how long the scale takes.

## 5. Opdenbusch, J. C., Hielscher, J., & Sasse, M. A. (2025). "Where are we on cyber?": A qualitative study on boards' cybersecurity risk decision making. In *Proceedings of the 2025 Network and Distributed System Security Symposium*. Internet Society. https://doi.org/10.14722/ndss.2025.240595

- **Venue and year:** NDSS 2025. Score 16/25.
- **Summary from:** publisher abstract and the NDSS page.
- **Full text:** in `../references/`.
- **Design:** in-depth interviews with 18 C-level managers, board members, CISOs and consultants at some of the largest UK-based companies.
- **Finding:** a power imbalance: board members often do not ask executives and CISOs the right questions because they fear being exposed as IT novices, so boards depend on whoever supplies their cyber information and reduce cyber risk to budget decisions.
- **Why it matters:** grounds the default silence in our design and the prediction that a checker gives board members a reason to ask (H4a). The budget framing matches our scenario.
- **Check when reading:** the exact "IT novices" wording before quoting it.

## 6. Chae, J. H., & Tewksbury, D. (2026). Perceiving AI intervention does not compromise the persuasive effect of fact-checking. *New Media & Society*, *28*(1), 191–211. https://doi.org/10.1177/14614448241286881

- **Venue and year:** New Media & Society, 2026 (online 2024). Score about 16/25 in the shortlist.
- **Summary from:** the published abstract, read from the PDF on 2026-10-18.
- **Full text:** in `../references/` (open access, 21 pages).
- **Design:** three online experiments, N = 3,978. Participants read fact-checks identified as human-generated or AI-assisted (AI fully generating the fact-check, or AI retrieving human fact-checks).
- **Finding:** the fact-checks raised truth discernment whether labelled AI or human, even among participants without a positive prior attitude toward AI. Neither AI-generated nor human fact-checks significantly affected participants' feelings toward, or their perceptions of the competence of, the targeted politicians.
- **Why it matters:** the direct counter-result to H1: a check can change beliefs about claims without touching the standing of the person checked. The likely moderator, read next to Prike et al. (2025), is that the targets were well known.
- **Check when reading:** how competence was measured; which politicians; whether a flag on an unknown expert would be expected to behave differently.

## 7. Thorson, E. (2016). Belief echoes: The persistent effects of corrected misinformation. *Political Communication*, *33*(3), 460–480. https://doi.org/10.1080/10584609.2015.1102187

- **Venue and year:** Political Communication, 2016. Score 14/25.
- **Summary from:** publisher abstract.
- **Full text:** the article is missing (SFU: https://proxy.lib.sfu.ca/login?url=https://doi.org/10.1080/10584609.2015.1102187). The 2013 dissertation with the same experiments is in `../references/` as a fallback; do not cite page numbers from it.
- **Design:** three experiments. Participants read negative information about a political candidate that was then discredited, in some conditions immediately. Attitudes toward the candidate were measured afterwards.
- **Finding:** the negative information kept lowering attitudes after it was discredited and after participants accepted the correction, even when the correction came immediately. The echoes arise through both automatic and deliberate routes.
- **Why it matters:** the main argument against H2's "above the start": even a convincing answer should leave a residue. All three reviewers suggest making "above the start" exploratory.
- **Check when reading:** the size of the residue after an immediate correction; whether the corrected person or a third party delivered the correction.

## 8. Lowry, M. R., Vance, A., & Vance, M. D. (2026). Inexpert supervision: Field evidence on boards' oversight of cybersecurity. *Management Science*, *72*(2), 783–804. https://doi.org/10.1287/mnsc.2023.04147

- **Venue and year:** Management Science, 2026. Score 15/25.
- **Summary from:** publisher abstract (Crossref).
- **Full text:** in `../references/` (published version, CC BY-ND). A 2021 draft is also saved; do not cite it.
- **Design:** field study with interviews of directors with and without cyber expertise, cybersecurity executives and consultants who support boards (a search summary gives 20 directors, 11 executives, 7 consultants; confirm in the paper).
- **Finding:** non-expert directors tried to oversee diligently, but even when they did the same activities as expert directors their oversight was "mostly symbolic", and they did not see this. Expert directors, CISOs and consultants saw the gap.
- **Why it matters:** the strongest scholarly ground for why a board member would reach for a checker and why they could not judge its flags. Pairs with the SEC (2023) decision not to require board cyber expertise.
- **Check when reading:** the "mostly symbolic" wording; the interview counts.

## 9. Gupta, C., Aritonang, N. V., Rajendran, D. P. D., Danry, V., Maes, P., & Nanayakkara, S. (2026). Feeling the Facts: Real-time wearable fact-checkers can use nudges to reduce user belief in false information. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems*. Association for Computing Machinery. https://doi.org/10.1145/3772318.3791679

- **Venue and year:** CHI 2026. In the earlier shortlist as a gap citation.
- **Summary from:** the abstract and measures list, read from the arXiv PDF on 2026-10-18.
- **Full text:** in `../references/` (arXiv version; check page numbers against the ACM version).
- **Design:** controlled study, N = 34, simulating a wearable that listens for verifiable claims, verifies them on the web, and gives a haptic nudge with a glanceable overview; compared with a no-support baseline while watching videos.
- **Finding:** instant, body-integrated feedback improved real-time truth discernment and increased verification activity, but introduced over-reliance when the system erred by missing a false claim or flagging a true one as false. The measures include trust in the AI, topic relatability and openness to updating beliefs; there is no rating of the speaker.
- **Why it matters:** the nearest real-time wearable checker, with the false-positive problem built in. It measures belief in claims and trust in the AI, not trust in the person speaking, which is the gap we fill.
- **Check when reading:** the system's error rate (an 87.5 percent figure appears in the text) and exactly what "trust" was measured on; confirm no speaker rating exists.

## 10. Ferrin, D. L., Kim, P. H., Cooper, C. D., & Dirks, K. T. (2007). Silence speaks volumes: The effectiveness of reticence in comparison to apology and denial for responding to integrity- and competence-based trust violations. *Journal of Applied Psychology*, *92*(4), 893–908. https://doi.org/10.1037/0021-9010.92.4.893

- **Venue and year:** J. Applied Psychology, 2007. Score 15/25.
- **Summary from:** publisher abstract (OpenAlex).
- **Full text:** in `../references/` as `Ferrin-2007-silence-speaks-volumes-JAP.pdf` (published PDF, added by Oliver 2026-10-19 from SFU PsycArticles).
- **Design:** two laboratory studies in two countries. An accused party responded to a competence or integrity violation with apology, denial, or reticence (declining to confirm or deny). Measured trust.
- **Finding:** reticence did worse than the better of apology and denial in each case: for integrity violations it, like apology, failed to address guilt; for competence violations it, like denial, failed to signal that the person would improve.
- **Why it matters:** reticence is the nearest tested analogue to a reply with no information, so it is the better support for H3 than Kim et al. (2004). With Kim et al. (2004) it decides whether "denial with no reason" lowers trust or only restores less.
- **Check when reading:** the exact wording of the reticence response; whether the competence case had an innocence condition.

---

## Next three if time allows

- **Mayer, R. C., Davis, J. H., & Schoorman, F. D. (1995).** The ability, benevolence, integrity definition of trust. AMR 20(3), 709–734. https://doi.org/10.2307/258792. Full text in `../references/`.
- **Schiff, K. J., Schiff, D. S., & Bueno, N. S. (2025).** Five experiments, over 15,000 adults: a politician who falsely called a scandal "fake news" gained more support than one who stayed silent or apologised. A counter-result to H3 from politics. APSR 119(1), 71–90. https://doi.org/10.1017/S0003055423001454. Full text in `../references/`.
- **Proudfoot, J. G., Cram, W. A., Madnick, S., & Coden, M. (2023).** 35 interviews: boards judge CISO competence by how well the CISO communicates; a credible-seeming but weak CISO can be "(dangerously)" overrated. MISQE 22(4), 235–250. https://doi.org/10.17705/2msqe.00084. Full text in `../references/`.
