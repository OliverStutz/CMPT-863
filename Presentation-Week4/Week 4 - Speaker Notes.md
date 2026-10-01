# Week 4 – Speaker Notes

Spoken notes for each slide in [[Week 4 - Slides]]. Written to be said aloud at about 130 words per minute; each block is kept within its slot. Each block has: **Say** (the script), **Transition** (the sentence that leads to the next slide), **Watch out** (things that are easy to get wrong), **Source** (where the claim lives). Page numbers are the printed ones (1083–1092). Background on every term is in [[Week 4 - Presentation Notes]].

Revision 2 (2026-09-30): all fixes from `reports/slides-review-cross-examination.md` applied. Spoken total about 3,100 words for 25:45 of slides; every block checked against its slot.

Ground rules for the whole talk, from the course brief:
- Everyone has read the paper. Part A is a fast shared map, not a summary. If the room nods, speed up.
- Say "we think" when it is our judgment, "the paper says" when it is theirs. Slides 11–21 mix both.
- Four things never to say in class: that the gesture set was validated (it was not, p. 1091); that this paper invented the agreement score (EdgeWrite 2005 did, ref. 33); that an agreement of 0.32 means 32% of people agreed (the floor is about 0.05); and that the paper promised participants with no prior habits ("unrevised behavior" on p. 1085 means not revised by recognizer feedback, nothing more).
- Classmates' commentaries go in the **[classmate]** slots on slides 14, 17 and 21, and on slide 22. Credit by name.

---

# Part A – Rundown

## Slide 1 — Title (0:30)

**Say:** Good morning. This week's paper is Wobbrock, Morris and Wilson, "User-Defined Gestures for Surface Computing," CHI 2009. Wobbrock was at the University of Washington; Morris and Wilson at Microsoft Research, which matters later. Ten minutes of shared map, twelve on the questions the course asks, three to set up the discussion. The first part is quick, because you have read it.

**Transition:** Here is the whole paper on one slide.

**Watch out:** Fill in the date and co-presenter before submission.

---

## Slide 2 — The paper in one slide (1:00)

**Say:** The question is simple: put someone in front of a touch table, never tell them what gesture to use, and watch what they do. Twenty adults with no technical background sat at a Microsoft Surface prototype. For each of 27 commands the table showed the effect, say a block disappearing, and asked for the gesture that would have caused it. One hand and two hands, thinking aloud. The table never recognized anything; that is the key design decision and we come back to it. Out came 1,080 gestures and three things: a taxonomy of what people do on a surface, an agreement score per command, and a set of 48 gestures covering 22 of the 27 commands.

**Transition:** Why 2008 and not 2004 or 2012?

**Watch out:** Four commands got no gesture of their own (insert, maximize, task switch, close, p. 1087); "22 of 27" is the paper's count with the two move commands merged. 1,080 is before six discards.

**Source:** p. 1083 abstract, p. 1085–1086 procedure, p. 1087–1088 set.

---

## Slide 3 — Why 2009 was the moment (0:45)

**Say:** Jeff Han's multi-touch demo went around the internet in 2006. The iPhone shipped in 2007 with pinch-to-zoom; Microsoft announced Surface the same year and sold it from 2008. The paper does not date the study, but it had to be done by the autumn 2008 deadline. Windows 7 shipped with multi-touch that October. So gesture vocabularies were being decided inside product teams at exactly this moment, and two of the three authors worked at Microsoft Research, on their own table.

**Transition:** And the way those vocabularies were being decided is the problem.

**Watch out:** Dates outside the paper are general knowledge. Say the authors had a stake, not that they wanted to sell Surface.

**Source:** p. 1083 ("timely"), p. 1085 (public spaces, Surface prototype).

---

## Slide 4 — The problem (1:00)

**Say:** In the authors' words, most surface gestures were defined by system designers, giving "somewhat arbitrary gesture sets whose members may be chosen out of concern for reliable recognition." Malik and colleagues gave different meanings to one, two, three and five fingers. A camera counts fingers easily; whether a person thinks in finger counts, nobody had asked. Long and colleagues had shown that even people designing gestures are sometimes poor at picking ones that are easy to tell apart. And the space is huge: a mouse pins you to one point, your hands can do almost anything. Add one designer's guess and you have Furnas's vocabulary problem. So: gesture sets designed for the recognizer, not the person. A few studies had asked users, but nobody had built a full multi-touch set from the answers.

