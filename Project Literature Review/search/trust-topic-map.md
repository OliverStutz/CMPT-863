# Trust topic map: where our study sits, and what lies around it

Written 2026-10-18. Purpose: widen the literature review beyond the six search themes and show research directions the project has not yet considered. Every anchor reference was checked against Crossref (`https://api.crossref.org/works/<DOI>`), OpenAlex or the publisher page on 2026-10-18 unless marked otherwise. "(found)" means the paper is already in `_master-index.md`; "(known)" means it is in `_known-papers.md`. Findings are stated only where an abstract was read; otherwise the claim is kept general.

## 1. The phenomenon decomposed

Two humans, one tool. The design space has nine dimensions:

| Dimension | Values |
|---|---|
| Who holds the tool | the speaker (expert), the listener (judge), a third party (moderator, broadcaster, platform), both |
| Power direction | tool holder above the checked person, equal, below |
| Expertise gap | holder is the novice, peers, holder is the expert |
| Tool correctness | right, wrong, right by the textbook but wrong in context, mixed |
| Visibility of the check | private to the holder, visible to the checked person, public to an audience |
| Reply | checked person can answer, cannot answer, answers only later (appeal) |
| Synchrony | live, after the fact (minutes, transcript, recording) |
| Stakes | low (a quiz), professional (a budget, a diagnosis), legal or life-and-death |
| Modality | text, audio, vision (camera on slides), multimodal |

**Our study in one line:** the listener holds a vision checker, has more power and less expertise, the flags are wrong in context, the check is private but can be voiced, the expert replies live, the stakes are a corporate budget.

**Three unstudied cells that would make good follow-ups:**

1. **Visible check.** The same wrong flags, but shown on a shared screen so the CISO and the other directors see them. This adds audience and face (Section 3, M5) and resembles VAR replays on a stadium screen and the live "worm" in TV debates. We found no study of trust in an expert when the AI flag is public to a group.
2. **Reversed power.** A junior employee or analyst holds the checker while the senior executive presents. This is sousveillance (Mann et al., 2003, known) with an AI. The literature on speaking up (aviation, psychological safety) predicts the tool will be used to raise concerns that the junior would otherwise swallow; nobody has tested whether it then lowers trust in the senior.
3. **The expert holds the checker too.** The CISO pre-empts with "I ran these slides through the same checker; here is why it flags slide 3". This is inoculation by the checked party. It also tests whether disclosing AI use costs the expert trust (Reif et al., 2025, known) when the audience already uses AI.

## 2. Adjacent topics (27)

Each entry: the phenomenon, why it is adjacent (what transfers, what does not), anchors, coverage.

### 2.1 Sports officiating aids: VAR, Hawk-Eye and player challenges
Referees are now checked live by video and ball-tracking systems, and in tennis a player may challenge the umpire a limited number of times. **Transfers:** a machine verdict replaces a human judgement in public, and players ration challenges like our board member choosing to challenge. **Does not:** the machine is usually right and has institutional authority; the referee rarely gets to reply. Collins and Evans argue that Hawk-Eye's reconstructions hide measurement error and make naive viewers overestimate how well a machine settles human disagreement, which is the risk of our checker's confident "Numbers do not add up". Players' challenges track stakes and the value of challenges left, close to optimal.
- Collins, H., & Evans, R. (2008). You cannot be serious! Public understanding of technology with special reference to "Hawk-Eye". *Public Understanding of Science*, 17(3). https://doi.org/10.1177/0963662508093370
- Abramitzky, R., Einav, L., Kolkowitz, S., & Mill, R. (2012). On the optimality of line call challenges in professional tennis. *International Economic Review*, 53(3). https://doi.org/10.1111/j.1468-2354.2012.00706.x
- Scanlon, C., Griggs, G., & McGillick, C. (2022). "It's not football anymore": Perceptions of the video assistant referee by English Premier League football fans. *Soccer & Society*. https://doi.org/10.1080/14660970.2022.2033731 (100 fans; themes of accuracy, flow, being "in the dark")
- **Coverage: new direction.**

