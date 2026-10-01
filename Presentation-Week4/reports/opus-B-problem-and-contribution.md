# Opus report B: problem and contribution

Independent report by a second Claude Opus agent, 2026-09-29, on questions 1 and 2 of the course rubric for Wobbrock, Morris & Wilson, "User-Defined Gestures for Surface Computing" (CHI 2009). Same brief as report A, written without sight of it. The GI 2010 follow-up it mentions under D was cited from memory and has not been verified.

---

Report on Q1 (problem) and Q2 (contribution): Wobbrock, Morris & Wilson, "User-Defined Gestures for Surface Computing," CHI 2009

I read the whole paper, including the references, and looked at Figures 2 and 3 as rendered pages. Page numbers below are the printed ones (1083 to 1092).

**A. The problem**

**As the authors state it.** Most surface gestures "have been defined by system designers, who personally employ them or teach them to user-testers," which yields "somewhat arbitrary gesture sets whose members may be chosen out of concern for reliable recognition" (p.1083). That criterion "is not useful for determining which gestures match those that would be chosen by users" (p.1083).

**The deeper problem.** The field had no way to find out what users would actually do, so the vocabulary was designed from the recognizer's side. The related work section is built to show this.
- Systems like RoomPlanner and Wu et al.'s gesture framework "were not elicited from users, although usability studies were conducted" (p.1085). Users only ever evaluated gestures that designers had already chosen.
- Malik et al. told 1-, 2-, 3- and 5-finger gestures apart, "a feature our current findings suggest may be problematic for users" (p.1085). Designers used finger count because a sensor can count fingers, not because people care about it.
- Principled approaches (Wu, Poggi) exist, but the authors claim to be "the first to employ users, rather than principles, in the development of a gesture set" (p.1084).
- Gesture is open-ended: "almost anything one can do with one's hands could be a potential gesture" (p.1083), unlike a mouse, which "inherently constrain[s] human motion." A large design space combined with designer-chosen gestures makes the vocabulary problem much worse.

**Why it mattered then.** From the paper: tabletops were spreading across research uses (single users, groups, multi-display setups, entertainment; p.1083). The study ran on a Microsoft Surface prototype (p.1085). The authors screened out anyone who had used "the Apple iPhone" (p.1084), which shows that consumer multi-touch was already shaping how people expect touch to work. They call the timing "timely" (p.1083). They also point out that tables "may be used in public spaces, where the importance of immediate usability is high" (p.1085).

My own knowledge, not in the paper:
- The iPhone (2007) made pinch-to-zoom mainstream.
- Jeff Han's 2006 TED demo made multi-touch widely known.
- Microsoft Surface shipped commercially in 2008.
- Windows 7, with native multi-touch, came out in 2009.
- Two of the three authors were at Microsoft Research, which had a direct stake in the question.

So the gesture vocabularies were about to become fixed in commercial products. It was the last cheap moment to ask users.

**Why it still matters.** Any input channel that is not physically constrained has the same problem: many possible actions and no obvious mapping to commands. Concrete cases:
- mid-air gestures (Kinect, Leap Motion)
- AR/VR hand input (HoloLens, Quest and Vision Pro pinch gestures)
- smartwatch and trackpad gestures
- voice command phrasing
- in-car gestures
- today, deciding what users will naturally say or do with LLM agents

My own knowledge: this paper started a whole genre of "gesture elicitation studies," and hundreds of them followed in these domains.

**Slide phrasings:**
1. "Tabletop gestures were designed for the recognizer, not the person."
2. "When anything your hands do could be a command, who decides which gesture means what?"

**B. Main idea and contribution**

**Core idea in one sentence.** Show non-technical users the effect of a command (the referent), ask them to perform the gesture they think caused it (the sign), then build the gesture set from where users agree (p.1083, p.1085).

Why it is designed this way:
- **Elicit from users.** Designers organize gestures "in a principled, logical fashion," but "user behavior is rarely so systematic" (p.1083). Users are "not designers," so their behavior has to be elicited in a form that is "profitable for design" (p.1085).
- **Show the effect first.** This fixes the meaning and leaves only the form open. Every participant answers the same question, so their gestures can be compared and agreement can be counted.
- **Remove the recognizer.** They "sought to remove the gulf of execution ... creating, in essence, a monologue in which the user's behavior is always acceptable" (p.1085). With no recognition feedback, users never learn to adapt to what the system can detect, so the study sees "unrevised behavior."

**Contributions as listed (p.1084):**
1. a quantitative and qualitative characterization, including a taxonomy
2. a user-defined gesture set
3. insight into mental models
4. implications for technology and design

