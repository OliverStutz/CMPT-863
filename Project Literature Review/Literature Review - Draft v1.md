# Related work

A board member points a camera assistant at the CISO's slide, the tool flags a fault, and the CISO answers. Five strands of work bear on this moment: board oversight, the checkers on offer, the cost of wrong flags, trust repair, and the gap where they meet.

## Boards oversee what they cannot judge

Regulation asks non-experts to oversee an expert. The U.S. Securities and Exchange Commission (SEC, 2023) chose not to require disclosure of board cyber expertise, holding that "directors with broad-based skills in risk management and strategy often effectively oversee management's efforts without specific subject matter expertise". In Lowry et al. (2026), non-expert directors' oversight was "mostly symbolic" even when they did what expert directors did, and one expert director said that for most boards "all they can do is assess the credibility of the executives that are put in front of them". Proudfoot et al. (2023) agree: the board's view of the CISO's competence is "heavily determined by the CISO's ability to effectively communicate". Directors are "too shy to ask questions since they do not want to get embarrassed for their missing knowledge", and most interviewees named the budget as the board's only cyber decision (Opdenbusch et al., 2025). Directors pushed back more confidently where independent reviews were routine (Aamir et al., 2026).

The sources disagree on power. Opdenbusch et al. (2025) conclude that "the CISOs are the ones in power", because they decide what reaches the board; Lowry et al. (2026) trace the deference to low expertise instead. We read this as formal power with the board and informational power with the CISO. In exchange experiments, the party with better alternatives trusted its partner less, and from the same starting level its trust fell further during a negotiation (Schilke et al., 2015). Their power was having alternatives, not authority. The board evidence is from interviews, with no tool and no trust measure; we put an independent check in the board member's hands and measure trust in the CISO.

## Checkers are available, and often wrong

The listener-held checker already exists as a product. Ray-Ban Meta glasses answer questions about what the wearer is looking at (Meta, 2024), Apple's Visual Intelligence lets users ask for "an explanation of a diagram from class notes" (Apple, 2024), and Copilot in Microsoft Teams suggests the private prompt "Where are some holes in [a meeting participant's] argument?" (Microsoft, n.d.). Such checkers are often wrong. Audio glasses that tell the wearer whether a heard claim is backed by evidence were right on 63.5 percent of statements when running live (Danry et al., 2020). At the 2020 State of the Union the live checker Squash showed 14 relevant fact-checks and 20 irrelevant ones, and its lab began building an interface for a human editor to approve each pop-up (Adair, 2020). GPT-4V scored at about the 16th percentile of humans on a visualization literacy test and showed some knowledge of design best practice, yet answered only 13 percent of questions that needed a value read off a chart when it lacked the data (Bendeck & Stasko, 2025). A flag can thus be right by the textbook and wrong for this company. None of this work reports what a check does to trust in the speaker; Danry et al. rated recorded statements, not a person. Our flags are drafted with a real VLM and are all wrong for the company.

## People follow a wrong flag, and the person flagged pays

Wrong flags are believed. A wrong AI warning on a true headline cut trust in it by 0.30 on a scale from −1 to +1 (Horne & Nevo, 2025). An AI credibility indicator had a larger effect on accuracy than an expert's vote, and in the authors' words people "lack the capability to differentiate correct and incorrect AI predictions" (Lu et al., 2025). Flags also move people with power to act. A fictitious report of a high AI-detection rate made 214 teachers grade an identical paper 1.8 points lower on a 10-point scale, and red highlighting made the report seem more credible and the drop larger (Du et al., 2026). People who chose to consult a lie-detection algorithm accused in almost 85 percent of cases when it said "lie" (von Schenk et al., 2024).

The person flagged pays. Even checks that find nothing can cost trust: more supervisors trusted a rarely checked worker than an often checked one (Strickland, 1958). A human fact-check lowered an unfamiliar candidate's ratings of competence and morality (Prike et al., 2025). When an AI and a radiologist disagreed, 1,606 adults agreed less with the radiologist who recommended less testing than the AI and rated that radiologist lowest of the four conditions; one who recommended more paid no penalty (Madanay et al., 2025). An AI assistant's disagreement made a doctor seem lazier and lowered trust in him (Chen et al., 2026). Employees described as getting help from AI were rated lazier and less competent than those who got the same help elsewhere (Reif et al., 2025), so our CISO never mentions AI.

Two results set limits. In D. Kim et al. (2026) the advisor's error lowered trust, but how the AI was involved changed nothing, so their data tie the loss to the error, not to the AI's contradiction. People resisted confident AI accusations of lying when the model's accuracy and confidence were on screen (Loconte et al., 2026); our flags concern the slide, not honesty, and show neither figure. Results on well-known politicians, whose reputations are settled, are not carried over. In all of these studies the expert stays silent and trust is rated once.

