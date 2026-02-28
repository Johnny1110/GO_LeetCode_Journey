package coin_change

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {

		currentBest := math.MaxInt32

		for _, coin := range coins {

			remainAmt := i - coin

			if remainAmt == 0 {
				// spot on.
				currentBest = 1
				break
			} else if remainAmt > 0 {

				tmp := dp[remainAmt]
				if tmp == -1 {
					continue // no answer
				}

				currentBest = min(currentBest, tmp+1)

			}
		}

		if currentBest == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = currentBest
		}

	}

	return dp[amount]
}
