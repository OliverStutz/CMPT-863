# Academic-register check of Draft v6, checker B (2026-10-20)

Inputs read: Draft v5, Draft v6, `search/_academic-brief.md`, `search/_academic-check-brief.md`, `search/verified-view.md` (Lowry entry), `search/fact-check-2026-10-19.md`. Full texts checked with `pdftotext -layout`: SEC 2023, Lowry et al. 2026, Opdenbusch et al. 2025, Lu et al. 2025; Sharma et al. 2023 read from the .txt file. Other checkers' files were not opened.

## 1. Verdict

**Accept with fixes.** v6 keeps every claim, number, source and "how our study differs" statement of v5, the reference list is byte-identical to v5, and all rules pass. Three of the new topic sentences overreach slightly (one is a meaning change), two wording swaps lose nuance, and the Lowry and Opdenbusch quotations need a small attribution or locator adjustment.

## 2. Fidelity findings

All "how our study differs" statements survive: no trust measure or tool in the board evidence ("the present study adds both"); Danry rated recordings, not a person; flags are wrong for the company; CISO never mentions AI; flags concern the slide, show neither figure; politicians not carried over; silent and rated once, "our CISO responds"; Kim's AI always correct, no reply, no tool or power; Gupta never measured trust in the speaker; within-subject ratings before, after violation, after reply; neutral group. Verified wordings (Madanay "lowest of the four conditions", Squash 14 and 20, Kim as bound and closest design, Du not on the 10-point scale for red highlighting, Aamir "felt more confident", Sharma "further development of trust" and 8 of 94, Pezzo one sentence, Kempt and Nagel "cannot settle", "may meet this burden" as our hypothesis, Schilke alternatives not authority) are all intact.

| v6 sentence (short) | v5 sentence (short) | Verdict | Proposed fix |
|---|---|---|---|
| "A board member directs a content checker built on a vision-language model (VLM) at a slide ... the checker flags a purported fault" | "A board member points a camera assistant at the slide ... the tool flags a fault" | Nuance loss: the camera (the tool looks at the slide) is dropped and the tool is redefined as a "content checker" | "A board member points a camera-based checker built on a vision-language model (VLM) at a slide presented by the chief information security officer (CISO); the checker flags a purported fault, and the CISO responds." |
| "Their accuracy under live conditions is limited." | (no counterpart; v5 heading "Checkers are available, and often wrong") | Meaning change: "Their" refers to the Ray-Ban Meta glasses and Copilot, whose accuracy no cited source measures; the GPT-4V test was not live | "Evaluations of related checkers report limited accuracy." |
| "Our flags are generated with a real VLM" | "Our flags are drafted with a real vision-language model" | Nuance loss: "drafted" says the model produced a draft; "generated" implies raw model output | "Our flags are drafted with a real VLM and are all incorrect for the company presented." |
| "Flags also prompt those with power over the person assessed to act." | "Flags also move people with power to act." | Nuance loss: v5 meant power to act; v6 asserts power over the person assessed, which fits the teachers (Du) but not the participants who accused in von Schenk et al. | "Flags also prompt those in a position to act on them to do so." |
| "Across domains, people follow AI warnings even when the warnings are wrong." | (no counterpart; v5 heading "People follow a wrong flag") | Fine in substance, but an overclaim in register (see section 4) | See register item R1 |
| "The costs fall on the person checked." | (no counterpart; v5 heading "the person flagged pays") | Fine in substance; categorical where the next paragraph bounds it (see R2) | See R2 |
| "In practice, such oversight appears limited." | (no counterpart) | Fine: hedged, and supported by Lowry and Opdenbusch in the next sentences | None |
| "Power also bears on trust: ... although that power derived from alternatives rather than from authority." | "that power came from having alternatives, not from authority" | Fine | None |
| "Sharma et al. (2023) term a rise beyond the initial level 'further development of trust' rather than repair" | "call a rise beyond the starting level 'further development of trust'" | Fine: "rather than repair" matches the source ("should be considered not as 'repair' but as further development of trust") | None |
| "because even a denial with an account was the weaker reply" | "since even a denial with an account lost" | Fine | None |
| "their data attribute the loss to the error rather than to the AI's contradiction" | "their data tie the loss to the error, not to the AI's contradiction" | Fine | None |
| "The design thereby tests whether ... a reason recovers it." | "This shows whether ... a reason wins it back." | Fine | None |
| All other sentences | | Fine: same claim, same number, same hedge | None |

## 3. Page locators

