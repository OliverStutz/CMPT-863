import math, sys, os
from xml.sax.saxutils import escape
from PIL import ImageFont

OUT = sys.argv[1]
FONT = 'Inter, Helvetica, Arial, sans-serif'
BRAND = '#14325A'; ACCENT = '#D9822B'; INK = '#1F2933'; MUTED = '#5F6B7A'; TINT = '#EEF2F7'; LINE = '#C9D2DE'
W, H = 1920, 1080
_fonts = {}
def width(text, size, bold=False):
    m = 0
    for base in (['/usr/share/fonts/gsfonts/NimbusSans-%s.otf', '/usr/share/fonts/noto/NotoSans-%s.ttf']):
        p = base % ('Bold' if bold else 'Regular')
        k = (p, size)
        if k not in _fonts: _fonts[k] = ImageFont.truetype(p, size)
        m = max(m, _fonts[k].getlength(text))
    return m
def wrap(text, size, maxw, bold=False):
    words = text.split(' '); lines = []; cur = ''
    for w in words:
        t = (cur + ' ' + w).strip()
        if width(t, size, bold) <= maxw: cur = t
        else: lines.append(cur); cur = w
    lines.append(cur); return lines
def e(s): return escape(s)

def frame(n, title):
    s = [f'<svg xmlns="http://www.w3.org/2000/svg" width="{W}" height="{H}" viewBox="0 0 {W} {H}" font-family="{FONT}">',
         f'<rect width="{W}" height="{H}" fill="#FFFFFF"/>',
         f'<rect width="{W}" height="150" fill="{BRAND}"/>',
         f'<rect y="150" width="{W}" height="6" fill="{ACCENT}"/>',
         f'<text id="s{n}.title" x="80" y="97" font-size="58" font-weight="700" fill="#FFFFFF">{e(title)}</text>',
         f'<g id="brand"><rect x="1556" y="62" width="22" height="22" fill="{ACCENT}"/>'
         f'<text x="1840" y="81" font-size="26" font-weight="600" fill="#FFFFFF" text-anchor="end" letter-spacing="0.5">Meridian Logistics</text></g>',
         f'<line x1="80" y1="1012" x2="1840" y2="1012" stroke="{LINE}" stroke-width="2"/>',
         f'<text id="slide-number" x="1840" y="1050" font-size="24" fill="{MUTED}" text-anchor="end">{n}</text>']
    return s

def bullets(items, x, y, maxw, size=38, lh=52, gap=34):
    out = []
    for sid, text in items:
        lines = wrap(text, size, maxw - 44)
        out.append(f'<g id="{sid}"><rect x="{x}" y="{y - size*0.42 - 7}" width="14" height="14" fill="{ACCENT}"/>')
        out.append(f'<text x="{x+44}" y="{y}" font-size="{size}" fill="{INK}">' +
                   ''.join(f'<tspan x="{x+44}" dy="{0 if i==0 else lh}">{e(l)}</tspan>' for i, l in enumerate(lines)) + '</text></g>')
        y += lh * len(lines) + gap
    return out, y

def save(n, parts):
    parts.append('</svg>')
    open(os.path.join(OUT, f's{n}.svg'), 'w').write('\n'.join(parts) + '\n')

# Slide 1
p = frame(1, 'Where we stand today')
b, _ = bullets([('s1.b1', 'Our 12 warehouses and head office run on one network: orders, scanners, trucks and payroll'),
                ('s1.b2', 'This year we stopped two attempted break-ins into that network'),
                ('s1.b3', 'Attackers are getting better'),
                ('s1.b4', 'Our insurer and our largest customer now expect stronger protection')], 80, 290, 1040)
p += b
cx, cy, cw, ch = 1220, 250, 620, 360
lab = wrap('budget request for approval today', 40, cw - 90)
p.append(f'<g id="s1.num1"><rect x="{cx}" y="{cy}" width="{cw}" height="{ch}" fill="{TINT}"/>'
         f'<rect x="{cx}" y="{cy}" width="10" height="{ch}" fill="{ACCENT}"/>'
         f'<text x="{cx+50}" y="{cy+150}" font-size="80" font-weight="700" fill="{BRAND}">$2.4 million:</text>'
         f'<text x="{cx+50}" y="{cy+230}" font-size="40" fill="{INK}">' +
         ''.join(f'<tspan x="{cx+50}" dy="{0 if i==0 else 54}">{e(l)}</tspan>' for i, l in enumerate(lab)) + '</text></g>')
