package word_break

import (
	"strings"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true, // "leet" + "code"
		},
		{
			name:     "Example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true, // "apple" + "pen" + "apple"
		},
		{
			name:     "Example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false, // cannot be segmented
		},
		{
			name:     "single character match",
			s:        "a",
			wordDict: []string{"a"},
			expected: true,
		},
		{
			name:     "single character no match",
			s:        "a",
			wordDict: []string{"b"},
			expected: false,
		},
		{
			name:     "empty string",
			s:        "",
			wordDict: []string{"a", "b"},
			expected: true, // empty string can always be segmented
		},
		{
			name:     "empty dictionary",
			s:        "abc",
			wordDict: []string{},
			expected: false,
		},
		{
			name:     "exact match",
			s:        "hello",
			wordDict: []string{"hello"},
			expected: true,
		},
		{
			name:     "multiple ways to break",
			s:        "abcd",
			wordDict: []string{"a", "abc", "b", "cd"},
			expected: true, // "a" + "b" + "cd" or "abc" + "d" (if "d" was in dict)
		},
		{
			name:     "overlapping words",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			expected: true, // "aaaa" + "aaa"
		},
		{
			name:     "repeated pattern",
			s:        "abab",
			wordDict: []string{"ab"},
			expected: true, // "ab" + "ab"
		},
		{
			name:     "partial match insufficient",
			s:        "cars",
			wordDict: []string{"car", "ca", "rs"},
			expected: true, // "car" + "s" (if "s" in dict) or "ca" + "rs"
		},
		{
			name:     "long string with short words",
			s:        "goalspecial",
			wordDict: []string{"go", "goal", "goals", "special"},
			expected: true, // "goal" + "special"
		},
		{
			name:     "cannot break - missing pieces",
			s:        "abcdef",
			wordDict: []string{"abc", "def", "gh"},
			expected: true, // "abc" + "def"
		},
		{
			name:     "single letter dictionary",
			s:        "abc",
			wordDict: []string{"a", "b", "c"},
			expected: true, // "a" + "b" + "c"
		},
		{
			name:     "no valid segmentation",
			s:        "abcdef",
			wordDict: []string{"ab", "cde", "gh"},
			expected: false, // "ab" + "cde" leaves "f" unmatched
		},
		{
			name:     "word longer than string",
			s:        "cat",
			wordDict: []string{"cats"},
			expected: false,
		},
		{
			name:     "duplicate words in dictionary",
			s:        "abab",
			wordDict: []string{"ab", "ab", "ba"},
			expected: true, // "ab" + "ab" (duplicates don't matter)
		},
		{
			name:     "case sensitive",
			s:        "Leetcode",
			wordDict: []string{"leet", "code"},
			expected: false, // case matters
		},
		{
			name:     "concatenated word equals input",
			s:        "helloworld",
			wordDict: []string{"hello", "world"},
			expected: true, // "hello" + "world"
		},
		{
			name:     "complex overlapping",
			s:        "abcdefg",
			wordDict: []string{"ab", "abc", "cd", "def", "abcd"},
			expected: false, // "abcd" + "ef" + "g" but "ef" and "g" not in dict
		},
		{
			name:     "tricky segmentation",
			s:        "aaab",
			wordDict: []string{"a", "aa"},
			expected: false, // "aa" + "a" + "b" but "b" not in dict
		},
		{
			name:     "all single chars",
			s:        "ab",
			wordDict: []string{"a", "b"},
			expected: true, // "a" + "b"
		},
		{
			name:     "recursive pattern",
			s:        "ababa",
			wordDict: []string{"ab", "ba", "a"},
			expected: true, // "ab" + "a" + "ba" or other combinations
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordBreak(tt.s, tt.wordDict)
			if result != tt.expected {
				t.Errorf("wordBreak(%q, %v) = %v, expected %v", tt.s, tt.wordDict, result, tt.expected)
			}
		})
	}
}