- SEC (2023) p. 85: **confirmed.** Printed footer "85"; text "directors with broad-based skills in risk management and strategy often effectively oversee management's efforts without specific subject matter expertise, as they do with other sophisticated technical matters". Quotation verbatim.
- Lowry et al. (2026) p. 783, "mostly symbolic": **confirmed.** Abstract on the first page, printed footer "783": "efforts lack substance and therefore are mostly symbolic". Verbatim.
- Lowry et al. (2026) p. 783, "all they can do is assess the credibility of the executives that are put in front of them": **verbatim on p. 783, but better cited to p. 798.** On p. 783 it is an epigraph signed "Former NASDAQ Board Director", with no press citation attached (the "(Rundle 2023)" beside it belongs to the S&P 500 sentence in the other column). On p. 798 (header "798") the same words appear in the findings, signed "(D-9 expert)", that is, one of the authors' own interviewees, coded as a director with cybersecurity expertise (Table 1, D-9 "Yes"). So the brief's premise that the words come from a press source is not borne out by the full text: they are an interview participant's words. Under APA 7, a participant quoted in a cited study is cited to that study with the page, so "one expert director they quote" is acceptable; "they interviewed" is more precise. The label "expert" is only supported on p. 798, so the locator should be p. 798. Proposed: "and one expert director they interviewed observed that for most boards "all they can do is assess the credibility of the executives that are put in front of them" (p. 798)."
- Opdenbusch et al. (2025) p. 2, "the CISOs are the ones in power": **confirmed.** Printed footer "2": "(III) We show that the CISOs are the ones in power". Verbatim.
- Opdenbusch et al. (2025) p. 6, "too shy to ask questions": **page and wording confirmed, attribution imprecise** (also in v5). Printed footer "6": "Some participants (CISO4; EX6; NED3,4) pointed out that NEDs lack a sufficient understanding of cybersecurity topics. Hence, they would be too shy to ask questions". Four participants is correct, but the quoted words are the authors' summary, not the participants' own words, so "four participants described directors as 'too shy to ask questions'" puts the authors' words in the participants' mouths. Proposed in section 6.
- Lu et al. (2025) p. 12: **confirmed.** The CHI PDF carries no printed page numbers, but the article runs pp. 1 to 19 as in the reference entry, and the quotation is on article page 12: "we found that people lack the capability to differentiate correct and incorrect AI predictions, regardless of the level of performance of the AI". Verbatim. (A near paraphrase on p. 16 reads "differentiate the correctness of AI predictions", so p. 12 is the right one.)
- Sharma et al. (2023), "What Are the Limits of Trust Repair?" section: **confirmed.** The heading exists verbatim (line 69 of the .txt), and "should be considered not as "repair" but as further development of trust" is in that section. APA 7 form for a source without page numbers (section name in quotation marks plus "section") is correct.
- Not in the brief's list: the Copilot prompt "Where are some holes in [a meeting participant's] argument?" (Microsoft, n.d.) is a direct quotation with no locator. See R5.

## 4. Register findings

