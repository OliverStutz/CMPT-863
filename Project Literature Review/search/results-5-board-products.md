# Results 5: Boardroom, cybersecurity governance, CISO-board communication, products and VLM reliability

Agent 5. Searched and verified 2026-10-18. All papers checked against Crossref (api.crossref.org) on 2026-10-18 unless noted; abstracts read at the URL given. Products and trade sources accessed 2026-10-18.

26 entries in four parts: (A) scholarly work on boards, CISOs and cyber oversight (8), (B) industry and trade sources (5), (C) products: camera assistants, meeting assistants and live fact-checkers (6), (D) VLM and LLM reliability on charts, documents and claims (7).

Short take for the review:
- Boards are non-expert overseers who depend on the CISO and avoid questions so they do not look like novices (Opdenbusch et al., 2025; Lowry et al., 2026). Board members judge the CISO's competence by how well the CISO communicates (Proudfoot et al., 2023; Hooper & McKissack, 2016). The CISO's authority rests on interpreting an opaque field for the board (Da Silva & Jensen, 2022). This is the asymmetry our study puts a checker into.
- The SEC (2023) chose not to require board cyber expertise because directors "often effectively oversee management's efforts without specific subject matter expertise". This means the regulator assumes boards will stay non-expert, which is what makes an AI checker tempting.
- Products that let a listener point a camera at content and ask about it have been on the market since 2024 (Ray-Ban Meta, Apple Visual Intelligence, Gemini Live, ChatGPT video). Microsoft Copilot in Teams suggests the prompt "Where are some holes in [a meeting participant's] argument?". Live fact-checkers (Squash, Full Fact AI, Factiverse, ClaimBuster) check speakers against databases or the web. None of these products comes with published evidence on what the check does to the audience's trust in the speaker.
- VLMs misread charts and documents. GPT-4o got 47.1% of chart-reasoning questions right where humans got 80.5% (CharXiv). GPT-4V scored at about the 16th percentile of humans on a visualization literacy test (Bendeck & Stasko). Without context, LLM fact-checks miss many claims that are true (Quelle & Bovet). On the other side, multimodal LLMs can spot textbook chart problems when prompted (Lo & Qu). Together these support our premise that a flag can be "right by the textbook and wrong for this company".

---

## A. Board-level cybersecurity governance and the CISO-board relationship (scholarly)

### Lowry, M. R., Vance, A., & Vance, M. D. (2026). Inexpert supervision: Field evidence on boards' oversight of cybersecurity. *Management Science*, *72*(2), 783–804. https://doi.org/10.1287/mnsc.2023.04147
- **Type:** management or economics
- **Verified:** Crossref 2026-10-18 (vol. 72, issue 2, pp. 783–804, published February 2026; OpenAlex lists online-first/SSRN 2021); abstract read at https://api.crossref.org/works/10.1287/mnsc.2023.04147
- **Design:** Field study (interviews) with board directors who have cyber expertise and directors who do not, plus cybersecurity executives and consultants who support boards. Compared the substance of oversight by expert and non-expert directors.
- **Finding:** Non-expert directors tried to oversee cybersecurity diligently. Even when they performed the same oversight activities as expert directors, their oversight lacked substance and was "mostly symbolic". Non-expert directors did not see that their efforts were symbolic and insufficient. Expert directors, CISOs and consultants saw the gap clearly and argued for more cyber experts on boards.
- **Relation:** Context and gap. This is the strongest scholarly evidence that board members cannot judge a CISO's technical claims on their own. That is why a checker would appeal to them, and why its flags would be hard to evaluate. Lowry et al. describe oversight without any tool; we test what happens to trust when a non-expert director gets an AI checker.
- **Access:** Open access preprint https://doi.org/10.2139/ssrn.4002794 | published version SFU library
- **Score:** 1=2 2=0 3=5 4=5 5=3 total=15/25

### Opdenbusch, J. C., Hielscher, J., & Sasse, M. A. (2025). "Where are we on cyber?": A qualitative study on boards' cybersecurity risk decision making. In *Proceedings of the 2025 Network and Distributed System Security Symposium (NDSS 2025)*. Internet Society. https://doi.org/10.14722/ndss.2025.240595
- **Type:** other HCI (usable security, NDSS)
- **Verified:** Crossref 2026-10-18; abstract read at https://api.openalex.org/works/doi:10.14722/ndss.2025.240595 and https://www.ndss-symposium.org/ndss-paper/where-are-we-on-cyber-a-qualitative-study-on-boards-cybersecurity-risk-decision-making/
- **Design:** In-depth interviews with 18 C-level managers, board members, CISOs and C-level consultants at some of the largest UK-based companies, on how boards decide once they receive cyber information.
- **Finding:** The authors report a power imbalance. Board members often do not ask executives and CISOs the right questions because they "fear being exposed as IT novices". This makes boards highly dependent on whoever supplies their cyber information, and they lose their oversight function. Boards reduce cybersecurity risk to budget decisions and are not otherwise involved in cyber strategy.
- **Relation:** Context and H4. Our design rests on this finding: board members usually stay silent. A checker could give them a reason to ask (H4, more questioning). The budget framing fits our scenario, in which the CISO asks for a budget. The study is qualitative; ours is an experiment that measures whether a checker changes the choice to question and how trust follows.
- **Access:** Open access https://www.ndss-symposium.org/wp-content/uploads/2025-595-paper.pdf
- **Score:** 1=2 2=0 3=5 4=5 5=4 total=16/25

