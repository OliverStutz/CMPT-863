
## What approaches, technologies and theories already exist

### Introduction

Company boards are obligated to make Cyber Security decisions, and most do so without the expertise to judge it (Lowry et al., 2026). Regulators like the U.S. Securities and Exchange Commission (SEC, 2023) require public companies to document board oversight processes related to cyber risk. In the European Union's NIS2 Directive it states that members of management bodies have to be trained to assess cybersecurity risk management themselves (Directive EU 2022/2555, 2022, Art. 20). Agency theory (Lowry et al., 2026) expects the board to monitor executives whose interests deviate from the companies like CISO's filtering their reports to look better. The institutional theory adopts practices given by regulators to make the oversight appear more legitimate(while it should not be for that sole purpose). Tools that could augment board members are already sold, like Meta's smart glasses that answer specific questions about what the wearer sees (Meta, 2024) and meeting assistants which transcribe the voice and evaluate it in real time like Microsoft Teams with Copilot (Microsoft, n.d.). Unfortunately current papers haven't studied the augmented human (with checker) to human trust impact.

### Trust theory and its measurement

Trust is the willingness to be vulnerable to another party (Mayer et al., 1995). The same willingness can be placed in a device, but the studies reviewed here concern trust between people, which rests on the other person's perceived ability, benevolence and integrity and shows itself in risk taking within the relationship; in this study that act is the board's decision whether to grant the CISO's budget (Mayer et al., 1995). Our three trust items follow the Muenster Epistemic Trustworthiness Inventory (METI), which measures laypeople's trust in an expert on the same three factors (Hendriks et al., 2015).

## What we know from prior work

### Boards without expertise

Field studies agree that directors without expertise ask questions but cannot judge the answers. Lowry et al. (2026) interviewed directors, security executives and consultants and found that expertise decides which theory fits: boards with expertise have been monitoring, while boards without it performed practices that looked legitimate, this was refered as symbolic oversight. Non-expert directors believed their business experience let them ask the right questions; expert participants saw them unable to judge the answers and left to assess the credibility of the executive in front of them. Thirty-three surveyed CISOs in the same study estimated, at the median, that 40% of CISOs filter their reports to make themselves or their superiors look better; most non-expert directors thought filtering was unlikely. In the United Kingdom, directors held back questions rather than possibly exposing their lack of knowledge, the budget was by most accounts the only cyber decision boards made, and that reluctance left CISOs to decide what reached the board (Opdenbusch et al., 2025). Lowry et al. read the same finding as a gap in expertise rather than power; both studies agree that the CISO controls what the board hears. An independent view changes the picture: Australian directors challenged management more confidently where external assurance was routine, and without cyber literacy boards often accepted management's account (Aamir et al., 2026).

### Checkers and their accuracy

Checking tools held by the listener exist as products (as mentioned above) and as research systems, and evaluations of related checkers report limited accuracy. Audio recording glasses, using machine learning, that judged whether a heard claim was backed by evidence were right on 63.5% of statements in live use (Danry et al., 2020). At the 2020 State of the Union address, the live checker Squash showed 14 relevant and 20 irrelevant fact-checks (Adair, 2020). GPT-4, given chart images and multiple-choice questions about them, scored lower than about five in six people. It knew some rules of good chart design, but when asked to read one number off a chart it was right in only 3 of 13 cases (Bendeck & Stasko, 2025). A flag can therefore cite a sound rule and still misread the slide.

### Trust repair with human and machine accusers

With human accusers, how the accused replies shapes how far trust recovers, a process the literature calls trust repair. After a hearsay accusation of incompetence, an apology restored more trust than a denial with an account (P. H. Kim et al., 2004), and a reply that neither confirmed nor denied the accusation did no better than the weaker of the two (Ferrin et al., 2007). When a third party exonerated an accused researcher, trust returned to the level of a researcher never accused, but not reliably above it (Greitemeyer & Sagioglou, 2015). Sharma et al. (2023) count a rise beyond the starting level as further development rather than repair. They also note that most studies of trust repair find trust still below where it started, and that denial works less well after a competence violation than after an integrity violation. With a machine accuser the picture may change. Kempt and Nagel (2022) argue, in a point Chen et al. (2026) take up, that an AI's contradiction shifts the burden of proof to the expert, and that giving reasons cannot settle the dispute because the machine does not respond to them. The nearest evidence on an expert's reason about a machine is indirect: after a poor outcome, non-expert readers rated a physician who had not consulted an available diagnostic aid as more competent when the physician believed they already knew what it would say than when they distrusted a consultation to the aid (Pezzo et al., 2022).

