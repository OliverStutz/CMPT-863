# Related work

*Draft v9 (2026-10-20). Full review: section 1 from the v8 review round; the opening and sections 2 to 5 rebuilt from the verified v6 text in the same four-move shape (what exists and what we know, then what is unresolved and how this study differs). Every number and verified wording is carried from v6.*

A board member points a camera-based checker at a slide; the checker flags a fault that is not there; the chief information security officer (CISO) answers. Five bodies of work bear on this moment: board oversight by non-experts, the checkers now on the market and their accuracy, what wrong flags do to judgment and to the person flagged, how trust in an expert is repaired, and the designs closest to ours.

## Cybersecurity oversight by non-expert boards

Boards must oversee cybersecurity, and most do so without the expertise to judge it (Lowry et al., 2026). The U.S. Securities and Exchange Commission (SEC, 2023) requires public companies to describe how the board oversees cyber risk, but it dropped the plan to disclose whether any director has cyber expertise, reasoning that directors with broad risk skills often oversee cyber effectively, as they do other technical matters (p. 85). Two theories frame what follows (Lowry et al., 2026). Agency theory expects the board to monitor executives whose interests may differ from the firm's, for instance by filtering what the board learns; institutional theory expects boards to adopt the practices that make oversight look legitimate. Tools that could give directors an independent view are already sold: smart glasses that answer questions about what the wearer sees (Meta, 2024) and a meeting assistant that, on request, looks for holes in a participant's argument (Microsoft, n.d.). Neither knows the company behind the slide.

Field studies agree that directors without expertise ask questions but cannot judge the answers. Lowry et al. (2026) interviewed directors, security executives and consultants and found that expertise decides which theory fits: boards with expertise monitored, while boards without it performed practices that looked legitimate, oversight the authors call mostly symbolic. Non-expert directors believed their business experience let them ask the right questions; expert participants saw them unable to judge the answers and left to assess the credibility of the executive in front of them. Thirty-three surveyed CISOs in the same study estimated, at the median, that 40% of CISOs filter their reports to make themselves or their superiors look better; most non-expert directors thought filtering unlikely. In the United Kingdom, directors held back questions rather than expose their lack of knowledge, the budget was by most accounts the only cyber decision boards made, and that reticence left CISOs to decide what reached the board (Opdenbusch et al., 2025). Lowry et al. read the same deference as a gap in expertise rather than power; both studies agree that the CISO controls what the board hears. An independent view changes the picture: Australian directors challenged management more confidently where external assurance was routine, and without cyber literacy boards often accepted management's account (Aamir et al., 2026).

Three things are missing. First, all of this is interview and survey evidence; we found no study that put a checking tool in a director's hand and measured what it does to trust between two human beings. Second, the two research fields that could answer the question do not meet. Trust research is interpersonal psychology: a human accuses, a human replies, and no tool is involved (Ferrin et al., 2007; P. H. Kim et al., 2004). HCI research on checkers mostly studies the tool and measures belief in claims or trust in the AI (Danry et al., 2020; Lu et al., 2025); the few studies that measured trust in the person checked had the expert consult the AI or let it run on its own (Chen et al., 2026; D. Kim et al., 2026), never a listener holding it. Third, the consequences for the board-CISO relationship are unknown. A board that cannot judge the substance has little to weigh a flag against, so a flag may carry weight whether or not it is right, and power itself works against trust: in exchange experiments, the party with better alternatives trusted its partner less (Schilke et al., 2015). A wrong flag may cost a correct CISO the board's trust, and with it the budget and the protection it would buy; a well-answered flag may strengthen trust, and with it approval for a project the company did not need. Whether a checker improves the board's questions is secondary here, although the board studies treat asking the right questions as the mark of competent oversight; what a checker does to trust is the open question.

This study adds the experiment to the board research and the tool to the trust research. A participant acts as a board member who decides the CISO's budget and holds a checker during the briefing; every flag is right by the textbook but wrong for this company, and the CISO answers each one. Trust in the CISO is rated after every flag and every answer and compared with board members who hear the same briefing without a checker. The result is a measure neither field has taken: trust between two people when one of them is augmented with a checker.

