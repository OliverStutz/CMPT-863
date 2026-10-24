// Package export writes the CSV and JSON exports of section 8.4.
package export

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
	"time"

	"study/api/internal/configs"
	"study/api/internal/domain"
	"study/api/internal/store"
	sc "study/api/pkg/studyconfig"
)

// Options of GET /api/admin/export.
type Options struct {
	RunID          string
	IncludeInvalid bool
	IncludeIP      bool
}

// SessionData is everything exported for one session.
type SessionData struct {
	Session *domain.Session
	Events  []*domain.Event
	Ratings map[string]*domain.Rating
	Answers map[string]*domain.AnswerDoc
	Content *domain.Content // nil when the version is not loaded
}

// Collect reads the sessions (oldest first) with their records.
func Collect(ctx context.Context, st store.Store, reg *configs.Registry, o Options) ([]*SessionData, error) {
	sessions, err := st.ListSessions(ctx, store.SessionFilter{RunID: o.RunID})
	if err != nil {
		return nil, err
	}
	sort.Slice(sessions, func(i, j int) bool { return sessions[i].OpenedAt.Before(sessions[j].OpenedAt) })
	var out []*SessionData
	for _, s := range sessions {
		d := &SessionData{Session: s, Ratings: map[string]*domain.Rating{}, Answers: map[string]*domain.AnswerDoc{}}
		if d.Events, err = st.Events(ctx, s.ID); err != nil {
			return nil, err
		}
		rts, err := st.Ratings(ctx, s.ID)
		if err != nil {
			return nil, err
		}
		for _, r := range rts {
			d.Ratings[r.Point] = r
		}
		ans, err := st.Answers(ctx, s.ID)
		if err != nil {
			return nil, err
		}
		for _, a := range ans {
			d.Answers[a.Kind] = a
		}
		if e, ok := reg.Get(s.ConfigVersion); ok {
			d.Content = e.Content
		}
		out = append(out, d)
	}
	return out, nil
}

var baseColumns = []string{"sessionId", "runCode", "source", "configVersion", "group", "denialSlide", "status",
	"openedAt", "consentAt", "startedAt", "finishedAt", "durationSec", "audioCheckCorrect", "audioCheckPlays",
	"reloadCount", "flags", "point", "slideId", "phase", "isDenialSlide", "trust", "competence", "benevolence",
	"integrity", "trustStatus", "competenceStatus", "benevolenceStatus", "integrityStatus", "ratingStatus",
	"lateSubmission", "latencyMs", "actionAtSlide", "actionTimedOut", "questionBy", "answerType", "budgetChoice",
	"budgetConfidence", "budgetStatus"}

var ratingItems = []string{"trust", "competence", "benevolence", "integrity"}

func ts(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.UTC().Format(time.RFC3339)
}

func str(v any) string {
	switch x := v.(type) {
	case nil:
		return ""
	case string:
		return x
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case int64:
		return strconv.FormatInt(x, 10)
	case int:
		return strconv.Itoa(x)
	case bool:
		return strconv.FormatBool(x)
	case map[string]any:
		if id, ok := x["optionId"]; ok {
			if t, ok := x["text"]; ok && str(t) != "" {
				return str(id) + ": " + str(t)
			}
			return str(id)
		}
	}
	b, _ := json.Marshal(v)
	return string(b)
}

// backgroundColumns is the union of background item ids (honeypots excluded)
// in first-seen order.
func backgroundColumns(data []*SessionData) []string {
	seen := map[string]bool{}
	var out []string
	for _, d := range data {
		if d.Content == nil {
			continue
		}
		bg := d.Content.Config.SceneOfKind(sc.SceneBackground)
		if bg == nil {
			continue
		}
		for _, it := range bg.Items {
			if it.Kind != sc.ItemHoneypot && !seen[it.ID] {
				seen[it.ID] = true
				out = append(out, it.ID)
			}
		}
	}
	return out
}

