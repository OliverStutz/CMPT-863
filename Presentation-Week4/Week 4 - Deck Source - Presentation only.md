# Week 4 deck source, presentation only: sections 01 to 12

Generated 2026-10-01 from the section files in `slides/` by `slides/_build-deck-source.py`. Edit the section files, not this file, then regenerate. Speaker notes, verification logs and open items live in the section files; the full version with notes is `Week 4 - Deck Source.md`. Run sheet, placeholders and figure assets are in `slides/00-Overview.md`.

## Slide 0: Title

**Time:** 0:30 · **Speaker:** Oliver · **Visual:** Figure 1 (PDF page 1, printed p. 1083), the participant panning with a flat hand

**On the slide**
- **User-Defined Gestures for Surface Computing**
- Jacob O. Wobbrock (University of Washington, Information School), Meredith Ringel Morris and Andrew D. Wilson (Microsoft Research)
- CHI 2009, Boston, pp. 1083 to 1092
- CMPT 863 Human-Computer Interaction · Week 4 · [date]
- Presented by Oliver Stutz

---

# 01 Background, 2008

**Purpose:** Put the room in 2008, show which touch and gesture devices an ordinary adult could already have used, then give the paper's motivation and the study setup as the base for everything after. **Slides:** 2. **Time:** 2:30. **Speaker:** Oliver. **Previous section:** Title slide. **Next section:** 02 The problem.

## Slide 01.1: 2008, where touch stood

**Time:** 1:15 · **Visual:** horizontal timeline 1985 to 2009, the study marked as a red band ending at August 2008; below it a single row of five small device icons with years. No figure from the paper.

**On the slide**
- 1985: Buxton's multi-touch tablet, Toronto
- 2001 DiamondTouch (MERL), 2002 SmartSkin (Sony)
- Feb 2006: Jeff Han's multi-touch demo at TED
- Jan 2007: iPhone announced; June 2007: on sale, with pinch
- May 2007: Microsoft Surface unveiled; April 2008: in AT&T stores
- **Before Aug 2008: this study** (patent filed 4 Aug 2008)
- April 2009: CHI, Boston; Oct 2009: Windows 7 with multi-touch
- Already in ordinary hands: laptop touchpads (1994), Palm Graffiti (1996), iPod touch wheel (2002), Nintendo DS (2004), Wii (2006)

## Slide 01.2: Why this study, and what it was

**Time:** 1:15 · **Visual:** left column "Why" with four short lines; right column "What" with Figure 1 (participant at the Surface, prompt "Pan, 1 hand," PDF page 1) as a small photo above four setup lines.

**On the slide**
- **Why (the paper says)**
- Tabletops for individuals, groups, multi-display rooms, fun (p. 1083)
- Keyboard and mouse "no longer preferable" (p. 1083)
- Gestures "defined by system designers" (p. 1083)
- Public tables: "the importance of immediate usability is high" (p. 1085)
- **What**
- 20 non-technical adults at a Surface prototype (p. 1085)
- 27 commands: see the effect, perform the cause; 1 and 2 hands; think aloud; no recognizer (p. 1085)
- 1,080 gestures (p. 1086)
- Out: taxonomy, agreement scores, 48 gestures for 22 commands (pp. 1086 to 1088)

---

# 02 The problem

**Purpose:** State the problem the paper addresses, why it mattered in 2008 and why it still matters, with analysis rather than summary. **Slides:** 1. **Time:** 1:15. **Speaker:** Oliver. **Previous section:** 01 Background, 2008. **Next section:** 03 Pressure and ownership.

## Slide 02.1: The problem: gestures built for the recognizer

**Time:** 1:15 · **Visual:** quote card at the top with the p. 1083 sentence; below it two columns, "Built from the recognizer outward" and "Users asked before"; tagline across the bottom. No figure from the paper.

**On the slide**
- The paper says: gestures "defined by system designers," "chosen out of concern for reliable recognition" (p. 1083)
- Yet the space is vast: "almost anything one can do with one's hands" could be a gesture (p. 1083)
- Built from the recognizer outward:
  - Wu's tabletops: "not elicited from users" (p. 1085)
  - Malik et al.: 1, 2, 3 or 5 fingers, "may be problematic" (p. 1085)
