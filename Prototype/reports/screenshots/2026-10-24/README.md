# Screenshot walk-through, 2026-10-24

Live service `https://study-api-pju4jz7m4a-nn.a.run.app`, run `tvfe8g` (Fast walk-through, config `test-fast-1.4.0`: same screens as the real study, short timers). Headless Firefox, 1366 x 768, audio muted. Taken by `e2e/tests/screenshots.spec.ts`, which drives each session by the step the server returns and holds the request that leaves a screen until that screen is captured. Slide steps have two shots: the slide as the clip starts and `-marked` when the clip ends (checker markings drawn; the neutral shot is the same moment).

Sessions: checker `ufjj3378nqu7` (actions challenge, explain, nothing; denial slide s4), neutral `baprdwnn4e8c` (Yes, No, Yes).

**Re-run, same day (local):** after the fixes listed at the end (slide 3 diagram, checker marks kept after the slide, action card in the boardroom, wrapping Flags column, device gate shot), every `checker-*` and `neutral-*` frame and `admin-04-participants.png` were replaced by the same frames from a local run (`q8cecx` on a local API with the Firestore emulator, config `test-fast-1.4.0`; checker `qh7xsaycud4w`, neutral `afgbyeckgayx`, same actions). `checker-21-postView-s3-marked.png` is new. The other admin frames and `contact-sheet.png` are still from the live run; session ids and times in the texts below refer to the live run.

Contact sheet of the checker session (tall pages cropped to the window): `contact-sheet.png`

![[reports/screenshots/2026-10-24/contact-sheet.png]]

## Checker session

1. `checker-01-deviceGate-deviceGate.png`: Device gate step, after the scene pack has loaded: the gate screen is an empty card (its heading is for screen readers only) while it measures width, orientation and touch and sends device_checked.

    ![[reports/screenshots/2026-10-24/checker-01-deviceGate-deviceGate.png]]

2. `checker-02-consent-consent.png`: Consent: PDF placeholder box, the checkbox (ticked here) and the accept button, now enabled. Look at the placeholder texts still waiting for the ethics form.

    ![[reports/screenshots/2026-10-24/checker-02-consent-consent.png]]

3. `checker-03-audioCheck-setup.audioCheck.png`: Audio check: Play button for the test clip and the four word choices (one answer only). Taken before Play is pressed.

    ![[reports/screenshots/2026-10-24/checker-03-audioCheck-setup.audioCheck.png]]

4. `checker-04-example-setup.exampleScene.png`: Example scene: the boardroom with an agenda slide and the three speakers (Chair, Board member, CISO), so the participant learns the layout before the meeting.

    ![[reports/screenshots/2026-10-24/checker-04-example-setup.exampleScene.png]]

5. `checker-05-preStart-setup.preStart.png`: Pre-start, checker version: look for the extra sentence "During this meeting, a checker reads each slide and marks anything it finds."

    ![[reports/screenshots/2026-10-24/checker-05-preStart-setup.preStart.png]]

6. `checker-06-opening-opening.png`: Opening: the presenter profile card for Daniel Hale (CISO, Meridian Logistics) on the screen while the chair introduces him.

    ![[reports/screenshots/2026-10-24/checker-06-opening-opening.png]]

7. `checker-07-postView-opening.png`: Post-view: the profile card stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/checker-07-postView-opening.png]]

8. `checker-08-rating-rating.T0.png`: Rating T0, after the opening, before any slide (baseline): trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-08-rating-rating.T0.png]]

9. `checker-09-slide-s1.png`: Slide 1 "Where we stand today" as it appears when the clip starts (checker panel shows only its title so far).

    ![[reports/screenshots/2026-10-24/checker-09-slide-s1.png]]

10. `checker-09-slide-s1-marked.png`: Slide 1 at the end of the clip: the panel shows "No issues found" with a tick.

    ![[reports/screenshots/2026-10-24/checker-09-slide-s1-marked.png]]

11. `checker-10-postView-s1.png`: Post-view: slide 1 stays on screen with the "Time left" bar until the server opens the next step. The checker panel (and on a flagged slide the markings) stays on screen during this view.

    ![[reports/screenshots/2026-10-24/checker-10-postView-s1.png]]

12. `checker-11-rating-rating.T1.png`: Rating T1, after slide 1: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-11-rating-rating.T1.png]]

