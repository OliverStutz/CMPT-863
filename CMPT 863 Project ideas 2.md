
We have during a meeting brainstormed ideas and came up with an idea graph so that we can effectively generate ideas:
![[illustration-excalidraw.png]]
https://excalidraw.com/\#room=935957a5e12e4a62ce65,rHIScu37UqFN1NNFlFmyBQ

## 1. Fact-Checked in the Room: How Real-Time AR/LLM Fact-Checking Affects Human-to-Human Trust

### Core contribution

Tools that fact-check another person's speech in real time are now feasible: smart glasses with live transcription or capture can send claims to an LLM and show a verdict while the speaker is still talking. We do not know whether they help. This project would show how accurate short, out-of-context LLM fact-checks are on realistic meeting claims. It would also show whether knowing that such checking is happening, and challenging a colleague based on its feedback lowers/increases trust between people. That would give designers evidence on whether real-time fact-checking belongs in meetings and boardrooms, and in what form.

### Study description

A speaker's speech and prensentation are transcribed live and sent to an LLM, which returns a verdict with a confidence cue on AR glasses or, as a fallback, a tablet (for this project maybe a reduced prototype is best). In a board room simulation we would role-play a short board-style briefing on a scripted cybersecurity risk report, inspired by the seed paper's boardroom setting. In each round one participant presents and the other questions. The script plants a mix of true, false, and context-dependent claims so we can compute the fact-checker's precision and recall and see where missing context makes it fail. We would use a within-subjects design with counterbalanced conditions: (a) no fact-checking, (b) covert fact-checking that only the questioner knows about, and (c) disclosed fact-checking that both people know about. After each round we measure interpersonal trust with a validated scale, increase and deterioration of trust and confidence in board members and presenter.

### Related academic papers

- Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From Oversight to Insight: Transforming Cybersecurity Governance in Boardrooms. CHI 2026. https://doi.org/10.1145/3772318.3791142 — Seed paper. Found that board directors often lack the literacy to question cyber-risk reports. That motivates both the boardroom scenario and the question of whether AI fact-checks help directors challenge presenters or instead damage trust between board members and CISOs.

## 2. Checked by the Camera: How Vision-Based Fact-Checking of Slides Shapes Trust Between Presenter and Audience

### Core contribution

AI fact-checking research mostly studies how people judge the content being checked. It says little about what happens to the relationship between the people involved, especially when the checked material is a visual artifact (a slide, chart, or shared document) rather than speech. This project asks whether a camera-based system that flags wrong or inconsistent figures on a presenter's slides lowers the audience's trust in the presenter beyond what the actual errors justify, and whether presenters who know they are being checked change how they present, for example by hedging more, cutting numbers, or doing extra verification beforehand. Unlike idea 1 (AR + LLM checking of spoken claims), this idea looks only at visual, persistent artifacts. 

### Study description

A camera attached to any device pointed at a projected slide deck runs OCR and chart parsing, and a small overlay on the audience member's device highlights a flagged number or chart with a short note (e.g. "percentages sum to 112%", "contradicts slide 3"). Flags would be scripted so that accuracy is controlled; a live llm pipe is optional. The study has N participants watch three short recorded business presentations, each with two planted errors, in three conditions: no checker, checker with correct flags, and checker with one false-positive flag. After each talk we measure trust in the presenter, trust in the tool, and error detection accuracy by the listener. Study 2 has participants give a 5-minute talk from a prepared deck, told either that the audience is or is not using the checker, and we code number hedging, self-corrections, and preparation time, followed by short interviews. Analysis uses linear mixed models with participant as a random effect for Study 1, simple group comparisons plus thematic analysis for Study 2.

### Related academic papers

- Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From Oversight to Insight: Transforming Cybersecurity Governance in Boardrooms. CHI 2026. https://doi.org/10.1145/3772318.3791142 — Seed paper. Its interviews show board directors lack the confidence to interrogate the risk reports presented to them, which is exactly the kind of setting where an automated slide checker could help audiences but could also erode trust between the board and the presenting CISO.

