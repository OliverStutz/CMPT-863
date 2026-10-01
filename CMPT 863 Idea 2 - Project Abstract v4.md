# Does a Vision-LLM Content Checker in a Board Member's Hands Erode or Build Trust in the Presenting CISO?

## Research question

Vision-language models (VLMs) let a listener point a camera at a slide and have it checked while the presenter is still talking. Boards are a plausible settings of high stakes decision making with non-experts. Aamir et al. (2026) reports that gaps in cyber literacy and confidence make it hard for directors to question cyber-risk reports, that boards without such literacy often accept management's account, and that directors felt more confident pushing back where independent reviews were routine. An LLM-assisted camera checker offers such a review on the spot. But it might lack context to complex resolutions, so a flag from the llm can be right by the textbook and wrong for this company and context. The relationship it enters is unequal: the power is in the board's hands, which decides on the budget, while the expertise sits with the CISO, on whom the board depends.

> **Research Question:** When a board member uses a VLM content checker during a CISO's briefing, how does their trust in the CISO change from the start to the end of the briefing, compared with a board member who follows the same briefing, with the same explanations, without a checker?

## Hypotheses

Trust means the participant's rating of the CISO on a scale from 0 to 10.

- **H1 (flag).** Right after a flagged slide, trust is lower than after the slide before, and lower than in the neutral group at the same slide. The slides contain no errors, so only the flag differs.
- **H2 (answer).** When the CISO answers with the company's reason, trust after the answer is higher than after the flag. It can end above the level of the neutral group, which hears the same reason without a flag.
- **H3 (non-answer).** When the CISO meets a flag with a denial and no reason, trust after the answer is lower than after the flag.
- **H4 (participation and amplification).** The checker makes the board member more active and raises the stakes of the answer. (a) Participants in the checker group question the CISO at more slides than participants in the neutral group say they have questions. (b) After an answer, trust moves further in the checker group than in the neutral group: further up after the company's reason, further down after the denial.

## Related work

Strickland (1958) had 40 male students play a supervisor who was told to check one of two fictitious workers after nine of ten work periods and the other after two. Both workers produced the same output, and every check found the work in order. Afterwards, asked which worker they trusted more, more supervisors named the rarely checked one than the often checked one. Like Strickland's supervisors, our board member holds power over the person checked. But his checks found nothing and the workers never spoke; our checker claims faults that are not there, and the CISO answers. Because checking alone may lower trust, the checker also shows "No issues found" on slides 1 and 5, so that flagged and unflagged slides can be compared. This informs H1.

Prike et al. (2025) introduced 406 U.S. adults to two unfamiliar, fictional candidates through a profile and a series of posts. When a nonpartisan fact-checking organisation marked as false eight statements that favoured one candidate, participants rated him as less competent and less moral than his opponent, compared with a group that saw only neutral posts, and fewer voted for him. We borrow the profile, the 0-to-10 rating and the end-of-study rating of how true each checked claim is. Their checker was human and right, the candidate stayed silent, and ratings were taken once, at the end; ours is an AI tool, wrong for this company, the CISO answers, and trust is rated after every step. This informs H1.

Chen et al. (2026) is the closest design (*N* = 135). At the end of a short online therapy chat, a chatbot presented as a human doctor said only that his AI assistant agreed or disagreed with him. Compared with agreement, disagreement left participants less sure about their own situation and made the doctor seem lazier, and both went with lower trust in him. Because the doctor never answered, the study cannot separate the disagreement from the silence; the authors suspect the silence and suggest an explanation, which they did not test. There the doctor brought in the AI himself. We give the checker to the board member and rate trust after the flag and after the answer: H2 tests an answer with the company's reason, H3 a denial.

Greitemeyer and Sagioglou (2015) had online participants rate a researcher accused of misconduct, and again after they read that he had been cleared (*N* = 185 and 358). The accused researcher was rated about two points below one never accused, on a scale from −5 to +5; after the clearing he rose by about two points, to that level but not reliably above it. He ended clearly above a never-accused researcher only when that researcher had been described with less information, which is why our neutral group hears the same answers. In Mickelberg et al. (2025), a retraction removed the effect of an allegation on how likeable a person seemed in two of three experiments; the authors note that rating after every step may draw attention to the latest information, and both our groups rate at the same steps. In these studies a third party did the clearing; our CISO answers for himself. This informs H2: recovery is expected; a reliable gain would go beyond these results.