13. `checker-12-slide-s2.png`: Slide 2 "The risk and what it would cost us" as it appears when the clip starts (checker panel shows only its title so far).

    ![[reports/screenshots/2026-10-24/checker-12-slide-s2.png]]

14. `checker-12-slide-s2-marked.png`: Slide 2 with the marking drawn: panel flag "$8 million does not follow from shown figures" (the three figure rows boxed).

    ![[reports/screenshots/2026-10-24/checker-12-slide-s2-marked.png]]

15. `checker-13-postView-s2.png`: Post-view: slide 2 stays on screen with the "Time left" bar until the server opens the next step. The checker panel (and on a flagged slide the markings) stays on screen during this view.

    ![[reports/screenshots/2026-10-24/checker-13-postView-s2.png]]

16. `checker-14-rating-rating.T2a.png`: Rating T2a, after slide 2, before the question round: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-14-rating-rating.T2a.png]]

17. `checker-15-action-s2.action.png`: Action for slide 2: "Say nothing", "Ask Daniel to explain", "Challenge Daniel with the checker's finding", with the time bar. The choice card sits at the bottom of the boardroom; the slide, its markings and the checker panel stay visible above it. This walk chose challenge.

    ![[reports/screenshots/2026-10-24/checker-15-action-s2.action.png]]

18. `checker-16-question-s2.question.png`: Voiced question on slide 2: the slide stays up while the question is played (audio only, nothing on screen names it). The checker panel and markings stay on screen.

    ![[reports/screenshots/2026-10-24/checker-16-question-s2.question.png]]

19. `checker-17-answer-s2.answer.png`: Answer on slide 2: Daniel answers over the same slide; the screen is the same for the reason and the denial variant.

    ![[reports/screenshots/2026-10-24/checker-17-answer-s2.answer.png]]

20. `checker-18-postView-s2.answer.png`: Post-view: slide 2 after the answer stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/checker-18-postView-s2.answer.png]]

21. `checker-19-rating-rating.T2b.png`: Rating T2b, after the answer on slide 2: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-19-rating-rating.T2b.png]]

22. `checker-20-slide-s3.png`: Slide 3 "How we will protect our 13 sites" as it appears when the clip starts (checker panel shows only its title so far).

    ![[reports/screenshots/2026-10-24/checker-20-slide-s3.png]]

23. `checker-20-slide-s3-marked.png`: Slide 3 with the marking drawn: panel flag "Head office is a single point of failure" (the network diagram and two bullets boxed). The diagram is `content/images/s3-network-diagram.svg`, extracted from the v1.3 slide.

    ![[reports/screenshots/2026-10-24/checker-20-slide-s3-marked.png]]

24. `checker-21-postView-s3.png`: Post-view: slide 3 stays on screen with the "Time left" bar until the server opens the next step. The checker panel (and on a flagged slide the markings) stays on screen during this view.

    ![[reports/screenshots/2026-10-24/checker-21-postView-s3.png]]

    `checker-21-postView-s3-marked.png` (new): the same view once the marks have finished drawing: the boxes on the diagram and the two bullets and the panel flag are still there after the clip ended.

    ![[reports/screenshots/2026-10-24/checker-21-postView-s3-marked.png]]

25. `checker-22-rating-rating.T3a.png`: Rating T3a, after slide 3, before the question round: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-22-rating-rating.T3a.png]]

26. `checker-23-action-s3.action.png`: Action for slide 3: "Say nothing", "Ask Daniel to explain", "Challenge Daniel with the checker's finding", with the time bar. The choice card sits at the bottom of the boardroom; the slide, its markings and the checker panel stay visible above it. This walk chose explain.

    ![[reports/screenshots/2026-10-24/checker-23-action-s3.action.png]]

27. `checker-24-question-s3.question.png`: Voiced question on slide 3: the slide stays up while the question is played (audio only, nothing on screen names it). The checker panel and markings stay on screen.

    ![[reports/screenshots/2026-10-24/checker-24-question-s3.question.png]]

28. `checker-25-answer-s3.answer.png`: Answer on slide 3: Daniel answers over the same slide; the screen is the same for the reason and the denial variant.

    ![[reports/screenshots/2026-10-24/checker-25-answer-s3.answer.png]]

