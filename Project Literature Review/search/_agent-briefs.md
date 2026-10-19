# Briefs for the literature-search agents (Project Literature Review)

Written 2026-10-18. Six agents search in parallel, one theme each. Each writes its results to `search/results-<theme>.md` in this folder. A verification pass and a cross-examination follow.

## The project in short

A CISO (chief information security officer) presents a five-slide cybersecurity briefing to a company board and asks for a budget. The participant plays a board member. The board member holds the formal power and the budget; the CISO holds the expertise the board depends on. Boards are non-experts making high-stakes decisions.

Two groups, about 40 adults online, 10 minutes. The **checker group** has a simulated vision-language-model (VLM) content checker: a panel shows the camera image of each slide, "No issues found" on slides 1 and 5, and on slides 2 to 4 a flag such as "Architecture likely wrong" or "Numbers do not add up". Every flag is wrong for this company (right by the textbook, wrong in context). After each flag the participant rates trust in the CISO (0 to 10), then says nothing, asks the CISO to explain, or challenges the CISO with the checker's finding; if nothing, the board chair asks a routine question. The CISO answers: two answers give the company's reason, one is a denial with no reason. Trust is rated again. The **neutral group** follows the same briefing and hears the same answers, without the checker. The participant may be asked to rate trust as competence, benevolence and integrity.

Hypotheses: **H1** trust drops after a flag. **H2** a reasoned answer restores trust, possibly above the start. **H3** a denial with no reason lowers trust further. **H4** checker participants question the CISO more, and the same answer then moves trust further in either direction (amplification).

Research question: when a board member uses a VLM content checker during a CISO's briefing, how does their trust in the CISO change from start to end, compared with a board member without a checker?

## What the literature review must do (course instructions)

Expand the related work into a 1,000 to 1,500 word review with thematic subsections, at least 15 sources, APA citations. Scholarly research first; industry articles, books, trade publications and existing products where appropriate. The review must say what exists, what we know, what is unresolved, and how this project differs.

## Common rules for every agent

1. **Read `search/_known-papers.md` first.** Those papers are already known. Do not spend effort re-finding them. If you meet one and have a materially new reading (a finding we got wrong, a detail that matters for H1 to H4), add it in a short "Notes on known papers" section at the end of your file. Otherwise skip it.
2. **Prefer SIGCHI venues** (CHI, CSCW / PACM HCI, UIST, IUI, DIS, CHI EA, TOCHI, IJHCS, FAccT, CHI PLAY is unlikely) for the HCI themes, then other HCI journals, then psychology, management, economics, communication and medicine journals. For the industry theme, trade publications, reports and products count.
3. **Verify every entry.** For a paper, fetch `https://api.crossref.org/works/<DOI>` or the publisher page and check title, authors, year, venue, volume, issue, pages or article number. Read the abstract (publisher page, PubMed, PMC, arXiv, Semantic Scholar, OpenAlex, the ACM DL) before writing the finding. Never invent a DOI. If a DOI cannot be resolved, write "DOI unverified" and give the URL you used. For a product or industry article, give the URL and the date you accessed it.
4. **Plain wording.** Name what was measured, on whom, and with what result, with numbers. Write "participants rated the candidate as less competent" and not "d = 0.6" alone; give both where the abstract gives both. No statistics terms as a stand-in for a finding.
5. **No em-dashes** anywhere. Use commas, full stops or colons.
6. **Do not pad.** 10 to 20 strong entries beat 40 weak ones. An entry earns its place if a marker reading our review would expect it, or if it supplies a finding, a measure, a design element or a counter-result that our review needs.
7. **Score every entry** on the rubric below so the three-way cross-examination can compare across agents.
8. **Flag access.** Say whether the full text is open access (URL) or needs the SFU library.
9. Use web search and fetch freely. Do not use your memory alone for any finding; fetch the abstract.

## Rubric (0 to 5 each, total out of 25)