### 2.2 Expert witnesses, juries and cross-examination
Lay jurors must judge experts whose content they cannot evaluate, and opposing counsel attacks the expert in front of them. **Transfers:** non-expert judge, adversarial challenge, expert reply. **Does not:** the challenger is a human advocate, not a neutral-seeming tool. Cooper, Bennett and Sukel found that when testimony was complex, mock jurors leaned on the expert's credentials rather than the content, which predicts that our board member judges the CISO's reason by cues, not substance. Bernstein et al. (2025, found) bring AI into this setting.
- Cooper, J., Bennett, E. A., & Sukel, H. L. (1996). Complex scientific testimony: How do jurors make decisions? *Law and Human Behavior*, 20(4). https://doi.org/10.1007/BF01498976
- **Coverage: partly (Bernstein found); jury-expert literature new.**

### 2.3 Interpreters and machine translation between two people
A tool sits between two speakers and changes what each hears. **Transfers:** the tool shapes one person's impression of the other in a live conversation; errors get attributed to the person (Vignovic & Thompson, 2010, known). **Does not:** the tool does not judge the speaker. Mehandru et al. studied clinicians using machine translation with patients and how they guard against its errors.
- Hara, K., & Iqbal, S. T. (2015). Effect of machine translation in interlingual conversation. *Proc. CHI 2015*. https://doi.org/10.1145/2702123.2702407
- Mehandru, N., Robertson, S., & Salehi, N. (2022). Reliable and safe use of machine translation in medical settings. *Proc. FAccT 2022*. https://doi.org/10.1145/3531146.3533244
- **Coverage: new direction.**

### 2.4 Aviation: crew resource management, speaking up, automation that contradicts the crew
Cockpits have a steep authority gradient, and first officers often hint rather than challenge. Automation (alerts, TCAS) can contradict the captain. **Transfers:** a checker can lower the cost of challenging upward, and automation bias leads to errors of commission when the automation is wrong. **Does not:** in the cockpit the challenger is the junior; in our boardroom the challenger is senior but less expert. Mosier et al. documented automation bias in pilots; Bienefeld and Grote studied why aircrew do not speak up.
- Mosier, K. L., Skitka, L. J., Heers, S., & Burdick, M. (1998). Automation bias: Decision making and performance in high-tech cockpits. *International Journal of Aviation Psychology*, 8(1). https://doi.org/10.1207/s15327108ijap0801_3
- Helmreich, R. L., Merritt, A. C., & Wilhelm, J. A. (1999). The evolution of crew resource management training in commercial aviation. *International Journal of Aviation Psychology*, 9(1). https://doi.org/10.1207/s15327108ijap0901_2
- Bienefeld, N., & Grote, G. (2012). Silence that may kill: When aircrew members don't speak up and why. *Aviation Psychology and Applied Human Factors*, 2(1). https://doi.org/10.1027/2192-0923/a000021
- **Coverage: automation bias covered (Parasuraman & Manzey, Goddard, found); crew authority and speaking up new.**

### 2.5 False alarms and the "cry wolf" effect
When an alerting system is often wrong, operators discount it. **Transfers:** our checker is wrong on three of five slides; repeated wrong flags may make the board member discount the tool, which would let trust in the CISO recover by slide 4. **Does not:** studies measure response to the alarm, not trust in a person. In air traffic control 45% of 495 conflict alerts were false, yet controllers did not ignore true alerts; they anticipated conflicts instead.
- Wickens, C. D., et al. (2009). False alerts in air traffic control conflict alerting system: Is there a "cry wolf" effect? *Human Factors*, 51(4). https://doi.org/10.1177/0018720809344720
- Ancker, J. S., et al. (2017). Effects of workload, work complexity, and repeated alerts on alert fatigue in a clinical decision support system. *BMC Medical Informatics and Decision Making*, 17. https://doi.org/10.1186/s12911-017-0430-8
- **Coverage: new direction.**

