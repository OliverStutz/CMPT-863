# Opus G: review of the presentation-only deck (on-slide text, no notes)

Reviewed 2026-09-30. Source read: `Week 4 - Deck Source - Presentation only.md` (19 slides, title to 12.3). Notes in the section files and the paper (PDF, printed pp. 1083 to 1092) used only to check meaning and numbers. No file edited.

Every "Current" string below is quoted exactly as it appears in the section file named, so it can be used as a find-and-replace target. Where a replacement removes a line, "Replacement" says so.

## Summary

- **Slides that stand alone as they are:** 5 of 19 (Title, 01.1, 05.1, 05.2, 12.1). Three more (01.2, 06.1, 08.1) need only a small fix. The other 11 depend on the notes in at least one bullet, carry an unlabelled judgment, or contain a dangling reference.
- **Em-dashes:** none found (none in the presentation-only deck or the section files 01 to 12; no en-dashes either).
- **Accuracy:** every number on the slides was checked against the paper and matches (1,080 / 1,074; 0.37 from 11, 3, 3, 3; 57.0%; 22 of 27; 43.5% / 60.9% / 19.1%; 31 and 17 of 48, 35.4%; 25 of 27; r = -0.52, 0.71, -0.59; 5.6 vs 4.9; 8 s and 15 s; 3.4 s and 2.0 s; 43.9%; about 72%; 13 of 20 and 2; 43.2, sd 15.6; 24" x 18", 1024 x 768; mean complexity 2.70; "ease of recognition" in the Conclusion, p. 1091; [32] is Wilson et al. 2008). One factual problem: slide 09.1 calls the authors "Microsoft authors", but Wobbrock was at the University of Washington. One rounding problem: slide 12.2 drops the paper's "about" from 72%.
- **Word counts** (on-slide text, page references and [classmate] slots excluded): 12.2 has 189 words and 12.3 has 109. Everything else is between 66 and 85. 03.1 (85), 08.1 (82) and 09.1 (81) are at the upper edge but readable.

### Counts per category

| Category | Must | Should | Optional | Total |
|---|---|---|---|---|
| (1) Depends on notes | 3 | 5 | 0 | 8 |
| (2) Dangling references | 2 | 1 | 0 | 3 |
| (3) "Paper says" vs "we argue" labels | 0 | 5 | 1 | 6 |
| (4) Ambiguous or misleading | 3 | 7 | 3 | 13 |
| (5) Too dense | 2 | 0 | 0 | 2 |
| (6) Visual directions | 0 | 4 | 0 | 4 |
| (7) Term consistency | 0 | 3 | 6 | 9 |
| (8) Em-dashes | 0 | 0 | 0 | 0 |
| (9) Accuracy | 1 | 0 | 1 | 2 |
| **Total** | **11** | **25** | **11** | **47** |

Each item below is filed under one category, the main one.

### Slide-by-slide verdict

| Slide | Stands alone? | Main issue |
|---|---|---|
| 0 Title | Yes | none |
| 01.1 | Yes | "patent" used before it is introduced (optional) |
| 01.2 | Nearly | "Out:" and "commands" vs "referents" (optional) |
| 02.1 | No | tagline is our claim but unlabelled; "almost anything" bullet sits under the wrong heading |
| 03.1 | No | "lawsuit avoidance" answers a hypothesis that only the notes state |
| 04.1 | No | "(section 03)" |
| 04.2 | No | "1/\|P\|", "Cannot fill close, task switch", "57.0% covered" of what |
| 05.1 | Yes | visual: whole Table 1 too small to read (should) |
| 05.2 | Yes | small grammar fix (optional) |
| 06.1 | Nearly | culture bullet never says why it matters |
| 07.1 | No | tag chips have no key; "makers" misreadable |
| 07.2 | No | "Rated easy took about 3.4 s" reads as thinking time; "hard means precise" cryptic |
| 08.1 | Nearly | "[32]" |
| 09.1 | No | "(see 06)", "(see 05, 06)", "Microsoft authors" wrong |
| 10.1 | No | Weiser line garbled; "gulf of execution" without Hutchins; Morris line cryptic |
| 11.1 | No | follow-up papers unnamed; "fixed"; "Twist" unlabelled |
| 12.1 | Yes | none |
| 12.2 | No | 189 words in 1:15 |
| 12.3 | No | production note printed on the slide; "top of Figure 3" |

