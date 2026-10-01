# Reviewer report: Aamir, Psaroulis, Grobler & Janicke (2026)

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. CHI '26. https://doi.org/10.1145/3772318.3791142

Source: `papers/aamir.md`, read in full (Sections 1 to 8, Appendix A.1 to A.3, footnotes, reference list). I ignored the ACM "AI Summary" block, which holds only a paywall notice and no summary text. Section numbers below are the paper's own. Tables are numbered as the paper's text refers to them (Table 2 = participants, Table 3 = board-pack items, Table 4 = conditions and signals, Table 5 = codebook).

---

## 1. The paper in one paragraph

This is a qualitative, semi-structured "problem-centred" interview study. It contains one study and no experiment (Section 3.1). The authors interviewed 13 people from 13 Australian organisations of different sectors and sizes. All held a board-level role at the time or had held one in the previous five years: 9 executive and 4 non-executive directors (Table 2, Section 3.2). Participants were recruited through professional networks between March and November 2023, and each interview ran 45 to 60 minutes online with two facilitators (Section 3.2). Nothing was manipulated and nothing was rated on a scale. The interview guide covered cyber literacy, strategy, board expertise and cadence with the CISO, incident governance, and resilience (Section 3.3, Appendix A.1). The analysis was reflexive thematic analysis with computer support (NLTK word splitting and LDA topic modelling). Two coders coded a calibration subset and agreed at about κ = 0.68, a figure the authors used only to start discussion (Section 3.4). The results are themes plus a few counts of how many directors mentioned something:
- Items that reach board packs: posture snapshots (10 of 13), incident trend lines (9), third-party exposure (8), compliance or exception logs (7) (Section 4.1).
- Laws named: Privacy Act (7), Notifiable Data Breaches scheme (4), GDPR (4), CPS 234 (3), SOCI Act (3), ISO 27001 and NIST CSF (3 each) (Section 4.1).

Main findings (Section 4.4):
- Directors find it hard to turn technical and compliance reports into decisions, and they distrust maturity scores.
- In directors' accounts, better decisions went with three things: a cyber-literate director, regular closed-door ("in-camera") sessions with the CISO, and independent assurance.
- Incident governance was the weakest link.

The authors turn this into a four-part "Board Cyber Governance Model" with a MUST/SHOULD/COULD checklist for directors (Section 5).

---

## 2. Verdict per hypothesis

Applies to all four hypotheses: the paper has no rating scale, no AI or automated checker, and no measure of how much directors trust the CISO. "Trust" appears only as stakeholders' trust in the organisation, for example P012 in Section 4.3 and "preserve trust with regulators, customers, and investors" in Section 5.1, Dimension 3. "AI" appears once, in a cited source about companies disclosing "cyber and AI oversight" (Section 2.1, Disclosure). Every verdict below therefore rests on what directors said about checking and challenging management, not on any measured change. None of the statements used below is counted; the paper says "several", "a few" or "many", or names example P-numbers.

### H1: trust drops when the checker's claim is shown

- **Verdict:** `Indirect support (by analogy)`. The support is weak.
- **Evidence:** Directors describe outside checks that contradict management as what lets them doubt management's account and push back:
  - Authors, Section 6.2 ("Challenge"): external assurance "gave directors confidence to challenge assumptions and calibrate risk appetite". They add: "where external challenges are routine, directors feel more confident in their judgments; where they are absent, decision-making remains anchored to internal narratives".
  - Authors, Section 4.1: "Where external assurance was routine, directors reported greater confidence in challenging management perspectives".
  - P008 (Section 6.2, repeated as the exemplar quote for the "Independent challenge" code in Appendix A.2, Table 5): "External reviews help us test whether management's view is too optimistic or whether we're actually prepared."
  - P013 (Section 4.2): "External assurance keeps us honest… it helps the board push back and ask: are we resilient, or are we just compliant?"
  - P002 (Table 4, Section 4.2): tightened vendor due-diligence requirements "after an external audit contradicted management's view".
  - P008 (Section 6.5): an external red-team style review "forced management to rethink its optimism". Within weeks the board approved more investment in monitoring and third-party due diligence.

  This is the behavioural core of H1: when an outside source contradicts the presenter, directors give less weight to the presenter's account. The paper does not show any drop in trust in the CISO as a person, and it has no numbers.
