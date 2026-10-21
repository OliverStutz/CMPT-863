# Related work: structure v2 (2026-10-19, after three-reviewer validation)

Changes from v1 are those in `design-tally.md`: all three alternatives rejected, order kept, section 4 repaired, four sources added, two dropped, four facts corrected. Budget 1,330 words. Sources to read before citing are marked "read first".

## Constraints and author decisions

As in v1: 1,000 to 1,500 words; thematic synthesis answering what exists, what we know, what is unresolved, how we differ; at least 15 sources; APA. Politician studies not used as evidence, with one sentence saying so; trust in the AI not an outcome; Gupta et al. (2026) in the gap only; Dobbs declined; Kim et al. (2026) closest.

## 0. Opening (60 words)

Two sentences of scenario: a board member points a camera assistant at a CISO's slide, the tool flags it, the CISO answers. One sentence naming the five themes.

## 1. Boards oversee what they cannot judge (240 words)

- **Bundle:** SEC (2023) final rule: boards must oversee cyber risk, and the requirement of board expertise was dropped because "directors with broad-based skills" suffice. Lowry et al. (2026): non-expert directors' oversight is "mostly symbolic" even when they do what expert directors do; they believe they ask the right questions, often off professional bodies' lists; CISOs filter reports (median estimate 40 percent); "all they can do is assess the credibility of the executives in front of them". Opdenbusch et al. (2025): UK boards stay silent for fear of exposure, decisions reduce to budget, CISOs hold informational power. Proudfoot et al. (2023): boards judge CISO competence from communication, management "always on audition", 10 to 15 minutes a few times a year, "don't show up the CEO". Aamir et al. (2026): directors challenge more confidently where independent reviews are routine. Schilke, Reimann & Cook (2015), one clause, read first: the more powerful party in an exchange trusts the other less, which is why the holder's power is a variable and not scenery. Optional, only if words allow: Da Silva & Jensen (2022), the CISO as interpreter of a "dark art", read first.
- **What exists:** regulation and professional norms that ask non-experts to oversee an expert; the norm is carried by Lowry's field evidence (NACD handbook dropped as unread and redundant).
- **What we know:** formal power with the board, informational power with the CISO; the board judges the presenter's credibility, not the substance; silence by default, from fear of exposure; an independent check raises willingness to challenge; power itself lowers trust.
- **Unresolved:** interview and field evidence only; no experiment, no tool, no trust measure.
- **How we differ:** an on-the-spot independent check in the board member's hands, with trust in the CISO, the act of asking (H4a) and the budget decision as outcomes.

## 2. Checkers are available, and they are often wrong (170 words)

- **Bundle:** Products: camera assistants (Meta 2024; Apple 2024; Google 2025) and Microsoft Copilot in Teams ("Where are some holes in [a participant's] argument?"). Research system: Danry et al. (2020), Wearable Reasoner, a wearable that judges heard claims for the wearer and changed how users rated a speaker's arguments. Live checking in practice: Squash (Adair 2020): at the 2020 State of the Union, 20 of 34 pop-ups were irrelevant matches, and the lab began building an interface for a human editor to approve each pop-up. VLM reliability: Bendeck & Stasko (2025): GPT-4V at about the 16th percentile of humans on visualization literacy, knows textbook design rules, cannot read values from a chart image without the data. Wang et al. (2024, CharXiv) only if room: 47 percent on chart reasoning against 81 percent for humans.
- **What exists:** tools that let any listener check a presenter live; one research prototype; benchmarks of chart reading.
- **What we know:** a VLM looking at a slide is often wrong on values and confident on rules; live checkers produce irrelevant or wrong flags in practice.
- **Unresolved:** none of these comes with evidence on what a check does to the audience's trust in the speaker.
- **How we differ:** flags drafted with a real VLM, fixed, and wrong for the company; the false flag is the normal case.

## 3. People follow a wrong flag, and the person flagged pays (340 words)

- **Part A, wrong flags are followed:** Horne & Nevo (2025): a wrong AI warning on a true headline cut trust by 0.30 on a two-point range, all-wrong labels pushed discernment below chance, the "rare failure effect". Lu et al. (2025): an AI credibility indicator outweighed an expert's vote and was followed when wrong; people could not tell right flags from wrong. Du et al. (2026): a wrong AI-detection report made teachers grade an identical paper 1.8 points lower on 10; red highlighting made the report more credible and the drop larger. von Schenk et al. (2024): people with the power to accuse followed a fallible lie detector into action, accusing in almost 85 percent of cases when it said "lie"; the accused could not answer (moved from section 4; carries H4a with Du's intention-to-intervene result).
- **Part B, the flagged person's standing falls:** Strickland (1958): checking alone lowers trust in the person checked (grounds the "No issues found" slides). Prike et al. (2025): a human fact-check lowered competence and morality ratings of an unfamiliar, fictional candidate; target silent. Madanay et al. (2025): 1,606 adults; the radiologist who recommended less than the AI was agreed with less and rated lowest of the four conditions; the one who recommended more paid nothing. Chen et al. (2026): AI disagreement made the doctor seem lazier and lowered trust; doctor silent. Reif et al. (2025), one clause: being seen to use AI costs a professional, which is why the CISO never mentions AI.
- **Counter-results and boundary:** Kim et al. (2026): the advisor's error lowered trust, but how the AI was involved changed nothing, so trust stayed anchored to the human; this bounds the claim that an AI's contradiction as such costs the expert. Loconte et al. (2026): people resisted confident AI accusations of lying when the model's accuracy and confidence were on screen; our flags are competence-type and the checker is opaque. One sentence: results on known politicians (settled reputations, loyal audiences) are not carried over.
- **What we know:** flags are believed by default, even by evaluators with power, and they move people to act; an AI's contradiction costs the contradicted expert; in every study the expert stays silent and trust is rated once.
- **Unresolved:** what a wrong flag does to trust in a correct expert, and what the expert's reply does afterwards.
- **How we differ:** wrong flags on a correct expert, the expert answers, the listener holds the tool and the power, trust rated at each step. Maps to H1 and its expected size (1.5 to 1.8 points).

