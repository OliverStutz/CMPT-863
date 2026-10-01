# Reviewer C: Related work, abstract v4

I read the Research question, Hypotheses, Related work and References of `ABSTRACT-v4.md`, then the full texts of every paper cited in Related work and the Aamir sentences. I checked each claim against the paper first and used the earlier per-paper reports only as a second opinion. Locations give the section or page of each paper. Anything I say from outside the files is marked **[outside the files]**.

**Summary.** The facts are mostly right. Only one claim is wrong, and it is about our own design, not a paper. The weak part is the link to the hypotheses. H1 is grounded three times. H3 rests only on Chen's doctor, who stays silent where our CISO gives a denial. **H4 is not grounded at all**: the Reif paragraph says "This informs H4", but Reif bears on which action participants choose, which is neither H4a nor H4b. Score: **2 of 3**. Fixes 1 to 4 below would make it a 3, for about 150 more words, most of them in the new H4 paragraph.

---

## A. Accuracy, claim by claim

Verdicts: **Accurate**, **Imprecise** (true in direction, but it drops a qualifier a marker could catch, or overstates), **Inaccurate**.

### Aamir et al. (2026), in the Research question

| # | Claim | Verdict | Location | Correction / note |
|---|---|---|---|---|
| 1 | "In interviews with 13 Australian directors" | Accurate | §3.2 Participants; Table 2 | The participants were directors and senior executives from 13 organisations, all with a current board role or one held in the past five years. Nine of the 13 were executive directors. "Directors" is fair. |
| 2 | "gaps in cyber literacy and confidence make it hard for directors to question cyber-risk reports" | Accurate | Abstract ("limits their ability to interrogate risk reports, challenge assumptions") | — |
| 3 | "boards without such literacy often accept management's account" | Accurate | §6.2 ("Where this literacy was missing, boards often defaulted to accepting management narratives") | — |
| 4 | "directors rely on independent reviews to push back" | **Imprecise** | §4.1 ("Where external assurance was routine, directors reported greater confidence in challenging management perspectives"); §6.2 ("Yet not every board sought such input … preferred to rely on management") | This generalises to all directors. Better: "and directors felt more confident pushing back where independent reviews were routine." |
| 5 | "the board decides on the budget, and the CISO holds the expertise the board depends on" | Accurate as a paraphrase | §2.1 ("they often depend on CISOs to translate technical risk"); Figure 1 (board decisions on budgets) | The point that CISOs lack decision-making power comes from work Aamir et al. cite ([34]), not from their own interviews. |
| 6 | "Research on AI fact-checking measures belief in the checked claim." | **Imprecise** | Contradicted by our own cited papers: Chen §3.7.3 measured trust in the doctor after an AI's contrary opinion; Prike p. 12 rated the checked candidates | The few studies that rate the person rate once, at the end, and the person never answers. That is the real gap. **[outside the files, from SHORTLIST]** Chae & Tewksbury (2026) also rated politicians after AI-labelled fact-checks. |

### Strickland (1958)

| # | Claim | Verdict | Location | Correction / note |
|---|---|---|---|---|
| 1 | "40 students" | Accurate | p. 202 | All 40 were male undergraduates. |
| 2 | "play a supervisor who was told to check one of two workers after nine of ten work periods and the other after two" | Accurate | p. 203 (A after periods 1–4 and 6–10; B after 5 and 10) | The workers were fictitious. Worth saying, because our CISO is simulated too. |
| 3 | "Both workers produced the same output" | Accurate | Table 1, p. 204 (totals 52/52 and 50/50) | — |
| 4 | "every check found the work in order" | Accurate | p. 204 | In Group II each worker fell short once, but only in a period that was not checked. The supervisor saw this later on the slips. |
| 5 | "Afterwards, more supervisors trusted the rarely checked worker than the often checked one." | Accurate | Hypothesis 2, Table 3, pp. 206–207 | This was a single relative item (A more, B more, or equal), asked once after the checking, with ties split (p = .039 and < .001). |
| 6 | "Checking alone lowered trust while the work stayed the same." | **Imprecise (overstated)** | pp. 206, 213 | "Lowered" implies a before-and-after drop, but trust was measured once, relative to the other worker, with no unchecked condition. The author himself doubts "absolute distrust" (p. 213). The checker could also fine the workers, and the proposed cause (the supervisor credits his own checking) was only "indicated" (pp. 208–209). Better: "The often checked worker was trusted less, although his work was the same." |
| 7 | "Our slides contain no errors, so the groups differ only in the checking." | **Inaccurate** (about our design and the analogy) | Our Method: "The groups differ only in the checker, in the questions it proposes, and in who asks" | Our groups also differ in the **flags**, which claim errors. Strickland's checks never claimed any fault. The sentence hides the main difference between his study and ours. |