29. `checker-26-postView-s3.answer.png`: Post-view: slide 3 after the answer stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/checker-26-postView-s3.answer.png]]

30. `checker-27-rating-rating.T3b.png`: Rating T3b, after the answer on slide 3: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-27-rating-rating.T3b.png]]

31. `checker-28-slide-s4.png`: Slide 4 "What the $2.4 million pays for" as it appears when the clip starts (checker panel shows only its title so far).

    ![[reports/screenshots/2026-10-24/checker-28-slide-s4.png]]

32. `checker-28-slide-s4-marked.png`: Slide 4 with the marking drawn: panel flag "Firewall cost per site looks high" (the firewall row boxed).

    ![[reports/screenshots/2026-10-24/checker-28-slide-s4-marked.png]]

33. `checker-29-postView-s4.png`: Post-view: slide 4 stays on screen with the "Time left" bar until the server opens the next step. The checker panel (and on a flagged slide the markings) stays on screen during this view.

    ![[reports/screenshots/2026-10-24/checker-29-postView-s4.png]]

34. `checker-30-rating-rating.T4a.png`: Rating T4a, after slide 4, before the question round: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-30-rating-rating.T4a.png]]

35. `checker-31-action-s4.action.png`: Action for slide 4: "Say nothing", "Ask Daniel to explain", "Challenge Daniel with the checker's finding", with the time bar. The choice card sits at the bottom of the boardroom; the slide, its markings and the checker panel stay visible above it. This walk chose nothing.

    ![[reports/screenshots/2026-10-24/checker-31-action-s4.action.png]]

36. `checker-32-question-s4.question.png`: Voiced question on slide 4: the slide stays up while the question is played (audio only, nothing on screen names it). The checker panel and markings stay on screen.

    ![[reports/screenshots/2026-10-24/checker-32-question-s4.question.png]]

37. `checker-33-answer-s4.answer.png`: Answer on slide 4: Daniel answers over the same slide; the screen is the same for the reason and the denial variant.

    ![[reports/screenshots/2026-10-24/checker-33-answer-s4.answer.png]]

38. `checker-34-postView-s4.answer.png`: Post-view: slide 4 after the answer stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/checker-34-postView-s4.answer.png]]

39. `checker-35-rating-rating.T4b.png`: Rating T4b, after the answer on slide 4: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-35-rating-rating.T4b.png]]

40. `checker-36-slide-s5.png`: Slide 5 "The decision today" as it appears when the clip starts (checker panel shows only its title so far).

    ![[reports/screenshots/2026-10-24/checker-36-slide-s5.png]]

41. `checker-36-slide-s5-marked.png`: Slide 5 at the end of the clip: the panel shows "No issues found" with a tick.

    ![[reports/screenshots/2026-10-24/checker-36-slide-s5-marked.png]]

42. `checker-37-postView-s5.png`: Post-view: slide 5 stays on screen with the "Time left" bar until the server opens the next step. The checker panel (and on a flagged slide the markings) stays on screen during this view.

    ![[reports/screenshots/2026-10-24/checker-37-postView-s5.png]]

43. `checker-38-rating-rating.T5.png`: Rating T5, after slide 5: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/checker-38-rating-rating.T5.png]]

44. `checker-39-budget-budget.png`: Budget vote: approve or reject the $2.4 million request, then "How sure are you?" 0 to 10, with the time bar.

    ![[reports/screenshots/2026-10-24/checker-39-budget-budget.png]]

45. `checker-40-background-background.png`: Background questions (full page, taller than the window): field, country, age, job level, education, cybersecurity knowledge 1 to 5 and AI tool use, each with "Prefer not to say".

    ![[reports/screenshots/2026-10-24/checker-40-background-background.png]]

46. `checker-41-debrief-debrief.png`: Debrief, checker version: "You were in the group with the checker", the fixed and wrong flags, placeholders for contact and ethics text, the do-not-discuss line and the session code.

    ![[reports/screenshots/2026-10-24/checker-41-debrief-debrief.png]]

47. `checker-42-done-done.png`: Done step: the client keeps the debrief on screen after sending finished (identical to the debrief shot).

    ![[reports/screenshots/2026-10-24/checker-42-done-done.png]]

## Neutral session

