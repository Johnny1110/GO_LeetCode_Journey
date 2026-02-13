package add_and_search_words

import (
	"testing"
)

func TestWordDictionaryExample1(t *testing.T) {
	// LeetCode Example 1
	dict := Constructor()

	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")

	if dict.Search("pad") {
		t.Error("Search(\"pad\") = true, expected false")
	}

	if !dict.Search("bad") {
		t.Error("Search(\"bad\") = false, expected true")
	}

	if !dict.Search(".ad") {
		t.Error("Search(\".ad\") = false, expected true")
	}

	if !dict.Search("b..") {
		t.Error("Search(\"b..\") = false, expected true")
	}
}

func TestWordDictionaryBasicOperations(t *testing.T) {
	t.Run("Empty dictionary operations", func(t *testing.T) {
		dict := Constructor()

		if dict.Search("anything") {
			t.Error("Search on empty dictionary should return false")
		}

		if dict.Search(".") {
			t.Error("Search(\".\") on empty dictionary should return false")
		}

		if dict.Search("...") {
			t.Error("Search(\"...\") on empty dictionary should return false")
		}
	})

	t.Run("Add and search exact words", func(t *testing.T) {
		dict := Constructor()

		words := []string{"hello", "world", "test", "code", "leetcode"}
		for _, word := range words {
			dict.AddWord(word)
		}

		// Test all added words
		for _, word := range words {
			if !dict.Search(word) {
				t.Errorf("Search(\"%s\") = false, expected true", word)
			}
		}

		// Test non-existent words
		nonWords := []string{"hell", "worlds", "testing", "codes"}
		for _, word := range nonWords {
			if dict.Search(word) {
				t.Errorf("Search(\"%s\") = true, expected false", word)
			}
		}
	})

	t.Run("Single character words", func(t *testing.T) {
		dict := Constructor()

		dict.AddWord("a")
		dict.AddWord("b")
		dict.AddWord("z")

		if !dict.Search("a") {
			t.Error("Search(\"a\") = false, expected true")
		}

		if !dict.Search(".") {
			t.Error("Search(\".\") = false, expected true (should match single chars)")
		}

		if dict.Search("ab") {
			t.Error("Search(\"ab\") = true, expected false")
		}
	})

	t.Run("Add same word multiple times", func(t *testing.T) {
		dict := Constructor()

		dict.AddWord("duplicate")
		dict.AddWord("duplicate")
		dict.AddWord("duplicate")

		if !dict.Search("duplicate") {
			t.Error("Search(\"duplicate\") = false, expected true")
		}

		if !dict.Search("........") { // 9 dots for "duplicate"
			t.Error("Search(\"........\") = false, expected true")
		}
	})
}

func TestWordDictionaryWildcardSearch(t *testing.T) {
	t.Run("Single dot wildcard", func(t *testing.T) {
		dict := Constructor()

		dict.AddWord("cat")
		dict.AddWord("car")
		dict.AddWord("dog")

		// Test single dot replacements
		if !dict.Search("c.t") {
			t.Error("Search(\"c.t\") = false, expected true (should match cat)")
		}

		if !dict.Search("c.r") {
			t.Error("Search(\"c.r\") = false, expected true (should match car)")
		}

		if !dict.Search("d.g") {
			t.Error("Search(\"d.g\") = false, expected true (should match dog)")
		}

		if dict.Search("c.d") {
			t.Error("Search(\"c.d\") = true, expected false")
		}
	})

	t.Run("Multiple dot wildcards", func(t *testing.T) {
		dict := Constructor()

		dict.AddWord("abc")
		dict.AddWord("def")
		dict.AddWord("ghi")

		// Test multiple dots
		if !dict.Search("...") {
			t.Error("Search(\"...\") = false, expected true (should match all 3-char words)")
		}

		if !dict.Search("a..") {
			t.Error("Search(\"a..\") = false, expected true (should match abc)")
		}

		if !dict.Search(".e.") {
			t.Error("Search(\".e.\") = false, expected true (should match def)")
		}

		if !dict.Search("..i") {
			t.Error("Search(\"..i\") = false, expected true (should match ghi)")
		}

		if dict.Search("....") {
			t.Error("Search(\"....\") = true, expected false (no 4-char words)")
		}
	})

	t.Run("All dots pattern", func(t *testing.T) {
		dict := Constructor()

		dict.AddWord("hello")
		dict.AddWord("world")
		dict.AddWord("test")

		// Test all dots matching word length
		if !dict.Search(".....") { // 5 dots for "hello", "world"
			t.Error("Search(\".....\") = false, expected true")
		}

		if !dict.Search("....") { // 4 dots for "test"
			t.Error("Search(\"....\") = false, expected true")
		}

		if dict.Search("......") { // 6 dots - no matches
			t.Error("Search(\"......\") = true, expected false")
		}
	})

	t.Run("Mixed wildcards and characters", func(t *testing.T) {
		dict := Constructor()

		dict.AddWord("apple")
		dict.AddWord("apply")
		dict.AddWord("apron")

		if !dict.Search("app..") {
			t.Error("Search(\"app..\") = false, expected true")
		}

		if !dict.Search("ap...") {
			t.Error("Search(\"ap...\") = false, expected true")
		}

		if !dict.Search("a....") {
			t.Error("Search(\"a....\") = false, expected true")
		}

		if !dict.Search(".pp..") {
			t.Error("Search(\".pp..\") = false, expected true")
		}

		if dict.Search("app...") {
			t.Error("Search(\"app...\") = true, expected false (wrong length)")
		}
	})
}