### 2.6 Patients who bring outside information to the doctor
The internet, and now chatbots, give the less expert party material to question the expert. **Transfers:** the layperson brings an outside verdict into a live consultation; the expert must answer. This is the closest everyday analogue to our board member. Murray et al. surveyed 3,209 US adults: 8% had taken internet information to their physician, most wanted the physician's opinion rather than a specific intervention, and the effect on the relationship was likely positive when the physician responded well. Li, Li and Kuang (2026) study chatbot pre-consultations and later trust in the physician.
- Murray, E., et al. (2003). The impact of health information on the internet on the physician-patient relationship. *Archives of Internal Medicine*, 163(14). https://doi.org/10.1001/archinte.163.14.1727
- Li, ?, Li, ?, & Kuang, ? (2026). Bridge or barrier? The impact of GAI chatbot-based pre-consultations on patients' subsequent trust in physicians. *Internet Research*. https://doi.org/10.1108/INTR-09-2025-1539 (metadata verified; abstract not read, first names not checked)
- **Coverage: AI second opinion covered (Chen known, Madanay found); patient-initiated challenge new.**

### 2.7 Audit, professional skepticism and auditors with AI
Auditors exist to doubt management; the "audit society" turns verification into ritual (Power, 1997, known). **Transfers:** the board's role is oversight, and the checker gives it an audit tool. **Does not:** auditors are experts themselves. Commerford et al. found auditors proposed smaller adjustments to management's estimates when contradictory evidence came from the firm's AI rather than a human specialist, an algorithm aversion that would weaken our H1 if board members behave like auditors.
- Nelson, M. W. (2009). A model and literature review of professional skepticism in auditing. *Auditing: A Journal of Practice & Theory*, 28(2). https://doi.org/10.2308/aud.2009.28.2.1
- Commerford, B. P., Dennis, S. A., Joe, J. R., & Ulla, J. W. (2022). Man versus machine: Complex estimates and auditor reliance on artificial intelligence. *Journal of Accounting Research*, 60(1). https://doi.org/10.1111/1475-679X.12407 (finding from the SSRN abstract)
- **Coverage: Power known; audit experiments new.**

### 2.8 Lie detection, polygraphs and the truth default
People detect lies only a little above chance and default to believing others. Machine lie detectors change who accuses whom. **Transfers:** a checker breaks the truth default; von Schenk et al. (2024, known) and Loconte et al. (2026, found) are already ours. **Does not:** our flags concern facts, not honesty, but "numbers do not add up" can be heard as an integrity charge.
- Bond, C. F., & DePaulo, B. M. (2006). Accuracy of deception judgments. *Personality and Social Psychology Review*, 10(3). https://doi.org/10.1207/s15327957pspr1003_2
- Levine, T. R. (2014). Truth-default theory (TDT). *Journal of Language and Social Psychology*, 33(4). https://doi.org/10.1177/0261927X14535916
- National Research Council. (2003). *The polygraph and lie detection*. https://doi.org/10.17226/10420
- **Coverage: AI lie detection found; truth default and polygraph new.**

### 2.9 Body cameras, sousveillance and trust in police
A recording device watches an authority during an encounter. **Transfers:** monitoring changes the behaviour of the watched (Ariel et al.: use of force and complaints fell) and the perceived legitimacy of the watched. Demir et al. found motorists stopped by camera-wearing officers rated procedural justice and legitimacy higher, which is the opposite sign to Strickland (known): verification can raise trust when it is seen as protecting the citizen. **Does not:** the camera records, it does not judge.
- Ariel, B., Farrar, W. A., & Sutherland, A. (2015). The effect of police body-worn cameras on use of force and citizens' complaints against the police. *Journal of Quantitative Criminology*, 31. https://doi.org/10.1007/s10940-014-9236-3
- Demir, M., Apel, R., Braga, A. A., Brunson, R. K., & Ariel, B. (2020). Body worn cameras, procedural justice, and police legitimacy. *Justice Quarterly*, 37(1). https://doi.org/10.1080/07418825.2018.1495751
- **Coverage: sousveillance known (Mann 2003); body cameras new.**

