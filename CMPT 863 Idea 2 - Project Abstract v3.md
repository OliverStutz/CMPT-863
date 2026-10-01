# Does a Vision-LLM Content Checker in a Board Member's Hands Erode or Build Trust in the Presenting CISO?

## Research question

Vision-language models (VLMs) let a listener point a camera at a slide and have it checked while the presenter is still talking. Boards are a likely setting: Aamir et al. (2026) found that gaps in directors' knowledge and confidence limit their ability to interrogate cyber-risk reports. But a checker lacks the company's context, so a flag can be right by the textbook and wrong for this company. The relationship it enters is unequal: the director holds the formal power, the CISO the expertise.

> **RQ.** When a board member uses a VLM content checker during a CISO's briefing, how does their trust in the CISO change from the start to the end of the briefing, compared with a board member who follows the same briefing without a checker?

Research on AI fact-checking measures belief in the checked claim. We measure trust in the checked person, step by step.

## Hypotheses

Trust means the participant's rating of the CISO/Presenter on a scale from 0 to 10.

- **H1:** The trust between the participant and the CISO drops when a simulated VLM claim is displayed.
- **H2:** The trust between the participant and the CISO can recover and increase if the simulated VLM claim is followed up with a proposed LLM question which has a positive response which reinforces the context knowledge of the CISO
- **H3:** The trust between the participant and the CISO can decrease heavily if the simulated VLM claim is responded with negation or non-answers #1
- **H4**: The trust change of participant and CISO is heavily dependent on the participants actions and the CISO's responses.

## Related work

In Strickland (1958), participants who monitored one of two subordinates more often came to trust that one less. The monitor there was a supervisor. Ours is a machine held by the person in authority. This informs H1.

In Prike et al. (2025), participants met an unfamiliar political candidate through a profile and rated the candidate. After fact-checks they rated the candidate as less competent and less moral (*N* = 406). We borrow the set-up, an unfamiliar person who is introduced by a profile and then rated. We rate trust, our flags are wrong for this company, and the CISO can reply. This informs H1.

Chen et al. (2026) is the closest design (*N* = 135). When an AI second opinion disagreed with a simulated doctor, participants who saw the doctor as human-like perceived more uncertainty and laziness, which carried over to trust. There the patient depends on the expert, and the doctor's reply was not tested. We reverse the power relation and let the expert answer, which is what H2 tests.

Reif et al. (2025) found that people who use AI at work are rated lazier and less competent and disclose it less willingly. We log whether participants name the checker when they question the CISO.

## Method

### Study design

Two groups, assigned at random. Participants read a scenario in which they sit on the board and decide on the CISO's budget request. They read a profile that establishes the CISO's experience and give a first trust rating. Then the CISO presents five slides in a web-based simulated boardroom. The slides contain no errors and are the same in both groups.

- **Checker group.** On three slides a checker panel shows the camera image of the slide with a marked region and a flag, for example "Architecture likely wrong" or "Numbers do not add up". The participant rates trust and then picks an action: say nothing, ask the CISO to explain, or challenge the CISO with the checker's finding. After a question the CISO gives a pre-defined answer, and the participant rates trust again.
- **Neutral group.** No checker and no questions. The participant rates trust after every slide.

The checker is simulated. Flags and questions are drafted with a real VLM from images of the slides and then fixed, so every participant sees the same output. Each flag is wrong for this company, and each answer explains why.

### Measures

- **Trust in the CISO** (primary): one rating from 0 to 10, after the profile, after every slide, after every answer, and at the end.
- **Trust questionnaire** at the start and the end, adapted from Mayer and Davis (1999), to check that the single rating measures trust.
- **Actions:** the system logs the action chosen at each flag.
- **Budget decision:** share of the requested budget approved.
- **Trust in the checker** (Jian et al., 2000), checker group only.
- **Background:** field of study or work, and self-rated knowledge of cybersecurity.

### Validity

Both groups see the same profile and slides, and the slides contain no errors, so a difference between the groups comes from the checker and not from the presenter. The first rating gives every participant their own starting point, which shows whether trust improved. The ratings after each flag and each answer show which step moves trust. The recorded background shows whether knowledge of cybersecurity changes the reaction to a flag. Limits: participants choose their actions, so a comparison of those who asked with those who said nothing describes them and does not prove a cause. Students stand in for directors.

## Recruitment

We will recruit 60 graduate and upper-year undergraduate students at SFU, 30 per group, through course mailing lists, department channels and word of mouth. All fields are welcome, and we record each participant's background. The study runs online in about 20 minutes. Thirty per group detects a difference of *d* ≈ 0.6 at 80% power. The simulation is a web page with fixed flags and answers, which keeps the build within two weeks. We pilot with four to five people first.

