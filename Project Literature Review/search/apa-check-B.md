# APA 7 audit, reviewer B

File audited: `Literature Review - Final v13.md` (the brief named `Literature Review - Draft v13.md`, which does not exist in the folder; `Final v13` is the only v13 file). Line numbers below refer to that file. Style audit only: no source facts were checked, nothing was fetched, the draft was not modified.

## 1. Verdict

**Compliant with minor fixes.** Must fix: 0. Should fix: 3. Optional: 5.

The reference list is clean: all 31 entries have correct author inversion, "&" before the last author, sentence-case titles, italic journal names and volumes, plain issue numbers, en-dashes in page ranges, "Article" numbers where applicable, and a uniform `https://doi.org/` DOI style with no closing period. Ordering is correct. Every in-text citation has an entry and every entry is cited. The two Kim first authors carry initials in all six citations. The fixes concern an abbreviation used only once, a "NIS2" / "NIS 2" mismatch, and appendix labelling.

## 2. Table 1, reference list

| Entry | Problem | APA 7 rule | Corrected text |
|---|---|---|---|
| Aamir et al., 2026; Gupta et al., 2026; Lu et al., 2025; Danry et al., 2020 (ACM proceedings) | Optional. ACM proceedings papers carry an article number rather than true page numbers; "pp. 1–22" etc. reproduces the ACM pagination within the article. Acceptable as written. | 9.27 and 10.5 (use an article number when the work has one) | If the ACM Digital Library lists an article number, e.g. "In *Proceedings of the 2026 CHI Conference on Human Factors in Computing Systems* (Article 123). Association for Computing Machinery. https://doi.org/..." Otherwise keep as is. Do it for all four or none. |
| Microsoft, n.d. | Optional. Title reads "Microsoft Copilot in Teams" while the URL slug reads "microsoft-365-copilot". APA asks for the title as shown on the page. The retrieval date itself is correct: an undated support page that is updated over time. | 9.16 (retrieval dates), 10.16 (webpages) | Confirm the on-page title on the next visit; if it says "Microsoft 365 Copilot", write: *Catch up on meetings with Microsoft 365 Copilot in Teams*. |
| Directive (EU) 2022/2555, 2022 | No change required (see Judgement calls). Acceptable adaptation for an EU instrument. | Ch. 11 (11.1, 11.10); APA defers to Bluebook-style legal form, which is US-oriented | Keep. |
| U.S. Securities and Exchange Commission, 2023 | No change required (see Judgement calls). Correct as a government report with a report number. | 10.4 (reports by a government agency); 9.11 (group author); 9.29 (report numbers) | Keep. |

**Entries inspected with no problems found** (authors, date, title case, italics, volume and issue, page range or article number, DOI, punctuation): Adair (2020), Bendeck & Stasko (2025), Chen et al. (2026), Du et al. (2026), Ferrin et al. (2007), Greitemeyer & Sagioglou (2015), Hendriks et al. (2015), Horne & Nevo (2025), Kempt & Nagel (2022), D. Kim et al. (2026), P. H. Kim et al. (2004), Loconte et al. (2026), Lowry et al. (2026), Madanay et al. (2025), Mayer et al. (1995), Meta (2024), Opdenbusch et al. (2025), Pezzo et al. (2022), Prike et al. (2025), Reif et al. (2025), Schilke et al. (2015), Sharma et al. (2023), Strickland (1958), von Schenk et al. (2024). Together with the four ACM entries, Microsoft, the Directive and the SEC above, all 31 entries were inspected.

Notes on individual entries that were checked and pass:
- **D. Kim et al. (2026), AIES.** Formatted as a journal article (italic proceedings name, italic volume 9, issue 3). Correct, because AAAI publishes the AIES proceedings in journal form with volume and issue (10.5: proceedings published regularly in a journal format are cited as journal articles). The difference from the CHI entries is therefore justified.
- **Opdenbusch et al. (2025), NDSS.** No page range; NDSS papers are unpaginated, so omitting pages is correct. "A" after the question mark in the title is correctly capitalized.
- **Adair (2020).** Dated web page with author, full date, italic title, site name (differs from author) and URL. Correct (10.16). "State of the Union" and "AI" are correctly capitalized.
- **Meta (2024).** Group author is also the site owner, so the site name is correctly omitted (10.16, "when the author and site name are the same, omit the site name"). No retrieval date is correct for a dated news post.
- **SEC.** Release numbers in plain text in parentheses after the italic title; publisher omitted because it equals the author. Correct.
- **Madanay et al. (2025).** The en-dash in "artificial intelligence–clinician" is the journal's own title form, keep it.
- **von Schenk et al. (2024).** Hyphenated initials "J.-F." are correct (9.8).
- **PNAS.** Schilke (2015) has a page range and Reif (2025) an "Article e..." number; this matches each article's actual form, so it is consistent, not a conflict.