**Transition:** The paper's answer is to turn the process around.

**Watch out:** Keep "sometimes" in the Long et al. line; the paper says "users are sometimes poor." The paper hedges Malik's finger distinctions as "may be problematic."

**Source:** p. 1083 quote; p. 1085 Malik; p. 1084 Long et al. and the prior elicitation work; ref. 8 Furnas.

---

## Slide 5 — The core idea: referent → sign (1:30)

**Say:** Instead of designing a gesture and testing whether users can learn it, show users the effect and watch what they do. The effect is the referent, the gesture is the sign; the terms come from McNeill. Do this with twenty people and twenty-seven commands, group identical gestures, and let agreement decide the set.

Three choices make this work. Elicit from users, because they are the ones who will have to guess; designers organize gestures logically, "user behavior is rarely so systematic." Show the effect first, in a neutral world of plain blocks, so everyone answers the same question, and because Epps and colleagues had conceded that their Windows screenshots may have pushed people toward the mouse. And no recognizer. The authors call it a monologue in which the user's behaviour is always acceptable. Norman's gulf of execution is removed on purpose, so the system never trains the participant during the session. They call that "unrevised behavior." Note what it means: not revised by feedback. It says nothing about what people bring in with them, and that distinction matters in the second half.

**Transition:** Here is how it ran.

**Watch out:** The procedure and the score are borrowed from the authors' 2005 EdgeWrite paper; say so.

**Source:** p. 1083 (systematic quote), p. 1084–1085 (Epps "may have biased", blocks world), p. 1085 (monologue, unrevised behavior, McNeill footnote), ref. 33.

---

## Slide 6 — Method at a glance (0:45)

**Say:** Twenty paid volunteers from the general public: nine women, average age 43, all right-handed, no tabletop or iPhone experience, no computer scientists. A restaurant host, a musician, a steelworker. A Surface prototype played a voice and an animation for each command, then reset. Twenty-seven commands from desktop and tabletop systems, each pre-rated for complexity by the three authors. For every command: a one-hand gesture, a two-hand gesture, think aloud, a preference, two ratings. Everything logged and filmed. And a baseline: before the study, each author designed their own set for the same commands.

**Transition:** So, 1,080 gestures. How do you turn that into a set?

**Watch out:** The paper does not state whether one-hand always came before two-hand. The authors' sets are a baseline, not a control condition.

**Source:** p. 1085–1086; designer baseline p. 1090.

---

## Slide 7 — From 1,080 gestures to a gesture set (1:30)

**Say:** Three steps. One: for each command, the authors sort the twenty gestures into groups they judge identical. The paper does not say what counted as identical; the caption of Figure 4 implies that one, two or three fingers were treated as the same thing. Hold that thought. Two: the agreement score. For one command, take each group's share of the twenty people, square it, add them up, then average over commands. The paper's own example, "move a little" with two hands: groups of twelve, three, three and two give 0.42. Everyone identical scores one; everyone different scores one over twenty, 0.05. Three: the largest group's gesture wins the command. If the same gesture would win two commands, the command with the larger group keeps it. What happened to the other command the paper does not say; in the final set the same gesture comes back bound to a different target. Result: 48 gestures, 22 commands, no conflicts, covering 57 percent of everything proposed.

**Transition:** Before the results, what did the gestures look like?

**Watch out:** The 0.42 is the two-hand score; one-hand "move a little" is 1.0 on Figure 3. The range of A is 1/|Pr| to 1; the −1 in the paper is an exponent.

**Source:** p. 1087 equations and examples; p. 1087–1088 set and coverage; Figure 4 caption p. 1089.

---

## Slide 8 — Taxonomy: what people actually did (0:45)

