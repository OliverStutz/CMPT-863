#!/usr/bin/env python3
"""Concatenate slides3 section files 01 to 06 into deck source files (with notes, and presentation only).
Drops the per-file 'Verification log' and 'Open items' sections. Re-run after editing: python3 slides3/_build-deck-source.py (from week7/)
"""
import re, glob, datetime, pathlib
here = pathlib.Path(__file__).parent
out = here.parent / "Week 7 - Deck Source v3.md"
files = sorted(f for f in glob.glob(str(here / "0*-*.md")) if re.search(r"/0[1-6]-", f))
TITLE = """
## Slide 1: What Should We Expect From Research Through Design?

**Time:** 0:30 · **Visual:** the title question large across the upper half; a strip cropped from Figure 1 of the paper across the lower third with "Figure 1, Gaver 2012" in small type beneath it; the body lines in small type between.
**Images:** `assets/rams-fig1-annotated.png`, cropped to a horizontal strip.
**Build:** all at once.

**On the slide**
- William Gaver, Interaction Research Studio, Goldsmiths, University of London
- CHI 2012, Austin, Texas, session Critical Perspectives on Design
- CMPT 863 Human-Computer Interaction · Week 7 · 22 October 2026
- Presented by Oliver Stutz
- Figure 1, Gaver 2012

**Speaker notes**
- **Say:** Good morning. In 2012 William Gaver, at the Interaction Research Studio at Goldsmiths in London, answered colleagues who wanted research through design to adopt standards. Twenty-five minutes from me, in six parts: where the field came from, what people had tried, what Gaver argues, whether it holds, what became of it, and what it is worth now. I will need your hands four times. We start with a definition and an object. (about 65 words)
- **Transition:** First, what research through design is, and a coffee table.
"""
parts = ["# Week 7 deck source, version 3: sections 01 to 06\n",
         f"Generated {datetime.date.today().isoformat()} from the section files in `slides3/` by `slides3/_build-deck-source.py`. "
         "Edit the section files, not this file, then regenerate. Verification logs and open items live in the section files. "
         "Run sheet and plan are in `slides3/00-Overview.md`; the slide text alone is extracted to `slides3/slides-only/` by `slides3/_extract-slides-only.py`.\n",
         TITLE]
for f in files:
    text = pathlib.Path(f).read_text(encoding="utf-8")
    text = re.split(r"\n## Verification log\n", text, maxsplit=1)[0].rstrip() + "\n"
    parts.append("\n---\n\n" + text)
full = "\n".join(parts)
out.write_text(full, encoding="utf-8"); print("wrote", out, "from", len(files), "section files")
pres = re.sub(r"\n\*\*Speaker notes\*\*\n.*?(?=\n## |\n---\n|\Z)", "\n", full, flags=re.S)
pres = pres.replace("# Week 7 deck source, version 3: sections 01 to 06", "# Week 7 deck source, version 3, presentation only: sections 01 to 06", 1)
pres = re.sub(r"\n{3,}", "\n\n", pres)
out2 = here.parent / "Week 7 - Deck Source v3 - Presentation only.md"
out2.write_text(pres, encoding="utf-8"); print("wrote", out2)