func TestWordBreakEdgeCases(t *testing.T) {
	t.Run("very long string with simple pattern", func(t *testing.T) {
		s := strings.Repeat("a", 100)
		wordDict := []string{"a"}
		expected := true
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(100 'a's, [\"a\"]) = %v, expected %v", result, expected)
		}
	})

	t.Run("long string impossible to break", func(t *testing.T) {
		s := strings.Repeat("a", 50) + "b"
		wordDict := []string{"a"}
		expected := false // last 'b' cannot be matched
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(50 'a's + 'b', [\"a\"]) = %v, expected %v", result, expected)
		}
	})

	t.Run("alternating pattern", func(t *testing.T) {
		s := "ababababab"
		wordDict := []string{"ab", "ba"}
		expected := true // "ab" + "ab" + "ab" + "ab" + "ab"
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})

	t.Run("dictionary with empty string", func(t *testing.T) {
		s := "abc"
		wordDict := []string{"", "abc"}
		expected := true // "abc" matches directly
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})

	t.Run("max length constraints", func(t *testing.T) {
		s := strings.Repeat("x", 300) // test with longer string
		wordDict := []string{"x", "xx", "xxx"}
		expected := true
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(300 'x's with dict [\"x\", \"xx\", \"xxx\"]) = %v, expected %v", result, expected)
		}
	})
}

func TestWordBreakComplexScenarios(t *testing.T) {
	t.Run("greedy approach fails", func(t *testing.T) {
		// This test case where greedy (always pick longest match) fails
		s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
		wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
		expected := false // cannot match the final 'b'
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(long string ending with 'b') = %v, expected %v", result, expected)
		}
	})

	t.Run("multiple valid segmentations", func(t *testing.T) {
		s := "pineapplepenapple"
		wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
		expected := true // "pineapple" + "pen" + "apple" or "pine" + "apple" + "pen" + "apple"
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})

	t.Run("nested word segments", func(t *testing.T) {
		s := "catsanddog"
		wordDict := []string{"cat", "cats", "and", "sand", "dog"}
		expected := true // "cats" + "and" + "dog"
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})

	t.Run("requires backtracking", func(t *testing.T) {
		s := "abcd"
		wordDict := []string{"a", "abc", "b", "cd"}
		expected := true // need to try "a" + "b" + "cd" after "abc" + ? fails
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})

	t.Run("fibonacci-like word building", func(t *testing.T) {
		s := "abaaba"
		wordDict := []string{"a", "ab", "ba", "aba", "aba"}
		expected := true // multiple ways: "a" + "ba" + "aba" or "aba" + "aba"
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})
}

func TestWordBreakPerformance(t *testing.T) {
	t.Run("large dictionary", func(t *testing.T) {
		s := "thequickbrownfox"
		wordDict := make([]string, 0, 1000)
		// Add many irrelevant words
		for i := 0; i < 990; i++ {
			wordDict = append(wordDict, strings.Repeat("z", i%10+1))
		}
		// Add the actual words we need
		wordDict = append(wordDict, "the", "quick", "brown", "fox")

		expected := true
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak with large dictionary = %v, expected %v", result, expected)
		}
	})

	t.Run("repeated subproblems", func(t *testing.T) {
		// This should test memoization effectiveness
		s := strings.Repeat("a", 20) + strings.Repeat("b", 20)
		wordDict := []string{"a", "b", "aa", "bb", "aaa", "bbb"}
		expected := true
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(repeated pattern) = %v, expected %v", result, expected)
		}
	})
}

func TestWordBreakSpecialWords(t *testing.T) {
	t.Run("palindromes", func(t *testing.T) {
		s := "raceacar"
		wordDict := []string{"race", "a", "car"}
		expected := true // "race" + "a" + "car"
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})

	t.Run("numeric strings", func(t *testing.T) {
		s := "123456"
		wordDict := []string{"123", "456", "12", "34", "56"}
		expected := true // "123" + "456" or "12" + "34" + "56"
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})

	t.Run("special characters", func(t *testing.T) {
		s := "hello-world"
		wordDict := []string{"hello", "-", "world"}
		expected := true // "hello" + "-" + "world"
		result := wordBreak(s, wordDict)
		if result != expected {
			t.Errorf("wordBreak(%q, %v) = %v, expected %v", s, wordDict, result, expected)
		}
	})
}

func BenchmarkWordBreak(b *testing.B) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wordBreak(s, wordDict)
	}
}

func BenchmarkWordBreakComplex(b *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wordBreak(s, wordDict)
	}
}

func BenchmarkWordBreakLargeDictionary(b *testing.B) {
	s := "thequickbrownfox"
	wordDict := make([]string, 1000)
	// Fill with random words
	for i := 0; i < 996; i++ {
		wordDict[i] = strings.Repeat("z", i%10+1)
	}
	// Add target words
	wordDict[996] = "the"
	wordDict[997] = "quick"
	wordDict[998] = "brown"
	wordDict[999] = "fox"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wordBreak(s, wordDict)
	}
}

func BenchmarkWordBreakWorstCase(b *testing.B) {
	// Worst case: many ways to start but only one way to finish
	s := strings.Repeat("a", 50)
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wordBreak(s, wordDict)
	}
}
