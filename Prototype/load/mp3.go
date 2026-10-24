package main

import (
	"encoding/binary"
	"errors"
)

// mp3DurationMs estimates an MP3's duration: the Xing/Info frame count when
// present (VBR), otherwise the constant bitrate of the first frame. The TTS
// tool writes MPEG-1 Layer III files, so only that table is needed; other
// layers fall back to an error and the caller relies on too_early/retryInMs.
func mp3DurationMs(b []byte) (int64, error) {
	off := 0
	if len(b) > 10 && string(b[:3]) == "ID3" {
		sz := int(b[6]&0x7f)<<21 | int(b[7]&0x7f)<<14 | int(b[8]&0x7f)<<7 | int(b[9]&0x7f)
		off = 10 + sz
	}
	for ; off+4 <= len(b); off++ {
		if b[off] == 0xff && b[off+1]&0xe0 == 0xe0 {
			break
		}
	}
	if off+4 > len(b) {
		return 0, errors.New("no frame sync")
	}
	h := b[off : off+4]
	version := (h[1] >> 3) & 3 // 3 = MPEG-1, 2 = MPEG-2, 0 = MPEG-2.5
	layer := (h[1] >> 1) & 3   // 1 = Layer III
	if layer != 1 {
		return 0, errors.New("not layer III")
	}
	brIdx, srIdx := h[2]>>4, (h[2]>>2)&3
	mono := h[3]>>6 == 3
	var br []int
	var sr []int
	samples := 1152
	if version == 3 {
		br = []int{0, 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320}
		sr = []int{44100, 48000, 32000}
	} else {
		br = []int{0, 8, 16, 24, 32, 40, 48, 56, 64, 80, 96, 112, 128, 144, 160}
		sr = []int{22050, 24000, 16000}
		if version == 0 {
			sr = []int{11025, 12000, 8000}
		}
		samples = 576
	}
	if int(brIdx) >= len(br) || brIdx == 0 || int(srIdx) >= len(sr) {
		return 0, errors.New("bad header")
	}
	rate := sr[srIdx]
	// Xing/Info header sits after the side information.
	side := 32
	if version == 3 && mono {
		side = 17
	} else if version != 3 {
		side = 17
		if mono {
			side = 9
		}
	}
	x := off + 4 + side
	if x+12 <= len(b) && (string(b[x:x+4]) == "Xing" || string(b[x:x+4]) == "Info") {
		flags := binary.BigEndian.Uint32(b[x+4 : x+8])
		if flags&1 == 1 {
			frames := binary.BigEndian.Uint32(b[x+8 : x+12])
			return int64(frames) * int64(samples) * 1000 / int64(rate), nil
		}
	}
	kbps := br[brIdx]
	return int64(len(b)-off) * 8 / int64(kbps), nil
}
