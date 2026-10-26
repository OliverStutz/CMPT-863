package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Public reports whether a content path may be served to anyone. Only these
// prefixes are public: audio/<version>/*.mp3, video/<version>/*.mp4, images/*,
// slides/**/*.svg, slides/**/*.png and consent/*.pdf. Everything else under the content root
// (configs, the schema, vlm-drafts, manifests, notes, scripts) is never served.
func Public(p string) bool {
	cp := CleanPath(p)
	if cp == "" {
		return false
	}
	segs := strings.Split(cp, "/")
	ext := strings.ToLower(path.Ext(cp))
	switch segs[0] {
	case "audio":
		return len(segs) == 3 && ext == ".mp3"
	case "video":
		// Rendered talking-head clips (2026-10-25); the manifest stays private.
		return len(segs) == 3 && ext == ".mp4"
	case "images":
		if len(segs) < 2 {
			return false
		}
		switch ext {
		case ".png", ".jpg", ".jpeg", ".webp", ".svg", ".gif", ".avif":
			return true
		}
		return false
	case "slides":
		return len(segs) >= 2 && (ext == ".svg" || ext == ".png")
	case "consent":
		return len(segs) == 2 && ext == ".pdf"
	}
	return false
}

var hashedName = regexp.MustCompile(`^[0-9a-f]{64}\.[A-Za-z0-9]+$`)

// The container image has no /etc/mime.types, so the media types the study
// serves are registered here rather than sniffed.
func init() {
	_ = mime.AddExtensionType(".mp3", "audio/mpeg")
	_ = mime.AddExtensionType(".mp4", "video/mp4")
	_ = mime.AddExtensionType(".wav", "audio/wav")
}

// Generated serves content that the server builds itself: the public scene
// packs (scenes/<version>.json) and, for dev without TTS, silent clips.
type Generated interface {
	ScenePack(version string) ([]byte, bool)
}

// Handler serves /assets/<path> (and the /a/<path> alias of the document).
type Handler struct {
	Source    Source
	Generated Generated
	// Bundle is the embedded participant SPA. Vite writes its hashed bundle
	// files to assets/, so /assets/<file> is looked up there first.
	Bundle fs.FS
	Prefix string // "/assets/"
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		http.NotFound(w, r)
		return
	}
	rel := strings.TrimPrefix(r.URL.Path, h.Prefix)
	cp := CleanPath(rel)
	if cp == "" || rel != cp {
		http.NotFound(w, r)
		return
	}
	if h.Bundle != nil {
		if b, err := fs.ReadFile(h.Bundle, "assets/"+cp); err == nil {
			// Vite bundle names carry a content hash.
			w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			serveBytes(w, r, cp, b, time.Time{})
			return
		}
	}
	if strings.HasPrefix(cp, "scenes/") && strings.HasSuffix(cp, ".json") && h.Generated != nil {
		v := strings.TrimSuffix(strings.TrimPrefix(cp, "scenes/"), ".json")
		if b, ok := h.Generated.ScenePack(v); ok {
			sum := sha256.Sum256(b)
			w.Header().Set("ETag", `"`+hex.EncodeToString(sum[:16])+`"`)
			w.Header().Set("Cache-Control", "no-cache")
			serveBytes(w, r, cp, b, time.Time{})
			return
		}
		http.NotFound(w, r)
		return
	}
	if strings.HasPrefix(cp, "_dev/silence/") && strings.HasSuffix(cp, ".wav") {
		n, err := strconv.Atoi(strings.TrimSuffix(strings.TrimPrefix(cp, "_dev/silence/"), ".wav"))
		if err != nil || n < 0 || n > 600000 {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
		serveBytes(w, r, cp, SilenceWAV(n), time.Time{})
		return
	}
	if !Public(cp) {
		http.NotFound(w, r)
		return
	}
	obj, err := h.Source.Read(r.Context(), cp)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	if hashedName.MatchString(path.Base(cp)) {
		w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
	} else {
		w.Header().Set("Cache-Control", "no-cache")
	}
	serveBytes(w, r, cp, obj.Data, obj.ModTime)
}

// serveBytes serves with http.ServeContent (byte ranges, ETag and
// If-None-Match) and gzips JSON and HTML when the client accepts it and asks
// for the whole file.
func serveBytes(w http.ResponseWriter, r *http.Request, name string, b []byte, mod time.Time) {
	ct := mime.TypeByExtension(path.Ext(name))
	if ct == "" {
		ct = http.DetectContentType(b)
	}
	w.Header().Set("Content-Type", ct)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if compressible(ct) && r.Header.Get("Range") == "" && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)
		_, _ = zw.Write(b)
		_ = zw.Close()
		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Add("Vary", "Accept-Encoding")
		if et := w.Header().Get("ETag"); et != "" {
			w.Header().Set("ETag", strings.TrimSuffix(et, `"`)+`-gz"`)
		}
		http.ServeContent(w, r, "", mod, bytes.NewReader(buf.Bytes()))
		return
	}
	http.ServeContent(w, r, "", mod, bytes.NewReader(b))
}

func compressible(ct string) bool {
	return strings.HasPrefix(ct, "application/json") || strings.HasPrefix(ct, "text/html")
}

// SilenceWAV returns a silent 8 kHz 8-bit mono WAV of n milliseconds.
// DECISION: when the audio manifest is missing (dev without TTS), audio steps
// point to /assets/_dev/silence/<ms>.wav, generated here with the estimated
// duration, so the vertical slice runs end to end without TTS.
func SilenceWAV(n int) []byte {
	const rate = 8000
	samples := rate * n / 1000
	var b bytes.Buffer
	b.WriteString("RIFF")
	_ = binary.Write(&b, binary.LittleEndian, uint32(36+samples))
	b.WriteString("WAVEfmt ")
	_ = binary.Write(&b, binary.LittleEndian, uint32(16))
	_ = binary.Write(&b, binary.LittleEndian, uint16(1)) // PCM
	_ = binary.Write(&b, binary.LittleEndian, uint16(1)) // mono
	_ = binary.Write(&b, binary.LittleEndian, uint32(rate))
	_ = binary.Write(&b, binary.LittleEndian, uint32(rate))
	_ = binary.Write(&b, binary.LittleEndian, uint16(1))
	_ = binary.Write(&b, binary.LittleEndian, uint16(8))
	b.WriteString("data")
	_ = binary.Write(&b, binary.LittleEndian, uint32(samples))
	b.Write(bytes.Repeat([]byte{0x80}, samples))
	return b.Bytes()
}
