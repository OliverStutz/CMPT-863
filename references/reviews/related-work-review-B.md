# Reviewer B: Related work in abstract v4

**Score: 2 of 3.** The five papers are relevant, and most claims check out against the full texts. H1 to H3 each have a link a marker can follow. H4, which was rewritten after the section, has no paper behind it. The Reif paragraph says "This informs H4", but it bears on which question a participant asks, not on whether they question more (H4a) or whether answers count for more (H4b). There are also four factual slips, the largest in the Prike paragraph, where the comparison group is wrong. The sentence in the Research question that states the gap is contradicted by Chen et al., which the abstract itself cites. With fixes 1 to 6 below, the section would earn a 3.

Sources: every claim was checked in the full texts in `papers/`. The per-paper reports in `reports/` were used as a second opinion only; I agree with them on every point I checked. Anything from outside the files is marked **[outside the files]**.

---

## A. Accuracy, claim by claim

Verdicts: **Accurate**, **Imprecise** (true in substance but missing a qualifier a marker could catch), **Inaccurate**.

### A0. Aamir et al. (2026), in the Research question

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | "interviews with 13 Australian directors" | Accurate | Abstract; §3.2: 13 people, 9 executive and 4 non-executive directors, in role at the time or within the past five years | none |
| 2 | "gaps in cyber literacy and confidence make it hard for directors to question cyber-risk reports" | Accurate | Abstract: the deficit "limits their ability to interrogate risk reports, challenge assumptions" | none |
| 3 | "boards without such literacy often accept management's account" | Accurate | §6.2 Capability: "Where this literacy was missing, boards often defaulted to accepting management narratives." | none |
| 4 | "directors rely on independent reviews to push back" | Imprecise | §6.2 Challenge: external assurance "gave directors confidence to challenge assumptions", but "not every board sought such input… A few relied almost entirely on management updates"; smaller boards found reviews "costly or disruptive" | "and that directors who used independent reviews, such as audits and red-team tests, said these gave them the confidence to challenge management." |
| 5 | "A llm assisted camera-capture checker offers such a review on the spot." | Framing, fair | Aamir's reviews are periodic and done by people (the §5 checklist: "Commission an independent annual challenge"). The next sentence already says the checker lacks context. | Typos: "An LLM-assisted", and "in the board's hands" in the last sentence of the paragraph |
| 6 | "Research on AI fact-checking measures belief in the checked claim." | **Inaccurate** | Chen et al. (cited in this abstract) measured trust in the checked doctor. Prike et al. measured competence and morality of the checked candidate. According to SHORTLIST.md, Chae & Tewksbury (2026) measured feelings toward, and competence of, politicians checked by fact-checks labelled as AI. | See fix 2 |

### A1. Strickland (1958)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | 40 students | Accurate | p. 202: "40 male students at the University of North Carolina" | none |
| 2 | played a supervisor | Accurate | p. 202 | none |
| 3 | told to check one worker after nine of ten work periods and the other after two | Accurate | p. 203: A after periods 1–4 and 6–10, B after 5 and 10; the experimenter set the schedule | none |
| 4 | both workers produced the same output | Accurate | Table 1: totals 52 and 52 (Group I), 50 and 50 (Group II). The workers were fictitious (p. 203). | Optional: "two workers, who were in fact fictitious" |
| 5 | every check found the work in order | Accurate | p. 204: in Group II each worker fell short once, but only in a period that was not checked | none |
| 6 | afterwards, more supervisors trusted the rarely checked worker than the often checked one | Accurate | Hypothesis 2, Table 3, pp. 206–207. The item was relative (A more, B more, or equal), asked once at the end of Part I. | none |
| 7 | "Checking alone lowered trust while the work stayed the same." | **Imprecise** | There was no trust rating before checking and no group without checking. The measure was relative. Strickland himself, p. 213: "It is questionable whether or not S actually had something like an absolute distrust of A… since he evaluated his subordinates relative to one another." | "The often-checked worker was trusted less than the rarely checked one, although both did the same work." |
| 8 | "Our slides contain no errors, so the groups differ only in the checking." | **Inaccurate as an analogy** | Our groups differ in the checking and in three wrong flags. Strickland's checks were clean and never claimed a fault. | See fix 4 |

