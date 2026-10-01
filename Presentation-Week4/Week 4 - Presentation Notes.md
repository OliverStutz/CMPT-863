# Week 4: User-Defined Gestures for Surface Computing

Wobbrock, J. O., Morris, M. R., and Wilson, A. D. (2009). User-Defined Gestures for Surface Computing. *Proceedings of CHI 2009*, Boston, pp. 1083–1092.
Authors: Wobbrock (University of Washington, Information School), Morris and Wilson (Microsoft Research). Page numbers below are the printed ones.

---

## Background: the minimum needed to follow the paper

### What surface computing was in 2009
- **Interactive tabletop.** A table-sized horizontal display that senses touch, usually with a camera looking up from underneath. Systems the paper cites: DigitalDesk (1993), DiamondTouch (2001), SmartSkin (2002), and the Microsoft Surface prototype used in this study.
- **Multi-touch.** Several fingers or both hands can touch at once. There is no single cursor.
- **Why gestures at all.** On a table, keyboard, mouse and mouse-style widgets are "no longer preferable" (p. 1083). Commands have to be given by hand. A mouse constrains motion to one point; hands do not, so "almost anything one can do with one's hands could be a potential gesture" (p. 1083).
- **How gestures were chosen before this paper.** System builders invented them, often so a recognizer could tell them apart, for example separate meanings for 1-, 2-, 3- and 5-finger touches (p. 1085). Users only ever tested gestures that had already been designed.

