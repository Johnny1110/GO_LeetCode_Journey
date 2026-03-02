package unique_paths

// m |, n ---
func uniquePaths(m int, n int) int {
	if (m == 1 || n == 1) {
		// one stright way. only 1 option
		return 1
	}

	// 2 option, go down or go right.

	// define DP
	dp := make([]int, n)
	for idx := range len(dp) {
		dp[idx] = 1
	}


	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1] 
		}
	}

	return dp[n-1]
}
