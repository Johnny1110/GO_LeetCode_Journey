package coin_change

import "sort"

func coinChange(coins []int, amount int) int {

	// order coins desc.
	sort.Slice(coins, func(i, j int) bool { return coins[i] > coins[j] })

	return 0
}