### Terms the paper uses
- **Referent** = the command or effect (delete, zoom in). **Sign** = the gesture that causes it. The words come from McNeill's work on human gesture (p. 1085, footnote).
- **Guessability study.** Show the referent, ask the participant to invent the sign. Measures how well an input can be guessed by someone who has never used the system. Comes from the authors' 2005 EdgeWrite work (ref. 33).
- **Elicitation.** Collecting the gestures users produce, instead of testing gestures a designer made.
- **Agreement score, A.** For one command, take each group of participants who made the same gesture, square the group's share of all participants, and add the squares. Then average over all commands. 1.0 means everyone did the same thing. The floor is 1/20 = 0.05 when all 20 people differ, so it is not a percentage (p. 1087).
- **Think-aloud protocol.** Participants say what they are thinking while they act. This is how the paper gets at mental models.
- **Gulf of execution.** The gap between what a user wants to do and what the system accepts (Hutchins, Hollan and Norman, ref. 10). The study removed it by recognizing nothing, so every gesture was "acceptable" (p. 1085).
- **Aliasing.** Giving one command several gestures. Known to raise guessability because people rarely pick the same word or sign for the same thing (Furnas's vocabulary problem, ref. 8).
- **Widget.** An on-screen button, menu or bar. The alternative to a gesture when people cannot agree on one.
- **Taxonomy dimensions** (Table 2, p. 1086). *Form*: what the hand does (static pose, dynamic pose, pose with a path, one-point touch, one-point path). *Nature*: symbolic (draws a symbol), physical (acts on objects as if real), metaphorical (acts like something else, e.g. a magnifying glass), abstract (arbitrary mapping). *Binding*: where the gesture must happen (on the object, relative to the world, anywhere, or mixed). *Flow*: discrete (system responds after the gesture) or continuous (responds during it).

### Timing (outside the paper, for context)
- Before 2007: single-touch tapping was already everyday experience. ATMs, ticket and check-in kiosks, point-of-sale terminals, stylus PDAs (Palm, from the mid-1990s), Nintendo DS (2004), portable GPS units, and in-car touchscreens in navigation systems (the first dates to 1986). These were tap and drag interfaces, not multi-touch. Full list in section 5, "Gesture input before 2007".
- 2006: Jeff Han's multi-touch demo made the idea widely known.
- 2007: iPhone shipped with pinch-to-zoom. Microsoft Surface announced.
- 2008: Microsoft Surface sold commercially.
- 2009: this paper (April). Windows 7 with built-in multi-touch shipped in October.
Gesture vocabularies were being fixed in products at exactly this moment. Two of the three authors worked at Microsoft Research and the study ran on a Surface prototype.

---

## 1. What problem does the paper address, and why does it matter?

**Stated problem.** "To date, most surface gestures have been defined by system designers," giving "somewhat arbitrary gesture sets whose members may be chosen out of concern for reliable recognition" (p. 1083). That criterion is fine for prototypes but "not useful for determining which gestures match those that would be chosen by users."

**Deeper problem.** The field built gesture vocabularies from the recognizer outward, not from the user inward, and had no systematic, measurable way to find out what people would do on their own. The related work is arranged to show this: RoomPlanner and Wu et al.'s gesture framework "were not elicited from users, although usability studies were conducted" (p. 1085). Long et al. had already shown designers are poor at picking gestures that are easy to tell apart (p. 1084).

**Why it mattered then.** Tabletops were spreading to single users, groups, multi-display rooms and entertainment (p. 1083). Tables in public spaces need "immediate usability" (p. 1085). The authors screened out anyone who had used an iPhone (p. 1084), a sign they knew consumer touch was about to set expectations. Commercial gesture conventions were being decided by engineers at that moment.

**Why it still matters.** The problem appears whenever a new input channel has no established vocabulary: mid-air gestures (Kinect, Leap Motion), hand tracking in AR and VR headsets, smartwatch and ring gestures, in-car gestures, voice command phrasing. This paper started the genre of "gesture elicitation studies," which now numbers in the hundreds.

**Slide phrasing.**
- "Gesture sets were designed for the recognizer, not the person. Nobody had asked what people would actually do."
- "When anything your hands can do could be a command, who decides which gesture means what?"

---

## 2. What is the paper's main idea or contribution?

**Core idea in one sentence.** Show non-technical users the effect of a command, ask them to perform the gesture that would cause it, and let agreement across many users decide the gesture set.

**Why it is built this way.**
- *Elicit from users* because users are the ones who have to guess. Designers organize gestures "in a principled, logical fashion," but "user behavior is rarely so systematic" (p. 1083).
- *Show the effect first* so the meaning is fixed and only the form is open. Everyone answers the same question, so gestures can be compared and counted. Prompts used a neutral world of 2D blocks with no Windows or Mac elements, because Epps et al.'s Windows screenshots had pushed people toward mouse emulation (p. 1084, 1085).
- *Remove the recognizer* so recognition failures cannot reshape behaviour. The study becomes "a monologue in which the user's behavior is always acceptable" and shows "unrevised behavior" (p. 1085).

**Contributions as the authors list them (p. 1084).**
1. A quantitative and qualitative characterization of user-defined gestures, including a taxonomy.
2. A user-defined gesture set.
3. Insight into users' mental models.
4. Implications for surface technology and interface design.

**Sorted by kind.**
- *Method*: effect-then-cause elicitation with no recognizer, then conflict resolution by largest group.
- *Artifact*: 48 gestures for 22 commands, conflict-free, covering 57.0% of everything proposed (p. 1087–1088).
- *Instruments*: the agreement score and the four-dimension taxonomy.
- *Findings*: see section 4.

**What matters most.** The method. The gesture set was never validated and is left as "an important next step" (p. 1091). The taxonomy is situational. The procedure, and agreement as the thing to measure, is what later work copied. Note that the authors' own list of four contributions does not include the method.

**Novel versus borrowed.** The referent-to-sign procedure and the agreement score both come from the authors' 2005 EdgeWrite work (ref. 33), which covered single-finger strokes. Good et al. (1984) used the same principle for a typed email interface. Epps et al. (2006) elicited tabletop hand shapes from static Windows images. So "the first to employ users, rather than principles" (p. 1084) really means first to build a complete, conflict-free set from user consensus, for two-handed, continuous, multi-touch input, with animated neutral prompts. The paper does not say how it differs from Nielsen et al. (2004), so do not overclaim.

**The one finding to remember.** Before the study, each author designed their own gestures for the 27 commands. Each author alone matched 43.5% of the user-defined set. All three combined matched 60.9%. And 19.1% of each author's gestures were never tried by any participant (p. 1090). Expert designers, including the people running the study, could not predict what users would do.

---

## 3. How was the research conducted?

| Element | Detail |
|---|---|
| Participants | 20 paid volunteers from the general public, 9 women, mean age 43.2 (sd 15.6), all right-handed. None had used a tabletop, iPhone or similar. No computer scientists or designers. Occupations included restaurant host, musician, author, steelworker (p. 1085). |
| Apparatus | Microsoft Surface prototype, 24 by 18 inches, 1024 by 768. Custom C# software played a recorded voice and an animation for each command, then reset the scene and waited (p. 1085). |
| Commands | 27 referents drawn from desktop and tabletop systems, from "move a little" to "undo" (Table 1). The three authors rated each one's conceptual complexity from 1 to 5 beforehand. |
| Procedure | Commands in random order. For each, the participant made a 1-hand gesture and a 2-hand gesture while thinking aloud, said which they preferred, and rated each gesture on two 7-point scales: "good match for its purpose" and "easy to perform" (p. 1086, 1088). |
| Data | 20 x 27 x 2 = 1080 gestures, 6 discarded for confusion. All touches logged as ovals with millisecond timestamps. Video from four angles. Two authors took notes on the think-aloud. |
| Analysis | Authors manually coded every gesture on the four taxonomy dimensions. Gestures for each command were grouped into identical sets and the agreement score computed. The largest group won each command; where one gesture would have served two commands, the command with the larger group kept it (p. 1087). Statistics: logistic regression and correlations against the complexity ratings. |
| Control comparison | Before the study, the three authors independently designed their own gesture sets for the same 27 commands (p. 1090). |

**What the system did not do.** It never recognized a gesture or gave feedback. That is a deliberate design choice, not a limitation of the hardware.

**The screening criterion, worth scrutinizing.** The only stated exclusion is that "no participant had used an interactive tabletop, Apple iPhone, or similar" (p. 1085). Nothing is said about ATMs, kiosks, PDAs, handheld game consoles, GPS units or in-car screens, and nothing about having *seen* multi-touch in adverts or online. The paper reports no touch history for the participants at all. See the highlight in section 5.

---

## 4. What are the most important findings or takeaways?

**About the gestures people make**
- Finger count rarely matters. 13 of 20 varied how many fingers they used; only 2 said it meant anything. One to three fingers read as "a single point," a whole palm as something more, four fingers as a grey area (p. 1090). This contradicts several earlier systems.
- One hand beats two. One hand was preferred for 25 of 27 commands, and two hands were never preferred (p. 1088).
- The desktop is deeply rooted. About 72% of gestures were mouse-like single-point touches or paths, even with no Windows imagery on screen. People tapped imaginary close buttons and taskbars. "I'm a child of the mouse" (p. 1090).
- Simple commands get physical gestures; complex ones get metaphorical or symbolic gestures (p. 1087).
- People imagine space beyond the screen edge, dragging off-screen to delete and in from off-screen to paste, and some gestured in the air above the table (p. 1090).
- Opposite commands get reversible gestures, even though the pairs were never shown together (p. 1090).

**About agreement**
- Overall agreement was 0.32 for one hand and 0.28 for two (p. 1087).
- Agreement falls as command complexity rises (r = -0.52). Move commands had perfect agreement; undo, help, task switch and cut sat near 0.10 to 0.15 (Figure 3).
- Gestures that more people chose were also rated better by the people who chose them. The authors take this as validating the approach (p. 1088).
- Longer planning time went with lower goodness ratings: good gestures came to mind quickly (p. 1088).

**The resulting set (Figure 4)**
- 48 gestures for 22 commands, 31 one-handed and 17 two-handed. Ten commands have one gesture, the rest have two to five aliases.
- The set reuses gestures and disambiguates by target: spreading fingers on an object enlarges it, on the background zooms in, on a folder opens it.
- Four commands got no gesture. Insert and maximize reused drag and enlarge. Close and task switch relied on imaginary widgets.

**Implications the authors draw (p. 1091)**
- Do not distinguish gestures by finger count. Treat four fingers as the boundary between a point and a hand.
- Give low-agreement commands (menu, help, undo, accept, reject, zoom) a widget, not only a gesture.
- Reuse gestures across related commands and disambiguate by where they land.
- A physics engine handles the 44% of gestures that were physical, but not the rest.
- Sense a little beyond the display edge and above the table.
- Avoid demands for precise positioning: only 2 of 14 people resized along the diagonal.

---

## 5. What are the paper's strengths and limitations?

**Strengths**
- Removes recognizer bias by design, so the data shows what people do before a system trains them.
- Neutral blocks world avoids the Windows bias that undermined Epps et al.
- Combines counts (1080 logged gestures, timing, ratings) with think-aloud and video, so the numbers come with explanations.
- Reports what to leave out: which commands need widgets. "Useful not just for what it contains, but also for what it omits" (p. 1088).
- Tests its own premise by comparing the users' set with the authors' own designs.
- The method is reusable in any modality, and the set has reversibility, reuse and aliasing built in.

**Limitations the authors acknowledge (p. 1091)**
- Participants could not revise earlier gestures after seeing later commands.
- No application context, and no organizational or cultural context.
- All participants were literate American adults; children and other cultures would differ.
- The set is not validated. The reverse study (show the gesture, guess the command) and a working recognizer are left as next steps.

**Limitations worth adding**
- Grouping gestures into "identical" clusters and coding the taxonomy were author judgments with no inter-rater reliability reported.
- Goodness ratings were given by participants for gestures they had just invented, so "popularity identifies better gestures" is partly circular.
- Consensus among novices may lock in legacy habits (72% mouse-like). Pinch still appeared even with iPhone users screened out. See the highlight below.
- The 27 commands came from desktop systems (cut, paste, minimize, task switch), which may have invited desktop thinking.
- 57% coverage also means 43% of what users proposed is not accommodated.
- The agreement score's floor is 0.05, not 0, so scores are not percentages. Later work (Vatavu and Wobbrock, CHI 2015) rescaled it.
- The paper claims "ease of recognition" for the set (p. 1091) without having built a recognizer.
- Small sample: 20 people, all right-handed, mean age 43. Two authors had a commercial stake through Microsoft.
- Minor: the discussion cites Figure 2 for the agreement cutoff after seven commands; it means Figure 3.

### Highlight: an unbiased participant did not exist in 2008

The paper's whole design rests on capturing "unrevised behavior" from people whose hands have not been trained by a touch system. It treats the threat to that as the desktop: the bias section is titled "It's a Windows World" and the mouse gets the blame for the 72% of single-point gestures (p. 1090). The screening question matches that view: no tabletop, no iPhone.

But by 2008 the mouse was not the only prior. Tap-and-drag touch had been ordinary for over a decade through ATMs, kiosks, point-of-sale terminals, stylus PDAs, the Nintendo DS (2004) and GPS units. Every one of those teaches the same lesson the mouse does: one finger, one point, tap the button. So the 72% figure cannot be read as "desktop habit" alone. It is the sum of all single-point interfaces people had used, and the paper never asked which ones.

Gesture literacy came through media as well as ownership. The iPhone had been advertised with pinch-to-zoom for over a year, and Jeff Han's 2006 multi-touch demo was widely circulated. Screening out iPhone owners removes the device, not the exposure. The set that came out of the study contains a pinch and a reverse pinch (Figure 4, p. 1089 and p. 1091). Where did those come from, if not from having seen them?

Two things follow for the class.
- The method's promise of unrevised behaviour is weaker than it sounds. Elicitation captures whatever repertoire the participant walks in with. The paper's own finding that legacy dominates is also a finding about the limits of its method.
- The problem has only grown. In 2008 the prior was tap and drag. Today every participant arrives with a full smartphone gesture vocabulary, so elicitation for new modalities (mid-air, VR, wearables) is swamped by touchscreen habits. This became its own research thread: Morris, Danielescu, Drucker, Fisher, Lee, schraefel and Wobbrock, "Reducing Legacy Bias in Gesture Elicitation Studies," *Interactions*, May–June 2014, propose asking for several gestures per command, priming participants with the new device's capabilities, and eliciting in pairs.

A correction to a claim that sounds plausible: production cars did not have gesture control in 2008. In-car touchscreens existed, but camera-based gesture control in a production car first shipped in the BMW 7 Series in 2015.

### Gesture input before 2007: what a participant could already have used

Almost all of it taught a single point of contact or a whole-device motion. Multi-finger gestures existed only in niche products. Grouped by how many people would have met them.

**Mass-market, likely in a 2008 participant's hands**

| Product | Date | Gesture |
|---|---|---|
| Palm PDAs with Graffiti | 1996 | Unistroke letters drawn with a stylus. Direct ancestor of the unistroke work (Goldberg and Richardson 1993, EdgeWrite 2003) behind this paper's guessability method. |
| Apple Newton | 1993 | Handwriting recognition, scrub-out stroke to delete. |
| Laptop touchpads (Synaptics, Alps) | late 1990s | Tap to click, edge scrolling. Two-finger scrolling on Apple laptops from 2005. |
| iPod touch wheel, then click wheel | 2002, 2004 | Circular finger stroke on a capacitive surface. Tens of millions of units. |
| Nintendo DS | 2004 | Stylus drag, flick and draw in games. |
| Sony EyeToy for PlayStation 2 | 2003 | Camera tracks arm and body movement. About ten million sold. |
| Nintendo Wii | Nov 2006 | Swing, point and shake with the remote. Enormous by 2008. |
| SMART Board whiteboards | 1991 | Finger or pen touch in classrooms, single point. |
| ATMs, kiosks, point-of-sale, GPS units | 1980s onward | Tap only. |
| In-car touchscreens | 1986 onward | Tap only, in navigation systems. No production gesture control until BMW, 2015. |

**Niche but real multi-touch**
- FingerWorks TouchStream keyboard and iGesture Pad, about 2001 to 2005. Ten-finger tracking with pinch, swipe and rotate. Sold to enthusiasts and people with repetitive strain injury. Apple bought the company in 2005 and the same engineers built the iPhone's gestures.
- JazzMutant Lemur, 2005. Multi-touch controller for musicians.
- Windows Tablet PCs, 2002. Pen gestures, with flicks added in Vista in 2006.
- Mouse gestures in the Opera browser from 2001 and in games such as Black & White (2001) and Okami (2006).
- Samsung SCH-S310, Korea, 2005. Write a digit in the air with the phone to dial it, shake twice to hang up. First accelerometer-gesture phone, never sold outside Korea.

**Research the paper cites or descends from**
- Sketchpad light pen, 1963. Krueger's Videoplace camera interaction, 1970s. Bolt's "Put That There" pointing plus speech, 1980. Buxton's multi-touch tablet, 1985. DataGlove 1987 and Nintendo Power Glove 1989. Charade 1993 (cited). DiamondTouch 2001 and SmartSkin 2002 (both cited). Han's FTIR demo, 2005 to 2006.

**What this does to the argument.** A 43-year-old American in 2008 had very probably used an iPod wheel, a laptop touchpad and an ATM, and quite possibly a Wii or a Palm. Every one of those teaches a single point of contact or a whole-hand motion. None teaches multi-finger gestures. That matches the paper's data almost exactly: 72% single-point touches, many whole-hand physical gestures, and pinch appearing only as one alias among four for enlarge. The pre-2007 legacy did not come from the mouse alone. It came from a decade of consumer devices that all reward the same one-finger behaviour, and the paper never asked about any of them.

Sources checked 2026-09-29: Samsung SCH-S310 announcement coverage (photographyblog.com, gearlive.com); FingerWorks acquisition coverage (dealroom.co, appleinsider.com); BMW 7 Series gesture control coverage (bmwblog.com, digitaltrends.com). Other dates are from general knowledge; confirm before putting on a slide.

---

## 6. How does it connect to other research or topics we have discussed?

*(To be filled in.)*

---

## 7. What new questions or research directions does it suggest?

**Named by the authors (p. 1091)**
- Validate the set by reversing the study: show unlabeled video of user-defined and designer-defined gestures and see whether new people can guess the command.
- Implement the set with a vision-based recognizer and measure recognition rates.
- Let participants see all commands first, then choose gestures in their own order.
- Repeat with application context, and with children, other cultures and less educated participants.
- Sense above the table and beyond the display edge.

**Questions the paper raises without answering**
- Should the most popular gesture win? Guessability favours what novices already know. Learnability, efficiency or expert performance might favour different gestures. Which goal should a gesture set serve?
- Is unrevised behaviour the right design target when real users adapt to feedback within minutes?
- Who should choose the referents, and does the command list itself carry a desktop bias?
- Does agreement predict anything about memorability or performance in actual use, over weeks rather than minutes?
- How should a physics engine and a symbolic recognizer be combined, given that gestures were 44% physical and 56% other?
- Does the method transfer to input with no surface at all (mid-air, VR, wearables), where the design space is even larger?
- Can elicitation still find anything new once every participant carries a smartphone gesture vocabulary? Is the "unbiased user" a fiction, and if so, what should the method aim for instead?
- Where does the four-finger boundary come from, and does it hold across hand sizes and cultures?

**Follow-up that exists (all verified 2026-09-30)**
- Morris, Wobbrock and Wilson, "Understanding Users' Preferences for Surface Gestures," Graphics Interface 2010, pp. 261–268: 22 participants, blind to authorship, judged 81 gestures on a Surface. They preferred gestures authored by larger groups of people (end-user elicitation, or more than one researcher), partly because the researchers' own gestures were more physically and conceptually complex. Do not shorten this to "users preferred user-authored gestures."
- Vatavu and Wobbrock, "Formalizing Agreement Analysis for Elicitation Studies: New Measures, Significance Test, and Toolkit," CHI 2015: agreement rate, disagreement and coagreement rates, a significance test and a toolkit. Honorable Mention.
- Morris et al., "Reducing Legacy Bias in Gesture Elicitation Studies," *Interactions* 2014: the legacy-bias finding of this paper turned into a known weakness of the method, with proposed fixes.
- Villarreal-Narvaez, Vanderdonckt, Vatavu and Wobbrock, "A Systematic Review of Gesture Elicitation Studies: What Can We Learn from 216 Studies?", DIS 2020: the source for "hundreds of elicitation studies."

Slides built from these notes: [[Week 4 - Slides]] and [[Week 4 - Speaker Notes]].
