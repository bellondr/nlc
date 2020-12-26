class Solution:
    res = False
    def helper(self, s, wordDict):
        if len(s) == 0:
            self.res = True
            return
        for wd in wordDict:
            l = len(wd)
            if s[:l] == wd:
                self.helper(s[l:], wordDict)


    def wordBreakViolence(self, s, wordDict):
        self.helper(s, wordDict)
        return self.res

    def wordBreak(self, s, wordDict):
        dp = [False] * (len(s) + 1)
        dp[0] = True
        for i in range(len(s)+1):
            for wd in wordDict:
                if len(wd) <= i and dp[i-len(wd)]:
                    if s[i-len(wd):i] == wd:
                        dp[i] = True

        return dp[len(s)]


if __name__ == '__main__':
    s = Solution()
    print(s.wordBreak("leetcode", ["leet", "code"]))
    print(s.wordBreak("cars", ["car", "ca", "rs"]))
