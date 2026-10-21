# Brief: draft the literature review

Written 2026-10-19. One writer. The draft goes to `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/Project Literature Review/Literature Review - Draft v1.md`. Three independent checkers will verify every sentence against the sources, so write nothing you cannot point to.

## Read, in this order

1. `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/Project Literature Review/Instructions.md` (the assignment).
2. `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/CMPT 863 Project Abstract v8.md` (the study the review situates; its appendix is the related work you are expanding; its reference list sets the APA style).
3. `search/_review-design-v2.md` (the structure: five sections, bundles, the four questions each section answers, word budget, source list, hypothesis map). This is binding unless a read file below says a paper does not support its planned use.
4. `search/verified-view.md` (what each read paper says, checked against the full text; the 22 design rules; the author's decisions and exclusions). Facts about the read papers come from here.
5. `search/read-pezzo-schilke-sharma.md` and `search/read-danry-kempt-dasilva.md` (the sources read today; follow their "Consequence" sections).
6. `search/fact-check-2026-10-19.md` (four corrected wordings; use the corrected ones).
7. `search/novelty-tally.md` (how the gap should be stated; Kim et al. 2026 is the closest design, rule 22 in verified-view has the sentence).
8. For products, regulation and benchmarks: the entries in `search/results-5-board-products.md`. For any other source: its entry in `search/results-1` to `results-6`.
9. Full texts in `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/Project Literature Review/references/` (pdftotext) when you need to confirm a wording.

## Rules

- Body length 1,000 to 1,500 words, target 1,330, counted without the title, section headings and the reference list. At least 15 sources; the structure gives 29.
- Five thematic subsections in the order of the structure; you may improve the titles. Each section answers, in prose and without labels: what exists, what we know, what is unresolved, how our project differs and extends. Synthesize across papers; never summarize sources one after another.
- In-text citations APA 7 (Author, Year; Author et al., Year), and a reference list in exactly the style of Abstract v8. Every in-text citation appears in the list and every list entry is cited. Take reference entries from Abstract v8, verified-view, the read files and the results files; do not invent details. Products and web sources are cited as in `results-5-board-products.md`.
- Plain English. No em-dashes. Short sentences. Numbers only where they change what the reader concludes (a sample size, an effect that sets a benchmark). No statistics jargon beyond means, sample sizes and "about".
- Author's decisions: politician studies (Chae & Tewksbury, Thorson, Schiff) are not cited as evidence; one sentence says results on known politicians are not carried over. Trust in the AI is not an outcome. Gupta et al. (2026) appears only in the gap section. Dobbs et al. (2026) is not cited. Kim et al. (2026) is the closest design; Chen et al. (2026) second.
- Corrected wordings: the burden-of-proof point is Kempt and Nagel's (2022), which Chen et al. take up; Madanay's undercalling radiologist was "rated lowest of the four conditions"; Squash showed 14 relevant and 20 irrelevant matches and the lab was building an interface for a human editor; Kim et al. (2026) bounds the AI-penalty claim (the advisor's error lowered trust, the AI's involvement changed nothing).
- Every factual claim must trace to verified-view, a read file, a results entry or a full text. If you are unsure of a claim, leave it out or put it in the writer's notes, never in the text.
- Do not edit any file other than the draft.

## Output file layout

1. Title line: `# Related work` (or a better title).
2. The five sections with `##` headings.
3. `## References`, APA 7, alphabetical, style of Abstract v8.
4. `## Writer's notes`: word count per section and total (body only); number of sources cited; any deviation from the structure and why; any claim you were unsure of and left out.
