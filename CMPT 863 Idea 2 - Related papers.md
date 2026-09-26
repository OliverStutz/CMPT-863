# Idea 2 "Checked by the Camera": related work on trust between humans when one checks the other's work

Compiled 2026-09-25 from four literature searches. Each entry gives the key reference, what the paper is about, and how it relates to Idea 2 (an automated slide checker in the audience's hands, its effect on trust in the presenter, and on the presenter's own behaviour). Papers are grouped by the question they answer. Three validator passes checked every entry against Crossref, Semantic Scholar, PubMed or the publisher page (75 entries, all resolved); corrections were applied and the remaining flags are listed at the end.

**Headline finding of the search:** nobody has yet run a study where the *audience* holds an automated checker, the checker contradicts a live human presenter, and trust in that presenter is the measured outcome. The closest designs are Chen et al. (2026, AI second opinion contradicting a doctor), Gupta et al. (CHI 2026, wearable fact-checker of spoken claims, measures belief not trust), and Garrett & Weeks (CSCW 2013, real-time inline corrections shift source trust). Everything else below supplies the mechanisms and the pre-AI baseline.

---

## A. Pre-AI baseline: monitoring, verification and control read as distrust

- **Strickland, L. H. (1958). Surveillance and trust. *Journal of Personality*, 26(2), 200–215.** https://doi.org/10.1111/j.1467-6494.1958.tb01580.x
  *About:* Participants acting as supervisors oversaw two subordinates with equal output but monitored one more often. They ended up trusting the less-monitored subordinate more and attributed the monitored one's good behaviour to the surveillance, not to him.
  *Relation:* The founding result for our core hypothesis. The act of checking itself produces a discount on the checked person even when outcomes are identical, which is exactly the "beyond what the errors justify" effect we want to measure.

- **Kruglanski, A. W. (1970). Attributing trustworthiness in supervisor-worker relations. *Journal of Experimental Social Psychology*, 6(2), 214–232.** https://doi.org/10.1016/0022-1031(70)90088-0
  *About:* Extends Strickland on the supervisor side. Supervisors who surveil attribute compliance to the surveillance, rate the highly monitored worker as less trustworthy, and then surveil more. (A worker-side behavioural response is often attributed to this paper in secondary sources but is not confirmed; check the full text before citing it for that.)
  *Relation:* Gives us the self-reinforcing loop on the audience side: checking lowers attributed trustworthiness, which justifies more checking. Pairs with Falk & Kosfeld for the presenter side.

- **Enzle, M. E., & Anderson, S. C. (1993). Surveillant intentions and intrinsic motivation. *Journal of Personality and Social Psychology*, 64(2), 257–266.** https://doi.org/10.1037/0022-3514.64.2.257
  *About:* Surveillance whose purpose was unexplained was read by the watched person as signalling distrust and an intention to evaluate, and it undermined intrinsic motivation. Surveillance framed as non-controlling did not.
  *Relation:* Framing lever for Study 2. Whether we tell presenters the checker is "there to help the audience follow" or "there to catch errors" should change hedging and preparation.

- **Frey, B. S. (1993). Does monitoring increase work effort? The rivalry with trust and loyalty. *Economic Inquiry*, 31(4), 663–670.** https://doi.org/10.1111/j.1465-7295.1993.tb00897.x
  *About:* Argues that where a personal relationship exists, intensified monitoring is read as distrust and crowds out effort; the disciplining effect wins only in impersonal relations.
  *Relation:* Predicts that presenter behaviour change should be strongest in small, personal audiences (a board, a seminar) rather than anonymous ones.

- **Falk, A., & Kosfeld, M. (2006). The hidden costs of control. *American Economic Review*, 96(5), 1611–1630.** https://doi.org/10.1257/aer.96.5.1611
  *About:* Principal-agent experiment. When the principal imposed a minimum performance requirement, most agents reduced effort and explicitly reported reading the control as a signal of distrust.
  *Relation:* The canonical experimental demonstration that control backfires through a distrust inference. Our checker is a form of control imposed by the audience.

- **Ziegelmeyer, A., Schmelz, K., & Ploner, M. (2012). Hidden costs of control: Four repetitions and an extension. *Experimental Economics*, 15(2), 323–340.** https://doi.org/10.1007/s10683-011-9302-8
  *About:* Four replications confirm hidden costs of control exist but are usually smaller than the incentive effect; with flat pay the hidden costs dominate.
  *Relation:* Calibration warning. A low-stakes classroom talk (flat "pay") may show a larger distrust effect than a high-stakes board presentation.

- **Dickinson, D., & Villeval, M.-C. (2008). Does monitoring decrease work effort? The complementarity between agency and crowding-out theories. *Games and Economic Behavior*, 63(1), 56–76.** https://doi.org/10.1016/j.geb.2007.08.004
  *About:* Most agents raise effort under monitoring, but effort is crowded out above a monitoring intensity threshold, and only when the relationship is interpersonal.
  *Relation:* Suggests a dose-response: how many flags the checker raises, and how visible they are, should matter, which supports having a "correct flags" and a "false-positive flag" condition.

