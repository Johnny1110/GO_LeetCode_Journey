package edit_distance

func minDistance(word1 string, word2 string) int {
	// Levenshtein distance
	// a metric measuring the minimum number of single-character edits—insertions, deletions, or substitutions—required to transform one string into another.

	if word1 == word2 {
		return 0
	} else if word1 == "" {
		return len(word2)
	} else if word2 == "" {
		return len(word1)
	}

	// define DP dp[i][j] -> edit distance for word1[0:i] to word2[0:j]
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i // delete i times
	}

	// init dp[0][j]
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}

	// strat transform
	// skip i = 0 & j = 0

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {

			if word1[i-1] == word2[j-1] {
				// 2 same chars found, no cost needed, ex: from "a" to "a", are same as "" to ""
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 2 chars are not the same.

				// insert: "a" -> "ab" based on "a" -> "a"
				insertCnt := dp[i][j-1] + 1
				// delete: "ab" -> "a" based on "a" -> "a"
				deleteCnt := dp[i-1][j] + 1
				// replace: "ab" -> "ac" based on "a" -> "a"
				replaceCnt := dp[i-1][j-1] + 1

				dp[i][j] = min(insertCnt, deleteCnt, replaceCnt)
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