save(1, p)

# Slide 2
p = frame(2, 'The risk and what it would cost us')
rows = [('s2.num1', '$1.2 million:', 'goods we ship each working day'),
        ('s2.num2', '5 days:', 'how long a serious attack stops a company like ours'),
        ('s2.num3', '$8 million:', 'cost to us of one serious attack')]
y = 220
for sid, fig, lab in rows:
    p.append(f'<g id="{sid}"><rect x="80" y="{y}" width="1760" height="118" fill="{TINT}"/>'
             f'<rect x="80" y="{y}" width="10" height="118" fill="{ACCENT}"/>'
             f'<text y="{y+80}" fill="{INK}"><tspan x="130" font-size="64" font-weight="700" fill="{BRAND}">{e(fig)}</tspan>'
             f'<tspan x="640" font-size="40"> {e(lab)}</tspan></text></g>')
    y += 140
b, _ = bullets([('s2.b1', 'Our insurance renews in 10 months: without stronger protection, the premium rises or cover ends'),
                ('s2.b2', "Our largest customer's new contract requires the same protection")], 80, 720, 1760)
p += b
save(2, p)

# Slide 3
p = frame(3, 'How we will protect our 13 sites')
def shield(x, y, s=1.0):
    return (f'<path transform="translate({x},{y}) scale({s})" d="M0,-15 L13,-10 L13,0 C13,9 7,14 0,17 C-7,14 -13,9 -13,0 L-13,-10 Z" '
            f'fill="{ACCENT}" stroke="#FFFFFF" stroke-width="2"/>'
            f'<path transform="translate({x},{y}) scale({s})" d="M-5,1 L-1,5 L6,-4" fill="none" stroke="#FFFFFF" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"/>')
hx, hy = 560, 610
rx, ry = 430, 235
g = ['<g id="s3.img1">', f'<rect x="60" y="186" width="1010" height="810" fill="#F7F9FC"/>']
pos = []
for i in range(12):
    a = math.radians(15 + 30*i)
    pos.append((hx + rx*math.cos(a), hy - ry*math.sin(a)))
# cloud and laptop positions
clx, cly = hx, 262
lpx, lpy = hx, 930
g.append(f'<g id="s3.img1.links">')
for (x, y) in pos:
    g.append(f'<line x1="{hx}" y1="{hy}" x2="{x:.0f}" y2="{y:.0f}" stroke="{BRAND}" stroke-width="3"/>')
g.append(f'<line x1="{hx}" y1="{hy}" x2="{clx}" y2="{cly}" stroke="{BRAND}" stroke-width="4"/>')
g.append(f'<line x1="{hx}" y1="{hy}" x2="{lpx}" y2="{lpy-30}" stroke="{BRAND}" stroke-width="3" stroke-dasharray="10 9"/>')
g.append('</g>')
for i, (x, y) in enumerate(pos):
    x0, y0 = x - 66, y - 38
    g.append(f'<g id="s3.img1.wh{i+1}"><rect x="{x0:.0f}" y="{y0:.0f}" width="132" height="76" rx="6" fill="#FFFFFF" stroke="{BRAND}" stroke-width="2.5"/>'
             f'<path d="M{x-26:.0f},{y+6:.0f} L{x-26:.0f},{y-10:.0f} L{x-13:.0f},{y-18:.0f} L{x-13:.0f},{y-10:.0f} L{x:.0f},{y-18:.0f} L{x:.0f},{y-10:.0f} L{x+13:.0f},{y-18:.0f} L{x+13:.0f},{y-10:.0f} L{x+26:.0f},{y-18:.0f} L{x+26:.0f},{y+6:.0f} Z" fill="{BRAND}"/>'
             f'<text x="{x:.0f}" y="{y+29:.0f}" font-size="18" fill="{INK}" text-anchor="middle">Warehouse {i+1}</text>'
             + shield(x0 + 132, y0, 0.9) + '</g>')
# head office
g.append(f'<g id="s3.img1.hq"><rect x="{hx-120}" y="{hy-62}" width="240" height="124" rx="8" fill="{BRAND}"/>'
         f'<text x="{hx}" y="{hy-12}" font-size="26" font-weight="700" fill="#FFFFFF" text-anchor="middle">Head office</text>'
         f'<text x="{hx}" y="{hy+24}" font-size="20" fill="#FFFFFF" text-anchor="middle">Main data centre</text>'
         + shield(hx - 80, hy - 62, 1.2) + '</g>')