- **Schweitzer, M. E., & Ho, T.-H. (2005). Trust but verify: Monitoring in interdependent relationships. In *Experimental and Behavioral Economics* (Advances in Applied Microeconomics, Vol. 13). Emerald.** https://doi.org/10.1016/S0278-0984(05)13004-1
  *About:* Repeated trust games. Anticipated monitoring raised trust in the monitored rounds but lowered trust overall, and trustworthy behaviour shown under expected monitoring was discounted.
  *Relation:* Directly parallels a presenter who knows the checker is on: even a clean, flag-free talk may earn less trust credit because the audience attributes accuracy to the checker.

- **Schweitzer, M. E., Ho, T.-H., & Zhang, X. (2018). How monitoring influences trust: A tale of two faces. *Management Science*, 64(1), 253–270.** https://doi.org/10.1287/mnsc.2016.2586
  *About:* Trustees behave strategically when they know they are monitored; trustors misattribute monitored compliance to genuine trustworthiness. Its literature review maps the Strickland line of work.
  *Relation:* Pre-LLM but the most direct experimental heir of Strickland. Useful as the bridge citation between the 1958 work and our AI setting.

- **Fehr, E., & Rockenbach, B. (2003). Detrimental effects of sanctions on human altruism. *Nature*, 422, 137–140.** https://doi.org/10.1038/nature01474
  *About:* Trust game with an optional fine. When the investor imposed the fine, trustee paybacks fell, almost to zero when the sanction signalled greed; sanctions seen as fair left reciprocity intact. When the investor had the fine available but chose not to use it, paybacks were higher than with no fine at all.
  *Relation:* Using a check lowers reciprocity, but visibly renouncing an available check raises it. Suggests a Study 2 condition where the audience has the checker but the presenter is told it is switched off.

- **Bohnet, I., Frey, B. S., & Huck, S. (2001). More order with less law: On contract enforcement, trust, and crowding. *American Political Science Review*, 95(1), 131–144.** https://doi.org/10.1017/S0003055401000211
  *About:* Trustworthiness is high under weak and under strong enforcement but crowded out at intermediate enforcement.
  *Relation:* An imperfect, probabilistic checker (some errors caught, some missed, some false positives) is precisely the intermediate regime this paper predicts is worst for trust.

- **Tenbrunsel, A. E., & Messick, D. M. (1999). Sanctioning systems, decision frames, and cooperation. *Administrative Science Quarterly*, 44(4), 684–707.** https://doi.org/10.2307/2667052
  *About:* A weak surveillance and sanction system produced less cooperation than none, because it reframed the decision from an ethical one to an instrumental one.
  *Relation:* Predicts that a checker can turn presenting from a communicative act into a compliance game, which is one interpretation of "cutting numbers to avoid flags".

- **Mulder, L. B., van Dijk, E., De Cremer, D., & Wilke, H. A. M. (2006). Undermining trust and cooperation: The paradox of sanctioning systems in social dilemmas. *Journal of Experimental Social Psychology*, 42(2), 147–162.** https://doi.org/10.1016/j.jesp.2005.03.002
  *About:* People who had experienced a sanctioning system trusted others' internal motivation less once it was removed.
  *Relation:* Suggests a carry-over effect: audiences who have used a checker may trust later, unchecked presenters less. A possible follow-up measure.

- **Malhotra, D., & Murnighan, J. K. (2002). The effects of contracts on interpersonal trust. *Administrative Science Quarterly*, 47(3), 534–559.** https://doi.org/10.2307/3094850
  *About:* Binding contracts led parties to attribute the other's cooperation to the contract rather than the person, so interpersonal trust failed to develop.
  *Relation:* Closest existing analogue to "the checker, not the presenter, gets credit for accuracy".

- **Sitkin, S. B., & Roth, N. L. (1993). Explaining the limited effectiveness of legalistic "remedies" for trust/distrust. *Organization Science*, 4(3), 367–392.** https://doi.org/10.1287/orsc.4.3.367
  *About:* Distinguishes task-specific reliability from value congruence. Formal controls address reliability but can escalate distrust when they spill into value concerns.
  *Relation:* Frames our key question: does a flagged number stay a reliability issue ("got a figure wrong") or spill into integrity ("is misleading us")?

- **Kramer, R. M. (1999). Trust and distrust in organizations: Emerging perspectives, enduring questions. *Annual Review of Psychology*, 50, 569–598.** https://doi.org/10.1146/annurev.psych.50.1.569
  *About:* Review that synthesises the "surveillance breeds distrust" line of work. Companion: Kramer (1998), Paranoid cognition in social systems, *PSPR*, 2(4), 251–275, https://doi.org/10.1207/s15327957pspr0204_3, on hypervigilance in scrutinised people.
  *Relation:* The citable review to anchor the pre-AI literature in one sentence.

- **Power, M. (1997). *The Audit Society: Rituals of Verification*. Oxford University Press.**
  *About:* Argues audit institutionalises distrust, shifts effort toward what is auditable, and can erode the trust it claims to certify.
  *Relation:* Theoretical basis for the prediction that presenters optimise their slides for the checker (fewer numbers, safer claims) rather than for the audience.

- **Nagin, D. S., Rebitzer, J. B., Sanders, S., & Taylor, L. J. (2002). Monitoring, motivation, and management: The determinants of opportunistic behavior in a field experiment. *American Economic Review*, 92(4), 850–873.** https://doi.org/10.1257/00028280260344498
  *About:* Field experiment varying monitoring of call-centre workers. Many did not react to monitoring at all; heterogeneity tracked perceived fairness of treatment.
  *Relation:* Supports adding perceived fairness or legitimacy of the checker as a moderator in Study 2.