**Say:** Every gesture was coded on four dimensions: form, what the hand does; nature, whether it draws a symbol, acts physically, works as a metaphor, or is arbitrary; binding, where it has to happen; flow, whether the system responds after or during. The number to remember from Figure 2: one-point paths plus one-point touches are 72 percent of gestures, for either hand. That is the "mouse-like" figure the paper keeps returning to. Physical gestures are 44 percent, symbolic only 10, and 70 percent are bound to an object.

**Transition:** Now the result the paper hangs on: where did people agree?

**Watch out:** Figure 2 reports form per hand; the slide shows the right hand. Exact values: .450, .274, .439, .702.

**Source:** Table 2 p. 1086; Figure 2 p. 1087.

---

## Slide 9 — Results: where people agreed (1:30)

**Say:** Figure 3, commands sorted by one-hand agreement. "Move a little" and "move a lot" are at one point zero: a single group each. Then a cliff: duplicate at about 0.5, rotate, select group and insert near 0.48, select single at 0.37. After the first seven commands nothing is above about 0.3, and the tail falls steadily from task switch at 0.21 to cut at 0.11. The two-hand line is noisier: it peaks for insert, maximize and enlarge, and dips for pan, close and task switch. Overall agreement is 0.32 for one hand and 0.28 for two, and it falls as the authors' complexity rating rises, correlation minus 0.52.

The authors draw two recommendations. Beyond the first seven commands, give the command a widget as well as a gesture. And because enough people tapped buttons that were not there, consider widgets for delete, zoom, accept, reject, menu and help.

**Transition:** Here is what won.

**Watch out:** Values are read from the chart; say "about." The Discussion cites "Figure 2" for the cutoff; it means Figure 3. Do not put undo on the widget list.

**Source:** Figure 3 p. 1088; A values and r p. 1087; recommendations p. 1091.

---

## Slide 10 — The user-defined gesture set (1:00)

**Say:** Forty-eight gestures, twenty-two commands, thirty-one one-handed and seventeen two-handed. The set works by reuse, with the target deciding the meaning: the four resize gestures on an object resize it, on the background they zoom, on a folder they open it. Drag an object off-screen to delete it, drag a dialog off-screen to reject it. Two idioms are worth noticing. The two-hand "jump": hold the object, tap where it should go, reused for delete, paste and minimize. And drawn marks: a check, an X, a question mark, a scratch-out for undo. Four commands are missing: insert and maximize reuse other gestures, and for close and task switch people tapped buttons that did not exist. One design implication: do not demand precision. Of fourteen people enlarging with two hands, only two pulled along the diagonal.

**Transition:** That is the map. Now, what matters in it?

**Watch out:** "Hold and tap" is Figure 4's label for select group; the reused idiom is Move₂ "jump." Menu is "pull out" from a held object, not from the screen edge. Never say the set was tested with users.

**Source:** Figure 4 p. 1089; counts p. 1088; missing commands p. 1087; diagonal p. 1091.

---

# Part B – Key questions

## Slide 11 — Findings, ranked, and the one to remember (1:45)

**Say:** Start with the three numbers at the top. Before the study, each author designed the gestures they thought users would make. Each author's set matched 43.5 percent of what users eventually defined; all three together, 60.9 percent; and 19.1 percent of each author's gestures were never tried by anyone. The authors joke they were either thinking creatively or hopelessly lost. We think it is the finding to remember, with one caveat: the yardstick is the users' set, so designers cannot win by construction, and the paper itself concedes three experts cannot match the breadth of twenty people. It measures overlap, not quality.

Then the rest, ranked by how solid we think they are. Finger count rarely matters: thirteen of twenty varied it, two said it meant anything; robust, though the paper never says how fingers were handled when grouping. One hand beats two for 25 of 27 commands; robust, though every two-hand gesture was required. Seventy-two percent mouse-like with no desktop on screen; a robust number with an open interpretation, which is slide 17. Agreement falls with complexity; the complexity was rated by the authors, but it also predicts how long people took to start, so it tracks something real. And weakest: popular gestures were rated better by the people who made them.

**Transition:** So what did the paper contribute?

**Watch out:** Do not say "the designers could not predict users." Say "covered under half." Nothing about how these notes were produced belongs in the room.