## Listener-held checking tools and their accuracy

Checking tools held by the listener exist as products (above) and as research systems, and their accuracy is limited. Audio glasses that judged whether a heard claim was backed by evidence were right on 63.5% of statements in live use (Danry et al., 2020). At the 2020 State of the Union address, the live checker Squash showed 14 relevant and 20 irrelevant fact-checks (Adair, 2020). GPT-4V scored at about the 16th percentile of humans on a visualization literacy test and knew some design rules, yet answered only 23% of questions asking for a single value when it lacked the data (Bendeck & Stasko, 2025). A flag can therefore cite a sound rule and still misread the slide.

What none of this work resolves is what a check does to the audience's trust in the speaker: no study measured it, and Danry et al.'s participants rated recorded statements, not a person. Our checker treats the false flag as the normal case rather than the exception. Its flags were drafted with a real vision-language model (VLM) from images of the slides, fixed so that every participant sees the same output, and all of them are wrong for this company.

## Effects of wrong flags on judgment and on the person flagged

People follow AI warnings even when the warnings are wrong, and those with power act on them. A wrong AI warning on a true headline cut trust in it by 0.30 on a scale from −1 to +1 (Horne & Nevo, 2025). An AI credibility indicator moved accuracy more than an expert's vote, and people could not tell correct AI predictions from incorrect ones (Lu et al., 2025). A fictitious report of a high AI-detection rate made teachers grade an identical piece of student writing 1.8 points lower on a 10-point scale, and red highlighting lowered the grade further (Du et al., 2026). People who consulted a lie-detection algorithm accused in almost 85% of the cases in which it called the statement a lie (von Schenk et al., 2024). The person checked pays. Even checks that find nothing can cost trust (Strickland, 1958); a human fact-check lowered an unfamiliar candidate's ratings of competence and morality (Prike et al., 2025); a radiologist who recommended less testing than an AI was rated lowest of four conditions, while one who recommended more paid no penalty (Madanay et al., 2025); an AI assistant's disagreement made a doctor seem lazier and lowered trust in him (Chen et al., 2026); and employees described as getting AI help were rated lazier and less competent than those who got the same help elsewhere (Reif et al., 2025). Two results bound this. In D. Kim et al. (2026) the advisor's error lowered trust and the AI's involvement changed nothing, and people resisted confident AI accusations of lying when the model's accuracy and confidence were on screen (Loconte et al., 2026). Results on well-known politicians, whose reputations are settled, are not carried over.

What remains unresolved is what a wrong flag does to trust in an expert who is right, and what the expert's reply does afterwards: in every study above, the person judged stays silent and is rated once. Our study differs on each point. The flags are wrong and the CISO is right; the CISO answers; the listener holds both the tool and the budget; and trust is rated at every step. Because accusations of dishonesty and visible confidence figures change how people respond, every flag concerns the slide and never the CISO's honesty, the panel shows no accuracy or confidence figure, and the CISO never mentions AI.

## Trust in an expert and its repair

Trust is the willingness to be vulnerable to another on the basis of perceived ability, benevolence and integrity, and it leads to risk taking, here the budget decision (Mayer et al., 1995). Our three trust items follow the Muenster Epistemic Trustworthiness Inventory (METI), which applies these factors to laypeople judging an expert (Hendriks et al., 2015). With human accusers, the reply decides the repair. After a hearsay accusation of incompetence, an apology repaired trust better than a denial with an account (P. H. Kim et al., 2004), and a reply that neither confirmed nor denied did no better than the weaker of the two (Ferrin et al., 2007). Exoneration by a third party restored an accused researcher to the level of one never accused, but not reliably above it (Greitemeyer & Sagioglou, 2015). Sharma et al. (2023) call a rise beyond the starting level further development rather than repair, and note that most repair studies find trust short of the start, and denial weaker after competence than after integrity violations. With a machine accuser the picture changes. Kempt and Nagel (2022) argue, in a point Chen et al. (2026) take up, that an AI's contradiction shifts the burden of proof to the expert, and that reasons cannot settle the dispute because the machine does not respond to them. The nearest evidence on an expert's reason is indirect: a physician who skipped an available diagnostic aid was rated higher for believing they knew what it would say than for distrusting it (Pezzo et al., 2022).