1. Trust in a *human* (or credibility, competence, trustworthiness of a person) is a measured outcome.
2. An AI or automated tool checks, verifies, or contradicts that human, or a human who uses such a tool is evaluated.
3. Setting: a non-expert judges an expert, a power asymmetry, or a high-stakes professional context (board, medicine, law, governance).
4. Direct use for our review: supplies a theme anchor, a definition, a measure, a design element, a finding for H1 to H4, a counter-result, or the gap statement.
5. Confidence in your summary: 5 = full text or detailed abstract read, 3 = abstract only, 1 = title and metadata only.

## Output format (strict)

File header, then entries grouped under the sub-themes you find, then "Notes on known papers", then "Searches run" (the queries you used, so gaps can be seen).

Each entry:

```
### Author, A. B., Author, C. D., & Author, E. F. (Year). Title in sentence case. *Journal or Proceedings*, *volume*(issue), pages or Article number. https://doi.org/...
- **Type:** SIGCHI | other HCI | psychology | management or economics | communication | medical | industry or trade | product | book
- **Verified:** Crossref <date> | publisher page <URL>; abstract read at <URL>
- **Design:** N and who, what was manipulated, what was measured on whom and on what scale
- **Finding:** plain words with numbers
- **Relation:** sub-theme; which of H1 to H4, or definition, measure, context, product, gap; one sentence on how our study differs
- **Access:** open access <URL> | SFU library
- **Score:** 1=_ 2=_ 3=_ 4=_ 5=_ total=_/25
```

APA 7 reference style as in the example above (authors with initials, year, sentence-case title, journal in italics with volume in italics, issue in parentheses, pages, DOI as URL). For conference papers: In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (Article 1599, pp. 1–22). Association for Computing Machinery. https://doi.org/...

---

## Agent 1: SIGCHI and HCI, AI tools that check, verify or contradict a human in front of another human

Scope: HCI studies in which an automated or AI tool fact-checks, verifies, flags, or gives a second opinion about what a human says or shows, and the effect on how other humans judge that person or their claims. Includes real-time fact-checking during speech, debates, meetings or presentations; wearable, AR or camera-based (VLM) verification assistants; LLM second opinions; AI credibility indicators next to a human expert's view; AI-mediated communication where perceived AI involvement changes interpersonal trust; systems papers (UIST, CHI) that built live checkers and what their user studies measured. Years 2013 to 2026, with weight on 2022 to 2026.

Known already (skip): Gupta et al. CHI 2026, Lu et al. CHI 2025, Jia et al. CHI 2026, Liao et al. 2026, Garrett & Weeks CSCW 2013, Jakesch et al. CHI 2019, Hohenstein & Jung 2020, Hohenstein et al. 2023, Detjen et al. CHI 2025, LiveFC WSDM 2025, Chen et al. 2026, DeVerna et al. 2024.

Look especially for: ACM DL searches "real-time fact-checking" "live fact-checking" "fact-checking assistant" "verification assistant" "AI second opinion" "conversational fact-checking" "wearable reasoner" "AR fact check" "smart glasses assistant trust" "AI-mediated communication trust" "human-AI-human" "LLM disagreement expert" "credibility indicator expert"; Danry, Pataranutaporn and Maes (MIT Media Lab) wearable reasoning work; Hancock, Naaman and Levy "AI-mediated communication" (JCMC 2020); Mieczkowski et al. CSCW 2021; Sundar's machine heuristic (CHI 2019) and MAIN model; any CHI, CSCW, IUI or UIST 2024 to 2026 paper on LLM fact-checking of spoken or visual content with a user study. Also search for any study where the *audience* or *recipient* holds the checker rather than the speaker.

Output file: `search/results-1-hci-checkers.md`.

## Agent 2: Human-AI decision making, reliance, automation bias and trust calibration