---

## MUST FIX (11)

### M1. Slide 04.1 (category 2, dangling reference). File `04-Method.md`
- **Current:** `- Patent application on the procedure, August 2008 (section 03)`
- **Replacement:** `- Microsoft filed a patent application on this procedure in August 2008`
- **Reason:** the audience has no section numbers; the ownership slide already showed the date, so the line needs no pointer.

### M2. Slide 09.1 (category 2, dangling reference). File `09-Limitations.md`
- **Current:** `- "children, Eastern, or uneducated participants would behave differently" (see 06)`
- **Replacement:** `- "children, Eastern, or uneducated participants would behave differently"`
- **Reason:** "(see 06)" means nothing on screen; the page is already given in the column heading (p. 1091).

### M3. Slide 09.1 (category 1, depends on notes, plus a dangling reference). File `09-Limitations.md`
- **Current:** `- n = 20 (p. 1085): strong consensus yes, subgroups no (see 05, 06)`
- **Replacement:** `- n = 20 (p. 1085): enough to show strong agreement, too few to compare groups such as age or culture`
- **Reason:** "strong consensus yes, subgroups no" is shorthand only the speaker can unpack, and "(see 05, 06)" dangles.

### M4. Slide 09.1 (category 9, accuracy). File `09-Limitations.md`
- **Current:** `- Microsoft authors, own prototype; patent unmentioned`
- **Replacement:** `- Two of three authors at Microsoft, a Microsoft prototype; the patent application goes unmentioned`
- **Reason:** Wobbrock was at the University of Washington (p. 1083, and the title slide says so); "Microsoft authors" contradicts the title slide.

### M5. Slide 03.1 (category 1, depends on notes). File `03-Pressure-and-Ownership.md`
- **Current:** `- **We argue:** a race to own conventions, not lawsuit avoidance`
- **Replacement:** `- **We argue:** a race to own gesture conventions; Apple's public lawsuit threat came only in 2009, after the study`
- **Reason:** "not lawsuit avoidance" answers a hypothesis stated only in the notes; the replacement gives the reason that is on the slide's own timeline (Cook, Jan 2009).

### M6. Slide 04.2 (category 4, misleading). File `04-Method.md`
- **Current:** `- **Cannot** fill close, task switch; never tests guessability (p. 1091)`
- **Replacement:** `- **Limits** no shared gesture for close or task switch, people tapped imaginary buttons (p. 1087); guessability of the set never tested (p. 1091)`
- **Reason:** "Cannot fill close" reads as a closing operation that fails; the paper says close and task switch got no gesture because participants used imaginary widgets (p. 1087).

### M7. Slide 07.1 (category 1, depends on notes). File `07-Findings.md`
- **Current:** `1. Desktop "deeply rooted": about 72% mouse-like one-point touches or paths (p. 1090). *Coded*`
- **Replacement:** insert one line above it, keep the line itself:
  `*Tags: Counted = tallied from logs or ratings; Coded = sorted by the authors; Observed = authors' notes and video; Self-rated = participants judged their own gestures*`
- **Reason:** the coloured chips carry the slide's argument about how solid each finding is, but nothing on screen says what the four tags mean. The line adds 25 words (79 to 104). If that is too many, put the key in a footer in smaller type.

### M8. Slide 07.1 (category 4, misleading). File `07-Findings.md`
- **Current:** `8. Popular gestures rated better by makers (p. 1088). *Self-rated*`
- **Replacement:** `8. Gestures many people chose got higher "good match" ratings from the people who made them (p. 1088). *Self-rated*`
- **Reason:** in a deck about designers versus users, "makers" will be read as the designers or authors.

### M9. Slide 07.2 (category 4, misleading). File `07-Findings.md`
- **Current:** `- Rated easy took about 3.4 s; rated difficult about 2.0 s (p. 1088)`
- **Replacement:** `- Doing: gestures rated easy took about 3.4 s to perform, those rated difficult about 2.0 s (p. 1088)`
- **Reason:** the result is counterintuitive, and on a slide about thinking time the audience will read these as planning times. They are articulation times.

