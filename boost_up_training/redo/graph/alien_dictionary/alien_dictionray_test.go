package alien_dictionary

import "testing"

func TestForeignDictionary(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected string
	}{
		{
			name:     "Example 1",
			words:    []string{"hrn", "hrf", "er", "enn", "rfnn"},
			expected: "hernf",
		},

		{
			name:     "Basic",
			words:    []string{"wrt", "wrf", "er", "ett", "rftt"},
			expected: "wertf",
		},
		{
			name:     "Simple",
			words:    []string{"z", "x"},
			expected: "zx",
		},

		{
			name:     "Single",
			words:    []string{"a"},
			expected: "a",
		},
		{
			name:     "Empty input",
			words:    []string{},
			expected: "",
		},
		{
			name:     "Multiple",
			words:    []string{"ac", "ab", "b"},
			expected: "acb", // Can be "acb" or "cab" - will validate in custom check
		},
		{
			name:     "Complex",
			words:    []string{"za", "zb", "ca", "cb"},
			expected: "zabc", // Should result in valid ordering like "azbc"
		},
		{
			name:     "Prefix violation - invalid",
			words:    []string{"abc", "ab"},
			expected: "",
		},
		{
			name:     "Valid prefix ordering",
			words:    []string{"ab", "abc"},
			expected: "abc",
		},
		{
			name:     "All same characters",
			words:    []string{"aa", "aa"},
			expected: "a",
		},
		{
			name:     "Long chains",
			words:    []string{"baa", "abcd", "abca", "cab", "cad"},
			expected: "", // Complex case - validate separately
		},
		{
			name:     "Single letter words",
			words:    []string{"a", "b", "c", "d"},
			expected: "abcd",
		},
		{
			name:     "Reverse alphabet",
			words:    []string{"d", "c", "b", "a"},
			expected: "dcba",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := foreignDictionary(tt.words)

			// For complex cases where multiple valid answers exist
			if tt.name == "Multiple valid orderings - ac before b" {
				// Validate that result is a valid topological sort
				if !isValidOrdering(tt.words, result) {
					t.Errorf("foreignDictionary(%v) = %q; want valid ordering", tt.words, result)
				}
				return
			}

			if tt.name == "Complex ordering" {
				// Should be a valid ordering containing all characters
				if !isValidOrdering(tt.words, result) || result == "" {
					t.Errorf("foreignDictionary(%v) = %q; want valid ordering", tt.words, result)
				}
				return
			}

			if tt.name == "Long chains" {
				// Should be a valid ordering or empty if impossible
				if result != "" && !isValidOrdering(tt.words, result) {
					t.Errorf("foreignDictionary(%v) = %q; want valid ordering or empty", tt.words, result)
				}
				return
			}

			if result != tt.expected {
				t.Errorf("foreignDictionary(%v) = %q; want %q", tt.words, result, tt.expected)
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
			name:     "Cycle detection",
			words:    []string{"a", "b", "a"},
			expected: "",
		},
		{
			name:     "Impossible ordering - prefix after longer word",
			words:    []string{"abcd", "ab"},
			expected: "",
		},
		{
			name:     "Same word repeated",
			words:    []string{"hello", "hello", "hello"},
			expected: "helo", // All unique characters in any valid order
		},
		{
			name:     "Two character ordering",
			words:    []string{"ba", "ab"},
			expected: "", // Impossible - contradictory ordering
		},
		{
			name:     "Disconnected components",
			words:    []string{"a", "c", "b", "d"},
			expected: "", // Multiple valid orderings possible
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := foreignDictionary(tt.words)

			// For cases where multiple valid answers exist or disconnected components
			if tt.name == "Same word repeated" {
				// Should contain all unique characters
				if !containsAllUniqueChars(tt.words, result) {
					t.Errorf("foreignDictionary(%v) = %q; want all unique characters", tt.words, result)
				}
				return
			}

			if tt.name == "Disconnected components" {
				// Should be a valid ordering containing all characters
				if !isValidOrdering(tt.words, result) || result == "" {
					t.Errorf("foreignDictionary(%v) = %q; want valid ordering", tt.words, result)
				}
				return
			}

			if result != tt.expected {
				t.Errorf("foreignDictionary(%v) = %q; want %q", tt.words, result, tt.expected)
			}
		})
	}
}

// Helper function to validate if a character ordering is valid for given words
func isValidOrdering(words []string, ordering string) bool {
	if len(words) == 0 {
		return ordering == ""
	}

	if ordering == "" {
		return false
	}

	// Create position map for characters
	pos := make(map[byte]int)
	for i, char := range []byte(ordering) {
		pos[char] = i
	}

	// Check if words are sorted according to this ordering
	for i := 0; i < len(words)-1; i++ {
		if !isLexicographicallySmaller(words[i], words[i+1], pos) {
			return false
		}
	}

	return true
}

// Helper function to check if word1 < word2 according to character positions
func isLexicographicallySmaller(word1, word2 string, pos map[byte]int) bool {
	minLen := len(word1)
	if len(word2) < minLen {
		minLen = len(word2)
	}

	for i := 0; i < minLen; i++ {
		pos1, exists1 := pos[word1[i]]
		pos2, exists2 := pos[word2[i]]

		// If character not in ordering, it's invalid
		if !exists1 || !exists2 {
			return false
		}

		if pos1 < pos2 {
			return true
		} else if pos1 > pos2 {
			return false
		}
	}

	// If all compared characters are equal, shorter word should come first
	return len(word1) <= len(word2)
}

// Helper function to check if result contains all unique characters from words
func containsAllUniqueChars(words []string, result string) bool {
	expected := make(map[byte]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			expected[word[i]] = true
		}
	}

	actual := make(map[byte]bool)
	for i := 0; i < len(result); i++ {
		actual[result[i]] = true
	}

	for char := range expected {
		if !actual[char] {
			return false
		}
	}

	return true
}

func BenchmarkForeignDictionary(b *testing.B) {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}

	for i := 0; i < b.N; i++ {
		foreignDictionary(words)
	}
}
