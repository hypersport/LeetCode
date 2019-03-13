func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minimum := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			minimum = min(minimum, prices[i-1])
			profit = max(profit, prices[i]-minimum)
		}
	}
	return profit
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
