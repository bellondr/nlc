class Solution:
    def numDecodings(self, s):
        if s is None or len(s) == 0:
            return 0
        dp = [0] * (len(s) + 1)
        dp[0] = 1
        dp[1] = 0 if s[0] == '0' else 1
        for i in range(2, len(s)+1):
            one = s[i-2]
            two = s[i-2:i]
            if '27' > two >= '10':
                dp[i] = dp[i-2]

            if '9' >= one >= '1':
                print(dp)
                dp[i] = dp[i] + dp[i-1]

        return dp[-1]


if __name__ == '__main__':
    s = Solution()
    print(s.numDecodings("1201234"))