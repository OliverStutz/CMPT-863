## Editor's notes

Draft v5, 2026-10-19, from Draft v4 with the fixes in `search/condense-check-A.md` and `search/condense-check-B.md` (section C of each). Restored wordings are taken from Draft v2. No claim, source or number that is not in v2.

### Changes, old then new

1. Checker A fix 1 (referent of the interview sentence).
    - Old: "This evidence comes from interviews, with no tool or trust measure; we add both."
    - New: "The board evidence comes from interviews, with no tool or trust measure; we add both."
2. Strickland (1958) and Prike et al. (2025) restored in section 3, second paragraph, before the politician sentence (Checker B fix 2). Gives the politician sentence its context again.
    - Old: "People agreed less with a radiologist who recommended less testing than an AI"
    - New: "Even checks that find nothing can cost trust (Strickland, 1958), and a human fact-check lowered an unfamiliar candidate's ratings of competence and morality (Prike et al., 2025). People agreed less with a radiologist who recommended less testing than an AI"
3. Reif et al. (2025) restored with the design clause, in the v2 wording that keeps the control condition (Checker B fix 7, long form).
    - Old: "An AI assistant's disagreement made a doctor seem lazier and lowered trust in him (Chen et al., 2026)."
    - New: "An AI assistant's disagreement made a doctor seem lazier and lowered trust in him (Chen et al., 2026). Employees described as getting help from AI were rated lazier and less competent than those who got the same help elsewhere (Reif et al., 2025), so our CISO never mentions AI."
4. Sharma et al. competence/denial point restored in the full v2 wording (Checker B fix 1, long form).
    - Old: "... find that trust falls short of the start."
    - New: "... find that trust falls short of the start and that denial works less well after competence than after integrity violations."
5. "by a third party" restored (Checker B fix 5).
    - Old: "Exoneration restored an accused researcher"
    - New: "Exoneration by a third party restored an accused researcher"
6. Reference entries for Prike et al. (2025), Reif et al. (2025) and Strickland (1958) added back exactly as in v2, in alphabetical order (Prike and Reif after Pezzo; Strickland after Sharma). The stray blank line after `## References` removed (Checker A fix 2, Checker B fix 8).
7. Schilke et al. made followable (Checker B fix 3 plus the caveat as its own clause). Same claims: the setting, the result, and the caveat that the power was alternatives, not authority.
    - Old: "The party with better alternatives trusted its partner less (Schilke et al., 2015), though that power was having alternatives, not authority."
    - New: "In exchange experiments, the party with better alternatives trusted its partner less (Schilke et al., 2015); that power came from having alternatives, not from authority."
8. Pezzo et al. shortened (Checker B fix 6, partly). Dropped "stated" and "already"; "not trusting" to "distrusting". Kept "available" and "in neither case below one with no aid mentioned" (the verified v2 sense), since Checker B's "never below one with no aid" was flagged for confirmation and "available" is part of the finding. Still one cautious sentence.
    - Old: "The nearest evidence on an expert's stated reason is indirect: a physician who skipped an available diagnostic aid before a bad outcome was rated higher for believing they already knew what it would say than for not trusting it, and in neither case below one with no aid mentioned (Pezzo et al., 2022)."
    - New: "The nearest evidence on an expert's reason is indirect: a physician who skipped an available diagnostic aid before a bad outcome was rated higher for believing they knew what it would say than for distrusting it, and in neither case below one with no aid mentioned (Pezzo et al., 2022)."
9. Step 4, power disagreement restored in one clause (Checker B fix 4, v2 wording "trace the deference to low expertise"). Body after steps 1 to 3 was 1,263, so the gate was met; with this clause it is 1,274.
    - Old: "because they decide what reaches the board; we read this as formal power"
    - New: "because they decide what reaches the board, while Lowry et al. (2026) trace the deference to low expertise; we read this as formal power"

### Body word count (split on whitespace, same method as v1 to v4)

Opening 48; boards 209; checkers 176; wrong flags 325; trust and repair 327; gap 189; total 1,274 (v4: 1,187; after steps 1 to 3: 1,263). A grep for the em-dash character finds none.

### References

30 cited, 30 listed (v4: 27; added Prike, Reif, Strickland). Checked both ways: every in-text citation has an entry with the same year, and every entry is cited.

### Not applied, and why

