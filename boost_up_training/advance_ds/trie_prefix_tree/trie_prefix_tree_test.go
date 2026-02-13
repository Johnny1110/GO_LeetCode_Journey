package trie_prefix_tree

import (
	"testing"
)

func TestTrieExample1(t *testing.T) {
	// LeetCode Example 1
	trie := Constructor()

	trie.Insert("apple")
	if !trie.Search("apple") {
		t.Error("Search(\"apple\") = false, expected true")
	}

	if trie.Search("app") {
		t.Error("Search(\"app\") = true, expected false")
	}

	if !trie.StartsWith("app") {
		t.Error("StartsWith(\"app\") = false, expected true")
	}

	trie.Insert("app")
	if !trie.Search("app") {
		t.Error("Search(\"app\") = false, expected true after inserting")
	}
}

func TestTrieBasicOperations(t *testing.T) {
	t.Run("Empty trie operations", func(t *testing.T) {
		trie := Constructor()

		if trie.Search("anything") {
			t.Error("Search on empty trie should return false")
		}

		if trie.StartsWith("anything") {
			t.Error("StartsWith on empty trie should return false")
		}
	})

	t.Run("Single character words", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("a")
		if !trie.Search("a") {
			t.Error("Search(\"a\") = false, expected true")
		}

		if !trie.StartsWith("a") {
			t.Error("StartsWith(\"a\") = false, expected true")
		}

		if trie.Search("ab") {
			t.Error("Search(\"ab\") = true, expected false")
		}
	})

	t.Run("Insert same word multiple times", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("test")
		trie.Insert("test")
		trie.Insert("test")

		if !trie.Search("test") {
			t.Error("Search(\"test\") = false, expected true")
		}

		if !trie.StartsWith("test") {
			t.Error("StartsWith(\"test\") = false, expected true")
		}
	})

	t.Run("Empty string operations", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("")
		if !trie.Search("") {
			t.Error("Search(\"\") = false, expected true after inserting empty string")
		}

		if !trie.StartsWith("") {
			t.Error("StartsWith(\"\") = false, expected true")
		}
	})
}

func TestTrieSearchOperations(t *testing.T) {
	t.Run("Search exact words", func(t *testing.T) {
		trie := Constructor()

		words := []string{"cat", "car", "card", "care", "careful"}
		for _, word := range words {
			trie.Insert(word)
		}

		// Test all inserted words
		for _, word := range words {
			if !trie.Search(word) {
				t.Errorf("Search(\"%s\") = false, expected true", word)
			}
		}

		// Test non-existent words
		nonWords := []string{"ca", "cars", "cards", "careless", "dog"}
		for _, word := range nonWords {
			if trie.Search(word) {
				t.Errorf("Search(\"%s\") = true, expected false", word)
			}
		}
	})

	t.Run("Search with overlapping words", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("abc")
		trie.Insert("abcd")
		trie.Insert("abcde")

		if !trie.Search("abc") {
			t.Error("Search(\"abc\") = false, expected true")
		}
		if !trie.Search("abcd") {
			t.Error("Search(\"abcd\") = false, expected true")
		}
		if !trie.Search("abcde") {
			t.Error("Search(\"abcde\") = false, expected true")
		}

		if trie.Search("ab") {
			t.Error("Search(\"ab\") = true, expected false")
		}
		if trie.Search("abcdef") {
			t.Error("Search(\"abcdef\") = true, expected false")
		}
	})
}

func TestTrieStartsWithOperations(t *testing.T) {
	t.Run("StartsWith basic functionality", func(t *testing.T) {
		trie := Constructor()

		words := []string{"apple", "app", "apricot", "application"}
		for _, word := range words {
			trie.Insert(word)
		}

		// Test valid prefixes
		validPrefixes := []string{"app", "appl", "apple", "apr", "apric", "apricot"}
		for _, prefix := range validPrefixes {
			if !trie.StartsWith(prefix) {
				t.Errorf("StartsWith(\"%s\") = false, expected true", prefix)
			}
		}

		// Test invalid prefixes
		invalidPrefixes := []string{"b", "ap", "appz", "apricots", "xyz"}
		for _, prefix := range invalidPrefixes {
			if trie.StartsWith(prefix) {
				t.Errorf("StartsWith(\"%s\") = true, expected false", prefix)
			}
		}
	})

	t.Run("StartsWith vs Search differences", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("programming")

		// These should be valid prefixes but not complete words
		prefixes := []string{"p", "pr", "pro", "prog", "program"}
		for _, prefix := range prefixes {
			if !trie.StartsWith(prefix) {
				t.Errorf("StartsWith(\"%s\") = false, expected true", prefix)
			}
			if trie.Search(prefix) {
				t.Errorf("Search(\"%s\") = true, expected false", prefix)
			}
		}

		// This should be both a valid prefix and a complete word
		if !trie.StartsWith("programming") {
			t.Error("StartsWith(\"programming\") = false, expected true")
		}
		if !trie.Search("programming") {
			t.Error("Search(\"programming\") = false, expected true")
		}
	})
}