**Ordering (9.44 to 9.48).** The full list is in correct letter-by-letter order. Specific checks:
- "Directive (EU) 2022/2555" files under its first word, between Danry and Du (Da < Di < Du). Correct (legal works without an author file by title, 9.12).
- Madanay < Mayer < Meta < Microsoft < Opdenbusch. Correct.
- Kim, D. (2026) before Kim, P. H. (2004): when first authors share a surname, order by first initial, not by year (9.47). Correct.
- Strickland < U.S. Securities and Exchange Commission < von Schenk. Correct if "von" is alphabetized as part of the surname; see Judgement calls.
- Mechanics: every element ends in a period, no double spaces or trailing spaces, all page ranges use en-dashes (the only hyphens between digits are the SEC release numbers, which are correct), all 27 DOIs use `https://doi.org/` with no closing period.

## 3. Table 2, in-text citations

| Where | Problem | Rule | Corrected text |
|---|---|---|---|
| Introduction, "Regulators like the U.S. Securities and Exchange Commission (SEC, 2023) require..." | Should fix. The form of the first narrative citation with an abbreviation is correct (8.21, Table 8.1 shows "Organization Name (ABBR, Year)"), but "SEC" is never used again. APA asks to abbreviate only when the abbreviation is used at least three times. | 6.25, 8.21 | "Regulators like the U.S. Securities and Exchange Commission (2023) require public companies..." |
| Introduction, "In the European Union's NIS2 Directive it states..." | Should fix. The text says "NIS2" while the reference title says "(NIS 2 Directive)". Keep the abbreviation identical to the official form. | 6.25 (consistent abbreviations), 9.12 | "...the European Union's NIS 2 Directive..." |
| Introduction, same sentence, "(Directive [EU] 2022/2555, 2022, Art. 20)" | Optional. Square brackets for parentheses inside parentheses are the correct APA treatment (6.10), and "Art. 20" is an adequate locator (8.13 permits section or article locators; APA has no fixed abbreviation for an article of an EU act). The bracketed "[EU]" does, however, alter the official name. A narrative form avoids the nesting altogether. | 6.10, 8.13, 11.3 | Either keep as is, or: "Article 20 of the NIS 2 Directive, Directive (EU) 2022/2555 (2022), states that members of management bodies..." |
| "What remains unresolved", "Wrong-flag studies (see the Appendix) leave the person judged silent..." | Should fix. The paper has two appendices ("Appendix Memo and AI Transparency" and "Appendix: Wrong flags, and the person flagged"). With more than one appendix APA labels them Appendix A, Appendix B, each with the label on its own line above a bold title, and cross-references name the letter. | 2.14 | Headings: "Appendix A" / "Memo and AI Transparency" and "Appendix B" / "Wrong Flags and the Person Flagged". Text: "Wrong-flag studies (see Appendix B) leave..." |

