package alien_dictionary

import "testing"

func TestForeignDictionary(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected string
	}{
		{
			name:     "basic example",
			words:    []string{"wrt", "wrf", "er", "ett", "rftt"},
			expected: "wertf",
		},
		{
			name:     "simple ordering",
			words:    []string{"z", "x"},
			expected: "zx",
		},
		{
			name:     "no valid ordering",
			words:    []string{"z", "x", "z"},
			expected: "",
		},
		{
			name:     "single word",
			words:    []string{"abc"},
			expected: "",
		},
		{
			name:     "empty input",
			words:    []string{},
			expected: "",
		},
		{
			name:     "prefix contradiction",
			words:    []string{"abc", "ab"},
			expected: "",
		},
		{
			name:     "complex ordering",
			words:    []string{"ac", "ab", "zc", "zb"},
			expected: "acbz",
		},
		{
			name:     "same length words",
			words:    []string{"ab", "ac"},
			expected: "abc",
		},
		{
			name:     "cycle detection",
			words:    []string{"a", "b", "a"},
			expected: "",
		},
		{
			name:     "multiple valid orderings",
			words:    []string{"z", "x", "y"},
			expected: "zxy",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := foreignDictionary(tt.words)
			if result != tt.expected {
				t.Errorf("foreignDictionary(%v) = %q, want %q", tt.words, result, tt.expected)
			}
		})
	}
}

func TestForeignDictionaryEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected string
	}{
		{
			name:     "identical words",
			words:    []string{"abc", "abc"},
			expected: "",
		},
		{
			name:     "single character words",
			words:    []string{"a", "b", "c"},
			expected: "abc",
		},
		{
			name:     "reverse alphabet",
			words:    []string{"z", "y", "x"},
			expected: "zyx",
		},
		{
			name:     "no ordering info",
			words:    []string{"abc", "def"},
			expected: "",
		},
		{
			name:     "valid prefix ordering",
			words:    []string{"ab", "abc"},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := foreignDictionary(tt.words)
			if result != tt.expected {
				t.Errorf("foreignDictionary(%v) = %q, want %q", tt.words, result, tt.expected)
			}
		})
	}
}
