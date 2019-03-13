class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) <= 1:
            return 0

        minimum = prices[0]
        profit = 0
        for i in range(1, len(prices)):
            if prices[i] > prices[i-1]:
                minimum = min(minimum, prices[i-1])
                profit = max(profit, prices[i]-minimum)
        return profit
