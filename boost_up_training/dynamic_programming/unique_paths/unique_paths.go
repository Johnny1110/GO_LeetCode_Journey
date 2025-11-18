package unique_paths

// uniquePaths
func uniquePaths(m int, n int) int {
	// 1. define 1D array as DP
	dp := make([]int, n)
	// 2. init dp (skip first, from idx: 1 to idx n-1)
	for i := 0; i < n; i++ {
		// first row always be 1 (only go right side, 1 option, no other choice)
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}

	return dp[n-1]
}
