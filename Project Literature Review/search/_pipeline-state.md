# Pipeline state: literature review drafting (keep current; update after every phase)

Last updated 2026-10-20, phase 7 (academic register, Draft v6) done; Oliver's read of v6 pending. If the session is interrupted, resume from the first phase marked "running" or "not started". Every phase's inputs and outputs are files in this folder, so nothing depends on chat history.

## Phases

| Phase | What | Inputs | Outputs | Status |
|---|---|---|---|---|
| 0 | Structure validated by three Opus reviewers | `_review-design-v1.md`, `_design-brief.md` | `design-review-A/B/C.md`, `design-tally.md`, `_review-design-v2.md` | done 2026-10-19 |
| 0b | Four factual corrections verified and swept | `fact-check-2026-10-19.md` | corrections in v8, verified-view, results, reduced-view, `_review-design-v2.md` | done 2026-10-19 |
| 1 | Fetch and read the "read first" sources (two Opus agents) | `_review-design-v2.md`, access routes in results files | `read-pezzo-schilke-sharma.md`, `read-danry-kempt-dasilva.md`, PDFs copied into `../references/` | done 2026-10-19. Pezzo 2022 abstract only (all routes blocked; needs SFU) and it cuts against H2; Sharma says "further development of trust", not "over-repair", and only 8 of 94 repair studies were within-subject; Schilke's power is alternatives not authority; Kempt & Nagel support the burden shift but not "a reason meets it"; Danry rated no speaker, device 63.5 percent right |
| 2 | Draft the review (one Opus writer) | `_writer-brief.md` and everything it lists; phase 1 outputs | `../Literature Review - Draft v1.md` | done 2026-10-19: 1,472 body words, 34 sources; deviations listed in the file's Writer's notes (Pezzo cautious clause only; Sharma wording; Kempt as our hypothesis; Jongsma & Sand added; Google Gemini, Da Silva, Hancock, CharXiv left out; Kim papers with initials) |
| 3 | Check the draft (three independent Opus checkers) | `_check-brief.md`, the draft | `check-A.md`, `check-B.md`, `check-C.md` | done 2026-10-19: all three "minor fixes", 0 wrong, 0 unsupported, 8 overstated claims, 1,472 words |
| 4 | Tally the checks and apply fixes | the three check files | `check-tally.md`, `../Literature Review - Draft v2.md` | done 2026-10-19: 26 fixes applied, 6 not applied (reasons in the file's Editor's notes), body 1,448 words, 31 references (Apple, Jongsma & Sand, Mickelberg cut for length) |
| 5 | Final read by the main session (Claude) and by Oliver | Draft v2 | `../Literature Review - Final.md` after Oliver's approval | Claude's read done 2026-10-19: body 1,448 words, 0 em-dashes, 31 citations all listed and all cited, no H-labels, Copilot "private pane" and Lu 2025 DOI verified against sources. Open for Oliver: approve; decide on Mickelberg (in Abstract v8, cut from the review), Pezzo (abstract only, needs SFU full text), and the Du wording ("lowered the grade further" states the main effect of red highlighting). Final file = v2 without the Editor's notes |

## How to resume each phase

- **Phase 1 incomplete:** check which of the two read files exists. Re-run the missing one with the prompt pattern: fetch each PDF into a fresh scratch dir, verify with pdfinfo/pdftotext, copy to `../references/` with the Author-Year-title-Venue name, write a verified-view-style entry (Agent reading / What the full text says / Consequence / Open, plus the checked APA reference). Papers: Pezzo et al. 2022 (USF repository), Schilke et al. 2015 (PMC 4620861), Sharma et al. 2023 (Sage, SFU if blocked), Danry et al. 2020 (ACM open), Kempt & Nagel 2022 (BMJ open), Da Silva & Jensen 2022 (already in references, arXiv).
- **Phase 2:** give one Opus agent `_writer-brief.md`; it writes `../Literature Review - Draft v1.md` with References and Writer's notes.
- **Phase 3:** give three Opus agents `_check-brief.md`, outputs `check-A.md`, `check-B.md`, `check-C.md`; they must not read each other's files.
- **Phase 4:** tally into `check-tally.md` (claim verdicts, compliance, ranked fixes, votes), then have the writer (or a fresh Opus agent with the tally and the draft) produce `../Literature Review - Draft v2.md`.
- **Phase 5:** the main session reads v2 in full: word count of the body, every in-text citation in the list and back, em-dash check, the author's rules, spot-check of claims against verified-view. Then Oliver reads.

