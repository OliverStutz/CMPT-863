# Related work: structure proposal v1 (2026-10-19)

Drafted from `../Instructions.md`, the reading pass (`verified-view.md`, 22 design rules), the selection reviewers' tally (`../Literature Review - Candidate papers.md`), the novelty check (`novelty-tally.md`) and Abstract v8 (`../../CMPT 863 Project Abstract v8.md`). This is a plan for the review, not the review. To be validated by three independent reviewers before drafting.

## Constraints (from Instructions.md)

- 1,000 to 1,500 words, thematic subsections, synthesis not one-by-one summaries.
- Each theme answers four questions: what exists; what we know; what is unresolved; how our project differs and extends.
- At least 15 sources; scholarly plus, where appropriate, industry, trade and products.
- APA as in the abstract.

## Author's decisions that bind the structure

- Politician studies are not evidence for or against H1 to H3 (Chae & Tewksbury 2026, Thorson 2016, Schiff et al. 2025 excluded). Prike et al. (2025) stays because it is an abstract-phase design source (fictional, unfamiliar candidates; the 0-to-10 rating and profile are borrowed from it), not a result about known politicians.
- Trust in the AI is not an outcome (rule 10); the review is about human-to-human trust when one human is augmented by a checker.
- Gupta et al. (2026) appears only in the gap statement, as the nearest real-time checker that never rated the speaker.
- Dobbs et al. (2026) declined.
- Kim et al. (2026) is the closest design (rule 22); Chen et al. (2026) second.

## Source pool and status