1. `neutral-01-deviceGate-deviceGate.png`: Device gate step, after the scene pack has loaded: the gate screen is an empty card (its heading is for screen readers only) while it measures width, orientation and touch and sends device_checked.

    ![[reports/screenshots/2026-10-24/neutral-01-deviceGate-deviceGate.png]]

2. `neutral-02-consent-consent.png`: Consent: PDF placeholder box, the checkbox (ticked here) and the accept button, now enabled. Look at the placeholder texts still waiting for the ethics form.

    ![[reports/screenshots/2026-10-24/neutral-02-consent-consent.png]]

3. `neutral-03-audioCheck-setup.audioCheck.png`: Audio check: Play button for the test clip and the four word choices (one answer only). Taken before Play is pressed.

    ![[reports/screenshots/2026-10-24/neutral-03-audioCheck-setup.audioCheck.png]]

4. `neutral-04-example-setup.exampleScene.png`: Example scene: the boardroom with an agenda slide and the three speakers (Chair, Board member, CISO), so the participant learns the layout before the meeting.

    ![[reports/screenshots/2026-10-24/neutral-04-example-setup.exampleScene.png]]

5. `neutral-05-preStart-setup.preStart.png`: Pre-start, neutral version: only the "hear each speaker once" line; compare with the checker screen, which has one more sentence.

    ![[reports/screenshots/2026-10-24/neutral-05-preStart-setup.preStart.png]]

6. `neutral-06-opening-opening.png`: Opening: the presenter profile card for Daniel Hale (CISO, Meridian Logistics) on the screen while the chair introduces him.

    ![[reports/screenshots/2026-10-24/neutral-06-opening-opening.png]]

7. `neutral-07-postView-opening.png`: Post-view: the profile card stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-07-postView-opening.png]]

8. `neutral-08-rating-rating.T0.png`: Rating T0, after the opening, before any slide (baseline): trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-08-rating-rating.T0.png]]

9. `neutral-09-slide-s1.png`: Slide 1 "Where we stand today" as it appears when the clip starts (no checker panel in this group).

    ![[reports/screenshots/2026-10-24/neutral-09-slide-s1.png]]

10. `neutral-09-slide-s1-marked.png`: Slide 1 at the moment the markings appear for the checker group: nothing changes here, no panel and no marks.

    ![[reports/screenshots/2026-10-24/neutral-09-slide-s1-marked.png]]

11. `neutral-10-postView-s1.png`: Post-view: slide 1 stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-10-postView-s1.png]]

12. `neutral-11-rating-rating.T1.png`: Rating T1, after slide 1: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-11-rating-rating.T1.png]]

13. `neutral-12-slide-s2.png`: Slide 2 "The risk and what it would cost us" as it appears when the clip starts (no checker panel in this group).

    ![[reports/screenshots/2026-10-24/neutral-12-slide-s2.png]]

14. `neutral-12-slide-s2-marked.png`: Slide 2 at the moment the markings appear for the checker group: nothing changes here, no panel and no marks.

    ![[reports/screenshots/2026-10-24/neutral-12-slide-s2-marked.png]]

15. `neutral-13-postView-s2.png`: Post-view: slide 2 stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-13-postView-s2.png]]

16. `neutral-14-rating-rating.T2a.png`: Rating T2a, after slide 2, before the question round: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-14-rating-rating.T2a.png]]

17. `neutral-15-action-s2.action.png`: Action for slide 2: "Do you have a question about this slide?" Yes / No. The choice card sits at the bottom of the boardroom with the slide visible above it. This walk chose Yes.

    ![[reports/screenshots/2026-10-24/neutral-15-action-s2.action.png]]

18. `neutral-16-question-s2.question.png`: Voiced question on slide 2: the slide stays up while the question is played (audio only, nothing on screen names it).

    ![[reports/screenshots/2026-10-24/neutral-16-question-s2.question.png]]

19. `neutral-17-answer-s2.answer.png`: Answer on slide 2: Daniel answers over the same slide; the screen is the same for the reason and the denial variant.

    ![[reports/screenshots/2026-10-24/neutral-17-answer-s2.answer.png]]

20. `neutral-18-postView-s2.answer.png`: Post-view: slide 2 after the answer stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-18-postView-s2.answer.png]]

21. `neutral-19-rating-rating.T2b.png`: Rating T2b, after the answer on slide 2: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-19-rating-rating.T2b.png]]

