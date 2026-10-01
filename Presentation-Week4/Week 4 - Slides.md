# Week 4 – Slides: User-Defined Gestures for Surface Computing

Slide-by-slide content for the 30-minute presentation on Wobbrock, Morris and Wilson, "User-Defined Gestures for Surface Computing," CHI 2009, pp. 1083–1092. What to say for each slide is in [[Week 4 - Speaker Notes]]. Every number below is sourced in [[Week 4 - Presentation Notes]] and [[progress]]; page numbers are the printed ones (1083–1092).

Revision 2 (2026-09-30): all fixes from `reports/slides-review-cross-examination.md` applied. Old slides 11 and 12 merged; slide 18 rewritten as "Whose habits did the study capture?"; Figure 3 readings corrected from a re-render; 22 slides, 25:45 planned, 4 min buffer.

Structure:
- **Part A – Rundown** (slides 2–10, about 10 min): what the paper is, how it was done, what came out.
- **Part B – Key questions** (slides 11–19, about 12.5 min): the course's seven questions, organized by what matters.
- **Part C – Bridges to discussion** (slides 20–22, about 3 min): our thesis and the tensions the class will argue about.

Open placeholders are marked **[TODO: …]**. Places to credit a classmate's commentary are marked **[classmate]**.

---

## Run sheet

| # | Slide | Part | Time | Speaker |
|---|---|---|---|---|
| 1 | Title | – | 0:30 | A |
| 2 | The paper in one slide | A | 1:00 | A |
| 3 | Why 2009 was the moment | A | 0:45 | A |
| 4 | The problem | A | 1:00 | A |
| 5 | The core idea: referent → sign | A | 1:30 | A |
| 6 | Method at a glance | A | 0:45 | B |
| 7 | From 1,080 gestures to a gesture set | A | 1:30 | B |
| 8 | Taxonomy: what people actually did | A | 0:45 | B |
| 9 | Results: where people agreed | A | 1:30 | B |
| 10 | The user-defined gesture set | A | 1:00 | B |
| 11 | Findings, ranked, and the one to remember | B | 1:45 | A |
| 12 | What is the contribution, really? | B | 1:30 | A |
| 13 | Strengths | B | 0:45 | B |
| 14 | Limitations: theirs and ours | B | 1:30 | B |
| 15 | Three numbers that are not the same | B | 1:15 | B |
| 16 | Design decisions that shaped the result | B | 1:30 | B |
| 17 | Whose habits did the study capture? | B | 1:45 | A |
| 18 | Connections to other research and course topics | B | 1:00 | A |
| 19 | What came next, and what it suggests | B | 1:30 | B |
| 20 | Our reading of the paper | C | 0:45 | B |
| 21 | Tensions worth arguing about | C | 1:15 | A + B |
| 22 | Over to you | C | 1:00 | A + B |
| | **Total** | | **25:45** | about 4 min buffer |

Speaker A = Oliver, Speaker B = co-presenter. Slides 17 and 18 are Oliver's material; everything else can move. Spoken script: about 3,100 words in total, which fits 25:45 at about 120 words per minute; no block exceeds 137.

## Placeholders to fill before the draft submission (day before class, 11:59 PM)

- **[TODO: presentation date]** and **[TODO: co-presenter name]** on slide 1.
- **Slide 18** – section 6 of the notes (connections to course topics), Oliver's to write.
- **Slide 22** – classmates' commentary themes, with names. The **[classmate]** slots on slides 14, 17 and 21 take one credited line each.
- **Figure assets** – see the list below. Nothing has been rendered for the deck yet.

## Assets to prepare

Render from `Week 4 - User-Defined Gestures for Surface Computing.pdf` (PDF page numbers, not printed ones):

| Asset | PDF page | Printed page | Used on slide |
|---|---|---|---|
| Figure 1, participant at the Surface, prompt "Pan, 1 hand" | 1 | 1083 | 1, 2 |
| Table 1, conceptual complexity of the 27 commands | 4 | 1086 | 6 |
| Table 2, taxonomy | 4 | 1086 | 8 |
| Agreement equation and worked examples (Eq. 1–3) | 5 | 1087 | 7, 15 |
| Figure 2, taxonomy shares (stacked bars) | 5 | 1087 | 8 |
| Figure 3, agreement per command (line chart) | 6 | 1088 | 9 |
| Figure 4, the 48-gesture set (line drawings) | 7 | 1089 | 10, 16, 17 |

