# Results, Agent 2: Human-AI decision making, reliance, automation bias and trust calibration

Searched 2026-10-18 by Agent 2. 24 scored entries, plus 13 lower-priority entries (also verified) listed in short form at the end. Every DOI was resolved against `https://api.crossref.org/works/<DOI>` on 2026-10-18; title, authors, year, venue, volume, issue and pages below are taken from the Crossref record. Abstracts were read from Crossref, OpenAlex, PubMed, arXiv, the ACL Anthology or the publisher page, as noted per entry.

Overall reading for our study. This literature measures trust in and reliance on the *AI*, almost never trust in a *human* the AI disagrees with. Its usable findings for us are: (a) people follow wrong AI advice even when it contradicts context they hold, and explanations attached to AI output raise acceptance whether the output is right or wrong (supports the mechanism behind H1 and H4); (b) a second opinion that disagrees with the AI cuts over-reliance (relevant to H2, the CISO's reasoned answer acting as the counter-opinion); (c) in judge-advisor studies, a human expert advisor and an algorithm are weighted differently depending on framing and perceived competence, and bad advice costs the algorithm more than the human (relevant to who loses standing when the two conflict). Only Hou and Jung (2021) measures perceived competence of the human agent as a mediator. The gap: no study in this stream has a non-expert judge holding the AI, the AI contradicting a human expert in front of that judge, and trust in the *human* measured before and after the human's reply.

---

## Sub-theme A: Algorithm versus human advisor, and advice conflict (judge-advisor systems)

### Hou, Y. T.-Y., & Jung, M. F. (2021). Who is the expert? Reconciling algorithm aversion and algorithm appreciation in AI-supported decision making. *Proceedings of the ACM on Human-Computer Interaction*, *5*(CSCW2), 1–25. https://doi.org/10.1145/3479864
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1145/3479864
- **Design:** Three online decision-making experiments in which participants received input from a human agent and an algorithmic agent. Only the textual description (framing) of the human and the algorithm was manipulated. Measured: which agent's input participants relied on, and the agent's perceived competence (expert power) as a mediator.
- **Finding:** Changing only the descriptions produced algorithm appreciation in Experiment 1 and algorithm aversion in Experiment 2; Experiment 3 showed how different framing choices reproduce the inconsistent results in earlier studies. The preference for one agent over the other ran through how competent each agent was perceived to be. The authors propose French and Raven's power sources (expert power) as the frame for algorithmic decision support.
- **Relation:** Advice conflict; H1 and H4 (whether a checker flag outweighs the CISO depends on how each is framed and on perceived competence); also a design element: our checker's description ("vision-language model") and the CISO's introduction are framings that will set the starting balance. Our study differs because the algorithm and the human are not alternatives offered to the judge: the algorithm comments on the human, and we measure trust in the human, not choice of advisor.
- **Access:** SFU library (ACM DL); author preprint may exist
- **Score:** 1=2 2=3 3=2 4=5 5=3 total=15/25

### Önkal, D., Goodwin, P., Thomson, M., Gönül, S., & Pollock, A. (2009). The relative influence of advice from human experts and statistical methods on forecast adjustments. *Journal of Behavioral Decision Making*, *22*(4), 390–409. https://doi.org/10.1002/bdm.637
- **Type:** psychology
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1002/bdm.637
- **Design:** Stock-price forecasting. Study 1: two groups received identical point and interval forecasts, labelled either as from a human expert or from a statistical method, and adjusted their own forecasts. Study 2: participants received advice from two sources, both human, both statistical, or one of each. Measured: how far participants moved toward each piece of advice.
- **Finding:** In Study 1 the advice improved accuracy in both groups, but advice labelled as statistical was discounted much more. In Study 2, when both sources were the same kind, they had equal influence; when one was a human expert and one a statistical method, participants paid much more attention to the human expert.
- **Relation:** Advice conflict; counter-result for H1: when a human expert and an algorithm sit side by side, judges have leaned toward the human. Our design puts them in direct conflict on the same claim, and the human is also the target of the algorithm's flag, which these studies never tested.
- **Access:** open access https://onlinelibrary.wiley.com/doi/pdfdirect/10.1002/bdm.637
- **Score:** 1=1 2=3 3=2 4=4 5=4 total=14/25

### Prahl, A., & Van Swol, L. (2017). Understanding algorithm aversion: When is advice from automation discounted? *Journal of Forecasting*, *36*(6), 691–702. https://doi.org/10.1002/for.2464
- **Type:** management or economics
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1002/for.2464
- **Design:** Forecasting experiment in which advice came from a human advisor or from automation, including rounds of bad advice. Measured: advice utilization and perceived similarity to the advisor.
- **Finding:** Human advice was not used more than automated advice overall (the earlier result did not replicate). After bad advice, use of the automated advice fell significantly more than use of the human advice. Participants described themselves as having much more in common with the human advisor than with the automation, although there was no relationship with either.
- **Relation:** Advice conflict; H2 and H3 context: an error is forgiven more readily in a human than in an algorithm, so a checker that is shown to be wrong (the CISO's reasoned answer) may lose more standing than the CISO. Our study measures the human's standing after the algorithm errs about the human, which this work does not.
- **Access:** SFU library
- **Score:** 1=2 2=2 3=1 4=4 5=4 total=13/25

### Bonaccio, S., & Dalal, R. S. (2006). Advice taking and decision-making: An integrative literature review, and implications for the organizational sciences. *Organizational Behavior and Human Decision Processes*, *101*(2), 127–151. https://doi.org/10.1016/j.obhdp.2006.07.001
- **Type:** psychology
- **Verified:** Crossref 2026-10-18. Abstract not available from Crossref, OpenAlex or Semantic Scholar (publisher elided), and ScienceDirect returned 403. Summary is from secondary descriptions: the ALNAP library record (https://alnap.org/help-library/resources/advice-taking-and-decision-making-an-integrative-literature-review-and-implications-for/) and Bailey et al. (2022, below), which uses its framework.
- **Design:** Integrative review of the advice-giving and advice-taking literature built on the judge-advisor system (JAS): a judge makes the final decision, one or more advisors give input.
- **Finding:** Organises the field into an input-process-output model: inputs at the individual, JAS and environment level; outputs are the weight the judge puts on advice and the judge's post-advice accuracy and confidence. Reviews advice utilization, confidence, accuracy, advisor-judge differences and design choices such as the number of advisors, and questions the common operational definition of "advice".
- **Relation:** Definition and theme anchor: the board member is a judge, the CISO and the checker are two advisors who disagree. The JAS literature measures weight on advice, not trust in the advisor, which is the gap our trust ratings fill.
- **Access:** SFU library
- **Score:** 1=1 2=0 3=2 4=4 5=2 total=9/25

### Bailey, P. E., Leon, T., Ebner, N. C., Moustafa, A. A., & Weidemann, G. (2023). A meta-analysis of the weight of advice in decision-making. *Current Psychology*, *42*(28), 24516–24541. https://doi.org/10.1007/s12144-022-03573-2
- **Type:** psychology
- **Verified:** Crossref 2026-10-18 (online 2022-08-09, issue 2023); abstract read at https://api.openalex.org/works/doi:10.1007/s12144-022-03573-2
- **Design:** Meta-analysis of 346 effect sizes from 129 judge-advisor datasets (N = 17,296).
- **Finding:** The only unique predictor of how much people took advice was information about the advisor that suggested advice quality. People moved their estimates 32%, 37% and 48% toward advice from advisors described as low, neutral or high quality. Advice was taken more for subjective and uncertain estimates than for objective ones. Age, gender and individualism had no effect.
- **Relation:** Advice conflict; H4 measure context: how the checker is described will drive how much its flag weighs, and our board judgements (is this architecture right for this company) are the subjective, uncertain kind where advice weighs most. Weight on advice is a behaviour measure; we measure trust in the human advisor.
- **Access:** open access https://link.springer.com/content/pdf/10.1007/s12144-022-03573-2.pdf
- **Score:** 1=1 2=0 3=1 4=4 5=4 total=10/25

---

## Sub-theme B: Automation bias and wrong AI advice

### Parasuraman, R., & Manzey, D. H. (2010). Complacency and bias in human use of automation: An attentional integration. *Human Factors*, *52*(3), 381–410. https://doi.org/10.1177/0018720810376055
- **Type:** other HCI (human factors)
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1177/0018720810376055
- **Design:** Review of empirical studies on automation complacency and automation bias, with an integrated model.
- **Finding:** Automation bias produces both omission errors (missing what the aid did not flag) and commission errors (following a wrong flag) when decision aids are imperfect. It occurs in naive and expert participants, is not prevented by training or instructions, and affects teams as well as individuals. Complacency appears under multiple-task load. The two are overlapping effects driven by attention.
- **Relation:** Definition of automation bias; H1 (a wrong flag is followed) and H4 (checker participants act on the flag by questioning). Our participants are not operators supervising automation: the automation's target is another person, and the commission error is a loss of trust in that person.
- **Access:** open access https://doi.org/10.14279/depositonce-8052
- **Score:** 1=0 2=2 3=3 4=5 5=4 total=14/25

### Goddard, K., Roudsari, A., & Wyatt, J. C. (2012). Automation bias: A systematic review of frequency, effect mediators, and mitigators. *Journal of the American Medical Informatics Association*, *19*(1), 121–127. https://doi.org/10.1136/amiajnl-2011-000089
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1136/amiajnl-2011-000089
- **Design:** Systematic review, healthcare focus; 74 of 13,821 retrieved papers met inclusion criteria.
- **Finding:** Automation bias is mediated by user factors (cognitive style, task-specific experience), by trust and confidence, and by environment (workload, task complexity, time pressure). Mitigators included training, emphasising user accountability, where the advice sits on the screen, updated confidence levels on the output, and giving information rather than a recommendation.
- **Relation:** Definition and design element: time pressure and low task experience (board members in a 10-minute briefing) are named mediators; our flags are worded as verdicts ("Architecture likely wrong"), the recommendation form the review links to more bias. Our study measures the social cost of the bias for the human expert, not diagnostic error.
- **Access:** open access https://academic.oup.com/jamia/article-pdf/19/1/121/5911703/19-1-121.pdf
- **Score:** 1=0 2=2 3=4 4=4 5=4 total=14/25

### Dratsch, T., Chen, X., Rezazade Mehrizi, M., Kloeckner, R., Mähringer-Kunz, A., Püsken, M., Baeßler, B., Sauer, S., Maintz, D., & Pinto dos Santos, D. (2023). Automation bias in mammography: The impact of artificial intelligence BI-RADS suggestions on reader performance. *Radiology*, *307*(4), Article e222176. https://doi.org/10.1148/radiol.222176
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1148/radiol.222176
- **Design:** 27 radiologists (inexperienced, moderately experienced, very experienced) read 50 mammograms with a purported AI system suggesting a BI-RADS category; in the 40-case test set the AI suggested a wrong category for 12. Measured: share of mammograms rated correctly, degree of bias, perceived AI accuracy, reader confidence.
- **Finding:** When the AI was right versus wrong, the share of correctly rated mammograms fell from 79.7% to 19.8% for inexperienced readers, from 81.3% to 24.8% for moderately experienced, and from 82.3% to 45.5% for very experienced readers. Inexperienced readers followed a wrongly higher AI category more than very experienced readers (mean bias 4.0 vs 1.2).
- **Relation:** Automation bias; H1 and H4: a wrong flag pulls judgement even among experts, and most among the least experienced, which is our board member's position. Our study asks what the wrong flag does to trust in the human whose work was flagged, not to the reader's accuracy.
- **Access:** open access https://research.vu.nl/files/291201748/Dratsch_et_al._2023_-_Automation_Bias_in_Mammography_-_The_Impact_of_Artificial_Intelligence_BI-RADS_Suggestions_on_Reader_Performance.pdf
- **Score:** 1=0 2=3 3=4 4=4 5=5 total=16/25

### Gaube, S., Suresh, H., Raue, M., Merritt, A., Berkowitz, S. J., Lermer, E., Coughlin, J. F., Guttag, J. V., Colak, E., & Ghassemi, M. (2021). Do as AI say: Susceptibility in deployment of clinical decision-aids. *npj Digital Medicine*, *4*(1), Article 31. https://doi.org/10.1038/s41746-021-00385-9
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1038/s41746-021-00385-9
- **Design:** Physicians (radiologists and physicians with less task expertise) received chest X-rays with diagnostic advice, some of it inaccurate. All advice was written by human experts; the label (AI system vs human expert) was manipulated. Measured: rated advice quality and diagnostic accuracy.
- **Finding:** Radiologists rated advice as lower quality when it was labelled as from AI; physicians with less task expertise did not. Diagnostic accuracy was significantly worse after inaccurate advice whatever the stated source.
- **Relation:** Advice source labelling; H1 and H4: non-experts do not discount AI-labelled advice the way domain experts do, and wrong advice harms judgement regardless of label. Our board member is the non-expert in this pattern; we measure trust in the person contradicted, not advice quality.
- **Access:** open access https://www.nature.com/articles/s41746-021-00385-9.pdf
- **Score:** 1=1 2=3 3=4 4=4 5=4 total=16/25

### Klingbeil, A., Grützner, C., & Schreck, P. (2024). Trust and reliance on AI: An experimental study on the extent and costs of overreliance on AI. *Computers in Human Behavior*, *160*, Article 108352. https://doi.org/10.1016/j.chb.2024.108352
- **Type:** other HCI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1016/j.chb.2024.108352
- **Design:** Incentivized, interactive, domain-independent behavioural experiment in which participants received advice in an uncertain situation; the advice source was labelled as AI or not. Measured: whether participants followed advice that contradicted the contextual information available to them and their own assessment, and the costs to themselves and to third parties.
- **Finding:** Merely knowing the advice came from an AI made people over-rely on it: they followed AI advice even when it contradicted the available contextual information and their own assessment, often to their own cost and with undesired effects on third parties. (Sample size and effect sizes are in the full text, not the abstract.)
- **Relation:** Automation bias against context; H1 and H4: the closest finding to our "right by the textbook, wrong in context" flags. The third party who pays in our study is the CISO, whose trust is measured.
- **Access:** open access https://doi.org/10.1016/j.chb.2024.108352
- **Score:** 1=0 2=3 3=2 4=5 5=4 total=14/25

### De-Arteaga, M., Fogliato, R., & Chouldechova, A. (2020). A case for humans-in-the-loop: Decisions in the presence of erroneous algorithmic scores. In *Proceedings of the 2020 CHI Conference on Human Factors in Computing Systems* (pp. 1–12). Association for Computing Machinery. https://doi.org/10.1145/3313831.3376638
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1145/3313831.3376638
- **Design:** Field data from a deployed algorithmic risk tool for child-maltreatment hotline screening, including a period in which the displayed score was an incorrect estimate of risk. Measured: call screeners' adherence to the tool's recommendation.
- **Finding:** Screeners changed their decisions when the tool was deployed, and were less likely to follow its recommendation when the displayed score was wrong, even when overriding it required supervisory approval.
- **Relation:** Counter-result to automation bias: professionals with domain context can detect and override wrong scores. Supports H2 (a reasoned account can win against the flag). Our judges are not domain experts; the person with the context is the CISO, so the override must come through the CISO's answer.
- **Access:** open access https://dl.acm.org/doi/pdf/10.1145/3313831.3376638
- **Score:** 1=0 2=3 3=4 4=4 5=4 total=15/25

### Holstein, K., De-Arteaga, M., Tumati, L., & Cheng, Y. (2023). Toward supporting perceptual complementarity in human-AI collaboration via reflection on unobservables. *Proceedings of the ACM on Human-Computer Interaction*, *7*(CSCW1), 1–20. https://doi.org/10.1145/3579628
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1145/3579628
- **Design:** Online experiment in which participants made predictions with model outputs; some conditions prompted them about "unobservables", features that may influence the outcome but are unavailable to the model. Measured: how participants combined model output and unobservables, and their accuracy.
- **Finding:** Prompts about unobservables changed how people integrated model outputs with the unobservables, but did not necessarily improve performance; effects varied with prior domain expertise.
- **Relation:** Design element and gap: our checker's flags are wrong precisely because of unobservables (the company's context only the CISO knows). The CISO's reasoned answer is in effect a disclosure of unobservables; this paper shows that telling people about them does not guarantee they use them well. We test this socially, with the unobservable voiced by the flagged human.
- **Access:** open access https://dl.acm.org/doi/pdf/10.1145/3579628
- **Score:** 1=0 2=2 3=2 4=4 5=4 total=12/25

---

## Sub-theme C: Over-reliance, explanations, confidence and second opinions (SIGCHI)

### Bansal, G., Wu, T., Zhou, J., Fok, R., Nushi, B., Kamar, E., Ribeiro, M. T., & Weld, D. (2021). Does the whole exceed its parts? The effect of AI explanations on complementary team performance. In *Proceedings of the 2021 CHI Conference on Human Factors in Computing Systems* (pp. 1–16). Association for Computing Machinery. https://doi.org/10.1145/3411764.3445717
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1145/3411764.3445717 and https://arxiv.org/abs/2006.14779
- **Design:** Mixed-method user studies on three datasets; an AI with accuracy comparable to humans helped participants, with or without explanations. Measured: team accuracy and acceptance of AI recommendations.
- **Finding:** AI assistance gave some complementary improvement, but explanations did not add to it. Explanations raised the chance that people accepted the AI's recommendation whether it was right or wrong. (No counts in the abstract.)
- **Relation:** Over-reliance; H1 and H4: a flag with a stated reason ("Numbers do not add up") should be accepted more readily even when wrong. Our checker flags a person in a social setting, not a classification label.
- **Access:** open access https://arxiv.org/pdf/2006.14779
- **Score:** 1=0 2=2 3=1 4=4 5=4 total=11/25

### Buçinca, Z., Malaya, M. B., & Gajos, K. Z. (2021). To trust or to think: Cognitive forcing functions can reduce overreliance on AI in AI-assisted decision-making. *Proceedings of the ACM on Human-Computer Interaction*, *5*(CSCW1), 1–21. https://doi.org/10.1145/3449287
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18 (Crossref lists the short title "To Trust or to Think"; subtitle from the arXiv record); abstract read at https://api.crossref.org/works/10.1145/3449287
- **Design:** Experiment, N = 199, comparing three cognitive forcing designs (making people engage before or while seeing the AI answer) with two simple explainable-AI designs and a no-AI baseline. Measured: over-reliance, subjective ratings, moderation by Need for Cognition.
- **Finding:** Cognitive forcing reduced over-reliance compared with simple explanations, but participants rated the designs that cut over-reliance most as least favourable. The benefit went mostly to participants high in Need for Cognition.
- **Relation:** Over-reliance and dual-process account; H4 design element: our trust rating before the CISO answers and the forced choice (nothing, explain, challenge) work like a cognitive forcing step. Need for Cognition is a candidate covariate. Our study measures trust in the human, not decision accuracy.
- **Access:** open access https://arxiv.org/pdf/2102.09692
- **Score:** 1=0 2=2 3=1 4=4 5=4 total=11/25

### Vasconcelos, H., Jörke, M., Grunde-McLaughlin, M., Gerstenberg, T., Bernstein, M. S., & Krishna, R. (2023). Explanations can reduce overreliance on AI systems during decision-making. *Proceedings of the ACM on Human-Computer Interaction*, *7*(CSCW1), 1–38. https://doi.org/10.1145/3579605
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1145/3579605
- **Design:** Five studies, N = 731, maze task with a simulated AI; manipulated task difficulty, explanation difficulty and monetary reward. Measured: over-reliance (agreeing with the AI when it is wrong).
- **Finding:** People choose strategically whether to check the AI: harder tasks, harder explanations and lower rewards increased over-reliance, and explanations reduced over-reliance when they lowered the cost of verifying the AI's answer.
- **Relation:** Over-reliance as a cost-benefit choice; H4: board members who cannot verify a slide cheaply will lean on the flag; asking the CISO is a low-cost verification route, which our "ask to explain" option offers. We measure trust in the human after verification, not agreement with the AI.
- **Access:** open access https://dl.acm.org/doi/pdf/10.1145/3579605
- **Score:** 1=0 2=2 3=1 4=3 5=4 total=10/25

### Zhang, Y., Liao, Q. V., & Bellamy, R. K. E. (2020). Effect of confidence and explanation on accuracy and trust calibration in AI-assisted decision making. In *Proceedings of the 2020 Conference on Fairness, Accountability, and Transparency* (pp. 295–305). Association for Computing Machinery. https://doi.org/10.1145/3351095.3372852
- **Type:** SIGCHI (FAT*/FAccT)
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1145/3351095.3372852
- **Design:** Two human experiments in which humans and AI had comparable accuracy alone; manipulated display of a per-case confidence score and a local explanation. Measured: trust calibration (case-by-case reliance) and joint accuracy.
- **Finding:** Showing a confidence score helped people calibrate trust in the AI, but calibration alone did not improve decisions; that also depended on whether the human brought enough unique knowledge to complement the AI's errors. Local explanations did not help.
- **Relation:** Trust calibration definition; design element (our flags carry "likely", a verbal confidence cue). The "unique knowledge" condition maps onto the CISO, not the board member, which is what makes our setting different: the knowledge that corrects the AI sits with the person being flagged.
- **Access:** open access https://arxiv.org/pdf/2001.02114
- **Score:** 1=0 2=2 3=2 4=4 5=4 total=12/25

### Lai, V., Chen, C., Smith-Renner, A., Liao, Q. V., & Tan, C. (2023). Towards a science of human-AI decision making: An overview of design space in empirical human-subject studies. In *Proceedings of the 2023 ACM Conference on Fairness, Accountability, and Transparency* (pp. 1369–1385). Association for Computing Machinery. https://doi.org/10.1145/3593013.3594087
- **Type:** SIGCHI (FAccT)
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1145/3593013.3594087
- **Design:** Survey of over 100 empirical human-subject studies of AI-assisted decision making, coded for decision tasks, AI assistance elements and evaluation metrics.
- **Finding:** Summarises trends and gaps in tasks, assistance elements and metrics, and calls for common frameworks so studies can be compared and generalised.
- **Relation:** Theme anchor and gap: the surveyed designs have a human deciding with AI help; none in the abstract involve the AI assessing another human. Lets us state that our human-AI-human design falls outside the mapped design space.
- **Access:** open access https://dl.acm.org/doi/pdf/10.1145/3593013.3594087
- **Score:** 1=0 2=1 3=2 4=4 5=4 total=11/25

### Lu, Z., Wang, D., & Yin, M. (2024). Does more advice help? The effects of second opinions in AI-assisted decision making. *Proceedings of the ACM on Human-Computer Interaction*, *8*(CSCW1), 1–31. https://doi.org/10.1145/3653708
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1145/3653708 and https://arxiv.org/abs/2401.07058
- **Design:** Three pre-registered randomized experiments; participants received an AI recommendation plus, in some conditions, a second opinion from a peer or another AI, either always shown or requested on demand. Measured: over-reliance and under-reliance on the AI, and performance.
- **Finding:** When the AI's recommendation and a second opinion were always shown together, people relied less on the AI when it was wrong but also more often rejected it when it was right, whether the second opinion came from a peer or another AI. When people chose when to ask a peer, the requests could cut over-reliance without raising under-reliance in some cases. The source of the second opinion had no significant main effect on reliance.
- **Relation:** Second opinions; H2 and H4: the CISO's answer functions as a human second opinion against the checker, and asking for it is the "on demand" condition. Our study differs because the second opinion comes from the person under scrutiny, and we measure trust in that person.
- **Access:** open access https://arxiv.org/pdf/2401.07058
- **Score:** 1=1 2=3 3=1 4=4 5=4 total=13/25

---

## Sub-theme D: Large language models that are confidently wrong

### Kim, S. S. Y., Liao, Q. V., Vorvoreanu, M., Ballard, S., & Vaughan, J. W. (2024). "I'm not sure, but...": Examining the impact of large language models' uncertainty expression on user reliance and trust. In *Proceedings of the 2024 ACM Conference on Fairness, Accountability, and Transparency* (pp. 822–835). Association for Computing Machinery. https://doi.org/10.1145/3630106.3658941
- **Type:** SIGCHI (FAccT)
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1145/3630106.3658941
- **Design:** Pre-registered experiment, N = 404, participants answered medical questions with or without a fictional LLM-infused search engine; manipulated natural-language uncertainty expressions (first person "I'm not sure, but..." vs general "It's not clear, but..." vs none). Measured: reliance (agreement), confidence in the system, trust, accuracy.
- **Finding:** First-person uncertainty lowered participants' confidence in the system and their agreement with its answers, and raised their accuracy, mainly by reducing (not removing) over-reliance on wrong answers. General-perspective uncertainty had weaker, non-significant effects.
- **Relation:** Design element for the checker's wording ("Architecture likely wrong" is a hedged, impersonal verdict). Predicts that a stronger or weaker hedge will change how much the flag moves trust (H1). Our study measures trust in the flagged human.
- **Access:** open access https://doi.org/10.1145/3630106.3658941
- **Score:** 1=0 2=2 3=2 4=4 5=4 total=12/25

### Kim, S. S. Y., Vaughan, J. W., Liao, Q. V., Lombrozo, T., & Russakovsky, O. (2025). Fostering appropriate reliance on large language models: The role of explanations, sources, and inconsistencies. In *Proceedings of the 2025 CHI Conference on Human Factors in Computing Systems* (pp. 1–19). Association for Computing Machinery. https://doi.org/10.1145/3706598.3714020
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1145/3706598.3714020
- **Design:** Think-aloud study, then a pre-registered experiment, N = 308, with an LLM-infused application answering objective questions; manipulated presence of explanations, sources, and inconsistencies within explanations. Measured: reliance, accuracy and other measures.
- **Finding:** Explanations increased reliance on both correct and incorrect LLM answers. Reliance on incorrect answers dropped when sources were shown or when the explanation contained inconsistencies.
- **Relation:** H1 and H2: a flag with a reason attached will be accepted even when wrong; a counter-account that exposes an inconsistency (the CISO naming the company context) is the kind of cue that lowers reliance on the wrong answer. Our flags have no sources, by design.
- **Access:** open access https://doi.org/10.1145/3706598.3714020
- **Score:** 1=0 2=2 3=1 4=4 5=4 total=11/25

### Steyvers, M., Tejeda, H., Kumar, A., Belem, C., Karny, S., Hu, X., Mayer, L. W., & Smyth, P. (2025). What large language models know and what people think they know. *Nature Machine Intelligence*, *7*(2), 221–231. https://doi.org/10.1038/s42256-024-00976-7
- **Type:** other HCI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1038/s42256-024-00976-7
- **Design:** Experiments with multiple-choice and short-answer questions; participants judged LLM answers with default or adjusted explanations. Measured: calibration gap (human confidence in the LLM answer vs the model's own confidence) and discrimination gap (telling correct from incorrect answers).
- **Finding:** Users overestimated the accuracy of LLM answers when given default explanations, and longer explanations raised user confidence even when the extra length did not improve accuracy. Adjusting explanations to reflect the model's internal confidence narrowed both gaps.
- **Relation:** Calibration; H1 and H4: people over-trust fluent model output, which supports a drop in trust in the CISO after a flag. Our checker gives terse verdicts, which this work suggests may be trusted less than long ones.
- **Access:** open access https://www.nature.com/articles/s42256-024-00976-7.pdf
- **Score:** 1=0 2=2 3=1 4=4 5=4 total=11/25

---

## Sub-theme E: Trust in automation and in AI: definitions, models and measurement

### Lee, J. D., & See, K. A. (2004). Trust in automation: Designing for appropriate reliance. *Human Factors*, *46*(1), 50–80. https://doi.org/10.1518/hfes.46.1.50_30392
- **Type:** other HCI (human factors)
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.1518/hfes.46.1.50_30392
- **Design:** Conceptual review drawing on organizational, sociological, interpersonal, psychological and neurological views of trust.
- **Finding:** Because people respond to technology socially, trust guides reliance on automation, especially when complexity makes full understanding impractical. Trust forms through analytic, analogical and affective processes, depends on context and display, and must be calibrated to the automation's capability. Discusses the limits of carrying trust-in-people concepts over to automation.
- **Relation:** Definition of appropriate reliance and calibration. Our study runs the transfer the other way: trust in automation shaping trust in a person.
- **Access:** SFU library
- **Score:** 1=0 2=1 3=2 4=5 5=4 total=12/25

### Glikson, E., & Woolley, A. W. (2020). Human trust in artificial intelligence: Review of empirical research. *Academy of Management Annals*, *14*(2), 627–660. https://doi.org/10.5465/annals.2018.0057
- **Type:** management or economics
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.5465/annals.2018.0057
- **Design:** Review of 20 years of empirical research across disciplines on the determinants of human trust in AI.
- **Finding:** Form of AI (robot, virtual, embedded) and level of machine intelligence shape trust. Tangibility, transparency, reliability and immediacy drive cognitive trust; anthropomorphism drives emotional trust. The evidence base relies on short-term, small, experimental studies with varied trust measures, unlike longer, higher-stakes field settings.
- **Relation:** Definition and theme anchor for trust in the checker (embedded AI, cognitive trust); management venue suits the board framing. Gap: the review covers trust in AI, not how AI shapes trust between people.
- **Access:** SFU library
- **Score:** 1=0 2=1 3=2 4=5 5=4 total=12/25

### Vereschak, O., Bailly, G., & Caramiaux, B. (2021). How to evaluate trust in AI-assisted decision making? A survey of empirical methodologies. *Proceedings of the ACM on Human-Computer Interaction*, *5*(CSCW2), 1–39. https://doi.org/10.1145/3476068
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18; abstract read at https://api.crossref.org/works/10.1145/3476068
- **Design:** Survey of empirical methods for studying trust in AI-assisted decision making, analysed by the elements of an experimental protocol.
- **Finding:** Most protocols do not build a definition of trust into their design, which leads to findings that are overclaimed or hard to compare. Gives practical guidelines drawn from human-human trust research.
- **Relation:** Measure: guides our trust measurement (state the definition, separate trust from reliance, measure before and after). Supports measuring trust in the CISO as competence, benevolence and integrity rather than a single item alone.
- **Access:** open access https://hal.sorbonne-universite.fr/hal-03280969
- **Score:** 1=1 2=1 3=1 4=5 5=4 total=12/25

---

## Lower-priority entries (verified, abstract read, not scored in full)

These are verified by Crossref on 2026-10-18 with abstracts read (OpenAlex, Crossref or ACL Anthology). They duplicate a finding above or sit further from our design. Use if a subsection needs another citation.

- Schemmer, M., Kuehl, N., Benz, C., Bartos, A., & Satzger, G. (2023). Appropriate reliance on AI advice: Conceptualization and the effect of explanations. In *Proceedings of the 28th International Conference on Intelligent User Interfaces* (pp. 410–422). ACM. https://doi.org/10.1145/3581641.3584066. Defines Appropriateness of Reliance as a two-dimensional measure (relying on correct, rejecting incorrect advice); experiment N = 200 on explanations. Measure only. Score about 10/25.
- Lai, V., & Tan, C. (2019). On human predictions with explanations and predictions of machine learning models: A case study on deception detection. In *Proceedings of the Conference on Fairness, Accountability, and Transparency* (pp. 29–38). ACM. https://doi.org/10.1145/3287560.3287590. Showing predicted labels raised human accuracy by more than 20% (relative); stating the machine's strong performance raised it further; explanations alone helped slightly. Abstract via arXiv 1811.07901. Score about 9/25.
- Lu, Z., & Yin, M. (2021). Human reliance on machine learning models when performance feedback is limited: Heuristics and risks. In *Proceedings of the 2021 CHI Conference on Human Factors in Computing Systems* (pp. 1–16). ACM. https://doi.org/10.1145/3411764.3445562. Without performance information, people judge a model by how often it agrees with them on cases they are confident about. Relevant to how board members judge the checker on slides 1 and 5 ("No issues found"). Score about 11/25.
- Si, C., Goyal, N., Wu, T., Zhao, C., Feng, S., Daumé III, H., & Boyd-Graber, J. (2024). Large language models help humans verify truthfulness, except when they are convincingly wrong. In *Proceedings of NAACL-HLT 2024* (Vol. 1, pp. 1459–1474). Association for Computational Linguistics. https://doi.org/10.18653/v1/2024.naacl-long.81. 80 crowdworkers fact-checked claims; LLM explanations made them faster with similar accuracy to search, but they over-relied on wrong explanations; contrastive explanations reduced this. Score about 11/25.
- Hoff, K. A., & Bashir, M. (2015). Trust in automation: Integrating empirical evidence on factors that influence trust. *Human Factors*, *57*(3), 407–434. https://doi.org/10.1177/0018720814547570. Systematic review of 127 studies (2002 to 2013); three layers of trust: dispositional, situational, learned. Score about 11/25.
- Jussupow, E., Spohrer, K., Heinzl, A., & Gawlitza, J. (2021). Augmenting medical diagnosis decisions? An investigation into physicians' decision-making process with artificial intelligence. *Information Systems Research*, *32*(3), 713–735. https://doi.org/10.1287/isre.2020.0980. 68 novice and 12 experienced physicians made more wrong diagnoses with incorrect AI advice than with no advice. Score about 12/25.
- Cecil, J., Lermer, E., Hudecek, M. F. C., Sauer, J., & Gaube, S. (2024). Explainability does not mitigate the negative impact of incorrect AI advice in a personnel selection task. *Scientific Reports*, *14*, Article 9736. https://doi.org/10.1038/s41598-024-60220-5. Five pre-registered experiments, N = 1,403 students and HR employees; incorrect advice harmed performance, source (human vs AI) and explainability had limited effects. Score about 12/25.
- Alon-Barkat, S., & Busuioc, M. (2023). Human-AI interactions in public sector decision making: "Automation bias" and "selective adherence" to algorithmic advice. *Journal of Public Administration Research and Theory*, *33*(1), 153–169. https://doi.org/10.1093/jopart/muac007. Three experiments (N = 605, 904, 1,345 civil servants): no evidence that algorithmic advice was followed more than equivalent human-expert advice; adherence rose when advice matched stereotypes. Counter-result for H1, governance context. Score about 12/25.
- You, S., Yang, C. L., & Li, X. (2022). Algorithmic versus human advice: Does presenting prediction performance matter for algorithm appreciation? *Journal of Management Information Systems*, *39*(2), 336–365. https://doi.org/10.1080/07421222.2022.2063553. Five JAS experiments: people followed algorithmic advice more than identical human advice because they trusted the algorithmic advisor more; aggregate error information did not reduce this. Score about 11/25.
- Burton, J. W., Stein, M.-K., & Jensen, T. B. (2020). A systematic review of algorithm aversion in augmented decision making. *Journal of Behavioral Decision Making*, *33*(2), 220–239. https://doi.org/10.1002/bdm.2155. Review of 61 articles; five themes of causes (expectations and expertise, decision autonomy, incentives, cognitive compatibility, divergent rationalities). Score about 9/25.
- Pálfi, B., Arora, K., & Kostopoulou, O. (2022). Algorithm-based advice taking and clinical judgement: Impact of advice distance and algorithm information. *Cognitive Research: Principles and Implications*, *7*, Article 70. https://doi.org/10.1186/s41235-022-00421-6. GPs weighed their own estimate and a cancer risk algorithm equally on average, kept their estimate 29% of the time and fully updated 27% of the time; slight discounting when advice was far from their own estimate. Score about 9/25.
- Chong, L., Zhang, G., Goucher-Lambert, K., Kotovsky, K., & Cagan, J. (2022). Human confidence in artificial intelligence and in themselves: The evolution and impact of confidence on adoption of AI advice. *Computers in Human Behavior*, *127*, Article 107018. https://doi.org/10.1016/j.chb.2021.107018. Self-confidence, not confidence in the AI, drove acceptance of AI suggestions; people blamed themselves for AI errors and kept relying on a poor AI. Score about 10/25.
- Bogert, E., Schecter, A., & Watson, R. T. (2021). Humans rely more on algorithms than social influence as a task becomes more difficult. *Scientific Reports*, *11*, Article 8028. https://doi.org/10.1038/s41598-021-87480-9. Three pre-registered experiments: reliance on algorithmic advice rose with task difficulty, relative to peer advice; inaccurate algorithmic advice was disregarded more strongly than equally inaccurate peer advice. Score about 10/25.

---

## Notes on known papers

- Chen, Sun, Liao and Sundar (2026, IJHCS, known) is the patient-side counterpart of Lu, Wang and Yin (2024, CSCW, above): Lu et al. show that a second opinion shifts reliance away from the AI; Chen et al. show that AI disagreement with a doctor raises perceived medical uncertainty and lowers trust in the doctor. Together they frame H1 and H2 from both directions. No new reading of Chen et al. itself.
- Logg, Minson and Moore (2019, known): the secondary description seen during search notes that experienced professionals relied less on algorithmic advice than lay people did. Relevant only in contrast: our board members are lay with respect to security, so the appreciation side applies.

## Not verified or not included

- Dzindolet, M. T., Peterson, S. A., Pomranky, R. A., Pierce, L. G., & Beck, H. P. (2003). The role of trust in automation reliance. *International Journal of Human-Computer Studies*, *58*(6), 697–718. https://doi.org/10.1016/S1071-5819(03)00038-7. Crossref metadata verified; abstract not obtainable (ScienceDirect 403, no abstract in OpenAlex or Semantic Scholar). Secondary sources say that telling operators why the automation might err maintained trust and use. Not scored.
- Skitka, L. J., Mosier, K. L., & Burdick, M. (1999). Does automation bias decision-making? *International Journal of Human-Computer Studies*, *51*(5), 991–1006. https://doi.org/10.1006/ijhc.1999.0252. Crossref metadata verified; abstract not read (only a Semantic Scholar machine summary). Classic source of the term "automation bias"; read before citing.
- Tschandl, P., et al. (2020). Human-computer collaboration for skin cancer recognition. *Nature Medicine*, *26*(8), 1229–1234. https://doi.org/10.1038/s41591-020-0942-0. Verified, abstract read on PubMed (PMID 32572267): faulty AI misled clinicians at all levels including experts. Left out because Dratsch et al. (2023) gives the same point with numbers.
- Tatlidil et al. (2020), SJDM poster on human vs algorithm advisors giving conflicting hiring advice (https://dlab.sauder.ubc.ca/sjdm/presentations/2020-Poster-Tatlidil-Kemal_Semir-Algorithms-Advice-Trust.pdf): poster only, no DOI, not peer reviewed. Reports that agreement with an advisor predicted trust in that advisor, for both human and algorithm.

## Searches run

- Crossref lookups by DOI for all entries in the brief's Agent 2 section and for candidates from memory (2026-10-18); Crossref bibliographic title queries for Lu, Wang and Yin (2024), Bailey et al. (2022), Pálfi et al. (2022), You et al. (2022).
- OpenAlex and Semantic Scholar abstract lookups by DOI; PubMed efetch for Tschandl et al.; arXiv and ACL Anthology pages for Lai and Tan, Bansal et al., Si et al.
- Web: "experiment perceived competence of human expert advisor when algorithm disagrees with advisor judge-advisor"
- Web: "conflicting advice" human advisor and algorithm advisor disagree trust experiment 2023 OR 2024 OR 2025
- Web: AI second opinion contradicts human expert observers trust in the expert decreases experiment financial advisor OR doctor OR judge
- Web: "algorithmic advice" "human advice" conflict OR disagreement "weight of advice" judge-advisor 2023 2024 journal
- Web: Bonaccio Dalal 2006 abstract judge-advisor system; Dzindolet 2003 abstract
- Gap observed: no search returned a study that measures trust in a *human* advisor after an algorithm disagrees with that advisor in front of a third-party judge, other than the known Chen et al. (2026). The judge-advisor and human-AI decision literatures measure reliance on advice, not the advisor's standing.