- Asked before: Good 1984, Epps 2006, Nielsen 2004 (p. 1084)
- Claim: "first to employ users, rather than principles" (p. 1084)
- **We argue: users had been asked before; nobody had built a whole surface gesture set from their agreement.**

---

# 03 Pressure and ownership

**Purpose:** Show who owned multi-touch research and gesture patents around 2005 to 2009, and test the hypothesis that Microsoft was under pressure, keeping documented facts apart from our interpretation. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 02 The problem. **Next section:** 04 The method.

## Slide 03.1: Who owned the gestures?

**Time:** 1:30 · **Visual:** a horizontal timeline, 2005 to 2009, Apple events above the line, Microsoft events below, a vertical marker "study done (before Aug 2008)". Right-hand box split in two: "Documented" and "We argue". No figure from the paper.

**On the slide**
- **Documented.** Apple: buys FingerWorks (2005); Jobs, Jan 2007, on multi-touch: "boy, have we patented it"; patent 7,479,949 granted Jan 2009; Cook, Jan 2009: "we will not stand for having our IP ripped off"
- Microsoft: hires Buxton (Dec 2005); Surface shown May 2007, in stores Apr 2008; Windows 7 touch demo May 2008
- Aug 2008: Microsoft files a patent application on this study's method (inventors Morris, Wobbrock, Wilson); abandoned 2014
- Apple and Microsoft: patent cross-licence since 1997
- **We argue:** a race to own gesture conventions; Apple's public lawsuit threat came only in 2009, after the study

---

# 04 The method

**Purpose:** Show how the guessability method works, where it comes from, why each design choice was made, and what it can and cannot deliver. **Slides:** 2. **Time:** 3:30. **Speaker:** Oliver. **Previous section:** 03 Pressure and ownership. **Next section:** 05 How the research was conducted.

## Slide 04.1: The method: a guessability study, from referent to sign

**Time:** 1:40 · **Visual:** left, a lineage strip with four dated stops (1984 Good, 2004 Nielsen, 2005 Wobbrock EdgeWrite, 2009 this paper). Right, a three-step loop drawn as a one-way arrow: "animation of the effect (referent)" → "participant performs a gesture (sign), thinking aloud" → "logged, no response from the system". Optional thumbnail of Figure 1 (PDF p. 1, printed p. 1083), the pan prompt.

**On the slide**
- Lineage: Good et al. 1984, "user-derived" email commands (p. 1084); Nielsen et al. 2004, "a similar approach" (p. 1084); Wobbrock et al. 2005, guessability and agreement score, EdgeWrite (p. 1087)
- Referent = a command's effect, shown as an animation; sign = the gesture performed (McNeill, p. 1085)
- Three choices:
  1. Ask users: "user behavior is rarely so systematic" (p. 1083)
  2. Effect first, plain 2D shapes, nothing from Windows or Mac (p. 1085)
  3. No recognizer: "a monologue in which the user's behavior is always acceptable" (p. 1085)
- Microsoft filed a patent application on this procedure in August 2008

## Slide 04.2: From 1,080 gestures to a set, step by step

**Time:** 1:50 · **Visual:** a five-step funnel across the top (1,080 → group → score → assign → alias → 48 gestures). Below it, Eq. 1 and Eq. 3 reproduced from p. 1087 (PDF p. 5), with a stacked bar of 20 people split 11 / 3 / 3 / 3 next to Eq. 3.

**On the slide**
- 1,080 gestures, 1,074 kept (p. 1086)
- **Group** identical gestures per referent; "identical" undefined (p. 1087)
- **Score** agreement score A = mean over referents of Σ (group share)² (Eq. 1)
  - Select single, 1 hand: 11, 3, 3, 3 → 0.37 (Eq. 3, p. 1087)
  - Range 1/20 (everyone different) to 1 (everyone the same). Not a percentage