- **Fein, S. (1996). Effects of suspicion on attributional thinking and the correspondence bias. *Journal of Personality and Social Psychology*, 70(6), 1164–1184.** https://doi.org/10.1037/0022-3514.70.6.1164
  *About:* Suspicion of ulterior motives triggers effortful, sceptical attributional processing.
  *Relation:* A live checker may prime this sceptical mode in the audience for the whole talk, not just for flagged slides.

## B. Trust repair after a caught error

- **Kim, P. H., Ferrin, D. L., Cooper, C. D., & Dirks, K. T. (2004). Removing the shadow of suspicion: The effects of apology versus denial for repairing competence- versus integrity-based trust violations. *Journal of Applied Psychology*, 89(1), 104–118.** https://doi.org/10.1037/0021-9010.89.1.104
  *About:* Interaction result: apology repaired trust better than denial after a competence-based violation, whereas denial repaired trust better than apology after an integrity-based violation. Apology helped when guilt was later confirmed, denial when innocence was confirmed.
  *Relation:* Tells us to separate competence trust from integrity trust, and that the presenter's reaction to a flag (apologise, correct, deny) should be coded in Study 2. Denying a false-positive flag is predicted to repair trust; apologising for it may not.

- **Schweitzer, M. E., Hershey, J. C., & Bradlow, E. T. (2006). Promises and lies: Restoring violated trust. *Organizational Behavior and Human Decision Processes*, 101(1), 1–19.** https://doi.org/10.1016/j.obhdp.2006.05.005
  *About:* Trust damaged by untrustworthy action recovers with consistent behaviour; trust damaged by deception never fully recovers.
  *Relation:* If the audience reads a flagged slide as a lie rather than a mistake, we should expect no recovery within the session. Note-wording of the flag ("inconsistent" vs "false") may decide which reading they take.

## C. Fact-checking and the person being checked (mostly human fact-checkers)

- **Nyhan, B., & Reifler, J. (2015). The effect of fact-checking on elites: A field experiment on U.S. state legislators. *American Journal of Political Science*, 59(3), 628–640.** https://doi.org/10.1111/ajps.12162
  *About:* Legislators randomly warned that they might be fact-checked were substantially less likely to receive a negative rating or have their accuracy questioned publicly.
  *Relation:* Direct evidence that the checked party changes behaviour when it knows it is being checked. The template for Study 2's "told the audience is using the checker" manipulation.

- **Mattozzi, A., Nocito, S., & Sobbrio, F. (2022, rev. 2024). Fact-checking politicians. CEPR Discussion Paper 17710 (working paper).** https://ideas.repec.org/p/cpr/ceprdp/17710.html
  *About:* Field experiment with 55 Italian MPs. Fact-checking reduced false statements for a few weeks, but politicians substituted them with no statements or unverifiable ones.
  *Relation:* The closest published evidence for our hedging and "cutting numbers" hypothesis. Not yet peer-reviewed.

- **Wintersieck, A. L. (2017). Debating the truth: The impact of fact-checking during electoral debates. *American Politics Research*, 45(2), 304–331.** https://doi.org/10.1177/1532673X16686555
  *About:* Fact-checks shown alongside a debate: "true" verdicts raised evaluations of the candidate, "mostly false" lowered them, "half true" did nothing.
  *Relation:* Shows live fact-checks move evaluation of the speaker, not just belief in the claim, and only when the verdict is decisive. Argues for unambiguous flag wording in our prototype.

- **Fridkin, K., Kenney, P. J., & Wintersieck, A. (2015). Liar, liar, pants on fire: How fact-checking influences citizens' reactions to negative advertising. *Political Communication*, 32(1), 127–151.** https://doi.org/10.1080/10584609.2014.914613
  *About:* Fact-checks saying an ad was accurate vs inaccurate changed perceived accuracy, usefulness and tone; effects strongest among sophisticated viewers.
  *Relation:* Template for the "correct flag" vs "false-positive flag" manipulation and a reason to record audience expertise.

- **Garrett, R. K., & Weeks, B. E. (2013). The promise and peril of real-time corrections to political misperceptions. *Proc. CSCW 2013*, 1047–1058.** https://doi.org/10.1145/2441776.2441895
  *About:* Inline real-time corrections vs delayed ones. Real-time was slightly better overall, but people who liked the false claim distrusted the source more when corrected in real time.
  *Relation:* The earliest HCI precedent that inline flagging shifts trust in the source, and that the effect depends on whether the flagged claim supported the viewer's prior attitudes.

- **Nyhan, B., Porter, E., Reifler, J., & Wood, T. J. (2020). Taking fact-checks literally but not seriously? The effects of journalistic fact-checking on factual beliefs and candidate favorability. *Political Behavior*, 42, 939–960.** https://doi.org/10.1007/s11109-019-09528-x
  *About:* Fact-checks of Trump's claims corrected beliefs even among supporters but did not change attitudes toward him.
  *Relation:* The counter-case. With strong prior attachment, flags may not dent speaker evaluation. Our presenters are strangers, so the next paper is the closer analogue.

- **Prike, T., Baker, J., & Ecker, U. K. H. (2025). Fact-checking election-campaign misinformation: Impacts on noncommitted voters' feelings and behavior. *Political Psychology*.** https://doi.org/10.1111/pops.13059
  *About:* Simulated election with unfamiliar fictional candidates. Corrective fact-checks substantially reduced favourability and vote share for the corrected candidate.
  *Relation:* When the audience has no prior bond with the speaker, as with a recorded business presenter, corrections transfer to the person. Supports our expected main effect.