## Trust in an expert, and how it is repaired

Mayer et al. (1995) define trust as the willingness to be vulnerable to another party, based on perceived ability, benevolence and integrity, and trust leads to risk taking, here the budget decision. Our three items come from the METI, built on the same factors for laypeople judging an expert (Hendriks et al., 2015).

After a hearsay accusation of a competence violation, an apology repaired trust better than a denial, and that denial came with an account (P. H. Kim et al., 2004). A reply that neither confirmed nor denied did no better than the weaker of the other two (Ferrin et al., 2007). Clearing by a third party restored an accused researcher to the level of one never accused, but not reliably above it (Greitemeyer & Sagioglou, 2015). Rating after every step may draw attention to the latest information (Mickelberg et al., 2025), so both our groups rate at the same steps. Sharma et al. (2023) call a return to the starting level complete repair and a rise beyond it "further development of trust". Most studies they review find that repair falls short of the start, and denial works less well after competence than after integrity violations. Hence H2 treats a rise above the first rating as exploratory, and H3 expects a bare denial to repair less.

In these studies the accuser is human. Kempt and Nagel (2022) argue, in a point Chen et al. (2026) take up, that when an AI contradicts the expert in charge the burden of proof moves to the expert. The expert cannot settle the disagreement by reasons, because the machine does not respond to them, so they propose that a second human settle it. Jongsma and Sand (2022) reply that agreement carries the same burden. Our own hypothesis is that the CISO's reason, addressed to a board member who can weigh it, may meet the burden in the listener's eyes, while a bare denial does not. A physician who had not consulted an available diagnostic aid before a bad outcome was rated higher when the physician believed he or she already knew what it would say than when the physician did not trust it, and in neither case below a physician with no aid mentioned (Pezzo et al., 2022).

## The gap and this study

No study combines these pieces. D. Kim et al. (2026) is the closest design: an AI monitored a human advisor in front of the student who depended on him, and trust in the advisor was measured with METI (*N* = 157). Their AI was always right, the advisor never replied, and the student held neither the tool nor any power over the advisor. Gupta et al. (2026), the nearest real-time wearable checker, measured belief in claims and trust in the AI and never rated the person speaking. Only 8 of 94 empirical repair studies used any within-subject design, though a claim of repair needs ratings before the violation, after it and after the reply (Sharma et al., 2023). None is set in a boardroom. We bring the pieces together: the checker is in the hands of the non-expert who holds power over the expert, its flags are plausible but wrong for the company, the expert answers each flag with a reason or a bare denial, and trust is rated after each flag and each answer, against a neutral group that hears the same answers without flags.

## References

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–22). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791142

Adair, B. (2020, February 23). *Squash report card: Improvements during State of the Union … and how humans will make our AI smarter*. Duke Reporters' Lab. https://reporterslab.org/2020/02/23/squash-report-card-improvements-during-state-of-the-union-and-how-humans-will-make-our-ai-smarter/

Apple. (2024, December 11). *Apple Intelligence now features Image Playground, Genmoji, and more*. https://www.apple.com/newsroom/2024/12/apple-intelligence-now-features-image-playground-genmoji-and-more/

Bendeck, A., & Stasko, J. (2025). An empirical evaluation of the GPT-4 multimodal language model on visualization literacy tasks. *IEEE Transactions on Visualization and Computer Graphics*, *31*(1), 1105–1115. https://doi.org/10.1109/TVCG.2024.3456155

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, *213*, Article 103824. https://doi.org/10.1016/j.ijhcs.2026.103824

Danry, V., Pataranutaporn, P., Mao, Y., & Maes, P. (2020). Wearable Reasoner: Towards enhanced human rationality through a wearable device with an explainable AI assistant. In *Proceedings of the Augmented Humans International Conference* (pp. 1–12). Association for Computing Machinery. https://doi.org/10.1145/3384657.3384799

Du, P., Liu, T., & Xian, X. (2026). Automation bias in teachers' evaluation of student writing: Effects of algorithmic warnings and visual risk cues in AI detection reports. *Frontiers in Psychology*, *17*, Article 1889402. https://doi.org/10.3389/fpsyg.2026.1889402

Ferrin, D. L., Kim, P. H., Cooper, C. D., & Dirks, K. T. (2007). Silence speaks volumes: The effectiveness of reticence in comparison to apology and denial for responding to integrity- and competence-based trust violations. *Journal of Applied Psychology*, *92*(4), 893–908. https://doi.org/10.1037/0021-9010.92.4.893

Greitemeyer, T., & Sagioglou, C. (2015). Does exonerating an accused researcher restore the researcher's credibility? *PLOS ONE*, *10*(5), Article e0126316. https://doi.org/10.1371/journal.pone.0126316