22. `neutral-20-slide-s3.png`: Slide 3 "How we will protect our 13 sites" as it appears when the clip starts (no checker panel in this group).

    ![[reports/screenshots/2026-10-24/neutral-20-slide-s3.png]]

23. `neutral-20-slide-s3-marked.png`: Slide 3 at the moment the markings appear for the checker group: nothing changes here, no panel and no marks. The network diagram is shown.

    ![[reports/screenshots/2026-10-24/neutral-20-slide-s3-marked.png]]

24. `neutral-21-postView-s3.png`: Post-view: slide 3 stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-21-postView-s3.png]]

25. `neutral-22-rating-rating.T3a.png`: Rating T3a, after slide 3, before the question round: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-22-rating-rating.T3a.png]]

26. `neutral-23-action-s3.action.png`: Action for slide 3: "Do you have a question about this slide?" Yes / No. The choice card sits at the bottom of the boardroom with the slide visible above it. This walk chose No.

    ![[reports/screenshots/2026-10-24/neutral-23-action-s3.action.png]]

27. `neutral-24-question-s3.question.png`: Voiced question on slide 3: the slide stays up while the question is played (audio only, nothing on screen names it).

    ![[reports/screenshots/2026-10-24/neutral-24-question-s3.question.png]]

28. `neutral-25-answer-s3.answer.png`: Answer on slide 3: Daniel answers over the same slide; the screen is the same for the reason and the denial variant.

    ![[reports/screenshots/2026-10-24/neutral-25-answer-s3.answer.png]]

29. `neutral-26-postView-s3.answer.png`: Post-view: slide 3 after the answer stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-26-postView-s3.answer.png]]

30. `neutral-27-rating-rating.T3b.png`: Rating T3b, after the answer on slide 3: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-27-rating-rating.T3b.png]]

31. `neutral-28-slide-s4.png`: Slide 4 "What the $2.4 million pays for" as it appears when the clip starts (no checker panel in this group).

    ![[reports/screenshots/2026-10-24/neutral-28-slide-s4.png]]

32. `neutral-28-slide-s4-marked.png`: Slide 4 at the moment the markings appear for the checker group: nothing changes here, no panel and no marks.

    ![[reports/screenshots/2026-10-24/neutral-28-slide-s4-marked.png]]

33. `neutral-29-postView-s4.png`: Post-view: slide 4 stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-29-postView-s4.png]]

34. `neutral-30-rating-rating.T4a.png`: Rating T4a, after slide 4, before the question round: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-30-rating-rating.T4a.png]]

35. `neutral-31-action-s4.action.png`: Action for slide 4: "Do you have a question about this slide?" Yes / No. The choice card sits at the bottom of the boardroom with the slide visible above it. This walk chose Yes.

    ![[reports/screenshots/2026-10-24/neutral-31-action-s4.action.png]]

36. `neutral-32-question-s4.question.png`: Voiced question on slide 4: the slide stays up while the question is played (audio only, nothing on screen names it).

    ![[reports/screenshots/2026-10-24/neutral-32-question-s4.question.png]]

37. `neutral-33-answer-s4.answer.png`: Answer on slide 4: Daniel answers over the same slide; the screen is the same for the reason and the denial variant.

    ![[reports/screenshots/2026-10-24/neutral-33-answer-s4.answer.png]]

38. `neutral-34-postView-s4.answer.png`: Post-view: slide 4 after the answer stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-34-postView-s4.answer.png]]

39. `neutral-35-rating-rating.T4b.png`: Rating T4b, after the answer on slide 4: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-35-rating-rating.T4b.png]]

40. `neutral-36-slide-s5.png`: Slide 5 "The decision today" as it appears when the clip starts (no checker panel in this group).

    ![[reports/screenshots/2026-10-24/neutral-36-slide-s5.png]]

41. `neutral-36-slide-s5-marked.png`: Slide 5 at the moment the markings appear for the checker group: nothing changes here, no panel and no marks.

    ![[reports/screenshots/2026-10-24/neutral-36-slide-s5-marked.png]]

42. `neutral-37-postView-s5.png`: Post-view: slide 5 stays on screen with the "Time left" bar until the server opens the next step.

    ![[reports/screenshots/2026-10-24/neutral-37-postView-s5.png]]