- **Hoes, E., Aitken, B., Zhang, J., Gackowski, T., & Wojcieszak, M. (2024). Prominent misinformation interventions reduce misperceptions but increase scepticism. *Nature Human Behaviour*.** https://doi.org/10.1038/s41562-024-01884-x
  *About:* Across three countries, fact-checking and similar interventions lowered belief in false claims but also lowered credibility of accurate information.
  *Relation:* The mechanism behind "trust lowered beyond what the errors justify": a generalised scepticism spillover onto correct slides.

- **Chae, J. H., & Tewksbury, D. (2026, online 2024). Perceiving AI intervention does not compromise the persuasive effect of fact-checking. *New Media & Society*.** https://doi.org/10.1177/14614448241286881
  *About:* AI-labelled fact-checks corrected beliefs as well as human ones; neither changed feelings toward or perceived competence of the targeted politicians.
  *Relation:* Useful null result. It suggests the speaker penalty may not depend on whether the checker is AI or human, at least for well-known targets.

## D. Automated checkers and AI-mediated communication

- **DeVerna, M. R., Yan, H. Y., Yang, K.-C., & Menczer, F. (2024). Fact-checking information from large language models can decrease headline discernment. *PNAS*, 121(50), e2322823121.** https://doi.org/10.1073/pnas.2322823121
  *About:* LLM fact-checks did not improve discernment on average; they raised belief in false headlines the LLM was unsure about and lowered belief in true headlines it mislabelled.
  *Relation:* Direct evidence that an automated checker's false positives damage correct claims. Justifies our false-positive condition.

- **von Schenk, A., Klockmann, V., Bonnefon, J.-F., Rahwan, I., & Köbis, N. (2024). Lie detection algorithms disrupt the social dynamics of accusation behavior. *iScience*, 27(7), 110201.** https://doi.org/10.1016/j.isci.2024.110201
  *About:* Without an algorithm people rarely accused others of lying. Those who requested a 67 percent accurate lie detector accused in most flagged cases, raising false accusations.
  *Relation:* An automated verifier lowers the social threshold for treating a speaker as dishonest. Closest evidence that a tool changes how humans treat each other, not just what they believe.

- **Gupta, C., Aritonang, N. V., Rajendran, D. P. D., Danry, V., Maes, P., & Nanayakkara, S. (2026). Feeling the Facts: Real-time wearable fact-checkers can use nudges to reduce user belief in false information. *Proc. CHI 2026*.** https://doi.org/10.1145/3772318.3791679
  *About:* Haptic nudges on a smartwatch while watching spoken claims improved real-time discernment, but users over-relied on the system and were more misled when it missed or wrongly flagged a claim.
  *Relation:* The closest HCI analogue to a live slide checker, including the false-positive problem. It measures belief in claims, not trust in the speaker. That is the gap Idea 2 fills.

- **Lu, Z., Li, P., Wang, W., & Yin, M. (2025). Understanding the effects of AI-based credibility indicators when people are influenced by both peers and experts. *Proc. CHI 2025*.** https://doi.org/10.1145/3706598.3713871
  *About:* An AI credibility indicator's influence was consistently larger when it disagreed with a human expert than when it agreed.
  *Relation:* AI-versus-expert disagreement is exactly the condition where the AI signal dominates. That is the risk case when our checker contradicts the presenter.

- **Jakesch, M., French, M., Ma, X., Hancock, J. T., & Naaman, M. (2019). AI-mediated communication: How the perception that profile text was written by AI affects trustworthiness. *Proc. CHI 2019*, Paper 239.** https://doi.org/10.1145/3290605.3300469
  *About:* Host profiles believed to be AI-written were rated less trustworthy when mixed with human ones.
  *Relation:* Foundational evidence that perceived AI involvement in a human's self-presentation shifts trust in the human.

- **Hohenstein, J., & Jung, M. F. (2020). AI as a moral crumple zone: The effects of AI-mediated communication on attribution and trust. *Computers in Human Behavior*, 106, 106190.** https://doi.org/10.1016/j.chb.2019.106190
  *About:* Smart replies increased interpersonal trust, and when conversations failed, blame shifted to the AI rather than the partner.
  *Relation:* A mediating AI can absorb responsibility. If the checker is framed as fallible, audiences might blame it rather than the presenter for a disputed flag.

- **Hohenstein, J., Kizilcec, R. F., DiFranzo, D., Aghajari, Z., Mieczkowski, H., Levy, K., Naaman, M., Hancock, J., & Jung, M. F. (2023). Artificial intelligence in communication impacts language and social relationships. *Scientific Reports*, 13, 5487.** https://doi.org/10.1038/s41598-023-30938-9
  *About:* Smart replies made conversations faster and more positive, but partners suspected of using AI were evaluated more negatively.
  *Relation:* Attribution of AI involvement, not the AI itself, drives the interpersonal penalty. Relevant if presenters start using AI to pre-check their own slides.