## References

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–22). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791142

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, *213*, Article 103824. https://doi.org/10.1016/j.ijhcs.2026.103824

Jian, J.-Y., Bisantz, A. M., & Drury, C. G. (2000). Foundations for an empirically determined scale of trust in automated systems. *International Journal of Cognitive Ergonomics*, *4*(1), 53–71. https://doi.org/10.1207/S15327566IJCE0401_04

Mayer, R. C., & Davis, J. H. (1999). The effect of the performance appraisal system on trust for management: A field quasi-experiment. *Journal of Applied Psychology*, *84*(1), 123–136. https://doi.org/10.1037/0021-9010.84.1.123

Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, *46*(5), 1081–1100. https://doi.org/10.1111/pops.13059

Reif, J. A., Larrick, R. P., & Soll, J. B. (2025). Evidence of a social evaluation penalty for using AI. *Proceedings of the National Academy of Sciences*, *122*(19), Article e2426766122. https://doi.org/10.1073/pnas.2426766122

Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, *26*(2), 200–215. https://doi.org/10.1111/j.1467-6494.1958.tb01580.x

## AI transparency

We used AI to search for and screen related papers, to check our claims against the sources, and to draft this text. The team reads all cited papers before the final submission.

---

# Working notes (delete before submission)

Version 3, written 2026-09-29 from your five points. Version 2 is [[CMPT 863 Idea 2 - Project Abstract v2]] and is unchanged. Submission is due 2026-09-30, 11:59 pm.

## 1. What changed from version 2

| Topic | Version 2 | Version 3 |
|---|---|---|
| References in the hypotheses | Four citations | None. The reasons are given in plain words, and Related work names the hypothesis each paper informs |
| Reference list | 10 entries | 7 entries: your four papers, the seed paper and two scales. Greitemeyer and Sagioglou, Thorson and Kelley are out, because they were cited only in the hypotheses |
| Groups | Three conditions: no checker, flag resolved, flag deflected | Two groups: checker and neutral |
| Start of the study | Scenario only | Scenario and a profile that gives the CISO credibility, followed by the first trust rating |
| Briefing | One recorded five-minute talk, one flag | Five slides, three of them flagged |
| Participant's role | Asks the one suggested question | Picks an action at each flag: say nothing, ask to explain, or challenge |
| CISO's answers | Resolving or deflecting, set by the condition | Pre-defined and resolving. What varies is whether the participant asks |
| Trust measurement | Three times | After the profile, after every slide, after every answer, and at the end |
| Main outcome | Trust after the answer, compared between conditions | Change in trust from start to end, compared between groups |
| Participants | 90, from outside computer security | 60, all fields, background recorded |

## 2. Decisions that are yours to confirm

1. **Are the flags wrong, or does the CISO make real mistakes?** Your examples ("numbers don't add up") can be read both ways. I kept the slides free of errors, so every flag is wrong for this company and every answer explains why. This keeps the separation between the presenter's mistakes and the AI's intervention that Xueying asked for. The alternative is a mix of wrong and correct flags, which is closer to a real checker but needs a second kind of answer, in which the CISO admits a mistake.
2. **All answers resolve the flag.** Version 2 had a condition in which the CISO deflects. With two groups there is no place for it. The friction side of your idea now shows in three places: the drop at each flag, the participants who say nothing, and the participants who pick the challenge. If you want weak answers back, the simplest way is to make one of the three answers a deflecting one.
3. **Two kinds of question.** You asked for pre-defined questions. I made them two per flag: a neutral one that asks the CISO to explain, and a challenge that names the checker's finding. Both lead to the same answer. Which one a participant picks is itself a result, and it is the link to Reif et al.
4. **One rating from 0 to 10 for trust.** This follows your point about Prike et al. and keeps the repeated rating short. The questionnaire at the start and the end is my addition, to show that the single rating measures trust. Drop it if the session gets too long.
5. **Kept from version 2 without your asking:** the budget decision and the trust-in-the-checker scale. Both are asked once at the end. Say so if you want them out.
6. **H3 predicts a gain.** Prior work found that a cleared person is rated at best as before (section 7). The hypothesis is yours to set, and the study is informative whichever way it comes out.
7. **Slides with recorded explanation, or a video of the CISO?** I wrote "the CISO presents five slides" and left the form open. Slides with recorded audio are much faster to build and to change after the pilot than video.
8. **Participants are told the checker is a live AI.** The flags are fixed in advance, which is mild deception. It needs a debrief at the end and should be checked against the course's ethics rules.