## What remains unresolved or underexplored

I found no study that gave a director a checking tool and measured its effect on trust between two people, and the two fields psychology and HCI do not meet. Trust research has a human accuse and a human reply, with no tool (Ferrin et al., 2007; P. H. Kim et al., 2004). Checker research in human-computer interaction mostly measures belief in claims or trust in the AI (Danry et al., 2020; Lu et al., 2025), and the few that rated the person checked had the expert consult the AI or let it run alone (Chen et al., 2026; D. Kim et al., 2026), never a listener use it. Wrong-flag studies (see the Appendix) leave the person judged silent and rated once, so it is unknown what a wrong flag does to trust in a correct expert, whether trust repairs after a machine accusation, whether a bare denial lowers trust further or only repairs less, and whether a reason meets the burden of proof a machine's contradiction places on the expert (Kempt & Nagel, 2022).

A board that cannot judge the substance may heed a flag whether or not it is right, and power works against trust: in exchange experiments, the party with better alternatives trusted its partner less (Schilke et al., 2015). A wrong flag may cost a correct CISO trust and budget; a well-answered one may win approval for an unneeded project. No study combines a wrong flag on a correct expert, a checker held by a non-expert with power over that expert, the expert's reply, and trust rated at each step, in a boardroom.

## How this study differs from and extends existing research

### The closest designs

The closest design is D. Kim et al. (2026): an AI checked a human advisor in front of the student who depended on him, and trust in the advisor was measured with METI. Their AI was always right, the advisor never replied, and the student held neither the tool nor any power over the advisor. In Chen et al. (2026) the doctor brought in the AI himself and never answered. Gupta et al. (2026), the nearest real-time wearable checker, measured belief in claims and trust in the AI, never trust in the speaker. And only 8 of 94 empirical repair studies used any within-subject design, although a claim of repair needs ratings before the violation, after it and after the reply (Sharma et al., 2023).

### What this study adds

This study adds the experiment to the board research and the tool to the trust research. A participant acts as a board member who approves/disapproves the CISO's budget and holds a checker during the briefing; every flag is a false positive, and the CISO answers. Trust in the CISO is rated before and after every slide in the simulation and will be compared with board members who hear the same briefing without a checker (second study group). The result is a measure neither field has taken: trust between two people when one of them is augmented with a checker.

The checker will highlight pre-defined text on the presentation, with pre-defined information for the board member. So that all participants will experience the same "AI" behaviour.

Our study differs from these designs on each point. The flags are wrong and the CISO is right; the CISO answers; the listener holds both the tool and the budget; and trust is rated at every step. Because accusations of dishonesty and visible confidence figures change how people respond (Loconte et al., 2026), every flag concerns the slide and never the CISO's honesty, the panel shows no accuracy or confidence figure, and the CISO never mentions AI (Reif et al., 2025).

Our study puts the three open questions on trust recovery to the test. The CISO answers each flag with a reason or a bare denial, delivered by the same person in the same tone; trust is rated after the flag and after the answer; a neutral group hears the same answers without flags; and a rise above the first rating is treated as exploratory, since repaired trust rarely exceeded its starting level (Greitemeyer & Sagioglou, 2015; Sharma et al., 2023). A bare denial is expected to repair less, since even a denial with an account was the weaker reply (P. H. Kim et al., 2004).

This study combines all four elements of the gap, and asks whether a wrong flag costs a correct expert trust and whether a reason recovers it. It will shed light to the question if checker augmented humans with no expert knowledge will by the act of using the checker lower the trust in the CISO. This is a very common situation nowadays and it was surprising not to find exact research in this area, where many consulting firms depend on exactly this constellation to get further engagement.

## References

Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From oversight to insight: Transforming cybersecurity governance in boardrooms. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–22). Association for Computing Machinery. https://doi.org/10.1145/3772318.3791142