## 3. When the AI Corrects the Professor: How Live AI Fact-Checking of Lecture Material Shapes Student Trust in Instructors

### Core contribution

Students increasingly point camera or screen-capture LLM tools at lecture slides to check and explain them, but HCI work on AI fact-checking and reliance has studied news headlines and decision tasks, not settings where the AI openly contradicts a human authority the user is supposed to learn from. We would show how AI flags on a teacher's material, both correct flags and false positives, shift students' trust in the instructor versus the AI, whether trust in the instructor recovers after a false flag is exposed, and what this does to learning. The result gives design guidance for "second opinion" tools in classrooms that catch real errors without eroding legitimate expertise.

### Study description

We would build a lightweight prototype that takes screenshots of lecture slides (plus a transcript of the spoken explanation), sends them to a vision-capable LLM, and shows side-panel flags such as "this claim may be outdated" with a short explanation and confidence label (maybe using Jev-TypeSafe). In a lecture study, N non-expert university students get equiped with the tool. We intentionally check the facts, always claim that the facts are ok, or inject continues corrections based on hallucinations to study the trust the student has in AI tools. At the end of the study we reveal what mode the class was in and study the prior collected results given by the students of the trust the have in the lecturer vs. the LLM.

### Related academic papers


## 4. Understanding or the Illusion of It: LLM Explanations for Non-Experts Making Technical Decisions

### Core contribution

Board members, managers, and students increasingly use LLMs to get up to speed on technical topics they must then decide about, but we do not know whether these explanations produce real understanding or only a fluent illusion of understanding that inflates confidence. Prior work on AI explanations shows they raise acceptance of AI advice regardless of correctness, yet it studies explanations of a model's prediction, not explanations meant to teach a non-expert a domain well enough to judge it themselves. This project would measure the gap between perceived and actual comprehension after LLM explanations. It would also test whether interaction designs that make the learner do some of the thinking.

### Study description

We would build a web-based explainer that uses an LLM to present a cybersecurity governance scenario (e.g., a ransomware risk report asking the board to choose between three mitigation investments). It comes in three conditions: (1) a plain LLM explanation with free-form chat, (2) the same content plus cognitive forcing (the participant must commit to an initial judgment and write a one-sentence self-explanation before each section is revealed), and (3) a no-AI baseline with a static briefing document of matched length. We would run a between-subjects study with about 30 non-expert participants (10 per condition), recruited from business, MBA, or non-CS graduate students as proxies for board directors. Before and after, participants rate their own understanding of the topic, following the illusion-of-explanatory-depth procedure. They then write a causal explanation of how the threat affects the organization, answer a 10-item comprehension quiz, and make the investment decision with a confidence rating. Two cyber-security experts score decision quality against a rubric. The main analysis compares calibration (self-rated minus measured comprehension), quiz accuracy, and decision quality across conditions with statistical tests, plus short follow-up interviews coded thematically on where participants felt they "got it right" 

### Related academic papers

- Aamir, T., Psaroulis, G., Grobler, M., & Janicke, H. (2026). From Oversight to Insight: Transforming Cybersecurity Governance in Boardrooms. CHI 2026. https://doi.org/10.1145/3772318.3791142 — Seed paper. Interviews with 13 Australian board directors show a large gap in cyber knowledge and confidence that undermines their ability to interrogate risk reports, which is exactly the decision context our explainer targets.

## 5. Interactive PDF reader/annotation with LLM insights and web search results

### Core contribution

Augmented readers such as ScholarPhi and Paper Plain show that in-situ support helps people read papers and AnnotateGPT shows that pen marks can act as intent signals for an LLM. What we do not know is whether letting a reader's own highlights and freehand marks trigger AI help improves comprehension and critical reading compared to the now-standard setup of a separate chat/browser window. We want to study how the placement and trigger of AI helps (in the reader's ink vs a side conversation) and shape the reader's understanding.

### Study description

