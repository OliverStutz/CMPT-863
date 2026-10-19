# Project Literature Review: candidate papers

Status 2026-10-18. Step 1 of the literature review: find the papers. Six search agents (Opus, one theme each, shared brief in `search/_agent-briefs.md`) found and verified 129 new sources on top of the 96 already known from the abstract phase. Every DOI was re-checked against Crossref by a script; all resolve except arXiv and medRxiv preprints, which Crossref does not hold. Three independent selection reviewers are now ranking the combined pool; their cross-examination is appended at the end once done.

The assignment (`Instructions.md`): expand the related work into 1,000 to 1,500 words, at least 15 sources, thematic subsections that synthesize (what exists, what we know, what is unresolved, how our project differs), APA citations as in the abstract.

## Where everything is

| File | What |
|---|---|
| `search/_agent-briefs.md` | Project summary, common rules, rubric, one brief per theme |
| `search/_known-papers.md` | The 96 sources known before this search |
| `search/results-1-hci-checkers.md` | SIGCHI and HCI: AI checks or contradicts a human in front of another human (21 entries) |
| `search/results-2-reliance.md` | Human-AI decision making, automation bias, trust calibration (24 scored + 13 short) |
| `search/results-3-trust-theory.md` | Trust definition, measurement, repair, continued influence, power (20) |
| `search/results-4-professionals.md` | Observers rate professionals who use, follow, defy or are contradicted by AI (18) |
| `search/results-5-board-products.md` | Boards and CISOs, trade sources, products, VLM reliability (26) |
| `search/results-6-factchecking.md` | Fact-checking effects on the speaker, wrong flags, AI labels (20) |
| `search/_master-index.md` and `.csv` | All 129 new entries in one table, sorted by rubric score |
| `search/_reviewer-brief.md` | Brief for the three selection reviewers |
| `search/review-A.md`, `review-B.md`, `review-C.md` | The three independent selections; tally in the cross-examination section below |
| `search/reduced-view.md` | The ten papers to read first, one structured summary each, with what to check when reading |
| `search/verified-view.md` | Oliver's reading of each paper checked against the full text, with design rules collected as he goes (15 papers planned) |
| `search/trust-topic-map.md` | Opus map of the design space, 27 adjacent topics with verified anchors, 8 cross-cutting mechanisms, 10 research questions |
| `references/` | 20 full texts fetched 2026-10-18, with a README of what still needs SFU |

Each entry in the result files has: APA reference, type, how it was verified, design (N, who, manipulation, measure), finding in plain words with numbers, relation to H1 to H4, access route, and a score out of 25 on the shared rubric (trust in a human measured; an AI checks the human; power or high-stakes setting; direct use; confidence in the summary).

## What the search found, theme by theme

**1. SIGCHI and HCI checkers.** In the SIGCHI literature the checker almost always judges content (headlines, statements, arguments), and the outcome is belief in the content or trust in the AI. The strand that measures trust in a person is AI-mediated communication (Liu et al. CHI 2022; Navajas Fernández, Hancock & Jakesch CHI 2026, N = 2,000; Mieczkowski et al. CSCW 2021), and there the AI acts on the speaker's side. The closest listener-side system is Danry et al. (2020) Wearable Reasoner, which judges heard claims but measures agreement with statements, not trust in the speaker. Two newer results matter for our wrong flags: Horne & Nevo (2025) found people follow warning labels even when the labels are wrong, and Loconte et al. (2026) found people move away from an AI's verdict the more confident it is, most of all when it calls someone a liar. Serra-Garcia & Gneezy (2025) found people rely far more on an algorithm's flag shown before they watch a speaker than after, which bears on our flag timing.

