# Brief: reference relevance scan (agent R) and reference correctness audit (agent C), 2026-10-20

Two independent Opus agents. Agent R writes `search/reference-scan-v7.md`; agent C writes `search/reference-audit-v7.md`. Neither edits any other file, and neither reads the other's output.

## The text under review

`Literature Review - Draft v7.md`, section "Cybersecurity oversight by non-expert boards" (two paragraphs, about 760 words), rewritten today. The rest of the review will follow from `Literature Review - Draft v6.md`, which is the verified baseline (claim-by-claim checks in `search/check-tally.md`, `search/academic-tally.md`). The study it situates is in `../CMPT 863 Project Abstract v8.md`.

The argument of section 1, claim by claim:
1. Boards must oversee cybersecurity professionals (employed or contracted) and mostly lack the expertise to judge their work.
2. The SEC (2023) made board oversight of cyber risk mandatory for listed companies but dropped the proposed disclosure of directors' cyber expertise, reasoning that broadly skilled directors can oversee without it (p. 85).
3. Field evidence: non-expert directors' oversight is "mostly symbolic"; they believe they ask the right questions but cannot interpret the answers; they can only judge the credibility of the executive in front of them (Lowry et al., 2026, pp. 783, 798).
4. Directors stay silent for fear of exposure; the budget is the only cyber decision most boards make (Opdenbusch et al., 2025, p. 6).
5. Routine external assurance makes directors more confident in challenging management; without literacy, boards accept management's account (Aamir et al., 2026).
6. Such a board has an evident need for a check that does not depend on its own expertise, and AI tools put one within reach of almost every director (Meta, 2024; Microsoft, n.d.).
7. The study's question is what a checker in one person's hand does to that person's trust in the person checked, not whether it improves questioning.
8. The board holds formal power (oversight, budget); the CISO holds informational power: CISOs decide what reaches the board (Opdenbusch et al., 2025, p. 2) and survey respondents estimated that a median 40% of CISOs filter their reports to look better, a risk most non-expert directors considered unlikely (Lowry et al., 2026, p. 796).
9. A non-expert board that does not expect filtering has no view beyond what a checker gives it, so the checker's verdicts are likely to carry weight whether or not they are right.
10. Power lowers trust: the party with better alternatives trusted its partner less (Schilke et al., 2015), where power meant alternatives rather than authority.
11. The consequences must be known before adoption: a wrong flag may cost a correct CISO the board's trust and with it the budget; a well-answered flag may strengthen it.
12. The evidence is interview-based, with no tool and no trust measure; trust studies are psychological and purely interpersonal (P. H. Kim et al., 2004; Ferrin et al., 2007); HCI checker studies measure belief in claims or trust in the AI, not trust in the person checked (Danry et al., 2020; Lu et al., 2025); the augmented human, a person checking another person with a tool in hand, has not been examined directly.
13. The present study examines exactly this scenario (design sentences from Abstract v8).

## Source pool (everything "we looked at")

- `search/_master-index.md` and `_master-index.csv`: the 129 sources from the six search agents, with rubric scores.
- `search/results-1-hci-checkers.md` to `results-6-factchecking.md`: the search entries (design, finding, relation, access, score).
- `search/results-7-regulation-and-board-ai.md`: today's regulation and board-AI research (23 verified entries).
- `search/_known-papers.md`: papers known before the search (abstract phase).
- `search/verified-view.md`: the 25 papers read or listed in the reading pass, with Oliver's reading, what the full text says, consequences, and 22 design rules.
- `search/read-pezzo-schilke-sharma.md`, `search/read-danry-kempt-dasilva.md`: full-text reads of six sources.
- `search/trust-topic-map.md`: 27 adjacent topics with sources.
- `Literature Review - Candidate papers.md`: the selection tally (25 chosen by all three reviewers, 6 by two, 3 by one) and the hypothesis warnings.
- `search/_review-design-v2.md`: the validated structure, with the sources planned per section and the optional ones (Proudfoot 2023, Da Silva & Jensen 2022, NACD & ISA, Hancock et al. 2020, Shaw et al. 2003, Wegner et al. 1981, Mickelberg 2025 were planned or optional and later cut for length).
- `references/`: full texts (pdftotext) and `references/README.md`.
- Author's fixed exclusions: politician studies (Chae & Tewksbury 2026, Thorson 2016, Schiff et al. 2025) are not evidence; Dobbs et al. 2026 declined; Gupta et al. 2026 only in the gap section; trust in the AI is not an outcome.

