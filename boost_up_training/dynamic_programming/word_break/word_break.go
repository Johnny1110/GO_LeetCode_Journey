package word_break

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}

	// 1. dump wordDict into a map
	wordSet := make(map[string]bool)
	wordLenSet := make(map[int]bool)
	for _, word := range wordDict {
		wordSet[word] = true
		wordLenSet[len(word)] = true
	}

	// 2. init dp: dp[i] -> from s[0...i] can be broken into valid words.
	dp := make([]bool, len(s)+1)
	dp[0] = true // empty string is valid as default.

	// example> s := "catsandog"

	for i := 0; i < len(s)+1; i++ { // `i` is input s's index

		for j := 0; j < i; j++ {

			if dp[j] && wordLenSet[i-j] {
				targetWord := s[j:i]
				if wordSet[targetWord] {
					dp[i] = true
				}
			}
		}

	}

	return dp[len(s)]
}