### M10. Slide 12.2 (category 5, too dense). File `12-Discussion-Bridge.md`
189 words to read in 1:15 while someone speaks. Either split the slide into 12.2a (four questions) and 12.2b (three smaller things), or replace line by line as follows. This brings it to about 100 words; the cut sentences move into the speaker notes.
- **Current:** `1. **Should the most popular gesture win?** 72% of gestures were mouse-like; "I'm a child of the mouse" (p. 1090). Does consensus from novices lock in old habits? Guessability against learnability against efficiency.`
  **Replacement:** `1. **Should the most popular gesture win?** About 72% of gestures were mouse-like (p. 1090). Does novice consensus lock in old habits?`
- **Current:** `2. **Did the commands shape the answers?** Cut, paste, minimize, maximize and task switch are desktop concepts (p. 1085). Who should choose the referents?`
  **Replacement:** `2. **Did the commands shape the answers?** Cut, paste, minimize, maximize and task switch are desktop concepts (p. 1085).`
- **Current:** `3. **Does popularity identify better gestures, or gestures people enjoyed inventing?** The ratings are self-reports (p. 1088). What would independent validation look like?`
  **Replacement:** `3. **Does popularity identify better gestures, or ones people enjoyed inventing?** Only the gestures' own makers rated them (p. 1088).`
- **Current:** `4. **Is behaviour without feedback the right design target?** Real users adapt to a system's responses. Participants gestured in the air and off-screen (p. 1090).`
  **Replacement:** `4. **Is behavior without feedback the right design target?** Participants gestured in the air and off-screen (p. 1090).`
- **Current:** `- **Floor or ceiling?** 19.1% of each author's gestures were never tried (p. 1090). Does elicitation find the most guessable gesture or the best one?`
  **Replacement:** `- **Most guessable, or best?** 19.1% of each author's gestures were never tried by any participant (p. 1090).`
- **Current:** `- **Symbolic gestures import writing.** The check, the X and the question mark are Western marks (Figure 4). The paper's own caveat names "Eastern" participants (p. 1091).`
  **Replacement:** `- **Symbolic gestures import writing.** Check, X and question mark are Western marks (Figure 4, p. 1089).`
- **Current:** `- **Single users on a group device.** Minimize means move to the bottom of the surface "as defined by user's seating position" (Figure 4), a different direction for every person at a table.`
  **Replacement:** `- **Single users on a group device.** Minimize moves to the bottom "as defined by user's seating position" (Figure 4), different for each person at the table.`
- **Reason:** too much to read during 1:15 of speech. The replacements also restore the paper's "about" before 72% (p. 1090), make "Floor or ceiling?" readable without the notes, and spell "behavior" the way the paper does.

### M11. Slide 12.3 (category 5, density; a production note printed on the slide). File `12-Discussion-Bridge.md`
- **Current:** `- **What you wrote.** **[classmate]** **[classmate]** **[classmate]** (three or four recurring themes or sharp observations, each credited by name and paired with the tension above that it speaks to; filled after the commentaries are posted)`
- **Replacement:** `- **What you wrote.** **[classmate]** **[classmate]** **[classmate]**`. Then move the parenthesis into the Visual line, so it reads: `**Visual:** three or four short quotes from classmates' commentaries with names, each paired with the tension it speaks to (filled after the commentaries are posted), then the three-step plan`
- **Reason:** this is a build instruction, not slide text. If the placeholder survives into the deck it ends up on screen.

---

## SHOULD FIX (25)

### S1. Slide 02.1 (category 3, label). File `02-Problem.md`
- **Current:** `- **Users had been asked before. Nobody had built a whole surface gesture set from their agreement.**`
- **Replacement:** `- **We argue: users had been asked before; nobody had built a whole surface gesture set from their agreement.**`
- **Reason:** the tagline is our reading of the novelty claim and sits right below the paper's own "first to employ users"; without a label the two blur.

### S2. Slide 02.1 (category 3, label). File `02-Problem.md`
- **Current:** `- Stated: gestures "defined by system designers,"`
- **Replacement:** `- The paper says: gestures "defined by system designers,"`
- **Reason:** "Stated" does not say who stated it; this matches the "(the paper says)" label on 01.2.

