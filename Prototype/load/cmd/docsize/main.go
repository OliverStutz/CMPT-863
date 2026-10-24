// Command docsize measures the Firestore documents a load run left in the
// emulator: document counts and sizes per collection, computed with
// Firestore's storage size rules (document name + fields + 32 bytes; string
// = UTF-8 length + 1; number, timestamp = 8; boolean, null = 1; map and array
// = sum of their parts). It uses only the emulator's REST API.
//
//	go run ./cmd/docsize -emulator localhost:8081 -project demo-study
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

type doc struct {
	Name   string                     `json:"name"`
	Fields map[string]json.RawMessage `json:"fields"`
}

func valueSize(raw json.RawMessage) int {
	var v map[string]json.RawMessage
	if json.Unmarshal(raw, &v) != nil {
		return 0
	}
	for k, x := range v {
		switch k {
		case "stringValue", "bytesValue", "referenceValue":
			var s string
			_ = json.Unmarshal(x, &s)
			return len(s) + 1
		case "integerValue", "doubleValue", "timestampValue", "geoPointValue":
			return 8
		case "booleanValue", "nullValue":
			return 1
		case "mapValue":
			var m struct {
				Fields map[string]json.RawMessage `json:"fields"`
			}
			_ = json.Unmarshal(x, &m)
			return fieldsSize(m.Fields)
		case "arrayValue":
			var a struct {
				Values []json.RawMessage `json:"values"`
			}
			_ = json.Unmarshal(x, &a)
			n := 0
			for _, e := range a.Values {
				n += valueSize(e)
			}
			return n
		}
	}
	return 0
}

func fieldsSize(f map[string]json.RawMessage) int {
	n := 0
	for k, v := range f {
		n += len(k) + 1 + valueSize(v)
	}
	return n
}

func docSize(d doc) int {
	// Name: each collection id and document id + 1, plus 16.
	_, rel, _ := strings.Cut(d.Name, "/documents/")
	n := 16
	for _, seg := range strings.Split(rel, "/") {
		n += len(seg) + 1
	}
	return n + fieldsSize(d.Fields) + 32
}

type agg struct {
	n, total, max int
	fields        map[string]int // per top-level field, summed
}

func main() {
	emu := flag.String("emulator", "localhost:8081", "Firestore emulator host:port")
	project := flag.String("project", "demo-study", "project id")
	maxSessions := flag.Int("sessions", 0, "only inspect the first n sessions (0 = all)")
	flag.Parse()
	base := fmt.Sprintf("http://%s/v1/projects/%s/databases/(default)/documents", *emu, *project)
	stats := map[string]*agg{}
	add := func(kind string, d doc) {
		a := stats[kind]
		if a == nil {
			a = &agg{fields: map[string]int{}}
			stats[kind] = a
		}
		s := docSize(d)
		a.n++
		a.total += s
		a.max = max(a.max, s)
		for k, v := range d.Fields {
			a.fields[k] += len(k) + 1 + valueSize(v)
		}
	}
	list := func(path string) []doc {
		var out []doc
		tok := ""
		for {
			u := base + "/" + path + "?pageSize=300"
			if tok != "" {
				u += "&pageToken=" + url.QueryEscape(tok)
			}
			resp, err := http.Get(u)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			var page struct {
				Documents     []doc  `json:"documents"`
				NextPageToken string `json:"nextPageToken"`
			}
			_ = json.NewDecoder(resp.Body).Decode(&page)
			resp.Body.Close()
			out = append(out, page.Documents...)
			if page.NextPageToken == "" {
				return out
			}
			tok = page.NextPageToken
		}
	}
	for _, c := range []string{"runs", "counters", "configs"} {
		for _, d := range list(c) {
			add(c, d)
		}
	}
	sessions := list("sessions")
	if *maxSessions > 0 && len(sessions) > *maxSessions {
		sessions = sessions[:*maxSessions]
	}
	for _, s := range sessions {
		add("sessions", s)
		_, rel, _ := strings.Cut(s.Name, "/documents/")
		for _, sub := range []string{"events", "ratings", "answers"} {
			for _, d := range list(rel + "/" + sub) {
				add("sessions/*/"+sub, d)
			}
		}
	}
	var kinds []string
	for k := range stats {
		kinds = append(kinds, k)
	}
	sort.Strings(kinds)
	fmt.Printf("| Collection | Documents | Avg bytes | Max bytes | Total KB | Per session |\n|---|---:|---:|---:|---:|---:|\n")
	for _, k := range kinds {
		a := stats[k]
		per := ""
		if len(sessions) > 0 && strings.HasPrefix(k, "sessions") {
			per = fmt.Sprintf("%.1f docs", float64(a.n)/float64(len(sessions)))
		}
		fmt.Printf("| %s | %d | %d | %d | %.1f | %s |\n", k, a.n, a.total/a.n, a.max, float64(a.total)/1024, per)
	}
	for _, k := range []string{"sessions", "sessions/*/events"} {
		a := stats[k]
		if a == nil {
			continue
		}
		type kv struct {
			k string
			v int
		}
		var fs []kv
		for f, v := range a.fields {
			fs = append(fs, kv{f, v / a.n})
		}
		sort.Slice(fs, func(i, j int) bool { return fs[i].v > fs[j].v })
		var parts []string
		for _, x := range fs[:min(6, len(fs))] {
			parts = append(parts, fmt.Sprintf("%s %d", x.k, x.v))
		}
		fmt.Printf("\nLargest fields of %s (avg bytes): %s\n", k, strings.Join(parts, ", "))
	}
}
