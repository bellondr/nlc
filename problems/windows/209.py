#!/usr/bin/env python
# -*- coding: utf-8 -*-

class Solution:
    def minSubArrayLen(self, s: int, nums: [int]) -> int:
        l = 0
        n = len(nums)
        tmp = 0
        res = []
        for r in range(n):
            tmp += nums[r]
            if tmp < s:
                continue

            while tmp >= s and l <= r:
                tmp = tmp - nums[l]
                if len(res) == 0:
                    res = nums[l:r+1]
                elif len(res) > r+1-l:
                    res = nums[l:r+1]
                l += 1
        
        return len(res)
                
if __name__ == '__main__':
    s = Solution()
    print(s.minSubArrayLen(11, [1,2,3,4,5]))
