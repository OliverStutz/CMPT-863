# Reviewer report: Reif, Larrick & Soll (2025), "Evidence of a social evaluation penalty for using AI", PNAS 122(19), e2426766122

Source: the full Markdown text (`papers/reif.md`, converted from Europe PMC XML, PMC12088386), read in full. Figures 1–4, the SI Appendix (Tables S2, S5, Figs. S1, S2) and the footnote text are not in the file. "md line" means the line number in `reif.md`.

### 1. The paper in one paragraph

Four preregistered online experiments test whether people who get help from AI at work are judged worse than people who get the same help from another source. The total is N = 4,439 recruited, not counting 801 real "candidates" in Study 3. Participants were adults on the US online panels Prolific and CloudResearch Connect, mean age 37.7 to 43.2, most with a bachelor's degree or higher, not students. In every study the person being judged is the one who **uses** the AI. Nobody is checked or criticised by an AI, the AI's output is never shown, and trust is never measured.

- **Study 1** (N = 497; Results md 37–45, Methods md 119–123). Participants imagined they had started using a "generative AI tool" or a "dashboard creation tool" to write customer reports. They then rated how others would see them and how likely they were to tell others about the tool. Compared with the dashboard group, the AI group expected to be seen as:
  - lazier (3.25 vs 2.72)
  - more replaceable (3.39 vs 2.83)
  - less competent (4.72 vs 5.45)
  - less diligent (4.66 vs 5.25)

  All four differences had p < 0.001. The AI group also said they were less likely to tell their manager (4.91 vs 5.25, p = 0.016) or colleagues (4.85 vs 5.17, p = 0.026). The text does not state the scale range; the values fit a 1–7 scale.
- **Study 2** (N = 1,203; md 49–59, 127–131). Observers rated one of 384 short employee descriptions. The descriptions varied gender, 8 occupations, age and the source of help: AI, the same help from a non-AI source (e.g. a paralegal instead of generative AI), or no help mentioned. Employees helped by AI were rated lazier (2.50) than employees helped by a non-AI source (2.16, p < 0.001) or with no help mentioned (2.02, p < 0.001). They were also rated less competent, less diligent, less independent and less self-assured, but not less ambitious or dominant. The text gives no means for these other traits; significance is shown only in Fig. 2, which is not in the file. The target's gender, age and occupation did not change the effect.
- **Study 3** (801 real candidates, 1,668 "managers"; md 63–71, 135–145). This was a hiring task with real money at stake. Managers saw one real candidate whose profile said they use generative AI daily or never. They rated the candidate's task fit (three items, 1–7). They then chose to hire that candidate or a random one, and were paid according to how well the hired candidate had actually done the task. On average, AI users and non-users got the same task-fit rating (p = 0.516). The result depended on the manager's own AI use (1–10 scale). Managers who rarely use AI favoured the non-user, and managers who often use AI favoured the daily user. The size of this dependence was b = 0.06 per scale step for task fit (p < 0.001) and b = 0.15 per step for hiring (p < 0.001).
- **Study 4** (N = 1,002; md 75–87, 149–153). Participants imagined hiring a gig worker, "Mike".
  - Mike was described as using either AI tools or Microsoft Word and PowerPoint.
  - Participants rated his laziness (four items) before they learned the task.
  - The task was either writing 750 handwritten thank-you notes ("manual") or 750 personal emails ("digital"). The digital version added: "Generative AI tools might be useful for this task, and the worker is welcome to use them."
  - They then rated Mike's task fit (three items) and recommended hiring him or not.

  Results:
  - The AI user was rated lazier, by 0.47 points (p < 0.001).
  - For the manual task, the AI user got lower task fit (4.18 vs 4.75, p < 0.001).
  - For the digital task, the difference was not significant (5.75 vs 5.58, p = 0.112).
  - Raters who themselves use AI weekly or daily did not rate the AI user as significantly lazier.

### 2. Verdict per hypothesis

**H1 (trust in the CISO drops when a flag is shown)**
- **Verdict:** `Silent`
- **Evidence:** No study shows a person being flagged, checked or criticised by an AI, and no study measures trust. The only related content is theory the paper cites but does not test: observers tend to explain behaviour by the person's traits rather than by the situation (Introduction, para 1, md 19; "Attribution Theory and Emerging Technologies", para 2, md 29). One could extend this to our case: a participant might read a flag as the CISO's fault rather than as a company-specific reason. That extension is the reviewer's own. The paper applies the theory only to people who use AI.
- **How their situation differs from ours:**
  - In the paper, the AI helps its user do the user's own work, and the person judged is that user. In ours, the AI checks another person's work, and the person judged is the one being checked.
  - The judges in Studies 3–4 hold hiring power over the judged person, much as our board member holds power over the CISO. But the judged person never replies.
  - The AI's output is never shown, so whether it is correct never comes up.
  - Measured: laziness, competence, diligence, independence, self-assurance, ambition, dominance, replaceability, task fit and hiring. Trust is not measured.
