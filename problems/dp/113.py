# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None

class Solution:
    def pathSum(self, root: TreeNode, sum: int) -> [[int]]:
        res = []

        def dfs(node, target, nums):
            if not node:
                return
            if not node.left and not node.right and target - node.val == 0:
                res.append(nums + [node.val])
                return
            dfs(node.left, target - node.val, nums + [node.val])
            dfs(node.right, target - node.val, nums + [node.val])

        dfs(root, sum, [])
        return res


if __name__ == '__main__':
    root = TreeNode(5)
    root.left = TreeNode(4)
    root.left.left = TreeNode(11)
    root.left.left.left = TreeNode(7)
    root.left.left.right = TreeNode(2)

    root.right = TreeNode(8)
    root.right.left = TreeNode(13)
    root.right.right = TreeNode(4)
    root.right.right.left = TreeNode(5)
    root.right.right.right = TreeNode(1)

    s = Solution()
    res = s.pathSum(root,22)
    print(res)

