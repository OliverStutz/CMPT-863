# C-09 independent check of `09-Connections.md`

Checked 2026-10-19 against the readings themselves, not against report R. Copies used: Gaver PDF (`week7/`), Tangible Bits and Zooids (`week5/`), Wobbrock, Morris and Wilson (`Presentation-Week4/`), and freshly downloaded copies of the URLs named in the file: Weiser (Guelph PDF), Wobbrock and Kientz (interactions-16.pdf), Greenberg and Buxton (reading18.pdf), Olsen (Wayback pre-print, 7 pages), Villarreal-Narvaez et al. (dis-20.01.pdf), Shin et al. (Wayback KAIST PDF). Text was extracted page by page with pdftotext and every phrase was matched on the stated page. Breazeal was not opened (alternates only; scanned PDF).

## 1. Claims on the slide and in the Say script

| # | Claim | Where checked | Verdict |
|---|---|---|---|
| 1 | Olsen: "If one has used the tool to solve three diverse problems ... the stronger the claim," section "Generality" | Olsen pre-print PDF p. 4, section "Generality" | Confirmed, exact |
| 2 | Slide: Olsen "the space between" | same | Confirmed |
| 3 | Gaver: portfolio "establishes an area" (p. 944), dashed "a portfolio" omitted | Gaver PDF 8, p. 944 | Confirmed; omission handled correctly |
| 4 | Say: "Olsen supplies what Gaver asserts but never shows: how a portfolio could be extensible" | Gaver pp. 944, 945 | Overstated. On p. 944 Gaver does sketch a mechanism: comparing items makes clear "both the dimensions along which a designer's choices may range and the invariances among them." What he does not give is Olsen's point that more and more diverse examples strengthen the claim. See should-fix 2 |
| 5 | Gaver: "the 'extensibility and verifiability'" (p. 945) | Gaver PDF 9 | Confirmed |
| 6 | Tangible Bits: "one user interface metaphor that can be applied seamlessly across the metaDESK, ambientROOM, and transBOARD platforms" (p. 240) | TB PDF 7, printed p. 240 | Confirmed, exact |
| 7 | Gaver: annotated portfolio balances "specific, detailed examples of design practice, and articulations of the issues, values and themes which characterise the relations among the collection" (p. 944) | Gaver PDF 8 | Confirmed (sentence continues "and to which the examples suggest answers"; stopping early is fine) |
| 8 | Fifteen years, 1997 to 2012 | TB footer "CHI 97"; Gaver CHI 2012 | Confirmed |
| 9 | Greenberg and Buxton: "the existence proof as used in HCI is confirmative hypothesis testing, for the evaluator is seeking confirmatory evidence" (p. 113) | reading18.pdf PDF 3, printed p. 113 | Confirmed, exact |
| 10 | Greenberg and Buxton: "For really novel innovations ... risky hypothesis testing." (p. 118) | reading18.pdf PDF 8, printed p. 118 | Confirmed, exact |
| 11 | Say: Greenberg and Buxton make the diagnosis "for all of HCI" | reading18.pdf p. 113 | Slight overreach: they discuss HCI evaluation of new techniques (the existence proof), not all HCI research. See should-fix 3 |
| 12 | Gaver: "practice is usually intended to confirm theories, not falsify them" (p. 940) | Gaver PDF 4 | Confirmed |
| 13 | Gaver: "testimony to the fertility of the overarching theory"; "productive research programmes" (p. 941) | Gaver PDF 5 | Confirmed |
| 14 | Say: "Gaver drops refutation for design theory" | Gaver p. 941, "The difficulty of verifying design theory, at least through falsification, is not a flaw" | Fair paraphrase |
| 15 | Wobbrock and Kientz: "while we should be wary of constraining our imaginations, we should embrace giving definition to the knowledge we produce." (p. 44, Conclusion) | interactions-16.pdf PDF 7, printed p. 44 | Confirmed, exact |
| 16 | Wobbrock and Kientz: "Theories should be testable and falsifiable" (p. 41), endnote 6 Popper | interactions-16.pdf PDF 4, printed p. 41 | Confirmed |
| 17 | Gaver: "the design research community should be wary of impulses towards convergence and standardisation" (p. 937) | Gaver PDF 1 | Confirmed |
| 18 | Wobbrock, Morris and Wilson: agreement score "reflects, in a single number, the degree of consensus among participants" (p. 1087) | Week 4 PDF 5, printed p. 1087 | Confirmed, exact |
| 19 | DIS 2020 review: title, four authors, "N=216 gesture elicitation studies," p. 855 | dis-20.01.pdf p. 1 (printed 855) | Confirmed. Note: the same page says elicitation studies "were introduced in 2005 by Wobbrock et al. [122]" (guessability) and the 2009 paper was "the first application ... to hand-gesture input." The 2009 paper itself says its agreement process "was adopted from prior work [33]" (p. 1087). See should-fix 4 |
| 20 | Gaver: Lakatos programmes "predict novel facts" (p. 940) | Gaver PDF 4 (Lakatos quoted, [19] p. 24) | Confirmed; it is Gaver quoting Lakatos, which the Watch out handles correctly |
| 21 | Say: "it passes Gaver's own fertility test, so fertility does not separate design theory from science" | Gaver pp. 940, 941 | Does not follow as a point against Gaver. Lakatos is an account of science; Gaver uses it to show design theory need not fail by Popper, and says the verdict "depends on the account of science one adopts" (p. 941). A science paper passing a philosophy-of-science test is expected. The Watch out concedes this, the Say does not. See should-fix 4 |
| 22 | Shin et al.: "We conducted the second study in a real-life setting for ecological validity and observed posture correction for four hours." (p. 5) | Shin PDF p. 5 | Confirmed, exact |
| 23 | Shin et al.: "t= 3.54, p= 0.002" (p. 7) | Shin PDF p. 7 | Confirmed |
| 24 | Shin et al.: "all of the participants guessed that the monitor was trying to correct their posture, except P2" (p. 7) | Shin PDF p. 7 | Confirmed, exact; Say "nearly every participant" is fair |
| 25 | Gaver: "tests which are simultaneously unambiguous and ecologically valid are inconceivable" (p. 940) | Gaver PDF 4 | Confirmed; the sentence is followed by the ambiguity example, as the Watch out says |
| 26 | Say: "as I predicted in my Week 6 commentary" | `week6/Toward Sociable Robots.md` line 26 (Slow Robots section, line 20 on) | Quotation exact. "Predicted" is wrong in kind: the commentary was written after reading the paper, which itself reports the guessing on p. 7. A classmate can say so. See should-fix 5 |
| 27 | Footer and Watch out: Gaver cites none of the nine readings | Gaver reference list p. 946, searched for Weiser, Ishii, Greenberg, Buxton, Olsen, Wobbrock, Kientz, Breazeal, Shin, Le Goc | Confirmed |
| 28 | Transition: "These readings all came before 2012" | Wobbrock and Kientz 2016; Shin et al. 2019 | False. Two of the six readings on the slide postdate the essay. See blocker 1 |
| 29 | Olsen pre-print is seven pages, March 2007; UIST pages 251 to 258 | pdfinfo of downloaded copy: 7 pages | Page count confirmed; UIST pages not checked |
| 30 | Section 02 quotations (G&B p. 111; Olsen abstract) | not re-opened for this check | Not checked (outside brief) |

