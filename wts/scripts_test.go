package wts_test

import (
	"path/filepath"
	"testing"

	"github.com/tchaudhry91/wts-service/wts"
)

var samples []string = []string{
	filepath.Join("samples", "list.sh"),
	filepath.Join("samples", "test.py"),
}

func buildSamples() ([]*wts.Script, error) {
	scripts := []*wts.Script{}
	// Files
	for _, fname := range samples {
		sb := wts.NewScriptBuilder(fname)
		sb.FromFile(fname)
		if filepath.Ext(fname) == "py" {
			sb.InterpretterPrefix("/usr/bin/env python")
		}
		s, err := sb.Build()
		if err != nil {
			return scripts, err
		}
		scripts = append(scripts, s)
	}

	// Inlines

	// Python Script - Inline
	{
		sb := wts.NewScriptBuilder("inline-python")
		sb.InterpretterPrefix("/usr/bin/python")
		sb.Inline("print('Works!')")
		s, err := sb.Build()
		if err != nil {
			return scripts, err
		}
		scripts = append(scripts, s)
	}

	// Echo Alias - Inline
	{
		sb := wts.NewScriptBuilder("inline-echo")
		sb.Inline("echo \"Works!\"")
		s, err := sb.Build()
		if err != nil {
			return scripts, err
		}
		scripts = append(scripts, s)
	}

	return scripts, nil
}

func TestBuilder(t *testing.T) {
	type TestCase struct {
		Name         string
		FName        string
		InlineB      bool
		Inline       string
		Interpretter string
		Fail         bool
	}
	cases := []TestCase{
		{"list-sh", samples[0], false, "", "", false},
		{"test-py", samples[1], false, "", "/usr/bin/python", false},
		{"non-existent", "abc.xyz", false, "", "", true},
		{"inline-good", "", true, "ps -aux", "", false},
		{"inline-bad", "", true, "", "", true},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			sb := wts.NewScriptBuilder(c.Name)
			if c.InlineB {
				sb.Inline(c.Inline)
			} else {
				sb.FromFile(c.FName)
			}
			sb.InterpretterPrefix(c.Interpretter)
			script, err := sb.Build()
			if !c.Fail && err != nil {
				t.Fatalf("Failed to build Script: %s", err.Error())
			}
			if c.Fail && err == nil {
				t.Fatalf("Failed to Error on faulty builder")
			}
			if !c.Fail {
				t.Logf("Script %s created with sum %s", script.Name, script.Checksum)
			}
		})
	}
}

func TestRun(t *testing.T) {
	samples, err := buildSamples()
	if err != nil {
		t.Fatalf("Failed to build samples: %s", err.Error())
	}

	for _, s := range samples {
		t.Run(s.Name, func(t *testing.T) {
			err = s.Run(false)
			if err != nil {
				t.Fatalf("Failed on script because: %s", err.Error())
			}
			// Re-Run Silent
			err = s.Run(true)
			if err != nil {
				t.Fatalf("Failed to run script silently because: %s", err.Error())
			}
		})
	}
}
