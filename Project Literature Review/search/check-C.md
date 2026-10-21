# Check C of Literature Review - Draft v1 (2026-10-19)

Checker C, working alone from `_check-brief.md`. I did not read check-A.md or check-B.md. Every claim with a full text in `references/` (or in `../../references/` from the abstract phase) was checked against that text with pdftotext; the rest against verified-view, the two read files, results-5 and the fact-check record.

## A. Verdict

**Minor fixes.** All 55 cited claims hold in substance, none is wrong, and the four corrections in the fact-check record are all in their corrected form. Four sentences say a little more than the source does (Du, Aamir, Sharma, Bendeck). The other fixes are about readability: three acronyms are never spelled out, the Pezzo sentence has no lead-in, section 3 has no "how we differ" sentence, and the review ends on the design rather than on the contribution. The body is 1,472 words, so every fix has to be word-neutral or paid for by a cut.

## B. Claim table

Classes: V = verified, W = wrong, U = unsupported, O = overstated. n/a = the authors' own framing, not a cited claim.

### Opening and section 1, Boards

| # | Claim (short) | Source checked | Class | Correction |
|---|---|---|---|---|
| 1 | SEC chose not to require disclosure of board cyber expertise; quote "directors with broad-based skills ... without specific subject matter expertise" | SEC 2023 PDF p. 85 ("we are not adopting proposed Item 407(j)"; quote verbatim) | V | |
| 2 | Lowry: non-expert oversight "mostly symbolic" even when they did what experts did | Lowry PDF abstract p. 783 | V | |
| 3 | Lowry: an expert director said for most boards "all they can do is assess the credibility of the executives that are put in front of them" | Lowry PDF, D-9 quote (section 4) | V | |
| 4 | Proudfoot: perceived CISO competence "heavily determined by the CISO's ability to effectively communicate" | Proudfoot PDF p. 240 | V | |
| 5 | Opdenbusch: directors "too shy to ask questions since they do not want to get embarrassed for their missing knowledge" | Opdenbusch PDF p. 6 (CISO4, EX6, NED3, NED4) and p. 13 discussion | V | Four participants say it; the authors generalise it in their own discussion, so the general wording is acceptable. |
| 6 | Opdenbusch: most interviewees named budget as the board's only cyber decision | Opdenbusch PDF p. 5 ("Most participants (NED1,3,4; EX1,3-6; CISO1-3) ... the only (abstracted) decision") | V | 11 of 18. |
| 7 | Aamir: directors "pushed back more confidently" where independent reviews were routine | Aamir PDF pp. 7, 13, 14 ("directors reported greater confidence in challenging management perspectives"; "It gave directors confidence to challenge") | O | This is self-reported confidence in interviews, not observed pushing back. Abstract v8 line 61 already has the right wording ("felt more confident challenging"). See fix 2. |
| 8 | Opdenbusch: "the CISOs are the ones in power", because they decide what reaches the board | Opdenbusch PDF p. 2 contribution III and p. 12 discussion | V | |
| 9 | Lowry trace the deference to low expertise instead | verified-view entry 8 | V | |
| 10 | Formal power with the board, informational power with the CISO | own reading (verified-view rule 7) | n/a | Marked "We read this as", correctly. |
| 11 | Schilke: the party with better alternatives trusted its partner less | Schilke PDF p. 12951, Study 1 (alternatives manipulation), Studies 2 to 4 | V | |
| 12 | Schilke: same starting level, trust fell further during the negotiation | Schilke PDF p. 12951 (2.59 and 2.59 at T1; 2.12 against 2.46 at T2; interaction F = 39.47) | V | |
| 13 | Their power was having alternatives, not authority | Schilke PDF p. 12950 (structural dependence) | V | |
| 14 | The board evidence is from interviews, with no tool and no trust measure | Opdenbusch, Lowry, Proudfoot, Aamir full texts | V | Lowry adds a CISO survey and a proxy-statement count, but neither has a tool or a trust measure, so the claim holds. |

### Section 2, Checkers

