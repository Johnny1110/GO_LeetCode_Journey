package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	// define DP
	dp := make([]int, len(nums))
	finalMaxVal := 1

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		currVal := nums[i]

		for j := 0; j < i; j++ {
			if currVal > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])

				if dp[i] > finalMaxVal {
					finalMaxVal = dp[i]
				}
			}
		}
	}

	return finalMaxVal
}
