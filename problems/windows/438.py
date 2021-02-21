#!/usr/bin/env python
# -*- coding: utf-8 -*-


from collections import defaultdict
from collections import Counter

class Solution:
    def findAnagrams(self, s: str, p: str) -> [int]:
        res = []
        if len(s) < len(p):
            return res
        l = len(p)
        windows = defaultdict(lambda:0)

        c = Counter(p)

        def contains(c, w):
            for k in c:
                if k not in w or c[k] != w[k]:
                    return False
            return True


        for i in range(len(s)):
            windows[s[i]] += 1
            index = i - len(p) + 1

            if index >= 0:
                print(windows)
                if contains(c, windows):
                    res.append(index)
                windows[s[index]] -= 1

        return res


if __name__ == '__main__':
    s = Solution()
    print(s.findAnagrams('abab', 'ab'))