- **Assign** largest group; conflicts go to the larger group (p. 1087)
- **Alias** 48 gestures for 22 referents, covering 57.0% of all gestures proposed (pp. 1087, 1088)
- **Around it** taxonomy (p. 1086); authors' own sets (p. 1090)
- **Limits** no shared gesture for close or task switch, people tapped imaginary buttons (p. 1087); guessability of the set never tested (p. 1091)

---

# 05 How the research was conducted

**Purpose:** show how the elicitation study was actually run, then put on screen three results the run produced: the designers' poor prediction of users' gestures, one hand versus two, and finger use. **Slides:** 2. **Time:** 3:00. **Speaker:** Oliver. **Previous section:** 04 The method. **Next section:** 06 Sample and culture.

## Slide 05.1: The study as run

**Time:** 1:15 · **Visual:** left, a smaller Figure 1 photo (PDF page 1, p. 1083), or replace it with a simple diagram of one trial (animation plays, participant gestures, nothing responds) since Figure 1 already appears on the title slide and 01.2; right, Table 1 (PDF page 4, p. 1086), cropped to the two columns of referent names and ratings, at least half the slide's height, with the MEAN row (2.70) highlighted; if it cannot be read at that size, show only the five lowest and five highest rows plus MEAN. A small pipeline strip along the bottom: hear and see effect, reset, 1-hand gesture, 2-hand gesture, ratings.

**On the slide**
- 20 paid volunteers, 9 women, mean age 43.2 (sd 15.6), all right-handed, none had used a tabletop or iPhone (p. 1085)
- Microsoft Surface prototype, 24 x 18 inches, 1024 x 768 (p. 1085)
- 27 referents, random order; authors rated complexity 1 to 5, mean 2.70 (p. 1086)
- Per referent: 1-hand and 2-hand gesture, think aloud, preference, two 7-point ratings (pp. 1086, 1088)
- 1,080 gestures, 6 discarded; touch logs, video, notes (p. 1086)
- Baseline: each author's own gestures, made before the study (p. 1090)

## Slide 05.2: Three things the run revealed

**Time:** 1:45 · **Visual:** three columns. Column 1: three big numbers 43.5% / 60.9% / 19.1% with the quote card "the authors are either thinking creatively or are hopelessly lost!" (p. 1090). Column 2: Figure 3 (p. 1088) with the 2-hand line highlighted at insert, maximize, enlarge. Column 3: a hand icon with 1 to 3 fingers marked "single point", 4 marked "gray area", 5 or palm marked "something more" (p. 1090).

**On the slide**
- **Designers missed.** Each author matched 43.5% of the user set; all three 60.9%; 19.1% of each author's gestures no participant tried (p. 1090)
- **One hand wins.** Preferred for 25 of 27 referents; tie for insert and maximize; two hands never preferred (p. 1088). Yet 17 of 48 set gestures (35.4%) are 2-hand (p. 1088)
- **Fingers barely matter.** 13 of 20 varied finger count; 2 said it mattered. 1 to 3 fingers: "single point"; 4: "gray area" (p. 1090)

---

# 06 Sample and culture

**Purpose:** show who was missing from the sample (younger and older people, other ethnicities and cultures), how culture is built into some gestures of the set, and why the 2008 adult arrived already shaped by the mouse and by single-point touch. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 05 How the research was conducted. **Next section:** 07 Findings.

## Slide 06.1: Who was in the room, and who was not

**Time:** 1:30 · **Visual:** left, a "sample card" with the participant line (p. 1085) and the limitation quote (p. 1091); right, two crops of Figure 4 (PDF page 7, p. 1089): row 3, the Accept "draw check", Reject "draw 'X'" and Help "draw '?'" panels; and the bottom row, right-most panel, Next "draw line across object", with three small culture icons underneath (head nod, ✓ versus ○, thumbs-up). **[classmate]** slot at the bottom.

