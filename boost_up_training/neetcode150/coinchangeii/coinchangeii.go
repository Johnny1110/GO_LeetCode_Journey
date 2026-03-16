package coinchangeii

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1) // i: coin can use, j: amount
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1 // whatever how many coin take, 0 amount always be 1.
	}

	for i := 1; i <= len(coins); i++ {

		coin := coins[i-1]

		for amt := 0; amt <= amount; amt++ {
			// skip this coin:
			dp[i][amt] = dp[i-1][amt]

			// take this coin:
			if amt >= coin {
				dp[i][amt] += dp[i][amt-coin]
			}
		}

	}

	return dp[len(coins)][amount]
}
