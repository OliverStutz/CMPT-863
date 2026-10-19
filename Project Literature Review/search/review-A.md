# Review A: selection for the literature review

Reviewer A, written 2026-10-18. I read the brief, the assignment, the abstract v7, the agent briefs, the known-papers list, all six result files and the earlier 75-entry list in full. No web search. Where I add knowledge from outside the files, I say so.

## 1. Proposed thematic subsections

**1.1 Boards depend on a CISO they cannot check (context).** Directors without cyber expertise oversee the CISO symbolically, rarely ask questions for fear of looking like novices, and judge the CISO's competence by how well the CISO explains; regulation assumes this will stay so, and camera and meeting assistants now offer a check on the spot (Aamir et al., 2026; Lowry et al., 2026; Opdenbusch et al., 2025; Da Silva & Jensen, 2022; Proudfoot et al., 2023; SEC, 2023; Microsoft Copilot and camera assistants).

**1.2 AI checkers in HCI judge content, are followed when wrong, and can be wrong about slides.** HCI checkers and credibility labels move belief in the checked content toward the AI's verdict even when the verdict is wrong, and VLMs misread charts and diagrams while knowing textbook rules; but these studies measure belief in content or trust in the AI, and AI-mediated communication research places the AI on the speaker's side, not in the listener's hand (Gupta et al., 2026; Lu et al., 2025; Horne & Nevo, 2025; Hancock et al., 2020; Bendeck & Stasko, 2025).

**1.3 Checking, flagging and AI disagreement lower the standing of the person checked, though not always.** Checking alone lowers trust in the person checked (Strickland, 1958); AI disagreement or an AI-detected error lowers lay ratings of an expert (Kim et al., 2026; Chen et al., 2026; Madanay et al., 2025), a correction costs a trusted professional more than a politician (Dobbs et al., 2026; Prike et al., 2025), a wrong automated flag lowers an evaluator's rating of identical work (Du et al., 2026), and holding an algorithm makes people accuse (von Schenk et al., 2024); against this, AI and human fact-checks left politicians' perceived competence unchanged (Chae & Tewksbury, 2026), and AI use itself carries a penalty the design has to avoid (Reif et al., 2025).

**1.4 Trust in an expert: what it is, how a layperson rates it, and how it is repaired.** Trust rests on perceived ability, benevolence and integrity, which laypeople distinguish when judging experts (Mayer et al., 1995; Hendriks et al., 2015); after a suspicion, denial and reticence work differently for competence and integrity (Kim et al., 2004; Ferrin et al., 2007), a third-party clearing restores credibility to about the level of the never-accused but not reliably above it (Greitemeyer & Sagioglou, 2015; Mickelberg et al., 2025), corrected negative information can leave a residue (Thorson, 2016), and an unsupported denial can work better than silence or apology (Schiff et al., 2025).

**1.5 Gap: no study gives the non-expert with power a wrong AI checker and lets the expert answer.** The closest designs either make the AI right, keep the expert silent, put the AI in the expert's hands, or measure belief instead of trust in the person (Kim et al., 2026; Chen et al., 2026; Madanay et al., 2025; Gupta et al., 2026; Du et al., 2026). This short closing subsection carries the differentiation; it reuses papers from 1.1 to 1.4 and needs no new ones.

## 2. Selected papers

30 papers (one row groups three closely related products). 14 must cite, 13 should cite, 3 optional.

### 1.1 Boards and CISOs

- **Aamir et al., 2026** (known, in abstract). Anchor. Must cite. The only CHI paper on board cyber governance in the files, and the source of "directors pushed back more confidently where independent reviews were routine", which motivates both the checker and H4.
- **Lowry et al., 2026.** Anchor, context. Must cite. Field evidence that non-expert directors' oversight is "mostly symbolic" and that they do not notice it; the strongest scholarly basis for "boards cannot judge the CISO's claims alone".
- **Opdenbusch et al., 2025.** Context, H4. Must cite. Board members do not ask the CISO questions because they "fear being exposed as IT novices", and reduce cyber risk to budget decisions; this is the silent baseline that H4a expects a checker to change, and it matches our budget scenario.
- **Da Silva & Jensen, 2022.** Context, SIGCHI anchor. Should cite. CSCW paper naming the CISO as the interpreter of an opaque field, in a role that is "superior and precarious"; it describes the interpretive monopoly that a checker breaks.
- **Proudfoot et al., 2023.** Context, supports H2 and H3. Should cite. Interviewees say the board's perceived competence of the CISO is "heavily determined by the CISO's ability to effectively communicate", which is the practitioner basis for expecting the reasoned answer to matter more than the technical facts. Preferred over Hooper & McKissack (2016), which makes the same point with a thinner evidence base.
- **SEC, 2023.** Industry or trade (regulation). Should cite. The SEC dropped the board-expertise disclosure because directors "often effectively oversee management's efforts without specific subject matter expertise"; one sentence that fixes the non-expert overseer as the regulatory norm. Preferred over the NACD handbook (not read in full) and the Proofpoint and Gartner surveys (vendor self-reports).
- **Products: Microsoft Copilot in Teams; camera assistants (Ray-Ban Meta, Apple Visual Intelligence, Gemini Live).** Product. Should cite. Copilot suggests the private prompt "Where are some holes in [a meeting participant's] argument?", and the camera assistants already let a listener point at a diagram and ask about it; together they show our checker is a realistic scenario, not a speculative one, and that no vendor reports what it does to trust in the speaker.

