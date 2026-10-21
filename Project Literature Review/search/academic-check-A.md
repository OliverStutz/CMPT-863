# Academic check A: Draft v6 against Draft v5 and the sources (2026-10-20)

Checker A. Inputs read: `Literature Review - Draft v5.md`, `Literature Review - Draft v6.md`, `search/_academic-brief.md`, `search/_academic-check-brief.md`, `search/verified-view.md` (Lowry, Opdenbusch entries), and the full texts in `references/` for SEC, Lowry, Opdenbusch, Lu (`pdftotext -layout`, page footers and headers inspected) and Sharma (.txt). Other checkers' files not read.

## 1. Verdict

**Accept with fixes.** v6 keeps every number, comparison and "how our study differs" statement from v5, the reference list is byte-identical to v5, and no rule fails. Two added topic sentences overreach ("Their accuracy under live conditions is limited", "Across domains"), one verb change weakens a design fact ("drafted" became "generated"), the director quotation's locator should point to the page where Lowry et al. identify the speaker as an expert interviewee (p. 798, not p. 783), and the Microsoft quotation has no locator.

## 2. Fidelity findings

Every v6 sentence was aligned with its v5 counterpart. Sentences not listed are "fine" (register change only, same claim, same numbers, same hedges).

| # | v6 sentence (short) | v5 sentence (short) | Verdict | Proposed fix |
|---|---|---|---|---|
| F1 | "Their accuracy under live conditions is limited." (new topic sentence after Ray-Ban Meta and Copilot) | none (v5 heading "Checkers are available, and often wrong") | **Meaning change.** "Their" refers to the two commercial products, but none of the evidence that follows measures Ray-Ban Meta or Copilot. Bendeck and Stasko (2025) is a test, not a live condition, so "under live conditions" also overstates. | "Evaluated checkers, however, have often been wrong." |
| F2 | "Our flags are generated with a real VLM and are all incorrect for the company presented." | "Our flags are drafted with a real vision-language model and are all wrong for the company." | **Nuance loss.** "Drafted" says the VLM output is a draft the researchers shape; "generated" reads as raw model output. | "Our flags are drafted with a real VLM and are all incorrect for the company presented." |
| F3 | "Across domains, people follow AI warnings even when the warnings are wrong." (new topic sentence) | heading "People follow a wrong flag" | **Nuance loss (overclaim).** Four studies in news, grading and lie detection do not license "across domains" in present tense as a general law. | "In several domains, people have followed AI warnings even when the warnings were wrong." |
| F4 | "A board member directs a content checker built on a vision-language model (VLM) at a slide ..." | "A board member points a camera assistant at the slide ..." | **Nuance loss.** The camera (point-and-ask, the link to Ray-Ban Meta) is lost. "Built on a VLM" is consistent with v5's later VLM sentence, so not a new claim. | "A board member points a camera-based checker built on a vision-language model (VLM) at a slide presented by the chief information security officer (CISO); the checker flags a purported fault, and the CISO responds." |
| F5 | "The costs fall on the person checked." (new topic sentence) | heading "... and the person flagged pays" | **Fine, borderline.** Same claim as the v5 heading. Unhedged, but the paragraph supports it; optional softening in section 4 (R4). | optional: "Costs can also fall on the person checked." |
| F6 | "Regulation assigns the oversight of an expert function to non-experts." | "Regulation asks non-experts to oversee an expert." | **Nuance loss (minor).** v5 says a person (the CISO) is overseen; "an expert function" moves the object from a person to a function, and the review's outcome is trust in a person. | "Regulation assigns non-experts the oversight of an expert." |
| F7 | "In practice, such oversight appears limited." (new) | none | Fine. Hedged ("appears") topic sentence, supported by Lowry and Opdenbusch in the next two sentences. | none |
| F8 | "Flags also prompt those with power over the person assessed to act." | "Flags also move people with power to act." | Fine. "Over the person assessed" narrows v5 slightly but matches both studies (teachers over students; accusers in von Schenk et al. impose a cost on senders). | none |
| F9 | Sharma: "term a rise beyond the initial level 'further development of trust' rather than repair" | "call a rise beyond the starting level 'further development of trust'" | Fine. "Rather than repair" is in the source ("should be considered not as 'repair' but as further development of trust"), so it adds no claim. | none |
| F10 | "because even a denial with an account was the weaker reply" | "since even a denial with an account lost" | Fine. Same comparison against apology. | none |
| F11 | "the checker is held by the non-expert who holds power over the expert" | "held by the non-expert with power over the expert" | Fine in meaning; register issue (held / holds), see R5. | see R5 |
| F12 | "The design thereby tests whether ... a reason recovers it." | "This shows whether ... a reason wins it back." | Fine. "Tests" is more accurate than "shows". | none |

