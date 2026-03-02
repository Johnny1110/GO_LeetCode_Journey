package edit_distance

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}
	if word1 == "" {
		return len(word2)
	}
	if word2 == "" {
		return len(word1)
	}

	// 1. define DP as 2D array: [][]int -> dp[i][j], how many distance from word1[0:i] to word[0:j].
	dp := make([][]int, len(word1)+1)
	for idx := range len(dp) {
		dp[idx] = make([]int, len(word2)+1)
	}

	// 2. init DP - 1
	// Allowed operations: [replace, insert, delete]
	// "" -> ""  cost 0
	// "a" -> "" cost 1
	// "" -> "a" cost 1
	dp[0][0], dp[1][0], dp[0][1] = 0, 1, 1

	// 3. init DP - 2
	// "" -> "a", "ab", "abc" // from empty to K len words just insert K times.
	// [0][j] = j
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}

	// "a", "ab", "abc" -> "" // from K len words to empty just delete K times.
	// [i][0] = i
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}

	// 3. we start from i:1 j:1
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			// cost from word1[0:i] to word2[0:j]

			// 3-1. if chars matched. just no step needed.
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 3-2. if chars mismatched, choose 3 operations.
				// replace
				replaceCnt := dp[i-1][j-1] + 1
				// insert
				insertCnt := dp[i][j-1] + 1
				// delete
				deleteCnt := dp[i-1][j] + 1

				// bestApproach
				dp[i][j] = min(replaceCnt, insertCnt, deleteCnt)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
