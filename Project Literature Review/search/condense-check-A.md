# Condense check A: fidelity of Draft v4 against Draft v2

Checker A, 2026-10-19. Compared `../Literature Review - Draft v4.md` with `../Literature Review - Draft v2.md` sentence by sentence. Sources checked in full text with pdftotext where a wording changed: SEC (2023) quotation, Danry et al. (2020), Schilke et al. (2015), Gupta et al. (2026). Did not read condense-check-B.md.

## A. Verdict

Accept with fixes: no cited sentence changes meaning or adds a claim, every "must stay" wording is intact and the reference list is exact, but one uncited sentence ("This evidence comes from interviews") now follows Schilke et al.'s experiments and can be read as calling them interviews.

## B. Findings table

### B1. Cited sentences (33)

| # | v4 sentence (short) | v2 source sentence | Class | Note |
|---|---|---|---|---|
| 1 | SEC (2023): "broadly skilled directors" "often effectively oversee ..." | SEC full quotation | same | Quote shortened; remaining words verbatim (SEC PDF p. with "directors with broad-based skills ... often effectively oversee management's efforts without specific subject matter expertise"). "broadly skilled" is a fair paraphrase. "(SEC, 2023)" abbreviation dropped; not used elsewhere. |
| 2 | Lowry: "mostly symbolic"; "all they can do is assess the credibility ..." | same | same | "said that for" to "said for". |
| 3 | Opdenbusch: four interviewees "too shy to ask questions"; budget | same | same | |
| 4 | Aamir: "felt more confident challenging management" | same | same | Verified wording intact. |
| 5 | Opdenbusch: "the CISOs are the ones in power"; our reading | same plus Lowry counterpoint | narrowed | "The sources disagree on power" and Lowry's low-expertise explanation dropped; remaining claim unchanged. |
| 6 | Schilke: better alternatives trusted less; "though that power was having alternatives, not authority" | same plus "trust fell further during a negotiation" | narrowed | Required caveat kept. |
| 7 | Meta; Microsoft Copilot "in a private pane" | same | same | |
| 8 | Danry: audio glasses "judge whether a heard claim is backed by evidence", 63.5 percent live | "tell the wearer whether ..." | same | Full text: evidence-type classification of heard arguments; "judge" is accurate. |
| 9 | Squash 14 relevant, 20 irrelevant | same plus human-editor interface | narrowed | Verified figures intact. |
| 10 | Bendeck and Stasko: 16th percentile, some best-practice knowledge, 23 percent | same | same | |
| 11 | None of this work reports trust in the speaker; Danry et al.'s participants rated recorded statements | same | same | |
| 12 | Horne and Nevo 0.30 on -1 to +1 | same | same | |
| 13 | Lu et al.: larger effect than expert vote; quotation | same, minus "in the authors' words" | same | Quotation marks still attribute. |
| 14 | Du: 1.8 points lower on 10-point scale; red highlighting "lowered the grade further" | same plus "214 teachers" | narrowed | Extra drop still not claimed on the 10-point scale. |
| 15 | von Schenk: almost 85 percent | same | same | |
| 16 | Madanay: agreed less, "rated that radiologist lowest of the four conditions"; more-testing no penalty | same plus "1,606 adults", "When an AI and a radiologist disagreed" | narrowed | Disagreement remains implied by "than an AI". |
| 17 | Chen: lazier, lower trust | same | same | |
| 18 | D. Kim (2026) as limit: error, not AI's contradiction | same | same | Bound on AI-penalty claim intact. |
| 19 | Loconte: resisted confident accusations with accuracy and confidence shown | same | same | |
| 20 | Mayer et al.: definition; "here the budget decision" | same | same | "another party" to "another"; narrative to parenthetical citation. |
| 21 | METI (Hendriks): "applies these factors to laypeople judging an expert" | "Our three trust items ... built on the same factors" | narrowed | "three" dropped; "applies" equals "built on". |
| 22 | P. H. Kim 2004: apology beat a denial with an account after hearsay accusation of incompetence | same | same | |
| 23 | Ferrin: no better than the weaker of the two | same | same | |
| 24 | Greitemeyer: exoneration restored to never-accused level, not reliably above | "Clearing by a third party" | same | |
| 25 | Sharma: "further development of trust"; most studies fall short | same plus "complete repair" and denial competence/integrity point | narrowed | Verified wordings intact. |
| 26 | Kempt and Nagel (with Chen): burden of proof moves to expert | same | same | |
| 27 | Kempt and Nagel: "Reasons cannot settle the disagreement"; second human | same | same | |
| 28 | Pezzo: "nearest evidence on an expert's stated reason is indirect"; skipped aid, rated higher for believing they knew, neither below no-aid | "on an expert's reason about a machine"; "had not consulted"; "he or she" | narrowed | "about a machine" dropped from the lead-in; the sentence itself still says diagnostic aid. One cautious sentence. |
| 29 | Hypothesis: reason "may meet the burden in the listener's eyes" | "Our own hypothesis" | same | Marked as ours. |
| 30 | D. Kim closest design, METI | same plus "(N = 157)" | narrowed | |
| 31 | Their AI always right, advisor never replied, student held no tool or power | same | same | |
| 32 | Gupta: "measured belief in claims and trust in the AI, never in the speaker" | "... and never rated the person speaking" | narrowed | v4 says trust in the speaker was never measured, weaker than v2's "never rated the person"; full text confirms only trust in AI and its information was measured. |
| 33 | Sharma: 8 of 94 within-subject | same | same | |

