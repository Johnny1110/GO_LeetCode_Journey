package best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_1(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	assert.Equal(t, 5, res)
}

func Test_2(t *testing.T) {
	prices := []int{7, 6, 4, 3, 1}
	res := maxProfit(prices)
	assert.Equal(t, 0, res)
}
