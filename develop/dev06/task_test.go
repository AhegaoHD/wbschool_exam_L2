package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseFields(t *testing.T) {
	testCases := []struct {
		name     string
		fields   string
		expected []int
		wantErr  bool
	}{
		{"Single field", "1", []int{0}, false},
		{"Multiple fields", "1,2,3", []int{0, 1, 2}, false},
		{"Invalid field", "a", nil, true},
		{"Empty field", "", []int{}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parseFields(tc.fields)
			if (err != nil) != tc.wantErr {
				t.Fatalf("parseFields(%q) error = %v, wantErr %v", tc.fields, err, tc.wantErr)
			}
			if !compareSlices(got, tc.expected) {
				t.Errorf("parseFields(%q) = %v, want %v", tc.fields, got, tc.expected)
			}
		})
	}
}

func TestSelectFields(t *testing.T) {
	columns := []string{"one", "two", "three", "four"}
	testCases := []struct {
		name     string
		fields   []int
		expected string
	}{
		{"Select first", []int{0}, "one"},
		{"Select first and third", []int{0, 2}, "one,three"},
		{"Select non-existent field", []int{10}, ""},
	}

	delimiter := "," // For test simplicity, we'll use a comma

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			selectedFields := selectFields(columns, tc.fields)
			got := strings.Join(selectedFields, delimiter)
			if got != tc.expected {
				t.Errorf("selectFields(%v) = %v, want %v", tc.fields, got, tc.expected)
			}
		})
	}
}

// Helper function to compare two slices
func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// TestCutCommand simulates the command execution with given input and flags
func TestCutCommand(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		fields     string
		delimiter  string
		onlySep    bool
		wantOutput string
	}{
		{"No flags", "one\ttwo\tthree", "", "\t", false, "one\ttwo\tthree\n"},
		{"Field flag", "one\ttwo\tthree", "1,3", "\t", false, "one\tthree\n"},
		{"Delimiter flag", "one,two,three", "2", ",", false, "two\n"},
		{"Only separated", "one\ttwo", "1", "\t", true, "one\n"},
		{"No separator in line", "one two", "1", "\t", true, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a buffer to capture output
			var output bytes.Buffer

			// Set up flags
			fieldsFlag := tc.fields
			delimiterFlag := tc.delimiter
			onlySeparatedFlag := tc.onlySep

			// Create an input reader
			input := strings.NewReader(tc.input)

			// Call the function that implements the cut command logic
			err := cutCommand(input, &output, fieldsFlag, delimiterFlag, onlySeparatedFlag)
			if err != nil {
				t.Errorf("cutCommand returned an error: %v", err)
			}

			// Compare the output
			if output.String() != tc.wantOutput {
				t.Errorf("cutCommand output = %q, want %q", output.String(), tc.wantOutput)
			}
		})
	}
}
