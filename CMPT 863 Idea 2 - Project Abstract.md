# Does an AI content checker in the Audience's Hands Cost the Presenter Trust They Did Not Lose?

## Research question

Vision-capable LLMs make it easy for anyone in a meeting to point a phone or laptop camera at a projected slide and have the content checked while the presenter is still talking. Boards are an obvious place for this. Aamir et al. (2026) interviewed board directors about cybersecurity governance and found that many lack the literacy and the confidence to question the cyber-risk reports presented to them. An automated checker could close that gap. It could also change how the directors see the person presenting(human to human trust relationship).

Research on AI fact-checking mostly asks whether people believe the checked *claim*. We ask what happens to trust in the *person*:

> **RQ.** When an audience member uses an automated slide checker, does a flag lower their trust in the presenter beyond what the presenter's actual errors justify?

The contribution is a test of whether a verification tool acts as a social signal about a person and not only as information about a number. If it does, the accuracy of a checker is not enough to evaluate it, because its false positives are paid for by the presenter.

## Hypothesis

We compare three conditions: no checker, a checker whose flags are all correct, and a checker that raises the same correct flags plus one false-positive flag on a correct figure.

- **H1 (flag penalty).** Trust in the presenter is lower with correct flags than with no checker. People who are fact-checked lose standing when the audience has no prior bond with them (Prike et al., 2025), and an AI that disagrees with a professional lowers trust in that professional (Chen et al., 2026).
- **H2 (unjustified penalty).** Trust in the presenter is lower with the additional false-positive flag than with correct flags only, although the presenter made exactly the same errors in both conditions. People who consult a fallible algorithm act on its verdict about another person: with a lie detector that was right about two times in three, accusations rose from 19% to about 85% of flagged cases (von Schenk et al., 2024). We expect most participants to treat the false flag as a third error.

H2 is the core test. Any drop between these two conditions cannot be attributed to the presenter's errors, because those are held constant.

We also ask, without a directional prediction, whether the penalty falls on perceived ability or on perceived integrity.

## Related work

**Checking itself lowers trust.** Strickland (1958) had supervisors monitor one of two subordinates more often although both produced the same output. Supervisors ended up trusting the less-monitored subordinate more. This is the pre-AI baseline for our claim that a check changes how a person is judged independently of what the check finds. Strickland's monitor was a supervisor with authority. In our setting the audience holds the tool and the verdict comes from a machine.

**A tool attached to a professional costs them standing.** In four experiments, Arkes et al. (2007) found that a physician who used a computerised diagnostic aid was rated lower in diagnostic ability than one who diagnosed unaided, whether the physician followed the aid or overruled it. There the professional chose to use the tool. We move the tool to the audience, so the presenter has no say in being checked.

**An AI that contradicts the professional.** Chen et al. (2026) is the closest existing design. An AI second opinion that disagreed with a doctor raised perceived uncertainty and perceived laziness, and through them lowered trust in the doctor. The doctor in that study was simulated, asked for the second opinion, and spoke to one patient. We extend this to a presenter in front of an audience, to a persistent visual artifact, and to a checker that is sometimes wrong.

**Fallible checkers change how people treat each other.** Von Schenk et al. (2024) show that an imperfect lie-detection algorithm makes people far more willing to accuse another person of lying. The study measured accusations, not trust, in a low-stakes game. We measure trust in a specific person after a specific false flag.

**Corrections transfer to the person, but not always.** Prike et al. (2025) found that fact-checks sharply reduced favourability and vote share for unfamiliar fictional candidates. Chae and Tewksbury (2026) found that neither AI-labelled nor human fact-checks changed feelings toward well-known politicians. Familiarity is the likely difference. Our presenters are strangers to the participants, as a new CISO or an external consultant is to a board, so we expect the penalty to appear.

None of these studies puts the checker in the audience's hands, lets it contradict a presenter, includes false positives, and measures trust in the presenter. That combination is our contribution.

## Method

### Study design

A within-subjects experiment with one factor, checker condition, at three levels:

1. **No checker.** The participant watches the talk.
2. **Correct flags.** An overlay flags the two planted errors.
3. **Correct flags plus one false positive.** The same overlay flags the two planted errors and one figure that is correct.

Stimuli are three recorded briefings of about four minutes each, in which a presenter reports a cybersecurity risk topic to a board (for example ransomware readiness, phishing exposure, vendor risk) and asks for a budget. Each talk has a different presenter, so trust does not carry over between talks, and two planted numeric errors of the same types (percentages that do not sum, a figure that contradicts an earlier slide).

Flags are scripted and time-synced to the video. Participants are told that an AI slide checker reads the slides through the camera. Scripting the flags lets us control the checker's accuracy exactly, which a live OCR and LLM pipeline would not. Each flag carries a short reason, for example "contradicts slide 3".

Participants are given a scenario in writing: they sit on the board and must decide on the funding request after each briefing. The pairing of talk and condition and the order of conditions are counterbalanced with a Latin square. A session takes about 35 minutes. We will pilot with three to four people before collecting data.