### Da Silva, J., & Jensen, R. B. (2022). "Cyber security is a dark art": The CISO as soothsayer. *Proceedings of the ACM on Human-Computer Interaction*, *6*(CSCW2), 1–31. https://doi.org/10.1145/3555090
- **Type:** SIGCHI (CSCW)
- **Verified:** Crossref 2026-10-18 (vol. 6, CSCW2, pp. 1–31, November 2022); abstract read at https://arxiv.org/abs/2202.12755 (via OpenAlex)
- **Design:** Semi-structured interviews with 15 CISOs and 6 senior organisational leaders (October 2019 to July 2020), interpretive analysis drawing on ontological security and identity work.
- **Finding:** Cybersecurity works as an expert system that makes the CISO the interpreter of something "mystical, unknown and fearful to the uninitiated". The CISO acts as a "modern-day soothsayer" for senior management, in a role that is both superior and precarious. That precarious position motivates self-serving behaviour, which the authors call "cyber sophistry".
- **Relation:** Context and theme anchor (SIGCHI). It names the epistemic asymmetry our study manipulates. An AI checker breaks the CISO's monopoly on interpretation. The study describes the relationship; we measure how trust changes when a third, machine voice enters it.
- **Access:** Open access https://arxiv.org/pdf/2202.12755
- **Score:** 1=2 2=0 3=5 4=5 5=4 total=16/25