- **How their situation differs from ours:**
  - *Who checks whom.* Human experts (internal audit, external auditors, red teams, consultants) check management's reporting, which includes the CISO's. In our study a simulated AI checks the CISO's slides live.
  - *Who holds the power.* The board, which uses the check. This matches our study.
  - *Could the checked person reply?* The paper never describes management answering a finding. In the one detailed case (P008, Section 6.5), management gave in.
  - *Was the check correct?* In every case reported, the check was right and management had been too optimistic. The paper has no case of a wrong check. In our study every flag is wrong.
  - *Timing.* Checks are periodic: semi-annual or annual in Table 3, and "independent annual challenge" is a MUST in Section 5.2. They do not happen live during a briefing.
  - *Target of the check.* In the paper the check targets management's optimism about risk. Our flags target the correctness of the CISO's own work ("Architecture likely wrong").
  - *What was measured.* Directors' self-reported confidence to challenge, and the decisions they said followed. Trust was not measured.
  - *Counterpoint.* The paper treats challenge as normal good practice ("a culture of constructive challenge", Section 4.2). For experienced directors, questioning a CISO may not signal distrust at all.
- **Quote:** "External assurance, whether through internal audit, red-teaming, or third-party reviews, was seen as a critical counterweight to management narratives." (Section 6.2, "Challenge")

### H2: trust recovers or rises after a question gets a knowledgeable answer

- **Verdict:** `Silent`
- **Evidence:** The paper describes no check that turned out wrong, and no case of a CISO answering a challenge, so it says nothing about how directors then saw the CISO. The closest material is below. Both parts concern directors' confidence in their own decisions, not their trust in the presenter, and neither follows a flag.
  - *Directors want explanations tied to their own business.*
    - P009 (Section 4.1): "We can read the traffic lights, but we need to understand what they mean for the business next quarter."
    - P005 (Section 4.2): "Dashboards have to tell me what matters for this business, not just cyber jargon—keep it short, keep it in context."
    - Authors (Section 6.1): with colours alone, directors "struggled to know what was at stake"; with a short consequence story "they felt better equipped to make informed decisions".
  - *Directors value regular dialogue with the CISO.*
    - Authors (Section 4.2): in-camera sessions with the CISO were "enabling candid discussions of risks and clearer lines of accountability".
    - Authors (Section 6.2): these sessions were "where the real issues surfaced", and regular dialogue "allows directors to build confidence in their own questions".
    - P013 (Table 4): after in-camera CISO sessions, the board approved controls in the same meeting instead of deferring them.
- **How their situation differs from ours:** The paper has no flag-then-answer sequence and no measure before and after an answer. The outcome is directors' sense of being equipped to decide, not a trust rating. The explanation comes from routine reporting, not from a reply to a challenge.
- **Quote:** None supports H2. The closest is: "Traffic-light dashboards were remembered by almost every participant, but their value depended on the accompanying narrative: when presented with only colours, directors admitted they struggled to know what was at stake; when paired with a short consequence story, such as what "red" meant for service continuity in the next quarter, they felt better equipped to make informed decisions." (Section 6.1)

### H3: trust falls further after a denial or non-answer

- **Verdict:** `Silent`
- **Evidence:** No denial of a check and no non-answer is described. The closest material shows general scepticism toward reporting that is unexplained or too optimistic, not toward a presenter who fails to answer:
  - Authors (Section 6.2): "A few relied almost entirely on management updates, which left them exposed to optimistic reporting."
  - P006 (Sections 4.1 and 4.2): "Maturity scores can be gamed".
  - Authors (Section 6.1): directors were frustrated with "static artefacts that look polished but do not explain what they mean".
  - P003 (Table 4): "queried why several cyber risks remained unchanged for 18 months despite repeated incidents".