- **Jia, C., Gondimalla, A., Zhang, A., Mullings, D. J., Boltz, A., & Lee, M. K. (2026). Beyond accuracy: Experts see AI fact-checks as accurate but less useful. *Proc. CHI 2026*.** https://doi.org/10.1145/3772318.3791625
  *About:* Media experts rated LLM fact-check reports as accurate as human ones but less useful, especially when the source was undisclosed.
  *Relation:* Supports measuring trust in the tool separately from trust in the presenter, and disclosing that flags are automated.

- **Liao, M., Lee, S., Dooley, A., Xiong, A., & Sundar, S. S. (2026). When an AI says it is false: User responses to misinformation flagging by automated vs. human fact-checkers. *Media Psychology*.** https://doi.org/10.1080/15213269.2026.2659876
  *About:* 291 US participants saw headlines labelled as fact-checked by AI or by humans, with evidence-based, feature-based or no explanation. Trust in AI and human fact-checkers was equal but for different reasons (AI seen as good at spotting surface red flags, humans at corroborating evidence), and any explanation beat a black box.
  *Relation:* Supports giving each flag a short reason, as our overlay does, and is the reference to pull if we add a "human co-presenter catches the error" control condition.

## E. Observers penalise humans who use, or are contradicted by, a tool

- **Arkes, H. R., Shaffer, V. A., & Medow, M. A. (2007). Patients derogate physicians who use a computer-assisted diagnostic aid. *Medical Decision Making*, 27(2), 189–202.** https://doi.org/10.1177/0272989X06297391
  *About:* Across four vignette experiments, a physician who consulted a computerised diagnostic aid was rated less capable and professional than one who diagnosed unaided, whether the physician heeded or defied the aid.
  *Relation:* Pre-LLM precedent that the mere presence of a checking tool lowers the judged competence of the human it is attached to.

- **Shaffer, V. A., Probst, C. A., Merkle, E. C., Arkes, H. R., & Medow, M. A. (2013). Why do patients derogate physicians who use a computer-based diagnostic support system? *Medical Decision Making*, 33(1), 108–118.** https://doi.org/10.1177/0272989X12453501
  *About:* The unaided physician was rated better than one using a computer aid, but not better than one consulting a human colleague. Attitudes toward statistics did not moderate the penalty; a stronger internal health locus of control did.
  *Relation:* Argues for a control condition where a human audience member catches the same error. Audience numeracy is still worth measuring, but this paper does not support it as a moderator.

- **Palmeira, M., & Spassova, G. (2015). Consumer reactions to professionals who use decision aids. *European Journal of Marketing*, 49(3/4), 302–326.** https://doi.org/10.1108/EJM-07-2013-0390
  *About:* Professionals were penalised only when the decision aid replaced their judgement; a hybrid of expertise plus aid was not penalised.
  *Relation:* Framing the checker as augmenting rather than replacing the presenter may blunt the trust penalty.

- **Palmeira, M., & Romero Lopez, M. (2023). The opposing impacts of advice use on perceptions of competence. *Journal of Behavioral Decision Making*, 36(4), e2318.** https://doi.org/10.1002/bdm.2318
  *About:* Moderate use of advice raises perceived competence, full deference lowers it.
  *Relation:* Predicts different audience reactions to a presenter who accepts every flag versus one who pushes back on a false positive.

- **Reif, J. A., Larrick, R. P., & Soll, J. B. (2025). Evidence of a social evaluation penalty for using AI. *PNAS*, 122(19), e2426766122.** https://doi.org/10.1073/pnas.2426766122
  *About:* Four preregistered experiments. People who use AI at work are judged lazier and less competent, anticipate this, and adjust behaviour.
  *Relation:* Nearest modern analogue for both halves of Idea 2: an evaluation penalty tied to AI, and anticipatory self-censorship by the evaluated person.

- **Chen, C., Sun, Y., Liao, M., & Sundar, S. S. (2026). When AI disagrees: The effect of second opinion on patients' trust in doctors. *International Journal of Human-Computer Studies*, 213, 103824.** https://doi.org/10.1016/j.ijhcs.2026.103824
  *About:* A simulated doctor consults an AI second opinion. AI agreement raised credibility; AI disagreement raised perceived uncertainty and perceived laziness, lowering trust in the doctor.
  *Relation:* The closest published design to "automated checker contradicts the human in front of the audience". Idea 2 moves this from a dyad to a presenter-audience setting with a persistent visual artifact.

- **Detjen, H., Densky, L., von Kalckreuth, N., & Kopka, M. (2025). Who is trusted for a second opinion? Comparing collective advice from a medical AI and physicians in biopsy decisions after mammography screening. *Proc. CHI 2025*.** https://doi.org/10.1145/3706598.3713898
  *About:* Participants said they trusted physicians more, yet under conflict many followed the AI.
  *Relation:* Stated trust and behavioural deference can diverge. We should measure both (trust ratings and whether listeners act on the flag).

- **Dietvorst, B. J., Simmons, J. P., & Massey, C. (2015). Algorithm aversion: People erroneously avoid algorithms after seeing them err. *Journal of Experimental Psychology: General*, 144(1), 114–126.** https://doi.org/10.1037/xge0000033
  *About:* Confidence in an algorithm fell faster than in a human after both made the same error.
  *Relation:* The flip side of our false-positive condition: a checker that visibly false-flags may lose the audience's trust faster than the presenter does.