- **R1 (overclaim, tense).** "Across domains, people follow AI warnings even when the warnings are wrong." generalises from four studies in present tense. Replace with: "In several studies, people followed AI warnings even when the warnings were wrong."
- **R2 (missing hedge).** "The costs fall on the person checked." is categorical, yet the paragraph after it ("Two findings bound these conclusions") limits it. Replace with: "Costs can also fall on the person checked."
- **R3 (spelling consistency).** v6 uses American spelling elsewhere ("visualization") but British "judgement" in the section 3 heading and "judgements" in the Mayer paragraph. APA uses American spelling. Replace "Effects of erroneous flags on judgement and on the person flagged" with "Effects of erroneous flags on judgment and on the person flagged", and "laypeople's judgements of an expert" with "laypeople's judgments of an expert".
- **R4 (awkward phrasing).** "In the interviews of Opdenbusch et al. (2025), four participants described directors as" reads as if the authors were interviewed. Combined with the attribution fix in section 3, replace with: "Opdenbusch et al. (2025) report that, according to four participants, directors were "too shy to ask questions" (p. 6), and that most participants identified the budget as the board's only cybersecurity decision."
- **R5 (APA 7 citation error).** A direct quotation from a web page needs a locator (section name or paragraph number); "(Microsoft, n.d.)" has none, and the page is not in `references/` to check one. Paraphrasing avoids the need without changing the claim: replace "suggests, in a private pane, the prompt "Where are some holes in [a meeting participant's] argument?" (Microsoft, n.d.)" with "suggests, in a private pane, a prompt asking where the holes lie in a meeting participant's argument (Microsoft, n.d.)". If the author prefers to keep the quotation, add the section heading of the support page as "(Microsoft, n.d., "<section heading>" section)" after checking it on the live page.
- **R6 (unclear antecedent, repetition).** "The present study combines them: the checker is held by the non-expert who holds power over the expert" has no clear antecedent for "them" (the nearest plural is "ratings") and repeats "held ... holds". Replace with: "The present study combines these elements: the checker is held by the non-expert with power over the expert".
- Tense otherwise consistent: past for findings, present for theory (Mayer, Kempt and Nagel, Sharma's definitions) and for the present study. Headings are all noun phrases. "Danry et al.'s (2020) participants" and "D. Kim" / "P. H. Kim" initials are correct APA 7.

## 5. Rules checklist

| Rule | Result |
|---|---|
| No em-dashes | Pass (0 U+2014 in the whole file; "−1 to +1" uses a minus sign; reference pages use en dashes) |
| Body 1,000 to 1,500 words | Pass: **1,450** words (lines 3 to 33 without headings, split on whitespace; v5 was 1,281). Only 50 words of headroom; the fixes above are net neutral within a few words |
| One sentence that politician results are not carried over | Pass ("Findings on well-known politicians, whose reputations are settled, are not carried over.") |
| Trust in the AI is not an outcome | Pass (trust in the AI appears only as what Gupta et al. measured) |
| Gupta et al. (2026) only in the gap section | Pass |
| Dobbs absent | Pass |
| No H1 to H4 labels | Pass |
| Every citation has an entry and every entry is cited | Pass (all 30 entries cited; no citation without entry) |
| Reference list identical to v5 | Pass (diff empty) |
| CISO, VLM, METI spelled out at first use | Pass (CISO and VLM in the opening sentence, METI in the Mayer paragraph; SEC is never abbreviated, so needs no introduction) |

## 6. Ranked fixes

1. Meaning change, section 2 topic sentence.
   Old: "Their accuracy under live conditions is limited."
   New: "Evaluations of related checkers report limited accuracy."

2. Lowry attribution and locator.
   Old: "and one expert director they quote observed that for most boards "all they can do is assess the credibility of the executives that are put in front of them" (p. 783)."
   New: "and one expert director they interviewed observed that for most boards "all they can do is assess the credibility of the executives that are put in front of them" (p. 798)."

3. Power to act (von Schenk does not fit "power over the person").
   Old: "Flags also prompt those with power over the person assessed to act."
   New: "Flags also prompt those in a position to act on them to do so."

4. Opdenbusch attribution (authors' words, not participants').
   Old: "In the interviews of Opdenbusch et al. (2025), four participants described directors as "too shy to ask questions" (p. 6), and most identified the budget as the board's only cybersecurity decision."
   New: "Opdenbusch et al. (2025) report that, according to four participants, directors were "too shy to ask questions" (p. 6), and that most participants identified the budget as the board's only cybersecurity decision."

5. Overclaim in section 3 topic sentence.
   Old: "Across domains, people follow AI warnings even when the warnings are wrong."
   New: "In several studies, people followed AI warnings even when the warnings were wrong."

6. Opening scenario (restore the camera).
   Old: "A board member directs a content checker built on a vision-language model (VLM) at a slide presented by the chief information security officer (CISO); the checker flags a purported fault, and the CISO responds."
   New: "A board member points a camera-based checker built on a vision-language model (VLM) at a slide presented by the chief information security officer (CISO); the checker flags a purported fault, and the CISO responds."

7. Drafted, not generated.
   Old: "Our flags are generated with a real VLM and are all incorrect for the company presented."
   New: "Our flags are drafted with a real VLM and are all incorrect for the company presented."

8. Copilot quotation without locator.
   Old: "suggests, in a private pane, the prompt "Where are some holes in [a meeting participant's] argument?" (Microsoft, n.d.)."
   New: "suggests, in a private pane, a prompt asking where the holes lie in a meeting participant's argument (Microsoft, n.d.)."

9. Hedge on costs.
   Old: "The costs fall on the person checked."
   New: "Costs can also fall on the person checked."

10. Antecedent and repetition in the gap section.
    Old: "The present study combines them: the checker is held by the non-expert who holds power over the expert,"
    New: "The present study combines these elements: the checker is held by the non-expert with power over the expert,"

11. American spelling.
    Old: "## Effects of erroneous flags on judgement and on the person flagged" and "laypeople's judgements of an expert"
    New: "## Effects of erroneous flags on judgment and on the person flagged" and "laypeople's judgments of an expert"