- **Quote:** "Attribution theory posits that observers tend to favor dispositional explanations relative to situational ones (5, 6), suggesting they may attribute the use of assistance to personal deficits rather than circumstantial causes." (Introduction, para 1, md 19). Cite it only as background theory, and note that the paper applies it to people who use AI.

**H2 (trust recovers, and can end higher, after the CISO's explanation)**
- **Verdict:** `Silent`
- **Evidence:** Nobody being judged ever answers anything. The only possible analogy is Study 4 (Results, Study 4, paras 2–3, md 77–79). There, information about the situation (AI is useful and allowed for this task) closed the task-fit gap between the AI user and the non-user. In the digital task the gap was 5.75 vs 5.58 (p = 0.112); in the manual task it stayed (4.18 vs 4.75, p < 0.001). The analogy is weak for three reasons:
  1. The situational information did not undo the negative judgment of the person. Laziness was rated before the task was described. The AI user was rated lazier (by 0.47, p < 0.001) in both task groups, and in the digital task that laziness pulled task fit down even more strongly. The gap closed only because a separate positive effect of AI use on task fit (+0.39) cancelled it.
  2. The AI user did not end significantly above the non-user (p = 0.112). So there is no counterpart to H2's "increase above the never-flagged level".
  3. The information came from the experimenter to a separate group, not from the judged person answering a doubt. There is no before/after rating of the same person.
- **How their situation differs from ours:** The judged person is the AI user, not someone the AI checked. The judged person gives no reply. Each participant gives one rating of the target, not repeated ratings. The outcome measured is task fit for a gig job, not trust.
- **Quote:** "However, this penalty was entirely offset for the Digital task for which AI was described as useful." (Results, Study 4, para 3, md 79). If cited at all, cite it for task fit only, not for laziness, and not as recovery after a challenge.

**H3 (trust falls further after a denial or non-answer)**
- **Verdict:** `Silent`
- **Evidence:** There is no reply of any kind from the judged person: no denial, no non-answer, no explanation.
- **How their situation differs from ours:** The judged person is a written description or profile and never speaks. Nobody is flagged.
- **Quote:** none.

**H4 (the trust change depends on the participant's action and the answers heard)**
- **Verdict:** `Silent`
- **Evidence:** The paper does not test whether a judge's own actions change how the judge later sees another person. What it offers bears on **which action participants choose**, not on the trust change. In Study 1, people who imagined using an AI tool at work expected to be seen as lazier (3.25 vs 2.72) and less competent (4.72 vs 5.45). They also said they were less likely to tell their manager (4.91 vs 5.25, p = 0.016) and colleagues (4.85 vs 5.17, p = 0.026) (Results, Study 1, paras 2–4, md 39–45). In our design, "challenge" names the checker and shows that the participant relies on AI, while "ask to explain" does not. Participants who pick "ask" may therefore do so partly to avoid looking reliant on AI, not only because of how much they trust the CISO. This inference is the reviewer's own; the paper does not test it. It means the action is partly a self-presentation choice, and the H4 analysis should allow for that.
- **How their situation differs from ours:**
  - In Study 1 the AI does the participant's own job (writing reports), which casts doubt on the participant's own effort and ability. In ours, the board member uses a checker supplied by the setup to question someone else's work, as part of an oversight role.
  - In Study 1 the audience is the manager (above) and colleagues (same level). In ours it is the CISO, who has less formal power than the board member, plus the rest of the board by implication.
  - Study 1 measured a stated likelihood of telling others in an imagined scenario, not an observed choice. The differences are small, about 0.3 points.
- **Quote:** "This is an important oversight because people care about how their actions will be perceived by others (7): people may choose not to use AI—or not to disclose their use of AI—if they expect to incur a social penalty." (Discussion, Implications, para 2, md 97)

### 3. Other relevant content

**Direction (a): how the participant expects to be seen when naming the checker**
- **Study 1 is the relevant finding.** People expected to be judged lazier, more replaceable, less competent and less diligent for using a generative AI tool, compared with a non-AI tool. They also said they would less readily tell their manager and colleagues (md 39–45).
- **The anticipated penalty appeared even under favourable conditions.** Participants in both groups had read that "there were no rules prohibiting the use of this tool" and that they "still ensure that the reports are accurate and meet quality standards" (Methods, Study 1, md 121).
- **Saying that others use AI too did not help.** A supplementary study found no difference in expected or actual penalty whether AI use was described as rare (only one person uses it) or common (many colleagues also use it) (Limitations, md 107; SI Fig. S1, not in the file). Implication for us: telling participants that the board routinely uses the checker may not remove any hesitation to name it.
- **The judge's own AI use matters.** Two findings support this:
  - Study 3 (md 65–67): only managers who rarely use AI favoured the non-user.
  - Study 4 (md 81): the more the rater used AI, the smaller the laziness gap (b = −0.23 per step on a 1–5 scale, p < 0.001). The AI user was rated significantly lazier by raters at levels 1–3 (never, once or twice, occasional), but not by raters at levels 4–5 (weekly, daily).

  Background, not from the paper: SFU students are likely to use AI often, so any such penalty may be small in our sample.
- **The paper's theory predicts less penalty when AI use is not a free choice.** The paper frames the penalty as observers reading the *choice* to use help as a trait rather than a response to the situation (md 19). In our study the checker is provided by the setup, which by that theory should invite fewer trait-based inferences. The paper does not test imposed versus chosen use, and the rare-versus-common result above points the other way.
- **The paper does not test the opposite possibility.** Naming the AI ("the checker says...") might let a participant shift a challenge away from themselves and onto the tool. This is the reviewer's own speculation.
- **Design caution for the action log:** "ask to explain" and "challenge with the checker's finding" differ in two ways at once: whether the AI is named, and neutral versus confronting tone. A lower share of "challenge" choices therefore cannot be pinned on hiding AI use. To separate the two, either add an option such as "ask to explain, mention the checker", or ask after the task why the participant chose each action. A further option: add a short post-task item adapted from Study 1, "How do you think the CISO / the board saw you when you used the checker?", rated on competent and lazy. Study 1's exact item wording is not in the file; the authors say it is on OSF (md 115).

**Direction (b): how the checked person is judged**
- **Nothing in the paper.** No study shows a person whose work was checked, flagged or corrected by an AI, and no study measures how such a person is judged. The introduction mentions only one other case, patients who assume that physicians who use diagnostic aids are less capable (md 27, refs 15–16). That is again about the person who uses the AI.
- **Design caution that follows from the paper:** the CISO must not appear to use AI. By Study 2, the CISO would draw a laziness and competence penalty from observers if the profile, the slides or the scripted answers suggested the CISO used AI to prepare the briefing (for example, "our AI tool estimated..."). That would be a second, unwanted source of lower trust.

**Study 4 task-fit result, stated precisely** (Results, Study 4, paras 2–3, md 77–79; Methods md 151–153)
- **Design:** 2 × 2 between groups. The candidate used AI tools or used Word and PowerPoint; the task was manual (handwritten notes) or digital (emails, with AI said to be useful and allowed). Laziness was rated before the task was revealed; task fit (three items, α = 0.92) was rated after. The text gives no scale range for Study 4's task fit; Study 3's task fit was 1–7.
- **Analysis of variance:** there were three significant effects:
  - Tool: F(1, 998) = 7.43, p = 0.007
  - Task: F(1, 998) = 256.28, p < 0.001
  - Tool × Task interaction: F(1, 998) = 24.68, p < 0.001
- **Manual task:** AI user 4.18 (SD 1.50) vs Word/PowerPoint user 4.75 (SD 1.14), t(998) = 5.44, p < 0.001. The AI user was rated a poorer fit.
- **Digital task:** AI user 5.75 (SD 1.05) vs Word/PowerPoint user 5.58 (SD 0.98), t(998) = 1.59, p = 0.112. The difference is not significant; it points in favour of the AI user.
- **Path through laziness:**
  - The AI user was rated lazier: a = 0.47, SE 0.07, p < 0.001. This is the same in both task groups, because laziness was rated before the task was known.
  - Higher laziness lowered task fit more in the digital task (b = −0.43, SE 0.05) than in the manual task (b = −0.21, SE 0.05), both p < 0.001.
  - The indirect effect through laziness was significant in both task groups. The exact values are in footnote ¶, which is not in the file. The reviewer's computation of a × b gives about −0.10 (manual) and −0.20 (digital).
- **Remaining effect of AI use once laziness is held constant:** negative in the manual task (c′ = −0.49, SE 0.10, p < 0.001), positive in the digital task (c′ = +0.39, SE 0.10, p < 0.001). The reviewer's check: −0.10 − 0.49 ≈ −0.59, close to the observed manual gap of −0.57; −0.20 + 0.39 ≈ +0.19, close to the observed digital gap of +0.17.
- **What this means:** when AI was described as useful and allowed for the task, the AI user was *still* seen as lazier, and that still cost him task fit. A separate gain of about the same size cancelled it, so the net task-fit gap was zero within error. "Offset" is the right word; "removed" is not.
- **Hiring (yes/no):** the pattern was the same in direction. The interaction between Tool and the participant's own AI use was not significant (p = 0.103) (md 85). No numbers for Tool × Task on hiring are given in the text.

**Measures worth borrowing**
- **Own AI use frequency.** Our measure list has field and cybersecurity knowledge but no AI use, although in two studies the judge's own AI use changed how AI users were judged. The item is a plausible moderator of trust in the checker and of whether participants name it. Two versions are available:
  - Study 4's five levels: "1 = has never used AI, 2 = tried AI once or twice, 3 = occasional use of AI, 4 = weekly use of AI, 5 = daily use of AI" (md 81).
  - Study 3's 1–10 scale, "1 (Never) to 10 (Multiple times per day)" (Fig. 3 caption, md 69).
- **Scale reliabilities (item wording on OSF, not in the file):**
  - laziness: 4 items, α = 0.94 (Study 4)
  - task fit: 3 items, α = 0.90–0.92
  - competence: α = 0.96 (Study 1) and 0.90 (Study 2)
  - willingness to tell the manager: 3 items, α = 0.91; willingness to tell colleagues: 3 items, α = 0.93
- **Exclusion rules set in advance** (md 119, 127, 137, 149):
  - identical answers to every rating item
  - minimum time on the study (30 s in Study 1, 45 s in Study 3)
  - a comprehension check (Study 3)

**Design elements**
- **Only the source varied.** In Study 2 the help was described word for word the same, and only its source changed (md 49). Our neutral group sees no flag at all. H1 as designed therefore tests "AI flag vs no flag", not "AI flag vs the same flag from a human reviewer". The abstract should not claim the drop is specific to AI unless a same-content, non-AI flag is added. Otherwise, state this as a limitation.
- **Timing of the first rating.** Study 4 rated laziness before giving task details "to capture participants' general impressions of the candidate independent of the task requirements" (md 151). This matches our trust rating after the profile and before the slides.
- **The description stayed visible.** The target's description was repeated on the rating page in Studies 2 and 4 (md 131, 153). We could do the same by keeping a short CISO profile visible on the rating screens.
- **Real stakes.** Study 3 paid managers according to the hired candidate's real performance (md 141–145). Our budget-approval measure is hypothetical; if it matters, a small bonus tied to the outcome would follow this model.

**Effect sizes, for scale**

The Cohen's d values are in Figs. 1–2, which are not in the file. The values below are the reviewer's own, computed from the reported means and SDs and approximate. In plain terms, the first line reads: the AI group expected to be seen as lazier by about a third of the spread of answers.
- Study 1:
  - expected laziness ≈ 0.37
  - expected replaceability ≈ 0.34
  - expected competence ≈ 0.54
  - expected diligence ≈ 0.43
  - willingness to tell the manager ≈ 0.22
  - willingness to tell colleagues ≈ 0.20
- Study 2: laziness ≈ 0.27 against non-AI help and ≈ 0.40 against no help mentioned.
- Study 4: task fit ≈ 0.43 for the manual task; ≈ 0.17 for the digital task, not significant.
- Background (the reviewer's calculation, not from the paper): with 30 per group, a two-group comparison has about 80% power only for differences of about d = 0.73 or larger. Effects of Reif's size (0.2–0.55) would need roughly 55 to 390 per group. These are different constructs from trust in a checked person, so they are not a basis for our main power calculation. They matter only for the choice of naming the checker, which in our design sits within the checker group (30 people × 3 flags = 90 choices).

**Stated limitations** (md 103–107)
- The samples were online convenience samples; the findings still need to be shown in established organisations.
- AI was described in general terms; judgments may depend on whether the AI augments work or automates it, and on whether it is embedded in a familiar tool.
- The data were collected between March 2024 and February 2025, so the penalty may shrink as AI becomes common.

**Future work the authors call for:** "future research could systematically examine how the type of support AI provides shapes social evaluations of its users" (md 105). Our checker is one type of support: it checks work rather than doing it. Our study would only carry out this call if we measured how the participant expects to be seen for using the checker (see direction (a)). As designed, it does not.

**Notes on the text file and on the paper's own wording**
- **Missing material.** Figures 1–4, the SI tables and figures, and the footnote text are missing. The footnote markers (*, †, ‡, §, ¶, #, ||) appear in the text with nothing behind them. As a result, the file lacks the Cohen's d values for Studies 1–2, the means for competence and the other traits in Study 2, the per-group laziness means in Study 4, and the size of the Study 4 indirect effects.
- **Study 3 sample size.** The Methods give 1,668 as the final sample (md 137), while the Fig. 3 caption gives N = 1,667 (md 69) and the test uses df 1664. The difference is minor.
- **Condition names.** Study 4's comparison group is called "Non-AI Tool" in the Results and "Traditional Tool" in the Methods. It is the same group.
- **Scale ranges.** The text does not state the scale ranges for Studies 1, 2 and 4.
- **The paper's summary sentences overstate Study 4.** Cite the numbers, not these two sentences:
  - "When AI is described as being useful for the task, candidates who use AI are perceived as having higher Task Fit" (md 87). The overall difference was not significant (p = 0.112). Only the effect with laziness held constant was positive.
  - "perceptions of laziness can be offset for tasks in which AI is clearly useful" (md 91). Laziness was rated before the task was described, so the task could not have changed it. What was offset was the task-fit penalty.

### 4. Check of our current sentence

Draft: "Reif et al. (2025) found that people who use AI at work are rated lazier and less competent and disclose it less willingly. We log whether participants name the checker when they question the CISO."

| Claim | Accurate? | Note |
|---|---|---|
| "people who use AI at work are rated lazier" | Accurate but incomplete | Study 2: observers rated written descriptions of employees. The finding is a comparison: AI-helped employees (2.50) were rated lazier than employees given the same help from a non-AI source (2.16) or with no help mentioned (2.02). The comparison should be named. |
| "and less competent" | Accurate | Study 2 (md 57). They were also rated less diligent, independent and self-assured. The text gives no means. In Study 1, lower competence was *expected* by the users, not rated by others. |
| "and disclose it less willingly" | Not accurate as worded | Study 1 measured how likely participants **said** they would be to tell their manager and colleagues, in an imagined scenario. No actual disclosure was observed. The difference was small (4.91 vs 5.25 for the manager, 4.85 vs 5.17 for colleagues). This is also a different study with different people (users imagining themselves) from the "rated lazier" result (observers). The sentence merges the two. |
| (omitted) | Missing qualifiers that matter for us | Raters who use AI weekly or daily did not rate AI users as significantly lazier (Study 4). The task-fit penalty was offset when AI was clearly useful and allowed for the task (Study 4). In Study 3, hiring did not differ on average; it depended on the manager's own AI use. |
| "We log whether participants name the checker when they question the CISO." | A design statement, consistent with our design | Linking it to Reif is an inference: Reif studied how AI users expect to be seen by managers and colleagues, not a board member using a checker in front of the person being checked. Also, "challenge" differs from "ask to explain" in tone as well as in naming the AI, so the log alone cannot separate the two. |

Suggested wording:

> Reif et al. (2025) found that employees described as getting help from generative AI were rated lazier and less competent than employees who got the same help from another source, and that people who imagined using an AI tool at work expected to be judged this way and said they were less likely to tell their manager and colleagues. Raters who used AI weekly or daily did not judge AI users as lazier. Challenging the CISO with the checker's finding shows that the board member relies on the AI, while asking the CISO to explain does not, so we log which of the two participants choose.

Shorter version, if space is tight:

> Reif et al. (2025) found that people who get help from AI at work are rated lazier and less competent than people who get the same help from another source, and that people expect this and say they are less likely to tell their manager. We therefore log whether participants name the checker when they question the CISO.

### 5. Bottom line

This paper does not test H1–H4: nobody in it is checked by an AI, the judged person never replies, and trust is not measured. It is worth citing in one role, as background for the action choice. It supports logging whether participants name the checker, and it motivates adding the participant's own AI-use frequency as a measure. The fact that matters most is Study 1: people who imagined using AI at work expected to be seen as lazier and less competent and said they were less likely to tell their manager and colleagues. However, raters who use AI weekly or daily did not judge AI users as lazier (Study 4), so among AI-savvy students the pull away from naming the checker may be small.
