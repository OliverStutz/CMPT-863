"""The cast: one synthetic portrait per role (no real person), prompts for Flux and for the speech-to-video
model. Roles follow the participant app's seats: ciso, chair, member2..member6 (member = the participant's
own seat, never given a face). Change a prompt or seed here and re-run the batch for that role."""
STYLE = ("photorealistic head-and-shoulders portrait photograph, looking straight at the camera, mouth closed, neutral calm "
         "expression, soft even lighting, blurred modern boardroom background, 85 mm lens, natural skin texture, centred, "
         "shoulders visible, sharp focus")
CAST = {
    "ciso":    {"seed": 20261025, "look": "a man in his mid-forties, chief information security officer, short dark hair with a little grey at the temples, clean-shaven, navy suit, white shirt, no tie",
                "talk": "A man in his mid-forties in a navy suit sits at a boardroom table and speaks calmly and confidently to the board, looking at the camera, lip movements matching the speech, small natural head movements and blinks, static camera, soft even lighting, photorealistic."},
    "chair":   {"seed": 20261027, "look": "a woman in her early sixties of East Asian heritage, chair of the board, short silver-grey hair, dark grey suit, light blouse",
                "talk": "A woman in her early sixties in a dark grey suit chairs a board meeting and speaks warmly and clearly to the board, looking at the camera, lip movements matching the speech, small natural head movements and blinks, static camera, soft even lighting, photorealistic."},
    "member2": {"seed": 20261026, "look": "a woman in her late fifties, company board director, shoulder-length grey hair, dark blazer, light blouse"},
    "member3": {"seed": 20261028, "look": "a man in his fifties of South Asian heritage, company board director, short black hair with grey at the temples, charcoal suit, white shirt, no tie"},
    "member4": {"seed": 20261029, "look": "a Black woman in her mid-forties, company board director, natural short hair, navy blazer, cream blouse"},
    "member5": {"seed": 20261030, "look": "a man in his mid-sixties, company board director, bald with a trimmed grey beard, glasses, dark blue suit, light blue shirt"},
    "member6": {"seed": 20261031, "look": "a woman in her late thirties of Hispanic heritage, company board director, dark hair tied back, black blazer, white top"},
}
def portrait_prompt(role): return f"{CAST[role]['look']}, {STYLE}"
def idle_prompt(role):
    return (f"{CAST[role]['look'].split(',')[0].capitalize()} sits at a boardroom table listening attentively, mouth closed, does not speak, "
            "occasional blinks and very small natural head movements, static camera, soft even lighting, photorealistic.")