Three questions are open: whether trust repairs after a machine accusation at all, whether a bare denial lowers trust further or only repairs less, and whether a reason addressed to a listener who can weigh it meets the burden of proof. Our study puts all three to the test. The CISO answers each flag with a reason or a bare denial, delivered by the same person in the same tone; trust is rated after the flag and after the answer; a neutral group hears the same answers without flags; and a rise above the first rating is treated as exploratory, since even a denial with an account was the weaker reply.

## Research gap and the present study

The closest design is D. Kim et al. (2026): an AI checked a human advisor in front of the student who depended on him, and trust in the advisor was measured with METI. Their AI was always right, the advisor never replied, and the student held neither the tool nor any power over the advisor. Chen et al. (2026) is next: the doctor brought in the AI himself and never answered. Gupta et al. (2026), the nearest real-time wearable checker, measured belief in claims and trust in the AI, never trust in the speaker. And only 8 of 94 empirical repair studies used any within-subject design, although a claim of repair needs ratings before the violation, after it and after the reply (Sharma et al., 2023).

No study combines a wrong flag on a correct expert, a checker in the hands of the non-expert who holds power over that expert, the expert's own reply, and trust rated at each step, in a boardroom. This study does, and asks whether a wrong flag costs a correct expert trust and whether a reason recovers it.

## References

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–22). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791142

Adair, B. (2020, February 23). *Squash report card: Improvements during State of the Union … and how humans will make our AI smarter*. Duke Reporters' Lab. https://reporterslab.org/2020/02/23/squash-report-card-improvements-during-state-of-the-union-and-how-humans-will-make-our-ai-smarter/

Bendeck, A., & Stasko, J. (2025). An empirical evaluation of the GPT-4 multimodal language model on visualization literacy tasks. *IEEE Transactions on Visualization and Computer Graphics*, *31*(1), 1105–1115. https://doi.org/10.1109/TVCG.2024.3456155

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, *213*, Article 103824. https://doi.org/10.1016/j.ijhcs.2026.103824

Danry, V., Pataranutaporn, P., Mao, Y., & Maes, P. (2020). Wearable Reasoner: Towards enhanced human rationality through a wearable device with an explainable AI assistant. In *Proceedings of the Augmented Humans International Conference* (pp. 1–12). Association for Computing Machinery. https://doi.org/10.1145/3384657.3384799

Du, P., Liu, T., & Xian, X. (2026). Automation bias in teachers' evaluation of student writing: Effects of algorithmic warnings and visual risk cues in AI detection reports. *Frontiers in Psychology*, *17*, Article 1889402. https://doi.org/10.3389/fpsyg.2026.1889402

Ferrin, D. L., Kim, P. H., Cooper, C. D., & Dirks, K. T. (2007). Silence speaks volumes: The effectiveness of reticence in comparison to apology and denial for responding to integrity- and competence-based trust violations. *Journal of Applied Psychology*, *92*(4), 893–908. https://doi.org/10.1037/0021-9010.92.4.893

Greitemeyer, T., & Sagioglou, C. (2015). Does exonerating an accused researcher restore the researcher's credibility? *PLOS ONE*, *10*(5), Article e0126316. https://doi.org/10.1371/journal.pone.0126316

Gupta, C., Aritonang, N. V., Rajendran, D. P. D., Danry, V., Maes, P., & Nanayakkara, S. (2026). Feeling the facts: Real-time wearable fact-checkers can use nudges to reduce user belief in false information. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–20). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791679

Hendriks, F., Kienhues, D., & Bromme, R. (2015). Measuring laypeople's trust in experts in a digital age: The Muenster Epistemic Trustworthiness Inventory (METI). *PLOS ONE*, *10*(10), Article e0139309. https://doi.org/10.1371/journal.pone.0139309