Gupta, C., Aritonang, N. V., Rajendran, D. P. D., Danry, V., Maes, P., & Nanayakkara, S. (2026). Feeling the Facts: Real-time wearable fact-checkers can use nudges to reduce user belief in false information. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems*. Association for Computing Machinery. https://doi.org/10.1145/3772318.3791679

Hendriks, F., Kienhues, D., & Bromme, R. (2015). Measuring laypeople's trust in experts in a digital age: The Muenster Epistemic Trustworthiness Inventory (METI). *PLOS ONE*, *10*(10), Article e0139309. https://doi.org/10.1371/journal.pone.0139309

Horne, B. D., & Nevo, D. (2025). People adhere to content warning labels even when they are wrong due to ecologically rational adaptations. *Scientific Reports*, *15*(1), Article 13896. https://doi.org/10.1038/s41598-025-98221-7

Jongsma, K. R., & Sand, M. (2022). Agree to disagree: The symmetry of burden of proof in human–AI collaboration. *Journal of Medical Ethics*, *48*(4), 230–231. https://doi.org/10.1136/medethics-2022-108242

Kempt, H., & Nagel, S. K. (2022). Responsibility, second opinions and peer-disagreement: Ethical and epistemological challenges of using AI in clinical diagnostic contexts. *Journal of Medical Ethics*, *48*(4), 222–229. https://doi.org/10.1136/medethics-2021-107440

Kim, D., Daneshi, R., Draper, B., & Sreedharan, S. (2026). Asymmetric trust effects of corrective AI in expert advisory workflows under epistemic dependence. *Proceedings of the AAAI/ACM Conference on AI, Ethics, and Society*, *9*(3), 1514–1525. https://doi.org/10.1609/aies.v9i3.43244

Kim, P. H., Ferrin, D. L., Cooper, C. D., & Dirks, K. T. (2004). Removing the shadow of suspicion: The effects of apology versus denial for repairing competence- versus integrity-based trust violations. *Journal of Applied Psychology*, *89*(1), 104–118. https://doi.org/10.1037/0021-9010.89.1.104

Loconte, R., Monaro, M., Pietrini, P., Verschuere, B., & Kleinberg, B. (2026). Humans incorrectly reject confident accusatory AI judgments. *Computers in Human Behavior*, *182*, Article 109019. https://doi.org/10.1016/j.chb.2026.109019

Lowry, M. R., Vance, A., & Vance, M. D. (2026). Inexpert supervision: Field evidence on boards' oversight of cybersecurity. *Management Science*, *72*(2), 783–804. https://doi.org/10.1287/mnsc.2023.04147

Lu, Z., Li, P., Wang, W., & Yin, M. (2025). Understanding the effects of AI-based credibility indicators when people are influenced by both peers and experts. In *Proceedings of the 2025 CHI Conference on Human Factors in Computing Systems* (pp. 1–19). Association for Computing Machinery. https://doi.org/10.1145/3706598.3713871

Madanay, F. L., O'Donohue, L. S., & Zikmund-Fisher, B. J. (2025). Patient reactions to artificial intelligence–clinician discrepancies: Web-based randomized experiment. *Journal of Medical Internet Research*, *27*, Article e68823. https://doi.org/10.2196/68823

Mayer, R. C., Davis, J. H., & Schoorman, F. D. (1995). An integrative model of organizational trust. *Academy of Management Review*, *20*(3), 709–734. https://doi.org/10.2307/258792

Meta. (2024, April 23). *New Ray-Ban | Meta smart glasses styles and Meta AI updates*. https://about.fb.com/news/2024/04/new-ray-ban-meta-smart-glasses-styles-and-meta-ai-updates/

Mickelberg, A. J., Walker, B., Ecker, U. K. H., Howe, P. D. L., Perfors, A., & Fay, N. (2025). Did he or didn't he? Mixed evidence for the continued influence of retracted misinformation on person impressions. *PLOS ONE*, *20*(5), Article e0322045. https://doi.org/10.1371/journal.pone.0322045

Microsoft. (n.d.). *Use Copilot in Microsoft Teams meetings* [Support article]. Retrieved October 18, 2026, from https://support.microsoft.com/en-us/office/use-copilot-in-microsoft-teams-meetings-0bf9dd3c-96f7-44e2-8bb8-790bedf066b1

Opdenbusch, J. C., Hielscher, J., & Sasse, M. A. (2025). "Where are we on cyber?": A qualitative study on boards' cybersecurity risk decision making. In *Proceedings of the 2025 Network and Distributed System Security Symposium (NDSS 2025)*. Internet Society. https://doi.org/10.14722/ndss.2025.240595

