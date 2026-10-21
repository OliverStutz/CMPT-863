# Condense check B: instructions and readability

Checker B, 2026-10-19. Compared `Literature Review - Draft v4.md` with `Literature Review - Draft v2.md`, `Instructions.md` and `CMPT 863 Project Abstract v8.md`. Did not read `condense-check-A.md`.

## A. Verdict

Accept with fixes: v4 meets every point of the instructions at 1,187 words and keeps the author's rules, but condensing cut three things a reader needs (the competence/denial qualification from Sharma et al., the context that makes the politician sentence make sense, and the "check alone costs trust" result behind the "No issues found" slides), and all three come back for about 60 words, ending near 1,250.

## B. Findings

### B1. Word count

Body (opening plus the five sections, without title, headings, references and notes), split on whitespace: opening 48, boards 193, checkers 176, wrong flags 267, trust and repair 314, gap 189, total **1,187**. This matches the condensing notes. v2 by the same method: 1,448.

### B2. Instructions.md checklist

| Point | Status | Note |
|---|---|---|
| 1,000 to 1,500 words | Met | 1,187 |
| At least 15 sources, scholarly plus industry and products where fitting | Met | 27 cited; Meta, Microsoft, Squash (Adair) and the SEC rule cover industry, products and regulation |
| Thematic subsections | Met | Four themes plus the gap |
| Synthesis, not one-by-one summary | Met, weakest in two places | Section 2 and the first paragraph of section 3 are still one sentence per paper. Each paragraph has a claim the papers serve ("Checkers ... often wrong"; "Flags also move people with power to act"), so it passes, but section 3 lost its opening topic sentence and now starts directly on Horne and Nevo |
| Each theme: what exists | Met | Regulation; products; studies of warnings and AI disagreement; trust models and repair studies |
| Each theme: what we know | Met | |
| Each theme: what is unresolved | Met, thinner in section 1 | v2's "The sources disagree on power" (Opdenbusch et al. against Lowry et al.) was the one open question in section 1; v4 keeps only "interviews, with no tool or trust measure" |
| Each theme: how the project differs | Met | "we add both"; "drafted with a real vision-language model"; "our CISO answers"; the hypothesis sentence |
| Clear gap and contribution | Met | Gap section names D. Kim et al. (2026) as closest, lists the three differences, ends on the contribution sentence |
| APA as in Abstract v8 | Met | Reference entries have the same form as v8. In-text "D. Kim et al." and "P. H. Kim et al." differ from v8's bare "Kim et al." but are the correct APA 7 form for two first authors with the same surname (keep v4, and fix v8 if it is reused). All 27 in-text citations are in the list and all 27 entries are cited. One stray blank line after `## References` |

### B3. Author's rules

| Rule | Status |
|---|---|
| No em-dashes | Met (grep finds none; the en-dashes in page ranges and the minus sign in "−1 to +1" are correct) |
| Plain English | Met, with two hard sentences (Schilke et al. and Pezzo et al., see B5) |
| Politician sentence present | Present, but now orphaned: no politician or candidate study is cited anywhere in v4, so "Results on well-known politicians ... are not carried over" answers a question the reader was never asked. In v2 it followed Prike et al.'s unfamiliar candidate |
| Trust in the AI not an outcome | Met (trust in the AI appears only as what Gupta et al. measured) |
| Gupta only in the gap | Met |
| Dobbs absent | Met |
| CISO and METI spelled out at first use | Met |
| No H labels | Met |

### B4. Losses from v2 to v4