**Citations inspected with no problems found:**
- **Parenthetical vs narrative, "&" vs "and".** All correct: "(Bendeck & Stasko, 2025)", "(Greitemeyer & Sagioglou, 2015)", "(Kempt & Nagel, 2022)", "(Horne & Nevo, 2025)" inside parentheses; "Kempt and Nagel (2022) argue" in running text.
- **"et al." from the first citation.** Applied to every work with three or more authors (Lowry, Mayer, Hendriks, Opdenbusch, Aamir, Danry, Kim, Ferrin, Sharma, Chen, Pezzo, Lu, Schilke, Gupta, Loconte, Reif, Horne, Du, von Schenk, Prike, Madanay). Both names given for every two-author work every time.
- **Kim initials (8.20).** All six citations carry initials: "(P. H. Kim et al., 2004)" three times (Trust repair; What remains unresolved; What this study adds), "(D. Kim et al., 2026)" in What remains unresolved, "D. Kim et al. (2026)" in The closest designs and in the appendix. P. H. Kim as second author of Ferrin et al. (2007) needs no initials, since the rule covers first authors only.
- **Multiple sources in one parenthesis.** All alphabetical and semicolon-separated: (Ferrin et al., 2007; P. H. Kim et al., 2004), (Danry et al., 2020; Lu et al., 2025), (Chen et al., 2026; D. Kim et al., 2026), (Greitemeyer & Sagioglou, 2015; Sharma et al., 2023). Initials do not affect order (8.12).
- **Repeated narrative citations in one paragraph (8.16).** "Lowry et al. (2026) interviewed..." followed later in the same paragraph by "Lowry et al. read the same finding..." correctly omits the year. "Sharma et al. (2023) count... They also note..." correct. Parenthetical repeats, e.g. "(Mayer et al., 1995)" twice in one paragraph, correctly keep the year.
- **Undated work.** "(Microsoft, n.d.)" correct.
- **Group authors.** "(Meta, 2024)", "(Microsoft, n.d.)" correct.
- **Lowercase prefix.** "(von Schenk et al., 2024)" mid-sentence, lowercase "von" correct; it would become "Von" only at the start of a sentence.
- **Possessives.** No possessive citation such as "Danry et al.'s" appears; nothing to check.
- **Direct quotations.** None. The only quotation marks surround "AI" (scare quotes) and section titles in the memo, so no page or paragraph locators are required.

## 4. Two-way match

- Entries never cited: **none.** All 31 entries are cited at least once; Horne, Du, von Schenk, Strickland, Prike and Madanay are cited only in the appendix, which counts, since APA treats appendix citations like body citations.
- Citations lacking an entry: **none.** Surname spelling and year match in every case (including Opdenbusch 2025, Greitemeyer & Sagioglou 2015, Pezzo 2022, Schilke 2015).

## 5. Judgement calls

1. **EU directive.** APA 7 Chapter 11 follows the Bluebook and gives templates only for US cases, statutes, regulations, constitutions and treaties (11.1 to 11.10); it has no EU template, and the APA Style team's guidance for foreign legal material is to follow that jurisdiction's own citation conventions or, failing that, to give the reader title, date, official source and a URL. The entry gives the full official title (including "(NIS 2 Directive)"), the year, the *Official Journal of the European Union* with series and issue "L 333" and pages 80–152, and the EUR-Lex ELI link. This is the form most university APA guides adopt for EU law. **Recommendation: keep.** The in-text short form "Directive (EU) 2022/2555, 2022" is the right short title (11.3). Only the "NIS2" spelling in the text needs aligning (Table 2).
2. **SEC release.** APA 11.7 handles US federal regulations in legal form, for a final rule usually via the Federal Register, e.g. "Cybersecurity Risk Management, Strategy, Governance, and Incident Disclosure, 88 Fed. Reg. 51896 (Aug. 4, 2023) (to be codified at 17 C.F.R. pts. 229, 232, 239, 240, 249)". That form is correct but heavy, and its in-text citation would be the long rule title. The paper instead cites the SEC's own release PDF as a government report (10.4), with group author, year, italic title, release numbers and URL, which is fully compliant for the document actually consulted. **Recommendation: keep the report form.** Switch to the Federal Register form only if the instructor asks for legal citation of regulations.
3. **"von Schenk" placement.** APA 7 asks writers to follow the author's own usage for surname prefixes. The authors publish as "von Schenk" and the paper cites them that way in text, so filing under V (last, after U.S.) is consistent. The alternative German convention would file it as "Schenk, A. von" between Reif and Schilke, but then the in-text citation would have to change too. **Recommendation: keep under V.**
4. **ACM page ranges vs article numbers.** "pp. 1–22" is not wrong, but ACM's own metadata uses article numbers. **Recommendation: leave as is unless the instructor is strict; if you change it, change all four ACM entries together.**
5. **SEC abbreviation.** Some instructors accept a defined abbreviation used once. Strict APA (6.25) says to drop it. **Recommendation: drop "SEC", the sentence loses nothing.**

Outside the APA scope, but noticed while reading the citation sentences: "legitimate(while" lacks a space, "refered" and "CISOS" are misspelled, and the memo uses lowercase "i".