Command: `pdftoppm -r 200 -f 5 -l 5 -png "Week 4 - User-Defined Gestures for Surface Computing.pdf" assets/p5` then crop. Pages 2, 3, 8, 9, 10 are text only.

---

# Part A – Rundown: what the paper is about

## Slide 1 — Title

**Time:** 0:30 · **Speaker:** A · **Visual:** Figure 1 photo, full bleed, or a crop of Figure 4

On the slide:
- **User-Defined Gestures for Surface Computing**
- Jacob O. Wobbrock (University of Washington, Information School), Meredith Ringel Morris and Andrew D. Wilson (Microsoft Research)
- CHI 2009, Boston, pp. 1083–1092
- CMPT 863 · Week 4 · **[TODO: date]**
- Presented by Oliver Stutz and **[TODO: co-presenter]**

---

## Slide 2 — The paper in one slide

**Time:** 1:00 · **Speaker:** A · **Visual:** Figure 1 (participant, flat hand dragging, prompt "Pan, 1 hand")

On the slide:
- **Question.** What gestures would people invent for a touch table if nobody told them what to do?
- **Who.** 20 non-technical adults, Microsoft Surface prototype, 27 commands.
- **How.** See the effect of a command → perform the gesture that would cause it. One hand and two hands. Think aloud. **No recognizer.**
- **Data.** 1,080 gestures, logged, filmed from four angles.
- **Output.**
  1. A taxonomy of surface gestures
  2. An agreement score per command
  3. A 48-gesture user-defined set covering 22 of the 27 commands

---

## Slide 3 — Why 2009 was the moment

**Time:** 0:45 · **Speaker:** A · **Visual:** horizontal timeline

On the slide (timeline):
- 2005–06 · Jeff Han's multi-touch demo goes viral
- 2007 · iPhone ships with pinch-to-zoom · Microsoft Surface announced
- 2008 · Surface sold commercially · the study runs some time before the CHI deadline in autumn (the paper gives no date)
- April 2009 · **this paper**
- October 2009 · Windows 7 ships with built-in multi-touch

Below the timeline:
- Tabletops were moving into group, multi-display and **public** settings, where "the importance of immediate usability is high" (p. 1085).
- Two of the three authors were at Microsoft Research; the study ran on a Surface prototype.
- **Gesture vocabularies were being fixed in products at exactly this moment.**

---

## Slide 4 — The problem

**Time:** 1:00 · **Speaker:** A · **Visual:** quote card

On the slide:
> "To date, most surface gestures have been defined by system designers … somewhat arbitrary gesture sets whose members may be chosen out of concern for reliable recognition." (p. 1083)

- **Example.** Malik et al. gave 1-, 2-, 3- and 5-finger touches different meanings (p. 1085). A camera can count fingers. People, it turns out, mostly do not.
- **Known since Long et al.** Even people designing gestures "are sometimes poor at picking easily differentiable gestures" (p. 1084).
- **Why it is hard.** A mouse constrains motion to one point. Hands do not: "almost anything one can do with one's hands could be a potential gesture" (p. 1083). Huge design space + one designer's guess = the vocabulary problem (Furnas, ref. 8).

Tagline: **Gesture sets were designed for the recognizer, not the person. Few had asked what people would do, and nobody had built a full multi-touch set from the answers.**

---

## Slide 5 — The core idea: referent → sign

**Time:** 1:30 · **Speaker:** A · **Visual:** flow diagram, build in two steps

On the slide, diagram (step 1):
- **Referent** (the effect: a block disappears) → shown to participant → **Sign** (the gesture they think caused it)

Diagram (step 2):
- 20 participants × 27 referents → group identical signs → **agreement score** → largest group wins → **gesture set**

Three design choices, each with its reason:
1. **Elicit from users**, because they are the ones who have to guess. Designers organize gestures "in a principled, logical fashion," but "user behavior is rarely so systematic" (p. 1083).
2. **Show the effect first, in a neutral world of 2D blocks.** Fixes the meaning, leaves only the form open. No Windows or Mac imagery, because Epps et al. had conceded their Windows screenshots "may have biased" people toward mouse emulation (p. 1084–1085).
3. **No recognizer.** "A monologue in which the user's behavior is always acceptable" → the study sees "unrevised behavior": behaviour not reshaped by recognition feedback (p. 1085). The gulf of execution (Norman) is removed on purpose.