- **Logg, J. M., Minson, J. A., & Moore, D. A. (2019). Algorithm appreciation: People prefer algorithmic to human judgment. *Organizational Behavior and Human Decision Processes*, 151, 90–103.** https://doi.org/10.1016/j.obhdp.2018.12.005
  *About:* Lay people weighted advice more when told it came from an algorithm; experts discounted it.
  *Relation:* Background for expecting non-expert audiences (board members, students) to side with the checker over the presenter.

- **Longoni, C., Bonezzi, A., & Morewedge, C. K. (2019). Resistance to medical artificial intelligence. *Journal of Consumer Research*, 46(4), 629–650.** https://doi.org/10.1093/jcr/ucz013
  *About:* Consumers avoid AI providers because they feel AI neglects their uniqueness.
  *Relation:* Background only. It is about choosing AI versus a human, not about judging a human who is checked by AI.

## F. Visible errors, numbers, charts and hedging

- **Kreiner, D. S., Schnakenberg, S. D., Green, A. G., Costello, M. J., & McClin, A. F. (2002). Effects of spelling errors on the perception of writers. *Journal of General Psychology*, 129(1), 5–17.** https://doi.org/10.1080/00221300209602029
  *About:* More spelling errors lowered ratings of the writer's ability in a dose-dependent way, spilling into judgements of general ability at high counts.
  *Relation:* Baseline that surface errors leak into competence judgements, and a dose-response we can mirror with the number of planted errors.

- **Vignovic, J. A., & Thompson, L. F. (2010). Computer-mediated cross-cultural collaboration: Attributing communication errors to the person versus the situation. *Journal of Applied Psychology*, 95(2), 265–276.** https://doi.org/10.1037/a0018628
  *About:* Emails with errors lowered ratings of the sender's intelligence and trustworthiness; situational information removed the penalty for grammar errors.
  *Relation:* Whether the audience attributes a flagged slide error to the presenter or to circumstances (OCR misread, stale data) should moderate the trust hit. A flag that names a plausible cause may soften it.

- **Appelman, A., & Bolls, P. (2011). Article recall, credibility lower with grammar errors. *Newspaper Research Journal*, 32(2), 50–62.** https://doi.org/10.1177/073953291103200205
  *About:* Ten errors in a news story lowered perceived credibility and recall.
  *Relation:* Media-credibility analogue for error visibility; the later 2017 replication needed a larger error dose, so two planted errors per talk may be at the low end.

- **Appelman, A. (2024). Numbers in news articles: Effects of presence, errors, and (false) recall. *Electronic News*, 18(2), 67–82.** https://doi.org/10.1177/19312431231190846
  *About:* Readers who noticed, or only believed they noticed, mathematical errors rated both content and source lower.
  *Relation:* The most direct evidence that numeric errors damage source credibility and that the mere belief that an error occurred is enough. That belief is exactly what a false-positive flag creates.

- **Clark, J. L., Green, M. C., & Simons, J. J. P. (2019). Narrative warmth and quantitative competence: Message type affects impressions of a speaker. *PLoS ONE*, 14(12), e0226713.** https://doi.org/10.1371/journal.pone.0226713
  *About:* Statistical arguments raise perceived speaker competence, narratives raise warmth.
  *Relation:* Presenters who cut numbers to avoid flags may pay a competence cost. Worth coding "numbers removed" in Study 2.

- **Pandey, A. V., Rall, K., Satterthwaite, M. L., Nov, O., & Bertini, E. (2015). How deceptive are deceptive visualizations? An empirical analysis of common distortion techniques. *Proc. CHI 2015*, 1469–1478.** https://doi.org/10.1145/2702123.2702608
  *About:* Truncated axes, area-as-quantity and similar distortions changed viewers' takeaways.
  *Relation:* Establishes that chart errors change beliefs. It does not measure trust in the presenter, so that is open for us.

- **Lo, L. Y.-H., Gupta, A., Shigyo, K., Wu, A., Bertini, E., & Qu, H. (2022). Misinformed by visualization: What do we learn from misinformative visualizations? *Computer Graphics Forum*, 41(3), 515–525.** https://doi.org/10.1111/cgf.14559
  *About:* Taxonomy of 74 issue types from over 1,000 real charts reported as misleading; many are numeric or label inconsistencies.
  *Relation:* Ready-made taxonomy for which slide errors an OCR/LLM checker should flag, and for choosing our planted errors.

- **Wang, H. W., Lin, K., Cohen, A., Kennedy, R., Zwald, Z., Nobre, C., & Xiong Bearfield, C. (2026). Do you "trust" this visualization? An inventory to measure trust in visualizations. *IEEE TVCG*.** https://doi.org/10.1109/TVCG.2025.3646847 (see also Elhamdadi et al. 2023, Vistrust, https://doi.org/10.1109/TVCG.2023.3326579)
  *About:* Validated short inventory for trust in a visualization.
  *Relation:* A ready dependent measure for trust in the flagged chart, separate from trust in the presenter and in the tool.

- **Jensen, J. D. (2008). Scientific uncertainty in news coverage of cancer research: Effects of hedging on scientists' and journalists' credibility. *Human Communication Research*, 34(3), 347–369.** https://doi.org/10.1111/j.1468-2958.2008.00324.x
  *About:* Hedged reports increased perceived trustworthiness of scientists, especially when the hedging came from the scientists themselves.
  *Relation:* Counterpoint to our assumption that hedging under checker scrutiny is costly. Hedging may help trust if it reads as honesty.

