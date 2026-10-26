"""Self-contained preview of the boardroom scene with the round-one clips, using the participant app's
real stage geometry (BoardroomScene.tsx / .module.css): 1600x900 stage, screen at (40,36) 1088x612, CISO
frame at (1240,470) 260 wide, seven 130 px seats at y=660. Everything is embedded as data URIs."""
import base64, io, os
from PIL import Image
LAB = os.environ.get("FACES_LAB", "/home/oliver/.cache/cmpt863-faces"); P = os.environ.get("PROTOTYPE_DIR", os.path.join(os.path.dirname(os.path.abspath(__file__)), "..", "..", ".."))
def b64(path, mime): return f"data:{mime};base64," + base64.b64encode(open(path, "rb").read()).decode()
def poster(path, size=400):
    im = Image.open(path).convert("RGB").resize((size, size)); buf = io.BytesIO(); im.save(buf, "JPEG", quality=85)
    return "data:image/jpeg;base64," + base64.b64encode(buf.getvalue()).decode()
slide = b64(f"{P}/content/slides/v1.3/s2.png", "image/png")
ciso = b64(f"{LAB}/out/s2v/ciso_00001__trim.mp4", "video/mp4")
member = b64(f"{LAB}/out/s2v/member_00001_.mp4", "video/mp4")
ciso_poster = poster(f"{LAB}/inputs/ref_ciso.png"); member_poster = poster(f"{LAB}/inputs/ref_member.png")
sil = lambda bg, fg, r: (f'<svg viewBox="0 0 200 200"><rect width="200" height="200" fill="{bg}"/><circle cx="100" cy="78" r="{r}" fill="{fg}"/>'
                         f'<path d="M28 200c4-52 34-78 72-78s68 26 72 78z" fill="{fg}"/></svg>')
TONES = [("#4a5262","#9aa3b2",38),("#4d5468","#aab3c1",37),("#465064","#8e98a8",39),("#505866","#a3abb8",36),("#434b5c","#97a1b0",40),("#4b5366","#a6afbd",37),("#485263","#919bab",39)]
seats = [("chair","Chair",30,0,None),("you","Board member (You)",208,1,None),("m2","Board member",385,2,"member"),
         ("m3","Board member",563,3,None),("m4","Board member",740,4,None),("m5","Board member",918,5,None),("m6","Board member",1095,6,None)]
seat_html = ""
for key, label, left, tone, video in seats:
    bg, fg, r = TONES[tone]
    inner = (f'<video class="idle" muted loop autoplay playsinline poster="{member_poster}" src="{member}"></video>' if video
             else sil(bg, fg, r))
    badge = '<span class="badge">You</span>' if key == "you" else ""
    seat_html += f'<figure class="portrait seat {key}" style="left:{left}px"><div class="img">{inner}{badge}</div><figcaption>{label}</figcaption></figure>\n'
