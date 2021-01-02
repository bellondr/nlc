class Solution:
    def rob(self, nums):
        dp = [0] * len(nums)
        for i in range(0, len(nums)):
            if i == 0:
                dp[0] = nums[i]
            elif i == 1:
               dp[i] = max(dp[0], nums[1])
            else:
                dp[i] = max(dp[i-1], dp[i-2] + nums[i])

        return dp[-1]

if __name__ == '__main__':
    s = Solution()
    a = [1,2,3,1]
    print(s.rob(a))