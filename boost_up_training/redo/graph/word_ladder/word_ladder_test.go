package word_ladder

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			name:      "Classic",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},

		{
			name:      "hotdog",
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "dog"},
			expected:  0,
		},

		{
			name:      "No solution",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
		{
			name:      "Single step transformation",
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			expected:  2,
		},
		{
			name:      "Same begin and end word",
			beginWord: "hit",
			endWord:   "hit",
			wordList:  []string{"hot", "hit"},
			expected:  1,
		},
		{
			name:      "Empty word list",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{},
			expected:  0,
		},
		{
			name:      "No possible transformation",
			beginWord: "hit",
			endWord:   "abc",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  0,
		},
		{
			name:      "Longer transformation chain",
			beginWord: "red",
			endWord:   "tax",
			wordList:  []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
			expected:  4,
		},
		{
			name:      "Single character words",
			beginWord: "a",
			endWord:   "z",
			wordList:  []string{"b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
			expected:  26,
		},
		{
			name:      "Complex graph with multiple paths",
			beginWord: "qa",
			endWord:   "sq",
			wordList:  []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"},
			expected:  5,
		},
		{
			name:      "Direct transformation impossible",
			beginWord: "dog",
			endWord:   "cat",
			wordList:  []string{"bat", "bag", "big", "dig", "fig"},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if result != tt.expected {
				t.Errorf("ladderLength(%q, %q, %v) = %d; want %d",
					tt.beginWord, tt.endWord, tt.wordList, result, tt.expected)
			}
		})
	}
}

func TestLadderLengthEdgeCases(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			name:      "BeginWord same as endWord",
			beginWord: "hit",
			endWord:   "hit",
			wordList:  []string{"hit"},
			expected:  1,
		},
		{
			name:      "Single word in list not matching endWord",
			beginWord: "hit",
			endWord:   "hot",
			wordList:  []string{"bit"},
			expected:  0,
		},
		{
			name:      "Two words with exact match",
			beginWord: "cat",
			endWord:   "bat",
			wordList:  []string{"bat"},
			expected:  2,
		},
		{
			name:      "Large word list with no solution",
			beginWord: "aaa",
			endWord:   "zzz",
			wordList:  []string{"bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh"},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if result != tt.expected {
				t.Errorf("ladderLength(%q, %q, %v) = %d; want %d",
					tt.beginWord, tt.endWord, tt.wordList, result, tt.expected)
			}
		})
	}
}

func BenchmarkLadderLength(b *testing.B) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	for i := 0; i < b.N; i++ {
		ladderLength(beginWord, endWord, wordList)
	}
}
