# Does a Vision-LLM Content Checker in a Board Member's Hands Erode or Build Trust in the Presenting CISO?

## Research question

A VLM (vision-language model) lets a listener point a camera at a slide and have it checked while the presenter still talks. Boards are non-experts making high-stakes decisions. Aamir et al. (2026) report that directors lacking cyber literacy struggle to question cyber-risk reports, often accept management's account, and pushed back more confidently where independent reviews were routine. A VLM checker offers such a review on the spot but without the company's context: a flag can be right by the textbook and wrong for this company. The relationship is unequal: the board holds the power and the budget, the CISO the expertise the board depends on.

> **Research Question:** When a board member uses a VLM content checker during a CISO's briefing, how does their trust in the CISO change from the start to the end of the briefing, compared with a board member who follows the same briefing and explanations without a checker?

## Hypotheses

Trust is the participant's 0-to-10 rating of the CISO.

- **H1 (flag).** 
	- Trust decreases after a flagged slide and is lower than the prior slide and the neutral group.
- **H2 (answer).** 
	- Post answer with CISO's reason, trust is higher than prior to flagging and increases beyond initial trust score opposed to the neutral group hearing reason but not getting the flag.
- **H3 (non-answer).** 
	- After a denial with no reason, trust is lower than after the flag.
- **H4 (participation and amplification).** 
	- (a) Checker participants question the CISO at more slides than neutral participants say they have questions; 
	- (b) after an answer, trust moves further in the checker group than in the neutral group, (b1) further up after the reason, (b2) further down after the denial.

## Method

### Study design

Online simulated boardroom, about 10 minutes, in two versions; sign-ups fill the two groups in turn (round robin, not at random). Asymetric whenever they want to take the boardroom simulation. Round robin is for equal participation numbers, they will not know which group they fall into.

1. Scenario: the participant sits on the board and decides on the CISO's budget request. Profile establishing the CISO's experience. First trust rating. Checker group: the checker is introduced.
2. The CISO presents five error-free slides, identical in both groups. Checker group: a panel shows the camera image of each slide, "No issues found" on slides 1 and 5, and on slides 2–4 a marked region with a flag such as "Architecture likely wrong" or "Numbers do not add up".
3. After each of slides 2–4: trust rating. Checker group: the participant says nothing, asks the CISO to explain, or challenges the CISO with the checker's finding; if nothing, the board chair asks a routine question. Neutral group: the participant says whether they have questions about the slide; the chair asks the routine question. The CISO's answer is the same whoever asks: two answers give the company's reason, one is a denial with no reason, on a slide balanced across participants within each group. Second trust rating.
4. Budget decision and background questions.
5. Debrief: no slide had a real issue, the flags were fixed in advance, and the study measures how LLM flags affect non-experts' trust in the presenter.

The checker is simulated: flags were drafted with a real VLM from images of the slides, then fixed so all participants see the same output. Every flag is wrong for this company, and each full answer explains why.

### Measures

Trust in the CISO, 0 to 10, at each step above; the first rating is each participant's own baseline, and each flag and answer is scored as the change from the rating just before it. The action chosen at each flag. The budget decision, to see whether a change in trust reaches an actual decision. Education, age group, field of expertise, self-rated cybersecurity knowledge (to split out prior expertise) and own AI use. Comparing the groups separates the effect of the checker's flags from that of the participant's own interaction.

Review impact: For the study we don't want that people overthink their decision, it will end up in fabricated decisions. I would like to define trust as benevolence, integrity, competence and ask for all three whenever mentioning "trust".

### Validity

The change from the first to the last rating answers the research question; each hypothesis compares the groups at the same step, or a participant with their own previous rating. The groups differ only in the flags and in who asks (checker participants may ask the CISO themselves; in the neutral group the chair always asks). Slide 1, with no issue reported, shows whether the checker's presence alone moves trust; we check that the groups start alike on the first rating. Limit: participants choose whether to ask, so comparisons between actions only describe them.

## Recruitment

I will recruit 40 adults of any education level, background and field online through our personal and professional networks, mailing lists and LinkedIn at no compensation.

## Appendix: Related work

