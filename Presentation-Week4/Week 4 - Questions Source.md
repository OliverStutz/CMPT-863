# Week 4 question sets, presenter version: files 13 to 15

Generated 2026-09-30 from `slides/13-Questions-Paper.md`, `14-Questions-Presentation.md` and `15-Questions-Synthesis.md` by `slides/_build-deck-source.py`. Edit the section files, then regenerate. Verification logs and open items are left in the section files. The class-facing version with question stems only is `Week 4 - Questions - Class only.md`.


---

# 13 Questions: Did they read the paper?

**Purpose:** test whether classmates read Wobbrock, Morris and Wilson (CHI 2009) closely enough to answer from the paper what the presentation does not say. **Count:** 12 questions (4 quick checks, 5 short answers, 3 analysis). **Discriminates:** people who read the paper, including its figures, captions and Discussion, from people who rely on the talk or on a skim of the abstract. **Use:** quick checks as a quiet written check or show of hands around the warm-up; short answers as probes when the group discussion stalls; analysis questions as pair prompts (see 3.3 and "How to use in class" below).

Page numbers are the printed pages (PDF page N = printed page 1082 + N). "Section files" means `01-Background-2008.md` to `12-Discussion-Bridge.md`. Where a section's presenter notes ("Watch out") contain part of an answer, that is named, because the class hears only the on-slide text and the "Say" script unless a presenter answers a question from the notes.

## Q13.1: The paper lists its contributions as four numbered items. What are they, and what does the abstract present that the numbered list leaves out?

- **Form:** quick check (sentence)
- **Tests:** recall and comprehension: what the authors claim to contribute, and the gap between the list and the abstract.
- **Model answer:** The four contributions are "(1) a quantitative and qualitative characterization of user-defined surface gestures, including a taxonomy, (2) a user-defined gesture set, (3) insight into users' mental models when making surface gestures, and (4) an understanding of implications for surface computing technology and user interface design" (p. 1084). The method is not on the list. The abstract opens with it: "We present an approach to designing tabletop gestures that relies on eliciting gestures from non-technical users by first portraying the effect of a gesture, and then asking users to perform its cause" (p. 1083). The agreement score is not claimed as new either; the process "was adopted from prior work [33]" (p. 1087).
- **Why it discriminates:** Section 10's slide and script say only that "the core is a method" and classify the contributions in Wobbrock and Kientz's terms. The four items appear only in section 10's Watch out notes. A student who skipped the paper will list the talk's framing: "a method, a gesture set, an agreement score."
- **Wrong answers that reveal:** "The agreement score" as a contribution shows the student does not know it came from the 2005 EdgeWrite work (p. 1087). "The method" as item 1 shows the student is reciting the presentation's reading, not the paper's list.
- **Follow-up probe:** Why might authors leave the method off their own contribution list while opening the abstract with it?

## Q13.2: The Discussion says "Figure 2 indicates that agreement is low after the first seven referents." Which figure actually shows this, which seven referents are meant, and what design advice does the paper draw from it?

- **Form:** quick check (sentence)
- **Tests:** recall and close reading of a figure: catching the slip and reading Figure 3.
- **Model answer:** The sentence is on p. 1091. Agreement per referent is Figure 3 (p. 1088); Figure 2 is the taxonomy breakdown (p. 1087), so "Figure 2" is a slip for Figure 3. Reading Figure 3's x-axis, the first seven by one-hand agreement are move a little, move a lot, duplicate, rotate, select group, insert and select single (p. 1088; chart reading: move a little and move a lot about 1.0, duplicate about 0.52, rotate, select group and insert about 0.48, select single 0.37). The advice: "referents beyond this point may benefit from an on-screen widget as well as a gesture" (p. 1091).
- **Why it discriminates:** Section 09 mentions the slip only in its Watch out notes, "only if asked." No section lists the seven referents or the "widget as well as a gesture" advice. Section 12.3 refers to "the top of Figure 3" without naming them.
- **Wrong answers that reveal:** "Figure 2 is right" shows the student did not compare the figures. Naming the seven least complex referents from Table 1 (move a little, move a lot, select single, rotate, shrink, delete, enlarge) shows the student confused the authors' complexity ranking with the agreement ranking, which the paper says differ (p. 1090).
- **Follow-up probe:** Insert is among the seven most agreed one-hand referents, yet it has no dedicated gesture in the set (p. 1087). How can both be true?

## Q13.3: When asked to zoom in, one participant said "Oh, that's the same as enlarge." What does the paper conclude from simplifications like this, and how does the final set use them?

- **Form:** quick check (sentence)
- **Tests:** comprehension: the "Simplified Mental Models" observation and its design consequence.
- **Model answer:** The paper reports that participants "often made simplifying assumptions," with similar mental models for enlarge and maximize, shrink and minimize, and pan and move (p. 1090). Because the authors' complexity order (Table 1) and the agreement order (Figure 3) differ, "participants and the authors did not always regard the same referents as 'complex'" (p. 1090). The set uses this to "unify the gesture set and disambiguate the effects of gestures based on where they occur" (p. 1090): the four enlarge gestures mean enlarge on an object, zoom in on the background, and open on a container such as a folder (p. 1088; Figure 4, p. 1089). Splaying five fingers enlarges an object but zooms in on the background (p. 1091).
- **Why it discriminates:** No section quotes the participant or reports the "Simplified Mental Models" observation. Section 04's Watch out notes mention reuse by target in one line; the slides and scripts do not.
- **Wrong answers that reveal:** "Zoom in has its own pinch gesture" shows the student did not look at Figure 4, where zoom in reuses the enlarge gestures on the background. "Participants found zoom in confusing" turns a design opportunity into a failure and misses the paper's point.
- **Follow-up probe:** If participants did not share the authors' sense of what is complex, what does that do to the correlations built on the authors' complexity ratings (r = -0.52, 0.71, -0.59, pp. 1087 to 1088)?

## Q13.4: Name the four dimensions of the paper's taxonomy, and explain in one sentence what the flow dimension distinguishes.

- **Form:** quick check (sentence)
- **Tests:** recall and comprehension of the descriptive layer of the paper.
- **Model answer:** Form, nature, binding and flow; "the authors manually classified each gesture" along them (p. 1086, Table 2). Flow is discrete when the "response occurs after the user acts," the gesture being "performed, delimited, recognized, and responded to as an event," such as tracing a question mark for help; it is continuous when the "response occurs while the user acts," as in most resize gestures (Table 2, p. 1086; text, p. 1087). Figure 2 puts discrete at .535 and continuous at .465 (p. 1087).
- **Why it discriminates:** Section 04 names the taxonomy and, in its Watch out notes only, the four dimension names. No section defines flow, binding or nature categories. A student who skipped the paper may know that a taxonomy exists but not what it describes.
- **Wrong answers that reveal:** "One hand versus two hands" as a dimension shows confusion of the study conditions with the taxonomy. Describing flow as "fast versus slow" shows the student guessed from the word, not from Table 2.
- **Follow-up probe:** Which flow does a recognizer find harder to support in real time, and what does that imply when 46.5% of all proposed gestures were continuous (Figure 2, p. 1087) and the set's taxonomy breakdown is close to that of all proposals (p. 1088)?

## Q13.5: The paper describes using a computer as a dialogue and then deliberately turns the study into a monologue. Reconstruct its argument from dialogue and feedback, and name the second thing, besides Windows and Macintosh elements, that the prompts deliberately left out.

