package main

import (
	"bytes"
	"os"
	"testing"
)

func readFixture(t *testing.T, name string) []byte {
	t.Helper()
	b, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestDurationGoogleClip(t *testing.T) {
	// 24000 bytes of MPEG-2 Layer III, 64 kbps, 24 kHz mono: 125 frames of
	// 192 bytes, 576 samples each = 72000 samples = 3000 ms (ffprobe agrees).
	b := readFixture(t, "google-neural2-j.mp3")
	d, err := durationMs(b)
	if err != nil {
		t.Fatal(err)
	}
	if d != 3000 {
		t.Fatalf("duration %d ms, want 3000", d)
	}
	s, err := parseMP3(b)
	if err != nil {
		t.Fatal(err)
	}
	if len(s.frames) != 125 || s.info.sampleRate != 24000 || !s.info.mono || s.info.version != 2 {
		t.Fatalf("parsed %d frames, %+v", len(s.frames), s.info)
	}
}

func TestDurationLameClipSkipsInfoFrame(t *testing.T) {
	// A LAME file starts with an Info frame, which parseMP3 drops.
	b := readFixture(t, "lame-sine-1s.mp3")
	s, err := parseMP3(b)
	if err != nil {
		t.Fatal(err)
	}
	want := int64(float64(len(s.frames))*s.info.frameDurationMs() + 0.5)
	d, err := durationMs(s.bytes())
	if err != nil {
		t.Fatal(err)
	}
	if d != want {
		t.Fatalf("duration %d ms, want %d", d, want)
	}
	if d < 1000 || d > 1100 {
		t.Fatalf("duration %d ms for a 1 s tone", d)
	}
}

func TestDurationEmptyAndGarbage(t *testing.T) {
	if d, err := durationMs(nil); err != nil || d != 0 {
		t.Fatalf("empty: %d %v", d, err)
	}
	if _, err := parseMP3([]byte("not an mp3 file at all")); err == nil {
		t.Fatal("garbage parsed")
	}
}

func TestJoinSharedPartsAreByteIdentical(t *testing.T) {
	g := readFixture(t, "google-neural2-j.mp3")
	tone, _ := parseMP3(readFixture(t, "lame-sine-1s.mp3"))
	// Re-encode-free stand-ins for opening, body and closing. The tone file is
	// MPEG-2 24 kHz mono too, so frames can be mixed.
	opening := g
	body := tone.bytes()
	closing := g[:192*50] // first 50 frames

	reason, rStarts, err := joinMP3([][]byte{opening, body, closing}, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	denial, dStarts, err := joinMP3([][]byte{opening, closing}, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(reason, opening) || !bytes.HasPrefix(denial, opening) {
		t.Fatal("opening not byte-identical at the start of both clips")
	}
	if !bytes.HasSuffix(reason, closing) || !bytes.HasSuffix(denial, closing) {
		t.Fatal("closing not byte-identical at the end of both clips")
	}
	if rStarts[0] != 0 || rStarts[1] != 3000 || rStarts[2] != 3000+int64(len(tone.frames))*24 {
		t.Fatalf("reason starts %v", rStarts)
	}
	if dStarts[1] != 3000 {
		t.Fatalf("denial starts %v", dStarts)
	}
	rd, _ := durationMs(reason)
	dd, _ := durationMs(denial)
	if rd != 3000+int64(len(tone.frames))*24+1200 || dd != 4200 {
		t.Fatalf("durations %d %d", rd, dd)
	}
}

func TestJoinGapAndPadding(t *testing.T) {
	g := readFixture(t, "google-neural2-j.mp3")
	out, starts, err := joinMP3([][]byte{g, g}, 600, 0)
	if err != nil {
		t.Fatal(err)
	}
	if starts[1] != 3600 {
		t.Fatalf("second part starts at %d, want 3600", starts[1])
	}
	d, _ := durationMs(out)
	if d != 6600 {
		t.Fatalf("joined duration %d, want 6600", d)
	}
	// Pad 1000 ms: 1000/24 = 41.7 -> 42 frames = 1008 ms.
	padded, _, err := joinMP3([][]byte{g}, 0, 1000)
	if err != nil {
		t.Fatal(err)
	}
	d, _ = durationMs(padded)
	if d != 4008 {
		t.Fatalf("padded duration %d, want 4008", d)
	}
	if !bytes.HasPrefix(padded, g) {
		t.Fatal("padding changed the audio")
	}
}

func TestSilenceFrameDecodesToSilence(t *testing.T) {
	s, _ := parseMP3(readFixture(t, "google-neural2-j.mp3"))
	sil := bytes.Join(s.silenceFrames(10), nil)
	pcm := decodeAll(t, sil)
	for i, v := range pcm {
		if v != 0 {
			t.Fatalf("sample %d is %d, want 0", i, v)
		}
	}
	if len(pcm) != 10*576*2 { // stereo output
		t.Fatalf("got %d samples", len(pcm))
	}
}

func TestJoinRejectsMixedFormats(t *testing.T) {
	g := readFixture(t, "google-neural2-j.mp3")
	// MPEG-1 44.1 kHz stereo header, 128 kbps: frame size 417 bytes.
	hdr := []byte{0xFF, 0xFB, 0x90, 0x00}
	f := make([]byte, 417)
	copy(f, hdr)
	if _, _, err := joinMP3([][]byte{g, f}, 0, 0); err == nil {
		t.Fatal("mixed formats joined")
	}
}
