package main

import (
	"testing"
)

func TestSortLines(t *testing.T) {
	tests := []struct {
		name          string
		lines         []string
		columnIndex   int
		numericSort   bool
		reverse       bool
		unique        bool
		expectedLines []string
	}{
		{
			name:          "Simple string sort",
			lines:         []string{"c", "a", "b"},
			columnIndex:   0,
			numericSort:   false,
			reverse:       false,
			unique:        false,
			expectedLines: []string{"a", "b", "c"},
		},
		{
			name:          "Numeric sort",
			lines:         []string{"1", "10", "2"},
			columnIndex:   0,
			numericSort:   true,
			reverse:       false,
			unique:        false,
			expectedLines: []string{"1", "2", "10"},
		},
		{
			name:          "Reverse sort",
			lines:         []string{"a", "b", "c"},
			columnIndex:   0,
			numericSort:   false,
			reverse:       true,
			unique:        false,
			expectedLines: []string{"c", "b", "a"},
		},
		{
			name:          "Unique sort",
			lines:         []string{"a", "b", "b"},
			columnIndex:   0,
			numericSort:   false,
			reverse:       false,
			unique:        true,
			expectedLines: []string{"a", "b"},
		},
		{
			name:          "Column sort",
			lines:         []string{"a 3", "a 1", "a 2"},
			columnIndex:   1,
			numericSort:   false,
			reverse:       false,
			unique:        false,
			expectedLines: []string{"a 1", "a 2", "a 3"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sortedLines := sortLines(test.lines, test.columnIndex, test.numericSort, test.reverse, test.unique)
			for i, line := range sortedLines {
				if line != test.expectedLines[i] {
					t.Errorf("Test %s failed: line %d sorted incorrectly, got %s, want %s", test.name, i, line, test.expectedLines[i])
				}
			}
		})
	}
}