### A2. Prike et al. (2025)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | 406 participants | Accurate | p. 9 (455 recruited, 49 excluded) | none |
| 2 | two unfamiliar politicians, introduced through a profile and a series of posts | Accurate | pp. 5–6, 11–12: two fictional candidates in Malta; profile with name, party, headshot; social-media simulator | none |
| 3 | "a fact-checking organisation marked eight of one candidate's statements as false" | **Imprecise** | p. 8: eight false statements "all associated with one of the candidates, designed to enhance that candidate's perceived competence and morality or reduce the perceived competence and morality of their opponent". They came from the candidates or from a fictional newspaper. The checker also marked two true statements per candidate as true. | "a fact-checking organisation marked as false eight posts that favoured one candidate" |
| 4 | participants rated him as less competent and less moral than his opponent | Accurate | pp. 17–19. The analysis used the difference between the two candidates. Warmth and votes fell too (pp. 21–23). | none |
| 5 | "compared with a group that saw no fact-checks" | **Inaccurate** | pp. 16–18: every comparison was with the control group, which saw only the 16 neutral posts, with no false posts and no fact-checks. Two other groups saw the same false posts without fact-checks, but they were not the comparison. The reader will think of those groups. | "compared with a control group that saw only the neutral posts" |
| 6 | "We borrow the profile and the 0-to-10 rating." | Accurate | p. 12: competence and morality on 0–10 sliders set at the midpoint | Add that the design also borrows Prike's end-of-study belief rating, with the control group as a baseline (p. 15). The Measures section uses it. |
| 7 | their fact-checks were correct; the candidate stayed silent | Accurate | Table 2: each false statement is corrected with the facts. No candidate replies. | none |
| 8 | (not stated) ratings taken once, at the end | Missing | p. 16: impressions, vote and belief were rated after the whole simulation | Worth adding as a way our study extends theirs (fix 3) |

### A3. Chen et al. (2026)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | "the closest design" | Fair judgement | SHORTLIST.md: the only paper in the pool in which an AI contradicts a professional and trust in that person is the outcome | none |
| 2 | N = 135 | Accurate | §3.6 (260 recruited, 125 excluded) | none |
| 3 | "A simulated doctor told patients…" | Accurate, could be sharper | §3.1: a GPT-3.5 chatbot presented as "Dr. Alex"; participants were told it was a chatbot only at the debrief (§3.3). "Patients" were online panel participants in a free stress-management session (§3.3–3.4). | "An LLM chatbot, presented as a human doctor, held a short stress-management session with online participants" |
| 4 | he said his AI assistant agreed or disagreed with his recommendation, with no explanation | Accurate | Appendix, steps 6–7; §5.6 lists the wordings ("CareBot disagrees with me") | none |
| 5 | "After disagreement, patients felt less sure of their situation and saw the doctor as lazier" | **Imprecise** | There were only two groups, disagreement and agreement (§3). No group heard no second opinion. Agreement itself raised the credibility of the recommendation (§4.2.2, §5.3), so the gap mixes a gain with a loss. | "Compared with agreement, disagreement left participants less sure about their situation and made the doctor seem lazier" |
| 6 | "both went with lower trust in the doctor" | Accurate | §4.2.1 (laziness) and §4.2.3 (uncertainty): both linked to lower cognitive, affective and behavioural trust. No overall difference in trust between the groups is reported. | none |
| 7 | "The authors suspect the missing explanation" | Accurate | §5.2: "This finding may be due to the doctor not providing any explanations after the AI's disagreement" | none |
| 8 | "and suggest a short one, which they did not test" | Accurate, one nuance | §5.1 and §5.5: they suggest explaining *why the AI disagreed*, or raising it indirectly. §5.5 also recommends that doctors "avoid explicitly disclosing" a disagreement. Our CISO explains his own decision instead. | "and suggest explaining the disagreement, which they did not test" |
| 9 | "There the expert brought in the AI himself." | Accurate | Abstract; Appendix step 6: the doctor offers CareBot's opinion and the participant accepts | none |

