package coin_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	//  Input: coins = [1,2,5], amount = 11
	//	Output: 3
	//	Explanation: 11 = 5 + 5 + 1

	coins := []int{1, 2, 5}
	assert.Equal(t, 3, coinChange(coins, 11))
}

func Test_2(t *testing.T) {
	//  Input: coins = [2], amount = 3
	//	Output: 0

	coins := []int{2}
	assert.Equal(t, -1, coinChange(coins, 3))
}

func Test_3(t *testing.T) {
	//  Input: coins = [1], amount = 0
	//	Output: 0

	coins := []int{2}
	assert.Equal(t, 0, coinChange(coins, 0))
}
