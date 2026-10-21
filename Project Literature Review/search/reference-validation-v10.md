# Reference validation: Literature Review Draft v10 (2026-10-20)

Brief: `search/_v10-reference-validation-brief.md`. Files read: `Literature Review - Draft v10.md`, `Literature Review - Draft v9.md` (to separate section 1 from the rebuilt text), `Literature Review - Draft v6.md` (body and reference list), `search/reference-audit-v7.md` (the three fixes), `search/academic-tally.md`, `search/verified-view.md`, `search/read-pezzo-schilke-sharma.md`, `../references/reviews/chen.md`. Full texts read where v10 differs from v6 in substance: D. Kim et al. 2026, Kempt and Nagel 2022, Ferrin et al. 2007, Sharma et al. 2023 (.txt), Madanay et al. 2025 (.txt), plus number spot-checks in Danry et al. 2020 and Bendeck and Stasko 2025 (all PDFs through `pdftotext -layout`). Crossref API for six DOIs; the Meta and Microsoft pages fetched with curl today.

v10 body text is v9's sentences in a new order (checked: every body sentence of v10 appears verbatim in v9). So every fidelity finding below is also present in v9; the placement findings are new to v10.

## 1. Summary

- **Citations:** 30 distinct sources cited (44 in-text occurrences); 30 matched both ways; 0 citations without an entry; 0 entries without a citation; 0 year or spelling mismatches; 0 order errors inside the three multi-source citations. 1 minor form problem: "Danry et al.'s participants" (line 43) has no year in a new paragraph.
- **Rebuilt sentences:** 43 checked. Same claim 31; nuance lost 9; meaning changed 3.
- **Placement problems:** 7 (two of them are also the "meaning changed" items), plus one optional note.
- **Entries:** 30. Correct 30; fix needed 0. The three fixes of today are present and correct. Six Crossref spot-checks match. Meta URL resolves (200, matching title); Microsoft URL resolves (200, matching title, prompt present).
- **Rules:**
  - No em-dashes: pass (0 in the whole file).
  - Gupta et al. (2026) only in the closest-designs material: pass (body line 61 only).
  - Dobbs absent: pass (body); the name appears only in the working notes ("Dobbs absent").
  - Politician sentence once: pass (line 29).
  - Trust in the AI not an outcome: pass (the only mentions describe other studies, lines 39 and 61; the study's outcome is trust in the CISO, line 65).
  - No H1 to H4 labels: pass in the body; the working notes (line 150) still say "(H1 and H2 in plain words)".
  - Abbreviations spelled out at first use in v10's reading order: pass. CISO line 5; SEC line 11 ("U.S. Securities and Exchange Commission (SEC, 2023)"); METI line 15; VLM line 67 (VLM no longer appears earlier, since the opening dropped it). Not on the list but noted: "HCI" (line 39) is never spelled out.

## 2. Citation table

| Source (in-text form) | Body lines in v10 | Entry present | Year match | Notes |
|---|---|---|---|---|
| Aamir et al., 2026 | 21 | yes | yes | |
| Adair, 2020 | 25 | yes | yes | |
| Bendeck & Stasko, 2025 | 25 | yes | yes | |
| Chen et al., 2026 | 29, 33, 39, 61 | yes | yes | |
| Danry et al., 2020 | 25, 39, 43 | yes | yes | line 43 "Danry et al.'s participants" has no year; it opens a new section, so APA wants "(2020)" |
| Du et al., 2026 | 29 | yes | yes | |
| Ferrin et al., 2007 | 33, 39 | yes | yes | |
| Greitemeyer & Sagioglou, 2015 | 33 | yes | yes | |
| Gupta et al., 2026 | 61 | yes | yes | closest designs only |
| Hendriks et al., 2015 | 15 | yes | yes | |
| Horne & Nevo, 2025 | 29 | yes | yes | |
| Kempt and Nagel, 2022 | 33 | yes | yes | |
| D. Kim et al., 2026 | 29, 39, 61 | yes | yes | initial used to separate from P. H. Kim; correct APA practice |
| P. H. Kim et al., 2004 | 33, 39 | yes | yes | |
| Loconte et al., 2026 | 29 | yes | yes | |
| Lowry et al., 2026 | 11 (twice), 21 (twice) | yes | yes | second line-21 mention ("Lowry et al. read") omits the year inside the same paragraph; allowed |
| Lu et al., 2025 | 29, 39 | yes | yes | |
| Madanay et al., 2025 | 29 | yes | yes | |
| Mayer et al., 1995 | 15 | yes | yes | |
| Meta, 2024 | 11 | yes | yes | |
| Microsoft, n.d. | 11 | yes | n.d. both | |
| Opdenbusch et al., 2025 | 21 | yes | yes | |
| Pezzo et al., 2022 | 33 | yes | yes | |
| Prike et al., 2025 | 29 | yes | yes | |
| Reif et al., 2025 | 29 | yes | yes | |
| Schilke et al., 2015 | 39 | yes | yes | |
| SEC, 2023 | 11 | yes ("U.S. Securities and Exchange Commission") | yes | group author with abbreviation at first citation; correct |
| Sharma et al., 2023 | 33, 61 | yes | yes | |
| Strickland, 1958 | 29 | yes | yes | |
| von Schenk et al., 2024 | 29 | yes | yes | |

Multi-source citations: (Ferrin et al., 2007; P. H. Kim et al., 2004), (Danry et al., 2020; Lu et al., 2025), (Chen et al., 2026; D. Kim et al., 2026). All alphabetical.

Reference-list order: alphabetical throughout (U.S. Securities and Exchange Commission before von Schenk; D. Kim before P. H. Kim).

## 3. Fidelity table (rebuilt sentences)

Line numbers are v10's. "v6" quotes `Literature Review - Draft v6.md`.

| # | v10 sentence (start) | v6 counterpart | Verdict | Evidence and exact fix |
|---|---|---|---|---|
| R1 | 5 "A board member points a camera-based checker at a slide; the checker flags a fault that is not there; ..." | opening, "flags a purported fault, and the CISO responds" | Same claim | "not there" matches the design (all flags wrong). VLM moved to line 67. |
| R2 | 5 "Five bodies of work bear on this moment ..." | "Four bodies of work ..." | Same claim | Fifth body (closest designs) added; structural. |
| R3 | 15 "Trust is the willingness to be vulnerable ... (Mayer et al., 1995)." | "Following Mayer et al. (1995), trust is ..." | Same claim | |
| R4 | 15 "Our three trust items follow the ... (METI) ..." | "Our trust items are drawn from ... (METI) ..." | Same claim | "three" matches Abstract v8 (three items: competence, benevolence, integrity). |
| R5 | 25 "Checking tools held by the listener exist as products (above) and as research systems, and their accuracy is limited." | "Evaluations of related checkers, however, report limited accuracy." | Nuance lost | v10 extends "limited accuracy" to the Meta and Microsoft products, which no cited study evaluated. Fix: replace "and their accuracy is limited." with "and evaluations of related checkers report limited accuracy." |
| R6 | 25 Danry 63.5% | same | Same claim | Danry full text: "of 63.5%" (Results), "pre-recorded arguments" (method). |
| R7 | 25 Squash 14 and 20 | same | Same claim | |
| R8 | 25 GPT-4V 16th percentile, 23% | same | Same claim | Bendeck text: "the 16th percentile of humans"; lookup 23% in the task table. |
| R9 | 25 "A flag can therefore cite a sound rule and still misread the slide." | same | Same claim | |
| R10 | 29 "People follow AI warnings even when the warnings are wrong, and those with power act on them." | "In several studies, people followed AI warnings ..." and "Flags also prompt action from those with the power to act." | Nuance lost | The hedge "In several studies" is gone, so the sentence states a general law. Fix: replace "People follow AI warnings even when the warnings are wrong, and those with power act on them." with "In several studies, people followed AI warnings even when the warnings were wrong, and those with power acted on them." (Carried-over issue, not new: `academic-tally.md` item 14 notes von Schenk's accusers held no power over the writer; v6 kept the clause, so it is left as is.) |
| R11 | 29 Horne 0.30 on −1 to +1 | same | Same claim | |
| R12 | 29 Lu indicator, could not tell correct from incorrect | same, with quotation (p. 12) | Same claim | Quotation paraphrased faithfully. |
| R13 | 29 Du 1.8 points of 10, red highlighting | same | Same claim | |
| R14 | 29 "People who consulted a lie-detection algorithm accused ..." | "Participants who chose to consult ..." | Nuance lost | The self-selection is lost; in von Schenk the 85% applies to those who chose to see the algorithm. Fix: replace "People who consulted a lie-detection algorithm" with "People who chose to consult a lie-detection algorithm". |
| R15 | 29 "The person checked pays." | "Costs can also fall on the person checked." | Nuance lost | Stated as a certainty; several results that follow are conditional (Madanay: only the radiologist who recommended less testing). Fix: replace "The person checked pays." with "The person checked can pay too." |
| R16 | 29 Strickland; Prike; Madanay; Chen; Reif (one sentence) | five v6 sentences | Same claim | Madanay: the "agreed less" clause is dropped, but "rated lowest of four conditions" and "no penalty" match the abstract (least likely to recommend and positively rate the undercalling radiologist, no differences among the others). Reif's "another source" became "elsewhere"; same. |
| R17 | 29 "Two results bound this." | "Two findings bound these conclusions." | Same claim | |
| R18 | 29 "In D. Kim et al. (2026) the advisor's error lowered trust and the AI's involvement changed nothing, and people resisted ... (Loconte et al., 2026)." | "the manner of the AI's involvement made no difference, so their data attribute the loss to the error rather than to the AI's contradiction" | Nuance lost | D. Kim p. 1520: "AI modality did not significantly change trust evaluations of the advisor" (modality includes no AI); p. 1521: AI-supported workflows "may modestly attenuate trust declines ... not statistically significant". "Changed nothing" overstates a null result. Fix: replace "and the AI's involvement changed nothing," with "and the AI's involvement made no significant difference," |
| R19 | 29 politicians not carried over | same | Same claim | |
| R20 | 33 "With human accusers, the reply decides the repair." | "the form of the reply shapes repair" | Nuance lost | "Decides" is stronger than the cited studies show (differences between replies, not full determination). Fix: replace "the reply decides the repair." with "the form of the reply shapes the repair." |
| R21 | 33 P. H. Kim 2004 apology over denial; Ferrin reticence no better than the weaker | two v6 sentences | Same claim | Joining puts Ferrin under the incompetence frame; Ferrin abstract (p. 893) says reticence is suboptimal after a competence violation because, like denial, it fails to signal redemption, so the frame holds. |
| R22 | 33 Greitemeyer exoneration | same | Same claim | |
| R23 | 33 Sharma further development; most studies short of start; denial weaker after competence | same, with quotation | Same claim | Sharma .txt: increase above pre-transgression level "should be considered not as 'repair' but as further development of trust". |
| R24 | 33 "With a machine accuser the picture changes." | "In these studies the accuser is human." | Nuance lost | v10 asserts a change; the only support that follows is an argument (Kempt and Nagel) and indirect evidence (Pezzo). Fix: replace "With a machine accuser the picture changes." with "With a machine accuser the picture may change." |
| R25 | 33 Kempt and Nagel burden of proof; reasons cannot settle the dispute | same | Same claim | Kempt p. 226: the physician "cannot overturn it in any meaningful way based on an exchange of reasons, as the AI is ... not receptive for diagnostic reasons. Thus, the burden of proof ... moves ... towards the proponent of the initial opinion." The remedy (a second human) was cut by design. |
| R26 | 33 "a physician who skipped an available diagnostic aid was rated higher ..." | "who had not consulted an available diagnostic aid before a poor outcome ..." | Nuance lost | The bad outcome is the setting of the whole comparison (Experiment 2 used only a bad outcome; abstract, per `read-pezzo-schilke-sharma.md`). Without it the reader may take the physician to be right, the opposite of the study. Full text could not verify (no copy reachable); the abstract supports the fix. Fix: replace "a physician who skipped an available diagnostic aid was rated higher" with "a physician who skipped an available diagnostic aid before a poor outcome was rated higher". |
| R27 | 43 "What none of this work resolves is what a check does to the audience's trust in the speaker: no study measured it, and Danry et al.'s participants rated recorded statements, not a person." | "None of this work reports the effect of a check on trust in the speaker; Danry et al.'s (2020) participants ..." | Meaning changed | "No study measured it" now reads as universal, and line 39 (four lines above) cites Chen and D. Kim as studies that measured trust in the person checked. Year also dropped. Fix: see P1. |
| R28 | 47 "... in every study above, the person judged stays silent and is rated once." | "Wherever these studies judge a person, that person remains silent and is rated once." | Meaning changed | After the regrouping, "every study above" includes P. H. Kim 2004 and Ferrin 2007 (the accused replies) and Greitemeyer (exoneration). Fix: see P2. |
| R29 | 51 three open questions | v6 hypothesis sentence and "we expect a bare denial to repair less" | Same claim | The study's own statement, turned into questions; no source claim. |
| R30 | 55 gap in one sentence | v6 "No study combines these elements." | Same claim | |
| R31 | 61 D. Kim closest design, METI | same ("checked a human advisor's work") | Same claim | D. Kim pp. 1515: proactive and reactive AI verification of the advisor's schedule; H1 uses Hendriks et al. (METI). |
| R32 | 61 "Their AI was always right, the advisor never replied, ..." | same | Same claim | D. Kim p. 1521: "it was always correct". |
| R33 | 61 "Chen et al. (2026) is next: the doctor brought in the AI himself and never answered." | not in v6's gap paragraph; in v9 | Same claim (stands alone correctly) | `chen.md`: the doctor asked whether the patient wanted to hear his AI assistant's opinion, reported only agree or disagree, gave no explanation (Appendix, p. 14); "He did not" reply. Ranking as second closest matches `verified-view.md` rule 22. |
| R34 | 61 Gupta measured belief and trust in the AI | same | Same claim | |
| R35 | 61 Sharma 8 of 94, within-subject | same | Same claim | Sharma .txt: "Only eight articles (out of 94 empirical articles) incorporated a complete or partial within-subject design". |
| R36 | 67 "Our checker treats the false flag as the normal case rather than the exception." | none (own) | Same claim | Study's own statement. |
| R37 | 67 flags drafted with a real VLM ... all wrong for this company | "Our flags are drafted with a real VLM and are all incorrect for the company presented." | Same claim | Added design details (slide images, fixed output) are the study's own. |
| R38 | 69 "Our study differs on each point." | v6 gap paragraph "The present study combines these elements" | Same claim | Placement problem, see P3. |
| R39 | 69 flags wrong, CISO answers, listener holds tool and budget, trust at every step | same | Same claim | |
| R40 | 69 "Because accusations of dishonesty and visible confidence figures change how people respond, ... and the CISO never mentions AI." | v6 tied each to its source: Loconte ("our flags concern the slide ... display neither figure") and Reif ("accordingly, our CISO never mentions AI") | Nuance lost | The reasons now carry no source. Fix: see P4. |
| R41 | 71 "Our study puts all three to the test." | none | Same claim | Placement problem, see P5. |
| R42 | 71 "... and a rise above the first rating is treated as exploratory, since even a denial with an account was the weaker reply." | "we treat a rise above the first rating as exploratory, and because even a denial with an account was the weaker reply, we expect a bare denial to repair less." | Meaning changed | v6 gave two claims with two reasons: the rise is exploratory (because trust rarely returns above its start: Greitemeyer, Sharma); a bare denial should repair less (because denial with an account was the weaker reply: P. H. Kim 2004). v10 joins the first claim to the second reason, which does not support it. Fix: see P6. |
| R43 | 73 "This study does, and asks whether ..." | v6 "The design thereby tests whether ..." | Same claim | Placement problem, see P7. |

Totals: 43 sentences; same claim 31; nuance lost 9 (R5, R10, R14, R15, R18, R20, R24, R26, R40); meaning changed 3 (R27, R28, R42).

## 4. Placement findings

| # | Line | Problem | Exact fix |
|---|---|---|---|
| P1 | 43 | "this work" and "no study" lost their scope (in v9 the sentence followed the checker paragraph directly). It now sits below line 39, which cites Chen and D. Kim as measuring trust in the person checked, so it reads as a contradiction. Year missing on Danry. | Replace "What none of this work resolves is what a check does to the audience's trust in the speaker: no study measured it, and Danry et al.'s participants rated recorded statements, not a person." with "What none of the checker studies resolves is what a check does to the audience's trust in the speaker: none of them measured it, and Danry et al.'s (2020) participants rated recorded statements, not a person." |
| P2 | 47 | "in every study above" now covers the trust-repair studies, where the accused does reply. | Replace "in every study above, the person judged stays silent and is rated once." with "in every wrong-flag study above, the person judged stays silent and is rated once." |
| P3 | 69 | "each point" has no nearby referent; the points are D. Kim's limits at line 61, two paragraphs earlier. | Replace "Our study differs on each point." with "Our study differs from these designs on each point." |
| P4 | 69 | The reasons for the flag wording, the hidden confidence figure and the CISO's silence about AI came from Loconte et al. and Reif et al., now 40 lines away and uncited here. | Replace "Because accusations of dishonesty and visible confidence figures change how people respond, every flag" with "Because accusations of dishonesty and visible confidence figures change how people respond (Loconte et al., 2026), every flag", and "and the CISO never mentions AI." with "and the CISO never mentions AI (Reif et al., 2025)." |
| P5 | 71 | "all three" refers to the three open questions at line 51, twenty lines and two headings away. | Replace "Our study puts all three to the test." with "Our study puts the three open questions on repair to the test." |
| P6 | 71 | The reason clause depends on P. H. Kim et al. (2004), no longer nearby, and is attached to the wrong claim (R42). | Replace "and a rise above the first rating is treated as exploratory, since even a denial with an account was the weaker reply." with "and a rise above the first rating is treated as exploratory, since repaired trust rarely exceeded its starting level (Greitemeyer & Sagioglou, 2015; Sharma et al., 2023)." If the denial point should stay, add after it: "A bare denial is expected to repair less, since even a denial with an account was the weaker reply (P. H. Kim et al., 2004)." |
| P7 | 73 | "This study does" now follows the repair paragraph; its referent ("No study combines ...") is at line 55, two headings up. | Replace "This study does, and asks whether" with "This study combines all four elements of the gap, and asks whether" |
| Note | 51 | "the burden of proof" refers to Kempt and Nagel, three headings above. Readable, since the heading names trust repair; optional fix: replace "meets the burden of proof." with "meets the burden of proof that a machine's contradiction places on the expert (Kempt & Nagel, 2022)." | optional |

Checked and fine: line 39's citations keep their v9 context (section 1 text, unchanged); line 61's Chen sentence stands alone correctly (see R33); line 25's "products (above)" still points upward (products are named at line 11); line 61's D. Kim and Gupta sentences carry their own citations.

## 5. Entry findings

**The three fixes of today** (compared with `reference-audit-v7.md` items 18, 21, 22 and the v6 list):

- Madanay: v6 "Madanay, F. L.," is now "Madanay, F.,". Correct: the article byline in the Europe PMC text reads "©Farrah Madanay, Laura S O'Donohue, Brian J Zikmund-Fisher".
- Opdenbusch: v6 '"Where are we on cyber?": A qualitative study' is now '"Where are we on cyber?" A qualitative study'. Correct.
- Microsoft: now "Microsoft. (n.d.). *Catch up on meetings with Microsoft Copilot in Teams*. Retrieved October 20, 2026, from https://support.microsoft.com/en-us/teams/copilot/catch-up-on-meetings-with-microsoft-365-copilot-in-teams". Correct: fetched today, HTTP 200, page title "Catch up on meetings with Microsoft Copilot in Teams | Microsoft Support", and the prompt "Where are some holes in [a meeting participant's] argument?" is on the page. The old v6 URL returns 301 to this one.

No other entry differs between v6 and v10 (diff of the two lists). v9 and v10 lists are identical.

**Crossref spot-checks** (`api.crossref.org/works/<DOI>`):

| Entry | Crossref | Verdict |
|---|---|---|
| Kempt & Nagel (2022), JME 48(4), 222–229 | Kempt, Hendrik; Nagel, Saskia K; JME 48(4) 222-229; online 2021-12-14 | Correct; the issue year 2022 is the printed one ("J Med Ethics 2022;48:222–229", PDF footer) |
| Pezzo et al. (2022), MDM 42(1), 94–104 | Pezzo, Mark V.; Nash, Brenton E. D.; Vieux; Foster-Grammer, Hannah W.; 42(1) 94-104; online 2021-05-08 | Correct (issue year 2022) |
| Strickland (1958), J Pers 26(2), 200–215 | Strickland, Lloyd H.; 26(2) 200-215; 1958 | Correct |
| Schilke et al. (2015), PNAS 112(42), 12950–12955 | same authors, 112(42) 12950-12955; 2015 | Correct |
| Gupta et al. (2026), CHI 2026, pp. 1–20 | six authors as listed; pp. 1-20; 2026 | Correct. Crossref splits "Victoria Aritonang" as family name; the PDF byline is "Nadia Victoria Aritonang" and the entry's "Aritonang, N. V." is a defensible reading. The PDF byline spells "Ranjendran"; Crossref and the entry have "Rajendran". No change needed. |
| D. Kim et al. (2026), AIES 9(3), 1514–1525 | Kim, Dennis; Daneshi, Roya; Draper, Bruce; Sreedharan, Sarath; 9(3) 1514-1525 | Correct |

**URLs:** Meta: HTTP 200, title "New Ray-Ban | Meta Smart Glasses Styles and Meta AI Updates", matches the entry. Microsoft: HTTP 200 (the first automated fetch returned 403 "Service unavailable"; two retries with a browser user agent returned 200), as above.

Result: 30 entries, 30 correct, 0 fixes needed.

## 6. Ranked fixes

1. **R42 / P6 (meaning changed): wrong reason for the exploratory rise, line 71.**
   Old: "and a rise above the first rating is treated as exploratory, since even a denial with an account was the weaker reply."
   New: "and a rise above the first rating is treated as exploratory, since repaired trust rarely exceeded its starting level (Greitemeyer & Sagioglou, 2015; Sharma et al., 2023)."
2. **R27 / P1 (meaning changed): "no study measured it" contradicts line 39, line 43.**
   Old: "What none of this work resolves is what a check does to the audience's trust in the speaker: no study measured it, and Danry et al.'s participants rated recorded statements, not a person."
   New: "What none of the checker studies resolves is what a check does to the audience's trust in the speaker: none of them measured it, and Danry et al.'s (2020) participants rated recorded statements, not a person."
3. **R28 / P2 (meaning changed): "every study above" now includes repair studies, line 47.**
   Old: "in every study above, the person judged stays silent and is rated once."
   New: "in every wrong-flag study above, the person judged stays silent and is rated once."
4. **R26 (nuance lost): Pezzo's poor outcome, line 33.**
   Old: "a physician who skipped an available diagnostic aid was rated higher"
   New: "a physician who skipped an available diagnostic aid before a poor outcome was rated higher"
5. **R40 / P4 (nuance lost): design reasons lost their sources, line 69.**
   Old: "Because accusations of dishonesty and visible confidence figures change how people respond, every flag"
   New: "Because accusations of dishonesty and visible confidence figures change how people respond (Loconte et al., 2026), every flag"
   Old: "and the CISO never mentions AI."
   New: "and the CISO never mentions AI (Reif et al., 2025)."
6. **P7 (placement): dangling "This study does", line 73.**
   Old: "This study does, and asks whether"
   New: "This study combines all four elements of the gap, and asks whether"
7. **P5 (placement): "all three", line 71.**
   Old: "Our study puts all three to the test."
   New: "Our study puts the three open questions on repair to the test."
8. **P3 (placement): "each point", line 69.**
   Old: "Our study differs on each point."
   New: "Our study differs from these designs on each point."
9. **R18 (nuance lost): D. Kim null result, line 29.**
   Old: "and the AI's involvement changed nothing,"
   New: "and the AI's involvement made no significant difference,"
10. **R5 (nuance lost): limited accuracy extended to products, line 25.**
    Old: "and their accuracy is limited."
    New: "and evaluations of related checkers report limited accuracy."
11. **R10 (nuance lost): hedge dropped, line 29.**
    Old: "People follow AI warnings even when the warnings are wrong, and those with power act on them."
    New: "In several studies, people followed AI warnings even when the warnings were wrong, and those with power acted on them."
12. **R15 (nuance lost): certainty, line 29.**
    Old: "The person checked pays."
    New: "The person checked can pay too."
13. **R14 (nuance lost): self-selection, line 29.**
    Old: "People who consulted a lie-detection algorithm"
    New: "People who chose to consult a lie-detection algorithm"
14. **R20 (nuance lost): "decides", line 33.**
    Old: "With human accusers, the reply decides the repair."
    New: "With human accusers, the form of the reply shapes the repair."
15. **R24 (nuance lost): asserted change, line 33.**
    Old: "With a machine accuser the picture changes."
    New: "With a machine accuser the picture may change."
16. **Optional, line 51:** Old: "meets the burden of proof." New: "meets the burden of proof that a machine's contradiction places on the expert (Kempt & Nagel, 2022)."
17. **Housekeeping (working notes only, not body):** line 150 "(H1 and H2 in plain words)" could read "(the two predictions in plain words)"; line 157 "Gupta only in the gap" is now "Gupta only in the closest designs". "HCI" (line 39) could be spelled out as "human-computer interaction (HCI)" at first use.