Adair, B. (2020, February 23). *Squash report card: Improvements during State of the Union … and how humans will make our AI smarter*. Duke Reporters' Lab. https://reporterslab.org/2020/02/23/squash-report-card-improvements-during-state-of-the-union-and-how-humans-will-make-our-ai-smarter/

Bendeck, A., & Stasko, J. (2025). An empirical evaluation of the GPT-4 multimodal language model on visualization literacy tasks. *IEEE Transactions on Visualization and Computer Graphics*, *31*(1), 1105–1115. https://doi.org/10.1109/TVCG.2024.3456155

Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, *213*, Article 103824. https://doi.org/10.1016/j.ijhcs.2026.103824

Danry, V., Pataranutaporn, P., Mao, Y., & Maes, P. (2020). Wearable Reasoner: Towards enhanced human rationality through a wearable device with an explainable AI assistant. In *Proceedings of the Augmented Humans International Conference* (pp. 1–12). Association for Computing Machinery. https://doi.org/10.1145/3384657.3384799

Directive (EU) 2022/2555 of the European Parliament and of the Council of 14 December 2022 on measures for a high common level of cybersecurity across the Union, amending Regulation (EU) No 910/2014 and Directive (EU) 2018/1972, and repealing Directive (EU) 2016/1148 (NIS 2 Directive). (2022). *Official Journal of the European Union*, L 333, 80–152. https://eur-lex.europa.eu/eli/dir/2022/2555/oj

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

## Appendix Memo and AI Transparency

In my literature review, i've used AI to structure, formulate, edit and rewrite my ideas.
I've had to move some section due to size into the Appendix "What we know from prior work" contained "Wrong flags and the person contained." By the time i was finished i realised it was too long but i did already invest the effort in the section and the research and therefore i kept it in the appendix.

## Appendix: Wrong flags, and the person flagged

Four studies show that people believe AI warnings even when the warnings are wrong, and that people with power over the person checked act on them. Readers trusted a true news headline less when it carried a wrong AI warning calling it false, by 0.30 on a scale from −1 to +1 (Horne & Nevo, 2025). An AI label saying whether a headline was credible moved readers' judgments more than an expert's vote did, and readers could not tell the AI's correct predictions from its wrong ones (Lu et al., 2025). Teachers who read a report from a fictitious AI detector claiming that 87% of a student paper was AI-written graded the identical paper 1.8 points lower on a 10-point scale than teachers whose report claimed 7%, and marking the suspect sentences in red lowered the grade further (Du et al., 2026). People judging whether a written statement was a lie accused the writer in 19% of cases on their own, but those who chose to consult a lie-detection algorithm accused in almost 85% of the cases in which it called the statement a lie (von Schenk et al., 2024).

The person checked pays too, even when the check is right or finds nothing wrong. Supervisors who had watched one worker closely and another rarely trusted the closely watched worker less, although both had produced the same work (Strickland, 1958). A correct fact-check of a fictional candidate in an unfamiliar country made participants rate the candidate as less competent and less moral (Prike et al., 2025). People imagining a cancer screening rated a radiologist who recommended less testing than an AI lowest of four conditions, while a radiologist who recommended more than the AI paid no penalty (Madanay et al., 2025). When a doctor consulted an AI assistant and the AI disagreed, participants saw the doctor as lazier and trusted the doctor less (Chen et al., 2026). Employees described as getting help from AI were rated lazier and less competent opposed to the ones described as getting help from other sources (Reif et al., 2025).

Two studies show that a flag does not always cost trust. In D. Kim et al. (2026), an AI corrected a human advisor's mistake in front of a student. Trust in the advisor fell when the advisor had erred, and it made no significant difference whether the AI was absent, called in by the advisor, or watching on its own. The mistake cost the trust, not the AI's correction. In Loconte et al. (2026), an AI judged written statements as true or lies, with its accuracy and its confidence shown on screen. The more confident the AI was that a statement was a lie, the less people followed it. A flag can be too strong to be believed. Studies that fact-check well-known politicians are left out of this review: people already hold fixed views of those figures, so those results say nothing about a listener judging a presenter whose reputation is still forming.