Footer: Procedure and agreement score adopted from the authors' EdgeWrite guessability work, 2005 (ref. 33). "Referent" and "sign" from McNeill.

---

## Slide 6 — Method at a glance

**Time:** 0:45 · **Speaker:** B · **Visual:** four-row table, Table 1 complexity bands as a small inset

On the slide:

| | |
|---|---|
| **Participants** | 20 paid volunteers from the general public. 9 women. Mean age 43.2 (sd 15.6). All right-handed. No tabletop or iPhone use, no computer scientists or designers. Restaurant host, musician, author, steelworker. |
| **Setup** | Microsoft Surface prototype, 24 × 18 in, 1024 × 768. Recorded voice plus an animation of the effect, then the scene resets and waits. 27 referents from desktop and tabletop systems (Table 1), each pre-rated for conceptual complexity by the three authors, 1 to 5, mean 2.70. |
| **Per command** | Random order. A one-hand gesture and a two-hand gesture, thinking aloud. Which do you prefer? Two 7-point ratings: *good match for its purpose*, *easy to perform*. |
| **Data and baseline** | 20 × 27 × 2 = 1,080 gestures (6 discarded). Touch logs with millisecond timestamps, four video angles, think-aloud notes. **Designer baseline:** before the study, each author designed their own set for the same 27 commands. |

Callout: **The system never recognized anything and never gave feedback. A design choice, not a hardware limit.**

---

## Slide 7 — From 1,080 gestures to a gesture set

**Time:** 1:30 · **Speaker:** B · **Visual:** the equation from p. 1087, worked example as a build

On the slide:

**Step 1 – Group.** For each command, the authors sort the 20 gestures into groups of "identical" gestures (p. 1087). The paper does not say what counted as identical; the Figure 4 caption implies that one to three fingers were treated as one.

**Step 2 – Score.** Agreement for one command = sum over groups of (group size ÷ 20)². Average over all commands.

Worked example from the paper, *move a little*, **two hands**: groups of 12, 3, 3 and 2 people
→ (12/20)² + (3/20)² + (3/20)² + (2/20)² = **0.42**

- Everyone identical → **1.0**. Everyone different → **0.05** (= 1/20).

**Step 3 – Assign.** The largest group wins the command. If one gesture would win two commands, "the referent with the largest group won the gesture" (p. 1087). What happened to the other command is not described; in the final set the same gesture reappears bound to a different target (slide 10).

**Result.** 48 gestures, 22 commands, no conflicts. Covers **57.0%** of everything participants proposed.

---

## Slide 8 — Taxonomy: what people actually did

**Time:** 0:45 · **Speaker:** B · **Visual:** Figure 2 (stacked bars), Table 2 headings beside it

On the slide:

Four dimensions (Table 2): **Form** (what the hand does) · **Nature** (symbolic, physical, metaphorical, abstract) · **Binding** (on the object, relative to the world, anywhere, mixed) · **Flow** (discrete or continuous)

Shares of all 1,080 gestures (Figure 2; Form is reported per hand, right hand shown):
- **Form.** One-point path 45% + one-point touch 27% = **72% "mouse-like."** Static pose + path 14%, dynamic pose 9%, static pose 4%. The left hand, present only in two-hand gestures, is also about 72% one-point.
- **Nature.** Physical 44% · abstract 27% · metaphorical 19% · symbolic 10%.
- **Binding.** Object-centric **70%** · world-dependent 18% · world-independent 9% · mixed 3%.
- **Flow.** Discrete 54% · continuous 46%.

Callout: Simple commands got physical gestures. Complex ones got metaphorical or symbolic ones (p. 1087).

---

## Slide 9 — Results: where people agreed

**Time:** 1:30 · **Speaker:** B · **Visual:** Figure 3, full width

On the slide, annotations on the chart (one-hand line unless stated):
- *Move a little* and *move a lot*: **A = 1.0**. One group each.
- Then a cliff. *Duplicate* 0.52; *rotate, select group, insert* about 0.48; *select single* 0.37.
- After the first seven commands nothing is above about 0.32. The tail falls steadily: *task switch* 0.21 … *help* 0.16 … *cut* 0.11.
- Two-hand line: peaks at *insert* and *maximize* (0.58) and *enlarge* (0.52); dips at *pan* and *close* (0.19) and *task switch* (0.13).