- Mickelberg et al. (2025) clause (step 4). The shortest form of the v2 wording ("rating after every step may draw attention to the latest information (Mickelberg et al., 2025)", best placed at the neutral group in the gap section to avoid v2's "both our groups" problem) adds about 16 words, which would take the body to about 1,290, past the 1,280 ceiling. Its reference entry is also only in v1, not v2. Not added; no entry added.
- Checker B fix 6, "about a machine" in the Pezzo lead-in: not restored, because step 3 asked for a shorter sentence; the sentence names the diagnostic aid itself.
- Checker B fix 6, "never below one with no aid": not used; kept the verified "in neither case below one with no aid mentioned".
- Checker B fix 7, short form of Reif: not used; the v2 wording keeps the control condition ("than those who got the same help elsewhere") and "never mentions AI" as asked.
- Checker B fix 1, short form of Sharma ("after competence violations"): not used; kept the full v2 comparison with integrity violations.
- Checker B's optional topic sentence for section 3 and merging the Mayer and METI sentences: not in the task; not applied.
- Proudfoot et al. (2023): not restored (both checkers agree the Lowry et al. quotation carries the point; not in the task).

### Main-session edit after reading v5 (2026-10-19)

- Section 2, last sentence: old "Our flags are drafted with a real vision-language model." new "Our flags are drafted with a real vision-language model and are all wrong for the company." Reason: the clause is the section's "how we differ" point and was lost in condensing. Body is now 1,282 words.

## Draft v6, 2026-10-20: academic register

From Draft v5 (body 1,282 words). Register only: no claim, source, number or comparison added or removed, and the reference list is byte-identical to v5 (30 entries, all cited, checked both ways). Brief: `search/_academic-brief.md`. Three independent Opus checks: `search/academic-check-A.md`, `-B.md`, `-C.md` (brief `search/_academic-check-brief.md`), all "accept with fixes", 0 rule failures; tally and decisions in `search/academic-tally.md`.

### Changes, by kind

1. Headings became noun phrases: "Boards oversee what they cannot judge" to "Board oversight of cybersecurity by non-experts"; "Checkers are available, and often wrong" to "Listener-held checking tools and their accuracy"; "People follow a wrong flag, and the person flagged pays" to "Effects of erroneous flags on judgment and on the person flagged"; "Trust in an expert, and how it is repaired" to "Trust in experts and its repair"; "The gap and this study" to "Research gap and the present study".
2. Topic sentences added where a section opened on a bare finding ("In practice, such oversight appears limited."; "Evaluations of related checkers, however, report limited accuracy."; "In several studies, people followed AI warnings even when the warnings were wrong."; "Costs can also fall on the person checked."; "With human accusers, the form of the reply shapes repair."), each hedged after the checks.
3. Colloquial wording replaced: "pays" and "wins it back" (now "costs" and "recovers it"), "lost" (now "was the weaker reply"), "checkers on offer", "shows whether" (now "tests whether"), "we read this as" (now "we interpret these accounts as"), "right on" (now "correct on"), "bare" kept as the study's own term.
4. APA locators on every direct quotation: SEC (2023) p. 85; Lowry et al. (2026) p. 783 for "mostly symbolic" and p. 798 for the director's words; Opdenbusch et al. (2025) p. 2 and p. 6; Lu et al. (2025) p. 12; Sharma et al. (2023) "What Are the Limits of Trust Repair?" section (no page numbers in the available text); Microsoft (n.d.) "Have a more effective meeting" section, checked on the live support page 2026-10-20 (prompt verbatim).
5. Lowry attribution corrected. v5 cited the director's words to Lowry et al. without a page; the words appear on p. 783 as an epigraph signed "Former NASDAQ Board Director" and on p. 798 as interview data from "D-9 expert". v6 says "one expert director they interviewed" with p. 798. (The main session's first reading, that it was a press quotation via Rundle 2023, was wrong: the Rundle citation belongs to the 12 percent statistic in the next column. All three checkers caught this.)
6. Opdenbusch attribution made exact: the words "too shy to ask questions" are the authors' summary of four participants, so v6 reads "Opdenbusch et al. (2025) report that, according to four participants, directors would be ..." (the source says "would be").
7. "percent" to "%" after numerals (APA 7); "judgement" to "judgment" (APA, American spelling as in the reference list).
8. Mayer definition recast ("Following Mayer et al. (1995), trust is ... and it leads to risk taking, which in the present study is the budget decision.").
9. The 14 post-check fixes in `search/academic-tally.md` (3 of 3 on seven items, 2 of 3 on one, single-checker items applied where they were correct APA or closer to v5).

### Body word count (split on whitespace, same method as v1 to v5)

Opening 66; boards 258; checkers 200; wrong flags 368; trust and repair 376; gap 195; total 1,463 (v5: 1,282; limit 1,500). No em-dash in the file.

### Open for Oliver

- Approve v6 as the final text (the file has no notes section; Final = v6 as it stands).
- The v5 open items are unchanged: Mickelberg et al. (2025) is in Abstract v8 but not in the review; Pezzo et al. (2022) is cited from the abstract only.
- Spelling choice: v6 uses American spelling ("judgment") to match APA and the reference list; Abstract v8 has one British form ("organisation"). Either is fine if used consistently.
