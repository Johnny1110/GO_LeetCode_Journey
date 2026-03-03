package word_break

// Input: s = "leetcode", wordDict = ["leet","code"]
// Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
func wordBreak(s string, wordDict []string) bool {

	// make dict
	dict := make(map[string]bool)
	for _, word := range wordDict {
		dict[word] = true
	}

	// define dp
	dp := make([]bool, len(s)+1)
	dp[0] = true // is always valid as start idx

	for idx := 1; idx < len(dp); idx++ {

		for x := 0; x <= idx; x++ {

			if !dp[x] {
				continue
			}

			tmp := s[x:idx]
			if dict[tmp] {
				// matched
				dp[idx] = true
				break
			}

		}
	}
	return dp[len(s)]
}