### Prike et al. (2025)

| # | Claim | Verdict | Location | Correction / note |
|---|---|---|---|---|
| 1 | "406 participants" | Accurate | p. 9 | Politically moderate U.S. adults on Prolific. |
| 2 | "two unfamiliar politicians through a profile and a series of posts" | Accurate | pp. 5–6, 11–12 | The candidates were fictional. Ratings were taken once, at the end (p. 16). |
| 3 | "a fact-checking organisation marked eight of one candidate's statements as false" | **Imprecise** | p. 8; Table 2, p. 13 | The eight false statements *favoured* one candidate. Some were his own posts, some came from a fictional newspaper, and some attacked his opponent (the loans and "communist organization" items). The checker also confirmed two true statements per candidate. It was "introduced as nonpartisan and objective" (p. 8). Better: "marked as false eight statements that favoured one candidate". |
| 4 | "participants rated him as less competent and less moral than his opponent" | Accurate | pp. 17–19 | The measure was the gap between the two candidates' ratings. The drop was large; the authors say "much less favorably" (p. 26). |
| 5 | "compared with a group that saw no fact-checks" | **Imprecise** | Table 1, p. 10; p. 16 | The control group saw only the 16 neutral posts: no false statements and no fact-checks. As written, the reader will assume the comparison group saw the same statements without checks, and that comparison is not reported. Better: "compared with a group that saw only neutral posts". |
| 6 | "We borrow the profile and the 0-to-10 rating." | Accurate | p. 12 (0–10 sliders, starting at the midpoint) | Their 0–10 scales rated competence, morality and warmth, not trust. |
| 7 | "Their fact-checks were correct and the candidate stayed silent" | Accurate | Table 2; no reply condition | The paragraph leaves out that their checker was a human organisation while ours is an AI tool in the rater's hands. |

### Chen et al. (2026)

| # | Claim | Verdict | Location | Correction / note |
|---|---|---|---|---|
| 1 | "the closest design" | A judgement, defensible | — | — |
| 2 | "*N* = 135" | Accurate | §3.6 (68 agree, 67 disagree, from 260 recruited) | — |
| 3 | "A simulated doctor told patients" | **Imprecise** | §3.1, §3.3, Appendix | The "doctor" was a GPT-3.5 chatbot presented as a human therapist, in a short online stress-management chat. The "patients" were online adults. Better: "a chatbot presented as a human doctor". |
| 4 | "that his AI assistant agreed or disagreed with his recommendation, and gave no explanation" | Accurate | Appendix, prompt steps 6–8 | The doctor first asked whether the patient wanted to hear the AI's opinion. |
| 5 | "After disagreement, patients felt less sure of their situation and saw the doctor as lazier" | **Imprecise** | §3 (two groups, agree vs disagree); §4.2.1, §4.2.3 | The comparison is with **agreement**; there was no group without the AI. Add "compared with agreement". "Less sure of their situation" is a fair plain rendering of "perceived medical uncertainty" (§2.7, §3.7.1). |
| 6 | "and both went with lower trust in the doctor" | Accurate | §4.2.1, §4.2.3 | This is an association, measured once after the chat. The paper reports no overall difference in trust between the groups, only effects running through each reason. "Went with" is the right wording. |
| 7 | "The authors suspect the missing explanation and suggest a short one, which they did not test." | Accurate | §5.2 ("may be due to the doctor not providing any explanations"); §5.1; §5.5 | Their example explanation accounts for *why the AI disagrees* and offers to review it together. It is not a rebuttal like our CISO's reason. Their first recommendation is not to disclose the disagreement openly at all (§5.5). |
| 8 | "There the expert brought in the AI himself." | Accurate | §1; Appendix step 6 | — |

