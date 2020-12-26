# 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
#
# 注意:
#
# 每个数组中的元素不会超过 100
# 数组的大小不会超过 200
# 示例 1:
#
# 输入: [1, 5, 11, 5]
#
# 输出: true
#
# 解释: 数组可以分割成 [1, 5, 5] 和 [11].
#  
#
# 示例 2:
#
# 输入: [1, 2, 3, 5]
#
# 输出: false
#
# 解释: 数组不能分割成两个元素和相等的子集.


class Solution:
    def canPartition(self, nums) -> bool:
        if len(nums) == 0:
            return True
        total = sum(nums)
        if total % 2:
            return False
        half = int(total / 2)
        dp = [0] * (half + 1)
        dp[0] = 1

        for n in nums:
            if n == half:
                return True
            if n > half:
                continue
            r = range(half + 1)
            for i in reversed(r):
                if dp[i] == 1 and i + n <= half:
                    dp[i + n] = 1
        return dp[-1] == 1