- **Form:** short answer (paragraph)
- **Tests:** comprehension of the three design choices and the paper's own reasons.
- **Model answer:** The argument (p. 1085): using a system is "a user-computer dialogue [6], a conversation mediated by a language of inputs and outputs." When something is misunderstood it may be rephrased; feedback "either endorses or deters a user's action, causing the user to revise his or her mental model and possibly take a new action." The authors "did not want the vicissitudes of gesture recognition to influence users' behavior," so they removed the gulf of execution [10], "creating, in essence, a monologue in which the user's behavior is always acceptable." This lets them "observe users' unrevised behavior, and drive system design to accommodate it," which matters because tabletops "may be used in public spaces, where the importance of immediate usability is high" (p. 1085). Unrevised means not revised by recognizer feedback, not free of habits. The two things not shown: "To avoid bias [5], no elements specific to Windows or the Macintosh were shown," and "no specific application domain was assumed"; instead participants saw "a simple blocks world of 2D shapes" (p. 1085). Behind all this is the choice to ask users at all: designers may be principled, but "user behavior is rarely so systematic" (p. 1083).
- **Why it discriminates:** Section 04.1 gives the three choices, the "monologue" quote and a one-line summary ("Nothing fails, so nothing pushes the participant to change"), so a listener can answer part of it. The paper-only parts are the dialogue and rephrasing argument, "endorses or deters," and "no specific application domain was assumed." Section 09 lists "no application context" as a limitation but not as a deliberate choice. A non-reader will say "they didn't have a recognizer yet."
- **Wrong answers that reveal:** "The prototype could not recognize gestures" shows the student reads the choice as a technical gap, not a design decision. "Unrevised means participants had no prior habits" repeats the misreading the paper itself rules out with its desktop findings (p. 1090).
- **Follow-up probe:** The paper admits on p. 1091 that this choice has "drawbacks." Which ones does it name?

## Q13.6: Participants rated each gesture for "good match" and for "easy to perform." For each rating, which predictors did the paper test and find had no significant effect, and what does that pattern of null results show?

- **Form:** short answer (paragraph)
- **Tests:** comprehension of the evidence: a double dissociation between "good match" and "easy to perform."
- **Model answer:** Scale 1: "The gesture I picked is a good match for its intended purpose." Scale 2: "The gesture I picked is easy to perform." Both ran from 1 = strongly disagree to 7 = strongly agree (p. 1088). Goodness was predicted by group size (larger groups rated higher, chi-square 34.10), by complexity (r = -0.59; simple referents 5.6, complex 4.9) and by planning time (longer planning, lower goodness, chi-square 38.98). Articulation time "did not significantly affect goodness ratings." Ease was predicted by articulation time (gestures rated easy took about 3.4 s, difficult about 2.0 s, chi-square 17.00) and by touch-event counts (chi-square 21.82). Planning time "did not affect perceptions of gesture ease," and complexity "did not correlate significantly with average ratings of gesture ease" (all p. 1088). So thinking predicts fit, doing predicts ease, and neither crosses over.
- **Why it discriminates:** Section 05.1 gives both scale names, section 07.2 the positive links (planning to goodness, articulation to ease, complexity to goodness) and section 07.1 group size to goodness, so the stem now asks only for what the talk leaves out: the three null results, which appear only in section 07's Watch out notes. A non-reader will say "complex commands got harder gestures," which the paper contradicts.
- **Wrong answers that reveal:** "Complex referents were rated harder to perform" shows the student merged the two scales; complexity predicts goodness, not ease (p. 1088). "Easy gestures were faster" reverses the paper's surprising direction.
- **Follow-up probe:** The touch-event count is offered as objective corroboration of the ease result. The paper itself says "long lived gestures will have more touch events" (p. 1088). Is it independent evidence?

## Q13.7: Four referents got no dedicated gesture. Which, why, in two different ways, and what does the paper say this tells us about the set?

- **Form:** short answer (paragraph)
- **Tests:** comprehension of the gesture-set construction and its limits.
- **Model answer:** Insert, maximize, task switch and close (p. 1087). Two reasons. For insert and maximize, "the action most participants took comprised more primitive gestures: insert used dragging, and maximize used enlarging," so they needed nothing new. For task switch and close, "participants relied on imaginary widgets; a common gesture was not feasible": most tapped an imaginary taskbar button for task switch and an imaginary button in the top-right corner for close (p. 1087). The paper's conclusion: "Our user-defined set is useful, therefore, not just for what it contains, but also for what it omits" (p. 1088). It also notes that insert and maximize are exactly the two referents where hand preference was evenly split, "neither of which were included in the user-defined gesture set because they reused existing gestures" (p. 1088). With the two move referents combined, 22 of 27 referents got dedicated gestures (p. 1087).
- **Why it discriminates:** Section 04.2's slide and script name only close and task switch; section 05.2's slide and script mention the insert and maximize tie in hand preference. The two-way explanation is in Watch out notes of sections 01, 04 and 05, not in the scripts. The "what it omits" sentence appears in no section. A non-reader will say "the method failed for four commands."
- **Wrong answers that reveal:** "Nobody agreed on insert" shows the student did not check Figure 3, where one-hand insert is among the highest at about 0.48 (p. 1088). "Close and task switch were too complex" ignores that close is rated 2.00 in Table 1 (p. 1086), below the mean of 2.70.
- **Follow-up probe:** Is "relied on imaginary widgets" evidence about users, or about the desktop habits the paper documents under "It's a Windows World" (p. 1090)?

## Q13.8: Under "A Land Beyond the Screen," what did participants do, which gestures in the final set depend on that space, and does the paper's reassurance about tables without edge sensing hold?

- **Form:** short answer (paragraph)
- **Tests:** comprehension of a qualitative finding and its trace in Figure 4; evaluation of a claim.
- **Model answer:** Participants "conceived of a world beyond the edges of the table's projected screen": they dragged in from off-screen "treating it as the clipboard" and dragged off-screen to delete and reject; one participant gave different edges different meanings (off the top was delete, off the left was cut, and she dragged paste in from the left to pair it with cut) (p. 1090). In Figure 4, Delete1 is "drag offscreen," Delete2 is Move2 to an off-screen destination, Paste2 is "drag from offscreen," Paste3 is Move2 from an off-screen source, and Reject2 and Reject3 reuse Delete1 or Delete2 (p. 1089). The paper suggests "a low-resolution sensing boundary beyond the high-resolution display area" and adds: "Conveniently, these gestures have alternatives in the user-defined set for tables without a sensing boundary" (p. 1091). That holds for paste (Paste1: tap) and reject (Reject: draw 'X', the unnumbered first reject gesture). We read Figure 4 as giving delete no on-screen alternative: both delete gestures leave the screen (p. 1089). This is our reading of the figure, not a statement in the paper.
- **Why it discriminates:** Section 08's script mentions dragging off the screen edge and the sensing border; section 07.1 lists "off-screen drags." No section names the off-screen gestures in Figure 4, the edge-meaning participant, or the "Conveniently" sentence.
- **Wrong answers that reveal:** "Off-screen gestures were dropped because the table could not sense them" shows the student did not look at Figure 4. "Every off-screen gesture has an alternative" repeats the paper's claim without checking it against the figure.
- **Follow-up probe:** The paper also recommends considering a widget for delete (p. 1091). Does that fill the gap, and is it a coincidence?

## Q13.9: The paper makes two separate recommendations about on-screen widgets. What are they, what evidence is each based on, and which question from the method section do they answer?

- **Form:** short answer (paragraph)
- **Tests:** comprehension: separating two recommendations with different evidence.
- **Model answer:** (1) Because agreement is low after the first seven referents in Figure 3 (the text says Figure 2), "referents beyond this point may benefit from an on-screen widget as well as a gesture" (p. 1091). The evidence is agreement scores. (2) Because "enough participants acted on imaginary widgets," designers "might consider using widgets along with gestures for delete, zoom in, zoom out, accept, reject, menu access, and help" (p. 1091). The evidence is observed behaviour. Undo is not on that list. Both answer a question the paper posed in its method section: "If we are to choose a mix of gestures and widgets, how should they be assigned?" (p. 1085). The abstract already flags it: some commands "elicit little gestural agreement, suggesting the need for on-screen widgets" (p. 1083). In both cases widgets are added to gestures, not instead of them.
- **Why it discriminates:** No section gives either recommendation. Sections 04, 07 and 10 mention imaginary widgets only as a finding (close, task switch, the desktop habit). A non-reader will guess "use widgets for close and task switch," which the paper does not list.
- **Wrong answers that reveal:** One merged list ("widgets for all low-agreement commands, including undo") shows the student did not see that the second list rests on imaginary-widget behaviour, not agreement, and leaves undo out. "Replace gestures with widgets" misses "as well as" and "along with."
- **Follow-up probe:** Undo has the highest complexity rating (5.00, p. 1086) and low agreement (about 0.18, p. 1088), yet it is not on the second list. Which recommendation still covers it, and why might it be missing from the other?

## Q13.10: Figure 2 reports the form dimension separately for the right and the left hand. Why is it split, how do the two hands differ, and what does that do to the paper's "about 72% mouse-like" claim?

