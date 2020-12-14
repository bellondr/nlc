# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:

    def pruneTree(self, root: TreeNode) -> TreeNode:
        def dfs_sum(root):
            if root is None:
                return 0
            l = dfs_sum(root.left)
            r = dfs_sum(root.right)
            return root.val + l + r

        def dfs_delete(root):
            if root is None:
                return
            dfs_delete(root.left)
            dfs_delete(root.right)
            if dfs_sum(root.left) == 0:
                root.left = None
            if dfs_sum(root.right) == 0:
                root.right = None

        n = TreeNode(0)
        n.left = root
        dfs_delete(n)
        return n.left