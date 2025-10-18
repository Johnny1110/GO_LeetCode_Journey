package house_robber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	// init dp
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		// 1. profit from rob 2 house behind and rob this house
		profitA := dp[i-2] + nums[i]
		// 2, profit from rob previous house and skip this house
		profitB := dp[i-1]
		dp[i] = max(profitA, profitB)
	}

	return dp[len(nums)-1]
}