**My sorting of what they produced:**
- **Method:** referent-then-sign elicitation on a surface, with think-aloud, logging, and 1- versus 2-hand trials.
- **Artifact:** 48 gestures for 22 referents, conflict-free, covering 57.0% of proposed gestures (p.1087, p.1088).
- **Measurement instruments:** the agreement score A (Eq. 1, p.1087) and a four-dimension taxonomy (form, nature, binding, flow; p.1086).
- **Empirical findings:**
  - Finger count rarely matters.
  - One hand is preferred for 25 of 27 referents (p.1088).
  - About 72% of gestures were mouse-like (p.1090).
  - Agreement falls as command complexity rises (r = -.52, p.1087).

**What matters most.** The method, together with agreement as the thing to measure. Nobody ships this 48-gesture set; it came from 20 Americans in a world of 2D blocks. What spread was the procedure, which can be carried to any modality, and the idea that a design team can measure consensus and then decide: give a command a gesture where agreement is high, and a widget where it is low. The paper itself says the set is useful "not just for what it contains, but also for what it omits" (p.1088). The four listed contributions leave the method out, which is odd given that it had the biggest lasting impact.

**Novel versus borrowed:**
- **Good et al. 1984** used referent-to-sign elicitation for a command-line email interface, with text commands and a hidden operator (p.1084).
- **Wobbrock et al. 2005** introduced guessability and the agreement score for EdgeWrite unistrokes. That covered one stroke with one finger (p.1084, p.1087, "adopted from prior work [33]").
- **Nielsen et al. 2004** described "a similar approach" as a procedure for gesture interfaces (p.1084). The paper does not say what differs.
- **Epps et al. 2006** is the nearest precedent: tabletop hand shapes elicited from static Windows screenshots. Its own authors admitted the prompts may have led users to "simply emulate the mouse" (p.1084).

What this paper adds:
1. The method applied to multi-touch, two-handed, continuous input, where the design space is far larger.
2. Animated, application-neutral prompts in a blocks world instead of Windows imagery (p.1085).
3. A complete, conflict-resolved gesture set rather than frequency counts.
4. A taxonomy of surface gestures grounded in user behavior (p.1086).
5. A direct comparison with the designers' own gestures (p.1090).

Beringer (2002) and Voida (2005) had also elicited gestures (p.1084). So "first to employ users" really means first to build a complete set from users.

**The one finding to remember.** Working independently, each author came up with only 43.5% of the gestures users defined, and all three together covered 60.9%. Meanwhile, 19.1% of each author's gestures were never tried by any participant (p.1090). I chose this because it is direct evidence for the problem in Part A: expert HCI designers, including the very people running the study, cannot predict what users will do. It sums up the paper, and the authors' own joke ("thinking creatively or ... hopelessly lost!") makes it memorable. A close second is that the 72% mouse-like gestures persisted even after all desktop imagery was removed (p.1090).

**C. Discussion questions**

1. **Should the most popular gesture win?** Legacy bias: 72% of gestures were mouse-like, and one participant said "I'm a child of the mouse" (p.1090). Majority agreement may simply lock in old habits and make it impossible to discover better, novel gestures. Is guessability the right goal for a device people will use every day, rather than learnability or efficiency?
2. **Is the validation circular?** The claim that "popularity does, in fact, identify better gestures" (p.1088) rests on goodness ratings that participants gave to their own gestures, right after inventing them. Does that validate the method, or just show that people like what they thought of? The authors admit the set is not yet validated (p.1091).
3. **Can "unrevised behavior" be the right design target?** Removing feedback removes bias, but real use is a dialogue, and people adapt quickly once a system responds. Participants gestured in the air, above the table, and off-screen (p.1090). Which should designers follow: what users do with no feedback, or what they settle into once they have it?

**D. Surprises and likely mistakes**

- **Agreement is not the same as good.** A is a consensus measure, and its minimum is 1/|Pr| (0.05 here), not 0 (p.1087). The formula counts each gesture as agreeing with itself. My own knowledge: Vatavu and Wobbrock (CHI 2015) replaced it with a corrected "agreement rate" for this reason. Do not present A = 0.32 as "32% agreed."
- **Figure reference error.** On p.1091 the text says "Figure 2 indicates that agreement is low after the first seven referents." It means Figure 3.
- **Mismatched worked example.** The move a little example (0.42) is the 2-hand score, while Figure 3 shows 1-hand move at 1.0. Say which one you are quoting.
- **Not every command got a gesture.** Four referents got no gesture. For close and task switch, users tapped imaginary widgets, which is itself a finding in favour of widgets.
- **Pinch was still elicited.** Even though iPhone users were screened out, "pinch" and "reverse pinch" appeared (Figure 4, p.1089; p.1091). Whether this reflects prior exposure is worth raising.
- **Narrow sample.** 20 right-handed American adults with an average age of 43 (p.1085). The authors flag this themselves (p.1091).
- **Heads-up for the discussion.** Some classmates may know the later GI 2010 follow-up (Morris, Wobbrock and Wilson), which I believe found that users preferred user-authored gestures to designer-authored ones. That is from memory; check it before citing.