- **Form:** analysis (pair discussion)
- **Tests:** analysis of a figure: reading the data behind a headline number and interpreting it with care.
- **Model answer:** The paper says "the scope of the form dimension is within one hand. It is applied separately to each hand in a 2-hand gesture" (p. 1086); the caption adds that "all participants were right-handed" (p. 1087). Figure 2: right hand one-point path .450, one-point touch .274, static pose and path .140, dynamic pose .092, static pose .037, dynamic pose and path .007; left hand one-point path .379, one-point touch .343, static pose and path .177, static pose .056, dynamic pose .037, dynamic pose and path .007 (p. 1087). Both hands sum to about 72% one-point, which matches "about 72% of gestures were mouse-like one-point touches or paths" (p. 1090). The difference is in the mix: the left hand touches in one place more often (.343 against .274) and moves a single point less (.379 against .450). We argue this fits the paper's own example of a two-hand move, "holding the object with the left index finger while tapping the destination with the right" (p. 1085): the non-dominant hand anchors, the dominant hand acts. The paper does not interpret the asymmetry, and it does not say which hand performed one-hand gestures. Two cautions on the 72%: it is a share of hand forms, not of participants; and a gesture counts as one-point "even if the user casually touches with more than one finger at the same point" (p. 1086), so the coding rule raises the mouse-like share.
- **Why it discriminates:** Section 07.1 cites the 72% and shows a Figure 2 inset; its Watch out notes give both hands' sums. No section explains why form is per hand, compares the hands, or links the asymmetry to the p. 1085 example. A non-reader will say "72% of participants used one finger."
- **Wrong answers that reveal:** "72% of participants were mouse-like" confuses gestures with people. "The left column is left-handed participants" shows the student missed the caption (all right-handed).
- **Follow-up probe:** Some participants "tapped an object first to select it, then gestured on top of the very same object," which the paper says negates "a key benefit of gestures that couples selection and action" (p. 1090). Is that a second kind of mouse-likeness that the form coding cannot see?

## Q13.11: The paper uses the figure "each author proposed only 43.5%" for two different arguments. What are they, and does the comparison support the claim the authors set out to test?

- **Form:** analysis (pair discussion)
- **Tests:** evaluation of how the paper reads its own designer comparison.
- **Model answer:** Before the study the three authors each designed gestures for all 27 referents, and "it was hypothesized that the 'wisdom of crowds' would generate a better set than the authors" (p. 1090). First use: each author matched only 43.5% of the user-defined set and all three together 60.9%, which "suggests that three experts cannot generate the scope of gestures that 20 participants can"; 19.1% of each author's gestures were never tried by any participant, "the authors are either thinking creatively or are hopelessly lost!" (p. 1090). Second use: the set seems "in hindsight, to be sensible choices," but "the gestures are not, in fact, 'obvious'," because "each author proposed only 43.5% of the gestures in their own designs" (pp. 1090 to 1091). The paper also says a coherent set "was not clear a priori" (p. 1090). Evaluation: the comparison measures overlap with the users' own set, so it shows scope and non-obviousness, not that the users' set is better. The hypothesis says "better," but the conclusion the paper actually draws is the narrower one about scope (p. 1090).
- **Why it discriminates:** Section 05.2 gives the three numbers, the joke and the overlap-not-quality reading, so the first use is known from the talk. The second use (the set is "not obvious") and the "a priori" remark appear in no section. A listener will give only the first argument.
- **Wrong answers that reveal:** "The users' set was proven better than the designers'" turns a scope result into a quality result. "The designers' gestures were never tried, so they were bad" ignores that the paper leaves "creative" open.
- **Follow-up probe:** If the same number shows both that designers miss user gestures and that the user gestures are not obvious, could any result of this comparison have counted against the method?

## Q13.12: On p. 1091 the paper recommends sensing above the table, and on the same page it admits what removing the dialogue cost. Put the two together: is "acting above the table" a fact about users or a product of the method, and what does the paper itself concede about how feedback would change it?

- **Form:** analysis (pair discussion)
- **Tests:** evaluation: weighing the paper's technology implications against its own limitations section.
- **Model answer:** The behaviour: "We instructed participants to touch the table while gesturing. Even so, some participants gestured in ways few tables could detect": a palm-up beckon for help, hands crossed in an "X" with the top hand about 3 inches up, a two-hand "lift" to the clipboard, "glue" applied in the air (p. 1090). The implication: tables "may benefit from the ability to look down or sideways at users' hands, rather than just up," which "provides robustness for users who forget to remain in contact with the surface at all times" (p. 1091). Then the paper's own concession: "Of course, interactive systems that provide feedback will implicitly remind users to remain in contact with the table, but users' unaltered tendencies clearly suggest a use for off-table sensing" (p. 1091). Its limitations paragraph says the study removed the dialogue to avoid "the inevitable bias and behavior change that comes from recognizer performance and feedback. But there are drawbacks to this approach" (p. 1091), and one participant added fingers "to ensure that I was pressing," which the paper links to "the lack of feedback from the table" (p. 1090). Case for "about users": the instruction was given and people still left the surface, so the tendency is real. Case for "about the method": the paper itself predicts that feedback would reduce it, so the size of the effect in a real system is unknown. The paper settles on unaltered tendencies; the evidence supports the tendency, not how often it would survive feedback.
- **Why it discriminates:** Section 08's script gives the above-table examples, the instruction and the sensing suggestions, and section 04's notes list costs of no feedback. The "Of course, interactive systems that provide feedback" concession and the "drawbacks" sentence of the limitations are in no section. A listener will argue only one side.
- **Wrong answers that reveal:** "Participants were never told to touch the table" contradicts p. 1090. "The paper ignores feedback" misses that it raises the point itself on p. 1091.
- **Follow-up probe:** The limitations paragraph ends "Thankfully, even with a lack of application context and upfront knowledge of all referents, participants still exhibited a substantial level of agreement" (p. 1091). Does agreement answer the worry about feedback, or a different worry?

## How to use in class

- **Warm-up (5 min, after the live elicitation):** Q13.2 and Q13.3 tie straight to the elicitation the class has just done (Figure 3, where agreement falls off; zoom in "same as enlarge"). Ask as a show of hands, then take one sentence from a volunteer.
- **Quiet written check (one minute, optional):** Q13.1 and Q13.4 on a slip of paper or a poll. They take a sentence each and show at a glance who read the paper.
- **Pairs (5 min):** give each pair one of Q13.10, Q13.11 or Q13.12. Each has a defensible answer on both sides. Q13.11 pairs with tension 3 on slide 12.2 (does popularity identify better gestures); Q13.12 with tension 4 (behaviour without feedback).
- **Group (20 min), held for lulls:** Q13.6 when the timing results come up; Q13.7 and Q13.9 when the talk turns to what the set leaves out or to widgets; Q13.8 when someone raises sensing; Q13.5 if the group questions the no-recognizer choice.
- **Participation grading:** note who gives a paper-only detail (the "what it omits" sentence, the second use of 43.5%, the feedback concession on p. 1091, the delete gap in Figure 4). Those details show reading, not listening. Record the name and the question number during the session.


---

# 14 Questions: Did they listen to the presentation?

**Purpose:** check whether classmates took in what the presentation added to the paper: outside facts, corrected readings, and the presenters' own arguments. **Count:** 10 questions (4 quick checks, 4 short answers, 2 analysis). **Discriminates:** people who followed the talk from people who only read the paper. Every question turns on a fact or an argument that is in the section files and not in the paper. A paper-only answer is listed under "Why it discriminates" for each. **Use:** quick checks as a written check right after the talk or as warm-up probes; short answers as stall probes in the group phase; the two analysis questions as pair prompts (see "How to use in class").

## Q14.1: The paper gives no date for its study. How did we pin down when it was done, and by when was it finished?

