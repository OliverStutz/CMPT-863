# Reviewer report: Chen, Sun, Liao & Sundar (2026)

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies, 213*, 103824.

Source read: the full pdftotext extraction (pp. 1–17), including the appendix with the chatbot prompt. Page numbers below are the journal pages. Figures 1–6 appear only as captions in the text file, and Tables 1 and 2 are broken up by the extraction (see the end of Section 3).

---

## 1. The paper in one paragraph

A pre-registered online experiment with two groups, assigned between subjects (Section 3, p. 5). The authors recruited 260 US adults through CloudResearch Connect (MTurk panel), paid $2 for a 15-minute study. After dropping 35 who failed attention checks, 16 with no matching chat log, and 74 whose chats contained LLM errors, 135 remained: 68 in the agreement group and 67 in the disagreement group, aged 18–71 (mean 36.5) (Section 3.6, p. 6). Each participant chatted with "Dr. Alex", a therapist who was really a GPT-3.5 chatbot presented as a human doctor. The chat was a short cognitive behavioural therapy (CBT) session about the participant's own stress (Sections 3.1 and 3.3, p. 5; Appendix, p. 14). At the end the doctor recommended CBT and asked whether the patient wanted to hear the opinion of his AI assistant, CareBot. He then said only that CareBot agreed or disagreed with him, and went on to offer a full session. He gave no explanation (Appendix, p. 14). That agree/disagree statement was the only manipulation. After the chat, participants gave one round of ratings, all on 7-point scales (Section 3.7, pp. 6–8):
- trust in the doctor, as cognitive, affective and behavioural trust;
- four possible reasons for a change in trust: how credible the recommendation was, how lazy the doctor seemed, how uncertain they felt about their own medical situation, and how independent the doctor seemed;
- how human-like the doctor seemed;
- belief that machines are more objective than humans, taken before the chat.

Participants clearly noticed the manipulation. They rated the doctor–AI disagreement at 5.43 in the disagreement group and 1.68 in the agreement group (d = 3.33) (Section 4.1, p. 8). Main results, compared with agreement (Section 4.2, p. 9; Table 2, p. 12): disagreement made participants see the doctor as lazier, feel less sure about their medical situation, and find the recommendation less credible. Each of these went with lower trust in the doctor on all three trust measures. For cognitive trust, the part of the group difference that runs through each reason was:
- laziness: B = −.33, 90% CI [−.64, −.06], p = .041;
- uncertainty: B = −.44, 90% CI [−.75, −.18], p = .006;
- credibility: B = −.31, p = .023.

How independent the doctor seemed explained nothing (B = −.02, p = .535). The uncertainty route was stronger for people who believe machines are more reliable than humans (Section 4.3, p. 9). In exploratory analyses, the laziness, uncertainty and credibility routes were stronger for participants who found the doctor more human-like (Section 4.4, pp. 9–11). **The paper never reports mean trust per group or the overall difference in trust between the groups.** Only indirect effects (effects that run through one of the reasons) are reported. There is one study.

---

## 2. Verdict per hypothesis

### H1: trust in the CISO drops when a VLM claim is displayed

- **Verdict:** Indirect support (by analogy)
- **Evidence:** When the AI disagreed rather than agreed, participants saw the doctor as lazier and felt more uncertain, and both went with lower trust in the doctor. The indirect effects on cognitive / affective / behavioural trust were:
  - through laziness: −.33 / −.24 / −.28 (p = .041 / .034 / .040; Section 4.2.1, p. 9);
  - through uncertainty: −.44 / −.19 / −.37 (all p = .006; Section 4.2.3, p. 9);
  - through lower credibility of the recommendation: −.31 / −.12 / −.33 (all p = .023; Section 4.2.2, p. 9).

  All effects held after correction for multiple tests (Benjamini-Hochberg). The paper does not say whether B is in raw scale points. If it is, each route accounts for roughly a third of a point on the 7-point trust scale. There is no trust rating before the second opinion, no group without an AI, and no reported overall trust difference between the groups.
- **How their situation differs from ours:**
  - **Comparison.** The paper compares disagreement with agreement, not a flag with no flag. Agreement itself probably raised trust (the credibility route, Section 5.3, p. 11), so the gap mixes a gain in one group with a loss in the other. Our neutral group is a cleaner baseline.
  - **Who checks whom.** The AI is the expert's own tool. The doctor offers the second opinion and reports it himself. The patient only agrees to hear it. In our study the AI belongs to the person with power and checks the expert.
  - **Power.** The patient depends on the doctor. The paper frames trust as reliance "in vulnerable and uncertain situations" (Section 2.4, p. 3).
  - **Was the check right?** The disagreement is bare. The patient never hears what CareBot recommended or why, and nobody knows whether it was right. Our flags have specific content and are wrong.
  - **Could the expert reply?** He did not. The disagreement is the last thing said about the recommendation.
  - **What was measured.** Multi-item 7-point trust scales (cognitive, affective, behavioural), once, after the chat. We use a single 0–10 rating, repeated.
