package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sort"
	"time"
)

// Admin talks to the admin API (8.2): gate key cookie, password login, and
// the X-Study-Admin header.
type Admin struct {
	g  *Gen
	hc *http.Client
}

func NewAdmin(g *Gen, gateKey, password string) (*Admin, error) {
	if gateKey == "" || password == "" {
		return nil, fmt.Errorf("gate key or password not set")
	}
	jar, _ := cookiejar.New(nil)
	a := &Admin{g: g, hc: &http.Client{Transport: g.transport, Jar: jar, Timeout: 120 * time.Second,
		CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}}
	resp, err := a.hc.Get(g.base + "/admin/?k=" + url.QueryEscape(gateKey))
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusFound {
		return nil, fmt.Errorf("gate key: HTTP %d", resp.StatusCode)
	}
	b, _ := json.Marshal(map[string]string{"password": password})
	if _, status, err := a.call("POST", "/api/admin/login", b); err != nil || status != 200 {
		return nil, fmt.Errorf("login: HTTP %d %v", status, err)
	}
	return a, nil
}

func (a *Admin) call(method, path string, body []byte) ([]byte, int, error) {
	req, _ := http.NewRequest(method, a.g.base+path, bytes.NewReader(body))
	req.Header.Set("X-Study-Admin", "1")
	req.Header.Set("Content-Type", "application/json")
	resp, err := a.hc.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	return b, resp.StatusCode, err
}

func (a *Admin) CreateRun(name, source, config string) (string, error) {
	b, _ := json.Marshal(map[string]string{"name": name, "source": source, "configVersion": config})
	out, status, err := a.call("POST", "/api/admin/runs", b)
	if err != nil || status != 201 {
		return "", fmt.Errorf("HTTP %d %s %v", status, out, err)
	}
	var r struct {
		Code string `json:"code"`
	}
	return r.Code, json.Unmarshal(out, &r)
}

type exportSession struct {
	SessionID     string   `json:"sessionId"`
	Group         string   `json:"group"`
	DenialSlide   string   `json:"denialSlide"`
	Status        string   `json:"status"`
	LastClientSeq int64    `json:"lastClientSeq"`
	LastEventNo   int64    `json:"lastEventNo"`
	Flags         []string `json:"flags"`
	Events        []struct {
		EventNo   int64  `json:"eventNo"`
		ClientSeq *int64 `json:"clientSeq"`
		Source    string `json:"source"`
		Valid     bool   `json:"valid"`
		Consumed  bool   `json:"consumed"`
		Reason    string `json:"reason"`
	} `json:"events"`
	Ratings []json.RawMessage `json:"ratings"`
}

