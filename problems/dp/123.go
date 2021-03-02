package dp

import "math"

func maxProfit(prices []int) int {
	buy1, buy2 := math.MinInt64, math.MinInt64
	sell1, sell2 := 0, 0
	for _, v := range prices {
		if buy1 < 0 - v {
			buy1 = 0 - v
		}
		if sell1 < buy1 + v {
			sell1 = buy1 + v
		}
		if buy2 < sell1 - v {
			buy2 = sell1 - v
		}
		if sell2 < buy2 + v {
			sell2 = buy2 + v
		}
	}
	return sell2
}
