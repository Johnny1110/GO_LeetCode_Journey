package edit_distance

import "fmt"

func minDistance(word1 string, word2 string) int {
	// 1. define 2D array:
	dp := make([][]int, len(word2)+1)
	for i := range dp {
		dp[i] = make([]int, len(word1)+1)
		dp[i][0] = i
		// init dp:
		if i == 0 {
			for j := range dp[i] {
				dp[i][j] = j
			}
		}
	}

	fmt.Println("after init dp:", dp)

	for i := 1; i < len(dp); i++ {
		w2Idx := i - 1

		for j := 1; j < len(dp[i]); j++ {
			w1Idx := j - 1

			if word1[w1Idx] == word2[w2Idx] {
				// both are matched:
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// both are not matched.

				// 1. try replace:
				replaceCnt := dp[i-1][j-1]
				// 2. try insert:
				insertCnt := dp[i][j-1]
				// 3. try delete:
				deleteCnt := dp[i-1][j]

				dp[i][j] = min(replaceCnt, insertCnt, deleteCnt) + 1
			}

		}
	}

	fmt.Println("dp:", dp)

	return dp[len(word1)][len(word2)]
}
