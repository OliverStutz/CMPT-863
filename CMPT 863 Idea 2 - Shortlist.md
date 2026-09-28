# Idea 2 "Checked by the Camera": reading shortlist for the project abstract

Compiled 2026-09-26. Three independent reviewers (Opus agents, identical rubric) each ranked the 75 entries in [[CMPT 863 Idea 2 - Related papers]] and checked their top picks against the real abstracts or full texts. This file is the cross-examination of their three reports.

**Rubric (0 to 5 each, total out of 25):** trust in a *human* is the measured outcome / an AI or automated tool checks or contradicts that human / high-stakes professional context / direct use for the abstract / quality and accuracy of our summary.

**Coverage roles the final set should fill:** (a) pre-AI baseline that checking itself lowers trust, (b) AI checker contradicting a human, (c) false positives, (d) behaviour change of the checked person.

## Cross-examination

| Paper | Reviewer A | Reviewer B | Reviewer C | In top 7 | Role |
|---|---|---|---|---|---|
| Chen et al. (2026) | 1st (23) | 1st (22) | 1st (22) | 3 of 3 | b |
| Arkes et al. (2007) | 2nd (21) | 2nd (20) | 2nd (21) | 3 of 3 | b, partly a |
| von Schenk et al. (2024) | 4th (18) | 5th (18) | 5th (18) | 3 of 3 | c |
| Nyhan & Reifler (2015) | 7th (15) | 6th (16) | 6th (16) | 3 of 3 | d |
| Chae & Tewksbury (2026) | not ranked | 3rd (20) | 3rd (20) | 2 of 3 | b, counter-evidence |
| Prike et al. (2025) | runner-up (15) | 4th (18) | 4th (18) | 2 of 3 | main effect on the person |
| Strickland (1958) | runner-up (14) | 7th (15) | 7th (15) | 2 of 3 | a |
| Reif et al. (2025) | 5th (18) | runner-up (18) | runner-up (19) | 1 of 3 | partial a, d |
| Wintersieck (2017) | 3rd (18) | runner-up (17) | runner-up (16) | 1 of 3 | design template |
| Schweitzer, Ho & Zhang (2018) | 6th (16) | rejected | not ranked | 1 of 3 | disputed, see below |

Reviewers B and C recommended the identical final five: Strickland, Chen, Prike, von Schenk, Nyhan & Reifler. Reviewer A recommended Chen, Arkes, Wintersieck, von Schenk, Nyhan & Reifler.

## The seven to read

### Decisive winners (all three reviewers)

1. **Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, 213, 103824.** https://doi.org/10.1016/j.ijhcs.2026.103824
   - Why: the only paper in the list where an AI contradicts a professional and trust in that person is the outcome (cognitive, affective and behavioural trust). Ranked first by all three.
   - Caveats found in verification: the "doctor" was an LLM-simulated persona in a stress-management session, N = 135. The doctor requested the second opinion. The trust effect is indirect (disagreement raised perceived uncertainty and perceived laziness, which lowered trust) and held only when the doctor was seen as human-like.
   - Access: no free full text found. Use the SFU library.

2. **Arkes, H. R., Shaffer, V. A., & Medow, M. A. (2007). Patients derogate physicians who use a computer-assisted diagnostic aid. *Medical Decision Making*, 27(2), 189–202.** https://doi.org/10.1177/0272989X06297391
   - Why: four experiments in which the physician who used no aid was always rated highest. Experiment 4 compares heeding the aid with defying it, a manipulation we could borrow.
   - Caveat: the outcome is rated diagnostic ability, not a trust scale.
   - Access: not open access. SFU library or ResearchGate.

3. **von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, 27(7), 110201.** https://doi.org/10.1016/j.isci.2024.110201
   - Why: a fallible tool (66.86% accurate) lowers the threshold for treating a person as dishonest. Accusation rate was 19% without the algorithm and about 85% among those who requested it and got a "lie" prediction.
   - Caveats: only about 31% requested the algorithm. The outcome is accusation behaviour, not a trust rating. Low-stakes lab setting.
   - Access: free, https://pmc.ncbi.nlm.nih.gov/articles/PMC11302104/
[[]]
1. **Nyhan, B., & Reifler, J. (2015). The effect of fact-checking on elites: A field experiment on U.S. state legislators. *American Journal of Political Science*, 59(3), 628–640.** https://doi.org/10.1111/ajps.12162
   - Why: the cleanest causal evidence that the checked party changes behaviour. Legislators sent letters about the reputational risk of being fact-checked were less likely to receive a negative rating.
   - Caveat: no trust measure and no AI. It matters only if the abstract keeps the presenter-side study.
   - Access: free, https://bpb-us-e1.wpmucdn.com/sites.dartmouth.edu/dist/5/2293/files/2021/03/fact-checking-elites.pdf

### Strong (two of three reviewers)

5. **Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, 46(5), 1081–1100.** https://doi.org/10.1111/pops.13059
   - Why: with unfamiliar fictional candidates (N = 406), fact-checked candidates were rated much less favourably and got far fewer votes. Our recorded presenters are also strangers, so this grounds the expected main effect.
   - Caveat: human fact-checker, no AI.
   - Access: free, https://api.research-repository.uwa.edu.au/ws/files/463479508/Prike.2024.PoPs_pp.pdf