// WriteCSV writes one row per rating point (8.4). A session without a
// compiled plan (no consent yet) gets one row with empty rating fields.
func WriteCSV(w io.Writer, data []*SessionData, o Options) error {
	bgCols := backgroundColumns(data)
	header := append([]string{}, baseColumns...)
	for _, id := range bgCols {
		header = append(header, "bg_"+id)
	}
	header = append(header, "backgroundStatus")
	if o.IncludeIP {
		header = append(header, "ip")
	}
	cw := csv.NewWriter(w)
	if err := cw.Write(header); err != nil {
		return err
	}
	for _, d := range data {
		s := d.Session
		dur := ""
		if s.ConsentAt != nil && s.FinishedAt != nil {
			dur = strconv.FormatFloat(s.FinishedAt.Sub(*s.ConsentAt).Seconds(), 'f', 1, 64)
		}
		acCorrect, acPlays := "", ""
		if s.AudioCheck != nil {
			acPlays = strconv.Itoa(s.AudioCheck.Plays)
			if s.AudioCheck.Correct != nil {
				acCorrect = strconv.FormatBool(*s.AudioCheck.Correct)
			}
		}
		budget := d.Answers[domain.AnswersBudget]
		bChoice, bConf, bStatus := "", "", ""
		if budget != nil {
			bChoice, bConf, bStatus = str(budget.Values["optionId"]), str(budget.Values["confidence"]), budget.Status
		}
		bg := d.Answers[domain.AnswersBackground]
		var actions map[string]domain.ActionRecord
		if a := d.Answers[domain.AnswersActions]; a != nil {
			actions = a.Actions
		}
		session := []string{s.ID, s.RunCode, s.Source, s.ConfigVersion, s.Group, s.DenialSlide, s.Status,
			s.OpenedAt.UTC().Format(time.RFC3339), ts(s.ConsentAt), ts(s.StartedAt), ts(s.FinishedAt), dur,
			acCorrect, acPlays, strconv.Itoa(s.ReloadCount), strings.Join(s.Flags, ";")}
		tail := []string{bChoice, bConf, bStatus}
		for _, id := range bgCols {
			v := ""
			if bg != nil {
				v = str(bg.Values[id])
			}
			tail = append(tail, v)
		}
		bgStatus := ""
		if bg != nil {
			bgStatus = bg.Status
		}
		tail = append(tail, bgStatus)
		if o.IncludeIP {
			tail = append(tail, s.IP)
		}

		var points []domain.PlanStep
		if d.Content != nil && len(s.Plan) > len(domain.PreAllocationSteps) {
			points = d.Content.Plan.RatingPoints()
		}
		if len(points) == 0 {
			row := append(append([]string{}, session...), make([]string, len(baseColumns)-len(session)-3)...)
			if err := cw.Write(append(row, tail...)); err != nil {
				return err
			}
			continue
		}
		for _, p := range points {
			row := append([]string{}, session...)
			isDenial := ""
			if p.SlideID != "" {
				isDenial = strconv.FormatBool(p.SlideID == s.DenialSlide)
			}
			row = append(row, p.Point, p.SlideID, p.Phase, isDenial)
			rt := d.Ratings[p.Point]
			for _, it := range ratingItems {
				v := ""
				if rt != nil && rt.Values[it] != nil {
					v = strconv.FormatInt(*rt.Values[it], 10)
				}
				row = append(row, v)
			}
			for _, it := range ratingItems {
				v := ""
				if rt != nil {
					v = rt.ItemStatus[it]
				}
				row = append(row, v)
			}
			rStatus, late, lat := "", "", ""
			if rt != nil {
				rStatus = rt.Status
				late = strconv.FormatBool(rt.LateEventNo != nil)
				if rt.LatencyMs != nil {
					lat = strconv.FormatInt(*rt.LatencyMs, 10)
				}
			}
			row = append(row, rStatus, late, lat)
			act, timedOut, by, ans := "", "", "", ""
			if rec, ok := actions[p.SlideID]; ok {
				if rec.ActionID != nil {
					act = *rec.ActionID
				}
				timedOut, by, ans = strconv.FormatBool(rec.TimedOut), rec.QuestionBy, rec.AnswerType
			}
			row = append(row, act, timedOut, by, ans)
			if err := cw.Write(append(row, tail...)); err != nil {
				return err
			}
		}
	}
	cw.Flush()
	return cw.Error()
}

// WriteJSON writes one object per session with nested ratings, answers and
// the event log (invalid events only with IncludeInvalid; ip only with
// IncludeIP).
func WriteJSON(w io.Writer, data []*SessionData, o Options) error {
	out := make([]map[string]any, 0, len(data))
	for _, d := range data {
		var sess map[string]any
		b, err := json.Marshal(d.Session)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &sess); err != nil {
			return err
		}
		if !o.IncludeIP {
			delete(sess, "ip")
		}
		evs := []*domain.Event{}
		for _, e := range d.Events {
			if e.Valid || o.IncludeInvalid {
				evs = append(evs, e)
			}
		}
		ratings := []*domain.Rating{}
		if d.Content != nil {
			for _, p := range d.Content.Plan.RatingPoints() {
				if r := d.Ratings[p.Point]; r != nil {
					ratings = append(ratings, r)
				}
			}
		}
		sess["ratings"] = ratings
		sess["answers"] = d.Answers
		sess["events"] = evs
		out = append(out, sess)
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(out); err != nil {
		return fmt.Errorf("export json: %w", err)
	}
	return nil
}