Verified wordings required by `_academic-brief.md`: Madanay "rated that radiologist lowest of the four conditions" kept; Squash 14 relevant and 20 irrelevant kept; D. Kim et al. (2026) as the bound on the AI-penalty claim and as the closest design kept; Du's red highlighting not placed on the 10-point scale kept; Aamir "felt more confident" kept; Sharma "further development of trust" and 8 of 94 kept; Pezzo in one cautious sentence ("The nearest evidence ... is indirect") kept; Kempt and Nagel "an exchange of reasons cannot settle the disagreement" kept, with "may meet this burden" as our hypothesis; Schilke's power "derived from alternatives rather than from authority" kept.

Numbers: 63.5%, 14 and 20, 16th percentile, 23%, 0.30 on −1 to +1, 1.8 points on a 10-point scale, almost 85%, 8 of 94, four interviewees: all unchanged.

"How our study differs" statements, all survive: board evidence has no tool or trust measure, "the present study adds both"; no checker study reports trust in the speaker, Danry's participants rated recordings; flags drafted with a real VLM and wrong for the company (verb changed, F2); CISO never mentions AI (after Reif); flags concern the slide and show neither figure (after Loconte); studied persons stay silent and are rated once, "our CISO responds"; rise above first rating exploratory, bare denial expected to repair less; reason may meet the burden; the gap paragraph's three contrasts with D. Kim et al. and the within-subject design.

## 3. Page locators

- **SEC (2023) p. 85: confirmed.** Printed footer "85" on PDF page 85. Quotation verbatim: "often effectively oversee management's efforts without specific subject matter expertise".
- **Lowry et al. (2026) p. 783, "mostly symbolic": confirmed.** Abstract on printed p. 783 (journal header pp. 783–804, footer 783): "efforts lack substance and therefore are mostly symbolic, even when they perform the same oversight activities as expert directors". Verbatim.
- **Lowry et al. (2026) p. 783, "all they can do ...": verbatim, but p. 798 is the better locator.** On p. 783 the line is an epigraph attributed to "Former NASDAQ Board Director"; it carries no press citation (the adjacent "(Rundle 2023)" belongs to the 12% S&P 500 statistic in the right-hand column, not to the epigraph). On printed p. 798 Lowry et al. quote the same words from their own interview data: "an expert director strongly challenged the idea ... stating: [M]ost boards are simply completely incapable of overseeing cyber risk. ... all they can do is assess the credibility of the executives that are put in front of them. (D-9 expert)". So this is primary interview data, not a secondary press quotation, and the "expert" label that v6 uses appears only on p. 798. Under APA 7, quoting a participant's words as reported in a study is cited to the study with the page where the words appear, which makes "one expert director they quote" acceptable. More precise: "one expert director they interviewed" with "(p. 798)". Exact fix in section 6, fix 1.
- **Opdenbusch et al. (2025) p. 2, "the CISOs are the ones in power": confirmed.** Footer "2"; contribution (III): "We show that the CISOs are the ones in power". Verbatim.
- **Opdenbusch et al. (2025) p. 6, "too shy to ask questions": confirmed.** Footer "6"; "Some participants (CISO4; EX6; NED3,4) pointed out that NEDs lack a sufficient understanding ... Hence, they would be too shy to ask questions". Verbatim, and the count of four participants is right. Note: the words are the authors' paraphrase of four participants, not the participants' own words. "Four participants described directors as 'too shy to ask questions' (p. 6)" is acceptable because the quotation is cited to the authors. A stricter form is optional (R6). (The phrase recurs on p. 13 as "too shy to ask the right questions"; p. 6 is the correct match.)
- **Lu et al. (2025) p. 12: confirmed, with caveat.** The PDF prints no page numbers (CHI article, 19 pages, reference gives pp. 1–19). PDF page 12 reads "we found that people lack the capability to differentiate correct and incorrect AI predictions, regardless of ...". Verbatim. Page 12 of the article's own 1–19 numbering is the correct APA locator. (A near-variant on PDF page 16, "differentiate the correctness of AI predictions", is not the quoted text.)
- **Sharma et al. (2023), "What Are the Limits of Trust Repair?" section: confirmed.** The heading is line 69 of the .txt; the same section contains "should be considered not as 'repair' but as further development of trust". Verbatim. The text has no page numbers, so a section locator is the APA 7 form.
- **Microsoft (n.d.): locator missing.** v6 quotes "Where are some holes in [a meeting participant's] argument?" with no locator, and APA 7 asks for one on every direct quotation (a section heading or paragraph number for a web page). No local copy of the support page exists to confirm a heading. Two options are given in R7.