## 3. Session flow, about 20 minutes

| Step | Checker group | Neutral group |
|---|---|---|
| 1 | Consent, scenario, introduction of the checker | Consent, scenario |
| 2 | CISO profile | CISO profile |
| 3 | First trust rating and questionnaire | Same |
| 4 | Slide 1, then trust rating | Same |
| 5 | Slide 2 with flag, then trust rating, action menu, answer if asked, trust rating | Slide 2, then trust rating |
| 6 | Slide 3 with flag, same steps | Slide 3, then trust rating |
| 7 | Slide 4 with flag, same steps | Slide 4, then trust rating |
| 8 | Slide 5, then trust rating | Same |
| 9 | Last trust rating and questionnaire, budget decision | Same |
| 10 | Trust in the checker | Not asked |
| 11 | Background questions, debrief | Same |

The rating directly after each slide is taken in both groups at the same point, so the groups can be compared slide by slide.

## 4. Stimulus sketch

Three flags, built from your examples. The wording is a starting point for the team.

| Slide | Content | Flag | The CISO's answer |
|---|---|---|---|
| 1 | Threat overview | None | |
| 2 | Network architecture | "Architecture likely wrong: the payment server sits outside the standard security zones." | "It is isolated on purpose. Patching that server means two days of downtime, about $4 million in lost revenue. We placed it in its own segment with extra monitoring, which cost $5,000." |
| 3 | Incident figures | "Numbers do not add up: 46 incidents in the chart, 52 in the summary." | "The chart covers headquarters. The summary includes the two subsidiaries. Both figures are correct." |
| 4 | Cost of the programme | "Cost looks too low: comparable setups usually end up costing more." | "The price is fixed in a signed three-year contract, and the hardware is already in place from last year's upgrade." |
| 5 | Budget request | None | |

Action menu at each flag, shown here for slide 2:

| Action | Text |
|---|---|
| Say nothing | The briefing continues |
| Ask to explain | "Can you explain the architecture on this slide?" |
| Challenge | "My AI checker says this architecture is likely wrong. Why did you choose it?" |

## 5. Notes on the measures

- **The single trust rating** could read: "How much do you trust this CISO's judgement?", from 0 (not at all) to 10 (completely). Prike et al. used the same kind of scale, from 0 to 10, for competence and morality.
- **Mayer and Davis (1999)** has 6 ability, 5 benevolence and 6 integrity items on a 5-point agree/disagree scale. The items refer to "top management" and have to be reworded for "the CISO". A selection of six items is enough for the check.
- **Jian et al. (2000)** has 12 items. The 7-point format was confirmed through secondary sources only.
- **Background** should be asked at the end, so that the questions about cybersecurity knowledge do not change how participants follow the briefing.

## 6. Analysis and sample size

| Hypothesis | Comparison |
|---|---|
| H1 | Rating directly after slides 2, 3 and 4, checker group against neutral group, with the first rating taken into account |
| H2 | In the checker group: rating after the answer against rating after the flag, for the participants who asked. For those who said nothing, the ratings are reported as they are |
| H3 | Change from the first to the last rating, checker group against neutral group |
| Background | Whether participants with cybersecurity knowledge react less to a flag. Exploratory |

Smallest effect that can be detected at 80% power, two-sided, alpha = .05. The group comparison assumes a correlation of 0.6 between the first rating and later ratings.

| Participants | H1 and H3, between groups | H2, within the checker group |
|---|---|---|
| 40 (20 per group) | d = 0.73 | d = 0.66, if all 20 ask |
| 50 (25 per group) | d = 0.65 | d = 0.58, if all 25 ask |
| 60 (30 per group) | d = 0.59 | d = 0.53, if all 30 ask |
| 80 (40 per group) | d = 0.51 | |

H2 depends on how many participants ask. If half of the checker group says nothing, 15 participants are left for the comparison and only a large effect (d = 0.78) can be detected.

## 7. What prior work says about the hypotheses

The hypotheses carry no references. This is what the two investigations of 2026-09-29 found, so that the team knows where each one stands.

| Hypothesis | Prior work | Support |
|---|---|---|
| H1 | Prike et al. (2025): fact-checked candidate rated less competent and less moral. Chen et al. (2026): doctor seen as more uncertain when an AI disagreed. Strickland (1958): the more closely monitored person was trusted less | Supported by analogy. No study used an AI flag that is wrong |
| H2 | Greitemeyer and Sagioglou (2015): a researcher who was accused and then cleared was rated as credible as one never accused | Supported by analogy |
| H3 | The same study found recovery but no gain. Thorson (2016): a politician was still rated lower after a false claim had been corrected | Not supported. The gain is our own prediction |