Strickland (1958) had 40 male students play a supervisor who was told to check one of two fictitious workers after nine of ten work periods and the other after two. Both workers produced the same output, and every check found the work in order. Afterwards, asked which worker they trusted more, more supervisors named the rarely checked one than the often checked one. Like Strickland's supervisors, our board member holds power over the person checked. But his checks found nothing and the workers never spoke; our checker claims faults that are not there, and the CISO answers. Because checking alone may lower trust, the checker also shows "No issues found" on slides 1 and 5, so that flagged and unflagged slides can be compared. 

Prike et al. (2025) introduced 406 U.S. adults to two unfamiliar, fictional candidates through a profile and a series of posts. When a nonpartisan fact-checking organisation marked as false eight statements that favoured one candidate, participants rated him as less competent and less moral than his opponent, compared with a group that saw only neutral posts, and fewer voted for him. We borrow the profile, the 0-to-10 rating and the end-of-study rating of how true each checked claim is. Their checker was human and right, the candidate stayed silent, and ratings were taken once, at the end; ours is an AI tool and we check at each step.

Chen et al. (2026) is the closest design (*N* = 135). At the end of a short online therapy chat, a chatbot presented as a human doctor said only that his AI assistant agreed or disagreed with him. Compared with agreement, disagreement left participants less sure about their own situation and made the doctor seem lazier, and both went with lower trust in him. Because the doctor never answered, the study cannot separate the disagreement from the silence. There the doctor brought in the AI himself. We give the checker to the board member and rate trust after the flag and after the answer.

Greitemeyer and Sagioglou (2015) had online participants rate a researcher accused of misconduct, and again after they read that he had been cleared (*N* = 185 and 358). The accused researcher was rated about two points below one never accused, on a scale from −5 to +5; after the clearing he rose by about two points, to that level but not reliably above it. He ended clearly above a never-accused researcher only when that researcher had been described with less information, which is why our neutral group hears the same answers. In Mickelberg et al. (2025), a retraction removed the effect of an allegation on how likeable a person seemed in two of three experiments; the authors note that rating after every step may draw attention to the latest information, and both our groups rate at the same steps. In these studies a third party did the clearing; our CISO answers for himself. This informs H2: recovery is expected; a reliable gain would go beyond these results.

A checker may also change what the board member does. Directors in Aamir et al. (2026) felt more confident challenging management where independent reviews were routine. In von Schenk et al. (2024), 2,040 people each judged whether a stranger's written statement was a lie, knowing that an accusation would cost the writer money. Without help, about one in five accused. Those who chose to consult a lie-detection algorithm that was right about two times in three mostly followed it: when it said "lie", almost 85% accused. The accused could not answer. Chen et al. note that an AI's contrary opinion shifts the burden of proof onto the expert. After a flag, a reason meets that burden and a denial fails it, so the same answer should move trust further than in the neutral group, where nothing has to be proved. Reif et al. (2025) found that employees described as getting help from generative AI were rated lazier and less competent than employees who got the same help from another source, so the CISO never mentions AI; and because naming the checker may look like relying on AI, we log which question participants choose and how often they use AI themselves.


## References

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–22). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791142

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, *213*, Article 103824. https://doi.org/10.1016/j.ijhcs.2026.103824

Greitemeyer, T., & Sagioglou, C. (2015). Does exonerating an accused researcher restore the researcher's credibility? *PLOS ONE*, *10*(5), Article e0126316. https://doi.org/10.1371/journal.pone.0126316

Mickelberg, A. J., Walker, B., Ecker, U. K. H., Howe, P. D. L., Perfors, A., & Fay, N. (2025). Did he or didn't he? Mixed evidence for the continued influence of retracted misinformation on person impressions. *PLOS ONE*, *20*(5), Article e0322045. https://doi.org/10.1371/journal.pone.0322045

Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, *46*(5), 1081–1100. https://doi.org/10.1111/pops.13059

Reif, J. A., Larrick, R. P., & Soll, J. B. (2025). Evidence of a social evaluation penalty for using AI. *Proceedings of the National Academy of Sciences*, *122*(19), Article e2426766122. https://doi.org/10.1073/pnas.2426766122

Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, *26*(2), 200–215. https://doi.org/10.1111/j.1467-6494.1958.tb01580.x

von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, *27*(7), Article 110201. https://doi.org/10.1016/j.isci.2024.110201


## AI transparency

AI was used to help me to articulate, challenge, correct my work.
AI was used to help me research papers which could have a relationship and summarize them prior to me reading them partially or fully.
