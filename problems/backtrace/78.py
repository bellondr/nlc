# 给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
#
# 说明：解集不能包含重复的子集。
#
# 示例:
#
# 输入: nums = [1,2,3]
# 输出:
# [
#   [3],
#   [1],
#   [2],
#   [1,2,3],
#   [1,3],
#   [2,3],
#   [1,2],
#   []
# ]

class Solution:

    def subsets(self, nums: [int]) -> [[int]]:
        self.res = []
        self.trace = []
        self.helper(nums, 0, len(nums))
        return self.res

    def helper(self, nums, left, right):
        self.res.append(self.trace.copy())
        for i in range(left, right):
            self.trace.append(nums[i])
            self.helper(nums, i + 1, right)
            self.trace.pop()

    def subsetsDfs(self, nums: [int]) -> [[int]]:
        res =[]

        def dfs(index, tmp):
            res.append(tmp)
            for i in range(index, len(nums)):
                dfs(i + 1, tmp + [nums[i]])

        dfs(0, [])
        return res
