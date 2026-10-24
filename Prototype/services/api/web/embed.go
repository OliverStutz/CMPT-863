// Package web embeds the two built SPAs: web/participant and web/admin, which
// `make build-web` fills from apps/participant/dist and apps/admin/dist.
//
// The pattern "*" (rather than naming the two directories) keeps the build
// working when a directory is missing, for example in a fresh clone where the
// git-ignored placeholders do not exist; the server then serves a built-in
// placeholder page.
package web

import (
	"embed"
	"io/fs"
)

//go:embed *
var files embed.FS

func sub(dir string) fs.FS {
	if _, err := fs.Stat(files, dir+"/index.html"); err != nil {
		return nil
	}
	s, err := fs.Sub(files, dir)
	if err != nil {
		return nil
	}
	return s
}

// Participant returns the participant bundle, or nil when not built.
func Participant() fs.FS { return sub("participant") }

// Admin returns the admin bundle, or nil when not built.
func Admin() fs.FS { return sub("admin") }