**On the slide**
- 20 adults, mean age 43.2, sd 15.6 (p. 1085); no range, no results by age
- No ethnicity reported. The authors: "all non-technical literate American adults; undoubtedly, children, Eastern, or uneducated participants would behave differently" (p. 1091)
- Symbols in the set: check, X, ?, a line drawn left to right (Fig. 4, p. 1089)
- Meanings vary: in Japan ○ is correct and ✓ can mark wrong; in Bulgaria a nod means no; thumbs-up insults in Iran. We argue: the set's symbols may not travel
- No iPhone or tabletop users (p. 1085), yet "I’m a child of the mouse" (p. 1090)
- **[classmate]**

---

# 07 Findings

**Purpose:** Rank the findings that matter most, then show what the timing data say about thinking versus doing. **Slides:** 2. **Time:** 2:30. **Speaker:** Oliver. **Previous section:** 06 Sample and culture. **Next section:** 08 Technical obstacles.

## Slide 07.1: What people did

**Time:** 1:10 · **Visual:** ranked list of eight findings, each with a one-word solidity tag in a coloured chip (Counted, Coded, Observed, Self-rated). Small inset of Figure 2, Form columns only (PDF page 5, p. 1087), with the one-point path and one-point touch segments highlighted.

**On the slide**
1. Desktop "deeply rooted": about 72% mouse-like one-point touches or paths (p. 1090). *Coded*
2. One hand preferred for 25 of 27 referents (p. 1088). *Counted*
3. Agreement falls with complexity, r = -0.52 (p. 1087). *Counted*
4. 43.9% of gestures physical, treating shapes as real objects; complex commands drew metaphors and symbols (p. 1087). *Coded*
5. Finger count rarely matters: 13 of 20 varied it, 2 cared (p. 1090). *Observed*
6. Opposite commands (enlarge, shrink) got mirror-image gestures, though never shown as pairs (p. 1090). *Observed*
7. Off-screen drags, mid-air gestures, imaginary widgets (p. 1090). *Observed*
8. Gestures many people chose got higher "good match" ratings from the people who made them (p. 1088). *Self-rated*

Footer, small type: Tags: Counted = tallied from logs or ratings; Coded = sorted by the authors; Observed = authors' notes and video; Self-rated = participants judged their own gestures

## Slide 07.2: Time tells a story: finding the gesture costs more than making it

**Time:** 1:20 · **Visual:** one horizontal time axis from 0 to 16 s with two bars. Top bar "thinking" (planning time) with marks at about 8 s (simple commands) and about 15 s (complex). Bottom bar "doing" (articulation time) spanning about 2 to 3.4 s. A small inset at the right: "rated easy 3.4 s / rated difficult 2.0 s." All numbers from p. 1088 (PDF page 6).

**On the slide**
- Thinking: about 8 s before touching for simple commands, about 15 s for complex ones; r = 0.71 with complexity (p. 1088)
- Doing: about 2 to 3.4 s, and complexity made no difference to it (p. 1088)
- Longer thinking went with lower "good match" ratings: good gestures came to mind fast (p. 1088)
- We argue: on a walk-up table the cost is in finding the gesture, not making it. That is what guessability is for
- Speed is not ease: gestures rated easy took about 3.4 s, rated difficult about 2.0 s (p. 1088). Comfortable gestures were slow and smooth, awkward ones quick and precise (we argue)
- Caveat: averages only, and the tests count 1,074 gestures as if 1,074 people had made them; 20 did

---

# 08 Technical Obstacles

**Purpose:** Show what the 2008 table could and could not sense, and how those limits shaped the study and its recommendations. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 07 Findings. **Next section:** 09 Limitations.

## Slide 08.1: What the 2008 table could not see

**Time:** 1:30 · **Visual:** left, a scale bar comparing pixel density, about 43 ppi (2008 Surface prototype) against 264 ppi (current iPad Pro); right, a side view of a table with three zones labelled "sensed: contact from below," "not sensed: above the table," "not sensed: beyond the edge." Optional inset: Figure 1 photo, PDF page 1 (p. 1083).

