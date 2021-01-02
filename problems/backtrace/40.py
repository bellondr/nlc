# 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
#
# candidates 中的每个数字在每个组合中只能使用一次。
#
# 说明：
#
# 所有数字（包括目标数）都是正整数。
# 解集不能包含重复的组合。 
# 示例 1:
#
# 输入: candidates = [10,1,2,7,6,1,5], target = 8,
# 所求解集为:
# [
#   [1, 7],
#   [1, 2, 5],
#   [2, 6],
#   [1, 1, 6]
# ]
# 示例 2:
#
# 输入: candidates = [2,5,2,1,2], target = 5,
# 所求解集为:
# [
#   [1,2,2],
#   [5]
# ]

class Solution:
    def combinationSum2Dfs(self, candidates: [], target: int) -> []:
        candidates.sort()
        res = []

        def dfs(sub, target, path):
            if not sub or target < sub[0]:
                return

            for i in range(len(sub)):
                if target == sub[i]:
                    p = path + [sub[i]]
                    if p not in res:
                        res.append(p)
                    return
                if i > 0 and sub[i] == sub[i - 1]:
                    continue

                if target > sub[i]:
                    dfs(sub[i + 1:], target - sub[i], path + [sub[i]])
                else:
                    break

        dfs(candidates, target, [])
        return res

    def combinationSum2(self, candidates: [], target: int) -> []:
        candidates.sort()

        size = len(candidates)
        if size == 0:
            return []
        res = []
        self.find_path(candidates, target, res, [], 0, size)
        return res


    def find_path(self, candidates, target, res, path, left, right):
        if target == 0:
            res.append(path.copy())
        else:
            for i in range(left, right):
                left_num = target - candidates[i]
                if left_num < 0:
                    break
                if i > left and candidates[i] == candidates[i-1]:
                    continue
                path.append(candidates[i])
                self.find_path(candidates, left_num, res, path, i+1, right)
                path.pop()


if __name__ == '__main__':
    s = Solution()
    print(s.combinationSum2Dfs([10,1,2,7,6,1,5], 8))