## Agent R: which other references apply to the current section 1

For each of the 13 claims, search the pool for sources not already cited in section 1 that would support, sharpen or contradict the claim, including sources the structure cut for length. Candidates to look at first: Proudfoot et al. (2023, MISQE; "always on audition", boards judge CISO competence from communication), Da Silva & Jensen (2022, CSCW; CISO as interpreter of a "dark art"), NACD & ISA (2026 handbook, "challenge management's technical assumptions"), NIS2 Art. 20 and the UK Cyber Governance Code (results-7), Nolen (2026) and PwC (2025) director surveys (results-7), Hendriks et al. (2016) on a flaw pointed out by another person, Reis, Reis & Kunde (2025) and Schilke & Reimann (2025) on the AI disclosure penalty, Wegner et al. (1981) and Letourneau & Gawronski (2024) on innuendo, Hancock et al. (2020) on AI-mediated communication, Kim et al. (2026) and Chen et al. (2026) for claim 12, Gale et al. (2022) and other board-governance papers in results-5, and anything in the master index scoring 15 or more on the rubric that touches boards, CISOs, power, filtering, or human-tool-human trust.

For each candidate: verify the entry (Crossref `https://api.crossref.org/works/<DOI>` or the publisher page; full text in `references/` if present), state in plain words what it found, which claim and sentence it would attach to, whether it is better than or additional to the current citation, and a recommendation: add, replace, or leave out, with the reason. Rank the recommendations. At most 15 candidates, with complete APA 7 entries in the style of Abstract v8 for every "add" or "replace". Also list, in one short section, any current section 1 citation that a better source in the pool would replace.

Output layout: 1. Summary (five lines: what section 1 is missing, the top three additions). 2. Ranked table of candidates (source, claim number, what it says, verdict, reason). 3. APA entries for the recommended ones. 4. Sources considered and left out, one line each.

## Agent C: how many of the current references are correct

Part 1, section 1 of v7, claim level. For every in-text citation in section 1 (SEC 2023; Lowry 2026, three locators; Opdenbusch 2025, two locators; Aamir 2026; Meta 2024; Microsoft n.d.; Schilke 2015; P. H. Kim 2004; Ferrin 2007; Danry 2020; Lu 2025): open the full text in `references/` (pdftotext -layout; .txt for Madanay and Sharma; Meta and Microsoft are web pages, fetch them) and check that the sentence states what the source says, that every quotation is verbatim, that every page or section locator is right (look at the printed page footer, not the PDF page index), and that the attribution is right (who said it: authors, an interviewee, a survey respondent). Verdict per claim: correct, imprecise (with the exact fix), or wrong (with the exact fix).

Part 2, reference list, metadata level. For all 30 entries in the v6 reference list (`Literature Review - Draft v6.md`, which v7 will inherit) plus the section 1 list in v7: check authors and initials, year, title, venue, volume, issue, pages or article number, and that the DOI resolves to that item, against Crossref (`curl -s https://api.crossref.org/works/<DOI>`) or the publisher page; for web sources check the URL resolves and the retrieval date form. Check APA 7 formatting details (italics, sentence-case titles, "Article" form, "In Proceedings ..." form, "n.d." with retrieval date, group authors). Verdict per entry: correct, imprecise (exact fix), wrong (exact fix).

Output layout: 1. Summary with counts (claims: correct / imprecise / wrong; entries: correct / imprecise / wrong). 2. Claim table for section 1. 3. Entry table for the reference list. 4. The exact fixes to apply, ranked.

## Rules for both

Plain wording, no em-dashes anywhere, nothing from memory alone: every statement about a source rests on a fetched page, the full text or a results entry that names its source. Say "could not verify" when a page is unreachable (eur-lex.europa.eu, pwc.com and apastyle.apa.org blocked automated fetches today; sec.gov may too).