| # | In v2, absent from v4 | CHI reviewer would miss? |
|---|---|---|
| 1 | Sharma et al.: "denial works less well after competence than after integrity violations" | Yes: our flags are competence-type, so this is the strongest published reason to expect the bare denial to fail, and the denial prediction now rests only on P. H. Kim et al. |
| 2 | Prike et al. (2025): a human fact-check lowered an unfamiliar candidate's competence and morality ratings | Yes: the only result where a fact-check costs a person standing, and the reason the politician sentence exists; the abstract borrows its profile and 0-to-10 rating |
| 3 | Strickland (1958): checks that find nothing still cost trust | Yes: a reviewer will ask whether checking alone lowers trust, and the abstract's "No issues found" slides and slide-1 validity check answer exactly that |
| 4 | "The sources disagree on power ... Lowry et al. (2026) trace the deference to low expertise instead" | Yes, mildly: it is section 1's unresolved question and gives "we read this as" something to reconcile; without it the reading looks like a gloss on one paper |
| 5 | Greitemeyer and Sagioglou: clearing came "by a third party" (v4: "Exoneration") | Yes, mildly: the contrast with a CISO who answers for himself is the point of citing it; costs 4 words |
| 6 | Reif et al. (2025): AI-assisted employees rated lazier and less competent, "so our CISO never mentions AI" | No for a CHI reviewer (a design note that belongs in the method); the instructor may notice it left the abstract's related work |
| 7 | Proudfoot et al. (2023): board's view of the CISO's competence set by communication | No: the Lowry et al. credibility quotation carries the same point |
| 8 | Schilke et al.: "In exchange experiments" and "from the same starting level its trust fell further during a negotiation" | No for the second clause; the lost "In exchange experiments" is a readability loss (see B5) |
| 9 | Squash: the lab began building a human-editor approval step | No: interesting, but nothing in v4 turns on it |
| 10 | Sharma et al.: "a return to the starting level complete repair" | No: "further development" carries the exploratory point |
| 11 | "Such checkers are often wrong." as a pivot | No: the heading carries it |
| 12 | Section 3: "and is rated after each step" | No: stated in the gap section |
| 13 | Section 2: flags "are all wrong for the company" | No: stated in the gap section |
| 14 | Pezzo et al.: "reason about a machine" (v4: "stated reason") | No, but "about a machine" says why Pezzo et al. is relevant at all; free to restore by swapping words |
| 15 | Sample sizes (214 teachers, 1,606 adults, N = 157) and "(SEC, 2023)" | No: not benchmarks |
| 16 | SEC quotation shortened ("directors with broad-based skills in risk management and strategy") | No: "broadly skilled directors" keeps the meaning |

Losses listed: 16. A CHI reviewer would miss 5 (items 1 to 5), of which 2 are mild.

### B5. Over-explained and too compressed

Too compressed to follow:
- Schilke et al.: "The party with better alternatives trusted its partner less (Schilke et al., 2015), though that power was having alternatives, not authority." With "In exchange experiments" cut, the reader does not know what setting or what partner, and the sentence never says what it means for us (the board, holding power, may start out trusting the CISO less).
- The politician sentence (B3): orphaned.
- Section 3 opens with no claim; the reader meets a 0.30 shift before being told what the paragraph argues. Restoring a four-word topic sentence would help, but the heading nearly does it; optional.

Still over-explained:
- Pezzo et al. sentence (about 50 words, three clauses of conditions) is the hardest sentence in the review. It can lose about 10 words.
- Section 4 is the longest section (314). The Mayer et al. sentence and the METI sentence could be one, but each carries a needed point; leave unless words are needed.
- The Loconte et al. clause "and show neither figure" and the Kempt and Nagel second sentence are justified; leave.

### B6. The four sources dropped in condensing

Three are cited in Abstract v8's related work (Strickland, Prike, Reif); Proudfoot is not. Mickelberg et al. (2025), also cited in v8, was already dropped in v2 (cut 23); not reviewed here, but the instructor may notice that four of v8's sources are gone from a review that is meant to expand it.

