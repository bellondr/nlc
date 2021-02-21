class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        res = 0
        for i in range(32):
            cnt = 0
            bit = 1 << i
            for n in nums:
                cnt += bit & n
            if cnt % 3 != 0:
                res |= bit
        return res - 2**32 if res > 2**31-1 else res