### S3. Slide 02.1 (category 4, ambiguous). File `02-Problem.md`
- **Current:** `  - "almost anything one can do with one's hands" (p. 1083)`
- **Replacement (outdented to a top-level bullet, placed right after the "Stated" line):** `- Yet the space is vast: "almost anything one can do with one's hands" could be a gesture (p. 1083)`
- **Reason:** under "Built from the recognizer outward" the quote reads as one more example of recognizer-driven design; in the paper it describes how large the gesture space is.

### S4. Slide 03.1 (category 3, label). File `03-Pressure-and-Ownership.md`
- **Current:** `- Apple: buys FingerWorks (2005);`
- **Replacement:** `- **Documented.** Apple: buys FingerWorks (2005);`
- **Reason:** the Visual line has a "Documented" box, but the bullets mark only "We argue". The label makes the split visible in the text itself, not only in the layout.

### S5. Slide 04.2 (category 1, depends on notes). File `04-Method.md`
- **Current:** `  - Range 1/|P| to 1. Not a percentage`
- **Replacement:** `  - Range 1/20 (everyone different) to 1 (everyone the same). Not a percentage`
- **Reason:** |P| is never defined on screen. With 20 participants per referent, the paper's lower bound |Pr|^-1 is 1/20.

### S6. Slide 04.2 (category 4, ambiguous). File `04-Method.md`
- **Current:** `- **Alias** 48 gestures, 22 referents, 57.0% covered (pp. 1087, 1088)`
- **Replacement:** `- **Alias** 48 gestures for 22 referents, covering 57.0% of all gestures proposed (pp. 1087, 1088)`
- **Reason:** "57.0% covered" could be read as 57% of referents or of participants; the paper says of all gestures proposed (p. 1087).

### S7. Slide 04.2 (category 7, term). File `04-Method.md`
- **Current:** `- **Score** A = mean over referents of Σ (group share)² (Eq. 1)`
- **Replacement:** `- **Score** agreement score A = mean over referents of Σ (group share)² (Eq. 1)`
- **Reason:** 01.2 and 12.3 say "agreement score", 11.1 says "agreement rate" (a different measure), and the bare "A" is never tied to either.

### S8. Slide 04.1 (category 7, term). File `04-Method.md`
- **Current:** `- Referent = effect shown; sign = gesture performed (McNeill, p. 1085)`
- **Replacement:** `- Referent = a command's effect, shown as an animation; sign = the gesture performed (McNeill, p. 1085)`
- **Reason:** 01.2 says "27 commands" and 04.2 onward says "referents". The bridge the paper itself makes ("the effects of 27 commands (i.e., the referents)", p. 1085) belongs on screen.

### S9. Slide 06.1 (category 3, label). File `06-Sample-and-Culture.md`
- **Current:** `- Meanings vary: in Japan ○ is correct and ✓ can mark wrong; in Bulgaria a nod means no; thumbs-up insults in Iran`
- **Replacement:** `- Meanings vary: in Japan ○ is correct and ✓ can mark wrong; in Bulgaria a nod means no; thumbs-up insults in Iran. We argue: the set's symbols may not travel`
- **Reason:** without the speaker, the line is trivia; the consequence for the set is our claim and needs to be said and labelled.

### S10. Slide 07.1 (category 4, ambiguous). File `07-Findings.md`
- **Current:** `4. 43.9% physical; complex commands drew metaphors and symbols (p. 1087). *Coded*`
- **Replacement:** `4. 43.9% of gestures physical, treating shapes as real objects; complex commands drew metaphors and symbols (p. 1087). *Coded*`
- **Reason:** "43.9% physical" does not say 43.9% of what, nor what "physical" means; the taxonomy is never shown on screen.

### S11. Slide 07.1 (category 4, ambiguous). File `07-Findings.md`
- **Current:** `6. Opposites got reversible gestures, never shown as pairs (p. 1090). *Observed*`
- **Replacement:** `6. Opposite commands (enlarge, shrink) got mirror-image gestures, though never shown as pairs (p. 1090). *Observed*`
- **Reason:** "Opposites" and "never shown as pairs" both need the speaker; one example fixes both.

### S12. Slide 07.2 (category 1, depends on notes). File `07-Findings.md`
- **Current:** `- We argue: hard means precise, not long`
- **Replacement:** `- We argue: gestures felt hard when they were quick and precise, not when they were long`
- **Reason:** "hard means precise" is a riddle without the spoken explanation.

