package wildcard_matching

import (
	"fmt"
	"go_leetcode_journey/common"
)

func isMatch(s string, pattern string) bool {
	// init dp
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(pattern)+1)
	}

	// init dp[0][0] = True (empty s match empty pattern)
	dp[0][0] = true
	// init dp[i][0] = False

	// dp[0][j] depends on whether p[0:j] consists only of '*'.
	for j := 1; j <= len(pattern); j++ {
		if pattern[j-1] == '*' && dp[0][j-1] {
			dp[0][j] = true
		}
	}

	fmt.Println(dp)

	return false
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(isMatch("", "*"), true)
	common.Assert_answer(isMatch("", "***"), true)
	common.Assert_answer(isMatch("aa", "a"), false)
	common.Assert_answer(isMatch("a", "a"), true)
	common.Assert_answer(isMatch("aa", "a?"), true)
	common.Assert_answer(isMatch("aa", "a*"), true)
	common.Assert_answer(isMatch("aaaaab", "a*b"), true)
	common.Assert_answer(isMatch("aaaaabbbbb", "a*b"), false)
	common.Assert_answer(isMatch("aaaaabbbbb", "a*b*"), true)
	common.Assert_answer(isMatch("aabb", "a?b*"), true)
	common.Assert_answer(isMatch("aabb", "a?b"), false)
}