| Group | Sources | Status |
|---|---|---|
| In Abstract v8 (14) | Aamir 2026; Chen 2026; Ferrin 2007; Greitemeyer & Sagioglou 2015; Hendriks 2015; Kim D. 2026; Kim P. 2004; Madanay 2025; Mayer 1995; Mickelberg 2025; Prike 2025; Reif 2025; Strickland 1958; von Schenk 2024 | all full texts checked (abstract phase or verified-view) |
| Read in the pass, not yet in the abstract (7) | Opdenbusch 2025; Lowry 2026; Proudfoot 2023; Du 2026; Lu 2025; Loconte 2026; Horne & Nevo 2025 | verified-view entries 5, 8, 13, 15, 16, 17, 18 |
| Gap-only (1) | Gupta 2026 | one sentence |
| Industry, regulation, products | SEC 2023 final rule; NACD & ISA 2023 handbook; camera assistants (Meta 2024, Apple 2024, Google Gemini Live 2025); Microsoft Copilot in Teams (support page); Squash (Adair 2020, Duke Reporters' Lab) | entries in `results-5-board-products.md`, verified 2026-10-18 |
| VLM reliability | Wang et al. 2024 (CharXiv, NeurIPS); Bendeck & Stasko 2025 (TVCG) | results-5, Bendeck full text in references |
| Optional, unread (use only if a subsection needs it) | Da Silva & Jensen 2022 (CSCW); Shaw, Wild & Colquitt 2003 (JAP meta-analysis); Sharma, Schoorman & Ballinger 2023 (J. Management review); Pezzo et al. 2022; Hancock et al. 2020 | results files; full texts partly in references |

## Proposed structure

### 0. Opening (about 60 words)

The scenario in two sentences: a board member points a camera assistant at a CISO's slide; the tool flags it; the CISO answers. Then the five themes in one sentence each.

### 1. Boards oversee what they cannot judge (about 260 words)

- **Bundle:** SEC 2023 (regulation requires board oversight of cyber risk and dropped the requirement of board expertise, "directors with broad-based skills" suffice); NACD & ISA 2023 (the professional norm is that directors ask management scripted questions); Lowry et al. 2026 (non-expert directors' oversight is "mostly symbolic" even when they do what experts do; they believe they ask the right questions; CISOs filter reports, median estimate 40 percent; "all they can do is assess the credibility of the executives in front of them"); Opdenbusch et al. 2025 (UK boards stay silent for fear of exposure as novices; decisions reduce to budget; CISOs hold informational power); Proudfoot et al. 2023 (boards judge CISO competence from communication; management "always on audition"; 10 to 15 minutes a few times a year; "don't show up the CEO"); Aamir et al. 2026 (directors challenge more confidently where independent reviews are routine). Optional: Da Silva & Jensen 2022 (the CISO as interpreter of a "dark art").
- **What exists:** regulation and norms that ask non-experts to oversee an expert.
- **What we know:** the board holds formal power and the CISO informational power; the board judges the presenter's credibility, not the substance; silence is the default and its cause is fear of exposure; an independent check raises willingness to challenge.
- **Unresolved:** all of this is interview or field evidence; no experiment, no tool, no trust measure.
- **How we differ:** we put an on-the-spot independent check in the board member's hands and measure what it does to trust in the CISO and to the act of asking (H4a) and to the budget decision.

### 2. Checkers are available, and they are often wrong (about 220 words)

- **Bundle:** products: camera assistants (Meta 2024; Apple 2024; Google 2025) and Microsoft Copilot in Teams, whose suggested prompt is "Where are some holes in [a participant's] argument?"; Squash (Adair 2020: a live checker that produced more wrong matches than right ones at the 2020 State of the Union, after which a human editor was added). VLM reliability: Wang et al. 2024 (GPT-4o 47 percent on chart reasoning against 81 percent for humans); Bendeck & Stasko 2025 (GPT-4V at about the 16th percentile of humans on visualization literacy; knows textbook design rules, cannot read values without the data). Gupta et al. 2026 goes to section 5, not here.
- **What exists:** tools that let any listener check a presenter live, and benchmarks of how they read charts.
- **What we know:** a VLM looking at a slide image is often wrong on numbers while confident on textbook rules; live checkers produce false flags in practice.
- **Unresolved:** none of these products or benchmarks comes with evidence on what a check does to the audience's trust in the speaker.
- **How we differ:** our flags are drafted with a real VLM, fixed, and wrong for the company; the design treats the false flag as the normal case, not the exception.

### 3. People follow a wrong flag, and the person flagged pays (about 300 words)

- **Bundle, part A, wrong flags are followed:** Horne & Nevo 2025 (a wrong AI warning on a true headline cut trust by 0.30 on a two-point range; all-wrong labels pushed discernment below chance; "rare failure effect"); Lu et al. 2025 (an AI credibility indicator outweighed an expert's vote and was followed when wrong; people cannot tell right from wrong flags); Du et al. 2026 (a wrong AI-detection report made teachers grade an identical paper 1.8 points lower on 10; red highlighting made the report more credible and the drop larger).
- **Bundle, part B, the flagged person's standing falls:** Strickland 1958 (checking alone lowers trust in the person checked); Prike et al. 2025 (a human fact-check lowered competence and morality ratings of an unfamiliar candidate; silent target); Madanay et al. 2025 (1,606 adults: the radiologist who recommended less than the AI was rated the worst provider; the one who recommended more paid nothing); Chen et al. 2026 (AI disagreement made the doctor seem lazier and lowered trust; doctor silent); Kim et al. 2026 (AI monitoring of an advisor: errors lowered trust, the AI's involvement did not change the penalty, trust stayed anchored to the human).
- **Counter-result and boundary:** Loconte et al. 2026 (people resisted confident AI accusations of lying when the model's accuracy and confidence were on screen; our flags are competence-type and the checker is opaque).
- **What we know:** flags are believed by default, even by evaluators with power; an AI's contradiction costs the contradicted expert; in every study the expert stays silent and trust is rated once.
- **Unresolved:** what a wrong flag does to trust in a correct expert, and what the expert's reply does afterwards.
- **How we differ:** wrong flags on a correct expert, the expert answers, the listener holds the tool and the power, trust rated at each step. Maps to H1 and gives its expected size.

### 4. Trust in an expert, and how it is repaired (about 300 words)

- **Bundle, definition and measure:** Mayer et al. 1995 (trust as willingness to be vulnerable; ability, benevolence, integrity; propensity before any relationship; integrity weighs most early; trust leads to risk taking in the relationship, which is our budget decision); Hendriks et al. 2015 (METI: laypeople's judgement of an expert forms expertise, integrity and benevolence; source of our three items).
- **Bundle, repair:** Kim et al. 2004 (hiring-manager role; apology repairs competence violations better than denial; their denial carried an account; evidence of innocence raised every measure); Ferrin et al. 2007 (reticence, a reply with no information, did no better than the weaker reply; no reply moved perceived competence); Greitemeyer & Sagioglou 2015 (clearing restores to baseline, not reliably above); Mickelberg et al. 2025 (retraction removed the effect in two of three; rating at every step draws attention to the latest information). Optional: Shaw et al. 2003 (excuses beat justifications) and Sharma et al. 2023 (partial versus complete repair) if words allow.
- **Bundle, the burden of proof and the AI penalty:** von Schenk et al. 2024 (a fallible lie detector raised accusations to 85 percent when it said "lie"; the accused could not answer); Chen et al. 2026's point that AI disagreement shifts the burden of proof onto the expert; Reif et al. 2025 (using AI costs the professional; why the CISO never mentions AI).
- **What we know:** a reply repairs trust when it supplies what the listener weighs most; for competence accusations that is a signal of reliability; bare replies do not; recovery reaches baseline, rarely above; replies move integrity, not perceived competence.
- **Unresolved:** repair after a machine accusation; the within-session path from flag to answer; whether a bare denial lowers trust further or only repairs less.
- **How we differ:** the expert answers an AI flag with a reason or a bare denial; trust is rated after each; a neutral group hears the same answers without flags. Maps to H2, H3 and H4b, and sets the H2 "above the start" claim as exploratory.

### 5. The gap and this study (about 150 words)

Element by element: no published study has a wrong AI flag on a correct expert with trust in that expert as the outcome; none puts the checker in the hands of the non-expert with power; none lets the expert answer; none rates trust at each step; none is in a boardroom. Kim et al. 2026 is the closest and differs on the first four; Kim et al. 2004 supplies the repair paradigm with a human accuser; Gupta et al. 2026 is the nearest real-time checker and rated trust in the AI, never the speaker. Close with the contribution sentence from Abstract v8.

## Word budget

| Section | Words |
|---|---|
| 0 Opening | 60 |
| 1 Boards | 260 |
| 2 Checkers | 220 |
| 3 Wrong flags and the flagged person | 300 |
| 4 Trust and repair | 300 |
| 5 Gap | 150 |
| Total | 1,290 (room to 1,500 for the optional sources) |

Source count without optional items: 14 (abstract) + 7 (read) + 1 (Gupta) + 5 product or regulatory sources (SEC, NACD, camera assistants as one, Copilot, Squash) + 2 VLM benchmarks = 29.

## Alternative bundlings to decide

- **Alt A:** merge section 2 into section 3 as its opening (checkers exist and err, then people follow them). Saves the section break; loses the product paragraph's visibility.
- **Alt B:** give H4 its own short section, "What a checker does to the person holding it" (von Schenk 2024; Aamir 2026; Opdenbusch 2025 on silence; Loconte 2026; Du 2026 on intention to intervene), taking those papers out of sections 1, 3 and 4. The selection reviewers proposed this as an option. Cost: words and some double-citing.
- **Alt C:** fold section 5 into the end of section 4 and put the element-by-element gap into the opening instead.

## Hypothesis map

| Hypothesis | Sections that carry it |
|---|---|
| H1 flag lowers trust | 3 (size from Horne & Nevo, Du; mechanism from Lu; boundary from Loconte) |
| H2 reasoned answer repairs, above start exploratory | 4 (Kim 2004, Ferrin 2007, Greitemeyer, Mickelberg) |
| H3 bare denial repairs less | 4 (Ferrin 2007 reticence, Kim 2004 denial with account) |
| H4a checker holders ask more | 1 (Opdenbusch silence, Aamir, Proudfoot norm) and 4 (von Schenk) |
| H4b amplification, exploratory | 4 (burden of proof) with Kim 2026's null stated |
| Setting and measures | 1 (boards), 4 (Mayer, Hendriks) |