Scope: why a non-expert follows an AI flag, and what explanations, confidence displays and AI errors do to reliance and to judgement of a human advisor. Includes human-AI complementarity and over-reliance (Bansal et al. CHI 2021; Buçinca, Malaya & Gajos CSCW 2021; Vasconcelos et al. CSCW 2023; Zhang, Liao & Bellamy FAT* 2020; Schemmer et al. IUI 2023; Lai & Tan FAT* 2019; Lai, Chen, Smith-Renner, Liao & Tan FAccT 2023 survey), algorithm aversion and appreciation when AI and a human expert disagree (Hou & Jung CSCW 2021; Chong et al. 2022; Logg, Minson & Moore 2019 is known), automation bias reviews (Parasuraman & Manzey 2010; Goddard, Roudsari & Wyatt 2012), trust in automation (Lee & See 2004; Hoff & Bashir 2015), human trust in AI reviews (Glikson & Woolley 2020; Vereschak, Bailly & Caramiaux CSCW 2021 on how to evaluate trust in AI-assisted decision making), judge-advisor system research where one advisor is an algorithm (Bonaccio & Dalal 2006 as the base), and any 2024 to 2026 CHI/CSCW/IUI work on LLM advice that is confidently wrong and what users do with it. Also any work on how a *human's* standing changes in the eyes of a judge when an algorithm disagrees with the human (advice conflict).

Known already (skip): Dietvorst et al. 2015, Logg et al. 2019, Jian et al. 2000, Lu et al. CHI 2025, Detjen et al. CHI 2025.

Output file: `search/results-2-reliance.md`.

## Agent 3: Interpersonal trust: definition, measurement, repair, and continued influence

