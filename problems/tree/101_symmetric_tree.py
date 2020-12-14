# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution():
    def isSymmetric(self, root: TreeNode) -> bool:
        if root is None:
            return True

        def mirror(left, right):
            if left is None or right is None:
                return left == right
            return left.val == right.val and mirror(left.left, right.right) and mirror(left.right, right.left)

        return mirror(root.left, root.right)