func TestWordDictionaryComplexScenarios(t *testing.T) {
	t.Run("Different word lengths", func(t *testing.T) {
		dict := Constructor()

		words := []string{"a", "ab", "abc", "abcd", "abcde"}
		for _, word := range words {
			dict.AddWord(word)
		}

		// Test exact searches
		for _, word := range words {
			if !dict.Search(word) {
				t.Errorf("Search(\"%s\") = false, expected true", word)
			}
		}

		// Test wildcard searches
		patterns := []struct {
			pattern string
			should  bool
		}{
			{".", true},     // matches "a"
			{"..", true},    // matches "ab"
			{"...", true},   // matches "abc"
			{"....", true},  // matches "abcd"
			{".....", true}, // matches "abcde"
			{"......", false}, // no 6-char words
			{"a", true},
			{"a.", true},    // matches "ab"
			{"a..", true},   // matches "abc"
			{"a...", true},  // matches "abcd"
			{"a....", true}, // matches "abcde"
			{"b", false},
			{"b.", false},
		}

		for _, test := range patterns {
			if got := dict.Search(test.pattern); got != test.should {
				t.Errorf("Search(\"%s\") = %v, expected %v", test.pattern, got, test.should)
			}
		}
	})

	t.Run("Large dictionary with patterns", func(t *testing.T) {
		dict := Constructor()

		// Add various words
		words := []string{
			"cat", "car", "card", "care", "careful",
			"bat", "bar", "bark", "barn",
			"rat", "ran", "rain", "rainbow",
		}

		for _, word := range words {
			dict.AddWord(word)
		}

		// Test complex patterns
		testCases := []struct {
			pattern string
			should  bool
		}{
			{"c..", true},     // cat, car
			{"c...", true},    // card, care
			{"c......", true}, // careful
			{".ar", true},     // car, bar
			{".a.", true},     // cat, car, bat, bar, rat, ran
			{"..r", true},     // car, bar
			{"....", true},    // card, care, bark, barn, rain
			{".....w", true},  // rainbow
			{"z..", false},    // no words starting with z
			{".......", true}, // careful, rainbow
			{"........", false}, // no 8-char words
		}

		for _, test := range testCases {
			if got := dict.Search(test.pattern); got != test.should {
				t.Errorf("Search(\"%s\") = %v, expected %v", test.pattern, got, test.should)
			}
		}
	})

	t.Run("Mixed operations sequence", func(t *testing.T) {
		dict := Constructor()

		operations := []struct {
			op      string
			word    string
			expect  bool
		}{
			{"add", "word", false},
			{"search", "word", true},
			{"search", "wor.", true},
			{"search", ".ord", true},
			{"search", "w..d", true},
			{"search", "....", true},
			{"search", "w.rd", true},
			{"search", "word.", false}, // wrong length
			{"add", "world", false},
			{"search", "wor..", true}, // matches both word and world
			{"search", "wo...", true}, // matches world
			{"search", ".o...", true}, // matches world
			{"add", "work", false},
			{"search", "wor.", true}, // matches word, work
			{"search", "w...", true}, // matches word, work
		}

		for i, op := range operations {
			switch op.op {
			case "add":
				dict.AddWord(op.word)
			case "search":
				if got := dict.Search(op.word); got != op.expect {
					t.Errorf("Operation %d: Search(\"%s\") = %v, expected %v", i, op.word, got, op.expect)
				}
			}
		}
	})
}

