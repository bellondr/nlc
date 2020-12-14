

class TreeNode():
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

import copy

class Solution():

    def sumRootToLeaf(self, root):
        def help(root, val):
            if root is None:
                return 0
            val = val * 2 + root.val
            if root.left is None and root.right is None:
                return val
            return help(root.left, val) + help(root.right, val)

        return help(root, 0)


if __name__ == '__main__':
    root = TreeNode(1)
    root.left = TreeNode(0)
    root.right = TreeNode(1)
    root.left.left = TreeNode(0)
    root.left.right = TreeNode(1)
    root.right.left = TreeNode(0)
    root.right.right = TreeNode(1)
    s = Solution()
    print(s.sumRootToLeaf(root))
