package main

import (
	"bytes"
	"errors"
	"fmt"

	gomp3 "github.com/hajimehoshi/go-mp3"
)

// Frame-level MP3 handling. Joined clips are built by concatenating the
// Layer III frames of their parts (no re-encoding), so a shared sentence is
// byte-identical in every clip that contains it. Silence is made of frames
// whose side information is all zero, which decode to digital silence.

// frameInfo describes one MPEG audio Layer III frame header.
type frameInfo struct {
	version     int // 1 = MPEG-1, 2 = MPEG-2, 25 = MPEG-2.5
	bitrate     int // bits per second
	sampleRate  int
	padding     bool
	crc         bool
	mono        bool
	size        int // bytes, header included
	samples     int // samples per channel in this frame
	sideInfoLen int
}

var (
	bitratesV1 = [16]int{0, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320, -1}
	bitratesV2 = [16]int{0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160, -1}
	rates      = map[int][3]int{1: {44100, 48000, 32000}, 2: {22050, 24000, 16000}, 25: {11025, 12000, 8000}}
)

func parseHeader(h []byte) (frameInfo, error) {
	var f frameInfo
	if len(h) < 4 || h[0] != 0xFF || h[1]&0xE0 != 0xE0 {
		return f, errors.New("no frame sync")
	}
	switch (h[1] >> 3) & 3 {
	case 3:
		f.version = 1
	case 2:
		f.version = 2
	case 0:
		f.version = 25
	default:
		return f, errors.New("reserved MPEG version")
	}
	if (h[1]>>1)&3 != 1 {
		return f, errors.New("not Layer III")
	}
	f.crc = h[1]&1 == 0
	bi := int(h[2] >> 4)
	si := int((h[2] >> 2) & 3)
	if bi == 0 || bi == 15 || si == 3 {
		return f, errors.New("free-format or invalid bitrate or sample rate")
	}
	if f.version == 1 {
		f.bitrate = bitratesV1[bi] * 1000
		f.samples = 1152
	} else {
		f.bitrate = bitratesV2[bi] * 1000
		f.samples = 576
	}
	f.sampleRate = rates[f.version][si]
	f.padding = (h[2]>>1)&1 == 1
	f.mono = h[3]>>6 == 3
	pad := 0
	if f.padding {
		pad = 1
	}
	f.size = f.samples/8*f.bitrate/f.sampleRate + pad
	switch {
	case f.version == 1 && f.mono:
		f.sideInfoLen = 17
	case f.version == 1:
		f.sideInfoLen = 32
	case f.mono:
		f.sideInfoLen = 9
	default:
		f.sideInfoLen = 17
	}
	return f, nil
}

// sameStream reports whether two frames can follow each other in one stream.
func sameStream(a, b frameInfo) bool {
	return a.version == b.version && a.sampleRate == b.sampleRate && a.mono == b.mono
}

// mp3Stream is the audio frames of one MP3 file, without ID3 tags and
// without a Xing, Info or VBRI header frame.
type mp3Stream struct {
	frames [][]byte
	info   frameInfo // of the first audio frame
}

// parseMP3 splits an MP3 file into its audio frames.
func parseMP3(b []byte) (*mp3Stream, error) {
	i := 0
	if len(b) >= 10 && string(b[:3]) == "ID3" {
		size := int(b[6]&0x7F)<<21 | int(b[7]&0x7F)<<14 | int(b[8]&0x7F)<<7 | int(b[9]&0x7F)
		i = 10 + size
		if b[5]&0x10 != 0 {
			i += 10
		}
	}
	s := &mp3Stream{}
	for i < len(b) {
		if len(b)-i >= 3 && string(b[i:i+3]) == "TAG" && len(b)-i == 128 {
			break // ID3v1 at the end
		}
		f, err := parseHeader(b[i:])
		if err != nil {
			return nil, fmt.Errorf("offset %d: %w", i, err)
		}
		if i+f.size > len(b) {
			return nil, fmt.Errorf("offset %d: truncated frame (%d of %d bytes)", i, len(b)-i, f.size)
		}
		fr := b[i : i+f.size]
		i += f.size
		if isInfoFrame(fr, f) {
			continue
		}
		if len(s.frames) == 0 {
			s.info = f
		} else if !sameStream(s.info, f) {
			return nil, fmt.Errorf("offset %d: frame format changes within the file", i-f.size)
		}
		s.frames = append(s.frames, fr)
	}
	if len(s.frames) == 0 {
		return nil, errors.New("no audio frames")
	}
	return s, nil
}