html = f"""<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>Boardroom preview</title>
<style>
html,body{{margin:0;height:100%;background:#14171d;font-family:Inter,Helvetica,Arial,sans-serif}}
.outer{{position:fixed;inset:0;overflow:hidden}}
.stage{{position:absolute;left:50%;top:50%;width:1600px;height:900px;transform-origin:50% 50%}}
.room{{position:absolute;inset:0;width:1600px;height:900px}}
.screenFrame{{position:absolute;left:30px;top:26px;width:1108px;height:632px;background:#0c0e12;border-radius:6px;box-shadow:0 10px 40px rgba(0,0,0,.55)}}
.screen{{position:absolute;left:40px;top:36px;width:1088px;height:612px;background:#fff;overflow:hidden}}
.screen img{{width:100%;height:100%;display:block}}
.table{{position:absolute;left:0;top:812px;width:1600px;height:88px;background:linear-gradient(#5a4130,#3e2c20);box-shadow:0 -4px 18px rgba(0,0,0,.5)}}
.portrait{{position:absolute;margin:0;text-align:center;color:#f4f6f8;font-size:20px;font-weight:600}}
.portrait figcaption{{margin-top:8px;line-height:1.15;text-shadow:0 1px 2px rgba(0,0,0,.8)}}
.img{{position:relative;width:100%;aspect-ratio:1/1;border-radius:12px;overflow:hidden;border:3px solid rgba(255,255,255,.15);transition:box-shadow .4s,border-color .4s,filter .4s;filter:brightness(.85);background:#2a2f39}}
.img svg,.img img,.img video{{width:100%;height:100%;display:block;object-fit:cover}}
.speaking .img{{border-color:#f2c879;box-shadow:0 0 0 4px rgba(242,200,121,.35),0 0 36px rgba(242,200,121,.55);filter:none}}
.bars{{position:absolute;right:10px;bottom:10px;display:none;gap:4px;align-items:flex-end;height:26px}}
.speaking .bars{{display:flex}}
.bars i{{display:block;width:6px;background:#f2c879;border-radius:2px;animation:bar .9s ease-in-out infinite}}
.bars i:nth-child(2){{animation-delay:.15s}}.bars i:nth-child(3){{animation-delay:.3s}}
@keyframes bar{{0%,100%{{height:6px}}50%{{height:24px}}}}
.ciso{{left:1240px;top:470px;width:260px;transition:left .3s,top .3s,width .3s}}
.large .ciso{{left:1170px;top:250px;width:400px}}
.seat{{top:660px;width:130px;font-size:18px}}
.you .img{{border-color:rgba(143,208,255,.85);box-shadow:0 0 0 3px rgba(143,208,255,.25);filter:brightness(.95)}}
.you figcaption{{color:#cfe9ff}}
.badge{{position:absolute;left:8px;top:8px;padding:2px 8px;border-radius:999px;background:#8fd0ff;color:#10233a;font-size:16px;font-weight:700;box-shadow:0 1px 3px rgba(0,0,0,.5)}}
.controls{{position:absolute;left:630px;top:676px;width:450px;min-height:110px;background:rgba(255,255,255,.96);color:#1f2933;border-radius:10px;padding:14px 22px;box-sizing:border-box;font-size:18px;display:flex;flex-direction:column;gap:8px;justify-content:center}}
.controls button{{font:inherit;font-weight:600;padding:8px 14px;border-radius:8px;border:1px solid #14325a;background:#14325a;color:#fff;cursor:pointer}}
.controls label{{font-size:16px;display:flex;gap:8px;align-items:center}}
.note{{position:absolute;left:1160px;top:36px;width:400px;color:#cfd5de;font-size:16px;line-height:1.4;background:rgba(0,0,0,.35);border-radius:8px;padding:12px 14px;box-sizing:border-box}}
</style></head><body><div class="outer"><div class="stage" id="stage">
<svg class="room" viewBox="0 0 1600 900"><defs><linearGradient id="wall" x1="0" y1="0" x2="0" y2="1"><stop offset="0" stop-color="#3b4250"/><stop offset="1" stop-color="#2a2f39"/></linearGradient></defs>
<rect width="1600" height="900" fill="url(#wall)"/><rect y="660" width="1600" height="240" fill="#1d2129"/>
{''.join(f'<rect x="{i*200+8}" y="0" width="184" height="660" fill="#fff" opacity="0.025"/>' for i in range(8))}</svg>
<div class="screenFrame"></div><div class="screen"><img src="{slide}" alt="Slide 2"></div>
<div class="note"><b>Preview, round one.</b> Real stage geometry of the participant app. Slide 2 is shown; the CISO answers the challenge on this slide. The board member in seat 3 is the silent idle loop. Other seats keep today's placeholders.</div>
<figure class="portrait ciso" id="ciso"><div class="img"><video id="cisoVideo" playsinline preload="auto" poster="{ciso_poster}" src="{ciso}"></video><span class="bars"><i></i><i></i><i></i></span></div><figcaption>CISO</figcaption></figure>
<div class="table"></div>
{seat_html}
<div class="controls"><button id="play">Play the CISO's answer</button>
<label><input type="checkbox" id="large"> Larger speaker frame (400 px) beside the slide</label></div>
</div></div>
<script>
const stage=document.getElementById('stage');function fit(){{const s=Math.min(innerWidth/1600,innerHeight/900);stage.style.transform=`translate(-50%,-50%) scale(${{s}})`}}addEventListener('resize',fit);fit();
const v=document.getElementById('cisoVideo'),fig=document.getElementById('ciso'),btn=document.getElementById('play');
btn.onclick=()=>{{v.currentTime=0;v.muted=false;v.play();}};
v.onplay=()=>{{fig.classList.add('speaking');btn.textContent='Playing…'}};v.onended=()=>{{fig.classList.remove('speaking');btn.textContent='Replay the answer'}};v.onpause=()=>{{if(!v.ended){{fig.classList.remove('speaking');btn.textContent='Play the CISO\\'s answer'}}}};
document.getElementById('large').onchange=e=>stage.classList.toggle('large',e.target.checked);
document.querySelectorAll('video.idle').forEach(x=>x.play().catch(()=>{{}}));
</script></body></html>"""
out = f"{LAB}/out/preview/boardroom-preview.html"; open(out, "w").write(html); print(out, round(len(html)/1e6, 1), "MB")
