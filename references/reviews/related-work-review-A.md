# Reviewer A: Related work in ABSTRACT-v4

(The typo path `…-bd7-…/rw-reviews/` does not exist, so this report is saved at `…-bed7-…/rw-reviews/A.md`.)

**Verdict in brief.** Most claims are accurate, and the five paragraphs are well chosen for H1 and H2. Three things hold the section at 2 out of 3. H4, as rewritten, has no paper that actually serves it: the Reif paragraph is about something else. No cited paper lets the accused person reply, although the reply is what H2 and H3 manipulate, and the trust-repair studies that test such replies are missing. And the Strickland paragraph's link to our design ("the groups differ only in the checking") is wrong. All three can be fixed mostly by swapping material, with a modest increase in length (see the length check at the end).

Page numbers refer to the files in `papers/`: Prike is the accepted manuscript and Ecker & Rodricks the in-press manuscript, so their page numbers are manuscript pages.

---

## A. Accuracy, claim by claim

### A0. Aamir et al. (2026), sentences in the Research question

| # | Claim | Verdict | Where in the paper | Correction |
|---|---|---|---|---|
| 1 | "In interviews with 13 Australian directors" | Accurate | §3.2, Table 2: 13 participants, each with an executive or non-executive director role, held now or within the past five years | Optional: 9 of the 13 were executive directors, so they were also part of management (Table 2). |
| 2 | "gaps in cyber literacy and confidence make it hard for directors to question cyber-risk reports" | Accurate (as the authors' report) | Abstract: the gap "limits their ability to interrogate risk reports, challenge assumptions" | None. "report" is the right verb for 13 self-reports. |
| 3 | "boards without such literacy often accept management's account" | Accurate | §6.2: "Where this literacy was missing, boards often defaulted to accepting management narratives." | None. |
| 4 | "directors rely on independent reviews to push back" | Imprecise | §4.1: "Where external assurance was routine, directors reported greater confidence in challenging management perspectives." §6.2: "not every board sought such input"; smaller boards found outside checks "costly or disruptive" | "directors who had regular outside reviews, such as audits and red-team tests, felt more confident to challenge management, though not every board sought them." |
| 5 | "A llm assisted camera-capture checker offers such a review on the spot. But it lacks the company's context …" | Our own argument, not a claim about the paper | Every review the directors describe was done by people who knew the company (§6.2) | Keep it. Fix the typing: "An LLM-assisted camera checker"; "in the board's hands"; "sits with the CISO". |
| 6 | "the board decides on the budget, and the CISO holds the expertise the board depends on" | Accurate in substance | §2.1: boards "depend on CISOs to translate technical risk"; §1: CISOs "lack decision-making power" (from cited work [34], not the interviews) | None needed. |
| 7 | "Research on AI fact-checking measures belief in the checked claim. We measure trust in the checked person, step by step." | Imprecise (too broad) | Chen et al. rated trust in the checked doctor; Prike et al. rated the checked candidates; per the shortlist, Chae & Tewksbury rated politicians checked by fact-checks labelled as AI | See Fix 7. The real gap is that these studies rate the person once, at the end, and the person never answers. |

### A1. Strickland (1958)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | "had 40 students play a supervisor" | Accurate | p. 202: "40 male students at the University of North Carolina" | Optional: "male"; the supervisor could fine the workers (p. 203). |
| 2 | "who was told to check one of two workers after nine of ten work periods and the other after two" | Accurate | p. 203: A checked after periods 1–4 and 6–10, B after 5 and 10; the schedule was given by the experimenter | Add that the workers were fictitious (p. 203). |
| 3 | "Both workers produced the same output" | Accurate | Table 1, pp. 203–204: equal totals (52/52 in Group I, 50/50 in Group II) | None. |
| 4 | "every check found the work in order" | Accurate | pp. 203–204: no shortfall in any checked period. In Group II, the full record handed over afterwards showed each worker fell short once in an unchecked period | None needed. |
| 5 | "Afterwards, more supervisors trusted the rarely checked worker than the often checked one." | Accurate | Table 3, pp. 206–207: one question (trust A more, B more, or equal); ties split; p = .039 (Group I) and < .001 (Group II) | None. The item compared the two workers. |
| 6 | "Checking alone lowered trust while the work stayed the same." | Imprecise | Trust was asked once, as a comparison of A with B. There was no earlier rating and no unchecked group. Strickland says absolute distrust is untested and "must await an experimental variation in which S has only high or low surveillance" (p. 213) | "The often checked worker was trusted less than the rarely checked one, although their work was the same." |
| 7 | "Our slides contain no errors, so the groups differ only in the checking." | Inaccurate (about our own design) | The Method says the groups differ "in the checker, in the questions it proposes, and in who asks", and the checker raises three flags. Strickland's checks never alleged a fault. | The analogy fits the checked slides without a flag ("No issues found"), not the flagged ones. See Fix 3. |

### A2. Prike et al. (2025)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | "introduced 406 participants to two unfamiliar politicians through a profile and a series of posts" | Accurate | p. 9 (N = 406, politically moderate U.S. Prolific users); p. 11 (profiles with name, party, headshot); pp. 8, 12 (social-media posts) | Optional: "fictional". |
| 2 | "When a fact-checking organisation marked eight of one candidate's statements as false" | Imprecise | p. 8: eight false statements "all associated with one of the candidates", which either boosted him or ran down his opponent. They came from the candidates or from a fictional newspaper (Table 2, p. 13: two of them attack the opponent). The checker also confirmed two true statements per candidate. | "labelled as false eight statements that favoured one candidate" |
| 3 | "participants rated him as less competent and less moral than his opponent, compared with a group that saw no fact-checks" | Imprecise | pp. 17–19: the comparison group was the control, which saw only the 16 general campaign posts and no false statements at all (Table 1, p. 10). Ratings were taken once, at the end (p. 16). The drop was large (competence d = 0.90 and 1.14; morality d = 1.11 and 1.31). | "… compared with a control group that saw only the candidates' general campaign posts. Ratings were taken once, at the end." |
| 4 | "We borrow the profile and the 0-to-10 rating." | Accurate | p. 12: 0–10 sliders for competence, morality and warmth, starting at the midpoint | None. Their scale measured competence and morality, not trust. |
| 5 | "Their fact-checks were correct and the candidate stayed silent" | Accurate | Table 2, p. 13 (each FALSE label gives the correct facts); the candidates never reply | None. |

### A3. Chen et al. (2026)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | "is the closest design (N = 135)" | Accurate (N); "closest" is a fair judgement | §3.6: 135 analysed (67 disagreement, 68 agreement) of 260 recruited | None. |
| 2 | "A simulated doctor told patients that his AI assistant agreed or disagreed with his recommendation, and gave no explanation." | Imprecise | §3.1, §3.3, Appendix steps 6–7: the "doctor" was a GPT-3.5 chatbot presented as a human therapist, in a brief online stress-management chat with U.S. adults (MTurk panel). He asked whether the patient wanted to hear the AI's opinion, then said only that it agreed or disagreed. | "a doctor, played by a chatbot and presented as human, told patients in a short stress-management chat that his AI assistant agreed or disagreed with him, and gave no explanation." |
| 3 | "After disagreement, patients felt less sure of their situation and saw the doctor as lazier" | Imprecise | §4.2.1, §4.2.3: the comparison is disagreement against agreement. There was no group without an AI and no rating before the second opinion. | "Compared with patients told that the AI agreed, patients told that it disagreed felt less sure …" |
| 4 | "and both went with lower trust in the doctor" | Accurate | §4.2.1, §4.2.3, §4.5: both routes linked to lower cognitive, affective and behavioural trust. The paper reports no overall trust difference between the groups. | Optional: "rated once after the chat". |
| 5 | "The authors suspect the missing explanation" | Accurate | §5.2: "This finding may be due to the doctor not providing any explanations after the AI's disagreement, which could be seen as a lack of effort." | None. |
| 6 | "and suggest a short one, which they did not test" | Imprecise | §5.1 and §5.5: they suggest explaining *the AI's disagreement*, for example "Let me take a look at what specifically led it in that direction", or raising it only indirectly. They also advise doctors not to disclose the disagreement openly. That is not the expert defending his own choice. | "and suggest, without testing it, that explaining the disagreement could limit the loss of trust." |
| 7 | "There the expert brought in the AI himself." | Accurate | Appendix step 6; §3.1 | None. |

### A4. Greitemeyer & Sagioglou (2015), with Ecker & Rodricks (2020) and Mickelberg et al. (2025)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | "had participants rate a researcher accused of misconduct, and again after a report cleared him (N = 185 and 358)" | Imprecise ("a report") | pp. 3–4, 7: the ethics committee of his university had demanded that one of his articles be retracted. Participants were later told that "None of the allegations were confirmed" and the article was not retracted. No report is named. N = 185 and 358 are correct (MTurk). The rating averaged three questions: how trustworthy the researcher was, whether his findings were believable, and whether he should get more funding (p. 4). | "… and again after they learned that the allegations were not confirmed and the article would stay" |
| 2 | "The accused researcher was rated about two points lower than one never accused, on a scale from −5 to +5." | Accurate | Study 1: +1.54 vs −0.16 and −0.30 (p. 5). Study 2, against the never-accused group that read the article: +2.35 vs +0.38 and +0.15 (p. 8) | Optional: "about 1.7 to 2 points". |
| 3 | "After the clearing he rose by about two points" | Accurate | Study 1: −0.16 to +1.85 (p. 6); Study 2: +0.38 to +2.47 (p. 9) | None. |
| 4 | "to the level of the never-accused researcher, but not reliably above it" | Accurate (against the matched group) | Study 1: +1.85 vs +1.63, p = .536 (p. 6). Study 2: +2.47 vs +2.27, p = .508 (p. 9). In Study 2 he did end well above the never-accused group that had not read about his article (+2.47 vs +0.96). The authors trace that gap to the extra information about the article, not to the clearing (pp. 9–10). | Keep, and use the Study 2 point to justify our neutral group (Fix 5). |
| 5 | "A retracted allegation likewise stopped affecting trait ratings of a person in Ecker and Rodricks (2020)" | Accurate | Exp. 1 (N = 142 UWA students): likability after retraction equalled control, F < 1 (p. 16); aggressiveness likewise; Exp. 2 (N = 287) replicated. Face ratings showed a tentative lasting effect (Abstract; p. 28) | "did not lower ratings of the person once it was retracted" is closer to the design: one rating, taken after the retraction, compared between groups. |
| 6 | "and in two of three experiments by Mickelberg et al. (2025)" | Accurate | Abstract; Discussion, p. 14: Experiments 1 and 3 found no lasting effect, Experiment 2 did. The measure was the Reysen likability scale (p. 5); about 300 Prolific adults per experiment | "trait ratings" should read "ratings of how likeable the person was". |
| 7 | "There a third party did the clearing." | Accurate | All three papers: the retraction came from the study materials, never from the accused | None. |

### A5. Reif et al. (2025)

| # | Claim | Verdict | Where | Correction |
|---|---|---|---|---|
| 1 | "employees described as getting help from generative AI were rated lazier and less competent than employees getting the same help from another source" | Accurate | Study 2 (N = 1,203): lazier 2.50 vs 2.16; also less competent, diligent, independent and self-assured; also below employees with no help mentioned | None. |
| 2 | "people who imagined using AI at work said they were less likely to tell their manager" | Imprecise (missing comparison) | Study 1 (N = 497): 4.91 vs 5.25, against people who imagined using a non-AI "dashboard creation tool" | "… less likely to tell their manager than people who imagined using a non-AI tool" |
| 3 | "Raters who used AI weekly or daily did not judge AI users as lazier." | Accurate | Study 4 (N = 1,002): the gap in laziness ratings was not significant for weekly or daily users; the comparison worker used Microsoft Office | Optional: "not reliably lazier than a worker using Microsoft Office". |
| 4 | "Challenging the CISO with the checker's finding shows reliance on AI; asking the CISO to explain does not." | Our inference, reasonable | The paper studies people who use AI for their own work, not people who use AI to check someone else | None. Label it as our reasoning. |

**Count:** 37 rows checked. 24 are accurate, 10 imprecise (one of them the over-broad gap claim, A0 #7) and 1 inaccurate (A1 #7). The other 2 (A0 #5, A5 #4) are our own reasoning, not claims about a paper. No number is wrong.

---

## B. Link to the hypotheses

| Paragraph | Stated link | Does it fit? | Can a marker follow why? | Does it say how we extend the paper? |
|---|---|---|---|---|
| Strickland | H1 | **Partly.** Strickland's checks never alleged anything, so the paper speaks to being checked, not to being flagged. That is the comparison between flagged slides and "No issues found" slides within the checker group, which the Validity section already plans. The reason given for the link ("the groups differ only in the checking") contradicts our own Method. | Not quite: the sentence that makes the link is the wrong one. | **No.** Two real extensions go unsaid. First, Strickland left absolute trust untested and asked for a design with only high or only low checking (p. 213); our neutral group and repeated 0–10 ratings do that. Second, his supervisor held power over the workers, as our board member holds the budget, and this is the only cited paper where the checker holds power over the person checked. |
| Prike | H1 | **Yes**, by analogy: correct checks, no reply, competence and morality rather than trust. | Yes. | Yes (wrong flags, the CISO answers). It could add that we rate after each step, where they rated once. |
| Chen | H2, H3 | **Yes for H2** as motivation (an untested suggestion). **H3 fits, but the reason is not given**: the link is the authors' reading that the missing explanation made the doctor look lazy, which the paragraph states but never connects to H3. The paper's strongest role is **H1** (an AI contradicting an expert went with lower trust in him), and the paragraph does not claim it. It also offers **H4b** support that goes unused: the authors note that an AI's disagreement puts the burden of proof on the doctor (§5.1, citing Kempt & Nagel). | Mostly. | Yes: the AI goes to the person in power, and the expert answers. Clear. |
| Greitemeyer + 2 | H2 | **Yes.** It is the only cited paper with a rating after the doubt and again after it is cleared, and "recovery is expected, a gain would be new" is the right reading. | Yes. | Half. "Here the CISO clears the flag in person" names the difference but draws no consequence. The authors credit the full recovery partly to a believable source of the clearing (p. 10), and an accused person clearing himself may be believed less. The Study 2 finding (a gain appeared only against a group that lacked information) is the reason our neutral group hears the same answer. That link is in the Working notes but not in the section. |
| Reif | H4 | **No.** H4 now claims (a) the checker group questions the CISO at more slides and (b) the same answer moves trust further after a flag. Reif bears on neither. It bears on *which* action participants pick (ask or challenge), which H4a counts together, and on the rule that the CISO never mentions AI. | No. A marker reading "This informs H4" will not find the link. | Partly: "We log which one participants choose". |

**Is H4 served?** Not by any paragraph. Support exists but is scattered or missing:
- **H4a** (more questioning with a checker): Aamir §4.1 (routine outside review gave directors "greater confidence in challenging management") sits in the Research question and is not tied to H4. The one paper that tests this directly, von Schenk et al. (2024), was on the shortlist and is not cited.
- **H4b** (the same answer moves trust further after a flag): Chen's burden-of-proof remark, and Strickland's argument that a superior "cannot know first-hand the nature of the loyalty of his subordinates until he perceives that they have had an opportunity to be disloyal" (p. 213). A flag gives the CISO a chance to fail, and an answer that passes is stronger evidence than the same answer to a routine question. Neither is used.
- A note for the authors, outside the section itself. Greitemeyer's data already show the "further up" half of H4b: only the accused group moved, by about 2 points, and it ended at the never-accused level. So "further up" follows from H1 plus full recovery. The new part of H4b is "further down after the denial", and no cited paper supports it. Greitemeyer's accused-and-never-cleared group held steady and did not fall further (Study 1: −0.30 to −0.28; Study 2: +0.15 to +0.27).

**H2's "gain" and H3.** The section is honest that a gain would be new. H3 rests only on Chen's interpretation. The trust-repair work in C1 is the literature that tests denials and non-replies.

---

## C. Coverage

**C1. Missing major work: trust repair after an accusation (outside the files; from my own knowledge, check before citing).** H2 and H3 manipulate the accused person's own reply. None of the cited papers lets the accused reply: Chen's doctor, Greitemeyer's researcher, Prike's candidate and Strickland's workers are all silent. Organisational psychology has tested exactly this:
- Kim, P. H., Ferrin, D. L., Cooper, C. D., & Dirks, K. T. (2004). Removing the shadow of suspicion: The effects of apology versus denial for repairing competence- versus integrity-based trust violations. *Journal of Applied Psychology, 89*(1), 104–118. https://doi.org/10.1037/0021-9010.89.1.104. Participants watched a recorded job interview. The applicant, accused of an accounting error at a previous job, either apologised or denied it. When the error was framed as a lack of skill, apology restored trust better than denial; when it was framed as dishonesty, denial did better.
- Ferrin, D. L., Kim, P. H., Cooper, C. D., & Dirks, K. T. (2007). Silence speaks volumes: The effectiveness of reticence in comparison to apology and denial for responding to integrity- and competence-based trust violations. *Journal of Applied Psychology, 92*(4), 893–908. https://doi.org/10.1037/0021-9010.92.4.893. A reply that neither admits nor denies restored trust no better than the weaker of the two.

Our denial ("That's not a problem, and I'd rather not go into the details here") is half denial, half reticence. Our flags are about competence, the case where denial did worst. Our full answer is a third kind of reply, not tested there: a reason that shows the flag was wrong. This is the most important addition. It gives H3 a direct basis and places H2's reply among known replies.

**C2. H4a: von Schenk et al. (2024), from the shortlist.** A fallible lie-detection algorithm in the judge's hands made accusing much more common. The shortlist gives about 19% of judgements as accusations without the algorithm and about 85% among judges who asked for it and got a "lie" prediction. Only about 31% asked for it, and accusations also became more accurate (51% vs 41%). The outcome is accusing, not trust, and the accused never replies. That is exactly the gap our design fills. It serves the rewritten H4 far better than Reif. The full text is free (PMC11302104). It is not in `papers/`, so check the numbers before citing.

**C3. Chae & Tewksbury (2026), from the shortlist.** With well-known politicians, neither AI-labelled nor human-labelled fact-checks changed feelings toward them or their perceived competence (N = 3,978). This is worth one clause beside Prike: it is the counter-evidence, it speaks to whether an AI source matters, and it supports the choice of an unfamiliar CISO introduced by a profile, since Prike et al. trace their large effects to the absence of prior opinions (p. 28). It is counter-evidence, not a core paper.

**C4. Arkes et al. (2007).** Not needed. Like Reif, it is about the *user* of a decision aid being rated less able, and our CISO never uses AI. It shares Chen's medical setting and measures rated diagnostic ability, not trust. It would only strengthen the stimulus rule, which Reif already covers.

**C5. The passing citations (Ecker & Rodricks; Mickelberg et al.).** They earn a small place. They show that Greitemeyer's full recovery is the usual result for ratings of a person, and Mickelberg's one exception is a fair caveat. But they add a third "an outside party retracts" study from the same research group. Their allegations are moral (hitting a partner, an affair, kicking a dog), not competence, and they take a rating only after the retraction. If words are short, cut them before cutting C1.

**C6. Weakest cited paper.** For the hypotheses, it is Reif: it tests none of H1–H4. Keep it as one or two sentences that justify logging the action and AI use and the rule that the CISO never mentions AI, and move H4 to C2. Strickland is old and uses a relative, single-item judgement. But it is the only paper on being checked without an accusation, and the only one where the checker holds power over the person checked. Keep it, reframed (Fix 3).

**C7. Board and cybersecurity setting.** Aamir remains the only paper, which matches the shortlist's finding and supports the novelty claim.

**C8. Theory of trust (from my own knowledge).** The section never says which part of trust the flags attack. Our flags question competence ("Architecture likely wrong"), which in the Mayer and Davis framework is "ability". One clause would link the hypotheses to the questionnaire already used in the Method. Optional.

---

## D. Score and fixes

**Score: 2 of 3.** The section discusses five relevant papers accurately, and their links to H1 and H2 are clear. But the rewritten H4 has no paper that serves it, the studies that test an accused person's own reply (the core of H2 and H3) are missing, and the Strickland link rests on a sentence that contradicts our own design. Fixes 1 to 4 would, in my view, lift it to 3, at the cost of roughly 150 to 250 more words unless the optional sentences are cut.

### Fix 1: Give H4 its own paper (replace the Reif paragraph)

**What.** Replace the Reif paragraph with a von Schenk paragraph for H4a, and keep Reif as two sentences. Check the numbers against the full text first; they come from the shortlist.

**Wording.**
> In von Schenk et al. (2024), participants judged whether statements written by other people were true, and could ask a lie-detection algorithm that was right about two times in three. Without it, they called about one statement in five a lie; when they asked for it and it said "lie", they did so about five times in six. Only about a third asked. A fallible tool in the judge's hands made accusing easier, and the accused could not answer. Aamir et al. (2026) describe the same pattern in boardrooms: directors with regular outside reviews felt more confident to challenge management. This informs H4a: we expect board members with the checker to question the CISO at more slides than board members without it say they have questions. Reif et al. (2025) found that employees described as getting help from generative AI were rated lazier and less competent than employees who got the same help from another source, and that people who imagined using an AI tool at work said they were less likely to tell their manager than people who imagined a non-AI tool. A challenge that names the checker shows that the board member relies on AI, and a request to explain does not, so we log which one participants choose and how often they use AI themselves.

**Why.** H4 was rewritten after the section, and "This informs H4" no longer holds for Reif (B, C2, C6).

### Fix 2: Add the trust-repair studies for H2 and H3

**What.** Add two sentences, best placed after the Chen paragraph or at the end of the Greitemeyer paragraph. If length is tight, cut the Ecker & Rodricks / Mickelberg sentence to make room. Check the details against the papers first; they are not in `papers/`.

**Wording.**
> Studies of trust repair test the accused person's own reply. In Kim et al. (2004), participants watched a recorded job interview in which an applicant accused of an accounting error at a previous job either apologised or denied it; when the error was framed as a lack of skill, the apology restored trust better than the denial. Ferrin et al. (2007) found that a reply that neither admits nor denies restored trust no better than the weaker of the two. Our flags question the CISO's skill, and our denial is close to such a non-reply, which informs H3. Our full answer is a third kind of reply that these studies did not test: a reason that shows the flag was wrong (H2).

Add both to the References.

**Why.** The reply is the core manipulation of H2 and H3, and every cited paper has a silent accused (C1). This is the "major related work" a marker in trust research would look for first.

### Fix 3: Reframe Strickland

**What.** Remove "so the groups differ only in the checking". Link the paper to the unflagged checked slides, state the relative-trust limit and our extension, and use the power parallel and the "chance to fail" argument.

**Wording.**
> Strickland (1958) had 40 students play a supervisor who could fine two workers, who were in fact fictitious. Each supervisor was told to check one worker after nine of ten work periods and the other after two. Both produced the same output, and every check found the work in order. Afterwards, more supervisors said they trusted the rarely checked worker more than the often checked one than said the reverse. The supervisors compared the two workers once, and Strickland left open whether checking lowers trust in absolute terms. He argued that a superior can only learn whether to trust someone after seeing him act well when he had the chance to fail. As in our study, the one who checks holds power over the one checked. Our checker also checks slides without flagging them ("No issues found"), and our neutral group has no checker, so we can separate being checked from being flagged, on a rating of one person over time. This informs H1. His argument about passing a test also informs H2 and H4b.

**Why.** A1 #6 and #7. It also supplies the missing "how we extend" (B) and gives H4b a basis.

### Fix 4: Tighten Chen and state all its links

**What.** Name the comparison (agreement) and what the "doctor" was, correct what the authors suggested, and state the H1, H3 and H4b links.

**Wording.**
> Chen et al. (2026) is the closest design (*N* = 135). In a short online stress-management chat, a doctor, played by a chatbot and presented as human, told patients that his own AI assistant agreed or disagreed with his recommendation, and gave no explanation. Compared with patients told that it agreed, patients told that it disagreed felt less sure about their situation and saw the doctor as lazier, and both went with lower trust in the doctor, rated once after the chat (H1). The authors think the doctor looked lazy because he did not explain, which points to H3. They suggest, without testing it, that explaining the disagreement could limit the loss, and note that an AI's disagreement puts the burden of proof on the doctor, which makes his answer matter more (H4b). There the expert brought in the AI himself and never answered it. We give the checker to the person who decides the budget and let the expert answer with his own reason (H2).

**Why.** A3 #2, #3 and #6. The comparison with agreement matters because Chen had no group without an AI. It also serves H4b (B).

### Fix 5: Draw the consequences in the Greitemeyer paragraph

**What.** Replace "a report cleared him". Add the source-of-clearing point, the Study 2 reason for our neutral group, and the "never cleared" baseline for H3. Reword the passing citations, or cut them for Fix 2.

**Wording.**
> Greitemeyer and Sagioglou (2015) had online participants (*N* = 185 and 358) rate a researcher whose university's ethics committee had demanded that one of his articles be retracted, and rate him again after some learned that the allegations were not confirmed. The rating combined how trustworthy he was, how believable his findings were and whether he should be funded, on a scale from −5 to +5. The accused researcher was rated about two points lower than one never accused. After the clearing he rose by about two points, to the level of the never-accused researcher but not reliably above it; when he was never cleared, his rating stayed low but did not fall further. In Study 2 the cleared researcher did end above a never-accused one, but only against a group that had not read about his article, and the authors trace that gap to the extra information. They credit the full recovery partly to a believable source of the clearing. Here the accused clears himself, and our neutral group hears the same reason, so a gain cannot come from extra information. This informs H2: recovery is expected, a gain would be new. *(Optional:)* A retracted allegation also did not lower ratings of how likeable a person was in Ecker and Rodricks (2020), and in two of three experiments by Mickelberg et al. (2025).

**Why.** A4 #1, #5 and #6. The "how we extend" becomes a real one, and the design choice is grounded in the literature rather than only in the Working notes.

### Fix 6: Correct the Prike comparison

**What.** Name the control group, what was labelled false, and the single rating at the end. Optionally add Chae & Tewksbury as the counter-case (check it first).

**Wording.**
> Prike et al. (2025) introduced 406 U.S. adults to two unfamiliar, fictional politicians through a profile and a series of social-media posts, and asked them at the end to rate each on competence and morality from 0 to 10. When a fact-checking organisation labelled as false eight statements that favoured one candidate, participants rated him as much less competent and less moral than his opponent, compared with a control group that saw only the candidates' general campaign posts. *(Optional:)* With well-known politicians, fact-checks labelled as coming from AI or from people changed neither feelings toward them nor their rated competence (Chae & Tewksbury, 2026); Prike et al. trace their larger effect to the absence of prior opinions. We borrow the unfamiliar person, the profile and the 0-to-10 rating. Their fact-checks were correct, the candidate stayed silent, and they rated once; our flags are wrong for this company, the CISO answers, and we rate after every step. This informs H1.

**Why.** A2 #2 and #3. The optional clause answers the obvious objection (C3) and justifies the profile.

### Fix 7: Correct the Aamir sentence and the gap statement in the Research question

**What.** Qualify "rely on independent reviews" and narrow the gap claim. Fix the typing.

**Wording.**
> … and that directors who had regular outside reviews, such as audits and red-team tests, felt more confident to challenge management, though not every board sought them. An LLM-assisted camera checker offers such a review on the spot. … The power is in the board's hands while the expertise sits with the CISO.

> Most research on fact-checking measures belief in the checked claim. The few studies that rate the checked person rate once, at the end, and the person never answers the check (Prike et al., 2025; Chen et al., 2026). We measure trust in the checked person after each step, and the person answers.

**Why.** A0 #4 and #7. The new gap sentence states what the Working notes (§7) identify as the real gap, "nobody has let a professional answer an AI flag in front of the person holding it", and a marker looks for exactly that under "how your work extends them".

### Fix 8: Carry the changes into the short version

**What.** Working notes §10 repeats the old Strickland, Chen and Reif wording ("supervisors told to check one worker more often trusted that worker less"; "patients then saw him as lazier and trusted him less"; Reif for naming the checker). Update it to match Fixes 1 to 6 if the short version is submitted.

**Why.** Otherwise the submitted text keeps the errors fixed here.

---

**Length check.** Fixes 1, 3, 4 and 5 lengthen their paragraphs by about 150 words together, and Fix 2 adds about 90. Cutting the Ecker & Rodricks / Mickelberg sentence and the optional Chae & Tewksbury clause keeps the section near seven cited works, with five discussed in full (Strickland, Prike, Chen, Greitemeyer, von Schenk) plus Kim/Ferrin and Reif in two sentences each. That stays close to the brief's "4–5 papers" while covering all four hypotheses.
