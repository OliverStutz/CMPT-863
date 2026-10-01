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