Pezzo, M. V., Nash, B. E. D., Vieux, P., & Foster-Grammer, H. W. (2022). Effect of having, but not consulting, a computerized diagnostic aid. *Medical Decision Making*, *42*(1), 94–104. https://doi.org/10.1177/0272989X211011160

Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*, *46*(5), 1081–1100. https://doi.org/10.1111/pops.13059

Proudfoot, J. G., Cram, W. A., Madnick, S., & Coden, M. (2023). The importance of board member actions for cybersecurity governance and risk management. *MIS Quarterly Executive*, *22*(4), 235–250. https://doi.org/10.17705/2msqe.00084

Reif, J. A., Larrick, R. P., & Soll, J. B. (2025). Evidence of a social evaluation penalty for using AI. *Proceedings of the National Academy of Sciences*, *122*(19), Article e2426766122. https://doi.org/10.1073/pnas.2426766122

Schilke, O., Reimann, M., & Cook, K. S. (2015). Power decreases trust in social exchange. *Proceedings of the National Academy of Sciences*, *112*(42), 12950–12955. https://doi.org/10.1073/pnas.1517057112

Sharma, K., Schoorman, F. D., & Ballinger, G. A. (2023). How can it be made right again? A review of trust repair research. *Journal of Management*, *49*(1), 363–399. https://doi.org/10.1177/01492063221089897

Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, *26*(2), 200–215. https://doi.org/10.1111/j.1467-6494.1958.tb01580.x

U.S. Securities and Exchange Commission. (2023). *Cybersecurity risk management, strategy, governance, and incident disclosure* (Release Nos. 33-11216; 34-97989). https://www.sec.gov/files/rules/final/2023/33-11216.pdf

von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, *27*(7), Article 110201. https://doi.org/10.1016/j.isci.2024.110201

## Writer's notes

**Word count (body only, without title, headings and references):** opening 47; boards 278; checkers 226; wrong flags 360; trust and repair 375; gap 186; total 1,472. This is above the 1,330 target and under the 1,500 limit. The easiest further cuts are the Apple sentence, the Jongsma and Sand clause and the Mickelberg sentence. Kim et al. (2004) and Ferrin et al. (2007) rating trust once, after the reply, was cut for length; the gap section carries the point through Sharma et al. (2023).

**Sources cited:** 34 (the 29 core sources of the structure, with the camera assistants as two entries, Meta and Apple; plus Microsoft, Adair and Bendeck and Stasko counted among them; plus Jongsma and Sand, 2022). Every in-text citation is in the list and every list entry is cited.

**Deviations from the structure, and why:**
- Pezzo et al. (2022): only the abstract was read, so only the cautious clause from `read-pezzo-schilke-sharma.md` is used, reworded closer to the abstract ("believing", "not trusting", "no aid mentioned"). It is not called "the only study", gives no numbers, and is not presented as a reply paradigm or as support for H2. The gap section no longer names it.
- Sharma et al. (2023): "further development of trust" replaces "over-repair"; the 8 of 94 within-subject count is in the gap section. The optional Shaw et al. (2003) clause is dropped.
- Kempt and Nagel (2022): presented as the burden moving to the expert, who cannot settle it by reasons, with a second human to settle it; "a reason meets the burden" is stated as our own hypothesis. Jongsma and Sand (2022) added in one clause as the stated disagreement, as the read file suggests.
- Danry et al. (2020): used only for the 63.5 percent live accuracy and the fact that statements, not a person, were rated; "changed how users rated a speaker's arguments" is not used.
- Schilke et al. (2015): used for the party with better alternatives trusting less and for trust falling further during the interaction for the powerful party, with a sentence that their power is alternatives, not authority; not used to call power a variable in our study.
- Optional sources Da Silva and Jensen (2022), Hancock et al. (2020) and Wang et al. (2024) left out for length. Google Gemini Live left out because the results entry cites only a secondary source.
- In-text citations of the two Kim papers carry initials (D. Kim et al., 2026; P. H. Kim et al., 2004), as APA 7 requires for different first authors with the same surname; Abstract v8 does not do this.
- Section 1 states "how we differ" briefly (interviews, no tool, no trust measure); the outcomes of our study are named in the gap section rather than repeated.

**Claims left out because unsure:** the "No issues found" prediction from Proudfoot's green dashboards; Chen et al. listing "the expert answers" as future work (reported only by a novelty reviewer); Pezzo's "consulted but ignored" physician (abstract only, and the brief allows only the cautious clause); Schilke's speculation that low-power people are less forgiving; Lowry's 40 percent filtering estimate and the NACD handbook point (not needed for the argument); Horne and Nevo's below-chance result with all-wrong labels (cut for length, not for doubt).
