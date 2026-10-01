# Does a Vision-LLM Content Checker in a Board Member's Hands Erode or Build Trust in the Presenting CISO?

## Research question

Vision-language models (VLMs) let a listener point a camera at a slide and have it checked while the presenter is still talking. Boards are a likely setting: Aamir et al. (2026) found that gaps in directors' knowledge and confidence limit their ability to interrogate cyber-risk reports. But a checker lacks the company's context, so a flag can be right by the textbook and wrong for this company. The relationship it enters is unequal: the director or board member holds the formal power, the CISO the expertise.

> **Research Question:** When a board member challenges a CISO with a question prompted by a VLM checker, does trust in the CISO fall, and does the CISO's answer restore it or raise it above the level reached without a checker?

Research on AI fact-checking measures belief in the checked claim. We measure trust in the checked person.

## Hypotheses

Trust means the participant's rating of the CISO on the trust questionnaire (see Measures).

- **H1 (flag penalty).** Right after the flagged slide, before the CISO answers, participants with the checker rate their trust in the CISO lower than participants without it, although the briefing contains no error. Participants cannot tell that the flag is wrong for this company, so it should act like a correct one. When a fact-checking organisation marked claims in a political candidate's favour as false, participants rated the candidate as less competent and less moral (Prike et al., 2025). We expect a smaller drop, because that study used eight fact-checks and we use one flag.
- **H2 (recovery).** After the answer, participants who heard the resolving answer rate their trust in the CISO higher than participants who heard the deflecting answer. A resolving answer removes the doubt: a researcher who was accused and then cleared was rated as credible as one never accused (Greitemeyer & Sagioglou, 2015). A deflecting answer leaves the flag standing.
- **H3 (gain).** After the resolving answer, participants with the checker rate their trust in the CISO higher than participants who heard the same explanation without a checker. Earlier studies do not support this: a cleared person was rated at best as before, and a politician was still rated lower after a false claim about the politician had been corrected (Thorson, 2016). But there someone else did the clearing. Here the CISO removes the doubt in person, and people give more credit for an outcome achieved against an obstacle (Kelley, 1973).

## Related work

In Strickland (1958), participants who monitored one of two subordinates more often came to trust that one less. The monitor there was a supervisor. Ours is a machine held by the person in authority.

In Prike et al. (2025), fact-checks made participants rate a political candidate as less competent and less moral (*N* = 406). The fact-checked claims there were false, and the candidate could not reply. Our flag is wrong for this company, and the CISO replies.

Chen et al. (2026) is the closest design (*N* = 135). When an AI second opinion disagreed with a simulated doctor, participants who saw the doctor as human-like perceived more uncertainty and laziness, which carried over to trust. There the patient depends on the expert, and the doctor's reply was not tested. We reverse the power relation and let the expert answer.

Reif et al. (2025) found that people who use AI at work are rated lazier and less competent and disclose it less willingly. We log whether participants name the checker as the source of their question.

## Method

### Study design

Between-subjects, three conditions. Participants read a scenario in which they sit on the board and decide on the CISO's budget request, then watch one recorded five-minute briefing. It contains no errors and is identical in all conditions.

1. **No checker.** The chair asks a routine question. The CISO explains why a patch was deferred: a $5,000 compensating control instead of $4 million of downtime.
2. **Flag, resolved.** A checker panel flags the deferred patch and suggests a question. The participant asks it and gets the same explanation.
3. **Flag, deflected.** Same flag and question, but the CISO deflects.

The checker is simulated. Its output is drafted with a real VLM from a photo of the slide and then frozen, so every participant sees the same flag and we control its accuracy.

### Measures

- **Trust in the presenter** (primary): ability, benevolence and integrity items adapted from Mayer and Davis (1999), before the flag (T0), after it (T1) and after the answer (T2).
- **Behavioural trust:** share of the requested budget approved.
- **Judgement of the flagged decision** as sound or not, which shows whether the checker misled.
- **Trust in the checker** (Jian et al., 2000).
- **Logs:** opening the flag, naming the checker when asking.

### Validity

The briefing is constant, so no difference can come from errors by the presenter. Conditions 1 and 2 share the footage of the explanation, so a difference in final trust comes from the challenge, not the information. Conditions 2 and 3 differ only in the answer. The design is between-subjects because a participant who has seen a flag refuted reads later flags differently. Limits: students stand in for directors, one presenter, one topic.

## Recruitment

