# Week 5: Zooids — Building Blocks for Swarm User Interfaces

Le Goc, M., Kim, L. H., Parsaei, A., Fekete, J.-D., Dragicevic, P., and Follmer, S. (2016). Zooids: Building Blocks for Swarm User Interfaces. *Proceedings of UIST 2016*, Tokyo, pp. 97–106. DOI: 10.1145/2984511.2984547. Open hardware/software: https://github.com/PhysicalInteractionLab/SwarmUI/

Page numbers below are the printed ones (97–106). Companion reading this week: Ishii & Ullmer (1997), *Tangible Bits*.

---

## 1. Hardware: a specification built to shrink

Reading the hardware section from a specification standpoint, I can tell that nearly every design decision was made to strip requirements off the robot itself so it could be "small-sized" to 26 mm diameter, 21 mm height and about 12 g (p. 101). The two FA-GM6-3V-25 micro motors are deliberately placed non-colinearly to minimise diameter, with the authors noting the net force and moment still match a colinear layout (p. 101, Figure 5). Most telling is what is not on board: there is no camera or vision; position sensing is pushed off the robot onto a 3000 Hz DLP projector flashing gray-coded patterns that two cheap photodiodes decode locally at roughly 73 Hz, which also removes network latency from the PID loop (p. 101, Figure 6). Everything else is similarly minimal: a 48 MHz STM32F051C8, an nRF24L01+ radio constrained to 12-byte packets and 10 slaves per master (safety factor of about 2 from a measured maximum of 18), a single colour LED for identity and feedback, and a flexible capacitive electrode wrapped inside the enclosure rather than richer sensing (p. 101). The power budget follows the same logic, with 40 mA idle and 100 mA moving on a 100 mAh LiPo yielding one hour of motion (p. 101), and the whole part arriving at about $50 per robot, or $20 at volume (p. 98). I appreciate that the authors are explicit about the floor they hit: going smaller means leaving geared DC motors for piezo actuation, where power electronics remain the bottleneck, or moving to printable, foldable robots (p. 105).

## 2. Critique: at this size, what is it for?

The concept I liked most is zooids as physical pixels, yet at their current size I see no real application for it. Each robot is 26 mm wide, 21 mm tall and 12 g (p. 101), the tracked area is only 1 m × 0.63 m (p. 101), and the authors admit that at about $50 per unit, swarms beyond 30–40 robots are cost-prohibitive outside research (pp. 98, 105). A circle of 16 elements (Figure 11) is nothing I would prefer over a laptop. If zooids were far smaller, which the authors say requires replacing geared DC motors with piezo actuators (p. 105), I could imagine genuinely useful tangible charts. But what price would we pay compared to a screen? A screen gives millions of elements, instant refresh and no moving parts for free; zooids move at 44 cm/s (p. 101), need a 3000 Hz DLP projector 1.25 m above the table that fails under occlusion and ambient light (pp. 101, 105), last one hour on a 100 mAh battery docked by hand (pp. 101, 105), and still need paper or a tablet for text (p. 104). The paper concedes that free positioning "comes at the cost of slower response time and higher engineering complexity" (p. 103) and that GUIs "currently support far higher information resolution" (p. 100). Shrinking the robots multiplies these costs, and the missing user study against GUIs (p. 105) leaves the trade unquantified.

## 3. From surface gestures to a third dimension

