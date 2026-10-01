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

**Speaker notes**
- **Say:** The table in this study was coarse by today's standards. Twenty-four by eighteen inches at ten twenty-four by seven sixty-eight is about forty-three pixels per inch. A current iPad Pro has two hundred sixty-four, roughly six times finer in each direction. The paper sensed touch with a vision system watching the hands from beneath the table and logged each contact as an oval with a timestamp. The four video cameras were a separate recording of the session. There was no recognizer and, as one participant showed by adding fingers "to ensure that I was pressing," no feedback that a touch had landed. The paper says participants were instructed to touch the table. Even so, some beckoned palm-up, lifted objects with two hands, or dragged things off the screen edge. The authors read this as a sensing gap: they suggest looking down or sideways at hands and adding a low-resolution sensing border. They also argue the many physical gestures support a physics engine, but not alone. We argue the "ease of recognition" claim is a hypothesis, since no recognizer was built.
- **Transition:** A table that could not see the air, and a set that was never run through a recognizer, lead straight into the study's limitations.
- **Watch out:**
  - The touch instruction is not in the Procedure section (pp. 1085 to 1086). It is one sentence under "Acting above the Table" on p. 1090: "We instructed participants to touch the table while gesturing." Quote it exactly; the paper does not say how or when the instruction was given.
  - The paper never names resolution as a problem. Linking 43 ppi to the advice that "demands for precise positioning should be avoided" (p. 1091) is our inference, and the paper's reason for that advice is user behaviour (12 of 14 resized sideways), not pixels.
  - Do not say the Surface used infrared cameras or projectors underneath: the paper says only "vision system" "from beneath the table" (p. 1086) and "the table's projected screen" (p. 1090). Extra hardware detail would need an outside source.
  - 43.9% physical matches the Figure 2 nature value .439 (pitfall 20).
  - If asked whether the iPad comparison is fair: it is a hand-held screen, not a table; the point is density only.
- **Sources:** pp. 1085, 1086, 1090, 1091, 1092 (ref. 32); https://www.apple.com/ipad-pro/specs/

## Verification log

| Claim | Where checked | Status |
|---|---|---|
| Surface prototype 24" x 18" at 1024 x 768 | p. 1085, Apparatus | OK |
| About 43 ppi | computed: 1024 / 24 = 768 / 18 = 1280 / 30 (diagonal) = 42.7 | OK |
| Current iPad Pro (11-inch and 13-inch) 264 ppi | https://www.apple.com/ipad-pro/specs/ (fetched 2026-09-30) | OK |
| About six times finer per direction | computed: 264 / 42.7 = 6.2 | OK |
| "The Surface vision system watched participants' hands from beneath the table and reported contact information to our software" | p. 1086 | OK |
| "All contacts were logged as ovals having millisecond timestamps" | p. 1086 | OK |
| Hands "videotaped from four angles"; two authors took notes on think-aloud | p. 1086 | OK |
| "The system did not attempt to recognize users' gestures, but did track and log all hand contact" | p. 1085 | OK |
| No feedback: participant used more fingers "to ensure that I was pressing"; "lack of feedback from the table when it was being touched" | p. 1090 | OK |
| "We instructed participants to touch the table while gesturing." | p. 1090, Acting above the Table (not in Procedure, pp. 1085 to 1086) | OK |
| Beckoning palm-up for help, "X" with top hand about 3" above, two-hand "lifted", "glue" in the air | p. 1090 | OK |
| Dragging from and to off-screen (clipboard, delete, reject; cut left, delete top) | p. 1090, A Land Beyond the Screen | OK |
| "the table's projected screen" | p. 1090 | OK |
| Recommendation: "look down or sideways at users' hands, rather than just up"; "a low-resolution sensing boundary beyond the high-resolution display area" | p. 1091 | OK |
| Ref. 32: Wilson, Izadi, Hilliges, Garcia-Mendoza, Kirk (2008) "Bringing physics to the surface," Proc. UIST '08, 67 to 76 | p. 1092 | OK |
| 43.9% physical gestures support a physics engine | p. 1091 | OK |
| "Seven participants, for example, expected intervening objects to move out of the way" | p. 1091 | OK |
| "Four participants "threw" an object off-screen to delete or reject it" | p. 1091 | OK |
| "a physics engine alone will probably not suffice as an adequate recognizer for all surface gestures" | p. 1091 | OK |
| "ease of recognition, consistency, reversibility, and versatility through aliasing" | p. 1091, Conclusion | OK |
| Recognizer is a next step ("implemented with a vision-based gesture recognizer") | p. 1091 | OK |
| "demands for precise positioning should be avoided"; 12 of 14 resized sideways | p. 1091 | OK |

## Open items
- Conflict with the presenter's belief resolved in the presenter's favour: the paper does state a touch instruction, but on p. 1090, not in the Procedure. The brief's fallback line ("the paper does not say whether touching was required") must not be used.
- Ref. 32 also appears in Related Work on p. 1085 ("used a physics engine with Microsoft Surface"); section 10 may cross-reference it.
- No [UNVERIFIED] items.
