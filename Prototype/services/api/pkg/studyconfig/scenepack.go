package studyconfig

import "encoding/json"

// DECISION: the server builds the scene pack from the loaded config and serves
// it at /assets/scenes/<version>.json, instead of reading one written at
// publish, so dev and prod serve the same pack.
//
// PublicScenePack builds the public scene pack of section 4.3 from a config:
// only text that is the same for both groups and gives nothing away.
//
// Removed: actions (per-group options), the panel labels and colours, every
// slide's checker panel, routine
// question, participant questions and answers, setup.preStartText.checker,
// the debrief scene, the voices, the audio check's clip and correct answer
// (the clip text is the answer), and audio fields of every line (audio URLs
// come in the step descriptor). Author notes ("_" keys) are not in the Go
// types and so never reach the pack.
func PublicScenePack(c *StudyConfig) ([]byte, error) {
	raw, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	delete(m, "actions")
	delete(m, "voices")
	// The checker panel's texts and colours give the condition away; they
	// come with the overlay of a checker session's slide step instead.
	delete(m, "colors")
	if labels, ok := m["labels"].(map[string]any); ok {
		for _, k := range []string{"panelTitle", "panelIdle", "noIssues"} {
			delete(labels, k)
		}
	}
	if setup, ok := m["setup"].(map[string]any); ok {
		if pst, ok := setup["preStartText"].(map[string]any); ok {
			delete(pst, "checker")
		}
		if ac, ok := setup["audioCheck"].(map[string]any); ok {
			delete(ac, "clip")
			delete(ac, "correctOptionId")
		}
		if ex, ok := setup["exampleScene"].(map[string]any); ok {
			stripLines(ex["lines"])
		}
	}
	var scenes []any
	if list, ok := m["scenes"].([]any); ok {
		for _, s := range list {
			sm, ok := s.(map[string]any)
			if !ok {
				continue
			}
			switch sm["kind"] {
			case SceneDebrief:
				continue
			case SceneSlide:
				for _, k := range []string{"checker", "routineQuestion", "participantQuestions", "answers"} {
					delete(sm, k)
				}
				if n, ok := sm["narration"].(map[string]any); ok {
					stripLine(n)
				}
			case SceneOpening:
				stripLines(sm["lines"])
			}
			scenes = append(scenes, sm)
		}
	}
	m["scenes"] = scenes
	return json.Marshal(m)
}

func stripLines(v any) {
	if list, ok := v.([]any); ok {
		for _, l := range list {
			if lm, ok := l.(map[string]any); ok {
				stripLine(lm)
			}
		}
	}
}

func stripLine(l map[string]any) {
	for _, k := range []string{"audioKey", "durationMs", "marks", "ssml"} {
		delete(l, k)
	}
}
