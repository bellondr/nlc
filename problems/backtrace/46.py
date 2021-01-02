# 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
#
# 示例:
#
# 输入: [1,2,3]
# 输出:
# [
#   [1,2,3],
#   [1,3,2],
#   [2,1,3],
#   [2,3,1],
#   [3,1,2],
#   [3,2,1]
# ]

class Solution:
    def permuteDfs(self, nums) -> []:
        res = []

        def dfs(nums, tmp):
            if len(nums) == 0:
                res.append(tmp.copy())

            for i in range(len(nums)):
                p = tmp[:]
                p.append(i)
                left = nums[:]
                left.remove(i)
                dfs(left, p)

        dfs(nums, [])
        return res

    def permute(self, nums) -> []:
        res = []
        size = len(nums)
        def backtrace(tmp):
            if len(tmp) == size:
                res.append(tmp)
            else:
                for n in nums:
                    if n in tmp:
                        continue
                    else:
                        backtrace(tmp+[n])
        return res
