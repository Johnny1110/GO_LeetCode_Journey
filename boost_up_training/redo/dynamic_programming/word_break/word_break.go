package word_break

// Input: s = "leetcode", wordDict = ["leet","code"]
// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
func wordBreak(s string, wordDict []string) bool {

	// 1. make dict
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	// 2. make DP, start from ""
	dp := make([]bool, len(s)+1)
	dp[0] = true // idx-0 = "" -> always be true

	for idx := 1; idx <= len(s); idx++ {

		for x := 0; x <= idx; x++ {
			// iterate all true string's idx
			if dp[x] {
				target := s[x:idx]
				if dict[target] {
					dp[idx] = true
					break
				}
			}
		}

	}

	return dp[len(s)]
}