No study from 2023 to 2026 was found in which observers rate a professional who rebuts an AI flag.

A marker who asks for grounding of the hypotheses in prior work can be pointed to Related work, where each paper names the hypothesis it informs. The rubric's top score for the hypothesis mentions this grounding, so removing the citations carries a small risk.

## 8. To verify before the final submission

| Claim in the abstract | Status | Where to check |
|---|---|---|
| Aamir et al.: gaps in knowledge and confidence limit the ability to interrogate risk reports | Confirmed from the abstract. The full text could not be downloaded automatically | Open access: https://ro.ecu.edu.au/ecuworks2022-2026/8241 |
| Prike et al.: profile, ratings of competence and morality from 0 to 10, *N* = 406 | Confirmed by me in the accepted manuscript | https://api.research-repository.uwa.edu.au/ws/files/463479508/Prike.2024.PoPs_pp.pdf |
| Chen et al.: the doctor's reply was not tested | Not confirmed. Nothing in the abstract or press release describes a reply | Full text through the SFU library |
| Chen et al.: the effect "carried over to trust" | Direction not seen in the sources read | Full text through the SFU library |
| Strickland: the more closely monitored subordinate was trusted less | Confirmed through a secondary source only | You have listened to the original, so check the wording against it |
| Reif et al.: AI users disclose less willingly | Read by the agent in the full text. The difference is small (4.91 against 5.25 on willingness to disclose to a manager) | https://www.pnas.org/doi/pdf/10.1073/pnas.2426766122 |

## 9. Timeline

| Weeks | Work |
|---|---|
| 1–2 | Write the CISO profile and the five slides, draft flags and questions with a VLM, write and record the answers, build the web page with logging |
| 3 | Pilot with four to five people, revise |
| 4–6 | Data collection |
| 7–8 | Analysis |
| 9–10 | Paper |

## 10. Short version, about 500 words

The brief asks for approximately 500 words. If the staff hold the draft to that, submit this text with the title, reference list and AI transparency statement from above. It makes the same claims as the long version.

**Research question.** Vision-language models (VLMs) let a listener point a camera at a slide and have it checked while the presenter is still talking. Boards are a likely setting: Aamir et al. (2026) found that gaps in directors' knowledge and confidence limit their ability to interrogate cyber-risk reports. But a checker lacks the company's context, so a flag can be right by the textbook and wrong for this company. We ask: when a board member, who holds the formal power, uses a VLM content checker during the briefing of a CISO, who holds the expertise, how does their trust in the CISO change from start to end, compared with a board member without a checker?

**Hypotheses.** Trust means the participant's rating of the CISO from 0 to 10. H1: at the slides with a flag, trust drops in the checker group and not in the neutral group, because a non-expert cannot tell whether a flag is right. H2: when a participant asks a question and hears the answer, trust is higher than right after the flag. When the participant says nothing, it stays low. H3: from the first rating to the last, trust rises more in the checker group, because each answer lets the CISO show knowledge that the neutral group never sees.

**Related work.** Strickland (1958) found that supervisors came to trust a more closely monitored subordinate less. In Prike et al. (2025), participants rated an unfamiliar political candidate as less competent and less moral after fact-checks, but the claims were false and the candidate could not reply. Chen et al. (2026) is the closest design: an AI second opinion that disagreed with a simulated doctor made the doctor seem more uncertain and lazier to participants who saw the doctor as human-like. The doctor's reply was not tested. We reverse the power relation and let the expert answer. Reif et al. (2025) found that people who use AI at work are rated lazier and less competent and disclose it less willingly, so we log whether participants name the checker when asking.

**Method.** *Design.* Two groups, assigned at random. Participants act as board members, read a profile of the CISO, and follow five error-free slides in a simulated boardroom. In the checker group, a simulated VLM checker flags three slides. At each flag the participant says nothing, asks the CISO to explain, or challenges the CISO with the finding, and hears a pre-defined answer. The neutral group sees the same slides without a checker. *Measures.* Trust in the CISO, rated after the profile, after every slide, after every answer and at the end; a trust questionnaire at start and end (Mayer & Davis, 1999); the actions chosen; the share of the budget approved; trust in the checker (Jian et al., 2000); the participant's background. *Validity.* Both groups see the same material, so differences come from the checker. The first rating gives each participant a starting point. Participants choose their actions, so comparisons between them are descriptive.

**Recruitment.** We will recruit 60 SFU graduate and upper-year undergraduate students from all fields, 30 per group, through course mailing lists, department channels and word of mouth. The study runs online in about 20 minutes, after a pilot with four to five people.