# cloud
c = (f'M{clx-90},{cly+38} C{clx-130},{cly+38} {clx-130},{cly-14} {clx-88},{cly-12} '
     f'C{clx-84},{cly-52} {clx-30},{cly-60} {clx-10},{cly-30} C{clx+10},{cly-62} {clx+70},{cly-56} {clx+74},{cly-14} '
     f'C{clx+124},{cly-18} {clx+130},{cly+38} {clx+88},{cly+38} Z')
g.append(f'<g id="s3.img1.internet"><path d="{c}" fill="#FFFFFF" stroke="{BRAND}" stroke-width="3"/>'
         f'<text x="{clx}" y="{cly+18}" font-size="24" font-weight="700" fill="{BRAND}" text-anchor="middle">Internet</text></g>')
# laptop
g.append(f'<g id="s3.img1.remote"><rect x="{lpx-40}" y="{lpy-30}" width="80" height="50" rx="4" fill="#FFFFFF" stroke="{BRAND}" stroke-width="3"/>'
         f'<rect x="{lpx-32}" y="{lpy-23}" width="64" height="36" fill="{TINT}"/>'
         f'<path d="M{lpx-56},{lpy+22} L{lpx+56},{lpy+22} L{lpx+48},{lpy+32} L{lpx-48},{lpy+32} Z" fill="{BRAND}"/>'
         f'<text x="{lpx+72}" y="{lpy+12}" font-size="22" font-weight="700" fill="{BRAND}">Remote staff</text></g>')
g.append('</g>')
p += g
b, _ = bullets([('s3.b1', 'A firewall at every site keeps warehouse machines apart from office computers'),
                ('s3.b2', 'All 12 warehouses connect to our data centre at head office'),
                ('s3.b3', 'One internet connection, at head office, watched day and night'),
                ('s3.b4', 'Staff at home or on the road use secure remote access')], 1130, 270, 710, size=34, lh=46, gap=40)
p += b
save(3, p)

# Slide 4
p = frame(4, 'What the $2.4 million pays for')
rows = [('s4.num1', 'Firewalls for 13 sites, with 3 years of support:', '$780,000'),
        ('s4.num2', '24-hour monitoring of our network, 2 years:', '$960,000'),
        ('s4.num3', 'Secure remote access for 1,400 staff:', '$310,000'),
        ('s4.num4', 'Installation and staff training:', '$350,000')]
y = 230; rh = 126
for i, (sid, d, a) in enumerate(rows):
    fill = '#FFFFFF' if i % 2 else '#F7F9FC'
    p.append(f'<g id="{sid}"><rect x="80" y="{y}" width="1760" height="{rh}" fill="{fill}"/>'
             f'<line x1="80" y1="{y+rh}" x2="1840" y2="{y+rh}" stroke="{LINE}" stroke-width="2"/>'
             f'<text y="{y+77}" font-size="44" fill="{INK}"><tspan x="120">{e(d)}</tspan>'
             f'<tspan x="1800" text-anchor="end"> {e(a)}</tspan></text></g>')
    y += rh
y += 14
p.append(f'<g id="s4.num5"><rect x="80" y="{y}" width="1760" height="{rh}" fill="{TINT}"/>'
         f'<rect x="80" y="{y}" width="1760" height="5" fill="{BRAND}"/>'
         f'<text y="{y+80}" font-size="48" font-weight="700" fill="{BRAND}"><tspan x="120">Total:</tspan>'
         f'<tspan x="1800" text-anchor="end"> $2,400,000</tspan></text></g>')
save(4, p)

# Slide 5
p = frame(5, 'The decision today')
p.append(f'<g id="s5.num1"><rect x="80" y="230" width="1760" height="200" fill="{BRAND}"/>'
         f'<rect x="80" y="230" width="14" height="200" fill="{ACCENT}"/>'
         f'<text x="150" y="358" font-size="80" fill="#FFFFFF"><tspan font-weight="700">$2.4 million</tspan> for approval today</text></g>')
b, _ = bullets([('s5.b1', 'Firewalls at all 13 sites, 24-hour monitoring, secure remote access and staff training'),
                ('s5.b2', 'Ready in 8 months, before our insurance renews in 10 months'),
                ('s5.b3', "If we wait: a higher premium or no cover, and our largest customer's contract at risk")], 80, 540, 1760, size=40, lh=54, gap=50)
p += b
save(5, p)