### A4. Greitemeyer & Sagioglou (2015), with Ecker & Rodricks (2020) and Mickelberg et al. (2025)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | participants rated a researcher accused of misconduct | Accurate | pp. 4, 7: "accused of reporting inconclusive data", and his university's ethics committee "demanded the retraction" of an article. The rating was three items from −5 to +5: is he a trustworthy researcher, do you believe his findings, should he get further funding. | Optional: name the three items |
| 2 | "and again after a report cleared him" | **Imprecise** | p. 4: participants read "that it turned out that Sanna and his colleagues were wrongly accused. None of the allegations were confirmed and so, eventually, the article was not retracted." No report or source is named. | "and again after they read that he had been wrongly accused" |
| 3 | N = 185 and 358 | Accurate | pp. 3, 7 | none |
| 4 | the accused researcher was rated about two points lower than one never accused, on a scale from −5 to +5 | Accurate (rounded) | Study 1: +1.54 vs −0.16, 1.7 points (p. 5). Study 2, against the control group that read the same article summary: +2.35 vs +0.38, 2.0 points (p. 8) | none |
| 5 | after the clearing he rose by about two points | Accurate | −0.16 to +1.85 (p. 6); +0.38 to +2.47 (p. 9) | none |
| 6 | to the level of the never-accused researcher, but not reliably above it | Accurate | +1.85 vs +1.63, p = .536 (p. 6); +2.47 vs +2.27, p = .508 (p. 9) | "not clearly above it" is plainer |
| 7 | (not stated) the cleared researcher did end above a never-accused one who was described with less information | Missing, and it matters | Study 2 (p. 9): +2.47 vs +0.96 for the control group that did not read the article summary. The authors put the gap down to the extra information, not to the clearing (p. 10). | This is the reason the neutral group must hear the same answers. See fix 5. |
| 8 | "A retracted allegation likewise stopped affecting trait ratings of a person in Ecker and Rodricks (2020)" | Accurate in substance | Two experiments (N = 142 and 287). A retracted allegation (he slapped his girlfriend) left ratings of likability and aggressiveness, and predictions of behaviour, at the level of the control group (Exp. 1 and 2 results; abstract). The authors found tentative signs of a lasting effect on ratings of his face. People were rated once, at the end, and groups were compared, so "stopped affecting" is loose. | "left no lasting mark on how likable and how aggressive the person seemed" |
| 9 | "…and in two of three experiments by Mickelberg et al. (2025)" | Accurate, measure loosely named | Abstract; Exp. 1 and 3 found no lasting effect, Exp. 2 did. The measures were a likability scale and inference questions, not trait ratings. The authors' own equivalence tests for Exp. 1 and 3 were inconclusive. | "and in two of three experiments by Mickelberg et al. (2025)", with "likability" in place of "trait ratings" |
| 10 | "There a third party did the clearing." | Accurate | Greitemeyer: an unnamed outcome. Ecker & Rodricks, Mickelberg: an update from the person's acquaintances. | none |

### A5. Reif et al. (2025)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | employees described as getting help from generative AI were rated lazier and less competent than employees getting the same help from another source | Accurate | Study 2 (N = 1,203). Also less diligent, independent and self-assured, and also lower than employees described with no help. | none |
| 2 | "people who imagined using AI at work said they were less likely to tell their manager" | **Imprecise** | Study 1 (N = 497). The comparison is with people who imagined a non-AI dashboard tool: 4.91 vs 5.25 on a 7-point scale. | "said they were less likely to tell their manager than people who imagined using a non-AI tool" |
| 3 | raters who used AI weekly or daily did not judge AI users as lazier | Accurate | Study 4 (hiring a gig worker): the difference "was not significant for participants with high levels of AI use (weekly or daily)" | "did not clearly judge" is more exact |
| 4 | "Challenging the CISO with the checker's finding shows reliance on AI; asking the CISO to explain does not." | Our inference, stated as fact | Reif's AI users used AI for their own work. Nobody in the paper uses AI to check someone else's work. | "Naming the checker may look like relying on AI" |

---

## B. Link to the hypotheses

**Strickland → H1.** The link is right but badly worded. Strickland shows the effect of *being checked*, with checks that found nothing. H1 is about a *flag*. The paragraph's own bridge ("so the groups differ only in the checking") hides that difference. The paper has a better link that the design already uses: the checker reports "No issues found" on slides 1 and 5, and the Validity section compares flagged with unflagged slides "which separates the flags from the presence of the checker". Strickland is the reason for that comparison. There is a second link the RQ would welcome: his hypothesis explicitly requires that the checker "has high power vis-à-vis a subordinate" (p. 202), which matches "the power is in the board's hands". **How our study extends it:** not stated. The paragraph gives only similarities.