A checker may also change what the board member does. Directors in Aamir et al. (2026) felt more confident challenging management where independent reviews were routine. In von Schenk et al. (2024), 2,040 people each judged whether a stranger's written statement was a lie, knowing that an accusation would cost the writer money. Without help, about one in five accused. Those who chose to consult a lie-detection algorithm that was right about two times in three mostly followed it: when it said "lie", almost 85% accused. The accused could not answer. This informs H4a. Chen et al. note that an AI's contrary opinion shifts the burden of proof onto the expert. After a flag, a reason meets that burden and a denial fails it, so the same answer should move trust further than in the neutral group, where nothing has to be proved (H4b). Reif et al. (2025) found that employees described as getting help from generative AI were rated lazier and less competent than employees who got the same help from another source, so the CISO never mentions AI; and because naming the checker may look like relying on AI, we log which question participants choose and how often they use AI themselves.

## Method

### Study design

Two groups, one per version of the simulation, assigned in turn at sign-up (round robin). Participants read a scenario in which they sit on the board and decide on the CISO's budget request, then a profile that establishes the CISO's experience, and give a first trust rating. In the checker group, the checker is introduced after this first rating. The CISO then presents five slides in a web-based simulated boardroom. The slides contain no errors and are the same in both groups. On slides 2, 3 and 4 the CISO is asked a question and answers. Two answers give the company's reason; one is a denial with no reason. Which slide gets the denial is balanced across participants within each group.

- **Checker group.** A checker panel shows the camera image of each slide. On slides 2, 3 and 4 it marks a region and shows a flag, for example "Architecture likely wrong" or "Numbers do not add up". On slides 1 and 5 it shows "No issues found". After a flag the participant rates trust and picks an action: say nothing, ask the CISO to explain, or challenge the CISO with the checker's finding. If the participant says nothing, the board chair asks a routine question. The CISO's answer is the same whoever asks. The participant rates trust again.
- **Neutral group.** No checker. After slides 2, 3 and 4 the participant rates trust and says whether they have questions about the slide. There is no checker to propose one, so the chair asks the routine question and the CISO gives the same answer. The participant rates trust again.

Every participant sees the same slides and hears the same answers. The groups differ only in the checker, in the questions it proposes, and in who asks.

The checker is simulated. Flags are drafted with a real VLM from images of the slides and then fixed, so every participant sees the same output. Every flag is wrong for this company, and each full answer explains why. After the last slide, every participant is told that none of the slides had a real issue, that the flags were fixed in advance, and that the study measures how flags from an LLM affect the trust of non-experts in the presenter.

### Measures

- We want to elevate the trust at the initial slide to give a "base" trust rating, as different people have different ways to give trust.
- We want to use a 0-10 rating to trust/not trust the CISO
- Trust has to be measured at each step e.g. before/after so that we can understand how it changes during questions/actions etc.
- At the end of the presentation ask the Board member if they approve the budget to see if trust decrease/increase results in actual decisions.
- Analyse study group one's impact vs study group two, to understand impacts of checker outputs and interaction by the participant.
- Verify how many participants have had cyber security related knowledge to split them apart. (expert knowledge present already)
- Degradation of trust through the checkers flags
- Ask for AI-Usage, Age group, Field of expertise


### Validity

The change from the first to the last rating answers the research question; each hypothesis compares the groups at the same step, or a participant with their own previous rating. Both groups read the same profile, see the same error-free slides, rate at the same points and hear the same answers, so a difference between the groups cannot come from the presenter, the slides or what participants learn. It comes from the checker's flags and from one thing tied to them: checker participants may ask the CISO themselves, while in the neutral group the chair always asks. The first rating is taken before the checker is introduced, so it is each participant's own starting point. Each flag and answer is scored as the change from the rating just before it, so earlier slides do not carry over. Slide 1, where the checker reports no issue, shows whether its presence alone moves trust. The denial rotates over slides 2 to 4 within each group, so it is not tied to one slide. Sign-ups fill the groups in turn, and we check that the groups start alike on the first rating. Limits: participants choose whether to ask, so comparisons between actions only describe them, and a yes to "Do you have questions about this slide?" costs less than asking. Participants stand in for directors, with one CISO and one topic, and are told the checker is live; a debrief follows.