**2. Reliance and automation bias.** This literature measures reliance on the AI, almost never trust in a human the AI contradicts. Usable points: wrong AI advice is followed even against context the person holds (Klingbeil et al. 2024), by experts and most of all by the least experienced (Dratsch et al. 2023, radiologists' correct readings fell from 79.7% to 19.8% under wrong AI suggestions), and explanations attached to AI output raise acceptance whether right or wrong (Bansal et al. CHI 2021; Kim et al. CHI 2025). Hou & Jung (CSCW 2021) is the one paper measuring perceived competence of the human agent when framed against an algorithm. A counter-result: Önkal et al. (2009) found people favour a human expert over a statistical method when both advise.

**3. Trust theory.** Mayer, Davis & Schoorman (1995) give ability, benevolence and integrity; Hendriks, Kienhues & Bromme (2015) give the METI scale for laypeople's trust in experts (expertise, integrity, benevolence), which is the measure Kim et al. (2026) also used. Trust repair: Kim et al. (2004) apology versus denial; Ferrin et al. (2007) reticence does worst; Sharma et al. (2023) defines partial, complete and over-repair, which H2 needs. Continued influence: Thorson (2016) belief echoes and Ecker et al. (2022) predict a residue after a good answer. Wegner et al. (1981), replicated by Letourneau & Gawronski (2024), show a hedged, question-like accusation lowers impressions of a person on its own. Schilke, Reimann & Cook (2015, PNAS): more power, less trust in the partner.

**4. Professionals and AI.** Closest to our flag: Madanay et al. (2025, N = 1,606) found patients rated a radiologist lower when the AI was more cautious than the radiologist, and not when the reverse held. Pezzo & Pezzo (2006) and Pezzo et al. (2022): an expert who defies a decision aid loses the aid's protection, and the stated reason for setting the aid aside changes the judgement. Du et al. (2026): 214 teachers graded the same paper lower when a fake AI-detector report said 87%, the nearest analogue to a wrong flag held by the person with power. Disclosure penalty: Reis, Reis & Kunde (2025, JAMA Network Open) and Schilke & Reimann (2025, 13 experiments) show that any mention of AI use lowers trust in the professional. Kempt & Nagel (2022) name the mechanism behind H2 and H3: an AI second opinion shifts the burden of proof onto the expert.

**5. Boards, trade sources, products, VLM reliability.** Opdenbusch, Hielscher & Sasse (NDSS 2025): board members do not question CISOs because they fear being exposed as novices. Proudfoot et al. (2023, 35 interviews): boards judge CISO competence by how well the CISO communicates. Lowry, Vance & Vance (2026, Management Science): non-expert directors' oversight is mostly symbolic. Da Silva & Jensen (CSCW 2022): the CISO as interpreter of a "dark art". SEC (2023) dropped the board cyber-expertise disclosure rule because directors oversee "without specific subject matter expertise". Products: Ray-Ban Meta, Apple Visual Intelligence, Gemini Live and ChatGPT video all let a listener point a camera at a slide since 2024 or 2025; Microsoft Copilot in Teams suggests the prompt "Where are some holes in [a participant's] argument?" during live meetings; Duke's Squash live checker showed more wrong matches than right ones in 2020. VLM reliability: GPT-4V scored at about the 16th percentile of humans on a visualization literacy test and read values off a chart with 13% accuracy without the data (Bendeck & Stasko 2025); GPT-4o got 47.1% of chart-reasoning questions right against 80.5% for humans (CharXiv); LLM fact-checkers lean toward calling claims false, especially without context (Quelle & Bovet 2024).

**6. Fact-checking effects.** Dobbs, Butler & Swire-Thompson (2026): a correction cut credibility for a doctor more than for a politician. Aird et al. (2018) and Swire-Thompson et al. (2020): feelings toward the speaker drop only when corrections outnumber affirmations, which makes our 3-to-2 flag ratio a design variable. Schiff, Schiff & Bueno (2025, APSR, N over 15,000): a politician who falsely calls a story fake news gains more support than one who stays silent or apologises, a counter-result to H3. Freeze et al. (2021) tainted truth and He, Horne & Nevo (2026): a wrong "false" label lowers trust in true content (about one point on a 5-point scale for an LLM checker). Drolsbach et al. (2024): flags with a reason are trusted more than bare flags.