Scope: psychology and management literature we need to define and measure trust in the CISO and to ground H2 and H3. (a) Definition and dimensions: Mayer, Davis & Schoorman 1995 (ability, benevolence, integrity), McAllister 1995 (cognition- and affect-based trust), Colquitt, Scott & LePine 2007 meta-analysis, Rousseau et al. 1998, Hendriks, Kienhues & Bromme 2015 METI (laypeople's trust in experts: expertise, integrity, benevolence), Bromme and colleagues on laypeople judging experts. (b) Trust repair after a competence or integrity question: Kim, Ferrin, Cooper & Dirks 2004 (apology vs denial), Kim, Dirks, Cooper & Ferrin 2006, Ferrin, Kim, Cooper & Dirks 2007 (reticence), Tomlinson & Mayer 2009, Lewicki & Brinsfield 2017 review, Dirks, Lewicki & Zaheer 2009, Kramer & Lewicki 2010, Schweitzer, Hershey & Bradlow 2006 (known, skip). (c) Explanations and accounts: Shaw, Wild & Colquitt 2003 meta-analysis on explanations, Bies & Shapiro 1987 causal accounts, Sitkin & Bies 1993. (d) Continued influence and belief echoes: Thorson 2016 belief echoes (Political Communication), Ecker et al. 2022 Nature Reviews Psychology, Lewandowsky et al. 2012. (e) Power and trust: Schilke, Reimann & Cook 2015 PNAS, Kruglanski 1970 (known), any work on supervisors or principals checking agents and the trust consequence beyond Strickland. (f) Source credibility of experts judged by non-experts: Hovland & Weiss 1951, Pornpitakpan 2004 review, Hendriks et al. 2016.

Known already (skip unless you add a new reading): Strickland 1958, Kruglanski 1970, Greitemeyer & Sagioglou 2015, Mickelberg et al. 2025, Ecker & Rodricks 2020, Mayer & Davis 1999, Schweitzer et al. 2006, Kramer 1999, Sitkin & Roth 1993, Fein 1996, Falk & Kosfeld 2006, Schweitzer Ho & Zhang 2018.

Output file: `search/results-3-trust-theory.md`.

## Agent 4: Evaluation of professionals who use, or are contradicted by, AI (medicine, law, work)

Scope: studies in which observers (patients, clients, colleagues, the public) rate a professional who uses an AI or decision aid, follows or defies it, or is contradicted by it. Medicine: Arkes, Shaffer & Medow 2007 and Shaffer et al. 2013 (known), Pezzo & Pezzo 2006 (physician evaluation after error with or without a decision aid), Palmeira & Spassova 2015 (known), Zondag et al. 2025 JAMA Network Open (public perception of physicians who use AI), Reis, Reis & Wörner 2024 Nature Medicine (believed AI involvement lowers trust in medical advice), Yokoi et al. 2021, Kempt & Nagel 2022 J Med Ethics (second opinions and burden of proof), Longoni et al. 2019 (known), any 2023 to 2026 experiment on "doctor uses AI" or "doctor overrides AI" and patient trust. Work and law: Reif et al. 2025 (known), any study on colleagues or managers who use AI being rated, "AI use disclosure" penalties, lawyers or advisors using AI (e.g., 2024 to 2026 experiments on disclosed AI use and perceived competence), and studies on experts who *push back* on AI and how observers judge that (hybrid judgement, Palmeira & Romero Lopez 2023 known). Also the flip side: people who are *flagged* by AI (AI detectors in education: Gorichanaz 2023 known; look for newer work on students accused by AI detectors and teacher-student trust, 2024 to 2026).

Output file: `search/results-4-professionals.md`.

## Agent 5: Boardroom, cybersecurity governance, CISO-board communication, products and VLM reliability

Scope, three parts. (a) Scholarly: board-level cybersecurity governance and the CISO-board relationship (Gale, Bongiovanni & Slapnicar 2022 Computers & Security; Karanja 2017; Hooper & McKissack 2016; any paper on how boards judge cyber-risk reports, cyber literacy of directors, CISO credibility or reporting to the board, 2015 to 2026), and research on non-expert decision makers evaluating expert technical presentations. (b) Industry and trade: NACD Director's Handbook on Cyber-Risk Oversight (latest edition), Harvard Business Review on boards and cybersecurity (Pearlson & Hetner 2023 and others), the SEC 2023 cybersecurity disclosure rule (board oversight), Gartner, Deloitte, PwC, Heidrick & Struggles or Proofpoint "CISO and board" surveys (give the report title, year, URL, and the one number that matters). (c) Existing products that let a listener point a camera at content and have it checked or explained, and live AI fact-checkers: Ray-Ban Meta glasses with Meta AI ("look and ask"), Google Gemini Live with camera, ChatGPT with vision or voice, Apple Visual Intelligence, Microsoft Copilot in Teams meetings, Otter.ai meeting agent, Factiverse, Full Fact AI, ClaimBuster, Duke Reporters' Lab Squash, X Community Notes (Wojcik et al. 2022 for the scholarly side). For each product: what it does, since when, source URL, access date. (d) VLM reliability: evidence that vision-language models misread charts, numbers or diagrams and give context-free verdicts (chart QA benchmarks with error rates, hallucination studies, Quelle & Bovet 2024 on LLM fact-checking reliability, any 2025 to 2026 evaluation of VLMs on slides or documents). This part grounds "a flag can be right by the textbook and wrong for this company".

Output file: `search/results-5-board-products.md`.

## Agent 6: Fact-checking effects on the person checked and on the audience, since 2015, with AI labels

Scope: the fact-checking and corrections literature, with emphasis on what a correction or flag does to evaluation of the *speaker* and on AI versus human fact-checkers. Meta-analyses (Walter, Cohen, Holbert & Morag 2020 "Fact-checking: A meta-analysis of what works and for whom"; Walter & Murphy 2018; Chan et al. 2017), belief echoes (Thorson 2016), fact-check effects on candidate evaluation (Wintersieck 2017 known; Fridkin et al. 2015 known; Nyhan et al. 2020 known; Prike et al. 2025 known; Chae & Tewksbury 2026 known), fact-checking during live debates or speeches with second-screen or on-screen verdicts (2016 to 2026), AI-labelled fact-checks and their credibility (2023 to 2026, e.g., studies on "AI-generated fact-check" labels, LLM corrections, Community Notes vs AI), corrections that are themselves wrong (false corrections, overcorrection, what a wrong flag does to trust in the flagged source and in the checker), and the effect of a fact-check on the checked party's behaviour (Nyhan & Reifler 2015 known; look for replications or newer field studies). Also the "tainted truth" or "implied truth" effects (Pennycook et al. 2020 implied truth; Freeze et al. 2021 "Fake claims of fake news" on warnings that are wrong) because our flags are wrong.

Known already (skip): Nyhan & Reifler 2015, Mattozzi et al., Wintersieck 2017, Fridkin et al. 2015, Garrett & Weeks 2013, Nyhan et al. 2020, Prike et al. 2025, Hoes et al. 2024, Chae & Tewksbury 2026, DeVerna et al. 2024, Appling et al. 2022.

Output file: `search/results-6-factchecking.md`.