### 2.10 Peer review and automated error checkers (statcheck)
In 2016 statcheck scanned about 50,000 psychology papers and posted flags publicly on PubPeer. Authors reported feeling exposed; the German Psychological Society objected that false positives could not be answered or removed. **Transfers:** an automated public flag on an expert's work, some flags wrong (for example corrected p-values), and a fight over the right to reply. **Does not:** after the fact, not live.
- Nuijten, M. B., et al. (2016). The prevalence of statistical reporting errors in psychology (1985-2013). *Behavior Research Methods*, 48. https://doi.org/10.3758/s13428-015-0664-2
- Retraction Watch (2016, 25 October). Psychological society wants end to posting error-finding algorithm results publicly. https://retractionwatch.com/2016/10/25/psychological-society-wants-end-to-posting-error-finding-algorithm-results-publicly/ (trade source, accessed 2026-10-18)
- **Coverage: new direction.**

### 2.11 AI note-takers and meeting assistants
Meeting assistants transcribe, summarise and can be asked to critique a participant. **Transfers:** a live meeting, a tool any attendee can query, power shifting toward whoever holds AI output. **Does not:** most work studies productivity, not trust in colleagues. Microsoft's Copilot prompt "Where are some holes in [a participant's] argument?" (found) is the product anchor. A University of Sydney thesis reports that meeting assistants shift power toward people with AI outputs (thesis, not peer reviewed; https://ses.library.usyd.edu.au/handle/2123/33975, unverified beyond the record page).
- **Coverage: product found; scholarly work thin, a gap in itself.**

### 2.12 Deepfake detection and the liar's dividend
Detectors label video as fake; liars exploit the existence of fakes to deny real evidence. **Transfers:** Groh et al. (15,016 participants) found people with access to a detector's prediction were more accurate overall, but wrong model predictions often lowered their accuracy, our wrong-flag case. Chesney and Citron named the liar's dividend, which Schiff et al. (2025, found) tested: a bare denial can pay.
- Groh, M., Epstein, Z., Firestone, C., & Picard, R. (2022). Deepfake detection by human crowds, machines, and machine-informed crowds. *PNAS*, 119(1). https://doi.org/10.1073/pnas.2110013119
- Chesney, R., & Citron, D. (2018). Deep fakes: A looming challenge for privacy, democracy, and national security. SSRN (later *California Law Review*). https://doi.org/10.2139/ssrn.3213954
- **Coverage: Schiff found; detector-assisted judgement new.**

### 2.13 AI detectors in education
Teachers hold detectors that flag students. **Transfers:** the powerful party holds a fallible flag about a less powerful one. **Does not:** the expertise gap runs the other way. Du et al. (2026, found), Chu et al. (2024, found), Gorichanaz (2023, known), Liang et al. (2023, known).
- **Coverage: covered.**

### 2.14 Algorithmic risk scores in courts and lending
Judges and loan officers receive a risk score and may override it. **Transfers:** a decision maker weighs a machine judgement of a third person. Green and Chen found people given a risk assessment did not reach its accuracy and were biased in how they used it; Dressel and Farid found untrained crowd workers about as accurate as COMPAS. **Does not:** the judged person has no voice in the moment.
- Green, B., & Chen, Y. (2019). Disparate interactions: An algorithm-in-the-loop analysis of fairness in risk assessments. *Proc. FAT* 2019*. https://doi.org/10.1145/3287560.3287563
- Dressel, J., & Farid, H. (2018). The accuracy, fairness, and limits of predicting recidivism. *Science Advances*, 4(1). https://doi.org/10.1126/sciadv.aao5580
- **Coverage: De-Arteaga et al. (2020, found) close; courts new.**

### 2.15 Algorithmic management and electronic monitoring
Workers managed and rated by algorithms. **Transfers:** monitoring as a signal of distrust, contested terrain of control. Lee et al. (2015, known), Kellogg et al. (2020, known), Ravid et al. (known).
- **Coverage: covered.**