Horne, B. D., & Nevo, D. (2025). People adhere to content warning labels even when they are wrong due to ecologically rational adaptations. *Scientific Reports*, *15*(1), Article 13896. https://doi.org/10.1038/s41598-025-98221-7

Kempt, H., & Nagel, S. K. (2022). Responsibility, second opinions and peer-disagreement: Ethical and epistemological challenges of using AI in clinical diagnostic contexts. *Journal of Medical Ethics*, *48*(4), 222–229. https://doi.org/10.1136/medethics-2021-107440

Kim, D., Daneshi, R., Draper, B., & Sreedharan, S. (2026). Asymmetric trust effects of corrective AI in expert advisory workflows under epistemic dependence. *Proceedings of the AAAI/ACM Conference on AI, Ethics, and Society*, *9*(3), 1514–1525. https://doi.org/10.1609/aies.v9i3.43244

Kim, P. H., Ferrin, D. L., Cooper, C. D., & Dirks, K. T. (2004). Removing the shadow of suspicion: The effects of apology versus denial for repairing competence- versus integrity-based trust violations. *Journal of Applied Psychology*, *89*(1), 104–118. https://doi.org/10.1037/0021-9010.89.1.104

Loconte, R., Monaro, M., Pietrini, P., Verschuere, B., & Kleinberg, B. (2026). Humans incorrectly reject confident accusatory AI judgments. *Computers in Human Behavior*, *182*, Article 109019. https://doi.org/10.1016/j.chb.2026.109019

Lowry, M. R., Vance, A., & Vance, M. D. (2026). Inexpert supervision: Field evidence on boards' oversight of cybersecurity. *Management Science*, *72*(2), 783–804. https://doi.org/10.1287/mnsc.2023.04147

Lu, Z., Li, P., Wang, W., & Yin, M. (2025). Understanding the effects of AI-based credibility indicators when people are influenced by both peers and experts. In *Proceedings of the 2025 CHI Conference on Human Factors in Computing Systems* (pp. 1–19). Association for Computing Machinery. https://doi.org/10.1145/3706598.3713871

Madanay, F., O'Donohue, L. S., & Zikmund-Fisher, B. J. (2025). Patient reactions to artificial intelligence–clinician discrepancies: Web-based randomized experiment. *Journal of Medical Internet Research*, *27*, Article e68823. https://doi.org/10.2196/68823

Mayer, R. C., Davis, J. H., & Schoorman, F. D. (1995). An integrative model of organizational trust. *Academy of Management Review*, *20*(3), 709–734. https://doi.org/10.2307/258792

Meta. (2024, April 23). *New Ray-Ban | Meta smart glasses styles and Meta AI updates*. https://about.fb.com/news/2024/04/new-ray-ban-meta-smart-glasses-styles-and-meta-ai-updates/

Microsoft. (n.d.). *Catch up on meetings with Microsoft Copilot in Teams*. Retrieved October 20, 2026, from https://support.microsoft.com/en-us/teams/copilot/catch-up-on-meetings-with-microsoft-365-copilot-in-teams

Opdenbusch, J. C., Hielscher, J., & Sasse, M. A. (2025). "Where are we on cyber?" A qualitative study on boards' cybersecurity risk decision making. In *Proceedings of the 2025 Network and Distributed System Security Symposium (NDSS 2025)*. Internet Society. https://doi.org/10.14722/ndss.2025.240595

Pezzo, M. V., Nash, B. E. D., Vieux, P., & Foster-Grammer, H. W. (2022). Effect of having, but not consulting, a computerized diagnostic aid. *Medical Decision Making*, *42*(1), 94–104. https://doi.org/10.1177/0272989X211011160

Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, *46*(5), 1081–1100. https://doi.org/10.1111/pops.13059

Reif, J. A., Larrick, R. P., & Soll, J. B. (2025). Evidence of a social evaluation penalty for using AI. *Proceedings of the National Academy of Sciences*, *122*(19), Article e2426766122. https://doi.org/10.1073/pnas.2426766122

Schilke, O., Reimann, M., & Cook, K. S. (2015). Power decreases trust in social exchange. *Proceedings of the National Academy of Sciences*, *112*(42), 12950–12955. https://doi.org/10.1073/pnas.1517057112

