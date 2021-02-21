# 你将会得到一份单词表 words，一个字母表 letters （可能会有重复字母），以及每个字母对应的得分情况表 score。
#
# 请你帮忙计算玩家在单词拼写游戏中所能获得的「最高得分」：能够由 letters 里的字母拼写出的 任意 属于 words 单词子集中，分数最高的单词集合的得分。
#
# 单词拼写游戏的规则概述如下：
#
# 玩家需要用字母表 letters 里的字母来拼写单词表 words 中的单词。
# 可以只使用字母表 letters 中的部分字母，但是每个字母最多被使用一次。
# 单词表 words 中每个单词只能计分（使用）一次。
# 根据字母得分情况表score，字母 'a', 'b', 'c', ... , 'z' 对应的得分分别为 score[0], score[1], ..., score[25]。
# 本场游戏的「得分」是指：玩家所拼写出的单词集合里包含的所有字母的得分之和。
#  
#
# 示例 1：
#
# 输入：words = ["dog","cat","dad","good"], letters = ["a","a","c","d","d","d","g","o","o"], score = [1,0,9,5,0,0,3,0,0,0,0,0,0,0,2,0,0,0,0,0,0,0,0,0,0,0]
# 输出：23
# 解释：
# 字母得分为  a=1, c=9, d=5, g=3, o=2
# 使用给定的字母表 letters，我们可以拼写单词 "dad" (5+1+5)和 "good" (3+2+2+5)，得分为 23 。
# 而单词 "dad" 和 "dog" 只能得到 21 分。


class Solution:
    def maxScoreWords(self, words: [str], letters: [str], score: [int]) -> int:
        wordScores = [sum(score[ord(c) - ord('a')] for c in word) for word in words]
        import collections
        wordCounters = [collections.Counter(word) for word in words]

        self.res = 0

        def backtrace(start, cur, counter):
            if start > len(words):
                return
            self.res = max(self.res, cur)
            for j, w in enumerate(wordCounters[start:], start):
                if all(n <= counter.get(c, 0) for c, n in w.items()):
                    backtrace(j + 1, cur + wordScores[j], counter - w)

        backtrace(0, 0, collections.Counter(letters))
        return self.res