### 2.16 Organisational silence, voice and psychological safety
Employees and directors withhold concerns because speaking up feels unsafe. **Transfers:** Opdenbusch et al. (2025, found) found boards do not question CISOs for fear of being exposed as novices. Westphal and Bednar found outside directors underestimate how far fellow directors share their concerns, so they stay silent (pluralistic ignorance). A checker may break that silence by giving the director a reason that is "not mine". **Does not:** voice research studies the subordinate, not the superior.
- Morrison, E. W., & Milliken, F. J. (2000). Organizational silence. *Academy of Management Review*, 25(4). https://doi.org/10.2307/259200
- Edmondson, A. (1999). Psychological safety and learning behavior in work teams. *Administrative Science Quarterly*, 44(2). https://doi.org/10.2307/2666999
- Detert, J. R., & Edmondson, A. C. (2011). Implicit voice theories. *Academy of Management Journal*, 54(3). https://doi.org/10.5465/amj.2011.61967925
- Westphal, J. D., & Bednar, M. K. (2005). Pluralistic ignorance in corporate boards and firms' strategic persistence. *Administrative Science Quarterly*, 50(2). https://doi.org/10.2189/asqu.2005.50.2.262
- **Coverage: Opdenbusch found; voice and board silence theory new.**

### 2.17 Boards as monitors and advisers: the information-sharing trade-off
Adams and Ferreira model the board's two roles: an executive who shares information gets better advice but is monitored harder, so may withhold from a tough board. **Transfers:** a checker makes the board a tougher monitor, which predicts that a real CISO would share less, the long-run cost our one-shot study cannot see. **Does not:** formal model, no trust measure.
- Adams, R. B., & Ferreira, D. (2007). A theory of friendly boards. *Journal of Finance*, 62(1). https://doi.org/10.1111/j.1540-6261.2007.01206.x
- **Coverage: new direction (Lowry et al. 2026, found, is the nearest).**

### 2.18 Epistemology of testimony and epistemic dependence
Philosophers ask how a novice can rationally trust an expert, and how to choose between two experts who disagree. **Transfers:** the checker is a second "expert" contradicting the first; Goldman's novice-two-experts problem lists the cues a novice can use (arguments, agreement of others, credentials, track record), which maps onto what the CISO's reason and the checker's record supply. Sperber et al. describe epistemic vigilance, the mechanisms that check communicated information. **Does not:** normative, not empirical.
- Hardwig, J. (1985). Epistemic dependence. *Journal of Philosophy*, 82(7). https://doi.org/10.2307/2026523
- Goldman, A. I. (2001). Experts: Which ones should you trust? *Philosophy and Phenomenological Research*, 63(1). https://doi.org/10.1111/j.1933-1592.2001.tb00093.x
- Sperber, D., et al. (2010). Epistemic vigilance. *Mind & Language*, 25(4). https://doi.org/10.1111/j.1468-0017.2010.01394.x
- **Coverage: new direction.**

### 2.19 Epistemic and testimonial injustice
Testimonial injustice is giving a speaker less credibility because of who they are. **Transfers:** a checker that systematically flags certain presenters (non-native speakers, unusual architectures, small firms) would deflate their credibility; Liang et al. (2023, known) show detectors flag non-native writers. **Does not:** our CISO is one person, no identity manipulation, which is a follow-up cell.
- Fricker, M. (2007). *Epistemic injustice: Power and the ethics of knowing*. Oxford University Press. https://doi.org/10.1093/acprof:oso/9780198237907.001.0001
- Symons, J., & Alvarado, R. (2022). Epistemic injustice and data science technologies. *Synthese*, 200. https://doi.org/10.1007/s11229-022-03631-z
- **Coverage: new direction.**

### 2.20 "Trust but verify" in diplomacy and arms control
Arms control rests on inspection between parties who do not fully trust each other. **Transfers:** a body of argument that verification and trust are interdependent rather than opposites: accepting inspection can signal trustworthiness. This is the counterweight to Strickland (known). **Does not:** states, not persons; verification is agreed in advance.
- Abbott, K. W. (1993). "Trust but verify": The production of information in arms control treaties and other international agreements. *Cornell International Law Journal*, 26(1). https://scholarship.law.cornell.edu/cilj/vol26/iss1/1 (no DOI; publisher page found, not read)
- Larson, D. W. (1997). Trust and missed opportunities in international relations. *Political Psychology*, 18(3). https://doi.org/10.1111/0162-895X.00075
- **Coverage: new direction.**

