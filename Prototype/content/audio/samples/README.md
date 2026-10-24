# Voice samples (decision 3.6)

One sample per candidate voice, US English, from the real v1.4 content. Every candidate is a Neural2 or WaveNet voice, and every sample was synthesised with one SSML `<mark>`; all nine returned the mark's timepoint, so all nine support timed flags.

Sample sentences:

- CISO: the greeting and the first sentence of narration 1: "Thank you, Margaret. Good morning, everyone. I will keep this short, because I am asking you for a decision today."
- Chair: the first opening line: "Good morning, everyone, and thank you for being here. Let us move to item four on our agenda: the protection of our network."
- Board member: the participant's explain question on slide 2: "Daniel, could you explain how you get to eight million dollars for one attack?"

| Role | File | Voice id | Length | Rough median pitch |
|---|---|---|---|---|
| CISO | `ciso-en-US-Neural2-J.mp3` | en-US-Neural2-J | 6.9 s | 146 Hz |
| CISO | `ciso-en-US-Neural2-D.mp3` | en-US-Neural2-D | 6.6 s | 142 Hz |
| CISO | `ciso-en-US-Wavenet-B.mp3` | en-US-Wavenet-B | 7.7 s | 148 Hz |
| Chair | `chair-en-US-Neural2-F.mp3` | en-US-Neural2-F | 8.0 s | 208 Hz |
| Chair | `chair-en-US-Neural2-H.mp3` | en-US-Neural2-H | 7.2 s | 232 Hz |
| Chair | `chair-en-US-Neural2-C.mp3` | en-US-Neural2-C | 7.8 s | 154 Hz |
| Board member | `member-en-US-Neural2-E.mp3` | en-US-Neural2-E | 4.8 s | 172 Hz |
| Board member | `member-en-US-Neural2-A.mp3` | en-US-Neural2-A | 4.8 s | 158 Hz |
| Board member | `member-en-US-Neural2-I.mp3` | en-US-Neural2-I | 4.5 s | 133 Hz |

The pitch column is a crude autocorrelation estimate, useful only to compare voices with each other. `samples.json` has the durations and mark offsets.

Note: en-US-Wavenet-D and en-US-Wavenet-J return byte-identical audio to en-US-Neural2-D and en-US-Neural2-J, so they are the same voices under another name. en-US-Wavenet-B replaces Wavenet-D as the third CISO candidate.

## Recommendation

These are picked without a listening test; Oliver decides by listening.

- CISO: **en-US-Neural2-J**, a mature, measured male voice that suits a senior executive asking for a budget.
- Chair: **en-US-Neural2-F**, a clear female voice, about 60 Hz higher than the CISO and a different timbre, so the handover is unmistakable.
- Board member: **en-US-Neural2-E**, a lower female voice that sits between the other two in pitch and differs in sex from the CISO, so the question and the CISO's answer that follows it never sound like one speaker.

The full v1.4 set in `content/audio/1.4.0/` uses these three voices (set in `tools/tts/voices.default.json`).

## Regenerating

From `Prototype/`:

```
go run ./tools/tts -config content/study-config.v1.4.json -samples tools/tts/voice-candidates.json
```

Edit `tools/tts/voice-candidates.json` to try other voices. The MP3 files are git-ignored (root `.gitignore`), so they exist only where they were generated.