// Check verifies through the admin API that the generator's sessions are
// Finished, that groups and denial slides are balanced, and that no event was
// lost or duplicated. It also times the admin read paths.
func (a *Admin) Check(ss []*Session) bool {
	fmt.Println("\nAdmin check")
	t0 := time.Now()
	out, status, err := a.call("GET", "/api/admin/runs", nil)
	if err != nil || status != 200 {
		fmt.Printf("  runs: HTTP %d %v\n", status, err)
		return false
	}
	tRuns := time.Since(t0)
	var runs []struct {
		ID   string `json:"id"`
		Code string `json:"code"`
	}
	_ = json.Unmarshal(out, &runs)
	runID := ""
	for _, r := range runs {
		if r.Code == a.g.runCode {
			runID = r.ID
		}
	}
	if runID == "" {
		fmt.Println("  run not found")
		return false
	}
	t0 = time.Now()
	_, status, _ = a.call("GET", "/api/admin/participants?runId="+runID, nil)
	tList := time.Since(t0)
	t0 = time.Now()
	csv, cstatus, _ := a.call("GET", "/api/admin/export?format=csv&runId="+runID, nil)
	tCSV := time.Since(t0)
	t0 = time.Now()
	out, jstatus, err := a.call("GET", "/api/admin/export?format=json&includeInvalid=true&runId="+runID, nil)
	tJSON := time.Since(t0)
	fmt.Printf("  admin timings: runs %s ms, participants %s ms (HTTP %d), export csv %s ms (HTTP %d, %d KB), export json %s ms (HTTP %d, %d KB)\n",
		ms(tRuns), ms(tList), status, ms(tCSV), cstatus, len(csv)/1024, ms(tJSON), jstatus, len(out)/1024)
	if err != nil || jstatus != 200 {
		return false
	}
	var data []exportSession
	if err := json.Unmarshal(out, &data); err != nil {
		fmt.Printf("  export json: %v\n", err)
		return false
	}
	byID := map[string]*exportSession{}
	for i := range data {
		byID[data[i].SessionID] = &data[i]
	}
	pass := true
	groups := map[string]int{}
	denial := map[string]map[string]int{}
	statuses := map[string]int{}
	lost, dup, gaps, missing, unallocated := 0, 0, 0, 0, 0
	for _, s := range ss {
		if s.ID == "" {
			missing++
			continue
		}
		d := byID[s.ID]
		if d == nil {
			missing++
			continue
		}
		statuses[d.Status]++
		if d.Group == "" {
			unallocated++
		} else {
			groups[d.Group]++
		}
		if d.Group != "" {
			if denial[d.Group] == nil {
				denial[d.Group] = map[string]int{}
			}
			denial[d.Group][d.DenialSlide]++
		}
		// Every consumed clientSeq 1..lastClientSeq exactly once; eventNo 1..lastEventNo exactly once.
		seqs := map[int64]int{}
		nos := map[int64]int{}
		accepted := 0
		for _, e := range d.Events {
			nos[e.EventNo]++
			if e.Source == "client" && e.Consumed && e.ClientSeq != nil {
				seqs[*e.ClientSeq]++
				if e.Valid {
					accepted++
				}
			}
		}
		for k := int64(1); k <= d.LastClientSeq; k++ {
			switch seqs[k] {
			case 0:
				lost++
			case 1:
			default:
				dup++
			}
		}
		for k := int64(1); k <= d.LastEventNo; k++ {
			if nos[k] != 1 {
				gaps++
			}
		}
		if int64(len(nos)) != d.LastEventNo {
			gaps++
		}
		if accepted != int(s.Events.Load()) {
			fmt.Printf("  session %s: generator saw %d accepted events, server stored %d valid consumed client events\n", s.ID, s.Events.Load(), accepted)
			pass = false
		}
	}
	fmt.Printf("  sessions found %d of %d; statuses %v\n", len(ss)-missing, len(ss), statuses)
	if missing > 0 || statuses["Finished"] != len(ss) {
		pass = false
		fmt.Println("  FAIL: not every session is Finished")
	}
	diff := func(m map[string]int) int {
		lo, hi := -1, 0
		for _, v := range m {
			if lo < 0 || v < lo {
				lo = v
			}
			if v > hi {
				hi = v
			}
		}
		return hi - max(lo, 0)
	}
	fmt.Printf("  groups %v (difference %d), not allocated %d\n", groups, diff(groups), unallocated)
	if diff(groups) > 1 {
		pass = false
		fmt.Println("  FAIL: group counts differ by more than one")
	}
	// A slide no session of a group got counts as zero for that group.
	for _, m := range denial {
		for slide := range m {
			for _, m2 := range denial {
				if _, ok := m2[slide]; !ok {
					m2[slide] = 0
				}
			}
		}
	}
	var gs []string
	for g := range denial {
		gs = append(gs, g)
	}
	sort.Strings(gs)
	for _, g := range gs {
		fmt.Printf("  denial slides in %s: %v (difference %d)\n", g, denial[g], diff(denial[g]))
		if diff(denial[g]) > 1 {
			pass = false
			fmt.Printf("  FAIL: denial slides in %s differ by more than one\n", g)
		}
	}
	fmt.Printf("  events: lost clientSeq %d, duplicated clientSeq %d, eventNo gaps or duplicates %d\n", lost, dup, gaps)
	if lost+dup+gaps > 0 {
		pass = false
	}
	return pass
}