### 2.21 Devil's advocacy and red teams
Organisations assign a dissenter to test a plan. **Transfers:** the checker is a mechanised devil's advocate. Nemeth et al. found authentic minority dissent produced more and better solutions than any form of assigned devil's advocacy, which suggests a checker's objection may be discounted as "role play" if people see it as built to object. Chiang et al. (2024, found) built an LLM devil's advocate. **Does not:** studies target decision quality, not trust in the presenter.
- Schwenk, C. R. (1990). Effects of devil's advocacy and dialectical inquiry on decision making: A meta-analysis. *OBHDP*, 47(1). https://doi.org/10.1016/0749-5978(90)90051-A
- Nemeth, C., Brown, K., & Rogers, J. (2001). Devil's advocate versus authentic dissent. *European Journal of Social Psychology*, 31(6). https://doi.org/10.1002/ejsp.58
- **Coverage: Chiang found; dissent theory new.**

### 2.22 Chess and engine-based anti-cheating
Engines check whether a player's moves match machine play; accusations follow statistical flags. **Transfers:** a statistical flag about a person, high reputational stakes, false positives that are hard to answer. **Does not:** the accused rarely replies live.
- Regan, K. W., & Haworth, G. M. (2011). Intrinsic chess ratings. *Proc. AAAI*, 25(1). https://doi.org/10.1609/aaai.v25i1.7951
- Barnes, D. J., & Hernandez-Castro, J. (2015). On the limits of engine analysis for cheating detection in chess. *Computers & Security*, 48. https://doi.org/10.1016/j.cose.2014.10.002 (abstract not read)
- **Coverage: new direction.**

### 2.23 Triadic trust: the third party in a dyad
Balance theory: if I like the tool and the tool dislikes you, I should like you less. **Transfers:** predicts that the size of H1 depends on trust in the checker, and that a CISO who discredits the checker restores balance. Erengin et al. found a human teammate's trust in an AI teammate predicted a focal employee's trust in that AI, but only when the human teammate was seen as trustworthy: trust flows through people. Hohenstein and Jung (2020, known) show blame shifting to the AI.
- Heider, F. (1946). Attitudes and cognitive organization. *Journal of Psychology*, 21(1). https://doi.org/10.1080/00223980.1946.9917275
- Erengin, T., Briker, R., & de Jong, S. B. (2025). You, me, and the AI: The role of third-party human teammates for trust formation toward AI teammates. *Journal of Organizational Behavior*, 47(8). (initials of co-authors not checked) https://doi.org/10.1002/job.2857
- Etzrodt, K. (2022). The third party will make a difference. *International Journal of Human-Computer Studies*, 168. https://doi.org/10.1016/j.ijhcs.2022.102901 (abstract not read)
- **Coverage: new direction.**

### 2.24 Face threat and challenging an expert in public
Questioning someone in front of others threatens their face and the questioner's. **Transfers:** "the checker says the numbers do not add up" is a face threat with an off-record shield; it lets the director challenge while attributing the challenge to the tool. **Does not:** our study is one-on-one online, so the audience is imagined.
- Goffman, E. (1955). On face-work. *Psychiatry*, 18(3). https://doi.org/10.1080/00332747.1955.11023008
- Brown, P., & Levinson, S. C. (1987). *Politeness: Some universals in language usage*. Cambridge University Press. https://doi.org/10.1017/CBO9780511813085
- **Coverage: new direction.**