**Source:** p. 1090 (designers, fingers, mouse), p. 1088 (hands, ratings, r = .71, 8 s vs 15 s, 3.4 s vs 2.0 s), p. 1087 (r = −.52).

---

## Slide 12 — What is the contribution, really? (1:30)

**Say:** The authors list four contributions: a characterization with a taxonomy, the gesture set, mental-model insight, and design implications. Sort them by kind and you get a method, two instruments, an artifact and findings. Which lasted? Not the set: the guessability test the authors proposed was never run, and no product adopted the set as published, though many of its members are now standard. What spread was the method and the habit of measuring agreement. By 2020 a systematic review covered 216 gesture elicitation studies. The numbered list leaves the method out because the authors credit it to their own 2005 work; the abstract still opens with it.

On novelty, be careful. The procedure and the score come from the 2005 EdgeWrite paper on single-finger strokes. Good and colleagues did the same for typed email commands in 1984. Epps elicited tabletop hand shapes in 2006. Nielsen described "a similar approach" in 2004, and the paper does not say how it differs. So "first to employ users rather than principles" means: first complete, conflict-free set from consensus, for two-handed continuous multi-touch, with animated neutral prompts. A real step, narrower than the sentence sounds.

**Transition:** With that framing, strengths and limitations.

**Watch out:** Do not claim a difference from Nielsen. Say "216 studies reviewed," not "descended from this paper."

**Source:** p. 1083 abstract, p. 1084 (contributions, prior work), p. 1087 (ref. 33), p. 1091 (next step). GI 2010 and DIS 2020 verified 2026-09-30.

---

## Slide 13 — Strengths (0:45)

**Say:** Four things the paper does well. It removes the recognizer, so nothing in the session trains the participant. It pairs counts with explanations: logged gestures, timing and ratings next to think-aloud and video. It tells you what to leave out, which commands need widgets; "useful not just for what it contains, but also for what it omits." And it checks the users' set against a designer baseline, with a procedure that nothing ties to tables.

**Transition:** Now the other column.

**Watch out:** Do not say the blocks world "avoided" the Windows bias. The desktop got in anyway (p. 1090), which is itself a finding.

**Source:** p. 1085, p. 1088, p. 1090.

---

## Slide 14 — Limitations: theirs and ours (1:30)

**Say:** The authors are candid. Participants could not go back and change an earlier gesture. There was no application or cultural context; their words are that children, Eastern or uneducated participants would behave differently. And the set is not validated.

We would add a few. The grouping into identical gestures and the taxonomy coding were done by the authors, with no inter-rater reliability, and the paper never says what counted as identical. The claim that popularity identifies better gestures rests on people rating their own inventions, with no independent judges, and the authors say it "goes a long way to validating" the approach, which is more than it carries. Every chi-square test treats 1,074 gestures from twenty people as independent observations. The commands came from desktop and tabletop systems, and at least five are desktop concepts: cut, paste, minimize, maximize, task switch. Fifty-seven percent coverage means 43 percent of proposals are not in the set. The paper claims ease of recognition without a recognizer. Twenty right-handed people, and two authors at Microsoft on their own prototype, which the paper does not discuss. **[classmate]**

**Transition:** Three of those come from mixing up three numbers.

**Watch out:** Not "circular": participants did not know the group sizes when rating. Keep the tone fair; the authors flag most of the big ones.

**Source:** p. 1091 (acknowledged limits, Eastern, ease of recognition), p. 1087–1088 (grouping, χ² with N = 1074, ratings), p. 1085 (referents, sample).

---

## Slide 15 — Three numbers that are not the same (1:15)

**Say:** Agreement, coverage and guessability get blurred, in the paper and in how it is cited. Agreement measures how concentrated the consensus is. The paper's own second example: select single, one hand, groups of eleven, three, three and three, giving 0.37. Eleven people agreed; nine did something else. Because the score squares the shares and a gesture always agrees with itself, the floor is about 0.05. "A equals 0.32" does not mean 32 percent agreed. Vatavu and Wobbrock fixed this at CHI 2015 with a corrected agreement rate and a significance test. Coverage, 57 percent, is the share of proposals matched by any gesture in the final set, aliases included; it rewards a large set. And guessability, the chance that a new user guesses right, is what the method is named after, and it is never measured.