- **Quote:** "This finding suggests that patients take the AI assistant's second opinion seriously, so much so that when the assistant's opinion diverges from the doctor's initial recommendation, it raises uncertainty about their medical situation." (Section 5.1, p. 11)

### H2: trust recovers or rises after the CISO answers with an explanation

- **Verdict:** Silent
- **Evidence:** No group heard the doctor answer the AI. The prompt tells the chatbot only to "mention that your AI medical system CareBot [agrees/disagrees] with you" and then to offer a full session (Appendix, steps 7–8, p. 14). The authors *propose*, without testing it, that an explanation could reduce the loss of trust (Section 5.1, p. 11). They also give two example wordings a doctor could use (Section 5.5, p. 13).
- **How their situation differs from ours:** The same differences as under H1. In addition, H2 needs an expert who replies after the check, and the paper has none. Its disagreement is also never shown to be wrong, while in our study the CISO's answer shows the flag missed the company's reason.
- **Quote:** "Given that a second opinion from an AI assistant can shift the burden of proof onto doctors (Kempt & Nagel, 2022), the resulting reduction of trust due to increased medical uncertainty could potentially be mitigated by providing explanations for the assistant's disagreement." (Section 5.1, p. 11)

### H3: trust falls further after a denial or a non-answer

- **Verdict:** Silent
- **Evidence:** In effect, the disagreement group got *disagreement followed by no answer*. But there is no group that got disagreement with an answer, and no group rated trust before the doctor's (non-)reply. So the paper cannot show that a non-answer makes trust fall further than the disagreement alone. The authors guess that the missing explanation is why the doctor seemed lazy (Section 5.2, p. 11). That points in the direction of H3, but it is an interpretation, not a tested result. Their laziness items fit a non-answer well (Section 3.7.1, pp. 6–7):
  - "The doctor did not show the level of effort that corresponds to his profession"
  - "The doctor did not live up to the expected expertise"
  - "The doctor was lazy when providing the recommendation"
- **How their situation differs from ours:** The same differences as under H1. The doctor's silence was not a choice he made in reply to a question. It was built into the script, and the patient never asked anything.
- **Quote:** "This finding may be due to the doctor not providing any explanations after the AI's disagreement, which could be seen as a lack of effort." (Section 5.2, p. 11)

### H4: trust change depends on the participant's action and the CISO's answer

- **Verdict:** Silent
- **Evidence:**
  - **Participant's action.** The only choice participants had was whether to hear CareBot's opinion at all. Refusals are listed among the incomplete chat logs (Section 3.5, p. 6). The text does not say how many refused or whether any are in the final 135 (Figure 2 is missing from the extraction).
  - **Expert's answer.** The doctor's response never varied.
  - **What did change the effect.** Who the participant was:
    - Belief that machines are more reliable than humans strengthened the uncertainty route (moderated indirect effects: cognitive trust B = −.17, p = .045; affective −.05, p = .024; behavioural −.12, p = .042; Section 4.3, p. 9; Figure 3).
    - In exploratory analyses, how human-like the doctor seemed strengthened the laziness route (B = .35 / .30 / .31, all p = .001; Section 4.4.1, p. 9; Figure 4).
  - **An untested suggestion that maps onto our action choice.** The authors recommend that doctors not say openly that the AI disagreed, but raise it indirectly or with a short explanation (Section 5.5, p. 13). This mirrors our comparison of a neutral question (AI not named) with a challenge that names the checker.
- **How their situation differs from ours:** The same differences as under H1. There was only one path through the interaction, with no questions and no replies.
- **Quote:** "Informed by this finding, we recommend that doctors avoid explicitly disclosing an AI assistant's disagreeing opinion to their patients." (Section 5.5, p. 13)

---

## 3. Other relevant content

**Design lessons to borrow**
- **Script the messages instead of letting an LLM generate them.** 74 of 260 participants (28%) were dropped because of LLM errors (Section 3.6, p. 6). Only 72% of the complete chat logs followed the prompt without errors (Section 3.5, p. 6). The most common errors were:
  - robotic replies (18);
  - no recommendation given (16);
  - not asking whether the patient wanted CareBot's opinion (11);
  - two made-up stress sources.

  The authors conclude: "researchers should consider addressing the stochastic nature of LLM outputs by employing pre-scripted responses or more controlled prompt engineering" (Section 5.6, p. 13). This supports our fixed CISO answers and simulated checker.
