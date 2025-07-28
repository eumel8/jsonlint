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
	}` // Missing closing bracket

	var output bytes.Buffer
	err := LintJSON(strings.NewReader(input), &output)
	if err == nil {
		t.Error("Expected error for malformed JSON, got nil")
	} else if !strings.Contains(err.Error(), "line 4") {
		t.Errorf("Expected error to mention line number, got: %v", err)
	}
}

func TestLintJSON_Empty(t *testing.T) {
	var output bytes.Buffer
	err := LintJSON(strings.NewReader(""), &output)
	if err == nil {
		t.Error("Expected error for empty input, got nil")
	}
}

