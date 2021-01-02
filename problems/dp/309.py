class Solution:
    def maxProfit(self, prices):
        buy = [0] * len(prices)
        sell = [0] * len(prices)
        buy[0] = 0 - prices[0]
        buy[1] = max(0 - prices[0], 0 - prices[1])
        sell[0] = 0
        sell[1] = max(0, prices[1] - prices[0])
        print(buy)
        print(sell)
        for i in range(2, len(prices)):
            buy[i] = max(buy[i - 1], sell[i - 2] - prices[i])
            sell[i] = max(sell[i - 1], buy[i - 1] + prices[i])
            print(buy)
            print(sell)
        return max(buy[-1], sell[-1], 0)

if __name__ == '__main__':
    s = Solution()
    s.maxProfit([1,2,3,0,2])
