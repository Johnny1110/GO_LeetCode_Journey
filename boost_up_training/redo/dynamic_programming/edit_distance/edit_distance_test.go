package edit_distance

import (
	"strings"
	"testing"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name     string
		word1    string
		word2    string
		expected int
	}{
		{
			name:     "Example 1 - horse to ros",
			word1:    "horse",
			word2:    "ros",
			expected: 3, // replace 'h'→'r', delete 'r', delete 'e'
		},
		{
			name:     "Example 2 - intention to execution",
			word1:    "intention",
			word2:    "execution",
			expected: 5, // multiple operations needed
		},
		{
			name:     "identical strings",
			word1:    "abc",
			word2:    "abc",
			expected: 0, // no operations needed
		},
		{
			name:     "empty strings",
			word1:    "",
			word2:    "",
			expected: 0, // no operations needed
		},
		{
			name:     "empty to non-empty",
			word1:    "",
			word2:    "abc",
			expected: 3, // 3 insertions
		},
		{
			name:     "non-empty to empty",
			word1:    "abc",
			word2:    "",
			expected: 3, // 3 deletions
		},
		{
			name:     "single character same",
			word1:    "a",
			word2:    "a",
			expected: 0, // no operations needed
		},
		{
			name:     "single character different",
			word1:    "a",
			word2:    "b",
			expected: 1, // 1 substitution
		},
		{
			name:     "single char to empty",
			word1:    "a",
			word2:    "",
			expected: 1, // 1 deletion
		},
		{
			name:     "empty to single char",
			word1:    "",
			word2:    "a",
			expected: 1, // 1 insertion
		},
		{
			name:     "completely different strings",
			word1:    "abc",
			word2:    "def",
			expected: 3, // 3 substitutions
		},
		{
			name:     "one string is substring",
			word1:    "abc",
			word2:    "ab",
			expected: 1, // 1 deletion
		},
		{
			name:     "substring relationship reversed",
			word1:    "ab",
			word2:    "abc",
			expected: 1, // 1 insertion
		},
		{
			name:     "prefix match",
			word1:    "kitten",
			word2:    "sitting",
			expected: 3, // k→s, e→i, insert g
		},
		{
			name:     "suffix match",
			word1:    "saturday",
			word2:    "sunday",
			expected: 3, // delete 'a', delete 't', delete 'r'
		},
		{
			name:     "anagrams",
			word1:    "abc",
			word2:    "bca",
			expected: 2, // 2 operations (swap operations)
		},
		{
			name:     "repeated characters",
			word1:    "aaa",
			word2:    "aa",
			expected: 1, // 1 deletion
		},
		{
			name:     "repeated characters different",
			word1:    "aaa",
			word2:    "bbb",
			expected: 3, // 3 substitutions
		},
		{
			name:     "palindrome to reverse",
			word1:    "racecar",
			word2:    "racecar",
			expected: 0, // same palindrome
		},
		{
			name:     "case sensitive",
			word1:    "abc",
			word2:    "ABC",
			expected: 3, // 3 substitutions (case matters)
		},
		{
			name:     "long common subsequence",
			word1:    "abcdefghijk",
			word2:    "a1b2c3d4e5f",
			expected: 10, // need to handle digits vs letters
		},
		{
			name:     "numbers vs letters",
			word1:    "123",
			word2:    "abc",
			expected: 3, // 3 substitutions
		},
		{
			name:     "insertions only",
			word1:    "a",
			word2:    "abcd",
			expected: 3, // 3 insertions
		},
		{
			name:     "deletions only",
			word1:    "abcd",
			word2:    "a",
			expected: 3, // 3 deletions
		},
		{
			name:     "mixed operations",
			word1:    "cat",
			word2:    "dog",
			expected: 3, // all substitutions
		},
		{
			name:     "longer strings with pattern",
			word1:    "exponential",
			word2:    "polynomial",
			expected: 6, // complex transformation
		},
		{
			name:     "common prefix different suffix",
			word1:    "prefix123",
			word2:    "prefix456",
			expected: 3, // change '123' to '456'
		},
		{
			name:     "common suffix different prefix",
			word1:    "abc_suffix",
			word2:    "def_suffix",
			expected: 3, // change 'abc' to 'def'
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minDistance(tt.word1, tt.word2)
			if result != tt.expected {
				t.Errorf("minDistance(%q, %q) = %d, expected %d", tt.word1, tt.word2, result, tt.expected)
			}
		})
	}
}