Totals: same 23; narrowed without loss 10; changed meaning 0; new claim 0.

### B2. Uncited sentences with a fidelity issue

| v4 | v2 | Issue |
|---|---|---|
| "This evidence comes from interviews, with no tool or trust measure; we add both." | "The board evidence is from interviews, with no tool and no trust measure; ..." | Now directly follows Schilke et al. (2015), which is laboratory experiments with a trust measure. "This evidence" can be read as including Schilke, which would be false. v2's "The board evidence" ruled that out. |
| "Results on well-known politicians, whose reputations are settled, are not carried over." | same | Not a fidelity error (no citation in v2 either), but the sentence now has no politician study anywhere in the text to refer to; v2 at least had Prike et al.'s unfamiliar candidate nearby. Readability, not fidelity; left to Checker B. |

Other uncited changes (dropped "Such checkers are often wrong", "Wrong flags are believed", "The person flagged pays", "and are all wrong for the company", "and is rated after each step") lose no source claim; the gap section still states the flags are wrong for the company and rated after each flag and answer.

### B3. "What must stay"

| Wording | Status |
|---|---|
| Kempt and Nagel for the burden of proof | intact |
| Madanay "rated that radiologist lowest of the four conditions" | intact |
| Squash 14 relevant and 20 irrelevant | intact |
| D. Kim et al. (2026) as bound and as closest design | intact (both) |
| Du extra drop not claimed on the 10-point scale | intact ("lowered the grade further") |
| Aamir "felt more confident" | intact |
| Sharma "further development of trust" and 8 of 94 | intact (both) |
| Pezzo in one cautious sentence | intact (one sentence, "indirect" lead-in) |
| Kempt and Nagel "Reasons cannot settle it"; "meet the burden" as our hypothesis | intact |
| Author's rules checked here: no em-dash (grep: 0), CISO and METI spelled out at first use, no H1 to H4, Gupta only in the gap, Dobbs absent | all met |

### B4. Reference list

27 in-text sources, 27 entries; every entry cited and every citation listed. The 27 remaining entries are byte-identical to v2 (diff of the reference blocks shows only the four removed entries: Prike, Proudfoot, Reif, Strickland). Cosmetic: an extra blank line after `## References` in v4.

### B5. Dropped sources

| Source | v2 job | Claim in v4 now unsupported? | Text still needs it? |
|---|---|---|---|
| Proudfoot et al. (2023) | CISO competence judged by communication | No | No; the credibility point is carried by the Lowry et al. quotation. |
| Prike et al. (2025) | Human fact-check lowered unfamiliar candidate's ratings | No | No; the politician sentence never cited it. |
| Strickland (1958) | Checks that find nothing cost trust | No | No; the claim was removed with the source. |
| Reif et al. (2025) | AI-use penalty, "so our CISO never mentions AI" | No | No; the design clause was removed with it, and no sentence mentions the CISO using AI. |

## C. Ranked fix list

1. Referent of the interview sentence (+1 word).
   - Old: "This evidence comes from interviews, with no tool or trust measure; we add both."
   - New: "The board evidence comes from interviews, with no tool or trust measure; we add both."
2. Optional, cosmetic (0 words): delete the extra blank line after `## References`.