### S13. Slide 07.2 (category 7, term). File `07-Findings.md`
- **Current:** `- Complexity vs goodness r = -0.59: 5.6 vs 4.9 of 7 (p. 1088)`
- **Replacement:** `- Complex referents got lower "good match" ratings: 4.9 vs 5.6 of 7 for simple ones, r = -0.59 (p. 1088)`
- **Reason:** "goodness" is used nowhere else on screen; the line above it calls the same scale "good match", and the bare pair "5.6 vs 4.9" does not say which number belongs to which group.

### S14. Slide 08.1 (category 2, dangling reference). File `08-Technical-Obstacles.md`
- **Current:** `- Physics engine [32]: supported by 43.9% physical gestures, "will probably not suffice" alone (p. 1091).`
- **Replacement:** `- A physics engine (Wilson et al. 2008): supported by 43.9% physical gestures, but "will probably not suffice" alone (p. 1091).`
- **Reason:** a bracketed reference number means nothing to the audience; [32] is Wilson et al., UIST 2008.

### S15. Slide 09.1 (category 1, depends on notes). File `09-Limitations.md`
- **Current:** `- Chi-square on N = 1074 non-independent gestures (pp. 1087 to 1088)`
- **Replacement:** `- Tests treat 1,074 gestures from 20 people as independent (pp. 1087 to 1088)`
- **Reason:** "non-independent" needs the explanation that each person made about 54 of them. Also fixes "1074" vs "1,074" (04.2).

### S16. Slide 09.1 (category 4, ambiguous). File `09-Limitations.md`
- **Current:** `- Only self-ratings, yet "goes a long way to validating" (p. 1088)`
- **Replacement:** `- Gestures judged only by the people who made them, yet "goes a long way to validating" (p. 1088)`
- **Reason:** "self-ratings" of what is unclear on screen, and the slide must not suggest the ratings were circular (participants did not know group sizes, per the notes).

### S17. Slide 10.1 (category 4, ambiguous). File `10-Connections.md`
- **Current:** `- **Weiser 1991:** boards for "casual use," 1024x768, like this Surface, meant for public "immediate usability" (p. 1085).`
- **Replacement:** `- **Weiser 1991:** wall boards at 1024 x 768 for walk-up "casual use." This Surface: same resolution, for public tables where "immediate usability" matters (p. 1085).`
- **Reason:** as written it is unclear which phrase belongs to Weiser and which to the paper. The notes warn that only resolution and walk-up use match, not form factor.

### S18. Slide 10.1 (category 1, depends on notes). File `10-Connections.md`
- **Current:** `- **Roots:** Furnas, aliasing (p. 1088); gulf of execution (p. 1085); Morris, group gestures, single-user study (p. 1084).`
- **Replacement:** `- **Roots:** Furnas et al., aliasing (p. 1088); Hutchins et al., gulf of execution (p. 1085); Morris et al., gestures for groups, yet this study tested one person at a time (p. 1084).`
- **Reason:** the visual names three roots but the text names only two; "group gestures, single-user study" needs the speaker to supply the tension.

### S19. Slide 11.1 (category 1, depends on notes). File `11-Research-Directions.md`
- **Current:** `- **What followed.** GI 2010: preference, not guessability. 2014: legacy bias, fixed by production, priming, partners. CHI 2015: agreement rate. DIS 2020: 216 studies.`
- **Replacement:** `- **What followed.** Morris et al., GI 2010: tested preference, not guessability. Morris et al. 2014: "legacy bias," reduced by production, priming, partners. Vatavu and Wobbrock, CHI 2015: agreement rate, a revised agreement score. Villarreal-Narvaez et al., DIS 2020: review of 216 elicitation studies.`
- **Reason:** venue-year labels with no authors cannot be followed or looked up. "Fixed" overclaims (the 2014 piece proposes ways to reduce legacy bias). "Agreement rate" needs tying to the score shown on 04.2. This adds about 18 words; trim the agenda bullet if needed.

