
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution():
    m = 0
    def diameterOfBinaryTree(self, root: TreeNode) -> int:
        def dfs(node):
            if node is None:
                return 0
            if node.left is None and node.right is None:
                return 1
            l = dfs(node.left)
            r = dfs(node.right)
            if l + r + 1 > self.m:
                self.m = l + r + 1
            return max(l, r) + 1
        dfs(root)
        return max(self.m - 1, 0)