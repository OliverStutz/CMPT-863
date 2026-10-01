# Opus report A: problem and contribution

Independent report by a Claude Opus agent, 2026-09-29, on questions 1 and 2 of the course rubric for Wobbrock, Morris & Wilson, "User-Defined Gestures for Surface Computing" (CHI 2009). Brief: read the whole paper, ground claims in page numbers, do not summarize section by section. Written without sight of report B.

**Correction after verification:** the claim under D that the paper misprints the range of A as "[|Pr|-1, 1]" is wrong. The paper prints the -1 as a superscript, so the range is 1/|Pr| to 1, which is correct. The apparent error came from text extraction. Everything else stands as the agent wrote it.

---

Report: Wobbrock, Morris & Wilson, "User-Defined Gestures for Surface Computing" (CHI 2009). I read all ten pages (1083 to 1092), including related work, discussion, limitations and references.

**A. The problem**

**As the authors state it.** "To date, most surface gestures have been defined by system designers," which produces "somewhat arbitrary gesture sets whose members may be chosen out of concern for reliable recognition" (p.1083). The authors say this is fine for early prototypes but "not useful for determining which gestures match those that would be chosen by users" (p.1083).

**The deeper problem.** The field was building gesture vocabularies from the recognizer outward, not from the user inward. The related work sorts prior systems in exactly this way:
- RoomPlanner and Wu et al.'s registration/relaxation/reuse gestures "were not elicited from users, although usability studies were conducted" (p.1084 to 1085). The users only tested the gestures. They never generated them.
- Malik et al. "distinguished among 1-, 2-, 3-, and 5-finger gestures, a feature our current findings suggest may be problematic for users" (p.1085). Counting fingers is easy for a vision system but means nothing to most people.
- Even the "principled" approaches [20,35] use principles rather than people (p.1084).
- Long et al. showed that designers are poor at picking gestures that are easy to tell apart (p.1084).

So the underlying issue is methodological. The field had no systematic, measurable way to find out what gestures people would produce on their own. Gesture sets were "designed" the way command names were designed before Furnas et al.'s vocabulary problem [8]: one designer's guess stood in for the population.

**Why it mattered then.** From the paper: tabletops were spreading across individual, group, multi-display and entertainment uses (p.1083). Keyboard and mouse were "no longer preferable" (p.1083). And surface gestures are unconstrained: "almost anything one can do with one's hands could be a potential gesture" (p.1083 to 1084). The authors also stress public-space tabletops, "where the importance of immediate usability is high" (p.1085). They exclude anyone who has used "the Apple iPhone" (p.1084), which shows they knew consumer touch was about to shape user expectations.

My own knowledge (not in the paper, flag it on the slide):
- The iPhone shipped in 2007.
- Microsoft Surface was announced in 2007 and sold commercially from 2008.
- Jeff Han's FTIR multi-touch demos ran from 2005 to 2006.
- Windows 7 with built-in multi-touch was in development through 2009.

Gesture vocabularies were being fixed commercially right then, so conventions set by engineers could have become permanent. Also note that two of the authors are at Microsoft Research and the study ran on a Surface prototype. That is a stake worth mentioning.

**Why it still matters.** The problem is general. Whenever a new input channel has no established vocabulary, someone has to decide which action maps to which command. Concrete cases:
- mid-air gestures (Kinect, Leap Motion)
- AR/VR hand tracking (Quest, Vision Pro pinch)
- smartwatch and ring gestures
- in-car gestures
- voice-assistant command phrasing
- deformable or on-body interfaces

The paper's own caveat that off-table gestures few tables could detect (p.1090) already points toward mid-air input. My own knowledge: this paper started a large genre of "gesture elicitation studies." Later reviews count well over 200 of them, and Vatavu & Wobbrock (CHI 2015) revised the agreement measure.

**Slide phrasings:**
- "Gesture sets were designed for the recognizer, not for the user. Nobody had asked what people would actually do."
- "Who should define a gesture language: the engineer who has to recognize it, or the person who has to guess it?"

**B. The main idea and contribution**

**Core idea in one sentence.** Show non-technical users the effect of a command, ask them to perform the gesture that would cause it, and let the agreement among many users decide the gesture set.

**The reasoning:**
- **Why elicit from users:** "users are not designers" (p.1085), but they are the people who have to guess the gestures. Consensus is a way to predict how guessable a gesture is.
- **Why show the effect first:** This is referent-to-sign elicitation, borrowed from linguistics via McNeill (p.1085, footnote). Showing the result without naming a gesture avoids suggesting one. The animations used a neutral "blocks world" with no Windows or Mac elements (p.1085), because Epps et al.'s Windows screenshots biased people toward mouse emulation (p.1084).
- **Why remove the recognizer:** The authors "did not want the vicissitudes of gesture recognition to influence users' behavior." Removing the gulf of execution turns the dialogue into "a monologue in which the user's behavior is always acceptable," so they can observe "users' unrevised behavior" (p.1085). They state the cost in the limitations (p.1091).

**Contributions as the authors list them (p.1084):**
1. A quantitative and qualitative characterization of user-defined gestures, including a taxonomy.
2. A user-defined gesture set.
3. Insight into users' mental models.
4. Implications for technology and UI design.