6. **Chae, J. H., & Tewksbury, D. (2026, online 2024). Perceiving AI intervention does not compromise the persuasive effect of fact-checking. *New Media & Society*.** https://doi.org/10.1177/14614448241286881
   - Why: the null result the abstract has to address. Three experiments (N = 3,978): neither AI-labelled nor human fact-checks changed feelings toward, or perceived competence of, the politicians being checked. The targets were well known, which is the likely moderator when read next to Prike et al.
   - Access: open access at the DOI.

7. **Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, 26(2), 200–215.** https://doi.org/10.1111/j.1467-6494.1958.tb01580.x
   - Why: the founding result that monitoring itself lowers trust in the monitored person when output is equal.
   - Caveat: no reviewer could read the original. The finding is confirmed only through secondary sources. Get the original through the SFU library before discussing it in depth.
   - Access: paywalled. Secondary account in Schweitzer, Ho & Zhang (2018), free at http://www.teckho.com/uploads/pdfs/HowMonitoringInfluencesTrust-ATaleofTwoFaces.pdf

### Alternates if one of the seven disappoints

- **Reif, Larrick & Soll (2025), PNAS.** Four preregistered studies, N = 4,439. Strong evidence of an interpersonal penalty, but for people who *use* AI, not people checked by it. Free at https://www.pnas.org/doi/pdf/10.1073/pnas.2426766122
- **Wintersieck (2017), American Politics Research.** Fact-checks shown alongside a recorded debate moved evaluations of the candidate. The closest format match to Study 1, but it overlaps with Prike et al. and is not open access.

## Where the reviewers disagreed

- **Schweitzer, Ho & Zhang (2018).** Reviewer A ranked it sixth as the modern heir of Strickland. Reviewer B rejected it because its finding runs the other way: trustors *over*-trusted compliance shown under monitoring. Both descriptions of the finding agree, so B's objection stands. Cite it as nuance, not as the baseline.
- **Arkes et al. (2007).** All three ranked it second, but B and C left it out of their final five because it shares the medical setting with Chen et al. It is a choice between a second professional-setting study (Arkes) and a broader spread of settings (Prike).
- **Chae & Tewksbury.** A did not rank it. B and C ranked it third and treat it as counter-evidence to cite, not as one of the core five.

## Rejected by at least two reviewers

- **Garrett & Weeks (2013).** All three, from the full text: the credibility measure concerns the fact-checking correction, not the person who made the claim. See corrections.
- **Lu et al. (2025).** Measures veracity judgements and sharing. Trust in the expert is never measured.
- **Gupta et al. (2026).** N = 34, measures belief in claims, no speaker evaluated. Good as a one-line gap citation for the device.
- **Falk & Kosfeld (2006).** Abstract principal-agent game. Nyhan & Reifler covers the same role with real professionals.
- **Mattozzi et al.** Working paper, not peer-reviewed.
- **Detjen et al. (2025).** Compares AI and physicians as advisers. The checked person is not evaluated.

## Gap: boardroom and cybersecurity

None of the three found a paper, inside or outside the list, set in a boardroom or cybersecurity context with trust in a person as the outcome. The seed paper (Aamir et al., CHI 2026) stays the only one. That supports the novelty claim. One outside paper surfaced twice (Libby et al., 2024, *Contemporary Accounting Research*, https://doi.org/10.1111/1911-3846.12941, on auditors using AI), but neither reviewer verified it and it is about tool use, not contradiction.

## Corrections to the related-papers file

Not yet applied to [[CMPT 863 Idea 2 - Related papers]].

| Entry | Problem | Found by |
|---|---|---|
| Garrett & Weeks (2013), and the headline finding at the top of the file | "Distrusted the source" refers to the credibility of the fact-checking correction, not trust in the corrected speaker. It should not be listed as one of the three closest designs. | A, B, C |
| Chen et al. (2026) | Omits that the doctor was LLM-simulated, the setting was stress management, N = 135, and the trust effect is indirect and moderated by anthropomorphism. | A, B, C |
| Reif et al. (2025) | "Adjust behaviour" is not supported by the abstract. Only anticipation of the penalty is. | A, C |
| Wintersieck (2017) | "Half true did nothing" is not in the abstract. Unverified. | A, B, C |
| von Schenk et al. (2024) | "Raising false accusations" is one-sided. Accusations rose, and accusation accuracy also rose (51% with a prediction, 41% baseline). Only about 31% requested the algorithm. | A, C |
| Arkes et al. (2007) | The measure is rated diagnostic ability, not trust. Experiment 4 used medical students. | A, B |
| Schweitzer, Ho & Zhang (2018) | "Direct heir of Strickland" hides that the result is misplaced trust, not reduced trust. | B |
| Lu et al. (2025) | The larger influence under disagreement refers to veracity judgements, not trust in the expert. | C |
| Liao et al. (2026) | Now open access. "Equal trust" comes from a press release. "Any explanation beat a black box" is not supported by the abstract. | C |
| Frey (1993) | Citation garbled. Correct: *Economic Inquiry*, 31(4), 663–670. | C |
| Prike et al. (2025) | Published online November 2024, print is 46(5), 1081–1100. | B, C |
| Schweitzer & Ho (2005) | Could not be verified. | B |