| # | Claim (short) | Source checked | Class | Correction |
|---|---|---|---|---|
| 15 | Ray-Ban Meta glasses answer questions about what the wearer is looking at (Meta, 2024) | results-5 lines 142 to 150 | V | |
| 16 | Apple Visual Intelligence: "an explanation of a diagram from class notes" (Apple, 2024) | results-5 line 151 | V | The quote is the ChatGPT route inside Visual Intelligence; acceptable. |
| 17 | Copilot in Teams suggests "Where are some holes in [a meeting participant's] argument?" | results-5 line 165 | V | "the private prompt" is loose: the pane is private, not the prompt. See fix 9. |
| 18 | Wearable Reasoner was right on 63.5 percent of statements when running live | Danry PDF section 5.1 (63.5% live against 85.11% offline) | V | |
| 19 | Danry et al. rated recorded statements, not a person | Danry PDF section 4.3.3 ("A speaker played pre-recorded arguments"; participants rated each argument) | V | Strictly it was the participants who rated; the wording is fine for a review. |
| 20 | Squash: 14 relevant fact-checks and 20 irrelevant ones; the lab began building an interface for a human editor to approve each pop-up | fact-check record item 3; results-5 line 176 | V | Corrected wording is used. |
| 21 | GPT-4V at about the 16th percentile of humans on visualization literacy | Bendeck PDF section 4 ("around the 16th percentile of humans") | V | |
| 22 | GPT-4V showed some knowledge of design best practice | Bendeck PDF section 6 ("basic knowledge of visualization design best practices") | V | |
| 23 | GPT-4V answered only 13 percent of questions that needed a value read off a chart when it lacked the data | Bendeck PDF section 5.3 (13% on "non-visual lookup" questions of the CQA set without data; 23.1% on VLAT "retrieve value") | O | 13 percent is one question category, not all value-reading questions. See fix 4. |
| 24 | None of this work reports what a check does to trust in the speaker | Danry, Adair, Bendeck, results-5 | V | |

### Section 3, Wrong flags and the person flagged

| # | Claim (short) | Source checked | Class | Correction |
|---|---|---|---|---|
| 25 | A wrong AI warning on a true headline cut trust by 0.30 on a scale from −1 to +1 | Horne PDF Table 2 (error labeling −0.298, p < .001; label "Warning: An AI Tool Says This Story is False") | V | |
| 26 | Lu: AI indicator had a larger effect on accuracy than an expert's vote | Lu PDF abstract and section 6 ("exploratory analyses suggest ... larger than that of the expert's") | V | Exploratory, and the "expert" was an inserted vote; "vote" says that. |
| 27 | Lu quote: people "lack the capability to differentiate correct and incorrect AI predictions" | Lu PDF section 6 | V | |
| 28 | Du: 214 teachers; identical paper graded 1.8 points lower on a 10-point scale | Du PDF Table 1 (6.83 to 5.02 without red highlighting) | V | Between subjects, so about half the 214 saw the high rate; "made 214 teachers grade" is loose. Folded into fix 1. |
| 29 | Du: red highlighting made the report seem more credible and the drop larger | Du PDF section 4.3 and Table 2 (1-to-10 interaction F = 3.229, p = .074, "not interpreted as evidence of an interaction"; 0-to-100 score interaction F = 32.68, p < .001; report perception main effect of highlighting significant) | O | On the 10-point scale the sentence quotes, the larger drop is only descriptive and the authors say not to read it as an effect. It holds for the 100-point score. See fix 1. |
| 30 | von Schenk: people who chose to consult the algorithm accused in almost 85 percent of cases when it said "lie" | von Schenk full text (84.91%; "climbed to almost 85%") | V | |
| 31 | Strickland: more supervisors trusted a rarely checked worker than an often checked one | Strickland PDF pp. 206, 215 (Table 3; finding 2) | V | |
| 32 | Even checks that find nothing can cost trust | Strickland PDF (both workers complied; checks found the work in order) | V | |
| 33 | Prike: a human fact-check lowered an unfamiliar candidate's competence and morality ratings | Prike preprint pp. 18 to 19 (fact-check contrasts d = 0.90 to 1.14 for competence, d = 1.11 for morality) | V | |
| 34 | Madanay: 1,606 adults; less agreement with the radiologist who recommended less testing than the AI, rated lowest of four; recommending more cost nothing | Madanay txt Table 2; fact-check item 2 (413 + 401 + 408 + 384 = 1,606) | V | Corrected wording is used. |
| 35 | Chen: AI disagreement made a doctor seem lazier and lowered trust in him | Chen PDF abstract, section 4.5 | V | The doctor was an LLM presented as human, "Dr. Alex", referred to as "his"; the trust effect is indirect, through laziness and uncertainty. Fine at review level. |
| 36 | Reif: employees getting help from AI rated lazier and less competent than those who got the same help elsewhere | Reif full text, Study 2 ("than people who get help from other sources") | V | |
| 37 | Kim 2026: advisor's error lowered trust, AI involvement changed nothing, so the loss follows the error | Kim 2026 PDF pp. 1520 to 1521; fact-check item 4 | V | Corrected wording is used. |
| 38 | Loconte: people resisted confident AI accusations of lying when accuracy and confidence were on screen | Loconte CHB PDF abstract and section 2; verified-view entry 17 | V | |
| 39 | Results on well-known politicians are not carried over | author rule | n/a | The required statement is present. |
| 40 | In all of these studies the expert stays silent and trust is rated once | Kim 2026 p. 1517 ("Post-task"), Madanay, Chen, Prike, Strickland, Du | V | Horne, Lu and Loconte have no expert, so "expert" is loose for them; acceptable. |