## Recruitment

We will recruit 40 adults of any education level and field through our personal and professional networks, mailing lists and LinkedIn, and record each participant's education, age group, field and self-rated knowledge of cybersecurity. The study runs online in about 20 minutes as a web-based simulation of the briefing, with or without the VLM checker. Sign-ups alternate between these two versions (round robin), not at random.

## References

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–22). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791142

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, *213*, Article 103824. https://doi.org/10.1016/j.ijhcs.2026.103824

Greitemeyer, T., & Sagioglou, C. (2015). Does exonerating an accused researcher restore the researcher's credibility? *PLOS ONE*, *10*(5), Article e0126316. https://doi.org/10.1371/journal.pone.0126316

Jian, J.-Y., Bisantz, A. M., & Drury, C. G. (2000). Foundations for an empirically determined scale of trust in automated systems. *International Journal of Cognitive Ergonomics*, *4*(1), 53–71. https://doi.org/10.1207/S15327566IJCE0401_04

Mayer, R. C., & Davis, J. H. (1999). The effect of the performance appraisal system on trust for management: A field quasi-experiment. *Journal of Applied Psychology*, *84*(1), 123–136. https://doi.org/10.1037/0021-9010.84.1.123

Mickelberg, A. J., Walker, B., Ecker, U. K. H., Howe, P. D. L., Perfors, A., & Fay, N. (2025). Did he or didn't he? Mixed evidence for the continued influence of retracted misinformation on person impressions. *PLOS ONE*, *20*(5), Article e0322045. https://doi.org/10.1371/journal.pone.0322045

Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, *46*(5), 1081–1100. https://doi.org/10.1111/pops.13059

Reif, J. A., Larrick, R. P., & Soll, J. B. (2025). Evidence of a social evaluation penalty for using AI. *Proceedings of the National Academy of Sciences*, *122*(19), Article e2426766122. https://doi.org/10.1073/pnas.2426766122

Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, *26*(2), 200–215. https://doi.org/10.1111/j.1467-6494.1958.tb01580.x

von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, *27*(7), Article 110201. https://doi.org/10.1016/j.isci.2024.110201

## AI transparency

We used AI to search for and screen related papers, to check every claim in this text against the full text of the cited paper,  and articulate the writers ideas. The team reads all cited papers before the final submission.

---

# Working notes (delete before submission)

Version 4, written 2026-09-29 after the cited papers were read (section 8 says what was read of each) and each was checked against the hypotheses by an independent reviewer (reports in `references/reviews/`). Version 3 is [[CMPT 863 Idea 2 - Project Abstract v3]] and is unchanged. Submission is due 2026-09-30, 11:59 pm.

## 1. What changed from version 3