My Wobbrock, Morris and Wilson (2009) presentation connects directly: a zooid senses only capacitive touch around its rim and being displaced, handled by re-targeting its goal to the new position (pp. 101, 104), and pushing many at once with the arm is a "sweeping" gesture (pp. 98, 100). In Wobbrock's terms the referents are rich but the signs are designer-defined; the authors concede their LED colour mappings are "arbitrary" and assume a familiar user (p. 104), so a novice elicitation study is an obvious extension. But I would rather build it in three dimensions. The authors admit rubber wheels confine Zooids to flat surfaces and 2D interaction, and name Kushleyev et al.'s micro-quadrotor swarm [33] as the route to "fully 3D interfaces" (p. 105). Their review shows the cost: BitDrones [62] can be grasped mid-air, but each drone is 8.9 cm and turbulence caps the prototype at three; Drone 100 [16] flies a hundred spaxels, but its volume rules out direct manipulation; and camera-tracked DisplaySwarm [3] gestures fall into a "gray area" because the elements no longer sense the user (p. 99). The IEEE drone work I linked (https://ieeexplore.ieee.org/document/6603259) points the same way. Keeping on-body touch sensing while leaving the table would preserve what makes it a swarm UI, unlike the 2.5D stacking the authors propose (p. 105).

## 4. Objects, particles, atoms: foundations for matter that forms itself

What I find most exciting is where the authors place their own work. Figure 9 draws a continuum from "things" to "stuff", with objects, particles and atoms along the bottom: classical TUIs sit on the left, made of individual solid things, while swarm UIs occupy the right half, where matter can be reshaped, divided and merged (p. 103). Zooids, as a "low-resolution swarm UI", stand in the grey area in between and "have both the affordance of things and stuff" (p. 103). The analogy with 80s displays makes the trajectory explicit: pixels were once discernible particles like zooids, and then shrank into invisible atoms (p. 103, Figure 10). Since the authors frame swarm UIs as a "coarse-grained version" of Sutherland's and Ishii's programmable matter (pp. 97–98), I read Zooids as foundational research for the dynamic formation of objects we can actually touch. One might imagine that future comes with wings, but the authors' own path is ground-based: smaller elements driven by piezo actuation, and swarms that, like ants, reconfigure into 2.5D or 3D by interweaving, connecting, rolling on top of each other, or using ramp robots and passive building blocks (p. 105). Once elements reach particle or atom size, that is matter that forms itself, without wings.

---

## Appendix A: Facts and page references used

### Hardware (paragraph 1)
- 26 mm diameter, 21 mm height, ~12 g; 100 mAh LiPo (p. 101)
- Motors placed non-colinearly to minimise diameter; same net force and moment as colinear (p. 101, Figure 5 exploded view: PCB, battery, motor holder, motors and wheels, touch sensor, enclosure, caster wheels)
- Two FA-GM6-3V-25 micro motors, Allegro A3901 driver; max ~74 cm/s, operated at ~44 cm/s (p. 101)
- Projector-based tracking: DLP LightCrafter at 3000 Hz, gray-coded patterns, two photodiodes decode position on-robot; ~73 Hz refresh; no network latency in local feedback control (p. 101, Figure 6)
- 48 MHz ARM STM32F051C8; nRF24L01+ 2.4 GHz radio; up to 18 slaves per master at 12-byte packets, Zooids use 10 for a safety factor of ~2 (p. 101)
- Single colour LED between photodiodes for identification and feedback (p. 101)
- Flexible electrode wrapped inside enclosure for capacitive touch, Atmel AT42QT1070 (p. 101)
- ~40 mA stationary, ~100 mA moving; one hour of movement on 100 mAh (p. 101)
- ~$50 USD each, down to ~$20 if mass manufactured (p. 98)
- Future work: smaller elements need piezo actuation instead of geared DC motors; power electronics at that scale remain challenging; printable/foldable robots as cheaper fabrication (p. 105)

### Critique (paragraph 2)
- Projection/tracking area 1 m × 0.63 m, projector 1.25 m above the table (p. 101)
- "$50 USD per robot in cost for parts and assembly … swarms larger than 30–40 cost prohibitive outside of research applications" (p. 105)
- 16-element circle (Figure 11, p. 103)
- Tracking adds cost and setup, suffers from occlusion, fails in strong ambient light (p. 105)
- Charging requires placing each robot manually in an individual charger (p. 105)
- Extra visual information (text, axes) needs paper templates, a display underneath, or top projection (p. 104, "Environment: Extra Visual Feedback"; Figure 8)
- "comes at the cost of slower response time and higher engineering complexity, with algorithmic problems such as collision avoidance and optimal element-target assignment" (p. 103, "Display: Fixed vs. Movable Elements")
- "although GUIs currently support far higher information resolution, Zooids enable richer physical gestures" (p. 100, Bezier curves)
- User studies comparing swarm UIs with GUIs are left as future work (p. 105)

### Gestures and drones (paragraph 3)
- A zooid is "a small wheel-propelled robot with position and touch sensing capabilities" (p. 97); users manipulate many at once using "sweeping" gestures [35] (p. 98); curves deformed "by pushing them with the side of the arm" (p. 100)
- System "registering when a zooid is touched and … constantly updating its goal based on its position" (p. 104)
- Input limited to capacitive touch around the circumference; when many are moved at once only those touching the hand register (p. 105)
- "the color mappings are arbitrary and this approach assumes a user who is already familiar with the system" (p. 104)
- Wheels limit the system to 2D; "work on swarms of aerial drones [33] present opportunities to make fully 3D interfaces"; authors prefer ground-based 2.5D/3D via interweaving, connecting, rolling on top of each other [8, 59], ramp robots, passive blocks [79] (p. 105)
- BitDrones [62]: mid-air 3D physical display, direct manipulation, 8.9 cm drones, turbulence limits prototype to 3 (p. 99)
- Drone 100 [16]: a hundred light-equipped "spaxels"; "the large operating volume prevents direct manipulation" (p. 99)
- Swarm UI definition requires elements to react to user input themselves; vision-based mid-air gesture systems "such as DisplaySwarm [3] are in a gray area" (p. 99)
- Wobbrock et al. (2009): referent = command, sign = gesture; elicitation from users vs designer-defined gestures chosen for recogniser reliability (see Presentation-Week4 notes)
- Zooids reference [33]: Kushleyev, Mellinger, Powers, Kumar. "Towards a swarm of agile micro quadrotors." *Autonomous Robots* 35(4), 2013, 287–300.

### Things/stuff continuum (paragraph 4)
- Sutherland's Ultimate Display ("a room within which the computer can control the existence of matter") and Ishii's Radical Atoms ("a new kind of matter capable of changing form dynamically") (p. 97)
- Swarm UIs "can be seen as a coarse-grained version of Sutherland's and Ishii's futuristic visions of user interfaces based on programmable matter" (p. 98)
- Figure 9 "The continuum between 'things' and 'stuff'": objects → particles → atoms; classical TUIs left, swarm UIs right, zooids in the grey area (p. 103)
- Things = "physical entities experienced as individual, solid objects"; stuff = "shapes and material that can be reshaped, divided, merged, or temporarily solidified to emulate things"; particles visible, atoms invisible (p. 103)
- "As a low-resolution swarm UI implementation, zooids stand in the gray area of the continuum and have both the affordance of things and stuff" (p. 103)
- 80s displays (ZX-81, TRS-80, 64 × 48 px): "pixels were discernible particles much like the zooids"; ultra-high resolution made them "atoms" (p. 103, Figure 10)
- "smaller and more elements may allow for radically different and richer styles of interaction with 'stuff' instead of 'things'" (p. 105)

## Appendix B: Notes on sources
- All quotes and numbers were checked against the PDF text of `Week 5 - Zooids.pdf`.
- The IEEE Xplore link (document 6603259) could not be fetched during writing (the site blocks non-browser access), so paragraph 3 cites it by URL only and does not describe its content. Replace "The IEEE drone work I linked" with the paper's title and authors once confirmed.
- Each paragraph was drafted independently from a separate slice of the author's notes and then assembled; paragraph 2 is the single critical paragraph.