### Section 4, Trust and repair

| # | Claim (short) | Source checked | Class | Correction |
|---|---|---|---|---|
| 41 | Mayer: willingness to be vulnerable; ability, benevolence, integrity; trust leads to risk taking | Mayer PDF pp. 712, 717, 724 to 726; verified-view entry 11 | V | |
| 42 | METI built on the same factors for laypeople judging an expert | Hendriks PDF (expertise, integrity, benevolence, from Mayer) | V | |
| 43 | Kim 2004: after a hearsay competence accusation, apology repaired better than denial; their denial came with an account | Kim 2004 PDF pp. 107 to 110 ("hearsay"); verified-view entry 2 | V | |
| 44 | Ferrin: a reply that neither confirmed nor denied did no better than the weaker of the other two | Ferrin PDF abstract; verified-view entry 10 | V | |
| 45 | Greitemeyer: clearing restored the researcher to the never-accused level, not reliably above | Greitemeyer PDF abstract ("similar in the exoneration and the control conditions") | V | |
| 46 | Mickelberg: rating after every step may draw attention to the latest information | Mickelberg PDF pp. 2 to 3 | V | Hedged, as in the source. |
| 47 | Sharma: complete repair = return to start; beyond it "further development of trust" | Sharma web text, "What Are the Limits of Trust Repair?" | V | |
| 48 | Sharma: "Most studies they review find that repair falls short of the start" | Sharma web text, introduction ("most of the studies have shown that repair tactics do not return the trust levels to the pre-transgression levels (Kim, 2018)") | O | The point is Kim (2018), cited by Sharma, not a finding of their own review. Their review also says baselines are seldom measured. See fix 3. |
| 49 | Sharma: denial works less well after competence than after integrity violations | Sharma web text, "Denial" ("more effective for integrity-based violations than for competence-based violations") | V | |
| 50 | In these studies the accuser is human | Kim 2004, Ferrin, Greitemeyer, Mickelberg | V | |
| 51 | Kempt and Nagel: when an AI contradicts the expert in charge, the burden of proof moves to the expert; Chen et al. take it up | Kempt PDF pp. 222, 226, 227; Chen PDF p. 11; fact-check item 1 | V | Corrected attribution is used. |
| 52 | The expert cannot settle it by reasons because the machine does not respond; a second human settles it | Kempt PDF p. 226 ("not receptive for diagnostic reasons"; "a second human opinion ought to be requested") | V | |
| 53 | Jongsma and Sand reply that agreement carries the same burden | read-danry-kempt-dasilva.md (accepted manuscript read in full) | V | |
| 54 | Pezzo: unconsulted aid, bad outcome; rated higher when he believed he knew what it would say than when he did not trust it; never below no-aid | read-pezzo-schilke-sharma.md (abstract only) | V | Abstract level only. The wording is the safe clause from the read file. |