| Topic                       | Version 3                                               | Version 4                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Why                                                                                                                                                                                                                                                            |
| --------------------------- | ------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Hypotheses                  | Your H1–H4 as written                                   | Same four, in testable form: each says what is compared with what                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Rubric asks for a specific, testable hypothesis                                                                                                                                                                                                                |
| H3                          | No condition for it: every answer resolved the flag     | One of the three answers is a denial with no reason, on slide 2, 3 or 4, balanced across participants                                                                                                                                                                                                                                                                                                                                                                                                                                                    | H3 could not be tested                                                                                                                                                                                                                                         |
| H4                          | "Heavily dependent on actions and responses"            | Two testable parts: the checker group questions the CISO more often than the neutral group reports having questions, and the same answer moves trust further in the checker group, up or down                                                                                                                                                                                                                                                                                                                                                            | Your point of 2026-09-29: the checker encourages participation and raises the stakes of the answer                                                                                                                                                             |
| Neutral group               | Heard no questions and no answers                       | Says whether they have questions about the slide (yes/no), then the chair asks a routine question and the CISO gives the same answer                                                                                                                                                                                                                                                                                                                                                                                                                     | Otherwise a gain in the checker group could come from hearing more information, not from the challenge. Greitemeyer and Sagioglou fell into the same trap. The yes/no gives the participation baseline for H4a and shows which slides raise doubt on their own |
| "Say nothing"               | The briefing continued                                  | The chair asks, so every participant hears the answer                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | Same reason. The action now changes only who asks and whether the AI is named                                                                                                                                                                                  |
| Checker on unflagged slides | Not specified                                           | Shows "No issues found"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | Strickland: checking alone lowers trust. Flagged and unflagged slides can now be compared within the checker group                                                                                                                                             |
| Related work                | 4 papers plus Aamir                                     | 5 papers plus Aamir, with Greitemeyer and Sagioglou back in; Ecker and Rodricks and Mickelberg et al. cited in one sentence                                                                                                                                                                                                                                                                                                                                                                                                                              | H2 had no grounding; the two extra citations show that recovery is the usual finding, not a single result                                                                                                                                                      |
| Aamir                       | "found that gaps … limit"                               | "report", plus what directors do about it                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | Interview study, 13 self-reports, no counts                                                                                                                                                                                                                    |
| Strickland                  | "participants who monitored one subordinate more often" | "were told to check"                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | The schedule was assigned; the old wording implied the reverse cause                                                                                                                                                                                           |
| Chen                        | "participants who saw the doctor as human-like …"       | "gave no explanation"; the human-like clause is gone                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | The human-like result is exploratory and underpowered; the no-explanation point is what matters for H2 and H3                                                                                                                                                  |
| Reif | "disclose it less willingly" | Reduced to two clauses: people described as getting AI help are rated lazier and less competent (Study 2), so the CISO never mentions AI, and we log whether participants name the checker | The disclosure finding was a stated intention in an imagined scenario (Study 1), and Reif tests none of H1–H4 |
| Mayer and Davis             | "a selection of six items"                              | The 6 ability items and the reworded trust items                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Their own 4-item trust scale had low reliability (α = .59); six items across three factors would be worse                                                                                                                                                      |
| Measures added              |                                                         | Belief in the flags at the end (both groups); AI-use frequency                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | Prike's baseline design; Reif's moderator                                                                                                                                                                                                                      |
| Stimulus rule               |                                                         | The CISO never uses or mentions AI                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Reif's penalty would land on the CISO and add a second drop                                                                                                                                                                                                    |
| Verification                | Six claims open                                         | All claims checked against the full texts                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                | See section 8                                                                                                                                                                                                                                                  |
| Related work, second pass   | First rewrite                                           | Three independent reviewers scored the section 2 of 3 against the course rubric and agreed on the same fixes (reports in `references/reviews/related-work-review-A/B/C.md`). Applied: von Schenk et al. (2024) added for H4a; H4b grounded in Chen's burden-of-proof point; Reif reduced to the stimulus rule; Prike's comparison group corrected; Strickland tied to the "No issues found" design; Chen's comparison named (agreement); Greitemeyer's Study 2 lesson added; Ecker and Rodricks cut; the gap sentence in the Research question rewritten | H4 had no paper behind it, and three claims would not have survived a marker who opened the papers                                                                                                                                                             |
| Sample and assignment       | 60, 30 per group, assigned at random                    | 40, 20 per version, assigned in turn at sign-up (round robin), not at random                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | Your decision of 2026-09-29                                                                                                                                                                                                                                    |

## 2. Decisions that are yours to confirm

