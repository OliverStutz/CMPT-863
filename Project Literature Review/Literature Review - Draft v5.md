# Related work

A board member points a camera assistant at the slide of the chief information security officer (CISO), the tool flags a fault, and the CISO answers. Four strands of work bear on this moment: board oversight, the checkers on offer, the cost of wrong flags, and trust repair.

## Boards oversee what they cannot judge

Regulation asks non-experts to oversee an expert. The U.S. Securities and Exchange Commission (2023) chose not to require disclosure of board cyber expertise, holding that broadly skilled directors "often effectively oversee management's efforts without specific subject matter expertise". In Lowry et al. (2026), non-expert directors' oversight was "mostly symbolic" even when they did what expert directors did, and one expert director said for most boards "all they can do is assess the credibility of the executives that are put in front of them". Four interviewees said directors are "too shy to ask questions", and most named the budget as the board's only cyber decision (Opdenbusch et al., 2025). Directors felt more confident challenging management where independent reviews were routine (Aamir et al., 2026).

Opdenbusch et al. (2025) conclude that "the CISOs are the ones in power", because they decide what reaches the board, while Lowry et al. (2026) trace the deference to low expertise; we read this as formal power with the board and informational power with the CISO. In exchange experiments, the party with better alternatives trusted its partner less (Schilke et al., 2015); that power came from having alternatives, not from authority. The board evidence comes from interviews, with no tool or trust measure; we add both.

## Checkers are available, and often wrong

Listener-held checkers exist as products. Ray-Ban Meta glasses answer questions about what the wearer is looking at (Meta, 2024), and Copilot in Microsoft Teams suggests, in a private pane, the prompt "Where are some holes in [a meeting participant's] argument?" (Microsoft, n.d.). Audio glasses that judge whether a heard claim is backed by evidence were right on 63.5 percent of statements live (Danry et al., 2020). At the 2020 State of the Union the live checker Squash showed 14 relevant fact-checks and 20 irrelevant ones (Adair, 2020). GPT-4V scored at about the 16th percentile of humans on a visualization literacy test and showed some knowledge of design best practice, yet answered only 23 percent of questions asking for a single value when it lacked the data (Bendeck & Stasko, 2025). A flag can thus cite a sound rule and still misread the slide. None of this work reports what a check does to trust in the speaker; Danry et al.'s participants rated recorded statements, not a person. Our flags are drafted with a real vision-language model and are all wrong for the company.

## People follow a wrong flag, and the person flagged pays

A wrong AI warning on a true headline cut trust in it by 0.30 on a scale from −1 to +1 (Horne & Nevo, 2025). An AI credibility indicator had a larger effect on accuracy than an expert's vote, and people "lack the capability to differentiate correct and incorrect AI predictions" (Lu et al., 2025). Flags also move people with power to act. A fictitious report of a high AI-detection rate made teachers grade an identical paper 1.8 points lower on a 10-point scale, and red highlighting lowered the grade further (Du et al., 2026). People who chose to consult a lie-detection algorithm accused in almost 85 percent of cases when it said "lie" (von Schenk et al., 2024).

Even checks that find nothing can cost trust (Strickland, 1958), and a human fact-check lowered an unfamiliar candidate's ratings of competence and morality (Prike et al., 2025). People agreed less with a radiologist who recommended less testing than an AI and rated that radiologist lowest of the four conditions; one who recommended more paid no penalty (Madanay et al., 2025). An AI assistant's disagreement made a doctor seem lazier and lowered trust in him (Chen et al., 2026). Employees described as getting help from AI were rated lazier and less competent than those who got the same help elsewhere (Reif et al., 2025), so our CISO never mentions AI.

Two results set limits. In D. Kim et al. (2026) the advisor's error lowered trust, but how the AI was involved changed nothing, so their data tie the loss to the error, not to the AI's contradiction. People resisted confident AI accusations of lying when the model's accuracy and confidence were on screen (Loconte et al., 2026); our flags concern the slide, not honesty, and show neither figure. Results on well-known politicians, whose reputations are settled, are not carried over. Wherever these studies judge a person, that person stays silent and is rated once; our CISO answers.

## Trust in an expert, and how it is repaired

Trust is the willingness to be vulnerable to another, based on perceived ability, benevolence and integrity; it leads to risk taking, here the budget decision (Mayer et al., 1995). Our trust items come from the Muenster Epistemic Trustworthiness Inventory (METI), which applies these factors to laypeople judging an expert (Hendriks et al., 2015).

