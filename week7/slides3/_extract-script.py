#!/usr/bin/env python3
"""Extract the speaker script alone (per slide: title, time, Say with its cues, Transition) from the slides3 section files
and the title slide in the build script, into slides3/Week 7 - Speaker Script v3.md. Run from week7/: python3 slides3/_extract-script.py
"""
import re, glob, pathlib
here = pathlib.Path(__file__).parent
build = (here / "_build-deck-source.py").read_text(encoding="utf-8")
blocks = [re.search(r'TITLE = """(.*?)"""', build, re.S).group(1)]
for f in sorted(glob.glob(str(here / "0*-*.md"))):
    if re.search(r"/0[1-6]-", f):
        blocks.append(pathlib.Path(f).read_text(encoding="utf-8").split("\n## Verification log\n")[0])
slides = []
for b in blocks:
    for m in re.finditer(r"## Slide (\d+): (.+?)\n\n\*\*Time:\*\* ([\d:]+).*?- \*\*Say:\*\* (.*?)\n- \*\*Transition:\*\* (.*?)\n", b, re.S):
        num, title, tm, say, trans = m.groups()
        say = re.sub(r"\s*\(about \d+ words\)\s*$", "", say.strip())
        slides.append((int(num), title.strip(), tm, say, trans.strip()))
slides.sort()
total = sum(len(s[3].split()) for s in slides if s[0] <= 20)
out = ["# Week 7 speaker script, version 3\n",
       f"Gaver, \"What Should We Expect From Research Through Design?\", CHI 2012. {len(slides)} slides. For each slide: the title, its time cap, the words to say, and the line that leads into the next slide. [click] and [pause] are cues, not words. Slides 1 to 20 are the talk, about {total:,} words; 21 and 22 open the discussion. The full notes (what to watch out for, likely questions, sources) are in `../Week 7 - Deck Source v3.md`.\n", "---\n"]
for num, title, tm, say, trans in slides:
    out.append(f"## Slide {num}: {title}\n\n*Time cap {tm}, about {len(say.split())} words*\n")
    out.append(say + "\n")
    out.append(f"**Then:** {trans}\n")
    out.append("---\n")
dest = here / "Week 7 - Speaker Script v3.md"
dest.write_text("\n".join(out), encoding="utf-8")
print("wrote", dest, "with", len(slides), "slides,", total, "words in slides 1 to 20")