### Proudfoot, J. G., Cram, W. A., Madnick, S., & Coden, M. (2023). The importance of board member actions for cybersecurity governance and risk management. *MIS Quarterly Executive*, *22*(4), 235–250. https://doi.org/10.17705/2msqe.00084
- **Type:** management or economics (IS practitioner journal)
- **Verified:** Crossref 2026-10-18 (pages 235–250; volume and issue from https://aisel.aisnet.org/misqe/vol22/iss4/6); full text read at https://cams.mit.edu/wp-content/uploads/MISQE_Board_Member_Actions_Proudfoot_Dec23.pdf
- **Design:** In-depth interviews with 35 cybersecurity experts and current board members (executives, CISOs, CTOs, compliance officers, board advisors) across finance, technology, healthcare, critical infrastructure and other sectors. The authors group the findings into four challenges and give ten recommended actions.
- **Finding:** Challenge 2 is "board-executive interaction dynamics". Interviewees said the "board's perceived competence of the CISO" is "heavily determined by the CISO's ability to effectively communicate". A highly expert CISO who cannot explain plainly may be seen as incompetent. A marginally qualified CISO "who appears credible" may "(dangerously) be perceived as highly qualified". A board member described "blank stares and no follow-up questions" when CISOs cannot present.
- **Relation:** H2 and H3, context. This is direct practitioner evidence that boards judge CISO competence from the explanation, not from technical substance. That predicts a reasoned answer will restore trust (H2) and a bare denial will not (H3). The paper reports interview claims; we test them experimentally, and with a checker added.
- **Access:** Open access author copy https://cams.mit.edu/wp-content/uploads/MISQE_Board_Member_Actions_Proudfoot_Dec23.pdf
- **Score:** 1=3 2=0 3=5 4=5 5=5 total=18/25

### Gale, M., Bongiovanni, I., & Slapničar, S. (2022). Governing cybersecurity from the boardroom: Challenges, drivers, and ways ahead. *Computers & Security*, *121*, Article 102840. https://doi.org/10.1016/j.cose.2022.102840
- **Type:** other (information security journal)
- **Verified:** Crossref 2026-10-18 (vol. 121, article 102840). The publisher page returned 403 and Crossref and Semantic Scholar carry no abstract. Abstract sentences were read in search snippets of the ScienceDirect and ResearchGate pages (https://www.sciencedirect.com/science/article/abs/pii/S0167404822002346).
- **Design:** 18 interviews with non-executive directors who sit on 43 organisations' boards, analysed through neo-institutional theory (coercive, normative and mimetic pressures).
- **Finding:** Boards are less engaged in cybersecurity than in other areas of oversight. Regulation is the strongest driver of engagement. Directors are not always aware of their cyber duties and liability. Boards often over-rely on the one member who has cyber experience. Secrecy around cyber incidents limits learning from other firms. The authors call for clearer regulation, director training and clearer reporting. A later review cites the paper for the lack of metrics to assess cyber investments.
- **Relation:** Context. It shows that directors lack the knowledge to evaluate cyber reports and lean on whoever holds it. The checker in our study could take the place of the "one cyber-literate member". Gale et al. describe governance structures; we measure individual trust.
- **Access:** SFU library
- **Score:** 1=1 2=0 3=5 4=4 5=2 total=12/25

### Ruth, K., Rivera, V. A., Akiwate, G., Fass, A., Kelley, P. G., Thomas, K., & Durumeric, Z. (2025). "Perfect is the enemy of good": The CISO's role in enterprise security as a business enabler. In *Proceedings of the 2025 CHI Conference on Human Factors in Computing Systems* (pp. 1–19). Association for Computing Machinery. https://doi.org/10.1145/3706598.3713895
- **Type:** SIGCHI
- **Verified:** Crossref 2026-10-18 (CHI 2025, pp. 1–19); full text read at https://zakird.com/papers/chi25-ciso.pdf
- **Design:** Interviews with 16 current and former CISOs on how they build a security strategy and what obstacles they face.
- **Finding:** CISOs frame security as business enablement. With the board they hold a "compromise conversation" in which they trade cyber risk against other business risks. Much of the work is "educational, it's not always technical". CISOs were careful about what they asked for so they would not be dismissed as "crying wolf". They lacked a "harmonized way of expressing risk" to non-expert business leaders. Getting the board's buy-in depended on "build[ing] effective business cases".
- **Relation:** Context and design element. It supports a briefing that ends with a budget ask, and a CISO whose answers give business reasons ("the company's reason"), as in our reasoned-answer condition. It is a SIGCHI anchor for the CISO side. Our study takes the board's side and adds a checker.
- **Access:** Open access https://zakird.com/papers/chi25-ciso.pdf
- **Score:** 1=1 2=0 3=5 4=4 5=5 total=15/25

### Hooper, V., & McKissack, J. (2016). The emerging role of the CISO. *Business Horizons*, *59*(6), 585–591. https://doi.org/10.1016/j.bushor.2016.07.004
- **Type:** management or economics
- **Verified:** Crossref 2026-10-18 (vol. 59, issue 6, pp. 585–591). The abstract is elided in Crossref and Semantic Scholar; a summary of it was read at https://ideas.repec.org/a/eee/bushor/v59y2016i6p585-591.html
- **Design:** Practitioner-oriented article drawing on data from the US, Canada and New Zealand on the CISO role and CISO-board communication.
- **Finding:** Organisations must define clearly what they expect of the CISO. Because the CISO is a senior executive, the role calls for an excellent communicator more than a specialised technical expert. The CISO must communicate effectively with the board, and the article suggests ways to make that communication more effective. Organisations should also build cybersecurity concern into the selection criteria for board members.
- **Relation:** Context, supporting H2. CISO standing with the board depends on communication. This is an early source the brief expected. Its evidence base is thin compared with Proudfoot et al. (2023).
- **Access:** SFU library
- **Score:** 1=1 2=0 3=4 4=3 5=2 total=10/25

### Karanja, E. (2017). The role of the chief information security officer in the management of IT security. *Information and Computer Security*, *25*(3), 300–329. https://doi.org/10.1108/ICS-02-2016-0013
- **Type:** other (information security journal)
- **Verified:** Crossref 2026-10-18 (vol. 25, issue 3, pp. 300–329); abstract read at https://api.openalex.org/works/doi:10.1108/ics-02-2016-0013
- **Design:** Content analysis, framed by agency theory, of US firms that had IT security breaches between 2009 and 2015 and how they reacted.
- **Finding:** After a breach, many firms reorganised their security strategy and hired a CISO. Firms that hired a CISO for the first time did not agree on where the CISO should report: most had the CISO report to either the CEO or the CIO.
- **Relation:** Context only (agency framing of CISO and principals). Marginal for H1 to H4. Included because the brief names it and markers may expect it.
- **Access:** SFU library
- **Score:** 1=0 2=0 3=3 4=2 5=3 total=8/25

---

## B. Industry and trade sources

### U.S. Securities and Exchange Commission. (2023). *Cybersecurity risk management, strategy, governance, and incident disclosure* (Release Nos. 33-11216; 34-97989). https://www.sec.gov/files/rules/final/2023/33-11216.pdf
- **Type:** industry or trade (regulation)
- **Verified:** Full release PDF and fact sheet (https://www.sec.gov/files/33-11216-fact-sheet.pdf) downloaded and read 2026-10-18
- **Design:** Final rule adopted 26 July 2023. New Regulation S-K Item 106 requires registrants to describe "the board of directors' oversight of risks from cybersecurity threats and management's role and expertise". A new Form 8-K Item 1.05 requires disclosure of a material incident generally within four business days of determining that it is material. Item 106 disclosures apply from fiscal years ending on or after 15 December 2023.
- **Finding (the one fact that matters):** The SEC dropped the proposed Item 407(j), which would have required disclosure of board cyber expertise. The release states (PDF p. 85): "we are not adopting proposed Item 407(j). We are persuaded that effective cybersecurity processes are designed and administered largely at the management level, and that directors with broad-based skills in risk management and strategy often effectively oversee management's efforts without specific subject matter expertise."
- **Relation:** Context and gap. Regulation now requires boards to oversee cyber risk and assumes they will do it without subject expertise. This sets up the non-expert overseer in our scenario. Lowry et al. (2026) suggest that such oversight tends to be symbolic.
- **Access:** Open access (URL above). Accessed 2026-10-18.
- **Score:** 1=0 2=0 3=5 4=4 5=5 total=14/25

### National Association of Corporate Directors & Internet Security Alliance. (2023). *Director's handbook on cyber-risk oversight* (4th ed.). NACD. https://www.nacdonline.org/all-governance/governance-resources/trending-oversight-topics/cybersecurity/
- **Type:** industry or trade (handbook)
- **Verified:** Press release read at https://www.nacdonline.org/about/newsroom/press-release/press-release/nacd-and-isa-launch-2023-cyber-risk-oversight-handbook-featuring-cisa-and-fbi/ (dated 22 March 2023). Principle wording read at https://www.fairinstitute.org/blog/nacd-directors-handbook-cyber-risk-quantification. The handbook PDF itself could not be downloaded; the full text is free in digital form but needs NACD registration.
- **Design:** Fourth edition (the first was in 2014), developed with the US Department of Homeland Security and the Department of Justice, with a foreword by CISA Director Jen Easterly. It sets out six principles, each with questions for directors to ask management.
- **Finding (the one fact that matters):** Principle 1: "Directors need to understand and approach cybersecurity as a strategic, enterprise risk, not just an IT risk." Principle 5: board-management discussions "should include identification and quantification of financial exposure". The handbook gives directors scripted questions to ask management, such as how the organisation measures the threat environment and what its cyber risk profile is. Secondary sources list Principle 3 as "Board oversight structure and access to expertise".
- **Relation:** Context and design element. The professional norm is that directors should question the CISO. Our "ask the CISO to explain" and "challenge" options turn that norm into an action participants can take. The handbook's question lists are a human-authored alternative to an AI checker.
- **Access:** Free digital version with NACD registration. Accessed 2026-10-18.
- **Score:** 1=0 2=0 3=5 4=3 5=2 total=10/25

### Proofpoint. (2023). *Cybersecurity: The 2023 board perspective*. Proofpoint. https://www.proofpoint.com/us/newsroom/press-releases/proofpoints-second-annual-board-perspective-report-reveals-nearly-three-five
- **Type:** industry or trade (consultancy survey)
- **Verified:** Press release read 2026-10-18 (dated 6 September 2023); 2022 numbers checked at https://www.helpnetsecurity.com/2022/10/10/board-cybersecurity-perceptions/
- **Design:** Survey of 659 board members at organisations with 5,000 or more employees in 12 countries, June 2023. Compared with Proofpoint's CISO survey. The 2022 edition was produced with Cybersecurity at MIT Sloan (about 600 directors).
- **Finding (the one number that matters):** 65% of board members said they "see eye to eye" with their CISO, against 62% of CISOs who said the same about their board. In 2022 the figures were 69% of board members and 51% of CISOs. 73% of directors felt at risk of a material cyber attack and 53% saw their organisation as unprepared. 59% saw generative AI as a security risk.
- **Relation:** Context. A third of board members and CISOs report misalignment, so the relationship is not secure and a flag could tip it. It is a vendor survey, so these are self-reports from a commercial source.
- **Access:** Open access press release; full report free behind a form. Accessed 2026-10-18.
- **Score:** 1=2 2=0 3=5 4=3 5=3 total=13/25

### Milică, L., & Pearlson, K. (2023, May 2). Boards are having the wrong conversations about cybersecurity. *Harvard Business Review*. https://hbr.org/2023/05/boards-are-having-the-wrong-conversations-about-cybersecurity
- **Type:** industry or trade (HBR)
- **Verified:** HBR page read 2026-10-18 (paywalled; only the summary was visible). The author list and date were confirmed in HBR search results, which also list Pearlson & Hetner (2022, 11 November), "Is your board prepared for new cybersecurity regulations?" (https://hbr.org/2022/11/is-your-board-prepared-for-new-cybersecurity-regulations), and Proudfoot & Madnick (2026, 2 April), "Boards are falling short on cybersecurity" (https://hbr.org/2026/04/boards-are-falling-short-on-cybersecurity).
- **Design:** Opinion and analysis by Proofpoint's Global Resident CISO and an MIT Sloan researcher, drawing on the Proofpoint and MIT CAMS board surveys.
- **Finding (the one fact that matters):** Boards say cybersecurity is a priority but focus on prevention rather than resilience, and "by not focusing on resilience, boards fail their companies". Directors with cyber expertise are trying to get their colleagues' attention.
- **Relation:** Context. This is the HBR source the brief asked for. Note: the brief's "Pearlson & Hetner 2023" is in fact HBR 11 November 2022, and Proudfoot & Madnick (2026) is the newest HBR piece on the topic. I read only the HBR summaries, so cite these for framing, not for findings.
- **Access:** HBR paywall (SFU library via HBR in Business Source Complete). Accessed 2026-10-18.
- **Score:** 1=0 2=0 3=5 4=2 5=1 total=8/25

### Gartner. (2021, November 18). *Gartner survey finds 88 percent of boards of directors view cybersecurity as a business risk* [Press release]. https://www.gartner.com/en/newsroom/press-releases/2021-11-18-gartner-survey-finds-88-percent-of-boards-of-directors-view-cybersecurity-as-a-business-risk
- **Type:** industry or trade (consultancy survey)
- **Verified:** Search results quoting the press release, 2026-10-18 (the page itself was not fetched)
- **Design:** Gartner's 2022 Board of Directors Survey.
- **Finding (the one number that matters):** 88% of boards view cybersecurity as a business risk rather than a technology risk (up from 58% five years earlier). Only 12% have a dedicated board-level cybersecurity committee. The CIO, CISO or equivalent is the person held accountable for cybersecurity in 85% of organisations.
- **Relation:** Context. Boards now own cyber risk, but they place accountability for it on the CISO. That is the expert on whom the board depends in our scenario.
- **Access:** Open access press release. Accessed 2026-10-18.
- **Score:** 1=0 2=0 3=4 4=2 5=2 total=8/25

---

## C. Products: camera assistants, meeting assistants and live fact-checkers

### Camera-based "look and ask" assistants: Ray-Ban Meta with Meta AI (2024), Apple Visual Intelligence (2024), ChatGPT Advanced Voice with video (2024), Google Gemini Live with camera (2025)
- **Type:** product
- **Verified (accessed 2026-10-18):**
  - Meta. (2024, April 23). *New Ray-Ban | Meta smart glasses styles and Meta AI updates*. https://about.fb.com/news/2024/04/new-ray-ban-meta-smart-glasses-styles-and-meta-ai-updates/
  - Apple. (2024, December 11). *Apple Intelligence now features Image Playground, Genmoji, and more*. https://www.apple.com/newsroom/2024/12/apple-intelligence-now-features-image-playground-genmoji-and-more/
  - Neowin (2024, December 12). *ChatGPT mobile apps now support video and screensharing in Advanced Voice Mode*. https://www.neowin.net/news/chatgpt-mobile-apps-now-support-video-and-screensharing-in-advanced-voice-mode/ (secondary; OpenAI's help page returned 403)
  - 9to5Google (2025, April 7). *Gemini Live camera and screen sharing*. https://9to5google.com/2025/04/07/gemini-live-camera-screen-sharing/ (secondary)
- **What they do and since when:**
  - Ray-Ban Meta: the wearer says "Hey Meta" and asks about what they are looking at. The glasses' camera and Meta AI answer hands-free, for example by translating a menu. Multimodal beta for all Ray-Ban Meta glasses in the US and Canada from 23 April 2024, after testing that began in December 2023.
  - Apple Visual Intelligence (iOS 18.2, 11 December 2024, iPhone 16 Camera Control): summarises, copies and translates text. Users "can ask ChatGPT for help with complex visuals, such as an explanation of a diagram from class notes".
  - ChatGPT Advanced Voice Mode: live video and screen sharing rolled out from 12 December 2024 to Plus, Pro and Team users (EU later), with daily limits.
  - Gemini Live: rear-camera and screen sharing so users can talk about what they see. Rolled out from April 2025 (Pixel 9, Galaxy S25, then all Android users), and free in the Gemini app from May 2025.
- **Finding:** The interaction in our checker condition already exists in consumer products: point a camera at a slide or diagram and get a spoken or written verdict. Apple's own marketing example is explaining a diagram. None of these vendors reports accuracy on business slides, and none shows the user the speaker's context.
- **Relation:** Product context and ecological validity of the simulated checker. Our study asks what happens when the listener uses such a tool on a human presenter, which no vendor or study has measured.
- **Access:** Open web pages, accessed 2026-10-18.
- **Score:** 1=0 2=4 3=1 4=4 5=4 total=13/25

### Microsoft. (n.d.). *Use Copilot in Microsoft Teams meetings* [Support article]; and Otter.ai Meeting Agent (2025)
- **Type:** product
- **Verified (accessed 2026-10-18):**
  - https://support.microsoft.com/en-us/office/use-copilot-in-microsoft-teams-meetings-0bf9dd3c-96f7-44e2-8bb8-790bedf066b1
  - UC Today (2025, March 26). *Otter revolutionises meetings with AI agent that speaks up during calls*. https://www.uctoday.com/unified-communications/otter-revolutionises-meetings-with-ai-agent-that-speaks-up-during-calls/ (Otter's Business Wire release of 25 March 2025 returned 403)
- **What they do and since when:**
  - Copilot in Teams answers a participant's questions in real time during a meeting, in a private pane, from the transcript and chat. Microsoft's own suggested prompts include "Where do we disagree on this topic?", "How did [a meeting participant] respond to this proposal?" and "Where are some holes in [a meeting participant's] argument?". Follow-up suggestions reached general availability in March 2025 (admin message centre via m365admin.handsontek.net). It does not work in meetings hosted outside the organisation and needs transcription.
  - Otter Meeting Agent, launched 25 to 26 March 2025, is a "voice-activated" participant that answers questions live from the organisation's meeting database ("Hey Otter"). It started on Zoom, with Teams and Google Meet planned.
- **Finding:** Microsoft markets a private AI critique of a colleague's argument, visible only to the asker, during a live meeting. That is close to our checker, with the transcript taking the place of the camera image.
- **Relation:** Product context and design element. A board member using Copilot during a CISO's briefing is a realistic scenario today. There is no published evidence on what such prompts do to trust in the speaker.
- **Access:** Open web pages, accessed 2026-10-18.
- **Score:** 1=0 2=5 3=3 4=5 5=4 total=17/25

### Adair, B. (2020, February 23). *Squash report card: Improvements during State of the Union … and how humans will make our AI smarter*. Duke Reporters' Lab. https://reporterslab.org/2020/02/23/squash-report-card-improvements-during-state-of-the-union-and-how-humans-will-make-our-ai-smarter/
- **Type:** product (research prototype)
- **Verified:** Page read 2026-10-18
- **Design:** Squash, from the Tech & Check Cooperative, listens to a speech, converts it to text, filters factual claims with ClaimBuster, and matches them with Elasticsearch against published fact-checks (Washington Post, FactCheck.org, PolitiFact). Matches pop up on the video screen.
- **Finding (the one fact that matters):** During the 2020 State of the Union, Squash showed 14 relevant fact-checks and 20 non-relevant ones, up from 6 relevant pop-ups in 2019. The lab responded by adding a human editor who approves each pop-up. A UX test found viewers liked the idea of real-time fact-checks but were often confused by the pop-up design.
- **Relation:** Product and counter-evidence on checker accuracy. A fully automated live checker produced more wrong matches than right ones, and its makers added a human in the loop. This is a real-world precedent for flags that are wrong in context, as all of ours are.
- **Access:** Open access, accessed 2026-10-18.
- **Score:** 1=0 2=4 3=2 4=4 5=4 total=14/25

### Hassan, N., Arslan, F., Li, C., & Tremayne, M. (2017). Toward automated fact-checking: Detecting check-worthy factual claims by ClaimBuster. In *Proceedings of the 23rd ACM SIGKDD International Conference on Knowledge Discovery and Data Mining* (pp. 1803–1812). Association for Computing Machinery. https://doi.org/10.1145/3097983.3098131
- **Type:** product (system paper, KDD)
- **Verified:** Crossref 2026-10-18 (title registered as "Toward Automated Fact-Checking", pp. 1803–1812); abstract read at https://api.openalex.org/works/doi:10.1145/3097983.3098131
- **Design:** System description. A supervised claim-spotting model trained on human-labelled check-worthy sentences from US general-election debate transcripts. It was deployed live during the 2016 US presidential debates and to monitor social media and the Australian Hansard.
- **Finding:** ClaimBuster scores sentences for check-worthiness in real time and matches them with a repository of professional fact-checks. A 2015 Poynter report (secondary, https://poynter.org/?p=379687) said it agreed with human analysts on 74% of check-worthy statements.
- **Relation:** Product. This is the canonical live claim spotter and the filter inside Squash. It decides what to check but gives no verdict tied to the speaker's context. No audience trust outcomes.
- **Access:** Open access author copy https://ranger.uta.edu/~cli/pubs/2017/claimbuster-kdd17-hassan.pdf
- **Score:** 1=0 2=4 3=2 4=2 5=3 total=11/25

### Full Fact AI (since 2016) and Factiverse (live debate fact-checking, 2024)
- **Type:** product
- **Verified (accessed 2026-10-18):**
  - Full Fact. (n.d.). *About Full Fact AI*. https://fullfact.org/ai/about/
  - Bellan, R. (2024, November 17). *Norwegian startup Factiverse wants to fight disinformation with AI*. TechCrunch. https://techcrunch.com/2024/11/17/norwegian-startup-factiverse-wants-to-fight-disinformation-with-ai
- **What they do and since when:**
  - Full Fact has used machine learning since 2016. Its tools take in news, live TV speech, podcasts and social media, split them into sentences, flag checkable claims with a BERT-based classifier, and match repeats of claims already checked. They are used daily by "over 40 fact checking organisations" in 30 countries and three languages, and supported monitoring of 12 national elections through 2024.
  - Factiverse sells B2B live fact-checking of text, audio and video. It says it is "not an LLM" but a retrieval model that ranks sources by credibility. It fact-checked the 2024 US presidential debates live for media partners. The CEO claims an 80% success rate (self-reported, not independently verified).
- **Finding:** Live, AI-driven checking of what a speaker says is commercially deployed. Its reported accuracy (around 80% by the vendor's own claim) implies frequent wrong verdicts, and these products are built for journalists, not for a listener judging a colleague.
- **Relation:** Product context. These show the live fact-checker as a working genre. Our checker differs in three ways: it reads visual slides, it sits with a single audience member, and it judges a colleague in a private, high-stakes setting rather than a public politician.
- **Access:** Open web pages, accessed 2026-10-18.
- **Score:** 1=0 2=4 3=2 4=3 5=3 total=12/25

### Wojcik, S., Hilgard, S., Judd, N., Mocanu, D., Ragain, S., Hunzaker, M. B. F., Coleman, K., & Baxter, J. (2022). *Birdwatch: Crowd wisdom and bridging algorithms can inform understanding and reduce the spread of misinformation* (arXiv:2210.15723). arXiv. https://doi.org/10.48550/arXiv.2210.15723
- **Type:** product (platform research, preprint)
- **Verified:** arXiv abstract page read 2026-10-18 (https://arxiv.org/abs/2210.15723; submitted 27 October 2022). DataCite DOI pattern for arXiv; not checked in Crossref because arXiv DOIs are registered with DataCite.
- **Design:** Contributors write notes on posts and rate each other's notes. A matrix-factorisation "bridging-based ranking" shows only notes rated helpful by raters who usually disagree. Tested in a survey experiment and in deployment on Twitter (now X Community Notes).
- **Finding:** Bridging-selected notes did better than overall-average and crowd-generated baselines on the survey's key indicators. On Twitter, users who saw these notes were significantly less likely to reshare the annotated post. (The abstract gives no numbers.)
- **Relation:** Product and design contrast. Community Notes shows a verdict only after agreement across viewpoints. Our checker gives a single machine verdict with no such safeguard and no context from the speaker. The two make a useful contrast for the discussion.
- **Access:** Open access https://arxiv.org/abs/2210.15723
- **Score:** 1=0 2=3 3=1 4=3 5=3 total=10/25

---

## D. Reliability of vision-language models and LLM fact-checkers

### Bendeck, A., & Stasko, J. (2025). An empirical evaluation of the GPT-4 multimodal language model on visualization literacy tasks. *IEEE Transactions on Visualization and Computer Graphics*, *31*(1), 1105–1115. https://doi.org/10.1109/TVCG.2024.3456155
- **Type:** other HCI (IEEE VIS / TVCG)
- **Verified:** Crossref 2026-10-18 (vol. 31, issue 1, pp. 1105–1115); full author version read at https://faculty.cc.gatech.edu/~john.stasko/papers/vis24-llm.pdf
- **Design:** GPT-4V answered four task sets: the 53-item Visualization Literacy Assessment Test (VLAT), a chart question-answering set with and without the underlying data, questions on deceptive design choices, and questions on misaligned chart titles.
- **Finding:** On the VLAT, GPT-4 scored 19.67, against a human mean of 28.82. That puts it at about the 16th percentile of humans. It recognised trends and extreme values and showed some knowledge of design best practice. It struggled to retrieve simple values when it did not have the data: on chart questions needing a value read off the chart, accuracy was 13% without the data and nearly perfect with it. It could not reliably tell colours apart, and it sometimes hallucinated or answered inconsistently.
- **Relation:** VLM reliability, supporting the premise of wrong flags. A VLM looking only at a slide image misreads numbers but "knows" textbook design rules. That is exactly the profile behind a flag such as "Numbers do not add up" that is wrong in context.
- **Access:** Open access author version (URL above); OSF materials https://doi.org/10.17605/OSF.IO/F39J6
- **Score:** 1=0 2=3 3=1 4=5 5=5 total=14/25

### Lo, L. Y.-H., & Qu, H. (2025). How good (or bad) are LLMs at detecting misleading visualizations? *IEEE Transactions on Visualization and Computer Graphics*, *31*(1), 1116–1125. https://doi.org/10.1109/TVCG.2024.3456333
- **Type:** other HCI (IEEE VIS / TVCG)
- **Verified:** Crossref 2026-10-18 (vol. 31, issue 1, pp. 1116–1125); abstract read at https://arxiv.org/abs/2407.17291
- **Design:** Four multimodal LLMs were given a dataset of misleading charts collected from the internet, with nine prompts ranging from simple to complex. Three experiments expanded detection from 5 to 21 chart issues.
- **Finding:** The authors conclude that multimodal LLMs have "a strong capability for chart comprehension and critical thinking in data interpretation" and could help counter misleading charts. Prompt design strongly affected what the models found, and covering many issue types raised scalability problems.
- **Relation:** VLM reliability, as a counter-result. LLMs do detect textbook chart problems, which makes our checker's flags plausible ("right by the textbook"). Lo is a co-author of the known paper Lo et al. (2022). Together with Bendeck & Stasko this frames the checker as good at rule-based critique and blind to company context.
- **Access:** Open access preprint https://arxiv.org/abs/2407.17291
- **Score:** 1=0 2=3 3=1 4=4 5=3 total=11/25

### Wang, Z., Xia, M., He, L., Chen, H., Liu, Y., Zhu, R., Liang, K., Wu, X., Liu, H., Malladi, S., Chevalier, A., Arora, S., & Chen, D. (2024). CharXiv: Charting gaps in realistic chart understanding in multimodal LLMs. In *Advances in Neural Information Processing Systems 37* (pp. 113569–113697). Curran Associates. https://doi.org/10.52202/079017-3609
- **Type:** other (machine learning benchmark, NeurIPS)
- **Verified:** Crossref 2026-10-18 (13 authors, pp. 113569–113697); abstract read at https://arxiv.org/abs/2406.18521
- **Design:** 2,323 natural, diverse charts from arXiv papers, each hand-picked and verified by human experts, with descriptive and reasoning questions. Tested on proprietary and open-source multimodal LLMs and compared with humans.
- **Finding:** On reasoning questions, the strongest model (GPT-4o) was 47.1% accurate, the strongest open-source model 29.2%, and humans 80.5%. Slightly changing the charts or questions in older benchmarks cut performance by up to 34.5%, so earlier benchmarks overstated progress.
- **Relation:** VLM reliability. On realistic charts, the best model gets more than half of the reasoning questions wrong. This is a citable error rate for "a VLM reading a slide chart is often wrong".
- **Access:** Open access https://arxiv.org/abs/2406.18521
- **Score:** 1=0 2=2 3=0 4=4 5=4 total=10/25

### Ma, Y., Zang, Y., Chen, L., Chen, M., Jiao, Y., Li, X., Lu, X., Liu, Z., Ma, Y., Dong, X., Zhang, P., Pan, L., Jiang, Y.-G., Wang, J., Cao, Y., & Sun, A. (2024). MMLongBench-Doc: Benchmarking long-context document understanding with visualizations. In *Advances in Neural Information Processing Systems 37* (pp. 95963–96010). Curran Associates. https://doi.org/10.52202/079017-3041
- **Type:** other (machine learning benchmark, NeurIPS Datasets and Benchmarks)
- **Verified:** Crossref 2026-10-18 (16 authors, pp. 95963–96010); abstract read at https://arxiv.org/abs/2407.01523
- **Design:** 1,062 expert-annotated questions on 130 PDF documents (on average 49.4 pages), with evidence in text, images, charts, tables and layout. 33.2% of questions need evidence across pages, and 22.8% are deliberately unanswerable to detect hallucination. 14 large vision-language models were tested.
- **Finding:** The best model, GPT-4o, reached an F1 of 42.7%, and GPT-4V 31.4%. Twelve of the 14 models did worse than their text-only counterparts given lossy OCR text.
- **Relation:** VLM reliability on documents. Even the best models fail on most questions about multi-page business-style documents. A checker that sees one slide at a time also lacks the cross-page context, which is how our flags go wrong.
- **Access:** Open access https://arxiv.org/abs/2407.01523
- **Score:** 1=0 2=2 3=0 4=4 5=4 total=10/25

### Guan, T., Liu, F., Wu, X., Xian, R., Li, Z., Liu, X., Wang, X., Chen, L., Huang, F., Yacoob, Y., Manocha, D., & Zhou, T. (2024). HallusionBench: An advanced diagnostic suite for entangled language hallucination and visual illusion in large vision-language models. In *Proceedings of the 2024 IEEE/CVF Conference on Computer Vision and Pattern Recognition (CVPR)* (pp. 14375–14385). IEEE. https://doi.org/10.1109/CVPR52733.2024.01363
- **Type:** other (computer vision, CVPR)
- **Verified:** Crossref 2026-10-18 (pp. 14375–14385); abstract read at https://arxiv.org/abs/2310.14566
- **Design:** 346 images with 1,129 expert-written questions in paired control structures (including charts and tables, some edited), to test whether models reason from the image or from prior knowledge. 15 models tested, including GPT-4V, Gemini Pro Vision, Claude 3 and LLaVA-1.5.
- **Finding:** GPT-4V answered both questions of a pair correctly only 31.42% of the time, and every other model scored below 16%. The main failure modes were "language hallucination" (answering from prior knowledge rather than the image) and "visual illusion".
- **Relation:** VLM reliability. Language hallucination is the mechanism behind "right by the textbook, wrong for this company": the model answers from what is usually true rather than from what the slide shows in context.
- **Access:** Open access https://arxiv.org/abs/2310.14566
- **Score:** 1=0 2=2 3=0 4=4 5=4 total=10/25

### Rahmanzadehgervi, P., Bolton, L., Taesiri, M. R., & Nguyen, A. T. (2024). Vision language models are blind. In *Computer Vision – ACCV 2024* (Lecture Notes in Computer Science, pp. 293–309). Springer. https://doi.org/10.1007/978-981-96-0917-8_17
- **Type:** other (computer vision, ACCV)
- **Verified:** Crossref 2026-10-18 (pp. 293–309; LNCS volume number not given in Crossref); abstract read at https://arxiv.org/abs/2407.06581
- **Design:** BlindTest: seven very simple visual tasks (do two circles overlap, how many times do two lines cross, which letter is circled, how many circles are in a logo), given to four state-of-the-art VLMs.
- **Finding:** The four VLMs were 58.07% accurate on average, with the best (Claude 3.5 Sonnet) at 77.84%, against an expected human accuracy of 100%. Models fail when shapes overlap or sit close together, and reach near 100% when spacing is added. Linear probes show the information is in the vision encoder, but the language model fails to decode it.
- **Relation:** VLM reliability. Architecture diagrams with overlapping boxes and lines (our slide 2, "Architecture likely wrong") are exactly where VLMs fail. This supports the plausibility of a wrong architecture flag.
- **Access:** Open access https://arxiv.org/abs/2407.06581
- **Score:** 1=0 2=2 3=0 4=4 5=4 total=10/25

### Quelle, D., & Bovet, A. (2024). The perils and promises of fact-checking with large language models. *Frontiers in Artificial Intelligence*, *7*, Article 1341697. https://doi.org/10.3389/frai.2024.1341697
- **Type:** other (AI journal)
- **Verified:** Crossref 2026-10-18 (vol. 7, article 1341697, February 2024); abstract and full text read at https://www.frontiersin.org/journals/artificial-intelligence/articles/10.3389/frai.2024.1341697/full
- **Design:** LLM agents (GPT-3.5, GPT-4) wrote search queries, retrieved context, and judged 3,000 PolitiFact claims (500 per verdict category), with and without retrieved context, giving reasons and sources. Multilingual tests drew on Data Commons claims in 78 languages.
- **Finding:** Accuracy rose when the models had context. Both models were much better at spotting false claims than at confirming true ones. Without context, GPT-3.5 labelled 58.2% of all claims false. Its accuracy on claims PolitiFact rated "True" was 49.11% without context and 67.61% with it, while GPT-4 reached 71.30% and 84.92%. Both models found mixed verdicts (half-true, mostly false) hardest. Accuracy also varied with the language of the claim. The authors conclude that "caution is essential due to inconsistent accuracy".
- **Relation:** Checker reliability, as the evidence behind the brief's request. LLM checkers lean towards calling statements false, especially without context, which is the condition of a camera-only checker. A true statement from the CISO is therefore at real risk of being flagged.
- **Access:** Open access (URL above)
- **Score:** 1=0 2=3 3=1 4=5 5=5 total=14/25

---

## Notes on known papers

- **Aamir, Psaroulis, Grobler & Janicke (CHI 2026)** came up in OpenAlex searches (https://doi.org/10.1145/3772318.3791142, pp. 1–22, open access). No new reading. The four board papers above (Lowry et al., 2026; Opdenbusch et al., 2025; Proudfoot et al., 2023; Da Silva & Jensen, 2022) are the scholarly neighbours it should be cited with.
- **Lo et al. (2022), "Misinformed by visualization"**: its first author also wrote Lo & Qu (2025, entry above), which tests whether multimodal LLMs can detect the misleading-chart categories from the 2022 paper. Citing them together connects the human and machine sides of chart critique.

## Leads not taken up (for the verification pass)

- Vance, A., Lowry, M., & Sahin, Z. (2022). Taking a seat at the table: The quest for CISO legitimacy. ICIS 2022 proceedings (AIS eLibrary). On CISO legitimacy in the eyes of the board. Cited in Ruth et al. (2025); not verified.
- Sahin, Z., & Vance, A. (2025). What do we need to know about the Chief Information Security Officer? A literature review and research agenda. *Computers & Security*, 148, 104063. https://doi.org/10.1016/j.cose.2024.104063. Crossref verified; abstract not available (elided), so not entered.
- Piazza, A., Vasudevan, S., & Carr, M. (2024). Am I hired as a firefighter? ICCR 2024. https://doi.org/10.1109/ICCR61006.2024.10532977. Interviews with 24 UK CISOs: board engagement through "cyber communication" reduces CISO job stress. Abstract read; left out as peripheral.
- Prommegger, B., et al. (2026). Hiring the right CISO. *Business Horizons*. https://doi.org/10.1016/j.bushor.2026.07.003. 41 CISO interviews, three CISO sub-roles. Peripheral.
- Ereddia, L. E. (2026). Firm transparency of risk oversight. *Journal of Information Systems*. https://doi.org/10.2308/isys-2024-012. Only 57% of 8,384 proxy statements (2019 to 2021) disclosed governance of cyber risk. Peripheral.
- Masry et al. (2022). ChartQA. Findings of ACL 2022, pp. 2263–2279. https://doi.org/10.18653/v1/2022.findings-acl.177. Crossref verified; superseded for our purpose by CharXiv.
- Tanaka et al. (2023). SlideVQA (arXiv:2301.04883, AAAI 2023): 2.6k slide decks and 14.5k questions; models lag far behind humans. Could anchor "VLMs on slides" if wanted; the published DOI was not checked.

## Searches run

Crossref and OpenAlex API (bibliographic and keyword):
- "Gale Bongiovanni Slapnicar governing cybersecurity from the boardroom"; "Karanja role of chief information security officer"; "Hooper McKissack cybersecurity board"; "board of directors cyber risk oversight director"; "board cybersecurity expertise directors firm breach"
- "CISO board communication"; "board cybersecurity oversight"; "directors cyber risk literacy"; "cybersecurity governance board of directors interview"; "CISO business enabler interviews chief information security officers CHI"
- "Inexpert supervision boards oversight of cybersecurity Lowry Vance"; "What do we need to know about the Chief Information Security Officer literature review"; "Taking a seat at the table CISO Vance Lowry Sahin"
- "Quelle Bovet perils and promises"; "Bendeck Stasko GPT-4 visualization literacy"; "HallusionBench"; "ChartQA"; "CharXiv"; "MMLongBench-Doc"; "Vision language models are blind"; "How good (or bad) are LLMs at detecting misleading visualizations"; "multimodal large language models misleading charts benchmark"
- arXiv API: HallusionBench, CharXiv, VLMs are blind, MMLongBench-Doc, Lo & Qu, SlideVQA

Web search:
- Hooper McKissack abstract; "Governing cybersecurity from the boardroom" neo-institutional findings; Sahin Vance CISO review abstract; Proudfoot Cram Madnick Coden MISQE volume and issue
- NACD Director's Handbook on Cyber-Risk Oversight 2023 principles; Pearlson Hetner HBR boards cybersecurity; HBR search "pearlson board cybersecurity"; SEC final rule 33-11216 board oversight; SEC dropped Item 407(j); Proofpoint Board Perspective "see eye to eye"; Proofpoint 2024/2025 Board Perspective (none found); Gartner 88% boards business risk
- Ray-Ban Meta "look and ask" April 2024; Gemini Live camera rollout 2025; ChatGPT Advanced Voice video December 2024; Apple Visual Intelligence iOS 18.2; Microsoft Copilot in Teams meetings prompts; Otter.ai Meeting Agent 2025; Factiverse live fact-checking; Duke Reporters' Lab Squash State of the Union; ClaimBuster KDD 2017; Full Fact AI; Birdwatch Wojcik 2022

Not searched in depth: research on non-expert decision makers judging expert technical presentations outside cybersecurity. This overlaps with Agent 3 (METI, laypeople judging experts).
