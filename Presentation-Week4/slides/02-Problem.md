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

**Speaker notes**
- **Say:** In the authors' words, most surface gestures were "defined by system designers," producing "somewhat arbitrary gesture sets whose members may be chosen out of concern for reliable recognition." The paper grants "skillful design." We argue the deeper problem is direction. Vocabularies were built from the recognizer outward. Wu's tabletop gestures were tested for usability, but never elicited from users. Malik's system told 1, 2, 3 and 5 fingers apart, which a camera does easily and, the paper suggests, people may not. And almost anything a hand can do is a candidate. Furnas showed in 1987 that two people pick the same word for the same thing less than 20 percent of the time. One designer's guess is a weak bet. Users had been asked before, by Good, Epps, Nielsen. What was missing was a full surface gesture set built from their agreement. And it mattered: public tables need immediate usability, and products were fixing conventions right then.
- **Transition:** If conventions were being fixed in products right then, the next question is who was fixing them, and whether Microsoft was under pressure.
- **Watch out:**
  - Do not say "nobody had asked users." The paper itself lists Good et al. 1984, Nielsen et al. 2004, Epps et al. 2006, Beringer 2002 and Voida et al. 2005 (p. 1084). The "first" claim is narrower: first to build a gesture set from users "rather than principles," where principles refers to Poggi and Wu et al. (refs. 20, 35, p. 1084).
  - Long et al.: the paper says "users are sometimes poor at picking easily differentiable gestures" (p. 1084), not designers. If someone raises it, it cuts both ways: it is the paper's own reason for resolving conflicts by agreement.
  - Malik: the paper hedges, "a feature our current findings suggest may be problematic for users" (p. 1085). Keep "may."
  - Furnas: the paper cites ref. 8 only to justify aliasing (p. 1088, p. 1091). The phrase "vocabulary problem" is the title of ref. 8, not a claim in the paper's text. Linking it to the problem is our reading. The 20 percent figure comes from the Furnas abstract, not from this paper.
  - A sharp classmate may notice that ref. 19, Nielsen et al., is cited both for recognition-driven gesture choice (p. 1083) and as "a similar approach" to elicitation (p. 1084). The paper does not say how its method differs from Nielsen's. Answer honestly: we cannot tell from this paper.
  - Wu et al. got usability studies. The critique is that the gestures were not elicited, not that nobody tested them.
  - "Products were fixing conventions" is our framing, supported in section 01 (timeline) and section 03, not a claim of the paper. The paper's own urgency is public spaces, "where the importance of immediate usability is high" (p. 1085), and "It is therefore timely" (p. 1083).
  - Why it still matters (if asked): every new input channel without an established vocabulary, such as mid-air hand tracking or voice, faces the same choice. Section 11 picks this up. **[classmate]** slot here if a commentary raised the recognizer-first point.
- **Sources:** p. 1083 (designer quote, "despite skillful design," design space, "timely"); p. 1084 ("first to employ users," Good, Nielsen, Epps, Beringer, Voida, Long et al.); p. 1085 (Wu "not elicited," Malik, public spaces); p. 1088 and p. 1091 (aliasing, refs. 8 and 33); p. 1092 (reference years). Furnas et al. 1987 abstract: https://scrapbox.io/UIPedia/George_Furnas:_The_vocabulary_problem_in_human-system_communication (cross-checked by search summary; ACM page https://dl.acm.org/doi/10.1145/32206.32212 returned 403).

## Verification log

| Claim | Where checked | Status |
|---|---|---|
| "defined by system designers" (full: "most surface gestures have been defined by system designers, who personally employ them or teach them to user-testers") | p. 1083, PDF p. 1 | OK |
| "somewhat arbitrary gesture sets whose members may be chosen out of concern for reliable recognition [19]" | p. 1083 | OK |
| "Despite skillful design" (quoted as "skillful design") | p. 1083 | OK |
| "almost anything one can do with one's hands could be a potential gesture" | p. 1083 (sentence runs across columns) | OK |
| "It is therefore timely" | p. 1083 | OK |
| "ours is the first to employ users, rather than principles, in the development of a gesture set"; principled approach cites [20,35] | p. 1083 to 1084 (sentence starts at foot of p. 1083, quote words on p. 1084) | OK |
| Good et al. used referent-to-sign approach for a command-line email interface; ref. 9 dated 1984 | p. 1084; p. 1092 | OK |
| Nielsen et al. "describe a similar approach"; ref. 19 dated 2004 | p. 1084; p. 1092 | OK |
| Epps et al., limited similar study, Windows prompts "may have biased participants to simply emulate the mouse"; ref. 5 dated 2006 | p. 1084; p. 1092 | OK |
| Beringer (ref. 2, 2002) and Voida et al. (ref. 28, 2005) elicited gestures | p. 1084; p. 1092 | OK |
| Long et al.: "users are sometimes poor at picking easily differentiable gestures" | p. 1084 | OK |
| Wu's systems: "were not elicited from users, although usability studies were conducted" | p. 1085 (sentence starts on p. 1084) | OK |
| Malik et al. distinguished "1-, 2-, 3-, and 5-finger gestures, a feature our current findings suggest may be problematic for users" | p. 1085 | OK |
| Public spaces, "where the importance of immediate usability is high" | p. 1085 | OK |
| Furnas et al. ref. 8 cited for aliasing only ("Aliasing has been shown to dramatically increase input guessability [8,33]") | p. 1088; also p. 1091 | OK |
| Furnas et al. 1987, CACM 30(11), 964 to 971, "The vocabulary problem in human-system communication" | p. 1092 | OK |
| Furnas: "In every case two people favored the same term with probability <0.20" | https://scrapbox.io/UIPedia/George_Furnas:_The_vocabulary_problem_in_human-system_communication and web search summary | OK (secondary copy of abstract; ACM page blocked) |

17 claims checked.

## Open items
- No conflict with the pitfalls list. Pitfalls 8 (Long et al.) and 9 (Epps et al.) are respected in the notes.
- The Furnas 0.20 figure is verified only through a secondary copy of the abstract, since the ACM Digital Library returned 403. If the orchestrator has SFU library access, a one-line check of the original abstract would make it primary.
- Ref. 19 (Nielsen et al.) is cited both for recognition-driven gesture choice (p. 1083) and as "a similar approach" (p. 1084). Worth flagging to section 04 (method lineage) so the two sections answer the same way.
- Cross-references: the transition hands to section 03; "products fixing conventions" relies on section 01's timeline; "still matters" points to section 11.
- On-slide text is about 85 words including quotes. The Wu and Nielsen details are shortened on the slide and complete in the notes.
