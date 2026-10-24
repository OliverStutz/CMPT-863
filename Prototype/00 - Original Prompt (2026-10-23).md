# Prototype: original prompt from Oliver (2026-10-23)

Saved verbatim as the source of truth for the application design. Source documents: `CMPT 863 Project Abstract v8.md` and `Project Literature Review/Literature Review - Final v13.md` in the repo root.

---

With the Project Abstract v8, and the Literature Review Final v13, we have to make a new folder for the Prototype,

The goal is to create a "browser" simulation to perform the study, i would like to "automate" the entire study with that.


We will have:

1. Welcome Page for participants of the study
- Containing an Inline PDF Document and a Button at the end with "I accept the terms and would like to participate in the study", the button can only be used once the user "marked" the box properly
- After we ask all relevant questions which we need to gather from the participant (Selection of field of expertise, Country, Age-bracket, Job Level: Student, Employee, Leadership, C-Level) , passively we want to record IP Address and Browser Type so we know that nobody is botting us 
2. Setup time, we ensure the listener has audio activated, we play a "presentation example to test the video playback of the participant.
3. After the real configureable presentation starts, we will show like described in Abstract v8 the "Presentation Slide" as well as the audio, we define a time within the window to mark specific words/images on the presentation in a configureable color, with additional information e.g. Firewalls likely too expensive, cheaper ones available. The CISO slides should show the typical boardroom approval, i want you to generate 5 slide examples (which should be configureable!) and configuration to "mark" specific text segments/pictures and configureable text for it so we can play around and evolve it once it is created. Make absolutely sure the study design is followed as given in Abstract 8 !
- 1. Opening Boardmeeting, CISO Introduction, Topic Introduction  (no markings!)
- 2. Slide 1-4 with the Questions ahead of the start, a slide will always be "Rate, Competence, Benevolence and Integrity" of the CISO
- 3. Each Slide should have the presentation embedded, preferably also showing the presentation, the CISO and the Boardroom and the "checker" is like we are wearing "AR" glasses, just marking the text dynamically
- 4. a Slide should always be "The Slide and possible animation during it, the audio recording (which we need to create by doing text to audio)  the user can't replay the audio
- 5. After the audio is finished the person has 10 more seconds to look at the slide and it swaps automatically to showing the selection of the rating, which is only available for 15 seconds
- 6. Each time the presentation starts, audio is finished with playing and the user is selecting the points or the time expired for any, we send a message to the backend indicating the progress, so we have it in realtime and nobody can just post a fake message with all of it at our backend.
4. Last slide, will tell the participant that there have been no issues on the slides themselves


Your goal is the following;

- Saveguard the prompt above for the actual application design
- Create the Architecture for our "study", i would propose we use a Typescript driven React APP, with a go backend which runs on cloudrun (GCP) and push our artefacts to google CDN, deployment prefered with terraform
- We need a part for the study runs which provide a specific identifier which passes the information on how to load the app
- We need a backend e.g. admin menu, which can create specific identifiers + enter a string for the source e.g. Linkedin, Whatsapp etc.
- We need in the admin menu a tab which shows the list of participants , when they took it, if they stepped through entirely e.g. 1/X ... Started, Ongoing, Finished, Timer how long it took so far from start to finish (or running when still going)
- Save everything from the Typescript app via. the go backend
- Terraform for the full setup should be good


Dont assume, when you are creating the Architecture and you have an open point whichdoesnt match ask me. Lets get started, feel free to use OPUS agents to support you and split up the work