### Section 5, Gap

| # | Claim (short) | Source checked | Class | Correction |
|---|---|---|---|---|
| 55 | Kim 2026 closest: an AI monitored an advisor in front of a dependent student; METI; N = 157 | Kim 2026 PDF pp. 1515 to 1517 | V | "Monitored" fits only the reactive conditions; in the proactive ones the advisor called the AI (fact-check C6). Optional fix 11. |
| 56 | Their AI was always right, the advisor never replied, the student had no tool and no power | Kim 2026 PDF section 4.1 | V | |
| 57 | Gupta: nearest real-time wearable checker; measured belief in claims and trust in the AI, never the speaker | Gupta arXiv PDF abstract, Table of measures ("Trust in AI"), section 6.3 | V | |
| 58 | Sharma: 8 of 94 empirical studies used any within-subject design; repair needs three ratings | Sharma web text, "Measurement Issues and Recommendations" | V | The paper says 97 empirical articles in one place and 94 in another; 94 is the one that goes with the count. |

**Totals over the 55 cited claims (rows 10 and 39 are n/a; row 28 is counted once and handled under fix 1):** verified 51, wrong 0, unsupported 0, overstated 4 (rows 7, 23, 29, 48). Rows 1 to 58 minus the two n/a rows and row 28 make 55.

Fact-check record: all four wordings are correct in the draft. Kempt and Nagel are given the burden-of-proof argument with Chen taking it up (row 51). Madanay is "lowest of the four conditions", not "worst provider" (row 34). Squash is 14 relevant and 20 irrelevant, and the lab "began building an interface" (row 20). Kim 2026 bounds the claim rather than supporting it (row 37) and is not cited for the AI-penalty claim.

## C. Compliance

### Item 2, Instructions.md

| Requirement | Status | Evidence |
|---|---|---|
| Under 1,500 words | Met | My count of the body, without the title, headings, references and writer's notes: opening 47, boards 278, checkers 226, wrong flags 360, trust and repair 375, gap 186, **total 1,472**. This matches the writer's count. There is 28 words of headroom, so the fixes below are costed. |
| At least 15 sources | Met | 34 cited, all in the list. Scholarly, industry (Meta, Apple, Microsoft, Adair) and regulation (SEC) are all present, as the instructions ask. |
| Thematic subsections | Met | Five themed sections plus an opening. |
| Synthesis, not one-by-one | Mostly met | Sections 1, 3 (paragraphs 1 and 3) and 5 synthesise. Three runs are still one paper per sentence: section 2, sentences 2 to 4 (Meta, Apple, Microsoft); section 3, paragraph 2 (Strickland, Prike, Madanay, Chen, Reif); section 4, paragraph 2 (Kim 2004, Ferrin, Greitemeyer, Mickelberg, Sharma). Each run opens with a claim sentence ("The person flagged pays"), which makes them acceptable. |
| Each theme: exists, known, unresolved, how we differ | Mostly met | Section 1 has all four (differ: "we put an independent check in the board member's hands"). Section 2 has all four (differ: "Our flags are drafted with a real VLM"). Section 3 has exists, known and unresolved ("expert stays silent and trust is rated once") but no explicit "how we differ"; only the asides "so our CISO never mentions AI" and "our flags concern the slide". Section 4 has all four (unresolved: "the accuser is human"; differ: "Our own hypothesis", "both our groups rate at the same steps"). See fix 7. |
| Gap and contribution clear by the end | Partly met | The gap is clear ("No study combines these pieces", with Kim 2026 contrasted on three points). The last sentence describes the design but does not say what we will learn, and `_review-design-v2.md` asked for a closing contribution sentence. See fix 8. |
| APA as in Abstract v8 | Met, with two small issues | The format matches v8. "D. Kim" and "P. H. Kim" initials follow APA 7 and differ from v8 for a good reason (noted by the writer). Gupta et al. has no page range and a title-case "Feeling the Facts" (fix 10). |