## The closest designs and the gap

The gap statement from the abstract phase survives, but it has to be worded more carefully because of one new paper.

| Paper | What they did | How ours differs |
|---|---|---|
| **Kim, Daneshi, Draper & Sreedharan (2026), AIES** (full text checked) | N = 157; a student receives course advice from a human advisor; the advisor errs or not; an AI is absent, invoked by the advisor, or monitors and flags automatically; all end correct. Trust in the advisor measured with METI, an entity trust scale, a 100-point rating and reuse intention. Advisor errors lowered every measure; how the AI was involved changed nothing; trust stayed anchored to the human. The AI was always right (their stated limitation). | Our flags are wrong, the expert answers, the listener holds the power and the tool, and trust is rated before and after each step. |
| Chen et al. (2026), IJHCS (known) | AI second opinion disagrees with a simulated doctor; trust in the doctor falls via uncertainty and perceived laziness; no explanation given. | Listener holds the AI; expert answers; repeated ratings. |
| Madanay et al. (2025), JMIR | Radiologist recommends less than the AI; patients rate the radiologist lower. AI right by convention. | Expert answers; flags wrong; live sequence. |
| Du et al. (2026), Frontiers in Psychology | Teachers grade the same paper lower under a wrong 87% AI-detector report. | Flagged person can answer; trust, not grades. |
| Hendriks et al. (2016) | An expert's flaw pointed out by another person lowers rated expertise; integrity and benevolence depend on who discloses. | Flag from an AI, wrong, answered live. |
| Gupta et al. (CHI 2026), Danry et al. (2020) | Wearable checkers in the listener's hand; outcome is belief in claims or trust in the AI. | Trust in the speaker is the outcome. |

No study found in six themes puts a wrong AI flag in the listener's hand during a live expert briefing, lets the expert answer, and measures trust in the expert at each step. Two reviewers of the reliance and fact-checking literatures reached the same conclusion independently.

## Warnings for the hypotheses

- **H3 (denial).** Kim et al. (2004): when later evidence shows the accused was innocent, denial repairs trust better than apology. Our CISO is right by design, so a bare denial may repair rather than lower trust. Schiff et al. (2025) found an unsupported "that is fake" raised support for the accused. Loconte et al. (2026) found people resist a confident AI accusation against a person. Ferrin et al. (2007) is the better support for H3: a reply with no information (reticence) did worst after a competence question. Consider rewording H3 as "a denial without a reason restores less trust than a reasoned answer".
- **H2 (above the start).** Thorson (2016) belief echoes, Ecker et al. (2022) and Hovland & Weiss (1951) all predict a residue after a good answer. Greitemeyer & Sagioglou (2015, known) found recovery to the never-accused level but not reliably above it. Shaw, Wild & Colquitt (2003): justifications ("this is right for our company") help less than excuses. Expect recovery, not a gain, unless the pilot shows otherwise.
- **H1 (flag lowers trust).** Counter-results exist: Chae & Tewksbury (2026) found neither AI nor human fact-checks changed perceived competence of well-known politicians; Swire et al. (2017) and Barrera et al. (2020) found beliefs changed but support did not; Önkal et al. (2009) found judges favour the human expert over the algorithm. Supporting: Madanay 2025, Du 2026, Dobbs 2026, Horne & Nevo 2025, Kim et al. 2026, Prike 2025 (known).
- **H4 (amplification).** De Brito Duarte et al. (CHI 2025): the social setting magnified both good and bad responses to the same AI input. Opdenbusch et al. (2025): boards stay silent by default, so a rise in questioning is plausible. Aird 2018 and Agadjanian 2019: flags accumulate.

## Corrections to earlier assumptions

