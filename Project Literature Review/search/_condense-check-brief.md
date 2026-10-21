# Brief: check the condensed review

Written 2026-10-19 for two independent checkers with different jobs. Each works alone and does not read the other's file. The condensed text is `../Literature Review - Draft v4.md`; the reference text it was condensed from is `../Literature Review - Draft v2.md`, whose claims were verified against full texts by three checkers (`check-A.md`, `check-B.md`, `check-C.md`, tally in `check-tally.md`).

## Checker A: fidelity

For every sentence in v4 that cites a source, find the sentence in v2 it came from and classify: same claim (wording may differ) / narrowed without loss / changed meaning / new claim not in v2. For "changed meaning" and "new claim", check the source itself (full text in `../references/` via pdftotext, or the entry in `verified-view.md`, `read-pezzo-schilke-sharma.md`, `read-danry-kempt-dasilva.md`, or `results-1` to `results-6`) and say whether the v4 sentence is still true to the source; give a corrected sentence where it is not. Confirm that every verified wording listed in `_condense-brief.md` under "What must stay" is intact. Confirm the reference list matches the in-text citations both ways and that each remaining entry is unchanged from v2.

## Checker B: instructions and readability

Count the body words yourself. Check `../Instructions.md` point by point: 1,000 to 1,500 words; at least 15 sources; thematic subsections; synthesis rather than one-by-one summaries; each theme answers what exists, what we know, what is unresolved, how the project differs; a clear gap and contribution; APA as in `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/CMPT 863 Project Abstract v8.md`. Then compare v4 with v2 for loss: list every argument, qualification or counter-result present in v2 and absent from v4, and say for each whether a CHI reviewer would miss it (yes / no, one line why). Check the author's rules (no em-dashes, plain English, politician sentence present, trust in the AI not an outcome, Gupta only in the gap, Dobbs absent, CISO and METI spelled out, no H labels). Say where v4 still over-explains and where it is now too compressed to follow.

## Output (both)

Write to the file named in your instructions: A. Verdict: accept v4 / accept with fixes / return to v2 and recondense, one sentence why. B. Findings table. C. Ranked fix list, each as old text then new text, with the word cost of each. No em-dashes. Do not edit the drafts or any other file.
