#!/usr/bin/env python3
"""Extract the slide text alone from the slides3 section files (and the title slide in the build script)
into slides3/slides-only/Week 7 - Slides v3.md: per slide the title, Layout, Images, Build, Text and Quote lines. Nothing else.
Run from week7/: python3 slides3/_extract-slides-only.py
"""
import re, glob, pathlib
here = pathlib.Path(__file__).parent
build = (here / "_build-deck-source.py").read_text(encoding="utf-8")
title_block = re.search(r'TITLE = """(.*?)"""', build, re.S).group(1)
files = sorted(f for f in glob.glob(str(here / "0*-*.md")) if re.search(r"/0[1-6]-", f))
blocks = [title_block] + [pathlib.Path(f).read_text(encoding="utf-8").split("\n## Verification log\n")[0] for f in files]
slides = []
for b in blocks:
    for m in re.finditer(r"## Slide (\d+): (.+?)\n\n\*\*Time:\*\* [\d:]+ · \*\*Visual:\*\* (.*?)\n\*\*Images:\*\* (.*?)\n\*\*Build:\*\* (.*?)\n\n\*\*On the slide\*\*\n(.*?)\n\*\*Speaker notes\*\*", b, re.S):
        num, title, visual, images, buildline, on = m.groups()
        text, quotes = [], []
        for line in on.strip().splitlines():
            s = line.rstrip()
            if re.match(r"\s*- Quote: ", s):
                quotes.append(re.sub(r"^\s*- Quote: ", "", s))
            else:
                text.append(s)
        slides.append((int(num), title.strip(), visual.strip(), images.strip(), buildline.strip(), text, quotes))
slides.sort()
n = len(slides)
out = ["# Week 7 slides, version 3: slide content only\n",
       f'Gaver, "What Should We Expect From Research Through Design?", CHI 2012. {n} slides. Each heading is the slide title. **Text** is the slide body, bullet by bullet. **Quote** is a pull quote with its attribution in small type. **Layout**, **Images** and **Build** give the arrangement, the image file and the click order.\n', "---\n"]
for num, title, visual, images, buildline, text, quotes in slides:
    out.append(f"## Slide {num} of {n}: {title}\n")
    out.append(f"**Layout:** {visual}\n**Images:** {images}\n**Build:** {buildline}\n")
    out.append("**Text:**\n" + "\n".join(text) + "\n")
    for q in quotes:
        out.append(f"**Quote:** {q}\n")
    out.append("---\n")
dest = here / "slides-only" / "Week 7 - Slides v3.md"
dest.write_text("\n".join(out), encoding="utf-8")
print("wrote", dest, "with", n, "slides")