### Greitemeyer and Sagioglou (2015), Ecker and Rodricks (2020), Mickelberg et al. (2025)

| # | Claim | Verdict | Location | Correction / note |
|---|---|---|---|---|
| 1 | "had participants rate a researcher accused of misconduct" | Accurate | pp. 4, 7 | He was accused of "reporting inconclusive data", and the ethics committee demanded a retraction. The rating combined three items: a trustworthy researcher, belief in his findings, and whether he deserves funding. |
| 2 | "and again after a report cleared him" | **Imprecise** | p. 4 ("it turned out that Sanna and his colleagues were wrongly accused. None of the allegations were confirmed … the article was not retracted") | No report and no source are named. Better: "after they read that he had been cleared". |
| 3 | "(*N* = 185 and 358)" | Accurate | pp. 3, 7 | — |
| 4 | "about two points lower than one never accused, on a scale from −5 to +5" | Accurate | Study 1: +1.54 vs −0.16 (p. 5); Study 2, matched control: +2.35 vs +0.38 (p. 8) | The gap was 1.7 points in Study 1 and 2.0 in Study 2. |
| 5 | "After the clearing he rose by about two points" | Accurate | +2.01 and +2.09 (pp. 6, 9) | — |
| 6 | "to the level of the never-accused researcher, but not reliably above it" | Accurate, with an omission | p. 6 (+1.85 vs +1.63, p = .536); p. 9 (+2.47 vs +2.27, p = .508) | Left out: in Study 2 the cleared researcher *did* end well above a never-accused researcher whose comparison group had not been told about his published article (+2.47 vs +0.96). The authors put that gap down to the extra information (pp. 9–10). This is the paper's most useful lesson for our design (see B). |
| 7 | "A retracted allegation likewise stopped affecting trait ratings of a person in Ecker and Rodricks (2020)" | Accurate | Exp. 1 and 2: likability, aggressiveness and predicted behaviour, retraction = control | Two things are left out. There was a tentative lasting effect on ratings of the person's *face* (Bayesian table). And in Exp. 2, likability after the retraction was slightly *above* control (p = .04, "marginal" in the authors' words). |
| 8 | "and in two of three experiments by Mickelberg et al. (2025)" | Accurate | Exp. 1 and 3: no difference; Exp. 2: a small lasting effect on likability; internal meta-analysis inconclusive | The measure was a likability scale. "Ratings of the person" is safer than "trait ratings". |
| 9 | "There a third party did the clearing." | Accurate | Greitemeyer p. 4; E&R "REJECTION" notice; Mickelberg "RETRACTION" notice | — |
| 10 | "recovery is expected, a gain would be new" | **Imprecise** | Greitemeyer pp. 6, 9; E&R Exp. 2 likability | Greitemeyer and Sagioglou tested the gain comparison and found none reliably; the cleared group was slightly higher both times (+0.2). E&R Exp. 2 showed a slight overshoot too. "New" overclaims. **[outside the files]** Marketing research on the "service recovery paradox" has long tested whether a well-handled failure leaves people better disposed than no failure at all, with mixed results. Better: "a reliable gain above the neutral group would go beyond these results". |

### Reif et al. (2025)

| # | Claim | Verdict | Location | Correction / note |
|---|---|---|---|---|
| 1 | "employees described as getting help from generative AI were rated lazier and less competent than employees getting the same help from another source" | Accurate | Study 2 (N = 1,203), md 55–57 | These employees were also rated below employees who got no help. |
| 2 | "people who imagined using AI at work said they were less likely to tell their manager" | Accurate, comparison missing | Study 1 (N = 497): 4.91 vs 5.25, md 45 | The comparison was with people imagining a non-AI dashboard tool. The difference was small. |
| 3 | "Raters who used AI weekly or daily did not judge AI users as lazier." | Accurate | Study 4, md 81 ("this difference was not significant for participants with high levels of AI use") | This is from Study 4 only, with a gig-worker candidate. |
| 4 | "Challenging the CISO with the checker's finding shows reliance on AI; asking the CISO to explain does not." | Our inference; reasonable | — | Reif studied AI doing a person's *own* work, not a tool for checking someone else's. The authors call for exactly that comparison (md 105). |