- **Form:** quick check (sentence)
- **Tests:** recall, plus comprehension of an inference: how an outside document dates the study.
- **Model answer:** Microsoft filed a patent application on the study's method, US 2010/0031202, "User-Defined Gesture Set for Surface Computing," with Morris, Wobbrock and Wilson as inventors, on 4 August 2008. So the study was complete before August 2008. The application was published in February 2010 and abandoned in December 2014; it was never granted. The paper does not mention it. (01.1 timeline and Watch out; 03.1; 04.1 last bullet and Watch out. Paper: no study date anywhere, CHI dates only, p. 1083.)
- **Why it discriminates:** the paper names no date and no patent. A paper-only reader can say only "before CHI in April 2009" or "2008 or so."
- **Wrong answers that reveal:**
  - "It was done in 2009." Confuses publication with data collection; missed 01.1.
  - "Microsoft patented the gesture set." Missed that it was an application, never granted, and that its claim is about the procedure (03.1 Transition and verification log).
- **Follow-up probe:** Why does it matter that the study was finished before Apple's public patent threats of January 2009? (03.1: the threats came after the study, so they cannot have prompted it.)

## Q14.2: What pixel density did we compute for the study's Surface prototype, what did we compare it with, and does the paper itself blame resolution for anything?

- **Form:** quick check (sentence)
- **Tests:** recall and comprehension: the gap between a measured fact and our inference.
- **Model answer:** About 43 pixels per inch (1024 pixels across 24 inches, p. 1085), against 264 ppi on a current iPad Pro, roughly six times finer in each direction (08.1). The paper never names resolution as a problem. Its advice that "demands for precise positioning should be avoided" (p. 1091) rests on user behaviour (12 of 14 people resized sideways for two-hand enlarge), not on pixels; linking it to 43 ppi is our inference (08.1 Watch out). The table also could not sense hands above it or drags beyond its edge (08.1, p. 1090).
- **Why it discriminates:** the paper gives the size and resolution but no density, no comparison, and no claim about resolution. A paper-only reader would either not know the figure or would say the paper blames the coarse display.
- **Wrong answers that reveal:**
  - "The paper says the low resolution forced coarse gestures." Attributes our inference to the authors.
  - "About 43 percent." Mixes the ppi figure with the 43.5% designer overlap or the 43.9% physical gestures; a sign of half-listening.
- **Follow-up probe:** We said the iPad comparison is fair only in one respect. Which, and why? (08.1 Watch out: density only; the iPad is hand-held, not a table.)

## Q14.3: Name one gesture in the final set that we called culture-bound, and one verified example of a mark or movement whose meaning changes across cultures.

- **Form:** quick check (sentence)
- **Tests:** recall and comprehension: how culture enters a gesture set through its symbolic part.
- **Model answer:** Culture-bound gestures in the set: accept is a check mark, reject an X, help a question mark, next and previous a line drawn across the object, drawn left to right in Figure 4 (p. 1089; 06.1, 12.2). Verified examples from 06.1: in Japan a circle marks correct and a check can mark a wrong answer; in Bulgaria a nod means no; a thumbs-up is an insult in Iran. Our point: the most abstract commands got the most culture-bound gestures (06.1 Say). That next follows reading direction is our inference (06.1 Watch out).
- **Why it discriminates:** the paper says only that "children, Eastern, or uneducated participants would behave differently" (p. 1091). It never examines its own marks for cultural loading and gives no examples.
- **Wrong answers that reveal:**
  - "Eastern participants would gesture differently." True but paper-only; it does not answer either part.
  - "The thumbs-up gesture for accept." Thumbs-up is not in the set (06.1 Watch out); the respondent remembers the example but not the point.
- **Follow-up probe:** We used the Bulgarian head nod with a caveat. What was the caveat? (06.1 Watch out: it is a body gesture, not a surface gesture; it only shows that the meaning of a movement is learned. The check mark is the example that touches the set directly.)

## Q14.4: Which course reading did we use to argue that elicitation can pull toward the familiar, and what evidence from the paper did we pair it with?

- **Form:** quick check (sentence)
- **Tests:** recall and comprehension of a connection to the course readings.
- **Model answer:** Greenberg and Buxton (CHI 2008): early evaluation converges on "that which we are already familiar with" (their p. 115). We argued elicitation carries the same risk, with about 72% of gestures mouse-like one-point touches or paths (p. 1090; 10.1). We also said their excuse for skipping evaluation of early designs only partly applies, because the paper calls its set "a good candidate for deployment" (p. 1091; 10.1 Watch out).
- **Why it discriminates:** the paper does not cite Greenberg and Buxton. Someone who read only the paper cannot make the link. Someone who did the course readings could guess it, so the pairing with the 72% and the partial-excuse point is what shows they listened.
- **Wrong answers that reveal:**
  - "Olsen." Mixes up the readings; Olsen was used for "new design participants" (10.1).
  - "Greenberg and Buxton say elicitation is harmful." They are about usability evaluation, not elicitation; the link is our analogy (10.1 Watch out).
- **Follow-up probe:** Which reading did we use to say methods need "repeated validation," and what did the paper offer instead? (Wobbrock and Kientz 2016, p. 41; one study whose own ratings "go a long way to validating," p. 1088; 10.1.)

## Q14.5: Someone asks, "So was the study done to protect Microsoft from lawsuits?" What did we answer, and on what evidence?

- **Form:** short answer (paragraph)
- **Tests:** evaluation: separating documented facts from interpretation, and refusing to attribute a motive.
- **Model answer:** We found no evidence of a lawsuit motive and claimed none (03.1). Documented: Apple bought FingerWorks in 2005; Jobs said of multi-touch in January 2007, "and boy, have we patented it"; Microsoft hired Buxton in December 2005, showed Surface in May 2007, put it in AT&T stores in April 2008 and demonstrated Windows 7 touch in May 2008; Microsoft filed a patent application on this study's method in August 2008. Against a lawsuit motive: Apple and Microsoft had a patent cross-licence since 1997, and Apple's public threat (Tim Cook, January 2009) came after the study was finished. The paper itself gives academic reasons and builds on Wobbrock's 2005 guessability work (pp. 1083 to 1084). What we argued is narrower and labelled "we argue": Microsoft was racing to establish and own its gesture conventions while Apple's were being fixed. The filing fits offensive, defensive or routine corporate practice; we cannot tell which (03.1 Watch out).
- **Why it discriminates:** the paper contains none of the business timeline, no patent and no cross-licence. A paper-only reader can only cite the paper's academic motives or speculate.
- **Wrong answers that reveal:**
  - "Yes, the patent shows it was defensive." Turns a filing into a motive; missed the "we argue" line and the cross-licence.
  - "The authors hid the patent." We said only that the paper does not mention it (03.1 Watch out, 09.1 Watch out).
- **Follow-up probe:** Why does the 1997 cross-licence weaken the lawsuit story for Apple and Microsoft in particular, but not for multi-touch in general? (03.1: it covers those two firms; Elan sued Apple in April 2009, so multi-touch IP was contested elsewhere.)

## Q14.6: The paper found that gestures rated easy took longer to perform than gestures rated difficult. How did we read the timing data, and how does our explanation differ from the paper's?

- **Form:** short answer (paragraph)
- **Tests:** analysis: interpreting a counterintuitive result and judging the evidence behind it.
- **Model answer:** The numbers are the paper's (p. 1088): planning time rises with complexity (r = 0.71), about 8 s for simple and about 15 s for complex referents; articulation time does not track complexity; gestures rated easy took about 3.4 s, those rated difficult about 2.0 s. The paper's guess is "perhaps because they were smoother or less hasty." Our reading (07.2): (1) thinking takes several times as long as doing (8 to 15 s against 2 to 3.4 s), so for a walk-up table the cost is in finding the gesture, not making it; (2) a motor explanation: a long drag is easy and slow, a quick tap on a precise spot is hard and fast, so "hard means precise, not long." Caveats we raised: the two bars split the data differently (complexity versus ease rating), so "several times" is loose (about 2.4 to 7.5); the touch-event count the paper offers as corroboration is largely the same duration measured again; the chi-square tests treat 1,074 gestures from 20 people as independent (07.2 Watch out).
- **Why it discriminates:** the paper gives the numbers and "smoother or less hasty" but neither the thinking-versus-doing comparison, nor the precision explanation, nor the critique of the touch-event corroboration.
- **Wrong answers that reveal:**
  - "Easy gestures were smoother or less hasty." The paper's hedge only; the respondent did not hear our reading.
  - "Complex gestures took longer to perform." Confuses planning with articulation; the paper says complexity did not correlate with articulation time (p. 1088).