### 1.2 AI checkers and AI-mediated communication

- **Gupta et al., 2026** (known). Anchor, gap. Must cite. CHI wearable checker of spoken claims with imperfect accuracy: the nearest listener-held, real-time checker, and it measured belief in the claims (and trust in the AI's information), not trust in the speaker.
- **Lu et al., 2025** (known). Supporting finding (H1). Should cite. CHI paper in which an AI credibility indicator swayed people more when it disagreed with a human expert than when it agreed; the closest SIGCHI evidence that AI-versus-expert conflict is where the AI dominates. Cited instead of its CSCW 2022 predecessor, which makes the same point without the expert.
- **Horne & Nevo, 2025.** Supporting finding (H1). Should cite. People relied on warning labels even when the labels were wrong, explained as a learned shortcut from mostly correct human labels; the direct evidence for our manipulation (every flag wrong). Preferred over Freeze et al. (2021) and He et al. (2026, preprint), which make the same point for news, the latter not peer reviewed.
- **Hancock et al., 2020.** Definition. Optional. Defines AI-mediated communication as an agent acting for a communicator; one sentence lets us say our checker is receiver-side mediation, outside the definition. Preferred over Jakesch et al. (2019) and Navajas Fernández et al. (2026) because the review needs the definition more than a third sender-side finding.
- **Bendeck & Stasko, 2025.** Design element (premise of wrong flags). Should cite. GPT-4V scored at about the 16th percentile of humans on a visualization literacy test, read values from charts badly without data, yet knew design rules: exactly the "right by the textbook, wrong for this company" profile, measured on charts like ours. Preferred over CharXiv and HallusionBench (machine-learning benchmarks, less legible to an HCI marker).

### 1.3 Flags and the standing of the person flagged

- **Strickland, 1958** (known, in abstract). Anchor (founding mechanism). Must cite. Checking alone lowered trust in the checked worker when every check found nothing; it is the reason for the "No issues found" slides.
- **Kim et al., 2026** (Kim, Daneshi, Draper & Sreedharan, AIES). Anchor, closest design, measure. Must cite. A human advisor's errors lowered trust on METI expertise, integrity and benevolence, and an AI that automatically monitored and flagged the errors did not change the penalty; this is the nearest design in all files and supplies the measure we will use.
- **Chen et al., 2026** (known, in abstract). Anchor, closest design. Must cite. AI disagreement with a doctor lowered trust in the doctor while the doctor stayed silent; the design that our answer step extends.
- **Madanay et al., 2025.** Supporting finding (H1), counter-detail. Must cite. 1,606 adults rated a radiologist lower when the radiologist recommended less testing than the AI, but not when the radiologist recommended more; the penalty depends on the AI looking more cautious, which our "numbers do not add up" flags do. Chosen over Bernstein et al. (2025), whose numbers come from press coverage.
- **Dobbs et al., 2026.** Supporting finding (H1), measure. Should cite. Corrections lowered rated credibility (trustworthiness and expertise) of both a doctor and a senator, and the doctor lost more; the only file entry showing a correction costs a trusted expert more, and the loss was largest when the error was framed as intentional (integrity).
- **Prike et al., 2025** (known, in abstract). Supporting finding (H1), design element. Should cite. Fact-checks lowered rated competence and morality of an unfamiliar candidate and his vote share; the source of our profile and 0 to 10 rating, and the evidence that corrections transfer to an unfamiliar person.
- **Du et al., 2026.** Supporting finding (H1), power setting. Should cite. 214 teachers given a fictitious 87% AI-detection report graded an identical paper lower; the evaluator with power held a wrong automated flag, our arrangement, though grades rather than trust were measured.
- **von Schenk et al., 2024** (known, in abstract). Supporting finding (H4). Must cite. People who chose to consult a two-in-three accurate lie detector accused in most flagged cases, against about one in five unaided; the evidence that holding a tool changes what one does to another person.
- **Chae & Tewksbury, 2026** (known). Counter-result (H1). Must cite. Across three experiments (N = 3,978 per Agent 6's note) neither AI nor human fact-checks changed feelings toward or perceived competence of the targeted politicians; the clearest file evidence that a check can leave the person's standing untouched.
- **Reif et al., 2025** (known, in abstract). Design element. Optional. Employees getting help from AI were rated lazier and less competent; it justifies the CISO never mentioning AI, but that is one design sentence. I keep it rather than argue it out, at optional priority; Schilke & Reimann (2025) would be the stronger alternative if the review wants a disclosure-penalty anchor.

### 1.4 Trust definition, measure and repair

- **Mayer et al., 1995.** Definition (ability, benevolence, integrity). Must cite. The founding definition; its phrase "irrespective of the ability to monitor or control" frames why a checker matters at all.
- **Hendriks et al., 2015 (METI).** Measure. Must cite. Laypeople's trust in experts forms three factors (expertise, integrity, benevolence) that respond to manipulated source information; it is the scale Kim et al. (2026) used and the one our three-part trust question should follow.
- **Kim et al., 2004** (Kim, Ferrin, Cooper & Dirks). Anchor (repair), counter-prediction for H3. Must cite. Denial repaired trust better than apology when the accused was later shown innocent; our CISO is innocent by design, so a bare denial may not fail, which H3 must face.
- **Ferrin et al., 2007.** Supporting finding (H3). Should cite. Reticence did worse than the better of apology and denial, and for competence violations it failed to signal improvement; our "denial with no reason" shares this weakness, so this is the better support for H3.
- **Greitemeyer & Sagioglou, 2015** (known, in abstract). Anchor (H2). Must cite. Clearing an accused researcher restored credibility to about the never-accused level, not reliably above; it sets the expected ceiling of H2 and motivates the neutral group hearing the same answers.
- **Mickelberg et al., 2025** (known, in abstract). Supporting finding (H2), measurement caution. Optional (kept from the abstract; Greitemeyer & Sagioglou carries the main H2 point). Retraction removed the allegation's effect on likeability in two of three experiments, and repeated ratings may draw attention to the latest information; both points shape our design.
- **Thorson, 2016.** Counter-result (H2). Should cite. Negative information about a candidate kept lowering attitudes after an immediate, accepted correction ("belief echoes"); the main reason to expect the "above the start" part of H2 to fail. Preferred over Ecker et al. (2022), a review making the general point without a person as target.
- **Schiff et al., 2025.** Counter-result (H3). Should cite. Politicians who falsely claimed a scandal story was fake gained more support than those who stayed silent or apologised (over 15,000 adults, five experiments); an unsupported denial can help the accused, which cuts against H3 directly.

### Considered and left out (for the cross-examination)

Loconte et al. (2026, CHB: observers reject confident accusatory AI verdicts) is the strongest omitted counter-result and my first replacement if a row is dropped. Pezzo et al. (2022) shows that the stated reason for setting an aid aside changes ratings of the physician, a close H2 and H3 analogue; second replacement. Garrett & Weeks (2013, CSCW) is the earliest HCI real-time correction study and could replace Lu et al. (2025) if the review wants an older SIGCHI source. Sharma et al. (2023) defines partial versus complete repair and would sharpen H2 if the word budget allows. Kim et al. (2006) predicts external accounts help less after competence violations, a counterweight to H2. Hendriks et al. (2016) is a close lay-expert analogue but its findings come from secondary sources only. The pre-AI monitoring economics (Falk & Kosfeld, 2006; Schweitzer et al., 2018) and the visible-error papers (Appelman, Kreiner, Vignovic) serve the earlier presenter-side Idea 2, not this study.

## 3. Must-read list (full text before writing)

1. **Chen et al., 2026.** The abstract v7 says disagreement "made the doctor seem lazier". Agent 4's note (from press coverage) says *agreement* made the doctor look lazier and disagreement raised perceived uncertainty and lowered perceived autonomy; the earlier list says disagreement raised uncertainty and laziness. Check which condition produced laziness, the exact trust measure, N = 135, and whether the "AI shifts the burden of proof" sentence attributed to Chen et al. in the abstract is theirs or Kempt & Nagel's (2022).
2. **Kim et al., 2026 (AIES).** The closest design. Check: whether the AI-flagged errors were real errors (so the study shows the penalty for being right-flagged, not wrong-flagged), the per-cell sample (157 over six cells), how the "AI monitored and flagged" condition was presented, and whether any condition let the advisor answer. Its null on AI involvement changes how H1 should be argued (see section 5).
3. **Kim et al., 2004.** Agent 3 read only the abstract. Check whether the denial carried any account or was a bare claim of innocence, whether "evidence of innocence" came from a third party, and whether competence-type violations with innocence evidence favoured denial in both studies. This decides whether it contradicts H3.
4. **Strickland, 1958.** The earlier list says findings were taken from secondary sources because the original is paywalled; the abstract v7 gives specific numbers (40 male students, checks after nine of ten versus two periods). Confirm these against the original.
5. **Chae & Tewksbury, 2026.** Agent 6 read the abstract via a UCLA page, not the journal. Check how competence was measured, that the targets were well-known politicians (which would make the null less relevant to an unfamiliar CISO), and the N of 3,978.
6. **Gupta et al., 2026.** Agent 1's reading that trust was measured in the AI's information, not the speaker, comes from a search snippet, "not re-read in full". The gap statement leans on it.
7. **Schiff et al., 2025.** Check how large the "claiming misinformation" advantage over silence and apology was, and that it held for text but not video; our slides are closer to visual evidence than to a text report, which may shrink the counter-result.
8. **Mickelberg et al., 2025** and **Greitemeyer & Sagioglou, 2015.** The abstract v7 gives specific numbers (two of three experiments; about two points on a scale from minus 5 to plus 5; N = 185 and 358) and attributes a measurement caution to Mickelberg et al.; confirm the numbers and the wording of the caution.
9. **Ferrin et al., 2007.** Abstract only; check how reticence was worded, to judge how close it is to a denial with no reason.
10. **Microsoft Copilot support page.** The quoted prompt "Where are some holes in [a meeting participant's] argument?" should be re-checked close to submission, since product pages change.
11. **Lu et al., 2025.** The finding ("larger when it disagreed with a human expert") comes from the earlier list's one-line summary; read the result section before using it as the SIGCHI anchor for AI-versus-expert conflict.

## 4. Gap and differentiation

The files support a narrow but clear gap: no study gives a non-expert who holds power over an expert an AI checker whose flags are wrong, lets the expert answer each flag, and rates trust in the expert before and after each step. The closest design is Kim et al. (2026), where an AI that monitored and flagged a human advisor's errors left trust on METI dimensions as low as when there was no AI; but the errors were real, the advisor did not answer, and the participant was a dependent student, not the party with the budget. Chen et al. (2026) has an AI contradict a doctor, but the doctor brought in the AI, stayed silent, and trust was rated once. Madanay et al. (2025) shows an asymmetric penalty for an expert who disagrees with an AI, in a static vignette without an answer or a baseline. Gupta et al. (2026) puts a live checker in the listener's hand but measures belief in claims, and Du et al. (2026) gives an evaluator with power a wrong automated flag but measures grades, not trust, and the student cannot respond. No paper closes the gap. Kim et al. (2026) does force a narrower claim: the review cannot say no one has measured trust in an expert after an AI flag; it should say the novelty is the combination of a wrong flag, a checker held by the powerful non-expert, and the expert's answer as the repair step. Kim et al. (2026) also suggests the trust loss may follow the apparent error rather than the AI as such, which our neutral group (same answers, no flag) is placed to test.

## 5. Hypothesis check

**H1 (a flag lowers trust).** Supported by Kim et al. (2026), Chen et al. (2026), Madanay et al. (2025), Dobbs et al. (2026), Prike et al. (2025), Du et al. (2026), Horne & Nevo (2025) and Strickland (1958). Cut against by Chae & Tewksbury (2026, no change in politicians' perceived competence) and, outside the selection, Loconte et al. (2026, observers resist confident AI accusations). No rewording needed, but the review should note that the flags are hedged ("likely") and that three of five slides are flagged, both of which the files tie to the size of the drop.

**H2 (a reasoned answer restores trust, possibly above the start).** Recovery is supported by Greitemeyer & Sagioglou (2015), Mickelberg et al. (2025) and Proudfoot et al. (2023); going above the start is cut against by Thorson (2016) and by Greitemeyer & Sagioglou's own result (back to level, not above). The files suggest rewording: "after a reasoned answer, trust rises from its post-flag level toward the starting level", with "above the start" stated as an open question rather than the prediction. The current wording in v7 ("higher than prior to flagging and increases beyond initial trust") predicts more than any file supports.

**H3 (a denial without a reason restores less or lowers further).** Supported by Ferrin et al. (2007) and indirectly by Proudfoot et al. (2023). Cut against by Kim et al. (2004, denial works when the accused is innocent, as our CISO is) and Schiff et al. (2025, an unsupported denial beat silence and apology). The files favour the weaker form: "a denial without a reason restores less trust than a reasoned answer"; the v7 form ("lower than after the flag") should be kept only as a secondary, exploratory prediction.

**H4 (checker participants question more, and answers move trust further).** H4a is supported by Opdenbusch et al. (2025), Aamir et al. (2026) and von Schenk et al. (2024); H4b (amplification) has no direct support in the selection, and Chae & Tewksbury (2026) suggests AI labels need not add weight. The files suggest rewording H4a so both groups report the same thing (questions wanted per slide), since v7 compares actions taken with questions stated, and presenting H4b as exploratory.

## Selection table

| Short citation | Subsection | Role | Priority | New or known |
|---|---|---|---|---|
| Aamir et al., 2026 | 1.1 Boards and CISOs | Anchor, context (H4) | Must cite | Known (abstract) |
| Lowry et al., 2026 | 1.1 Boards and CISOs | Anchor, context | Must cite | New |
| Opdenbusch et al., 2025 | 1.1 Boards and CISOs | Context, H4 | Must cite | New |
| Da Silva & Jensen, 2022 | 1.1 Boards and CISOs | Context (SIGCHI) | Should cite | New |
| Proudfoot et al., 2023 | 1.1 Boards and CISOs | Context, supports H2 and H3 | Should cite | New |
| SEC, 2023 | 1.1 Boards and CISOs | Industry or trade (regulation) | Should cite | New |
| Microsoft Copilot in Teams; Ray-Ban Meta, Apple Visual Intelligence, Gemini Live | 1.1 Boards and CISOs | Product | Should cite | New |
| Gupta et al., 2026 | 1.2 AI checkers | Anchor, gap | Must cite | Known |
| Lu et al., 2025 | 1.2 AI checkers | Supporting finding (H1) | Should cite | Known |
| Horne & Nevo, 2025 | 1.2 AI checkers | Supporting finding (H1, wrong flags) | Should cite | New |
| Hancock et al., 2020 | 1.2 AI checkers | Definition (AI-MC) | Optional | New |
| Bendeck & Stasko, 2025 | 1.2 AI checkers | Design element (VLM reliability) | Should cite | New |
| Strickland, 1958 | 1.3 Flags and the person | Anchor (mechanism) | Must cite | Known (abstract) |
| Kim, Daneshi, Draper & Sreedharan, 2026 | 1.3 Flags and the person | Anchor, closest design, measure | Must cite | New |
| Chen et al., 2026 | 1.3 Flags and the person | Anchor, closest design | Must cite | Known (abstract) |
| Madanay et al., 2025 | 1.3 Flags and the person | Supporting finding (H1) | Must cite | New |
| Dobbs et al., 2026 | 1.3 Flags and the person | Supporting finding (H1), measure | Should cite | New |
| Prike et al., 2025 | 1.3 Flags and the person | Supporting finding (H1), design element | Should cite | Known (abstract) |
| Du et al., 2026 | 1.3 Flags and the person | Supporting finding (H1, power) | Should cite | New |
| von Schenk et al., 2024 | 1.3 Flags and the person | Supporting finding (H4) | Must cite | Known (abstract) |
| Chae & Tewksbury, 2026 | 1.3 Flags and the person | Counter-result (H1) | Must cite | Known |
| Reif et al., 2025 | 1.3 Flags and the person | Design element | Optional | Known (abstract) |
| Mayer et al., 1995 | 1.4 Trust and repair | Definition | Must cite | New |
| Hendriks et al., 2015 | 1.4 Trust and repair | Measure (METI) | Must cite | New |
| Kim, Ferrin, Cooper & Dirks, 2004 | 1.4 Trust and repair | Anchor, counter-prediction (H3) | Must cite | Known (not yet read) |
| Ferrin et al., 2007 | 1.4 Trust and repair | Supporting finding (H3) | Should cite | Known (not yet read) |
| Greitemeyer & Sagioglou, 2015 | 1.4 Trust and repair | Anchor (H2) | Must cite | Known (abstract) |
| Mickelberg et al., 2025 | 1.4 Trust and repair | Supporting finding (H2), measurement caution | Optional | Known (abstract) |
| Thorson, 2016 | 1.4 Trust and repair | Counter-result (H2) | Should cite | New |
| Schiff et al., 2025 | 1.4 Trust and repair | Counter-result (H3) | Should cite | New |