func isInfoFrame(fr []byte, f frameInfo) bool {
	off := 4 + f.sideInfoLen
	if f.crc {
		off += 2
	}
	if len(fr) >= off+4 {
		tag := string(fr[off : off+4])
		if tag == "Xing" || tag == "Info" {
			return true
		}
	}
	return len(fr) >= 40 && string(fr[36:40]) == "VBRI"
}

// frameDurationMs is the duration of one frame in ms (24 ms at 24 kHz MPEG-2).
func (f frameInfo) frameDurationMs() float64 {
	return float64(f.samples) * 1000 / float64(f.sampleRate)
}

// silenceFrame returns one frame of digital silence in the format of f: the
// same header without CRC or padding, then zeros (all-zero side information
// means no main data and main_data_begin 0, so it never borrows from the bit
// reservoir and nothing borrows from it).
func silenceFrame(f frameInfo, header []byte) []byte {
	g := f
	g.padding = false
	g.crc = false
	h := []byte{header[0], header[1] | 0x01, header[2] &^ 0x02, header[3]}
	size := g.samples / 8 * g.bitrate / g.sampleRate
	fr := make([]byte, size)
	copy(fr, h)
	return fr
}

// silenceFrames returns n silent frames matching s.
func (s *mp3Stream) silenceFrames(n int) [][]byte {
	out := make([][]byte, n)
	sf := silenceFrame(s.info, s.frames[0][:4])
	for i := range out {
		out[i] = sf
	}
	return out
}

// framesFor returns how many frames of s come closest to ms.
func (s *mp3Stream) framesFor(ms int64) int {
	if ms <= 0 {
		return 0
	}
	d := s.info.frameDurationMs()
	return int(float64(ms)/d + 0.5)
}

func (s *mp3Stream) bytes() []byte {
	var buf bytes.Buffer
	for _, f := range s.frames {
		buf.Write(f)
	}
	return buf.Bytes()
}

// joinMP3 concatenates parts with gapMs of silence between consecutive parts
// and padMs of silence at the end. All parts must share one format. It
// returns the file and the start offset of each part in ms, measured by
// decoding the frames before it.
func joinMP3(parts [][]byte, gapMs, padMs int64) ([]byte, []int64, error) {
	if len(parts) == 0 {
		return nil, nil, errors.New("nothing to join")
	}
	streams := make([]*mp3Stream, len(parts))
	for i, p := range parts {
		s, err := parseMP3(p)
		if err != nil {
			return nil, nil, fmt.Errorf("part %d: %w", i, err)
		}
		if i > 0 && !sameStream(streams[0].info, s.info) {
			return nil, nil, fmt.Errorf("part %d: format differs from part 0 (%d Hz, MPEG %d)", i, s.info.sampleRate, s.info.version)
		}
		streams[i] = s
	}
	out := &mp3Stream{info: streams[0].info}
	starts := make([]int64, len(parts))
	gap := streams[0].framesFor(gapMs)
	for i, s := range streams {
		if i > 0 {
			out.frames = append(out.frames, streams[0].silenceFrames(gap)...)
		}
		ms, err := durationMs(out.bytes())
		if err != nil {
			return nil, nil, err
		}
		starts[i] = ms
		out.frames = append(out.frames, s.frames...)
	}
	out.frames = append(out.frames, streams[0].silenceFrames(streams[0].framesFor(padMs))...)
	return out.bytes(), starts, nil
}

// durationMs decodes an MP3 completely and returns its duration in ms,
// rounded to the nearest ms. An empty input has duration 0.
func durationMs(b []byte) (int64, error) {
	if len(b) == 0 {
		return 0, nil
	}
	d, err := gomp3.NewDecoder(bytes.NewReader(b))
	if err != nil {
		return 0, fmt.Errorf("decode mp3: %w", err)
	}
	n := d.Length() // bytes of 16-bit stereo PCM
	if n < 0 {
		return 0, errors.New("decode mp3: unknown length")
	}
	samples := n / 4
	return (samples*1000 + int64(d.SampleRate())/2) / int64(d.SampleRate()), nil
}