### 2.25 Live audience-response displays in debates
TV debates overlay a live "worm" of audience reactions. Davis et al. manipulated the worm on a real UK debate (150 participants): most did not notice, yet it shifted who they thought won, their preferred prime minister and voting intention. **Transfers:** a live, confident, machine-looking signal next to a speaker changes judgement of the speaker, even when false. **Does not:** the worm reports other people, not a verdict on facts.
- Davis, C. J., Bowers, J. S., & Memon, A. (2011). Social influence in televised election debates: A potential distortion of democracy. *PLoS ONE*, 6(3), e18154. https://doi.org/10.1371/journal.pone.0018154
- **Coverage: new direction; complements live fact-checking (Wintersieck, known).**

### 2.26 Persuasion knowledge and customers who check salespeople
A CISO asking for budget is persuading. Customers now check salespeople's claims on their phones in the shop ("showrooming"). **Transfers:** a checker may activate persuasion knowledge ("she wants money, so of course the slide is rosy"), turning a factual flag into a motive attribution and hitting benevolence and integrity, not only ability. **Does not:** retail stakes are low and repeated.
- Friestad, M., & Wright, P. (1994). The persuasion knowledge model. *Journal of Consumer Research*, 21(1). https://doi.org/10.1086/209380
- Rapp, A., Baker, T. L., Bachrach, D. G., Ogilvie, J., & Beitelspacher, L. S. (2015). Perceived customer showrooming behavior and the effect on retail salesperson self-efficacy and performance. *Journal of Retailing*, 91(2). https://doi.org/10.1016/j.jretai.2014.12.007 (abstract not read; the checked party's side)
- **Coverage: new direction.**

### 2.27 Contestability of automated flags: content moderation and Wikipedia
People whose posts or edits are flagged by algorithms can sometimes appeal. Vaccaro et al. found none of their appeal designs improved perceived fairness, accountability, trustworthiness or control compared with no appeal. Halfaker et al. tie Wikipedia's newcomer decline partly to algorithmic tools that reject contributions. **Transfers:** the right to reply is the heart of H2 and H3, and these studies say a reply channel alone does not repair perceptions. **Does not:** the judged party is the one rating, not an observer.
- Vaccaro, K., Sandvig, C., & Karahalios, K. (2020). "At the end of the day Facebook does what it wants". *PACM HCI*, 4(CSCW2), Article 167. https://doi.org/10.1145/3415238
- Lyons, H., Velloso, E., & Miller, T. (2021). Conceptualising contestability. *PACM HCI*, 5(CSCW1). https://doi.org/10.1145/3449180
- Halfaker, A., Geiger, R. S., Morgan, J. T., & Riedl, J. (2013). The rise and decline of an open collaboration system. *American Behavioral Scientist*, 57(5). https://doi.org/10.1177/0002764212469365
- **Coverage: new direction.**

## 3. Mechanisms that cut across topics