### Item 3, Structure fidelity to `_review-design-v2.md`

| Deviation | Harm |
|---|---|
| Opening has one scenario sentence, not two | None. |
| Section 1: Lowry's "we ask the right questions", question lists and 40 percent filtering dropped; Proudfoot's "always on audition", 10 to 15 minutes and "don't show up the CEO" dropped | Low. The anchor quotes remain. Filtering was the reason a flag seems plausible; that link is lost, but it is not needed for the argument. |
| Section 1: "How we differ" names trust in the CISO but not the act of asking (H4a) or the budget decision | Low. Budget comes back in section 4 through Mayer ("here the budget decision"); asking does not come back anywhere in the review. |
| Section 1: Schilke used for "trust fell further" and "power was alternatives, not authority", not "power is a variable" | None. This is the correction in read-pezzo-schilke-sharma.md. |
| Section 2: Google Gemini Live dropped | None. It had only a secondary source. |
| Section 2: Danry used only for 63.5 percent and recorded statements; "changed how users rated a speaker's arguments" dropped | None. This follows the read file. |
| Section 2: Wang et al. (2024) optional, dropped | None. |
| Section 3: Horne and Nevo's below-chance discernment and the "rare failure effect" dropped | Low. The rare failure effect was the design assumption that three flags fall before anyone adapts; it is not needed for the review. |
| Section 3: von Schenk's "the accused could not answer" dropped | Low. "In all of these studies the expert stays silent" covers it. |
| Section 3: Strickland's link to the "No issues found" slides dropped | Low. That belongs in the method. |
| Section 3: no mapping to H1 or its expected size (1.5 to 1.8) | Low for a related-work section. |
| Section 4: Mayer's propensity and "integrity weighs most early" dropped; Kim 2004's "evidence of innocence raised every measure" dropped | Low. |
| Section 4: Ferrin's "no reply moved perceived competence" dropped | Medium-low. It is the reason to analyse the three items separately and to expect the competence item to stay down; it belongs in the method if not here. |
| Section 4: Pezzo reframed as an abstract-level clause; "only study" and "reply paradigm" dropped | None. This is the correction in the read file. The sentence now has no lead-in, though (fix 5). |
| Section 4: Shaw et al. (2003) optional, dropped | None. The read file advised dropping it. |
| Section 4: H4b and Kim 2026's null for it are not named | Low. The Kempt and Nagel paragraph carries the argument; the hypothesis label is missing. |
| Section 5: "Kim 2004 and Pezzo supply the reply paradigm" dropped | None. The read file showed this overstated Pezzo. |
| Section 5: Hancock et al. (2020) optional, dropped | None. |
| Section 5: no closing contribution sentence from Abstract v8 | Medium. See fix 8. |
| Added: Jongsma and Sand (2022) | Positive. It gives section 4 a stated disagreement. |
| Sources: 29 core minus Google, plus Jongsma | Matches the design. |

### Item 4, Author's rules

| Rule | Status | Evidence |
|---|---|---|
| No em-dashes | Met | grep finds no em-dash character in the file. Page ranges use en dashes, and "−1 to +1" uses minus signs. |
| Plain English | Mostly met | "VLM" is never spelled out; it first appears as "a real VLM" in section 2. "CISO" and "METI" are not spelled out either. "H2" and "H3" appear in section 4 without being defined in the review. No statistics jargon. See fix 6. |
| Politician studies not used as evidence, with the one-sentence statement | Met | "Results on well-known politicians, whose reputations are settled, are not carried over." Prike (fictional, unfamiliar candidates) is kept, as the design allows. |
| Trust in the AI not an outcome | Met | It is mentioned only as what Gupta measured. |
| Gupta only in the gap | Met | It is cited once, in section 5. |
| Dobbs not cited | Met | |
| Kim 2026 as closest design and as bounding the AI-penalty claim | Met | Section 3: "their data tie the loss to the error, not to the AI's contradiction". Section 5: "D. Kim et al. (2026) is the closest design". |