### S20. Slide 11.1 (category 3, label). File `11-Research-Directions.md`
- **Current:** `- **Twist.** Set not adopted as published; pinch universal anyway (Fig. 4, p. 1089).`
- **Replacement:** `- **Twist (we argue).** No product adopted the set as published; pinch became universal anyway (Fig. 4, p. 1089).`
- **Reason:** "not adopted" is our claim, not the paper's, and the page reference makes it look sourced from the paper.

### S21. Slide 12.3 (category 4, ambiguous). File `12-Discussion-Bridge.md`
- **Current:** `We show the effect of two commands, one from the top of Figure 3 and one from the tail.`
- **Replacement:** `We show the effect of two commands, one the paper's participants agreed on and one they did not (Figure 3).`
- **Reason:** "top" and "tail" of a chart the audience saw for a few seconds, 15 minutes earlier, cannot be decoded on screen.

### S22. Slide 05.1 (category 6, visual). File `05-How-Conducted.md`
- **Current:** `right, Table 1 with the 27 referents and complexity ratings (p. 1086), with the mean row 2.70 highlighted.`
- **Replacement:** `right, Table 1 (PDF page 4, p. 1086), cropped to the two columns of referent names and ratings, at least half the slide's height, with the MEAN row (2.70) highlighted; if it cannot be read at that size, show only the five lowest and five highest rows plus MEAN.`
- **Reason:** a full 27-row table at a quarter of a slide is unreadable. The page number and the crop were missing.

### S23. Slide 05.1 (category 6, visual). File `05-How-Conducted.md`
- **Current:** `left, Figure 1 photo of a participant panning (p. 1083);`
- **Replacement:** `left, a smaller Figure 1 photo (PDF page 1, p. 1083), or replace it with a simple diagram of one trial (animation plays, participant gestures, nothing responds) since Figure 1 already appears on the title slide and 01.2;`
- **Reason:** Figure 1 is specified on five slides (Title, 01.2, 04.1, 05.1, 08.1); the builder needs to know which uses to keep. Suggest keeping it on Title and 01.2 only.

### S24. Slide 06.1 (category 6, visual). File `06-Sample-and-Culture.md`
- **Current:** `right, a crop of the bottom rows of Figure 4 (PDF page 7) showing Accept "draw check", Reject "draw 'X'", Help "draw '?'", Next "draw line across object",`
- **Replacement:** `right, two crops of Figure 4 (PDF page 7, p. 1089): row 3, the Accept "draw check", Reject "draw 'X'" and Help "draw '?'" panels; and the bottom row, right-most panel, Next "draw line across object",`
- **Reason:** Accept, Reject and Help are in the third row of Figure 4, not the bottom rows; only Next is in the bottom row.

### S25. Slide 09.1 (category 6, visual). File `09-Limitations.md`
- **Current:** `a small footer card "n = 20" with the DIS 2020 figure;`
- **Replacement:** `a small footer card: "n = 20. Across 216 later elicitation studies the average was 25, and 20 the most common choice (Villarreal-Narvaez et al., DIS 2020, p. 859)";`
- **Reason:** "the DIS 2020 figure" does not tell the builder which figure or what text to put on the card; the numbers are in this file's verification table.

---

## OPTIONAL (11)

### O1. Slide 01.1 (category 4). File `01-Background-2008.md`
- **Current:** `- **Before Aug 2008: this study** (patent filed 4 Aug 2008)`
- **Replacement:** `- **Before Aug 2008: this study** (Microsoft patent application on its method, 4 Aug 2008)`
- **Reason:** a patent is mentioned before slide 03.1 explains it; the audience will ask "which patent?".

### O2. Slide 01.2 (category 7). File `01-Background-2008.md`
- **Current:** `- Out: taxonomy, agreement scores, 48 gestures for 22 commands (pp. 1086 to 1088)`
- **Replacement:** `- Results: taxonomy, agreement scores, 48 gestures for 22 commands (pp. 1086 to 1088)`
- **Reason:** "Out" is jargon. The "commands" here is fine if S8 lands, because 04.1 then equates the two terms.

### O3. Slide 02.1 (category 4). File `02-Problem.md`
- **Current:** `- Asked before: Good 1984, Epps 2006, Nielsen 2004 (p. 1084)`
- **Replacement:** `- Asked before: Good 1984, Nielsen 2004, Epps 2006 (p. 1084)`
- **Reason:** chronological order, as on 04.1's lineage strip.

