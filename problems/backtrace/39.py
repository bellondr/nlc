# 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
#
# candidates 中的数字可以无限制重复被选取。
#
# 说明：
#
# 所有数字（包括 target）都是正整数。
# 解集不能包含重复的组合。 
# 示例 1：
#
# 输入：candidates = [2,3,6,7], target = 7,
# 所求解集为：
# [
#   [7],
#   [2,2,3]
# ]
# 示例 2：
#
# 输入：candidates = [2,3,5], target = 8,
# 所求解集为：
# [
#   [2,2,2,2],
#   [2,3,3],
#   [3,5]
# ]
#  
#
class Solution:
    def combinationSumDFS(self, candidates: [], target: int) -> []:
        res = []

        def dfs(sub, target, path):
            if not sub or target < sub[0]:
                return
            for i in range(len(sub)):
                if target == sub[i]:
                    res.append(path + [sub[i]])
                elif target > sub[i]:
                    dfs(sub[i:], target - sub[i], path + [sub[i]])
                else:
                    break

        candidates.sort()
        dfs(candidates, target, [])
        return res

    def combinationSum(self, candidates: [], target: int) -> []:
        size = len(candidates)
        if size == 0:
            return []
        candidates.sort()
        res = []
        self.find_path(target, res, candidates, [], 0, size)
        return res


    def find_path(self, target, res, candidates, path, left, right):
        if target == 0:
            res.append(path.copy())
        else:
            for i in range(left, right):
                left_num = target - candidates[i]
                if left_num < 0:
                    break
                path.append(candidates[i])
                self.find_path(left_num, res, candidates, path, i, right)
                path.pop()


if __name__ == '__main__':
    s = Solution()
    print(s.combinationSumDFS([2,3,6,7], 7))
