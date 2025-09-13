package unique_paths

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	dp := make([]int, n)
	// init start position
	dp[0] = 1

	for i := 0; i < m; i++ {

		for j := 0; j < n; j++ {

			// first column always be 1
			if j == 0 {
				// check is barrier
				if obstacleGrid[i][j] == 1 {
					dp[j] = 0
				}
				continue
			}

			// other column
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] += dp[j-1]
			}
		}

	}

	return dp[n-1]
}