**Accuracy verdict:** Of the claims about the papers, 30 are accurate and 9 imprecise; none is wrong. The one inaccurate sentence (Strickland 7) is about our own design, and no number is wrong. Most imprecise claims share a pattern: the comparison group is not named (Prike 5, Chen 5, Reif 2), or a qualifier that matters is dropped (Strickland 6, Greitemeyer 10).

---

## B. Link to the hypotheses

| Paragraph | Stated link | Does it fit? | Can a marker follow it? | Does it say how we extend the paper? |
|---|---|---|---|---|
| Strickland | H1 | **Partly.** Strickland shows that *checking itself* can lower trust even when every check comes back clean. That is not a flag. It speaks to the presence of the checker, which is a rival explanation for the between-group part of H1 and for the RQ. It is also the reason the design shows "No issues found" on slides 1 and 5 (Validity section), but the paragraph never says so. | No. "The groups differ only in the checking" hides that our checker claims faults. | No. It draws a parallel ("our slides contain no errors") instead of naming the difference. |
| Prike | H1 | Yes, by analogy: a check lowered ratings of the checked person. | Yes. | Yes (wrong flags, the CISO answers). It misses the two biggest differences: an AI tool held by the rater instead of a human organisation, and ratings after every step instead of once. |
| Chen | H2, H3 | **H3: yes, but loosely.** Chen's disagreement condition is a contrary opinion followed by silence; ours is a flag followed by an explicit denial ("I'd rather not go into the details"). **H2: motivation only**, because the explanation the authors suggest is an account of the AI's view, not a rebuttal. H1 is also served, though not stated. | Mostly. | Partly ("we give the checker to the person in power and let the expert answer"). It misses the strongest extension. In Chen the disagreement and the silence always came together, so the study cannot tell which one lowered trust. Our ratings after the flag and after the answer separate the two. |
| Greitemeyer (+ two) | H2 | **Yes for recovery**: this is the only direct test of rating, doubt, clearing and rating again. For the gain it gives a boundary. | Yes. | Partly. "Here the CISO clears the flag in person" names the difference but not why it matters. A self-clearing may be believed less; the authors suggest the full recovery may be due to a believable source (p. 10). The paragraph also misses its best design link: the only time a cleared person beat a never-accused one, the comparison group had been told less. That is why our neutral group hears the same answers. |
| Reif | H4 | **No.** H4a is about how *often* the board member questions; H4b is about how *far* the answer moves trust. Reif bears on neither. It bears on *which* action is chosen (ask or challenge), which no hypothesis predicts, and on the stimulus rule that the CISO never mentions AI. The paragraph does not mention that rule. | No. A marker will not see how rating AI users as lazy leads to H4. | No. "We log which one participants choose" is a measure, not an extension. |

**H4 is served by no paragraph.** H4 was rewritten after the section, and the old link was never replaced. Grounding is available in the files already cited:
- **H4a.** Aamir et al. §4.1: "Where external assurance was routine, directors reported greater confidence in challenging management perspectives." A checker is a form of on-the-spot assurance. **[outside the files, from SHORTLIST]** von Schenk et al. (2024) is the direct evidence: with a lie-detection algorithm that was right about two times in three, people who asked for its verdict and heard "lie" accused far more often than people without it (about 85% against 19%).
- **H4b.** Chen et al. §5.1, citing Kempt & Nagel (2022): "a second opinion from an AI assistant can shift the burden of proof onto doctors" (also §2.8). After a flag, the CISO's answer has to carry that burden: a reason meets it and a denial fails it. In the neutral group nothing has to be proved. This grounds both directions of H4b. (Chen also cites expectancy violation theory, §2.5, but that supports only the upward half: in the neutral group a denial is itself a violation, so it predicts a larger drop there, not in the checker group.)

**H3 is thin.** The only support is the authors' own reading of Chen, and our denial is not Chen's silence. See C for the missing trust-repair literature.

**Balance.** H1 has three paragraphs (Strickland, Prike, Chen). H3 has half a sentence and H4 has none. Keeping Strickland and Prike short limits how much the section grows once H4 is added.

---

## C. Coverage