### O4. Slide 04.2 (category 3). File `04-Method.md`
- **Current:** `"identical" undefined (p. 1087)`
- **Replacement:** `"identical" never defined, our note (p. 1087)`
- **Reason:** a critique sitting inside a "how it works" step; a light label keeps the slide's descriptive tone.

### O5. Slide 05.2 (category 4). File `05-How-Conducted.md`
- **Current:** `19.1% of each author's gestures no participant tried (p. 1090)`
- **Replacement:** `19.1% of each author's gestures were tried by no participant (p. 1090)`
- **Reason:** the current word order makes readers stop and parse it twice.

### O6. Slide 08.1 vs 05.1 (category 7). File `08-Technical-Obstacles.md`
- **Current:** `- Display: 24" x 18" at 1024 x 768 (p. 1085)`
- **Replacement:** `- Display: 24 x 18 inches at 1024 x 768 (p. 1085)`
- **Reason:** 05.1 writes "24 x 18 inches"; use one form throughout.

### O7. Slide 03.1 (category 7). File `03-Pressure-and-Ownership.md`
- **Current:** `- Apple and Microsoft: patent cross-licence since 1997`
- **Replacement:** `- Apple and Microsoft: patent cross-license since 1997`
- **Reason:** the rest of the deck and the paper use American spelling (behavior, color). The Visual line of 07.1 has "coloured"; it is not on screen but could be fixed in the same pass.

### O8. Slide 06.1 (category 7). File `06-Sample-and-Culture.md`
- **Current:** `"I’m a child of the mouse"` (curly apostrophe)
- **Replacement:** `"I'm a child of the mouse"`
- **Reason:** 12.2 uses a straight apostrophe; the two copies of the same quote should match.

### O9. Slides 06.1, 11.1 vs 12.2 (category 7)
- **Current:** `(Fig. 4, p. 1089)` (06.1, 11.1) and `(Figure 4)` (12.2)
- **Replacement:** pick one, preferably `(Figure 4, p. 1089)` everywhere.
- **Reason:** consistency. 12.2 also lacks the page.

### O10. Slides 05.1, 06.1 (category 7, plain wording)
- **Current (06.1):** `- 20 adults, mean age 43.2, sd 15.6 (p. 1085); no range, no results by age`
- **Replacement:** `- 20 adults, mean age 43.2, spread (sd) 15.6 (p. 1085); no age range, no results by age`
- **Reason:** follows the project's plain-wording preference; "sd" alone may not register with everyone.

### O11. Slide 07.1 (category 9, accuracy of tag). File `07-Findings.md`
- **Current:** `3. Agreement falls with complexity, r = -0.52 (p. 1087). *Counted*`
- **Replacement:** `3. Agreement falls with complexity as the authors rated it, r = -0.52 (p. 1087). *Counted*`
- **Reason:** agreement is counted, but complexity was rated by the three authors (Table 1, p. 1086). The "Counted" chip alone overstates how solid the finding is.

---

## Checks that passed (no change needed)

- **Em-dashes or en-dashes:** none in the presentation-only deck or the section files 01 to 12.
- **Quotes checked verbatim against the PDF:** "no longer preferable", "defined by system designers", "chosen out of concern for reliable recognition", "the importance of immediate usability is high", "not elicited from users", "may be problematic", "first to employ users, rather than principles", "user behavior is rarely so systematic", "a monologue in which the user's behavior is always acceptable", "deeply rooted", "I'm a child of the mouse", "single point", "gray area", "the authors are either thinking creatively or are hopelessly lost!", "goes a long way to validating", "will probably not suffice", "children, Eastern, or uneducated participants would behave differently", "as defined by user's seating position", "ease of recognition". All match.
- **Paper-internal note:** p. 1091 says "Figure 2 indicates that agreement is low after the first seven referents"; it means Figure 3. No slide repeats this, but do not quote that sentence with "Figure 2".
- **Not verifiable from the paper (taken from the section verification tables, not re-checked here):** the 03.1 timeline (FingerWorks, Jobs and Cook quotes, patent 7,479,949, the Microsoft application and its abandonment), the 01.1 device dates, the course-reading quotes on 10.1, and the 06.1 cultural examples.
