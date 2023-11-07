package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGrep(t *testing.T) {
	tests := []struct {
		name           string
		cfg            *GrepConfig
		input          string
		expectedOutput string
	}{
		{
			name: "Test Case Sensitive Match",
			cfg: &GrepConfig{
				pattern: "test",
			},
			input:          "This is a test\nAnother Test line\n",
			expectedOutput: "This is a test\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			if err := grep(tt.cfg, input, &output); err != nil {
				t.Fatalf("grep() error = %v", err)
			}

			if gotOutput := output.String(); gotOutput != tt.expectedOutput {
				t.Errorf("grep() = %v, want %v", gotOutput, tt.expectedOutput)
			}
		})
	}
}