- **Manipulation check to adapt** (Section 4.1, p. 8). Four items, 7-point, Cronbach's α = .92:
  - "CareBot did not agree with the doctor's decision"
  - "CareBot did not support the decision made by the doctor"
  - "CareBot expressed a second opinion differing from the doctor's decision"
  - "CareBot provided an alternative recommendation to the doctor's decision"

  A checker-group version could check that participants saw the flag as a challenge to the CISO.
- **How the expert was introduced.** "Dr. Alex" introduced himself by name only, with no profile. Only about half of the participants (n = 63) saw him as human-like (Section 5.7, p. 13). Our written CISO profile establishes expertise more firmly. A short realism check on the CISO would be cheap to add.
- **Deception and debrief.** Participants were told afterwards that the "doctor" was a chatbot and why (Section 3.3, p. 5). The authors add that participants "should not be misled about the true source of the communication" in healthcare (Section 5.7, p. 13). Our debrief should say that the checker was simulated and that its flags were wrong on purpose.

**Measures with exact wording (all 7-point; Section 3.7, pp. 6–8)**
- **Cognitive trust** (α = .94): four adjective pairs, undependable/dependable, dishonest/honest, insincere/sincere, untrustworthy/trustworthy.
- **Behavioural trust** (α = .97), for example "I plan to follow the doctor's recommendation to take care of my mental health." Our share-of-budget-approved measure plays the same role.
- **Affective trust** (α = .92), for example "The doctor displayed a warm and caring attitude towards me."
- **Laziness** (α = .86), items quoted under H3. Our flags ("Numbers do not add up") suggest sloppy work, so an effort or laziness item on the CISO could explain *why* trust drops. This is my suggestion, not the paper's.
- **Medical uncertainty** (α = .90), for example "I was not confident about the decision made by the doctor" and "The doctor said things to me that were confusing."
- **Belief that machines beat humans** (Yang & Sundar, 2024; 7 items, α = .95), for example "AIs are more reliable than humans" and "AIs are more dependable than humans". The sample averaged 3.04 of 7. It was measured *before* the interaction and moderated the effect (H4 of the paper).

  Suggestion, not the paper's: measure this before the slides in both groups. Our Jian et al. (2000) trust-in-checker score exists only in the checker group and is taken after the flags, so it cannot serve as a clean moderator.

**Effect sizes for our power calculation**
- The authors assumed a medium effect (f = .25) and needed 128 people (Section 3.4, p. 5). The effects they found were "smaller than what we had originally assumed". After the fact, their power was only .37 and .64 for the two human-likeness interactions (Section 5.7, p. 13).
- The paper gives **no effect size for the trust difference between groups**, so it cannot feed our power calculation directly.
- The large d = 3.33 is for *noticing* the disagreement, not for trust.
- Background, my calculation rather than the paper's: with 30 per group, a two-group comparison has 80% power only for differences of about d = 0.72. The before-flag and after-flag ratings within each person are where our study has power.

**Timing.** The whole study took 15 minutes and the chat 10–15 minutes. Trust was measured once, right after (Section 5.7, p. 14).

**Stated limitations** (Section 5.7, pp. 13–14)
- MTurk sample.
- Mental-health consultation only.
- An LLM stood in for a human doctor.
- Only half the participants saw the doctor as human-like.
- Human-likeness was measured, not manipulated, so its effects are "conditional associations".
- The interaction was brief, so trust "may be unwarranted and temporary".
- Role-play with low stakes.
- Familiarity with therapy was not measured.

**Future work the authors call for that our study carries out**
1. The expert explains the AI's disagreement (Sections 5.1 and 5.5).
2. Trust tracked over time, because "trust formation is a dynamic and fluid process" (Section 5.7, p. 14).
3. Scripted messages instead of free LLM output (Section 5.6).
4. Contexts other than mental health (Section 5.7).
5. Measuring how people see the AI itself: "Future research should directly measure patients' perceptions of the AI's recommendation, such as its usefulness and credibility" (Section 5.3, p. 12). Our Jian et al. checker-trust scale does this.

**Theory the paper uses that we could cite**
- Expectancy violations (Burgoon, 1993).
- The effort heuristic (Kruger et al., 2004).
- The machine heuristic (Sundar & Kim, 2019).
- Warranting theory (Walther & Park, 2002).