Numbers:
- Overall agreement: **0.32** one hand, **0.28** two hands (p. 1087).
- Agreement falls as the authors' complexity rating rises: **r = −0.52**.

Two separate recommendations the authors draw (p. 1091):
1. Beyond the first seven commands, **a widget as well as a gesture.**
2. Because enough participants tapped imaginary widgets: consider widgets for **delete, zoom in, zoom out, accept, reject, menu access, help.**

Footnote: the Discussion cites "Figure 2" for the cutoff after seven commands (p. 1091). It means this figure.

---

## Slide 10 — The user-defined gesture set

**Time:** 1:00 · **Speaker:** B · **Visual:** Figure 4, with three crops highlighted

On the slide:
- **48 gestures for 22 commands.** 31 one-handed, 17 two-handed (35%). Ten commands have one gesture, the rest have 2–5 aliases.
- **One gesture, many meanings, resolved by target.** The four enlarge/shrink gestures (pull apart with hands, pull apart with fingers, pinch, splay) *on an object* = resize · *on the background* = zoom · *atop an "openable" object* = open. Drag off-screen: *object* = delete · *dialog* = reject.
- **Idioms that emerged.**
  - *Jump* (Move₂, two hands): hold the object, tap the destination. Reused for delete (off-screen destination), paste (off-screen source), minimize (bottom edge). *Select group*: hold one object and tap the others.
  - *Drawn marks*: ✓ accept · ✗ reject · ? help · scratch-out = undo · slash = cut.
  - *Menu*: "pull out", a finger of the second hand dragged out from the held object.
- **Not in the set.** Insert and maximize reuse drag and enlarge. **Close and task switch got no gesture: participants tapped imaginary widgets.**
- **A design implication the authors draw:** do not demand precision. For two-hand enlarge, 2 of 14 resized along the diagonal and 12 went sideways (p. 1091).

Footer: any one-finger gesture may be done with 1–3 fingers (caption, Figure 4).

---

# Part B – Key questions

## Slide 11 — Findings, ranked, and the one to remember

**Time:** 1:45 · **Speaker:** A · **Visual:** three large numbers across the top, ranked table below

Top band:

**43.5%** · the share of the user-defined set that any one author had predicted · **60.9%** · all three authors combined · **19.1%** · of each author's gestures that no participant ever tried (p. 1090)

> The authors on their own misses: either "thinking creatively" or "hopelessly lost." (p. 1090)

Ranked table:

| | Finding | How solid |
|---|---|---|
| 1 | **Designers ≠ users.** Three HCI experts, including the study's own authors, covered under half of the user set individually and about 60% together. | The yardstick is the users' set, so designers cannot win by construction, and the paper concedes three experts cannot match the breadth of 20 people. It measures overlap, not quality. |
| 2 | **Finger count rarely matters.** 13 of 20 varied it; only 2 said it meant anything. 1–3 fingers = a point, 5 fingers or a palm = something more, four = grey area (p. 1090). | Robust. The paper does not say how fingers were treated when grouping gestures (slide 16). |
| 3 | **One hand beats two.** Preferred for 25 of 27 commands; two hands never preferred (p. 1088). | Robust. Two-hand gestures were required (slide 16). |
| 4 | **"It's a Windows world."** 72% of gestures were mouse-like, with no desktop imagery on screen. "I'm a child of the mouse" (p. 1090). | Robust number, open interpretation (slide 17). |
| 5 | **Agreement falls with complexity**, r = −0.52 (p. 1087). | Complexity was rated by the authors, but it also predicts planning time (r = 0.71; 8 s vs 15 s, p. 1088), so the ratings track something real. |
| 6 | **Popular gestures were rated better** by the people who made them (p. 1088). | Weakest: self-rated, no independent judges (slide 14). |

Also seen: dragging off-screen and gesturing above the table; opposite commands get reversible gestures although pairs were never shown together; gestures rated *easy* took longer (about 3.4 s) than those rated *hard* (about 2.0 s).

---

## Slide 12 — What is the contribution, really?

**Time:** 1:30 · **Speaker:** A · **Visual:** two columns, "as listed" vs "sorted by kind"

On the slide:

**As the authors list it (p. 1084)**
1. Quantitative and qualitative characterization, with a taxonomy
2. A user-defined gesture set
3. Insight into users' mental models
4. Implications for technology and UI design

**Sorted by kind**
- **Method:** effect-then-cause elicitation, no recognizer, consensus decides
- **Instruments:** the agreement score, the four-dimension taxonomy
- **Artifact:** 48 gestures for 22 commands, 57% coverage
- **Findings:** slide 11

**What lasted.** The method and agreement as the metric. The guessability test the authors proposed for the set (p. 1091) was never run; a 2010 follow-up tested preference instead (slide 19). No product adopted the set as published, though many of its members are now standard. The method became a genre: a 2020 systematic review covered **216 gesture elicitation studies** (Villarreal-Narvaez, Vanderdonckt, Vatavu and Wobbrock, DIS 2020). The numbered list leaves the method out because the authors credit it to their own 2005 work; the abstract still opens with "an approach … that relies on eliciting gestures."

**Novel vs. borrowed.** Procedure and agreement score from EdgeWrite 2005 (ref. 33). Good et al. 1984 did it for typed commands. Epps et al. 2006 elicited tabletop hand shapes from Windows screenshots. Nielsen et al. 2004 "describe a similar approach," and the paper does not say how it differs (p. 1084). "First to employ users, rather than principles" (p. 1084) means: first complete, conflict-free set from consensus, for two-handed continuous multi-touch, with animated neutral prompts.

---

## Slide 13 — Strengths

**Time:** 0:45 · **Speaker:** B · **Visual:** four tiles

On the slide:
- **Behaviour before feedback, by design.** No recognizer, so nothing in the session trains the participant.
- **Counts with explanations.** 1,080 logged gestures, timing, ratings, plus think-aloud and four-angle video.
- **Reports what to leave out.** Which commands need widgets. "Useful not just for what it contains, but also for what it omits" (p. 1088).
- **A designer baseline, and a method that travels.** The users' set is checked against the authors' own designs, and nothing in the procedure is tied to tables.

---

## Slide 14 — Limitations: theirs and ours

**Time:** 1:30 · **Speaker:** B · **Visual:** two columns · **[classmate]** slot for one credited critique

On the slide:

**Acknowledged by the authors (p. 1091)**
- Participants could not revise earlier gestures after seeing later commands
- No application, organizational or cultural context; in the paper's words, "children, Eastern, or uneducated participants would behave differently"
- **The set is not validated.** The reverse guessability study and a working recognizer are left as next steps

**Worth adding**
- Grouping into "identical" gestures and taxonomy coding were author judgments. The paper does not say what counted as identical, and **no inter-rater reliability is reported.**
- "Popularity identifies better gestures" (p. 1088) rests on participants rating their own gestures. No independent judges. The authors' "goes a long way to validating" overstates it.
- Every χ² test uses N = 1,074 gestures from 20 people as if they were independent observations (pp. 1087–1088).
- Commands were drawn from desktop and tabletop systems (p. 1085). At least five are desktop concepts: cut, paste, minimize, maximize, task switch.
- 57% coverage also means **43% of what users proposed is not accommodated.**
- Claims "ease of recognition" (p. 1091) without having built a recognizer.
- n = 20, all right-handed, mean age 43. Two authors at Microsoft Research, on their own prototype; the paper does not discuss it.

---

## Slide 15 — Three numbers that are not the same

**Time:** 1:15 · **Speaker:** B · **Visual:** three cards, Eq. 3 from p. 1087

On the slide:

**Agreement, A = 0.32.** How concentrated the consensus is. The paper's own second example (Eq. 3, p. 1087): *select single*, one hand, groups of 11, 3, 3 and 3 → **A = 0.37.** Eleven people agreed; nine did something else. The floor is about 0.05, not 0, because a gesture always "agrees with itself." **"A = 0.32" does not mean "32% agreed."** Vatavu and Wobbrock (CHI 2015) replaced it with a corrected agreement rate for this reason.

**Coverage, 57.0%.** Share of all proposals matched by *any* gesture in the final set, aliases included. Rewards the set for being large.

**Guessability, ?** The chance that a new user guesses the right gesture for a command. **Never measured.** It is what the method is named after and what the proposed validation would test.

---

## Slide 16 — Design decisions that shaped the result