- "Zondag et al. 2025, JAMA Network Open" does not exist. The paper is Reis, Reis & Kunde (2025). The 2024 Nature Medicine paper is also Reis, Reis & Kunde, not "Reis, Reis & Wörner".
- "Yokoi et al. 2021" was not found; the nearest is Yokoi & Nakayachi (2018) on trust in medical AI, which does not fit.
- The HBR piece by Pearlson & Hetner is from 11 November 2022, not 2023. Milică & Pearlson (May 2023) and Proudfoot & Madnick (April 2026) are the other HBR pieces.
- Lowry, Vance & Vance "Inexpert supervision" is now published in Management Science 72(2), 2026, not only on SSRN.
- Mattozzi et al. has a 2026 IZA Discussion Paper version (DP 18534); cite the latest.
- Chen et al. (2026): a press-based note from the search said AI *agreement* made the doctor look lazier. The full text (`references/Chen-2026-when-AI-disagrees.pdf`, abstract and section 5) says AI disagreement increased perceived medical uncertainty and doctor laziness, and AI agreement raised the recommendation's credibility. Abstract v7 is correct; the press note is wrong. Settled 2026-10-18.

## Full texts already collected

Twenty full texts of the selected papers are in `references/` (see `references/README.md`): Kim et al. 2026 (AIES), Chae & Tewksbury 2026, Lowry et al. 2026 (published, CC BY-ND), Opdenbusch 2025, Proudfoot 2023, Aamir 2026 (PDF), Hendriks 2015, Mayer et al. 1995, Schiff 2025, Gupta 2026 (arXiv), Lu et al. 2025, Loconte 2026 (published), Bendeck & Stasko 2025, Dobbs 2026, Du 2026, Horne & Nevo 2025, Da Silva & Jensen 2022 (arXiv), Madanay 2025 (text from Europe PMC), the SEC 2023 rule, and Thorson's 2013 dissertation (not the 2016 article). The abstract-phase full texts (Chen, Strickland, Prike, Reif, von Schenk, Greitemeyer, Mickelberg) are in `../references/`.

## Still to fetch by hand (every script route is blocked)

Three of the selected papers: Ferrin et al. (2007) via SFU PsycArticles; Thorson (2016), the Political Communication article, via SFU; Hendriks, Kienhues & Bromme (2016) via SFU. Direct links are in `references/README.md`. Kim, Ferrin, Cooper & Dirks (2004) was added by Oliver on 2026-10-18 (`references/Kim-2004-apology-vs-denial-JAP.pdf`).

## Pages nobody could reach during the search (for manual fetching through SFU)

Agents report these as unverified or read from secondary sources only. Fetching the full text or abstract would settle them. Only the ones that made the reviewers' selection (Hendriks 2016, Bernstein 2025 for context) matter for the review; the rest are for completeness.