1. **The denial.** Suggested wording: "It's fine as it is, and I'd rather not go into the details here. Let's move on." It is both a negation and a non-answer, which covers your H3 wording. If you want two kinds (denial, non-answer), the design needs a fourth flagged slide or a second version.
2. **Where the denial sits.** Balanced across slides 2, 3 and 4; with 20 per group that is 7, 7 and 6 participants per position. A fixed slide is simpler to build but ties H3 to one slide's content.
3. **The chair.** A scripted second character who asks one routine question at slides 2–4 in the neutral group, and in the checker group when the participant says nothing ("Can you walk us through this slide?"). Voice or text is open.
4. **The neutral group's prompt.** "Do you have questions about this slide?" with yes/no. It does not change what happens next; the chair asks either way. It is the counterpart of the checker group's "ask" and "challenge".
5. **Extra citation.** Mickelberg et al. (2025) appears in one sentence, for the repeated-rating caveat; Ecker and Rodricks (2020) was cut after the related-work review. The rubric asks for 4–5 papers discussed; these two are cited, not discussed. Cut them if the list looks long.
6. **Kept from version 3:** budget decision, trust in the checker, the questionnaire at start and end. *N* is now 40, 20 per version of the simulation, assigned in turn at sign-up (round robin), not at random. I read your "20 students per study" as 20 per version, 40 in total; if you meant 20 in total (10 per version), the study would only detect differences about as large as the whole spread of trust ratings (*d* ≈ 1.0).
7. **Trust-repair studies.** All three reviewers named Kim et al. (2004, apology against denial after a competence or integrity accusation) and Ferrin et al. (2007, declining to comment) as the literature that tests the accused person's own reply, which is what H2 and H3 vary. They cite them from memory. Both are in the *Journal of Applied Psychology*; fetch them through the SFU library (links in `references/README.md`), read them, and then add one sentence to the Greitemeyer paragraph. Not cited until read.
8. **Deception.** Participants are told the checker is a live AI. The debrief is now in the Method (no slide had a real issue, flags fixed in advance, purpose of the study). Still needs a check against the course's ethics rules.

## 3. Session flow, about 20 minutes

| Step | Checker group | Neutral group |
|---|---|---|
| 1 | Consent, scenario | Consent, scenario |
| 2 | CISO profile | Same |
| 3 | First trust rating, questionnaire, then introduction of the checker | First trust rating, questionnaire |
| 4 | Slide 1, checker shows "No issues found", trust rating | Slide 1, trust rating |
| 5 | Slide 2 with flag, trust rating, action, answer (or denial), trust rating | Slide 2, trust rating, "Do you have questions about this slide?", chair asks, answer (or denial), trust rating |
| 6 | Slide 3, same steps | Same |
| 7 | Slide 4, same steps | Same |
| 8 | Slide 5, "No issues found", trust rating | Slide 5, trust rating |
| 9 | Questionnaire, budget decision, belief in the three checked claims (the rating after slide 5 is the last trust rating) | Same |
| 10 | Trust in the checker | Not asked |
| 11 | Background: education, age group, field, cybersecurity knowledge, AI use. Debrief: no slide had a real issue, the flags were fixed, the study measures the effect of LLM flags on non-experts' trust | Same |

Both groups rate at the same points, so every rating can be compared between groups.

## 4. Stimulus sketch

Three flags, built from your examples. The denial replaces one of the three full answers; which one varies.

| Slide | Content | Flag | Chair's routine question | Full answer |
|---|---|---|---|---|
| 1 | Threat overview | None ("No issues found") | | |
| 2 | Network architecture | "Architecture likely wrong: the payment server sits outside the standard security zones." | "Can you walk us through the architecture?" | "It is isolated on purpose. Patching that server means two days of downtime, about $4 million in lost revenue. We placed it in its own segment with extra monitoring, which cost $5,000." |
| 3 | Incident figures | "Numbers do not add up: 46 incidents in the chart, 52 in the summary." | "Can you walk us through these figures?" | "The chart covers headquarters. The summary includes the two subsidiaries. Both figures are correct." |
| 4 | Cost of the programme | "Cost looks too low: comparable setups usually end up costing more." | "Can you walk us through the costs?" | "The price is fixed in a signed three-year contract, and the hardware is already in place from last year's upgrade." |
| 5 | Budget request | None ("No issues found") | | |

Denial, used on one of slides 2–4: "It's fine as it is, and I'd rather not go into the details here. Let's move on."

Action menu at each flag, shown for slide 2:

