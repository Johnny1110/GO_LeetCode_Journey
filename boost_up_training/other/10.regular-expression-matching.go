/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start

func isMatch(s string, p string) bool {
	// define dp
	dp := make([][]bool, len(s)+1) // i: from s[0:i], j: from p[0:j] bool -> valid
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	dp[0][0] = true                // empty string empty pattern are true
	for j := 2; j <= len(p); j++ { // handle empty string with a*, a*b*, a*b*c*
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {

			char, pattern := s[i-1], p[j-1]

			if pattern == '.' || char == pattern {
				dp[i][j] = dp[i-1][j-1]
			} else if pattern == '*' {
				// 1. ALWAYS try the "Zero Match" case (skip the character and the asterisk)
				dp[i][j] = dp[i][j-2]
				// 2. Try the "Match One or More" case (if the preceding character matches)
				prePattern := p[j-2]
				if prePattern == char || prePattern == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}

		}
	}

	return dp[len(s)][len(p)]
}

// @lc code=end

