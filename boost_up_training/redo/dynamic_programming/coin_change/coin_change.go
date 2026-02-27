package coin_change

import "math"

func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }

    dp := make([]int, amount+1)

    for i := 1; i < amount+1; i++ {
        
        tmp := math.MaxInt32

        // iterate through the coins
        for _, coin := range coins {
            remainAmount := i - coin

            if remainAmount == 0 {
                tmp = 1
                break // jack pock
            } else if remainAmount > 0 {

                if dp[remainAmount] > 0 {
                    tmp = min(1 + dp[remainAmount], tmp)
                }
            }
        }

        if tmp == math.MaxInt32 {
            dp[i] = -1
        } else {
             dp[i] = tmp
        }

    }

    return dp[amount]
}