### Item 5, Reference list

| Check | Status | Evidence |
|---|---|---|
| Every in-text citation has an entry, and every entry is cited | Met | 34 and 34, matched one by one. |
| Abstract v8 entries (Aamir, Chen, Ferrin, Greitemeyer, Hendriks, Kempt, Kim D., Kim P. H., Madanay, Mayer, Mickelberg, Prike, Reif, Strickland) | Met | Identical to v8 lines 66 to 92. |
| Lowry 72(2), 783–804, 10.1287/mnsc.2023.04147 | Met | PDF header. |
| Opdenbusch, NDSS 2025, 10.14722/ndss.2025.240595 | Met | PDF. |
| Proudfoot 22(4), 235–250, 10.17705/2msqe.00084 | Met | PDF pages 235 and 250. |
| Lu, CHI 2025, pp. 1–19, 10.1145/3706598.3713871 | Met | PDF; README gives 19 pages. |
| Horne and Nevo 15, 13896, 10.1038/s41598-025-98221-7 | Met | PDF footer. |
| Du 17, 1889402, 10.3389/fpsyg.2026.1889402 | Met | PDF. |
| Loconte, five authors, CHB 182, 109019 | Met | Published PDF byline and DOI. |
| Bendeck and Stasko, TVCG 31(1), 1105–1115 | Met | results-5 line 216; DOI on the author PDF. |
| Danry, Kempt, Jongsma, Schilke, Sharma, Pezzo | Met | Read files (Sharma DOI and 49(1), 363-399 also in the web text). |
| von Schenk, iScience 27(7), 110201 | Met | Full-text header. |
| Gupta et al. | Minor issues | DOI is correct (PDF). There is no page range, while Aamir in the same format has one; the PDF says 20 pages. The title is in title case ("Feeling the Facts") where APA wants sentence case. The arXiv byline spells the third author "Ranjendran", not "Rajendran"; check this against the ACM record before submitting. See fix 10. |
| Meta, Apple, Microsoft, Adair, SEC | Met | results-5 entries and SEC PDF. |

### Item 6, Readability for a CHI reviewer

