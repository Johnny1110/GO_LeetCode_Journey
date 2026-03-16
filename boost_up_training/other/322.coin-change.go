/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	// [1, 2] 2

	for amt := 1; amt <= amount; amt++ {

		tmp := math.MaxInt64

		for _, coin := range coins {

			remainAmt := amt - coin
			if remainAmt == 0 {
				tmp = 1
				break // early leave the coins try
			}

			if remainAmt > 0 && dp[remainAmt] > 0 {
				tmp = min(tmp, dp[remainAmt]+1)
			}
		}

		if tmp < math.MaxInt64 {
			dp[amt] = tmp
		}

	}

	result := dp[amount]
	if result == 0 {
		return -1
	} else {
		return result
	}
}

// @lc code=end

