package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	err := LintJSON(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// LintJSON reads from input, parses JSON, writes formatted JSON to output,
// and returns error with line/column info if parsing fails.
func LintJSON(input io.Reader, output io.Writer) error {
	reader := bufio.NewReader(input)
	var data map[string]interface{}

	var inputBuffer []byte
	tee := io.TeeReader(reader, &bufferWriter{&inputBuffer})

	decoder := json.NewDecoder(tee)
	err := decoder.Decode(&data)
	if err != nil {
		offset := decoder.InputOffset()
		line, col := getLineAndColumn(inputBuffer, offset)
		return fmt.Errorf("JSON parse error at line %d, column %d: %v", line, col, err)
	}

	encoder := json.NewEncoder(output)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "   ")
	err = encoder.Encode(data)
	if err != nil {
		return fmt.Errorf("JSON encoding error: %v", err)
	}
	return nil
}

type bufferWriter struct {
	buf *[]byte
}

func (bw *bufferWriter) Write(p []byte) (int, error) {
	*bw.buf = append(*bw.buf, p...)
	return len(p), nil
}

func getLineAndColumn(data []byte, offset int64) (line, column int) {
	line = 1
	column = 1
	for i := int64(0); i < offset && i < int64(len(data)); i++ {
		if data[i] == '\n' {
			line++
			column = 1
		} else {
			column++
		}
	}
	return
}

