package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	maxPnl := 0
	minEntryPrice := prices[0]

	for i := 1; i < len(prices); i++ {

		price := prices[i]
		if price < minEntryPrice { // have better entry price
			minEntryPrice = price // update best entry price
		} else {
			maxPnl = max(maxPnl, price-minEntryPrice) // calculate best pnl
		}
	}

	return maxPnl
}