**Against the shortlist's coverage roles** (SHORTLIST.md): (a) a pre-AI baseline that checking lowers trust is filled by Strickland. (b) An AI contradicting a human is filled by Chen. (c) **False positives are not filled**, although every flag in our study is a false positive. von Schenk et al. (2024) was ranked 4th or 5th by all three shortlist reviewers for this role. (d) Behaviour change of the checked person (Nyhan & Reifler) is no longer needed, because only the board member is measured.

**Major work missing, from my own knowledge [outside the files; verify before citing]:**
1. **Trust repair after an accusation.** This is the literature closest to H2 and H3, and it is absent. Kim, Ferrin, Cooper & Dirks (2004, *Journal of Applied Psychology*) compared an accused job candidate apologising with denying, after a competence-based or an integrity-based accusation, and found the better reply depends on the kind of accusation. Ferrin, Kim, Cooper & Dirks (2007, *JAP*) added reticence (declining to confirm or deny) and found it a weak way to repair trust. Our flags are competence accusations. Our full answer is a denial backed by a reason, and our denial is a denial plus reticence. These studies measure what the cited papers do not: the accused person's own reply. Greitemeyer, Ecker & Rodricks and Mickelberg all have a third party do the clearing. A marker who knows organisational behaviour will notice the gap.
2. **Following fallible automated advice** (automation bias and algorithm appreciation; for example Logg, Minson & Moore, 2019). This bears on whether participants believe a wrong flag at all, which H1 needs. It is lower priority, and the end-of-session belief-in-flags measure already covers it empirically.
3. **Expectancy violation theory** (Burgoon, 1993). It is already quoted in Chen §2.5 and could be cited through Chen for the upward half of H4b.

**Would a listed alternative serve better?**
- **von Schenk et al. (2024) should replace Reif as the fifth discussed paper.** It fills role (c) and grounds H4a directly. Reif should stay as a single clause justifying the stimulus rule and the logging of the chosen action. von Schenk is open access (PMC11302104), so it can be read before the deadline.
- **Chae & Tewksbury (2026)** would strengthen, not replace, Prike. In three experiments with well-known politicians, neither AI-labelled nor human fact-checks changed feelings toward them or how competent they seemed (from SHORTLIST, not read). It adds the AI label, which Prike lacks, and it explains why we, like Prike, use a stranger. One clause would do, but only after reading it.
- **Arkes et al. (2007)** is not needed. It would add a second medical setting next to Chen, and Reif covers the penalty for using an aid.
- **Strickland** is old and small (40 students, a relative trust item), but it is the classic result and it justifies the "No issues found" design. Keep it, with the link fixed.

**Do the two passing citations earn their place?** As written, barely. They only repeat Greitemeyer's conclusion, and the rubric counts discussed papers, not citations. **Mickelberg et al. earns its place if used for its measurement point.** The authors argue that rating a person after every step gives the latest information extra weight. Their earlier study, which used repeated ratings, found full recovery, and a repeated rating is exactly our design (Mickelberg pp. 3, 15, citing the belief-adjustment model). Our within-person recovery (H2) may therefore look larger than a single end rating would show. The between-group comparisons are protected, because both groups rate at the same steps. **Ecker & Rodricks is optional.** Keep it for breadth if words allow, or cut it. Its slight overshoot in Exp. 2 is too weak to cite for the gain.

---

## D. Score and fixes

**Score: 2 of 3.** Five papers plus Aamir are discussed, the facts are mostly right, and each paragraph names a difference from our study, so much of the relevant work is present. But H4 has no grounding (the paragraph that claims it bears on which action is chosen), H3's denial rests only on Chen's silent doctor with the trust-repair literature missing, and Strickland's link to H1 is misstated, so the relationship to the project is weak in places.

Fixes 1 to 4 would bring it to a 3. Fixes 5 to 8 are refinements. The replacement paragraphs below are written to be pasted in this order: Strickland, Prike, Chen, Greitemeyer, H4. Together they come to about 625 words, against 475 now. The H4 paragraph accounts for most of the increase. Cut options are at the end.