## 4. Register findings

- **R1 (overclaim, = F1).** Old: "Their accuracy under live conditions is limited." New: "Evaluated checkers, however, have often been wrong."
- **R2 (missing hedge, inconsistent tense, = F3).** Old: "Across domains, people follow AI warnings even when the warnings are wrong." New: "In several domains, people have followed AI warnings even when the warnings were wrong."
- **R3 (spelling consistency).** v6 uses British "judgement" (heading of section 3) and "judgements" (METI sentence), while the reference list and APA use American spelling ("Behavior"). Old: "Effects of erroneous flags on judgement and on the person flagged" New: "Effects of erroneous flags on judgment and on the person flagged". Old: "laypeople's judgements of an expert" New: "laypeople's judgments of an expert". (Keep British if the course prefers it, but then use it everywhere.)
- **R4 (unhedged topic sentence, optional, = F5).** Old: "The costs fall on the person checked." New: "Costs can also fall on the person checked."
- **R5 (repetition and vague antecedent in the gap paragraph).** Old: "The present study combines them: the checker is held by the non-expert who holds power over the expert," New: "The present study combines these elements: the checker is held by the non-expert who has power over the expert,"
- **R6 (attribution of a paraphrase, optional).** Old: "four participants described directors as "too shy to ask questions" (p. 6)" New: "the authors report, from four participants, that directors were "too shy to ask questions" (p. 6)". Optional, because the citation already names the authors.
- **R7 (APA: quotation without locator).** Option a (preferred, no lookup needed): Old: "suggests, in a private pane, the prompt "Where are some holes in [a meeting participant's] argument?" (Microsoft, n.d.)." New: "suggests, in a private pane, a prompt asking where the holes lie in a meeting participant's argument (Microsoft, n.d.)." Option b: keep the quotation and add the section heading of the live support page in the form "(Microsoft, n.d., "<Heading>" section)" once checked on the page; earlier reviews (review-A item 10, review-C item 12) already ask for that page to be rechecked before submission.
- **R8 (= F6).** Old: "Regulation assigns the oversight of an expert function to non-experts." New: "Regulation assigns non-experts the oversight of an expert."
- **Voice (no fix required).** v6 alternates "we"/"our" with "the present study". APA 7 permits both; the mix is acceptable because "the present study" is used for design statements and "we" for interpretation and hypotheses.
- **Tense: acceptable.** Findings are in the past tense; authors' arguments and conclusions (Opdenbusch "conclude", Lowry "trace", Sharma "term ... and note", Kempt and Nagel "argue") and theory (Mayer) are in the present tense; the scenario and the present study are in the present tense. No mixed tense within a sentence except R2.
- **Headings:** all five are noun phrases. **In-text citations:** D. Kim / P. H. Kim initials correct (same surname, different first authors); "Danry et al.'s (2020)" possessive form correct; "(Lu et al., 2025, p. 12)" correct.

## 5. Rules checklist

| Rule | Result |
|---|---|
| No em-dashes | Pass (0 in the file; no en-dashes in the body either; the minus sign in "−1 to +1" is U+2212) |
| Body 1,000 to 1,500 words | **Pass: 1,450 words** (opening plus five sections, title, headings and references excluded, split on whitespace). v5 was 1,281. The margin is 50 words, so fixes must stay close to word-neutral; the fixes in section 6 change the count by about +3. |
| One sentence: politician results not carried over | Pass ("Findings on well-known politicians, whose reputations are settled, are not carried over.") |
| Trust in the AI is not an outcome | Pass (only as Gupta et al.'s measure) |
| Gupta et al. (2026) only in the gap section | Pass (one mention, gap section) |
| Dobbs absent | Pass |
| No H1 to H4 labels | Pass |
| Every in-text citation has an entry and every entry is cited | Pass (all 30 entries cited at least once; no citation without an entry) |
| Reference list identical to v5 | Pass (diff empty) |
| CISO, VLM, METI spelled out at first use | Pass (CISO and VLM in the opening sentence, METI in the Hendriks sentence). SEC is never abbreviated, so no definition is needed. |

## 6. Ranked fixes (exact old and new text)

1. **Lowry locator and attribution.**
   Old: `and one expert director they quote observed that for most boards "all they can do is assess the credibility of the executives that are put in front of them" (p. 783).`
   New: `and one expert director they interviewed observed that for most boards "all they can do is assess the credibility of the executives that are put in front of them" (p. 798).`
2. **Unsupported topic sentence on product accuracy.**
   Old: `Their accuracy under live conditions is limited.`
   New: `Evaluated checkers, however, have often been wrong.`
3. **Flags drafted, not generated.**
   Old: `Our flags are generated with a real VLM and are all incorrect for the company presented.`
   New: `Our flags are drafted with a real VLM and are all incorrect for the company presented.`
4. **Overgeneralised topic sentence.**
   Old: `Across domains, people follow AI warnings even when the warnings are wrong.`
   New: `In several domains, people have followed AI warnings even when the warnings were wrong.`
5. **Microsoft quotation without locator (paraphrase option).**
   Old: `suggests, in a private pane, the prompt "Where are some holes in [a meeting participant's] argument?" (Microsoft, n.d.).`
   New: `suggests, in a private pane, a prompt asking where the holes lie in a meeting participant's argument (Microsoft, n.d.).`
6. **Camera restored in the opening.**
   Old: `A board member directs a content checker built on a vision-language model (VLM) at a slide presented by the chief information security officer (CISO);`
   New: `A board member points a camera-based checker built on a vision-language model (VLM) at a slide presented by the chief information security officer (CISO);`
7. **Gap paragraph wording.**
   Old: `The present study combines them: the checker is held by the non-expert who holds power over the expert,`
   New: `The present study combines these elements: the checker is held by the non-expert who has power over the expert,`
8. **Object of oversight.**
   Old: `Regulation assigns the oversight of an expert function to non-experts.`
   New: `Regulation assigns non-experts the oversight of an expert.`
9. **Spelling consistency.**
   Old: `on judgement and on the person flagged` New: `on judgment and on the person flagged`
   Old: `laypeople's judgements of an expert` New: `laypeople's judgments of an expert`
10. **Optional hedge.**
    Old: `The costs fall on the person checked.`
    New: `Costs can also fall on the person checked.`
