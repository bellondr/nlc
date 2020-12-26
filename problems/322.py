class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
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
    
    def coinChange1(self, coins: List[int], amount: int) -> int:
        if amount == 0:
            return 0
        dp = [-1] * (amount + 1)
        for i in range(len(dp)):
            for c in coins:
                if i == c:
                    dp[i] = 1
                if i > c:
                    if dp[i - c] == -1:
                        continue
                    elif dp[i] != -1:
                        dp[i] = min(dp[i], dp[i - c] + 1)
                    else:
                        dp[i] = dp[i - c] + 1
        return dp[-1]