- **Follow-up probe:** If the cost is in finding the gesture, what would you change about a walk-up table: the gestures, or something around them? (Links to the paper's widget advice, p. 1091, and to the synthesis set (file 15) on what a walk-up table should optimise.)

## Q14.7: Each author matched only 43.5% of the users' set. We said this measures overlap, not quality, and we tied it to one of Olsen's steps for "empowering new design participants." Which step does it support, which claim does it leave open, and what did we say the 19.1% leaves undecided?

- **Form:** short answer (paragraph)
- **Tests:** evaluation: what a built-in baseline can and cannot show.
- **Model answer:** The yardstick is the users' own set, so the designers are scored by how much they resemble the users and cannot win by definition (05.2 Say and Watch out). The 19.1% of each author's gestures that no participant tried could be creative or "hopelessly lost," as the authors joke (p. 1090); the comparison cannot tell which, because nothing measured how guessable or usable those gestures were. The paper's "wisdom of crowds" hypothesis promised a "better" set (p. 1090), but "better" is judged against the crowd itself. The paper's own conclusion is the modest one: "three experts cannot generate the scope of gestures that 20 participants can" (p. 1090). In 10.1 we used the same numbers for Olsen's "new design participants": they arguably support his second step (existing designer-made gestures are not enough) but not the claim that the resulting designs are "better." We also asked whether elicitation is a floor or a ceiling, since the 19.1% was never tested (12.2).
- **Why it discriminates:** a careful paper reader can reach "overlap, not quality" unaided (the paper's own conclusion is the scope sentence on p. 1090, and Q13.11 expects that answer from the paper alone), so that part does not discriminate. The Olsen mapping (10.1) and the "floor or ceiling?" framing of the 19.1% (12.2) are only in the talk. A paper-only reader will stop at "three experts cannot generate the scope of gestures that 20 participants can" or at "users beat the designers."
- **Wrong answers that reveal:**
  - "It proves the user-defined set is better than the designers'." Takes the paper's framing at face value.
  - "The designers were bad at gesture design." Confuses overlap with quality in the other direction.
- **Follow-up probe:** What test would turn overlap into quality? (The reverse study the paper proposes: unlabeled clips of user and designer gestures shown to 20 new people, p. 1091; never run, 11.1 and 12.1.)

## Q14.8: What can a sample of 20 people support, what can it not, and why did we say "norm is not the same as sufficient"?

- **Form:** short answer (paragraph)
- **Tests:** evaluation: sample size against the claims made, and the field's norm.
- **Model answer:** Twenty people are enough to see strong consensus, such as the move referents, but too few to split results by age or culture (09.1). The paper breaks down nothing by age and reports no ethnicity; the sd of 15.6 implies only that at least one person was under about 28 or over about 58 (06.1). The DIS 2020 review of 216 elicitation studies found an average of 25 participants (SD 4), with 20 the most frequent choice (36% of studies), which the review says is "likely explained by the fact that Wobbrock et al. [123] employed 20 participants" (09.1 Watch out, review p. 859). So 20 became the norm partly because of this paper; that makes it conventional, not adequate.
- **Why it discriminates:** the paper states n = 20 (p. 1085) and the limitations sentence (p. 1091) but nothing about the field's norm or about this paper setting it. The circularity is the presenters' point.
- **Wrong answers that reveal:**
  - "Twenty is too small for anything." Ignores that consensus on simple referents is robust (09.1).
  - "Twenty is standard, so it is fine." Heard the norm but missed that the norm traces back to this paper.
- **Follow-up probe:** Besides sample size, which statistical choice did we say inflates the paper's confidence? (Chi-square on N = 1074 gestures from 20 people treated as independent, pp. 1087 to 1088; 09.1, 07.2.)

## Q14.9: About 72% of gestures were mouse-like one-point touches or paths. We gave more than one source for that number. Name the sources, and explain why the study's data cannot tell them apart.

- **Form:** pair discussion (analysis)
- **Tests:** analysis: competing explanations for one observation, and what a study would need to separate them.
- **Model answer:** Three sources were put forward across the talk. (1) The desktop and the mouse, the paper's own reading: two-finger "clicking," imaginary close buttons and taskbar, "I'm a child of the mouse" (p. 1090; 06.1, 07.1). (2) Single-point touch devices a 2008 adult could already have used: laptop trackpads (from 1994), Palm Graffiti (1996), touch-screen ticket machines (New York, 1999), the Nintendo DS (2004); every one teaches one finger or one point (06.1 Watch out, 01.1). (3) The method itself: novices, no feedback, and several desktop referents (cut, paste, minimize, maximize, task switch) capture the repertoire people bring in; the paper calls that a finding about users, we argued it is also a property of the method (06.1, 12.1, 12.2 question 2; Greenberg and Buxton in 10.1). The data cannot separate them because the screening only named "an interactive tabletop, Apple iPhone, or similar" (p. 1085) and no touch history was recorded (06.1); every source predicts the same one-point behaviour; and there is no comparison condition with feedback or with non-desktop referents. "Unrevised behavior" (p. 1085) means not revised by recognizer feedback; it never promised habit-free participants (04.1 Watch out).
- **Why it discriminates:** the paper offers one explanation, the desktop (p. 1090). The device list, the "property of the method" reading and the "cannot separate" argument are only in the talk.
- **Wrong answers that reveal:**
  - "Because they were used to Windows." Only the paper's explanation.
  - "The study assumed participants had no prior habits." Repeats the corrected misreading of "unrevised behavior"; the paper reports the desktop bias itself as a finding.
- **Follow-up probe:** What study would separate the three? (11.1 open question 1: record prior device exposure and analyse it as a covariate; Morris et al. 2014 legacy-bias remedies, production, priming and partners, 11.1.)

## Q14.10: Our one-sentence reading was "The lasting contribution is the method. The evidence in this paper for the method is thinner than its influence." Make the case for each half, and say what we conceded.

- **Form:** pair discussion (analysis)
- **Tests:** evaluation: reconstructing the presenters' thesis from pieces spread over several sections.
- **Model answer:** Influence (04.2, 09.1, 10.1, 11.1): the DIS 2020 review counted 216 gesture elicitation studies, calls this paper "the first application of the guessability method to hand-gesture input," and lists it as the most cited, 1,070 Google Scholar citations as of August 2019; 20 participants became the field's most frequent sample size; Vatavu and Wobbrock (CHI 2015) reworked the agreement measure; Wobbrock's 2012 whitepaper lists the paper as a methodological contribution. Thin evidence (12.1): the guessability test the authors proposed (p. 1091) was never run; the GI 2010 follow-up tested preference, not guessability (22 participants blind to authorship judged 81 gestures and preferred those authored by larger groups, partly because the researchers' gestures were more physically and conceptually complex); "popularity does, in fact, identify better gestures" (p. 1088) rests on participants rating their own gestures, with no independent judges; and the study measured the repertoire people brought in. Wobbrock and Kientz say methods need "repeated validation" (10.1). Concessions: the designer comparison and GI 2010 are real evidence, the strongest on offer, though neither measures guessability (12.1 Watch out); the self-ratings are not circular, since participants did not know group sizes when rating (09.1, 12.2); and "not a bad paper" (12.1 Say).
- **Why it discriminates:** the paper contains none of the influence figures, none of the follow-up studies and no thesis of this kind. A paper-only reader can cite "goes a long way to validating" (p. 1088) as support but cannot weigh it against the influence.
- **Wrong answers that reveal:**
  - "The method was shown to fail validation." The set was never tested for guessability; nothing failed (09.1 Watch out, pitfall 6).
  - "The 2010 study showed users prefer user-authored gestures." The over-short version we warned against; the finding is about group size and the researchers' more complex gestures (11.1 Watch out).
- **Follow-up probe:** If you had to strengthen the evidence with one study, which would you run first, the reverse guessability test or a recognizer, and why? (p. 1091; 11.1.)

## How to use in class

- **Right after the talk, before the warm-up (2 minutes, written):** Q14.1, Q14.2, Q14.3 and Q14.4 as a quiet one-line check on a slip of paper or a show of hands. These give the presenters a quick read on who followed the talk.
- **Pairs (5 minutes, 12.3 step 2):** Q14.9 and Q14.10 are the pair prompts if the four tensions on 12.2 do not take. Q14.9 pairs naturally with tension 1 (should the most popular gesture win); Q14.10 with tension 3 (does popularity identify better gestures).
- **Group phase, held for a lull:** Q14.5 (lawsuits) reliably restarts discussion; Q14.6 (timing) moves the group from critique to design; Q14.7 and Q14.8 work as probes when someone defends the paper uncritically.
- **Participation grading:** note who answers Q14.5 to Q14.10 with a section-specific point (an outside fact, a "we argue" line or a caveat) rather than the paper-only answer listed under "Wrong answers that reveal." A section-specific point shows engagement with the talk; the paper-only answer shows reading but not listening. Record names against question numbers during the session.


---

# 15 Questions: Both, and worth arguing

**Purpose:** eight questions that set a claim of the paper against a reading from the presentation, so that a good answer needs both and more than one answer can be defended. **Count:** 8 questions. **Discriminates:** students who read the paper and listened to the talk, and can weigh one against the other, from students who can only repeat the paper's claims or only repeat the presenters' critique. **Use:** pair phase (5 min) and group phase (20 min) of the discussion, after the live elicitation warm-up (`12-Discussion-Bridge.md`, slide 12.3). See "How to use in class."

Reading key for the "Feeds" lines. The four tensions on slide 12.2 are: **T1** Should the most popular gesture win? **T2** Did the commands shape the answers? **T3** Does popularity identify better gestures, or gestures people enjoyed inventing? **T4** Is behaviour without feedback the right design target? The three smaller items are: **S1** Floor or ceiling? **S2** Symbolic gestures import writing. **S3** Single users on a group device. **R** is the presenters' reading on slide 12.1.

Each question below gives the strongest answer on each side. Neither side is the "right" one. A good answer picks a side and shows it knows the best case against it.

## Q15.1: Does the designer comparison support the method, or does it only show that twenty people cover more ground than three?

- **Form:** pair discussion; can open with a show of hands ("supports the method" or "only shows 20 beat 3")
- **Tests:** evaluation of evidence; whether the student sees what the yardstick in the comparison is
- **Model answer, side A (it supports the method):** The paper says each author alone came up with only 43.5% of the user-defined gestures, all three combined only 60.9%, and that the gestures are "not, in fact, 'obvious'" (pp. 1090 to 1091). These are three HCI experts who designed their sets before the study (p. 1090). If experts cannot predict what users converge on, then asking users adds information a designer does not have, which is exactly the method's premise ("user behavior is rarely so systematic," p. 1083). The presentation concedes that this comparison and the GI 2010 follow-up are "the strongest the paper and its follow-up offer" (12.1, Watch out). The GI 2010 study adds independent judges: 22 participants blind to authorship judged 81 gestures and preferred those authored by larger groups (04.2 and 11.1). On Olsen's test for "empowering new design participants," the comparison is the paper's argument that existing designer practice is not good enough (10.1, Watch out).
- **Model answer, side B (it only shows that 20 cover more than 3):** The yardstick is the users' own set. The 43.5% and 60.9% count how many gestures of the users' set the designers happened to propose; designer gestures outside that set (the 19.1%) earn no credit, so the designers can at best match the users' set and cannot beat it by construction (05.2, Say and Watch out: "this measures overlap with the users' set, not quality"). The paper's own conclusion is the narrow one: "three experts cannot generate the scope of gestures that 20 participants can" (p. 1090), a statement about coverage. The 19.1% of each author's gestures that no participant tried could be creative or "hopelessly lost" (p. 1090); the comparison cannot tell which, which is slide 12.2's "Floor or ceiling?" question. The test that would settle it, showing unlabeled clips of user and designer gestures to 20 new people (p. 1091), was proposed and not run in the paper (04.2, 11.1, 12.1). GI 2010 tested preference, partly explained by the researchers' gestures being more physically and conceptually complex (04.2), not guessability.
- **Why it discriminates:** a paper-only reader tends to repeat the authors' "wisdom of crowds" framing (p. 1090) and treat 43.5% as a defeat for designers. A talk-only reader repeats "overlap, not quality" but cannot say what the paper itself concedes, or that the paper planned a designer-versus-user guessability test on p. 1091.
- **Weak answer:** "Users beat the designers, so the method works," or the opposite, "it is circular, so it proves nothing." Both ignore that the comparison does show something (experts do not predict consensus) and that what it shows is coverage, not quality.
- **Follow-up probes:**
  1. The validation the authors proposed includes "clips of designers' gestures" (p. 1091). Design the result that would make side B concede.
  2. If the designers' 19.1% were tested and turned out more guessable than the users' gestures, what would that say about elicitation as a floor or a ceiling?
- **Feeds:** T3 and S1 (12.2); also R (12.1, Watch out).

## Q15.2: Does a gesture elicitation study measure the users, or does it measure the method?

- **Form:** pair discussion, then group
- **Tests:** analysis; whether the student can separate a property of the population from a property of the procedure
- **Model answer, side A (it measures the users):** The paper went out of its way to remove method-made bias: a blocks world of 2D shapes, "no elements specific to Windows or the Macintosh," no application domain (p. 1085). Even so, about 72% of gestures were mouse-like, participants tapped imaginary close buttons and taskbars, and one said "I'm a child of the mouse" (p. 1090). A bias that survives the removal of its cue is a fact about the people, and the abstract reports it as a finding: "desktop idioms strongly influence users' mental models" (p. 1083). For a public walk-up table "where the importance of immediate usability is high" (p. 1085), those ingrained habits are exactly what a designer needs to know. The presentation agrees that "unrevised behavior" means unrevised by recognizer feedback, not habit-free, and that the paper does not claim otherwise (04.1, Watch out; 12.1).
- **Model answer, side B (it measures the method):** The output depends on choices the authors made. The 27 referents were "obtained from desktop and tabletop systems" (p. 1085) and include cut, paste, minimize, maximize and task switch, which are desktop concepts (09.1, Watch out; T2). Every participant had to make a two-hand gesture for every referent, which, the presentation argued, is one reason 17 of 48 set gestures are two-handed although two hands were never preferred (p. 1088; 05.2, Say). No feedback meant one participant added fingers "to ensure that I was pressing" (p. 1090; 05.2). Participants could not revise earlier gestures (p. 1091). The screening named only "an interactive tabletop, Apple iPhone, or similar" (p. 1085), and no touch history was recorded, so habits from touchpads, Palm Graffiti or a Nintendo DS could not be measured (06.1, 01.1). The presentation's reading: the study measured the repertoire participants brought in, which is "also a property of the method" (12.1). Greenberg and Buxton's warning that evaluation converges on "that which we are already familiar with" applies here too (10.1).
- **Why it discriminates:** a paper-only reader says "it measures users, the paper removed bias" and stops at p. 1085. A talk-only reader says "it measures legacy bias" without knowing the paper itself reports the desktop pull as a finding, or that the Windows elements were deliberately removed.
- **Weak answer:** "The participants were biased, so the study is invalid," or "unrevised behavior means the participants had no prior habits." The second repeats an error the presentation corrected (pitfall 1; 04.1, Watch out).
- **Follow-up probes:**
  1. In the warm-up we just ran, which gestures in this room came from smartphones? Does that make our agreement score a fact about us or about the prompt we chose?
  2. Name one change to the procedure that would separate the two readings. (Section 11 suggests recording prior device exposure as a covariate; Morris et al. 2014 propose production, priming and partners.)
- **Feeds:** T1 and T2 (12.2); R, third support (12.1).

## Q15.3: Does the 2010 preference study close the validation gap the paper left open?

- **Form:** short answer (paragraph); works as a one-minute written check if discussion stalls
- **Tests:** comprehension of what "validation" meant in the paper, and evaluation of partial evidence
- **Model answer, side A (it largely closes it):** The weakest link in the paper's own validation is that goodness was self-rated: the authors say larger-group gestures got higher ratings and that this "goes a long way to validating" the approach (p. 1088), but the raters were the gestures' authors (09.1; 12.1). GI 2010 fixes exactly that weakness: 22 new participants, blind to authorship, judged 81 gestures and preferred those authored by larger groups (04.2, 11.1). That is independent judges confirming the same direction as the self-ratings. Section 09 notes this explicitly ("The GI 2010 follow-up added independent judges," 09.1, Watch out).
- **Model answer, side B (it does not):** The paper defined the next step precisely: show unlabeled clips to 20 new participants and see whether they "can guess which gestures perform which commands," reversing the study "from signs to referents" (p. 1091), then build a recognizer and measure recognition rates (p. 1091). GI 2010 measured preference, not guessability, and the preference was partly explained by the researchers' gestures being more physically and conceptually complex (04.2; pitfall 6). The paper itself shows that more complex referents went with lower goodness ratings (r = -0.59, p. 1088). That correlation concerns the complexity of the commands, as rated by the authors, not of the gestures, so it only hints at the rival explanation that "people prefer simpler gestures"; it does not show it. Neither the guessability test nor the recognizer appears in this paper, and the presentation did not check whether either was later run for this set (11.1, Watch out). Hence the presentation's verdict: "partial support" (12.1).
- **Why it discriminates:** the paper alone does not mention GI 2010 at all. The talk alone gives the GI 2010 result but not the exact form of validation the authors promised on p. 1091 or the complexity correlation on p. 1088 that offers the rival explanation.
- **Weak answer:** "Users preferred user-authored gestures, so the set is validated." This is the shortened version the presentation warns against (04.2, 11.1; pitfall 6), and it confuses preference with guessability.
- **Follow-up probes:**
  1. Suppose the reverse test were run and the user-defined set were guessed only slightly better than the designers' sets. Would that rescue or sink the method?
  2. Is preference ever enough evidence for a walk-up public table, where people use a gesture once?
- **Feeds:** T3 (12.2) and R, first support (12.1).

## Q15.4: Do Greenberg and Buxton excuse the missing validation?

- **Form:** short answer, then group
- **Tests:** application of a course reading to this paper; evaluation of the paper's own framing of its contribution
- **Model answer, side A (yes, largely):** Greenberg and Buxton argue that evaluating an early design can "mute creative ideas that do not conform to current interface norms" (CHI 2008, p. 111; 10.1, Watch out). Wobbrock and Kientz say artifacts "are often accompanied by empirical studies but do not have to be, and sometimes should not be," citing Greenberg and Buxton (Interactions 2016, p. 40; 10.1, log). The paper's abstract presents "an approach to designing tabletop gestures" (p. 1083), and Wobbrock's 2012 whitepaper lists this paper as a methodological contribution (10.1). Read that way, the main contribution is the method, the set is its first output, and the paper is honest that validating the set is "an important next step" (p. 1091). Demanding a full guessability study in the same paper would be the premature evaluation Greenberg and Buxton warn about.
- **Model answer, side B (no, the excuse only partly applies):** The paper does not frame the set as an early sketch. It calls it "a good candidate for deployment in tabletop systems" with "ease of recognition" (p. 1091), although no recognizer was built (08.1; 09.1). That is a late-stage claim, so the early-stage excuse does not cover it (10.1, Watch out). For the method, Wobbrock and Kientz say methodological contributions need "repeated validation" (p. 41; 10.1), and this paper has one study whose validation rests on self-ratings (p. 1088). And Greenberg and Buxton's central warning, convergence on "that which we are already familiar with" (p. 115), cuts against elicitation itself: about 72% of gestures were mouse-like (p. 1090; 10.1).
- **Why it discriminates:** the paper alone gives the "good candidate for deployment" and "next step" sentences but not the course readings. The talk alone gives the reading quotes but a student needs the paper's own Conclusion wording to judge whether the set was presented as early-stage.
- **Weak answer:** "Greenberg and Buxton say evaluation is harmful, so no validation was needed." This misreads the title's "some of the time" and ignores that the reading is about usability evaluation, not elicitation (10.1, Watch out).
- **Follow-up probes:**
  1. Which of the paper's four listed contributions (p. 1084) would Greenberg and Buxton excuse, and which would they not?
  2. Greenberg and Buxton list "participatory critique" among other ways to validate (p. 118; 10.1). Is elicitation itself a form of participatory critique, or is it something else?
- **Feeds:** T3 (12.2) and R (12.1); connects to section 10.

## Q15.5: Is a sample of 20 a weakness of this paper, or the field's norm?

- **Form:** quick check (show of hands: weakness or norm), then short answer
- **Tests:** evaluation of sample size against the claims made; whether "norm" and "sufficient" are kept apart
- **Model answer, side A (it is the norm, and enough for the purpose):** Twenty was enough to show strong consensus where it existed: move a little and move a lot reached the maximum agreement of 1.0 for one hand (Figure 3, p. 1088), far above the floor of about 0.05 that twenty different gestures would give (p. 1087; 04.2). The DIS 2020 review of 216 elicitation studies found an average of 25 participants (SD 4) and 20 as the most frequent choice, in 36% of studies (09.1, Watch out). The authors even planned their validation with 20 new participants (p. 1091). For a method that looks for the largest group, the sample size matches the job.
- **Model answer, side B (it is a weakness, and the norm is circular):** The same review says 20 is common "likely" because Wobbrock et al. employed 20 (09.1), so the norm partly comes from this paper and cannot justify it. "Norm is not the same as sufficient" (09.1). Twenty is too few to split by age or culture (09.1, Say), and the paper reports no result by age and no ethnicity (06.1). The tail of Figure 3 sits at about 0.11 to 0.21 (p. 1088; chart reading), where, we argue, a few people changing their gesture would reorder the referents. The chi-square tests treat 1,074 gestures from 20 people as independent (pp. 1087 to 1088; 07.2, 09.1). And the sample was "all non-technical literate American adults" (p. 1091), so twenty more of the same would not fix the culture problem.
- **Why it discriminates:** the paper alone gives n = 20 and the agreement values but not the DIS 2020 figures or the self-reference. The talk alone gives the DIS 2020 numbers but a student needs Figure 3 and Eq. 1 to argue what twenty can and cannot show.
- **Weak answer:** "20 is too small" with no reason, or "20 is standard so it is fine." Both skip what the sample is being asked to show (strong consensus versus subgroup differences).
- **Follow-up probes:**
  1. Which conclusion in the paper would most likely change with 200 participants of the same kind, and which with 20 participants from a different culture?
  2. If a later study adopts n = 20 because this paper did, who carries the burden of showing it is enough?
- **Feeds:** T1 (12.2, consensus from one group of novices) and S2; links to sections 06 and 09.

## Q15.6: Do culture-bound marks disqualify the gesture set, or only its symbolic part?

- **Form:** pair discussion, then group
- **Tests:** analysis of where culture enters the set; evaluation of how far a limitation reaches
- **Model answer, side A (only the symbolic part, and the paper already hedges it):** The culture-bound marks are a small, identifiable part of the set: accept "draw check," reject "draw 'X'," help "draw '?'" (Figure 4, p. 1089). Symbolic gestures were .100 of all gestures (Figure 2, p. 1087), and the set's taxonomy breakdown is close to that of all proposals, an average difference of 6.7 percentage points (p. 1088). Physical gestures were the largest nature category, 43.9% of all proposals (Figure 2, p. 1087; p. 1091), and many gestures in the set are direct manipulation: drag, pinch, pull apart, drag offscreen (Figure 4, p. 1089). The paper gives no nature breakdown for the set itself, only the 6.7-point average difference. The paper itself suggests widgets for "delete, zoom in, zoom out, accept, reject, menu access, and help" (p. 1091), so the culture-loaded referents already have a fallback. The fix is local: swap the marks per locale, as the presentation's Japanese circle-versus-check example implies (06.1).
- **Model answer, side B (culture reaches further than the marks):** The presentation's point is that culture also sits in gestures that look physical or neutral. Next is drawn left to right in Figure 4, which we read as reading direction (06.1, Watch out; an inference). Close and task switch were imaginary Windows widgets at objects' top right and the screen's bottom (p. 1090), a software culture. Minimize means drag to the bottom of the surface (Figure 4, p. 1089), which we read as echoing the Windows taskbar at the screen's bottom that participants imagined for task switch (p. 1090); that link is our inference. (Minimize2's "as defined by user's seating position" is a multi-user issue, S3, not a cultural one.) More basically, the agreement score itself was computed within one culture, so what counts as "consensus" is cultural; the authors say "children, Eastern, or uneducated participants would behave differently" (p. 1091) and that culture could affect gesture choice (p. 1091). The meaning of a movement is learned, as the Bulgarian head nod shows (06.1). A later cross-country study reportedly found more similarity than difference except in symbolic gestures (06.1) [UNVERIFIED: Mauney et al., CHI 2010 EA, findings known only from a secondary summary], which would support side A for physical gestures and side B for symbols.
- **Why it discriminates:** the paper alone gives the limitation sentence and Figure 4 but no cultural examples or the reading-direction argument. The talk alone gives the Japan and Bulgaria examples but not the Figure 2 proportion or the p. 1091 widget list that limits the damage.
- **Weak answer:** "The set is Western, so it is useless," or "the authors admitted it, so it does not matter." The first ignores that only the symbolic marks are plainly culture-bound and that physical gestures were the largest category; the second treats acknowledging a limit as removing it (06.1, Watch out: "Our point is the consequence for the set, not that the authors missed it").
- **Follow-up probes:**
  1. Which referents in Figure 4 would you expect to keep their gesture in a Japanese replication, and which to change? Give a reason for each.
  2. Should a cross-cultural set be one set with local aliases, or separate sets per culture? What does aliasing (p. 1088) make possible here?
- **Feeds:** S2 and S3 (12.2); T2 (who chose the referents and their framing).

## Q15.7: Does Microsoft's patent application on the method change how the paper should be read?

- **Form:** group discussion only (needs care; not for pairs)
- **Tests:** evaluation of context against content; separating documented fact from attributed motive
- **Model answer, side A (no, judge the paper on what it reports):** The method is published in full at CHI, with the agreement formula and worked examples (p. 1087), and it builds openly on academic work: Good et al. 1984 and Wobbrock's own 2005 guessability study (pp. 1084, 1087; 04.1). Anyone can and did replicate it: a 2020 review counted 216 elicitation studies (04.2, 11.1). The application was abandoned in 2014 and never granted (03.1). The presentation itself found "no evidence" that the study was run to avoid lawsuits, and says the filing "is consistent with offensive, defensive or routine corporate IP practice" (03.1, Say and Watch out). The evidence in the paper stands or falls on its own.
- **Model answer, side B (yes, as context for how claims are weighed):** Two of three authors were at Microsoft Research (p. 1083), the study ran on a Microsoft Surface prototype (p. 1085), and Microsoft filed an application on this very procedure on 2008-08-04, naming the three authors as inventors (03.1, 04.1). The paper does not mention it (03.1, 09.1). That does not make the results wrong, but it is a reason to read the strongest framing sentences more carefully: "the first to employ users, rather than principles" (p. 1084), "goes a long way to validating" from self-ratings (p. 1088), and "a good candidate for deployment" with "ease of recognition" and no recognizer (p. 1091). It also sharpens T3's question of what independent validation would look like: independent of the authors' employer as well as of the gesture makers.
- **Why it discriminates:** the patent application is not in the paper; only the talk gives it (03.1, 04.1). But judging whether it matters needs the paper's own claims and affiliations, and the paper's academic lineage on p. 1084.
- **Weak answer:** "They patented it, so the study was done to make money" or "to avoid lawsuits." Both attribute a motive as fact, which the presentation explicitly refuses (03.1, Watch out). Also weak: calling the application "a patent"; it was never granted (03.1).
- **Follow-up probes:**
  1. Which sentence in the paper, if any, would you reword knowing about the filing, and why that one?
  2. Industry labs produce much of HCI's systems research. What disclosure, if any, should a paper like this make, and would it change your reading of the evidence?
- **Feeds:** R (12.1) and T3 (12.2). No tension on slide 12.2 is about the patent directly; see Open items.

## Q15.8: From the paper's timing data, the presentation argued that thinking took several times longer than doing. Does that change what a walk-up table should optimise?

- **Form:** pair discussion; good closer for the group phase
- **Tests:** analysis of the timing data and its conditions; transfer to a design decision
- **Model answer, side A (yes, optimise for finding the gesture):** Planning time rose with complexity, about 8 s for simple and about 15 s for complex referents (r = 0.71), while complexity did not change articulation time (p. 1088). Gestures rated easy took about 3.4 s and difficult about 2.0 s (p. 1088), so execution is short in either case; thinking is several times doing (07.2). Long planning went with lower goodness: "good gestures were those most quickly apparent to participants" (p. 1088). For a public table "where the importance of immediate usability is high" (p. 1085), the cost is in finding the gesture, not making it (07.2, Say). The paper's own advice points the same way: widgets as well as gestures beyond the first seven referents, and widgets for delete, zoom in, zoom out, accept, reject, menu access and help (p. 1091). Agreement also falls with complexity, r = -0.52 (p. 1087).
- **Model answer, side B (no, or not from this evidence):** The participants were inventing a gesture from an animation, alone, with no feedback (p. 1085). A walk-up user faces a different task: guessing the designer's gesture, which is the reverse study the paper proposed and did not run (p. 1091). Planning time here measures invention, not discovery. After first use the cost shifts to execution, and the presentation's motor reading says "hard means precise, not long" (07.2), which points to optimising precision tolerance, as the paper also recommends ("demands for precise positioning should be avoided," p. 1091). The numbers also need care: the "several times" ratio crosses two different splits of the data, by complexity and by ease rating, and runs from about 2.4 to 7.5 (07.2, Watch out); the means come with no spread, and the time-on-rating tests treat 1,074 gestures as independent (p. 1088; 07.2). And because there was no feedback, people did not learn, so the data say nothing about the second use (T4).
- **Why it discriminates:** the paper alone gives the numbers but never sets thinking against doing or draws a walk-up conclusion; that synthesis is the presentation's (07.2). The talk alone gives the conclusion, but defending or attacking it needs the paper's procedure (p. 1085) and the reverse-study proposal (p. 1091).
- **Weak answer:** "Users are slow, so gestures should be simpler," or quoting 8 s versus 2 s as a precise ratio. The first misses the difference between planning and execution; the second ignores that the two numbers come from different splits.
- **Follow-up probes:**
  1. What would you measure on a deployed table to tell discovery cost from invention cost?
  2. If easy gestures are long and continuous, should a walk-up table favour continuous gestures even for discrete commands like accept and help? What does Figure 2's flow split (.535 discrete, .465 continuous, p. 1087) suggest?
- **Feeds:** T4 and T1 (12.2: guessability against learnability against efficiency).

## How to use in class

**Warm-up (5 min, slide 12.3).** No question from this set. The live elicitation feeds Q15.2: note on the board any smartphone-era gestures the room produces (12.3, Watch out), and come back to them in the group phase.

**Pair phase (5 min).** Slide 12.3 asks pairs to pick one of the four tensions and take opposite sides. Hand each pair a card with the matching question and sides already assigned (side A, side B), so nobody has to choose:
- T1: Q15.2 (users or method)
- T2: Q15.2 or Q15.6 (culture), alternate between pairs
- T3: Q15.3 (GI 2010) or Q15.1 (designer comparison), alternate
- T4: Q15.8 (thinking versus doing)

Each pair writes down one piece of evidence from the paper (with a page) and one from the talk (with a section) for its side. This is what we collect for grading.

**Group phase (20 min).** Order, assuming pairs used Q15.1, Q15.2, Q15.3, Q15.6, Q15.8:
1. Open with two opposed pairs from the same card reporting (about 6 min). Q15.2 is the best opener because it ties to the warm-up.
2. Q15.5 (sample size) as a show of hands, then one reason per side (about 3 min). It moves quickly and links culture (Q15.6) to statistics.
3. Q15.4 (Greenberg and Buxton) to bring in the course readings, which the grading rubric rewards (about 4 min).
4. Q15.7 (patent) only with the whole group and only if time allows (about 4 min). Restate the rule first: documented facts on one side, interpretation labelled on the other, no motive as fact.
5. Close on Q15.8 if it was not a pair card, or on its second probe (about 3 min).

**For a lull.** Use the first follow-up probe of whichever question is live; probes are written to be asked as they stand. Q15.3 also works as a one-minute written answer.

**Grading participation.** For each contribution note the student, the question, and three marks: cited the paper with a page, cited the talk with a section, acknowledged the strongest point of the other side. A student who does all three on one question has shown both reading and listening. Weak answers listed above are not penalised if the student revises after a probe; note the revision. Credit classmates whose commentaries anticipated a side by name on slide 12.3 (**[classmate]**).