**Time:** 1:30 · **Speaker:** B · **Visual:** four tiles, Figure 4 crop of the four enlarge variants

On the slide:

1. **What counted as "identical"?** The paper does not say (p. 1087). The coding rules (p. 1086) and the Figure 4 caption imply that one to three fingers were merged. If so, that merges groups and raises agreement, and the finding "fingers don't matter" and the scoring rule lean on each other.
2. **Conflict-free because the target is part of the gesture.** The paper's rule: the referent with the larger group wins (p. 1087). It does not say what happened to the other referent. In the set, the same four gestures mean resize on an object, zoom on the background, open on a folder. Our reading: a design move, not an observation. The authors could answer that binding is a taxonomy dimension, so these were never "the same gesture."
3. **Two-hand gestures were required.** Everyone had to produce one for every command, although one hand was preferred for 25 of 27. The jump idiom and 35% of the set came from this. Related to the "production" technique Morris et al. (2014) later proposed against legacy bias: ask for several proposals per command.
4. **Users collapsed categories.** Prompted to zoom in, a participant said "Oh, that's the same as enlarge" (p. 1090). Maximize folded into enlarge (p. 1087). Users paired minimize with shrink, yet minimize kept its own gesture in the set. Participants and authors "did not always regard the same referents as complex": Figure 3's order differs from Table 1's.

---

## Slide 17 — Whose habits did the study capture?

**Time:** 1:45 · **Speaker:** A · **Visual:** the two screening quotes, a four-row table, Figure 4 crop of the pinch · **[classmate]** slot

On the slide:

**Who the paper says sat down.** "Non-technical people without prior experience using touch screens (e.g., the Apple iPhone)" (p. 1084). "No participant had used an interactive tabletop, Apple iPhone, or similar" (p. 1085). No screening instrument and no touch history are reported.

**What the paper blames: the desktop.** "It's a Windows World" (p. 1090): clicking with index and middle finger as on a two-button mouse, imaginary close buttons at the top right, an imaginary taskbar, "I'm a child of the mouse."

**What a 2008 adult had already touched** (dates from general knowledge)

| | Since | Teaches |
|---|---|---|
| ATMs, kiosks, point-of-sale terminals | 1980s | Tap the button |
| Palm PDAs with Graffiti | 1996 | Stylus strokes on a screen |
| Laptop touchpads | late 1990s | Tap, drag |
| Nintendo DS | 2004 | Stylus tap, drag, draw |

"No prior experience using touch screens" is hard to believe for 2008 adults. Every row teaches one finger, one point.