We will recruit 90 graduate and upper-year undergraduate students at SFU from outside computer security, 30 per condition, through course mailing lists, department channels and word of mouth. Like directors, they judge a technical report without being experts in it. The study runs online in about 15 minutes. Thirty per condition detects effects of *d* ≈ 0.5 to 0.6 at 80% power. One script, one presenter and a scripted flag keep the build within two weeks. We pilot with four to five people first.

## References

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–22). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791142

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, *213*, Article 103824. https://doi.org/10.1016/j.ijhcs.2026.103824

Greitemeyer, T., & Sagioglou, C. (2015). Does exonerating an accused researcher restore the researcher's credibility? *PLOS ONE*, *10*(5), Article e0126316. https://doi.org/10.1371/journal.pone.0126316

Jian, J.-Y., Bisantz, A. M., & Drury, C. G. (2000). Foundations for an empirically determined scale of trust in automated systems. *International Journal of Cognitive Ergonomics*, *4*(1), 53–71. https://doi.org/10.1207/S15327566IJCE0401_04

Kelley, H. H. (1973). The processes of causal attribution. *American Psychologist*, *28*(2), 107–128. https://doi.org/10.1037/h0034225

Mayer, R. C., & Davis, J. H. (1999). The effect of the performance appraisal system on trust for management: A field quasi-experiment. *Journal of Applied Psychology*, *84*(1), 123–136. https://doi.org/10.1037/0021-9010.84.1.123

Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, *46*(5), 1081–1100. https://doi.org/10.1111/pops.13059

Reif, J. A., Larrick, R. P., & Soll, J. B. (2025). Evidence of a social evaluation penalty for using AI. *Proceedings of the National Academy of Sciences*, *122*(19), Article e2426766122. https://doi.org/10.1073/pnas.2426766122

Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, *26*(2), 200–215. https://doi.org/10.1111/j.1467-6494.1958.tb01580.x

Thorson, E. (2016). Belief echoes: The persistent effects of corrected misinformation. *Political Communication*, *33*(3), 460–480. https://doi.org/10.1080/10584609.2015.1102187

## AI transparency

We used AI to search for and screen related papers, to check our claims against the sources, and to draft this text. The team reads all cited papers before the final submission.

---

# Working notes (delete before submission)

Version 2, written 2026-09-29. Version 1 is [[CMPT 863 Idea 2 - Project Abstract]]. Submission is due 2026-09-30, 11:59 pm.

## 1. What changed from version 1

| Topic | Version 1 | Version 2 |
|---|---|---|
| Question | Does a flag cost the presenter trust beyond what their errors justify? | Does an AI-prompted challenge lower trust in the CISO, and does the CISO's answer restore or raise it? |
| Setting | Generic audience and presenter | Board member (formal power) and CISO (expertise) |
| Presenter errors | Two planted errors per talk | None. The briefing is sound, so every trust change comes from the checker and the answer |
| Flag | Correct flags plus one false positive | One context-blind flag: right by the textbook, wrong for this company |
| Presenter's response | Not possible, listed as the main limit | The manipulation: resolved or deflected |
| Checker | Scripted overlay | Simulated VLM: output drafted with a real VLM from a photo of the slide, then frozen |
| Design | Within-subjects, three talks, three presenters, N = 30 | Between-subjects, one talk, one presenter, N = 90 |
| Trust measurement | Once per talk | Three times in one talk (T0, T1, T2) |
| Papers | Strickland, Arkes, Chen, von Schenk, Prike, Chae and Tewksbury | Your four pre-selected papers and the seed paper, plus three new ones for H2 and H3 |
| Length | 1,345 words | 929 words, and a short version of 629 words in section 11. Counts cover the text from Research question to Recruitment, without title, references and AI statement |

## 2. Decisions that are yours to confirm