- **One argument?** Mostly. The section titles read as a chain: boards cannot judge, checkers exist and are wrong, wrong flags are believed and cost the person, repair is possible but studied only with human accusers, and nobody has combined these. The gap section closes the chain well.
- **Where it still lists papers:** the three runs named under "Synthesis" above. The worst is section 4, paragraph 2, where five sources each get one sentence and only the last two sentences ("Hence H2 ...") say what they add up to. Moving that "Hence" sentence to the front of the paragraph would make it read as an argument at no word cost.
- **Where it is vague or dangles:**
  - The Pezzo sentence ends section 4 with no link to what comes before it. A reviewer will not see why an unconsulted diagnostic aid matters (fix 5).
  - "Their power was having alternatives, not authority" is true but leaves the reader to draw the consequence (so the result may not carry over to a board's authority). It is acceptable as it stands.
  - "Hence H2 ... and H3 ..." uses hypothesis labels the review never defines.
  - "Danry et al. rated recorded statements" makes the authors the raters; the participants were.
  - The final sentence lists design features but not what the study will tell us (fix 8).

## D. Ranked fix list

The word changes are for the body. Fixes 1 to 4 and 6 add 15 words (1,487). Fix 5 adds 13 more (1,500, right at the limit), so make it together with its Apple offset (−17), which brings the body to 1,483. Fixes 7 and 8 need further cuts, given with them.

**1. Du et al.: overstated "drop larger" on the 10-point scale (+2 words).**
Old: "A fictitious report of a high AI-detection rate made 214 teachers grade an identical paper 1.8 points lower on a 10-point scale, and red highlighting made the report seem more credible and the drop larger (Du et al., 2026)."
New: "In a study of 214 teachers, a fictitious high AI-detection rate cut an identical paper's grade by 1.8 points on a 10-point scale; red highlighting made the report seem more credible and cut the 100-point score further (Du et al., 2026)."

**2. Aamir et al.: self-reported confidence, not observed pushing back (+1).**
Old: "Directors pushed back more confidently where independent reviews were routine (Aamir et al., 2026)."
New: "Directors felt more confident challenging management where independent reviews were routine (Aamir et al., 2026)."

**3. Sharma et al.: "most studies fall short" is a cited claim, not their own finding (+3).**
Old: "Most studies they review find that repair falls short of the start, and denial works less well after competence than after integrity violations."
New: "Citing Kim (2018), they note that most studies find repair falls short of the start, and denial works less well after competence than after integrity violations."
(Kim, 2018, is a secondary mention through Sharma; it needs no reference-list entry if worded as "citing", but if the course wants it in the list, use "as cited in" style: "most studies find repair falls short of the start (Kim, 2018, as cited in Sharma et al., 2023)".)

**4. Bendeck and Stasko: 13 percent is one question category (−3).**
Old: "yet answered only 13 percent of questions that needed a value read off a chart when it lacked the data (Bendeck & Stasko, 2025)."
New: "yet answered only 13 percent of one type of value-lookup question when it lacked the underlying data (Bendeck & Stasko, 2025)."

**5. Pezzo et al.: give the sentence a lead-in so it joins the argument (+13; pay with the Apple cut below).**
Old: "A physician who had not consulted an available diagnostic aid before a bad outcome was rated higher"
New: "The nearest lay evidence on a stated reason about a machine is indirect. A physician who had not consulted an available diagnostic aid before a bad outcome was rated higher"
Offset (−17, and the Apple entry leaves the reference list, giving 33 sources):
Old: "Ray-Ban Meta glasses answer questions about what the wearer is looking at (Meta, 2024), Apple's Visual Intelligence lets users ask for \"an explanation of a diagram from class notes\" (Apple, 2024), and Copilot"
New: "Ray-Ban Meta glasses answer questions about what the wearer is looking at (Meta, 2024), and Copilot"
(The writer's notes already list the Apple sentence as the easiest cut.)

**6. Spell out the three acronyms at first use (+12).**
- Old: "points a camera assistant at the CISO's slide" / New: "points a camera assistant at the slide of the chief information security officer (CISO)" (+6)
- Old: "Our flags are drafted with a real VLM and are all wrong for the company." / New: "Our flags are drafted with a real vision-language model (VLM) and are all wrong for the company." (+2)
- Old: "Our three items come from the METI, built on the same factors" / New: "Our three items come from the Muenster Epistemic Trustworthiness Inventory (METI), built on the same factors" (+4)

**7. Section 3: add the "how we differ" line (+10; optional unless a cut is found).**
Old: "In all of these studies the expert stays silent and trust is rated once."
New: "In all of these studies the expert stays silent and trust is rated once; our CISO answers, and trust is rated after each step."
Offset if needed: cut "None is set in a boardroom." from section 5 (−5), since section 1 already says the board evidence has no tool or trust measure, and the scenario makes the boardroom obvious.

**8. End on the contribution, as `_review-design-v2.md` asks (+25; needs a cut).**
Add after the final sentence: "The study thus asks whether a wrong flag from a listener-held checker costs a correct expert trust, and whether the expert's reason wins it back."
Offset options: the Mickelberg sentence (−22, but it carries a design choice) or the Jongsma and Sand clause (−9, but it gives section 4 its stated disagreement). My choice would be to shorten the gap sentence "We bring the pieces together: the checker is in the hands of the non-expert who holds power over the expert" to "Our study combines them: the checker is held by the non-expert with power over the expert" (−5), and to cut "None is set in a boardroom." (−5), which still leaves 15 words to find.

**9. Copilot wording (+3; cosmetic).**
Old: "Copilot in Microsoft Teams suggests the private prompt"
New: "Copilot in Microsoft Teams suggests, in a private pane, the prompt"

**10. Gupta et al. reference (list only).**
Old: "Gupta, C., Aritonang, N. V., Rajendran, D. P. D., Danry, V., Maes, P., & Nanayakkara, S. (2026). Feeling the Facts: Real-time wearable fact-checkers can use nudges to reduce user belief in false information. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems*. Association for Computing Machinery."
New: "Gupta, C., Aritonang, N. V., Rajendran, D. P. D., Danry, V., Maes, P., & Nanayakkara, S. (2026). Feeling the facts: Real-time wearable fact-checkers can use nudges to reduce user belief in false information. In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (pp. 1–20). Association for Computing Machinery."
Also check the spelling of the third author against the ACM record; the arXiv byline reads "Ranjendran".

**11. Kim et al. (2026) "monitored" (0 words; optional).**
Old: "an AI monitored a human advisor in front of the student who depended on him"
New: "an AI checked a human advisor's work in front of the student who depended on him"
(This is word-neutral apart from +1. "Monitored" fits only the reactive conditions; in the proactive ones the advisor invoked the AI.)

**12. Danry: name the raters (0 words; optional).**
Old: "Danry et al. rated recorded statements, not a person."
New: "Danry et al.'s participants rated recorded statements, not a person." (+1)

## E. Evidence log

Full texts read (pdftotext -layout, into the scratchpad):
- SEC (2023) final rule, pp. 82 to 85 (Item 407(j) decision and the "broad-based skills" passage).
- Lowry et al. (2026): abstract p. 783; section 4 (D-9 quote).
- Proudfoot et al. (2023): pp. 235, 240 (the "heavily determined" passage) and 250.
- Opdenbusch et al. (2025): abstract, contribution list p. 2, results IV-A and IV-B pp. 5 to 6, discussion pp. 10 to 13.
- Aamir et al. (2026): pp. 7, 13 and 14 (external assurance and confidence passages).
- Schilke et al. (2015): pp. 12950 to 12951 (theory and Study 1 statistics).
- Danry et al. (2020): sections 4.3 and 5.1.
- Bendeck and Stasko (2025): sections 4, 5.2, 5.3 and 6 (16th percentile, 13%, best practice, weaknesses).
- Horne and Nevo (2025): Study 1 design and Table 2; label wording.
- Lu et al. (2025): abstract, introduction, section 5 and the conclusion (the quote appears three times).
- Du et al. (2026): abstract, sections 3.4, 4.1 to 4.3, Tables 1 and 2, section 5 (interaction caveat).
- von Schenk et al. (2024): results (84.91%) and discussion ("almost 85%").
- Strickland (1958): pp. 206 and 215.
- Prike et al. (2025) preprint: abstract and results pp. 17 to 19.
- Madanay et al. (2025) txt: methods, cell sizes, Table 2.
- Chen et al. (2026): abstract, sections 4.2, 4.5 and 5.1.
- Reif et al. (2025) markdown: Significance, Abstract, Study 2.
- Kim et al. (2026): pp. 1515 to 1517 (design, N, METI, post-task measure).
- Loconte et al. (2026), CHB: abstract.
- Kim et al. (2004): pp. 107 to 110 (hearsay accusation).
- Ferrin et al. (2007): abstract.
- Greitemeyer and Sagioglou (2015): abstract.
- Mickelberg et al. (2025): pp. 2 to 3.
- Sharma et al. (2023) web text: introduction, definition, denial, measurement and conclusion sections.
- Kempt and Nagel (2022): pp. 226 to 227.
- Gupta et al. (2026) arXiv: byline, abstract, measures table, section 6.3.

Record files read: `_check-brief.md`, `Instructions.md`, `Literature Review - Draft v1.md`, `_review-design-v2.md`, `fact-check-2026-10-19.md`, `verified-view.md` (entries 2, 5, 8, 10, 15 to 18, and rule 22), `read-pezzo-schilke-sharma.md` (Pezzo, Schilke, Sharma), `read-danry-kempt-dasilva.md` (Kempt, Jongsma), `results-5-board-products.md` (products, Squash, Bendeck), `results-1-hci-checkers.md` (Gupta line 214), `CMPT 863 Project Abstract v8.md` (hypotheses, related-work appendix, references), `references/README.md`.

Not read: check-A.md and check-B.md. The Pezzo full text is not in the project. The Jongsma and Sand text was taken from the read file, not read directly.