**Reporting problems and extraction issues**
- **Table 2 labels are scrambled.** In the extraction, the "Supported / Not supported" labels of Table 2 (p. 12) are out of order. From the text: RQ1, H1, H2 and H4 were supported; H3 and H5 were not.
- **The paper contradicts itself on the sign of the uncertainty route.** One row of Table 2 says uncertainty "was associated with higher patient trust", and Section 4.4.3 (pp. 9–10) gives positive B values for uncertainty → trust. Section 4.2.3 and the Discussion say uncertainty lowered trust. This inconsistency is in the paper, not caused by the extraction.
- **Uncertainty SEs do not match their CIs.** In Section 4.2.3 (p. 9) the standard errors are .05 / .04 / .08, but the 90% CIs imply about .17 / .10 / .16 (my calculation). The laziness SEs do match their CIs.
- **Manipulation check does not fit N = 135.** It is reported as t(98) = 16.61 (p. 8). With 67 and 68 per group and d = 3.33, t would be about 19. The reported t fits about 100 people (my calculation).
- **"Human-like" overlaps with trust.** How human-like the doctor seemed was rated after the chat and correlates .77 with cognitive trust and −.63 with laziness (Table 1, p. 8). "Saw the doctor as human-like" therefore overlaps heavily with "trusted the doctor". This is my reading of Table 1.
- **Two statements about the uncertainty route.** The Discussion says the uncertainty effect "only occurred when patients perceived the LLM-simulated doctor as more humanlike" (Section 5.1, p. 11). The pre-registered test, however, reports the effect across the whole sample (Section 4.2.3).

---

## 4. Check of our current sentence

> "Chen et al. (2026) is the closest design (N = 135). When an AI second opinion disagreed with a simulated doctor, participants who saw the doctor as human-like perceived more uncertainty and laziness, which carried over to trust. There the patient depends on the expert, and the doctor's reply was not tested. We reverse the power relation and let the expert answer, which is what H2 tests."

| Claim | Accurate? | Comment |
|---|---|---|
| "the closest design" | A judgement, defensible | What is close: an AI second opinion contradicts a human expert in front of a lay person, and the outcome is trust in the *expert*. What differs: disagreement vs agreement with no no-AI group; the expert brings in his own AI; the disagreement has no content; trust is measured once, afterwards. |
| "N = 135" | Yes | This is the analysed sample: 67 disagreement, 68 agreement, from 260 recruited (Section 3.6, p. 6). |
| "an AI second opinion disagreed with a simulated doctor" | Yes, but incomplete | The comparison is with the AI *agreeing*, not with no AI. The "simulated doctor" was an LLM chatbot presented to participants as a human therapist. He consulted his own AI assistant and reported its verdict himself. |
| "participants who saw the doctor as human-like perceived more uncertainty and laziness" | Partly | The pre-registered results hold across all participants: disagreement raised uncertainty and laziness (Sections 4.2.1 and 4.2.3). The human-like part is exploratory, rests on a rating taken after the chat, is underpowered (post-hoc power .37 and .64), and the authors call it a "conditional association" (Section 5.7). The paper's abstract does lead with the human-like version, which is probably where the wording came from. "Uncertainty" means the participants' uncertainty about *their own medical situation*, not about the doctor. |
| "which carried over to trust" | Direction yes, wording too causal | Both were linked to lower cognitive, affective and behavioural trust. The reasons and trust were rated at the same time after the chat, so this is an association. The paper reports no overall trust difference between groups. |
| "There the patient depends on the expert" | Yes | This matches the paper's framing of patient trust "in vulnerable and uncertain situations" (Section 2.4, p. 3). The paper does not discuss power as such. |
| "the doctor's reply was not tested" | Yes, and understated | The doctor never replied to the disagreement at all. The authors suspect this silence made him look lazy (Section 5.2) and suggest an explanation could limit the trust loss (Section 5.1). |
| "We reverse the power relation and let the expert answer, which is what H2 tests" | Our claim; consistent with the paper | Worth adding that we also move the AI from the expert's hands to the power-holder's. Chen's "disagreement then silence" is the nearest analogue to our H3 non-answer case. |

**Suggested wording (plain):**

> Chen et al. (2026) is the closest design (N = 135). At the end of a short therapy chat, a doctor (a chatbot presented as a human) consulted his own AI assistant and told the patient that it agreed or disagreed with him. Compared with agreement, disagreement made patients less sure about their situation and made the doctor seem lazier, and both went with lower trust in the doctor, rated once after the chat; these effects were stronger among patients who found the doctor human-like. There the patient depends on the expert, the expert brings in the AI himself, and he never answered the disagreement; the authors suggest that an explanation could limit the loss of trust. We reverse the power relation, give the checker to the board member, and let the expert answer, which is what H2 tests.

---

## 5. Bottom line

Cite this paper for **H1 as the closest analogue** (indirect support): an AI contradicting an expert went with lower trust in the expert. Also cite it as **motivation for H2 and H3**, because the authors propose, without testing, that the expert's explanation could repair trust. The fact that matters most for us: **the doctor never replied to the AI's disagreement, and the authors think this silence is why he looked lazy.** The expert's reply is exactly the open question our study tests. Do not use the paper for the power calculation. It has no neutral baseline and reports no overall trust difference between its groups.