- **How their situation differs from ours:** There is no challenge-and-reply exchange and no measure of trust. The scepticism is aimed at artefacts such as dashboards and maturity scores, not at a person's answer.
- **Quote:** None supports H3. The closest is: "What stood out here was the frustration with static artefacts that look polished but do not explain what they mean for business continuity or risk appetite in the next quarter." (Section 6.1)
- **Reviewer note (about our design, not the paper):** As the brief describes it, the CISO always answers with the company's reason. Our design has no denial or non-answer condition, so H3 cannot be tested as it stands.

### H4: the trust change depends on the participant's action and the CISO's answers

- **Verdict:** `Silent`
- **Evidence:** The paper never connects directors' actions to trust. It does speak to two related points.
  - *What drives the action: literacy.*
    - Authors (Section 6.2, "Capability"): "Where this literacy was missing, boards often defaulted to accepting management narratives."
    - P013 (Section 6.2): directors "need to be literate enough to interrogate management without getting lost in jargon".
    - Authors (Section 4.2): boards with at least one member with prior technical or cyber exposure and "a culture of constructive challenge" "reported more probing questioning".
    - Authors (Section 6.4): boards relying on generic awareness sessions or briefings led by management "tend to struggle to interrogate metrics".
  - *Outcomes depend on both sides.*
    - Authors (Section 4.4): "information bottlenecks do not originate solely within the boardroom".
    - P002, P006, P008 and P010 (Section 4.4) said reporting from management or risk teams "lacked clarity or prioritisation".

  So the paper treats oversight quality as depending on both the director's questioning and management's reporting, which has the same two-sided shape as H4. But its outcome is decision quality as directors describe it, not trust.