**Three explanations for 72% single-point, which this data cannot separate**
1. The mouse (the paper's reading)
2. A decade of single-point touch interfaces (ours)
3. Plain pointing, the default human gesture. Beringer, Voida et al. and Epps et al. all found pointing dominant (p. 1084)

**The pinch.** It is in the set as one of four resize aliases; 5 of 20 used a one-hand reverse pinch to resize (p. 1091). Prior exposure through a year of iPhone adverts and Han's demo? The study cannot tell.

**What follows.** Elicitation captures the repertoire the participant brings in. The paper reports that as a finding about users. We read it also as a limit of the method, and the problem has grown: today every participant arrives with a smartphone vocabulary. Morris, Wobbrock and colleagues, "Reducing Legacy Bias in Gesture Elicitation Studies," *Interactions* 2014.

---

## Slide 18 — Connections to other research and course topics

**Time:** 1:00 · **Speaker:** A · **Visual:** map or list

**[TODO: Oliver – section 6 of the notes. At least two explicit links to the weeks 1–3 readings, each with a "so what."]**

Hooks already in the paper:
- Furnas et al., the vocabulary problem (ref. 8): people rarely pick the same word for the same thing → aliasing
- Hutchins, Hollan and Norman, gulf of execution (ref. 10): removed on purpose here
- Good et al. 1984: a "user-derived" command-line interface, the same principle for typed commands
- Wobbrock et al. 2005, EdgeWrite guessability: where the procedure and the score come from
- Long et al.: even gesture designers are sometimes poor at picking easily differentiable gestures
- Morris: cooperative gestures for tabletops (same author, cited). This study elicited from single users on a device built for groups.
- Wilson: physics-based interaction on the Surface, UIST 2008 (same lab, previous year). The 44% physical gestures are read as support for it.

---

## Slide 19 — What came next, and what it suggests

**Time:** 1:30 · **Speaker:** B · **Visual:** two columns, "the authors' agenda" vs "what actually followed"

On the slide:

**The authors' agenda (p. 1091)**
- *Next steps:* reverse the study (show gesture videos, ask people to guess the command); build a vision-based recognizer and measure recognition rates
- *Limitations left for later:* let participants see all commands first; add application context; study children and other cultures
- *Technology implications:* sense above the table and beyond the display edge

**What actually followed (all verified 2026-09-30)**
- **Morris, Wobbrock, Wilson, GI 2010.** 22 new participants, blind to authorship, judged 81 gestures. They preferred gestures authored by larger groups, partly because the researchers' own gestures were more physically and conceptually complex.
- **Vatavu and Wobbrock, CHI 2015.** Corrected agreement rate, disagreement and co-agreement rates, a significance test, a toolkit. Honorable Mention.
- **Morris et al., *Interactions* 2014.** Legacy bias named as a weakness of the method. Remedies: several gestures per command, priming with the device's capabilities, eliciting in pairs.
- **Villarreal-Narvaez et al., DIS 2020.** Systematic review of 216 elicitation studies.

**The twist.** The method spread; no product adopted the set as published. Yet drag, double-tap, pinch and spread are now universal. Apple shipped pinch-to-zoom in 2007; this study elicited it a year later from people recruited for having never used a touch screen. Elicitation and the market converged.

**Open today.** Mid-air, AR/VR hand tracking, wearables, voice and LLM command phrasing. Can elicitation still find anything new once every participant carries a smartphone gesture vocabulary?

---

# Part C – Bridges to discussion

## Slide 20 — Our reading of the paper

**Time:** 0:45 · **Speaker:** B · **Visual:** one sentence, three supports

On the slide:

> **The lasting contribution is the method. The evidence in this paper for the method is thinner than its influence.**

- The guessability test the authors proposed was never run. GI 2010 tested preference, which is partial support.
- "Popularity identifies better gestures" rests on participants' own ratings, with no independent judges.
- The study measured the repertoire participants brought in. The paper calls that a finding about users; it is also a property of the method.

---

## Slide 21 — Tensions worth arguing about

**Time:** 1:15 · **Speaker:** A + B · **Visual:** four big questions, three small ones · **[classmate]** slots

On the slide:

**Four questions with more than one defensible answer**
1. **Should the most popular gesture win?** 72% mouse-like, "I'm a child of the mouse." Does consensus from novices lock in old habits? Guessability vs. learnability vs. efficiency.
2. **Did the commands shape the answers?** Cut, paste, minimize, maximize, task switch are desktop concepts. Who should choose the referents?
3. **Does popularity identify better gestures, or gestures people enjoyed inventing?** The ratings are self-reports. What would independent validation look like?
4. **Is behaviour without feedback the right design target?** Real users adapt to a system's responses. Participants gestured in the air and off-screen (p. 1090).

**Three smaller things we found interesting**
- **Floor or ceiling?** 19.1% of each author's gestures were never tried. Does elicitation find the most guessable gesture or the best one?
- **Symbolic gestures import writing.** ✓, ✗, ? are Western marks. The most abstract commands got the most culture-bound gestures; the paper's own caveat says "Eastern" participants would differ.
- **Single users on a group device.** *Minimize* = move to the bottom of the surface "as defined by user's seating position" (Figure 4). A different direction for every person at a table.

---

## Slide 22 — Over to you

**Time:** 1:00 · **Speaker:** A + B · **Visual:** commentary quotes, then the plan for the next 30 minutes

On the slide:

**What you wrote** **[TODO: after reading the commentaries – three to four recurring themes or sharp observations, each credited by name. Pair each with the tension on slide 21 it speaks to.]**

**How the next 30 minutes run**
1. **Warm-up, 5 min: a live elicitation.** We show the effect of two commands on screen (one from the top of Figure 3, one from the tail). On three, everyone performs the gesture they would use on the desk. We count the groups and compute agreement on the spot.
2. **Pairs, 5 min:** pick one of the four questions, take opposite sides.
3. **Full group, 20 min:** we follow what comes up.

Footer: Discussion plan with prompts, follow-ups and timing is in **[TODO: Week 4 - Discussion Plan, to be written]**.