### Measures

Collected after each talk:

- **Trust in the presenter** (primary). Ability, benevolence and integrity items adapted from Mayer and Davis (1999), 7-point scale.
- **Behavioural trust.** The share of the requested budget the participant approves, 0 to 100%.
- **Error identification.** The participant marks which figures in the talk were wrong. This gives hits and false alarms, and shows whether they accepted the false flag.
- **Trust in the tool.** Short form of the Jian et al. (2000) trust in automation scale, in the two checker conditions.
- **One open question** on what drove their rating of the presenter.

The system logs when participants open or dismiss a flag.

### Validity

The H2 contrast isolates the unjustified part of the penalty by design: the presenter's errors are identical in conditions 2 and 3, and only the checker's output differs. The H1 contrast mixes two things, the checker and the participant's awareness of the errors. The error identification task lets us separate them by comparing participants who noticed the errors unaided with those who did not. Using a rating scale and a budget decision guards against the known gap between stated trust and behaviour. We analyse ratings with a linear mixed model, condition as fixed effect, participant and talk as random intercepts, and two planned contrasts (2 vs 1, 3 vs 2).

The main limits are that recorded talks remove the presenter's chance to respond to a flag, and that students stand in for directors.

## Recruitment

We will recruit 30 participants, graduate and upper-year undergraduate students at SFU from outside computing security, through course mailing lists, department channels and word of mouth. Like the directors in Aamir et al. (2026), they must judge a technical report without being experts in it. Participation is voluntary. Sessions run in person or over video call with screen sharing.

## Feasibility

| Weeks | Work |
|---|---|
| 1–2 | Write the three scripts, build the decks, record the talks, build the overlay |
| 3 | Pilot, revise materials |
| 4–6 | Data collection |
| 7–8 | Analysis |
| 9–10 | Paper |

No live AI pipeline is needed for the study.

## References

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems*, 1–22. https://doi.org/10.1145/3772318.3791142

Arkes, H. R., Shaffer, V. A., & Medow, M. A. (2007). Patients derogate physicians who use a computer-assisted diagnostic aid. *Medical Decision Making*, *27*(2), 189–202. https://doi.org/10.1177/0272989X06297391

Chae, J. H., & Tewksbury, D. (2026). Perceiving AI intervention does not compromise the persuasive effect of fact-checking. *New Media & Society*. https://doi.org/10.1177/14614448241286881

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, *213*, 103824. https://doi.org/10.1016/j.ijhcs.2026.103824

Jian, J.-Y., Bisantz, A. M., & Drury, C. G. (2000). Foundations for an empirically determined scale of trust in automated systems. *International Journal of Cognitive Ergonomics*, *4*(1), 53–71. https://doi.org/10.1207/S15327566IJCE0401_04

Mayer, R. C., & Davis, J. H. (1999). The effect of the performance appraisal system on trust for management: A field quasi-experiment. *Journal of Applied Psychology*, *84*(1), 123–136. https://doi.org/10.1037/0021-9010.84.1.123

Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, *46*(5), 1081–1100. https://doi.org/10.1111/pops.13059

Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, *26*(2), 200–215. https://doi.org/10.1111/j.1467-6494.1958.tb01580.x

von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, *27*(7), 110201. https://doi.org/10.1016/j.isci.2024.110201

## AI transparency

We used AI to search for and screen related papers and to draft this text. All cited papers are to be read by the team before submission.

---

## Draft notes (delete before submission)

**Decisions I made that are yours to confirm**

1. **The presenter-side study is left out.** The ideas file had a second study where participants give a talk while believing they are being checked. I scoped the abstract to the audience study only, because the rubric penalises designs that are "overly complicated" and awards feasibility for fitting in ten weeks. If you want it back, Nyhan and Reifler (2015) becomes a core paper.
2. **Paper set.** Core: Strickland, Arkes, Chen, von Schenk, Prike. Counter-evidence: Chae and Tewksbury. This is one paper more than the brief asks for. Drop Arkes or Prike to get to five.
3. **N = 30, SFU students, 35-minute sessions.** Placeholders. Adjust to what the team can run.
4. **Three different presenters.** This needs three people willing to be recorded. The alternative is one presenter for all talks, which is easier but lets trust carry over between talks.

**Claims to verify while reading**

- Strickland (1958): the finding is taken from secondary sources. Read the original before submitting.
- Chen et al. (2026): the simulated doctor and the indirect effect are from the abstract and press release, not the full text.
- Arkes et al. (2007): "whether the physician followed the aid or overruled it" rests on Experiment 4, which used medical students.
- Aamir et al. (2026): the description follows our ideas file. Check it against the paper.
- Mayer and Davis (1999) and Jian et al. (2000): the citations are confirmed in Crossref. Neither scale was part of the literature search, so check the items fit a presenter and a slide checker.