## Fixed decisions the drafts must respect (from `verified-view.md` rules and the author)

Five sections in the order of `_review-design-v2.md`; 1,000 to 1,500 words, target 1,330; at least 15 sources; APA as in Abstract v8; no em-dashes; plain English; politician studies not evidence (one sentence saying so); trust in the AI not an outcome; Gupta 2026 gap-only; Dobbs not cited; Kim et al. 2026 closest design and a bound on the AI-penalty claim; corrected wordings in `fact-check-2026-10-19.md`.

## Where the other state lives

- Reading pass and 22 design rules: `verified-view.md`.
- Novelty verdict: `novelty-tally.md`.
- Abstract: `../../CMPT 863 Project Abstract v8.md` with `... v8 - changes from v7.md`.
- Git: nothing from 2026-10-19 is committed yet (about 25 files). Commit when Oliver asks.

## Phase 6 (added 2026-10-19 at Oliver's request): condense

| Step | What | Inputs | Outputs | Status |
|---|---|---|---|---|
| 6a | Condensing pass 1 (Opus) | `_condense-brief.md`, Draft v2 | `../Literature Review - Draft v3.md` (1,100 to 1,200 words) | done 2026-10-19: 1,198 words, 27 sources; dropped Proudfoot, Prike, Strickland, Reif (three are in Abstract v8; checkers to rule on Strickland) |
| 6b | Condensing pass 2 (different Opus agent) | brief, Draft v3 | `../Literature Review - Draft v4.md` (1,000 to 1,150 words if lossless) | done 2026-10-19: stopped at 1,187 words (two heading-repeating sentences removed); the condenser judged every further cut lossy |
| 6c | Two independent checkers: fidelity of v4 to v2 and sources; instructions and readability | `_condense-check-brief.md`, v2, v4 | `condense-check-A.md`, `condense-check-B.md` | done 2026-10-19: both "accept with fixes"; A: 0 meaning changes, 1 wording fix; B: all Instructions.md points met, 5 losses a reviewer would miss, restore Prike, Strickland (and Reif for the instructor) |
| 6d | Editor applies agreed fixes | the two checks | `../Literature Review - Draft v5.md` | done 2026-10-19: 1,274 words, 30 references; Prike, Strickland, Reif, the Sharma competence point, "by a third party" and the power-disagreement clause restored; Mickelberg left out |
| 6e | Claude's read, then Oliver's | v5 | `../Literature Review - Final.md` | Claude's read done 2026-10-19: restored the "all wrong for the company" clause in section 2 (body 1,282); 0 em-dashes; 30 citations matched both ways; no H labels. Oliver's read pending; Final = v5 without the notes sections |

## Phase 7 (added 2026-10-20 at Oliver's request): academic register

| Step | What | Inputs | Outputs | Status |
|---|---|---|---|---|
| 7a | Rewrite v5 in academic register (main session) | `_academic-brief.md`, Draft v5, full texts for page locators | `../Literature Review - Draft v6.md` | done 2026-10-20: noun-phrase headings, topic sentences, formal wording, APA page or section locators on every quotation, % for percent; first pass 1,581 words, tightened to 1,450 |
| 7b | Three independent Opus checkers: fidelity to v5, locators against the full texts, register, rules | `_academic-check-brief.md`, v5, v6, `references/` | `academic-check-A/B/C.md` | done 2026-10-20: all three "accept with fixes", 0 rule failures, 1 meaning change (a new topic sentence), 2 to 4 nuance losses, locator fix (Lowry p. 798, not p. 783: interviewee D-9 expert; p. 783 is the epigraph) |
| 7c | Tally and apply | the three checks | `academic-tally.md`, v6 edited in place, entry in `../Edit Log.md` | done 2026-10-20: 14 fixes applied; Copilot quotation given its section locator from the live page; body 1,463 words, 0 em-dashes, 30 citations both ways, reference list identical to v5 |
| 7d | Oliver's read | v6 | `../Literature Review - Final.md` | pending. Final = v6 as it stands (no notes section in the file) |
