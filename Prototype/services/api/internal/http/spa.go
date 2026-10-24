package http

import (
	"bytes"
	"io/fs"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"
)

// hashedBundleName matches Vite's content-hashed file names, e.g.
// "index-NMdTJX92.js".
var hashedBundleName = regexp.MustCompile(`-[A-Za-z0-9_-]{8,}\.[a-z0-9]+$`)

const placeholderIndex = `<!doctype html><html lang="en"><head><meta charset="utf-8"><title>Study</title></head>` +
	`<body><p>The app has not been built yet. Run <code>make build-web</code>.</p></body></html>`

// spaHandler serves a built Vite app from fsys under prefix: existing files
// as they are, anything else as index.html (SPA fallback, 9.1). index.html is
// Cache-Control: no-cache; content-hashed bundle files are immutable (12.4).
func spaHandler(fsys fs.FS, prefix string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		rel := strings.TrimPrefix(r.URL.Path, prefix)
		if rel != "" && rel != "index.html" && fsys != nil {
			if b, err := fs.ReadFile(fsys, rel); err == nil {
				if hashedBundleName.MatchString(path.Base(rel)) {
					w.Header().Set("Cache-Control", "public, max-age=31536000, immutable")
				} else {
					w.Header().Set("Cache-Control", "no-cache")
				}
				http.ServeContent(w, r, path.Base(rel), time.Time{}, bytes.NewReader(b))
				return
			}
			// A missing bundle file is a 404, not the SPA page.
			if path.Ext(rel) == ".js" || path.Ext(rel) == ".css" || path.Ext(rel) == ".map" {
				http.NotFound(w, r)
				return
			}
		}
		b := []byte(placeholderIndex)
		if fsys != nil {
			if x, err := fs.ReadFile(fsys, "index.html"); err == nil {
				b = x
			}
		}
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeContent(w, r, "index.html", time.Time{}, bytes.NewReader(b))
	})
}