| Source | What is missing | Why it matters |
|---|---|---|
| Hendriks, Kienhues & Bromme (2016), *Studies in Communication Sciences* 16(2), 124–131, https://doi.org/10.1016/j.scoms.2016.10.003 | Abstract and N; findings from a blog summary | Closest lay-expert analogue to a flag on an expert's work |
| Bernstein et al. (2025), *NEJM AI* 2(6), https://doi.org/10.1056/AIoa2400785 | Publisher page 403; numbers from press coverage, which disagree slightly | Jurors held a radiologist liable more often when the AI had caught what they missed |
| Bonaccio & Dalal (2006), *OBHDP* 101(2), https://doi.org/10.1016/j.obhdp.2006.07.001 | Abstract (ScienceDirect 403) | Judge-advisor framing of board member, CISO and checker |
| Dzindolet et al. (2003), *IJHCS* 58(6), https://doi.org/10.1016/S1071-5819(03)00038-7 | Abstract | Telling users why automation errs maintained trust and use |
| Skitka, Mosier & Burdick (1999), *IJHCS* 51(5), https://doi.org/10.1006/ijhc.1999.0252 | Abstract | Origin of the term automation bias |
| Kim, Dirks, Cooper & Ferrin (2006), *OBHDP* 99(1), https://doi.org/10.1016/j.obhdp.2005.07.002 | Publisher abstract (repository summary only) | External accounts help less for competence violations (H2 counterweight) |
| Bies & Shapiro (1987), *Social Justice Research* 1(2), https://doi.org/10.1007/BF01048016 | Abstract via search extract only | Adequate causal accounts raise fairness ratings |
| Wegner et al. (1981), *JPSP* 40(5), https://doi.org/10.1037/0022-3514.40.5.822 | Abstract via ResearchGate extract | Innuendo lowers impressions; denial repeats the charge |
| Lewicki & Brinsfield (2017), *Annual Review of Org. Psychology* 4, https://doi.org/10.1146/annurev-orgpsych-032516-113147 | Publisher blocked; abstract via OpenAlex | Trust repair review |
| Lount & Pettit (2012), *OBHDP* 117(1), https://doi.org/10.1016/j.obhdp.2011.07.005 | Abstract; secondary sources disagree on direction | Status and trust |
| Freeze et al. (2021), *Political Behavior* 43(4), https://doi.org/10.1007/s11109-020-09597-3 | Full text (Springer login); N and procedure unchecked | Tainted truth: false warnings lower credibility of accurate news |
| Wang & Ophir (2026), *Computers in Human Behavior* 176, 108872, https://doi.org/10.1016/j.chb.2025.108872 | Publisher 403; abstract from a university portal | AI label does not make a fact-check more effective |
| Gale, Bongiovanni & Slapničar (2022), *Computers & Security* 121, https://doi.org/10.1016/j.cose.2022.102840 | Abstract from search snippets | Boards over-rely on the one cyber-literate member |
| Hooper & McKissack (2016), *Business Horizons* 59(6), https://doi.org/10.1016/j.bushor.2016.07.004 | Abstract summary from RePEc only | CISO as communicator |
| Sahin & Vance (2025), *Computers & Security* 148, 104063, https://doi.org/10.1016/j.cose.2024.104063 | Abstract | CISO literature review; may cite board-CISO trust work we missed |
| "Fact-checking in the age of AI" (2024), *Technology in Society*, https://doi.org/10.1016/j.techsoc.2024.102760 | Abstract (403) | Unknown fit |
| NACD & ISA (2023), *Director's Handbook on Cyber-Risk Oversight*, 4th ed. | PDF needs NACD registration; principles from press release | Industry source for "directors should question the CISO" |
| HBR: Milică & Pearlson (2023); Pearlson & Hetner (2022); Proudfoot & Madnick (2026) | Paywalled; summaries only (SFU has HBR in Business Source Complete) | Trade source for framing |
| Gartner (2021) board survey press release | Not fetched; 88% figure from search results | Trade number |
| Renault, Mosleh & Rand (2025), working paper on "@grok is this true" on X | Not located | Nearest real-world case of an audience-held checker |
| Joshi & Wolf (2011), JAIS-listed, Indian physicians judging peers who use diagnostic aids | No DOI found | Extends the derogation effect to peers |
| Kremer et al., managers blaming people who override algorithms | Seen only as a seminar listing | Would bear on H4 |
| From the abstract phase: Kruglanski (1970) original; Arkes et al. (2007); Wintersieck (2017); Liao et al. (2026) abstract | Paywalled or metadata only | Known candidates still summarised from secondary sources |

## Cross-examination of the three selection reviewers

Three Opus reviewers (A, B, C) worked independently from the same brief (`search/_reviewer-brief.md`), each choosing 30 papers from the 225 candidates and grouping them into subsections. Reports: `search/review-A.md`, `review-B.md`, `review-C.md`. 34 distinct papers were chosen in all. Priority: M = must cite, S = should cite, O = optional, in the order A, B, C.

### Decisive winners: chosen by all three (25)