**Prike → H1.** The link fits and is easy to follow: an unfamiliar person loses standing when a third party marks their statements false. **Extension:** stated (wrong flags, the CISO answers). It could add that Prike rated once at the end, while we rate after every step. One nuance for the authors: Prike's candidate was caught spreading falsehoods, which explains the drop in *morality*. Our flags question competence, so only the competence part of Prike's result carries over.

**Chen → H2 and H3.** This is the clearest paragraph. The closing sentence ("H2 tests the explanation, H3 the non-answer") tells the marker exactly what we add. It can be made sharper for H3. In Chen, the disagreement and the doctor's silence came together, so the loss of trust cannot be split between them. Our ratings after the flag and after the answer split them. That is the precise way the study extends Chen. The Working notes (§7) also list Chen as support for H4 via "burden of proof". That is Chen citing Kempt & Nagel (§5.1), not a finding, so do not use it for H4.

**Greitemeyer (+2) → H2.** The link fits, and "recovery is expected, a gain would be new" is exactly the right framing. **Extension:** stated (the CISO clears the flag himself). The paragraph leaves out the paper's most useful point for H2's "above the neutral group" clause. In Study 2 the cleared researcher *did* end above a never-accused researcher, but only above the one about whom participants had been told less. The authors traced that gap to the extra information. That is the design reason our neutral group hears the same answers (Working notes §1 says so: "Greitemeyer and Sagioglou fell into the same trap"). One sentence in Related work would make the H2 "gain" test look deliberate rather than hopeful.

**Reif → H4.** **This does not fit.** H4 now says (a) the checker group questions the CISO at more slides and (b) the same answer moves trust further in the checker group. Reif speaks to neither. At most it bears on *which* question a participant chooses (ask or challenge). It also backs two design choices: the CISO never mentions AI (Working notes §4), and H4a counts asking and challenging together. A marker who reads "This informs H4" will look for the connection and not find it. That is the "relationship to this project is weak" case in the rubric.

**Is H4 served by any paragraph?** No. H4a has an analogue already in the abstract: Aamir's directors said independent reviews gave them the confidence to challenge management. It is in the Research question, though, and never tied to H4. von Schenk et al. (2024), from the shortlist, is direct evidence for H4a (see C). For H4b, nothing in the files tests it. Note for the authors: if H1 and the recovery part of H2 hold, the checker group answers from a lower rating, so a larger rise after the reason follows almost automatically. The part of H4b that needs its own grounding is the larger *fall* after the denial. An honest sentence ("we found no study that tests this") is better than a stretched citation.

**H3** rests only on the Chen authors' own reading. That is acceptable but thin. See C for the literature that tests it directly.

---

## C. Coverage

**What is well covered.** H1 has three papers (Strickland, Prike, Chen) plus Aamir, possibly more than it needs. The recovery part of H2 is well grounded. The boardroom setting has Aamir. SHORTLIST.md confirms that no paper on trust in a person in a boardroom or cyber setting was found.

**What is missing.**

1. **Evidence for H4a: von Schenk et al. (2024), from the shortlist.** According to SHORTLIST.md (the paper is not in the files), people had access to a lie-detection algorithm that was right about two times in three. About 31% asked for its verdict. Among those told "lie", about 85% accused the author, against 19% without the algorithm. That is the closest published analogue to "a fallible checker makes the holder question more", and it is about AI. Caveats to state: the people who asked chose themselves, and the outcome is accusing, not trust. **Read it before citing.**
2. **How an accused person's own reply repairs trust [outside the files].** H2 and H3 are about what the CISO says back. The trust-repair literature in organisational psychology tests exactly this: participants watch a job candidate answer an allegation, and trust in the candidate is measured. Kim, Ferrin, Cooper & Dirks (2004, *Journal of Applied Psychology*, 89(1), 104–118) compared apology with denial and found that which one works better depends on whether the allegation concerns competence or integrity. Ferrin, Kim, Cooper & Dirks (2007, *Journal of Applied Psychology*, 92(4), 893–908, "Silence speaks volumes") added reticence, a refusal to discuss the allegation. As I recall, they found it repaired trust less well than the better of the other two responses. Our denial ("I'd rather not go into the details here") is a mix of denial and reticence, and our flags question competence. This is the most directly relevant body of work that is absent. A marker who knows it would call it major. I cite these from memory; check the findings before using them.
3. **Counter-evidence on AI-labelled fact-checks: Chae & Tewksbury (2026), from the shortlist.** According to SHORTLIST.md, three experiments (N = 3,978) found that neither AI-labelled nor human fact-checks changed feelings toward, or perceived competence of, the politicians being checked. The targets were well known. One clause next to Prike would show that the authors know the null result and why a stranger CISO differs. It also contradicts the RQ sentence "Research on AI fact-checking measures belief in the checked claim" (fix 2).
4. **A theory for the "gain" and for H4b (optional, in the files).** Chen et al. (§2.5) draw on expectancy violations theory (Burgoon, 1993): doing better than expected leaves a better impression than simply meeting expectations. That gives H2's "above neutral" clause and H4b a reason, even without data. One clause is enough.
5. **Optional, for the setting.** Aamir et al. cite Opdenbusch, Hielscher & Sasse (2025, NDSS), "Where Are We On Cyber?", on directors' limited technical confidence and information gaps. A second boardroom source would strengthen the Research question. [outside the files for its content]

