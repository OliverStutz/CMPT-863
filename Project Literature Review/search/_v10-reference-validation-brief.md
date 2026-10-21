# Brief: validate the references of Draft v10 (2026-10-20)

One Opus agent. Output: `search/reference-validation-v10.md`. Edit no other file.

## What changed and why a check is needed

`Literature Review - Draft v10.md` is the full review (about 1,950 body words, 30 reference entries). Section 1 of the earlier drafts was checked claim by claim today (`search/reference-audit-v7.md`, `search/review-v8-A.md`, `search/review-v8-B.md`, `search/product-refs-check.md`). The opening and the other four themes (checkers, wrong flags, trust repair, closest designs) were rebuilt today by condensing `Literature Review - Draft v6.md`, whose text was verified claim by claim earlier (`search/check-tally.md`, `search/academic-tally.md`). v10 then regrouped every sentence under the four assignment questions, which separated some citations from the context they sat in. Three things can have gone wrong: a condensed sentence no longer says what v6 said; a citation now sits next to a claim it does not support; an in-text citation and the reference list no longer match.

## Tasks

1. **Citations both ways.** List every in-text citation in v10 (author-year, including the D. Kim and P. H. Kim forms and the n.d. entry). Confirm each has exactly one matching entry in the reference list and that every entry is cited at least once. Report any citation without an entry, any entry without a citation, any year or spelling mismatch, and any alphabetical-order error inside a multi-source citation.
2. **Claim fidelity for the rebuilt text.** For every sentence in v10 that is not in section 1 of `Literature Review - Draft v9.md` (that is, the opening and all sentences on checkers, wrong flags, trust repair and the closest designs), find its counterpart in v6 and check that the condensed wording keeps the same claim, number, qualifier and attribution. Where v10 differs from v6 in substance, check the full text in `references/` (pdftotext -layout; .txt for Madanay and Sharma; Chen has no PDF, use `references/reviews/chen.md` under `/mnt/RAID_STORAGE/SFU/Projects/CMPT-863/references/reviews/` and `search/verified-view.md`). Verdict per sentence: same claim, nuance lost (exact fix), or meaning changed (exact fix).
3. **Citation placement after regrouping.** For each citation in the "unresolved" and "differs" sections, confirm the sentence it is attached to is one the source supports in that form (for example "Danry et al.'s participants rated recorded statements, not a person" now stands apart from the Danry accuracy sentence; "Chen et al. (2026) is next: the doctor brought in the AI himself and never answered" stands alone). Flag any sentence whose claim now depends on context that moved elsewhere.
4. **Entry metadata.** The 30 entries are those of v6 with three fixes applied today (Madanay first author "F."; the Opdenbusch title without a colon after the question mark; the Microsoft entry retitled and re-dated). Confirm the three fixes are present and correct, spot-check five other entries against Crossref (`curl -s https://api.crossref.org/works/<DOI>`), and check that the Meta and Microsoft URLs resolve.
5. **The author's rules.** No em-dashes anywhere; Gupta et al. (2026) only in the closest-designs material; Dobbs absent; the one sentence that politician results are not carried over; trust in the AI not an outcome; no H1 to H4 labels; CISO, VLM, METI, SEC spelled out at first use in reading order (note that v10's order differs from v9's).

## Output layout

1. Summary with counts: citations (n, matched both ways, problems); rebuilt sentences (n, same claim / nuance lost / meaning changed); placement problems (n); entries (30, correct / fix needed); rules (pass or fail per item).
2. Citation table.
3. Fidelity table with exact fixes.
4. Placement findings with exact fixes.
5. Entry findings.
6. Ranked fixes, each with exact old and new text. No em-dashes in your proposals; no new claim or source.

Plain wording. Work from the files, not from memory. Say "could not verify" when a page is unreachable.
