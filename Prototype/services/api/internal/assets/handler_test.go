package assets

import "testing"

func TestPublicPrefixes(t *testing.T) {
	cases := map[string]bool{
		"audio/1.4.0/" + sha + ".mp3":    true,
		"video/1.4.0/" + sha + ".mp4":    true,
		"video/1.4.0/manifest.json":      false, // the manifest is never served
		"video/1.4.0/" + sha + ".webm":   false,
		"video/" + sha + ".mp4":          false, // must sit under a version
		"video/1.4.0/x/" + sha + ".mp4":  false,
		"audio/1.4.0/manifest.json":      false,
		"images/ciso.png":                true,
		"slides/v1.3/s2.png":             true,
		"consent/placeholder.pdf":        true,
		"configs/study-config.v1.4.json": false,
		"../video/1.4.0/" + sha + ".mp4": false,
	}
	for p, want := range cases {
		if got := Public(p); got != want {
			t.Errorf("Public(%q) = %v, want %v", p, got, want)
		}
	}
}

const sha = "71ac796629f071af50c166270376cf0715133299426c77d291c89d90c82646d3"
