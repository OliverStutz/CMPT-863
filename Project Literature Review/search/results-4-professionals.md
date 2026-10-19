# Results, Agent 4: Evaluation of professionals who use, follow, defy, or are contradicted by AI

Written 2026-10-18 by Agent 4. 18 entries in four sub-themes. Every DOI was checked against api.crossref.org on 2026-10-18 (the arXiv DOI is a DataCite DOI and was checked on OpenAlex and arXiv instead). Abstracts were read on Europe PMC, PubMed Central, OpenAlex (which holds the publisher abstract) or the publisher page, as listed per entry. Known papers from `_known-papers.md` are not repeated.

Short map of what this file adds for our review:

- **AI contradicts the expert, lay observers judge the expert** (closest to our checker flag): Madanay et al. 2025, Bernstein et al. 2025, Song et al. 2026, Pezzo & Pezzo 2006, Pezzo et al. 2022, Nurek & Kostopoulou 2023, Tobia et al. 2021, Kempt & Nagel 2022.
- **Simply using AI costs the professional trust** (a baseline penalty that our checker group does not test, because there the board member, not the CISO, uses the tool): Reis et al. 2024, Reis et al. 2025, Yang et al. 2025, Cavalier et al. 2025, Schilke & Reimann 2025, Fine et al. 2025, Zhang et al. 2026, and a counter-result, Calluso et al. 2026.
- **People flagged by an AI detector**: Du et al. 2026 (a wrong detector report lowers teachers' ratings of an identical paper) and Chu et al. 2024.

---

## Sub-theme A: An AI disagrees with the professional, and lay observers judge the professional

### Madanay, F. L., O'Donohue, L. S., & Zikmund-Fisher, B. J. (2025). Patient reactions to artificial intelligence–clinician discrepancies: Web-based randomized experiment. *Journal of Medical Internet Research*, *27*, Article e68823. https://doi.org/10.2196/68823
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://www.ebi.ac.uk/europepmc/ (PMC12141964)
- **Design:** 1,606 US adults imagined a lung-cancer screening CT and were randomized to results from (1) a radiologist only, (2) AI and radiologist in agreement, (3) a radiologist who recommended more testing than the AI, or (4) a radiologist who recommended less testing than the AI. Measured on the radiologist: agreement with the recommendation, likelihood of recommending the radiologist to family and friends, and how good a provider the radiologist is. Medical maximizing-minimizing preference measured as moderator.
- **Finding:** When the radiologist recommended less testing than the AI, agreement with the radiologist fell to a mean of 4.01, against 4.55 to 4.63 in the other three conditions, which did not differ from each other. Participants were also least likely to recommend this radiologist and rated them as the worst provider. A radiologist who recommended more testing than the AI paid no penalty. Patients who prefer aggressive care sided with the radiologist who overcalled the AI and against the one who undercalled it.
- **Relation:** AI contradicts the expert; supports H1 (a disagreeing AI lowers lay evaluation of the expert) and shows the penalty is asymmetric: it depends on whether the AI looks more cautious than the expert. Our checker's flags ("numbers do not add up") play the cautious role, so this predicts a drop. Our study differs in that the observer holds the checker live, the expert can answer, and trust is measured before and after the answer.
- **Access:** open access https://pmc.ncbi.nlm.nih.gov/articles/PMC12141964/
- **Score:** 1=4 2=5 3=5 4=5 5=5 total=24/25

### Pezzo, M. V., & Pezzo, S. P. (2006). Physician evaluation after medical errors: Does having a computer decision aid help or hurt in hindsight? *Medical Decision Making*, *26*(1), 48–56. https://doi.org/10.1177/0272989X05282644
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/https://doi.org/10.1177/0272989x05282644 (publisher abstract)
- **Design:** Experiment 1: 59 undergraduates read about a doctor whose diagnosis was right or wrong, with or without a decision aid; rated decision quality, likelihood of recommending the doctor, and after a bad outcome negligence and likelihood of suing. Experiment 2: 166 medical students and 154 undergraduates read bad-outcome scenarios in which the doctor agreed with the aid, followed the aid against his own opinion, defied the aid in favour of his own opinion, or used no aid; rated fault and competence.
- **Finding:** Using an aid made a good outcome look less good and a bad outcome look less bad. After a bad outcome, a doctor who agreed with or followed the aid was blamed less, while a doctor who defied the aid was blamed about as much as one who used no aid. Medical students judged more mildly than undergraduates but correctly predicted the lay judgement.
- **Relation:** follows or defies the aid; context for H1 and H3: an expert who overrides a tool loses the tool's protection. In our study the CISO is right and the checker wrong, and the CISO's answer, not the outcome, is what the observer sees; this paper judges after a known bad outcome.
- **Access:** SFU library (author copy listed at https://digitalcommons.usf.edu/fac_publications/1224)
- **Score:** 1=4 2=5 3=4 4=5 5=4 total=22/25

### Pezzo, M. V., Nash, B. E. D., Vieux, P., & Foster-Grammer, H. W. (2022). Effect of having, but not consulting, a computerized diagnostic aid. *Medical Decision Making*, *42*(1), 94–104. https://doi.org/10.1177/0272989X211011160
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/https://doi.org/10.1177/0272989X211011160
- **Design:** Experiment 1: participants read about a diagnosis by a physician or an auto mechanic; half were told a computerized diagnostic aid was available but not consulted. Experiment 2 (bad outcome only): two stated reasons for not consulting the aid, plus accuracy information showing the aid beat the physician. Competence and negligence ratings.
- **Finding:** Not consulting an available aid did not lower the physician's competence ratings. Ratings were lower when the physician said he did not trust the aid and, unexpectedly, higher when he said he already knew what the aid would say. Consistent with the 2006 paper, ratings were high when the physician consulted and followed the aid, and low when he consulted it and then ignored it.
- **Relation:** follows or defies the aid; a design element for H2 and H3: the stated reason for setting the tool aside changes the verdict ("I knew what it would say" helps, "I don't trust it" hurts). Our CISO's reasoned answer versus bare denial is a close analogue; our study adds a live observer who holds the tool.
- **Access:** open access https://digital.stpetersburg.usf.edu/fac_publications/4109
- **Score:** 1=4 2=4 3=4 4=5 5=4 total=21/25

### Bernstein, M. H., Sheppard, B., Bruno, M. A., Lay, P. S., & Baird, G. L. (2025). Randomized study of the impact of AI on perceived legal liability for radiologists. *NEJM AI*, *2*(6). https://doi.org/10.1056/AIoa2400785
- **Type:** medical
- **Verified:** Crossref 2026-10-18 (vol 2, issue 6, published 2025-05-22; no article number in Crossref); the publisher abstract was blocked (HTTP 403), so the design and numbers were read in coverage at https://www.2minutemedicine.com/juror-perception-of-radiologist-liability-can-be-affected-by-artificial-intelligence-ai-use-in-diagnosis and https://medical.brown.edu/news/2025-07-28/radiology-artificial-intelligence-malpractice-study
- **Design:** 1,334 US adults eligible for jury duty read two malpractice cases (missed brain bleed, missed lung cancer) and were randomized to one of five conditions: no AI; AI agreed with the radiologist (both missed it); AI disagreed (AI found the abnormality the radiologist missed); AI disagreed plus a stated AI false-discovery rate of 50%; AI agreed plus a stated false-omission rate of 1%. Outcome: did the radiologist meet the duty of care (yes or no).
- **Finding:** Brain bleed: 72.9% found the radiologist liable when the AI had flagged the bleed, against 50.0% when the AI had also missed it (about 56% with no AI). Lung cancer: 78.7% against 63.5%. Giving jurors the AI's error rates reduced liability findings. A follow-up preprint by the same group (https://doi.org/10.1101/2025.06.14.25329278, 282 adults) found that liability after disagreeing with a correct AI fell from 74.7% to 52.9% when the radiologist had read the scan once before seeing the AI.
- **Relation:** AI contradicts the expert; H1 (an AI flag against the expert lowers lay judgement) and a design lever for our materials: telling observers the checker's error rate weakens the flag. Here the AI was right; in our study every flag is wrong, so we test whether the penalty survives when the expert explains why.
- **Access:** SFU library
- **Score:** 1=3 2=5 3=4 4=5 5=3 total=20/25

### Song, E. C., Bernstein, M. H., Sheppard, B., Bruno, M. A., & Baird, G. L. (2026). *An experimental investigation of the relationship between AI-human workflow design and legal liability for radiologists: The erroneous-change penalty and omission bias* [Preprint]. medRxiv. https://doi.org/10.64898/2026.05.20.26353717
- **Type:** medical (preprint, not peer reviewed)
- **Verified:** Crossref 2026-10-18 (posted 2026-05-22); abstract read at https://www.medrxiv.org/content/10.64898/2026.05.20.26353717
- **Design:** 2,347 US adults acting as jurors in a missed-brain-bleed case, randomized to a no-AI control or one of 21 AI conditions varying single or double read, whether the radiologist's first reading was documented, whether the radiologist changed the reading after seeing the AI, whether the AI found the bleed, and whether AI error rates were shown. Outcome: duty of care met or not.
- **Finding:** Liability was highest when the radiologist first saw the bleed and then changed to "normal" after the AI called the scan normal (an "erroneous-change penalty"). Liability was lowest with documented double reads. Showing jurors the AI error rates lowered liability, and liability was lower when the AI was wrong.
- **Relation:** follows a wrong AI; counter-side of H4: an expert who abandons a correct judgement to agree with a wrong tool is judged most harshly. Our CISO keeps the correct position; this preprint suggests observers reward that, if they learn the tool was wrong.
- **Access:** open access https://www.medrxiv.org/content/10.64898/2026.05.20.26353717v1.full.pdf
- **Score:** 1=2 2=5 3=4 4=4 5=4 total=19/25

### Nurek, M., & Kostopoulou, O. (2023). How the UK public views the use of diagnostic decision aids by physicians: A vignette-based experiment. *Journal of the American Medical Informatics Association*, *30*(5), 888–898. https://doi.org/10.1093/jamia/ocad019
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/https://doi.org/10.1093/jamia/ocad019
- **Design:** 730 UK adults imagined a consultation in which a computerized diagnostic aid recommended a test to rule out serious disease. Manipulated: invasiveness of the test, whether the doctor followed the aid's advice, severity of the disease. Measured before and after severity was revealed: satisfaction with the consultation, likelihood of recommending the doctor, preferred frequency of aid use.
- **Finding:** Satisfaction and likelihood of recommending the doctor were higher when the doctor followed the aid, at both time points, and this effect was stronger when participants were worried and when the disease turned out to be serious. Most wanted doctors to use such aids "frequently" (43%) or "always" (17 to 21%).
- **Relation:** follows or defies the aid; counter to a pure use penalty: the public rewards following a cautious aid. Supports the asymmetry in Madanay et al. Our checker flags are not about a test to order but about the expert's content, and the expert is right.
- **Access:** open access https://academic.oup.com/jamia/article/30/5/888/7084009
- **Score:** 1=3 2=5 3=4 4=4 5=5 total=21/25

### Tobia, K., Nielsen, A., & Stremitzer, A. (2021). When does physician use of AI increase liability? *Journal of Nuclear Medicine*, *62*(1), 17–21. https://doi.org/10.2967/jnumed.120.256032
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://jnm.snmjournals.org/content/62/1/17
- **Design:** Nationally representative sample of 2,000 US adults; 2 x 2 vignettes: the AI recommended standard or nonstandard care, and the physician accepted or rejected it; the decision then caused harm. Outcome: agreement that most physicians and a reasonable physician could have made the decision (higher means less liability).
- **Finding:** When the AI recommended standard care, a physician who accepted it was judged less liable than one who rejected it. When the AI recommended nonstandard care, rejecting it and giving standard care gave no comparable protection.
- **Relation:** follows or defies the AI; context for H1 and H3: lay judges hold the expert to the tool's view when the tool sides with convention. Our flags are "right by the textbook", the same pattern as standard-care advice, so observers may side with the checker. Our study measures trust, not liability, and lets the expert give a reason.
- **Access:** open access https://jnm.snmjournals.org/content/jnumed/62/1/17.full.pdf
- **Score:** 1=2 2=5 3=4 4=4 5=5 total=20/25

### Kempt, H., & Nagel, S. K. (2022). Responsibility, second opinions and peer-disagreement: Ethical and epistemological challenges of using AI in clinical diagnostic contexts. *Journal of Medical Ethics*, *48*(4), 222–229. https://doi.org/10.1136/medethics-2021-107440
- **Type:** medical (ethics, conceptual)
- **Verified:** Crossref 2026-10-18; abstract read at https://jme.bmj.com/content/48/4/222
- **Design:** Conceptual paper. Classifies kinds of second opinions, argues several AI systems are accurate enough to act as one, and proposes a "rule of disagreement".
- **Finding:** Treats AI second opinions as peer opinions; when the AI disagrees with the physician in charge, the physician bears an increased burden to justify the decision and a human second opinion should settle the disagreement. Replies in the same issue (Jongsma & Sand; Lang) dispute whether the burden of proof is really asymmetric.
- **Relation:** definition and gap: names the mechanism behind H2 and H3, a contradicting tool shifts the burden of proof onto the expert, so a reasoned answer discharges it and a bare denial does not. Our study tests this claim empirically with a lay observer.
- **Access:** open access https://jme.bmj.com/content/medethics/48/4/222.full.pdf
- **Score:** 1=1 2=5 3=4 4=4 5=4 total=18/25

---

## Sub-theme B: Disclosed AI use lowers trust in the professional

### Reis, M., Reis, F., & Kunde, W. (2025). Public perception of physicians who use artificial intelligence. *JAMA Network Open*, *8*(7), Article e2521643. https://doi.org/10.1001/jamanetworkopen.2025.21643
- **Type:** medical
- **Verified:** Crossref 2026-10-18; full text read at https://www.ebi.ac.uk/europepmc/ (PMC12272287)
- **Design:** 1,276 US adults (January 2025) saw a simulated advertisement for a family doctor that mentioned no AI, or AI used for administrative, diagnostic or therapeutic work. Rated on 1 to 5 scales: competence, trustworthiness, empathy, willingness to book an appointment.
- **Finding:** Every AI mention lowered every rating. Trustworthiness fell from 3.88 (no AI) to 3.61 to 3.66; competence from 3.85 to 3.58 to 3.71; empathy from 4.00 to 3.72 to 3.82; willingness to book from 3.61 to 3.15 to 3.32. The three kinds of AI use did not differ from each other, so even administrative use carried the penalty. Effects were small (d about 0.16 to 0.41).
- **Relation:** use penalty; measure (competence, trustworthiness and empathy as separate ratings, close to ability, integrity and benevolence). Note: the brief lists this as "Zondag et al. 2025"; the JAMA Network Open paper on this topic is by Reis, Reis and Kunde. In our study the expert does not use AI; the observer does, which inverts this setting.
- **Access:** open access https://jamanetwork.com/journals/jamanetworkopen/fullarticle/10.1001/jamanetworkopen.2025.21643
- **Score:** 1=5 2=4 3=4 4=4 5=5 total=22/25

### Reis, M., Reis, F., & Kunde, W. (2024). Influence of believed AI involvement on the perception of digital medical advice. *Nature Medicine*, *30*(11), 3098–3100. https://doi.org/10.1038/s41591-024-03180-7
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://www.ncbi.nlm.nih.gov/pmc/articles/PMC11564086/
- **Design:** Two preregistered studies, 2,280 participants, scenarios of patients obtaining medical advice. The advice text was identical; only its stated source changed ("AI", "human physician", "human + AI"). Rated reliability, empathy and willingness to follow the advice.
- **Finding:** Advice labelled "AI" or "human + AI" was rated less reliable and less empathetic than identical advice labelled "human physician", and participants were less willing to follow it, even when a physician was said to supervise the AI.
- **Relation:** use penalty; context: a label alone, with identical content, shifts judgement of the expert's advice, which supports using identical CISO answers across groups. Our study holds the CISO's content constant and varies whether the observer has a checker.
- **Access:** open access https://www.ncbi.nlm.nih.gov/pmc/articles/PMC11564086/
- **Score:** 1=4 2=4 3=4 4=4 5=4 total=20/25

### Yang, H., Dai, T., Mathioudakis, N., Knight, A. M., Nakayasu, Y., & Wolf, R. M. (2025). Peer perceptions of clinicians using generative AI in medical decision-making. *npj Digital Medicine*, *8*, Article 530. https://doi.org/10.1038/s41746-025-01901-x
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://www.ebi.ac.uk/europepmc/ (PMC12361413)
- **Design:** Randomized experiment, 276 practising clinicians, each rated one vignette of a physician who used no generative AI, used it as the primary decision tool, or used it to verify a decision. Rated clinical skill, overall healthcare experience and competence on 1 to 7 scales.
- **Finding:** Clinical skill was rated 5.93 with no AI, 3.79 when AI was the primary tool, and 4.99 when AI was used only to verify. Similar patterns for competence and experience. Clinicians still agreed that generative AI improves accuracy (4.30) and rated an institution-customized tool more favourably (4.96).
- **Relation:** use penalty, with a framing moderator: "used it to check" softens but does not remove the penalty. Relevant to how observers read a checker as verification. Here peers judge a peer; in our study a non-expert with power judges an expert.
- **Access:** open access https://www.nature.com/articles/s41746-025-01901-x
- **Score:** 1=4 2=4 3=4 4=5 5=4 total=21/25

### Schilke, O., & Reimann, M. (2025). The transparency dilemma: How AI disclosure erodes trust. *Organizational Behavior and Human Decision Processes*, *188*, Article 104405. https://doi.org/10.1016/j.obhdp.2025.104405
- **Type:** management or economics
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/https://doi.org/10.1016/j.obhdp.2025.104405; study examples from https://news.arizona.edu/news/disclosing-ai-use-can-backfire-research-shows
- **Design:** 13 experiments, over 5,000 participants in total, across tasks (communication, analytics, artistry) and actors (supervisors, subordinates, professors, analysts, creatives, investment funds). Manipulated whether the actor disclosed AI use, the framing of the disclosure, voluntary versus mandatory disclosure, and exposure by a third party. Measured trust in the actor and perceived legitimacy.
- **Finding:** In all 13 experiments, actors who disclosed AI use were trusted less than those who did not. Examples: students trusted a professor 16% less when told the professor graded with AI; clients trusted a designer 20% less. The drop ran through lower perceived legitimacy, held for softer framings and for mandatory disclosure, and was larger when a third party exposed the AI use. It was smaller, but not gone, among evaluators who like technology and believe AI is accurate.
- **Relation:** use penalty; theme anchor for "AI disclosure penalty" and for legitimacy as a mechanism. Our study inverts the roles: the evaluator uses the AI, and we ask whether the evaluator's own tool erodes trust in the person evaluated.
- **Access:** open access https://doi.org/10.1016/j.obhdp.2025.104405
- **Score:** 1=5 2=4 3=3 4=5 5=4 total=21/25

### Fine, A., Berthelot, E. R., & Marsh, S. (2025). Public perceptions of judges' use of AI tools in courtroom decision-making: An examination of legitimacy, fairness, trust, and procedural justice. *Behavioral Sciences*, *15*(4), Article 476. https://doi.org/10.3390/bs15040476
- **Type:** psychology
- **Verified:** Crossref 2026-10-18; abstract read at https://www.ebi.ac.uk/europepmc/ (PMC12024057)
- **Design:** 1,800 US adults, stratified by race or ethnicity and gender, read bail and sentencing vignettes in which a judge used varying levels of AI assistance or relied only on expertise. Measured perceived judicial legitimacy, procedural justice and trust in AI; open-ended responses.
- **Finding:** Judges who relied on their own expertise were generally rated as more legitimate than judges who used AI. Black participants rated AI-assisted decisions as fairer and trusted the AI more than White and Hispanic participants did.
- **Relation:** use penalty in a high-stakes governance setting (law); a power-asymmetric institution. Supports a legitimacy reading of the CISO's standing; our CISO does not use AI.
- **Access:** open access https://www.mdpi.com/2076-328X/15/4/476
- **Score:** 1=4 2=4 3=4 4=3 5=4 total=19/25

### Cavalier, J. S., Goldstein, B. A., Ravitsky, V., Bélisle-Pipon, J.-C., Bedoya, A., Maddocks, J., Klotman, S., Roman, M., Sperling, J., Xu, C., Poon, E. G., & Chowdhury, A. (2025). Ethics in patient preferences for artificial intelligence–drafted responses to electronic messages. *JAMA Network Open*, *8*(3), Article e250449. https://doi.org/10.1001/jamanetworkopen.2025.0449
- **Type:** medical
- **Verified:** Crossref 2026-10-18; abstract read at https://www.ebi.ac.uk/europepmc/ (PMC11897835)
- **Design:** Surveys of a Duke health-system patient advisory panel (2,511 invited, 1,455 responded, 2023). Varied the message author (AI or clinician), the disclosure (AI, human, none) and topic seriousness. Rated satisfaction, usefulness and feeling cared for on 5-point scales.
- **Finding:** Without being told, patients slightly preferred AI-drafted replies (satisfaction 0.30 points higher, feeling cared for 0.43 higher). When AI authorship was disclosed, satisfaction dropped slightly: 0.13 points below human disclosure and 0.09 below no disclosure. Over 75% were satisfied whatever the author or disclosure.
- **Relation:** use penalty, small; counter-weight showing the disclosure penalty can be small in a real patient population. Our study measures trust in a person, not satisfaction with a message.
- **Access:** open access https://jamanetwork.com/journals/jamanetworkopen/fullarticle/10.1001/jamanetworkopen.2025.0449
- **Score:** 1=2 2=4 3=3 4=3 5=5 total=17/25

### Calluso, C., Iacopino, V., & Piazza, A. (2026). Gendered perceptions and the use of generative AI in academia: An experimental vignette study. *Gender, Work & Organization*. Advance online publication. https://doi.org/10.1111/gwao.70275
- **Type:** management or economics
- **Verified:** Crossref 2026-10-18 (online 2026-10-10, no volume yet); abstract read at https://api.openalex.org/works/https://doi.org/10.1111/gwao.70275
- **Design:** 2 x 2 x 2 vignette experiment, 651 participants: academic's gender, whether the academic used generative AI, and task (teaching or research). Rated perceived competence and warmth.
- **Finding:** Using generative AI generally raised perceived competence (women gained most for research tasks) but lowered perceived warmth, and the warmth loss was larger for women, especially in research.
- **Relation:** counter-result to the competence penalty in Reif et al. and Reis et al.: the effect of AI use splits by dimension (competence up, warmth down). Supports measuring trust by dimension (competence, benevolence, integrity) in our study.
- **Access:** open access https://onlinelibrary.wiley.com/doi/10.1111/gwao.70275
- **Score:** 1=3 2=4 3=2 4=4 5=4 total=17/25

### Zhang, Z., Geng, J., & Qi, C. (2026). Observer perceptions of GenAI use and interpersonal trust: The roles of warmth, competence, and AI literacy. *Behavioral Sciences*, *16*(7), Article 1221. https://doi.org/10.3390/bs16071221
- **Type:** psychology
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/https://doi.org/10.3390/bs16071221
- **Design:** 406 participants; observers' perception of a target person's generative AI use, perceived warmth and competence of the target, interpersonal trust in the target, and the target's perceived AI literacy. The abstract reports associations, so the design appears correlational.
- **Finding:** Perceived AI use went with lower interpersonal trust, through lower perceived warmth and competence. The direct link held when the target seemed low in AI literacy and disappeared when the target seemed AI-literate.
- **Relation:** use penalty; measure (warmth and competence as mediators of interpersonal trust). Weak design for our purposes; include only if space allows.
- **Access:** open access https://www.mdpi.com/2076-328X/16/7/1221
- **Score:** 1=4 2=3 3=1 4=2 5=3 total=13/25

---

## Sub-theme C: People flagged by an AI detector

### Du, P., Liu, T., & Xian, X. (2026). Automation bias in teachers' evaluation of student writing: Effects of algorithmic warnings and visual risk cues in AI detection reports. *Frontiers in Psychology*, *17*, Article 1889402. https://doi.org/10.3389/fpsyg.2026.1889402
- **Type:** psychology
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/https://doi.org/10.3389/fpsyg.2026.1889402
- **Design:** 2 x 2 between-subjects experiment, 214 university teachers in China. All graded the same medium-quality social-science course paper with a fictitious AI-detection report showing a 7% or 87% AI detection rate, with or without red highlighting. Measured perceived likelihood of AI authorship, quality ratings (overall, percentage score, originality, language, logic), and intention to intervene.
- **Finding:** An 87% detection report raised perceived AI authorship and risk and lowered the paper's overall quality rating, percentage score, originality, language and logical-structure ratings, for an identical text. Red highlighting further raised suspicion and intention to intervene, and interacted with the warning level on several scores.
- **Relation:** flagged by AI; closest analogue to H1 outside medicine: a wrong automated flag, held by the evaluator with power over the person evaluated, lowers the evaluation of identical work. Also a design element: the visual salience of the flag matters (our panel's flag wording and colour). Our study adds the evaluated person's answer and measures trust, not grades.
- **Access:** open access https://www.frontiersin.org/journals/psychology/articles/10.3389/fpsyg.2026.1889402/full
- **Score:** 1=2 2=5 3=4 4=5 5=4 total=20/25

### Chu, H., Men, L. R., Liu, S., Yuan, S., & Sun, Y. (2024). *Nationality, race, and ethnicity biases in and consequences of detecting AI-generated self-presentations* [Preprint]. arXiv. https://doi.org/10.48550/arXiv.2412.18647
- **Type:** communication (preprint)
- **Verified:** DataCite DOI, checked on OpenAlex and https://arxiv.org/abs/2412.18647 on 2026-10-18 (Crossref does not hold arXiv DOIs); abstract read at https://arxiv.org/abs/2412.18647
- **Design:** Preregistered experiment, nationally representative US sample of 644. Participants judged college-application personal statements in which linguistic style and the applicant's race, ethnicity and nationality were varied; they judged whether AI was used and evaluated the applicant.
- **Finding:** Writing style drove AI detection most, but international applicants were more often judged as AI users, and Asian and Hispanic applicants labelled domestic were too. Once AI use was attributed, the statement was rated lower in quality and authenticity, and the applicant lower in competence, sociability, morality and expected success.
- **Relation:** flagged by AI (here, by a human suspicion of AI); H1 analogue: an AI-use attribution lowers ratings of the person across competence and morality. In our study the flag is about content accuracy, not authorship.
- **Access:** open access https://arxiv.org/pdf/2412.18647
- **Score:** 1=3 2=3 3=3 4=3 5=4 total=16/25

---

## Notes on known papers

- **Brief item "Zondag et al. 2025 JAMA Network Open":** I could not find a paper by Zondag on this topic. The JAMA Network Open 2025 paper on public perception of physicians who use AI is Reis, Reis & Kunde (entry above). The Nature Medicine 2024 paper is also by Reis, Reis & Kunde (not "Reis, Reis & Wörner").
- **Brief item "Yokoi et al. 2021":** not verified. The closest record found is Yokoi & Nakayachi (2018) in the *Japanese Journal of Social Psychology* on value similarity and trust in medical AI (https://www.jstage.jst.go.jp/article/jssp/34/1/34_1736/_article), which compares trust in AI, not evaluation of a physician who uses it. Not included.
- **Chen, Sun, Liao & Sundar (2026), IJHCS:** press coverage (https://www.psu.edu/news/research/story/ai-disagreement-may-shake-patient-trust-doctors) reports N = 135, a chatbot role-playing a doctor in a mental-health consultation, and that AI agreement made the doctor look lazier while AI disagreement raised perceived uncertainty and lowered perceived autonomy. If our review cites it, check these details against the paper; the "agreement makes the doctor look lazy" point is a second route to a use penalty and fits with Reif et al. 2025.
- **Arkes et al. 2007 and Shaffer et al. 2013:** an unverified 2011 JAIS-listed record (Joshi & Wolf, "Do Indian physicians derogate peers who use computer-based diagnostic aids?") appears to extend the derogation effect to physicians judging peers; no DOI found, not included. Yang et al. 2025 above is the peer-reviewed modern version of that peer-judgement question.

## Searches run

OpenAlex API (`api.openalex.org/works?search=`): "Pezzo physician evaluation decision aid error"; "Zondag public perception physicians artificial intelligence"; "Reis Wörner AI involvement medical advice trust"; "Kempt Nagel responsibility second opinions artificial intelligence"; "Yokoi artificial intelligence medical diagnosis trust physician"; "students accused AI detection false positive academic integrity trust"; "student perceptions AI detection tools generative AI writing fairness"; "teacher suspicion AI-generated student essays evaluation bias"; "AI text detectors false accusations students non-native"; "transparency dilemma how AI disclosure erodes trust"; "Randomized study of the impact of AI on perceived legal liability for radiologists"; "Madanay radiologist AI recommendation patient confidence"; "Chu Men detecting AI-generated self-presentations"; "Pezzo not consulting decision aid available"; "Perceptions of justice by algorithms Yalcin"; "managers blame decision makers overriding algorithmic recommendations Kremer"; "Do Indian physicians derogate peers computer-based diagnostic aids"; "Bernstein Baird artificial intelligence radiologist liability jurors brain bleed"; "Cavalier patient-clinician messaging AI disclosure satisfaction"; "patient perceptions clinician use of large language model trust experiment"; "perceived competence of professional using AI assistance observers experiment"; "doctor consults ChatGPT patient trust experiment vignette"; "AI-assisted writing disclosure author perceived trustworthiness experiment".

Web search: "Pezzo Pezzo 2006 physician evaluation after medical errors decision aid"; "JAMA Network Open 2025 public perceptions of physicians who use artificial intelligence"; "Reis Reis Wörner 2024 Nature Medicine believed AI involvement"; "experiment physician overrides AI recommendation patient trust blame liability vignette"; "disclosing AI use perceived competence trust experiment disclosure penalty legitimacy"; "peer perceptions clinicians using generative AI npj Digital Medicine"; "Tobia Nielsen Stremitzer 2021 physician use of AI liability"; "Yokoi 2021 AI trusted less than a doctor perceived care value similarity"; "students falsely accused AI detector trust teacher study 2024 2025"; "lawyer uses generative AI client trust experiment perceived competence"; "Bernstein mock jurors radiologist liability AI NEJM AI"; "observers judge experts who rely on algorithmic advice less competent"; "patient trust physician disagrees with AI diagnosis experiment 2024 2025"; "Kempt Nagel 2022 J Med Ethics second opinions burden of proof"; "observers evaluate expert who overrides algorithm versus follows algorithm blame"; "public perceptions judges using algorithmic risk assessment legitimacy"; "AI use coworkers perceived lazy less competent experiment"; "patient preferences AI-drafted responses disclosure JAMA Network Open Cavalier"; "students perceive instructor who uses AI for feedback or grading less credible"; "financial advisor uses AI client trust perceived competence experiment".

Gaps: no peer-reviewed experiment found on lawyers or financial advisors whose AI use is rated by clients (only an undergraduate project and trade surveys); no study found in which a manager or board member holds an AI checker while evaluating a subordinate expert's presentation (Du et al. 2026 with teachers is the nearest); the Kremer et al. study on managers blaming decision makers who override algorithms was seen only as a seminar listing and could not be verified.