- **How their situation differs from ours:** In the paper the "actions" are standing board habits (accepting management's account versus seeking challenge), not a single choice after a flag. The outcome is self-reported decision quality, not a trust rating. The participants are experienced directors.
- **Quote:** None supports H4 as worded. This sentence is useful for our knowledge moderator: "Where this literacy was missing, boards often defaulted to accepting management narratives." (Section 6.2, "Capability")

---

## 3. Other relevant content

### Who said what on challenge, management's account and independent checks

| Who | What | Where |
|---|---|---|
| P013 | "External assurance keeps us honest… it helps the board push back and ask: are we resilient, or are we just compliant?" | 4.2 |
| P008 | "External reviews help us test whether management's view is too optimistic or whether we're actually prepared." | 6.2; Table 5 (A.2) |
| P008 | Red-team style review "forced management to rethink its optimism"; the board approved investment within weeks | 6.5 |
| P002 | Tightened vendor due diligence after an external audit contradicted management's view | Table 4 |
| P013 | "…literate enough to interrogate management without getting lost in jargon." | 6.2 |
| P008 | "…we do need someone who can ask the second-order questions." | 4.2 |
| P006, P008 | Example questions: "What is the MTTR trend?" (P006); "Which vendor access routes failed last time?" (P008) | Table 4 |
| P007 | "Show me what changes if this KPI goes red next month. What decision do we need to make?" | 4.2 |
| P001 | "…you need to hear from others to really understand where the risks are." | 6.1 |
| P005 | Tracked simulation and training metrics "to test improvement claims" | Table 4 |
| P002, P005, P007 | Acknowledged gaps in cyber literacy, especially with technical or legal material | 5.1, Dimension 1 |
| P004, P011 | Targeted briefings and education improved their ability to contribute | 5.1, Dimension 1 |
| P002, P006, P008, P010 | Reporting from management or risk teams lacked clarity or prioritisation | 4.4 |
| (no P-number) | "boards often defaulted to accepting management narratives" | 6.2 |
| (no P-number) | "A few relied almost entirely on management updates, which left them exposed to optimistic reporting." | 6.2 |
| (no P-number) | Smaller or resource-constrained boards saw external checks as "costly or disruptive" and "preferred to rely on management's internal reporting unless regulators demanded more" | 6.2 |
| (no P-number) | An objective briefing from consultants "cut through internal optimism and provided language directors could use in challenge" (the authors attach citations [32, 36], mixing interview data with literature) | 6.2 |

### Design elements to borrow

- **Realistic slide content.** Table 3 lists what real board packs contain, with purpose, frequency and "decision it informs": posture snapshots as traffic lights or heatmaps, "crown jewels", third-party exposure, incident trends, compliance exceptions, external assurance, and regulatory change. Posture snapshots inform "Budget prioritisation; risk acceptance/deferral". This is a ready template for our five slides.
- **What directors want to see.**
  - Outcome measures such as mean time to detect, mean time to restore, and restore-to-SLA, rather than maturity scores (P006, Sections 4.1 and 4.2).
  - Plain language in the company's own context (P005, Section 4.2).
  - A one-page impact narrative with any red KPI: "What changes if the trend continues? What decision is requested?" (Section 5.2, SHOULD). This could shape our budget-request slide.
- **Wording for the neutral question.** P006 "What is the MTTR trend?", P008 "Which vendor access routes failed last time?" (Table 4) and P007 (Section 4.2) are real director questions. They are short, concrete and non-accusing, and fit our "ask the CISO to explain" option.
- **Our "challenge with the checker's finding" action has a real analogue.** Directors use outside briefings as "language directors could use in challenge" (Section 6.2).
- **Our flags echo a real director concern.** Directors already distrust maturity scores that "can be gamed" (P006). A flag such as "Numbers do not add up" on a maturity-type slide may seem plausible to participants. This is my inference, not the paper's.
- **Budget measure caution.** In the paper, an independent finding that there was a gap *increased* investment. P008 (Section 6.5): the board approved more spending within weeks. Section 6.2: external checks "gave the board defensible grounds for investment or escalation". Our flags question the CISO's competence rather than reveal a need, so their effect on budget may run the other way. Keep this difference in mind when interpreting budget share.

### Measures and exact wording

- No rating scales and no trust items.
- The interview prompts in Appendix A.1 that come closest to our topic:
  - "How does the board interact with the CISO or security lead (including in-camera)?" (A.1.4)
  - "How does independent challenge (internal audit, external consultants) influence board oversight?" (A.1.4)
  - "Where do you feel confident, and where do you see gaps?" (A.1.2)
- The paper treats "higher-quality decisions" as "timely, evidence-based, consequence-aware, and auditable" (Sections 1 and 6.2).

### Effect sizes for power calculation

None. The only number resembling a statistic is κ ≈ 0.68 between two coders on a calibration subset (Section 3.4). The rest are counts out of 13 (Section 4.1). Coding patterns "stabilised by interview 10" (Appendix A.3).

### Moderators

- **Cyber literacy.** This is the main individual difference in the paper: low literacy goes with accepting management's account, higher literacy with probing questions (Sections 4.2, 6.2 and 6.4). It supports keeping our self-rated cybersecurity knowledge measure. My inference: a student sample will mostly be low in cyber literacy, which by this paper's account could push choices toward "say nothing".
- **Organisation size and sector.** Regulated or large boards have routine CISO sessions and external assurance. Startup, education and small boards engage only intermittently and see external checks as costly (Sections 6.1, 6.2 and 6.5).
- **Attitude to external challenge.** Directors differ: some welcome it, others prefer management's reporting (Section 6.2).

### Stated limitations (Section 7)

- Small, non-random sample.
- Reliance on self-report, with social-desirability and recall bias.
- A single country.
- No direct performance outcomes.
- The sample may lean toward large, regulated, well-governed entities.
- "As an interview study, we did not observe meetings directly; future observational work could validate the routines described."

### Future work that our study would partly carry out (Section 8)

- The authors call for "Longitudinal board-meeting ethnography", which would include analysing "discourse (challenge, framing)".
- They also call for "Before/after studies on briefing redesign", run as field experiments.
- Our study records an actual choice (say nothing, ask, challenge) during a simulated briefing rather than a self-report. It can be framed as one step toward what Section 7 says is missing, observed behaviour instead of recalled practice, but it is a simulation with students, not observation of real boards.

### The board–CISO relationship as the paper describes it (for the framing check in Section 4)

- Directors have legal fiduciary duties (Sections 1 and 3.2).
- "Boards hold accountability for oversight but are not expected to perform technical analysis"; interpretation is a "shared responsibility" in which "boards set expectations for clarity, while subject-matter experts … operationalise interpretation" (Section 5.1).
- Boards "often depend on CISOs to translate technical risk" (Section 2.1, citing [2]).
- CISOs "lack decision-making power" (Section 1) and have "limited organisational authority" (Section 2.1). Both points come from cited work [34], not from these interviews.
- The relationship is also described as "risk-based partnership" (Section 2.1, citing [39, 48]) and as "regularity of dialogue" (Section 6.2).

### Leads in the reference list (not read; described here only as Aamir et al. describe them)

- [61] Opdenbusch, Hielscher & Sasse (2025, NDSS), "Where Are We On Cyber?". Aamir et al. say it documents "information asymmetries, limited technical confidence among directors, and a tendency to rely on high-level budget decisions rather than structured oversight", and difficulties "scrutinising management assurances" (Section 2.1). This may be closer to our budget scenario than Aamir et al.
- [34] Hielscher et al. (2023, USENIX Security): the source for CISOs lacking decision-making power.
- [24] Da Silva & Jensen (2022, CSCW), "The CISO as soothsayer".
- [44] Lowry, Vance & Vance, "Inexpert supervision: Field evidence on boards' oversight of cybersecurity".
- [45] Lowry, Sahin & Vance (2022, ICIS), "The quest for CISO legitimacy".

### Cautions on the text

- **Participant quotes are not stable across the paper.** The same P006 "Maturity scores can be gamed…" quote appears in four different wordings (Sections 4.1 and 4.2, Table 5, memo M-07 in A.3). P004's "More compliance does not equal less risk…" differs slightly between Sections 4.1 and 6.1. The quotes seem condensed or edited. If we quote a participant, cite one location and its exact wording.
- **Internal inconsistencies:**
  - The calibration subset is "two transcripts" in Section 3.4 but "three transcripts" in Appendix A.3.
  - Appendix A.1.6 mentions "RQ4", though there are only three research questions.
  - Section 7 calls the sample "directors and C-suite", while Section 3.2 says all participants held a board role.
- **Abstract versus body.** The abstract's "substantial gap in directors' knowledge and confidence" is backed in the body by "Several participants (e.g., P002, P005, P007)" and "Many reported…" (Section 5.1). No count is given.
- **Conversion errors, none of which changes a verdict:**
  - Tables are flattened, with column headers run together.
  - Figures appear only as their alt-text descriptions.
  - Footnote numbers are glued to words ("decision1", "(AICD)5").
  - "indicates that  90% of thematic patterns" (Section 3.2) has lost a symbol, probably "~".

---

## 4. Check of our current sentence

> "Boards are a likely setting: Aamir et al. (2026) found that gaps in directors' knowledge and confidence limit their ability to interrogate cyber-risk reports."

| Claim | Accurate? | Comment |
|---|---|---|
| "found that gaps in directors' knowledge and confidence limit their ability to interrogate … reports" | **Accurate as a paraphrase of the abstract.** | The abstract says: "Findings reveal a substantial gap in directors' knowledge and confidence, undermining effective oversight and informed decision-making. This deficit limits their ability to interrogate risk reports, challenge assumptions, and steer investment…". The body backs it only qualitatively: P002, P005 and P007 acknowledged literacy gaps (Section 5.1); "boards often defaulted to accepting management narratives" (Section 6.2). Three refinements are needed: (a) this is self-report from 13 Australian directors, so "report" or "describe" is safer than "found"; (b) the paper says the problem is not only in directors, because reporting from management "lacked clarity or prioritisation" (Section 4.4: P002, P006, P008, P010) and "The inconsistency is not due to lack of awareness" (Section 6); (c) in the body, "confidence" is mostly tied to independent assurance: directors felt more confident challenging management where external assurance was routine (Sections 4.1, 6.1, 6.2). |
| "cyber-risk reports" | Accurate. | The abstract says "risk reports". In context this means cyber-risk reporting (board packs, dashboards, metrics). |
| "Boards are a likely setting" (implied: for an AI content checker) | **Not from this paper.** | The paper never discusses AI or automated checking tools. Its only "AI" is a cited source on companies disclosing "cyber and AI oversight" (Section 2.1). The paper does show that boards already seek outside checks on management's account: "independent annual challenge" is a MUST (Section 5.2); an external observer is a COULD (Section 5.2); directors use consultants for "plain-speak" interpretation (Section 6.2). That gives a better-grounded reason why a checker fits this setting. |

**Suggested wording (short):**
"Boards are a plausible setting: in interviews with 13 Australian directors, Aamir et al. (2026) report that gaps in cyber literacy and confidence make it hard for directors to question cyber-risk reports, and that directors rely on independent reviews to push back against management's account."

**Suggested wording (longer, if space allows):**
"Boards are a plausible setting for such a checker. In interviews with 13 Australian directors, Aamir et al. (2026) report that gaps in cyber literacy and confidence, together with unclear reporting from management, make it hard for directors to question cyber-risk reports. Boards without this literacy often accept management's account, and directors describe independent reviews (audits, external reviews, red teams) as what lets them push back."

### Framing check: the board member as "the person holding the formal power"

**Consistent in substance, with four caveats.** The paper supports the split "board holds the authority, CISO holds the expertise":
- Directors carry legal fiduciary duties (Sections 1 and 3.2).
- Boards "hold accountability for oversight" and decide on budgets and risk acceptance (Section 5.1; Table 3; Figure 1 description).
- Boards "often depend on CISOs to translate technical risk" (Section 2.1).
- CISOs "lack decision-making power" (Section 1) and have "limited organisational authority" (Section 2.1).

The caveats:
1. **Wording.** The paper speaks of "accountability", "fiduciary duties" and "oversight", never of "formal power". The one power statement, that CISOs lack decision-making power, comes from cited work [34] (Hielscher et al., 2023), not from Aamir et al.'s interviews. For that claim, cite [34] directly.
2. **Authority lies with the board as a body, not with one director.** The paper warns that a single "cyber seat" "could distort accountability" (Section 6.4) and that "expertise should not be confined to one seat" (Section 6.2). Our participant plays one member. Say "a board member, who with the board decides on the budget", not "the person holding the formal power".
3. **The relationship is described as dependence and dialogue, not only as a hierarchy.** The paper uses "risk-based partnership" (Section 2.1), "regularity of dialogue" and in-camera sessions "where the real issues surfaced" (Section 6.2), and "a culture of constructive challenge" (Section 4.2). Challenging the CISO is presented as normal oversight, not as a sign of distrust.
4. **The CISO sits inside "management".** What directors challenge in the paper is "management narratives". Also, 9 of the 13 participants were executive directors, so they were themselves part of management (Table 2). The line between board and CISO is less sharp in the paper than in our scenario.

**Suggested wording:** "The board member is accountable for the decision and votes on the budget; the CISO holds the expertise the board depends on to understand cyber risk (Aamir et al., 2026). CISOs themselves report having little decision-making power (Hielscher et al., 2023, as cited in Aamir et al.)." Verify Hielscher et al. directly before citing it.

---

## 5. Bottom line

Cite this paper as **background for the setting**, not as evidence for any hypothesis. It shows that directors depend on the CISO's expertise, often accept management's account when they lack literacy, and use independent checks to push back. It gives only weak support to H1 (by analogy) and says nothing on H2, H3 or H4, because it measures no trust and involves no AI.

The fact that matters most for us: directors call independent checks "a critical counterweight to management narratives" (Section 6.2), but every check they describe was done by human experts (auditors, external reviewers, red teams), came periodically, and was right. The paper has nothing on a check that is wrong. That is exactly the gap our study, with a wrong checker flag followed by the CISO's answer, would fill.