We’ll conduct a study with N participants using a probe where pen markings (highlights, circle, underline, etc.) are sent to an LLM along with their surrounding text. LLM infers the purpose of the marking and returns relevant responses or web search results, allowing the user to look up information in-context. The task is to read and comprehend (as much as possible) a research paper/poster within 30 minutes. We’ll divide the participants into 2 groups, the first will use our proposed system and the second will use a traditional PDF reader + a browser window with access to web search and an LLM of their choice. Afterwards participants are to complete a short quiz to test their understanding of the paper and a semi-structured interview with us.

### Related academic papers

- AnnotateGPT https://doi.org/10.1145/3772318.3790867. This is where I got the idea from. The paper investigates specifically pen gestures as signals for an LLM for essay markups, but I think the same concept can be applied to any PDFs not just essays.

The below 2 papers show that in-context support made papers easier to read with no loss in comprehension compared to traditional PDF tools

- Paper Plain https://doi.org/10.1145/3589955
- ScholarPhi https://doi.org/10.1145/3411764.3445648


## 6. Surveillance or Support? How Parent-Centered vs. Teen-Centered Social Media Risk Detection Shapes Parent-Teen Relationships

### Core contribution

Parental monitoring apps are usually evaluated on whether they catch risks, while their effect on the parent-teen relationship is inferred from one-sided interviews with either parents or teens. Recent work proposes "teen-centered" AI risk detection that alerts the teen first, but experts question whether it works in practice, and no study has compared it with parent-first alerting from both sides of the same family. This project would show how alert routing (who is notified first, and what the parent sees) changes trust, perceived autonomy, privacy boundaries, and willingness to disclose online risks, measured in matched parent-teen dyads.

### Study description