func TestWordDictionaryEdgeCases(t *testing.T) {
	t.Run("Empty string handling", func(t *testing.T) {
		dict := Constructor()

		dict.AddWord("")

		if !dict.Search("") {
			t.Error("Search(\"\") = false, expected true after adding empty string")
		}

		if dict.Search(".") {
			t.Error("Search(\".\") = true, expected false (empty string doesn't match dot)")
		}
	})

	t.Run("Only dots in dictionary", func(t *testing.T) {
		dict := Constructor()

		// Add words that look like dots (but they're not wildcards when added)
		dict.AddWord(".")
		dict.AddWord("..")
		dict.AddWord("...")

		if !dict.Search(".") {
			t.Error("Search(\".\") = false, expected true")
		}

		if !dict.Search("..") {
			t.Error("Search(\"..\") = false, expected true")
		}

		if !dict.Search("...") {
			t.Error("Search(\"...\") = false, expected true")
		}
	})

	t.Run("Very long words", func(t *testing.T) {
		dict := Constructor()

		longWord := "supercalifragilisticexpialidocious"
		dict.AddWord(longWord)

		if !dict.Search(longWord) {
			t.Error("Search long word failed")
		}

		// Create pattern with all dots
		allDots := ""
		for i := 0; i < len(longWord); i++ {
			allDots += "."
		}

		if !dict.Search(allDots) {
			t.Error("Search with all dots pattern failed")
		}

		// Create mixed pattern
		mixedPattern := "super....................docious"
		if !dict.Search(mixedPattern) {
			t.Error("Search with mixed pattern failed")
		}
	})

	t.Run("All single characters", func(t *testing.T) {
		dict := Constructor()

		// Add all lowercase letters
		for ch := 'a'; ch <= 'z'; ch++ {
			dict.AddWord(string(ch))
		}

		// Test individual characters
		for ch := 'a'; ch <= 'z'; ch++ {
			if !dict.Search(string(ch)) {
				t.Errorf("Search(\"%c\") = false, expected true", ch)
			}
		}

		// Test dot wildcard should match any single character
		if !dict.Search(".") {
			t.Error("Search(\".\") = false, expected true")
		}
	})
}

func BenchmarkWordDictionary(b *testing.B) {
	b.Run("AddWord operations", func(b *testing.B) {
		dict := Constructor()
		words := generateTestWords(1000)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dict.AddWord(words[i%1000])
		}
	})

	b.Run("Search exact words", func(b *testing.B) {
		dict := Constructor()
		words := generateTestWords(1000)

		// Pre-populate dictionary
		for _, word := range words {
			dict.AddWord(word)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dict.Search(words[i%1000])
		}
	})

	b.Run("Search with wildcards", func(b *testing.B) {
		dict := Constructor()
		words := generateTestWords(1000)
		patterns := make([]string, 1000)

		// Pre-populate dictionary
		for i, word := range words {
			dict.AddWord(word)
			// Create pattern with some dots
			pattern := ""
			for j, ch := range word {
				if j%3 == 0 { // Every 3rd character becomes a dot
					pattern += "."
				} else {
					pattern += string(ch)
				}
			}
			patterns[i] = pattern
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dict.Search(patterns[i%1000])
		}
	})

	b.Run("Search all dots patterns", func(b *testing.B) {
		dict := Constructor()
		words := generateTestWords(1000)
		dotPatterns := make([]string, 1000)

		// Pre-populate dictionary
		for i, word := range words {
			dict.AddWord(word)
			// Create all-dots pattern
			pattern := ""
			for range word {
				pattern += "."
			}
			dotPatterns[i] = pattern
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dict.Search(dotPatterns[i%1000])
		}
	})
}

// Helper function to generate test words
func generateTestWords(count int) []string {
	words := make([]string, count)
	for i := 0; i < count; i++ {
		length := (i % 10) + 1 // Words of length 1-10
		word := ""
		for j := 0; j < length; j++ {
			word += string('a' + byte((i+j)%26))
		}
		words[i] = word
	}
	return words
}