**Transition:** The other limitations come from decisions the authors made.

**Watch out:** The 0.37 is printed in the paper as Equation 3; cite it. "Tapped" is inferred from Figure 4.

**Source:** p. 1087; Vatavu and Wobbrock CHI 2015, verified 2026-09-30.

---

## Slide 16 — Design decisions that shaped the result (1:30)

**Say:** Four decisions, each defensible, each with consequences. One: what counted as identical? The paper does not say. The coding rules and the Figure 4 caption imply one to three fingers were merged. If so, that merges groups and raises agreement, so the finding that fingers do not matter and the scoring rule lean on each other. Two: the set is conflict-free because the target is part of the gesture. The paper's rule is that the larger group wins; what happened to the loser it does not say, but in the set the same four gestures mean resize, zoom or open depending on where you land. We read that as a design move. The authors could reply that binding is a taxonomy dimension, so these were never the same gesture. Three: two-hand gestures were required for every command, though one hand was preferred for 25 of 27. The jump idiom and a third of the set came out of that; it is related to the "production" technique the same group later proposed against legacy bias. Four: users collapsed categories. Asked to zoom in, one said "Oh, that's the same as enlarge." Maximize folded into enlarge. Users paired minimize with shrink, yet minimize kept its own gesture.

**Transition:** And then there is the decision they could not fully make: who to recruit.

**Watch out:** Items one and two are our inferences from unreported procedure; say "the paper does not say." Minimize did not become shrink in the set.

**Source:** p. 1086 (coding), p. 1087 (grouping, rule, maximize), p. 1088 (two-hand share), p. 1090 (quote, complexity mismatch), Figure 4; Morris et al. 2014.

---

## Slide 17 — Whose habits did the study capture? (1:45)

**Say:** This is the part we most want to argue about. Who sat down? The paper says non-technical people "without prior experience using touch screens," and nobody who had used a tabletop, an iPhone "or similar." No screening instrument, no touch history. What the paper blames for the 72 percent single-point gestures is the desktop: "It's a Windows World," with specific evidence. People clicked with index and middle finger as on a two-button mouse, tapped imaginary close buttons, an imaginary taskbar. "I'm a child of the mouse."

Now look at what a 2008 adult had already touched: ATMs and kiosks since the eighties, Palm PDAs with stylus strokes since 1996, laptop touchpads, the Nintendo DS. "No prior experience using touch screens" is hard to believe, and every one of those teaches one finger, one point.

So the 72 percent has three candidate explanations this data cannot separate: the mouse, a decade of single-point touch, or plain pointing, which the paper's own related work found dominant everywhere. The pinch is a smaller puzzle: five of twenty used a one-hand reverse pinch to resize, after a year of iPhone adverts. Exposure? The study cannot tell.

What follows: elicitation captures the repertoire the participant brings in. The paper reports that as a finding about users. We read it also as a limit of the method, and the gap has widened since. Morris, Wobbrock and colleagues named it "legacy bias" in 2014. **[classmate]**

**Transition:** This connects to several things we have read.

**Watch out:** Do not say the paper promised unbiased participants. Device dates are general knowledge. Say "could also come from," not "is." The 2014 paper shares Morris and Wobbrock, not Wilson.

**Source:** p. 1084 (screening, Beringer, Voida, Epps on pointing), p. 1085 ("or similar"), p. 1090 ("It's a Windows World"), p. 1091 (pinch counts), Figure 4; Morris et al. 2014.

---

## Slide 18 — Connections to other research and course topics (1:00)

**Say:** **[TODO: Oliver – section 6. Suggested shape: one connection to each of the weeks 1–3 readings with a "so what," then two of the hooks below.]**

Hooks already in the paper, pick two: Furnas's vocabulary problem is why the set uses aliases. Norman's gulf of execution is what the no-recognizer design removes. Good and colleagues did the same elicitation for typed commands in 1984. Morris's own cooperative-gestures work raises the question of why a group device was studied one user at a time. Wilson's physics-based Surface work from the year before is what the "44 percent physical" finding is used to support.

