package studyconfig

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// ContentDir is the repository's content directory, relative to this package.
const contentDir = "../../../../content"

func validator(t *testing.T) *Validator {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(contentDir, "schema", "study-config.schema.json"))
	if err != nil {
		t.Fatal(err)
	}
	v, err := NewValidator(b)
	if err != nil {
		t.Fatal(err)
	}
	return v
}

func stripNotes(v any) any {
	switch x := v.(type) {
	case map[string]any:
		for k, val := range x {
			if strings.HasPrefix(k, "_") {
				delete(x, k)
				continue
			}
			x[k] = stripNotes(val)
		}
	case []any:
		for i := range x {
			x[i] = stripNotes(x[i])
		}
	}
	return v
}

// TestRealConfigsRoundTrip loads the current configs against the schema and
// checks that the Go types round-trip them (author notes aside).
func TestRealConfigsRoundTrip(t *testing.T) {
	v := validator(t)
	for _, name := range []string{"study-config.v1.3.json", "study-config.v1.4.json"} {
		t.Run(name, func(t *testing.T) {
			data, err := os.ReadFile(filepath.Join(contentDir, name))
			if err != nil {
				t.Fatal(err)
			}
			c, err := Parse(data, v)
			if err != nil {
				t.Fatal(err)
			}
			out, err := json.Marshal(c)
			if err != nil {
				t.Fatal(err)
			}
			var a, b any
			_ = json.Unmarshal(data, &a)
			_ = json.Unmarshal(out, &b)
			if !reflect.DeepEqual(stripNotes(a), b) {
				t.Fatalf("round trip differs for %s", name)
			}
			if got := c.FlaggedSlideIDs(); !reflect.DeepEqual(got, []string{"s2", "s3", "s4"}) {
				t.Fatalf("flagged = %v", got)
			}
		})
	}
}

func TestSchemaRejectsUnknownKeys(t *testing.T) {
	v := validator(t)
	data, _ := os.ReadFile(filepath.Join(contentDir, "study-config.v1.4.json"))
	var m map[string]any
	_ = json.Unmarshal(data, &m)
	m["_anyNote"] = "ok"
	ok, _ := json.Marshal(m)
	if _, err := Parse(ok, v); err != nil {
		t.Fatalf("underscore key rejected: %v", err)
	}
	m["unknownKey"] = 1
	bad, _ := json.Marshal(m)
	if _, err := Parse(bad, v); err == nil {
		t.Fatal("unknown key accepted")
	}
}

func TestScenePackHidesGroupContent(t *testing.T) {
	data, _ := os.ReadFile(filepath.Join(contentDir, "study-config.v1.4.json"))
	c, err := Parse(data, nil)
	if err != nil {
		t.Fatal(err)
	}
	pack, err := PublicScenePack(c)
	if err != nil {
		t.Fatal(err)
	}
	s := string(pack)
	for _, bad := range []string{`"actions"`, `"answers"`, `"markings"`, `"routineQuestion"`, `"participantQuestions"`, `"correctOptionId"`, `"debrief"`} {
		if strings.Contains(s, bad) {
			t.Errorf("scene pack contains %s", bad)
		}
	}
	for _, line := range c.Setup.PreStartText.Checker {
		if strings.Contains(s, line) {
			t.Errorf("scene pack contains the checker pre-start sentence")
		}
	}
}
