package coin_change

import "math"

func coinChange(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	// define DP
	dp := make([]int, amount+1)

	for amt := 1; amt <= amount; amt++ {

		betterCnt := math.MaxInt32
		for _, coin := range coins {
			remainAmt := amt - coin

			if remainAmt == 0 {
				betterCnt = 1
				break
			} else if remainAmt > 0 && dp[remainAmt] > 0 {
				betterCnt = min(betterCnt, dp[remainAmt]+1)
			}
		}

		if betterCnt == math.MaxInt32 {
			dp[amt] = -1
		} else {
			dp[amt] = betterCnt
		}
	}

	return dp[amount]
}
