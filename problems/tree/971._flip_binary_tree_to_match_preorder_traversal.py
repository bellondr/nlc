# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
from problems.tree.tree import TreeNode


class Solution:
    def flipMatchVoyage(self, root: TreeNode, voyage):
        self.index = 0
        res = []
        if self.dfs(root, voyage, res):
            return res
        return [-1]

    def dfs(self, node, voyage, res):
        if node is None:
            return True
        if node.val != voyage[self.index]:
            return False
        self.index += 1

        if node.left and node.left.val == voyage[self.index]:
            return self.dfs(node.left, voyage, res) and self.dfs(node.right, voyage, res)
        elif node.right and node.right.val == voyage[self.index]:
            if node.left:
                res.append(node.val)
            return self.dfs(node.right, voyage, res) and self.dfs(node.left, voyage, res)

        return node.left is None and node.right is None
