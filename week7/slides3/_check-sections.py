#!/usr/bin/env python3
"""Mechanical checks on the slides3 section files: dashes, Say words per minute, on-slide words, [classmate] slots.
Run from anywhere: python3 slides3/_check-sections.py
"""
import re, glob, pathlib
here = pathlib.Path(__file__).parent
files = sorted(f for f in glob.glob(str(here / "0*-*.md")) if re.search(r"/0[1-6]-", f))
def mmss(s):
    m, sec = s.split(":"); return int(m) + int(sec) / 60
total_time = 0.0; total_say = 0; total_on = 0
for f in files:
    t = pathlib.Path(f).read_text(encoding="utf-8")
    em = t.count("—"); en = t.count("–")
    body = t.split("\n## Verification log\n")[0]
    print(f"{pathlib.Path(f).name}: em={em} en={en} classmate={body.count('[classmate]')}")
    for m in re.finditer(r"## Slide (\d+): (.+?)\n\n\*\*Time:\*\* (\d+:\d\d)(.*?)(?=\n## Slide |\n## Verification log|\Z)", body, re.S):
        num, title, tm, rest = m.groups()
        on = re.search(r"\*\*On the slide\*\*\n(.*?)\n\*\*Speaker notes\*\*", rest, re.S)
        on_words = len(re.sub(r"\(Gaver, p\. ?[\d, to]+\)", "", on.group(1)).split()) if on else -1
        say = re.search(r"\*\*Say:\*\* (.*?)(?=\n- \*\*Transition:\*\*)", rest, re.S)
        say_words = len(re.sub(r"\(about \d+ words\)", "", say.group(1)).split()) if say else -1
        limit = int(130 * mmss(tm)); total_time += mmss(tm); total_say += max(say_words, 0); total_on += max(on_words, 0)
        flag = "" if say_words <= limit else "  <-- OVER"
        flag2 = "" if on_words <= 65 else "  <-- SLIDE DENSE"
        print(f"   {num:>2} {tm}  on-slide {on_words:3d}w{flag2}  say {say_words:3d}w / limit {limit}{flag}   {title[:55]}")
print(f"\nSections total: {int(total_time)}:{int(round((total_time%1)*60)):02d}, {total_say} spoken words, {total_on} on-slide words")