Directions and "so what" per row: O1 supports, follows (with the wording fix). T1 instance, follows; labelled "we think." G1 supports diagnosis, challenges cure, follows. K3 challenges, follows; the Watch out correctly says Wobbrock and Kientz do not reply to Gaver. M2 does not follow as stated (row 21). S1 half right, follows, and the behaviour versus meaning caveat in the Watch out is the right defence.

## 2. Gaver pages cited in the file

All Gaver pages cited on the slide, in the notes and in the alternates were phrase-searched on the stated printed page: 937 (wary, self-policing, "clear expectations and standards," "some sympathy"), 938 (manifestos' "primary function," "borrow theories to justify it," "inappropriate 'scientific' models," borrowed theories "used both to inspire"), 939 ("hard core"), 940 (confirm not falsify, "inconceivable," "predict novel facts," "dynamic machine," "designing for X," "is generative," "creating what might be," synthetic nature "incompatible"), 941 (fertility, productive programmes, "ability to inspire new designs," Stolterman "prepared-for-action"), 942 (shared values, "route to discovery"), 943 (iPad, "fresh beginnings"), 944 ("establishes an area," annotation definition, "theory nexus," "as opposed to demonstrators," "design space around themselves," portfolios to programmes, "embody the myriad"), 945 ("extensibility and verifiability," "subsume," "definite facts"), 946 ("consensus can come to look like constraint," "usefully subvert"). All confirmed.

## 3. Alternates

Completeness: report R has 31 connections (W1 to W4, K1 to K4, G1 to G4, O1 to O3, M1 to M3, T1 to T4, Z1 to Z3, B1 to B3, S1 to S3). The slide uses six (O1, T1, G1, K3, M2, S1). The alternates list holds the other 25, each exactly once. Complete.

Reading pages in the alternates, phrase-checked on the stated page: Weiser Guelph pp. 1, 5, 8; Wobbrock and Kientz pp. 40, 42, 44; Greenberg and Buxton pp. 113, 114, 115, 117, 118, 119; Olsen pre-print pp. 3, 6 to 7, 7; Wobbrock, Morris and Wilson pp. 1083, 1085, 1090; Tangible Bits pp. 234, 235, 236, 238; Zooids pp. 98, 101, 103, 105; Shin pp. 2, 6, 8. All confirmed. Breazeal pp. 169 to 174 not re-read (the file says so honestly).

Counts: Say script 227 words, against a ceiling of 227.5 (130 x 1.75). Passes, with no slack. On-slide bullets 99 words including the bold names, over the "about 90" guide. No em-dashes or en-dashes in the file.

Summary: 30 slide and Say claims checked. 25 confirmed as written, 5 need wording (rows 4, 11, 21, 26, 28). Of the quotations alone (20 from readings and Gaver on slide and Say), 20 of 20 are exact with correct pages.

## 4. Blockers

1. **Transition states a false fact.** "These readings all came before 2012" is wrong: Wobbrock and Kientz is 2016 and Shin et al. is 2019, both on the slide. Replace the Transition with:
   "Two of these readings, Wobbrock and Kientz and Shin, already came after the essay; here is what else followed from 2012 on."

## 5. Should-fix (exact replacement wording)

1. **Slide row 3 mixes page systems.** "Same diagnosis (p. 940), split cure (p. 118)" puts a Gaver page and a Greenberg and Buxton page side by side with no label. Replace the row with:
   "- **Greenberg and Buxton 2008:** "confirmatory evidence" (p. 113), split cure (p. 118). Gaver: same diagnosis (p. 940)."

2. **Olsen "never shows" overstates (row 4).** Replace "We argue Olsen supplies what Gaver asserts but never shows: how a portfolio could be extensible." with:
   "We argue Olsen states what Gaver only sketches: more diverse examples make a stronger general claim."
   In the Watch out, replace "Gaver asserts on p. 945 that portfolios allow "the 'extensibility and verifiability'" the critics asked for, but does not say how" with: "Gaver asserts on p. 945 that portfolios allow "the 'extensibility and verifiability'" the critics asked for, and on p. 944 says comparison shows "the dimensions along which a designer's choices may range," but never says why more or more varied examples make the claim stronger".

3. **Greenberg and Buxton scope (row 11).** Replace "for all of HCI" with "for HCI evaluation".

4. **M2 credit and logic (rows 19, 21).** Replace "Wobbrock, Morris and Wilson is an empirical method paper, and a 2020 review covers 216 gesture elicitation studies. We argue it passes Gaver's own fertility test, so fertility does not separate design theory from science." with:
   "Wobbrock, Morris and Wilson extended an empirical method begun in 2005; a 2020 review covers 216 such studies. By Gaver's fertility test that programme thrives, so we argue fertility rescues design theory without setting it apart from science."
   On the slide, replace "216 studies reviewed, 2020." with "216 such studies by 2020." Add to the Watch out: "The elicitation method dates from Wobbrock et al. 2005 (DIS 2020 review, p. 855); the 2009 paper adopted its agreement process "from prior work [33]" (p. 1087). Credit the programme, not this paper alone."

5. **"Predicted" (row 26).** Replace "as I predicted in my Week 6 commentary" with "as I flagged in Week 6". In the Watch out, replace "Oliver's Week 6 commentary predicted the confound" with "Oliver's Week 6 commentary flagged the confound".

6. **Rhetorical question on the slide.** Section 1.5 forbids questions in on-slide text unless the brief asks. Replace ""Inconceivable" (p. 940)?" with "Gaver half right (p. 940)."

With fixes 2 to 5 applied together the Say script is 227 words (counted), so it stays inside the limit.

## 6. Nice-to-have

- Trim the slide from 99 to about 90 words: drop "("Generality")" from the slide (it stays in the notes) and shorten row 6 to "**Shin et al. 2019:** "ecological validity" (p. 5), yet purpose guessed (p. 7). Gaver half right (p. 940)."
- The Watch out for S1 could add that the study was a two-day baseline versus treatment comparison (p. 5), which is why "inconceivable" looks too strong for behavioural claims.
- Verification log rows 102 and 103 say "per report R"; both are now independently confirmed (commentary line 26; reference list p. 946) and can read "OK."
- Breazeal pages (alternates B1 to B3) remain checked only by report R; re-read the scanned pages visually if Oliver swaps one in.

## 7. Single most valuable improvement

Fix the M2 line (should-fix 4). As written, the Say presents "fertility does not separate design theory from science" as a hit on Gaver, but Lakatos is a philosophy of science and Gaver uses him exactly to say the design and science line depends on the account one adopts (p. 941). A sharp classmate will make that reply in one sentence. The replacement keeps the point (fertility rescues design theory but does not make it distinct), credits the 2005 origin of the method the DIS review counts, and costs no extra words.
