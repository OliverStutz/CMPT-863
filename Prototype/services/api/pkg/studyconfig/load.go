package studyconfig

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Validator validates config JSON against study-config.schema.json.
type Validator struct {
	schema *jsonschema.Schema
}

const schemaURL = "https://example.invalid/study-config.schema.json"

// NewValidator compiles the JSON Schema (draft 2020-12).
func NewValidator(schemaJSON []byte) (*Validator, error) {
	doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(schemaJSON))
	if err != nil {
		return nil, fmt.Errorf("parse schema: %w", err)
	}
	c := jsonschema.NewCompiler()
	if err := c.AddResource(schemaURL, doc); err != nil {
		return nil, fmt.Errorf("add schema: %w", err)
	}
	sch, err := c.Compile(schemaURL)
	if err != nil {
		return nil, fmt.Errorf("compile schema: %w", err)
	}
	return &Validator{schema: sch}, nil
}

// Validate checks data against the schema.
func (v *Validator) Validate(data []byte) error {
	inst, err := jsonschema.UnmarshalJSON(bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("parse config: %w", err)
	}
	if err := v.schema.Validate(inst); err != nil {
		return fmt.Errorf("schema: %w", err)
	}
	return nil
}

var idPattern = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*$`)

// Parse decodes a config. When v is non-nil the JSON is first validated
// against the schema; with a nil validator only the Go-side structural checks
// in CheckIDs run (used when the schema file is not available).
func Parse(data []byte, v *Validator) (*StudyConfig, error) {
	if v != nil {
		if err := v.Validate(data); err != nil {
			return nil, err
		}
	}
	var c StudyConfig
	if err := json.Unmarshal(data, &c); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}
	if c.Version == "" {
		return nil, errors.New("config has no version")
	}
	if err := CheckIDs(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

// CheckIDs enforces section 4.1: every Id matches the pattern and is unique
// within the config. Ids of the rating items and of action options are fixed
// enum values and choices are scoped to their list, so they are not included.
func CheckIDs(c *StudyConfig) error {
	seen := map[string]string{}
	var errs []error
	add := func(id, where string) {
		if id == "" {
			errs = append(errs, fmt.Errorf("%s: empty id", where))
			return
		}
		if !idPattern.MatchString(id) {
			errs = append(errs, fmt.Errorf("%s: id %q does not match %s", where, id, idPattern))
		}
		if prev, ok := seen[id]; ok {
			errs = append(errs, fmt.Errorf("%s: id %q already used by %s", where, id, prev))
		}
		seen[id] = where
	}
	line := func(l *AudioLine, where string) {
		if l != nil {
			add(l.ID, where)
		}
	}
	blocks := func(bs []Block, where string) {
		for _, b := range bs {
			add(b.ID, where+" block")
			for _, sp := range b.Text.Spans() {
				add(sp.ID, where+" span")
			}
			for _, sp := range b.Value.Spans() {
				add(sp.ID, where+" span")
			}
			for _, sp := range b.Caption.Spans() {
				add(sp.ID, where+" span")
			}
			if b.Data != nil {
				for _, s := range b.Data.Series {
					add(s.ID, where+" series")
				}
				for _, m := range b.Data.Markers {
					add(m.ID, where+" marker")
				}
			}
		}
	}
	line(&c.Setup.AudioCheck.Clip, "setup.audioCheck.clip")
	blocks(c.Setup.ExampleScene.Layout, "setup.exampleScene")
	for i := range c.Setup.ExampleScene.Lines {
		line(&c.Setup.ExampleScene.Lines[i], "setup.exampleScene.lines")
	}
	for i := range c.Scenes {
		s := &c.Scenes[i]
		where := "scene " + s.ID
		add(s.ID, where)
		blocks(s.Layout, where)
		for j := range s.Lines {
			line(&s.Lines[j], where)
		}
		line(s.Narration, where)
		line(s.RoutineQuestion, where)
		if q := s.ParticipantQuestions; q != nil {
			line(&q.Explain, where)
			line(&q.Challenge, where)
		}
		if a := s.Answers; a != nil {
			line(a.Reason, where)
			line(a.Denial, where)
			line(a.Opening, where)
			line(a.ReasonBody, where)
			line(a.Closing, where)
		}
		if s.Checker != nil {
			for _, m := range s.Checker.Markings {
				add(m.ID, where+" marking")
			}
		}
		for _, it := range s.Items {
			add(it.ID, where+" item")
		}
	}
	return errors.Join(errs...)
}