**Would a listed alternative serve better than a cited paper?**
- **Reif → von Schenk for the H4 slot.** Reif studies how people who *use* AI for their own work are judged. Nobody in our study judges the board member, so Reif's finding lands only on a design rule. von Schenk speaks directly to H4a. Keep Reif as one clause for the design rule.
- **Arkes et al. (2007)** does not beat Chen. It shares the medical setting, its outcome is rated diagnostic ability rather than trust, and its point (a professional who uses an aid is downgraded) overlaps Reif. Not needed.
- **Strickland** is old and uses a relative measure, but it is the only direct evidence that being checked lowers trust when performance is equal. It justifies the flagged-vs-unflagged comparison. Keep it.

**Are the two passing citations earning their place?** Only just. They show that full recovery is the usual finding, but:
- both involve moral allegations about a private person (slapping a girlfriend, an affair, kicking a dog), retracted by an outside source, and rated once;
- neither bears on a gain or on a person answering for themselves;
- Ecker & Rodricks adds nothing that Greitemeyer does not already show.

Mickelberg et al. can earn its place for a different reason that concerns our design. Their introduction and discussion raise the possibility that rating a person repeatedly, as the task goes on, draws attention to the latest information and so favours recovery. Their own results were mixed on this. Our within-person recovery (H2) uses exactly such repeated ratings. Comparisons between the groups are protected because both groups rate at the same points. My recommendation: cut Ecker & Rodricks. Either cut Mickelberg too, or keep it for the repeated-rating caution (fix 9).

---

## D. Score and fixes