**My sorting of these:**
- **Method:** effect-then-cause elicitation with no recognizer, then conflict resolution by largest group.
- **Artifact:** a set of 48 gestures for 22 referents, with aliases. It covers 57.0% of all proposed gestures (p.1087 to 1088).
- **Measurement instruments:** the agreement score A and a four-dimension taxonomy (form, nature, binding, flow; Table 2, p.1086).
- **Empirical findings:**
  - number of fingers rarely matters
  - one hand preferred for 25 of 27 referents
  - about 72% of gestures were mouse-like
  - agreement falls as command complexity rises (r = -.52)
  - off-screen and above-table gestures appear

**What matters most, in my view,** is the method packaged as a reusable, quantified procedure. The paper's own treatment of the gesture set undercuts it as the lasting contribution: it was never validated and is left as the "important next step" (p.1091). The taxonomy has been cited but is situational. The method is what hundreds of later papers copied. The key evidence for the method is on p.1088: gestures in larger agreement groups got higher goodness ratings, "indicating that popularity does, in fact, identify better gestures," which the authors say "goes a long way to validating this user-driven approach."

**Novel versus borrowed:**
- **Wobbrock et al. 2005 (EdgeWrite, [33]):** The referent/sign procedure, the agreement score ("adopted from prior work [33]," p.1087) and the largest-group assignment all come from here. That work was a 4-page extended abstract on unistroke alphabet symbols. This paper adds: whole-hand, two-dimensional, continuous and bimanual gestures; the 1-hand versus 2-hand manipulation; think-aloud mental models; timing and Likert measures tied to agreement; a taxonomy; and a comparison with designers.
- **Good et al. 1984 [9]:** A "user-derived" command-line email interface, built by collecting users' unconstrained commands. Same principle, but it was typed language and it evolved through iteration with a hidden operator. This paper moves the idea to gestures and makes the consensus measurable.
- **Nielsen et al. 2004 [19]:** A procedure for "intuitive and ergonomic" gesture interfaces, described as "a similar approach" (p.1084). The paper does not say how it differs. Be careful not to overclaim here.
- **Epps et al. 2006 [5]:** The closest prior study. It was tabletop hand-shape elicitation from static Windows images, and the authors of that study admitted mouse bias. This paper adds animated effects, neutral content, a full gesture set and agreement scores.

The novelty claim, "the first to employ users, rather than principles, in the development of a gesture set" (p.1084), is therefore narrow. Epps, Beringer and Voida had already elicited gestures. What is new is producing a complete, conflict-free set from user consensus.

**The one finding to remember:** before the study, the three authors designed their own gestures for the same 27 commands. Each author matched only 43.5% of the user-defined set, and all three combined covered only 60.9%. Meanwhile, 19.1% of each author's gestures were never tried by any participant (p.1090). I picked this because it tests the paper's own premise directly: expert designers, including these HCI experts, do not predict what users do.

A second option: finger count. Of the 13 participants who varied finger count, only 2 said it mattered (p.1090). That contradicts several prior systems outright.

**C. Discussion-lead questions**

1. **Is agreement the right target?** Average agreement was low (A = 0.32 for one hand, 0.28 for two; p.1087). About 72% of gestures copied the mouse (p.1090), and participants said "I'm a child of the mouse." Does eliciting gestures from novices lock in legacy habits rather than find better interactions? Is a designer's 19.1% of "never tried" gestures creativity or noise (p.1090)?
2. **Did the commands shape the answers?** The 27 referents were taken "from desktop and tabletop systems" (p.1085): cut, paste, minimize, maximize, task switch, undo. Can a study that asks for desktop commands conclude "it's a Windows world"? Who should choose the referents?
3. **Does removing the recognizer produce valid data, or data that cannot be built?** With no feedback, people gestured in the air and off-screen (p.1090). The paper still claims the set has "ease of recognition" (p.1091), even though no recognizer was built. Is "unrevised behavior" the right design input when real users adapt to feedback within minutes?

**D. Surprises and likely pitfalls**

- **Do not present the gesture set as validated.** Validation is explicitly future work (p.1091).
- **Do not credit this paper with inventing the agreement score.** It came from [33].
- **Coverage cuts both ways.** "Covers 57.0%" (p.1087) also means 43% of what users proposed is not accommodated.
- **The user set is 35.4% two-handed** (p.1088), even though no referent had a majority preference for two hands.
- **Grouping into "identical" gestures was a judgment call by the authors.** The paper does not report inter-rater reliability for it or for the taxonomy coding.
- **Minor errors in the paper** worth knowing so they do not trip the student up:
  - The discussion says "Figure 2 indicates that agreement is low after the first seven referents" (p.1091). It means Figure 3.
  - The stated range of A, "[|Pr|-1, 1]" (p.1087), should read 1/|Pr| to 1. [See correction at the top of this file: the paper is right, the extraction was wrong.]
- **The sample is narrow:** 20 right-handed American adults, mean age 43. The authors acknowledge that culture and age would change results (p.1091).
- **The four-finger "gray area"** (p.1090) is a nice, specific detail if the student wants a concrete takeaway for the audience.
