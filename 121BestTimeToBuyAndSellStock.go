package demo

//method one: violently calculate
func BestTimeToBuyAndSellStockViolence(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	indexValue, profit := 0, 0
	for i := 0; i < len(prices); i++ {
		indexValue = prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-indexValue > profit {
				profit = prices[j] - indexValue
			}
		}

	}
	return profit
}

//method two: it costs lower(lower price) or gets higher profit
func BestTimeToBuyAndSellStock(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	rawCost := 1 << 32
	profit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < rawCost {
			rawCost = prices[i]
		} else if prices[i]-rawCost > profit {
			profit = prices[i] - rawCost

		}

	}
	return profit
}