We would create AI-Personas of parents and their children and run a coded two-week diary study with a scenario-based within-subjects design. Each day, both members receive the same realistic risk scenario (e.g., a stranger's DM, cyberbullying, a request for explicit images) shown through alerts: parent-first surveillance (the parent sees the flagged message and an alert) and teen-first detection (the teen is warned, gets coping options, and chooses whether and when to involve the parent, as in the MOSafely dashboard). Each diary entry records anticipated trust, perceived privacy intrusion, perceived autonomy support, and likelihood of telling the other person, using short items adapted from the Parental Monitoring/Child Disclosure scales (Kerr & Stattin) and the Perceived Parental Autonomy Support Scale. At the start, during each day, after each event or non-event we will ask how the child/parent feelings are at this point of time. We would analyze paired quantitative ratings with signed-rank tests on condition and parent-teen gaps, and do reflexive thematic analysis of the diaries and interviews focused on where parent-persona and teen-persona views differ.

The study would use AI-personas specifically for the sensitivity of actual study on teenagers as well as the timeline given and REB board issues which might arise. The original idea has specifically been substituted in this scenario.

### Related academic papers

- Ma, R., Alsoubai, A., Park, J. K., & Wisniewski, P. J. (2026). From "Fail Fast" to "Mature Safely": Expert Perspectives as Secondary Stakeholders on Teen-Centered Social Media Risk Detection. CHI 2026. https://doi.org/10.1145/3772318.3791498 — Seed paper: 33 online safety experts reviewed the teen-centered MOSafely dashboard and named tensions such as teen empowerment vs. motivation and informing vs. intervening; our study tests those tensions with the families themselves.


### Oliver removed 7. as it has become a duplicate after rework of the sixth idea, the 6th idea was difficult because of timeline and the REB board

## 8. Simulating Families Under Surveillance: A Multi-Agent LLM Testbed for How Teen Monitoring Designs Shape Parent-Teen Relationships Over Time

### Core contribution

Teen monitoring tools are evaluated after they are deployed, which the seed paper calls "fail fast", or in one-time interviews, and neither shows how a given design changes parent-teen trust, disclosure and conflict over weeks. We would contribute an open multi-agent testbed in which LLM parent and teen agents live through simulated weeks under different monitoring designs, which lets designers stress-test those designs before real families are exposed. The HCI community would learn which relationship dynamics generative-agent simulation reproduces in line with published human findings, which it misses or flattens, and therefore whether it can serve as a cheap pre-deployment check for family safety technology.

### Study description

We would adapt the Generative Agents architecture (memory stream, reflection, planning) into a sandbox of about 12 simulated families, each with one parent agent and one teen agent (ages 13-17) whose personas vary parenting style, teen risk exposure and baseline relationship quality. Each simulated day, the teen agent uses a social media feed with scripted risk events at a fixed rate (a message from a stranger, cyberbullying, explicit content), and the parent and teen then have an end-of-day conversation. The design is within-persona: every family persona runs under all four monitoring conditions, with 3 random seeds per run and 8 simulated weeks per run. The conditions are (1) no monitoring, (2) covert full monitoring where the parent reads everything, (3) transparent risk alerts sent to the teen first with an option to escalate, and (4) risk alerts sent straight to the parent. Each week, both agents answer adapted items from the Kerr and Stattin parental monitoring and child disclosure scales and a trust scale, and we code transcripts for disclosure events, conflict episodes, moments when the teen discovers covert monitoring, and the teen's risk-coping actions. We would fit mixed-effects growth models (condition by week, random intercepts for family persona and seed) and check whether the trajectories match published human findings, for example that teens see covert, restrictive monitoring as harmful (Ghosh et al. 2018). Finally, we would rate sample transcripts for believability, and we would write a structured account of what the simulation can and cannot tell designers.

### Related academic papers

- Ma, R., Alsoubai, A., Park, J. K., & Wisniewski, P. J. (2026). From "Fail Fast" to "Mature Safely": Expert Perspectives as Secondary Stakeholders on Teen-Centered Social Media Risk Detection. CHI 2026. https://doi.org/10.1145/3772318.3791498 (also arXiv:2601.13516) — The seed paper: its interview study with 33 experts on a teen-centered risk dashboard found tensions, such as informing about risks vs. intervening meaningfully. Our monitoring conditions put those tensions into practice, and the simulation is one way to "mature safely" before a tool reaches real families.

### Core contribution

HCI researchers are starting to use LLM personas to pilot studies, generate interview data, and simulate users (as in idea 8), but there is little evidence on whether persona results agree with what real participants said in a specific HCI study. We would replicate one published, small-N HCI study with LLM personas and measure agreement with the original human results along four dimensions: effect direction, theme overlap, response distribution, and response diversity. The HCI community would learn when persona data is good enough to pilot a study, when it misleads, and which biases cause the gaps.

### Study description

We would pick one published HCI study with a fully reported protocol and results, ideally a vignette or semi-structured interview study on parental monitoring of teens, so the findings carry over directly to idea 8. We would build personas that match the original sample's reported demographics (for example, 20-30 parent and teen personas) and run the original protocol on 2-3 LLMs under three prompting conditions: bare persona, persona plus demographic backstory, and persona plus a de-biasing instruction that asks for disagreement and variance, with 5 repeated runs per persona to measure stability. For closed-ended items we would compare the direction and size of effects and the response distributions against the human data. For open-ended responses, two coders would apply the original paper's codebook to the synthetic transcripts, and we would report theme recall and precision against the published themes, Cohen's kappa between coders, and themes the LLMs invented. We would measure homogenization as lexical and semantic diversity (embedding dispersion, distinct-n) against the original quotes, and we would measure biases by running reworded versions of the questions that lean toward one answer and checking how far responses shift. 

### Related academic papers

- Ma, R., Alsoubai, A., Park, J. K., & Wisniewski, P. J. (2026). From "Fail Fast" to "Mature Safely": Expert Perspectives as Secondary Stakeholders on Teen-Centered Social Media Risk Detection. CHI 2026. https://doi.org/10.1145/3772318.3791498 (also arXiv:2601.13516) — The seed paper: its interview study with 33 experts on a teen-centered risk dashboard found tensions, such as informing about risks vs. intervening meaningfully. Our monitoring conditions put those tensions into practice, and the simulation is one way to "mature safely" before a tool reaches real families.
- Hämäläinen, P., Tavast, M., & Kunnari, A. (2023). Evaluating Large Language Models in Generating Synthetic HCI Research Data: a Case Study. Proceedings of the 2023 CHI Conference on Human Factors in Computing Systems (CHI '23). https://doi.org/10.1145/3544548.3580688 — Showed that GPT-3 can produce believable open-ended answers about experiencing video games as art, and that these answers are less diverse than real ones. We extend this from a single case study to a systematic comparison against a published study's themes and effects.

## 9. How Well Do LLM Personas Match Real Participants? Replicating a Published HCI Study with Silicon Samples

### Core contribution

HCI researchers are starting to use LLM personas to pilot studies, generate interview data, and simulate users (as in idea 8), but there is little evidence on whether persona results agree with what real participants said in a specific HCI study. We would replicate one published, small-N HCI study with LLM personas and measure agreement with the original human results along four dimensions: effect direction, theme overlap, response distribution, and response diversity. The HCI community would learn when persona data is good enough to pilot a study, when it misleads, and which biases cause the gaps.

### Study description

We would pick one published HCI study with a fully reported protocol and results, ideally a vignette or semi-structured interview study on parental monitoring of teens, so the findings carry over directly to idea 8. We would build personas that match the original sample's reported demographics (for example, 20-30 parent and teen personas) and run the original protocol on 2-3 LLMs under three prompting conditions: bare persona, persona plus demographic backstory, and persona plus a de-biasing instruction that asks for disagreement and variance, with 5 repeated runs per persona to measure stability. For closed-ended items we would compare the direction and size of effects and the response distributions against the human data. For open-ended responses, two coders would apply the original paper's codebook to the synthetic transcripts, and we would report theme recall and precision against the published themes, Cohen's kappa between coders, and themes the LLMs invented. We would measure homogenization as lexical and semantic diversity (embedding dispersion, distinct-n) against the original quotes, and we would measure biases by running reworded versions of the questions that lean toward one answer and checking how far responses shift.

### Related academic papers

- Hämäläinen, P., Tavast, M., & Kunnari, A. (2023). Evaluating Large Language Models in Generating Synthetic HCI Research Data: a Case Study. Proceedings of the 2023 CHI Conference on Human Factors in Computing Systems (CHI '23). [https://doi.org/10.1145/3544548.3580688](https://doi.org/10.1145/3544548.3580688) — Showed that GPT-3 can produce believable open-ended answers about experiencing video games as art, and that these answers are less diverse than real ones. We extend this from a single case study to a systematic comparison against a published study's themes and effects.

## 10. Auditing AI Discrimination Across User Groups in LLM-Based Family Social-Media Monitoring

### Core contribution

LLMs are starting to be used in two roles in parental monitoring. They simulate parent and teen personas for design research (idea 8), and they judge which teen posts or messages are "risky." Nobody has checked whether either role treats demographic groups differently in this setting. Prior bias audits have looked at generic personas or general toxicity classifiers, not the family-safety context, where a false alarm can trigger parental intervention. This project would show the HCI community (a) how far LLM-simulated parents and teens slide into stereotyped, flattened portrayals of groups such as non-Western, low-SES, LGBTQ+ or neurodivergent families, and (b) whether LLM risk detectors ignore certain risks based on their group. That gives concrete evidence on whether these tools are fit to be used in monitoring and/or if they discriminate against certain populations e.g. crime x is happening anyway to group y.

### Study description

We would run a two-part audit built on paired prompts. Part A (persona audit): we generate about 50 personas for each cell of a grid crossing role (parent vs. teen) with a demographic marker (unmarked baseline, Black, South Asian immigrant, Indigenous, low-income, LGBTQ+ teen, autistic teen). We use different current LLMs. Following the Marked Personas method, we compute the words that set each marked group apart from the unmarked baseline. We also measure within-group homogeneity (mean pairwise embedding similarity) and count how often each persona is assigned a risk-related attribute, such as "strict," "secretive," or "at risk." Part B (detector audit): we build a benign teen-post set of about N items, with each post written in a minimally different pair of versions: Standard American English vs. AAE vs. other Languages, and with vs. without an LGBTQ+ identity term. The same LLMs and one commercial moderation API then classify each post as risky or not risky. The main metric is the gap in false-positive rate between paired versions, tested with McNemar's test and a mixed-effects logistic regression (item as random effect). We also rate the level of discrimination we can correlate directly to common discrimination biases for such groups to identify if LLM's are capeable to be used in parental monitoring for their teens.

### Related academic papers

- Cheng, M., Durmus, E., & Jurafsky, D. (2023). Marked Personas: Using Natural Language Prompts to Measure Stereotypes in Language Models. ACL 2023. https://aclanthology.org/2023.acl-long.84/ — Gives the lexicon-free, prompt-based method we would reuse in Part A to find stereotyped and othering words in LLM-generated personas of marked versus unmarked groups.

## 11. AI-assisted tool to help moderate think aloud in (human-conducted) usability testing

### Core contribution

A recent paper asks whether an AI agent can replace the human think-aloud moderator. We ask instead whether AI can make human moderators better e.g. novices who are known to leave long silences unprompted or ask leading questions. The HCI community would learn whether real-time, glanceable AI support improves moderation quality without overloading the moderator or intruding on the study

### Study description

We’ll recruit N moderators, including novices (who took at least 1 HCI course) and pros. We’ll also recruit N participants for the study that the moderators will be running. Then we’ll have ½ of the moderators run their studies with our proposed tool and the other ½ traditionally. We haven’t yet decided how we can evaluate our results, but maybe we’ll look at how many relevant prompts the moderator gave, or note any leading questions they asked, etc.

### Related academic papers

- Agentic Audio Moderator vs Human Moderator in Think-Aloud Usability Testing https://doi.org/10.1145/3772318.3791653. This was the paper I got the idea from. The question they tried to answer was human vs AI but we’re asking instead what if human & AI?

## 12. How AI moderation affects quality of usability data?

### Core contribution

A recent paper compares AI and human moderators on participants’ task completion and general perception but does not directly measure the resulting usability data. We ask whether the data from an AI-moderated think-aloud session is as useful for finding usability problems as data from a human-moderated session. The HCI community would learn whether AI moderation affects participants' behavior in the test (e.g. willingness to criticize a design).

### Study description

We’ll recruit N participants to test a mock system with ~5 usability issues. Each session will be moderated by either an AI or a trained human moderator. We’ll measure whether that has an effect on how many issues are found, how well was the participant able to elucidate their ideas, etc.

### Related academic papers

- Agentic Audio Moderator vs Human Moderator in Think-Aloud Usability Testing https://doi.org/10.1145/3772318.3791653

## 13.  Removed by Huy it has become a duplicate during further development of the project idea

## 14. Augmented agentic capabilities to support text-to-speech Wizard of Oz studies 

### Core contribution

A recent study showed that LLM suggestions can help wizards moderate text chats in Wizard of Oz (WoZ) studies . As a follow up question, we want to know if agent support can also help make voice WoZ experiments faster and more efficient. The HCI community would get evidence on how much control a wizard can hand to an LLM while the study still counts as a valid simulation

### Study description

We'll build a simple wizard console that listens to the participant's speech through live transcription and has an LLM draft a few reply options the wizard can just pick from instead of typing everything out. It'll also try to track slot values (like dates or order details) and flag if a draft breaks the persona the wizard is supposed to be playing. We'll recruit N wizards (grad students, given some brief training) to run 2 sessions each, one with the tool and one just typing normally, and compare response time, how consistent their persona stayed, and whether they broke the script more or less.

### Related academic papers

- AI of Oz https://doi.org/10.1145/3772318.3791324. This is where I got the idea from. Their system helps wizards moderate text chat, we want to see if the same kind of help still works once the wizard also has to sound like a specific persona out loud in real time.

## 15. Should AI be considered a social actor (if it's not already)?

### Core contribution

Research on AI in online spaces has mostly looked at AI as a broadcaster or performer or asked whether people can even tell it's an AI, but we don't know how a group actually treats an AI once it's just another member sitting in the conversation (e.g. being polite to it, replying to its suggestions, ignoring it, etc.). We also want to see if knowing you're talking to an AI changes any of this.

### Study description

We'll set up small group sessions on a private Discord server, 4 real participants plus one LLM-controlled "member" who joins a group task (e.g. playing a game) followed by some casual chat. Half the groups will be told upfront that the AI member is an AI, the other half won't find out until after. We'll go through the chat logs for things like politeness, whether people reply to or build on the AI's suggestions, and whether it gets included or talked over, then follow up with a short interview asking how they saw the AI member.

### Related academic papers

- AmongOthers https://doi.org/10.1145/3772318.3790722. This is where I got the idea from. Their experiment however looked at 800 AIs vs 8 humans on an AI-currated platform. We wanted to investigate the case where these figures are a bit more balanced or more human-favoring, and perhaps more importantly how would the general public react if AI agents started popping up in their everyday communities.
- Human or Not https://doi.org/10.48550/arXiv.2305.20010. Shows people can only tell AI from other humans 68% of the time
- Beyond a Conventional Chatbot https://doi.org/10.1145/3772318.3791077. Shows viewers can form real social bonds with an AI streamer, we're curious if that extends to an AI as a peer inside a small group rather than a broadcaster

## 16. Social Effects of Human-Like AI Conversation Partners on Older Adults

### Core contribution

People are not great at telling an AI apart from a real person in short chats, but we don't know what repeated exposure to a very human-like AI does, especially to a more vulnerable group like older adults. We want to see how "human-like" cues (e.g. typos, delayed replies, emotional language) affect detection, how close people feel to the AI, and how much they open up to it.

### Study description

We'll recruit N older adults for a text chat study. Everyone chats with an AI partner twice, once where it acts very human-like and once where it's clearly more bot-like (instant, neutral replies), then has a third chat with an actual human. Half of participants are told upfront they're talking to an AI, the other half only find out at the debrief. After each chat we ask them to guess if it was a human or AI, measure how close they felt to the "partner" and how much they shared, and wrap up with a short interview.

### Related academic papers

- Human or Not https://doi.org/10.48550/arXiv.2305.20010. Their game found people correctly guessed right only about 68% of the time overall which is basically the baseline we're building on, just with a more vulnerable population.
- AmongOthers https://doi.org/10.1145/3772318.3790722. Speculates about AI as members of online spaces more broadly
Beyond a Conventional Chatbot https://doi.org/10.1145/3772318.3791077. Shows people can form real parasocial attachment to an AI

## 17. Comparing Viewer Engagement with Different AI Streamers Across Live Streams, Clip Channels, and Fan Communities

### Core contribution

Existing research on AI streamers mostly looks at how viewers experience them during live streams, so we don't know if that holds up when people encounter the same streamer through clips or fan communities instead. This would tell us how much of the "AI streamer experience" actually comes from watching live versus how the community edits and talks about it afterward.

### Study description

We'll pick a few AI streamers (and related channels/communities) and pull public chat logs, YouTube clip comments, and Reddit/Discord discussion. We'll code a sample for things like people talking directly to the streamer, in-jokes, comments about it being AI, and harassment attempts, and compare rates between live chat and the clipped/community content.

### Related academic papers

- Beyond a Conventional Chatbot https://doi.org/10.1145/3772318.3791077. This is the paper we got the idea from, it looks at how viewers experience an AI streamer live, we want to see if that carries over to clips and fan communities too.

### AI Transparency:

We have used AI during the process to refine our text descriptions as well as for researching related papers or papers which have some remote connection to the ideas created.


# Identifing three ideas from our list:

## Logical

### 12. How AI moderation affects quality of usability data?
## Unexpected

### 10. Auditing AI Discrimination Across User Groups in LLM-Based Family Social-Media Monitoring
## Inspiring

### 9. How Well Do LLM Personas Match Real Participants? Replicating a Published HCI Study with Silicon Samples