### Fix 1: Replace the Reif paragraph with an H4 paragraph and keep Reif as a design clause
**Change:** Delete "This informs H4" and the sentence about which action shows reliance on AI. Ground H4a in Aamir and H4b in Chen's burden-of-proof point, then turn Reif into the reason for the stimulus rule.
**Replacement (whole paragraph, 150 words):**
> Two sources inform H4. Aamir et al. (2026) report that directors felt more confident challenging management where independent reviews were routine; the checker offers such a review on the spot, so we expect the checker group to question the CISO more often than the neutral group reports having questions (H4a). Chen et al. note that an AI's contrary opinion shifts the burden of proof onto the expert. After a flag, a reason meets that burden and a denial fails it, so the same answer should move trust further than in the neutral group, where nothing has to be proved (H4b). Reif et al. (2025) found that employees described as getting help from generative AI were rated lazier and less competent than employees getting the same help from another source, so our CISO never mentions AI, and we log whether participants name the checker and how often they use AI themselves.

**Optional, only after reading von Schenk et al.** Insert after the H4a sentence:
> In von Schenk et al. (2024), people who asked a lie-detection algorithm for its verdict and were told "lie" accused the speaker far more often than people without it, although the algorithm was right only about two times in three.

Add the reference: von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, *27*(7), Article 110201. https://doi.org/10.1016/j.isci.2024.110201

**Why:** H4 is a quarter of the hypotheses and currently has no support. A marker checking whether the "relationship is clear" will find that the Reif-to-H4 link does not hold. Both new grounds come from papers already cited and read: Aamir §4.1, and Chen §5.1 and §2.8, where the burden-of-proof point is the authors' reading of earlier work, which is why the text says "note". Attribute it that way.

### Fix 2: Correct the Strickland analogy and tie it to the "No issues found" design
**Change:** Replace the last three sentences ("Checking alone lowered trust …", "Our slides contain no errors …", "This informs H1.") and add "male" and "fictitious".
**Replacement (whole paragraph, 97 words):**
> Strickland (1958) had 40 male students play a supervisor who was told to check one of two fictitious workers after nine of ten work periods and the other after two. Both workers produced the same output, and every check found the work in order. Afterwards, more supervisors trusted the rarely checked worker than the often checked one. Our checker goes further: it claims faults that are not there. Because checking alone may lower trust, it also shows "No issues found" on slides 1 and 5, so that flagged and unflagged slides can be compared. This informs H1.

**Why:** "The groups differ only in the checking" is wrong, because the groups also differ in the flags. "Checking alone lowered trust" overstates a single relative rating with no baseline (p. 213). The replacement turns the paper into the reason for a design choice the Validity section already relies on.

### Fix 3: Make the Greitemeyer paragraph carry its design lesson, and give Mickelberg a job
**Replacement (whole paragraph, 161 words):**
> Greitemeyer and Sagioglou (2015) had online participants rate a researcher accused of misconduct, and again after they read that he had been cleared (*N* = 185 and 358). The accused researcher was rated about two points below one never accused, on a scale from −5 to +5; after the clearing he rose by about two points, to that level but not reliably above it. He ended clearly above a never-accused researcher only when that researcher had been described with less information, which is why our neutral group hears the same answers. In Mickelberg et al. (2025), a retraction removed the effect of an allegation on how likeable a person seemed in two of three experiments; the authors note that rating after every step may give the latest information extra weight, and both our groups rate at the same steps. In these studies a third party did the clearing; our CISO answers for himself. This informs H2: recovery is expected; a reliable gain would go beyond these results.

This version drops Ecker and Rodricks (2020). Remove its reference too, or put it back as "in Ecker and Rodricks (2020) and" (6 words) if the list does not look long.

**Why:** It fixes "a report cleared him" (no report is named, p. 4) and "a gain would be new" (Greitemeyer and Sagioglou tested the gain). It gives the paper's most useful lesson for our design, the Study 2 comparison with the group told less (pp. 9–10). And it gives the passing citation a job that bears on our step-by-step ratings, instead of repeating the main finding (Mickelberg pp. 3, 15).