| Paper | A B C | Role the reviewers give it | Status |
|---|---|---|---|
| Aamir et al. (2026), CHI | M M M | Board context, seed paper, H4 | In abstract v7 |
| Lowry, Vance & Vance (2026), Management Science | M M M | Non-expert oversight is symbolic | New |
| Opdenbusch, Hielscher & Sasse (2025), NDSS | M M M | Boards stay silent for fear of looking like novices; H4 | New |
| Gupta et al. (2026), CHI | M M M | Nearest wearable checker; measures belief, not trust in speaker; gap | Known |
| Strickland (1958) | M M M | Checking itself lowers trust | In abstract v7 |
| Kim, Daneshi, Draper & Sreedharan (2026), AIES | M M M | Closest design; METI measure; narrows the novelty claim | New, full text checked |
| Chen et al. (2026), IJHCS | M M M | AI disagreement lowers trust in the doctor | In abstract v7 |
| Madanay et al. (2025), JMIR | M M M | AI contradicts radiologist; asymmetric penalty; H1 | New |
| Chae & Tewksbury (2026), New Media & Society | M M M | Counter-result to H1: no change in perceived competence | Known |
| von Schenk et al. (2024), iScience | M M M | A fallible tool lowers the threshold to accuse; H4 | In abstract v7 |
| Mayer, Davis & Schoorman (1995), AMR | M M M | Definition: ability, benevolence, integrity | New |
| Hendriks, Kienhues & Bromme (2015), PLOS ONE | M M M | METI: laypeople's trust in experts, three factors | New |
| Kim, Ferrin, Cooper & Dirks (2004), J. Applied Psychology | M M M | Apology vs denial; denial works when innocent; H3 counter-prediction | Known, not yet read |
| Greitemeyer & Sagioglou (2015), PLOS ONE | M M M | Clearing restores to baseline, not above; H2 | In abstract v7 |
| Prike, Baker & Ecker (2025), Political Psychology | S M M | Fact-check lowers ratings of an unfamiliar candidate; design source | In abstract v7 |
| Thorson (2016), Political Communication | S M M | Belief echoes; counter-result to H2 "above the start" | New |
| Ferrin et al. (2007), J. Applied Psychology | S M S | Reticence does worst; best support for H3 | Known, not yet read |
| Lu et al. (2025), CHI | S S M | AI indicator weighs more when it disagrees with an expert | Known |
| Proudfoot et al. (2023), MISQE | S S S | Boards judge CISO competence by communication; H2, H3 | New |
| SEC (2023) final rule | S S S | Industry source: boards oversee without subject expertise | New |
| Microsoft Copilot in Teams, with camera assistants | S S S | Product: "Where are some holes in [a participant's] argument?" | New |
| Bendeck & Stasko (2025), TVCG | S S O | VLM misreads charts; "right by the textbook" | New |
| Reif, Larrick & Soll (2025), PNAS | O S S | AI-use penalty; why the CISO never mentions AI | In abstract v7 |
| Mickelberg et al. (2025), PLOS ONE | O S S | Retraction removes the effect in 2 of 3; rating-at-every-step caution | In abstract v7 |
| Schiff, Schiff & Bueno (2025), APSR | S S S | Counter-result to H3: an unsupported denial can help | New |

### Chosen by two of three (6)

| Paper | Chosen by | Role |
|---|---|---|
| Da Silva & Jensen (2022), CSCW | A, B (S S) | SIGCHI anchor for the CISO as interpreter of a "dark art" |
| Dobbs, Butler & Swire-Thompson (2026), Scientific Reports | A, B (S S) | A correction cuts a doctor's credibility more than a politician's; H1 |
| Du, Liu & Xian (2026), Frontiers in Psychology | A, C (S S) | Wrong AI-detector report lowers grades of an identical paper; power holder uses the flag |
| Horne & Nevo (2025), Scientific Reports | A, C (S S) | People follow warning labels even when wrong; H1 |
| Loconte et al. (2026), Computers in Human Behavior | B, C (S S) | Counter-result: people resist a confident AI accusation |
| Hancock, Naaman & Levy (2020), JCMC | A, B (O O) | Definition of AI-mediated communication; we extend it to the receiver side |