Sharma, K., Schoorman, F. D., & Ballinger, G. A. (2023). How can it be made right again? A review of trust repair research. *Journal of Management*, *49*(1), 363–399. https://doi.org/10.1177/01492063221089897

Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, *26*(2), 200–215. https://doi.org/10.1111/j.1467-6494.1958.tb01580.x

U.S. Securities and Exchange Commission. (2023). *Cybersecurity risk management, strategy, governance, and incident disclosure* (Release Nos. 33-11216; 34-97989). https://www.sec.gov/files/rules/final/2023/33-11216.pdf

von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, *27*(7), Article 110201. https://doi.org/10.1016/j.isci.2024.110201

## Working notes

### Section 1

- **From v8 to v9:** the fixes in `search/review-v8-tally.md` (15 items, both reviewers agreed on 12). Factual corrections: Lowry's conclusion is that expertise is the boundary condition between agency theory and institutional theory (pp. 785, 799), not that one theory wins; agency theory is about diverging interests, not power; Opdenbusch's directors "would be too shy to ask questions" (4 of 18 participants, p. 6), and their fear is what gives CISOs the power to decide what reaches the board (pp. 2, 11); SEC "often effectively oversee" and "public companies"; Lowry interviewed security executives with varied titles (p. 788).
- **Oliver's edits to v8 saved at 20:25** (paragraph 3) carried into v9: "what it does to trust between two human beings"; the consequences sentence now names the stakes on both sides (the protection a budget would buy; approval for a project the company did not need); the questioning sentence restored in his sense (the board studies treat asking the right questions as the mark of competent oversight, our question is trust). His "Several things are missing" kept as "Three", because three follow; change on request.
- **Oliver's three earlier edits to v8** are carried in substance: "no matter the actual context" became our own sentence "Neither knows the company behind the slide" (the Microsoft page says Copilot reasons over the meeting, so the clause could not hang on that citation); the survey sentence now says who estimated and better to whom (33 CISOs, p. 796); the budget clause uses the sourced wording (p. 5).
- **Structure:** paragraph 1 what exists (regulation, two theories, tools); paragraph 2 what we know, claim-led; paragraph 3 what is unresolved (no tool in any study; the two fields do not meet; consequences unknown, with Schilke tied to the flag); paragraph 4 how we extend.
- **The study's own statements, hedged:** "we found no study"; "a flag may carry weight"; the wrong-flag and well-answered-flag sentence (H1 and H2 in plain words); "Neither knows the company behind the slide" (the abstract's premise).
- **Cuts:** "The augmented human ... falls between the two" (the idea stays in the last sentence); "Whether the checker improves the board's questions is secondary ..."; the product sentence shortened and the Copilot prompt left to section 2; paragraph 4 merged.
- **Open for Oliver:** agent R's three optional additions (Proudfoot 2023; NACD & ISA 2026; Nolen 2026, entries in `search/reference-scan-v7.md`); whether section 5 or each section carries the design paragraph (both reviewers note paragraph 4 previews the gap section).
- **Word count:** 746 words for the section (v8 as edited: about 730; v6: 258).

### Opening and sections 2 to 5 (rebuilt 2026-10-20)

- **Source of every claim:** `Literature Review - Draft v6.md` (verified claim by claim in `search/check-tally.md` and `search/academic-tally.md`). No new source, number or comparison; Gupta only in the gap; Dobbs absent; politician sentence kept; trust in the AI not an outcome; no hypothesis labels.
- **Shape:** each section has one paragraph of what exists and what we know, and one of what is unresolved and how the study differs, so that a marker finds the four answers in the same place in every section.
- **Cuts from v6 to fit the budget:** the Copilot prompt and the Meta sentence (now in section 1); the APA page locators on quotations (no quotations remain); the Kempt and Nagel remedy (a second human) and our hypothesis sentence (now the third open question); the Pezzo "in neither case below one with no aid mentioned" clause.
- **Reference list:** the 30 entries of v6 with the three audit fixes (Madanay, F.; Opdenbusch title; Microsoft entry).