| Source | Would a CHI reviewer or the instructor miss it? | Clause to restore, with word cost |
|---|---|---|
| Proudfoot et al. (2023) | No: Lowry et al.'s "assess the credibility of the executives" makes the same point, and it was not in the abstract | None |
| Prike et al. (2025) | Yes: the only study of a fact-check costing a person standing, in v8, and the setup the politician sentence needs | Section 3, second paragraph, first sentence: "A human fact-check lowered an unfamiliar candidate's ratings of competence and morality (Prike et al., 2025)." 16 words |
| Strickland (1958) | Yes: answers the reviewer's "does checking alone cost trust?" and justifies the "No issues found" slides in v8 | Same paragraph, before Prike: "Even checks that find nothing can cost trust (Strickland, 1958)." 10 words |
| Reif et al. (2025) | Instructor yes, CHI reviewer no: a known penalty on trust in a person who uses AI, cited in v8 and behind the rule that the CISO never mentions AI | After Chen et al.: "Merely being described as using AI made employees seem lazier and less competent (Reif et al., 2025), so our CISO never mentions it." 23 words; optional |

## C. Ranked fix list

1. Restore the competence/denial qualification (loss 1). Cost +12.
   - Old: "Sharma et al. (2023) call a rise beyond the starting level "further development of trust", and most repair studies, they note, find that trust falls short of the start."
   - New: "Sharma et al. (2023) call a rise beyond the starting level "further development of trust", and most repair studies, they note, find that trust falls short of the start and that denial works less well after competence violations."
   - Note for Checker A or the author: v2 said "less well after competence than after integrity violations"; keep that full wording (+3) if the shorter form is judged to change meaning.

2. Restore Strickland and Prike, which also gives the politician sentence its context (losses 2 and 3). Cost +26.
   - Old: "People agreed less with a radiologist who recommended less testing than an AI"
   - New: "Even checks that find nothing can cost trust (Strickland, 1958), and a human fact-check lowered an unfamiliar candidate's ratings of competence and morality (Prike et al., 2025). People agreed less with a radiologist who recommended less testing than an AI"
   - Add both entries back to the reference list from v2 (unchanged).

3. Restore the "In exchange experiments" setting for Schilke et al. (readability). Cost +3.
   - Old: "The party with better alternatives trusted its partner less (Schilke et al., 2015)"
   - New: "In exchange experiments, the party with better alternatives trusted its partner less (Schilke et al., 2015)"

4. Restore the power disagreement (loss 4). Cost +11.
   - Old: "because they decide what reaches the board; we read this as formal power"
   - New: "because they decide what reaches the board, while Lowry et al. (2026) trace the deference to low expertise; we read this as formal power"

5. Restore "by a third party" (loss 5). Cost +4.
   - Old: "Exoneration restored an accused researcher"
   - New: "Exoneration by a third party restored an accused researcher"

6. Trim Pezzo et al. and restore "about a machine". Saves about 8.
   - Old: "The nearest evidence on an expert's stated reason is indirect: a physician who skipped an available diagnostic aid before a bad outcome was rated higher for believing they already knew what it would say than for not trusting it, and in neither case below one with no aid mentioned (Pezzo et al., 2022)."
   - New: "The nearest evidence on an expert's reason about a machine is indirect: a physician who skipped a diagnostic aid before a bad outcome was rated higher for believing they knew what it would say than for distrusting it, and never below one with no aid (Pezzo et al., 2022)."
   - Checker A or the author should confirm "never below one with no aid" against the verified v2 wording before taking this.

7. Optional: restore Reif et al. (loss 6). Cost +23, plus the reference entry.
   - Old: "An AI assistant's disagreement made a doctor seem lazier and lowered trust in him (Chen et al., 2026)."
   - New: "An AI assistant's disagreement made a doctor seem lazier and lowered trust in him (Chen et al., 2026). Merely being described as using AI made employees seem lazier and less competent (Reif et al., 2025), so our CISO never mentions it."
   - v2's verified wording was "Employees described as getting help from AI were rated lazier and less competent than those who got the same help elsewhere"; use it (+5) if the shorter form is judged to drop the control condition.

8. Remove the stray blank line after `## References`. Cost 0.

Fixes 1 to 5 add 56 words (body about 1,243, 29 sources); with fix 6, about 1,235; with fix 7 as well, about 1,258. All stay within the limit.