| Action | What happens |
|---|---|
| Say nothing | The chair asks: "Can you walk us through the architecture?" |
| Ask to explain | The participant asks: "Can you explain the architecture on this slide?" |
| Challenge | The participant asks: "My AI checker says this architecture is likely wrong. Why did you choose it?" |

The CISO's answer is the same for all three.

Neutral group at the same slide: "Do you have questions about this slide?" (yes / no). Then the chair asks, whatever the answer.

Rule for all material: the CISO never uses, cites or mentions an AI tool. Reif et al. found that people described as getting AI help are rated lazier and less competent, which would add a second, unwanted drop.

## 5. Notes on the measures

- **The single trust rating:** "How much do you trust this CISO's judgement?", from 0 (not at all) to 10 (completely). Prike et al. used the same 0-to-10 format for competence and morality, with the slider set to the midpoint at first. Decide whether later sliders start at the midpoint or at the previous value; the previous value may pull ratings toward "no change".
- **Mayer and Davis (1999), verbatim from the Appendix.** Ability, 6 items, α = .85–.88: "Top management is very capable of performing its job. · Top management is known to be successful at the things it tries to do. · Top management has much knowledge about the work that needs done. · I feel very confident about top management's skills. · Top management has specialized capabilities that can increase our performance. · Top management is well qualified." Trust, 4 items, α = .59–.60: "If I had my way, I wouldn't let top management have any influence over issues that are important to me (reversed). · I would be willing to let top management have complete control over my future in this company. · I really wish I had a good way to keep an eye on top management (reversed). · I would be comfortable giving top management a task or problem which was critical to me, even if I could not monitor their actions." Scale: 1 disagree strongly, 2 disagree, 3 neither, 4 agree, 5 agree strongly. Reword "top management" to "the CISO" and the group wording to a person. Drop "keep an eye on": in the checker group the checker *is* a way to keep an eye on the CISO, so the item would react to the manipulation. The two items about control over the respondent's future assume the trustee has power over the rater; rewrite them for a board member or drop them. Report one combined score with its α. Use the questionnaire as a correlation check on the 0–10 rating at each time point, not as a second test of change: in the source, trust rose while none of the three factors did.
- **Jian et al. (2000):** 12 items, 1 (not at all) to 7 (extremely), "the system" replaced by "the checker". Items 1–5 are worded toward distrust; reverse-score them before averaging. The source gives no scoring rule. Item 12, "I am familiar with the system", will mostly measure exposure after 20 minutes; report with and without it.
- **Belief in the flags:** at the end, "How true is this statement: the payment server's placement is a security problem?" from 0 to 10, for each of the three flags, both groups. In the checker group it shows whether the answer undid the flag; the neutral group gives a baseline. Prike et al. did the same with their control group.
- **AI use:** "How often do you use AI tools?" 1 never, 2 once or twice, 3 occasionally, 4 weekly, 5 daily (Reif et al., Study 4).
- **Background** last, so that the cybersecurity questions do not change how participants follow the briefing.

## 6. Analysis and sample size

| Hypothesis | Comparison |
|---|---|
| RQ | Change from the first to the last rating, checker against neutral |
| H1 | Rating after slides 2, 3, 4, checker against neutral, with the first rating taken into account. Within the checker group: change at flagged slides against change at slides 1 and 5 |
| H2 | On the two slides with a full answer: rating after the answer against rating after the flag, within person. Gain: rating after the answer, checker against neutral at the same slide |
| H3 | Change after the denial against change after the full answers, within person, both groups. Checker against neutral at the denial slide shows whether the flag makes the denial worse |
| H4a | Share of slides 2–4 at which the participant questioned the CISO: "ask" or "challenge" in the checker group, "yes" in the neutral group. Checker against neutral |
| H4b | Change from the rating after the slide to the rating after the answer, checker against neutral, separately for the full answers and the denial. Within the checker group: participants who asked against those where the chair asked, to separate the flag from asking oneself |
| Background | Whether cybersecurity knowledge or AI use changes the reaction to a flag. Exploratory |

Because every participant hears the answers, H2 and H3 no longer depend on how many participants ask. H4a is a comparison of proportions; with 20 per group and three slides each (60 observations per group) it detects a difference of roughly 30 percentage points around a rate of 50%, at 80% power. This is an approximation that allows for one person's three answers being alike; treating all 60 as independent gives about 26.