## 4. Trust in an expert, and how it is repaired (350 words)

- **Definition and measure (about 40 words):** Mayer et al. (1995): willingness to be vulnerable; ability, benevolence, integrity; propensity before any relationship; integrity weighs most early; trust leads to risk taking in the relationship, our budget decision. Hendriks et al. (2015): METI, the three factors for laypeople judging an expert, source of our three items.
- **Repair after a human accusation:** Kim et al. (2004): apology repairs a competence violation better than denial; their denial carried an account; evidence of innocence raised every measure. Ferrin et al. (2007): a reply with no information did no better than the weaker reply; no reply moved perceived competence. Greitemeyer & Sagioglou (2015): clearing restores to baseline, not reliably above. Mickelberg et al. (2025), one clause: rating at every step draws attention to the latest information. Sharma et al. (2023), one clause, read first: partial, complete and over-repair, measured against a pre-violation baseline; why "above the start" is exploratory. Optional: Shaw et al. (2003), justifications help less than excuses, which tempers H2, read first.
- **The expert and the machine:** Kempt & Nagel (2022): when an AI disagrees with the physician in charge, the physician bears an increased burden to justify the decision (the replies in the same issue dispute the asymmetry); Chen et al. (2026) take this up. Pezzo et al. (2022), read first: the only study in which a professional's own stated reason about a computer aid changed lay ratings; "I knew what it would say" raised competence ratings, "I don't trust it" lowered them. This is the nearest machine-side precedent for a reasoned answer versus a bare one, and it narrows section 5's "none lets the expert answer" to "none lets the expert answer a flag on live content".
- **What we know:** a reply repairs trust when it supplies what the listener weighs most; for competence accusations that is a signal of reliability; bare replies do not; recovery reaches baseline, rarely above; replies move integrity, not perceived competence; a stated reason about a machine changes how the expert is judged.
- **Unresolved:** repair after a machine accusation of live content; the within-session path from flag to answer; whether a bare denial lowers trust further or only repairs less.
- **How we differ:** the expert answers an AI flag with a reason or a bare denial; trust after each; a neutral group hears the same answers without flags. Maps to H2, H3 and H4b, with "above the start" and H4b exploratory.

## 5. The gap and this study (170 words)

Element by element: no published study has a wrong AI flag on a correct expert with trust in that expert as the outcome; none puts the checker in the hands of the non-expert with power; none lets the expert answer a flag on content the audience is watching; none rates trust at each step; none is in a boardroom. Kim et al. (2026) is the closest and differs on the first four; Kim et al. (2004) and Pezzo et al. (2022) supply the reply paradigm with a human accuser or an unused aid; Gupta et al. (2026) is the nearest real-time checker and rated trust in the AI, never the speaker. Optional, one clause: Hancock et al. (2020), the checker as receiver-side AI-mediated communication. Close with the contribution sentence from Abstract v8.

## Word budget

| Section | Words |
|---|---|
| 0 Opening | 60 |
| 1 Boards | 240 |
| 2 Checkers | 170 |
| 3 Wrong flags and the flagged person | 340 |
| 4 Trust and repair | 350 |
| 5 Gap | 170 |
| Total | 1,330 |

## Sources (29 core, up to 33 with optional)

In Abstract v8 (15 after today's fix): Aamir 2026; Chen 2026; Ferrin 2007; Greitemeyer & Sagioglou 2015; Hendriks 2015; Kempt & Nagel 2022; Kim D. 2026; Kim P. 2004; Madanay 2025; Mayer 1995; Mickelberg 2025; Prike 2025; Reif 2025; Strickland 1958; von Schenk 2024.
Read in the pass (7): Opdenbusch 2025; Lowry 2026; Proudfoot 2023; Du 2026; Lu 2025; Loconte 2026; Horne & Nevo 2025.
To read before citing (4): Pezzo et al. 2022 (open access, USF repository); Schilke et al. 2015 (open access, PMC); Danry et al. 2020 (ACM, open); Sharma et al. 2023 (SFU).
Gap only (1): Gupta 2026.
Products and regulation (5): SEC 2023; camera assistants (Meta, Apple, Google, one entry); Microsoft Copilot in Teams; Adair 2020 (Squash); Bendeck & Stasko 2025.
Optional (4): Da Silva & Jensen 2022; Shaw et al. 2003; Hancock et al. 2020; Wang et al. 2024.

## Hypothesis map

| Hypothesis | Sections |
|---|---|
| H1 flag lowers trust | 3 (size: Horne & Nevo, Du; mechanism: Lu; action: von Schenk; boundaries: Kim 2026, Loconte) |
| H2 reasoned answer repairs; above start exploratory | 4 (Kim 2004, Ferrin, Greitemeyer, Sharma, Pezzo) |
| H3 bare denial repairs less | 4 (Ferrin reticence; Kim 2004 denial with account; Pezzo "I don't trust it") |
| H4a checker holders ask more | 1 (Opdenbusch, Proudfoot norm, Aamir) and 3 (von Schenk, Du) |
| H4b amplification, exploratory | 4 (Kempt & Nagel) with Kim 2026's null stated in 3 |
| Setting and measures | 1 (boards, Schilke), 4 (Mayer, Hendriks) |

## Next steps

1. Fetch and read the four "read first" sources (three are open access).
2. Draft the review from this structure, 1,330 words.
3. Independent cross-check of every sentence against the full texts before submission.