**Score: 2.** The section discusses five relevant papers accurately enough, and H1 to H3 each have a stated link and a stated extension. However, H4 has no supporting paper (the Reif paragraph's link does not match the rewritten H4), the literature on how an accused person's own explanation or denial affects trust is absent, and three claims (the Prike comparison group, Strickland's "lowered trust", the RQ's gap sentence) would not survive a marker who opens the papers.

### Fixes, most important first

**1. Replace the Reif paragraph so that H4 is served.**
Replacement (uses von Schenk; read it first):
> A checker may also change what the board member does. Directors in Aamir et al. (2026) said independent reviews gave them the confidence to challenge management. In von Schenk et al. (2024), people who consulted a lie-detection algorithm that was right about two times in three, and were told a statement was a lie, mostly accused its author; without the algorithm, about one in five did. This informs H4a. We found no study of whether the same answer moves trust further after a flag than without one (H4b). Reif et al. (2025) found that employees described as getting help from generative AI were rated lazier and less competent than employees getting the same help from another source. So the CISO never uses or mentions AI, and because naming the checker may look like relying on AI, H4a counts asking and challenging together.

If there is no time to read von Schenk, delete its sentence. Aamir alone then carries H4a, and the rest stands.
*Why:* H4 is the only hypothesis with no paper, and "This informs H4" under Reif is the weakest link in the section. The replacement ties Reif to what it actually supports (a design rule and the way H4a is counted) and says openly that H4b is new.

**2. Correct the gap sentence in the Research question.**
Replace "Research on AI fact-checking measures belief in the checked claim. We measure trust in the checked person, step by step." with:
> Studies of fact-checks and of AI second opinions have measured how the checked person is seen (Prike et al., 2025; Chen et al., 2026), but the checked person never answered, and ratings were taken once, at the end. We let the CISO answer and measure trust after every step.

*Why:* The current sentence is contradicted by Chen et al., whom the abstract calls "the closest design" and who measured trust in the checked doctor. The replacement states the real gap, which the Working notes (§7) already name: nobody has let a professional answer an AI flag.

**3. Prike: correct what was flagged and what it was compared with.**
Replace sentences 2–3 with:
> In some groups, eight false posts favoured one candidate, and a fact-checking organisation marked each of them false. At the end, participants rated that candidate as less competent and less moral than his opponent, compared with a control group that saw only neutral posts, and fewer voted for him. We borrow the profile, the 0-to-10 rating and the final rating of how true each flagged claim is. Their fact-checks were correct, the candidate stayed silent, and ratings were taken once, at the end; our flags are wrong for this company, the CISO answers, and trust is rated after every step. This informs H1.

*Why:* "compared with a group that saw no fact-checks" points the reader to the wrong group (pp. 16–18). The false posts were not all the candidate's own statements (p. 8). Adding "rated once, at the end" gives a second, concrete way we extend the paper.

**4. Strickland: soften the finding, fix the analogy, state the extension.**
Replace from "Afterwards…" to the end with:
> Afterwards, asked which worker they trusted more, more supervisors named the rarely checked one than the often checked one, although both had done the same work. Like Strickland's supervisors, our board member holds power over the person checked, and our slides, like the workers' output, contain no errors. But Strickland's checks found nothing and the workers never spoke; our checker raises wrong flags and the CISO answers. The checker also reports "No issues found" on slides 1 and 5, so we can separate being checked from being flagged. This informs H1.

*Why:* "Checking alone lowered trust" implies a drop from a baseline that was never measured; the item was relative, and Strickland says so himself (p. 213). "The groups differ only in the checking" is not true of our design. The new version ties the paper to the power asymmetry in the RQ and to the flagged-vs-unflagged comparison in Validity.

**5. Greitemeyer: fix "a report", add the lesson about extra information.**
Replace the first sentence and add one sentence before "There a third party…":
> Greitemeyer and Sagioglou (2015) had participants rate a researcher accused of misconduct on how trustworthy he was, whether they believed his findings and whether he should be funded, and rate him again after they read that he had been wrongly accused (N = 185 and 358). […] In the second study he did end above a never-accused researcher about whom participants had been told less; the authors trace that gap to the extra information. This is why our neutral group hears the same answers.

*Why:* No report is named in the paper (p. 4). The added sentence explains the one design choice that makes the H2 "gain" test clean, and a marker will see that the authors learned from the paper rather than just citing it.

**6. Chen: name the comparison, and say what the design splits.**
Replace "After disagreement, patients felt…" with "Compared with agreement, disagreement left participants less sure about their situation and made the doctor seem lazier". Replace "There the expert brought in the AI himself." with:
> There the doctor brought in the AI himself, and the disagreement and his silence came together, so the loss of trust cannot be split between them. We give the checker to the person in power, rate trust after the flag and again after the answer, and vary the answer: H2 tests the explanation, H3 the non-answer.

Optional: "An LLM chatbot, presented as a human doctor, held a short stress-management session with online participants" in place of "A simulated doctor told patients".
*Why:* There was no group without a second opinion (§3). Naming the confound between the disagreement and the silence is the sharpest statement of how our design extends Chen, and it gives H3 a proper link.

**7. Add the trust-repair literature for H2 and H3 [outside the files; read first].**
One sentence after the Chen or Greitemeyer paragraph, for example:
> Studies of trust repair show that how an accused professional answers matters: in Ferrin et al. (2007), a job candidate who declined to discuss an allegation repaired trust less well than one who gave the better of an apology or a denial.

*Why:* H2 and H3 test the accused person's own reply, and this is the literature that tests it. Its absence is the most likely "major related work is missing" objection. Do this only if the paper can be read before submission. The wording above comes from my memory of the paper.

**8. Aamir: soften "rely on"; fix typos.**
Replace "and that directors rely on independent reviews to push back" with "and that directors who used independent reviews, such as audits and red-team tests, said these gave them the confidence to challenge management". Change "A llm assisted" to "An LLM-assisted" and "the boards hands" to "the board's hands".
*Why:* Aamir §6.2 says not every board sought such reviews, and some avoided them as costly. The statements are uncounted self-reports.

**9. Passing citations: cut or give them a job.**
Either delete the Ecker & Rodricks / Mickelberg sentence, or replace it with:
> Retracted allegations about a person's behaviour also left no lasting mark on how likable he seemed in most such experiments (Mickelberg et al., 2025, two of three), although these authors caution that rating someone repeatedly may draw attention to the latest information and so favour recovery; both our groups rate at the same points.

and drop Ecker & Rodricks from the references.
*Why:* As written, they repeat Greitemeyer for a different kind of allegation. "Trait ratings" misnames Mickelberg's measure. The caution about repeated ratings is the one thing they add that bears on this design.

**10. If Reif's Study 1 stays anywhere: add the comparison.**
"said they were less likely to tell their manager than people who imagined using a non-AI tool".
*Why:* Study 1 compared AI with a non-AI dashboard tool (4.91 vs 5.25 out of 7); without the comparison the claim sounds absolute.

### Combined version with fixes 1–6 (about 620 words, against 475 now)

> Strickland (1958) had 40 students act as supervisors of two workers, who were in fact fictitious. Each supervisor could fine the workers and was told to check one after nine of ten work periods and the other after two. Both workers produced the same output, and every check found the work in order. Afterwards, asked which worker they trusted more, more supervisors named the rarely checked one than the often checked one. Like Strickland's supervisors, our board member holds power over the person checked, and our slides, like the workers' output, contain no errors. But Strickland's checks found nothing and the workers never spoke; our checker raises wrong flags and the CISO answers. The checker also reports "No issues found" on slides 1 and 5, so we can separate being checked from being flagged. This informs H1.
>
> Prike et al. (2025) introduced 406 participants to two unfamiliar politicians through a profile and a series of social-media posts. In some groups, eight false posts favoured one candidate, and a fact-checking organisation marked each of them false. At the end, participants rated that candidate as less competent and less moral than his opponent, compared with a control group that saw only neutral posts, and fewer voted for him. We borrow the profile, the 0-to-10 rating and the final rating of how true each flagged claim is. Their fact-checks were correct, the candidate stayed silent, and ratings were taken once, at the end; our flags are wrong for this company, the CISO answers, and trust is rated after every step. This informs H1.
>
> Chen et al. (2026) is the closest design (*N* = 135). An LLM chatbot, presented as a human doctor, held a short stress-management session with online participants, then offered his AI assistant's opinion and said only that it agreed or disagreed with him. Compared with agreement, disagreement left participants less sure about their situation and made the doctor seem lazier, and both went with lower trust in him. The authors think the missing explanation made him look lazy and suggest explaining the disagreement, which they did not test. There the doctor brought in the AI himself, and the disagreement and his silence came together, so the loss of trust cannot be split between them. We give the checker to the person in power, rate trust after the flag and again after the answer, and vary the answer: H2 tests the explanation, H3 the non-answer.
>
> Greitemeyer and Sagioglou (2015) had participants rate a researcher accused of misconduct on how trustworthy he was, whether they believed his findings and whether he should be funded, and rate him again after they read that he had been wrongly accused (*N* = 185 and 358). The accused researcher was rated about two points lower than one never accused, on a scale from −5 to +5. After the clearing he rose by about two points, to the level of the never-accused researcher, but not clearly above it. In the second study he did end above a never-accused researcher about whom participants had been told less; the authors trace that gap to the extra information. This is why our neutral group hears the same answers. There an outside finding cleared the researcher; here the CISO answers the flag himself. This informs H2: recovery is expected, a gain would be new.
>
> A checker may also change what the board member does. Directors in Aamir et al. (2026) said independent reviews gave them the confidence to challenge management. In von Schenk et al. (2024), people who consulted a lie-detection algorithm that was right about two times in three, and were told a statement was a lie, mostly accused its author; without the algorithm, about one in five did. This informs H4a. We found no study of whether the same answer moves trust further after a flag than without one (H4b). Reif et al. (2025) found that employees described as getting help from generative AI were rated lazier and less competent than employees getting the same help from another source. So the CISO never uses or mentions AI, and because naming the checker may look like relying on AI, H4a counts asking and challenging together.

References to update: add von Schenk et al. (2024) (full reference in SHORTLIST.md, item 3). Drop Ecker & Rodricks and Mickelberg unless fix 9 keeps Mickelberg. If fix 7 is used, add Ferrin et al. (2007) after checking it.

The short version in Working notes §10 has the same Reif link and the same loose Prike comparison. If that version is submitted, apply fixes 1–3 there too.