After a hearsay accusation of incompetence, an apology repaired trust better than a denial with an account (P. H. Kim et al., 2004). A reply that neither confirmed nor denied did no better than the weaker of the two (Ferrin et al., 2007). Exoneration by a third party restored an accused researcher to the level of one never accused, but not reliably above it (Greitemeyer & Sagioglou, 2015). Sharma et al. (2023) call a rise beyond the starting level "further development of trust", and most repair studies, they note, find that trust falls short of the start and that denial works less well after competence than after integrity violations. Hence we treat a rise above the first rating as exploratory, and since even a denial with an account lost, we expect a bare denial to repair less.

These accusers are human. Kempt and Nagel (2022) argue, in a point Chen et al. (2026) take up, that when an AI contradicts the expert in charge the burden of proof moves to the expert. Reasons cannot settle the disagreement, because the machine does not respond to them, so they propose that a second human settle it. The nearest evidence on an expert's reason is indirect: a physician who skipped an available diagnostic aid before a bad outcome was rated higher for believing they knew what it would say than for distrusting it, and in neither case below one with no aid mentioned (Pezzo et al., 2022). Our hypothesis is that the CISO's reason, addressed to a board member who can weigh it, may meet the burden in the listener's eyes, while a bare denial does not.

## The gap and this study

No study combines these pieces. D. Kim et al. (2026) is the closest design: an AI checked a human advisor's work in front of the student who depended on him, and trust in the advisor was measured with METI. Their AI was always right, the advisor never replied, and the student held neither the tool nor any power over the advisor. Gupta et al. (2026), the nearest real-time wearable checker, measured belief in claims and trust in the AI, never in the speaker. Only 8 of 94 empirical repair studies used any within-subject design, though a claim of repair needs ratings before the violation, after it and after the reply (Sharma et al., 2023). We combine them: the checker is held by the non-expert with power over the expert, its flags are plausible but wrong for the company, the expert answers each flag with a reason or a bare denial, and trust is rated after each flag and each answer, against a neutral group that hears the same answers without flags. This shows whether a wrong flag costs a correct expert trust and whether a reason wins it back.

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

Madanay, F. L., O'Donohue, L. S., & Zikmund-Fisher, B. J. (2025). Patient reactions to artificial intelligence–clinician discrepancies: Web-based randomized experiment. *Journal of Medical Internet Research*, *27*, Article e68823. https://doi.org/10.2196/68823

Mayer, R. C., Davis, J. H., & Schoorman, F. D. (1995). An integrative model of organizational trust. *Academy of Management Review*, *20*(3), 709–734. https://doi.org/10.2307/258792

Meta. (2024, April 23). *New Ray-Ban | Meta smart glasses styles and Meta AI updates*. https://about.fb.com/news/2024/04/new-ray-ban-meta-smart-glasses-styles-and-meta-ai-updates/

Microsoft. (n.d.). *Use Copilot in Microsoft Teams meetings* [Support article]. Retrieved October 18, 2026, from https://support.microsoft.com/en-us/office/use-copilot-in-microsoft-teams-meetings-0bf9dd3c-96f7-44e2-8bb8-790bedf066b1

Opdenbusch, J. C., Hielscher, J., & Sasse, M. A. (2025). "Where are we on cyber?": A qualitative study on boards' cybersecurity risk decision making. In *Proceedings of the 2025 Network and Distributed System Security Symposium (NDSS 2025)*. Internet Society. https://doi.org/10.14722/ndss.2025.240595

Pezzo, M. V., Nash, B. E. D., Vieux, P., & Foster-Grammer, H. W. (2022). Effect of having, but not consulting, a computerized diagnostic aid. *Medical Decision Making*, *42*(1), 94–104. https://doi.org/10.1177/0272989X211011160

Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, *46*(5), 1081–1100. https://doi.org/10.1111/pops.13059

Reif, J. A., Larrick, R. P., & Soll, J. B. (2025). Evidence of a social evaluation penalty for using AI. *Proceedings of the National Academy of Sciences*, *122*(19), Article e2426766122. https://doi.org/10.1073/pnas.2426766122

Schilke, O., Reimann, M., & Cook, K. S. (2015). Power decreases trust in social exchange. *Proceedings of the National Academy of Sciences*, *112*(42), 12950–12955. https://doi.org/10.1073/pnas.1517057112

Sharma, K., Schoorman, F. D., & Ballinger, G. A. (2023). How can it be made right again? A review of trust repair research. *Journal of Management*, *49*(1), 363–399. https://doi.org/10.1177/01492063221089897

Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, *26*(2), 200–215. https://doi.org/10.1111/j.1467-6494.1958.tb01580.x

U.S. Securities and Exchange Commission. (2023). *Cybersecurity risk management, strategy, governance, and incident disclosure* (Release Nos. 33-11216; 34-97989). https://www.sec.gov/files/rules/final/2023/33-11216.pdf

von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, *27*(7), Article 110201. https://doi.org/10.1016/j.isci.2024.110201