- **Durik, A. M., Britt, M. A., Reynolds, R., & Storey, J. (2008). The effects of hedges in persuasive arguments: A nuanced analysis of language. *Journal of Language and Social Psychology*, 27(3), 217–234.** https://doi.org/10.1177/0261927X08317947
  *About:* Hedges on data statements lowered evaluations of the source, most strongly among readers with scientific training; colloquial hedges hurt more than professional ones.
  *Relation:* Predicts that a presenter who starts saying "roughly" on numbers to dodge flags pays a credibility cost with numerate audiences. Gives us a coding scheme for hedge type.

## G. Detection and monitoring tools and the checker-checked relationship

- **Introna, L. D., & Hayes, N. (2011). On sociomaterial imbrications: What plagiarism detection systems reveal and why it matters. *Information and Organization*, 21(2), 107–122.** https://doi.org/10.1016/j.infoandorg.2011.03.001
  *About:* Turnitin detects copied-text patterns, not plagiarism; delegating judgement to it labels some students, especially non-native writers, as plagiarists who are not.
  *Relation:* An OCR/LLM checker also detects a narrow proxy of "error". Audiences may read its flags as judgements of the presenter.

- **Introna, L. D. (2016). Algorithms, governance, and governmentality: On governing academic writing. *Science, Technology, & Human Values*, 41(1), 17–49.** https://doi.org/10.1177/0162243915587360
  *About:* Turnitin reshapes what "original writing" means and how students self-govern to fit the detector.
  *Relation:* Supports the hypothesis that checked presenters restructure content around what the checker can score.

- **Zwagerman, S. (2008). The Scarlet P: Plagiarism, panopticism, and the rhetoric of academic integrity. *College Composition and Communication*, 59(4), 676–710.** https://eric.ed.gov/?id=EJ802718
  *About:* Argues detection technology installs a culture of suspicion that damages the teacher-student relationship regardless of actual cheating rates.
  *Relation:* Source for the "trust penalty beyond what the errors justify" framing in a pre-LLM educational setting.

- **Liang, W., Yuksekgonul, M., Mao, Y., Wu, E., & Zou, J. (2023). GPT detectors are biased against non-native English writers. *Patterns*, 4(7), 100779.** https://doi.org/10.1016/j.patter.2023.100779
  *About:* Short opinion piece. AI-text detectors flagged most TOEFL essays as AI-written while sparing native-speaker essays.
  *Relation:* Automated checkers produce false positives correlated with attributes of the checked person, so the trust cost may fall unevenly across presenters.

- **Gorichanaz, T. (2023). Accused: How students respond to allegations of using ChatGPT on assessments. *Learning: Research and Practice*, 9(2), 183–196.** https://doi.org/10.1080/23735082.2023.2254787
  *About:* Thematic analysis of student accounts of being accused by AI detectors; a core theme is teachers trusting the detector over the student.
  *Relation:* Closest empirical account of a checker's verdict displacing interpersonal trust.

- **Ravid, D. M., Tomczak, D. L., White, J. C., & Behrend, T. S. (2020). EPM 20/20: A review, framework, and research agenda for electronic performance monitoring. *Journal of Management*, 46(1), 100–126.** https://doi.org/10.1177/0149206319869435
  *About:* Framework organising monitoring effects by purpose, transparency, invasiveness and control.
  *Relation:* Gives the design space for our checker: announced vs covert, developmental vs evaluative framing.

- **Ravid, D. M., White, J. C., Tomczak, D. L., Miles, A. F., & Behrend, T. S. (2023). A meta-analysis of the effects of electronic performance monitoring on work outcomes. *Personnel Psychology*, 76(1), 5–40.** https://doi.org/10.1111/peps.12514
  *About:* Meta-analysis of 94 studies. No evidence monitoring improves performance; it raises stress regardless of design; transparent, less invasive monitoring yields better attitudes.
  *Relation:* Baseline for "being checked" having a cost even when no error is flagged.

- **Siegel, R., König, C. J., & Lazar, V. (2022). The impact of electronic monitoring on employees' job satisfaction, stress, performance, and counterproductive work behavior: A meta-analysis. *Computers in Human Behavior Reports*, 8, 100227.** https://doi.org/10.1016/j.chbr.2022.100227
  *About:* Independent meta-analysis: monitoring slightly lowers satisfaction, raises stress and counterproductive behaviour, no performance gain.
  *Relation:* Replicates the previous entry from a different team; useful as a pair.

- **Thiel, C. E., Bonner, J., Bush, J. T., Welsh, D. T., & Garud, N. (2023). Stripped of agency: The paradoxical effect of employee monitoring on deviance. *Journal of Management*, 49(2), 709–740.** https://doi.org/10.1177/01492063211053224
  *About:* Monitored people broke more rules, via a reduced sense of agency and moral responsibility; perceived fairness buffered this.
  *Relation:* Predicts an alternative presenter response: offloading responsibility for accuracy onto the checker rather than preparing more.

- **Alge, B. J. (2001). Effects of computer surveillance on perceptions of privacy and procedural justice. *Journal of Applied Psychology*, 86(4), 797–804.** https://doi.org/10.1037/0021-9010.86.4.797
  *About:* Monitoring only job-relevant activity and giving the monitored party input reduced perceived invasion and raised procedural justice.
  *Relation:* Design levers: restrict the checker to verifiable facts and let presenters respond to flags.