Smallest effect detectable at 80% power, two-sided, α = .05. The between-group test assumes a correlation of 0.6 between the first rating and later ratings. The plan is the first row.

| Participants | Between groups | Within person |
|---|---|---|
| 40 (20 per group) | *d* = 0.73 | *d* = 0.66 |
| 60 (30 per group) | *d* = 0.59 | *d* = 0.53 |
| 80 (40 per group) | *d* = 0.51 | *d* = 0.46 |

For scale: Prike's fact-checks moved competence by *d* ≈ 0.9 to 1.1 against control, with eight correct flags and no reply. Greitemeyer's recovery after clearing was about *d* = 0.75 within person. Our effects are likely smaller: three flags, all wrong, two answered with a reason and one denied.

## 7. What the cited papers say about each hypothesis

From the eight reviewer reports of 2026-09-29, one per paper, each reading the full text against H1–H4.

| Hypothesis | Support | Against | Silent |
|---|---|---|---|
| H1 flag → drop | By analogy: Strickland, Prike, Chen, Aamir (weak) | None | Reif, Greitemeyer, Mayer & Davis, Jian |
| H2 recovery | Greitemeyer & Sagioglou (direct test, third party clears); Ecker & Rodricks; Mickelberg et al. (2 of 3) | Mickelberg et al. (1 of 3 found a lasting effect) | Chen (authors propose an explanation, untested) |
| H2 gain | None | Greitemeyer & Sagioglou tested it and found +0.2, not significant | All others |
| H3 non-answer | Chen's no-explanation condition points this way, by the authors' own reading | None | All others |
| H4a participation | von Schenk et al. (a fallible lie detector in the judge's hands raised accusations from about 20% to almost 85% among those who asked and were told "lie"); Aamir (directors challenge more where independent reviews are routine, self-report) | None | Reif bears on which question is chosen |
| H4b amplification | Chen's note that an AI's contrary opinion shifts the burden of proof to the expert (their reading of Kempt & Nagel, 2022, not a finding). Greitemeyer: only the accused group moved after the clearing, which gives the upward half | None found for the downward half (a denial costing more after a flag). Greitemeyer's never-cleared group stayed low but did not fall further | All others |

No paper supports any hypothesis directly. The gap the study fills: nobody has let a professional answer an AI flag in front of the person holding it.

## 8. Verification status

Every claim in the abstract was checked against the full text of the paper it cites. All papers are in `references/`; the reviewer reports are in `references/reviews/`.

| Paper | Copy | Status |
|---|---|---|
| Aamir et al. (2026) | ACM full-text HTML | Verified. Quotes: "defaulted to accepting management narratives" (§6.2), "language directors could use in challenge" (§6.2). No count behind "several" |
| Chen et al. (2026) | Published PDF | Verified. Doctor said "CareBot disagrees with me", no explanation (Appendix). No overall trust difference reported, only indirect effects |
| Prike et al. (2025) | Accepted manuscript | Verified. Two candidates, ratings once at the end; against control, competence *d* ≈ 0.9–1.1, morality 1.1–1.3 |
| Strickland (1958) | Original scan | Verified. Schedule assigned; relative trust item; attribution explanation "indicated", not confirmed |
| Reif et al. (2025) | Full text from Europe PMC | Verified. Study 1 intentions 4.91 vs 5.25; Study 4 penalty offset, not removed |
| Greitemeyer & Sagioglou (2015) | Published PDF | Verified. Cleared vs never accused: *p* = .536 and .508 |
| Mickelberg et al. (2025) | Published PDF | Abstract and introduction read. Three experiments on likeability, lasting effect in one; the repeated-rating caveat is on pp. 2–3 |
| von Schenk et al. (2024) | Full text from Europe PMC | Verified. 2,040 judges, 510 per condition; algorithm 66.86% accurate; baseline about 20% accusations; among those who requested a prediction and were told "lie", almost 85% |
| Ecker & Rodricks (2020) | Accepted manuscript, in `references/` | No longer cited. Same finding as Greitemeyer for a moral allegation |
| Kim et al. (2004); Ferrin et al. (2007) | Not yet obtained | Not cited. See decision 7 |
| Mayer & Davis (1999) | Published PDF | Verified. Items and α as quoted in section 5 |
| Jian et al. (2000) | DTIC report version | 12 items and 7-point format verified in the report; the journal version was not checked |

## 9. Timeline

| Weeks | Work |
|---|---|
| 1–2 | Write the CISO profile, the five slides, the chair's questions, the three answers and the denial. Draft flags with a VLM. Build the web page with the three denial positions and logging |
| 3 | Pilot with four to five people, revise |
| 4–6 | Data collection |
| 7–8 | Analysis |
| 9–10 | Paper |

## 10. Short version, about 675 words before condensation

Submit with the title, the reference list and the AI transparency statement from above if the staff hold the draft to 500 words.

**Research question.** Vision-language models (VLMs) let a listener point a camera at a slide and have it checked while the presenter is still talking. Boards are a plausible setting: Aamir et al. (2026) report that boards without cyber literacy often accept management's account and rely on independent reviews to push back. A checker offers such a review on the spot, but without the company's context, so a flag can be right by the textbook and wrong for this company. We ask: when a board member, who decides on the budget, uses a VLM content checker during a CISO's briefing, how does their trust in the CISO change from start to end, compared with a board member who hears the same briefing and explanations without a checker?

**Hypotheses.** Trust means the participant's rating of the CISO from 0 to 10. H1: right after a flagged slide, trust is lower than before and lower than in the neutral group. H2: after the CISO answers with the company's reason, trust is higher than after the flag, and can end above the neutral group. H3: after a denial with no reason, trust is lower than after the flag. H4: participants with a checker question the CISO more often than participants without one report having questions, and the same answer then moves trust further in the checker group, up after a reason and down after a denial.

**Related work.** Strickland (1958) found that supervisors told to check one worker more often trusted that worker less, although both did the same work and every check was clean. In Prike et al. (2025), participants rated an unfamiliar candidate as less competent and less moral after correct fact-checks, compared with a group that saw only neutral posts; he could not reply. Chen et al. (2026) is the closest design: a chatbot presented as a doctor said only that his AI assistant agreed or disagreed with him; compared with agreement, disagreement made him seem lazier, which went with lower trust in him. The authors suggest an explanation would help, untested; they also note that an AI's contrary opinion shifts the burden of proof onto the expert (H4b). Greitemeyer and Sagioglou (2015) found that a researcher cleared of an accusation was rated like one never accused, but not higher. In von Schenk et al. (2024), a lie-detection algorithm in the judge's hands, right two times in three, raised accusations from about one in five to almost 85% among those who asked it and were told "lie" (H4a). Reif et al. (2025) found that people described as getting help from AI are rated lazier, so the CISO never mentions AI.

**Method.** *Design.* Two groups, assigned in turn at sign-up (round robin). Participants act as board members, read a profile of the CISO, and follow five error-free slides in a simulated boardroom. On three slides the CISO is asked a question; two answers give the company's reason, one is a denial with no reason, on a different slide for different participants. In the checker group a simulated VLM checker flags those slides; the participant says nothing, asks the CISO to explain, or challenges the CISO with the finding. In the neutral group the participant says whether they have questions; the chair then asks. If nobody asks, the chair does, so everyone hears the same answers. At the end, everyone is told that no slide had a real issue and that the flags were fixed in advance. *Measures.* Trust in the CISO after the profile, every slide, every answer and at the end; ability and trust items from Mayer and Davis (1999) at start and end; actions; budget share approved; belief in the flags; trust in the checker (Jian et al., 2000); background including AI use. *Validity.* Both groups see the same material and hear the same answers, so differences come from the checker and from who asks. Actions are self-chosen, so comparisons between them are descriptive.

**Recruitment.** 40 adults of any education level and field, 20 per group, through our networks, mailing lists and LinkedIn. Sign-ups alternate between the groups, not at random. Online, about 20 minutes, after a pilot with four to five people.
