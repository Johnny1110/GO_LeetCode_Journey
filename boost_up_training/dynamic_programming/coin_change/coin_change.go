package coin_change

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}

	// 1. define the dp:
	dp := make([]int, amount+1)
	// 2. init dp:
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // init with a impossible number
	}

	// 3. state transform:
	for i := 1; i < len(dp); i++ {
		// i is current amount, for each every coin:
		for _, coin := range coins {
			remaining := i - coin
			if coin <= i {
				dp[i] = min(dp[remaining]+1, dp[i])
			}
		}
	}

	if dp[amount] >= amount+1 {
		return -1
	}
	return dp[amount]
}