func TestTrieComplexScenarios(t *testing.T) {
	t.Run("Mixed operations sequence", func(t *testing.T) {
		trie := Constructor()

		operations := []struct {
			op     string
			word   string
			expect bool
		}{
			{"insert", "word", false},
			{"search", "word", true},
			{"search", "wor", false},
			{"startswith", "wor", true},
			{"insert", "wor", false},
			{"search", "wor", true},
			{"startswith", "wo", true},
			{"search", "wo", false},
			{"insert", "world", false},
			{"startswith", "wor", true},
			{"search", "world", true},
			{"search", "words", false},
		}

		for i, op := range operations {
			switch op.op {
			case "insert":
				trie.Insert(op.word)
			case "search":
				if got := trie.Search(op.word); got != op.expect {
					t.Errorf("Operation %d: Search(\"%s\") = %v, expected %v", i, op.word, got, op.expect)
				}
			case "startswith":
				if got := trie.StartsWith(op.word); got != op.expect {
					t.Errorf("Operation %d: StartsWith(\"%s\") = %v, expected %v", i, op.word, got, op.expect)
				}
			}
		}
	})

	t.Run("Dictionary-like operations", func(t *testing.T) {
		trie := Constructor()

		// Insert a dictionary of words
		dictionary := []string{
			"the", "a", "there", "answer", "any", "by", "bye", "their", "hero", "heroic",
		}

		for _, word := range dictionary {
			trie.Insert(word)
		}

		// Test all words exist
		for _, word := range dictionary {
			if !trie.Search(word) {
				t.Errorf("Search(\"%s\") = false, expected true", word)
			}
		}

		// Test prefix queries
		prefixTests := []struct {
			prefix string
			should bool
		}{
			{"th", true},
			{"the", true},
			{"ther", true},
			{"an", true},
			{"ans", true},
			{"b", true},
			{"by", true},
			{"h", true},
			{"he", true},
			{"her", true},
			{"hero", true},
			{"z", false},
			{"xyz", false},
		}

		for _, test := range prefixTests {
			if got := trie.StartsWith(test.prefix); got != test.should {
				t.Errorf("StartsWith(\"%s\") = %v, expected %v", test.prefix, got, test.should)
			}
		}
	})

	t.Run("Long words and deep trie", func(t *testing.T) {
		trie := Constructor()

		longWords := []string{
			"supercalifragilisticexpialidocious",
			"antidisestablishmentarianism",
			"pneumonoultramicroscopicsilicovolcanoconiosiss",
		}

		for _, word := range longWords {
			trie.Insert(word)
		}

		for _, word := range longWords {
			if !trie.Search(word) {
				t.Errorf("Search(\"%s\") = false, expected true", word)
			}
		}

		// Test prefixes of long words
		for _, word := range longWords {
			for i := 1; i <= len(word); i++ {
				prefix := word[:i]
				if !trie.StartsWith(prefix) {
					t.Errorf("StartsWith(\"%s\") = false, expected true", prefix)
				}
			}
		}
	})
}

func TestTrieEdgeCases(t *testing.T) {
	t.Run("Single character variations", func(t *testing.T) {
		trie := Constructor()

		// Insert all single characters
		for ch := 'a'; ch <= 'z'; ch++ {
			trie.Insert(string(ch))
		}

		// Test all single characters
		for ch := 'a'; ch <= 'z'; ch++ {
			if !trie.Search(string(ch)) {
				t.Errorf("Search(\"%c\") = false, expected true", ch)
			}
			if !trie.StartsWith(string(ch)) {
				t.Errorf("StartsWith(\"%c\") = false, expected true", ch)
			}
		}
	})

	t.Run("Nested word structures", func(t *testing.T) {
		trie := Constructor()

		// Create nested structure: a -> ab -> abc -> abcd
		words := []string{"a", "ab", "abc", "abcd"}
		for _, word := range words {
			trie.Insert(word)
		}

		// All should be searchable
		for _, word := range words {
			if !trie.Search(word) {
				t.Errorf("Search(\"%s\") = false, expected true", word)
			}
		}

		// All should be valid prefixes
		for _, prefix := range words {
			if !trie.StartsWith(prefix) {
				t.Errorf("StartsWith(\"%s\") = false, expected true", prefix)
			}
		}

		// Additional prefixes should work
		if !trie.StartsWith("") {
			t.Error("StartsWith(\"\") = false, expected true")
		}
	})

	t.Run("Case sensitivity", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("Test")
		trie.Insert("test")

		if !trie.Search("Test") {
			t.Error("Search(\"Test\") = false, expected true")
		}
		if !trie.Search("test") {
			t.Error("Search(\"test\") = false, expected true")
		}

		if trie.Search("TEST") {
			t.Error("Search(\"TEST\") = true, expected false")
		}
	})
}

func BenchmarkTrie(b *testing.B) {
	b.Run("Insert operations", func(b *testing.B) {
		trie := Constructor()
		words := make([]string, 1000)

		// Generate test words
		for i := 0; i < 1000; i++ {
			words[i] = generateWord(i%20 + 1) // Words of length 1-20
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			trie.Insert(words[i%1000])
		}
	})

	b.Run("Search operations", func(b *testing.B) {
		trie := Constructor()
		words := make([]string, 1000)

		// Generate and insert test words
		for i := 0; i < 1000; i++ {
			words[i] = generateWord(i%20 + 1)
			trie.Insert(words[i])
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			trie.Search(words[i%1000])
		}
	})

	b.Run("StartsWith operations", func(b *testing.B) {
		trie := Constructor()
		words := make([]string, 1000)
		prefixes := make([]string, 1000)

		// Generate test words and prefixes
		for i := 0; i < 1000; i++ {
			words[i] = generateWord(i%20 + 1)
			trie.Insert(words[i])

			// Create prefix (half the word length)
			wordLen := len(words[i])
			prefixLen := wordLen / 2
			if prefixLen == 0 {
				prefixLen = 1
			}
			prefixes[i] = words[i][:prefixLen]
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			trie.StartsWith(prefixes[i%1000])
		}
	})
}

// Helper function to generate test words
func generateWord(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = 'a' + byte(i%26)
	}
	return string(result)
}
