class Solution:
    def maxProfit(self, prices):
        buy = [0] * 2
        sell = [0] * 2
        for i in range(len(prices)):
            if i == 0:
                buy[0] = 0 - prices[i]
            elif i == 1:
                buy[0] = max(buy[0], 0 - prices[i])
                sell[0] = max(sell[0], prices[i] + buy[0])
            else:
                buy[0] = max(buy[0], 0 - prices[i])
                sell[0] = max(sell[0], prices[i] + buy[0])

                buy[1] = max(buy[1], sell[0] - prices[i])
                sell[1] = max(sell[1], prices[i] + buy[1])

        return sell[1]