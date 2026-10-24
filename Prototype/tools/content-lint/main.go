// Command content-lint validates a study config against the JSON Schema and
// applies the wording and structure rules of section 4.2 of
// "02 - Architecture.md". Exit code 1 on any failure, 0 otherwise (warnings
// included).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	sc "study/api/pkg/studyconfig"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("content-lint", flag.ContinueOnError)
	fs.SetOutput(stderr)
	schema := fs.String("schema", "", "JSON Schema (default <config dir>/schema/study-config.schema.json)")
	rulesPath := fs.String("rules", "", "rules file with the deny lists (default: the built-in rules.json)")
	manifest := fs.String("manifest", "", "audio manifest; enables the post-TTS checks (default <config dir>/audio/<version>/manifest.json if it exists)")
	noManifest := fs.Bool("no-manifest", false, "skip the post-TTS checks even if a manifest exists")
	resolved := fs.Bool("resolved", false, "the config is a resolved config (audioKey, durationMs and marks filled)")
	strict := fs.Bool("strict", false, "treat missing vlmRun as a failure (before a config is locked for a run)")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	if fs.NArg() == 0 {
		fmt.Fprintln(stderr, "usage: content-lint [flags] <config.json> ...")
		fs.PrintDefaults()
		return 2
	}
	rb := defaultRules
	if *rulesPath != "" {
		b, err := os.ReadFile(*rulesPath)
		if err != nil {
			fmt.Fprintln(stderr, err)
			return 2
		}
		rb = b
	}
	rules, err := parseRules(rb)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 2
	}
	code := 0
	for _, path := range fs.Args() {
		rep := lintFile(path, *schema, *manifest, *noManifest, rules, Options{Resolved: *resolved, Strict: *strict})
		for _, f := range rep.Failures {
			fmt.Fprintf(stdout, "%s: FAIL %s\n", path, f)
		}
		for _, w := range rep.Warnings {
			fmt.Fprintf(stdout, "%s: WARN %s\n", path, w)
		}
		fmt.Fprintf(stdout, "%s: %d failures, %d warnings\n", path, len(rep.Failures), len(rep.Warnings))
		if len(rep.Failures) > 0 {
			code = 1
		}
	}
	return code
}

func lintFile(path, schemaPath, manifestPath string, noManifest bool, rules *Rules, o Options) *Report {
	rep := &Report{}
	data, err := os.ReadFile(path)
	if err != nil {
		rep.fail("%v", err)
		return rep
	}
	if schemaPath == "" {
		schemaPath = filepath.Join(filepath.Dir(path), "schema", "study-config.schema.json")
	}
	sb, err := os.ReadFile(schemaPath)
	if err != nil {
		rep.fail("schema: %v", err)
		return rep
	}
	v, err := sc.NewValidator(sb)
	if err != nil {
		rep.fail("%v", err)
		return rep
	}
	if err := v.Validate(data); err != nil {
		rep.fail("%v", err)
		return rep
	}
	cfg, err := sc.Parse(data, nil)
	if err != nil {
		rep.fail("%v", err)
		return rep
	}
	if !noManifest {
		mp := manifestPath
		if mp == "" {
			mp = filepath.Join(filepath.Dir(path), "audio", cfg.Version, "manifest.json")
			if _, err := os.Stat(mp); err != nil {
				mp = ""
			}
		}
		if mp != "" {
			b, err := os.ReadFile(mp)
			if err != nil {
				rep.fail("manifest: %v", err)
				return rep
			}
			var m sc.Manifest
			if err := json.Unmarshal(b, &m); err != nil {
				rep.fail("manifest %s: %v", mp, err)
				return rep
			}
			o.Manifest = &m
		}
	}
	r := Lint(cfg, rules, o)
	r.Failures = append(rep.Failures, r.Failures...)
	return r
}