- **Bernstein, E. S. (2012). The transparency paradox: A role for privacy in organizational learning and operational control. *Administrative Science Quarterly*, 57(2), 181–216.** https://doi.org/10.1177/0001839212453028
  *About:* Field experiment in a factory: observability led workers to conceal activities and suppress productive deviation.
  *Relation:* Supports the prediction that observed presenters cut risky-but-informative content such as estimates.

- **Kellogg, K. C., Valentine, M. A., & Christin, A. (2020). Algorithms at work: The new contested terrain of control. *Academy of Management Annals*, 14(1), 366–410.** https://doi.org/10.5465/annals.2018.0174
  *About:* Synthesises algorithmic control and worker responses such as gaming and obfuscation.
  *Relation:* Vocabulary for presenter counter-strategies (hedging, formatting numbers so the OCR cannot parse them).

- **Wessel, M., Serebrenik, A., Wiese, I., Steinmacher, I., & Gerosa, M. A. (2022). Quality gatekeepers: Investigating the effects of code review bots on pull request activities. *Empirical Software Engineering*, 27(5), 108.** https://doi.org/10.1007/s10664-022-10130-9
  *About:* Adopting automated review bots increased merged pull requests but reduced human communication between maintainers and contributors. Companion: Wessel et al. (2021), "Don't disturb me", *PACM HCI* 5(CSCW2), https://doi.org/10.1145/3476042.
  *Relation:* Evidence that inserting a checking tool between author and reviewer restructures the human relationship, in a pre-LLM software setting.

- **Penney, J. W. (2016). Chilling effects: Online surveillance and Wikipedia use. *Berkeley Technology Law Journal*, 31(1), 117–182.** https://btlj.org/data/articles2016/vol31/31_1/0117_0182_Penney_ChillingEffects_WEB.pdf
  *About:* Sustained drop in views of privacy-sensitive Wikipedia articles after the 2013 surveillance revelations.
  *Relation:* Canonical empirical chilling-effect study for the "presenters avoid checkable content" prediction.

- **Marder, B., Joinson, A., Shankar, A., & Houghton, D. (2016). The extended 'chilling' effect of Facebook: The cold reality of ubiquitous social networking. *Computers in Human Behavior*, 60, 582–592.** https://doi.org/10.1016/j.chb.2016.02.097
  *About:* Anticipated peer surveillance constrains self-presentation even offline.
  *Relation:* Relevant because in Idea 2 the audience, not a manager, is the monitor.

- **Mak, M. H. C., & Walasek, L. (2025). Style, sentiment, and quality of undergraduate writing in the AI era. *Computers and Education: Artificial Intelligence*, 9, 100507.** https://doi.org/10.1016/j.caeai.2025.100507
  *About:* ChatGPT-marker words in student reports rose in 2023-24 then fell in 2025, read as students masking AI traces.
  *Relation:* Longitudinal evidence of the checked party altering output to evade a detector.

## H. Lower-priority items found along the way (verified DOIs, one line each)

- Lee, Kusbit, Metsky & Dabbish (2015), Working with machines, CHI, https://doi.org/10.1145/2702123.2702548. Algorithmic management and driver trust.
- Stoycheff (2016), Under surveillance, *J&MCQ* 93(2), https://doi.org/10.1177/1077699016630255. Self-censorship under perceived surveillance.
- Jeske & Santuzzi (2015), *New Technology, Work and Employment* 30(1), https://doi.org/10.1111/ntwe.12039. Continuous monitoring lowers self-efficacy; unpredictable, individual-focused monitoring lowers attitudes and citizenship behaviour.
- Bacchelli & Bird (2013), Expectations, outcomes, and challenges of modern code review, ICSE, https://doi.org/10.1109/ICSE.2013.6606617.
- Mann, Nolan & Wellman (2003), Sousveillance, *Surveillance & Society* 1(3), https://doi.org/10.24908/ss.v1i3.3344. Bottom-up watching provokes counter-performances by the watched; nearest concept to "audience holds the checker".
- Appling, Bruckman & De Choudhury (2022), Reactions to fact checking, *PACM HCI* (CSCW), https://doi.org/10.1145/3555128. Qualitative, sharers' reactions.
- LiveFC (WSDM 2025), https://doi.org/10.1145/3701551.3704128. Live audio-stream fact-checking system, no audience-trust measures.
- Kramer (1998), Paranoid cognition in social systems, *PSPR* 2(4), https://doi.org/10.1207/s15327957pspr0204_3.

## Open flags from the search

- Kruglanski (1970) and Strickland (1958): citations confirmed via Crossref, findings summarised from secondary sources (Schweitzer, Ho & Zhang 2018; Kramer 1999; Mayer et al. 1995) because the originals are paywalled. Kruglanski's worker-side response is unconfirmed.
- Mattozzi, Nocito & Sobbrio is a working paper, not peer-reviewed.
- Liao et al. (2026): metadata only, abstract not retrieved.
- Zwagerman (2008) has no DOI; cite via ERIC or JSTOR.
- Detjen et al. (2025): "many followed the AI" holds only when the AI recommended biopsy; otherwise most followed the physician.
- Validator corrections applied on 2026-09-25: Wessel et al. 2022 author list (the search had attached the authors of a different Wessel paper), Hohenstein et al. 2023 author order, Palmeira & Romero Lopez issue number, Fehr & Rockenbach and Kim et al. findings, Shaffer et al. moderator, Kruglanski worker-side claim.