1. **Between-subjects with 90 participants.** This is the largest change. Once a participant has seen the CISO refute a flag, they read every later flag differently, and counterbalancing cannot remove that. The price is three times the participants. Statistical power is about the same in both designs (section 6). If 90 is out of reach, the fallback is the version 1 structure: within-subjects, three talks with three presenters, N = 30, with the first talk analysed separately as a check.
2. **The gain (H3) is not supported by prior work.** The Gemini conversation treats the "trust multiplier" as a likely outcome. The literature we found says recovery to the baseline is the best documented outcome, and a lingering deficit is common (section 7). I kept the gain as a hypothesis because it is your idea and it has a theoretical basis, and the hypothesis itself says that prior work found recovery at best. The study is informative either way: if no gain appears, a checker costs the expert trust even when the expert is right and answers well.
3. **Three new citations that you have not read.** Greitemeyer and Sagioglou (2015), Thorson (2016), Kelley (1973). They are needed only for H2 and H3. Access is in section 8.
4. **Papers removed.** Arkes et al. (2007), von Schenk et al. (2024) and Chae and Tewksbury (2026) are not in your preparation file, so I left them out. Arkes et al. is evidence against H3 (a physician who overruled the decision aid was still rated below one who used none), so it is the first one to bring back.
5. **No planted errors.** Version 1 measured whether participants spotted errors. Version 2 has no errors to spot, so that measure is replaced by the judgement of the flagged decision.
6. **Who asks in the no-checker condition.** I let the chair ask, which matches the situation Aamir et al. describe, where directors find it hard to challenge. The alternative is that the participant asks a routine question taken from the board papers. That would make the participant equally active in all conditions, but it is further from the status quo.
7. **Participants are told the checker is a live AI.** The flag is scripted, which is mild deception. It needs a debrief at the end and should be checked against the course's ethics rules.
8. **The presenter-side study stays out.** The link to Reif et al. is the logged choice of whether the participant names the checker as the source of the question.

## 3. Stimulus sketch

| Element | Content |
|---|---|
| Slide 4 | "Payment server: critical patch deferred to the March maintenance window. Compensating controls in place." |
| Checker panel | Camera image of slide 4 with the patch line marked. "A critical vulnerability is left unpatched. Security guidance calls for prompt patching of critical vulnerabilities." |
| Suggested question | "Why has this server not been patched?" |
| Chair's routine question (condition 1) | "Can you say more about the deferred patch on slide 4?" |
| Resolving answer (conditions 1 and 2, same footage) | "Patching means taking payments offline for about two days, roughly $4 million in lost revenue. We isolated the server on its own network segment and added monitoring, which cost $5,000. The patch is scheduled for the March maintenance window." |
| Deflecting answer (condition 3) | "That is an operational detail. My team has it under control, and I can follow up after the meeting." |

The resolving answer must not mention the checker, otherwise conditions 1 and 2 cannot share the footage.

## 4. Procedure, about 15 minutes

1. Consent and the written scenario. In conditions 2 and 3 the scenario introduces the checker.
2. Briefing, slides 1 to 3. Then T0: six trust items.
3. Briefing, slide 4. In conditions 2 and 3 the flag appears. Then T1: the same six items, and one item on confidence to question the report.
4. Question and answer. In conditions 2 and 3 the participant chooses between two phrasings, with or without naming the checker, and the recorded answer plays.
5. T2: the six items again, the full Mayer and Davis item set, the budget decision, the judgement of the flagged decision, trust in the checker (conditions 2 and 3), and one open question on what drove the rating.
6. Debrief.

## 5. Notes on the measures

- **Mayer and Davis (1999)** has 6 ability, 5 benevolence and 6 integrity items on a 5-point agree/disagree scale. The items refer to "top management" and have to be reworded for "the presenter". Version 1 said 7-point, which was wrong for the original.
- **The six-item short form** (two items per dimension) is our own selection. Check its reliability in the pilot.
- **Jian et al. (2000)** has 12 items. The 7-point format was confirmed through secondary sources only.

## 6. Analysis and sample size

Linear model with T0 as covariate and three planned contrasts:

| Hypothesis | Time | Contrast |
|---|---|---|
| H1 | T1 | Conditions 2 and 3 against condition 1 |
| H2 | T2 | Condition 2 against condition 3 |
| H3 | T2 | Condition 2 against condition 1 |

Smallest effect that can be detected at 80% power, two-sided, alpha = .05. The between-subjects figures assume a correlation of 0.6 between T0 and the later rating. The within-subjects figures assume a correlation of 0.5 between conditions.

| Design | Participants | H1 | H2 and H3 |
|---|---|---|---|
| Between-subjects, 20 per condition | 60 | d = 0.62 | d = 0.73 |
| Between-subjects, 25 per condition | 75 | d = 0.56 | d = 0.65 |
| Between-subjects, 30 per condition | 90 | d = 0.51 | d = 0.59 |
| Between-subjects, 40 per condition | 120 | d = 0.44 | d = 0.51 |
| Within-subjects (fallback) | 30 | d = 0.53 | d = 0.53 |