**On the slide**
- Display: 24" x 18" at 1024 x 768 (p. 1085), about 43 pixels per inch. A current iPad Pro: 264 ppi.
- Sensing: a vision system "from beneath the table"; contacts logged as ovals with millisecond timestamps (p. 1086). Cameras recorded the session separately (p. 1086).
- No recognizer, no touch feedback (p. 1085, p. 1090).
- Participants were told to touch the table (p. 1090). Some still gestured in the air and beyond the screen edge (p. 1090).
- A physics engine (Wilson et al. 2008): supported by 43.9% physical gestures, but "will probably not suffice" alone (p. 1091).
- Set claims "ease of recognition," never tested (p. 1091).

---

# 09 Limitations

**Purpose:** Separate the limitations the authors state from the ones we add, and size up what a sample of 20 can carry. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 08 Technical obstacles. **Next section:** 10 Connections.

## Slide 09.1: Limitations, theirs and ours

**Time:** 1:30 · **Visual:** two columns, "The authors say (p. 1091)" left, "We add" right; a small footer card: "n = 20. Across 216 later elicitation studies the average was 25, and 20 the most common choice (Villarreal-Narvaez et al., DIS 2020, p. 859)"; **[classmate]** slot under the right column

**On the slide**

*The authors say (p. 1091)*
- No revising earlier gestures
- No application, organizational or cultural context
- "children, Eastern, or uneducated participants would behave differently"
- Set not validated; no recognizer

*We add*
- n = 20 (p. 1085): enough to show strong agreement, too few to compare groups such as age or culture
- Authors grouped and coded; no inter-rater reliability (pp. 1086 to 1087)
- The statistics count 1,074 gestures as if 1,074 different people had made them. Twenty did, so the p-values are too confident (p. 1088)
- Gestures judged only by the people who made them, yet "goes a long way to validating" (p. 1088)
- Five of the 27 commands are desktop ideas: cut, paste, minimize, maximize, task switch (p. 1085). Desktop questions invite desktop gestures
- The set matches 57% of what participants did; the other 43% of their gestures get nothing (p. 1087)
- "ease of recognition" without a recognizer (p. 1091)
- Two of three authors at Microsoft, a Microsoft prototype; the patent application goes unmentioned
- **[classmate]**

---

# 10 Connections to other research

**Purpose:** Place the paper against the four course readings from weeks 1 to 3 and against three of its own references, each with a "so what" for how we judge it. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 09 Limitations. **Next section:** 11 Research directions.

## Slide 10.1: Four readings, three roots

**Time:** 1:30 · **Visual:** a two-column card grid. Left column: four cards, one per course reading (author, year, one short quote). Right column: the matching point in the paper with its page. Bottom strip: three small "roots" chips (Furnas, Hutchins et al., Morris). No figure from the paper.

**On the slide**
- **Wobbrock and Kientz 2016:** methods need "repeated validation" (p. 41). Here one study (p. 1088); set validation a "next step" (p. 1091).
- **Greenberg and Buxton 2008:** early evaluation converges on "that which we are already familiar with" (p. 115). Elicitation too: about 72% mouse-like (p. 1090).
- **Olsen 2007:** "new design participants." Non-technical users author the set (p. 1084).
- **Weiser 1991:** wall boards at 1024 x 768 for walk-up "casual use." This Surface: same resolution, for public tables where "immediate usability" matters (p. 1085).
- **Roots:** Furnas et al., aliasing (p. 1088); Hutchins et al., gulf of execution (p. 1085); Morris et al., gestures for groups, yet this study tested one person at a time (p. 1084).

---

# 11 Research directions

**Purpose:** show the research agenda the paper sets itself, what actually followed, and which questions are still open for a student today. **Slides:** 1. **Time:** 1:30. **Speaker:** Oliver. **Previous section:** 10 Connections to other research. **Next section:** 12 Discussion bridge.

## Slide 11.1: What the paper asked for, and what came next

**Time:** 1:30 · **Visual:** two columns, "The authors' agenda (p. 1091)" on the left and "What followed" on the right; a strip along the bottom with the Figure 4 "pinch" panel (Enlarge (Shrink)3, p. 1089) and three open questions