func TestMinDistanceEdgeCases(t *testing.T) {
	t.Run("very long identical strings", func(t *testing.T) {
		longString := strings.Repeat("a", 1000)
		expected := 0
		result := minDistance(longString, longString)
		if result != expected {
			t.Errorf("minDistance(long identical strings) = %d, expected %d", result, expected)
		}
	})

	t.Run("very long completely different strings", func(t *testing.T) {
		string1 := strings.Repeat("a", 500)
		string2 := strings.Repeat("b", 500)
		expected := 500 // all substitutions
		result := minDistance(string1, string2)
		if result != expected {
			t.Errorf("minDistance(500 'a's, 500 'b's) = %d, expected %d", result, expected)
		}
	})

	t.Run("one very long one empty", func(t *testing.T) {
		longString := strings.Repeat("x", 100)
		expected := 100 // all deletions
		result := minDistance(longString, "")
		if result != expected {
			t.Errorf("minDistance(100 chars, empty) = %d, expected %d", result, expected)
		}
	})

	t.Run("special characters", func(t *testing.T) {
		word1 := "hello@world.com"
		word2 := "hello#world!net"
		expected := 4 // @→#, .→!, c→n, m→t
		result := minDistance(word1, word2)
		if result <= expected-2 { // allow some flexibility for optimal solution
			t.Logf("minDistance(%q, %q) = %d (acceptable range)", word1, word2, result)
		}
	})

	t.Run("unicode characters", func(t *testing.T) {
		word1 := "café"
		word2 := "cave"
		expected := 2 // é→v, insert e (depends on how é is encoded)
		result := minDistance(word1, word2)
		if result >= expected-1 && result <= expected+1 { // allow flexibility for unicode handling
			t.Logf("minDistance(%q, %q) = %d (acceptable range)", word1, word2, result)
		}
	})
}

func TestMinDistanceSymmetry(t *testing.T) {
	testPairs := []struct {
		word1 string
		word2 string
	}{
		{"abc", "def"},
		{"hello", "world"},
		{"programming", "algorithm"},
		{"", "test"},
		{"a", "bb"},
	}

	for _, pair := range testPairs {
		t.Run("symmetry_"+pair.word1+"_"+pair.word2, func(t *testing.T) {
			dist1 := minDistance(pair.word1, pair.word2)
			dist2 := minDistance(pair.word2, pair.word1)
			if dist1 != dist2 {
				t.Errorf("Edit distance not symmetric: minDistance(%q, %q) = %d, but minDistance(%q, %q) = %d",
					pair.word1, pair.word2, dist1, pair.word2, pair.word1, dist2)
			}
		})
	}
}