43. `neutral-38-rating-rating.T5.png`: Rating T5, after slide 5: trust 0 to 10 plus the three 1 to 7 scales (competence, benevolence, integrity), with the time bar. Taken before any answer is clicked.

    ![[reports/screenshots/2026-10-24/neutral-38-rating-rating.T5.png]]

44. `neutral-39-budget-budget.png`: Budget vote: approve or reject the $2.4 million request, then "How sure are you?" 0 to 10, with the time bar.

    ![[reports/screenshots/2026-10-24/neutral-39-budget-budget.png]]

45. `neutral-40-background-background.png`: Background questions (full page, taller than the window): field, country, age, job level, education, cybersecurity knowledge 1 to 5 and AI tool use, each with "Prefer not to say".

    ![[reports/screenshots/2026-10-24/neutral-40-background-background.png]]

46. `neutral-41-debrief-debrief.png`: Debrief, neutral version: "You were in the group without the checker" and which slides the other group saw flagged; same placeholders and session code.

    ![[reports/screenshots/2026-10-24/neutral-41-debrief-debrief.png]]

47. `neutral-42-done-done.png`: Done step: the client keeps the debrief on screen after sending finished (identical to the debrief shot).

    ![[reports/screenshots/2026-10-24/neutral-42-done-done.png]]

## Admin

1. `admin-01-404-without-key.png`: `/admin/` without the gate key: the plain "404 page not found", so the admin area is invisible without the key.

    ![[reports/screenshots/2026-10-24/admin-01-404-without-key.png]]

2. `admin-02-login.png`: After opening `/admin/?k=<key>`: the login dialog over the Runs tab, password field still empty.

    ![[reports/screenshots/2026-10-24/admin-02-login.png]]

3. `admin-03-runs.png`: Runs tab: new-run form and the run list with link, Copy and QR, config version, status, the O/S/On/F/A/B counts and finished per group. Look at the tvfe8g row.

    ![[reports/screenshots/2026-10-24/admin-03-runs.png]]

4. `admin-04-participants.png`: Participants tab (local re-run `q8cecx`): the two walk sessions Finished, progress 32/32, last step done:done. The Flags column now fits at 1366 px: flag chips wrap inside their cell, and the table still scrolls sideways inside its card on narrower windows. (The live shot had the column cut at the right edge with rate_limited and possible_duplicate.)

    ![[reports/screenshots/2026-10-24/admin-04-participants.png]]

5. `admin-05-session-detail.png`: Session detail drawer for the checker session: status, run, config, group, denial slide s4, the timestamps, browser, flags, and the start of the event timeline (51 events, 5 rejected, highlighted in red).

    ![[reports/screenshots/2026-10-24/admin-05-session-detail.png]]

6. `admin-06-session-detail-timeline.png`: Same drawer scrolled down: events 31 to 51 with the state transitions (slide, postView, rating, action, question, answer, budget, background, debrief, done) and the start of the Ratings JSON.

    ![[reports/screenshots/2026-10-24/admin-06-session-detail-timeline.png]]

7. `admin-07-export.png`: Export tab: CSV (one row per rating point) and JSON per run or for all runs, with the "Include invalid events" and "Include IP address" options.

    ![[reports/screenshots/2026-10-24/admin-07-export.png]]

8. `admin-08-configs.png`: Configs tab: versions 1.3.0, 1.4.0 (validated) and test-fast-1.4.0 (locked), five slides each, flagged slides 2, 3, 4 and a lint report button.

    ![[reports/screenshots/2026-10-24/admin-08-configs.png]]

## What did not render (live run) and what was fixed

- Device gate: both live shots showed "Loading...". The spec now lets the landing's GET /next through on the device gate step (it was held by the screenshot gate, so "Loading..." could never clear) and waits for "Loading..." to disappear; the re-run frames show the gate screen.
- Slide 3 image: `s3-network-diagram.svg` did not exist. It is now `content/images/s3-network-diagram.svg` (served from `/assets/images/`), the diagram group of `content/slides/v1.3/s3.svg` with its own viewBox.
- Checker group: the panel and markings were shown only while the slide clip played (the server sent `overlay` on slide steps only). Checker sessions now get the slide's overlay on the post-view, action, question, answer and post-answer view steps of that slide, so marks and panel stay until the next slide (rating stays a separate screen without them). The action choice is now made in the boardroom with the slide in view.
- Admin Flags column: chips wrap inside the cell.