Prike et al. report large effects (d of about 0.9 to 1.3), but those come from eight fact-checks of false claims. One flag on a sound decision will probably produce a smaller drop, so H1 may be harder to detect than those numbers suggest. Two or three flags in the briefing would strengthen it. H3 is likely to be small if it exists, so this study can detect a clear gain but not a subtle one. That belongs in the limits of the final paper.

## 7. What the two investigations found

Two Opus agents ran on 2026-09-29. Afterwards I checked every reference against Crossref, and the Greitemeyer and Sagioglou result against the full text.

**Evidence on H2 and H3**

| Paper | Finding | Points to |
|---|---|---|
| Greitemeyer and Sagioglou (2015), two studies, N = 185 and 358 | A cleared researcher was rated like a never-accused one (p = .54 and p = .51). Slightly higher, not significant | Recovery, no gain |
| Thorson (2016) | A corrected false claim restored belief but left the politician rated lower than the control | Lingering deficit |
| Arkes et al. (2007), Experiment 4 | The physician who used no aid was always rated highest, also against one who overruled the aid | Lingering deficit |
| Wintersieck (2017), N = 321 | Fact-checks confirming a candidate's accuracy raised evaluations. No doubt had been raised first | Gain, in a different situation |
| de Matos et al. (2007), meta-analysis | Good recovery after a real service failure raised satisfaction above the no-failure level, but not repurchase intention or image | Gain, weak |

No study from 2023 to 2026 was found in which observers rate a professional who correctly rebuts an AI flag. That is the gap this project fills.

**Corrections to version 1 and to our notes**

| Paper | We said | The source supports |
|---|---|---|
| Aamir et al. (2026) | Interviewed board directors | 13 Australian participants, directors and senior executives with a current or recent board-level role. It did not study trust or power between board and CISO, and it proposes no AI tool |
| Chen et al. (2026) | The doctor asked for the second opinion | The doctor offered to consult the AI. Effects held only when the doctor was seen as human-like |
| Strickland (1958) | Both subordinates produced the same output, and supervisors attributed the performance to the surveillance | Tasks were "similar". The attribution is a later author's explanation, not a reported finding |
| Reif et al. (2025) | A general penalty for AI users | The penalty shrank or vanished when the evaluator used AI often or the task suited AI. N = 4,439 is the number recruited |
| Prike et al. (2025) | Reduced favourability | Confirmed. One experiment, no trust scale, candidates never replied |

## 8. To verify before the final submission

| Claim in the abstract | Status | Where to check |
|---|---|---|
| Aamir et al.: the gap "limits their ability to interrogate risk reports" | Confirmed from the abstract. The full text could not be downloaded automatically | Open access: https://ro.ecu.edu.au/ecuworks2022-2026/8241 |
| Chen et al.: no reply by the doctor was tested | Not confirmed. Nothing in the abstract or press release describes a reply | Full text through the SFU library |
| Chen et al.: the effect "carried over to trust" | Direction not seen in the sources read | Full text through the SFU library |
| Strickland: monitored subordinate trusted less | Confirmed through a secondary source only | You have listened to the original, so check the wording against it |
| Prike et al.: candidate rated less competent and less moral after fact-checks | Confirmed by me in the accepted manuscript (competence d = 0.90 and 1.14, morality d = 1.11 and 1.31, each rated from 0 to 10). Trust was not measured | https://api.research-repository.uwa.edu.au/ws/files/463479508/Prike.2024.PoPs_pp.pdf |
| Reif et al.: AI users disclose less willingly | Read by the agent in the full text. The difference is small (4.91 against 5.25 on willingness to disclose to a manager) | https://www.pnas.org/doi/pdf/10.1073/pnas.2426766122 |
| Thorson: corrected politician still rated lower | Read by the agent in the full text, not checked by me | https://doi.org/10.1080/10584609.2015.1102187 |
| Greitemeyer and Sagioglou: cleared equals never accused | Confirmed by me in the full text | Open access at the DOI |
| Kelley: success against an obstacle is credited to the person (augmentation principle) | Reference confirmed in Crossref. The content is from my knowledge of the paper | https://doi.org/10.1037/h0034225 |

## 9. Optional citations for the power framing

Neither is in the abstract. Add one if the staff ask for a source on the power relation.