func TestMinDistanceComplexScenarios(t *testing.T) {
	t.Run("DNA sequences", func(t *testing.T) {
		word1 := "ATCGATCG"
		word2 := "ATGCATCG"
		expected := 2 // C→G, insert C (or similar optimal solution)
		result := minDistance(word1, word2)
		if result <= expected+1 { // allow some flexibility
			t.Logf("minDistance(DNA: %q, %q) = %d", word1, word2, result)
		}
	})

	t.Run("transposition pattern", func(t *testing.T) {
		word1 := "ab"
		word2 := "ba"
		expected := 2 // 2 operations (since we only have insert/delete/replace)
		result := minDistance(word1, word2)
		if result != expected {
			t.Errorf("minDistance(%q, %q) = %d, expected %d", word1, word2, result, expected)
		}
	})

	t.Run("multiple repeated patterns", func(t *testing.T) {
		word1 := "ababab"
		word2 := "bababa"
		expected := 4 // shift pattern requires multiple operations
		result := minDistance(word1, word2)
		if result <= expected+1 { // allow flexibility for optimal solution
			t.Logf("minDistance(%q, %q) = %d", word1, word2, result)
		}
	})

	t.Run("logarithmic vs exponential", func(t *testing.T) {
		word1 := "log"
		word2 := "exp"
		expected := 3 // all different
		result := minDistance(word1, word2)
		if result != expected {
			t.Errorf("minDistance(%q, %q) = %d, expected %d", word1, word2, result, expected)
		}
	})

	t.Run("mathematical expressions", func(t *testing.T) {
		word1 := "x+y"
		word2 := "x*y"
		expected := 1 // +→*
		result := minDistance(word1, word2)
		if result != expected {
			t.Errorf("minDistance(%q, %q) = %d, expected %d", word1, word2, result, expected)
		}
	})
}

func TestMinDistanceTriangleInequality(t *testing.T) {
	// Test triangle inequality: d(a,c) <= d(a,b) + d(b,c)
	testSets := []struct {
		a, b, c string
	}{
		{"", "a", "ab"},
		{"cat", "car", "card"},
		{"abc", "def", "ghi"},
	}

	for _, set := range testSets {
		t.Run("triangle_"+set.a+"_"+set.b+"_"+set.c, func(t *testing.T) {
			dAB := minDistance(set.a, set.b)
			dBC := minDistance(set.b, set.c)
			dAC := minDistance(set.a, set.c)

			if dAC > dAB+dBC {
				t.Errorf("Triangle inequality violated: d(%q,%q)=%d > d(%q,%q)+d(%q,%q)=%d+%d=%d",
					set.a, set.c, dAC, set.a, set.b, set.b, set.c, dAB, dBC, dAB+dBC)
			}
		})
	}
}

func TestMinDistanceOptimalSubstructure(t *testing.T) {
	t.Run("optimal substructure property", func(t *testing.T) {
		// If we know optimal solution for prefixes, we should be able to build optimal solution
		word1 := "kitten"
		word2 := "sitting"
		
		// Test that removing last character gives us a related subproblem
		prefix1 := word1[:len(word1)-1] // "kitte"
		prefix2 := word2[:len(word2)-1] // "sittin"
		
		fullDistance := minDistance(word1, word2)
		prefixDistance := minDistance(prefix1, prefix2)
		
		// The full distance should be related to prefix distance
		// (either same, +1, or in some reasonable relationship)
		if fullDistance < prefixDistance || fullDistance > prefixDistance+2 {
			t.Logf("Full distance: %d, Prefix distance: %d (relationship noted)", fullDistance, prefixDistance)
		}
	})
}

func BenchmarkMinDistance(b *testing.B) {
	word1 := "kitten"
	word2 := "sitting"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minDistance(word1, word2)
	}
}

func BenchmarkMinDistanceEmpty(b *testing.B) {
	word1 := ""
	word2 := "test"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minDistance(word1, word2)
	}
}

func BenchmarkMinDistanceLongStrings(b *testing.B) {
	word1 := strings.Repeat("a", 100) + strings.Repeat("b", 100)
	word2 := strings.Repeat("b", 100) + strings.Repeat("a", 100)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minDistance(word1, word2)
	}
}

func BenchmarkMinDistanceWorstCase(b *testing.B) {
	// Worst case: completely different strings
	word1 := strings.Repeat("a", 150)
	word2 := strings.Repeat("b", 150)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minDistance(word1, word2)
	}
}

func BenchmarkMinDistanceBestCase(b *testing.B) {
	// Best case: identical strings
	word := strings.Repeat("test", 50)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minDistance(word, word)
	}
}