### Fix 4: Name Chen's comparison and state the separation our design achieves
**Replacement (whole paragraph, 126 words):**
> Chen et al. (2026) is the closest design (*N* = 135). At the end of a short online therapy chat, a chatbot presented as a human doctor said only that his AI assistant agreed or disagreed with him. Compared with agreement, disagreement left participants less sure about their own situation and made the doctor seem lazier, and both went with lower trust in him. Because the doctor never answered, the study cannot separate the disagreement from the silence; the authors suspect the silence and suggest an explanation, which they did not test. There the doctor brought in the AI himself. We give the checker to the board member and rate trust after the flag and after the answer: H2 tests an answer with the company's reason, H3 a denial.

**Why:** "After disagreement" reads as a drop from a baseline, but the comparison is with agreement (§3). "Simulated doctor" hides that it was a chatbot (§3.1). The new sentence about separating disagreement from silence is the clearest statement of how we extend the closest paper, and it gives H3 a proper link.

### Fix 5: Name the trust-repair literature for H2 and H3 (only after reading) [outside the files]
**Change:** Add one sentence at the end of the Greitemeyer paragraph.
**Replacement (draft; check it against the two abstracts before use):**
> Work on trust repair has tested the accused person's own reply: Kim et al. (2004) compared apology with denial after an accusation of incompetence or dishonesty, and Ferrin et al. (2007) found that declining to comment repaired trust poorly; our denial combines a denial with declining to explain.

**Why:** H2 and H3 are about how the accused person's own reply changes trust. All the cited papers have a third party do the clearing, and this is the literature that tests the reply itself. If there is no time to read these two papers before the deadline, leave them out rather than cite them unread (the AI transparency statement promises the team reads every cited paper).

### Fix 6: Tighten Prike's comparison and name the AI difference
**Replacement (whole paragraph, 91 words):**
> Prike et al. (2025) introduced 406 U.S. adults to two unfamiliar, fictional candidates through a profile and a series of posts. When a nonpartisan fact-checking organisation marked as false eight statements that favoured one candidate, participants rated him as less competent and less moral than his opponent, compared with a group that saw only neutral posts. We borrow the profile and the 0-to-10 rating. Their checker was human and right, and the candidate stayed silent; ours is an AI tool, wrong for this company, and the CISO answers. This informs H1.

**Optional, after reading Chae & Tewksbury.** Add:
> With well-known politicians, neither AI-labelled nor human fact-checks changed how they were rated (Chae & Tewksbury, 2026), which is why we, like Prike et al., use a person participants do not know.

**Why:** "A group that saw no fact-checks" suggests the same statements shown without checks, but the control group saw no false statements at all (Table 1). "Eight of one candidate's statements" is not what the paper did (p. 8). The AI source is the most obvious way we extend Prike, and the current text leaves it out.

### Fix 7: Aamir and the gap sentence in the Research question
**Replacements:**
- "…and that directors rely on independent reviews to push back." → "…and that directors felt more confident pushing back where independent reviews were routine."
- "Research on AI fact-checking measures belief in the checked claim. We measure trust in the checked person, step by step." → "Most studies of fact-checking measure belief in the checked claim. The few that rate the checked person, such as Prike et al. (2025) and Chen et al. (2026), rate once, at the end, and the person never answers. We rate trust in the person after every step, and the person answers."

**Why:** Not every board sought independent reviews (§6.2). The gap sentence is contradicted by two of our own cited papers. The replacement states the actual gap and sets up the whole section.

### Fix 8: Carry the changes to the short version, and fix the wording errors
- The short Related work in Working notes §10 needs the same H4 grounding: one sentence, Aamir for H4a and Chen's burden of proof for H4b. Its Chen sentence should say "compared with agreement".
- In the Research question, "A llm assisted camera-capture checker" should read "An LLM-assisted camera checker". Delete "The power is in the boards hands while the expertise sits at the CISO.": it repeats the sentence before it and lacks the apostrophe.

**Why:** These are small, but they are the first things a marker reads.

### Cut options if length matters (in order)
1. Move Reif to the Method sentence "The CISO never uses or mentions AI", as "because people who get help from AI at work are rated lazier and less competent (Reif et al., 2025)". This saves about 40 words in Related work. Four papers are still discussed, plus Aamir, which meets the brief.
2. Drop "Two sources inform H4." (4 words) and "There the doctor brought in the AI himself." (8 words; "We give the checker to the board member" still carries the contrast).
3. Drop the Mickelberg sentence (about 40 words). This loses the measurement caveat and the passing citation.
