package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestLintJSON_Valid(t *testing.T) {
	input := `{
		"name": "ChatGPT",
		"language": "Go",
		"version": 4
	}`
	var output bytes.Buffer
	err := LintJSON(strings.NewReader(input), &output)
	if err != nil {
		t.Errorf("Expected valid JSON, got error: %v", err)
	}
	if !strings.Contains(output.String(), `"name": "ChatGPT"`) {
		t.Errorf("Output JSON missing expected content: %s", output.String())
	}
}

func TestLintJSON_Invalid(t *testing.T) {
	input := `{
"name": "ChatGPT",
"skills": ["Go", "Python"
}`

	var output bytes.Buffer
	err := LintJSON(strings.NewReader(input), &output)
	if err == nil {
		t.Error("Expected error for malformed JSON, got nil")
	} else if !strings.Contains(err.Error(), "line 1") {
		t.Errorf("Expected error to mention line 1 got: %v", err)
	}
}

func TestLintJSON_Empty(t *testing.T) {
	var output bytes.Buffer
	err := LintJSON(strings.NewReader(""), &output)
	if err == nil {
		t.Error("Expected error for empty input, got nil")
	}
}

