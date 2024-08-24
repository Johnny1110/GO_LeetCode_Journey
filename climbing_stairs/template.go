package climbing_stairs

import "go_leetcode_journey/common"

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1 // init only 1 stair
	dp[1] = 2 // init only 2 stair

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n-1]
}

// Go call this func in main.go
func Go() {
	common.Assert_answer(climbStairs(3), 3)
}