### Chosen by one (3)

Sharma, Schoorman & Ballinger (2023) (C: definition of partial vs complete repair for H2); Pezzo et al. (2022) (C: the stated reason for setting an aid aside changes the verdict); Shaw, Wild & Colquitt (2003) (B: justifications help less than excuses).

### Where the reviewers agree beyond the paper list

- **Subsections.** All three converge on five: (1) boards as non-expert overseers of the CISO, with regulation and products; (2) AI checkers in HCI judge content, people follow them even when wrong, VLMs misread charts; (3) a check, correction or AI disagreement lowers the standing of the person checked, with Chae & Tewksbury as the counter-result; (4) trust in an expert, its three dimensions, its repair, and the residue after a good answer; (5) the gap, or a short section on how a checker changes what its holder does (H4).
- **All 8 abstract v7 papers stay**, two at lower priority (Reif, Mickelberg).
- **Novelty claim must narrow** because of Kim et al. (2026): not "first to measure trust in an expert after an AI flag", but the combination of a wrong flag, a checker in the hands of the more powerful non-expert, the expert's answer as the repair step, and trust rated at each step.
- **H3**: all three recommend "a denial without a reason restores less trust than a reasoned answer" instead of "lower than after the flag" (Kim et al. 2004; Schiff et al. 2025).
- **H2**: all three recommend treating "above the start" as exploratory (Greitemeyer & Sagioglou 2015; Thorson 2016).
- **H4b (amplification)**: all three find no direct support in the files and recommend stating it as exploratory; Kim et al. (2026) found the AI's involvement did not change the trust penalty, which cuts against it. A and C also note H4a compares two different acts (asking the CISO vs saying one has questions).

### Must-read before writing, after removing what the abstract phase already checked

The reviewers did not see `../references/reviews/v6-related-work-validation.md`, which already verified Strickland, Prike, Chen, Greitemeyer, Mickelberg, Aamir, von Schenk and Reif claim by claim against full texts (72 claims, none wrong). Chen's laziness direction is settled above. What remains:

1. Kim, Ferrin, Cooper & Dirks (2004): was the denial bare or with an account; what showed innocence. Decides H3. SFU PsycArticles.
2. Ferrin et al. (2007): exact wording of reticence; competence case. SFU PsycArticles.
3. Chae & Tewksbury (2026): how competence was measured; whether targets were well-known politicians. Open access at the DOI.
4. Gupta et al. (2026): confirm trust was measured in the AI's information, not the speaker (gap statement rests on it). arXiv 2603.27289.
5. Lu et al. (2025): confirm "larger when it disagreed with a human expert" and what the expert was. ACM DL.
6. Kim et al. (2026), AIES: cell sizes, whether the no-error monitoring condition moved trust (bears on our "No issues found" slides). Full text in `references/`.
7. Madanay et al. (2025): which direction of discrepancy maps onto our flags. PMC.
8. Schiff et al. (2025): size of the advantage over silence; text vs video. Open access.
9. Opdenbusch et al. (2025) and Lowry et al. (2026): confirm the "IT novices" and "mostly symbolic" wording. Open access.
10. Loconte et al. (2026): whether participants judged statements or persons. Open access.
11. Microsoft Copilot support page: re-check the quoted prompt close to submission.

## Next steps

1. Read the must-read list above (open-access PDFs are being collected in `references/`; the SFU-only ones are in the manual-fetch table).
2. Draft the review in the five agreed subsections, 1,000 to 1,500 words, APA, from the 25 decisive winners plus the 2-of-3 papers where a subsection needs them.
3. Independent cross-check of the draft against the full texts before submission.
