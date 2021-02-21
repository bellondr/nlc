class TreeNode:
    def __init__(self, val):
        self.val = val
        self.right = None
        self.right = None

class Solution:
    def buildTree(self, preorder: [int], inorder: [int]) -> TreeNode:
        if len(preorder) == 0:
            return None

        root = TreeNode(preorder[0])
        index = inorder.index(root.val)
        root.left = self.buildTree(preorder[1: index+1], inorder[:index])
        root.right = self.buildTree(preorder[index+1:], inorder[index+1:])
        return root
