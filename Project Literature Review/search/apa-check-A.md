# APA 7 check A, Literature Review v13

File checked: `Literature Review - Final v13.md`. The brief named `Literature Review - Draft v13.md`, which does not exist; `Final v13` is the only v13 file in the folder (125 lines, 31 references). If a different file was meant, this check must be rerun.

## 1. Verdict

**Compliant with minor fixes.** Must fix: 0. Should fix: 2. Optional: 9.

## 2. Table A: reference list issues

| Entry | Issue | APA 7 rule | Exact fix |
|---|---|---|---|
| Aamir et al., 2026; Gupta et al., 2026; Lu et al., 2025 (CHI papers) | ACM CHI papers are paginated per article and carry an article number. "pp. 1–22" is not wrong, but APA prefers the article number where one exists. Editors of the proceedings are also left out (APA 10.5, Example 46 lists them), which is common practice and tolerated. | 9.29, 10.5 | Optional: replace "(pp. 1–22)" with "(Article NNN)" using the article number on the ACM DL page, e.g. "In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (Article NNN). Association for Computing Machinery." Same for Gupta and Lu. Keep pages if the number is not at hand. |
| Danry et al., 2020 | Same point as above (AHs '20 papers also carry article numbers). | 10.5 | Optional: "(Article N)" instead of "(pp. 1–12)". |
| Kim, D. et al., 2026 | Formatted as a periodical (italic container, volume 9, issue 3) while the other four proceedings are formatted as edited books. This is allowed: APA 10.5 says proceedings published regularly in a journal-like format are cited as periodicals, and AIES publishes with volume and issue. | 10.5 | No change needed. Optional only for visual consistency. |
| Opdenbusch et al., 2025 | No page range or article number. NDSS papers are not paginated, so omitting the locator is correct. | 10.5 | No change. |
| Adair, 2020 | Treated as a web page (italic title, site name). Reporters' Lab items are blog-style posts; APA 10.1 (Example 17) formats blog posts like articles, with the post title plain and the blog name italic. Either is defensible. | 10.1, 10.16 | Optional alternative: "Adair, B. (2020, February 23). Squash report card: Improvements during State of the Union … and how humans will make our AI smarter. *Duke Reporters' Lab*. https://reporterslab.org/..." Keeping the current form is acceptable. |
| Microsoft, n.d. | Format is correct (n.d. plus retrieval date for a page designed to change; author equals site, so site name omitted, 9.24). The title says "Microsoft Copilot" while the URL slug says "microsoft-365-copilot"; the page title may be "...with Microsoft 365 Copilot in Teams". Not verifiable offline. | 9.16, 9.24 | Optional: check the live page title and copy it exactly. |
| Reif et al., 2025; Schilke et al., 2015 | Journal given as "Proceedings of the National Academy of Sciences"; the full title is "...of the United States of America". APA says use the title as the journal presents it; the short form is widely accepted and is used consistently in both entries. | 9.25 | Optional: "*Proceedings of the National Academy of Sciences of the United States of America*" in both entries, or leave as is. |
| Directive (EU) 2022/2555, 2022 | See Notes. Acceptable author-date adaptation. | 11.1, 11.4 | No required change. |
| U.S. Securities and Exchange Commission, 2023 | See Notes. Acceptable as a government report with report numbers (10.4). | 10.4, 11.10 | No required change. |
| Whole list | Hanging indent (0.5 in) and double spacing cannot be shown in Markdown. | 9.43 | Optional: apply hanging indent when exporting to Word or PDF. |

**Alphabetical order (A.8): checked, no issue.** Directive sorts by its first word between Danry and Du ("Di" before "Du"). Meta before Microsoft ("Me" before "Mi"), both after Mayer. Kim, D. (2026) before Kim, P. H. (2004), correct because entries by different first authors with the same surname are ordered by initials (9.48), not by year. U.S. Securities and Exchange Commission sorts under U, before von Schenk. von Schenk under V is correct when "von" is treated as part of the surname, which matches the author's own usage and the in-text form (see Notes).

**Punctuation and DOI format (A.9): checked, no issue.** All page ranges use en dashes; no double spaces; every DOI is "https://doi.org/..." with no trailing period; "Article" numbers are used consistently for all article-numbered journals (Chen, Du, Greitemeyer, Hendriks, Horne, Loconte, Madanay, Reif, von Schenk); volume numbers are italic and issue numbers plain in every journal entry.

**Checked, no issue (author, date, title, source elements):** Aamir 2026 (apart from the optional article number), Bendeck & Stasko 2025, Chen 2026, Danry 2020 ("Wearable Reasoner" correctly kept capitalized as a system name), Du 2026, Ferrin 2007, Greitemeyer & Sagioglou 2015 (question-mark title, no added period, correct), Gupta 2026, Hendriks 2015, Horne & Nevo 2025, Kempt & Nagel 2022, Kim D. 2026, Kim P. H. 2004, Loconte 2026, Lowry 2026, Lu 2025, Madanay 2025 (en dash in "intelligence–clinician" follows the original title), Mayer 1995, Meta 2024 (author equals site, site omitted, correct), Opdenbusch 2025 (capital "A" after the quoted question, correct), Pezzo 2022, Prike 2025, Schilke 2015, Sharma 2023, Strickland 1958, von Schenk 2024. All use "Surname, I. I." form, ampersand before the last author, under 20 authors, sentence-case titles with a capital after colons, title-case italic journal names.

## 3. Table B: in-text citation issues

| Location | Issue | Rule | Exact fix |
|---|---|---|---|
| Introduction, "Regulators like the U.S. Securities and Exchange Commission (SEC, 2023) require..." | The form itself is APA's model for introducing a group-author abbreviation. But the abbreviation is never used again, and APA asks that an abbreviation be introduced only if used often (at least three times). | 6.25, 8.21 | Should fix: "Regulators like the U.S. Securities and Exchange Commission (2023) require..." |
| Introduction, "In the European Union's NIS2 Directive it states..." | The body calls it "NIS2 Directive"; the reference entry and the official title say "NIS 2 Directive". Same source, two names. | Consistency, 8.14 | Should fix: "In the European Union's NIS 2 Directive it states..." (or change the entry; the official title uses "NIS 2"). |
| Trust theory and its measurement, "Trust is the willingness..." and "The same willingness..." | Mayer et al., 1995 cited in two neighbouring sentences of the same paragraph with no other source between them. Not an error, but APA advises not repeating the same parenthetical citation when context makes the source clear. | 8.1 | Optional: drop the second "(Mayer et al., 1995)" if the sentence clearly continues the same source. |
| Introduction, "(Directive [EU] 2022/2555, 2022, Art. 20)" | Square brackets for the nested parentheses are correct (6.10). "Art." as a locator is acceptable for a legal provision. Shortened title in text is allowed for long legal titles. | 6.10, 8.13, 11.4 | No change. |

**Checked, no issue:**
- Kim citations: all five use initials: "(P. H. Kim et al., 2004)" in Trust repair and twice more (Unresolved, What this study adds); "D. Kim et al." in Unresolved, The closest designs, and Appendix. No bare "Kim et al." anywhere.
- "et al." from first citation for every work with three or more authors; two-author works (Bendeck & Stasko, Greitemeyer & Sagioglou, Horne & Nevo, Kempt & Nagel) always name both.
- "&" in parentheses, "and" in narrative ("Kempt and Nagel (2022)").
- Multiple citations ordered alphabetically and separated by semicolons: (Ferrin...; P. H. Kim...), (Danry...; Lu...), (Chen...; D. Kim...), (Greitemeyer & Sagioglou...; Sharma...).
- "Lowry et al. read the same finding..." omits the year correctly, as a repeat narrative citation within the paragraph after "Lowry et al. (2026)" (8.16).
- "(Microsoft, n.d.)" and "(Meta, 2024)" correct.
- No possessive narrative citations occur. No direct quotations from sources occur in body or appendix (the quoted "AI" in "What this study adds" is a scare quote, not a quotation), so no page or paragraph locators are needed.

## 4. Both-ways check

- Entries not cited: **none.** All 31 entries are cited. Six are cited only in the Appendix (Du 2026, Horne & Nevo 2025, Madanay 2025, Prike 2025, Strickland 1958, von Schenk 2024); this is correct, as appendix sources belong in the single reference list (2.14).
- Citations without an entry: **none.** Every surname and year matches.

## 5. Notes

1. **EU directive.** APA 7 Chapter 11 covers US legal materials and, for foreign and international law, refers writers to the conventions of that jurisdiction or the Bluebook. A strict Bluebook form would be "Directive 2022/2555, of the European Parliament and of the Council of 14 December 2022 on Measures for a High Common Level of Cybersecurity Across the Union..., 2022 O.J. (L 333) 80 (EU)". The draft instead uses the author-date adaptation that most university APA guides give for EU law: full title as the author position, year, *Official Journal of the European Union*, L series number, pages, URL. This matches the in-text form, and the reader can find the source. Recommendation: keep it.
2. **SEC release.** APA 11.10 would cite a final rule by its Federal Register publication, roughly "Cybersecurity Risk Management, Strategy, Governance, and Incident Disclosure, 88 Fed. Reg. 51896 (Aug. 4, 2023) (to be codified at 17 C.F.R. pts. 229, 232, 239, 240, 249)", with the in-text citation using the title. Verify volume and page before using this. The draft cites the SEC's own release PDF as an agency report with release numbers (10.4, report with number). That is a valid APA form for the document actually used. Recommendation: keep it.
3. **von Schenk.** APA 9.47 says to sort particles by the person's usage. Alicia von Schenk publishes as "von Schenk", so sorting under V with a lowercase "von" is defensible. If you would rather sort under S, write the entry "Schenk, A. von" and place it between Reif and Schilke ("Schenk" comes before "Schilke"). Recommendation: keep it under V. If a sentence ever starts with this citation, capitalize it: "Von Schenk et al. (2024)".
4. **Proceedings formats.** Mixing periodical form (AIES) and book form (CHI, AHs, NDSS) follows how each venue publishes, which APA 10.5 allows.
5. Outside the citation scope, seen while checking: "legitimate(while" in the Introduction needs a space before the parenthesis.
