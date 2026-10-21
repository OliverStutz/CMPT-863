# Brief: check Draft v6 (academic register) against Draft v5 and the sources (2026-10-20)

You are one of three independent checkers (A, B, C). Do not read the other checkers' files (`search/academic-check-*.md`). Write your report to `search/academic-check-<your letter>.md` and edit no other file.

## Inputs (absolute paths under `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/Project Literature Review/`)

- `Literature Review - Draft v5.md`: the verified text (checked claim by claim against the sources by three earlier checkers, then condensed and re-checked).
- `Literature Review - Draft v6.md`: the rewrite under check. It was meant to change register only.
- `search/_academic-brief.md`: what v6 was allowed to change.
- `references/`: full texts (use `pdftotext -layout` on PDFs; two are .txt files). Use them for page locators and quotations.
- `search/verified-view.md`, `search/fact-check-2026-10-19.md`, `search/read-pezzo-schilke-sharma.md`, `search/read-danry-kempt-dasilva.md`: what each paper says, as verified earlier.

## Tasks

A. Fidelity, sentence by sentence. Align every v6 sentence with its v5 counterpart. Report any claim, number, comparison, qualifier, attribution or hedge that v6 adds, drops, strengthens or weakens. Rate each as "meaning change", "nuance loss" or "fine". Check that every "how our study differs" statement in v5 survives in v6.

B. Page locators. Confirm each locator against the full text: U.S. Securities and Exchange Commission (2023) p. 85; Lowry et al. (2026) p. 783 for both quotations; Opdenbusch et al. (2025) p. 2 ("the ones in power") and p. 6 ("too shy to ask questions"); Lu et al. (2025) p. 12; Sharma et al. (2023) section heading "What Are the Limits of Trust Repair?" (the available text has no page numbers). Confirm each quotation is verbatim. Note that the Lowry "all they can do ..." quotation is a director's words that Lowry et al. reproduce from a press source; judge whether v6's attribution ("one expert director they quote") is acceptable under APA 7, and propose the exact alternative if not.

C. Academic register. Flag any remaining colloquial phrasing, inconsistent tense, missing hedge or overclaim, heading that is not a noun phrase, abbreviation not introduced at first use, or APA 7 in-text citation error. Propose exact replacement wording for each. Do not propose wording that adds a claim or a source.

D. Rules. No em-dashes anywhere. Body 1,000 to 1,500 words (opening plus the five sections, excluding the title, headings and reference list; split on whitespace). One sentence stating that results on well-known politicians are not carried over. Trust in the AI is not an outcome. Gupta et al. (2026) appears only in the gap section. Dobbs is absent. No H1 to H4 labels. Every in-text citation has a reference entry and every entry is cited. Reference list identical to v5. CISO, VLM, METI spelled out at first use.

## Output file layout

1. Verdict: accept / accept with fixes / revise.
2. Fidelity findings as a table: v6 sentence (short), v5 sentence (short), verdict, proposed fix.
3. Page locator results, one line per locator.
4. Register findings with exact replacement wording.
5. Rules checklist, pass or fail per item, with the body word count you measured.
6. Ranked list of the fixes you would apply, each with exact old text and new text. No em-dashes in your proposals.