**M1. Burden of proof shifts to the checked person.** Shown in AI second opinions (Kempt & Nagel, 2022, found, https://doi.org/10.1136/medethics-2021-107440), cross-examination (2.2), statcheck (2.10), contestability (2.27). Predicts: H1 yes; H3 holds because a bare denial does not discharge the burden; H2 depends on whether the reason is judged adequate by a novice.

**M2. Machine heuristic and automation bias.** Machines seen as objective (Sundar & Kim, 2019, found, https://doi.org/10.1145/3290605.3300768); commission errors under wrong automation (2.4; Dratsch et al., 2023, found); wrong detector output lowers accuracy (2.12); Hawk-Eye overrated (2.1). Predicts: strong H1, more challenges (H4a). Counter: algorithm aversion among auditors (2.7) and people resisting confident AI accusations (Loconte et al., 2026, found).

**M3. Innuendo and belief echoes.** An accusation leaves a residue after it is answered (Wegner et al., 1981, found; Thorson, 2016, found). Seen again in the liar's dividend (2.12) and in public statcheck flags that "cannot be removed" (2.10). Predicts: H2 recovery short of the start.

**M4. Triadic balance and cry wolf.** Trust in the person depends on trust in the tool (2.23), and trust in the tool falls with repeated false alerts (2.5). Predicts: H1 shrinks from slide 2 to slide 4; participants who distrust the checker show no H1; a reason that exposes the checker's error (H2) may overshoot because it lowers the tool and raises the person at once. This is a mechanism for H2 "above the start" that the trust-repair literature does not offer.

**M5. Face threat and attribution of the challenge to the tool.** Novices stay silent to protect face (2.16, 2.24; Opdenbusch et al., 2025, found). A tool gives an off-record shield. Predicts: H4a (more questions) and an attribution question: does the CISO's reply restore trust in the CISO, or merely lower trust in the tool?

**M6. Monitoring signals distrust and changes the monitored.** Strickland (1958, known), Falk and Kosfeld (2006, known), body cameras (2.9), friendly boards (2.17). Predicts nothing about H1 to H4 directly but warns that the real cost may be the CISO withholding information later; and the body-camera result shows the sign can flip when monitoring is seen as protective.

**M7. Novices judge experts by cues, not content.** Jurors use credentials (2.2), Goldman's cue list (2.18), boards judge CISOs by communication (Proudfoot et al., 2023, found). Predicts: H2 depends on delivery as much as substance; a confident bare denial may work better than expected (Schiff et al., 2025, found), which threatens H3.

**M8. Live signals as social proof.** A confident display next to a speaker changes judgements even when false and unnoticed (2.25). Predicts: H1 even among participants who say they ignored the checker; supports measuring trust implicitly as well.

## 4. Research questions the map suggests

1. When the checker's flags are shown on a shared screen to the CISO and other directors, does the trust drop after a flag grow, and does the CISO's reason repair more or less than in private (cell 1, 2.24, 2.25)?
2. When a junior employee holds the checker during a senior executive's briefing, do they challenge more, and does the senior lose more trust than in the board case (cell 2, 2.4, 2.16)?
3. Does a CISO who pre-empts the checker ("I ran these slides through the same tool") lose trust for disclosing AI use or gain it as an inoculation (cell 3, 2.26)?
4. Does trust in the CISO recover across slides 2 to 4 as the board member's trust in the checker falls, and does trust in the checker mediate the effect of each flag (2.5, 2.23, M4)?
5. Is a CISO's reply that explains why the checker is wrong ("the tool assumes a textbook network") more restorative than a reply that only justifies the company's choice (M4, 2.18)?
6. Do board members phrase challenges as "the tool says" rather than "I think", and does that attribution change how much trust the CISO's answer recovers (M5, 2.24)?
7. Does a flag that suggests a motive ("figures look inflated") lower benevolence and integrity more than a flag about accuracy, through persuasion knowledge (2.26, M2)?
8. Does a limited number of permitted challenges, as in tennis, make board members use the checker more selectively and reduce wrongful trust loss (2.1)?
9. Does the effect of a wrong flag on trust in the expert differ when the expert is a member of a group the checker is known to over-flag (2.19)?
10. After a briefing under a visible checker, does a CISO disclose less in a later briefing, as the friendly-board model predicts (2.17, M6)?

## 5. Verification notes

- All DOIs above resolved on Crossref on 2026-10-18 with matching title and first author, except the arms-control law review (no DOI; URL given) and the Sydney thesis (record page only). Abstracts were read through OpenAlex or the publisher for Collins & Evans, Abramitzky et al., Scanlon et al., Davis et al., Adams & Ferreira, Westphal & Bednar, Groh et al., Murray et al., Nemeth et al., Halfaker et al., Vaccaro et al., Wickens et al., Erengin et al.; Commerford et al. via the SSRN abstract; Demir et al. and the statcheck controversy via search summaries of the publisher and Retraction Watch pages.
- Not read beyond metadata, so claims kept general: Li, Li & Kuang (2026), Barnes & Hernandez-Castro (2015), Etzrodt (2022), Rapp et al. (2015), Symons & Alvarado (2022), Lyons et al. (2021), Ancker et al. (2017), Hara & Iqbal (2015), Mehandru et al. (2022), Green & Chen (2019) and Dressel & Farid (2018) (findings stated from prior knowledge of widely cited results; re-check before citing), Abbott (1993).
- Author first names for Li, Li & Kuang (2026) not checked; the Crossref record lists only family names in our query.
