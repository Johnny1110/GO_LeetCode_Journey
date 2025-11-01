package word_break

func wordBreak(s string, wordDict []string) bool {
	// dump wordDict into a set, optimize finding word.
	wordDictSet := make(map[string]bool)
	for _, word := range wordDict {
		wordDictSet[word] = true
	}

	// define DP
	dp := make([]bool, len(s)+1)
	dp[0] = true // empty string always be true

	// iterate through the s.
	for currentDPIdx := 1; currentDPIdx < len(dp); currentDPIdx++ {

		for j := 0; j <= currentDPIdx; j++ {
			if dp[j] {
				targetWord := s[j:currentDPIdx]
				if _, ok := wordDictSet[targetWord]; ok {
					dp[currentDPIdx] = true
					break
				}
			}
		}

	}

	return dp[len(s)]
}