- See, K. E., Morrison, E. W., Rothman, N. B., & Soll, J. B. (2011). The detrimental effects of power on confidence, advice taking, and accuracy. *Organizational Behavior and Human Decision Processes*, *116*(2), 272–285. https://doi.org/10.1016/j.obhdp.2011.07.006. People with more power took less advice, explained by higher confidence. Supports "the director may not listen to the expert".
- Da Silva, J., & Jensen, R. B. (2022). "Cyber security is a dark art": The CISO as soothsayer. *Proceedings of the ACM on Human-Computer Interaction*, *6*(CSCW2), 1–31. https://doi.org/10.1145/3555090. Interviews with 15 CISOs and 6 senior leaders. The abstract describes the CISO's role as "precarious and, at the same time, superior". This is the closest source for the power relation that Aamir et al. do not cover.

## 10. Timeline

| Weeks | Work |
|---|---|
| 1–2 | Write the script, build the deck, draft the flag with a VLM, record the briefing and both answers, build the study page with logging |
| 3 | Pilot with four to five people, revise |
| 4–6 | Data collection |
| 7–8 | Analysis |
| 9–10 | Paper |

## 11. Short version, about 500 words

The brief asks for approximately 500 words. If the staff hold the draft to that, submit this text with the title, reference list and AI transparency statement from above. It makes the same claims as the long version.

**Research question.** Vision-language models (VLMs) let a listener point a camera at a slide and have it checked while the presenter is still talking. Boards are a likely setting: Aamir et al. (2026) found that gaps in directors' knowledge and confidence limit their ability to interrogate cyber-risk reports. But a checker lacks the company's context, so a flag can be right by the textbook and wrong for this company. We ask: when a board member, who holds the formal power, challenges a CISO, who holds the expertise, with a question prompted by a VLM checker, does trust in the CISO fall, and does the CISO's answer restore it or raise it above the level reached without a checker? Research on AI fact-checking measures belief in the checked claim. We measure trust in the checked person.

**Hypotheses.** Trust means the participant's rating of the CISO on the trust questionnaire. H1: right after the flag, before the CISO answers, participants with the checker rate their trust in the CISO lower than participants without it, although the briefing contains no error. After fact-checks, participants rated a political candidate as less competent and less moral (Prike et al., 2025). H2: after the answer, trust is higher when the CISO resolved the flag than when the CISO deflected. A researcher who was accused and then cleared was rated as credible as one never accused (Greitemeyer & Sagioglou, 2015). H3: after a resolved flag, trust is higher than when the same explanation is given without a checker. Earlier studies found recovery at best (Thorson, 2016), but there someone else did the clearing, and people give more credit for an outcome achieved against an obstacle (Kelley, 1973).

**Related work.** Strickland (1958) found that supervisors came to trust a more closely monitored subordinate less. In Prike et al. (2025), fact-checks made participants rate a political candidate as less competent and less moral, but the fact-checked claims were false and the candidate could not reply. Chen et al. (2026) is the closest design: an AI second opinion that disagreed with a simulated doctor made the doctor seem more uncertain and lazier to participants who saw the doctor as human-like. There the patient depends on the expert, and the doctor's reply was not tested. We reverse the power relation and let the expert answer. Reif et al. (2025) found that people who use AI at work are rated lazier and less competent and disclose it less willingly, so we log whether participants name the checker when asking.

**Method.** *Design.* Between-subjects, three conditions. Participants act as board members deciding on the CISO's budget request and watch one recorded, error-free briefing. (1) No checker: the chair asks a routine question and the CISO explains why a patch was deferred. (2) Flag, resolved: a simulated VLM checker flags the deferred patch and suggests a question, the participant asks it, and the CISO gives the same explanation. (3) Flag, deflected: the CISO deflects. The flag is drafted with a real VLM and then frozen, so we control its accuracy. *Measures.* Trust in the presenter (Mayer & Davis, 1999) before the flag, after it and after the answer; share of the budget approved; judgement of the flagged decision; trust in the checker (Jian et al., 2000); interaction logs. *Validity.* The briefing is constant, conditions 1 and 2 share the footage of the explanation, and conditions 2 and 3 differ only in the answer. The design is between-subjects because a participant who has seen a flag refuted reads later flags differently.

**Recruitment.** We will recruit 90 SFU graduate and upper-year undergraduate students from outside computer security, 30 per condition, through course mailing lists, department channels and word of mouth. The study runs online in about 15 minutes, after a pilot with four to five people.
