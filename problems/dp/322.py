import math
class Solution:
    def coinChange(self, coins, amount: int) -> int:
        n = len(coins)
        coins.sort(reverse=True)     # 先给硬币排序，降序
        self.res = float("inf")      # 定义最小的使用硬币的数量为self.res

        def dfs(index,target,count):   # 定义深度优先搜索算法
            coin = coins[index]
            if math.ceil(target/coin)+count>=self.res:
                return
            if target%coin==0:
                self.res = count+target//coin
            if index==n-1:return
            for j in range(target//coin,-1,-1):
                dfs(index+1,target-j*coin,count+j)

        dfs(0,amount,0)
        return int(self.res) if self.res!=float("inf") else -1


    def coinChangeDP(self, coins, amount: int) -> int:
        dp = [-1] * (amount + 1)
        dp[0] = 0
        for i in range(1, amount + 1):
            for c in coins:
                if c > i:
                    continue
                elif c == i:
                    dp[i] = 1
                elif dp[i - c] == -1:
                    continue
                elif dp[i] == -1:
                    dp[i] = dp[i - c] + 1
                else:
                    dp[i] = min(dp[i], dp[i - c] + 1)

        return dp[-1]