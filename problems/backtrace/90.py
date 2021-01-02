# 给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
#
# 说明：解集不能包含重复的子集。
#
# 示例:
#
# 输入: [1,2,2]
# 输出:
# [
#   [2],
#   [1],
#   [1,2,2],
#   [2,2],
#   [1,2],
#   []
# ]

class Solution:

    def subsetsWithDup(self, nums) -> []:
        res = []
        nums.sort()

        def dfs(index, tmp):
            if tmp not in res:
                res.append(tmp)
            for i in range(index, len(nums)):
                dfs(i+1, tmp+[nums[i]])
        dfs(0,[])
        return res

    def subsetsWithDup1(self, nums) -> []:
        self.res = []
        self.trace = []
        nums.sort()
        self.helper(nums, 0, len(nums))
        return self.res

    def helper(self, nums, left, right):
        if self.trace not in self.res:
            self.res.append(self.trace.copy())
        for i in range(left, right):
            self.trace.append(nums[i])
            self.helper(nums, i + 1, right)
            self.trace.pop()