**Source:** refs. 8, 9, 10; notes section 6; progress points 7 and 8.

---

## Slide 19 — What came next, and what it suggests (1:30)

**Say:** The authors' agenda has three layers. Next steps proper: reverse the study by showing gesture videos and asking people to guess the command, and build a recognizer. Limitations left for later: let participants see all commands first, add context, study children and other cultures. And technology implications: sense above the table and beyond the edge.

Some of this happened, and we checked each item. In 2010 the same authors had 22 new participants, blind to who authored what, judge 81 gestures; they preferred gestures from larger groups, partly because the researchers' gestures were more complex. In 2015 Vatavu and Wobbrock formalized the agreement analysis. In 2014 Morris and colleagues named legacy bias as a weakness of the method and proposed remedies. By 2020 a review covered 216 elicitation studies.

Here is the twist. The method spread; nobody adopted the set as published; yet drag, double-tap, pinch and spread are now universal. Apple shipped pinch-to-zoom in 2007, and a year later this study elicited it from people recruited for never having used a touch screen. Elicitation and the market converged. Which leaves the open question for mid-air, VR, wearables and voice: can elicitation still find anything new when every participant carries a smartphone vocabulary?

**Transition:** That brings us to what we think.

**Watch out:** Not "users preferred user-authored gestures"; the finding is about group size and complexity. Do not claim the reverse study or the recognizer were ever built; we did not check.

**Source:** p. 1091; GI 2010 pp. 261–268; CHI 2015; Interactions 2014; DIS 2020. All verified 2026-09-30.

---

# Part C – Bridges to discussion

## Slide 20 — Our reading of the paper (0:45)

**Say:** One sentence: the lasting contribution is the method, and the evidence in this paper for the method is thinner than its influence. Three reasons. The guessability test the authors proposed was never run; the 2010 study tested preference, which is partial support. The claim that popularity identifies better gestures rests on people rating their own ideas. And the study measured what participants brought in; the paper calls that a finding about users, and it is also a property of the method. None of this makes it a bad paper. It makes it a paper whose influence outran its evidence.

**Transition:** So here is what we want to argue about.

**Watch out:** Say "thinner than its influence," not "weaker than its reputation"; the deck never describes a reputation.

---

## Slide 21 — Tensions worth arguing about (1:15)

**Say:** Four questions with more than one defensible answer. Should the most popular gesture win, when 72 percent were mouse-like and consensus from novices may lock in old habits? Did the commands shape the answers, with five desktop concepts on the list? Does popularity identify better gestures, or gestures people enjoyed inventing, and what would independent validation look like? And is behaviour without feedback the right target, when real users adapt to a system in minutes?

Three smaller things. Nineteen percent of the designers' gestures were never tried: is elicitation a floor or a ceiling? The check mark, the X and the question mark are Western writing, and the paper itself says Eastern participants would differ. And minimize means drag toward your own edge of the table, a different direction for every person at a group device. **[classmate]**

**Transition:** Before we open it up, here is what you wrote.

**Watch out:** Read these as invitations, not conclusions. Pair each with a classmate who raised it.

**Source:** Four prompts from the Opus reports; small items from progress points 6, 7, 9 (p. 1090–1091, Figure 4).

---

## Slide 22 — Over to you (1:00)

**Say:** **[TODO: after the commentaries – name three or four recurring themes or sharp observations, credit each person, and say which tension on slide 21 it feeds.]**

Here is how the next half hour runs. First a five-minute live elicitation: we show you the effect of two commands, one from the top of Figure 3 and one from the tail. On three, everyone performs the gesture they would use, on the desk. We count the groups and compute the agreement score together. Then five minutes in pairs: pick one of the four questions and take opposite sides. Then twenty minutes together, and we follow whatever is interesting.

**Watch out:** Choose "move" and "undo" or "help." Expect near-total agreement on move and a spread on undo, including smartphone-era gestures such as shake, which proves the legacy point in the room. Tally on the board. The full discussion plan is a separate document still to be written.