**On the slide**
- **Authors' agenda (p. 1091).** Next: reverse the study (clips to commands); build a vision-based recognizer. Later: see all commands first; application context; "children, Eastern, or uneducated participants." Technology: sense above the table and beyond the edge.
- **What followed: the method became a standard**
  - **2020, Villarreal-Narvaez, Vanderdonckt, Vatavu and Wobbrock:** a review of 216 studies that used this method. It calls this paper the first of its kind and the most cited, and notes that 20 participants became the norm because of it
  - **2010, the same three authors:** 22 new people judged the gestures without knowing who made them. They preferred the gestures of larger groups. That tests liking, not guessing
  - **2014, Morris, Wobbrock and colleagues:** named the old-habits problem "legacy bias" and proposed fixes: several gestures per command, show the device's abilities first, work in pairs
  - **2015, Vatavu and Wobbrock:** corrected the agreement score (this paper's 0.32 becomes 0.28) and added a significance test
- **Twist (we argue).** No product adopted the set as published; pinch became universal anyway (Fig. 4, p. 1089).
- **Open.** Prior exposure as a variable? The reverse test? Other cultures, mid-air, VR?

---

# 12 Discussion bridge: our reading, the tensions, the hand-over

**Purpose:** close the presentation with the presenters' own reading of the paper, the tensions the class will argue about, and the plan for the 30-minute discussion. **Slides:** 3. **Time:** 3:00. **Speaker:** Oliver (sole presenter). **Previous section:** 11 Research directions. **Next:** the discussion itself (plan in `16-Discussion-Guide.md`).

Source material: revision 2 of `Week 4 - Slides.md`, slides 20 to 22, which passed two independent Opus reviews after correction. Written by the orchestrator, 2026-09-30.

## Slide 12.1: Our reading of the paper

**Time:** 0:45 · **Visual:** one sentence large, three supports beneath

**On the slide**
- **The lasting contribution is the method. The evidence in this paper for the method is thinner than its influence.**
- The guessability test the authors proposed was never run (p. 1091). The 2010 follow-up tested preference, which is partial support.
- "Popularity identifies better gestures" rests on participants rating their own gestures (p. 1088). No independent judges.
- The study measured the repertoire participants brought in. The paper calls that a finding about users. We argue it is also a property of the method.

## Slide 12.2: Tensions worth arguing about

**Time:** 1:15 · **Visual:** four large questions, three smaller items beneath, with **[classmate]** credits beside the ones classmates raised

**On the slide**

Four questions with more than one defensible answer
1. **Should the most popular gesture win?** About 72% of gestures were mouse-like (p. 1090). Does novice consensus lock in old habits?
2. **Did the commands shape the answers?** Cut, paste, minimize, maximize and task switch are desktop concepts (p. 1085).
3. **Does popularity identify better gestures, or ones people enjoyed inventing?** Only the gestures' own makers rated them (p. 1088).
4. **Is behavior without feedback the right design target?** Participants gestured in the air and off-screen (p. 1090).

Three smaller things we found interesting
- **Most guessable, or best?** 19.1% of each author's gestures were never tried by any participant (p. 1090).
- **Symbolic gestures import writing.** Check, X and question mark are Western marks (Figure 4, p. 1089).
- **Single users on a group device.** Minimize moves to the bottom "as defined by user's seating position" (Figure 4), different for each person at the table.

## Slide 12.3: Over to you

**Time:** 1:00 · **Visual:** three or four short quotes from classmates' commentaries with names, each paired with the tension it speaks to (filled after the commentaries are posted), then the three-step plan

**On the slide**
- **What you wrote.** **[classmate]** **[classmate]** **[classmate]**
- **How the next 30 minutes run**
  1. **Warm-up, 5 min: a live elicitation.** We show the effect of two commands, one the paper's participants agreed on and one they did not (Figure 3). On three, everyone performs the gesture they would use, on the desk. We count the groups and compute the agreement score together.
  2. **Pairs, 5 min:** one card per pair, a question and two roles, one each.
  3. **Full group, 20 min:** we follow what comes up.

